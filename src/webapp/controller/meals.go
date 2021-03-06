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
	storage        model.Storage
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

	//meals := model.GetAllMeals()
	meals := m.storage.GetAllMeals()
	vm := viewmodel.NewMeals(meals)
	err := m.listTemplate.Execute(w, vm)
	if err != nil {
		log.Println("Could not execute template", m.listTemplate.Name(), err)
	}
}

func (m meals) handleDetail(w http.ResponseWriter, r *http.Request, mealId int) {
	meal := m.storage.GetMealById(mealId)

	if r.Method == http.MethodPost {
		m.parseMealData(&meal, r)
		m.storage.UpdateMeal(meal)
		http.Redirect(w, r, "/meals", http.StatusTemporaryRedirect)
		return
	}

	vm := viewmodel.NewMealDetail(meal)
	vm.Categories = m.storage.GetAllCategories()
	err := m.detailTemplate.Execute(w, vm)
	if err != nil {
		log.Println("Could not execute template", m.detailTemplate.Name(), err)
	}
}

func (m meals) handleNew(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var newMeal model.Meal
		m.parseMealData(&newMeal, r)
		m.storage.AddMeal(newMeal)
		http.Redirect(w, r, "/meals", http.StatusTemporaryRedirect)
		return
	}
	vm := viewmodel.NewMealDetail(model.Meal{})
	vm.Categories = m.storage.GetAllCategories()
	err := m.detailTemplate.Execute(w, vm)
	if err != nil {
		log.Println("Could not execute template", m.detailTemplate.Name(), err)
	}
}

func (m meals) handleDelete(w http.ResponseWriter, r *http.Request, mealId int) {
	if r.Method == http.MethodPost {
		m.storage.DeleteMeal(mealId)
		http.Redirect(w, r, "/meals", http.StatusTemporaryRedirect)
		return
	}

	vm := viewmodel.NewDeleteViewModel()
	vm.Active = "meals"
	vm.Content = "meal"
	vm.Name = m.storage.GetMealById(mealId).Name
	vm.ReturnPath = "/meals"

	err := m.deleteTemplate.Execute(w, vm)
	if err != nil {
		log.Println("Could not execute template", m.deleteTemplate.Name(), err)
	}
}

func (m meals) parseMealData(meal *model.Meal, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("Could not parse mealDetail form", err)
	}

	meal.Name = r.Form.Get("meal-name")
	meal.Description = r.Form.Get("meal-description")

	categoryId, err := strconv.Atoi(r.Form.Get("categories"))
	if err != nil {
		log.Println("Could not parse categoryId", err)
	}
	meal.Category = m.storage.GetCategoryById(categoryId)

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
