package chat

import (
	"context"
	"os"
	"strconv"
	"strings"

	"github.com/dkZzzz/quality_hub/config"
	"github.com/dkZzzz/quality_hub/db/mysql"
	"github.com/dkZzzz/quality_hub/db/redis"
	"github.com/dkZzzz/quality_hub/proto/chatpb"
)

type ChatServerImpl struct {
	chatpb.UnimplementedChatServer
}

var (
	token_error         = "token验证失败"
	get_issue_error     = "获取issue失败"
	issue_not_exist     = "issue不存在"
	get_codeline_error  = "获取代码行失败"
	chat_eror           = "ChatGPT交互失败"
	create_advice_error = "创建建议失败"

	sent_issue_success = "发送issue成功"
)

// 具体业务逻辑实现

// 发送单个issue
func (s *ChatServerImpl) SentSingleIssue(ctx context.Context, in *chatpb.SentSingleIssueReq) (*chatpb.SentSingleIssueResp, error) {
	if !redis.JWTMatch(ctx, in.Username, in.Token) {
		return &chatpb.SentSingleIssueResp{
			Code:    401,
			Message: token_error,
			Data:    nil,
		}, nil
	}

	issue, err := mysql.GetIssueByID(ctx, int(in.IssueId))
	if err != nil {
		return &chatpb.SentSingleIssueResp{
			Code:    500,
			Message: get_issue_error,
			Data:    nil,
		}, nil
	}

	if issue == nil {
		return &chatpb.SentSingleIssueResp{
			Code:    404,
			Message: issue_not_exist,
			Data:    nil,
		}, nil
	}

	code, err := find(issue)
	if err != nil {
		return &chatpb.SentSingleIssueResp{
			Code:    500,
			Message: get_codeline_error,
			Data:    nil,
		}, nil
	}

	advice, err := Chat(code, issue.Message)
	if err != nil {
		return &chatpb.SentSingleIssueResp{
			Code:    500,
			Message: chat_eror,
			Data:    nil,
		}, nil

	}

	adviceID, err := mysql.CraeteAdvice(ctx, int(in.IssueId), issue.ProjectName, advice)
	if err != nil {
		return &chatpb.SentSingleIssueResp{
			Code:    500,
			Message: create_advice_error,
			Data:    nil,
		}, nil
	}

	return &chatpb.SentSingleIssueResp{
		Code:    200,
		Message: sent_issue_success,
		Data: map[string]string{
			"advice_id": strconv.Itoa(adviceID),
			"advice":    advice,
		},
	}, nil
}

// 发送项目issue
func (s *ChatServerImpl) SentProjectIssue(ctx context.Context, in *chatpb.SentProjectIssueReq) (*chatpb.SentProjectIssueResp, error) {
	if !redis.JWTMatch(ctx, in.Username, in.Token) {
		return &chatpb.SentProjectIssueResp{
			Code:    401,
			Message: token_error,
			Data:    nil,
		}, nil
	}

	return &chatpb.SentProjectIssueResp{
		Code:    200,
		Message: "",
		Data:    nil,
	}, nil
}

// 通过issue的位置信息，定位到代码行
// 提取出issue范围的代码行
// 返回代码行
func find(issue *mysql.Issue) (string, error) {
	absolutepath := config.Cfg.CodeStorePath + issue.ProjectName + "/" + issue.File
	fileContent, err := os.ReadFile(absolutepath)
	if err != nil {
		return "", err
	}
	lines := strings.Split(string(fileContent), "\n")

	var codeLines []string
	for i := issue.StartLine - 1; i < issue.EndLine; i++ {
		line := lines[i]
		if len(line) >= issue.StartOffset && len(line) >= issue.EndOffset {
			codeLines = append(codeLines, line[issue.StartOffset:issue.EndOffset])
		}
	}

	// 将代码行连接为一个字符串
	code := strings.Join(codeLines, "\n")
	return code, nil
}
