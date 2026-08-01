package Data_Decide

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Divide "gopurs/output/Data.Divide"
	pkg_Data_Either "gopurs/output/Data.Either"
)

var cache_identity gopurs_runtime.Value
var once_identity sync.Once
func Get_identity() gopurs_runtime.Value {
	once_identity.Do(func() {
		cache_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_identity(x_0_box)
})
	})
	return cache_identity
}

var cache_choosePredicate gopurs_runtime.Value
var once_choosePredicate sync.Once
func Get_choosePredicate() gopurs_runtime.Value {
	once_choosePredicate.Do(func() {
		cache_choosePredicate = gopurs_runtime.RecordDict2("Divide0", "choose", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Divide.Get_dividePredicate()
}), gopurs_runtime.Func4(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value, x_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_0 := gopurs_runtime.Apply(f_0, x_3)
_ = __local_var_4_0
var __t1 gopurs_runtime.Value
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 3711209382) {
__t1 = gopurs_runtime.Apply(v_1, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_0.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 2465973597) {
__t1 = gopurs_runtime.Apply(v1_2, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_0.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}))
	})
	return cache_choosePredicate
}

var cache_chooseOp gopurs_runtime.Value
var once_chooseOp sync.Once
func Get_chooseOp() gopurs_runtime.Value {
	once_chooseOp.Do(func() {
		cache_chooseOp = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_chooseOp(dictSemigroup_0_box)
})
	})
	return cache_chooseOp
}

var cache_chooseEquivalence gopurs_runtime.Value
var once_chooseEquivalence sync.Once
func Get_chooseEquivalence() gopurs_runtime.Value {
	once_chooseEquivalence.Do(func() {
		cache_chooseEquivalence = gopurs_runtime.RecordDict2("Divide0", "choose", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Divide.Get_divideEquivalence()
}), gopurs_runtime.Func5(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value) gopurs_runtime.Value {
v2_5_0 := gopurs_runtime.Apply(f_0, a_3)
_ = v2_5_0
var __t1 gopurs_runtime.Value
{
if (v2_5_0.Type == 9 && v2_5_0.IntVal == 3711209382) {
v3_6_2 := gopurs_runtime.Apply(f_0, b_4)
_ = v3_6_2
var __t3 gopurs_runtime.Value
{
if (v3_6_2.Type == 9 && v3_6_2.IntVal == 3711209382) {
__t3 = gopurs_runtime.Apply2(v_1, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5_0.UnsafePtr).V0, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v3_6_2.UnsafePtr).V0)
goto end_branch_3
} else {

}
}
{
if (v3_6_2.Type == 9 && v3_6_2.IntVal == 2465973597) {
__t3 = gopurs_runtime.Bool(false)
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t1 = __t3
goto end_branch_1
} else {

}
}
{
if (v2_5_0.Type == 9 && v2_5_0.IntVal == 2465973597) {
v3_6_4 := gopurs_runtime.Apply(f_0, b_4)
_ = v3_6_4
var __t5 gopurs_runtime.Value
{
if (v3_6_4.Type == 9 && v3_6_4.IntVal == 3711209382) {
__t5 = gopurs_runtime.Bool(false)
goto end_branch_5
} else {

}
}
{
if (v3_6_4.Type == 9 && v3_6_4.IntVal == 2465973597) {
__t5 = gopurs_runtime.Apply2(v1_2, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5_0.UnsafePtr).V0, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v3_6_4.UnsafePtr).V0)
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
__t1 = __t5
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}))
	})
	return cache_chooseEquivalence
}

