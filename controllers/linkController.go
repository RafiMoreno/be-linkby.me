package controllers

import (
	"net/http"

	"github.com/RafiMoreno/be-linkby.me/initializers"
	"github.com/RafiMoreno/be-linkby.me/models"
	"github.com/gin-gonic/gin"
)

// Create Link             godoc
// @Summary      Create Link
// @Description  Create a link iitem for a profile owned by a user
// @Tags         link
// @Produce      json
// @Success      200
// @Router       /profile/:username/link-create [get]
func CreateLink(c *gin.Context) {
	var body struct {
		Url        string
		Title      string
		ClickCount string
		IconUrl    string
	}

	c.Bind(&body)

	username := c.Param("username")
	currUser, _ := c.Get("user")
	currUsername := currUser.(models.User).Username

	if username != currUsername {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized User"})

		return
	}

	var user models.User

	initializers.DB.Where("username = ?", username).Preload("Profile").First(&user)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})

		return
	}

	c.JSON(200, gin.H{"profile": user.Profile})
}

// Edit Profile             godoc
// @Summary      Edit Profile
// @Description  Edit a user profile based on username
// @Tags         profile
// @Produce      json
// @Success      200
// @Router       /profile/:username [put]
// func EditLink(c *gin.Context) {
// 	var body struct {
// 		DisplayName    string
// 		PrimaryColor   string
// 		SecondaryColor string
// 		Description    string
// 		DisplayPicture string
// 	}

// 	c.Bind(&body)

// 	username := c.Param("username")
// 	currUser, _ := c.Get("user")
// 	currUsername := currUser.(models.User).Username

// 	if username != currUsername {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized User"})

// 		return
// 	}

// 	var user models.User

// 	initializers.DB.Where("username = ?", username).Preload("Profile").First(&user)

// 	if user.ID == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})

// 		return
// 	}

// 	initializers.DB.Model(&user.Profile).Updates(
// 		models.Profile{
// 			DisplayName:    body.DisplayName,
// 			PrimaryColor:   body.PrimaryColor,
// 			SecondaryColor: body.SecondaryColor,
// 			Description:    body.Description,
// 			DisplayPicture: body.DisplayPicture},
// 	)

// 	c.JSON(200, gin.H{"profile": user.Profile})
// }
