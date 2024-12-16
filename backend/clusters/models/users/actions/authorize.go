package usersDB_actions

import (
	"database/sql"
	"net/http"
)

func Authorize(r *http.Request, db *sql.DB) bool {
	// Retrieve the username from the cookie
	usernameCookie, err := r.Cookie("username")
	if err != nil || usernameCookie == nil {
		return false
	}
	username := usernameCookie.Value

	user, err := Search(username, db)
	if err != nil {
		return false
	}

	sessionToken, err := r.Cookie("session_token")
	if err != nil || sessionToken.Value == "" || sessionToken.Value != user.SessionToken {
		return false
	}

	//if the tokens is not in the header, that is a signal that is the first request
	//a non tokenized request, the first request is validated with only the cookie
	//the subsequential request are all validated with the header only
	//NOTE: i don't know if this is a flaw security wise
	csrfTokenHeader := r.Header.Get("X-CSRF-Token")
	if csrfTokenHeader != user.CsrfToken || csrfTokenHeader == "" {

		csrfTokenCookie, err := r.Cookie("csrf_token")
		if err != nil || csrfTokenCookie.Value == "" || csrfTokenCookie.Value != user.CsrfToken {
			return false
		}
	}

	return true
}
