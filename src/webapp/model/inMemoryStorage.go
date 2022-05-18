package model

import (
	"math/rand"
	"time"
)

type InMemoryStorage struct {
}

func (s InMemoryStorage) GetAllMeals() []Meal {
	return meals
}

func (s InMemoryStorage) GetMealById(id int) Meal {
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

func (s InMemoryStorage) UpdateMeal(newMeal Meal) Meal {
	meal := s.GetMealById(newMeal.Id)

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

func (s InMemoryStorage) AddMeal(newMeal Meal) []Meal {
	newMeal.Id = meals[len(meals)-1].Id + 1
	meals = append(meals, newMeal)

	return meals
}

func (s InMemoryStorage) DeleteMeal(mealId int) {
	for i := 0; i < len(meals); i++ {
		if meals[i].Id == mealId {
			meals = append(meals[:i], meals[i+1:]...)
		}
	}
}

func (s InMemoryStorage) GetAllCategories() []Category {
	return categories
}

func (s InMemoryStorage) GetCategoryById(id int) Category {
	for _, c := range categories {
		if c.Id == id {
			return c
		}
	}

	return Category{}
}

func (s InMemoryStorage) UpdateCategory(newCategory Category) Category {
	category := s.GetCategoryById(newCategory.Id)

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

func (s InMemoryStorage) AddCategory(newCategory Category) []Category {
	newCategory.Id = categories[len(categories)-1].Id + 1
	categories = append(categories, newCategory)

	return categories
}

func (s InMemoryStorage) DeleteCategory(categoryId int) {
	for i := 0; i < len(categories); i++ {
		if categories[i].Id == categoryId {
			categories = append(categories[:i], categories[i+1:]...)
		}
	}
}

func (s InMemoryStorage) GetAllSchedules() []Schedule {
	return schedules
}

func (s InMemoryStorage) GetScheduleById(id int) Schedule {
	for _, s := range schedules {
		if s.Id == id {
			return s
		}
	}

	return Schedule{}
}

func (s InMemoryStorage) UpdateSchedule(newSchedule Schedule) Schedule {
	schedule := s.GetScheduleById(newSchedule.Id)

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

func (s InMemoryStorage) AddSchedule(newSchedule Schedule) []Schedule {
	if len(schedules) == 0 {
		newSchedule.Id = 0
	} else {
		newSchedule.Id = schedules[len(schedules)-1].Id + 1
	}
	schedules = append(schedules, newSchedule)

	return schedules
}

func (s InMemoryStorage) DeleteSchedule(scheduleId int) {
	for i := 0; i < len(schedules); i++ {
		if schedules[i].Id == scheduleId {
			schedules = append(schedules[:i], schedules[i+1:]...)
		}
	}
}
