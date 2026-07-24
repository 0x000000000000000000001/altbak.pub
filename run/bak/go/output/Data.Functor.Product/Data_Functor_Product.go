package Data_Functor_Product

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var Product gopurs_runtime.Value
var once_Product sync.Once
func Get_Product() gopurs_runtime.Value {
	once_Product.Do(func() {
		Product = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0_loop
}()
})
	})
	return Product
}

var showProduct gopurs_runtime.Value
var once_showProduct sync.Once
func Get_showProduct() gopurs_runtime.Value {
	once_showProduct.Do(func() {
		showProduct = gopurs_runtime.Func2(Call_showProduct)
	})
	return showProduct
}

var product gopurs_runtime.Value
var once_product sync.Once
func Get_product() gopurs_runtime.Value {
	once_product.Do(func() {
		product = gopurs_runtime.Func2(Call_product)
	})
	return product
}

var newtypeProduct gopurs_runtime.Value
var once_newtypeProduct sync.Once
func Get_newtypeProduct() gopurs_runtime.Value {
	once_newtypeProduct.Do(func() {
		newtypeProduct = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return newtypeProduct
}

var functorProduct gopurs_runtime.Value
var once_functorProduct sync.Once
func Get_functorProduct() gopurs_runtime.Value {
	once_functorProduct.Do(func() {
		functorProduct = gopurs_runtime.Func2(Call_functorProduct)
	})
	return functorProduct
}

var eq1Product gopurs_runtime.Value
var once_eq1Product sync.Once
func Get_eq1Product() gopurs_runtime.Value {
	once_eq1Product.Do(func() {
		eq1Product = gopurs_runtime.Func2(Call_eq1Product)
	})
	return eq1Product
}

var eqProduct gopurs_runtime.Value
var once_eqProduct sync.Once
func Get_eqProduct() gopurs_runtime.Value {
	once_eqProduct.Do(func() {
		eqProduct = gopurs_runtime.Func3(Call_eqProduct)
	})
	return eqProduct
}

var ord1Product gopurs_runtime.Value
var once_ord1Product sync.Once
func Get_ord1Product() gopurs_runtime.Value {
	once_ord1Product.Do(func() {
		ord1Product = gopurs_runtime.Func(func(dictOrd1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd1_0 gopurs_runtime.Value = dictOrd1_0_loop
_ = dictOrd1_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0_loop, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictOrd11_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd11_2, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_3_1
eq1Product2_4_2 := gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_4 gopurs_runtime.Value) gopurs_runtime.Value {
eq12_5_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "eq1"), dictEq_4)
_ = eq12_5_3
eq13_6_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_1, "eq1"), dictEq_4)
_ = eq13_6_4
return gopurs_runtime.Func2(func(v_7 gopurs_runtime.Value, v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Apply2(eq12_5_3, (*[1024]gopurs_runtime.Value)(v_7.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_8.UnsafePtr)[0]).IntVal != 0 && gopurs_runtime.Apply2(eq13_6_4, (*[1024]gopurs_runtime.Value)(v_7.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v1_8.UnsafePtr)[1]).IntVal != 0)
})
}))
_ = eq1Product2_4_2
return gopurs_runtime.RecordDict2("compare1", "Eq10", gopurs_runtime.Func(func(dictOrd_5 gopurs_runtime.Value) gopurs_runtime.Value {
compare12_6_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0_loop, "compare1"), dictOrd_5)
_ = compare12_6_5
compare13_7_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd11_2, "compare1"), dictOrd_5)
_ = compare13_7_6
return gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
v2_10_7 := gopurs_runtime.Apply2(compare12_6_5, (*[1024]gopurs_runtime.Value)(v_8.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_9.UnsafePtr)[0])
_ = v2_10_7
var __t8 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v2_10_7.StrVal == "EQ").IntVal != 0 {
__t8 = gopurs_runtime.Apply2(compare13_7_6, (*[1024]gopurs_runtime.Value)(v_8.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v1_9.UnsafePtr)[1])
goto end_branch_8
} else {

}
}
{
__t8 = v2_10_7
}
end_branch_8:
return __t8
})
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return eq1Product2_4_2
}))
})
}()
})
	})
	return ord1Product
}

