package models

import (
	"time"
)

type Favorite struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `json:"user_id"`
	User      User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
	RiddleID  uint      `json:"riddle_id"`
	Riddle    Riddle    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"riddle"`
	CreatedAt time.Time `json:"created_at"`
}