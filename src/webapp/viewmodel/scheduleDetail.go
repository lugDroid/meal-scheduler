package viewmodel

import "lugdroid/mealsScheduler/webapp/model"

type ScheduleDetail struct {
	Title       string
	Schedule    model.Schedule
	LunchMeals  []model.Meal
	DinnerMeals []model.Meal
	Active      string
}

func NewScheduleDetail(schedule model.Schedule) ScheduleDetail {
	return ScheduleDetail{
		Title:    "Meal Scheduler - Schedule Detail",
		Schedule: schedule,
		Active:   "schedules",
	}
}
