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
		cache_Control_Monad_ST_Internal_functorST = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Functor{1, Get_Control_Monad_ST_Internal_map_()})}
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
		cache_Control_Monad_ST_Internal_monadST = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Monad{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Monad_ST_Internal_applicativeST()))}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Control_Monad_ST_Internal_bindST()))}
})})}
	})
	return cache_Control_Monad_ST_Internal_monadST
}

var cache_Control_Monad_ST_Internal_bindST gopurs_runtime.Value
var once_Control_Monad_ST_Internal_bindST sync.Once
func Get_Control_Monad_ST_Internal_bindST() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_bindST.Do(func() {
		cache_Control_Monad_ST_Internal_bindST = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Control_Monad_ST_Internal_applyST()))}
}), Get_Control_Monad_ST_Internal_bind_()})}
	})
	return cache_Control_Monad_ST_Internal_bindST
}

var cache_Control_Monad_ST_Internal_applyST gopurs_runtime.Value
var once_Control_Monad_ST_Internal_applyST sync.Once
func Get_Control_Monad_ST_Internal_applyST() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_applyST.Do(func() {
		cache_Control_Monad_ST_Internal_applyST = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Control_Monad_ST_Internal_functorST()))}
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(f_0, gopurs_runtime.Value{})
_ = __local_var_2_0
__local_var_3_1 := gopurs_runtime.Apply(a_1, gopurs_runtime.Value{})
_ = __local_var_3_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Monad_ST_Internal_applicativeST()).V1), gopurs_runtime.Apply(__local_var_2_0, __local_var_3_1)), gopurs_runtime.Value{})
})
})
})})}
	})
	return cache_Control_Monad_ST_Internal_applyST
}

var cache_Control_Monad_ST_Internal_applicativeST gopurs_runtime.Value
var once_Control_Monad_ST_Internal_applicativeST sync.Once
func Get_Control_Monad_ST_Internal_applicativeST() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_applicativeST.Do(func() {
		cache_Control_Monad_ST_Internal_applicativeST = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Control_Monad_ST_Internal_applyST()))}
}), Get_Control_Monad_ST_Internal_pure_()})}
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
		cache_Control_Monad_ST_Internal_monadRecST = gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_MonadRec{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](Get_Control_Monad_ST_Internal_monadST()))}
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): fromDone_2_0 -> gopurs_runtime.Value
fromDone_2_0 := gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 60402430) {
__t1 = (*Constructor_Control_Monad_Rec_Class_Done)(v_3.UnsafePtr).V0
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
})
}))
_ = fromDone_2_0
// TAST (Let): __local_var_3_3 -> gopurs_runtime.Value
__local_var_3_3 := gopurs_runtime.Apply(f_0, a_1)
_ = __local_var_3_3
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_4 := gopurs_runtime.Apply(__local_var_3_3, gopurs_runtime.Value{})
_ = __local_var_4_4
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_newImpl(), __local_var_4_4), gopurs_runtime.Value{})
})
_ = __local_var_3_2
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_5 := gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Value{})
_ = __local_var_4_5
__local_var_5_6 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Control_Monad_ST_Internal_while(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_7 := (*(__local_var_4_5.PtrVal().(*interface{}))).(gopurs_runtime.Value)
_ = __local_var_5_7
var __t8 bool
{
if (__local_var_5_7.Type == 9 && __local_var_5_7.IntVal == 525585346) {
__t8 = true
goto end_branch_8
} else {

}
}
{
__t8 = false
}
end_branch_8:
return gopurs_runtime.Bool(__t8)
}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_9 := (*(__local_var_4_5.PtrVal().(*interface{}))).(gopurs_runtime.Value)
_ = __local_var_5_9
var __t12 gopurs_runtime.Value
{
if (__local_var_5_9.Type == 9 && __local_var_5_9.IntVal == 525585346) {
__t12 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_10 := gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, (*Constructor_Control_Monad_Rec_Class_Loop)(__local_var_5_9.UnsafePtr).V0), gopurs_runtime.Value{})
_ = __local_var_6_10
*(__local_var_4_5.PtrVal().(*interface{})) = __local_var_6_10
__local_var_7_11 := __local_var_6_10
_ = __local_var_7_11
return Get_Data_Unit_unit()
})
goto end_branch_12
} else {

}
}
{
if (__local_var_5_9.Type == 9 && __local_var_5_9.IntVal == 60402430) {
__t12 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
})
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_12:
return gopurs_runtime.Apply(__t12, gopurs_runtime.Value{})
})), gopurs_runtime.Value{})
_ = __local_var_5_6
__local_var_6_13 := (*(__local_var_4_5.PtrVal().(*interface{}))).(gopurs_runtime.Value)
_ = __local_var_6_13
return gopurs_runtime.Apply(fromDone_2_0, __local_var_6_13)
})
})
})})}
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
		cache_Control_Monad_ST_Internal_applicativeST__3091537981 = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Control_Monad_ST_Internal_applyST()))}
}), Get_Control_Monad_ST_Internal_pure_()})}
	})
	return cache_Control_Monad_ST_Internal_applicativeST__3091537981
}

