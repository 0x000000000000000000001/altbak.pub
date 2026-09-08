package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_Gen_monoidAdditive gopurs_runtime.Value
var once_Control_Monad_Gen_monoidAdditive sync.Once
func Get_Control_Monad_Gen_monoidAdditive() gopurs_runtime.Value {
	once_Control_Monad_Gen_monoidAdditive.Do(func() {
		cache_Control_Monad_Gen_monoidAdditive = func() gopurs_runtime.Value {
// TAST (Let): semigroupAdditive1_0_0 shape=LitRecord bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar a)])
semigroupAdditive1_0_0 := (&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float((v_0.FloatVal()) + (v1_1.FloatVal()))
})})
_ = semigroupAdditive1_0_0
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupAdditive1_0_0)}
}), gopurs_runtime.Float(0.0)}))}
}()
	})
	return cache_Control_Monad_Gen_monoidAdditive
}

var cache_Control_Monad_Gen_Cons gopurs_runtime.Value
var once_Control_Monad_Gen_Cons sync.Once
func Get_Control_Monad_Gen_Cons() gopurs_runtime.Value {
	once_Control_Monad_Gen_Cons.Do(func() {
		cache_Control_Monad_Gen_Cons = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]{1, value0, gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]](value1)}))}
})
})
	})
	return cache_Control_Monad_Gen_Cons
}

var cache_Control_Monad_Gen_Nil gopurs_runtime.Value
var once_Control_Monad_Gen_Nil sync.Once
func Get_Control_Monad_Gen_Nil() gopurs_runtime.Value {
	once_Control_Monad_Gen_Nil.Do(func() {
		cache_Control_Monad_Gen_Nil = gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value])(nil))}
	})
	return cache_Control_Monad_Gen_Nil
}

var cache_Control_Monad_Gen_FreqSemigroup gopurs_runtime.Value
var once_Control_Monad_Gen_FreqSemigroup sync.Once
func Get_Control_Monad_Gen_FreqSemigroup() gopurs_runtime.Value {
	once_Control_Monad_Gen_FreqSemigroup.Do(func() {
		cache_Control_Monad_Gen_FreqSemigroup = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Control_Monad_Gen_FreqSemigroup(x_0_box)
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()
})
	})
	return cache_Control_Monad_Gen_FreqSemigroup
}

var cache_Control_Monad_Gen_unfoldable gopurs_runtime.Value
var once_Control_Monad_Gen_unfoldable sync.Once
func Get_Control_Monad_Gen_unfoldable() gopurs_runtime.Value {
	once_Control_Monad_Gen_unfoldable.Do(func() {
		cache_Control_Monad_Gen_unfoldable = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_unfoldable(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_1_box))
})
	})
	return cache_Control_Monad_Gen_unfoldable
}

var cache_Control_Monad_Gen_semigroupFreqSemigroup gopurs_runtime.Value
var once_Control_Monad_Gen_semigroupFreqSemigroup sync.Once
func Get_Control_Monad_Gen_semigroupFreqSemigroup() gopurs_runtime.Value {
	once_Control_Monad_Gen_semigroupFreqSemigroup.Do(func() {
		cache_Control_Monad_Gen_semigroupFreqSemigroup = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, pos_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v2_3_0 shape=App(Other) bindingType=(ADT ["Data","Tuple","Tuple"] [(ADT ["Data","Maybe","Maybe"] [Number]), (TypeVar a)])
v2_3_0 := Rebox_Control_Monad_Gen_138441832_1087394609(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(v_0, gopurs_runtime.Float(pos_2.FloatVal()))))
_ = v2_3_0
var __t2 *Constructor_Data_Tuple_Tuple[*Constructor_Data_Maybe_Just[float64], gopurs_runtime.Value]
{
var __t_tag_1 *Constructor_Data_Maybe_Just[float64] = (v2_3_0).V0
if (__t_tag_1 != nil) {
__t2 = Rebox_Control_Monad_Gen_138441832_1087394609(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(v1_1, gopurs_runtime.Float(((v2_3_0).V0).V0))))
goto end_branch_2
} else {

}
}
{
__t2 = v2_3_0
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_Gen_1087394609_138441832(__t2))}
})}))}
	})
	return cache_Control_Monad_Gen_semigroupFreqSemigroup
}

