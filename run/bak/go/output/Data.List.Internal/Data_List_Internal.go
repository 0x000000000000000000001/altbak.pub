package Data_List_Internal

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var Leaf gopurs_runtime.Value
var once_Leaf sync.Once
func Get_Leaf() gopurs_runtime.Value {
	once_Leaf.Do(func() {
		Leaf = gopurs_runtime.Constructor0("Leaf")
	})
	return Leaf
}

var Two gopurs_runtime.Value
var once_Two sync.Once
func Get_Two() gopurs_runtime.Value {
	once_Two.Do(func() {
		Two = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor3("Two", value0, value1, value2)
})
})
})
	})
	return Two
}

var Three gopurs_runtime.Value
var once_Three sync.Once
func Get_Three() gopurs_runtime.Value {
	once_Three.Do(func() {
		Three = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor5("Three", value0, value1, value2, value3, value4)
})
})
})
})
})
	})
	return Three
}

var TwoLeft gopurs_runtime.Value
var once_TwoLeft sync.Once
func Get_TwoLeft() gopurs_runtime.Value {
	once_TwoLeft.Do(func() {
		TwoLeft = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("TwoLeft", value0, value1)
})
})
	})
	return TwoLeft
}

var TwoRight gopurs_runtime.Value
var once_TwoRight sync.Once
func Get_TwoRight() gopurs_runtime.Value {
	once_TwoRight.Do(func() {
		TwoRight = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("TwoRight", value0, value1)
})
})
	})
	return TwoRight
}

var ThreeLeft gopurs_runtime.Value
var once_ThreeLeft sync.Once
func Get_ThreeLeft() gopurs_runtime.Value {
	once_ThreeLeft.Do(func() {
		ThreeLeft = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor4("ThreeLeft", value0, value1, value2, value3)
})
})
})
})
	})
	return ThreeLeft
}

var ThreeMiddle gopurs_runtime.Value
var once_ThreeMiddle sync.Once
func Get_ThreeMiddle() gopurs_runtime.Value {
	once_ThreeMiddle.Do(func() {
		ThreeMiddle = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor4("ThreeMiddle", value0, value1, value2, value3)
})
})
})
})
	})
	return ThreeMiddle
}

var ThreeRight gopurs_runtime.Value
var once_ThreeRight sync.Once
func Get_ThreeRight() gopurs_runtime.Value {
	once_ThreeRight.Do(func() {
		ThreeRight = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor4("ThreeRight", value0, value1, value2, value3)
})
})
})
})
	})
	return ThreeRight
}

var KickUp gopurs_runtime.Value
var once_KickUp sync.Once
func Get_KickUp() gopurs_runtime.Value {
	once_KickUp.Do(func() {
		KickUp = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor3("KickUp", value0, value1, value2)
})
})
})
	})
	return KickUp
}

