package Data_Bounded

import "gopurs/output/gopurs_runtime"

import "math"
var TopInt = 2147483647
var BottomInt = -2147483648
var TopChar = string(rune(65535))
var BottomChar = string(rune(0))
var TopNumber = math.Inf(1)
var BottomNumber = math.Inf(-1)


// --- Auto-generated FFI wrappers ---
var _Gopurs_TopInt = gopurs_runtime.Box(TopInt)
var _Gopurs_BottomInt = gopurs_runtime.Box(BottomInt)
var _Gopurs_TopChar = gopurs_runtime.Box(TopChar)
var _Gopurs_BottomChar = gopurs_runtime.Box(BottomChar)
var _Gopurs_TopNumber = gopurs_runtime.Box(TopNumber)
var _Gopurs_BottomNumber = gopurs_runtime.Box(BottomNumber)
