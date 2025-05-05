package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	ServerPath        string
	ServerPort        string
	SecretKey         string
	JWTSecretKey      string
	JWTAccessTokenExp int
	MongoUser         string
	MongoPassword     string
	MongoHost         string
	MongoDB           string
	MongoURI          string
	CORSOrigin        string
	CORSHeaders       string
}

var AppConfig *Config

func init() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: .env file not found, using default or system environment variables")
	}

	// Initialize the AppConfig struct
	AppConfig = &Config{
		ServerPath:        "/api",
		ServerPort:        "8080",
		SecretKey:         os.Getenv("SECRET_KEY"),
		JWTSecretKey:      os.Getenv("JWT_SECRET_KEY"),
		JWTAccessTokenExp: 3600,
		MongoUser:         os.Getenv("MONGO_USER"),
		MongoPassword:     os.Getenv("MONGO_PASSWORD"),
		MongoHost:         os.Getenv("MONGO_HOST"),
		MongoDB:           os.Getenv("MONGO_DB"),
		MongoURI:          fmt.Sprintf("mongodb+srv://%s:%s@%s/%s", os.Getenv("MONGO_USER"), os.Getenv("MONGO_PASSWORD"), os.Getenv("MONGO_HOST"), os.Getenv("MONGO_DB")),
		CORSOrigin:        os.Getenv("CORS_ORIGIN"),
		CORSHeaders:       "Content-Type",
	}
}

func ShowBanner() {
	// Get the content of the banner file
	bannerFile, err := os.ReadFile("banner.txt")
	if err != nil {
		log.Fatalf("Error reading banner file: %v", err)
	}

	// Replace the placeholders in the banner with the actual values
	bannerLog := string(bannerFile)
	bannerLog = strings.ReplaceAll(bannerLog, "package.name", "ms-cards")
	bannerLog = strings.ReplaceAll(bannerLog, "package.version", "1.0.0")
	bannerLog = strings.ReplaceAll(bannerLog, "go.version", strings.Replace(runtime.Version(), "go", "", 1))
	bannerLog = strings.ReplaceAll(bannerLog, "gin.version", gin.Version)
	bannerLog = strings.ReplaceAll(bannerLog, "server.path", AppConfig.ServerPath)
	bannerLog = strings.ReplaceAll(bannerLog, "server.port", AppConfig.ServerPort)

	// Print the banner
	fmt.Println(bannerLog)
}

func Connect() (*mongo.Database, error) {
	// Set the MongoDB connection string
	mongoUrl := fmt.Sprintf("mongodb+srv://%s:%s@%s/%s",
		AppConfig.MongoUser,
		AppConfig.MongoPassword,
		AppConfig.MongoHost,
		AppConfig.MongoDB,
	)

	// Config the options for the MongoDB client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoUrl).SetServerAPIOptions(serverAPI)

	ctx, _ := context.WithTimeout(context.Background(), time.Duration(5000)*time.Second)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	// Verify the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("Error pinging MongoDB: %v", err)
	}

	log.Println("Successfully connected to MongoDB")

	// Return the database connection
	return client.Database("db_cards"), err
}
