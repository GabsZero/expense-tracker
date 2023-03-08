package routes

import (
	"github.com/gabszero/expense-tracker/pkg/controllers"
	"github.com/gin-gonic/gin"
)

type ExpenseRoutes struct {
}

func (e *ExpenseRoutes) RegisterRoutes() {
	router := gin.Default()
	expenseController := controllers.ExpenseController{}
	v1 := router.Group("/api/v1/expenses")
	{
		v1.GET("", expenseController.Index)
	}
	router.Run() // listen and serve on 0.0.0.0:8080
}
