package Data_Functor_Product

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Bifunctor "gopurs/output/Data.Bifunctor"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	unsafe "unsafe"
)

var cache_Product gopurs_runtime.Value
var once_Product sync.Once
func Get_Product() gopurs_runtime.Value {
	once_Product.Do(func() {
		cache_Product = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Product(x_0_box)
})
	})
	return cache_Product
}

var cache_showProduct gopurs_runtime.Value
var once_showProduct sync.Once
func Get_showProduct() gopurs_runtime.Value {
	once_showProduct.Do(func() {
		cache_showProduct = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, dictShow1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showProduct(dictShow_0_box, dictShow1_1_box)
})
	})
	return cache_showProduct
}

var cache_product gopurs_runtime.Value
var once_product sync.Once
func Get_product() gopurs_runtime.Value {
	once_product.Do(func() {
		cache_product = gopurs_runtime.Func2(func(fa_0_box gopurs_runtime.Value, ga_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_product(fa_0_box, ga_1_box)
})
	})
	return cache_product
}

var cache_product__gopurs_runtime_Value_1903793087 gopurs_runtime.Value
var once_product__gopurs_runtime_Value_1903793087 sync.Once
func Get_product__gopurs_runtime_Value_1903793087() gopurs_runtime.Value {
	once_product__gopurs_runtime_Value_1903793087.Do(func() {
		cache_product__gopurs_runtime_Value_1903793087 = gopurs_runtime.Func2(func(fa_0_box gopurs_runtime.Value, ga_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_product__gopurs_runtime_Value_1903793087(fa_0_box, ga_1_box)
})
	})
	return cache_product__gopurs_runtime_Value_1903793087
}

var cache_product__gopurs_runtime_Value_1470022684 gopurs_runtime.Value
var once_product__gopurs_runtime_Value_1470022684 sync.Once
func Get_product__gopurs_runtime_Value_1470022684() gopurs_runtime.Value {
	once_product__gopurs_runtime_Value_1470022684.Do(func() {
		cache_product__gopurs_runtime_Value_1470022684 = gopurs_runtime.Func2(func(fa_0_box gopurs_runtime.Value, ga_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_product__gopurs_runtime_Value_1470022684(fa_0_box, ga_1_box)
})
	})
	return cache_product__gopurs_runtime_Value_1470022684
}

var cache_product__gopurs_runtime_Value_2238817533 gopurs_runtime.Value
var once_product__gopurs_runtime_Value_2238817533 sync.Once
func Get_product__gopurs_runtime_Value_2238817533() gopurs_runtime.Value {
	once_product__gopurs_runtime_Value_2238817533.Do(func() {
		cache_product__gopurs_runtime_Value_2238817533 = gopurs_runtime.Func2(func(fa_0_box gopurs_runtime.Value, ga_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_product__gopurs_runtime_Value_2238817533(fa_0_box, ga_1_box)
})
	})
	return cache_product__gopurs_runtime_Value_2238817533
}

var cache_product__gopurs_runtime_Value_3006469466 gopurs_runtime.Value
var once_product__gopurs_runtime_Value_3006469466 sync.Once
func Get_product__gopurs_runtime_Value_3006469466() gopurs_runtime.Value {
	once_product__gopurs_runtime_Value_3006469466.Do(func() {
		cache_product__gopurs_runtime_Value_3006469466 = gopurs_runtime.Func2(func(fa_0_box gopurs_runtime.Value, ga_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_product__gopurs_runtime_Value_3006469466(fa_0_box, ga_1_box)
})
	})
	return cache_product__gopurs_runtime_Value_3006469466
}

var cache_product__gopurs_runtime_Value_4078581051 gopurs_runtime.Value
var once_product__gopurs_runtime_Value_4078581051 sync.Once
func Get_product__gopurs_runtime_Value_4078581051() gopurs_runtime.Value {
	once_product__gopurs_runtime_Value_4078581051.Do(func() {
		cache_product__gopurs_runtime_Value_4078581051 = gopurs_runtime.Func2(func(fa_0_box gopurs_runtime.Value, ga_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_product__gopurs_runtime_Value_4078581051(fa_0_box, ga_1_box)
})
	})
	return cache_product__gopurs_runtime_Value_4078581051
}

