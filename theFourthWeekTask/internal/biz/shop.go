package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Goods struct {
	Id    uint
	Name  string
	Price float64
}

type GoodsRepo interface {
	GetGoods(ctx context.Context, id uint64) (*Goods, error)
	CreateGoods(ctx context.Context, g *Goods) (uint64, error)
}

type GoodsUsecase struct {
	repo GoodsRepo
	log  *log.Helper
}

func NewGoodsUsecase(repo GoodsRepo, logger log.Logger) *GoodsUsecase {
	return &GoodsUsecase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "biz/Goods")),
	}
}

func (uc *GoodsUsecase) GetGoods(ctx context.Context, id uint64) (*Goods, error) {
	return uc.repo.GetGoods(ctx, id)
}

func (uc *GoodsUsecase) CreateGoods(ctx context.Context, g *Goods) (uint64, error) {
	return uc.repo.CreateGoods(ctx, g)
}
