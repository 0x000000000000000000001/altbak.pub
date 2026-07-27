package Data_Function

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Ord "gopurs/output/Data.Ord"
)

var cache_lessThanOrEq gopurs_runtime.Value
var once_lessThanOrEq sync.Once
func Get_lessThanOrEq() gopurs_runtime.Value {
	once_lessThanOrEq.Do(func() {
		cache_lessThanOrEq = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}))
_ = __local_var_0_0
return gopurs_runtime.Func2(func(a1_1 gopurs_runtime.Value, a2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Apply2(__local_var_0_0, a1_1, a2_2)
return gopurs_runtime.Bool(((__t_tag_1.Type == 9 && __t_tag_1.IntVal == 380165415)) != (true))
})
}()
	})
	return cache_lessThanOrEq
}

var cache_on gopurs_runtime.Value
var once_on sync.Once
func Get_on() gopurs_runtime.Value {
	once_on.Do(func() {
		cache_on = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_on(func(inner_arg0 interface{}, inner_arg1 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(f_0_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1)))
}, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(g_1_box, gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(x_2_box), gopurs_runtime.UnboxAny(y_3_box)))
})
	})
	return cache_on
}

var cache_on__func_func_interface____interface____bool__func_interface____interface____interface____interface____bool_1062875992 gopurs_runtime.Value
var once_on__func_func_interface____interface____bool__func_interface____interface____interface____interface____bool_1062875992 sync.Once
func Get_on__func_func_interface____interface____bool__func_interface____interface____interface____interface____bool_1062875992() gopurs_runtime.Value {
	once_on__func_func_interface____interface____bool__func_interface____interface____interface____interface____bool_1062875992.Do(func() {
		cache_on__func_func_interface____interface____bool__func_interface____interface____interface____interface____bool_1062875992 = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_on__func_func_interface____interface____bool__func_interface____interface____interface____interface____bool_1062875992(func(inner_arg0 interface{}, inner_arg1 interface{}) bool {
return (gopurs_runtime.Apply2(f_0_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1)).IntVal) != (0)
}, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(g_1_box, gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(x_2_box), gopurs_runtime.UnboxAny(y_3_box)))
})
	})
	return cache_on__func_func_interface____interface____bool__func_interface____interface____interface____interface____bool_1062875992
}

var cache_on__func_func_interface____interface____interface____func_interface____interface____interface____interface____interface___2003439736 gopurs_runtime.Value
var once_on__func_func_interface____interface____interface____func_interface____interface____interface____interface____interface___2003439736 sync.Once
func Get_on__func_func_interface____interface____interface____func_interface____interface____interface____interface____interface___2003439736() gopurs_runtime.Value {
	once_on__func_func_interface____interface____interface____func_interface____interface____interface____interface____interface___2003439736.Do(func() {
		cache_on__func_func_interface____interface____interface____func_interface____interface____interface____interface____interface___2003439736 = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_on__func_func_interface____interface____interface____func_interface____interface____interface____interface____interface___2003439736(func(inner_arg0 interface{}, inner_arg1 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(f_0_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1)))
}, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(g_1_box, gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(x_2_box), gopurs_runtime.UnboxAny(y_3_box)))
})
	})
	return cache_on__func_func_interface____interface____interface____func_interface____interface____interface____interface____interface___2003439736
}

var cache_flip gopurs_runtime.Value
var once_flip sync.Once
func Get_flip() gopurs_runtime.Value {
	once_flip.Do(func() {
		cache_flip = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_flip(func(inner_arg0 interface{}, inner_arg1 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(f_0_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1)))
}, gopurs_runtime.UnboxAny(b_1_box), gopurs_runtime.UnboxAny(a_2_box)))
})
	})
	return cache_flip
}

