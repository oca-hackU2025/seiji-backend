package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KENKUN-1031/seiji-backend/lib/firebase"
	"github.com/gin-gonic/gin"
)

func FirebaseAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		idToken := c.GetHeader("Authorization")
		if idToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No ID token provided"})
			return
		}

		client, err := firebase.App.Auth(context.Background())
		if err != nil {
			fmt.Println("❌ Firebase client error:", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Firebase client error"})
			return
		}

		token, err := client.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			fmt.Println("❌ Token verification failed:", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid ID token"})
			return
		}

		// UIDをコンテキストに保存して、次の処理へ
		c.Set("uid", token.UID)
		c.Next()
	}
}
