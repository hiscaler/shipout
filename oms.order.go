package shipout

import (
	"context"
	"errors"
	mapset "github.com/deckarep/golang-set"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hiscaler/shipout-go/constant"
	"github.com/hiscaler/shipout-go/entity"
	jsoniter "github.com/json-iterator/go"
)

// 订单
// https://open.shipout.com/portal/zh/api/order.html
type orderService service

type BatchSubmitResult struct {
	OrderId       string   `json:"order_id"`
	OrderNo       string   `json:"orderNo"`
	FulfillResult []string `json:"fulfillResult"`
}

// OrderSummary 订单摘要
type OrderSummary struct {
	Age           int    `json:"age,omitempty"`
	CreateTime    string `json:"createTime,omitempty"`
	NoteFromBuyer string `json:"note_from_buyer,omitempty"`
	OrderDate     string `json:"orderDate"` // 订单创建时间,格式:yyyy-MM-dd HH:mm:ss
}

func (m OrderSummary) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.OrderDate,
			validation.Required.Error("订单创建时间不能为空"),
			validation.When(m.OrderDate != "", validation.Date(constant.DatetimeFormat).Error("无效的订单创建时间格式")),
		),
	)
}

// ShipmentFormInternational 发货单国际单补充
type ShipmentFormInternational struct {
	EEIType                int    `json:"eeiType"`                          // EEI类型 1.需要豁免 2.已有ITN Number; 一般总申报金额低于2500时需要豁免
	EinOrSsn               string `json:"einOrSsn"`                         // 找仓库要的ein码，在美国发国外用的一个雇主识别码，不发国外的话不需要
	ForeignTradeRegulation string `json:"foreignTradeRegulation,omitempty"` // FTR码, 如果是去往非加拿大国家的货物,没有超过2500时豁免码为30.37(a); 如果是去往加拿大,没有超过2500时豁免码为30.36
	ITNNumber              string `json:"itnNumber,omitempty"`              // ITN Number，EEI类型为已有ITN Number时必传
}

// ShipmentOutboundLabel label 信息
type ShipmentOutboundLabel struct {
	LabelURL []string `json:"labelUrl"` // 运单url地址
}

// ShipmentFormOutboundInfo 出库单信息
type ShipmentFormOutboundInfo struct {
	AdditionService       []string                `json:"additionService"`       // 附加服务
	Remark                string                  `json:"remark"`                // 备注
	ShipmentOutboundLabel []ShipmentOutboundLabel `json:"shipmentOutboundLabel"` // label 信息
	SysServiceId          string                  `json:"sysServiceId"`          // 运输类型 6.客户自己上传物流单号和运单地址 7.不需要打单（客户去仓库自提） 9.使用仓库选择的服务,即要打物流单
	TrackingNumber        string                  `json:"trackingNumber"`        // 物流跟踪号
}

// ShipmentFormProduct 发货单产品列表
type ShipmentFormProduct struct {
	Price    float64 `json:"price"` // 单价
	Quantity int     `json:"qty"`   // 数量
	SKUId    string  `json:"skuId"` // 系统产品主键
}

// ShipmentFormShippingInfo 发货单基本信息
type ShipmentFormShippingInfo struct {
	CarrierId     int    `json:"carrierId,omitempty"` // 运输商: 1. USPS 2. UPS 3. FedEx 4. DHL 9. Other
	ShipDate      string `json:"shipDate"`            // 计划执运日期，即计划发货日期，格式：“yyyy-MM-dd 00:00:00”
	ShipmentSid   string `json:"shipmentSid"`         // shipment序号
	SignatureType int    `json:"signatureType"`       // 签名类型：1.Indirect (FedEx,UPS only) 2.DIRECT 3.ADULT 4.SERVICE_DEFAULT(default)
}

// ShipmentForm 执运信息，即发货表单
type ShipmentForm struct {
	International ShipmentFormInternational `json:"international"` // 发货单国际单补充
	OutboundInfo  ShipmentFormOutboundInfo  `json:"outboundInfo"`  // 出库单信息
	ProductList   []ShipmentFormProduct     `json:"productList"`   // 发货单产品列表
	ShippingInfo  ShipmentFormShippingInfo  `json:"shippingInfo"`  // 发货单基本信息
}

