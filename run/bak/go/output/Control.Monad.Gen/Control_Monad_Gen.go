package Control_Monad_Gen

import (
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Control_Monad_Gen_Class "gopurs/output/Control.Monad.Gen.Class"
	pkg_Control_Monad_Rec_Class "gopurs/output/Control.Monad.Rec.Class"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Newtype "gopurs/output/Data.Newtype"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semigroup_Foldable "gopurs/output/Data.Semigroup.Foldable"
	pkg_Data_Semigroup_Last "gopurs/output/Data.Semigroup.Last"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unfoldable "gopurs/output/Data.Unfoldable"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

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
return gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, value0, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](value1)})}
})
})
	})
	return cache_Cons
}

var cache_Nil gopurs_runtime.Value
var once_Nil sync.Once
func Get_Nil() gopurs_runtime.Value {
	once_Nil.Do(func() {
		cache_Nil = gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}
	})
	return cache_Nil
}

var cache_unfoldable gopurs_runtime.Value
var once_unfoldable sync.Once
func Get_unfoldable() gopurs_runtime.Value {
	once_unfoldable.Do(func() {
		cache_unfoldable = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfoldable(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dictMonadGen_1_box))
})
	})
	return cache_unfoldable
}

var cache_semigroupFreqSemigroup gopurs_runtime.Value
var once_semigroupFreqSemigroup sync.Once
func Get_semigroupFreqSemigroup() gopurs_runtime.Value {
	once_semigroupFreqSemigroup.Do(func() {
		cache_semigroupFreqSemigroup = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(pos_2 gopurs_runtime.Value) gopurs_runtime.Value {
v2_3_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[*pkg_Data_Maybe.Constructor_Just[float64], gopurs_runtime.Value]](gopurs_runtime.Apply(v_0, pos_2))
_ = v2_3_0
var __t2 *pkg_Data_Tuple.Constructor_Tuple[*pkg_Data_Maybe.Constructor_Just[float64], gopurs_runtime.Value]
{
var __t_tag_1 gopurs_runtime.Value = (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v2_3_0)}.UnsafePtr).V0
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 930809136 && __t_tag_1.UnsafePtr != nil) {
__t2 = gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[*pkg_Data_Maybe.Constructor_Just[float64], gopurs_runtime.Value]](gopurs_runtime.Apply(v1_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v2_3_0)}.UnsafePtr).V0.UnsafePtr).V0))
goto end_branch_2
} else {

}
}
{
__t2 = v2_3_0
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(__t2)}
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
return Call_fromIndex(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box))
})
	})
	return cache_fromIndex
}

var cache_oneOf gopurs_runtime.Value
var once_oneOf sync.Once
func Get_oneOf() gopurs_runtime.Value {
	once_oneOf.Do(func() {
		cache_oneOf = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_oneOf(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dictMonadGen_0_box))
})
	})
	return cache_oneOf
}

var cache_freqSemigroup gopurs_runtime.Value
var once_freqSemigroup sync.Once
func Get_freqSemigroup() gopurs_runtime.Value {
	once_freqSemigroup.Do(func() {
		cache_freqSemigroup = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_freqSemigroup(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[float64, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_freqSemigroup
}

var cache_frequency gopurs_runtime.Value
var once_frequency sync.Once
func Get_frequency() gopurs_runtime.Value {
	once_frequency.Do(func() {
		cache_frequency = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_frequency(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dictMonadGen_0_box))
})
	})
	return cache_frequency
}

var cache_filtered gopurs_runtime.Value
var once_filtered sync.Once
func Get_filtered() gopurs_runtime.Value {
	once_filtered.Do(func() {
		cache_filtered = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_filtered(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dictMonadGen_1_box))
})
	})
	return cache_filtered
}

var cache_suchThat gopurs_runtime.Value
var once_suchThat sync.Once
func Get_suchThat() gopurs_runtime.Value {
	once_suchThat.Do(func() {
		cache_suchThat = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_suchThat(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dictMonadGen_1_box))
})
	})
	return cache_suchThat
}

var cache_elements gopurs_runtime.Value
var once_elements sync.Once
func Get_elements() gopurs_runtime.Value {
	once_elements.Do(func() {
		cache_elements = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_elements(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dictMonadGen_0_box))
})
	})
	return cache_elements
}

var cache_choose gopurs_runtime.Value
var once_choose sync.Once
func Get_choose() gopurs_runtime.Value {
	once_choose.Do(func() {
		cache_choose = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_choose(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dictMonadGen_0_box))
})
	})
	return cache_choose
}

