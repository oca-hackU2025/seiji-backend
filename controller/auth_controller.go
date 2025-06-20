package controller

import (
	"fmt"
	"net/http"

	"github.com/KENKUN-1031/seiji-backend/lib/firebase"
	"github.com/KENKUN-1031/seiji-backend/lib/jwt"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	SessionID string `json:"idToken"`
}

func Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("❌ Failed to parse JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// fmt.Println("✅ Received session ID:", req.SessionID)

	// ✅ 共通ライブラリの関数を使って Firebase トークン検証
	token, err := firebase.VerifyIDToken(req.SessionID)
	if err != nil {
		fmt.Println("❌ Token verification failed:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Firebase ID token"})
		return
	}

	fmt.Println("✅ Firebase Auth success: UID =", token.UID)

	// 🔐 JWT 生成
	jwtToken, err := jwt.GenerateToken(token.UID)
	if err != nil {
		fmt.Println("❌ Failed to generate JWT:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate JWT"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Login successful",
		"accessToken": jwtToken,
	})
}
