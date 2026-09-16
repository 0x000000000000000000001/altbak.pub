package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Function_on gopurs_runtime.Value
var once_Data_Function_on sync.Once
func Get_Data_Function_on() gopurs_runtime.Value {
	once_Data_Function_on.Do(func() {
		cache_Data_Function_on = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_on(f_0_box, g_1_box, x_2_box, y_3_box)
})
	})
	return cache_Data_Function_on
}

var cache_Data_Function_flip gopurs_runtime.Value
var once_Data_Function_flip sync.Once
func Get_Data_Function_flip() gopurs_runtime.Value {
	once_Data_Function_flip.Do(func() {
		cache_Data_Function_flip = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_Data_Function_flip
}

var cache_Data_Function_flip__1501157206 gopurs_runtime.Value
var once_Data_Function_flip__1501157206 sync.Once
func Get_Data_Function_flip__1501157206() gopurs_runtime.Value {
	once_Data_Function_flip__1501157206.Do(func() {
		cache_Data_Function_flip__1501157206 = gopurs_runtime.Func3(func(f_unused_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(Call_Data_Function_flip__1501157206(f_unused_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](b_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](a_2_box)))}
})
	})
	return cache_Data_Function_flip__1501157206
}

var cache_Data_Function_flip__379318895 gopurs_runtime.Value
var once_Data_Function_flip__379318895 sync.Once
func Get_Data_Function_flip__379318895() gopurs_runtime.Value {
	once_Data_Function_flip__379318895.Do(func() {
		cache_Data_Function_flip__379318895 = gopurs_runtime.Func3(func(f_unused_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_Function_flip__379318895(f_unused_0_box, Rebox_Data_Function_849153993_2442833393(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_1_box)), a_2_box))}
})
	})
	return cache_Data_Function_flip__379318895
}

var cache_Data_Function_flip__1483642665 gopurs_runtime.Value
var once_Data_Function_flip__1483642665 sync.Once
func Get_Data_Function_flip__1483642665() gopurs_runtime.Value {
	once_Data_Function_flip__1483642665.Do(func() {
		cache_Data_Function_flip__1483642665 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Function_flip__1483642665(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](b_1_box), a_2_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Function_flip__1483642665
}

var cache_Data_Function_go__const gopurs_runtime.Value
var once_Data_Function_go__const sync.Once
func Get_Data_Function_go__const() gopurs_runtime.Value {
	once_Data_Function_go__const.Do(func() {
		cache_Data_Function_go__const = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_go__const(a_0_box, v_1_box)
})
	})
	return cache_Data_Function_go__const
}

var cache_Data_Function_const__3126182883 gopurs_runtime.Value
var once_Data_Function_const__3126182883 sync.Once
func Get_Data_Function_const__3126182883() gopurs_runtime.Value {
	once_Data_Function_const__3126182883.Do(func() {
		cache_Data_Function_const__3126182883 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Function_const__3126182883(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](a_0_box), v_1_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Function_const__3126182883
}

var cache_Data_Function_const__1198217881 gopurs_runtime.Value
var once_Data_Function_const__1198217881 sync.Once
func Get_Data_Function_const__1198217881() gopurs_runtime.Value {
	once_Data_Function_const__1198217881.Do(func() {
		cache_Data_Function_const__1198217881 = gopurs_runtime.Func2(func(a_unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_const__1198217881(a_unused_0_box, v_1_box)
})
	})
	return cache_Data_Function_const__1198217881
}

var cache_Data_Function_const__1677613395 gopurs_runtime.Value
var once_Data_Function_const__1677613395 sync.Once
func Get_Data_Function_const__1677613395() gopurs_runtime.Value {
	once_Data_Function_const__1677613395.Do(func() {
		cache_Data_Function_const__1677613395 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_const__1677613395(a_0_box, v_1_box)
})
	})
	return cache_Data_Function_const__1677613395
}

