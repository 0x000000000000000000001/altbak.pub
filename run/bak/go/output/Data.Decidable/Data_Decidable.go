package Data_Decidable

import (
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Comparison "gopurs/output/Data.Comparison"
	pkg_Data_Decide "gopurs/output/Data.Decide"
	pkg_Data_Divide "gopurs/output/Data.Divide"
	pkg_Data_Divisible "gopurs/output/Data.Divisible"
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

var cache_lose gopurs_runtime.Value
var once_lose sync.Once
func Get_lose() gopurs_runtime.Value {
	once_lose.Do(func() {
		cache_lose = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lose(gopurs_runtime.CoerceToStruct[Constructor_Decidable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_lose
}

var cache_lost gopurs_runtime.Value
var once_lost sync.Once
func Get_lost() gopurs_runtime.Value {
	once_lost.Do(func() {
		cache_lost = gopurs_runtime.Func(func(dictDecidable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lost(dictDecidable_0_box)
})
	})
	return cache_lost
}

var cache_decidablePredicate gopurs_runtime.Value
var once_decidablePredicate sync.Once
func Get_decidablePredicate() gopurs_runtime.Value {
	once_decidablePredicate.Do(func() {
		cache_decidablePredicate = gopurs_runtime.RecordDict3("Decide0", "Divisible1", "lose", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Decide.Get_choosePredicate()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Divisible.Get_divisiblePredicate()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((Call_absurd__2082956474(gopurs_runtime.Apply(f_0, a_1)).IntVal) != (0))
})
}))
	})
	return cache_decidablePredicate
}

var cache_decidableOp gopurs_runtime.Value
var once_decidableOp sync.Once
func Get_decidableOp() gopurs_runtime.Value {
	once_decidableOp.Do(func() {
		cache_decidableOp = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_decidableOp(dictMonoid_0_box)
})
	})
	return cache_decidableOp
}

var cache_decidableEquivalence gopurs_runtime.Value
var once_decidableEquivalence sync.Once
func Get_decidableEquivalence() gopurs_runtime.Value {
	once_decidableEquivalence.Do(func() {
		cache_decidableEquivalence = gopurs_runtime.RecordDict3("Decide0", "Divisible1", "lose", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Decide.Get_chooseEquivalence()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Divisible.Get_divisibleEquivalence()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
var spin_2_0_1 gopurs_runtime.Value
spin_2_0_1 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_3_loop gopurs_runtime.Value = v_3_loop_val
spin_2_0_1:
for {
if false { continue spin_2_0_1 }
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
v_3_loop = v_3
continue spin_2_0_1
return gopurs_runtime.Value{}
}
}()
})
return gopurs_runtime.Apply(spin_2_0_1, gopurs_runtime.Apply(f_0, a_1))
})
}))
	})
	return cache_decidableEquivalence
}

var cache_decidableComparison gopurs_runtime.Value
var once_decidableComparison sync.Once
func Get_decidableComparison() gopurs_runtime.Value {
	once_decidableComparison.Do(func() {
		cache_decidableComparison = gopurs_runtime.RecordDict3("Decide0", "Divisible1", "lose", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Decide.Get_chooseComparison()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Divisible.Get_divisibleComparison()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Call_absurd__977285408(gopurs_runtime.Apply(f_0, a_1)).IntVal)), UnsafePtr: nil}
})
})
}))
	})
	return cache_decidableComparison
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

