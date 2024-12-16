package usersDB_actions

import (
	"database/sql"
	"errors"

	"github.com/Mateus-MS/CustomAuthInGolang/utils"
)

func Create(username, password string, db *sql.DB) error {
	if _, err := Search(username, db); err == nil {
		return errors.New("user already exist")
	}

	hashPass, _ := utils.HashPassword(password)

	if _, err := db.Exec("INSERT INTO tb_users (username, password_hash, session_token, csrf_token) VALUES ($1, $2, '', '')", username, hashPass); err != nil {
		return errors.New("error while saving in usersDB")
	}

	return nil
}
