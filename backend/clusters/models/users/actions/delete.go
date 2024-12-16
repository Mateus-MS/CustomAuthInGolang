package usersDB_actions

import (
	"database/sql"
	"errors"
)

func Delete(username string, db *sql.DB) error {
	if _, err := Search(username, db); err != nil {
		return errors.New("user do not exists")
	}

	if _, err := db.Exec("DELETE FROM tb_users WHERE username = $1;", username); err != nil {
		return errors.New("error while deleting user on DB")
	}

	return nil
}
