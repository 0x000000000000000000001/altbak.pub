package main

import "fmt"

type Constructor_T struct {
	Rc uint32
	V0 interface{}
}

func main() {
	var v *Constructor_T = nil
	if v != nil {
		v.V0 = nil
		fmt.Println("Inside if!")
	} else {
		fmt.Println("Inside else!")
	}
}
