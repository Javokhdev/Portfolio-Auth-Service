package middleware

import (
	"net/http"
	"strings"

	t "github.com/Javokhdev/Auth-Service/api/token"
	"github.com/gin-gonic/gin"
)

func MiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authourization")
		url := ctx.Request.URL.Path
		if strings.Contains(url, "swagger") || url == "/user/login" || url == "/user/registr"{
			ctx.Next()
			return
		} else if _, err := t.ExtractClaim(token); err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.Next()

	}
}
