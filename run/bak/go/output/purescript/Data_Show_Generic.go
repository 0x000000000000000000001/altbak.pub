package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Show_Generic_GenericShowArgs_dollarDict gopurs_runtime.Value
var once_Data_Show_Generic_GenericShowArgs_dollarDict sync.Once
func Get_Data_Show_Generic_GenericShowArgs_dollarDict() gopurs_runtime.Value {
	once_Data_Show_Generic_GenericShowArgs_dollarDict.Do(func() {
		cache_Data_Show_Generic_GenericShowArgs_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_Generic_GenericShowArgs_dollarDict(x_0_box)
})
	})
	return cache_Data_Show_Generic_GenericShowArgs_dollarDict
}

var cache_Data_Show_Generic_GenericShow_dollarDict gopurs_runtime.Value
var once_Data_Show_Generic_GenericShow_dollarDict sync.Once
func Get_Data_Show_Generic_GenericShow_dollarDict() gopurs_runtime.Value {
	once_Data_Show_Generic_GenericShow_dollarDict.Do(func() {
		cache_Data_Show_Generic_GenericShow_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_Generic_GenericShow_dollarDict(x_0_box)
})
	})
	return cache_Data_Show_Generic_GenericShow_dollarDict
}

var cache_Data_Show_Generic_genericShowArgsNoArguments gopurs_runtime.Value
var once_Data_Show_Generic_genericShowArgsNoArguments sync.Once
func Get_Data_Show_Generic_genericShowArgsNoArguments() gopurs_runtime.Value {
	once_Data_Show_Generic_genericShowArgsNoArguments.Do(func() {
		cache_Data_Show_Generic_genericShowArgsNoArguments = gopurs_runtime.RecordDict1("genericShowArgs", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := func() []string {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr { unboxed[i] = v.StrVal() }
					return unboxed
				}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Str(v) }
					return gopurs_runtime.Array(boxed)
				}()
}))
	})
	return cache_Data_Show_Generic_genericShowArgsNoArguments
}

var cache_Data_Show_Generic_genericShowArgsArgument gopurs_runtime.Value
var once_Data_Show_Generic_genericShowArgsArgument sync.Once
func Get_Data_Show_Generic_genericShowArgsArgument() gopurs_runtime.Value {
	once_Data_Show_Generic_genericShowArgsArgument.Do(func() {
		cache_Data_Show_Generic_genericShowArgsArgument = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_Generic_genericShowArgsArgument(dictShow_0_box)
})
	})
	return cache_Data_Show_Generic_genericShowArgsArgument
}

var cache_Data_Show_Generic_genericShowArgs gopurs_runtime.Value
var once_Data_Show_Generic_genericShowArgs sync.Once
func Get_Data_Show_Generic_genericShowArgs() gopurs_runtime.Value {
	once_Data_Show_Generic_genericShowArgs.Do(func() {
		cache_Data_Show_Generic_genericShowArgs = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_Generic_genericShowArgs(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Generic_GenericShowArgs](dict_0_box))
})
	})
	return cache_Data_Show_Generic_genericShowArgs
}

var cache_Data_Show_Generic_genericShowArgsProduct gopurs_runtime.Value
var once_Data_Show_Generic_genericShowArgsProduct sync.Once
func Get_Data_Show_Generic_genericShowArgsProduct() gopurs_runtime.Value {
	once_Data_Show_Generic_genericShowArgsProduct.Do(func() {
		cache_Data_Show_Generic_genericShowArgsProduct = gopurs_runtime.Func2(func(dictGenericShowArgs_0_box gopurs_runtime.Value, dictGenericShowArgs1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_Generic_genericShowArgsProduct(dictGenericShowArgs_0_box, dictGenericShowArgs1_1_box)
})
	})
	return cache_Data_Show_Generic_genericShowArgsProduct
}

