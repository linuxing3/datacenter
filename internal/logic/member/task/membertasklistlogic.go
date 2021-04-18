package logic

import (
	"context"
	"fmt"
	"go-zero-admin/rpc/ums/umsclient"

	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberTaskListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberTaskListLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberTaskListLogic {
	return MemberTaskListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberTaskListLogic) MemberTaskList(req types.ListMemberTaskReq) (*types.ListMemberTaskResp, error) {
	resp, err := l.svcCtx.UmsRpc.MemberTaskList(l.ctx, &umsclient.MemberTaskListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		return nil, err
	}

	for _, data := range resp.List {
		fmt.Println(data)
	}
	var list []*types.ListtMemberTaskData

	for _, item := range resp.List {
		list = append(list, &types.ListtMemberTaskData{
			Id:           item.Id,
			Name:         item.Name,
			Growth:       item.Growth,
			Intergration: item.Intergration,
			Type:         item.Type,
		})
	}

	return &types.ListMemberTaskResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "",
	}, nil
}
