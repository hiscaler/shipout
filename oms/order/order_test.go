package order

import (
	"fmt"
	"github.com/hiscaler/shipout-go"
	"github.com/hiscaler/shipout-go/config"
	jsoniter "github.com/json-iterator/go"
	"os"
	"testing"
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
