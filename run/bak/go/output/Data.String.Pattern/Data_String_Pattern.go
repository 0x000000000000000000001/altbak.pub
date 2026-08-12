package Data_String_Pattern

import (
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Show "gopurs/output/Data.Show"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Replacement gopurs_runtime.Value
var once_Replacement sync.Once
func Get_Replacement() gopurs_runtime.Value {
	once_Replacement.Do(func() {
		cache_Replacement = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Replacement(x_0_box)
})
	})
	return cache_Replacement
}

var cache_Pattern gopurs_runtime.Value
var once_Pattern sync.Once
func Get_Pattern() gopurs_runtime.Value {
	once_Pattern.Do(func() {
		cache_Pattern = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Pattern(x_0_box)
})
	})
	return cache_Pattern
}

var cache_showReplacement gopurs_runtime.Value
var once_showReplacement sync.Once
func Get_showReplacement() gopurs_runtime.Value {
	once_showReplacement.Do(func() {
		cache_showReplacement = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str("(Replacement "), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(Call_show__3756561682(gopurs_runtime.Str(v_0.StrVal())).StrVal()), gopurs_runtime.Str(")")).StrVal())).StrVal())
}))
	})
	return cache_showReplacement
}

var cache_showPattern gopurs_runtime.Value
var once_showPattern sync.Once
func Get_showPattern() gopurs_runtime.Value {
	once_showPattern.Do(func() {
		cache_showPattern = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str("(Pattern "), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(Call_show__3756561682(gopurs_runtime.Str(v_0.StrVal())).StrVal()), gopurs_runtime.Str(")")).StrVal())).StrVal())
}))
	})
	return cache_showPattern
}

var cache_newtypeReplacement gopurs_runtime.Value
var once_newtypeReplacement sync.Once
func Get_newtypeReplacement() gopurs_runtime.Value {
	once_newtypeReplacement.Do(func() {
		cache_newtypeReplacement = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeReplacement
}

var cache_newtypePattern gopurs_runtime.Value
var once_newtypePattern sync.Once
func Get_newtypePattern() gopurs_runtime.Value {
	once_newtypePattern.Do(func() {
		cache_newtypePattern = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypePattern
}

var cache_eqReplacement gopurs_runtime.Value
var once_eqReplacement sync.Once
func Get_eqReplacement() gopurs_runtime.Value {
	once_eqReplacement.Do(func() {
		cache_eqReplacement = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((Call_eq__472317769(gopurs_runtime.Str(x_0.StrVal()), gopurs_runtime.Str(y_1.StrVal())).IntVal) != (0))
})
}))
	})
	return cache_eqReplacement
}

var cache_ordReplacement gopurs_runtime.Value
var once_ordReplacement sync.Once
func Get_ordReplacement() gopurs_runtime.Value {
	once_ordReplacement.Do(func() {
		cache_ordReplacement = gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eqReplacement()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Call_compare__882312371(gopurs_runtime.Str(x_0.StrVal()), gopurs_runtime.Str(y_1.StrVal())).IntVal)), UnsafePtr: nil}
})
}))
	})
	return cache_ordReplacement
}

var cache_eqPattern gopurs_runtime.Value
var once_eqPattern sync.Once
func Get_eqPattern() gopurs_runtime.Value {
	once_eqPattern.Do(func() {
		cache_eqPattern = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((Call_eq__472317769(gopurs_runtime.Str(x_0.StrVal()), gopurs_runtime.Str(y_1.StrVal())).IntVal) != (0))
})
}))
	})
	return cache_eqPattern
}

var cache_ordPattern gopurs_runtime.Value
var once_ordPattern sync.Once
func Get_ordPattern() gopurs_runtime.Value {
	once_ordPattern.Do(func() {
		cache_ordPattern = gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eqPattern()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Call_compare__882312371(gopurs_runtime.Str(x_0.StrVal()), gopurs_runtime.Str(y_1.StrVal())).IntVal)), UnsafePtr: nil}
})
}))
	})
	return cache_ordPattern
}

var cache_eq__472317769 gopurs_runtime.Value
var once_eq__472317769 sync.Once
func Get_eq__472317769() gopurs_runtime.Value {
	once_eq__472317769.Do(func() {
		cache_eq__472317769 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq__472317769(__eta0_0_box, __eta1_1_box)
})
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
		cache_compare__882312371 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compare__882312371(__eta0_0_box, __eta1_1_box)
})
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
		cache_append__493084344 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__493084344(__eta0_0_box, __eta1_1_box)
})
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
		cache_show__3756561682 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_show__3756561682(__eta0_0_box)
})
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

func Call_Replacement(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Pattern(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_eq__472317769(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Bool((__eta0_0.StrVal()) == (__eta1_1.StrVal()))
}

func Call_eq__2384498378(dict_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_compare__882312371(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply5(pkg_Data_Ord.Get_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __eta0_0, __eta1_1)
}

func Call_compare__821463600(dict_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_append__493084344(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Str((__eta0_0.StrVal()) + (__eta1_1.StrVal()))
}

func Call_append__1230318264(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_show__3756561682(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Str(gopurs_runtime.Apply(pkg_Data_Show.Get_showStringImpl(), __eta0_0).StrVal())
}

func Call_show__2742601362(dict_0_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}


