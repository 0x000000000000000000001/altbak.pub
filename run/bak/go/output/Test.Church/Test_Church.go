package Test_Church

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Effect "gopurs/output/Effect"
	pkg_Bench "gopurs/output/Bench"
	pkg_Data_Show "gopurs/output/Data.Show"
)

var cache_zeroC gopurs_runtime.Value
var once_zeroC sync.Once
func Get_zeroC() gopurs_runtime.Value {
	once_zeroC.Do(func() {
		cache_zeroC = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_zeroC(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(v_0_box, gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(x_1_box)))
})
	})
	return cache_zeroC
}

var cache_zeroC__func_func_int64__int64__int64__int64_4066693242 gopurs_runtime.Value
var once_zeroC__func_func_int64__int64__int64__int64_4066693242 sync.Once
func Get_zeroC__func_func_int64__int64__int64__int64_4066693242() gopurs_runtime.Value {
	once_zeroC__func_func_int64__int64__int64__int64_4066693242.Do(func() {
		cache_zeroC__func_func_int64__int64__int64__int64_4066693242 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_zeroC__func_func_int64__int64__int64__int64_4066693242(func(inner_arg0 int64) int64 {
return gopurs_runtime.Apply(v_0_box, gopurs_runtime.Int(inner_arg0)).IntVal
}, x_1_box.IntVal))
})
	})
	return cache_zeroC__func_func_int64__int64__int64__int64_4066693242
}

var cache_toInt gopurs_runtime.Value
var once_toInt sync.Once
func Get_toInt() gopurs_runtime.Value {
	once_toInt.Do(func() {
		cache_toInt = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_toInt(func(inner_arg0 func(int64) int64, inner_arg1 int64) int64 {
return gopurs_runtime.Apply2(n_0_box, gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(inner_arg0(arg0.IntVal))
}), gopurs_runtime.Int(inner_arg1)).IntVal
}))
})
	})
	return cache_toInt
}

var cache_succC gopurs_runtime.Value
var once_succC sync.Once
func Get_succC() gopurs_runtime.Value {
	once_succC.Do(func() {
		cache_succC = gopurs_runtime.Func3(func(n_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_succC(func(inner_arg0 func(interface{}) interface{}, inner_arg1 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(n_0_box, gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(inner_arg0(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Any(inner_arg1)))
}, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_1_box, gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(x_2_box)))
})
	})
	return cache_succC
}

var cache_succC__func_func_func_int64__int64__int64__int64__func_int64__int64__int64__int64_4206869978 gopurs_runtime.Value
var once_succC__func_func_func_int64__int64__int64__int64__func_int64__int64__int64__int64_4206869978 sync.Once
func Get_succC__func_func_func_int64__int64__int64__int64__func_int64__int64__int64__int64_4206869978() gopurs_runtime.Value {
	once_succC__func_func_func_int64__int64__int64__int64__func_int64__int64__int64__int64_4206869978.Do(func() {
		cache_succC__func_func_func_int64__int64__int64__int64__func_int64__int64__int64__int64_4206869978 = gopurs_runtime.Func3(func(n_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_succC__func_func_func_int64__int64__int64__int64__func_int64__int64__int64__int64_4206869978(func(inner_arg0 func(int64) int64, inner_arg1 int64) int64 {
return gopurs_runtime.Apply2(n_0_box, gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(inner_arg0(arg0.IntVal))
}), gopurs_runtime.Int(inner_arg1)).IntVal
}, func(inner_arg0 int64) int64 {
return gopurs_runtime.Apply(f_1_box, gopurs_runtime.Int(inner_arg0)).IntVal
}, x_2_box.IntVal))
})
	})
	return cache_succC__func_func_func_int64__int64__int64__int64__func_int64__int64__int64__int64_4206869978
}

