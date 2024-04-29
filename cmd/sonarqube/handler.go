package sonarqube

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/dkZzzz/quality_hub/config"
	"github.com/dkZzzz/quality_hub/db/mysql"
	"github.com/dkZzzz/quality_hub/db/redis"
	"github.com/dkZzzz/quality_hub/pkg/clone"
	"github.com/dkZzzz/quality_hub/pkg/sonarapi"
	"github.com/dkZzzz/quality_hub/proto/sonarqubepb"
)

var wg sync.WaitGroup

var (
	token_error            = "token验证失败"
	create_project_error   = "创建项目失败"
	clone_project_error    = "克隆项目失败"
	create_token_error     = "创建项目token失败"
	scanning_error         = "扫描代码失败"
	get_issue_error        = "获取issue失败"
	get_hotspot_error      = "获取hotspot失败"
	get_duplication_error  = "获取重复度失败"
	get_coverage_error     = "获取测试覆盖率失败"
	get_project_error      = "获取项目失败"
	get_report_error       = "获取报告失败"
	get_project_list_error = "获取项目列表失败"
	get_report_list_error  = "获取报告列表失败"

	get_issue_succ        = "获取issue成功"
	create_project_succ   = "创建项目成功"
	get_project_succ      = "获取项目成功"
	get_report_succ       = "获取报告成功"
	get_project_list_succ = "获取项目列表成功"
	get_report_list_succ  = "获取报告列表成功"
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
			Data:    nil,
		}, nil
	}

	// 向sonarqube发送请求
	err := sonarapi.CreateProject(in.ProjectName, in.ProjectName)
	if err != nil {
		return &sonarqubepb.CreateProjectResp{
			Code:    500,
			Message: create_project_error,
			Data:    nil,
		}, nil
	}

	// 在本地库创建记录
	project, _ := mysql.CreateProject(ctx, in.Username, in.ProjectName, in.BranchName, in.Url, in.Token)

	// 处理url
	tmp := strings.Split(in.Url, "/")
	pName := tmp[len(tmp)-1]
	pName = pName[:len(pName)-4]
	path := config.Cfg.CodeStorePath + pName

	// 检测目录是否存在
	if _, err := os.Stat(path); err == nil {
		// 目录存在，删除目录
		if err := os.RemoveAll(path); err != nil {
			// 删除目录失败
			// 处理错误，可以输出日志或者返回错误给调用者
			return &sonarqubepb.CreateProjectResp{
				Code:    500,
				Message: clone_project_error,
				Data:    nil,
			}, nil
		}
	}

	// 克隆代码
	err = clone.Clone(in.Url, pName)
	if err != nil {
		return &sonarqubepb.CreateProjectResp{
			Code:    500,
			Message: clone_project_error,
			Data:    nil,
		}, nil
	}

	// 创建token
	token, err := sonarapi.GenerateProjectToken(in.ProjectName, in.ProjectName)
	if err != nil {
		return &sonarqubepb.CreateProjectResp{
			Code:    500,
			Message: create_token_error,
			Data:    nil,
		}, nil
	}

	// 扫描代码
	err = Scan(path, in.ProjectName, token)
	if err != nil {
		return &sonarqubepb.CreateProjectResp{
			Code:    500,
			Message: scanning_error,
			Data:    nil,
		}, nil
	}

	// 等待3s
	time.Sleep(3 * time.Second)

	// 获取issue到本地库
	response, err := sonarapi.GetIssueByProject(in.ProjectName)
	if err != nil {
		return &sonarqubepb.CreateProjectResp{
			Code:    500,
			Message: get_issue_error,
			Data:    nil,
		}, nil
	}
	issueCnt := int(response["paging"].(map[string]interface{})["total"].(float64))

	wg.Add(1)
	go func() {
		defer wg.Done()
		Isus := ParseIssue(response)
		for _, isu := range Isus {
			mysql.CreateIssue(ctx, in.ProjectName, isu.Type, isu.File, isu.StartLine, isu.EndLine, isu.StartOffset, isu.EndOffset, isu.Message)
		}
	}()
	wg.Wait()

	// 获取hotspot到本地库
	response, err = sonarapi.GetHotspotByProject(in.ProjectName)
	if err != nil {
		return &sonarqubepb.CreateProjectResp{
			Code:    500,
			Message: get_hotspot_error,
			Data:    nil,
		}, nil
	}

	hotspotCnt := int(response["paging"].(map[string]interface{})["total"].(float64))

	wg.Add(1)
	go func() {
		defer wg.Done()
		Isus := PaserHotspot(response)
		for _, isu := range Isus {
			mysql.CreateIssue(ctx, in.ProjectName, isu.Type, isu.File, isu.StartLine, isu.EndLine, isu.StartOffset, isu.EndOffset, isu.Message)
		}
	}()
	wg.Wait()

	// 获取重复度
	duplication, err := sonarapi.GetDuplicationsByProject(in.ProjectName)
	if err != nil {
		return &sonarqubepb.CreateProjectResp{
			Code:    500,
			Message: get_duplication_error,
		}, nil
	}

	// 获取测试覆盖率
	coverage, err := sonarapi.GetCoverageByProject(in.ProjectName)
	if err != nil {
		return &sonarqubepb.CreateProjectResp{
			Code:    500,
			Message: get_coverage_error,
		}, nil
	}

	// 创建report
	wg.Add(1)
	go func() {
		defer wg.Done()
		mysql.CreateReport(ctx, in.Username, in.ProjectName, issueCnt, hotspotCnt, duplication, coverage)
	}()
	wg.Wait()

	data := &sonarqubepb.Project{
		Id:          int32(project.ID),
		ProjectName: project.ProjectName,
		BranchName:  project.BranchName,
		Url:         project.Url,
	}

	return &sonarqubepb.CreateProjectResp{
		Code:    200,
		Message: create_project_succ,
		Data:    data,
	}, nil
}

