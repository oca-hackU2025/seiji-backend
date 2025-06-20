package controller

import (
	"fmt"
	"net/http"

	"github.com/KENKUN-1031/seiji-backend/db"
	"github.com/KENKUN-1031/seiji-backend/models"
	"github.com/gin-gonic/gin"
)

type LikeRequest struct {
	SeijikaID uint `json:"seijika_id"`
}

func CreateLike(c *gin.Context) {
	// JWTミドルウェアで保存された userID を取得
	userIDInterface, exists := c.Get("userID")
	if !exists {
		fmt.Println("❌ userID not found in context")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}
	userID := userIDInterface.(uint) // 型アサーション
	fmt.Println("✅ userID from context:", userID)

	// リクエストから seijikaID を取得
	var req LikeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("❌ Failed to bind JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	fmt.Println("✅ seijikaID from request:", req.SeijikaID)

	// 重複チェック
	var existing models.Like
	if err := db.DB.Where("user_id = ? AND seijika_id = ?", userID, req.SeijikaID).First(&existing).Error; err == nil {
		fmt.Println("⚠️ Like already exists for userID:", userID, "seijikaID:", req.SeijikaID)
		c.JSON(http.StatusConflict, gin.H{"error": "Like already exists"})
		return
	}

	// Like作成
	like := models.Like{
		UserID:    userID,
		SeijikaID: req.SeijikaID,
	}
	if err := db.DB.Create(&like).Error; err != nil {
		fmt.Println("❌ Failed to create Like:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create like"})
		return
	}

	fmt.Println("✅ Like created successfully:", like)
	c.JSON(http.StatusOK, gin.H{"message": "Like created successfully"})
}
