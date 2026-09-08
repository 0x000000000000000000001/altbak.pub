package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Date_Component_Gen_toEnum gopurs_runtime.Value
var once_Data_Date_Component_Gen_toEnum sync.Once
func Get_Data_Date_Component_Gen_toEnum() gopurs_runtime.Value {
	once_Data_Date_Component_Gen_toEnum.Do(func() {
		cache_Data_Date_Component_Gen_toEnum = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[int64]](Get_Data_Date_Component_boundedEnumYear()).V4)
	})
	return cache_Data_Date_Component_Gen_toEnum
}

var cache_Data_Date_Component_Gen_genYear gopurs_runtime.Value
var once_Data_Date_Component_Gen_genYear sync.Once
func Get_Data_Date_Component_Gen_genYear() gopurs_runtime.Value {
	once_Data_Date_Component_Gen_genYear.Do(func() {
		cache_Data_Date_Component_Gen_genYear = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Date_Component_Gen_genYear(dictMonadGen_0_box)
})
	})
	return cache_Data_Date_Component_Gen_genYear
}

var cache_Data_Date_Component_Gen_genWeekday gopurs_runtime.Value
var once_Data_Date_Component_Gen_genWeekday sync.Once
func Get_Data_Date_Component_Gen_genWeekday() gopurs_runtime.Value {
	once_Data_Date_Component_Gen_genWeekday.Do(func() {
		cache_Data_Date_Component_Gen_genWeekday = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Date_Component_Gen_genWeekday(dictMonadGen_0_box)
})
	})
	return cache_Data_Date_Component_Gen_genWeekday
}

var cache_Data_Date_Component_Gen_genMonth gopurs_runtime.Value
var once_Data_Date_Component_Gen_genMonth sync.Once
func Get_Data_Date_Component_Gen_genMonth() gopurs_runtime.Value {
	once_Data_Date_Component_Gen_genMonth.Do(func() {
		cache_Data_Date_Component_Gen_genMonth = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Date_Component_Gen_genMonth(dictMonadGen_0_box)
})
	})
	return cache_Data_Date_Component_Gen_genMonth
}

var cache_Data_Date_Component_Gen_genDay gopurs_runtime.Value
var once_Data_Date_Component_Gen_genDay sync.Once
func Get_Data_Date_Component_Gen_genDay() gopurs_runtime.Value {
	once_Data_Date_Component_Gen_genDay.Do(func() {
		cache_Data_Date_Component_Gen_genDay = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Date_Component_Gen_genDay(dictMonadGen_0_box)
})
	})
	return cache_Data_Date_Component_Gen_genDay
}

func Call_Data_Date_Component_Gen_genYear(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 int64
{
if ((x_1.IntVal) >= (int64(-271820))) && ((x_1.IntVal) <= (int64(275759))) {
__t0 = gopurs_runtime.Int(x_1.IntVal).IntVal
goto end_branch_0
} else {

}
}
{
__t0 = func() int64 { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Int(__t0)
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(int64(1900)), gopurs_runtime.Int(int64(2100))))
}

func Call_Data_Date_Component_Gen_genWeekday(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(Get_Data_Enum_Gen_genBoundedEnum(), dictMonadGen_0, gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_Gen_4021906832_123048125(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[uint32]](Get_Data_Date_Component_boundedEnumWeekday())))})
}

func Call_Data_Date_Component_Gen_genMonth(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(Get_Data_Enum_Gen_genBoundedEnum(), dictMonadGen_0, gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_Gen_4021906832_123048125(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[uint32]](Get_Data_Date_Component_boundedEnumMonth())))})
}

func Call_Data_Date_Component_Gen_genDay(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(Get_Data_Enum_Gen_genBoundedEnum(), dictMonadGen_0, gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_Gen_1306125126_123048125(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[int64]](Get_Data_Date_Component_boundedEnumDay())))})
}

func Rebox_Data_Date_Component_Gen_1306125126_123048125(in *Constructor_Data_Enum_BoundedEnum[int64]) *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
		out.V4 = in.V4
	return out
}

func Rebox_Data_Date_Component_Gen_4021906832_123048125(in *Constructor_Data_Enum_BoundedEnum[uint32]) *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
		out.V4 = in.V4
	return out
}


