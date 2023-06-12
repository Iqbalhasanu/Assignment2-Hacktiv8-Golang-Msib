package item_repository

import (
	"github.com/Iqbalhasanu/Assignment2-Hacktiv8-Golang-Msib/entity"
	"github.com/Iqbalhasanu/Assignment2-Hacktiv8-Golang-Msib/pkg/errs"
)

type ItemRepository interface {
	FindItemsByItemCodes(itemCodes []string) ([]*entity.Item, errs.MessageErr)
}
