package db

import (
	"fmt"

	"github.com/KENKUN-1031/seiji-backend/models" // モデルのパスをプロジェクトに合わせて調整
)

// Migrate runs the database schema migrations
func Migrate() error {
	if DB == nil {
		return fmt.Errorf("DB connection is not initialized")
	}

	// ↓ models/register.goに入ってるmodelを全てここでmigrateしてる
	err := DB.AutoMigrate(models.AllModels()...)

	if err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}

	fmt.Println("✅ Database migration completed successfully")
	return nil
}
