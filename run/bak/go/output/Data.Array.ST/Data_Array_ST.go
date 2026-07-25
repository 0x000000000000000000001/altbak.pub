package Data_Array_ST

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Monad_ST_Uncurried "gopurs/output/Control.Monad.ST.Uncurried"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Control_Monad_ST_Internal "gopurs/output/Control.Monad.ST.Internal"
	unsafe "unsafe"
)

var cache_unshiftAll gopurs_runtime.Value
var once_unshiftAll sync.Once
func Get_unshiftAll() gopurs_runtime.Value {
	once_unshiftAll.Do(func() {
		cache_unshiftAll = gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn2(), Get_unshiftAllImpl())
	})
	return cache_unshiftAll
}

var cache_unshift gopurs_runtime.Value
var once_unshift sync.Once
func Get_unshift() gopurs_runtime.Value {
	once_unshift.Do(func() {
		cache_unshift = gopurs_runtime.Func(func(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return gopurs_runtime.Apply2(pkg_Control_Monad_ST_Uncurried.Get_runSTFn2(), Get_unshiftAllImpl(), gopurs_runtime.Array([]gopurs_runtime.Value{a_0}))
}()
})
	})
	return cache_unshift
}

var cache_unsafeThaw gopurs_runtime.Value
var once_unsafeThaw sync.Once
func Get_unsafeThaw() gopurs_runtime.Value {
	once_unsafeThaw.Do(func() {
		cache_unsafeThaw = gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn1(), Get_unsafeThawImpl())
	})
	return cache_unsafeThaw
}

var cache_unsafeFreeze gopurs_runtime.Value
var once_unsafeFreeze sync.Once
func Get_unsafeFreeze() gopurs_runtime.Value {
	once_unsafeFreeze.Do(func() {
		cache_unsafeFreeze = gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn1(), Get_unsafeFreezeImpl())
	})
	return cache_unsafeFreeze
}

var cache_toAssocArray gopurs_runtime.Value
var once_toAssocArray sync.Once
func Get_toAssocArray() gopurs_runtime.Value {
	once_toAssocArray.Do(func() {
		cache_toAssocArray = gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn1(), Get_toAssocArrayImpl())
	})
	return cache_toAssocArray
}

var cache_thaw gopurs_runtime.Value
var once_thaw sync.Once
func Get_thaw() gopurs_runtime.Value {
	once_thaw.Do(func() {
		cache_thaw = gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn1(), Get_thawImpl())
	})
	return cache_thaw
}

var cache_withArray gopurs_runtime.Value
var once_withArray sync.Once
func Get_withArray() gopurs_runtime.Value {
	once_withArray.Do(func() {
		cache_withArray = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_withArray(f_0_box, xs_1_box)
})
	})
	return cache_withArray
}

var cache_splice gopurs_runtime.Value
var once_splice sync.Once
func Get_splice() gopurs_runtime.Value {
	once_splice.Do(func() {
		cache_splice = gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn4(), Get_spliceImpl())
	})
	return cache_splice
}

var cache_sortBy gopurs_runtime.Value
var once_sortBy sync.Once
func Get_sortBy() gopurs_runtime.Value {
	once_sortBy.Do(func() {
		cache_sortBy = gopurs_runtime.Func(func(comp_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var comp_0 gopurs_runtime.Value = comp_0_loop
_ = comp_0
return gopurs_runtime.Apply3(pkg_Control_Monad_ST_Uncurried.Get_runSTFn3(), Get_sortByImpl(), comp_0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_1.Type == 9 && v_1.IntVal == 380165415) {
__t0 = gopurs_runtime.Int(1)
goto end_branch_0
} else {

}
}
{
if (v_1.Type == 9 && v_1.IntVal == 902936544) {
__t0 = gopurs_runtime.Int(0)
goto end_branch_0
} else {

}
}
{
if (v_1.Type == 9 && v_1.IntVal == 1527465420) {
__t0 = gopurs_runtime.Int(-1)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}))
}()
})
	})
	return cache_sortBy
}

var cache_sortWith gopurs_runtime.Value
var once_sortWith sync.Once
func Get_sortWith() gopurs_runtime.Value {
	once_sortWith.Do(func() {
		cache_sortWith = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sortWith(dictOrd_0_box, f_1_box)
})
	})
	return cache_sortWith
}

var cache_sort gopurs_runtime.Value
var once_sort sync.Once
func Get_sort() gopurs_runtime.Value {
	once_sort.Do(func() {
		cache_sort = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_sortBy(), gopurs_runtime.RecordGet(dictOrd_0, "compare"))
}()
})
	})
	return cache_sort
}

var cache_shift gopurs_runtime.Value
var once_shift sync.Once
func Get_shift() gopurs_runtime.Value {
	once_shift.Do(func() {
		cache_shift = gopurs_runtime.Apply3(pkg_Control_Monad_ST_Uncurried.Get_runSTFn3(), Get_shiftImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})})
	})
	return cache_shift
}

var cache_run gopurs_runtime.Value
var once_run sync.Once
func Get_run() gopurs_runtime.Value {
	once_run.Do(func() {
		cache_run = gopurs_runtime.Func(func(st_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var st_0 gopurs_runtime.Value = st_0_loop
_ = st_0
return gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(st_0, gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_unsafeFreeze(), __local_var_1_0), gopurs_runtime.Value{})
}))
}()
})
	})
	return cache_run
}

var cache_pushAll gopurs_runtime.Value
var once_pushAll sync.Once
func Get_pushAll() gopurs_runtime.Value {
	once_pushAll.Do(func() {
		cache_pushAll = gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn2(), Get_pushAllImpl())
	})
	return cache_pushAll
}

