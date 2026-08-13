package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Array_ST_unshiftAll gopurs_runtime.Value
var once_Data_Array_ST_unshiftAll sync.Once
func Get_Data_Array_ST_unshiftAll() gopurs_runtime.Value {
	once_Data_Array_ST_unshiftAll.Do(func() {
		cache_Data_Array_ST_unshiftAll = gopurs_runtime.Apply(Get_Control_Monad_ST_Uncurried_runSTFn2(), Get_Data_Array_ST_unshiftAllImpl())
	})
	return cache_Data_Array_ST_unshiftAll
}

var cache_Data_Array_ST_unshift gopurs_runtime.Value
var once_Data_Array_ST_unshift sync.Once
func Get_Data_Array_ST_unshift() gopurs_runtime.Value {
	once_Data_Array_ST_unshift.Do(func() {
		cache_Data_Array_ST_unshift = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_ST_unshift(a_0_box)
})
	})
	return cache_Data_Array_ST_unshift
}

var cache_Data_Array_ST_unsafeThaw gopurs_runtime.Value
var once_Data_Array_ST_unsafeThaw sync.Once
func Get_Data_Array_ST_unsafeThaw() gopurs_runtime.Value {
	once_Data_Array_ST_unsafeThaw.Do(func() {
		cache_Data_Array_ST_unsafeThaw = gopurs_runtime.Apply(Get_Control_Monad_ST_Uncurried_runSTFn1(), Get_Data_Array_ST_unsafeThawImpl())
	})
	return cache_Data_Array_ST_unsafeThaw
}

var cache_Data_Array_ST_unsafeFreeze gopurs_runtime.Value
var once_Data_Array_ST_unsafeFreeze sync.Once
func Get_Data_Array_ST_unsafeFreeze() gopurs_runtime.Value {
	once_Data_Array_ST_unsafeFreeze.Do(func() {
		cache_Data_Array_ST_unsafeFreeze = gopurs_runtime.Apply(Get_Control_Monad_ST_Uncurried_runSTFn1(), Get_Data_Array_ST_unsafeFreezeImpl())
	})
	return cache_Data_Array_ST_unsafeFreeze
}

var cache_Data_Array_ST_toAssocArray gopurs_runtime.Value
var once_Data_Array_ST_toAssocArray sync.Once
func Get_Data_Array_ST_toAssocArray() gopurs_runtime.Value {
	once_Data_Array_ST_toAssocArray.Do(func() {
		cache_Data_Array_ST_toAssocArray = gopurs_runtime.Apply(Get_Control_Monad_ST_Uncurried_runSTFn1(), Get_Data_Array_ST_toAssocArrayImpl())
	})
	return cache_Data_Array_ST_toAssocArray
}

var cache_Data_Array_ST_thaw gopurs_runtime.Value
var once_Data_Array_ST_thaw sync.Once
func Get_Data_Array_ST_thaw() gopurs_runtime.Value {
	once_Data_Array_ST_thaw.Do(func() {
		cache_Data_Array_ST_thaw = gopurs_runtime.Apply(Get_Control_Monad_ST_Uncurried_runSTFn1(), Get_Data_Array_ST_thawImpl())
	})
	return cache_Data_Array_ST_thaw
}

