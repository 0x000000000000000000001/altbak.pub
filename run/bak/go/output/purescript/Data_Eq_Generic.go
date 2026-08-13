package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Eq_Generic_GenericEq_dollarDict gopurs_runtime.Value
var once_Data_Eq_Generic_GenericEq_dollarDict sync.Once
func Get_Data_Eq_Generic_GenericEq_dollarDict() gopurs_runtime.Value {
	once_Data_Eq_Generic_GenericEq_dollarDict.Do(func() {
		cache_Data_Eq_Generic_GenericEq_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_Generic_GenericEq_dollarDict(x_0_box)
})
	})
	return cache_Data_Eq_Generic_GenericEq_dollarDict
}

var cache_Data_Eq_Generic_genericEqNoConstructors gopurs_runtime.Value
var once_Data_Eq_Generic_genericEqNoConstructors sync.Once
func Get_Data_Eq_Generic_genericEqNoConstructors() gopurs_runtime.Value {
	once_Data_Eq_Generic_genericEqNoConstructors.Do(func() {
		cache_Data_Eq_Generic_genericEqNoConstructors = gopurs_runtime.RecordDict1("genericEq'", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})
}))
	})
	return cache_Data_Eq_Generic_genericEqNoConstructors
}

var cache_Data_Eq_Generic_genericEqNoArguments gopurs_runtime.Value
var once_Data_Eq_Generic_genericEqNoArguments sync.Once
func Get_Data_Eq_Generic_genericEqNoArguments() gopurs_runtime.Value {
	once_Data_Eq_Generic_genericEqNoArguments.Do(func() {
		cache_Data_Eq_Generic_genericEqNoArguments = gopurs_runtime.RecordDict1("genericEq'", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})
}))
	})
	return cache_Data_Eq_Generic_genericEqNoArguments
}

var cache_Data_Eq_Generic_genericEqArgument gopurs_runtime.Value
var once_Data_Eq_Generic_genericEqArgument sync.Once
func Get_Data_Eq_Generic_genericEqArgument() gopurs_runtime.Value {
	once_Data_Eq_Generic_genericEqArgument.Do(func() {
		cache_Data_Eq_Generic_genericEqArgument = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_Generic_genericEqArgument(dictEq_0_box)
})
	})
	return cache_Data_Eq_Generic_genericEqArgument
}

var cache_Data_Eq_Generic_genericEq_prime gopurs_runtime.Value
var once_Data_Eq_Generic_genericEq_prime sync.Once
func Get_Data_Eq_Generic_genericEq_prime() gopurs_runtime.Value {
	once_Data_Eq_Generic_genericEq_prime.Do(func() {
		cache_Data_Eq_Generic_genericEq_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_Generic_genericEq_prime(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Generic_GenericEq[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Eq_Generic_genericEq_prime
}

var cache_Data_Eq_Generic_genericEqConstructor gopurs_runtime.Value
var once_Data_Eq_Generic_genericEqConstructor sync.Once
func Get_Data_Eq_Generic_genericEqConstructor() gopurs_runtime.Value {
	once_Data_Eq_Generic_genericEqConstructor.Do(func() {
		cache_Data_Eq_Generic_genericEqConstructor = gopurs_runtime.Func(func(dictGenericEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_Generic_genericEqConstructor(dictGenericEq_0_box)
})
	})
	return cache_Data_Eq_Generic_genericEqConstructor
}

var cache_Data_Eq_Generic_genericEqProduct gopurs_runtime.Value
var once_Data_Eq_Generic_genericEqProduct sync.Once
func Get_Data_Eq_Generic_genericEqProduct() gopurs_runtime.Value {
	once_Data_Eq_Generic_genericEqProduct.Do(func() {
		cache_Data_Eq_Generic_genericEqProduct = gopurs_runtime.Func2(func(dictGenericEq_0_box gopurs_runtime.Value, dictGenericEq1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_Generic_genericEqProduct(dictGenericEq_0_box, dictGenericEq1_1_box)
})
	})
	return cache_Data_Eq_Generic_genericEqProduct
}

var cache_Data_Eq_Generic_genericEqSum gopurs_runtime.Value
var once_Data_Eq_Generic_genericEqSum sync.Once
func Get_Data_Eq_Generic_genericEqSum() gopurs_runtime.Value {
	once_Data_Eq_Generic_genericEqSum.Do(func() {
		cache_Data_Eq_Generic_genericEqSum = gopurs_runtime.Func2(func(dictGenericEq_0_box gopurs_runtime.Value, dictGenericEq1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_Generic_genericEqSum(dictGenericEq_0_box, dictGenericEq1_1_box)
})
	})
	return cache_Data_Eq_Generic_genericEqSum
}

var cache_Data_Eq_Generic_genericEq gopurs_runtime.Value
var once_Data_Eq_Generic_genericEq sync.Once
func Get_Data_Eq_Generic_genericEq() gopurs_runtime.Value {
	once_Data_Eq_Generic_genericEq.Do(func() {
		cache_Data_Eq_Generic_genericEq = gopurs_runtime.Func4(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericEq_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Eq_Generic_genericEq(gopurs_runtime.CoerceToStruct[Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value]](dictGeneric_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Generic_GenericEq[gopurs_runtime.Value]](dictGenericEq_1_box), x_2_box, y_3_box))
})
	})
	return cache_Data_Eq_Generic_genericEq
}

var cache_Data_Eq_Generic_genericEq_prime__1422345493 gopurs_runtime.Value
var once_Data_Eq_Generic_genericEq_prime__1422345493 sync.Once
func Get_Data_Eq_Generic_genericEq_prime__1422345493() gopurs_runtime.Value {
	once_Data_Eq_Generic_genericEq_prime__1422345493.Do(func() {
		cache_Data_Eq_Generic_genericEq_prime__1422345493 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_Generic_genericEq_prime__1422345493(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Generic_GenericEq[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Eq_Generic_genericEq_prime__1422345493
}

type Constructor_Data_Eq_Generic_GenericEq[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[106035173] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Eq_Generic_GenericEq[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "genericEq'": return gopurs_runtime.Box(c.V0)
		default: panic("Key not found in dictionary Constructor_Data_Eq_Generic_GenericEq: " + key)
		}
	}
}


func Call_Data_Eq_Generic_GenericEq_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Eq_Generic_genericEqArgument(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.RecordDict1("genericEq'", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), v_1, v1_2).IntVal) != (0))
})
}))
}

