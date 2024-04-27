package notice

import (
	"context"

	"github.com/dkZzzz/quality_hub/db/mysql"
	"github.com/dkZzzz/quality_hub/db/redis"
	"github.com/dkZzzz/quality_hub/proto/noticepb"
)

type NoticeServerImpl struct {
	noticepb.UnimplementedNoticeServer
}

var (
	token_error      = "token验证失败"
	get_advice_error = "获取建议失败"

	get_advice_succ = "获取建议成功"
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

	advice, err := mysql.GetAdviceByID(ctx, int(in.AdviceId))
	if err != nil {
		return &noticepb.GetSingleAdviceResp{
			Code:    500,
			Message: get_advice_error,
		}, nil
	}

	data := &noticepb.Advice{
		Id:          int32(advice.ID),
		IssueId:     int32(advice.IssueID),
		ProjectName: advice.ProjectName,
		Advice:      advice.Advice,
	}

	return &noticepb.GetSingleAdviceResp{
		Code:    200,
		Message: get_advice_succ,
		Data:    data,
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

	advices, err := mysql.GetProjectAdvice(ctx, in.ProjectName)
	if err != nil {
		return &noticepb.GetProjectAdviceResp{
			Code:    500,
			Message: get_advice_error,
			Data:    nil,
		}, nil
	}

	data := []*noticepb.Advice{}
	for _, ad := range advices {
		advice := &noticepb.Advice{
			Id:          int32(ad.ID),
			IssueId:     int32(ad.IssueID),
			ProjectName: ad.ProjectName,
			Advice:      ad.Advice,
		}
		data = append(data, advice)
	}

	return &noticepb.GetProjectAdviceResp{
		Code:    200,
		Message: get_advice_succ,
		Data:    data,
	}, nil
}
