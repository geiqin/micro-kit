package auth

import (
	"github.com/gin-gonic/gin"
)

type HeaderGrantInfo struct {
	Mode            string `json:"mode"`
	ManagerId       string `json:"manager_id"`
	StoreId         string `json:"store_id"`
	StoreUserId     string `json:"store_user_id"`
	StoreSellerId   string `json:"store_seller_id"`
	StoreCustomerId string `json:"store_customer_id"`
	SessionId       string `json:"session_id"`
	ClientId        string `json:"client_id"`
}

func GinJWTMiddleware(ctx *gin.Context) {
	g := &HeaderGrantInfo{
		Mode:            ctx.GetHeader("Auth-Mode"),
		StoreId:         ctx.GetHeader("Auth-Store-Id"),
		ManagerId:       ctx.GetHeader("Auth-Manager-Id"),
		StoreSellerId:   ctx.GetHeader("Auth-Store-Seller-Id"),
		StoreUserId:     ctx.GetHeader("Auth-Store-User-Id"),
		StoreCustomerId: ctx.GetHeader("Auth-Store-Customer-Id"),
		SessionId:       ctx.GetHeader("Session-Id"),
		ClientId:        ctx.GetHeader("Client-Id"),
	}

	if ctx.Keys == nil {
		ctx.Keys = make(map[string]interface{})
	}

	if g.StoreUserId != "" {
		ctx.Keys["user_id"] = g.StoreUserId
	}
	if g.StoreId != "" {
		ctx.Keys["store_id"] = g.StoreId
	}
	if g.ManagerId != "" {
		ctx.Keys["manager_id"] = g.ManagerId
	}
	if g.StoreSellerId != "" {
		ctx.Keys["store_seller_id"] = g.StoreSellerId
	}
	if g.StoreUserId != "" {
		ctx.Keys["store_user_id"] = g.StoreUserId
	}
	if g.StoreCustomerId != "" {
		ctx.Keys["store_customer_id"] = g.StoreCustomerId
	}
	if g.SessionId != "" {
		ctx.Keys["session_id"] = g.SessionId
	}
	if g.Mode != "" {
		ctx.Keys["mode"] = g.Mode
	}
	if g.ClientId != "" {
		ctx.Keys["client_id"] = g.ClientId
	}

	//取值方式
	//ctx.Get(key)

	//Pass on
	ctx.Next()
}
