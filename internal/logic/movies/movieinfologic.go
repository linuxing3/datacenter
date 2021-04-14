package logic

import (
	"context"
	"fmt"
	"log"

	"datacenter/internal/svc"
	"datacenter/internal/types"
	// FIXED: 调用rpcclient
	"datacenter/movies/rpc/movieclient"
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

func (l *MovieInfoLogic) MovieInfo(req types.MovieReq) (*movieclient.MovieListResp, error) {
	// TODO: 调用rpc的movieclient，像rpc服务器发送请求
	id:= fmt.Sprintf("%s", req.Id)
	log.Printf("%v", id);
	ids := []string{id}
	return l.svcCtx.MoviesRpc.Movies(l.ctx, &movieclient.MovieListReq{
		Ids: ids,
	})
	// return &types.MovieReply{
	// 	Title: "movies",
	// }, nil
}
