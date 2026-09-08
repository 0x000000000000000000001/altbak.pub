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

var cache_Data_Function_flip__1757181342 gopurs_runtime.Value
var once_Data_Function_flip__1757181342 sync.Once
func Get_Data_Function_flip__1757181342() gopurs_runtime.Value {
	once_Data_Function_flip__1757181342.Do(func() {
		cache_Data_Function_flip__1757181342 = gopurs_runtime.Func3(func(f_unused_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(Call_Data_Function_flip__1757181342(f_unused_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](b_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](a_2_box)))}
})
	})
	return cache_Data_Function_flip__1757181342
}

var cache_Data_Function_flip__2357332471 gopurs_runtime.Value
var once_Data_Function_flip__2357332471 sync.Once
func Get_Data_Function_flip__2357332471() gopurs_runtime.Value {
	once_Data_Function_flip__2357332471.Do(func() {
		cache_Data_Function_flip__2357332471 = gopurs_runtime.Func3(func(f_unused_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_Function_flip__2357332471(f_unused_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]](b_1_box), a_2_box))}
})
	})
	return cache_Data_Function_flip__2357332471
}

var cache_Data_Function_flip__4138356370 gopurs_runtime.Value
var once_Data_Function_flip__4138356370 sync.Once
func Get_Data_Function_flip__4138356370() gopurs_runtime.Value {
	once_Data_Function_flip__4138356370.Do(func() {
		cache_Data_Function_flip__4138356370 = gopurs_runtime.Func3(func(f_unused_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Function_flip__4138356370(f_unused_0_box, b_1_box, a_2_box)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Function_flip__4138356370
}

var cache_Data_Function_flip__965230301 gopurs_runtime.Value
var once_Data_Function_flip__965230301 sync.Once
func Get_Data_Function_flip__965230301() gopurs_runtime.Value {
	once_Data_Function_flip__965230301.Do(func() {
		cache_Data_Function_flip__965230301 = gopurs_runtime.Func3(func(f_unused_0_box gopurs_runtime.Value, b_unused_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Function_flip__965230301(f_unused_0_box, b_unused_1_box, a_2_box)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Function_flip__965230301
}

var cache_Data_Function_flip__746081925 gopurs_runtime.Value
var once_Data_Function_flip__746081925 sync.Once
func Get_Data_Function_flip__746081925() gopurs_runtime.Value {
	once_Data_Function_flip__746081925.Do(func() {
		cache_Data_Function_flip__746081925 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Function_flip__746081925(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](b_1_box), a_2_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Function_flip__746081925
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

var cache_Data_Function_const__254090813 gopurs_runtime.Value
var once_Data_Function_const__254090813 sync.Once
func Get_Data_Function_const__254090813() gopurs_runtime.Value {
	once_Data_Function_const__254090813.Do(func() {
		cache_Data_Function_const__254090813 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Function_const__254090813(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](a_0_box), v_1_box)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Function_const__254090813
}

var cache_Data_Function_const__2942423693 gopurs_runtime.Value
var once_Data_Function_const__2942423693 sync.Once
func Get_Data_Function_const__2942423693() gopurs_runtime.Value {
	once_Data_Function_const__2942423693.Do(func() {
		cache_Data_Function_const__2942423693 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_const__2942423693(a_0_box, v_1_box)
})
	})
	return cache_Data_Function_const__2942423693
}

var cache_Data_Function_const__235433565 gopurs_runtime.Value
var once_Data_Function_const__235433565 sync.Once
func Get_Data_Function_const__235433565() gopurs_runtime.Value {
	once_Data_Function_const__235433565.Do(func() {
		cache_Data_Function_const__235433565 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_const__235433565(a_0_box, v_1_box)
})
	})
	return cache_Data_Function_const__235433565
}

var cache_Data_Function_const__549142733 gopurs_runtime.Value
var once_Data_Function_const__549142733 sync.Once
func Get_Data_Function_const__549142733() gopurs_runtime.Value {
	once_Data_Function_const__549142733.Do(func() {
		cache_Data_Function_const__549142733 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_const__549142733(a_0_box, v_1_box)
})
	})
	return cache_Data_Function_const__549142733
}

var cache_Data_Function_const__2135255069 gopurs_runtime.Value
var once_Data_Function_const__2135255069 sync.Once
func Get_Data_Function_const__2135255069() gopurs_runtime.Value {
	once_Data_Function_const__2135255069.Do(func() {
		cache_Data_Function_const__2135255069 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_const__2135255069(a_0_box, v_1_box)
})
	})
	return cache_Data_Function_const__2135255069
}

var cache_Data_Function_const__4293180849 gopurs_runtime.Value
var once_Data_Function_const__4293180849 sync.Once
func Get_Data_Function_const__4293180849() gopurs_runtime.Value {
	once_Data_Function_const__4293180849.Do(func() {
		cache_Data_Function_const__4293180849 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Function_const__4293180849((a_0_box.IntVal) != (0), uint32(v_1_box.IntVal)))
})
	})
	return cache_Data_Function_const__4293180849
}

var cache_Data_Function_const__1421211688 gopurs_runtime.Value
var once_Data_Function_const__1421211688 sync.Once
func Get_Data_Function_const__1421211688() gopurs_runtime.Value {
	once_Data_Function_const__1421211688.Do(func() {
		cache_Data_Function_const__1421211688 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Function_const__1421211688((a_0_box.IntVal) != (0), v_1_box.IntVal))
})
	})
	return cache_Data_Function_const__1421211688
}

var cache_Data_Function_const__2519857944 gopurs_runtime.Value
var once_Data_Function_const__2519857944 sync.Once
func Get_Data_Function_const__2519857944() gopurs_runtime.Value {
	once_Data_Function_const__2519857944.Do(func() {
		cache_Data_Function_const__2519857944 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Function_const__2519857944((a_0_box.IntVal) != (0), v_1_box.FloatVal()))
})
	})
	return cache_Data_Function_const__2519857944
}

