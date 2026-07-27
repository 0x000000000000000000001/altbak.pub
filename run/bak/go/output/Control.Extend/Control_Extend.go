package Control_Extend

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Functor "gopurs/output/Data.Functor"
)

var cache_identity gopurs_runtime.Value
var once_identity sync.Once
func Get_identity() gopurs_runtime.Value {
	once_identity.Do(func() {
		cache_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_identity(gopurs_runtime.UnboxAny(x_0_box)))
})
	})
	return cache_identity
}

var cache_extendFn gopurs_runtime.Value
var once_extendFn sync.Once
func Get_extendFn() gopurs_runtime.Value {
	once_extendFn.Do(func() {
		cache_extendFn = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_extendFn(dictSemigroup_0_box))
})
	})
	return cache_extendFn
}

var cache_extendArray gopurs_runtime.Value
var once_extendArray sync.Once
func Get_extendArray() gopurs_runtime.Value {
	once_extendArray.Do(func() {
		cache_extendArray = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorArray()
}), Get_arrayExtend())))
	})
	return cache_extendArray
}

var cache_extend gopurs_runtime.Value
var once_extend sync.Once
func Get_extend() gopurs_runtime.Value {
	once_extend.Do(func() {
		cache_extend = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_extend(dict_0_box)
})
	})
	return cache_extend
}

var cache_extend__func_gopurs_runtime_Value__func_interface____interface____interface____interface___186489329 gopurs_runtime.Value
var once_extend__func_gopurs_runtime_Value__func_interface____interface____interface____interface___186489329 sync.Once
func Get_extend__func_gopurs_runtime_Value__func_interface____interface____interface____interface___186489329() gopurs_runtime.Value {
	once_extend__func_gopurs_runtime_Value__func_interface____interface____interface____interface___186489329.Do(func() {
		cache_extend__func_gopurs_runtime_Value__func_interface____interface____interface____interface___186489329 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_extend__func_gopurs_runtime_Value__func_interface____interface____interface____interface___186489329(dict_0_box)
})
	})
	return cache_extend__func_gopurs_runtime_Value__func_interface____interface____interface____interface___186489329
}

var cache_extendFlipped gopurs_runtime.Value
var once_extendFlipped sync.Once
func Get_extendFlipped() gopurs_runtime.Value {
	once_extendFlipped.Do(func() {
		cache_extendFlipped = gopurs_runtime.Func3(func(dictExtend_0_box gopurs_runtime.Value, w_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_extendFlipped(dictExtend_0_box, gopurs_runtime.UnboxAny(w_1_box), func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_2_box, gopurs_runtime.Any(inner_arg0)))
}))
})
	})
	return cache_extendFlipped
}

var cache_duplicate gopurs_runtime.Value
var once_duplicate sync.Once
func Get_duplicate() gopurs_runtime.Value {
	once_duplicate.Do(func() {
		cache_duplicate = gopurs_runtime.Func(func(dictExtend_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_duplicate(dictExtend_0_box)
})
	})
	return cache_duplicate
}

var cache_composeCoKleisliFlipped gopurs_runtime.Value
var once_composeCoKleisliFlipped sync.Once
func Get_composeCoKleisliFlipped() gopurs_runtime.Value {
	once_composeCoKleisliFlipped.Do(func() {
		cache_composeCoKleisliFlipped = gopurs_runtime.Func4(func(dictExtend_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value, w_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_composeCoKleisliFlipped(dictExtend_0_box, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_1_box, gopurs_runtime.Any(inner_arg0)))
}, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(g_2_box, gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(w_3_box)))
})
	})
	return cache_composeCoKleisliFlipped
}

var cache_composeCoKleisliFlipped__func_gopurs_runtime_Value__func_interface____interface____func_interface____interface____interface____interface___3863740592 gopurs_runtime.Value
var once_composeCoKleisliFlipped__func_gopurs_runtime_Value__func_interface____interface____func_interface____interface____interface____interface___3863740592 sync.Once
func Get_composeCoKleisliFlipped__func_gopurs_runtime_Value__func_interface____interface____func_interface____interface____interface____interface___3863740592() gopurs_runtime.Value {
	once_composeCoKleisliFlipped__func_gopurs_runtime_Value__func_interface____interface____func_interface____interface____interface____interface___3863740592.Do(func() {
		cache_composeCoKleisliFlipped__func_gopurs_runtime_Value__func_interface____interface____func_interface____interface____interface____interface___3863740592 = gopurs_runtime.Func4(func(dictExtend_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value, w_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_composeCoKleisliFlipped__func_gopurs_runtime_Value__func_interface____interface____func_interface____interface____interface____interface___3863740592(dictExtend_0_box, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_1_box, gopurs_runtime.Any(inner_arg0)))
}, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(g_2_box, gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(w_3_box)))
})
	})
	return cache_composeCoKleisliFlipped__func_gopurs_runtime_Value__func_interface____interface____func_interface____interface____interface____interface___3863740592
}

