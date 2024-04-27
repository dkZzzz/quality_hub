package handlers

import (
	"net/http"

	"github.com/dkZzzz/quality_hub/cmd/api/rpc"
	"github.com/dkZzzz/quality_hub/db/etcd"
	"github.com/dkZzzz/quality_hub/proto/chatpb"
	"github.com/gin-gonic/gin"
)

// 从GIN接受参数，发送给rpc服务
func SentSingleIssue(c *gin.Context) {
	var param SentSingleIssueParam

	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  param_error,
			"data": nil,
		})
		return
	}

	// etcd服务发现
	ok, err := etcd.Get("chat")
	if !ok || err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  service_discovery_error,
			"data": nil,
		})
		return
	}

	req := chatpb.SentSingleIssueReq{
		Username: param.Username,
		Token:    param.Token,
		IssueId:  int32(param.IssueID),
	}

	rsp, _ := rpc.SentSingleIssue(&req)
	c.JSON(int(rsp.Code), gin.H{
		"code": rsp.Code,
		"msg":  rsp.Message,
		"data": rsp.Data,
	})
}

func SentProjectIssue(c *gin.Context) {
	var param SentProjectIssueParam

	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  param_error,
			"data": nil,
		})
		return
	}

	// etcd服务发现
	ok, err := etcd.Get("chat")
	if !ok || err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  service_discovery_error,
			"data": nil,
		})
		return
	}

	req := chatpb.SentProjectIssueReq{
		Username:    param.Username,
		Token:       param.Token,
		ProjectName: param.ProjectName,
	}

	rsp, _ := rpc.SentProjectIssue(&req)
	c.JSON(int(rsp.Code), gin.H{
		"code": rsp.Code,
		"msg":  rsp.Message,
		"data": rsp.Data,
	})
}
