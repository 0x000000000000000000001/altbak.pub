package Test_ListOps

import (
	pkg_Bench "gopurs/output/Bench"
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
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

var cache_Nil gopurs_runtime.Value
var once_Nil sync.Once
func Get_Nil() gopurs_runtime.Value {
	once_Nil.Do(func() {
		cache_Nil = gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}
	})
	return cache_Nil
}

var cache_Cons gopurs_runtime.Value
var once_Cons sync.Once
func Get_Cons() gopurs_runtime.Value {
	once_Cons.Do(func() {
		cache_Cons = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, value0, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](value1)})}
})
})
	})
	return cache_Cons
}

var cache_go__range gopurs_runtime.Value
var once_go__range sync.Once
func Get_go__range() gopurs_runtime.Value {
	once_go__range.Do(func() {
		cache_go__range = gopurs_runtime.Func2(func(start_0_box gopurs_runtime.Value, end_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(Call_go__range(start_0_box.IntVal, end_1_box.IntVal))}
})
	})
	return cache_go__range
}

var cache_foldl gopurs_runtime.Value
var once_foldl sync.Once
func Get_foldl() gopurs_runtime.Value {
	once_foldl.Do(func() {
		cache_foldl = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_foldl
}

var cache_filterEvens gopurs_runtime.Value
var once_filterEvens sync.Once
func Get_filterEvens() gopurs_runtime.Value {
	once_filterEvens.Do(func() {
		cache_filterEvens = gopurs_runtime.Func(func(lst_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(Call_filterEvens(gopurs_runtime.CoerceToStruct[Constructor_Cons[int64]](lst_0_box)))}
})
	})
	return cache_filterEvens
}

var cache_sumEvens gopurs_runtime.Value
var once_sumEvens sync.Once
func Get_sumEvens() gopurs_runtime.Value {
	once_sumEvens.Do(func() {
		cache_sumEvens = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_sumEvens(n_0_box.IntVal))
})
	})
	return cache_sumEvens
}

var cache_describe gopurs_runtime.Value
var once_describe sync.Once
func Get_describe() gopurs_runtime.Value {
	once_describe.Do(func() {
		cache_describe = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("List Processing (900 elements):"))
	})
	return cache_describe
}

var cache_act gopurs_runtime.Value
var once_act sync.Once
func Get_act() gopurs_runtime.Value {
	once_act.Do(func() {
		cache_act = Call_bind__3550378017(gopurs_runtime.Apply(pkg_Bench.Get_opaque(), gopurs_runtime.Int(900)), gopurs_runtime.Func(func(dummy_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Show.Get_showInt(), "show"), gopurs_runtime.Int(Call_foldl(pkg_Data_Semiring.Get_intAdd(), gopurs_runtime.Int(0), gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[int64]](gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(Call_filterEvens(gopurs_runtime.CoerceToStruct[Constructor_Cons[int64]](gopurs_runtime.Apply2(Get_go__range(), gopurs_runtime.Int(1), gopurs_runtime.Int(dummy_0.IntVal)))))}))})).IntVal)).StrVal()))
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

