package controllers

import (
	"context"
	"net/http"
	"os"

	"github.com/RafiMoreno/be-linkby.me/initializers"
	"github.com/RafiMoreno/be-linkby.me/models"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
)

// Get Profile             godoc
// @Summary      Get Profile
// @Description  Get a user profile based on username
// @Tags         profile
// @Produce      json
// @Success      200
// @Router       /profile/:username [get]
func GetProfile(c *gin.Context) {
	username := c.Param("username")
	var user models.User
	initializers.DB.Where("username = ?", username).Preload("Profile").First(&user)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})

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
func EditProfile(c *gin.Context) {
	var body struct {
		DisplayName    string
		PrimaryColor   string
		SecondaryColor string
		Description    string
	}

	c.Bind(&body)

	username := c.Param("username")
	currUser, _ := c.Get("user")
	currUsername := currUser.(models.User).Username

	if username != currUsername {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized User"})

		return
	}

	var user models.User

	initializers.DB.Where("username = ?", username).Preload("Profile").First(&user)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})

		return
	}

	result := initializers.DB.Model(&user.Profile).Updates(
		models.Profile{
			DisplayName:    body.DisplayName,
			PrimaryColor:   body.PrimaryColor,
			SecondaryColor: body.SecondaryColor,
			Description:    body.Description,
		},
	)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})

		return
	}

	c.JSON(200, gin.H{"profile": user.Profile})
}

func UploadImage(c * gin.Context){
	username := c.Param("username")
	currUser, _ := c.Get("user")
	currUsername := currUser.(models.User).Username

	if username != currUsername {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized User"})

		return
	}

	fileHeader, _ := c.FormFile("image")
	file, _ := fileHeader.Open()

	ctx := context.Background()

	cld, err := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Image Storage Error"})

		return
	}

	resp, _ := cld.Upload.Upload(ctx, file, uploader.UploadParams{})

	var user models.User

	initializers.DB.Where("username = ?", username).Preload("Profile").First(&user)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})

		return
	}

	result := initializers.DB.Model(&user.Profile).Updates(
		models.Profile{
			DisplayPicture: resp.SecureURL,
		},
	)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update image"})

		return
	}

	c.JSON(200, gin.H{"displayPicture": user.Profile.DisplayPicture})
}