func Call_Data_Eq_Generic_genericEq_prime(dict_0_loop *Constructor_Data_Eq_Generic_GenericEq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Eq_Generic_GenericEq[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Eq_Generic_genericEqConstructor(dictGenericEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericEq_0 gopurs_runtime.Value = dictGenericEq_0_loop
_ = dictGenericEq_0
return gopurs_runtime.RecordDict1("genericEq'", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericEq_0, "genericEq'"), v_1, v1_2).IntVal) != (0))
})
}))
}

func Call_Data_Eq_Generic_genericEqProduct(dictGenericEq_0_loop gopurs_runtime.Value, dictGenericEq1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericEq_0 gopurs_runtime.Value = dictGenericEq_0_loop
_ = dictGenericEq_0
var dictGenericEq1_1 gopurs_runtime.Value = dictGenericEq1_1_loop
_ = dictGenericEq1_1
return gopurs_runtime.RecordDict1("genericEq'", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericEq_0, "genericEq'"), (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericEq1_1, "genericEq'"), (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1).IntVal) != (0)))
})
}))
}

func Call_Data_Eq_Generic_genericEqSum(dictGenericEq_0_loop gopurs_runtime.Value, dictGenericEq1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericEq_0 gopurs_runtime.Value = dictGenericEq_0_loop
_ = dictGenericEq_0
var dictGenericEq1_1 gopurs_runtime.Value = dictGenericEq1_1_loop
_ = dictGenericEq1_1
return gopurs_runtime.RecordDict1("genericEq'", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 bool
{
if (v_2.Type == 9 && v_2.IntVal == 3478632216) {
var __t0 bool
{
if (v1_3.Type == 9 && v1_3.IntVal == 3478632216) {
__t0 = (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericEq_0, "genericEq'"), (*Constructor_Data_Generic_Rep_Inl[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*Constructor_Data_Generic_Rep_Inl[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0).IntVal) != (0)
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
__t1 = __t0
goto end_branch_1
} else {

}
}
{
if ((v_2.Type == 9 && v_2.IntVal == 492034566)) && ((v1_3.Type == 9 && v1_3.IntVal == 492034566)) {
__t1 = (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericEq1_1, "genericEq'"), (*Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0).IntVal) != (0)
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
return gopurs_runtime.Bool(__t1)
})
}))
}

func Call_Data_Eq_Generic_genericEq(dictGeneric_0_loop *Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value], dictGenericEq_1_loop *Constructor_Data_Eq_Generic_GenericEq[gopurs_runtime.Value], x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) bool {
var dictGeneric_0 *Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value] = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericEq_1 *Constructor_Data_Eq_Generic_GenericEq[gopurs_runtime.Value] = dictGenericEq_1_loop
_ = dictGenericEq_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return (gopurs_runtime.Apply3(Get_Data_Eq_Generic_genericEq_prime__1422345493(), gopurs_runtime.Value{Type: 9, IntVal: 106035173, UnsafePtr: unsafe.Pointer(dictGenericEq_1)}, gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V0), x_2), gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V0), y_3)).IntVal) != (0)
}

func Call_Data_Eq_Generic_genericEq_prime__1422345493(dict_0_loop *Constructor_Data_Eq_Generic_GenericEq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Eq_Generic_GenericEq[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}


