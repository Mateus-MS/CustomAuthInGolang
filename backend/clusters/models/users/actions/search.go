package usersDB_actions

import (
	"database/sql"
	"fmt"

	usersDB "github.com/Mateus-MS/CustomAuthInGolang/clusters/models/users"
)

func Search(username string, db *sql.DB) (usersDB.User, error) {
	user := usersDB.User{}

	rows, err := db.Query("SELECT id, created_at, username, password_hash, session_token, csrf_token FROM tb_users WHERE username = $1", username)
	if err != nil {
		return user, fmt.Errorf("error executing user search query")
	}
	defer rows.Close()

	if !rows.Next() {
		return user, fmt.Errorf("user { %s } not finded", username)
	}

	if err := rows.Scan(&user.Id, &user.CreatedAt, &user.Username, &user.Passwordhash, &user.SessionToken, &user.CsrfToken); err != nil {
		return user, fmt.Errorf("error scanning queryed row")
	}

	return user, nil
}
