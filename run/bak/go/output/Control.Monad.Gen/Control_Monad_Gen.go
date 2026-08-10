package Control_Monad_Gen

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Control_Monad_Rec_Class "gopurs/output/Control.Monad.Rec.Class"
	pkg_Data_Semigroup_Last "gopurs/output/Data.Semigroup.Last"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	unsafe "unsafe"
)

var cache_lessThanOrEq gopurs_runtime.Value
var once_lessThanOrEq sync.Once
func Get_lessThanOrEq() gopurs_runtime.Value {
	once_lessThanOrEq.Do(func() {
		cache_lessThanOrEq = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})
_ = __local_var_0_0
return gopurs_runtime.Func(func(a1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Apply2(__local_var_0_0, a1_1, a2_2)
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 380165415) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Bool(true)
}
end_branch_1:
return gopurs_runtime.Bool((__t1.IntVal) != (0))
})
})
}()
	})
	return cache_lessThanOrEq
}

var cache_monoidAdditive gopurs_runtime.Value
var once_monoidAdditive sync.Once
func Get_monoidAdditive() gopurs_runtime.Value {
	once_monoidAdditive.Do(func() {
		cache_monoidAdditive = func() gopurs_runtime.Value {
semigroupAdditive1_0_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float((v_0.FloatVal()) + (v1_1.FloatVal()))
})
}))
_ = semigroupAdditive1_0_0
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupAdditive1_0_0
}), gopurs_runtime.Float(0.0))
}()
	})
	return cache_monoidAdditive
}

var cache_Cons gopurs_runtime.Value
var once_Cons sync.Once
func Get_Cons() gopurs_runtime.Value {
	once_Cons.Do(func() {
		cache_Cons = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, value0, (*Constructor_Cons[gopurs_runtime.Value])(value1.UnsafePtr)})}
})
})
	})
	return cache_Cons
}

var cache_Nil gopurs_runtime.Value
var once_Nil sync.Once
func Get_Nil() gopurs_runtime.Value {
	once_Nil.Do(func() {
		cache_Nil = gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: nil}
	})
	return cache_Nil
}

var cache_unfoldable gopurs_runtime.Value
var once_unfoldable sync.Once
func Get_unfoldable() gopurs_runtime.Value {
	once_unfoldable.Do(func() {
		cache_unfoldable = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfoldable(dictMonadRec_0_box, dictMonadGen_1_box)
})
	})
	return cache_unfoldable
}

var cache_unfoldable__gopurs_runtime_Value_542018773 gopurs_runtime.Value
var once_unfoldable__gopurs_runtime_Value_542018773 sync.Once
func Get_unfoldable__gopurs_runtime_Value_542018773() gopurs_runtime.Value {
	once_unfoldable__gopurs_runtime_Value_542018773.Do(func() {
		cache_unfoldable__gopurs_runtime_Value_542018773 = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfoldable__gopurs_runtime_Value_542018773(dictMonadRec_0_box, dictMonadGen_1_box)
})
	})
	return cache_unfoldable__gopurs_runtime_Value_542018773
}

var cache_semigroupFreqSemigroup gopurs_runtime.Value
var once_semigroupFreqSemigroup sync.Once
func Get_semigroupFreqSemigroup() gopurs_runtime.Value {
	once_semigroupFreqSemigroup.Do(func() {
		cache_semigroupFreqSemigroup = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(pos_2 gopurs_runtime.Value) gopurs_runtime.Value {
v2_3_0 := gopurs_runtime.Apply(v_0, pos_2)
_ = v2_3_0
var __t1 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_3_0.UnsafePtr).V0
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 930809136 && __t_tag_2.UnsafePtr != nil) {
__t1 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_3_0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
__t1 = v2_3_0
}
end_branch_1:
return __t1
})
})
}))
	})
	return cache_semigroupFreqSemigroup
}

var cache_fromIndex gopurs_runtime.Value
var once_fromIndex sync.Once
func Get_fromIndex() gopurs_runtime.Value {
	once_fromIndex.Do(func() {
		cache_fromIndex = gopurs_runtime.Func(func(dictFoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromIndex(dictFoldable1_0_box)
})
	})
	return cache_fromIndex
}

var cache_fromIndex__gopurs_runtime_Value_3111933544 gopurs_runtime.Value
var once_fromIndex__gopurs_runtime_Value_3111933544 sync.Once
func Get_fromIndex__gopurs_runtime_Value_3111933544() gopurs_runtime.Value {
	once_fromIndex__gopurs_runtime_Value_3111933544.Do(func() {
		cache_fromIndex__gopurs_runtime_Value_3111933544 = gopurs_runtime.Func(func(dictFoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromIndex__gopurs_runtime_Value_3111933544(dictFoldable1_0_box)
})
	})
	return cache_fromIndex__gopurs_runtime_Value_3111933544
}

