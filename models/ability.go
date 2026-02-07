package models

// Ability represents a card's special ability or attack
// @Description Special ability or attack that a card can perform
type Ability struct {
	Name             string `json:"name" example:"Thunder Jolt"`                                                         // Name of the ability
	Description      string `json:"description" example:"Flip a coin. If tails, Pokémon also does 10 damage to itself."` // Description of what the ability does
	Damage           int    `json:"damage" example:"30"`                                                                 // Base damage value
	DamageMultiplier string `json:"damageMultiplier" example:"x"`                                                        // Damage multiplier (e.g., 2x, 3x)
}
