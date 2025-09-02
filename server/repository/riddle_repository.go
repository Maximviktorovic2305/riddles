package repository

import (
	"riddles-server/database"
	"riddles-server/models"
)

type RiddleRepository interface {
	FindAll() ([]models.Riddle, error)
	FindByID(id uint) (*models.Riddle, error)
	FindByCategory(categoryID uint) ([]models.Riddle, error)
	FindByDifficulty(difficulty string) ([]models.Riddle, error)
	FindByCategoryAndDifficulty(categoryID uint, difficulty string) ([]models.Riddle, error)
	Search(query string) ([]models.Riddle, error)
}

type riddleRepository struct{}

func NewRiddleRepository() RiddleRepository {
	return &riddleRepository{}
}

func (r *riddleRepository) FindAll() ([]models.Riddle, error) {
	var riddles []models.Riddle
	err := database.DB.Preload("Category").Find(&riddles).Error
	return riddles, err
}

func (r *riddleRepository) FindByID(id uint) (*models.Riddle, error) {
	var riddle models.Riddle
	err := database.DB.Preload("Category").First(&riddle, id).Error
	return &riddle, err
}

func (r *riddleRepository) FindByCategory(categoryID uint) ([]models.Riddle, error) {
	var riddles []models.Riddle
	err := database.DB.Preload("Category").Where("category_id = ?", categoryID).Find(&riddles).Error
	return riddles, err
}

func (r *riddleRepository) FindByDifficulty(difficulty string) ([]models.Riddle, error) {
	var riddles []models.Riddle
	err := database.DB.Preload("Category").Where("difficulty = ?", difficulty).Find(&riddles).Error
	return riddles, err
}

func (r *riddleRepository) FindByCategoryAndDifficulty(categoryID uint, difficulty string) ([]models.Riddle, error) {
	var riddles []models.Riddle
	err := database.DB.Preload("Category").Where("category_id = ? AND difficulty = ?", categoryID, difficulty).Find(&riddles).Error
	return riddles, err
}

func (r *riddleRepository) Search(query string) ([]models.Riddle, error) {
	var riddles []models.Riddle
	err := database.DB.Preload("Category").Where("title ILIKE ? OR description ILIKE ?", "%"+query+"%", "%"+query+"%").Find(&riddles).Error
	return riddles, err
}