package shipout

import (
	jsoniter "github.com/json-iterator/go"
)

// 基础信息
// https://open.shipout.com/portal/zh/api/base-info.html

type baseInfoService service

type Warehouse struct {
	OrgId             string `json:"orgId"`             // 仓库机构ID
	TimeZone          string `json:"timeZone"`          // 仓库时区
	WarehouseAddr1    string `json:"warehouseAddr1"`    // 仓库地址1
	WarehouseAddr2    string `json:"warehouseAddr2"`    // 仓库地址2
	WarehouseCity     string `json:"warehouseCity"`     // 仓库所在城市
	WarehouseContacts string `json:"warehouseContacts"` // 仓库联系人
	WarehouseCountry  string `json:"warehouseCountry"`  // 仓库所在国家
	WarehouseEmail    string `json:"warehouseEmail"`    // 仓库联系 Email
	WarehouseId       string `json:"warehouseId"`       // 仓库编号
	WarehouseName     string `json:"warehouseName"`     // 仓库名称
	WarehousePhone    string `json:"warehousePhone"`    // 仓库联系电话
	WarehouseProvince string `json:"warehouseProvince"` // 仓库所在州
	WarehouseZipCode  string `json:"warehouseZipCode"`  // 仓库邮编
}

// Warehouses 仓库列表获取接口
func (s baseInfoService) Warehouses() (items []Warehouse, err error) {
	res := struct {
		NormalResponse
		Data []Warehouse `json:"data"`
	}{}
	resp, err := s.httpClient.R().Get("/open-api/oms/info/warehouse/list")
	if err != nil {
		return
	}

	if err = jsoniter.Unmarshal(resp.Body(), &res); err == nil {
		items = res.Data
	}
	return
}