var cache_Data_Function_const__3022713485 gopurs_runtime.Value
var once_Data_Function_const__3022713485 sync.Once
func Get_Data_Function_const__3022713485() gopurs_runtime.Value {
	once_Data_Function_const__3022713485.Do(func() {
		cache_Data_Function_const__3022713485 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_const__3022713485(a_0_box, v_1_box)
})
	})
	return cache_Data_Function_const__3022713485
}

var cache_Data_Function_const__3052583336 gopurs_runtime.Value
var once_Data_Function_const__3052583336 sync.Once
func Get_Data_Function_const__3052583336() gopurs_runtime.Value {
	once_Data_Function_const__3052583336.Do(func() {
		cache_Data_Function_const__3052583336 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Function_const__3052583336(a_0_box.IntVal, v_1_box.IntVal))
})
	})
	return cache_Data_Function_const__3052583336
}

var cache_Data_Function_const__1592876733 gopurs_runtime.Value
var once_Data_Function_const__1592876733 sync.Once
func Get_Data_Function_const__1592876733() gopurs_runtime.Value {
	once_Data_Function_const__1592876733.Do(func() {
		cache_Data_Function_const__1592876733 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Function_const__1592876733(a_0_box.IntVal, v_1_box))
})
	})
	return cache_Data_Function_const__1592876733
}

var cache_Data_Function_const__3126880413 gopurs_runtime.Value
var once_Data_Function_const__3126880413 sync.Once
func Get_Data_Function_const__3126880413() gopurs_runtime.Value {
	once_Data_Function_const__3126880413.Do(func() {
		cache_Data_Function_const__3126880413 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_Function_const__3126880413(a_0_box.StrVal(), v_1_box))
})
	})
	return cache_Data_Function_const__3126880413
}

var cache_Data_Function_const__2623567125 gopurs_runtime.Value
var once_Data_Function_const__2623567125 sync.Once
func Get_Data_Function_const__2623567125() gopurs_runtime.Value {
	once_Data_Function_const__2623567125.Do(func() {
		cache_Data_Function_const__2623567125 = gopurs_runtime.Func2(func(a_unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_const__2623567125(a_unused_0_box, v_1_box.StrVal())
})
	})
	return cache_Data_Function_const__2623567125
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

var cache_Data_Function_apply__3981898766 gopurs_runtime.Value
var once_Data_Function_apply__3981898766 sync.Once
func Get_Data_Function_apply__3981898766() gopurs_runtime.Value {
	once_Data_Function_apply__3981898766.Do(func() {
		cache_Data_Function_apply__3981898766 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(Call_Data_Function_apply__3981898766(f_unused_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](x_1_box)))}
})
	})
	return cache_Data_Function_apply__3981898766
}

var cache_Data_Function_apply__3720661445 gopurs_runtime.Value
var once_Data_Function_apply__3720661445 sync.Once
func Get_Data_Function_apply__3720661445() gopurs_runtime.Value {
	once_Data_Function_apply__3720661445.Do(func() {
		cache_Data_Function_apply__3720661445 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(Call_Data_Function_apply__3720661445(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Time_Time]](x_1_box)))}
})
	})
	return cache_Data_Function_apply__3720661445
}

var cache_Data_Function_apply__2361363086 gopurs_runtime.Value
var once_Data_Function_apply__2361363086 sync.Once
func Get_Data_Function_apply__2361363086() gopurs_runtime.Value {
	once_Data_Function_apply__2361363086.Do(func() {
		cache_Data_Function_apply__2361363086 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Function_apply__2361363086(f_unused_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](x_1_box)))
})
	})
	return cache_Data_Function_apply__2361363086
}

var cache_Data_Function_apply__2704629163 gopurs_runtime.Value
var once_Data_Function_apply__2704629163 sync.Once
func Get_Data_Function_apply__2704629163() gopurs_runtime.Value {
	once_Data_Function_apply__2704629163.Do(func() {
		cache_Data_Function_apply__2704629163 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Function_apply__2704629163(f_unused_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](x_1_box)))
})
	})
	return cache_Data_Function_apply__2704629163
}

var cache_Data_Function_apply__1359309341 gopurs_runtime.Value
var once_Data_Function_apply__1359309341 sync.Once
func Get_Data_Function_apply__1359309341() gopurs_runtime.Value {
	once_Data_Function_apply__1359309341.Do(func() {
		cache_Data_Function_apply__1359309341 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_apply__1359309341(f_unused_0_box, x_1_box)
})
	})
	return cache_Data_Function_apply__1359309341
}

var cache_Data_Function_apply__429089108 gopurs_runtime.Value
var once_Data_Function_apply__429089108 sync.Once
func Get_Data_Function_apply__429089108() gopurs_runtime.Value {
	once_Data_Function_apply__429089108.Do(func() {
		cache_Data_Function_apply__429089108 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_apply__429089108(f_unused_0_box, x_unused_1_box)
})
	})
	return cache_Data_Function_apply__429089108
}

var cache_Data_Function_apply__204083284 gopurs_runtime.Value
var once_Data_Function_apply__204083284 sync.Once
func Get_Data_Function_apply__204083284() gopurs_runtime.Value {
	once_Data_Function_apply__204083284.Do(func() {
		cache_Data_Function_apply__204083284 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_apply__204083284(f_unused_0_box, x_unused_1_box)
})
	})
	return cache_Data_Function_apply__204083284
}

var cache_Data_Function_apply__1790022004 gopurs_runtime.Value
var once_Data_Function_apply__1790022004 sync.Once
func Get_Data_Function_apply__1790022004() gopurs_runtime.Value {
	once_Data_Function_apply__1790022004.Do(func() {
		cache_Data_Function_apply__1790022004 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_apply__1790022004(f_unused_0_box, x_unused_1_box)
})
	})
	return cache_Data_Function_apply__1790022004
}

var cache_Data_Function_apply__352423700 gopurs_runtime.Value
var once_Data_Function_apply__352423700 sync.Once
func Get_Data_Function_apply__352423700() gopurs_runtime.Value {
	once_Data_Function_apply__352423700.Do(func() {
		cache_Data_Function_apply__352423700 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_apply__352423700(f_unused_0_box, x_unused_1_box)
})
	})
	return cache_Data_Function_apply__352423700
}

