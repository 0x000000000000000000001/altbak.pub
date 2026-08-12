package Data_Date_Gen

import (
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Control_Monad_Gen_Class "gopurs/output/Control.Monad.Gen.Class"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Date "gopurs/output/Data.Date"
	pkg_Data_Date_Component "gopurs/output/Data.Date.Component"
	pkg_Data_Date_Component_Gen "gopurs/output/Data.Date.Component.Gen"
	pkg_Data_Enum "gopurs/output/Data.Enum"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Int "gopurs/output/Data.Int"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Partial_Unsafe "gopurs/output/Partial.Unsafe"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_genDate gopurs_runtime.Value
var once_genDate sync.Once
func Get_genDate() gopurs_runtime.Value {
	once_genDate.Do(func() {
		cache_genDate = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genDate(dictMonadGen_0_box)
})
	})
	return cache_genDate
}

var cache_pure__779048344 gopurs_runtime.Value
var once_pure__779048344 sync.Once
func Get_pure__779048344() gopurs_runtime.Value {
	once_pure__779048344.Do(func() {
		cache_pure__779048344 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__779048344(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_pure__779048344
}

var cache_pure__4037597590 gopurs_runtime.Value
var once_pure__4037597590 sync.Once
func Get_pure__4037597590() gopurs_runtime.Value {
	once_pure__4037597590.Do(func() {
		cache_pure__4037597590 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__4037597590(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__4037597590
}

var cache_pure__3215807376 gopurs_runtime.Value
var once_pure__3215807376 sync.Once
func Get_pure__3215807376() gopurs_runtime.Value {
	once_pure__3215807376.Do(func() {
		cache_pure__3215807376 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__3215807376(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__3215807376
}

var cache_pure__3181299446 gopurs_runtime.Value
var once_pure__3181299446 sync.Once
func Get_pure__3181299446() gopurs_runtime.Value {
	once_pure__3181299446.Do(func() {
		cache_pure__3181299446 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__3181299446(__eta0_0_box)
})
	})
	return cache_pure__3181299446
}

var cache_apply__2174094052 gopurs_runtime.Value
var once_apply__2174094052 sync.Once
func Get_apply__2174094052() gopurs_runtime.Value {
	once_apply__2174094052.Do(func() {
		cache_apply__2174094052 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__2174094052(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_apply__2174094052
}

var cache_apply__652936900 gopurs_runtime.Value
var once_apply__652936900 sync.Once
func Get_apply__652936900() gopurs_runtime.Value {
	once_apply__652936900.Do(func() {
		cache_apply__652936900 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__652936900(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_apply__652936900
}

var cache_apply__1982519530 gopurs_runtime.Value
var once_apply__1982519530 sync.Once
func Get_apply__1982519530() gopurs_runtime.Value {
	once_apply__1982519530.Do(func() {
		cache_apply__1982519530 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__1982519530(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_apply__1982519530
}

var cache_apply__353515660 gopurs_runtime.Value
var once_apply__353515660 sync.Once
func Get_apply__353515660() gopurs_runtime.Value {
	once_apply__353515660.Do(func() {
		cache_apply__353515660 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__353515660(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_apply__353515660
}

var cache_apply__1471729482 gopurs_runtime.Value
var once_apply__1471729482 sync.Once
func Get_apply__1471729482() gopurs_runtime.Value {
	once_apply__1471729482.Do(func() {
		cache_apply__1471729482 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_apply__1471729482(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](v1_1_box)))}
})
	})
	return cache_apply__1471729482
}

var cache_apply__1572009162 gopurs_runtime.Value
var once_apply__1572009162 sync.Once
func Get_apply__1572009162() gopurs_runtime.Value {
	once_apply__1572009162.Do(func() {
		cache_apply__1572009162 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_apply__1572009162(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[uint32]](v1_1_box)))}
})
	})
	return cache_apply__1572009162
}

var cache_bind__3254602343 gopurs_runtime.Value
var once_bind__3254602343 sync.Once
func Get_bind__3254602343() gopurs_runtime.Value {
	once_bind__3254602343.Do(func() {
		cache_bind__3254602343 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__3254602343(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__3254602343
}

var cache_bind__4011257415 gopurs_runtime.Value
var once_bind__4011257415 sync.Once
func Get_bind__4011257415() gopurs_runtime.Value {
	once_bind__4011257415.Do(func() {
		cache_bind__4011257415 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__4011257415(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__4011257415
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

var cache_bind__4062037089 gopurs_runtime.Value
var once_bind__4062037089 sync.Once
func Get_bind__4062037089() gopurs_runtime.Value {
	once_bind__4062037089.Do(func() {
		cache_bind__4062037089 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_bind__4062037089(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](v_0_box), v1_1_box))}
})
	})
	return cache_bind__4062037089
}

var cache_bind__3444930753 gopurs_runtime.Value
var once_bind__3444930753 sync.Once
func Get_bind__3444930753() gopurs_runtime.Value {
	once_bind__3444930753.Do(func() {
		cache_bind__3444930753 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_bind__3444930753(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Date.Constructor_Date]](v_0_box), v1_1_box))}
})
	})
	return cache_bind__3444930753
}

var cache_bindFlipped__1485397639 gopurs_runtime.Value
var once_bindFlipped__1485397639 sync.Once
func Get_bindFlipped__1485397639() gopurs_runtime.Value {
	once_bindFlipped__1485397639.Do(func() {
		cache_bindFlipped__1485397639 = gopurs_runtime.Func3(func(dictBind_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindFlipped__1485397639(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dictBind_0_box), b_1_box, a_2_box)
})
	})
	return cache_bindFlipped__1485397639
}

var cache_bindFlipped__3917280577 gopurs_runtime.Value
var once_bindFlipped__3917280577 sync.Once
func Get_bindFlipped__3917280577() gopurs_runtime.Value {
	once_bindFlipped__3917280577.Do(func() {
		cache_bindFlipped__3917280577 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindFlipped__3917280577(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_bindFlipped__3917280577
}

var cache_bindFlipped__1454086721 gopurs_runtime.Value
var once_bindFlipped__1454086721 sync.Once
func Get_bindFlipped__1454086721() gopurs_runtime.Value {
	once_bindFlipped__1454086721.Do(func() {
		cache_bindFlipped__1454086721 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindFlipped__1454086721(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_bindFlipped__1454086721
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

var cache_compose__1636404804 gopurs_runtime.Value
var once_compose__1636404804 sync.Once
func Get_compose__1636404804() gopurs_runtime.Value {
	once_compose__1636404804.Do(func() {
		cache_compose__1636404804 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__1636404804(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[uint32]](dict_0_box))
})
	})
	return cache_compose__1636404804
}

var cache_compose__1774599291 gopurs_runtime.Value
var once_compose__1774599291 sync.Once
func Get_compose__1774599291() gopurs_runtime.Value {
	once_compose__1774599291.Do(func() {
		cache_compose__1774599291 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__1774599291(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[int64]](dict_0_box))
})
	})
	return cache_compose__1774599291
}

var cache_compose__1543665403 gopurs_runtime.Value
var once_compose__1543665403 sync.Once
func Get_compose__1543665403() gopurs_runtime.Value {
	once_compose__1543665403.Do(func() {
		cache_compose__1543665403 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__1543665403(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[int64]](dict_0_box))
})
	})
	return cache_compose__1543665403
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

var cache_fromEnum__3599151655 gopurs_runtime.Value
var once_fromEnum__3599151655 sync.Once
func Get_fromEnum__3599151655() gopurs_runtime.Value {
	once_fromEnum__3599151655.Do(func() {
		cache_fromEnum__3599151655 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_fromEnum__3599151655(v_0_box.IntVal))
})
	})
	return cache_fromEnum__3599151655
}

var cache_fromEnum__1637084359 gopurs_runtime.Value
var once_fromEnum__1637084359 sync.Once
func Get_fromEnum__1637084359() gopurs_runtime.Value {
	once_fromEnum__1637084359.Do(func() {
		cache_fromEnum__1637084359 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromEnum__1637084359(gopurs_runtime.CoerceToStruct[pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_fromEnum__1637084359
}

var cache_fromEnum__1196942535 gopurs_runtime.Value
var once_fromEnum__1196942535 sync.Once
func Get_fromEnum__1196942535() gopurs_runtime.Value {
	once_fromEnum__1196942535.Do(func() {
		cache_fromEnum__1196942535 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_fromEnum__1196942535(uint32(v_0_box.IntVal)))
})
	})
	return cache_fromEnum__1196942535
}

var cache_pred__2914940949 gopurs_runtime.Value
var once_pred__2914940949 sync.Once
func Get_pred__2914940949() gopurs_runtime.Value {
	once_pred__2914940949.Do(func() {
		cache_pred__2914940949 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pred__2914940949(__eta0_0_box)
})
	})
	return cache_pred__2914940949
}

