package shipout

import (
	"fmt"
	"github.com/hiscaler/gox/jsonx"
	"testing"
)

func TestWarehouses(t *testing.T) {
	params := WarehousesQueryParams{}
	warehouses, _, err := shipOutClient.OMS.BaseInfo.Warehouses(params)
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(jsonx.ToPrettyJson(warehouses))
	}
}
