package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/KENKUN-1031/seiji-backend/lib/firebase"
	"github.com/KENKUN-1031/seiji-backend/lib/jwt"
	"github.com/gin-gonic/gin"
)

func FirebaseAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		idToken := c.GetHeader("Authorization")
		if idToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No ID token provided"})
			return
		}

		// ✅ 共通関数でトークンを検証
		token, err := firebase.VerifyIDToken(idToken)
		if err != nil {
			fmt.Println("❌ Token verification failed:", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid ID token"})
			return
		}

		// UID を Gin コンテキストに保存
		c.Set("uid", token.UID)
		c.Next() //controllerに処理をつなげるため
	}
}

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		uid, err := jwt.VerifyTokenAndExtractUID(tokenStr) //jwt認証してuidを出力してる
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid JWT"})
			return
		}

		c.Set("uid", uid)
		c.Next()
	}
}
