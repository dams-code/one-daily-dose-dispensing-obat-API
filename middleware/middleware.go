package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func MiddlewareJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status_auth_login": "error",
				"message":           "Akses gagal, tidak ada token",
			})
			return
		}

		Token := strings.TrimPrefix(authHeader, "Bearer ")

		setClaims, err := ValidasiJWT(Token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status_auth_login": "error",
				"message":           err.Error(),
			})

			return
		}

		ctx.Set("username_sekarang", setClaims.Username)
		ctx.Set("IdUser_sekarang", setClaims.UserID)
		ctx.Set("RoleUser", setClaims.Role)

		ctx.Next()
	}
}

func RoleUserObat(role string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		UserObatRole := ctx.MustGet("RoleUser").(string)

		if UserObatRole == role {
			ctx.Next()
			return
		}

		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"status_auth_login": "error",
			"message":           "Anda tidak memiliki akses",
		})
	}
}
