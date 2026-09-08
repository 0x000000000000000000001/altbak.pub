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
		cache_Data_Decidable_identity = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1)
	})
	return cache_Data_Decidable_identity
}

var cache_Data_Decidable_Decidable_dollar_Dict gopurs_runtime.Value
var once_Data_Decidable_Decidable_dollar_Dict sync.Once
func Get_Data_Decidable_Decidable_dollar_Dict() gopurs_runtime.Value {
	once_Data_Decidable_Decidable_dollar_Dict.Do(func() {
		cache_Data_Decidable_Decidable_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 336732346, UnsafePtr: unsafe.Pointer(Call_Data_Decidable_Decidable_dollar_Dict(func() struct{
	Decide0 gopurs_runtime.Value
	Divisible1 gopurs_runtime.Value
	lose gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Decide0 gopurs_runtime.Value
	Divisible1 gopurs_runtime.Value
	lose gopurs_runtime.Value
}{}
					clone.Decide0 = gopurs_runtime.RecordGet(orig, "Decide0")
					clone.Divisible1 = gopurs_runtime.RecordGet(orig, "Divisible1")
					clone.lose = gopurs_runtime.RecordGet(orig, "lose")
					return clone
				}()))}
})
	})
	return cache_Data_Decidable_Decidable_dollar_Dict
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
		cache_Data_Decidable_decidablePredicate = gopurs_runtime.Value{Type: 9, IntVal: 336732346, UnsafePtr: unsafe.Pointer((&Constructor_Data_Decidable_Decidable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1618621146, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Decide_Decide[gopurs_runtime.Value]](Get_Data_Decide_choosePredicate()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2265116602, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Divisible_Divisible[gopurs_runtime.Value]](Get_Data_Divisible_divisiblePredicate()))}
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_Decidable_spin__2182866177_2_0_0 func(gopurs_runtime.Value) bool
_ = Call_local_Data_Decidable_spin__2182866177_2_0_0
var spin__2182866177_2_0_0 gopurs_runtime.Value
_ = spin__2182866177_2_0_0
Call_local_Data_Decidable_spin__2182866177_2_0_0 = func(v_3_loop gopurs_runtime.Value) bool {
spin__2182866177_2_0_0:
for {
if false { continue spin__2182866177_2_0_0 }
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
v_3_loop = v_3
continue spin__2182866177_2_0_0
return func() bool { panic("unreachable") }()
}
}
spin__2182866177_2_0_0 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_local_Data_Decidable_spin__2182866177_2_0_0(v_3_loop_val))
})
var spin_3_1_1 gopurs_runtime.Value
_ = spin_3_1_1
// FALLBACK TCO: isLoop=false len=1
spin_3_1_1 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_local_Data_Decidable_spin__2182866177_2_0_0(v_4))
})
return gopurs_runtime.Bool(Call_local_Data_Decidable_spin__2182866177_2_0_0(gopurs_runtime.Apply(f_0, a_1)))
})}))}
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
		cache_Data_Decidable_decidableEquivalence = gopurs_runtime.Value{Type: 9, IntVal: 336732346, UnsafePtr: unsafe.Pointer((&Constructor_Data_Decidable_Decidable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1618621146, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Decide_Decide[gopurs_runtime.Value]](Get_Data_Decide_chooseEquivalence()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2265116602, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Divisible_Divisible[gopurs_runtime.Value]](Get_Data_Divisible_divisibleEquivalence()))}
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_Decidable_spin__1769020947_2_0_4 func(gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_Decidable_spin__1769020947_2_0_4
var spin__1769020947_2_0_4 gopurs_runtime.Value
_ = spin__1769020947_2_0_4
Call_local_Data_Decidable_spin__1769020947_2_0_4 = func(v_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
spin__1769020947_2_0_4:
for {
if false { continue spin__1769020947_2_0_4 }
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
v_3_loop = v_3
continue spin__1769020947_2_0_4
return func() gopurs_runtime.Value { panic("unreachable") }()
}
}
spin__1769020947_2_0_4 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_Decidable_spin__1769020947_2_0_4(v_3_loop_val)
})
var spin_3_1_5 gopurs_runtime.Value
_ = spin_3_1_5
// FALLBACK TCO: isLoop=false len=1
spin_3_1_5 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_Decidable_spin__1769020947_2_0_4(v_4)
})
return gopurs_runtime.Bool((Call_local_Data_Decidable_spin__1769020947_2_0_4(gopurs_runtime.Apply(f_0, a_1)).IntVal) != (0))
})}))}
	})
	return cache_Data_Decidable_decidableEquivalence
}

