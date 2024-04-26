package rpc

import (
	"context"

	"github.com/dkZzzz/quality_hub/proto/sonarqubepb"
	"github.com/dkZzzz/quality_hub/srv/sonarqubesrv"
)

var (
	sonarqubeClient sonarqubepb.SonarQubeClient
)

// 调用客户端，给服务端发送请求

// CreateProject 创建项目
func CreateProject(req *sonarqubepb.CreateProjectReq) (*sonarqubepb.CreateProjectResp, error) {
	sonarqubeClient = sonarqubesrv.Client
	ctx := context.Background()
	rsp, err := sonarqubeClient.CreateProject(ctx, req)
	if err != nil {
		return nil, err
	} else {
		return rsp, nil
	}
}

// GetProject 获取项目
func GetProject(req *sonarqubepb.GetProjectReq) (*sonarqubepb.GetProjectResp, error) {
	sonarqubeClient = sonarqubesrv.Client
	ctx := context.Background()
	rsp, err := sonarqubeClient.GetProject(ctx, req)
	if err != nil {
		return nil, err
	} else {
		return rsp, nil
	}
}

// GetProjectList 获取项目列表
func GetProjectList(req *sonarqubepb.GetProjectListReq) (*sonarqubepb.GetProjectListResp, error) {
	sonarqubeClient = sonarqubesrv.Client
	ctx := context.Background()
	rsp, err := sonarqubeClient.GetProjectList(ctx, req)
	if err != nil {
		return nil, err
	} else {
		return rsp, nil
	}
}

// GetReport 获取报告
func GetReport(req *sonarqubepb.GetReportReq) (*sonarqubepb.GetReportResp, error) {
	sonarqubeClient = sonarqubesrv.Client
	ctx := context.Background()
	rsp, err := sonarqubeClient.GetReport(ctx, req)
	if err != nil {
		return nil, err
	} else {
		return rsp, nil
	}
}

func GetReportList(req *sonarqubepb.GetReportListReq) (*sonarqubepb.GetReportListResp, error) {
	sonarqubeClient = sonarqubesrv.Client
	ctx := context.Background()
	rsp, err := sonarqubeClient.GetReportList(ctx, req)
	if err != nil {
		return nil, err
	} else {
		return rsp, nil
	}
}

func GetIssue(req *sonarqubepb.GetIssueReq) (*sonarqubepb.GetIssueResp, error) {
	sonarqubeClient = sonarqubesrv.Client
	ctx := context.Background()
	rsp, err := sonarqubeClient.GetIssue(ctx, req)
	if err != nil {
		return nil, err
	} else {
		return rsp, nil
	}
}