var cache_Data_Show_Generic_genericShowConstructor gopurs_runtime.Value
var once_Data_Show_Generic_genericShowConstructor sync.Once
func Get_Data_Show_Generic_genericShowConstructor() gopurs_runtime.Value {
	once_Data_Show_Generic_genericShowConstructor.Do(func() {
		cache_Data_Show_Generic_genericShowConstructor = gopurs_runtime.Func2(func(dictGenericShowArgs_0_box gopurs_runtime.Value, dictIsSymbol_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_Generic_genericShowConstructor(dictGenericShowArgs_0_box, dictIsSymbol_1_box)
})
	})
	return cache_Data_Show_Generic_genericShowConstructor
}

var cache_Data_Show_Generic_genericShow_prime gopurs_runtime.Value
var once_Data_Show_Generic_genericShow_prime sync.Once
func Get_Data_Show_Generic_genericShow_prime() gopurs_runtime.Value {
	once_Data_Show_Generic_genericShow_prime.Do(func() {
		cache_Data_Show_Generic_genericShow_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_Generic_genericShow_prime(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Generic_GenericShow](dict_0_box))
})
	})
	return cache_Data_Show_Generic_genericShow_prime
}

var cache_Data_Show_Generic_genericShowNoConstructors gopurs_runtime.Value
var once_Data_Show_Generic_genericShowNoConstructors sync.Once
func Get_Data_Show_Generic_genericShowNoConstructors() gopurs_runtime.Value {
	once_Data_Show_Generic_genericShowNoConstructors.Do(func() {
		cache_Data_Show_Generic_genericShowNoConstructors = gopurs_runtime.RecordDict1("genericShow'", gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_Generic_genericShow_prime__1538939821(), a_0).StrVal())
}))
	})
	return cache_Data_Show_Generic_genericShowNoConstructors
}

var cache_Data_Show_Generic_genericShowSum gopurs_runtime.Value
var once_Data_Show_Generic_genericShowSum sync.Once
func Get_Data_Show_Generic_genericShowSum() gopurs_runtime.Value {
	once_Data_Show_Generic_genericShowSum.Do(func() {
		cache_Data_Show_Generic_genericShowSum = gopurs_runtime.Func2(func(dictGenericShow_0_box gopurs_runtime.Value, dictGenericShow1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_Generic_genericShowSum(dictGenericShow_0_box, dictGenericShow1_1_box)
})
	})
	return cache_Data_Show_Generic_genericShowSum
}

var cache_Data_Show_Generic_genericShow gopurs_runtime.Value
var once_Data_Show_Generic_genericShow sync.Once
func Get_Data_Show_Generic_genericShow() gopurs_runtime.Value {
	once_Data_Show_Generic_genericShow.Do(func() {
		cache_Data_Show_Generic_genericShow = gopurs_runtime.Func3(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericShow_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_Show_Generic_genericShow(gopurs_runtime.CoerceToStruct[Constructor_Data_Generic_Rep_Generic](dictGeneric_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Generic_GenericShow](dictGenericShow_1_box), x_2_box))
})
	})
	return cache_Data_Show_Generic_genericShow
}

var cache_Data_Show_Generic_genericShow_prime__3102496331 gopurs_runtime.Value
var once_Data_Show_Generic_genericShow_prime__3102496331 sync.Once
func Get_Data_Show_Generic_genericShow_prime__3102496331() gopurs_runtime.Value {
	once_Data_Show_Generic_genericShow_prime__3102496331.Do(func() {
		cache_Data_Show_Generic_genericShow_prime__3102496331 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_Generic_genericShow_prime__3102496331(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Generic_GenericShow](dict_0_box))
})
	})
	return cache_Data_Show_Generic_genericShow_prime__3102496331
}

var cache_Data_Show_Generic_genericShow_prime__2423230701 gopurs_runtime.Value
var once_Data_Show_Generic_genericShow_prime__2423230701 sync.Once
func Get_Data_Show_Generic_genericShow_prime__2423230701() gopurs_runtime.Value {
	once_Data_Show_Generic_genericShow_prime__2423230701.Do(func() {
		cache_Data_Show_Generic_genericShow_prime__2423230701 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_Generic_genericShow_prime__2423230701(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Generic_GenericShow](dict_0_box))
})
	})
	return cache_Data_Show_Generic_genericShow_prime__2423230701
}