var cache_Data_Function_apply__2706583508 gopurs_runtime.Value
var once_Data_Function_apply__2706583508 sync.Once
func Get_Data_Function_apply__2706583508() gopurs_runtime.Value {
	once_Data_Function_apply__2706583508.Do(func() {
		cache_Data_Function_apply__2706583508 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_apply__2706583508(f_unused_0_box, x_unused_1_box)
})
	})
	return cache_Data_Function_apply__2706583508
}

var cache_Data_Function_apply__818091412 gopurs_runtime.Value
var once_Data_Function_apply__818091412 sync.Once
func Get_Data_Function_apply__818091412() gopurs_runtime.Value {
	once_Data_Function_apply__818091412.Do(func() {
		cache_Data_Function_apply__818091412 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_apply__818091412(f_unused_0_box, x_unused_1_box)
})
	})
	return cache_Data_Function_apply__818091412
}

var cache_Data_Function_apply__2406532916 gopurs_runtime.Value
var once_Data_Function_apply__2406532916 sync.Once
func Get_Data_Function_apply__2406532916() gopurs_runtime.Value {
	once_Data_Function_apply__2406532916.Do(func() {
		cache_Data_Function_apply__2406532916 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_apply__2406532916(f_unused_0_box, x_unused_1_box)
})
	})
	return cache_Data_Function_apply__2406532916
}

var cache_Data_Function_apply__2745903828 gopurs_runtime.Value
var once_Data_Function_apply__2745903828 sync.Once
func Get_Data_Function_apply__2745903828() gopurs_runtime.Value {
	once_Data_Function_apply__2745903828.Do(func() {
		cache_Data_Function_apply__2745903828 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_apply__2745903828(f_unused_0_box, x_unused_1_box)
})
	})
	return cache_Data_Function_apply__2745903828
}

var cache_Data_Function_apply__4001191444 gopurs_runtime.Value
var once_Data_Function_apply__4001191444 sync.Once
func Get_Data_Function_apply__4001191444() gopurs_runtime.Value {
	once_Data_Function_apply__4001191444.Do(func() {
		cache_Data_Function_apply__4001191444 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_apply__4001191444(f_unused_0_box, x_unused_1_box)
})
	})
	return cache_Data_Function_apply__4001191444
}

var cache_Data_Function_apply__4162337140 gopurs_runtime.Value
var once_Data_Function_apply__4162337140 sync.Once
func Get_Data_Function_apply__4162337140() gopurs_runtime.Value {
	once_Data_Function_apply__4162337140.Do(func() {
		cache_Data_Function_apply__4162337140 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_apply__4162337140(f_unused_0_box, x_unused_1_box)
})
	})
	return cache_Data_Function_apply__4162337140
}

var cache_Data_Function_apply__522457044 gopurs_runtime.Value
var once_Data_Function_apply__522457044 sync.Once
func Get_Data_Function_apply__522457044() gopurs_runtime.Value {
	once_Data_Function_apply__522457044.Do(func() {
		cache_Data_Function_apply__522457044 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_apply__522457044(f_unused_0_box, x_unused_1_box)
})
	})
	return cache_Data_Function_apply__522457044
}

var cache_Data_Function_apply__865585108 gopurs_runtime.Value
var once_Data_Function_apply__865585108 sync.Once
func Get_Data_Function_apply__865585108() gopurs_runtime.Value {
	once_Data_Function_apply__865585108.Do(func() {
		cache_Data_Function_apply__865585108 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_apply__865585108(f_unused_0_box, x_unused_1_box)
})
	})
	return cache_Data_Function_apply__865585108
}

var cache_Data_Function_apply__2214658004 gopurs_runtime.Value
var once_Data_Function_apply__2214658004 sync.Once
func Get_Data_Function_apply__2214658004() gopurs_runtime.Value {
	once_Data_Function_apply__2214658004.Do(func() {
		cache_Data_Function_apply__2214658004 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_apply__2214658004(f_unused_0_box, x_unused_1_box)
})
	})
	return cache_Data_Function_apply__2214658004
}

var cache_Data_Function_apply__1683855444 gopurs_runtime.Value
var once_Data_Function_apply__1683855444 sync.Once
func Get_Data_Function_apply__1683855444() gopurs_runtime.Value {
	once_Data_Function_apply__1683855444.Do(func() {
		cache_Data_Function_apply__1683855444 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_apply__1683855444(f_unused_0_box, x_unused_1_box)
})
	})
	return cache_Data_Function_apply__1683855444
}

var cache_Data_Function_apply__295373620 gopurs_runtime.Value
var once_Data_Function_apply__295373620 sync.Once
func Get_Data_Function_apply__295373620() gopurs_runtime.Value {
	once_Data_Function_apply__295373620.Do(func() {
		cache_Data_Function_apply__295373620 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_apply__295373620(f_unused_0_box, x_unused_1_box)
})
	})
	return cache_Data_Function_apply__295373620
}

var cache_Data_Function_apply__3214919060 gopurs_runtime.Value
var once_Data_Function_apply__3214919060 sync.Once
func Get_Data_Function_apply__3214919060() gopurs_runtime.Value {
	once_Data_Function_apply__3214919060.Do(func() {
		cache_Data_Function_apply__3214919060 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_apply__3214919060(f_unused_0_box, x_unused_1_box)
})
	})
	return cache_Data_Function_apply__3214919060
}

var cache_Data_Function_apply__568573268 gopurs_runtime.Value
var once_Data_Function_apply__568573268 sync.Once
func Get_Data_Function_apply__568573268() gopurs_runtime.Value {
	once_Data_Function_apply__568573268.Do(func() {
		cache_Data_Function_apply__568573268 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_apply__568573268(f_unused_0_box, x_unused_1_box)
})
	})
	return cache_Data_Function_apply__568573268
}

var cache_Data_Function_apply__2635643092 gopurs_runtime.Value
var once_Data_Function_apply__2635643092 sync.Once
func Get_Data_Function_apply__2635643092() gopurs_runtime.Value {
	once_Data_Function_apply__2635643092.Do(func() {
		cache_Data_Function_apply__2635643092 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_apply__2635643092(f_unused_0_box, x_unused_1_box)
})
	})
	return cache_Data_Function_apply__2635643092
}

