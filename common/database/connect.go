package database

import (
	"context"
	"fmt"
	"github.com/geiqin/micro-kit/app"
	"github.com/geiqin/micro-kit/auth"
	dbs "github.com/geiqin/micro-kit/database"
	"gorm.io/gorm"
	"log"
)

func Connect(ctx context.Context) *gorm.DB {
	conf := GetConfigure(ctx)
	db, err := Setup(conf)
	if err != nil {
		log.Fatal("Connect db error:", err)
	}
	return db
}

// GetConnect 获取数据库连接
func GetConfigure(ctx context.Context) *Configure {
	cfg := dbs.GetConnectCfg(app.Flag())
	appCfg := app.GetConfig()
	dbName := cfg.Database
	if app.Private() {
		storeId := auth.GetStoreId(ctx)
		if storeId > 0 {
			dbName = auth.GetStoreFlag(storeId, appCfg.DbPrefix+"store_")
		}
	}
	serverAddr := cfg.Host + ":" + cfg.Port
	source := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=1000ms", cfg.Username, cfg.Password, serverAddr, dbName)

	return &Configure{
		Driver:      "mysql",
		Source:      source,
		Database:    dbName,
		TablePrefix: cfg.Prefix,
		IsPools:     true,
	}
}