var cache_pred__3199041328 gopurs_runtime.Value
var once_pred__3199041328 sync.Once
func Get_pred__3199041328() gopurs_runtime.Value {
	once_pred__3199041328.Do(func() {
		cache_pred__3199041328 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pred__3199041328(gopurs_runtime.CoerceToStruct[pkg_Data_Enum.Constructor_Enum[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pred__3199041328
}

var cache_pred__2010692236 gopurs_runtime.Value
var once_pred__2010692236 sync.Once
func Get_pred__2010692236() gopurs_runtime.Value {
	once_pred__2010692236.Do(func() {
		cache_pred__2010692236 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pred__2010692236(__eta0_0_box)
})
	})
	return cache_pred__2010692236
}

var cache_succ__412946465 gopurs_runtime.Value
var once_succ__412946465 sync.Once
func Get_succ__412946465() gopurs_runtime.Value {
	once_succ__412946465.Do(func() {
		cache_succ__412946465 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_succ__412946465(gopurs_runtime.CoerceToStruct[pkg_Data_Enum.Constructor_Enum[*pkg_Data_Date.Constructor_Date]](dict_0_box))
})
	})
	return cache_succ__412946465
}

var cache_succ__2914940949 gopurs_runtime.Value
var once_succ__2914940949 sync.Once
func Get_succ__2914940949() gopurs_runtime.Value {
	once_succ__2914940949.Do(func() {
		cache_succ__2914940949 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_succ__2914940949(__eta0_0_box)
})
	})
	return cache_succ__2914940949
}

var cache_succ__3199041328 gopurs_runtime.Value
var once_succ__3199041328 sync.Once
func Get_succ__3199041328() gopurs_runtime.Value {
	once_succ__3199041328.Do(func() {
		cache_succ__3199041328 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_succ__3199041328(gopurs_runtime.CoerceToStruct[pkg_Data_Enum.Constructor_Enum[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_succ__3199041328
}

var cache_succ__2010692236 gopurs_runtime.Value
var once_succ__2010692236 sync.Once
func Get_succ__2010692236() gopurs_runtime.Value {
	once_succ__2010692236.Do(func() {
		cache_succ__2010692236 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_succ__2010692236(__eta0_0_box)
})
	})
	return cache_succ__2010692236
}

var cache_succ__2858180024 gopurs_runtime.Value
var once_succ__2858180024 sync.Once
func Get_succ__2858180024() gopurs_runtime.Value {
	once_succ__2858180024.Do(func() {
		cache_succ__2858180024 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_succ__2858180024(gopurs_runtime.CoerceToStruct[pkg_Data_Date.Constructor_Date](v_0_box)))}
})
	})
	return cache_succ__2858180024
}

var cache_toEnum__2203070892 gopurs_runtime.Value
var once_toEnum__2203070892 sync.Once
func Get_toEnum__2203070892() gopurs_runtime.Value {
	once_toEnum__2203070892.Do(func() {
		cache_toEnum__2203070892 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toEnum__2203070892(gopurs_runtime.CoerceToStruct[pkg_Data_Enum.Constructor_BoundedEnum[*pkg_Data_Date.Constructor_Date]](dict_0_box))
})
	})
	return cache_toEnum__2203070892
}

var cache_toEnum__2099864294 gopurs_runtime.Value
var once_toEnum__2099864294 sync.Once
func Get_toEnum__2099864294() gopurs_runtime.Value {
	once_toEnum__2099864294.Do(func() {
		cache_toEnum__2099864294 = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_toEnum__2099864294(n_0_box.IntVal))}
})
	})
	return cache_toEnum__2099864294
}

var cache_toEnum__3317293286 gopurs_runtime.Value
var once_toEnum__3317293286 sync.Once
func Get_toEnum__3317293286() gopurs_runtime.Value {
	once_toEnum__3317293286.Do(func() {
		cache_toEnum__3317293286 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toEnum__3317293286(gopurs_runtime.CoerceToStruct[pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_toEnum__3317293286
}

var cache_toEnum__2309750950 gopurs_runtime.Value
var once_toEnum__2309750950 sync.Once
func Get_toEnum__2309750950() gopurs_runtime.Value {
	once_toEnum__2309750950.Do(func() {
		cache_toEnum__2309750950 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_toEnum__2309750950(v_0_box.IntVal))}
})
	})
	return cache_toEnum__2309750950
}

var cache_eq__3259097883 gopurs_runtime.Value
var once_eq__3259097883 sync.Once
func Get_eq__3259097883() gopurs_runtime.Value {
	once_eq__3259097883.Do(func() {
		cache_eq__3259097883 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq__3259097883(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[*pkg_Data_Date.Constructor_Date]](dict_0_box))
})
	})
	return cache_eq__3259097883
}

var cache_eq__2843686287 gopurs_runtime.Value
var once_eq__2843686287 sync.Once
func Get_eq__2843686287() gopurs_runtime.Value {
	once_eq__2843686287.Do(func() {
		cache_eq__2843686287 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq__2843686287(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_eq__2843686287
}

var cache_eq__2384498378 gopurs_runtime.Value
var once_eq__2384498378 sync.Once
func Get_eq__2384498378() gopurs_runtime.Value {
	once_eq__2384498378.Do(func() {
		cache_eq__2384498378 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq__2384498378(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_eq__2384498378
}

var cache_eq__3887832182 gopurs_runtime.Value
var once_eq__3887832182 sync.Once
func Get_eq__3887832182() gopurs_runtime.Value {
	once_eq__3887832182.Do(func() {
		cache_eq__3887832182 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_eq__3887832182(uint32(x_0_box.IntVal), uint32(y_1_box.IntVal)))
})
	})
	return cache_eq__3887832182
}

var cache_eq__1204755874 gopurs_runtime.Value
var once_eq__1204755874 sync.Once
func Get_eq__1204755874() gopurs_runtime.Value {
	once_eq__1204755874.Do(func() {
		cache_eq__1204755874 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_eq__1204755874(gopurs_runtime.CoerceToStruct[pkg_Data_Date.Constructor_Date](x_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Date.Constructor_Date](y_1_box)))
})
	})
	return cache_eq__1204755874
}

