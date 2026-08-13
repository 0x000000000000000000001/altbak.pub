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
// TAST (Let): Bind1_0_0 -> *Constructor_Control_Bind_Bind[gopurs_runtime.Value]
Bind1_0_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_monadST(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
// TAST (Let): Applicative0_1_1 -> *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_monadST(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_functorST()
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_0_0.V1), f_2, gopurs_runtime.Func(func(f_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_0_0.V1), a_3, gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_1.V1), gopurs_runtime.Apply(f_prime_4, a_prime_5))
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
return Get_Control_Monad_ST_Internal_applyST()
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
return Get_Control_Monad_ST_Internal_monadST()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): fromDone_2_0 -> gopurs_runtime.Value
fromDone_2_0 := gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 60402430) {
__t1 = (*Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0
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
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Apply(f_0, a_1), Get_Control_Monad_ST_Internal_newImpl()), gopurs_runtime.Func(func(r_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Control_Bind_discardUnit(), "discard"), gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Control_Monad_ST_Internal_bindST()))}, gopurs_runtime.Apply2(Get_Control_Monad_ST_Internal_while(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_functorST(), "map"), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 bool
{
if (v_4.Type == 9 && v_4.IntVal == 525585346) {
__t2 = true
goto end_branch_2
} else {

}
}
{
__t2 = false
}
end_branch_2:
return gopurs_runtime.Bool(__t2)
}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return (*(r_3.PtrVal().(*interface{}))).(gopurs_runtime.Value)
})), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return (*(r_3.PtrVal().(*interface{}))).(gopurs_runtime.Value)
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 525585346) {
__t3 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Apply(f_0, (*Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), gopurs_runtime.Func(func(e_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_functorST(), "map"), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
*(r_3.PtrVal().(*interface{})) = e_5
return e_5
}))
}))
goto end_branch_3
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 60402430) {
__t3 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_applicativeST(), "pure"), Get_Data_Unit_unit())
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}))), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_functorST(), "map"), fromDone_2_0, gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return (*(r_3.PtrVal().(*interface{}))).(gopurs_runtime.Value)
}))
}))
}))
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
return Get_Control_Monad_ST_Internal_applyST()
}), Get_Control_Monad_ST_Internal_pure_())
	})
	return cache_Control_Monad_ST_Internal_applicativeST__3091537981
}

var cache_Control_Monad_ST_Internal_applicativeST__2868811880 gopurs_runtime.Value
var once_Control_Monad_ST_Internal_applicativeST__2868811880 sync.Once
func Get_Control_Monad_ST_Internal_applicativeST__2868811880() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_applicativeST__2868811880.Do(func() {
		cache_Control_Monad_ST_Internal_applicativeST__2868811880 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applyST()
}), Get_Control_Monad_ST_Internal_pure_())
	})
	return cache_Control_Monad_ST_Internal_applicativeST__2868811880
}

var cache_Control_Monad_ST_Internal_applyST__2796778301 gopurs_runtime.Value
var once_Control_Monad_ST_Internal_applyST__2796778301 sync.Once
func Get_Control_Monad_ST_Internal_applyST__2796778301() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_applyST__2796778301.Do(func() {
		cache_Control_Monad_ST_Internal_applyST__2796778301 = func() gopurs_runtime.Value {
// TAST (Let): Bind1_0_0 -> *Constructor_Control_Bind_Bind[gopurs_runtime.Value]
Bind1_0_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_monadST(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
// TAST (Let): Applicative0_1_1 -> *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_monadST(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_functorST()
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_0_0.V1), f_2, gopurs_runtime.Func(func(f_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_0_0.V1), a_3, gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_1.V1), gopurs_runtime.Apply(f_prime_4, a_prime_5))
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
// TAST (Let): Bind1_0_0 -> *Constructor_Control_Bind_Bind[gopurs_runtime.Value]
Bind1_0_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_monadST(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
// TAST (Let): Applicative0_1_1 -> *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_monadST(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_functorST()
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_0_0.V1), f_2, gopurs_runtime.Func(func(f_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_0_0.V1), a_3, gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_1.V1), gopurs_runtime.Apply(f_prime_4, a_prime_5))
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
// TAST (Let): Bind1_0_0 -> *Constructor_Control_Bind_Bind[gopurs_runtime.Value]
Bind1_0_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_monadST(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
// TAST (Let): Applicative0_1_1 -> *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_monadST(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_functorST()
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_0_0.V1), f_2, gopurs_runtime.Func(func(f_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_0_0.V1), a_3, gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_1.V1), gopurs_runtime.Apply(f_prime_4, a_prime_5))
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
// TAST (Let): Bind1_0_0 -> *Constructor_Control_Bind_Bind[gopurs_runtime.Value]
Bind1_0_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_monadST(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
// TAST (Let): Applicative0_1_1 -> *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_monadST(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_functorST()
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_0_0.V1), f_2, gopurs_runtime.Func(func(f_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_0_0.V1), a_3, gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_1.V1), gopurs_runtime.Apply(f_prime_4, a_prime_5))
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
return Get_Control_Monad_ST_Internal_applyST()
}), Get_Control_Monad_ST_Internal_bind_())
	})
	return cache_Control_Monad_ST_Internal_bindST__2435660861
}

var cache_Control_Monad_ST_Internal_bindST__4187656679 gopurs_runtime.Value
var once_Control_Monad_ST_Internal_bindST__4187656679 sync.Once
func Get_Control_Monad_ST_Internal_bindST__4187656679() gopurs_runtime.Value {
	once_Control_Monad_ST_Internal_bindST__4187656679.Do(func() {
		cache_Control_Monad_ST_Internal_bindST__4187656679 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_applyST()
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
return Get_Control_Monad_ST_Internal_applicativeST()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_ST_Internal_bindST()
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
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_applyST(), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.RecordGet(dictSemigroup_0, "append")
_ = __local_var_2_1
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_applyST(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), __local_var_2_1, a_3), b_4)
})
}))
}

func Call_Control_Monad_ST_Internal_monoidST(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): semigroupST1_1_0 -> gopurs_runtime.Value
semigroupST1_1_0 := Call_Control_Monad_ST_Internal_semigroupST(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = semigroupST1_1_0
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupST1_1_0
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_applicativeST(), "pure"), gopurs_runtime.RecordGet(dictMonoid_0, "mempty")))
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
// TAST (Let): s_prime_2_0 -> *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]
s_prime_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](s_1))}))
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
