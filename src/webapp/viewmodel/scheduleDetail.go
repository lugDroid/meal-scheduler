package viewmodel

import "lugdroid/mealsScheduler/webapp/model"

type ScheduleDetail struct {
	Title    string
	Schedule model.Schedule
	Meals    []model.Meal
	Active   string
}

func NewScheduleDetail(schedule model.Schedule, meals []model.Meal) ScheduleDetail {
	return ScheduleDetail{
		Title:    "Meal Scheduler - Schedule Detail",
		Schedule: schedule,
		Meals:    meals,
		Active:   "schedules",
	}
}
