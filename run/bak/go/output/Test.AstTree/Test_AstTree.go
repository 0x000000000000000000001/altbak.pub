package Test_AstTree

import (
	pkg_Bench "gopurs/output/Bench"
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Effect "gopurs/output/Effect"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Val gopurs_runtime.Value
var once_Val sync.Once
func Get_Val() gopurs_runtime.Value {
	once_Val.Do(func() {
		cache_Val = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 245188743, UnsafePtr: unsafe.Pointer(&Constructor_Val{1, value0.IntVal})}
})
	})
	return cache_Val
}

var cache_Add gopurs_runtime.Value
var once_Add sync.Once
func Get_Add() gopurs_runtime.Value {
	once_Add.Do(func() {
		cache_Add = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2937956733, UnsafePtr: unsafe.Pointer(&Constructor_Add{1, value0, value1})}
})
})
	})
	return cache_Add
}

var cache_Mul gopurs_runtime.Value
var once_Mul sync.Once
func Get_Mul() gopurs_runtime.Value {
	once_Mul.Do(func() {
		cache_Mul = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3406566728, UnsafePtr: unsafe.Pointer(&Constructor_Mul{1, value0, value1})}
})
})
	})
	return cache_Mul
}

var cache_Sub gopurs_runtime.Value
var once_Sub sync.Once
func Get_Sub() gopurs_runtime.Value {
	once_Sub.Do(func() {
		cache_Sub = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2029887576, UnsafePtr: unsafe.Pointer(&Constructor_Sub{1, value0, value1})}
})
})
	})
	return cache_Sub
}

var cache_eval gopurs_runtime.Value
var once_eval sync.Once
func Get_eval() gopurs_runtime.Value {
	once_eval.Do(func() {
		cache_eval = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_eval(v_0_box))
})
	})
	return cache_eval
}

var cache_describe gopurs_runtime.Value
var once_describe sync.Once
func Get_describe() gopurs_runtime.Value {
	once_describe.Do(func() {
		cache_describe = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("AST Evaluation:"))
	})
	return cache_describe
}

var cache_buildTree gopurs_runtime.Value
var once_buildTree sync.Once
func Get_buildTree() gopurs_runtime.Value {
	once_buildTree.Do(func() {
		cache_buildTree = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_buildTree(v_0_box.IntVal)
})
	})
	return cache_buildTree
}

var cache_act gopurs_runtime.Value
var once_act sync.Once
func Get_act() gopurs_runtime.Value {
	once_act.Do(func() {
		cache_act = Call_bind__3550378017(gopurs_runtime.Apply(pkg_Bench.Get_opaque(), gopurs_runtime.Int(3)), gopurs_runtime.Func(func(dummy_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Show.Get_showInt(), "show"), gopurs_runtime.Int(gopurs_runtime.Int(Call_eval(Call_buildTree(dummy_0.IntVal))).IntVal)).StrVal()))
}))
	})
	return cache_act
}