var cache_flip__func_func_arrinterface____interface____arrinterface____interface____arrinterface____arrinterface___2010786042 gopurs_runtime.Value
var once_flip__func_func_arrinterface____interface____arrinterface____interface____arrinterface____arrinterface___2010786042 sync.Once
func Get_flip__func_func_arrinterface____interface____arrinterface____interface____arrinterface____arrinterface___2010786042() gopurs_runtime.Value {
	once_flip__func_func_arrinterface____interface____arrinterface____interface____arrinterface____arrinterface___2010786042.Do(func() {
		cache_flip__func_func_arrinterface____interface____arrinterface____interface____arrinterface____arrinterface___2010786042 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_flip__func_func_arrinterface____interface____arrinterface____interface____arrinterface____arrinterface___2010786042(func(inner_arg0 []interface{}, inner_arg1 interface{}) []interface{} {
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(f_0_box, func() gopurs_runtime.Value {
					arr := inner_arg0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}(), gopurs_runtime.Any(inner_arg1)).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}, gopurs_runtime.UnboxAny(b_1_box), func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(a_2_box.UnsafePtr)
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
	return cache_flip__func_func_arrinterface____interface____arrinterface____interface____arrinterface____arrinterface___2010786042
}

var cache_flip__func_func_interface____arrinterface____arrinterface____arrinterface____interface____arrinterface___662810362 gopurs_runtime.Value
var once_flip__func_func_interface____arrinterface____arrinterface____arrinterface____interface____arrinterface___662810362 sync.Once
func Get_flip__func_func_interface____arrinterface____arrinterface____arrinterface____interface____arrinterface___662810362() gopurs_runtime.Value {
	once_flip__func_func_interface____arrinterface____arrinterface____arrinterface____interface____arrinterface___662810362.Do(func() {
		cache_flip__func_func_interface____arrinterface____arrinterface____arrinterface____interface____arrinterface___662810362 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_flip__func_func_interface____arrinterface____arrinterface____arrinterface____interface____arrinterface___662810362(func(inner_arg0 interface{}, inner_arg1 []interface{}) []interface{} {
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(f_0_box, gopurs_runtime.Any(inner_arg0), func() gopurs_runtime.Value {
					arr := inner_arg1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(b_1_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}(), gopurs_runtime.UnboxAny(a_2_box))
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_flip__func_func_interface____arrinterface____arrinterface____arrinterface____interface____arrinterface___662810362
}

var cache_flip__func_func_interface____interface____interface____interface____interface____interface___3946185210 gopurs_runtime.Value
var once_flip__func_func_interface____interface____interface____interface____interface____interface___3946185210 sync.Once
func Get_flip__func_func_interface____interface____interface____interface____interface____interface___3946185210() gopurs_runtime.Value {
	once_flip__func_func_interface____interface____interface____interface____interface____interface___3946185210.Do(func() {
		cache_flip__func_func_interface____interface____interface____interface____interface____interface___3946185210 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_flip__func_func_interface____interface____interface____interface____interface____interface___3946185210(func(inner_arg0 interface{}, inner_arg1 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(f_0_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1)))
}, gopurs_runtime.UnboxAny(b_1_box), gopurs_runtime.UnboxAny(a_2_box)))
})
	})
	return cache_flip__func_func_interface____interface____interface____interface____interface____interface___3946185210
}

var cache_flip__func_func_interface____interface____interface____interface____interface____interface___2301875002 gopurs_runtime.Value
var once_flip__func_func_interface____interface____interface____interface____interface____interface___2301875002 sync.Once
func Get_flip__func_func_interface____interface____interface____interface____interface____interface___2301875002() gopurs_runtime.Value {
	once_flip__func_func_interface____interface____interface____interface____interface____interface___2301875002.Do(func() {
		cache_flip__func_func_interface____interface____interface____interface____interface____interface___2301875002 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_flip__func_func_interface____interface____interface____interface____interface____interface___2301875002(func(inner_arg0 interface{}, inner_arg1 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(f_0_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1)))
}, gopurs_runtime.UnboxAny(b_1_box), gopurs_runtime.UnboxAny(a_2_box)))
})
	})
	return cache_flip__func_func_interface____interface____interface____interface____interface____interface___2301875002
}

var cache_flip__func_func_interface____interface____interface____interface____interface____interface___1042884218 gopurs_runtime.Value
var once_flip__func_func_interface____interface____interface____interface____interface____interface___1042884218 sync.Once
func Get_flip__func_func_interface____interface____interface____interface____interface____interface___1042884218() gopurs_runtime.Value {
	once_flip__func_func_interface____interface____interface____interface____interface____interface___1042884218.Do(func() {
		cache_flip__func_func_interface____interface____interface____interface____interface____interface___1042884218 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_flip__func_func_interface____interface____interface____interface____interface____interface___1042884218(func(inner_arg0 interface{}, inner_arg1 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(f_0_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1)))
}, gopurs_runtime.UnboxAny(b_1_box), gopurs_runtime.UnboxAny(a_2_box)))
})
	})
	return cache_flip__func_func_interface____interface____interface____interface____interface____interface___1042884218
}

var cache_flip__func_func_interface____interface____interface____interface____interface____interface___1841051514 gopurs_runtime.Value
var once_flip__func_func_interface____interface____interface____interface____interface____interface___1841051514 sync.Once
func Get_flip__func_func_interface____interface____interface____interface____interface____interface___1841051514() gopurs_runtime.Value {
	once_flip__func_func_interface____interface____interface____interface____interface____interface___1841051514.Do(func() {
		cache_flip__func_func_interface____interface____interface____interface____interface____interface___1841051514 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_flip__func_func_interface____interface____interface____interface____interface____interface___1841051514(func(inner_arg0 interface{}, inner_arg1 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(f_0_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1)))
}, gopurs_runtime.UnboxAny(b_1_box), gopurs_runtime.UnboxAny(a_2_box)))
})
	})
	return cache_flip__func_func_interface____interface____interface____interface____interface____interface___1841051514
}

var cache_flip__func_func_interface____interface____interface____interface____interface____interface___1752871994 gopurs_runtime.Value
var once_flip__func_func_interface____interface____interface____interface____interface____interface___1752871994 sync.Once
func Get_flip__func_func_interface____interface____interface____interface____interface____interface___1752871994() gopurs_runtime.Value {
	once_flip__func_func_interface____interface____interface____interface____interface____interface___1752871994.Do(func() {
		cache_flip__func_func_interface____interface____interface____interface____interface____interface___1752871994 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_flip__func_func_interface____interface____interface____interface____interface____interface___1752871994(func(inner_arg0 interface{}, inner_arg1 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(f_0_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1)))
}, gopurs_runtime.UnboxAny(b_1_box), gopurs_runtime.UnboxAny(a_2_box)))
})
	})
	return cache_flip__func_func_interface____interface____interface____interface____interface____interface___1752871994
}

var cache_flip__func_func_interface____interface____interface____interface____interface____interface___2551194106 gopurs_runtime.Value
var once_flip__func_func_interface____interface____interface____interface____interface____interface___2551194106 sync.Once
func Get_flip__func_func_interface____interface____interface____interface____interface____interface___2551194106() gopurs_runtime.Value {
	once_flip__func_func_interface____interface____interface____interface____interface____interface___2551194106.Do(func() {
		cache_flip__func_func_interface____interface____interface____interface____interface____interface___2551194106 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_flip__func_func_interface____interface____interface____interface____interface____interface___2551194106(func(inner_arg0 interface{}, inner_arg1 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(f_0_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1)))
}, gopurs_runtime.UnboxAny(b_1_box), gopurs_runtime.UnboxAny(a_2_box)))
})
	})
	return cache_flip__func_func_interface____interface____interface____interface____interface____interface___2551194106
}

