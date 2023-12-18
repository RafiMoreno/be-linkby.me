package controllers

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/RafiMoreno/be-linkby.me/initializers"
	"github.com/RafiMoreno/be-linkby.me/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// SignUp             godoc
// @Summary      Sign Up
// @Description  Create user using username and password
// @Tags         auth
// @Produce      json
// @Success      200
// @Router       /sign-up [post]
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})

		return
	}

	var existingUser models.User

	initializers.DB.First(&existingUser, "username= ?", body.Username)

	if existingUser.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username " + body.Username + " already exists"})

		return
	}

	user := models.User{
		Username: body.Username,
		Password: string(hashed_pw),
		Profile:  models.Profile{DisplayName: body.Username},
	}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})

		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully created user"})
}

// Login             godoc
// @Summary      Login
// @Description  Generate JWT for authentication
// @Tags         auth
// @Produce      json
// @Success      200
// @Router       /login [post]
func Login(c *gin.Context) {
	var body struct {
		Username string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to fetch body"})

		return
	}
	var user models.User
	initializers.DB.First(&user, "username= ?", body.Username)

	if user.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})

		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})

		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})

		return
	}
	isSecure, err := strconv.ParseBool(os.Getenv("COOKIE_SECURE"))

	if err != nil {
		isSecure = true
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("token", tokenString, 3600*24*30, "/", "", isSecure, true)

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

// Validate             godoc
// @Summary      Validate
// @Description  Validate authentication
// @Tags         auth
// @Produce      json
// @Success      200
// @Router       /validate [get]
func Validate(c *gin.Context) {
	user, _ := c.Get("user")
	username := user.(models.User).Username

	c.JSON(http.StatusOK, gin.H{"message": "User is logged in", "username": username})
}