var cache_oneOf gopurs_runtime.Value
var once_oneOf sync.Once
func Get_oneOf() gopurs_runtime.Value {
	once_oneOf.Do(func() {
		cache_oneOf = gopurs_runtime.Func2(func(dictMonadGen_0_box gopurs_runtime.Value, dictFoldable1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_oneOf(dictMonadGen_0_box, dictFoldable1_1_box)
})
	})
	return cache_oneOf
}

var cache_oneOf__gopurs_runtime_Value_2265861353 gopurs_runtime.Value
var once_oneOf__gopurs_runtime_Value_2265861353 sync.Once
func Get_oneOf__gopurs_runtime_Value_2265861353() gopurs_runtime.Value {
	once_oneOf__gopurs_runtime_Value_2265861353.Do(func() {
		cache_oneOf__gopurs_runtime_Value_2265861353 = gopurs_runtime.Func2(func(dictMonadGen_0_box gopurs_runtime.Value, dictFoldable1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_oneOf__gopurs_runtime_Value_2265861353(dictMonadGen_0_box, dictFoldable1_1_box)
})
	})
	return cache_oneOf__gopurs_runtime_Value_2265861353
}

var cache_freqSemigroup gopurs_runtime.Value
var once_freqSemigroup sync.Once
func Get_freqSemigroup() gopurs_runtime.Value {
	once_freqSemigroup.Do(func() {
		cache_freqSemigroup = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_freqSemigroup(v_0_box)
})
	})
	return cache_freqSemigroup
}

var cache_freqSemigroup__gopurs_runtime_Value_772548326 gopurs_runtime.Value
var once_freqSemigroup__gopurs_runtime_Value_772548326 sync.Once
func Get_freqSemigroup__gopurs_runtime_Value_772548326() gopurs_runtime.Value {
	once_freqSemigroup__gopurs_runtime_Value_772548326.Do(func() {
		cache_freqSemigroup__gopurs_runtime_Value_772548326 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_freqSemigroup__gopurs_runtime_Value_772548326(v_0_box)
})
	})
	return cache_freqSemigroup__gopurs_runtime_Value_772548326
}

var cache_frequency gopurs_runtime.Value
var once_frequency sync.Once
func Get_frequency() gopurs_runtime.Value {
	once_frequency.Do(func() {
		cache_frequency = gopurs_runtime.Func2(func(dictMonadGen_0_box gopurs_runtime.Value, dictFoldable1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_frequency(dictMonadGen_0_box, dictFoldable1_1_box)
})
	})
	return cache_frequency
}

var cache_filtered gopurs_runtime.Value
var once_filtered sync.Once
func Get_filtered() gopurs_runtime.Value {
	once_filtered.Do(func() {
		cache_filtered = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_filtered(dictMonadRec_0_box, dictMonadGen_1_box)
})
	})
	return cache_filtered
}

var cache_suchThat gopurs_runtime.Value
var once_suchThat sync.Once
func Get_suchThat() gopurs_runtime.Value {
	once_suchThat.Do(func() {
		cache_suchThat = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_suchThat(dictMonadRec_0_box, dictMonadGen_1_box)
})
	})
	return cache_suchThat
}

var cache_elements gopurs_runtime.Value
var once_elements sync.Once
func Get_elements() gopurs_runtime.Value {
	once_elements.Do(func() {
		cache_elements = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_elements(dictMonadGen_0_box)
})
	})
	return cache_elements
}

var cache_elements__gopurs_runtime_Value_1644372582 gopurs_runtime.Value
var once_elements__gopurs_runtime_Value_1644372582 sync.Once
func Get_elements__gopurs_runtime_Value_1644372582() gopurs_runtime.Value {
	once_elements__gopurs_runtime_Value_1644372582.Do(func() {
		cache_elements__gopurs_runtime_Value_1644372582 = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_elements__gopurs_runtime_Value_1644372582(dictMonadGen_0_box)
})
	})
	return cache_elements__gopurs_runtime_Value_1644372582
}

var cache_choose gopurs_runtime.Value
var once_choose sync.Once
func Get_choose() gopurs_runtime.Value {
	once_choose.Do(func() {
		cache_choose = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_choose(dictMonadGen_0_box)
})
	})
	return cache_choose
}

type Constructor_Cons[T_a any] struct {
	Rc uint32
	V0 T_a
	V1 *Constructor_Cons[gopurs_runtime.Value]
}


