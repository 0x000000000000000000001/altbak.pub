package Data_Time_Component

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	unsafe "unsafe"
)

var cache_showSecond gopurs_runtime.Value
var once_showSecond sync.Once
func Get_showSecond() gopurs_runtime.Value {
	once_showSecond.Do(func() {
		cache_showSecond = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Second ") + (gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), v_0).StrVal())) + (")"))
}))
	})
	return cache_showSecond
}

var cache_showMinute gopurs_runtime.Value
var once_showMinute sync.Once
func Get_showMinute() gopurs_runtime.Value {
	once_showMinute.Do(func() {
		cache_showMinute = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Minute ") + (gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), v_0).StrVal())) + (")"))
}))
	})
	return cache_showMinute
}

var cache_showMillisecond gopurs_runtime.Value
var once_showMillisecond sync.Once
func Get_showMillisecond() gopurs_runtime.Value {
	once_showMillisecond.Do(func() {
		cache_showMillisecond = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Millisecond ") + (gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), v_0).StrVal())) + (")"))
}))
	})
	return cache_showMillisecond
}

var cache_showHour gopurs_runtime.Value
var once_showHour sync.Once
func Get_showHour() gopurs_runtime.Value {
	once_showHour.Do(func() {
		cache_showHour = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Hour ") + (gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), v_0).StrVal())) + (")"))
}))
	})
	return cache_showHour
}

var cache_ordSecond gopurs_runtime.Value
var once_ordSecond sync.Once
func Get_ordSecond() gopurs_runtime.Value {
	once_ordSecond.Do(func() {
		cache_ordSecond = pkg_Data_Ord.Get_ordInt()
	})
	return cache_ordSecond
}

var cache_ordMinute gopurs_runtime.Value
var once_ordMinute sync.Once
func Get_ordMinute() gopurs_runtime.Value {
	once_ordMinute.Do(func() {
		cache_ordMinute = pkg_Data_Ord.Get_ordInt()
	})
	return cache_ordMinute
}

var cache_ordMillisecond gopurs_runtime.Value
var once_ordMillisecond sync.Once
func Get_ordMillisecond() gopurs_runtime.Value {
	once_ordMillisecond.Do(func() {
		cache_ordMillisecond = pkg_Data_Ord.Get_ordInt()
	})
	return cache_ordMillisecond
}

var cache_ordHour gopurs_runtime.Value
var once_ordHour sync.Once
func Get_ordHour() gopurs_runtime.Value {
	once_ordHour.Do(func() {
		cache_ordHour = pkg_Data_Ord.Get_ordInt()
	})
	return cache_ordHour
}

var cache_eqSecond gopurs_runtime.Value
var once_eqSecond sync.Once
func Get_eqSecond() gopurs_runtime.Value {
	once_eqSecond.Do(func() {
		cache_eqSecond = pkg_Data_Eq.Get_eqInt()
	})
	return cache_eqSecond
}

var cache_eqMinute gopurs_runtime.Value
var once_eqMinute sync.Once
func Get_eqMinute() gopurs_runtime.Value {
	once_eqMinute.Do(func() {
		cache_eqMinute = pkg_Data_Eq.Get_eqInt()
	})
	return cache_eqMinute
}

var cache_eqMillisecond gopurs_runtime.Value
var once_eqMillisecond sync.Once
func Get_eqMillisecond() gopurs_runtime.Value {
	once_eqMillisecond.Do(func() {
		cache_eqMillisecond = pkg_Data_Eq.Get_eqInt()
	})
	return cache_eqMillisecond
}

var cache_eqHour gopurs_runtime.Value
var once_eqHour sync.Once
func Get_eqHour() gopurs_runtime.Value {
	once_eqHour.Do(func() {
		cache_eqHour = pkg_Data_Eq.Get_eqInt()
	})
	return cache_eqHour
}

var cache_boundedSecond gopurs_runtime.Value
var once_boundedSecond sync.Once
func Get_boundedSecond() gopurs_runtime.Value {
	once_boundedSecond.Do(func() {
		cache_boundedSecond = gopurs_runtime.RecordDict3("bottom", "top", "Ord0", gopurs_runtime.Int(0), gopurs_runtime.Int(59), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordInt()
}))
	})
	return cache_boundedSecond
}

var cache_boundedMinute gopurs_runtime.Value
var once_boundedMinute sync.Once
func Get_boundedMinute() gopurs_runtime.Value {
	once_boundedMinute.Do(func() {
		cache_boundedMinute = gopurs_runtime.RecordDict3("bottom", "top", "Ord0", gopurs_runtime.Int(0), gopurs_runtime.Int(59), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordInt()
}))
	})
	return cache_boundedMinute
}

var cache_boundedMillisecond gopurs_runtime.Value
var once_boundedMillisecond sync.Once
func Get_boundedMillisecond() gopurs_runtime.Value {
	once_boundedMillisecond.Do(func() {
		cache_boundedMillisecond = gopurs_runtime.RecordDict3("bottom", "top", "Ord0", gopurs_runtime.Int(0), gopurs_runtime.Int(999), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordInt()
}))
	})
	return cache_boundedMillisecond
}

var cache_boundedHour gopurs_runtime.Value
var once_boundedHour sync.Once
func Get_boundedHour() gopurs_runtime.Value {
	once_boundedHour.Do(func() {
		cache_boundedHour = gopurs_runtime.RecordDict3("bottom", "top", "Ord0", gopurs_runtime.Int(0), gopurs_runtime.Int(23), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordInt()
}))
	})
	return cache_boundedHour
}

