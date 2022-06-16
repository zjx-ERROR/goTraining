package data

import (
	"context"

	"terraria/app/battle/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type weaponRepo struct {
	data *Data
	log  *log.Helper
}

// NewWeaponRepo .
func NewWeaponRepo(data *Data, logger log.Logger) biz.WeaponRepo {
	return &weaponRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *weaponRepo) Save(ctx context.Context, g *biz.Weapon) (*biz.Weapon, error) {
	return g, nil
}

func (r *weaponRepo) Update(ctx context.Context, g *biz.Weapon) (*biz.Weapon, error) {
	return g, nil
}

func (r *weaponRepo) FindByID(context.Context, int64) (*biz.Weapon, error) {
	return nil, nil
}

func (r *weaponRepo) ListByName(context.Context, string) ([]*biz.Weapon, error) {
	return nil, nil
}

func (r *weaponRepo) ListAll(context.Context) ([]*biz.Weapon, error) {
	return nil, nil
}
