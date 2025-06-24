package models

type Career struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	SeijikaID uint   `gorm:"not null" json:"seijika_id"`
	Year      string `gorm:"size:255" json:"year"`
	Desc      string `gorm:"size:255" json:"description"`
	IconType  int    `json:"icon_type"`
}
