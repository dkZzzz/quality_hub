package sonarqube

import (
	"context"

	"github.com/dkZzzz/quality_hub/proto/sonarqubepb"
)

// rpc服务具体接口业务逻辑

type SonarQubeServerImpl struct {
	sonarqubepb.UnimplementedSonarQubeServer
}

func (s *SonarQubeServerImpl) CreateProject(ctx context.Context, in *sonarqubepb.CreateProjectReq) (*sonarqubepb.CreateProjectResp, error) {
	// 业务逻辑

	return &sonarqubepb.CreateProjectResp{}, nil
}
