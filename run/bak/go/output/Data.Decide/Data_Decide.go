package Data_Decide

import (
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Comparison "gopurs/output/Data.Comparison"
	pkg_Data_Divide "gopurs/output/Data.Divide"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_Equivalence "gopurs/output/Data.Equivalence"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Predicate "gopurs/output/Data.Predicate"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
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
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
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
})
})
})
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
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
v2_5_0 := gopurs_runtime.Apply(f_0, a_3)
_ = v2_5_0
var __t5 gopurs_runtime.Value
{
if (v2_5_0.Type == 9 && v2_5_0.IntVal == 3711209382) {
v3_6_1 := gopurs_runtime.Apply(f_0, b_4)
_ = v3_6_1
var __t2 gopurs_runtime.Value
{
if (v3_6_1.Type == 9 && v3_6_1.IntVal == 3711209382) {
__t2 = gopurs_runtime.Bool((gopurs_runtime.Apply2(v_1, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5_0.UnsafePtr).V0, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v3_6_1.UnsafePtr).V0).IntVal) != (0))
goto end_branch_2
} else {

}
}
{
if (v3_6_1.Type == 9 && v3_6_1.IntVal == 2465973597) {
__t2 = gopurs_runtime.Bool(false)
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t5 = gopurs_runtime.Bool((__t2.IntVal) != (0))
goto end_branch_5
} else {

}
}
{
if (v2_5_0.Type == 9 && v2_5_0.IntVal == 2465973597) {
v3_6_3 := gopurs_runtime.Apply(f_0, b_4)
_ = v3_6_3
var __t4 gopurs_runtime.Value
{
if (v3_6_3.Type == 9 && v3_6_3.IntVal == 3711209382) {
__t4 = gopurs_runtime.Bool(false)
goto end_branch_4
} else {

}
}
{
if (v3_6_3.Type == 9 && v3_6_3.IntVal == 2465973597) {
__t4 = gopurs_runtime.Bool((gopurs_runtime.Apply2(v1_2, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5_0.UnsafePtr).V0, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v3_6_3.UnsafePtr).V0).IntVal) != (0))
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t5 = gopurs_runtime.Bool((__t4.IntVal) != (0))
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Bool((__t5.IntVal) != (0))
})
})
})
})
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
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
v2_5_0 := gopurs_runtime.Apply(f_0, a_3)
_ = v2_5_0
var __t5 gopurs_runtime.Value
{
if (v2_5_0.Type == 9 && v2_5_0.IntVal == 3711209382) {
v3_6_1 := gopurs_runtime.Apply(f_0, b_4)
_ = v3_6_1
var __t2 gopurs_runtime.Value
{
if (v3_6_1.Type == 9 && v3_6_1.IntVal == 3711209382) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(v_1, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5_0.UnsafePtr).V0, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v3_6_1.UnsafePtr).V0).IntVal)), UnsafePtr: nil}
goto end_branch_2
} else {

}
}
{
if (v3_6_1.Type == 9 && v3_6_1.IntVal == 2465973597) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t5 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__t2.IntVal)), UnsafePtr: nil}
goto end_branch_5
} else {

}
}
{
if (v2_5_0.Type == 9 && v2_5_0.IntVal == 2465973597) {
v3_6_3 := gopurs_runtime.Apply(f_0, b_4)
_ = v3_6_3
var __t4 gopurs_runtime.Value
{
if (v3_6_3.Type == 9 && v3_6_3.IntVal == 3711209382) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
goto end_branch_4
} else {

}
}
{
if (v3_6_3.Type == 9 && v3_6_3.IntVal == 2465973597) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(v1_2, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5_0.UnsafePtr).V0, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v3_6_3.UnsafePtr).V0).IntVal)), UnsafePtr: nil}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t5 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__t4.IntVal)), UnsafePtr: nil}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__t5.IntVal)), UnsafePtr: nil}
})
})
})
})
}))
	})
	return cache_chooseComparison
}