var cache_product__gopurs_runtime_Value_1134324115 gopurs_runtime.Value
var once_product__gopurs_runtime_Value_1134324115 sync.Once
func Get_product__gopurs_runtime_Value_1134324115() gopurs_runtime.Value {
	once_product__gopurs_runtime_Value_1134324115.Do(func() {
		cache_product__gopurs_runtime_Value_1134324115 = gopurs_runtime.Func2(func(fa_0_box gopurs_runtime.Value, ga_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_product__gopurs_runtime_Value_1134324115(fa_0_box, ga_1_box)
})
	})
	return cache_product__gopurs_runtime_Value_1134324115
}

var cache_product__gopurs_runtime_Value_3644810648 gopurs_runtime.Value
var once_product__gopurs_runtime_Value_3644810648 sync.Once
func Get_product__gopurs_runtime_Value_3644810648() gopurs_runtime.Value {
	once_product__gopurs_runtime_Value_3644810648.Do(func() {
		cache_product__gopurs_runtime_Value_3644810648 = gopurs_runtime.Func2(func(fa_0_box gopurs_runtime.Value, ga_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_product__gopurs_runtime_Value_3644810648(fa_0_box, ga_1_box)
})
	})
	return cache_product__gopurs_runtime_Value_3644810648
}

var cache_product__gopurs_runtime_Value_118638201 gopurs_runtime.Value
var once_product__gopurs_runtime_Value_118638201 sync.Once
func Get_product__gopurs_runtime_Value_118638201() gopurs_runtime.Value {
	once_product__gopurs_runtime_Value_118638201.Do(func() {
		cache_product__gopurs_runtime_Value_118638201 = gopurs_runtime.Func2(func(fa_0_box gopurs_runtime.Value, ga_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_product__gopurs_runtime_Value_118638201(fa_0_box, ga_1_box)
})
	})
	return cache_product__gopurs_runtime_Value_118638201
}

var cache_product__gopurs_runtime_Value_188821718 gopurs_runtime.Value
var once_product__gopurs_runtime_Value_188821718 sync.Once
func Get_product__gopurs_runtime_Value_188821718() gopurs_runtime.Value {
	once_product__gopurs_runtime_Value_188821718.Do(func() {
		cache_product__gopurs_runtime_Value_188821718 = gopurs_runtime.Func2(func(fa_0_box gopurs_runtime.Value, ga_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_product__gopurs_runtime_Value_188821718(fa_0_box, ga_1_box)
})
	})
	return cache_product__gopurs_runtime_Value_188821718
}

var cache_product__gopurs_runtime_Value_1260933303 gopurs_runtime.Value
var once_product__gopurs_runtime_Value_1260933303 sync.Once
func Get_product__gopurs_runtime_Value_1260933303() gopurs_runtime.Value {
	once_product__gopurs_runtime_Value_1260933303.Do(func() {
		cache_product__gopurs_runtime_Value_1260933303 = gopurs_runtime.Func2(func(fa_0_box gopurs_runtime.Value, ga_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_product__gopurs_runtime_Value_1260933303(fa_0_box, ga_1_box)
})
	})
	return cache_product__gopurs_runtime_Value_1260933303
}

var cache_product__gopurs_runtime_Value_827162900 gopurs_runtime.Value
var once_product__gopurs_runtime_Value_827162900 sync.Once
func Get_product__gopurs_runtime_Value_827162900() gopurs_runtime.Value {
	once_product__gopurs_runtime_Value_827162900.Do(func() {
		cache_product__gopurs_runtime_Value_827162900 = gopurs_runtime.Func2(func(fa_0_box gopurs_runtime.Value, ga_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_product__gopurs_runtime_Value_827162900(fa_0_box, ga_1_box)
})
	})
	return cache_product__gopurs_runtime_Value_827162900
}

