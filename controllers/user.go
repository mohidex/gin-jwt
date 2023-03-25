package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohidex/identity-service/models"
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

func (u UserController) Register(c *gin.Context) {
	var input models.RegistrationInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := models.User{
		Name:     input.Name,
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
		PhotoURL: input.PhotoURL,
	}
	savedUser, err := user.Save()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"user": savedUser,
	})
}
