package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_Gen_monoidAdditive gopurs_runtime.Value
var once_Control_Monad_Gen_monoidAdditive sync.Once
func Get_Control_Monad_Gen_monoidAdditive() gopurs_runtime.Value {
	once_Control_Monad_Gen_monoidAdditive.Do(func() {
		cache_Control_Monad_Gen_monoidAdditive = func() gopurs_runtime.Value {
// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
__local_var_0_0 := gopurs_runtime.RecordDict4("add", "mul", "one", "zero", Get_Data_Semiring_numAdd(), Get_Data_Semiring_numMul(), gopurs_runtime.Float(1.0), gopurs_runtime.Float(0.0))
_ = __local_var_0_0
// TAST (Let): semigroupAdditive1_1_1 -> gopurs_runtime.Value
semigroupAdditive1_1_1 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_0_0, "add"), v_1, v1_2)
})
}))
_ = semigroupAdditive1_1_1
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupAdditive1_1_1
}), gopurs_runtime.RecordGet(__local_var_0_0, "zero"))
}()
	})
	return cache_Control_Monad_Gen_monoidAdditive
}

var cache_Control_Monad_Gen_Cons gopurs_runtime.Value
var once_Control_Monad_Gen_Cons sync.Once
func Get_Control_Monad_Gen_Cons() gopurs_runtime.Value {
	once_Control_Monad_Gen_Cons.Do(func() {
		cache_Control_Monad_Gen_Cons = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Gen_Cons{1, value0, gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons](value1)})}
})
})
	})
	return cache_Control_Monad_Gen_Cons
}

var cache_Control_Monad_Gen_Nil gopurs_runtime.Value
var once_Control_Monad_Gen_Nil sync.Once
func Get_Control_Monad_Gen_Nil() gopurs_runtime.Value {
	once_Control_Monad_Gen_Nil.Do(func() {
		cache_Control_Monad_Gen_Nil = gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(nil))}
	})
	return cache_Control_Monad_Gen_Nil
}

var cache_Control_Monad_Gen_FreqSemigroup gopurs_runtime.Value
var once_Control_Monad_Gen_FreqSemigroup sync.Once
func Get_Control_Monad_Gen_FreqSemigroup() gopurs_runtime.Value {
	once_Control_Monad_Gen_FreqSemigroup.Do(func() {
		cache_Control_Monad_Gen_FreqSemigroup = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_FreqSemigroup(x_0_box)
})
	})
	return cache_Control_Monad_Gen_FreqSemigroup
}

var cache_Control_Monad_Gen_unfoldable gopurs_runtime.Value
var once_Control_Monad_Gen_unfoldable sync.Once
func Get_Control_Monad_Gen_unfoldable() gopurs_runtime.Value {
	once_Control_Monad_Gen_unfoldable.Do(func() {
		cache_Control_Monad_Gen_unfoldable = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_unfoldable(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dictMonadRec_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen](dictMonadGen_1_box))
})
	})
	return cache_Control_Monad_Gen_unfoldable
}

var cache_Control_Monad_Gen_semigroupFreqSemigroup gopurs_runtime.Value
var once_Control_Monad_Gen_semigroupFreqSemigroup sync.Once
func Get_Control_Monad_Gen_semigroupFreqSemigroup() gopurs_runtime.Value {
	once_Control_Monad_Gen_semigroupFreqSemigroup.Do(func() {
		cache_Control_Monad_Gen_semigroupFreqSemigroup = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(pos_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v2_3_0 -> *Constructor_Data_Tuple_Tuple
v2_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(v_0, gopurs_runtime.Float(pos_2.FloatVal())))
_ = v2_3_0
var __t2 *Constructor_Data_Tuple_Tuple
{
var __t_tag_1 gopurs_runtime.Value = (v2_3_0).V0
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 930809136 && __t_tag_1.UnsafePtr != nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(v1_1, gopurs_runtime.Float((*Constructor_Data_Maybe_Just)((v2_3_0).V0.UnsafePtr).V0.FloatVal())))
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
	return cache_Control_Monad_Gen_semigroupFreqSemigroup
}

var cache_Control_Monad_Gen_getFreqVal gopurs_runtime.Value
var once_Control_Monad_Gen_getFreqVal sync.Once
func Get_Control_Monad_Gen_getFreqVal() gopurs_runtime.Value {
	once_Control_Monad_Gen_getFreqVal.Do(func() {
		cache_Control_Monad_Gen_getFreqVal = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_getFreqVal(v_0_box, x_1_box.FloatVal())
})
	})
	return cache_Control_Monad_Gen_getFreqVal
}

var cache_Control_Monad_Gen_fromIndex gopurs_runtime.Value
var once_Control_Monad_Gen_fromIndex sync.Once
func Get_Control_Monad_Gen_fromIndex() gopurs_runtime.Value {
	once_Control_Monad_Gen_fromIndex.Do(func() {
		cache_Control_Monad_Gen_fromIndex = gopurs_runtime.Func(func(dictFoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_fromIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](dictFoldable1_0_box))
})
	})
	return cache_Control_Monad_Gen_fromIndex
}

var cache_Control_Monad_Gen_oneOf gopurs_runtime.Value
var once_Control_Monad_Gen_oneOf sync.Once
func Get_Control_Monad_Gen_oneOf() gopurs_runtime.Value {
	once_Control_Monad_Gen_oneOf.Do(func() {
		cache_Control_Monad_Gen_oneOf = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_oneOf(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen](dictMonadGen_0_box))
})
	})
	return cache_Control_Monad_Gen_oneOf
}

var cache_Control_Monad_Gen_freqSemigroup gopurs_runtime.Value
var once_Control_Monad_Gen_freqSemigroup sync.Once
func Get_Control_Monad_Gen_freqSemigroup() gopurs_runtime.Value {
	once_Control_Monad_Gen_freqSemigroup.Do(func() {
		cache_Control_Monad_Gen_freqSemigroup = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_freqSemigroup(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_0_box))
})
	})
	return cache_Control_Monad_Gen_freqSemigroup
}

var cache_Control_Monad_Gen_frequency gopurs_runtime.Value
var once_Control_Monad_Gen_frequency sync.Once
func Get_Control_Monad_Gen_frequency() gopurs_runtime.Value {
	once_Control_Monad_Gen_frequency.Do(func() {
		cache_Control_Monad_Gen_frequency = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_frequency(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen](dictMonadGen_0_box))
})
	})
	return cache_Control_Monad_Gen_frequency
}

var cache_Control_Monad_Gen_filtered gopurs_runtime.Value
var once_Control_Monad_Gen_filtered sync.Once
func Get_Control_Monad_Gen_filtered() gopurs_runtime.Value {
	once_Control_Monad_Gen_filtered.Do(func() {
		cache_Control_Monad_Gen_filtered = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_filtered(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dictMonadRec_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen](dictMonadGen_1_box))
})
	})
	return cache_Control_Monad_Gen_filtered
}

var cache_Control_Monad_Gen_suchThat gopurs_runtime.Value
var once_Control_Monad_Gen_suchThat sync.Once
func Get_Control_Monad_Gen_suchThat() gopurs_runtime.Value {
	once_Control_Monad_Gen_suchThat.Do(func() {
		cache_Control_Monad_Gen_suchThat = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_suchThat(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dictMonadRec_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen](dictMonadGen_1_box))
})
	})
	return cache_Control_Monad_Gen_suchThat
}