var ordProduct gopurs_runtime.Value
var once_ordProduct sync.Once
func Get_ordProduct() gopurs_runtime.Value {
	once_ordProduct.Do(func() {
		ordProduct = gopurs_runtime.Func(func(dictOrd1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd1_0 gopurs_runtime.Value = dictOrd1_0_loop
_ = dictOrd1_0
ord1Product1_1_0 := gopurs_runtime.Apply(Get_ord1Product(), dictOrd1_0_loop)
_ = ord1Product1_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0_loop, "Eq10"), gopurs_runtime.Value{})
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
eqProduct3_7_4 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(v_9 gopurs_runtime.Value, v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Apply2(eq12_7_5, (*[1024]gopurs_runtime.Value)(v_9.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_10.UnsafePtr)[0]).IntVal != 0 && gopurs_runtime.Apply2(eq13_8_6, (*[1024]gopurs_runtime.Value)(v_9.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v1_10.UnsafePtr)[1]).IntVal != 0)
}))
_ = eqProduct3_7_4
return gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(ord1Product1_1_0, dictOrd11_3), "compare1"), dictOrd_5), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return eqProduct3_7_4
}))
})
})
}()
})
	})
	return ordProduct
}

var bihoistProduct gopurs_runtime.Value
var once_bihoistProduct sync.Once
func Get_bihoistProduct() gopurs_runtime.Value {
	once_bihoistProduct.Do(func() {
		bihoistProduct = gopurs_runtime.Func3(Call_bihoistProduct)
	})
	return bihoistProduct
}

var applyProduct gopurs_runtime.Value
var once_applyProduct sync.Once
func Get_applyProduct() gopurs_runtime.Value {
	once_applyProduct.Do(func() {
		applyProduct = gopurs_runtime.Func(func(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0_loop, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictApply1_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply1_2, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_1
functorProduct2_4_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), f_4, (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[0]), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_1, "map"), f_4, (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[1]))
}))
_ = functorProduct2_4_2
return gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0_loop, "apply"), (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_6.UnsafePtr)[0]), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply1_2, "apply"), (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v1_6.UnsafePtr)[1]))
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorProduct2_4_2
}))
})
}()
})
	})
	return applyProduct
}

var bindProduct gopurs_runtime.Value
var once_bindProduct sync.Once
func Get_bindProduct() gopurs_runtime.Value {
	once_bindProduct.Do(func() {
		bindProduct = gopurs_runtime.Func(func(dictBind_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind_0_loop, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func(func(dictBind1_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind1_3, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_4_2
__local_var_5_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_2, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_3
functorProduct2_6_5 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "map"), f_6, (*[1024]gopurs_runtime.Value)(v_7.UnsafePtr)[0]), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_3, "map"), f_6, (*[1024]gopurs_runtime.Value)(v_7.UnsafePtr)[1]))
}))
_ = functorProduct2_6_5
applyProduct2_6_4 := gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func2(func(v_7 gopurs_runtime.Value, v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "apply"), (*[1024]gopurs_runtime.Value)(v_7.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_8.UnsafePtr)[0]), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_2, "apply"), (*[1024]gopurs_runtime.Value)(v_7.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v1_8.UnsafePtr)[1]))
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return functorProduct2_6_5
}))
_ = applyProduct2_6_4
return gopurs_runtime.RecordDict2("bind", "Apply0", gopurs_runtime.Func2(func(v_7 gopurs_runtime.Value, f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_0_loop, "bind"), (*[1024]gopurs_runtime.Value)(v_7.UnsafePtr)[0], gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return (*[1024]gopurs_runtime.Value)(gopurs_runtime.Apply(f_8, x_9).UnsafePtr)[0]
})), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind1_3, "bind"), (*[1024]gopurs_runtime.Value)(v_7.UnsafePtr)[1], gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return (*[1024]gopurs_runtime.Value)(gopurs_runtime.Apply(f_8, x_9).UnsafePtr)[1]
})))
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applyProduct2_6_4
}))
})
}()
})
	})
	return bindProduct
}

var applicativeProduct gopurs_runtime.Value
var once_applicativeProduct sync.Once
func Get_applicativeProduct() gopurs_runtime.Value {
	once_applicativeProduct.Do(func() {
		applicativeProduct = gopurs_runtime.Func(func(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0_loop, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func(func(dictApplicative1_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative1_3, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_4_2
__local_var_5_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_2, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_3
functorProduct2_6_5 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "map"), f_6, (*[1024]gopurs_runtime.Value)(v_7.UnsafePtr)[0]), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_3, "map"), f_6, (*[1024]gopurs_runtime.Value)(v_7.UnsafePtr)[1]))
}))
_ = functorProduct2_6_5
applyProduct2_6_4 := gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func2(func(v_7 gopurs_runtime.Value, v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "apply"), (*[1024]gopurs_runtime.Value)(v_7.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_8.UnsafePtr)[0]), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_2, "apply"), (*[1024]gopurs_runtime.Value)(v_7.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v1_8.UnsafePtr)[1]))
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return functorProduct2_6_5
}))
_ = applyProduct2_6_4
return gopurs_runtime.RecordDict2("pure", "Apply0", gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0_loop, "pure"), a_7), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative1_3, "pure"), a_7))
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applyProduct2_6_4
}))
})
}()
})
	})
	return applicativeProduct
}

