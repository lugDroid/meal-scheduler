package model

type Meal struct {
	Id             int
	UserId         string
	Name           string
	Description    string
	Servings       int
	MainIngredient Ingredient
	Type           MealType
}
