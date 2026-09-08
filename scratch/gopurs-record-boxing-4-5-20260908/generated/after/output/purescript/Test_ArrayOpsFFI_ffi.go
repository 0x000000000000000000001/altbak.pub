package purescript

import "gopurs/output/gopurs_runtime"



func Test_ArrayOpsFFI_RunArrayOpsFFI(limit int) int {
	n := int(limit)
	
	arr := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		arr = append(arr, i)
	}
	
	evens := make([]int, 0)
	for _, x := range arr {
		if x%2 == 0 {
			evens = append(evens, x)
		}
	}
	
	sum := 0
	for _, x := range evens {
		sum += x
	}
	
	return (sum)
}


// --- Auto-generated FFI wrappers ---
var _Gopurs_Test_ArrayOpsFFI_RunArrayOpsFFI = // TAST: (Func [Int] Int)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_res := Test_ArrayOpsFFI_RunArrayOpsFFI(go_arg0)
	return gopurs_runtime.Int(int64(go_res))
})