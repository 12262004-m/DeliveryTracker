package main

import (
	"delivery-tracker-go/internal/models"
)

type Handler struct {
	orders *models.OrderModel
}

func NewHandler(dbmodel *models.DBModel) *Handler {
	return &Handler{
		orders: &dbmodel.Order,
	}
}
