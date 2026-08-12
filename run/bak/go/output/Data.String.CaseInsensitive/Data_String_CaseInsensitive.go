package Data_String_CaseInsensitive

import (
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_String_Common "gopurs/output/Data.String.Common"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_CaseInsensitiveString gopurs_runtime.Value
var once_CaseInsensitiveString sync.Once
func Get_CaseInsensitiveString() gopurs_runtime.Value {
	once_CaseInsensitiveString.Do(func() {
		cache_CaseInsensitiveString = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_CaseInsensitiveString(x_0_box)
})
	})
	return cache_CaseInsensitiveString
}

var cache_showCaseInsensitiveString gopurs_runtime.Value
var once_showCaseInsensitiveString sync.Once
func Get_showCaseInsensitiveString() gopurs_runtime.Value {
	once_showCaseInsensitiveString.Do(func() {
		cache_showCaseInsensitiveString = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Str("(CaseInsensitiveString "), gopurs_runtime.Str(gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_show__3756561682(), gopurs_runtime.Str(v_0.StrVal())).StrVal()), gopurs_runtime.Str(")")).StrVal())).StrVal())
}))
	})
	return cache_showCaseInsensitiveString
}

var cache_newtypeCaseInsensitiveString gopurs_runtime.Value
var once_newtypeCaseInsensitiveString sync.Once
func Get_newtypeCaseInsensitiveString() gopurs_runtime.Value {
	once_newtypeCaseInsensitiveString.Do(func() {
		cache_newtypeCaseInsensitiveString = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeCaseInsensitiveString
}

var cache_eqCaseInsensitiveString gopurs_runtime.Value
var once_eqCaseInsensitiveString sync.Once
func Get_eqCaseInsensitiveString() gopurs_runtime.Value {
	once_eqCaseInsensitiveString.Do(func() {
		cache_eqCaseInsensitiveString = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(Get_eq__472317769(), gopurs_runtime.Str(gopurs_runtime.Apply(pkg_Data_String_Common.Get_toLower(), gopurs_runtime.Str(v_0.StrVal())).StrVal()), gopurs_runtime.Str(gopurs_runtime.Apply(pkg_Data_String_Common.Get_toLower(), gopurs_runtime.Str(v1_1.StrVal())).StrVal())).IntVal) != (0))
})
}))
	})
	return cache_eqCaseInsensitiveString
}

var cache_ordCaseInsensitiveString gopurs_runtime.Value
var once_ordCaseInsensitiveString sync.Once
func Get_ordCaseInsensitiveString() gopurs_runtime.Value {
	once_ordCaseInsensitiveString.Do(func() {
		cache_ordCaseInsensitiveString = gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eqCaseInsensitiveString()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(Get_compare__882312371(), gopurs_runtime.Str(gopurs_runtime.Apply(pkg_Data_String_Common.Get_toLower(), gopurs_runtime.Str(v_0.StrVal())).StrVal()), gopurs_runtime.Str(gopurs_runtime.Apply(pkg_Data_String_Common.Get_toLower(), gopurs_runtime.Str(v1_1.StrVal())).StrVal())).IntVal)), UnsafePtr: nil}
})
}))
	})
	return cache_ordCaseInsensitiveString
}

var cache_eq__472317769 gopurs_runtime.Value
var once_eq__472317769 sync.Once
func Get_eq__472317769() gopurs_runtime.Value {
	once_eq__472317769.Do(func() {
		cache_eq__472317769 = gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqString(), "eq")
	})
	return cache_eq__472317769
}

var cache_eq__2384498378 gopurs_runtime.Value
var once_eq__2384498378 sync.Once
func Get_eq__2384498378() gopurs_runtime.Value {
	once_eq__2384498378.Do(func() {
		cache_eq__2384498378 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq__2384498378(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_eq__2384498378
}

var cache_compare__882312371 gopurs_runtime.Value
var once_compare__882312371 sync.Once
func Get_compare__882312371() gopurs_runtime.Value {
	once_compare__882312371.Do(func() {
		cache_compare__882312371 = gopurs_runtime.RecordGet(pkg_Data_Ord.Get_ordString(), "compare")
	})
	return cache_compare__882312371
}

var cache_compare__821463600 gopurs_runtime.Value
var once_compare__821463600 sync.Once
func Get_compare__821463600() gopurs_runtime.Value {
	once_compare__821463600.Do(func() {
		cache_compare__821463600 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compare__821463600(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compare__821463600
}

var cache_append__493084344 gopurs_runtime.Value
var once_append__493084344 sync.Once
func Get_append__493084344() gopurs_runtime.Value {
	once_append__493084344.Do(func() {
		cache_append__493084344 = gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append")
	})
	return cache_append__493084344
}

var cache_append__1230318264 gopurs_runtime.Value
var once_append__1230318264 sync.Once
func Get_append__1230318264() gopurs_runtime.Value {
	once_append__1230318264.Do(func() {
		cache_append__1230318264 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__1230318264(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_append__1230318264
}

var cache_show__3756561682 gopurs_runtime.Value
var once_show__3756561682 sync.Once
func Get_show__3756561682() gopurs_runtime.Value {
	once_show__3756561682.Do(func() {
		cache_show__3756561682 = gopurs_runtime.RecordGet(pkg_Data_Show.Get_showString(), "show")
	})
	return cache_show__3756561682
}

var cache_show__2742601362 gopurs_runtime.Value
var once_show__2742601362 sync.Once
func Get_show__2742601362() gopurs_runtime.Value {
	once_show__2742601362.Do(func() {
		cache_show__2742601362 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_show__2742601362(gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_show__2742601362
}

func Call_CaseInsensitiveString(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_eq__2384498378(dict_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_compare__821463600(dict_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_append__1230318264(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_show__2742601362(dict_0_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}


