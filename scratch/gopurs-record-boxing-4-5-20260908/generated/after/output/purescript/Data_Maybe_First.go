package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Maybe_First_First gopurs_runtime.Value
var once_Data_Maybe_First_First sync.Once
func Get_Data_Maybe_First_First() gopurs_runtime.Value {
	once_Data_Maybe_First_First.Do(func() {
		cache_Data_Maybe_First_First = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_First_First(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](x_0_box))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Maybe_First_First
}

var cache_Data_Maybe_First_showFirst gopurs_runtime.Value
var once_Data_Maybe_First_showFirst sync.Once
func Get_Data_Maybe_First_showFirst() gopurs_runtime.Value {
	once_Data_Maybe_First_showFirst.Do(func() {
		cache_Data_Maybe_First_showFirst = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_First_showFirst(dictShow_0_box)
})
	})
	return cache_Data_Maybe_First_showFirst
}

var cache_Data_Maybe_First_semigroupFirst gopurs_runtime.Value
var once_Data_Maybe_First_semigroupFirst sync.Once
func Get_Data_Maybe_First_semigroupFirst() gopurs_runtime.Value {
	once_Data_Maybe_First_semigroupFirst.Do(func() {
		cache_Data_Maybe_First_semigroupFirst = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_First_125234255_4179793454((&Constructor_Data_Semigroup_Semigroup[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_0)
if (__t_tag_0 != nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_0)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_1)
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
})})))}
	})
	return cache_Data_Maybe_First_semigroupFirst
}

var cache_Data_Maybe_First_ordFirst gopurs_runtime.Value
var once_Data_Maybe_First_ordFirst sync.Once
func Get_Data_Maybe_First_ordFirst() gopurs_runtime.Value {
	once_Data_Maybe_First_ordFirst.Do(func() {
		cache_Data_Maybe_First_ordFirst = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_First_ordFirst(dictOrd_0_box)
})
	})
	return cache_Data_Maybe_First_ordFirst
}

var cache_Data_Maybe_First_ord1First gopurs_runtime.Value
var once_Data_Maybe_First_ord1First sync.Once
func Get_Data_Maybe_First_ord1First() gopurs_runtime.Value {
	once_Data_Maybe_First_ord1First.Do(func() {
		cache_Data_Maybe_First_ord1First = gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_First_170825214_3985601471(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord1[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Maybe_ord1Maybe())))}
	})
	return cache_Data_Maybe_First_ord1First
}

var cache_Data_Maybe_First_newtypeFirst gopurs_runtime.Value
var once_Data_Maybe_First_newtypeFirst sync.Once
func Get_Data_Maybe_First_newtypeFirst() gopurs_runtime.Value {
	once_Data_Maybe_First_newtypeFirst.Do(func() {
		cache_Data_Maybe_First_newtypeFirst = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_First_1892488872_385277032((&Constructor_Data_Newtype_Newtype[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})))}
	})
	return cache_Data_Maybe_First_newtypeFirst
}

var cache_Data_Maybe_First_monoidFirst gopurs_runtime.Value
var once_Data_Maybe_First_monoidFirst sync.Once
func Get_Data_Maybe_First_monoidFirst() gopurs_runtime.Value {
	once_Data_Maybe_First_monoidFirst.Do(func() {
		cache_Data_Maybe_First_monoidFirst = gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_First_1089607855_1201789390((&Constructor_Data_Monoid_Monoid[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_First_125234255_4179793454(Rebox_Data_Maybe_First_4179793454_125234255(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Get_Data_Maybe_First_semigroupFirst()))))}
}), gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())})))}
	})
	return cache_Data_Maybe_First_monoidFirst
}

var cache_Data_Maybe_First_monadFirst gopurs_runtime.Value
var once_Data_Maybe_First_monadFirst sync.Once
func Get_Data_Maybe_First_monadFirst() gopurs_runtime.Value {
	once_Data_Maybe_First_monadFirst.Do(func() {
		cache_Data_Maybe_First_monadFirst = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_First_1642601656_2568689657(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Maybe_monadMaybe())))}
	})
	return cache_Data_Maybe_First_monadFirst
}

