package Data_Functor_Product2

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var Product2 gopurs_runtime.Value
var once_Product2 sync.Once
func Get_Product2() gopurs_runtime.Value {
	once_Product2.Do(func() {
		Product2 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Product2", value0, value1)
})
})
	})
	return Product2
}

var showProduct2 gopurs_runtime.Value
var once_showProduct2 sync.Once
func Get_showProduct2() gopurs_runtime.Value {
	once_showProduct2.Do(func() {
		showProduct2 = gopurs_runtime.Func2(func(dictShow_0 gopurs_runtime.Value, dictShow1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Str(gopurs_runtime.Str(gopurs_runtime.Str(gopurs_runtime.Str("(Product2 ").StrVal + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), gopurs_runtime.ConstructorGet(v_2, 0)).StrVal).StrVal + gopurs_runtime.Str(" ").StrVal).StrVal + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), gopurs_runtime.ConstructorGet(v_2, 1)).StrVal).StrVal + gopurs_runtime.Str(")").StrVal)
}))
})
	})
	return showProduct2
}

var profunctorProduct2 gopurs_runtime.Value
var once_profunctorProduct2 sync.Once
func Get_profunctorProduct2() gopurs_runtime.Value {
	once_profunctorProduct2.Do(func() {
		profunctorProduct2 = gopurs_runtime.Func2(func(dictProfunctor_0 gopurs_runtime.Value, dictProfunctor1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("dimap", gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, g_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Product2", gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictProfunctor_0, "dimap"), f_2, g_3, gopurs_runtime.ConstructorGet(v_4, 0)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictProfunctor1_1, "dimap"), f_2, g_3, gopurs_runtime.ConstructorGet(v_4, 1)))
}))
})
	})
	return profunctorProduct2
}

var functorProduct2 gopurs_runtime.Value
var once_functorProduct2 sync.Once
func Get_functorProduct2() gopurs_runtime.Value {
	once_functorProduct2.Do(func() {
		functorProduct2 = gopurs_runtime.Func2(func(dictFunctor_0 gopurs_runtime.Value, dictFunctor1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Product2", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_2, gopurs_runtime.ConstructorGet(v_3, 0)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor1_1, "map"), f_2, gopurs_runtime.ConstructorGet(v_3, 1)))
}))
})
	})
	return functorProduct2
}

var eqProduct2 gopurs_runtime.Value
var once_eqProduct2 sync.Once
func Get_eqProduct2() gopurs_runtime.Value {
	once_eqProduct2.Do(func() {
		eqProduct2 = gopurs_runtime.Func2(func(dictEq_0 gopurs_runtime.Value, dictEq1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), gopurs_runtime.ConstructorGet(x_2, 0), gopurs_runtime.ConstructorGet(y_3, 0)).IntVal != 0 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), gopurs_runtime.ConstructorGet(x_2, 1), gopurs_runtime.ConstructorGet(y_3, 1)).IntVal != 0)
}))
})
	})
	return eqProduct2
}

var ordProduct2 gopurs_runtime.Value
var once_ordProduct2 sync.Once
func Get_ordProduct2() gopurs_runtime.Value {
	once_ordProduct2.Do(func() {
		ordProduct2 = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictOrd1_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_2, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_3_1
eqProduct22_4_2 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_4 gopurs_runtime.Value, y_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "eq"), gopurs_runtime.ConstructorGet(x_4, 0), gopurs_runtime.ConstructorGet(y_5, 0)).IntVal != 0 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_1, "eq"), gopurs_runtime.ConstructorGet(x_4, 1), gopurs_runtime.ConstructorGet(y_5, 1)).IntVal != 0)
}))
_ = eqProduct22_4_2
return gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(x_5 gopurs_runtime.Value, y_6 gopurs_runtime.Value) gopurs_runtime.Value {
v_7_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), gopurs_runtime.ConstructorGet(x_5, 0), gopurs_runtime.ConstructorGet(y_6, 0))
_ = v_7_3
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_7_3.StrVal == "LT")).IntVal != 0 {
__t4 = gopurs_runtime.Constructor0("LT")
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(v_7_3.StrVal == "GT")).IntVal != 0 {
__t4 = gopurs_runtime.Constructor0("GT")
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), gopurs_runtime.ConstructorGet(x_5, 1), gopurs_runtime.ConstructorGet(y_6, 1))
}
end_branch_4:
return __t4
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return eqProduct22_4_2
}))
})
})
	})
	return ordProduct2
}

