package purescript

import "gopurs/output/gopurs_runtime"



type Lazy func() int

func force(l Lazy) int {
	return l()
}

func deferFunc(f func() int) Lazy {
	return f
}

func buildThunks(depth int, acc Lazy) Lazy {
	if depth == 0 {
		return acc
	}
	return buildThunks(depth-1, deferFunc(func() int {
		return force(acc) + 1
	}))
}

func runManyTimes(times int, acc int) int {
	if times == 0 {
		return acc
	}
	return runManyTimes(times-1, acc+force(buildThunks(1000, deferFunc(func() int {
		return 0
	}))))
}

func Test_LazyEvaluationFFI_RunLazyEvaluationFFI(limit int) int {
	n := int(limit)
	return (runManyTimes(n, 0))
}


// --- Auto-generated FFI wrappers ---
var _Gopurs_Test_LazyEvaluationFFI_RunLazyEvaluationFFI = // TAST: (Func [Int] Int)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_res := Test_LazyEvaluationFFI_RunLazyEvaluationFFI(go_arg0)
	return gopurs_runtime.Int(int64(go_res))
})