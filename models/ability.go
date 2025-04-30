package models

type Ability struct {
	Name             string `json:"name"`
	Description      string `json:"description"`
	Damage           int    `json:"damage"`
	DamageMultiplier string `json:"damageMultiplier"`
}
