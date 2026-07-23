package Data_Functor_Coproduct

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var Coproduct gopurs_runtime.Value
var once_Coproduct sync.Once
func Get_Coproduct() gopurs_runtime.Value {
	once_Coproduct.Do(func() {
		Coproduct = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Coproduct
}

var showCoproduct gopurs_runtime.Value
var once_showCoproduct sync.Once
func Get_showCoproduct() gopurs_runtime.Value {
	once_showCoproduct.Do(func() {
		showCoproduct = gopurs_runtime.Func2(func(dictShow_0 gopurs_runtime.Value, dictShow1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_2.StrVal == "Left")).IntVal != 0 {
__t0 = gopurs_runtime.Str(gopurs_runtime.Str(gopurs_runtime.Str("(left ").StrVal + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), gopurs_runtime.ConstructorGet(v_2, 0)).StrVal).StrVal + gopurs_runtime.Str(")").StrVal)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_2.StrVal == "Right")).IntVal != 0 {
__t0 = gopurs_runtime.Str(gopurs_runtime.Str(gopurs_runtime.Str("(right ").StrVal + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), gopurs_runtime.ConstructorGet(v_2, 0)).StrVal).StrVal + gopurs_runtime.Str(")").StrVal)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}))
})
	})
	return showCoproduct
}

var right gopurs_runtime.Value
var once_right sync.Once
func Get_right() gopurs_runtime.Value {
	once_right.Do(func() {
		right = gopurs_runtime.Func(func(ga_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor1("Right", ga_0)
})
	})
	return right
}

var newtypeCoproduct gopurs_runtime.Value
var once_newtypeCoproduct sync.Once
func Get_newtypeCoproduct() gopurs_runtime.Value {
	once_newtypeCoproduct.Do(func() {
		newtypeCoproduct = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return newtypeCoproduct
}

var left gopurs_runtime.Value
var once_left sync.Once
func Get_left() gopurs_runtime.Value {
	once_left.Do(func() {
		left = gopurs_runtime.Func(func(fa_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor1("Left", fa_0)
})
	})
	return left
}

var functorCoproduct gopurs_runtime.Value
var once_functorCoproduct sync.Once
func Get_functorCoproduct() gopurs_runtime.Value {
	once_functorCoproduct.Do(func() {
		functorCoproduct = gopurs_runtime.Func2(func(dictFunctor_0 gopurs_runtime.Value, dictFunctor1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_2)
_ = __local_var_4_0
__local_var_5_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor1_1, "map"), f_2)
_ = __local_var_5_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_3.StrVal == "Left")).IntVal != 0 {
__t2 = gopurs_runtime.Constructor1("Left", gopurs_runtime.Apply(__local_var_4_0, gopurs_runtime.ConstructorGet(v_3, 0)))
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(v_3.StrVal == "Right")).IntVal != 0 {
__t2 = gopurs_runtime.Constructor1("Right", gopurs_runtime.Apply(__local_var_5_1, gopurs_runtime.ConstructorGet(v_3, 0)))
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}))
})
	})
	return functorCoproduct
}

var eq1Coproduct gopurs_runtime.Value
var once_eq1Coproduct sync.Once
func Get_eq1Coproduct() gopurs_runtime.Value {
	once_eq1Coproduct.Do(func() {
		eq1Coproduct = gopurs_runtime.Func2(func(dictEq1_0 gopurs_runtime.Value, dictEq11_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_2 gopurs_runtime.Value) gopurs_runtime.Value {
eq12_3_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), dictEq_2)
_ = eq12_3_0
eq13_4_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEq11_1, "eq1"), dictEq_2)
_ = eq13_4_1
return gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_5.StrVal == "Left")).IntVal != 0 {
__t2 = gopurs_runtime.Bool(gopurs_runtime.Bool(v1_6.StrVal == "Left").IntVal != 0 && gopurs_runtime.Apply2(eq12_3_0, gopurs_runtime.ConstructorGet(v_5, 0), gopurs_runtime.ConstructorGet(v1_6, 0)).IntVal != 0)
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Bool(gopurs_runtime.Bool(v_5.StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.Bool(v1_6.StrVal == "Right").IntVal != 0 && gopurs_runtime.Apply2(eq13_4_1, gopurs_runtime.ConstructorGet(v_5, 0), gopurs_runtime.ConstructorGet(v1_6, 0)).IntVal != 0).IntVal != 0)
}
end_branch_2:
return __t2
})
}))
})
	})
	return eq1Coproduct
}