var fromZipper gopurs_runtime.Value
var once_fromZipper sync.Once
func Get_fromZipper() gopurs_runtime.Value {
	once_fromZipper.Do(func() {
		fromZipper = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
fromZipper:
for {
if false { continue fromZipper }
var v_0 = v_0_loop
_ = v_0
var v1_1 = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.StrVal == "Nil")).IntVal != 0 {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.StrVal == "Cons")).IntVal != 0 {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.ConstructorGet(v_0, 0).StrVal == "TwoLeft")).IntVal != 0 {
__t1 = gopurs_runtime.Apply2(Get_fromZipper(), gopurs_runtime.ConstructorGet(v_0, 1), gopurs_runtime.Constructor3("Two", v1_1, gopurs_runtime.ConstructorGet(gopurs_runtime.ConstructorGet(v_0, 0), 0), gopurs_runtime.ConstructorGet(gopurs_runtime.ConstructorGet(v_0, 0), 1)))
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.ConstructorGet(v_0, 0).StrVal == "TwoRight")).IntVal != 0 {
__t1 = gopurs_runtime.Apply2(Get_fromZipper(), gopurs_runtime.ConstructorGet(v_0, 1), gopurs_runtime.Constructor3("Two", gopurs_runtime.ConstructorGet(gopurs_runtime.ConstructorGet(v_0, 0), 0), gopurs_runtime.ConstructorGet(gopurs_runtime.ConstructorGet(v_0, 0), 1), v1_1))
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.ConstructorGet(v_0, 0).StrVal == "ThreeLeft")).IntVal != 0 {
__t1 = gopurs_runtime.Apply2(Get_fromZipper(), gopurs_runtime.ConstructorGet(v_0, 1), gopurs_runtime.Constructor5("Three", v1_1, gopurs_runtime.ConstructorGet(gopurs_runtime.ConstructorGet(v_0, 0), 0), gopurs_runtime.ConstructorGet(gopurs_runtime.ConstructorGet(v_0, 0), 1), gopurs_runtime.ConstructorGet(gopurs_runtime.ConstructorGet(v_0, 0), 2), gopurs_runtime.ConstructorGet(gopurs_runtime.ConstructorGet(v_0, 0), 3)))
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.ConstructorGet(v_0, 0).StrVal == "ThreeMiddle")).IntVal != 0 {
__t1 = gopurs_runtime.Apply2(Get_fromZipper(), gopurs_runtime.ConstructorGet(v_0, 1), gopurs_runtime.Constructor5("Three", gopurs_runtime.ConstructorGet(gopurs_runtime.ConstructorGet(v_0, 0), 0), gopurs_runtime.ConstructorGet(gopurs_runtime.ConstructorGet(v_0, 0), 1), v1_1, gopurs_runtime.ConstructorGet(gopurs_runtime.ConstructorGet(v_0, 0), 2), gopurs_runtime.ConstructorGet(gopurs_runtime.ConstructorGet(v_0, 0), 3)))
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.ConstructorGet(v_0, 0).StrVal == "ThreeRight")).IntVal != 0 {
__t1 = gopurs_runtime.Apply2(Get_fromZipper(), gopurs_runtime.ConstructorGet(v_0, 1), gopurs_runtime.Constructor5("Three", gopurs_runtime.ConstructorGet(gopurs_runtime.ConstructorGet(v_0, 0), 0), gopurs_runtime.ConstructorGet(gopurs_runtime.ConstructorGet(v_0, 0), 1), gopurs_runtime.ConstructorGet(gopurs_runtime.ConstructorGet(v_0, 0), 2), gopurs_runtime.ConstructorGet(gopurs_runtime.ConstructorGet(v_0, 0), 3), v1_1))
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
__t0 = __t1
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
}()
})
})
	})
	return fromZipper
}

