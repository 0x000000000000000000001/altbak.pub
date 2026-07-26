package Control_Monad_ST_Internal

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Control_Monad_Rec_Class "gopurs/output/Control.Monad.Rec.Class"
)

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

var cache_void gopurs_runtime.Value
var once_void sync.Once
func Get_void() gopurs_runtime.Value {
	once_void.Do(func() {
		cache_void = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_functorST(), "map"), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}))
	})
	return cache_void
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
__local_var_0_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadST(), "Bind1"), gopurs_runtime.Value{})
_ = __local_var_0_0
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorST()
}), gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_0_0, "bind"), f_1, gopurs_runtime.Func(func(f_prime_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_0_0, "bind"), a_2, gopurs_runtime.Func(func(a_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadST(), "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Apply(f_prime_3, a_prime_4))
}))
}))
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

var cache_lift2 gopurs_runtime.Value
var once_lift2 sync.Once
func Get_lift2() gopurs_runtime.Value {
	once_lift2.Do(func() {
		cache_lift2 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lift2(f_0_box, a_1_box, b_2_box)
})
	})
	return cache_lift2
}

var cache_discard gopurs_runtime.Value
var once_discard sync.Once
func Get_discard() gopurs_runtime.Value {
	once_discard.Do(func() {
		cache_discard = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), Get_bindST())
	})
	return cache_discard
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
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_bindST(), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_bindST(), "bind"), gopurs_runtime.Apply(f_0, a_1), Get_new_()), gopurs_runtime.Func(func(r_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_discard(), gopurs_runtime.Apply2(Get_while(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorST(), "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((v_3.Type == 9 && v_3.IntVal == 525585346))
}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return *(r_2.PtrVal().(*gopurs_runtime.Value))
})), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return *(r_2.PtrVal().(*gopurs_runtime.Value))
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 525585346) {
__t0 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_bindST(), "bind"), gopurs_runtime.Apply(f_0, (*pkg_Control_Monad_Rec_Class.Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), gopurs_runtime.Func(func(e_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_void(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
*(r_2.PtrVal().(*gopurs_runtime.Value)) = e_4
return e_4
}))
}))
goto end_branch_0
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 60402430) {
__t0 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applicativeST(), "pure"), pkg_Data_Unit.Get_unit())
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}))), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorST(), "map"), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 60402430) {
__t1 = (*pkg_Control_Monad_Rec_Class.Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return *(r_2.PtrVal().(*gopurs_runtime.Value))
}))
}))
}))
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

func Call_modify(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
return gopurs_runtime.Apply(Get_modifyImpl(), gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
s_prime_2_0 := gopurs_runtime.Apply(f_0, s_1)
_ = s_prime_2_0
return gopurs_runtime.RecordDict2("state", "value", s_prime_2_0, s_prime_2_0)
}))
}

func Call_lift2(f_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_applyST(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applyST(), "Functor0"), gopurs_runtime.Value{}), "map"), f_0, a_1), b_2)
}

func Call_semigroupST(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Apply(Get_lift2(), ((*gopurs_runtime.RecordData1)(dictSemigroup_0.UnsafePtr)).V0))
}

func Call_monoidST(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Apply(Get_lift2(), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}), "append")))
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applicativeST(), "pure"), ((*gopurs_runtime.RecordData1)(dictMonoid_0.UnsafePtr)).V0))
}

func Get_bind_() gopurs_runtime.Value {
	return _Gopurs_Bind_
}

func Get_for_() gopurs_runtime.Value {
	return _Gopurs_For_
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

func Get_new_() gopurs_runtime.Value {
	return _Gopurs_New_
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
