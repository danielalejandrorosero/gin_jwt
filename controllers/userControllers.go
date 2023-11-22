package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/danielalejandrorosero/jwt_gin/initialize"
	"github.com/danielalejandrorosero/jwt_gin/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var data struct {
		email    string
		password string
	}

	if c.Bind(&data) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "400 bad request",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(data.password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Faild to hash password",
		})
		return
	}

	user := models.User{Email: data.email, Password: string(hash)}
	result := initialize.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Faild to create user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func Login(c *gin.Context) {
	var data struct {
		email    string
		password string
	}

	if c.Bind(&data) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "400 bad request",
		})
		return
	}

	var user models.User
	initialize.DB.First(&user, "email = ?", data.email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "No user found",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "No user found",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "faild to create token",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}
