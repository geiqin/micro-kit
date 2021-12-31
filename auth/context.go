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
	userId := header.Get("Auth-User-Id")
	storeId := header.Get("Auth-Store-Id")
	storeManagerId := header.Get("Auth-Store-Manager-Id")
	storeSellerId := header.Get("Auth-Store-Seller-Id")
	storeHumanId := header.Get("Auth-Store-Human-Id")
	storeCustomerId := header.Get("Auth-Customer-Id")

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
	if userId != "" {
		ctx = context.WithValue(ctx, "user_id", userId)
		ctx = metadata.NewContext(ctx, map[string]string{
			"Auth-User-Id": userId,
		})
	}
	if storeId != "" {
		ctx = context.WithValue(ctx, "store_id", storeId)
		ctx = metadata.NewContext(ctx, map[string]string{
			"Auth-Store-Id": storeId,
		})
	}
	if storeManagerId != "" {
		ctx = context.WithValue(ctx, "store_manager_id", storeManagerId)
		ctx = metadata.NewContext(ctx, map[string]string{
			"Auth-Store-Manager-Id": storeManagerId,
		})
	}
	if storeSellerId != "" {
		ctx = context.WithValue(ctx, "store_seller_id", storeSellerId)
		ctx = metadata.NewContext(ctx, map[string]string{
			"Auth-Store-Seller-Id": storeSellerId,
		})
	}
	if storeHumanId != "" {
		ctx = context.WithValue(ctx, "store_human_id", storeHumanId)
		ctx = metadata.NewContext(ctx, map[string]string{
			"Auth-Store-Human-Id": storeHumanId,
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