var cache_Data_Decidable_decidableComparison gopurs_runtime.Value
var once_Data_Decidable_decidableComparison sync.Once
func Get_Data_Decidable_decidableComparison() gopurs_runtime.Value {
	once_Data_Decidable_decidableComparison.Do(func() {
		cache_Data_Decidable_decidableComparison = gopurs_runtime.Value{Type: 9, IntVal: 336732346, UnsafePtr: unsafe.Pointer((&Constructor_Data_Decidable_Decidable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1618621146, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Decide_Decide[gopurs_runtime.Value]](Get_Data_Decide_chooseComparison()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2265116602, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Divisible_Divisible[gopurs_runtime.Value]](Get_Data_Divisible_divisibleComparison()))}
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_Decidable_spin__1133204955_3_0_6 func(gopurs_runtime.Value) uint32
_ = Call_local_Data_Decidable_spin__1133204955_3_0_6
var spin__1133204955_3_0_6 gopurs_runtime.Value
_ = spin__1133204955_3_0_6
Call_local_Data_Decidable_spin__1133204955_3_0_6 = func(v_4_loop gopurs_runtime.Value) uint32 {
spin__1133204955_3_0_6:
for {
if false { continue spin__1133204955_3_0_6 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
v_4_loop = v_4
continue spin__1133204955_3_0_6
return func() uint32 { panic("unreachable") }()
}
}
spin__1133204955_3_0_6 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_local_Data_Decidable_spin__1133204955_3_0_6(v_4_loop_val)), UnsafePtr: nil}
})
var spin_4_1_7 gopurs_runtime.Value
_ = spin_4_1_7
// FALLBACK TCO: isLoop=false len=1
spin_4_1_7 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_local_Data_Decidable_spin__1133204955_3_0_6(v_5)), UnsafePtr: nil}
})
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_local_Data_Decidable_spin__1133204955_3_0_6(gopurs_runtime.Apply(f_0, a_1))), UnsafePtr: nil}
})}))}
	})
	return cache_Data_Decidable_decidableComparison
}

type Constructor_Data_Decidable_Decidable[T_f any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[336732346] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Decidable_Decidable[any])(ptr)
		_ = c
		switch key {
		case "Decide0": return gopurs_runtime.Box(c.V0)
		case "Divisible1": return gopurs_runtime.Box(c.V1)
		case "lose": return gopurs_runtime.Box(c.V2)
		default: panic("Key not found in dictionary Constructor_Data_Decidable_Decidable: " + key)
		}
	}
}


func Call_Data_Decidable_Decidable_dollar_Dict(x_0_loop struct{
	Decide0 gopurs_runtime.Value
	Divisible1 gopurs_runtime.Value
	lose gopurs_runtime.Value
}) *Constructor_Data_Decidable_Decidable[gopurs_runtime.Value] {
var x_0 struct{
	Decide0 gopurs_runtime.Value
	Divisible1 gopurs_runtime.Value
	lose gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Decidable_Decidable[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict3("Decide0", "Divisible1", "lose", orig.Decide0, orig.Divisible1, orig.lose)
				}())
}

func Call_Data_Decidable_lose(dict_0_loop *Constructor_Data_Decidable_Decidable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Decidable_Decidable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Decidable_lost(dictDecidable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictDecidable_0 gopurs_runtime.Value = dictDecidable_0_loop
_ = dictDecidable_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDecidable_0, "lose"), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1))
}

func Call_Data_Decidable_decidableOp(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): chooseOp_1_0 shape=App(Var) bindingType=(ADT ["Data","Decide","Decide"] [(Func [(TypeVar b)] (TypeVar r))])
chooseOp_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Decide_Decide[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Decide_chooseOp(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})))
_ = chooseOp_1_0
// TAST (Let): divideOp_2_2 shape=App(Var) bindingType=(ADT ["Data","Divide","Divide"] [(Func [(TypeVar b)] (TypeVar r))])
divideOp_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Divide_Divide[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Divide_divideOp(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})))
_ = divideOp_2_2
// TAST (Let): __local_var_3_3 shape=Other bindingType=Any
__local_var_3_3 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = __local_var_3_3
// TAST (Let): divisibleOp_2_1 shape=Let(LitRecord) bindingType=(ADT ["Data","Divisible","Divisible"] [(Func [(TypeVar b)] (TypeVar r))])
divisibleOp_2_1 := (&Constructor_Data_Divisible_Divisible[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2642321722, UnsafePtr: unsafe.Pointer(divideOp_2_2)}
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_3_3
})})
_ = divisibleOp_2_1
return gopurs_runtime.Value{Type: 9, IntVal: 336732346, UnsafePtr: unsafe.Pointer((&Constructor_Data_Decidable_Decidable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1618621146, UnsafePtr: unsafe.Pointer(chooseOp_1_0)}
}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2265116602, UnsafePtr: unsafe.Pointer(divisibleOp_2_1)}
}), gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_Decidable_spin__1769020947_5_4_2 func(gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_Decidable_spin__1769020947_5_4_2
var spin__1769020947_5_4_2 gopurs_runtime.Value
_ = spin__1769020947_5_4_2
Call_local_Data_Decidable_spin__1769020947_5_4_2 = func(v_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
spin__1769020947_5_4_2:
for {
if false { continue spin__1769020947_5_4_2 }
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
v_6_loop = v_6
continue spin__1769020947_5_4_2
return func() gopurs_runtime.Value { panic("unreachable") }()
}
}
spin__1769020947_5_4_2 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_Decidable_spin__1769020947_5_4_2(v_6_loop_val)
})
var spin_6_5_3 gopurs_runtime.Value
_ = spin_6_5_3
// FALLBACK TCO: isLoop=false len=1
spin_6_5_3 = gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_Decidable_spin__1769020947_5_4_2(v_7)
})
return Call_local_Data_Decidable_spin__1769020947_5_4_2(gopurs_runtime.Apply(f_3, a_4))
})}))}
}


