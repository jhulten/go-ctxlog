package ctxlog

import (
	"fmt"
	"reflect"
	"strings"
)

type LogFormatter interface {
	Format(v ...interface{}) string
}

type stringFormatter struct{}

func (f *stringFormatter) Format(i interface{}) string {
	v := reflect.ValueOf(i)
	return formatValue(v)
}

func formatValue(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Slice:
		h, t := head(v)
		if t.IsNil() {
			return formatValue(h)
		}
		return strings.Join([]string{formatValue(h), formatValue(t)}, ", ")
	case reflect.String, reflect.Int, reflect.Bool:
		return v.String()
	default:
		return fmt.Sprintf("%v[%#v]", v.Kind(), v.Interface())
	}
}

func head(v reflect.Value) (reflect.Value, reflect.Value) {
	var rest reflect.Value
	if v.Len() > 0 {
		rest = v.Slice(1, v.Len()-1)
	}
	return v.Index(0), rest
}
