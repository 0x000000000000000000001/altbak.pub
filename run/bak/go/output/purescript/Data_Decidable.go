package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Decidable_identity gopurs_runtime.Value
var once_Data_Decidable_identity sync.Once
func Get_Data_Decidable_identity() gopurs_runtime.Value {
	once_Data_Decidable_identity.Do(func() {
		cache_Data_Decidable_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Decidable_identity(x_0_box)
})
	})
	return cache_Data_Decidable_identity
}

var cache_Data_Decidable_Decidable_dollarDict gopurs_runtime.Value
var once_Data_Decidable_Decidable_dollarDict sync.Once
func Get_Data_Decidable_Decidable_dollarDict() gopurs_runtime.Value {
	once_Data_Decidable_Decidable_dollarDict.Do(func() {
		cache_Data_Decidable_Decidable_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Decidable_Decidable_dollarDict(x_0_box)
})
	})
	return cache_Data_Decidable_Decidable_dollarDict
}

var cache_Data_Decidable_lose gopurs_runtime.Value
var once_Data_Decidable_lose sync.Once
func Get_Data_Decidable_lose() gopurs_runtime.Value {
	once_Data_Decidable_lose.Do(func() {
		cache_Data_Decidable_lose = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Decidable_lose(gopurs_runtime.CoerceToStruct[Constructor_Data_Decidable_Decidable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Decidable_lose
}

var cache_Data_Decidable_lost gopurs_runtime.Value
var once_Data_Decidable_lost sync.Once
func Get_Data_Decidable_lost() gopurs_runtime.Value {
	once_Data_Decidable_lost.Do(func() {
		cache_Data_Decidable_lost = gopurs_runtime.Func(func(dictDecidable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Decidable_lost(dictDecidable_0_box)
})
	})
	return cache_Data_Decidable_lost
}

var cache_Data_Decidable_decidablePredicate gopurs_runtime.Value
var once_Data_Decidable_decidablePredicate sync.Once
func Get_Data_Decidable_decidablePredicate() gopurs_runtime.Value {
	once_Data_Decidable_decidablePredicate.Do(func() {
		cache_Data_Decidable_decidablePredicate = gopurs_runtime.RecordDict3("Decide0", "Divisible1", "lose", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Decide_choosePredicate()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Divisible_divisiblePredicate()
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
return gopurs_runtime.Bool((gopurs_runtime.Value{}.IntVal) != (0))
}
}()
})
return gopurs_runtime.Bool((gopurs_runtime.Apply(spin_2_0_0, gopurs_runtime.Apply(f_0, a_1)).IntVal) != (0))
})
}))
	})
	return cache_Data_Decidable_decidablePredicate
}

var cache_Data_Decidable_decidableOp gopurs_runtime.Value
var once_Data_Decidable_decidableOp sync.Once
func Get_Data_Decidable_decidableOp() gopurs_runtime.Value {
	once_Data_Decidable_decidableOp.Do(func() {
		cache_Data_Decidable_decidableOp = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Decidable_decidableOp(dictMonoid_0_box)
})
	})
	return cache_Data_Decidable_decidableOp
}

var cache_Data_Decidable_decidableEquivalence gopurs_runtime.Value
var once_Data_Decidable_decidableEquivalence sync.Once
func Get_Data_Decidable_decidableEquivalence() gopurs_runtime.Value {
	once_Data_Decidable_decidableEquivalence.Do(func() {
		cache_Data_Decidable_decidableEquivalence = gopurs_runtime.RecordDict3("Decide0", "Divisible1", "lose", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Decide_chooseEquivalence()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Divisible_divisibleEquivalence()
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
	return cache_Data_Decidable_decidableEquivalence
}

var cache_Data_Decidable_decidableComparison gopurs_runtime.Value
var once_Data_Decidable_decidableComparison sync.Once
func Get_Data_Decidable_decidableComparison() gopurs_runtime.Value {
	once_Data_Decidable_decidableComparison.Do(func() {
		cache_Data_Decidable_decidableComparison = gopurs_runtime.RecordDict3("Decide0", "Divisible1", "lose", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Decide_chooseComparison()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Divisible_divisibleComparison()
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
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{}.IntVal)), UnsafePtr: nil}
}
}()
})
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply(spin_3_0_3, gopurs_runtime.Apply(f_0, a_1)).IntVal)), UnsafePtr: nil}
})
})
}))
	})
	return cache_Data_Decidable_decidableComparison
}

var cache_Data_Decidable_lose__3306256519 gopurs_runtime.Value
var once_Data_Decidable_lose__3306256519 sync.Once
func Get_Data_Decidable_lose__3306256519() gopurs_runtime.Value {
	once_Data_Decidable_lose__3306256519.Do(func() {
		cache_Data_Decidable_lose__3306256519 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Decidable_lose__3306256519(gopurs_runtime.CoerceToStruct[Constructor_Data_Decidable_Decidable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Decidable_lose__3306256519
}

var cache_Data_Decidable_lose__2926568423 gopurs_runtime.Value
var once_Data_Decidable_lose__2926568423 sync.Once
func Get_Data_Decidable_lose__2926568423() gopurs_runtime.Value {
	once_Data_Decidable_lose__2926568423.Do(func() {
		cache_Data_Decidable_lose__2926568423 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Decidable_lose__2926568423(gopurs_runtime.CoerceToStruct[Constructor_Data_Decidable_Decidable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Decidable_lose__2926568423
}

type Constructor_Data_Decidable_Decidable[T_f any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[336732346] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Decidable_Decidable[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Decide0": return gopurs_runtime.Box(c.V0)
		case "Divisible1": return gopurs_runtime.Box(c.V1)
		case "lose": return gopurs_runtime.Box(c.V2)
		default: panic("Key not found in dictionary Constructor_Data_Decidable_Decidable: " + key)
		}
	}
}


func Call_Data_Decidable_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Decidable_Decidable_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Decidable_lose(dict_0_loop *Constructor_Data_Decidable_Decidable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Decidable_Decidable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Decidable_lost(dictDecidable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictDecidable_0 gopurs_runtime.Value = dictDecidable_0_loop
_ = dictDecidable_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDecidable_0, "lose"), Get_Data_Decidable_identity())
}

func Call_Data_Decidable_decidableOp(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): chooseOp_1_0 -> gopurs_runtime.Value
chooseOp_1_0 := gopurs_runtime.Apply(Get_Data_Decide_chooseOp(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = chooseOp_1_0
// TAST (Let): divisibleOp_2_1 -> gopurs_runtime.Value
divisibleOp_2_1 := gopurs_runtime.Apply(Get_Data_Divisible_divisibleOp(), dictMonoid_0)
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

func Call_Data_Decidable_lose__3306256519(dict_0_loop *Constructor_Data_Decidable_Decidable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Decidable_Decidable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Decidable_lose__2926568423(dict_0_loop *Constructor_Data_Decidable_Decidable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Decidable_Decidable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}


