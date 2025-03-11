package middlewares

import (
	"github.com/gin-gonic/gin"
	"shop-api/user-web/models"
)

func IsAdminAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, _ := ctx.Get("claims")
		currentUser, _ := claims.(*models.CustomClaims)

		if currentUser.AuthorityId != 2 {
			ctx.JSON(200, gin.H{
				"msg": "无权限",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
