package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Test_Church_logShow gopurs_runtime.Value
var once_Test_Church_logShow sync.Once
func Get_Test_Church_logShow() gopurs_runtime.Value {
	once_Test_Church_logShow.Do(func() {
		cache_Test_Church_logShow = gopurs_runtime.Apply(Get_Effect_Console_logShow(), Get_Data_Show_showInt())
	})
	return cache_Test_Church_logShow
}

var cache_Test_Church_zeroC gopurs_runtime.Value
var once_Test_Church_zeroC sync.Once
func Get_Test_Church_zeroC() gopurs_runtime.Value {
	once_Test_Church_zeroC.Do(func() {
		cache_Test_Church_zeroC = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_Church_zeroC(v_0_box, x_1_box)
})
	})
	return cache_Test_Church_zeroC
}

var cache_Test_Church_toInt gopurs_runtime.Value
var once_Test_Church_toInt sync.Once
func Get_Test_Church_toInt() gopurs_runtime.Value {
	once_Test_Church_toInt.Do(func() {
		cache_Test_Church_toInt = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Test_Church_toInt(n_0_box))
})
	})
	return cache_Test_Church_toInt
}

var cache_Test_Church_succC gopurs_runtime.Value
var once_Test_Church_succC sync.Once
func Get_Test_Church_succC() gopurs_runtime.Value {
	once_Test_Church_succC.Do(func() {
		cache_Test_Church_succC = gopurs_runtime.Func3(func(n_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_Church_succC(n_0_box, f_1_box, x_2_box)
})
	})
	return cache_Test_Church_succC
}

var cache_Test_Church_mulC gopurs_runtime.Value
var once_Test_Church_mulC sync.Once
func Get_Test_Church_mulC() gopurs_runtime.Value {
	once_Test_Church_mulC.Do(func() {
		cache_Test_Church_mulC = gopurs_runtime.Func4(func(m_0_box gopurs_runtime.Value, n_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value, x_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_Church_mulC(m_0_box, n_1_box, f_2_box, x_3_box)
})
	})
	return cache_Test_Church_mulC
}

var cache_Test_Church_fromInt gopurs_runtime.Value
var once_Test_Church_fromInt sync.Once
func Get_Test_Church_fromInt() gopurs_runtime.Value {
	once_Test_Church_fromInt.Do(func() {
		cache_Test_Church_fromInt = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_Church_fromInt(v_0_box.IntVal)
})
	})
	return cache_Test_Church_fromInt
}

var cache_Test_Church_describe gopurs_runtime.Value
var once_Test_Church_describe sync.Once
func Get_Test_Church_describe() gopurs_runtime.Value {
	once_Test_Church_describe.Do(func() {
		cache_Test_Church_describe = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Church Numerals (100k Closure Applications):"))
	})
	return cache_Test_Church_describe
}

var cache_Test_Church_c10 gopurs_runtime.Value
var once_Test_Church_c10 sync.Once
func Get_Test_Church_c10() gopurs_runtime.Value {
	once_Test_Church_c10.Do(func() {
		cache_Test_Church_c10 = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_Church_c10(n_0_box.IntVal)
})
	})
	return cache_Test_Church_c10
}

var cache_Test_Church_c100 gopurs_runtime.Value
var once_Test_Church_c100 sync.Once
func Get_Test_Church_c100() gopurs_runtime.Value {
	once_Test_Church_c100.Do(func() {
		cache_Test_Church_c100 = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_Church_c100(n_0_box.IntVal)
})
	})
	return cache_Test_Church_c100
}

var cache_Test_Church_c10k gopurs_runtime.Value
var once_Test_Church_c10k sync.Once
func Get_Test_Church_c10k() gopurs_runtime.Value {
	once_Test_Church_c10k.Do(func() {
		cache_Test_Church_c10k = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_Church_c10k(n_0_box.IntVal)
})
	})
	return cache_Test_Church_c10k
}

var cache_Test_Church_c100k gopurs_runtime.Value
var once_Test_Church_c100k sync.Once
func Get_Test_Church_c100k() gopurs_runtime.Value {
	once_Test_Church_c100k.Do(func() {
		cache_Test_Church_c100k = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_Church_c100k(n_0_box.IntVal)
})
	})
	return cache_Test_Church_c100k
}

var cache_Test_Church_addC gopurs_runtime.Value
var once_Test_Church_addC sync.Once
func Get_Test_Church_addC() gopurs_runtime.Value {
	once_Test_Church_addC.Do(func() {
		cache_Test_Church_addC = gopurs_runtime.Func4(func(m_0_box gopurs_runtime.Value, n_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value, x_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_Church_addC(m_0_box, n_1_box, f_2_box, x_3_box)
})
	})
	return cache_Test_Church_addC
}

