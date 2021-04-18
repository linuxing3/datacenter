package logic

import (
	"context"
	"go-zero-admin/rpc/ums/umsclient"

	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberTagAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberTagAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberTagAddLogic {
	return MemberTagAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberTagAddLogic) MemberTagAdd(req types.AddMemberTagReq) (*types.AddMemberTagResp, error) {
	_, err := l.svcCtx.UmsRpc.MemberTagAdd(l.ctx, &umsclient.MemberTagAddReq{
		Name:              req.Name,
		FinishOrderCount:  req.FinishOrderCount,
		FinishOrderAmount: int64(req.FinishOrderAmount),
	})

	if err != nil {
		return nil, err
	}

	return &types.AddMemberTagResp{
		Code:    "000000",
		Message: "",
	}, nil
}
