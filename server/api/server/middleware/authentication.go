package middleware

import (
	"fmt"
	"strings"

	"captioner.com.ng/api/server/utils"
	"github.com/gin-gonic/gin"
)

func IsAuthenticated() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session_token := ctx.GetHeader("Authorization")
		refresh_token := ctx.GetHeader("x-refresh")
		if refresh_token == "" && session_token == "" {
			ctx.AbortWithStatus(404)
			return
		}

		authToken := strings.Split(session_token, " ")[1]

		fmt.Println(authToken)

		token, err := utils.VerifyToken(authToken)
		if err != nil {
			ctx.AbortWithStatus(404)
			return
		}

		fmt.Println(token)

		// ctx.Next()
		ctx.Next()
	}

}
