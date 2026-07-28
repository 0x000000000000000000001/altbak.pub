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
__local_var_0_0 := gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})
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
return Call_on(f_0_box, g_1_box, x_2_box, y_3_box)
})
	})
	return cache_on
}

var cache_on__gopurs_runtime_Value_1062875992 gopurs_runtime.Value
var once_on__gopurs_runtime_Value_1062875992 sync.Once
func Get_on__gopurs_runtime_Value_1062875992() gopurs_runtime.Value {
	once_on__gopurs_runtime_Value_1062875992.Do(func() {
		cache_on__gopurs_runtime_Value_1062875992 = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_on__gopurs_runtime_Value_1062875992(f_0_box, g_1_box, x_2_box, y_3_box))
})
	})
	return cache_on__gopurs_runtime_Value_1062875992
}

var cache_on__gopurs_runtime_Value_2003439736 gopurs_runtime.Value
var once_on__gopurs_runtime_Value_2003439736 sync.Once
func Get_on__gopurs_runtime_Value_2003439736() gopurs_runtime.Value {
	once_on__gopurs_runtime_Value_2003439736.Do(func() {
		cache_on__gopurs_runtime_Value_2003439736 = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_on__gopurs_runtime_Value_2003439736(f_0_box, g_1_box, x_2_box, y_3_box)
})
	})
	return cache_on__gopurs_runtime_Value_2003439736
}