var cache_newtypeProduct gopurs_runtime.Value
var once_newtypeProduct sync.Once
func Get_newtypeProduct() gopurs_runtime.Value {
	once_newtypeProduct.Do(func() {
		cache_newtypeProduct = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeProduct
}

var cache_functorProduct gopurs_runtime.Value
var once_functorProduct sync.Once
func Get_functorProduct() gopurs_runtime.Value {
	once_functorProduct.Do(func() {
		cache_functorProduct = gopurs_runtime.Func2(func(dictFunctor_0_box gopurs_runtime.Value, dictFunctor1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_functorProduct(dictFunctor_0_box, dictFunctor1_1_box)
})
	})
	return cache_functorProduct
}

var cache_eq1Product gopurs_runtime.Value
var once_eq1Product sync.Once
func Get_eq1Product() gopurs_runtime.Value {
	once_eq1Product.Do(func() {
		cache_eq1Product = gopurs_runtime.Func2(func(dictEq1_0_box gopurs_runtime.Value, dictEq11_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq1Product(dictEq1_0_box, dictEq11_1_box)
})
	})
	return cache_eq1Product
}

var cache_eqProduct gopurs_runtime.Value
var once_eqProduct sync.Once
func Get_eqProduct() gopurs_runtime.Value {
	once_eqProduct.Do(func() {
		cache_eqProduct = gopurs_runtime.Func3(func(dictEq1_0_box gopurs_runtime.Value, dictEq11_1_box gopurs_runtime.Value, dictEq_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqProduct(dictEq1_0_box, dictEq11_1_box, dictEq_2_box)
})
	})
	return cache_eqProduct
}

var cache_ord1Product gopurs_runtime.Value
var once_ord1Product sync.Once
func Get_ord1Product() gopurs_runtime.Value {
	once_ord1Product.Do(func() {
		cache_ord1Product = gopurs_runtime.Func(func(dictOrd1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ord1Product(dictOrd1_0_box)
})
	})
	return cache_ord1Product
}

var cache_ordProduct gopurs_runtime.Value
var once_ordProduct sync.Once
func Get_ordProduct() gopurs_runtime.Value {
	once_ordProduct.Do(func() {
		cache_ordProduct = gopurs_runtime.Func(func(dictOrd1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordProduct(dictOrd1_0_box)
})
	})
	return cache_ordProduct
}

var cache_bihoistProduct gopurs_runtime.Value
var once_bihoistProduct sync.Once
func Get_bihoistProduct() gopurs_runtime.Value {
	once_bihoistProduct.Do(func() {
		cache_bihoistProduct = gopurs_runtime.Func3(func(natF_0_box gopurs_runtime.Value, natG_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bihoistProduct(natF_0_box, natG_1_box, v_2_box)
})
	})
	return cache_bihoistProduct
}

var cache_applyProduct gopurs_runtime.Value
var once_applyProduct sync.Once
func Get_applyProduct() gopurs_runtime.Value {
	once_applyProduct.Do(func() {
		cache_applyProduct = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applyProduct(dictApply_0_box)
})
	})
	return cache_applyProduct
}

var cache_bindProduct gopurs_runtime.Value
var once_bindProduct sync.Once
func Get_bindProduct() gopurs_runtime.Value {
	once_bindProduct.Do(func() {
		cache_bindProduct = gopurs_runtime.Func(func(dictBind_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindProduct(dictBind_0_box)
})
	})
	return cache_bindProduct
}

var cache_applicativeProduct gopurs_runtime.Value
var once_applicativeProduct sync.Once
func Get_applicativeProduct() gopurs_runtime.Value {
	once_applicativeProduct.Do(func() {
		cache_applicativeProduct = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applicativeProduct(dictApplicative_0_box)
})
	})
	return cache_applicativeProduct
}

var cache_monadProduct gopurs_runtime.Value
var once_monadProduct sync.Once
func Get_monadProduct() gopurs_runtime.Value {
	once_monadProduct.Do(func() {
		cache_monadProduct = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadProduct(dictMonad_0_box)
})
	})
	return cache_monadProduct
}

func Call_Product(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_showProduct(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
_ = dictShow1_1
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(product "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str(" "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1), gopurs_runtime.Str(")")))))
}))
}

func Call_product(fa_0_loop gopurs_runtime.Value, ga_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var fa_0 gopurs_runtime.Value = fa_0_loop
_ = fa_0
var ga_1 gopurs_runtime.Value = ga_1_loop
_ = ga_1
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{fa_0, ga_1})}
}

func Call_product__gopurs_runtime_Value_1903793087(fa_0_loop gopurs_runtime.Value, ga_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var fa_0 gopurs_runtime.Value = fa_0_loop
_ = fa_0
var ga_1 gopurs_runtime.Value = ga_1_loop
_ = ga_1
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{fa_0, ga_1})}
}

func Call_product__gopurs_runtime_Value_1470022684(fa_0_loop gopurs_runtime.Value, ga_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var fa_0 gopurs_runtime.Value = fa_0_loop
_ = fa_0
var ga_1 gopurs_runtime.Value = ga_1_loop
_ = ga_1
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{fa_0, ga_1})}
}

func Call_product__gopurs_runtime_Value_2238817533(fa_0_loop gopurs_runtime.Value, ga_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var fa_0 gopurs_runtime.Value = fa_0_loop
_ = fa_0
var ga_1 gopurs_runtime.Value = ga_1_loop
_ = ga_1
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{fa_0, ga_1})}
}