var cache_Control_Monad_Gen_getFreqVal gopurs_runtime.Value
var once_Control_Monad_Gen_getFreqVal sync.Once
func Get_Control_Monad_Gen_getFreqVal() gopurs_runtime.Value {
	once_Control_Monad_Gen_getFreqVal.Do(func() {
		cache_Control_Monad_Gen_getFreqVal = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_getFreqVal(v_0_box, x_1_box.FloatVal())
})
	})
	return cache_Control_Monad_Gen_getFreqVal
}

var cache_Control_Monad_Gen_fromIndex gopurs_runtime.Value
var once_Control_Monad_Gen_fromIndex sync.Once
func Get_Control_Monad_Gen_fromIndex() gopurs_runtime.Value {
	once_Control_Monad_Gen_fromIndex.Do(func() {
		cache_Control_Monad_Gen_fromIndex = gopurs_runtime.Func(func(dictFoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_fromIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box))
})
	})
	return cache_Control_Monad_Gen_fromIndex
}

var cache_Control_Monad_Gen_fromIndex__436997833 gopurs_runtime.Value
var once_Control_Monad_Gen_fromIndex__436997833 sync.Once
func Get_Control_Monad_Gen_fromIndex__436997833() gopurs_runtime.Value {
	once_Control_Monad_Gen_fromIndex__436997833.Do(func() {
		cache_Control_Monad_Gen_fromIndex__436997833 = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_fromIndex__436997833(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](dictFoldable1_0_box), __eta_norm_0_1_box.IntVal)
})
	})
	return cache_Control_Monad_Gen_fromIndex__436997833
}

var cache_Control_Monad_Gen_oneOf gopurs_runtime.Value
var once_Control_Monad_Gen_oneOf sync.Once
func Get_Control_Monad_Gen_oneOf() gopurs_runtime.Value {
	once_Control_Monad_Gen_oneOf.Do(func() {
		cache_Control_Monad_Gen_oneOf = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_oneOf(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_0_box))
})
	})
	return cache_Control_Monad_Gen_oneOf
}

var cache_Control_Monad_Gen_oneOf__2664930405 gopurs_runtime.Value
var once_Control_Monad_Gen_oneOf__2664930405 sync.Once
func Get_Control_Monad_Gen_oneOf__2664930405() gopurs_runtime.Value {
	once_Control_Monad_Gen_oneOf__2664930405.Do(func() {
		cache_Control_Monad_Gen_oneOf__2664930405 = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_oneOf__2664930405(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_0_box))
})
	})
	return cache_Control_Monad_Gen_oneOf__2664930405
}

var cache_Control_Monad_Gen_freqSemigroup gopurs_runtime.Value
var once_Control_Monad_Gen_freqSemigroup sync.Once
func Get_Control_Monad_Gen_freqSemigroup() gopurs_runtime.Value {
	once_Control_Monad_Gen_freqSemigroup.Do(func() {
		cache_Control_Monad_Gen_freqSemigroup = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Control_Monad_Gen_freqSemigroup(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[float64, gopurs_runtime.Value]](v_0_box))
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()
})
	})
	return cache_Control_Monad_Gen_freqSemigroup
}

var cache_Control_Monad_Gen_frequency gopurs_runtime.Value
var once_Control_Monad_Gen_frequency sync.Once
func Get_Control_Monad_Gen_frequency() gopurs_runtime.Value {
	once_Control_Monad_Gen_frequency.Do(func() {
		cache_Control_Monad_Gen_frequency = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_frequency(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_0_box))
})
	})
	return cache_Control_Monad_Gen_frequency
}

