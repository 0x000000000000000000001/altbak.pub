package Data_Array_ST

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Monad_ST_Uncurried "gopurs/output/Control.Monad.ST.Uncurried"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Control_Monad_ST_Internal "gopurs/output/Control.Monad.ST.Internal"
)

var unshiftAll gopurs_runtime.Value
var once_unshiftAll sync.Once
func Get_unshiftAll() gopurs_runtime.Value {
	once_unshiftAll.Do(func() {
		unshiftAll = gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn2(), Get_unshiftAllImpl())
	})
	return unshiftAll
}

var unshift gopurs_runtime.Value
var once_unshift sync.Once
func Get_unshift() gopurs_runtime.Value {
	once_unshift.Do(func() {
		unshift = gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Control_Monad_ST_Uncurried.Get_runSTFn2(), Get_unshiftAllImpl(), gopurs_runtime.Array([]gopurs_runtime.Value{a_0}))
})
	})
	return unshift
}

var unsafeThaw gopurs_runtime.Value
var once_unsafeThaw sync.Once
func Get_unsafeThaw() gopurs_runtime.Value {
	once_unsafeThaw.Do(func() {
		unsafeThaw = gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn1(), Get_unsafeThawImpl())
	})
	return unsafeThaw
}

var unsafeFreeze gopurs_runtime.Value
var once_unsafeFreeze sync.Once
func Get_unsafeFreeze() gopurs_runtime.Value {
	once_unsafeFreeze.Do(func() {
		unsafeFreeze = gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn1(), Get_unsafeFreezeImpl())
	})
	return unsafeFreeze
}

var toAssocArray gopurs_runtime.Value
var once_toAssocArray sync.Once
func Get_toAssocArray() gopurs_runtime.Value {
	once_toAssocArray.Do(func() {
		toAssocArray = gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn1(), Get_toAssocArrayImpl())
	})
	return toAssocArray
}

var thaw gopurs_runtime.Value
var once_thaw sync.Once
func Get_thaw() gopurs_runtime.Value {
	once_thaw.Do(func() {
		thaw = gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn1(), Get_thawImpl())
	})
	return thaw
}

var withArray gopurs_runtime.Value
var once_withArray sync.Once
func Get_withArray() gopurs_runtime.Value {
	once_withArray.Do(func() {
		withArray = gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(Get_thaw(), xs_1)
_ = __local_var_2_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
result_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
_ = result_3_1
_dollar__unused_4_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, result_3_1), gopurs_runtime.Value{})
_ = _dollar__unused_4_2
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_unsafeFreeze(), result_3_1), gopurs_runtime.Value{})
})
})
	})
	return withArray
}

var splice gopurs_runtime.Value
var once_splice sync.Once
func Get_splice() gopurs_runtime.Value {
	once_splice.Do(func() {
		splice = gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn4(), Get_spliceImpl())
	})
	return splice
}

var sortBy gopurs_runtime.Value
var once_sortBy sync.Once
func Get_sortBy() gopurs_runtime.Value {
	once_sortBy.Do(func() {
		sortBy = gopurs_runtime.Func(func(comp_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(pkg_Control_Monad_ST_Uncurried.Get_runSTFn3(), Get_sortByImpl(), comp_0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_1.StrVal == "GT").IntVal != 0 {
__t0 = gopurs_runtime.Int(1)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_1.StrVal == "EQ").IntVal != 0 {
__t0 = gopurs_runtime.Int(0)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_1.StrVal == "LT").IntVal != 0 {
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
})
	})
	return sortBy
}

var sortWith gopurs_runtime.Value
var once_sortWith sync.Once
func Get_sortWith() gopurs_runtime.Value {
	once_sortWith.Do(func() {
		sortWith = gopurs_runtime.Func2(func(dictOrd_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_sortBy(), gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), gopurs_runtime.Apply(f_1, x_2), gopurs_runtime.Apply(f_1, y_3))
}))
})
	})
	return sortWith
}

