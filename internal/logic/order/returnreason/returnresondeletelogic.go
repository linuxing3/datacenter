package logic

import (
	"context"
	"datacenter/internal/svc"
	"datacenter/internal/types"
	"go-zero-admin/rpc/oms/omsclient"

	"github.com/tal-tech/go-zero/core/logx"
)

type ReturnResonDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReturnResonDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) ReturnResonDeleteLogic {
	return ReturnResonDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReturnResonDeleteLogic) ReturnResonDelete(req types.DeleteReturnResonReq) (*types.DeleteReturnResonResp, error) {
	_, _ = l.svcCtx.OmsRpc.OrderReturnReasonDelete(l.ctx, &omsclient.OrderReturnReasonDeleteReq{
		Id: req.Id,
	})

	return &types.DeleteReturnResonResp{
		Code:    "000000",
		Message: "",
	}, nil
}
