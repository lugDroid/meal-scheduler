package controller

import (
	"html/template"
	"net/http"
)

var (
	mealsController       meals
	ingredientsController ingredients
	schedulesController   schedules
)

func StartUp(templates map[string]*template.Template) {
	mealsController.listTemplate = templates["meals.html"]
	mealsController.detailTemplate = templates["meal-detail.html"]
	mealsController.deleteTemplate = templates["delete.html"]

	ingredientsController.listTemplate = templates["ingredients.html"]
	ingredientsController.detailTemplate = templates["ingredient-detail.html"]
	ingredientsController.deleteTemplate = templates["delete.html"]

	schedulesController.listTemplate = templates["schedules.html"]
	schedulesController.detailTemplate = templates["schedule-detail.html"]
	schedulesController.deleteTemplate = templates["delete.html"]

	mealsController.registerRoutes()
	ingredientsController.registerRoutes()
	schedulesController.registerRoutes()

	http.Handle("/img/", http.FileServer(http.Dir("../../public")))
	http.Handle("/css/", http.FileServer(http.Dir("../../public")))
}
