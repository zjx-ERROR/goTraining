package biz

import (
	"context"

	v1 "terraria/api/admin/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	ErrNotFound = errors.NotFound(v1.ErrorReason_NOT_FOUND.String(), "not found")
)

type Admin struct {
	Id   uint64
	Name string
}

type AdminRepo interface {
	Save(context.Context, *Admin) (*Admin, error)
	Update(context.Context, *Admin) (*Admin, error)
	FindByID(context.Context, int64) (*Admin, error)
	ListByName(context.Context, string) ([]*Admin, error)
	ListAll(context.Context) ([]*Admin, error)
}

type AdminUsecase struct {
	repo AdminRepo
	log  *log.Helper
}

func NewAdminUsecase(repo AdminRepo, logger log.Logger) *AdminUsecase {
	return &AdminUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *AdminUsecase) CreateAdmin(ctx context.Context, g *Admin) (*Admin, error) {
	uc.log.WithContext(ctx).Infof("CreateAdmin: %v", g.Name)
	return uc.repo.Save(ctx, g)
}
