package models

// migrateしたいモデルをどんどん追加していく
func AllModels() []interface{} {
	return []interface{}{
		&User{},
		// &Post{}, ←例
	}
}