var cache_Control_Monad_Gen_filtered gopurs_runtime.Value
var once_Control_Monad_Gen_filtered sync.Once
func Get_Control_Monad_Gen_filtered() gopurs_runtime.Value {
	once_Control_Monad_Gen_filtered.Do(func() {
		cache_Control_Monad_Gen_filtered = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_filtered(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_1_box))
})
	})
	return cache_Control_Monad_Gen_filtered
}

var cache_Control_Monad_Gen_suchThat gopurs_runtime.Value
var once_Control_Monad_Gen_suchThat sync.Once
func Get_Control_Monad_Gen_suchThat() gopurs_runtime.Value {
	once_Control_Monad_Gen_suchThat.Do(func() {
		cache_Control_Monad_Gen_suchThat = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_suchThat(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_1_box))
})
	})
	return cache_Control_Monad_Gen_suchThat
}

var cache_Control_Monad_Gen_elements gopurs_runtime.Value
var once_Control_Monad_Gen_elements sync.Once
func Get_Control_Monad_Gen_elements() gopurs_runtime.Value {
	once_Control_Monad_Gen_elements.Do(func() {
		cache_Control_Monad_Gen_elements = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_elements(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_0_box))
})
	})
	return cache_Control_Monad_Gen_elements
}

var cache_Control_Monad_Gen_choose gopurs_runtime.Value
var once_Control_Monad_Gen_choose sync.Once
func Get_Control_Monad_Gen_choose() gopurs_runtime.Value {
	once_Control_Monad_Gen_choose.Do(func() {
		cache_Control_Monad_Gen_choose = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_choose(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_0_box))
})
	})
	return cache_Control_Monad_Gen_choose
}

type Constructor_Control_Monad_Gen_Cons[T_a any] struct {
	Rc uint32
	V0 T_a
	V1 *Constructor_Control_Monad_Gen_Cons[T_a]
}


type Constructor_Control_Monad_Gen_Nil[T_a any] struct {
	Rc uint32
}


func Call_Control_Monad_Gen_FreqSemigroup(x_0_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value} {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return func() struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value} {
				_v := x_0
				_p := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)
				return struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{V0: _p.V0, V1: _p.V1}
			}()
}

func Call_Control_Monad_Gen_unfoldable(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value], dictMonadGen_1_loop *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value] = dictMonadGen_1_loop
_ = dictMonadGen_1
// TAST (Let): Monad0_2_0 shape=App(Other) bindingType=Any
Monad0_2_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_1.V0), gopurs_runtime.Value{})
_ = Monad0_2_0
// TAST (Let): pure_3_1 shape=Other bindingType=(Func [(ADT ["Control","Monad","Rec","Class","Step"] [(ADT ["Data","Tuple","Tuple"] [(ADT ["Control","Monad","Gen","LL"] [(TypeVar a)]), Int]), (ADT ["Control","Monad","Gen","LL"] [(TypeVar a)])])] (TypeApp (TypeVar m) [(ADT ["Control","Monad","Rec","Class","Step"] [(ADT ["Data","Tuple","Tuple"] [(ADT ["Control","Monad","Gen","LL"] [(TypeVar a)]), Int]), (ADT ["Control","Monad","Gen","LL"] [(TypeVar a)])])]))
pure_3_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_1
// TAST (Let): Bind1_4_2 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_2
// TAST (Let): Functor0_5_3 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_3
return gopurs_runtime.Func2(func(dictUnfoldable_6 gopurs_runtime.Value, gen_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_7 shape=App(Other) bindingType=Any
__local_var_8_7 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadRec_0.V1), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t10 gopurs_runtime.Value
{
if ((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V1.IntVal) <= (int64(0)) {
__t10 = gopurs_runtime.Apply(pure_3_1, gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_Gen_1429920250_3603546092((&Constructor_Control_Monad_Rec_Class_Done[*Constructor_Data_Tuple_Tuple[*Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value], int64], *Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]]((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0)})))})
goto end_branch_10
} else {

}
}
{
// TAST (Let): __local_var_9_8 shape=Other bindingType=Any
__local_var_9_8 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0
_ = __local_var_9_8
// TAST (Let): __local_var_10_9 shape=Other bindingType=Any
__local_var_10_9 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V1
_ = __local_var_10_9
__t10 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_2.V1), gen_7, gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_1, gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_Gen_3474068486_4008603408((&Constructor_Control_Monad_Rec_Class_Loop[*Constructor_Data_Tuple_Tuple[*Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value], int64], *Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]]{1, Rebox_Control_Monad_Gen_138441832_4216845252((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]{1, x_11, gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]](__local_var_9_8)}))}, gopurs_runtime.Int((__local_var_10_9.IntVal) - (int64(1)))}))})))})
}))
}
end_branch_10:
return __t10
}))
_ = __local_var_8_7
// TAST (Let): __local_var_9_11 shape=App(Var) bindingType=(Func [Int] (ADT ["Data","Tuple","Tuple"] [(ADT ["Control","Monad","Gen","LL"] [(TypeVar a)]), Int]))
__local_var_9_11 := gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value])(nil))}))})
_ = __local_var_9_11
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_3.V0), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictUnfoldable_6, "unfoldr"), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]]]
{
var __t_tag_4 *Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]](v_8)
if (__t_tag_4 == nil) {
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_6
} else {

}
}
{
var __t_tag_5 *Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]](v_8)
if (__t_tag_5 != nil) {
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value])(v_8.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value])(v_8.UnsafePtr).V1)}}))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_6
} else {

}
}
{
__t6 = func() *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]]] { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_Gen_3081497822_3094389156(__t6))}
})), gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_1.V5), gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_7, gopurs_runtime.Apply(__local_var_9_11, x_10))
})))
})
}

