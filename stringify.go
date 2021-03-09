package stringify

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type String interface {
	String() string
}

func ToString(i interface{}, subs ...string) string {

	valueElement := getReflectValue(i)
	valueTypeKind := valueElement.Type().Kind()

	if e, ok := valueElement.Interface().(String); ok {
		return e.String()
	}

	if valueTypeKind == reflect.String {
		return valueElement.String()
	}

	if valueTypeKind == reflect.Bool {
		if valueElement.Bool() == true {
			return "true"
		}
		return "false"
	}

	if valueTypeKind <= reflect.Int64 {
		return strconv.FormatInt(valueElement.Int(), 10)
	}

	if valueTypeKind <= reflect.Uintptr {
		return strconv.FormatUint(valueElement.Uint(), 10)
	}

	if valueTypeKind <= reflect.Float64 {
		if len(subs) > 0 && subs[0] != "" {
			f, _ := strconv.ParseInt(subs[0], 10, 64)
			return strconv.FormatFloat(valueElement.Float(), 'f', int(f), 64)
		}
		return fmt.Sprintf("%v", valueElement.Float())
	}

	if valueTypeKind == reflect.Slice {

		if _, ok := valueElement.Interface().([]string); !ok {
			return ""
		}
		sub := ","
		if len(subs) > 0 && subs[0] != "" {
			sub = subs[0]
		}
		return strings.Join((valueElement.Slice(0, valueElement.Cap()).Interface()).([]string), sub)
	}

	return ""
}

func getReflectValue(value interface{}) reflect.Value {
	if e, ok := value.(reflect.Value); ok {
		return e
	} else {
		valueElement := reflect.ValueOf(value)
		valueElement = reflect.ValueOf(valueElement.Interface())
		valueTypeKind := valueElement.Type().Kind()
		if valueTypeKind == reflect.Ptr {
			return valueElement.Elem()
		} else {
			return valueElement
		}
	}
}
