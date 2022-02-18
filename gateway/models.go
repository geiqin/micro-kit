package gateway

import "github.com/geiqin/gotools/model"

//授权频道
type GrantChannel struct {
	Id     int           `json:"id"`                                 //id
	Name   string        `json:"name"  gorm:"size:50;unique"`        //频道名称
	Title  string        `json:"title"  gorm:"size:50"`              //频道标题
	Type   string        `json:"type" gorm:"type:enum('API','WEB')"` //类型
	Routes []*GrantRoute `json:"routes"`
	model.Timestamps
}

//授权路由
type GrantRoute struct {
	Id             int    `json:"id"`                                 //id
	GrantChannelId int    `json:"grant_channel_id"`                   //channelId
	Name           string `json:"name" gorm:"size:50"`                //名称
	Type           string `json:"type" gorm:"type:enum('API','WEB')"` //类型
	Path           string `json:"path" gorm:"size:200"`               //路径
	Method         string `json:"method" gorm:"size:10"`              //请求方式 ALL GET POST DELETE
	HasStore       bool   `json:"has_store" gorm:"default:0"`         //是否店铺API
	OnlyMaster     bool   `json:"only_master" gorm:"default:0"`       //只允许主铺
	AllowManager   bool   `json:"allow_manager" gorm:"default:0"`     //允许主铺管理员
	AllowCustomer  bool   `json:"allow_customer" gorm:"default:0"`    //允许商铺客户
	AllowSeller    bool   `json:"allow_seller" gorm:"default:0"`      //允许商铺卖家
	Unlimited      bool   `json:"unlimited" gorm:"default:0"`         //无限制（任何人都可以访问）
	model.Timestamps
}
