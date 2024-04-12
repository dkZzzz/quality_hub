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
