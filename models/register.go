package models

// migrateしたいモデルをどんどん追加していく
func AllModels() []interface{} {
	return []interface{}{
		&User{},
		&Seijika{},
		&Career{},
		&Sns{},
		&Like{},
		// &Post{}, ←例
	}
}
