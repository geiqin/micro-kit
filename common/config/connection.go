package config

import "github.com/geiqin/xconfig/model"

var dbConnections map[string]*model.DatabaseInfo

func Load(connections map[string]*model.DatabaseInfo) {
	dbConnections = connections
}

//获取数据库连接配置
func GetConnectCfg(name string, storeFlag ...string) *model.DatabaseInfo {
	cfg := dbConnections[name]
	if &cfg != nil && storeFlag != nil {
		cfg.Database = storeFlag[0]
	}
	return cfg
}
