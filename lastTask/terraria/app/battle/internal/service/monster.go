package service

import (
	"context"

	pb "terraria/api/battle/v1"
	"terraria/app/battle/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type MonsterService struct {
	pb.UnimplementedMonsterServer
	log *log.Helper
	uc  *biz.MonsterUsecase
}

func NewMonsterService(uc *biz.MonsterUsecase, logger log.Logger) *MonsterService {
	return &MonsterService{uc: uc, log: log.NewHelper(logger)}
}

func (s *MonsterService) CreateMonster(ctx context.Context, req *pb.CreateMonsterRequest) (*pb.CreateMonsterReply, error) {
	return &pb.CreateMonsterReply{}, nil
}
func (s *MonsterService) UpdateMonster(ctx context.Context, req *pb.UpdateMonsterRequest) (*pb.UpdateMonsterReply, error) {
	return &pb.UpdateMonsterReply{}, nil
}
func (s *MonsterService) DeleteMonster(ctx context.Context, req *pb.DeleteMonsterRequest) (*pb.DeleteMonsterReply, error) {
	return &pb.DeleteMonsterReply{}, nil
}
func (s *MonsterService) GetMonster(ctx context.Context, req *pb.GetMonsterRequest) (*pb.GetMonsterReply, error) {
	return &pb.GetMonsterReply{}, nil
}
func (s *MonsterService) ListMonster(ctx context.Context, req *pb.ListMonsterRequest) (*pb.ListMonsterReply, error) {
	return &pb.ListMonsterReply{}, nil
}
