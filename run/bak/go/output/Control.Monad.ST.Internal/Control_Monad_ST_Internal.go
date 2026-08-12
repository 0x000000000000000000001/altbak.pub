package Control_Monad_ST_Internal

import (
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Control_Monad_Rec_Class "gopurs/output/Control.Monad.Rec.Class"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Partial_Unsafe "gopurs/output/Partial.Unsafe"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_new gopurs_runtime.Value
var once_new sync.Once
func Get_new() gopurs_runtime.Value {
	once_new.Do(func() {
		cache_new = Get_newImpl()
	})
	return cache_new
}

var cache_modify_prime gopurs_runtime.Value
var once_modify_prime sync.Once
func Get_modify_prime() gopurs_runtime.Value {
	once_modify_prime.Do(func() {
		cache_modify_prime = Get_modifyImpl()
	})
	return cache_modify_prime
}

var cache_modify gopurs_runtime.Value
var once_modify sync.Once
func Get_modify() gopurs_runtime.Value {
	once_modify.Do(func() {
		cache_modify = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_modify(f_0_box)
})
	})
	return cache_modify
}

var cache_functorST gopurs_runtime.Value
var once_functorST sync.Once
func Get_functorST() gopurs_runtime.Value {
	once_functorST.Do(func() {
		cache_functorST = gopurs_runtime.RecordDict1("map", Get_map_())
	})
	return cache_functorST
}

var cache_go__for gopurs_runtime.Value
var once_go__for sync.Once
func Get_go__for() gopurs_runtime.Value {
	once_go__for.Do(func() {
		cache_go__for = Get_forImpl()
	})
	return cache_go__for
}

var cache_monadST gopurs_runtime.Value
var once_monadST sync.Once
func Get_monadST() gopurs_runtime.Value {
	once_monadST.Do(func() {
		cache_monadST = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeST()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindST()
}))
	})
	return cache_monadST
}

var cache_bindST gopurs_runtime.Value
var once_bindST sync.Once
func Get_bindST() gopurs_runtime.Value {
	once_bindST.Do(func() {
		cache_bindST = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyST()
}), Get_bind_())
	})
	return cache_bindST
}

var cache_applyST gopurs_runtime.Value
var once_applyST sync.Once
func Get_applyST() gopurs_runtime.Value {
	once_applyST.Do(func() {
		cache_applyST = func() gopurs_runtime.Value {
Bind1_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadST(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadST(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorST()
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
	return cache_applyST
}

var cache_applicativeST gopurs_runtime.Value
var once_applicativeST sync.Once
func Get_applicativeST() gopurs_runtime.Value {
	once_applicativeST.Do(func() {
		cache_applicativeST = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyST()
}), Get_pure_())
	})
	return cache_applicativeST
}

var cache_semigroupST gopurs_runtime.Value
var once_semigroupST sync.Once
func Get_semigroupST() gopurs_runtime.Value {
	once_semigroupST.Do(func() {
		cache_semigroupST = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semigroupST(dictSemigroup_0_box)
})
	})
	return cache_semigroupST
}

var cache_monadRecST gopurs_runtime.Value
var once_monadRecST sync.Once
func Get_monadRecST() gopurs_runtime.Value {
	once_monadRecST.Do(func() {
		cache_monadRecST = gopurs_runtime.RecordDict2("Monad0", "tailRecM", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_monadST()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_bindST(), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_bindST(), "bind"), gopurs_runtime.Apply(f_0, a_1), Get_newImpl()), gopurs_runtime.Func(func(r_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), Get_bindST(), gopurs_runtime.Apply2(Get_while(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorST(), "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 bool
{
if (v_3.Type == 9 && v_3.IntVal == 525585346) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
return gopurs_runtime.Bool(__t0)
}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return (*(r_2.PtrVal().(*interface{}))).(gopurs_runtime.Value)
})), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return (*(r_2.PtrVal().(*interface{}))).(gopurs_runtime.Value)
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 525585346) {
__t1 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_bindST(), "bind"), gopurs_runtime.Apply(f_0, (*pkg_Control_Monad_Rec_Class.Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), gopurs_runtime.Func(func(e_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorST(), "map"), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
*(r_2.PtrVal().(*interface{})) = e_4
return e_4
}))
}))
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 60402430) {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applicativeST(), "pure"), pkg_Data_Unit.Get_unit())
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}))), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorST(), "map"), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 60402430) {
__t2 = (*pkg_Control_Monad_Rec_Class.Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return (*(r_2.PtrVal().(*interface{}))).(gopurs_runtime.Value)
}))
}))
}))
})
}))
	})
	return cache_monadRecST
}