var cache_Control_Monad_Gen_elements gopurs_runtime.Value
var once_Control_Monad_Gen_elements sync.Once
func Get_Control_Monad_Gen_elements() gopurs_runtime.Value {
	once_Control_Monad_Gen_elements.Do(func() {
		cache_Control_Monad_Gen_elements = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_elements(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen](dictMonadGen_0_box))
})
	})
	return cache_Control_Monad_Gen_elements
}

var cache_Control_Monad_Gen_choose gopurs_runtime.Value
var once_Control_Monad_Gen_choose sync.Once
func Get_Control_Monad_Gen_choose() gopurs_runtime.Value {
	once_Control_Monad_Gen_choose.Do(func() {
		cache_Control_Monad_Gen_choose = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_choose(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen](dictMonadGen_0_box))
})
	})
	return cache_Control_Monad_Gen_choose
}

var cache_Control_Monad_Gen_elements__1644372582 gopurs_runtime.Value
var once_Control_Monad_Gen_elements__1644372582 sync.Once
func Get_Control_Monad_Gen_elements__1644372582() gopurs_runtime.Value {
	once_Control_Monad_Gen_elements__1644372582.Do(func() {
		cache_Control_Monad_Gen_elements__1644372582 = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_elements__1644372582(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen](dictMonadGen_0_box))
})
	})
	return cache_Control_Monad_Gen_elements__1644372582
}

var cache_Control_Monad_Gen_elements__2000025632 gopurs_runtime.Value
var once_Control_Monad_Gen_elements__2000025632 sync.Once
func Get_Control_Monad_Gen_elements__2000025632() gopurs_runtime.Value {
	once_Control_Monad_Gen_elements__2000025632.Do(func() {
		cache_Control_Monad_Gen_elements__2000025632 = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_elements__2000025632(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen](dictMonadGen_0_box))
})
	})
	return cache_Control_Monad_Gen_elements__2000025632
}

var cache_Control_Monad_Gen_freqSemigroup__772548326 gopurs_runtime.Value
var once_Control_Monad_Gen_freqSemigroup__772548326 sync.Once
func Get_Control_Monad_Gen_freqSemigroup__772548326() gopurs_runtime.Value {
	once_Control_Monad_Gen_freqSemigroup__772548326.Do(func() {
		cache_Control_Monad_Gen_freqSemigroup__772548326 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_freqSemigroup__772548326(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_0_box))
})
	})
	return cache_Control_Monad_Gen_freqSemigroup__772548326
}

var cache_Control_Monad_Gen_fromIndex__3111933544 gopurs_runtime.Value
var once_Control_Monad_Gen_fromIndex__3111933544 sync.Once
func Get_Control_Monad_Gen_fromIndex__3111933544() gopurs_runtime.Value {
	once_Control_Monad_Gen_fromIndex__3111933544.Do(func() {
		cache_Control_Monad_Gen_fromIndex__3111933544 = gopurs_runtime.Func(func(dictFoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_fromIndex__3111933544(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](dictFoldable1_0_box))
})
	})
	return cache_Control_Monad_Gen_fromIndex__3111933544
}

var cache_Control_Monad_Gen_fromIndex__4031440680 gopurs_runtime.Value
var once_Control_Monad_Gen_fromIndex__4031440680 sync.Once
func Get_Control_Monad_Gen_fromIndex__4031440680() gopurs_runtime.Value {
	once_Control_Monad_Gen_fromIndex__4031440680.Do(func() {
		cache_Control_Monad_Gen_fromIndex__4031440680 = gopurs_runtime.Func(func(dictFoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_fromIndex__4031440680(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](dictFoldable1_0_box))
})
	})
	return cache_Control_Monad_Gen_fromIndex__4031440680
}

var cache_Control_Monad_Gen_fromIndex__289821390 gopurs_runtime.Value
var once_Control_Monad_Gen_fromIndex__289821390 sync.Once
func Get_Control_Monad_Gen_fromIndex__289821390() gopurs_runtime.Value {
	once_Control_Monad_Gen_fromIndex__289821390.Do(func() {
		cache_Control_Monad_Gen_fromIndex__289821390 = gopurs_runtime.Func(func(dictFoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_fromIndex__289821390(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](dictFoldable1_0_box))
})
	})
	return cache_Control_Monad_Gen_fromIndex__289821390
}

var cache_Control_Monad_Gen_fromIndex__462059246 gopurs_runtime.Value
var once_Control_Monad_Gen_fromIndex__462059246 sync.Once
func Get_Control_Monad_Gen_fromIndex__462059246() gopurs_runtime.Value {
	once_Control_Monad_Gen_fromIndex__462059246.Do(func() {
		cache_Control_Monad_Gen_fromIndex__462059246 = gopurs_runtime.Func(func(dictFoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_fromIndex__462059246(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](dictFoldable1_0_box))
})
	})
	return cache_Control_Monad_Gen_fromIndex__462059246
}

var cache_Control_Monad_Gen_getFreqVal__2886942013 gopurs_runtime.Value
var once_Control_Monad_Gen_getFreqVal__2886942013 sync.Once
func Get_Control_Monad_Gen_getFreqVal__2886942013() gopurs_runtime.Value {
	once_Control_Monad_Gen_getFreqVal__2886942013.Do(func() {
		cache_Control_Monad_Gen_getFreqVal__2886942013 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_getFreqVal__2886942013(v_0_box, x_1_box.FloatVal())
})
	})
	return cache_Control_Monad_Gen_getFreqVal__2886942013
}

var cache_Control_Monad_Gen_getFreqVal__3389400221 gopurs_runtime.Value
var once_Control_Monad_Gen_getFreqVal__3389400221 sync.Once
func Get_Control_Monad_Gen_getFreqVal__3389400221() gopurs_runtime.Value {
	once_Control_Monad_Gen_getFreqVal__3389400221.Do(func() {
		cache_Control_Monad_Gen_getFreqVal__3389400221 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_getFreqVal__3389400221(v_0_box, x_1_box.FloatVal())
})
	})
	return cache_Control_Monad_Gen_getFreqVal__3389400221
}

var cache_Control_Monad_Gen_oneOf__2265861353 gopurs_runtime.Value
var once_Control_Monad_Gen_oneOf__2265861353 sync.Once
func Get_Control_Monad_Gen_oneOf__2265861353() gopurs_runtime.Value {
	once_Control_Monad_Gen_oneOf__2265861353.Do(func() {
		cache_Control_Monad_Gen_oneOf__2265861353 = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_oneOf__2265861353(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen](dictMonadGen_0_box))
})
	})
	return cache_Control_Monad_Gen_oneOf__2265861353
}

var cache_Control_Monad_Gen_oneOf__1433237231 gopurs_runtime.Value
var once_Control_Monad_Gen_oneOf__1433237231 sync.Once
func Get_Control_Monad_Gen_oneOf__1433237231() gopurs_runtime.Value {
	once_Control_Monad_Gen_oneOf__1433237231.Do(func() {
		cache_Control_Monad_Gen_oneOf__1433237231 = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_oneOf__1433237231(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen](dictMonadGen_0_box))
})
	})
	return cache_Control_Monad_Gen_oneOf__1433237231
}