var cache_bind__1858449959 gopurs_runtime.Value
var once_bind__1858449959 sync.Once
func Get_bind__1858449959() gopurs_runtime.Value {
	once_bind__1858449959.Do(func() {
		cache_bind__1858449959 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__1858449959(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__1858449959
}

var cache_bind__10508807 gopurs_runtime.Value
var once_bind__10508807 sync.Once
func Get_bind__10508807() gopurs_runtime.Value {
	once_bind__10508807.Do(func() {
		cache_bind__10508807 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__10508807(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__10508807
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

var cache_bind__1801470023 gopurs_runtime.Value
var once_bind__1801470023 sync.Once
func Get_bind__1801470023() gopurs_runtime.Value {
	once_bind__1801470023.Do(func() {
		cache_bind__1801470023 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__1801470023(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__1801470023
}

var cache_chooseFloat__1964853975 gopurs_runtime.Value
var once_chooseFloat__1964853975 sync.Once
func Get_chooseFloat__1964853975() gopurs_runtime.Value {
	once_chooseFloat__1964853975.Do(func() {
		cache_chooseFloat__1964853975 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_chooseFloat__1964853975(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_chooseFloat__1964853975
}

var cache_chooseInt__1063828903 gopurs_runtime.Value
var once_chooseInt__1063828903 sync.Once
func Get_chooseInt__1063828903() gopurs_runtime.Value {
	once_chooseInt__1063828903.Do(func() {
		cache_chooseInt__1063828903 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_chooseInt__1063828903(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_chooseInt__1063828903
}

var cache_sized__3147117623 gopurs_runtime.Value
var once_sized__3147117623 sync.Once
func Get_sized__3147117623() gopurs_runtime.Value {
	once_sized__3147117623.Do(func() {
		cache_sized__3147117623 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sized__3147117623(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sized__3147117623
}

var cache_sized__4206899191 gopurs_runtime.Value
var once_sized__4206899191 sync.Once
func Get_sized__4206899191() gopurs_runtime.Value {
	once_sized__4206899191.Do(func() {
		cache_sized__4206899191 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sized__4206899191(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sized__4206899191
}

var cache_freqSemigroup__772548326 gopurs_runtime.Value
var once_freqSemigroup__772548326 sync.Once
func Get_freqSemigroup__772548326() gopurs_runtime.Value {
	once_freqSemigroup__772548326.Do(func() {
		cache_freqSemigroup__772548326 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_freqSemigroup__772548326(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[float64, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_freqSemigroup__772548326
}

var cache_fromIndex__3111933544 gopurs_runtime.Value
var once_fromIndex__3111933544 sync.Once
func Get_fromIndex__3111933544() gopurs_runtime.Value {
	once_fromIndex__3111933544.Do(func() {
		cache_fromIndex__3111933544 = gopurs_runtime.Func(func(dictFoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromIndex__3111933544(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box))
})
	})
	return cache_fromIndex__3111933544
}

var cache_fromIndex__4031440680 gopurs_runtime.Value
var once_fromIndex__4031440680 sync.Once
func Get_fromIndex__4031440680() gopurs_runtime.Value {
	once_fromIndex__4031440680.Do(func() {
		cache_fromIndex__4031440680 = gopurs_runtime.Func(func(dictFoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromIndex__4031440680(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box))
})
	})
	return cache_fromIndex__4031440680
}

var cache_getFreqVal__2886942013 gopurs_runtime.Value
var once_getFreqVal__2886942013 sync.Once
func Get_getFreqVal__2886942013() gopurs_runtime.Value {
	once_getFreqVal__2886942013.Do(func() {
		cache_getFreqVal__2886942013 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_getFreqVal__2886942013(v_0_box, x_1_box.FloatVal())
})
	})
	return cache_getFreqVal__2886942013
}

var cache_getFreqVal__3389400221 gopurs_runtime.Value
var once_getFreqVal__3389400221 sync.Once
func Get_getFreqVal__3389400221() gopurs_runtime.Value {
	once_getFreqVal__3389400221.Do(func() {
		cache_getFreqVal__3389400221 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_getFreqVal__3389400221(v_0_box, x_1_box.FloatVal())
})
	})
	return cache_getFreqVal__3389400221
}

var cache_semigroupFreqSemigroup__3762318396 gopurs_runtime.Value
var once_semigroupFreqSemigroup__3762318396 sync.Once
func Get_semigroupFreqSemigroup__3762318396() gopurs_runtime.Value {
	once_semigroupFreqSemigroup__3762318396.Do(func() {
		cache_semigroupFreqSemigroup__3762318396 = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(pos_2 gopurs_runtime.Value) gopurs_runtime.Value {
v2_3_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[*pkg_Data_Maybe.Constructor_Just[float64], gopurs_runtime.Value]](gopurs_runtime.Apply(v_0, pos_2))
_ = v2_3_0
var __t2 *pkg_Data_Tuple.Constructor_Tuple[*pkg_Data_Maybe.Constructor_Just[float64], gopurs_runtime.Value]
{
var __t_tag_1 gopurs_runtime.Value = (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v2_3_0)}.UnsafePtr).V0
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 930809136 && __t_tag_1.UnsafePtr != nil) {
__t2 = gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[*pkg_Data_Maybe.Constructor_Just[float64], gopurs_runtime.Value]](gopurs_runtime.Apply(v1_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v2_3_0)}.UnsafePtr).V0.UnsafePtr).V0))
goto end_branch_2
} else {

}
}
{
__t2 = v2_3_0
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(__t2)}
})
})
}))
	})
	return cache_semigroupFreqSemigroup__3762318396
}

var cache_semigroupFreqSemigroup__4221302400 gopurs_runtime.Value
var once_semigroupFreqSemigroup__4221302400 sync.Once
func Get_semigroupFreqSemigroup__4221302400() gopurs_runtime.Value {
	once_semigroupFreqSemigroup__4221302400.Do(func() {
		cache_semigroupFreqSemigroup__4221302400 = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(pos_2 gopurs_runtime.Value) gopurs_runtime.Value {
v2_3_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[*pkg_Data_Maybe.Constructor_Just[float64], gopurs_runtime.Value]](gopurs_runtime.Apply(v_0, pos_2))
_ = v2_3_0
var __t2 *pkg_Data_Tuple.Constructor_Tuple[*pkg_Data_Maybe.Constructor_Just[float64], gopurs_runtime.Value]
{
var __t_tag_1 gopurs_runtime.Value = (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v2_3_0)}.UnsafePtr).V0
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 930809136 && __t_tag_1.UnsafePtr != nil) {
__t2 = gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[*pkg_Data_Maybe.Constructor_Just[float64], gopurs_runtime.Value]](gopurs_runtime.Apply(v1_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v2_3_0)}.UnsafePtr).V0.UnsafePtr).V0))
goto end_branch_2
} else {

}
}
{
__t2 = v2_3_0
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(__t2)}
})
})
}))
	})
	return cache_semigroupFreqSemigroup__4221302400
}

