package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (u UserController) Retrieve(c *gin.Context) {
	if c.Param("id") != "" {
		c.JSON(http.StatusOK, gin.H{"message": "User founded!", "username": "Joe Black"})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
}
