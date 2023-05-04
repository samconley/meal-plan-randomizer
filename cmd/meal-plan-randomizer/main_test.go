package main

import (
	"meal-plan-randomizer/internal/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMain_GetEligibleMeals(t *testing.T) {

	now := time.Now()

	tests := []struct {
		name                  string
		inputMealList         []*model.Meal
		limitDays             int
		expectedEligibleMeals []*model.Meal
	}{
		{
			name: `All meals are too recent, none are eligible`,
			inputMealList: []*model.Meal{
				{
					Name:        "Spaghetti and Meatballs",
					Ingredients: []string{},
					LastUsed:    now.Add(10 * time.Minute),
				},
				{
					Name:        "Pot Roast",
					Ingredients: []string{},
					LastUsed:    now.Add(15 * time.Minute),
				},
				{
					Name:        "Tacos",
					Ingredients: []string{},
					LastUsed:    now,
				},
			},
			limitDays:             21,
			expectedEligibleMeals: []*model.Meal{},
		},
		{
			name: `Two meals too recent, one remains eligible`,
			inputMealList: []*model.Meal{
				{
					Name:        "Spaghetti and Meatballs",
					Ingredients: []string{},
					LastUsed:    now.Add(10 * time.Minute),
				},
				{
					Name:        "Pot Roast",
					Ingredients: []string{},
					LastUsed:    now.AddDate(0, 0, (-15)),
				},
				{
					Name:        "Tacos",
					Ingredients: []string{},
					LastUsed:    now,
				},
			},
			limitDays: 7,
			expectedEligibleMeals: []*model.Meal{
				{
					Name:        "Pot Roast",
					Ingredients: []string{},
					LastUsed:    now.AddDate(0, 0, (-15)),
				},
			},
		},
		{
			name: `All meals are eligible`,
			inputMealList: []*model.Meal{
				{
					Name:        "Spaghetti and Meatballs",
					Ingredients: []string{},
					LastUsed:    now.Add(-700 * time.Hour),
				},
				{
					Name:        "Pot Roast",
					Ingredients: []string{},
					LastUsed:    now.Add(-800 * time.Hour),
				},
				{
					Name:        "Tacos",
					Ingredients: []string{},
					LastUsed:    now.Add(-600 * time.Hour),
				},
			},
			limitDays: 3,
			expectedEligibleMeals: []*model.Meal{
				{
					Name:        "Spaghetti and Meatballs",
					Ingredients: []string{},
					LastUsed:    now.Add(-700 * time.Hour),
				},
				{
					Name:        "Pot Roast",
					Ingredients: []string{},
					LastUsed:    now.Add(-800 * time.Hour),
				},
				{
					Name:        "Tacos",
					Ingredients: []string{},
					LastUsed:    now.Add(-600 * time.Hour),
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actualEligibleMeals := getEligibleMeals(tc.inputMealList, tc.limitDays)
			assert.ElementsMatch(t, tc.expectedEligibleMeals, actualEligibleMeals)
		})
	}
}
