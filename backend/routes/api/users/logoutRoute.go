package api_users

import (
	"database/sql"
	"net/http"
	"time"

	usersDB_actions "github.com/Mateus-MS/CustomAuthInGolang/clusters/models/users/actions"
)

func LogoutRoute(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if isAuthorized := usersDB_actions.Authorize(r, db); !isAuthorized {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Retrieve the username from the cookie
	usernameCookie, err := r.Cookie("username")
	if err != nil || usernameCookie == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	username := usernameCookie.Value

	//Clear cookies
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour * 24),
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour * 24),
		HttpOnly: false,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "username",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour * 24),
		HttpOnly: true,
	})

	if err := usersDB_actions.AtualizeTokens(username, "", "", db); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
