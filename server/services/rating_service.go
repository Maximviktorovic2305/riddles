package services

import (
	"errors"
	"riddles-server/models"
	"riddles-server/repository"
)

type RatingService interface {
	RateRiddle(userID, riddleID uint, rating int) error // rating: -1 for dislike, 1 for like
	RemoveRating(userID, riddleID uint) error
	GetRiddleRatings(riddleID uint) (int, int, error) // likes, dislikes
	GetUserRating(userID, riddleID uint) (int, error) // -1, 0, or 1
}

type ratingService struct {
	ratingRepo repository.RatingRepository
	riddleRepo repository.RiddleRepository
}

func NewRatingService(ratingRepo repository.RatingRepository, riddleRepo repository.RiddleRepository) RatingService {
	return &ratingService{
		ratingRepo: ratingRepo,
		riddleRepo: riddleRepo,
	}
}

func (s *ratingService) RateRiddle(userID, riddleID uint, rating int) error {
	// Validate rating value
	if rating != -1 && rating != 1 {
		return errors.New("rating must be either -1 (dislike) or 1 (like)")
	}

	// Check if riddle exists
	_, err := s.riddleRepo.FindByID(riddleID)
	if err != nil {
		return err
	}

	// Create or update rating
	riddleRating := &models.RiddleRating{
		UserID:   userID,
		RiddleID: riddleID,
		Rating:   rating,
	}

	return s.ratingRepo.CreateOrUpdate(riddleRating)
}

func (s *ratingService) RemoveRating(userID, riddleID uint) error {
	return s.ratingRepo.Delete(userID, riddleID)
}

func (s *ratingService) GetRiddleRatings(riddleID uint) (int, int, error) {
	return s.ratingRepo.GetRiddleRatings(riddleID)
}

func (s *ratingService) GetUserRating(userID, riddleID uint) (int, error) {
	return s.ratingRepo.GetUserRating(userID, riddleID)
}