var insertAndLookupBy gopurs_runtime.Value
var once_insertAndLookupBy sync.Once
func Get_insertAndLookupBy() gopurs_runtime.Value {
	once_insertAndLookupBy.Do(func() {
		insertAndLookupBy = gopurs_runtime.Func3(func(comp_0 gopurs_runtime.Value, k_1 gopurs_runtime.Value, orig_2 gopurs_runtime.Value) gopurs_runtime.Value {
var up_3_0 gopurs_runtime.Value
up_3_0 = gopurs_runtime.Func(func(v_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
up_3_0:
for {
if false { continue up_3_0 }
var v_4 = v_4_loop
_ = v_4
var v1_5 = v1_5_loop
_ = v1_5
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_4.StrVal == "Nil")).IntVal != 0 {
__t1 = gopurs_runtime.Constructor3("Two", gopurs_runtime.ConstructorGet(v1_5, 0), gopurs_runtime.ConstructorGet(v1_5, 1), gopurs_runtime.ConstructorGet(v1_5, 2))
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_4.StrVal == "Cons")).IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.ConstructorGet(v_4, 0).StrVal == "TwoLeft")).IntVal != 0 {
__t2 = gopurs_runtime.Apply2(Get_fromZipper(), gopurs_runtime.ConstructorGet(v_4, 1), gopurs_runtime.Constructor5("Three", gopurs_runtime.ConstructorGet(v1_5, 0), gopurs_runtime.ConstructorGet(v1_5, 1), gopurs_runtime.ConstructorGet(v1_5, 2), gopurs_runtime.ConstructorGet(gopurs_runtime.ConstructorGet(v_4, 0), 0), gopurs_runtime.ConstructorGet(gopurs_runtime.ConstructorGet(v_4, 0), 1)))
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.ConstructorGet(v_4, 0).StrVal == "TwoRight")).IntVal != 0 {
__t2 = gopurs_runtime.Apply2(Get_fromZipper(), gopurs_runtime.ConstructorGet(v_4, 1), gopurs_runtime.Constructor5("Three", gopurs_runtime.ConstructorGet(gopurs_runtime.ConstructorGet(v_4, 0), 0), gopurs_runtime.ConstructorGet(gopurs_runtime.ConstructorGet(v_4, 0), 1), gopurs_runtime.ConstructorGet(v1_5, 0), gopurs_runtime.ConstructorGet(v1_5, 1), gopurs_runtime.ConstructorGet(v1_5, 2)))
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.ConstructorGet(v_4, 0).StrVal == "ThreeLeft")).IntVal != 0 {
v_4_loop = gopurs_runtime.ConstructorGet(v_4, 1)
v1_5_loop = gopurs_runtime.Constructor3("KickUp", gopurs_runtime.Constructor3("Two", gopurs_runtime.ConstructorGet(v1_5, 0), gopurs_runtime.ConstructorGet(v1_5, 1), gopurs_runtime.ConstructorGet(v1_5, 2)), gopurs_runtime.ConstructorGet(gopurs_runtime.ConstructorGet(v_4, 0), 0), gopurs_runtime.Constructor3("Two", gopurs_runtime.ConstructorGet(gopurs_runtime.ConstructorGet(v_4, 0), 1), gopurs_runtime.ConstructorGet(gopurs_runtime.ConstructorGet(v_4, 0), 2), gopurs_runtime.ConstructorGet(gopurs_runtime.ConstructorGet(v_4, 0), 3)))
continue up_3_0
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.ConstructorGet(v_4, 0).StrVal == "ThreeMiddle")).IntVal != 0 {
v_4_loop = gopurs_runtime.ConstructorGet(v_4, 1)
v1_5_loop = gopurs_runtime.Constructor3("KickUp", gopurs_runtime.Constructor3("Two", gopurs_runtime.ConstructorGet(gopurs_runtime.ConstructorGet(v_4, 0), 0), gopurs_runtime.ConstructorGet(gopurs_runtime.ConstructorGet(v_4, 0), 1), gopurs_runtime.ConstructorGet(v1_5, 0)), gopurs_runtime.ConstructorGet(v1_5, 1), gopurs_runtime.Constructor3("Two", gopurs_runtime.ConstructorGet(v1_5, 2), gopurs_runtime.ConstructorGet(gopurs_runtime.ConstructorGet(v_4, 0), 2), gopurs_runtime.ConstructorGet(gopurs_runtime.ConstructorGet(v_4, 0), 3)))
continue up_3_0
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.ConstructorGet(v_4, 0).StrVal == "ThreeRight")).IntVal != 0 {
v_4_loop = gopurs_runtime.ConstructorGet(v_4, 1)
v1_5_loop = gopurs_runtime.Constructor3("KickUp", gopurs_runtime.Constructor3("Two", gopurs_runtime.ConstructorGet(gopurs_runtime.ConstructorGet(v_4, 0), 0), gopurs_runtime.ConstructorGet(gopurs_runtime.ConstructorGet(v_4, 0), 1), gopurs_runtime.ConstructorGet(gopurs_runtime.ConstructorGet(v_4, 0), 2)), gopurs_runtime.ConstructorGet(gopurs_runtime.ConstructorGet(v_4, 0), 3), gopurs_runtime.Constructor3("Two", gopurs_runtime.ConstructorGet(v1_5, 0), gopurs_runtime.ConstructorGet(v1_5, 1), gopurs_runtime.ConstructorGet(v1_5, 2)))
continue up_3_0
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t1 = __t2
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
}()
})
})
var down_4_3 gopurs_runtime.Value
down_4_3 = gopurs_runtime.Func(func(v_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
down_4_3:
for {
if false { continue down_4_3 }
var v_5 = v_5_loop
_ = v_5
var v1_6 = v1_6_loop
_ = v1_6
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v1_6.StrVal == "Leaf")).IntVal != 0 {
__t4 = gopurs_runtime.RecordDict2("found", "result", gopurs_runtime.Bool(false), gopurs_runtime.Apply2(up_3_0, v_5, gopurs_runtime.Constructor3("KickUp", gopurs_runtime.Constructor0("Leaf"), k_1, gopurs_runtime.Constructor0("Leaf"))))
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(v1_6.StrVal == "Two")).IntVal != 0 {
v2_7_5 := gopurs_runtime.Apply2(comp_0, k_1, gopurs_runtime.ConstructorGet(v1_6, 1))
_ = v2_7_5
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v2_7_5.StrVal == "EQ")).IntVal != 0 {
__t6 = gopurs_runtime.RecordDict2("found", "result", gopurs_runtime.Bool(true), orig_2)
goto end_branch_6
} else {

}
}
{
if (gopurs_runtime.Bool(v2_7_5.StrVal == "LT")).IntVal != 0 {
v_5_loop = gopurs_runtime.Constructor2("Cons", gopurs_runtime.Constructor2("TwoLeft", gopurs_runtime.ConstructorGet(v1_6, 1), gopurs_runtime.ConstructorGet(v1_6, 2)), v_5)
v1_6_loop = gopurs_runtime.ConstructorGet(v1_6, 0)
continue down_4_3
__t6 = gopurs_runtime.Value{}
goto end_branch_6
} else {

}
}
{
v_5_loop = gopurs_runtime.Constructor2("Cons", gopurs_runtime.Constructor2("TwoRight", gopurs_runtime.ConstructorGet(v1_6, 0), gopurs_runtime.ConstructorGet(v1_6, 1)), v_5)
v1_6_loop = gopurs_runtime.ConstructorGet(v1_6, 2)
continue down_4_3
__t6 = gopurs_runtime.Value{}
}
end_branch_6:
__t4 = __t6
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(v1_6.StrVal == "Three")).IntVal != 0 {
v2_7_7 := gopurs_runtime.Apply2(comp_0, k_1, gopurs_runtime.ConstructorGet(v1_6, 1))
_ = v2_7_7
var __t10 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v2_7_7.StrVal == "EQ")).IntVal != 0 {
__t10 = gopurs_runtime.RecordDict2("found", "result", gopurs_runtime.Bool(true), orig_2)
goto end_branch_10
} else {

}
}
{
v3_8_8 := gopurs_runtime.Apply2(comp_0, k_1, gopurs_runtime.ConstructorGet(v1_6, 3))
_ = v3_8_8
var __t9 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v3_8_8.StrVal == "EQ")).IntVal != 0 {
__t9 = gopurs_runtime.RecordDict2("found", "result", gopurs_runtime.Bool(true), orig_2)
goto end_branch_9
} else {

}
}
{
if (gopurs_runtime.Bool(v2_7_7.StrVal == "LT")).IntVal != 0 {
v_5_loop = gopurs_runtime.Constructor2("Cons", gopurs_runtime.Constructor4("ThreeLeft", gopurs_runtime.ConstructorGet(v1_6, 1), gopurs_runtime.ConstructorGet(v1_6, 2), gopurs_runtime.ConstructorGet(v1_6, 3), gopurs_runtime.ConstructorGet(v1_6, 4)), v_5)
v1_6_loop = gopurs_runtime.ConstructorGet(v1_6, 0)
continue down_4_3
__t9 = gopurs_runtime.Value{}
goto end_branch_9
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(v2_7_7.StrVal == "GT").IntVal != 0 && gopurs_runtime.Bool(v3_8_8.StrVal == "LT").IntVal != 0)).IntVal != 0 {
v_5_loop = gopurs_runtime.Constructor2("Cons", gopurs_runtime.Constructor4("ThreeMiddle", gopurs_runtime.ConstructorGet(v1_6, 0), gopurs_runtime.ConstructorGet(v1_6, 1), gopurs_runtime.ConstructorGet(v1_6, 3), gopurs_runtime.ConstructorGet(v1_6, 4)), v_5)
v1_6_loop = gopurs_runtime.ConstructorGet(v1_6, 2)
continue down_4_3
__t9 = gopurs_runtime.Value{}
goto end_branch_9
} else {

}
}
{
v_5_loop = gopurs_runtime.Constructor2("Cons", gopurs_runtime.Constructor4("ThreeRight", gopurs_runtime.ConstructorGet(v1_6, 0), gopurs_runtime.ConstructorGet(v1_6, 1), gopurs_runtime.ConstructorGet(v1_6, 2), gopurs_runtime.ConstructorGet(v1_6, 3)), v_5)
v1_6_loop = gopurs_runtime.ConstructorGet(v1_6, 4)
continue down_4_3
__t9 = gopurs_runtime.Value{}
}
end_branch_9:
__t10 = __t9
}
end_branch_10:
__t4 = __t10
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
return gopurs_runtime.Apply2(down_4_3, gopurs_runtime.Constructor0("Nil"), orig_2)
})
	})
	return insertAndLookupBy
}

var emptySet gopurs_runtime.Value
var once_emptySet sync.Once
func Get_emptySet() gopurs_runtime.Value {
	once_emptySet.Do(func() {
		emptySet = gopurs_runtime.Constructor0("Leaf")
	})
	return emptySet
}


