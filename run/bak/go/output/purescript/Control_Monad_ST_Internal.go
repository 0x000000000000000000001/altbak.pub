package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_ST_Internal_new gopurs_runtime.Value
var once_Control_Monad_ST_Internal_new sync.Once
func Get_Control_Monad_ST_Internal_new() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_new.Do(func() {
		cache_Control_Monad_ST_Internal_new = Get_Control_Monad_ST_Internal_newImpl()
	})
	return cache_Control_Monad_ST_Internal_new
}

var cache_Control_Monad_ST_Internal_modify_prime gopurs_runtime.Value
var once_Control_Monad_ST_Internal_modify_prime sync.Once
func Get_Control_Monad_ST_Internal_modify_prime() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_modify_prime.Do(func() {
		cache_Control_Monad_ST_Internal_modify_prime = Get_Control_Monad_ST_Internal_modifyImpl()
	})
	return cache_Control_Monad_ST_Internal_modify_prime
}

var cache_Control_Monad_ST_Internal_modify gopurs_runtime.Value
var once_Control_Monad_ST_Internal_modify sync.Once
func Get_Control_Monad_ST_Internal_modify() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_modify.Do(func() {
		cache_Control_Monad_ST_Internal_modify = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_ST_Internal_modify(f_0_box)
})
	})
	return cache_Control_Monad_ST_Internal_modify
}

var cache_Control_Monad_ST_Internal_functorST gopurs_runtime.Value
var once_Control_Monad_ST_Internal_functorST sync.Once
func Get_Control_Monad_ST_Internal_functorST() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_functorST.Do(func() {
		cache_Control_Monad_ST_Internal_functorST = gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
	})
	return cache_Control_Monad_ST_Internal_functorST
}

var cache_Control_Monad_ST_Internal_go__for gopurs_runtime.Value
var once_Control_Monad_ST_Internal_go__for sync.Once
func Get_Control_Monad_ST_Internal_go__for() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_go__for.Do(func() {
		cache_Control_Monad_ST_Internal_go__for = Get_Control_Monad_ST_Internal_forImpl()
	})
	return cache_Control_Monad_ST_Internal_go__for
}

var cache_Control_Monad_ST_Internal_monadST gopurs_runtime.Value
var once_Control_Monad_ST_Internal_monadST sync.Once
func Get_Control_Monad_ST_Internal_monadST() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_monadST.Do(func() {
		cache_Control_Monad_ST_Internal_monadST = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applicativeST()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_bindST()
}))
	})
	return cache_Control_Monad_ST_Internal_monadST
}

var cache_Control_Monad_ST_Internal_bindST gopurs_runtime.Value
var once_Control_Monad_ST_Internal_bindST sync.Once
func Get_Control_Monad_ST_Internal_bindST() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_bindST.Do(func() {
		cache_Control_Monad_ST_Internal_bindST = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applyST()
}), Get_Control_Monad_ST_Internal_bind_())
	})
	return cache_Control_Monad_ST_Internal_bindST
}

var cache_Control_Monad_ST_Internal_applyST gopurs_runtime.Value
var once_Control_Monad_ST_Internal_applyST sync.Once
func Get_Control_Monad_ST_Internal_applyST() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_applyST.Do(func() {
		cache_Control_Monad_ST_Internal_applyST = func() gopurs_runtime.Value {
// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
__local_var_0_0 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applicativeST()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applyST()
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_0_0
// TAST (Let): Bind1_1_1 -> *Constructor_Control_Bind_Bind
Bind1_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_0_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_1
// TAST (Let): Applicative0_2_2 -> *Constructor_Control_Applicative_Applicative
Applicative0_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_0_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_2_2
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_1.V1), f_3, gopurs_runtime.Func(func(f_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_1.V1), a_4, gopurs_runtime.Func(func(a_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_2.V1), gopurs_runtime.Apply(f_prime_5, a_prime_6))
}))
}))
})
}))
}()
	})
	return cache_Control_Monad_ST_Internal_applyST
}

var cache_Control_Monad_ST_Internal_applicativeST gopurs_runtime.Value
var once_Control_Monad_ST_Internal_applicativeST sync.Once
func Get_Control_Monad_ST_Internal_applicativeST() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_applicativeST.Do(func() {
		cache_Control_Monad_ST_Internal_applicativeST = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applicativeST()
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applyST()
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_1_0
// TAST (Let): Bind1_2_1 -> *Constructor_Control_Bind_Bind
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
// TAST (Let): Applicative0_3_2 -> *Constructor_Control_Applicative_Applicative
Applicative0_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_3_2
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), f_4, gopurs_runtime.Func(func(f_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), a_5, gopurs_runtime.Func(func(a_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_2.V1), gopurs_runtime.Apply(f_prime_6, a_prime_7))
}))
}))
})
}))
}), Get_Control_Monad_ST_Internal_pure_())
	})
	return cache_Control_Monad_ST_Internal_applicativeST
}

var cache_Control_Monad_ST_Internal_semigroupST gopurs_runtime.Value
var once_Control_Monad_ST_Internal_semigroupST sync.Once
func Get_Control_Monad_ST_Internal_semigroupST() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_semigroupST.Do(func() {
		cache_Control_Monad_ST_Internal_semigroupST = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_ST_Internal_semigroupST(dictSemigroup_0_box)
})
	})
	return cache_Control_Monad_ST_Internal_semigroupST
}

