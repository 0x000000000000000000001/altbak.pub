package Test_Records

import (
	pkg_Bench "gopurs/output/Bench"
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Effect "gopurs/output/Effect"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_updateRec gopurs_runtime.Value
var once_updateRec sync.Once
func Get_updateRec() gopurs_runtime.Value {
	once_updateRec.Do(func() {
		cache_updateRec = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_updateRec(v_0_box.IntVal, v1_1_box)
})
	})
	return cache_updateRec
}

var cache_initial gopurs_runtime.Value
var once_initial sync.Once
func Get_initial() gopurs_runtime.Value {
	once_initial.Do(func() {
		cache_initial = gopurs_runtime.RecordDict2("a", "b", gopurs_runtime.Int(0), gopurs_runtime.RecordDict2("c", "d", gopurs_runtime.Int(0), gopurs_runtime.RecordDict2("e", "f", gopurs_runtime.Int(0), gopurs_runtime.Int(0))))
	})
	return cache_initial
}

var cache_describe gopurs_runtime.Value
var once_describe sync.Once
func Get_describe() gopurs_runtime.Value {
	once_describe.Do(func() {
		cache_describe = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("Deep Record Updates (10k iterations):"))
	})
	return cache_describe
}

var cache_act gopurs_runtime.Value
var once_act sync.Once
func Get_act() gopurs_runtime.Value {
	once_act.Do(func() {
		cache_act = gopurs_runtime.Apply2(Get_bind__3550378017(), gopurs_runtime.Apply(pkg_Bench.Get_opaque(), gopurs_runtime.Int(10000)), gopurs_runtime.Func(func(dummy_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_logShow__2885109999(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(Call_updateRec(dummy_0.IntVal, Get_initial()), "b"), "d"), "f"))
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

var cache_mod__2185172824 gopurs_runtime.Value
var once_mod__2185172824 sync.Once
func Get_mod__2185172824() gopurs_runtime.Value {
	once_mod__2185172824.Do(func() {
		cache_mod__2185172824 = gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingInt(), "mod")
	})
	return cache_mod__2185172824
}

var cache_mod__2579358968 gopurs_runtime.Value
var once_mod__2579358968 sync.Once
func Get_mod__2579358968() gopurs_runtime.Value {
	once_mod__2579358968.Do(func() {
		cache_mod__2579358968 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mod__2579358968(gopurs_runtime.CoerceToStruct[pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_mod__2579358968
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

var cache_show__2742601362 gopurs_runtime.Value
var once_show__2742601362 sync.Once
func Get_show__2742601362() gopurs_runtime.Value {
	once_show__2742601362.Do(func() {
		cache_show__2742601362 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_show__2742601362(gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_show__2742601362
}

var cache_logShow__2885109999 gopurs_runtime.Value
var once_logShow__2885109999 sync.Once
func Get_logShow__2885109999() gopurs_runtime.Value {
	once_logShow__2885109999.Do(func() {
		cache_logShow__2885109999 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_logShow__2885109999(a_0_box)
})
	})
	return cache_logShow__2885109999
}

var cache_logShow__339054415 gopurs_runtime.Value
var once_logShow__339054415 sync.Once
func Get_logShow__339054415() gopurs_runtime.Value {
	once_logShow__339054415.Do(func() {
		cache_logShow__339054415 = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_logShow__339054415(gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]](dictShow_0_box), a_1_box)
})
	})
	return cache_logShow__339054415
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

func Call_updateRec(v_0_loop int64, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
updateRec:
for {
if false { continue updateRec }
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
v1_1_loop = func() gopurs_runtime.Value {
origVal := v1_1
if origVal.Type != gopurs_runtime.TypeRecord2 {
return gopurs_runtime.RecordUpdateDict(origVal, []string{"a", "b"}, []gopurs_runtime.Value{gopurs_runtime.Int(gopurs_runtime.Apply2(Get_add__560788792(), gopurs_runtime.RecordGet(v1_1, "a"), gopurs_runtime.Int(1)).IntVal), gopurs_runtime.RecordUpdate2(gopurs_runtime.RecordGet(v1_1, "b"), "c", gopurs_runtime.Int(gopurs_runtime.Apply2(Get_add__560788792(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_1, "b"), "c"), gopurs_runtime.Int(2)).IntVal), "d", gopurs_runtime.RecordUpdate2(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_1, "b"), "d"), "e", gopurs_runtime.Int(gopurs_runtime.Apply2(Get_add__560788792(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_1, "b"), "d"), "e"), gopurs_runtime.Int(3)).IntVal), "f", gopurs_runtime.Int(gopurs_runtime.Apply2(Get_add__560788792(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_1, "b"), "d"), "f"), gopurs_runtime.Apply2(Get_mod__2185172824(), gopurs_runtime.Int(v_0), gopurs_runtime.Int(5))).IntVal)))})
}
clone := *((*gopurs_runtime.RecordData2)(origVal.UnsafePtr))
clone.V0 = gopurs_runtime.Int(gopurs_runtime.Apply2(Get_add__560788792(), gopurs_runtime.RecordGet(v1_1, "a"), gopurs_runtime.Int(1)).IntVal)
clone.V1 = gopurs_runtime.RecordUpdate2(gopurs_runtime.RecordGet(v1_1, "b"), "c", gopurs_runtime.Int(gopurs_runtime.Apply2(Get_add__560788792(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_1, "b"), "c"), gopurs_runtime.Int(2)).IntVal), "d", gopurs_runtime.RecordUpdate2(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_1, "b"), "d"), "e", gopurs_runtime.Int(gopurs_runtime.Apply2(Get_add__560788792(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_1, "b"), "d"), "e"), gopurs_runtime.Int(3)).IntVal), "f", gopurs_runtime.Int(gopurs_runtime.Apply2(Get_add__560788792(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_1, "b"), "d"), "f"), gopurs_runtime.Apply2(Get_mod__2185172824(), gopurs_runtime.Int(v_0), gopurs_runtime.Int(5))).IntVal)))
return gopurs_runtime.Value{Type: gopurs_runtime.TypeRecord2, UnsafePtr: unsafe.Pointer(&clone)}
}()
continue updateRec
__t0 = gopurs_runtime.Value{}
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

func Call_mod__2579358968(dict_0_loop *pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
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

func Call_show__2742601362(dict_0_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_logShow__2885109999(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Show.Get_showInt(), "show"), a_0))
}

func Call_logShow__339054415(dictShow_0_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value], a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dictShow_0_loop
_ = dictShow_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
return gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(dictShow_0.V0, a_1))
}


