package biz

import (
	"context"

	v1 "terraria/api/bff/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	ErrNotFound = errors.NotFound(v1.ErrorReason_NOT_FOUND.String(), "not found")
)

type Bff struct {
	Id   uint64
	Name string
}

type BffRepo interface {
	Save(context.Context, *Bff) (*Bff, error)
	Update(context.Context, *Bff) (*Bff, error)
	FindByID(context.Context, int64) (*Bff, error)
	ListByName(context.Context, string) ([]*Bff, error)
	ListAll(context.Context) ([]*Bff, error)
}

type BffUsecase struct {
	repo BffRepo
	log  *log.Helper
}

func NewBffUsecase(repo BffRepo, logger log.Logger) *BffUsecase {
	return &BffUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *BffUsecase) CreateBff(ctx context.Context, g *Bff) (*Bff, error) {
	uc.log.WithContext(ctx).Infof("CreateBff: %v", g.Name)
	return uc.repo.Save(ctx, g)
}
