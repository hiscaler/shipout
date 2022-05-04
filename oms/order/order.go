package order

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hiscaler/shipout-go"
	"github.com/hiscaler/shipout-go/pkg/cast"
	jsoniter "github.com/json-iterator/go"
)

// 订单
// https://open.shipout.com/portal/zh/api/order.html

type BatchSubmitResult struct {
	OrderId       string   `json:"order_id"`
	OrderNo       string   `json:"orderNo"`
	FulfillResult []string `json:"fulfillResult"`
}

// OrderSummary 订单摘要
type OrderSummary struct {
	Age        int    `json:"age"`
	CreateTime string `json:"createTime"`
	OrderDate  string `json:"orderDate"` // 订单创建时间,格式:yyyy-MM-dd HH:mm:ss
}

// ShipmentFormInternational 发货单国际单补充
type ShipmentFormInternational struct {
	EEIType                int    `json:"eeiType"`                // EEI类型 1.需要豁免 2.已有ITN Number; 一般总申报金额低于2500时需要豁免
	EinOrSsn               string `json:"einOrSsn"`               // 找仓库要的ein码，在美国发国外用的一个雇主识别码，不发国外的话不需要
	ForeignTradeRegulation string `json:"foreignTradeRegulation"` // FTR码, 如果是去往非加拿大国家的货物,没有超过2500时豁免码为30.37(a); 如果是去往加拿大,没有超过2500时豁免码为30.36
	ITNNumber              string `json:"itnNumber"`              // ITN Number，EEI类型为已有ITN Number时必传
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
	CarrierId     int    `json:"carrierId"`     // 运输商: 1. USPS 2. UPS 3. FedEx 4. DHL 9. Other
	ShipDate      string `json:"shipDate"`      // 计划执运日期，即计划发货日期，格式：“yyyy-MM-dd 00:00:00”
	ShipmentSid   string `json:"shipmentSid"`   // shipment序号
	SignatureType int    `json:"signatureType"` // 签名类型：1.Indirect (FedEx,UPS only) 2.DIRECT 3.ADULT 4.SERVICE_DEFAULT(default)
}

// ShipmentForm 执运信息，即发货表单
type ShipmentForm struct {
	International ShipmentFormInternational `json:"international"` // 发货单国际单补充
	OutboundInfo  ShipmentFormOutboundInfo  `json:"outboundInfo"`  // 出库单信息
	ProductList   []ShipmentFormProduct     `json:"productList"`   // 发货单产品列表
	ShippingInfo  ShipmentFormShippingInfo  `json:"shippingInfo"`  // 发货单基本信息
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

type BatchSubmitOrderRequest struct {
	OrderNo       string         `json:"orderNo"`       // openapi允许客户传入的订单编号
	OrderSummary  OrderSummary   `json:"OrderSummary"`  // 订单摘要
	ShipmentForms []ShipmentForm `json:"shipmentForms"` // 执运信息，即发货表单
	SID           int            `json:"sid"`           // 当前请求序号，当前请求内不能重复; 是批量传订单的一个标识,只当前请求有效，无业务含义
	ToAddress     ToAddress      `json:"toAddress"`     // 收货地址
	WarehouseId   string         `json:"warehouseId"`   // 仓库 ID
}

func (m BatchSubmitOrderRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.WarehouseId, validation.Required.Error("仓库 ID 不能为空")),
	)
}

func (s service) BatchSubmit(req BatchSubmitOrderRequest) (items []BatchSubmitResult, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	res := struct {
		shipout.NormalResponse
		Data []BatchSubmitResult `json:"data"`
	}{}

	resp, err := s.shipOut.Client.R().SetBody(req).Post("/open-api/oms/order/batchSubmit")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		if err = shipout.ErrorWrap(res.ErrorCode, res.Message); err == nil {
			items = res.Data
		}
	} else {
		if e := jsoniter.Unmarshal(resp.Body(), &res); e == nil {
			err = shipout.ErrorWrap(res.ErrorCode, res.Message)
		} else {
			err = errors.New(resp.Status())
		}
	}
	return
}

// 订单列表

type OrdersResRecordShipment struct {
	FeesDetail     string  `json:"feesDetail"`     // 金额详情，格式为{“费用名”: 金额，“费用名”: 金额}
	LabelStatus    int     `json:"labelStatus"`    // label状态：0,1,2,3,4,5：无须打单,未打单,已打单,运输中,已运输,异常
	OutboundStatus int     `json:"outboundStatus"` // 0,1,2,3,4：无须建立出库单,未建立,已建立,仓库处理中,已发货
	OutbundingId   string  `json:"outbundingId"`   // outbundingId
	Rate           float64 `json:"rate"`           // 金额
	ShipmentId     string  `json:"shipmentId"`     // shipmentId
	ShippingRate   float64 `json:"shippingRate"`   // 物流金额
	TrackingNumber string  `json:"trackingNumber"` // 物流追踪号
	WarehouseId    string  `json:"warehouseId"`    // 仓库编号
}

