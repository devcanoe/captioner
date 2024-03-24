package middleware

import (
	"strings"
	"time"

	"captioner.com.ng/api/server/utils"
	"captioner.com.ng/api/types"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func IsAuthenticated() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session_token := ctx.GetHeader("Authorization")
		refresh_token := ctx.GetHeader("x-refresh")
		if refresh_token == "" && session_token == "" {
			ctx.AbortWithStatus(404)
			return
		}
		refreshToken, err := utils.VerifyToken(refresh_token)
		if err != nil {
			ctx.AbortWithStatus(404)
			return
		}
		if refreshToken.Expiration().Unix() < time.Now().Unix() {
			ctx.AbortWithStatus(404)
			return
		}

		id, found := refreshToken.Get("user_id")
		if !found {
			ctx.AbortWithStatus(400)
			return
		}

		userID, _ := primitive.ObjectIDFromHex(id.(string))

		authToken, err := utils.VerifyToken(strings.Split(session_token, " ")[1])
		if err != nil {
			ctx.AbortWithStatus(404)
			return
		}
		if authToken.Expiration().Unix() < time.Now().Unix() {
			payload := utils.JWTPayload{
				UserID: userID,
				Type:   "",
			}
			newSession, err := utils.CreateToken(payload, time.Now().Add(300*time.Second))
			if err != nil {
				ctx.AbortWithError(400, err)
				return
			}
			ctx.SetCookie("x-session", newSession, types.SESSION_TOKEN_EXPIRE, "*", "*", true, true)
		}

		ctx.Next()
	}

}
