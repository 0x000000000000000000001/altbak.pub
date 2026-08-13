package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Decide_identity gopurs_runtime.Value
var once_Data_Decide_identity sync.Once
func Get_Data_Decide_identity() gopurs_runtime.Value {
	once_Data_Decide_identity.Do(func() {
		cache_Data_Decide_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Decide_identity(x_0_box)
})
	})
	return cache_Data_Decide_identity
}

var cache_Data_Decide_Decide_dollarDict gopurs_runtime.Value
var once_Data_Decide_Decide_dollarDict sync.Once
func Get_Data_Decide_Decide_dollarDict() gopurs_runtime.Value {
	once_Data_Decide_Decide_dollarDict.Do(func() {
		cache_Data_Decide_Decide_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Decide_Decide_dollarDict(x_0_box)
})
	})
	return cache_Data_Decide_Decide_dollarDict
}

var cache_Data_Decide_choosePredicate gopurs_runtime.Value
var once_Data_Decide_choosePredicate sync.Once
func Get_Data_Decide_choosePredicate() gopurs_runtime.Value {
	once_Data_Decide_choosePredicate.Do(func() {
		cache_Data_Decide_choosePredicate = gopurs_runtime.Value{Type: 9, IntVal: 1618621146, UnsafePtr: unsafe.Pointer(&Constructor_Data_Decide_Decide{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2642321722, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Divide_Divide](Get_Data_Divide_dividePredicate()))}
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_0 -> gopurs_runtime.Value
__local_var_4_0 := gopurs_runtime.Apply(f_0, x_3)
_ = __local_var_4_0
var __t1 gopurs_runtime.Value
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 3711209382) {
__t1 = gopurs_runtime.Apply(v_1, (*Constructor_Data_Either_Left)(__local_var_4_0.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 2465973597) {
__t1 = gopurs_runtime.Apply(v1_2, (*Constructor_Data_Either_Right)(__local_var_4_0.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
})
})
})
})})}
	})
	return cache_Data_Decide_choosePredicate
}

var cache_Data_Decide_chooseOp gopurs_runtime.Value
var once_Data_Decide_chooseOp sync.Once
func Get_Data_Decide_chooseOp() gopurs_runtime.Value {
	once_Data_Decide_chooseOp.Do(func() {
		cache_Data_Decide_chooseOp = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Decide_chooseOp(dictSemigroup_0_box)
})
	})
	return cache_Data_Decide_chooseOp
}

var cache_Data_Decide_chooseEquivalence gopurs_runtime.Value
var once_Data_Decide_chooseEquivalence sync.Once
func Get_Data_Decide_chooseEquivalence() gopurs_runtime.Value {
	once_Data_Decide_chooseEquivalence.Do(func() {
		cache_Data_Decide_chooseEquivalence = gopurs_runtime.Value{Type: 9, IntVal: 1618621146, UnsafePtr: unsafe.Pointer(&Constructor_Data_Decide_Decide{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2642321722, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Divide_Divide](Get_Data_Divide_divideEquivalence()))}
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v2_5_0 -> gopurs_runtime.Value
v2_5_0 := gopurs_runtime.Apply(f_0, a_3)
_ = v2_5_0
var __t5 bool
{
if (v2_5_0.Type == 9 && v2_5_0.IntVal == 3711209382) {
// TAST (Let): v3_6_1 -> gopurs_runtime.Value
v3_6_1 := gopurs_runtime.Apply(f_0, b_4)
_ = v3_6_1
var __t2 bool
{
if (v3_6_1.Type == 9 && v3_6_1.IntVal == 3711209382) {
__t2 = (gopurs_runtime.Apply2(v_1, (*Constructor_Data_Either_Left)(v2_5_0.UnsafePtr).V0, (*Constructor_Data_Either_Left)(v3_6_1.UnsafePtr).V0).IntVal) != (0)
goto end_branch_2
} else {

}
}
{
if (v3_6_1.Type == 9 && v3_6_1.IntVal == 2465973597) {
__t2 = false
goto end_branch_2
} else {

}
}
{
__t2 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_2:
__t5 = __t2
goto end_branch_5
} else {

}
}
{
if (v2_5_0.Type == 9 && v2_5_0.IntVal == 2465973597) {
// TAST (Let): v3_6_3 -> gopurs_runtime.Value
v3_6_3 := gopurs_runtime.Apply(f_0, b_4)
_ = v3_6_3
var __t4 bool
{
if (v3_6_3.Type == 9 && v3_6_3.IntVal == 3711209382) {
__t4 = false
goto end_branch_4
} else {

}
}
{
if (v3_6_3.Type == 9 && v3_6_3.IntVal == 2465973597) {
__t4 = (gopurs_runtime.Apply2(v1_2, (*Constructor_Data_Either_Right)(v2_5_0.UnsafePtr).V0, (*Constructor_Data_Either_Right)(v3_6_3.UnsafePtr).V0).IntVal) != (0)
goto end_branch_4
} else {

}
}
{
__t4 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
__t5 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_5:
return gopurs_runtime.Bool(__t5)
})
})
})
})
})})}
	})
	return cache_Data_Decide_chooseEquivalence
}

