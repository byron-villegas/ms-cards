package models

// Attribute represents the various attributes of a card, such as element, rarity, mana cost, color, and type-specific attributes.
// @Description Various attributes of a card, such as element, rarity, mana cost, color, and type-specific attributes
type Attribute struct {
	Element      string `json:"element" example:""`      // Elemental type of the card (e.g., Fire, Water, Earth)
	Rarity       string `json:"rarity" example:"Basic"`  // Rarity level of the card (e.g., Common, Uncommon, Rare, Epic)
	ManaCost     int    `json:"manaCost" example:"0"`    // Mana cost required to play the card
	Color        string `json:"color" example:""`        // Color of the card (e.g., Red, Blue, Green)
	TypeSpecific string `json:"typeSpecific" example:""` // Type-specific attribute (e.g., Flying for creatures, Spell for spells)
}