type Constructor_Nil[T_a any] struct {
	Rc uint32
}


func Call_unfoldable(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
_ = dictMonadGen_1
Monad0_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_2_0
pure_3_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_1
Bind1_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{})
_ = Bind1_4_2
return gopurs_runtime.Func(func(dictUnfoldable_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(gen_6 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "tailRecM"), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_lessThanOrEq(), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V1, gopurs_runtime.Int(0)).IntVal) != (0) {
__t7 = gopurs_runtime.Apply(pure_3_1, gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0})})
goto end_branch_7
} else {

}
}
{
__local_var_8_5 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0
_ = __local_var_8_5
__local_var_9_6 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V1
_ = __local_var_9_6
__t7 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_4_2, "bind"), gen_6, gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_1, gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, x_10, (*Constructor_Cons[gopurs_runtime.Value])(__local_var_8_5.UnsafePtr)})}, gopurs_runtime.Int((__local_var_9_6.IntVal) - (1))})}})})
}))
}
end_branch_7:
return __t7
}))
_ = __local_var_7_4
__local_var_8_8 := gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: nil})
_ = __local_var_8_8
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_4_2, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictUnfoldable_5, "unfoldr"), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 759514854 && v_7.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}.UnsafePtr))}
goto end_branch_3
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 759514854 && v_7.UnsafePtr != nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Cons[gopurs_runtime.Value])(v_7.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v_7.UnsafePtr).V1)}})}})}.UnsafePtr))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__t3.UnsafePtr))}
})), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "sized"), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_4, gopurs_runtime.Apply(__local_var_8_8, x_9))
})))
})
})
}

func Call_unfoldable__gopurs_runtime_Value_542018773(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
_ = dictMonadGen_1
Monad0_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_2_0
pure_3_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_1
Bind1_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{})
_ = Bind1_4_2
return gopurs_runtime.Func(func(dictUnfoldable_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(gen_6 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "tailRecM"), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_lessThanOrEq(), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V1, gopurs_runtime.Int(0)).IntVal) != (0) {
__t7 = gopurs_runtime.Apply(pure_3_1, gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0})})
goto end_branch_7
} else {

}
}
{
__local_var_8_5 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0
_ = __local_var_8_5
__local_var_9_6 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V1
_ = __local_var_9_6
__t7 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_4_2, "bind"), gen_6, gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_1, gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, x_10, (*Constructor_Cons[gopurs_runtime.Value])(__local_var_8_5.UnsafePtr)})}, gopurs_runtime.Int((__local_var_9_6.IntVal) - (1))})}})})
}))
}
end_branch_7:
return __t7
}))
_ = __local_var_7_4
__local_var_8_8 := gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: nil})
_ = __local_var_8_8
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_4_2, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictUnfoldable_5, "unfoldr"), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 759514854 && v_7.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}.UnsafePtr))}
goto end_branch_3
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 759514854 && v_7.UnsafePtr != nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Cons[gopurs_runtime.Value])(v_7.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v_7.UnsafePtr).V1)}})}})}.UnsafePtr))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__t3.UnsafePtr))}
})), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "sized"), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_4, gopurs_runtime.Apply(__local_var_8_8, x_9))
})))
})
})
}

func Call_fromIndex(dictFoldable1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable1_0 gopurs_runtime.Value = dictFoldable1_0_loop
_ = dictFoldable1_0
foldMap1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_0, "foldMap1"), pkg_Data_Semigroup_Last.Get_semigroupLast())
_ = foldMap1_1_0
return gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_1_0 gopurs_runtime.Value
go__go_4_1_0 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop gopurs_runtime.Value = v_5_loop_val
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_1_0:
for {
if false { continue go__go_4_1_0 }
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t2 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 759514854 && v1_6.UnsafePtr != nil) {
var __t3 gopurs_runtime.Value
{
var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)}
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 759514854 && __t_tag_4.UnsafePtr == nil) {
__t3 = (*Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V0
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Apply2(Get_lessThanOrEq(), v_5, gopurs_runtime.Int(0)).IntVal) != (0) {
__t3 = (*Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V0
goto end_branch_3
} else {

}
}
{
v_5_loop = gopurs_runtime.Int((v_5.IntVal) - (1))
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)}
continue go__go_4_1_0
__t3 = gopurs_runtime.Value{}
}
end_branch_3:
__t2 = __t3
goto end_branch_2
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 759514854 && v1_6.UnsafePtr == nil) {
__t2 = gopurs_runtime.Apply2(foldMap1_1_0, pkg_Data_Semigroup_Last.Get_Last(), xs_3)
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}
}()
})
})
return gopurs_runtime.Apply2(go__go_4_1_0, i_2, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_0, "Foldable0"), gopurs_runtime.Value{}), "foldr"), Get_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: nil}, xs_3))
})
})
}

