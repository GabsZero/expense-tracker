package main

import (
	"github.com/gabszero/expense-tracker/pkg/routes"
)

func main() {
	router := routes.ExpenseRoutes{}
	router.RegisterRoutes()

}
