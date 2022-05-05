package order

import (
	"fmt"
	"github.com/hiscaler/gox/jsonx"
	"github.com/hiscaler/shipout-go"
	"github.com/hiscaler/shipout-go/config"
	"github.com/hiscaler/shipout-go/constant"
	jsoniter "github.com/json-iterator/go"
	"os"
	"testing"
	"time"
)

var soInstance *shipout.ShipOut
var soService Service

func TestMain(m *testing.M) {
	b, err := os.ReadFile("../../config/config_test.json")
	if err != nil {
		panic(fmt.Sprintf("Read config error: %s", err.Error()))
	}
	var c config.Config
	err = jsoniter.Unmarshal(b, &c)
	if err != nil {
		panic(fmt.Sprintf("Parse config file error: %s", err.Error()))
	}

	soInstance = shipout.NewShipOut(c)
	soService = NewService(soInstance)
	m.Run()
}

func TestService_BatchSubmit(t *testing.T) {
	req := BatchSubmitOrderRequest{
		{
			OrderNo: "001",
			OrderSummary: OrderSummary{
				OrderDate: time.Now().Format(constant.DatetimeFormat),
			},
			ShipmentForms: []ShipmentForm{
				{
					International: ShipmentFormInternational{
						EEIType:   2,
						EinOrSsn:  "123",
						ITNNumber: "121212",
					},
					ProductList: []ShipmentFormProduct{
						{
							SKUId:    "111",
							Quantity: 1,
							Price:    1,
						},
					},
					ShippingInfo: ShipmentFormShippingInfo{
						ShipDate:      time.Now().Format(constant.DatetimeFormat),
						ShipmentSid:   "1",
						SignatureType: 2,
					},
				},
			},
			SID: int(time.Now().Unix()),
			ToAddress: ToAddress{
				ZipCode: "10010",
			},
			WarehouseId: "1",
		},
	}
	results, err := soService.BatchSubmit(req)
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(jsonx.ToPrettyJson(results))
	}
}

func TestOrders(t *testing.T) {
	params := OrdersQueryParams{
		CurPageNo: 1,
		PageSize:  100,
	}
	body := OrdersQueryBody{
		OrgId:       "",
		ServerOrgId: "",
		Status:      9,
		WarehouseId: "",
	}
	orders, _, err := soService.Orders(params, body)
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(fmt.Sprintf("%#v", orders))
	}
}

func TestOrder(t *testing.T) {
	params := OrderQueryParams{OrderId: "1"}
	order, err := soService.Order(params)
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(fmt.Sprintf("%#v", order))
	}
}