var bifunctorProduct2 gopurs_runtime.Value
var once_bifunctorProduct2 sync.Once
func Get_bifunctorProduct2() gopurs_runtime.Value {
	once_bifunctorProduct2.Do(func() {
		bifunctorProduct2 = gopurs_runtime.Func2(func(dictBifunctor_0 gopurs_runtime.Value, dictBifunctor1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, g_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Product2", gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictBifunctor_0, "bimap"), f_2, g_3, gopurs_runtime.ConstructorGet(v_4, 0)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictBifunctor1_1, "bimap"), f_2, g_3, gopurs_runtime.ConstructorGet(v_4, 1)))
}))
})
	})
	return bifunctorProduct2
}

var biapplyProduct2 gopurs_runtime.Value
var once_biapplyProduct2 sync.Once
func Get_biapplyProduct2() gopurs_runtime.Value {
	once_biapplyProduct2.Do(func() {
		biapplyProduct2 = gopurs_runtime.Func(func(dictBiapply_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBiapply_0, "Bifunctor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictBiapply1_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBiapply1_2, "Bifunctor0"), gopurs_runtime.Value{})
_ = __local_var_3_1
bifunctorProduct22_4_2 := gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func3(func(f_4 gopurs_runtime.Value, g_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Product2", gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_0, "bimap"), f_4, g_5, gopurs_runtime.ConstructorGet(v_6, 0)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_3_1, "bimap"), f_4, g_5, gopurs_runtime.ConstructorGet(v_6, 1)))
}))
_ = bifunctorProduct22_4_2
return gopurs_runtime.RecordDict2("biapply", "Bifunctor0", gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Product2", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBiapply_0, "biapply"), gopurs_runtime.ConstructorGet(v_5, 0), gopurs_runtime.ConstructorGet(v1_6, 0)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBiapply1_2, "biapply"), gopurs_runtime.ConstructorGet(v_5, 1), gopurs_runtime.ConstructorGet(v1_6, 1)))
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorProduct22_4_2
}))
})
})
	})
	return biapplyProduct2
}

var biapplicativeProduct2 gopurs_runtime.Value
var once_biapplicativeProduct2 sync.Once
func Get_biapplicativeProduct2() gopurs_runtime.Value {
	once_biapplicativeProduct2.Do(func() {
		biapplicativeProduct2 = gopurs_runtime.Func(func(dictBiapplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBiapplicative_0, "Biapply0"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Bifunctor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func(func(dictBiapplicative1_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBiapplicative1_3, "Biapply0"), gopurs_runtime.Value{})
_ = __local_var_4_2
__local_var_5_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_2, "Bifunctor0"), gopurs_runtime.Value{})
_ = __local_var_5_3
bifunctorProduct22_6_5 := gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func3(func(f_6 gopurs_runtime.Value, g_7 gopurs_runtime.Value, v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Product2", gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_1, "bimap"), f_6, g_7, gopurs_runtime.ConstructorGet(v_8, 0)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_5_3, "bimap"), f_6, g_7, gopurs_runtime.ConstructorGet(v_8, 1)))
}))
_ = bifunctorProduct22_6_5
biapplyProduct22_6_4 := gopurs_runtime.RecordDict2("biapply", "Bifunctor0", gopurs_runtime.Func2(func(v_7 gopurs_runtime.Value, v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Product2", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "biapply"), gopurs_runtime.ConstructorGet(v_7, 0), gopurs_runtime.ConstructorGet(v1_8, 0)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_2, "biapply"), gopurs_runtime.ConstructorGet(v_7, 1), gopurs_runtime.ConstructorGet(v1_8, 1)))
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorProduct22_6_5
}))
_ = biapplyProduct22_6_4
return gopurs_runtime.RecordDict2("bipure", "Biapply0", gopurs_runtime.Func2(func(a_7 gopurs_runtime.Value, b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Product2", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBiapplicative_0, "bipure"), a_7, b_8), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBiapplicative1_3, "bipure"), a_7, b_8))
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return biapplyProduct22_6_4
}))
})
})
	})
	return biapplicativeProduct2
}


