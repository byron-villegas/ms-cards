package models

type Card struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	Name        string    `json:"name"`
	ImageUrl    string    `json:"imageUrl"`
	Number      int       `json:"number"`
	Game        string    `json:"game"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
	Attributes  Attribute `json:"attributes"`
	Stats       Stat      `json:"stats"`
	Abilities   []Ability `json:"abilities"`
	Effects     []Effect  `json:"effects"`
	Set         Set       `json:"set"`
	Artist      string    `json:"artist"`
	Count       int       `json:"count"`
}
