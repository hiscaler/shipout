package shipout

import (
	"fmt"
	"github.com/hiscaler/gox/jsonx"
	"github.com/hiscaler/shipout-go/constant"
	"testing"
	"time"
)

func TestService_Bulk(t *testing.T) {
	req := BulkOrderRequests{
		{
			OrderNo: "001",
			OrderSummary: BulkOrderSummary{
				OrderDate: time.Now().Format(constant.DatetimeFormat),
			},
			ShipmentForms: []BulkShipmentForm{
				{
					International: BulkShipmentFormInternational{
						EEIType:   2,
						EinOrSsn:  "123",
						ITNNumber: "121212",
					},
					ProductList: []BulkShipmentFormProduct{
						{
							SkuId:    "1521695517976563714",
							Quantity: 1,
							Price:    1,
						},
					},
					OutboundInfo: BulkShipmentFormOutboundInfo{
						SysServiceId:   "6",
						TrackingNumber: "SF001",
						ShipmentOutboundLabel: BulkShipmentFormOutboundLabel{
							LabelURL: "https://www.a.com/sf001.pdf",
						},
					},
					ShippingInfo: BulkShipmentFormShippingInfo{
						ShipDate:      time.Now().Format(constant.DatetimeFormat),
						ShipmentSid:   "1",
						CarrierId:     1,
						SignatureType: 2,
					},
				},
			},
			SID: int(time.Now().Unix()),
			ToAddress: BulkToAddress{
				Name:         "Haroutun Karadjian",
				AddressLine1: "3384 Route 14 Rum Creek Road Stollings WV US",
				CountryCode:  "US",
				City:         "Stollings",
				StateCode:    "WV",
				ZipCode:      "25646",
				Phone:        "+1 123-456-7890 ext. 62214",
			},
			WarehouseId: "1260418082065616897",
		},
	}
	results, err := shipOutClient.OMS.Order.Bulk(req)
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
