package controller

import (
	"html/template"
	"lugdroid/mealsScheduler/webapp/model"
	"net/http"
)

var (
	mealsController      meals
	categoriesController categories
	schedulesController  schedules
)

func StartUp(templates map[string]*template.Template, storage model.Storage) {
	mealsController.listTemplate = templates["meals.html"]
	mealsController.detailTemplate = templates["meal-detail.html"]
	mealsController.deleteTemplate = templates["delete.html"]

	categoriesController.listTemplate = templates["categories.html"]
	categoriesController.detailTemplate = templates["category-detail.html"]
	categoriesController.deleteTemplate = templates["delete.html"]

	schedulesController.listTemplate = templates["schedules.html"]
	schedulesController.detailTemplate = templates["schedule-detail.html"]
	schedulesController.deleteTemplate = templates["delete.html"]

	mealsController.storage = storage
	categoriesController.storage = storage
	schedulesController.storage = storage

	mealsController.registerRoutes()
	categoriesController.registerRoutes()
	schedulesController.registerRoutes()

	http.Handle("/img/", http.FileServer(http.Dir("../../public")))
	http.Handle("/css/", http.FileServer(http.Dir("../../public")))
}
