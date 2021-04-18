package logic

import (
	"context"
	"datacenter/internal/svc"
	"datacenter/internal/types"
	"go-zero-admin/rpc/oms/omsclient"

	"github.com/tal-tech/go-zero/core/logx"
)

type CompayAddressDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCompayAddressDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) CompayAddressDeleteLogic {
	return CompayAddressDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CompayAddressDeleteLogic) CompayAddressDelete(req types.DeleteCompayAddressReq) (*types.DeleteCompayAddressResp, error) {
	_, _ = l.svcCtx.OmsRpc.CompanyAddressDelete(l.ctx, &omsclient.CompanyAddressDeleteReq{
		Id: req.Id,
	})

	return &types.DeleteCompayAddressResp{
		Code:    "000000",
		Message: "",
	}, nil
}
