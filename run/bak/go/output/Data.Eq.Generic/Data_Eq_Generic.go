package Data_Eq_Generic

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Generic_Rep "gopurs/output/Data.Generic.Rep"
)

var cache_genericEqNoConstructors gopurs_runtime.Value
var once_genericEqNoConstructors sync.Once
func Get_genericEqNoConstructors() gopurs_runtime.Value {
	once_genericEqNoConstructors.Do(func() {
		cache_genericEqNoConstructors = gopurs_runtime.RecordDict1("genericEq'", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})
}))
	})
	return cache_genericEqNoConstructors
}

var cache_genericEqNoArguments gopurs_runtime.Value
var once_genericEqNoArguments sync.Once
func Get_genericEqNoArguments() gopurs_runtime.Value {
	once_genericEqNoArguments.Do(func() {
		cache_genericEqNoArguments = gopurs_runtime.RecordDict1("genericEq'", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})
}))
	})
	return cache_genericEqNoArguments
}

var cache_genericEqArgument gopurs_runtime.Value
var once_genericEqArgument sync.Once
func Get_genericEqArgument() gopurs_runtime.Value {
	once_genericEqArgument.Do(func() {
		cache_genericEqArgument = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericEqArgument(dictEq_0_box)
})
	})
	return cache_genericEqArgument
}

var cache_genericEq_prime gopurs_runtime.Value
var once_genericEq_prime sync.Once
func Get_genericEq_prime() gopurs_runtime.Value {
	once_genericEq_prime.Do(func() {
		cache_genericEq_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericEq_prime(dict_0_box)
})
	})
	return cache_genericEq_prime
}

var cache_genericEq_prime__gopurs_runtime_Value_2165983829 gopurs_runtime.Value
var once_genericEq_prime__gopurs_runtime_Value_2165983829 sync.Once
func Get_genericEq_prime__gopurs_runtime_Value_2165983829() gopurs_runtime.Value {
	once_genericEq_prime__gopurs_runtime_Value_2165983829.Do(func() {
		cache_genericEq_prime__gopurs_runtime_Value_2165983829 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericEq_prime__gopurs_runtime_Value_2165983829(dict_0_box)
})
	})
	return cache_genericEq_prime__gopurs_runtime_Value_2165983829
}

var cache_genericEqConstructor gopurs_runtime.Value
var once_genericEqConstructor sync.Once
func Get_genericEqConstructor() gopurs_runtime.Value {
	once_genericEqConstructor.Do(func() {
		cache_genericEqConstructor = gopurs_runtime.Func(func(dictGenericEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericEqConstructor(dictGenericEq_0_box)
})
	})
	return cache_genericEqConstructor
}

var cache_genericEqProduct gopurs_runtime.Value
var once_genericEqProduct sync.Once
func Get_genericEqProduct() gopurs_runtime.Value {
	once_genericEqProduct.Do(func() {
		cache_genericEqProduct = gopurs_runtime.Func2(func(dictGenericEq_0_box gopurs_runtime.Value, dictGenericEq1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericEqProduct(dictGenericEq_0_box, dictGenericEq1_1_box)
})
	})
	return cache_genericEqProduct
}

var cache_genericEqSum gopurs_runtime.Value
var once_genericEqSum sync.Once
func Get_genericEqSum() gopurs_runtime.Value {
	once_genericEqSum.Do(func() {
		cache_genericEqSum = gopurs_runtime.Func2(func(dictGenericEq_0_box gopurs_runtime.Value, dictGenericEq1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericEqSum(dictGenericEq_0_box, dictGenericEq1_1_box)
})
	})
	return cache_genericEqSum
}

var cache_genericEq gopurs_runtime.Value
var once_genericEq sync.Once
func Get_genericEq() gopurs_runtime.Value {
	once_genericEq.Do(func() {
		cache_genericEq = gopurs_runtime.Func4(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericEq_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_genericEq(dictGeneric_0_box, dictGenericEq_1_box, x_2_box, y_3_box))
})
	})
	return cache_genericEq
}

func Call_genericEqArgument(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.RecordDict1("genericEq'", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), v_1, v1_2).IntVal) != (0))
})
}))
}

func Call_genericEq_prime(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericEq'")
}

func Call_genericEq_prime__gopurs_runtime_Value_2165983829(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericEq'")
}

func Call_genericEqConstructor(dictGenericEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericEq_0 gopurs_runtime.Value = dictGenericEq_0_loop
_ = dictGenericEq_0
return gopurs_runtime.RecordDict1("genericEq'", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericEq_0, "genericEq'"), v_1, v1_2).IntVal) != (0))
})
}))
}

func Call_genericEqProduct(dictGenericEq_0_loop gopurs_runtime.Value, dictGenericEq1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericEq_0 gopurs_runtime.Value = dictGenericEq_0_loop
_ = dictGenericEq_0
var dictGenericEq1_1 gopurs_runtime.Value = dictGenericEq1_1_loop
_ = dictGenericEq1_1
return gopurs_runtime.RecordDict1("genericEq'", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericEq_0, "genericEq'"), (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericEq1_1, "genericEq'"), (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1)).IntVal) != (0))
})
}))
}

func Call_genericEqSum(dictGenericEq_0_loop gopurs_runtime.Value, dictGenericEq1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericEq_0 gopurs_runtime.Value = dictGenericEq_0_loop
_ = dictGenericEq_0
var dictGenericEq1_1 gopurs_runtime.Value = dictGenericEq1_1_loop
_ = dictGenericEq1_1
return gopurs_runtime.RecordDict1("genericEq'", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 3478632216) {
var __t1 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 3478632216) {
__t1 = gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericEq_0, "genericEq'"), (*pkg_Data_Generic_Rep.Constructor_Inl[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*pkg_Data_Generic_Rep.Constructor_Inl[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0).IntVal) != (0))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Bool(false)
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
if ((v_2.Type == 9 && v_2.IntVal == 492034566)) && ((v1_3.Type == 9 && v1_3.IntVal == 492034566)) {
__t0 = gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericEq1_1, "genericEq'"), (*pkg_Data_Generic_Rep.Constructor_Inr[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*pkg_Data_Generic_Rep.Constructor_Inr[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0).IntVal) != (0))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Bool(false)
}
end_branch_0:
return gopurs_runtime.Bool((__t0.IntVal) != (0))
})
}))
}

func Call_genericEq(dictGeneric_0_loop gopurs_runtime.Value, dictGenericEq_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) bool {
var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericEq_1 gopurs_runtime.Value = dictGenericEq_1_loop
_ = dictGenericEq_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericEq_1, "genericEq'"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "from"), x_2), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "from"), y_3)).IntVal) != (0)
}