var eqCoproduct gopurs_runtime.Value
var once_eqCoproduct sync.Once
func Get_eqCoproduct() gopurs_runtime.Value {
	once_eqCoproduct.Do(func() {
		eqCoproduct = gopurs_runtime.Func3(func(dictEq1_0 gopurs_runtime.Value, dictEq11_1 gopurs_runtime.Value, dictEq_2 gopurs_runtime.Value) gopurs_runtime.Value {
eq12_3_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), dictEq_2)
_ = eq12_3_0
eq13_4_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEq11_1, "eq1"), dictEq_2)
_ = eq13_4_1
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_5.StrVal == "Left")).IntVal != 0 {
__t2 = gopurs_runtime.Bool(gopurs_runtime.Bool(v1_6.StrVal == "Left").IntVal != 0 && gopurs_runtime.Apply2(eq12_3_0, gopurs_runtime.ConstructorGet(v_5, 0), gopurs_runtime.ConstructorGet(v1_6, 0)).IntVal != 0)
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Bool(gopurs_runtime.Bool(v_5.StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.Bool(v1_6.StrVal == "Right").IntVal != 0 && gopurs_runtime.Apply2(eq13_4_1, gopurs_runtime.ConstructorGet(v_5, 0), gopurs_runtime.ConstructorGet(v1_6, 0)).IntVal != 0).IntVal != 0)
}
end_branch_2:
return __t2
}))
})
	})
	return eqCoproduct
}

var ord1Coproduct gopurs_runtime.Value
var once_ord1Coproduct sync.Once
func Get_ord1Coproduct() gopurs_runtime.Value {
	once_ord1Coproduct.Do(func() {
		ord1Coproduct = gopurs_runtime.Func(func(dictOrd1_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictOrd11_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd11_2, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_3_1
eq1Coproduct2_4_2 := gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_4 gopurs_runtime.Value) gopurs_runtime.Value {
eq12_5_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "eq1"), dictEq_4)
_ = eq12_5_3
eq13_6_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_1, "eq1"), dictEq_4)
_ = eq13_6_4
return gopurs_runtime.Func2(func(v_7 gopurs_runtime.Value, v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_7.StrVal == "Left")).IntVal != 0 {
__t5 = gopurs_runtime.Bool(gopurs_runtime.Bool(v1_8.StrVal == "Left").IntVal != 0 && gopurs_runtime.Apply2(eq12_5_3, gopurs_runtime.ConstructorGet(v_7, 0), gopurs_runtime.ConstructorGet(v1_8, 0)).IntVal != 0)
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Bool(gopurs_runtime.Bool(v_7.StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.Bool(v1_8.StrVal == "Right").IntVal != 0 && gopurs_runtime.Apply2(eq13_6_4, gopurs_runtime.ConstructorGet(v_7, 0), gopurs_runtime.ConstructorGet(v1_8, 0)).IntVal != 0).IntVal != 0)
}
end_branch_5:
return __t5
})
}))
_ = eq1Coproduct2_4_2
return gopurs_runtime.RecordDict2("compare1", "Eq10", gopurs_runtime.Func(func(dictOrd_5 gopurs_runtime.Value) gopurs_runtime.Value {
compare12_6_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "compare1"), dictOrd_5)
_ = compare12_6_6
compare13_7_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd11_2, "compare1"), dictOrd_5)
_ = compare13_7_7
return gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_8.StrVal == "Left")).IntVal != 0 {
var __t9 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v1_9.StrVal == "Left")).IntVal != 0 {
__t9 = gopurs_runtime.Apply2(compare12_6_6, gopurs_runtime.ConstructorGet(v_8, 0), gopurs_runtime.ConstructorGet(v1_9, 0))
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.Constructor0("LT")
}
end_branch_9:
__t8 = __t9
goto end_branch_8
} else {

}
}
{
if (gopurs_runtime.Bool(v1_9.StrVal == "Left")).IntVal != 0 {
__t8 = gopurs_runtime.Constructor0("GT")
goto end_branch_8
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(v_8.StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool(v1_9.StrVal == "Right").IntVal != 0)).IntVal != 0 {
__t8 = gopurs_runtime.Apply2(compare13_7_7, gopurs_runtime.ConstructorGet(v_8, 0), gopurs_runtime.ConstructorGet(v1_9, 0))
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return __t8
})
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return eq1Coproduct2_4_2
}))
})
})
	})
	return ord1Coproduct
}

var ordCoproduct gopurs_runtime.Value
var once_ordCoproduct sync.Once
func Get_ordCoproduct() gopurs_runtime.Value {
	once_ordCoproduct.Do(func() {
		ordCoproduct = gopurs_runtime.Func(func(dictOrd1_0 gopurs_runtime.Value) gopurs_runtime.Value {
ord1Coproduct1_1_0 := gopurs_runtime.Apply(Get_ord1Coproduct(), dictOrd1_0)
_ = ord1Coproduct1_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func(func(dictOrd11_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd11_3, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_4_2
return gopurs_runtime.Func(func(dictOrd_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_5, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_6_3
eq12_7_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "eq1"), __local_var_6_3)
_ = eq12_7_5
eq13_8_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_2, "eq1"), __local_var_6_3)
_ = eq13_8_6
eqCoproduct3_7_4 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(v_9 gopurs_runtime.Value, v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_9.StrVal == "Left")).IntVal != 0 {
__t7 = gopurs_runtime.Bool(gopurs_runtime.Bool(v1_10.StrVal == "Left").IntVal != 0 && gopurs_runtime.Apply2(eq12_7_5, gopurs_runtime.ConstructorGet(v_9, 0), gopurs_runtime.ConstructorGet(v1_10, 0)).IntVal != 0)
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.Bool(gopurs_runtime.Bool(v_9.StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.Bool(v1_10.StrVal == "Right").IntVal != 0 && gopurs_runtime.Apply2(eq13_8_6, gopurs_runtime.ConstructorGet(v_9, 0), gopurs_runtime.ConstructorGet(v1_10, 0)).IntVal != 0).IntVal != 0)
}
end_branch_7:
return __t7
}))
_ = eqCoproduct3_7_4
return gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(ord1Coproduct1_1_0, dictOrd11_3), "compare1"), dictOrd_5), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return eqCoproduct3_7_4
}))
})
})
})
	})
	return ordCoproduct
}

