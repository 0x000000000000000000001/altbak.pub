package Control_Monad_Gen_Common

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Identity "gopurs/output/Data.Identity"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Control_Monad_Gen "gopurs/output/Control.Monad.Gen"
	pkg_Data_NonEmpty "gopurs/output/Data.NonEmpty"
)

var max gopurs_runtime.Value
var once_max sync.Once
func Get_max() gopurs_runtime.Value {
	once_max.Do(func() {
		max = gopurs_runtime.Func2(Call_max)
	})
	return max
}

var genTuple gopurs_runtime.Value
var once_genTuple sync.Once
func Get_genTuple() gopurs_runtime.Value {
	once_genTuple.Do(func() {
		genTuple = gopurs_runtime.Func3(Call_genTuple)
	})
	return genTuple
}

var genNonEmpty gopurs_runtime.Value
var once_genNonEmpty sync.Once
func Get_genNonEmpty() gopurs_runtime.Value {
	once_genNonEmpty.Do(func() {
		genNonEmpty = gopurs_runtime.Func2(Call_genNonEmpty)
	})
	return genNonEmpty
}

var genMaybe_prime gopurs_runtime.Value
var once_genMaybe_prime sync.Once
func Get_genMaybe_prime() gopurs_runtime.Value {
	once_genMaybe_prime.Do(func() {
		genMaybe_prime = gopurs_runtime.Func(func(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0_loop, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
Bind1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{})
_ = Bind1_2_1
return gopurs_runtime.Func2(func(bias_3 gopurs_runtime.Value, gen_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_2_1, "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0_loop, "chooseFloat"), gopurs_runtime.Float(0.0), gopurs_runtime.Float(1.0)), gopurs_runtime.Func(func(n_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if n_5.FloatVal() < bias_3.FloatVal() {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_2_1, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Maybe.Get_Just(), gen_4)
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Constructor0("Nothing"))
}
end_branch_2:
return __t2
}))
})
}()
})
	})
	return genMaybe_prime
}

var genMaybe gopurs_runtime.Value
var once_genMaybe sync.Once
func Get_genMaybe() gopurs_runtime.Value {
	once_genMaybe.Do(func() {
		genMaybe = gopurs_runtime.Func(func(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(Get_genMaybe_prime(), dictMonadGen_0_loop, gopurs_runtime.Float(0.75))
}()
})
	})
	return genMaybe
}

var genIdentity gopurs_runtime.Value
var once_genIdentity sync.Once
func Get_genIdentity() gopurs_runtime.Value {
	once_genIdentity.Do(func() {
		genIdentity = gopurs_runtime.Func(func(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0_loop, "map"), pkg_Data_Identity.Get_Identity())
}()
})
	})
	return genIdentity
}

var genEither_prime gopurs_runtime.Value
var once_genEither_prime sync.Once
func Get_genEither_prime() gopurs_runtime.Value {
	once_genEither_prime.Do(func() {
		genEither_prime = gopurs_runtime.Func(func(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
Bind1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0_loop, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{})
_ = Bind1_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_1_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func3(func(bias_3 gopurs_runtime.Value, genA_4 gopurs_runtime.Value, genB_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_1_0, "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0_loop, "chooseFloat"), gopurs_runtime.Float(0.0), gopurs_runtime.Float(1.0)), gopurs_runtime.Func(func(n_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if n_6.FloatVal() < bias_3.FloatVal() {
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
}()
})
	})
	return genEither_prime
}

var genEither gopurs_runtime.Value
var once_genEither sync.Once
func Get_genEither() gopurs_runtime.Value {
	once_genEither.Do(func() {
		genEither = gopurs_runtime.Func(func(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(Get_genEither_prime(), dictMonadGen_0_loop, gopurs_runtime.Float(0.5))
}()
})
	})
	return genEither
}

func Call_max(x_0_loop gopurs_runtime.Value, y_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var y_1 gopurs_runtime.Value = y_1_loop
_ = y_1
v_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ord.Get_ordInt(), "compare"), x_0_loop, y_1_loop)
_ = v_2_0
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_2_0.StrVal == "LT").IntVal != 0 {
__t1 = y_1_loop
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v_2_0.StrVal == "EQ").IntVal != 0 {
__t1 = x_0_loop
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v_2_0.StrVal == "GT").IntVal != 0 {
__t1 = x_0_loop
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}

func Call_genTuple(dictApply_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0_loop, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0_loop, "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Tuple.Get_Tuple(), a_1_loop), b_2_loop)
}

func Call_genNonEmpty(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
_ = dictMonadGen_1
Apply0_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1_loop, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{})
_ = Apply0_2_0
unfoldable1_3_1 := gopurs_runtime.Apply2(pkg_Control_Monad_Gen.Get_unfoldable(), dictMonadRec_0_loop, dictMonadGen_1_loop)
_ = unfoldable1_3_1
return gopurs_runtime.Func(func(dictUnfoldable_4 gopurs_runtime.Value) gopurs_runtime.Value {
unfoldable2_5_2 := gopurs_runtime.Apply(unfoldable1_3_1, dictUnfoldable_4)
_ = unfoldable2_5_2
return gopurs_runtime.Func(func(gen_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_2_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_2_0, "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_NonEmpty.Get_NonEmpty(), gen_6), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_1_loop, "resize"), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_max(), gopurs_runtime.Int(0), gopurs_runtime.Int(x_7.IntVal - 1))
}), gopurs_runtime.Apply(unfoldable2_5_2, gen_6)))
})
})
}


