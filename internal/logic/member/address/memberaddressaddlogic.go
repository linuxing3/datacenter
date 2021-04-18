package logic

import (
	"context"
	"datacenter/internal/svc"
	"datacenter/internal/types"
	"go-zero-admin/rpc/ums/umsclient"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberAddressAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberAddressAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberAddressAddLogic {
	return MemberAddressAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberAddressAddLogic) MemberAddressAdd(req types.AddMemberAddressReq) (*types.AddMemberAddressResp, error) {
	_, err := l.svcCtx.UmsRpc.MemberReceiveAddressAdd(l.ctx, &umsclient.MemberReceiveAddressAddReq{
		MemberId:      req.MemberId,
		Name:          req.Name,
		PhoneNumber:   req.PhoneNumber,
		DefaultStatus: req.DefaultStatus,
		PostCode:      req.PostCode,
		Province:      req.Province,
		City:          req.City,
		Region:        req.Region,
		DetailAddress: req.DetailAddress,
	})

	if err != nil {
		return nil, err
	}

	return &types.AddMemberAddressResp{
		Code:    "000000",
		Message: "",
	}, nil
}
