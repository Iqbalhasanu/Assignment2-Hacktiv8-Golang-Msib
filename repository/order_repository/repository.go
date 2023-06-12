package order_repository

import (
	"github.com/Iqbalhasanu/Assignment2-Hacktiv8-Golang-Msib/entity"
	"github.com/Iqbalhasanu/Assignment2-Hacktiv8-Golang-Msib/pkg/errs"
)

type OrderRepository interface {
	CreateOrder(orderPayload entity.Order, itemsPayload []entity.Item) (*entity.Order, errs.MessageErr)
	GetAllOrder() (*[]OrderItem, errs.MessageErr)
	UpdateOrder(orderPayload entity.Order, itemsPayload []entity.Item) (*OrderItem, errs.MessageErr)
	DeleteOrder(orderId int) errs.MessageErr
}
