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
			fmt.Println("Error scanning GetAllMeals query results", err)
		}
		m.Category = s.GetCategoryById(categoryId)

		meals = append(meals, m)
	}

	return meals
}

func (s *DbStorage) GetMealById(mealId int) Meal {
	m := Meal{}
	var categoryId int

	row := s.db.QueryRow(`
		SELECT id, name, description, servings, type, category_id
		FROM meal
		WHERE id = $1
	`, mealId)
	err := row.Scan(&m.Id, &m.Name, &m.Description, &m.Servings, &m.Type, &categoryId)
	m.Category = s.GetCategoryById(categoryId)

	if err != nil {
		fmt.Println("GetMealById query failed", err)
	}

	return m
}

func (s *DbStorage) UpdateMeal(m Meal) Meal {
	_, err := s.db.Exec(`
		UPDATE meal
		SET name = $1, description = $2, servings = $3, type = $4, category_id = $5
		WHERE id = $6
	`, m.Name, m.Description, m.Servings, m.Type, m.Category.Id, m.Id)
	if err != nil {
		fmt.Println("UpdateMeal query failed", err)
	}

	return s.GetMealById(m.Id)
}

func (s *DbStorage) AddMeal(nm Meal) Meal {
	err := s.db.QueryRow(`
		INSERT INTO meal (name, description, servings, type, category_id)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`, nm.Name, nm.Description, nm.Servings, nm.Type, nm.Category.Id).Scan(&nm.Id)
	if err != nil {
		fmt.Println("AddMeal query failed", err)
	}

	return nm
}

func (s *DbStorage) DeleteMeal(mealId int) {
	_, err := s.db.Exec(`
		DELETE FROM meal
		WHERE id = $1
	`, mealId)
	if err != nil {
		fmt.Println("DeleteMeal query failed")
	}
}

func (s *DbStorage) GetAllCategories() []Category {
	rows, err := s.db.Query(`
		SELECT id, name, description, servings_per_week
		FROM category
	`)
	if err != nil {
		fmt.Println("GetAllCategories query failed", err)
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		c := Category{}
		err := rows.Scan(&c.Id, &c.Name, &c.Description, &c.ServingsPerWeek)
		if err != nil {
			fmt.Println("Error scanning GetAllCategories query result", err)
		}

		categories = append(categories, c)
	}

	return categories
}

func (s *DbStorage) GetCategoryById(categoryId int) Category {
	c := Category{}

	row := s.db.QueryRow(`
		SELECT id, name, description, servings_per_week
		FROM category
		WHERE id = $1
	`, categoryId)
	err := row.Scan(&c.Id, &c.Name, &c.Description, &c.ServingsPerWeek)
	if err != nil {
		fmt.Println("GetCategoryById query failed", err)
	}

	return c
}

func (s *DbStorage) UpdateCategory(c Category) Category {
	_, err := s.db.Exec(`
		UPDATE category
		SET name = $1, description = $2, servings_per_week = $3
		WHERE id = $4
	`, c.Name, c.Description, c.ServingsPerWeek, c.Id)
	if err != nil {
		fmt.Println("UpdateCategory query failed", err)
	}

	return s.GetCategoryById(c.Id)
}

func (s *DbStorage) AddCategory(newCategory Category) Category {
	err := s.db.QueryRow(`
		INSERT INTO category (name, description, servings_per_week)
		VALUES ($1, $2, $3)
		RETURNING id
	`, newCategory.Name, newCategory.Description, newCategory.ServingsPerWeek).Scan(&newCategory.Id)
	if err != nil {
		fmt.Println("AddCategory query failed", err)
	}

	return newCategory
}

func (s *DbStorage) DeleteCategory(categoryId int) {
	_, err := s.db.Exec(`
		DELETE FROM category
		WHERE id = $1
	`, categoryId)
	if err != nil {
		fmt.Println("DeleteCategory query failed", err)
	}
}

func (s *DbStorage) GetAllSchedules() []Schedule {
	rows, err := s.db.Query(`
		SELECT id, name, lunch_meals, dinner_meals
		FROM schedule
	`)

	if err != nil {
		fmt.Println("GetAllSchedules query failed", err)
	}
	defer rows.Close()

	var schedules []Schedule
	for rows.Next() {
		sc := Schedule{}
		var lunchMealsId, dinnerMealsId int
		err := rows.Scan(&sc.Id, &sc.Name, &lunchMealsId, &dinnerMealsId)
		if err != nil {
			fmt.Println("Error scanning GetAllSchedules query result", err)
		}

		sc.LunchMeals = s.getMealListById("lunch_meals", lunchMealsId)
		sc.DinnerMeals = s.getMealListById("dinner_meals", dinnerMealsId)

		schedules = append(schedules, sc)
	}

	return schedules
}

func (s *DbStorage) GetScheduleById(scheduleId int) Schedule {
	sc := Schedule{}
	var lunchMealsId, dinnerMealsId int
	row := s.db.QueryRow(`
		SELECT id, name, lunch_meals, dinner_meals
		FROM schedule
		WHERE id = $1
	`, scheduleId)
	err := row.Scan(&sc.Id, &sc.Name, &lunchMealsId, &dinnerMealsId)
	if err != nil {
		fmt.Println("GetScheduleById query failed", err)
	}

	sc.LunchMeals = s.getMealListById("lunch_meals", lunchMealsId)
	sc.DinnerMeals = s.getMealListById("dinner_meals", dinnerMealsId)

	return sc
}