var cache_Data_Array_ST_withArray gopurs_runtime.Value
var once_Data_Array_ST_withArray sync.Once
func Get_Data_Array_ST_withArray() gopurs_runtime.Value {
	once_Data_Array_ST_withArray.Do(func() {
		cache_Data_Array_ST_withArray = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_ST_withArray(f_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
	})
	return cache_Data_Array_ST_withArray
}

var cache_Data_Array_ST_splice gopurs_runtime.Value
var once_Data_Array_ST_splice sync.Once
func Get_Data_Array_ST_splice() gopurs_runtime.Value {
	once_Data_Array_ST_splice.Do(func() {
		cache_Data_Array_ST_splice = gopurs_runtime.Apply(Get_Control_Monad_ST_Uncurried_runSTFn4(), Get_Data_Array_ST_spliceImpl())
	})
	return cache_Data_Array_ST_splice
}

var cache_Data_Array_ST_sortBy gopurs_runtime.Value
var once_Data_Array_ST_sortBy sync.Once
func Get_Data_Array_ST_sortBy() gopurs_runtime.Value {
	once_Data_Array_ST_sortBy.Do(func() {
		cache_Data_Array_ST_sortBy = gopurs_runtime.Func(func(comp_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_ST_sortBy(comp_0_box)
})
	})
	return cache_Data_Array_ST_sortBy
}

var cache_Data_Array_ST_sortWith gopurs_runtime.Value
var once_Data_Array_ST_sortWith sync.Once
func Get_Data_Array_ST_sortWith() gopurs_runtime.Value {
	once_Data_Array_ST_sortWith.Do(func() {
		cache_Data_Array_ST_sortWith = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_ST_sortWith(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), f_1_box)
})
	})
	return cache_Data_Array_ST_sortWith
}

var cache_Data_Array_ST_sort gopurs_runtime.Value
var once_Data_Array_ST_sort sync.Once
func Get_Data_Array_ST_sort() gopurs_runtime.Value {
	once_Data_Array_ST_sort.Do(func() {
		cache_Data_Array_ST_sort = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_ST_sort(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
})
	})
	return cache_Data_Array_ST_sort
}

var cache_Data_Array_ST_shift gopurs_runtime.Value
var once_Data_Array_ST_shift sync.Once
func Get_Data_Array_ST_shift() gopurs_runtime.Value {
	once_Data_Array_ST_shift.Do(func() {
		cache_Data_Array_ST_shift = gopurs_runtime.Apply3(Get_Control_Monad_ST_Uncurried_runSTFn3(), Get_Data_Array_ST_shiftImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
	})
	return cache_Data_Array_ST_shift
}

var cache_Data_Array_ST_run gopurs_runtime.Value
var once_Data_Array_ST_run sync.Once
func Get_Data_Array_ST_run() gopurs_runtime.Value {
	once_Data_Array_ST_run.Do(func() {
		cache_Data_Array_ST_run = gopurs_runtime.Func(func(st_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_ST_run(st_0_box))
})
	})
	return cache_Data_Array_ST_run
}

var cache_Data_Array_ST_pushAll gopurs_runtime.Value
var once_Data_Array_ST_pushAll sync.Once
func Get_Data_Array_ST_pushAll() gopurs_runtime.Value {
	once_Data_Array_ST_pushAll.Do(func() {
		cache_Data_Array_ST_pushAll = gopurs_runtime.Apply(Get_Control_Monad_ST_Uncurried_runSTFn2(), Get_Data_Array_ST_pushAllImpl())
	})
	return cache_Data_Array_ST_pushAll
}

var cache_Data_Array_ST_push gopurs_runtime.Value
var once_Data_Array_ST_push sync.Once
func Get_Data_Array_ST_push() gopurs_runtime.Value {
	once_Data_Array_ST_push.Do(func() {
		cache_Data_Array_ST_push = gopurs_runtime.Apply(Get_Control_Monad_ST_Uncurried_runSTFn2(), Get_Data_Array_ST_pushImpl())
	})
	return cache_Data_Array_ST_push
}

