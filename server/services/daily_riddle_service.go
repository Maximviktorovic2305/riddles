package services

import (
	"time"
	"riddles-server/models"
	"riddles-server/repository"
)

type DailyRiddleService interface {
	GetTodayRiddle() (*models.DailyRiddle, error)
	GetRiddleByDate(date time.Time) (*models.DailyRiddle, error)
	GetRiddlesForDateRange(startDate, endDate time.Time) ([]models.DailyRiddle, error)
}

type dailyRiddleService struct {
	dailyRiddleRepo repository.DailyRiddleRepository
}

func NewDailyRiddleService(dailyRiddleRepo repository.DailyRiddleRepository) DailyRiddleService {
	return &dailyRiddleService{
		dailyRiddleRepo: dailyRiddleRepo,
	}
}

func (s *dailyRiddleService) GetTodayRiddle() (*models.DailyRiddle, error) {
	return s.dailyRiddleRepo.FindToday()
}

func (s *dailyRiddleService) GetRiddleByDate(date time.Time) (*models.DailyRiddle, error) {
	return s.dailyRiddleRepo.FindByDate(date)
}

func (s *dailyRiddleService) GetRiddlesForDateRange(startDate, endDate time.Time) ([]models.DailyRiddle, error) {
	return s.dailyRiddleRepo.GetRiddlesForDateRange(startDate, endDate)
}