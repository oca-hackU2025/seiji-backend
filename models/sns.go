package models

import (
	"github.com/lib/pq"
)

type Sns struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	SeijikaID  uint           `gorm:"not null" json:"seijika_id"`
	Twitter    string         `gorm:"size:255" json:"twitter"`
	Instagram  string         `gorm:"size:255" json:"instagram"`
	Youtube    string         `gorm:"size:255" json:"youtube"`
	Activities pq.StringArray `gorm:"type:text[]" json:"activities"`
}
