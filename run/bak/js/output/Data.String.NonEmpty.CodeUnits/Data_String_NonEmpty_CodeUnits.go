package Data_String_NonEmpty_CodeUnits

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_String_CodeUnits "gopurs/output/Data.String.CodeUnits"
	pkg_Data_String_Unsafe "gopurs/output/Data.String.Unsafe"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Control_Category "gopurs/output/Control.Category"
)

var snoc gopurs_runtime.Value
var once_snoc sync.Once
func Get_snoc() gopurs_runtime.Value {
	once_snoc.Do(func() {
		snoc = gopurs_runtime.Func2(func(c_0 gopurs_runtime.Value, s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(s_1.StrVal + gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_singleton(), c_0).StrVal)
})
	})
	return snoc
}

var singleton gopurs_runtime.Value
var once_singleton sync.Once
func Get_singleton() gopurs_runtime.Value {
	once_singleton.Do(func() {
		singleton = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_singleton(), x_0)
})
	})
	return singleton
}

var takeWhile gopurs_runtime.Value
var once_takeWhile sync.Once
func Get_takeWhile() gopurs_runtime.Value {
	once_takeWhile.Do(func() {
		takeWhile = gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_take(), gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_countPrefix(), f_0, x_1), x_1)
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if __local_var_2_0.StrVal == "" {
__t1 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor1("Just", __local_var_2_0)
}
end_branch_1:
return __t1
})
	})
	return takeWhile
}

var lastIndexOf_prime gopurs_runtime.Value
var once_lastIndexOf_prime sync.Once
func Get_lastIndexOf_prime() gopurs_runtime.Value {
	once_lastIndexOf_prime.Do(func() {
		lastIndexOf_prime = gopurs_runtime.Func(func(pat_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_lastIndexOf_prime(), pat_0)
})
	})
	return lastIndexOf_prime
}

var lastIndexOf gopurs_runtime.Value
var once_lastIndexOf sync.Once
func Get_lastIndexOf() gopurs_runtime.Value {
	once_lastIndexOf.Do(func() {
		lastIndexOf = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_lastIndexOf(), x_0)
})
	})
	return lastIndexOf
}

var indexOf_prime gopurs_runtime.Value
var once_indexOf_prime sync.Once
func Get_indexOf_prime() gopurs_runtime.Value {
	once_indexOf_prime.Do(func() {
		indexOf_prime = gopurs_runtime.Func(func(pat_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_indexOf_prime(), pat_0)
})
	})
	return indexOf_prime
}

var indexOf gopurs_runtime.Value
var once_indexOf sync.Once
func Get_indexOf() gopurs_runtime.Value {
	once_indexOf.Do(func() {
		indexOf = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_indexOf(), x_0)
})
	})
	return indexOf
}

var length gopurs_runtime.Value
var once_length sync.Once
func Get_length() gopurs_runtime.Value {
	once_length.Do(func() {
		length = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_length(), x_0)
})
	})
	return length
}

var splitAt gopurs_runtime.Value
var once_splitAt sync.Once
func Get_splitAt() gopurs_runtime.Value {
	once_splitAt.Do(func() {
		splitAt = gopurs_runtime.Func2(func(i_0 gopurs_runtime.Value, nes_1 gopurs_runtime.Value) gopurs_runtime.Value {
v_2_0 := gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_splitAt(), i_0, nes_1)
_ = v_2_0
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet(v_2_0, "before").StrVal == "" {
__t1 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor1("Just", gopurs_runtime.RecordGet(v_2_0, "before"))
}
end_branch_1:
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet(v_2_0, "after").StrVal == "" {
__t2 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Constructor1("Just", gopurs_runtime.RecordGet(v_2_0, "after"))
}
end_branch_2:
return gopurs_runtime.RecordDict2("before", "after", __t1, __t2)
})
	})
	return splitAt
}

var take gopurs_runtime.Value
var once_take sync.Once
func Get_take() gopurs_runtime.Value {
	once_take.Do(func() {
		take = gopurs_runtime.Func2(func(i_0 gopurs_runtime.Value, nes_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if i_0.IntVal < 1 {
__t0 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_take(), i_0, nes_1))
}
end_branch_0:
return __t0
})
	})
	return take
}

var takeRight gopurs_runtime.Value
var once_takeRight sync.Once
func Get_takeRight() gopurs_runtime.Value {
	once_takeRight.Do(func() {
		takeRight = gopurs_runtime.Func2(func(i_0 gopurs_runtime.Value, nes_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if i_0.IntVal < 1 {
__t0 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_drop(), gopurs_runtime.Int(gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_length(), nes_1).IntVal - i_0.IntVal), nes_1))
}
end_branch_0:
return __t0
})
	})
	return takeRight
}

var toChar gopurs_runtime.Value
var once_toChar sync.Once
func Get_toChar() gopurs_runtime.Value {
	once_toChar.Do(func() {
		toChar = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_toChar(), x_0)
})
	})
	return toChar
}