var cache_Control_Monad_Gen_semigroupFreqSemigroup__3762318396 gopurs_runtime.Value
var once_Control_Monad_Gen_semigroupFreqSemigroup__3762318396 sync.Once
func Get_Control_Monad_Gen_semigroupFreqSemigroup__3762318396() gopurs_runtime.Value {
	once_Control_Monad_Gen_semigroupFreqSemigroup__3762318396.Do(func() {
		cache_Control_Monad_Gen_semigroupFreqSemigroup__3762318396 = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(pos_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v2_3_0 -> *Constructor_Data_Tuple_Tuple
v2_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(v_0, gopurs_runtime.Float(pos_2.FloatVal())))
_ = v2_3_0
var __t2 *Constructor_Data_Tuple_Tuple
{
var __t_tag_1 gopurs_runtime.Value = (v2_3_0).V0
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 930809136 && __t_tag_1.UnsafePtr != nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(v1_1, gopurs_runtime.Float((*Constructor_Data_Maybe_Just)((v2_3_0).V0.UnsafePtr).V0.FloatVal())))
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
	return cache_Control_Monad_Gen_semigroupFreqSemigroup__3762318396
}

var cache_Control_Monad_Gen_semigroupFreqSemigroup__4221302400 gopurs_runtime.Value
var once_Control_Monad_Gen_semigroupFreqSemigroup__4221302400 sync.Once
func Get_Control_Monad_Gen_semigroupFreqSemigroup__4221302400() gopurs_runtime.Value {
	once_Control_Monad_Gen_semigroupFreqSemigroup__4221302400.Do(func() {
		cache_Control_Monad_Gen_semigroupFreqSemigroup__4221302400 = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(pos_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v2_3_0 -> *Constructor_Data_Tuple_Tuple
v2_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(v_0, gopurs_runtime.Float(pos_2.FloatVal())))
_ = v2_3_0
var __t2 *Constructor_Data_Tuple_Tuple
{
var __t_tag_1 gopurs_runtime.Value = (v2_3_0).V0
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 930809136 && __t_tag_1.UnsafePtr != nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(v1_1, gopurs_runtime.Float((*Constructor_Data_Maybe_Just)((v2_3_0).V0.UnsafePtr).V0.FloatVal())))
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
	return cache_Control_Monad_Gen_semigroupFreqSemigroup__4221302400
}

var cache_Control_Monad_Gen_unfoldable__1024209525 gopurs_runtime.Value
var once_Control_Monad_Gen_unfoldable__1024209525 sync.Once
func Get_Control_Monad_Gen_unfoldable__1024209525() gopurs_runtime.Value {
	once_Control_Monad_Gen_unfoldable__1024209525.Do(func() {
		cache_Control_Monad_Gen_unfoldable__1024209525 = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_unfoldable__1024209525(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dictMonadRec_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen](dictMonadGen_1_box))
})
	})
	return cache_Control_Monad_Gen_unfoldable__1024209525
}

var cache_Control_Monad_Gen_unfoldable__542018773 gopurs_runtime.Value
var once_Control_Monad_Gen_unfoldable__542018773 sync.Once
func Get_Control_Monad_Gen_unfoldable__542018773() gopurs_runtime.Value {
	once_Control_Monad_Gen_unfoldable__542018773.Do(func() {
		cache_Control_Monad_Gen_unfoldable__542018773 = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_unfoldable__542018773(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dictMonadRec_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen](dictMonadGen_1_box))
})
	})
	return cache_Control_Monad_Gen_unfoldable__542018773
}

var cache_Control_Monad_Gen_unfoldable__4218841939 gopurs_runtime.Value
var once_Control_Monad_Gen_unfoldable__4218841939 sync.Once
func Get_Control_Monad_Gen_unfoldable__4218841939() gopurs_runtime.Value {
	once_Control_Monad_Gen_unfoldable__4218841939.Do(func() {
		cache_Control_Monad_Gen_unfoldable__4218841939 = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_unfoldable__4218841939(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dictMonadRec_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen](dictMonadGen_1_box))
})
	})
	return cache_Control_Monad_Gen_unfoldable__4218841939
}

type Constructor_Control_Monad_Gen_Cons struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 *Constructor_Control_Monad_Gen_Cons
}


type Constructor_Control_Monad_Gen_Nil struct {
	Rc uint32
}


func Call_Control_Monad_Gen_FreqSemigroup(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Monad_Gen_unfoldable(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec, dictMonadGen_1_loop *Constructor_Control_Monad_Gen_Class_MonadGen) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 *Constructor_Control_Monad_Gen_Class_MonadGen = dictMonadGen_1_loop
_ = dictMonadGen_1
// TAST (Let): Monad0_2_0 -> gopurs_runtime.Value
Monad0_2_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_1.V0), gopurs_runtime.Value{})
_ = Monad0_2_0
// TAST (Let): pure_3_1 -> gopurs_runtime.Value
pure_3_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_1
// TAST (Let): Bind1_4_2 -> *Constructor_Control_Bind_Bind
Bind1_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_2
// TAST (Let): Functor0_5_3 -> *Constructor_Data_Functor_Functor
Functor0_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_3
return gopurs_runtime.Func(func(dictUnfoldable_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(gen_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_5 -> gopurs_runtime.Value
__local_var_8_5 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadRec_0.V1), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 gopurs_runtime.Value
{
var __t8 bool
{
if ((*Constructor_Data_Tuple_Tuple)(v_8.UnsafePtr).V1.IntVal) > (0) {
__t8 = false
goto end_branch_8
} else {

}
}
{
__t8 = true
}
end_branch_8:
if __t8 {
__t9 = gopurs_runtime.Apply(pure_3_1, gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done{1, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons]((*Constructor_Data_Tuple_Tuple)(v_8.UnsafePtr).V0))}})})
goto end_branch_9
} else {

}
}
{
// TAST (Let): __local_var_9_6 -> gopurs_runtime.Value
__local_var_9_6 := (*Constructor_Data_Tuple_Tuple)(v_8.UnsafePtr).V0
_ = __local_var_9_6
// TAST (Let): __local_var_10_7 -> gopurs_runtime.Value
__local_var_10_7 := (*Constructor_Data_Tuple_Tuple)(v_8.UnsafePtr).V1
_ = __local_var_10_7
__t9 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_2.V1), gen_7, gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_1, gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Gen_Cons{1, x_11, gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons](__local_var_9_6)})}, gopurs_runtime.Int((__local_var_10_7.IntVal) - (1))})}})})
}))
}
end_branch_9:
return __t9
}))
_ = __local_var_8_5
// TAST (Let): __local_var_9_10 -> gopurs_runtime.Value
__local_var_9_10 := gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(nil))})
_ = __local_var_9_10
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_3.V0), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictUnfoldable_6, "unfoldr"), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 *Constructor_Data_Maybe_Just
{
if (v_8.Type == 9 && v_8.IntVal == 759514854 && v_8.UnsafePtr == nil) {
__t4 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_4
} else {

}
}
{
if (v_8.Type == 9 && v_8.IntVal == 759514854 && v_8.UnsafePtr != nil) {
__t4 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Control_Monad_Gen_Cons)(v_8.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(v_8.UnsafePtr).V1)}})}}
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t4)}
})), gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_1.V5), gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_5, gopurs_runtime.Apply(__local_var_9_10, x_10))
})))
})
})
}

func Call_Control_Monad_Gen_getFreqVal(v_0_loop gopurs_runtime.Value, x_1_loop float64) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var x_1 float64 = x_1_loop
_ = x_1
return (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply(v_0, gopurs_runtime.Float(x_1)).UnsafePtr).V1
}

