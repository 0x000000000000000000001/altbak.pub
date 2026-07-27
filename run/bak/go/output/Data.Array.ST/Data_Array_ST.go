package Data_Array_ST

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Monad_ST_Uncurried "gopurs/output/Control.Monad.ST.Uncurried"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Control_Monad_ST_Internal "gopurs/output/Control.Monad.ST.Internal"
	unsafe "unsafe"
)

var cache_unshiftAll gopurs_runtime.Value
var once_unshiftAll sync.Once
func Get_unshiftAll() gopurs_runtime.Value {
	once_unshiftAll.Do(func() {
		cache_unshiftAll = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return gopurs_runtime.Int(func(inner_arg0 []interface{}, inner_arg1 gopurs_runtime.Value) func() int64 {
return func() int64 {
return gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn2(), Get_unshiftAllImpl()), func() gopurs_runtime.Value {
					arr := inner_arg0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}(), inner_arg1), nil).IntVal
}
}(func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(arg0.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}(), arg1)())
})
})
	})
	return cache_unshiftAll
}

var cache_unshift gopurs_runtime.Value
var once_unshift sync.Once
func Get_unshift() gopurs_runtime.Value {
	once_unshift.Do(func() {
		cache_unshift = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unshift(gopurs_runtime.UnboxAny(a_0_box))
})
	})
	return cache_unshift
}

var cache_unsafeThaw gopurs_runtime.Value
var once_unsafeThaw sync.Once
func Get_unsafeThaw() gopurs_runtime.Value {
	once_unsafeThaw.Do(func() {
		cache_unsafeThaw = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return func(inner_arg0 []interface{}) func() gopurs_runtime.Value {
return func() gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn1(), Get_unsafeThawImpl()), func() gopurs_runtime.Value {
					arr := inner_arg0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()), nil)
}
}(func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(arg0.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}())()
})
})
	})
	return cache_unsafeThaw
}

var cache_unsafeFreeze gopurs_runtime.Value
var once_unsafeFreeze sync.Once
func Get_unsafeFreeze() gopurs_runtime.Value {
	once_unsafeFreeze.Do(func() {
		cache_unsafeFreeze = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := func(inner_arg0 gopurs_runtime.Value) func() []interface{} {
return func() []interface{} {
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn1(), Get_unsafeFreezeImpl()), inner_arg0), nil).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}
}(arg0)()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
})
	})
	return cache_unsafeFreeze
}

var cache_toAssocArray gopurs_runtime.Value
var once_toAssocArray sync.Once
func Get_toAssocArray() gopurs_runtime.Value {
	once_toAssocArray.Do(func() {
		cache_toAssocArray = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := func(inner_arg0 gopurs_runtime.Value) func() []interface{} {
return func() []interface{} {
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn1(), Get_toAssocArrayImpl()), inner_arg0), nil).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}
}(arg0)()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
})
	})
	return cache_toAssocArray
}

var cache_thaw gopurs_runtime.Value
var once_thaw sync.Once
func Get_thaw() gopurs_runtime.Value {
	once_thaw.Do(func() {
		cache_thaw = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return func(inner_arg0 []interface{}) func() gopurs_runtime.Value {
return func() gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn1(), Get_thawImpl()), func() gopurs_runtime.Value {
					arr := inner_arg0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()), nil)
}
}(func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(arg0.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}())()
})
})
	})
	return cache_thaw
}

var cache_thaw__func_arrinterface____func___gopurs_runtime_Value_4258121115 gopurs_runtime.Value
var once_thaw__func_arrinterface____func___gopurs_runtime_Value_4258121115 sync.Once
func Get_thaw__func_arrinterface____func___gopurs_runtime_Value_4258121115() gopurs_runtime.Value {
	once_thaw__func_arrinterface____func___gopurs_runtime_Value_4258121115.Do(func() {
		cache_thaw__func_arrinterface____func___gopurs_runtime_Value_4258121115 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return func(inner_arg0 []interface{}) func() gopurs_runtime.Value {
return func() gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn1(), Get_thawImpl()), func() gopurs_runtime.Value {
					arr := inner_arg0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()), nil)
}
}(func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(arg0.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}())()
})
})
	})
	return cache_thaw__func_arrinterface____func___gopurs_runtime_Value_4258121115
}