var cache_Data_Array_ST_pop gopurs_runtime.Value
var once_Data_Array_ST_pop sync.Once
func Get_Data_Array_ST_pop() gopurs_runtime.Value {
	once_Data_Array_ST_pop.Do(func() {
		cache_Data_Array_ST_pop = gopurs_runtime.Apply3(Get_Control_Monad_ST_Uncurried_runSTFn3(), Get_Data_Array_ST_popImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
	})
	return cache_Data_Array_ST_pop
}

var cache_Data_Array_ST_poke gopurs_runtime.Value
var once_Data_Array_ST_poke sync.Once
func Get_Data_Array_ST_poke() gopurs_runtime.Value {
	once_Data_Array_ST_poke.Do(func() {
		cache_Data_Array_ST_poke = gopurs_runtime.Apply(Get_Control_Monad_ST_Uncurried_runSTFn3(), Get_Data_Array_ST_pokeImpl())
	})
	return cache_Data_Array_ST_poke
}

var cache_Data_Array_ST_peek gopurs_runtime.Value
var once_Data_Array_ST_peek sync.Once
func Get_Data_Array_ST_peek() gopurs_runtime.Value {
	once_Data_Array_ST_peek.Do(func() {
		cache_Data_Array_ST_peek = gopurs_runtime.Apply3(Get_Control_Monad_ST_Uncurried_runSTFn4(), Get_Data_Array_ST_peekImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
	})
	return cache_Data_Array_ST_peek
}

var cache_Data_Array_ST_new gopurs_runtime.Value
var once_Data_Array_ST_new sync.Once
func Get_Data_Array_ST_new() gopurs_runtime.Value {
	once_Data_Array_ST_new.Do(func() {
		cache_Data_Array_ST_new = Get_Data_Array_ST_newImpl()
	})
	return cache_Data_Array_ST_new
}

var cache_Data_Array_ST_modify gopurs_runtime.Value
var once_Data_Array_ST_modify sync.Once
func Get_Data_Array_ST_modify() gopurs_runtime.Value {
	once_Data_Array_ST_modify.Do(func() {
		cache_Data_Array_ST_modify = gopurs_runtime.Func3(func(i_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_ST_modify(i_0_box.IntVal, f_1_box, xs_2_box)
})
	})
	return cache_Data_Array_ST_modify
}

var cache_Data_Array_ST_length gopurs_runtime.Value
var once_Data_Array_ST_length sync.Once
func Get_Data_Array_ST_length() gopurs_runtime.Value {
	once_Data_Array_ST_length.Do(func() {
		cache_Data_Array_ST_length = gopurs_runtime.Apply(Get_Control_Monad_ST_Uncurried_runSTFn1(), Get_Data_Array_ST_lengthImpl())
	})
	return cache_Data_Array_ST_length
}

var cache_Data_Array_ST_freeze gopurs_runtime.Value
var once_Data_Array_ST_freeze sync.Once
func Get_Data_Array_ST_freeze() gopurs_runtime.Value {
	once_Data_Array_ST_freeze.Do(func() {
		cache_Data_Array_ST_freeze = gopurs_runtime.Apply(Get_Control_Monad_ST_Uncurried_runSTFn1(), Get_Data_Array_ST_freezeImpl())
	})
	return cache_Data_Array_ST_freeze
}

var cache_Data_Array_ST_clone gopurs_runtime.Value
var once_Data_Array_ST_clone sync.Once
func Get_Data_Array_ST_clone() gopurs_runtime.Value {
	once_Data_Array_ST_clone.Do(func() {
		cache_Data_Array_ST_clone = gopurs_runtime.Apply(Get_Control_Monad_ST_Uncurried_runSTFn1(), Get_Data_Array_ST_cloneImpl())
	})
	return cache_Data_Array_ST_clone
}

var cache_Data_Array_ST_modify__3999472078 gopurs_runtime.Value
var once_Data_Array_ST_modify__3999472078 sync.Once
func Get_Data_Array_ST_modify__3999472078() gopurs_runtime.Value {
	once_Data_Array_ST_modify__3999472078.Do(func() {
		cache_Data_Array_ST_modify__3999472078 = gopurs_runtime.Func3(func(i_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_ST_modify__3999472078(i_0_box.IntVal, f_1_box, xs_2_box)
})
	})
	return cache_Data_Array_ST_modify__3999472078
}

var cache_Data_Array_ST_new__361268277 gopurs_runtime.Value
var once_Data_Array_ST_new__361268277 sync.Once
func Get_Data_Array_ST_new__361268277() gopurs_runtime.Value {
	once_Data_Array_ST_new__361268277.Do(func() {
		cache_Data_Array_ST_new__361268277 = Get_Data_Array_ST_newImpl()
	})
	return cache_Data_Array_ST_new__361268277
}

var cache_Data_Array_ST_new__108694771 gopurs_runtime.Value
var once_Data_Array_ST_new__108694771 sync.Once
func Get_Data_Array_ST_new__108694771() gopurs_runtime.Value {
	once_Data_Array_ST_new__108694771.Do(func() {
		cache_Data_Array_ST_new__108694771 = Get_Data_Array_ST_newImpl()
	})
	return cache_Data_Array_ST_new__108694771
}

var cache_Data_Array_ST_peek__1114289340 gopurs_runtime.Value
var once_Data_Array_ST_peek__1114289340 sync.Once
func Get_Data_Array_ST_peek__1114289340() gopurs_runtime.Value {
	once_Data_Array_ST_peek__1114289340.Do(func() {
		cache_Data_Array_ST_peek__1114289340 = gopurs_runtime.Apply3(Get_Control_Monad_ST_Uncurried_runSTFn4(), Get_Data_Array_ST_peekImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
	})
	return cache_Data_Array_ST_peek__1114289340
}

var cache_Data_Array_ST_poke__1035212550 gopurs_runtime.Value
var once_Data_Array_ST_poke__1035212550 sync.Once
func Get_Data_Array_ST_poke__1035212550() gopurs_runtime.Value {
	once_Data_Array_ST_poke__1035212550.Do(func() {
		cache_Data_Array_ST_poke__1035212550 = gopurs_runtime.Apply(Get_Control_Monad_ST_Uncurried_runSTFn3(), Get_Data_Array_ST_pokeImpl())
	})
	return cache_Data_Array_ST_poke__1035212550
}

var cache_Data_Array_ST_push__1557574173 gopurs_runtime.Value
var once_Data_Array_ST_push__1557574173 sync.Once
func Get_Data_Array_ST_push__1557574173() gopurs_runtime.Value {
	once_Data_Array_ST_push__1557574173.Do(func() {
		cache_Data_Array_ST_push__1557574173 = gopurs_runtime.Apply(Get_Control_Monad_ST_Uncurried_runSTFn2(), Get_Data_Array_ST_pushImpl())
	})
	return cache_Data_Array_ST_push__1557574173
}

var cache_Data_Array_ST_push__2784843933 gopurs_runtime.Value
var once_Data_Array_ST_push__2784843933 sync.Once
func Get_Data_Array_ST_push__2784843933() gopurs_runtime.Value {
	once_Data_Array_ST_push__2784843933.Do(func() {
		cache_Data_Array_ST_push__2784843933 = gopurs_runtime.Apply(Get_Control_Monad_ST_Uncurried_runSTFn2(), Get_Data_Array_ST_pushImpl())
	})
	return cache_Data_Array_ST_push__2784843933
}

var cache_Data_Array_ST_push__3387684349 gopurs_runtime.Value
var once_Data_Array_ST_push__3387684349 sync.Once
func Get_Data_Array_ST_push__3387684349() gopurs_runtime.Value {
	once_Data_Array_ST_push__3387684349.Do(func() {
		cache_Data_Array_ST_push__3387684349 = gopurs_runtime.Apply(Get_Control_Monad_ST_Uncurried_runSTFn2(), Get_Data_Array_ST_pushImpl())
	})
	return cache_Data_Array_ST_push__3387684349
}

var cache_Data_Array_ST_run__1673253202 gopurs_runtime.Value
var once_Data_Array_ST_run__1673253202 sync.Once
func Get_Data_Array_ST_run__1673253202() gopurs_runtime.Value {
	once_Data_Array_ST_run__1673253202.Do(func() {
		cache_Data_Array_ST_run__1673253202 = gopurs_runtime.Func(func(st_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_ST_run__1673253202(st_0_box))
})
	})
	return cache_Data_Array_ST_run__1673253202
}

var cache_Data_Array_ST_sortBy__2466834470 gopurs_runtime.Value
var once_Data_Array_ST_sortBy__2466834470 sync.Once
func Get_Data_Array_ST_sortBy__2466834470() gopurs_runtime.Value {
	once_Data_Array_ST_sortBy__2466834470.Do(func() {
		cache_Data_Array_ST_sortBy__2466834470 = gopurs_runtime.Func(func(comp_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_ST_sortBy__2466834470(comp_0_box)
})
	})
	return cache_Data_Array_ST_sortBy__2466834470
}

var cache_Data_Array_ST_thaw__2850628855 gopurs_runtime.Value
var once_Data_Array_ST_thaw__2850628855 sync.Once
func Get_Data_Array_ST_thaw__2850628855() gopurs_runtime.Value {
	once_Data_Array_ST_thaw__2850628855.Do(func() {
		cache_Data_Array_ST_thaw__2850628855 = gopurs_runtime.Apply(Get_Control_Monad_ST_Uncurried_runSTFn1(), Get_Data_Array_ST_thawImpl())
	})
	return cache_Data_Array_ST_thaw__2850628855
}

var cache_Data_Array_ST_thaw__772033143 gopurs_runtime.Value
var once_Data_Array_ST_thaw__772033143 sync.Once
func Get_Data_Array_ST_thaw__772033143() gopurs_runtime.Value {
	once_Data_Array_ST_thaw__772033143.Do(func() {
		cache_Data_Array_ST_thaw__772033143 = gopurs_runtime.Apply(Get_Control_Monad_ST_Uncurried_runSTFn1(), Get_Data_Array_ST_thawImpl())
	})
	return cache_Data_Array_ST_thaw__772033143
}

var cache_Data_Array_ST_unsafeFreeze__2265612059 gopurs_runtime.Value
var once_Data_Array_ST_unsafeFreeze__2265612059 sync.Once
func Get_Data_Array_ST_unsafeFreeze__2265612059() gopurs_runtime.Value {
	once_Data_Array_ST_unsafeFreeze__2265612059.Do(func() {
		cache_Data_Array_ST_unsafeFreeze__2265612059 = gopurs_runtime.Apply(Get_Control_Monad_ST_Uncurried_runSTFn1(), Get_Data_Array_ST_unsafeFreezeImpl())
	})
	return cache_Data_Array_ST_unsafeFreeze__2265612059
}

var cache_Data_Array_ST_unsafeFreeze__3261307415 gopurs_runtime.Value
var once_Data_Array_ST_unsafeFreeze__3261307415 sync.Once
func Get_Data_Array_ST_unsafeFreeze__3261307415() gopurs_runtime.Value {
	once_Data_Array_ST_unsafeFreeze__3261307415.Do(func() {
		cache_Data_Array_ST_unsafeFreeze__3261307415 = gopurs_runtime.Apply(Get_Control_Monad_ST_Uncurried_runSTFn1(), Get_Data_Array_ST_unsafeFreezeImpl())
	})
	return cache_Data_Array_ST_unsafeFreeze__3261307415
}

var cache_Data_Array_ST_unsafeFreeze__3249621719 gopurs_runtime.Value
var once_Data_Array_ST_unsafeFreeze__3249621719 sync.Once
func Get_Data_Array_ST_unsafeFreeze__3249621719() gopurs_runtime.Value {
	once_Data_Array_ST_unsafeFreeze__3249621719.Do(func() {
		cache_Data_Array_ST_unsafeFreeze__3249621719 = gopurs_runtime.Apply(Get_Control_Monad_ST_Uncurried_runSTFn1(), Get_Data_Array_ST_unsafeFreezeImpl())
	})
	return cache_Data_Array_ST_unsafeFreeze__3249621719
}

var cache_Data_Array_ST_unsafeFreeze__4066092631 gopurs_runtime.Value
var once_Data_Array_ST_unsafeFreeze__4066092631 sync.Once
func Get_Data_Array_ST_unsafeFreeze__4066092631() gopurs_runtime.Value {
	once_Data_Array_ST_unsafeFreeze__4066092631.Do(func() {
		cache_Data_Array_ST_unsafeFreeze__4066092631 = gopurs_runtime.Apply(Get_Control_Monad_ST_Uncurried_runSTFn1(), Get_Data_Array_ST_unsafeFreezeImpl())
	})
	return cache_Data_Array_ST_unsafeFreeze__4066092631
}

var cache_Data_Array_ST_unsafeThaw__1529848091 gopurs_runtime.Value
var once_Data_Array_ST_unsafeThaw__1529848091 sync.Once
func Get_Data_Array_ST_unsafeThaw__1529848091() gopurs_runtime.Value {
	once_Data_Array_ST_unsafeThaw__1529848091.Do(func() {
		cache_Data_Array_ST_unsafeThaw__1529848091 = gopurs_runtime.Apply(Get_Control_Monad_ST_Uncurried_runSTFn1(), Get_Data_Array_ST_unsafeThawImpl())
	})
	return cache_Data_Array_ST_unsafeThaw__1529848091
}

var cache_Data_Array_ST_withArray__126410905 gopurs_runtime.Value
var once_Data_Array_ST_withArray__126410905 sync.Once
func Get_Data_Array_ST_withArray__126410905() gopurs_runtime.Value {
	once_Data_Array_ST_withArray__126410905.Do(func() {
		cache_Data_Array_ST_withArray__126410905 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_ST_withArray__126410905(f_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
	})
	return cache_Data_Array_ST_withArray__126410905
}

var cache_Data_Array_ST_withArray__3965688428 gopurs_runtime.Value
var once_Data_Array_ST_withArray__3965688428 sync.Once
func Get_Data_Array_ST_withArray__3965688428() gopurs_runtime.Value {
	once_Data_Array_ST_withArray__3965688428.Do(func() {
		cache_Data_Array_ST_withArray__3965688428 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_ST_withArray__3965688428(f_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
	})
	return cache_Data_Array_ST_withArray__3965688428
}

var cache_Data_Array_ST_withArray__2779921436 gopurs_runtime.Value
var once_Data_Array_ST_withArray__2779921436 sync.Once
func Get_Data_Array_ST_withArray__2779921436() gopurs_runtime.Value {
	once_Data_Array_ST_withArray__2779921436.Do(func() {
		cache_Data_Array_ST_withArray__2779921436 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_ST_withArray__2779921436(f_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
	})
	return cache_Data_Array_ST_withArray__2779921436
}

var cache_Data_Array_ST_withArray__1455372447 gopurs_runtime.Value
var once_Data_Array_ST_withArray__1455372447 sync.Once
func Get_Data_Array_ST_withArray__1455372447() gopurs_runtime.Value {
	once_Data_Array_ST_withArray__1455372447.Do(func() {
		cache_Data_Array_ST_withArray__1455372447 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_ST_withArray__1455372447(f_0_box, func() [][]gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([][]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}() }
					return unboxed
				}())
})
	})
	return cache_Data_Array_ST_withArray__1455372447
}

var cache_Data_Array_ST_withArray__2734111834 gopurs_runtime.Value
var once_Data_Array_ST_withArray__2734111834 sync.Once
func Get_Data_Array_ST_withArray__2734111834() gopurs_runtime.Value {
	once_Data_Array_ST_withArray__2734111834.Do(func() {
		cache_Data_Array_ST_withArray__2734111834 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_ST_withArray__2734111834(f_0_box, func() [][]gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([][]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}() }
					return unboxed
				}())
})
	})
	return cache_Data_Array_ST_withArray__2734111834
}