var cache_Data_Maybe_First_invariantFirst gopurs_runtime.Value
var once_Data_Maybe_First_invariantFirst sync.Once
func Get_Data_Maybe_First_invariantFirst() gopurs_runtime.Value {
	once_Data_Maybe_First_invariantFirst.Do(func() {
		cache_Data_Maybe_First_invariantFirst = gopurs_runtime.Value{Type: 9, IntVal: 2396985522, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_First_3818752775_2241343270(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Invariant_Invariant[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Maybe_invariantMaybe())))}
	})
	return cache_Data_Maybe_First_invariantFirst
}

var cache_Data_Maybe_First_functorFirst gopurs_runtime.Value
var once_Data_Maybe_First_functorFirst sync.Once
func Get_Data_Maybe_First_functorFirst() gopurs_runtime.Value {
	once_Data_Maybe_First_functorFirst.Do(func() {
		cache_Data_Maybe_First_functorFirst = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_First_3689823567_2812149806(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Maybe_functorMaybe())))}
	})
	return cache_Data_Maybe_First_functorFirst
}

var cache_Data_Maybe_First_extendFirst gopurs_runtime.Value
var once_Data_Maybe_First_extendFirst sync.Once
func Get_Data_Maybe_First_extendFirst() gopurs_runtime.Value {
	once_Data_Maybe_First_extendFirst.Do(func() {
		cache_Data_Maybe_First_extendFirst = gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_First_570732504_3290176857(gopurs_runtime.CoerceToStruct[Constructor_Control_Extend_Extend[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Maybe_extendMaybe())))}
	})
	return cache_Data_Maybe_First_extendFirst
}

var cache_Data_Maybe_First_eqFirst gopurs_runtime.Value
var once_Data_Maybe_First_eqFirst sync.Once
func Get_Data_Maybe_First_eqFirst() gopurs_runtime.Value {
	once_Data_Maybe_First_eqFirst.Do(func() {
		cache_Data_Maybe_First_eqFirst = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_First_eqFirst(dictEq_0_box)
})
	})
	return cache_Data_Maybe_First_eqFirst
}

var cache_Data_Maybe_First_eq1First gopurs_runtime.Value
var once_Data_Maybe_First_eq1First sync.Once
func Get_Data_Maybe_First_eq1First() gopurs_runtime.Value {
	once_Data_Maybe_First_eq1First.Do(func() {
		cache_Data_Maybe_First_eq1First = gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_First_1662389854_1766074591(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Maybe_eq1Maybe())))}
	})
	return cache_Data_Maybe_First_eq1First
}

var cache_Data_Maybe_First_boundedFirst gopurs_runtime.Value
var once_Data_Maybe_First_boundedFirst sync.Once
func Get_Data_Maybe_First_boundedFirst() gopurs_runtime.Value {
	once_Data_Maybe_First_boundedFirst.Do(func() {
		cache_Data_Maybe_First_boundedFirst = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_First_boundedFirst(dictBounded_0_box)
})
	})
	return cache_Data_Maybe_First_boundedFirst
}

var cache_Data_Maybe_First_bindFirst gopurs_runtime.Value
var once_Data_Maybe_First_bindFirst sync.Once
func Get_Data_Maybe_First_bindFirst() gopurs_runtime.Value {
	once_Data_Maybe_First_bindFirst.Do(func() {
		cache_Data_Maybe_First_bindFirst = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_First_3887487416_2748095225(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Maybe_bindMaybe())))}
	})
	return cache_Data_Maybe_First_bindFirst
}

var cache_Data_Maybe_First_applyFirst gopurs_runtime.Value
var once_Data_Maybe_First_applyFirst sync.Once
func Get_Data_Maybe_First_applyFirst() gopurs_runtime.Value {
	once_Data_Maybe_First_applyFirst.Do(func() {
		cache_Data_Maybe_First_applyFirst = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_First_3552963512_3741347833(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Maybe_applyMaybe())))}
	})
	return cache_Data_Maybe_First_applyFirst
}

var cache_Data_Maybe_First_applicativeFirst gopurs_runtime.Value
var once_Data_Maybe_First_applicativeFirst sync.Once
func Get_Data_Maybe_First_applicativeFirst() gopurs_runtime.Value {
	once_Data_Maybe_First_applicativeFirst.Do(func() {
		cache_Data_Maybe_First_applicativeFirst = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_First_649684152_1439734649(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Maybe_applicativeMaybe())))}
	})
	return cache_Data_Maybe_First_applicativeFirst
}

