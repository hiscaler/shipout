package entity

type BundleDetail struct {
	ASIN             string  `json:"asin"`
	AuditStatus      int     `json:"auditStatus"`
	BatteryFlag      bool    `json:"batteryFlag"`
	Brand            string  `json:"brand"`
	ChildQuantity    int     `json:"childQuantity"`
	ChildSkuId       string  `json:"childSkuId"`
	CreateTime       string  `json:"createTime"`
	DeclaredValue    float64 `json:"declaredValue"`
	Dimension        string  `json:"dimension"`
	DistanceUnit     string  `json:"distanceUnit"`
	Ean              string  `json:"ean"`
	ExtraBarcode1    string  `json:"extraBarcode1"`
	FnSku            string  `json:"fnSku"`
	Height           float64 `json:"height"`
	Length           float64 `json:"length"`
	MassUnit         string  `json:"massUnit"`
	Note             string  `json:"note"`
	OmsSku           string  `json:"omsSku"`
	ParentSkuId      string  `json:"parentSkuId"`
	ScheduleB        string  `json:"scheduleB"`
	SerialNumberFlag bool    `json:"serialNumberFlag"`
	ShippingType     int     `json:"shippingType"`
	ShortCode        string  `json:"shortCode"`
	SkuId            string  `json:"skuId"`
	SkuNameCN        string  `json:"skuNameCN"`
	SkuNameEN        string  `json:"skuNameEN"`
	Upc              string  `json:"upc"`
	UpdateTime       string  `json:"updateTime"`
	Weight           float64 `json:"weight"`
	Weight1          float64 `json:"weight1"`
	Weight2          float64 `json:"weight2"`
	Width            float64 `json:"width"`
}
