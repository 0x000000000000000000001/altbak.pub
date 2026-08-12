package Data_String_Gen

import (
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Control_Monad_Gen "gopurs/output/Control.Monad.Gen"
	pkg_Control_Monad_Gen_Class "gopurs/output/Control.Monad.Gen.Class"
	pkg_Control_Monad_Rec_Class "gopurs/output/Control.Monad.Rec.Class"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Char_Gen "gopurs/output/Data.Char.Gen"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_String_CodeUnits "gopurs/output/Data.String.CodeUnits"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unfoldable "gopurs/output/Data.Unfoldable"
	pkg_Data_Unfoldable1 "gopurs/output/Data.Unfoldable1"
	pkg_Partial_Unsafe "gopurs/output/Partial.Unsafe"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_genString gopurs_runtime.Value
var once_genString sync.Once
func Get_genString() gopurs_runtime.Value {
	once_genString.Do(func() {
		cache_genString = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genString(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dictMonadGen_1_box))
})
	})
	return cache_genString
}

var cache_genUnicodeString gopurs_runtime.Value
var once_genUnicodeString sync.Once
func Get_genUnicodeString() gopurs_runtime.Value {
	once_genUnicodeString.Do(func() {
		cache_genUnicodeString = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genUnicodeString(dictMonadRec_0_box, dictMonadGen_1_box)
})
	})
	return cache_genUnicodeString
}

var cache_genDigitString gopurs_runtime.Value
var once_genDigitString sync.Once
func Get_genDigitString() gopurs_runtime.Value {
	once_genDigitString.Do(func() {
		cache_genDigitString = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genDigitString(dictMonadRec_0_box, dictMonadGen_1_box)
})
	})
	return cache_genDigitString
}

var cache_genAsciiString_prime gopurs_runtime.Value
var once_genAsciiString_prime sync.Once
func Get_genAsciiString_prime() gopurs_runtime.Value {
	once_genAsciiString_prime.Do(func() {
		cache_genAsciiString_prime = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genAsciiString_prime(dictMonadRec_0_box, dictMonadGen_1_box)
})
	})
	return cache_genAsciiString_prime
}

var cache_genAsciiString gopurs_runtime.Value
var once_genAsciiString sync.Once
func Get_genAsciiString() gopurs_runtime.Value {
	once_genAsciiString.Do(func() {
		cache_genAsciiString = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genAsciiString(dictMonadRec_0_box, dictMonadGen_1_box)
})
	})
	return cache_genAsciiString
}

var cache_genAlphaUppercaseString gopurs_runtime.Value
var once_genAlphaUppercaseString sync.Once
func Get_genAlphaUppercaseString() gopurs_runtime.Value {
	once_genAlphaUppercaseString.Do(func() {
		cache_genAlphaUppercaseString = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genAlphaUppercaseString(dictMonadRec_0_box, dictMonadGen_1_box)
})
	})
	return cache_genAlphaUppercaseString
}

var cache_genAlphaString gopurs_runtime.Value
var once_genAlphaString sync.Once
func Get_genAlphaString() gopurs_runtime.Value {
	once_genAlphaString.Do(func() {
		cache_genAlphaString = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genAlphaString(dictMonadRec_0_box, dictMonadGen_1_box)
})
	})
	return cache_genAlphaString
}

var cache_genAlphaLowercaseString gopurs_runtime.Value
var once_genAlphaLowercaseString sync.Once
func Get_genAlphaLowercaseString() gopurs_runtime.Value {
	once_genAlphaLowercaseString.Do(func() {
		cache_genAlphaLowercaseString = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genAlphaLowercaseString(dictMonadRec_0_box, dictMonadGen_1_box)
})
	})
	return cache_genAlphaLowercaseString
}