var cache_choose gopurs_runtime.Value
var once_choose sync.Once
func Get_choose() gopurs_runtime.Value {
	once_choose.Do(func() {
		cache_choose = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_choose(gopurs_runtime.CoerceToStruct[Constructor_Decide[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_choose
}

var cache_chosen gopurs_runtime.Value
var once_chosen sync.Once
func Get_chosen() gopurs_runtime.Value {
	once_chosen.Do(func() {
		cache_chosen = gopurs_runtime.Func(func(dictDecide_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_chosen(gopurs_runtime.CoerceToStruct[Constructor_Decide[gopurs_runtime.Value]](dictDecide_0_box))
})
	})
	return cache_chosen
}

var cache_compose__858342840 gopurs_runtime.Value
var once_compose__858342840 sync.Once
func Get_compose__858342840() gopurs_runtime.Value {
	once_compose__858342840.Do(func() {
		cache_compose__858342840 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__858342840(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__858342840
}

var cache_semigroupoidFn__2387483462 gopurs_runtime.Value
var once_semigroupoidFn__2387483462 sync.Once
func Get_semigroupoidFn__2387483462() gopurs_runtime.Value {
	once_semigroupoidFn__2387483462.Do(func() {
		cache_semigroupoidFn__2387483462 = gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(g_1, x_2))
})
})
}))
	})
	return cache_semigroupoidFn__2387483462
}

var cache_contravariantComparison__1065380147 gopurs_runtime.Value
var once_contravariantComparison__1065380147 sync.Once
func Get_contravariantComparison__1065380147() gopurs_runtime.Value {
	once_contravariantComparison__1065380147.Do(func() {
		cache_contravariantComparison__1065380147 = gopurs_runtime.RecordDict1("cmap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(v_1, gopurs_runtime.Apply(f_0, x_2), gopurs_runtime.Apply(f_0, y_3))
})
})
})
}))
	})
	return cache_contravariantComparison__1065380147
}

var cache_choose__2139889126 gopurs_runtime.Value
var once_choose__2139889126 sync.Once
func Get_choose__2139889126() gopurs_runtime.Value {
	once_choose__2139889126.Do(func() {
		cache_choose__2139889126 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_choose__2139889126(gopurs_runtime.CoerceToStruct[Constructor_Decide[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_choose__2139889126
}

var cache_choose__3147709126 gopurs_runtime.Value
var once_choose__3147709126 sync.Once
func Get_choose__3147709126() gopurs_runtime.Value {
	once_choose__3147709126.Do(func() {
		cache_choose__3147709126 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_choose__3147709126(gopurs_runtime.CoerceToStruct[Constructor_Decide[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_choose__3147709126
}

var cache_divideComparison__295028591 gopurs_runtime.Value
var once_divideComparison__295028591 sync.Once
func Get_divideComparison__295028591() gopurs_runtime.Value {
	once_divideComparison__295028591.Do(func() {
		cache_divideComparison__295028591 = gopurs_runtime.RecordDict2("Contravariant0", "divide", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Comparison.Get_contravariantComparison()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
v2_5_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, a_3))
_ = v2_5_0
v3_6_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, b_4))
_ = v3_6_1
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(Get_append__868515608(), gopurs_runtime.Apply2(v_1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v2_5_0)}.UnsafePtr).V0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v3_6_1)}.UnsafePtr).V0), gopurs_runtime.Apply2(v1_2, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v2_5_0)}.UnsafePtr).V1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v3_6_1)}.UnsafePtr).V1)).IntVal)), UnsafePtr: nil}
})
})
})
})
}))
	})
	return cache_divideComparison__295028591
}

