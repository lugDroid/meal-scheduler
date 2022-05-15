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

	schedules := model.GetAllSchedules()

	vm := viewmodel.NewSchedules(schedules)
	err := s.listTemplate.Execute(w, vm)
	if err != nil {
		log.Println("Could not execute template", s.listTemplate.Name(), err)
	}
}

func (s schedules) handleDetail(w http.ResponseWriter, r *http.Request, scheduleId int) {
	schedule := model.GetScheduleById(scheduleId)

	if r.Method == http.MethodPost {
		parseFormData(&schedule, r)
		model.UpdateSchedule(schedule)
		http.Redirect(w, r, "/schedules", http.StatusTemporaryRedirect)
		return
	}

	vm := viewmodel.NewScheduleDetail(schedule, model.GetAllMeals())
	err := s.detailTemplate.Execute(w, vm)
	if err != nil {
		log.Println("Could not execute template", s.detailTemplate.Name(), err)
	}
}

func (s schedules) handleDelete(w http.ResponseWriter, r *http.Request, scheduleId int) {
	if r.Method == http.MethodPost {
		model.DeleteSchedule(scheduleId)
		http.Redirect(w, r, "/schedules", http.StatusTemporaryRedirect)
		return
	}

	vm := viewmodel.NewDeleteViewModel("schedule", model.GetScheduleById(scheduleId).Name, "/schedules")
	err := s.deleteTemplate.Execute(w, vm)
	if err != nil {
		log.Println("Could not execute template", s.deleteTemplate.Name(), err)
	}
}

func (s schedules) handleNew(w http.ResponseWriter, r *http.Request) {
	schedule := model.Schedule{}
	meals := model.GetAllMeals()

	if r.Method == http.MethodPost {
		parseFormData(&schedule, r)
		model.AddSchedule(schedule)
		http.Redirect(w, r, "/schedules", http.StatusTemporaryRedirect)
		return
	}

	schedule.PopulateMeals(meals)
	vm := viewmodel.NewScheduleDetail(schedule, meals)
	err := s.detailTemplate.Execute(w, vm)
	if err != nil {
		log.Println("Could not execute template", s.detailTemplate.Name(), err)
	}
}

func parseFormData(schedule *model.Schedule, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("Could not parse scheduleDetail form", err)
	}

	schedule.Name = r.Form.Get("schedule-name")
	schedule.LunchMeals = parseMeals("lunch-meal", r)
	schedule.DinnerMeals = parseMeals("dinner-meal", r)
}

func parseMeals(idText string, r *http.Request) [7]model.Meal {
	var mealList [7]model.Meal

	for i := range mealList {
		fieldName := idText + "-" + strconv.Itoa(i)

		mealId, err := strconv.Atoi(r.Form.Get(fieldName))
		if err != nil {
			log.Println("Could not parse mealId", err)
		}

		mealList[i] = model.GetMealById(mealId)
	}

	return mealList
}
