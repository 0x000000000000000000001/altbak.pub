package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_Gen_Common_genTuple gopurs_runtime.Value
var once_Control_Monad_Gen_Common_genTuple sync.Once
func Get_Control_Monad_Gen_Common_genTuple() gopurs_runtime.Value {
	once_Control_Monad_Gen_Common_genTuple.Do(func() {
		cache_Control_Monad_Gen_Common_genTuple = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_Common_genTuple(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dictApply_0_box))
})
	})
	return cache_Control_Monad_Gen_Common_genTuple
}

var cache_Control_Monad_Gen_Common_genNonEmpty gopurs_runtime.Value
var once_Control_Monad_Gen_Common_genNonEmpty sync.Once
func Get_Control_Monad_Gen_Common_genNonEmpty() gopurs_runtime.Value {
	once_Control_Monad_Gen_Common_genNonEmpty.Do(func() {
		cache_Control_Monad_Gen_Common_genNonEmpty = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_Common_genNonEmpty(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dictMonadRec_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen](dictMonadGen_1_box))
})
	})
	return cache_Control_Monad_Gen_Common_genNonEmpty
}

var cache_Control_Monad_Gen_Common_genMaybe_prime gopurs_runtime.Value
var once_Control_Monad_Gen_Common_genMaybe_prime sync.Once
func Get_Control_Monad_Gen_Common_genMaybe_prime() gopurs_runtime.Value {
	once_Control_Monad_Gen_Common_genMaybe_prime.Do(func() {
		cache_Control_Monad_Gen_Common_genMaybe_prime = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_Common_genMaybe_prime(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen](dictMonadGen_0_box))
})
	})
	return cache_Control_Monad_Gen_Common_genMaybe_prime
}

var cache_Control_Monad_Gen_Common_genMaybe gopurs_runtime.Value
var once_Control_Monad_Gen_Common_genMaybe sync.Once
func Get_Control_Monad_Gen_Common_genMaybe() gopurs_runtime.Value {
	once_Control_Monad_Gen_Common_genMaybe.Do(func() {
		cache_Control_Monad_Gen_Common_genMaybe = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_Common_genMaybe(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen](dictMonadGen_0_box))
})
	})
	return cache_Control_Monad_Gen_Common_genMaybe
}

var cache_Control_Monad_Gen_Common_genIdentity gopurs_runtime.Value
var once_Control_Monad_Gen_Common_genIdentity sync.Once
func Get_Control_Monad_Gen_Common_genIdentity() gopurs_runtime.Value {
	once_Control_Monad_Gen_Common_genIdentity.Do(func() {
		cache_Control_Monad_Gen_Common_genIdentity = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_Common_genIdentity(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_0_box))
})
	})
	return cache_Control_Monad_Gen_Common_genIdentity
}

var cache_Control_Monad_Gen_Common_genEither_prime gopurs_runtime.Value
var once_Control_Monad_Gen_Common_genEither_prime sync.Once
func Get_Control_Monad_Gen_Common_genEither_prime() gopurs_runtime.Value {
	once_Control_Monad_Gen_Common_genEither_prime.Do(func() {
		cache_Control_Monad_Gen_Common_genEither_prime = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_Common_genEither_prime(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen](dictMonadGen_0_box))
})
	})
	return cache_Control_Monad_Gen_Common_genEither_prime
}

var cache_Control_Monad_Gen_Common_genEither gopurs_runtime.Value
var once_Control_Monad_Gen_Common_genEither sync.Once
func Get_Control_Monad_Gen_Common_genEither() gopurs_runtime.Value {
	once_Control_Monad_Gen_Common_genEither.Do(func() {
		cache_Control_Monad_Gen_Common_genEither = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_Common_genEither(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen](dictMonadGen_0_box))
})
	})
	return cache_Control_Monad_Gen_Common_genEither
}

var cache_Control_Monad_Gen_Common_genEither_prime__1946557461 gopurs_runtime.Value
var once_Control_Monad_Gen_Common_genEither_prime__1946557461 sync.Once
func Get_Control_Monad_Gen_Common_genEither_prime__1946557461() gopurs_runtime.Value {
	once_Control_Monad_Gen_Common_genEither_prime__1946557461.Do(func() {
		cache_Control_Monad_Gen_Common_genEither_prime__1946557461 = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_Common_genEither_prime__1946557461(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen](dictMonadGen_0_box))
})
	})
	return cache_Control_Monad_Gen_Common_genEither_prime__1946557461
}

