package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Functor_Product_unwrap gopurs_runtime.Value
var once_Data_Functor_Product_unwrap sync.Once
func Get_Data_Functor_Product_unwrap() gopurs_runtime.Value {
	once_Data_Functor_Product_unwrap.Do(func() {
		cache_Data_Functor_Product_unwrap = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Data_Functor_Product_unwrap
}

var cache_Data_Functor_Product_unwrap1 gopurs_runtime.Value
var once_Data_Functor_Product_unwrap1 sync.Once
func Get_Data_Functor_Product_unwrap1() gopurs_runtime.Value {
	once_Data_Functor_Product_unwrap1.Do(func() {
		cache_Data_Functor_Product_unwrap1 = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Data_Functor_Product_unwrap1
}

var cache_Data_Functor_Product_Product gopurs_runtime.Value
var once_Data_Functor_Product_Product sync.Once
func Get_Data_Functor_Product_Product() gopurs_runtime.Value {
	once_Data_Functor_Product_Product.Do(func() {
		cache_Data_Functor_Product_Product = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Product_Product(x_0_box)
})
	})
	return cache_Data_Functor_Product_Product
}

var cache_Data_Functor_Product_showProduct gopurs_runtime.Value
var once_Data_Functor_Product_showProduct sync.Once
func Get_Data_Functor_Product_showProduct() gopurs_runtime.Value {
	once_Data_Functor_Product_showProduct.Do(func() {
		cache_Data_Functor_Product_showProduct = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, dictShow1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Product_showProduct(dictShow_0_box, dictShow1_1_box)
})
	})
	return cache_Data_Functor_Product_showProduct
}

var cache_Data_Functor_Product_product gopurs_runtime.Value
var once_Data_Functor_Product_product sync.Once
func Get_Data_Functor_Product_product() gopurs_runtime.Value {
	once_Data_Functor_Product_product.Do(func() {
		cache_Data_Functor_Product_product = gopurs_runtime.Func2(func(fa_0_box gopurs_runtime.Value, ga_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Data_Functor_Product_product(fa_0_box, ga_1_box))}
})
	})
	return cache_Data_Functor_Product_product
}