var cache_Data_Function_const__320519811 gopurs_runtime.Value
var once_Data_Function_const__320519811 sync.Once
func Get_Data_Function_const__320519811() gopurs_runtime.Value {
	once_Data_Function_const__320519811.Do(func() {
		cache_Data_Function_const__320519811 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_const__320519811(a_0_box, v_1_box)
})
	})
	return cache_Data_Function_const__320519811
}

var cache_Data_Function_const__3823496979 gopurs_runtime.Value
var once_Data_Function_const__3823496979 sync.Once
func Get_Data_Function_const__3823496979() gopurs_runtime.Value {
	once_Data_Function_const__3823496979.Do(func() {
		cache_Data_Function_const__3823496979 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_const__3823496979(a_0_box, v_1_box)
})
	})
	return cache_Data_Function_const__3823496979
}

var cache_Data_Function_const__1320027331 gopurs_runtime.Value
var once_Data_Function_const__1320027331 sync.Once
func Get_Data_Function_const__1320027331() gopurs_runtime.Value {
	once_Data_Function_const__1320027331.Do(func() {
		cache_Data_Function_const__1320027331 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_const__1320027331(a_0_box, v_1_box)
})
	})
	return cache_Data_Function_const__1320027331
}

var cache_Data_Function_const__3225710767 gopurs_runtime.Value
var once_Data_Function_const__3225710767 sync.Once
func Get_Data_Function_const__3225710767() gopurs_runtime.Value {
	once_Data_Function_const__3225710767.Do(func() {
		cache_Data_Function_const__3225710767 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Function_const__3225710767((a_0_box.IntVal) != (0), uint32(v_1_box.IntVal)))
})
	})
	return cache_Data_Function_const__3225710767
}

var cache_Data_Function_const__318080054 gopurs_runtime.Value
var once_Data_Function_const__318080054 sync.Once
func Get_Data_Function_const__318080054() gopurs_runtime.Value {
	once_Data_Function_const__318080054.Do(func() {
		cache_Data_Function_const__318080054 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Function_const__318080054((a_0_box.IntVal) != (0), v_1_box.IntVal))
})
	})
	return cache_Data_Function_const__318080054
}

var cache_Data_Function_const__2630302150 gopurs_runtime.Value
var once_Data_Function_const__2630302150 sync.Once
func Get_Data_Function_const__2630302150() gopurs_runtime.Value {
	once_Data_Function_const__2630302150.Do(func() {
		cache_Data_Function_const__2630302150 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Function_const__2630302150((a_0_box.IntVal) != (0), v_1_box.FloatVal()))
})
	})
	return cache_Data_Function_const__2630302150
}

var cache_Data_Function_const__2823118163 gopurs_runtime.Value
var once_Data_Function_const__2823118163 sync.Once
func Get_Data_Function_const__2823118163() gopurs_runtime.Value {
	once_Data_Function_const__2823118163.Do(func() {
		cache_Data_Function_const__2823118163 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_const__2823118163(a_0_box, v_1_box)
})
	})
	return cache_Data_Function_const__2823118163
}

var cache_Data_Function_const__3916468150 gopurs_runtime.Value
var once_Data_Function_const__3916468150 sync.Once
func Get_Data_Function_const__3916468150() gopurs_runtime.Value {
	once_Data_Function_const__3916468150.Do(func() {
		cache_Data_Function_const__3916468150 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Function_const__3916468150(a_0_box.IntVal, v_1_box.IntVal))
})
	})
	return cache_Data_Function_const__3916468150
}

var cache_Data_Function_const__3422218083 gopurs_runtime.Value
var once_Data_Function_const__3422218083 sync.Once
func Get_Data_Function_const__3422218083() gopurs_runtime.Value {
	once_Data_Function_const__3422218083.Do(func() {
		cache_Data_Function_const__3422218083 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Function_const__3422218083(a_0_box.IntVal, v_1_box))
})
	})
	return cache_Data_Function_const__3422218083
}