func Call_fromIndex__gopurs_runtime_Value_3111933544(dictFoldable1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable1_0 gopurs_runtime.Value = dictFoldable1_0_loop
_ = dictFoldable1_0
foldMap1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_0, "foldMap1"), pkg_Data_Semigroup_Last.Get_semigroupLast())
_ = foldMap1_1_0
return gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_1_1 gopurs_runtime.Value
go__go_4_1_1 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop gopurs_runtime.Value = v_5_loop_val
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_1_1:
for {
if false { continue go__go_4_1_1 }
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t2 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 759514854 && v1_6.UnsafePtr != nil) {
var __t3 gopurs_runtime.Value
{
var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)}
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 759514854 && __t_tag_4.UnsafePtr == nil) {
__t3 = (*Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V0
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Apply2(Get_lessThanOrEq(), v_5, gopurs_runtime.Int(0)).IntVal) != (0) {
__t3 = (*Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V0
goto end_branch_3
} else {

}
}
{
v_5_loop = gopurs_runtime.Int((v_5.IntVal) - (1))
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)}
continue go__go_4_1_1
__t3 = gopurs_runtime.Value{}
}
end_branch_3:
__t2 = __t3
goto end_branch_2
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 759514854 && v1_6.UnsafePtr == nil) {
__t2 = gopurs_runtime.Apply2(foldMap1_1_0, pkg_Data_Semigroup_Last.Get_Last(), xs_3)
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}
}()
})
})
return gopurs_runtime.Apply2(go__go_4_1_1, i_2, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_0, "Foldable0"), gopurs_runtime.Value{}), "foldr"), Get_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: nil}, xs_3))
})
})
}

func Call_oneOf(dictMonadGen_0_loop gopurs_runtime.Value, dictFoldable1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
var dictFoldable1_1 gopurs_runtime.Value = dictFoldable1_1_loop
_ = dictFoldable1_1
length_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_1, "Foldable0"), gopurs_runtime.Value{}), "foldl"), gopurs_runtime.Func(func(c_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((1) + (c_2.IntVal))
})
}), gopurs_runtime.Int(0))
_ = length_2_0
fromIndex1_3_1 := Call_fromIndex(dictFoldable1_1)
_ = fromIndex1_3_1
return gopurs_runtime.Func(func(xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(0), gopurs_runtime.Int((gopurs_runtime.Apply(length_2_0, xs_4).IntVal) - (1))), gopurs_runtime.Func(func(n_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(fromIndex1_3_1, n_5, xs_4)
}))
})
}

func Call_oneOf__gopurs_runtime_Value_2265861353(dictMonadGen_0_loop gopurs_runtime.Value, dictFoldable1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
var dictFoldable1_1 gopurs_runtime.Value = dictFoldable1_1_loop
_ = dictFoldable1_1
length_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_1, "Foldable0"), gopurs_runtime.Value{}), "foldl"), gopurs_runtime.Func(func(c_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((1) + (c_2.IntVal))
})
}), gopurs_runtime.Int(0))
_ = length_2_0
fromIndex1_3_1 := Call_fromIndex(dictFoldable1_1)
_ = fromIndex1_3_1
return gopurs_runtime.Func(func(xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(0), gopurs_runtime.Int((gopurs_runtime.Apply(length_2_0, xs_4).IntVal) - (1))), gopurs_runtime.Func(func(n_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(fromIndex1_3_1, n_5, xs_4)
}))
})
}

func Call_freqSemigroup(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
__local_var_1_0 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0
_ = __local_var_1_0
__local_var_2_1 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1
_ = __local_var_2_1
return gopurs_runtime.Func(func(pos_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
var __t3 gopurs_runtime.Value
{
if (pos_3.FloatVal()) < (__local_var_1_0.FloatVal()) {
__t3 = gopurs_runtime.Bool(false)
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Bool(true)
}
end_branch_3:
if (__t3.IntVal) != (0) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Float((pos_3.FloatVal()) - (__local_var_1_0.FloatVal()))})}, __local_var_2_1})}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}, __local_var_2_1})}
}
end_branch_2:
return __t2
})
}

