package main

import (
	"authentication/data"
	"authentication/services"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = "80"

type Config struct {
	DB            *sql.DB
	Models        data.Models
	authService   services.AuthService
	loggerService services.LoggerService
}

func main() {
	log.Println("Connecting to DB..")
	conn := connectToDB()
	if conn == nil {
		log.Panic("Failed to connect to DB after multiple retries")
	}
	log.Println("Connected to DB..")
	log.Println("Starting Authentication Service...")
	models := data.New(conn)
	app := Config{
		DB:            conn,
		Models:        models,
		authService:   services.NewAuthService(models),
		loggerService: services.NewLoggerService(),
	}
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic("Failed to start server ", err)
	}
	log.Println("Started Authentication Serivce...")
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Panic("Failed to connect to DB ", err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Println("Failed to ping to DB ", err)
		return nil, err
	}
	return db, nil
}

func connectToDB() *sql.DB {
	maxRetries := 10
	dsn := os.Getenv("DSN")
	for i := 0; i < maxRetries; i++ {
		conn, err := openDB(dsn)
		if err != nil {
			log.Println("Failed to connect to DB, retrying", err)
			time.Sleep(2 * time.Second)
		} else {
			return conn
		}
	}
	return nil
}
