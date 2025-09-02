package repository

import (
	"riddles-server/database"
	"riddles-server/models"
)

type FavoriteRepository interface {
	Create(favorite *models.Favorite) error
	Delete(userID, riddleID uint) error
	FindByUserAndRiddle(userID, riddleID uint) (*models.Favorite, error)
	FindByUserID(userID uint) ([]models.Favorite, error)
	IsFavorite(userID, riddleID uint) bool
}

type favoriteRepository struct{}

func NewFavoriteRepository() FavoriteRepository {
	return &favoriteRepository{}
}

func (r *favoriteRepository) Create(favorite *models.Favorite) error {
	return database.DB.Create(favorite).Error
}

func (r *favoriteRepository) Delete(userID, riddleID uint) error {
	return database.DB.Where("user_id = ? AND riddle_id = ?", userID, riddleID).Delete(&models.Favorite{}).Error
}

func (r *favoriteRepository) FindByUserAndRiddle(userID, riddleID uint) (*models.Favorite, error) {
	var favorite models.Favorite
	err := database.DB.Where("user_id = ? AND riddle_id = ?", userID, riddleID).First(&favorite).Error
	return &favorite, err
}

func (r *favoriteRepository) FindByUserID(userID uint) ([]models.Favorite, error) {
	var favorites []models.Favorite
	err := database.DB.Where("user_id = ?", userID).Find(&favorites).Error
	return favorites, err
}

func (r *favoriteRepository) IsFavorite(userID, riddleID uint) bool {
	var count int64
	database.DB.Model(&models.Favorite{}).Where("user_id = ? AND riddle_id = ?", userID, riddleID).Count(&count)
	return count > 0
}