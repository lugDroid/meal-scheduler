package model

type Storage interface {
	GetAllMeals() []Meal
	GetMealById(id int) Meal
	UpdateMeal(newMeal Meal) Meal
	AddMeal(newMeal Meal) Meal
	DeleteMeal(mealId int)
	GetAllCategories() []Category
	GetCategoryById(id int) Category
	UpdateCategory(newCategory Category) Category
	AddCategory(newCategory Category) Category
	DeleteCategory(categoryId int)
	GetAllSchedules() []Schedule
	GetScheduleById(id int) Schedule
	UpdateSchedule(newSchedule Schedule) Schedule
	AddSchedule(newSchedule Schedule) Schedule
	DeleteSchedule(scheduleId int)
}