var cache_tailRecM__2220253896 gopurs_runtime.Value
var once_tailRecM__2220253896 sync.Once
func Get_tailRecM__2220253896() gopurs_runtime.Value {
	once_tailRecM__2220253896.Do(func() {
		cache_tailRecM__2220253896 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRecM__2220253896(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_tailRecM__2220253896
}

var cache_tailRecM__3865988408 gopurs_runtime.Value
var once_tailRecM__3865988408 sync.Once
func Get_tailRecM__3865988408() gopurs_runtime.Value {
	once_tailRecM__3865988408.Do(func() {
		cache_tailRecM__3865988408 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRecM__3865988408(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_tailRecM__3865988408
}

var cache_tailRecM__2063592441 gopurs_runtime.Value
var once_tailRecM__2063592441 sync.Once
func Get_tailRecM__2063592441() gopurs_runtime.Value {
	once_tailRecM__2063592441.Do(func() {
		cache_tailRecM__2063592441 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRecM__2063592441(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_tailRecM__2063592441
}

var cache_compose__858342840 gopurs_runtime.Value
var once_compose__858342840 sync.Once
func Get_compose__858342840() gopurs_runtime.Value {
	once_compose__858342840.Do(func() {
		cache_compose__858342840 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__858342840(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__858342840
}

var cache_semigroupoidFn__2387483462 gopurs_runtime.Value
var once_semigroupoidFn__2387483462 sync.Once
func Get_semigroupoidFn__2387483462() gopurs_runtime.Value {
	once_semigroupoidFn__2387483462.Do(func() {
		cache_semigroupoidFn__2387483462 = gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(g_1, x_2))
})
})
}))
	})
	return cache_semigroupoidFn__2387483462
}

var cache_foldl__2151204251 gopurs_runtime.Value
var once_foldl__2151204251 sync.Once
func Get_foldl__2151204251() gopurs_runtime.Value {
	once_foldl__2151204251.Do(func() {
		cache_foldl__2151204251 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__2151204251(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldl__2151204251
}

var cache_foldr__2151204251 gopurs_runtime.Value
var once_foldr__2151204251 sync.Once
func Get_foldr__2151204251() gopurs_runtime.Value {
	once_foldr__2151204251.Do(func() {
		cache_foldr__2151204251 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__2151204251(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldr__2151204251
}

var cache_foldr__3675782427 gopurs_runtime.Value
var once_foldr__3675782427 sync.Once
func Get_foldr__3675782427() gopurs_runtime.Value {
	once_foldr__3675782427.Do(func() {
		cache_foldr__3675782427 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__3675782427(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldr__3675782427
}

var cache_length__854370588 gopurs_runtime.Value
var once_length__854370588 sync.Once
func Get_length__854370588() gopurs_runtime.Value {
	once_length__854370588.Do(func() {
		cache_length__854370588 = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_length__854370588(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box))
})
	})
	return cache_length__854370588
}

var cache_length__1958096179 gopurs_runtime.Value
var once_length__1958096179 sync.Once
func Get_length__1958096179() gopurs_runtime.Value {
	once_length__1958096179.Do(func() {
		cache_length__1958096179 = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_length__1958096179(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box))
})
	})
	return cache_length__1958096179
}

var cache_length__949294460 gopurs_runtime.Value
var once_length__949294460 sync.Once
func Get_length__949294460() gopurs_runtime.Value {
	once_length__949294460.Do(func() {
		cache_length__949294460 = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictSemiring_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_length__949294460(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](dictSemiring_1_box))
})
	})
	return cache_length__949294460
}

var cache_map__2199395572 gopurs_runtime.Value
var once_map__2199395572 sync.Once
func Get_map__2199395572() gopurs_runtime.Value {
	once_map__2199395572.Do(func() {
		cache_map__2199395572 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2199395572(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2199395572
}

var cache_map__1504457012 gopurs_runtime.Value
var once_map__1504457012 sync.Once
func Get_map__1504457012() gopurs_runtime.Value {
	once_map__1504457012.Do(func() {
		cache_map__1504457012 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__1504457012(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__1504457012
}

var cache_mapFlipped__4215217780 gopurs_runtime.Value
var once_mapFlipped__4215217780 sync.Once
func Get_mapFlipped__4215217780() gopurs_runtime.Value {
	once_mapFlipped__4215217780.Do(func() {
		cache_mapFlipped__4215217780 = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, fa_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapFlipped__4215217780(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), fa_1_box, f_2_box)
})
	})
	return cache_mapFlipped__4215217780
}

var cache_mapFlipped__509401044 gopurs_runtime.Value
var once_mapFlipped__509401044 sync.Once
func Get_mapFlipped__509401044() gopurs_runtime.Value {
	once_mapFlipped__509401044.Do(func() {
		cache_mapFlipped__509401044 = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, fa_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapFlipped__509401044(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), fa_1_box, f_2_box)
})
	})
	return cache_mapFlipped__509401044
}

var cache_mapFlipped__3249733428 gopurs_runtime.Value
var once_mapFlipped__3249733428 sync.Once
func Get_mapFlipped__3249733428() gopurs_runtime.Value {
	once_mapFlipped__3249733428.Do(func() {
		cache_mapFlipped__3249733428 = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, fa_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapFlipped__3249733428(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), fa_1_box, f_2_box)
})
	})
	return cache_mapFlipped__3249733428
}

var cache_alaF__734584620 gopurs_runtime.Value
var once_alaF__734584620 sync.Once
func Get_alaF__734584620() gopurs_runtime.Value {
	once_alaF__734584620.Do(func() {
		cache_alaF__734584620 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alaF__734584620(v_0_box)
})
	})
	return cache_alaF__734584620
}

