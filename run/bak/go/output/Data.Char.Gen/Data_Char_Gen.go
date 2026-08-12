package Data_Char_Gen

import (
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Control_Monad_Gen "gopurs/output/Control.Monad.Gen"
	pkg_Control_Monad_Gen_Class "gopurs/output/Control.Monad.Gen.Class"
	pkg_Data_Bounded "gopurs/output/Data.Bounded"
	pkg_Data_Enum "gopurs/output/Data.Enum"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Newtype "gopurs/output/Data.Newtype"
	pkg_Data_NonEmpty "gopurs/output/Data.NonEmpty"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Semigroup_Foldable "gopurs/output/Data.Semigroup.Foldable"
	pkg_Data_Semigroup_Last "gopurs/output/Data.Semigroup.Last"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_foldable1NonEmpty gopurs_runtime.Value
var once_foldable1NonEmpty sync.Once
func Get_foldable1NonEmpty() gopurs_runtime.Value {
	once_foldable1NonEmpty.Do(func() {
		cache_foldable1NonEmpty = gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup_Foldable.Constructor_Foldable1[*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.Apply(pkg_Data_NonEmpty.Get_foldable1NonEmpty(), pkg_Data_Foldable.Get_foldableArray())))}
	})
	return cache_foldable1NonEmpty
}

var cache_genUnicodeChar gopurs_runtime.Value
var once_genUnicodeChar sync.Once
func Get_genUnicodeChar() gopurs_runtime.Value {
	once_genUnicodeChar.Do(func() {
		cache_genUnicodeChar = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genUnicodeChar(dictMonadGen_0_box)
})
	})
	return cache_genUnicodeChar
}

var cache_genDigitChar gopurs_runtime.Value
var once_genDigitChar sync.Once
func Get_genDigitChar() gopurs_runtime.Value {
	once_genDigitChar.Do(func() {
		cache_genDigitChar = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genDigitChar(dictMonadGen_0_box)
})
	})
	return cache_genDigitChar
}

var cache_genAsciiChar_prime gopurs_runtime.Value
var once_genAsciiChar_prime sync.Once
func Get_genAsciiChar_prime() gopurs_runtime.Value {
	once_genAsciiChar_prime.Do(func() {
		cache_genAsciiChar_prime = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genAsciiChar_prime(dictMonadGen_0_box)
})
	})
	return cache_genAsciiChar_prime
}

var cache_genAsciiChar gopurs_runtime.Value
var once_genAsciiChar sync.Once
func Get_genAsciiChar() gopurs_runtime.Value {
	once_genAsciiChar.Do(func() {
		cache_genAsciiChar = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genAsciiChar(dictMonadGen_0_box)
})
	})
	return cache_genAsciiChar
}

var cache_genAlphaUppercase gopurs_runtime.Value
var once_genAlphaUppercase sync.Once
func Get_genAlphaUppercase() gopurs_runtime.Value {
	once_genAlphaUppercase.Do(func() {
		cache_genAlphaUppercase = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genAlphaUppercase(dictMonadGen_0_box)
})
	})
	return cache_genAlphaUppercase
}

var cache_genAlphaLowercase gopurs_runtime.Value
var once_genAlphaLowercase sync.Once
func Get_genAlphaLowercase() gopurs_runtime.Value {
	once_genAlphaLowercase.Do(func() {
		cache_genAlphaLowercase = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genAlphaLowercase(dictMonadGen_0_box)
})
	})
	return cache_genAlphaLowercase
}

var cache_genAlpha gopurs_runtime.Value
var once_genAlpha sync.Once
func Get_genAlpha() gopurs_runtime.Value {
	once_genAlpha.Do(func() {
		cache_genAlpha = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genAlpha(dictMonadGen_0_box)
})
	})
	return cache_genAlpha
}