var cache_withArray gopurs_runtime.Value
var once_withArray sync.Once
func Get_withArray() gopurs_runtime.Value {
	once_withArray.Do(func() {
		cache_withArray = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_withArray(func(inner_arg0 gopurs_runtime.Value) func() interface{} {
return func() interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Apply(f_0_box, inner_arg0), nil))
}
}, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}())()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
})
	})
	return cache_withArray
}

var cache_withArray__func_func_gopurs_runtime_Value__func___interface____arrinterface____func___arrinterface___2856329898 gopurs_runtime.Value
var once_withArray__func_func_gopurs_runtime_Value__func___interface____arrinterface____func___arrinterface___2856329898 sync.Once
func Get_withArray__func_func_gopurs_runtime_Value__func___interface____arrinterface____func___arrinterface___2856329898() gopurs_runtime.Value {
	once_withArray__func_func_gopurs_runtime_Value__func___interface____arrinterface____func___arrinterface___2856329898.Do(func() {
		cache_withArray__func_func_gopurs_runtime_Value__func___interface____arrinterface____func___arrinterface___2856329898 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_withArray__func_func_gopurs_runtime_Value__func___interface____arrinterface____func___arrinterface___2856329898(func(inner_arg0 gopurs_runtime.Value) func() interface{} {
return func() interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Apply(f_0_box, inner_arg0), nil))
}
}, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}())()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
})
	})
	return cache_withArray__func_func_gopurs_runtime_Value__func___interface____arrinterface____func___arrinterface___2856329898
}

var cache_splice gopurs_runtime.Value
var once_splice sync.Once
func Get_splice() gopurs_runtime.Value {
	once_splice.Do(func() {
		cache_splice = gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := func(inner_arg0 int64, inner_arg1 int64, inner_arg2 []interface{}, inner_arg3 gopurs_runtime.Value) func() []interface{} {
return func() []interface{} {
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(gopurs_runtime.Apply4(gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn4(), Get_spliceImpl()), gopurs_runtime.Int(inner_arg0), gopurs_runtime.Int(inner_arg1), func() gopurs_runtime.Value {
					arr := inner_arg2
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}(), inner_arg3), nil).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}
}(arg0.IntVal, arg1.IntVal, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(arg2.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}(), arg3)()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
})
	})
	return cache_splice
}

var cache_sortBy gopurs_runtime.Value
var once_sortBy sync.Once
func Get_sortBy() gopurs_runtime.Value {
	once_sortBy.Do(func() {
		cache_sortBy = gopurs_runtime.Func(func(comp_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sortBy(func(inner_arg0 interface{}, inner_arg1 interface{}) gopurs_runtime.Value {
return gopurs_runtime.Apply2(comp_0_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1))
})
})
	})
	return cache_sortBy
}

var cache_sortBy__func_func_interface____interface____gopurs_runtime_Value__gopurs_runtime_Value__func___gopurs_runtime_Value_2740821710 gopurs_runtime.Value
var once_sortBy__func_func_interface____interface____gopurs_runtime_Value__gopurs_runtime_Value__func___gopurs_runtime_Value_2740821710 sync.Once
func Get_sortBy__func_func_interface____interface____gopurs_runtime_Value__gopurs_runtime_Value__func___gopurs_runtime_Value_2740821710() gopurs_runtime.Value {
	once_sortBy__func_func_interface____interface____gopurs_runtime_Value__gopurs_runtime_Value__func___gopurs_runtime_Value_2740821710.Do(func() {
		cache_sortBy__func_func_interface____interface____gopurs_runtime_Value__gopurs_runtime_Value__func___gopurs_runtime_Value_2740821710 = gopurs_runtime.Func(func(comp_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sortBy__func_func_interface____interface____gopurs_runtime_Value__gopurs_runtime_Value__func___gopurs_runtime_Value_2740821710(func(inner_arg0 interface{}, inner_arg1 interface{}) gopurs_runtime.Value {
return gopurs_runtime.Apply2(comp_0_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1))
})
})
	})
	return cache_sortBy__func_func_interface____interface____gopurs_runtime_Value__gopurs_runtime_Value__func___gopurs_runtime_Value_2740821710
}

