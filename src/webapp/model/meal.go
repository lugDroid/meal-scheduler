package model

type Meal struct {
	Id          int
	UserId      string
	Name        string
	Description string
	Servings    int
	Category    Category
	Type        MealType
}
