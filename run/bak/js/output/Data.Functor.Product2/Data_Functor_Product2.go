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
return gopurs_runtime.Str("(Product2 " + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0]).StrVal + " " + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[1]).StrVal + ")")
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
return gopurs_runtime.Constructor2("Product2", gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictProfunctor_0, "dimap"), f_2, g_3, (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[0]), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictProfunctor1_1, "dimap"), f_2, g_3, (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[1]))
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
return gopurs_runtime.Constructor2("Product2", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_2, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[0]), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor1_1, "map"), f_2, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[1]))
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
return gopurs_runtime.Boolean(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*[1024]gopurs_runtime.Value)(x_2.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(y_3.UnsafePtr)[0]).IntVal != 0 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*[1024]gopurs_runtime.Value)(x_2.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(y_3.UnsafePtr)[1]).IntVal != 0)
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
return gopurs_runtime.Boolean(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "eq"), (*[1024]gopurs_runtime.Value)(x_4.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(y_5.UnsafePtr)[0]).IntVal != 0 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_1, "eq"), (*[1024]gopurs_runtime.Value)(x_4.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(y_5.UnsafePtr)[1]).IntVal != 0)
}))
_ = eqProduct22_4_2
return gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(x_5 gopurs_runtime.Value, y_6 gopurs_runtime.Value) gopurs_runtime.Value {
v_7_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*[1024]gopurs_runtime.Value)(x_5.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(y_6.UnsafePtr)[0])
_ = v_7_3
var __t4 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_7_3.StrVal == "LT").IntVal != 0 {
__t4 = gopurs_runtime.Constructor0("LT")
goto end_branch_4
} else {

}
}
{
if gopurs_runtime.Bool(v_7_3.StrVal == "GT").IntVal != 0 {
__t4 = gopurs_runtime.Constructor0("GT")
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*[1024]gopurs_runtime.Value)(x_5.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(y_6.UnsafePtr)[1])
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
return gopurs_runtime.Constructor2("Product2", gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictBifunctor_0, "bimap"), f_2, g_3, (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[0]), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictBifunctor1_1, "bimap"), f_2, g_3, (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[1]))
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
return gopurs_runtime.Constructor2("Product2", gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_0, "bimap"), f_4, g_5, (*[1024]gopurs_runtime.Value)(v_6.UnsafePtr)[0]), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_3_1, "bimap"), f_4, g_5, (*[1024]gopurs_runtime.Value)(v_6.UnsafePtr)[1]))
}))
_ = bifunctorProduct22_4_2
return gopurs_runtime.RecordDict2("biapply", "Bifunctor0", gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Product2", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBiapply_0, "biapply"), (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_6.UnsafePtr)[0]), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBiapply1_2, "biapply"), (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v1_6.UnsafePtr)[1]))
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
return gopurs_runtime.Constructor2("Product2", gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_1, "bimap"), f_6, g_7, (*[1024]gopurs_runtime.Value)(v_8.UnsafePtr)[0]), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_5_3, "bimap"), f_6, g_7, (*[1024]gopurs_runtime.Value)(v_8.UnsafePtr)[1]))
}))
_ = bifunctorProduct22_6_5
biapplyProduct22_6_4 := gopurs_runtime.RecordDict2("biapply", "Bifunctor0", gopurs_runtime.Func2(func(v_7 gopurs_runtime.Value, v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Product2", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "biapply"), (*[1024]gopurs_runtime.Value)(v_7.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_8.UnsafePtr)[0]), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_2, "biapply"), (*[1024]gopurs_runtime.Value)(v_7.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v1_8.UnsafePtr)[1]))
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