var cache_sortWith gopurs_runtime.Value
var once_sortWith sync.Once
func Get_sortWith() gopurs_runtime.Value {
	once_sortWith.Do(func() {
		cache_sortWith = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sortWith(dictOrd_0_box, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_1_box, gopurs_runtime.Any(inner_arg0)))
})
})
	})
	return cache_sortWith
}

var cache_sort gopurs_runtime.Value
var once_sort sync.Once
func Get_sort() gopurs_runtime.Value {
	once_sort.Do(func() {
		cache_sort = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sort(dictOrd_0_box)
})
	})
	return cache_sort
}

var cache_shift gopurs_runtime.Value
var once_shift sync.Once
func Get_shift() gopurs_runtime.Value {
	once_shift.Do(func() {
		cache_shift = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(func(inner_arg0 gopurs_runtime.Value) func() *pkg_Data_Maybe.Constructor_Just[interface{}] {
return func() *pkg_Data_Maybe.Constructor_Just[interface{}] {
return (*pkg_Data_Maybe.Constructor_Just[interface{}])(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply3(pkg_Control_Monad_ST_Uncurried.Get_runSTFn3(), Get_shiftImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})), inner_arg0), nil).UnsafePtr)
}
}(arg0)())}
})
})
	})
	return cache_shift
}

var cache_run gopurs_runtime.Value
var once_run sync.Once
func Get_run() gopurs_runtime.Value {
	once_run.Do(func() {
		cache_run = gopurs_runtime.Func(func(st_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_run(func() gopurs_runtime.Value {
return gopurs_runtime.Apply(st_0_box, nil)
})
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_run
}

var cache_run__func_func___gopurs_runtime_Value__arrinterface___1582179099 gopurs_runtime.Value
var once_run__func_func___gopurs_runtime_Value__arrinterface___1582179099 sync.Once
func Get_run__func_func___gopurs_runtime_Value__arrinterface___1582179099() gopurs_runtime.Value {
	once_run__func_func___gopurs_runtime_Value__arrinterface___1582179099.Do(func() {
		cache_run__func_func___gopurs_runtime_Value__arrinterface___1582179099 = gopurs_runtime.Func(func(st_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_run__func_func___gopurs_runtime_Value__arrinterface___1582179099(func() gopurs_runtime.Value {
return gopurs_runtime.Apply(st_0_box, nil)
})
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_run__func_func___gopurs_runtime_Value__arrinterface___1582179099
}

var cache_pushAll gopurs_runtime.Value
var once_pushAll sync.Once
func Get_pushAll() gopurs_runtime.Value {
	once_pushAll.Do(func() {
		cache_pushAll = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return gopurs_runtime.Int(func(inner_arg0 []interface{}, inner_arg1 gopurs_runtime.Value) func() int64 {
return func() int64 {
return gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn2(), Get_pushAllImpl()), func() gopurs_runtime.Value {
					arr := inner_arg0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}(), inner_arg1), nil).IntVal
}
}(func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(arg0.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}(), arg1)())
})
})
	})
	return cache_pushAll
}

var cache_push gopurs_runtime.Value
var once_push sync.Once
func Get_push() gopurs_runtime.Value {
	once_push.Do(func() {
		cache_push = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return gopurs_runtime.Int(func(inner_arg0 interface{}, inner_arg1 gopurs_runtime.Value) func() int64 {
return func() int64 {
return gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn2(), Get_pushImpl()), gopurs_runtime.Any(inner_arg0), inner_arg1), nil).IntVal
}
}(gopurs_runtime.UnboxAny(arg0), arg1)())
})
})
	})
	return cache_push
}

var cache_push__func_arrinterface____gopurs_runtime_Value__func___int64_2602532183 gopurs_runtime.Value
var once_push__func_arrinterface____gopurs_runtime_Value__func___int64_2602532183 sync.Once
func Get_push__func_arrinterface____gopurs_runtime_Value__func___int64_2602532183() gopurs_runtime.Value {
	once_push__func_arrinterface____gopurs_runtime_Value__func___int64_2602532183.Do(func() {
		cache_push__func_arrinterface____gopurs_runtime_Value__func___int64_2602532183 = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return gopurs_runtime.Int(func(inner_arg0 []interface{}, inner_arg1 gopurs_runtime.Value) func() int64 {
return func() int64 {
return gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn2(), Get_pushImpl()), func() gopurs_runtime.Value {
					arr := inner_arg0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}(), inner_arg1), nil).IntVal
}
}(func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(arg0.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}(), arg1)())
})
})
	})
	return cache_push__func_arrinterface____gopurs_runtime_Value__func___int64_2602532183
}

