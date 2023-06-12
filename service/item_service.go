package service

import (
	"github.com/Iqbalhasanu/Assignment2-Hacktiv8-Golang-Msib/entity"
	"github.com/Iqbalhasanu/Assignment2-Hacktiv8-Golang-Msib/pkg/errs"
	"github.com/Iqbalhasanu/Assignment2-Hacktiv8-Golang-Msib/repository/item_repository"
)

type itemService struct {
	itemRepo item_repository.ItemRepository
}

type ItemService interface {
	FindItemsByItemCodes(itemCodes []string) ([]*entity.Item, errs.MessageErr)
}

func NewItemService(itemRepo item_repository.ItemRepository) ItemService {
	return &itemService{
		itemRepo: itemRepo,
	}
}

func (i *itemService) FindItemsByItemCodes(itemCodes []string) ([]*entity.Item, errs.MessageErr) {
	items, err := i.itemRepo.FindItemsByItemCodes(itemCodes)

	if err != nil {
		return nil, err
	}

	for _, eachItemCode := range itemCodes {
		isFound := false

		for _, eachItem := range items {
			if eachItemCode == eachItem.ItemCode {
				isFound = true
				break
			}
		}

		if !isFound {
			return nil, errs.NewNotFoundError("notfound")
		}
	}

	return items, err
}