var cache_Data_Decide_chooseComparison gopurs_runtime.Value
var once_Data_Decide_chooseComparison sync.Once
func Get_Data_Decide_chooseComparison() gopurs_runtime.Value {
	once_Data_Decide_chooseComparison.Do(func() {
		cache_Data_Decide_chooseComparison = gopurs_runtime.Value{Type: 9, IntVal: 1618621146, UnsafePtr: unsafe.Pointer(&Constructor_Data_Decide_Decide{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2642321722, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Divide_Divide](Get_Data_Divide_divideComparison()))}
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v2_5_0 -> gopurs_runtime.Value
v2_5_0 := gopurs_runtime.Apply(f_0, a_3)
_ = v2_5_0
var __t5 uint32
{
if (v2_5_0.Type == 9 && v2_5_0.IntVal == 3711209382) {
// TAST (Let): v3_6_1 -> gopurs_runtime.Value
v3_6_1 := gopurs_runtime.Apply(f_0, b_4)
_ = v3_6_1
var __t2 uint32
{
if (v3_6_1.Type == 9 && v3_6_1.IntVal == 3711209382) {
__t2 = uint32(gopurs_runtime.Apply2(v_1, (*Constructor_Data_Either_Left)(v2_5_0.UnsafePtr).V0, (*Constructor_Data_Either_Left)(v3_6_1.UnsafePtr).V0).IntVal)
goto end_branch_2
} else {

}
}
{
if (v3_6_1.Type == 9 && v3_6_1.IntVal == 2465973597) {
__t2 = 1527465420
goto end_branch_2
} else {

}
}
{
__t2 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_2:
__t5 = __t2
goto end_branch_5
} else {

}
}
{
if (v2_5_0.Type == 9 && v2_5_0.IntVal == 2465973597) {
// TAST (Let): v3_6_3 -> gopurs_runtime.Value
v3_6_3 := gopurs_runtime.Apply(f_0, b_4)
_ = v3_6_3
var __t4 uint32
{
if (v3_6_3.Type == 9 && v3_6_3.IntVal == 3711209382) {
__t4 = 380165415
goto end_branch_4
} else {

}
}
{
if (v3_6_3.Type == 9 && v3_6_3.IntVal == 2465973597) {
__t4 = uint32(gopurs_runtime.Apply2(v1_2, (*Constructor_Data_Either_Right)(v2_5_0.UnsafePtr).V0, (*Constructor_Data_Either_Right)(v3_6_3.UnsafePtr).V0).IntVal)
goto end_branch_4
} else {

}
}
{
__t4 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
__t5 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t5), UnsafePtr: nil}
})
})
})
})
})})}
	})
	return cache_Data_Decide_chooseComparison
}

