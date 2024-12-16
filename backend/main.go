package main

import (
	"fmt"
	"net/http"

	usersDB "github.com/Mateus-MS/CustomAuthInGolang/clusters/models/users"
	api_users "github.com/Mateus-MS/CustomAuthInGolang/routes/api/users"
	"github.com/Mateus-MS/CustomAuthInGolang/utils"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/api/users/check", func(w http.ResponseWriter, r *http.Request) {
		utils.HandleWithCORS(func(w http.ResponseWriter, r *http.Request) {
			api_users.CheckRoute(w, r, usersDB.GetInstance())
		})(w, r)
	})

	router.HandleFunc("/api/users/create", func(w http.ResponseWriter, r *http.Request) {
		utils.HandleWithCORS(func(w http.ResponseWriter, r *http.Request) {
			api_users.CreateRoute(w, r, usersDB.GetInstance())
		})(w, r)
	})

	router.HandleFunc("/api/users/login", func(w http.ResponseWriter, r *http.Request) {
		utils.HandleWithCORS(func(w http.ResponseWriter, r *http.Request) {
			api_users.LoginRoute(w, r, usersDB.GetInstance())
		})(w, r)
	})

	router.HandleFunc("/api/users/logout", func(w http.ResponseWriter, r *http.Request) {
		utils.HandleWithCORS(func(w http.ResponseWriter, r *http.Request) {
			api_users.LogoutRoute(w, r, usersDB.GetInstance())
		})(w, r)
	})

	startServer(router)
}

func startServer(router *http.ServeMux) {
	certManager := &autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		Cache:      autocert.DirCache("certs"),
		HostPolicy: autocert.HostWhitelist("localhost"),
	}

	go func() {
		httpServer := &http.Server{
			Addr:    ":80",
			Handler: certManager.HTTPHandler(nil),
		}
		fmt.Println("Starting HTTP server on port 80 for certificate challenges and redirection to HTTPS")
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("HTTP server error: %s", err)
		}
	}()

	httpsServer := &http.Server{
		Addr:      ":443",
		Handler:   router,
		TLSConfig: certManager.TLSConfig(),
	}

	fmt.Println("Starting HTTPS server on port 443")
	if err := httpsServer.ListenAndServeTLS("", ""); err != nil {
		fmt.Println("HTTPS server error:", err)
	}
}