var cache_Data_Function_apply__475876628 gopurs_runtime.Value
var once_Data_Function_apply__475876628 sync.Once
func Get_Data_Function_apply__475876628() gopurs_runtime.Value {
	once_Data_Function_apply__475876628.Do(func() {
		cache_Data_Function_apply__475876628 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_apply__475876628(f_unused_0_box, x_unused_1_box)
})
	})
	return cache_Data_Function_apply__475876628
}

var cache_Data_Function_apply__3933388244 gopurs_runtime.Value
var once_Data_Function_apply__3933388244 sync.Once
func Get_Data_Function_apply__3933388244() gopurs_runtime.Value {
	once_Data_Function_apply__3933388244.Do(func() {
		cache_Data_Function_apply__3933388244 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_apply__3933388244(f_unused_0_box, x_unused_1_box)
})
	})
	return cache_Data_Function_apply__3933388244
}

var cache_Data_Function_apply__2893877524 gopurs_runtime.Value
var once_Data_Function_apply__2893877524 sync.Once
func Get_Data_Function_apply__2893877524() gopurs_runtime.Value {
	once_Data_Function_apply__2893877524.Do(func() {
		cache_Data_Function_apply__2893877524 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_apply__2893877524(f_unused_0_box, x_unused_1_box)
})
	})
	return cache_Data_Function_apply__2893877524
}

var cache_Data_Function_apply__2381085460 gopurs_runtime.Value
var once_Data_Function_apply__2381085460 sync.Once
func Get_Data_Function_apply__2381085460() gopurs_runtime.Value {
	once_Data_Function_apply__2381085460.Do(func() {
		cache_Data_Function_apply__2381085460 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_apply__2381085460(f_unused_0_box, x_unused_1_box)
})
	})
	return cache_Data_Function_apply__2381085460
}

var cache_Data_Function_apply__48294676 gopurs_runtime.Value
var once_Data_Function_apply__48294676 sync.Once
func Get_Data_Function_apply__48294676() gopurs_runtime.Value {
	once_Data_Function_apply__48294676.Do(func() {
		cache_Data_Function_apply__48294676 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_apply__48294676(f_unused_0_box, x_unused_1_box)
})
	})
	return cache_Data_Function_apply__48294676
}

var cache_Data_Function_apply__2583419188 gopurs_runtime.Value
var once_Data_Function_apply__2583419188 sync.Once
func Get_Data_Function_apply__2583419188() gopurs_runtime.Value {
	once_Data_Function_apply__2583419188.Do(func() {
		cache_Data_Function_apply__2583419188 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_apply__2583419188(f_unused_0_box, x_unused_1_box)
})
	})
	return cache_Data_Function_apply__2583419188
}

var cache_Data_Function_apply__3463278420 gopurs_runtime.Value
var once_Data_Function_apply__3463278420 sync.Once
func Get_Data_Function_apply__3463278420() gopurs_runtime.Value {
	once_Data_Function_apply__3463278420.Do(func() {
		cache_Data_Function_apply__3463278420 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_apply__3463278420(f_unused_0_box, x_unused_1_box)
})
	})
	return cache_Data_Function_apply__3463278420
}

var cache_Data_Function_apply__1305127060 gopurs_runtime.Value
var once_Data_Function_apply__1305127060 sync.Once
func Get_Data_Function_apply__1305127060() gopurs_runtime.Value {
	once_Data_Function_apply__1305127060.Do(func() {
		cache_Data_Function_apply__1305127060 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_apply__1305127060(f_unused_0_box, x_unused_1_box)
})
	})
	return cache_Data_Function_apply__1305127060
}

var cache_Data_Function_apply__1570647764 gopurs_runtime.Value
var once_Data_Function_apply__1570647764 sync.Once
func Get_Data_Function_apply__1570647764() gopurs_runtime.Value {
	once_Data_Function_apply__1570647764.Do(func() {
		cache_Data_Function_apply__1570647764 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_apply__1570647764(f_unused_0_box, x_unused_1_box)
})
	})
	return cache_Data_Function_apply__1570647764
}

var cache_Data_Function_apply__2218396692 gopurs_runtime.Value
var once_Data_Function_apply__2218396692 sync.Once
func Get_Data_Function_apply__2218396692() gopurs_runtime.Value {
	once_Data_Function_apply__2218396692.Do(func() {
		cache_Data_Function_apply__2218396692 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_apply__2218396692(f_unused_0_box, x_unused_1_box)
})
	})
	return cache_Data_Function_apply__2218396692
}

var cache_Data_Function_apply__3626011640 gopurs_runtime.Value
var once_Data_Function_apply__3626011640 sync.Once
func Get_Data_Function_apply__3626011640() gopurs_runtime.Value {
	once_Data_Function_apply__3626011640.Do(func() {
		cache_Data_Function_apply__3626011640 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_apply__3626011640(f_unused_0_box, x_1_box)
})
	})
	return cache_Data_Function_apply__3626011640
}

var cache_Data_Function_apply__1597510678 gopurs_runtime.Value
var once_Data_Function_apply__1597510678 sync.Once
func Get_Data_Function_apply__1597510678() gopurs_runtime.Value {
	once_Data_Function_apply__1597510678.Do(func() {
		cache_Data_Function_apply__1597510678 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_Data_Function_apply__1597510678(f_unused_0_box, func() []int64 {
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
	return cache_Data_Function_apply__1597510678
}

var cache_Data_Function_apply__508702885 gopurs_runtime.Value
var once_Data_Function_apply__508702885 sync.Once
func Get_Data_Function_apply__508702885() gopurs_runtime.Value {
	once_Data_Function_apply__508702885.Do(func() {
		cache_Data_Function_apply__508702885 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Function_apply__508702885(f_0_box, (x_1_box.IntVal) != (0)))
})
	})
	return cache_Data_Function_apply__508702885
}

var cache_Data_Function_apply__1422799008 gopurs_runtime.Value
var once_Data_Function_apply__1422799008 sync.Once
func Get_Data_Function_apply__1422799008() gopurs_runtime.Value {
	once_Data_Function_apply__1422799008.Do(func() {
		cache_Data_Function_apply__1422799008 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_apply__1422799008(f_unused_0_box, x_1_box)
})
	})
	return cache_Data_Function_apply__1422799008
}

