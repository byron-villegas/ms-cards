package models

type Stat struct {
	Attack         int    `json:"attack"`
	Defense        int    `json:"defense"`
	Health         int    `json:"health"`
	Power          int    `json:"power"`
	Weakness       int    `json:"weakness"`
	WeaknessType   string `json:"weaknessType"`
	Resistance     int    `json:"resistance"`
	ResistanceType string `json:"resistanceType"`
	RetreatCost    int    `json:"retreatCost"`
}
