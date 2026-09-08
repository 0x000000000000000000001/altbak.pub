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
		cache_Control_Bind_identity = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1)
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

var cache_Control_Bind_discard__770006763 gopurs_runtime.Value
var once_Control_Bind_discard__770006763 sync.Once
func Get_Control_Bind_discard__770006763() gopurs_runtime.Value {
	once_Control_Bind_discard__770006763.Do(func() {
		cache_Control_Bind_discard__770006763 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__770006763(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__770006763
}

var cache_Control_Bind_discard__731412779 gopurs_runtime.Value
var once_Control_Bind_discard__731412779 sync.Once
func Get_Control_Bind_discard__731412779() gopurs_runtime.Value {
	once_Control_Bind_discard__731412779.Do(func() {
		cache_Control_Bind_discard__731412779 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__731412779(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__731412779
}

var cache_Control_Bind_discard__2329751179 gopurs_runtime.Value
var once_Control_Bind_discard__2329751179 sync.Once
func Get_Control_Bind_discard__2329751179() gopurs_runtime.Value {
	once_Control_Bind_discard__2329751179.Do(func() {
		cache_Control_Bind_discard__2329751179 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__2329751179(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__2329751179
}

var cache_Control_Bind_discard__3906178915 gopurs_runtime.Value
var once_Control_Bind_discard__3906178915 sync.Once
func Get_Control_Bind_discard__3906178915() gopurs_runtime.Value {
	once_Control_Bind_discard__3906178915.Do(func() {
		cache_Control_Bind_discard__3906178915 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__3906178915(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_discard__3906178915
}

var cache_Control_Bind_bindProxy gopurs_runtime.Value
var once_Control_Bind_bindProxy sync.Once
func Get_Control_Bind_bindProxy() gopurs_runtime.Value {
	once_Control_Bind_bindProxy.Do(func() {
		cache_Control_Bind_bindProxy = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(Rebox_Control_Bind_4014717844_2748095225((&Constructor_Control_Bind_Bind[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(Rebox_Control_Bind_1303115796_3741347833(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[uint32]](Get_Control_Apply_applyProxy())))}
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

var cache_Control_Bind_bind__4070834660 gopurs_runtime.Value
var once_Control_Bind_bind__4070834660 sync.Once
func Get_Control_Bind_bind__4070834660() gopurs_runtime.Value {
	once_Control_Bind_bind__4070834660.Do(func() {
		cache_Control_Bind_bind__4070834660 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__4070834660(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__4070834660
}

var cache_Control_Bind_bind__982086404 gopurs_runtime.Value
var once_Control_Bind_bind__982086404 sync.Once
func Get_Control_Bind_bind__982086404() gopurs_runtime.Value {
	once_Control_Bind_bind__982086404.Do(func() {
		cache_Control_Bind_bind__982086404 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__982086404(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__982086404
}

var cache_Control_Bind_bind__2667112063 gopurs_runtime.Value
var once_Control_Bind_bind__2667112063 sync.Once
func Get_Control_Bind_bind__2667112063() gopurs_runtime.Value {
	once_Control_Bind_bind__2667112063.Do(func() {
		cache_Control_Bind_bind__2667112063 = gopurs_runtime.Func2(func(__eta_norm_0_unused_0_box gopurs_runtime.Value, __eta_norm_1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__2667112063(__eta_norm_0_unused_0_box, __eta_norm_1_1_box)
})
	})
	return cache_Control_Bind_bind__2667112063
}

var cache_Control_Bind_bind__3327316868 gopurs_runtime.Value
var once_Control_Bind_bind__3327316868 sync.Once
func Get_Control_Bind_bind__3327316868() gopurs_runtime.Value {
	once_Control_Bind_bind__3327316868.Do(func() {
		cache_Control_Bind_bind__3327316868 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__3327316868(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__3327316868
}

var cache_Control_Bind_bind__837570900 gopurs_runtime.Value
var once_Control_Bind_bind__837570900 sync.Once
func Get_Control_Bind_bind__837570900() gopurs_runtime.Value {
	once_Control_Bind_bind__837570900.Do(func() {
		cache_Control_Bind_bind__837570900 = gopurs_runtime.Func2(func(__eta_norm_0_unused_0_box gopurs_runtime.Value, __eta_norm_1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__837570900(__eta_norm_0_unused_0_box, __eta_norm_1_1_box)
})
	})
	return cache_Control_Bind_bind__837570900
}

var cache_Control_Bind_bind__695877027 gopurs_runtime.Value
var once_Control_Bind_bind__695877027 sync.Once
func Get_Control_Bind_bind__695877027() gopurs_runtime.Value {
	once_Control_Bind_bind__695877027.Do(func() {
		cache_Control_Bind_bind__695877027 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__695877027(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__695877027
}

var cache_Control_Bind_bind__2027585355 gopurs_runtime.Value
var once_Control_Bind_bind__2027585355 sync.Once
func Get_Control_Bind_bind__2027585355() gopurs_runtime.Value {
	once_Control_Bind_bind__2027585355.Do(func() {
		cache_Control_Bind_bind__2027585355 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__2027585355(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__2027585355
}

var cache_Control_Bind_bind__3801120152 gopurs_runtime.Value
var once_Control_Bind_bind__3801120152 sync.Once
func Get_Control_Bind_bind__3801120152() gopurs_runtime.Value {
	once_Control_Bind_bind__3801120152.Do(func() {
		cache_Control_Bind_bind__3801120152 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__3801120152(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__3801120152
}

var cache_Control_Bind_bind__64791260 gopurs_runtime.Value
var once_Control_Bind_bind__64791260 sync.Once
func Get_Control_Bind_bind__64791260() gopurs_runtime.Value {
	once_Control_Bind_bind__64791260.Do(func() {
		cache_Control_Bind_bind__64791260 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__64791260(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__64791260
}

var cache_Control_Bind_bind__2198421916 gopurs_runtime.Value
var once_Control_Bind_bind__2198421916 sync.Once
func Get_Control_Bind_bind__2198421916() gopurs_runtime.Value {
	once_Control_Bind_bind__2198421916.Do(func() {
		cache_Control_Bind_bind__2198421916 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__2198421916(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__2198421916
}

var cache_Control_Bind_bind__1179148028 gopurs_runtime.Value
var once_Control_Bind_bind__1179148028 sync.Once
func Get_Control_Bind_bind__1179148028() gopurs_runtime.Value {
	once_Control_Bind_bind__1179148028.Do(func() {
		cache_Control_Bind_bind__1179148028 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1179148028(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__1179148028
}

var cache_Control_Bind_bind__1590050396 gopurs_runtime.Value
var once_Control_Bind_bind__1590050396 sync.Once
func Get_Control_Bind_bind__1590050396() gopurs_runtime.Value {
	once_Control_Bind_bind__1590050396.Do(func() {
		cache_Control_Bind_bind__1590050396 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1590050396(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__1590050396
}

var cache_Control_Bind_bind__1664930140 gopurs_runtime.Value
var once_Control_Bind_bind__1664930140 sync.Once
func Get_Control_Bind_bind__1664930140() gopurs_runtime.Value {
	once_Control_Bind_bind__1664930140.Do(func() {
		cache_Control_Bind_bind__1664930140 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1664930140(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__1664930140
}

var cache_Control_Bind_bind__1510544860 gopurs_runtime.Value
var once_Control_Bind_bind__1510544860 sync.Once
func Get_Control_Bind_bind__1510544860() gopurs_runtime.Value {
	once_Control_Bind_bind__1510544860.Do(func() {
		cache_Control_Bind_bind__1510544860 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1510544860(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__1510544860
}

var cache_Control_Bind_bind__2916363964 gopurs_runtime.Value
var once_Control_Bind_bind__2916363964 sync.Once
func Get_Control_Bind_bind__2916363964() gopurs_runtime.Value {
	once_Control_Bind_bind__2916363964.Do(func() {
		cache_Control_Bind_bind__2916363964 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__2916363964(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__2916363964
}

var cache_Control_Bind_bind__1583830364 gopurs_runtime.Value
var once_Control_Bind_bind__1583830364 sync.Once
func Get_Control_Bind_bind__1583830364() gopurs_runtime.Value {
	once_Control_Bind_bind__1583830364.Do(func() {
		cache_Control_Bind_bind__1583830364 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1583830364(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__1583830364
}

var cache_Control_Bind_bind__2171072156 gopurs_runtime.Value
var once_Control_Bind_bind__2171072156 sync.Once
func Get_Control_Bind_bind__2171072156() gopurs_runtime.Value {
	once_Control_Bind_bind__2171072156.Do(func() {
		cache_Control_Bind_bind__2171072156 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__2171072156(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__2171072156
}

var cache_Control_Bind_bind__3345233820 gopurs_runtime.Value
var once_Control_Bind_bind__3345233820 sync.Once
func Get_Control_Bind_bind__3345233820() gopurs_runtime.Value {
	once_Control_Bind_bind__3345233820.Do(func() {
		cache_Control_Bind_bind__3345233820 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__3345233820(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__3345233820
}

var cache_Control_Bind_bind__1272295772 gopurs_runtime.Value
var once_Control_Bind_bind__1272295772 sync.Once
func Get_Control_Bind_bind__1272295772() gopurs_runtime.Value {
	once_Control_Bind_bind__1272295772.Do(func() {
		cache_Control_Bind_bind__1272295772 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1272295772(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__1272295772
}

var cache_Control_Bind_bind__4188900348 gopurs_runtime.Value
var once_Control_Bind_bind__4188900348 sync.Once
func Get_Control_Bind_bind__4188900348() gopurs_runtime.Value {
	once_Control_Bind_bind__4188900348.Do(func() {
		cache_Control_Bind_bind__4188900348 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__4188900348(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__4188900348
}

var cache_Control_Bind_bind__754360668 gopurs_runtime.Value
var once_Control_Bind_bind__754360668 sync.Once
func Get_Control_Bind_bind__754360668() gopurs_runtime.Value {
	once_Control_Bind_bind__754360668.Do(func() {
		cache_Control_Bind_bind__754360668 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__754360668(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__754360668
}

var cache_Control_Bind_bind__454683484 gopurs_runtime.Value
var once_Control_Bind_bind__454683484 sync.Once
func Get_Control_Bind_bind__454683484() gopurs_runtime.Value {
	once_Control_Bind_bind__454683484.Do(func() {
		cache_Control_Bind_bind__454683484 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__454683484(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__454683484
}

var cache_Control_Bind_bind__3834475196 gopurs_runtime.Value
var once_Control_Bind_bind__3834475196 sync.Once
func Get_Control_Bind_bind__3834475196() gopurs_runtime.Value {
	once_Control_Bind_bind__3834475196.Do(func() {
		cache_Control_Bind_bind__3834475196 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__3834475196(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__3834475196
}

var cache_Control_Bind_bind__637019804 gopurs_runtime.Value
var once_Control_Bind_bind__637019804 sync.Once
func Get_Control_Bind_bind__637019804() gopurs_runtime.Value {
	once_Control_Bind_bind__637019804.Do(func() {
		cache_Control_Bind_bind__637019804 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__637019804(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__637019804
}

var cache_Control_Bind_bind__2515926748 gopurs_runtime.Value
var once_Control_Bind_bind__2515926748 sync.Once
func Get_Control_Bind_bind__2515926748() gopurs_runtime.Value {
	once_Control_Bind_bind__2515926748.Do(func() {
		cache_Control_Bind_bind__2515926748 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__2515926748(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__2515926748
}

var cache_Control_Bind_bind__2869226556 gopurs_runtime.Value
var once_Control_Bind_bind__2869226556 sync.Once
func Get_Control_Bind_bind__2869226556() gopurs_runtime.Value {
	once_Control_Bind_bind__2869226556.Do(func() {
		cache_Control_Bind_bind__2869226556 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__2869226556(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__2869226556
}

var cache_Control_Bind_bind__365045148 gopurs_runtime.Value
var once_Control_Bind_bind__365045148 sync.Once
func Get_Control_Bind_bind__365045148() gopurs_runtime.Value {
	once_Control_Bind_bind__365045148.Do(func() {
		cache_Control_Bind_bind__365045148 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__365045148(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__365045148
}

var cache_Control_Bind_bind__1668159612 gopurs_runtime.Value
var once_Control_Bind_bind__1668159612 sync.Once
func Get_Control_Bind_bind__1668159612() gopurs_runtime.Value {
	once_Control_Bind_bind__1668159612.Do(func() {
		cache_Control_Bind_bind__1668159612 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1668159612(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__1668159612
}

var cache_Control_Bind_bind__1762716060 gopurs_runtime.Value
var once_Control_Bind_bind__1762716060 sync.Once
func Get_Control_Bind_bind__1762716060() gopurs_runtime.Value {
	once_Control_Bind_bind__1762716060.Do(func() {
		cache_Control_Bind_bind__1762716060 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1762716060(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__1762716060
}

var cache_Control_Bind_bind__3189147804 gopurs_runtime.Value
var once_Control_Bind_bind__3189147804 sync.Once
func Get_Control_Bind_bind__3189147804() gopurs_runtime.Value {
	once_Control_Bind_bind__3189147804.Do(func() {
		cache_Control_Bind_bind__3189147804 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__3189147804(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__3189147804
}

var cache_Control_Bind_bind__4056826780 gopurs_runtime.Value
var once_Control_Bind_bind__4056826780 sync.Once
func Get_Control_Bind_bind__4056826780() gopurs_runtime.Value {
	once_Control_Bind_bind__4056826780.Do(func() {
		cache_Control_Bind_bind__4056826780 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__4056826780(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__4056826780
}

var cache_Control_Bind_bind__2547376988 gopurs_runtime.Value
var once_Control_Bind_bind__2547376988 sync.Once
func Get_Control_Bind_bind__2547376988() gopurs_runtime.Value {
	once_Control_Bind_bind__2547376988.Do(func() {
		cache_Control_Bind_bind__2547376988 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__2547376988(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__2547376988
}

var cache_Control_Bind_bind__2015572188 gopurs_runtime.Value
var once_Control_Bind_bind__2015572188 sync.Once
func Get_Control_Bind_bind__2015572188() gopurs_runtime.Value {
	once_Control_Bind_bind__2015572188.Do(func() {
		cache_Control_Bind_bind__2015572188 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__2015572188(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__2015572188
}

var cache_Control_Bind_bind__2286040444 gopurs_runtime.Value
var once_Control_Bind_bind__2286040444 sync.Once
func Get_Control_Bind_bind__2286040444() gopurs_runtime.Value {
	once_Control_Bind_bind__2286040444.Do(func() {
		cache_Control_Bind_bind__2286040444 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__2286040444(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__2286040444
}

var cache_Control_Bind_bind__1031802972 gopurs_runtime.Value
var once_Control_Bind_bind__1031802972 sync.Once
func Get_Control_Bind_bind__1031802972() gopurs_runtime.Value {
	once_Control_Bind_bind__1031802972.Do(func() {
		cache_Control_Bind_bind__1031802972 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1031802972(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__1031802972
}

var cache_Control_Bind_bind__1634036188 gopurs_runtime.Value
var once_Control_Bind_bind__1634036188 sync.Once
func Get_Control_Bind_bind__1634036188() gopurs_runtime.Value {
	once_Control_Bind_bind__1634036188.Do(func() {
		cache_Control_Bind_bind__1634036188 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1634036188(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__1634036188
}

var cache_Control_Bind_bind__1859829227 gopurs_runtime.Value
var once_Control_Bind_bind__1859829227 sync.Once
func Get_Control_Bind_bind__1859829227() gopurs_runtime.Value {
	once_Control_Bind_bind__1859829227.Do(func() {
		cache_Control_Bind_bind__1859829227 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1859829227(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__1859829227
}

var cache_Control_Bind_bind__4101075371 gopurs_runtime.Value
var once_Control_Bind_bind__4101075371 sync.Once
func Get_Control_Bind_bind__4101075371() gopurs_runtime.Value {
	once_Control_Bind_bind__4101075371.Do(func() {
		cache_Control_Bind_bind__4101075371 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__4101075371(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__4101075371
}

var cache_Control_Bind_bind__2487957236 gopurs_runtime.Value
var once_Control_Bind_bind__2487957236 sync.Once
func Get_Control_Bind_bind__2487957236() gopurs_runtime.Value {
	once_Control_Bind_bind__2487957236.Do(func() {
		cache_Control_Bind_bind__2487957236 = gopurs_runtime.Func2(func(dict_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__2487957236(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](dict_0_box), __eta_norm_0_1_box)
})
	})
	return cache_Control_Bind_bind__2487957236
}

var cache_Control_Bind_bindFlipped gopurs_runtime.Value
var once_Control_Bind_bindFlipped sync.Once
func Get_Control_Bind_bindFlipped() gopurs_runtime.Value {
	once_Control_Bind_bindFlipped.Do(func() {
		cache_Control_Bind_bindFlipped = gopurs_runtime.Func3(func(dictBind_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bindFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](dictBind_0_box), b_1_box, a_2_box)
})
	})
	return cache_Control_Bind_bindFlipped
}

var cache_Control_Bind_bindFlipped__892068708 gopurs_runtime.Value
var once_Control_Bind_bindFlipped__892068708 sync.Once
func Get_Control_Bind_bindFlipped__892068708() gopurs_runtime.Value {
	once_Control_Bind_bindFlipped__892068708.Do(func() {
		cache_Control_Bind_bindFlipped__892068708 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Control_Bind_bindFlipped__892068708(__eta_norm_1_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]](__eta_norm_0_1_box))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Control_Bind_bindFlipped__892068708
}

var cache_Control_Bind_bindFlipped__892410344 gopurs_runtime.Value
var once_Control_Bind_bindFlipped__892410344 sync.Once
func Get_Control_Bind_bindFlipped__892410344() gopurs_runtime.Value {
	once_Control_Bind_bindFlipped__892410344.Do(func() {
		cache_Control_Bind_bindFlipped__892410344 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Control_Bind_bindFlipped__892410344(__eta_norm_1_unused_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]](__eta_norm_0_1_box))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Control_Bind_bindFlipped__892410344
}

var cache_Control_Bind_bindFlipped__2379315604 gopurs_runtime.Value
var once_Control_Bind_bindFlipped__2379315604 sync.Once
func Get_Control_Bind_bindFlipped__2379315604() gopurs_runtime.Value {
	once_Control_Bind_bindFlipped__2379315604.Do(func() {
		cache_Control_Bind_bindFlipped__2379315604 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Control_Bind_bindFlipped__2379315604(__eta_norm_1_unused_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[string]](__eta_norm_0_1_box))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Control_Bind_bindFlipped__2379315604
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

var cache_Control_Bind_composeKleisliFlipped__1579089109 gopurs_runtime.Value
var once_Control_Bind_composeKleisliFlipped__1579089109 sync.Once
func Get_Control_Bind_composeKleisliFlipped__1579089109() gopurs_runtime.Value {
	once_Control_Bind_composeKleisliFlipped__1579089109.Do(func() {
		cache_Control_Bind_composeKleisliFlipped__1579089109 = gopurs_runtime.Func3(func(f_unused_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Control_Bind_composeKleisliFlipped__1579089109(f_unused_0_box, g_1_box, a_2_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Control_Bind_composeKleisliFlipped__1579089109
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
return gopurs_runtime.RecordGet(dictBind_0, "bind")
})})))}
	})
	return cache_Control_Bind_discardProxy
}

var cache_Control_Bind_discardUnit gopurs_runtime.Value
var once_Control_Bind_discardUnit sync.Once
func Get_Control_Bind_discardUnit() gopurs_runtime.Value {
	once_Control_Bind_discardUnit.Do(func() {
		cache_Control_Bind_discardUnit = gopurs_runtime.Value{Type: 9, IntVal: 2260728934, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Discard[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictBind_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictBind_0, "bind")
})}))}
	})
	return cache_Control_Bind_discardUnit
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

var cache_Control_Bind_join__637077509 gopurs_runtime.Value
var once_Control_Bind_join__637077509 sync.Once
func Get_Control_Bind_join__637077509() gopurs_runtime.Value {
	once_Control_Bind_join__637077509.Do(func() {
		cache_Control_Bind_join__637077509 = gopurs_runtime.Func(func(m_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Control_Bind_join__637077509(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]]](m_0_box))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Control_Bind_join__637077509
}

type Constructor_Control_Bind_Bind[T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[4032919565] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Bind_Bind[any])(ptr)
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
		c := (*Constructor_Control_Bind_Discard[any])(ptr)
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
				return gopurs_runtime.RecordDict([]string{"Apply0", "bind"}, []gopurs_runtime.Value{orig.Apply0, orig.bind})
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
				return gopurs_runtime.RecordDict([]string{"discard"}, []gopurs_runtime.Value{orig.discard})
				}())
}

func Call_Control_Bind_discard(dict_0_loop *Constructor_Control_Bind_Discard[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Discard[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Bind_discard__770006763(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__770006763:
for {
if false { continue discard__770006763 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(__eta_norm_1_0, gopurs_runtime.Value{})
_ = __local_var_2_0
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta_norm_0_1, __local_var_2_0), gopurs_runtime.Value{})
})
}
}

func Call_Control_Bind_discard__731412779(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__731412779:
for {
if false { continue discard__731412779 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(__eta_norm_1_0, gopurs_runtime.Value{})
_ = __local_var_2_0
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta_norm_0_1, __local_var_2_0), gopurs_runtime.Value{})
})
}
}

func Call_Control_Bind_discard__2329751179(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__2329751179:
for {
if false { continue discard__2329751179 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(__eta_norm_1_0, gopurs_runtime.Value{})
_ = __local_var_2_0
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta_norm_0_1, __local_var_2_0), gopurs_runtime.Value{})
})
}
}

func Call_Control_Bind_discard__3906178915(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__3906178915:
for {
if false { continue discard__3906178915 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(Get_App_warmup(), gopurs_runtime.Value{})
_ = __local_var_2_0
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta_norm_0_1, __local_var_2_0), gopurs_runtime.Value{})
})
}
}

func Call_Control_Bind_bind(dict_0_loop *Constructor_Control_Bind_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__4070834660(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__4070834660:
for {
if false { continue bind__4070834660 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
var __t2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_1_0)
if (__t_tag_0 != nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(__eta_norm_0_1, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__eta_norm_1_0.UnsafePtr).V0))
goto end_branch_2
} else {

}
}
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_1_0)
if (__t_tag_1 == nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_2
} else {

}
}
{
__t2 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)}
}
}

func Call_Control_Bind_bind__982086404(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__982086404:
for {
if false { continue bind__982086404 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
var __t2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_1_0)
if (__t_tag_0 != nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(__eta_norm_0_1, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__eta_norm_1_0.UnsafePtr).V0))
goto end_branch_2
} else {

}
}
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_1_0)
if (__t_tag_1 == nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_2
} else {

}
}
{
__t2 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)}
}
}

func Call_Control_Bind_bind__2667112063(__eta_norm_0_unused_0_loop gopurs_runtime.Value, __eta_norm_1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__2667112063:
for {
if false { continue bind__2667112063 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
var __eta_norm_1_1 gopurs_runtime.Value = __eta_norm_1_1_loop
_ = __eta_norm_1_1
var __t2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_1_1)
if (__t_tag_0 != nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]]((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__eta_norm_1_1.UnsafePtr).V0)
goto end_branch_2
} else {

}
}
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_1_1)
if (__t_tag_1 == nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_2
} else {

}
}
{
__t2 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)}
}
}

func Call_Control_Bind_bind__3327316868(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__3327316868:
for {
if false { continue bind__3327316868 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
var __t2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_1_0)
if (__t_tag_0 != nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(__eta_norm_0_1, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__eta_norm_1_0.UnsafePtr).V0))
goto end_branch_2
} else {

}
}
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_1_0)
if (__t_tag_1 == nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_2
} else {

}
}
{
__t2 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)}
}
}

func Call_Control_Bind_bind__837570900(__eta_norm_0_unused_0_loop gopurs_runtime.Value, __eta_norm_1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__837570900:
for {
if false { continue bind__837570900 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
var __eta_norm_1_1 gopurs_runtime.Value = __eta_norm_1_1_loop
_ = __eta_norm_1_1
var __t3 gopurs_runtime.Value
{
var __t_tag_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_1_1)
if (__t_tag_0 != nil) {
var __t1 gopurs_runtime.Value
{
if ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__eta_norm_1_1.UnsafePtr).V0.StrVal()) == ("") {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Bind_742090555_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[string]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Str((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__eta_norm_1_1.UnsafePtr).V0.StrVal())}))}
}
end_branch_1:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Bind_742090555_3094389156(Rebox_Control_Bind_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t1))))}
goto end_branch_3
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_1_1)
if (__t_tag_2 == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }())}
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t3))}
}
}

func Call_Control_Bind_bind__695877027(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__695877027:
for {
if false { continue bind__695877027 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), __eta_norm_1_0, __eta_norm_0_1)
}
}

func Call_Control_Bind_bind__2027585355(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__2027585355:
for {
if false { continue bind__2027585355 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(__eta_norm_1_0, gopurs_runtime.Value{})
_ = __local_var_2_0
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta_norm_0_1, __local_var_2_0), gopurs_runtime.Value{})
})
}
}

func Call_Control_Bind_bind__3801120152(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__3801120152:
for {
if false { continue bind__3801120152 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_2_0
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta_norm_0_1, __local_var_2_0), gopurs_runtime.Value{})
})
}
}

func Call_Control_Bind_bind__64791260(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__64791260:
for {
if false { continue bind__64791260 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
__local_var_2_0 := gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_AckermannFFICheatcode_describe(), Get_Test_AckermannFFICheatcode_act())
_ = __local_var_2_0
__local_var_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
_ = __local_var_3_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta_norm_0_1, __local_var_3_1), gopurs_runtime.Value{})
})
}
}

func Call_Control_Bind_bind__2198421916(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__2198421916:
for {
if false { continue bind__2198421916 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
__local_var_2_0 := gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_Ackermann_describe(), Get_Test_Ackermann_act())
_ = __local_var_2_0
__local_var_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
_ = __local_var_3_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta_norm_0_1, __local_var_3_1), gopurs_runtime.Value{})
})
}
}

func Call_Control_Bind_bind__1179148028(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__1179148028:
for {
if false { continue bind__1179148028 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
__local_var_2_0 := gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_ArrayOpsFFICheatcode_describe(), Get_Test_ArrayOpsFFICheatcode_act())
_ = __local_var_2_0
__local_var_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
_ = __local_var_3_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta_norm_0_1, __local_var_3_1), gopurs_runtime.Value{})
})
}
}

func Call_Control_Bind_bind__1590050396(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__1590050396:
for {
if false { continue bind__1590050396 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
__local_var_2_0 := gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_ArrayOps_describe(), Get_Test_ArrayOps_act())
_ = __local_var_2_0
__local_var_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
_ = __local_var_3_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta_norm_0_1, __local_var_3_1), gopurs_runtime.Value{})
})
}
}

func Call_Control_Bind_bind__1664930140(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__1664930140:
for {
if false { continue bind__1664930140 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
__local_var_2_0 := gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_AstTreeFFICheatcode_describe(), Get_Test_AstTreeFFICheatcode_act())
_ = __local_var_2_0
__local_var_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
_ = __local_var_3_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta_norm_0_1, __local_var_3_1), gopurs_runtime.Value{})
})
}
}

func Call_Control_Bind_bind__1510544860(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__1510544860:
for {
if false { continue bind__1510544860 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
__local_var_2_0 := gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_AstTree_describe(), Get_Test_AstTree_act())
_ = __local_var_2_0
__local_var_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
_ = __local_var_3_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta_norm_0_1, __local_var_3_1), gopurs_runtime.Value{})
})
}
}

func Call_Control_Bind_bind__2916363964(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__2916363964:
for {
if false { continue bind__2916363964 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
__local_var_2_0 := gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_ChurchFFICheatcode_describe(), Get_Test_ChurchFFICheatcode_act())
_ = __local_var_2_0
__local_var_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
_ = __local_var_3_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta_norm_0_1, __local_var_3_1), gopurs_runtime.Value{})
})
}
}

func Call_Control_Bind_bind__1583830364(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__1583830364:
for {
if false { continue bind__1583830364 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
__local_var_2_0 := gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_Church_describe(), Get_Test_Church_act())
_ = __local_var_2_0
__local_var_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
_ = __local_var_3_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta_norm_0_1, __local_var_3_1), gopurs_runtime.Value{})
})
}
}

func Call_Control_Bind_bind__2171072156(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__2171072156:
for {
if false { continue bind__2171072156 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
__local_var_2_0 := gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_FibFFICheatcode_describe(), Get_Test_FibFFICheatcode_act())
_ = __local_var_2_0
__local_var_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
_ = __local_var_3_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta_norm_0_1, __local_var_3_1), gopurs_runtime.Value{})
})
}
}

func Call_Control_Bind_bind__3345233820(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__3345233820:
for {
if false { continue bind__3345233820 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
__local_var_2_0 := gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_Fib_describe(), Get_Test_Fib_act())
_ = __local_var_2_0
__local_var_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
_ = __local_var_3_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta_norm_0_1, __local_var_3_1), gopurs_runtime.Value{})
})
}
}

func Call_Control_Bind_bind__1272295772(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__1272295772:
for {
if false { continue bind__1272295772 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
__local_var_2_0 := gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_LazyEvaluationFFICheatcode_describe(), Get_Test_LazyEvaluationFFICheatcode_act())
_ = __local_var_2_0
__local_var_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
_ = __local_var_3_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta_norm_0_1, __local_var_3_1), gopurs_runtime.Value{})
})
}
}

func Call_Control_Bind_bind__4188900348(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__4188900348:
for {
if false { continue bind__4188900348 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
__local_var_2_0 := gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_LazyEvaluation_describe(), Get_Test_LazyEvaluation_act())
_ = __local_var_2_0
__local_var_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
_ = __local_var_3_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta_norm_0_1, __local_var_3_1), gopurs_runtime.Value{})
})
}
}

func Call_Control_Bind_bind__754360668(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__754360668:
for {
if false { continue bind__754360668 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
__local_var_2_0 := gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_ListOpsFFICheatcode_describe(), Get_Test_ListOpsFFICheatcode_act())
_ = __local_var_2_0
__local_var_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
_ = __local_var_3_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta_norm_0_1, __local_var_3_1), gopurs_runtime.Value{})
})
}
}

func Call_Control_Bind_bind__454683484(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__454683484:
for {
if false { continue bind__454683484 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
__local_var_2_0 := gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_ListOps_describe(), Get_Test_ListOps_act())
_ = __local_var_2_0
__local_var_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
_ = __local_var_3_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta_norm_0_1, __local_var_3_1), gopurs_runtime.Value{})
})
}
}

func Call_Control_Bind_bind__3834475196(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__3834475196:
for {
if false { continue bind__3834475196 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
__local_var_2_0 := gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_PolymorphismFFICheatcode_describe(), Get_Test_PolymorphismFFICheatcode_act())
_ = __local_var_2_0
__local_var_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
_ = __local_var_3_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta_norm_0_1, __local_var_3_1), gopurs_runtime.Value{})
})
}
}

func Call_Control_Bind_bind__637019804(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__637019804:
for {
if false { continue bind__637019804 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
__local_var_2_0 := gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_Polymorphism_describe(), Get_Test_Polymorphism_act())
_ = __local_var_2_0
__local_var_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
_ = __local_var_3_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta_norm_0_1, __local_var_3_1), gopurs_runtime.Value{})
})
}
}

