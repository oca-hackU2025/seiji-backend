package models

import "time"

type User struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	FirebaseUserID string    `gorm:"size:255;not null" json:"firebase_id"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
}
