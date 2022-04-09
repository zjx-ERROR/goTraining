package service

import (
	"context"
	pb "fwt/api/shop/v1"
	"fwt/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type GoodsService struct {
	pb.UnimplementedShopServer
	uc  *biz.GoodsUsecase
	log *log.Helper
}

func NewGoodsService(uc *biz.GoodsUsecase, logger log.Logger) *GoodsService {
	return &GoodsService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service/goods")),
	}
}

func (s *GoodsService) GetGoods(ctx context.Context, req *pb.GetGoodsReq) (*pb.GetGoodsReply, error) {
	goods, err := s.uc.GetGoods(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetGoodsReply{Name: goods.Name, Price: goods.Price}, nil
}

func (s *GoodsService) CreateGoods(ctx context.Context, req *pb.CreateGoodsReq) (*pb.CreateGoodsReply, error) {
	tmp := biz.Goods{
		Name:  req.Name,
		Price: req.Price,
	}
	id, err := s.uc.CreateGoods(ctx, &tmp)
	if err != nil {
		return nil, err
	}
	return &pb.CreateGoodsReply{Id: id}, nil
}
