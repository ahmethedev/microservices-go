package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = "8080"

type Config struct{}

func main() {
	app := Config{}
	log.Printf("Starting service on port: %s", PORT)

	//define https server

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", PORT),
		Handler: app.routes(),
	}

	//start server

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
