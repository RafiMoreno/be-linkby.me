package controllers

import (
	"net/http"

	"example/be-linkby.me/initializers"
	"example/be-linkby.me/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var body struct {
		Username string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to fetch body"})

		return
	}

	hashed_pw, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to hash password"})

		return
	}

	user := models.User{
		Username: body.Username,
		Password: string(hashed_pw),
	}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create user"})

		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
