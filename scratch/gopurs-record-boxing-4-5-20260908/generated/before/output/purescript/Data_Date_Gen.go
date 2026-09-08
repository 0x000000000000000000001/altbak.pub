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
		cache_Data_Date_Gen_bottom1 = gopurs_runtime.Int(int64(1))
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
// TAST (Let): Monad0_1_0 shape=App(Other) bindingType=Any
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): Bind1_2_1 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
// TAST (Let): Functor0_3_2 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_2
// TAST (Let): pure_4_3 shape=Other bindingType=(Func [(ADT ["Data","Date","Date"] [])] (TypeApp (TypeVar m) [(ADT ["Data","Date","Date"] [])]))
pure_4_3 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_4_3
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 int64
{
if ((x_5.IntVal) >= (int64(-271820))) && ((x_5.IntVal) <= (int64(275759))) {
__t4 = gopurs_runtime.Int(x_5.IntVal).IntVal
goto end_branch_4
} else {

}
}
{
__t4 = func() int64 { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Int(__t4)
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(int64(1900)), gopurs_runtime.Int(int64(2100)))), gopurs_runtime.Func(func(year_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 int64
{
if (((year_5.IntVal) % (int64(4))) == (int64(0))) && ((((year_5.IntVal) % (int64(400))) == (int64(0))) || ((((year_5.IntVal) % (int64(100))) == (int64(0))) != (true))) {
__t5 = int64(365)
goto end_branch_5
} else {

}
}
{
__t5 = int64(364)
}
end_branch_5:
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_2.V0), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(gopurs_runtime.Apply(Get_Data_Int_toNumber(), x_6).FloatVal())
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(int64(0)), gopurs_runtime.Int(__t5))), gopurs_runtime.Func(func(days_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_6 shape=App(Var) bindingType=Any
__local_var_7_6 := gopurs_runtime.Apply3(Get_Data_Date_exactDate(), gopurs_runtime.Int(year_5.IntVal), gopurs_runtime.Value{Type: 9, IntVal: int64(1908470532), UnsafePtr: nil}, gopurs_runtime.Int(int64(1)))
_ = __local_var_7_6
var __t8 gopurs_runtime.Value
{
if (__local_var_7_6.Type == 9 && __local_var_7_6.IntVal == 930809136 && __local_var_7_6.UnsafePtr != nil) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Gen_2280409795_3094389156(Rebox_Data_Date_Gen_3094389156_2280409795(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(Get_Data_Date_adjust(), gopurs_runtime.Float(days_6.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date]((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_7_6.UnsafePtr).V0))})))))}
goto end_branch_8
} else {

}
}
{
if (__local_var_7_6.Type == 9 && __local_var_7_6.IntVal == 930809136 && __local_var_7_6.UnsafePtr == nil) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }())}
}
end_branch_8:
// TAST (Let): __local_var_8_7 shape=Branch(App(Var), Other, def=Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b)])
__local_var_8_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t8)
_ = __local_var_8_7
var __t9 *Constructor_Data_Date_Date
{
if (__local_var_8_7 != nil) {
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date]((__local_var_8_7).V0)
goto end_branch_9
} else {

}
}
{
__t9 = func() *Constructor_Data_Date_Date { panic("Failed pattern match") }()
}
end_branch_9:
return gopurs_runtime.Apply(pure_4_3, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(__t9)})
}))
}))
}

func Rebox_Data_Date_Gen_2280409795_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}

func Rebox_Data_Date_Gen_3094389156_2280409795(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]{}
		out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](in.V0)
	return out
}