var cache_monoidST gopurs_runtime.Value
var once_monoidST sync.Once
func Get_monoidST() gopurs_runtime.Value {
	once_monoidST.Do(func() {
		cache_monoidST = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monoidST(dictMonoid_0_box)
})
	})
	return cache_monoidST
}

var cache_pure__3215807376 gopurs_runtime.Value
var once_pure__3215807376 sync.Once
func Get_pure__3215807376() gopurs_runtime.Value {
	once_pure__3215807376.Do(func() {
		cache_pure__3215807376 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__3215807376(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__3215807376
}

var cache_pure__3079134646 gopurs_runtime.Value
var once_pure__3079134646 sync.Once
func Get_pure__3079134646() gopurs_runtime.Value {
	once_pure__3079134646.Do(func() {
		cache_pure__3079134646 = gopurs_runtime.RecordGet(Get_applicativeST(), "pure")
	})
	return cache_pure__3079134646
}

var cache_apply__353515660 gopurs_runtime.Value
var once_apply__353515660 sync.Once
func Get_apply__353515660() gopurs_runtime.Value {
	once_apply__353515660.Do(func() {
		cache_apply__353515660 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__353515660(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_apply__353515660
}

var cache_lift2__2762258480 gopurs_runtime.Value
var once_lift2__2762258480 sync.Once
func Get_lift2__2762258480() gopurs_runtime.Value {
	once_lift2__2762258480.Do(func() {
		cache_lift2__2762258480 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lift2__2762258480(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dictApply_0_box))
})
	})
	return cache_lift2__2762258480
}

var cache_lift2__650234614 gopurs_runtime.Value
var once_lift2__650234614 sync.Once
func Get_lift2__650234614() gopurs_runtime.Value {
	once_lift2__650234614.Do(func() {
		cache_lift2__650234614 = func() gopurs_runtime.Value {
Functor0_0_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applyST(), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_0_0
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_applyST(), "apply"), gopurs_runtime.Apply2(Functor0_0_0.V0, f_1, a_2), b_3)
})
})
})
}()
	})
	return cache_lift2__650234614
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

var cache_bind__1887881377 gopurs_runtime.Value
var once_bind__1887881377 sync.Once
func Get_bind__1887881377() gopurs_runtime.Value {
	once_bind__1887881377.Do(func() {
		cache_bind__1887881377 = gopurs_runtime.RecordGet(Get_bindST(), "bind")
	})
	return cache_bind__1887881377
}

var cache_bind__2981096353 gopurs_runtime.Value
var once_bind__2981096353 sync.Once
func Get_bind__2981096353() gopurs_runtime.Value {
	once_bind__2981096353.Do(func() {
		cache_bind__2981096353 = gopurs_runtime.RecordGet(Get_bindST(), "bind")
	})
	return cache_bind__2981096353
}

var cache_bindFlipped__1485397639 gopurs_runtime.Value
var once_bindFlipped__1485397639 sync.Once
func Get_bindFlipped__1485397639() gopurs_runtime.Value {
	once_bindFlipped__1485397639.Do(func() {
		cache_bindFlipped__1485397639 = gopurs_runtime.Func3(func(dictBind_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindFlipped__1485397639(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dictBind_0_box), b_1_box, a_2_box)
})
	})
	return cache_bindFlipped__1485397639
}

var cache_bindFlipped__3235594689 gopurs_runtime.Value
var once_bindFlipped__3235594689 sync.Once
func Get_bindFlipped__3235594689() gopurs_runtime.Value {
	once_bindFlipped__3235594689.Do(func() {
		cache_bindFlipped__3235594689 = gopurs_runtime.Func2(func(b_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindFlipped__3235594689(b_0_box, a_1_box)
})
	})
	return cache_bindFlipped__3235594689
}