func Call_product__gopurs_runtime_Value_3006469466(fa_0_loop gopurs_runtime.Value, ga_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var fa_0 gopurs_runtime.Value = fa_0_loop
_ = fa_0
var ga_1 gopurs_runtime.Value = ga_1_loop
_ = ga_1
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{fa_0, ga_1})}
}

func Call_product__gopurs_runtime_Value_4078581051(fa_0_loop gopurs_runtime.Value, ga_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var fa_0 gopurs_runtime.Value = fa_0_loop
_ = fa_0
var ga_1 gopurs_runtime.Value = ga_1_loop
_ = ga_1
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{fa_0, ga_1})}
}

func Call_product__gopurs_runtime_Value_1134324115(fa_0_loop gopurs_runtime.Value, ga_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var fa_0 gopurs_runtime.Value = fa_0_loop
_ = fa_0
var ga_1 gopurs_runtime.Value = ga_1_loop
_ = ga_1
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{fa_0, ga_1})}
}

func Call_product__gopurs_runtime_Value_3644810648(fa_0_loop gopurs_runtime.Value, ga_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var fa_0 gopurs_runtime.Value = fa_0_loop
_ = fa_0
var ga_1 gopurs_runtime.Value = ga_1_loop
_ = ga_1
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{fa_0, ga_1})}
}

func Call_product__gopurs_runtime_Value_118638201(fa_0_loop gopurs_runtime.Value, ga_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var fa_0 gopurs_runtime.Value = fa_0_loop
_ = fa_0
var ga_1 gopurs_runtime.Value = ga_1_loop
_ = ga_1
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{fa_0, ga_1})}
}

func Call_product__gopurs_runtime_Value_188821718(fa_0_loop gopurs_runtime.Value, ga_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var fa_0 gopurs_runtime.Value = fa_0_loop
_ = fa_0
var ga_1 gopurs_runtime.Value = ga_1_loop
_ = ga_1
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{fa_0, ga_1})}
}

func Call_product__gopurs_runtime_Value_1260933303(fa_0_loop gopurs_runtime.Value, ga_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var fa_0 gopurs_runtime.Value = fa_0_loop
_ = fa_0
var ga_1 gopurs_runtime.Value = ga_1_loop
_ = ga_1
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{fa_0, ga_1})}
}

func Call_product__gopurs_runtime_Value_827162900(fa_0_loop gopurs_runtime.Value, ga_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var fa_0 gopurs_runtime.Value = fa_0_loop
_ = fa_0
var ga_1 gopurs_runtime.Value = ga_1_loop
_ = ga_1
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{fa_0, ga_1})}
}

func Call_functorProduct(dictFunctor_0_loop gopurs_runtime.Value, dictFunctor1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var dictFunctor1_1 gopurs_runtime.Value = dictFunctor1_1_loop
_ = dictFunctor1_1
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_Bifunctor.Get_bifunctorTuple(), "bimap"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_2), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor1_1, "map"), f_2), v_3)
}))
}

func Call_eq1Product(dictEq1_0_loop gopurs_runtime.Value, dictEq11_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
_ = dictEq1_0
var dictEq11_1 gopurs_runtime.Value = dictEq11_1_loop
_ = dictEq11_1
return gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_2 gopurs_runtime.Value) gopurs_runtime.Value {
eq12_3_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), dictEq_2)
_ = eq12_3_0
eq13_4_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEq11_1, "eq1"), dictEq_2)
_ = eq13_4_1
return gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(eq12_3_0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V0), gopurs_runtime.Apply2(eq13_4_1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V1))
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
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_eq1Product(dictEq1_0, dictEq11_1), "eq1"), dictEq_2))
}