func Call_Control_Bind_bind__2515926748(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__2515926748:
for {
if false { continue bind__2515926748 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
__local_var_2_0 := gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_PrimesFFICheatcode_describe(), Get_Test_PrimesFFICheatcode_act())
_ = __local_var_2_0
__local_var_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
_ = __local_var_3_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta_norm_0_1, __local_var_3_1), gopurs_runtime.Value{})
})
}
}

func Call_Control_Bind_bind__2869226556(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__2869226556:
for {
if false { continue bind__2869226556 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
__local_var_2_0 := gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_Primes_describe(), Get_Test_Primes_act())
_ = __local_var_2_0
__local_var_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
_ = __local_var_3_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta_norm_0_1, __local_var_3_1), gopurs_runtime.Value{})
})
}
}

func Call_Control_Bind_bind__365045148(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__365045148:
for {
if false { continue bind__365045148 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
__local_var_2_0 := gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_RBTreeFFICheatcode_describe(), Get_Test_RBTreeFFICheatcode_act())
_ = __local_var_2_0
__local_var_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
_ = __local_var_3_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta_norm_0_1, __local_var_3_1), gopurs_runtime.Value{})
})
}
}

func Call_Control_Bind_bind__1668159612(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__1668159612:
for {
if false { continue bind__1668159612 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
__local_var_2_0 := gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_RBTree_describe(), Get_Test_RBTree_act())
_ = __local_var_2_0
__local_var_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
_ = __local_var_3_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta_norm_0_1, __local_var_3_1), gopurs_runtime.Value{})
})
}
}

