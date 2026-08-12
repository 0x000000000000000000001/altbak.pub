package Control_Monad_Gen_Common

import (
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Control_Monad_Gen "gopurs/output/Control.Monad.Gen"
	pkg_Control_Monad_Gen_Class "gopurs/output/Control.Monad.Gen.Class"
	pkg_Control_Monad_Rec_Class "gopurs/output/Control.Monad.Rec.Class"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Identity "gopurs/output/Data.Identity"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_NonEmpty "gopurs/output/Data.NonEmpty"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_genTuple gopurs_runtime.Value
var once_genTuple sync.Once
func Get_genTuple() gopurs_runtime.Value {
	once_genTuple.Do(func() {
		cache_genTuple = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genTuple(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dictApply_0_box))
})
	})
	return cache_genTuple
}

var cache_genNonEmpty gopurs_runtime.Value
var once_genNonEmpty sync.Once
func Get_genNonEmpty() gopurs_runtime.Value {
	once_genNonEmpty.Do(func() {
		cache_genNonEmpty = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genNonEmpty(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dictMonadGen_1_box))
})
	})
	return cache_genNonEmpty
}

var cache_genMaybe_prime gopurs_runtime.Value
var once_genMaybe_prime sync.Once
func Get_genMaybe_prime() gopurs_runtime.Value {
	once_genMaybe_prime.Do(func() {
		cache_genMaybe_prime = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genMaybe_prime(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dictMonadGen_0_box))
})
	})
	return cache_genMaybe_prime
}

var cache_genMaybe_prime__gopurs_runtime_Value_1561363431 gopurs_runtime.Value
var once_genMaybe_prime__gopurs_runtime_Value_1561363431 sync.Once
func Get_genMaybe_prime__gopurs_runtime_Value_1561363431() gopurs_runtime.Value {
	once_genMaybe_prime__gopurs_runtime_Value_1561363431.Do(func() {
		cache_genMaybe_prime__gopurs_runtime_Value_1561363431 = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genMaybe_prime__gopurs_runtime_Value_1561363431(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dictMonadGen_0_box))
})
	})
	return cache_genMaybe_prime__gopurs_runtime_Value_1561363431
}

var cache_genMaybe gopurs_runtime.Value
var once_genMaybe sync.Once
func Get_genMaybe() gopurs_runtime.Value {
	once_genMaybe.Do(func() {
		cache_genMaybe = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genMaybe(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dictMonadGen_0_box))
})
	})
	return cache_genMaybe
}

var cache_genIdentity gopurs_runtime.Value
var once_genIdentity sync.Once
func Get_genIdentity() gopurs_runtime.Value {
	once_genIdentity.Do(func() {
		cache_genIdentity = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genIdentity(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dictFunctor_0_box))
})
	})
	return cache_genIdentity
}

var cache_genEither_prime gopurs_runtime.Value
var once_genEither_prime sync.Once
func Get_genEither_prime() gopurs_runtime.Value {
	once_genEither_prime.Do(func() {
		cache_genEither_prime = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genEither_prime(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dictMonadGen_0_box))
})
	})
	return cache_genEither_prime
}

var cache_genEither_prime__gopurs_runtime_Value_1946557461 gopurs_runtime.Value
var once_genEither_prime__gopurs_runtime_Value_1946557461 sync.Once
func Get_genEither_prime__gopurs_runtime_Value_1946557461() gopurs_runtime.Value {
	once_genEither_prime__gopurs_runtime_Value_1946557461.Do(func() {
		cache_genEither_prime__gopurs_runtime_Value_1946557461 = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genEither_prime__gopurs_runtime_Value_1946557461(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dictMonadGen_0_box))
})
	})
	return cache_genEither_prime__gopurs_runtime_Value_1946557461
}

var cache_genEither gopurs_runtime.Value
var once_genEither sync.Once
func Get_genEither() gopurs_runtime.Value {
	once_genEither.Do(func() {
		cache_genEither = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genEither(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dictMonadGen_0_box))
})
	})
	return cache_genEither
}

