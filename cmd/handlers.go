package main

import (
	"delivery-tracker-go/internal/models"
)

type Handler struct {
	orders *models.OrderModel
	users  *models.UserModel
}

func NewHandler(dbmodel *models.DBModel) *Handler {
	return &Handler{
		orders: &dbmodel.Order,
		users:  &dbmodel.User,
	}
}
