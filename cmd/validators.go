package main

import (
	"delivery-tracker-go/internal/models"
	"slices"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
)

func RegisterCustomValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("valid_order_type", createSliceValidator(models.OrderStatuses))
		v.RegisterValidation("valid_order_size", createSliceValidator(models.OrderSizes))
	}
}

func createSliceValidator(allowedValues []string) validator.Func {
	return func(fl validator.FieldLevel) bool {
		return slices.Contains(allowedValues, fl.Field().String())
	}
}
