package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Bind_identity gopurs_runtime.Value
var once_Control_Bind_identity sync.Once
func Get_Control_Bind_identity() gopurs_runtime.Value {
	once_Control_Bind_identity.Do(func() {
		cache_Control_Bind_identity = Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
	})
	return cache_Control_Bind_identity
}

var cache_Control_Bind_Bind_dollar_Dict gopurs_runtime.Value
var once_Control_Bind_Bind_dollar_Dict sync.Once
func Get_Control_Bind_Bind_dollar_Dict() gopurs_runtime.Value {
	once_Control_Bind_Bind_dollar_Dict.Do(func() {
		cache_Control_Bind_Bind_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(Call_Control_Bind_Bind_dollar_Dict(func() struct{
	Apply0 gopurs_runtime.Value
	bind gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Apply0 gopurs_runtime.Value
	bind gopurs_runtime.Value
}{}
					clone.Apply0 = gopurs_runtime.RecordGet(orig, "Apply0")
					clone.bind = gopurs_runtime.RecordGet(orig, "bind")
					return clone
				}()))}
})
	})
	return cache_Control_Bind_Bind_dollar_Dict
}

var cache_Control_Bind_Discard_dollar_Dict gopurs_runtime.Value
var once_Control_Bind_Discard_dollar_Dict sync.Once
func Get_Control_Bind_Discard_dollar_Dict() gopurs_runtime.Value {
	once_Control_Bind_Discard_dollar_Dict.Do(func() {
		cache_Control_Bind_Discard_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2260728934, UnsafePtr: unsafe.Pointer(Call_Control_Bind_Discard_dollar_Dict(func() struct{
	discard gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	discard gopurs_runtime.Value
}{}
					clone.discard = gopurs_runtime.RecordGet(orig, "discard")
					return clone
				}()))}
})
	})
	return cache_Control_Bind_Discard_dollar_Dict
}

var cache_Control_Bind_discard gopurs_runtime.Value
var once_Control_Bind_discard sync.Once
func Get_Control_Bind_discard() gopurs_runtime.Value {
	once_Control_Bind_discard.Do(func() {
		cache_Control_Bind_discard = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Discard[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Bind_discard
}

var cache_Control_Bind_bind gopurs_runtime.Value
var once_Control_Bind_bind sync.Once
func Get_Control_Bind_bind() gopurs_runtime.Value {
	once_Control_Bind_bind.Do(func() {
		cache_Control_Bind_bind = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Bind_bind
}

var cache_Control_Bind_discardUnit gopurs_runtime.Value
var once_Control_Bind_discardUnit sync.Once
func Get_Control_Bind_discardUnit() gopurs_runtime.Value {
	once_Control_Bind_discardUnit.Do(func() {
		cache_Control_Bind_discardUnit = gopurs_runtime.Value{Type: 9, IntVal: 2260728934, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Discard[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictBind_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](dictBind_0))
})}))}
	})
	return cache_Control_Bind_discardUnit
}

