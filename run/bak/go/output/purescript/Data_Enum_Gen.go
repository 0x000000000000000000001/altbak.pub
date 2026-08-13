package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Enum_Gen_foldable1NonEmpty gopurs_runtime.Value
var once_Data_Enum_Gen_foldable1NonEmpty sync.Once
func Get_Data_Enum_Gen_foldable1NonEmpty() gopurs_runtime.Value {
	once_Data_Enum_Gen_foldable1NonEmpty.Do(func() {
		cache_Data_Enum_Gen_foldable1NonEmpty = gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.Apply(Get_Data_NonEmpty_foldable1NonEmpty(), Get_Data_Foldable_foldableArray())))}
	})
	return cache_Data_Enum_Gen_foldable1NonEmpty
}

var cache_Data_Enum_Gen_genBoundedEnum gopurs_runtime.Value
var once_Data_Enum_Gen_genBoundedEnum sync.Once
func Get_Data_Enum_Gen_genBoundedEnum() gopurs_runtime.Value {
	once_Data_Enum_Gen_genBoundedEnum.Do(func() {
		cache_Data_Enum_Gen_genBoundedEnum = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_Gen_genBoundedEnum(dictMonadGen_0_box)
})
	})
	return cache_Data_Enum_Gen_genBoundedEnum
}

func Call_Data_Enum_Gen_genBoundedEnum(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
// TAST (Let): Applicative0_1_0 -> *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_0
return gopurs_runtime.Func(func(dictBoundedEnum_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Enum1_3_1 -> *Constructor_Data_Enum_Enum[gopurs_runtime.Value]
Enum1_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_2, "Enum1"), gopurs_runtime.Value{}))
_ = Enum1_3_1
// TAST (Let): Bounded0_4_2 -> *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]
Bounded0_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_2, "Bounded0"), gopurs_runtime.Value{}))
_ = Bounded0_4_2
// TAST (Let): v_5_3 -> gopurs_runtime.Value
v_5_3 := gopurs_runtime.Apply(gopurs_runtime.Box(Enum1_3_1.V2), gopurs_runtime.Box(Bounded0_4_2.V1))
_ = v_5_3
var __t4 gopurs_runtime.Value
{
if (v_5_3.Type == 9 && v_5_3.IntVal == 930809136 && v_5_3.UnsafePtr != nil) {
__t4 = gopurs_runtime.Apply3(Get_Control_Monad_Gen_elements(), gopurs_runtime.Value{Type: 9, IntVal: 2254593219, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_0))}, gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_Enum_Gen_foldable1NonEmpty()))}, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Box(Bounded0_4_2.V1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply4(Get_Data_Enum_enumFromTo(), gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Enum1_3_1)}, gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]](Get_Data_Unfoldable1_unfoldable1Array()))}, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_5_3.UnsafePtr).V0, gopurs_runtime.Box(Bounded0_4_2.V2)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())})})
goto end_branch_4
} else {

}
}
{
if (v_5_3.Type == 9 && v_5_3.IntVal == 930809136 && v_5_3.UnsafePtr == nil) {
__t4 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), gopurs_runtime.Box(Bounded0_4_2.V1))
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
})
}