var cache_Test_Church_act gopurs_runtime.Value
var once_Test_Church_act sync.Once
func Get_Test_Church_act() gopurs_runtime.Value {
	once_Test_Church_act.Do(func() {
		cache_Test_Church_act = func() gopurs_runtime.Value {
// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
__local_var_0_0 := gopurs_runtime.Apply(Get_Bench_opaque(), gopurs_runtime.Int(10))
_ = __local_var_0_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
_ = __local_var_1_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(gopurs_runtime.Apply2(Call_Test_Church_c100k(__local_var_1_1.IntVal), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((x_2.IntVal) + (1))
}), gopurs_runtime.Int(0)).IntVal)).StrVal())), gopurs_runtime.Value{})
})
}()
	})
	return cache_Test_Church_act
}

var cache_Test_Church_mulC__1746928225 gopurs_runtime.Value
var once_Test_Church_mulC__1746928225 sync.Once
func Get_Test_Church_mulC__1746928225() gopurs_runtime.Value {
	once_Test_Church_mulC__1746928225.Do(func() {
		cache_Test_Church_mulC__1746928225 = gopurs_runtime.Func4(func(m_0_box gopurs_runtime.Value, n_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value, x_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Test_Church_mulC__1746928225(m_0_box, n_1_box, f_2_box, x_3_box.IntVal))
})
	})
	return cache_Test_Church_mulC__1746928225
}

var cache_Test_Church_mulC__3596604257 gopurs_runtime.Value
var once_Test_Church_mulC__3596604257 sync.Once
func Get_Test_Church_mulC__3596604257() gopurs_runtime.Value {
	once_Test_Church_mulC__3596604257.Do(func() {
		cache_Test_Church_mulC__3596604257 = gopurs_runtime.Func4(func(m_0_box gopurs_runtime.Value, n_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value, x_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_Church_mulC__3596604257(m_0_box, n_1_box, f_2_box, x_3_box)
})
	})
	return cache_Test_Church_mulC__3596604257
}

var cache_Test_Church_succC__952275393 gopurs_runtime.Value
var once_Test_Church_succC__952275393 sync.Once
func Get_Test_Church_succC__952275393() gopurs_runtime.Value {
	once_Test_Church_succC__952275393.Do(func() {
		cache_Test_Church_succC__952275393 = gopurs_runtime.Func3(func(n_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Test_Church_succC__952275393(n_0_box, f_1_box, x_2_box.IntVal))
})
	})
	return cache_Test_Church_succC__952275393
}

var cache_Test_Church_succC__1461826241 gopurs_runtime.Value
var once_Test_Church_succC__1461826241 sync.Once
func Get_Test_Church_succC__1461826241() gopurs_runtime.Value {
	once_Test_Church_succC__1461826241.Do(func() {
		cache_Test_Church_succC__1461826241 = gopurs_runtime.Func3(func(n_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_Church_succC__1461826241(n_0_box, f_1_box, x_2_box)
})
	})
	return cache_Test_Church_succC__1461826241
}

var cache_Test_Church_zeroC__4066693242 gopurs_runtime.Value
var once_Test_Church_zeroC__4066693242 sync.Once
func Get_Test_Church_zeroC__4066693242() gopurs_runtime.Value {
	once_Test_Church_zeroC__4066693242.Do(func() {
		cache_Test_Church_zeroC__4066693242 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Test_Church_zeroC__4066693242(v_0_box, x_1_box.IntVal))
})
	})
	return cache_Test_Church_zeroC__4066693242
}

func Call_Test_Church_zeroC(v_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return x_1
}

func Call_Test_Church_toInt(n_0_loop gopurs_runtime.Value) int64 {
var n_0 gopurs_runtime.Value = n_0_loop
_ = n_0
return gopurs_runtime.Apply2(n_0, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((x_1.IntVal) + (1))
}), gopurs_runtime.Int(0)).IntVal
}

func Call_Test_Church_succC(n_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var n_0 gopurs_runtime.Value = n_0_loop
_ = n_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(f_1, gopurs_runtime.Apply2(n_0, f_1, x_2))
}

func Call_Test_Church_mulC(m_0_loop gopurs_runtime.Value, n_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value, x_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var m_0 gopurs_runtime.Value = m_0_loop
_ = m_0
var n_1 gopurs_runtime.Value = n_1_loop
_ = n_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
var x_3 gopurs_runtime.Value = x_3_loop
_ = x_3
return gopurs_runtime.Apply2(m_0, gopurs_runtime.Apply(n_1, f_2), x_3)
}

func Call_Test_Church_fromInt(v_0_loop int64) gopurs_runtime.Value {
fromInt:
for {
if false { continue fromInt }
var v_0 int64 = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0) == (0) {
__t0 = Get_Test_Church_zeroC()
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Apply(Get_Test_Church_succC__952275393(), Call_Test_Church_fromInt((v_0) - (1)))
}
end_branch_0:
return __t0
}
}

func Call_Test_Church_c10(n_0_loop int64) gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var __t0 gopurs_runtime.Value
{
if (n_0) == (0) {
__t0 = Get_Test_Church_zeroC()
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Apply(Get_Test_Church_succC__952275393(), Call_Test_Church_fromInt((n_0) - (1)))
}
end_branch_0:
return __t0
}

