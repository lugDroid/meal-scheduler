package model

import (
	"database/sql"
	"fmt"
)

type DbStorage struct {
	db *sql.DB
}

func NewDbStorage(database *sql.DB) DbStorage {
	return DbStorage{
		db: database,
	}
}

func (s *DbStorage) GetAllMeals() []Meal {
	rows, err := s.db.Query(`
		SELECT id, name, description, servings, type, category_id
		FROM meal
	`)
	if err != nil {
		fmt.Println("GetAllMeals query failed", err)
	}
	defer rows.Close()

	var meals []Meal
	for rows.Next() {
		m := Meal{}
		var categoryId int
		err := rows.Scan(&m.Id, &m.Name, &m.Description, &m.Servings, &m.Type, &categoryId)
		if err != nil {
			fmt.Println("Error scanning query results", err)
		}
		m.Category = GetCategoryById(categoryId)

		meals = append(meals, m)
	}

	return meals
}
