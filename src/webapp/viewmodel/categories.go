package viewmodel

import "lugdroid/mealsScheduler/webapp/model"

type Categories struct {
	Title      string
	Categories []model.Category
	Active     string
}

func NewCategories(categories []model.Category) Categories {
	return Categories{
		Title:      "Meal Scheduler - Categories",
		Categories: categories,
		Active:     "categories",
	}
}
