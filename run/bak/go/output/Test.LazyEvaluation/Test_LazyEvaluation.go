package Test_LazyEvaluation

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

var cache_Lazy gopurs_runtime.Value
var once_Lazy sync.Once
func Get_Lazy() gopurs_runtime.Value {
	once_Lazy.Do(func() {
		cache_Lazy = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Lazy(x_0_box)
})
	})
	return cache_Lazy
}

var cache_force gopurs_runtime.Value
var once_force sync.Once
func Get_force() gopurs_runtime.Value {
	once_force.Do(func() {
		cache_force = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_force(v_0_box)
})
	})
	return cache_force
}

var cache_describe gopurs_runtime.Value
var once_describe sync.Once
func Get_describe() gopurs_runtime.Value {
	once_describe.Do(func() {
		cache_describe = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("Lazy Evaluation (1M Thunks Forced, 1k Depth):"))
	})
	return cache_describe
}

var cache_go__defer gopurs_runtime.Value
var once_go__defer sync.Once
func Get_go__defer() gopurs_runtime.Value {
	once_go__defer.Do(func() {
		cache_go__defer = Get_Lazy()
	})
	return cache_go__defer
}

var cache_buildThunks gopurs_runtime.Value
var once_buildThunks sync.Once
func Get_buildThunks() gopurs_runtime.Value {
	once_buildThunks.Do(func() {
		cache_buildThunks = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_buildThunks(v_0_box.IntVal, v1_1_box)
})
	})
	return cache_buildThunks
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
		cache_act = gopurs_runtime.Apply2(Get_bind__3550378017(), gopurs_runtime.Apply(pkg_Bench.Get_opaque(), gopurs_runtime.Int(1000)), gopurs_runtime.Func(func(dummy_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Show.Get_showInt(), "show"), gopurs_runtime.Int(Call_runManyTimes(dummy_0.IntVal, 0))))
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

var cache_defer__3520065601 gopurs_runtime.Value
var once_defer__3520065601 sync.Once
func Get_defer__3520065601() gopurs_runtime.Value {
	once_defer__3520065601.Do(func() {
		cache_defer__3520065601 = Get_Lazy()
	})
	return cache_defer__3520065601
}

var cache_defer__3363737377 gopurs_runtime.Value
var once_defer__3363737377 sync.Once
func Get_defer__3363737377() gopurs_runtime.Value {
	once_defer__3363737377.Do(func() {
		cache_defer__3363737377 = Get_Lazy()
	})
	return cache_defer__3363737377
}

var cache_force__3902501304 gopurs_runtime.Value
var once_force__3902501304 sync.Once
func Get_force__3902501304() gopurs_runtime.Value {
	once_force__3902501304.Do(func() {
		cache_force__3902501304 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_force__3902501304(v_0_box)
})
	})
	return cache_force__3902501304
}

var cache_force__721037880 gopurs_runtime.Value
var once_force__721037880 sync.Once
func Get_force__721037880() gopurs_runtime.Value {
	once_force__721037880.Do(func() {
		cache_force__721037880 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_force__721037880(v_0_box)
})
	})
	return cache_force__721037880
}

func Call_Lazy(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_force(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Apply(v_0, pkg_Data_Unit.Get_unit())
}

func Call_buildThunks(v_0_loop int64, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
buildThunks:
for {
if false { continue buildThunks }
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (v_0) == (0) {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
v_0_loop = gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Int(v_0), gopurs_runtime.Int(1)).IntVal
v1_1_loop = gopurs_runtime.Apply(Get_defer__3520065601(), gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(gopurs_runtime.Apply2(Get_add__560788792(), Call_force__3902501304(v1_1), gopurs_runtime.Int(1)).IntVal)
}))
continue buildThunks
__t0 = gopurs_runtime.Value{}
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
v_0_loop = gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Int(v_0), gopurs_runtime.Int(1)).IntVal
v1_1_loop = gopurs_runtime.Apply2(Get_add__560788792(), gopurs_runtime.Int(v1_1), Call_force__3902501304(Call_buildThunks(1000, gopurs_runtime.Apply(Get_defer__3520065601(), gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(0)
}))))).IntVal
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

func Call_force__3902501304(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Apply(v_0, pkg_Data_Unit.Get_unit())
}

func Call_force__721037880(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Apply(v_0, pkg_Data_Unit.Get_unit())
}