var cache_Data_Decide_choose gopurs_runtime.Value
var once_Data_Decide_choose sync.Once
func Get_Data_Decide_choose() gopurs_runtime.Value {
	once_Data_Decide_choose.Do(func() {
		cache_Data_Decide_choose = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Decide_choose(gopurs_runtime.CoerceToStruct[Constructor_Data_Decide_Decide](dict_0_box))
})
	})
	return cache_Data_Decide_choose
}

var cache_Data_Decide_chosen gopurs_runtime.Value
var once_Data_Decide_chosen sync.Once
func Get_Data_Decide_chosen() gopurs_runtime.Value {
	once_Data_Decide_chosen.Do(func() {
		cache_Data_Decide_chosen = gopurs_runtime.Func(func(dictDecide_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Decide_chosen(gopurs_runtime.CoerceToStruct[Constructor_Data_Decide_Decide](dictDecide_0_box))
})
	})
	return cache_Data_Decide_chosen
}

var cache_Data_Decide_choose__2139889126 gopurs_runtime.Value
var once_Data_Decide_choose__2139889126 sync.Once
func Get_Data_Decide_choose__2139889126() gopurs_runtime.Value {
	once_Data_Decide_choose__2139889126.Do(func() {
		cache_Data_Decide_choose__2139889126 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Decide_choose__2139889126(gopurs_runtime.CoerceToStruct[Constructor_Data_Decide_Decide](dict_0_box))
})
	})
	return cache_Data_Decide_choose__2139889126
}

var cache_Data_Decide_choose__3147709126 gopurs_runtime.Value
var once_Data_Decide_choose__3147709126 sync.Once
func Get_Data_Decide_choose__3147709126() gopurs_runtime.Value {
	once_Data_Decide_choose__3147709126.Do(func() {
		cache_Data_Decide_choose__3147709126 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Decide_choose__3147709126(gopurs_runtime.CoerceToStruct[Constructor_Data_Decide_Decide](dict_0_box))
})
	})
	return cache_Data_Decide_choose__3147709126
}

var cache_Data_Decide_chooseComparison__519342811 gopurs_runtime.Value
var once_Data_Decide_chooseComparison__519342811 sync.Once
func Get_Data_Decide_chooseComparison__519342811() gopurs_runtime.Value {
	once_Data_Decide_chooseComparison__519342811.Do(func() {
		cache_Data_Decide_chooseComparison__519342811 = gopurs_runtime.Value{Type: 9, IntVal: 1618621146, UnsafePtr: unsafe.Pointer(&Constructor_Data_Decide_Decide{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2642321722, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Divide_Divide](Get_Data_Divide_divideComparison()))}
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v2_5_0 -> gopurs_runtime.Value
v2_5_0 := gopurs_runtime.Apply(f_0, a_3)
_ = v2_5_0
var __t5 uint32
{
if (v2_5_0.Type == 9 && v2_5_0.IntVal == 3711209382) {
// TAST (Let): v3_6_1 -> gopurs_runtime.Value
v3_6_1 := gopurs_runtime.Apply(f_0, b_4)
_ = v3_6_1
var __t2 uint32
{
if (v3_6_1.Type == 9 && v3_6_1.IntVal == 3711209382) {
__t2 = uint32(gopurs_runtime.Apply2(v_1, (*Constructor_Data_Either_Left)(v2_5_0.UnsafePtr).V0, (*Constructor_Data_Either_Left)(v3_6_1.UnsafePtr).V0).IntVal)
goto end_branch_2
} else {

}
}
{
if (v3_6_1.Type == 9 && v3_6_1.IntVal == 2465973597) {
__t2 = 1527465420
goto end_branch_2
} else {

}
}
{
__t2 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_2:
__t5 = __t2
goto end_branch_5
} else {

}
}
{
if (v2_5_0.Type == 9 && v2_5_0.IntVal == 2465973597) {
// TAST (Let): v3_6_3 -> gopurs_runtime.Value
v3_6_3 := gopurs_runtime.Apply(f_0, b_4)
_ = v3_6_3
var __t4 uint32
{
if (v3_6_3.Type == 9 && v3_6_3.IntVal == 3711209382) {
__t4 = 380165415
goto end_branch_4
} else {

}
}
{
if (v3_6_3.Type == 9 && v3_6_3.IntVal == 2465973597) {
__t4 = uint32(gopurs_runtime.Apply2(v1_2, (*Constructor_Data_Either_Right)(v2_5_0.UnsafePtr).V0, (*Constructor_Data_Either_Right)(v3_6_3.UnsafePtr).V0).IntVal)
goto end_branch_4
} else {

}
}
{
__t4 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
__t5 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t5), UnsafePtr: nil}
})
})
})
})
})})}
	})
	return cache_Data_Decide_chooseComparison__519342811
}

