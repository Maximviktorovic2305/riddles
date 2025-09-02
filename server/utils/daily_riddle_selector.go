package utils

import (
	"math/rand"
	"time"
	"riddles-server/database"
	"riddles-server/models"
)

// SelectDailyRiddles selects 6 riddles for the current date
func SelectDailyRiddles() error {
	// Get today's date
	today := time.Now().Truncate(24 * time.Hour)
	
	// Check if we already have daily riddles for today
	var count int64
	database.DB.Model(&models.DailyRiddle{}).Where("featured_date = ?", today).Count(&count)
	
	// If we already have daily riddles for today, don't select new ones
	if count > 0 {
		return nil
	}
	
	// Get all riddles
	var riddles []models.Riddle
	if err := database.DB.Find(&riddles).Error; err != nil {
		return err
	}
	
	// If we don't have enough riddles, return error
	if len(riddles) < 6 {
		return nil // Not enough riddles, skip for now
	}
	
	// Shuffle riddles and select first 6
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(riddles), func(i, j int) {
		riddles[i], riddles[j] = riddles[j], riddles[i]
	})
	
	// Create daily riddle entries
	for i := 0; i < 6 && i < len(riddles); i++ {
		dailyRiddle := models.DailyRiddle{
			RiddleID:     riddles[i].ID,
			FeaturedDate: today,
		}
		
		if err := database.DB.Create(&dailyRiddle).Error; err != nil {
			return err
		}
	}
	
	return nil
}

// GetTodaysRiddles returns the riddles selected for today
func GetTodaysRiddles() ([]models.DailyRiddle, error) {
	today := time.Now().Truncate(24 * time.Hour)
	
	var dailyRiddles []models.DailyRiddle
	err := database.DB.Preload("Riddle.Category").Where("featured_date = ?", today).Find(&dailyRiddles).Error
	
	return dailyRiddles, err
}