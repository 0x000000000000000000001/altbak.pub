package main

import (
	"gopurs/output/App"
	"gopurs/output/gopurs_runtime"
)

func main() {
	gopurs_runtime.Apply(App.Get_main(), gopurs_runtime.Value{})
}