var cache_Control_Monad_ST_Internal_applyST__2796778301 gopurs_runtime.Value
var once_Control_Monad_ST_Internal_applyST__2796778301 sync.Once
func Get_Control_Monad_ST_Internal_applyST__2796778301() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_applyST__2796778301.Do(func() {
		cache_Control_Monad_ST_Internal_applyST__2796778301 = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Control_Monad_ST_Internal_functorST()))}
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(f_0, gopurs_runtime.Value{})
_ = __local_var_2_0
__local_var_3_1 := gopurs_runtime.Apply(a_1, gopurs_runtime.Value{})
_ = __local_var_3_1
return gopurs_runtime.Apply(__local_var_2_0, __local_var_3_1)
})
})
})})}
	})
	return cache_Control_Monad_ST_Internal_applyST__2796778301
}

var cache_Control_Monad_ST_Internal_bindST__2435660861 gopurs_runtime.Value
var once_Control_Monad_ST_Internal_bindST__2435660861 sync.Once
func Get_Control_Monad_ST_Internal_bindST__2435660861() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_bindST__2435660861.Do(func() {
		cache_Control_Monad_ST_Internal_bindST__2435660861 = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Control_Monad_ST_Internal_applyST()))}
}), Get_Control_Monad_ST_Internal_bind_()})}
	})
	return cache_Control_Monad_ST_Internal_bindST__2435660861
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
		cache_Control_Monad_ST_Internal_functorST__4062753802 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Functor{1, Get_Control_Monad_ST_Internal_map_()})}
	})
	return cache_Control_Monad_ST_Internal_functorST__4062753802
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

var cache_Control_Monad_ST_Internal_monadST__2440522045 gopurs_runtime.Value
var once_Control_Monad_ST_Internal_monadST__2440522045 sync.Once
func Get_Control_Monad_ST_Internal_monadST__2440522045() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_monadST__2440522045.Do(func() {
		cache_Control_Monad_ST_Internal_monadST__2440522045 = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Monad{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Monad_ST_Internal_applicativeST()))}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Control_Monad_ST_Internal_bindST()))}
})})}
	})
	return cache_Control_Monad_ST_Internal_monadST__2440522045
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
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(a_1, gopurs_runtime.Value{})
_ = __local_var_3_1
__local_var_3_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), __local_var_3_1)
_ = __local_var_3_0
__local_var_4_2 := gopurs_runtime.Apply(b_2, gopurs_runtime.Value{})
_ = __local_var_4_2
return gopurs_runtime.Apply(__local_var_3_0, __local_var_4_2)
})
})
})})}
}

func Call_Control_Monad_ST_Internal_monoidST(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): semigroupST1_1_0 -> *Constructor_Data_Semigroup_Semigroup
semigroupST1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_3 := gopurs_runtime.Apply(a_2, gopurs_runtime.Value{})
_ = __local_var_4_3
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "append"), __local_var_4_3)
_ = __local_var_4_2
__local_var_5_4 := gopurs_runtime.Apply(b_3, gopurs_runtime.Value{})
_ = __local_var_5_4
return gopurs_runtime.Apply(__local_var_4_2, __local_var_5_4)
})
})
})))
_ = semigroupST1_1_0
// TAST (Let): __local_var_2_5 -> gopurs_runtime.Value
__local_var_2_5 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = __local_var_2_5
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(&Constructor_Data_Monoid_Monoid{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupST1_1_0)}
}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_2_5
})})}
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
