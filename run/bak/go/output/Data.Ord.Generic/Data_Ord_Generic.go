package Data_Ord_Generic

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Ordering "gopurs/output/Data.Ordering"
	pkg_Data_Generic_Rep "gopurs/output/Data.Generic.Rep"
	unsafe "unsafe"
)

var genericOrdNoConstructors gopurs_runtime.Value
var once_genericOrdNoConstructors sync.Once
func Get_genericOrdNoConstructors() gopurs_runtime.Value {
	once_genericOrdNoConstructors.Do(func() {
		genericOrdNoConstructors = gopurs_runtime.RecordDict1("genericCompare'", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1111389260, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
}))
	})
	return genericOrdNoConstructors
}

var genericOrdNoArguments gopurs_runtime.Value
var once_genericOrdNoArguments sync.Once
func Get_genericOrdNoArguments() gopurs_runtime.Value {
	once_genericOrdNoArguments.Do(func() {
		genericOrdNoArguments = gopurs_runtime.RecordDict1("genericCompare'", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1111389260, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
}))
	})
	return genericOrdNoArguments
}

var genericOrdArgument gopurs_runtime.Value
var once_genericOrdArgument sync.Once
func Get_genericOrdArgument() gopurs_runtime.Value {
	once_genericOrdArgument.Do(func() {
		genericOrdArgument = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.RecordDict1("genericCompare'", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), v_1, v1_2)
}))
}()
})
	})
	return genericOrdArgument
}

var genericCompare_prime gopurs_runtime.Value
var once_genericCompare_prime sync.Once
func Get_genericCompare_prime() gopurs_runtime.Value {
	once_genericCompare_prime.Do(func() {
		genericCompare_prime = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericCompare'")
}()
})
	})
	return genericCompare_prime
}

var genericOrdConstructor gopurs_runtime.Value
var once_genericOrdConstructor sync.Once
func Get_genericOrdConstructor() gopurs_runtime.Value {
	once_genericOrdConstructor.Do(func() {
		genericOrdConstructor = gopurs_runtime.Func(func(dictGenericOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictGenericOrd_0 gopurs_runtime.Value = dictGenericOrd_0_loop
_ = dictGenericOrd_0
return gopurs_runtime.RecordDict1("genericCompare'", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericOrd_0, "genericCompare'"), v_1, v1_2)
}))
}()
})
	})
	return genericOrdConstructor
}

var genericOrdProduct gopurs_runtime.Value
var once_genericOrdProduct sync.Once
func Get_genericOrdProduct() gopurs_runtime.Value {
	once_genericOrdProduct.Do(func() {
		genericOrdProduct = gopurs_runtime.Func2(func(dictGenericOrd_0_box gopurs_runtime.Value, dictGenericOrd1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericOrdProduct(dictGenericOrd_0_box, dictGenericOrd1_1_box)
})
	})
	return genericOrdProduct
}

var genericOrdSum gopurs_runtime.Value
var once_genericOrdSum sync.Once
func Get_genericOrdSum() gopurs_runtime.Value {
	once_genericOrdSum.Do(func() {
		genericOrdSum = gopurs_runtime.Func2(func(dictGenericOrd_0_box gopurs_runtime.Value, dictGenericOrd1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericOrdSum(dictGenericOrd_0_box, dictGenericOrd1_1_box)
})
	})
	return genericOrdSum
}

var genericCompare gopurs_runtime.Value
var once_genericCompare sync.Once
func Get_genericCompare() gopurs_runtime.Value {
	once_genericCompare.Do(func() {
		genericCompare = gopurs_runtime.Func4(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericOrd_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericCompare(dictGeneric_0_box, dictGenericOrd_1_box, x_2_box, y_3_box)
})
	})
	return genericCompare
}

func Call_genericOrdProduct(dictGenericOrd_0_loop gopurs_runtime.Value, dictGenericOrd1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericOrd_0 gopurs_runtime.Value = dictGenericOrd_0_loop
_ = dictGenericOrd_0
var dictGenericOrd1_1 gopurs_runtime.Value = dictGenericOrd1_1_loop
_ = dictGenericOrd1_1
return gopurs_runtime.RecordDict1("genericCompare'", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
v2_4_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericOrd_0, "genericCompare'"), (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product)(v_2.UnsafePtr).V0, (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product)(v1_3.UnsafePtr).V0)
_ = v2_4_0
var __t1 gopurs_runtime.Value
{
if (v2_4_0.Type == 9 && v2_4_0.IntVal == 1111389260) {
__t1 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericOrd1_1, "genericCompare'"), (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product)(v_2.UnsafePtr).V1, (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product)(v1_3.UnsafePtr).V1)
goto end_branch_1
} else {

}
}
{
__t1 = v2_4_0
}
end_branch_1:
return __t1
}))
}

func Call_genericOrdSum(dictGenericOrd_0_loop gopurs_runtime.Value, dictGenericOrd1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericOrd_0 gopurs_runtime.Value = dictGenericOrd_0_loop
_ = dictGenericOrd_0
var dictGenericOrd1_1 gopurs_runtime.Value = dictGenericOrd1_1_loop
_ = dictGenericOrd1_1
return gopurs_runtime.RecordDict1("genericCompare'", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 164387955) {
var __t1 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 164387955) {
__t1 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericOrd_0, "genericCompare'"), (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Inl)(v_2.UnsafePtr).V0, (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Inl)(v1_3.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 4051932077) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3866105248, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 4051932077) {
var __t2 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 4051932077) {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericOrd1_1, "genericCompare'"), (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Inr)(v_2.UnsafePtr).V0, (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Inr)(v1_3.UnsafePtr).V0)
goto end_branch_2
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 164387955) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 2098047435, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t0 = __t2
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
}

func Call_genericCompare(dictGeneric_0_loop gopurs_runtime.Value, dictGenericOrd_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericOrd_1 gopurs_runtime.Value = dictGenericOrd_1_loop
_ = dictGenericOrd_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericOrd_1, "genericCompare'"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "from"), x_2), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "from"), y_3))
}


