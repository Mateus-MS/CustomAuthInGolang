package api_users

import (
	"database/sql"
	"net/http"

	usersDB_actions "github.com/Mateus-MS/CustomAuthInGolang/clusters/models/users/actions"
)

func CreateRoute(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := r.FormValue("username")
	pass := r.FormValue("password")

	if status := usersDB_actions.Create(user, pass, db); status == nil {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusConflict)
	}
}
