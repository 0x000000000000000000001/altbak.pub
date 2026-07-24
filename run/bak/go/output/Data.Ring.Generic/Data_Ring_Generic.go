package Data_Ring_Generic

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Generic_Rep "gopurs/output/Data.Generic.Rep"
	unsafe "unsafe"
)

var genericSub_prime gopurs_runtime.Value
var once_genericSub_prime sync.Once
func Get_genericSub_prime() gopurs_runtime.Value {
	once_genericSub_prime.Do(func() {
		genericSub_prime = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericSub'")
}()
})
	})
	return genericSub_prime
}

var genericSub gopurs_runtime.Value
var once_genericSub sync.Once
func Get_genericSub() gopurs_runtime.Value {
	once_genericSub.Do(func() {
		genericSub = gopurs_runtime.Func4(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericRing_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericSub(dictGeneric_0_box, dictGenericRing_1_box, x_2_box, y_3_box)
})
	})
	return genericSub
}

var genericRingProduct gopurs_runtime.Value
var once_genericRingProduct sync.Once
func Get_genericRingProduct() gopurs_runtime.Value {
	once_genericRingProduct.Do(func() {
		genericRingProduct = gopurs_runtime.Func2(func(dictGenericRing_0_box gopurs_runtime.Value, dictGenericRing1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericRingProduct(dictGenericRing_0_box, dictGenericRing1_1_box)
})
	})
	return genericRingProduct
}

var genericRingNoArguments gopurs_runtime.Value
var once_genericRingNoArguments sync.Once
func Get_genericRingNoArguments() gopurs_runtime.Value {
	once_genericRingNoArguments.Do(func() {
		genericRingNoArguments = gopurs_runtime.RecordDict1("genericSub'", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2433704057, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Data_Data_Generic_Rep_NoArguments{})}
}))
	})
	return genericRingNoArguments
}

var genericRingConstructor gopurs_runtime.Value
var once_genericRingConstructor sync.Once
func Get_genericRingConstructor() gopurs_runtime.Value {
	once_genericRingConstructor.Do(func() {
		genericRingConstructor = gopurs_runtime.Func(func(dictGenericRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictGenericRing_0 gopurs_runtime.Value = dictGenericRing_0_loop
_ = dictGenericRing_0
return gopurs_runtime.RecordDict1("genericSub'", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericRing_0, "genericSub'"), v_1, v1_2)
}))
}()
})
	})
	return genericRingConstructor
}

var genericRingArgument gopurs_runtime.Value
var once_genericRingArgument sync.Once
func Get_genericRingArgument() gopurs_runtime.Value {
	once_genericRingArgument.Do(func() {
		genericRingArgument = gopurs_runtime.Func(func(dictRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictRing_0 gopurs_runtime.Value = dictRing_0_loop
_ = dictRing_0
return gopurs_runtime.RecordDict1("genericSub'", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictRing_0, "sub"), v_1, v1_2)
}))
}()
})
	})
	return genericRingArgument
}

func Call_genericSub(dictGeneric_0_loop gopurs_runtime.Value, dictGenericRing_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericRing_1 gopurs_runtime.Value = dictGenericRing_1_loop
_ = dictGenericRing_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "to"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericRing_1, "genericSub'"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "from"), x_2), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "from"), y_3)))
}

func Call_genericRingProduct(dictGenericRing_0_loop gopurs_runtime.Value, dictGenericRing1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericRing_0 gopurs_runtime.Value = dictGenericRing_0_loop
_ = dictGenericRing_0
var dictGenericRing1_1 gopurs_runtime.Value = dictGenericRing1_1_loop
_ = dictGenericRing1_1
return gopurs_runtime.RecordDict1("genericSub'", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2309107923, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericRing_0, "genericSub'"), (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product)(v_2.UnsafePtr).V0, (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product)(v1_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericRing1_1, "genericSub'"), (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product)(v_2.UnsafePtr).V1, (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product)(v1_3.UnsafePtr).V1)})}
}))
}


