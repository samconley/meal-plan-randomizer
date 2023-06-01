package main

import (
	"log"
	"meal-plan-randomizer/internal/model"
	"meal-plan-randomizer/internal/service"
	"time"
)

func main() {
	log.Println("Starting meal-plan-randomizer...")

	smsConfig := model.NewSmsConfig()
	smsMessageService := service.NewSmsMessageService(smsConfig)

	mealConfig, err := model.NewMealConfig()
	if err != nil {
		log.Panic(err)
	}

	fileService := service.NewFileService(mealConfig)
	data := fileService.ReadMealsFromFile()

	numberOfMealsToSend := mealConfig.NumberOfMealsToSend

	if err != nil || numberOfMealsToSend > len(data.Meals) {
		log.Println("number of meals to send not correctly configured or too many")
		return
	}

	eligibleMeals := getEligibleMeals(data.Meals, mealConfig.LessRecentThanDays)
	if len(eligibleMeals) < numberOfMealsToSend {
		log.Println("error: not enough meals to send")
		return
	}

	randomIndices := service.GetRandomIndices(len(eligibleMeals), numberOfMealsToSend)
	for _, idx := range randomIndices {
		meal := eligibleMeals[idx]
		smsMessageService.ComposeMessageFromMeal(meal)
		meal.LastUsed = time.Now().UTC()
	}

	fileService.SaveUpdatedMeals(data.Meals)

	log.Println("\nDone")
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
