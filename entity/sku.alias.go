package entity

// SkuAlias sku 别名
type SkuAlias struct {
	SkuId          string `json:"skuId"`
	Sku            string `json:"sku"`
	StoreAliasesId string `json:"storeAliasesId"`
	StoreId        string `json:"storeId"`
}