var cache_Data_Functor_Product_newtypeProduct gopurs_runtime.Value
var once_Data_Functor_Product_newtypeProduct sync.Once
func Get_Data_Functor_Product_newtypeProduct() gopurs_runtime.Value {
	once_Data_Functor_Product_newtypeProduct.Do(func() {
		cache_Data_Functor_Product_newtypeProduct = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_Data_Functor_Product_newtypeProduct
}

var cache_Data_Functor_Product_functorProduct gopurs_runtime.Value
var once_Data_Functor_Product_functorProduct sync.Once
func Get_Data_Functor_Product_functorProduct() gopurs_runtime.Value {
	once_Data_Functor_Product_functorProduct.Do(func() {
		cache_Data_Functor_Product_functorProduct = gopurs_runtime.Func2(func(dictFunctor_0_box gopurs_runtime.Value, dictFunctor1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Product_functorProduct(dictFunctor_0_box, dictFunctor1_1_box)
})
	})
	return cache_Data_Functor_Product_functorProduct
}

var cache_Data_Functor_Product_eq1Product gopurs_runtime.Value
var once_Data_Functor_Product_eq1Product sync.Once
func Get_Data_Functor_Product_eq1Product() gopurs_runtime.Value {
	once_Data_Functor_Product_eq1Product.Do(func() {
		cache_Data_Functor_Product_eq1Product = gopurs_runtime.Func2(func(dictEq1_0_box gopurs_runtime.Value, dictEq11_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Product_eq1Product(dictEq1_0_box, dictEq11_1_box)
})
	})
	return cache_Data_Functor_Product_eq1Product
}

var cache_Data_Functor_Product_eqProduct gopurs_runtime.Value
var once_Data_Functor_Product_eqProduct sync.Once
func Get_Data_Functor_Product_eqProduct() gopurs_runtime.Value {
	once_Data_Functor_Product_eqProduct.Do(func() {
		cache_Data_Functor_Product_eqProduct = gopurs_runtime.Func3(func(dictEq1_0_box gopurs_runtime.Value, dictEq11_1_box gopurs_runtime.Value, dictEq_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Product_eqProduct(dictEq1_0_box, dictEq11_1_box, dictEq_2_box)
})
	})
	return cache_Data_Functor_Product_eqProduct
}

var cache_Data_Functor_Product_ord1Product gopurs_runtime.Value
var once_Data_Functor_Product_ord1Product sync.Once
func Get_Data_Functor_Product_ord1Product() gopurs_runtime.Value {
	once_Data_Functor_Product_ord1Product.Do(func() {
		cache_Data_Functor_Product_ord1Product = gopurs_runtime.Func(func(dictOrd1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Product_ord1Product(dictOrd1_0_box)
})
	})
	return cache_Data_Functor_Product_ord1Product
}

var cache_Data_Functor_Product_ordProduct gopurs_runtime.Value
var once_Data_Functor_Product_ordProduct sync.Once
func Get_Data_Functor_Product_ordProduct() gopurs_runtime.Value {
	once_Data_Functor_Product_ordProduct.Do(func() {
		cache_Data_Functor_Product_ordProduct = gopurs_runtime.Func(func(dictOrd1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Product_ordProduct(dictOrd1_0_box)
})
	})
	return cache_Data_Functor_Product_ordProduct
}

var cache_Data_Functor_Product_bihoistProduct gopurs_runtime.Value
var once_Data_Functor_Product_bihoistProduct sync.Once
func Get_Data_Functor_Product_bihoistProduct() gopurs_runtime.Value {
	once_Data_Functor_Product_bihoistProduct.Do(func() {
		cache_Data_Functor_Product_bihoistProduct = gopurs_runtime.Func3(func(natF_0_box gopurs_runtime.Value, natG_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Product_bihoistProduct(natF_0_box, natG_1_box, v_2_box)
})
	})
	return cache_Data_Functor_Product_bihoistProduct
}

var cache_Data_Functor_Product_applyProduct gopurs_runtime.Value
var once_Data_Functor_Product_applyProduct sync.Once
func Get_Data_Functor_Product_applyProduct() gopurs_runtime.Value {
	once_Data_Functor_Product_applyProduct.Do(func() {
		cache_Data_Functor_Product_applyProduct = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Product_applyProduct(dictApply_0_box)
})
	})
	return cache_Data_Functor_Product_applyProduct
}

var cache_Data_Functor_Product_bindProduct gopurs_runtime.Value
var once_Data_Functor_Product_bindProduct sync.Once
func Get_Data_Functor_Product_bindProduct() gopurs_runtime.Value {
	once_Data_Functor_Product_bindProduct.Do(func() {
		cache_Data_Functor_Product_bindProduct = gopurs_runtime.Func(func(dictBind_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Product_bindProduct(dictBind_0_box)
})
	})
	return cache_Data_Functor_Product_bindProduct
}

var cache_Data_Functor_Product_applicativeProduct gopurs_runtime.Value
var once_Data_Functor_Product_applicativeProduct sync.Once
func Get_Data_Functor_Product_applicativeProduct() gopurs_runtime.Value {
	once_Data_Functor_Product_applicativeProduct.Do(func() {
		cache_Data_Functor_Product_applicativeProduct = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Product_applicativeProduct(dictApplicative_0_box)
})
	})
	return cache_Data_Functor_Product_applicativeProduct
}

var cache_Data_Functor_Product_monadProduct gopurs_runtime.Value
var once_Data_Functor_Product_monadProduct sync.Once
func Get_Data_Functor_Product_monadProduct() gopurs_runtime.Value {
	once_Data_Functor_Product_monadProduct.Do(func() {
		cache_Data_Functor_Product_monadProduct = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Product_monadProduct(dictMonad_0_box)
})
	})
	return cache_Data_Functor_Product_monadProduct
}

var cache_Data_Functor_Product_product__2764631669 gopurs_runtime.Value
var once_Data_Functor_Product_product__2764631669 sync.Once
func Get_Data_Functor_Product_product__2764631669() gopurs_runtime.Value {
	once_Data_Functor_Product_product__2764631669.Do(func() {
		cache_Data_Functor_Product_product__2764631669 = gopurs_runtime.Func2(func(fa_0_box gopurs_runtime.Value, ga_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Data_Functor_Product_product__2764631669(fa_0_box, ga_1_box))}
})
	})
	return cache_Data_Functor_Product_product__2764631669
}

