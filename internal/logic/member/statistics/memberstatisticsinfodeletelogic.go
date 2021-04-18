package logic

import (
	"context"
	"go-zero-admin/rpc/ums/umsclient"

	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberStatisticsInfoDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberStatisticsInfoDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberStatisticsInfoDeleteLogic {
	return MemberStatisticsInfoDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberStatisticsInfoDeleteLogic) MemberStatisticsInfoDelete(req types.DeleteMemberStatisticsInfoReq) (*types.DeleteMemberStatisticsInfoResp, error) {
	_, _ = l.svcCtx.UmsRpc.MemberStatisticsInfoDelete(l.ctx, &umsclient.MemberStatisticsInfoDeleteReq{
		Id: req.Id,
	})

	return &types.DeleteMemberStatisticsInfoResp{
		Code:    "000000",
		Message: "",
	}, nil
}
