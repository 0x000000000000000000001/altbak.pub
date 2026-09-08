package purescript

import "gopurs/output/gopurs_runtime"



func fib_cheatcode(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib_cheatcode(n-1) + fib_cheatcode(n-2)
}

func Test_FibFFICheatcode_RunFibFFICheatcode(limit int) int {
	dummy := limit
	return (fib_cheatcode(dummy))
}


// --- Auto-generated FFI wrappers ---
var _Gopurs_Test_FibFFICheatcode_RunFibFFICheatcode = // TAST: (Func [Int] Int)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_res := Test_FibFFICheatcode_RunFibFFICheatcode(go_arg0)
	return gopurs_runtime.Int(int64(go_res))
})