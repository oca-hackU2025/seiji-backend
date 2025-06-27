package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/KENKUN-1031/seiji-backend/db"
	"github.com/KENKUN-1031/seiji-backend/lib/firebase"
	"github.com/KENKUN-1031/seiji-backend/routes"
	"github.com/joho/godotenv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Print("🔥" + os.Getenv("GIN_MODE"))
	if os.Getenv("GIN_MODE") != "release" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("❌ .envファイルの読み込みに失敗しました")
		}
	}

	// firebaseの初期化
	firebase.InitFirebase()

	// dbの初期化
	db.Init()

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	routes.DefineRoutes(router)

	for _, route := range router.Routes() {
		fmt.Printf("Method: %s, Path: %s\n", route.Method, route.Path)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // ローカル開発用のデフォルト
	}
	log.Printf("✅ Starting server on port %s\n", port)
	router.Run("0.0.0.0:" + port)
}