var cache_push__func_interface____gopurs_runtime_Value__func___int64_236408145 gopurs_runtime.Value
var once_push__func_interface____gopurs_runtime_Value__func___int64_236408145 sync.Once
func Get_push__func_interface____gopurs_runtime_Value__func___int64_236408145() gopurs_runtime.Value {
	once_push__func_interface____gopurs_runtime_Value__func___int64_236408145.Do(func() {
		cache_push__func_interface____gopurs_runtime_Value__func___int64_236408145 = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return gopurs_runtime.Int(func(inner_arg0 interface{}, inner_arg1 gopurs_runtime.Value) func() int64 {
return func() int64 {
return gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn2(), Get_pushImpl()), gopurs_runtime.Any(inner_arg0), inner_arg1), nil).IntVal
}
}(gopurs_runtime.UnboxAny(arg0), arg1)())
})
})
	})
	return cache_push__func_interface____gopurs_runtime_Value__func___int64_236408145
}

var cache_push__func_interface____gopurs_runtime_Value__func___int64_2223234226 gopurs_runtime.Value
var once_push__func_interface____gopurs_runtime_Value__func___int64_2223234226 sync.Once
func Get_push__func_interface____gopurs_runtime_Value__func___int64_2223234226() gopurs_runtime.Value {
	once_push__func_interface____gopurs_runtime_Value__func___int64_2223234226.Do(func() {
		cache_push__func_interface____gopurs_runtime_Value__func___int64_2223234226 = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return gopurs_runtime.Int(func(inner_arg0 interface{}, inner_arg1 gopurs_runtime.Value) func() int64 {
return func() int64 {
return gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn2(), Get_pushImpl()), gopurs_runtime.Any(inner_arg0), inner_arg1), nil).IntVal
}
}(gopurs_runtime.UnboxAny(arg0), arg1)())
})
})
	})
	return cache_push__func_interface____gopurs_runtime_Value__func___int64_2223234226
}

var cache_pop gopurs_runtime.Value
var once_pop sync.Once
func Get_pop() gopurs_runtime.Value {
	once_pop.Do(func() {
		cache_pop = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(func(inner_arg0 gopurs_runtime.Value) func() *pkg_Data_Maybe.Constructor_Just[interface{}] {
return func() *pkg_Data_Maybe.Constructor_Just[interface{}] {
return (*pkg_Data_Maybe.Constructor_Just[interface{}])(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply3(pkg_Control_Monad_ST_Uncurried.Get_runSTFn3(), Get_popImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})), inner_arg0), nil).UnsafePtr)
}
}(arg0)())}
})
})
	})
	return cache_pop
}

var cache_poke gopurs_runtime.Value
var once_poke sync.Once
func Get_poke() gopurs_runtime.Value {
	once_poke.Do(func() {
		cache_poke = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return gopurs_runtime.Bool(func(inner_arg0 int64, inner_arg1 interface{}, inner_arg2 gopurs_runtime.Value) func() bool {
return func() bool {
return (gopurs_runtime.Apply(gopurs_runtime.Apply3(gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn3(), Get_pokeImpl()), gopurs_runtime.Int(inner_arg0), gopurs_runtime.Any(inner_arg1), inner_arg2), nil).IntVal) != (0)
}
}(arg0.IntVal, gopurs_runtime.UnboxAny(arg1), arg2)())
})
})
	})
	return cache_poke
}

