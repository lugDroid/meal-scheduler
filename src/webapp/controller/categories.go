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

type categories struct {
	listTemplate   *template.Template
	detailTemplate *template.Template
	deleteTemplate *template.Template
}

func (c categories) registerRoutes() {
	http.HandleFunc("/categories", c.handleCategories)
	http.HandleFunc("/categories/", c.handleCategories)
}

func (c categories) handleCategories(w http.ResponseWriter, r *http.Request) {
	idPattern, _ := regexp.Compile(`/categories/(\d+)`)
	idMatches := idPattern.FindStringSubmatch(r.URL.Path)
	if len(idMatches) > 0 {
		categoryId, _ := strconv.Atoi(idMatches[1])
		c.handleDetail(w, r, categoryId)
		return
	}

	newPattern, _ := regexp.Compile(`/categories/new$`)
	newMatches := newPattern.FindStringSubmatch(r.URL.Path)
	if len(newMatches) > 0 {
		c.handleNew(w, r)
		return
	}

	deletePattern, _ := regexp.Compile(`/categories/delete/(\d+)`)
	deleteMatches := deletePattern.FindStringSubmatch(r.URL.Path)
	if len(deleteMatches) > 0 {
		categoryId, _ := strconv.Atoi(deleteMatches[1])
		c.handleDelete(w, r, categoryId)
		return
	}

	categories := model.GetAllCategories()
	vm := viewmodel.NewCategories(categories)
	err := c.listTemplate.Execute(w, vm)
	if err != nil {
		log.Println("Could not execute template", c.listTemplate.Name(), err)
	}
}

func (c categories) handleDetail(w http.ResponseWriter, r *http.Request, categoryId int) {
	category := model.GetCategoryById(categoryId)

	if r.Method == http.MethodPost {
		parseCategoryData(&category, r)
		model.UpdateCategory(category)
		http.Redirect(w, r, "/categories", http.StatusTemporaryRedirect)
		return
	}

	vm := viewmodel.NewCategoryDetail(category)
	err := c.detailTemplate.Execute(w, vm)
	if err != nil {
		log.Println("Could not execute template", c.detailTemplate.Name(), err)
	}
}

func (i categories) handleNew(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var newCategory model.Category
		parseCategoryData(&newCategory, r)
		model.AddCategory(newCategory)
		http.Redirect(w, r, "/categories", http.StatusTemporaryRedirect)
		return
	}

	vm := viewmodel.NewCategoryDetail(model.Category{})
	err := i.detailTemplate.Execute(w, vm)
	if err != nil {
		log.Println("Could not execute template", i.detailTemplate.Name(), err)
	}
}

func (c categories) handleDelete(w http.ResponseWriter, r *http.Request, categoryId int) {
	if r.Method == http.MethodPost {
		model.DeleteCategory(categoryId)
		http.Redirect(w, r, "/categories", http.StatusTemporaryRedirect)
		return
	}

	vm := viewmodel.NewDeleteViewModel()
	vm.Active = "categories"
	vm.Content = "category"
	vm.Name = model.GetCategoryById(categoryId).Name
	vm.ReturnPath = "/categories"

	err := c.deleteTemplate.Execute(w, vm)
	if err != nil {
		log.Println("Could not execute template", c.deleteTemplate.Name(), err)
	}
}

func parseCategoryData(c *model.Category, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Print("Could not parse categoryDetail form", err)
	}

	c.Name = r.Form.Get("category-name")
	c.Description = r.Form.Get("category-desc")

	categoryServings, err := strconv.Atoi(r.Form.Get("category-servings"))
	if err != nil {
		log.Println("Could not parse categoryServings", err)
	}
	c.ServingsPerWeek = categoryServings
}
