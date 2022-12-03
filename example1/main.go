package main

import (
	"reflect"
)

type TestU8 struct {
	A uint8
	B uint8
	C uint8
	D uint8
	E uint8 // fails if this additional field is added to the struct
}

//go:noinline
func check(value interface{}) {
	t := reflect.TypeOf(value)
	v := reflect.ValueOf(value)

	println("num fields: ", v.NumField())

	for i := 0; i < v.NumField(); i++ {
		println("i: ", i)

		field := v.Field(i)

		if t.Field(i).IsExported() {
			field.Interface()
		}
	}
}

func main() {
	check(TestU8{})
}
