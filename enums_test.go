package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/rvIceBreaker/goclack/keycodes"
)

func TestEnumReflection(t *testing.T) {
	var structToMap func(interface{}) (map[string]interface{})
	structToMap = func(obj interface{}) (map[string]interface{}) {
		res := make(map[string]interface{})
		val := reflect.ValueOf(obj)

		typ := val.Type()

		for i:= 0; i < val.NumField(); i++ {
			fieldName := typ.Field(i).Name
			fieldKind := val.Field(i).Kind()
			var fieldVal interface{}

			if fieldKind == reflect.Struct {
				fieldVal = structToMap(val.Field(i).Interface())
			} else {
				fieldVal = val.Field(i).Interface()
			}

			res[fieldName] = fieldVal
		}

		return res
	}

	m := structToMap(*keycodes.VKeys)
	fmt.Println(m)
}