package viewmodel

import "lugdroid/mealsScheduler/webapp/model"

type Ingredients struct {
	Title       string
	Ingredients []model.Ingredient
	Active      string
}

func NewIngredients(ingredients []model.Ingredient) Ingredients {
	return Ingredients{
		Title:       "Meal Scheduler - Ingredients",
		Ingredients: ingredients,
		Active:      "ingredients",
	}
}
