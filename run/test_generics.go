package main

import "fmt"
import "unsafe"

type Constructor[T any] struct {
	Rc uint32
	V0 interface{}
}

func Coerce[T any](val interface{}) *T {
	return (*T)(unsafe.Pointer(&val)) // dummy
}

func Call[T any](dict *Constructor[T]) interface{} {
	return dict.V0
}

func Get() interface{} {
	return func(dict_box interface{}) interface{} {
		// Use `interface{}` instead of `any`, without explicit instantiation on Call
		return Call(Coerce[Constructor[interface{}]](dict_box))
	}
}

func main() {
	fmt.Println("Compiles!")
}
