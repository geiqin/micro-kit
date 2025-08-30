package auth

import (
	"context"
	"github.com/geiqin/micro-kit/utils"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/metadata"
	"net/http"
)

func StoreContext(storeId int64) context.Context {
	storeIdStr := utils.Int64ToString(storeId)
	ctx := context.WithValue(context.Background(), StoreIdKey, storeIdStr)
	ctx = metadata.NewContext(ctx, map[string]string{
		StoreIdKey: storeIdStr,
	})
	return ctx
}

func StoreContextByString(storeId string) context.Context {
	ctx := context.WithValue(context.Background(), StoreIdKey, storeId)
	ctx = metadata.NewContext(ctx, map[string]string{
		StoreIdKey: storeId,
	})
	return ctx
}

func StoreContextByBroker(p broker.Event) context.Context {
	if p != nil && p.Message().Header != nil {
		storeId := p.Message().Header["store_id"]
		if storeId != "" {
			sid := utils.StringToInt64(storeId)
			if sid > 0 {
				ctx := context.WithValue(context.Background(), StoreIdKey, storeId)
				ctx = metadata.NewContext(ctx, map[string]string{
					StoreIdKey: storeId,
				})
				return ctx
			}
		}
	}
	return context.Background()
}

//主要提供给 WEB模式下使用
func StoreContextByHeader(header http.Header) context.Context {
	ctx := context.Background()
	for _, k := range inContainKeys {
		val := header.Get(k)
		if val != "" {
			ctx = context.WithValue(ctx, k, val)
			ctx = metadata.NewContext(ctx, map[string]string{
				k: val,
			})
		}
	}

	/*
		storeId := header.Get("Auth-Store-Id")
		if storeId != "" {
			ctx = context.WithValue(ctx, "store_id", storeId)
			ctx = metadata.NewContext(ctx, map[string]string{
				"Auth-Store-Id": storeId,
			})
		}
	*/

	return ctx
}