var cache_flip__func_func_interface____interface____interface____interface____interface____interface___2108834362 gopurs_runtime.Value
var once_flip__func_func_interface____interface____interface____interface____interface____interface___2108834362 sync.Once
func Get_flip__func_func_interface____interface____interface____interface____interface____interface___2108834362() gopurs_runtime.Value {
	once_flip__func_func_interface____interface____interface____interface____interface____interface___2108834362.Do(func() {
		cache_flip__func_func_interface____interface____interface____interface____interface____interface___2108834362 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_flip__func_func_interface____interface____interface____interface____interface____interface___2108834362(func(inner_arg0 interface{}, inner_arg1 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(f_0_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1)))
}, gopurs_runtime.UnboxAny(b_1_box), gopurs_runtime.UnboxAny(a_2_box)))
})
	})
	return cache_flip__func_func_interface____interface____interface____interface____interface____interface___2108834362
}

var cache_flip__func_func_interface____interface____interface____interface____interface____interface___3361344762 gopurs_runtime.Value
var once_flip__func_func_interface____interface____interface____interface____interface____interface___3361344762 sync.Once
func Get_flip__func_func_interface____interface____interface____interface____interface____interface___3361344762() gopurs_runtime.Value {
	once_flip__func_func_interface____interface____interface____interface____interface____interface___3361344762.Do(func() {
		cache_flip__func_func_interface____interface____interface____interface____interface____interface___3361344762 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_flip__func_func_interface____interface____interface____interface____interface____interface___3361344762(func(inner_arg0 interface{}, inner_arg1 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(f_0_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1)))
}, gopurs_runtime.UnboxAny(b_1_box), gopurs_runtime.UnboxAny(a_2_box)))
})
	})
	return cache_flip__func_func_interface____interface____interface____interface____interface____interface___3361344762
}

var cache_flip__func_func_interface____interface____interface____interface____interface____interface___1114146170 gopurs_runtime.Value
var once_flip__func_func_interface____interface____interface____interface____interface____interface___1114146170 sync.Once
func Get_flip__func_func_interface____interface____interface____interface____interface____interface___1114146170() gopurs_runtime.Value {
	once_flip__func_func_interface____interface____interface____interface____interface____interface___1114146170.Do(func() {
		cache_flip__func_func_interface____interface____interface____interface____interface____interface___1114146170 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_flip__func_func_interface____interface____interface____interface____interface____interface___1114146170(func(inner_arg0 interface{}, inner_arg1 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(f_0_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1)))
}, gopurs_runtime.UnboxAny(b_1_box), gopurs_runtime.UnboxAny(a_2_box)))
})
	})
	return cache_flip__func_func_interface____interface____interface____interface____interface____interface___1114146170
}

var cache_flip__func_func_interface____interface____interface____interface____interface____interface___2645119546 gopurs_runtime.Value
var once_flip__func_func_interface____interface____interface____interface____interface____interface___2645119546 sync.Once
func Get_flip__func_func_interface____interface____interface____interface____interface____interface___2645119546() gopurs_runtime.Value {
	once_flip__func_func_interface____interface____interface____interface____interface____interface___2645119546.Do(func() {
		cache_flip__func_func_interface____interface____interface____interface____interface____interface___2645119546 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_flip__func_func_interface____interface____interface____interface____interface____interface___2645119546(func(inner_arg0 interface{}, inner_arg1 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(f_0_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1)))
}, gopurs_runtime.UnboxAny(b_1_box), gopurs_runtime.UnboxAny(a_2_box)))
})
	})
	return cache_flip__func_func_interface____interface____interface____interface____interface____interface___2645119546
}

var cache_const_ gopurs_runtime.Value
var once_const_ sync.Once
func Get_const_() gopurs_runtime.Value {
	once_const_.Do(func() {
		cache_const_ = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_const_(gopurs_runtime.UnboxAny(a_0_box), gopurs_runtime.UnboxAny(v_1_box)))
})
	})
	return cache_const_
}

var cache_const__func_int64__interface____int64_1014520284 gopurs_runtime.Value
var once_const__func_int64__interface____int64_1014520284 sync.Once
func Get_const__func_int64__interface____int64_1014520284() gopurs_runtime.Value {
	once_const__func_int64__interface____int64_1014520284.Do(func() {
		cache_const__func_int64__interface____int64_1014520284 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_const__func_int64__interface____int64_1014520284(a_0_box.IntVal, gopurs_runtime.UnboxAny(v_1_box)))
})
	})
	return cache_const__func_int64__interface____int64_1014520284
}

var cache_const__func_bool__interface____bool_3313219100 gopurs_runtime.Value
var once_const__func_bool__interface____bool_3313219100 sync.Once
func Get_const__func_bool__interface____bool_3313219100() gopurs_runtime.Value {
	once_const__func_bool__interface____bool_3313219100.Do(func() {
		cache_const__func_bool__interface____bool_3313219100 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_const__func_bool__interface____bool_3313219100((a_0_box.IntVal) != (0), gopurs_runtime.UnboxAny(v_1_box)))
})
	})
	return cache_const__func_bool__interface____bool_3313219100
}

var cache_const__func_func_interface____interface____interface____func_interface____interface___2448935612 gopurs_runtime.Value
var once_const__func_func_interface____interface____interface____func_interface____interface___2448935612 sync.Once
func Get_const__func_func_interface____interface____interface____func_interface____interface___2448935612() gopurs_runtime.Value {
	once_const__func_func_interface____interface____interface____func_interface____interface___2448935612.Do(func() {
		cache_const__func_func_interface____interface____interface____func_interface____interface___2448935612 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__func_func_interface____interface____interface____func_interface____interface___2448935612(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(a_0_box, gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(v_1_box))
})
	})
	return cache_const__func_func_interface____interface____interface____func_interface____interface___2448935612
}