func Call_freqSemigroup__gopurs_runtime_Value_772548326(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
__local_var_1_0 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0
_ = __local_var_1_0
__local_var_2_1 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1
_ = __local_var_2_1
return gopurs_runtime.Func(func(pos_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
var __t3 gopurs_runtime.Value
{
if (pos_3.FloatVal()) < (__local_var_1_0.FloatVal()) {
__t3 = gopurs_runtime.Bool(false)
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Bool(true)
}
end_branch_3:
if (__t3.IntVal) != (0) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Float((pos_3.FloatVal()) - (__local_var_1_0.FloatVal()))})}, __local_var_2_1})}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}, __local_var_2_1})}
}
end_branch_2:
return __t2
})
}

func Call_frequency(dictMonadGen_0_loop gopurs_runtime.Value, dictFoldable1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
var dictFoldable1_1 gopurs_runtime.Value = dictFoldable1_1_loop
_ = dictFoldable1_1
foldMap_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_1, "Foldable0"), gopurs_runtime.Value{}), "foldMap"), Get_monoidAdditive())
_ = foldMap_2_0
foldMap1_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_1, "foldMap1"), Get_semigroupFreqSemigroup())
_ = foldMap1_3_1
return gopurs_runtime.Func(func(xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_2 := gopurs_runtime.Apply2(foldMap1_3_1, Get_freqSemigroup(), xs_4)
_ = __local_var_5_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseFloat"), gopurs_runtime.Float(0.0), gopurs_runtime.Apply2(foldMap_2_0, pkg_Data_Tuple.Get_fst(), xs_4)), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(__local_var_5_2, x_6).UnsafePtr).V1
}))
})
}

func Call_filtered(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
_ = dictMonadGen_1
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_0
return gopurs_runtime.Func(func(gen_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadRec_0, "tailRecM"), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_0, "map"), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (a_5.Type == 9 && a_5.IntVal == 930809136 && a_5.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, pkg_Data_Unit.Get_unit()})}
goto end_branch_1
} else {

}
}
{
if (a_5.Type == 9 && a_5.IntVal == 930809136 && a_5.UnsafePtr != nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(a_5.UnsafePtr).V0})}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}), gen_3)
}), pkg_Data_Unit.Get_unit())
})
}

func Call_suchThat(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
_ = dictMonadGen_1
filtered2_2_0 := Call_filtered(dictMonadRec_0, dictMonadGen_1)
_ = filtered2_2_0
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_1
return gopurs_runtime.Func(func(gen_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(pred_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(filtered2_2_0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_1, "map"), gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(pred_5, a_6).IntVal) != (0) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, a_6})}.UnsafePtr))}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__t2.UnsafePtr))}
}), gen_4))
})
})
}

func Call_elements(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
pure_2_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_2_1
return gopurs_runtime.Func(func(dictFoldable1_3 gopurs_runtime.Value) gopurs_runtime.Value {
length_4_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_3, "Foldable0"), gopurs_runtime.Value{}), "foldl"), gopurs_runtime.Func(func(c_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((1) + (c_4.IntVal))
})
}), gopurs_runtime.Int(0))
_ = length_4_2
fromIndex1_5_3 := Call_fromIndex(dictFoldable1_3)
_ = fromIndex1_5_3
return gopurs_runtime.Func(func(xs_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(0), gopurs_runtime.Int((gopurs_runtime.Apply(length_4_2, xs_6).IntVal) - (1))), gopurs_runtime.Func(func(n_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_2_1, gopurs_runtime.Apply2(fromIndex1_5_3, n_7, xs_6))
}))
})
})
}

func Call_elements__gopurs_runtime_Value_1644372582(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
pure_2_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_2_1
return gopurs_runtime.Func(func(dictFoldable1_3 gopurs_runtime.Value) gopurs_runtime.Value {
length_4_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_3, "Foldable0"), gopurs_runtime.Value{}), "foldl"), gopurs_runtime.Func(func(c_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((1) + (c_4.IntVal))
})
}), gopurs_runtime.Int(0))
_ = length_4_2
fromIndex1_5_3 := Call_fromIndex(dictFoldable1_3)
_ = fromIndex1_5_3
return gopurs_runtime.Func(func(xs_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(0), gopurs_runtime.Int((gopurs_runtime.Apply(length_4_2, xs_6).IntVal) - (1))), gopurs_runtime.Func(func(n_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_2_1, gopurs_runtime.Apply2(fromIndex1_5_3, n_7, xs_6))
}))
})
})
}

func Call_choose(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
chooseBool_1_0 := gopurs_runtime.RecordGet(dictMonadGen_0, "chooseBool")
_ = chooseBool_1_0
return gopurs_runtime.Func(func(genA_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(genB_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "bind"), chooseBool_1_0, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_4.IntVal) != (0) {
__t1 = genA_2
goto end_branch_1
} else {

}
}
{
__t1 = genB_3
}
end_branch_1:
return __t1
}))
})
})
}


