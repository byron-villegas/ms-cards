package repository

import (
	"context"
	"fe-cards/config"
	"fe-cards/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var mongoDatabase *mongo.Database
var err error

func init() {
	mongoDatabase, err = config.Connect()
	if err != nil {
		panic(err)
	}
}

type CardRepository struct{}

func (c CardRepository) GetCards() []models.Card {
	cursor, err := mongoDatabase.Collection("cards").Find(context.TODO(), bson.D{{}})

	if err != nil {
		return []models.Card{}
	}

	// Map results
	var bsonCards []bson.M
	if err = cursor.All(context.TODO(), &bsonCards); err != nil {
		return []models.Card{}
	}

	// Convert bson.M to models.Card
	var cards []models.Card
	for _, bsonCard := range bsonCards {
		var card models.Card
		bsonBytes, _ := bson.Marshal(bsonCard)
		_ = bson.Unmarshal(bsonBytes, &card)
		cards = append(cards, card)
	}

	return cards
}

func (c CardRepository) GetCardsBySaga(saga string) []models.Card {
	cursor, err := mongoDatabase.Collection("cards").Find(context.TODO(), bson.D{{Key: "saga", Value: saga}})

	if err != nil {
		return []models.Card{}
	}

	// Map results
	var bsonCards []bson.M
	if err = cursor.All(context.TODO(), &bsonCards); err != nil {
		return []models.Card{}
	}

	// Convert bson.M to models.Card
	var cards []models.Card
	for _, bsonCard := range bsonCards {
		var card models.Card
		bsonBytes, _ := bson.Marshal(bsonCard)
		_ = bson.Unmarshal(bsonBytes, &card)
		cards = append(cards, card)
	}

	if cards == nil {
		return []models.Card{}
	}

	return cards
}

func (c CardRepository) GetCardByID(id string) models.Card {
	var card models.Card

	// Convert the string ID to ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		// Return an empty card if the ID is invalid
		return models.Card{}
	}

	// Query using the ObjectID
	err = mongoDatabase.Collection("cards").FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&card)
	if err != nil {
		return models.Card{}
	}

	return card
}
