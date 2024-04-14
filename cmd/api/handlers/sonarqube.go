package handlers

import (
	"net/http"

	"github.com/dkZzzz/quality_hub/cmd/api/rpc"
	"github.com/dkZzzz/quality_hub/db/etcd"
	"github.com/dkZzzz/quality_hub/proto/sonarqubepb"
	"github.com/gin-gonic/gin"
)

// 从Gin接受参数，发送给rpc服务
func CreateProject(c *gin.Context) {
	var param CreateProjectParam

	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "参数错误",
			"error": err.Error(),
		})
		return
	}
	// etcd服务发现
	ok, err := etcd.Get("sonarqube")
	if !ok || err != nil {
		c.JSON(500, gin.H{
			"msg": "服务发现失败",
		})
		return
	}

	req := sonarqubepb.CreateProjectReq{
		Username:    param.Username,
		ProjectName: param.ProjectName,
		BranchName:  param.BranchName,
		Url:         param.Url,
		Token:       param.Token,
	}
	rpc.CreateProject(&req)
	c.JSON(http.StatusOK, gin.H{})
}
