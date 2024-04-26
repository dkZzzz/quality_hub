package chat

import (
	"context"

	"github.com/dkZzzz/quality_hub/db/redis"
	"github.com/dkZzzz/quality_hub/proto/chatpb"
)

type ChatServerImpl struct {
	chatpb.UnimplementedChatServer
}

var (
	token_error = "token验证失败"
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

	return &chatpb.SentSingleIssueResp{
		Code:    200,
		Message: "",
		Data:    nil,
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
