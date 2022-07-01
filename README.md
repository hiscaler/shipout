ShipOut SDK for golang
======================

## 文档

### OpenAPI

https://open.shipout.com/portal/documentation/

### 签名

https://shenyu.apache.org/zh/docs/plugin-center/authority-and-certification/sign-plugin#%E9%89%B4%E6%9D%83%E4%BD%BF%E7%94%A8%E6%8C%87%E5%8D%97

## 使用
```go
b, err := os.ReadFile("./config/config_test.json")
if err != nil {
    panic(fmt.Sprintf("Read config error: %s", err.Error()))
}
var c config.Config
err = jsoniter.Unmarshal(b, &c)
if err != nil {
    panic(fmt.Sprintf("Parse config file error: %s", err.Error()))
}

shipOutClient := NewShipOut(c)
params := OrdersQueryParams{
    CurPageNo: 1,
    PageSize:  100,
}
orders, _, err := shipOutClient.OMS.Order.All(params)
fmt.Println(orders)
```

## ShipOut OMS API

### 订单

```go
shipOutClient.OMS.Order.Xyz()
```

- Bulk() 批量提交订单
- All() 查询订单列表
- One() 查询单个订单
- Cancel() 取消订单

## 基础信息

```go
shipOutClient.OMS.BaseInfo.Xyz()
```

- Warehouses() 仓库列表