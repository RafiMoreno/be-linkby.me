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
// @Router       /profile/:username/link [post]
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

// Create Link             godoc
// @Summary      Get Link
// @Description  Get link items for a profile owned by a user
// @Tags         link
// @Produce      json
// @Success      200
// @Router       /profile/:username/link [get]
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

// Update Link             godoc
// @Summary      Update Link
// @Description  Update a link item for a profile owned by a user
// @Tags         link
// @Produce      json
// @Success      200
// @Router       /profile/:username/link/:linkID [put]
func UpdateLink(c *gin.Context) {
	var body struct {
		Url     string
		Title   string
		IconUrl string
	}

	c.Bind(&body)

	username := c.Param("username")
	linkID := c.Param("linkID")

	currUser, _ := c.Get("user")
	currUsername := currUser.(models.User).Username

	if username != currUsername {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized User"})

		return
	}

	var user models.User

	initializers.DB.Where("username = ?", username).First(&user)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	var oldLink models.Link

	initializers.DB.Where("Profile_ID = ?", user.ID).Where("ID = ?", linkID).First(&oldLink)

	if oldLink.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Link not found"})
		return
	}

	link := models.Link{
		Url:     body.Url,
		Title:   body.Title,
		IconUrl: body.IconUrl,
	}

	initializers.DB.Model(&oldLink).Updates(&link)

	var profile models.Profile

	initializers.DB.Where("ID = ?", user.ID).Preload("Links").First(&profile)

	c.JSON(200, gin.H{
		"username":  user.Username,
		"profileID": profile.ID,
		"links":     profile.Links,
	})

}

// Delete Link             godoc
// @Summary      Delete Link
// @Description  Delete a link item for a profile owned by a user
// @Tags         link
// @Produce      json
// @Success      200
// @Router       /profile/:username/link/:linkID [delete]
func DeleteLink(c *gin.Context) {
	var body struct {
		Url     string
		Title   string
		IconUrl string
	}

	c.Bind(&body)

	username := c.Param("username")
	linkID := c.Param("linkID")

	currUser, _ := c.Get("user")
	currUsername := currUser.(models.User).Username

	if username != currUsername {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized User"})

		return
	}

	var user models.User

	initializers.DB.Where("username = ?", username).First(&user)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	initializers.DB.Where("Profile_ID = ?", user.ID).Delete(&models.Link{}, linkID)

	var profile models.Profile

	initializers.DB.Where("ID = ?", user.ID).Preload("Links").First(&profile)

	c.JSON(200, gin.H{
		"username":  user.Username,
		"profileID": profile.ID,
		"links":     profile.Links,
	})

}

// Increment Click Count           godoc
// @Summary      Increment Click Counter
// @Description  Increment click count of a link item for a profile owned by a user
// @Tags         link
// @Produce      json
// @Success      200
// @Router       /profile/:username/link/:linkID/increment-count [delete]
func IncrementCounter(c *gin.Context) {
	var body struct{}

	c.Bind(&body)

	username := c.Param("username")
	linkID := c.Param("linkID")

	var user models.User

	initializers.DB.Where("username = ?", username).First(&user)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	var oldLink models.Link

	initializers.DB.Where("Profile_ID = ?", user.ID).Where("ID = ?", linkID).First(&oldLink)

	if oldLink.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Link not found"})
		return
	}

	link := models.Link{
		ClickCount: oldLink.ClickCount + 1,
	}

	initializers.DB.Model(&oldLink).Updates(&link)

	var newLink models.Link
	initializers.DB.Where("Profile_ID = ?", user.ID).Where("ID = ?", linkID).First(&newLink)

	c.JSON(200, newLink)

}
