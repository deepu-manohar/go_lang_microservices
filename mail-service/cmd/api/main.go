package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

type Config struct {
	Mailer Mail
}

func main() {
	app := Config{
		Mailer: CreateMailer(),
	}
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}
	log.Println("Starting mail server ..")
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic("Failed to start server ", err)
	}
}