var cache_bind__2668924679 gopurs_runtime.Value
var once_bind__2668924679 sync.Once
func Get_bind__2668924679() gopurs_runtime.Value {
	once_bind__2668924679.Do(func() {
		cache_bind__2668924679 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__2668924679(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__2668924679
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

var cache_resize__4113973243 gopurs_runtime.Value
var once_resize__4113973243 sync.Once
func Get_resize__4113973243() gopurs_runtime.Value {
	once_resize__4113973243.Do(func() {
		cache_resize__4113973243 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_resize__4113973243(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_resize__4113973243
}

var cache_resize__1050945947 gopurs_runtime.Value
var once_resize__1050945947 sync.Once
func Get_resize__1050945947() gopurs_runtime.Value {
	once_resize__1050945947.Do(func() {
		cache_resize__1050945947 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_resize__1050945947(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_resize__1050945947
}

var cache_sized__120035991 gopurs_runtime.Value
var once_sized__120035991 sync.Once
func Get_sized__120035991() gopurs_runtime.Value {
	once_sized__120035991.Do(func() {
		cache_sized__120035991 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sized__120035991(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sized__120035991
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

var cache_unfoldable__1024209525 gopurs_runtime.Value
var once_unfoldable__1024209525 sync.Once
func Get_unfoldable__1024209525() gopurs_runtime.Value {
	once_unfoldable__1024209525.Do(func() {
		cache_unfoldable__1024209525 = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfoldable__1024209525(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dictMonadGen_1_box))
})
	})
	return cache_unfoldable__1024209525
}

var cache_unfoldable__542018773 gopurs_runtime.Value
var once_unfoldable__542018773 sync.Once
func Get_unfoldable__542018773() gopurs_runtime.Value {
	once_unfoldable__542018773.Do(func() {
		cache_unfoldable__542018773 = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfoldable__542018773(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dictMonadGen_1_box))
})
	})
	return cache_unfoldable__542018773
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

var cache_const__1243414737 gopurs_runtime.Value
var once_const__1243414737 sync.Once
func Get_const__1243414737() gopurs_runtime.Value {
	once_const__1243414737.Do(func() {
		cache_const__1243414737 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__1243414737(a_0_box, v_1_box)
})
	})
	return cache_const__1243414737
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

var cache_map__38773460 gopurs_runtime.Value
var once_map__38773460 sync.Once
func Get_map__38773460() gopurs_runtime.Value {
	once_map__38773460.Do(func() {
		cache_map__38773460 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__38773460(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__38773460
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

var cache_isNothing__2591355336 gopurs_runtime.Value
var once_isNothing__2591355336 sync.Once
func Get_isNothing__2591355336() gopurs_runtime.Value {
	once_isNothing__2591355336.Do(func() {
		cache_isNothing__2591355336 = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_isNothing__2591355336(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_0_box)))
})
	})
	return cache_isNothing__2591355336
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

var cache_max__2538992856 gopurs_runtime.Value
var once_max__2538992856 sync.Once
func Get_max__2538992856() gopurs_runtime.Value {
	once_max__2538992856.Do(func() {
		cache_max__2538992856 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_max__2538992856(x_0_box, y_1_box)
})
	})
	return cache_max__2538992856
}

var cache_max__2767602680 gopurs_runtime.Value
var once_max__2767602680 sync.Once
func Get_max__2767602680() gopurs_runtime.Value {
	once_max__2767602680.Do(func() {
		cache_max__2767602680 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_max__2767602680(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), x_1_box, y_2_box)
})
	})
	return cache_max__2767602680
}

var cache_sub__1043827704 gopurs_runtime.Value
var once_sub__1043827704 sync.Once
func Get_sub__1043827704() gopurs_runtime.Value {
	once_sub__1043827704.Do(func() {
		cache_sub__1043827704 = pkg_Data_Ring.Get_intSub()
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

var cache_genString__3531165174 gopurs_runtime.Value
var once_genString__3531165174 sync.Once
func Get_genString__3531165174() gopurs_runtime.Value {
	once_genString__3531165174.Do(func() {
		cache_genString__3531165174 = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genString__3531165174(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dictMonadGen_1_box))
})
	})
	return cache_genString__3531165174
}

var cache_fst__20422131 gopurs_runtime.Value
var once_fst__20422131 sync.Once
func Get_fst__20422131() gopurs_runtime.Value {
	once_fst__20422131.Do(func() {
		cache_fst__20422131 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fst__20422131(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_fst__20422131
}

var cache_snd__20422131 gopurs_runtime.Value
var once_snd__20422131 sync.Once
func Get_snd__20422131() gopurs_runtime.Value {
	once_snd__20422131.Do(func() {
		cache_snd__20422131 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_snd__20422131(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_snd__20422131
}

var cache_unfoldableArray__644327338 gopurs_runtime.Value
var once_unfoldableArray__644327338 sync.Once
func Get_unfoldableArray__644327338() gopurs_runtime.Value {
	once_unfoldableArray__644327338.Do(func() {
		cache_unfoldableArray__644327338 = gopurs_runtime.RecordDict2("Unfoldable10", "unfoldr", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unfoldable1.Get_unfoldable1Array()
}), gopurs_runtime.Apply4(pkg_Data_Unfoldable.Get_unfoldrArrayImpl(), pkg_Data_Maybe.Get_isNothing(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr != nil) {
__t0 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_0.UnsafePtr).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), pkg_Data_Tuple.Get_fst(), pkg_Data_Tuple.Get_snd()))
	})
	return cache_unfoldableArray__644327338
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

var cache_unfoldable1Array__4196906331 gopurs_runtime.Value
var once_unfoldable1Array__4196906331 sync.Once
func Get_unfoldable1Array__4196906331() gopurs_runtime.Value {
	once_unfoldable1Array__4196906331.Do(func() {
		cache_unfoldable1Array__4196906331 = gopurs_runtime.RecordDict1("unfoldr1", gopurs_runtime.Apply4(pkg_Data_Unfoldable1.Get_unfoldr1ArrayImpl(), pkg_Data_Maybe.Get_isNothing(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr != nil) {
__t0 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_0.UnsafePtr).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), pkg_Data_Tuple.Get_fst(), pkg_Data_Tuple.Get_snd()))
	})
	return cache_unfoldable1Array__4196906331
}

var cache_unsafePartial__1306634845 gopurs_runtime.Value
var once_unsafePartial__1306634845 sync.Once
func Get_unsafePartial__1306634845() gopurs_runtime.Value {
	once_unsafePartial__1306634845.Do(func() {
		cache_unsafePartial__1306634845 = pkg_Partial_Unsafe.Get__unsafePartial()
	})
	return cache_unsafePartial__1306634845
}

var cache_unsafePartial__1090765981 gopurs_runtime.Value
var once_unsafePartial__1090765981 sync.Once
func Get_unsafePartial__1090765981() gopurs_runtime.Value {
	once_unsafePartial__1090765981.Do(func() {
		cache_unsafePartial__1090765981 = pkg_Partial_Unsafe.Get__unsafePartial()
	})
	return cache_unsafePartial__1090765981
}

func Call_genString(dictMonadRec_0_loop *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value], dictMonadGen_1_loop *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadRec_0 *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value] = dictMonadGen_1_loop
_ = dictMonadGen_1
Monad0_2_0 := gopurs_runtime.Apply(dictMonadGen_1.V0, gopurs_runtime.Value{})
_ = Monad0_2_0
Bind1_3_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_1
Functor0_4_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_2
return gopurs_runtime.Func(func(genChar_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictMonadGen_1.V5, gopurs_runtime.Func(func(size_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_3_1.V1, gopurs_runtime.Apply2(dictMonadGen_1.V3, gopurs_runtime.Int(1), Call_max__2538992856(gopurs_runtime.Int(1), size_6)), gopurs_runtime.Func(func(newSize_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictMonadGen_1.V4, gopurs_runtime.Apply(Get_const__1243414737(), newSize_7), gopurs_runtime.Apply2(Functor0_4_2.V0, pkg_Data_String_CodeUnits.Get_fromCharArray(), gopurs_runtime.Apply4(pkg_Control_Monad_Gen.Get_unfoldable(), gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer(dictMonadRec_0)}, gopurs_runtime.Value{Type: 9, IntVal: 2254593219, UnsafePtr: unsafe.Pointer(dictMonadGen_1)}, pkg_Data_Unfoldable.Get_unfoldableArray(), genChar_5)))
}))
}))
})
}

func Call_genUnicodeString(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
_ = dictMonadGen_1
return gopurs_runtime.Apply(Call_genString(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dictMonadRec_0), gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dictMonadGen_1)), gopurs_runtime.Apply(pkg_Data_Char_Gen.Get_genUnicodeChar(), dictMonadGen_1))
}