var cache_mod__2185172824 gopurs_runtime.Value
var once_mod__2185172824 sync.Once
func Get_mod__2185172824() gopurs_runtime.Value {
	once_mod__2185172824.Do(func() {
		cache_mod__2185172824 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mod__2185172824(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_mod__2185172824
}

var cache_mod__2579358968 gopurs_runtime.Value
var once_mod__2579358968 sync.Once
func Get_mod__2579358968() gopurs_runtime.Value {
	once_mod__2579358968.Do(func() {
		cache_mod__2579358968 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mod__2579358968(gopurs_runtime.CoerceToStruct[pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_mod__2579358968
}

var cache_const__220790420 gopurs_runtime.Value
var once_const__220790420 sync.Once
func Get_const__220790420() gopurs_runtime.Value {
	once_const__220790420.Do(func() {
		cache_const__220790420 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__220790420(a_0_box, v_1_box)
})
	})
	return cache_const__220790420
}

var cache_const__641934996 gopurs_runtime.Value
var once_const__641934996 sync.Once
func Get_const__641934996() gopurs_runtime.Value {
	once_const__641934996.Do(func() {
		cache_const__641934996 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__641934996(a_0_box, v_1_box)
})
	})
	return cache_const__641934996
}

var cache_flip__1826582752 gopurs_runtime.Value
var once_flip__1826582752 sync.Once
func Get_flip__1826582752() gopurs_runtime.Value {
	once_flip__1826582752.Do(func() {
		cache_flip__1826582752 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__1826582752(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__1826582752
}

var cache_flip__3525081280 gopurs_runtime.Value
var once_flip__3525081280 sync.Once
func Get_flip__3525081280() gopurs_runtime.Value {
	once_flip__3525081280.Do(func() {
		cache_flip__3525081280 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__3525081280(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__3525081280
}

var cache_flip__2253242624 gopurs_runtime.Value
var once_flip__2253242624 sync.Once
func Get_flip__2253242624() gopurs_runtime.Value {
	once_flip__2253242624.Do(func() {
		cache_flip__2253242624 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__2253242624(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__2253242624
}

var cache_map__1165794789 gopurs_runtime.Value
var once_map__1165794789 sync.Once
func Get_map__1165794789() gopurs_runtime.Value {
	once_map__1165794789.Do(func() {
		cache_map__1165794789 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__1165794789(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_map__1165794789
}

var cache_map__1171202917 gopurs_runtime.Value
var once_map__1171202917 sync.Once
func Get_map__1171202917() gopurs_runtime.Value {
	once_map__1171202917.Do(func() {
		cache_map__1171202917 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__1171202917(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_map__1171202917
}

var cache_map__1307706501 gopurs_runtime.Value
var once_map__1307706501 sync.Once
func Get_map__1307706501() gopurs_runtime.Value {
	once_map__1307706501.Do(func() {
		cache_map__1307706501 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__1307706501(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__1307706501
}

var cache_map__1408505925 gopurs_runtime.Value
var once_map__1408505925 sync.Once
func Get_map__1408505925() gopurs_runtime.Value {
	once_map__1408505925.Do(func() {
		cache_map__1408505925 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__1408505925(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_map__1408505925
}

var cache_map__243231988 gopurs_runtime.Value
var once_map__243231988 sync.Once
func Get_map__243231988() gopurs_runtime.Value {
	once_map__243231988.Do(func() {
		cache_map__243231988 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__243231988(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__243231988
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

var cache_map__3447677596 gopurs_runtime.Value
var once_map__3447677596 sync.Once
func Get_map__3447677596() gopurs_runtime.Value {
	once_map__3447677596.Do(func() {
		cache_map__3447677596 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_map__3447677596(v_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](v1_1_box)))}
})
	})
	return cache_map__3447677596
}

var cache_map__4155962236 gopurs_runtime.Value
var once_map__4155962236 sync.Once
func Get_map__4155962236() gopurs_runtime.Value {
	once_map__4155962236.Do(func() {
		cache_map__4155962236 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_map__4155962236(v_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](v1_1_box)))}
})
	})
	return cache_map__4155962236
}

var cache_map__901270812 gopurs_runtime.Value
var once_map__901270812 sync.Once
func Get_map__901270812() gopurs_runtime.Value {
	once_map__901270812.Do(func() {
		cache_map__901270812 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_map__901270812(v_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v1_1_box)))}
})
	})
	return cache_map__901270812
}

var cache_conj__2927892844 gopurs_runtime.Value
var once_conj__2927892844 sync.Once
func Get_conj__2927892844() gopurs_runtime.Value {
	once_conj__2927892844.Do(func() {
		cache_conj__2927892844 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_conj__2927892844(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[int64]](dict_0_box))
})
	})
	return cache_conj__2927892844
}

var cache_conj__4093645121 gopurs_runtime.Value
var once_conj__4093645121 sync.Once
func Get_conj__4093645121() gopurs_runtime.Value {
	once_conj__4093645121.Do(func() {
		cache_conj__4093645121 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_conj__4093645121(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[*pkg_Data_Date.Constructor_Date]](dict_0_box))
})
	})
	return cache_conj__4093645121
}

var cache_conj__3676519832 gopurs_runtime.Value
var once_conj__3676519832 sync.Once
func Get_conj__3676519832() gopurs_runtime.Value {
	once_conj__3676519832.Do(func() {
		cache_conj__3676519832 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_conj__3676519832(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_conj__3676519832
}

var cache_conj__3472268504 gopurs_runtime.Value
var once_conj__3472268504 sync.Once
func Get_conj__3472268504() gopurs_runtime.Value {
	once_conj__3472268504.Do(func() {
		cache_conj__3472268504 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_conj__3472268504(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_conj__3472268504
}

var cache_disj__3676519832 gopurs_runtime.Value
var once_disj__3676519832 sync.Once
func Get_disj__3676519832() gopurs_runtime.Value {
	once_disj__3676519832.Do(func() {
		cache_disj__3676519832 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_disj__3676519832(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_disj__3676519832
}

var cache_disj__3472268504 gopurs_runtime.Value
var once_disj__3472268504 sync.Once
func Get_disj__3472268504() gopurs_runtime.Value {
	once_disj__3472268504.Do(func() {
		cache_disj__3472268504 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_disj__3472268504(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_disj__3472268504
}

var cache_not__3201284355 gopurs_runtime.Value
var once_not__3201284355 sync.Once
func Get_not__3201284355() gopurs_runtime.Value {
	once_not__3201284355.Do(func() {
		cache_not__3201284355 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_not__3201284355(__eta0_0_box)
})
	})
	return cache_not__3201284355
}

var cache_not__1505204753 gopurs_runtime.Value
var once_not__1505204753 sync.Once
func Get_not__1505204753() gopurs_runtime.Value {
	once_not__1505204753.Do(func() {
		cache_not__1505204753 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_not__1505204753(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_not__1505204753
}

var cache_applicativeMaybe__3016118221 gopurs_runtime.Value
var once_applicativeMaybe__3016118221 sync.Once
func Get_applicativeMaybe__3016118221() gopurs_runtime.Value {
	once_applicativeMaybe__3016118221.Do(func() {
		cache_applicativeMaybe__3016118221 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_applyMaybe()
}), pkg_Data_Maybe.Get_Just())
	})
	return cache_applicativeMaybe__3016118221
}

var cache_applyMaybe__3561700045 gopurs_runtime.Value
var once_applyMaybe__3561700045 sync.Once
func Get_applyMaybe__3561700045() gopurs_runtime.Value {
	once_applyMaybe__3561700045.Do(func() {
		cache_applyMaybe__3561700045 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_functorMaybe()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_0.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v1_1))})))}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_applyMaybe__3561700045
}

var cache_applyMaybe__3698865467 gopurs_runtime.Value
var once_applyMaybe__3698865467 sync.Once
func Get_applyMaybe__3698865467() gopurs_runtime.Value {
	once_applyMaybe__3698865467.Do(func() {
		cache_applyMaybe__3698865467 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_functorMaybe()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_0.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v1_1))})))}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_applyMaybe__3698865467
}

var cache_bindMaybe__1910292045 gopurs_runtime.Value
var once_bindMaybe__1910292045 sync.Once
func Get_bindMaybe__1910292045() gopurs_runtime.Value {
	once_bindMaybe__1910292045.Do(func() {
		cache_bindMaybe__1910292045 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_applyMaybe()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(v1_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_0.UnsafePtr).V0)))}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_bindMaybe__1910292045
}

var cache_fromJust__1577979644 gopurs_runtime.Value
var once_fromJust__1577979644 sync.Once
func Get_fromJust__1577979644() gopurs_runtime.Value {
	once_fromJust__1577979644.Do(func() {
		cache_fromJust__1577979644 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromJust__1577979644(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_fromJust__1577979644
}

var cache_fromJust__1791383420 gopurs_runtime.Value
var once_fromJust__1791383420 sync.Once
func Get_fromJust__1791383420() gopurs_runtime.Value {
	once_fromJust__1791383420.Do(func() {
		cache_fromJust__1791383420 = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromJust__1791383420(_dollar__unused_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v_1_box))
})
	})
	return cache_fromJust__1791383420
}

var cache_fromJust__4142563260 gopurs_runtime.Value
var once_fromJust__4142563260 sync.Once
func Get_fromJust__4142563260() gopurs_runtime.Value {
	once_fromJust__4142563260.Do(func() {
		cache_fromJust__4142563260 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromJust__4142563260(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_fromJust__4142563260
}

var cache_fromJust__965748316 gopurs_runtime.Value
var once_fromJust__965748316 sync.Once
func Get_fromJust__965748316() gopurs_runtime.Value {
	once_fromJust__965748316.Do(func() {
		cache_fromJust__965748316 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromJust__965748316(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_fromJust__965748316
}

var cache_fromMaybe__430429096 gopurs_runtime.Value
var once_fromMaybe__430429096 sync.Once
func Get_fromMaybe__430429096() gopurs_runtime.Value {
	once_fromMaybe__430429096.Do(func() {
		cache_fromMaybe__430429096 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromMaybe__430429096(a_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_1_box))
})
	})
	return cache_fromMaybe__430429096
}

var cache_fromMaybe__18840980 gopurs_runtime.Value
var once_fromMaybe__18840980 sync.Once
func Get_fromMaybe__18840980() gopurs_runtime.Value {
	once_fromMaybe__18840980.Do(func() {
		cache_fromMaybe__18840980 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromMaybe__18840980(a_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_1_box))
})
	})
	return cache_fromMaybe__18840980
}

var cache_fromMaybe__737056608 gopurs_runtime.Value
var once_fromMaybe__737056608 sync.Once
func Get_fromMaybe__737056608() gopurs_runtime.Value {
	once_fromMaybe__737056608.Do(func() {
		cache_fromMaybe__737056608 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromMaybe__737056608(a_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_1_box))
})
	})
	return cache_fromMaybe__737056608
}

var cache_functorMaybe__2569569018 gopurs_runtime.Value
var once_functorMaybe__2569569018 sync.Once
func Get_functorMaybe__2569569018() gopurs_runtime.Value {
	once_functorMaybe__2569569018.Do(func() {
		cache_functorMaybe__2569569018 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 930809136 && v1_1.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_1.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_functorMaybe__2569569018
}

var cache_functorMaybe__2097654001 gopurs_runtime.Value
var once_functorMaybe__2097654001 sync.Once
func Get_functorMaybe__2097654001() gopurs_runtime.Value {
	once_functorMaybe__2097654001.Do(func() {
		cache_functorMaybe__2097654001 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 930809136 && v1_1.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_1.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_functorMaybe__2097654001
}

var cache_isNothing__1358705270 gopurs_runtime.Value
var once_isNothing__1358705270 sync.Once
func Get_isNothing__1358705270() gopurs_runtime.Value {
	once_isNothing__1358705270.Do(func() {
		cache_isNothing__1358705270 = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_isNothing__1358705270(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_0_box)))
})
	})
	return cache_isNothing__1358705270
}

var cache_isNothing__4206805139 gopurs_runtime.Value
var once_isNothing__4206805139 sync.Once
func Get_isNothing__4206805139() gopurs_runtime.Value {
	once_isNothing__4206805139.Do(func() {
		cache_isNothing__4206805139 = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_isNothing__4206805139(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_0_box)))
})
	})
	return cache_isNothing__4206805139
}

