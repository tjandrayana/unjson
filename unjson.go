package unjson

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

var ErrNotStruct = errors.New("must pass a struct or pointer to struct")

func MarshalWithoutQuotes(data interface{}) (string, error) {
	bld := strings.Builder{}
	_, err := bld.WriteString("{")
	if err != nil {
		return "", err
	}

	values := reflect.ValueOf(data)
	values = reflect.Indirect(values)
	// prevent panic when call NumField()
	if values.Kind() != reflect.Struct {
		return "", ErrNotStruct
	}

	n := values.NumField()
	for i := 0; i < n; i++ {
		f := values.Type().Field(i)
		t := f.Tag.Get("json")

		// remove omitempty
		field := strings.Split(t, ",")[0]
		if field == "" {
			// use name if no tag
			field = f.Name
		}
		_, err = bld.WriteString(field + ": ")
		if err != nil {
			return "", err
		}

		switch f.Type.String() {
		case "string", "time.Time":
			_, err = bld.WriteString(fmt.Sprintf(`"%v"`, values.Field(i).Interface()))
			if err != nil {
				return "", err
			}
		default:
			_, err = bld.WriteString(fmt.Sprintf(`%v`, values.Field(i).Interface()))
			if err != nil {
				return "", err
			}
		}

		if i < n-1 {
			_, err = bld.WriteString(", ")
			if err != nil {
				return "", err
			}
		}
	}

	_, err = bld.WriteString("}")
	if err != nil {
		return "", err
	}

	return bld.String(), nil
}