var toCharArray gopurs_runtime.Value
var once_toCharArray sync.Once
func Get_toCharArray() gopurs_runtime.Value {
	once_toCharArray.Do(func() {
		toCharArray = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_toCharArray(), x_0)
})
	})
	return toCharArray
}

var toNonEmptyCharArray gopurs_runtime.Value
var once_toNonEmptyCharArray sync.Once
func Get_toNonEmptyCharArray() gopurs_runtime.Value {
	once_toNonEmptyCharArray.Do(func() {
		toNonEmptyCharArray = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_toCharArray(), x_0)
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Int(int64(len(__local_var_1_0.PtrVal.([]gopurs_runtime.Value)))).IntVal > 0 {
__t1 = __local_var_1_0
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
})
	})
	return toNonEmptyCharArray
}

var uncons gopurs_runtime.Value
var once_uncons sync.Once
func Get_uncons() gopurs_runtime.Value {
	once_uncons.Do(func() {
		uncons = gopurs_runtime.Func(func(nes_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_drop(), gopurs_runtime.Int(1), nes_0)
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if __local_var_1_0.StrVal == "" {
__t1 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor1("Just", __local_var_1_0)
}
end_branch_1:
return gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Apply2(pkg_Data_String_Unsafe.Get_charAt(), gopurs_runtime.Int(0), nes_0), __t1)
})
	})
	return uncons
}

var fromFoldable1 gopurs_runtime.Value
var once_fromFoldable1 sync.Once
func Get_fromFoldable1() gopurs_runtime.Value {
	once_fromFoldable1.Do(func() {
		fromFoldable1 = gopurs_runtime.Func(func(dictFoldable1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable1_0, "foldMap1"), pkg_Data_Semigroup.Get_semigroupString(), gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
})
	})
	return fromFoldable1
}

var fromCharArray gopurs_runtime.Value
var once_fromCharArray sync.Once
func Get_fromCharArray() gopurs_runtime.Value {
	once_fromCharArray.Do(func() {
		fromCharArray = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Int(int64(len(v_0.PtrVal.([]gopurs_runtime.Value)))).IntVal == 0 {
__t0 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_fromCharArray(), v_0))
}
end_branch_0:
return __t0
})
	})
	return fromCharArray
}

var fromNonEmptyCharArray gopurs_runtime.Value
var once_fromNonEmptyCharArray sync.Once
func Get_fromNonEmptyCharArray() gopurs_runtime.Value {
	once_fromNonEmptyCharArray.Do(func() {
		fromNonEmptyCharArray = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Int(int64(len(x_0.PtrVal.([]gopurs_runtime.Value)))).IntVal == 0 {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_fromCharArray(), x_0)
}
end_branch_0:
return __t0
})
	})
	return fromNonEmptyCharArray
}

var dropWhile gopurs_runtime.Value
var once_dropWhile sync.Once
func Get_dropWhile() gopurs_runtime.Value {
	once_dropWhile.Do(func() {
		dropWhile = gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_drop(), gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_countPrefix(), f_0, x_1), x_1)
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if __local_var_2_0.StrVal == "" {
__t1 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor1("Just", __local_var_2_0)
}
end_branch_1:
return __t1
})
	})
	return dropWhile
}

var dropRight gopurs_runtime.Value
var once_dropRight sync.Once
func Get_dropRight() gopurs_runtime.Value {
	once_dropRight.Do(func() {
		dropRight = gopurs_runtime.Func2(func(i_0 gopurs_runtime.Value, nes_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if i_0.IntVal >= gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_length(), nes_1).IntVal {
__t0 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_take(), gopurs_runtime.Int(gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_length(), nes_1).IntVal - i_0.IntVal), nes_1))
}
end_branch_0:
return __t0
})
	})
	return dropRight
}

var drop gopurs_runtime.Value
var once_drop sync.Once
func Get_drop() gopurs_runtime.Value {
	once_drop.Do(func() {
		drop = gopurs_runtime.Func2(func(i_0 gopurs_runtime.Value, nes_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if i_0.IntVal >= gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_length(), nes_1).IntVal {
__t0 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_drop(), i_0, nes_1))
}
end_branch_0:
return __t0
})
	})
	return drop
}

var countPrefix gopurs_runtime.Value
var once_countPrefix sync.Once
func Get_countPrefix() gopurs_runtime.Value {
	once_countPrefix.Do(func() {
		countPrefix = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_countPrefix(), x_0)
})
	})
	return countPrefix
}

var cons gopurs_runtime.Value
var once_cons sync.Once
func Get_cons() gopurs_runtime.Value {
	once_cons.Do(func() {
		cons = gopurs_runtime.Func2(func(c_0 gopurs_runtime.Value, s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_singleton(), c_0).StrVal + s_1.StrVal)
})
	})
	return cons
}

var charAt gopurs_runtime.Value
var once_charAt sync.Once
func Get_charAt() gopurs_runtime.Value {
	once_charAt.Do(func() {
		charAt = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_charAt(), x_0)
})
	})
	return charAt
}