func Call_Control_Monad_Gen_fromIndex(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1 = dictFoldable1_0_loop
_ = dictFoldable1_0
// TAST (Let): Foldable0_1_0 -> *Constructor_Data_Foldable_Foldable
Foldable0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](gopurs_runtime.Apply(gopurs_runtime.Box(dictFoldable1_0.V0), gopurs_runtime.Value{}))
_ = Foldable0_1_0
return gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_1_0 gopurs_runtime.Value
go__go_4_1_0 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop int64 = v_5_loop_val.IntVal
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_1_0:
for {
if false { continue go__go_4_1_0 }
var v_5 int64 = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t5 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 759514854 && v1_6.UnsafePtr != nil) {
var __t4 gopurs_runtime.Value
{
var __t_tag_2 *Constructor_Control_Monad_Gen_Cons = (*Constructor_Control_Monad_Gen_Cons)(v1_6.UnsafePtr).V1
if (__t_tag_2 == nil) {
__t4 = (*Constructor_Control_Monad_Gen_Cons)(v1_6.UnsafePtr).V0
goto end_branch_4
} else {

}
}
{
var __t3 bool
{
if (v_5) > (0) {
__t3 = false
goto end_branch_3
} else {

}
}
{
__t3 = true
}
end_branch_3:
if __t3 {
__t4 = (*Constructor_Control_Monad_Gen_Cons)(v1_6.UnsafePtr).V0
goto end_branch_4
} else {

}
}
{
v_5_loop = (v_5) - (1)
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(v1_6.UnsafePtr).V1)}
continue go__go_4_1_0
__t4 = gopurs_runtime.Value{}
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 759514854 && v1_6.UnsafePtr == nil) {
__t5 = gopurs_runtime.Apply3(gopurs_runtime.Box(dictFoldable1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_Semigroup_Last_semigroupLast()))}, Get_Data_Semigroup_Last_Last(), xs_3)
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}
}()
})
})
return gopurs_runtime.Apply2(go__go_4_1_0, gopurs_runtime.Int(i_2.IntVal), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons](gopurs_runtime.Apply3(gopurs_runtime.Box(Foldable0_1_0.V2), Get_Control_Monad_Gen_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(nil))}, xs_3)))})
})
})
}

func Call_Control_Monad_Gen_oneOf(dictMonadGen_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen) gopurs_runtime.Value {
var dictMonadGen_0 *Constructor_Control_Monad_Gen_Class_MonadGen = dictMonadGen_0_loop
_ = dictMonadGen_0
// TAST (Let): Bind1_1_0 -> *Constructor_Control_Bind_Bind
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_0.V0), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_0
return gopurs_runtime.Func(func(dictFoldable1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Foldable0_3_1 -> *Constructor_Data_Foldable_Foldable
Foldable0_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_2, "Foldable0"), gopurs_runtime.Value{}))
_ = Foldable0_3_1
return gopurs_runtime.Func(func(xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_2 -> *Constructor_Data_Semiring_Semiring
__local_var_5_2 := &Constructor_Data_Semiring_Semiring{1, Get_Data_Semiring_intAdd(), Get_Data_Semiring_intMul(), gopurs_runtime.Int(1), gopurs_runtime.Int(0)}
_ = __local_var_5_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadGen_0.V3), gopurs_runtime.Int(0), gopurs_runtime.Int((gopurs_runtime.Apply3(gopurs_runtime.Box(Foldable0_3_1.V1), gopurs_runtime.Func(func(c_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(__local_var_5_2.V0), gopurs_runtime.Box(__local_var_5_2.V2), c_6)
})
}), gopurs_runtime.Box(__local_var_5_2.V3), xs_4).IntVal) - (1))), gopurs_runtime.Func(func(n_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Monad_Gen_fromIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](dictFoldable1_2)), gopurs_runtime.Int(n_5.IntVal), xs_4)
}))
})
})
}

func Call_Control_Monad_Gen_freqSemigroup(v_0_loop *Constructor_Data_Tuple_Tuple) gopurs_runtime.Value {
var v_0 *Constructor_Data_Tuple_Tuple = v_0_loop
_ = v_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := (v_0).V0
_ = __local_var_1_0
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := (v_0).V1
_ = __local_var_2_1
return gopurs_runtime.Func(func(pos_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Tuple_Tuple
{
var __t2 bool
{
if (pos_3.FloatVal()) < (__local_var_1_0.FloatVal()) {
__t2 = false
goto end_branch_2
} else {

}
}
{
__t2 = true
}
end_branch_2:
if __t2 {
__t3 = &Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Float((pos_3.FloatVal()) - (__local_var_1_0.FloatVal()))})}, __local_var_2_1}
goto end_branch_3
} else {

}
}
{
__t3 = &Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, __local_var_2_1}
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(__t3)}
})
}

func Call_Control_Monad_Gen_frequency(dictMonadGen_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen) gopurs_runtime.Value {
var dictMonadGen_0 *Constructor_Control_Monad_Gen_Class_MonadGen = dictMonadGen_0_loop
_ = dictMonadGen_0
// TAST (Let): Bind1_1_0 -> *Constructor_Control_Bind_Bind
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_0.V0), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_0
return gopurs_runtime.Func(func(dictFoldable1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.RecordDict4("add", "mul", "one", "zero", Get_Data_Semiring_numAdd(), Get_Data_Semiring_numMul(), gopurs_runtime.Float(1.0), gopurs_runtime.Float(0.0))
_ = __local_var_3_2
// TAST (Let): semigroupAdditive1_4_3 -> gopurs_runtime.Value
semigroupAdditive1_4_3 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "add"), v_4, v1_5)
})
}))
_ = semigroupAdditive1_4_3
// TAST (Let): foldMap_3_1 -> gopurs_runtime.Value
foldMap_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_2, "Foldable0"), gopurs_runtime.Value{}), "foldMap"), gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupAdditive1_4_3
}), gopurs_runtime.RecordGet(__local_var_3_2, "zero")))
_ = foldMap_3_1
return gopurs_runtime.Func(func(xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_4 -> gopurs_runtime.Value
__local_var_5_4 := gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable1_2, "foldMap1"), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Control_Monad_Gen_semigroupFreqSemigroup()))}, Get_Control_Monad_Gen_freqSemigroup(), xs_4)
_ = __local_var_5_4
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadGen_0.V2), gopurs_runtime.Float(0.0), gopurs_runtime.Float(gopurs_runtime.Apply2(foldMap_3_1, Get_Data_Tuple_fst(), xs_4).FloatVal())), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply(__local_var_5_4, x_6).UnsafePtr).V1
}))
})
})
}

func Call_Control_Monad_Gen_filtered(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec, dictMonadGen_1_loop *Constructor_Control_Monad_Gen_Class_MonadGen) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 *Constructor_Control_Monad_Gen_Class_MonadGen = dictMonadGen_1_loop
_ = dictMonadGen_1
// TAST (Let): Functor0_2_0 -> *Constructor_Data_Functor_Functor
Functor0_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_1.V0), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_0
return gopurs_runtime.Func(func(gen_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadRec_0.V1), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_0.V0), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (a_5.Type == 9 && a_5.IntVal == 930809136 && a_5.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop{1, Get_Data_Unit_unit()})}
goto end_branch_1
} else {

}
}
{
if (a_5.Type == 9 && a_5.IntVal == 930809136 && a_5.UnsafePtr != nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done{1, (*Constructor_Data_Maybe_Just)(a_5.UnsafePtr).V0})}
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
}), Get_Data_Unit_unit())
})
}

