package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	dsn := os.Getenv("DATABASE_URL") //←デプロイする時はDB_URLをここに入れる
	if dsn == "" {
		// 環境変数が設定されていない場合、デフォルトのDSNを使用
		dsn = "host=db user=user password=password dbname=mydatabase port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db
	fmt.Println("Database connection established")
	return nil
}
