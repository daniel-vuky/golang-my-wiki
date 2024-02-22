package repository

import (
	"database/sql"
	"github.com/daniel-vuky/golang-my-wiki/model"
)

type UserRepository struct {
	Db *sql.DB
}

func (repository *UserRepository) GetUser(user *model.User) error {
	exec := "SELECT user_id, username, email, password FROM user WHERE email = ?"
	return repository.Db.QueryRow(exec, user.Email).Scan(
		&user.UserId,
		&user.UserName,
		&user.Email,
		&user.Password,
	)
}

func (repository *UserRepository) CreateUser(user *model.User) error {
	result, insertErr := repository.Db.Exec(
		"INSERT INTO user (username, email, password) VALUES (?, ?, ?)",
		user.UserName,
		user.Email,
		user.Password,
	)
	if insertErr != nil {
		return insertErr
	}
	lastInsertedUserId, lastInsertUserIdErr := result.LastInsertId()
	if lastInsertUserIdErr != nil {
		return lastInsertUserIdErr
	}
	user.UserId = uint64(lastInsertedUserId)

	return nil
}
