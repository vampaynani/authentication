package users

import (
	"github.com/gin-gonic/gin"
	"trucode.app/user_admin_api/shared"
)

func AddUserRoutes(r *gin.Engine){
	group := r.Group("/users")

	group.Use(shared.AuthenticateSession())
	
	group.GET("", GetAllUsers)
	group.GET("/me", GetMe)
	group.GET("/:id", GetUserById)
	group.POST("", CreateUser)
	group.DELETE("/:id", DeleteUser)
}