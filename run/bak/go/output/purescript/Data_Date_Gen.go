package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Date_Gen_bottom gopurs_runtime.Value
var once_Data_Date_Gen_bottom sync.Once
func Get_Data_Date_Gen_bottom() gopurs_runtime.Value {
	once_Data_Date_Gen_bottom.Do(func() {
		cache_Data_Date_Gen_bottom = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedMonth(), "bottom").IntVal)), UnsafePtr: nil}
	})
	return cache_Data_Date_Gen_bottom
}

var cache_Data_Date_Gen_bottom1 gopurs_runtime.Value
var once_Data_Date_Gen_bottom1 sync.Once
func Get_Data_Date_Gen_bottom1() gopurs_runtime.Value {
	once_Data_Date_Gen_bottom1.Do(func() {
		cache_Data_Date_Gen_bottom1 = gopurs_runtime.Int(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedDay(), "bottom").IntVal)
	})
	return cache_Data_Date_Gen_bottom1
}

var cache_Data_Date_Gen_genDate gopurs_runtime.Value
var once_Data_Date_Gen_genDate sync.Once
func Get_Data_Date_Gen_genDate() gopurs_runtime.Value {
	once_Data_Date_Gen_genDate.Do(func() {
		cache_Data_Date_Gen_genDate = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Date_Gen_genDate(dictMonadGen_0_box)
})
	})
	return cache_Data_Date_Gen_genDate
}

func Call_Data_Date_Gen_genDate(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
// TAST (Let): Monad0_1_0 -> gopurs_runtime.Value
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): Bind1_2_1 -> *Constructor_Control_Bind_Bind
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
// TAST (Let): Functor0_3_2 -> *Constructor_Data_Functor_Functor
Functor0_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_2
// TAST (Let): pure_4_3 -> gopurs_runtime.Value
pure_4_3 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_4_3
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply(Get_Data_Date_Component_Gen_genYear(), dictMonadGen_0), gopurs_runtime.Func(func(year_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 int64
{
if (gopurs_runtime.Apply(Get_Data_Date_isLeapYear(), gopurs_runtime.Int(year_5.IntVal)).IntVal) != (0) {
__t4 = 365
goto end_branch_4
} else {

}
}
{
__t4 = 364
}
end_branch_4:
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_2.V0), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Int_toNumber(), x_6)
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(0), gopurs_runtime.Int(__t4))), gopurs_runtime.Func(func(days_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_4_3, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_6 -> *Constructor_Data_Maybe_Just
__local_var_8_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply3(Get_Data_Date_exactDate(), gopurs_runtime.Int(year_5.IntVal), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedMonth(), "bottom").IntVal)), UnsafePtr: nil}, gopurs_runtime.Int(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedDay(), "bottom").IntVal)))
_ = __local_var_8_6
var __t7 gopurs_runtime.Value
{
if (__local_var_8_6 != nil) {
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(Get_Data_Date_adjust(), gopurs_runtime.Float(days_6.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date]((__local_var_8_6).V0))})))}
goto end_branch_7
} else {

}
}
{
if (__local_var_8_6 == nil) {
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
// TAST (Let): __local_var_8_5 -> *Constructor_Data_Maybe_Just
__local_var_8_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__t7)
_ = __local_var_8_5
var __t8 *Constructor_Data_Date_Date
{
if (__local_var_8_5 != nil) {
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date]((__local_var_8_5).V0)
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(__t8)}
}))))})
}))
}))
}


