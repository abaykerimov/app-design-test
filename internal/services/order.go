package services

import (
	"applicationDesignTest/internal/domain/order"
	"fmt"
	"time"
)

type Repository interface {
	CreateOrder(newOrder *order.Order) *order.Order
	GetOrderAvailabilities() []*order.RoomAvailability
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateOrder(entity *order.Order) (*order.Order, error) {
	daysToBook := daysBetween(entity.From, entity.To)

	unavailableDays := make(map[time.Time]struct{})
	for _, day := range daysToBook {
		unavailableDays[day] = struct{}{}
	}

	availabilities := s.repo.GetOrderAvailabilities()

	for _, dayToBook := range daysToBook {
		for i, availability := range availabilities {
			if !availability.Date.Equal(dayToBook) || availability.Quota < 1 {
				continue
			}
			availability.Quota -= 1
			availabilities[i] = availability
			delete(unavailableDays, dayToBook)
		}
	}

	if len(unavailableDays) != 0 {
		return nil, fmt.Errorf("Hotel room is not available for selected dates:\n%v\n%v", entity, unavailableDays)
	}

	return s.repo.CreateOrder(entity), nil
}

func daysBetween(from time.Time, to time.Time) []time.Time {
	if from.After(to) {
		return nil
	}

	days := make([]time.Time, 0)
	for d := toDay(from); !d.After(toDay(to)); d = d.AddDate(0, 0, 1) {
		days = append(days, d)
	}

	return days
}

func toDay(timestamp time.Time) time.Time {
	return time.Date(timestamp.Year(), timestamp.Month(), timestamp.Day(), 0, 0, 0, 0, time.UTC)
}
