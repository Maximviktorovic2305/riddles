package repository

import (
	"time"
	"riddles-server/database"
	"riddles-server/models"
)

type DailyRiddleRepository interface {
	Create(dailyRiddle *models.DailyRiddle) error
	FindByDate(date time.Time) (*models.DailyRiddle, error)
	FindToday() (*models.DailyRiddle, error)
	GetRiddlesForDateRange(startDate, endDate time.Time) ([]models.DailyRiddle, error)
}

type dailyRiddleRepository struct{}

func NewDailyRiddleRepository() DailyRiddleRepository {
	return &dailyRiddleRepository{}
}

func (r *dailyRiddleRepository) Create(dailyRiddle *models.DailyRiddle) error {
	return database.DB.Create(dailyRiddle).Error
}

func (r *dailyRiddleRepository) FindByDate(date time.Time) (*models.DailyRiddle, error) {
	var dailyRiddle models.DailyRiddle
	err := database.DB.Preload("Riddle.Category").Where("featured_date = ?", date.Format("2006-01-02")).First(&dailyRiddle).Error
	return &dailyRiddle, err
}

func (r *dailyRiddleRepository) FindToday() (*models.DailyRiddle, error) {
	return r.FindByDate(time.Now())
}

func (r *dailyRiddleRepository) GetRiddlesForDateRange(startDate, endDate time.Time) ([]models.DailyRiddle, error) {
	var dailyRiddles []models.DailyRiddle
	err := database.DB.Preload("Riddle.Category").Where("featured_date BETWEEN ? AND ?", startDate.Format("2006-01-02"), endDate.Format("2006-01-02")).Find(&dailyRiddles).Error
	return dailyRiddles, err
}