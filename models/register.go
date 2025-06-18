package models

// migrateしたいモデルをどんどん追加していく
func AllModels() []interface{} {
	return []interface{}{
		&User{},
		&Seijika{},
		&Career{},
		&Sns{},
		// &Post{}, ←例
	}
}
