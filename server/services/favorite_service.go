package services

import (
	"riddles-server/models"
	"riddles-server/repository"
)

type FavoriteService interface {
	AddFavorite(userID, riddleID uint) error
	RemoveFavorite(userID, riddleID uint) error
	GetUserFavorites(userID uint) ([]models.Favorite, error)
	IsFavorite(userID, riddleID uint) bool
}

type favoriteService struct {
	favoriteRepo repository.FavoriteRepository
	riddleRepo   repository.RiddleRepository
}

func NewFavoriteService(favoriteRepo repository.FavoriteRepository, riddleRepo repository.RiddleRepository) FavoriteService {
	return &favoriteService{
		favoriteRepo: favoriteRepo,
		riddleRepo:   riddleRepo,
	}
}

func (s *favoriteService) AddFavorite(userID, riddleID uint) error {
	// Check if riddle exists
	_, err := s.riddleRepo.FindByID(riddleID)
	if err != nil {
		return err
	}

	// Create favorite record
	favorite := &models.Favorite{
		UserID:   userID,
		RiddleID: riddleID,
	}

	return s.favoriteRepo.Create(favorite)
}

func (s *favoriteService) RemoveFavorite(userID, riddleID uint) error {
	return s.favoriteRepo.Delete(userID, riddleID)
}

func (s *favoriteService) GetUserFavorites(userID uint) ([]models.Favorite, error) {
	return s.favoriteRepo.FindByUserID(userID)
}

func (s *favoriteService) IsFavorite(userID, riddleID uint) bool {
	return s.favoriteRepo.IsFavorite(userID, riddleID)
}