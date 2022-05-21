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

type schedules struct {
	listTemplate   *template.Template
	detailTemplate *template.Template
	deleteTemplate *template.Template
	storage        model.Storage
}

func (s schedules) registerRoutes() {
	http.HandleFunc("/schedules", s.handleSchedules)
	http.HandleFunc("/schedules/", s.handleSchedules)
}

func (s schedules) handleSchedules(w http.ResponseWriter, r *http.Request) {
	idPattern, _ := regexp.Compile(`/schedules/(\d+)`)
	idMatches := idPattern.FindStringSubmatch(r.URL.Path)
	if len(idMatches) > 0 {
		scheduleId, _ := strconv.Atoi(idMatches[1])
		s.handleDetail(w, r, scheduleId)
		return
	}

	newPattern, _ := regexp.Compile(`/schedules/new$`)
	newMatches := newPattern.FindStringSubmatch(r.URL.Path)
	if len(newMatches) > 0 {
		s.handleNew(w, r)
		return
	}

	deletePattern, _ := regexp.Compile(`/schedules/delete/(\d+)`)
	deleteMatches := deletePattern.FindStringSubmatch(r.URL.Path)
	if len(deleteMatches) > 0 {
		scheduleId, _ := strconv.Atoi(deleteMatches[1])
		s.handleDelete(w, r, scheduleId)
		return
	}

	schedules := s.storage.GetAllSchedules()

	vm := viewmodel.NewSchedules(schedules)
	err := s.listTemplate.Execute(w, vm)
	if err != nil {
		log.Println("Could not execute template", s.listTemplate.Name(), err)
	}
}

func (s schedules) handleDetail(w http.ResponseWriter, r *http.Request, scheduleId int) {
	schedule := s.storage.GetScheduleById(scheduleId)

	if r.Method == http.MethodPost {
		s.parseFormData(&schedule, r)
		s.storage.UpdateSchedule(schedule)
		http.Redirect(w, r, "/schedules", http.StatusTemporaryRedirect)
		return
	}

	vm := viewmodel.NewScheduleDetail(schedule)
	vm.LunchMeals = s.storage.GetMealsByType(model.Lunch)
	vm.DinnerMeals = s.storage.GetMealsByType(model.Dinner)
	err := s.detailTemplate.Execute(w, vm)
	if err != nil {
		log.Println("Could not execute template", s.detailTemplate.Name(), err)
	}
}

func (s schedules) handleDelete(w http.ResponseWriter, r *http.Request, scheduleId int) {
	if r.Method == http.MethodPost {
		s.storage.DeleteSchedule(scheduleId)
		http.Redirect(w, r, "/schedules", http.StatusTemporaryRedirect)
		return
	}

	vm := viewmodel.NewDeleteViewModel()
	vm.Active = "schedules"
	vm.Content = "schedule"
	vm.Name = s.storage.GetScheduleById(scheduleId).Name
	vm.ReturnPath = "/schedules"

	err := s.deleteTemplate.Execute(w, vm)
	if err != nil {
		log.Println("Could not execute template", s.deleteTemplate.Name(), err)
	}
}

func (s schedules) handleNew(w http.ResponseWriter, r *http.Request) {
	schedule := model.Schedule{}
	meals := s.storage.GetAllMeals()

	if r.Method == http.MethodPost {
		s.parseFormData(&schedule, r)
		s.storage.AddSchedule(schedule)
		http.Redirect(w, r, "/schedules", http.StatusTemporaryRedirect)
		return
	}

	schedule.PopulateMeals(meals)
	vm := viewmodel.NewScheduleDetail(schedule)
	vm.LunchMeals = s.storage.GetMealsByType(model.Lunch)
	vm.DinnerMeals = s.storage.GetMealsByType(model.Dinner)
	err := s.detailTemplate.Execute(w, vm)
	if err != nil {
		log.Println("Could not execute template", s.detailTemplate.Name(), err)
	}
}

func (s schedules) parseFormData(schedule *model.Schedule, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("Could not parse scheduleDetail form", err)
	}

	schedule.Name = r.Form.Get("schedule-name")
	schedule.LunchMeals = s.parseMeals("lunch-meal", r)
	schedule.DinnerMeals = s.parseMeals("dinner-meal", r)
}

func (s schedules) parseMeals(idText string, r *http.Request) [7]model.Meal {
	var mealList [7]model.Meal

	for i := range mealList {
		fieldName := idText + "-" + strconv.Itoa(i)

		mealId, err := strconv.Atoi(r.Form.Get(fieldName))
		if err != nil {
			log.Println("Could not parse mealId", err)
		}

		mealList[i] = s.storage.GetMealById(mealId)
	}

	return mealList
}
