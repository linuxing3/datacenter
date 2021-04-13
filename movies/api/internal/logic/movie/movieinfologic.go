package logic

import (
	"context"

	"datacenter/movies/api/internal/svc"
	"datacenter/movies/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MovieInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMovieInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) MovieInfoLogic {
	return MovieInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MovieInfoLogic) MovieInfo(req types.MovieReq) (*types.MovieReply, error) {
	// todo: add your logic here and delete this line

	return &types.MovieReply{}, nil
}
