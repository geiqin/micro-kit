package auth

import (
	"context"
	"github.com/geiqin/gotools/helper"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/metadata"
	"net/http"
)

func StoreContext(storeId int64) context.Context {
	storeIdStr := helper.Int64ToString(storeId)
	ctx := context.WithValue(context.Background(), "store_id", storeIdStr)
	ctx = metadata.NewContext(ctx, map[string]string{
		"Auth-Store-Id": storeIdStr,
	})
	return ctx
}

func StoreContextByString(storeId string) context.Context {
	ctx := context.WithValue(context.Background(), "store_id", storeId)
	ctx = metadata.NewContext(ctx, map[string]string{
		"Auth-Store-Id": storeId,
	})
	return ctx
}

func StoreContextByBroker(p broker.Event) context.Context {
	if p != nil && p.Message().Header != nil {
		storeId := p.Message().Header["store_id"]
		if storeId != "" {
			sid := helper.StringToInt64(storeId)
			if sid > 0 {
				ctx := context.WithValue(context.Background(), "store_id", storeId)
				ctx = metadata.NewContext(ctx, map[string]string{
					"Auth-Store-Id": storeId,
				})
				return ctx
			}
		}
	}
	return context.Background()
}

//主要提供给 WEB模式下使用
func StoreContextByHeader(header http.Header) context.Context {
	mode := header.Get("Auth-Mode")
	sessionKey := header.Get("Auth-Session-Key")
	clientId := header.Get("Auth-Client-Id")
	nickname := header.Get("Auth-Nickname")
	userId := header.Get("Auth-User-Id")
	storeId := header.Get("Auth-Store-Id")
	storeShopId := header.Get("Auth-Store-Shop-Id")
	storeRegion := header.Get("Auth-Store-Region")
	platformId := header.Get("Auth-Platform-Id")
	permission := header.Get("Auth-Data-Permission")
	application := header.Get("Application")
	applicationClientType := header.Get("Application-Client-Type")

	ctx := context.Background()
	if mode != "" {
		ctx = context.WithValue(ctx, "mode", storeId)
		ctx = metadata.NewContext(ctx, map[string]string{
			"Auth-Mode": storeId,
		})
	}
	if sessionKey != "" {
		ctx = context.WithValue(ctx, "session_key", sessionKey)
		ctx = metadata.NewContext(ctx, map[string]string{
			"Auth-Session-Key": sessionKey,
		})
	}
	if nickname != "" {
		ctx = context.WithValue(ctx, "nickname", nickname)
		ctx = metadata.NewContext(ctx, map[string]string{
			"Auth-Nickname": sessionKey,
		})
	}
	if storeId != "" {
		ctx = context.WithValue(ctx, "store_id", storeId)
		ctx = metadata.NewContext(ctx, map[string]string{
			"Auth-Store-Id": storeId,
		})
	}
	if storeShopId != "" {
		ctx = context.WithValue(ctx, "store_shop_id", storeShopId)
		ctx = metadata.NewContext(ctx, map[string]string{
			"Auth-Store-Shop-Id": storeShopId,
		})
	}
	if platformId != "" {
		ctx = context.WithValue(ctx, "platform_id", platformId)
		ctx = metadata.NewContext(ctx, map[string]string{
			"Auth-Platform-Id": platformId,
		})
	}
	if storeRegion != "" {
		ctx = context.WithValue(ctx, "store_region", storeRegion)
		ctx = metadata.NewContext(ctx, map[string]string{
			"Auth-Store-Region": storeRegion,
		})
	}
	if userId != "" {
		ctx = context.WithValue(ctx, "user_id", userId)
		ctx = metadata.NewContext(ctx, map[string]string{
			"Auth-User-Id": userId,
		})
	}
	if clientId != "" {
		ctx = context.WithValue(ctx, "client_id", clientId)
		ctx = metadata.NewContext(ctx, map[string]string{
			"Auth-Client-Id": clientId,
		})
	}
	if permission != "" {
		ctx = context.WithValue(ctx, "data_permission", permission)
		ctx = metadata.NewContext(ctx, map[string]string{
			"Auth-DataPermission": permission,
		})
	}
	if application != "" {
		ctx = context.WithValue(ctx, "application", application)
		ctx = metadata.NewContext(ctx, map[string]string{
			"Application": application,
		})
	}
	if applicationClientType != "" {
		ctx = context.WithValue(ctx, "application_client_type", applicationClientType)
		ctx = metadata.NewContext(ctx, map[string]string{
			"Application-Client-Type": applicationClientType,
		})
	}

	return ctx
}