var cache_discard__1876171936 gopurs_runtime.Value
var once_discard__1876171936 sync.Once
func Get_discard__1876171936() gopurs_runtime.Value {
	once_discard__1876171936.Do(func() {
		cache_discard__1876171936 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), Get_bindST())
	})
	return cache_discard__1876171936
}

var cache_discard__317162198 gopurs_runtime.Value
var once_discard__317162198 sync.Once
func Get_discard__317162198() gopurs_runtime.Value {
	once_discard__317162198.Do(func() {
		cache_discard__317162198 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_discard__317162198(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Discard[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_discard__317162198
}

var cache_discardUnit__2687062302 gopurs_runtime.Value
var once_discardUnit__2687062302 sync.Once
func Get_discardUnit__2687062302() gopurs_runtime.Value {
	once_discardUnit__2687062302.Do(func() {
		cache_discardUnit__2687062302 = gopurs_runtime.RecordDict1("discard", gopurs_runtime.Func(func(dictBind_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictBind_0, "bind")
}))
	})
	return cache_discardUnit__2687062302
}

var cache_applicativeST__3091537981 gopurs_runtime.Value
var once_applicativeST__3091537981 sync.Once
func Get_applicativeST__3091537981() gopurs_runtime.Value {
	once_applicativeST__3091537981.Do(func() {
		cache_applicativeST__3091537981 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyST()
}), Get_pure_())
	})
	return cache_applicativeST__3091537981
}

var cache_applicativeST__2868811880 gopurs_runtime.Value
var once_applicativeST__2868811880 sync.Once
func Get_applicativeST__2868811880() gopurs_runtime.Value {
	once_applicativeST__2868811880.Do(func() {
		cache_applicativeST__2868811880 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyST()
}), Get_pure_())
	})
	return cache_applicativeST__2868811880
}

var cache_applyST__2796778301 gopurs_runtime.Value
var once_applyST__2796778301 sync.Once
func Get_applyST__2796778301() gopurs_runtime.Value {
	once_applyST__2796778301.Do(func() {
		cache_applyST__2796778301 = func() gopurs_runtime.Value {
Bind1_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadST(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadST(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorST()
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
	return cache_applyST__2796778301
}

var cache_applyST__2741064779 gopurs_runtime.Value
var once_applyST__2741064779 sync.Once
func Get_applyST__2741064779() gopurs_runtime.Value {
	once_applyST__2741064779.Do(func() {
		cache_applyST__2741064779 = func() gopurs_runtime.Value {
Bind1_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadST(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadST(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorST()
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
	return cache_applyST__2741064779
}

var cache_bindST__2435660861 gopurs_runtime.Value
var once_bindST__2435660861 sync.Once
func Get_bindST__2435660861() gopurs_runtime.Value {
	once_bindST__2435660861.Do(func() {
		cache_bindST__2435660861 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyST()
}), Get_bind_())
	})
	return cache_bindST__2435660861
}

var cache_bindST__4187656679 gopurs_runtime.Value
var once_bindST__4187656679 sync.Once
func Get_bindST__4187656679() gopurs_runtime.Value {
	once_bindST__4187656679.Do(func() {
		cache_bindST__4187656679 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyST()
}), Get_bind_())
	})
	return cache_bindST__4187656679
}

var cache_functorST__4062753802 gopurs_runtime.Value
var once_functorST__4062753802 sync.Once
func Get_functorST__4062753802() gopurs_runtime.Value {
	once_functorST__4062753802.Do(func() {
		cache_functorST__4062753802 = gopurs_runtime.RecordDict1("map", Get_map_())
	})
	return cache_functorST__4062753802
}

var cache_functorST__2441840241 gopurs_runtime.Value
var once_functorST__2441840241 sync.Once
func Get_functorST__2441840241() gopurs_runtime.Value {
	once_functorST__2441840241.Do(func() {
		cache_functorST__2441840241 = gopurs_runtime.RecordDict1("map", Get_map_())
	})
	return cache_functorST__2441840241
}

var cache_modify_prime__1497736571 gopurs_runtime.Value
var once_modify_prime__1497736571 sync.Once
func Get_modify_prime__1497736571() gopurs_runtime.Value {
	once_modify_prime__1497736571.Do(func() {
		cache_modify_prime__1497736571 = Get_modifyImpl()
	})
	return cache_modify_prime__1497736571
}

var cache_monadST__1413783571 gopurs_runtime.Value
var once_monadST__1413783571 sync.Once
func Get_monadST__1413783571() gopurs_runtime.Value {
	once_monadST__1413783571.Do(func() {
		cache_monadST__1413783571 = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeST()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindST()
}))
	})
	return cache_monadST__1413783571
}

var cache_new__3489595018 gopurs_runtime.Value
var once_new__3489595018 sync.Once
func Get_new__3489595018() gopurs_runtime.Value {
	once_new__3489595018.Do(func() {
		cache_new__3489595018 = Get_newImpl()
	})
	return cache_new__3489595018
}

var cache_const__4026847508 gopurs_runtime.Value
var once_const__4026847508 sync.Once
func Get_const__4026847508() gopurs_runtime.Value {
	once_const__4026847508.Do(func() {
		cache_const__4026847508 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__4026847508(a_0_box, v_1_box)
})
	})
	return cache_const__4026847508
}

