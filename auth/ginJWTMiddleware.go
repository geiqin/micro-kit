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
	ctx.Keys["user_id"] = user.UserId
	ctx.Keys["nickname"] = user.Nickname
	ctx.Keys["platform_id"] = user.PlatformId
	ctx.Keys["store_id"] = user.StoreId
	ctx.Keys["realstore_id"] = user.RealstoreId
	ctx.Keys["shop_id"] = user.ShopId
	ctx.Keys["session_key"] = user.SessionKey
	ctx.Keys["client_id"] = user.ClientId
	ctx.Keys["application"] = ctx.GetHeader("Application")
	ctx.Keys["application_client_type"] = ctx.GetHeader("Application-Client-Type")
	ctx.Keys["extends"] = user.Extends
	//Pass on
	ctx.Next()

}
