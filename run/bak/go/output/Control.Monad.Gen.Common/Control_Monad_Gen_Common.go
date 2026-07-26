package Control_Monad_Gen_Common

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Control_Monad_Gen "gopurs/output/Control.Monad.Gen"
	pkg_Data_NonEmpty "gopurs/output/Data.NonEmpty"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Identity "gopurs/output/Data.Identity"
	pkg_Data_Either "gopurs/output/Data.Either"
)

var cache_max gopurs_runtime.Value
var once_max sync.Once
func Get_max() gopurs_runtime.Value {
	once_max.Do(func() {
		cache_max = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})
_ = __local_var_0_0
return gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_1 := gopurs_runtime.Apply2(__local_var_0_0, x_1, y_2)
_ = v_3_1
var __t2 gopurs_runtime.Value
{
if (v_3_1.Type == 9 && v_3_1.IntVal == 1527465420) {
__t2 = y_2
goto end_branch_2
} else {

}
}
{
if (v_3_1.Type == 9 && v_3_1.IntVal == 902936544) {
__t2 = x_1
goto end_branch_2
} else {

}
}
{
if (v_3_1.Type == 9 && v_3_1.IntVal == 380165415) {
__t2 = x_1
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
}()
	})
	return cache_max
}

var cache_genTuple gopurs_runtime.Value
var once_genTuple sync.Once
func Get_genTuple() gopurs_runtime.Value {
	once_genTuple.Do(func() {
		cache_genTuple = gopurs_runtime.Func3(func(dictApply_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genTuple(dictApply_0_box, a_1_box, b_2_box)
})
	})
	return cache_genTuple
}

var cache_genNonEmpty gopurs_runtime.Value
var once_genNonEmpty sync.Once
func Get_genNonEmpty() gopurs_runtime.Value {
	once_genNonEmpty.Do(func() {
		cache_genNonEmpty = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genNonEmpty(dictMonadRec_0_box, dictMonadGen_1_box)
})
	})
	return cache_genNonEmpty
}

var cache_genMaybe_prime gopurs_runtime.Value
var once_genMaybe_prime sync.Once
func Get_genMaybe_prime() gopurs_runtime.Value {
	once_genMaybe_prime.Do(func() {
		cache_genMaybe_prime = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genMaybe_prime(dictMonadGen_0_box)
})
	})
	return cache_genMaybe_prime
}

var cache_genMaybe_prime__gopurs_runtime_Value_3295579624 gopurs_runtime.Value
var once_genMaybe_prime__gopurs_runtime_Value_3295579624 sync.Once
func Get_genMaybe_prime__gopurs_runtime_Value_3295579624() gopurs_runtime.Value {
	once_genMaybe_prime__gopurs_runtime_Value_3295579624.Do(func() {
		cache_genMaybe_prime__gopurs_runtime_Value_3295579624 = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genMaybe_prime__gopurs_runtime_Value_3295579624(dictMonadGen_0_box)
})
	})
	return cache_genMaybe_prime__gopurs_runtime_Value_3295579624
}

var cache_genMaybe gopurs_runtime.Value
var once_genMaybe sync.Once
func Get_genMaybe() gopurs_runtime.Value {
	once_genMaybe.Do(func() {
		cache_genMaybe = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genMaybe(dictMonadGen_0_box)
})
	})
	return cache_genMaybe
}

var cache_genIdentity gopurs_runtime.Value
var once_genIdentity sync.Once
func Get_genIdentity() gopurs_runtime.Value {
	once_genIdentity.Do(func() {
		cache_genIdentity = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genIdentity(dictFunctor_0_box)
})
	})
	return cache_genIdentity
}

var cache_genEither_prime gopurs_runtime.Value
var once_genEither_prime sync.Once
func Get_genEither_prime() gopurs_runtime.Value {
	once_genEither_prime.Do(func() {
		cache_genEither_prime = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genEither_prime(dictMonadGen_0_box)
})
	})
	return cache_genEither_prime
}

var cache_genEither_prime__gopurs_runtime_Value_4139285216 gopurs_runtime.Value
var once_genEither_prime__gopurs_runtime_Value_4139285216 sync.Once
func Get_genEither_prime__gopurs_runtime_Value_4139285216() gopurs_runtime.Value {
	once_genEither_prime__gopurs_runtime_Value_4139285216.Do(func() {
		cache_genEither_prime__gopurs_runtime_Value_4139285216 = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genEither_prime__gopurs_runtime_Value_4139285216(dictMonadGen_0_box)
})
	})
	return cache_genEither_prime__gopurs_runtime_Value_4139285216
}

