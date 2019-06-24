package service

import (
	"log"
	"project/configs"
	"project/validation"
	"net/http"
	"time"
)

func createServer() (server *http.Server) {

	server = &http.Server{
		Addr:         configs.ADDR_SERVER,
		IdleTimeout:  1000 * time.Millisecond,
		ReadTimeout:  1000 * time.Millisecond,
		WriteTimeout: 1000 * time.Millisecond,
	}

	return
}

func StartServer() {

	// creates a hanlder for the server
	h := createRouter()

	// creates an HTTP server
	server := createServer()

	//CORS
	c := createCORS()

	handler := c.Handler(h)

	// associates a handler to a server
	server.Handler = handler

	validation.CreateValidator()

	// starts the server and prints error messages
	log.Fatal(server.ListenAndServe())
}

func StopServer() {}
