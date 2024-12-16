package usersDB_actions

import (
	"database/sql"
)

func AtualizeTokens(username, sessionToken, csrfToken string, db *sql.DB) error {
	//Test if the user do exists in DB
	if _, err := Search(username, db); err != nil {
		return err
	}

	//Try update the values from the user
	if _, err := db.Exec("UPDATE tb_users SET session_token = $1, csrf_token = $2 WHERE username = $3", sessionToken, csrfToken, username); err != nil {
		return err
	}

	return nil
}
