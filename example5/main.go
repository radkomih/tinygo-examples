package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Mod interface {
	Id() int
	Fn() error
}

func main() {
	v := (*Mod)(nil)
	t := reflect.TypeOf(v).Elem()
	num := t.NumMethod()

	fmt.Println("number of methods:" + strconv.Itoa(num))
	for i := 0; i < num; i++ {
		fmt.Println("method name:", t.Method(i).Name)
	}
}
