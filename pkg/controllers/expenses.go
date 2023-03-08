package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExpenseController struct {
}

func (ec *ExpenseController) Index(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello from expense controller")
}