var cache_Control_Monad_ST_Internal_monadRecST gopurs_runtime.Value
var once_Control_Monad_ST_Internal_monadRecST sync.Once
func Get_Control_Monad_ST_Internal_monadRecST() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_monadRecST.Do(func() {
		cache_Control_Monad_ST_Internal_monadRecST = gopurs_runtime.RecordDict2("Monad0", "tailRecM", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applicativeST()
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_1 -> gopurs_runtime.Value
__local_var_5_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applicativeST()
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applyST()
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_5_1
// TAST (Let): Bind1_6_2 -> *Constructor_Control_Bind_Bind
Bind1_6_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_2
// TAST (Let): Applicative0_7_3 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_1, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_3
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_2.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_2.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_3.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
}))
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_3_0
// TAST (Let): Bind1_4_4 -> *Constructor_Control_Bind_Bind
Bind1_4_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_4
// TAST (Let): Applicative0_5_5 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_5
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_4.V1), f_6, gopurs_runtime.Func(func(f_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_4.V1), a_7, gopurs_runtime.Func(func(a_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_5.V1), gopurs_runtime.Apply(f_prime_8, a_prime_9))
}))
}))
})
}))
}), Get_Control_Monad_ST_Internal_pure_())
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_6 -> gopurs_runtime.Value
__local_var_3_6 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_7 -> gopurs_runtime.Value
__local_var_5_7 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applicativeST()
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applyST()
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_5_7
// TAST (Let): Bind1_6_8 -> *Constructor_Control_Bind_Bind
Bind1_6_8 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_8
// TAST (Let): Applicative0_7_9 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_9 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_9
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_8.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_8.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_9.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
}))
}), Get_Control_Monad_ST_Internal_pure_())
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applyST()
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_3_6
// TAST (Let): Bind1_4_10 -> *Constructor_Control_Bind_Bind
Bind1_4_10 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_6, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_10
// TAST (Let): Applicative0_5_11 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_11 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_6, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_11
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_10.V1), f_6, gopurs_runtime.Func(func(f_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_10.V1), a_7, gopurs_runtime.Func(func(a_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_11.V1), gopurs_runtime.Apply(f_prime_8, a_prime_9))
}))
}))
})
}))
}), Get_Control_Monad_ST_Internal_bind_())
}))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): fromDone_2_12 -> gopurs_runtime.Value
fromDone_2_12 := gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 60402430) {
__t13 = (*Constructor_Control_Monad_Rec_Class_Done)(v_3.UnsafePtr).V0
goto end_branch_13
} else {

}
}
{
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_13:
return __t13
})
}))
_ = fromDone_2_12
// TAST (Let): __local_var_3_15 -> gopurs_runtime.Value
__local_var_3_15 := gopurs_runtime.Apply(f_0, a_1)
_ = __local_var_3_15
// TAST (Let): __local_var_3_14 -> gopurs_runtime.Value
__local_var_3_14 := gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_16 := gopurs_runtime.Apply(__local_var_3_15, gopurs_runtime.Value{})
_ = __local_var_4_16
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_newImpl(), __local_var_4_16), gopurs_runtime.Value{})
})
_ = __local_var_3_14
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_17 := gopurs_runtime.Apply(__local_var_3_14, gopurs_runtime.Value{})
_ = __local_var_4_17
return gopurs_runtime.Apply(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Control_Bind_discardUnit(), "discard"), gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_18 -> gopurs_runtime.Value
__local_var_6_18 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_19 -> gopurs_runtime.Value
__local_var_8_19 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applicativeST()
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applyST()
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_8_19
// TAST (Let): Bind1_9_20 -> *Constructor_Control_Bind_Bind
Bind1_9_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_19, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_20
// TAST (Let): Applicative0_10_21 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_19, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_21
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_20.V1), f_11, gopurs_runtime.Func(func(f_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_20.V1), a_12, gopurs_runtime.Func(func(a_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_21.V1), gopurs_runtime.Apply(f_prime_13, a_prime_14))
}))
}))
})
}))
}), Get_Control_Monad_ST_Internal_pure_())
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applyST()
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_6_18
// TAST (Let): Bind1_7_22 -> *Constructor_Control_Bind_Bind
Bind1_7_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_18, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_22
// TAST (Let): Applicative0_8_23 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_18, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_23
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_22.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_22.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_23.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
}))
}), Get_Control_Monad_ST_Internal_bind_()})}, gopurs_runtime.Apply2(Get_Control_Monad_ST_Internal_while(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_24 := (*(__local_var_4_17.PtrVal().(*interface{}))).(gopurs_runtime.Value)
_ = __local_var_5_24
var __t25 bool
{
if (__local_var_5_24.Type == 9 && __local_var_5_24.IntVal == 525585346) {
__t25 = true
goto end_branch_25
} else {

}
}
{
__t25 = false
}
end_branch_25:
return gopurs_runtime.Bool(__t25)
}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_26 := (*(__local_var_4_17.PtrVal().(*interface{}))).(gopurs_runtime.Value)
_ = __local_var_5_26
var __t29 gopurs_runtime.Value
{
if (__local_var_5_26.Type == 9 && __local_var_5_26.IntVal == 525585346) {
__t29 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_27 := gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, (*Constructor_Control_Monad_Rec_Class_Loop)(__local_var_5_26.UnsafePtr).V0), gopurs_runtime.Value{})
_ = __local_var_6_27
*(__local_var_4_17.PtrVal().(*interface{})) = __local_var_6_27
__local_var_7_28 := __local_var_6_27
_ = __local_var_7_28
return Get_Data_Unit_unit()
})
goto end_branch_29
} else {

}
}
{
if (__local_var_5_26.Type == 9 && __local_var_5_26.IntVal == 60402430) {
__t29 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
})
goto end_branch_29
} else {

}
}
{
__t29 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_29:
return gopurs_runtime.Apply(__t29, gopurs_runtime.Value{})
})), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_30 := (*(__local_var_4_17.PtrVal().(*interface{}))).(gopurs_runtime.Value)
_ = __local_var_6_30
return gopurs_runtime.Apply(fromDone_2_12, __local_var_6_30)
})
})), gopurs_runtime.Value{})
})
})
}))
	})
	return cache_Control_Monad_ST_Internal_monadRecST
}

var cache_Control_Monad_ST_Internal_monoidST gopurs_runtime.Value
var once_Control_Monad_ST_Internal_monoidST sync.Once
func Get_Control_Monad_ST_Internal_monoidST() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_monoidST.Do(func() {
		cache_Control_Monad_ST_Internal_monoidST = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_ST_Internal_monoidST(dictMonoid_0_box)
})
	})
	return cache_Control_Monad_ST_Internal_monoidST
}

var cache_Control_Monad_ST_Internal_applicativeST__3091537981 gopurs_runtime.Value
var once_Control_Monad_ST_Internal_applicativeST__3091537981 sync.Once
func Get_Control_Monad_ST_Internal_applicativeST__3091537981() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_applicativeST__3091537981.Do(func() {
		cache_Control_Monad_ST_Internal_applicativeST__3091537981 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applicativeST()
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applyST()
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_3_1
// TAST (Let): Bind1_4_2 -> *Constructor_Control_Bind_Bind
Bind1_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_2
// TAST (Let): Applicative0_5_3 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_1, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_3
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_2.V1), f_6, gopurs_runtime.Func(func(f_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_2.V1), a_7, gopurs_runtime.Func(func(a_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_3.V1), gopurs_runtime.Apply(f_prime_8, a_prime_9))
}))
}))
})
}))
}), Get_Control_Monad_ST_Internal_pure_())
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applyST()
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_1_0
// TAST (Let): Bind1_2_4 -> *Constructor_Control_Bind_Bind
Bind1_2_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_4
// TAST (Let): Applicative0_3_5 -> *Constructor_Control_Applicative_Applicative
Applicative0_3_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_3_5
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_4.V1), f_4, gopurs_runtime.Func(func(f_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_4.V1), a_5, gopurs_runtime.Func(func(a_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_5.V1), gopurs_runtime.Apply(f_prime_6, a_prime_7))
}))
}))
})
}))
}), Get_Control_Monad_ST_Internal_pure_())
	})
	return cache_Control_Monad_ST_Internal_applicativeST__3091537981
}

