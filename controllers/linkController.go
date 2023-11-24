package controllers

import (
	"net/http"

	"github.com/RafiMoreno/be-linkby.me/initializers"
	"github.com/RafiMoreno/be-linkby.me/models"
	"github.com/gin-gonic/gin"
)

// Create Link             godoc
// @Summary      Create Link
// @Description  Create a link item for a profile owned by a user
// @Tags         link
// @Produce      json
// @Success      200
// @Router       /profile/:username/link-create [get]
func CreateLink(c *gin.Context) {
	var body struct {
		Url     string
		Title   string
		IconUrl string
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

	var profile models.Profile

	initializers.DB.Where("ID = ?", user.ID).Preload("Links").First(&profile)

	if profile.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User profile not found"})
		return
	}

	link := models.Link{
		Url:       body.Url,
		Title:     body.Title,
		IconUrl:   body.IconUrl,
		ProfileID: profile.ID,
	}

	initializers.DB.Create(&link)

	initializers.DB.Where("ID = ?", user.ID).Preload("Links").First(&profile)

	c.JSON(200, gin.H{
		"username":  user.Username,
		"profileID": profile.ID,
		"links":     profile.Links,
	})

}

func GetLink(c *gin.Context) {
	var body struct {
		Url     string
		Title   string
		IconUrl string
	}

	c.Bind(&body)

	username := c.Param("username")

	var user models.User

	initializers.DB.Where("username = ?", username).Preload("Profile").First(&user)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	var profile models.Profile

	initializers.DB.Where("ID = ?", user.ID).Preload("Links").First(&profile)

	if profile.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User profile not found"})
		return
	}

	initializers.DB.Where("ID = ?", user.ID).Preload("Links").First(&profile)

	c.JSON(200, gin.H{
		"username":  user.Username,
		"profileID": profile.ID,
		"links":     profile.Links,
	})

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
