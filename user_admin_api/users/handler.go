package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"trucode.app/user_admin_api/database"
	"trucode.app/user_admin_api/models"
)

func GetAllUsers(c *gin.Context){
	var users []models.User
	dbConn := database.CreateDbConnection()
	dbConn.Find(&users)
	c.JSON(http.StatusOK, users)
}

func GetUserById(c *gin.Context){
	var foundUser models.User
	dbConn := database.CreateDbConnection()
	tx := dbConn.Where("id=?", c.Param("id")).Find(&foundUser)
	if tx.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, foundUser)
}

func GetMe(c *gin.Context){
	sessionUser, _ := c.Get("authorizedUser")
	c.JSON(http.StatusOK, sessionUser)
}

func CreateUser(c *gin.Context){
	var user models.User
	c.BindJSON(&user)
	dbConn := database.CreateDbConnection()
	dbConn.Create(&user)
	c.JSON(http.StatusCreated, user)
}

func DeleteUser(c *gin.Context){
	id := c.Param("id")
	dbConn := database.CreateDbConnection()
	dbConn.Delete(&models.User{}, id)
	c.JSON(http.StatusNoContent, gin.H{"id" : id})
}