package controller

import (
	"fmt"
	"net/http"

	"github.com/KENKUN-1031/seiji-backend/db"
	"github.com/KENKUN-1031/seiji-backend/utils"

	"github.com/KENKUN-1031/seiji-backend/models"
	"github.com/gin-gonic/gin"
)

type LikeRequest struct {
	SeijikaID uint `json:"seijika_id"`
}

func CreateLike(c *gin.Context) {
	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
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

func GetLikedSeijikaList(c *gin.Context) {
	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("✅ userID from context:", userID)

	// ユーザーが Like している Seijika 一覧を取得（JOIN して取得）
	var likes []models.Like
	if err := db.DB.Preload("Seijika").Where("user_id = ?", userID).Find(&likes).Error; err != nil {
		fmt.Println("❌ Failed to fetch liked seijika:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch liked seijika"})
		return
	}

	// Seijika のみを抽出して返す（不要なデータ除去のため）
	var likedSeijikas []models.Seijika
	for _, like := range likes {
		if like.Seijika != nil {
			likedSeijikas = append(likedSeijikas, *like.Seijika)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"liked_seijika": likedSeijikas,
	})
}