var cache_Data_Function_apply__3557122740 gopurs_runtime.Value
var once_Data_Function_apply__3557122740 sync.Once
func Get_Data_Function_apply__3557122740() gopurs_runtime.Value {
	once_Data_Function_apply__3557122740.Do(func() {
		cache_Data_Function_apply__3557122740 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(Call_Data_Function_apply__3557122740(f_unused_0_box, x_1_box.FloatVal()))
})
	})
	return cache_Data_Function_apply__3557122740
}

var cache_Data_Function_apply__1427044821 gopurs_runtime.Value
var once_Data_Function_apply__1427044821 sync.Once
func Get_Data_Function_apply__1427044821() gopurs_runtime.Value {
	once_Data_Function_apply__1427044821.Do(func() {
		cache_Data_Function_apply__1427044821 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(Call_Data_Function_apply__1427044821(f_unused_0_box, x_1_box.FloatVal()))
})
	})
	return cache_Data_Function_apply__1427044821
}

var cache_Data_Function_apply__927689573 gopurs_runtime.Value
var once_Data_Function_apply__927689573 sync.Once
func Get_Data_Function_apply__927689573() gopurs_runtime.Value {
	once_Data_Function_apply__927689573.Do(func() {
		cache_Data_Function_apply__927689573 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Function_apply__927689573(f_0_box, x_1_box.StrVal())
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
})
	})
	return cache_Data_Function_apply__927689573
}

var cache_Data_Function_apply__1361677546 gopurs_runtime.Value
var once_Data_Function_apply__1361677546 sync.Once
func Get_Data_Function_apply__1361677546() gopurs_runtime.Value {
	once_Data_Function_apply__1361677546.Do(func() {
		cache_Data_Function_apply__1361677546 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_apply__1361677546(f_unused_0_box, x_1_box.StrVal())
})
	})
	return cache_Data_Function_apply__1361677546
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

func Call_Data_Function_flip__1757181342(f_unused_0_loop gopurs_runtime.Value, b_1_loop *Constructor_Data_Time_Time, a_2_loop *Constructor_Data_Date_Date) *Constructor_Data_DateTime_DateTime {
flip__1757181342:
for {
if false { continue flip__1757181342 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var b_1 *Constructor_Data_Time_Time = b_1_loop
_ = b_1
var a_2 *Constructor_Data_Date_Date = a_2_loop
_ = a_2
return (&Constructor_Data_DateTime_DateTime{1, a_2, b_1})
}
}

func Call_Data_Function_flip__2357332471(f_unused_0_loop gopurs_runtime.Value, b_1_loop *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]], a_2_loop gopurs_runtime.Value) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
flip__2357332471:
for {
if false { continue flip__2357332471 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var b_1 *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(Get_Data_Interval_Duration_Iso_foldMap(), a_2, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_Function_2442833393_849153993(b_1))}))
}
}

func Call_Data_Function_flip__4138356370(f_unused_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
flip__4138356370:
for {
if false { continue flip__4138356370 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
var __t2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](a_2)
if (__t_tag_0 != nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(b_1, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(a_2.UnsafePtr).V0))
goto end_branch_2
} else {

}
}
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](a_2)
if (__t_tag_1 == nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_2
} else {

}
}
{
__t2 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_2:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Function_flip__965230301(f_unused_0_loop gopurs_runtime.Value, b_unused_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
flip__965230301:
for {
if false { continue flip__965230301 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var b_unused_1 gopurs_runtime.Value = b_unused_1_loop
_ = b_unused_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
var __t3 gopurs_runtime.Value
{
var __t_tag_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](a_2)
if (__t_tag_0 != nil) {
var __t1 gopurs_runtime.Value
{
if ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(a_2.UnsafePtr).V0.StrVal()) == ("") {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Function_742090555_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[string]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Str((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(a_2.UnsafePtr).V0.StrVal())}))}
}
end_branch_1:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Function_742090555_3094389156(Rebox_Data_Function_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t1))))}
goto end_branch_3
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](a_2)
if (__t_tag_2 == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }())}
}
end_branch_3:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t3))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Function_flip__746081925(f_0_loop gopurs_runtime.Value, b_1_loop *Constructor_Data_Date_Date, a_2_loop int64) struct{V0 gopurs_runtime.Value; V1 bool} {
flip__746081925:
for {
if false { continue flip__746081925 }
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 *Constructor_Data_Date_Date = b_1_loop
_ = b_1
var a_2 int64 = a_2_loop
_ = a_2
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Function_2280409795_3094389156(Rebox_Data_Function_3094389156_2280409795(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(f_0, gopurs_runtime.Int(a_2), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(b_1)})))))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
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

func Call_Data_Function_const__254090813(a_0_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value], v_1_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
const__254090813:
for {
if false { continue const__254090813 }
var a_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(a_0)}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Function_const__2942423693(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
const__2942423693:
for {
if false { continue const__2942423693 }
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}
}

func Call_Data_Function_const__235433565(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
const__235433565:
for {
if false { continue const__235433565 }
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}
}

func Call_Data_Function_const__549142733(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
const__549142733:
for {
if false { continue const__549142733 }
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}
}

func Call_Data_Function_const__2135255069(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
const__2135255069:
for {
if false { continue const__2135255069 }
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}
}

func Call_Data_Function_const__4293180849(a_0_loop bool, v_1_loop uint32) bool {
const__4293180849:
for {
if false { continue const__4293180849 }
var a_0 bool = a_0_loop
_ = a_0
var v_1 uint32 = v_1_loop
_ = v_1
return a_0
}
}

func Call_Data_Function_const__1421211688(a_0_loop bool, v_1_loop int64) bool {
const__1421211688:
for {
if false { continue const__1421211688 }
var a_0 bool = a_0_loop
_ = a_0
var v_1 int64 = v_1_loop
_ = v_1
return a_0
}
}

func Call_Data_Function_const__2519857944(a_0_loop bool, v_1_loop float64) bool {
const__2519857944:
for {
if false { continue const__2519857944 }
var a_0 bool = a_0_loop
_ = a_0
var v_1 float64 = v_1_loop
_ = v_1
return a_0
}
}

