package Test_RBTree

import (
	pkg_Bench "gopurs/output/Bench"
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Effect "gopurs/output/Effect"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_R gopurs_runtime.Value
var once_R sync.Once
func Get_R() gopurs_runtime.Value {
	once_R.Do(func() {
		cache_R = gopurs_runtime.Value{Type: 9, IntVal: int64(3668501016), UnsafePtr: nil}
	})
	return cache_R
}

var cache_B gopurs_runtime.Value
var once_B sync.Once
func Get_B() gopurs_runtime.Value {
	once_B.Do(func() {
		cache_B = gopurs_runtime.Value{Type: 9, IntVal: int64(1583507464), UnsafePtr: nil}
	})
	return cache_B
}

var cache_E gopurs_runtime.Value
var once_E sync.Once
func Get_E() gopurs_runtime.Value {
	once_E.Do(func() {
		cache_E = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(nil))}
	})
	return cache_E
}

var cache_T gopurs_runtime.Value
var once_T sync.Once
func Get_T() gopurs_runtime.Value {
	once_T.Do(func() {
		cache_T = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, uint32(value0.IntVal), gopurs_runtime.CoerceToStruct[Constructor_T](value1), value2.IntVal, gopurs_runtime.CoerceToStruct[Constructor_T](value3)})}
})
})
})
})
	})
	return cache_T
}

var cache_max gopurs_runtime.Value
var once_max sync.Once
func Get_max() gopurs_runtime.Value {
	once_max.Do(func() {
		cache_max = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_max(x_0_box.IntVal, y_1_box.IntVal))
})
	})
	return cache_max
}

var cache_makeBlack gopurs_runtime.Value
var once_makeBlack sync.Once
func Get_makeBlack() gopurs_runtime.Value {
	once_makeBlack.Do(func() {
		cache_makeBlack = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(Call_makeBlack(gopurs_runtime.CoerceToStruct[Constructor_T](v_0_box)))}
})
	})
	return cache_makeBlack
}

var cache_describe gopurs_runtime.Value
var once_describe sync.Once
func Get_describe() gopurs_runtime.Value {
	once_describe.Do(func() {
		cache_describe = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("Red-Black Tree (100k Worst-Case Insertions):"))
	})
	return cache_describe
}

var cache_depth gopurs_runtime.Value
var once_depth sync.Once
func Get_depth() gopurs_runtime.Value {
	once_depth.Do(func() {
		cache_depth = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_depth(gopurs_runtime.CoerceToStruct[Constructor_T](v_0_box)))
})
	})
	return cache_depth
}

var cache_balance gopurs_runtime.Value
var once_balance sync.Once
func Get_balance() gopurs_runtime.Value {
	once_balance.Do(func() {
		cache_balance = gopurs_runtime.Func4(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value, v3_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(Call_balance(uint32(v_0_box.IntVal), gopurs_runtime.CoerceToStruct[Constructor_T](v1_1_box), v2_2_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_T](v3_3_box)))}
})
	})
	return cache_balance
}

var cache_ins gopurs_runtime.Value
var once_ins sync.Once
func Get_ins() gopurs_runtime.Value {
	once_ins.Do(func() {
		cache_ins = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(Call_ins(v_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_T](v1_1_box)))}
})
	})
	return cache_ins
}

var cache_insert gopurs_runtime.Value
var once_insert sync.Once
func Get_insert() gopurs_runtime.Value {
	once_insert.Do(func() {
		cache_insert = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(Call_insert(x_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_T](s_1_box)))}
})
	})
	return cache_insert
}

var cache_buildTree gopurs_runtime.Value
var once_buildTree sync.Once
func Get_buildTree() gopurs_runtime.Value {
	once_buildTree.Do(func() {
		cache_buildTree = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(Call_buildTree(v_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_T](v1_1_box)))}
})
	})
	return cache_buildTree
}

var cache_act gopurs_runtime.Value
var once_act sync.Once
func Get_act() gopurs_runtime.Value {
	once_act.Do(func() {
		cache_act = gopurs_runtime.Apply2(Get_bind__3550378017(), gopurs_runtime.Apply(pkg_Bench.Get_opaque(), gopurs_runtime.Int(100000)), gopurs_runtime.Func(func(dummy_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Show.Get_showInt(), "show"), gopurs_runtime.Int(Call_depth(gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(Call_buildTree(dummy_0.IntVal, gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(nil))})))})))))
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
		cache_bind__3550378017 = gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind")
	})
	return cache_bind__3550378017
}

