package data

import (
	"context"

	"terraria/app/admin/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type adminRepo struct {
	data *Data
	log  *log.Helper
}

// NewAdminRepo .
func NewAdminRepo(data *Data, logger log.Logger) biz.AdminRepo {
	return &adminRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *adminRepo) Save(ctx context.Context, g *biz.Admin) (*biz.Admin, error) {
	return g, nil
}

func (r *adminRepo) Update(ctx context.Context, g *biz.Admin) (*biz.Admin, error) {
	return g, nil
}

func (r *adminRepo) FindByID(context.Context, int64) (*biz.Admin, error) {
	return nil, nil
}

func (r *adminRepo) ListByName(context.Context, string) ([]*biz.Admin, error) {
	return nil, nil
}

func (r *adminRepo) ListAll(context.Context) ([]*biz.Admin, error) {
	return nil, nil
}