func Call_genTuple(dictApply_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApply_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dictApply_0_loop
_ = dictApply_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, pkg_Data_Tuple.Get_Tuple(), a_2), b_3)
})
})
}

func Call_genNonEmpty(dictMonadRec_0_loop *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value], dictMonadGen_1_loop *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadRec_0 *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value] = dictMonadGen_1_loop
_ = dictMonadGen_1
Bind1_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonadGen_1.V0, gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{})
_ = Bind1_2_0
Apply0_3_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_2_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_3_1
Functor0_4_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_2_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_2
return gopurs_runtime.Func(func(dictUnfoldable_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(gen_6 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_3 := gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: unsafe.Pointer(nil)}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: unsafe.Pointer(nil)}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: unsafe.Pointer(nil)})
_ = __local_var_7_3
return gopurs_runtime.Apply2(Apply0_3_1.V1, gopurs_runtime.Apply2(Functor0_4_2.V0, pkg_Data_NonEmpty.Get_NonEmpty(), gen_6), gopurs_runtime.Apply2(dictMonadGen_1.V4, gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_9_4 := (x_8.IntVal) - (1)
_ = __local_var_9_4
v_10_5 := gopurs_runtime.Apply2(__local_var_7_3, gopurs_runtime.Int(0), gopurs_runtime.Int(__local_var_9_4))
_ = v_10_5
var __t6 gopurs_runtime.Value
{
if (v_10_5.Type == 9 && v_10_5.IntVal == 1527465420) {
__t6 = gopurs_runtime.Int(__local_var_9_4)
goto end_branch_6
} else {

}
}
{
if (v_10_5.Type == 9 && v_10_5.IntVal == 902936544) {
__t6 = gopurs_runtime.Int(0)
goto end_branch_6
} else {

}
}
{
if (v_10_5.Type == 9 && v_10_5.IntVal == 380165415) {
__t6 = gopurs_runtime.Int(0)
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
}), gopurs_runtime.Apply4(pkg_Control_Monad_Gen.Get_unfoldable(), gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer(dictMonadRec_0)}, gopurs_runtime.Value{Type: 9, IntVal: 2254593219, UnsafePtr: unsafe.Pointer(dictMonadGen_1)}, dictUnfoldable_5, gen_6)))
})
})
}

func Call_genMaybe_prime(dictMonadGen_0_loop *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadGen_0 *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value] = dictMonadGen_0_loop
_ = dictMonadGen_0
Monad0_1_0 := gopurs_runtime.Apply(dictMonadGen_0.V0, gopurs_runtime.Value{})
_ = Monad0_1_0
Bind1_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
Functor0_3_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_2
Applicative0_4_3 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_3
return gopurs_runtime.Func(func(bias_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(gen_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply2(dictMonadGen_0.V2, gopurs_runtime.Float(0.0), gopurs_runtime.Float(1.0)), gopurs_runtime.Func(func(n_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
var __t5 gopurs_runtime.Value
{
if (n_7.FloatVal()) < (bias_5.FloatVal()) {
__t5 = gopurs_runtime.Bool(true)
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Bool(false)
}
end_branch_5:
if (__t5.IntVal) != (0) {
__t4 = gopurs_runtime.Apply2(Functor0_3_2.V0, pkg_Data_Maybe.Get_Just(), gen_6)
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Apply(Applicative0_4_3.V1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})
}
end_branch_4:
return __t4
}))
})
})
}

