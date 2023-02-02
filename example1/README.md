## Example1

Sometimes with certain layout of the struct fields, depending on their order and type, calling the exported `read_buffer` function fails (with `panic: (reflect.Value).Interface: unexported`). See the comments in `main.go`

`make`
