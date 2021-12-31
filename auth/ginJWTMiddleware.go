package auth

import (
	"github.com/gin-gonic/gin"
)

type HeaderGrantInfo struct {
	Mode            string `json:"mode"`
	UserId          string `json:"user_id"`
	StoreId         string `json:"store_id"`
	StoreManagerId  string `json:"store_manager_id"`
	StoreSellerId   string `json:"store_seller_id"`
	StoreHumanId    string `json:"store_human_id"`
	StoreCustomerId string `json:"store_customer_id"`
	SessionId       string `json:"session_id"`
	ClientId        string `json:"client_id"`
}

func GinJWTMiddleware(ctx *gin.Context) {
	g := &HeaderGrantInfo{
		Mode:            ctx.GetHeader("Auth-Mode"),
		UserId:          ctx.GetHeader("Auth-User-Id"),
		StoreId:         ctx.GetHeader("Auth-Store-Id"),
		StoreManagerId:  ctx.GetHeader("Auth-Store-Manager-Id"),
		StoreSellerId:   ctx.GetHeader("Auth-Store-Seller-Id"),
		StoreHumanId:    ctx.GetHeader("Auth-Store-Human-Id"),
		StoreCustomerId: ctx.GetHeader("Auth-Store-Customer-Id"),
		SessionId:       ctx.GetHeader("Session-Id"),
		ClientId:        ctx.GetHeader("Client-Id"),
	}

	if ctx.Keys == nil {
		ctx.Keys = make(map[string]interface{})
	}

	if g.UserId != "" {
		ctx.Keys["user_id"] = g.UserId
	}
	if g.StoreId != "" {
		ctx.Keys["store_id"] = g.StoreId
	}
	if g.StoreManagerId != "" {
		ctx.Keys["store_manager_id"] = g.StoreManagerId
	}
	if g.StoreSellerId != "" {
		ctx.Keys["store_seller_id"] = g.StoreSellerId
	}
	if g.StoreHumanId != "" {
		ctx.Keys["store_human_id"] = g.StoreHumanId
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