var cache_genEither gopurs_runtime.Value
var once_genEither sync.Once
func Get_genEither() gopurs_runtime.Value {
	once_genEither.Do(func() {
		cache_genEither = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genEither(dictMonadGen_0_box)
})
	})
	return cache_genEither
}

func Call_genTuple(dictApply_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Tuple.Get_Tuple(), a_1), b_2)
}

func Call_genNonEmpty(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
_ = dictMonadGen_1
Apply0_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{})
_ = Apply0_2_0
unfoldable1_3_1 := gopurs_runtime.Apply2(pkg_Control_Monad_Gen.Get_unfoldable(), dictMonadRec_0, dictMonadGen_1)
_ = unfoldable1_3_1
return gopurs_runtime.Func(func(dictUnfoldable_4 gopurs_runtime.Value) gopurs_runtime.Value {
unfoldable2_5_2 := gopurs_runtime.Apply(unfoldable1_3_1, dictUnfoldable_4)
_ = unfoldable2_5_2
return gopurs_runtime.Func(func(gen_6 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_3 := gopurs_runtime.Apply(Get_max(), gopurs_runtime.Int(0))
_ = __local_var_7_3
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_2_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_2_0, "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_NonEmpty.Get_NonEmpty(), gen_6), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_1, "resize"), gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_3, gopurs_runtime.Int((x_8.IntVal) - (1)))
}), gopurs_runtime.Apply(unfoldable2_5_2, gen_6)))
})
})
}

func Call_genMaybe_prime(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
Bind1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{})
_ = Bind1_2_1
return gopurs_runtime.Func2(func(bias_3 gopurs_runtime.Value, gen_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_2_1, "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseFloat"), gopurs_runtime.Float(0.0), gopurs_runtime.Float(1.0)), gopurs_runtime.Func(func(n_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (n_5.FloatVal()) < (bias_3.FloatVal()) {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_2_1, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Maybe.Get_Just(), gen_4)
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})
}
end_branch_2:
return __t2
}))
})
}

func Call_genMaybe_prime__gopurs_runtime_Value_3295579624(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
Bind1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{})
_ = Bind1_2_1
return gopurs_runtime.Func2(func(bias_3 gopurs_runtime.Value, gen_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_2_1, "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseFloat"), gopurs_runtime.Float(0.0), gopurs_runtime.Float(1.0)), gopurs_runtime.Func(func(n_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (n_5.FloatVal()) < (bias_3.FloatVal()) {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_2_1, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Maybe.Get_Just(), gen_4)
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})
}
end_branch_2:
return __t2
}))
})
}

func Call_genMaybe(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(Get_genMaybe_prime(), dictMonadGen_0, gopurs_runtime.Float(0.75))
}

func Call_genIdentity(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), pkg_Data_Identity.Get_Identity())
}

func Call_genEither_prime(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
Bind1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{})
_ = Bind1_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_1_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func3(func(bias_3 gopurs_runtime.Value, genA_4 gopurs_runtime.Value, genB_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_1_0, "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseFloat"), gopurs_runtime.Float(0.0), gopurs_runtime.Float(1.0)), gopurs_runtime.Func(func(n_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (n_6.FloatVal()) < (bias_3.FloatVal()) {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "map"), pkg_Data_Either.Get_Left(), genA_4)
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "map"), pkg_Data_Either.Get_Right(), genB_5)
}
end_branch_2:
return __t2
}))
})
}

func Call_genEither_prime__gopurs_runtime_Value_4139285216(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
Bind1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{})
_ = Bind1_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_1_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func3(func(bias_3 gopurs_runtime.Value, genA_4 gopurs_runtime.Value, genB_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_1_0, "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseFloat"), gopurs_runtime.Float(0.0), gopurs_runtime.Float(1.0)), gopurs_runtime.Func(func(n_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (n_6.FloatVal()) < (bias_3.FloatVal()) {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "map"), pkg_Data_Either.Get_Left(), genA_4)
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "map"), pkg_Data_Either.Get_Right(), genB_5)
}
end_branch_2:
return __t2
}))
})
}

func Call_genEither(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(Get_genEither_prime(), dictMonadGen_0, gopurs_runtime.Float(0.5))
}


