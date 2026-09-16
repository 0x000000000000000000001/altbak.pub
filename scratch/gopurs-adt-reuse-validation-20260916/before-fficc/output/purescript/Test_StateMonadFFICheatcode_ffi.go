package purescript

import "gopurs/output/gopurs_runtime"



func Test_StateMonadFFICheatcode_RunStateMonadFFICheatcode(limit int) int {
	// The shared FFI wrapper supplies the depth; evaluate 20 repetitions without closures.
	state := 0
	for i := 0; i < 20; i++ {
		for j := 0; j < limit; j++ {
			state += 1
		}
	}
	return (state)
}


// --- Auto-generated FFI wrappers ---
var _Gopurs_Test_StateMonadFFICheatcode_RunStateMonadFFICheatcode = // TAST: (Func [Int] Int)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_res := Test_StateMonadFFICheatcode_RunStateMonadFFICheatcode(go_arg0)
	return gopurs_runtime.Int(int64(go_res))
})