func Call_Control_Bind_bind__1762716060(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__1762716060:
for {
if false { continue bind__1762716060 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
__local_var_2_0 := gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_RecordsFFICheatcode_describe(), Get_Test_RecordsFFICheatcode_act())
_ = __local_var_2_0
__local_var_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
_ = __local_var_3_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta_norm_0_1, __local_var_3_1), gopurs_runtime.Value{})
})
}
}

func Call_Control_Bind_bind__3189147804(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__3189147804:
for {
if false { continue bind__3189147804 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
__local_var_2_0 := gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_Records_describe(), Get_Test_Records_act())
_ = __local_var_2_0
__local_var_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
_ = __local_var_3_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta_norm_0_1, __local_var_3_1), gopurs_runtime.Value{})
})
}
}

func Call_Control_Bind_bind__4056826780(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__4056826780:
for {
if false { continue bind__4056826780 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
__local_var_2_0 := gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_RowToListFFICheatcode_describe(), Get_Test_RowToListFFICheatcode_act())
_ = __local_var_2_0
__local_var_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
_ = __local_var_3_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta_norm_0_1, __local_var_3_1), gopurs_runtime.Value{})
})
}
}

func Call_Control_Bind_bind__2547376988(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__2547376988:
for {
if false { continue bind__2547376988 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
__local_var_2_0 := gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_RowToList_describe(), Get_Test_RowToList_act())
_ = __local_var_2_0
__local_var_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
_ = __local_var_3_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta_norm_0_1, __local_var_3_1), gopurs_runtime.Value{})
})
}
}

