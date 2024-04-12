package main

import (
	"github.com/dkZzzz/quality_hub/cmd/api/router"
	"github.com/dkZzzz/quality_hub/db/mysql"
	"github.com/dkZzzz/quality_hub/srv/usersrv"
	"github.com/gin-gonic/gin"
)

func main() {
	// 启动Gin，注册路由
	go usersrv.Init_server()
	go usersrv.Init_client()
	_, err := mysql.InitDB()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	router.RegisterRoutes(r)
	r.Run("localhost:8080")
}
