package handlers

import (
	"applicationDesignTest/internal/domain/order"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type orderRequest struct {
	HotelID   string `json:"hotel_id"`
	RoomID    string `json:"room_id"`
	UserEmail string `json:"email"`
	From      string `json:"from"`
	To        string `json:"to"`
}

func getOrderRequest(req *http.Request) (r *orderRequest, err error) {
	err = json.NewDecoder(req.Body).Decode(&r)
	if err != nil {
		return
	}
	return
}

func (r *orderRequest) getOrder() (*order.Order, error) {
	err := r.validate()
	if err != nil {
		return nil, err
	}

	entity := &order.Order{
		HotelID:   r.HotelID,
		RoomID:    r.RoomID,
		UserEmail: r.UserEmail,
	}

	if entity.From, err = time.Parse("2006-01-02", r.From); err != nil {
		return nil, err
	}

	if entity.To, err = time.Parse("2006-01-02", r.To); err != nil {
		return nil, err
	}

	return entity, nil
}

func (r *orderRequest) validate() error {
	if r.HotelID == "" {
		return errors.New("Отель не указан")
	}

	if r.RoomID == "" {
		return errors.New("Номер не указан")
	}

	if r.From == "" {
		return errors.New("Дата старта не указана")
	}

	if r.To == "" {
		return errors.New("Дата конца не указана")
	}

	return nil
}