var cache_bind__1858449959 gopurs_runtime.Value
var once_bind__1858449959 sync.Once
func Get_bind__1858449959() gopurs_runtime.Value {
	once_bind__1858449959.Do(func() {
		cache_bind__1858449959 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__1858449959(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__1858449959
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

var cache_fromIndex__3111933544 gopurs_runtime.Value
var once_fromIndex__3111933544 sync.Once
func Get_fromIndex__3111933544() gopurs_runtime.Value {
	once_fromIndex__3111933544.Do(func() {
		cache_fromIndex__3111933544 = gopurs_runtime.Func(func(dictFoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromIndex__3111933544(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box))
})
	})
	return cache_fromIndex__3111933544
}

var cache_fromIndex__4031440680 gopurs_runtime.Value
var once_fromIndex__4031440680 sync.Once
func Get_fromIndex__4031440680() gopurs_runtime.Value {
	once_fromIndex__4031440680.Do(func() {
		cache_fromIndex__4031440680 = gopurs_runtime.Func(func(dictFoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromIndex__4031440680(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box))
})
	})
	return cache_fromIndex__4031440680
}

var cache_oneOf__2265861353 gopurs_runtime.Value
var once_oneOf__2265861353 sync.Once
func Get_oneOf__2265861353() gopurs_runtime.Value {
	once_oneOf__2265861353.Do(func() {
		cache_oneOf__2265861353 = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_oneOf__2265861353(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dictMonadGen_0_box))
})
	})
	return cache_oneOf__2265861353
}

var cache_oneOf__1433237231 gopurs_runtime.Value
var once_oneOf__1433237231 sync.Once
func Get_oneOf__1433237231() gopurs_runtime.Value {
	once_oneOf__1433237231.Do(func() {
		cache_oneOf__1433237231 = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_oneOf__1433237231(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dictMonadGen_0_box))
})
	})
	return cache_oneOf__1433237231
}

var cache_bottom__338427193 gopurs_runtime.Value
var once_bottom__338427193 sync.Once
func Get_bottom__338427193() gopurs_runtime.Value {
	once_bottom__338427193.Do(func() {
		cache_bottom__338427193 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bottom__338427193(dict_0_box)
})
	})
	return cache_bottom__338427193
}

var cache_top__338427193 gopurs_runtime.Value
var once_top__338427193 sync.Once
func Get_top__338427193() gopurs_runtime.Value {
	once_top__338427193.Do(func() {
		cache_top__338427193 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_top__338427193(dict_0_box)
})
	})
	return cache_top__338427193
}

var cache_genAlphaLowercase__4294897069 gopurs_runtime.Value
var once_genAlphaLowercase__4294897069 sync.Once
func Get_genAlphaLowercase__4294897069() gopurs_runtime.Value {
	once_genAlphaLowercase__4294897069.Do(func() {
		cache_genAlphaLowercase__4294897069 = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genAlphaLowercase__4294897069(dictMonadGen_0_box)
})
	})
	return cache_genAlphaLowercase__4294897069
}

var cache_defaultPred__1581620096 gopurs_runtime.Value
var once_defaultPred__1581620096 sync.Once
func Get_defaultPred__1581620096() gopurs_runtime.Value {
	once_defaultPred__1581620096.Do(func() {
		cache_defaultPred__1581620096 = gopurs_runtime.Func3(func(toEnum_prime_0_box gopurs_runtime.Value, fromEnum_prime_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_defaultPred__1581620096(toEnum_prime_0_box, fromEnum_prime_1_box, a_2_box))}
})
	})
	return cache_defaultPred__1581620096
}

var cache_defaultPred__2204581824 gopurs_runtime.Value
var once_defaultPred__2204581824 sync.Once
func Get_defaultPred__2204581824() gopurs_runtime.Value {
	once_defaultPred__2204581824.Do(func() {
		cache_defaultPred__2204581824 = gopurs_runtime.Func3(func(toEnum_prime_0_box gopurs_runtime.Value, fromEnum_prime_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_defaultPred__2204581824(toEnum_prime_0_box, fromEnum_prime_1_box, a_2_box))}
})
	})
	return cache_defaultPred__2204581824
}

var cache_defaultSucc__1581620096 gopurs_runtime.Value
var once_defaultSucc__1581620096 sync.Once
func Get_defaultSucc__1581620096() gopurs_runtime.Value {
	once_defaultSucc__1581620096.Do(func() {
		cache_defaultSucc__1581620096 = gopurs_runtime.Func3(func(toEnum_prime_0_box gopurs_runtime.Value, fromEnum_prime_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_defaultSucc__1581620096(toEnum_prime_0_box, fromEnum_prime_1_box, a_2_box))}
})
	})
	return cache_defaultSucc__1581620096
}

