package Test_Church

import (
	pkg_Bench "gopurs/output/Bench"
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Effect "gopurs/output/Effect"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_zeroC gopurs_runtime.Value
var once_zeroC sync.Once
func Get_zeroC() gopurs_runtime.Value {
	once_zeroC.Do(func() {
		cache_zeroC = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zeroC(v_0_box, x_1_box)
})
	})
	return cache_zeroC
}

var cache_toInt gopurs_runtime.Value
var once_toInt sync.Once
func Get_toInt() gopurs_runtime.Value {
	once_toInt.Do(func() {
		cache_toInt = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_toInt(n_0_box))
})
	})
	return cache_toInt
}

var cache_succC gopurs_runtime.Value
var once_succC sync.Once
func Get_succC() gopurs_runtime.Value {
	once_succC.Do(func() {
		cache_succC = gopurs_runtime.Func3(func(n_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_succC(n_0_box, f_1_box, x_2_box)
})
	})
	return cache_succC
}

var cache_mulC gopurs_runtime.Value
var once_mulC sync.Once
func Get_mulC() gopurs_runtime.Value {
	once_mulC.Do(func() {
		cache_mulC = gopurs_runtime.Func4(func(m_0_box gopurs_runtime.Value, n_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value, x_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mulC(m_0_box, n_1_box, f_2_box, x_3_box)
})
	})
	return cache_mulC
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
		cache_describe = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("Church Numerals (100k Closure Applications):"))
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
return Call_addC(m_0_box, n_1_box, f_2_box, x_3_box)
})
	})
	return cache_addC
}

var cache_act gopurs_runtime.Value
var once_act sync.Once
func Get_act() gopurs_runtime.Value {
	once_act.Do(func() {
		cache_act = gopurs_runtime.Apply2(Get_bind__3550378017(), gopurs_runtime.Apply(pkg_Bench.Get_opaque(), gopurs_runtime.Int(10)), gopurs_runtime.Func(func(dummy_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Show.Get_showInt(), "show"), gopurs_runtime.Int(Call_mulC__1746928225(Call_c10k(dummy_0.IntVal), Call_fromInt(dummy_0.IntVal), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(gopurs_runtime.Apply2(Get_add__560788792(), gopurs_runtime.Int(x_1.IntVal), gopurs_runtime.Int(1)).IntVal)
}), gopurs_runtime.Int(0)).IntVal)).StrVal()))
}))
	})
	return cache_act
}

var cache_bind__2601835655 gopurs_runtime.Value
var once_bind__2601835655 sync.Once
func Get_bind__2601835655() gopurs_runtime.Value {
	once_bind__2601835655.Do(func() {
		cache_bind__2601835655 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__2601835655(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__2601835655
}

var cache_bind__3550378017 gopurs_runtime.Value
var once_bind__3550378017 sync.Once
func Get_bind__3550378017() gopurs_runtime.Value {
	once_bind__3550378017.Do(func() {
		cache_bind__3550378017 = gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind")
	})
	return cache_bind__3550378017
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

var cache_add__560788792 gopurs_runtime.Value
var once_add__560788792 sync.Once
func Get_add__560788792() gopurs_runtime.Value {
	once_add__560788792.Do(func() {
		cache_add__560788792 = pkg_Data_Semiring.Get_intAdd()
	})
	return cache_add__560788792
}

var cache_add__1614463960 gopurs_runtime.Value
var once_add__1614463960 sync.Once
func Get_add__1614463960() gopurs_runtime.Value {
	once_add__1614463960.Do(func() {
		cache_add__1614463960 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_add__1614463960(gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_add__1614463960
}

var cache_applyEffect__2014400020 gopurs_runtime.Value
var once_applyEffect__2014400020 sync.Once
func Get_applyEffect__2014400020() gopurs_runtime.Value {
	once_applyEffect__2014400020.Do(func() {
		cache_applyEffect__2014400020 = func() gopurs_runtime.Value {
Bind1_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_monadEffect(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_monadEffect(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect.Get_functorEffect()
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, f_2, gopurs_runtime.Func(func(f_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, a_3, gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_1_1.V1, gopurs_runtime.Apply(f_prime_4, a_prime_5))
}))
}))
})
}))
}()
	})
	return cache_applyEffect__2014400020
}

var cache_bindEffect__2113658466 gopurs_runtime.Value
var once_bindEffect__2113658466 sync.Once
func Get_bindEffect__2113658466() gopurs_runtime.Value {
	once_bindEffect__2113658466.Do(func() {
		cache_bindEffect__2113658466 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect.Get_applyEffect()
}), pkg_Effect.Get_bindE())
	})
	return cache_bindEffect__2113658466
}

var cache_functorEffect__3107547953 gopurs_runtime.Value
var once_functorEffect__3107547953 sync.Once
func Get_functorEffect__3107547953() gopurs_runtime.Value {
	once_functorEffect__3107547953.Do(func() {
		cache_functorEffect__3107547953 = func() gopurs_runtime.Value {
Apply0_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_0_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_0_0.V1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure"), f_1), a_2)
})
}))
}()
	})
	return cache_functorEffect__3107547953
}

