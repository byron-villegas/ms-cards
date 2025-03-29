package models

type Card struct {
	ID                  string `json:"id" bson:"_id,omitempty"`
	Name                string `json:"name"`
	NameFont            string `json:"nameFont"`
	Number              int    `json:"number"`
	Cost                Cost   `json:"cost"`
	Strength            int    `json:"strength"`
	Health              int    `json:"health"`
	DamageAndHealthType string `json:"damageAndHealthType"`
	Rarity              string `json:"rarity"`
	Image               string `json:"image"`
	Saga                string `json:"saga"`
	Type                string `json:"type"`
	Expansion           string `json:"expansion"`
	Text                string `json:"text"`
	AdditionalText      string `json:"additionalText"`
	Artist              string `json:"artist"`
	Count               int    `json:"count"`
}
