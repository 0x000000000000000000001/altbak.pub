package Data_Decidable

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Decide "gopurs/output/Data.Decide"
	pkg_Data_Divisible "gopurs/output/Data.Divisible"
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
return Call_lose(dict_0_box)
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
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value) gopurs_runtime.Value {
var spin_2_0 gopurs_runtime.Value
spin_2_0 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_3_loop gopurs_runtime.Value = v_3_loop_val
spin_2_0:
for {
if false { continue spin_2_0 }
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
v_3_loop = v_3
continue spin_2_0
return gopurs_runtime.Value{}
}
}()
})
return gopurs_runtime.Apply(spin_2_0, gopurs_runtime.Apply(f_0, a_1))
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
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value) gopurs_runtime.Value {
var spin_2_0 gopurs_runtime.Value
spin_2_0 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_3_loop gopurs_runtime.Value = v_3_loop_val
spin_2_0:
for {
if false { continue spin_2_0 }
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
v_3_loop = v_3
continue spin_2_0
return gopurs_runtime.Value{}
}
}()
})
return gopurs_runtime.Apply(spin_2_0, gopurs_runtime.Apply(f_0, a_1))
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
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var spin_3_0 gopurs_runtime.Value
spin_3_0 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
spin_3_0:
for {
if false { continue spin_3_0 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
v_4_loop = v_4
continue spin_3_0
return gopurs_runtime.Value{}
}
}()
})
return gopurs_runtime.Apply(spin_3_0, gopurs_runtime.Apply(f_0, a_1))
}))
	})
	return cache_decidableComparison
}

func Call_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_lose(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return ((*gopurs_runtime.RecordData1)(dict_0.UnsafePtr)).V0
}

func Call_lost(dictDecidable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictDecidable_0 gopurs_runtime.Value = dictDecidable_0_loop
_ = dictDecidable_0
return gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictDecidable_0.UnsafePtr)).V0, Get_identity())
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
}), gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value) gopurs_runtime.Value {
var spin_5_2 gopurs_runtime.Value
spin_5_2 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_6_loop gopurs_runtime.Value = v_6_loop_val
spin_5_2:
for {
if false { continue spin_5_2 }
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
v_6_loop = v_6
continue spin_5_2
return gopurs_runtime.Value{}
}
}()
})
return gopurs_runtime.Apply(spin_5_2, gopurs_runtime.Apply(f_3, a_4))
}))
}