type OrdersResRecord struct {
	FulfillCharge  float64                   `json:"fulfillCharge"` // 总金额,完成计费后才会有
	OrderDate      string                    `json:"orderDate"`     // 订单日期
	OrderId        string                    `json:"orderId"`
	OrderNO        string                    `json:"orderNO"`
	RecipientName  string                    `json:"recipientName"`
	ShipTo         string                    `json:"shipTo"`
	ShipmentList   []OrdersResRecordShipment `json:"shipmentList"`
	ShippingCharge float64                   `json:"shippingCharge"` // 运费总金额,打单后就会有
	Status         int                       `json:"status"`         // 状态
	ZipCode        string                    `json:"zipCode"`        // 邮编
}
type OrdersRes struct {
	CountId          string `json:"countId"`
	Current          int    `json:"current"`
	HitCount         bool   `json:"hitCount"`
	MaxLimit         int    `json:"maxLimit"`
	OptimizeCountSQL bool   `json:"optimizeCountSql"`
	Pages            int    `json:"pages"`
	SearchCount      bool   `json:"searchCount"`
	Size             int    `json:"size"`
	Total            int    `json:"total"`
}

type OrdersQueryParams struct {
	Asc         bool   `json:"asc,omitempty"`
	CurPageNo   int    `json:"curPageNo,omitempty"`
	HiDirection string `json:"hiDirection,omitempty"`
	Name        string `json:"name,omitempty"`
	OrderColumn string `json:"orderColumn,omitempty"`
	PageSize    int    `json:"pageSize,omitempty"`
}

func (m OrdersQueryParams) Validate() error {
	return nil
}

type OrdersQueryBody struct {
	OrgId       string `json:"orgId"`
	ServerOrgId string `json:"serverOrgId"`
	Status      int    `json:"status"`
	WarehouseId string `json:"warehouseId"`
}

func (s service) Orders(params OrdersQueryParams, body OrdersQueryBody) (items []OrdersRes, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	res := struct {
		shipout.NormalResponse
		Data []OrdersRes `json:"data"`
	}{}

	resp, err := s.shipOut.Client.R().
		SetQueryParamsFromValues(cast.StructToURLValues(params)).
		SetBody(&body).
		Get("/open-api/oms/order/queryList")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		if err = shipout.ErrorWrap(res.ErrorCode, res.Message); err == nil {
			if err = jsoniter.Unmarshal(resp.Body(), &res); err == nil {
				items = res.Data
			}
		}
	} else {
		if e := jsoniter.Unmarshal(resp.Body(), &res); e == nil {
			err = shipout.ErrorWrap(res.ErrorCode, res.Message)
		} else {
			err = errors.New(resp.Status())
		}
	}
	return
}

// 单个订单查询

// OrderRecipient 收货人信息
type OrderRecipient struct {
	AddressLine1 string `json:"addressLine1"` // 收件人地址行1
	AddressLine2 string `json:"addressLine2"` // 收件人地址行2
	City         string `json:"city"`         // 城市
	Company      string `json:"company"`      // 收件人公司
	CountryCode  string `json:"countryCode"`  // 国家编码，格式标准遵循ISO 3166-1 alpha-2
	Email        string `json:"email"`        // 收件人邮箱
	Name         string `json:"name"`         // 收件人姓名
	Phone        string `json:"phone"`        // 收件人联系电话
	Residential  bool   `json:"residential"`  //
	StateCode    string `json:"stateCode"`    // 州代码，美国为两位大写，如CA、NY
	ZipCode      string `json:"zipCode"`      // 邮政编码
}

// OrderShipment 发货信息
type OrderShipment struct {
}

type Order struct {
	OrderId        string          `json:"orderId"`        // 订单 ID
	OrderRecipient OrderRecipient  `json:"orderRecipient"` // 收货人信息
	OrderShipments []OrderShipment `json:"orderShipments"` // 发货列表
	OrderSummary   []OrderSummary  `json:"orderSummary"`   // 订单摘要
	Status         int             `json:"status"`         // 状态
}

type OrderQueryParams struct {
	Name    string `json:"name,omitempty"`
	OrderId string `json:"orderId"`
}

func (m OrderQueryParams) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.OrderId, validation.Required.Error("订单 ID 不能为空")),
	)
}

func (s service) Order(params OrderQueryParams) (item Order, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	res := struct {
		shipout.NormalResponse
		Data Order `json:"data"`
	}{}

	resp, err := s.shipOut.Client.R().SetBody(params).Get("/open-api/oms/order/query")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		if err = shipout.ErrorWrap(res.ErrorCode, res.Message); err == nil {
			item = res.Data
		}
	} else {
		if e := jsoniter.Unmarshal(resp.Body(), &res); e == nil {
			err = shipout.ErrorWrap(res.ErrorCode, res.Message)
		} else {
			err = errors.New(resp.Status())
		}
	}
	return
}