func Call_Control_Bind_bind__2015572188(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__2015572188:
for {
if false { continue bind__2015572188 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
__local_var_2_0 := gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_StateMonadFFICheatcode_describe(), Get_Test_StateMonadFFICheatcode_act())
_ = __local_var_2_0
__local_var_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
_ = __local_var_3_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta_norm_0_1, __local_var_3_1), gopurs_runtime.Value{})
})
}
}

func Call_Control_Bind_bind__2286040444(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__2286040444:
for {
if false { continue bind__2286040444 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
__local_var_2_0 := gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_StateMonad_describe(), Get_Test_StateMonad_act())
_ = __local_var_2_0
__local_var_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
_ = __local_var_3_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta_norm_0_1, __local_var_3_1), gopurs_runtime.Value{})
})
}
}

func Call_Control_Bind_bind__1031802972(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__1031802972:
for {
if false { continue bind__1031802972 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
__local_var_2_0 := gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_TCOFFICheatcode_describe(), Get_Test_TCOFFICheatcode_act())
_ = __local_var_2_0
__local_var_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
_ = __local_var_3_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta_norm_0_1, __local_var_3_1), gopurs_runtime.Value{})
})
}
}

func Call_Control_Bind_bind__1634036188(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__1634036188:
for {
if false { continue bind__1634036188 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
__local_var_2_0 := gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_TCO_describe(), Get_Test_TCO_act())
_ = __local_var_2_0
__local_var_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
_ = __local_var_3_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta_norm_0_1, __local_var_3_1), gopurs_runtime.Value{})
})
}
}

