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
		cache_Data_Date_Gen_bottom = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Call_Data_Bounded_bottom(gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Gen_832288803_2094947566(Rebox_Data_Date_Gen_2094947566_832288803(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Date_Component_boundedMonth()))))}).IntVal)), UnsafePtr: nil}
	})
	return cache_Data_Date_Gen_bottom
}

var cache_Data_Date_Gen_bottom1 gopurs_runtime.Value
var once_Data_Date_Gen_bottom1 sync.Once
func Get_Data_Date_Gen_bottom1() gopurs_runtime.Value {
	once_Data_Date_Gen_bottom1.Do(func() {
		cache_Data_Date_Gen_bottom1 = gopurs_runtime.Int(Call_Data_Bounded_bottom(gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Gen_3764732725_2094947566(Rebox_Data_Date_Gen_2094947566_3764732725(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Date_Component_boundedDay()))))}).IntVal)
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
// TAST (Let): Bind1_2_1 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m$scope1)])
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
// TAST (Let): Functor0_3_2 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m$scope1)])
Functor0_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_2
// TAST (Let): pure_4_3 shape=App(Var) bindingType=(Func [(ADT ["Data","Date","Date"] [])] (TypeApp (TypeVar m$scope1) [(ADT ["Data","Date","Date"] [])]))
pure_4_3 := Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{})))
_ = pure_4_3
return gopurs_runtime.Apply2(Bind1_2_1.V1, Call_Data_Date_Component_Gen_genYear(dictMonadGen_0), gopurs_runtime.Func(func(year_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 int64
{
if ((gopurs_runtime.IntMod(year_5.IntVal, int64(4))) == (int64(0))) && (((gopurs_runtime.IntMod(year_5.IntVal, int64(400))) == (int64(0))) || (((gopurs_runtime.IntMod(year_5.IntVal, int64(100))) == (int64(0))) != (true))) {
__t4 = int64(365)
goto end_branch_4
} else {

}
}
{
__t4 = int64(364)
}
end_branch_4:
return gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply2(Functor0_3_2.V0, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Time_Duration_Days(), Get_Data_Int_toNumber()), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(int64(0)), gopurs_runtime.Int(__t4))), gopurs_runtime.Func(func(days_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_4_3, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](gopurs_runtime.Apply(Get_Partial_Unsafe_unsafePartial(), gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_5 shape=App(Var) bindingType=Any
__local_var_8_5 := Call_Data_Date_exactDate(year_5.IntVal, uint32(Call_Data_Bounded_bottom(gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Gen_832288803_2094947566(Rebox_Data_Date_Gen_2094947566_832288803(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Date_Component_boundedMonth()))))}).IntVal), Call_Data_Bounded_bottom(gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Gen_3764732725_2094947566(Rebox_Data_Date_Gen_2094947566_3764732725(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Date_Component_boundedDay()))))}).IntVal)
_ = __local_var_8_5
var __t7 gopurs_runtime.Value
{
if __local_var_8_5.V1 {
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Gen_2280409795_3094389156(Rebox_Data_Date_Gen_3094389156_2280409795(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := Call_Data_Date_adjust(days_6.FloatVal(), gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date]((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(func() gopurs_runtime.Value {
				_v := __local_var_8_5
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}().UnsafePtr).V0))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
goto end_branch_7
} else {

}
}
{
if (!__local_var_8_5.V1) {
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))}
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }())}
}
end_branch_7:
// TAST (Let): __local_var_9_6 shape=Branch(App(Var), Other, def=Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b$scope103)])
__local_var_9_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t7)
_ = __local_var_9_6
var __t8 *Constructor_Data_Date_Date
{
if (__local_var_9_6 != nil) {
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date]((__local_var_9_6).V0)
goto end_branch_8
} else {

}
}
{
__t8 = func() *Constructor_Data_Date_Date { panic("Failed pattern match") }()
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(__t8)}
}))))})
}))
}))
}

func Rebox_Data_Date_Gen_2094947566_3764732725(in *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]) *Constructor_Data_Bounded_Bounded[int64] {
	if in == nil { return nil }
	out := &Constructor_Data_Bounded_Bounded[int64]{}
		out.V0 = in.V0
		out.V1 = in.V1.IntVal
		out.V2 = in.V2.IntVal
	return out
}

func Rebox_Data_Date_Gen_2094947566_832288803(in *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]) *Constructor_Data_Bounded_Bounded[uint32] {
	if in == nil { return nil }
	out := &Constructor_Data_Bounded_Bounded[uint32]{}
		out.V0 = in.V0
		out.V1 = uint32(in.V1.IntVal)
		out.V2 = uint32(in.V2.IntVal)
	return out
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

func Rebox_Data_Date_Gen_3764732725_2094947566(in *Constructor_Data_Bounded_Bounded[int64]) *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Int(in.V1)
		out.V2 = gopurs_runtime.Int(in.V2)
	return out
}

func Rebox_Data_Date_Gen_832288803_2094947566(in *Constructor_Data_Bounded_Bounded[uint32]) *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V1), UnsafePtr: nil}
		out.V2 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V2), UnsafePtr: nil}
	return out
}


