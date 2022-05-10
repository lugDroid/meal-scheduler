package model

type Schedule struct {
	Id          int
	UserId      string
	Title       string
	LunchMeals  [7]Meal
	DinnerMeals [7]Meal
}

func (s Schedule) NewSchedule(meals []Meal) {
	populate(&s.LunchMeals, Lunch)
	populate(&s.DinnerMeals, Dinner)

	/* 	for i, l := range s.LunchMeals {
	   		fmt.Println(i+1, l.Name, ", Servings", l.Servings, ", Type", l.Type.ToString(), l.MainIngredient.Name)
	   	}
	   	fmt.Println("##################")
	   	for i, l := range s.DinnerMeals {
	   		fmt.Println(i+1, l.Name, ", Servings", l.Servings, ", Type", l.Type.ToString(), l.MainIngredient.Name)
	   	} */
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

		usedAlready := getIngredientUses(selectedMeal.MainIngredient, *list)
		if usedAlready >= selectedMeal.MainIngredient.ServingsPerWeek {
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

func getIngredientUses(i Ingredient, list [7]Meal) int {
	count := 0

	for _, m := range list {
		if m.MainIngredient.Name == i.Name {
			count++
		}
	}

	return count
}
