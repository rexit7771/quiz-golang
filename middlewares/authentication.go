package middlewares

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const basicAuthprefix string = "Basic "

		auth := ctx.GetHeader("Authorization")
		fmt.Println(auth)
		if auth == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized")
			return
		}

		payload, err := base64.StdEncoding.DecodeString(auth[len(basicAuthprefix):])
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized")
			return
		}

		pair := strings.SplitN(string(payload), ":", 2)
		if len(pair) != 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized")
			return
		}

		ctx.Next()
	}
}
