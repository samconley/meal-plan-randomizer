package main

import (
	"fmt"
	"meal-plan-randomizer/internal/model"
	"meal-plan-randomizer/internal/service"
	"time"
)

func main() {
	fmt.Println("Starting meal-plan-randomizer...")

	smsConfig := model.NewSmsConfig()
	smsMessageService := service.NewSmsMessageService(smsConfig)

	mealConfig, err := model.NewMealConfig()
	if err != nil {
		panic(err)
	}

	fileService := service.NewFileService(mealConfig)
	data := fileService.ReadMealsFromFile()

	numberOfMealsToSend := mealConfig.NumberOfMealsToSend

	if err != nil || numberOfMealsToSend > len(data.Meals) {
		fmt.Println("number of meals to send not correctly configured or too many")
		return
	}

	eligibleMeals := getEligibleMeals(data.Meals, 7)
	if len(eligibleMeals) < numberOfMealsToSend {
		fmt.Println("error: not enough meals to send")
		return
	}

	randomIndices := service.GetRandomIndices(len(eligibleMeals), numberOfMealsToSend)
	for _, idx := range randomIndices {
		meal := eligibleMeals[idx]
		smsMessageService.ComposeMessageFromMeal(meal)
		meal.LastUsed = time.Now().UTC()
	}

	fileService.SaveUpdatedMeals(data.Meals)

	fmt.Println("\nDone")
}

func getEligibleMeals(mealList []*model.Meal, lessRecentThanDays int) []*model.Meal {
	var eligibleMeals []*model.Meal
	for i := 0; i < len(mealList); i++ {
		lastEligibleDate := time.Now().AddDate(0, 0, (-1 * lessRecentThanDays))

		if mealList[i].LastUsed.Before(lastEligibleDate) {
			eligibleMeals = append(eligibleMeals, mealList[i])
		}
	}

	return eligibleMeals
}
