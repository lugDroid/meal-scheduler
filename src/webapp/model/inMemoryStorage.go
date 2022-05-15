package model

import (
	"math/rand"
	"time"
)

func GetAllMeals() []Meal {
	return meals
}

func GetMealById(id int) Meal {
	for _, m := range meals {
		if m.Id == id {
			return m
		}
	}

	return Meal{}
}

func GetRandomMeal(mealType MealType) Meal {
	for {
		rand.Seed((time.Now().UnixNano()))
		index := rand.Intn(len(meals))
		if meals[index].Type == mealType || meals[index].Type == Both {
			return meals[index]
		}
	}
}

func UpdateMeal(newMeal Meal) Meal {
	meal := GetMealById(newMeal.Id)

	meal.Name = newMeal.Name
	meal.Description = newMeal.Description
	meal.Category = newMeal.Category
	meal.Servings = newMeal.Servings
	meal.Type = newMeal.Type

	for i := range meals {
		if meals[i].Id == meal.Id {
			meals[i] = meal
		}
	}

	return meal
}

func AddMeal(newMeal Meal) []Meal {
	newMeal.Id = meals[len(meals)-1].Id + 1
	meals = append(meals, newMeal)

	return meals
}

func DeleteMeal(mealId int) {
	for i := range meals {
		if meals[i].Id == mealId {
			meals = append(meals[:i], meals[i+1:]...)
		}
	}
}

func GetAllCategories() []Category {
	return categories
}

func GetCategoryById(id int) Category {
	for _, c := range categories {
		if c.Id == id {
			return c
		}
	}

	return Category{}
}

func UpdateCategory(newCategory Category) Category {
	category := GetCategoryById(newCategory.Id)

	category.Name = newCategory.Name
	category.Description = newCategory.Description
	category.ServingsPerWeek = newCategory.ServingsPerWeek

	for i := range categories {
		if categories[i].Id == category.Id {
			categories[i] = category
		}
	}

	return category
}

func AddCategory(newCategory Category) []Category {
	newCategory.Id = categories[len(categories)-1].Id + 1
	categories = append(categories, newCategory)

	return categories
}

func DeleteCategory(categoryId int) {
	for i := range categories {
		if categories[i].Id == categoryId {
			categories = append(categories[:i], categories[i+1:]...)
		}
	}
}

func GetAllSchedules() []Schedule {
	return schedules
}

func GetScheduleById(id int) Schedule {
	for _, s := range schedules {
		if s.Id == id {
			return s
		}
	}

	return Schedule{}
}

func UpdateSchedule(newSchedule Schedule) Schedule {
	schedule := GetScheduleById(newSchedule.Id)

	schedule.Name = newSchedule.Name
	schedule.LunchMeals = newSchedule.LunchMeals
	schedule.DinnerMeals = newSchedule.DinnerMeals

	for i := range schedules {
		if schedules[i].Id == schedule.Id {
			schedules[i] = schedule
		}
	}

	return schedule
}

func AddSchedule(newSchedule Schedule) []Schedule {
	if len(schedules) == 0 {
		newSchedule.Id = 0
	} else {
		newSchedule.Id = schedules[len(schedules)-1].Id + 1
	}
	schedules = append(schedules, newSchedule)

	return schedules
}

func DeleteSchedule(scheduleId int) {
	for i := range schedules {
		if schedules[i].Id == scheduleId {
			schedules = append(schedules[:i], schedules[i+1:]...)
		}
	}
}
