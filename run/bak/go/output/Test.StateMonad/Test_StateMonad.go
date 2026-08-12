package Test_StateMonad

import (
	pkg_Bench "gopurs/output/Bench"
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Effect "gopurs/output/Effect"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_State gopurs_runtime.Value
var once_State sync.Once
func Get_State() gopurs_runtime.Value {
	once_State.Do(func() {
		cache_State = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_State(x_0_box)
})
	})
	return cache_State
}

var cache_runState gopurs_runtime.Value
var once_runState sync.Once
func Get_runState() gopurs_runtime.Value {
	once_runState.Do(func() {
		cache_runState = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runState(v_0_box, s_1_box)
})
	})
	return cache_runState
}

var cache_put gopurs_runtime.Value
var once_put sync.Once
func Get_put() gopurs_runtime.Value {
	once_put.Do(func() {
		cache_put = gopurs_runtime.Func2(func(s_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_put(s_0_box, v_1_box)
})
	})
	return cache_put
}

var cache_pureState gopurs_runtime.Value
var once_pureState sync.Once
func Get_pureState() gopurs_runtime.Value {
	once_pureState.Do(func() {
		cache_pureState = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pureState(a_0_box, s_1_box)
})
	})
	return cache_pureState
}

var cache_get gopurs_runtime.Value
var once_get sync.Once
func Get_get() gopurs_runtime.Value {
	once_get.Do(func() {
		cache_get = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_get(s_0_box)
})
	})
	return cache_get
}

var cache_describe gopurs_runtime.Value
var once_describe sync.Once
func Get_describe() gopurs_runtime.Value {
	once_describe.Do(func() {
		cache_describe = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("State Monad (1.2k Binds, 60 Stack Depth):"))
	})
	return cache_describe
}

var cache_bindState gopurs_runtime.Value
var once_bindState sync.Once
func Get_bindState() gopurs_runtime.Value {
	once_bindState.Do(func() {
		cache_bindState = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindState(v_0_box, g_1_box, s_2_box)
})
	})
	return cache_bindState
}

var cache_modify gopurs_runtime.Value
var once_modify sync.Once
func Get_modify() gopurs_runtime.Value {
	once_modify.Do(func() {
		cache_modify = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_modify(f_0_box, s_1_box)
})
	})
	return cache_modify
}

var cache_chainModifications gopurs_runtime.Value
var once_chainModifications sync.Once
func Get_chainModifications() gopurs_runtime.Value {
	once_chainModifications.Do(func() {
		cache_chainModifications = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_chainModifications(v_0_box.IntVal)
})
	})
	return cache_chainModifications
}

var cache_runManyTimes gopurs_runtime.Value
var once_runManyTimes sync.Once
func Get_runManyTimes() gopurs_runtime.Value {
	once_runManyTimes.Do(func() {
		cache_runManyTimes = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_runManyTimes(v_0_box.IntVal, v1_1_box.IntVal))
})
	})
	return cache_runManyTimes
}

var cache_act gopurs_runtime.Value
var once_act sync.Once
func Get_act() gopurs_runtime.Value {
	once_act.Do(func() {
		cache_act = Call_bind__3550378017(gopurs_runtime.Apply(pkg_Bench.Get_opaque(), gopurs_runtime.Int(20)), gopurs_runtime.Func(func(dummy_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Show.Get_showInt(), "show"), gopurs_runtime.Int(gopurs_runtime.Int(Call_runManyTimes(dummy_0.IntVal, 0)).IntVal)).StrVal()))
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
		cache_bind__3550378017 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__3550378017(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_bind__3550378017
}

var cache_sub__1043827704 gopurs_runtime.Value
var once_sub__1043827704 sync.Once
func Get_sub__1043827704() gopurs_runtime.Value {
	once_sub__1043827704.Do(func() {
		cache_sub__1043827704 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__1043827704(__eta0_0_box, __eta1_1_box)
})
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
		cache_add__560788792 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_add__560788792(__eta0_0_box, __eta1_1_box)
})
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

var cache_bindState__567439955 gopurs_runtime.Value
var once_bindState__567439955 sync.Once
func Get_bindState__567439955() gopurs_runtime.Value {
	once_bindState__567439955.Do(func() {
		cache_bindState__567439955 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindState__567439955(v_0_box, g_1_box, s_2_box)
})
	})
	return cache_bindState__567439955
}

var cache_bindState__2171045075 gopurs_runtime.Value
var once_bindState__2171045075 sync.Once
func Get_bindState__2171045075() gopurs_runtime.Value {
	once_bindState__2171045075.Do(func() {
		cache_bindState__2171045075 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindState__2171045075(v_0_box, g_1_box, s_2_box)
})
	})
	return cache_bindState__2171045075
}

var cache_bindState__3267751411 gopurs_runtime.Value
var once_bindState__3267751411 sync.Once
func Get_bindState__3267751411() gopurs_runtime.Value {
	once_bindState__3267751411.Do(func() {
		cache_bindState__3267751411 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindState__3267751411(v_0_box, g_1_box, s_2_box)
})
	})
	return cache_bindState__3267751411
}

