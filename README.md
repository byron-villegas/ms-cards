# ms-cards

Microservice for my cards collection.

## Instalar paquetes
Se debe ejecutar el siguiente comando

```shell
go mod tidy
```

## Ejecutar aplicacion
Se debe ejecutar el siguiente comando

```shell
go run main.go
```

## Swagger
### Instalación
Debemos instalar swaggo/swag con el siguiente comando

```shell
go install github.com/swaggo/swag/cmd/swag@latest
```

### Agregar Dependencias
Debemos agregar las siguientes dependencias

```shell
go get github.com/swaggo/gin-swagger
go get github.com/swaggo/files
```

### Documentar Aplicacion
```go
// @title           MS Cards API
// @version         1.0
// @description     Microservice for managing trading cards
// @host            localhost:8080
// @BasePath        /api
// @schemes         http https
func main() {
}
```

### Documentar Endpoints
Debemos documentar los endpoints con comentarios ejemplo

```go
// GetCards godoc
// @Summary Get all cards
// @Description Get all cards, optionally filtered by game
// @Tags cards
// @Accept json
// @Produce json
// @Param game query string false "Game name to filter cards"
// @Success 200 {array} models.Card
// @Failure 500 {object} map[string]string
// @Router /cards [get]
func (ec *CardController) GetCards(c *gin.Context) {
	game := c.Query("game")

	if game != "" {
		cards := cardService.GetCardsByGame(game)
		c.JSON(http.StatusOK, cards)
		return
	}

	cards := cardService.GetCards()

	c.JSON(http.StatusOK, cards)
}
```

### Documentar DTO/Models

```go
// Ability represents a card's special ability or attack
// @Description Special ability or attack that a card can perform
type Ability struct {
	Name             string `json:"name" example:"Thunder Jolt"`                                                         // Name of the ability
	Description      string `json:"description" example:"Flip a coin. If tails, Pokémon also does 10 damage to itself."` // Description of what the ability does
	Damage           int    `json:"damage" example:"30"`                                                                 // Base damage value
	DamageMultiplier string `json:"damageMultiplier" example:"x"`                                                        // Damage multiplier (e.g., 2x, 3x)
}
```

Swag generará la documentacion de swagger en base a esos comentarios

### Generar Archivo Swagger
Para generar el archivo swagger simplemente ejecutamos el siguiente comando

```shell
swag init
```

Esto nos generará los archivos de swagger en la carpeta **/docs** con la siguiente estructura

/docs
├── docs.go
├── swagger.json
└── swagger.yml

### Configurar Swagger UI
Para configurar Swagger UI simplemente agregamos el siguiente codigo al archivo **main.go**

```go
import (
    "github.com/swaggo/gin-swagger"
    "github.com/swaggo/files"
    _ "ms-cards/docs"
)

r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
```

Quedando así

```go
package main

import (
	"ms-cards/config"
	"ms-cards/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "ms-cards/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           MS Cards API
// @version         1.0
// @description     Microservice for managing trading cards
// @host            localhost:8080
// @BasePath        /api
// @schemes         http https
func main() {
	// Show banner with application information
	config.ShowBanner()

	// Create a server instance
	r := gin.Default()

	// Config CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{config.AppConfig.CORSOrigin},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{config.AppConfig.CORSHeaders},
		AllowCredentials: true,
	}))

	// Config server path
	routerGroup := r.Group(config.AppConfig.ServerPath)

	// Config routes for the server
	routes.SetupRoutes(routerGroup)

	// Swagger documentation route
	routerGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start server with port
	r.Run(":" + config.AppConfig.ServerPort)
}
```

Cuando ejecutemos a la aplicacion debemos entrar a la pagina /swagger/index.html