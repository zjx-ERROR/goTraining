package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Weapon struct {
	Id   uint64
	Name string
}

type WeaponRepo interface {
	Save(context.Context, *Weapon) (*Weapon, error)
	Update(context.Context, *Weapon) (*Weapon, error)
	FindByID(context.Context, int64) (*Weapon, error)
	ListByName(context.Context, string) ([]*Weapon, error)
	ListAll(context.Context) ([]*Weapon, error)
}

type WeaponUsecase struct {
	repo WeaponRepo
	log  *log.Helper
}

func NewWeaponUsecase(repo WeaponRepo, logger log.Logger) *WeaponUsecase {
	return &WeaponUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *WeaponUsecase) CreateWeapon(ctx context.Context, g *Weapon) (*Weapon, error) {
	uc.log.WithContext(ctx).Infof("CreateWeapon: %v", g.Name)
	return uc.repo.Save(ctx, g)
}
