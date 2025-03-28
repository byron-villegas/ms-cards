package routes

import (
	"fe-cards/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(routerGroup *gin.RouterGroup) {
	cardController := controller.CardController{}

	routerGroup.GET("/cards", cardController.GetCards)
	routerGroup.GET("/cards/:id", cardController.GetCardByID)
}