var cache_Data_Functor_Product_product__346705816 gopurs_runtime.Value
var once_Data_Functor_Product_product__346705816 sync.Once
func Get_Data_Functor_Product_product__346705816() gopurs_runtime.Value {
	once_Data_Functor_Product_product__346705816.Do(func() {
		cache_Data_Functor_Product_product__346705816 = gopurs_runtime.Func2(func(fa_0_box gopurs_runtime.Value, ga_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Data_Functor_Product_product__346705816(fa_0_box, ga_1_box))}
})
	})
	return cache_Data_Functor_Product_product__346705816
}

var cache_Data_Functor_Product_product__2025901598 gopurs_runtime.Value
var once_Data_Functor_Product_product__2025901598 sync.Once
func Get_Data_Functor_Product_product__2025901598() gopurs_runtime.Value {
	once_Data_Functor_Product_product__2025901598.Do(func() {
		cache_Data_Functor_Product_product__2025901598 = gopurs_runtime.Func2(func(fa_0_box gopurs_runtime.Value, ga_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Data_Functor_Product_product__2025901598(fa_0_box, ga_1_box))}
})
	})
	return cache_Data_Functor_Product_product__2025901598
}

var cache_Data_Functor_Product_product__3868608679 gopurs_runtime.Value
var once_Data_Functor_Product_product__3868608679 sync.Once
func Get_Data_Functor_Product_product__3868608679() gopurs_runtime.Value {
	once_Data_Functor_Product_product__3868608679.Do(func() {
		cache_Data_Functor_Product_product__3868608679 = gopurs_runtime.Func2(func(fa_0_box gopurs_runtime.Value, ga_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Data_Functor_Product_product__3868608679(fa_0_box, ga_1_box))}
})
	})
	return cache_Data_Functor_Product_product__3868608679
}

var cache_Data_Functor_Product_product__3582060361 gopurs_runtime.Value
var once_Data_Functor_Product_product__3582060361 sync.Once
func Get_Data_Functor_Product_product__3582060361() gopurs_runtime.Value {
	once_Data_Functor_Product_product__3582060361.Do(func() {
		cache_Data_Functor_Product_product__3582060361 = gopurs_runtime.Func2(func(fa_0_box gopurs_runtime.Value, ga_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Data_Functor_Product_product__3582060361(fa_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](ga_1_box)))}
})
	})
	return cache_Data_Functor_Product_product__3582060361
}

var cache_Data_Functor_Product_product__3089971879 gopurs_runtime.Value
var once_Data_Functor_Product_product__3089971879 sync.Once
func Get_Data_Functor_Product_product__3089971879() gopurs_runtime.Value {
	once_Data_Functor_Product_product__3089971879.Do(func() {
		cache_Data_Functor_Product_product__3089971879 = gopurs_runtime.Func2(func(fa_0_box gopurs_runtime.Value, ga_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Data_Functor_Product_product__3089971879(fa_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](ga_1_box)))}
})
	})
	return cache_Data_Functor_Product_product__3089971879
}

var cache_Data_Functor_Product_product__573643081 gopurs_runtime.Value
var once_Data_Functor_Product_product__573643081 sync.Once
func Get_Data_Functor_Product_product__573643081() gopurs_runtime.Value {
	once_Data_Functor_Product_product__573643081.Do(func() {
		cache_Data_Functor_Product_product__573643081 = gopurs_runtime.Func2(func(fa_0_box gopurs_runtime.Value, ga_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Data_Functor_Product_product__573643081(fa_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](ga_1_box)))}
})
	})
	return cache_Data_Functor_Product_product__573643081
}

var cache_Data_Functor_Product_product__808561831 gopurs_runtime.Value
var once_Data_Functor_Product_product__808561831 sync.Once
func Get_Data_Functor_Product_product__808561831() gopurs_runtime.Value {
	once_Data_Functor_Product_product__808561831.Do(func() {
		cache_Data_Functor_Product_product__808561831 = gopurs_runtime.Func2(func(fa_0_box gopurs_runtime.Value, ga_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Data_Functor_Product_product__808561831(fa_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](ga_1_box)))}
})
	})
	return cache_Data_Functor_Product_product__808561831
}

