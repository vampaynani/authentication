package shared

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"trucode.app/user_admin_api/database"
	"trucode.app/user_admin_api/models"
)

func Cors() gin.HandlerFunc {
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
	return func (c *gin.Context){
		sessionToken, err := c.Cookie("session")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error":err.Error()})
			return
		}

		userData := Sessions[sessionToken]
		if userData.ExpiryTime.Before(time.Now()){
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "expired session"})
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