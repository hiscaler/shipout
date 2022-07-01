package entity

type Order struct {
	ID        string          `json:"orderId"`        // 订单 ID
	NO        string          `json:"orderNO"`        // 订单号
	Recipient OrderRecipient  `json:"orderRecipient"` // 收货人信息
	Shipments []OrderShipment `json:"orderShipments"` // 发货列表
	Summary   OrderSummary    `json:"orderSummary"`   // 订单摘要
	Status    int             `json:"status"`         // 状态
}

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

// OrderShipmentInternational 发货单国际单补充
type OrderShipmentInternational struct {
	EeiType                int    `json:"eeiType"`                // EEI类型 1.需要豁免 2.已有ITN Number; 一般总申报金额低于2500时需要豁免
	EinOrSsn               string `json:"einOrSsn"`               // 找仓库要的ein码，在美国发国外用的一个雇主识别码，不发国外的话不需要
	ForeignTradeRegulation string `json:"foreignTradeRegulation"` // FTR码, 如果是去往非加拿大国家的货物,没有超过2500时豁免码为30.37(a); 如果是去往加拿大,没有超过2500时豁免码为30.36
	ItnNumber              string `json:"itnNumber"`              // ITN Number，EEI类型为已有ITN Number时必传
}

type OrderShipmentPackageDimension struct {
	DistanceUnit         string  `json:"distanceUnit"`
	Height               float64 `json:"height"`
	LengthWithAroundSize float64 `json:"lenghtWithAroundSize"`
	Length               float64 `json:"length"`
	Width                float64 `json:"width"`
}

type OrderShipmentPackageWeight struct {
	MassUnit string  `json:"massUnit"`
	Weight   float64 `json:"weight"`
	Weight1  float64 `json:"weight1"`
	Weight2  float64 `json:"weight2"`
}

type OrderShipmentPackage struct {
	Dimension OrderShipmentPackageDimension `json:"dimension"`
	Weight    OrderShipmentPackageWeight    `json:"weight"`
}

type OrderShipmentProduct struct {
	OmsSKU   string  `json:"omsSku"` // 产品的sku(skuId和omsSku至少传一个)
	Price    float64 `json:"price"`  // 单价
	Quantity int     `json:"qty"`    // 数量（大于0）
	SKUId    string  `json:"skuId"`  // 系统产品主键(skuId和omsSku至少传一个)
}

// OrderShipment 发货信息
type OrderShipment struct {
	FeesDetail     string                     `json:"feesDetail"`     // 金额详情，格式为{“费用名”: 金额，“费用名”: 金额}
	International  OrderShipmentInternational `json:"international"`  // 发货单国际单补充
	LabelStatus    int                        `json:"labelStatus"`    // label状态：0,1,2,3,4,5：无须打单,未打单,已打单,运输中,已运输,异常
	OutboundNO     string                     `json:"outboundNO"`     // 出库单号
	OutboundStatus int                        `json:"outboundStatus"` // 0,1,2,3,4：无须建立出库单,未建立,已建立,仓库处理中,已发货
	OutboundingId  string                     `json:"outboundingId"`  // outboundingId
	OutbundingId   string                     `json:"outbundingId"`   // outbundingId
	PackageList    []OrderShipmentPackage     `json:"packageList"`    // 发货单包裹列表
	ProductList    []OrderShipmentProduct     `json:"productList"`    // 发货单产品列表
	Rate           float64                    `json:"rate"`           // 金额
	ShipmentId     string                     `json:"shipmentId"`     // shipmentId
	ShippingRate   float64                    `json:"shippingRate"`   // 物流金额
	SysServiceId   string                     `json:"sysServiceId"`   // 运输类型 7.不需要打单 9.使用仓库选择的服务
	TrackingNumber string                     `json:"trackingNumber"` // 物流追踪号
	WarehouseId    string                     `json:"warehouseId"`    // 仓库编号
}

// OrderSummary 订单摘要
type OrderSummary struct {
	Age           int    `json:"age,omitempty"`
	CreateTime    string `json:"createTime,omitempty"`
	NoteFromBuyer string `json:"note_from_buyer,omitempty"`
	OrderDate     string `json:"orderDate"` // 订单创建时间,格式:yyyy-MM-dd HH:mm:ss
}