func Call_Control_Bind_bind__1859829227(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__1859829227:
for {
if false { continue bind__1859829227 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(__eta_norm_1_0, gopurs_runtime.Value{})
_ = __local_var_2_0
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta_norm_0_1, __local_var_2_0), gopurs_runtime.Value{})
})
}
}

func Call_Control_Bind_bind__4101075371(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__4101075371:
for {
if false { continue bind__4101075371 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(__eta_norm_1_0, gopurs_runtime.Value{})
_ = __local_var_2_0
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta_norm_0_1, __local_var_2_0), gopurs_runtime.Value{})
})
}
}

func Call_Control_Bind_bind__2487957236(dict_0_loop *Constructor_Control_Bind_Bind[gopurs_runtime.Value], __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__2487957236:
for {
if false { continue bind__2487957236 }
var dict_0 *Constructor_Control_Bind_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply(gopurs_runtime.Box(dict_0.V1), __eta_norm_0_1)
}
}

func Call_Control_Bind_bindFlipped(dictBind_0_loop *Constructor_Control_Bind_Bind[gopurs_runtime.Value], b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 *Constructor_Control_Bind_Bind[gopurs_runtime.Value] = dictBind_0_loop
_ = dictBind_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictBind_0.V1), a_2, b_1)
}

