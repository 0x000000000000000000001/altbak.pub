package purescript

import "gopurs/output/gopurs_runtime"



func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func Test_FibFFI_RunFibFFI(limit int) int {
	return (fib(int(limit)))
}


// --- Auto-generated FFI wrappers ---
var _Gopurs_Test_FibFFI_RunFibFFI = // TAST: (Func [Int] Int)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_res := Test_FibFFI_RunFibFFI(go_arg0)
	return gopurs_runtime.Int(int64(go_res))
})