var cache_Data_Function_applyN gopurs_runtime.Value
var once_Data_Function_applyN sync.Once
func Get_Data_Function_applyN() gopurs_runtime.Value {
	once_Data_Function_applyN.Do(func() {
		cache_Data_Function_applyN = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_applyN(f_0_box)
})
	})
	return cache_Data_Function_applyN
}

var cache_Data_Function_applyFlipped gopurs_runtime.Value
var once_Data_Function_applyFlipped sync.Once
func Get_Data_Function_applyFlipped() gopurs_runtime.Value {
	once_Data_Function_applyFlipped.Do(func() {
		cache_Data_Function_applyFlipped = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_applyFlipped(x_0_box, f_1_box)
})
	})
	return cache_Data_Function_applyFlipped
}

var cache_Data_Function_apply gopurs_runtime.Value
var once_Data_Function_apply sync.Once
func Get_Data_Function_apply() gopurs_runtime.Value {
	once_Data_Function_apply.Do(func() {
		cache_Data_Function_apply = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_apply(f_0_box, x_1_box)
})
	})
	return cache_Data_Function_apply
}

var cache_Data_Function_apply__1173000582 gopurs_runtime.Value
var once_Data_Function_apply__1173000582 sync.Once
func Get_Data_Function_apply__1173000582() gopurs_runtime.Value {
	once_Data_Function_apply__1173000582.Do(func() {
		cache_Data_Function_apply__1173000582 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(Call_Data_Function_apply__1173000582(f_unused_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](x_1_box)))}
})
	})
	return cache_Data_Function_apply__1173000582
}

var cache_Data_Function_apply__26935913 gopurs_runtime.Value
var once_Data_Function_apply__26935913 sync.Once
func Get_Data_Function_apply__26935913() gopurs_runtime.Value {
	once_Data_Function_apply__26935913.Do(func() {
		cache_Data_Function_apply__26935913 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(Call_Data_Function_apply__26935913(f_0_box, Rebox_Data_Function_3094389156_3839235747(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](x_1_box))))}
})
	})
	return cache_Data_Function_apply__26935913
}

var cache_Data_Function_apply__2647861350 gopurs_runtime.Value
var once_Data_Function_apply__2647861350 sync.Once
func Get_Data_Function_apply__2647861350() gopurs_runtime.Value {
	once_Data_Function_apply__2647861350.Do(func() {
		cache_Data_Function_apply__2647861350 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Function_apply__2647861350(f_unused_0_box, Rebox_Data_Function_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](x_1_box))))
})
	})
	return cache_Data_Function_apply__2647861350
}

var cache_Data_Function_apply__2639368963 gopurs_runtime.Value
var once_Data_Function_apply__2639368963 sync.Once
func Get_Data_Function_apply__2639368963() gopurs_runtime.Value {
	once_Data_Function_apply__2639368963.Do(func() {
		cache_Data_Function_apply__2639368963 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Function_apply__2639368963(f_unused_0_box, Rebox_Data_Function_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](x_1_box))))
})
	})
	return cache_Data_Function_apply__2639368963
}

var cache_Data_Function_apply__1422622261 gopurs_runtime.Value
var once_Data_Function_apply__1422622261 sync.Once
func Get_Data_Function_apply__1422622261() gopurs_runtime.Value {
	once_Data_Function_apply__1422622261.Do(func() {
		cache_Data_Function_apply__1422622261 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_apply__1422622261(f_unused_0_box, x_1_box)
})
	})
	return cache_Data_Function_apply__1422622261
}

var cache_Data_Function_apply__1560082597 gopurs_runtime.Value
var once_Data_Function_apply__1560082597 sync.Once
func Get_Data_Function_apply__1560082597() gopurs_runtime.Value {
	once_Data_Function_apply__1560082597.Do(func() {
		cache_Data_Function_apply__1560082597 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_apply__1560082597(f_unused_0_box, x_1_box)
})
	})
	return cache_Data_Function_apply__1560082597
}