var cache_flip gopurs_runtime.Value
var once_flip sync.Once
func Get_flip() gopurs_runtime.Value {
	once_flip.Do(func() {
		cache_flip = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip
}

var cache_flip__gopurs_runtime_Value_2010786042 gopurs_runtime.Value
var once_flip__gopurs_runtime_Value_2010786042 sync.Once
func Get_flip__gopurs_runtime_Value_2010786042() gopurs_runtime.Value {
	once_flip__gopurs_runtime_Value_2010786042.Do(func() {
		cache_flip__gopurs_runtime_Value_2010786042 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_flip__gopurs_runtime_Value_2010786042(f_0_box, b_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(a_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_flip__gopurs_runtime_Value_2010786042
}

var cache_flip__gopurs_runtime_Value_662810362 gopurs_runtime.Value
var once_flip__gopurs_runtime_Value_662810362 sync.Once
func Get_flip__gopurs_runtime_Value_662810362() gopurs_runtime.Value {
	once_flip__gopurs_runtime_Value_662810362.Do(func() {
		cache_flip__gopurs_runtime_Value_662810362 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_flip__gopurs_runtime_Value_662810362(f_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(b_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), a_2_box))
})
	})
	return cache_flip__gopurs_runtime_Value_662810362
}

var cache_flip__gopurs_runtime_Value_3946185210 gopurs_runtime.Value
var once_flip__gopurs_runtime_Value_3946185210 sync.Once
func Get_flip__gopurs_runtime_Value_3946185210() gopurs_runtime.Value {
	once_flip__gopurs_runtime_Value_3946185210.Do(func() {
		cache_flip__gopurs_runtime_Value_3946185210 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__gopurs_runtime_Value_3946185210(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__gopurs_runtime_Value_3946185210
}

var cache_flip__gopurs_runtime_Value_2301875002 gopurs_runtime.Value
var once_flip__gopurs_runtime_Value_2301875002 sync.Once
func Get_flip__gopurs_runtime_Value_2301875002() gopurs_runtime.Value {
	once_flip__gopurs_runtime_Value_2301875002.Do(func() {
		cache_flip__gopurs_runtime_Value_2301875002 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__gopurs_runtime_Value_2301875002(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__gopurs_runtime_Value_2301875002
}

var cache_flip__gopurs_runtime_Value_1042884218 gopurs_runtime.Value
var once_flip__gopurs_runtime_Value_1042884218 sync.Once
func Get_flip__gopurs_runtime_Value_1042884218() gopurs_runtime.Value {
	once_flip__gopurs_runtime_Value_1042884218.Do(func() {
		cache_flip__gopurs_runtime_Value_1042884218 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__gopurs_runtime_Value_1042884218(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__gopurs_runtime_Value_1042884218
}

var cache_flip__gopurs_runtime_Value_1841051514 gopurs_runtime.Value
var once_flip__gopurs_runtime_Value_1841051514 sync.Once
func Get_flip__gopurs_runtime_Value_1841051514() gopurs_runtime.Value {
	once_flip__gopurs_runtime_Value_1841051514.Do(func() {
		cache_flip__gopurs_runtime_Value_1841051514 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__gopurs_runtime_Value_1841051514(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__gopurs_runtime_Value_1841051514
}

var cache_flip__gopurs_runtime_Value_1752871994 gopurs_runtime.Value
var once_flip__gopurs_runtime_Value_1752871994 sync.Once
func Get_flip__gopurs_runtime_Value_1752871994() gopurs_runtime.Value {
	once_flip__gopurs_runtime_Value_1752871994.Do(func() {
		cache_flip__gopurs_runtime_Value_1752871994 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__gopurs_runtime_Value_1752871994(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__gopurs_runtime_Value_1752871994
}

var cache_flip__gopurs_runtime_Value_2551194106 gopurs_runtime.Value
var once_flip__gopurs_runtime_Value_2551194106 sync.Once
func Get_flip__gopurs_runtime_Value_2551194106() gopurs_runtime.Value {
	once_flip__gopurs_runtime_Value_2551194106.Do(func() {
		cache_flip__gopurs_runtime_Value_2551194106 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__gopurs_runtime_Value_2551194106(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__gopurs_runtime_Value_2551194106
}

var cache_flip__gopurs_runtime_Value_2108834362 gopurs_runtime.Value
var once_flip__gopurs_runtime_Value_2108834362 sync.Once
func Get_flip__gopurs_runtime_Value_2108834362() gopurs_runtime.Value {
	once_flip__gopurs_runtime_Value_2108834362.Do(func() {
		cache_flip__gopurs_runtime_Value_2108834362 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__gopurs_runtime_Value_2108834362(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__gopurs_runtime_Value_2108834362
}

var cache_flip__gopurs_runtime_Value_3361344762 gopurs_runtime.Value
var once_flip__gopurs_runtime_Value_3361344762 sync.Once
func Get_flip__gopurs_runtime_Value_3361344762() gopurs_runtime.Value {
	once_flip__gopurs_runtime_Value_3361344762.Do(func() {
		cache_flip__gopurs_runtime_Value_3361344762 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__gopurs_runtime_Value_3361344762(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__gopurs_runtime_Value_3361344762
}

var cache_flip__gopurs_runtime_Value_1114146170 gopurs_runtime.Value
var once_flip__gopurs_runtime_Value_1114146170 sync.Once
func Get_flip__gopurs_runtime_Value_1114146170() gopurs_runtime.Value {
	once_flip__gopurs_runtime_Value_1114146170.Do(func() {
		cache_flip__gopurs_runtime_Value_1114146170 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__gopurs_runtime_Value_1114146170(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__gopurs_runtime_Value_1114146170
}

var cache_flip__gopurs_runtime_Value_2645119546 gopurs_runtime.Value
var once_flip__gopurs_runtime_Value_2645119546 sync.Once
func Get_flip__gopurs_runtime_Value_2645119546() gopurs_runtime.Value {
	once_flip__gopurs_runtime_Value_2645119546.Do(func() {
		cache_flip__gopurs_runtime_Value_2645119546 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__gopurs_runtime_Value_2645119546(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__gopurs_runtime_Value_2645119546
}

var cache_const_ gopurs_runtime.Value
var once_const_ sync.Once
func Get_const_() gopurs_runtime.Value {
	once_const_.Do(func() {
		cache_const_ = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const_(a_0_box, v_1_box)
})
	})
	return cache_const_
}

var cache_const__gopurs_runtime_Value_1014520284 gopurs_runtime.Value
var once_const__gopurs_runtime_Value_1014520284 sync.Once
func Get_const__gopurs_runtime_Value_1014520284() gopurs_runtime.Value {
	once_const__gopurs_runtime_Value_1014520284.Do(func() {
		cache_const__gopurs_runtime_Value_1014520284 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_const__gopurs_runtime_Value_1014520284(a_0_box.IntVal, v_1_box))
})
	})
	return cache_const__gopurs_runtime_Value_1014520284
}

var cache_const__gopurs_runtime_Value_3313219100 gopurs_runtime.Value
var once_const__gopurs_runtime_Value_3313219100 sync.Once
func Get_const__gopurs_runtime_Value_3313219100() gopurs_runtime.Value {
	once_const__gopurs_runtime_Value_3313219100.Do(func() {
		cache_const__gopurs_runtime_Value_3313219100 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_const__gopurs_runtime_Value_3313219100((a_0_box.IntVal) != (0), v_1_box))
})
	})
	return cache_const__gopurs_runtime_Value_3313219100
}

var cache_const__gopurs_runtime_Value_2448935612 gopurs_runtime.Value
var once_const__gopurs_runtime_Value_2448935612 sync.Once
func Get_const__gopurs_runtime_Value_2448935612() gopurs_runtime.Value {
	once_const__gopurs_runtime_Value_2448935612.Do(func() {
		cache_const__gopurs_runtime_Value_2448935612 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__gopurs_runtime_Value_2448935612(a_0_box, v_1_box)
})
	})
	return cache_const__gopurs_runtime_Value_2448935612
}

var cache_const__gopurs_runtime_Value_827936540 gopurs_runtime.Value
var once_const__gopurs_runtime_Value_827936540 sync.Once
func Get_const__gopurs_runtime_Value_827936540() gopurs_runtime.Value {
	once_const__gopurs_runtime_Value_827936540.Do(func() {
		cache_const__gopurs_runtime_Value_827936540 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__gopurs_runtime_Value_827936540(a_0_box, v_1_box)
})
	})
	return cache_const__gopurs_runtime_Value_827936540
}

var cache_const__gopurs_runtime_Value_2737562588 gopurs_runtime.Value
var once_const__gopurs_runtime_Value_2737562588 sync.Once
func Get_const__gopurs_runtime_Value_2737562588() gopurs_runtime.Value {
	once_const__gopurs_runtime_Value_2737562588.Do(func() {
		cache_const__gopurs_runtime_Value_2737562588 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__gopurs_runtime_Value_2737562588(a_0_box, v_1_box)
})
	})
	return cache_const__gopurs_runtime_Value_2737562588
}

var cache_const__gopurs_runtime_Value_2901144028 gopurs_runtime.Value
var once_const__gopurs_runtime_Value_2901144028 sync.Once
func Get_const__gopurs_runtime_Value_2901144028() gopurs_runtime.Value {
	once_const__gopurs_runtime_Value_2901144028.Do(func() {
		cache_const__gopurs_runtime_Value_2901144028 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__gopurs_runtime_Value_2901144028(a_0_box, v_1_box)
})
	})
	return cache_const__gopurs_runtime_Value_2901144028
}

var cache_const__gopurs_runtime_Value_1533104436 gopurs_runtime.Value
var once_const__gopurs_runtime_Value_1533104436 sync.Once
func Get_const__gopurs_runtime_Value_1533104436() gopurs_runtime.Value {
	once_const__gopurs_runtime_Value_1533104436.Do(func() {
		cache_const__gopurs_runtime_Value_1533104436 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__gopurs_runtime_Value_1533104436(a_0_box, v_1_box)
})
	})
	return cache_const__gopurs_runtime_Value_1533104436
}

var cache_const__gopurs_runtime_Value_4266931454 gopurs_runtime.Value
var once_const__gopurs_runtime_Value_4266931454 sync.Once
func Get_const__gopurs_runtime_Value_4266931454() gopurs_runtime.Value {
	once_const__gopurs_runtime_Value_4266931454.Do(func() {
		cache_const__gopurs_runtime_Value_4266931454 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__gopurs_runtime_Value_4266931454(a_0_box, v_1_box)
})
	})
	return cache_const__gopurs_runtime_Value_4266931454
}

var cache_const__gopurs_runtime_Value_87201948 gopurs_runtime.Value
var once_const__gopurs_runtime_Value_87201948 sync.Once
func Get_const__gopurs_runtime_Value_87201948() gopurs_runtime.Value {
	once_const__gopurs_runtime_Value_87201948.Do(func() {
		cache_const__gopurs_runtime_Value_87201948 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__gopurs_runtime_Value_87201948(a_0_box, v_1_box)
})
	})
	return cache_const__gopurs_runtime_Value_87201948
}

var cache_const__gopurs_runtime_Value_3060642877 gopurs_runtime.Value
var once_const__gopurs_runtime_Value_3060642877 sync.Once
func Get_const__gopurs_runtime_Value_3060642877() gopurs_runtime.Value {
	once_const__gopurs_runtime_Value_3060642877.Do(func() {
		cache_const__gopurs_runtime_Value_3060642877 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__gopurs_runtime_Value_3060642877(a_0_box, v_1_box)
})
	})
	return cache_const__gopurs_runtime_Value_3060642877
}

var cache_const__gopurs_runtime_Value_31775132 gopurs_runtime.Value
var once_const__gopurs_runtime_Value_31775132 sync.Once
func Get_const__gopurs_runtime_Value_31775132() gopurs_runtime.Value {
	once_const__gopurs_runtime_Value_31775132.Do(func() {
		cache_const__gopurs_runtime_Value_31775132 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__gopurs_runtime_Value_31775132(a_0_box, v_1_box)
})
	})
	return cache_const__gopurs_runtime_Value_31775132
}

var cache_const__gopurs_runtime_Value_4379098 gopurs_runtime.Value
var once_const__gopurs_runtime_Value_4379098 sync.Once
func Get_const__gopurs_runtime_Value_4379098() gopurs_runtime.Value {
	once_const__gopurs_runtime_Value_4379098.Do(func() {
		cache_const__gopurs_runtime_Value_4379098 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__gopurs_runtime_Value_4379098(a_0_box, v_1_box)
})
	})
	return cache_const__gopurs_runtime_Value_4379098
}

var cache_const__gopurs_runtime_Value_945658268 gopurs_runtime.Value
var once_const__gopurs_runtime_Value_945658268 sync.Once
func Get_const__gopurs_runtime_Value_945658268() gopurs_runtime.Value {
	once_const__gopurs_runtime_Value_945658268.Do(func() {
		cache_const__gopurs_runtime_Value_945658268 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__gopurs_runtime_Value_945658268(a_0_box, v_1_box)
})
	})
	return cache_const__gopurs_runtime_Value_945658268
}

var cache_const__gopurs_runtime_Value_4007400604 gopurs_runtime.Value
var once_const__gopurs_runtime_Value_4007400604 sync.Once
func Get_const__gopurs_runtime_Value_4007400604() gopurs_runtime.Value {
	once_const__gopurs_runtime_Value_4007400604.Do(func() {
		cache_const__gopurs_runtime_Value_4007400604 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__gopurs_runtime_Value_4007400604(a_0_box, v_1_box)
})
	})
	return cache_const__gopurs_runtime_Value_4007400604
}

var cache_const__gopurs_runtime_Value_144291955 gopurs_runtime.Value
var once_const__gopurs_runtime_Value_144291955 sync.Once
func Get_const__gopurs_runtime_Value_144291955() gopurs_runtime.Value {
	once_const__gopurs_runtime_Value_144291955.Do(func() {
		cache_const__gopurs_runtime_Value_144291955 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__gopurs_runtime_Value_144291955(a_0_box, v_1_box)
})
	})
	return cache_const__gopurs_runtime_Value_144291955
}