var cache_lose__3306256519 gopurs_runtime.Value
var once_lose__3306256519 sync.Once
func Get_lose__3306256519() gopurs_runtime.Value {
	once_lose__3306256519.Do(func() {
		cache_lose__3306256519 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lose__3306256519(gopurs_runtime.CoerceToStruct[Constructor_Decidable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_lose__3306256519
}

var cache_lose__2926568423 gopurs_runtime.Value
var once_lose__2926568423 sync.Once
func Get_lose__2926568423() gopurs_runtime.Value {
	once_lose__2926568423.Do(func() {
		cache_lose__2926568423 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lose__2926568423(gopurs_runtime.CoerceToStruct[Constructor_Decidable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_lose__2926568423
}

var cache_chooseComparison__253205372 gopurs_runtime.Value
var once_chooseComparison__253205372 sync.Once
func Get_chooseComparison__253205372() gopurs_runtime.Value {
	once_chooseComparison__253205372.Do(func() {
		cache_chooseComparison__253205372 = gopurs_runtime.RecordDict2("Divide0", "choose", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_chooseComparison__253205372
}

var cache_chooseEquivalence__667167580 gopurs_runtime.Value
var once_chooseEquivalence__667167580 sync.Once
func Get_chooseEquivalence__667167580() gopurs_runtime.Value {
	once_chooseEquivalence__667167580.Do(func() {
		cache_chooseEquivalence__667167580 = gopurs_runtime.RecordDict2("Divide0", "choose", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_chooseEquivalence__667167580
}

var cache_choosePredicate__1380088508 gopurs_runtime.Value
var once_choosePredicate__1380088508 sync.Once
func Get_choosePredicate__1380088508() gopurs_runtime.Value {
	once_choosePredicate__1380088508.Do(func() {
		cache_choosePredicate__1380088508 = gopurs_runtime.RecordDict2("Divide0", "choose", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_choosePredicate__1380088508
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
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(Get_append__868515608(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(v_1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v2_5_0)}.UnsafePtr).V0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v3_6_1)}.UnsafePtr).V0).IntVal)), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(v1_2, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v2_5_0)}.UnsafePtr).V1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v3_6_1)}.UnsafePtr).V1).IntVal)), UnsafePtr: nil}).IntVal)), UnsafePtr: nil}
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
return gopurs_runtime.Bool((gopurs_runtime.Apply2(Get_conj__3676519832(), gopurs_runtime.Bool((gopurs_runtime.Apply2(v_1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v2_5_0)}.UnsafePtr).V0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v3_6_1)}.UnsafePtr).V0).IntVal) != (0)), gopurs_runtime.Bool((gopurs_runtime.Apply2(v1_2, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v2_5_0)}.UnsafePtr).V1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v3_6_1)}.UnsafePtr).V1).IntVal) != (0))).IntVal) != (0))
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
return gopurs_runtime.Bool((gopurs_runtime.Apply2(Get_conj__3676519832(), gopurs_runtime.Bool((gopurs_runtime.Apply(v_1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v2_4_0)}.UnsafePtr).V0).IntVal) != (0)), gopurs_runtime.Bool((gopurs_runtime.Apply(v1_2, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v2_4_0)}.UnsafePtr).V1).IntVal) != (0))).IntVal) != (0))
})
})
})
}))
	})
	return cache_dividePredicate__3306073532
}

var cache_divisibleComparison__661164760 gopurs_runtime.Value
var once_divisibleComparison__661164760 sync.Once
func Get_divisibleComparison__661164760() gopurs_runtime.Value {
	once_divisibleComparison__661164760.Do(func() {
		cache_divisibleComparison__661164760 = gopurs_runtime.RecordDict2("Divide0", "conquer", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Divide.Get_divideComparison()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
})
}))
	})
	return cache_divisibleComparison__661164760
}

var cache_divisibleEquivalence__4236776696 gopurs_runtime.Value
var once_divisibleEquivalence__4236776696 sync.Once
func Get_divisibleEquivalence__4236776696() gopurs_runtime.Value {
	once_divisibleEquivalence__4236776696.Do(func() {
		cache_divisibleEquivalence__4236776696 = gopurs_runtime.RecordDict2("Divide0", "conquer", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Divide.Get_divideEquivalence()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})
}))
	})
	return cache_divisibleEquivalence__4236776696
}

var cache_divisiblePredicate__1930744184 gopurs_runtime.Value
var once_divisiblePredicate__1930744184 sync.Once
func Get_divisiblePredicate__1930744184() gopurs_runtime.Value {
	once_divisiblePredicate__1930744184.Do(func() {
		cache_divisiblePredicate__1930744184 = gopurs_runtime.RecordDict2("Divide0", "conquer", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Divide.Get_dividePredicate()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
}))
	})
	return cache_divisiblePredicate__1930744184
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

var cache_const__641934996 gopurs_runtime.Value
var once_const__641934996 sync.Once
func Get_const__641934996() gopurs_runtime.Value {
	once_const__641934996.Do(func() {
		cache_const__641934996 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__641934996(a_0_box, v_1_box)
})
	})
	return cache_const__641934996
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

var cache_absurd__2082956474 gopurs_runtime.Value
var once_absurd__2082956474 sync.Once
func Get_absurd__2082956474() gopurs_runtime.Value {
	once_absurd__2082956474.Do(func() {
		cache_absurd__2082956474 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_absurd__2082956474(a_0_box)
})
	})
	return cache_absurd__2082956474
}

