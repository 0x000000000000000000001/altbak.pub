package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Op_Op gopurs_runtime.Value
var once_Data_Op_Op sync.Once
func Get_Data_Op_Op() gopurs_runtime.Value {
	once_Data_Op_Op.Do(func() {
		cache_Data_Op_Op = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Op_Op(x_0_box)
})
	})
	return cache_Data_Op_Op
}

var cache_Data_Op_semigroupoidOp gopurs_runtime.Value
var once_Data_Op_semigroupoidOp sync.Once
func Get_Data_Op_semigroupoidOp() gopurs_runtime.Value {
	once_Data_Op_semigroupoidOp.Do(func() {
		cache_Data_Op_semigroupoidOp = gopurs_runtime.Value{Type: 9, IntVal: 350442445, UnsafePtr: unsafe.Pointer(&Constructor_Control_Semigroupoid_Semigroupoid{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_1, gopurs_runtime.Apply(v_0, x_2))
})
})
})})}
	})
	return cache_Data_Op_semigroupoidOp
}

var cache_Data_Op_semigroupOp gopurs_runtime.Value
var once_Data_Op_semigroupOp sync.Once
func Get_Data_Op_semigroupOp() gopurs_runtime.Value {
	once_Data_Op_semigroupOp.Do(func() {
		cache_Data_Op_semigroupOp = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Op_semigroupOp(dictSemigroup_0_box)
})
	})
	return cache_Data_Op_semigroupOp
}

var cache_Data_Op_newtypeOp gopurs_runtime.Value
var once_Data_Op_newtypeOp sync.Once
func Get_Data_Op_newtypeOp() gopurs_runtime.Value {
	once_Data_Op_newtypeOp.Do(func() {
		cache_Data_Op_newtypeOp = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(&Constructor_Data_Newtype_Newtype{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})}
	})
	return cache_Data_Op_newtypeOp
}

var cache_Data_Op_monoidOp gopurs_runtime.Value
var once_Data_Op_monoidOp sync.Once
func Get_Data_Op_monoidOp() gopurs_runtime.Value {
	once_Data_Op_monoidOp.Do(func() {
		cache_Data_Op_monoidOp = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Op_monoidOp(dictMonoid_0_box)
})
	})
	return cache_Data_Op_monoidOp
}

var cache_Data_Op_contravariantOp gopurs_runtime.Value
var once_Data_Op_contravariantOp sync.Once
func Get_Data_Op_contravariantOp() gopurs_runtime.Value {
	once_Data_Op_contravariantOp.Do(func() {
		cache_Data_Op_contravariantOp = gopurs_runtime.Value{Type: 9, IntVal: 85171506, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Contravariant_Contravariant{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(f_0, x_2))
})
})
})})}
	})
	return cache_Data_Op_contravariantOp
}

var cache_Data_Op_categoryOp gopurs_runtime.Value
var once_Data_Op_categoryOp sync.Once
func Get_Data_Op_categoryOp() gopurs_runtime.Value {
	once_Data_Op_categoryOp.Do(func() {
		cache_Data_Op_categoryOp = gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(&Constructor_Control_Category_Category{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 350442445, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](Get_Data_Op_semigroupoidOp()))}
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})})}
	})
	return cache_Data_Op_categoryOp
}

var cache_Data_Op_contravariantOp__1891937874 gopurs_runtime.Value
var once_Data_Op_contravariantOp__1891937874 sync.Once
func Get_Data_Op_contravariantOp__1891937874() gopurs_runtime.Value {
	once_Data_Op_contravariantOp__1891937874.Do(func() {
		cache_Data_Op_contravariantOp__1891937874 = gopurs_runtime.Value{Type: 9, IntVal: 85171506, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Contravariant_Contravariant{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(f_0, x_2))
})
})
})})}
	})
	return cache_Data_Op_contravariantOp__1891937874
}

var cache_Data_Op_semigroupoidOp__1787481579 gopurs_runtime.Value
var once_Data_Op_semigroupoidOp__1787481579 sync.Once
func Get_Data_Op_semigroupoidOp__1787481579() gopurs_runtime.Value {
	once_Data_Op_semigroupoidOp__1787481579.Do(func() {
		cache_Data_Op_semigroupoidOp__1787481579 = gopurs_runtime.Value{Type: 9, IntVal: 350442445, UnsafePtr: unsafe.Pointer(&Constructor_Control_Semigroupoid_Semigroupoid{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_1, gopurs_runtime.Apply(v_0, x_2))
})
})
})})}
	})
	return cache_Data_Op_semigroupoidOp__1787481579
}

func Call_Data_Op_Op(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Op_semigroupOp(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), gopurs_runtime.Apply(f_1, x_3), gopurs_runtime.Apply(g_2, x_3))
})
})
})})}
}

func Call_Data_Op_monoidOp(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): semigroupFn_1_0 -> *Constructor_Data_Semigroup_Semigroup
semigroupFn_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "append"), gopurs_runtime.Apply(f_2, x_4), gopurs_runtime.Apply(g_3, x_4))
})
})
})))
_ = semigroupFn_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(&Constructor_Data_Monoid_Monoid{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupFn_1_0)}
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
})})}
}