func Call_Control_Monad_Gen_getFreqVal(v_0_loop gopurs_runtime.Value, x_1_loop float64) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var x_1 float64 = x_1_loop
_ = x_1
return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(v_0, gopurs_runtime.Float(x_1)).UnsafePtr).V1
}

func Call_Control_Monad_Gen_fromIndex(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
// TAST (Let): Foldable0_1_0 shape=App(Other) bindingType=(ADT ["Data","Foldable","Foldable"] [(TypeVar f)])
Foldable0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictFoldable1_0.V0), gopurs_runtime.Value{}))
_ = Foldable0_1_0
return gopurs_runtime.Func2(func(i_2 gopurs_runtime.Value, xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Control_Monad_Gen_go__go_4_1_0 func(int64, *Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Control_Monad_Gen_go__go_4_1_0
var go__go_4_1_0 gopurs_runtime.Value
_ = go__go_4_1_0
Call_local_Control_Monad_Gen_go__go_4_1_0 = func(v_5_loop int64, v1_6_loop *Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_4_1_0:
for {
if false { continue go__go_4_1_0 }
var v_5 int64 = v_5_loop
_ = v_5
var v1_6 *Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value] = v1_6_loop
_ = v1_6
var __t4 gopurs_runtime.Value
{
if (v1_6 != nil) {
var __t3 gopurs_runtime.Value
{
var __t_tag_2 *Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value] = (v1_6).V1
if (__t_tag_2 == nil) {
__t3 = (v1_6).V0
goto end_branch_3
} else {

}
}
{
if (v_5) <= (int64(0)) {
__t3 = (v1_6).V0
goto end_branch_3
} else {

}
}
{
v_5_loop = (v_5) - (int64(1))
v1_6_loop = (v1_6).V1
continue go__go_4_1_0
__t3 = func() gopurs_runtime.Value { panic("unreachable") }()
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
if (v1_6 == nil) {
__t4 = gopurs_runtime.Apply3(gopurs_runtime.Box(dictFoldable1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Get_Data_Semigroup_Last_semigroupLast()))}, Get_Data_Semigroup_Last_Last(), xs_3)
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}
}
go__go_4_1_0 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_Gen_go__go_4_1_0(v_5_loop_val.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]](v1_6_loop_val))
})
})
return Call_local_Control_Monad_Gen_go__go_4_1_0(i_2.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.Box(Foldable0_1_0.V2), Get_Control_Monad_Gen_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value])(nil))}, xs_3)))
})
}