var cache_Control_Monad_ST_Internal_applicativeST__2868811880 gopurs_runtime.Value
var once_Control_Monad_ST_Internal_applicativeST__2868811880 sync.Once
func Get_Control_Monad_ST_Internal_applicativeST__2868811880() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_applicativeST__2868811880.Do(func() {
		cache_Control_Monad_ST_Internal_applicativeST__2868811880 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applicativeST()
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applyST()
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_3_1
// TAST (Let): Bind1_4_2 -> *Constructor_Control_Bind_Bind
Bind1_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_2
// TAST (Let): Applicative0_5_3 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_1, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_3
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_2.V1), f_6, gopurs_runtime.Func(func(f_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_2.V1), a_7, gopurs_runtime.Func(func(a_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_3.V1), gopurs_runtime.Apply(f_prime_8, a_prime_9))
}))
}))
})
}))
}), Get_Control_Monad_ST_Internal_pure_())
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applyST()
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_1_0
// TAST (Let): Bind1_2_4 -> *Constructor_Control_Bind_Bind
Bind1_2_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_4
// TAST (Let): Applicative0_3_5 -> *Constructor_Control_Applicative_Applicative
Applicative0_3_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_3_5
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_4.V1), f_4, gopurs_runtime.Func(func(f_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_4.V1), a_5, gopurs_runtime.Func(func(a_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_5.V1), gopurs_runtime.Apply(f_prime_6, a_prime_7))
}))
}))
})
}))
}), Get_Control_Monad_ST_Internal_pure_())
	})
	return cache_Control_Monad_ST_Internal_applicativeST__2868811880
}

var cache_Control_Monad_ST_Internal_applyST__2796778301 gopurs_runtime.Value
var once_Control_Monad_ST_Internal_applyST__2796778301 sync.Once
func Get_Control_Monad_ST_Internal_applyST__2796778301() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_applyST__2796778301.Do(func() {
		cache_Control_Monad_ST_Internal_applyST__2796778301 = func() gopurs_runtime.Value {
// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
__local_var_0_0 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applicativeST()
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applicativeST()
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applyST()
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_4_2
// TAST (Let): Bind1_5_3 -> *Constructor_Control_Bind_Bind
Bind1_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_3
// TAST (Let): Applicative0_6_4 -> *Constructor_Control_Applicative_Applicative
Applicative0_6_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_4
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_3.V1), f_7, gopurs_runtime.Func(func(f_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_3.V1), a_8, gopurs_runtime.Func(func(a_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_4.V1), gopurs_runtime.Apply(f_prime_9, a_prime_10))
}))
}))
})
}))
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_2_1
// TAST (Let): Bind1_3_5 -> *Constructor_Control_Bind_Bind
Bind1_3_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_5
// TAST (Let): Applicative0_4_6 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_6
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_5.V1), f_5, gopurs_runtime.Func(func(f_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_5.V1), a_6, gopurs_runtime.Func(func(a_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_6.V1), gopurs_runtime.Apply(f_prime_7, a_prime_8))
}))
}))
})
}))
}), Get_Control_Monad_ST_Internal_pure_())
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_7 -> gopurs_runtime.Value
__local_var_2_7 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_8 -> gopurs_runtime.Value
__local_var_4_8 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applicativeST()
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applyST()
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_4_8
// TAST (Let): Bind1_5_9 -> *Constructor_Control_Bind_Bind
Bind1_5_9 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_8, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_9
// TAST (Let): Applicative0_6_10 -> *Constructor_Control_Applicative_Applicative
Applicative0_6_10 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_8, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_10
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_9.V1), f_7, gopurs_runtime.Func(func(f_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_9.V1), a_8, gopurs_runtime.Func(func(a_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_10.V1), gopurs_runtime.Apply(f_prime_9, a_prime_10))
}))
}))
})
}))
}), Get_Control_Monad_ST_Internal_pure_())
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applyST()
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_2_7
// TAST (Let): Bind1_3_11 -> *Constructor_Control_Bind_Bind
Bind1_3_11 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_11
// TAST (Let): Applicative0_4_12 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_12
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_11.V1), f_5, gopurs_runtime.Func(func(f_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_11.V1), a_6, gopurs_runtime.Func(func(a_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_12.V1), gopurs_runtime.Apply(f_prime_7, a_prime_8))
}))
}))
})
}))
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_0_0
// TAST (Let): Bind1_1_13 -> *Constructor_Control_Bind_Bind
Bind1_1_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_0_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_13
// TAST (Let): Applicative0_2_14 -> *Constructor_Control_Applicative_Applicative
Applicative0_2_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_0_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_2_14
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_13.V1), f_3, gopurs_runtime.Func(func(f_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_13.V1), a_4, gopurs_runtime.Func(func(a_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_14.V1), gopurs_runtime.Apply(f_prime_5, a_prime_6))
}))
}))
})
}))
}()
	})
	return cache_Control_Monad_ST_Internal_applyST__2796778301
}

