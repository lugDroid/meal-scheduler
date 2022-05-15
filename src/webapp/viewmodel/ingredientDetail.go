package viewmodel

import "lugdroid/mealsScheduler/webapp/model"

type CategoryDetail struct {
	Title    string
	Category model.Category
	Active   string
}

func NewCategoryDetail(category model.Category) CategoryDetail {
	return CategoryDetail{
		Title:    "Meal Scheduler - Category Detail",
		Category: category,
		Active:   "categories",
	}
}
