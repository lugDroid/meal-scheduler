package controller

import (
	"html/template"
	"net/http"
)

var (
	mealsController       meals
	ingredientsController ingredients
)

func StartUp(templates map[string]*template.Template) {
	mealsController.mealsTemplate = templates["meals.html"]
	mealsController.mealDetailTemplate = templates["meal-detail.html"]
	ingredientsController.ingredientsTemplate = templates["ingredients.html"]
	ingredientsController.ingredientDetailTemplate = templates["ingredient-detail.html"]

	mealsController.registerRoutes()
	ingredientsController.registerRoutes()

	http.Handle("/img/", http.FileServer(http.Dir("../../public")))
	http.Handle("/css/", http.FileServer(http.Dir("../../public")))
}
