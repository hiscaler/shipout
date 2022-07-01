package shipout

import (
	"fmt"
	"testing"
)

func TestOmsValueAddedServiceWarehouses(t *testing.T) {
	params := WarehouseValueAddedServicesQueryParams{
		WarehouseId: "1260418082065616897",
	}
	orders, _, err := shipOutClient.OMS.ValueAddedService.Warehouses(params)
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(fmt.Sprintf("%#v", orders))
	}
}
