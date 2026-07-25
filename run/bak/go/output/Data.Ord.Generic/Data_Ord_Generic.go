package Data_Ord_Generic

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Generic_Rep "gopurs/output/Data.Generic.Rep"
)

var cache_genericOrdNoConstructors gopurs_runtime.Value
var once_genericOrdNoConstructors sync.Once
func Get_genericOrdNoConstructors() gopurs_runtime.Value {
	once_genericOrdNoConstructors.Do(func() {
		cache_genericOrdNoConstructors = gopurs_runtime.RecordDict1("genericCompare'", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}
}))
	})
	return cache_genericOrdNoConstructors
}

var cache_genericOrdNoArguments gopurs_runtime.Value
var once_genericOrdNoArguments sync.Once
func Get_genericOrdNoArguments() gopurs_runtime.Value {
	once_genericOrdNoArguments.Do(func() {
		cache_genericOrdNoArguments = gopurs_runtime.RecordDict1("genericCompare'", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}
}))
	})
	return cache_genericOrdNoArguments
}

var cache_genericOrdArgument gopurs_runtime.Value
var once_genericOrdArgument sync.Once
func Get_genericOrdArgument() gopurs_runtime.Value {
	once_genericOrdArgument.Do(func() {
		cache_genericOrdArgument = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericOrdArgument(dictOrd_0_box)
})
	})
	return cache_genericOrdArgument
}

var cache_genericCompare_prime gopurs_runtime.Value
var once_genericCompare_prime sync.Once
func Get_genericCompare_prime() gopurs_runtime.Value {
	once_genericCompare_prime.Do(func() {
		cache_genericCompare_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericCompare_prime(dict_0_box)
})
	})
	return cache_genericCompare_prime
}

var cache_genericOrdConstructor gopurs_runtime.Value
var once_genericOrdConstructor sync.Once
func Get_genericOrdConstructor() gopurs_runtime.Value {
	once_genericOrdConstructor.Do(func() {
		cache_genericOrdConstructor = gopurs_runtime.Func(func(dictGenericOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericOrdConstructor(dictGenericOrd_0_box)
})
	})
	return cache_genericOrdConstructor
}

var cache_genericOrdProduct gopurs_runtime.Value
var once_genericOrdProduct sync.Once
func Get_genericOrdProduct() gopurs_runtime.Value {
	once_genericOrdProduct.Do(func() {
		cache_genericOrdProduct = gopurs_runtime.Func2(func(dictGenericOrd_0_box gopurs_runtime.Value, dictGenericOrd1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericOrdProduct(dictGenericOrd_0_box, dictGenericOrd1_1_box)
})
	})
	return cache_genericOrdProduct
}

var cache_genericOrdSum gopurs_runtime.Value
var once_genericOrdSum sync.Once
func Get_genericOrdSum() gopurs_runtime.Value {
	once_genericOrdSum.Do(func() {
		cache_genericOrdSum = gopurs_runtime.Func2(func(dictGenericOrd_0_box gopurs_runtime.Value, dictGenericOrd1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericOrdSum(dictGenericOrd_0_box, dictGenericOrd1_1_box)
})
	})
	return cache_genericOrdSum
}

var cache_genericCompare gopurs_runtime.Value
var once_genericCompare sync.Once
func Get_genericCompare() gopurs_runtime.Value {
	once_genericCompare.Do(func() {
		cache_genericCompare = gopurs_runtime.Func4(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericOrd_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericCompare(dictGeneric_0_box, dictGenericOrd_1_box, x_2_box, y_3_box)
})
	})
	return cache_genericCompare
}

func Call_genericOrdArgument(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.RecordDict1("genericCompare'", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictOrd_0.UnsafePtr)).V0, v_1, v1_2)
}))
}

func Call_genericCompare_prime(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return ((*gopurs_runtime.RecordData1)(dict_0.UnsafePtr)).V0
}

func Call_genericOrdConstructor(dictGenericOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericOrd_0 gopurs_runtime.Value = dictGenericOrd_0_loop
_ = dictGenericOrd_0
return gopurs_runtime.RecordDict1("genericCompare'", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictGenericOrd_0.UnsafePtr)).V0, v_1, v1_2)
}))
}

func Call_genericOrdProduct(dictGenericOrd_0_loop gopurs_runtime.Value, dictGenericOrd1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericOrd_0 gopurs_runtime.Value = dictGenericOrd_0_loop
_ = dictGenericOrd_0
var dictGenericOrd1_1 gopurs_runtime.Value = dictGenericOrd1_1_loop
_ = dictGenericOrd1_1
return gopurs_runtime.RecordDict1("genericCompare'", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
v2_4_0 := gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictGenericOrd_0.UnsafePtr)).V0, (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product)(v_2.UnsafePtr).V0, (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product)(v1_3.UnsafePtr).V0)
_ = v2_4_0
var __t1 gopurs_runtime.Value
{
if (v2_4_0.Type == 9 && v2_4_0.IntVal == 902936544) {
__t1 = gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictGenericOrd1_1.UnsafePtr)).V0, (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product)(v_2.UnsafePtr).V1, (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product)(v1_3.UnsafePtr).V1)
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
if (v_2.Type == 9 && v_2.IntVal == 3478632216) {
var __t1 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 3478632216) {
__t1 = gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictGenericOrd_0.UnsafePtr)).V0, (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Inl)(v_2.UnsafePtr).V0, (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Inl)(v1_3.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 492034566) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}
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
if (v_2.Type == 9 && v_2.IntVal == 492034566) {
var __t2 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 492034566) {
__t2 = gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictGenericOrd1_1.UnsafePtr)).V0, (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Inr)(v_2.UnsafePtr).V0, (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Inr)(v1_3.UnsafePtr).V0)
goto end_branch_2
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 3478632216) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}
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
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictGenericOrd_1.UnsafePtr)).V0, gopurs_runtime.Apply(((*gopurs_runtime.RecordData2)(dictGeneric_0.UnsafePtr)).V0, x_2), gopurs_runtime.Apply(((*gopurs_runtime.RecordData2)(dictGeneric_0.UnsafePtr)).V0, y_3))
}


