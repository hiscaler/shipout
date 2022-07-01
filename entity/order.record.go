package entity

// OrderRecord 订单数据（用于订单列表）
// type OrderRecord struct {
// 	FulfillCharge float64 `json:"fulfillCharge"` // 总金额,完成计费后才会有
// 	OrderDate     string  `json:"orderDate"`     // 订单日期
// 	OrderId       string  `json:"orderId"`
// 	OrderNO       string  `json:"orderNO"`
// 	recipientName string  `json:"recipientName"` // 收货人名称
// 	shipTo        string  `json:"shipTo"`        // 收货地址
// }

type OrderRecord struct {
	FulfillCharge  float64                `json:"fulfillCharge"` // 总金额,完成计费后才会有
	OrderDate      string                 `json:"orderDate"`     // 订单日期
	OrderId        string                 `json:"orderId"`
	OrderNO        string                 `json:"orderNO"`
	RecipientName  string                 `json:"recipientName"`  // 收货人名称
	ShipTo         string                 `json:"shipTo"`         // 收货地址
	ShipmentList   []OrdersRecordShipment `json:"shipmentList"`   // shipment list(每条order内嵌一个ShipmentList的返回结果)
	ShippingCharge float64                `json:"shippingCharge"` // 运费总金额,打单后就会有
	Status         int                    `json:"status"`         // 状态
	ZipCode        string                 `json:"zipCode"`        // 邮编
}

type OrdersRecordShipment struct {
	FeesDetail     string  `json:"feesDetail"`     // 金额详情，格式为{“费用名”: 金额，“费用名”: 金额}
	LabelStatus    int     `json:"labelStatus"`    // label状态：0,1,2,3,4,5：无须打单,未打单,已打单,运输中,已运输,异常
	OutboundNO     string  `json:"outboundNO"`     // 出库单号
	OutboundStatus int     `json:"outboundStatus"` // 0,1,2,3,4：无须建立出库单,未建立,已建立,仓库处理中,已发货
	OutboundingId  string  `json:"outboundingId"`  // outboundingId
	OutbundingId   string  `json:"outbundingId"`   // outbundingId
	Rate           float64 `json:"rate"`           // 金额
	ShipmentId     string  `json:"shipmentId"`     // shipmentId
	ShippingRate   float64 `json:"shippingRate"`   // 物流金额
	TrackingNumber string  `json:"trackingNumber"` // 物流追踪号
	WarehouseId    string  `json:"warehouseId"`    // 仓库编号
}