func (m ShipmentFormInternational) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.EEIType, validation.In(1, 2).Error("无效的 EFI 类型")),
		validation.Field(&m.EinOrSsn, validation.Required.Error("EIN 码不能为空")),
		validation.Field(&m.ForeignTradeRegulation, validation.When(m.EEIType == 1, validation.Required.Error("FTR 码不能为空"))),
		validation.Field(&m.ITNNumber, validation.When(m.EEIType == 2, validation.Required.Error("ITN 不能为空"))),
	)
}
func (m ShipmentFormShippingInfo) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.ShipDate,
			validation.Required.Error("计划发货日期不能为空"),
			validation.Date(constant.DatetimeFormat).Error("无效的计划发货日期格式"),
		),
		validation.Field(&m.ShipmentSid, validation.Required.Error("shipment 序号不能为空")),
		validation.Field(&m.SignatureType, validation.Required.Error("签名类型不能为空"), validation.In(1, 2, 3, 4).Error("无效的签名类型")),
	)
}

type ToAddress struct {
	AddressLine1 string `json:"addressLine1"` // 收件人地址行1
	AddressLine2 string `json:"addressLine2"` // 收件人地址行2
	City         string `json:"city"`         // 城市
	Company      string `json:"company"`      // 收件人公司
	CountryCode  string `json:"countryCode"`  // 国家编码，格式标准遵循ISO 3166-1 alpha-2
	Email        string `json:"email"`        // 收件人邮箱
	Name         string `json:"name"`         // 收件人姓名
	Phone        string `json:"phone"`        // 收件人联系电话
	StateCode    string `json:"stateCode"`    // 州代码，美国为两位大写，如CA、NY
	ZipCode      string `json:"zipCode"`      // 邮政编码
}

func (m ToAddress) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.ZipCode, validation.Required.Error("收货地址邮编不能为空")),
	)
}

type SubmitOrderRequest struct {
	OrderNo       string         `json:"orderNo,omitempty"` // openapi 允许客户传入的订单编号
	OrderSummary  OrderSummary   `json:"OrderSummary"`      // 订单摘要
	ShipmentForms []ShipmentForm `json:"shipmentForms"`     // 执运信息，即发货表单
	SID           int            `json:"sid"`               // 当前请求序号，当前请求内不能重复; 是批量传订单的一个标识,只当前请求有效，无业务含义
	ToAddress     ToAddress      `json:"toAddress"`         // 收货地址
	WarehouseId   string         `json:"warehouseId"`       // 仓库 ID
}

type BatchSubmitOrderRequest []SubmitOrderRequest

func (m BatchSubmitOrderRequest) Validate() error {
	or := make([]SubmitOrderRequest, len(m))
	for i, request := range m {
		or[i] = request
	}
	return validation.Validate(or,
		validation.Required.Error("请求数据不能为空"),
		validation.By(func(value interface{}) (err error) {
			reqs, ok := value.([]SubmitOrderRequest)
			if !ok {
				return errors.New("无效的提交订单")
			}

			sids := mapset.NewSet()
			for _, req := range reqs {
				err = validation.ValidateStruct(&req,
					// validation.Field(&m.OrderNo, validation.Required.Error("订单编号不能为空")),
					validation.Field(&req.WarehouseId, validation.Required.Error("仓库 ID 不能为空")),
					validation.Field(&req.SID, validation.Required.Error("当前请求序号不能为空")),
					validation.Field(&req.OrderSummary),
					validation.Field(&req.ToAddress),
					validation.Field(&req.ShipmentForms,
						validation.Required.Error("发货单不能为空"),
						validation.Each(validation.WithContext(func(ctx context.Context, value interface{}) error {
							form, ok := value.(ShipmentForm)
							if !ok {
								return errors.New("无效的发货单数据")
							}
							return validation.ValidateStruct(&form,
								validation.Field(&form.International),
								validation.Field(&form.ProductList,
									validation.Required.Error("发货单产品不能为空"),
									validation.Each(validation.WithContext(func(ctx context.Context, value interface{}) error {
										product, ok := value.(ShipmentFormProduct)
										if !ok {
											return errors.New("无效的发货单商品")
										}
										return validation.ValidateStruct(&product,
											validation.Field(&product.SKUId, validation.Required.Error("系统产品主键不能为空")),
											validation.Field(&product.Quantity, validation.Min(1).Error("商品数量不能少于 {{.threshold}}")),
											validation.Field(&product.Price, validation.Min(0.01).Error("商品数量不能小于 {{.threshold}}")),
										)
									})),
								),
								validation.Field(&form.ShippingInfo),
							)
						}))),
				)
				if err != nil {
					return
				}
				sids.Add(req.SID)
			}
			if err == nil && sids.Cardinality() != len(reqs) {
				err = errors.New("SID 值重复")
			}
			return
		}),
	)
}