var cache_succC__func_func_func_interface____interface____interface____interface____func_interface____interface____interface____interface___3882721754 gopurs_runtime.Value
var once_succC__func_func_func_interface____interface____interface____interface____func_interface____interface____interface____interface___3882721754 sync.Once
func Get_succC__func_func_func_interface____interface____interface____interface____func_interface____interface____interface____interface___3882721754() gopurs_runtime.Value {
	once_succC__func_func_func_interface____interface____interface____interface____func_interface____interface____interface____interface___3882721754.Do(func() {
		cache_succC__func_func_func_interface____interface____interface____interface____func_interface____interface____interface____interface___3882721754 = gopurs_runtime.Func3(func(n_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_succC__func_func_func_interface____interface____interface____interface____func_interface____interface____interface____interface___3882721754(func(inner_arg0 func(interface{}) interface{}, inner_arg1 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(n_0_box, gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(inner_arg0(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Any(inner_arg1)))
}, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_1_box, gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(x_2_box)))
})
	})
	return cache_succC__func_func_func_interface____interface____interface____interface____func_interface____interface____interface____interface___3882721754
}

var cache_mulC gopurs_runtime.Value
var once_mulC sync.Once
func Get_mulC() gopurs_runtime.Value {
	once_mulC.Do(func() {
		cache_mulC = gopurs_runtime.Func4(func(m_0_box gopurs_runtime.Value, n_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value, x_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_mulC(func(inner_arg0 func(interface{}) interface{}, inner_arg1 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(m_0_box, gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(inner_arg0(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Any(inner_arg1)))
}, func(inner_arg0 func(interface{}) interface{}, inner_arg1 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(n_1_box, gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(inner_arg0(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Any(inner_arg1)))
}, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_2_box, gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(x_3_box)))
})
	})
	return cache_mulC
}

var cache_mulC__func_func_func_int64__int64__int64__int64__func_func_int64__int64__int64__int64__func_int64__int64__int64__int64_3478825786 gopurs_runtime.Value
var once_mulC__func_func_func_int64__int64__int64__int64__func_func_int64__int64__int64__int64__func_int64__int64__int64__int64_3478825786 sync.Once
func Get_mulC__func_func_func_int64__int64__int64__int64__func_func_int64__int64__int64__int64__func_int64__int64__int64__int64_3478825786() gopurs_runtime.Value {
	once_mulC__func_func_func_int64__int64__int64__int64__func_func_int64__int64__int64__int64__func_int64__int64__int64__int64_3478825786.Do(func() {
		cache_mulC__func_func_func_int64__int64__int64__int64__func_func_int64__int64__int64__int64__func_int64__int64__int64__int64_3478825786 = gopurs_runtime.Func4(func(m_0_box gopurs_runtime.Value, n_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value, x_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_mulC__func_func_func_int64__int64__int64__int64__func_func_int64__int64__int64__int64__func_int64__int64__int64__int64_3478825786(func(inner_arg0 func(int64) int64, inner_arg1 int64) int64 {
return gopurs_runtime.Apply2(m_0_box, gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(inner_arg0(arg0.IntVal))
}), gopurs_runtime.Int(inner_arg1)).IntVal
}, func(inner_arg0 func(int64) int64, inner_arg1 int64) int64 {
return gopurs_runtime.Apply2(n_1_box, gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(inner_arg0(arg0.IntVal))
}), gopurs_runtime.Int(inner_arg1)).IntVal
}, func(inner_arg0 int64) int64 {
return gopurs_runtime.Apply(f_2_box, gopurs_runtime.Int(inner_arg0)).IntVal
}, x_3_box.IntVal))
})
	})
	return cache_mulC__func_func_func_int64__int64__int64__int64__func_func_int64__int64__int64__int64__func_int64__int64__int64__int64_3478825786
}

