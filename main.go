package main

import (
	"fmt"
	"net/http"

	"golang.org/x/crypto/acme/autocert"
)

func main() {
	router := http.NewServeMux()

	startServer(router)
}

func startServer(router *http.ServeMux) {
	certManager := &autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		Cache:      autocert.DirCache("certs"),
		HostPolicy: autocert.HostWhitelist("cubonauta.com"),
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
