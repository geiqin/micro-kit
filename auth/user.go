package auth

import (
	"context"
	"fmt"
	"github.com/geiqin/gotools/helper"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

//当前授权用户
type User struct {
	Mode        string `json:"mode"`                    //授权模式
	SessionKey  string `json:"session_key,omitempty"`   //会话Key
	UserId      int64  `json:"user_id,omitempty"`       //用户ID
	PlatformId  int64  `json:"platform_id,omitempty"`   //平台ID
	StoreId     int64  `json:"store_id,omitempty"`      //店铺ID
	StoreShopId int64  `json:"store_shop_id,omitempty"` //店铺分店ID
	StoreRegion string `json:"store_region,omitempty"`  //店铺分区
	ClientId    string `json:"client_id,omitempty"`     //ClientID
	DisplayName string `json:"display_name,omitempty"`  //显示名称
	Username    string `json:"username,omitempty"`      //登录账号
}

//是否为超级管理员
func (a *User) HasAdmin() bool {
	list := []string{"master_admin", "store_admin"}
	if helper.InArray(list, a.Mode) {
		return true
	}
	return false
}

//是否为管理员（内部用户）
func (a *User) HasManager() bool {
	list := []string{
		"master_admin", "master_manager", "master_user",
		"store_admin", "store_manager", "store_user", "store_seller",
	}
	if helper.InArray(list, a.Mode) {
		return true
	}
	return false
}

//是否为客户/会员
func (a *User) HasCustomer() bool {
	list := []string{"master_customer", "master_member", "store_customer"}
	if helper.InArray(list, a.Mode) {
		return true
	}
	return false
}

//是否为网站
func (a *User) HasWebsite() bool {
	list := []string{"master_site", "store_site"}
	if helper.InArray(list, a.Mode) {
		return true
	}
	return false
}

//是否为商铺模式
func (a *User) HasStoreMode() bool {
	if strings.HasPrefix(a.Mode, "store_") {
		return true
	}
	return false
}

//是否为主铺模式
func (a *User) HasMasterMode() bool {
	if strings.HasPrefix(a.Mode, "master_") {
		return true
	}
	return false
}

//用户是否已登录(mode为store_site 除外)
func (a *User) HasLogin() bool {
	if a.UserId > 0 {
		return true
	}
	return false
}

//获得当前授权用户
func GetUser(ctx context.Context) *User {
	ret := &User{
		Mode:        getMetaValue(ctx, "mode"),
		DisplayName: getMetaValue(ctx, "display_name"),
		UserId:      helper.StringToInt64(getMetaValue(ctx, "user_id")),
		PlatformId:  helper.StringToInt64(getMetaValue(ctx, "platform_id")),
		StoreId:     helper.StringToInt64(getMetaValue(ctx, "store_id")),
		StoreShopId: helper.StringToInt64(getMetaValue(ctx, "store_shop_id")),
		StoreRegion: getMetaValue(ctx, "store_region"),
		SessionKey:  getMetaValue(ctx, "session_id"),
		ClientId:    getMetaValue(ctx, "client_id"),
	}
	return ret
}

//获得当前授权用户(通过HttpHeader)
func GetUserByHttpHeader(header http.Header) *User {
	ret := &User{
		Mode:        header.Get("Auth-Mode"),
		DisplayName: header.Get("Auth-Display-Name"),
		UserId:      helper.StringToInt64(header.Get("Auth-User-Id")),
		PlatformId:  helper.StringToInt64(header.Get("Auth-Platform-Id")),
		StoreId:     helper.StringToInt64(header.Get("Auth-Store-Id")),
		StoreShopId: helper.StringToInt64(header.Get("Auth-Store-Shop-Id")),
		StoreRegion: header.Get("Auth-Store-Region"),
		SessionKey:  header.Get("Auth-Session-Key"),
		ClientId:    header.Get("Auth-Client-Id"),
	}
	return ret
}

//获得当前授权用户(通过GinHeader)
func GetUserByGinHeader(ctx *gin.Context) *User {
	ret := &User{
		Mode:        ctx.GetHeader("Auth-Mode"),
		DisplayName: ctx.GetHeader("Auth-Display-Name"),
		UserId:      helper.StringToInt64(ctx.GetHeader("Auth-User-Id")),
		PlatformId:  helper.StringToInt64(ctx.GetHeader("Auth-Platform-Id")),
		StoreId:     helper.StringToInt64(ctx.GetHeader("Auth-Store-Id")),
		StoreShopId: helper.StringToInt64(ctx.GetHeader("Auth-Store-Shop-Id")),
		StoreRegion: ctx.GetHeader("Auth-Store-Region"),
		SessionKey:  ctx.GetHeader("Auth-Session-Key"),
		ClientId:    ctx.GetHeader("Auth-Client-Id"),
	}
	return ret
}

func SetUser(ctx *gin.Context, user *User) {
	if ctx.Keys == nil {
		ctx.Keys = make(map[string]interface{})
	}
	ctx.Keys["mode"] = user.Mode
	ctx.Keys["display_name"] = user.DisplayName
	ctx.Keys["user_id"] = user.UserId
	ctx.Keys["store_id"] = user.StoreId
	ctx.Keys["platform_id"] = user.PlatformId
	ctx.Keys["store_shop_id"] = user.StoreShopId
	ctx.Keys["store_region"] = user.StoreRegion
	ctx.Keys["session_key"] = user.SessionKey
	ctx.Keys["client_id"] = user.ClientId
	//Pass on
	ctx.Next()
}

//获得当前用户类型:（customer/user）
func GetUserType(ctx context.Context) string {
	u := GetUser(ctx)
	if u != nil {
		if u.HasCustomer() {
			return "customer"
		} else {
			return "user"
		}
	}
	return ""
}

//获得当前用户ID
func GetUserId(ctx context.Context) int64 {
	u := GetUser(ctx)
	if u != nil {
		return u.UserId
	}
	return 0
}

//获得当前客户ID
func GetCustomerId(ctx context.Context) int64 {
	u := GetUser(ctx)
	if u != nil && u.HasCustomer() {
		return u.UserId
	}
	return 0
}

//获得当前店铺ID
func GetStoreId(ctx context.Context) int64 {
	val := ctx.Value("store_id")
	if val != nil {
		v := helper.StringToInt64(helper.ToString(val))
		return v
	}
	return 0
}

func getMetaValue(ctx context.Context, key string) string {
	val := ctx.Value(key)
	if val != nil {
		v := helper.ToString(val)
		return v
	}
	return ""
}

func GetStoreFlag(id int64, prefix ...string) string {
	flag := fmt.Sprintf("%08d", id)
	if id == 1 {
		flag = "master"
	}
	p := "go_store_"
	if prefix != nil {
		p = prefix[0]
	}
	flag = p + flag
	return flag
}
