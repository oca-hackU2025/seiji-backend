package controller

import (
	"net/http"

	"github.com/KENKUN-1031/seiji-backend/service"
	"github.com/gin-gonic/gin"
)

func GetRandomSeijika(c *gin.Context) {
	seijika, err := service.GetRandomSeijika()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "データ取得に失敗しました"})
		return
	}

	if seijika == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "政治家データが存在しません"})
		return
	}

	c.JSON(http.StatusOK, seijika)
}
