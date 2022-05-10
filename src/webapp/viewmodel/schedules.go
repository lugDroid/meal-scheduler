package viewmodel

import "lugdroid/mealsScheduler/webapp/model"

type Schedules struct {
	Title     string
	Schedules []model.Schedule
	Active    string
}

func NewSchedules(schedules []model.Schedule) Schedules {
	return Schedules{
		Title:     "Meal Scheduler - Schedules",
		Schedules: schedules,
		Active:    "schedules",
	}
}
