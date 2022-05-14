package controller

import (
	"html/template"
	"log"
	"lugdroid/mealsScheduler/webapp/model"
	"lugdroid/mealsScheduler/webapp/viewmodel"
	"net/http"
	"regexp"
	"strconv"
)

type meals struct {
	listTemplate   *template.Template
	detailTemplate *template.Template
	deleteTemplate *template.Template
}

func (m meals) registerRoutes() {
	http.HandleFunc("/meals", m.handleMeals)
	http.HandleFunc("/", m.handleMeals)
	http.HandleFunc("/meals/", m.handleMeals)
}

func (m meals) handleMeals(w http.ResponseWriter, r *http.Request) {
	idPattern, _ := regexp.Compile(`/meals/(\d+)`)
	idMatches := idPattern.FindStringSubmatch(r.URL.Path)
	if len(idMatches) > 0 {
		mealId, _ := strconv.Atoi(idMatches[1])
		m.handleDetail(w, r, mealId)
		return
	}

	newPattern, _ := regexp.Compile(`/meals/new$`)
	newMatches := newPattern.FindStringSubmatch(r.URL.Path)
	if len(newMatches) > 0 {
		m.handleNew(w, r)
		return
	}

	deletePattern, _ := regexp.Compile(`/meals/delete/(\d+)`)
	deleteMatches := deletePattern.FindStringSubmatch(r.URL.Path)
	if len(deleteMatches) > 0 {
		mealId, _ := strconv.Atoi(deleteMatches[1])
		m.handleDelete(w, r, mealId)
		return
	}

	meals := model.GetAllMeals()
	vm := viewmodel.NewMeals(meals)
	err := m.listTemplate.Execute(w, vm)
	if err != nil {
		log.Println("Could not execute template", m.listTemplate.Name(), err)
	}
}

func (m meals) handleDetail(w http.ResponseWriter, r *http.Request, mealId int) {
	meal := model.GetMealById(mealId)

	if r.Method == http.MethodPost {
		parseMealData(&meal, r)
		model.UpdateMeal(meal)
		http.Redirect(w, r, "/meals", http.StatusTemporaryRedirect)
	}

	vm := viewmodel.NewMealDetail(meal)
	err := m.detailTemplate.Execute(w, vm)
	if err != nil {
		log.Println("Could not execute template", m.detailTemplate.Name(), err)
	}
}

func (m meals) handleNew(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var newMeal model.Meal
		parseMealData(&newMeal, r)
		model.AddMeal(newMeal)
		http.Redirect(w, r, "/meals", http.StatusTemporaryRedirect)
	}
	vm := viewmodel.NewMealDetail(model.Meal{})
	err := m.detailTemplate.Execute(w, vm)
	if err != nil {
		log.Println("Could not execute template", m.detailTemplate.Name(), err)
	}
}

func (m meals) handleDelete(w http.ResponseWriter, r *http.Request, mealId int) {
	if r.Method == http.MethodPost {
		model.RemoveMeal(mealId)
		http.Redirect(w, r, "/meals", http.StatusTemporaryRedirect)
	}

	vm := viewmodel.NewDeleteViewModel("meal", model.GetMealById(mealId).Name, "/meals")
	vm.Active = "schedules"
	err := m.deleteTemplate.Execute(w, vm)
	if err != nil {
		log.Println("Could not execute template", m.deleteTemplate.Name(), err)
	}
}

func parseMealData(meal *model.Meal, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("Could not parse mealDetail form", err)
	}

	meal.Name = r.Form.Get("meal-name")
	meal.Description = r.Form.Get("meal-description")

	ingredientId, err := strconv.Atoi(r.Form.Get("ingredients"))
	if err != nil {
		log.Println("Could not parse ingredientId", err)
	}
	meal.MainIngredient = model.GetIngredientById(ingredientId)

	mealServings, err := strconv.Atoi(r.Form.Get("meal-servings"))
	if err != nil {
		log.Println("Could not parse mealServings", err)
	}
	meal.Servings = mealServings

	mealType, err := strconv.Atoi(r.Form.Get("meal-type"))
	if err != nil {
		log.Println("Could not parse mealType", err)
	}
	meal.Type = model.MealType(mealType)
}