var cache_const__641934996 gopurs_runtime.Value
var once_const__641934996 sync.Once
func Get_const__641934996() gopurs_runtime.Value {
	once_const__641934996.Do(func() {
		cache_const__641934996 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__641934996(a_0_box, v_1_box)
})
	})
	return cache_const__641934996
}

var cache_flip__3525081280 gopurs_runtime.Value
var once_flip__3525081280 sync.Once
func Get_flip__3525081280() gopurs_runtime.Value {
	once_flip__3525081280.Do(func() {
		cache_flip__3525081280 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__3525081280(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__3525081280
}

var cache_flip__2253242624 gopurs_runtime.Value
var once_flip__2253242624 sync.Once
func Get_flip__2253242624() gopurs_runtime.Value {
	once_flip__2253242624.Do(func() {
		cache_flip__2253242624 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__2253242624(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__2253242624
}

var cache_map__3240628980 gopurs_runtime.Value
var once_map__3240628980 sync.Once
func Get_map__3240628980() gopurs_runtime.Value {
	once_map__3240628980.Do(func() {
		cache_map__3240628980 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__3240628980(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__3240628980
}

var cache_map__2199395572 gopurs_runtime.Value
var once_map__2199395572 sync.Once
func Get_map__2199395572() gopurs_runtime.Value {
	once_map__2199395572.Do(func() {
		cache_map__2199395572 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2199395572(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2199395572
}

var cache_map__1483545076 gopurs_runtime.Value
var once_map__1483545076 sync.Once
func Get_map__1483545076() gopurs_runtime.Value {
	once_map__1483545076.Do(func() {
		cache_map__1483545076 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__1483545076(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__1483545076
}

var cache_map__2174973445 gopurs_runtime.Value
var once_map__2174973445 sync.Once
func Get_map__2174973445() gopurs_runtime.Value {
	once_map__2174973445.Do(func() {
		cache_map__2174973445 = gopurs_runtime.RecordGet(Get_functorST(), "map")
	})
	return cache_map__2174973445
}

var cache_void__3020373336 gopurs_runtime.Value
var once_void__3020373336 sync.Once
func Get_void__3020373336() gopurs_runtime.Value {
	once_void__3020373336.Do(func() {
		cache_void__3020373336 = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_void__3020373336(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dictFunctor_0_box))
})
	})
	return cache_void__3020373336
}

var cache_void__2104786761 gopurs_runtime.Value
var once_void__2104786761 sync.Once
func Get_void__2104786761() gopurs_runtime.Value {
	once_void__2104786761.Do(func() {
		cache_void__2104786761 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_functorST(), "map"), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}))
	})
	return cache_void__2104786761
}

var cache_mempty__2312420373 gopurs_runtime.Value
var once_mempty__2312420373 sync.Once
func Get_mempty__2312420373() gopurs_runtime.Value {
	once_mempty__2312420373.Do(func() {
		cache_mempty__2312420373 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mempty__2312420373(dict_0_box)
})
	})
	return cache_mempty__2312420373
}

var cache_unsafePartial__1306634845 gopurs_runtime.Value
var once_unsafePartial__1306634845 sync.Once
func Get_unsafePartial__1306634845() gopurs_runtime.Value {
	once_unsafePartial__1306634845.Do(func() {
		cache_unsafePartial__1306634845 = pkg_Partial_Unsafe.Get__unsafePartial()
	})
	return cache_unsafePartial__1306634845
}

var cache_unsafePartial__254560477 gopurs_runtime.Value
var once_unsafePartial__254560477 sync.Once
func Get_unsafePartial__254560477() gopurs_runtime.Value {
	once_unsafePartial__254560477.Do(func() {
		cache_unsafePartial__254560477 = pkg_Partial_Unsafe.Get__unsafePartial()
	})
	return cache_unsafePartial__254560477
}

func Call_modify(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
return gopurs_runtime.Apply(Get_modifyImpl(), gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
s_prime_2_0 := gopurs_runtime.Apply(f_0, s_1)
_ = s_prime_2_0
return gopurs_runtime.RecordDict2("state", "value", s_prime_2_0, s_prime_2_0)
}))
}

func Call_semigroupST(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applyST(), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
__local_var_2_1 := gopurs_runtime.RecordGet(dictSemigroup_0, "append")
_ = __local_var_2_1
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_applyST(), "apply"), gopurs_runtime.Apply2(Functor0_1_0.V0, __local_var_2_1, a_3), b_4)
})
}))
}