var cache_alaF__4085337484 gopurs_runtime.Value
var once_alaF__4085337484 sync.Once
func Get_alaF__4085337484() gopurs_runtime.Value {
	once_alaF__4085337484.Do(func() {
		cache_alaF__4085337484 = gopurs_runtime.Func5(func(_dollar__unused_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, _dollar__unused_2_box gopurs_runtime.Value, _dollar__unused_3_box gopurs_runtime.Value, v_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alaF__4085337484(_dollar__unused_0_box, _dollar__unused_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](_dollar__unused_2_box), gopurs_runtime.CoerceToStruct[pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](_dollar__unused_3_box), v_4_box)
})
	})
	return cache_alaF__4085337484
}

var cache_un__3470773042 gopurs_runtime.Value
var once_un__3470773042 sync.Once
func Get_un__3470773042() gopurs_runtime.Value {
	once_un__3470773042.Do(func() {
		cache_un__3470773042 = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_un__3470773042(gopurs_runtime.CoerceToStruct[pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](_dollar__unused_0_box), v_1_box)
})
	})
	return cache_un__3470773042
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

var cache_greaterThanOrEq__1061005983 gopurs_runtime.Value
var once_greaterThanOrEq__1061005983 sync.Once
func Get_greaterThanOrEq__1061005983() gopurs_runtime.Value {
	once_greaterThanOrEq__1061005983.Do(func() {
		cache_greaterThanOrEq__1061005983 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThanOrEq__1061005983(a1_0_box, a2_1_box))
})
	})
	return cache_greaterThanOrEq__1061005983
}

