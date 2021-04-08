package main

import (
	"github.com/OJoklrO/forum/internal/routers"
	"log"
	"net/http"
)

// @title forum
// @version 1.0
// @description student forum api
func main() {
	router := routers.NewRouter()
	s := http.Server{
		Addr: ":8080",
		Handler: router,
	}
	log.Fatal(s.ListenAndServe())
}