var cache_const__gopurs_runtime_Value_4005387916 gopurs_runtime.Value
var once_const__gopurs_runtime_Value_4005387916 sync.Once
func Get_const__gopurs_runtime_Value_4005387916() gopurs_runtime.Value {
	once_const__gopurs_runtime_Value_4005387916.Do(func() {
		cache_const__gopurs_runtime_Value_4005387916 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__gopurs_runtime_Value_4005387916(a_0_box, v_1_box)
})
	})
	return cache_const__gopurs_runtime_Value_4005387916
}

var cache_const__gopurs_runtime_Value_3293692348 gopurs_runtime.Value
var once_const__gopurs_runtime_Value_3293692348 sync.Once
func Get_const__gopurs_runtime_Value_3293692348() gopurs_runtime.Value {
	once_const__gopurs_runtime_Value_3293692348.Do(func() {
		cache_const__gopurs_runtime_Value_3293692348 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__gopurs_runtime_Value_3293692348(a_0_box, v_1_box)
})
	})
	return cache_const__gopurs_runtime_Value_3293692348
}

var cache_const__gopurs_runtime_Value_2326386184 gopurs_runtime.Value
var once_const__gopurs_runtime_Value_2326386184 sync.Once
func Get_const__gopurs_runtime_Value_2326386184() gopurs_runtime.Value {
	once_const__gopurs_runtime_Value_2326386184.Do(func() {
		cache_const__gopurs_runtime_Value_2326386184 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__gopurs_runtime_Value_2326386184(a_0_box, v_1_box)
})
	})
	return cache_const__gopurs_runtime_Value_2326386184
}

