package viewmodel

import "lugdroid/mealsScheduler/webapp/model"

type IngredientDetail struct {
	Title      string
	Ingredient model.Ingredient
	Active     string
}

func NewIngredientDetail(ingredient model.Ingredient) IngredientDetail {
	return IngredientDetail{
		Title:      "Meal Scheduler - Ingredient Detail",
		Ingredient: ingredient,
		Active:     "ingredients",
	}
}