func Call_Control_Monad_Gen_fromIndex__436997833(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]], __eta_norm_0_1_loop int64) gopurs_runtime.Value {
fromIndex__436997833:
for {
if false { continue fromIndex__436997833 }
var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]] = dictFoldable1_0_loop
_ = dictFoldable1_0
var __eta_norm_0_1 int64 = __eta_norm_0_1_loop
_ = __eta_norm_0_1
// TAST (Let): Foldable0_2_0 shape=App(Other) bindingType=(ADT ["Data","Foldable","Foldable"] [(ADT ["Data","NonEmpty","NonEmpty"] [])])
Foldable0_2_0 := Rebox_Control_Monad_Gen_1680800814_3071895939(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictFoldable1_0.V0), gopurs_runtime.Value{})))
_ = Foldable0_2_0
return gopurs_runtime.Func(func(xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Control_Monad_Gen_go__3236717243_4_1_1 func(int64, *Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Control_Monad_Gen_go__3236717243_4_1_1
var go__3236717243_4_1_1 gopurs_runtime.Value
_ = go__3236717243_4_1_1
Call_local_Control_Monad_Gen_go__3236717243_4_1_1 = func(v_5_loop int64, v1_6_loop *Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__3236717243_4_1_1:
for {
if false { continue go__3236717243_4_1_1 }
var v_5 int64 = v_5_loop
_ = v_5
var v1_6 *Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value] = v1_6_loop
_ = v1_6
var __t4 gopurs_runtime.Value
{
if (v1_6 != nil) {
var __t3 gopurs_runtime.Value
{
var __t_tag_2 *Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value] = (v1_6).V1
if (__t_tag_2 == nil) {
__t3 = (v1_6).V0
goto end_branch_3
} else {

}
}
{
if (v_5) <= (int64(0)) {
__t3 = (v1_6).V0
goto end_branch_3
} else {

}
}
{
v_5_loop = (v_5) - (int64(1))
v1_6_loop = (v1_6).V1
continue go__3236717243_4_1_1
__t3 = func() gopurs_runtime.Value { panic("unreachable") }()
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
if (v1_6 == nil) {
__t4 = gopurs_runtime.Apply3(gopurs_runtime.Box(dictFoldable1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Get_Data_Semigroup_Last_semigroupLast()))}, Get_Data_Semigroup_Last_Last(), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](xs_3))})
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}
}
go__3236717243_4_1_1 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_Gen_go__3236717243_4_1_1(v_5_loop_val.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]](v1_6_loop_val))
})
})
var go__go_5_5_2 gopurs_runtime.Value
_ = go__go_5_5_2
// FALLBACK TCO: isLoop=false len=1
go__go_5_5_2 = gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t10 gopurs_runtime.Value
{
var __t_tag_6 *Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]](v1_7)
if (__t_tag_6 != nil) {
var __t8 gopurs_runtime.Value
{
var __t_tag_7 *Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value] = (*Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value])(v1_7.UnsafePtr).V1
if (__t_tag_7 == nil) {
__t8 = (*Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value])(v1_7.UnsafePtr).V0
goto end_branch_8
} else {

}
}
{
if (v_6.IntVal) <= (int64(0)) {
__t8 = (*Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value])(v1_7.UnsafePtr).V0
goto end_branch_8
} else {

}
}
{
__t8 = Call_local_Control_Monad_Gen_go__3236717243_4_1_1((v_6.IntVal) - (int64(1)), (*Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value])(v1_7.UnsafePtr).V1)
}
end_branch_8:
__t10 = __t8
goto end_branch_10
} else {

}
}
{
var __t_tag_9 *Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]](v1_7)
if (__t_tag_9 == nil) {
__t10 = gopurs_runtime.Apply3(gopurs_runtime.Box(dictFoldable1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Get_Data_Semigroup_Last_semigroupLast()))}, Get_Data_Semigroup_Last_Last(), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](xs_3))})
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
return __t10
})
return Call_local_Control_Monad_Gen_go__3236717243_4_1_1(__eta_norm_0_1, gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.Box(Foldable0_2_0.V2), Get_Control_Monad_Gen_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](xs_3))})))
})
}
}

