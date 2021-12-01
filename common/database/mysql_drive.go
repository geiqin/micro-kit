package database

import (
	"database/sql"
	"github.com/geiqin/micro-kit/common/config"
	"github.com/geiqin/micro-kit/common/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

// Mysql mysql配置结构体
type Mysql struct {
}

func (e *Mysql) RemoveSqlDB(cfg *Configure) {
	if cfg.IsPools {
		global.Cfg.SetDbs(cfg.Database, nil)
	} else {
		global.Cfg.SetDb(nil)
	}
}

func (e *Mysql) OpenSqlDB(cfg *Configure) (*sql.DB, error) {
	db, err := sql.Open(cfg.Driver, cfg.Source)
	if err != nil {
		log.Fatal(cfg.Driver+" connect error :", err)
		return nil, err
	}
	if err := db.Ping(); err != nil {
		defer db.Close()
		log.Fatal(cfg.Driver+"OpenSqlDB connect error :", err)
		return nil, err
	}
	if cfg.IsPools {
		global.Cfg.SetDbs(cfg.Database, &config.DBConfig{
			Driver: cfg.Driver,
			DB:     db,
		})
	} else {
		global.Cfg.SetDb(&config.DBConfig{
			Driver: cfg.Driver,
			DB:     db,
		})
	}
	return db, err
}
func (e *Mysql) GetSqlDB(cfg *Configure) (*sql.DB, error) {
	var db *sql.DB
	var err error

	//	key:= cfg.Database
	if cfg.IsPools {
		dbCfg := global.Cfg.GetDbByKey(cfg.Database)
		if dbCfg == nil {
			db, err = e.OpenSqlDB(cfg)
		} else {
			db = dbCfg.DB
		}
	} else {
		dbCfg := global.Cfg.GetDb()
		if dbCfg == nil {
			db, err = e.OpenSqlDB(cfg)
		} else {
			db = dbCfg.DB
		}
	}
	/*
		if err := db.Ping(); err != nil {
			defer db.Close()
			return nil, err
		}
	*/
	return db, err
}

// Setup 配置步骤
func (e *Mysql) Setup(cfg *Configure) (*gorm.DB, error) {
	//db, err := sql.Open(cfg.Driver, cfg.Source)
	db, err := e.GetSqlDB(cfg)
	if err != nil {
		defer db.Close()
		return nil, err
	}

	gormBb, err := e.Open(db, &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		CreateBatchSize:                          1000,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   cfg.TablePrefix,
			SingularTable: false,
		},
	})
	if err != nil {
		e.RemoveSqlDB(cfg)
		log.Fatal(" connect error :", err)
		return nil, err
	}
	if gormBb.Error != nil {
		e.RemoveSqlDB(cfg)
		log.Println(" connect gormBb fail:", gormBb.Error)
		return nil, err
	}

	return gormBb, nil
}

// Open 打开数据库连接
func (e *Mysql) Open(db *sql.DB, cfg *gorm.Config) (*gorm.DB, error) {
	return gorm.Open(mysql.New(mysql.Config{Conn: db}), cfg)
}
