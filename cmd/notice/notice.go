package notice

import (
	"context"

	"github.com/dkZzzz/quality_hub/db/redis"
	"github.com/dkZzzz/quality_hub/proto/noticepb"
)

type NoticeServerImpl struct {
	noticepb.UnimplementedNoticeServer
}

var (
	token_error = "token验证失败"
)

// 具体业务逻辑实现
func (s *NoticeServerImpl) GetSingleAdvice(ctx context.Context, in *noticepb.GetSingleAdviceReq) (*noticepb.GetSingleAdviceResp, error) {
	if !redis.JWTMatch(ctx, in.Username, in.Token) {
		return &noticepb.GetSingleAdviceResp{
			Code:    401,
			Message: token_error,
			Data:    nil,
		}, nil
	}

	return &noticepb.GetSingleAdviceResp{
		Code:    200,
		Message: "",
		Data:    nil,
	}, nil
}

func (s *NoticeServerImpl) GetProjectAdvice(ctx context.Context, in *noticepb.GetProjectAdviceReq) (*noticepb.GetProjectAdviceResp, error) {
	if !redis.JWTMatch(ctx, in.Username, in.Token) {
		return &noticepb.GetProjectAdviceResp{
			Code:    401,
			Message: token_error,
			Data:    nil,
		}, nil
	}

	return &noticepb.GetProjectAdviceResp{
		Code:    200,
		Message: "",
		Data:    nil,
	}, nil
}