var cache_Data_Decide_chooseEquivalence__3218367361 gopurs_runtime.Value
var once_Data_Decide_chooseEquivalence__3218367361 sync.Once
func Get_Data_Decide_chooseEquivalence__3218367361() gopurs_runtime.Value {
	once_Data_Decide_chooseEquivalence__3218367361.Do(func() {
		cache_Data_Decide_chooseEquivalence__3218367361 = gopurs_runtime.Value{Type: 9, IntVal: 1618621146, UnsafePtr: unsafe.Pointer(&Constructor_Data_Decide_Decide{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2642321722, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Divide_Divide](Get_Data_Divide_divideEquivalence()))}
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v2_5_0 -> gopurs_runtime.Value
v2_5_0 := gopurs_runtime.Apply(f_0, a_3)
_ = v2_5_0
var __t5 bool
{
if (v2_5_0.Type == 9 && v2_5_0.IntVal == 3711209382) {
// TAST (Let): v3_6_1 -> gopurs_runtime.Value
v3_6_1 := gopurs_runtime.Apply(f_0, b_4)
_ = v3_6_1
var __t2 bool
{
if (v3_6_1.Type == 9 && v3_6_1.IntVal == 3711209382) {
__t2 = (gopurs_runtime.Apply2(v_1, (*Constructor_Data_Either_Left)(v2_5_0.UnsafePtr).V0, (*Constructor_Data_Either_Left)(v3_6_1.UnsafePtr).V0).IntVal) != (0)
goto end_branch_2
} else {

}
}
{
if (v3_6_1.Type == 9 && v3_6_1.IntVal == 2465973597) {
__t2 = false
goto end_branch_2
} else {

}
}
{
__t2 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_2:
__t5 = __t2
goto end_branch_5
} else {

}
}
{
if (v2_5_0.Type == 9 && v2_5_0.IntVal == 2465973597) {
// TAST (Let): v3_6_3 -> gopurs_runtime.Value
v3_6_3 := gopurs_runtime.Apply(f_0, b_4)
_ = v3_6_3
var __t4 bool
{
if (v3_6_3.Type == 9 && v3_6_3.IntVal == 3711209382) {
__t4 = false
goto end_branch_4
} else {

}
}
{
if (v3_6_3.Type == 9 && v3_6_3.IntVal == 2465973597) {
__t4 = (gopurs_runtime.Apply2(v1_2, (*Constructor_Data_Either_Right)(v2_5_0.UnsafePtr).V0, (*Constructor_Data_Either_Right)(v3_6_3.UnsafePtr).V0).IntVal) != (0)
goto end_branch_4
} else {

}
}
{
__t4 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
__t5 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_5:
return gopurs_runtime.Bool(__t5)
})
})
})
})
})})}
	})
	return cache_Data_Decide_chooseEquivalence__3218367361
}

