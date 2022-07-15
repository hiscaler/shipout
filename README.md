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
params := OrdersQueryParams{}
params.CurPageNo = 1
params.PageSize = 100
orders, _, err := shipOutClient.OMS.Order.All(params)
fmt.Println(orders)
```

## 方法调用说明

### 带翻页的列表数据

如果是带有翻页的数据，对应的方法会有三个返回值（比如：All()），分别是：

- 获取到的数据
- 是否为最后一页
- 错误信息

您可以使用以下代码来获取所有的数据，例如：
```go
params := OrdersQueryParams{}
params.CurPageNo = 1
params.PageSize = 100
for {    
    orders, isLastPage, err := shipOutClient.OMS.Order.All(params)
    if err != nil {
        panic(err)
    }
    // 根据返回的 orders 进行业务处理
    if isLastPage {
        break
    }
    params.CurPageNo += 1
}
```

需要获取所有值，您可以使用 for 循环调用对应的方法，并且不断的累加 CurPageNo 参数，直到 isLastPage 返回 true 为止。

### 不带翻页的列表数据

如果不带有翻页，则方法调用后只会有两个返回值，分别为数据集和错误。

### 其他

在任何情况下，您都应该首先判断是否返回错误，然后再进行下一步的业务逻辑处理。

## ShipOut OMS API 方法实现

### 订单

- Bulk() 批量提交订单

```go
req := BulkOrderRequest{}
shipOutClient.OMS.Order.Bulk(req)
```

- All() 查询订单列表

```go
parmas := OrdersQueryParams()
orders, isLastPage, err := shipOutClient.OMS.Order.All(params)
```

- One() 查询单个订单

```go
params := OrderQueryParams{}
order, err := shipOutClient.OMS.Order.One(params)
```

- Cancel() 取消订单

```go
id, err := shipOutClient.OMS.Order.Cancel(orderId)
```

## 产品

- Bulk() 批量创建产品

```go
req := BulkProductRequest{}
err := shipOutClient.OMS.Product.Bulk(req)
```

- All() 产品列表

```go
params := ProductsQueryParams{}
products, isLastPage, err := shipOutClient.OMS.Product.All(params)
```

- Update() 更新商品

```go
req := UpdateProductRequest{}
err := shipOutClient.OMS.Product.Update(req)
```

## 增值服务

- Warehouses() 仓库增值服务

```go
warehouseValueAddedServices, err := shipOutClient.OMS.ValueAddedService.Warehouses()
```

## 基础信息

- Warehouses() 仓库列表

```go
warehouses, err := shipOutClient.OMS.BaseInfo.Warehouses()
```