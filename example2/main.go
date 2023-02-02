package main

type Foo struct {
	Bar func() any
}

type Baz struct{}

func main() {
	foo := Foo{Bar: func() any { return &Baz{} }}
	_ = foo.Bar().(*Baz)
}
