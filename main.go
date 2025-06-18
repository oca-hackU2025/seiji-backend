package main

import (
	"fmt"
	"log"
	"time"

	"github.com/KENKUN-1031/seiji-backend/db"
	"github.com/KENKUN-1031/seiji-backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	//dbと接続してる
	if err := db.Connect(); err != nil {
		log.Fatalf("DB接続失敗: %v", err)
	}
	//migrateしてる
	if err := db.Migrate(); err != nil {
		log.Fatalf("マイグレーション失敗: %v", err)
	}
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

	router.Run() // ← ここも router にする
}