var cache_Control_Monad_ST_Internal_applyST__3437405515 gopurs_runtime.Value
var once_Control_Monad_ST_Internal_applyST__3437405515 sync.Once
func Get_Control_Monad_ST_Internal_applyST__3437405515() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_applyST__3437405515.Do(func() {
		cache_Control_Monad_ST_Internal_applyST__3437405515 = func() gopurs_runtime.Value {
// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
__local_var_0_0 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applicativeST()
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applicativeST()
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applyST()
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_4_2
// TAST (Let): Bind1_5_3 -> *Constructor_Control_Bind_Bind
Bind1_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_3
// TAST (Let): Applicative0_6_4 -> *Constructor_Control_Applicative_Applicative
Applicative0_6_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_4
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_3.V1), f_7, gopurs_runtime.Func(func(f_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_3.V1), a_8, gopurs_runtime.Func(func(a_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_4.V1), gopurs_runtime.Apply(f_prime_9, a_prime_10))
}))
}))
})
}))
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_2_1
// TAST (Let): Bind1_3_5 -> *Constructor_Control_Bind_Bind
Bind1_3_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_5
// TAST (Let): Applicative0_4_6 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_6
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_5.V1), f_5, gopurs_runtime.Func(func(f_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_5.V1), a_6, gopurs_runtime.Func(func(a_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_6.V1), gopurs_runtime.Apply(f_prime_7, a_prime_8))
}))
}))
})
}))
}), Get_Control_Monad_ST_Internal_pure_())
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_7 -> gopurs_runtime.Value
__local_var_2_7 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_8 -> gopurs_runtime.Value
__local_var_4_8 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applicativeST()
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applyST()
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_4_8
// TAST (Let): Bind1_5_9 -> *Constructor_Control_Bind_Bind
Bind1_5_9 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_8, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_9
// TAST (Let): Applicative0_6_10 -> *Constructor_Control_Applicative_Applicative
Applicative0_6_10 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_8, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_10
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_9.V1), f_7, gopurs_runtime.Func(func(f_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_9.V1), a_8, gopurs_runtime.Func(func(a_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_10.V1), gopurs_runtime.Apply(f_prime_9, a_prime_10))
}))
}))
})
}))
}), Get_Control_Monad_ST_Internal_pure_())
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applyST()
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_2_7
// TAST (Let): Bind1_3_11 -> *Constructor_Control_Bind_Bind
Bind1_3_11 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_11
// TAST (Let): Applicative0_4_12 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_12
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_11.V1), f_5, gopurs_runtime.Func(func(f_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_11.V1), a_6, gopurs_runtime.Func(func(a_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_12.V1), gopurs_runtime.Apply(f_prime_7, a_prime_8))
}))
}))
})
}))
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_0_0
// TAST (Let): Bind1_1_13 -> *Constructor_Control_Bind_Bind
Bind1_1_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_0_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_13
// TAST (Let): Applicative0_2_14 -> *Constructor_Control_Applicative_Applicative
Applicative0_2_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_0_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_2_14
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_13.V1), f_3, gopurs_runtime.Func(func(f_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_13.V1), a_4, gopurs_runtime.Func(func(a_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_14.V1), gopurs_runtime.Apply(f_prime_5, a_prime_6))
}))
}))
})
}))
}()
	})
	return cache_Control_Monad_ST_Internal_applyST__3437405515
}

var cache_Control_Monad_ST_Internal_applyST__2741064779 gopurs_runtime.Value
var once_Control_Monad_ST_Internal_applyST__2741064779 sync.Once
func Get_Control_Monad_ST_Internal_applyST__2741064779() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_applyST__2741064779.Do(func() {
		cache_Control_Monad_ST_Internal_applyST__2741064779 = func() gopurs_runtime.Value {
// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
__local_var_0_0 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applicativeST()
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applicativeST()
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applyST()
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_4_2
// TAST (Let): Bind1_5_3 -> *Constructor_Control_Bind_Bind
Bind1_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_3
// TAST (Let): Applicative0_6_4 -> *Constructor_Control_Applicative_Applicative
Applicative0_6_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_4
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_3.V1), f_7, gopurs_runtime.Func(func(f_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_3.V1), a_8, gopurs_runtime.Func(func(a_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_4.V1), gopurs_runtime.Apply(f_prime_9, a_prime_10))
}))
}))
})
}))
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_2_1
// TAST (Let): Bind1_3_5 -> *Constructor_Control_Bind_Bind
Bind1_3_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_5
// TAST (Let): Applicative0_4_6 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_6
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_5.V1), f_5, gopurs_runtime.Func(func(f_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_5.V1), a_6, gopurs_runtime.Func(func(a_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_6.V1), gopurs_runtime.Apply(f_prime_7, a_prime_8))
}))
}))
})
}))
}), Get_Control_Monad_ST_Internal_pure_())
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_7 -> gopurs_runtime.Value
__local_var_2_7 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_8 -> gopurs_runtime.Value
__local_var_4_8 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applicativeST()
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applyST()
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_4_8
// TAST (Let): Bind1_5_9 -> *Constructor_Control_Bind_Bind
Bind1_5_9 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_8, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_9
// TAST (Let): Applicative0_6_10 -> *Constructor_Control_Applicative_Applicative
Applicative0_6_10 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_8, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_10
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_9.V1), f_7, gopurs_runtime.Func(func(f_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_9.V1), a_8, gopurs_runtime.Func(func(a_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_10.V1), gopurs_runtime.Apply(f_prime_9, a_prime_10))
}))
}))
})
}))
}), Get_Control_Monad_ST_Internal_pure_())
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applyST()
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_2_7
// TAST (Let): Bind1_3_11 -> *Constructor_Control_Bind_Bind
Bind1_3_11 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_11
// TAST (Let): Applicative0_4_12 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_12
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_11.V1), f_5, gopurs_runtime.Func(func(f_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_11.V1), a_6, gopurs_runtime.Func(func(a_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_12.V1), gopurs_runtime.Apply(f_prime_7, a_prime_8))
}))
}))
})
}))
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_0_0
// TAST (Let): Bind1_1_13 -> *Constructor_Control_Bind_Bind
Bind1_1_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_0_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_13
// TAST (Let): Applicative0_2_14 -> *Constructor_Control_Applicative_Applicative
Applicative0_2_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_0_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_2_14
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_13.V1), f_3, gopurs_runtime.Func(func(f_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_13.V1), a_4, gopurs_runtime.Func(func(a_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_14.V1), gopurs_runtime.Apply(f_prime_5, a_prime_6))
}))
}))
})
}))
}()
	})
	return cache_Control_Monad_ST_Internal_applyST__2741064779
}