func Call_Control_Monad_Gen_suchThat(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec, dictMonadGen_1_loop *Constructor_Control_Monad_Gen_Class_MonadGen) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 *Constructor_Control_Monad_Gen_Class_MonadGen = dictMonadGen_1_loop
_ = dictMonadGen_1
// TAST (Let): filtered2_2_0 -> gopurs_runtime.Value
filtered2_2_0 := Call_Control_Monad_Gen_filtered(dictMonadRec_0, dictMonadGen_1)
_ = filtered2_2_0
// TAST (Let): Functor0_3_1 -> *Constructor_Data_Functor_Functor
Functor0_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_1.V0), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_1
return gopurs_runtime.Func(func(gen_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(pred_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(filtered2_2_0, gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_1.V0), gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 *Constructor_Data_Maybe_Just
{
if (gopurs_runtime.Apply(pred_5, a_6).IntVal) != (0) {
__t2 = &Constructor_Data_Maybe_Just{1, a_6}
goto end_branch_2
} else {

}
}
{
__t2 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)}
}), gen_4))
})
})
}

func Call_Control_Monad_Gen_elements(dictMonadGen_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen) gopurs_runtime.Value {
var dictMonadGen_0 *Constructor_Control_Monad_Gen_Class_MonadGen = dictMonadGen_0_loop
_ = dictMonadGen_0
// TAST (Let): Monad0_1_0 -> gopurs_runtime.Value
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_0.V0), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): Bind1_2_1 -> *Constructor_Control_Bind_Bind
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
// TAST (Let): pure_3_2 -> gopurs_runtime.Value
pure_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_2
return gopurs_runtime.Func(func(dictFoldable1_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Foldable0_5_3 -> *Constructor_Data_Foldable_Foldable
Foldable0_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_4, "Foldable0"), gopurs_runtime.Value{}))
_ = Foldable0_5_3
return gopurs_runtime.Func(func(xs_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_4 -> *Constructor_Data_Semiring_Semiring
__local_var_7_4 := &Constructor_Data_Semiring_Semiring{1, Get_Data_Semiring_intAdd(), Get_Data_Semiring_intMul(), gopurs_runtime.Int(1), gopurs_runtime.Int(0)}
_ = __local_var_7_4
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadGen_0.V3), gopurs_runtime.Int(0), gopurs_runtime.Int((gopurs_runtime.Apply3(gopurs_runtime.Box(Foldable0_5_3.V1), gopurs_runtime.Func(func(c_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(__local_var_7_4.V0), gopurs_runtime.Box(__local_var_7_4.V2), c_8)
})
}), gopurs_runtime.Box(__local_var_7_4.V3), xs_6).IntVal) - (1))), gopurs_runtime.Func(func(n_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_2, gopurs_runtime.Apply2(Call_Control_Monad_Gen_fromIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](dictFoldable1_4)), gopurs_runtime.Int(n_7.IntVal), xs_6))
}))
})
})
}

func Call_Control_Monad_Gen_choose(dictMonadGen_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen) gopurs_runtime.Value {
var dictMonadGen_0 *Constructor_Control_Monad_Gen_Class_MonadGen = dictMonadGen_0_loop
_ = dictMonadGen_0
// TAST (Let): Bind1_1_0 -> *Constructor_Control_Bind_Bind
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_0.V0), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_0
// TAST (Let): chooseBool_2_1 -> gopurs_runtime.Value
chooseBool_2_1 := gopurs_runtime.Box(dictMonadGen_0.V1)
_ = chooseBool_2_1
return gopurs_runtime.Func(func(genA_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(genB_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), chooseBool_2_1, gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_Control_Monad_Gen_elements__1644372582(dictMonadGen_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen) gopurs_runtime.Value {
var dictMonadGen_0 *Constructor_Control_Monad_Gen_Class_MonadGen = dictMonadGen_0_loop
_ = dictMonadGen_0
// TAST (Let): Monad0_1_0 -> gopurs_runtime.Value
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_0.V0), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): Bind1_2_1 -> *Constructor_Control_Bind_Bind
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
// TAST (Let): pure_3_2 -> gopurs_runtime.Value
pure_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_2
return gopurs_runtime.Func(func(dictFoldable1_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Foldable0_5_3 -> *Constructor_Data_Foldable_Foldable
Foldable0_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_4, "Foldable0"), gopurs_runtime.Value{}))
_ = Foldable0_5_3
return gopurs_runtime.Func(func(xs_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_4 -> *Constructor_Data_Semiring_Semiring
__local_var_7_4 := &Constructor_Data_Semiring_Semiring{1, Get_Data_Semiring_intAdd(), Get_Data_Semiring_intMul(), gopurs_runtime.Int(1), gopurs_runtime.Int(0)}
_ = __local_var_7_4
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadGen_0.V3), gopurs_runtime.Int(0), gopurs_runtime.Int((gopurs_runtime.Apply3(gopurs_runtime.Box(Foldable0_5_3.V1), gopurs_runtime.Func(func(c_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(__local_var_7_4.V0), gopurs_runtime.Box(__local_var_7_4.V2), c_8)
})
}), gopurs_runtime.Box(__local_var_7_4.V3), xs_6).IntVal) - (1))), gopurs_runtime.Func(func(n_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_2, gopurs_runtime.Apply2(Call_Control_Monad_Gen_fromIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](dictFoldable1_4)), gopurs_runtime.Int(n_7.IntVal), xs_6))
}))
})
})
}

func Call_Control_Monad_Gen_elements__2000025632(dictMonadGen_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen) gopurs_runtime.Value {
var dictMonadGen_0 *Constructor_Control_Monad_Gen_Class_MonadGen = dictMonadGen_0_loop
_ = dictMonadGen_0
// TAST (Let): Monad0_1_0 -> gopurs_runtime.Value
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_0.V0), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): Bind1_2_1 -> *Constructor_Control_Bind_Bind
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
// TAST (Let): pure_3_2 -> gopurs_runtime.Value
pure_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_2
// TAST (Let): Foldable0_4_3 -> *Constructor_Data_Foldable_Foldable
Foldable0_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](Get_Data_Enum_Gen_foldable1NonEmpty()).V0), gopurs_runtime.Value{}))
_ = Foldable0_4_3
return gopurs_runtime.Func(func(xs_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_4 -> *Constructor_Data_Semiring_Semiring
__local_var_6_4 := &Constructor_Data_Semiring_Semiring{1, Get_Data_Semiring_intAdd(), Get_Data_Semiring_intMul(), gopurs_runtime.Int(1), gopurs_runtime.Int(0)}
_ = __local_var_6_4
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadGen_0.V3), gopurs_runtime.Int(0), gopurs_runtime.Int((gopurs_runtime.Apply3(gopurs_runtime.Box(Foldable0_4_3.V1), gopurs_runtime.Func(func(c_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(__local_var_6_4.V0), gopurs_runtime.Box(__local_var_6_4.V2), c_7)
})
}), gopurs_runtime.Box(__local_var_6_4.V3), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](xs_5))}).IntVal) - (1))), gopurs_runtime.Func(func(n_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_2, gopurs_runtime.Apply2(Call_Control_Monad_Gen_fromIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](Get_Data_Enum_Gen_foldable1NonEmpty()))})), gopurs_runtime.Int(n_6.IntVal), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](xs_5))}))
}))
})
}

func Call_Control_Monad_Gen_freqSemigroup__772548326(v_0_loop *Constructor_Data_Tuple_Tuple) gopurs_runtime.Value {
var v_0 *Constructor_Data_Tuple_Tuple = v_0_loop
_ = v_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := (v_0).V0
_ = __local_var_1_0
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := (v_0).V1
_ = __local_var_2_1
return gopurs_runtime.Func(func(pos_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Tuple_Tuple
{
var __t2 bool
{
if (pos_3.FloatVal()) < (__local_var_1_0.FloatVal()) {
__t2 = false
goto end_branch_2
} else {

}
}
{
__t2 = true
}
end_branch_2:
if __t2 {
__t3 = &Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Float((pos_3.FloatVal()) - (__local_var_1_0.FloatVal()))})}, __local_var_2_1}
goto end_branch_3
} else {

}
}
{
__t3 = &Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, __local_var_2_1}
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(__t3)}
})
}

