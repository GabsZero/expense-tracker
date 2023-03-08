package main

import (
	"log"

	"github.com/gabszero/expense-tracker/pkg/infra/repositories"
	"github.com/gabszero/expense-tracker/pkg/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database := repositories.Database{}
	database.AutoMigrate()

	router := routes.ExpenseRoutes{}
	router.RegisterRoutes()

}