var cache_Control_Monad_ST_Internal_applyST__1776171083 gopurs_runtime.Value
var once_Control_Monad_ST_Internal_applyST__1776171083 sync.Once
func Get_Control_Monad_ST_Internal_applyST__1776171083() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_applyST__1776171083.Do(func() {
		cache_Control_Monad_ST_Internal_applyST__1776171083 = func() gopurs_runtime.Value {
// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
__local_var_0_0 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applicativeST()
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applicativeST()
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applyST()
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_4_2
// TAST (Let): Bind1_5_3 -> *Constructor_Control_Bind_Bind
Bind1_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_3
// TAST (Let): Applicative0_6_4 -> *Constructor_Control_Applicative_Applicative
Applicative0_6_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_4
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_3.V1), f_7, gopurs_runtime.Func(func(f_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_3.V1), a_8, gopurs_runtime.Func(func(a_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_4.V1), gopurs_runtime.Apply(f_prime_9, a_prime_10))
}))
}))
})
}))
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_2_1
// TAST (Let): Bind1_3_5 -> *Constructor_Control_Bind_Bind
Bind1_3_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_5
// TAST (Let): Applicative0_4_6 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_6
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_5.V1), f_5, gopurs_runtime.Func(func(f_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_5.V1), a_6, gopurs_runtime.Func(func(a_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_6.V1), gopurs_runtime.Apply(f_prime_7, a_prime_8))
}))
}))
})
}))
}), Get_Control_Monad_ST_Internal_pure_())
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_7 -> gopurs_runtime.Value
__local_var_2_7 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_8 -> gopurs_runtime.Value
__local_var_4_8 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applicativeST()
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applyST()
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_4_8
// TAST (Let): Bind1_5_9 -> *Constructor_Control_Bind_Bind
Bind1_5_9 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_8, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_9
// TAST (Let): Applicative0_6_10 -> *Constructor_Control_Applicative_Applicative
Applicative0_6_10 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_8, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_10
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_9.V1), f_7, gopurs_runtime.Func(func(f_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_9.V1), a_8, gopurs_runtime.Func(func(a_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_10.V1), gopurs_runtime.Apply(f_prime_9, a_prime_10))
}))
}))
})
}))
}), Get_Control_Monad_ST_Internal_pure_())
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applyST()
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_2_7
// TAST (Let): Bind1_3_11 -> *Constructor_Control_Bind_Bind
Bind1_3_11 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_11
// TAST (Let): Applicative0_4_12 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_12
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_11.V1), f_5, gopurs_runtime.Func(func(f_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_11.V1), a_6, gopurs_runtime.Func(func(a_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_12.V1), gopurs_runtime.Apply(f_prime_7, a_prime_8))
}))
}))
})
}))
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_0_0
// TAST (Let): Bind1_1_13 -> *Constructor_Control_Bind_Bind
Bind1_1_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_0_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_13
// TAST (Let): Applicative0_2_14 -> *Constructor_Control_Applicative_Applicative
Applicative0_2_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_0_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_2_14
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_13.V1), f_3, gopurs_runtime.Func(func(f_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_13.V1), a_4, gopurs_runtime.Func(func(a_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_14.V1), gopurs_runtime.Apply(f_prime_5, a_prime_6))
}))
}))
})
}))
}()
	})
	return cache_Control_Monad_ST_Internal_applyST__1776171083
}

var cache_Control_Monad_ST_Internal_bindST__2435660861 gopurs_runtime.Value
var once_Control_Monad_ST_Internal_bindST__2435660861 sync.Once
func Get_Control_Monad_ST_Internal_bindST__2435660861() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_bindST__2435660861.Do(func() {
		cache_Control_Monad_ST_Internal_bindST__2435660861 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applicativeST()
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applyST()
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_3_1
// TAST (Let): Bind1_4_2 -> *Constructor_Control_Bind_Bind
Bind1_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_2
// TAST (Let): Applicative0_5_3 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_1, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_3
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_2.V1), f_6, gopurs_runtime.Func(func(f_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_2.V1), a_7, gopurs_runtime.Func(func(a_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_3.V1), gopurs_runtime.Apply(f_prime_8, a_prime_9))
}))
}))
})
}))
}), Get_Control_Monad_ST_Internal_pure_())
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applyST()
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_1_0
// TAST (Let): Bind1_2_4 -> *Constructor_Control_Bind_Bind
Bind1_2_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_4
// TAST (Let): Applicative0_3_5 -> *Constructor_Control_Applicative_Applicative
Applicative0_3_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_3_5
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_4.V1), f_4, gopurs_runtime.Func(func(f_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_4.V1), a_5, gopurs_runtime.Func(func(a_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_5.V1), gopurs_runtime.Apply(f_prime_6, a_prime_7))
}))
}))
})
}))
}), Get_Control_Monad_ST_Internal_bind_())
	})
	return cache_Control_Monad_ST_Internal_bindST__2435660861
}

var cache_Control_Monad_ST_Internal_bindST__4187656679 gopurs_runtime.Value
var once_Control_Monad_ST_Internal_bindST__4187656679 sync.Once
func Get_Control_Monad_ST_Internal_bindST__4187656679() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_bindST__4187656679.Do(func() {
		cache_Control_Monad_ST_Internal_bindST__4187656679 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applicativeST()
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applyST()
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_3_1
// TAST (Let): Bind1_4_2 -> *Constructor_Control_Bind_Bind
Bind1_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_2
// TAST (Let): Applicative0_5_3 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_1, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_3
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_2.V1), f_6, gopurs_runtime.Func(func(f_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_2.V1), a_7, gopurs_runtime.Func(func(a_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_3.V1), gopurs_runtime.Apply(f_prime_8, a_prime_9))
}))
}))
})
}))
}), Get_Control_Monad_ST_Internal_pure_())
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applyST()
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_1_0
// TAST (Let): Bind1_2_4 -> *Constructor_Control_Bind_Bind
Bind1_2_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_4
// TAST (Let): Applicative0_3_5 -> *Constructor_Control_Applicative_Applicative
Applicative0_3_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_3_5
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_4.V1), f_4, gopurs_runtime.Func(func(f_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_4.V1), a_5, gopurs_runtime.Func(func(a_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_5.V1), gopurs_runtime.Apply(f_prime_6, a_prime_7))
}))
}))
})
}))
}), Get_Control_Monad_ST_Internal_bind_())
	})
	return cache_Control_Monad_ST_Internal_bindST__4187656679
}

var cache_Control_Monad_ST_Internal_for__1203933728 gopurs_runtime.Value
var once_Control_Monad_ST_Internal_for__1203933728 sync.Once
func Get_Control_Monad_ST_Internal_for__1203933728() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_for__1203933728.Do(func() {
		cache_Control_Monad_ST_Internal_for__1203933728 = Get_Control_Monad_ST_Internal_forImpl()
	})
	return cache_Control_Monad_ST_Internal_for__1203933728
}

var cache_Control_Monad_ST_Internal_for__956375728 gopurs_runtime.Value
var once_Control_Monad_ST_Internal_for__956375728 sync.Once
func Get_Control_Monad_ST_Internal_for__956375728() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_for__956375728.Do(func() {
		cache_Control_Monad_ST_Internal_for__956375728 = Get_Control_Monad_ST_Internal_forImpl()
	})
	return cache_Control_Monad_ST_Internal_for__956375728
}

var cache_Control_Monad_ST_Internal_functorST__4062753802 gopurs_runtime.Value
var once_Control_Monad_ST_Internal_functorST__4062753802 sync.Once
func Get_Control_Monad_ST_Internal_functorST__4062753802() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_functorST__4062753802.Do(func() {
		cache_Control_Monad_ST_Internal_functorST__4062753802 = gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
	})
	return cache_Control_Monad_ST_Internal_functorST__4062753802
}

var cache_Control_Monad_ST_Internal_functorST__1148171281 gopurs_runtime.Value
var once_Control_Monad_ST_Internal_functorST__1148171281 sync.Once
func Get_Control_Monad_ST_Internal_functorST__1148171281() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_functorST__1148171281.Do(func() {
		cache_Control_Monad_ST_Internal_functorST__1148171281 = gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
	})
	return cache_Control_Monad_ST_Internal_functorST__1148171281
}

