package db

import (
	"fmt"

	"github.com/KENKUN-1031/seiji-backend/models" // モデルのパスをプロジェクトに合わせて調整
)

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

func ResetAndMigrate() error {
	if DB == nil {
		return fmt.Errorf("DB connection is not initialized")
	}

	// Drop all tables
	err := DB.Migrator().DropTable(models.AllModels()...)
	if err != nil {
		return fmt.Errorf("failed to drop tables: %w", err)
	}
	fmt.Println("🧹 All tables dropped.")

	// Recreate tables
	err = DB.AutoMigrate(models.AllModels()...)
	if err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}
	fmt.Println("✅ Database migration completed successfully")

	return nil
}
