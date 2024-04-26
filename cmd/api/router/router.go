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

	// sonarqube模块
	sonarqube := r.Group("/sonarqube")
	{
		sonarqube.POST("/create_project", handlers.CreateProject)

		sonarqube.POST("/get_project", handlers.GetProject)

		sonarqube.POST("/get_project_list", handlers.GetProjectList)

		sonarqube.POST("/get_report", handlers.GetReport)

		sonarqube.POST("/get_report_list", handlers.GetReportList)

		sonarqube.POST("/get_issue", handlers.GetIssue)

	}
}
