package order

import (
	"github.com/hiscaler/shipout-go"
)

type service struct {
	shipOut *shipout.ShipOut
}

type Service interface {
	BatchSubmit(req BatchSubmitOrderRequest) (items []Order, err error) // 批量订单提交
}

func NewService(client *shipout.ShipOut) Service {
	return service{client}
}
