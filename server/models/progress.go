package models

import (
	"time"
)

type UserRiddleProgress struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `json:"user_id"`
	User      User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
	RiddleID  uint      `json:"riddle_id"`
	Riddle    Riddle    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"riddle"`
	Solved    bool      `gorm:"default:false" json:"solved"`
	SolvedAt  time.Time `json:"solved_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}