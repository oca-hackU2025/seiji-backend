package db

import (
	"fmt"
	"log"
)

func Init() {
	// dbとの接続
	if err := Connect(); err != nil {
		log.Fatalf("❌ DB接続失敗: %v", err)
	}
	// migrationしてる
	if err := Migrate(); err != nil {
		log.Fatalf("❌ マイグレーション失敗: %v", err)
	}
	// seedでダミーデータ入れてる
	// if err := Seed(); err != nil {
	// 	log.Fatalf("❌ Seed失敗: %v", err)
	// }
	fmt.Println("✅ DB Initializer Successful")
}
