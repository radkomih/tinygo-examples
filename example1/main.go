package main

import (
	"bytes"
	"reflect"
)

// custom defined types, each having an Encode method attached

type U uint

func (v U) Encode(buffer *bytes.Buffer) {
	buffer.Write([]byte{0x00, 0x00})
}

type I int

func (v I) Encode(buffer *bytes.Buffer) {
	buffer.Write([]byte{0x01, 0x00})
}

type U8 uint8

func (v U8) Encode(buffer *bytes.Buffer) {
	buffer.Write([]byte{0x00, 0x08})
}

type I8 int8

func (v I8) Encode(buffer *bytes.Buffer) {
	buffer.Write([]byte{0x01, 0x08})
}

type U16 uint16

func (v U16) Encode(buffer *bytes.Buffer) {
	buffer.Write([]byte{0x00, 0x10})
}

type I16 int16

func (v I16) Encode(buffer *bytes.Buffer) {
	buffer.Write([]byte{0x01, 0x10})
}

type U32 uint32

func (v U32) Encode(buffer *bytes.Buffer) {
	buffer.Write([]byte{0x00, 0x20})
}

type I32 int32

func (v I32) Encode(buffer *bytes.Buffer) {
	buffer.Write([]byte{0x01, 0x20})
}

type U64 uint64

func (v U64) Encode(buffer *bytes.Buffer) {
	buffer.Write([]byte{0x00, 0x40})
}

type I64 int64

func (v I64) Encode(buffer *bytes.Buffer) {
	buffer.Write([]byte{0x01, 0x40})
}

type TestU8 struct {
	A U8
	B U8
	C U8
	D U8
	E U8 // fails if this additional field is added to the struct
}

type TestI8 struct {
	A I8
	B I8
	C I8
	D I8
	// E I8 // fails if this additional field is added to the struct
}

type TestU16 struct {
	A U16
	B U16
	// C U16 // fails if this additional field is added to the struct
}

type TestI16 struct {
	A I16
	B I16
	// C I16 // fails if this additional field is added to the struct
}

type TestU32 struct {
	A U32
	// B U32 // fails if this additional field is added to the struct
}

type TestI32 struct {
	A I32
	// B I32 // fails if this additional field is added to the struct
}

type TestU64 struct {
	A U64
	B U64
	C U64
	D U64
	E U64
}

type TestI64 struct {
	A I64
	B I64
	C I64
	D I64
	E I64
}

type TestU struct {
	A U
	// B U // fails if this additional field is added to the struct
}

type TestI struct {
	A I
	// B I // fails if this additional field is added to the struct
}

type TestU8U16 struct {
	A U8
	B U16
	// C U8 // fails if this additional field is added to the struct
}

func encode(value interface{}, buffer *bytes.Buffer) {
	t := reflect.TypeOf(value)
	v := reflect.ValueOf(value)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		if t.Field(i).IsExported() {
			switch field.Kind() {
			case reflect.Int8:
				(field.Interface().(I8)).Encode(buffer)
			case reflect.Uint8:
				(field.Interface().(U8)).Encode(buffer)
			case reflect.Int16:
				(field.Interface().(I16)).Encode(buffer)
			case reflect.Uint16:
				(field.Interface().(U16)).Encode(buffer)
			case reflect.Int32:
				(field.Interface().(I32)).Encode(buffer)
			case reflect.Uint32:
				(field.Interface().(U32)).Encode(buffer)
			case reflect.Int64:
				(field.Interface().(I64)).Encode(buffer)
			case reflect.Uint64:
				(field.Interface().(U64)).Encode(buffer)
			case reflect.Uint:
				(field.Interface().(U)).Encode(buffer)
			case reflect.Int:
				(field.Interface().(I)).Encode(buffer)
			default:
				panic("unreachable case")
			}
		}
	}
}

//export read_buffer
func readBuffer(i int) int {
	buffer := bytes.Buffer{}
	t := TestU8{}
	encode(t, &buffer)
	return int(buffer.Bytes()[i])
}

func main() {}
