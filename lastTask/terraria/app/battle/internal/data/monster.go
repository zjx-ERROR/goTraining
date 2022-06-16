package data

import (
	"context"

	"terraria/app/battle/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type monsterRepo struct {
	data *Data
	log  *log.Helper
}

// NewMonsterRepo .
func NewMonsterRepo(data *Data, logger log.Logger) biz.MonsterRepo {
	return &monsterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *monsterRepo) Save(ctx context.Context, g *biz.Monster) (*biz.Monster, error) {
	return g, nil
}

func (r *monsterRepo) Update(ctx context.Context, g *biz.Monster) (*biz.Monster, error) {
	return g, nil
}

func (r *monsterRepo) FindByID(context.Context, int64) (*biz.Monster, error) {
	return nil, nil
}

func (r *monsterRepo) ListByName(context.Context, string) ([]*biz.Monster, error) {
	return nil, nil
}

func (r *monsterRepo) ListAll(context.Context) ([]*biz.Monster, error) {
	return nil, nil
}
