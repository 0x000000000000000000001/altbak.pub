package Data_Array_Partial

import (
	pkg_Data_Array "gopurs/output/Data.Array"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_tail gopurs_runtime.Value
var once_tail sync.Once
func Get_tail() gopurs_runtime.Value {
	once_tail.Do(func() {
		cache_tail = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_tail(_dollar__unused_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_tail
}

var cache_last gopurs_runtime.Value
var once_last sync.Once
func Get_last() gopurs_runtime.Value {
	once_last.Do(func() {
		cache_last = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_last(_dollar__unused_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
	})
	return cache_last
}

var cache_init gopurs_runtime.Value
var once_init sync.Once
func Get_init() gopurs_runtime.Value {
	once_init.Do(func() {
		cache_init = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_init(_dollar__unused_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_init
}

var cache_head gopurs_runtime.Value
var once_head sync.Once
func Get_head() gopurs_runtime.Value {
	once_head.Do(func() {
		cache_head = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_head(_dollar__unused_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
	})
	return cache_head
}

var cache_slice__3011328576 gopurs_runtime.Value
var once_slice__3011328576 sync.Once
func Get_slice__3011328576() gopurs_runtime.Value {
	once_slice__3011328576.Do(func() {
		cache_slice__3011328576 = gopurs_runtime.Func3(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_slice__3011328576(__local_var_0_box.IntVal, __local_var_1_box.IntVal, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_slice__3011328576
}

var cache_unsafeIndex__2808089623 gopurs_runtime.Value
var once_unsafeIndex__2808089623 sync.Once
func Get_unsafeIndex__2808089623() gopurs_runtime.Value {
	once_unsafeIndex__2808089623.Do(func() {
		cache_unsafeIndex__2808089623 = gopurs_runtime.Func3(func(_dollar__unused_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeIndex__2808089623(_dollar__unused_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), __local_var_2_box.IntVal)
})
	})
	return cache_unsafeIndex__2808089623
}

var cache_sub__1043827704 gopurs_runtime.Value
var once_sub__1043827704 sync.Once
func Get_sub__1043827704() gopurs_runtime.Value {
	once_sub__1043827704.Do(func() {
		cache_sub__1043827704 = pkg_Data_Ring.Get_intSub()
	})
	return cache_sub__1043827704
}

var cache_sub__3675938712 gopurs_runtime.Value
var once_sub__3675938712 sync.Once
func Get_sub__3675938712() gopurs_runtime.Value {
	once_sub__3675938712.Do(func() {
		cache_sub__3675938712 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__3675938712(gopurs_runtime.CoerceToStruct[pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sub__3675938712
}

func Call_tail(_dollar__unused_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(pkg_Data_Array.Get_sliceImpl(), gopurs_runtime.Int(1), gopurs_runtime.Int(int64(len(xs_1))), gopurs_runtime.Array(xs_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_last(_dollar__unused_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return xs_1[gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Int(int64(len(xs_1))), gopurs_runtime.Int(1)).IntVal]
}

func Call_init(_dollar__unused_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(pkg_Data_Array.Get_sliceImpl(), gopurs_runtime.Int(0), gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Int(int64(len(xs_1))), gopurs_runtime.Int(1)), gopurs_runtime.Array(xs_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_head(_dollar__unused_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.ArrayAccess(gopurs_runtime.Array(xs_1), 0)
}

func Call_slice__3011328576(__local_var_0_loop int64, __local_var_1_loop int64, __local_var_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var __local_var_0 int64 = __local_var_0_loop
_ = __local_var_0
var __local_var_1 int64 = __local_var_1_loop
_ = __local_var_1
var __local_var_2 []gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(pkg_Data_Array.Get_sliceImpl(), gopurs_runtime.Int(__local_var_0), gopurs_runtime.Int(__local_var_1), gopurs_runtime.Array(__local_var_2)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_unsafeIndex__2808089623(_dollar__unused_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value, __local_var_2_loop int64) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 int64 = __local_var_2_loop
_ = __local_var_2
return __local_var_1[__local_var_2]
}

func Call_sub__3675938712(dict_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}