var cache_Data_Functor_Product_product__1448987465 gopurs_runtime.Value
var once_Data_Functor_Product_product__1448987465 sync.Once
func Get_Data_Functor_Product_product__1448987465() gopurs_runtime.Value {
	once_Data_Functor_Product_product__1448987465.Do(func() {
		cache_Data_Functor_Product_product__1448987465 = gopurs_runtime.Func2(func(fa_0_box gopurs_runtime.Value, ga_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Data_Functor_Product_product__1448987465(fa_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](ga_1_box)))}
})
	})
	return cache_Data_Functor_Product_product__1448987465
}

var cache_Data_Functor_Product_product__155426471 gopurs_runtime.Value
var once_Data_Functor_Product_product__155426471 sync.Once
func Get_Data_Functor_Product_product__155426471() gopurs_runtime.Value {
	once_Data_Functor_Product_product__155426471.Do(func() {
		cache_Data_Functor_Product_product__155426471 = gopurs_runtime.Func2(func(fa_0_box gopurs_runtime.Value, ga_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Data_Functor_Product_product__155426471(fa_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](ga_1_box)))}
})
	})
	return cache_Data_Functor_Product_product__155426471
}

var cache_Data_Functor_Product_product__2024275273 gopurs_runtime.Value
var once_Data_Functor_Product_product__2024275273 sync.Once
func Get_Data_Functor_Product_product__2024275273() gopurs_runtime.Value {
	once_Data_Functor_Product_product__2024275273.Do(func() {
		cache_Data_Functor_Product_product__2024275273 = gopurs_runtime.Func2(func(fa_0_box gopurs_runtime.Value, ga_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Data_Functor_Product_product__2024275273(fa_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](ga_1_box)))}
})
	})
	return cache_Data_Functor_Product_product__2024275273
}

var cache_Data_Functor_Product_product__1443041447 gopurs_runtime.Value
var once_Data_Functor_Product_product__1443041447 sync.Once
func Get_Data_Functor_Product_product__1443041447() gopurs_runtime.Value {
	once_Data_Functor_Product_product__1443041447.Do(func() {
		cache_Data_Functor_Product_product__1443041447 = gopurs_runtime.Func2(func(fa_0_box gopurs_runtime.Value, ga_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Data_Functor_Product_product__1443041447(fa_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](ga_1_box)))}
})
	})
	return cache_Data_Functor_Product_product__1443041447
}

var cache_Data_Functor_Product_product__3887050569 gopurs_runtime.Value
var once_Data_Functor_Product_product__3887050569 sync.Once
func Get_Data_Functor_Product_product__3887050569() gopurs_runtime.Value {
	once_Data_Functor_Product_product__3887050569.Do(func() {
		cache_Data_Functor_Product_product__3887050569 = gopurs_runtime.Func2(func(fa_0_box gopurs_runtime.Value, ga_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Data_Functor_Product_product__3887050569(fa_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](ga_1_box)))}
})
	})
	return cache_Data_Functor_Product_product__3887050569
}

func Call_Data_Functor_Product_Product(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Functor_Product_showProduct(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
_ = dictShow1_1
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((((("(product ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V0).StrVal())) + (" ")) + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V1).StrVal())) + (")"))
}))
}

func Call_Data_Functor_Product_product(fa_0_loop gopurs_runtime.Value, ga_1_loop gopurs_runtime.Value) *Constructor_Data_Tuple_Tuple {
var fa_0 gopurs_runtime.Value = fa_0_loop
_ = fa_0
var ga_1 gopurs_runtime.Value = ga_1_loop
_ = ga_1
return &Constructor_Data_Tuple_Tuple{1, fa_0, ga_1}
}

func Call_Data_Functor_Product_functorProduct(dictFunctor_0_loop gopurs_runtime.Value, dictFunctor1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var dictFunctor1_1 gopurs_runtime.Value = dictFunctor1_1_loop
_ = dictFunctor1_1
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_2, (*Constructor_Data_Tuple_Tuple)(v_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor1_1, "map"), f_2, (*Constructor_Data_Tuple_Tuple)(v_3.UnsafePtr).V1)})}
})
}))
}