var cache_mulC__func_func_func_interface____interface____interface____interface____func_func_interface____interface____interface____interface____func_interface____interface____interface____interface___2338611642 gopurs_runtime.Value
var once_mulC__func_func_func_interface____interface____interface____interface____func_func_interface____interface____interface____interface____func_interface____interface____interface____interface___2338611642 sync.Once
func Get_mulC__func_func_func_interface____interface____interface____interface____func_func_interface____interface____interface____interface____func_interface____interface____interface____interface___2338611642() gopurs_runtime.Value {
	once_mulC__func_func_func_interface____interface____interface____interface____func_func_interface____interface____interface____interface____func_interface____interface____interface____interface___2338611642.Do(func() {
		cache_mulC__func_func_func_interface____interface____interface____interface____func_func_interface____interface____interface____interface____func_interface____interface____interface____interface___2338611642 = gopurs_runtime.Func4(func(m_0_box gopurs_runtime.Value, n_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value, x_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_mulC__func_func_func_interface____interface____interface____interface____func_func_interface____interface____interface____interface____func_interface____interface____interface____interface___2338611642(func(inner_arg0 func(interface{}) interface{}, inner_arg1 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(m_0_box, gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(inner_arg0(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Any(inner_arg1)))
}, func(inner_arg0 func(interface{}) interface{}, inner_arg1 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(n_1_box, gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(inner_arg0(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Any(inner_arg1)))
}, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_2_box, gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(x_3_box)))
})
	})
	return cache_mulC__func_func_func_interface____interface____interface____interface____func_func_interface____interface____interface____interface____func_interface____interface____interface____interface___2338611642
}

var cache_fromInt gopurs_runtime.Value
var once_fromInt sync.Once
func Get_fromInt() gopurs_runtime.Value {
	once_fromInt.Do(func() {
		cache_fromInt = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromInt(v_0_box.IntVal)
})
	})
	return cache_fromInt
}

var cache_describe gopurs_runtime.Value
var once_describe sync.Once
func Get_describe() gopurs_runtime.Value {
	once_describe.Do(func() {
		cache_describe = gopurs_runtime.Func0(func() gopurs_runtime.Value {
return func() gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("Church Numerals (100k Closure Applications):")), nil)
}()
})
	})
	return cache_describe
}

var cache_c10 gopurs_runtime.Value
var once_c10 sync.Once
func Get_c10() gopurs_runtime.Value {
	once_c10.Do(func() {
		cache_c10 = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_c10(n_0_box.IntVal)
})
	})
	return cache_c10
}

var cache_c100 gopurs_runtime.Value
var once_c100 sync.Once
func Get_c100() gopurs_runtime.Value {
	once_c100.Do(func() {
		cache_c100 = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_c100(n_0_box.IntVal)
})
	})
	return cache_c100
}

var cache_c10k gopurs_runtime.Value
var once_c10k sync.Once
func Get_c10k() gopurs_runtime.Value {
	once_c10k.Do(func() {
		cache_c10k = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_c10k(n_0_box.IntVal)
})
	})
	return cache_c10k
}

var cache_c100k gopurs_runtime.Value
var once_c100k sync.Once
func Get_c100k() gopurs_runtime.Value {
	once_c100k.Do(func() {
		cache_c100k = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_c100k(n_0_box.IntVal)
})
	})
	return cache_c100k
}

var cache_addC gopurs_runtime.Value
var once_addC sync.Once
func Get_addC() gopurs_runtime.Value {
	once_addC.Do(func() {
		cache_addC = gopurs_runtime.Func4(func(m_0_box gopurs_runtime.Value, n_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value, x_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_addC(func(inner_arg0 func(interface{}) interface{}, inner_arg1 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(m_0_box, gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(inner_arg0(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Any(inner_arg1)))
}, func(inner_arg0 func(interface{}) interface{}, inner_arg1 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(n_1_box, gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(inner_arg0(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Any(inner_arg1)))
}, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_2_box, gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(x_3_box)))
})
	})
	return cache_addC
}

var cache_act gopurs_runtime.Value
var once_act sync.Once
func Get_act() gopurs_runtime.Value {
	once_act.Do(func() {
		cache_act = gopurs_runtime.Func0(func() gopurs_runtime.Value {
return func() gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), gopurs_runtime.Apply(pkg_Bench.Get_opaque(), gopurs_runtime.Int(10)), gopurs_runtime.Func(func(dummy_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Show.Get_showInt(), "show"), gopurs_runtime.Apply2(Call_c100k(dummy_0.IntVal), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((x_1.IntVal) + (1))
}), gopurs_runtime.Int(0))))
})), nil)
}()
})
	})
	return cache_act
}