func Call_monoidST(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
semigroupST1_1_0 := Call_semigroupST(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = semigroupST1_1_0
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupST1_1_0
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applicativeST(), "pure"), gopurs_runtime.RecordGet(dictMonoid_0, "mempty")))
}

func Call_pure__3215807376(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__353515660(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_lift2__2762258480(dictApply_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApply_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dictApply_0_loop
_ = dictApply_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, f_2, a_3), b_4)
})
})
})
}

func Call_bind__2601835655(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bindFlipped__1485397639(dictBind_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value], b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dictBind_0_loop
_ = dictBind_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(dictBind_0.V1, a_2, b_1)
}

func Call_bindFlipped__3235594689(b_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var b_0 gopurs_runtime.Value = b_0_loop
_ = b_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_bindST(), "bind"), a_1, b_0)
}

func Call_discard__317162198(dict_0_loop *pkg_Control_Bind.Constructor_Discard[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Discard[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_const__4026847508(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__641934996(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_flip__3525081280(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__2253242624(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_map__3240628980(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__2199395572(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__1483545076(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_void__3020373336(dictFunctor_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFunctor_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.Apply(dictFunctor_0.V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}))
}

func Call_mempty__2312420373(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "mempty")
}

func Get_bind_() gopurs_runtime.Value {
	return _Gopurs_Bind_
}

func Get_forImpl() gopurs_runtime.Value {
	return _Gopurs_ForImpl
}

func Get_foreach() gopurs_runtime.Value {
	return _Gopurs_Foreach
}

func Get_map_() gopurs_runtime.Value {
	return _Gopurs_Map_
}

func Get_modifyImpl() gopurs_runtime.Value {
	return _Gopurs_ModifyImpl
}

func Get_newImpl() gopurs_runtime.Value {
	return _Gopurs_NewImpl
}

func Get_pure_() gopurs_runtime.Value {
	return _Gopurs_Pure_
}

func Get_read() gopurs_runtime.Value {
	return _Gopurs_Read
}

func Get_run() gopurs_runtime.Value {
	return _Gopurs_Run
}

func Get_while() gopurs_runtime.Value {
	return _Gopurs_While
}

func Get_write() gopurs_runtime.Value {
	return _Gopurs_Write
}
