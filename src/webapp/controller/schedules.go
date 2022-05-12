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
	schedulesTemplate      *template.Template
	scheduleDetailTemplate *template.Template
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

	schedules := model.GetAllSchedules()

	vm := viewmodel.NewSchedules(schedules)
	s.schedulesTemplate.Execute(w, vm)
}

func (s schedules) handleDetail(w http.ResponseWriter, r *http.Request, scheduleId int) {
	schedule := model.GetScheduleById(scheduleId)

	if r.Method == http.MethodPost {
		parseFormData(&schedule, r)
		model.UpdateSchedule(schedule)
		http.Redirect(w, r, "/schedules", http.StatusTemporaryRedirect)
	}

	vm := viewmodel.NewScheduleDetail(schedule, model.GetAllMeals())
	s.scheduleDetailTemplate.Execute(w, vm)
}

func (s schedules) handleNew(w http.ResponseWriter, r *http.Request) {
	schedule := model.Schedule{}
	meals := model.GetAllMeals()

	if r.Method == http.MethodPost {
		parseFormData(&schedule, r)
		model.AddSchedule(schedule)
		http.Redirect(w, r, "/schedules", http.StatusTemporaryRedirect)
	}

	schedule.PopulateMeals(meals)
	vm := viewmodel.NewScheduleDetail(schedule, meals)
	s.scheduleDetailTemplate.Execute(w, vm)
}

func parseFormData(schedule *model.Schedule, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("Could not parse scheduleDetail form", err)
	}

	schedule.Title = r.Form.Get("schedule-name")
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
