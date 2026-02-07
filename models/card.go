package models

// Card represents a trading card with various attributes, stats, abilities, effects, and set information
// @Description Trading card with various attributes, stats, abilities, effects, and set information
type Card struct {
	ID          string    `json:"id" bson:"_id,omitempty" example:"680d7ea3abc6e3839bd4adb5"`                                                                    // Unique identifier for the card
	Name        string    `json:"name" example:"Pikachu"`                                                                                                        // Name of the card
	ImageUrl    string    `json:"imageUrl" example:"Pikachu-005-025.webp"`                                                                                       // URL to the card's image
	Number      int       `json:"number" example:"5"`                                                                                                            // Card number in the set
	Game        string    `json:"game" example:"Pokemon"`                                                                                                        // Game the card belongs to
	Type        string    `json:"type" example:"Lightning"`                                                                                                      // Type of the card
	Description string    `json:"description" example:"It has small electric sacs on both its cheeks. If threatened, it looses electric charges from the sacs."` // Description of the card
	Attributes  Attribute `json:"attributes"`                                                                                                                    // Various attributes of the card
	Stats       Stat      `json:"stats"`
	Abilities   []Ability `json:"abilities"`
	Effects     []Effect  `json:"effects"`
	Set         Set       `json:"set"`
	Artist      string    `json:"artist" example:"Mitsuhiro Arita"`
	Count       int       `json:"count" example:"1"` // Count of this card in the collection
}
