package model

import "database/sql"

// Household 가계부 모델
type Household struct {
	ID           int64  `db:"id" json:"id"`
	TypeOf       string `db:"typeof" json:"typeof" form:"typeof"`
	CategoryName string `db:"category_name" json:"categoryName" form:"categoryName"`
	Content      string `db:"content" json:"content" form:"content"`
	Money        int    `db:"money" json:"money" form:"money"`
	Others       string `db:"others" json:"others" form:"others"`
	Author       string `db:"author" json:"author" form:"author"`
	State        bool   `db:"state" json:"state" form:"state"`
	CreateAt     int64  `db:"create_at" json:"createAt"`
}

// GetHouseholdArray 모든 가계부 조회
func (db *DBPool) GetHouseholdArray(categoryName string) ([]Household, error) {
	query := "SELECT * FROM household WHERE category_name=?"
	householdArray := []Household{}
	err := db.Master.Select(&householdArray, query, categoryName)
	return householdArray, err
}

// GetHousehold 가계부의 id로 조회
func (db *DBPool) GetHousehold(id int64) (Household, error) {
	query := "SELECT * FROM household WHERE id=?"
	household := Household{}
	err := db.Master.Get(&household, query, id)
	return household, err
}

// SetHousehold 가계부 생성
func (db *DBPool) SetHousehold(household *Household) (sql.Result, error) {
	query := `
		INSERT INTO household 
			(typeof, category_name, content, money, others, author)
		VALUES
			(:typeof, :category_name, content, money, others, author)
	`
	result, err := db.Master.NamedExec(query, household)
	return result, err
}

// PutHousehold 가계부 수정
func (db *DBPool) PutHousehold(household *Household) (sql.Result, error) {
	query := `
		UPDATE household
			SET typeof=:typeof, category_name=:category_name,
				content=:content, money=:money, others=:others, author=:author
			WHERE id=:id
	`
	result, err := db.Master.NamedExec(query, household)
	return result, err
}

// DeleteHousehold 가계부 제거
func (db *DBPool) DeleteHousehold(id int64) (sql.Result, error) {
	query := "DELETE FROM household WHERE id=?"
	result, err := db.Master.Exec(query, id)
	return result, err
}
