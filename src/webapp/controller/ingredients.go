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

type ingredients struct {
	listTemplate   *template.Template
	detailTemplate *template.Template
	deleteTemplate *template.Template
}

func (i ingredients) registerRoutes() {
	http.HandleFunc("/ingredients", i.handleIngredients)
	http.HandleFunc("/ingredients/", i.handleIngredients)
}

func (i ingredients) handleIngredients(w http.ResponseWriter, r *http.Request) {
	idPattern, _ := regexp.Compile(`/ingredients/(\d+)`)
	idMatches := idPattern.FindStringSubmatch(r.URL.Path)
	if len(idMatches) > 0 {
		ingredientId, _ := strconv.Atoi(idMatches[1])
		i.handleDetail(w, r, ingredientId)
		return
	}

	newPattern, _ := regexp.Compile(`/ingredients/new$`)
	newMatches := newPattern.FindStringSubmatch(r.URL.Path)
	if len(newMatches) > 0 {
		i.handleNew(w, r)
		return
	}

	deletePattern, _ := regexp.Compile(`/ingredients/delete/(\d+)`)
	deleteMatches := deletePattern.FindStringSubmatch(r.URL.Path)
	if len(deleteMatches) > 0 {
		ingredientId, _ := strconv.Atoi(deleteMatches[1])
		i.handleDelete(w, r, ingredientId)
		return
	}

	ingredients := model.GetAllIngredients()
	vm := viewmodel.NewIngredients(ingredients)
	err := i.listTemplate.Execute(w, vm)
	if err != nil {
		log.Println("Could not execute template", i.listTemplate.Name(), err)
	}
}

func (i ingredients) handleDetail(w http.ResponseWriter, r *http.Request, ingredientId int) {
	ingredient := model.GetIngredientById(ingredientId)

	if r.Method == http.MethodPost {
		parseIngredientData(&ingredient, r)
		model.UpdateIngredient(ingredient)
		http.Redirect(w, r, "/ingredients", http.StatusTemporaryRedirect)
		return
	}

	vm := viewmodel.NewIngredientDetail(ingredient)
	err := i.detailTemplate.Execute(w, vm)
	if err != nil {
		log.Println("Could not execute template", i.detailTemplate.Name(), err)
	}
}

func (i ingredients) handleNew(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var newIngredient model.Ingredient
		parseIngredientData(&newIngredient, r)
		model.AddIngredient(newIngredient)
		http.Redirect(w, r, "/ingredients", http.StatusTemporaryRedirect)
		return
	}

	vm := viewmodel.NewIngredientDetail(model.Ingredient{})
	err := i.detailTemplate.Execute(w, vm)
	if err != nil {
		log.Println("Could not execute template", i.detailTemplate.Name(), err)
	}
}

func (i ingredients) handleDelete(w http.ResponseWriter, r *http.Request, ingredientId int) {
	if r.Method == http.MethodPost {
		model.DeleteIngredient(ingredientId)
		http.Redirect(w, r, "/ingredients", http.StatusTemporaryRedirect)
		return
	}

	vm := viewmodel.NewDeleteViewModel("ingredient", model.GetIngredientById(ingredientId).Name, "/ingredients")
	err := i.deleteTemplate.Execute(w, vm)
	if err != nil {
		log.Println("Could not execute template", i.deleteTemplate.Name(), err)
	}
}

func parseIngredientData(i *model.Ingredient, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Print("Could not parse ingredientDetail form", err)
	}

	i.Name = r.Form.Get("ingredient-name")
	i.Description = r.Form.Get("ingredient-desc")

	ingredientServings, err := strconv.Atoi(r.Form.Get("ingredient-servings"))
	if err != nil {
		log.Println("Could not parse ingredientServings", err)
	}
	i.ServingsPerWeek = ingredientServings
}
