package logic

import (
	"context"
	"go-zero-admin/rpc/ums/umsclient"

	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberLoginLogDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberLoginLogDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberLoginLogDeleteLogic {
	return MemberLoginLogDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberLoginLogDeleteLogic) MemberLoginLogDelete(req types.DeleteMemberLoginLogReq) (*types.DeleteMemberLoginLogResp, error) {
	_, _ = l.svcCtx.UmsRpc.MemberLoginLogDelete(l.ctx, &umsclient.MemberLoginLogDeleteReq{
		Id: req.Id,
	})

	return &types.DeleteMemberLoginLogResp{
		Code:    "000000",
		Message: "",
	}, nil
}
