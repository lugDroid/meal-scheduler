package controller

import (
	"html/template"
	"lugdroid/mealsScheduler/webapp/model"
	"lugdroid/mealsScheduler/webapp/viewmodel"
	"net/http"
)

type schedules struct {
	schedulesTemplate *template.Template
}

func (s schedules) registerRoutes() {
	http.HandleFunc("/schedules", s.HandleSchedules)
}

func (s schedules) HandleSchedules(w http.ResponseWriter, r *http.Request) {
	sc1 := model.Schedule{
		Title: "Schedule 1",
	}
	sc1.NewSchedule(model.GetAllMeals())
	sc2 := model.Schedule{
		Title: "Schedule 2",
	}
	sc2.NewSchedule(model.GetAllMeals())

	schedules := []model.Schedule{sc1, sc2}

	vm := viewmodel.NewSchedules(schedules)
	s.schedulesTemplate.Execute(w, vm)
}
