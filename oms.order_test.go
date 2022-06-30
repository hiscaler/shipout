package shipout

import (
	"fmt"
	"github.com/hiscaler/gox/jsonx"
	"github.com/hiscaler/shipout-go/constant"
	"testing"
	"time"
)

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
	results, err := shipOutClient.OMS.Order.BatchSubmit(req)
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
	orders, _, err := shipOutClient.OMS.Order.Orders(params, body)
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(fmt.Sprintf("%#v", orders))
	}
}

func TestOrder(t *testing.T) {
	params := OrderQueryParams{OrderId: "1522028289232097282"}
	order, err := shipOutClient.OMS.Order.Order(params)
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(fmt.Sprintf("%#v", order))
	}
}