func Call_Data_Functor_Product_eq1Product(dictEq1_0_loop gopurs_runtime.Value, dictEq11_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
_ = dictEq1_0
var dictEq11_1 gopurs_runtime.Value = dictEq11_1_loop
_ = dictEq11_1
return gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_2))}, (*Constructor_Data_Tuple_Tuple)(v_3.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_4.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq11_1, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_2))}, (*Constructor_Data_Tuple_Tuple)(v_3.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v1_4.UnsafePtr).V1).IntVal) != (0)))
})
})
}))
}

func Call_Data_Functor_Product_eqProduct(dictEq1_0_loop gopurs_runtime.Value, dictEq11_1_loop gopurs_runtime.Value, dictEq_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
_ = dictEq1_0
var dictEq11_1 gopurs_runtime.Value = dictEq11_1_loop
_ = dictEq11_1
var dictEq_2 gopurs_runtime.Value = dictEq_2_loop
_ = dictEq_2
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_2))}, (*Constructor_Data_Tuple_Tuple)(v_3.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_4.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq11_1, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_2))}, (*Constructor_Data_Tuple_Tuple)(v_3.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v1_4.UnsafePtr).V1).IntVal) != (0)))
})
}))
}

func Call_Data_Functor_Product_ord1Product(dictOrd1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd1_0 gopurs_runtime.Value = dictOrd1_0_loop
_ = dictOrd1_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictOrd11_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd11_2, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_3_2
// TAST (Let): eq1Product2_3_1 -> gopurs_runtime.Value
eq1Product2_3_1 := gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_4))}, (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_6.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_3_2, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_4))}, (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v1_6.UnsafePtr).V1).IntVal) != (0)))
})
})
}))
_ = eq1Product2_3_1
return gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eq1Product2_3_1
}), gopurs_runtime.Func(func(dictOrd_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v2_7_3 -> gopurs_runtime.Value
v2_7_3 := gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictOrd1_0, "compare1"), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_4))}, (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_6.UnsafePtr).V0)
_ = v2_7_3
var __t4 uint32
{
if (uint32(v2_7_3.IntVal) == 902936544) {
__t4 = uint32(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictOrd11_2, "compare1"), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_4))}, (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v1_6.UnsafePtr).V1).IntVal)
goto end_branch_4
} else {

}
}
{
__t4 = uint32(v2_7_3.IntVal)
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t4), UnsafePtr: nil}
})
})
}))
})
}

func Call_Data_Functor_Product_ordProduct(dictOrd1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd1_0 gopurs_runtime.Value = dictOrd1_0_loop
_ = dictOrd1_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictOrd11_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd11_2, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_3_1
return gopurs_runtime.Func(func(dictOrd_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_4, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_5_3
// TAST (Let): eqProduct3_5_2 -> gopurs_runtime.Value
eqProduct3_5_2 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](__local_var_5_3))}, (*Constructor_Data_Tuple_Tuple)(v_6.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_7.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_3_1, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](__local_var_5_3))}, (*Constructor_Data_Tuple_Tuple)(v_6.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v1_7.UnsafePtr).V1).IntVal) != (0)))
})
}))
_ = eqProduct3_5_2
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return eqProduct3_5_2
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v2_8_4 -> gopurs_runtime.Value
v2_8_4 := gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictOrd1_0, "compare1"), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_4))}, (*Constructor_Data_Tuple_Tuple)(v_6.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_7.UnsafePtr).V0)
_ = v2_8_4
var __t5 uint32
{
if (uint32(v2_8_4.IntVal) == 902936544) {
__t5 = uint32(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictOrd11_2, "compare1"), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_4))}, (*Constructor_Data_Tuple_Tuple)(v_6.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v1_7.UnsafePtr).V1).IntVal)
goto end_branch_5
} else {

}
}
{
__t5 = uint32(v2_8_4.IntVal)
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t5), UnsafePtr: nil}
})
}))
})
})
}

func Call_Data_Functor_Product_bihoistProduct(natF_0_loop gopurs_runtime.Value, natG_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var natF_0 gopurs_runtime.Value = natF_0_loop
_ = natF_0
var natG_1 gopurs_runtime.Value = natG_1_loop
_ = natG_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(natF_0, (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V0), gopurs_runtime.Apply(natG_1, (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V1)})}
}