var cache_greaterThanOrEq__1409282474 gopurs_runtime.Value
var once_greaterThanOrEq__1409282474 sync.Once
func Get_greaterThanOrEq__1409282474() gopurs_runtime.Value {
	once_greaterThanOrEq__1409282474.Do(func() {
		cache_greaterThanOrEq__1409282474 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThanOrEq__1409282474(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_greaterThanOrEq__1409282474
}

var cache_greaterThanOrEq__2065354949 gopurs_runtime.Value
var once_greaterThanOrEq__2065354949 sync.Once
func Get_greaterThanOrEq__2065354949() gopurs_runtime.Value {
	once_greaterThanOrEq__2065354949.Do(func() {
		cache_greaterThanOrEq__2065354949 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThanOrEq__2065354949(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_greaterThanOrEq__2065354949
}

var cache_lessThanOrEq__4087042607 gopurs_runtime.Value
var once_lessThanOrEq__4087042607 sync.Once
func Get_lessThanOrEq__4087042607() gopurs_runtime.Value {
	once_lessThanOrEq__4087042607.Do(func() {
		cache_lessThanOrEq__4087042607 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThanOrEq__4087042607(a1_0_box, a2_1_box))
})
	})
	return cache_lessThanOrEq__4087042607
}

var cache_lessThanOrEq__1409282474 gopurs_runtime.Value
var once_lessThanOrEq__1409282474 sync.Once
func Get_lessThanOrEq__1409282474() gopurs_runtime.Value {
	once_lessThanOrEq__1409282474.Do(func() {
		cache_lessThanOrEq__1409282474 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThanOrEq__1409282474(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_lessThanOrEq__1409282474
}

var cache_sub__1043827704 gopurs_runtime.Value
var once_sub__1043827704 sync.Once
func Get_sub__1043827704() gopurs_runtime.Value {
	once_sub__1043827704.Do(func() {
		cache_sub__1043827704 = pkg_Data_Ring.Get_intSub()
	})
	return cache_sub__1043827704
}

var cache_sub__1135378904 gopurs_runtime.Value
var once_sub__1135378904 sync.Once
func Get_sub__1135378904() gopurs_runtime.Value {
	once_sub__1135378904.Do(func() {
		cache_sub__1135378904 = pkg_Data_Ring.Get_numSub()
	})
	return cache_sub__1135378904
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

var cache_sub__871462840 gopurs_runtime.Value
var once_sub__871462840 sync.Once
func Get_sub__871462840() gopurs_runtime.Value {
	once_sub__871462840.Do(func() {
		cache_sub__871462840 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__871462840(gopurs_runtime.CoerceToStruct[pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sub__871462840
}

var cache_foldMap1__3342855683 gopurs_runtime.Value
var once_foldMap1__3342855683 sync.Once
func Get_foldMap1__3342855683() gopurs_runtime.Value {
	once_foldMap1__3342855683.Do(func() {
		cache_foldMap1__3342855683 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMap1__3342855683(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldMap1__3342855683
}

var cache_foldMap1__4160988333 gopurs_runtime.Value
var once_foldMap1__4160988333 sync.Once
func Get_foldMap1__4160988333() gopurs_runtime.Value {
	once_foldMap1__4160988333.Do(func() {
		cache_foldMap1__4160988333 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMap1__4160988333(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldMap1__4160988333
}

var cache_semigroupLast__2108226578 gopurs_runtime.Value
var once_semigroupLast__2108226578 sync.Once
func Get_semigroupLast__2108226578() gopurs_runtime.Value {
	once_semigroupLast__2108226578.Do(func() {
		cache_semigroupLast__2108226578 = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
})
}))
	})
	return cache_semigroupLast__2108226578
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

var cache_one__1204848985 gopurs_runtime.Value
var once_one__1204848985 sync.Once
func Get_one__1204848985() gopurs_runtime.Value {
	once_one__1204848985.Do(func() {
		cache_one__1204848985 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_one__1204848985(dict_0_box)
})
	})
	return cache_one__1204848985
}

var cache_zero__1204848985 gopurs_runtime.Value
var once_zero__1204848985 sync.Once
func Get_zero__1204848985() gopurs_runtime.Value {
	once_zero__1204848985.Do(func() {
		cache_zero__1204848985 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zero__1204848985(dict_0_box)
})
	})
	return cache_zero__1204848985
}

var cache_fst__549384412 gopurs_runtime.Value
var once_fst__549384412 sync.Once
func Get_fst__549384412() gopurs_runtime.Value {
	once_fst__549384412.Do(func() {
		cache_fst__549384412 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(Call_fst__549384412(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[float64, gopurs_runtime.Value]](v_0_box)))
})
	})
	return cache_fst__549384412
}

var cache_snd__4227940231 gopurs_runtime.Value
var once_snd__4227940231 sync.Once
func Get_snd__4227940231() gopurs_runtime.Value {
	once_snd__4227940231.Do(func() {
		cache_snd__4227940231 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_snd__4227940231(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[*pkg_Data_Maybe.Constructor_Just[float64], gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_snd__4227940231
}

var cache_unfoldr__1128708256 gopurs_runtime.Value
var once_unfoldr__1128708256 sync.Once
func Get_unfoldr__1128708256() gopurs_runtime.Value {
	once_unfoldr__1128708256.Do(func() {
		cache_unfoldr__1128708256 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfoldr__1128708256(gopurs_runtime.CoerceToStruct[pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_unfoldr__1128708256
}

var cache_unfoldr__457386988 gopurs_runtime.Value
var once_unfoldr__457386988 sync.Once
func Get_unfoldr__457386988() gopurs_runtime.Value {
	once_unfoldr__457386988.Do(func() {
		cache_unfoldr__457386988 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfoldr__457386988(gopurs_runtime.CoerceToStruct[pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_unfoldr__457386988
}

type Constructor_Cons[T_a any] struct {
	Rc uint32
	V0 T_a
	V1 *Constructor_Cons[gopurs_runtime.Value]
}


type Constructor_Nil[T_a any] struct {
	Rc uint32
}


func Call_unfoldable(dictMonadRec_0_loop *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value], dictMonadGen_1_loop *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadRec_0 *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value] = dictMonadGen_1_loop
_ = dictMonadGen_1
Monad0_2_0 := gopurs_runtime.Apply(dictMonadGen_1.V0, gopurs_runtime.Value{})
_ = Monad0_2_0
pure_3_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_1
Bind1_4_2 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_2
Functor0_5_3 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_3
return gopurs_runtime.Func(func(dictUnfoldable_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(gen_7 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_8_5 := gopurs_runtime.Apply(dictMonadRec_0.V1, gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(Call_lessThanOrEq__4087042607((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V1, gopurs_runtime.Int(0))).IntVal) != (0) {
__t8 = gopurs_runtime.Apply(pure_3_1, gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0})})
goto end_branch_8
} else {

}
}
{
__local_var_9_6 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0
_ = __local_var_9_6
__local_var_10_7 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V1
_ = __local_var_10_7
__t8 = gopurs_runtime.Apply2(Bind1_4_2.V1, gen_7, gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_1, gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, x_11, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](__local_var_9_6)})}, gopurs_runtime.Apply2(Get_sub__1043827704(), __local_var_10_7, gopurs_runtime.Int(1))})}})})
}))
}
end_branch_8:
return __t8
}))
_ = __local_var_8_5
__local_var_9_9 := gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))})
_ = __local_var_9_9
return gopurs_runtime.Apply2(Functor0_5_3.V0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictUnfoldable_6, "unfoldr"), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (v_8.Type == 9 && v_8.IntVal == 759514854 && v_8.UnsafePtr == nil) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_4
} else {

}
}
{
if (v_8.Type == 9 && v_8.IntVal == 759514854 && v_8.UnsafePtr != nil) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Cons[gopurs_runtime.Value])(v_8.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v_8.UnsafePtr).V1)}})}})}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, *Constructor_Cons[gopurs_runtime.Value]]]](__t4))}
})), gopurs_runtime.Apply(dictMonadGen_1.V5, gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_5, gopurs_runtime.Apply(__local_var_9_9, x_10))
})))
})
})
}