var cache_poke__func_int64__interface____gopurs_runtime_Value__func___bool_2144247274 gopurs_runtime.Value
var once_poke__func_int64__interface____gopurs_runtime_Value__func___bool_2144247274 sync.Once
func Get_poke__func_int64__interface____gopurs_runtime_Value__func___bool_2144247274() gopurs_runtime.Value {
	once_poke__func_int64__interface____gopurs_runtime_Value__func___bool_2144247274.Do(func() {
		cache_poke__func_int64__interface____gopurs_runtime_Value__func___bool_2144247274 = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return gopurs_runtime.Bool(func(inner_arg0 int64, inner_arg1 interface{}, inner_arg2 gopurs_runtime.Value) func() bool {
return func() bool {
return (gopurs_runtime.Apply(gopurs_runtime.Apply3(gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn3(), Get_pokeImpl()), gopurs_runtime.Int(inner_arg0), gopurs_runtime.Any(inner_arg1), inner_arg2), nil).IntVal) != (0)
}
}(arg0.IntVal, gopurs_runtime.UnboxAny(arg1), arg2)())
})
})
	})
	return cache_poke__func_int64__interface____gopurs_runtime_Value__func___bool_2144247274
}

var cache_peek gopurs_runtime.Value
var once_peek sync.Once
func Get_peek() gopurs_runtime.Value {
	once_peek.Do(func() {
		cache_peek = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(func(inner_arg0 int64, inner_arg1 gopurs_runtime.Value) func() *pkg_Data_Maybe.Constructor_Just[interface{}] {
return func() *pkg_Data_Maybe.Constructor_Just[interface{}] {
return (*pkg_Data_Maybe.Constructor_Just[interface{}])(gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.Apply3(pkg_Control_Monad_ST_Uncurried.Get_runSTFn4(), Get_peekImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})), gopurs_runtime.Int(inner_arg0), inner_arg1), nil).UnsafePtr)
}
}(arg0.IntVal, arg1)())}
})
})
	})
	return cache_peek
}

var cache_modify gopurs_runtime.Value
var once_modify sync.Once
func Get_modify() gopurs_runtime.Value {
	once_modify.Do(func() {
		cache_modify = gopurs_runtime.Func3(func(i_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_modify(i_0_box.IntVal, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_1_box, gopurs_runtime.Any(inner_arg0)))
}, xs_2_box)())
})
})
	})
	return cache_modify
}

var cache_modify__func_int64__func_interface____interface____gopurs_runtime_Value__func___bool_3704614031 gopurs_runtime.Value
var once_modify__func_int64__func_interface____interface____gopurs_runtime_Value__func___bool_3704614031 sync.Once
func Get_modify__func_int64__func_interface____interface____gopurs_runtime_Value__func___bool_3704614031() gopurs_runtime.Value {
	once_modify__func_int64__func_interface____interface____gopurs_runtime_Value__func___bool_3704614031.Do(func() {
		cache_modify__func_int64__func_interface____interface____gopurs_runtime_Value__func___bool_3704614031 = gopurs_runtime.Func3(func(i_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_modify__func_int64__func_interface____interface____gopurs_runtime_Value__func___bool_3704614031(i_0_box.IntVal, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_1_box, gopurs_runtime.Any(inner_arg0)))
}, xs_2_box)())
})
})
	})
	return cache_modify__func_int64__func_interface____interface____gopurs_runtime_Value__func___bool_3704614031
}

var cache_length gopurs_runtime.Value
var once_length sync.Once
func Get_length() gopurs_runtime.Value {
	once_length.Do(func() {
		cache_length = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return gopurs_runtime.Int(func(inner_arg0 gopurs_runtime.Value) func() int64 {
return func() int64 {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn1(), Get_lengthImpl()), inner_arg0), nil).IntVal
}
}(arg0)())
})
})
	})
	return cache_length
}

var cache_freeze gopurs_runtime.Value
var once_freeze sync.Once
func Get_freeze() gopurs_runtime.Value {
	once_freeze.Do(func() {
		cache_freeze = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := func(inner_arg0 gopurs_runtime.Value) func() []interface{} {
return func() []interface{} {
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn1(), Get_freezeImpl()), inner_arg0), nil).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}
}(arg0)()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
})
	})
	return cache_freeze
}

var cache_clone gopurs_runtime.Value
var once_clone sync.Once
func Get_clone() gopurs_runtime.Value {
	once_clone.Do(func() {
		cache_clone = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return func(inner_arg0 gopurs_runtime.Value) func() gopurs_runtime.Value {
return func() gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn1(), Get_cloneImpl()), inner_arg0), nil)
}
}(arg0)()
})
})
	})
	return cache_clone
}

