package model

var (
	i1 = Ingredient{
		Id:              0,
		Name:            "Patata",
		ServingsPerWeek: 4,
	}

	i2 = Ingredient{
		Id:              1,
		Name:            "Pollo",
		ServingsPerWeek: 2,
	}

	i3 = Ingredient{
		Id:              2,
		Name:            "Masa",
		ServingsPerWeek: 2,
	}

	ingredients = []Ingredient{i1, i2, i3}
)

var (
	meal1 = Meal{
		Id:             0,
		Name:           "Tortilla de Patata",
		MainIngredient: i1,
		Type:           Both,
		Servings:       1,
	}

	meal2 = Meal{
		Id:             1,
		Name:           "Fajitas de Pollo",
		MainIngredient: i2,
		Type:           Both,
		Servings:       1,
	}

	meal3 = Meal{
		Id:             2,
		Name:           "Pizzas de Pimientos y Atun",
		MainIngredient: i3,
		Type:           Dinner,
		Servings:       1,
	}

	meals = []Meal{meal1, meal2, meal3}
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

func UpdateMeal(newMeal Meal) Meal {
	meal := GetMealById(newMeal.Id)

	meal.Name = newMeal.Name
	meal.Description = newMeal.Description
	meal.MainIngredient = newMeal.MainIngredient
	meal.Servings = newMeal.Servings
	meal.Type = newMeal.Type

	for i := 0; i < len(meals); i++ {
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

	for i := 0; i < len(ingredients); i++ {
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
