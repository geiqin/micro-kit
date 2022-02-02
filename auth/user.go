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
	StoreId     int64  `json:"store_id,omitempty"`      //店铺ID
	StoreShopId int64  `json:"store_shop_id,omitempty"` //店铺分店ID
	ClientId    string `json:"client_id,omitempty"`     //ClientID
	DisplayName string `json:"display_name,omitempty"`  //显示名称
	Username    string `json:"username,omitempty"`      //登录账号
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

//是否为店铺模式
func (a *User) HasStoreMode() bool {
	if strings.HasPrefix(a.Mode, "store_") {
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

//是否为超级用户（分平台级和店铺级）
func (a *User) HasAdmin() bool {
	if strings.HasSuffix(a.Mode, "admin") {
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
		StoreId:     helper.StringToInt64(ctx.GetHeader("Auth-Store-Id")),
		StoreShopId: helper.StringToInt64(ctx.GetHeader("Auth-Store-Shop-Id")),
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

//获得当前店铺用户ID
func GetManagerId(ctx context.Context) int64 {
	u := GetUser(ctx)
	if u.HasManager() {
		return u.UserId
	}
	return 0
}

//获得当前店铺用户ID
func GetStoreUserId(ctx context.Context) int64 {
	u := GetUser(ctx)
	if u.HasStoreUser() {
		return u.UserId
	}
	return 0
}

//获得当前店铺客户ID
func GetStoreCustomerId(ctx context.Context) int64 {
	u := GetUser(ctx)
	if u.HasStoreCustomer() {
		return u.UserId
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
