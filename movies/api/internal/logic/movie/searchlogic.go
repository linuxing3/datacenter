package logic

import (
	"context"
	"strconv"
	"log"

	"datacenter/movies/api/internal/svc"
	"datacenter/movies/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) SearchLogic {
	return SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req types.MovieReq) (*types.MovieReply, error) {
	movieInt, err := strconv.ParseInt(req.Id, 10, 64)
	if err != nil {
		return nil, err
	}
	log.Println(movieInt)
	// todo: add your logic here and delete this line
	movieInfo, err := l.svcCtx.MovieModel.FindOne(movieInt)
	return &types.MovieReply{
		Id: movieInfo.Id,
		Title: movieInfo.Title,
		Description: movieInfo.Description,
		Url: movieInfo.Url,
	}, nil
}