var cache_defaultSucc__2204581824 gopurs_runtime.Value
var once_defaultSucc__2204581824 sync.Once
func Get_defaultSucc__2204581824() gopurs_runtime.Value {
	once_defaultSucc__2204581824.Do(func() {
		cache_defaultSucc__2204581824 = gopurs_runtime.Func3(func(toEnum_prime_0_box gopurs_runtime.Value, fromEnum_prime_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_defaultSucc__2204581824(toEnum_prime_0_box, fromEnum_prime_1_box, a_2_box))}
})
	})
	return cache_defaultSucc__2204581824
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

var cache_toEnumWithDefaults__3941305703 gopurs_runtime.Value
var once_toEnumWithDefaults__3941305703 sync.Once
func Get_toEnumWithDefaults__3941305703() gopurs_runtime.Value {
	once_toEnumWithDefaults__3941305703.Do(func() {
		cache_toEnumWithDefaults__3941305703 = func() gopurs_runtime.Value {
bottom2_0_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Enum.Get_boundedEnumChar(), "Bounded0"), gopurs_runtime.Value{}), "bottom")
_ = bottom2_0_0
return gopurs_runtime.Func(func(low_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(high_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
v_4_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Enum.Get_boundedEnumChar(), "toEnum"), gopurs_runtime.Int(x_3.IntVal))
_ = v_4_1
var __t3 gopurs_runtime.Value
{
if (v_4_1.Type == 9 && v_4_1.IntVal == 930809136 && v_4_1.UnsafePtr != nil) {
__t3 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_4_1.UnsafePtr).V0
goto end_branch_3
} else {

}
}
{
if (v_4_1.Type == 9 && v_4_1.IntVal == 930809136 && v_4_1.UnsafePtr == nil) {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(Call_lessThan__4087042607(gopurs_runtime.Int(x_3.IntVal), gopurs_runtime.Int(gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Enum.Get_boundedEnumChar(), "fromEnum"), bottom2_0_0).IntVal))).IntVal) != (0) {
__t2 = low_1
goto end_branch_2
} else {

}
}
{
__t2 = high_2
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
})
})
})
}()
	})
	return cache_toEnumWithDefaults__3941305703
}

var cache_toEnumWithDefaults__3558602759 gopurs_runtime.Value
var once_toEnumWithDefaults__3558602759 sync.Once
func Get_toEnumWithDefaults__3558602759() gopurs_runtime.Value {
	once_toEnumWithDefaults__3558602759.Do(func() {
		cache_toEnumWithDefaults__3558602759 = gopurs_runtime.Func(func(dictBoundedEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toEnumWithDefaults__3558602759(gopurs_runtime.CoerceToStruct[pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value]](dictBoundedEnum_0_box))
})
	})
	return cache_toEnumWithDefaults__3558602759
}

