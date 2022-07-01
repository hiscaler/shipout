package shipout

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hiscaler/shipout-go/entity"
	jsoniter "github.com/json-iterator/go"
)

// 增值服务
// https://open.shipout.com/portal/zh/api/add-value.html
type valueAddedService service

type WarehouseValueAddedServicesQueryParams struct {
	Typ         int    `url:"type,omitempty"` // 业务类型：11.入库/6.出库-小包/12.出库-大货/10.退货/99.通用
	WarehouseId string `url:"warehouseId"`    // 仓库 ID
}

func (m WarehouseValueAddedServicesQueryParams) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Typ, validation.When(m.Typ > 0, validation.In(11, 6, 12, 10, 99).Error("无效的业务类型"))),
		validation.Field(&m.WarehouseId, validation.Required.Error("仓库 ID 不能为空")),
	)
}

func (s valueAddedService) Warehouses(params WarehouseValueAddedServicesQueryParams) (items []entity.WarehouseValueAddedService, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	res := struct {
		NormalResponse
		Data []entity.WarehouseValueAddedService `json:"data"`
	}{}
	resp, err := s.httpClient.R().
		SetQueryParamsFromValues(toValues(params)).
		Get("/open-api/oms/addValueService/warehouse/query")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		if err = ErrorWrap(res.ErrorCode, res.Message); err == nil {
			if err = jsoniter.Unmarshal(resp.Body(), &res); err == nil {
				items = res.Data
				isLastPage = true
			}
		}
	} else {
		if e := jsoniter.Unmarshal(resp.Body(), &res); e == nil {
			err = ErrorWrap(res.ErrorCode, res.Message)
		} else {
			err = errors.New(resp.Status())
		}
	}
	return
}