var cache_compare__821463600 gopurs_runtime.Value
var once_compare__821463600 sync.Once
func Get_compare__821463600() gopurs_runtime.Value {
	once_compare__821463600.Do(func() {
		cache_compare__821463600 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compare__821463600(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compare__821463600
}

var cache_greaterThan__4087042607 gopurs_runtime.Value
var once_greaterThan__4087042607 sync.Once
func Get_greaterThan__4087042607() gopurs_runtime.Value {
	once_greaterThan__4087042607.Do(func() {
		cache_greaterThan__4087042607 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThan__4087042607(a1_0_box, a2_1_box))
})
	})
	return cache_greaterThan__4087042607
}

var cache_greaterThan__1409282474 gopurs_runtime.Value
var once_greaterThan__1409282474 sync.Once
func Get_greaterThan__1409282474() gopurs_runtime.Value {
	once_greaterThan__1409282474.Do(func() {
		cache_greaterThan__1409282474 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThan__1409282474(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_greaterThan__1409282474
}

var cache_lessThan__4087042607 gopurs_runtime.Value
var once_lessThan__4087042607 sync.Once
func Get_lessThan__4087042607() gopurs_runtime.Value {
	once_lessThan__4087042607.Do(func() {
		cache_lessThan__4087042607 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThan__4087042607(a1_0_box, a2_1_box))
})
	})
	return cache_lessThan__4087042607
}

