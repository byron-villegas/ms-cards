package controller

import (
	"fe-cards/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CardController struct{}

var cardService = service.CardService{}

func (ec *CardController) GetCards(c *gin.Context) {
	saga := c.Query("saga")

	if saga != "" {
		cards := cardService.GetCardsBySaga(saga)
		c.JSON(http.StatusOK, cards)
		return
	}

	cards := cardService.GetCards()

	c.JSON(http.StatusOK, cards)
}

func (ec *CardController) GetCardByID(c *gin.Context) {
	id := c.Param("id")
	card := cardService.GetCardByID(id)

	c.JSON(http.StatusOK, card)
}