func Call_Data_Functor_Product_applyProduct(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictApply1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply1_2, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_2
// TAST (Let): functorProduct2_3_1 -> gopurs_runtime.Value
functorProduct2_3_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), f_4, (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "map"), f_4, (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V1)})}
})
}))
_ = functorProduct2_3_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorProduct2_3_1
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_5.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply1_2, "apply"), (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v1_5.UnsafePtr).V1)})}
})
}))
})
}

func Call_Data_Functor_Product_bindProduct(dictBind_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_1_0
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func(func(dictBind1_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_3 -> gopurs_runtime.Value
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind1_3, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_4_3
// TAST (Let): __local_var_5_5 -> gopurs_runtime.Value
__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_5
// TAST (Let): functorProduct2_5_4 -> gopurs_runtime.Value
functorProduct2_5_4 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "map"), f_6, (*Constructor_Data_Tuple_Tuple)(v_7.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_5, "map"), f_6, (*Constructor_Data_Tuple_Tuple)(v_7.UnsafePtr).V1)})}
})
}))
_ = functorProduct2_5_4
// TAST (Let): applyProduct2_4_2 -> gopurs_runtime.Value
applyProduct2_4_2 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorProduct2_5_4
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "apply"), (*Constructor_Data_Tuple_Tuple)(v_6.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_7.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_3, "apply"), (*Constructor_Data_Tuple_Tuple)(v_6.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v1_7.UnsafePtr).V1)})}
})
}))
_ = applyProduct2_4_2
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return applyProduct2_4_2
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_0, "bind"), (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V0, gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply(f_6, x_7).UnsafePtr).V0
})), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind1_3, "bind"), (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V1, gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply(f_6, x_7).UnsafePtr).V1
}))})}
})
}))
})
}

func Call_Data_Functor_Product_applicativeProduct(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_1_0
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func(func(dictApplicative1_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_3 -> gopurs_runtime.Value
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative1_3, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_4_3
// TAST (Let): __local_var_5_5 -> gopurs_runtime.Value
__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_5
// TAST (Let): functorProduct2_5_4 -> gopurs_runtime.Value
functorProduct2_5_4 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "map"), f_6, (*Constructor_Data_Tuple_Tuple)(v_7.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_5, "map"), f_6, (*Constructor_Data_Tuple_Tuple)(v_7.UnsafePtr).V1)})}
})
}))
_ = functorProduct2_5_4
// TAST (Let): applyProduct2_4_2 -> gopurs_runtime.Value
applyProduct2_4_2 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorProduct2_5_4
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "apply"), (*Constructor_Data_Tuple_Tuple)(v_6.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_7.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_3, "apply"), (*Constructor_Data_Tuple_Tuple)(v_6.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v1_7.UnsafePtr).V1)})}
})
}))
_ = applyProduct2_4_2
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return applyProduct2_4_2
}), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), a_5), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative1_3, "pure"), a_5)})}
}))
})
}

