package purescript

import "gopurs/output/gopurs_runtime"


func Test_TCOFFICheatcode_RunTCOFFICheatcode(limit int) int {
	n := int(limit)
	acc := 0
	for n > 0 {
		acc += (n % 3)
		n--
	}
	return (acc)
}


// --- Auto-generated FFI wrappers ---
var _Gopurs_Test_TCOFFICheatcode_RunTCOFFICheatcode = // TAST: (Func [Int] Int)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_res := Test_TCOFFICheatcode_RunTCOFFICheatcode(go_arg0)
	return gopurs_runtime.Int(int64(go_res))
})