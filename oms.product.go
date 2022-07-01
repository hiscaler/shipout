package shipout

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hiscaler/shipout-go/entity"
	jsoniter "github.com/json-iterator/go"
)

// 产品
// https://open.shipout.com/portal/zh/api/product.html#

type productService service

// 更新产品

type UpdateProductRequest struct {
	Aliases             []entity.SKUAlias `json:"aliases"`             // 店铺SKU对照，只支持覆盖式的更新
	Asin                string            `json:"asin"`                // ASIN码
	BatteryFlag         bool              `json:"batteryFlag"`         // 是否包含电池
	BatteryType         string            `json:"batteryType"`         // 电池类型，含电池时必传
	Brand               string            `json:"brand"`               // 品牌
	CustDescription     string            `json:"custDescription"`     // 产品海关描述
	DeclaredValue       float64           `json:"declaredValue"`       // 货值
	DistanceUnit        string            `json:"distanceUnit"`        // 长度单位，仅支持“in”，“cm”，有仓库已审核的不可改
	Ean                 string            `json:"ean"`                 // EAN码
	ExtraBarcode1       string            `json:"extraBarcode1"`       // 额外编码
	FnSku               string            `json:"fnSku"`               // FNSKU码
	Height              float64           `json:"height"`              // 高度，有仓库已审核的不可改
	IndividualInventory int               `json:"individualInventory"` // 1.仅有独立库存 2.独立库存+非独立库存
	Length              float64           `json:"length"`              // 长度，有仓库已审核的不可改
	MassUnit            string            `json:"massUnit"`            // 重量单位，仅支持“lb”，“kg”，有仓库已审核的不可改
	OmsSku              string            `json:"omsSku"`              // 系统显示的SKU
	OriginCountry       string            `json:"originCountry"`       // 原产地国家，格式标准遵循ISO 3166-1 alpha-2
	PurchasingCost      float64           `json:"purchasingCost"`      // 采购金额
	QtyInOnePackage     int               `json:"qtyInOnePackage"`     // 单包裹产品数量,运输类型为直接运输有效
	QuantityUnit        string            `json:"quantityUnit"`        // 数量单位（来自于承运商要求）
	ScheduleB           string            `json:"scheduleB"`           // 海关协调码
	ShippingType        int               `json:"shippingType"`        // 运输类型 1.直接运输 2.需要仓库另外打包，有仓库已审核的不可改
	SkuId               string            `json:"skuId"`               // 系统的产品主键
	SkuNameCN           string            `json:"skuNameCN"`           // 商品中文名
	SkuNameEN           string            `json:"skuNameEN"`           // 商品英文名
	Typ                 int               `json:"type"`                // 类型 1.单个产品 2.组合产品
	Upc                 string            `json:"upc"`                 // UPC 码
	WarehouseIds        []string          `json:"warehouseIds"`        // String型数组，每一条记录均为仓库的warehouseId，不需要传已有的warehouseId，所有传递的会与已审核过的判断，新的仓库进行审核
	Weight1             float64           `json:"weight1"`             // 重量1 (KG,LB)，有仓库已审核的不可改
	Weight2             float64           `json:"weight2"`             // 重量2(g,oz)，有仓库已审核的不可改
	Width               float64           `json:"width"`               // 宽度，有仓库已审核的不可改
}

func (m UpdateProductRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.DistanceUnit,
			validation.Required.Error("长度单位不能为空"),
			validation.In("in", "cm").Error("无效的长度单位"),
		),
		validation.Field(&m.Height, validation.Required.Error("高度不能为空")),
		validation.Field(&m.Length, validation.Required.Error("长度不能为空")),
		validation.Field(&m.MassUnit,
			validation.Required.Error("重量单位不能为空"),
			validation.In("lb", "kg").Error("无效的重量单位"),
		),
		validation.Field(&m.OmsSku, validation.Required.Error("系统 SKU 不能为空")),
		validation.Field(&m.ScheduleB, validation.Required.Error("海关协调码不能为空")),
		validation.Field(&m.ShippingType,
			validation.Required.Error("运输类型不能为空"),
			validation.In(1, 2).Error("无效的运输类型"),
		),
		validation.Field(&m.SkuId, validation.Required.Error("系统主键不能为空")),
		validation.Field(&m.SkuNameEN, validation.Required.Error("商品英文名不能为空")),
		validation.Field(&m.Typ,
			validation.Required.Error("类型不能为空"),
			validation.In(1, 2).Error("无效的类型"),
		),
		validation.Field(&m.Weight1, validation.Required.Error("重量1不能为空")),
		validation.Field(&m.Weight2, validation.Required.Error("重量2不能为空")),
		validation.Field(&m.Width, validation.Required.Error("宽度不能为空")),
	)
}

func (s productService) Update(req UpdateProductRequest) error {
	var err error
	if err = req.Validate(); err != nil {
		return err
	}

	res := struct {
		NormalResponse
	}{}

	resp, err := s.httpClient.R().
		SetBody(req).
		Get("/open-api/oms/product/modify")
	if err != nil {
		return err
	}

	if resp.IsSuccess() {
		err = ErrorWrap(res.ErrorCode, res.Message)
	} else {
		if e := jsoniter.Unmarshal(resp.Body(), &res); e == nil {
			err = ErrorWrap(res.ErrorCode, res.Message)
		} else {
			err = errors.New(resp.Status())
		}
	}
	return err
}

// 产品列表获取

type ProductsQueryParams struct {
	Asc         bool     `url:"asc,omitempty"`
	auditStatus []string `url:"audit_status,omitempty"`
	CurPageNo   int      `url:"curPageNo,omitempty"`
	HiDirection string   `url:"hiDirection,omitempty"`
	Name        string   `url:"name,omitempty"`
	omsSku      string   `url:"omsSku,omitempty"`
	OrderColumn string   `url:"orderColumn,omitempty"`
	PageSize    int      `url:"pageSize,omitempty"`
	Status      int      `url:"status,omitempty"`
	Typ         int      `url:"type,omitempty"`
}

func (m ProductsQueryParams) Validate() error {
	return nil
}

func (s productService) All(params ProductsQueryParams) (items []entity.ProductRecord, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	res := struct {
		NormalResponse
		Data struct {
			Current     int                    `json:"current"`
			Pages       int                    `json:"pages"`
			Records     []entity.ProductRecord `json:"records"`
			SearchCount bool                   `json:"searchCount"`
			Size        string                 `json:"size"`
			Total       string                 `json:"total"`
		} `json:"data"`
	}{}

	resp, err := s.httpClient.R().
		SetQueryParamsFromValues(toValues(params)).
		Get("/open-api/oms/product/queryList")
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
