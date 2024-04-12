package router

import (
	"github.com/dkZzzz/quality_hub/cmd/api/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	// 用户模块
	user := r.Group("/user")
	{
		user.POST("/register", handlers.Register)

		user.POST("/login", handlers.Login)

		user.POST("/logout", handlers.Logout)

		user.POST("/modify_username", handlers.ModifyUsername)

		user.POST("/modify_email", handlers.ModifyEmail)

		user.POST("/modify_password", handlers.ModifyPassword)
	}

}
