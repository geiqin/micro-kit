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
	Mode        string `json:"mode"`
	SessionKey  string `json:"session_key"`
	UserId      int64  `json:"user_id"`
	StoreId     int64  `json:"store_id"`
	StoreShopId int64  `json:"store_shop_id"`
	ClientId    string `json:"client_id"`
	DisplayName string `json:"display_name"`       //显示名称
	Username    string `json:"username,omitempty"` //登录账号
}

//是否为店铺客户
func (a *User) HasStoreCustomer() bool {
	if a.Mode == "store_customer" {
		return true
	}
	return false
}

//是否为店铺网站
func (a *User) HasStoreSite() bool {
	if a.Mode == "store_site" {
		return true
	}
	return false
}

//是否为店铺用户
func (a *User) HasStoreUser() bool {
	if a.Mode == "store_admin" || a.Mode == "store_user" {
		return true
	}
	return false
}

//是否为平台用户
func (a *User) HasManager() bool {
	if a.Mode == "admin" || a.Mode == "manager" {
		return true
	}
	return false
}

//是否为店铺模式
func (a *User) HasStoreMode() bool {
	if strings.HasPrefix(a.Mode, "store_") {
		return true
	}
	return false
}

/*
//用户是否已登录(mode为store_site 除外)
func (a *User) GetDisplayName() string {
	if a.displayName == "" {
		return a.displayName
	}
	return ""
}

//获取用户登录账号
func (a *User) GetUsername() string {
	if a.username == "" {
		return a.username
	}
	return ""
}

//从缓存获取用户信息
func (a *User) loadUserInfo() string {
	if a.username == "" {
		return a.username
	}
	return ""
}

*/

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
		StoreId:     helper.StringToInt64(getMetaValue(ctx, "store_id")),
		StoreShopId: helper.StringToInt64(getMetaValue(ctx, "store_shop_id")),
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
		StoreId:     helper.StringToInt64(header.Get("Auth-Store-Id")),
		StoreShopId: helper.StringToInt64(header.Get("Auth-Store-Shop-Id")),
		SessionKey:  header.Get("Session-Key"),
		ClientId:    header.Get("Client-Id"),
	}
	return ret
}

//获得当前授权用户(通过GinHeader)
func GetUserByGinHeader(ctx *gin.Context) *User {
	ret := &User{
		Mode:        ctx.GetHeader("Auth-Mode"),
		DisplayName: ctx.GetHeader("Auth-Display-Name"),
		UserId:      helper.StringToInt64(ctx.GetHeader("Auth-User-Id")),
		StoreId:     helper.StringToInt64(ctx.GetHeader("Auth-Store-Id")),
		StoreShopId: helper.StringToInt64(ctx.GetHeader("Auth-Store-Shop-Id")),
		SessionKey:  ctx.GetHeader("Session-Key"),
		ClientId:    ctx.GetHeader("Client-Id"),
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
	ctx.Keys["store_shop_id"] = user.StoreShopId
	ctx.Keys["session_key"] = user.SessionKey
	ctx.Keys["client_id"] = user.ClientId
	//Pass on
	ctx.Next()
}

//获得当前店铺ID
func GetStoreId(ctx context.Context) int64 {
	val := ctx.Value("store_id")
	if val != nil {
		v := helper.StringToInt64(val.(string))
		return v
	}
	return 0
}

func getMetaValue(ctx context.Context, key string) string {
	val := ctx.Value(key)
	if val != nil {
		v := val.(string)
		return v
	}
	return ""
}

func GetStoreFlag(id int64, prefix ...string) string {
	flag := fmt.Sprintf("%08d", id)
	p := "go_store_"
	if prefix != nil {
		p = prefix[0]
	}
	flag = p + flag
	return flag
}