var cache_const__func_func_interface____interface____interface____func_interface____interface___827936540 gopurs_runtime.Value
var once_const__func_func_interface____interface____interface____func_interface____interface___827936540 sync.Once
func Get_const__func_func_interface____interface____interface____func_interface____interface___827936540() gopurs_runtime.Value {
	once_const__func_func_interface____interface____interface____func_interface____interface___827936540.Do(func() {
		cache_const__func_func_interface____interface____interface____func_interface____interface___827936540 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__func_func_interface____interface____interface____func_interface____interface___827936540(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(a_0_box, gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(v_1_box))
})
	})
	return cache_const__func_func_interface____interface____interface____func_interface____interface___827936540
}

var cache_const__func_func_interface____interface____interface____interface____func_interface____interface____interface___2737562588 gopurs_runtime.Value
var once_const__func_func_interface____interface____interface____interface____func_interface____interface____interface___2737562588 sync.Once
func Get_const__func_func_interface____interface____interface____interface____func_interface____interface____interface___2737562588() gopurs_runtime.Value {
	once_const__func_func_interface____interface____interface____interface____func_interface____interface____interface___2737562588.Do(func() {
		cache_const__func_func_interface____interface____interface____interface____func_interface____interface____interface___2737562588 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__func_func_interface____interface____interface____interface____func_interface____interface____interface___2737562588(func(inner_arg0 interface{}, inner_arg1 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(a_0_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1)))
}, gopurs_runtime.UnboxAny(v_1_box))
})
	})
	return cache_const__func_func_interface____interface____interface____interface____func_interface____interface____interface___2737562588
}

var cache_const__func_func_interface____interface____interface____interface____func_interface____interface____interface___2901144028 gopurs_runtime.Value
var once_const__func_func_interface____interface____interface____interface____func_interface____interface____interface___2901144028 sync.Once
func Get_const__func_func_interface____interface____interface____interface____func_interface____interface____interface___2901144028() gopurs_runtime.Value {
	once_const__func_func_interface____interface____interface____interface____func_interface____interface____interface___2901144028.Do(func() {
		cache_const__func_func_interface____interface____interface____interface____func_interface____interface____interface___2901144028 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__func_func_interface____interface____interface____interface____func_interface____interface____interface___2901144028(func(inner_arg0 interface{}, inner_arg1 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(a_0_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1)))
}, gopurs_runtime.UnboxAny(v_1_box))
})
	})
	return cache_const__func_func_interface____interface____interface____interface____func_interface____interface____interface___2901144028
}

var cache_const__func_func_interface____bool__interface____interface____bool_1533104436 gopurs_runtime.Value
var once_const__func_func_interface____bool__interface____interface____bool_1533104436 sync.Once
func Get_const__func_func_interface____bool__interface____interface____bool_1533104436() gopurs_runtime.Value {
	once_const__func_func_interface____bool__interface____interface____bool_1533104436.Do(func() {
		cache_const__func_func_interface____bool__interface____interface____bool_1533104436 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__func_func_interface____bool__interface____interface____bool_1533104436(func(inner_arg0 interface{}) bool {
return (gopurs_runtime.Apply(a_0_box, gopurs_runtime.Any(inner_arg0)).IntVal) != (0)
}, gopurs_runtime.UnboxAny(v_1_box))
})
	})
	return cache_const__func_func_interface____bool__interface____interface____bool_1533104436
}

var cache_const__func_interface____func_interface____interface____interface____interface___4266931454 gopurs_runtime.Value
var once_const__func_interface____func_interface____interface____interface____interface___4266931454 sync.Once
func Get_const__func_interface____func_interface____interface____interface____interface___4266931454() gopurs_runtime.Value {
	once_const__func_interface____func_interface____interface____interface____interface___4266931454.Do(func() {
		cache_const__func_interface____func_interface____interface____interface____interface___4266931454 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_const__func_interface____func_interface____interface____interface____interface___4266931454(gopurs_runtime.UnboxAny(a_0_box), func(inner_arg0 interface{}, inner_arg1 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(v_1_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1)))
}))
})
	})
	return cache_const__func_interface____func_interface____interface____interface____interface___4266931454
}

var cache_const__func_interface____interface____interface___87201948 gopurs_runtime.Value
var once_const__func_interface____interface____interface___87201948 sync.Once
func Get_const__func_interface____interface____interface___87201948() gopurs_runtime.Value {
	once_const__func_interface____interface____interface___87201948.Do(func() {
		cache_const__func_interface____interface____interface___87201948 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_const__func_interface____interface____interface___87201948(gopurs_runtime.UnboxAny(a_0_box), gopurs_runtime.UnboxAny(v_1_box)))
})
	})
	return cache_const__func_interface____interface____interface___87201948
}

var cache_const__func_interface____interface____interface___3060642877 gopurs_runtime.Value
var once_const__func_interface____interface____interface___3060642877 sync.Once
func Get_const__func_interface____interface____interface___3060642877() gopurs_runtime.Value {
	once_const__func_interface____interface____interface___3060642877.Do(func() {
		cache_const__func_interface____interface____interface___3060642877 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_const__func_interface____interface____interface___3060642877(gopurs_runtime.UnboxAny(a_0_box), gopurs_runtime.UnboxAny(v_1_box)))
})
	})
	return cache_const__func_interface____interface____interface___3060642877
}

