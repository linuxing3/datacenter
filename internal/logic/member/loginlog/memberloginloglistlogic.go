package logic

import (
	"context"
	"fmt"
	"go-zero-admin/rpc/ums/umsclient"

	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberLoginLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberLoginLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberLoginLogListLogic {
	return MemberLoginLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberLoginLogListLogic) MemberLoginLogList(req types.ListMemberLoginLogReq) (*types.ListMemberLoginLogResp, error) {
	resp, err := l.svcCtx.UmsRpc.MemberLoginLogList(l.ctx, &umsclient.MemberLoginLogListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		return nil, err
	}

	for _, data := range resp.List {
		fmt.Println(data)
	}
	var list []*types.ListtMemberLoginLogData

	for _, item := range resp.List {
		list = append(list, &types.ListtMemberLoginLogData{
			Id:         item.Id,
			MemberId:   item.MemberId,
			CreateTime: item.CreateTime,
			Ip:         item.Ip,
			City:       item.City,
			LoginType:  item.LoginType,
			Province:   item.Province,
		})
	}

	return &types.ListMemberLoginLogResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "",
	}, nil
}