func Call_ord1Product(dictOrd1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd1_0 gopurs_runtime.Value = dictOrd1_0_loop
_ = dictOrd1_0
eq1Product1_1_0 := gopurs_runtime.Apply(Get_eq1Product(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{}))
_ = eq1Product1_1_0
return gopurs_runtime.Func(func(dictOrd11_2 gopurs_runtime.Value) gopurs_runtime.Value {
eq1Product2_3_1 := gopurs_runtime.Apply(eq1Product1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd11_2, "Eq10"), gopurs_runtime.Value{}))
_ = eq1Product2_3_1
return gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eq1Product2_3_1
}), gopurs_runtime.Func(func(dictOrd_4 gopurs_runtime.Value) gopurs_runtime.Value {
compare12_5_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "compare1"), dictOrd_4)
_ = compare12_5_2
compare13_6_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd11_2, "compare1"), dictOrd_4)
_ = compare13_6_3
return gopurs_runtime.Func2(func(v_7 gopurs_runtime.Value, v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
v2_9_4 := gopurs_runtime.Apply2(compare12_5_2, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_8.UnsafePtr).V0)
_ = v2_9_4
var __t5 gopurs_runtime.Value
{
if (v2_9_4.Type == 9 && v2_9_4.IntVal == 902936544) {
__t5 = gopurs_runtime.Apply2(compare13_6_3, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_8.UnsafePtr).V1)
goto end_branch_5
} else {

}
}
{
__t5 = v2_9_4
}
end_branch_5:
return __t5
})
}))
})
}

func Call_ordProduct(dictOrd1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd1_0 gopurs_runtime.Value = dictOrd1_0_loop
_ = dictOrd1_0
ord1Product1_1_0 := Call_ord1Product(dictOrd1_0)
_ = ord1Product1_1_0
eqProduct1_2_1 := gopurs_runtime.Apply(Get_eqProduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{}))
_ = eqProduct1_2_1
return gopurs_runtime.Func(func(dictOrd11_3 gopurs_runtime.Value) gopurs_runtime.Value {
eqProduct2_4_2 := gopurs_runtime.Apply(eqProduct1_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd11_3, "Eq10"), gopurs_runtime.Value{}))
_ = eqProduct2_4_2
return gopurs_runtime.Func(func(dictOrd_5 gopurs_runtime.Value) gopurs_runtime.Value {
eqProduct3_6_3 := gopurs_runtime.Apply(eqProduct2_4_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_5, "Eq0"), gopurs_runtime.Value{}))
_ = eqProduct3_6_3
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return eqProduct3_6_3
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(ord1Product1_1_0, dictOrd11_3), "compare1"), dictOrd_5))
})
})
}

func Call_bihoistProduct(natF_0_loop gopurs_runtime.Value, natG_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var natF_0 gopurs_runtime.Value = natF_0_loop
_ = natF_0
var natG_1 gopurs_runtime.Value = natG_1_loop
_ = natG_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_Bifunctor.Get_bifunctorTuple(), "bimap"), natF_0, natG_1, v_2)
}

func Call_applyProduct(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
functorProduct1_1_0 := gopurs_runtime.Apply(Get_functorProduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = functorProduct1_1_0
return gopurs_runtime.Func(func(dictApply1_2 gopurs_runtime.Value) gopurs_runtime.Value {
functorProduct2_3_1 := gopurs_runtime.Apply(functorProduct1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply1_2, "Functor0"), gopurs_runtime.Value{}))
_ = functorProduct2_3_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorProduct2_3_1
}), gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply1_2, "apply"), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V1)})}
}))
})
}

func Call_bindProduct(dictBind_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
applyProduct1_1_0 := Call_applyProduct(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind_0, "Apply0"), gopurs_runtime.Value{}))
_ = applyProduct1_1_0
return gopurs_runtime.Func(func(dictBind1_2 gopurs_runtime.Value) gopurs_runtime.Value {
applyProduct2_3_1 := gopurs_runtime.Apply(applyProduct1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind1_2, "Apply0"), gopurs_runtime.Value{}))
_ = applyProduct2_3_1
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return applyProduct2_3_1
}), gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_0, "bind"), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(f_5, x_6).UnsafePtr).V0
})), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind1_2, "bind"), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(f_5, x_6).UnsafePtr).V1
}))})}
}))
})
}

func Call_applicativeProduct(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
applyProduct1_1_0 := Call_applyProduct(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = applyProduct1_1_0
return gopurs_runtime.Func(func(dictApplicative1_2 gopurs_runtime.Value) gopurs_runtime.Value {
applyProduct2_3_1 := gopurs_runtime.Apply(applyProduct1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative1_2, "Apply0"), gopurs_runtime.Value{}))
_ = applyProduct2_3_1
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return applyProduct2_3_1
}), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), a_4), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative1_2, "pure"), a_4)})}
}))
})
}

func Call_monadProduct(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
applicativeProduct1_1_0 := Call_applicativeProduct(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = applicativeProduct1_1_0
bindProduct1_2_1 := Call_bindProduct(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
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
}


