package model

import "database/sql"

// User 사용자 모델
type User struct {
	ID       int64   `db:"id" json:"id"`
	Name     string  `db:"name" json:"name" form:"name" binding:"required"`
	PassWD   string  `db:"passwd" json:"passwd" form:"passwd" binding:"required"`
	CreateAt []uint8 `db:"create_at" json:"createAt"`
}

// GetUserList 모든 사용자의 정보를 조회
func (db *DBPool) GetUserList() ([]User, error) {
	query := "SELECT * FROM user"
	userArray := []User{}
	err := db.Master.Select(&userArray, query)
	return userArray, err
}

// GetUser 사용자의 id로 조회
func (db *DBPool) GetUser(id int64) (User, error) {
	query := "SELECT * FROM user WHERE id=?"
	user := User{}
	err := db.Master.Get(&user, query, id)
	return user, err
}

// Auth 이름과 비밀번호를 대조하여 등록된 사용자인지 확인
func (db *DBPool) Auth(name string, passwd string) (User, error) {
	query := "SELECT * FROM user WHERE name=? AND passwd=?"
	user := User{}
	err := db.Master.Get(&user, query, name, passwd)
	return user, err
}

// SetUser 사용자 정보 등록
func (db *DBPool) SetUser(user *User) (sql.Result, error) {
	query := "INSERT INTO user (name, passwd) VALUES (:name, :passwd)"
	result, err := db.Master.NamedExec(query, user)
	return result, err
}

// PutUser 사용자 정보 수정
func (db *DBPool) PutUser(user *User) (sql.Result, error) {
	query := `
		UPDATE user
			SET name=:name,
				passwd=:passwd
			WHERE id=:id
	`
	result, err := db.Master.Exec(query, user)
	return result, err
}

// DeleteUser 사용자 정보 삭제
func (db *DBPool) DeleteUser(id int64) (sql.Result, error) {
	query := "DELTET FROM user WHERE id=?"
	result, err := db.Master.Exec(query, id)
	return result, err
}
