package Data_Comparison

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var Comparison gopurs_runtime.Value
var once_Comparison sync.Once
func Get_Comparison() gopurs_runtime.Value {
	once_Comparison.Do(func() {
		Comparison = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Comparison
}

var semigroupComparison gopurs_runtime.Value
var once_semigroupComparison sync.Once
func Get_semigroupComparison() gopurs_runtime.Value {
	once_semigroupComparison.Do(func() {
		semigroupComparison = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Apply(v_0, x_2)
_ = __local_var_3_0
__local_var_4_1 := gopurs_runtime.Apply(v1_1, x_2)
_ = __local_var_4_1
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_2 := gopurs_runtime.Apply(__local_var_3_0, x_5)
_ = __local_var_6_2
__local_var_7_3 := gopurs_runtime.Apply(__local_var_4_1, x_5)
_ = __local_var_7_3
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_6_2, "_tag").StrVal == "LT")).IntVal != 0 {
__t4 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("LT"))
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_6_2, "_tag").StrVal == "GT")).IntVal != 0 {
__t4 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("GT"))
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_6_2, "_tag").StrVal == "EQ")).IntVal != 0 {
__t4 = __local_var_7_3
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
})
}))
	})
	return semigroupComparison
}

var newtypeComparison gopurs_runtime.Value
var once_newtypeComparison sync.Once
func Get_newtypeComparison() gopurs_runtime.Value {
	once_newtypeComparison.Do(func() {
		newtypeComparison = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return newtypeComparison
}

var monoidComparison gopurs_runtime.Value
var once_monoidComparison sync.Once
func Get_monoidComparison() gopurs_runtime.Value {
	once_monoidComparison.Do(func() {
		monoidComparison = gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("EQ"))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupComparison()
}))
	})
	return monoidComparison
}

var defaultComparison gopurs_runtime.Value
var once_defaultComparison sync.Once
func Get_defaultComparison() gopurs_runtime.Value {
	once_defaultComparison.Do(func() {
		defaultComparison = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictOrd_0, "compare")
})
	})
	return defaultComparison
}

var contravariantComparison gopurs_runtime.Value
var once_contravariantComparison sync.Once
func Get_contravariantComparison() gopurs_runtime.Value {
	once_contravariantComparison.Do(func() {
		contravariantComparison = gopurs_runtime.RecordDict1("cmap", gopurs_runtime.Func4(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(v_1, gopurs_runtime.Apply(f_0, x_2), gopurs_runtime.Apply(f_0, y_3))
}))
	})
	return contravariantComparison
}