var cache_foldl__2151204251 gopurs_runtime.Value
var once_foldl__2151204251 sync.Once
func Get_foldl__2151204251() gopurs_runtime.Value {
	once_foldl__2151204251.Do(func() {
		cache_foldl__2151204251 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__2151204251(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldl__2151204251
}

var cache_foldr__2151204251 gopurs_runtime.Value
var once_foldr__2151204251 sync.Once
func Get_foldr__2151204251() gopurs_runtime.Value {
	once_foldr__2151204251.Do(func() {
		cache_foldr__2151204251 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__2151204251(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldr__2151204251
}

var cache_foldr__3675782427 gopurs_runtime.Value
var once_foldr__3675782427 sync.Once
func Get_foldr__3675782427() gopurs_runtime.Value {
	once_foldr__3675782427.Do(func() {
		cache_foldr__3675782427 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__3675782427(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldr__3675782427
}

var cache_length__1958096179 gopurs_runtime.Value
var once_length__1958096179 sync.Once
func Get_length__1958096179() gopurs_runtime.Value {
	once_length__1958096179.Do(func() {
		cache_length__1958096179 = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_length__1958096179(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box))
})
	})
	return cache_length__1958096179
}

var cache_length__949294460 gopurs_runtime.Value
var once_length__949294460 sync.Once
func Get_length__949294460() gopurs_runtime.Value {
	once_length__949294460.Do(func() {
		cache_length__949294460 = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictSemiring_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_length__949294460(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](dictSemiring_1_box))
})
	})
	return cache_length__949294460
}

var cache_map__2278567252 gopurs_runtime.Value
var once_map__2278567252 sync.Once
func Get_map__2278567252() gopurs_runtime.Value {
	once_map__2278567252.Do(func() {
		cache_map__2278567252 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2278567252(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2278567252
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

var cache_conj__3676519832 gopurs_runtime.Value
var once_conj__3676519832 sync.Once
func Get_conj__3676519832() gopurs_runtime.Value {
	once_conj__3676519832.Do(func() {
		cache_conj__3676519832 = gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj")
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
		cache_disj__3676519832 = gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "disj")
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
		cache_not__3201284355 = gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "not")
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

var cache_un__3470773042 gopurs_runtime.Value
var once_un__3470773042 sync.Once
func Get_un__3470773042() gopurs_runtime.Value {
	once_un__3470773042.Do(func() {
		cache_un__3470773042 = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_un__3470773042(gopurs_runtime.CoerceToStruct[pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](_dollar__unused_0_box), v_1_box)
})
	})
	return cache_un__3470773042
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

var cache_foldMap1__3342855683 gopurs_runtime.Value
var once_foldMap1__3342855683 sync.Once
func Get_foldMap1__3342855683() gopurs_runtime.Value {
	once_foldMap1__3342855683.Do(func() {
		cache_foldMap1__3342855683 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMap1__3342855683(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldMap1__3342855683
}

var cache_semigroupLast__2108226578 gopurs_runtime.Value
var once_semigroupLast__2108226578 sync.Once
func Get_semigroupLast__2108226578() gopurs_runtime.Value {
	once_semigroupLast__2108226578.Do(func() {
		cache_semigroupLast__2108226578 = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
})
}))
	})
	return cache_semigroupLast__2108226578
}

var cache_add__560788792 gopurs_runtime.Value
var once_add__560788792 sync.Once
func Get_add__560788792() gopurs_runtime.Value {
	once_add__560788792.Do(func() {
		cache_add__560788792 = pkg_Data_Semiring.Get_intAdd()
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

var cache_one__1204848985 gopurs_runtime.Value
var once_one__1204848985 sync.Once
func Get_one__1204848985() gopurs_runtime.Value {
	once_one__1204848985.Do(func() {
		cache_one__1204848985 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_one__1204848985(dict_0_box)
})
	})
	return cache_one__1204848985
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

func Call_genUnicodeChar(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(Get_toEnumWithDefaults__3941305703(), gopurs_runtime.Str(gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "bottom").StrVal()), gopurs_runtime.Str(gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "top").StrVal())), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(0), gopurs_runtime.Int(65536)))
}

func Call_genDigitChar(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(Get_toEnumWithDefaults__3941305703(), gopurs_runtime.Str(gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "bottom").StrVal()), gopurs_runtime.Str(gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "top").StrVal())), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(48), gopurs_runtime.Int(57)))
}

func Call_genAsciiChar_prime(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(Get_toEnumWithDefaults__3941305703(), gopurs_runtime.Str(gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "bottom").StrVal()), gopurs_runtime.Str(gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "top").StrVal())), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(0), gopurs_runtime.Int(127)))
}

func Call_genAsciiChar(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(Get_toEnumWithDefaults__3941305703(), gopurs_runtime.Str(gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "bottom").StrVal()), gopurs_runtime.Str(gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "top").StrVal())), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(32), gopurs_runtime.Int(127)))
}

func Call_genAlphaUppercase(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(Get_toEnumWithDefaults__3941305703(), gopurs_runtime.Str(gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "bottom").StrVal()), gopurs_runtime.Str(gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "top").StrVal())), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(65), gopurs_runtime.Int(90)))
}

func Call_genAlphaLowercase(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(Get_toEnumWithDefaults__3941305703(), gopurs_runtime.Str(gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "bottom").StrVal()), gopurs_runtime.Str(gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "top").StrVal())), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(97), gopurs_runtime.Int(122)))
}

