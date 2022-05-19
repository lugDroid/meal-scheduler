package model

import (
	"math/rand"
	"time"
)

type Schedule struct {
	Id          int
	UserId      string
	Name        string
	LunchMeals  [7]Meal
	DinnerMeals [7]Meal
}

func (s *Schedule) PopulateMeals(meals []Meal) {
	s.LunchMeals = populate(meals, Lunch)
	s.DinnerMeals = populate(meals, Dinner)
}

func populate(allMeals []Meal, mealType MealType) [7]Meal {
	i := 0
	var meals [7]Meal

out:
	for i < len(meals) {
		selectedMeal := getRandomMeal(allMeals, mealType)
		for _, m := range meals {
			if m.Name == selectedMeal.Name {
				continue out
			}
		}

		usedAlready := getCategoryUses(selectedMeal.Category, meals)
		if usedAlready >= selectedMeal.Category.ServingsPerWeek {
			continue
		}

		// add the meal number-of-servings times
		for j := 0; j < selectedMeal.Servings; j++ {
			meals[i] = selectedMeal
			i++

			if i == len(meals) {
				break
			}
		}
	}

	return meals
}

func getCategoryUses(c Category, list [7]Meal) int {
	count := 0

	for _, m := range list {
		if m.Category.Name == c.Name {
			count++
		}
	}

	return count
}

func getRandomMeal(allMeals []Meal, mealType MealType) Meal {
	for {
		rand.Seed((time.Now().UnixNano()))
		index := rand.Intn(len(allMeals))
		if allMeals[index].Type == mealType || allMeals[index].Type == Both {
			return allMeals[index]
		}
	}
}