var cache_Control_Bind_discard__691961821 gopurs_runtime.Value
var once_Control_Bind_discard__691961821 sync.Once
func Get_Control_Bind_discard__691961821() gopurs_runtime.Value {
	once_Control_Bind_discard__691961821.Do(func() {
		cache_Control_Bind_discard__691961821 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__691961821(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__691961821
}

var cache_Control_Bind_discard__865298732 gopurs_runtime.Value
var once_Control_Bind_discard__865298732 sync.Once
func Get_Control_Bind_discard__865298732() gopurs_runtime.Value {
	once_Control_Bind_discard__865298732.Do(func() {
		cache_Control_Bind_discard__865298732 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__865298732(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__865298732
}

var cache_Control_Bind_discard__520248428 gopurs_runtime.Value
var once_Control_Bind_discard__520248428 sync.Once
func Get_Control_Bind_discard__520248428() gopurs_runtime.Value {
	once_Control_Bind_discard__520248428.Do(func() {
		cache_Control_Bind_discard__520248428 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__520248428(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__520248428
}

var cache_Control_Bind_discard__3511935750 gopurs_runtime.Value
var once_Control_Bind_discard__3511935750 sync.Once
func Get_Control_Bind_discard__3511935750() gopurs_runtime.Value {
	once_Control_Bind_discard__3511935750.Do(func() {
		cache_Control_Bind_discard__3511935750 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__3511935750(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__3511935750
}

var cache_Control_Bind_discard__1066292146 gopurs_runtime.Value
var once_Control_Bind_discard__1066292146 sync.Once
func Get_Control_Bind_discard__1066292146() gopurs_runtime.Value {
	once_Control_Bind_discard__1066292146.Do(func() {
		cache_Control_Bind_discard__1066292146 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__1066292146(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__1066292146
}

var cache_Control_Bind_discard__2704726848 gopurs_runtime.Value
var once_Control_Bind_discard__2704726848 sync.Once
func Get_Control_Bind_discard__2704726848() gopurs_runtime.Value {
	once_Control_Bind_discard__2704726848.Do(func() {
		cache_Control_Bind_discard__2704726848 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__2704726848(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__2704726848
}

var cache_Control_Bind_discard__3651513677 gopurs_runtime.Value
var once_Control_Bind_discard__3651513677 sync.Once
func Get_Control_Bind_discard__3651513677() gopurs_runtime.Value {
	once_Control_Bind_discard__3651513677.Do(func() {
		cache_Control_Bind_discard__3651513677 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__3651513677(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__3651513677
}

var cache_Control_Bind_discard__3637513696 gopurs_runtime.Value
var once_Control_Bind_discard__3637513696 sync.Once
func Get_Control_Bind_discard__3637513696() gopurs_runtime.Value {
	once_Control_Bind_discard__3637513696.Do(func() {
		cache_Control_Bind_discard__3637513696 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__3637513696(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__3637513696
}

var cache_Control_Bind_discard__498166654 gopurs_runtime.Value
var once_Control_Bind_discard__498166654 sync.Once
func Get_Control_Bind_discard__498166654() gopurs_runtime.Value {
	once_Control_Bind_discard__498166654.Do(func() {
		cache_Control_Bind_discard__498166654 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__498166654(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__498166654
}

var cache_Control_Bind_discard__2627477652 gopurs_runtime.Value
var once_Control_Bind_discard__2627477652 sync.Once
func Get_Control_Bind_discard__2627477652() gopurs_runtime.Value {
	once_Control_Bind_discard__2627477652.Do(func() {
		cache_Control_Bind_discard__2627477652 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__2627477652(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__2627477652
}

var cache_Control_Bind_discard__1207148101 gopurs_runtime.Value
var once_Control_Bind_discard__1207148101 sync.Once
func Get_Control_Bind_discard__1207148101() gopurs_runtime.Value {
	once_Control_Bind_discard__1207148101.Do(func() {
		cache_Control_Bind_discard__1207148101 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__1207148101(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__1207148101
}

var cache_Control_Bind_discard__3906355111 gopurs_runtime.Value
var once_Control_Bind_discard__3906355111 sync.Once
func Get_Control_Bind_discard__3906355111() gopurs_runtime.Value {
	once_Control_Bind_discard__3906355111.Do(func() {
		cache_Control_Bind_discard__3906355111 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__3906355111(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__3906355111
}

var cache_Control_Bind_discard__3705517149 gopurs_runtime.Value
var once_Control_Bind_discard__3705517149 sync.Once
func Get_Control_Bind_discard__3705517149() gopurs_runtime.Value {
	once_Control_Bind_discard__3705517149.Do(func() {
		cache_Control_Bind_discard__3705517149 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__3705517149(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__3705517149
}

var cache_Control_Bind_discard__3734845108 gopurs_runtime.Value
var once_Control_Bind_discard__3734845108 sync.Once
func Get_Control_Bind_discard__3734845108() gopurs_runtime.Value {
	once_Control_Bind_discard__3734845108.Do(func() {
		cache_Control_Bind_discard__3734845108 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__3734845108(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__3734845108
}

var cache_Control_Bind_discard__2950130767 gopurs_runtime.Value
var once_Control_Bind_discard__2950130767 sync.Once
func Get_Control_Bind_discard__2950130767() gopurs_runtime.Value {
	once_Control_Bind_discard__2950130767.Do(func() {
		cache_Control_Bind_discard__2950130767 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__2950130767(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__2950130767
}

var cache_Control_Bind_discard__755549702 gopurs_runtime.Value
var once_Control_Bind_discard__755549702 sync.Once
func Get_Control_Bind_discard__755549702() gopurs_runtime.Value {
	once_Control_Bind_discard__755549702.Do(func() {
		cache_Control_Bind_discard__755549702 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__755549702(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__755549702
}

var cache_Control_Bind_discard__2757891740 gopurs_runtime.Value
var once_Control_Bind_discard__2757891740 sync.Once
func Get_Control_Bind_discard__2757891740() gopurs_runtime.Value {
	once_Control_Bind_discard__2757891740.Do(func() {
		cache_Control_Bind_discard__2757891740 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__2757891740(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__2757891740
}

var cache_Control_Bind_discard__4200878445 gopurs_runtime.Value
var once_Control_Bind_discard__4200878445 sync.Once
func Get_Control_Bind_discard__4200878445() gopurs_runtime.Value {
	once_Control_Bind_discard__4200878445.Do(func() {
		cache_Control_Bind_discard__4200878445 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__4200878445(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__4200878445
}

var cache_Control_Bind_discard__1521145623 gopurs_runtime.Value
var once_Control_Bind_discard__1521145623 sync.Once
func Get_Control_Bind_discard__1521145623() gopurs_runtime.Value {
	once_Control_Bind_discard__1521145623.Do(func() {
		cache_Control_Bind_discard__1521145623 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__1521145623(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__1521145623
}

var cache_Control_Bind_discard__41305420 gopurs_runtime.Value
var once_Control_Bind_discard__41305420 sync.Once
func Get_Control_Bind_discard__41305420() gopurs_runtime.Value {
	once_Control_Bind_discard__41305420.Do(func() {
		cache_Control_Bind_discard__41305420 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__41305420(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__41305420
}

var cache_Control_Bind_discard__842189721 gopurs_runtime.Value
var once_Control_Bind_discard__842189721 sync.Once
func Get_Control_Bind_discard__842189721() gopurs_runtime.Value {
	once_Control_Bind_discard__842189721.Do(func() {
		cache_Control_Bind_discard__842189721 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__842189721(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__842189721
}

var cache_Control_Bind_discard__1644983759 gopurs_runtime.Value
var once_Control_Bind_discard__1644983759 sync.Once
func Get_Control_Bind_discard__1644983759() gopurs_runtime.Value {
	once_Control_Bind_discard__1644983759.Do(func() {
		cache_Control_Bind_discard__1644983759 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__1644983759(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__1644983759
}

var cache_Control_Bind_discard__1382635538 gopurs_runtime.Value
var once_Control_Bind_discard__1382635538 sync.Once
func Get_Control_Bind_discard__1382635538() gopurs_runtime.Value {
	once_Control_Bind_discard__1382635538.Do(func() {
		cache_Control_Bind_discard__1382635538 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__1382635538(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__1382635538
}

var cache_Control_Bind_discard__2468631704 gopurs_runtime.Value
var once_Control_Bind_discard__2468631704 sync.Once
func Get_Control_Bind_discard__2468631704() gopurs_runtime.Value {
	once_Control_Bind_discard__2468631704.Do(func() {
		cache_Control_Bind_discard__2468631704 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__2468631704(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__2468631704
}

var cache_Control_Bind_discard__631035863 gopurs_runtime.Value
var once_Control_Bind_discard__631035863 sync.Once
func Get_Control_Bind_discard__631035863() gopurs_runtime.Value {
	once_Control_Bind_discard__631035863.Do(func() {
		cache_Control_Bind_discard__631035863 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__631035863(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__631035863
}

var cache_Control_Bind_discard__3057069797 gopurs_runtime.Value
var once_Control_Bind_discard__3057069797 sync.Once
func Get_Control_Bind_discard__3057069797() gopurs_runtime.Value {
	once_Control_Bind_discard__3057069797.Do(func() {
		cache_Control_Bind_discard__3057069797 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__3057069797(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__3057069797
}

var cache_Control_Bind_discard__574185058 gopurs_runtime.Value
var once_Control_Bind_discard__574185058 sync.Once
func Get_Control_Bind_discard__574185058() gopurs_runtime.Value {
	once_Control_Bind_discard__574185058.Do(func() {
		cache_Control_Bind_discard__574185058 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__574185058(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__574185058
}

var cache_Control_Bind_discard__258171883 gopurs_runtime.Value
var once_Control_Bind_discard__258171883 sync.Once
func Get_Control_Bind_discard__258171883() gopurs_runtime.Value {
	once_Control_Bind_discard__258171883.Do(func() {
		cache_Control_Bind_discard__258171883 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__258171883(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__258171883
}

var cache_Control_Bind_discard__1581641049 gopurs_runtime.Value
var once_Control_Bind_discard__1581641049 sync.Once
func Get_Control_Bind_discard__1581641049() gopurs_runtime.Value {
	once_Control_Bind_discard__1581641049.Do(func() {
		cache_Control_Bind_discard__1581641049 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__1581641049(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__1581641049
}

var cache_Control_Bind_discard__1364771502 gopurs_runtime.Value
var once_Control_Bind_discard__1364771502 sync.Once
func Get_Control_Bind_discard__1364771502() gopurs_runtime.Value {
	once_Control_Bind_discard__1364771502.Do(func() {
		cache_Control_Bind_discard__1364771502 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__1364771502(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__1364771502
}

var cache_Control_Bind_discard__1689051973 gopurs_runtime.Value
var once_Control_Bind_discard__1689051973 sync.Once
func Get_Control_Bind_discard__1689051973() gopurs_runtime.Value {
	once_Control_Bind_discard__1689051973.Do(func() {
		cache_Control_Bind_discard__1689051973 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__1689051973(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__1689051973
}

var cache_Control_Bind_discard__4277594644 gopurs_runtime.Value
var once_Control_Bind_discard__4277594644 sync.Once
func Get_Control_Bind_discard__4277594644() gopurs_runtime.Value {
	once_Control_Bind_discard__4277594644.Do(func() {
		cache_Control_Bind_discard__4277594644 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__4277594644(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__4277594644
}

var cache_Control_Bind_discard__1920595157 gopurs_runtime.Value
var once_Control_Bind_discard__1920595157 sync.Once
func Get_Control_Bind_discard__1920595157() gopurs_runtime.Value {
	once_Control_Bind_discard__1920595157.Do(func() {
		cache_Control_Bind_discard__1920595157 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__1920595157(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__1920595157
}

var cache_Control_Bind_discard__4142830707 gopurs_runtime.Value
var once_Control_Bind_discard__4142830707 sync.Once
func Get_Control_Bind_discard__4142830707() gopurs_runtime.Value {
	once_Control_Bind_discard__4142830707.Do(func() {
		cache_Control_Bind_discard__4142830707 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__4142830707(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__4142830707
}

var cache_Control_Bind_discard__1965855283 gopurs_runtime.Value
var once_Control_Bind_discard__1965855283 sync.Once
func Get_Control_Bind_discard__1965855283() gopurs_runtime.Value {
	once_Control_Bind_discard__1965855283.Do(func() {
		cache_Control_Bind_discard__1965855283 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__1965855283(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__1965855283
}

var cache_Control_Bind_discard__963184288 gopurs_runtime.Value
var once_Control_Bind_discard__963184288 sync.Once
func Get_Control_Bind_discard__963184288() gopurs_runtime.Value {
	once_Control_Bind_discard__963184288.Do(func() {
		cache_Control_Bind_discard__963184288 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__963184288(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__963184288
}

var cache_Control_Bind_discard__2569697110 gopurs_runtime.Value
var once_Control_Bind_discard__2569697110 sync.Once
func Get_Control_Bind_discard__2569697110() gopurs_runtime.Value {
	once_Control_Bind_discard__2569697110.Do(func() {
		cache_Control_Bind_discard__2569697110 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__2569697110(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__2569697110
}

var cache_Control_Bind_discard__2973373841 gopurs_runtime.Value
var once_Control_Bind_discard__2973373841 sync.Once
func Get_Control_Bind_discard__2973373841() gopurs_runtime.Value {
	once_Control_Bind_discard__2973373841.Do(func() {
		cache_Control_Bind_discard__2973373841 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__2973373841(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__2973373841
}

var cache_Control_Bind_discard__2695068429 gopurs_runtime.Value
var once_Control_Bind_discard__2695068429 sync.Once
func Get_Control_Bind_discard__2695068429() gopurs_runtime.Value {
	once_Control_Bind_discard__2695068429.Do(func() {
		cache_Control_Bind_discard__2695068429 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__2695068429(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__2695068429
}

var cache_Control_Bind_discard__2846824090 gopurs_runtime.Value
var once_Control_Bind_discard__2846824090 sync.Once
func Get_Control_Bind_discard__2846824090() gopurs_runtime.Value {
	once_Control_Bind_discard__2846824090.Do(func() {
		cache_Control_Bind_discard__2846824090 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__2846824090(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__2846824090
}

var cache_Control_Bind_discard__3007957918 gopurs_runtime.Value
var once_Control_Bind_discard__3007957918 sync.Once
func Get_Control_Bind_discard__3007957918() gopurs_runtime.Value {
	once_Control_Bind_discard__3007957918.Do(func() {
		cache_Control_Bind_discard__3007957918 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__3007957918(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__3007957918
}

var cache_Control_Bind_discard__2615273511 gopurs_runtime.Value
var once_Control_Bind_discard__2615273511 sync.Once
func Get_Control_Bind_discard__2615273511() gopurs_runtime.Value {
	once_Control_Bind_discard__2615273511.Do(func() {
		cache_Control_Bind_discard__2615273511 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__2615273511(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__2615273511
}

var cache_Control_Bind_discard__2177238399 gopurs_runtime.Value
var once_Control_Bind_discard__2177238399 sync.Once
func Get_Control_Bind_discard__2177238399() gopurs_runtime.Value {
	once_Control_Bind_discard__2177238399.Do(func() {
		cache_Control_Bind_discard__2177238399 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__2177238399(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__2177238399
}

var cache_Control_Bind_discard__2783083337 gopurs_runtime.Value
var once_Control_Bind_discard__2783083337 sync.Once
func Get_Control_Bind_discard__2783083337() gopurs_runtime.Value {
	once_Control_Bind_discard__2783083337.Do(func() {
		cache_Control_Bind_discard__2783083337 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__2783083337(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__2783083337
}

var cache_Control_Bind_discard__3012029984 gopurs_runtime.Value
var once_Control_Bind_discard__3012029984 sync.Once
func Get_Control_Bind_discard__3012029984() gopurs_runtime.Value {
	once_Control_Bind_discard__3012029984.Do(func() {
		cache_Control_Bind_discard__3012029984 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__3012029984(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__3012029984
}

var cache_Control_Bind_discard__2736724620 gopurs_runtime.Value
var once_Control_Bind_discard__2736724620 sync.Once
func Get_Control_Bind_discard__2736724620() gopurs_runtime.Value {
	once_Control_Bind_discard__2736724620.Do(func() {
		cache_Control_Bind_discard__2736724620 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__2736724620(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__2736724620
}

var cache_Control_Bind_bindProxy gopurs_runtime.Value
var once_Control_Bind_bindProxy sync.Once
func Get_Control_Bind_bindProxy() gopurs_runtime.Value {
	once_Control_Bind_bindProxy.Do(func() {
		cache_Control_Bind_bindProxy = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(Rebox_Control_Bind_4014717844_2748095225((&Constructor_Control_Bind_Bind[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(Rebox_Control_Bind_1303115796_3741347833(Rebox_Control_Bind_3741347833_1303115796(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Control_Apply_applyProxy()))))}
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
})})))}
	})
	return cache_Control_Bind_bindProxy
}

var cache_Control_Bind_bindFn gopurs_runtime.Value
var once_Control_Bind_bindFn sync.Once
func Get_Control_Bind_bindFn() gopurs_runtime.Value {
	once_Control_Bind_bindFn.Do(func() {
		cache_Control_Bind_bindFn = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Control_Apply_applyFn()))}
}), gopurs_runtime.Func3(func(m_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, gopurs_runtime.Apply(m_0, x_2), x_2)
})}))}
	})
	return cache_Control_Bind_bindFn
}

var cache_Control_Bind_bindArray gopurs_runtime.Value
var once_Control_Bind_bindArray sync.Once
func Get_Control_Bind_bindArray() gopurs_runtime.Value {
	once_Control_Bind_bindArray.Do(func() {
		cache_Control_Bind_bindArray = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Control_Apply_applyArray()))}
}), Get_Control_Bind_arrayBind()}))}
	})
	return cache_Control_Bind_bindArray
}

var cache_Control_Bind_bind__3949000746 gopurs_runtime.Value
var once_Control_Bind_bind__3949000746 sync.Once
func Get_Control_Bind_bind__3949000746() gopurs_runtime.Value {
	once_Control_Bind_bind__3949000746.Do(func() {
		cache_Control_Bind_bind__3949000746 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Control_Bind_bind__3949000746(Rebox_Control_Bind_3094389156_2280409795(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_1_0_box)), __eta_norm_0_1_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Control_Bind_bind__3949000746
}

var cache_Control_Bind_bind__211009738 gopurs_runtime.Value
var once_Control_Bind_bind__211009738 sync.Once
func Get_Control_Bind_bind__211009738() gopurs_runtime.Value {
	once_Control_Bind_bind__211009738.Do(func() {
		cache_Control_Bind_bind__211009738 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Control_Bind_bind__211009738(Rebox_Control_Bind_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_1_0_box)), __eta_norm_0_1_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Control_Bind_bind__211009738
}

var cache_Control_Bind_bind__3495580458 gopurs_runtime.Value
var once_Control_Bind_bind__3495580458 sync.Once
func Get_Control_Bind_bind__3495580458() gopurs_runtime.Value {
	once_Control_Bind_bind__3495580458.Do(func() {
		cache_Control_Bind_bind__3495580458 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Control_Bind_bind__3495580458(Rebox_Control_Bind_3094389156_3302008615(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_1_0_box)), __eta_norm_0_1_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Control_Bind_bind__3495580458
}

var cache_Control_Bind_bind__2273829194 gopurs_runtime.Value
var once_Control_Bind_bind__2273829194 sync.Once
func Get_Control_Bind_bind__2273829194() gopurs_runtime.Value {
	once_Control_Bind_bind__2273829194.Do(func() {
		cache_Control_Bind_bind__2273829194 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__2273829194(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__2273829194
}

var cache_Control_Bind_bind__901684298 gopurs_runtime.Value
var once_Control_Bind_bind__901684298 sync.Once
func Get_Control_Bind_bind__901684298() gopurs_runtime.Value {
	once_Control_Bind_bind__901684298.Do(func() {
		cache_Control_Bind_bind__901684298 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__901684298(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__901684298
}

var cache_Control_Bind_bind__3187162954 gopurs_runtime.Value
var once_Control_Bind_bind__3187162954 sync.Once
func Get_Control_Bind_bind__3187162954() gopurs_runtime.Value {
	once_Control_Bind_bind__3187162954.Do(func() {
		cache_Control_Bind_bind__3187162954 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__3187162954(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__3187162954
}

var cache_Control_Bind_bind__1550128635 gopurs_runtime.Value
var once_Control_Bind_bind__1550128635 sync.Once
func Get_Control_Bind_bind__1550128635() gopurs_runtime.Value {
	once_Control_Bind_bind__1550128635.Do(func() {
		cache_Control_Bind_bind__1550128635 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1550128635(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__1550128635
}

var cache_Control_Bind_bind__905501900 gopurs_runtime.Value
var once_Control_Bind_bind__905501900 sync.Once
func Get_Control_Bind_bind__905501900() gopurs_runtime.Value {
	once_Control_Bind_bind__905501900.Do(func() {
		cache_Control_Bind_bind__905501900 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__905501900(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__905501900
}

var cache_Control_Bind_bind__1892703400 gopurs_runtime.Value
var once_Control_Bind_bind__1892703400 sync.Once
func Get_Control_Bind_bind__1892703400() gopurs_runtime.Value {
	once_Control_Bind_bind__1892703400.Do(func() {
		cache_Control_Bind_bind__1892703400 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1892703400(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__1892703400
}

var cache_Control_Bind_bind__844793448 gopurs_runtime.Value
var once_Control_Bind_bind__844793448 sync.Once
func Get_Control_Bind_bind__844793448() gopurs_runtime.Value {
	once_Control_Bind_bind__844793448.Do(func() {
		cache_Control_Bind_bind__844793448 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__844793448(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__844793448
}

var cache_Control_Bind_bind__1703273960 gopurs_runtime.Value
var once_Control_Bind_bind__1703273960 sync.Once
func Get_Control_Bind_bind__1703273960() gopurs_runtime.Value {
	once_Control_Bind_bind__1703273960.Do(func() {
		cache_Control_Bind_bind__1703273960 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1703273960(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__1703273960
}

var cache_Control_Bind_bind__3238444488 gopurs_runtime.Value
var once_Control_Bind_bind__3238444488 sync.Once
func Get_Control_Bind_bind__3238444488() gopurs_runtime.Value {
	once_Control_Bind_bind__3238444488.Do(func() {
		cache_Control_Bind_bind__3238444488 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__3238444488(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__3238444488
}

var cache_Control_Bind_bind__1327903464 gopurs_runtime.Value
var once_Control_Bind_bind__1327903464 sync.Once
func Get_Control_Bind_bind__1327903464() gopurs_runtime.Value {
	once_Control_Bind_bind__1327903464.Do(func() {
		cache_Control_Bind_bind__1327903464 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1327903464(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__1327903464
}

var cache_Control_Bind_bind__2569443944 gopurs_runtime.Value
var once_Control_Bind_bind__2569443944 sync.Once
func Get_Control_Bind_bind__2569443944() gopurs_runtime.Value {
	once_Control_Bind_bind__2569443944.Do(func() {
		cache_Control_Bind_bind__2569443944 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__2569443944(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__2569443944
}

var cache_Control_Bind_bind__2159884072 gopurs_runtime.Value
var once_Control_Bind_bind__2159884072 sync.Once
func Get_Control_Bind_bind__2159884072() gopurs_runtime.Value {
	once_Control_Bind_bind__2159884072.Do(func() {
		cache_Control_Bind_bind__2159884072 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__2159884072(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__2159884072
}

var cache_Control_Bind_bind__1338100968 gopurs_runtime.Value
var once_Control_Bind_bind__1338100968 sync.Once
func Get_Control_Bind_bind__1338100968() gopurs_runtime.Value {
	once_Control_Bind_bind__1338100968.Do(func() {
		cache_Control_Bind_bind__1338100968 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1338100968(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__1338100968
}

var cache_Control_Bind_bind__3352327720 gopurs_runtime.Value
var once_Control_Bind_bind__3352327720 sync.Once
func Get_Control_Bind_bind__3352327720() gopurs_runtime.Value {
	once_Control_Bind_bind__3352327720.Do(func() {
		cache_Control_Bind_bind__3352327720 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__3352327720(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__3352327720
}

var cache_Control_Bind_bind__1106606024 gopurs_runtime.Value
var once_Control_Bind_bind__1106606024 sync.Once
func Get_Control_Bind_bind__1106606024() gopurs_runtime.Value {
	once_Control_Bind_bind__1106606024.Do(func() {
		cache_Control_Bind_bind__1106606024 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1106606024(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__1106606024
}

var cache_Control_Bind_bind__3889554344 gopurs_runtime.Value
var once_Control_Bind_bind__3889554344 sync.Once
func Get_Control_Bind_bind__3889554344() gopurs_runtime.Value {
	once_Control_Bind_bind__3889554344.Do(func() {
		cache_Control_Bind_bind__3889554344 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__3889554344(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__3889554344
}

var cache_Control_Bind_bind__2906096168 gopurs_runtime.Value
var once_Control_Bind_bind__2906096168 sync.Once
func Get_Control_Bind_bind__2906096168() gopurs_runtime.Value {
	once_Control_Bind_bind__2906096168.Do(func() {
		cache_Control_Bind_bind__2906096168 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__2906096168(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__2906096168
}

var cache_Control_Bind_bind__3506955112 gopurs_runtime.Value
var once_Control_Bind_bind__3506955112 sync.Once
func Get_Control_Bind_bind__3506955112() gopurs_runtime.Value {
	once_Control_Bind_bind__3506955112.Do(func() {
		cache_Control_Bind_bind__3506955112 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__3506955112(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__3506955112
}

var cache_Control_Bind_bind__2494544072 gopurs_runtime.Value
var once_Control_Bind_bind__2494544072 sync.Once
func Get_Control_Bind_bind__2494544072() gopurs_runtime.Value {
	once_Control_Bind_bind__2494544072.Do(func() {
		cache_Control_Bind_bind__2494544072 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__2494544072(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__2494544072
}

var cache_Control_Bind_bind__1865029352 gopurs_runtime.Value
var once_Control_Bind_bind__1865029352 sync.Once
func Get_Control_Bind_bind__1865029352() gopurs_runtime.Value {
	once_Control_Bind_bind__1865029352.Do(func() {
		cache_Control_Bind_bind__1865029352 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1865029352(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__1865029352
}

var cache_Control_Bind_bind__1677660904 gopurs_runtime.Value
var once_Control_Bind_bind__1677660904 sync.Once
func Get_Control_Bind_bind__1677660904() gopurs_runtime.Value {
	once_Control_Bind_bind__1677660904.Do(func() {
		cache_Control_Bind_bind__1677660904 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1677660904(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__1677660904
}

var cache_Control_Bind_bind__548632872 gopurs_runtime.Value
var once_Control_Bind_bind__548632872 sync.Once
func Get_Control_Bind_bind__548632872() gopurs_runtime.Value {
	once_Control_Bind_bind__548632872.Do(func() {
		cache_Control_Bind_bind__548632872 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__548632872(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__548632872
}

var cache_Control_Bind_bind__2202160712 gopurs_runtime.Value
var once_Control_Bind_bind__2202160712 sync.Once
func Get_Control_Bind_bind__2202160712() gopurs_runtime.Value {
	once_Control_Bind_bind__2202160712.Do(func() {
		cache_Control_Bind_bind__2202160712 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__2202160712(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__2202160712
}

var cache_Control_Bind_bind__2785545000 gopurs_runtime.Value
var once_Control_Bind_bind__2785545000 sync.Once
func Get_Control_Bind_bind__2785545000() gopurs_runtime.Value {
	once_Control_Bind_bind__2785545000.Do(func() {
		cache_Control_Bind_bind__2785545000 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__2785545000(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__2785545000
}

var cache_Control_Bind_bind__468113064 gopurs_runtime.Value
var once_Control_Bind_bind__468113064 sync.Once
func Get_Control_Bind_bind__468113064() gopurs_runtime.Value {
	once_Control_Bind_bind__468113064.Do(func() {
		cache_Control_Bind_bind__468113064 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__468113064(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__468113064
}

var cache_Control_Bind_bind__3763285544 gopurs_runtime.Value
var once_Control_Bind_bind__3763285544 sync.Once
func Get_Control_Bind_bind__3763285544() gopurs_runtime.Value {
	once_Control_Bind_bind__3763285544.Do(func() {
		cache_Control_Bind_bind__3763285544 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__3763285544(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__3763285544
}

var cache_Control_Bind_bind__1709860040 gopurs_runtime.Value
var once_Control_Bind_bind__1709860040 sync.Once
func Get_Control_Bind_bind__1709860040() gopurs_runtime.Value {
	once_Control_Bind_bind__1709860040.Do(func() {
		cache_Control_Bind_bind__1709860040 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1709860040(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__1709860040
}

var cache_Control_Bind_bind__3297983976 gopurs_runtime.Value
var once_Control_Bind_bind__3297983976 sync.Once
func Get_Control_Bind_bind__3297983976() gopurs_runtime.Value {
	once_Control_Bind_bind__3297983976.Do(func() {
		cache_Control_Bind_bind__3297983976 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__3297983976(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__3297983976
}

var cache_Control_Bind_bind__3933711464 gopurs_runtime.Value
var once_Control_Bind_bind__3933711464 sync.Once
func Get_Control_Bind_bind__3933711464() gopurs_runtime.Value {
	once_Control_Bind_bind__3933711464.Do(func() {
		cache_Control_Bind_bind__3933711464 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__3933711464(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__3933711464
}

var cache_Control_Bind_bind__3080250408 gopurs_runtime.Value
var once_Control_Bind_bind__3080250408 sync.Once
func Get_Control_Bind_bind__3080250408() gopurs_runtime.Value {
	once_Control_Bind_bind__3080250408.Do(func() {
		cache_Control_Bind_bind__3080250408 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__3080250408(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__3080250408
}

var cache_Control_Bind_bind__1471461544 gopurs_runtime.Value
var once_Control_Bind_bind__1471461544 sync.Once
func Get_Control_Bind_bind__1471461544() gopurs_runtime.Value {
	once_Control_Bind_bind__1471461544.Do(func() {
		cache_Control_Bind_bind__1471461544 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1471461544(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__1471461544
}

var cache_Control_Bind_bind__1691268040 gopurs_runtime.Value
var once_Control_Bind_bind__1691268040 sync.Once
func Get_Control_Bind_bind__1691268040() gopurs_runtime.Value {
	once_Control_Bind_bind__1691268040.Do(func() {
		cache_Control_Bind_bind__1691268040 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1691268040(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__1691268040
}

var cache_Control_Bind_bind__1372361832 gopurs_runtime.Value
var once_Control_Bind_bind__1372361832 sync.Once
func Get_Control_Bind_bind__1372361832() gopurs_runtime.Value {
	once_Control_Bind_bind__1372361832.Do(func() {
		cache_Control_Bind_bind__1372361832 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1372361832(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__1372361832
}

var cache_Control_Bind_bind__3940832552 gopurs_runtime.Value
var once_Control_Bind_bind__3940832552 sync.Once
func Get_Control_Bind_bind__3940832552() gopurs_runtime.Value {
	once_Control_Bind_bind__3940832552.Do(func() {
		cache_Control_Bind_bind__3940832552 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__3940832552(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__3940832552
}

var cache_Control_Bind_bind__927062664 gopurs_runtime.Value
var once_Control_Bind_bind__927062664 sync.Once
func Get_Control_Bind_bind__927062664() gopurs_runtime.Value {
	once_Control_Bind_bind__927062664.Do(func() {
		cache_Control_Bind_bind__927062664 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__927062664(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__927062664
}

var cache_Control_Bind_bind__2680013544 gopurs_runtime.Value
var once_Control_Bind_bind__2680013544 sync.Once
func Get_Control_Bind_bind__2680013544() gopurs_runtime.Value {
	once_Control_Bind_bind__2680013544.Do(func() {
		cache_Control_Bind_bind__2680013544 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__2680013544(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__2680013544
}

var cache_Control_Bind_bind__1437226920 gopurs_runtime.Value
var once_Control_Bind_bind__1437226920 sync.Once
func Get_Control_Bind_bind__1437226920() gopurs_runtime.Value {
	once_Control_Bind_bind__1437226920.Do(func() {
		cache_Control_Bind_bind__1437226920 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1437226920(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__1437226920
}

var cache_Control_Bind_bind__1577959528 gopurs_runtime.Value
var once_Control_Bind_bind__1577959528 sync.Once
func Get_Control_Bind_bind__1577959528() gopurs_runtime.Value {
	once_Control_Bind_bind__1577959528.Do(func() {
		cache_Control_Bind_bind__1577959528 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1577959528(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__1577959528
}

var cache_Control_Bind_bind__1025225448 gopurs_runtime.Value
var once_Control_Bind_bind__1025225448 sync.Once
func Get_Control_Bind_bind__1025225448() gopurs_runtime.Value {
	once_Control_Bind_bind__1025225448.Do(func() {
		cache_Control_Bind_bind__1025225448 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1025225448(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__1025225448
}

var cache_Control_Bind_bind__406106760 gopurs_runtime.Value
var once_Control_Bind_bind__406106760 sync.Once
func Get_Control_Bind_bind__406106760() gopurs_runtime.Value {
	once_Control_Bind_bind__406106760.Do(func() {
		cache_Control_Bind_bind__406106760 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__406106760(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__406106760
}

var cache_Control_Bind_bind__3158749992 gopurs_runtime.Value
var once_Control_Bind_bind__3158749992 sync.Once
func Get_Control_Bind_bind__3158749992() gopurs_runtime.Value {
	once_Control_Bind_bind__3158749992.Do(func() {
		cache_Control_Bind_bind__3158749992 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__3158749992(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__3158749992
}

var cache_Control_Bind_bind__1035761384 gopurs_runtime.Value
var once_Control_Bind_bind__1035761384 sync.Once
func Get_Control_Bind_bind__1035761384() gopurs_runtime.Value {
	once_Control_Bind_bind__1035761384.Do(func() {
		cache_Control_Bind_bind__1035761384 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1035761384(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__1035761384
}

var cache_Control_Bind_bind__4096788840 gopurs_runtime.Value
var once_Control_Bind_bind__4096788840 sync.Once
func Get_Control_Bind_bind__4096788840() gopurs_runtime.Value {
	once_Control_Bind_bind__4096788840.Do(func() {
		cache_Control_Bind_bind__4096788840 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__4096788840(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__4096788840
}

var cache_Control_Bind_bind__1242703944 gopurs_runtime.Value
var once_Control_Bind_bind__1242703944 sync.Once
func Get_Control_Bind_bind__1242703944() gopurs_runtime.Value {
	once_Control_Bind_bind__1242703944.Do(func() {
		cache_Control_Bind_bind__1242703944 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1242703944(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__1242703944
}

var cache_Control_Bind_bind__455159848 gopurs_runtime.Value
var once_Control_Bind_bind__455159848 sync.Once
func Get_Control_Bind_bind__455159848() gopurs_runtime.Value {
	once_Control_Bind_bind__455159848.Do(func() {
		cache_Control_Bind_bind__455159848 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__455159848(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__455159848
}

var cache_Control_Bind_bind__3614875944 gopurs_runtime.Value
var once_Control_Bind_bind__3614875944 sync.Once
func Get_Control_Bind_bind__3614875944() gopurs_runtime.Value {
	once_Control_Bind_bind__3614875944.Do(func() {
		cache_Control_Bind_bind__3614875944 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__3614875944(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__3614875944
}

var cache_Control_Bind_bind__489458088 gopurs_runtime.Value
var once_Control_Bind_bind__489458088 sync.Once
func Get_Control_Bind_bind__489458088() gopurs_runtime.Value {
	once_Control_Bind_bind__489458088.Do(func() {
		cache_Control_Bind_bind__489458088 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__489458088(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__489458088
}

var cache_Control_Bind_bind__2532035099 gopurs_runtime.Value
var once_Control_Bind_bind__2532035099 sync.Once
func Get_Control_Bind_bind__2532035099() gopurs_runtime.Value {
	once_Control_Bind_bind__2532035099.Do(func() {
		cache_Control_Bind_bind__2532035099 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__2532035099(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__2532035099
}

var cache_Control_Bind_bind__3972441179 gopurs_runtime.Value
var once_Control_Bind_bind__3972441179 sync.Once
func Get_Control_Bind_bind__3972441179() gopurs_runtime.Value {
	once_Control_Bind_bind__3972441179.Do(func() {
		cache_Control_Bind_bind__3972441179 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__3972441179(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__3972441179
}

var cache_Control_Bind_bindFlipped gopurs_runtime.Value
var once_Control_Bind_bindFlipped sync.Once
func Get_Control_Bind_bindFlipped() gopurs_runtime.Value {
	once_Control_Bind_bindFlipped.Do(func() {
		cache_Control_Bind_bindFlipped = gopurs_runtime.Func(func(dictBind_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bindFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](dictBind_0_box))
})
	})
	return cache_Control_Bind_bindFlipped
}

var cache_Control_Bind_bindFlipped__3709447330 gopurs_runtime.Value
var once_Control_Bind_bindFlipped__3709447330 sync.Once
func Get_Control_Bind_bindFlipped__3709447330() gopurs_runtime.Value {
	once_Control_Bind_bindFlipped__3709447330.Do(func() {
		cache_Control_Bind_bindFlipped__3709447330 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Control_Bind_bindFlipped__3709447330(__eta_norm_1_unused_0_box, Rebox_Control_Bind_3094389156_2280409795(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_0_1_box)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Control_Bind_bindFlipped__3709447330
}

var cache_Control_Bind_bindFlipped__2900220074 gopurs_runtime.Value
var once_Control_Bind_bindFlipped__2900220074 sync.Once
func Get_Control_Bind_bindFlipped__2900220074() gopurs_runtime.Value {
	once_Control_Bind_bindFlipped__2900220074.Do(func() {
		cache_Control_Bind_bindFlipped__2900220074 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Control_Bind_bindFlipped__2900220074(__eta_norm_1_0_box, Rebox_Control_Bind_3094389156_2280409795(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_0_1_box)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Control_Bind_bindFlipped__2900220074
}

var cache_Control_Bind_composeKleisliFlipped gopurs_runtime.Value
var once_Control_Bind_composeKleisliFlipped sync.Once
func Get_Control_Bind_composeKleisliFlipped() gopurs_runtime.Value {
	once_Control_Bind_composeKleisliFlipped.Do(func() {
		cache_Control_Bind_composeKleisliFlipped = gopurs_runtime.Func4(func(dictBind_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value, a_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_composeKleisliFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](dictBind_0_box), f_1_box, g_2_box, a_3_box)
})
	})
	return cache_Control_Bind_composeKleisliFlipped
}

var cache_Control_Bind_composeKleisliFlipped__1603318047 gopurs_runtime.Value
var once_Control_Bind_composeKleisliFlipped__1603318047 sync.Once
func Get_Control_Bind_composeKleisliFlipped__1603318047() gopurs_runtime.Value {
	once_Control_Bind_composeKleisliFlipped__1603318047.Do(func() {
		cache_Control_Bind_composeKleisliFlipped__1603318047 = gopurs_runtime.Func3(func(f_unused_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Control_Bind_composeKleisliFlipped__1603318047(f_unused_0_box, g_1_box, a_2_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Control_Bind_composeKleisliFlipped__1603318047
}

var cache_Control_Bind_composeKleisli gopurs_runtime.Value
var once_Control_Bind_composeKleisli sync.Once
func Get_Control_Bind_composeKleisli() gopurs_runtime.Value {
	once_Control_Bind_composeKleisli.Do(func() {
		cache_Control_Bind_composeKleisli = gopurs_runtime.Func4(func(dictBind_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value, a_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_composeKleisli(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](dictBind_0_box), f_1_box, g_2_box, a_3_box)
})
	})
	return cache_Control_Bind_composeKleisli
}

var cache_Control_Bind_discardProxy gopurs_runtime.Value
var once_Control_Bind_discardProxy sync.Once
func Get_Control_Bind_discardProxy() gopurs_runtime.Value {
	once_Control_Bind_discardProxy.Do(func() {
		cache_Control_Bind_discardProxy = gopurs_runtime.Value{Type: 9, IntVal: 2260728934, UnsafePtr: unsafe.Pointer(Rebox_Control_Bind_3559868223_177467250((&Constructor_Control_Bind_Discard[uint32]{1, gopurs_runtime.Func(func(dictBind_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](dictBind_0))
})})))}
	})
	return cache_Control_Bind_discardProxy
}

var cache_Control_Bind_ifM gopurs_runtime.Value
var once_Control_Bind_ifM sync.Once
func Get_Control_Bind_ifM() gopurs_runtime.Value {
	once_Control_Bind_ifM.Do(func() {
		cache_Control_Bind_ifM = gopurs_runtime.Func4(func(dictBind_0_box gopurs_runtime.Value, cond_1_box gopurs_runtime.Value, t_2_box gopurs_runtime.Value, f_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_ifM(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](dictBind_0_box), cond_1_box, t_2_box, f_3_box)
})
	})
	return cache_Control_Bind_ifM
}

var cache_Control_Bind_join gopurs_runtime.Value
var once_Control_Bind_join sync.Once
func Get_Control_Bind_join() gopurs_runtime.Value {
	once_Control_Bind_join.Do(func() {
		cache_Control_Bind_join = gopurs_runtime.Func2(func(dictBind_0_box gopurs_runtime.Value, m_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_join(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](dictBind_0_box), m_1_box)
})
	})
	return cache_Control_Bind_join
}

var cache_Control_Bind_join__1287623801 gopurs_runtime.Value
var once_Control_Bind_join__1287623801 sync.Once
func Get_Control_Bind_join__1287623801() gopurs_runtime.Value {
	once_Control_Bind_join__1287623801.Do(func() {
		cache_Control_Bind_join__1287623801 = gopurs_runtime.Func(func(m_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Control_Bind_join__1287623801(Rebox_Control_Bind_3094389156_3516139266(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](m_0_box)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Control_Bind_join__1287623801
}

type Constructor_Control_Bind_Bind[T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[4032919565] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Bind_Bind[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Apply0": return gopurs_runtime.Box(c.V0)
		case "bind": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Control_Bind_Bind: " + key)
		}
	}
}


type Constructor_Control_Bind_Discard[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2260728934] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Bind_Discard[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "discard": return gopurs_runtime.Box(c.V0)
		default: panic("Key not found in dictionary Constructor_Control_Bind_Discard: " + key)
		}
	}
}


func Call_Control_Bind_Bind_dollar_Dict(x_0_loop struct{
	Apply0 gopurs_runtime.Value
	bind gopurs_runtime.Value
}) *Constructor_Control_Bind_Bind[gopurs_runtime.Value] {
var x_0 struct{
	Apply0 gopurs_runtime.Value
	bind gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict2("Apply0", "bind", orig.Apply0, orig.bind)
				}())
}

func Call_Control_Bind_Discard_dollar_Dict(x_0_loop struct{
	discard gopurs_runtime.Value
}) *Constructor_Control_Bind_Discard[gopurs_runtime.Value] {
var x_0 struct{
	discard gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Discard[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict1("discard", orig.discard)
				}())
}

func Call_Control_Bind_discard(dict_0_loop *Constructor_Control_Bind_Discard[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Discard[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_Control_Bind_bind(dict_0_loop *Constructor_Control_Bind_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_Control_Bind_discard__691961821(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__691961821:
for {
if false { continue discard__691961821 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_Aff_bindAff())), __eta_norm_1_0, __eta_norm_0_1)
}
}

func Call_Control_Bind_discard__865298732(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__865298732:
for {
if false { continue discard__865298732 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), __eta_norm_1_0, __eta_norm_0_1)
}
}

func Call_Control_Bind_discard__520248428(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__520248428:
for {
if false { continue discard__520248428 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), __eta_norm_1_0, __eta_norm_0_1)
}
}

func Call_Control_Bind_discard__3511935750(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__3511935750:
for {
if false { continue discard__3511935750 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_ArrayOpsFFI_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_discard__1066292146(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__1066292146:
for {
if false { continue discard__1066292146 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_ChurchFFI_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_discard__2704726848(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__2704726848:
for {
if false { continue discard__2704726848 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_LazyEvaluationFFI_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_discard__3651513677(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__3651513677:
for {
if false { continue discard__3651513677 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_PolymorphismFFI_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_discard__3637513696(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__3637513696:
for {
if false { continue discard__3637513696 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_TCO_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_discard__498166654(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__498166654:
for {
if false { continue discard__498166654 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_StateMonadFFICheatcode_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_discard__2627477652(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__2627477652:
for {
if false { continue discard__2627477652 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_RecordsFFICheatcode_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_discard__1207148101(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__1207148101:
for {
if false { continue discard__1207148101 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_LazyEvaluation_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_discard__3906355111(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__3906355111:
for {
if false { continue discard__3906355111 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_RecordsFFI_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_discard__3705517149(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__3705517149:
for {
if false { continue discard__3705517149 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_ListOpsFFI_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_discard__3734845108(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__3734845108:
for {
if false { continue discard__3734845108 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_AstTreeFFICheatcode_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_discard__2950130767(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__2950130767:
for {
if false { continue discard__2950130767 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_RBTreeFFI_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_discard__755549702(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__755549702:
for {
if false { continue discard__755549702 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_Ackermann_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_discard__2757891740(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__2757891740:
for {
if false { continue discard__2757891740 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_ListOpsFFICheatcode_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_discard__4200878445(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__4200878445:
for {
if false { continue discard__4200878445 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_ArrayOps_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_discard__1521145623(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__1521145623:
for {
if false { continue discard__1521145623 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_RBTreeFFICheatcode_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_discard__41305420(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__41305420:
for {
if false { continue discard__41305420 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_Records_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_discard__842189721(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__842189721:
for {
if false { continue discard__842189721 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_Fib_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_discard__1644983759(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__1644983759:
for {
if false { continue discard__1644983759 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_Primes_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_discard__1382635538(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__1382635538:
for {
if false { continue discard__1382635538 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_TCOFFICheatcode_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_discard__2468631704(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__2468631704:
for {
if false { continue discard__2468631704 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_PolymorphismFFICheatcode_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_discard__631035863(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__631035863:
for {
if false { continue discard__631035863 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_PrimesFFICheatcode_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_discard__3057069797(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__3057069797:
for {
if false { continue discard__3057069797 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_RBTree_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_discard__574185058(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__574185058:
for {
if false { continue discard__574185058 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_ChurchFFICheatcode_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_discard__258171883(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__258171883:
for {
if false { continue discard__258171883 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_StateMonad_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_discard__1581641049(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__1581641049:
for {
if false { continue discard__1581641049 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_AstTreeFFI_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_discard__1364771502(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__1364771502:
for {
if false { continue discard__1364771502 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_AckermannFFI_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_discard__1689051973(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__1689051973:
for {
if false { continue discard__1689051973 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_FibFFI_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_discard__4277594644(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__4277594644:
for {
if false { continue discard__4277594644 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_FibFFICheatcode_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_discard__1920595157(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__1920595157:
for {
if false { continue discard__1920595157 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_AstTree_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_discard__4142830707(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__4142830707:
for {
if false { continue discard__4142830707 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_LazyEvaluationFFICheatcode_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_discard__1965855283(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__1965855283:
for {
if false { continue discard__1965855283 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_ArrayOpsFFICheatcode_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_discard__963184288(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__963184288:
for {
if false { continue discard__963184288 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_StateMonadFFI_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_discard__2569697110(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__2569697110:
for {
if false { continue discard__2569697110 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_ListOps_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_discard__2973373841(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__2973373841:
for {
if false { continue discard__2973373841 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_Polymorphism_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_discard__2695068429(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__2695068429:
for {
if false { continue discard__2695068429 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_Church_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_discard__2846824090(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__2846824090:
for {
if false { continue discard__2846824090 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_TCOFFI_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_discard__3007957918(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__3007957918:
for {
if false { continue discard__3007957918 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_PrimesFFI_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_discard__2615273511(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__2615273511:
for {
if false { continue discard__2615273511 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_AckermannFFICheatcode_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_discard__2177238399(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__2177238399:
for {
if false { continue discard__2177238399 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Get_AppFFICheatcode_warmup(), __eta_norm_0_1)
}
}

func Call_Control_Bind_discard__2783083337(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__2783083337:
for {
if false { continue discard__2783083337 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Get_AppFFI_warmup(), __eta_norm_0_1)
}
}

func Call_Control_Bind_discard__3012029984(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__3012029984:
for {
if false { continue discard__3012029984 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Get_App_warmup(), __eta_norm_0_1)
}
}

func Call_Control_Bind_discard__2736724620(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__2736724620:
for {
if false { continue discard__2736724620 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), __eta_norm_1_0, __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__3949000746(__eta_norm_1_0_loop *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date], __eta_norm_0_1_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
bind__3949000746:
for {
if false { continue bind__3949000746 }
var __eta_norm_1_0 *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date] = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Apply2(Rebox_Control_Bind_2748095225_3887487416(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_Maybe_bindMaybe())).V1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Bind_2280409795_3094389156(__eta_norm_1_0))}, __eta_norm_0_1)
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Control_Bind_bind__211009738(__eta_norm_1_0_loop *Constructor_Data_Maybe_Just[int64], __eta_norm_0_1_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
bind__211009738:
for {
if false { continue bind__211009738 }
var __eta_norm_1_0 *Constructor_Data_Maybe_Just[int64] = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Apply2(Rebox_Control_Bind_2748095225_3887487416(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_Maybe_bindMaybe())).V1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Bind_1170268447_3094389156(__eta_norm_1_0))}, __eta_norm_0_1)
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Control_Bind_bind__3495580458(__eta_norm_1_0_loop *Constructor_Data_Maybe_Just[struct{
	day int64
	hour int64
	millisecond int64
	minute int64
	month int64
	second int64
	year int64
}], __eta_norm_0_1_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
bind__3495580458:
for {
if false { continue bind__3495580458 }
var __eta_norm_1_0 *Constructor_Data_Maybe_Just[struct{
	day int64
	hour int64
	millisecond int64
	minute int64
	month int64
	second int64
	year int64
}] = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Apply2(Rebox_Control_Bind_2748095225_3887487416(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_Maybe_bindMaybe())).V1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Bind_3302008615_3094389156(__eta_norm_1_0))}, __eta_norm_0_1)
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Control_Bind_bind__2273829194(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__2273829194:
for {
if false { continue bind__2273829194 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_Aff_bindAff()).V1, __eta_norm_1_0, __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__901684298(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__901684298:
for {
if false { continue bind__901684298 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_Aff_bindAff()).V1, __eta_norm_1_0, __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__3187162954(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__3187162954:
for {
if false { continue bind__3187162954 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_Aff_bindAff()).V1, __eta_norm_1_0, __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__1550128635(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__1550128635:
for {
if false { continue bind__1550128635 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, __eta_norm_1_0, __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__905501900(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__905501900:
for {
if false { continue bind__905501900 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, Get_Bench_benchNow(), __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__1892703400(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__1892703400:
for {
if false { continue bind__1892703400 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_AckermannFFICheatcode_describe(), Get_Test_AckermannFFICheatcode_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__844793448(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__844793448:
for {
if false { continue bind__844793448 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_AckermannFFI_describe(), Get_Test_AckermannFFI_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__1703273960(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__1703273960:
for {
if false { continue bind__1703273960 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_Ackermann_describe(), Get_Test_Ackermann_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__3238444488(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__3238444488:
for {
if false { continue bind__3238444488 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_ArrayOpsFFICheatcode_describe(), Get_Test_ArrayOpsFFICheatcode_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__1327903464(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__1327903464:
for {
if false { continue bind__1327903464 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_ArrayOpsFFI_describe(), Get_Test_ArrayOpsFFI_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__2569443944(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__2569443944:
for {
if false { continue bind__2569443944 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_ArrayOps_describe(), Get_Test_ArrayOps_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__2159884072(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__2159884072:
for {
if false { continue bind__2159884072 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_AstTreeFFICheatcode_describe(), Get_Test_AstTreeFFICheatcode_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__1338100968(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__1338100968:
for {
if false { continue bind__1338100968 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_AstTreeFFI_describe(), Get_Test_AstTreeFFI_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__3352327720(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__3352327720:
for {
if false { continue bind__3352327720 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_AstTree_describe(), Get_Test_AstTree_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__1106606024(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__1106606024:
for {
if false { continue bind__1106606024 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_ChurchFFICheatcode_describe(), Get_Test_ChurchFFICheatcode_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__3889554344(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__3889554344:
for {
if false { continue bind__3889554344 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_ChurchFFI_describe(), Get_Test_ChurchFFI_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__2906096168(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__2906096168:
for {
if false { continue bind__2906096168 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_Church_describe(), Get_Test_Church_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__3506955112(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__3506955112:
for {
if false { continue bind__3506955112 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_FibFFICheatcode_describe(), Get_Test_FibFFICheatcode_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__2494544072(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__2494544072:
for {
if false { continue bind__2494544072 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_FibFFI_describe(), Get_Test_FibFFI_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__1865029352(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__1865029352:
for {
if false { continue bind__1865029352 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_Fib_describe(), Get_Test_Fib_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__1677660904(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__1677660904:
for {
if false { continue bind__1677660904 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_LazyEvaluationFFICheatcode_describe(), Get_Test_LazyEvaluationFFICheatcode_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__548632872(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__548632872:
for {
if false { continue bind__548632872 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_LazyEvaluationFFI_describe(), Get_Test_LazyEvaluationFFI_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__2202160712(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__2202160712:
for {
if false { continue bind__2202160712 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_LazyEvaluation_describe(), Get_Test_LazyEvaluation_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__2785545000(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__2785545000:
for {
if false { continue bind__2785545000 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_ListOpsFFICheatcode_describe(), Get_Test_ListOpsFFICheatcode_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__468113064(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__468113064:
for {
if false { continue bind__468113064 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_ListOpsFFI_describe(), Get_Test_ListOpsFFI_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__3763285544(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__3763285544:
for {
if false { continue bind__3763285544 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_ListOps_describe(), Get_Test_ListOps_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__1709860040(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__1709860040:
for {
if false { continue bind__1709860040 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_PolymorphismFFICheatcode_describe(), Get_Test_PolymorphismFFICheatcode_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__3297983976(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__3297983976:
for {
if false { continue bind__3297983976 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_PolymorphismFFI_describe(), Get_Test_PolymorphismFFI_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__3933711464(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__3933711464:
for {
if false { continue bind__3933711464 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_Polymorphism_describe(), Get_Test_Polymorphism_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__3080250408(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__3080250408:
for {
if false { continue bind__3080250408 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_PrimesFFICheatcode_describe(), Get_Test_PrimesFFICheatcode_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__1471461544(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__1471461544:
for {
if false { continue bind__1471461544 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_PrimesFFI_describe(), Get_Test_PrimesFFI_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__1691268040(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__1691268040:
for {
if false { continue bind__1691268040 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_Primes_describe(), Get_Test_Primes_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__1372361832(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__1372361832:
for {
if false { continue bind__1372361832 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_RBTreeFFICheatcode_describe(), Get_Test_RBTreeFFICheatcode_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__3940832552(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__3940832552:
for {
if false { continue bind__3940832552 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_RBTreeFFI_describe(), Get_Test_RBTreeFFI_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__927062664(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__927062664:
for {
if false { continue bind__927062664 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_RBTree_describe(), Get_Test_RBTree_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__2680013544(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__2680013544:
for {
if false { continue bind__2680013544 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_RecordsFFICheatcode_describe(), Get_Test_RecordsFFICheatcode_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__1437226920(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__1437226920:
for {
if false { continue bind__1437226920 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_RecordsFFI_describe(), Get_Test_RecordsFFI_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__1577959528(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__1577959528:
for {
if false { continue bind__1577959528 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_Records_describe(), Get_Test_Records_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__1025225448(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__1025225448:
for {
if false { continue bind__1025225448 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_RowToListFFICheatcode_describe(), Get_Test_RowToListFFICheatcode_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__406106760(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__406106760:
for {
if false { continue bind__406106760 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_RowToListFFI_describe(), Get_Test_RowToListFFI_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__3158749992(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__3158749992:
for {
if false { continue bind__3158749992 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_RowToList_describe(), Get_Test_RowToList_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__1035761384(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__1035761384:
for {
if false { continue bind__1035761384 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_StateMonadFFICheatcode_describe(), Get_Test_StateMonadFFICheatcode_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__4096788840(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__4096788840:
for {
if false { continue bind__4096788840 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_StateMonadFFI_describe(), Get_Test_StateMonadFFI_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__1242703944(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__1242703944:
for {
if false { continue bind__1242703944 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_StateMonad_describe(), Get_Test_StateMonad_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__455159848(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__455159848:
for {
if false { continue bind__455159848 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_TCOFFICheatcode_describe(), Get_Test_TCOFFICheatcode_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__3614875944(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__3614875944:
for {
if false { continue bind__3614875944 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_TCOFFI_describe(), Get_Test_TCOFFI_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__489458088(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__489458088:
for {
if false { continue bind__489458088 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_TCO_describe(), Get_Test_TCO_act()), __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__2532035099(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__2532035099:
for {
if false { continue bind__2532035099 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, __eta_norm_1_0, __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__3972441179(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__3972441179:
for {
if false { continue bind__3972441179 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, __eta_norm_1_0, __eta_norm_0_1)
}
}

func Call_Control_Bind_bindFlipped(dictBind_0_loop *Constructor_Control_Bind_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBind_0 *Constructor_Control_Bind_Bind[gopurs_runtime.Value] = dictBind_0_loop
_ = dictBind_0
// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
__local_var_1_0 := Call_Control_Bind_bind(dictBind_0)
_ = __local_var_1_0
return gopurs_runtime.Func2(func(b_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_1_0, a_3, b_2)
})
}

func Call_Control_Bind_bindFlipped__3709447330(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]) struct{V0 gopurs_runtime.Value; V1 bool} {
bindFlipped__3709447330:
for {
if false { continue bindFlipped__3709447330 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date] = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(Call_Control_Bind_bind(Rebox_Control_Bind_3887487416_2748095225(Rebox_Control_Bind_2748095225_3887487416(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_Maybe_bindMaybe())))), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Bind_2280409795_3094389156(__eta_norm_0_1))}, Get_Data_Date_pred())))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Control_Bind_bindFlipped__2900220074(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]) struct{V0 gopurs_runtime.Value; V1 bool} {
bindFlipped__2900220074:
for {
if false { continue bindFlipped__2900220074 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date] = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(Call_Control_Bind_bind(Rebox_Control_Bind_3887487416_2748095225(Rebox_Control_Bind_2748095225_3887487416(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_Maybe_bindMaybe())))), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Bind_2280409795_3094389156(__eta_norm_0_1))}, __eta_norm_1_0)))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Control_Bind_composeKleisliFlipped(dictBind_0_loop *Constructor_Control_Bind_Bind[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value, a_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 *Constructor_Control_Bind_Bind[gopurs_runtime.Value] = dictBind_0_loop
_ = dictBind_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
var a_3 gopurs_runtime.Value = a_3_loop
_ = a_3
return gopurs_runtime.Apply2(Call_Control_Bind_bind(dictBind_0), gopurs_runtime.Apply(g_2, a_3), f_1)
}

func Call_Control_Bind_composeKleisliFlipped__1603318047(f_unused_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, a_2_loop string) struct{V0 gopurs_runtime.Value; V1 bool} {
composeKleisliFlipped__1603318047:
for {
if false { continue composeKleisliFlipped__1603318047 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var g_1 gopurs_runtime.Value = g_1_loop
_ = g_1
var a_2 string = a_2_loop
_ = a_2
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Apply2(Call_Control_Bind_bind(Rebox_Control_Bind_3887487416_2748095225(Rebox_Control_Bind_2748095225_3887487416(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_Maybe_bindMaybe())))), gopurs_runtime.Apply(g_1, gopurs_runtime.Str(a_2)), Get_Data_String_NonEmpty_Internal_fromString())
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Control_Bind_composeKleisli(dictBind_0_loop *Constructor_Control_Bind_Bind[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value, a_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 *Constructor_Control_Bind_Bind[gopurs_runtime.Value] = dictBind_0_loop
_ = dictBind_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
var a_3 gopurs_runtime.Value = a_3_loop
_ = a_3
return gopurs_runtime.Apply2(dictBind_0.V1, gopurs_runtime.Apply(f_1, a_3), g_2)
}

func Call_Control_Bind_ifM(dictBind_0_loop *Constructor_Control_Bind_Bind[gopurs_runtime.Value], cond_1_loop gopurs_runtime.Value, t_2_loop gopurs_runtime.Value, f_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 *Constructor_Control_Bind_Bind[gopurs_runtime.Value] = dictBind_0_loop
_ = dictBind_0
var cond_1 gopurs_runtime.Value = cond_1_loop
_ = cond_1
var t_2 gopurs_runtime.Value = t_2_loop
_ = t_2
var f_3 gopurs_runtime.Value = f_3_loop
_ = f_3
return gopurs_runtime.Apply2(dictBind_0.V1, cond_1, gopurs_runtime.Func(func(cond_prime__4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (cond_prime__4.IntVal) != (0) {
__t0 = t_2
goto end_branch_0
} else {

}
}
{
__t0 = f_3
}
end_branch_0:
return __t0
}))
}

func Call_Control_Bind_join(dictBind_0_loop *Constructor_Control_Bind_Bind[gopurs_runtime.Value], m_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 *Constructor_Control_Bind_Bind[gopurs_runtime.Value] = dictBind_0_loop
_ = dictBind_0
var m_1 gopurs_runtime.Value = m_1_loop
_ = m_1
return gopurs_runtime.Apply2(dictBind_0.V1, m_1, Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}))
}

func Call_Control_Bind_join__1287623801(m_0_loop *Constructor_Data_Maybe_Just[*Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]]) struct{V0 gopurs_runtime.Value; V1 bool} {
join__1287623801:
for {
if false { continue join__1287623801 }
var m_0 *Constructor_Data_Maybe_Just[*Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]] = m_0_loop
_ = m_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Apply2(Rebox_Control_Bind_2748095225_3887487416(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_Maybe_bindMaybe())).V1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Bind_3516139266_3094389156(m_0))}, Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}))
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Rebox_Control_Bind_1170268447_3094389156(in *Constructor_Data_Maybe_Just[int64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Int(in.V0)
	return out
}

func Rebox_Control_Bind_1303115796_3741347833(in *Constructor_Control_Apply_Apply[uint32]) *Constructor_Control_Apply_Apply[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Apply_Apply[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Control_Bind_2280409795_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}

func Rebox_Control_Bind_2748095225_3887487416(in *Constructor_Control_Bind_Bind[gopurs_runtime.Value]) *Constructor_Control_Bind_Bind[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Control_Bind_Bind[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Control_Bind_3094389156_1170268447(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[int64]{}
		out.V0 = in.V0.IntVal
	return out
}

func Rebox_Control_Bind_3094389156_2280409795(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]{}
		out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](in.V0)
	return out
}

func Rebox_Control_Bind_3094389156_3302008615(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[struct{
	day int64
	hour int64
	millisecond int64
	minute int64
	month int64
	second int64
	year int64
}] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[struct{
	day int64
	hour int64
	millisecond int64
	minute int64
	month int64
	second int64
	year int64
}]{}
		out.V0 = func() struct{
	day int64
	hour int64
	millisecond int64
	minute int64
	month int64
	second int64
	year int64
} {
					orig := in.V0
					_ = orig
					clone := struct{
	day int64
	hour int64
	millisecond int64
	minute int64
	month int64
	second int64
	year int64
}{}
					clone.day = gopurs_runtime.RecordGet(orig, "day").IntVal
					clone.hour = gopurs_runtime.RecordGet(orig, "hour").IntVal
					clone.millisecond = gopurs_runtime.RecordGet(orig, "millisecond").IntVal
					clone.minute = gopurs_runtime.RecordGet(orig, "minute").IntVal
					clone.month = gopurs_runtime.RecordGet(orig, "month").IntVal
					clone.second = gopurs_runtime.RecordGet(orig, "second").IntVal
					clone.year = gopurs_runtime.RecordGet(orig, "year").IntVal
					return clone
				}()
	return out
}

func Rebox_Control_Bind_3094389156_3516139266(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]]{}
		out.V0 = Rebox_Control_Bind_3094389156_2280409795(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](in.V0))
	return out
}

func Rebox_Control_Bind_3302008615_3094389156(in *Constructor_Data_Maybe_Just[struct{
	day int64
	hour int64
	millisecond int64
	minute int64
	month int64
	second int64
	year int64
}]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = func() gopurs_runtime.Value {
				orig := in.V0
				_ = orig
				return gopurs_runtime.RecordDict([]string{"day", "hour", "millisecond", "minute", "month", "second", "year"}, []gopurs_runtime.Value{gopurs_runtime.Int(orig.day), gopurs_runtime.Int(orig.hour), gopurs_runtime.Int(orig.millisecond), gopurs_runtime.Int(orig.minute), gopurs_runtime.Int(orig.month), gopurs_runtime.Int(orig.second), gopurs_runtime.Int(orig.year)})
				}()
	return out
}

func Rebox_Control_Bind_3516139266_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Bind_2280409795_3094389156(in.V0))}
	return out
}

func Rebox_Control_Bind_3559868223_177467250(in *Constructor_Control_Bind_Discard[uint32]) *Constructor_Control_Bind_Discard[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Bind_Discard[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Control_Bind_3741347833_1303115796(in *Constructor_Control_Apply_Apply[gopurs_runtime.Value]) *Constructor_Control_Apply_Apply[uint32] {
	if in == nil { return nil }
	out := &Constructor_Control_Apply_Apply[uint32]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Control_Bind_3887487416_2748095225(in *Constructor_Control_Bind_Bind[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Bind_Bind[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Bind_Bind[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Control_Bind_4014717844_2748095225(in *Constructor_Control_Bind_Bind[uint32]) *Constructor_Control_Bind_Bind[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Bind_Bind[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Get_Control_Bind_arrayBind() gopurs_runtime.Value {
	return _Gopurs_Control_Bind_ArrayBind
}