var cache_Data_Maybe_First_altFirst gopurs_runtime.Value
var once_Data_Maybe_First_altFirst sync.Once
func Get_Data_Maybe_First_altFirst() gopurs_runtime.Value {
	once_Data_Maybe_First_altFirst.Do(func() {
		cache_Data_Maybe_First_altFirst = gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_First_3904200120_3421983481((&Constructor_Control_Alt_Alt[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_First_3689823567_2812149806(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Maybe_functorMaybe())))}
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_0)
if (__t_tag_0 != nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_0)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_1)
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
})})))}
	})
	return cache_Data_Maybe_First_altFirst
}

var cache_Data_Maybe_First_plusFirst gopurs_runtime.Value
var once_Data_Maybe_First_plusFirst sync.Once
func Get_Data_Maybe_First_plusFirst() gopurs_runtime.Value {
	once_Data_Maybe_First_plusFirst.Do(func() {
		cache_Data_Maybe_First_plusFirst = gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_First_267397720_3706288089((&Constructor_Control_Plus_Plus[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_First_3904200120_3421983481(gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Maybe_First_altFirst())))}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}})))}
	})
	return cache_Data_Maybe_First_plusFirst
}

var cache_Data_Maybe_First_alternativeFirst gopurs_runtime.Value
var once_Data_Maybe_First_alternativeFirst sync.Once
func Get_Data_Maybe_First_alternativeFirst() gopurs_runtime.Value {
	once_Data_Maybe_First_alternativeFirst.Do(func() {
		cache_Data_Maybe_First_alternativeFirst = gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_First_673668088_2307501113((&Constructor_Control_Alternative_Alternative[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_First_649684152_1439734649(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Maybe_applicativeMaybe())))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_First_267397720_3706288089(gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Maybe_First_plusFirst())))}
})})))}
	})
	return cache_Data_Maybe_First_alternativeFirst
}

func Call_Data_Maybe_First_First(x_0_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 bool} {
var x_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = x_0_loop
_ = x_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(x_0)}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Maybe_First_showFirst(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
// TAST (Let): showMaybe_1_0 shape=LitRecord bindingType=(ADT ["Data","Show","Show"] [(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])])
showMaybe_1_0 := (&Constructor_Data_Show_Show[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 string
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_1)
if (__t_tag_1 != nil) {
__t3 = (("(Just ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_1.UnsafePtr).V0).StrVal())) + (")")
goto end_branch_3
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_1)
if (__t_tag_2 == nil) {
__t3 = "Nothing"
goto end_branch_3
} else {

}
}
{
__t3 = func() string { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Str(__t3)
})})
_ = showMaybe_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_First_345859663_1386611502((&Constructor_Data_Show_Show[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("First (") + (gopurs_runtime.Apply(gopurs_runtime.Box(showMaybe_1_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_2))}).StrVal())) + (")"))
})})))}
}

func Call_Data_Maybe_First_ordFirst(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): eqMaybe1_1_0 shape=Let(LitRecord) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])])
eqMaybe1_1_0 := (&Constructor_Data_Eq_Eq[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 bool
{
var __t_tag_5 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](x_2)
if (__t_tag_5 == nil) {
var __t_tag_6 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](y_3)
__t7 = (__t_tag_6 == nil)
goto end_branch_7
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](x_2)
var __t_and_4 bool = false
if (__t_tag_2 != nil) {

var __t_tag_3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](y_3)
__t_and_4 = ((__t_tag_3 != nil)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "eq"), (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal) != (0))
}
__t7 = __t_and_4
}
end_branch_7:
return gopurs_runtime.Bool(__t7)
})})
_ = eqMaybe1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_First_2155612431_4177771502((&Constructor_Data_Ord_Ord[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_First_3508461103_3790796878(eqMaybe1_1_0))}
}), gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t15 uint32
{
var __t_tag_8 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](x_2)
if (__t_tag_8 == nil) {
var __t10 uint32
{
var __t_tag_9 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](y_3)
if (__t_tag_9 == nil) {
__t10 = 902936544
goto end_branch_10
} else {

}
}
{
__t10 = 1527465420
}
end_branch_10:
__t15 = __t10
goto end_branch_15
} else {

}
}
{
var __t_tag_11 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](y_3)
if (__t_tag_11 == nil) {
__t15 = 380165415
goto end_branch_15
} else {

}
}
{
var __t_tag_12 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](x_2)
var __t_and_14 bool = false
if (__t_tag_12 != nil) {

var __t_tag_13 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](y_3)
__t_and_14 = (__t_tag_13 != nil)
}
if __t_and_14 {
__t15 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal)
goto end_branch_15
} else {

}
}
{
__t15 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_15:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t15), UnsafePtr: nil}
})})))}
}