var cache_isNothing__2787066607 gopurs_runtime.Value
var once_isNothing__2787066607 sync.Once
func Get_isNothing__2787066607() gopurs_runtime.Value {
	once_isNothing__2787066607.Do(func() {
		cache_isNothing__2787066607 = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_isNothing__2787066607(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_0_box)))
})
	})
	return cache_isNothing__2787066607
}

var cache_isNothing__323776123 gopurs_runtime.Value
var once_isNothing__323776123 sync.Once
func Get_isNothing__323776123() gopurs_runtime.Value {
	once_isNothing__323776123.Do(func() {
		cache_isNothing__323776123 = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_isNothing__323776123(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_0_box)))
})
	})
	return cache_isNothing__323776123
}

var cache_maybe__1594528518 gopurs_runtime.Value
var once_maybe__1594528518 sync.Once
func Get_maybe__1594528518() gopurs_runtime.Value {
	once_maybe__1594528518.Do(func() {
		cache_maybe__1594528518 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe__1594528518(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_maybe__1594528518
}

var cache_maybe__3658316244 gopurs_runtime.Value
var once_maybe__3658316244 sync.Once
func Get_maybe__3658316244() gopurs_runtime.Value {
	once_maybe__3658316244.Do(func() {
		cache_maybe__3658316244 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe__3658316244(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_maybe__3658316244
}

var cache_compare__669572705 gopurs_runtime.Value
var once_compare__669572705 sync.Once
func Get_compare__669572705() gopurs_runtime.Value {
	once_compare__669572705.Do(func() {
		cache_compare__669572705 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compare__669572705(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compare__669572705
}

var cache_compare__372254389 gopurs_runtime.Value
var once_compare__372254389 sync.Once
func Get_compare__372254389() gopurs_runtime.Value {
	once_compare__372254389.Do(func() {
		cache_compare__372254389 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compare__372254389(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_compare__372254389
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

var cache_compare__696857420 gopurs_runtime.Value
var once_compare__696857420 sync.Once
func Get_compare__696857420() gopurs_runtime.Value {
	once_compare__696857420.Do(func() {
		cache_compare__696857420 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_compare__696857420(uint32(x_0_box.IntVal), uint32(y_1_box.IntVal))), UnsafePtr: nil}
})
	})
	return cache_compare__696857420
}

var cache_greaterThan__3259097883 gopurs_runtime.Value
var once_greaterThan__3259097883 sync.Once
func Get_greaterThan__3259097883() gopurs_runtime.Value {
	once_greaterThan__3259097883.Do(func() {
		cache_greaterThan__3259097883 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThan__3259097883(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_greaterThan__3259097883
}

var cache_greaterThan__4087042607 gopurs_runtime.Value
var once_greaterThan__4087042607 sync.Once
func Get_greaterThan__4087042607() gopurs_runtime.Value {
	once_greaterThan__4087042607.Do(func() {
		cache_greaterThan__4087042607 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThan__4087042607(a1_0_box, a2_1_box))
})
	})
	return cache_greaterThan__4087042607
}

var cache_greaterThan__1409282474 gopurs_runtime.Value
var once_greaterThan__1409282474 sync.Once
func Get_greaterThan__1409282474() gopurs_runtime.Value {
	once_greaterThan__1409282474.Do(func() {
		cache_greaterThan__1409282474 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThan__1409282474(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_greaterThan__1409282474
}

var cache_greaterThan__2400628110 gopurs_runtime.Value
var once_greaterThan__2400628110 sync.Once
func Get_greaterThan__2400628110() gopurs_runtime.Value {
	once_greaterThan__2400628110.Do(func() {
		cache_greaterThan__2400628110 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThan__2400628110(a1_0_box, a2_1_box))
})
	})
	return cache_greaterThan__2400628110
}

var cache_greaterThanOrEq__1710332219 gopurs_runtime.Value
var once_greaterThanOrEq__1710332219 sync.Once
func Get_greaterThanOrEq__1710332219() gopurs_runtime.Value {
	once_greaterThanOrEq__1710332219.Do(func() {
		cache_greaterThanOrEq__1710332219 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThanOrEq__1710332219(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_greaterThanOrEq__1710332219
}

var cache_greaterThanOrEq__4087042607 gopurs_runtime.Value
var once_greaterThanOrEq__4087042607 sync.Once
func Get_greaterThanOrEq__4087042607() gopurs_runtime.Value {
	once_greaterThanOrEq__4087042607.Do(func() {
		cache_greaterThanOrEq__4087042607 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThanOrEq__4087042607(a1_0_box, a2_1_box))
})
	})
	return cache_greaterThanOrEq__4087042607
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

var cache_lessThan__4087042607 gopurs_runtime.Value
var once_lessThan__4087042607 sync.Once
func Get_lessThan__4087042607() gopurs_runtime.Value {
	once_lessThan__4087042607.Do(func() {
		cache_lessThan__4087042607 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThan__4087042607(a1_0_box, a2_1_box))
})
	})
	return cache_lessThan__4087042607
}

var cache_lessThan__1409282474 gopurs_runtime.Value
var once_lessThan__1409282474 sync.Once
func Get_lessThan__1409282474() gopurs_runtime.Value {
	once_lessThan__1409282474.Do(func() {
		cache_lessThan__1409282474 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThan__1409282474(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_lessThan__1409282474
}

var cache_lessThanOrEq__1710332219 gopurs_runtime.Value
var once_lessThanOrEq__1710332219 sync.Once
func Get_lessThanOrEq__1710332219() gopurs_runtime.Value {
	once_lessThanOrEq__1710332219.Do(func() {
		cache_lessThanOrEq__1710332219 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThanOrEq__1710332219(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_lessThanOrEq__1710332219
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

var cache_negate__2635823316 gopurs_runtime.Value
var once_negate__2635823316 sync.Once
func Get_negate__2635823316() gopurs_runtime.Value {
	once_negate__2635823316.Do(func() {
		cache_negate__2635823316 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_negate__2635823316(__eta0_0_box)
})
	})
	return cache_negate__2635823316
}

var cache_negate__1364373265 gopurs_runtime.Value
var once_negate__1364373265 sync.Once
func Get_negate__1364373265() gopurs_runtime.Value {
	once_negate__1364373265.Do(func() {
		cache_negate__1364373265 = gopurs_runtime.Func(func(dictRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_negate__1364373265(gopurs_runtime.CoerceToStruct[pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]](dictRing_0_box))
})
	})
	return cache_negate__1364373265
}

var cache_sub__2927892844 gopurs_runtime.Value
var once_sub__2927892844 sync.Once
func Get_sub__2927892844() gopurs_runtime.Value {
	once_sub__2927892844.Do(func() {
		cache_sub__2927892844 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__2927892844(gopurs_runtime.CoerceToStruct[pkg_Data_Ring.Constructor_Ring[int64]](dict_0_box))
})
	})
	return cache_sub__2927892844
}

var cache_sub__1124926121 gopurs_runtime.Value
var once_sub__1124926121 sync.Once
func Get_sub__1124926121() gopurs_runtime.Value {
	once_sub__1124926121.Do(func() {
		cache_sub__1124926121 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__1124926121(gopurs_runtime.CoerceToStruct[pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sub__1124926121
}

var cache_sub__1841809173 gopurs_runtime.Value
var once_sub__1841809173 sync.Once
func Get_sub__1841809173() gopurs_runtime.Value {
	once_sub__1841809173.Do(func() {
		cache_sub__1841809173 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__1841809173(gopurs_runtime.CoerceToStruct[pkg_Data_Ring.Constructor_Ring[uint32]](dict_0_box))
})
	})
	return cache_sub__1841809173
}

var cache_sub__1043827704 gopurs_runtime.Value
var once_sub__1043827704 sync.Once
func Get_sub__1043827704() gopurs_runtime.Value {
	once_sub__1043827704.Do(func() {
		cache_sub__1043827704 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__1043827704(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_sub__1043827704
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

var cache_add__2927892844 gopurs_runtime.Value
var once_add__2927892844 sync.Once
func Get_add__2927892844() gopurs_runtime.Value {
	once_add__2927892844.Do(func() {
		cache_add__2927892844 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_add__2927892844(gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[int64]](dict_0_box))
})
	})
	return cache_add__2927892844
}

var cache_add__1841809173 gopurs_runtime.Value
var once_add__1841809173 sync.Once
func Get_add__1841809173() gopurs_runtime.Value {
	once_add__1841809173.Do(func() {
		cache_add__1841809173 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_add__1841809173(gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[uint32]](dict_0_box))
})
	})
	return cache_add__1841809173
}

var cache_add__560788792 gopurs_runtime.Value
var once_add__560788792 sync.Once
func Get_add__560788792() gopurs_runtime.Value {
	once_add__560788792.Do(func() {
		cache_add__560788792 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_add__560788792(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_add__560788792
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

var cache_zero__1556010056 gopurs_runtime.Value
var once_zero__1556010056 sync.Once
func Get_zero__1556010056() gopurs_runtime.Value {
	once_zero__1556010056.Do(func() {
		cache_zero__1556010056 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zero__1556010056(dict_0_box)
})
	})
	return cache_zero__1556010056
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

var cache_unsafePartial__3861213094 gopurs_runtime.Value
var once_unsafePartial__3861213094 sync.Once
func Get_unsafePartial__3861213094() gopurs_runtime.Value {
	once_unsafePartial__3861213094.Do(func() {
		cache_unsafePartial__3861213094 = pkg_Partial_Unsafe.Get__unsafePartial()
	})
	return cache_unsafePartial__3861213094
}

var cache_unsafePartial__1306634845 gopurs_runtime.Value
var once_unsafePartial__1306634845 sync.Once
func Get_unsafePartial__1306634845() gopurs_runtime.Value {
	once_unsafePartial__1306634845.Do(func() {
		cache_unsafePartial__1306634845 = pkg_Partial_Unsafe.Get__unsafePartial()
	})
	return cache_unsafePartial__1306634845
}

var cache_unsafePartial__1130268957 gopurs_runtime.Value
var once_unsafePartial__1130268957 sync.Once
func Get_unsafePartial__1130268957() gopurs_runtime.Value {
	once_unsafePartial__1130268957.Do(func() {
		cache_unsafePartial__1130268957 = pkg_Partial_Unsafe.Get__unsafePartial()
	})
	return cache_unsafePartial__1130268957
}

var cache_unsafePartial__1059034269 gopurs_runtime.Value
var once_unsafePartial__1059034269 sync.Once
func Get_unsafePartial__1059034269() gopurs_runtime.Value {
	once_unsafePartial__1059034269.Do(func() {
		cache_unsafePartial__1059034269 = pkg_Partial_Unsafe.Get__unsafePartial()
	})
	return cache_unsafePartial__1059034269
}

func Call_genDate(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
Bind1_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
Functor0_3_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_2
pure_4_3 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_4_3
return gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply(pkg_Data_Date_Component_Gen.Get_genYear(), dictMonadGen_0), gopurs_runtime.Func(func(year_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 int64
{
if (gopurs_runtime.Apply(pkg_Data_Date.Get_isLeapYear(), gopurs_runtime.Int(year_5.IntVal)).IntVal) != (0) {
__t4 = 365
goto end_branch_4
} else {

}
}
{
__t4 = 364
}
end_branch_4:
return gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply2(Functor0_3_2.V0, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Int.Get_toNumber(), x_6)
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(0), gopurs_runtime.Int(__t4))), gopurs_runtime.Func(func(days_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_4_3, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Date.Constructor_Date](gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Date.Constructor_Date](Call_fromJust__965748316(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Date.Constructor_Date]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_bind__3444930753(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Date.Constructor_Date]](gopurs_runtime.Apply3(pkg_Data_Date.Get_exactDate(), gopurs_runtime.Int(year_5.IntVal), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedMonth(), "bottom").IntVal)), UnsafePtr: nil}, gopurs_runtime.Int(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedDay(), "bottom").IntVal))), gopurs_runtime.Func(func(janFirst_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Date.Constructor_Date]](gopurs_runtime.Apply2(pkg_Data_Date.Get_adjust(), gopurs_runtime.Float(days_6.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Date.Constructor_Date](janFirst_8))})))}
})))}))}))))}
}))))})
}))
}))
}

