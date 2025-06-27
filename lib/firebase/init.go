package firebase

import (
	"context"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

var App *firebase.App

func InitFirebase() {
	credPath := os.Getenv("FIREBASE_CREDENTIALS")
	if credPath == "" {
		fmt.Println("❌ 環境変数 FIREBASE_CREDENTIALS が見つかりません")
		log.Fatal("❌ FIREBASE_CREDENTIALS が設定されていません")
	}

	// fmt.Println("✅ Firebase認証ファイルパス:", credPath)

	opt := option.WithCredentialsFile(credPath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		fmt.Printf("❌ Firebase 初期化エラー: %v\n", err)
		log.Fatalf("❌ Firebase 初期化エラー: %v", err)
	}

	fmt.Println("✅ Firebase App 初期化に成功しました")
	App = app
}
