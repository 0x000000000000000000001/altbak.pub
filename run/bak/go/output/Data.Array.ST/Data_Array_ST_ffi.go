package Data_Array_ST

import "gopurs/output/gopurs_runtime"

func New_() func() interface{} {
	return func() interface{} {
		arr := make([]interface{}, 0)
		return &arr
	}
}

func PeekImpl(just func(interface{}) interface{}, nothing interface{}, i int64, arr interface{}) func() interface{} {
	return func() interface{} {
		a := arr.(*[]interface{})
		if i >= 0 && i < int64(len(*a)) {
			return just((*a)[i])
		}
		return nothing
	}
}

func PokeImpl(i int64, a interface{}, arr interface{}) func() interface{} {
	return func() interface{} {
		ptr := arr.(*[]interface{})
		if i >= 0 && i < int64(len(*ptr)) {
			(*ptr)[i] = a
			return true
		}
		return false
	}
}

func LengthImpl(arr interface{}) func() interface{} {
	return func() interface{} {
		ptr := arr.(*[]interface{})
		return int64(len(*ptr))
	}
}

func PopImpl(just func(interface{}) interface{}, nothing interface{}, arr interface{}) func() interface{} {
	return func() interface{} {
		ptr := arr.(*[]interface{})
		if len(*ptr) > 0 {
			last := (*ptr)[len(*ptr)-1]
			*ptr = (*ptr)[:len(*ptr)-1]
			return just(last)
		}
		return nothing
	}
}

func PushAllImpl(xs []interface{}, arr interface{}) func() interface{} {
	return func() interface{} {
		ptr := arr.(*[]interface{})
		*ptr = append(*ptr, xs...)
		return int64(len(*ptr))
	}
}

func PushImpl(x interface{}, arr interface{}) func() interface{} {
	return func() interface{} {
		ptr := arr.(*[]interface{})
		*ptr = append(*ptr, x)
		return int64(len(*ptr))
	}
}

func ShiftImpl(just func(interface{}) interface{}, nothing interface{}, arr interface{}) func() interface{} {
	return func() interface{} {
		ptr := arr.(*[]interface{})
		if len(*ptr) > 0 {
			first := (*ptr)[0]
			*ptr = (*ptr)[1:]
			return just(first)
		}
		return nothing
	}
}

func UnshiftAllImpl(xs []interface{}, arr interface{}) func() interface{} {
	return func() interface{} {
		ptr := arr.(*[]interface{})
		*ptr = append(xs, *ptr...)
		return int64(len(*ptr))
	}
}

func UnshiftImpl(x interface{}, arr interface{}) func() interface{} {
	return func() interface{} {
		ptr := arr.(*[]interface{})
		*ptr = append([]interface{}{x}, *ptr...)
		return int64(len(*ptr))
	}
}

func SpliceImpl(start int64, count int64, xs []interface{}, arr interface{}) func() interface{} {
	return func() interface{} {
		ptr := arr.(*[]interface{})
		removed := make([]interface{}, count)
		copy(removed, (*ptr)[start:start+count])
		
		newArr := make([]interface{}, 0, len(*ptr) - int(count) + len(xs))
		newArr = append(newArr, (*ptr)[:start]...)
		newArr = append(newArr, xs...)
		newArr = append(newArr, (*ptr)[start+count:]...)
		*ptr = newArr
		return removed
	}
}

func UnsafeFreezeImpl(arr interface{}) func() interface{} {
	return func() interface{} {
		return *(arr.(*[]interface{}))
	}
}

func UnsafeThawImpl(xs []interface{}) func() interface{} {
	return func() interface{} {
		return &xs
	}
}

func FreezeImpl(arr interface{}) func() interface{} {
	return func() interface{} {
		ptr := arr.(*[]interface{})
		res := make([]interface{}, len(*ptr))
		copy(res, *ptr)
		return res
	}
}

func ThawImpl(xs []interface{}) func() interface{} {
	return func() interface{} {
		res := make([]interface{}, len(xs))
		copy(res, xs)
		return &res
	}
}

func CloneImpl(arr interface{}) func() interface{} {
	return func() interface{} {
		ptr := arr.(*[]interface{})
		res := make([]interface{}, len(*ptr))
		copy(res, *ptr)
		return &res
	}
}

func SortByImpl(f func(interface{}, interface{}) interface{}, toInt func(interface{}) int64, arr interface{}) func() interface{} {
	return func() interface{} {
		panic("Not implemented: sortByImpl")
	}
}

func ToAssocArrayImpl(arr interface{}) func() interface{} {
	return func() interface{} {
		panic("Not implemented: toAssocArrayImpl")
	}
}


