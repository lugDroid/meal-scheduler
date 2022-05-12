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
	mealsTemplate      *template.Template
	mealDetailTemplate *template.Template
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

	meals := model.GetAllMeals()
	vm := viewmodel.NewMeals(meals)
	m.mealsTemplate.Execute(w, vm)
}

func (m meals) handleDetail(w http.ResponseWriter, r *http.Request, mealId int) {
	meal := model.GetMealById(mealId)

	if r.Method == http.MethodPost {
		parseMealData(&meal, r)
		model.UpdateMeal(meal)
		http.Redirect(w, r, "/meals", http.StatusTemporaryRedirect)
	}

	vm := viewmodel.NewMealDetail(meal)
	m.mealDetailTemplate.Execute(w, vm)
}

func (m meals) handleNew(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var newMeal model.Meal
		parseMealData(&newMeal, r)
		model.AddMeal(newMeal)
		http.Redirect(w, r, "/meals", http.StatusTemporaryRedirect)
	}
	vm := viewmodel.NewMealDetail(model.Meal{})
	m.mealDetailTemplate.Execute(w, vm)
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
