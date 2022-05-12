package controller

import (
	"html/template"
	"lugdroid/mealsScheduler/webapp/model"
	"lugdroid/mealsScheduler/webapp/viewmodel"
	"net/http"
	"regexp"
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
	/* 	idPattern, _ := regexp.Compile(`/schedules/(\d+)`)
	   	idMatches := idPattern.FindStringSubmatch(r.URL.Path)
	   	if len(idMatches) > 0 {
	   		scheduleId, _ := strconv.Atoi(idMatches[1])
	   		s.handleDetail(w, r, scheduleId)
	   		return
	   	} */

	newPattern, _ := regexp.Compile(`/schedules/new$`)
	newMatches := newPattern.FindStringSubmatch(r.URL.Path)
	if len(newMatches) > 0 {
		s.handleNew(w, r)
		return
	}

	// To-Do: add get all schedules method to memory storage
	schedules := []model.Schedule{}

	vm := viewmodel.NewSchedules(schedules)
	s.schedulesTemplate.Execute(w, vm)
}

/* func (s schedules) handleDetail(w http.ResponseWriter, r *http.Request, scheduleId int) {
	schedule := model.GetScheduleById(scheduleId)

	if r.Method == http.MethodPost {

	}
}

func parseScheduleData(schedule *model.Schedule, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("Could not parse scheduleDetail form", err)
	}

	schedule.Title = r.Form.Get("schedule-name")

	for i, m := range schedule.LunchMeals {
		fieldName := "lunchmeal-" + strconv.Itoa(i)
		schedule.LunchMeals[i] = r.Form.Get()
	}
} */

func (s schedules) handleNew(w http.ResponseWriter, r *http.Request) {
	meals := model.GetAllMeals()

	sc := model.Schedule{}
	sc.PopulateMeals(meals)

	vm := viewmodel.NewScheduleDetail(sc, meals)
	s.scheduleDetailTemplate.Execute(w, vm)
}
