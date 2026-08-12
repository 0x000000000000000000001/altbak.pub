package Data_Decidable

import (
	pkg_Data_Decide "gopurs/output/Data.Decide"
	pkg_Data_Divisible "gopurs/output/Data.Divisible"
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

var cache_lose__gopurs_runtime_Value_3306256519 gopurs_runtime.Value
var once_lose__gopurs_runtime_Value_3306256519 sync.Once
func Get_lose__gopurs_runtime_Value_3306256519() gopurs_runtime.Value {
	once_lose__gopurs_runtime_Value_3306256519.Do(func() {
		cache_lose__gopurs_runtime_Value_3306256519 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lose__gopurs_runtime_Value_3306256519(gopurs_runtime.CoerceToStruct[Constructor_Decidable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_lose__gopurs_runtime_Value_3306256519
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
var spin_2_0_0 gopurs_runtime.Value
spin_2_0_0 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_3_loop gopurs_runtime.Value = v_3_loop_val
spin_2_0_0:
for {
if false { continue spin_2_0_0 }
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
v_3_loop = v_3
continue spin_2_0_0
return gopurs_runtime.Value{}
}
}()
})
return gopurs_runtime.Bool((gopurs_runtime.Apply(spin_2_0_0, gopurs_runtime.Apply(f_0, a_1)).IntVal) != (0))
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
var spin_2_0_2 gopurs_runtime.Value
spin_2_0_2 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_3_loop gopurs_runtime.Value = v_3_loop_val
spin_2_0_2:
for {
if false { continue spin_2_0_2 }
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
v_3_loop = v_3
continue spin_2_0_2
return gopurs_runtime.Value{}
}
}()
})
return gopurs_runtime.Apply(spin_2_0_2, gopurs_runtime.Apply(f_0, a_1))
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
var spin_3_0_3 gopurs_runtime.Value
spin_3_0_3 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
spin_3_0_3:
for {
if false { continue spin_3_0_3 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
v_4_loop = v_4
continue spin_3_0_3
return gopurs_runtime.Value{}
}
}()
})
return gopurs_runtime.Apply(spin_3_0_3, gopurs_runtime.Apply(f_0, a_1))
})
})
}))
	})
	return cache_decidableComparison
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

func Call_lose__gopurs_runtime_Value_3306256519(dict_0_loop *Constructor_Decidable[gopurs_runtime.Value]) gopurs_runtime.Value {
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
var spin_5_2_1 gopurs_runtime.Value
spin_5_2_1 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_6_loop gopurs_runtime.Value = v_6_loop_val
spin_5_2_1:
for {
if false { continue spin_5_2_1 }
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
v_6_loop = v_6
continue spin_5_2_1
return gopurs_runtime.Value{}
}
}()
})
return gopurs_runtime.Apply(spin_5_2_1, gopurs_runtime.Apply(f_3, a_4))
})
}))
}


