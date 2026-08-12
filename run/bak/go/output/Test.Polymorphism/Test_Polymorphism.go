package Test_Polymorphism

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
	unsafe "unsafe"
)

var cache_mempty_ gopurs_runtime.Value
var once_mempty_ sync.Once
func Get_mempty_() gopurs_runtime.Value {
	once_mempty_.Do(func() {
		cache_mempty_ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mempty_(dict_0_box)
})
	})
	return cache_mempty_
}

var cache_mappend_ gopurs_runtime.Value
var once_mappend_ sync.Once
func Get_mappend_() gopurs_runtime.Value {
	once_mappend_.Do(func() {
		cache_mappend_ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mappend_(gopurs_runtime.CoerceToStruct[Constructor_Monoidish[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_mappend_
}

var cache_polyLoop gopurs_runtime.Value
var once_polyLoop sync.Once
func Get_polyLoop() gopurs_runtime.Value {
	once_polyLoop.Do(func() {
		cache_polyLoop = gopurs_runtime.Func3(func(dictMonoidish_0_box gopurs_runtime.Value, n_init_1_box gopurs_runtime.Value, acc_init_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_polyLoop(gopurs_runtime.CoerceToStruct[Constructor_Monoidish[gopurs_runtime.Value]](dictMonoidish_0_box), n_init_1_box.IntVal, acc_init_2_box)
})
	})
	return cache_polyLoop
}

var cache_intMonoidish gopurs_runtime.Value
var once_intMonoidish sync.Once
func Get_intMonoidish() gopurs_runtime.Value {
	once_intMonoidish.Do(func() {
		cache_intMonoidish = gopurs_runtime.RecordDict2("mappend_", "mempty_", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_add__560788792(gopurs_runtime.Int(x_0.IntVal), gopurs_runtime.Int(y_1.IntVal)).IntVal)
})
}), gopurs_runtime.Int(1))
	})
	return cache_intMonoidish
}

var cache_describe gopurs_runtime.Value
var once_describe sync.Once
func Get_describe() gopurs_runtime.Value {
	once_describe.Do(func() {
		cache_describe = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("Polymorphism (10M Type Class Dict Lookups):"))
	})
	return cache_describe
}

var cache_act gopurs_runtime.Value
var once_act sync.Once
func Get_act() gopurs_runtime.Value {
	once_act.Do(func() {
		cache_act = Call_bind__3550378017(gopurs_runtime.Apply(pkg_Bench.Get_opaque(), gopurs_runtime.Int(10000000)), gopurs_runtime.Func(func(dummy_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Show.Get_showInt(), "show"), gopurs_runtime.Int(Call_polyLoop__1533381815(dummy_0.IntVal, gopurs_runtime.Int(0)).IntVal)).StrVal()))
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

var cache_sub__2927892844 gopurs_runtime.Value
var once_sub__2927892844 sync.Once
func Get_sub__2927892844() gopurs_runtime.Value {
	once_sub__2927892844.Do(func() {
		cache_sub__2927892844 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__2927892844(gopurs_runtime.CoerceToStruct[pkg_Data_Ring.Constructor_Ring[int64]](dict_0_box))
})
	})
	return cache_sub__2927892844
}

