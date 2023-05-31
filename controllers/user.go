package controllers

import (
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
	"net/http"
	// "skid_go/models"
	"time"
)

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"not null" json:"username"`
	Password string `gorm:"not null" json:"-"`
}
type UserController struct {
	db *gorm.DB // Your database connection
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{
		db: db,
	}
}

func (uc *UserController) Authenticate(ctx *gin.Context) {
	var credentials struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// Bind the request body to the credentials struct
	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
		return
	}

	// Check if the username exists in the database
	var user User
	err := uc.db.Where("username = ?", credentials.Username).First(&user).Error
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	// Generate and sign the JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(2 * time.Minute).Unix()

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte("thisismysecretkey"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Return the access token
	ctx.JSON(http.StatusOK, gin.H{"access_token": tokenString})
}
