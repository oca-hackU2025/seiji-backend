package service

import (
	"math/rand"
	"time"

	"github.com/KENKUN-1031/seiji-backend/db"
	"github.com/KENKUN-1031/seiji-backend/models"
)

func GetRandomSeijika() (*models.Seijika, error) {
	var seijikas []models.Seijika

	// 全ての政治家データを取得
	if err := db.DB.Find(&seijikas).Error; err != nil {
		return nil, err
	}

	// データが存在しない場合
	if len(seijikas) == 0 {
		return nil, nil
	}

	// ランダムに1件選ぶ
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := r.Intn(len(seijikas))

	return &seijikas[index], nil
}
