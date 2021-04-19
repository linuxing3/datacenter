package logic

import (
	"context"
	"errors"
	"go-zero-admin/rpc/sys/sysclient"
	"strings"

	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserLoginLogic {
	return UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//根据用户名和密码登录
func (l *UserLoginLogic) UserLogin(req types.LoginReq, ip string) (*types.LoginResp, error) {

	if len(strings.TrimSpace(req.UserName)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errors.New("用户名或密码不能为空")
	}

	resp, err := l.svcCtx.SysRpc.Login(l.ctx, &sysclient.LoginReq{
		UserName: req.UserName,
		Password: req.Password,
	})

	if err != nil {
		return nil, err
	}

	//保存登录日志
	_, _ = l.svcCtx.SysRpc.LoginLogAdd(l.ctx, &sysclient.LoginLogAddReq{
		UserName: resp.UserName,
		Status:   "login",
		Ip:       ip,
		CreateBy: resp.UserName,
	})

	return &types.LoginResp{
		Code:             "000000",
		Message:          "登录成功",
		Status:           resp.Status,
		CurrentAuthority: resp.CurrentAuthority,
		Id:               resp.Id,
		UserName:         resp.UserName,
		AccessToken:      resp.AccessToken,
		AccessExpire:     resp.AccessExpire,
		RefreshAfter:     resp.RefreshAfter,
	}, nil
}
