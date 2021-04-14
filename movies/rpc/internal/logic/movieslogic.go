package logic

import (
	"context"
	"log"

	"datacenter/movies/rpc/internal/svc"
	"datacenter/movies/rpc/movie"
	"datacenter/shared"

	"github.com/tal-tech/go-zero/core/logx"
)

type MoviesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMoviesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MoviesLogic {
	return &MoviesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MoviesLogic) Movies(in *movie.MovieListReq) (*movie.MovieListResp, error) {
	// FIXED 这里接受网关传来的数据，调用model的方法查找并返回结果
	// 检查 internal/logic/movies/movieinfologic.go
	movieInt := shared.StrToInt64(in.Ids[0])
	log.Printf("获得网关的请求id： %v， 开始查询。。。", movieInt)

	movieInfo, _ := l.svcCtx.MoviesModel.FindOne(movieInt)

	list := make([]*movie.MovieResp, 0)
	list = append(list, &movie.MovieResp{
		Title: movieInfo.Title,
		Description: movieInfo.Description,
		Url: movieInfo.Url,
	})
	return &movie.MovieListResp{
		Data: list,
	}, nil
	// return &movie.MovieListResp{}, nil
}