func Call_genDigitString(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
_ = dictMonadGen_1
return gopurs_runtime.Apply(Call_genString(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dictMonadRec_0), gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dictMonadGen_1)), gopurs_runtime.Apply(pkg_Data_Char_Gen.Get_genDigitChar(), dictMonadGen_1))
}

func Call_genAsciiString_prime(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
_ = dictMonadGen_1
return gopurs_runtime.Apply(Call_genString(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dictMonadRec_0), gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dictMonadGen_1)), gopurs_runtime.Apply(pkg_Data_Char_Gen.Get_genAsciiChar_prime(), dictMonadGen_1))
}

func Call_genAsciiString(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
_ = dictMonadGen_1
return gopurs_runtime.Apply(Call_genString(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dictMonadRec_0), gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dictMonadGen_1)), gopurs_runtime.Apply(pkg_Data_Char_Gen.Get_genAsciiChar(), dictMonadGen_1))
}

func Call_genAlphaUppercaseString(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
_ = dictMonadGen_1
return gopurs_runtime.Apply(Call_genString(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dictMonadRec_0), gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dictMonadGen_1)), gopurs_runtime.Apply(pkg_Data_Char_Gen.Get_genAlphaUppercase(), dictMonadGen_1))
}

func Call_genAlphaString(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
_ = dictMonadGen_1
return gopurs_runtime.Apply(Call_genString(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dictMonadRec_0), gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dictMonadGen_1)), gopurs_runtime.Apply(pkg_Data_Char_Gen.Get_genAlpha(), dictMonadGen_1))
}

