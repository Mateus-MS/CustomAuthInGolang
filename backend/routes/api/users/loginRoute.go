package api_users

import (
	"database/sql"
	"net/http"
	"time"

	usersDB_actions "github.com/Mateus-MS/CustomAuthInGolang/clusters/models/users/actions"
	"github.com/Mateus-MS/CustomAuthInGolang/utils"
)

func LoginRoute(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	//Look for the request name in the DB
	user, err := usersDB_actions.Search(username, db)

	//IF users does not exists AND If password received miss match the one in DB
	if err != nil || !utils.CheckPassordHash(password, user.Passwordhash) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	sessionToken := utils.GenerateToken(32)
	csrfToken := utils.GenerateToken(32)

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
		Secure:   true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    csrfToken,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: false,
		Secure:   true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "username",
		Value:    username,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
		Secure:   true,
	})

	if err := usersDB_actions.AtualizeTokens(username, sessionToken, csrfToken, db); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