func Call_Control_Bind_bindFlipped__892068708(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]) struct{V0 gopurs_runtime.Value; V1 bool} {
bindFlipped__892068708:
for {
if false { continue bindFlipped__892068708 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date] = __eta_norm_0_1_loop
_ = __eta_norm_0_1
var __t0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__eta_norm_0_1 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(__eta_norm_1_0, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((__eta_norm_0_1).V0)}))
goto end_branch_0
} else {

}
}
{
if (__eta_norm_0_1 == nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_0
} else {

}
}
{
__t0 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_0:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Control_Bind_bindFlipped__892410344(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]) struct{V0 gopurs_runtime.Value; V1 bool} {
bindFlipped__892410344:
for {
if false { continue bindFlipped__892410344 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date] = __eta_norm_0_1_loop
_ = __eta_norm_0_1
var __t0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__eta_norm_0_1 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Date_pred(), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((__eta_norm_0_1).V0)}))
goto end_branch_0
} else {

}
}
{
if (__eta_norm_0_1 == nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_0
} else {

}
}
{
__t0 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_0:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Control_Bind_bindFlipped__2379315604(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop *Constructor_Data_Maybe_Just[string]) struct{V0 gopurs_runtime.Value; V1 bool} {
bindFlipped__2379315604:
for {
if false { continue bindFlipped__2379315604 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 *Constructor_Data_Maybe_Just[string] = __eta_norm_0_1_loop
_ = __eta_norm_0_1
var __t1 gopurs_runtime.Value
{
if (__eta_norm_0_1 != nil) {
var __t0 gopurs_runtime.Value
{
if ((__eta_norm_0_1).V0) == ("") {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Bind_742090555_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[string]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Str((__eta_norm_0_1).V0)}))}
}
end_branch_0:
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Bind_742090555_3094389156(Rebox_Control_Bind_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t0))))}
goto end_branch_1
} else {

}
}
{
if (__eta_norm_0_1 == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }())}
}
end_branch_1:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t1))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
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
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictBind_0.V1), gopurs_runtime.Apply(g_2, a_3), f_1)
}

