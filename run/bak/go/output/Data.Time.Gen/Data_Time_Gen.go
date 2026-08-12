package Data_Time_Gen

import (
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Data_Enum_Gen "gopurs/output/Data.Enum.Gen"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Time "gopurs/output/Data.Time"
	pkg_Data_Time_Component "gopurs/output/Data.Time.Component"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_genTime gopurs_runtime.Value
var once_genTime sync.Once
func Get_genTime() gopurs_runtime.Value {
	once_genTime.Do(func() {
		cache_genTime = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genTime(dictMonadGen_0_box)
})
	})
	return cache_genTime
}

var cache_apply__4140882028 gopurs_runtime.Value
var once_apply__4140882028 sync.Once
func Get_apply__4140882028() gopurs_runtime.Value {
	once_apply__4140882028.Do(func() {
		cache_apply__4140882028 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__4140882028(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_apply__4140882028
}

var cache_apply__2008680204 gopurs_runtime.Value
var once_apply__2008680204 sync.Once
func Get_apply__2008680204() gopurs_runtime.Value {
	once_apply__2008680204.Do(func() {
		cache_apply__2008680204 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__2008680204(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_apply__2008680204
}

var cache_apply__2092855244 gopurs_runtime.Value
var once_apply__2092855244 sync.Once
func Get_apply__2092855244() gopurs_runtime.Value {
	once_apply__2092855244.Do(func() {
		cache_apply__2092855244 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__2092855244(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_apply__2092855244
}

var cache_apply__353515660 gopurs_runtime.Value
var once_apply__353515660 sync.Once
func Get_apply__353515660() gopurs_runtime.Value {
	once_apply__353515660.Do(func() {
		cache_apply__353515660 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__353515660(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_apply__353515660
}

var cache_map__358342900 gopurs_runtime.Value
var once_map__358342900 sync.Once
func Get_map__358342900() gopurs_runtime.Value {
	once_map__358342900.Do(func() {
		cache_map__358342900 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__358342900(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__358342900
}

var cache_map__2199395572 gopurs_runtime.Value
var once_map__2199395572 sync.Once
func Get_map__2199395572() gopurs_runtime.Value {
	once_map__2199395572.Do(func() {
		cache_map__2199395572 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2199395572(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2199395572
}

func Call_genTime(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
Bind1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{})
_ = Bind1_1_0
Apply0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_1_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_2_1
return gopurs_runtime.Apply2(Apply0_2_1.V1, gopurs_runtime.Apply2(Apply0_2_1.V1, gopurs_runtime.Apply2(Apply0_2_1.V1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_1_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Time.Get_Time(), gopurs_runtime.Apply2(pkg_Data_Enum_Gen.Get_genBoundedEnum(), dictMonadGen_0, pkg_Data_Time_Component.Get_boundedEnumHour())), gopurs_runtime.Apply2(pkg_Data_Enum_Gen.Get_genBoundedEnum(), dictMonadGen_0, pkg_Data_Time_Component.Get_boundedEnumMinute())), gopurs_runtime.Apply2(pkg_Data_Enum_Gen.Get_genBoundedEnum(), dictMonadGen_0, pkg_Data_Time_Component.Get_boundedEnumSecond())), gopurs_runtime.Apply2(pkg_Data_Enum_Gen.Get_genBoundedEnum(), dictMonadGen_0, pkg_Data_Time_Component.Get_boundedEnumMillisecond()))
}

func Call_apply__4140882028(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__2008680204(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__2092855244(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__353515660(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_map__358342900(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__2199395572(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}


