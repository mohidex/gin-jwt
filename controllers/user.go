package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohidex/identity-service/auth"
	"github.com/mohidex/identity-service/db"
	"github.com/mohidex/identity-service/models"
)

type UserController struct {
	DB   db.Database
	Auth auth.Authenticator
}

func (uh UserController) Register(c *gin.Context) {
	var input models.RegistrationInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newUser := models.User{
		Name:     input.Name,
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
		PhotoURL: input.PhotoURL,
	}
	savedUser, err := uh.DB.SaveUser(context.Background(), &newUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user": savedUser.ToUserResponse(),
	})
}

func (uh UserController) Login(c *gin.Context) {
	var input models.LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	user, err := uh.DB.GetUserByUsername(context.Background(), input.Username)

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

	token, err := uh.Auth.GenerateToken(context.Background(), user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"jwt": token})

}

func (uh UserController) AutorizeToken(c *gin.Context) {
	requestUser := c.MustGet("user").(*models.RequestUser)
	user, err := uh.DB.GetUserByID(context.Background(), requestUser.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": user.ToUserResponse(),
	})
}
