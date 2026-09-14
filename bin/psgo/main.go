package main

import (
    . "github.com/purescript-native/go-runtime"
    _ "project.localhost/purescript-native/ffi-loader"
    "project.localhost/purescript-native/output/App"
)

func main() { App.Ꞌmain().(func() Any)() }