var cache_Control_Monad_ST_Internal_functorST__2441840241 gopurs_runtime.Value
var once_Control_Monad_ST_Internal_functorST__2441840241 sync.Once
func Get_Control_Monad_ST_Internal_functorST__2441840241() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_functorST__2441840241.Do(func() {
		cache_Control_Monad_ST_Internal_functorST__2441840241 = gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
	})
	return cache_Control_Monad_ST_Internal_functorST__2441840241
}

var cache_Control_Monad_ST_Internal_functorST__1848446545 gopurs_runtime.Value
var once_Control_Monad_ST_Internal_functorST__1848446545 sync.Once
func Get_Control_Monad_ST_Internal_functorST__1848446545() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_functorST__1848446545.Do(func() {
		cache_Control_Monad_ST_Internal_functorST__1848446545 = gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
	})
	return cache_Control_Monad_ST_Internal_functorST__1848446545
}

var cache_Control_Monad_ST_Internal_modify__3866314397 gopurs_runtime.Value
var once_Control_Monad_ST_Internal_modify__3866314397 sync.Once
func Get_Control_Monad_ST_Internal_modify__3866314397() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_modify__3866314397.Do(func() {
		cache_Control_Monad_ST_Internal_modify__3866314397 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_ST_Internal_modify__3866314397(f_0_box)
})
	})
	return cache_Control_Monad_ST_Internal_modify__3866314397
}

var cache_Control_Monad_ST_Internal_modify__781734141 gopurs_runtime.Value
var once_Control_Monad_ST_Internal_modify__781734141 sync.Once
func Get_Control_Monad_ST_Internal_modify__781734141() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_modify__781734141.Do(func() {
		cache_Control_Monad_ST_Internal_modify__781734141 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_ST_Internal_modify__781734141(f_0_box)
})
	})
	return cache_Control_Monad_ST_Internal_modify__781734141
}

var cache_Control_Monad_ST_Internal_modify__2563484957 gopurs_runtime.Value
var once_Control_Monad_ST_Internal_modify__2563484957 sync.Once
func Get_Control_Monad_ST_Internal_modify__2563484957() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_modify__2563484957.Do(func() {
		cache_Control_Monad_ST_Internal_modify__2563484957 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_ST_Internal_modify__2563484957(f_0_box)
})
	})
	return cache_Control_Monad_ST_Internal_modify__2563484957
}

var cache_Control_Monad_ST_Internal_modify_prime__2036662078 gopurs_runtime.Value
var once_Control_Monad_ST_Internal_modify_prime__2036662078 sync.Once
func Get_Control_Monad_ST_Internal_modify_prime__2036662078() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_modify_prime__2036662078.Do(func() {
		cache_Control_Monad_ST_Internal_modify_prime__2036662078 = Get_Control_Monad_ST_Internal_modifyImpl()
	})
	return cache_Control_Monad_ST_Internal_modify_prime__2036662078
}

var cache_Control_Monad_ST_Internal_modify_prime__2600755550 gopurs_runtime.Value
var once_Control_Monad_ST_Internal_modify_prime__2600755550 sync.Once
func Get_Control_Monad_ST_Internal_modify_prime__2600755550() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_modify_prime__2600755550.Do(func() {
		cache_Control_Monad_ST_Internal_modify_prime__2600755550 = Get_Control_Monad_ST_Internal_modifyImpl()
	})
	return cache_Control_Monad_ST_Internal_modify_prime__2600755550
}

var cache_Control_Monad_ST_Internal_modify_prime__1497736571 gopurs_runtime.Value
var once_Control_Monad_ST_Internal_modify_prime__1497736571 sync.Once
func Get_Control_Monad_ST_Internal_modify_prime__1497736571() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_modify_prime__1497736571.Do(func() {
		cache_Control_Monad_ST_Internal_modify_prime__1497736571 = Get_Control_Monad_ST_Internal_modifyImpl()
	})
	return cache_Control_Monad_ST_Internal_modify_prime__1497736571
}

var cache_Control_Monad_ST_Internal_modify_prime__3123720915 gopurs_runtime.Value
var once_Control_Monad_ST_Internal_modify_prime__3123720915 sync.Once
func Get_Control_Monad_ST_Internal_modify_prime__3123720915() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_modify_prime__3123720915.Do(func() {
		cache_Control_Monad_ST_Internal_modify_prime__3123720915 = Get_Control_Monad_ST_Internal_modifyImpl()
	})
	return cache_Control_Monad_ST_Internal_modify_prime__3123720915
}

var cache_Control_Monad_ST_Internal_modify_prime__3525811603 gopurs_runtime.Value
var once_Control_Monad_ST_Internal_modify_prime__3525811603 sync.Once
func Get_Control_Monad_ST_Internal_modify_prime__3525811603() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_modify_prime__3525811603.Do(func() {
		cache_Control_Monad_ST_Internal_modify_prime__3525811603 = Get_Control_Monad_ST_Internal_modifyImpl()
	})
	return cache_Control_Monad_ST_Internal_modify_prime__3525811603
}

var cache_Control_Monad_ST_Internal_monadST__1413783571 gopurs_runtime.Value
var once_Control_Monad_ST_Internal_monadST__1413783571 sync.Once
func Get_Control_Monad_ST_Internal_monadST__1413783571() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_monadST__1413783571.Do(func() {
		cache_Control_Monad_ST_Internal_monadST__1413783571 = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applicativeST()
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applicativeST()
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applyST()
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_4_1
// TAST (Let): Bind1_5_2 -> *Constructor_Control_Bind_Bind
Bind1_5_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_2
// TAST (Let): Applicative0_6_3 -> *Constructor_Control_Applicative_Applicative
Applicative0_6_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_1, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_3
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_2.V1), f_7, gopurs_runtime.Func(func(f_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_2.V1), a_8, gopurs_runtime.Func(func(a_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_3.V1), gopurs_runtime.Apply(f_prime_9, a_prime_10))
}))
}))
})
}))
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_2_0
// TAST (Let): Bind1_3_4 -> *Constructor_Control_Bind_Bind
Bind1_3_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_4
// TAST (Let): Applicative0_4_5 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_5
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_4.V1), f_5, gopurs_runtime.Func(func(f_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_4.V1), a_6, gopurs_runtime.Func(func(a_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_5.V1), gopurs_runtime.Apply(f_prime_7, a_prime_8))
}))
}))
})
}))
}), Get_Control_Monad_ST_Internal_pure_())
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_6 -> gopurs_runtime.Value
__local_var_2_6 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_7 -> gopurs_runtime.Value
__local_var_4_7 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applicativeST()
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applyST()
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_4_7
// TAST (Let): Bind1_5_8 -> *Constructor_Control_Bind_Bind
Bind1_5_8 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_8
// TAST (Let): Applicative0_6_9 -> *Constructor_Control_Applicative_Applicative
Applicative0_6_9 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_9
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_8.V1), f_7, gopurs_runtime.Func(func(f_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_8.V1), a_8, gopurs_runtime.Func(func(a_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_9.V1), gopurs_runtime.Apply(f_prime_9, a_prime_10))
}))
}))
})
}))
}), Get_Control_Monad_ST_Internal_pure_())
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applyST()
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_2_6
// TAST (Let): Bind1_3_10 -> *Constructor_Control_Bind_Bind
Bind1_3_10 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_6, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_10
// TAST (Let): Applicative0_4_11 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_11 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_6, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_11
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_10.V1), f_5, gopurs_runtime.Func(func(f_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_10.V1), a_6, gopurs_runtime.Func(func(a_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_11.V1), gopurs_runtime.Apply(f_prime_7, a_prime_8))
}))
}))
})
}))
}), Get_Control_Monad_ST_Internal_bind_())
}))
	})
	return cache_Control_Monad_ST_Internal_monadST__1413783571
}