func Call_genAlpha(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply3(pkg_Control_Monad_Gen.Get_oneOf(), gopurs_runtime.Value{Type: 9, IntVal: 2254593219, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dictMonadGen_0))}, gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup_Foldable.Constructor_Foldable1[*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_foldable1NonEmpty()))}, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, Call_genAlphaLowercase(gopurs_runtime.Value{Type: 9, IntVal: 2254593219, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dictMonadGen_0))}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{Call_genAlphaUppercase(dictMonadGen_0)}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())})})
}

func Call_bind__1858449959(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__2601835655(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_chooseInt__1063828903(dict_0_loop *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_fromIndex__3111933544(dictFoldable1_0_loop *pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
Foldable0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](gopurs_runtime.Apply(dictFoldable1_0.V0, gopurs_runtime.Value{}))
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
var __t4 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 759514854 && v1_6.UnsafePtr != nil) {
var __t3 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*pkg_Control_Monad_Gen.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)}
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 759514854 && __t_tag_2.UnsafePtr == nil) {
__t3 = (*pkg_Control_Monad_Gen.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V0
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(Call_lessThanOrEq__4087042607(gopurs_runtime.Int(v_5), gopurs_runtime.Int(0))).IntVal) != (0) {
__t3 = (*pkg_Control_Monad_Gen.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V0
goto end_branch_3
} else {

}
}
{
v_5_loop = gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Int(v_5), gopurs_runtime.Int(1)).IntVal
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*pkg_Control_Monad_Gen.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)}
continue go__go_4_1_0
__t3 = gopurs_runtime.Value{}
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 759514854 && v1_6.UnsafePtr == nil) {
__t4 = gopurs_runtime.Apply3(dictFoldable1_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](pkg_Data_Semigroup_Last.Get_semigroupLast()))}, pkg_Data_Semigroup_Last.Get_Last(), xs_3)
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}
}()
})
})
return gopurs_runtime.Apply2(go__go_4_1_0, gopurs_runtime.Int(i_2.IntVal), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(Foldable0_1_0.V2, pkg_Control_Monad_Gen.Get_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*pkg_Control_Monad_Gen.Constructor_Cons[gopurs_runtime.Value])(nil))}, xs_3)))})
})
})
}

func Call_fromIndex__4031440680(dictFoldable1_0_loop *pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
Foldable0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](gopurs_runtime.Apply(dictFoldable1_0.V0, gopurs_runtime.Value{}))
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
var __t4 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 759514854 && v1_6.UnsafePtr != nil) {
var __t3 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*pkg_Control_Monad_Gen.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)}
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 759514854 && __t_tag_2.UnsafePtr == nil) {
__t3 = (*pkg_Control_Monad_Gen.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V0
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(Call_lessThanOrEq__4087042607(gopurs_runtime.Int(v_5), gopurs_runtime.Int(0))).IntVal) != (0) {
__t3 = (*pkg_Control_Monad_Gen.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V0
goto end_branch_3
} else {

}
}
{
v_5_loop = gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Int(v_5), gopurs_runtime.Int(1)).IntVal
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*pkg_Control_Monad_Gen.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)}
continue go__go_4_1_1
__t3 = gopurs_runtime.Value{}
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 759514854 && v1_6.UnsafePtr == nil) {
__t4 = gopurs_runtime.Apply3(dictFoldable1_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](pkg_Data_Semigroup_Last.Get_semigroupLast()))}, pkg_Data_Semigroup_Last.Get_Last(), xs_3)
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}
}()
})
})
return gopurs_runtime.Apply2(go__go_4_1_1, gopurs_runtime.Int(i_2.IntVal), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(Foldable0_1_0.V2, pkg_Control_Monad_Gen.Get_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*pkg_Control_Monad_Gen.Constructor_Cons[gopurs_runtime.Value])(nil))}, xs_3)))})
})
})
}