func Call_Test_Church_c100(n_0_loop int64) gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var __t2 gopurs_runtime.Value
{
if (n_0) == (0) {
__t2 = Get_Test_Church_zeroC()
goto end_branch_2
} else {

}
}
{
// TAST (Let): __local_var_1_0 -> int64
__local_var_1_0 := (n_0) - (1)
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (__local_var_1_0) == (0) {
__t1 = Get_Test_Church_zeroC()
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Apply(Get_Test_Church_succC__952275393(), Call_Test_Church_fromInt((__local_var_1_0) - (1)))
}
end_branch_1:
__t2 = gopurs_runtime.Apply(Get_Test_Church_succC__952275393(), __t1)
}
end_branch_2:
var __t5 gopurs_runtime.Value
{
if (n_0) == (0) {
__t5 = Get_Test_Church_zeroC()
goto end_branch_5
} else {

}
}
{
// TAST (Let): __local_var_1_3 -> int64
__local_var_1_3 := (n_0) - (1)
_ = __local_var_1_3
var __t4 gopurs_runtime.Value
{
if (__local_var_1_3) == (0) {
__t4 = Get_Test_Church_zeroC()
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Apply(Get_Test_Church_succC__952275393(), Call_Test_Church_fromInt((__local_var_1_3) - (1)))
}
end_branch_4:
__t5 = gopurs_runtime.Apply(Get_Test_Church_succC__952275393(), __t4)
}
end_branch_5:
return gopurs_runtime.Apply2(Get_Test_Church_mulC__1746928225(), __t2, __t5)
}

func Call_Test_Church_c10k(n_0_loop int64) gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
return gopurs_runtime.Apply2(Get_Test_Church_mulC__1746928225(), Call_Test_Church_c100(n_0), Call_Test_Church_c100(n_0))
}

func Call_Test_Church_c100k(n_0_loop int64) gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var __t2 gopurs_runtime.Value
{
if (n_0) == (0) {
__t2 = Get_Test_Church_zeroC()
goto end_branch_2
} else {

}
}
{
// TAST (Let): __local_var_1_0 -> int64
__local_var_1_0 := (n_0) - (1)
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (__local_var_1_0) == (0) {
__t1 = Get_Test_Church_zeroC()
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Apply(Get_Test_Church_succC__952275393(), Call_Test_Church_fromInt((__local_var_1_0) - (1)))
}
end_branch_1:
__t2 = gopurs_runtime.Apply(Get_Test_Church_succC__952275393(), __t1)
}
end_branch_2:
return gopurs_runtime.Apply2(Get_Test_Church_mulC__1746928225(), gopurs_runtime.Apply2(Get_Test_Church_mulC__1746928225(), Call_Test_Church_c100(n_0), Call_Test_Church_c100(n_0)), __t2)
}

func Call_Test_Church_addC(m_0_loop gopurs_runtime.Value, n_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value, x_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var m_0 gopurs_runtime.Value = m_0_loop
_ = m_0
var n_1 gopurs_runtime.Value = n_1_loop
_ = n_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
var x_3 gopurs_runtime.Value = x_3_loop
_ = x_3
return gopurs_runtime.Apply2(m_0, f_2, gopurs_runtime.Apply2(n_1, f_2, x_3))
}

func Call_Test_Church_mulC__1746928225(m_0_loop gopurs_runtime.Value, n_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value, x_3_loop int64) int64 {
var m_0 gopurs_runtime.Value = m_0_loop
_ = m_0
var n_1 gopurs_runtime.Value = n_1_loop
_ = n_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
var x_3 int64 = x_3_loop
_ = x_3
return gopurs_runtime.Apply2(m_0, gopurs_runtime.Apply(n_1, f_2), gopurs_runtime.Int(x_3)).IntVal
}

func Call_Test_Church_mulC__3596604257(m_0_loop gopurs_runtime.Value, n_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value, x_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var m_0 gopurs_runtime.Value = m_0_loop
_ = m_0
var n_1 gopurs_runtime.Value = n_1_loop
_ = n_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
var x_3 gopurs_runtime.Value = x_3_loop
_ = x_3
return gopurs_runtime.Apply2(m_0, gopurs_runtime.Apply(n_1, f_2), x_3)
}

func Call_Test_Church_succC__952275393(n_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, x_2_loop int64) int64 {
var n_0 gopurs_runtime.Value = n_0_loop
_ = n_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var x_2 int64 = x_2_loop
_ = x_2
return gopurs_runtime.Apply(f_1, gopurs_runtime.Int(gopurs_runtime.Apply2(n_0, f_1, gopurs_runtime.Int(x_2)).IntVal)).IntVal
}

func Call_Test_Church_succC__1461826241(n_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var n_0 gopurs_runtime.Value = n_0_loop
_ = n_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(f_1, gopurs_runtime.Apply2(n_0, f_1, x_2))
}

func Call_Test_Church_zeroC__4066693242(v_0_loop gopurs_runtime.Value, x_1_loop int64) int64 {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var x_1 int64 = x_1_loop
_ = x_1
return x_1
}


