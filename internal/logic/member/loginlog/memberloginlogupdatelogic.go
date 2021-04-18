package logic

import (
	"context"
	"go-zero-admin/rpc/ums/umsclient"

	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberLoginLogUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberLoginLogUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberLoginLogUpdateLogic {
	return MemberLoginLogUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberLoginLogUpdateLogic) MemberLoginLogUpdate(req types.UpdateMemberLoginLogReq) (*types.UpdateMemberLoginLogResp, error) {
	_, err := l.svcCtx.UmsRpc.MemberLoginLogUpdate(l.ctx, &umsclient.MemberLoginLogUpdateReq{
		Id:         req.Id,
		MemberId:   req.MemberId,
		CreateTime: req.CreateTime,
		Ip:         req.Ip,
		City:       req.City,
		LoginType:  req.LoginType,
		Province:   req.Province,
	})

	if err != nil {
		return nil, err
	}

	return &types.UpdateMemberLoginLogResp{
		Code:    "000000",
		Message: "",
	}, nil
}