func Call_zeroC(v_0_loop func(interface{}) interface{}, x_1_loop interface{}) interface{} {
var v_0 func(interface{}) interface{} = v_0_loop
_ = v_0
var x_1 interface{} = x_1_loop
_ = x_1
return x_1
}

func Call_zeroC__func_func_int64__int64__int64__int64_4066693242(v_0_loop func(int64) int64, x_1_loop int64) int64 {
var v_0 func(int64) int64 = v_0_loop
_ = v_0
var x_1 int64 = x_1_loop
_ = x_1
return x_1
}

func Call_toInt(n_0_loop func(func(int64) int64, int64) int64) int64 {
var n_0 func(func(int64) int64, int64) int64 = n_0_loop
_ = n_0
return gopurs_runtime.Int(n_0(func(inner_arg0 int64) int64 {
return gopurs_runtime.Apply(gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((x_1.IntVal) + (1))
}), gopurs_runtime.Int(inner_arg0)).IntVal
}, 0)).IntVal
}

func Call_succC(n_0_loop func(func(interface{}) interface{}, interface{}) interface{}, f_1_loop func(interface{}) interface{}, x_2_loop interface{}) interface{} {
var n_0 func(func(interface{}) interface{}, interface{}) interface{} = n_0_loop
_ = n_0
var f_1 func(interface{}) interface{} = f_1_loop
_ = f_1
var x_2 interface{} = x_2_loop
_ = x_2
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(f_1(gopurs_runtime.UnboxAny(gopurs_runtime.Any(n_0(f_1, x_2))))))
}

func Call_succC__func_func_func_int64__int64__int64__int64__func_int64__int64__int64__int64_4206869978(n_0_loop func(func(int64) int64, int64) int64, f_1_loop func(int64) int64, x_2_loop int64) int64 {
var n_0 func(func(int64) int64, int64) int64 = n_0_loop
_ = n_0
var f_1 func(int64) int64 = f_1_loop
_ = f_1
var x_2 int64 = x_2_loop
_ = x_2
return gopurs_runtime.Int(f_1(gopurs_runtime.Int(n_0(f_1, x_2)).IntVal)).IntVal
}

func Call_succC__func_func_func_interface____interface____interface____interface____func_interface____interface____interface____interface___3882721754(n_0_loop func(func(interface{}) interface{}, interface{}) interface{}, f_1_loop func(interface{}) interface{}, x_2_loop interface{}) interface{} {
var n_0 func(func(interface{}) interface{}, interface{}) interface{} = n_0_loop
_ = n_0
var f_1 func(interface{}) interface{} = f_1_loop
_ = f_1
var x_2 interface{} = x_2_loop
_ = x_2
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(f_1(gopurs_runtime.UnboxAny(gopurs_runtime.Any(n_0(f_1, x_2))))))
}

func Call_mulC(m_0_loop func(func(interface{}) interface{}, interface{}) interface{}, n_1_loop func(func(interface{}) interface{}, interface{}) interface{}, f_2_loop func(interface{}) interface{}, x_3_loop interface{}) interface{} {
var m_0 func(func(interface{}) interface{}, interface{}) interface{} = m_0_loop
_ = m_0
var n_1 func(func(interface{}) interface{}, interface{}) interface{} = n_1_loop
_ = n_1
var f_2 func(interface{}) interface{} = f_2_loop
_ = f_2
var x_3 interface{} = x_3_loop
_ = x_3
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(m_0(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(n_1(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(arg0, gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(arg1)))
}), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(f_2(gopurs_runtime.UnboxAny(arg0)))
})), gopurs_runtime.Any(inner_arg0)))
}, x_3)))
}

