package controllers

import (
	"net/http"
	"github.com/RafiMoreno/be-linkby.me/initializers"
	"github.com/RafiMoreno/be-linkby.me/models"
	"github.com/gin-gonic/gin"
)

// Get Profile             godoc
// @Summary      Get Profile  
// @Description  Get a user profile
// @Tags         profile
// @Produce      json
// @Success      200
// @Router       /profile/:id [get]
func GetProfile(c *gin.Context) {
	username := c.Param("username")
	var user models.User
	initializers.DB.Where("username = ?", username).Preload("Profile").First(&user)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})

		return
	}

	c.JSON(200, gin.H{ "profile" : user.Profile })
}

// func EditProfile(c *gin.Context) {
// 	var body struct{
// 		DisplayName string
// 		PrimaryColor string
// 		SecondaryColor string
// 		Description string
// 		DisplayPicture string
// 	}

// 	c.Bind(&body)

// 	username := c.Param("username")
// 	var user models.User
// 	var profile models.Profile
// 	initializers.DB.Model(&user).Where("username = ?", username).Association("Profile").Find(&profile)

// }