var cache_lessThan__1409282474 gopurs_runtime.Value
var once_lessThan__1409282474 sync.Once
func Get_lessThan__1409282474() gopurs_runtime.Value {
	once_lessThan__1409282474.Do(func() {
		cache_lessThan__1409282474 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThan__1409282474(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_lessThan__1409282474
}

var cache_sub__1043827704 gopurs_runtime.Value
var once_sub__1043827704 sync.Once
func Get_sub__1043827704() gopurs_runtime.Value {
	once_sub__1043827704.Do(func() {
		cache_sub__1043827704 = pkg_Data_Ring.Get_intSub()
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
		cache_add__560788792 = pkg_Data_Semiring.Get_intAdd()
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

type Constructor_R struct {
	Rc uint32
}


type Constructor_B struct {
	Rc uint32
}


type Constructor_E struct {
	Rc uint32
}


type Constructor_T struct {
	Rc uint32
	V0 uint32
	V1 *Constructor_T
	V2 int64
	V3 *Constructor_T
}


func Call_max(x_0_loop int64, y_1_loop int64) int64 {
var x_0 int64 = x_0_loop
_ = x_0
var y_1 int64 = y_1_loop
_ = y_1
var __t0 int64
{
if (gopurs_runtime.Bool(Call_greaterThan__4087042607(gopurs_runtime.Int(x_0), gopurs_runtime.Int(y_1))).IntVal) != (0) {
__t0 = x_0
goto end_branch_0
} else {

}
}
{
__t0 = y_1
}
end_branch_0:
return __t0
}

func Call_makeBlack(v_0_loop *Constructor_T) *Constructor_T {
var v_0 *Constructor_T = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 3983586014 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, 1583507464, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V3})}
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 3983586014 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(nil))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[Constructor_T](__t0)
}

func Call_depth(v_0_loop *Constructor_T) int64 {
depth:
for {
if false { continue depth }
var v_0 *Constructor_T = v_0_loop
_ = v_0
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 3983586014 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t3 = gopurs_runtime.Int(0)
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 3983586014 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
__local_var_1_0 := gopurs_runtime.Int(Call_depth((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1))
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Int(Call_depth((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V3))
_ = __local_var_2_1
var __t2 int64
{
if (gopurs_runtime.Bool(Call_greaterThan__4087042607(__local_var_1_0, __local_var_2_1)).IntVal) != (0) {
__t2 = __local_var_1_0.IntVal
goto end_branch_2
} else {

}
}
{
__t2 = __local_var_2_1.IntVal
}
end_branch_2:
__t3 = gopurs_runtime.Int(gopurs_runtime.Apply2(Get_add__560788792(), gopurs_runtime.Int(1), gopurs_runtime.Int(__t2)).IntVal)
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3.IntVal
}
}

func Call_balance(v_0_loop uint32, v1_1_loop *Constructor_T, v2_2_loop int64, v3_3_loop *Constructor_T) *Constructor_T {
var v_0 uint32 = v_0_loop
_ = v_0
var v1_1 *Constructor_T = v1_1_loop
_ = v1_1
var v2_2 int64 = v2_2_loop
_ = v2_2
var v3_3 *Constructor_T = v3_3_loop
_ = v3_3
var __t85 *Constructor_T
{
if (v_0 == 1583507464) {
var __t84 *Constructor_T
{
if (gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 3983586014 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr != nil) {
var __t71 *Constructor_T
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0), UnsafePtr: nil}
if (uint32(__t_tag_0.IntVal) == 3668501016) {
var __t58 *Constructor_T
{
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V1)}
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 3983586014 && __t_tag_1.UnsafePtr != nil) {
var __t30 *Constructor_T
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V1)}.UnsafePtr).V0), UnsafePtr: nil}
if (uint32(__t_tag_2.IntVal) == 3668501016) {
__t30 = &Constructor_T{1, 3668501016, gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, 1583507464, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V1)}.UnsafePtr).V1, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V1)}.UnsafePtr).V2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V1)}.UnsafePtr).V3})}), (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, 1583507464, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V3, v2_2, v3_3})})}
goto end_branch_30
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V3)}
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 3983586014 && __t_tag_3.UnsafePtr != nil) {
var __t17 *Constructor_T
{
var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V3)}.UnsafePtr).V0), UnsafePtr: nil}
if (uint32(__t_tag_4.IntVal) == 3668501016) {
__t17 = &Constructor_T{1, 3668501016, gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, 1583507464, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V1, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V3)}.UnsafePtr).V1})}), (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V3)}.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, 1583507464, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V3)}.UnsafePtr).V3, v2_2, v3_3})})}
goto end_branch_17
} else {

}
}
{
var __t_and_6 bool = false
if (gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.IntVal == 3983586014 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr != nil) {

var __t_tag_5 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V0), UnsafePtr: nil}
__t_and_6 = (uint32(__t_tag_5.IntVal) == 3668501016)
}
if __t_and_6 {
var __t16 *Constructor_T
{
var __t_tag_7 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}
if (__t_tag_7.Type == 9 && __t_tag_7.IntVal == 3983586014 && __t_tag_7.UnsafePtr != nil) {
var __t12 *Constructor_T
{
var __t_tag_8 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr).V0), UnsafePtr: nil}
if (uint32(__t_tag_8.IntVal) == 3668501016) {
__t12 = &Constructor_T{1, 3668501016, gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, 1583507464, v1_1, v2_2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr).V1})}), (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, 1583507464, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr).V3, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3})})}
goto end_branch_12
} else {

}
}
{
var __t_tag_9 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}
var __t_and_11 bool = false
if (__t_tag_9.Type == 9 && __t_tag_9.IntVal == 3983586014 && __t_tag_9.UnsafePtr != nil) {

var __t_tag_10 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V0), UnsafePtr: nil}
__t_and_11 = (uint32(__t_tag_10.IntVal) == 3668501016)
}
if __t_and_11 {
__t12 = &Constructor_T{1, 3668501016, gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, 1583507464, v1_1, v2_2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1})}), (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, 1583507464, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V1, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V3})})}
goto end_branch_12
} else {

}
}
{
__t12 = &Constructor_T{1, v_0, v1_1, v2_2, v3_3}
}
end_branch_12:
__t16 = __t12
goto end_branch_16
} else {

}
}
{
var __t_tag_13 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}
var __t_and_15 bool = false
if (__t_tag_13.Type == 9 && __t_tag_13.IntVal == 3983586014 && __t_tag_13.UnsafePtr != nil) {

var __t_tag_14 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V0), UnsafePtr: nil}
__t_and_15 = (uint32(__t_tag_14.IntVal) == 3668501016)
}
if __t_and_15 {
__t16 = &Constructor_T{1, 3668501016, gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, 1583507464, v1_1, v2_2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1})}), (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, 1583507464, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V1, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V3})})}
goto end_branch_16
} else {

}
}
{
__t16 = &Constructor_T{1, v_0, v1_1, v2_2, v3_3}
}
end_branch_16:
__t17 = __t16
goto end_branch_17
} else {

}
}
{
__t17 = &Constructor_T{1, v_0, v1_1, v2_2, v3_3}
}
end_branch_17:
__t30 = __t17
goto end_branch_30
} else {

}
}
{
var __t_and_19 bool = false
if (gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.IntVal == 3983586014 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr != nil) {

var __t_tag_18 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V0), UnsafePtr: nil}
__t_and_19 = (uint32(__t_tag_18.IntVal) == 3668501016)
}
if __t_and_19 {
var __t29 *Constructor_T
{
var __t_tag_20 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}
if (__t_tag_20.Type == 9 && __t_tag_20.IntVal == 3983586014 && __t_tag_20.UnsafePtr != nil) {
var __t25 *Constructor_T
{
var __t_tag_21 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr).V0), UnsafePtr: nil}
if (uint32(__t_tag_21.IntVal) == 3668501016) {
__t25 = &Constructor_T{1, 3668501016, gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, 1583507464, v1_1, v2_2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr).V1})}), (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, 1583507464, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr).V3, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3})})}
goto end_branch_25
} else {

}
}
{
var __t_tag_22 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}
var __t_and_24 bool = false
if (__t_tag_22.Type == 9 && __t_tag_22.IntVal == 3983586014 && __t_tag_22.UnsafePtr != nil) {

var __t_tag_23 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V0), UnsafePtr: nil}
__t_and_24 = (uint32(__t_tag_23.IntVal) == 3668501016)
}
if __t_and_24 {
__t25 = &Constructor_T{1, 3668501016, gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, 1583507464, v1_1, v2_2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1})}), (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, 1583507464, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V1, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V3})})}
goto end_branch_25
} else {

}
}
{
__t25 = &Constructor_T{1, v_0, v1_1, v2_2, v3_3}
}
end_branch_25:
__t29 = __t25
goto end_branch_29
} else {

}
}
{
var __t_tag_26 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}
var __t_and_28 bool = false
if (__t_tag_26.Type == 9 && __t_tag_26.IntVal == 3983586014 && __t_tag_26.UnsafePtr != nil) {

var __t_tag_27 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V0), UnsafePtr: nil}
__t_and_28 = (uint32(__t_tag_27.IntVal) == 3668501016)
}
if __t_and_28 {
__t29 = &Constructor_T{1, 3668501016, gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, 1583507464, v1_1, v2_2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1})}), (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, 1583507464, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V1, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V3})})}
goto end_branch_29
} else {

}
}
{
__t29 = &Constructor_T{1, v_0, v1_1, v2_2, v3_3}
}
end_branch_29:
__t30 = __t29
goto end_branch_30
} else {

}
}
{
__t30 = &Constructor_T{1, v_0, v1_1, v2_2, v3_3}
}
end_branch_30:
__t58 = __t30
goto end_branch_58
} else {

}
}
{
var __t_tag_31 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V3)}
if (__t_tag_31.Type == 9 && __t_tag_31.IntVal == 3983586014 && __t_tag_31.UnsafePtr != nil) {
var __t45 *Constructor_T
{
var __t_tag_32 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V3)}.UnsafePtr).V0), UnsafePtr: nil}
if (uint32(__t_tag_32.IntVal) == 3668501016) {
__t45 = &Constructor_T{1, 3668501016, gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, 1583507464, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V1, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V3)}.UnsafePtr).V1})}), (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V3)}.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, 1583507464, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V3)}.UnsafePtr).V3, v2_2, v3_3})})}
goto end_branch_45
} else {

}
}
{
var __t_and_34 bool = false
if (gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.IntVal == 3983586014 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr != nil) {

var __t_tag_33 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V0), UnsafePtr: nil}
__t_and_34 = (uint32(__t_tag_33.IntVal) == 3668501016)
}
if __t_and_34 {
var __t44 *Constructor_T
{
var __t_tag_35 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}
if (__t_tag_35.Type == 9 && __t_tag_35.IntVal == 3983586014 && __t_tag_35.UnsafePtr != nil) {
var __t40 *Constructor_T
{
var __t_tag_36 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr).V0), UnsafePtr: nil}
if (uint32(__t_tag_36.IntVal) == 3668501016) {
__t40 = &Constructor_T{1, 3668501016, gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, 1583507464, v1_1, v2_2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr).V1})}), (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, 1583507464, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr).V3, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3})})}
goto end_branch_40
} else {

}
}
{
var __t_tag_37 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}
var __t_and_39 bool = false
if (__t_tag_37.Type == 9 && __t_tag_37.IntVal == 3983586014 && __t_tag_37.UnsafePtr != nil) {

var __t_tag_38 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V0), UnsafePtr: nil}
__t_and_39 = (uint32(__t_tag_38.IntVal) == 3668501016)
}
if __t_and_39 {
__t40 = &Constructor_T{1, 3668501016, gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, 1583507464, v1_1, v2_2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1})}), (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, 1583507464, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V1, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V3})})}
goto end_branch_40
} else {

}
}
{
__t40 = &Constructor_T{1, v_0, v1_1, v2_2, v3_3}
}
end_branch_40:
__t44 = __t40
goto end_branch_44
} else {

}
}
{
var __t_tag_41 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}
var __t_and_43 bool = false
if (__t_tag_41.Type == 9 && __t_tag_41.IntVal == 3983586014 && __t_tag_41.UnsafePtr != nil) {

var __t_tag_42 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V0), UnsafePtr: nil}
__t_and_43 = (uint32(__t_tag_42.IntVal) == 3668501016)
}
if __t_and_43 {
__t44 = &Constructor_T{1, 3668501016, gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, 1583507464, v1_1, v2_2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1})}), (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, 1583507464, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V1, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V3})})}
goto end_branch_44
} else {

}
}
{
__t44 = &Constructor_T{1, v_0, v1_1, v2_2, v3_3}
}
end_branch_44:
__t45 = __t44
goto end_branch_45
} else {

}
}
{
__t45 = &Constructor_T{1, v_0, v1_1, v2_2, v3_3}
}
end_branch_45:
__t58 = __t45
goto end_branch_58
} else {

}
}
{
var __t_and_47 bool = false
if (gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.IntVal == 3983586014 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr != nil) {

var __t_tag_46 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V0), UnsafePtr: nil}
__t_and_47 = (uint32(__t_tag_46.IntVal) == 3668501016)
}
if __t_and_47 {
var __t57 *Constructor_T
{
var __t_tag_48 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}
if (__t_tag_48.Type == 9 && __t_tag_48.IntVal == 3983586014 && __t_tag_48.UnsafePtr != nil) {
var __t53 *Constructor_T
{
var __t_tag_49 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr).V0), UnsafePtr: nil}
if (uint32(__t_tag_49.IntVal) == 3668501016) {
__t53 = &Constructor_T{1, 3668501016, gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, 1583507464, v1_1, v2_2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr).V1})}), (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, 1583507464, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr).V3, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3})})}
goto end_branch_53
} else {

}
}
{
var __t_tag_50 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}
var __t_and_52 bool = false
if (__t_tag_50.Type == 9 && __t_tag_50.IntVal == 3983586014 && __t_tag_50.UnsafePtr != nil) {

var __t_tag_51 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V0), UnsafePtr: nil}
__t_and_52 = (uint32(__t_tag_51.IntVal) == 3668501016)
}
if __t_and_52 {
__t53 = &Constructor_T{1, 3668501016, gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, 1583507464, v1_1, v2_2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1})}), (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, 1583507464, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V1, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V3})})}
goto end_branch_53
} else {

}
}
{
__t53 = &Constructor_T{1, v_0, v1_1, v2_2, v3_3}
}
end_branch_53:
__t57 = __t53
goto end_branch_57
} else {

}
}
{
var __t_tag_54 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}
var __t_and_56 bool = false
if (__t_tag_54.Type == 9 && __t_tag_54.IntVal == 3983586014 && __t_tag_54.UnsafePtr != nil) {

var __t_tag_55 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V0), UnsafePtr: nil}
__t_and_56 = (uint32(__t_tag_55.IntVal) == 3668501016)
}
if __t_and_56 {
__t57 = &Constructor_T{1, 3668501016, gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, 1583507464, v1_1, v2_2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1})}), (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, 1583507464, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V1, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V3})})}
goto end_branch_57
} else {

}
}
{
__t57 = &Constructor_T{1, v_0, v1_1, v2_2, v3_3}
}
end_branch_57:
__t58 = __t57
goto end_branch_58
} else {

}
}
{
__t58 = &Constructor_T{1, v_0, v1_1, v2_2, v3_3}
}
end_branch_58:
__t71 = __t58
goto end_branch_71
} else {

}
}
{
var __t_and_60 bool = false
if (gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.IntVal == 3983586014 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr != nil) {

var __t_tag_59 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V0), UnsafePtr: nil}
__t_and_60 = (uint32(__t_tag_59.IntVal) == 3668501016)
}
if __t_and_60 {
var __t70 *Constructor_T
{
var __t_tag_61 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}
if (__t_tag_61.Type == 9 && __t_tag_61.IntVal == 3983586014 && __t_tag_61.UnsafePtr != nil) {
var __t66 *Constructor_T
{
var __t_tag_62 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr).V0), UnsafePtr: nil}
if (uint32(__t_tag_62.IntVal) == 3668501016) {
__t66 = &Constructor_T{1, 3668501016, gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, 1583507464, v1_1, v2_2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr).V1})}), (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, 1583507464, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr).V3, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3})})}
goto end_branch_66
} else {

}
}
{
var __t_tag_63 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}
var __t_and_65 bool = false
if (__t_tag_63.Type == 9 && __t_tag_63.IntVal == 3983586014 && __t_tag_63.UnsafePtr != nil) {

var __t_tag_64 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V0), UnsafePtr: nil}
__t_and_65 = (uint32(__t_tag_64.IntVal) == 3668501016)
}
if __t_and_65 {
__t66 = &Constructor_T{1, 3668501016, gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, 1583507464, v1_1, v2_2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1})}), (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, 1583507464, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V1, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V3})})}
goto end_branch_66
} else {

}
}
{
__t66 = &Constructor_T{1, v_0, v1_1, v2_2, v3_3}
}
end_branch_66:
__t70 = __t66
goto end_branch_70
} else {

}
}
{
var __t_tag_67 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}
var __t_and_69 bool = false
if (__t_tag_67.Type == 9 && __t_tag_67.IntVal == 3983586014 && __t_tag_67.UnsafePtr != nil) {

var __t_tag_68 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V0), UnsafePtr: nil}
__t_and_69 = (uint32(__t_tag_68.IntVal) == 3668501016)
}
if __t_and_69 {
__t70 = &Constructor_T{1, 3668501016, gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, 1583507464, v1_1, v2_2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1})}), (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, 1583507464, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V1, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V3})})}
goto end_branch_70
} else {

}
}
{
__t70 = &Constructor_T{1, v_0, v1_1, v2_2, v3_3}
}
end_branch_70:
__t71 = __t70
goto end_branch_71
} else {

}
}
{
__t71 = &Constructor_T{1, v_0, v1_1, v2_2, v3_3}
}
end_branch_71:
__t84 = __t71
goto end_branch_84
} else {

}
}
{
var __t_and_73 bool = false
if (gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.IntVal == 3983586014 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr != nil) {

var __t_tag_72 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V0), UnsafePtr: nil}
__t_and_73 = (uint32(__t_tag_72.IntVal) == 3668501016)
}
if __t_and_73 {
var __t83 *Constructor_T
{
var __t_tag_74 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}
if (__t_tag_74.Type == 9 && __t_tag_74.IntVal == 3983586014 && __t_tag_74.UnsafePtr != nil) {
var __t79 *Constructor_T
{
var __t_tag_75 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr).V0), UnsafePtr: nil}
if (uint32(__t_tag_75.IntVal) == 3668501016) {
__t79 = &Constructor_T{1, 3668501016, gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, 1583507464, v1_1, v2_2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr).V1})}), (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, 1583507464, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr).V3, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3})})}
goto end_branch_79
} else {

}
}
{
var __t_tag_76 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}
var __t_and_78 bool = false
if (__t_tag_76.Type == 9 && __t_tag_76.IntVal == 3983586014 && __t_tag_76.UnsafePtr != nil) {

var __t_tag_77 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V0), UnsafePtr: nil}
__t_and_78 = (uint32(__t_tag_77.IntVal) == 3668501016)
}
if __t_and_78 {
__t79 = &Constructor_T{1, 3668501016, gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, 1583507464, v1_1, v2_2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1})}), (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, 1583507464, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V1, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V3})})}
goto end_branch_79
} else {

}
}
{
__t79 = &Constructor_T{1, v_0, v1_1, v2_2, v3_3}
}
end_branch_79:
__t83 = __t79
goto end_branch_83
} else {

}
}
{
var __t_tag_80 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}
var __t_and_82 bool = false
if (__t_tag_80.Type == 9 && __t_tag_80.IntVal == 3983586014 && __t_tag_80.UnsafePtr != nil) {

var __t_tag_81 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V0), UnsafePtr: nil}
__t_and_82 = (uint32(__t_tag_81.IntVal) == 3668501016)
}
if __t_and_82 {
__t83 = &Constructor_T{1, 3668501016, gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, 1583507464, v1_1, v2_2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1})}), (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, 1583507464, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V1, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V3})})}
goto end_branch_83
} else {

}
}
{
__t83 = &Constructor_T{1, v_0, v1_1, v2_2, v3_3}
}
end_branch_83:
__t84 = __t83
goto end_branch_84
} else {

}
}
{
__t84 = &Constructor_T{1, v_0, v1_1, v2_2, v3_3}
}
end_branch_84:
__t85 = __t84
goto end_branch_85
} else {

}
}
{
__t85 = &Constructor_T{1, v_0, v1_1, v2_2, v3_3}
}
end_branch_85:
return gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(__t85)})
}

