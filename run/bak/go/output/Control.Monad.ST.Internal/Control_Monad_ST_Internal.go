package Control_Monad_ST_Internal

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Unit "gopurs/output/Data.Unit"
)

var modify_prime gopurs_runtime.Value
var once_modify_prime sync.Once
func Get_modify_prime() gopurs_runtime.Value {
	once_modify_prime.Do(func() {
		modify_prime = Get_modifyImpl()
	})
	return modify_prime
}

var modify gopurs_runtime.Value
var once_modify sync.Once
func Get_modify() gopurs_runtime.Value {
	once_modify.Do(func() {
		modify = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_modifyImpl(), gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
s_prime_2_0 := gopurs_runtime.Apply(f_0, s_1)
_ = s_prime_2_0
return gopurs_runtime.RecordDict2("state", "value", s_prime_2_0, s_prime_2_0)
}))
})
	})
	return modify
}

var functorST gopurs_runtime.Value
var once_functorST sync.Once
func Get_functorST() gopurs_runtime.Value {
	once_functorST.Do(func() {
		functorST = gopurs_runtime.RecordDict1("map", Get_map_())
	})
	return functorST
}

var monadST gopurs_runtime.Value
var once_monadST sync.Once
func Get_monadST() gopurs_runtime.Value {
	once_monadST.Do(func() {
		monadST = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeST()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindST()
}))
	})
	return monadST
}

var bindST gopurs_runtime.Value
var once_bindST sync.Once
func Get_bindST() gopurs_runtime.Value {
	once_bindST.Do(func() {
		bindST = gopurs_runtime.RecordDict2("bind", "Apply0", Get_bind_(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyST()
}))
	})
	return bindST
}

var applyST gopurs_runtime.Value
var once_applyST sync.Once
func Get_applyST() gopurs_runtime.Value {
	once_applyST.Do(func() {
		applyST = gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
f_prime_2_0 := gopurs_runtime.Apply(f_0, gopurs_runtime.Value{})
_ = f_prime_2_0
a_prime_3_1 := gopurs_runtime.Apply(a_1, gopurs_runtime.Value{})
_ = a_prime_3_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applicativeST(), "pure"), gopurs_runtime.Apply(f_prime_2_0, a_prime_3_1)), gopurs_runtime.Value{})
})
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorST()
}))
	})
	return applyST
}

var applicativeST gopurs_runtime.Value
var once_applicativeST sync.Once
func Get_applicativeST() gopurs_runtime.Value {
	once_applicativeST.Do(func() {
		applicativeST = gopurs_runtime.RecordDict2("pure", "Apply0", Get_pure_(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyST()
}))
	})
	return applicativeST
}

var semigroupST gopurs_runtime.Value
var once_semigroupST sync.Once
func Get_semigroupST() gopurs_runtime.Value {
	once_semigroupST.Do(func() {
		semigroupST = gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(a_1 gopurs_runtime.Value, b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Apply(a_1, gopurs_runtime.Value{})
_ = __local_var_3_0
a_prime_4_1 := gopurs_runtime.Apply(b_2, gopurs_runtime.Value{})
_ = a_prime_4_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), __local_var_3_0, a_prime_4_1)
})
}))
})
	})
	return semigroupST
}

var monadRecST gopurs_runtime.Value
var once_monadRecST sync.Once
func Get_monadRecST() gopurs_runtime.Value {
	once_monadRecST.Do(func() {
		monadRecST = gopurs_runtime.RecordDict2("tailRecM", "Monad0", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(f_0, a_1)
_ = __local_var_2_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
_ = __local_var_3_1
__local_ref_3 := __local_var_3_1
_ = __local_ref_3
r_4_2 := gopurs_runtime.Value{PtrVal: &__local_ref_3}
_ = r_4_2
_dollar__unused_5_4 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_while(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_5 := *(r_4_2.PtrVal.(*gopurs_runtime.Value))
_ = __local_var_5_5
return gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_5_5, "_tag").StrVal == "Loop")
}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
v_5_6 := *(r_4_2.PtrVal.(*gopurs_runtime.Value))
_ = v_5_6
var __t7 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_5_6, "_tag").StrVal == "Loop")).IntVal != 0 {
__t7 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
e_6_8 := gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, gopurs_runtime.RecordGet(v_5_6, "value0")), gopurs_runtime.Value{})
_ = e_6_8
*(r_4_2.PtrVal.(*gopurs_runtime.Value)) = e_6_8
__local_var_7_9 := e_6_8
_ = __local_var_7_9
return pkg_Data_Unit.Get_unit()
})
goto end_branch_7
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_5_6, "_tag").StrVal == "Done")).IntVal != 0 {
__t7 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
})
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
return gopurs_runtime.Apply(__t7, gopurs_runtime.Value{})
})), gopurs_runtime.Value{})
_ = _dollar__unused_5_4
__local_var_6_10 := *(r_4_2.PtrVal.(*gopurs_runtime.Value))
_ = __local_var_6_10
var __t11 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_6_10, "_tag").StrVal == "Done")).IntVal != 0 {
__t11 = gopurs_runtime.RecordGet(__local_var_6_10, "value0")
goto end_branch_11
} else {

}
}
{
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_11:
return __t11
})
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_monadST()
}))
	})
	return monadRecST
}

var monoidST gopurs_runtime.Value
var once_monoidST sync.Once
func Get_monoidST() gopurs_runtime.Value {
	once_monoidST.Do(func() {
		monoidST = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_1_0
semigroupST1_2_1 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(a_2, gopurs_runtime.Value{})
_ = __local_var_4_2
a_prime_5_3 := gopurs_runtime.Apply(b_3, gopurs_runtime.Value{})
_ = a_prime_5_3
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "append"), __local_var_4_2, a_prime_5_3)
})
}))
_ = semigroupST1_2_1
__local_var_3_4 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = __local_var_3_4
return gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_3_4
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupST1_2_1
}))
})
	})
	return monoidST
}

func Get_bind_() gopurs_runtime.Value {
	return Bind_
}

func Get_for_() gopurs_runtime.Value {
	return For_
}

func Get_foreach() gopurs_runtime.Value {
	return Foreach
}

func Get_map_() gopurs_runtime.Value {
	return Map_
}

func Get_modifyImpl() gopurs_runtime.Value {
	return ModifyImpl
}

func Get_new_() gopurs_runtime.Value {
	return New_
}

func Get_pure_() gopurs_runtime.Value {
	return Pure_
}

func Get_read() gopurs_runtime.Value {
	return Read
}

func Get_run() gopurs_runtime.Value {
	return Run
}

func Get_while() gopurs_runtime.Value {
	return While
}

func Get_write() gopurs_runtime.Value {
	return Write
}
