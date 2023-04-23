package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/smtp"
	"os"
	"strings"
	"time"
)

type Meal struct {
	Name        string   `json:"name"`
	Ingredients []string `json:"ingredients"`
	Sides       []string `json:"sides"`
}

type MealsDTO struct {
	Meals []*Meal `json:"meals"`
}

func main() {
	fmt.Println("Starting meal-plan-randomizer...")
	data := readMealsFromFile("etc/meals.json")
	smsSendRandomMeals(data.Meals, 3)
	fmt.Println("\nDone")
}

func readMealsFromFile(filePath string) *MealsDTO {
	content, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	var data MealsDTO
	if err = json.Unmarshal(content, &data); err != nil {
		panic(err)
	}

	return &data
}

func smsSendRandomMeals(mealList []*Meal, numberOfMealsToSend int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < numberOfMealsToSend; i++ {
		randIndex := r.Intn(len(mealList))
		meal := mealList[randIndex]
		composeMessageFromMeal(meal)
		mealList = removeMealAtIndex(mealList, randIndex)
	}
}

func removeMealAtIndex(meals []*Meal, index int) []*Meal {
	return append(meals[:index], meals[index+1:]...)
}

func composeMessageFromMeal(meal *Meal) {
	sb := &strings.Builder{}

	sb.WriteString("\n" + meal.Name)

	sb.WriteString("\nSupplies - ")
	for _, ingredient := range meal.Ingredients {
		sb.WriteString(ingredient + ",")
	}

	if len(meal.Sides) > 0 {
		sb.WriteString("\nSides - ")

		for _, side := range meal.Sides {
			sb.WriteString(side + ",")
		}
	}

	sendMsg(sb.String())
}

// TODO: DO NOT CHECK-IN!!! Set as github secrets and pull from environment.
func sendMsg(msg string) {
	from := os.Getenv("FROM_EMAIL")
	password := os.Getenv("EMAIL_PASSWORD")
	toList := strings.Split(os.Getenv("TO_LIST"), ",")
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_HOST_PORT")
	auth := smtp.PlainAuth("", from, password, host)
	if err := smtp.SendMail(host+":"+port, auth, from, toList, []byte(msg)); err != nil {
		fmt.Printf("error: %v", err)
	}
}
