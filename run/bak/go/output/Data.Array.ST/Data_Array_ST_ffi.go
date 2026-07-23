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
var _Gopurs_New_ = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0.PtrVal
	go_res := New_(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_PeekImpl = gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0.PtrVal
	go_arg1 := arg1.PtrVal
	go_arg2 := arg2.PtrVal
	go_arg3 := arg3.PtrVal
	go_res := PeekImpl(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_PokeImpl = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0.PtrVal
	go_arg1 := arg1.PtrVal
	go_arg2 := arg2.PtrVal
	go_res := PokeImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_LengthImpl = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0.PtrVal
	go_res := LengthImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_PopImpl = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0.PtrVal
	go_arg1 := arg1.PtrVal
	go_arg2 := arg2.PtrVal
	go_res := PopImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_PushAllImpl = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0.PtrVal
	go_arg1 := arg1.PtrVal
	go_res := PushAllImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_ShiftImpl = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0.PtrVal
	go_arg1 := arg1.PtrVal
	go_arg2 := arg2.PtrVal
	go_res := ShiftImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_UnshiftAllImpl = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0.PtrVal
	go_arg1 := arg1.PtrVal
	go_res := UnshiftAllImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_SpliceImpl = gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0.PtrVal
	go_arg1 := arg1.PtrVal
	go_arg2 := arg2.PtrVal
	go_arg3 := arg3.PtrVal
	go_res := SpliceImpl(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_UnsafeFreezeImpl = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0.PtrVal
	go_res := UnsafeFreezeImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_UnsafeThawImpl = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0.PtrVal
	go_res := UnsafeThawImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_FreezeImpl = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0.PtrVal
	go_res := FreezeImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_ThawImpl = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0.PtrVal
	go_res := ThawImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_CloneImpl = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0.PtrVal
	go_res := CloneImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_SortByImpl = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0.PtrVal
	go_arg1 := arg1.PtrVal
	go_arg2 := arg2.PtrVal
	go_res := SortByImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_ToAssocArrayImpl = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0.PtrVal
	go_res := ToAssocArrayImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_PushImpl = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0.PtrVal
	go_arg1 := arg1.PtrVal
	go_res := PushImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
