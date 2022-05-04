package cast

import (
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

func StructToURLValues(i interface{}) url.Values {
	values := url.Values{}
	vo := reflect.ValueOf(i)
	if vo.Kind() != reflect.Struct {
		return values
	}

	t := vo.Type()
	for j := 0; j < vo.NumField(); j++ {
		key := t.Field(j).Tag.Get("json")
		if key == "" {
			continue
		}

		field := vo.Field(j)
		ignoreEmpty := false
		tags := strings.Split(key, ",")
		if len(tags) > 1 {
			key = tags[0]
			if tags[1] == "omitempty" {
				ignoreEmpty = true
			}
		}
		switch field.Kind() {
		case reflect.String:
			v := field.String()
			if !ignoreEmpty || v != "" {
				values.Add(key, v)
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			v := field.Int()
			if !ignoreEmpty || v != 0 {
				values.Add(key, strconv.Itoa(int(v)))
			}
		case reflect.Float32, reflect.Float64:
			v := field.Float()
			if !ignoreEmpty || v != 0 {
				values.Add(key, strconv.FormatFloat(v, 2, -1, 64))
			}
		}
	}
	return values
}
