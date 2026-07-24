package Data_Monoid

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Ordering "gopurs/output/Data.Ordering"
	pkg_Record_Unsafe "gopurs/output/Record.Unsafe"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
)

var monoidUnit gopurs_runtime.Value
var once_monoidUnit sync.Once
func Get_monoidUnit() gopurs_runtime.Value {
	once_monoidUnit.Do(func() {
		monoidUnit = gopurs_runtime.RecordDict2("mempty", "Semigroup0", pkg_Data_Unit.Get_unit(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Semigroup.Get_semigroupUnit()
}))
	})
	return monoidUnit
}

var monoidString gopurs_runtime.Value
var once_monoidString sync.Once
func Get_monoidString() gopurs_runtime.Value {
	once_monoidString.Do(func() {
		monoidString = gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Str(""), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Semigroup.Get_semigroupString()
}))
	})
	return monoidString
}

var monoidRecordNil gopurs_runtime.Value
var once_monoidRecordNil sync.Once
func Get_monoidRecordNil() gopurs_runtime.Value {
	once_monoidRecordNil.Do(func() {
		monoidRecordNil = gopurs_runtime.RecordDict2("memptyRecord", "SemigroupRecord0", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict0()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Semigroup.Get_semigroupRecordNil()
}))
	})
	return monoidRecordNil
}

var monoidOrdering gopurs_runtime.Value
var once_monoidOrdering sync.Once
func Get_monoidOrdering() gopurs_runtime.Value {
	once_monoidOrdering.Do(func() {
		monoidOrdering = gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Constructor0("EQ"), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ordering.Get_semigroupOrdering()
}))
	})
	return monoidOrdering
}

var monoidArray gopurs_runtime.Value
var once_monoidArray sync.Once
func Get_monoidArray() gopurs_runtime.Value {
	once_monoidArray.Do(func() {
		monoidArray = gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Array([]gopurs_runtime.Value{}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Semigroup.Get_semigroupArray()
}))
	})
	return monoidArray
}

var memptyRecord gopurs_runtime.Value
var once_memptyRecord sync.Once
func Get_memptyRecord() gopurs_runtime.Value {
	once_memptyRecord.Do(func() {
		memptyRecord = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "memptyRecord")
})
	})
	return memptyRecord
}

var monoidRecord gopurs_runtime.Value
var once_monoidRecord sync.Once
func Get_monoidRecord() gopurs_runtime.Value {
	once_monoidRecord.Do(func() {
		monoidRecord = gopurs_runtime.Func2(func(_dollar__unused_0 gopurs_runtime.Value, dictMonoidRecord_1 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupRecord1_2_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoidRecord_1, "SemigroupRecord0"), gopurs_runtime.Value{}), "appendRecord"), gopurs_runtime.Constructor0("Proxy")))
_ = semigroupRecord1_2_0
return gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoidRecord_1, "memptyRecord"), gopurs_runtime.Constructor0("Proxy")), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupRecord1_2_0
}))
})
	})
	return monoidRecord
}

var mempty gopurs_runtime.Value
var once_mempty sync.Once
func Get_mempty() gopurs_runtime.Value {
	once_mempty.Do(func() {
		mempty = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "mempty")
})
	})
	return mempty
}

var monoidFn gopurs_runtime.Value
var once_monoidFn sync.Once
func Get_monoidFn() gopurs_runtime.Value {
	once_monoidFn.Do(func() {
		monoidFn = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty1_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty1_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_1
semigroupFn_3_2 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func3(func(f_3 gopurs_runtime.Value, g_4 gopurs_runtime.Value, x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), gopurs_runtime.Apply(f_3, x_5), gopurs_runtime.Apply(g_4, x_5))
}))
_ = semigroupFn_3_2
return gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return mempty1_1_0
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupFn_3_2
}))
})
	})
	return monoidFn
}

var monoidRecordCons gopurs_runtime.Value
var once_monoidRecordCons sync.Once
func Get_monoidRecordCons() gopurs_runtime.Value {
	once_monoidRecordCons.Do(func() {
		monoidRecordCons = gopurs_runtime.Func2(func(dictIsSymbol_0 gopurs_runtime.Value, dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
mempty1_2_0 := gopurs_runtime.RecordGet(dictMonoid_1, "mempty")
_ = mempty1_2_0
Semigroup0_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{})
_ = Semigroup0_3_1
return gopurs_runtime.Func2(func(_dollar__unused_4 gopurs_runtime.Value, dictMonoidRecord_5 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupRecordCons1_6_2 := gopurs_runtime.Apply4(pkg_Data_Semigroup.Get_semigroupRecordCons(), dictIsSymbol_0, gopurs_runtime.Value{}, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoidRecord_5, "SemigroupRecord0"), gopurs_runtime.Value{}), Semigroup0_3_1)
_ = semigroupRecordCons1_6_2
return gopurs_runtime.RecordDict2("memptyRecord", "SemigroupRecord0", gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(pkg_Record_Unsafe.Get_unsafeSet(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Constructor0("Proxy")), mempty1_2_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoidRecord_5, "memptyRecord"), gopurs_runtime.Constructor0("Proxy")))
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupRecordCons1_6_2
}))
})
})
	})
	return monoidRecordCons
}

var power gopurs_runtime.Value
var once_power sync.Once
func Get_power() gopurs_runtime.Value {
	once_power.Do(func() {
		power = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty1_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty1_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__4_2 gopurs_runtime.Value
_ = go__4_2
go__4_2 = gopurs_runtime.Func(func(p_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if p_5.IntVal <= 0 {
__t4 = mempty1_1_0
goto end_branch_4
} else {

}
}
{
if p_5.IntVal == 1 {
__t4 = x_3
goto end_branch_4
} else {

}
}
{
if gopurs_runtime.Apply2(pkg_Data_EuclideanRing.Get_intMod(), p_5, gopurs_runtime.Int(2)).IntVal == 0 {
x_prime_6_5 := gopurs_runtime.Apply(go__4_2, gopurs_runtime.Int(p_5.IntVal / 2))
_ = x_prime_6_5
__t4 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), x_prime_6_5, x_prime_6_5)
goto end_branch_4
} else {

}
}
{
x_prime_6_3 := gopurs_runtime.Apply(go__4_2, gopurs_runtime.Int(p_5.IntVal / 2))
_ = x_prime_6_3
__t4 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), x_prime_6_3, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), x_prime_6_3, x_3))
}
end_branch_4:
return __t4
})
return go__4_2
})
})
	})
	return power
}

var guard gopurs_runtime.Value
var once_guard sync.Once
func Get_guard() gopurs_runtime.Value {
	once_guard.Do(func() {
		guard = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty1_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty1_1_0
return gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if v_2.IntVal != 0 {
__t1 = v1_3
goto end_branch_1
} else {

}
}
{
__t1 = mempty1_1_0
}
end_branch_1:
return __t1
})
})
	})
	return guard
}