var cache_mulC__1746928225 gopurs_runtime.Value
var once_mulC__1746928225 sync.Once
func Get_mulC__1746928225() gopurs_runtime.Value {
	once_mulC__1746928225.Do(func() {
		cache_mulC__1746928225 = gopurs_runtime.Func4(func(m_0_box gopurs_runtime.Value, n_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value, x_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mulC__1746928225(m_0_box, n_1_box, f_2_box, x_3_box)
})
	})
	return cache_mulC__1746928225
}

var cache_mulC__3596604257 gopurs_runtime.Value
var once_mulC__3596604257 sync.Once
func Get_mulC__3596604257() gopurs_runtime.Value {
	once_mulC__3596604257.Do(func() {
		cache_mulC__3596604257 = gopurs_runtime.Func4(func(m_0_box gopurs_runtime.Value, n_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value, x_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mulC__3596604257(m_0_box, n_1_box, f_2_box, x_3_box)
})
	})
	return cache_mulC__3596604257
}

var cache_succC__952275393 gopurs_runtime.Value
var once_succC__952275393 sync.Once
func Get_succC__952275393() gopurs_runtime.Value {
	once_succC__952275393.Do(func() {
		cache_succC__952275393 = gopurs_runtime.Func3(func(n_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_succC__952275393(n_0_box, f_1_box, x_2_box)
})
	})
	return cache_succC__952275393
}

var cache_succC__1461826241 gopurs_runtime.Value
var once_succC__1461826241 sync.Once
func Get_succC__1461826241() gopurs_runtime.Value {
	once_succC__1461826241.Do(func() {
		cache_succC__1461826241 = gopurs_runtime.Func3(func(n_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_succC__1461826241(n_0_box, f_1_box, x_2_box)
})
	})
	return cache_succC__1461826241
}

var cache_zeroC__4066693242 gopurs_runtime.Value
var once_zeroC__4066693242 sync.Once
func Get_zeroC__4066693242() gopurs_runtime.Value {
	once_zeroC__4066693242.Do(func() {
		cache_zeroC__4066693242 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_zeroC__4066693242(v_0_box, x_1_box.IntVal))
})
	})
	return cache_zeroC__4066693242
}

func Call_zeroC(v_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return x_1
}

func Call_toInt(n_0_loop gopurs_runtime.Value) int64 {
var n_0 gopurs_runtime.Value = n_0_loop
_ = n_0
return gopurs_runtime.Apply2(n_0, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(gopurs_runtime.Apply2(Get_add__560788792(), gopurs_runtime.Int(x_1.IntVal), gopurs_runtime.Int(1)).IntVal)
}), gopurs_runtime.Int(0)).IntVal
}

func Call_succC(n_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var n_0 gopurs_runtime.Value = n_0_loop
_ = n_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(f_1, gopurs_runtime.Apply2(n_0, f_1, x_2))
}

func Call_mulC(m_0_loop gopurs_runtime.Value, n_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value, x_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_fromInt(v_0_loop int64) gopurs_runtime.Value {
fromInt:
for {
if false { continue fromInt }
var v_0 int64 = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0) == (0) {
__t0 = Get_zeroC()
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Apply(Get_succC__952275393(), Call_fromInt(gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Int(v_0), gopurs_runtime.Int(1)).IntVal))
}
end_branch_0:
return __t0
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
return gopurs_runtime.Apply2(Get_mulC__1746928225(), Call_fromInt(n_0), Call_fromInt(n_0))
}

func Call_c10k(n_0_loop int64) gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
return gopurs_runtime.Apply2(Get_mulC__1746928225(), gopurs_runtime.Apply2(Get_mulC__1746928225(), Call_fromInt(n_0), Call_fromInt(n_0)), gopurs_runtime.Apply2(Get_mulC__1746928225(), Call_fromInt(n_0), Call_fromInt(n_0)))
}

func Call_c100k(n_0_loop int64) gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
return gopurs_runtime.Apply2(Get_mulC__1746928225(), Call_c10k(n_0), Call_fromInt(n_0))
}

func Call_addC(m_0_loop gopurs_runtime.Value, n_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value, x_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_bind__2601835655(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_sub__3675938712(dict_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_add__1614463960(dict_0_loop *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_mulC__1746928225(m_0_loop gopurs_runtime.Value, n_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value, x_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_mulC__3596604257(m_0_loop gopurs_runtime.Value, n_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value, x_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_succC__952275393(n_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var n_0 gopurs_runtime.Value = n_0_loop
_ = n_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(f_1, gopurs_runtime.Apply2(n_0, f_1, x_2))
}

func Call_succC__1461826241(n_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var n_0 gopurs_runtime.Value = n_0_loop
_ = n_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(f_1, gopurs_runtime.Apply2(n_0, f_1, x_2))
}

func Call_zeroC__4066693242(v_0_loop gopurs_runtime.Value, x_1_loop int64) int64 {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var x_1 int64 = x_1_loop
_ = x_1
return x_1
}


