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

func TestOrderAll(t *testing.T) {
	params := OrdersQueryParams{
		CurPageNo: 1,
		PageSize:  100,
	}
	orders, _, err := shipOutClient.OMS.Order.All(params)
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(fmt.Sprintf("%#v", orders))
	}
}

func TestOrderOne(t *testing.T) {
	params := OrderQueryParams{OrderId: "1521695670086434817"}
	order, err := shipOutClient.OMS.Order.One(params)
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(fmt.Sprintf("%#v", order))
	}
}
