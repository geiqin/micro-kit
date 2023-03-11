package globalData

import (
	"context"
	"github.com/geiqin/gotools/helper"
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
		v := helper.ToString(val)
		return v
	}
	return ""
}
