package models

import "github.com/lib/pq"

type Seijika struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	Name         string         `gorm:"size:255" json:"name"`
	NameFurigana string         `gorm:"size:255" json:"name_furigana"`
	Age          int            `json:"age"`
	PartyName    string         `gorm:"size:255" json:"party_name"`
	Term         int            `json:"term"`
	CatchPhrase  string         `gorm:"size:255" json:"catch_phrase"`
	MainImg      string         `gorm:"size:255" json:"main_img"`
	GraphImg     string         `gorm:"size:255" json:"graph_img"`
	Activities   pq.StringArray `gorm:"type:text[]" json:"activities"`

	Sns    Sns    `gorm:"foreignKey:SeijikaID" json:"sns"`    //seijikaからSnsをpreloadで呼べる (逆は無理)
	Career Career `gorm:"foreignKey:SeijikaID" json:"career"` //seijikaからCareerをpreloadで呼べる (逆は無理)
}
