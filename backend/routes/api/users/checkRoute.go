package api_users

import (
	"database/sql"
	"net/http"

	usersDB_actions "github.com/Mateus-MS/CustomAuthInGolang/clusters/models/users/actions"
	"github.com/Mateus-MS/CustomAuthInGolang/utils"
)

func CheckRoute(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	username, err := utils.QueryFromURL[string]("username", r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = usersDB_actions.Search(username, db)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