var cache_const__func_interface____interface____interface___31775132 gopurs_runtime.Value
var once_const__func_interface____interface____interface___31775132 sync.Once
func Get_const__func_interface____interface____interface___31775132() gopurs_runtime.Value {
	once_const__func_interface____interface____interface___31775132.Do(func() {
		cache_const__func_interface____interface____interface___31775132 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_const__func_interface____interface____interface___31775132(gopurs_runtime.UnboxAny(a_0_box), gopurs_runtime.UnboxAny(v_1_box)))
})
	})
	return cache_const__func_interface____interface____interface___31775132
}

var cache_const__func_interface____interface____interface___4379098 gopurs_runtime.Value
var once_const__func_interface____interface____interface___4379098 sync.Once
func Get_const__func_interface____interface____interface___4379098() gopurs_runtime.Value {
	once_const__func_interface____interface____interface___4379098.Do(func() {
		cache_const__func_interface____interface____interface___4379098 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_const__func_interface____interface____interface___4379098(gopurs_runtime.UnboxAny(a_0_box), gopurs_runtime.UnboxAny(v_1_box)))
})
	})
	return cache_const__func_interface____interface____interface___4379098
}

var cache_const__func_interface____interface____interface___945658268 gopurs_runtime.Value
var once_const__func_interface____interface____interface___945658268 sync.Once
func Get_const__func_interface____interface____interface___945658268() gopurs_runtime.Value {
	once_const__func_interface____interface____interface___945658268.Do(func() {
		cache_const__func_interface____interface____interface___945658268 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_const__func_interface____interface____interface___945658268(gopurs_runtime.UnboxAny(a_0_box), gopurs_runtime.UnboxAny(v_1_box)))
})
	})
	return cache_const__func_interface____interface____interface___945658268
}

var cache_const__func_interface____interface____interface___4007400604 gopurs_runtime.Value
var once_const__func_interface____interface____interface___4007400604 sync.Once
func Get_const__func_interface____interface____interface___4007400604() gopurs_runtime.Value {
	once_const__func_interface____interface____interface___4007400604.Do(func() {
		cache_const__func_interface____interface____interface___4007400604 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_const__func_interface____interface____interface___4007400604(gopurs_runtime.UnboxAny(a_0_box), gopurs_runtime.UnboxAny(v_1_box)))
})
	})
	return cache_const__func_interface____interface____interface___4007400604
}

var cache_const__func_interface____interface____interface___144291955 gopurs_runtime.Value
var once_const__func_interface____interface____interface___144291955 sync.Once
func Get_const__func_interface____interface____interface___144291955() gopurs_runtime.Value {
	once_const__func_interface____interface____interface___144291955.Do(func() {
		cache_const__func_interface____interface____interface___144291955 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_const__func_interface____interface____interface___144291955(gopurs_runtime.UnboxAny(a_0_box), gopurs_runtime.UnboxAny(v_1_box)))
})
	})
	return cache_const__func_interface____interface____interface___144291955
}

var cache_const__func_interface____interface____interface___4005387916 gopurs_runtime.Value
var once_const__func_interface____interface____interface___4005387916 sync.Once
func Get_const__func_interface____interface____interface___4005387916() gopurs_runtime.Value {
	once_const__func_interface____interface____interface___4005387916.Do(func() {
		cache_const__func_interface____interface____interface___4005387916 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_const__func_interface____interface____interface___4005387916(gopurs_runtime.UnboxAny(a_0_box), gopurs_runtime.UnboxAny(v_1_box)))
})
	})
	return cache_const__func_interface____interface____interface___4005387916
}

var cache_const__func_interface____interface____interface___3293692348 gopurs_runtime.Value
var once_const__func_interface____interface____interface___3293692348 sync.Once
func Get_const__func_interface____interface____interface___3293692348() gopurs_runtime.Value {
	once_const__func_interface____interface____interface___3293692348.Do(func() {
		cache_const__func_interface____interface____interface___3293692348 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_const__func_interface____interface____interface___3293692348(gopurs_runtime.UnboxAny(a_0_box), gopurs_runtime.UnboxAny(v_1_box)))
})
	})
	return cache_const__func_interface____interface____interface___3293692348
}

var cache_const__func_interface____interface____interface___2326386184 gopurs_runtime.Value
var once_const__func_interface____interface____interface___2326386184 sync.Once
func Get_const__func_interface____interface____interface___2326386184() gopurs_runtime.Value {
	once_const__func_interface____interface____interface___2326386184.Do(func() {
		cache_const__func_interface____interface____interface___2326386184 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_const__func_interface____interface____interface___2326386184(gopurs_runtime.UnboxAny(a_0_box), gopurs_runtime.UnboxAny(v_1_box)))
})
	})
	return cache_const__func_interface____interface____interface___2326386184
}

var cache_applyN gopurs_runtime.Value
var once_applyN sync.Once
func Get_applyN() gopurs_runtime.Value {
	once_applyN.Do(func() {
		cache_applyN = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applyN(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_0_box, gopurs_runtime.Any(inner_arg0)))
})
})
	})
	return cache_applyN
}

var cache_applyFlipped gopurs_runtime.Value
var once_applyFlipped sync.Once
func Get_applyFlipped() gopurs_runtime.Value {
	once_applyFlipped.Do(func() {
		cache_applyFlipped = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_applyFlipped(gopurs_runtime.UnboxAny(x_0_box), func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_1_box, gopurs_runtime.Any(inner_arg0)))
}))
})
	})
	return cache_applyFlipped
}

var cache_apply gopurs_runtime.Value
var once_apply sync.Once
func Get_apply() gopurs_runtime.Value {
	once_apply.Do(func() {
		cache_apply = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_apply(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_0_box, gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(x_1_box)))
})
	})
	return cache_apply
}

