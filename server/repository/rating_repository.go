package repository

import (
	"riddles-server/database"
	"riddles-server/models"
)

type RatingRepository interface {
	CreateOrUpdate(rating *models.RiddleRating) error
	Delete(userID, riddleID uint) error
	FindByUserAndRiddle(userID, riddleID uint) (*models.RiddleRating, error)
	GetRiddleRatings(riddleID uint) (int, int, error) // likes, dislikes
	GetUserRating(userID, riddleID uint) (int, error)
}

type ratingRepository struct{}

func NewRatingRepository() RatingRepository {
	return &ratingRepository{}
}

func (r *ratingRepository) CreateOrUpdate(rating *models.RiddleRating) error {
	// Check if rating already exists
	existingRating, err := r.FindByUserAndRiddle(rating.UserID, rating.RiddleID)
	if err != nil {
		// No existing rating, create new one
		return database.DB.Create(rating).Error
	}
	
	// Update existing rating
	existingRating.Rating = rating.Rating
	return database.DB.Save(existingRating).Error
}

func (r *ratingRepository) Delete(userID, riddleID uint) error {
	return database.DB.Where("user_id = ? AND riddle_id = ?", userID, riddleID).Delete(&models.RiddleRating{}).Error
}

func (r *ratingRepository) FindByUserAndRiddle(userID, riddleID uint) (*models.RiddleRating, error) {
	var rating models.RiddleRating
	err := database.DB.Where("user_id = ? AND riddle_id = ?", userID, riddleID).First(&rating).Error
	if err != nil {
		return nil, err
	}
	return &rating, nil
}

func (r *ratingRepository) GetRiddleRatings(riddleID uint) (int, int, error) {
	var likes, dislikes int64
	
	// Count likes (rating = 1)
	err := database.DB.Model(&models.RiddleRating{}).Where("riddle_id = ? AND rating = ?", riddleID, 1).Count(&likes).Error
	if err != nil {
		return 0, 0, err
	}
	
	// Count dislikes (rating = -1)
	err = database.DB.Model(&models.RiddleRating{}).Where("riddle_id = ? AND rating = ?", riddleID, -1).Count(&dislikes).Error
	if err != nil {
		return 0, 0, err
	}
	
	return int(likes), int(dislikes), nil
}

func (r *ratingRepository) GetUserRating(userID, riddleID uint) (int, error) {
	var rating models.RiddleRating
	err := database.DB.Select("rating").Where("user_id = ? AND riddle_id = ?", userID, riddleID).First(&rating).Error
	if err != nil {
		return 0, err
	}
	return rating.Rating, nil
}