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
	mealsController.mealsTemplate = templates["meals.html"]
	mealsController.mealDetailTemplate = templates["meal-detail.html"]
	ingredientsController.ingredientsTemplate = templates["ingredients.html"]
	ingredientsController.ingredientDetailTemplate = templates["ingredient-detail.html"]
	schedulesController.schedulesTemplate = templates["schedules.html"]
	schedulesController.scheduleDetailTemplate = templates["schedule-detail.html"]

	mealsController.registerRoutes()
	ingredientsController.registerRoutes()
	schedulesController.registerRoutes()

	http.Handle("/img/", http.FileServer(http.Dir("../../public")))
	http.Handle("/css/", http.FileServer(http.Dir("../../public")))
}