var cache_Data_Function_apply__483590224 gopurs_runtime.Value
var once_Data_Function_apply__483590224 sync.Once
func Get_Data_Function_apply__483590224() gopurs_runtime.Value {
	once_Data_Function_apply__483590224.Do(func() {
		cache_Data_Function_apply__483590224 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_apply__483590224(f_unused_0_box, x_1_box)
})
	})
	return cache_Data_Function_apply__483590224
}

var cache_Data_Function_apply__1378432126 gopurs_runtime.Value
var once_Data_Function_apply__1378432126 sync.Once
func Get_Data_Function_apply__1378432126() gopurs_runtime.Value {
	once_Data_Function_apply__1378432126.Do(func() {
		cache_Data_Function_apply__1378432126 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_Data_Function_apply__1378432126(f_unused_0_box, func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}())
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_Data_Function_apply__1378432126
}

var cache_Data_Function_apply__2114047305 gopurs_runtime.Value
var once_Data_Function_apply__2114047305 sync.Once
func Get_Data_Function_apply__2114047305() gopurs_runtime.Value {
	once_Data_Function_apply__2114047305.Do(func() {
		cache_Data_Function_apply__2114047305 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Function_apply__2114047305(f_0_box, (x_1_box.IntVal) != (0)))
})
	})
	return cache_Data_Function_apply__2114047305
}

var cache_Data_Function_apply__3402160552 gopurs_runtime.Value
var once_Data_Function_apply__3402160552 sync.Once
func Get_Data_Function_apply__3402160552() gopurs_runtime.Value {
	once_Data_Function_apply__3402160552.Do(func() {
		cache_Data_Function_apply__3402160552 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_apply__3402160552(f_unused_0_box, x_1_box)
})
	})
	return cache_Data_Function_apply__3402160552
}

var cache_Data_Function_apply__355458268 gopurs_runtime.Value
var once_Data_Function_apply__355458268 sync.Once
func Get_Data_Function_apply__355458268() gopurs_runtime.Value {
	once_Data_Function_apply__355458268.Do(func() {
		cache_Data_Function_apply__355458268 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(Call_Data_Function_apply__355458268(f_unused_0_box, x_1_box.FloatVal()))
})
	})
	return cache_Data_Function_apply__355458268
}

var cache_Data_Function_apply__2245704861 gopurs_runtime.Value
var once_Data_Function_apply__2245704861 sync.Once
func Get_Data_Function_apply__2245704861() gopurs_runtime.Value {
	once_Data_Function_apply__2245704861.Do(func() {
		cache_Data_Function_apply__2245704861 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(Call_Data_Function_apply__2245704861(f_unused_0_box, x_1_box.FloatVal()))
})
	})
	return cache_Data_Function_apply__2245704861
}

var cache_Data_Function_apply__2515510025 gopurs_runtime.Value
var once_Data_Function_apply__2515510025 sync.Once
func Get_Data_Function_apply__2515510025() gopurs_runtime.Value {
	once_Data_Function_apply__2515510025.Do(func() {
		cache_Data_Function_apply__2515510025 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Function_apply__2515510025(f_0_box, x_1_box.StrVal())
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
})
	})
	return cache_Data_Function_apply__2515510025
}

var cache_Data_Function_apply__2067995554 gopurs_runtime.Value
var once_Data_Function_apply__2067995554 sync.Once
func Get_Data_Function_apply__2067995554() gopurs_runtime.Value {
	once_Data_Function_apply__2067995554.Do(func() {
		cache_Data_Function_apply__2067995554 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_apply__2067995554(f_unused_0_box, x_1_box.StrVal())
})
	})
	return cache_Data_Function_apply__2067995554
}

