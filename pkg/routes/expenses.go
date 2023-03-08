package routes

import "github.com/gin-gonic/gin"

type ExpenseRoutes struct {
}

func (e *ExpenseRoutes) RegisterRoutes() {
	router := gin.Default()
	v1 := router.Group("/api/v1/expenses")
	{
		v1.GET("", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "hello from expenses",
			})
		})
	}
	router.Run() // listen and serve on 0.0.0.0:8080
}