func Call_Data_Array_ST_unshift(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return gopurs_runtime.Apply2(Get_Control_Monad_ST_Uncurried_runSTFn2(), Get_Data_Array_ST_unshiftAllImpl(), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{a_0}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
}

func Call_Data_Array_ST_withArray(f_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(Get_Data_Array_ST_thawImpl(), gopurs_runtime.Array(xs_1))
}), gopurs_runtime.Func(func(result_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Apply(f_0, result_2), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeFreezeImpl(), result_2)
})
}))
}))
}

func Call_Data_Array_ST_sortBy(comp_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var comp_0 gopurs_runtime.Value = comp_0_loop
_ = comp_0
return gopurs_runtime.Apply3(Get_Control_Monad_ST_Uncurried_runSTFn3(), Get_Data_Array_ST_sortByImpl(), comp_0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 int64
{
if (uint32(v_1.IntVal) == 380165415) {
__t0 = 1
goto end_branch_0
} else {

}
}
{
if (uint32(v_1.IntVal) == 902936544) {
__t0 = 0
goto end_branch_0
} else {

}
}
{
if (uint32(v_1.IntVal) == 1527465420) {
__t0 = -1
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_0:
return gopurs_runtime.Int(__t0)
}))
}

func Call_Data_Array_ST_sortWith(dictOrd_0_loop *Constructor_Data_Ord_Ord, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return Call_Data_Array_ST_sortBy(gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), gopurs_runtime.Apply(f_1, x_2), gopurs_runtime.Apply(f_1, y_3)).IntVal)), UnsafePtr: nil}
})
}))
}