var cache_get__676984528 gopurs_runtime.Value
var once_get__676984528 sync.Once
func Get_get__676984528() gopurs_runtime.Value {
	once_get__676984528.Do(func() {
		cache_get__676984528 = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_get__676984528(s_0_box)
})
	})
	return cache_get__676984528
}

var cache_modify__1175978184 gopurs_runtime.Value
var once_modify__1175978184 sync.Once
func Get_modify__1175978184() gopurs_runtime.Value {
	once_modify__1175978184.Do(func() {
		cache_modify__1175978184 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_modify__1175978184(f_0_box, s_1_box)
})
	})
	return cache_modify__1175978184
}

var cache_modify__3050914184 gopurs_runtime.Value
var once_modify__3050914184 sync.Once
func Get_modify__3050914184() gopurs_runtime.Value {
	once_modify__3050914184.Do(func() {
		cache_modify__3050914184 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_modify__3050914184(f_0_box, s_1_box)
})
	})
	return cache_modify__3050914184
}

var cache_pureState__608762702 gopurs_runtime.Value
var once_pureState__608762702 sync.Once
func Get_pureState__608762702() gopurs_runtime.Value {
	once_pureState__608762702.Do(func() {
		cache_pureState__608762702 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pureState__608762702(a_0_box, s_1_box)
})
	})
	return cache_pureState__608762702
}

var cache_pureState__1329830318 gopurs_runtime.Value
var once_pureState__1329830318 sync.Once
func Get_pureState__1329830318() gopurs_runtime.Value {
	once_pureState__1329830318.Do(func() {
		cache_pureState__1329830318 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pureState__1329830318(a_0_box, s_1_box)
})
	})
	return cache_pureState__1329830318
}

var cache_put__3685210848 gopurs_runtime.Value
var once_put__3685210848 sync.Once
func Get_put__3685210848() gopurs_runtime.Value {
	once_put__3685210848.Do(func() {
		cache_put__3685210848 = gopurs_runtime.Func2(func(s_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_put__3685210848(s_0_box, v_1_box)
})
	})
	return cache_put__3685210848
}

var cache_runState__2373419117 gopurs_runtime.Value
var once_runState__2373419117 sync.Once
func Get_runState__2373419117() gopurs_runtime.Value {
	once_runState__2373419117.Do(func() {
		cache_runState__2373419117 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runState__2373419117(v_0_box, s_1_box)
})
	})
	return cache_runState__2373419117
}

var cache_runState__3059282509 gopurs_runtime.Value
var once_runState__3059282509 sync.Once
func Get_runState__3059282509() gopurs_runtime.Value {
	once_runState__3059282509.Do(func() {
		cache_runState__3059282509 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runState__3059282509(v_0_box, s_1_box)
})
	})
	return cache_runState__3059282509
}

func Call_State(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_runState(v_0_loop gopurs_runtime.Value, s_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var s_1 gopurs_runtime.Value = s_1_loop
_ = s_1
return gopurs_runtime.Apply(v_0, s_1)
}

func Call_put(s_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var s_0 gopurs_runtime.Value = s_0_loop
_ = s_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.RecordDict2("state", "val", s_0, pkg_Data_Unit.Get_unit())
}

func Call_pureState(a_0_loop gopurs_runtime.Value, s_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var s_1 gopurs_runtime.Value = s_1_loop
_ = s_1
return gopurs_runtime.RecordDict2("state", "val", s_1, a_0)
}

func Call_get(s_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var s_0 gopurs_runtime.Value = s_0_loop
_ = s_0
return gopurs_runtime.RecordDict2("state", "val", s_0, s_0)
}

func Call_bindState(v_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, s_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var g_1 gopurs_runtime.Value = g_1_loop
_ = g_1
var s_2 gopurs_runtime.Value = s_2_loop
_ = s_2
r1_3_0 := gopurs_runtime.Apply(v_0, s_2)
_ = r1_3_0
return gopurs_runtime.Apply2(g_1, gopurs_runtime.RecordGet(r1_3_0, "val"), gopurs_runtime.RecordGet(r1_3_0, "state"))
}

func Call_modify(f_0_loop gopurs_runtime.Value, s_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var s_1 gopurs_runtime.Value = s_1_loop
_ = s_1
return gopurs_runtime.RecordDict2("state", "val", gopurs_runtime.Apply(f_0, s_1), pkg_Data_Unit.Get_unit())
}

func Call_chainModifications(v_0_loop int64) gopurs_runtime.Value {
chainModifications:
for {
if false { continue chainModifications }
var v_0 int64 = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0) == (0) {
__t0 = gopurs_runtime.Apply(Get_pureState__608762702(), pkg_Data_Unit.Get_unit())
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Apply2(Get_bindState__567439955(), gopurs_runtime.Apply(Get_modify__1175978184(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_add__560788792(gopurs_runtime.Int(x_1.IntVal), gopurs_runtime.Int(1)).IntVal)
})), gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_chainModifications(Call_sub__1043827704(gopurs_runtime.Int(v_0), gopurs_runtime.Int(1)).IntVal)
}))
}
end_branch_0:
return __t0
}
}