func Call_oneOf__2265861353(dictMonadGen_0_loop *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadGen_0 *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value] = dictMonadGen_0_loop
_ = dictMonadGen_0
Bind1_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonadGen_0.V0, gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_0
return gopurs_runtime.Func(func(dictFoldable1_2 gopurs_runtime.Value) gopurs_runtime.Value {
Foldable0_3_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_2, "Foldable0"), gopurs_runtime.Value{}))
_ = Foldable0_3_1
return gopurs_runtime.Func(func(xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_2 := &pkg_Data_Semiring.Constructor_Semiring[int64]{1, pkg_Data_Semiring.Get_intAdd(), pkg_Data_Semiring.Get_intMul(), 1, 0}
_ = __local_var_5_2
return gopurs_runtime.Apply2(Bind1_1_0.V1, gopurs_runtime.Apply2(dictMonadGen_0.V3, gopurs_runtime.Int(0), gopurs_runtime.Int(gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Int(gopurs_runtime.Apply3(Foldable0_3_1.V1, gopurs_runtime.Func(func(c_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_5_2.V0, __local_var_5_2.V2, c_6)
})
}), __local_var_5_2.V3, xs_4).IntVal), gopurs_runtime.Int(1)).IntVal)), gopurs_runtime.Func(func(n_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(pkg_Control_Monad_Gen.Get_fromIndex(), gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value]](dictFoldable1_2))}, gopurs_runtime.Int(n_5.IntVal), xs_4)
}))
})
})
}

func Call_oneOf__1433237231(dictMonadGen_0_loop *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadGen_0 *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value] = dictMonadGen_0_loop
_ = dictMonadGen_0
Bind1_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonadGen_0.V0, gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_0
Foldable0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup_Foldable.Constructor_Foldable1[*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_foldable1NonEmpty()).V0, gopurs_runtime.Value{}))
_ = Foldable0_2_1
return gopurs_runtime.Func(func(xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := &pkg_Data_Semiring.Constructor_Semiring[int64]{1, pkg_Data_Semiring.Get_intAdd(), pkg_Data_Semiring.Get_intMul(), 1, 0}
_ = __local_var_4_2
return gopurs_runtime.Apply2(Bind1_1_0.V1, gopurs_runtime.Apply2(dictMonadGen_0.V3, gopurs_runtime.Int(0), gopurs_runtime.Int(gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Int(gopurs_runtime.Apply3(Foldable0_2_1.V1, gopurs_runtime.Func(func(c_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_4_2.V0, __local_var_4_2.V2, c_5)
})
}), __local_var_4_2.V3, xs_3).IntVal), gopurs_runtime.Int(1)).IntVal)), gopurs_runtime.Func(func(n_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(pkg_Control_Monad_Gen.Get_fromIndex(), gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup_Foldable.Constructor_Foldable1[*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_foldable1NonEmpty()))}, gopurs_runtime.Int(n_4.IntVal), xs_3)
}))
})
}

func Call_bottom__338427193(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "bottom")
}

func Call_top__338427193(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "top")
}

func Call_genAlphaLowercase__4294897069(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(Get_toEnumWithDefaults__3941305703(), gopurs_runtime.Str(gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "bottom").StrVal()), gopurs_runtime.Str(gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "top").StrVal())), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(97), gopurs_runtime.Int(122)))
}

func Call_defaultPred__1581620096(toEnum_prime_0_loop gopurs_runtime.Value, fromEnum_prime_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var toEnum_prime_0 gopurs_runtime.Value = toEnum_prime_0_loop
_ = toEnum_prime_0
var fromEnum_prime_1 gopurs_runtime.Value = fromEnum_prime_1_loop
_ = fromEnum_prime_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(toEnum_prime_0, gopurs_runtime.Int(gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Int(gopurs_runtime.Apply(fromEnum_prime_1, a_2).IntVal), gopurs_runtime.Int(1)).IntVal)))
}

func Call_defaultPred__2204581824(toEnum_prime_0_loop gopurs_runtime.Value, fromEnum_prime_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var toEnum_prime_0 gopurs_runtime.Value = toEnum_prime_0_loop
_ = toEnum_prime_0
var fromEnum_prime_1 gopurs_runtime.Value = fromEnum_prime_1_loop
_ = fromEnum_prime_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(toEnum_prime_0, gopurs_runtime.Int(gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Int(gopurs_runtime.Apply(fromEnum_prime_1, a_2).IntVal), gopurs_runtime.Int(1)).IntVal)))
}

