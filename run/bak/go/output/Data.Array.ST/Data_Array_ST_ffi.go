package Data_Array_ST

import "gopurs/output/gopurs_runtime"

func New_(_ interface{}) interface{} {
	panic("Not implemented: new_")
}

func PeekImpl(_ interface{}, _ interface{}, _ interface{}, _ interface{}) interface{} {
	panic("Not implemented: peekImpl")
}

func PokeImpl(_ interface{}, _ interface{}, _ interface{}) interface{} {
	panic("Not implemented: pokeImpl")
}

func LengthImpl(_ interface{}) interface{} {
	panic("Not implemented: lengthImpl")
}

func PopImpl(_ interface{}, _ interface{}, _ interface{}) interface{} {
	panic("Not implemented: popImpl")
}

func PushAllImpl(_ interface{}, _ interface{}) interface{} {
	panic("Not implemented: pushAllImpl")
}

func ShiftImpl(_ interface{}, _ interface{}, _ interface{}) interface{} {
	panic("Not implemented: shiftImpl")
}

func UnshiftAllImpl(_ interface{}, _ interface{}) interface{} {
	panic("Not implemented: unshiftAllImpl")
}

func SpliceImpl(_ interface{}, _ interface{}, _ interface{}, _ interface{}) interface{} {
	panic("Not implemented: spliceImpl")
}

func UnsafeFreezeImpl(_ interface{}) interface{} {
	panic("Not implemented: unsafeFreezeImpl")
}

func UnsafeThawImpl(_ interface{}) interface{} {
	panic("Not implemented: unsafeThawImpl")
}

func FreezeImpl(_ interface{}) interface{} {
	panic("Not implemented: freezeImpl")
}

func ThawImpl(_ interface{}) interface{} {
	panic("Not implemented: thawImpl")
}

func CloneImpl(_ interface{}) interface{} {
	panic("Not implemented: cloneImpl")
}

func SortByImpl(_ interface{}, _ interface{}, _ interface{}) interface{} {
	panic("Not implemented: sortByImpl")
}

func ToAssocArrayImpl(_ interface{}) interface{} {
	panic("Not implemented: toAssocArrayImpl")
}

func PushImpl(_ interface{}, _ interface{}) interface{} {
	panic("Not implemented: pushImpl")
}


// --- Auto-generated FFI wrappers ---
func Call_new_(arg0 interface{}) interface{} {
	return New_(arg0)
}
var _Gopurs_New_ = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := New_(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_peekImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}, arg3 interface{}) interface{} {
	return PeekImpl(arg0, arg1, arg2, arg3)
}
var _Gopurs_PeekImpl = gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_res := PeekImpl(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
func Call_pokeImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}) interface{} {
	return PokeImpl(arg0, arg1, arg2)
}
var _Gopurs_PokeImpl = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := PokeImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
func Call_lengthImpl(arg0 interface{}) interface{} {
	return LengthImpl(arg0)
}
var _Gopurs_LengthImpl = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := LengthImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_popImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}) interface{} {
	return PopImpl(arg0, arg1, arg2)
}
var _Gopurs_PopImpl = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := PopImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
func Call_pushAllImpl(arg0 interface{}, arg1 interface{}) interface{} {
	return PushAllImpl(arg0, arg1)
}
var _Gopurs_PushAllImpl = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := PushAllImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call_shiftImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}) interface{} {
	return ShiftImpl(arg0, arg1, arg2)
}
var _Gopurs_ShiftImpl = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := ShiftImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
func Call_unshiftAllImpl(arg0 interface{}, arg1 interface{}) interface{} {
	return UnshiftAllImpl(arg0, arg1)
}
var _Gopurs_UnshiftAllImpl = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := UnshiftAllImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call_spliceImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}, arg3 interface{}) interface{} {
	return SpliceImpl(arg0, arg1, arg2, arg3)
}
var _Gopurs_SpliceImpl = gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_res := SpliceImpl(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
func Call_unsafeFreezeImpl(arg0 interface{}) interface{} {
	return UnsafeFreezeImpl(arg0)
}
var _Gopurs_UnsafeFreezeImpl = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := UnsafeFreezeImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_unsafeThawImpl(arg0 interface{}) interface{} {
	return UnsafeThawImpl(arg0)
}
var _Gopurs_UnsafeThawImpl = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := UnsafeThawImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_freezeImpl(arg0 interface{}) interface{} {
	return FreezeImpl(arg0)
}
var _Gopurs_FreezeImpl = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := FreezeImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_thawImpl(arg0 interface{}) interface{} {
	return ThawImpl(arg0)
}
var _Gopurs_ThawImpl = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := ThawImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_cloneImpl(arg0 interface{}) interface{} {
	return CloneImpl(arg0)
}
var _Gopurs_CloneImpl = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := CloneImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_sortByImpl(arg0 interface{}, arg1 interface{}, arg2 interface{}) interface{} {
	return SortByImpl(arg0, arg1, arg2)
}
var _Gopurs_SortByImpl = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := SortByImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
func Call_toAssocArrayImpl(arg0 interface{}) interface{} {
	return ToAssocArrayImpl(arg0)
}
var _Gopurs_ToAssocArrayImpl = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := ToAssocArrayImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_pushImpl(arg0 interface{}, arg1 interface{}) interface{} {
	return PushImpl(arg0, arg1)
}
var _Gopurs_PushImpl = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := PushImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