func Call_Data_Array_ST_sort(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
return Call_Data_Array_ST_sortBy(gopurs_runtime.Box(dictOrd_0.V1))
}

func Call_Data_Array_ST_run(st_0_loop gopurs_runtime.Value) []gopurs_runtime.Value {
var st_0 gopurs_runtime.Value = st_0_loop
_ = st_0
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), st_0, Get_Data_Array_ST_unsafeFreeze())).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_ST_modify(i_0_loop int64, f_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var i_0 int64 = i_0_loop
_ = i_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp4(Get_Data_Array_ST_peekImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Int(i_0), xs_2)
}), gopurs_runtime.Func(func(entry_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (entry_3.Type == 9 && entry_3.IntVal == 930809136 && entry_3.UnsafePtr != nil) {
// TAST (Let): __local_var_4_0 -> gopurs_runtime.Value
__local_var_4_0 := gopurs_runtime.Apply(f_1, (*Constructor_Data_Maybe_Just)(entry_3.UnsafePtr).V0)
_ = __local_var_4_0
__t1 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp3(Get_Data_Array_ST_pokeImpl(), gopurs_runtime.Int(i_0), __local_var_4_0, xs_2)
})
goto end_branch_1
} else {

}
}
{
if (entry_3.Type == 9 && entry_3.IntVal == 930809136 && entry_3.UnsafePtr == nil) {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_applicativeST(), "pure"), gopurs_runtime.Bool(false))
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}))
}

