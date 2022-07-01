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

## 说明

All() 方法会返回三个值，分别是：

- 获取到的数据
- 是否为最后一页
- 错误信息

您可以使用以下代码来获取所有的数据，例如：
```go
params := OrdersQueryParams{
    CurPageNo: 1,
    PageSize:  100,
}
for {    
    orders, isLastPage, err := shipOutClient.OMS.Order.All(params)
    if err != nil {
        panic(err)	
    }
    // orders 业务处理
    if isLastPage {
        break
    }
    params.CurPageNo += 1
}
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

## 产品

```go
shipOutClient.OMS.Product.Xyz()
```

- All() 产品列表
- Update() 更新商品

## 增值服务

```go
shipOutClient.OMS.ValueAddedService.Xyz()
```

- Warehouses() 仓库增值服务

## 基础信息

```go
shipOutClient.OMS.BaseInfo.Xyz()
```

- Warehouses() 仓库列表