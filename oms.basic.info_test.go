package shipout

import (
	"fmt"
	"github.com/hiscaler/gox/jsonx"
	"testing"
)

func TestWarehouses(t *testing.T) {
	warehouses, err := shipOutClient.OMS.BaseInfo.Warehouses()
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(jsonx.ToPrettyJson(warehouses))
	}
}
