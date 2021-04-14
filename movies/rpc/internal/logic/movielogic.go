package logic

import (
	"context"

	"datacenter/movies/rpc/internal/svc"
	"datacenter/movies/rpc/movie"
	"datacenter/shared"

	"github.com/tal-tech/go-zero/core/logx"
)

type MovieLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMovieLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MovieLogic {
	return &MovieLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MovieLogic) Movie(in *movie.MovieReq) (*movie.MovieResp, error) {
	// todo: add your logic here and delete this line
	movieInt := shared.StrToInt64(in.Id)
	movieInfo, _ := l.svcCtx.MoviesModel.FindOne(movieInt)
	return &movie.MovieResp{
		Title: movieInfo.Title,
		Description: movieInfo.Description,
		Url: movieInfo.Url,
	}, nil
	// return &movie.MovieResp{}, nil
}