func Call_Control_Monad_Gen_oneOf(dictMonadGen_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadGen_0 *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value] = dictMonadGen_0_loop
_ = dictMonadGen_0
// TAST (Let): Bind1_1_0 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_0.V0), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_0
return gopurs_runtime.Func(func(dictFoldable1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Foldable0_3_1 shape=App(Other) bindingType=(ADT ["Data","Foldable","Foldable"] [(TypeVar f)])
Foldable0_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_2, "Foldable0"), gopurs_runtime.Value{}))
_ = Foldable0_3_1
return gopurs_runtime.Func(func(xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadGen_0.V3), gopurs_runtime.Int(int64(0)), gopurs_runtime.Int((gopurs_runtime.Apply3(gopurs_runtime.Box(Foldable0_3_1.V1), gopurs_runtime.Func2(func(c_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((int64(1)) + (c_5.IntVal))
}), gopurs_runtime.Int(int64(0)), xs_4).IntVal) - (int64(1)))), gopurs_runtime.Func(func(n_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Monad_Gen_fromIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](dictFoldable1_2)), gopurs_runtime.Int(n_5.IntVal), xs_4)
}))
})
})
}

func Call_Control_Monad_Gen_oneOf__2664930405(dictMonadGen_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
oneOf__2664930405:
for {
if false { continue oneOf__2664930405 }
var dictMonadGen_0 *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value] = dictMonadGen_0_loop
_ = dictMonadGen_0
// TAST (Let): Bind1_1_0 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(ADT ["Prim","Array"] [])])
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_0.V0), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_0
// TAST (Let): Foldable0_2_1 shape=App(Other) bindingType=(ADT ["Data","Foldable","Foldable"] [(ADT ["Data","NonEmpty","NonEmpty"] [])])
Foldable0_2_1 := Rebox_Control_Monad_Gen_1680800814_3071895939(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_Char_Gen_foldable1NonEmpty()).V0), gopurs_runtime.Value{})))
_ = Foldable0_2_1
return gopurs_runtime.Func(func(xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadGen_0.V3), gopurs_runtime.Int(int64(0)), gopurs_runtime.Int((gopurs_runtime.Apply3(gopurs_runtime.Box(Foldable0_2_1.V1), gopurs_runtime.Func2(func(c_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((int64(1)) + (c_4.IntVal))
}), gopurs_runtime.Int(int64(0)), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](xs_3))}).IntVal) - (int64(1)))), gopurs_runtime.Func(func(n_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Monad_Gen_fromIndex(Rebox_Control_Monad_Gen_4217626592_4151366573(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_Char_Gen_foldable1NonEmpty()))), gopurs_runtime.Int(n_4.IntVal), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](xs_3))})
}))
})
}
}

func Call_Control_Monad_Gen_freqSemigroup(v_0_loop *Constructor_Data_Tuple_Tuple[float64, gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value} {
var v_0 *Constructor_Data_Tuple_Tuple[float64, gopurs_runtime.Value] = v_0_loop
_ = v_0
// TAST (Let): __local_var_1_0 shape=Other bindingType=Any
__local_var_1_0 := (v_0).V0
_ = __local_var_1_0
// TAST (Let): __local_var_2_1 shape=Other bindingType=Any
__local_var_2_1 := (v_0).V1
_ = __local_var_2_1
return func() struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value} {
				_v := gopurs_runtime.Func(func(pos_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Tuple_Tuple[*Constructor_Data_Maybe_Just[float64], gopurs_runtime.Value]
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, pos_3, gopurs_runtime.Float(__local_var_1_0))
if ((uint32(__t_tag_2.IntVal) == 1527465420)) != (true) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[*Constructor_Data_Maybe_Just[float64], gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Float((pos_3.FloatVal()) - (__local_var_1_0))}))}, __local_var_2_1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[*Constructor_Data_Maybe_Just[float64], gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}, __local_var_2_1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_Gen_1087394609_138441832(__t3))}
})
				_p := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)
				return struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{V0: _p.V0, V1: _p.V1}
			}()
}

