package models

import (
	"time"
)

type DailyRiddle struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	RiddleID     uint      `json:"riddle_id"`
	Riddle       Riddle    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"riddle"`
	FeaturedDate time.Time `json:"featured_date"`
	CreatedAt    time.Time `json:"created_at"`
}