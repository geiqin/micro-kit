package database

import (
	"errors"
	"gorm.io/gorm"
)

// Setup 配置数据库
func Setup(cfg *Configure) (*gorm.DB, error) {
	if cfg.Driver == "mysql" {
		var dba = new(Mysql)
		return dba.Setup(cfg)
	}
	return nil, errors.New("not set db configure")
}
