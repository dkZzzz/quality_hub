package rpc

import (
	"context"

	"github.com/dkZzzz/quality_hub/proto/noticepb"
	"github.com/dkZzzz/quality_hub/srv/noticesrv"
)

var (
	noticeClient noticepb.NoticeClient
)

// 调用客户端，给服务端发送请求

// 获取单个建议
func GetSingleAdvice(req *noticepb.GetSingleAdviceReq) (*noticepb.GetSingleAdviceResp, error) {
	noticeClient = noticesrv.Client
	ctx := context.Background()
	rsp, err := noticeClient.GetSingleAdvice(ctx, req)
	if err != nil {
		return nil, err
	} else {
		return rsp, nil
	}
}

// 获取项目建议
func GetProjectAdvice(req *noticepb.GetProjectAdviceReq) (*noticepb.GetProjectAdviceResp, error) {
	noticeClient = noticesrv.Client
	ctx := context.Background()
	rsp, err := noticeClient.GetProjectAdvice(ctx, req)
	if err != nil {
		return nil, err
	} else {
		return rsp, nil
	}
}