func Call_Data_Function_on(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_Data_Function_flip(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_Data_Function_flip__1501157206(f_unused_0_loop gopurs_runtime.Value, b_1_loop *Constructor_Data_Time_Time, a_2_loop *Constructor_Data_Date_Date) *Constructor_Data_DateTime_DateTime {
flip__1501157206:
for {
if false { continue flip__1501157206 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var b_1 *Constructor_Data_Time_Time = b_1_loop
_ = b_1
var a_2 *Constructor_Data_Date_Date = a_2_loop
_ = a_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_DateTime_DateTime](gopurs_runtime.Apply2(Get_Data_DateTime_DateTime(), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(a_2)}, gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(b_1)}))
}
}

func Call_Data_Function_flip__379318895(f_unused_0_loop gopurs_runtime.Value, b_1_loop *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]], a_2_loop gopurs_runtime.Value) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
flip__379318895:
for {
if false { continue flip__379318895 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var b_1 *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(Get_Data_Interval_Duration_Iso_foldMap(), a_2, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_Function_2442833393_849153993(b_1))}))
}
}

func Call_Data_Function_flip__1483642665(f_0_loop gopurs_runtime.Value, b_1_loop *Constructor_Data_Date_Date, a_2_loop int64) struct{V0 gopurs_runtime.Value; V1 bool} {
flip__1483642665:
for {
if false { continue flip__1483642665 }
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 *Constructor_Data_Date_Date = b_1_loop
_ = b_1
var a_2 int64 = a_2_loop
_ = a_2
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Apply2(f_0, gopurs_runtime.Int(a_2), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(b_1)})
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Function_go__const(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__3126182883(a_0_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value], v_1_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
const__3126182883:
for {
if false { continue const__3126182883 }
var a_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(a_0)}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Function_const__1198217881(a_unused_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
const__1198217881:
for {
if false { continue const__1198217881 }
var a_unused_0 gopurs_runtime.Value = a_unused_0_loop
_ = a_unused_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply(Get_Control_Applicative_pure__3122131407(), Get_Data_Unit_unit())
}
}

func Call_Data_Function_const__1677613395(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
const__1677613395:
for {
if false { continue const__1677613395 }
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}
}

func Call_Data_Function_const__320519811(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
const__320519811:
for {
if false { continue const__320519811 }
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}
}

func Call_Data_Function_const__3823496979(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
const__3823496979:
for {
if false { continue const__3823496979 }
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}
}

func Call_Data_Function_const__1320027331(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
const__1320027331:
for {
if false { continue const__1320027331 }
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}
}

func Call_Data_Function_const__3225710767(a_0_loop bool, v_1_loop uint32) bool {
const__3225710767:
for {
if false { continue const__3225710767 }
var a_0 bool = a_0_loop
_ = a_0
var v_1 uint32 = v_1_loop
_ = v_1
return a_0
}
}

func Call_Data_Function_const__318080054(a_0_loop bool, v_1_loop int64) bool {
const__318080054:
for {
if false { continue const__318080054 }
var a_0 bool = a_0_loop
_ = a_0
var v_1 int64 = v_1_loop
_ = v_1
return a_0
}
}

func Call_Data_Function_const__2630302150(a_0_loop bool, v_1_loop float64) bool {
const__2630302150:
for {
if false { continue const__2630302150 }
var a_0 bool = a_0_loop
_ = a_0
var v_1 float64 = v_1_loop
_ = v_1
return a_0
}
}

func Call_Data_Function_const__2823118163(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
const__2823118163:
for {
if false { continue const__2823118163 }
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}
}

func Call_Data_Function_const__3916468150(a_0_loop int64, v_1_loop int64) int64 {
const__3916468150:
for {
if false { continue const__3916468150 }
var a_0 int64 = a_0_loop
_ = a_0
var v_1 int64 = v_1_loop
_ = v_1
return a_0
}
}

func Call_Data_Function_const__3422218083(a_0_loop int64, v_1_loop gopurs_runtime.Value) int64 {
const__3422218083:
for {
if false { continue const__3422218083 }
var a_0 int64 = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}
}

func Call_Data_Function_applyN(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var Call_local_Data_Function_go__go_1_0_0 func(int64, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_Function_go__go_1_0_0
var go__go_1_0_0 gopurs_runtime.Value
_ = go__go_1_0_0
Call_local_Data_Function_go__go_1_0_0 = func(n_2_loop int64, acc_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_1_0_0:
for {
if false { continue go__go_1_0_0 }
var n_2 int64 = n_2_loop
_ = n_2
var acc_3 gopurs_runtime.Value = acc_3_loop
_ = acc_3
var __t1 gopurs_runtime.Value
{
if (n_2) <= (int64(0)) {
__t1 = acc_3
goto end_branch_1
} else {

}
}
{
n_2_loop = (n_2) - (int64(1))
acc_3_loop = gopurs_runtime.Apply(f_0, acc_3)
continue go__go_1_0_0
__t1 = func() gopurs_runtime.Value { panic("unreachable") }()
}
end_branch_1:
return __t1
}
}
go__go_1_0_0 = gopurs_runtime.Func(func(n_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_Function_go__go_1_0_0(n_2_loop_val.IntVal, acc_3_loop_val)
})
})
return go__go_1_0_0
}

func Call_Data_Function_applyFlipped(x_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply(f_1, x_0)
}

func Call_Data_Function_apply(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(f_0, x_1)
}

func Call_Data_Function_apply__1173000582(f_unused_0_loop gopurs_runtime.Value, x_1_loop *Constructor_Data_Date_Date) *Constructor_Data_Date_Date {
apply__1173000582:
for {
if false { continue apply__1173000582 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var x_1 *Constructor_Data_Date_Date = x_1_loop
_ = x_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](gopurs_runtime.Apply(Get_Partial_Unsafe_unsafePartial__3011796551(), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(x_1)}))
}
}