var cache_Control_Monad_Gen_Common_genMaybe_prime__1561363431 gopurs_runtime.Value
var once_Control_Monad_Gen_Common_genMaybe_prime__1561363431 sync.Once
func Get_Control_Monad_Gen_Common_genMaybe_prime__1561363431() gopurs_runtime.Value {
	once_Control_Monad_Gen_Common_genMaybe_prime__1561363431.Do(func() {
		cache_Control_Monad_Gen_Common_genMaybe_prime__1561363431 = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_Common_genMaybe_prime__1561363431(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen](dictMonadGen_0_box))
})
	})
	return cache_Control_Monad_Gen_Common_genMaybe_prime__1561363431
}

func Call_Control_Monad_Gen_Common_genTuple(dictApply_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dictApply_0 *Constructor_Control_Apply_Apply = dictApply_0_loop
_ = dictApply_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(dictApply_0.V0), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictApply_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Tuple_Tuple(), a_2), b_3)
})
})
}

func Call_Control_Monad_Gen_Common_genNonEmpty(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec, dictMonadGen_1_loop *Constructor_Control_Monad_Gen_Class_MonadGen) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 *Constructor_Control_Monad_Gen_Class_MonadGen = dictMonadGen_1_loop
_ = dictMonadGen_1
// TAST (Let): Bind1_2_0 -> gopurs_runtime.Value
Bind1_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_1.V0), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{})
_ = Bind1_2_0
// TAST (Let): Apply0_3_1 -> *Constructor_Control_Apply_Apply
Apply0_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_2_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_3_1
// TAST (Let): Functor0_4_2 -> *Constructor_Data_Functor_Functor
Functor0_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_2_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_2
return gopurs_runtime.Func(func(dictUnfoldable_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(gen_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Monad0_7_3 -> gopurs_runtime.Value
Monad0_7_3 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_1.V0), gopurs_runtime.Value{})
_ = Monad0_7_3
// TAST (Let): pure_8_4 -> gopurs_runtime.Value
pure_8_4 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_7_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_8_4
// TAST (Let): Bind1_9_5 -> *Constructor_Control_Bind_Bind
Bind1_9_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_7_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_5
// TAST (Let): __local_var_10_7 -> gopurs_runtime.Value
__local_var_10_7 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadRec_0.V1), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t11 gopurs_runtime.Value
{
var __t10 bool
{
if ((*Constructor_Data_Tuple_Tuple)(v_10.UnsafePtr).V1.IntVal) > (0) {
__t10 = false
goto end_branch_10
} else {

}
}
{
__t10 = true
}
end_branch_10:
if __t10 {
__t11 = gopurs_runtime.Apply(pure_8_4, gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done{1, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons]((*Constructor_Data_Tuple_Tuple)(v_10.UnsafePtr).V0))}})})
goto end_branch_11
} else {

}
}
{
// TAST (Let): __local_var_11_8 -> gopurs_runtime.Value
__local_var_11_8 := (*Constructor_Data_Tuple_Tuple)(v_10.UnsafePtr).V0
_ = __local_var_11_8
// TAST (Let): __local_var_12_9 -> gopurs_runtime.Value
__local_var_12_9 := (*Constructor_Data_Tuple_Tuple)(v_10.UnsafePtr).V1
_ = __local_var_12_9
__t11 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_5.V1), gen_6, gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_8_4, gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Gen_Cons{1, x_13, gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons](__local_var_11_8)})}, gopurs_runtime.Int((__local_var_12_9.IntVal) - (1))})}})})
}))
}
end_branch_11:
return __t11
}))
_ = __local_var_10_7
// TAST (Let): __local_var_11_12 -> gopurs_runtime.Value
__local_var_11_12 := gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(nil))})
_ = __local_var_11_12
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_3_1.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_2.V0), Get_Data_NonEmpty_NonEmpty(), gen_6), gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadGen_1.V4), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Data_Ord_max__2538992856(), gopurs_runtime.Int(0), gopurs_runtime.Int((x_7.IntVal) - (1)))
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_7_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictUnfoldable_5, "unfoldr"), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 *Constructor_Data_Maybe_Just
{
if (v_10.Type == 9 && v_10.IntVal == 759514854 && v_10.UnsafePtr == nil) {
__t6 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_6
} else {

}
}
{
if (v_10.Type == 9 && v_10.IntVal == 759514854 && v_10.UnsafePtr != nil) {
__t6 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Control_Monad_Gen_Cons)(v_10.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(v_10.UnsafePtr).V1)}})}}
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t6)}
})), gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_1.V5), gopurs_runtime.Func(func(x_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_7, gopurs_runtime.Apply(__local_var_11_12, x_12))
})))))
})
})
}

