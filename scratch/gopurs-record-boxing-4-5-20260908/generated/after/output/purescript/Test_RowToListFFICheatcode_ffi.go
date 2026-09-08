package purescript

import "gopurs/output/gopurs_runtime"


func Test_RowToListFFICheatcode_RunRowToListFFICheatcode(limit int) int {
	return 5
}


// --- Auto-generated FFI wrappers ---
var _Gopurs_Test_RowToListFFICheatcode_RunRowToListFFICheatcode = // TAST: (Func [Int] Int)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_res := Test_RowToListFFICheatcode_RunRowToListFFICheatcode(go_arg0)
	return gopurs_runtime.Int(int64(go_res))
})