var cache_cloneImpl gopurs_runtime.Value
var once_cloneImpl sync.Once
func Get_cloneImpl() gopurs_runtime.Value {
	once_cloneImpl.Do(func() {
		cache_cloneImpl = CloneImpl
	})
	return cache_cloneImpl
}

var cache_freezeImpl gopurs_runtime.Value
var once_freezeImpl sync.Once
func Get_freezeImpl() gopurs_runtime.Value {
	once_freezeImpl.Do(func() {
		cache_freezeImpl = FreezeImpl
	})
	return cache_freezeImpl
}

var cache_lengthImpl gopurs_runtime.Value
var once_lengthImpl sync.Once
func Get_lengthImpl() gopurs_runtime.Value {
	once_lengthImpl.Do(func() {
		cache_lengthImpl = LengthImpl
	})
	return cache_lengthImpl
}

var cache_new_ gopurs_runtime.Value
var once_new_ sync.Once
func Get_new_() gopurs_runtime.Value {
	once_new_.Do(func() {
		cache_new_ = gopurs_runtime.Func0(func() gopurs_runtime.Value {
return New_()
})
	})
	return cache_new_
}

var cache_peekImpl gopurs_runtime.Value
var once_peekImpl sync.Once
func Get_peekImpl() gopurs_runtime.Value {
	once_peekImpl.Do(func() {
		cache_peekImpl = PeekImpl
	})
	return cache_peekImpl
}

var cache_pokeImpl gopurs_runtime.Value
var once_pokeImpl sync.Once
func Get_pokeImpl() gopurs_runtime.Value {
	once_pokeImpl.Do(func() {
		cache_pokeImpl = PokeImpl
	})
	return cache_pokeImpl
}

var cache_popImpl gopurs_runtime.Value
var once_popImpl sync.Once
func Get_popImpl() gopurs_runtime.Value {
	once_popImpl.Do(func() {
		cache_popImpl = PopImpl
	})
	return cache_popImpl
}

var cache_pushAllImpl gopurs_runtime.Value
var once_pushAllImpl sync.Once
func Get_pushAllImpl() gopurs_runtime.Value {
	once_pushAllImpl.Do(func() {
		cache_pushAllImpl = PushAllImpl
	})
	return cache_pushAllImpl
}

var cache_pushImpl gopurs_runtime.Value
var once_pushImpl sync.Once
func Get_pushImpl() gopurs_runtime.Value {
	once_pushImpl.Do(func() {
		cache_pushImpl = PushImpl
	})
	return cache_pushImpl
}

var cache_shiftImpl gopurs_runtime.Value
var once_shiftImpl sync.Once
func Get_shiftImpl() gopurs_runtime.Value {
	once_shiftImpl.Do(func() {
		cache_shiftImpl = ShiftImpl
	})
	return cache_shiftImpl
}

var cache_sortByImpl gopurs_runtime.Value
var once_sortByImpl sync.Once
func Get_sortByImpl() gopurs_runtime.Value {
	once_sortByImpl.Do(func() {
		cache_sortByImpl = SortByImpl
	})
	return cache_sortByImpl
}

var cache_spliceImpl gopurs_runtime.Value
var once_spliceImpl sync.Once
func Get_spliceImpl() gopurs_runtime.Value {
	once_spliceImpl.Do(func() {
		cache_spliceImpl = SpliceImpl
	})
	return cache_spliceImpl
}

var cache_thawImpl gopurs_runtime.Value
var once_thawImpl sync.Once
func Get_thawImpl() gopurs_runtime.Value {
	once_thawImpl.Do(func() {
		cache_thawImpl = ThawImpl
	})
	return cache_thawImpl
}

var cache_toAssocArrayImpl gopurs_runtime.Value
var once_toAssocArrayImpl sync.Once
func Get_toAssocArrayImpl() gopurs_runtime.Value {
	once_toAssocArrayImpl.Do(func() {
		cache_toAssocArrayImpl = ToAssocArrayImpl
	})
	return cache_toAssocArrayImpl
}

