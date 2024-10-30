package auth

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"trucode.app/user_admin_api/database"
	"trucode.app/user_admin_api/models"
	"trucode.app/user_admin_api/shared"
)

func Register(c *gin.Context){
	var userInput models.UserInput
	c.BindJSON(&userInput)
	user := models.User{
		Username: userInput.Username,
		Password: userInput.Password,
	}
	dbConn := database.CreateDbConnection()
	if tx := dbConn.Create(&user); tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrDuplicatedKey) {
			c.JSON(http.StatusConflict, gin.H{"error":"Username already exists"})
			return
		}
	}
	c.JSON(http.StatusCreated, gin.H{
		"username":user.Username,
		"id":user.ID,
	})
}

func Login(c *gin.Context){
	var input models.UserInput
	var user models.User

	c.BindJSON(&input)
	dbConn := database.CreateDbConnection()
	dbConn.Where("username=?",input.Username).Find(&user)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error":"Invalid credentials"})
		return
	}

	sessionToken := uuid.NewV5(uuid.UUID{}, "session").String()

	session := shared.Session{
		Uid: user.ID,
		ExpiryTime: time.Now().Add(1 * time.Minute),
	}

	shared.Sessions[sessionToken] = session

	c.SetCookie("session", sessionToken, 1 * 60, "/", "localhost", false, false)
	c.JSON(http.StatusOK, gin.H{"session":sessionToken})
}