func Call_mulC__func_func_func_int64__int64__int64__int64__func_func_int64__int64__int64__int64__func_int64__int64__int64__int64_3478825786(m_0_loop func(func(int64) int64, int64) int64, n_1_loop func(func(int64) int64, int64) int64, f_2_loop func(int64) int64, x_3_loop int64) int64 {
var m_0 func(func(int64) int64, int64) int64 = m_0_loop
_ = m_0
var n_1 func(func(int64) int64, int64) int64 = n_1_loop
_ = n_1
var f_2 func(int64) int64 = f_2_loop
_ = f_2
var x_3 int64 = x_3_loop
_ = x_3
return gopurs_runtime.Int(m_0(func(inner_arg0 int64) int64 {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(n_1(func(inner_arg0 int64) int64 {
return gopurs_runtime.Apply(arg0, gopurs_runtime.Int(inner_arg0)).IntVal
}, arg1.IntVal))
}), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(f_2(arg0.IntVal))
})), gopurs_runtime.Int(inner_arg0)).IntVal
}, x_3)).IntVal
}

func Call_mulC__func_func_func_interface____interface____interface____interface____func_func_interface____interface____interface____interface____func_interface____interface____interface____interface___2338611642(m_0_loop func(func(interface{}) interface{}, interface{}) interface{}, n_1_loop func(func(interface{}) interface{}, interface{}) interface{}, f_2_loop func(interface{}) interface{}, x_3_loop interface{}) interface{} {
var m_0 func(func(interface{}) interface{}, interface{}) interface{} = m_0_loop
_ = m_0
var n_1 func(func(interface{}) interface{}, interface{}) interface{} = n_1_loop
_ = n_1
var f_2 func(interface{}) interface{} = f_2_loop
_ = f_2
var x_3 interface{} = x_3_loop
_ = x_3
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(m_0(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(n_1(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(arg0, gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(arg1)))
}), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(f_2(gopurs_runtime.UnboxAny(arg0)))
})), gopurs_runtime.Any(inner_arg0)))
}, x_3)))
}

func Call_fromInt(v_0_loop int64) gopurs_runtime.Value {
fromInt:
for {
if false { continue fromInt }
var v_0 int64 = v_0_loop
_ = v_0
var __t1 gopurs_runtime.Value
{
if (v_0) == (0) {
__t1 = Get_zeroC()
goto end_branch_1
} else {

}
}
{
__local_var_1_0 := Call_fromInt((v_0) - (1))
_ = __local_var_1_0
__t1 = gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, gopurs_runtime.Apply2(__local_var_1_0, f_2, x_3))
})
}
end_branch_1:
return __t1
}
}

func Call_c10(n_0_loop int64) gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
return Call_fromInt(n_0)
}

func Call_c100(n_0_loop int64) gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
__local_var_1_0 := Call_fromInt(n_0)
_ = __local_var_1_0
__local_var_2_1 := Call_fromInt(n_0)
_ = __local_var_2_1
return gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_1_0, gopurs_runtime.Apply(__local_var_2_1, f_3), x_4)
})
}

func Call_c10k(n_0_loop int64) gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
__local_var_1_0 := Call_c100(n_0)
_ = __local_var_1_0
__local_var_2_1 := Call_c100(n_0)
_ = __local_var_2_1
return gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_1_0, gopurs_runtime.Apply(__local_var_2_1, f_3), x_4)
})
}

func Call_c100k(n_0_loop int64) gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
__local_var_1_0 := Call_c10k(n_0)
_ = __local_var_1_0
__local_var_2_1 := Call_fromInt(n_0)
_ = __local_var_2_1
return gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_1_0, gopurs_runtime.Apply(__local_var_2_1, f_3), x_4)
})
}

func Call_addC(m_0_loop func(func(interface{}) interface{}, interface{}) interface{}, n_1_loop func(func(interface{}) interface{}, interface{}) interface{}, f_2_loop func(interface{}) interface{}, x_3_loop interface{}) interface{} {
var m_0 func(func(interface{}) interface{}, interface{}) interface{} = m_0_loop
_ = m_0
var n_1 func(func(interface{}) interface{}, interface{}) interface{} = n_1_loop
_ = n_1
var f_2 func(interface{}) interface{} = f_2_loop
_ = f_2
var x_3 interface{} = x_3_loop
_ = x_3
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(m_0(f_2, gopurs_runtime.UnboxAny(gopurs_runtime.Any(n_1(f_2, x_3))))))
}
