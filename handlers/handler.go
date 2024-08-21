package handlers

import (
	"applicationDesignTest/internal/domain/order"
	"applicationDesignTest/packages/log"
	"encoding/json"
	"net/http"
)

type Service interface {
	CreateOrder(newOrder *order.Order) (*order.Order, error)
}

type Handler struct {
	logger  *log.Logger
	service Service
}

func NewHandler(service Service, logger *log.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}

func (h *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	request, err := getOrderRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	entity, err := request.getOrder()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := h.service.CreateOrder(entity)
	if err != nil {
		h.logger.LogErrorf("Error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

	h.logger.LogInfoF("Order successfully created: %v", response)
}

func (h *Handler) LoadRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /orders", h.CreateOrder)
}
