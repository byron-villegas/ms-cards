package service

import (
	"fe-cards/models"
	"fe-cards/repository"
	"strings"
)

type CardService struct{}

var cardRepository = repository.CardRepository{}

func (c CardService) GetCards() []models.Card {
	cards := cardRepository.GetCards()

	return cards
}

func (c CardService) GetCardsBySaga(saga string) []models.Card {

	if strings.Contains(strings.ToLower(saga), "digimon") {
		saga = "Digimon"
	}

	if strings.Contains(strings.ToLower(saga), "pokemon") {
		saga = "Pokemon"
	}

	if strings.Contains(strings.ToLower(saga), "magic") {
		saga = "Magic The Gathering"
	}

	if strings.Contains(strings.ToLower(saga), "yu-gi-oh") {
		saga = "Yu-Gi-Oh"
	}

	if strings.Contains(strings.ToLower(saga), "mitos") {
		saga = "Mitos y Leyendas"
	}

	cards := cardRepository.GetCardsBySaga(saga)

	return cards
}

func (c CardService) GetCardByID(id string) models.Card {
	card := cardRepository.GetCardByID(id)

	return card
}