func Call_genAlphaLowercaseString(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
_ = dictMonadGen_1
return gopurs_runtime.Apply(Call_genString(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dictMonadRec_0), gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dictMonadGen_1)), gopurs_runtime.Apply(pkg_Data_Char_Gen.Get_genAlphaLowercase(), dictMonadGen_1))
}

func Call_bind__2668924679(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
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

func Call_chooseInt__1063828903(dict_0_loop *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_resize__4113973243(dict_0_loop *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V4
}

func Call_resize__1050945947(dict_0_loop *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V4
}

func Call_sized__120035991(dict_0_loop *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V5
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

func Call_unfoldable__1024209525(dictMonadRec_0_loop *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value], dictMonadGen_1_loop *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
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
return gopurs_runtime.Func(func(gen_6 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_5 := gopurs_runtime.Apply(dictMonadRec_0.V1, gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(Call_lessThanOrEq__4087042607((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V1, gopurs_runtime.Int(0))).IntVal) != (0) {
__t8 = gopurs_runtime.Apply(pure_3_1, gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0})})
goto end_branch_8
} else {

}
}
{
__local_var_8_6 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0
_ = __local_var_8_6
__local_var_9_7 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V1
_ = __local_var_9_7
__t8 = gopurs_runtime.Apply2(Bind1_4_2.V1, gen_6, gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_1, gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Gen.Constructor_Cons[gopurs_runtime.Value]{1, x_10, gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen.Constructor_Cons[gopurs_runtime.Value]](__local_var_8_6)})}, gopurs_runtime.Apply2(Get_sub__1043827704(), __local_var_9_7, gopurs_runtime.Int(1))})}})})
}))
}
end_branch_8:
return __t8
}))
_ = __local_var_7_5
__local_var_8_9 := gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*pkg_Control_Monad_Gen.Constructor_Cons[gopurs_runtime.Value])(nil))})
_ = __local_var_8_9
return gopurs_runtime.Apply2(Functor0_5_3.V0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Unfoldable.Get_unfoldableArray(), "unfoldr"), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 759514854 && v_7.UnsafePtr == nil) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_4
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 759514854 && v_7.UnsafePtr != nil) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Control_Monad_Gen.Constructor_Cons[gopurs_runtime.Value])(v_7.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*pkg_Control_Monad_Gen.Constructor_Cons[gopurs_runtime.Value])(v_7.UnsafePtr).V1)}})}})}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, *pkg_Control_Monad_Gen.Constructor_Cons[gopurs_runtime.Value]]]](__t4))}
})), gopurs_runtime.Apply(dictMonadGen_1.V5, gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_5, gopurs_runtime.Apply(__local_var_8_9, x_9))
})))
})
}

