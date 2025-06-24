package models

type Like struct {
	ID        uint `gorm:"primaryKey" json:"id"`
	UserID    uint `gorm:"not null;uniqueIndex:idx_user_seijika"`
	SeijikaID uint `gorm:"not null;uniqueIndex:idx_user_seijika"`

	// ↓ リレーション（必要に応じて）
	User    *User    `gorm:"foreignKey:UserID" json:"-"`
	Seijika *Seijika `gorm:"foreignKey:SeijikaID" json:"-"`
}
