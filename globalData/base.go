package globalData

import (
	"context"
	"github.com/geiqin/micro-kit/utils"
	"strings"
)

//获得当前应用
func GetApplication(ctx context.Context) string {
	return getMetaValue(ctx, "application")
}

//获得当前应用终端
func GetApplicationClientType(ctx context.Context) string {
	return getMetaValue(ctx, "application_client_type")
}

func getMetaValue(ctx context.Context, key string) string {
	val := ctx.Value(key)
	if val != nil {
		v := utils.ToString(val)
		return v
	}
	return ""
}

//是否微信环境
func IsWeixinEnv(ctx context.Context) bool {
	userAgent := GetHttpUserAgent(ctx)
	if userAgent != "" && strings.Index(userAgent, "MicroMessenger") != -1 {
		return true
	}
	return false
}

//是否微信环境
func IsDingdingEnv(ctx context.Context) bool {
	userAgent := GetHttpUserAgent(ctx)
	if userAgent != "" && strings.Index(userAgent, "DingTalk") != -1 {
		return true
	}
	return false
}

//是否QQ环境
func IsQQEnv(ctx context.Context) bool {
	userAgent := GetHttpUserAgent(ctx)
	if userAgent != "" && strings.Index(userAgent, "QQ/") != -1 {
		return true
	}
	return false
}

//是否支付宝环境
func IsAlipayEnv(ctx context.Context) bool {
	userAgent := GetHttpUserAgent(ctx)
	if userAgent != "" && strings.Index(userAgent, "AlipayClient") != -1 {
		return true
	}
	return false
}

//是否新浪微博环境
func IsWeiboEnv(ctx context.Context) bool {
	userAgent := GetHttpUserAgent(ctx)
	if userAgent != "" && strings.Index(userAgent, "Weibo") != -1 {
		return true
	}
	return false
}

func GetHttpUserAgent(ctx context.Context) string {
	val := ctx.Value("Http-User-Agent")
	if val != nil {
		v := utils.ToString(val)
		return v
	}
	return ""
}