func Call_Control_Monad_Gen_Common_genMaybe_prime(dictMonadGen_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen) gopurs_runtime.Value {
var dictMonadGen_0 *Constructor_Control_Monad_Gen_Class_MonadGen = dictMonadGen_0_loop
_ = dictMonadGen_0
// TAST (Let): Monad0_1_0 -> gopurs_runtime.Value
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_0.V0), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): Bind1_2_1 -> *Constructor_Control_Bind_Bind
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
// TAST (Let): Functor0_3_2 -> *Constructor_Data_Functor_Functor
Functor0_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_2
// TAST (Let): Applicative0_4_3 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_3
return gopurs_runtime.Func(func(bias_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(gen_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadGen_0.V2), gopurs_runtime.Float(0.0), gopurs_runtime.Float(1.0)), gopurs_runtime.Func(func(n_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
var __t5 bool
{
var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Float(n_7.FloatVal()), gopurs_runtime.Float(bias_5.FloatVal()))
if (uint32(__t_tag_4.IntVal) == 1527465420) {
__t5 = true
goto end_branch_5
} else {

}
}
{
__t5 = false
}
end_branch_5:
if __t5 {
__t6 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_2.V0), Get_Data_Maybe_Just(), gen_6)
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_3.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
}
end_branch_6:
return __t6
}))
})
})
}

func Call_Control_Monad_Gen_Common_genMaybe(dictMonadGen_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen) gopurs_runtime.Value {
var dictMonadGen_0 *Constructor_Control_Monad_Gen_Class_MonadGen = dictMonadGen_0_loop
_ = dictMonadGen_0
// TAST (Let): Monad0_1_0 -> gopurs_runtime.Value
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_0.V0), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): Bind1_2_1 -> *Constructor_Control_Bind_Bind
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
// TAST (Let): Functor0_3_2 -> *Constructor_Data_Functor_Functor
Functor0_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_2
// TAST (Let): Applicative0_4_3 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_3
return gopurs_runtime.Func(func(gen_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadGen_0.V2), gopurs_runtime.Float(0.0), gopurs_runtime.Float(1.0)), gopurs_runtime.Func(func(n_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
var __t5 bool
{
var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Float(n_6.FloatVal()), gopurs_runtime.Float(0.75))
if (uint32(__t_tag_4.IntVal) == 1527465420) {
__t5 = true
goto end_branch_5
} else {

}
}
{
__t5 = false
}
end_branch_5:
if __t5 {
__t6 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_2.V0), Get_Data_Maybe_Just(), gen_5)
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_3.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
}
end_branch_6:
return __t6
}))
})
}

func Call_Control_Monad_Gen_Common_genIdentity(dictFunctor_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.Apply(gopurs_runtime.Box(dictFunctor_0.V0), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
}

func Call_Control_Monad_Gen_Common_genEither_prime(dictMonadGen_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen) gopurs_runtime.Value {
var dictMonadGen_0 *Constructor_Control_Monad_Gen_Class_MonadGen = dictMonadGen_0_loop
_ = dictMonadGen_0
// TAST (Let): Monad0_1_0 -> gopurs_runtime.Value
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_0.V0), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): Bind1_2_1 -> *Constructor_Control_Bind_Bind
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
// TAST (Let): Functor0_3_2 -> *Constructor_Data_Functor_Functor
Functor0_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_2
return gopurs_runtime.Func(func(bias_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(genA_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(genB_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadGen_0.V2), gopurs_runtime.Float(0.0), gopurs_runtime.Float(1.0)), gopurs_runtime.Func(func(n_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
var __t4 bool
{
var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Float(n_7.FloatVal()), gopurs_runtime.Float(bias_4.FloatVal()))
if (uint32(__t_tag_3.IntVal) == 1527465420) {
__t4 = true
goto end_branch_4
} else {

}
}
{
__t4 = false
}
end_branch_4:
if __t4 {
__t5 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_2.V0), Get_Data_Either_Left(), genA_5)
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_2.V0), Get_Data_Either_Right(), genB_6)
}
end_branch_5:
return __t5
}))
})
})
})
}

