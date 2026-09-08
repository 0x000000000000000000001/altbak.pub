package purescript

import "gopurs/output/gopurs_runtime"



func Test_ChurchFFICheatcode_RunChurchFFICheatcode(limit int) int {
	// Native Church numerals are essentially functional loops.
	// We optimize it as a native loop for the FFICheatcode.
	n := int(limit)
	acc := 0
	for i := 0; i < n * 10000; i++ {
		acc++
	}
	return (acc)
}



// --- Auto-generated FFI wrappers ---
var _Gopurs_Test_ChurchFFICheatcode_RunChurchFFICheatcode = // TAST: (Func [Int] Int)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_res := Test_ChurchFFICheatcode_RunChurchFFICheatcode(go_arg0)
	return gopurs_runtime.Int(int64(go_res))
})