func Call_Data_Function_const__3022713485(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
const__3022713485:
for {
if false { continue const__3022713485 }
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}
}

func Call_Data_Function_const__3052583336(a_0_loop int64, v_1_loop int64) int64 {
const__3052583336:
for {
if false { continue const__3052583336 }
var a_0 int64 = a_0_loop
_ = a_0
var v_1 int64 = v_1_loop
_ = v_1
return a_0
}
}

func Call_Data_Function_const__1592876733(a_0_loop int64, v_1_loop gopurs_runtime.Value) int64 {
const__1592876733:
for {
if false { continue const__1592876733 }
var a_0 int64 = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}
}

func Call_Data_Function_const__3126880413(a_0_loop string, v_1_loop gopurs_runtime.Value) string {
const__3126880413:
for {
if false { continue const__3126880413 }
var a_0 string = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}
}

func Call_Data_Function_const__2623567125(a_unused_0_loop gopurs_runtime.Value, v_1_loop string) gopurs_runtime.Value {
const__2623567125:
for {
if false { continue const__2623567125 }
var a_unused_0 gopurs_runtime.Value = a_unused_0_loop
_ = a_unused_0
var v_1 string = v_1_loop
_ = v_1
return Get_Data_Unit_unit()
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

func Call_Data_Function_apply__3981898766(f_unused_0_loop gopurs_runtime.Value, x_1_loop *Constructor_Data_Date_Date) *Constructor_Data_Date_Date {
apply__3981898766:
for {
if false { continue apply__3981898766 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var x_1 *Constructor_Data_Date_Date = x_1_loop
_ = x_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(x_1)}))
}
}

func Call_Data_Function_apply__3720661445(f_0_loop gopurs_runtime.Value, x_1_loop *Constructor_Data_Maybe_Just[*Constructor_Data_Time_Time]) *Constructor_Data_Time_Time {
apply__3720661445:
for {
if false { continue apply__3720661445 }
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 *Constructor_Data_Maybe_Just[*Constructor_Data_Time_Time] = x_1_loop
_ = x_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Function_3839235747_3094389156(x_1))}))
}
}

