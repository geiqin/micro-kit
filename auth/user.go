package auth

import (
	"context"
	"fmt"
	"github.com/geiqin/micro-kit/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

//当前授权用户
type User struct {
	Mode        string            `json:"mode"`                   //授权模式
	SessionKey  string            `json:"session_key,omitempty"`  //会话Key
	UserId      int64             `json:"user_id,omitempty"`      //用户ID
	PlatformId  int64             `json:"platform_id,omitempty"`  //平台ID
	StoreId     int64             `json:"store_id,omitempty"`     //店铺ID
	RealstoreId int64             `json:"realstore_id,omitempty"` //多门店ID
	ShopId      int64             `json:"shop_id,omitempty"`      //多商户ID
	ClientId    string            `json:"client_id,omitempty"`    //ClientID
	Nickname    string            `json:"nickname,omitempty"`     //用户昵称
	Extends     map[string]string `json:"extends,omitempty"`      //附加信息
}

/*
type DataPermission struct {
	DataScope string `json:"data_scope"`
	StaffId   int64  `json:"staff_id"`
	DeptId    int32  `json:"dept_id"`
	RoleId    int32  `json:"role_id"`
	PostId    int32  `json:"post_id"`
}

*/

func (a *User) ExtendsToJson() string {
	var str string
	if a != nil && a.Extends != nil {
		str = utils.JsonEncode(a.Extends)
	}
	return str
}

func (a *User) ExtendsFromJson(strJson string) map[string]string {
	var ret map[string]string
	if strJson != "" {
		utils.JsonDecode(strJson, &ret)
		a.Extends = ret
	}
	return ret
}

//是否为超级管理员
func (a *User) HasAdmin() bool {
	list := []string{"master_admin", "store_admin"}
	if utils.InArray(list, a.Mode) {
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
	if utils.InArray(list, a.Mode) {
		return true
	}
	return false
}

//是否为客户/会员
func (a *User) HasMember() bool {
	list := []string{"master_member", "master_member", "store_member"}
	if utils.InArray(list, a.Mode) {
		return true
	}
	return false
}

//是否为网站
func (a *User) HasWebsite() bool {
	list := []string{"master_site", "store_site"}
	if utils.InArray(list, a.Mode) {
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

//获得当前用户类型:（member/user）
func (a *User) GetUserType() string {
	if a.UserId > 0 {
		if a.HasMember() {
			return "member"
		} else {
			return "user"
		}
	}
	return ""
}

//获得当前授权用户
func GetUser(ctx context.Context) *User {
	ret := &User{
		Mode:        getMetaValue(ctx, "Auth-Mode"),
		Nickname:    getMetaValue(ctx, "Auth-Nickname"),
		UserId:      utils.StringToInt64(getMetaValue(ctx, "Auth-User-Id")),
		PlatformId:  utils.StringToInt64(getMetaValue(ctx, "Auth-Platform-Id")),
		StoreId:     utils.StringToInt64(getMetaValue(ctx, "Auth-Store-Id")),
		RealstoreId: utils.StringToInt64(getMetaValue(ctx, "Auth-Realstore-Id")),
		ShopId:      utils.StringToInt64(getMetaValue(ctx, "Auth-Shop-Id")),
		SessionKey:  getMetaValue(ctx, "Auth-Session-Key"),
		ClientId:    getMetaValue(ctx, "Auth-Client-Id"),
	}
	dp := getMetaValue(ctx, "Auth-Extends")
	if dp != "" {
		utils.JsonDecode(dp, &ret.Extends)
	}
	return ret
}

//获得当前授权用户(通过HttpHeader)
func GetUserByHttpHeader(header http.Header) *User {
	ret := &User{
		Mode:        header.Get("Auth-Mode"),
		Nickname:    header.Get("Auth-Nickname"),
		UserId:      utils.StringToInt64(header.Get("Auth-User-Id")),
		PlatformId:  utils.StringToInt64(header.Get("Auth-Platform-Id")),
		StoreId:     utils.StringToInt64(header.Get("Auth-Store-Id")),
		RealstoreId: utils.StringToInt64(header.Get("Auth-Realstore-Id")),
		ShopId:      utils.StringToInt64(header.Get("Auth-Shop-Id")),
		SessionKey:  header.Get("Auth-Session-Key"),
		ClientId:    header.Get("Auth-Client-Id"),
	}
	dp := header.Get("Auth-Extends")
	if dp != "" {
		utils.JsonDecode(dp, &ret.Extends)
	}
	return ret
}

//获得当前授权用户(通过GinHeader)
func GetUserByGinHeader(ctx *gin.Context) *User {
	ret := &User{
		Mode:        ctx.GetHeader("Auth-Mode"),
		Nickname:    ctx.GetHeader("Auth-Nickname"),
		UserId:      utils.StringToInt64(ctx.GetHeader("Auth-User-Id")),
		PlatformId:  utils.StringToInt64(ctx.GetHeader("Auth-Platform-Id")),
		StoreId:     utils.StringToInt64(ctx.GetHeader("Auth-Store-Id")),
		RealstoreId: utils.StringToInt64(ctx.GetHeader("Auth-Realstore-Id")),
		ShopId:      utils.StringToInt64(ctx.GetHeader("Auth-Shop-Id")),
		SessionKey:  ctx.GetHeader("Auth-Session-Key"),
		ClientId:    ctx.GetHeader("Auth-Client-Id"),
	}
	dp := ctx.GetHeader("Auth-Extends")
	if dp != "" {
		utils.JsonDecode(dp, &ret.Extends)
	}
	return ret
}

func SetUser(ctx *gin.Context, user *User) {
	if ctx.Keys == nil {
		ctx.Keys = make(map[string]interface{})
	}
	ctx.Keys["Auth-Mode"] = user.Mode
	ctx.Keys["Auth-Nickname"] = user.Nickname
	ctx.Keys["Auth-User-Id"] = user.UserId
	ctx.Keys["Auth-Store-Id"] = user.StoreId
	ctx.Keys["Auth-Platform-Id"] = user.PlatformId
	ctx.Keys["Auth-Realstore-Id"] = user.RealstoreId
	ctx.Keys["Auth-Shop-Id"] = user.ShopId
	ctx.Keys["Auth-Session-Key"] = user.SessionKey
	ctx.Keys["Auth-Client-Id"] = user.ClientId
	ctx.Keys["Auth-Extends"] = user.Extends
	//Pass on
	ctx.Next()
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
func GetMemberId(ctx context.Context) int64 {
	u := GetUser(ctx)
	if u != nil && u.HasMember() {
		return u.UserId
	}
	return 0
}

//获得当前店铺ID
func GetStoreId(ctx context.Context) int64 {
	val := ctx.Value(StoreIdKey)
	if val != nil {
		v := utils.StringToInt64(utils.ToString(val))
		return v
	}
	return 0
}

func getMetaValue(ctx context.Context, key string) string {
	val := ctx.Value(key)
	if val != nil {
		v := utils.ToString(val)
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