var cache_unsafeFreezeImpl gopurs_runtime.Value
var once_unsafeFreezeImpl sync.Once
func Get_unsafeFreezeImpl() gopurs_runtime.Value {
	once_unsafeFreezeImpl.Do(func() {
		cache_unsafeFreezeImpl = UnsafeFreezeImpl
	})
	return cache_unsafeFreezeImpl
}

var cache_unsafeThawImpl gopurs_runtime.Value
var once_unsafeThawImpl sync.Once
func Get_unsafeThawImpl() gopurs_runtime.Value {
	once_unsafeThawImpl.Do(func() {
		cache_unsafeThawImpl = UnsafeThawImpl
	})
	return cache_unsafeThawImpl
}

var cache_unshiftAllImpl gopurs_runtime.Value
var once_unshiftAllImpl sync.Once
func Get_unshiftAllImpl() gopurs_runtime.Value {
	once_unshiftAllImpl.Do(func() {
		cache_unshiftAllImpl = UnshiftAllImpl
	})
	return cache_unshiftAllImpl
}

func Call_unshift(a_0_loop interface{}) gopurs_runtime.Value {
var a_0 interface{} = a_0_loop
_ = a_0
return gopurs_runtime.Apply2(pkg_Control_Monad_ST_Uncurried.Get_runSTFn2(), Get_unshiftAllImpl(), func() gopurs_runtime.Value {
					arr := []interface{}{a_0}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}())
}

func Call_withArray(f_0_loop func(gopurs_runtime.Value) func() interface{}, xs_1_loop []interface{}) func() []interface{} {
var f_0 func(gopurs_runtime.Value) func() interface{} = f_0_loop
_ = f_0
var xs_1 []interface{} = xs_1_loop
_ = xs_1
return func() []interface{} {
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(Get_thawImpl(), func() gopurs_runtime.Value {
					arr := xs_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}())
}), gopurs_runtime.Func(func(result_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Func0(func() gopurs_runtime.Value {
return gopurs_runtime.Any(f_0(result_2)())
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(Get_unsafeFreezeImpl(), result_2)
})
}))
})), nil).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}
}

func Call_withArray__func_func_gopurs_runtime_Value__func___interface____arrinterface____func___arrinterface___2856329898(f_0_loop func(gopurs_runtime.Value) func() interface{}, xs_1_loop []interface{}) func() []interface{} {
var f_0 func(gopurs_runtime.Value) func() interface{} = f_0_loop
_ = f_0
var xs_1 []interface{} = xs_1_loop
_ = xs_1
return func() []interface{} {
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(Get_thawImpl(), func() gopurs_runtime.Value {
					arr := xs_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}())
}), gopurs_runtime.Func(func(result_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Func0(func() gopurs_runtime.Value {
return gopurs_runtime.Any(f_0(result_2)())
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(Get_unsafeFreezeImpl(), result_2)
})
}))
})), nil).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}
}

func Call_sortBy(comp_0_loop func(interface{}, interface{}) gopurs_runtime.Value) gopurs_runtime.Value {
var comp_0 func(interface{}, interface{}) gopurs_runtime.Value = comp_0_loop
_ = comp_0
return gopurs_runtime.Apply3(pkg_Control_Monad_ST_Uncurried.Get_runSTFn3(), Get_sortByImpl(), gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return comp_0(gopurs_runtime.UnboxAny(arg0), gopurs_runtime.UnboxAny(arg1))
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_1.Type == 9 && v_1.IntVal == 380165415) {
__t0 = gopurs_runtime.Int(1)
goto end_branch_0
} else {

}
}
{
if (v_1.Type == 9 && v_1.IntVal == 902936544) {
__t0 = gopurs_runtime.Int(0)
goto end_branch_0
} else {

}
}
{
if (v_1.Type == 9 && v_1.IntVal == 1527465420) {
__t0 = gopurs_runtime.Int(-1)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}))
}