func Call_Control_Monad_Gen_fromIndex__3111933544(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1 = dictFoldable1_0_loop
_ = dictFoldable1_0
// TAST (Let): Foldable0_1_0 -> *Constructor_Data_Foldable_Foldable
Foldable0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](gopurs_runtime.Apply(gopurs_runtime.Box(dictFoldable1_0.V0), gopurs_runtime.Value{}))
_ = Foldable0_1_0
return gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_1_1 gopurs_runtime.Value
go__go_4_1_1 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop int64 = v_5_loop_val.IntVal
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_1_1:
for {
if false { continue go__go_4_1_1 }
var v_5 int64 = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t5 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 759514854 && v1_6.UnsafePtr != nil) {
var __t4 gopurs_runtime.Value
{
var __t_tag_2 *Constructor_Control_Monad_Gen_Cons = (*Constructor_Control_Monad_Gen_Cons)(v1_6.UnsafePtr).V1
if (__t_tag_2 == nil) {
__t4 = (*Constructor_Control_Monad_Gen_Cons)(v1_6.UnsafePtr).V0
goto end_branch_4
} else {

}
}
{
var __t3 bool
{
if (v_5) > (0) {
__t3 = false
goto end_branch_3
} else {

}
}
{
__t3 = true
}
end_branch_3:
if __t3 {
__t4 = (*Constructor_Control_Monad_Gen_Cons)(v1_6.UnsafePtr).V0
goto end_branch_4
} else {

}
}
{
v_5_loop = (v_5) - (1)
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(v1_6.UnsafePtr).V1)}
continue go__go_4_1_1
__t4 = gopurs_runtime.Value{}
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 759514854 && v1_6.UnsafePtr == nil) {
__t5 = gopurs_runtime.Apply3(gopurs_runtime.Box(dictFoldable1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_Semigroup_Last_semigroupLast()))}, Get_Data_Semigroup_Last_Last(), xs_3)
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}
}()
})
})
return gopurs_runtime.Apply2(go__go_4_1_1, gopurs_runtime.Int(i_2.IntVal), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons](gopurs_runtime.Apply3(gopurs_runtime.Box(Foldable0_1_0.V2), Get_Control_Monad_Gen_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(nil))}, xs_3)))})
})
})
}

func Call_Control_Monad_Gen_fromIndex__4031440680(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1 = dictFoldable1_0_loop
_ = dictFoldable1_0
// TAST (Let): Foldable0_1_0 -> *Constructor_Data_Foldable_Foldable
Foldable0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](gopurs_runtime.Apply(gopurs_runtime.Box(dictFoldable1_0.V0), gopurs_runtime.Value{}))
_ = Foldable0_1_0
return gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_1_2 gopurs_runtime.Value
go__go_4_1_2 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop int64 = v_5_loop_val.IntVal
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_1_2:
for {
if false { continue go__go_4_1_2 }
var v_5 int64 = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t5 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 759514854 && v1_6.UnsafePtr != nil) {
var __t4 gopurs_runtime.Value
{
var __t_tag_2 *Constructor_Control_Monad_Gen_Cons = (*Constructor_Control_Monad_Gen_Cons)(v1_6.UnsafePtr).V1
if (__t_tag_2 == nil) {
__t4 = (*Constructor_Control_Monad_Gen_Cons)(v1_6.UnsafePtr).V0
goto end_branch_4
} else {

}
}
{
var __t3 bool
{
if (v_5) > (0) {
__t3 = false
goto end_branch_3
} else {

}
}
{
__t3 = true
}
end_branch_3:
if __t3 {
__t4 = (*Constructor_Control_Monad_Gen_Cons)(v1_6.UnsafePtr).V0
goto end_branch_4
} else {

}
}
{
v_5_loop = (v_5) - (1)
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(v1_6.UnsafePtr).V1)}
continue go__go_4_1_2
__t4 = gopurs_runtime.Value{}
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 759514854 && v1_6.UnsafePtr == nil) {
__t5 = gopurs_runtime.Apply3(gopurs_runtime.Box(dictFoldable1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_Semigroup_Last_semigroupLast()))}, Get_Data_Semigroup_Last_Last(), xs_3)
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}
}()
})
})
return gopurs_runtime.Apply2(go__go_4_1_2, gopurs_runtime.Int(i_2.IntVal), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons](gopurs_runtime.Apply3(gopurs_runtime.Box(Foldable0_1_0.V2), Get_Control_Monad_Gen_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(nil))}, xs_3)))})
})
})
}

func Call_Control_Monad_Gen_fromIndex__289821390(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1 = dictFoldable1_0_loop
_ = dictFoldable1_0
// TAST (Let): Foldable0_1_0 -> *Constructor_Data_Foldable_Foldable
Foldable0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](gopurs_runtime.Apply(gopurs_runtime.Box(dictFoldable1_0.V0), gopurs_runtime.Value{}))
_ = Foldable0_1_0
return gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_1_3 gopurs_runtime.Value
go__go_4_1_3 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop int64 = v_5_loop_val.IntVal
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_1_3:
for {
if false { continue go__go_4_1_3 }
var v_5 int64 = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t5 string
{
if (v1_6.Type == 9 && v1_6.IntVal == 759514854 && v1_6.UnsafePtr != nil) {
var __t4 string
{
var __t_tag_2 *Constructor_Control_Monad_Gen_Cons = (*Constructor_Control_Monad_Gen_Cons)(v1_6.UnsafePtr).V1
if (__t_tag_2 == nil) {
__t4 = (*Constructor_Control_Monad_Gen_Cons)(v1_6.UnsafePtr).V0.StrVal()
goto end_branch_4
} else {

}
}
{
var __t3 bool
{
if (v_5) > (0) {
__t3 = false
goto end_branch_3
} else {

}
}
{
__t3 = true
}
end_branch_3:
if __t3 {
__t4 = (*Constructor_Control_Monad_Gen_Cons)(v1_6.UnsafePtr).V0.StrVal()
goto end_branch_4
} else {

}
}
{
v_5_loop = (v_5) - (1)
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(v1_6.UnsafePtr).V1)}
continue go__go_4_1_3
__t4 = gopurs_runtime.Value{}.StrVal()
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 759514854 && v1_6.UnsafePtr == nil) {
__t5 = gopurs_runtime.Apply3(gopurs_runtime.Box(dictFoldable1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_Semigroup_Last_semigroupLast()))}, Get_Data_Semigroup_Last_Last(), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](xs_3))}).StrVal()
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
}
end_branch_5:
return gopurs_runtime.Str(__t5)
}
}()
})
})
return gopurs_runtime.Str(gopurs_runtime.Apply2(go__go_4_1_3, gopurs_runtime.Int(i_2.IntVal), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons](gopurs_runtime.Apply3(gopurs_runtime.Box(Foldable0_1_0.V2), Get_Control_Monad_Gen_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](xs_3))})))}).StrVal())
})
})
}

