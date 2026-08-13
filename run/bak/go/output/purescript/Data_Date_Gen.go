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
		cache_Data_Date_Gen_bottom = gopurs_runtime.Value{Type: 9, IntVal: int64(1908470532), UnsafePtr: nil}
	})
	return cache_Data_Date_Gen_bottom
}

var cache_Data_Date_Gen_bottom1 gopurs_runtime.Value
var once_Data_Date_Gen_bottom1 sync.Once
func Get_Data_Date_Gen_bottom1() gopurs_runtime.Value {
	once_Data_Date_Gen_bottom1.Do(func() {
		cache_Data_Date_Gen_bottom1 = gopurs_runtime.Int(1)
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
// TAST (Let): __local_var_5_4 -> gopurs_runtime.Value
__local_var_5_4 := gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Maybe_fromJust__1577979644()
}))
_ = __local_var_5_4
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 *Constructor_Data_Maybe_Just
{
var __t5 bool
{
if (x_6.IntVal) < (-271820) {
__t5 = false
goto end_branch_5
} else {

}
}
{
__t5 = true
}
end_branch_5:
var __t_and_7 bool = false
if __t5 {

var __t6 bool
{
if (x_6.IntVal) > (275759) {
__t6 = false
goto end_branch_6
} else {

}
}
{
__t6 = true
}
end_branch_6:
__t_and_7 = __t6
}
if __t_and_7 {
__t8 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(x_6.IntVal)}
goto end_branch_8
} else {

}
}
{
__t8 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_8:
return gopurs_runtime.Apply(__local_var_5_4, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t8)})
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(1900), gopurs_runtime.Int(2100))), gopurs_runtime.Func(func(year_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 int64
{
if (gopurs_runtime.Apply(Get_Data_Date_isLeapYear(), gopurs_runtime.Int(year_5.IntVal)).IntVal) != (0) {
__t9 = 365
goto end_branch_9
} else {

}
}
{
__t9 = 364
}
end_branch_9:
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_2.V0), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Int_toNumber(), x_6)
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(0), gopurs_runtime.Int(__t9))), gopurs_runtime.Func(func(days_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_4_3, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_11 -> *Constructor_Data_Maybe_Just
__local_var_8_11 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply3(Get_Data_Date_exactDate(), gopurs_runtime.Int(year_5.IntVal), gopurs_runtime.Value{Type: 9, IntVal: int64(1908470532), UnsafePtr: nil}, gopurs_runtime.Int(1)))
_ = __local_var_8_11
var __t12 *Constructor_Data_Maybe_Just
{
if (__local_var_8_11 != nil) {
__t12 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(Get_Data_Date_adjust(), gopurs_runtime.Float(days_6.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date]((__local_var_8_11).V0))}))
goto end_branch_12
} else {

}
}
{
if (__local_var_8_11 == nil) {
__t12 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_12
} else {

}
}
{
__t12 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_12:
// TAST (Let): __local_var_8_10 -> *Constructor_Data_Maybe_Just
__local_var_8_10 := __t12
_ = __local_var_8_10
var __t13 *Constructor_Data_Date_Date
{
if (__local_var_8_10 != nil) {
__t13 = gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date]((__local_var_8_10).V0)
goto end_branch_13
} else {

}
}
{
__t13 = gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_13:
return gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(__t13)}
}))))})
}))
}))
}