func Call_fromIndex(dictFoldable1_0_loop *pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
Foldable0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](gopurs_runtime.Apply(dictFoldable1_0.V0, gopurs_runtime.Value{}))
_ = Foldable0_1_0
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
var __t4 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 759514854 && v1_6.UnsafePtr != nil) {
var __t3 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)}
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 759514854 && __t_tag_2.UnsafePtr == nil) {
__t3 = (*Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V0
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(Call_lessThanOrEq__4087042607(v_5, gopurs_runtime.Int(0))).IntVal) != (0) {
__t3 = (*Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V0
goto end_branch_3
} else {

}
}
{
v_5_loop = gopurs_runtime.Apply2(Get_sub__1043827704(), v_5, gopurs_runtime.Int(1))
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)}
continue go__go_4_1_0
__t3 = gopurs_runtime.Value{}
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 759514854 && v1_6.UnsafePtr == nil) {
__t4 = gopurs_runtime.Apply3(dictFoldable1_0.V1, pkg_Data_Semigroup_Last.Get_semigroupLast(), pkg_Data_Semigroup_Last.Get_Last(), xs_3)
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}
}()
})
})
return gopurs_runtime.Apply2(go__go_4_1_0, i_2, gopurs_runtime.Apply3(Foldable0_1_0.V2, Get_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}, xs_3))
})
})
}

func Call_oneOf(dictMonadGen_0_loop *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadGen_0 *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value] = dictMonadGen_0_loop
_ = dictMonadGen_0
Bind1_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonadGen_0.V0, gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_0
return gopurs_runtime.Func(func(dictFoldable1_2 gopurs_runtime.Value) gopurs_runtime.Value {
Foldable0_3_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_2, "Foldable0"), gopurs_runtime.Value{}))
_ = Foldable0_3_1
return gopurs_runtime.Func(func(xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_1_0.V1, gopurs_runtime.Apply2(dictMonadGen_0.V3, gopurs_runtime.Int(0), gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Apply3(Foldable0_3_1.V1, gopurs_runtime.Func(func(c_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((1) + (c_5.IntVal))
})
}), gopurs_runtime.Int(0), xs_4), gopurs_runtime.Int(1))), gopurs_runtime.Func(func(n_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_fromIndex(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value]](dictFoldable1_2)), n_5, xs_4)
}))
})
})
}

func Call_freqSemigroup(v_0_loop *pkg_Data_Tuple.Constructor_Tuple[float64, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *pkg_Data_Tuple.Constructor_Tuple[float64, gopurs_runtime.Value] = v_0_loop
_ = v_0
__local_var_1_0 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
_ = __local_var_1_0
__local_var_2_1 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1
_ = __local_var_2_1
return gopurs_runtime.Func(func(pos_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]
{
if (gopurs_runtime.Bool(Call_greaterThanOrEq__1061005983(pos_3, __local_var_1_0)).IntVal) != (0) {
__t2 = &pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply2(Get_sub__1135378904(), pos_3, __local_var_1_0)})}, __local_var_2_1}
goto end_branch_2
} else {

}
}
{
__t2 = &pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, __local_var_2_1}
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(__t2)}
})
}

func Call_frequency(dictMonadGen_0_loop *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadGen_0 *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value] = dictMonadGen_0_loop
_ = dictMonadGen_0
Bind1_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonadGen_0.V0, gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_0
return gopurs_runtime.Func(func(dictFoldable1_2 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_2, "Foldable0"), gopurs_runtime.Value{}), "foldMap"), Get_monoidAdditive())
_ = foldMap_3_1
return gopurs_runtime.Func(func(xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_2 := gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable1_2, "foldMap1"), Get_semigroupFreqSemigroup(), Get_freqSemigroup(), xs_4)
_ = __local_var_5_2
return gopurs_runtime.Apply2(Bind1_1_0.V1, gopurs_runtime.Apply2(dictMonadGen_0.V2, gopurs_runtime.Float(0.0), gopurs_runtime.Apply2(foldMap_3_1, pkg_Data_Tuple.Get_fst(), xs_4)), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(__local_var_5_2, x_6).UnsafePtr).V1
}))
})
})
}

func Call_filtered(dictMonadRec_0_loop *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value], dictMonadGen_1_loop *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadRec_0 *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value] = dictMonadGen_1_loop
_ = dictMonadGen_1
Functor0_2_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonadGen_1.V0, gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_0
return gopurs_runtime.Func(func(gen_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictMonadRec_0.V1, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_2_0.V0, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_suchThat(dictMonadRec_0_loop *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value], dictMonadGen_1_loop *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadRec_0 *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value] = dictMonadGen_1_loop
_ = dictMonadGen_1
filtered2_2_0 := Call_filtered(dictMonadRec_0, dictMonadGen_1)
_ = filtered2_2_0
Functor0_3_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonadGen_1.V0, gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_1
return gopurs_runtime.Func(func(gen_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(pred_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(filtered2_2_0, gopurs_runtime.Apply2(Functor0_3_1.V0, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(pred_5, a_6).IntVal) != (0) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, a_6})}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t2))}
}), gen_4))
})
})
}