func Call_Data_Function_apply__26935913(f_0_loop gopurs_runtime.Value, x_1_loop *Constructor_Data_Maybe_Just[*Constructor_Data_Time_Time]) *Constructor_Data_Time_Time {
apply__26935913:
for {
if false { continue apply__26935913 }
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 *Constructor_Data_Maybe_Just[*Constructor_Data_Time_Time] = x_1_loop
_ = x_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Function_3839235747_3094389156(x_1))}))
}
}

func Call_Data_Function_apply__2647861350(f_unused_0_loop gopurs_runtime.Value, x_1_loop *Constructor_Data_Maybe_Just[int64]) bool {
apply__2647861350:
for {
if false { continue apply__2647861350 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var x_1 *Constructor_Data_Maybe_Just[int64] = x_1_loop
_ = x_1
return (gopurs_runtime.Apply(Get_Data_Maybe_isJust__1937207052(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Function_1170268447_3094389156(x_1))}).IntVal) != (0)
}
}

func Call_Data_Function_apply__2639368963(f_unused_0_loop gopurs_runtime.Value, x_1_loop *Constructor_Data_Maybe_Just[int64]) bool {
apply__2639368963:
for {
if false { continue apply__2639368963 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var x_1 *Constructor_Data_Maybe_Just[int64] = x_1_loop
_ = x_1
return (gopurs_runtime.Apply(Get_Data_Maybe_isNothing__1937207052(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Function_1170268447_3094389156(x_1))}).IntVal) != (0)
}
}