var cache_divideEquivalence__2464526773 gopurs_runtime.Value
var once_divideEquivalence__2464526773 sync.Once
func Get_divideEquivalence__2464526773() gopurs_runtime.Value {
	once_divideEquivalence__2464526773.Do(func() {
		cache_divideEquivalence__2464526773 = gopurs_runtime.RecordDict2("Contravariant0", "divide", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Equivalence.Get_contravariantEquivalence()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
v2_5_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, a_3))
_ = v2_5_0
v3_6_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, b_4))
_ = v3_6_1
return gopurs_runtime.Bool((gopurs_runtime.Apply2(Get_conj__3676519832(), gopurs_runtime.Apply2(v_1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v2_5_0)}.UnsafePtr).V0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v3_6_1)}.UnsafePtr).V0), gopurs_runtime.Apply2(v1_2, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v2_5_0)}.UnsafePtr).V1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v3_6_1)}.UnsafePtr).V1)).IntVal) != (0))
})
})
})
})
}))
	})
	return cache_divideEquivalence__2464526773
}

var cache_dividePredicate__3306073532 gopurs_runtime.Value
var once_dividePredicate__3306073532 sync.Once
func Get_dividePredicate__3306073532() gopurs_runtime.Value {
	once_dividePredicate__3306073532.Do(func() {
		cache_dividePredicate__3306073532 = gopurs_runtime.RecordDict2("Contravariant0", "divide", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Predicate.Get_contravariantPredicate()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
v2_4_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, a_3))
_ = v2_4_0
return gopurs_runtime.Bool((gopurs_runtime.Apply2(Get_conj__3676519832(), gopurs_runtime.Apply(v_1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v2_4_0)}.UnsafePtr).V0), gopurs_runtime.Apply(v1_2, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v2_4_0)}.UnsafePtr).V1)).IntVal) != (0))
})
})
})
}))
	})
	return cache_dividePredicate__3306073532
}

var cache_either__1539695579 gopurs_runtime.Value
var once_either__1539695579 sync.Once
func Get_either__1539695579() gopurs_runtime.Value {
	once_either__1539695579.Do(func() {
		cache_either__1539695579 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_either__1539695579(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_either__1539695579
}

var cache_either__2158544585 gopurs_runtime.Value
var once_either__2158544585 sync.Once
func Get_either__2158544585() gopurs_runtime.Value {
	once_either__2158544585.Do(func() {
		cache_either__2158544585 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_either__2158544585(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_either__2158544585
}

var cache_contravariantEquivalence__506233683 gopurs_runtime.Value
var once_contravariantEquivalence__506233683 sync.Once
func Get_contravariantEquivalence__506233683() gopurs_runtime.Value {
	once_contravariantEquivalence__506233683.Do(func() {
		cache_contravariantEquivalence__506233683 = gopurs_runtime.RecordDict1("cmap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(v_1, gopurs_runtime.Apply(f_0, x_2), gopurs_runtime.Apply(f_0, y_3))
})
})
})
}))
	})
	return cache_contravariantEquivalence__506233683
}

var cache_on__3122155169 gopurs_runtime.Value
var once_on__3122155169 sync.Once
func Get_on__3122155169() gopurs_runtime.Value {
	once_on__3122155169.Do(func() {
		cache_on__3122155169 = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_on__3122155169(f_0_box, g_1_box, x_2_box, y_3_box)
})
	})
	return cache_on__3122155169
}

var cache_on__3980724833 gopurs_runtime.Value
var once_on__3980724833 sync.Once
func Get_on__3980724833() gopurs_runtime.Value {
	once_on__3980724833.Do(func() {
		cache_on__3980724833 = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_on__3980724833(f_0_box, g_1_box, x_2_box, y_3_box)
})
	})
	return cache_on__3980724833
}

var cache_on__3556844193 gopurs_runtime.Value
var once_on__3556844193 sync.Once
func Get_on__3556844193() gopurs_runtime.Value {
	once_on__3556844193.Do(func() {
		cache_on__3556844193 = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_on__3556844193(f_0_box, g_1_box, x_2_box, y_3_box)
})
	})
	return cache_on__3556844193
}