func Call_ins(v_0_loop int64, v1_1_loop *Constructor_T) *Constructor_T {
ins:
for {
if false { continue ins }
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 *Constructor_T = v1_1_loop
_ = v1_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 3983586014 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, 3668501016, gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(nil))}), v_0, gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(nil))})})}
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 3983586014 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr != nil) {
var __t1 *Constructor_T
{
if (gopurs_runtime.Bool(Call_lessThan__4087042607(gopurs_runtime.Int(v_0), gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V2))).IntVal) != (0) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(Call_balance((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(Call_ins(v_0, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V1))}), (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V3))})
goto end_branch_1
} else {

}
}
{
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(Call_greaterThan__4087042607(gopurs_runtime.Int(v_0), gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V2))).IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(Call_balance((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V1, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(Call_ins(v_0, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V3))})))}))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V1, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V3})}
}
end_branch_0:
__t1 = gopurs_runtime.CoerceToStruct[Constructor_T](__t0)
}
end_branch_1:
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(__t1)}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[Constructor_T](__t2)
}
}

func Call_insert(x_0_loop int64, s_1_loop *Constructor_T) *Constructor_T {
var x_0 int64 = x_0_loop
_ = x_0
var s_1 *Constructor_T = s_1_loop
_ = s_1
__local_var_2_0 := gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(Call_ins(x_0, s_1))}
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 3983586014 && __local_var_2_0.UnsafePtr != nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, 1583507464, (*Constructor_T)(__local_var_2_0.UnsafePtr).V1, (*Constructor_T)(__local_var_2_0.UnsafePtr).V2, (*Constructor_T)(__local_var_2_0.UnsafePtr).V3})}
goto end_branch_1
} else {

}
}
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 3983586014 && __local_var_2_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(nil))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[Constructor_T](__t1)
}