// --- Auto-generated FFI wrappers ---
func Call_new_() func() interface{} {
	return New_()
}
var _Gopurs_New_ = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
	go_res := New_()
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_peekImpl(arg0 func(interface{}) interface{}, arg1 interface{}, arg2 int64, arg3 interface{}) func() interface{} {
	return PeekImpl(arg0, arg1, arg2, arg3)
}
var _Gopurs_PeekImpl = gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 interface{}) interface{} {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[int64](arg2)
	go_arg3 := arg3
	go_res := PeekImpl(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_pokeImpl(arg0 int64, arg1 interface{}, arg2 interface{}) func() interface{} {
	return PokeImpl(arg0, arg1, arg2)
}
var _Gopurs_PokeImpl = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int64](arg0)
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := PokeImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_lengthImpl(arg0 interface{}) func() interface{} {
	return LengthImpl(arg0)
}
var _Gopurs_LengthImpl = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := LengthImpl(go_arg0)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_popImpl(arg0 func(interface{}) interface{}, arg1 interface{}, arg2 interface{}) func() interface{} {
	return PopImpl(arg0, arg1, arg2)
}
var _Gopurs_PopImpl = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 interface{}) interface{} {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := PopImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_pushAllImpl(arg0 []interface{}, arg1 interface{}) func() interface{} {
	return PushAllImpl(arg0, arg1)
}
var _Gopurs_PushAllImpl = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	arg0_arr := *(*[]gopurs_runtime.Value)(arg0.UnsafePtr)
	go_arg0 := make([]interface{}, len(arg0_arr))
	for i, v := range arg0_arr { go_arg0[i] = v }
	go_arg1 := arg1
	go_res := PushAllImpl(go_arg0, go_arg1)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_pushImpl(arg0 interface{}, arg1 interface{}) func() interface{} {
	return PushImpl(arg0, arg1)
}
var _Gopurs_PushImpl = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := PushImpl(go_arg0, go_arg1)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_shiftImpl(arg0 func(interface{}) interface{}, arg1 interface{}, arg2 interface{}) func() interface{} {
	return ShiftImpl(arg0, arg1, arg2)
}
var _Gopurs_ShiftImpl = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 interface{}) interface{} {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := ShiftImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_unshiftAllImpl(arg0 []interface{}, arg1 interface{}) func() interface{} {
	return UnshiftAllImpl(arg0, arg1)
}
var _Gopurs_UnshiftAllImpl = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	arg0_arr := *(*[]gopurs_runtime.Value)(arg0.UnsafePtr)
	go_arg0 := make([]interface{}, len(arg0_arr))
	for i, v := range arg0_arr { go_arg0[i] = v }
	go_arg1 := arg1
	go_res := UnshiftAllImpl(go_arg0, go_arg1)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_unshiftImpl(arg0 interface{}, arg1 interface{}) func() interface{} {
	return UnshiftImpl(arg0, arg1)
}
var _Gopurs_UnshiftImpl = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := UnshiftImpl(go_arg0, go_arg1)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_spliceImpl(arg0 int64, arg1 int64, arg2 []interface{}, arg3 interface{}) func() interface{} {
	return SpliceImpl(arg0, arg1, arg2, arg3)
}
var _Gopurs_SpliceImpl = gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int64](arg0)
	go_arg1 := gopurs_runtime.Unbox[int64](arg1)
	arg2_arr := *(*[]gopurs_runtime.Value)(arg2.UnsafePtr)
	go_arg2 := make([]interface{}, len(arg2_arr))
	for i, v := range arg2_arr { go_arg2[i] = v }
	go_arg3 := arg3
	go_res := SpliceImpl(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_unsafeFreezeImpl(arg0 interface{}) func() interface{} {
	return UnsafeFreezeImpl(arg0)
}
var _Gopurs_UnsafeFreezeImpl = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := UnsafeFreezeImpl(go_arg0)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_unsafeThawImpl(arg0 []interface{}) func() interface{} {
	return UnsafeThawImpl(arg0)
}
var _Gopurs_UnsafeThawImpl = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	arg0_arr := *(*[]gopurs_runtime.Value)(arg0.UnsafePtr)
	go_arg0 := make([]interface{}, len(arg0_arr))
	for i, v := range arg0_arr { go_arg0[i] = v }
	go_res := UnsafeThawImpl(go_arg0)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_freezeImpl(arg0 interface{}) func() interface{} {
	return FreezeImpl(arg0)
}
var _Gopurs_FreezeImpl = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := FreezeImpl(go_arg0)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_thawImpl(arg0 []interface{}) func() interface{} {
	return ThawImpl(arg0)
}
var _Gopurs_ThawImpl = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	arg0_arr := *(*[]gopurs_runtime.Value)(arg0.UnsafePtr)
	go_arg0 := make([]interface{}, len(arg0_arr))
	for i, v := range arg0_arr { go_arg0[i] = v }
	go_res := ThawImpl(go_arg0)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_cloneImpl(arg0 interface{}) func() interface{} {
	return CloneImpl(arg0)
}
var _Gopurs_CloneImpl = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := CloneImpl(go_arg0)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_sortByImpl(arg0 func(interface{}, interface{}) interface{}, arg1 func(interface{}) int64, arg2 interface{}) func() interface{} {
	return SortByImpl(arg0, arg1, arg2)
}
var _Gopurs_SortByImpl = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 interface{}, p0_1 interface{}) interface{} {
			return gopurs_runtime.Apply2(arg0, gopurs_runtime.Box(p0_0), gopurs_runtime.Box(p0_1))
		}
	go_arg1 := func(p0_0 interface{}) int64 {
			inner_res0 := gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
			return gopurs_runtime.Unbox[int64](inner_res0)
		}
	go_arg2 := arg2
	go_res := SortByImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_toAssocArrayImpl(arg0 interface{}) func() interface{} {
	return ToAssocArrayImpl(arg0)
}
var _Gopurs_ToAssocArrayImpl = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := ToAssocArrayImpl(go_arg0)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