var cache_conj__3472268504 gopurs_runtime.Value
var once_conj__3472268504 sync.Once
func Get_conj__3472268504() gopurs_runtime.Value {
	once_conj__3472268504.Do(func() {
		cache_conj__3472268504 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_conj__3472268504(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_conj__3472268504
}

var cache_disj__3676519832 gopurs_runtime.Value
var once_disj__3676519832 sync.Once
func Get_disj__3676519832() gopurs_runtime.Value {
	once_disj__3676519832.Do(func() {
		cache_disj__3676519832 = gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "disj")
	})
	return cache_disj__3676519832
}

var cache_disj__3472268504 gopurs_runtime.Value
var once_disj__3472268504 sync.Once
func Get_disj__3472268504() gopurs_runtime.Value {
	once_disj__3472268504.Do(func() {
		cache_disj__3472268504 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_disj__3472268504(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_disj__3472268504
}

var cache_not__3201284355 gopurs_runtime.Value
var once_not__3201284355 sync.Once
func Get_not__3201284355() gopurs_runtime.Value {
	once_not__3201284355.Do(func() {
		cache_not__3201284355 = gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "not")
	})
	return cache_not__3201284355
}

var cache_not__1505204753 gopurs_runtime.Value
var once_not__1505204753 sync.Once
func Get_not__1505204753() gopurs_runtime.Value {
	once_not__1505204753.Do(func() {
		cache_not__1505204753 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_not__1505204753(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_not__1505204753
}

var cache_contravariantPredicate__2354513683 gopurs_runtime.Value
var once_contravariantPredicate__2354513683 sync.Once
func Get_contravariantPredicate__2354513683() gopurs_runtime.Value {
	once_contravariantPredicate__2354513683.Do(func() {
		cache_contravariantPredicate__2354513683 = gopurs_runtime.RecordDict1("cmap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(f_0, x_2))
})
})
}))
	})
	return cache_contravariantPredicate__2354513683
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

type Constructor_Decide[T_f any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1618621146] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Decide[gopurs_runtime.Value])(ptr)
		switch key {
		case "Divide0": return c.V0
		case "choose": return c.V1
		default: panic("Key not found in dictionary Constructor_Decide: " + key)
		}
	}
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
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
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
})
})
})
}))
}

func Call_choose(dict_0_loop *Constructor_Decide[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Decide[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_chosen(dictDecide_0_loop *Constructor_Decide[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictDecide_0 *Constructor_Decide[gopurs_runtime.Value] = dictDecide_0_loop
_ = dictDecide_0
return gopurs_runtime.Apply(dictDecide_0.V1, Get_identity())
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_choose__2139889126(dict_0_loop *Constructor_Decide[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Decide[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_choose__3147709126(dict_0_loop *Constructor_Decide[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Decide[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_either__1539695579(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(v_0, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_either__2158544585(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(v_0, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_on__3122155169(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var g_1 gopurs_runtime.Value = g_1_loop
_ = g_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return gopurs_runtime.Apply2(f_0, gopurs_runtime.Apply(g_1, x_2), gopurs_runtime.Apply(g_1, y_3))
}

func Call_on__3980724833(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var g_1 gopurs_runtime.Value = g_1_loop
_ = g_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return gopurs_runtime.Apply2(f_0, gopurs_runtime.Apply(g_1, x_2), gopurs_runtime.Apply(g_1, y_3))
}

func Call_on__3556844193(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var g_1 gopurs_runtime.Value = g_1_loop
_ = g_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return gopurs_runtime.Apply2(f_0, gopurs_runtime.Apply(g_1, x_2), gopurs_runtime.Apply(g_1, y_3))
}

func Call_conj__3472268504(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_disj__3472268504(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_not__1505204753(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V4
}

func Call_append__1230318264(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}


