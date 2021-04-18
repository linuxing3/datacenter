package logic

import (
	"context"
	"go-zero-admin/rpc/ums/umsclient"

	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberTagDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberTagDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberTagDeleteLogic {
	return MemberTagDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberTagDeleteLogic) MemberTagDelete(req types.DeleteMemberTagReq) (*types.DeleteMemberTagResp, error) {
	_, _ = l.svcCtx.UmsRpc.MemberTagDelete(l.ctx, &umsclient.MemberTagDeleteReq{
		Id: req.Id,
	})

	return &types.DeleteMemberTagResp{
		Code:    "000000",
		Message: "",
	}, nil
}