var cache_composeCoKleisli gopurs_runtime.Value
var once_composeCoKleisli sync.Once
func Get_composeCoKleisli() gopurs_runtime.Value {
	once_composeCoKleisli.Do(func() {
		cache_composeCoKleisli = gopurs_runtime.Func4(func(dictExtend_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value, w_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_composeCoKleisli(dictExtend_0_box, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_1_box, gopurs_runtime.Any(inner_arg0)))
}, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(g_2_box, gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(w_3_box)))
})
	})
	return cache_composeCoKleisli
}

var cache_arrayExtend gopurs_runtime.Value
var once_arrayExtend sync.Once
func Get_arrayExtend() gopurs_runtime.Value {
	once_arrayExtend.Do(func() {
		cache_arrayExtend = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := ArrayExtend(func(inner_arg0 []interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(arg0, func() gopurs_runtime.Value {
					arr := inner_arg0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()))
}, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(arg1.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}())
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_arrayExtend
}

func Call_identity(x_0_loop interface{}) interface{} {
var x_0 interface{} = x_0_loop
_ = x_0
return x_0
}

func Call_extendFn(dictSemigroup_0_loop gopurs_runtime.Value) interface{} {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorFn()
}), gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, g_2 gopurs_runtime.Value, w_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, gopurs_runtime.Func(func(w_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(g_2, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), w_3, w_prime_4))
}))
})))
}

func Call_extend(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "extend")
}

func Call_extend__func_gopurs_runtime_Value__func_interface____interface____interface____interface___186489329(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "extend")
}

func Call_extendFlipped(dictExtend_0_loop gopurs_runtime.Value, w_1_loop interface{}, f_2_loop func(interface{}) interface{}) interface{} {
var dictExtend_0 gopurs_runtime.Value = dictExtend_0_loop
_ = dictExtend_0
var w_1 interface{} = w_1_loop
_ = w_1
var f_2 func(interface{}) interface{} = f_2_loop
_ = f_2
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictExtend_0, "extend"), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(f_2(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Any(w_1)))
}

func Call_duplicate(dictExtend_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictExtend_0 gopurs_runtime.Value = dictExtend_0_loop
_ = dictExtend_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend_0, "extend"), Get_identity())
}

func Call_composeCoKleisliFlipped(dictExtend_0_loop gopurs_runtime.Value, f_1_loop func(interface{}) interface{}, g_2_loop func(interface{}) interface{}, w_3_loop interface{}) interface{} {
var dictExtend_0 gopurs_runtime.Value = dictExtend_0_loop
_ = dictExtend_0
var f_1 func(interface{}) interface{} = f_1_loop
_ = f_1
var g_2 func(interface{}) interface{} = g_2_loop
_ = g_2
var w_3 interface{} = w_3_loop
_ = w_3
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(f_1(gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictExtend_0, "extend"), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(g_2(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Any(w_3))))))
}

func Call_composeCoKleisliFlipped__func_gopurs_runtime_Value__func_interface____interface____func_interface____interface____interface____interface___3863740592(dictExtend_0_loop gopurs_runtime.Value, f_1_loop func(interface{}) interface{}, g_2_loop func(interface{}) interface{}, w_3_loop interface{}) interface{} {
var dictExtend_0 gopurs_runtime.Value = dictExtend_0_loop
_ = dictExtend_0
var f_1 func(interface{}) interface{} = f_1_loop
_ = f_1
var g_2 func(interface{}) interface{} = g_2_loop
_ = g_2
var w_3 interface{} = w_3_loop
_ = w_3
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(f_1(gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictExtend_0, "extend"), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(g_2(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Any(w_3))))))
}

func Call_composeCoKleisli(dictExtend_0_loop gopurs_runtime.Value, f_1_loop func(interface{}) interface{}, g_2_loop func(interface{}) interface{}, w_3_loop interface{}) interface{} {
var dictExtend_0 gopurs_runtime.Value = dictExtend_0_loop
_ = dictExtend_0
var f_1 func(interface{}) interface{} = f_1_loop
_ = f_1
var g_2 func(interface{}) interface{} = g_2_loop
_ = g_2
var w_3 interface{} = w_3_loop
_ = w_3
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(g_2(gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictExtend_0, "extend"), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(f_1(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Any(w_3))))))
}