func Call_pure__779048344(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__4037597590(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__3215807376(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__3181299446(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, __eta0_0})}
}

func Call_apply__2174094052(dict_0_loop *pkg_Control_Apply.Constructor_Apply[*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__652936900(dict_0_loop *pkg_Control_Apply.Constructor_Apply[*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__1982519530(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__353515660(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__1471729482(v_0_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value], v1_1_loop *pkg_Data_Maybe.Constructor_Just[int64]) *pkg_Data_Maybe.Constructor_Just[*pkg_Data_Date.Constructor_Date] {
var v_0 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v_0_loop
_ = v_0
var v1_1 *pkg_Data_Maybe.Constructor_Just[int64] = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Date.Constructor_Date]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)})))}
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Date.Constructor_Date]](__t0)
}

func Call_apply__1572009162(v_0_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value], v1_1_loop *pkg_Data_Maybe.Constructor_Just[uint32]) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var v_0 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v_0_loop
_ = v_0
var v1_1 *pkg_Data_Maybe.Constructor_Just[uint32] = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)})))}
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0)
}

func Call_bind__3254602343(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__4011257415(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__2601835655(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__4062037089(v_0_loop *pkg_Data_Maybe.Constructor_Just[int64], v1_1_loop gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[*pkg_Data_Date.Constructor_Date] {
var v_0 *pkg_Data_Maybe.Constructor_Just[int64] = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Date.Constructor_Date]](gopurs_runtime.Apply(v1_1, gopurs_runtime.Int((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.IntVal))))}
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Date.Constructor_Date]](__t0)
}

func Call_bind__3444930753(v_0_loop *pkg_Data_Maybe.Constructor_Just[*pkg_Data_Date.Constructor_Date], v1_1_loop gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[*pkg_Data_Date.Constructor_Date] {
var v_0 *pkg_Data_Maybe.Constructor_Just[*pkg_Data_Date.Constructor_Date] = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Date.Constructor_Date]](gopurs_runtime.Apply(v1_1, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Date.Constructor_Date]((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0))})))}
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Date.Constructor_Date]](__t0)
}

func Call_bindFlipped__1485397639(dictBind_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value], b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dictBind_0_loop
_ = dictBind_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(dictBind_0.V1, a_2, b_1)
}

func Call_bindFlipped__3917280577(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_bindMaybe(), "bind"), __eta1_1, __eta0_0)
}

func Call_bindFlipped__1454086721(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_bindMaybe(), "bind"), __eta1_1, __eta0_0)
}

func Call_chooseInt__1063828903(dict_0_loop *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_compose__1636404804(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[uint32]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[uint32] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_compose__1774599291(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[int64]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[int64] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_compose__1543665403(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[int64]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[int64] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_fromEnum__3599151655(v_0_loop int64) int64 {
var v_0 int64 = v_0_loop
_ = v_0
return v_0
}

func Call_fromEnum__1637084359(dict_0_loop *pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_fromEnum__1196942535(v_0_loop uint32) int64 {
var v_0 uint32 = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0 == 1908470532) {
__t0 = gopurs_runtime.Int(1)
goto end_branch_0
} else {

}
}
{
if (v_0 == 2455627378) {
__t0 = gopurs_runtime.Int(2)
goto end_branch_0
} else {

}
}
{
if (v_0 == 4162469099) {
__t0 = gopurs_runtime.Int(3)
goto end_branch_0
} else {

}
}
{
if (v_0 == 1692989816) {
__t0 = gopurs_runtime.Int(4)
goto end_branch_0
} else {

}
}
{
if (v_0 == 330658827) {
__t0 = gopurs_runtime.Int(5)
goto end_branch_0
} else {

}
}
{
if (v_0 == 4067355978) {
__t0 = gopurs_runtime.Int(6)
goto end_branch_0
} else {

}
}
{
if (v_0 == 2276710548) {
__t0 = gopurs_runtime.Int(7)
goto end_branch_0
} else {

}
}
{
if (v_0 == 243771071) {
__t0 = gopurs_runtime.Int(8)
goto end_branch_0
} else {

}
}
{
if (v_0 == 215731793) {
__t0 = gopurs_runtime.Int(9)
goto end_branch_0
} else {

}
}
{
if (v_0 == 8639228) {
__t0 = gopurs_runtime.Int(10)
goto end_branch_0
} else {

}
}
{
if (v_0 == 49471444) {
__t0 = gopurs_runtime.Int(11)
goto end_branch_0
} else {

}
}
{
if (v_0 == 3889233761) {
__t0 = gopurs_runtime.Int(12)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0.IntVal
}

func Call_pred__2914940949(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum"), gopurs_runtime.Int((gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "fromEnum"), __eta0_0).IntVal) - (1)))))}
}

func Call_pred__3199041328(dict_0_loop *pkg_Data_Enum.Constructor_Enum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Enum.Constructor_Enum[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pred__2010692236(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[uint32]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumMonth(), "toEnum"), gopurs_runtime.Int((gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumMonth(), "fromEnum"), __eta0_0).IntVal) - (1)))))}
}