func Call_sortBy__func_func_interface____interface____gopurs_runtime_Value__gopurs_runtime_Value__func___gopurs_runtime_Value_2740821710(comp_0_loop func(interface{}, interface{}) gopurs_runtime.Value) gopurs_runtime.Value {
var comp_0 func(interface{}, interface{}) gopurs_runtime.Value = comp_0_loop
_ = comp_0
return gopurs_runtime.Apply3(pkg_Control_Monad_ST_Uncurried.Get_runSTFn3(), Get_sortByImpl(), gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return comp_0(gopurs_runtime.UnboxAny(arg0), gopurs_runtime.UnboxAny(arg1))
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_1.Type == 9 && v_1.IntVal == 380165415) {
__t0 = gopurs_runtime.Int(1)
goto end_branch_0
} else {

}
}
{
if (v_1.Type == 9 && v_1.IntVal == 902936544) {
__t0 = gopurs_runtime.Int(0)
goto end_branch_0
} else {

}
}
{
if (v_1.Type == 9 && v_1.IntVal == 1527465420) {
__t0 = gopurs_runtime.Int(-1)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}))
}

func Call_sortWith(dictOrd_0_loop gopurs_runtime.Value, f_1_loop func(interface{}) interface{}) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var f_1 func(interface{}) interface{} = f_1_loop
_ = f_1
return Call_sortBy(func(inner_arg0 interface{}, inner_arg1 interface{}) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), gopurs_runtime.Any(f_1(gopurs_runtime.UnboxAny(x_2))), gopurs_runtime.Any(f_1(gopurs_runtime.UnboxAny(y_3))))
}), gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1))
})
}

func Call_sort(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return Call_sortBy(func(inner_arg0 interface{}, inner_arg1 interface{}) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1))
})
}

func Call_run(st_0_loop func() gopurs_runtime.Value) []interface{} {
var st_0 func() gopurs_runtime.Value = st_0_loop
_ = st_0
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Func0(func() gopurs_runtime.Value {
return st_0()
}), Get_unsafeFreeze())).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}

func Call_run__func_func___gopurs_runtime_Value__arrinterface___1582179099(st_0_loop func() gopurs_runtime.Value) []interface{} {
var st_0 func() gopurs_runtime.Value = st_0_loop
_ = st_0
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Func0(func() gopurs_runtime.Value {
return st_0()
}), Get_unsafeFreeze())).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}

func Call_modify(i_0_loop int64, f_1_loop func(interface{}) interface{}, xs_2_loop gopurs_runtime.Value) func() bool {
var i_0 int64 = i_0_loop
_ = i_0
var f_1 func(interface{}) interface{} = f_1_loop
_ = f_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return func() bool {
return (gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp4(Get_peekImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}), gopurs_runtime.Int(i_0), xs_2)
}), gopurs_runtime.Func(func(entry_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (entry_3.Type == 9 && entry_3.IntVal == 930809136 && entry_3.UnsafePtr != nil) {
__local_var_4_1 := gopurs_runtime.Any(f_1(gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(entry_3.UnsafePtr).V0))))
_ = __local_var_4_1
__t0 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp3(Get_pokeImpl(), gopurs_runtime.Int(i_0), __local_var_4_1, xs_2)
})
goto end_branch_0
} else {

}
}
{
if (entry_3.Type == 9 && entry_3.IntVal == 930809136 && entry_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_applicativeST(), "pure"), gopurs_runtime.Bool(false))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})), nil).IntVal) != (0)
}
}

func Call_modify__func_int64__func_interface____interface____gopurs_runtime_Value__func___bool_3704614031(i_0_loop int64, f_1_loop func(interface{}) interface{}, xs_2_loop gopurs_runtime.Value) func() bool {
var i_0 int64 = i_0_loop
_ = i_0
var f_1 func(interface{}) interface{} = f_1_loop
_ = f_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return func() bool {
return (gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp4(Get_peekImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}), gopurs_runtime.Int(i_0), xs_2)
}), gopurs_runtime.Func(func(entry_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (entry_3.Type == 9 && entry_3.IntVal == 930809136 && entry_3.UnsafePtr != nil) {
__local_var_4_1 := gopurs_runtime.Any(f_1(gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(entry_3.UnsafePtr).V0))))
_ = __local_var_4_1
__t0 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp3(Get_pokeImpl(), gopurs_runtime.Int(i_0), __local_var_4_1, xs_2)
})
goto end_branch_0
} else {

}
}
{
if (entry_3.Type == 9 && entry_3.IntVal == 930809136 && entry_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_applicativeST(), "pure"), gopurs_runtime.Bool(false))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})), nil).IntVal) != (0)
}
}
