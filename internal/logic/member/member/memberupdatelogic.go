package logic

import (
	"context"
	"go-zero-admin/rpc/ums/umsclient"

	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberUpdateLogic {
	return MemberUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberUpdateLogic) MemberUpdate(req types.UpdateMemberReq) (*types.UpdateMemberResp, error) {
	_, err := l.svcCtx.UmsRpc.MemberUpdate(l.ctx, &umsclient.MemberUpdateReq{
		Id:                    req.Id,
		MemberLevelId:         req.MemberLevelId,
		Username:              req.Username,
		Password:              req.Password,
		Nickname:              req.Nickname,
		Phone:                 req.Phone,
		Status:                req.Status,
		CreateTime:            req.CreateTime,
		Icon:                  req.Icon,
		Gender:                req.Gender,
		Birthday:              req.Birthday,
		City:                  req.City,
		Job:                   req.Job,
		PersonalizedSignature: req.PersonalizedSignature,
		SourceType:            req.SourceType,
		Integration:           req.Integration,
		Growth:                req.Growth,
		LuckeyCount:           req.LuckeyCount,
		HistoryIntegration:    req.HistoryIntegration,
	})

	if err != nil {
		return nil, err
	}

	return &types.UpdateMemberResp{
		Code:    "000000",
		Message: "",
	}, nil
}
