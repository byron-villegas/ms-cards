package models

type Attribute struct {
	Element      string `json:"element"`
	Rarity       string `json:"rarity"`
	ManaCost     int    `json:"manaCost"`
	Color        string `json:"color"`
	TypeSpecific string `json:"typeSpecific"`
}
