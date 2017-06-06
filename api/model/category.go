package model

import "database/sql"

// Category 항목 모델
type Category struct {
	ID     int64  `db:"id" json:"id"`
	Name   string `db:"name" json:"name" form:"name" binding:"required"`
	Budget int    `db:"budget" json:"budget" form:"budget" binding:"required"`
	State  bool   `db:"state" json:"state" form:"state"`
}

// GetCategoryList 모든 항목 조회
func (db *DBPool) GetCategoryList() ([]Category, error) {
	query := "SELECT * FROM category"
	categoryArray := []Category{}
	err := db.Master.Select(&categoryArray, query)
	return categoryArray, err
}

// GetCategory 항목의 id로 조회
func (db *DBPool) GetCategory(id int64) (Category, error) {
	query := "SELECT * FROM category WHERE id=?"
	category := Category{}
	err := db.Master.Get(&category, query, id)
	return category, err
}

// SetCategory 항목 생성
func (db *DBPool) SetCategory(category *Category) (sql.Result, error) {
	query := `
		INSERT INTO category (name, budget)
			VALUES (:name, :budget)
	`
	result, err := db.Master.NamedExec(query, category)
	return result, err
}

// PutCategory 항목 수정
func (db *DBPool) PutCategory(category *Category) (sql.Result, error) {
	query := `
		UPDATE category
			SET name=:name, budget=:budget, state=:state
			WHERE id=:id
	`
	result, err := db.Master.NamedExec(query, category)
	return result, err
}

// DeleteCategory 항목 제거
func (db *DBPool) DeleteCategory(id int64) (sql.Result, error) {
	query := "DELETE FROM category WHERE name=?"
	result, err := db.Master.Exec(query, id)
	return result, err
}