func Call_Data_Function_apply__1422622261(f_unused_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__1422622261:
for {
if false { continue apply__1422622261 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(Get_Effect_Aff_void1(), x_1)
}
}

func Call_Data_Function_apply__1560082597(f_unused_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__1560082597:
for {
if false { continue apply__1560082597 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(Get_Bench_Aff_liftEffect(), x_1)
}
}

func Call_Data_Function_apply__483590224(f_unused_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__483590224:
for {
if false { continue apply__483590224 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(Get_Effect_Aff_liftEffect(), x_1)
}
}

func Call_Data_Function_apply__1378432126(f_unused_0_loop gopurs_runtime.Value, x_1_loop []int64) []int64 {
apply__1378432126:
for {
if false { continue apply__1378432126 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var x_1 []int64 = x_1_loop
_ = x_1
return func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Data_Array_NonEmpty_unsafeFromArray__4020493786(), func() gopurs_runtime.Value {
					arr := x_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}()
}
}

func Call_Data_Function_apply__2114047305(f_0_loop gopurs_runtime.Value, x_1_loop bool) bool {
apply__2114047305:
for {
if false { continue apply__2114047305 }
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 bool = x_1_loop
_ = x_1
return (gopurs_runtime.Apply(f_0, gopurs_runtime.Bool(x_1)).IntVal) != (0)
}
}

func Call_Data_Function_apply__3402160552(f_unused_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__3402160552:
for {
if false { continue apply__3402160552 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(Get_Effect_Aff_Compat_pure(), x_1)
}
}

func Call_Data_Function_apply__355458268(f_unused_0_loop gopurs_runtime.Value, x_1_loop float64) float64 {
apply__355458268:
for {
if false { continue apply__355458268 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var x_1 float64 = x_1_loop
_ = x_1
return gopurs_runtime.Float(x_1).FloatVal()
}
}

func Call_Data_Function_apply__2245704861(f_unused_0_loop gopurs_runtime.Value, x_1_loop float64) float64 {
apply__2245704861:
for {
if false { continue apply__2245704861 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var x_1 float64 = x_1_loop
_ = x_1
return gopurs_runtime.Float(x_1).FloatVal()
}
}

func Call_Data_Function_apply__2515510025(f_0_loop gopurs_runtime.Value, x_1_loop string) struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool} {
apply__2515510025:
for {
if false { continue apply__2515510025 }
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 string = x_1_loop
_ = x_1
return func() struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool} {
				_v := gopurs_runtime.Apply(f_0, gopurs_runtime.Str(x_1))
				if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{V0: gopurs_runtime.Value{}, V1: (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr).V0, V2: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{V0: (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: gopurs_runtime.Value{}, V2: false}
			}()
}
}

func Call_Data_Function_apply__2067995554(f_unused_0_loop gopurs_runtime.Value, x_1_loop string) gopurs_runtime.Value {
apply__2067995554:
for {
if false { continue apply__2067995554 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var x_1 string = x_1_loop
_ = x_1
return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(x_1))
}
}

func Rebox_Data_Function_1170268447_3094389156(in *Constructor_Data_Maybe_Just[int64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Int(in.V0)
	return out
}

func Rebox_Data_Function_138441832_3132786365(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[uint32, float64] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[uint32, float64]{}
		out.V0 = uint32(in.V0.IntVal)
		out.V1 = in.V1.FloatVal()
	return out
}

func Rebox_Data_Function_2442833393_849153993(in *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Function_3132786365_138441832(in.V0))}
		out.V1 = Rebox_Data_Function_2442833393_849153993(in.V1)
	return out
}

func Rebox_Data_Function_3094389156_1170268447(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[int64]{}
		out.V0 = in.V0.IntVal
	return out
}

func Rebox_Data_Function_3094389156_3839235747(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Time_Time] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Time_Time]{}
		out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](in.V0)
	return out
}

func Rebox_Data_Function_3132786365_138441832(in *Constructor_Data_Tuple_Tuple[uint32, float64]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V0), UnsafePtr: nil}
		out.V1 = gopurs_runtime.Float(in.V1)
	return out
}

func Rebox_Data_Function_3839235747_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Time_Time]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}

func Rebox_Data_Function_849153993_2442833393(in *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] {
	if in == nil { return nil }
	out := &Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]{}
		out.V0 = Rebox_Data_Function_138441832_3132786365(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0))
		out.V1 = Rebox_Data_Function_849153993_2442833393(in.V1)
	return out
}