var cache_Data_Decide_choosePredicate__2339472168 gopurs_runtime.Value
var once_Data_Decide_choosePredicate__2339472168 sync.Once
func Get_Data_Decide_choosePredicate__2339472168() gopurs_runtime.Value {
	once_Data_Decide_choosePredicate__2339472168.Do(func() {
		cache_Data_Decide_choosePredicate__2339472168 = gopurs_runtime.Value{Type: 9, IntVal: 1618621146, UnsafePtr: unsafe.Pointer(&Constructor_Data_Decide_Decide{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2642321722, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Divide_Divide](Get_Data_Divide_dividePredicate()))}
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_0 -> gopurs_runtime.Value
__local_var_4_0 := gopurs_runtime.Apply(f_0, x_3)
_ = __local_var_4_0
var __t1 gopurs_runtime.Value
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 3711209382) {
__t1 = gopurs_runtime.Apply(v_1, (*Constructor_Data_Either_Left)(__local_var_4_0.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 2465973597) {
__t1 = gopurs_runtime.Apply(v1_2, (*Constructor_Data_Either_Right)(__local_var_4_0.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
})
})
})
})})}
	})
	return cache_Data_Decide_choosePredicate__2339472168
}

type Constructor_Data_Decide_Decide struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1618621146] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Decide_Decide)(ptr)
		_ = c
		switch key {
		case "Divide0": return gopurs_runtime.Box(c.V0)
		case "choose": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Data_Decide_Decide: " + key)
		}
	}
}


func Call_Data_Decide_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Decide_Decide_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Decide_chooseOp(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
// TAST (Let): divideOp_1_0 -> *Constructor_Data_Divide_Divide
divideOp_1_0 := &Constructor_Data_Divide_Divide{1, gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 85171506, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Contravariant_Contravariant](Get_Data_Op_contravariantOp()))}
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v2_5_1 -> *Constructor_Data_Tuple_Tuple
v2_5_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(f_1, a_4))
_ = v2_5_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), gopurs_runtime.Apply(v_2, (v2_5_1).V0), gopurs_runtime.Apply(v1_3, (v2_5_1).V1))
})
})
})
})}
_ = divideOp_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1618621146, UnsafePtr: unsafe.Pointer(&Constructor_Data_Decide_Decide{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2642321722, UnsafePtr: unsafe.Pointer(divideOp_1_0)}
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_2 -> gopurs_runtime.Value
__local_var_6_2 := gopurs_runtime.Apply(f_2, x_5)
_ = __local_var_6_2
var __t3 gopurs_runtime.Value
{
if (__local_var_6_2.Type == 9 && __local_var_6_2.IntVal == 3711209382) {
__t3 = gopurs_runtime.Apply(v_3, (*Constructor_Data_Either_Left)(__local_var_6_2.UnsafePtr).V0)
goto end_branch_3
} else {

}
}
{
if (__local_var_6_2.Type == 9 && __local_var_6_2.IntVal == 2465973597) {
__t3 = gopurs_runtime.Apply(v1_4, (*Constructor_Data_Either_Right)(__local_var_6_2.UnsafePtr).V0)
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
})
})
})
})})}
}

func Call_Data_Decide_choose(dict_0_loop *Constructor_Data_Decide_Decide) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Decide_Decide = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Decide_chosen(dictDecide_0_loop *Constructor_Data_Decide_Decide) gopurs_runtime.Value {
var dictDecide_0 *Constructor_Data_Decide_Decide = dictDecide_0_loop
_ = dictDecide_0
return gopurs_runtime.Apply(gopurs_runtime.Box(dictDecide_0.V1), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
}

func Call_Data_Decide_choose__2139889126(dict_0_loop *Constructor_Data_Decide_Decide) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Decide_Decide = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Decide_choose__3147709126(dict_0_loop *Constructor_Data_Decide_Decide) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Decide_Decide = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}