func Call_Data_Array_ST_modify__3999472078(i_0_loop int64, f_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var i_0 int64 = i_0_loop
_ = i_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp4(Get_Data_Array_ST_peekImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Int(i_0), xs_2)
}), gopurs_runtime.Func(func(entry_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (entry_3.Type == 9 && entry_3.IntVal == 930809136 && entry_3.UnsafePtr != nil) {
// TAST (Let): __local_var_4_0 -> gopurs_runtime.Value
__local_var_4_0 := gopurs_runtime.Apply(f_1, (*Constructor_Data_Maybe_Just)(entry_3.UnsafePtr).V0)
_ = __local_var_4_0
__t1 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp3(Get_Data_Array_ST_pokeImpl(), gopurs_runtime.Int(i_0), __local_var_4_0, xs_2)
})
goto end_branch_1
} else {

}
}
{
if (entry_3.Type == 9 && entry_3.IntVal == 930809136 && entry_3.UnsafePtr == nil) {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_applicativeST(), "pure"), gopurs_runtime.Bool(false))
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}))
}

func Call_Data_Array_ST_run__1673253202(st_0_loop gopurs_runtime.Value) []gopurs_runtime.Value {
var st_0 gopurs_runtime.Value = st_0_loop
_ = st_0
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), st_0, Get_Data_Array_ST_unsafeFreeze())).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_ST_sortBy__2466834470(comp_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var comp_0 gopurs_runtime.Value = comp_0_loop
_ = comp_0
return gopurs_runtime.Apply3(Get_Control_Monad_ST_Uncurried_runSTFn3(), Get_Data_Array_ST_sortByImpl(), comp_0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 int64
{
if (uint32(v_1.IntVal) == 380165415) {
__t0 = 1
goto end_branch_0
} else {

}
}
{
if (uint32(v_1.IntVal) == 902936544) {
__t0 = 0
goto end_branch_0
} else {

}
}
{
if (uint32(v_1.IntVal) == 1527465420) {
__t0 = -1
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_0:
return gopurs_runtime.Int(__t0)
}))
}

