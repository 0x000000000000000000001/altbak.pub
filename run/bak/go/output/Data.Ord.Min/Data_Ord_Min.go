package Data_Ord_Min

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var Min gopurs_runtime.Value
var once_Min sync.Once
func Get_Min() gopurs_runtime.Value {
	once_Min.Do(func() {
		Min = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Min
}

var showMin gopurs_runtime.Value
var once_showMin sync.Once
func Get_showMin() gopurs_runtime.Value {
	once_showMin.Do(func() {
		showMin = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Str(gopurs_runtime.Str("(Min ").StrVal + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal).StrVal + gopurs_runtime.Str(")").StrVal)
}))
})
	})
	return showMin
}

var semigroupMin gopurs_runtime.Value
var once_semigroupMin sync.Once
func Get_semigroupMin() gopurs_runtime.Value {
	once_semigroupMin.Do(func() {
		semigroupMin = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), v_1, v1_2)
_ = v_3_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_3_0.StrVal == "LT")).IntVal != 0 {
__t1 = v_1
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_3_0.StrVal == "EQ")).IntVal != 0 {
__t1 = v_1
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_3_0.StrVal == "GT")).IntVal != 0 {
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
}))
})
	})
	return semigroupMin
}

var newtypeMin gopurs_runtime.Value
var once_newtypeMin sync.Once
func Get_newtypeMin() gopurs_runtime.Value {
	once_newtypeMin.Do(func() {
		newtypeMin = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return newtypeMin
}

var monoidMin gopurs_runtime.Value
var once_monoidMin sync.Once
func Get_monoidMin() gopurs_runtime.Value {
	once_monoidMin.Do(func() {
		monoidMin = gopurs_runtime.Func(func(dictBounded_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBounded_0, "Ord0"), gopurs_runtime.Value{})
_ = __local_var_1_0
semigroupMin1_2_1 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
v_4_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "compare"), v_2, v1_3)
_ = v_4_2
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_4_2.StrVal == "LT")).IntVal != 0 {
__t3 = v_2
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(v_4_2.StrVal == "EQ")).IntVal != 0 {
__t3 = v_2
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(v_4_2.StrVal == "GT")).IntVal != 0 {
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
return gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.RecordGet(dictBounded_0, "top"), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupMin1_2_1
}))
})
	})
	return monoidMin
}

var eqMin gopurs_runtime.Value
var once_eqMin sync.Once
func Get_eqMin() gopurs_runtime.Value {
	once_eqMin.Do(func() {
		eqMin = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictEq_0
})
	})
	return eqMin
}

var ordMin gopurs_runtime.Value
var once_ordMin sync.Once
func Get_ordMin() gopurs_runtime.Value {
	once_ordMin.Do(func() {
		ordMin = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), v_2, v1_3)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_1_0
}))
})
	})
	return ordMin
}