func Call_succ__412946465(dict_0_loop *pkg_Data_Enum.Constructor_Enum[*pkg_Data_Date.Constructor_Date]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Enum.Constructor_Enum[*pkg_Data_Date.Constructor_Date] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_succ__2914940949(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum"), gopurs_runtime.Int((gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "fromEnum"), __eta0_0).IntVal) + (1)))))}
}

func Call_succ__3199041328(dict_0_loop *pkg_Data_Enum.Constructor_Enum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Enum.Constructor_Enum[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_succ__2010692236(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[uint32]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumMonth(), "toEnum"), gopurs_runtime.Int((gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumMonth(), "fromEnum"), __eta0_0).IntVal) + (1)))))}
}

func Call_succ__2858180024(v_0_loop *pkg_Data_Date.Constructor_Date) *pkg_Data_Maybe.Constructor_Just[*pkg_Data_Date.Constructor_Date] {
var v_0 *pkg_Data_Date.Constructor_Date = v_0_loop
_ = v_0
sm_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[uint32]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_enumMonth(), "succ"), gopurs_runtime.Value{Type: 9, IntVal: int64((*pkg_Data_Date.Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1), UnsafePtr: nil}))
_ = sm_1_0
v1_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_enumDay(), "succ"), gopurs_runtime.Int((*pkg_Data_Date.Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V2)))
_ = v1_2_1
var __t5 gopurs_runtime.Value
{
var __t4 bool
{
var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[*pkg_Data_Maybe.Constructor_Just[int64]]](pkg_Data_Date.Get_ordMaybe()).V1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_2_1)}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(gopurs_runtime.Apply2(pkg_Data_Date.Get_lastDayOfMonth(), gopurs_runtime.Int((*pkg_Data_Date.Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: int64((*pkg_Data_Date.Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1), UnsafePtr: nil}).IntVal)})})
if (uint32(__t_tag_3.IntVal) == 380165415) {
__t4 = true
goto end_branch_4
} else {

}
}
{
__t4 = false
}
end_branch_4:
if __t4 {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_2_1)}
}
end_branch_5:
sd_3_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](__t5)
_ = sd_3_2
var __t8 gopurs_runtime.Value
{
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sd_3_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sd_3_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sd_3_2)}.UnsafePtr == nil) {
__t6 = gopurs_runtime.Bool(true)
goto end_branch_6
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sd_3_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sd_3_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sd_3_2)}.UnsafePtr != nil) {
__t6 = gopurs_runtime.Bool(false)
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
var __t7 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sm_1_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sm_1_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sm_1_0)}.UnsafePtr == nil) {
__t7 = gopurs_runtime.Bool(true)
goto end_branch_7
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sm_1_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sm_1_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sm_1_0)}.UnsafePtr != nil) {
__t7 = gopurs_runtime.Bool(false)
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Bool((__t6.IntVal) != (0)), gopurs_runtime.Bool((__t7.IntVal) != (0))).IntVal) != (0) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_enumYear(), "succ"), gopurs_runtime.Int((*pkg_Data_Date.Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0))))}
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int((*pkg_Data_Date.Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0)})}
}
end_branch_8:
var __t11 uint32
{
var __t9 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sd_3_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sd_3_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sd_3_2)}.UnsafePtr == nil) {
__t9 = gopurs_runtime.Bool(true)
goto end_branch_9
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sd_3_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sd_3_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sd_3_2)}.UnsafePtr != nil) {
__t9 = gopurs_runtime.Bool(false)
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
if (__t9.IntVal) != (0) {
var __t10 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sm_1_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sm_1_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sm_1_0)}.UnsafePtr == nil) {
__t10 = gopurs_runtime.Value{Type: 9, IntVal: int64(1908470532), UnsafePtr: nil}
goto end_branch_10
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sm_1_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sm_1_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sm_1_0)}.UnsafePtr != nil) {
__t10 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sm_1_0)}.UnsafePtr).V0
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
__t11 = uint32(__t10.IntVal)
goto end_branch_11
} else {

}
}
{
__t11 = (*pkg_Data_Date.Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1
}
end_branch_11:
var __t13 *pkg_Data_Maybe.Constructor_Just[int64]
{
var __t12 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sd_3_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sd_3_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sd_3_2)}.UnsafePtr == nil) {
__t12 = gopurs_runtime.Bool(true)
goto end_branch_12
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sd_3_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sd_3_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sd_3_2)}.UnsafePtr != nil) {
__t12 = gopurs_runtime.Bool(false)
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_12:
if (__t12.IntVal) != (0) {
__t13 = gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(1)))
goto end_branch_13
} else {

}
}
{
__t13 = sd_3_2
}
end_branch_13:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Date.Constructor_Date]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applyMaybe(), "apply"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applyMaybe(), "apply"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Date.Get_Date(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](__t8))})))}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[uint32]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applicativeMaybe(), "pure"), gopurs_runtime.Value{Type: 9, IntVal: int64(__t11), UnsafePtr: nil})))})))}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t13)}))
}

func Call_toEnum__2203070892(dict_0_loop *pkg_Data_Enum.Constructor_BoundedEnum[*pkg_Data_Date.Constructor_Date]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Enum.Constructor_BoundedEnum[*pkg_Data_Date.Constructor_Date] = dict_0_loop
_ = dict_0
return dict_0.V4
}

func Call_toEnum__2099864294(n_0_loop int64) *pkg_Data_Maybe.Constructor_Just[int64] {
var n_0 int64 = n_0_loop
_ = n_0
var __t2 gopurs_runtime.Value
{
var __t0 bool
{
if (gopurs_runtime.Int(n_0).IntVal) < (gopurs_runtime.Int(1).IntVal) {
__t0 = false
goto end_branch_0
} else {

}
}
{
__t0 = true
}
end_branch_0:
var __t1 bool
{
if (gopurs_runtime.Int(n_0).IntVal) > (gopurs_runtime.Int(31).IntVal) {
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = true
}
end_branch_1:
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Bool(__t0), gopurs_runtime.Bool(__t1)).IntVal) != (0) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(n_0)})}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](__t2)
}

func Call_toEnum__3317293286(dict_0_loop *pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V4
}

func Call_toEnum__2309750950(v_0_loop int64) *pkg_Data_Maybe.Constructor_Just[uint32] {
var v_0 int64 = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0) == (1) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(1908470532), UnsafePtr: nil}})}
goto end_branch_0
} else {

}
}
{
if (v_0) == (2) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(2455627378), UnsafePtr: nil}})}
goto end_branch_0
} else {

}
}
{
if (v_0) == (3) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(4162469099), UnsafePtr: nil}})}
goto end_branch_0
} else {

}
}
{
if (v_0) == (4) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(1692989816), UnsafePtr: nil}})}
goto end_branch_0
} else {

}
}
{
if (v_0) == (5) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(330658827), UnsafePtr: nil}})}
goto end_branch_0
} else {

}
}
{
if (v_0) == (6) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(4067355978), UnsafePtr: nil}})}
goto end_branch_0
} else {

}
}
{
if (v_0) == (7) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(2276710548), UnsafePtr: nil}})}
goto end_branch_0
} else {

}
}
{
if (v_0) == (8) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(243771071), UnsafePtr: nil}})}
goto end_branch_0
} else {

}
}
{
if (v_0) == (9) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(215731793), UnsafePtr: nil}})}
goto end_branch_0
} else {

}
}
{
if (v_0) == (10) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(8639228), UnsafePtr: nil}})}
goto end_branch_0
} else {

}
}
{
if (v_0) == (11) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(49471444), UnsafePtr: nil}})}
goto end_branch_0
} else {

}
}
{
if (v_0) == (12) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(3889233761), UnsafePtr: nil}})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[uint32]](__t0)
}