var cache_Data_Show_Generic_genericShow_prime__1538939821 gopurs_runtime.Value
var once_Data_Show_Generic_genericShow_prime__1538939821 sync.Once
func Get_Data_Show_Generic_genericShow_prime__1538939821() gopurs_runtime.Value {
	once_Data_Show_Generic_genericShow_prime__1538939821.Do(func() {
		cache_Data_Show_Generic_genericShow_prime__1538939821 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_Show_Generic_genericShow_prime__1538939821(a_0_box))
})
	})
	return cache_Data_Show_Generic_genericShow_prime__1538939821
}

var cache_Data_Show_Generic_genericShowArgs__157062380 gopurs_runtime.Value
var once_Data_Show_Generic_genericShowArgs__157062380 sync.Once
func Get_Data_Show_Generic_genericShowArgs__157062380() gopurs_runtime.Value {
	once_Data_Show_Generic_genericShowArgs__157062380.Do(func() {
		cache_Data_Show_Generic_genericShowArgs__157062380 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_Generic_genericShowArgs__157062380(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Generic_GenericShowArgs](dict_0_box))
})
	})
	return cache_Data_Show_Generic_genericShowArgs__157062380
}

type Constructor_Data_Show_Generic_GenericShowArgs struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1968625250] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Show_Generic_GenericShowArgs)(ptr)
		_ = c
		switch key {
		case "genericShowArgs": return gopurs_runtime.Box(c.V0)
		default: panic("Key not found in dictionary Constructor_Data_Show_Generic_GenericShowArgs: " + key)
		}
	}
}


type Constructor_Data_Show_Generic_GenericShow struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2730968613] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Show_Generic_GenericShow)(ptr)
		_ = c
		switch key {
		case "genericShow'": return gopurs_runtime.Box(c.V0)
		default: panic("Key not found in dictionary Constructor_Data_Show_Generic_GenericShow: " + key)
		}
	}
}


func Call_Data_Show_Generic_GenericShowArgs_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Show_Generic_GenericShow_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Show_Generic_genericShowArgsArgument(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("genericShowArgs", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := []string{gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal()}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Str(v) }
					return gopurs_runtime.Array(boxed)
				}()
}))
}

func Call_Data_Show_Generic_genericShowArgs(dict_0_loop *Constructor_Data_Show_Generic_GenericShowArgs) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Show_Generic_GenericShowArgs = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Show_Generic_genericShowArgsProduct(dictGenericShowArgs_0_loop gopurs_runtime.Value, dictGenericShowArgs1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericShowArgs_0 gopurs_runtime.Value = dictGenericShowArgs_0_loop
_ = dictGenericShowArgs_0
var dictGenericShowArgs1_1 gopurs_runtime.Value = dictGenericShowArgs1_1_loop
_ = dictGenericShowArgs1_1
return gopurs_runtime.RecordDict1("genericShowArgs", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := func() []string {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericShowArgs_0, "genericShowArgs"), (*Constructor_Data_Generic_Rep_Product)(v_2.UnsafePtr).V0), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericShowArgs1_1, "genericShowArgs"), (*Constructor_Data_Generic_Rep_Product)(v_2.UnsafePtr).V1)).UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr { unboxed[i] = v.StrVal() }
					return unboxed
				}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Str(v) }
					return gopurs_runtime.Array(boxed)
				}()
}))
}

