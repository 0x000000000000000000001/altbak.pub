package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Data_Equivalence_Equivalence gopurs_runtime.Value
var once_Data_Equivalence_Equivalence sync.Once
func Get_Data_Equivalence_Equivalence() gopurs_runtime.Value {
	once_Data_Equivalence_Equivalence.Do(func() {
		cache_Data_Equivalence_Equivalence = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Equivalence_Equivalence(x_0_box)
})
	})
	return cache_Data_Equivalence_Equivalence
}

var cache_Data_Equivalence_semigroupEquivalence gopurs_runtime.Value
var once_Data_Equivalence_semigroupEquivalence sync.Once
func Get_Data_Equivalence_semigroupEquivalence() gopurs_runtime.Value {
	once_Data_Equivalence_semigroupEquivalence.Do(func() {
		cache_Data_Equivalence_semigroupEquivalence = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(v_0, a_2, b_3).IntVal) != (0)) && ((gopurs_runtime.Apply2(v1_1, a_2, b_3).IntVal) != (0)))
})
})
})
}))
	})
	return cache_Data_Equivalence_semigroupEquivalence
}

var cache_Data_Equivalence_newtypeEquivalence gopurs_runtime.Value
var once_Data_Equivalence_newtypeEquivalence sync.Once
func Get_Data_Equivalence_newtypeEquivalence() gopurs_runtime.Value {
	once_Data_Equivalence_newtypeEquivalence.Do(func() {
		cache_Data_Equivalence_newtypeEquivalence = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_Data_Equivalence_newtypeEquivalence
}

var cache_Data_Equivalence_monoidEquivalence gopurs_runtime.Value
var once_Data_Equivalence_monoidEquivalence sync.Once
func Get_Data_Equivalence_monoidEquivalence() gopurs_runtime.Value {
	once_Data_Equivalence_monoidEquivalence.Do(func() {
		cache_Data_Equivalence_monoidEquivalence = gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Equivalence_semigroupEquivalence()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})
}))
	})
	return cache_Data_Equivalence_monoidEquivalence
}

var cache_Data_Equivalence_defaultEquivalence gopurs_runtime.Value
var once_Data_Equivalence_defaultEquivalence sync.Once
func Get_Data_Equivalence_defaultEquivalence() gopurs_runtime.Value {
	once_Data_Equivalence_defaultEquivalence.Do(func() {
		cache_Data_Equivalence_defaultEquivalence = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Equivalence_defaultEquivalence(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
})
	})
	return cache_Data_Equivalence_defaultEquivalence
}

var cache_Data_Equivalence_contravariantEquivalence gopurs_runtime.Value
var once_Data_Equivalence_contravariantEquivalence sync.Once
func Get_Data_Equivalence_contravariantEquivalence() gopurs_runtime.Value {
	once_Data_Equivalence_contravariantEquivalence.Do(func() {
		cache_Data_Equivalence_contravariantEquivalence = gopurs_runtime.RecordDict1("cmap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(v_1, gopurs_runtime.Apply(f_0, x_2), gopurs_runtime.Apply(f_0, y_3))
})
})
})
}))
	})
	return cache_Data_Equivalence_contravariantEquivalence
}

var cache_Data_Equivalence_comparisonEquivalence gopurs_runtime.Value
var once_Data_Equivalence_comparisonEquivalence sync.Once
func Get_Data_Equivalence_comparisonEquivalence() gopurs_runtime.Value {
	once_Data_Equivalence_comparisonEquivalence.Do(func() {
		cache_Data_Equivalence_comparisonEquivalence = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Equivalence_comparisonEquivalence(v_0_box, a_1_box, b_2_box))
})
	})
	return cache_Data_Equivalence_comparisonEquivalence
}

var cache_Data_Equivalence_contravariantEquivalence__506233683 gopurs_runtime.Value
var once_Data_Equivalence_contravariantEquivalence__506233683 sync.Once
func Get_Data_Equivalence_contravariantEquivalence__506233683() gopurs_runtime.Value {
	once_Data_Equivalence_contravariantEquivalence__506233683.Do(func() {
		cache_Data_Equivalence_contravariantEquivalence__506233683 = gopurs_runtime.RecordDict1("cmap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(v_1, gopurs_runtime.Apply(f_0, x_2), gopurs_runtime.Apply(f_0, y_3))
})
})
})
}))
	})
	return cache_Data_Equivalence_contravariantEquivalence__506233683
}

var cache_Data_Equivalence_semigroupEquivalence__2325462015 gopurs_runtime.Value
var once_Data_Equivalence_semigroupEquivalence__2325462015 sync.Once
func Get_Data_Equivalence_semigroupEquivalence__2325462015() gopurs_runtime.Value {
	once_Data_Equivalence_semigroupEquivalence__2325462015.Do(func() {
		cache_Data_Equivalence_semigroupEquivalence__2325462015 = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(v_0, a_2, b_3).IntVal) != (0)) && ((gopurs_runtime.Apply2(v1_1, a_2, b_3).IntVal) != (0)))
})
})
})
}))
	})
	return cache_Data_Equivalence_semigroupEquivalence__2325462015
}

func Call_Data_Equivalence_Equivalence(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Equivalence_defaultEquivalence(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Box(dictEq_0.V0)
}

func Call_Data_Equivalence_comparisonEquivalence(v_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value) bool {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
// TAST (Let): __local_var_3_0 -> uint32
__local_var_3_0 := uint32(gopurs_runtime.Apply2(v_0, a_1, b_2).IntVal)
_ = __local_var_3_0
var __t1 bool
{
if (__local_var_3_0 == 1527465420) {
__t1 = false
goto end_branch_1
} else {

}
}
{
if (__local_var_3_0 == 380165415) {
__t1 = false
goto end_branch_1
} else {

}
}
{
if (__local_var_3_0 == 902936544) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
return __t1
}


