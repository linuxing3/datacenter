package logic

import (
	"context"
	"go-zero-admin/rpc/sms/smsclient"

	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type HomeRecommendSubjectUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeRecommendSubjectUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeRecommendSubjectUpdateLogic {
	return HomeRecommendSubjectUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeRecommendSubjectUpdateLogic) HomeRecommendSubjectUpdate(req types.UpdateHomeRecommendSubjectReq) (*types.UpdateHomeRecommendSubjectResp, error) {
	_, err := l.svcCtx.SmsRpc.HomeRecommendSubjectUpdate(l.ctx, &smsclient.HomeRecommendSubjectUpdateReq{
		Id:              req.Id,
		SubjectId:       req.SubjectId,
		SubjectName:     req.SubjectName,
		RecommendStatus: req.RecommendStatus,
		Sort:            req.Sort,
	})

	if err != nil {
		return nil, err
	}

	return &types.UpdateHomeRecommendSubjectResp{
		Code:    "000000",
		Message: "",
	}, nil
}