var cache_applyN gopurs_runtime.Value
var once_applyN sync.Once
func Get_applyN() gopurs_runtime.Value {
	once_applyN.Do(func() {
		cache_applyN = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applyN(f_0_box)
})
	})
	return cache_applyN
}

var cache_applyFlipped gopurs_runtime.Value
var once_applyFlipped sync.Once
func Get_applyFlipped() gopurs_runtime.Value {
	once_applyFlipped.Do(func() {
		cache_applyFlipped = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applyFlipped(x_0_box, f_1_box)
})
	})
	return cache_applyFlipped
}

var cache_apply gopurs_runtime.Value
var once_apply sync.Once
func Get_apply() gopurs_runtime.Value {
	once_apply.Do(func() {
		cache_apply = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply(f_0_box, x_1_box)
})
	})
	return cache_apply
}

var cache_apply__gopurs_runtime_Value_559335802 gopurs_runtime.Value
var once_apply__gopurs_runtime_Value_559335802 sync.Once
func Get_apply__gopurs_runtime_Value_559335802() gopurs_runtime.Value {
	once_apply__gopurs_runtime_Value_559335802.Do(func() {
		cache_apply__gopurs_runtime_Value_559335802 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__gopurs_runtime_Value_559335802(f_0_box, x_1_box)
})
	})
	return cache_apply__gopurs_runtime_Value_559335802
}