func Call_Data_Array_ST_withArray__126410905(f_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(Get_Data_Array_ST_thawImpl(), gopurs_runtime.Array(xs_1))
}), gopurs_runtime.Func(func(result_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Apply(f_0, result_2), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeFreezeImpl(), result_2)
})
}))
}))
}

func Call_Data_Array_ST_withArray__3965688428(f_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(Get_Data_Array_ST_thawImpl(), gopurs_runtime.Array(xs_1))
}), gopurs_runtime.Func(func(result_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Apply(f_0, result_2), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeFreezeImpl(), result_2)
})
}))
}))
}

func Call_Data_Array_ST_withArray__2779921436(f_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(Get_Data_Array_ST_thawImpl(), gopurs_runtime.Array(xs_1))
}), gopurs_runtime.Func(func(result_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Apply(f_0, result_2), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeFreezeImpl(), result_2)
})
}))
}))
}

func Call_Data_Array_ST_withArray__1455372447(f_0_loop gopurs_runtime.Value, xs_1_loop [][]gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var xs_1 [][]gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(Get_Data_Array_ST_thawImpl(), func() gopurs_runtime.Value {
					arr := xs_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Array(v) }
					return gopurs_runtime.Array(boxed)
				}())
}), gopurs_runtime.Func(func(result_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Apply(f_0, result_2), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeFreezeImpl(), result_2)
})
}))
}))
}