var monadProduct gopurs_runtime.Value
var once_monadProduct sync.Once
func Get_monadProduct() gopurs_runtime.Value {
	once_monadProduct.Do(func() {
		monadProduct = gopurs_runtime.Func(func(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
applicativeProduct1_1_0 := gopurs_runtime.Apply(Get_applicativeProduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0_loop, "Applicative0"), gopurs_runtime.Value{}))
_ = applicativeProduct1_1_0
bindProduct1_2_1 := gopurs_runtime.Apply(Get_bindProduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0_loop, "Bind1"), gopurs_runtime.Value{}))
_ = bindProduct1_2_1
return gopurs_runtime.Func(func(dictMonad1_3 gopurs_runtime.Value) gopurs_runtime.Value {
applicativeProduct2_4_2 := gopurs_runtime.Apply(applicativeProduct1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad1_3, "Applicative0"), gopurs_runtime.Value{}))
_ = applicativeProduct2_4_2
bindProduct2_5_3 := gopurs_runtime.Apply(bindProduct1_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad1_3, "Bind1"), gopurs_runtime.Value{}))
_ = bindProduct2_5_3
return gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeProduct2_4_2
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return bindProduct2_5_3
}))
})
}()
})
	})
	return monadProduct
}

func Call_showProduct(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
_ = dictShow1_1
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("(product " + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0_loop, "show"), (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0]).StrVal + " " + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1_loop, "show"), (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[1]).StrVal + ")")
}))
}

func Call_product(fa_0_loop gopurs_runtime.Value, ga_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var fa_0 gopurs_runtime.Value = fa_0_loop
_ = fa_0
var ga_1 gopurs_runtime.Value = ga_1_loop
_ = ga_1
return gopurs_runtime.Constructor2("Tuple", fa_0_loop, ga_1_loop)
}

func Call_functorProduct(dictFunctor_0_loop gopurs_runtime.Value, dictFunctor1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var dictFunctor1_1 gopurs_runtime.Value = dictFunctor1_1_loop
_ = dictFunctor1_1
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0_loop, "map"), f_2, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[0]), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor1_1_loop, "map"), f_2, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[1]))
}))
}

func Call_eq1Product(dictEq1_0_loop gopurs_runtime.Value, dictEq11_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
_ = dictEq1_0
var dictEq11_1 gopurs_runtime.Value = dictEq11_1_loop
_ = dictEq11_1
return gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_2 gopurs_runtime.Value) gopurs_runtime.Value {
eq12_3_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEq1_0_loop, "eq1"), dictEq_2)
_ = eq12_3_0
eq13_4_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEq11_1_loop, "eq1"), dictEq_2)
_ = eq13_4_1
return gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Apply2(eq12_3_0, (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_6.UnsafePtr)[0]).IntVal != 0 && gopurs_runtime.Apply2(eq13_4_1, (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v1_6.UnsafePtr)[1]).IntVal != 0)
})
}))
}

func Call_eqProduct(dictEq1_0_loop gopurs_runtime.Value, dictEq11_1_loop gopurs_runtime.Value, dictEq_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
_ = dictEq1_0
var dictEq11_1 gopurs_runtime.Value = dictEq11_1_loop
_ = dictEq11_1
var dictEq_2 gopurs_runtime.Value = dictEq_2_loop
_ = dictEq_2
eq12_3_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEq1_0_loop, "eq1"), dictEq_2_loop)
_ = eq12_3_0
eq13_4_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEq11_1_loop, "eq1"), dictEq_2_loop)
_ = eq13_4_1
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Apply2(eq12_3_0, (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_6.UnsafePtr)[0]).IntVal != 0 && gopurs_runtime.Apply2(eq13_4_1, (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v1_6.UnsafePtr)[1]).IntVal != 0)
}))
}

func Call_bihoistProduct(natF_0_loop gopurs_runtime.Value, natG_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var natF_0 gopurs_runtime.Value = natF_0_loop
_ = natF_0
var natG_1 gopurs_runtime.Value = natG_1_loop
_ = natG_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply(natF_0_loop, (*[1024]gopurs_runtime.Value)(v_2_loop.UnsafePtr)[0]), gopurs_runtime.Apply(natG_1_loop, (*[1024]gopurs_runtime.Value)(v_2_loop.UnsafePtr)[1]))
}


