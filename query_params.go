package shipout

import (
	"github.com/google/go-querystring/query"
	"net/url"
	"strings"
)

const (
	SortAsc  = "asc"
	SortDesc = "desc"
)

type queryParams struct {
	CurPageNo   int    `url:"curPageNo,omitempty"`
	PageSize    int    `url:"pageSize,omitempty"`
	OrderColumn string `url:"orderColumn,omitempty"`
	HiDirection string `url:"hiDirection,omitempty"`
}

func (q *queryParams) TidyVars() *queryParams {
	if q.CurPageNo <= 0 {
		q.CurPageNo = 1
	}
	if q.PageSize <= 0 {
		q.PageSize = 20
	} else if q.PageSize > 500 {
		q.PageSize = 500
	}
	if q.HiDirection == "" {
		q.HiDirection = SortDesc
	} else {
		q.HiDirection = strings.ToLower(q.HiDirection)
		if q.HiDirection != SortDesc {
			q.HiDirection = SortAsc
		}
	}
	return q
}

// change to url.values
func toValues(i interface{}) (values url.Values) {
	values, _ = query.Values(i)
	return
}