var cache_absurd__3279552488 gopurs_runtime.Value
var once_absurd__3279552488 sync.Once
func Get_absurd__3279552488() gopurs_runtime.Value {
	once_absurd__3279552488.Do(func() {
		cache_absurd__3279552488 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_absurd__3279552488(a_0_box)
})
	})
	return cache_absurd__3279552488
}

var cache_absurd__977285408 gopurs_runtime.Value
var once_absurd__977285408 sync.Once
func Get_absurd__977285408() gopurs_runtime.Value {
	once_absurd__977285408.Do(func() {
		cache_absurd__977285408 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_absurd__977285408(a_0_box)
})
	})
	return cache_absurd__977285408
}

var cache_absurd__701346290 gopurs_runtime.Value
var once_absurd__701346290 sync.Once
func Get_absurd__701346290() gopurs_runtime.Value {
	once_absurd__701346290.Do(func() {
		cache_absurd__701346290 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_absurd__701346290(a_0_box)
})
	})
	return cache_absurd__701346290
}

type Constructor_Decidable[T_f any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[336732346] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Decidable[gopurs_runtime.Value])(ptr)
		switch key {
		case "Decide0": return c.V0
		case "Divisible1": return c.V1
		case "lose": return c.V2
		default: panic("Key not found in dictionary Constructor_Decidable: " + key)
		}
	}
}


func Call_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_lose(dict_0_loop *Constructor_Decidable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Decidable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_lost(dictDecidable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictDecidable_0 gopurs_runtime.Value = dictDecidable_0_loop
_ = dictDecidable_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDecidable_0, "lose"), Get_identity())
}

func Call_decidableOp(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
chooseOp_1_0 := gopurs_runtime.Apply(pkg_Data_Decide.Get_chooseOp(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = chooseOp_1_0
divisibleOp_2_1 := gopurs_runtime.Apply(pkg_Data_Divisible.Get_divisibleOp(), dictMonoid_0)
_ = divisibleOp_2_1
return gopurs_runtime.RecordDict3("Decide0", "Divisible1", "lose", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return chooseOp_1_0
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return divisibleOp_2_1
}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
var spin_5_2_0 gopurs_runtime.Value
spin_5_2_0 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_6_loop gopurs_runtime.Value = v_6_loop_val
spin_5_2_0:
for {
if false { continue spin_5_2_0 }
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
v_6_loop = v_6
continue spin_5_2_0
return gopurs_runtime.Value{}
}
}()
})
return gopurs_runtime.Apply(spin_5_2_0, gopurs_runtime.Apply(f_3, a_4))
})
}))
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_lose__3306256519(dict_0_loop *Constructor_Decidable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Decidable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_lose__2926568423(dict_0_loop *Constructor_Decidable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Decidable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
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

func Call_const__641934996(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
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

func Call_absurd__2082956474(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var spin_1_0_2 gopurs_runtime.Value
spin_1_0_2 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
spin_1_0_2:
for {
if false { continue spin_1_0_2 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
v_2_loop = v_2
continue spin_1_0_2
return gopurs_runtime.Value{}
}
}()
})
return gopurs_runtime.Apply(spin_1_0_2, a_0)
}

func Call_absurd__3279552488(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var spin_1_0_3 gopurs_runtime.Value
spin_1_0_3 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
spin_1_0_3:
for {
if false { continue spin_1_0_3 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
v_2_loop = v_2
continue spin_1_0_3
return gopurs_runtime.Value{}
}
}()
})
return gopurs_runtime.Apply(spin_1_0_3, a_0)
}

func Call_absurd__977285408(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var spin_1_0_4 gopurs_runtime.Value
spin_1_0_4 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
spin_1_0_4:
for {
if false { continue spin_1_0_4 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
v_2_loop = v_2
continue spin_1_0_4
return gopurs_runtime.Value{}
}
}()
})
return gopurs_runtime.Apply(spin_1_0_4, a_0)
}

func Call_absurd__701346290(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var spin_1_0_5 gopurs_runtime.Value
spin_1_0_5 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
spin_1_0_5:
for {
if false { continue spin_1_0_5 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
v_2_loop = v_2
continue spin_1_0_5
return gopurs_runtime.Value{}
}
}()
})
return gopurs_runtime.Apply(spin_1_0_5, a_0)
}


