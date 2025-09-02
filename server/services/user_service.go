package services

import (
	"riddles-server/models"
	"riddles-server/repository"
)

type UserService interface {
	GetProfile(userID uint) (*models.User, error)
	GetUserStats(userID uint) (int, int, error) // total riddles, solved riddles
}

type userService struct {
	userRepo     repository.UserRepository
	progressRepo repository.ProgressRepository
}

func NewUserService(userRepo repository.UserRepository, progressRepo repository.ProgressRepository) UserService {
	return &userService{
		userRepo:     userRepo,
		progressRepo: progressRepo,
	}
}

func (s *userService) GetProfile(userID uint) (*models.User, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}
	
	// Remove password from response
	user.Password = ""
	return user, nil
}

func (s *userService) GetUserStats(userID uint) (int, int, error) {
	return s.progressRepo.GetUserStats(userID)
}