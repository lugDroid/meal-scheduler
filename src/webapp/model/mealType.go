package model

type MealType int

const (
	Disabled MealType = iota
	Lunch
	Dinner
	Both
)

func (mealType MealType) ToString() string {
	names := []string{"Disabled", "Lunch", "Dinner", "Both"}

	if mealType < Disabled || mealType > Both {
		return "Unknown"
	}

	return names[mealType]
}
