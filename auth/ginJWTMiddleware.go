package auth

import (
	"github.com/gin-gonic/gin"
)

func GinJWTMiddleware(ctx *gin.Context) {
	user := GetUserByGinHeader(ctx)
	if ctx.Keys == nil {
		ctx.Keys = make(map[string]interface{})
	}
	ctx.Keys["mode"] = user.Mode
	ctx.Keys["display_name"] = user.DisplayName
	ctx.Keys["user_id"] = user.UserId
	ctx.Keys["store_id"] = user.StoreId
	ctx.Keys["store_shop_id"] = user.StoreShopId
	ctx.Keys["session_key"] = user.SessionKey
	ctx.Keys["client_id"] = user.ClientId
	//Pass on
	ctx.Next()

}