func Call_eq__3259097883(dict_0_loop *pkg_Data_Eq.Constructor_Eq[*pkg_Data_Date.Constructor_Date]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Eq.Constructor_Eq[*pkg_Data_Date.Constructor_Date] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_eq__2843686287(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Bool((__eta0_0.IntVal) == (__eta1_1.IntVal))
}

func Call_eq__2384498378(dict_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_eq__3887832182(x_0_loop uint32, y_1_loop uint32) bool {
var x_0 uint32 = x_0_loop
_ = x_0
var y_1 uint32 = y_1_loop
_ = y_1
var __t11 bool
{
if (x_0 == 1908470532) {
var __t0 bool
{
if (y_1 == 1908470532) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
__t11 = __t0
goto end_branch_11
} else {

}
}
{
if (x_0 == 2455627378) {
var __t1 bool
{
if (y_1 == 2455627378) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
__t11 = __t1
goto end_branch_11
} else {

}
}
{
if (x_0 == 4162469099) {
var __t2 bool
{
if (y_1 == 4162469099) {
__t2 = true
goto end_branch_2
} else {

}
}
{
__t2 = false
}
end_branch_2:
__t11 = __t2
goto end_branch_11
} else {

}
}
{
if (x_0 == 1692989816) {
var __t3 bool
{
if (y_1 == 1692989816) {
__t3 = true
goto end_branch_3
} else {

}
}
{
__t3 = false
}
end_branch_3:
__t11 = __t3
goto end_branch_11
} else {

}
}
{
if (x_0 == 330658827) {
var __t4 bool
{
if (y_1 == 330658827) {
__t4 = true
goto end_branch_4
} else {

}
}
{
__t4 = false
}
end_branch_4:
__t11 = __t4
goto end_branch_11
} else {

}
}
{
if (x_0 == 4067355978) {
var __t5 bool
{
if (y_1 == 4067355978) {
__t5 = true
goto end_branch_5
} else {

}
}
{
__t5 = false
}
end_branch_5:
__t11 = __t5
goto end_branch_11
} else {

}
}
{
if (x_0 == 2276710548) {
var __t6 bool
{
if (y_1 == 2276710548) {
__t6 = true
goto end_branch_6
} else {

}
}
{
__t6 = false
}
end_branch_6:
__t11 = __t6
goto end_branch_11
} else {

}
}
{
if (x_0 == 243771071) {
var __t7 bool
{
if (y_1 == 243771071) {
__t7 = true
goto end_branch_7
} else {

}
}
{
__t7 = false
}
end_branch_7:
__t11 = __t7
goto end_branch_11
} else {

}
}
{
if (x_0 == 215731793) {
var __t8 bool
{
if (y_1 == 215731793) {
__t8 = true
goto end_branch_8
} else {

}
}
{
__t8 = false
}
end_branch_8:
__t11 = __t8
goto end_branch_11
} else {

}
}
{
if (x_0 == 8639228) {
var __t9 bool
{
if (y_1 == 8639228) {
__t9 = true
goto end_branch_9
} else {

}
}
{
__t9 = false
}
end_branch_9:
__t11 = __t9
goto end_branch_11
} else {

}
}
{
if (x_0 == 49471444) {
var __t10 bool
{
if (y_1 == 49471444) {
__t10 = true
goto end_branch_10
} else {

}
}
{
__t10 = false
}
end_branch_10:
__t11 = __t10
goto end_branch_11
} else {

}
}
{
if ((x_0 == 3889233761)) && ((y_1 == 3889233761)) {
__t11 = true
goto end_branch_11
} else {

}
}
{
__t11 = false
}
end_branch_11:
return __t11
}

func Call_eq__1204755874(x_0_loop *pkg_Data_Date.Constructor_Date, y_1_loop *pkg_Data_Date.Constructor_Date) bool {
var x_0 *pkg_Data_Date.Constructor_Date = x_0_loop
_ = x_0
var y_1 *pkg_Data_Date.Constructor_Date = y_1_loop
_ = y_1
return (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_eqYear(), "eq"), gopurs_runtime.Int((*pkg_Data_Date.Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(x_0)}.UnsafePtr).V0), gopurs_runtime.Int((*pkg_Data_Date.Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(y_1)}.UnsafePtr).V0)).IntVal) != (0)), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_eqMonth(), "eq"), gopurs_runtime.Value{Type: 9, IntVal: int64((*pkg_Data_Date.Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(x_0)}.UnsafePtr).V1), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64((*pkg_Data_Date.Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(y_1)}.UnsafePtr).V1), UnsafePtr: nil}).IntVal) != (0))).IntVal) != (0)), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_eqDay(), "eq"), gopurs_runtime.Int((*pkg_Data_Date.Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(x_0)}.UnsafePtr).V2), gopurs_runtime.Int((*pkg_Data_Date.Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(y_1)}.UnsafePtr).V2)).IntVal) != (0))).IntVal) != (0)
}

func Call_mod__2185172824(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(pkg_Data_EuclideanRing.Get_intMod(), __eta0_0, __eta1_1)
}

