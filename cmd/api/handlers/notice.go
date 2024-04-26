package handlers

import (
	"net/http"

	"github.com/dkZzzz/quality_hub/cmd/api/rpc"
	"github.com/dkZzzz/quality_hub/db/etcd"
	"github.com/dkZzzz/quality_hub/proto/noticepb"
	"github.com/gin-gonic/gin"
)

// 从GIN接受参数，发送给rpc服务
func GetSingleAdvice(c *gin.Context) {
	var param GetSingleAdviceParam

	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  param_error,
			"data": nil,
		})
		return
	}

	// etcd服务发现
	ok, err := etcd.Get("notice")
	if !ok || err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  service_discovery_error,
			"data": nil,
		})
		return
	}

	req := noticepb.GetSingleAdviceReq{
		Username: param.Username,
		Token:    param.Token,
		AdviceId: int32(param.AdviceID),
	}

	rsp, _ := rpc.GetSingleAdvice(&req)
	c.JSON(int(rsp.Code), gin.H{
		"code": rsp.Code,
		"msg":  rsp.Message,
		"data": rsp.Data,
	})
}

func GetProjectAdvice(c *gin.Context) {
	var param GetProjectAdviceParam

	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  param_error,
			"data": nil,
		})
		return
	}

	// etcd服务发现
	ok, err := etcd.Get("notice")
	if !ok || err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  service_discovery_error,
			"data": nil,
		})
		return
	}

	req := noticepb.GetProjectAdviceReq{
		Username:    param.Username,
		Token:       param.Token,
		ProjectName: param.ProjectName,
	}

	rsp, _ := rpc.GetProjectAdvice(&req)
	c.JSON(int(rsp.Code), gin.H{
		"code": rsp.Code,
		"msg":  rsp.Message,
		"data": rsp.Data,
	})
}