func Call_Data_Array_ST_withArray__2734111834(f_0_loop gopurs_runtime.Value, xs_1_loop [][]gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var xs_1 [][]gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(Get_Data_Array_ST_thawImpl(), func() gopurs_runtime.Value {
					arr := xs_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Array(v) }
					return gopurs_runtime.Array(boxed)
				}())
}), gopurs_runtime.Func(func(result_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Apply(f_0, result_2), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeFreezeImpl(), result_2)
})
}))
}))
}

func Get_Data_Array_ST_cloneImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_ST_CloneImpl
}

func Get_Data_Array_ST_freezeImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_ST_FreezeImpl
}

func Get_Data_Array_ST_lengthImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_ST_LengthImpl
}

func Get_Data_Array_ST_newImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_ST_NewImpl
}

func Get_Data_Array_ST_peekImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_ST_PeekImpl
}

func Get_Data_Array_ST_pokeImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_ST_PokeImpl
}

func Get_Data_Array_ST_popImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_ST_PopImpl
}

func Get_Data_Array_ST_pushAllImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_ST_PushAllImpl
}

func Get_Data_Array_ST_pushImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_ST_PushImpl
}

func Get_Data_Array_ST_shiftImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_ST_ShiftImpl
}

func Get_Data_Array_ST_sortByImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_ST_SortByImpl
}

func Get_Data_Array_ST_spliceImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_ST_SpliceImpl
}

func Get_Data_Array_ST_thawImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_ST_ThawImpl
}

func Get_Data_Array_ST_toAssocArrayImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_ST_ToAssocArrayImpl
}

func Get_Data_Array_ST_unsafeFreezeImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_ST_UnsafeFreezeImpl
}

func Get_Data_Array_ST_unsafeThawImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_ST_UnsafeThawImpl
}

func Get_Data_Array_ST_unshiftAllImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_ST_UnshiftAllImpl
}
