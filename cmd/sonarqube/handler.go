package sonarqube

import (
	"context"
	"fmt"
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
	create_token_error   = "创建项目token失败"
	scanning_error       = "扫描代码失败"
	get_issue_error      = "获取issue失败"
	get_hotspot_error    = "获取hotspot失败"
	create_project_succ  = "创建项目成功"
)

// rpc服务具体接口业务逻辑

type SonarQubeServerImpl struct {
	sonarqubepb.UnimplementedSonarQubeServer
}

type Isu struct {
	Type        string `json:"type"`
	File        string `json:"file"`
	StartLine   int    `json:"start_line"`
	EndLine     int    `json:"end_line"`
	StartOffset int    `json:"start_offset"`
	EndOffset   int    `json:"end_offset"`
	Message     string `json:"message"`
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
	go func() {
		mysql.CreateProject(ctx, in.Username, in.ProjectName, in.BranchName, in.Url, in.Token)
	}()

	// 处理url
	tmp := strings.Split(in.Url, "/")
	pName := tmp[len(tmp)-1]
	pName = pName[:len(pName)-4]
	path := config.Cfg.CodeStorePath + pName

	// 克隆代码
	err = clone.Clone(in.Url, pName)
	if err != nil {
		return &sonarqubepb.CreateProjectResp{
			Code:    500,
			Message: create_project_error,
		}, nil
	}

	// 创建token
	token, err := sonarapi.GenerateProjectToken(in.ProjectName, in.ProjectName)
	if err != nil {
		return &sonarqubepb.CreateProjectResp{
			Code:    500,
			Message: create_token_error,
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

	// 获取issue到本地库
	response, err := sonarapi.GetIssueByProject(in.ProjectName)
	if err != nil {
		return &sonarqubepb.CreateProjectResp{
			Code:    500,
			Message: get_issue_error,
		}, nil
	}

	go func() {
		Isus := ParseIssue(response)
		for _, isu := range Isus {
			mysql.CreateIssue(ctx, in.ProjectName, isu.Type, isu.File, isu.StartLine, isu.EndLine, isu.StartOffset, isu.EndOffset, isu.Message)
		}
	}()

	// 获取hotspot到本地库
	response, err = sonarapi.GetHotspotByProject(in.ProjectName)
	if err != nil {
		return &sonarqubepb.CreateProjectResp{
			Code:    500,
			Message: get_hotspot_error,
		}, nil
	}

	go func() {
		Isus := PaserHotspot(response)
		for _, isu := range Isus {
			mysql.CreateIssue(ctx, in.ProjectName, isu.Type, isu.File, isu.StartLine, isu.EndLine, isu.StartOffset, isu.EndOffset, isu.Message)
		}
	}()

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

func ParseIssue(input map[string]interface{}) (output []Isu) {
	output = make([]Isu, 0)
	for _, issue := range input["issues"].([]interface{}) {
		issueMap := issue.(map[string]interface{})
		var isu Isu
		isu.Type = issueMap["type"].(string)
		isu.File = issueMap["component"].(string)
		isu.StartLine = int(issueMap["textRange"].(map[string]interface{})["startLine"].(float64))
		isu.EndLine = int(issueMap["textRange"].(map[string]interface{})["endLine"].(float64))
		isu.StartOffset = int(issueMap["textRange"].(map[string]interface{})["startOffset"].(float64))
		isu.EndOffset = int(issueMap["textRange"].(map[string]interface{})["endOffset"].(float64))
		isu.Message = issueMap["message"].(string)
		output = append(output, isu)
	}
	return
}

func PaserHotspot(input map[string]interface{}) (output []Isu) {
	output = make([]Isu, 0)
	for _, hotspot := range input["hotspots"].([]interface{}) {
		hotspotMap := hotspot.(map[string]interface{})
		var isu Isu
		isu.Type = "HOTSPOT"
		isu.File = hotspotMap["component"].(string)
		isu.StartLine = int(hotspotMap["textRange"].(map[string]interface{})["startLine"].(float64))
		isu.EndLine = int(hotspotMap["textRange"].(map[string]interface{})["endLine"].(float64))
		isu.StartOffset = int(hotspotMap["textRange"].(map[string]interface{})["startOffset"].(float64))
		isu.EndOffset = int(hotspotMap["textRange"].(map[string]interface{})["endOffset"].(float64))
		isu.Message = hotspotMap["message"].(string)
		output = append(output, isu)
	}
	return
}