var coproduct gopurs_runtime.Value
var once_coproduct sync.Once
func Get_coproduct() gopurs_runtime.Value {
	once_coproduct.Do(func() {
		coproduct = gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v2_2.StrVal == "Left")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(v_0, gopurs_runtime.ConstructorGet(v2_2, 0))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v2_2.StrVal == "Right")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(v1_1, gopurs_runtime.ConstructorGet(v2_2, 0))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})
	})
	return coproduct
}

var extendCoproduct gopurs_runtime.Value
var once_extendCoproduct sync.Once
func Get_extendCoproduct() gopurs_runtime.Value {
	once_extendCoproduct.Do(func() {
		extendCoproduct = gopurs_runtime.Func(func(dictExtend_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictExtend1_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend1_2, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_1
functorCoproduct2_4_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "map"), f_4)
_ = __local_var_6_3
__local_var_7_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_1, "map"), f_4)
_ = __local_var_7_4
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_5.StrVal == "Left")).IntVal != 0 {
__t5 = gopurs_runtime.Constructor1("Left", gopurs_runtime.Apply(__local_var_6_3, gopurs_runtime.ConstructorGet(v_5, 0)))
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(v_5.StrVal == "Right")).IntVal != 0 {
__t5 = gopurs_runtime.Constructor1("Right", gopurs_runtime.Apply(__local_var_7_4, gopurs_runtime.ConstructorGet(v_5, 0)))
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}))
_ = functorCoproduct2_4_2
return gopurs_runtime.RecordDict2("extend", "Functor0", gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend_0, "extend"), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_5, gopurs_runtime.Constructor1("Left", x_6))
}))
_ = __local_var_6_6
__local_var_7_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend1_2, "extend"), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_5, gopurs_runtime.Constructor1("Right", x_7))
}))
_ = __local_var_7_7
return gopurs_runtime.Func(func(v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v2_8.StrVal == "Left")).IntVal != 0 {
__t8 = gopurs_runtime.Constructor1("Left", gopurs_runtime.Apply(__local_var_6_6, gopurs_runtime.ConstructorGet(v2_8, 0)))
goto end_branch_8
} else {

}
}
{
if (gopurs_runtime.Bool(v2_8.StrVal == "Right")).IntVal != 0 {
__t8 = gopurs_runtime.Constructor1("Right", gopurs_runtime.Apply(__local_var_7_7, gopurs_runtime.ConstructorGet(v2_8, 0)))
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return __t8
})
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCoproduct2_4_2
}))
})
})
	})
	return extendCoproduct
}

var comonadCoproduct gopurs_runtime.Value
var once_comonadCoproduct sync.Once
func Get_comonadCoproduct() gopurs_runtime.Value {
	once_comonadCoproduct.Do(func() {
		comonadCoproduct = gopurs_runtime.Func(func(dictComonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
extendCoproduct1_1_0 := gopurs_runtime.Apply(Get_extendCoproduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "Extend0"), gopurs_runtime.Value{}))
_ = extendCoproduct1_1_0
return gopurs_runtime.Func(func(dictComonad1_2 gopurs_runtime.Value) gopurs_runtime.Value {
extendCoproduct2_3_1 := gopurs_runtime.Apply(extendCoproduct1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad1_2, "Extend0"), gopurs_runtime.Value{}))
_ = extendCoproduct2_3_1
return gopurs_runtime.RecordDict2("extract", "Extend0", gopurs_runtime.Func(func(v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v2_4.StrVal == "Left")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "extract"), gopurs_runtime.ConstructorGet(v2_4, 0))
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(v2_4.StrVal == "Right")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad1_2, "extract"), gopurs_runtime.ConstructorGet(v2_4, 0))
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return extendCoproduct2_3_1
}))
})
})
	})
	return comonadCoproduct
}

var bihoistCoproduct gopurs_runtime.Value
var once_bihoistCoproduct sync.Once
func Get_bihoistCoproduct() gopurs_runtime.Value {
	once_bihoistCoproduct.Do(func() {
		bihoistCoproduct = gopurs_runtime.Func3(func(natF_0 gopurs_runtime.Value, natG_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_2.StrVal == "Left")).IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Left", gopurs_runtime.Apply(natF_0, gopurs_runtime.ConstructorGet(v_2, 0)))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_2.StrVal == "Right")).IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Right", gopurs_runtime.Apply(natG_1, gopurs_runtime.ConstructorGet(v_2, 0)))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})
	})
	return bihoistCoproduct
}


