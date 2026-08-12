package Data_Ord_Max

import (
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Max gopurs_runtime.Value
var once_Max sync.Once
func Get_Max() gopurs_runtime.Value {
	once_Max.Do(func() {
		cache_Max = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Max(x_0_box)
})
	})
	return cache_Max
}

var cache_showMax gopurs_runtime.Value
var once_showMax sync.Once
func Get_showMax() gopurs_runtime.Value {
	once_showMax.Do(func() {
		cache_showMax = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showMax(dictShow_0_box)
})
	})
	return cache_showMax
}

var cache_semigroupMax gopurs_runtime.Value
var once_semigroupMax sync.Once
func Get_semigroupMax() gopurs_runtime.Value {
	once_semigroupMax.Do(func() {
		cache_semigroupMax = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semigroupMax(dictOrd_0_box)
})
	})
	return cache_semigroupMax
}

var cache_newtypeMax gopurs_runtime.Value
var once_newtypeMax sync.Once
func Get_newtypeMax() gopurs_runtime.Value {
	once_newtypeMax.Do(func() {
		cache_newtypeMax = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeMax
}

var cache_monoidMax gopurs_runtime.Value
var once_monoidMax sync.Once
func Get_monoidMax() gopurs_runtime.Value {
	once_monoidMax.Do(func() {
		cache_monoidMax = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monoidMax(dictBounded_0_box)
})
	})
	return cache_monoidMax
}

var cache_eqMax gopurs_runtime.Value
var once_eqMax sync.Once
func Get_eqMax() gopurs_runtime.Value {
	once_eqMax.Do(func() {
		cache_eqMax = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqMax(dictEq_0_box)
})
	})
	return cache_eqMax
}

var cache_ordMax gopurs_runtime.Value
var once_ordMax sync.Once
func Get_ordMax() gopurs_runtime.Value {
	once_ordMax.Do(func() {
		cache_ordMax = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordMax(dictOrd_0_box)
})
	})
	return cache_ordMax
}

func Call_Max(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_showMax(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(Max "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1), gopurs_runtime.Str(")"))).StrVal())
}))
}

func Call_semigroupMax(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), v_1, v1_2)
_ = v_3_0
var __t1 gopurs_runtime.Value
{
if (uint32(v_3_0.IntVal) == 1527465420) {
__t1 = v1_2
goto end_branch_1
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 902936544) {
__t1 = v_1
goto end_branch_1
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 380165415) {
__t1 = v_1
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
}

func Call_monoidMax(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBounded_0, "Ord0"), gopurs_runtime.Value{})
_ = __local_var_1_1
semigroupMax1_1_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
v_4_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "compare"), v_2, v1_3)
_ = v_4_2
var __t3 gopurs_runtime.Value
{
if (uint32(v_4_2.IntVal) == 1527465420) {
__t3 = v1_3
goto end_branch_3
} else {

}
}
{
if (uint32(v_4_2.IntVal) == 902936544) {
__t3 = v_2
goto end_branch_3
} else {

}
}
{
if (uint32(v_4_2.IntVal) == 380165415) {
__t3 = v_2
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
})
}))
_ = semigroupMax1_1_0
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupMax1_1_0
}), gopurs_runtime.RecordGet(dictBounded_0, "bottom"))
}

func Call_eqMax(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return dictEq_0
}

func Call_ordMax(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
eqMax1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = eqMax1_1_0
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return eqMax1_1_0
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), v_2, v1_3).IntVal)), UnsafePtr: nil}
})
}))
}


