package Data_Ord

import "gopurs/output/gopurs_runtime"

func OrdBooleanImpl(lt interface{}, eq interface{}, gt interface{}, x bool, y bool) interface{} {
	if !x && y {
		return lt
	} else if x == y {
		return eq
	}
	return gt
}
func OrdIntImpl(lt interface{}, eq interface{}, gt interface{}, x int64, y int64) interface{} {
	if x < y { return lt }
	if x == y { return eq }
	return gt
}
func OrdCharImpl(lt interface{}, eq interface{}, gt interface{}, x string, y string) interface{} {
	if x < y { return lt }
	if x == y { return eq }
	return gt
}
func OrdStringImpl(lt interface{}, eq interface{}, gt interface{}, x string, y string) interface{} {
	if x < y { return lt }
	if x == y { return eq }
	return gt
}
func OrdNumberImpl(lt interface{}, eq interface{}, gt interface{}, x float64, y float64) interface{} {
	if x < y { return lt }
	if x == y { return eq }
	return gt
}
func OrdArrayImpl(f func(interface{}, interface{}) int64, xs []interface{}, ys []interface{}) int64 {
	xlen := len(xs)
	ylen := len(ys)
	for i := 0; i < xlen && i < ylen; i++ {
		o := f(xs[i], ys[i])
		if o != 0 {
			return o
		}
	}
	if xlen == ylen {
		return 0
	} else if xlen > ylen {
		return 1
	} else {
		return -1
	}
}


// --- Auto-generated FFI wrappers ---
var _Gopurs_OrdArrayImpl = // TAST: (Func [(Func [(TypeVar a), (TypeVar a)] Int), (Array (TypeVar a)), (Array (TypeVar a))] Int)
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any, p0_1 any) int64 {
			inner_res0 := gopurs_runtime.Apply2(arg0, gopurs_runtime.Box(p0_0), gopurs_runtime.Box(p0_1))
			return gopurs_runtime.Unbox[int64](inner_res0)
		}
	arg1_arr := *(*[]gopurs_runtime.Value)(arg1.UnsafePtr)
	go_arg1 := make([]any, len(arg1_arr))
	for i, v := range arg1_arr { go_arg1[i] = v }
	arg2_arr := *(*[]gopurs_runtime.Value)(arg2.UnsafePtr)
	go_arg2 := make([]any, len(arg2_arr))
	for i, v := range arg2_arr { go_arg2[i] = v }
	go_res := OrdArrayImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_OrdBooleanImpl = // TAST: (Func [(ADT ["Data","Ordering","Ordering"] []), (ADT ["Data","Ordering","Ordering"] []), (ADT ["Data","Ordering","Ordering"] []), Boolean, Boolean] (ADT ["Data","Ordering","Ordering"] []))
gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := gopurs_runtime.Unbox[bool](arg3)
	go_arg4 := gopurs_runtime.Unbox[bool](arg4)
	go_res := OrdBooleanImpl(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_OrdCharImpl = // TAST: (Func [(ADT ["Data","Ordering","Ordering"] []), (ADT ["Data","Ordering","Ordering"] []), (ADT ["Data","Ordering","Ordering"] []), Char, Char] (ADT ["Data","Ordering","Ordering"] []))
gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := gopurs_runtime.Unbox[string](arg3)
	go_arg4 := gopurs_runtime.Unbox[string](arg4)
	go_res := OrdCharImpl(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_OrdIntImpl = // TAST: (Func [(ADT ["Data","Ordering","Ordering"] []), (ADT ["Data","Ordering","Ordering"] []), (ADT ["Data","Ordering","Ordering"] []), Int, Int] (ADT ["Data","Ordering","Ordering"] []))
gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := gopurs_runtime.Unbox[int64](arg3)
	go_arg4 := gopurs_runtime.Unbox[int64](arg4)
	go_res := OrdIntImpl(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_OrdNumberImpl = // TAST: (Func [(ADT ["Data","Ordering","Ordering"] []), (ADT ["Data","Ordering","Ordering"] []), (ADT ["Data","Ordering","Ordering"] []), Number, Number] (ADT ["Data","Ordering","Ordering"] []))
gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := gopurs_runtime.Unbox[float64](arg3)
	go_arg4 := gopurs_runtime.Unbox[float64](arg4)
	go_res := OrdNumberImpl(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_OrdStringImpl = // TAST: (Func [(ADT ["Data","Ordering","Ordering"] []), (ADT ["Data","Ordering","Ordering"] []), (ADT ["Data","Ordering","Ordering"] []), String, String] (ADT ["Data","Ordering","Ordering"] []))
gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := gopurs_runtime.Unbox[string](arg3)
	go_arg4 := gopurs_runtime.Unbox[string](arg4)
	go_res := OrdStringImpl(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return gopurs_runtime.Box(go_res)
})