var cache_apply__func_func_interface____interface____interface____interface___559335802 gopurs_runtime.Value
var once_apply__func_func_interface____interface____interface____interface___559335802 sync.Once
func Get_apply__func_func_interface____interface____interface____interface___559335802() gopurs_runtime.Value {
	once_apply__func_func_interface____interface____interface____interface___559335802.Do(func() {
		cache_apply__func_func_interface____interface____interface____interface___559335802 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_apply__func_func_interface____interface____interface____interface___559335802(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_0_box, gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(x_1_box)))
})
	})
	return cache_apply__func_func_interface____interface____interface____interface___559335802
}

func Call_on(f_0_loop func(interface{}, interface{}) interface{}, g_1_loop func(interface{}) interface{}, x_2_loop interface{}, y_3_loop interface{}) interface{} {
var f_0 func(interface{}, interface{}) interface{} = f_0_loop
_ = f_0
var g_1 func(interface{}) interface{} = g_1_loop
_ = g_1
var x_2 interface{} = x_2_loop
_ = x_2
var y_3 interface{} = y_3_loop
_ = y_3
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(f_0(gopurs_runtime.UnboxAny(gopurs_runtime.Any(g_1(x_2))), gopurs_runtime.UnboxAny(gopurs_runtime.Any(g_1(y_3))))))
}

func Call_on__func_func_interface____interface____bool__func_interface____interface____interface____interface____bool_1062875992(f_0_loop func(interface{}, interface{}) bool, g_1_loop func(interface{}) interface{}, x_2_loop interface{}, y_3_loop interface{}) bool {
var f_0 func(interface{}, interface{}) bool = f_0_loop
_ = f_0
var g_1 func(interface{}) interface{} = g_1_loop
_ = g_1
var x_2 interface{} = x_2_loop
_ = x_2
var y_3 interface{} = y_3_loop
_ = y_3
return (gopurs_runtime.Bool(f_0(gopurs_runtime.UnboxAny(gopurs_runtime.Any(g_1(x_2))), gopurs_runtime.UnboxAny(gopurs_runtime.Any(g_1(y_3))))).IntVal) != (0)
}

func Call_on__func_func_interface____interface____interface____func_interface____interface____interface____interface____interface___2003439736(f_0_loop func(interface{}, interface{}) interface{}, g_1_loop func(interface{}) interface{}, x_2_loop interface{}, y_3_loop interface{}) interface{} {
var f_0 func(interface{}, interface{}) interface{} = f_0_loop
_ = f_0
var g_1 func(interface{}) interface{} = g_1_loop
_ = g_1
var x_2 interface{} = x_2_loop
_ = x_2
var y_3 interface{} = y_3_loop
_ = y_3
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(f_0(gopurs_runtime.UnboxAny(gopurs_runtime.Any(g_1(x_2))), gopurs_runtime.UnboxAny(gopurs_runtime.Any(g_1(y_3))))))
}

func Call_flip(f_0_loop func(interface{}, interface{}) interface{}, b_1_loop interface{}, a_2_loop interface{}) interface{} {
var f_0 func(interface{}, interface{}) interface{} = f_0_loop
_ = f_0
var b_1 interface{} = b_1_loop
_ = b_1
var a_2 interface{} = a_2_loop
_ = a_2
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(f_0(a_2, b_1)))
}

func Call_flip__func_func_arrinterface____interface____arrinterface____interface____arrinterface____arrinterface___2010786042(f_0_loop func([]interface{}, interface{}) []interface{}, b_1_loop interface{}, a_2_loop []interface{}) []interface{} {
var f_0 func([]interface{}, interface{}) []interface{} = f_0_loop
_ = f_0
var b_1 interface{} = b_1_loop
_ = b_1
var a_2 []interface{} = a_2_loop
_ = a_2
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
					arr := f_0(a_2, b_1)
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}().UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}

func Call_flip__func_func_interface____arrinterface____arrinterface____arrinterface____interface____arrinterface___662810362(f_0_loop func(interface{}, []interface{}) []interface{}, b_1_loop []interface{}, a_2_loop interface{}) []interface{} {
var f_0 func(interface{}, []interface{}) []interface{} = f_0_loop
_ = f_0
var b_1 []interface{} = b_1_loop
_ = b_1
var a_2 interface{} = a_2_loop
_ = a_2
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
					arr := f_0(a_2, b_1)
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}().UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}

func Call_flip__func_func_interface____interface____interface____interface____interface____interface___3946185210(f_0_loop func(interface{}, interface{}) interface{}, b_1_loop interface{}, a_2_loop interface{}) interface{} {
var f_0 func(interface{}, interface{}) interface{} = f_0_loop
_ = f_0
var b_1 interface{} = b_1_loop
_ = b_1
var a_2 interface{} = a_2_loop
_ = a_2
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(f_0(a_2, b_1)))
}

func Call_flip__func_func_interface____interface____interface____interface____interface____interface___2301875002(f_0_loop func(interface{}, interface{}) interface{}, b_1_loop interface{}, a_2_loop interface{}) interface{} {
var f_0 func(interface{}, interface{}) interface{} = f_0_loop
_ = f_0
var b_1 interface{} = b_1_loop
_ = b_1
var a_2 interface{} = a_2_loop
_ = a_2
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(f_0(a_2, b_1)))
}

func Call_flip__func_func_interface____interface____interface____interface____interface____interface___1042884218(f_0_loop func(interface{}, interface{}) interface{}, b_1_loop interface{}, a_2_loop interface{}) interface{} {
var f_0 func(interface{}, interface{}) interface{} = f_0_loop
_ = f_0
var b_1 interface{} = b_1_loop
_ = b_1
var a_2 interface{} = a_2_loop
_ = a_2
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(f_0(a_2, b_1)))
}

