package logic

import (
	"context"
	"go-zero-admin/rpc/ums/umsclient"

	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberLevelDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberLevelDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberLevelDeleteLogic {
	return MemberLevelDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberLevelDeleteLogic) MemberLevelDelete(req types.DeleteMemberLevelReq) (*types.DeleteMemberLevelResp, error) {
	_, _ = l.svcCtx.UmsRpc.MemberLevelDelete(l.ctx, &umsclient.MemberLevelDeleteReq{
		Id: req.Id,
	})

	return &types.DeleteMemberLevelResp{
		Code:    "000000",
		Message: "",
	}, nil
}
