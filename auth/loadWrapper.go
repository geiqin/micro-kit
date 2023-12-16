package auth

import (
	"context"
	"errors"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/server"
)

// AuthWrapper 是一个高阶函数，入参是 ”下一步“ 函数，出参是认证函数
// 在返回的函数内部处理完认证逻辑后，再手动调用 fn() 进行下一步处理
// token 是从 consignment-ci 上下文中取出的，再调用 user-service 将其做验证
// 认证通过则 fn() 继续执行，否则报错
func LoadWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("no auth meta-data found in request")
		}

		mode := meta["Auth-Mode"]
		userId := meta["Auth-User-Id"]
		nickname := meta["Auth-Nickname"]
		sessionKey := meta["Auth-Session-Key"]
		platformId := meta["Auth-Platform-Id"]
		storeId := meta["Auth-Store-Id"]
		realstoreId := meta["Auth-Realstore-Id"]
		shopId := meta["Auth-Shop-Id"]
		clientId := meta["Auth-Client-Id"]
		extends := meta["Auth-Extends"]
		application := meta["Application"]
		applicationClientType := meta["Application-Client-Type"]

		if mode != "" {
			ctx = context.WithValue(ctx, "mode", mode)
		}
		if extends != "" {
			ctx = context.WithValue(ctx, "extends", extends)
		}
		if userId != "" {
			ctx = context.WithValue(ctx, "user_id", userId)
		}
		if nickname != "" {
			ctx = context.WithValue(ctx, "nickname", nickname)
		}
		if storeId != "" {
			ctx = context.WithValue(ctx, "store_id", storeId)
		}
		if realstoreId != "" {
			ctx = context.WithValue(ctx, "realstore_id", realstoreId)
		}
		if platformId != "" {
			ctx = context.WithValue(ctx, "platform_id", platformId)
		}
		if shopId != "" {
			ctx = context.WithValue(ctx, "shop_id", shopId)
		}
		if sessionKey != "" {
			ctx = context.WithValue(ctx, "session_key", sessionKey)
		}
		if clientId != "" {
			ctx = context.WithValue(ctx, "client_id", clientId)
		}
		if application != "" {
			ctx = context.WithValue(ctx, "application", application)
		}
		if applicationClientType != "" {
			ctx = context.WithValue(ctx, "application_client_type", applicationClientType)
		}

		//继续执行下一步处理
		err := fn(ctx, req, resp)
		return err
	}
}