func Call_Data_Maybe_First_eqFirst(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_First_3508461103_3790796878((&Constructor_Data_Eq_Eq[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 bool
{
var __t_tag_3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](x_1)
if (__t_tag_3 == nil) {
var __t_tag_4 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](y_2)
__t5 = (__t_tag_4 == nil)
goto end_branch_5
} else {

}
}
{
var __t_tag_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](x_1)
var __t_and_2 bool = false
if (__t_tag_0 != nil) {

var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](y_2)
__t_and_2 = ((__t_tag_1 != nil)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(x_1.UnsafePtr).V0, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(y_2.UnsafePtr).V0).IntVal) != (0))
}
__t5 = __t_and_2
}
end_branch_5:
return gopurs_runtime.Bool(__t5)
})})))}
}

func Call_Data_Maybe_First_boundedFirst(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_First_591333647_2094947566(Rebox_Data_Maybe_First_2094947566_591333647(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Maybe_boundedMaybe(), dictBounded_0)))))}
}

func Rebox_Data_Maybe_First_1089607855_1201789390(in *Constructor_Data_Monoid_Monoid[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(in.V1)}
	return out
}

func Rebox_Data_Maybe_First_125234255_4179793454(in *Constructor_Data_Semigroup_Semigroup[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Maybe_First_1642601656_2568689657(in *Constructor_Control_Monad_Monad[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Monad_Monad[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Monad_Monad[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Maybe_First_1662389854_1766074591(in *Constructor_Data_Eq_Eq1[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Eq_Eq1[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq1[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Maybe_First_170825214_3985601471(in *Constructor_Data_Ord_Ord1[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Ord_Ord1[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord1[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Maybe_First_1892488872_385277032(in *Constructor_Data_Newtype_Newtype[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Maybe_First_2094947566_591333647(in *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]) *Constructor_Data_Bounded_Bounded[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Bounded_Bounded[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](in.V1)
		out.V2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](in.V2)
	return out
}

func Rebox_Data_Maybe_First_2155612431_4177771502(in *Constructor_Data_Ord_Ord[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Maybe_First_267397720_3706288089(in *Constructor_Control_Plus_Plus[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Plus_Plus[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Plus_Plus[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Maybe_First_345859663_1386611502(in *Constructor_Data_Show_Show[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Maybe_First_3508461103_3790796878(in *Constructor_Data_Eq_Eq[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Maybe_First_3552963512_3741347833(in *Constructor_Control_Apply_Apply[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Apply_Apply[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Apply_Apply[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Maybe_First_3689823567_2812149806(in *Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Functor_Functor[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Functor_Functor[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Maybe_First_3818752775_2241343270(in *Constructor_Data_Functor_Invariant_Invariant[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Functor_Invariant_Invariant[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Functor_Invariant_Invariant[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Maybe_First_3887487416_2748095225(in *Constructor_Control_Bind_Bind[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Bind_Bind[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Bind_Bind[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Maybe_First_3904200120_3421983481(in *Constructor_Control_Alt_Alt[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Alt_Alt[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Alt_Alt[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Maybe_First_4179793454_125234255(in *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]) *Constructor_Data_Semigroup_Semigroup[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Semigroup_Semigroup[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Maybe_First_570732504_3290176857(in *Constructor_Control_Extend_Extend[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Extend_Extend[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Extend_Extend[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Maybe_First_591333647_2094947566(in *Constructor_Data_Bounded_Bounded[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(in.V1)}
		out.V2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(in.V2)}
	return out
}

func Rebox_Data_Maybe_First_649684152_1439734649(in *Constructor_Control_Applicative_Applicative[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Maybe_First_673668088_2307501113(in *Constructor_Control_Alternative_Alternative[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}


