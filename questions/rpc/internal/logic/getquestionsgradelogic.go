package logic

import (
	"context"

	"datacenter/questions/rpc/internal/svc"
	"datacenter/questions/rpc/questions"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetQuestionsGradeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetQuestionsGradeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetQuestionsGradeLogic {
	return &GetQuestionsGradeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// θ·εεΎε
func (l *GetQuestionsGradeLogic) GetQuestionsGrade(in *questions.GradeReq) (*questions.QuestionsAnswerResp, error) {
	// todo: add your logic here and delete this line
	answers, err := l.svcCtx.QuestionsAnswersModel.FindOne(in)
	if err != nil {
		return nil, err
	}
	return &questions.QuestionsAnswerResp{
		AnswerId: answers.Id,
		Score:    answers.Score,
	}, nil
}
