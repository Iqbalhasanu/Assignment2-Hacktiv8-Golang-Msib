package entity

import (
	"time"

	"github.com/IqbalHasanu/Assignment2-Hacktiv8-Golang-Msib/dto"
)

type Item struct {
	ItemId      int
	ItemCode    string
	Quantity    int
	Description string
	OrderId     int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (i *Item) ItemToItemResponse() dto.ItemResponse {
	return dto.ItemResponse{
		Id:          i.ItemId,
		ItemCode:    i.ItemCode,
		Quantity:    i.Quantity,
		Description: i.Description,
		OrderId:     i.OrderId,
		CreatedAt:   i.CreatedAt,
		UpdatedAt:   i.UpdatedAt,
	}
}
