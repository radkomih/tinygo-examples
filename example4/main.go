package main

import (
	"reflect"
)

func main() {
	var value interface{} = uint16(0)

	// get the pointer to the interface value
	inPtr := reflect.ValueOf(&value)

	// dereference to get the actual value (an interface)
	inElem := inPtr.Elem()

	// create a new value of the same concrete type
	uint16Type := reflect.TypeOf(uint16(0))
	newUint16Value := reflect.New(uint16Type).Elem()
	newUint16Value.Set(reflect.ValueOf(uint16(13)))

	// set the new value to the interface
	inElem.Set(newUint16Value)

	println(value.(uint16))
}
