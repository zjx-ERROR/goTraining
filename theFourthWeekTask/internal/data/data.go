package data

import (
	"fwt/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet(NewData, NewDB, NewGoodsRepo)

type Data struct {
	db  *gorm.DB
	log *log.Helper
}

func NewDB(conf *conf.Data, logger log.Logger) *gorm.DB {
	log := log.NewHelper(log.With(logger, "module", "data/gorm"))
	db, err := gorm.Open(postgres.Open(conf.Database.Source), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	if err := db.AutoMigrate(&Goods{}); err != nil {
		log.Fatal(err)
	}
	return db
}

func NewData(db *gorm.DB, logger log.Logger) (*Data, func(), error) {
	d := &Data{
		db:  db,
		log: log.NewHelper(log.With(logger, "module", "data")),
	}
	cleanup := func() {}
	return d, cleanup, nil
}