func Call_Control_Monad_Gen_frequency(dictMonadGen_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadGen_0 *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value] = dictMonadGen_0_loop
_ = dictMonadGen_0
// TAST (Let): Bind1_1_0 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_0.V0), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_0
return gopurs_runtime.Func(func(dictFoldable1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): semigroupAdditive1_3_2 shape=LitRecord bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar a)])
semigroupAdditive1_3_2 := (&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float((v_3.FloatVal()) + (v1_4.FloatVal()))
})})
_ = semigroupAdditive1_3_2
// TAST (Let): foldMap_3_1 shape=App(Other) bindingType=(Func [(Func [(ADT ["Data","Tuple","Tuple"] [Number, (TypeApp (TypeVar m) [(TypeVar a)])])] Number), (TypeApp (TypeVar f) [(ADT ["Data","Tuple","Tuple"] [Number, (TypeApp (TypeVar m) [(TypeVar a)])])])] Number)
foldMap_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_2, "Foldable0"), gopurs_runtime.Value{}), "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupAdditive1_3_2)}
}), gopurs_runtime.Float(0.0)}))})
_ = foldMap_3_1
return gopurs_runtime.Func(func(xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_3 shape=App(Other) bindingType=Any
__local_var_5_3 := gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable1_2, "foldMap1"), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Get_Control_Monad_Gen_semigroupFreqSemigroup()))}, Get_Control_Monad_Gen_freqSemigroup(), xs_4)
_ = __local_var_5_3
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadGen_0.V2), gopurs_runtime.Float(0.0), gopurs_runtime.Float(gopurs_runtime.Apply2(foldMap_3_1, Get_Data_Tuple_fst(), xs_4).FloatVal())), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(__local_var_5_3, x_6).UnsafePtr).V1
}))
})
})
}

func Call_Control_Monad_Gen_filtered(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value], dictMonadGen_1_loop *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value] = dictMonadGen_1_loop
_ = dictMonadGen_1
// TAST (Let): Functor0_2_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_1.V0), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_0
return gopurs_runtime.Func(func(gen_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadRec_0.V1), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_0.V0), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](a_5)
if (__t_tag_1 == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, Get_Data_Unit_unit()}))}
goto end_branch_3
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](a_5)
if (__t_tag_2 != nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(a_5.UnsafePtr).V0}))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}), gen_3)
}), Get_Data_Unit_unit())
})
}

func Call_Control_Monad_Gen_suchThat(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value], dictMonadGen_1_loop *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value] = dictMonadGen_1_loop
_ = dictMonadGen_1
// TAST (Let): filtered2_2_0 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m) [(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])])] (TypeApp (TypeVar m) [(TypeVar a)]))
filtered2_2_0 := Call_Control_Monad_Gen_filtered(dictMonadRec_0, dictMonadGen_1)
_ = filtered2_2_0
// TAST (Let): Functor0_3_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_1.V0), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_1
return gopurs_runtime.Func2(func(gen_4 gopurs_runtime.Value, pred_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(filtered2_2_0, gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_1.V0), gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (gopurs_runtime.Apply(pred_5, a_6).IntVal) != (0) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{a_6, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))})
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)}
}), gen_4))
})
}

