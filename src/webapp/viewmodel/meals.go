package viewmodel

import "lugdroid/mealsScheduler/webapp/model"

type Meals struct {
	Title  string
	Meals  []model.Meal
	Active string
}

func NewMeals(meals []model.Meal) Meals {
	return Meals{
		Title:  "Meal Scheduler - Meals",
		Meals:  meals,
		Active: "meals",
	}
}
