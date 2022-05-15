package model

type Schedule struct {
	Id          int
	UserId      string
	Name        string
	LunchMeals  [7]Meal
	DinnerMeals [7]Meal
}

func (s *Schedule) PopulateMeals(meals []Meal) {
	populate(&s.LunchMeals, Lunch)
	populate(&s.DinnerMeals, Dinner)
}

func populate(list *[7]Meal, mealType MealType) {
	i := 0

out:
	for i < len(*list) {
		selectedMeal := GetRandomMeal(mealType)
		for _, m := range *list {
			if m.Name == selectedMeal.Name {
				continue out
			}
		}

		usedAlready := getCategoryUses(selectedMeal.Category, *list)
		if usedAlready >= selectedMeal.Category.ServingsPerWeek {
			continue
		}

		// add the meal number-of-servings times
		for j := 0; j < selectedMeal.Servings; j++ {
			(*list)[i] = selectedMeal
			i++

			if i == len(*list) {
				break
			}
		}
	}
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
