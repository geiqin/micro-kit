package auth

import (
	"github.com/gin-gonic/gin"
)

func GinJWTMiddleware(ctx *gin.Context) {
	user := GetUserByGinHeader(ctx)
	if ctx.Keys == nil {
		ctx.Keys = make(map[string]interface{})
	}
	ctx.Keys["Auth-Mode"] = user.Mode
	ctx.Keys["Auth-User-Id"] = user.UserId
	ctx.Keys["Auth-Nickname"] = user.Nickname
	ctx.Keys["Auth-Platform-Id"] = user.PlatformId
	ctx.Keys["Auth-Store-Id"] = user.StoreId
	ctx.Keys["Auth-Realstore-Id"] = user.RealstoreId
	ctx.Keys["Auth-Shop-Id"] = user.ShopId
	ctx.Keys["Auth-Session-Key"] = user.SessionKey
	ctx.Keys["Auth-Client-Id"] = user.ClientId
	ctx.Keys["From-Type"] = ctx.GetHeader("From-Type")
	ctx.Keys["Routine-Type"] = ctx.GetHeader("Routine-Type")
	ctx.Keys["Auth-Extends"] = user.Extends

	//Pass on
	ctx.Next()

}
