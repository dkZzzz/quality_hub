package sonarqube

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/dkZzzz/quality_hub/config"
	"github.com/dkZzzz/quality_hub/db/mysql"
	"github.com/dkZzzz/quality_hub/db/redis"
	"github.com/dkZzzz/quality_hub/pkg/clone"
	"github.com/dkZzzz/quality_hub/pkg/sonarapi"
	"github.com/dkZzzz/quality_hub/proto/sonarqubepb"
)

var (
	token_error          = "token验证失败"
	create_project_error = "创建项目失败"
	scanning_error       = "扫描代码失败"
	create_project_succ  = "创建项目成功"
)

// rpc服务具体接口业务逻辑

type SonarQubeServerImpl struct {
	sonarqubepb.UnimplementedSonarQubeServer
}

func (s *SonarQubeServerImpl) CreateProject(ctx context.Context, in *sonarqubepb.CreateProjectReq) (*sonarqubepb.CreateProjectResp, error) {
	if !redis.JWTMatch(ctx, in.Username, in.Token) {
		return &sonarqubepb.CreateProjectResp{
			Code:    401,
			Message: token_error,
		}, nil
	}

	// 向sonarqube发送请求
	err := sonarapi.CreateProject(in.ProjectName, in.ProjectName)
	if err != nil {
		return &sonarqubepb.CreateProjectResp{
			Code:    500,
			Message: create_project_error,
		}, nil
	}

	// 在本地库创建记录
	err = mysql.CreateProject(ctx, in.Username, in.ProjectName, in.BranchName, in.Url, in.Token)
	if err != nil {
		return &sonarqubepb.CreateProjectResp{
			Code:    500,
			Message: create_project_error,
		}, nil
	}

	tmp := strings.Split(in.Url, "/")
	pName := tmp[len(tmp)-1]
	pName = pName[:len(pName)-4]
	path := config.Cfg.CodeStorePath + pName

	// 克隆代码
	err = clone.Clone(in.Url, pName)
	if err != nil {
		log.Println(err)
	}

	// 创建token
	token, err := sonarapi.GenerateProjectToken(in.ProjectName, in.ProjectName)
	if err != nil {
		return &sonarqubepb.CreateProjectResp{
			Code:    500,
			Message: create_project_error,
		}, nil
	}

	// 扫描代码
	err = Scan(path, in.ProjectName, token)
	if err != nil {
		return &sonarqubepb.CreateProjectResp{
			Code:    500,
			Message: scanning_error,
		}, nil
	}

	return &sonarqubepb.CreateProjectResp{
		Code:    200,
		Message: create_project_succ,
	}, nil
}

func Scan(path, projectKey, token string) error {
	cmd := exec.Command("sonar-scanner",
		fmt.Sprintf("-Dsonar.projectKey=%s", projectKey),
		fmt.Sprintf("-Dsonar.sources=%s", path),
		fmt.Sprintf("-Dsonar.projectBaseDir=%s", path),
		fmt.Sprintf("-Dsonar.host.url=%s", config.Cfg.SonarHost),
		fmt.Sprintf("-Dsonar.token=%s", token))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
