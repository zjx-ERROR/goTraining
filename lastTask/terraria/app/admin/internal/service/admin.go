package service

import (
	"context"

	pb "terraria/api/admin/v1"
)

type AdminService struct {
	pb.UnimplementedAdminServer
}

func NewAdminService() *AdminService {
	return &AdminService{}
}

func (s *AdminService) Createmonster(ctx context.Context, req *pb.CreateMonsterRequest) (*pb.CreateMonsterReply, error) {
	return &pb.CreateMonsterReply{}, nil
}
func (s *AdminService) Updatemonster(ctx context.Context, req *pb.UpdateMonsterRequest) (*pb.UpdateMonsterReply, error) {
	return &pb.UpdateMonsterReply{}, nil
}
func (s *AdminService) Deletemonster(ctx context.Context, req *pb.DeleteMonsterRequest) (*pb.DeleteMonsterReply, error) {
	return &pb.DeleteMonsterReply{}, nil
}
func (s *AdminService) Getmonster(ctx context.Context, req *pb.GetMonsterRequest) (*pb.GetMonsterReply, error) {
	return &pb.GetMonsterReply{}, nil
}
func (s *AdminService) Listmonster(ctx context.Context, req *pb.ListMonsterRequest) (*pb.ListMonsterReply, error) {
	return &pb.ListMonsterReply{}, nil
}
func (s *AdminService) Createweapon(ctx context.Context, req *pb.CreateWeaponRequest) (*pb.CreateWeaponReply, error) {
	return &pb.CreateWeaponReply{}, nil
}
func (s *AdminService) Updateweapon(ctx context.Context, req *pb.UpdateWeaponRequest) (*pb.UpdateWeaponReply, error) {
	return &pb.UpdateWeaponReply{}, nil
}
func (s *AdminService) Deleteweapon(ctx context.Context, req *pb.DeleteWeaponRequest) (*pb.DeleteWeaponReply, error) {
	return &pb.DeleteWeaponReply{}, nil
}
func (s *AdminService) Getweapon(ctx context.Context, req *pb.GetWeaponRequest) (*pb.GetWeaponReply, error) {
	return &pb.GetWeaponReply{}, nil
}
func (s *AdminService) Listweapon(ctx context.Context, req *pb.ListWeaponRequest) (*pb.ListWeaponReply, error) {
	return &pb.ListWeaponReply{}, nil
}
