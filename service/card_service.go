package service

import (
	"ms-cards/models"
	"ms-cards/repository"
	"strings"
)

type CardService struct{}

var cardRepository = repository.CardRepository{}

func (c CardService) GetCards() []models.Card {
	cards := cardRepository.GetCards()

	return cards
}

func (c CardService) GetCardsByGame(game string) []models.Card {

	if strings.Contains(strings.ToLower(game), "digimon") {
		game = "Digimon"
	}

	if strings.Contains(strings.ToLower(game), "pokemon") {
		game = "Pokemon"
	}

	if strings.Contains(strings.ToLower(game), "magic") {
		game = "Magic The Gathering"
	}

	if strings.Contains(strings.ToLower(game), "yu-gi-oh") {
		game = "Yu-Gi-Oh"
	}

	if strings.Contains(strings.ToLower(game), "mitos") {
		game = "Mitos y Leyendas"
	}

	cards := cardRepository.GetCardsByGame(game)

	return cards
}

func (c CardService) GetCardByID(id string) models.Card {
	card := cardRepository.GetCardByID(id)

	return card
}
