package models

import (
	"time"
)

type RiddleRating struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `json:"user_id"`
	User      User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
	RiddleID  uint      `json:"riddle_id"`
	Riddle    Riddle    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"riddle"`
	Rating    int       `gorm:"check:rating IN (-1, 1)" json:"rating"` // -1 for dislike, 1 for like
	CreatedAt time.Time `json:"created_at"`
}