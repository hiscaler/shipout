package shipout

import (
	"fmt"
	"testing"
)

func TestProductAll(t *testing.T) {
	params := ProductsQueryParams{}
	params.CurPageNo = 1
	params.PageSize = 100
	orders, _, err := shipOutClient.OMS.Product.All(params)
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(fmt.Sprintf("%#v", orders))
	}
}
