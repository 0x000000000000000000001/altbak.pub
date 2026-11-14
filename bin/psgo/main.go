package main

import (
    . "github.com/purescript-native/go-runtime"
    _ "project.localhost/purescript-native/ffi-loader"
    benchmark "project.localhost/purescript-native/output/App"
)

func main() { benchmark.Ꞌmain().(func() Any)() }