func Call_Control_Bind_composeKleisliFlipped__1579089109(f_unused_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, a_2_loop string) struct{V0 gopurs_runtime.Value; V1 bool} {
composeKleisliFlipped__1579089109:
for {
if false { continue composeKleisliFlipped__1579089109 }
var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
_ = f_unused_0
var g_1 gopurs_runtime.Value = g_1_loop
_ = g_1
var a_2 string = a_2_loop
_ = a_2
// TAST (Let): __local_var_3_0 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [String])
__local_var_3_0 := Rebox_Control_Bind_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(g_1, gopurs_runtime.Str(a_2))))
_ = __local_var_3_0
var __t2 gopurs_runtime.Value
{
if (__local_var_3_0 != nil) {
var __t1 gopurs_runtime.Value
{
if ((__local_var_3_0).V0) == ("") {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Bind_742090555_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[string]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Str((__local_var_3_0).V0)}))}
}
end_branch_1:
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Bind_742090555_3094389156(Rebox_Control_Bind_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t1))))}
goto end_branch_2
} else {

}
}
{
if (__local_var_3_0 == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }())}
}
end_branch_2:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t2))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
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
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictBind_0.V1), gopurs_runtime.Apply(f_1, a_3), g_2)
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
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictBind_0.V1), cond_1, gopurs_runtime.Func(func(cond_prime__4 gopurs_runtime.Value) gopurs_runtime.Value {
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
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictBind_0.V1), m_1, gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1))
}

