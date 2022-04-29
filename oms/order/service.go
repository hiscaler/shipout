package order

import (
	"github.com/hiscaler/shipout-go"
)

type service struct {
	shipOut *shipout.ShipOut
}

type Service interface {
	BatchSubmit(req BatchSubmitOrderRequest) (items []BatchSubmitResult, err error) // 批量订单提交
	Order(params OrderQueryParams) (item BatchSubmitResult, err error)              // 单个订单明细查询
}

func NewService(client *shipout.ShipOut) Service {
	return service{client}
}
