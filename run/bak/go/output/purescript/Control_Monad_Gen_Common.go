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
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_3_1.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_2.V0), Get_Data_NonEmpty_NonEmpty(), gen_6), gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadGen_1.V4), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Data_Ord_max__2538992856(), gopurs_runtime.Int(0), gopurs_runtime.Int((x_7.IntVal) - (1)))
}), gopurs_runtime.Apply4(Get_Control_Monad_Gen_unfoldable(), gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer(dictMonadRec_0)}, gopurs_runtime.Value{Type: 9, IntVal: 2254593219, UnsafePtr: unsafe.Pointer(dictMonadGen_1)}, gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable](dictUnfoldable_5))}, gen_6)))
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
var __t5 gopurs_runtime.Value
{
var __t4 bool
{
if (n_7.FloatVal()) < (bias_5.FloatVal()) {
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
__t5 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_2.V0), Get_Data_Maybe_Just(), gen_6)
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_3.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
}
end_branch_5:
return __t5
}))
})
})
}

func Call_Control_Monad_Gen_Common_genMaybe(dictMonadGen_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen) gopurs_runtime.Value {
var dictMonadGen_0 *Constructor_Control_Monad_Gen_Class_MonadGen = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(Get_Control_Monad_Gen_Common_genMaybe_prime(), gopurs_runtime.Value{Type: 9, IntVal: 2254593219, UnsafePtr: unsafe.Pointer(dictMonadGen_0)}, gopurs_runtime.Float(0.75))
}

func Call_Control_Monad_Gen_Common_genIdentity(dictFunctor_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.Apply(gopurs_runtime.Box(dictFunctor_0.V0), Get_Data_Identity_Identity())
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
var __t4 gopurs_runtime.Value
{
var __t3 bool
{
if (n_7.FloatVal()) < (bias_4.FloatVal()) {
__t3 = true
goto end_branch_3
} else {

}
}
{
__t3 = false
}
end_branch_3:
if __t3 {
__t4 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_2.V0), Get_Data_Either_Left(), genA_5)
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_2.V0), Get_Data_Either_Right(), genB_6)
}
end_branch_4:
return __t4
}))
})
})
})
}

func Call_Control_Monad_Gen_Common_genEither(dictMonadGen_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen) gopurs_runtime.Value {
var dictMonadGen_0 *Constructor_Control_Monad_Gen_Class_MonadGen = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(Get_Control_Monad_Gen_Common_genEither_prime(), gopurs_runtime.Value{Type: 9, IntVal: 2254593219, UnsafePtr: unsafe.Pointer(dictMonadGen_0)}, gopurs_runtime.Float(0.5))
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
var __t4 gopurs_runtime.Value
{
var __t3 bool
{
if (n_7.FloatVal()) < (bias_4.FloatVal()) {
__t3 = true
goto end_branch_3
} else {

}
}
{
__t3 = false
}
end_branch_3:
if __t3 {
__t4 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_2.V0), Get_Data_Either_Left(), genA_5)
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_2.V0), Get_Data_Either_Right(), genB_6)
}
end_branch_4:
return __t4
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
var __t5 gopurs_runtime.Value
{
var __t4 bool
{
if (n_7.FloatVal()) < (bias_5.FloatVal()) {
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
__t5 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_2.V0), Get_Data_Maybe_Just(), gen_6)
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_3.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
}
end_branch_5:
return __t5
}))
})
})
}


