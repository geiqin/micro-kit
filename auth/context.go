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
	sessionId := header.Get("Auth-Session-Id")
	managerId := header.Get("Auth-Manager-Id")
	storeId := header.Get("Auth-Store-Id")
	storeShopId := header.Get("Auth-Store-Shop-Id")
	storeUserId := header.Get("Auth-Store-User-Id")
	storeEmployeeId := header.Get("Auth-Store-Employee-Id")
	storeCustomerId := header.Get("Auth-Store-Customer-Id")

	ctx := context.Background()
	if mode != "" {
		ctx = context.WithValue(ctx, "mode", storeId)
		ctx = metadata.NewContext(ctx, map[string]string{
			"Auth-Mode": storeId,
		})
	}
	if sessionId != "" {
		ctx = context.WithValue(ctx, "session_id", sessionId)
		ctx = metadata.NewContext(ctx, map[string]string{
			"Auth-Session-Id": sessionId,
		})
	}
	if storeId != "" {
		ctx = context.WithValue(ctx, "store_id", storeId)
		ctx = metadata.NewContext(ctx, map[string]string{
			"Auth-Store-Id": storeId,
		})
	}
	if managerId != "" {
		ctx = context.WithValue(ctx, "manager_id", managerId)
		ctx = metadata.NewContext(ctx, map[string]string{
			"Auth-Manager-Id": managerId,
		})
	}
	if storeShopId != "" {
		ctx = context.WithValue(ctx, "store_shop_id", storeShopId)
		ctx = metadata.NewContext(ctx, map[string]string{
			"Auth-Store-Shop-Id": storeShopId,
		})
	}
	if storeEmployeeId != "" {
		ctx = context.WithValue(ctx, "store_employee_id", storeEmployeeId)
		ctx = metadata.NewContext(ctx, map[string]string{
			"Auth-Store-Employee-Id": storeEmployeeId,
		})
	}
	if storeUserId != "" {
		ctx = context.WithValue(ctx, "store_user_id", storeUserId)
		ctx = metadata.NewContext(ctx, map[string]string{
			"Auth-Store-User-Id": storeUserId,
		})
	}
	if storeCustomerId != "" {
		ctx = context.WithValue(ctx, "store_customer_id", storeCustomerId)
		ctx = metadata.NewContext(ctx, map[string]string{
			"Auth-Store-Customer-Id": storeCustomerId,
		})
	}
	return ctx
}