var cache_boundedEnumSecond gopurs_runtime.Value
var once_boundedEnumSecond sync.Once
func Get_boundedEnumSecond() gopurs_runtime.Value {
	once_boundedEnumSecond.Do(func() {
		cache_boundedEnumSecond = gopurs_runtime.RecordDict5("cardinality", "toEnum", "fromEnum", "Bounded0", "Enum1", gopurs_runtime.Int(60), gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if ((n_0.IntVal) >= (0)) && ((n_0.IntVal) <= (59)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{n_0})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_0:
return __t0
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_boundedSecond()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_enumSecond()
}))
	})
	return cache_boundedEnumSecond
}

var cache_enumSecond gopurs_runtime.Value
var once_enumSecond sync.Once
func Get_enumSecond() gopurs_runtime.Value {
	once_enumSecond.Do(func() {
		cache_enumSecond = gopurs_runtime.RecordDict3("succ", "pred", "Ord0", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := (x_0.IntVal) + (1)
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if ((__local_var_1_0) >= (0)) && ((__local_var_1_0) <= (59)) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Int(__local_var_1_0)})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_1:
return __t1
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_2 := (x_0.IntVal) - (1)
_ = __local_var_1_2
var __t3 gopurs_runtime.Value
{
if ((__local_var_1_2) >= (0)) && ((__local_var_1_2) <= (59)) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Int(__local_var_1_2)})}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_3:
return __t3
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordInt()
}))
	})
	return cache_enumSecond
}

var cache_boundedEnumMinute gopurs_runtime.Value
var once_boundedEnumMinute sync.Once
func Get_boundedEnumMinute() gopurs_runtime.Value {
	once_boundedEnumMinute.Do(func() {
		cache_boundedEnumMinute = gopurs_runtime.RecordDict5("cardinality", "toEnum", "fromEnum", "Bounded0", "Enum1", gopurs_runtime.Int(60), gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if ((n_0.IntVal) >= (0)) && ((n_0.IntVal) <= (59)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{n_0})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_0:
return __t0
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_boundedMinute()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_enumMinute()
}))
	})
	return cache_boundedEnumMinute
}

var cache_enumMinute gopurs_runtime.Value
var once_enumMinute sync.Once
func Get_enumMinute() gopurs_runtime.Value {
	once_enumMinute.Do(func() {
		cache_enumMinute = gopurs_runtime.RecordDict3("succ", "pred", "Ord0", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := (x_0.IntVal) + (1)
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if ((__local_var_1_0) >= (0)) && ((__local_var_1_0) <= (59)) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Int(__local_var_1_0)})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_1:
return __t1
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_2 := (x_0.IntVal) - (1)
_ = __local_var_1_2
var __t3 gopurs_runtime.Value
{
if ((__local_var_1_2) >= (0)) && ((__local_var_1_2) <= (59)) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Int(__local_var_1_2)})}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_3:
return __t3
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordInt()
}))
	})
	return cache_enumMinute
}

var cache_boundedEnumMillisecond gopurs_runtime.Value
var once_boundedEnumMillisecond sync.Once
func Get_boundedEnumMillisecond() gopurs_runtime.Value {
	once_boundedEnumMillisecond.Do(func() {
		cache_boundedEnumMillisecond = gopurs_runtime.RecordDict5("cardinality", "toEnum", "fromEnum", "Bounded0", "Enum1", gopurs_runtime.Int(1000), gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if ((n_0.IntVal) >= (0)) && ((n_0.IntVal) <= (999)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{n_0})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_0:
return __t0
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_boundedMillisecond()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_enumMillisecond()
}))
	})
	return cache_boundedEnumMillisecond
}

var cache_enumMillisecond gopurs_runtime.Value
var once_enumMillisecond sync.Once
func Get_enumMillisecond() gopurs_runtime.Value {
	once_enumMillisecond.Do(func() {
		cache_enumMillisecond = gopurs_runtime.RecordDict3("succ", "pred", "Ord0", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := (x_0.IntVal) + (1)
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if ((__local_var_1_0) >= (0)) && ((__local_var_1_0) <= (999)) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Int(__local_var_1_0)})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_1:
return __t1
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_2 := (x_0.IntVal) - (1)
_ = __local_var_1_2
var __t3 gopurs_runtime.Value
{
if ((__local_var_1_2) >= (0)) && ((__local_var_1_2) <= (999)) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Int(__local_var_1_2)})}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_3:
return __t3
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordInt()
}))
	})
	return cache_enumMillisecond
}

var cache_boundedEnumHour gopurs_runtime.Value
var once_boundedEnumHour sync.Once
func Get_boundedEnumHour() gopurs_runtime.Value {
	once_boundedEnumHour.Do(func() {
		cache_boundedEnumHour = gopurs_runtime.RecordDict5("cardinality", "toEnum", "fromEnum", "Bounded0", "Enum1", gopurs_runtime.Int(24), gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if ((n_0.IntVal) >= (0)) && ((n_0.IntVal) <= (23)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{n_0})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_0:
return __t0
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_boundedHour()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_enumHour()
}))
	})
	return cache_boundedEnumHour
}

var cache_enumHour gopurs_runtime.Value
var once_enumHour sync.Once
func Get_enumHour() gopurs_runtime.Value {
	once_enumHour.Do(func() {
		cache_enumHour = gopurs_runtime.RecordDict3("succ", "pred", "Ord0", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := (x_0.IntVal) + (1)
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if ((__local_var_1_0) >= (0)) && ((__local_var_1_0) <= (23)) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Int(__local_var_1_0)})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_1:
return __t1
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_2 := (x_0.IntVal) - (1)
_ = __local_var_1_2
var __t3 gopurs_runtime.Value
{
if ((__local_var_1_2) >= (0)) && ((__local_var_1_2) <= (23)) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Int(__local_var_1_2)})}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_3:
return __t3
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordInt()
}))
	})
	return cache_enumHour
}




