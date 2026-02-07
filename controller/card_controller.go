package controller

import (
	"ms-cards/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CardController struct{}

var cardService = service.CardService{}

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

// GetCardByID godoc
// @Summary Get a card by ID
// @Description Get a card by its ID
// @Tags cards
// @Accept json
// @Produce json
// @Param id path string true "Card ID"
// @Success 200 {object} models.Card
// @Failure 404 {object} map[string]string
// @Router /cards/{id} [get]
func (ec *CardController) GetCardByID(c *gin.Context) {
	id := c.Param("id")
	card := cardService.GetCardByID(id)

	c.JSON(http.StatusOK, card)
}
