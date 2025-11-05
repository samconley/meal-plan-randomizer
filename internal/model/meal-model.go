package model

import "time"

type Meal struct {
	Name        string    `json:"name"`
	Ingredients []string  `json:"ingredients"`
	Sides       []string  `json:"sides"`
	LastUsed    time.Time `json:"lastUsed"`
	Struck 		boolean	  `json:"struck"`
}

type MealsDTO struct {
	Meals []*Meal `json:"meals"`
}
