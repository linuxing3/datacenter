package logic

import (
	"context"
	"encoding/json"
	"go-zero-admin/rpc/sys/sysclient"
	"strings"

	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserInfoLogic {
	return UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (*types.UserInfoResp, error) {
	// 这里的key和生成jwt token时传入的key一致
	userId, _ := l.ctx.Value("userId").(json.Number).Int64()

	resp, err := l.svcCtx.SysRpc.UserInfo(l.ctx, &sysclient.InfoReq{
		UserId: userId,
	})

	if err != nil {
		return nil, err
	}

	var MenuTree []*types.ListMenuTree

	for _, item := range resp.MenuListTree {
		MenuTree = append(MenuTree, &types.ListMenuTree{
			Id:       item.Id,
			Path:     item.Path,
			Name:     item.Name,
			ParentId: item.ParentId,
			Icon:     item.Icon,
		})
	}

	var MenuTreeVue []*types.ListMenuTreeVue

	for _, item := range resp.MenuListTree {

		if len(strings.TrimSpace(item.VuePath)) != 0 {
			MenuTreeVue = append(MenuTreeVue, &types.ListMenuTreeVue{
				Id:           item.Id,
				ParentId:     item.ParentId,
				Title:        item.Name,
				Path:         item.VuePath,
				Name:         item.Name,
				Icon:         item.VueIcon,
				VueRedirect:  item.VueRedirect,
				VueComponent: item.VueComponent,
				Meta: types.MenuTreeMeta{
					Title: item.Name,
					Icon:  item.VueIcon,
				},
			})
		}

	}

	return &types.UserInfoResp{
		Code:        "000000",
		Message:     "获取个人信息成功",
		Avatar:      resp.Avatar,
		Name:        resp.Name,
		MenuTree:    MenuTree,
		MenuTreeVue: MenuTreeVue,
	}, nil
}