func Call_Control_Bind_join__637077509(m_0_loop *Constructor_Data_Maybe_Just[*Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]]) struct{V0 gopurs_runtime.Value; V1 bool} {
join__637077509:
for {
if false { continue join__637077509 }
var m_0 *Constructor_Data_Maybe_Just[*Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]] = m_0_loop
_ = m_0
var __t0 gopurs_runtime.Value
{
if (m_0 != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Bind_2280409795_3094389156((m_0).V0))}
goto end_branch_0
} else {

}
}
{
if (m_0 == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }())}
}
end_branch_0:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t0))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
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

func Rebox_Control_Bind_3094389156_742090555(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[string] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[string]{}
		out.V0 = in.V0.StrVal()
	return out
}

func Rebox_Control_Bind_3559868223_177467250(in *Constructor_Control_Bind_Discard[uint32]) *Constructor_Control_Bind_Discard[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Bind_Discard[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Control_Bind_4014717844_2748095225(in *Constructor_Control_Bind_Bind[uint32]) *Constructor_Control_Bind_Bind[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Bind_Bind[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Control_Bind_742090555_3094389156(in *Constructor_Data_Maybe_Just[string]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Str(in.V0)
	return out
}

func Get_Control_Bind_arrayBind() gopurs_runtime.Value {
	return _Gopurs_Control_Bind_ArrayBind
}
