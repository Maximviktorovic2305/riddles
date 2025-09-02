package services

import (
	"strings"
	"riddles-server/models"
	"riddles-server/repository"
)

type RiddleService interface {
	GetAllRiddles() ([]models.Riddle, error)
	GetRiddleByID(id uint) (*models.Riddle, error)
	GetRiddlesByCategory(categoryID uint) ([]models.Riddle, error)
	GetRiddlesByDifficulty(difficulty string) ([]models.Riddle, error)
	GetRiddlesByCategoryAndDifficulty(categoryID uint, difficulty string) ([]models.Riddle, error)
	SearchRiddles(query string) ([]models.Riddle, error)
	CheckAnswer(riddleID uint, userAnswer string) (bool, error)
	GetRiddleWithUserProgress(riddleID, userID uint) (*RiddleWithProgress, error)
	GetRiddlesWithUserProgress(riddles []models.Riddle, userID uint) ([]RiddleWithProgress, error)
}

type RiddleWithProgress struct {
	Riddle       models.Riddle `json:"riddle"`
	IsSolved     bool          `json:"is_solved"`
	IsFavorite   bool          `json:"is_favorite"`
	UserRating   int           `json:"user_rating"` // -1, 0, or 1
	Likes        int           `json:"likes"`
	Dislikes     int           `json:"dislikes"`
}

type riddleService struct {
	riddleRepo   repository.RiddleRepository
	progressRepo repository.ProgressRepository
	favoriteRepo repository.FavoriteRepository
	ratingRepo   repository.RatingRepository
}

func NewRiddleService(
	riddleRepo repository.RiddleRepository,
	progressRepo repository.ProgressRepository,
	favoriteRepo repository.FavoriteRepository,
	ratingRepo repository.RatingRepository,
) RiddleService {
	return &riddleService{
		riddleRepo:   riddleRepo,
		progressRepo: progressRepo,
		favoriteRepo: favoriteRepo,
		ratingRepo:   ratingRepo,
	}
}

func (s *riddleService) GetAllRiddles() ([]models.Riddle, error) {
	return s.riddleRepo.FindAll()
}

func (s *riddleService) GetRiddleByID(id uint) (*models.Riddle, error) {
	return s.riddleRepo.FindByID(id)
}

func (s *riddleService) GetRiddlesByCategory(categoryID uint) ([]models.Riddle, error) {
	return s.riddleRepo.FindByCategory(categoryID)
}

func (s *riddleService) GetRiddlesByDifficulty(difficulty string) ([]models.Riddle, error) {
	return s.riddleRepo.FindByDifficulty(difficulty)
}

func (s *riddleService) GetRiddlesByCategoryAndDifficulty(categoryID uint, difficulty string) ([]models.Riddle, error) {
	return s.riddleRepo.FindByCategoryAndDifficulty(categoryID, difficulty)
}

func (s *riddleService) SearchRiddles(query string) ([]models.Riddle, error) {
	return s.riddleRepo.Search(query)
}

func (s *riddleService) CheckAnswer(riddleID uint, userAnswer string) (bool, error) {
	riddle, err := s.riddleRepo.FindByID(riddleID)
	if err != nil {
		return false, err
	}

	// Case-insensitive comparison
	return strings.EqualFold(strings.TrimSpace(userAnswer), strings.TrimSpace(riddle.Answer)), nil
}

func (s *riddleService) GetRiddleWithUserProgress(riddleID, userID uint) (*RiddleWithProgress, error) {
	riddle, err := s.riddleRepo.FindByID(riddleID)
	if err != nil {
		return nil, err
	}

	result := &RiddleWithProgress{
		Riddle: *riddle,
	}

	// Get user progress
	progress, err := s.progressRepo.FindByUserAndRiddle(userID, riddleID)
	if err == nil {
		result.IsSolved = progress.Solved
	}

	// Check if favorite
	result.IsFavorite = s.favoriteRepo.IsFavorite(userID, riddleID)

	// Get user rating
	userRating, err := s.ratingRepo.GetUserRating(userID, riddleID)
	if err == nil {
		result.UserRating = userRating
	}

	// Get total ratings
	likes, dislikes, err := s.ratingRepo.GetRiddleRatings(riddleID)
	if err == nil {
		result.Likes = likes
		result.Dislikes = dislikes
	}

	return result, nil
}

func (s *riddleService) GetRiddlesWithUserProgress(riddles []models.Riddle, userID uint) ([]RiddleWithProgress, error) {
	result := make([]RiddleWithProgress, len(riddles))

	// Get all user progress for these riddles
	var riddleIDs []uint
	for _, riddle := range riddles {
		riddleIDs = append(riddleIDs, riddle.ID)
	}

	// This would be more efficient with a batch query, but for simplicity we'll do individual lookups
	for i, riddle := range riddles {
		result[i] = RiddleWithProgress{
			Riddle: riddle,
		}

		// Get user progress
		progress, err := s.progressRepo.FindByUserAndRiddle(userID, riddle.ID)
		if err == nil {
			result[i].IsSolved = progress.Solved
		}

		// Check if favorite
		result[i].IsFavorite = s.favoriteRepo.IsFavorite(userID, riddle.ID)

		// Get user rating
		userRating, err := s.ratingRepo.GetUserRating(userID, riddle.ID)
		if err == nil {
			result[i].UserRating = userRating
		}

		// Get total ratings
		likes, dislikes, err := s.ratingRepo.GetRiddleRatings(riddle.ID)
		if err == nil {
			result[i].Likes = likes
			result[i].Dislikes = dislikes
		}
	}

	return result, nil
}