func Call_flip__func_func_interface____interface____interface____interface____interface____interface___1841051514(f_0_loop func(interface{}, interface{}) interface{}, b_1_loop interface{}, a_2_loop interface{}) interface{} {
var f_0 func(interface{}, interface{}) interface{} = f_0_loop
_ = f_0
var b_1 interface{} = b_1_loop
_ = b_1
var a_2 interface{} = a_2_loop
_ = a_2
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(f_0(a_2, b_1)))
}

func Call_flip__func_func_interface____interface____interface____interface____interface____interface___1752871994(f_0_loop func(interface{}, interface{}) interface{}, b_1_loop interface{}, a_2_loop interface{}) interface{} {
var f_0 func(interface{}, interface{}) interface{} = f_0_loop
_ = f_0
var b_1 interface{} = b_1_loop
_ = b_1
var a_2 interface{} = a_2_loop
_ = a_2
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(f_0(a_2, b_1)))
}

func Call_flip__func_func_interface____interface____interface____interface____interface____interface___2551194106(f_0_loop func(interface{}, interface{}) interface{}, b_1_loop interface{}, a_2_loop interface{}) interface{} {
var f_0 func(interface{}, interface{}) interface{} = f_0_loop
_ = f_0
var b_1 interface{} = b_1_loop
_ = b_1
var a_2 interface{} = a_2_loop
_ = a_2
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(f_0(a_2, b_1)))
}

func Call_flip__func_func_interface____interface____interface____interface____interface____interface___2108834362(f_0_loop func(interface{}, interface{}) interface{}, b_1_loop interface{}, a_2_loop interface{}) interface{} {
var f_0 func(interface{}, interface{}) interface{} = f_0_loop
_ = f_0
var b_1 interface{} = b_1_loop
_ = b_1
var a_2 interface{} = a_2_loop
_ = a_2
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(f_0(a_2, b_1)))
}

func Call_flip__func_func_interface____interface____interface____interface____interface____interface___3361344762(f_0_loop func(interface{}, interface{}) interface{}, b_1_loop interface{}, a_2_loop interface{}) interface{} {
var f_0 func(interface{}, interface{}) interface{} = f_0_loop
_ = f_0
var b_1 interface{} = b_1_loop
_ = b_1
var a_2 interface{} = a_2_loop
_ = a_2
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(f_0(a_2, b_1)))
}

func Call_flip__func_func_interface____interface____interface____interface____interface____interface___1114146170(f_0_loop func(interface{}, interface{}) interface{}, b_1_loop interface{}, a_2_loop interface{}) interface{} {
var f_0 func(interface{}, interface{}) interface{} = f_0_loop
_ = f_0
var b_1 interface{} = b_1_loop
_ = b_1
var a_2 interface{} = a_2_loop
_ = a_2
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(f_0(a_2, b_1)))
}

func Call_flip__func_func_interface____interface____interface____interface____interface____interface___2645119546(f_0_loop func(interface{}, interface{}) interface{}, b_1_loop interface{}, a_2_loop interface{}) interface{} {
var f_0 func(interface{}, interface{}) interface{} = f_0_loop
_ = f_0
var b_1 interface{} = b_1_loop
_ = b_1
var a_2 interface{} = a_2_loop
_ = a_2
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(f_0(a_2, b_1)))
}

func Call_const_(a_0_loop interface{}, v_1_loop interface{}) interface{} {
var a_0 interface{} = a_0_loop
_ = a_0
var v_1 interface{} = v_1_loop
_ = v_1
return a_0
}

func Call_const__func_int64__interface____int64_1014520284(a_0_loop int64, v_1_loop interface{}) int64 {
var a_0 int64 = a_0_loop
_ = a_0
var v_1 interface{} = v_1_loop
_ = v_1
return a_0
}

func Call_const__func_bool__interface____bool_3313219100(a_0_loop bool, v_1_loop interface{}) bool {
var a_0 bool = a_0_loop
_ = a_0
var v_1 interface{} = v_1_loop
_ = v_1
return a_0
}

func Call_const__func_func_interface____interface____interface____func_interface____interface___2448935612(a_0_loop func(interface{}) interface{}, v_1_loop interface{}) gopurs_runtime.Value {
var a_0 func(interface{}) interface{} = a_0_loop
_ = a_0
var v_1 interface{} = v_1_loop
_ = v_1
return gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(a_0(gopurs_runtime.UnboxAny(arg0)))
})
}

func Call_const__func_func_interface____interface____interface____func_interface____interface___827936540(a_0_loop func(interface{}) interface{}, v_1_loop interface{}) gopurs_runtime.Value {
var a_0 func(interface{}) interface{} = a_0_loop
_ = a_0
var v_1 interface{} = v_1_loop
_ = v_1
return gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(a_0(gopurs_runtime.UnboxAny(arg0)))
})
}

func Call_const__func_func_interface____interface____interface____interface____func_interface____interface____interface___2737562588(a_0_loop func(interface{}, interface{}) interface{}, v_1_loop interface{}) gopurs_runtime.Value {
var a_0 func(interface{}, interface{}) interface{} = a_0_loop
_ = a_0
var v_1 interface{} = v_1_loop
_ = v_1
return gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(a_0(gopurs_runtime.UnboxAny(arg0), gopurs_runtime.UnboxAny(arg1)))
})
}