func Call_Data_Function_apply__2361363086(f_unused_0_loop gopurs_runtime.Value, x_1_loop *Constructor_Data_Maybe_Just[int64]) bool {
apply__2361363086:
for {
if false { continue apply__2361363086 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var x_1 *Constructor_Data_Maybe_Just[int64] = x_1_loop
_ = x_1
var __t0 gopurs_runtime.Value
{
if (x_1 == nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
if (x_1 != nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return (__t0.IntVal) != (0)
}
}

func Call_Data_Function_apply__2704629163(f_unused_0_loop gopurs_runtime.Value, x_1_loop *Constructor_Data_Maybe_Just[int64]) bool {
apply__2704629163:
for {
if false { continue apply__2704629163 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var x_1 *Constructor_Data_Maybe_Just[int64] = x_1_loop
_ = x_1
var __t0 gopurs_runtime.Value
{
if (x_1 == nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
if (x_1 != nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return (__t0.IntVal) != (0)
}
}

func Call_Data_Function_apply__1359309341(f_unused_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__1359309341:
for {
if false { continue apply__1359309341 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(Get_Effect_Aff_void1(), x_1)
}
}

func Call_Data_Function_apply__429089108(f_unused_0_loop gopurs_runtime.Value, x_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__429089108:
for {
if false { continue apply__429089108 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var x_unused_1 gopurs_runtime.Value = x_unused_1_loop
_ = x_unused_1
return gopurs_runtime.Apply(Get_AppFFI_void(), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_AckermannFFICheatcode_describe(), Get_Test_AckermannFFICheatcode_act()))
}
}

func Call_Data_Function_apply__204083284(f_unused_0_loop gopurs_runtime.Value, x_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__204083284:
for {
if false { continue apply__204083284 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var x_unused_1 gopurs_runtime.Value = x_unused_1_loop
_ = x_unused_1
return gopurs_runtime.Apply(Get_AppFFI_void(), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_AckermannFFI_describe(), Get_Test_AckermannFFI_act()))
}
}

func Call_Data_Function_apply__1790022004(f_unused_0_loop gopurs_runtime.Value, x_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__1790022004:
for {
if false { continue apply__1790022004 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var x_unused_1 gopurs_runtime.Value = x_unused_1_loop
_ = x_unused_1
return gopurs_runtime.Apply(Get_AppFFI_void(), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_ArrayOpsFFICheatcode_describe(), Get_Test_ArrayOpsFFICheatcode_act()))
}
}

func Call_Data_Function_apply__352423700(f_unused_0_loop gopurs_runtime.Value, x_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__352423700:
for {
if false { continue apply__352423700 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var x_unused_1 gopurs_runtime.Value = x_unused_1_loop
_ = x_unused_1
return gopurs_runtime.Apply(Get_AppFFI_void(), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_ArrayOpsFFI_describe(), Get_Test_ArrayOpsFFI_act()))
}
}

func Call_Data_Function_apply__2706583508(f_unused_0_loop gopurs_runtime.Value, x_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__2706583508:
for {
if false { continue apply__2706583508 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var x_unused_1 gopurs_runtime.Value = x_unused_1_loop
_ = x_unused_1
return gopurs_runtime.Apply(Get_AppFFI_void(), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_AstTreeFFICheatcode_describe(), Get_Test_AstTreeFFICheatcode_act()))
}
}

func Call_Data_Function_apply__818091412(f_unused_0_loop gopurs_runtime.Value, x_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__818091412:
for {
if false { continue apply__818091412 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var x_unused_1 gopurs_runtime.Value = x_unused_1_loop
_ = x_unused_1
return gopurs_runtime.Apply(Get_AppFFI_void(), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_AstTreeFFI_describe(), Get_Test_AstTreeFFI_act()))
}
}

func Call_Data_Function_apply__2406532916(f_unused_0_loop gopurs_runtime.Value, x_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__2406532916:
for {
if false { continue apply__2406532916 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var x_unused_1 gopurs_runtime.Value = x_unused_1_loop
_ = x_unused_1
return gopurs_runtime.Apply(Get_AppFFI_void(), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_ChurchFFICheatcode_describe(), Get_Test_ChurchFFICheatcode_act()))
}
}

func Call_Data_Function_apply__2745903828(f_unused_0_loop gopurs_runtime.Value, x_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__2745903828:
for {
if false { continue apply__2745903828 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var x_unused_1 gopurs_runtime.Value = x_unused_1_loop
_ = x_unused_1
return gopurs_runtime.Apply(Get_AppFFI_void(), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_ChurchFFI_describe(), Get_Test_ChurchFFI_act()))
}
}

func Call_Data_Function_apply__4001191444(f_unused_0_loop gopurs_runtime.Value, x_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__4001191444:
for {
if false { continue apply__4001191444 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var x_unused_1 gopurs_runtime.Value = x_unused_1_loop
_ = x_unused_1
return gopurs_runtime.Apply(Get_AppFFI_void(), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_FibFFICheatcode_describe(), Get_Test_FibFFICheatcode_act()))
}
}

func Call_Data_Function_apply__4162337140(f_unused_0_loop gopurs_runtime.Value, x_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__4162337140:
for {
if false { continue apply__4162337140 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var x_unused_1 gopurs_runtime.Value = x_unused_1_loop
_ = x_unused_1
return gopurs_runtime.Apply(Get_AppFFI_void(), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_FibFFI_describe(), Get_Test_FibFFI_act()))
}
}

func Call_Data_Function_apply__522457044(f_unused_0_loop gopurs_runtime.Value, x_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__522457044:
for {
if false { continue apply__522457044 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var x_unused_1 gopurs_runtime.Value = x_unused_1_loop
_ = x_unused_1
return gopurs_runtime.Apply(Get_AppFFI_void(), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_LazyEvaluationFFICheatcode_describe(), Get_Test_LazyEvaluationFFICheatcode_act()))
}
}

func Call_Data_Function_apply__865585108(f_unused_0_loop gopurs_runtime.Value, x_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__865585108:
for {
if false { continue apply__865585108 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var x_unused_1 gopurs_runtime.Value = x_unused_1_loop
_ = x_unused_1
return gopurs_runtime.Apply(Get_AppFFI_void(), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_LazyEvaluationFFI_describe(), Get_Test_LazyEvaluationFFI_act()))
}
}

func Call_Data_Function_apply__2214658004(f_unused_0_loop gopurs_runtime.Value, x_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__2214658004:
for {
if false { continue apply__2214658004 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var x_unused_1 gopurs_runtime.Value = x_unused_1_loop
_ = x_unused_1
return gopurs_runtime.Apply(Get_AppFFI_void(), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_ListOpsFFICheatcode_describe(), Get_Test_ListOpsFFICheatcode_act()))
}
}

func Call_Data_Function_apply__1683855444(f_unused_0_loop gopurs_runtime.Value, x_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__1683855444:
for {
if false { continue apply__1683855444 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var x_unused_1 gopurs_runtime.Value = x_unused_1_loop
_ = x_unused_1
return gopurs_runtime.Apply(Get_AppFFI_void(), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_ListOpsFFI_describe(), Get_Test_ListOpsFFI_act()))
}
}

func Call_Data_Function_apply__295373620(f_unused_0_loop gopurs_runtime.Value, x_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__295373620:
for {
if false { continue apply__295373620 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var x_unused_1 gopurs_runtime.Value = x_unused_1_loop
_ = x_unused_1
return gopurs_runtime.Apply(Get_AppFFI_void(), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_PolymorphismFFICheatcode_describe(), Get_Test_PolymorphismFFICheatcode_act()))
}
}

func Call_Data_Function_apply__3214919060(f_unused_0_loop gopurs_runtime.Value, x_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__3214919060:
for {
if false { continue apply__3214919060 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var x_unused_1 gopurs_runtime.Value = x_unused_1_loop
_ = x_unused_1
return gopurs_runtime.Apply(Get_AppFFI_void(), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_PolymorphismFFI_describe(), Get_Test_PolymorphismFFI_act()))
}
}

func Call_Data_Function_apply__568573268(f_unused_0_loop gopurs_runtime.Value, x_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__568573268:
for {
if false { continue apply__568573268 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var x_unused_1 gopurs_runtime.Value = x_unused_1_loop
_ = x_unused_1
return gopurs_runtime.Apply(Get_AppFFI_void(), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_PrimesFFICheatcode_describe(), Get_Test_PrimesFFICheatcode_act()))
}
}

func Call_Data_Function_apply__2635643092(f_unused_0_loop gopurs_runtime.Value, x_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__2635643092:
for {
if false { continue apply__2635643092 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var x_unused_1 gopurs_runtime.Value = x_unused_1_loop
_ = x_unused_1
return gopurs_runtime.Apply(Get_AppFFI_void(), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_PrimesFFI_describe(), Get_Test_PrimesFFI_act()))
}
}

func Call_Data_Function_apply__475876628(f_unused_0_loop gopurs_runtime.Value, x_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__475876628:
for {
if false { continue apply__475876628 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var x_unused_1 gopurs_runtime.Value = x_unused_1_loop
_ = x_unused_1
return gopurs_runtime.Apply(Get_AppFFI_void(), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_RBTreeFFICheatcode_describe(), Get_Test_RBTreeFFICheatcode_act()))
}
}

func Call_Data_Function_apply__3933388244(f_unused_0_loop gopurs_runtime.Value, x_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__3933388244:
for {
if false { continue apply__3933388244 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var x_unused_1 gopurs_runtime.Value = x_unused_1_loop
_ = x_unused_1
return gopurs_runtime.Apply(Get_AppFFI_void(), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_RBTreeFFI_describe(), Get_Test_RBTreeFFI_act()))
}
}

func Call_Data_Function_apply__2893877524(f_unused_0_loop gopurs_runtime.Value, x_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__2893877524:
for {
if false { continue apply__2893877524 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var x_unused_1 gopurs_runtime.Value = x_unused_1_loop
_ = x_unused_1
return gopurs_runtime.Apply(Get_AppFFI_void(), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_RecordsFFICheatcode_describe(), Get_Test_RecordsFFICheatcode_act()))
}
}

func Call_Data_Function_apply__2381085460(f_unused_0_loop gopurs_runtime.Value, x_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__2381085460:
for {
if false { continue apply__2381085460 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var x_unused_1 gopurs_runtime.Value = x_unused_1_loop
_ = x_unused_1
return gopurs_runtime.Apply(Get_AppFFI_void(), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_RecordsFFI_describe(), Get_Test_RecordsFFI_act()))
}
}

func Call_Data_Function_apply__48294676(f_unused_0_loop gopurs_runtime.Value, x_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__48294676:
for {
if false { continue apply__48294676 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var x_unused_1 gopurs_runtime.Value = x_unused_1_loop
_ = x_unused_1
return gopurs_runtime.Apply(Get_AppFFI_void(), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_RowToListFFICheatcode_describe(), Get_Test_RowToListFFICheatcode_act()))
}
}

func Call_Data_Function_apply__2583419188(f_unused_0_loop gopurs_runtime.Value, x_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__2583419188:
for {
if false { continue apply__2583419188 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var x_unused_1 gopurs_runtime.Value = x_unused_1_loop
_ = x_unused_1
return gopurs_runtime.Apply(Get_AppFFI_void(), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_RowToListFFI_describe(), Get_Test_RowToListFFI_act()))
}
}

func Call_Data_Function_apply__3463278420(f_unused_0_loop gopurs_runtime.Value, x_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__3463278420:
for {
if false { continue apply__3463278420 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var x_unused_1 gopurs_runtime.Value = x_unused_1_loop
_ = x_unused_1
return gopurs_runtime.Apply(Get_AppFFI_void(), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_StateMonadFFICheatcode_describe(), Get_Test_StateMonadFFICheatcode_act()))
}
}

func Call_Data_Function_apply__1305127060(f_unused_0_loop gopurs_runtime.Value, x_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__1305127060:
for {
if false { continue apply__1305127060 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var x_unused_1 gopurs_runtime.Value = x_unused_1_loop
_ = x_unused_1
return gopurs_runtime.Apply(Get_AppFFI_void(), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_StateMonadFFI_describe(), Get_Test_StateMonadFFI_act()))
}
}

func Call_Data_Function_apply__1570647764(f_unused_0_loop gopurs_runtime.Value, x_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__1570647764:
for {
if false { continue apply__1570647764 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var x_unused_1 gopurs_runtime.Value = x_unused_1_loop
_ = x_unused_1
return gopurs_runtime.Apply(Get_AppFFI_void(), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_TCOFFICheatcode_describe(), Get_Test_TCOFFICheatcode_act()))
}
}

func Call_Data_Function_apply__2218396692(f_unused_0_loop gopurs_runtime.Value, x_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__2218396692:
for {
if false { continue apply__2218396692 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var x_unused_1 gopurs_runtime.Value = x_unused_1_loop
_ = x_unused_1
return gopurs_runtime.Apply(Get_AppFFI_void(), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_TCOFFI_describe(), Get_Test_TCOFFI_act()))
}
}

func Call_Data_Function_apply__3626011640(f_unused_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__3626011640:
for {
if false { continue apply__3626011640 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), x_1)
}
}

func Call_Data_Function_apply__1597510678(f_unused_0_loop gopurs_runtime.Value, x_1_loop []int64) []int64 {
apply__1597510678:
for {
if false { continue apply__1597510678 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var x_1 []int64 = x_1_loop
_ = x_1
return func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
					arr := x_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}().UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}()
}
}

func Call_Data_Function_apply__508702885(f_0_loop gopurs_runtime.Value, x_1_loop bool) bool {
apply__508702885:
for {
if false { continue apply__508702885 }
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 bool = x_1_loop
_ = x_1
return (gopurs_runtime.Apply(f_0, gopurs_runtime.Bool(x_1)).IntVal) != (0)
}
}

func Call_Data_Function_apply__1422799008(f_unused_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__1422799008:
for {
if false { continue apply__1422799008 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
})
}
}

func Call_Data_Function_apply__3557122740(f_unused_0_loop gopurs_runtime.Value, x_1_loop float64) float64 {
apply__3557122740:
for {
if false { continue apply__3557122740 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var x_1 float64 = x_1_loop
_ = x_1
return x_1
}
}

func Call_Data_Function_apply__1427044821(f_unused_0_loop gopurs_runtime.Value, x_1_loop float64) float64 {
apply__1427044821:
for {
if false { continue apply__1427044821 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var x_1 float64 = x_1_loop
_ = x_1
return x_1
}
}

func Call_Data_Function_apply__927689573(f_0_loop gopurs_runtime.Value, x_1_loop string) struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool} {
apply__927689573:
for {
if false { continue apply__927689573 }
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 string = x_1_loop
_ = x_1
return func() struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool} {
				_v := gopurs_runtime.Apply(f_0, gopurs_runtime.Str(x_1))
				if _v.Type == 9 && _v.IntVal == 3234899973 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{V0: gopurs_runtime.Value{}, V1: (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr).V0, V2: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{V0: (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: gopurs_runtime.Value{}, V2: false}
			}()
}
}

func Call_Data_Function_apply__1361677546(f_unused_0_loop gopurs_runtime.Value, x_1_loop string) gopurs_runtime.Value {
apply__1361677546:
for {
if false { continue apply__1361677546 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var x_1 string = x_1_loop
_ = x_1
return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(x_1))
}
}

func Rebox_Data_Function_2280409795_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}

func Rebox_Data_Function_2442833393_849153993(in *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Function_3132786365_138441832(in.V0))}
		out.V1 = Rebox_Data_Function_2442833393_849153993(in.V1)
	return out
}

func Rebox_Data_Function_3094389156_2280409795(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]{}
		out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](in.V0)
	return out
}

func Rebox_Data_Function_3094389156_742090555(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[string] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[string]{}
		out.V0 = in.V0.StrVal()
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

func Rebox_Data_Function_742090555_3094389156(in *Constructor_Data_Maybe_Just[string]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Str(in.V0)
	return out
}


