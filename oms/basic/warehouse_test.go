package basic

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
	params := WarehousesQueryParams{}
	warehouses, _, err := soService.Warehouses(params)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fmt.Sprintf("%#v", warehouses))
	}
}