func (s *DbStorage) UpdateSchedule(sc Schedule) Schedule {
	// get current meal lists id's
	var lunchMealsId, dinnerMealsId int
	row := s.db.QueryRow(`
		SELECT lunch_meals, dinner_meals
		FROM schedule
		WHERE id = $1
	`, sc.Id)
	err := row.Scan(&lunchMealsId, &dinnerMealsId)
	if err != nil {
		fmt.Println("UpdateSchedule query failed", err)
	}

	// update meal lists with meals from schedule
	s.updateMealList("lunch_meals", lunchMealsId, sc.LunchMeals)
	s.updateMealList("dinner_meals", dinnerMealsId, sc.DinnerMeals)

	// update schedule name
	_, err = s.db.Exec(`
		UPDATE schedule
		SET name = $1
		WHERE id = $2
	`, sc.Name, sc.Id)
	if err != nil {
		fmt.Println("Update schedule name failed", err)
	}

	return s.GetScheduleById(sc.Id)
}

func (s *DbStorage) AddSchedule(sc Schedule) Schedule {
	lunchMealsId := s.addMealList("lunch_meals", sc.LunchMeals)
	dinnerMealsId := s.addMealList("dinner_meals", sc.DinnerMeals)

	err := s.db.QueryRow(`
		INSERT INTO schedule (name, lunch_meals, dinner_meals)
		VALUES ($1, $2, $3)
		RETURNING id
	`, sc.Name, lunchMealsId, dinnerMealsId).Scan(&sc.Id)
	if err != nil {
		fmt.Println("AddSchedule query failed", err)
	}

	return sc
}

func (s *DbStorage) DeleteSchedule(scheduleId int) {
	// Get meal lists id´s for later use
	var lunchMealsId, dinnerMealsId int
	row := s.db.QueryRow(`
	SELECT lunch_meals, dinner_meals
	FROM schedule
	WHERE id = $1
	`, scheduleId)
	err := row.Scan(&lunchMealsId, &dinnerMealsId)
	if err != nil {
		fmt.Println("DeleteSchedule meals query failed", err)
	}

	_, err = s.db.Exec(`
		DELETE FROM schedule
		WHERE id = $1
	`, scheduleId)
	if err != nil {
		fmt.Println("DeleteSchedule query failed", err)
	}

	// Delete also meal lists
	s.deleteMealList("lunch_meals", lunchMealsId)
	s.deleteMealList("dinner_meals", dinnerMealsId)
}

// Aux functions
func (s *DbStorage) getMealListById(listName string, listId int) [7]Meal {
	var mealIds [7]int
	var sqlStatement string

	switch {
	case listName == "lunch_meals":
		sqlStatement = `SELECT day_0, day_1, day_2, day_3, day_4, day_5, day_6
						FROM lunch_meals
						WHERE id = $1`
	case listName == "dinner_meals":
		sqlStatement = `SELECT day_0, day_1, day_2, day_3, day_4, day_5, day_6
						FROM dinner_meals
						WHERE id = $1`
	}

	row := s.db.QueryRow(sqlStatement, listId)
	err := row.Scan(&mealIds[0], &mealIds[1], &mealIds[2], &mealIds[3], &mealIds[4], &mealIds[5], &mealIds[6])
	if err != nil {
		fmt.Println("GetMealListById query failed", err)
	}

	var meals [7]Meal
	for i := range meals {
		meals[i] = s.GetMealById(mealIds[i])
	}

	return meals
}

func (s *DbStorage) updateMealList(listName string, listId int, list [7]Meal) {
	var sqlStatement string

	switch {
	case listName == "lunch_meals":
		sqlStatement = `UPDATE lunch_meals
						SET day_0 = $1, day_1 = $2, day_2 = $3, day_3 = $4, day_4 = $5, day_5 = $6, day_6 = $7
						WHERE id = $8`
	case listName == "dinner_meals":
		sqlStatement = `UPDATE dinner_meals
						SET day_0 = $1, day_1 = $2, day_2 = $3, day_3 = $4, day_4 = $5, day_5 = $6, day_6 = $7
						WHERE id = $8`
	}

	_, err := s.db.Exec(sqlStatement, list[0].Id, list[1].Id, list[2].Id, list[3].Id, list[4].Id, list[5].Id, list[6].Id, listId)
	if err != nil {
		fmt.Println("Update meal list failed", err)
	}
}

func (s *DbStorage) addMealList(listName string, list [7]Meal) int {
	var listId int
	var sqlStatement string

	switch {
	case listName == "lunch_meals":
		sqlStatement = `INSERT INTO lunch_meals (day_0, day_1, day_2, day_3, day_4, day_5, day_6)
						VALUES ($1, $2, $3, $4, $5, $6, $7)
						RETURNING id`
	case listName == "dinner_meals":
		sqlStatement = `INSERT INTO dinner_meals (day_0, day_1, day_2, day_3, day_4, day_5, day_6)
						VALUES ($1, $2, $3, $4, $5, $6, $7)
						RETURNING id`
	}
	err := s.db.QueryRow(sqlStatement, list[0].Id, list[1].Id, list[2].Id, list[3].Id, list[4].Id, list[5].Id, list[6].Id).Scan(&listId)
	if err != nil {
		fmt.Println("AddMealList query failed", err)
	}

	return listId
}

func (s *DbStorage) deleteMealList(listName string, listId int) {
	var sqlStatement string

	switch {
	case listName == "lunch_meals":
		sqlStatement = `DELETE FROM lunch_meals
						WHERE id = $1`
	case listName == "dinner_meals":
		sqlStatement = `DELETE FROM dinner_meals
						WHERE id = $1`
	}
	_, err := s.db.Exec(sqlStatement, listId)
	if err != nil {
		fmt.Println("deleteMealList query failed", err)
	}
}
