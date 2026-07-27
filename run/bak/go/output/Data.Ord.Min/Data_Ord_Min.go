package Data_Ord_Min

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
)

var cache_Min gopurs_runtime.Value
var once_Min sync.Once
func Get_Min() gopurs_runtime.Value {
	once_Min.Do(func() {
		cache_Min = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Min(x_0_box)
})
	})
	return cache_Min
}

var cache_showMin gopurs_runtime.Value
var once_showMin sync.Once
func Get_showMin() gopurs_runtime.Value {
	once_showMin.Do(func() {
		cache_showMin = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_showMin(dictShow_0_box))
})
	})
	return cache_showMin
}

var cache_semigroupMin gopurs_runtime.Value
var once_semigroupMin sync.Once
func Get_semigroupMin() gopurs_runtime.Value {
	once_semigroupMin.Do(func() {
		cache_semigroupMin = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_semigroupMin(dictOrd_0_box))
})
	})
	return cache_semigroupMin
}

var cache_newtypeMin gopurs_runtime.Value
var once_newtypeMin sync.Once
func Get_newtypeMin() gopurs_runtime.Value {
	once_newtypeMin.Do(func() {
		cache_newtypeMin = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))))
	})
	return cache_newtypeMin
}

var cache_monoidMin gopurs_runtime.Value
var once_monoidMin sync.Once
func Get_monoidMin() gopurs_runtime.Value {
	once_monoidMin.Do(func() {
		cache_monoidMin = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_monoidMin(dictBounded_0_box))
})
	})
	return cache_monoidMin
}

var cache_eqMin gopurs_runtime.Value
var once_eqMin sync.Once
func Get_eqMin() gopurs_runtime.Value {
	once_eqMin.Do(func() {
		cache_eqMin = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_eqMin(dictEq_0_box))
})
	})
	return cache_eqMin
}

var cache_ordMin gopurs_runtime.Value
var once_ordMin sync.Once
func Get_ordMin() gopurs_runtime.Value {
	once_ordMin.Do(func() {
		cache_ordMin = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_ordMin(dictOrd_0_box))
})
	})
	return cache_ordMin
}

func Call_Min(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_showMin(dictShow_0_loop gopurs_runtime.Value) interface{} {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(Min "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1), gopurs_runtime.Str(")")))
})))
}

func Call_semigroupMin(dictOrd_0_loop gopurs_runtime.Value) interface{} {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), v_1, v1_2)
_ = v_3_0
var __t1 gopurs_runtime.Value
{
if (v_3_0.Type == 9 && v_3_0.IntVal == 1527465420) {
__t1 = v_1
goto end_branch_1
} else {

}
}
{
if (v_3_0.Type == 9 && v_3_0.IntVal == 902936544) {
__t1 = v_1
goto end_branch_1
} else {

}
}
{
if (v_3_0.Type == 9 && v_3_0.IntVal == 380165415) {
__t1 = v1_2
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
})))
}

func Call_monoidMin(dictBounded_0_loop gopurs_runtime.Value) interface{} {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBounded_0, "Ord0"), gopurs_runtime.Value{})
_ = __local_var_1_0
semigroupMin1_2_1 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
v_4_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "compare"), v_2, v1_3)
_ = v_4_2
var __t3 gopurs_runtime.Value
{
if (v_4_2.Type == 9 && v_4_2.IntVal == 1527465420) {
__t3 = v_2
goto end_branch_3
} else {

}
}
{
if (v_4_2.Type == 9 && v_4_2.IntVal == 902936544) {
__t3 = v_2
goto end_branch_3
} else {

}
}
{
if (v_4_2.Type == 9 && v_4_2.IntVal == 380165415) {
__t3 = v1_3
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}))
_ = semigroupMin1_2_1
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupMin1_2_1
}), gopurs_runtime.RecordGet(dictBounded_0, "top")))
}

func Call_eqMin(dictEq_0_loop gopurs_runtime.Value) interface{} {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.UnboxAny(dictEq_0)
}

func Call_ordMin(dictOrd_0_loop gopurs_runtime.Value) interface{} {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_1_0
}), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), v_2, v1_3)
})))
}
