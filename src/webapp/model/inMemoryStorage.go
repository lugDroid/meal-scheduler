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
	meal.MainIngredient = newMeal.MainIngredient
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

func RemoveMeal(mealId int) {
	for i := range meals {
		if meals[i].Id == mealId {
			meals = append(meals[:i], meals[i+1:]...)
		}
	}
}

func GetAllIngredients() []Ingredient {
	return ingredients
}

func GetIngredientById(id int32) Ingredient {
	for _, i := range ingredients {
		if i.Id == id {
			return i
		}
	}

	return Ingredient{}
}

func UpdateIngredient(newIngredient Ingredient) Ingredient {
	ingredient := GetIngredientById(newIngredient.Id)

	ingredient.Name = newIngredient.Name
	ingredient.Description = newIngredient.Description
	ingredient.ServingsPerWeek = newIngredient.ServingsPerWeek

	for i := range ingredients {
		if ingredients[i].Id == ingredient.Id {
			ingredients[i] = ingredient
		}
	}

	return ingredient
}

func AddIngredient(newIngredient Ingredient) []Ingredient {
	newIngredient.Id = ingredients[len(ingredients)-1].Id + 1
	ingredients = append(ingredients, newIngredient)

	return ingredients
}

func RemoveIngredient(ingredientId int32) {
	for i := range ingredients {
		if ingredients[i].Id == ingredientId {
			ingredients = append(ingredients[:i], ingredients[i+1:]...)
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

func RemoveSchedule(scheduleId int) {
	for i := range schedules {
		if schedules[i].Id == scheduleId {
			schedules = append(schedules[:i], schedules[i+1:]...)
		}
	}
}
