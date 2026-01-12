package main

import (
	"delivery-tracker-go/internal/models"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerData struct {
	Title    string
	Order    models.Order
	Statuses []string
}

type OrderFormData struct {
	OrderType []string
	OrderSize []string
}

type OrderRequest struct {
	Name         string   `form:"name" binding:"required,min=2,max=100"`
	Phone        string   `form:"phone" binding:"required,min=2,max=20"`
	Address      string   `form:"address" binding:"required,min=2,max=200"`
	OrderSizes   []string `form:"size" binding:"required,min=1,dive,valid_order_size"`
	OrderTypes   []string `form:"order" binding:"required,min=1,dive,valid_order_type"`
	Instructions []string `form:"instructions" binding:"max=300"`
}

func (h *Handler) ServeNewOrderForm(c *gin.Context) {
	c.HTML(http.StatusOK, "order.tmpl", OrderFormData{
		OrderType: models.OrderTypes,
		OrderSize: models.OrderSizes,
	})
}

func (h *Handler) HandleNewOrderPost(c *gin.Context) {
	var form OrderRequest
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	orderItems := make([]models.OrderItem, len(form.OrderSizes))
	for i := range orderItems {
		orderItems[i] = models.OrderItem{
			Size:         form.OrderSizes[i],
			Package:      form.OrderTypes[i],
			Instructions: form.Instructions[i],
		}
	}

	order := models.Order{
		CustomerName: form.Name,
		Phone:        form.Phone,
		Address:      form.Address,
		Status:       models.OrderStatuses[0],
		Items:        orderItems,
	}

	if err := h.orders.CreateOrder(&order); err != nil {
		slog.Error("Failed to create order", "error", err)
		c.String(http.StatusInternalServerError, "Some went wrong")
		return
	}

	slog.Info("Order created", "orderId", order.ID, "customer", order.CustomerName)

	h.notificationManager.Notify("admin:new_orders", "new_order")
	c.Redirect(http.StatusSeeOther, "/customer/"+order.ID)
}

func (h *Handler) ServeCustomer(c *gin.Context) {
	orderID := c.Param("id")
	if orderID == "" {
		c.String(http.StatusBadRequest, "Order ID is required")
	}
	order, err := h.orders.GetOrder(orderID)
	if err != nil {
		c.String(http.StatusNotFound, "Order not found")
		return
	}

	c.HTML(http.StatusOK, "customer.tmpl", CustomerData{
		Title:    "Order status " + orderID,
		Order:    *order,
		Statuses: models.OrderStatuses,
	})
}
