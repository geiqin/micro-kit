package auth

import (
	"github.com/gin-gonic/gin"
)

type HeaderGrantInfo struct {
	Mode            string `json:"mode"`
	ManagerId       string `json:"manager_id"`
	StoreId         string `json:"store_id"`
	StoreUserId     string `json:"store_user_id"`
	StoreShopId     string `json:"store_shop_id"`
	StoreEmployeeId string `json:"store_employee_id"`
	StoreCustomerId string `json:"store_customer_id"`
	SessionId       string `json:"session_id"`
	ClientId        string `json:"client_id"`
}

func GinJWTMiddleware(ctx *gin.Context) {
	g := &HeaderGrantInfo{
		Mode:            ctx.GetHeader("Auth-Mode"),
		StoreId:         ctx.GetHeader("Auth-Store-Id"),
		ManagerId:       ctx.GetHeader("Auth-Manager-Id"),
		StoreShopId:     ctx.GetHeader("Auth-Store-Shop-Id"),
		StoreUserId:     ctx.GetHeader("Auth-Store-User-Id"),
		StoreEmployeeId: ctx.GetHeader("Auth-Store-Employee-Id"),
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
	if g.StoreShopId != "" {
		ctx.Keys["store_shop_id"] = g.StoreShopId
	}
	if g.StoreUserId != "" {
		ctx.Keys["store_user_id"] = g.StoreUserId
	}
	if g.StoreEmployeeId != "" {
		ctx.Keys["store_employee_id"] = g.StoreEmployeeId
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
