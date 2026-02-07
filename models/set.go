package models

// Set represents the set information of a card, including the set name and release year
// @Description Set information of a card, including the set name and release year
type Set struct {
	Name        string `json:"name" example:"Celebrations 25th Anniversary"`
	ReleaseYear int    `json:"releaseYear" example:"2021"`
}
