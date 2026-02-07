package models

// Effect represents a card's special effect or status condition
// @Description Special effect or status condition that a card can have
type Effect struct {
	Name        string `json:"name" example:"Paralysis"`
	Description string `json:"description" example:"Prevents the affected card from attacking during its next turn."`
}