func Call_Control_Monad_Gen_Common_genEither(dictMonadGen_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen) gopurs_runtime.Value {
var dictMonadGen_0 *Constructor_Control_Monad_Gen_Class_MonadGen = dictMonadGen_0_loop
_ = dictMonadGen_0
// TAST (Let): Monad0_1_0 -> gopurs_runtime.Value
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_0.V0), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): Bind1_2_1 -> *Constructor_Control_Bind_Bind
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
// TAST (Let): Functor0_3_2 -> *Constructor_Data_Functor_Functor
Functor0_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_2
return gopurs_runtime.Func(func(genA_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(genB_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadGen_0.V2), gopurs_runtime.Float(0.0), gopurs_runtime.Float(1.0)), gopurs_runtime.Func(func(n_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
var __t4 bool
{
var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Float(n_6.FloatVal()), gopurs_runtime.Float(0.5))
if (uint32(__t_tag_3.IntVal) == 1527465420) {
__t4 = true
goto end_branch_4
} else {

}
}
{
__t4 = false
}
end_branch_4:
if __t4 {
__t5 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_2.V0), Get_Data_Either_Left(), genA_4)
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_2.V0), Get_Data_Either_Right(), genB_5)
}
end_branch_5:
return __t5
}))
})
})
}

func Call_Control_Monad_Gen_Common_genEither_prime__1946557461(dictMonadGen_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen) gopurs_runtime.Value {
var dictMonadGen_0 *Constructor_Control_Monad_Gen_Class_MonadGen = dictMonadGen_0_loop
_ = dictMonadGen_0
// TAST (Let): Monad0_1_0 -> gopurs_runtime.Value
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_0.V0), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): Bind1_2_1 -> *Constructor_Control_Bind_Bind
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
// TAST (Let): Functor0_3_2 -> *Constructor_Data_Functor_Functor
Functor0_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_2
return gopurs_runtime.Func(func(bias_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(genA_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(genB_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadGen_0.V2), gopurs_runtime.Float(0.0), gopurs_runtime.Float(1.0)), gopurs_runtime.Func(func(n_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
var __t4 bool
{
var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Float(n_7.FloatVal()), gopurs_runtime.Float(bias_4.FloatVal()))
if (uint32(__t_tag_3.IntVal) == 1527465420) {
__t4 = true
goto end_branch_4
} else {

}
}
{
__t4 = false
}
end_branch_4:
if __t4 {
__t5 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_2.V0), Get_Data_Either_Left(), genA_5)
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_2.V0), Get_Data_Either_Right(), genB_6)
}
end_branch_5:
return __t5
}))
})
})
})
}

func Call_Control_Monad_Gen_Common_genMaybe_prime__1561363431(dictMonadGen_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen) gopurs_runtime.Value {
var dictMonadGen_0 *Constructor_Control_Monad_Gen_Class_MonadGen = dictMonadGen_0_loop
_ = dictMonadGen_0
// TAST (Let): Monad0_1_0 -> gopurs_runtime.Value
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_0.V0), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): Bind1_2_1 -> *Constructor_Control_Bind_Bind
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
// TAST (Let): Functor0_3_2 -> *Constructor_Data_Functor_Functor
Functor0_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_2
// TAST (Let): Applicative0_4_3 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_3
return gopurs_runtime.Func(func(bias_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(gen_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadGen_0.V2), gopurs_runtime.Float(0.0), gopurs_runtime.Float(1.0)), gopurs_runtime.Func(func(n_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
var __t5 bool
{
var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Float(n_7.FloatVal()), gopurs_runtime.Float(bias_5.FloatVal()))
if (uint32(__t_tag_4.IntVal) == 1527465420) {
__t5 = true
goto end_branch_5
} else {

}
}
{
__t5 = false
}
end_branch_5:
if __t5 {
__t6 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_2.V0), Get_Data_Maybe_Just(), gen_6)
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_3.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
}
end_branch_6:
return __t6
}))
})
})
}


