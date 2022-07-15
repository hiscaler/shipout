package entity

// WarehouseValueAddedService 仓库增值服务
type WarehouseValueAddedService struct {
	AdServiceId string `json:"adServiceId"` // 增值服务id
	ServiceName string `json:"serviceName"` // 增值服务名称
	Status      int    `json:"status"`      // 1.启用 2.停用 3.删除
	Type        int    `json:"type"`        // 1.产品 2.库存 3.入库 4.出库 5.退货
	WarehouseId string `json:"warehouseId"` // 仓库 ID
}