func Call_Data_Functor_Product_monadProduct(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_1_0
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_2_1
// TAST (Let): __local_var_3_3 -> gopurs_runtime.Value
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_3
// TAST (Let): applicativeProduct1_3_2 -> gopurs_runtime.Value
applicativeProduct1_3_2 := gopurs_runtime.Func(func(dictApplicative1_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_5 -> gopurs_runtime.Value
__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative1_4, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_5_5
// TAST (Let): __local_var_6_7 -> gopurs_runtime.Value
__local_var_6_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_7
// TAST (Let): functorProduct2_6_6 -> gopurs_runtime.Value
functorProduct2_6_6 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_3, "map"), f_7, (*Constructor_Data_Tuple_Tuple)(v_8.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_7, "map"), f_7, (*Constructor_Data_Tuple_Tuple)(v_8.UnsafePtr).V1)})}
})
}))
_ = functorProduct2_6_6
// TAST (Let): applyProduct2_5_4 -> gopurs_runtime.Value
applyProduct2_5_4 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return functorProduct2_6_6
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "apply"), (*Constructor_Data_Tuple_Tuple)(v_7.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_8.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_5, "apply"), (*Constructor_Data_Tuple_Tuple)(v_7.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v1_8.UnsafePtr).V1)})}
})
}))
_ = applyProduct2_5_4
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return applyProduct2_5_4
}), gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), a_6), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative1_4, "pure"), a_6)})}
}))
})
_ = applicativeProduct1_3_2
// TAST (Let): __local_var_4_8 -> gopurs_runtime.Value
__local_var_4_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_4_8
// TAST (Let): __local_var_5_9 -> gopurs_runtime.Value
__local_var_5_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_8, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_5_9
// TAST (Let): __local_var_6_11 -> gopurs_runtime.Value
__local_var_6_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_9, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_11
// TAST (Let): bindProduct1_6_10 -> gopurs_runtime.Value
bindProduct1_6_10 := gopurs_runtime.Func(func(dictBind1_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_13 -> gopurs_runtime.Value
__local_var_8_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind1_7, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_8_13
// TAST (Let): __local_var_9_15 -> gopurs_runtime.Value
__local_var_9_15 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_13, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_15
// TAST (Let): functorProduct2_9_14 -> gopurs_runtime.Value
functorProduct2_9_14 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_11, "map"), f_10, (*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_15, "map"), f_10, (*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V1)})}
})
}))
_ = functorProduct2_9_14
// TAST (Let): applyProduct2_8_12 -> gopurs_runtime.Value
applyProduct2_8_12 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return functorProduct2_9_14
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_9, "apply"), (*Constructor_Data_Tuple_Tuple)(v_10.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_11.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_8_13, "apply"), (*Constructor_Data_Tuple_Tuple)(v_10.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v1_11.UnsafePtr).V1)})}
})
}))
_ = applyProduct2_8_12
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return applyProduct2_8_12
}), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_8, "bind"), (*Constructor_Data_Tuple_Tuple)(v_9.UnsafePtr).V0, gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply(f_10, x_11).UnsafePtr).V0
})), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind1_7, "bind"), (*Constructor_Data_Tuple_Tuple)(v_9.UnsafePtr).V1, gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply(f_10, x_11).UnsafePtr).V1
}))})}
})
}))
})
_ = bindProduct1_6_10
return gopurs_runtime.Func(func(dictMonad1_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applicativeProduct2_8_16 -> gopurs_runtime.Value
applicativeProduct2_8_16 := gopurs_runtime.Apply(applicativeProduct1_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad1_7, "Applicative0"), gopurs_runtime.Value{}))
_ = applicativeProduct2_8_16
// TAST (Let): bindProduct2_9_17 -> gopurs_runtime.Value
bindProduct2_9_17 := gopurs_runtime.Apply(bindProduct1_6_10, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad1_7, "Bind1"), gopurs_runtime.Value{}))
_ = bindProduct2_9_17
return gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeProduct2_8_16
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return bindProduct2_9_17
}))
})
}

func Call_Data_Functor_Product_product__2764631669(fa_0_loop gopurs_runtime.Value, ga_1_loop gopurs_runtime.Value) *Constructor_Data_Tuple_Tuple {
var fa_0 gopurs_runtime.Value = fa_0_loop
_ = fa_0
var ga_1 gopurs_runtime.Value = ga_1_loop
_ = ga_1
return &Constructor_Data_Tuple_Tuple{1, fa_0, ga_1}
}

func Call_Data_Functor_Product_product__346705816(fa_0_loop gopurs_runtime.Value, ga_1_loop gopurs_runtime.Value) *Constructor_Data_Tuple_Tuple {
var fa_0 gopurs_runtime.Value = fa_0_loop
_ = fa_0
var ga_1 gopurs_runtime.Value = ga_1_loop
_ = ga_1
return &Constructor_Data_Tuple_Tuple{1, fa_0, ga_1}
}

func Call_Data_Functor_Product_product__2025901598(fa_0_loop gopurs_runtime.Value, ga_1_loop gopurs_runtime.Value) *Constructor_Data_Tuple_Tuple {
var fa_0 gopurs_runtime.Value = fa_0_loop
_ = fa_0
var ga_1 gopurs_runtime.Value = ga_1_loop
_ = ga_1
return &Constructor_Data_Tuple_Tuple{1, fa_0, ga_1}
}

