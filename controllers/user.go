package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohidex/identity-service/models"
	"github.com/mohidex/identity-service/utils"
)

type UserController struct{}

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

func (u UserController) Login(c *gin.Context) {
	var input models.LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	user, err := models.FindUserByUsername(input.Username)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if err := user.ValidatePassword(input.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jwt, err := utils.GenerateJwt(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"jwt": jwt})

}
