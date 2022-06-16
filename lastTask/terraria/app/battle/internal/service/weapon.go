package service

import (
	"context"

	pb "terraria/api/battle/v1"
	"terraria/app/battle/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type WeaponService struct {
	pb.UnimplementedWeaponServer
	log *log.Helper
	uc  *biz.WeaponUsecase
}

func NewWeaponService(uc *biz.WeaponUsecase, logger log.Logger) *WeaponService {
	return &WeaponService{uc: uc, log: log.NewHelper(logger)}
}

func (s *WeaponService) CreateWeapon(ctx context.Context, req *pb.CreateWeaponRequest) (*pb.CreateWeaponReply, error) {
	return &pb.CreateWeaponReply{}, nil
}
func (s *WeaponService) UpdateWeapon(ctx context.Context, req *pb.UpdateWeaponRequest) (*pb.UpdateWeaponReply, error) {
	return &pb.UpdateWeaponReply{}, nil
}
func (s *WeaponService) DeleteWeapon(ctx context.Context, req *pb.DeleteWeaponRequest) (*pb.DeleteWeaponReply, error) {
	return &pb.DeleteWeaponReply{}, nil
}
func (s *WeaponService) GetWeapon(ctx context.Context, req *pb.GetWeaponRequest) (*pb.GetWeaponReply, error) {
	return &pb.GetWeaponReply{}, nil
}
func (s *WeaponService) ListWeapon(ctx context.Context, req *pb.ListWeaponRequest) (*pb.ListWeaponReply, error) {
	return &pb.ListWeaponReply{}, nil
}
