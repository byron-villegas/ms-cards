package models

// Stat represents the various stats of a card, such as attack, defense, health, power, weakness, resistance, and retreat cost
// @Description Various stats of a card, such as attack, defense, health, power, weakness, resistance, and retreat cost
type Stat struct {
	Attack         int    `json:"attack" example:"0"`             // Base attack value
	Defense        int    `json:"defense" example:"0"`            // Base defense value
	Health         int    `json:"health" example:"60"`            // Base health value
	Power          int    `json:"power" example:"0"`              // Base power value (for non-creature cards)
	Weakness       int    `json:"weakness" example:"0"`           // Weakness value (additional damage taken from certain types)
	WeaknessType   string `json:"weaknessType" example:"Rock"`    // Type of weakness (e.g., Fire, Water, Grass)
	Resistance     int    `json:"resistance" example:"0"`         // Resistance value (reduced damage taken from certain types)
	ResistanceType string `json:"resistanceType" example:"Water"` // Type of resistance (e.g., Fire, Water, Grass)
	RetreatCost    int    `json:"retreatCost" example:"1"`        // Cost to retreat the card from battle
}