func Call_elements(dictMonadGen_0_loop *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadGen_0 *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value] = dictMonadGen_0_loop
_ = dictMonadGen_0
Monad0_1_0 := gopurs_runtime.Apply(dictMonadGen_0.V0, gopurs_runtime.Value{})
_ = Monad0_1_0
Bind1_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
pure_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_2
return gopurs_runtime.Func(func(dictFoldable1_4 gopurs_runtime.Value) gopurs_runtime.Value {
Foldable0_5_3 := gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_4, "Foldable0"), gopurs_runtime.Value{}))
_ = Foldable0_5_3
return gopurs_runtime.Func(func(xs_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply2(dictMonadGen_0.V3, gopurs_runtime.Int(0), gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Apply3(Foldable0_5_3.V1, gopurs_runtime.Func(func(c_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((1) + (c_7.IntVal))
})
}), gopurs_runtime.Int(0), xs_6), gopurs_runtime.Int(1))), gopurs_runtime.Func(func(n_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_2, gopurs_runtime.Apply2(Call_fromIndex(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value]](dictFoldable1_4)), n_7, xs_6))
}))
})
})
}

func Call_choose(dictMonadGen_0_loop *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadGen_0 *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value] = dictMonadGen_0_loop
_ = dictMonadGen_0
Bind1_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonadGen_0.V0, gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_0
chooseBool_2_1 := dictMonadGen_0.V1
_ = chooseBool_2_1
return gopurs_runtime.Func(func(genA_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(genB_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_1_0.V1, chooseBool_2_1, gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_5.IntVal) != (0) {
__t2 = genA_3
goto end_branch_2
} else {

}
}
{
__t2 = genB_4
}
end_branch_2:
return __t2
}))
})
})
}

func Call_bind__1858449959(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__10508807(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__2601835655(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__1801470023(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_chooseFloat__1964853975(dict_0_loop *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_chooseInt__1063828903(dict_0_loop *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_sized__3147117623(dict_0_loop *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V5
}

func Call_sized__4206899191(dict_0_loop *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V5
}

func Call_freqSemigroup__772548326(v_0_loop *pkg_Data_Tuple.Constructor_Tuple[float64, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *pkg_Data_Tuple.Constructor_Tuple[float64, gopurs_runtime.Value] = v_0_loop
_ = v_0
__local_var_1_0 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
_ = __local_var_1_0
__local_var_2_1 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1
_ = __local_var_2_1
return gopurs_runtime.Func(func(pos_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]
{
if (gopurs_runtime.Bool(Call_greaterThanOrEq__1061005983(pos_3, __local_var_1_0)).IntVal) != (0) {
__t2 = &pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply2(Get_sub__1135378904(), pos_3, __local_var_1_0)})}, __local_var_2_1}
goto end_branch_2
} else {

}
}
{
__t2 = &pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, __local_var_2_1}
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(__t2)}
})
}

func Call_fromIndex__3111933544(dictFoldable1_0_loop *pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
Foldable0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](gopurs_runtime.Apply(dictFoldable1_0.V0, gopurs_runtime.Value{}))
_ = Foldable0_1_0
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
var __t4 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 759514854 && v1_6.UnsafePtr != nil) {
var __t3 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)}
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 759514854 && __t_tag_2.UnsafePtr == nil) {
__t3 = (*Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V0
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(Call_lessThanOrEq__4087042607(v_5, gopurs_runtime.Int(0))).IntVal) != (0) {
__t3 = (*Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V0
goto end_branch_3
} else {

}
}
{
v_5_loop = gopurs_runtime.Apply2(Get_sub__1043827704(), v_5, gopurs_runtime.Int(1))
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)}
continue go__go_4_1_1
__t3 = gopurs_runtime.Value{}
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 759514854 && v1_6.UnsafePtr == nil) {
__t4 = gopurs_runtime.Apply3(dictFoldable1_0.V1, pkg_Data_Semigroup_Last.Get_semigroupLast(), pkg_Data_Semigroup_Last.Get_Last(), xs_3)
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}
}()
})
})
return gopurs_runtime.Apply2(go__go_4_1_1, i_2, gopurs_runtime.Apply3(Foldable0_1_0.V2, Get_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}, xs_3))
})
})
}

func Call_fromIndex__4031440680(dictFoldable1_0_loop *pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
Foldable0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](gopurs_runtime.Apply(dictFoldable1_0.V0, gopurs_runtime.Value{}))
_ = Foldable0_1_0
return gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_1_2 gopurs_runtime.Value
go__go_4_1_2 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop gopurs_runtime.Value = v_5_loop_val
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_1_2:
for {
if false { continue go__go_4_1_2 }
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t4 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 759514854 && v1_6.UnsafePtr != nil) {
var __t3 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)}
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 759514854 && __t_tag_2.UnsafePtr == nil) {
__t3 = (*Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V0
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(Call_lessThanOrEq__4087042607(v_5, gopurs_runtime.Int(0))).IntVal) != (0) {
__t3 = (*Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V0
goto end_branch_3
} else {

}
}
{
v_5_loop = gopurs_runtime.Apply2(Get_sub__1043827704(), v_5, gopurs_runtime.Int(1))
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)}
continue go__go_4_1_2
__t3 = gopurs_runtime.Value{}
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 759514854 && v1_6.UnsafePtr == nil) {
__t4 = gopurs_runtime.Apply3(dictFoldable1_0.V1, pkg_Data_Semigroup_Last.Get_semigroupLast(), pkg_Data_Semigroup_Last.Get_Last(), xs_3)
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}
}()
})
})
return gopurs_runtime.Apply2(go__go_4_1_2, i_2, gopurs_runtime.Apply3(Foldable0_1_0.V2, Get_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}, xs_3))
})
})
}

func Call_getFreqVal__2886942013(v_0_loop gopurs_runtime.Value, x_1_loop float64) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var x_1 float64 = x_1_loop
_ = x_1
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(v_0, gopurs_runtime.Float(x_1)).UnsafePtr).V1
}