func Call_Control_Monad_Gen_fromIndex__462059246(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1 = dictFoldable1_0_loop
_ = dictFoldable1_0
// TAST (Let): Foldable0_1_0 -> *Constructor_Data_Foldable_Foldable
Foldable0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](gopurs_runtime.Apply(gopurs_runtime.Box(dictFoldable1_0.V0), gopurs_runtime.Value{}))
_ = Foldable0_1_0
return gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_1_4 gopurs_runtime.Value
go__go_4_1_4 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop int64 = v_5_loop_val.IntVal
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_1_4:
for {
if false { continue go__go_4_1_4 }
var v_5 int64 = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t5 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 759514854 && v1_6.UnsafePtr != nil) {
var __t4 gopurs_runtime.Value
{
var __t_tag_2 *Constructor_Control_Monad_Gen_Cons = (*Constructor_Control_Monad_Gen_Cons)(v1_6.UnsafePtr).V1
if (__t_tag_2 == nil) {
__t4 = (*Constructor_Control_Monad_Gen_Cons)(v1_6.UnsafePtr).V0
goto end_branch_4
} else {

}
}
{
var __t3 bool
{
if (v_5) > (0) {
__t3 = false
goto end_branch_3
} else {

}
}
{
__t3 = true
}
end_branch_3:
if __t3 {
__t4 = (*Constructor_Control_Monad_Gen_Cons)(v1_6.UnsafePtr).V0
goto end_branch_4
} else {

}
}
{
v_5_loop = (v_5) - (1)
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(v1_6.UnsafePtr).V1)}
continue go__go_4_1_4
__t4 = gopurs_runtime.Value{}
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 759514854 && v1_6.UnsafePtr == nil) {
__t5 = gopurs_runtime.Apply3(gopurs_runtime.Box(dictFoldable1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_Semigroup_Last_semigroupLast()))}, Get_Data_Semigroup_Last_Last(), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](xs_3))})
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}
}()
})
})
return gopurs_runtime.Apply2(go__go_4_1_4, gopurs_runtime.Int(i_2.IntVal), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons](gopurs_runtime.Apply3(gopurs_runtime.Box(Foldable0_1_0.V2), Get_Control_Monad_Gen_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](xs_3))})))})
})
})
}

func Call_Control_Monad_Gen_getFreqVal__2886942013(v_0_loop gopurs_runtime.Value, x_1_loop float64) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var x_1 float64 = x_1_loop
_ = x_1
return (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply(v_0, gopurs_runtime.Float(x_1)).UnsafePtr).V1
}

func Call_Control_Monad_Gen_getFreqVal__3389400221(v_0_loop gopurs_runtime.Value, x_1_loop float64) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var x_1 float64 = x_1_loop
_ = x_1
return (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply(v_0, gopurs_runtime.Float(x_1)).UnsafePtr).V1
}

func Call_Control_Monad_Gen_oneOf__2265861353(dictMonadGen_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen) gopurs_runtime.Value {
var dictMonadGen_0 *Constructor_Control_Monad_Gen_Class_MonadGen = dictMonadGen_0_loop
_ = dictMonadGen_0
// TAST (Let): Bind1_1_0 -> *Constructor_Control_Bind_Bind
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_0.V0), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_0
return gopurs_runtime.Func(func(dictFoldable1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Foldable0_3_1 -> *Constructor_Data_Foldable_Foldable
Foldable0_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_2, "Foldable0"), gopurs_runtime.Value{}))
_ = Foldable0_3_1
return gopurs_runtime.Func(func(xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_2 -> *Constructor_Data_Semiring_Semiring
__local_var_5_2 := &Constructor_Data_Semiring_Semiring{1, Get_Data_Semiring_intAdd(), Get_Data_Semiring_intMul(), gopurs_runtime.Int(1), gopurs_runtime.Int(0)}
_ = __local_var_5_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadGen_0.V3), gopurs_runtime.Int(0), gopurs_runtime.Int((gopurs_runtime.Apply3(gopurs_runtime.Box(Foldable0_3_1.V1), gopurs_runtime.Func(func(c_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(__local_var_5_2.V0), gopurs_runtime.Box(__local_var_5_2.V2), c_6)
})
}), gopurs_runtime.Box(__local_var_5_2.V3), xs_4).IntVal) - (1))), gopurs_runtime.Func(func(n_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Monad_Gen_fromIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](dictFoldable1_2)), gopurs_runtime.Int(n_5.IntVal), xs_4)
}))
})
})
}

func Call_Control_Monad_Gen_oneOf__1433237231(dictMonadGen_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen) gopurs_runtime.Value {
var dictMonadGen_0 *Constructor_Control_Monad_Gen_Class_MonadGen = dictMonadGen_0_loop
_ = dictMonadGen_0
// TAST (Let): Bind1_1_0 -> *Constructor_Control_Bind_Bind
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_0.V0), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_0
// TAST (Let): Foldable0_2_1 -> *Constructor_Data_Foldable_Foldable
Foldable0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](Get_Data_Char_Gen_foldable1NonEmpty()).V0), gopurs_runtime.Value{}))
_ = Foldable0_2_1
return gopurs_runtime.Func(func(xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_2 -> *Constructor_Data_Semiring_Semiring
__local_var_4_2 := &Constructor_Data_Semiring_Semiring{1, Get_Data_Semiring_intAdd(), Get_Data_Semiring_intMul(), gopurs_runtime.Int(1), gopurs_runtime.Int(0)}
_ = __local_var_4_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadGen_0.V3), gopurs_runtime.Int(0), gopurs_runtime.Int((gopurs_runtime.Apply3(gopurs_runtime.Box(Foldable0_2_1.V1), gopurs_runtime.Func(func(c_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(__local_var_4_2.V0), gopurs_runtime.Box(__local_var_4_2.V2), c_5)
})
}), gopurs_runtime.Box(__local_var_4_2.V3), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](xs_3))}).IntVal) - (1))), gopurs_runtime.Func(func(n_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Monad_Gen_fromIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](Get_Data_Char_Gen_foldable1NonEmpty()))})), gopurs_runtime.Int(n_4.IntVal), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](xs_3))})
}))
})
}

func Call_Control_Monad_Gen_unfoldable__1024209525(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec, dictMonadGen_1_loop *Constructor_Control_Monad_Gen_Class_MonadGen) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 *Constructor_Control_Monad_Gen_Class_MonadGen = dictMonadGen_1_loop
_ = dictMonadGen_1
// TAST (Let): Monad0_2_0 -> gopurs_runtime.Value
Monad0_2_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_1.V0), gopurs_runtime.Value{})
_ = Monad0_2_0
// TAST (Let): pure_3_1 -> gopurs_runtime.Value
pure_3_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_1
// TAST (Let): Bind1_4_2 -> *Constructor_Control_Bind_Bind
Bind1_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_2
// TAST (Let): Functor0_5_3 -> *Constructor_Data_Functor_Functor
Functor0_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_3
return gopurs_runtime.Func(func(gen_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_5 -> gopurs_runtime.Value
__local_var_7_5 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadRec_0.V1), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 gopurs_runtime.Value
{
var __t8 bool
{
if ((*Constructor_Data_Tuple_Tuple)(v_7.UnsafePtr).V1.IntVal) > (0) {
__t8 = false
goto end_branch_8
} else {

}
}
{
__t8 = true
}
end_branch_8:
if __t8 {
__t9 = gopurs_runtime.Apply(pure_3_1, gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done{1, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons]((*Constructor_Data_Tuple_Tuple)(v_7.UnsafePtr).V0))}})})
goto end_branch_9
} else {

}
}
{
// TAST (Let): __local_var_8_6 -> gopurs_runtime.Value
__local_var_8_6 := (*Constructor_Data_Tuple_Tuple)(v_7.UnsafePtr).V0
_ = __local_var_8_6
// TAST (Let): __local_var_9_7 -> gopurs_runtime.Value
__local_var_9_7 := (*Constructor_Data_Tuple_Tuple)(v_7.UnsafePtr).V1
_ = __local_var_9_7
__t9 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_2.V1), gen_6, gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_1, gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Gen_Cons{1, gopurs_runtime.Str(x_10.StrVal()), gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons](gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons](__local_var_8_6))})})}, gopurs_runtime.Int((__local_var_9_7.IntVal) - (1))})}})})
}))
}
end_branch_9:
return __t9
}))
_ = __local_var_7_5
// TAST (Let): __local_var_8_10 -> gopurs_runtime.Value
__local_var_8_10 := gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(nil))})
_ = __local_var_8_10
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_3.V0), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Unfoldable_unfoldableArray(), "unfoldr"), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 *Constructor_Data_Maybe_Just
{
if (v_7.Type == 9 && v_7.IntVal == 759514854 && v_7.UnsafePtr == nil) {
__t4 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_4
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 759514854 && v_7.UnsafePtr != nil) {
__t4 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Str((*Constructor_Control_Monad_Gen_Cons)(v_7.UnsafePtr).V0.StrVal()), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(v_7.UnsafePtr).V1)}})}}
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t4)}
})), gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_1.V5), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_5, gopurs_runtime.Apply(__local_var_8_10, x_9))
})))
})
}

