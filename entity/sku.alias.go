package entity

// SKUAlias SKU 别名
type SKUAlias struct {
	SKUId          string `json:"skuId"`
	SKU            string `json:"sku"`
	StoreAliasesId string `json:"storeAliasesId"`
	StoreId        string `json:"storeId"`
}