func (s orderService) BatchSubmit(req BatchSubmitOrderRequest) (items []BatchSubmitResult, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	res := struct {
		NormalResponse
		Data []BatchSubmitResult `json:"data"`
	}{}

	resp, err := s.httpClient.R().
		SetBody(req).
		Post("/open-api/oms/order/batchSubmit")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		if err = ErrorWrap(res.ErrorCode, res.Message); err == nil {
			items = res.Data
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

// 订单列表

type OrdersQueryParams struct {
	Asc         bool   `url:"asc,omitempty"`
	CurPageNo   int    `url:"curPageNo,omitempty"`
	HiDirection string `url:"hiDirection,omitempty"`
	Name        string `url:"name,omitempty"`
	OrderColumn string `url:"orderColumn,omitempty"`
	OrderNo     string `url:"orderNO,omitempty"`
	OrgId       string `url:"orgId,omitempty"`
	ServerOrgId string `url:"serverOrgId,omitempty"`
	Status      int    `url:"status,omitempty"` // 1.Awaiting Payment 2.Awaiting Fulillment 3.Being Fulilled 4.Shipped 5.Delievered 6.Cancelled 7.RMA Initiated 8.RMA Processing 9.RMA Completed
	WarehouseId string `url:"warehouseId,omitempty"`
	PageSize    int    `url:"pageSize,omitempty"`
}

func (m OrdersQueryParams) Validate() error {
	return nil
}

func (s orderService) All(params OrdersQueryParams) (items []entity.OrderRecord, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	res := struct {
		NormalResponse
		Data struct {
			CountId          string               `json:"countId"`
			Current          int                  `json:"current"`
			HitCount         bool                 `json:"hitCount"`
			MaxLimit         int                  `json:"maxLimit"`
			OptimizeCountSQL bool                 `json:"optimizeCountSql"`
			Pages            int                  `json:"pages"`
			Records          []entity.OrderRecord `json:"records"`
			IsSearchCount    bool                 `json:"IsSearchCount"`
			Size             string               `json:"size"`
			Total            string               `json:"total"`
		} `json:"data"`
	}{}

	resp, err := s.httpClient.R().
		SetQueryParamsFromValues(toValues(params)).
		Get("/open-api/oms/order/queryListV2")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		if err = ErrorWrap(res.ErrorCode, res.Message); err == nil {
			if err = jsoniter.Unmarshal(resp.Body(), &res); err == nil {
				items = res.Data.Records
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

// 单个订单查询

// OrderQueryParams 订单查询参数
type OrderQueryParams struct {
	Name        string `url:"name,omitempty"`
	OrderId     string `url:"orderId,omitempty"`     // 订单id
	OutboundNO  string `url:"outboundNO,omitempty"`  // 出库单号
	WarehouseId string `url:"warehouseId,omitempty"` // 仓库 id
}

func (m OrderQueryParams) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.OrderId, validation.Required.Error("订单 ID 不能为空")),
	)
}

func (s orderService) One(params OrderQueryParams) (item entity.Order, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	res := struct {
		NormalResponse
		Data entity.Order `json:"data"`
	}{}
	resp, err := s.httpClient.R().
		SetQueryParamsFromValues(toValues(params)).
		Get("/open-api/oms/order/query")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		if err = jsoniter.Unmarshal(resp.Body(), &res); err == nil {
			if err = ErrorWrap(res.ErrorCode, res.Message); err == nil {
				item = res.Data
			}
		}
	} else {
		err = errors.New(resp.Status())
	}
	return
}

// 订单取消
// 取消成功后返回出库单 id

type CancelOrderRequest struct {
	Name    string `url:"name,omitempty"`
	OrderId string `url:"order_id"`
}

func (m CancelOrderRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.OrderId, validation.Required.Error("订单 ID 不能为空")),
	)
}

func (s orderService) Cancel(req CancelOrderRequest) (id string, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	res := struct {
		NormalResponse
		Data string `json:"data"`
	}{}
	resp, err := s.httpClient.R().
		SetQueryParamsFromValues(toValues(req)).
		Post("/open-api/oms/order/cancel")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		if err = jsoniter.Unmarshal(resp.Body(), &res); err == nil {
			if err = ErrorWrap(res.ErrorCode, res.Message); err == nil {
				id = res.Data
			}
		}
	} else {
		err = errors.New(resp.Status())
	}
	return
}
