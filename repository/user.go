package repository

import (
	"database/sql"
	"quiz-golang/structs"
)

func Register(db *sql.DB, user structs.User) (err error) {
	sql := "INSERT INTO users(username, password, created_by, modified_by) VALUES ($1, $2, $3, $4)"
	errs := db.QueryRow(sql, user.Username, user.Password, user.Created_by, user.Modified_by)
	return errs.Err()
}