func Call_mod__2579358968(dict_0_loop *pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_const__220790420(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__641934996(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_flip__1826582752(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__3525081280(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__2253242624(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_map__1165794789(dict_0_loop *pkg_Data_Functor.Constructor_Functor[*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__1171202917(dict_0_loop *pkg_Data_Functor.Constructor_Functor[*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__1307706501(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__1408505925(dict_0_loop *pkg_Data_Functor.Constructor_Functor[*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__243231988(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__2199395572(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__3447677596(v_0_loop gopurs_runtime.Value, v1_1_loop *pkg_Data_Maybe.Constructor_Just[int64]) *pkg_Data_Maybe.Constructor_Just[*pkg_Data_Date.Constructor_Date] {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *pkg_Data_Maybe.Constructor_Just[int64] = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Date.Constructor_Date](gopurs_runtime.Apply(v_0, gopurs_runtime.Int((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0.IntVal))))}})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Date.Constructor_Date]](__t0)
}

func Call_map__4155962236(v_0_loop gopurs_runtime.Value, v1_1_loop *pkg_Data_Maybe.Constructor_Just[int64]) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *pkg_Data_Maybe.Constructor_Just[int64] = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, gopurs_runtime.Int((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0.IntVal))})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0)
}

func Call_map__901270812(v_0_loop gopurs_runtime.Value, v1_1_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0)
}

func Call_conj__2927892844(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[int64]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[int64] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_conj__4093645121(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[*pkg_Data_Date.Constructor_Date]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[*pkg_Data_Date.Constructor_Date] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_conj__3676519832(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Bool(((__eta0_0.IntVal) != (0)) && ((__eta1_1.IntVal) != (0)))
}

func Call_conj__3472268504(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_disj__3676519832(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Bool(((__eta0_0.IntVal) != (0)) || ((__eta1_1.IntVal) != (0)))
}

func Call_disj__3472268504(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_not__3201284355(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Bool(((__eta0_0.IntVal) != (0)) != (true))
}

func Call_not__1505204753(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V4
}

func Call_fromJust__1577979644(v_0_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
__t0 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_fromJust__1791383420(_dollar__unused_0_loop gopurs_runtime.Value, v_1_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var v_1 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v_1_loop
_ = v_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr != nil) {
__t0 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_fromJust__4142563260(v_0_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
__t0 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_fromJust__965748316(v_0_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
__t0 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_fromMaybe__430429096(a_0_loop gopurs_runtime.Value, v2_1_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v2_1 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_1_loop
_ = v2_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr == nil) {
__t0 = a_0
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr != nil) {
__t0 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_fromMaybe__18840980(a_0_loop gopurs_runtime.Value, v2_1_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v2_1 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_1_loop
_ = v2_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr == nil) {
__t0 = a_0
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr != nil) {
__t0 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_fromMaybe__737056608(a_0_loop gopurs_runtime.Value, v2_1_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v2_1 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_1_loop
_ = v2_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr == nil) {
__t0 = a_0
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr != nil) {
__t0 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_isNothing__1358705270(v2_0_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) bool {
var v2_0 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return (__t0.IntVal) != (0)
}

func Call_isNothing__4206805139(v2_0_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) bool {
var v2_0 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return (__t0.IntVal) != (0)
}

func Call_isNothing__2787066607(v2_0_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) bool {
var v2_0 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return (__t0.IntVal) != (0)
}

func Call_isNothing__323776123(v2_0_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) bool {
var v2_0 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return (__t0.IntVal) != (0)
}

func Call_maybe__1594528518(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_maybe__3658316244(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_compare__669572705(dict_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_compare__372254389(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply5(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __eta0_0, __eta1_1)
}

func Call_compare__821463600(dict_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_compare__696857420(x_0_loop uint32, y_1_loop uint32) uint32 {
var x_0 uint32 = x_0_loop
_ = x_0
var y_1 uint32 = y_1_loop
_ = y_1
var __t11 gopurs_runtime.Value
{
if (x_0 == 1908470532) {
var __t0 uint32
{
if (y_1 == 1908470532) {
__t0 = 902936544
goto end_branch_0
} else {

}
}
{
__t0 = 1527465420
}
end_branch_0:
__t11 = gopurs_runtime.Value{Type: 9, IntVal: int64(__t0), UnsafePtr: nil}
goto end_branch_11
} else {

}
}
{
if (y_1 == 1908470532) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
goto end_branch_11
} else {

}
}
{
if (x_0 == 2455627378) {
var __t1 uint32
{
if (y_1 == 2455627378) {
__t1 = 902936544
goto end_branch_1
} else {

}
}
{
__t1 = 1527465420
}
end_branch_1:
__t11 = gopurs_runtime.Value{Type: 9, IntVal: int64(__t1), UnsafePtr: nil}
goto end_branch_11
} else {

}
}
{
if (y_1 == 2455627378) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
goto end_branch_11
} else {

}
}
{
if (x_0 == 4162469099) {
var __t2 uint32
{
if (y_1 == 4162469099) {
__t2 = 902936544
goto end_branch_2
} else {

}
}
{
__t2 = 1527465420
}
end_branch_2:
__t11 = gopurs_runtime.Value{Type: 9, IntVal: int64(__t2), UnsafePtr: nil}
goto end_branch_11
} else {

}
}
{
if (y_1 == 4162469099) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
goto end_branch_11
} else {

}
}
{
if (x_0 == 1692989816) {
var __t3 uint32
{
if (y_1 == 1692989816) {
__t3 = 902936544
goto end_branch_3
} else {

}
}
{
__t3 = 1527465420
}
end_branch_3:
__t11 = gopurs_runtime.Value{Type: 9, IntVal: int64(__t3), UnsafePtr: nil}
goto end_branch_11
} else {

}
}
{
if (y_1 == 1692989816) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
goto end_branch_11
} else {

}
}
{
if (x_0 == 330658827) {
var __t4 uint32
{
if (y_1 == 330658827) {
__t4 = 902936544
goto end_branch_4
} else {

}
}
{
__t4 = 1527465420
}
end_branch_4:
__t11 = gopurs_runtime.Value{Type: 9, IntVal: int64(__t4), UnsafePtr: nil}
goto end_branch_11
} else {

}
}
{
if (y_1 == 330658827) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
goto end_branch_11
} else {

}
}
{
if (x_0 == 4067355978) {
var __t5 uint32
{
if (y_1 == 4067355978) {
__t5 = 902936544
goto end_branch_5
} else {

}
}
{
__t5 = 1527465420
}
end_branch_5:
__t11 = gopurs_runtime.Value{Type: 9, IntVal: int64(__t5), UnsafePtr: nil}
goto end_branch_11
} else {

}
}
{
if (y_1 == 4067355978) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
goto end_branch_11
} else {

}
}
{
if (x_0 == 2276710548) {
var __t6 uint32
{
if (y_1 == 2276710548) {
__t6 = 902936544
goto end_branch_6
} else {

}
}
{
__t6 = 1527465420
}
end_branch_6:
__t11 = gopurs_runtime.Value{Type: 9, IntVal: int64(__t6), UnsafePtr: nil}
goto end_branch_11
} else {

}
}
{
if (y_1 == 2276710548) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
goto end_branch_11
} else {

}
}
{
if (x_0 == 243771071) {
var __t7 uint32
{
if (y_1 == 243771071) {
__t7 = 902936544
goto end_branch_7
} else {

}
}
{
__t7 = 1527465420
}
end_branch_7:
__t11 = gopurs_runtime.Value{Type: 9, IntVal: int64(__t7), UnsafePtr: nil}
goto end_branch_11
} else {

}
}
{
if (y_1 == 243771071) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
goto end_branch_11
} else {

}
}
{
if (x_0 == 215731793) {
var __t8 uint32
{
if (y_1 == 215731793) {
__t8 = 902936544
goto end_branch_8
} else {

}
}
{
__t8 = 1527465420
}
end_branch_8:
__t11 = gopurs_runtime.Value{Type: 9, IntVal: int64(__t8), UnsafePtr: nil}
goto end_branch_11
} else {

}
}
{
if (y_1 == 215731793) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
goto end_branch_11
} else {

}
}
{
if (x_0 == 8639228) {
var __t9 uint32
{
if (y_1 == 8639228) {
__t9 = 902936544
goto end_branch_9
} else {

}
}
{
__t9 = 1527465420
}
end_branch_9:
__t11 = gopurs_runtime.Value{Type: 9, IntVal: int64(__t9), UnsafePtr: nil}
goto end_branch_11
} else {

}
}
{
if (y_1 == 8639228) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
goto end_branch_11
} else {

}
}
{
if (x_0 == 49471444) {
var __t10 uint32
{
if (y_1 == 49471444) {
__t10 = 902936544
goto end_branch_10
} else {

}
}
{
__t10 = 1527465420
}
end_branch_10:
__t11 = gopurs_runtime.Value{Type: 9, IntVal: int64(__t10), UnsafePtr: nil}
goto end_branch_11
} else {

}
}
{
if (y_1 == 49471444) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
goto end_branch_11
} else {

}
}
{
if ((x_0 == 3889233761)) && ((y_1 == 3889233761)) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
goto end_branch_11
} else {

}
}
{
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_11:
return uint32(__t11.IntVal)
}

func Call_greaterThan__3259097883(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
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
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
return __t1
}

func Call_greaterThan__4087042607(a1_0_loop gopurs_runtime.Value, a2_1_loop gopurs_runtime.Value) bool {
var a1_0 gopurs_runtime.Value = a1_0_loop
_ = a1_0
var a2_1 gopurs_runtime.Value = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0.IntVal) > (a2_1.IntVal) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
return __t0
}

func Call_greaterThan__1409282474(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
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
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
return __t1
}

func Call_greaterThan__2400628110(a1_0_loop gopurs_runtime.Value, a2_1_loop gopurs_runtime.Value) bool {
var a1_0 gopurs_runtime.Value = a1_0_loop
_ = a1_0
var a2_1 gopurs_runtime.Value = a2_1_loop
_ = a2_1
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[*pkg_Data_Maybe.Constructor_Just[int64]]](pkg_Data_Date.Get_ordMaybe()).V1, a1_0, a2_1)
if (uint32(__t_tag_0.IntVal) == 380165415) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
return __t1
}

func Call_greaterThanOrEq__1710332219(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
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

func Call_greaterThanOrEq__4087042607(a1_0_loop gopurs_runtime.Value, a2_1_loop gopurs_runtime.Value) bool {
var a1_0 gopurs_runtime.Value = a1_0_loop
_ = a1_0
var a2_1 gopurs_runtime.Value = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0.IntVal) < (a2_1.IntVal) {
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

func Call_lessThan__4087042607(a1_0_loop gopurs_runtime.Value, a2_1_loop gopurs_runtime.Value) bool {
var a1_0 gopurs_runtime.Value = a1_0_loop
_ = a1_0
var a2_1 gopurs_runtime.Value = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0.IntVal) < (a2_1.IntVal) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
return __t0
}

func Call_lessThan__1409282474(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
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
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
return __t1
}

func Call_lessThanOrEq__1710332219(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
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

func Call_negate__2635823316(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Int(-(__eta0_0.IntVal))
}

func Call_negate__1364373265(dictRing_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictRing_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dictRing_0_loop
_ = dictRing_0
Semiring0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](gopurs_runtime.Apply(dictRing_0.V0, gopurs_runtime.Value{}))
_ = Semiring0_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictRing_0.V1, Semiring0_1_0.V3, a_2)
})
}

func Call_sub__2927892844(dict_0_loop *pkg_Data_Ring.Constructor_Ring[int64]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ring.Constructor_Ring[int64] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_sub__1124926121(dict_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_sub__1841809173(dict_0_loop *pkg_Data_Ring.Constructor_Ring[uint32]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ring.Constructor_Ring[uint32] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_sub__1043827704(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Int((__eta0_0.IntVal) - (__eta1_1.IntVal))
}

func Call_sub__3675938712(dict_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_add__2927892844(dict_0_loop *pkg_Data_Semiring.Constructor_Semiring[int64]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semiring.Constructor_Semiring[int64] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_add__1841809173(dict_0_loop *pkg_Data_Semiring.Constructor_Semiring[uint32]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semiring.Constructor_Semiring[uint32] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_add__560788792(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Int((__eta0_0.IntVal) + (__eta1_1.IntVal))
}

func Call_add__1614463960(dict_0_loop *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_zero__1556010056(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "zero")
}

func Call_zero__1204848985(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "zero")
}


