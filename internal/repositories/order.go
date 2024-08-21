package repositories

import (
	"applicationDesignTest/internal/domain/order"
	"time"
)

type OrderRepository struct {
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{}
}

func (repo *OrderRepository) CreateOrder(newOrder *order.Order) *order.Order {
	var orders = make([]*order.Order, 0)
	orders = append(orders, newOrder)

	return newOrder
}

func (repo *OrderRepository) GetOrderAvailabilities() []*order.RoomAvailability {
	return []*order.RoomAvailability{
		{"reddison", "lux", date(2024, 1, 1), 1},
		{"reddison", "lux", date(2024, 1, 2), 1},
		{"reddison", "lux", date(2024, 1, 3), 1},
		{"reddison", "lux", date(2024, 1, 4), 1},
		{"reddison", "lux", date(2024, 1, 5), 0},
	}
}

func date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
