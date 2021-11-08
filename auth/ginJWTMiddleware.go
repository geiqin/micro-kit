package auth

import (
	"github.com/gin-gonic/gin"
)

type HeaderGrantInfo struct {
	Mode       string `json:"mode"`
	UserId     string `json:"user_id"`
	StoreId    string `json:"store_id"`
	SellerId   string `json:"seller_id"`
	CustomerId string `json:"customer_id"`
	SessionId  string `json:"session_id"`
	ClientId   string `json:"client_id"`
}

func GinJWTMiddleware(ctx *gin.Context) {
	g := &HeaderGrantInfo{
		Mode:       ctx.GetHeader("Auth-Mode"),
		UserId:     ctx.GetHeader("Auth-User-Id"),
		StoreId:    ctx.GetHeader("Auth-Store-Id"),
		SellerId:   ctx.GetHeader("Auth-Seller-Id"),
		CustomerId: ctx.GetHeader("Auth-Customer-Id"),
		SessionId:  ctx.GetHeader("Session-Id"),
		ClientId:   ctx.GetHeader("Client-Id"),
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
	if g.SellerId != "" {
		ctx.Keys["seller_id"] = g.SellerId
	}
	if g.CustomerId != "" {
		ctx.Keys["customer_id"] = g.CustomerId
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
