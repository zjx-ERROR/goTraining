package data

import (
	"context"
	"fwt/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type Goods struct {
	Id    uint `gorm:"primarykey"`
	Name  string
	Price float64
}

type goodsRepo struct {
	data *Data
	log  *log.Helper
}

func NewGoodsRepo(data *Data, logger log.Logger) biz.GoodsRepo {
	return &goodsRepo{data: data, log: log.NewHelper(log.With(logger, "module", "data/goods"))}
}

func (r *goodsRepo) GetGoods(ctx context.Context, id uint64) (*biz.Goods, error) {
	return nil, nil
}

func (r *goodsRepo) CreateGoods(ctx context.Context, g *biz.Goods) (uint64, error) {
	return 0, nil
}
