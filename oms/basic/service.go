package basic

import (
	"github.com/hiscaler/shipout-go"
)

type service struct {
	shipOut *shipout.ShipOut
}

type Service interface {
	Warehouses(params WarehousesQueryParams) (items []Warehouse, isLastPage bool, err error) // 仓库列表
}

func NewService(client *shipout.ShipOut) Service {
	return service{client}
}
