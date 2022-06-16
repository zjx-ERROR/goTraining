package biz

import (
	"context"

	v1 "terraria/api/battle/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	ErrNotFound = errors.NotFound(v1.ErrorReason_NOT_FOUND.String(), "not found")
)

type Monster struct {
	Id   uint64
	Name string
}

type MonsterRepo interface {
	Save(context.Context, *Monster) (*Monster, error)
	Update(context.Context, *Monster) (*Monster, error)
	FindByID(context.Context, int64) (*Monster, error)
	ListByName(context.Context, string) ([]*Monster, error)
	ListAll(context.Context) ([]*Monster, error)
}

type MonsterUsecase struct {
	repo MonsterRepo
	log  *log.Helper
}

func NewMonsterUsecase(repo MonsterRepo, logger log.Logger) *MonsterUsecase {
	return &MonsterUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *MonsterUsecase) CreateMonster(ctx context.Context, g *Monster) (*Monster, error) {
	uc.log.WithContext(ctx).Infof("CreateMonster: %v", g.Name)
	return uc.repo.Save(ctx, g)
}
