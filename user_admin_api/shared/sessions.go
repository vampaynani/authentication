package shared

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"trucode.app/user_admin_api/database"
	"trucode.app/user_admin_api/models"
)

type Session struct {
	Uid uint
	ExpiryTime time.Time
}

var Sessions = map[string]Session{}

func SetCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func AuthenticateSession() gin.HandlerFunc {
	return func(c *gin.Context){
		sessionToken, err := c.Cookie("session")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "you don't have permission to access this resource"})
			return
		}

		userData := Sessions[sessionToken]
		if userData.ExpiryTime.Before(time.Now()){
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error":"session expired"})
			return
		}
	
		var user models.User
		dbConn := database.CreateDbConnection()
		tx := dbConn.Where("id=?", userData.Uid).Find(&user)
		if tx.Error != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": tx.Error.Error()})
			return
		}

		c.Set("authorizedUser", user)
		
		c.Next()
	}
}