var cache_Control_Monad_ST_Internal_new__3489595018 gopurs_runtime.Value
var once_Control_Monad_ST_Internal_new__3489595018 sync.Once
func Get_Control_Monad_ST_Internal_new__3489595018() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_new__3489595018.Do(func() {
		cache_Control_Monad_ST_Internal_new__3489595018 = Get_Control_Monad_ST_Internal_newImpl()
	})
	return cache_Control_Monad_ST_Internal_new__3489595018
}

var cache_Control_Monad_ST_Internal_new__3579768924 gopurs_runtime.Value
var once_Control_Monad_ST_Internal_new__3579768924 sync.Once
func Get_Control_Monad_ST_Internal_new__3579768924() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_new__3579768924.Do(func() {
		cache_Control_Monad_ST_Internal_new__3579768924 = Get_Control_Monad_ST_Internal_newImpl()
	})
	return cache_Control_Monad_ST_Internal_new__3579768924
}

var cache_Control_Monad_ST_Internal_new__122671164 gopurs_runtime.Value
var once_Control_Monad_ST_Internal_new__122671164 sync.Once
func Get_Control_Monad_ST_Internal_new__122671164() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_new__122671164.Do(func() {
		cache_Control_Monad_ST_Internal_new__122671164 = Get_Control_Monad_ST_Internal_newImpl()
	})
	return cache_Control_Monad_ST_Internal_new__122671164
}

var cache_Control_Monad_ST_Internal_new__2010968700 gopurs_runtime.Value
var once_Control_Monad_ST_Internal_new__2010968700 sync.Once
func Get_Control_Monad_ST_Internal_new__2010968700() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_new__2010968700.Do(func() {
		cache_Control_Monad_ST_Internal_new__2010968700 = Get_Control_Monad_ST_Internal_newImpl()
	})
	return cache_Control_Monad_ST_Internal_new__2010968700
}

var cache_Control_Monad_ST_Internal_new__8318140 gopurs_runtime.Value
var once_Control_Monad_ST_Internal_new__8318140 sync.Once
func Get_Control_Monad_ST_Internal_new__8318140() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_new__8318140.Do(func() {
		cache_Control_Monad_ST_Internal_new__8318140 = Get_Control_Monad_ST_Internal_newImpl()
	})
	return cache_Control_Monad_ST_Internal_new__8318140
}

func Call_Control_Monad_ST_Internal_modify(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
return gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_modifyImpl(), gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): s_prime_2_0 -> gopurs_runtime.Value
s_prime_2_0 := gopurs_runtime.Apply(f_0, s_1)
_ = s_prime_2_0
return gopurs_runtime.RecordDict2("state", "value", s_prime_2_0, s_prime_2_0)
}))
}

func Call_Control_Monad_ST_Internal_semigroupST(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applicativeST()
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applyST()
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_3_2
// TAST (Let): Bind1_4_3 -> *Constructor_Control_Bind_Bind
Bind1_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_3
// TAST (Let): Applicative0_5_4 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_4
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_3.V1), f_6, gopurs_runtime.Func(func(f_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_3.V1), a_7, gopurs_runtime.Func(func(a_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_4.V1), gopurs_runtime.Apply(f_prime_8, a_prime_9))
}))
}))
})
}))
}), Get_Control_Monad_ST_Internal_pure_())
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applyST()
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_1_1
// TAST (Let): Bind1_2_5 -> *Constructor_Control_Bind_Bind
Bind1_2_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_5
// TAST (Let): Applicative0_3_6 -> *Constructor_Control_Applicative_Applicative
Applicative0_3_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_3_6
// TAST (Let): __local_var_1_0 -> *Constructor_Control_Apply_Apply
__local_var_1_0 := &Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_5.V1), f_4, gopurs_runtime.Func(func(f_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_5.V1), a_5, gopurs_runtime.Func(func(a_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_6.V1), gopurs_runtime.Apply(f_prime_6, a_prime_7))
}))
}))
})
})}
_ = __local_var_1_0
// TAST (Let): Functor0_2_7 -> *Constructor_Data_Functor_Functor
Functor0_2_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_1_0.V0), gopurs_runtime.Value{}))
_ = Functor0_2_7
// TAST (Let): __local_var_3_8 -> gopurs_runtime.Value
__local_var_3_8 := gopurs_runtime.RecordGet(dictSemigroup_0, "append")
_ = __local_var_3_8
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(__local_var_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_7.V0), __local_var_3_8, a_4), b_5)
})
}))
}