func Call_Data_Functor_Product_product__3868608679(fa_0_loop gopurs_runtime.Value, ga_1_loop gopurs_runtime.Value) *Constructor_Data_Tuple_Tuple {
var fa_0 gopurs_runtime.Value = fa_0_loop
_ = fa_0
var ga_1 gopurs_runtime.Value = ga_1_loop
_ = ga_1
return &Constructor_Data_Tuple_Tuple{1, fa_0, ga_1}
}

func Call_Data_Functor_Product_product__3582060361(fa_0_loop gopurs_runtime.Value, ga_1_loop *Constructor_Data_Tuple_Tuple) *Constructor_Data_Tuple_Tuple {
var fa_0 gopurs_runtime.Value = fa_0_loop
_ = fa_0
var ga_1 *Constructor_Data_Tuple_Tuple = ga_1_loop
_ = ga_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, fa_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(ga_1)}})})
}

func Call_Data_Functor_Product_product__3089971879(fa_0_loop gopurs_runtime.Value, ga_1_loop *Constructor_Data_Tuple_Tuple) *Constructor_Data_Tuple_Tuple {
var fa_0 gopurs_runtime.Value = fa_0_loop
_ = fa_0
var ga_1 *Constructor_Data_Tuple_Tuple = ga_1_loop
_ = ga_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, fa_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(ga_1)}})})
}

func Call_Data_Functor_Product_product__573643081(fa_0_loop gopurs_runtime.Value, ga_1_loop *Constructor_Data_Tuple_Tuple) *Constructor_Data_Tuple_Tuple {
var fa_0 gopurs_runtime.Value = fa_0_loop
_ = fa_0
var ga_1 *Constructor_Data_Tuple_Tuple = ga_1_loop
_ = ga_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, fa_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(ga_1)}})})
}

func Call_Data_Functor_Product_product__808561831(fa_0_loop gopurs_runtime.Value, ga_1_loop *Constructor_Data_Tuple_Tuple) *Constructor_Data_Tuple_Tuple {
var fa_0 gopurs_runtime.Value = fa_0_loop
_ = fa_0
var ga_1 *Constructor_Data_Tuple_Tuple = ga_1_loop
_ = ga_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, fa_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(ga_1)}})})
}

func Call_Data_Functor_Product_product__1448987465(fa_0_loop gopurs_runtime.Value, ga_1_loop *Constructor_Data_Tuple_Tuple) *Constructor_Data_Tuple_Tuple {
var fa_0 gopurs_runtime.Value = fa_0_loop
_ = fa_0
var ga_1 *Constructor_Data_Tuple_Tuple = ga_1_loop
_ = ga_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, fa_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(ga_1)}})})
}

func Call_Data_Functor_Product_product__155426471(fa_0_loop gopurs_runtime.Value, ga_1_loop *Constructor_Data_Tuple_Tuple) *Constructor_Data_Tuple_Tuple {
var fa_0 gopurs_runtime.Value = fa_0_loop
_ = fa_0
var ga_1 *Constructor_Data_Tuple_Tuple = ga_1_loop
_ = ga_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, fa_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(ga_1)}})})
}

func Call_Data_Functor_Product_product__2024275273(fa_0_loop gopurs_runtime.Value, ga_1_loop *Constructor_Data_Tuple_Tuple) *Constructor_Data_Tuple_Tuple {
var fa_0 gopurs_runtime.Value = fa_0_loop
_ = fa_0
var ga_1 *Constructor_Data_Tuple_Tuple = ga_1_loop
_ = ga_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, fa_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(ga_1)}})})
}

func Call_Data_Functor_Product_product__1443041447(fa_0_loop gopurs_runtime.Value, ga_1_loop *Constructor_Data_Tuple_Tuple) *Constructor_Data_Tuple_Tuple {
var fa_0 gopurs_runtime.Value = fa_0_loop
_ = fa_0
var ga_1 *Constructor_Data_Tuple_Tuple = ga_1_loop
_ = ga_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, fa_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(ga_1)}})})
}

func Call_Data_Functor_Product_product__3887050569(fa_0_loop gopurs_runtime.Value, ga_1_loop *Constructor_Data_Tuple_Tuple) *Constructor_Data_Tuple_Tuple {
var fa_0 gopurs_runtime.Value = fa_0_loop
_ = fa_0
var ga_1 *Constructor_Data_Tuple_Tuple = ga_1_loop
_ = ga_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, fa_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(ga_1)}})})
}