var cache_eq__2843686287 gopurs_runtime.Value
var once_eq__2843686287 sync.Once
func Get_eq__2843686287() gopurs_runtime.Value {
	once_eq__2843686287.Do(func() {
		cache_eq__2843686287 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq__2843686287(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_eq__2843686287
}

var cache_eq__2384498378 gopurs_runtime.Value
var once_eq__2384498378 sync.Once
func Get_eq__2384498378() gopurs_runtime.Value {
	once_eq__2384498378.Do(func() {
		cache_eq__2384498378 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq__2384498378(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_eq__2384498378
}

var cache_mod__2185172824 gopurs_runtime.Value
var once_mod__2185172824 sync.Once
func Get_mod__2185172824() gopurs_runtime.Value {
	once_mod__2185172824.Do(func() {
		cache_mod__2185172824 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mod__2185172824(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_mod__2185172824
}

var cache_mod__2579358968 gopurs_runtime.Value
var once_mod__2579358968 sync.Once
func Get_mod__2579358968() gopurs_runtime.Value {
	once_mod__2579358968.Do(func() {
		cache_mod__2579358968 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mod__2579358968(gopurs_runtime.CoerceToStruct[pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_mod__2579358968
}

var cache_compare__669572705 gopurs_runtime.Value
var once_compare__669572705 sync.Once
func Get_compare__669572705() gopurs_runtime.Value {
	once_compare__669572705.Do(func() {
		cache_compare__669572705 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compare__669572705(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compare__669572705
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

var cache_foldl__3262866295 gopurs_runtime.Value
var once_foldl__3262866295 sync.Once
func Get_foldl__3262866295() gopurs_runtime.Value {
	once_foldl__3262866295.Do(func() {
		cache_foldl__3262866295 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__3262866295(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_foldl__3262866295
}

var cache_foldl__1255354935 gopurs_runtime.Value
var once_foldl__1255354935 sync.Once
func Get_foldl__1255354935() gopurs_runtime.Value {
	once_foldl__1255354935.Do(func() {
		cache_foldl__1255354935 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__1255354935(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_foldl__1255354935
}

type Constructor_Nil[T_a any] struct {
	Rc uint32
}


type Constructor_Cons[T_a any] struct {
	Rc uint32
	V0 T_a
	V1 *Constructor_Cons[gopurs_runtime.Value]
}


func Call_go__range(start_0_loop int64, end_1_loop int64) *Constructor_Cons[int64] {
var start_0 int64 = start_0_loop
_ = start_0
var end_1 int64 = end_1_loop
_ = end_1
var go__go_2_0_0 gopurs_runtime.Value
go__go_2_0_0 = gopurs_runtime.Func(func(curr_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var curr_3_loop int64 = curr_3_loop_val.IntVal
var acc_4_loop *Constructor_Cons[int64] = gopurs_runtime.CoerceToStruct[Constructor_Cons[int64]](acc_4_loop_val)
go__go_2_0_0:
for {
if false { continue go__go_2_0_0 }
var curr_3 int64 = curr_3_loop
_ = curr_3
var acc_4 *Constructor_Cons[int64] = acc_4_loop
_ = acc_4
var __t1 *Constructor_Cons[int64]
{
if (gopurs_runtime.Bool(Call_lessThan__4087042607(gopurs_runtime.Int(curr_3), gopurs_runtime.Int(start_0))).IntVal) != (0) {
__t1 = acc_4
goto end_branch_1
} else {

}
}
{
curr_3_loop = Call_sub__1043827704(gopurs_runtime.Int(curr_3), gopurs_runtime.Int(1)).IntVal
acc_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Cons[int64]](gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Int(curr_3), gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(acc_4)})})})
continue go__go_2_0_0
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Cons[int64]](gopurs_runtime.Value{})
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(__t1)}
}
}()
})
})
return gopurs_runtime.CoerceToStruct[Constructor_Cons[int64]](gopurs_runtime.Apply2(go__go_2_0_0, gopurs_runtime.Int(end_1), gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}))
}

func Call_foldl(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
foldl:
for {
if false { continue foldl }
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Cons[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 1127792131 && gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 1127792131 && gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
v_0_loop = v_0
v1_1_loop = gopurs_runtime.Apply2(v_0, v1_1, (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0)
v2_2_loop = (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V1
continue foldl
__t0 = gopurs_runtime.Value{}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}
}

func Call_filterEvens(lst_0_loop *Constructor_Cons[int64]) *Constructor_Cons[int64] {
var lst_0 *Constructor_Cons[int64] = lst_0_loop
_ = lst_0
var go__go_1_0_1 gopurs_runtime.Value
go__go_1_0_1 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
var v1_3_loop *Constructor_Cons[int64] = gopurs_runtime.CoerceToStruct[Constructor_Cons[int64]](v1_3_loop_val)
go__go_1_0_1:
for {
if false { continue go__go_1_0_1 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var v1_3 *Constructor_Cons[int64] = v1_3_loop
_ = v1_3
var __t2 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 1127792131 && v_2.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v1_3)}
goto end_branch_2
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 1127792131 && v_2.UnsafePtr != nil) {
var __t1 *Constructor_Cons[int64]
{
if (Call_eq__2843686287(gopurs_runtime.Int(Call_mod__2185172824(gopurs_runtime.Int((*Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0.IntVal), gopurs_runtime.Int(2)).IntVal), gopurs_runtime.Int(0)).IntVal) != (0) {
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)}
v1_3_loop = gopurs_runtime.CoerceToStruct[Constructor_Cons[int64]](gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Int((*Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0.IntVal), gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v1_3)})})})
continue go__go_1_0_1
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Cons[int64]](gopurs_runtime.Value{})
goto end_branch_1
} else {

}
}
{
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)}
v1_3_loop = v1_3
continue go__go_1_0_1
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Cons[int64]](gopurs_runtime.Value{})
}
end_branch_1:
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(__t1)}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[int64]](__t2))}
}
}()
})
})
return gopurs_runtime.CoerceToStruct[Constructor_Cons[int64]](gopurs_runtime.Apply2(go__go_1_0_1, gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(lst_0)}, gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}))
}

func Call_sumEvens(n_0_loop int64) int64 {
var n_0 int64 = n_0_loop
_ = n_0
return Call_foldl(pkg_Data_Semiring.Get_intAdd(), gopurs_runtime.Int(0), gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[int64]](gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(Call_filterEvens(gopurs_runtime.CoerceToStruct[Constructor_Cons[int64]](gopurs_runtime.Apply2(Get_go__range(), gopurs_runtime.Int(1), gopurs_runtime.Int(n_0)))))}))})).IntVal
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

func Call_eq__2843686287(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Bool((__eta0_0.IntVal) == (__eta1_1.IntVal))
}

func Call_eq__2384498378(dict_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_mod__2185172824(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(pkg_Data_EuclideanRing.Get_intMod(), __eta0_0, __eta1_1)
}

func Call_mod__2579358968(dict_0_loop *pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_compare__669572705(dict_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_compare__821463600(dict_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
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

func Call_foldl__3262866295(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Cons[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 1127792131 && gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 1127792131 && gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
__t0 = Call_foldl(v_0, gopurs_runtime.Apply2(v_0, v1_1, (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0), (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V1)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_foldl__1255354935(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Cons[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 1127792131 && gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 1127792131 && gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
__t0 = Call_foldl(v_0, gopurs_runtime.Apply2(v_0, v1_1, (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0), (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V1)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}


