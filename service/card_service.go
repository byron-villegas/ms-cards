package service

import (
	"fe-cards/models"
	"fe-cards/repository"
)

type CardService struct{}

var cardRepository = repository.CardRepository{}

func (c CardService) GetCards() []models.Card {
	cards := cardRepository.GetCards()

	return cards
}

func (c CardService) GetCardByID(id string) models.Card {
	card := cardRepository.GetCardByID(id)

	return card
}
