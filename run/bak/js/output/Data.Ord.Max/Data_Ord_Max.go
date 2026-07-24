package Data_Ord_Max

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var Max gopurs_runtime.Value
var once_Max sync.Once
func Get_Max() gopurs_runtime.Value {
	once_Max.Do(func() {
		Max = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Max
}

var showMax gopurs_runtime.Value
var once_showMax sync.Once
func Get_showMax() gopurs_runtime.Value {
	once_showMax.Do(func() {
		showMax = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("(Max " + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal + ")")
}))
})
	})
	return showMax
}

var semigroupMax gopurs_runtime.Value
var once_semigroupMax sync.Once
func Get_semigroupMax() gopurs_runtime.Value {
	once_semigroupMax.Do(func() {
		semigroupMax = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), v_1, v1_2)
_ = v_3_0
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_3_0.StrVal == "LT").IntVal != 0 {
__t1 = v1_2
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v_3_0.StrVal == "EQ").IntVal != 0 {
__t1 = v_1
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v_3_0.StrVal == "GT").IntVal != 0 {
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
}))
})
	})
	return semigroupMax
}

var newtypeMax gopurs_runtime.Value
var once_newtypeMax sync.Once
func Get_newtypeMax() gopurs_runtime.Value {
	once_newtypeMax.Do(func() {
		newtypeMax = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return newtypeMax
}

var monoidMax gopurs_runtime.Value
var once_monoidMax sync.Once
func Get_monoidMax() gopurs_runtime.Value {
	once_monoidMax.Do(func() {
		monoidMax = gopurs_runtime.Func(func(dictBounded_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBounded_0, "Ord0"), gopurs_runtime.Value{})
_ = __local_var_1_0
semigroupMax1_2_1 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
v_4_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "compare"), v_2, v1_3)
_ = v_4_2
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_4_2.StrVal == "LT").IntVal != 0 {
__t3 = v1_3
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool(v_4_2.StrVal == "EQ").IntVal != 0 {
__t3 = v_2
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool(v_4_2.StrVal == "GT").IntVal != 0 {
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
}))
_ = semigroupMax1_2_1
return gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.RecordGet(dictBounded_0, "bottom"), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupMax1_2_1
}))
})
	})
	return monoidMax
}

var eqMax gopurs_runtime.Value
var once_eqMax sync.Once
func Get_eqMax() gopurs_runtime.Value {
	once_eqMax.Do(func() {
		eqMax = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictEq_0
})
	})
	return eqMax
}

var ordMax gopurs_runtime.Value
var once_ordMax sync.Once
func Get_ordMax() gopurs_runtime.Value {
	once_ordMax.Do(func() {
		ordMax = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), v_2, v1_3)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_1_0
}))
})
	})
	return ordMax
}