var cache_bind__2601835655 gopurs_runtime.Value
var once_bind__2601835655 sync.Once
func Get_bind__2601835655() gopurs_runtime.Value {
	once_bind__2601835655.Do(func() {
		cache_bind__2601835655 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__2601835655(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__2601835655
}

var cache_bind__3550378017 gopurs_runtime.Value
var once_bind__3550378017 sync.Once
func Get_bind__3550378017() gopurs_runtime.Value {
	once_bind__3550378017.Do(func() {
		cache_bind__3550378017 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__3550378017(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_bind__3550378017
}

var cache_sub__1043827704 gopurs_runtime.Value
var once_sub__1043827704 sync.Once
func Get_sub__1043827704() gopurs_runtime.Value {
	once_sub__1043827704.Do(func() {
		cache_sub__1043827704 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__1043827704(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_sub__1043827704
}

var cache_sub__3675938712 gopurs_runtime.Value
var once_sub__3675938712 sync.Once
func Get_sub__3675938712() gopurs_runtime.Value {
	once_sub__3675938712.Do(func() {
		cache_sub__3675938712 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__3675938712(gopurs_runtime.CoerceToStruct[pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sub__3675938712
}

var cache_add__560788792 gopurs_runtime.Value
var once_add__560788792 sync.Once
func Get_add__560788792() gopurs_runtime.Value {
	once_add__560788792.Do(func() {
		cache_add__560788792 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_add__560788792(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_add__560788792
}

var cache_add__1614463960 gopurs_runtime.Value
var once_add__1614463960 sync.Once
func Get_add__1614463960() gopurs_runtime.Value {
	once_add__1614463960.Do(func() {
		cache_add__1614463960 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_add__1614463960(gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_add__1614463960
}

var cache_mul__560788792 gopurs_runtime.Value
var once_mul__560788792 sync.Once
func Get_mul__560788792() gopurs_runtime.Value {
	once_mul__560788792.Do(func() {
		cache_mul__560788792 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mul__560788792(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_mul__560788792
}

var cache_mul__1614463960 gopurs_runtime.Value
var once_mul__1614463960 sync.Once
func Get_mul__1614463960() gopurs_runtime.Value {
	once_mul__1614463960.Do(func() {
		cache_mul__1614463960 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mul__1614463960(gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_mul__1614463960
}

var cache_applyEffect__2014400020 gopurs_runtime.Value
var once_applyEffect__2014400020 sync.Once
func Get_applyEffect__2014400020() gopurs_runtime.Value {
	once_applyEffect__2014400020.Do(func() {
		cache_applyEffect__2014400020 = func() gopurs_runtime.Value {
Bind1_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_monadEffect(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_monadEffect(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect.Get_functorEffect()
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, f_2, gopurs_runtime.Func(func(f_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, a_3, gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_1_1.V1, gopurs_runtime.Apply(f_prime_4, a_prime_5))
}))
}))
})
}))
}()
	})
	return cache_applyEffect__2014400020
}

var cache_bindEffect__2113658466 gopurs_runtime.Value
var once_bindEffect__2113658466 sync.Once
func Get_bindEffect__2113658466() gopurs_runtime.Value {
	once_bindEffect__2113658466.Do(func() {
		cache_bindEffect__2113658466 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect.Get_applyEffect()
}), pkg_Effect.Get_bindE())
	})
	return cache_bindEffect__2113658466
}

var cache_functorEffect__3107547953 gopurs_runtime.Value
var once_functorEffect__3107547953 sync.Once
func Get_functorEffect__3107547953() gopurs_runtime.Value {
	once_functorEffect__3107547953.Do(func() {
		cache_functorEffect__3107547953 = func() gopurs_runtime.Value {
Apply0_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_0_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_0_0.V1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure"), f_1), a_2)
})
}))
}()
	})
	return cache_functorEffect__3107547953
}

type Constructor_Val struct {
	Rc uint32
	V0 int64
}


type Constructor_Add struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


type Constructor_Mul struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


type Constructor_Sub struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func Call_eval(v_0_loop gopurs_runtime.Value) int64 {
eval:
for {
if false { continue eval }
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 245188743) {
__t0 = gopurs_runtime.Int((*Constructor_Val)(v_0.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 2937956733) {
__t0 = gopurs_runtime.Int(Call_add__560788792(gopurs_runtime.Int(gopurs_runtime.Int(Call_eval((*Constructor_Add)(v_0.UnsafePtr).V0)).IntVal), gopurs_runtime.Int(gopurs_runtime.Int(Call_eval((*Constructor_Add)(v_0.UnsafePtr).V1)).IntVal)).IntVal)
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 3406566728) {
__t0 = gopurs_runtime.Int(Call_mul__560788792(gopurs_runtime.Int(gopurs_runtime.Int(Call_eval((*Constructor_Mul)(v_0.UnsafePtr).V0)).IntVal), gopurs_runtime.Int(gopurs_runtime.Int(Call_eval((*Constructor_Mul)(v_0.UnsafePtr).V1)).IntVal)).IntVal)
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 2029887576) {
__t0 = gopurs_runtime.Int(Call_sub__1043827704(gopurs_runtime.Int(gopurs_runtime.Int(Call_eval((*Constructor_Sub)(v_0.UnsafePtr).V0)).IntVal), gopurs_runtime.Int(gopurs_runtime.Int(Call_eval((*Constructor_Sub)(v_0.UnsafePtr).V1)).IntVal)).IntVal)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0.IntVal
}
}

func Call_buildTree(v_0_loop int64) gopurs_runtime.Value {
buildTree:
for {
if false { continue buildTree }
var v_0 int64 = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0) == (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 245188743, UnsafePtr: unsafe.Pointer(&Constructor_Val{1, 1})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2937956733, UnsafePtr: unsafe.Pointer(&Constructor_Add{1, gopurs_runtime.Value{Type: 9, IntVal: 3406566728, UnsafePtr: unsafe.Pointer(&Constructor_Mul{1, gopurs_runtime.Value{Type: 9, IntVal: 245188743, UnsafePtr: unsafe.Pointer(&Constructor_Val{1, v_0})}, Call_buildTree(Call_sub__1043827704(gopurs_runtime.Int(v_0), gopurs_runtime.Int(1)).IntVal)})}, gopurs_runtime.Value{Type: 9, IntVal: 2029887576, UnsafePtr: unsafe.Pointer(&Constructor_Sub{1, Call_buildTree(Call_sub__1043827704(gopurs_runtime.Int(v_0), gopurs_runtime.Int(1)).IntVal), gopurs_runtime.Value{Type: 9, IntVal: 245188743, UnsafePtr: unsafe.Pointer(&Constructor_Val{1, 1})}})}})}
}
end_branch_0:
return __t0
}
}

func Call_bind__2601835655(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__3550378017(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Value{})
_ = __local_var_2_0
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta1_1, __local_var_2_0), gopurs_runtime.Value{})
})
}

func Call_sub__1043827704(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Int((__eta0_0.IntVal) - (__eta1_1.IntVal))
}

func Call_sub__3675938712(dict_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_add__560788792(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Int((__eta0_0.IntVal) + (__eta1_1.IntVal))
}

func Call_add__1614463960(dict_0_loop *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_mul__560788792(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Int((__eta0_0.IntVal) * (__eta1_1.IntVal))
}

func Call_mul__1614463960(dict_0_loop *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}