var cache_sub__1124926121 gopurs_runtime.Value
var once_sub__1124926121 sync.Once
func Get_sub__1124926121() gopurs_runtime.Value {
	once_sub__1124926121.Do(func() {
		cache_sub__1124926121 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__1124926121(gopurs_runtime.CoerceToStruct[pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sub__1124926121
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

var cache_mappend___1124926121 gopurs_runtime.Value
var once_mappend___1124926121 sync.Once
func Get_mappend___1124926121() gopurs_runtime.Value {
	once_mappend___1124926121.Do(func() {
		cache_mappend___1124926121 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mappend___1124926121(gopurs_runtime.CoerceToStruct[Constructor_Monoidish[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_mappend___1124926121
}

var cache_mappend___3566619927 gopurs_runtime.Value
var once_mappend___3566619927 sync.Once
func Get_mappend___3566619927() gopurs_runtime.Value {
	once_mappend___3566619927.Do(func() {
		cache_mappend___3566619927 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mappend___3566619927(gopurs_runtime.CoerceToStruct[Constructor_Monoidish[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_mappend___3566619927
}

var cache_mempty___1556010056 gopurs_runtime.Value
var once_mempty___1556010056 sync.Once
func Get_mempty___1556010056() gopurs_runtime.Value {
	once_mempty___1556010056.Do(func() {
		cache_mempty___1556010056 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mempty___1556010056(dict_0_box)
})
	})
	return cache_mempty___1556010056
}

var cache_mempty___1540866998 gopurs_runtime.Value
var once_mempty___1540866998 sync.Once
func Get_mempty___1540866998() gopurs_runtime.Value {
	once_mempty___1540866998.Do(func() {
		cache_mempty___1540866998 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mempty___1540866998(dict_0_box)
})
	})
	return cache_mempty___1540866998
}

var cache_polyLoop__1533381815 gopurs_runtime.Value
var once_polyLoop__1533381815 sync.Once
func Get_polyLoop__1533381815() gopurs_runtime.Value {
	once_polyLoop__1533381815.Do(func() {
		cache_polyLoop__1533381815 = gopurs_runtime.Func2(func(n_init_0_box gopurs_runtime.Value, acc_init_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_polyLoop__1533381815(n_init_0_box.IntVal, acc_init_1_box)
})
	})
	return cache_polyLoop__1533381815
}

var cache_polyLoop__2675791634 gopurs_runtime.Value
var once_polyLoop__2675791634 sync.Once
func Get_polyLoop__2675791634() gopurs_runtime.Value {
	once_polyLoop__2675791634.Do(func() {
		cache_polyLoop__2675791634 = gopurs_runtime.Func3(func(dictMonoidish_0_box gopurs_runtime.Value, n_init_1_box gopurs_runtime.Value, acc_init_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_polyLoop__2675791634(gopurs_runtime.CoerceToStruct[Constructor_Monoidish[gopurs_runtime.Value]](dictMonoidish_0_box), n_init_1_box.IntVal, acc_init_2_box)
})
	})
	return cache_polyLoop__2675791634
}

type Constructor_Monoidish[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 T_a
}


func init() {
	gopurs_runtime.StructGetters[459160245] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Monoidish[gopurs_runtime.Value])(ptr)
		switch key {
		case "mappend_": return c.V0
		case "mempty_": return c.V1
		default: panic("Key not found in dictionary Constructor_Monoidish: " + key)
		}
	}
}


func Call_mempty_(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "mempty_")
}

func Call_mappend_(dict_0_loop *Constructor_Monoidish[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Monoidish[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_polyLoop(dictMonoidish_0_loop *Constructor_Monoidish[gopurs_runtime.Value], n_init_1_loop int64, acc_init_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoidish_0 *Constructor_Monoidish[gopurs_runtime.Value] = dictMonoidish_0_loop
_ = dictMonoidish_0
var n_init_1 int64 = n_init_1_loop
_ = n_init_1
var acc_init_2 gopurs_runtime.Value = acc_init_2_loop
_ = acc_init_2
var go__go_3_0_0 gopurs_runtime.Value
go__go_3_0_0 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop int64 = v_4_loop_val.IntVal
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_0_0:
for {
if false { continue go__go_3_0_0 }
var v_4 int64 = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t1 gopurs_runtime.Value
{
if (v_4) == (0) {
__t1 = v1_5
goto end_branch_1
} else {

}
}
{
v_4_loop = Call_sub__1043827704(gopurs_runtime.Int(v_4), gopurs_runtime.Int(1)).IntVal
v1_5_loop = gopurs_runtime.Apply2(dictMonoidish_0.V0, v1_5, dictMonoidish_0.V1)
continue go__go_3_0_0
__t1 = gopurs_runtime.Value{}
}
end_branch_1:
return __t1
}
}()
})
})
return gopurs_runtime.Apply2(go__go_3_0_0, gopurs_runtime.Int(n_init_1), acc_init_2)
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

func Call_sub__2927892844(dict_0_loop *pkg_Data_Ring.Constructor_Ring[int64]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ring.Constructor_Ring[int64] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_sub__1124926121(dict_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
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

func Call_mappend___1124926121(dict_0_loop *Constructor_Monoidish[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Monoidish[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_mappend___3566619927(dict_0_loop *Constructor_Monoidish[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Monoidish[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_mempty___1556010056(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "mempty_")
}

func Call_mempty___1540866998(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "mempty_")
}

func Call_polyLoop__1533381815(n_init_0_loop int64, acc_init_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var n_init_0 int64 = n_init_0_loop
_ = n_init_0
var acc_init_1 gopurs_runtime.Value = acc_init_1_loop
_ = acc_init_1
var go__go_2_0_1 gopurs_runtime.Value
go__go_2_0_1 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_3_loop int64 = v_3_loop_val.IntVal
var v1_4_loop gopurs_runtime.Value = v1_4_loop_val
go__go_2_0_1:
for {
if false { continue go__go_2_0_1 }
var v_3 int64 = v_3_loop
_ = v_3
var v1_4 gopurs_runtime.Value = v1_4_loop
_ = v1_4
var __t1 gopurs_runtime.Value
{
if (v_3) == (0) {
__t1 = v1_4
goto end_branch_1
} else {

}
}
{
v_3_loop = (v_3) - (1)
v1_4_loop = gopurs_runtime.Int((v1_4.IntVal) + (1))
continue go__go_2_0_1
__t1 = gopurs_runtime.Value{}
}
end_branch_1:
return __t1
}
}()
})
})
return gopurs_runtime.Apply2(go__go_2_0_1, gopurs_runtime.Int(n_init_0), acc_init_1)
}

func Call_polyLoop__2675791634(dictMonoidish_0_loop *Constructor_Monoidish[gopurs_runtime.Value], n_init_1_loop int64, acc_init_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoidish_0 *Constructor_Monoidish[gopurs_runtime.Value] = dictMonoidish_0_loop
_ = dictMonoidish_0
var n_init_1 int64 = n_init_1_loop
_ = n_init_1
var acc_init_2 gopurs_runtime.Value = acc_init_2_loop
_ = acc_init_2
var go__go_3_0_2 gopurs_runtime.Value
go__go_3_0_2 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop int64 = v_4_loop_val.IntVal
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_0_2:
for {
if false { continue go__go_3_0_2 }
var v_4 int64 = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t1 gopurs_runtime.Value
{
if (v_4) == (0) {
__t1 = v1_5
goto end_branch_1
} else {

}
}
{
v_4_loop = (v_4) - (1)
v1_5_loop = gopurs_runtime.Apply2(dictMonoidish_0.V0, v1_5, dictMonoidish_0.V1)
continue go__go_3_0_2
__t1 = gopurs_runtime.Value{}
}
end_branch_1:
return __t1
}
}()
})
})
return gopurs_runtime.Apply2(go__go_3_0_2, gopurs_runtime.Int(n_init_1), acc_init_2)
}


