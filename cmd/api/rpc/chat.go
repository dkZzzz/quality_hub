package rpc

import (
	"context"

	"github.com/dkZzzz/quality_hub/proto/chatpb"
	"github.com/dkZzzz/quality_hub/srv/chatsrv"
)

var (
	chatClient chatpb.ChatClient
)

// 调用客户端，给服务端发送请求

// 发送单个issue
func SentSingleIssue(req *chatpb.SentSingleIssueReq) (*chatpb.SentSingleIssueResp, error) {
	chatClient = chatsrv.Client
	ctx := context.Background()
	rsp, err := chatClient.SentSingleIssue(ctx, req)
	if err != nil {
		return nil, err
	} else {
		return rsp, nil
	}
}

func SentProjectIssue(req *chatpb.SentProjectIssueReq) (*chatpb.SentProjectIssueResp, error) {
	chatClient = chatsrv.Client
	ctx := context.Background()
	rsp, err := chatClient.SentProjectIssue(ctx, req)
	if err != nil {
		return nil, err
	} else {
		return rsp, nil
	}
}