func Call_buildTree(v_0_loop int64, v1_1_loop *Constructor_T) *Constructor_T {
buildTree:
for {
if false { continue buildTree }
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 *Constructor_T = v1_1_loop
_ = v1_1
var __t0 *Constructor_T
{
if (v_0) == (0) {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
v_0_loop = gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Int(v_0), gopurs_runtime.Int(1)).IntVal
v1_1_loop = gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(Call_insert(v_0, v1_1))})
continue buildTree
__t0 = gopurs_runtime.CoerceToStruct[Constructor_T](gopurs_runtime.Value{})
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

func Call_compare__821463600(dict_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_greaterThan__4087042607(a1_0_loop gopurs_runtime.Value, a2_1_loop gopurs_runtime.Value) bool {
var a1_0 gopurs_runtime.Value = a1_0_loop
_ = a1_0
var a2_1 gopurs_runtime.Value = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0.IntVal) > (a2_1.IntVal) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
return __t0
}

func Call_greaterThan__1409282474(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, a1_1, a2_2)
if (uint32(__t_tag_0.IntVal) == 380165415) {
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

func Call_lessThan__4087042607(a1_0_loop gopurs_runtime.Value, a2_1_loop gopurs_runtime.Value) bool {
var a1_0 gopurs_runtime.Value = a1_0_loop
_ = a1_0
var a2_1 gopurs_runtime.Value = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0.IntVal) < (a2_1.IntVal) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
return __t0
}

func Call_lessThan__1409282474(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, a1_1, a2_2)
if (uint32(__t_tag_0.IntVal) == 1527465420) {
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

func Call_sub__3675938712(dict_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_add__1614463960(dict_0_loop *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}