func Call_runManyTimes(v_0_loop int64, v1_1_loop int64) int64 {
runManyTimes:
for {
if false { continue runManyTimes }
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 int64 = v1_1_loop
_ = v1_1
var __t0 int64
{
if (v_0) == (0) {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
v_0_loop = Call_sub__1043827704(gopurs_runtime.Int(v_0), gopurs_runtime.Int(1)).IntVal
v1_1_loop = Call_add__560788792(gopurs_runtime.Int(v1_1), gopurs_runtime.Int(gopurs_runtime.RecordGet(Call_runState__2373419117(Call_chainModifications(60), gopurs_runtime.Int(0)), "state").IntVal)).IntVal
continue runManyTimes
__t0 = gopurs_runtime.Value{}.IntVal
}
end_branch_0:
return __t0
}
}

func Call_bind__2601835655(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__3550378017(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Value{})
_ = __local_var_2_0
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta1_1, __local_var_2_0), gopurs_runtime.Value{})
})
}

func Call_sub__1043827704(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Int((__eta0_0.IntVal) - (__eta1_1.IntVal))
}

func Call_sub__3675938712(dict_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_add__560788792(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Int((__eta0_0.IntVal) + (__eta1_1.IntVal))
}

func Call_add__1614463960(dict_0_loop *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_bindState__567439955(v_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, s_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var g_1 gopurs_runtime.Value = g_1_loop
_ = g_1
var s_2 gopurs_runtime.Value = s_2_loop
_ = s_2
r1_3_0 := gopurs_runtime.Apply(v_0, s_2)
_ = r1_3_0
return gopurs_runtime.Apply2(g_1, gopurs_runtime.RecordGet(r1_3_0, "val"), gopurs_runtime.RecordGet(r1_3_0, "state"))
}

func Call_bindState__2171045075(v_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, s_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var g_1 gopurs_runtime.Value = g_1_loop
_ = g_1
var s_2 gopurs_runtime.Value = s_2_loop
_ = s_2
r1_3_0 := gopurs_runtime.Apply(v_0, s_2)
_ = r1_3_0
return gopurs_runtime.Apply2(g_1, gopurs_runtime.RecordGet(r1_3_0, "val"), gopurs_runtime.RecordGet(r1_3_0, "state"))
}

func Call_bindState__3267751411(v_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, s_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var g_1 gopurs_runtime.Value = g_1_loop
_ = g_1
var s_2 gopurs_runtime.Value = s_2_loop
_ = s_2
r1_3_0 := gopurs_runtime.Apply(v_0, s_2)
_ = r1_3_0
return gopurs_runtime.Apply2(g_1, gopurs_runtime.RecordGet(r1_3_0, "val"), gopurs_runtime.RecordGet(r1_3_0, "state"))
}

func Call_get__676984528(s_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var s_0 gopurs_runtime.Value = s_0_loop
_ = s_0
return gopurs_runtime.RecordDict2("state", "val", s_0, s_0)
}

func Call_modify__1175978184(f_0_loop gopurs_runtime.Value, s_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var s_1 gopurs_runtime.Value = s_1_loop
_ = s_1
return gopurs_runtime.RecordDict2("state", "val", gopurs_runtime.Apply(f_0, s_1), pkg_Data_Unit.Get_unit())
}

func Call_modify__3050914184(f_0_loop gopurs_runtime.Value, s_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var s_1 gopurs_runtime.Value = s_1_loop
_ = s_1
return gopurs_runtime.RecordDict2("state", "val", gopurs_runtime.Apply(f_0, s_1), pkg_Data_Unit.Get_unit())
}

func Call_pureState__608762702(a_0_loop gopurs_runtime.Value, s_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var s_1 gopurs_runtime.Value = s_1_loop
_ = s_1
return gopurs_runtime.RecordDict2("state", "val", s_1, a_0)
}

func Call_pureState__1329830318(a_0_loop gopurs_runtime.Value, s_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var s_1 gopurs_runtime.Value = s_1_loop
_ = s_1
return gopurs_runtime.RecordDict2("state", "val", s_1, a_0)
}

func Call_put__3685210848(s_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var s_0 gopurs_runtime.Value = s_0_loop
_ = s_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.RecordDict2("state", "val", s_0, pkg_Data_Unit.Get_unit())
}

func Call_runState__2373419117(v_0_loop gopurs_runtime.Value, s_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var s_1 gopurs_runtime.Value = s_1_loop
_ = s_1
return gopurs_runtime.Apply(v_0, s_1)
}

func Call_runState__3059282509(v_0_loop gopurs_runtime.Value, s_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var s_1 gopurs_runtime.Value = s_1_loop
_ = s_1
return gopurs_runtime.Apply(v_0, s_1)
}