func Call_const__func_func_interface____interface____interface____interface____func_interface____interface____interface___2901144028(a_0_loop func(interface{}, interface{}) interface{}, v_1_loop interface{}) gopurs_runtime.Value {
var a_0 func(interface{}, interface{}) interface{} = a_0_loop
_ = a_0
var v_1 interface{} = v_1_loop
_ = v_1
return gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(a_0(gopurs_runtime.UnboxAny(arg0), gopurs_runtime.UnboxAny(arg1)))
})
}

func Call_const__func_func_interface____bool__interface____interface____bool_1533104436(a_0_loop func(interface{}) bool, v_1_loop interface{}) gopurs_runtime.Value {
var a_0 func(interface{}) bool = a_0_loop
_ = a_0
var v_1 interface{} = v_1_loop
_ = v_1
return gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(a_0(gopurs_runtime.UnboxAny(arg0)))
})
}

func Call_const__func_interface____func_interface____interface____interface____interface___4266931454(a_0_loop interface{}, v_1_loop func(interface{}, interface{}) interface{}) interface{} {
var a_0 interface{} = a_0_loop
_ = a_0
var v_1 func(interface{}, interface{}) interface{} = v_1_loop
_ = v_1
return a_0
}

func Call_const__func_interface____interface____interface___87201948(a_0_loop interface{}, v_1_loop interface{}) interface{} {
var a_0 interface{} = a_0_loop
_ = a_0
var v_1 interface{} = v_1_loop
_ = v_1
return a_0
}

func Call_const__func_interface____interface____interface___3060642877(a_0_loop interface{}, v_1_loop interface{}) interface{} {
var a_0 interface{} = a_0_loop
_ = a_0
var v_1 interface{} = v_1_loop
_ = v_1
return a_0
}

func Call_const__func_interface____interface____interface___31775132(a_0_loop interface{}, v_1_loop interface{}) interface{} {
var a_0 interface{} = a_0_loop
_ = a_0
var v_1 interface{} = v_1_loop
_ = v_1
return a_0
}

func Call_const__func_interface____interface____interface___4379098(a_0_loop interface{}, v_1_loop interface{}) interface{} {
var a_0 interface{} = a_0_loop
_ = a_0
var v_1 interface{} = v_1_loop
_ = v_1
return a_0
}

func Call_const__func_interface____interface____interface___945658268(a_0_loop interface{}, v_1_loop interface{}) interface{} {
var a_0 interface{} = a_0_loop
_ = a_0
var v_1 interface{} = v_1_loop
_ = v_1
return a_0
}

func Call_const__func_interface____interface____interface___4007400604(a_0_loop interface{}, v_1_loop interface{}) interface{} {
var a_0 interface{} = a_0_loop
_ = a_0
var v_1 interface{} = v_1_loop
_ = v_1
return a_0
}

func Call_const__func_interface____interface____interface___144291955(a_0_loop interface{}, v_1_loop interface{}) interface{} {
var a_0 interface{} = a_0_loop
_ = a_0
var v_1 interface{} = v_1_loop
_ = v_1
return a_0
}

func Call_const__func_interface____interface____interface___4005387916(a_0_loop interface{}, v_1_loop interface{}) interface{} {
var a_0 interface{} = a_0_loop
_ = a_0
var v_1 interface{} = v_1_loop
_ = v_1
return a_0
}

func Call_const__func_interface____interface____interface___3293692348(a_0_loop interface{}, v_1_loop interface{}) interface{} {
var a_0 interface{} = a_0_loop
_ = a_0
var v_1 interface{} = v_1_loop
_ = v_1
return a_0
}

func Call_const__func_interface____interface____interface___2326386184(a_0_loop interface{}, v_1_loop interface{}) interface{} {
var a_0 interface{} = a_0_loop
_ = a_0
var v_1 interface{} = v_1_loop
_ = v_1
return a_0
}

func Call_applyN(f_0_loop func(interface{}) interface{}) gopurs_runtime.Value {
var f_0 func(interface{}) interface{} = f_0_loop
_ = f_0
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(n_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var n_2_loop interface{} = gopurs_runtime.UnboxAny(n_2_loop_val)
var acc_3_loop interface{} = gopurs_runtime.UnboxAny(acc_3_loop_val)
go__1_0:
for {
if false { continue go__1_0 }
var n_2 interface{} = n_2_loop
_ = n_2
var acc_3 interface{} = acc_3_loop
_ = acc_3
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_lessThanOrEq(), n_2, gopurs_runtime.Int(0)).IntVal) != (0) {
__t1 = acc_3
goto end_branch_1
} else {

}
}
{
n_2_loop = gopurs_runtime.UnboxAny((n_2.IntVal) - (1))
acc_3_loop = gopurs_runtime.UnboxAny(gopurs_runtime.Any(f_0(gopurs_runtime.UnboxAny(acc_3))))
continue go__1_0
__t1 = gopurs_runtime.Value{}
}
end_branch_1:
return __t1
}
}()
})
})
return go__1_0
}

func Call_applyFlipped(x_0_loop interface{}, f_1_loop func(interface{}) interface{}) interface{} {
var x_0 interface{} = x_0_loop
_ = x_0
var f_1 func(interface{}) interface{} = f_1_loop
_ = f_1
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(f_1(x_0)))
}

func Call_apply(f_0_loop func(interface{}) interface{}, x_1_loop interface{}) interface{} {
var f_0 func(interface{}) interface{} = f_0_loop
_ = f_0
var x_1 interface{} = x_1_loop
_ = x_1
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(f_0(x_1)))
}

func Call_apply__func_func_interface____interface____interface____interface___559335802(f_0_loop func(interface{}) interface{}, x_1_loop interface{}) interface{} {
var f_0 func(interface{}) interface{} = f_0_loop
_ = f_0
var x_1 interface{} = x_1_loop
_ = x_1
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(f_0(x_1)))
}
