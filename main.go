package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/wawew/golang-simple-user/app"
	"github.com/wawew/golang-simple-user/repository"
	"github.com/wawew/golang-simple-user/usecase"
)

var db *sql.DB

func main() {
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	connectionString := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable", dbUser, dbName, dbPassword, dbHost, dbPort)

	var err error
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userRepository := repository.NewPostgresUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepository)
	handler := app.NewController(userUseCase)

	log.Fatal(handler.Listen(":3000"))
}