var sort gopurs_runtime.Value
var once_sort sync.Once
func Get_sort() gopurs_runtime.Value {
	once_sort.Do(func() {
		sort = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_sortBy(), gopurs_runtime.RecordGet(dictOrd_0, "compare"))
})
	})
	return sort
}

var shift gopurs_runtime.Value
var once_shift sync.Once
func Get_shift() gopurs_runtime.Value {
	once_shift.Do(func() {
		shift = gopurs_runtime.Apply3(pkg_Control_Monad_ST_Uncurried.Get_runSTFn3(), Get_shiftImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Constructor0("Nothing"))
	})
	return shift
}

var run gopurs_runtime.Value
var once_run sync.Once
func Get_run() gopurs_runtime.Value {
	once_run.Do(func() {
		run = gopurs_runtime.Func(func(st_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(st_0, gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_unsafeFreeze(), __local_var_1_0), gopurs_runtime.Value{})
}))
})
	})
	return run
}

var pushAll gopurs_runtime.Value
var once_pushAll sync.Once
func Get_pushAll() gopurs_runtime.Value {
	once_pushAll.Do(func() {
		pushAll = gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn2(), Get_pushAllImpl())
	})
	return pushAll
}

var push gopurs_runtime.Value
var once_push sync.Once
func Get_push() gopurs_runtime.Value {
	once_push.Do(func() {
		push = gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn2(), Get_pushImpl())
	})
	return push
}

var pop gopurs_runtime.Value
var once_pop sync.Once
func Get_pop() gopurs_runtime.Value {
	once_pop.Do(func() {
		pop = gopurs_runtime.Apply3(pkg_Control_Monad_ST_Uncurried.Get_runSTFn3(), Get_popImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Constructor0("Nothing"))
	})
	return pop
}

var poke gopurs_runtime.Value
var once_poke sync.Once
func Get_poke() gopurs_runtime.Value {
	once_poke.Do(func() {
		poke = gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn3(), Get_pokeImpl())
	})
	return poke
}

var peek gopurs_runtime.Value
var once_peek sync.Once
func Get_peek() gopurs_runtime.Value {
	once_peek.Do(func() {
		peek = gopurs_runtime.Apply3(pkg_Control_Monad_ST_Uncurried.Get_runSTFn4(), Get_peekImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Constructor0("Nothing"))
	})
	return peek
}

var modify gopurs_runtime.Value
var once_modify sync.Once
func Get_modify() gopurs_runtime.Value {
	once_modify.Do(func() {
		modify = gopurs_runtime.Func3(func(i_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Apply2(Get_peek(), i_0, xs_2)
_ = __local_var_3_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
entry_4_1 := gopurs_runtime.Apply(__local_var_3_0, gopurs_runtime.Value{})
_ = entry_4_1
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool(entry_4_1.StrVal == "Just").IntVal != 0 {
__t2 = gopurs_runtime.Apply3(Get_poke(), i_0, gopurs_runtime.Apply(f_1, (*[1024]gopurs_runtime.Value)(entry_4_1.UnsafePtr)[0]), xs_2)
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool(entry_4_1.StrVal == "Nothing").IntVal != 0 {
__t2 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Boolean(false)
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
})
	})
	return modify
}

var length gopurs_runtime.Value
var once_length sync.Once
func Get_length() gopurs_runtime.Value {
	once_length.Do(func() {
		length = gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn1(), Get_lengthImpl())
	})
	return length
}

var freeze gopurs_runtime.Value
var once_freeze sync.Once
func Get_freeze() gopurs_runtime.Value {
	once_freeze.Do(func() {
		freeze = gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn1(), Get_freezeImpl())
	})
	return freeze
}

var clone gopurs_runtime.Value
var once_clone sync.Once
func Get_clone() gopurs_runtime.Value {
	once_clone.Do(func() {
		clone = gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn1(), Get_cloneImpl())
	})
	return clone
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