func Call_Control_Monad_Gen_unfoldable__542018773(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec, dictMonadGen_1_loop *Constructor_Control_Monad_Gen_Class_MonadGen) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 *Constructor_Control_Monad_Gen_Class_MonadGen = dictMonadGen_1_loop
_ = dictMonadGen_1
// TAST (Let): Monad0_2_0 -> gopurs_runtime.Value
Monad0_2_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_1.V0), gopurs_runtime.Value{})
_ = Monad0_2_0
// TAST (Let): pure_3_1 -> gopurs_runtime.Value
pure_3_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_1
// TAST (Let): Bind1_4_2 -> *Constructor_Control_Bind_Bind
Bind1_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_2
// TAST (Let): Functor0_5_3 -> *Constructor_Data_Functor_Functor
Functor0_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_3
return gopurs_runtime.Func(func(dictUnfoldable_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(gen_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_5 -> gopurs_runtime.Value
__local_var_8_5 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadRec_0.V1), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 gopurs_runtime.Value
{
var __t8 bool
{
if ((*Constructor_Data_Tuple_Tuple)(v_8.UnsafePtr).V1.IntVal) > (0) {
__t8 = false
goto end_branch_8
} else {

}
}
{
__t8 = true
}
end_branch_8:
if __t8 {
__t9 = gopurs_runtime.Apply(pure_3_1, gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done{1, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons]((*Constructor_Data_Tuple_Tuple)(v_8.UnsafePtr).V0))}})})
goto end_branch_9
} else {

}
}
{
// TAST (Let): __local_var_9_6 -> gopurs_runtime.Value
__local_var_9_6 := (*Constructor_Data_Tuple_Tuple)(v_8.UnsafePtr).V0
_ = __local_var_9_6
// TAST (Let): __local_var_10_7 -> gopurs_runtime.Value
__local_var_10_7 := (*Constructor_Data_Tuple_Tuple)(v_8.UnsafePtr).V1
_ = __local_var_10_7
__t9 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_2.V1), gen_7, gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_1, gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Gen_Cons{1, x_11, gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons](__local_var_9_6)})}, gopurs_runtime.Int((__local_var_10_7.IntVal) - (1))})}})})
}))
}
end_branch_9:
return __t9
}))
_ = __local_var_8_5
// TAST (Let): __local_var_9_10 -> gopurs_runtime.Value
__local_var_9_10 := gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(nil))})
_ = __local_var_9_10
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_3.V0), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictUnfoldable_6, "unfoldr"), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 *Constructor_Data_Maybe_Just
{
if (v_8.Type == 9 && v_8.IntVal == 759514854 && v_8.UnsafePtr == nil) {
__t4 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_4
} else {

}
}
{
if (v_8.Type == 9 && v_8.IntVal == 759514854 && v_8.UnsafePtr != nil) {
__t4 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Control_Monad_Gen_Cons)(v_8.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(v_8.UnsafePtr).V1)}})}}
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t4)}
})), gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_1.V5), gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_5, gopurs_runtime.Apply(__local_var_9_10, x_10))
})))
})
})
}

func Call_Control_Monad_Gen_unfoldable__4218841939(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec, dictMonadGen_1_loop *Constructor_Control_Monad_Gen_Class_MonadGen) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 *Constructor_Control_Monad_Gen_Class_MonadGen = dictMonadGen_1_loop
_ = dictMonadGen_1
// TAST (Let): Monad0_2_0 -> gopurs_runtime.Value
Monad0_2_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_1.V0), gopurs_runtime.Value{})
_ = Monad0_2_0
// TAST (Let): pure_3_1 -> gopurs_runtime.Value
pure_3_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_1
// TAST (Let): Bind1_4_2 -> *Constructor_Control_Bind_Bind
Bind1_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_2
// TAST (Let): Functor0_5_3 -> *Constructor_Data_Functor_Functor
Functor0_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_3
return gopurs_runtime.Func(func(gen_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_5 -> gopurs_runtime.Value
__local_var_7_5 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadRec_0.V1), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 gopurs_runtime.Value
{
var __t8 bool
{
if ((*Constructor_Data_Tuple_Tuple)(v_7.UnsafePtr).V1.IntVal) > (0) {
__t8 = false
goto end_branch_8
} else {

}
}
{
__t8 = true
}
end_branch_8:
if __t8 {
__t9 = gopurs_runtime.Apply(pure_3_1, gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done{1, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons]((*Constructor_Data_Tuple_Tuple)(v_7.UnsafePtr).V0))}})})
goto end_branch_9
} else {

}
}
{
// TAST (Let): __local_var_8_6 -> gopurs_runtime.Value
__local_var_8_6 := (*Constructor_Data_Tuple_Tuple)(v_7.UnsafePtr).V0
_ = __local_var_8_6
// TAST (Let): __local_var_9_7 -> gopurs_runtime.Value
__local_var_9_7 := (*Constructor_Data_Tuple_Tuple)(v_7.UnsafePtr).V1
_ = __local_var_9_7
__t9 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_2.V1), gen_6, gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_1, gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Gen_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](x_10))}, gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons](gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons](__local_var_8_6))})})}, gopurs_runtime.Int((__local_var_9_7.IntVal) - (1))})}})})
}))
}
end_branch_9:
return __t9
}))
_ = __local_var_7_5
// TAST (Let): __local_var_8_10 -> gopurs_runtime.Value
__local_var_8_10 := gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(nil))})
_ = __local_var_8_10
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_3.V0), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_List_Types_unfoldableList(), "unfoldr"), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 *Constructor_Data_Maybe_Just
{
if (v_7.Type == 9 && v_7.IntVal == 759514854 && v_7.UnsafePtr == nil) {
__t4 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_4
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 759514854 && v_7.UnsafePtr != nil) {
__t4 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple]((*Constructor_Control_Monad_Gen_Cons)(v_7.UnsafePtr).V0))}, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(v_7.UnsafePtr).V1)}})}}
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t4)}
})), gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_1.V5), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_5, gopurs_runtime.Apply(__local_var_8_10, x_9))
})))
})
}


