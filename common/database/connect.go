package database

import (
	"context"
	"fmt"
	"github.com/geiqin/micro-kit/app"
	"github.com/geiqin/micro-kit/auth"
	"github.com/geiqin/micro-kit/common/config"

	//dbs "github.com/geiqin/micro-kit/database"
	"github.com/geiqin/xconfig/model"
	"gorm.io/gorm"
	"log"
)

func ConnectOther(ctx context.Context, cfg *model.DatabaseInfo) *gorm.DB {
	conf := GetConfigure(ctx, cfg)
	db, err := Setup(conf)
	if err != nil {
		log.Fatal("Connect db error:", err)
	}
	return db
}

func Connect(ctx context.Context, customAppFlag ...string) *gorm.DB {
	var flag string
	if customAppFlag != nil {
		flag = customAppFlag[0]
	} else {
		flag = app.Flag()
	}
	cfg := config.GetConnectCfg(flag)
	conf := GetConfigure(ctx, cfg)
	db, err := Setup(conf)
	if err != nil {
		log.Fatal("Connect db error:", err)
	}
	return db
}

func GetConfigure(ctx context.Context, cfg *model.DatabaseInfo) *Configure {
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
