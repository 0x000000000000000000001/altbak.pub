package Data_Ordering

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var LT gopurs_runtime.Value
var once_LT sync.Once
func Get_LT() gopurs_runtime.Value {
	once_LT.Do(func() {
		LT = gopurs_runtime.Constructor0("LT")
	})
	return LT
}

var GT gopurs_runtime.Value
var once_GT sync.Once
func Get_GT() gopurs_runtime.Value {
	once_GT.Do(func() {
		GT = gopurs_runtime.Constructor0("GT")
	})
	return GT
}

var EQ gopurs_runtime.Value
var once_EQ sync.Once
func Get_EQ() gopurs_runtime.Value {
	once_EQ.Do(func() {
		EQ = gopurs_runtime.Constructor0("EQ")
	})
	return EQ
}

var showOrdering gopurs_runtime.Value
var once_showOrdering sync.Once
func Get_showOrdering() gopurs_runtime.Value {
	once_showOrdering.Do(func() {
		showOrdering = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_0.StrVal == "LT").IntVal != 0 {
__t0 = gopurs_runtime.Str("LT")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_0.StrVal == "GT").IntVal != 0 {
__t0 = gopurs_runtime.Str("GT")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_0.StrVal == "EQ").IntVal != 0 {
__t0 = gopurs_runtime.Str("EQ")
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
	return showOrdering
}

var semigroupOrdering gopurs_runtime.Value
var once_semigroupOrdering sync.Once
func Get_semigroupOrdering() gopurs_runtime.Value {
	once_semigroupOrdering.Do(func() {
		semigroupOrdering = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_0.StrVal == "LT").IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("LT")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_0.StrVal == "GT").IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("GT")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_0.StrVal == "EQ").IntVal != 0 {
__t0 = v1_1
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
	return semigroupOrdering
}

var invert gopurs_runtime.Value
var once_invert sync.Once
func Get_invert() gopurs_runtime.Value {
	once_invert.Do(func() {
		invert = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_0.StrVal == "GT").IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("LT")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_0.StrVal == "EQ").IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("EQ")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_0.StrVal == "LT").IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("GT")
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})
	})
	return invert
}

var eqOrdering gopurs_runtime.Value
var once_eqOrdering sync.Once
func Get_eqOrdering() gopurs_runtime.Value {
	once_eqOrdering.Do(func() {
		eqOrdering = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_0.StrVal == "LT").IntVal != 0 {
__t0 = gopurs_runtime.Bool(v1_1.StrVal == "LT")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_0.StrVal == "GT").IntVal != 0 {
__t0 = gopurs_runtime.Bool(v1_1.StrVal == "GT")
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Bool(gopurs_runtime.Bool(v_0.StrVal == "EQ").IntVal != 0 && gopurs_runtime.Bool(v1_1.StrVal == "EQ").IntVal != 0)
}
end_branch_0:
return __t0
}))
	})
	return eqOrdering
}




