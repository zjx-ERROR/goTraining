package service

import (
	"context"

	battle "terraria/api/battle/v1"
	pb "terraria/api/bff/v1"
	"terraria/app/bff/internal/biz"
	"terraria/app/bff/internal/client"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

type BffService struct {
	pb.UnimplementedBffServer
	log          *log.Helper
	uc           *biz.BffUsecase
	grpClientOps client.GrpcClientOptions
}

func NewBffService(uc *biz.BffUsecase, grpClientOps client.GrpcClientOptions, logger log.Logger) *BffService {
	return &BffService{uc: uc, log: log.NewHelper(logger), grpClientOps: grpClientOps}
}

func (s *BffService) GetMonster(ctx context.Context, req *pb.GetMonsterRequest) (*pb.GetMonsterReply, error) {
	conn, err := grpc.DialInsecure(ctx, s.grpClientOps...)
	if err != nil {
		return nil, errors.New(500, err.Error(), "")
	}
	defer conn.Close()
	client := battle.NewMonsterClient(conn)
	reply, err := client.GetMonster(ctx, &battle.GetMonsterRequest{Id: req.Id})
	if err != nil {
		return nil, errors.New(500, err.Error(), "")
	}
	return &pb.GetMonsterReply{Name: reply.Name}, nil
}
func (s *BffService) ListMonster(ctx context.Context, req *pb.ListMonsterRequest) (*pb.ListMonsterReply, error) {
	conn, err := grpc.DialInsecure(ctx, s.grpClientOps...)
	if err != nil {
		return nil, errors.New(500, err.Error(), "")
	}
	defer conn.Close()
	client := battle.NewMonsterClient(conn)
	resp, err := client.ListMonster(ctx, &battle.ListMonsterRequest{})
	if err != nil {
		return nil, errors.New(500, err.Error(), "")
	}
	var reply []*pb.ListMonsterReply_Data
	for _, v := range resp.Data {
		reply = append(reply, &pb.ListMonsterReply_Data{Id: v.Id, Name: v.Name})
	}
	return &pb.ListMonsterReply{Data: reply}, nil
}
func (s *BffService) GetWeapon(ctx context.Context, req *pb.GetWeaponRequest) (*pb.GetWeaponReply, error) {
	conn, err := grpc.DialInsecure(ctx, s.grpClientOps...)
	if err != nil {
		return nil, errors.New(500, err.Error(), "")
	}
	defer conn.Close()
	client := battle.NewWeaponClient(conn)
	reply, err := client.GetWeapon(ctx, &battle.GetWeaponRequest{Id: req.Id})
	if err != nil {
		return nil, errors.New(500, err.Error(), "")
	}
	return &pb.GetWeaponReply{Name: reply.Name}, nil
}
func (s *BffService) ListWeapon(ctx context.Context, req *pb.ListWeaponRequest) (*pb.ListWeaponReply, error) {
	conn, err := grpc.DialInsecure(ctx, s.grpClientOps...)
	if err != nil {
		return nil, errors.New(500, err.Error(), "")
	}
	defer conn.Close()
	client := battle.NewWeaponClient(conn)
	resp, err := client.ListWeapon(ctx, &battle.ListWeaponRequest{})
	if err != nil {
		return nil, errors.New(500, err.Error(), "")
	}
	var reply []*pb.ListWeaponReply_Data
	for _, v := range resp.Data {
		reply = append(reply, &pb.ListWeaponReply_Data{Id: v.Id, Name: v.Name})
	}
	return &pb.ListWeaponReply{Data: reply}, nil
}