func Call_Data_Show_Generic_genericShowConstructor(dictGenericShowArgs_0_loop gopurs_runtime.Value, dictIsSymbol_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericShowArgs_0 gopurs_runtime.Value = dictGenericShowArgs_0_loop
_ = dictGenericShowArgs_0
var dictIsSymbol_1 gopurs_runtime.Value = dictIsSymbol_1_loop
_ = dictIsSymbol_1
return gopurs_runtime.RecordDict1("genericShow'", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): ctor_3_0 -> string
ctor_3_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_1, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}).StrVal()
_ = ctor_3_0
// TAST (Let): v1_4_1 -> gopurs_runtime.Value
v1_4_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericShowArgs_0, "genericShowArgs"), v_2)
_ = v1_4_1
var __t2 string
{
if (gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(v1_4_1))).IntVal) == (0) {
__t2 = ctor_3_0
goto end_branch_2
} else {

}
}
{
__t2 = (("(") + (gopurs_runtime.Apply2(Get_Data_Show_Generic_intercalate(), gopurs_runtime.Str(" "), func() gopurs_runtime.Value {
					arr := func() []string {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), func() gopurs_runtime.Value {
					arr := []string{ctor_3_0}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Str(v) }
					return gopurs_runtime.Array(boxed)
				}(), v1_4_1).UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr { unboxed[i] = v.StrVal() }
					return unboxed
				}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Str(v) }
					return gopurs_runtime.Array(boxed)
				}()).StrVal())) + (")")
}
end_branch_2:
return gopurs_runtime.Str(__t2)
}))
}

func Call_Data_Show_Generic_genericShow_prime(dict_0_loop *Constructor_Data_Show_Generic_GenericShow) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Show_Generic_GenericShow = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Show_Generic_genericShowSum(dictGenericShow_0_loop gopurs_runtime.Value, dictGenericShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericShow_0 gopurs_runtime.Value = dictGenericShow_0_loop
_ = dictGenericShow_0
var dictGenericShow1_1 gopurs_runtime.Value = dictGenericShow1_1_loop
_ = dictGenericShow1_1
return gopurs_runtime.RecordDict1("genericShow'", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 string
{
if (v_2.Type == 9 && v_2.IntVal == 3478632216) {
__t0 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericShow_0, "genericShow'"), (*Constructor_Data_Generic_Rep_Inl)(v_2.UnsafePtr).V0).StrVal()
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 492034566) {
__t0 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericShow1_1, "genericShow'"), (*Constructor_Data_Generic_Rep_Inr)(v_2.UnsafePtr).V0).StrVal()
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
}
end_branch_0:
return gopurs_runtime.Str(__t0)
}))
}

func Call_Data_Show_Generic_genericShow(dictGeneric_0_loop *Constructor_Data_Generic_Rep_Generic, dictGenericShow_1_loop *Constructor_Data_Show_Generic_GenericShow, x_2_loop gopurs_runtime.Value) string {
var dictGeneric_0 *Constructor_Data_Generic_Rep_Generic = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericShow_1 *Constructor_Data_Show_Generic_GenericShow = dictGenericShow_1_loop
_ = dictGenericShow_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply2(Get_Data_Show_Generic_genericShow_prime__2423230701(), gopurs_runtime.Value{Type: 9, IntVal: 2730968613, UnsafePtr: unsafe.Pointer(dictGenericShow_1)}, gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V0), x_2)).StrVal()
}

func Call_Data_Show_Generic_genericShow_prime__3102496331(dict_0_loop *Constructor_Data_Show_Generic_GenericShow) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Show_Generic_GenericShow = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Show_Generic_genericShow_prime__2423230701(dict_0_loop *Constructor_Data_Show_Generic_GenericShow) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Show_Generic_GenericShow = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Show_Generic_genericShow_prime__1538939821(a_0_loop gopurs_runtime.Value) string {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Show_Generic_genericShowNoConstructors(), "genericShow'"), a_0).StrVal()
}

func Call_Data_Show_Generic_genericShowArgs__157062380(dict_0_loop *Constructor_Data_Show_Generic_GenericShowArgs) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Show_Generic_GenericShowArgs = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Get_Data_Show_Generic_intercalate() gopurs_runtime.Value {
	return _Gopurs_Data_Show_Generic_Intercalate
}
