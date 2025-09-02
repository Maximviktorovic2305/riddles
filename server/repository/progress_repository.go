package repository

import (
	"riddles-server/database"
	"riddles-server/models"
)

type ProgressRepository interface {
	Create(progress *models.UserRiddleProgress) error
	FindByUserAndRiddle(userID, riddleID uint) (*models.UserRiddleProgress, error)
	FindByUserID(userID uint) ([]models.UserRiddleProgress, error)
	Update(progress *models.UserRiddleProgress) error
	GetUserStats(userID uint) (int, int, error) // total, solved
}

type progressRepository struct{}

func NewProgressRepository() ProgressRepository {
	return &progressRepository{}
}

func (r *progressRepository) Create(progress *models.UserRiddleProgress) error {
	return database.DB.Create(progress).Error
}

func (r *progressRepository) FindByUserAndRiddle(userID, riddleID uint) (*models.UserRiddleProgress, error) {
	var progress models.UserRiddleProgress
	err := database.DB.Where("user_id = ? AND riddle_id = ?", userID, riddleID).First(&progress).Error
	if err != nil {
		return nil, err
	}
	return &progress, nil
}

func (r *progressRepository) FindByUserID(userID uint) ([]models.UserRiddleProgress, error) {
	var progress []models.UserRiddleProgress
	err := database.DB.Where("user_id = ?", userID).Find(&progress).Error
	return progress, err
}

func (r *progressRepository) Update(progress *models.UserRiddleProgress) error {
	return database.DB.Save(progress).Error
}

func (r *progressRepository) GetUserStats(userID uint) (int, int, error) {
	var total, solved int64
	
	// Get total riddles for user
	err := database.DB.Model(&models.UserRiddleProgress{}).Where("user_id = ?", userID).Count(&total).Error
	if err != nil {
		return 0, 0, err
	}
	
	// Get solved riddles for user
	err = database.DB.Model(&models.UserRiddleProgress{}).Where("user_id = ? AND solved = ?", userID, true).Count(&solved).Error
	if err != nil {
		return 0, 0, err
	}
	
	return int(total), int(solved), nil
}