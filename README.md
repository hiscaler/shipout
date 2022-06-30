ShipOut SDK for golang
======================

## 文档

### OpenAPI

https://open.shipout.com/portal/documentation/

### 签名

https://shenyu.apache.org/zh/docs/plugin-center/authority-and-certification/sign-plugin#%E9%89%B4%E6%9D%83%E4%BD%BF%E7%94%A8%E6%8C%87%E5%8D%97

## ShipOut OMS API

### Base Info

- Warehouses(params WarehousesQueryParams) (items []Warehouse, isLastPage bool, err error) // 仓库列表

### Order

- BatchSubmit(req BatchSubmitOrderRequest) (items []BatchSubmitResult, err error)                          // 批量订单提交
- Orders(params OrdersQueryParams, body OrdersQueryBody) (items []OrderRecord, isLastPage bool, err error) // 订单列表
- Order(params OrderQueryParams) (item Order, err error)         