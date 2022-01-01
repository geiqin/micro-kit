package auth

import (
	"context"
	"fmt"
	"github.com/geiqin/gotools/helper"
)

//获得当前用户Mode
func GetMode(ctx context.Context) string {
	val := ctx.Value("mode")
	if val != nil {
		v := val.(string)
		return v
	}
	return ""
}

//获得当前用户会话ID
func GetSessionId(ctx context.Context) string {
	val := ctx.Value("session_id")
	if val != nil {
		v := val.(string)
		return v
	}
	return ""
}

//获得当前ClientID
func GetClientId(ctx context.Context) string {
	val := ctx.Value("client_id")
	if val != nil {
		v := val.(string)
		return v
	}
	return ""
}

//获得当前管理员ID
func GetManagerId(ctx context.Context) int64 {
	val := ctx.Value("manager_id")
	if val != nil {
		v := helper.StringToInt64(val.(string))
		return v
	}
	return 0
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

//获得当前店铺卖家ID
func GetStoreSellerId(ctx context.Context) int64 {
	val := ctx.Value("store_seller_id")
	if val != nil {
		v := helper.StringToInt64(val.(string))
		return v
	}
	return 0
}

//获得当前店铺用户ID
func GetStoreUserId(ctx context.Context) int64 {
	val := ctx.Value("store_user_id")
	if val != nil {
		v := helper.StringToInt64(val.(string))
		return v
	}
	return 0
}

//获得当前客户ID
func GetStoreCustomerId(ctx context.Context) int64 {
	val := ctx.Value("store_customer_id")
	if val != nil {
		v := helper.StringToInt64(val.(string))
		return v
	}
	return 0
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
