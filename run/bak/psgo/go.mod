module project.localhost/example

go 1.27.0

replace project.localhost/purescript-native/ffi-loader => ./purescript-native

replace project.localhost/purescript-native/output => ./output

replace github.com/purescript-native/go-ffi => github.com/joprice/purescript-native-go-ffi v0.0.0-20230319170919-df001aa880fd

require (
	github.com/dlclark/regexp2 v1.4.0 // indirect
	github.com/i-am-the-slime/go-ffi v0.0.0-20260725162152-03a066021db8 // indirect
	github.com/purescript-native/go-runtime v0.1.2 // indirect
	project.localhost/purescript-native/ffi-loader v0.0.0-00010101000000-000000000000 // indirect
	project.localhost/purescript-native/output v0.0.0-00010101000000-000000000000 // indirect
)