func Call_on(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var g_1 gopurs_runtime.Value = g_1_loop
_ = g_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return gopurs_runtime.Apply2(f_0, gopurs_runtime.Apply(g_1, x_2), gopurs_runtime.Apply(g_1, y_3))
}

func Call_on__gopurs_runtime_Value_1062875992(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) bool {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var g_1 gopurs_runtime.Value = g_1_loop
_ = g_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return (gopurs_runtime.Apply2(f_0, gopurs_runtime.Apply(g_1, x_2), gopurs_runtime.Apply(g_1, y_3)).IntVal) != (0)
}

func Call_on__gopurs_runtime_Value_2003439736(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var g_1 gopurs_runtime.Value = g_1_loop
_ = g_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return gopurs_runtime.Apply2(f_0, gopurs_runtime.Apply(g_1, x_2), gopurs_runtime.Apply(g_1, y_3))
}

func Call_flip(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__gopurs_runtime_Value_2010786042(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 []gopurs_runtime.Value = a_2_loop
_ = a_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(f_0, gopurs_runtime.Array(a_2), b_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_flip__gopurs_runtime_Value_662810362(f_0_loop gopurs_runtime.Value, b_1_loop []gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) []gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 []gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(f_0, a_2, gopurs_runtime.Array(b_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_flip__gopurs_runtime_Value_3946185210(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__gopurs_runtime_Value_2301875002(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__gopurs_runtime_Value_1042884218(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__gopurs_runtime_Value_1841051514(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__gopurs_runtime_Value_1752871994(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__gopurs_runtime_Value_2551194106(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__gopurs_runtime_Value_2108834362(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__gopurs_runtime_Value_3361344762(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__gopurs_runtime_Value_1114146170(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__gopurs_runtime_Value_2645119546(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_const_(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__gopurs_runtime_Value_1014520284(a_0_loop int64, v_1_loop gopurs_runtime.Value) int64 {
var a_0 int64 = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__gopurs_runtime_Value_3313219100(a_0_loop bool, v_1_loop gopurs_runtime.Value) bool {
var a_0 bool = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__gopurs_runtime_Value_2448935612(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__gopurs_runtime_Value_827936540(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__gopurs_runtime_Value_2737562588(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__gopurs_runtime_Value_2901144028(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__gopurs_runtime_Value_1533104436(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__gopurs_runtime_Value_4266931454(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__gopurs_runtime_Value_87201948(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__gopurs_runtime_Value_3060642877(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__gopurs_runtime_Value_31775132(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__gopurs_runtime_Value_4379098(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__gopurs_runtime_Value_945658268(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__gopurs_runtime_Value_4007400604(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__gopurs_runtime_Value_144291955(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__gopurs_runtime_Value_4005387916(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__gopurs_runtime_Value_3293692348(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__gopurs_runtime_Value_2326386184(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_applyN(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(n_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var n_2_loop gopurs_runtime.Value = n_2_loop_val
var acc_3_loop gopurs_runtime.Value = acc_3_loop_val
go__1_0:
for {
if false { continue go__1_0 }
var n_2 gopurs_runtime.Value = n_2_loop
_ = n_2
var acc_3 gopurs_runtime.Value = acc_3_loop
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
n_2_loop = gopurs_runtime.Int((n_2.IntVal) - (1))
acc_3_loop = gopurs_runtime.Apply(f_0, acc_3)
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

func Call_applyFlipped(x_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply(f_1, x_0)
}

func Call_apply(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(f_0, x_1)
}

func Call_apply__gopurs_runtime_Value_559335802(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(f_0, x_1)
}