func Call_genMaybe_prime__gopurs_runtime_Value_1561363431(dictMonadGen_0_loop *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadGen_0 *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value] = dictMonadGen_0_loop
_ = dictMonadGen_0
Monad0_1_0 := gopurs_runtime.Apply(dictMonadGen_0.V0, gopurs_runtime.Value{})
_ = Monad0_1_0
Bind1_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
Functor0_3_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_2
Applicative0_4_3 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_3
return gopurs_runtime.Func(func(bias_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(gen_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply2(dictMonadGen_0.V2, gopurs_runtime.Float(0.0), gopurs_runtime.Float(1.0)), gopurs_runtime.Func(func(n_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
var __t5 gopurs_runtime.Value
{
if (n_7.FloatVal()) < (bias_5.FloatVal()) {
__t5 = gopurs_runtime.Bool(true)
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Bool(false)
}
end_branch_5:
if (__t5.IntVal) != (0) {
__t4 = gopurs_runtime.Apply2(Functor0_3_2.V0, pkg_Data_Maybe.Get_Just(), gen_6)
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Apply(Applicative0_4_3.V1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})
}
end_branch_4:
return __t4
}))
})
})
}

func Call_genMaybe(dictMonadGen_0_loop *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadGen_0 *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value] = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(Get_genMaybe_prime(), gopurs_runtime.Value{Type: 9, IntVal: 2254593219, UnsafePtr: unsafe.Pointer(dictMonadGen_0)}, gopurs_runtime.Float(0.75))
}

func Call_genIdentity(dictFunctor_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFunctor_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.Apply(dictFunctor_0.V0, pkg_Data_Identity.Get_Identity())
}

func Call_genEither_prime(dictMonadGen_0_loop *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadGen_0 *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value] = dictMonadGen_0_loop
_ = dictMonadGen_0
Monad0_1_0 := gopurs_runtime.Apply(dictMonadGen_0.V0, gopurs_runtime.Value{})
_ = Monad0_1_0
Bind1_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
Functor0_3_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_2
return gopurs_runtime.Func(func(bias_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(genA_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(genB_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply2(dictMonadGen_0.V2, gopurs_runtime.Float(0.0), gopurs_runtime.Float(1.0)), gopurs_runtime.Func(func(n_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
var __t4 gopurs_runtime.Value
{
if (n_7.FloatVal()) < (bias_4.FloatVal()) {
__t4 = gopurs_runtime.Bool(true)
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Bool(false)
}
end_branch_4:
if (__t4.IntVal) != (0) {
__t3 = gopurs_runtime.Apply2(Functor0_3_2.V0, pkg_Data_Either.Get_Left(), genA_5)
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Apply2(Functor0_3_2.V0, pkg_Data_Either.Get_Right(), genB_6)
}
end_branch_3:
return __t3
}))
})
})
})
}

func Call_genEither_prime__gopurs_runtime_Value_1946557461(dictMonadGen_0_loop *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadGen_0 *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value] = dictMonadGen_0_loop
_ = dictMonadGen_0
Monad0_1_0 := gopurs_runtime.Apply(dictMonadGen_0.V0, gopurs_runtime.Value{})
_ = Monad0_1_0
Bind1_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
Functor0_3_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_2
return gopurs_runtime.Func(func(bias_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(genA_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(genB_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply2(dictMonadGen_0.V2, gopurs_runtime.Float(0.0), gopurs_runtime.Float(1.0)), gopurs_runtime.Func(func(n_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
var __t4 gopurs_runtime.Value
{
if (n_7.FloatVal()) < (bias_4.FloatVal()) {
__t4 = gopurs_runtime.Bool(true)
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Bool(false)
}
end_branch_4:
if (__t4.IntVal) != (0) {
__t3 = gopurs_runtime.Apply2(Functor0_3_2.V0, pkg_Data_Either.Get_Left(), genA_5)
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Apply2(Functor0_3_2.V0, pkg_Data_Either.Get_Right(), genB_6)
}
end_branch_3:
return __t3
}))
})
})
})
}

func Call_genEither(dictMonadGen_0_loop *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadGen_0 *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value] = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(Get_genEither_prime(), gopurs_runtime.Value{Type: 9, IntVal: 2254593219, UnsafePtr: unsafe.Pointer(dictMonadGen_0)}, gopurs_runtime.Float(0.5))
}


