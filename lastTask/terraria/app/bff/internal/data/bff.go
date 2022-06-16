package data

import (
	"context"

	"terraria/app/bff/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type bffRepo struct {
	data *Data
	log  *log.Helper
}

// NewBffRepo .
func NewBffRepo(data *Data, logger log.Logger) biz.BffRepo {
	return &bffRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *bffRepo) Save(ctx context.Context, g *biz.Bff) (*biz.Bff, error) {
	return g, nil
}

func (r *bffRepo) Update(ctx context.Context, g *biz.Bff) (*biz.Bff, error) {
	return g, nil
}

func (r *bffRepo) FindByID(context.Context, int64) (*biz.Bff, error) {
	return nil, nil
}

func (r *bffRepo) ListByName(context.Context, string) ([]*biz.Bff, error) {
	return nil, nil
}

func (r *bffRepo) ListAll(context.Context) ([]*biz.Bff, error) {
	return nil, nil
}
