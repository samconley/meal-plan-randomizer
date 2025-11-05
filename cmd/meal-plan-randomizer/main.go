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

	if numberOfMealsToSend > len(data.Meals) {
		log.Println("number of meals to send too many")
		return
	}

	eligibleMeals := getEligibleMeals(data.Meals)
	if len(eligibleMeals) < numberOfMealsToSend {
		for _, meal := range data.Meals {
			meal.Struck = false
		}
		eligibleMeals = getEligibleMeals(data.Meals)
	}

	preamble := "Behold! The meals you shall consume this week... "
	smsMessageService.SendMsg(preamble)

	randomIndices := service.GetRandomIndices(len(eligibleMeals), numberOfMealsToSend)
	for _, idx := range randomIndices {
		meal := eligibleMeals[idx]
		smsMessageService.ComposeMessageFromMeal(meal)
		meal.LastUsed = time.Now().UTC()
		meal.Struck = true
	}

	fileService.SaveUpdatedMeals(data.Meals)

	log.Println("\nDone")
}

func getEligibleMeals(mealList []*model.Meal) []*model.Meal {
	var unstruckMeals []*model.Meal
	for _, meal := range mealList {
		if !meal.Struck {
			unstruckMeals = append(unstruckMeals, meal)
		}
	}

	return unstruckMeals
}
