package logic

import (
	"context"
	"go-zero-admin/rpc/ums/umsclient"

	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberAddLogic {
	return MemberAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberAddLogic) MemberAdd(req types.AddMemberReq) (*types.AddMemberResp, error) {
	_, err := l.svcCtx.UmsRpc.MemberAdd(l.ctx, &umsclient.MemberAddReq{
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

	return &types.AddMemberResp{
		Code:    "000000",
		Message: "",
	}, nil
}
