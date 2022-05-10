package viewmodel

import (
	"lugdroid/mealsScheduler/webapp/model"
)

type MealDetail struct {
	Title       string
	Meal        model.Meal
	Ingredients []model.Ingredient
	MealTypes   map[string]int
	Active      string
}

func NewMealDetail(meal model.Meal) MealDetail {
	return MealDetail{
		Title:       "Meal Scheduler - Meal Detail",
		Meal:        meal,
		Ingredients: model.GetAllIngredients(),
		MealTypes: map[string]int{
			"Disabled": 0,
			"Lunch":    1,
			"Dinner":   2,
			"Both":     3,
		},
		Active: "meals",
	}
}
