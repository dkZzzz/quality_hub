package main

import (
	"github.com/dkZzzz/quality_hub/cmd/api/router"
	"github.com/dkZzzz/quality_hub/db/mysql"
	"github.com/dkZzzz/quality_hub/srv/chatsrv"
	"github.com/dkZzzz/quality_hub/srv/sonarqubesrv"
	"github.com/dkZzzz/quality_hub/srv/usersrv"
	"github.com/gin-gonic/gin"
)

func main() {
	go usersrv.Init_server()
	go usersrv.Init_client()

	go sonarqubesrv.Init_server()
	go sonarqubesrv.Init_client()

	go chatsrv.Init_server()
	go chatsrv.Init_client()

	_, err := mysql.InitDB()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	router.RegisterRoutes(r)
	r.Run("localhost:8080")
}
