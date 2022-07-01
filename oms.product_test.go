package shipout

import (
	"fmt"
	"testing"
)

func TestProductAll(t *testing.T) {
	params := ProductsQueryParams{
		CurPageNo: 1,
		PageSize:  100,
	}
	orders, _, err := shipOutClient.OMS.Product.All(params)
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(fmt.Sprintf("%#v", orders))
	}
}