func Call_unfoldable__542018773(dictMonadRec_0_loop *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value], dictMonadGen_1_loop *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
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
return gopurs_runtime.Apply(pure_3_1, gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Gen.Constructor_Cons[gopurs_runtime.Value]{1, x_11, gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen.Constructor_Cons[gopurs_runtime.Value]](__local_var_9_6)})}, gopurs_runtime.Apply2(Get_sub__1043827704(), __local_var_10_7, gopurs_runtime.Int(1))})}})})
}))
}
end_branch_8:
return __t8
}))
_ = __local_var_8_5
__local_var_9_9 := gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*pkg_Control_Monad_Gen.Constructor_Cons[gopurs_runtime.Value])(nil))})
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
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Control_Monad_Gen.Constructor_Cons[gopurs_runtime.Value])(v_8.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*pkg_Control_Monad_Gen.Constructor_Cons[gopurs_runtime.Value])(v_8.UnsafePtr).V1)}})}})}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, *pkg_Control_Monad_Gen.Constructor_Cons[gopurs_runtime.Value]]]](__t4))}
})), gopurs_runtime.Apply(dictMonadGen_1.V5, gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_5, gopurs_runtime.Apply(__local_var_9_9, x_10))
})))
})
})
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

func Call_const__1243414737(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_map__2199395572(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__38773460(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__1504457012(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
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

func Call_isNothing__2591355336(v2_0_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) bool {
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

func Call_compare__821463600(dict_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
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

func Call_max__2538992856(x_0_loop gopurs_runtime.Value, y_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var y_1 gopurs_runtime.Value = y_1_loop
_ = y_1
v_2_0 := gopurs_runtime.Apply5(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, x_0, y_1)
_ = v_2_0
var __t1 gopurs_runtime.Value
{
if (uint32(v_2_0.IntVal) == 1527465420) {
__t1 = y_1
goto end_branch_1
} else {

}
}
{
if (uint32(v_2_0.IntVal) == 902936544) {
__t1 = x_0
goto end_branch_1
} else {

}
}
{
if (uint32(v_2_0.IntVal) == 380165415) {
__t1 = x_0
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}

func Call_max__2767602680(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
v_3_0 := gopurs_runtime.Apply2(dictOrd_0.V1, x_1, y_2)
_ = v_3_0
var __t1 gopurs_runtime.Value
{
if (uint32(v_3_0.IntVal) == 1527465420) {
__t1 = y_2
goto end_branch_1
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 902936544) {
__t1 = x_1
goto end_branch_1
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 380165415) {
__t1 = x_1
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}

func Call_sub__3675938712(dict_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_genString__3531165174(dictMonadRec_0_loop *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value], dictMonadGen_1_loop *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadRec_0 *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value] = dictMonadGen_1_loop
_ = dictMonadGen_1
Monad0_2_0 := gopurs_runtime.Apply(dictMonadGen_1.V0, gopurs_runtime.Value{})
_ = Monad0_2_0
Bind1_3_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_1
Functor0_4_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_2
return gopurs_runtime.Func(func(genChar_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictMonadGen_1.V5, gopurs_runtime.Func(func(size_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_3_1.V1, gopurs_runtime.Apply2(dictMonadGen_1.V3, gopurs_runtime.Int(1), Call_max__2538992856(gopurs_runtime.Int(1), size_6)), gopurs_runtime.Func(func(newSize_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictMonadGen_1.V4, gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return newSize_7
}), gopurs_runtime.Apply2(Functor0_4_2.V0, pkg_Data_String_CodeUnits.Get_fromCharArray(), gopurs_runtime.Apply4(pkg_Control_Monad_Gen.Get_unfoldable(), gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer(dictMonadRec_0)}, gopurs_runtime.Value{Type: 9, IntVal: 2254593219, UnsafePtr: unsafe.Pointer(dictMonadGen_1)}, pkg_Data_Unfoldable.Get_unfoldableArray(), genChar_5)))
}))
}))
})
}

func Call_fst__20422131(v_0_loop *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
}

func Call_snd__20422131(v_0_loop *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
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