func Call_Control_Monad_Gen_elements(dictMonadGen_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadGen_0 *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value] = dictMonadGen_0_loop
_ = dictMonadGen_0
// TAST (Let): Monad0_1_0 shape=App(Other) bindingType=Any
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_0.V0), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): Bind1_2_1 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
// TAST (Let): pure_3_2 shape=Other bindingType=(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar a)]))
pure_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_2
return gopurs_runtime.Func(func(dictFoldable1_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Foldable0_5_3 shape=App(Other) bindingType=(ADT ["Data","Foldable","Foldable"] [(TypeVar f)])
Foldable0_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_4, "Foldable0"), gopurs_runtime.Value{}))
_ = Foldable0_5_3
return gopurs_runtime.Func(func(xs_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadGen_0.V3), gopurs_runtime.Int(int64(0)), gopurs_runtime.Int((gopurs_runtime.Apply3(gopurs_runtime.Box(Foldable0_5_3.V1), gopurs_runtime.Func2(func(c_7 gopurs_runtime.Value, v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((int64(1)) + (c_7.IntVal))
}), gopurs_runtime.Int(int64(0)), xs_6).IntVal) - (int64(1)))), gopurs_runtime.Func(func(n_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_2, gopurs_runtime.Apply2(Call_Control_Monad_Gen_fromIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](dictFoldable1_4)), gopurs_runtime.Int(n_7.IntVal), xs_6))
}))
})
})
}

func Call_Control_Monad_Gen_choose(dictMonadGen_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadGen_0 *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value] = dictMonadGen_0_loop
_ = dictMonadGen_0
// TAST (Let): Bind1_1_0 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_0.V0), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_0
// TAST (Let): chooseBool_2_1 shape=Other bindingType=(TypeApp (TypeVar m) [Boolean])
chooseBool_2_1 := gopurs_runtime.Box(dictMonadGen_0.V1)
_ = chooseBool_2_1
return gopurs_runtime.Func2(func(genA_3 gopurs_runtime.Value, genB_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), chooseBool_2_1, gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_5.IntVal) != (0) {
__t2 = genA_3
goto end_branch_2
} else {

}
}
{
__t2 = genB_4
}
end_branch_2:
return __t2
}))
})
}

func Rebox_Control_Monad_Gen_1080505919_138441832(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(in.V1)}
	return out
}

func Rebox_Control_Monad_Gen_1087394609_138441832(in *Constructor_Data_Tuple_Tuple[*Constructor_Data_Maybe_Just[float64], gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_Gen_3240988860_3094389156(in.V0))}
		out.V1 = in.V1
	return out
}

func Rebox_Control_Monad_Gen_138441832_1087394609(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[*Constructor_Data_Maybe_Just[float64], gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[*Constructor_Data_Maybe_Just[float64], gopurs_runtime.Value]{}
		out.V0 = Rebox_Control_Monad_Gen_3094389156_3240988860(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](in.V0))
		out.V1 = in.V1
	return out
}

func Rebox_Control_Monad_Gen_138441832_4216845252(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[*Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value], int64] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[*Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value], int64]{}
		out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]](in.V0)
		out.V1 = in.V1.IntVal
	return out
}

func Rebox_Control_Monad_Gen_1429920250_3603546092(in *Constructor_Control_Monad_Rec_Class_Done[*Constructor_Data_Tuple_Tuple[*Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value], int64], *Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]]) *Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}

func Rebox_Control_Monad_Gen_1680800814_3071895939(in *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) *Constructor_Data_Foldable_Foldable[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Foldable_Foldable[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Control_Monad_Gen_3081497822_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]]]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_Gen_1080505919_138441832(in.V0))}
	return out
}

func Rebox_Control_Monad_Gen_3094389156_3240988860(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[float64] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[float64]{}
		out.V0 = in.V0.FloatVal()
	return out
}

func Rebox_Control_Monad_Gen_3240988860_3094389156(in *Constructor_Data_Maybe_Just[float64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Float(in.V0)
	return out
}

func Rebox_Control_Monad_Gen_3474068486_4008603408(in *Constructor_Control_Monad_Rec_Class_Loop[*Constructor_Data_Tuple_Tuple[*Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value], int64], *Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]]) *Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_Gen_4216845252_138441832(in.V0))}
	return out
}

func Rebox_Control_Monad_Gen_4216845252_138441832(in *Constructor_Data_Tuple_Tuple[*Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value], int64]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(in.V0)}
		out.V1 = gopurs_runtime.Int(in.V1)
	return out
}

func Rebox_Control_Monad_Gen_4217626592_4151366573(in *Constructor_Data_Semigroup_Foldable_Foldable1[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}


