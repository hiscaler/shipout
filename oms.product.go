package shipout

import (
	"errors"
	"github.com/hiscaler/shipout-go/entity"
	jsoniter "github.com/json-iterator/go"
)

type productService service

type ProductsQueryParams struct {
	Asc         bool     `url:"asc,omitempty"`
	auditStatus []string `url:"audit_status,omitempty"`
	CurPageNo   int      `url:"curPageNo,omitempty"`
	HiDirection string   `url:"hiDirection,omitempty"`
	Name        string   `url:"name,omitempty"`
	omsSku      string   `url:"omsSku,omitempty"`
	OrderColumn string   `url:"orderColumn,omitempty"`
	PageSize    int      `url:"pageSize,omitempty"`
	Status      int      `url:"status,omitempty"`
	Typ         int      `url:"type,omitempty"`
}

func (m ProductsQueryParams) Validate() error {
	return nil
}

func (s productService) All(params ProductsQueryParams) (items []entity.ProductRecord, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	res := struct {
		NormalResponse
		Data struct {
			Current     int                    `json:"current"`
			Pages       int                    `json:"pages"`
			Records     []entity.ProductRecord `json:"records"`
			SearchCount bool                   `json:"searchCount"`
			Size        string                 `json:"size"`
			Total       string                 `json:"total"`
		} `json:"data"`
	}{}

	resp, err := s.httpClient.R().
		SetQueryParamsFromValues(toValues(params)).
		Get("/open-api/oms/product/queryList")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		if err = ErrorWrap(res.ErrorCode, res.Message); err == nil {
			if err = jsoniter.Unmarshal(resp.Body(), &res); err == nil {
				items = res.Data.Records
			}
		}
	} else {
		if e := jsoniter.Unmarshal(resp.Body(), &res); e == nil {
			err = ErrorWrap(res.ErrorCode, res.Message)
		} else {
			err = errors.New(resp.Status())
		}
	}
	return
}