func (s *SonarQubeServerImpl) GetProject(ctx context.Context, in *sonarqubepb.GetProjectReq) (*sonarqubepb.GetProjectResp, error) {
	if !redis.JWTMatch(ctx, in.Username, in.Token) {
		return &sonarqubepb.GetProjectResp{
			Code:    401,
			Message: token_error,
			Data:    nil,
		}, nil
	}

	project, err := mysql.GetProject(ctx, in.ProjectName)
	if err != nil {
		return &sonarqubepb.GetProjectResp{
			Code:    500,
			Message: get_project_error,
			Data:    nil,
		}, nil
	}

	data := &sonarqubepb.Project{
		Id:          int32(project.ID),
		ProjectName: project.ProjectName,
		BranchName:  project.BranchName,
		Url:         project.Url,
	}

	return &sonarqubepb.GetProjectResp{
		Code:    200,
		Message: get_project_succ,
		Data:    data,
	}, nil
}

func (s *SonarQubeServerImpl) GetProjectList(ctx context.Context, in *sonarqubepb.GetProjectListReq) (*sonarqubepb.GetProjectListResp, error) {
	if !redis.JWTMatch(ctx, in.Username, in.Token) {
		return &sonarqubepb.GetProjectListResp{
			Code:    401,
			Message: token_error,
			Data:    nil,
		}, nil
	}

	projects, err := mysql.GetProjectList(ctx, in.Username)
	if err != nil {
		return &sonarqubepb.GetProjectListResp{
			Code:    500,
			Message: get_project_list_error,
			Data:    nil,
		}, nil
	}
	var data []*sonarqubepb.Project
	for _, p := range projects {
		project := &sonarqubepb.Project{
			Id:          int32(p.ID),
			ProjectName: p.ProjectName,
			BranchName:  p.BranchName,
			Url:         p.Url,
		}
		data = append(data, project)
	}

	return &sonarqubepb.GetProjectListResp{
		Code:    200,
		Message: get_project_list_succ,
		Data:    data,
	}, nil
}

