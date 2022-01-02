package gateway

import "github.com/geiqin/gotools/model"

//授权频道
type GrantChannel struct {
	Id     int           `json:"id"`                   //id
	Name   string        `json:"name"  gorm:"size:50"` //频道名称
	Routes []*GrantRoute `json:"routes"`
	model.Timestamps
}

//授权路由
type GrantRoute struct {
	Id                 int    `json:"id"`                                    //id
	GrantChannelId     int    `json:"grant_channel_id"`                      //id
	Name               string `json:"name" gorm:"size:50"`                   //名称
	Type               string `json:"type" gorm:"type:enum('API','WEB')"`    //类型
	Path               string `json:"path" gorm:"size:100"`                  //路径
	Method             string `json:"method" gorm:"size:10"`                 //请求方式 ALL GET POST DELETE
	HasStore           bool   `json:"has_store" gorm:"default:0"`            //是否店铺API
	AllowManager       bool   `json:"allow_manager" gorm:"default:0"`        //允许平台管理员
	AllowStoreUser     bool   `json:"allow_store_user" gorm:"default:0"`     //允许店铺用户
	AllowStoreSeller   bool   `json:"allow_store_seller" gorm:"default:0"`   //允许店铺卖家
	AllowStoreCustomer bool   `json:"allow_store_customer" gorm:"default:0"` //允许店铺客户
	Unlimited          bool   `json:"unlimited" gorm:"default:0"`            //无限制（任何人都可以访问）
	model.Timestamps
}
