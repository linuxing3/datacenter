package logic

import (
	"context"
	"go-zero-admin/rpc/ums/umsclient"

	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberTagUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberTagUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberTagUpdateLogic {
	return MemberTagUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberTagUpdateLogic) MemberTagUpdate(req types.UpdateMemberTagReq) (*types.UpdateMemberTagResp, error) {
	_, err := l.svcCtx.UmsRpc.MemberTagUpdate(l.ctx, &umsclient.MemberTagUpdateReq{
		Id:                req.Id,
		Name:              req.Name,
		FinishOrderCount:  req.FinishOrderCount,
		FinishOrderAmount: int64(req.FinishOrderAmount),
	})

	if err != nil {
		return nil, err
	}

	return &types.UpdateMemberTagResp{
		Code:    "000000",
		Message: "",
	}, nil
}