var cache_push gopurs_runtime.Value
var once_push sync.Once
func Get_push() gopurs_runtime.Value {
	once_push.Do(func() {
		cache_push = gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn2(), Get_pushImpl())
	})
	return cache_push
}

var cache_pop gopurs_runtime.Value
var once_pop sync.Once
func Get_pop() gopurs_runtime.Value {
	once_pop.Do(func() {
		cache_pop = gopurs_runtime.Apply3(pkg_Control_Monad_ST_Uncurried.Get_runSTFn3(), Get_popImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})})
	})
	return cache_pop
}

var cache_poke gopurs_runtime.Value
var once_poke sync.Once
func Get_poke() gopurs_runtime.Value {
	once_poke.Do(func() {
		cache_poke = gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn3(), Get_pokeImpl())
	})
	return cache_poke
}

var cache_peek gopurs_runtime.Value
var once_peek sync.Once
func Get_peek() gopurs_runtime.Value {
	once_peek.Do(func() {
		cache_peek = gopurs_runtime.Apply3(pkg_Control_Monad_ST_Uncurried.Get_runSTFn4(), Get_peekImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})})
	})
	return cache_peek
}

var cache_modify gopurs_runtime.Value
var once_modify sync.Once
func Get_modify() gopurs_runtime.Value {
	once_modify.Do(func() {
		cache_modify = gopurs_runtime.Func3(func(i_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_modify(i_0_box, f_1_box, xs_2_box)
})
	})
	return cache_modify
}

var cache_length gopurs_runtime.Value
var once_length sync.Once
func Get_length() gopurs_runtime.Value {
	once_length.Do(func() {
		cache_length = gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn1(), Get_lengthImpl())
	})
	return cache_length
}

var cache_freeze gopurs_runtime.Value
var once_freeze sync.Once
func Get_freeze() gopurs_runtime.Value {
	once_freeze.Do(func() {
		cache_freeze = gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn1(), Get_freezeImpl())
	})
	return cache_freeze
}

var cache_clone gopurs_runtime.Value
var once_clone sync.Once
func Get_clone() gopurs_runtime.Value {
	once_clone.Do(func() {
		cache_clone = gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn1(), Get_cloneImpl())
	})
	return cache_clone
}

func Call_withArray(f_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
__local_var_2_0 := gopurs_runtime.Apply(Get_thaw(), xs_1)
_ = __local_var_2_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
result_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
_ = result_3_1
_dollar__unused_4_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, result_3_1), gopurs_runtime.Value{})
_ = _dollar__unused_4_2
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_unsafeFreeze(), result_3_1), gopurs_runtime.Value{})
})
}

func Call_sortWith(dictOrd_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply(Get_sortBy(), gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), gopurs_runtime.Apply(f_1, x_2), gopurs_runtime.Apply(f_1, y_3))
}))
}

func Call_modify(i_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var i_0 gopurs_runtime.Value = i_0_loop
_ = i_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
__local_var_3_0 := gopurs_runtime.Apply2(Get_peek(), i_0, xs_2)
_ = __local_var_3_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
entry_4_1 := gopurs_runtime.Apply(__local_var_3_0, gopurs_runtime.Value{})
_ = entry_4_1
var __t2 gopurs_runtime.Value
{
if (entry_4_1.Type == 9 && entry_4_1.IntVal == 930809136) {
__t2 = gopurs_runtime.Apply3(Get_poke(), i_0, gopurs_runtime.Apply(f_1, (*pkg_Data_Maybe.Data_Data_Maybe_Just)(entry_4_1.UnsafePtr).V0), xs_2)
goto end_branch_2
} else {

}
}
{
if (entry_4_1.Type == 9 && entry_4_1.IntVal == 3589588149) {
__t2 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(false)
})
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Apply(__t2, gopurs_runtime.Value{})
})
}

func Get_cloneImpl() gopurs_runtime.Value {
	return _Gopurs_CloneImpl
}

func Get_freezeImpl() gopurs_runtime.Value {
	return _Gopurs_FreezeImpl
}

func Get_lengthImpl() gopurs_runtime.Value {
	return _Gopurs_LengthImpl
}

func Get_new_() gopurs_runtime.Value {
	return _Gopurs_New_
}

func Get_peekImpl() gopurs_runtime.Value {
	return _Gopurs_PeekImpl
}

func Get_pokeImpl() gopurs_runtime.Value {
	return _Gopurs_PokeImpl
}

func Get_popImpl() gopurs_runtime.Value {
	return _Gopurs_PopImpl
}

func Get_pushAllImpl() gopurs_runtime.Value {
	return _Gopurs_PushAllImpl
}

func Get_pushImpl() gopurs_runtime.Value {
	return _Gopurs_PushImpl
}

func Get_shiftImpl() gopurs_runtime.Value {
	return _Gopurs_ShiftImpl
}

func Get_sortByImpl() gopurs_runtime.Value {
	return _Gopurs_SortByImpl
}

func Get_spliceImpl() gopurs_runtime.Value {
	return _Gopurs_SpliceImpl
}

func Get_thawImpl() gopurs_runtime.Value {
	return _Gopurs_ThawImpl
}

func Get_toAssocArrayImpl() gopurs_runtime.Value {
	return _Gopurs_ToAssocArrayImpl
}

func Get_unsafeFreezeImpl() gopurs_runtime.Value {
	return _Gopurs_UnsafeFreezeImpl
}

func Get_unsafeThawImpl() gopurs_runtime.Value {
	return _Gopurs_UnsafeThawImpl
}

func Get_unshiftAllImpl() gopurs_runtime.Value {
	return _Gopurs_UnshiftAllImpl
}