func Call_defaultSucc__1581620096(toEnum_prime_0_loop gopurs_runtime.Value, fromEnum_prime_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var toEnum_prime_0 gopurs_runtime.Value = toEnum_prime_0_loop
_ = toEnum_prime_0
var fromEnum_prime_1 gopurs_runtime.Value = fromEnum_prime_1_loop
_ = fromEnum_prime_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(toEnum_prime_0, gopurs_runtime.Int(gopurs_runtime.Apply2(Get_add__560788792(), gopurs_runtime.Int(gopurs_runtime.Apply(fromEnum_prime_1, a_2).IntVal), gopurs_runtime.Int(1)).IntVal)))
}

func Call_defaultSucc__2204581824(toEnum_prime_0_loop gopurs_runtime.Value, fromEnum_prime_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var toEnum_prime_0 gopurs_runtime.Value = toEnum_prime_0_loop
_ = toEnum_prime_0
var fromEnum_prime_1 gopurs_runtime.Value = fromEnum_prime_1_loop
_ = fromEnum_prime_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(toEnum_prime_0, gopurs_runtime.Int(gopurs_runtime.Apply2(Get_add__560788792(), gopurs_runtime.Int(gopurs_runtime.Apply(fromEnum_prime_1, a_2).IntVal), gopurs_runtime.Int(1)).IntVal)))
}

func Call_fromEnum__1637084359(dict_0_loop *pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_toEnum__3317293286(dict_0_loop *pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V4
}

func Call_toEnumWithDefaults__3558602759(dictBoundedEnum_0_loop *pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBoundedEnum_0 *pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value] = dictBoundedEnum_0_loop
_ = dictBoundedEnum_0
bottom2_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictBoundedEnum_0.V0, gopurs_runtime.Value{}), "bottom")
_ = bottom2_1_0
return gopurs_runtime.Func(func(low_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(high_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
v_5_1 := gopurs_runtime.Apply(dictBoundedEnum_0.V4, gopurs_runtime.Int(x_4.IntVal))
_ = v_5_1
var __t3 gopurs_runtime.Value
{
if (v_5_1.Type == 9 && v_5_1.IntVal == 930809136 && v_5_1.UnsafePtr != nil) {
__t3 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_5_1.UnsafePtr).V0
goto end_branch_3
} else {

}
}
{
if (v_5_1.Type == 9 && v_5_1.IntVal == 930809136 && v_5_1.UnsafePtr == nil) {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(Call_lessThan__4087042607(gopurs_runtime.Int(x_4.IntVal), gopurs_runtime.Int(gopurs_runtime.Apply(dictBoundedEnum_0.V3, bottom2_1_0).IntVal))).IntVal) != (0) {
__t2 = low_2
goto end_branch_2
} else {

}
}
{
__t2 = high_3
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
})
})
})
}

func Call_foldl__2151204251(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldr__2151204251(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_foldr__3675782427(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_length__1958096179(dictFoldable_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
return gopurs_runtime.Apply2(dictFoldable_0.V1, gopurs_runtime.Func(func(c_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((1) + (c_1.IntVal))
})
}), gopurs_runtime.Int(0))
}

func Call_length__949294460(dictFoldable_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value], dictSemiring_1_loop *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictSemiring_1 *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value] = dictSemiring_1_loop
_ = dictSemiring_1
return gopurs_runtime.Apply2(dictFoldable_0.V1, gopurs_runtime.Func(func(c_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictSemiring_1.V0, dictSemiring_1.V2, c_2)
})
}), dictSemiring_1.V3)
}

func Call_map__2278567252(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__2199395572(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_conj__3472268504(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_disj__3472268504(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_not__1505204753(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V4
}

func Call_un__3470773042(_dollar__unused_0_loop *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value], v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar__unused_0_loop
_ = _dollar__unused_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_compare__821463600(dict_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
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

func Call_sub__3675938712(dict_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldMap1__3342855683(dict_0_loop *pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_add__1614463960(dict_0_loop *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_one__1204848985(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "one")
}

func Call_zero__1204848985(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "zero")
}