var cache_chooseComparison gopurs_runtime.Value
var once_chooseComparison sync.Once
func Get_chooseComparison() gopurs_runtime.Value {
	once_chooseComparison.Do(func() {
		cache_chooseComparison = gopurs_runtime.RecordDict2("Divide0", "choose", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Divide.Get_divideComparison()
}), gopurs_runtime.Func5(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value) gopurs_runtime.Value {
v2_5_0 := gopurs_runtime.Apply(f_0, a_3)
_ = v2_5_0
var __t1 gopurs_runtime.Value
{
if (v2_5_0.Type == 9 && v2_5_0.IntVal == 3711209382) {
v3_6_2 := gopurs_runtime.Apply(f_0, b_4)
_ = v3_6_2
var __t3 gopurs_runtime.Value
{
if (v3_6_2.Type == 9 && v3_6_2.IntVal == 3711209382) {
__t3 = gopurs_runtime.Apply2(v_1, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5_0.UnsafePtr).V0, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v3_6_2.UnsafePtr).V0)
goto end_branch_3
} else {

}
}
{
if (v3_6_2.Type == 9 && v3_6_2.IntVal == 2465973597) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t1 = __t3
goto end_branch_1
} else {

}
}
{
if (v2_5_0.Type == 9 && v2_5_0.IntVal == 2465973597) {
v3_6_4 := gopurs_runtime.Apply(f_0, b_4)
_ = v3_6_4
var __t5 gopurs_runtime.Value
{
if (v3_6_4.Type == 9 && v3_6_4.IntVal == 3711209382) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}
goto end_branch_5
} else {

}
}
{
if (v3_6_4.Type == 9 && v3_6_4.IntVal == 2465973597) {
__t5 = gopurs_runtime.Apply2(v1_2, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5_0.UnsafePtr).V0, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v3_6_4.UnsafePtr).V0)
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
__t1 = __t5
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}))
	})
	return cache_chooseComparison
}

var cache_choose gopurs_runtime.Value
var once_choose sync.Once
func Get_choose() gopurs_runtime.Value {
	once_choose.Do(func() {
		cache_choose = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_choose(dict_0_box)
})
	})
	return cache_choose
}

var cache_choose__gopurs_runtime_Value_1032977354 gopurs_runtime.Value
var once_choose__gopurs_runtime_Value_1032977354 sync.Once
func Get_choose__gopurs_runtime_Value_1032977354() gopurs_runtime.Value {
	once_choose__gopurs_runtime_Value_1032977354.Do(func() {
		cache_choose__gopurs_runtime_Value_1032977354 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_choose__gopurs_runtime_Value_1032977354(dict_0_box)
})
	})
	return cache_choose__gopurs_runtime_Value_1032977354
}

var cache_chosen gopurs_runtime.Value
var once_chosen sync.Once
func Get_chosen() gopurs_runtime.Value {
	once_chosen.Do(func() {
		cache_chosen = gopurs_runtime.Func(func(dictDecide_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_chosen(dictDecide_0_box)
})
	})
	return cache_chosen
}

func Call_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_chooseOp(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
divideOp_1_0 := gopurs_runtime.Apply(pkg_Data_Divide.Get_divideOp(), dictSemigroup_0)
_ = divideOp_1_0
return gopurs_runtime.RecordDict2("Divide0", "choose", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return divideOp_1_0
}), gopurs_runtime.Func4(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value, x_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_1 := gopurs_runtime.Apply(f_2, x_5)
_ = __local_var_6_1
var __t2 gopurs_runtime.Value
{
if (__local_var_6_1.Type == 9 && __local_var_6_1.IntVal == 3711209382) {
__t2 = gopurs_runtime.Apply(v_3, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_6_1.UnsafePtr).V0)
goto end_branch_2
} else {

}
}
{
if (__local_var_6_1.Type == 9 && __local_var_6_1.IntVal == 2465973597) {
__t2 = gopurs_runtime.Apply(v1_4, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_6_1.UnsafePtr).V0)
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}))
}

func Call_choose(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "choose")
}

func Call_choose__gopurs_runtime_Value_1032977354(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "choose")
}

func Call_chosen(dictDecide_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictDecide_0 gopurs_runtime.Value = dictDecide_0_loop
_ = dictDecide_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDecide_0, "choose"), Get_identity())
}