func Call_getFreqVal__3389400221(v_0_loop gopurs_runtime.Value, x_1_loop float64) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var x_1 float64 = x_1_loop
_ = x_1
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(v_0, gopurs_runtime.Float(x_1)).UnsafePtr).V1
}

func Call_tailRecM__2220253896(dict_0_loop *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_tailRecM__3865988408(dict_0_loop *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_tailRecM__2063592441(dict_0_loop *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_foldl__2151204251(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldr__2151204251(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_foldr__3675782427(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_length__854370588(dictFoldable_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
return gopurs_runtime.Apply2(dictFoldable_0.V1, gopurs_runtime.Func(func(c_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((1) + (c_1.IntVal))
})
}), gopurs_runtime.Int(0))
}

func Call_length__1958096179(dictFoldable_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
return gopurs_runtime.Apply2(dictFoldable_0.V1, gopurs_runtime.Func(func(c_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((1) + (c_1.IntVal))
})
}), gopurs_runtime.Int(0))
}

func Call_length__949294460(dictFoldable_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value], dictSemiring_1_loop *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictSemiring_1 *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value] = dictSemiring_1_loop
_ = dictSemiring_1
return gopurs_runtime.Apply2(dictFoldable_0.V1, gopurs_runtime.Func(func(c_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictSemiring_1.V0, dictSemiring_1.V2, c_2)
})
}), dictSemiring_1.V3)
}

func Call_map__2199395572(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__1504457012(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_mapFlipped__4215217780(dictFunctor_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value], fa_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var fa_1 gopurs_runtime.Value = fa_1_loop
_ = fa_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
return gopurs_runtime.Apply2(dictFunctor_0.V0, f_2, fa_1)
}

func Call_mapFlipped__509401044(dictFunctor_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value], fa_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var fa_1 gopurs_runtime.Value = fa_1_loop
_ = fa_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
return gopurs_runtime.Apply2(dictFunctor_0.V0, f_2, fa_1)
}

func Call_mapFlipped__3249733428(dictFunctor_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value], fa_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var fa_1 gopurs_runtime.Value = fa_1_loop
_ = fa_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
return gopurs_runtime.Apply2(dictFunctor_0.V0, f_2, fa_1)
}

func Call_alaF__734584620(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_alaF__4085337484(_dollar__unused_0_loop gopurs_runtime.Value, _dollar__unused_1_loop gopurs_runtime.Value, _dollar__unused_2_loop *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value], _dollar__unused_3_loop *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value], v_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var _dollar__unused_1 gopurs_runtime.Value = _dollar__unused_1_loop
_ = _dollar__unused_1
var _dollar__unused_2 *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar__unused_2_loop
_ = _dollar__unused_2
var _dollar__unused_3 *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar__unused_3_loop
_ = _dollar__unused_3
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_un__3470773042(_dollar__unused_0_loop *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value], v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar__unused_0_loop
_ = _dollar__unused_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_compare__821463600(dict_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_greaterThanOrEq__1061005983(a1_0_loop gopurs_runtime.Value, a2_1_loop gopurs_runtime.Value) bool {
var a1_0 gopurs_runtime.Value = a1_0_loop
_ = a1_0
var a2_1 gopurs_runtime.Value = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0.FloatVal()) < (a2_1.FloatVal()) {
__t0 = false
goto end_branch_0
} else {

}
}
{
__t0 = true
}
end_branch_0:
return __t0
}

func Call_greaterThanOrEq__1409282474(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
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
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = true
}
end_branch_1:
return __t1
}

func Call_greaterThanOrEq__2065354949(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
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
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = true
}
end_branch_1:
return __t1
}

func Call_lessThanOrEq__4087042607(a1_0_loop gopurs_runtime.Value, a2_1_loop gopurs_runtime.Value) bool {
var a1_0 gopurs_runtime.Value = a1_0_loop
_ = a1_0
var a2_1 gopurs_runtime.Value = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0.IntVal) > (a2_1.IntVal) {
__t0 = false
goto end_branch_0
} else {

}
}
{
__t0 = true
}
end_branch_0:
return __t0
}

func Call_lessThanOrEq__1409282474(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
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
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = true
}
end_branch_1:
return __t1
}

func Call_sub__3675938712(dict_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_sub__871462840(dict_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldMap1__3342855683(dict_0_loop *pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldMap1__4160988333(dict_0_loop *pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Apply(dict_0.V1, Get_semigroupFreqSemigroup())
}

func Call_add__1614463960(dict_0_loop *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_one__1204848985(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "one")
}

func Call_zero__1204848985(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "zero")
}

func Call_fst__549384412(v_0_loop *pkg_Data_Tuple.Constructor_Tuple[float64, gopurs_runtime.Value]) float64 {
var v_0 *pkg_Data_Tuple.Constructor_Tuple[float64, gopurs_runtime.Value] = v_0_loop
_ = v_0
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.FloatVal()
}

func Call_snd__4227940231(v_0_loop *pkg_Data_Tuple.Constructor_Tuple[*pkg_Data_Maybe.Constructor_Just[float64], gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *pkg_Data_Tuple.Constructor_Tuple[*pkg_Data_Maybe.Constructor_Just[float64], gopurs_runtime.Value] = v_0_loop
_ = v_0
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1
}

func Call_unfoldr__1128708256(dict_0_loop *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_unfoldr__457386988(dict_0_loop *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}


