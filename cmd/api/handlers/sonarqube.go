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
			"code": 400,
			"msg":  param_error,
			"data": nil,
		})
		return
	}

	// etcd服务发现
	ok, err := etcd.Get("sonarqube")
	if !ok || err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  service_discovery_error,
			"data": nil,
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

	rsp, _ := rpc.CreateProject(&req)
	c.JSON(int(rsp.Code), gin.H{
		"code": rsp.Code,
		"msg":  rsp.Message,
		"data": rsp.Data,
	})
}

func GetProject(c *gin.Context) {
	var param GetProjectParam

	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  param_error,
			"data": nil,
		})
		return
	}

	// etcd服务发现
	ok, err := etcd.Get("sonarqube")
	if !ok || err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  service_discovery_error,
			"data": nil,
		})
		return
	}

	req := sonarqubepb.GetProjectReq{
		Username:    param.Username,
		Token:       param.Token,
		ProjectName: param.ProjectName,
	}

	rsp, _ := rpc.GetProject(&req)
	c.JSON(int(rsp.Code), gin.H{
		"code": rsp.Code,
		"msg":  rsp.Message,
		"data": rsp.Data,
	})
}

func GetProjectList(c *gin.Context) {
	var param GetProjectListParam

	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  param_error,
			"data": nil,
		})
		return
	}

	// etcd服务发现
	ok, err := etcd.Get("sonarqube")
	if !ok || err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  service_discovery_error,
			"data": nil,
		})
		return
	}

	req := sonarqubepb.GetProjectListReq{
		Username: param.Username,
		Token:    param.Token,
	}

	rsp, _ := rpc.GetProjectList(&req)
	c.JSON(int(rsp.Code), gin.H{
		"code": rsp.Code,
		"msg":  rsp.Message,
		"data": rsp.Data,
	})
}

func GetReport(c *gin.Context) {
	var param GetReportParam

	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  param_error,
			"data": nil,
		})
		return
	}

	// etcd服务发现
	ok, err := etcd.Get("sonarqube")
	if !ok || err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  service_discovery_error,
			"data": nil,
		})
		return
	}

	req := sonarqubepb.GetReportReq{
		Username: param.Username,
		Token:    param.Token,
		ReportId: int32(param.ReportID),
	}

	rsp, _ := rpc.GetReport(&req)
	c.JSON(int(rsp.Code), gin.H{
		"code": rsp.Code,
		"msg":  rsp.Message,
		"data": rsp.Data,
	})
}

func GetReportList(c *gin.Context) {
	var param GetReportListParam

	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  param_error,
			"data": nil,
		})
		return
	}

	// etcd服务发现
	ok, err := etcd.Get("sonarqube")
	if !ok || err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  service_discovery_error,
			"data": nil,
		})
		return
	}

	req := sonarqubepb.GetReportListReq{
		Username: param.Username,
		Token:    param.Token,
	}

	rsp, _ := rpc.GetReportList(&req)
	c.JSON(int(rsp.Code), gin.H{
		"code": rsp.Code,
		"msg":  rsp.Message,
		"data": rsp.Data,
	})
}

func GetIssue(c *gin.Context) {
	var param GetIssueParam

	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  param_error,
			"data": nil,
		})
		return
	}

	// etcd服务发现
	ok, err := etcd.Get("sonarqube")
	if !ok || err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  service_discovery_error,
			"data": nil,
		})
		return
	}

	req := sonarqubepb.GetIssueReq{
		Username:    param.Username,
		Token:       param.Token,
		ProjectName: param.ProjectName,
	}

	rsp, _ := rpc.GetIssue(&req)
	c.JSON(int(rsp.Code), gin.H{
		"code": rsp.Code,
		"msg":  rsp.Message,
		"data": rsp.Data,
	})
}
