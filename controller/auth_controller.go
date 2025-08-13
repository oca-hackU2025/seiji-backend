package controller

import (
	"fmt"
	"net/http"

	"github.com/KENKUN-1031/seiji-backend/db"
	"github.com/KENKUN-1031/seiji-backend/lib/firebase"
	"github.com/KENKUN-1031/seiji-backend/lib/jwt"
	"github.com/KENKUN-1031/seiji-backend/models"
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

	// 🔍 DBにすでに存在するか確認
	var user models.User
	result := db.DB.Where("firebase_user_id = ?", token.UID).First(&user)

	if result.Error != nil {
		// ユーザーが見つからなければ新規作成
		user = models.User{FirebaseUserID: token.UID}
		if err := db.DB.Create(&user).Error; err != nil {
			fmt.Println("❌ Failed to create user:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}
		fmt.Println("✅ New user created: ", user.ID)
	}

	// 🔐 JWT 生成（セキュリティ上、User.IDを使用）
	jwtToken, err := jwt.GenerateToken(user.ID)
	if err != nil {
		fmt.Println("❌ Failed to generate JWT:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate JWT"})
		return
	}
	fmt.Println("🔑JWToken: " + jwtToken)
	c.JSON(http.StatusOK, gin.H{
		"message":     "Login successful",
		"accessToken": jwtToken,
	})
}