func (s *SonarQubeServerImpl) GetReport(ctx context.Context, in *sonarqubepb.GetReportReq) (*sonarqubepb.GetReportResp, error) {
	if !redis.JWTMatch(ctx, in.Username, in.Token) {
		return &sonarqubepb.GetReportResp{
			Code:    401,
			Message: token_error,
			Data:    nil,
		}, nil
	}

	report, err := mysql.GetReport(ctx, int(in.ReportId))
	if err != nil {
		return &sonarqubepb.GetReportResp{
			Code:    500,
			Message: get_report_error,
			Data:    nil,
		}, nil
	}

	data := &sonarqubepb.Report{
		Id:          int32(report.ID),
		ProjectName: report.ProjectName,
		IssueNum:    int32(report.IssueNum),
		HotspotNum:  int32(report.HotspotNum),
		Duplication: report.Duplication,
		Coverage:    report.Coverage,
	}

	return &sonarqubepb.GetReportResp{
		Code:    200,
		Message: get_report_succ,
		Data:    data,
	}, nil
}

func (s *SonarQubeServerImpl) GetReportList(ctx context.Context, in *sonarqubepb.GetReportListReq) (*sonarqubepb.GetReportListResp, error) {
	if !redis.JWTMatch(ctx, in.Username, in.Token) {
		return &sonarqubepb.GetReportListResp{
			Code:    401,
			Message: token_error,
			Data:    nil,
		}, nil
	}

	reports, err := mysql.GetReportList(ctx, in.Username)
	if err != nil {
		return &sonarqubepb.GetReportListResp{
			Code:    500,
			Message: get_report_list_error,
			Data:    nil,
		}, nil
	}

	var data []*sonarqubepb.Report
	for _, r := range reports {
		report := &sonarqubepb.Report{
			Id:          int32(r.ID),
			ProjectName: r.ProjectName,
			IssueNum:    int32(r.IssueNum),
			HotspotNum:  int32(r.HotspotNum),
			Duplication: r.Duplication,
			Coverage:    r.Coverage,
		}
		data = append(data, report)
	}

	return &sonarqubepb.GetReportListResp{
		Code:    200,
		Message: get_report_list_succ,
		Data:    data,
	}, nil
}

func (s *SonarQubeServerImpl) GetIssue(ctx context.Context, in *sonarqubepb.GetIssueReq) (*sonarqubepb.GetIssueResp, error) {
	if !redis.JWTMatch(ctx, in.Username, in.Token) {
		return &sonarqubepb.GetIssueResp{
			Code:    401,
			Message: token_error,
			Data:    nil,
		}, nil
	}

	issues, err := mysql.GetIssue(ctx, in.ProjectName)
	if err != nil {
		return &sonarqubepb.GetIssueResp{
			Code:    500,
			Message: get_issue_error,
			Data:    nil,
		}, nil
	}

	var data []*sonarqubepb.Issue
	for _, i := range issues {
		issue := &sonarqubepb.Issue{
			Type:        i.Type,
			File:        i.File,
			StartLine:   int32(i.StartLine),
			EndLine:     int32(i.EndLine),
			StartOffset: int32(i.StartOffset),
			EndOffset:   int32(i.EndOffset),
			Message:     i.Message,
		}
		data = append(data, issue)
	}

	return &sonarqubepb.GetIssueResp{
		Code:    200,
		Message: get_issue_succ,
		Data:    data,
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
		isu.File = strings.Split(issueMap["component"].(string), ":")[1]
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
		isu.File = strings.Split(hotspotMap["component"].(string), ":")[1]
		isu.StartLine = int(hotspotMap["textRange"].(map[string]interface{})["startLine"].(float64))
		isu.EndLine = int(hotspotMap["textRange"].(map[string]interface{})["endLine"].(float64))
		isu.StartOffset = int(hotspotMap["textRange"].(map[string]interface{})["startOffset"].(float64))
		isu.EndOffset = int(hotspotMap["textRange"].(map[string]interface{})["endOffset"].(float64))
		isu.Message = hotspotMap["message"].(string)
		output = append(output, isu)
	}
	return
}
