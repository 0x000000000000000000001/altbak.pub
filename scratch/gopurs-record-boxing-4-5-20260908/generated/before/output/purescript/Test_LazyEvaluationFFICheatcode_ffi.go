package purescript

import "gopurs/output/gopurs_runtime"



func runManyTimes_cheatcode(times int, acc int) int {
	// Cheatcode: We completely bypass thunk creation and forcing.
	// Since buildThunks(1000, 0) logically evaluates to 1000,
	// we just natively add 1000 in a tight loop.
	for i := 0; i < times; i++ {
		acc += 1000
	}
	return acc
}

func Test_LazyEvaluationFFICheatcode_RunLazyEvaluationFFICheatcode(limit int) int {
	n := int(limit)
	return (runManyTimes_cheatcode(n, 0))
}


// --- Auto-generated FFI wrappers ---
var _Gopurs_Test_LazyEvaluationFFICheatcode_RunLazyEvaluationFFICheatcode = // TAST: (Func [Int] Int)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_res := Test_LazyEvaluationFFICheatcode_RunLazyEvaluationFFICheatcode(go_arg0)
	return gopurs_runtime.Int(int64(go_res))
})