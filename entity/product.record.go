package entity

type ProductRecordSKUAlias struct {
	OmsSKU  string `json:"omsSku"`  // 系统SKU
	SKU     string `json:"sku"`     // 店铺SKU
	StoreId string `json:"storeId"` // 店铺Id
}

// ProductRecordBundleDetail 组合详情
type ProductRecordBundleDetail struct {
	ParentSkuId   string `json:"parentSkuId"`   // 系统SKU
	ChildSkuId    string `json:"childSkuId"`    // 子SKU
	ChildQuantity int    `json:"childQuantity"` // 子SKU在组合中的数量
	Note          string `json:"note"`          // 组装信息
}

type ProductRecordWmsAllStatus struct {
	SkuId     string `json:"skuId"`
	AllStatus int    `json:"allStatus"` // 审核状态：1-Approval Pending，2-Approved，3-Rejected，4-Partial Approved，5-Partial Rejected
}

// ProductRecordWmsSkuBundleDetail
type ProductRecordWmsSkuBundleDetail struct {
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

type ProductRecordWmsSkuStockDetailWarehouseStock struct {
	AvailableQty   int    `json:"availableQty"`
	BrokenQty      int    `json:"brokenQty"`
	InboundingQty  int    `json:"inboundingQty"`
	OutboundingQty int    `json:"outboundingQty"`
	SkuId          string `json:"skuId"`
	StandardQty    int    `json:"standardQty"`
	WarehouseId    string `json:"warehouseId"`
	WarehouseName  string `json:"warehouseName"`
}

// ProductRecordWmsSkuStockDetail todo
type ProductRecordWmsSkuStockDetail struct {
	AvailableQty       int                                            `json:"availableQty"`
	BrokenQty          int                                            `json:"brokenQty"`
	DailyEstimateFee   float64                                        `json:"dailyEstimateFee"`
	InboundingQty      int                                            `json:"inboundingQty"`
	OutboundingQty     int                                            `json:"outboundingQty"`
	SkuId              string                                         `json:"skuId"`
	StandardQty        int                                            `json:"standardQty"`
	TotalStorageFee    float64                                        `json:"totalStorageFee"`
	WarehouseStockList []ProductRecordWmsSkuStockDetailWarehouseStock `json:"warehouseStockList"`
}

// ProductRecordWmsHistory 历史数据
type ProductRecordWmsHistory struct {
	Activity        string `json:"activity"`
	Operator        string `json:"operator"`
	OperatorId      string `json:"operatorId"`
	SkuId           string `json:"skuId"`
	Status          int    `json:"status"`
	UpdateTime      string `json:"updateTime"`
	WarehouseId     string `json:"warehouseId"`
	WarehouseName   string `json:"warehouseName"`
	WmsSkuHistoryId string `json:"wmsSkuHistoryId"`
	WmsSkuId        string `json:"wmsSkuId"`
}

// ProductRecordWmsSku 仓库
type ProductRecordWmsSku struct {
	Aliases               []SKUAlias                        `json:"aliases"`
	AllStatus             []ProductRecordWmsAllStatus       `json:"allStatus"`
	Asin                  string                            `json:"asin"`
	AuditRemark           string                            `json:"auditRemark"`
	AuditStatus           int                               `json:"auditStatus"`
	BatteryFlag           bool                              `json:"batteryFlag"`
	BatteryFlagStr        string                            `json:"batteryFlagStr"`
	BatteryType           string                            `json:"batteryType"`
	Brand                 string                            `json:"brand"`
	BundleDetail          []ProductRecordWmsSkuBundleDetail `json:"bundleDetail"`
	CaseDistanceUnit      string                            `json:"caseDistanceUnit"`
	CaseHeight            float64                           `json:"caseHeight"`
	CaseLength            float64                           `json:"caseLength"`
	CaseMassUnit          string                            `json:"caseMassUnit"`
	CaseWeight            float64                           `json:"caseWeight"`
	CaseWeight1           float64                           `json:"caseWeight1"`
	CaseWeight2           float64                           `json:"caseWeight2"`
	CaseWidth             float64                           `json:"caseWidth"`
	CreateTime            string                            `json:"createTime"`
	CustDescription       string                            `json:"custDescription"`
	DeclaredValue         float64                           `json:"declaredValue"`
	Dimension             string                            `json:"dimension"`
	DistanceUnit          string                            `json:"distanceUnit"`
	Ean                   string                            `json:"ean"`
	ExtraBarcode1         string                            `json:"extraBarcode1"`
	FnSku                 string                            `json:"fnSku"`
	HasOmsSkuHistory      string                            `json:"hasOmsSkuHistory"`
	Height                float64                           `json:"height"`
	IndividualInventory   int                               `json:"individualInventory"` // 1.仅有独立库存 2.独立库存+非独立库存
	InnerQuantity         int                               `json:"innerQuantity"`
	InventoryThreshold    int                               `json:"inventoryThreshold"`
	Length                float64                           `json:"length"`
	LsaQuantity           int                               `json:"lsaQuantity"`
	MassUnit              string                            `json:"massUnit"`
	MaterialName          string                            `json:"materialName"` // 包材名称
	Note                  string                            `json:"note"`
	OmsSku                string                            `json:"omsSku"`
	OrgId                 string                            `json:"orgId"`
	OrgName               string                            `json:"orgName"`
	OriginCountry         string                            `json:"originCountry"`
	PackagingMaterialId   string                            `json:"packagingMaterialId"` // 包材ID
	PackagingMaterialName string                            `json:"packagingMaterialName"`
	PicFileId             string                            `json:"picFileId"`
	PicUrl                string                            `json:"picUrl"`
	PickingLsaQuantity    int                               `json:"pickingLsaQuantity"`
	PurchasingCost        float64                           `json:"purchasingCost"`
	QtyInOnePackage       int                               `json:"qtyInOnePackage"`
	QuantityUnit          string                            `json:"quantityUnit"`
	ReceiveFirstFlag      bool                              `json:"receiveFirstFlag"`
	ScheduleB             string                            `json:"scheduleB"`
	SerialNumberFlag      bool                              `json:"serialNumberFlag"`
	SerialNumberFlagStr   string                            `json:"serialNumberFlagStr"`
	ShippingType          int                               `json:"shippingType"`
	ShortCode             string                            `json:"shortCode"`
	SkuCartonId           string                            `json:"skuCartonId"`
	SkuCustomDetailsId    string                            `json:"skuCustomDetailsId"`
	SkuId                 string                            `json:"skuId"`
	SkuNameCN             string                            `json:"skuNameCN"`
	SkuNameEN             string                            `json:"skuNameEN"`
	SpuId                 string                            `json:"spuId"`
	Status                int                               `json:"status"`
	StockDetail           ProductRecordWmsSkuStockDetail    `json:"stockDetail"`
	Typ                   int                               `json:"type"`
	Unit                  string                            `json:"unit"`
	Upc                   string                            `json:"upc"`
	UpdateTime            string                            `json:"updateTime"`
	ValueCurrency         string                            `json:"valueCurrency"`
	VariantName           string                            `json:"variantName"`
	WarehouseId           string                            `json:"warehouseId"`
	WarehouseIds          []string                          `json:"warehouseIds"`
	WarehouseName         string                            `json:"warehouseName"`
	WarehouseQuantity     int                               `json:"warehouseQuantity"`
	Weight                float64                           `json:"weight"`
	Weight1               float64                           `json:"weight1"`
	Weight2               float64                           `json:"weight2"`
	WeightDisp            string                            `json:"weightDisp"`
	Width                 float64                           `json:"width"`
	WmsHistories          []string                          `json:"wmsHistories"`
	WmsOrgId              string                            `json:"wmsOrgId"`
	WmsSkuCartonId        string                            `json:"wmsSkuCartonId"`
	WmsSkuId              string                            `json:"wmsSkuId"`
	WmsSkus               []string                          `json:"wmsSkus"`
}

type ProductRecord struct {
	ASIN                string                      `json:"asin"`                // ASIN 码
	Aliases             []ProductRecordSKUAlias     `json:"aliases"`             // 店铺别名详情
	AuditRemark         string                      `json:"auditRemark"`         // 审核备注
	BatteryFlag         bool                        `json:"batteryFlag"`         // 是否含电池
	BatteryType         string                      `json:"batteryType"`         // 电池类型，含电池时必传. 1.Lithium Ion-Cells or batteries ONLY;2.Lithium Ion-Packed with Equipment;3.Lithium Ion-Contained in Equipment4.Lithium Metal-Cells or batteries ONLY5.Lithium Metal-Packed with Equipment6.Lithium Metal-Contained in Equipment
	Brand               string                      `json:"brand"`               // 品牌
	BundleDetail        []ProductRecordBundleDetail `json:"bundleDetail"`        // 组合详情
	CustDescription     string                      `json:"custDescription"`     // 产品海关描述
	DeclaredValue       float64                     `json:"declaredValue"`       // 货值
	DistanceUnit        string                      `json:"distanceUnit"`        // 长度单位，仅支持“in”，“cm”
	Ean                 string                      `json:"ean"`                 // EAN码
	ExtraBarcode1       string                      `json:"extraBarcode1"`       // 额外编码
	FnSku               string                      `json:"fnSku"`               // FNSKU码
	Height              float64                     `json:"height"`              // 高度
	IndividualInventory int                         `json:"individualInventory"` // 1.仅有独立库存 2.独立库存+非独立库存
	Length              float64                     `json:"length"`              // 长度
	MassUnit            string                      `json:"massUnit"`            // 重量单位，仅支持“lb”，“kg”
	Note                string                      `json:"note"`                // 备注
	OmsSku              string                      `json:"omsSku"`              // 系统SKU（用户自定义）
	OrgId               string                      `json:"orgId"`               // 客户机构号
	OriginCountry       string                      `json:"originCountry"`       // 原产地国家，格式标准遵循ISO 3166-1 alpha-2
	PurchasingCost      float64                     `json:"purchasingCost"`      // 采购金额
	QtyInOnePackage     int                         `json:"qtyInOnePackage"`     // 单包裹产品数量,运输类型为直接运输有效
	QuantityUnit        string                      `json:"quantityUnit"`        // 数量单位（来自于承运商要求）
	ScheduleB           string                      `json:"scheduleB"`           // 海关协调码
	ShippingType        int                         `json:"shippingType"`        // 运输类型 1.直接运输 2.需要仓库另外打包
	SkuId               string                      `json:"skuId"`               // SKU ID
	SkuNameCN           string                      `json:"skuNameCN"`           // 产品中文名
	SkuNameEN           string                      `json:"skuNameEN"`           // 产品英文名
	Status              int                         `json:"status"`              // 状态 1.待审核 2.已审核 3.已归档
	Typ                 int                         `json:"type"`                // 类型 1.单个产品 2.组合产品
	Upc                 string                      `json:"upc"`                 // UPC 码
	ValueCurrency       string                      `json:"valueCurrency"`       // 货值币种，格式标准遵循ISO 4217
	WarehouseIds        []string                    `json:"warehouseIds"`        // 新增/修改 产品时选择的仓库列表
	Weight1             float64                     `json:"weight1"`             // 主重量，根据重量单位的不同，展示kg或者是lb
	Weight2             float64                     `json:"weight2"`             // 次重量，根据重量单位的不同，展示g或者是oz，与主重量相加即为总重量
	Width               float64                     `json:"width"`               // 宽度
}