func Call_Control_Monad_ST_Internal_monoidST(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): __local_var_1_2 -> gopurs_runtime.Value
__local_var_1_2 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_3 -> gopurs_runtime.Value
__local_var_3_3 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_4 -> gopurs_runtime.Value
__local_var_5_4 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applicativeST()
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_5 -> gopurs_runtime.Value
__local_var_7_5 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applicativeST()
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applyST()
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_7_5
// TAST (Let): Bind1_8_6 -> *Constructor_Control_Bind_Bind
Bind1_8_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_5, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_6
// TAST (Let): Applicative0_9_7 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_5, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_7
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_6.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_6.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_7.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_5_4
// TAST (Let): Bind1_6_8 -> *Constructor_Control_Bind_Bind
Bind1_6_8 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_8
// TAST (Let): Applicative0_7_9 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_9 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_9
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_8.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_8.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_9.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
}))
}), Get_Control_Monad_ST_Internal_pure_())
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_10 -> gopurs_runtime.Value
__local_var_5_10 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_11 -> gopurs_runtime.Value
__local_var_7_11 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applicativeST()
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applyST()
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_7_11
// TAST (Let): Bind1_8_12 -> *Constructor_Control_Bind_Bind
Bind1_8_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_11, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_12
// TAST (Let): Applicative0_9_13 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_11, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_13
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_12.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_12.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_13.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), Get_Control_Monad_ST_Internal_pure_())
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applyST()
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_5_10
// TAST (Let): Bind1_6_14 -> *Constructor_Control_Bind_Bind
Bind1_6_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_10, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_14
// TAST (Let): Applicative0_7_15 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_15 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_10, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_15
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_14.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_14.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_15.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
}))
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_3_3
// TAST (Let): Bind1_4_16 -> *Constructor_Control_Bind_Bind
Bind1_4_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_16
// TAST (Let): Applicative0_5_17 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_17
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_16.V1), f_6, gopurs_runtime.Func(func(f_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_16.V1), a_7, gopurs_runtime.Func(func(a_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_17.V1), gopurs_runtime.Apply(f_prime_8, a_prime_9))
}))
}))
})
}))
}), Get_Control_Monad_ST_Internal_pure_())
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_18 -> gopurs_runtime.Value
__local_var_3_18 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_19 -> gopurs_runtime.Value
__local_var_5_19 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applicativeST()
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applyST()
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_5_19
// TAST (Let): Bind1_6_20 -> *Constructor_Control_Bind_Bind
Bind1_6_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_19, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_20
// TAST (Let): Applicative0_7_21 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_19, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_21
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_20.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_20.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_21.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
}))
}), Get_Control_Monad_ST_Internal_pure_())
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applyST()
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_3_18
// TAST (Let): Bind1_4_22 -> *Constructor_Control_Bind_Bind
Bind1_4_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_18, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_22
// TAST (Let): Applicative0_5_23 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_18, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_23
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_22.V1), f_6, gopurs_runtime.Func(func(f_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_22.V1), a_7, gopurs_runtime.Func(func(a_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_23.V1), gopurs_runtime.Apply(f_prime_8, a_prime_9))
}))
}))
})
}))
}), Get_Control_Monad_ST_Internal_bind_())
}))
_ = __local_var_1_2
// TAST (Let): Bind1_2_24 -> *Constructor_Control_Bind_Bind
Bind1_2_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_24
// TAST (Let): Applicative0_3_25 -> *Constructor_Control_Applicative_Applicative
Applicative0_3_25 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_3_25
// TAST (Let): __local_var_1_1 -> *Constructor_Control_Apply_Apply
__local_var_1_1 := &Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", Get_Control_Monad_ST_Internal_map_())
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_24.V1), f_4, gopurs_runtime.Func(func(f_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_24.V1), a_5, gopurs_runtime.Func(func(a_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_25.V1), gopurs_runtime.Apply(f_prime_6, a_prime_7))
}))
}))
})
})}
_ = __local_var_1_1
// TAST (Let): Functor0_2_26 -> *Constructor_Data_Functor_Functor
Functor0_2_26 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_1_1.V0), gopurs_runtime.Value{}))
_ = Functor0_2_26
// TAST (Let): __local_var_3_27 -> gopurs_runtime.Value
__local_var_3_27 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}), "append")
_ = __local_var_3_27
// TAST (Let): semigroupST1_1_0 -> gopurs_runtime.Value
semigroupST1_1_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(__local_var_1_1.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_26.V0), __local_var_3_27, a_4), b_5)
})
}))
_ = semigroupST1_1_0
// TAST (Let): __local_var_2_28 -> gopurs_runtime.Value
__local_var_2_28 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = __local_var_2_28
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupST1_1_0
}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_2_28
}))
}

func Call_Control_Monad_ST_Internal_modify__3866314397(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
return gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_modifyImpl(), gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): s_prime_2_0 -> int64
s_prime_2_0 := gopurs_runtime.Apply(f_0, gopurs_runtime.Int(s_1.IntVal)).IntVal
_ = s_prime_2_0
return gopurs_runtime.RecordDict2("state", "value", gopurs_runtime.Int(s_prime_2_0), gopurs_runtime.Int(s_prime_2_0))
}))
}

func Call_Control_Monad_ST_Internal_modify__781734141(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
return gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_modifyImpl(), gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): s_prime_2_0 -> gopurs_runtime.Value
s_prime_2_0 := gopurs_runtime.Apply(f_0, s_1)
_ = s_prime_2_0
return gopurs_runtime.RecordDict2("state", "value", s_prime_2_0, s_prime_2_0)
}))
}

func Call_Control_Monad_ST_Internal_modify__2563484957(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
return gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_modifyImpl(), gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): s_prime_2_0 -> *Constructor_Data_Tuple_Tuple
s_prime_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](s_1))}))
_ = s_prime_2_0
return gopurs_runtime.RecordDict2("state", "value", gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(s_prime_2_0)}, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(s_prime_2_0)})
}))
}

func Get_Control_Monad_ST_Internal_bind_() gopurs_runtime.Value {
	return _Gopurs_Control_Monad_ST_Internal_Bind_
}

func Get_Control_Monad_ST_Internal_forImpl() gopurs_runtime.Value {
	return _Gopurs_Control_Monad_ST_Internal_ForImpl
}

func Get_Control_Monad_ST_Internal_foreach() gopurs_runtime.Value {
	return _Gopurs_Control_Monad_ST_Internal_Foreach
}

func Get_Control_Monad_ST_Internal_map_() gopurs_runtime.Value {
	return _Gopurs_Control_Monad_ST_Internal_Map_
}

func Get_Control_Monad_ST_Internal_modifyImpl() gopurs_runtime.Value {
	return _Gopurs_Control_Monad_ST_Internal_ModifyImpl
}

func Get_Control_Monad_ST_Internal_newImpl() gopurs_runtime.Value {
	return _Gopurs_Control_Monad_ST_Internal_NewImpl
}

func Get_Control_Monad_ST_Internal_pure_() gopurs_runtime.Value {
	return _Gopurs_Control_Monad_ST_Internal_Pure_
}

func Get_Control_Monad_ST_Internal_read() gopurs_runtime.Value {
	return _Gopurs_Control_Monad_ST_Internal_Read
}

func Get_Control_Monad_ST_Internal_run() gopurs_runtime.Value {
	return _Gopurs_Control_Monad_ST_Internal_Run
}

func Get_Control_Monad_ST_Internal_while() gopurs_runtime.Value {
	return _Gopurs_Control_Monad_ST_Internal_While
}

func Get_Control_Monad_ST_Internal_write() gopurs_runtime.Value {
	return _Gopurs_Control_Monad_ST_Internal_Write
}
