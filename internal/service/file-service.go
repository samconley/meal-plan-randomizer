package service

import (
	"encoding/json"
	"log"
	"meal-plan-randomizer/internal/model"
	"os"
)

type FileService struct {
	MealConfig *model.MealConfig
}

func NewFileService(cfg *model.MealConfig) *FileService {
	return &FileService{
		MealConfig: cfg,
	}
}

func (f *FileService) ReadMealsFromFile() *model.MealsDTO {
	content, err := os.ReadFile(f.MealConfig.SourceFileDir + "/meals.json")

	if err != nil {
		log.Panic(err)
	}

	var data model.MealsDTO
	if err = json.Unmarshal(content, &data); err != nil {
		log.Panic(err)
	}

	return &data
}

func (f *FileService) SaveUpdatedMeals(mealList []*model.Meal) {
	mealsJson, err := json.MarshalIndent(&model.MealsDTO{
		Meals: mealList,
	}, "", "    ")

	if err != nil {
		log.Println("error serializing updated meal list")
		return
	}

	out, err := os.Create(f.MealConfig.SourceFileDir + "/updated_meals.json")
	if err != nil {
		log.Println("error creating updated meals output file")
		return
	}

	defer out.Close()

	out.Write(mealsJson)
}
