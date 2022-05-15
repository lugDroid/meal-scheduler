package controller

import (
	"html/template"
	"net/http"
)

var (
	mealsController      meals
	categoriesController categories
	schedulesController  schedules
)

func StartUp(templates map[string]*template.Template) {
	mealsController.listTemplate = templates["meals.html"]
	mealsController.detailTemplate = templates["meal-detail.html"]
	mealsController.deleteTemplate = templates["delete.html"]

	categoriesController.listTemplate = templates["categories.html"]
	categoriesController.detailTemplate = templates["category-detail.html"]
	categoriesController.deleteTemplate = templates["delete.html"]

	schedulesController.listTemplate = templates["schedules.html"]
	schedulesController.detailTemplate = templates["schedule-detail.html"]
	schedulesController.deleteTemplate = templates["delete.html"]

	mealsController.registerRoutes()
	categoriesController.registerRoutes()
	schedulesController.registerRoutes()

	http.Handle("/img/", http.FileServer(http.Dir("../../public")))
	http.Handle("/css/", http.FileServer(http.Dir("../../public")))
}
