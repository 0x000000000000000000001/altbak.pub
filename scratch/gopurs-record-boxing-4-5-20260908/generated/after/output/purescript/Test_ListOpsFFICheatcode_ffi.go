package purescript

import "gopurs/output/gopurs_runtime"



func Test_ListOpsFFICheatcode_RunListOpsFFICheatcode(limit int) int {
	n := int(limit)
	sum := 0
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	return (sum)
}



// --- Auto-generated FFI wrappers ---
var _Gopurs_Test_ListOpsFFICheatcode_RunListOpsFFICheatcode = // TAST: (Func [Int] Int)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_res := Test_ListOpsFFICheatcode_RunListOpsFFICheatcode(go_arg0)
	return gopurs_runtime.Int(int64(go_res))
})