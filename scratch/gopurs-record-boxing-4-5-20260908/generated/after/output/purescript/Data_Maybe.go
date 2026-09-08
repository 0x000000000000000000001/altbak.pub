package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Maybe_identity gopurs_runtime.Value
var once_Data_Maybe_identity sync.Once
func Get_Data_Maybe_identity() gopurs_runtime.Value {
	once_Data_Maybe_identity.Do(func() {
		cache_Data_Maybe_identity = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1)
	})
	return cache_Data_Maybe_identity
}

var cache_Data_Maybe_Nothing gopurs_runtime.Value
var once_Data_Maybe_Nothing sync.Once
func Get_Data_Maybe_Nothing() gopurs_runtime.Value {
	once_Data_Maybe_Nothing.Do(func() {
		cache_Data_Maybe_Nothing = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
	})
	return cache_Data_Maybe_Nothing
}

var cache_Data_Maybe_Just gopurs_runtime.Value
var once_Data_Maybe_Just sync.Once
func Get_Data_Maybe_Just() gopurs_runtime.Value {
	once_Data_Maybe_Just.Do(func() {
		cache_Data_Maybe_Just = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, value0}))}
})
	})
	return cache_Data_Maybe_Just
}

var cache_Data_Maybe_Just__3521044077 gopurs_runtime.Value
var once_Data_Maybe_Just__3521044077 sync.Once
func Get_Data_Maybe_Just__3521044077() gopurs_runtime.Value {
	once_Data_Maybe_Just__3521044077.Do(func() {
		cache_Data_Maybe_Just__3521044077 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__3521044077(uint32(__eta_norm_0_unused_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Maybe_Just__3521044077
}

var cache_Data_Maybe_Just__2798349578 gopurs_runtime.Value
var once_Data_Maybe_Just__2798349578 sync.Once
func Get_Data_Maybe_Just__2798349578() gopurs_runtime.Value {
	once_Data_Maybe_Just__2798349578.Do(func() {
		cache_Data_Maybe_Just__2798349578 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__2798349578(uint32(__eta_norm_0_unused_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Maybe_Just__2798349578
}

var cache_Data_Maybe_Just__2825732244 gopurs_runtime.Value
var once_Data_Maybe_Just__2825732244 sync.Once
func Get_Data_Maybe_Just__2825732244() gopurs_runtime.Value {
	once_Data_Maybe_Just__2825732244.Do(func() {
		cache_Data_Maybe_Just__2825732244 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__2825732244(uint32(__eta_norm_0_unused_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Maybe_Just__2825732244
}

var cache_Data_Maybe_Just__3832816551 gopurs_runtime.Value
var once_Data_Maybe_Just__3832816551 sync.Once
func Get_Data_Maybe_Just__3832816551() gopurs_runtime.Value {
	once_Data_Maybe_Just__3832816551.Do(func() {
		cache_Data_Maybe_Just__3832816551 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__3832816551(uint32(__eta_norm_0_unused_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Maybe_Just__3832816551
}

var cache_Data_Maybe_Just__4271887697 gopurs_runtime.Value
var once_Data_Maybe_Just__4271887697 sync.Once
func Get_Data_Maybe_Just__4271887697() gopurs_runtime.Value {
	once_Data_Maybe_Just__4271887697.Do(func() {
		cache_Data_Maybe_Just__4271887697 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__4271887697(uint32(__eta_norm_0_unused_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Maybe_Just__4271887697
}

var cache_Data_Maybe_Just__1091121345 gopurs_runtime.Value
var once_Data_Maybe_Just__1091121345 sync.Once
func Get_Data_Maybe_Just__1091121345() gopurs_runtime.Value {
	once_Data_Maybe_Just__1091121345.Do(func() {
		cache_Data_Maybe_Just__1091121345 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__1091121345(uint32(__eta_norm_0_unused_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Maybe_Just__1091121345
}

var cache_Data_Maybe_Just__4216227551 gopurs_runtime.Value
var once_Data_Maybe_Just__4216227551 sync.Once
func Get_Data_Maybe_Just__4216227551() gopurs_runtime.Value {
	once_Data_Maybe_Just__4216227551.Do(func() {
		cache_Data_Maybe_Just__4216227551 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__4216227551(uint32(__eta_norm_0_unused_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Maybe_Just__4216227551
}

var cache_Data_Maybe_Just__3031626846 gopurs_runtime.Value
var once_Data_Maybe_Just__3031626846 sync.Once
func Get_Data_Maybe_Just__3031626846() gopurs_runtime.Value {
	once_Data_Maybe_Just__3031626846.Do(func() {
		cache_Data_Maybe_Just__3031626846 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__3031626846(uint32(__eta_norm_0_unused_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Maybe_Just__3031626846
}

var cache_Data_Maybe_Just__1154318014 gopurs_runtime.Value
var once_Data_Maybe_Just__1154318014 sync.Once
func Get_Data_Maybe_Just__1154318014() gopurs_runtime.Value {
	once_Data_Maybe_Just__1154318014.Do(func() {
		cache_Data_Maybe_Just__1154318014 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__1154318014(uint32(__eta_norm_0_unused_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Maybe_Just__1154318014
}

var cache_Data_Maybe_Just__1938051713 gopurs_runtime.Value
var once_Data_Maybe_Just__1938051713 sync.Once
func Get_Data_Maybe_Just__1938051713() gopurs_runtime.Value {
	once_Data_Maybe_Just__1938051713.Do(func() {
		cache_Data_Maybe_Just__1938051713 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__1938051713(uint32(__eta_norm_0_unused_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Maybe_Just__1938051713
}

var cache_Data_Maybe_Just__3180259945 gopurs_runtime.Value
var once_Data_Maybe_Just__3180259945 sync.Once
func Get_Data_Maybe_Just__3180259945() gopurs_runtime.Value {
	once_Data_Maybe_Just__3180259945.Do(func() {
		cache_Data_Maybe_Just__3180259945 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__3180259945(uint32(__eta_norm_0_unused_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Maybe_Just__3180259945
}

var cache_Data_Maybe_Just__2957110564 gopurs_runtime.Value
var once_Data_Maybe_Just__2957110564 sync.Once
func Get_Data_Maybe_Just__2957110564() gopurs_runtime.Value {
	once_Data_Maybe_Just__2957110564.Do(func() {
		cache_Data_Maybe_Just__2957110564 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__2957110564(uint32(__eta_norm_0_unused_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Maybe_Just__2957110564
}

var cache_Data_Maybe_Just__3561111658 gopurs_runtime.Value
var once_Data_Maybe_Just__3561111658 sync.Once
func Get_Data_Maybe_Just__3561111658() gopurs_runtime.Value {
	once_Data_Maybe_Just__3561111658.Do(func() {
		cache_Data_Maybe_Just__3561111658 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__3561111658(uint32(__eta_norm_0_unused_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Maybe_Just__3561111658
}

var cache_Data_Maybe_Just__3347273307 gopurs_runtime.Value
var once_Data_Maybe_Just__3347273307 sync.Once
func Get_Data_Maybe_Just__3347273307() gopurs_runtime.Value {
	once_Data_Maybe_Just__3347273307.Do(func() {
		cache_Data_Maybe_Just__3347273307 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__3347273307(uint32(__eta_norm_0_unused_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Maybe_Just__3347273307
}

var cache_Data_Maybe_Just__2779760438 gopurs_runtime.Value
var once_Data_Maybe_Just__2779760438 sync.Once
func Get_Data_Maybe_Just__2779760438() gopurs_runtime.Value {
	once_Data_Maybe_Just__2779760438.Do(func() {
		cache_Data_Maybe_Just__2779760438 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__2779760438(uint32(__eta_norm_0_unused_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Maybe_Just__2779760438
}

var cache_Data_Maybe_Just__107066655 gopurs_runtime.Value
var once_Data_Maybe_Just__107066655 sync.Once
func Get_Data_Maybe_Just__107066655() gopurs_runtime.Value {
	once_Data_Maybe_Just__107066655.Do(func() {
		cache_Data_Maybe_Just__107066655 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__107066655(uint32(__eta_norm_0_unused_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Maybe_Just__107066655
}

var cache_Data_Maybe_Just__508806911 gopurs_runtime.Value
var once_Data_Maybe_Just__508806911 sync.Once
func Get_Data_Maybe_Just__508806911() gopurs_runtime.Value {
	once_Data_Maybe_Just__508806911.Do(func() {
		cache_Data_Maybe_Just__508806911 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__508806911(uint32(__eta_norm_0_unused_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Maybe_Just__508806911
}

var cache_Data_Maybe_Just__1678136480 gopurs_runtime.Value
var once_Data_Maybe_Just__1678136480 sync.Once
func Get_Data_Maybe_Just__1678136480() gopurs_runtime.Value {
	once_Data_Maybe_Just__1678136480.Do(func() {
		cache_Data_Maybe_Just__1678136480 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__1678136480(uint32(__eta_norm_0_unused_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Maybe_Just__1678136480
}

var cache_Data_Maybe_Just__2937552249 gopurs_runtime.Value
var once_Data_Maybe_Just__2937552249 sync.Once
func Get_Data_Maybe_Just__2937552249() gopurs_runtime.Value {
	once_Data_Maybe_Just__2937552249.Do(func() {
		cache_Data_Maybe_Just__2937552249 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__2937552249(uint32(__eta_norm_0_unused_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Maybe_Just__2937552249
}

var cache_Data_Maybe_Just__3996651813 gopurs_runtime.Value
var once_Data_Maybe_Just__3996651813 sync.Once
func Get_Data_Maybe_Just__3996651813() gopurs_runtime.Value {
	once_Data_Maybe_Just__3996651813.Do(func() {
		cache_Data_Maybe_Just__3996651813 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__3996651813(gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](__eta_norm_0_0_box))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Maybe_Just__3996651813
}

var cache_Data_Maybe_Just__4227547111 gopurs_runtime.Value
var once_Data_Maybe_Just__4227547111 sync.Once
func Get_Data_Maybe_Just__4227547111() gopurs_runtime.Value {
	once_Data_Maybe_Just__4227547111.Do(func() {
		cache_Data_Maybe_Just__4227547111 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__4227547111(uint32(__eta_norm_0_unused_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Maybe_Just__4227547111
}

var cache_Data_Maybe_Just__3629863589 gopurs_runtime.Value
var once_Data_Maybe_Just__3629863589 sync.Once
func Get_Data_Maybe_Just__3629863589() gopurs_runtime.Value {
	once_Data_Maybe_Just__3629863589.Do(func() {
		cache_Data_Maybe_Just__3629863589 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__3629863589(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](__eta_norm_0_0_box))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Maybe_Just__3629863589
}

var cache_Data_Maybe_Just__2347572452 gopurs_runtime.Value
var once_Data_Maybe_Just__2347572452 sync.Once
func Get_Data_Maybe_Just__2347572452() gopurs_runtime.Value {
	once_Data_Maybe_Just__2347572452.Do(func() {
		cache_Data_Maybe_Just__2347572452 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__2347572452(uint32(__eta_norm_0_unused_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Maybe_Just__2347572452
}

var cache_Data_Maybe_Just__920868163 gopurs_runtime.Value
var once_Data_Maybe_Just__920868163 sync.Once
func Get_Data_Maybe_Just__920868163() gopurs_runtime.Value {
	once_Data_Maybe_Just__920868163.Do(func() {
		cache_Data_Maybe_Just__920868163 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__920868163(uint32(__eta_norm_0_unused_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Maybe_Just__920868163
}

var cache_Data_Maybe_Just__113341384 gopurs_runtime.Value
var once_Data_Maybe_Just__113341384 sync.Once
func Get_Data_Maybe_Just__113341384() gopurs_runtime.Value {
	once_Data_Maybe_Just__113341384.Do(func() {
		cache_Data_Maybe_Just__113341384 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__113341384(uint32(__eta_norm_0_unused_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Maybe_Just__113341384
}

var cache_Data_Maybe_Just__3120863845 gopurs_runtime.Value
var once_Data_Maybe_Just__3120863845 sync.Once
func Get_Data_Maybe_Just__3120863845() gopurs_runtime.Value {
	once_Data_Maybe_Just__3120863845.Do(func() {
		cache_Data_Maybe_Just__3120863845 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__3120863845(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, int64]](__eta_norm_0_0_box))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Maybe_Just__3120863845
}

var cache_Data_Maybe_Just__3356776037 gopurs_runtime.Value
var once_Data_Maybe_Just__3356776037 sync.Once
func Get_Data_Maybe_Just__3356776037() gopurs_runtime.Value {
	once_Data_Maybe_Just__3356776037.Do(func() {
		cache_Data_Maybe_Just__3356776037 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__3356776037((__eta_norm_0_0_box.IntVal) != (0))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Maybe_Just__3356776037
}

var cache_Data_Maybe_Just__2827584037 gopurs_runtime.Value
var once_Data_Maybe_Just__2827584037 sync.Once
func Get_Data_Maybe_Just__2827584037() gopurs_runtime.Value {
	once_Data_Maybe_Just__2827584037.Do(func() {
		cache_Data_Maybe_Just__2827584037 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__2827584037(__eta_norm_0_0_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Maybe_Just__2827584037
}

var cache_Data_Maybe_Just__983269893 gopurs_runtime.Value
var once_Data_Maybe_Just__983269893 sync.Once
func Get_Data_Maybe_Just__983269893() gopurs_runtime.Value {
	once_Data_Maybe_Just__983269893.Do(func() {
		cache_Data_Maybe_Just__983269893 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__983269893(__eta_norm_0_0_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Maybe_Just__983269893
}

var cache_Data_Maybe_Just__157193253 gopurs_runtime.Value
var once_Data_Maybe_Just__157193253 sync.Once
func Get_Data_Maybe_Just__157193253() gopurs_runtime.Value {
	once_Data_Maybe_Just__157193253.Do(func() {
		cache_Data_Maybe_Just__157193253 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__157193253(__eta_norm_0_0_box.FloatVal())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Maybe_Just__157193253
}

var cache_Data_Maybe_Just__436376869 gopurs_runtime.Value
var once_Data_Maybe_Just__436376869 sync.Once
func Get_Data_Maybe_Just__436376869() gopurs_runtime.Value {
	once_Data_Maybe_Just__436376869.Do(func() {
		cache_Data_Maybe_Just__436376869 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__436376869(__eta_norm_0_0_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Maybe_Just__436376869
}

var cache_Data_Maybe_Just__4131531262 gopurs_runtime.Value
var once_Data_Maybe_Just__4131531262 sync.Once
func Get_Data_Maybe_Just__4131531262() gopurs_runtime.Value {
	once_Data_Maybe_Just__4131531262.Do(func() {
		cache_Data_Maybe_Just__4131531262 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__4131531262(__eta_norm_0_unused_0_box)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Maybe_Just__4131531262
}

var cache_Data_Maybe_showMaybe gopurs_runtime.Value
var once_Data_Maybe_showMaybe sync.Once
func Get_Data_Maybe_showMaybe() gopurs_runtime.Value {
	once_Data_Maybe_showMaybe.Do(func() {
		cache_Data_Maybe_showMaybe = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_showMaybe(dictShow_0_box)
})
	})
	return cache_Data_Maybe_showMaybe
}

var cache_Data_Maybe_semigroupMaybe gopurs_runtime.Value
var once_Data_Maybe_semigroupMaybe sync.Once
func Get_Data_Maybe_semigroupMaybe() gopurs_runtime.Value {
	once_Data_Maybe_semigroupMaybe.Do(func() {
		cache_Data_Maybe_semigroupMaybe = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_semigroupMaybe(dictSemigroup_0_box)
})
	})
	return cache_Data_Maybe_semigroupMaybe
}

var cache_Data_Maybe_optional gopurs_runtime.Value
var once_Data_Maybe_optional sync.Once
func Get_Data_Maybe_optional() gopurs_runtime.Value {
	once_Data_Maybe_optional.Do(func() {
		cache_Data_Maybe_optional = gopurs_runtime.Func(func(dictAlt_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_optional(gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt[gopurs_runtime.Value]](dictAlt_0_box))
})
	})
	return cache_Data_Maybe_optional
}

var cache_Data_Maybe_monoidMaybe gopurs_runtime.Value
var once_Data_Maybe_monoidMaybe sync.Once
func Get_Data_Maybe_monoidMaybe() gopurs_runtime.Value {
	once_Data_Maybe_monoidMaybe.Do(func() {
		cache_Data_Maybe_monoidMaybe = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_monoidMaybe(dictSemigroup_0_box)
})
	})
	return cache_Data_Maybe_monoidMaybe
}

var cache_Data_Maybe_maybe_prime_ gopurs_runtime.Value
var once_Data_Maybe_maybe_prime_ sync.Once
func Get_Data_Maybe_maybe_prime_() gopurs_runtime.Value {
	once_Data_Maybe_maybe_prime_.Do(func() {
		cache_Data_Maybe_maybe_prime_ = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_maybe_prime_(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_Data_Maybe_maybe_prime_
}

var cache_Data_Maybe_maybe gopurs_runtime.Value
var once_Data_Maybe_maybe sync.Once
func Get_Data_Maybe_maybe() gopurs_runtime.Value {
	once_Data_Maybe_maybe.Do(func() {
		cache_Data_Maybe_maybe = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_maybe(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_Data_Maybe_maybe
}

var cache_Data_Maybe_maybe__4101309216 gopurs_runtime.Value
var once_Data_Maybe_maybe__4101309216 sync.Once
func Get_Data_Maybe_maybe__4101309216() gopurs_runtime.Value {
	once_Data_Maybe_maybe__4101309216.Do(func() {
		cache_Data_Maybe_maybe__4101309216 = gopurs_runtime.Func3(func(v_unused_0_box gopurs_runtime.Value, v1_unused_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Maybe_maybe__4101309216(uint32(v_unused_0_box.IntVal), v1_unused_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[uint32]](v2_2_box))), UnsafePtr: nil}
})
	})
	return cache_Data_Maybe_maybe__4101309216
}

var cache_Data_Maybe_maybe__3207254341 gopurs_runtime.Value
var once_Data_Maybe_maybe__3207254341 sync.Once
func Get_Data_Maybe_maybe__3207254341() gopurs_runtime.Value {
	once_Data_Maybe_maybe__3207254341.Do(func() {
		cache_Data_Maybe_maybe__3207254341 = gopurs_runtime.Func3(func(v_unused_0_box gopurs_runtime.Value, v1_unused_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Maybe_maybe__3207254341(uint32(v_unused_0_box.IntVal), v1_unused_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[uint32]](v2_2_box))), UnsafePtr: nil}
})
	})
	return cache_Data_Maybe_maybe__3207254341
}

var cache_Data_Maybe_maybe__539590268 gopurs_runtime.Value
var once_Data_Maybe_maybe__539590268 sync.Once
func Get_Data_Maybe_maybe__539590268() gopurs_runtime.Value {
	once_Data_Maybe_maybe__539590268.Do(func() {
		cache_Data_Maybe_maybe__539590268 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_maybe__539590268(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_0_box), v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_2_box))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Maybe_maybe__539590268
}

var cache_Data_Maybe_maybe__17375519 gopurs_runtime.Value
var once_Data_Maybe_maybe__17375519 sync.Once
func Get_Data_Maybe_maybe__17375519() gopurs_runtime.Value {
	once_Data_Maybe_maybe__17375519.Do(func() {
		cache_Data_Maybe_maybe__17375519 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Maybe_maybe__17375519((v_0_box.IntVal) != (0), v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[uint32]](v2_2_box)))
})
	})
	return cache_Data_Maybe_maybe__17375519
}

var cache_Data_Maybe_maybe__3303070975 gopurs_runtime.Value
var once_Data_Maybe_maybe__3303070975 sync.Once
func Get_Data_Maybe_maybe__3303070975() gopurs_runtime.Value {
	once_Data_Maybe_maybe__3303070975.Do(func() {
		cache_Data_Maybe_maybe__3303070975 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Maybe_maybe__3303070975((v_0_box.IntVal) != (0), v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](v2_2_box)))
})
	})
	return cache_Data_Maybe_maybe__3303070975
}

var cache_Data_Maybe_maybe__1344312831 gopurs_runtime.Value
var once_Data_Maybe_maybe__1344312831 sync.Once
func Get_Data_Maybe_maybe__1344312831() gopurs_runtime.Value {
	once_Data_Maybe_maybe__1344312831.Do(func() {
		cache_Data_Maybe_maybe__1344312831 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Maybe_maybe__1344312831((v_0_box.IntVal) != (0), v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[float64]](v2_2_box)))
})
	})
	return cache_Data_Maybe_maybe__1344312831
}

var cache_Data_Maybe_maybe__774193768 gopurs_runtime.Value
var once_Data_Maybe_maybe__774193768 sync.Once
func Get_Data_Maybe_maybe__774193768() gopurs_runtime.Value {
	once_Data_Maybe_maybe__774193768.Do(func() {
		cache_Data_Maybe_maybe__774193768 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Maybe_maybe__774193768(v_0_box.IntVal, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](v2_2_box)))
})
	})
	return cache_Data_Maybe_maybe__774193768
}

var cache_Data_Maybe_maybe__4030275671 gopurs_runtime.Value
var once_Data_Maybe_maybe__4030275671 sync.Once
func Get_Data_Maybe_maybe__4030275671() gopurs_runtime.Value {
	once_Data_Maybe_maybe__4030275671.Do(func() {
		cache_Data_Maybe_maybe__4030275671 = gopurs_runtime.Func3(func(v1_unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Maybe_maybe__4030275671(v1_unused_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](v_1_box), v2_2_box))
})
	})
	return cache_Data_Maybe_maybe__4030275671
}

var cache_Data_Maybe_isNothing gopurs_runtime.Value
var once_Data_Maybe_isNothing sync.Once
func Get_Data_Maybe_isNothing() gopurs_runtime.Value {
	once_Data_Maybe_isNothing.Do(func() {
		cache_Data_Maybe_isNothing = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Maybe_isNothing(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_0_box)))
})
	})
	return cache_Data_Maybe_isNothing
}

var cache_Data_Maybe_isNothing__2526948203 gopurs_runtime.Value
var once_Data_Maybe_isNothing__2526948203 sync.Once
func Get_Data_Maybe_isNothing__2526948203() gopurs_runtime.Value {
	once_Data_Maybe_isNothing__2526948203.Do(func() {
		cache_Data_Maybe_isNothing__2526948203 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Maybe_isNothing__2526948203(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[uint32]](__eta_norm_0_0_box)))
})
	})
	return cache_Data_Maybe_isNothing__2526948203
}

var cache_Data_Maybe_isNothing__684798802 gopurs_runtime.Value
var once_Data_Maybe_isNothing__684798802 sync.Once
func Get_Data_Maybe_isNothing__684798802() gopurs_runtime.Value {
	once_Data_Maybe_isNothing__684798802.Do(func() {
		cache_Data_Maybe_isNothing__684798802 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Maybe_isNothing__684798802(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](__eta_norm_0_0_box)))
})
	})
	return cache_Data_Maybe_isNothing__684798802
}

var cache_Data_Maybe_isJust gopurs_runtime.Value
var once_Data_Maybe_isJust sync.Once
func Get_Data_Maybe_isJust() gopurs_runtime.Value {
	once_Data_Maybe_isJust.Do(func() {
		cache_Data_Maybe_isJust = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Maybe_isJust(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_0_box)))
})
	})
	return cache_Data_Maybe_isJust
}

var cache_Data_Maybe_isJust__684798802 gopurs_runtime.Value
var once_Data_Maybe_isJust__684798802 sync.Once
func Get_Data_Maybe_isJust__684798802() gopurs_runtime.Value {
	once_Data_Maybe_isJust__684798802.Do(func() {
		cache_Data_Maybe_isJust__684798802 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Maybe_isJust__684798802(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](__eta_norm_0_0_box)))
})
	})
	return cache_Data_Maybe_isJust__684798802
}

var cache_Data_Maybe_isJust__912756578 gopurs_runtime.Value
var once_Data_Maybe_isJust__912756578 sync.Once
func Get_Data_Maybe_isJust__912756578() gopurs_runtime.Value {
	once_Data_Maybe_isJust__912756578.Do(func() {
		cache_Data_Maybe_isJust__912756578 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Maybe_isJust__912756578(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[float64]](__eta_norm_0_0_box)))
})
	})
	return cache_Data_Maybe_isJust__912756578
}

var cache_Data_Maybe_genericMaybe gopurs_runtime.Value
var once_Data_Maybe_genericMaybe sync.Once
func Get_Data_Maybe_genericMaybe() gopurs_runtime.Value {
	once_Data_Maybe_genericMaybe.Do(func() {
		cache_Data_Maybe_genericMaybe = gopurs_runtime.Value{Type: 9, IntVal: 1921946594, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_732223057_2818661616((&Constructor_Data_Generic_Rep_Generic[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
var __t_tag_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](x_0)
if (__t_tag_0 == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 3478632216, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_497217223_1323331594((&Constructor_Data_Generic_Rep_Inl[uint32, gopurs_runtime.Value]{1, 1454898258})))}
goto end_branch_2
} else {

}
}
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](x_0)
if (__t_tag_1 != nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 492034566, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_660678937_2687169876((&Constructor_Data_Generic_Rep_Inr[uint32, gopurs_runtime.Value]{1, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(x_0.UnsafePtr).V0})))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (x_0.Type == 9 && x_0.IntVal == 3478632216) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))})
goto end_branch_3
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 492034566) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{(*Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value])(x_0.UnsafePtr).V0, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_3
} else {

}
}
{
__t3 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)}
})})))}
	})
	return cache_Data_Maybe_genericMaybe
}

var cache_Data_Maybe_functorMaybe gopurs_runtime.Value
var once_Data_Maybe_functorMaybe sync.Once
func Get_Data_Maybe_functorMaybe() gopurs_runtime.Value {
	once_Data_Maybe_functorMaybe.Do(func() {
		cache_Data_Maybe_functorMaybe = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_3689823567_2812149806((&Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_1)
if (__t_tag_0 != nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply(v_0, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v1_1.UnsafePtr).V0), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
})})))}
	})
	return cache_Data_Maybe_functorMaybe
}

var cache_Data_Maybe_invariantMaybe gopurs_runtime.Value
var once_Data_Maybe_invariantMaybe sync.Once
func Get_Data_Maybe_invariantMaybe() gopurs_runtime.Value {
	once_Data_Maybe_invariantMaybe.Do(func() {
		cache_Data_Maybe_invariantMaybe = gopurs_runtime.Value{Type: 9, IntVal: 2396985522, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_3818752775_2241343270((&Constructor_Data_Functor_Invariant_Invariant[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_2)
if (__t_tag_0 != nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply(f_0, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v1_2.UnsafePtr).V0), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
})})))}
	})
	return cache_Data_Maybe_invariantMaybe
}

var cache_Data_Maybe_fromMaybe_prime_ gopurs_runtime.Value
var once_Data_Maybe_fromMaybe_prime_ sync.Once
func Get_Data_Maybe_fromMaybe_prime_() gopurs_runtime.Value {
	once_Data_Maybe_fromMaybe_prime_.Do(func() {
		cache_Data_Maybe_fromMaybe_prime_ = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_fromMaybe_prime_(a_0_box)
})
	})
	return cache_Data_Maybe_fromMaybe_prime_
}

var cache_Data_Maybe_fromMaybe gopurs_runtime.Value
var once_Data_Maybe_fromMaybe sync.Once
func Get_Data_Maybe_fromMaybe() gopurs_runtime.Value {
	once_Data_Maybe_fromMaybe.Do(func() {
		cache_Data_Maybe_fromMaybe = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_fromMaybe(a_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_1_box))
})
	})
	return cache_Data_Maybe_fromMaybe
}

var cache_Data_Maybe_fromMaybe__2288822913 gopurs_runtime.Value
var once_Data_Maybe_fromMaybe__2288822913 sync.Once
func Get_Data_Maybe_fromMaybe__2288822913() gopurs_runtime.Value {
	once_Data_Maybe_fromMaybe__2288822913.Do(func() {
		cache_Data_Maybe_fromMaybe__2288822913 = gopurs_runtime.Func2(func(a_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Maybe_fromMaybe__2288822913(uint32(a_unused_0_box.IntVal), gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[uint32]](__eta_norm_0_1_box))), UnsafePtr: nil}
})
	})
	return cache_Data_Maybe_fromMaybe__2288822913
}

var cache_Data_Maybe_fromMaybe__1106780484 gopurs_runtime.Value
var once_Data_Maybe_fromMaybe__1106780484 sync.Once
func Get_Data_Maybe_fromMaybe__1106780484() gopurs_runtime.Value {
	once_Data_Maybe_fromMaybe__1106780484.Do(func() {
		cache_Data_Maybe_fromMaybe__1106780484 = gopurs_runtime.Func2(func(a_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Maybe_fromMaybe__1106780484(uint32(a_unused_0_box.IntVal), gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[uint32]](__eta_norm_0_1_box))), UnsafePtr: nil}
})
	})
	return cache_Data_Maybe_fromMaybe__1106780484
}

var cache_Data_Maybe_fromMaybe__521152297 gopurs_runtime.Value
var once_Data_Maybe_fromMaybe__521152297 sync.Once
func Get_Data_Maybe_fromMaybe__521152297() gopurs_runtime.Value {
	once_Data_Maybe_fromMaybe__521152297.Do(func() {
		cache_Data_Maybe_fromMaybe__521152297 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Maybe_fromMaybe__521152297(a_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](__eta_norm_0_1_box)))
})
	})
	return cache_Data_Maybe_fromMaybe__521152297
}

var cache_Data_Maybe_fromJust gopurs_runtime.Value
var once_Data_Maybe_fromJust sync.Once
func Get_Data_Maybe_fromJust() gopurs_runtime.Value {
	once_Data_Maybe_fromJust.Do(func() {
		cache_Data_Maybe_fromJust = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_fromJust(_dollar___unused_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_1_box))
})
	})
	return cache_Data_Maybe_fromJust
}

var cache_Data_Maybe_fromJust__617296645 gopurs_runtime.Value
var once_Data_Maybe_fromJust__617296645 sync.Once
func Get_Data_Maybe_fromJust__617296645() gopurs_runtime.Value {
	once_Data_Maybe_fromJust__617296645.Do(func() {
		cache_Data_Maybe_fromJust__617296645 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Maybe_fromJust__617296645(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[uint32]](v_0_box))), UnsafePtr: nil}
})
	})
	return cache_Data_Maybe_fromJust__617296645
}

var cache_Data_Maybe_fromJust__3553976581 gopurs_runtime.Value
var once_Data_Maybe_fromJust__3553976581 sync.Once
func Get_Data_Maybe_fromJust__3553976581() gopurs_runtime.Value {
	once_Data_Maybe_fromJust__3553976581.Do(func() {
		cache_Data_Maybe_fromJust__3553976581 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Maybe_fromJust__3553976581(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[uint32]](v_0_box))), UnsafePtr: nil}
})
	})
	return cache_Data_Maybe_fromJust__3553976581
}

var cache_Data_Maybe_fromJust__3582419621 gopurs_runtime.Value
var once_Data_Maybe_fromJust__3582419621 sync.Once
func Get_Data_Maybe_fromJust__3582419621() gopurs_runtime.Value {
	once_Data_Maybe_fromJust__3582419621.Do(func() {
		cache_Data_Maybe_fromJust__3582419621 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(Call_Data_Maybe_fromJust__3582419621(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]](v_0_box)))}
})
	})
	return cache_Data_Maybe_fromJust__3582419621
}

var cache_Data_Maybe_fromJust__1283891301 gopurs_runtime.Value
var once_Data_Maybe_fromJust__1283891301 sync.Once
func Get_Data_Maybe_fromJust__1283891301() gopurs_runtime.Value {
	once_Data_Maybe_fromJust__1283891301.Do(func() {
		cache_Data_Maybe_fromJust__1283891301 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(Call_Data_Maybe_fromJust__1283891301(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Time_Time]](v_0_box)))}
})
	})
	return cache_Data_Maybe_fromJust__1283891301
}

var cache_Data_Maybe_fromJust__3527346789 gopurs_runtime.Value
var once_Data_Maybe_fromJust__3527346789 sync.Once
func Get_Data_Maybe_fromJust__3527346789() gopurs_runtime.Value {
	once_Data_Maybe_fromJust__3527346789.Do(func() {
		cache_Data_Maybe_fromJust__3527346789 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_Data_Maybe_fromJust__3527346789(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[[]string]](v_0_box))
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Str(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_Data_Maybe_fromJust__3527346789
}

var cache_Data_Maybe_fromJust__3648713157 gopurs_runtime.Value
var once_Data_Maybe_fromJust__3648713157 sync.Once
func Get_Data_Maybe_fromJust__3648713157() gopurs_runtime.Value {
	once_Data_Maybe_fromJust__3648713157.Do(func() {
		cache_Data_Maybe_fromJust__3648713157 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_Data_Maybe_fromJust__3648713157(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[[]int64]](v_0_box))
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_Data_Maybe_fromJust__3648713157
}

var cache_Data_Maybe_fromJust__1645345861 gopurs_runtime.Value
var once_Data_Maybe_fromJust__1645345861 sync.Once
func Get_Data_Maybe_fromJust__1645345861() gopurs_runtime.Value {
	once_Data_Maybe_fromJust__1645345861.Do(func() {
		cache_Data_Maybe_fromJust__1645345861 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Maybe_fromJust__1645345861(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](v_0_box)))
})
	})
	return cache_Data_Maybe_fromJust__1645345861
}

var cache_Data_Maybe_fromJust__1901609285 gopurs_runtime.Value
var once_Data_Maybe_fromJust__1901609285 sync.Once
func Get_Data_Maybe_fromJust__1901609285() gopurs_runtime.Value {
	once_Data_Maybe_fromJust__1901609285.Do(func() {
		cache_Data_Maybe_fromJust__1901609285 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_Maybe_fromJust__1901609285(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[string]](v_0_box)))
})
	})
	return cache_Data_Maybe_fromJust__1901609285
}

var cache_Data_Maybe_extendMaybe gopurs_runtime.Value
var once_Data_Maybe_extendMaybe sync.Once
func Get_Data_Maybe_extendMaybe() gopurs_runtime.Value {
	once_Data_Maybe_extendMaybe.Do(func() {
		cache_Data_Maybe_extendMaybe = gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_570732504_3290176857((&Constructor_Control_Extend_Extend[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_3689823567_2812149806(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Maybe_functorMaybe())))}
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_1)
if (__t_tag_0 == nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply(v_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_1))}), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
})})))}
	})
	return cache_Data_Maybe_extendMaybe
}

var cache_Data_Maybe_eqMaybe gopurs_runtime.Value
var once_Data_Maybe_eqMaybe sync.Once
func Get_Data_Maybe_eqMaybe() gopurs_runtime.Value {
	once_Data_Maybe_eqMaybe.Do(func() {
		cache_Data_Maybe_eqMaybe = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_eqMaybe(dictEq_0_box)
})
	})
	return cache_Data_Maybe_eqMaybe
}

var cache_Data_Maybe_ordMaybe gopurs_runtime.Value
var once_Data_Maybe_ordMaybe sync.Once
func Get_Data_Maybe_ordMaybe() gopurs_runtime.Value {
	once_Data_Maybe_ordMaybe.Do(func() {
		cache_Data_Maybe_ordMaybe = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_ordMaybe(dictOrd_0_box)
})
	})
	return cache_Data_Maybe_ordMaybe
}

var cache_Data_Maybe_eq1Maybe gopurs_runtime.Value
var once_Data_Maybe_eq1Maybe sync.Once
func Get_Data_Maybe_eq1Maybe() gopurs_runtime.Value {
	once_Data_Maybe_eq1Maybe.Do(func() {
		cache_Data_Maybe_eq1Maybe = gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_1662389854_1766074591((&Constructor_Data_Eq_Eq1[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func3(func(dictEq_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
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
	})
	return cache_Data_Maybe_eq1Maybe
}

var cache_Data_Maybe_ord1Maybe gopurs_runtime.Value
var once_Data_Maybe_ord1Maybe sync.Once
func Get_Data_Maybe_ord1Maybe() gopurs_runtime.Value {
	once_Data_Maybe_ord1Maybe.Do(func() {
		cache_Data_Maybe_ord1Maybe = gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_170825214_3985601471((&Constructor_Data_Ord_Ord1[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_1662389854_1766074591(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Maybe_eq1Maybe())))}
}), gopurs_runtime.Func3(func(dictOrd_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 uint32
{
var __t_tag_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](x_1)
if (__t_tag_0 == nil) {
var __t2 uint32
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](y_2)
if (__t_tag_1 == nil) {
__t2 = 902936544
goto end_branch_2
} else {

}
}
{
__t2 = 1527465420
}
end_branch_2:
__t7 = __t2
goto end_branch_7
} else {

}
}
{
var __t_tag_3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](y_2)
if (__t_tag_3 == nil) {
__t7 = 380165415
goto end_branch_7
} else {

}
}
{
var __t_tag_4 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](x_1)
var __t_and_6 bool = false
if (__t_tag_4 != nil) {

var __t_tag_5 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](y_2)
__t_and_6 = (__t_tag_5 != nil)
}
if __t_and_6 {
__t7 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(x_1.UnsafePtr).V0, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(y_2.UnsafePtr).V0).IntVal)
goto end_branch_7
} else {

}
}
{
__t7 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_7:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t7), UnsafePtr: nil}
})})))}
	})
	return cache_Data_Maybe_ord1Maybe
}

var cache_Data_Maybe_boundedMaybe gopurs_runtime.Value
var once_Data_Maybe_boundedMaybe sync.Once
func Get_Data_Maybe_boundedMaybe() gopurs_runtime.Value {
	once_Data_Maybe_boundedMaybe.Do(func() {
		cache_Data_Maybe_boundedMaybe = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_boundedMaybe(dictBounded_0_box)
})
	})
	return cache_Data_Maybe_boundedMaybe
}

var cache_Data_Maybe_applyMaybe gopurs_runtime.Value
var once_Data_Maybe_applyMaybe sync.Once
func Get_Data_Maybe_applyMaybe() gopurs_runtime.Value {
	once_Data_Maybe_applyMaybe.Do(func() {
		cache_Data_Maybe_applyMaybe = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_3552963512_3741347833((&Constructor_Control_Apply_Apply[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_3689823567_2812149806(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Maybe_functorMaybe())))}
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_0)
if (__t_tag_0 != nil) {
var __t2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_1)
if (__t_tag_1 != nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_0.UnsafePtr).V0, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v1_1.UnsafePtr).V0), true}
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
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_2:
__t4 = __t2
goto end_branch_4
} else {

}
}
{
var __t_tag_3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_0)
if (__t_tag_3 == nil) {
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_4
} else {

}
}
{
__t4 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t4)}
})})))}
	})
	return cache_Data_Maybe_applyMaybe
}

var cache_Data_Maybe_bindMaybe gopurs_runtime.Value
var once_Data_Maybe_bindMaybe sync.Once
func Get_Data_Maybe_bindMaybe() gopurs_runtime.Value {
	once_Data_Maybe_bindMaybe.Do(func() {
		cache_Data_Maybe_bindMaybe = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_3887487416_2748095225((&Constructor_Control_Bind_Bind[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_3552963512_3741347833(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Maybe_applyMaybe())))}
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_0)
if (__t_tag_0 != nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(v1_1, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_0.UnsafePtr).V0))
goto end_branch_2
} else {

}
}
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_0)
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
})})))}
	})
	return cache_Data_Maybe_bindMaybe
}

var cache_Data_Maybe_semiringMaybe gopurs_runtime.Value
var once_Data_Maybe_semiringMaybe sync.Once
func Get_Data_Maybe_semiringMaybe() gopurs_runtime.Value {
	once_Data_Maybe_semiringMaybe.Do(func() {
		cache_Data_Maybe_semiringMaybe = gopurs_runtime.Func(func(dictSemiring_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_semiringMaybe(dictSemiring_0_box)
})
	})
	return cache_Data_Maybe_semiringMaybe
}

var cache_Data_Maybe_applicativeMaybe gopurs_runtime.Value
var once_Data_Maybe_applicativeMaybe sync.Once
func Get_Data_Maybe_applicativeMaybe() gopurs_runtime.Value {
	once_Data_Maybe_applicativeMaybe.Do(func() {
		cache_Data_Maybe_applicativeMaybe = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_649684152_1439734649((&Constructor_Control_Applicative_Applicative[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_3552963512_3741347833(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Maybe_applyMaybe())))}
}), Get_Data_Maybe_Just()})))}
	})
	return cache_Data_Maybe_applicativeMaybe
}

var cache_Data_Maybe_monadMaybe gopurs_runtime.Value
var once_Data_Maybe_monadMaybe sync.Once
func Get_Data_Maybe_monadMaybe() gopurs_runtime.Value {
	once_Data_Maybe_monadMaybe.Do(func() {
		cache_Data_Maybe_monadMaybe = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_1642601656_2568689657((&Constructor_Control_Monad_Monad[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_649684152_1439734649(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Maybe_applicativeMaybe())))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_3887487416_2748095225(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Maybe_bindMaybe())))}
})})))}
	})
	return cache_Data_Maybe_monadMaybe
}

var cache_Data_Maybe_altMaybe gopurs_runtime.Value
var once_Data_Maybe_altMaybe sync.Once
func Get_Data_Maybe_altMaybe() gopurs_runtime.Value {
	once_Data_Maybe_altMaybe.Do(func() {
		cache_Data_Maybe_altMaybe = gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_3904200120_3421983481((&Constructor_Control_Alt_Alt[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_3689823567_2812149806(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Maybe_functorMaybe())))}
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_0)
if (__t_tag_0 == nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_1)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_0)
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
})})))}
	})
	return cache_Data_Maybe_altMaybe
}

var cache_Data_Maybe_plusMaybe gopurs_runtime.Value
var once_Data_Maybe_plusMaybe sync.Once
func Get_Data_Maybe_plusMaybe() gopurs_runtime.Value {
	once_Data_Maybe_plusMaybe.Do(func() {
		cache_Data_Maybe_plusMaybe = gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_267397720_3706288089((&Constructor_Control_Plus_Plus[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_3904200120_3421983481(gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Maybe_altMaybe())))}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}})))}
	})
	return cache_Data_Maybe_plusMaybe
}

var cache_Data_Maybe_alternativeMaybe gopurs_runtime.Value
var once_Data_Maybe_alternativeMaybe sync.Once
func Get_Data_Maybe_alternativeMaybe() gopurs_runtime.Value {
	once_Data_Maybe_alternativeMaybe.Do(func() {
		cache_Data_Maybe_alternativeMaybe = gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_673668088_2307501113((&Constructor_Control_Alternative_Alternative[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_649684152_1439734649(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Maybe_applicativeMaybe())))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_267397720_3706288089(gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Maybe_plusMaybe())))}
})})))}
	})
	return cache_Data_Maybe_alternativeMaybe
}

type Constructor_Data_Maybe_Nothing[T_a any] struct {
	Rc uint32
}


type Constructor_Data_Maybe_Just[T_a any] struct {
	Rc uint32
	V0 T_a
}


func Call_Data_Maybe_Just__3521044077(__eta_norm_0_unused_0_loop uint32) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__3521044077:
for {
if false { continue Just__3521044077 }
var __eta_norm_0_unused_0 uint32 = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(1692989816), UnsafePtr: nil}, true}
}
}

func Call_Data_Maybe_Just__2798349578(__eta_norm_0_unused_0_loop uint32) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__2798349578:
for {
if false { continue Just__2798349578 }
var __eta_norm_0_unused_0 uint32 = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(243771071), UnsafePtr: nil}, true}
}
}

func Call_Data_Maybe_Just__2825732244(__eta_norm_0_unused_0_loop uint32) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__2825732244:
for {
if false { continue Just__2825732244 }
var __eta_norm_0_unused_0 uint32 = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(3889233761), UnsafePtr: nil}, true}
}
}

func Call_Data_Maybe_Just__3832816551(__eta_norm_0_unused_0_loop uint32) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__3832816551:
for {
if false { continue Just__3832816551 }
var __eta_norm_0_unused_0 uint32 = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(2455627378), UnsafePtr: nil}, true}
}
}

func Call_Data_Maybe_Just__4271887697(__eta_norm_0_unused_0_loop uint32) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__4271887697:
for {
if false { continue Just__4271887697 }
var __eta_norm_0_unused_0 uint32 = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(1908470532), UnsafePtr: nil}, true}
}
}

func Call_Data_Maybe_Just__1091121345(__eta_norm_0_unused_0_loop uint32) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__1091121345:
for {
if false { continue Just__1091121345 }
var __eta_norm_0_unused_0 uint32 = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(2276710548), UnsafePtr: nil}, true}
}
}

func Call_Data_Maybe_Just__4216227551(__eta_norm_0_unused_0_loop uint32) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__4216227551:
for {
if false { continue Just__4216227551 }
var __eta_norm_0_unused_0 uint32 = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(4067355978), UnsafePtr: nil}, true}
}
}

func Call_Data_Maybe_Just__3031626846(__eta_norm_0_unused_0_loop uint32) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__3031626846:
for {
if false { continue Just__3031626846 }
var __eta_norm_0_unused_0 uint32 = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(4162469099), UnsafePtr: nil}, true}
}
}

func Call_Data_Maybe_Just__1154318014(__eta_norm_0_unused_0_loop uint32) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__1154318014:
for {
if false { continue Just__1154318014 }
var __eta_norm_0_unused_0 uint32 = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(330658827), UnsafePtr: nil}, true}
}
}

func Call_Data_Maybe_Just__1938051713(__eta_norm_0_unused_0_loop uint32) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__1938051713:
for {
if false { continue Just__1938051713 }
var __eta_norm_0_unused_0 uint32 = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(49471444), UnsafePtr: nil}, true}
}
}

func Call_Data_Maybe_Just__3180259945(__eta_norm_0_unused_0_loop uint32) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__3180259945:
for {
if false { continue Just__3180259945 }
var __eta_norm_0_unused_0 uint32 = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(8639228), UnsafePtr: nil}, true}
}
}

func Call_Data_Maybe_Just__2957110564(__eta_norm_0_unused_0_loop uint32) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__2957110564:
for {
if false { continue Just__2957110564 }
var __eta_norm_0_unused_0 uint32 = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(215731793), UnsafePtr: nil}, true}
}
}

func Call_Data_Maybe_Just__3561111658(__eta_norm_0_unused_0_loop uint32) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__3561111658:
for {
if false { continue Just__3561111658 }
var __eta_norm_0_unused_0 uint32 = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(2946274527), UnsafePtr: nil}, true}
}
}

func Call_Data_Maybe_Just__3347273307(__eta_norm_0_unused_0_loop uint32) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__3347273307:
for {
if false { continue Just__3347273307 }
var __eta_norm_0_unused_0 uint32 = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(2900196686), UnsafePtr: nil}, true}
}
}

func Call_Data_Maybe_Just__2779760438(__eta_norm_0_unused_0_loop uint32) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__2779760438:
for {
if false { continue Just__2779760438 }
var __eta_norm_0_unused_0 uint32 = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(1070786179), UnsafePtr: nil}, true}
}
}

func Call_Data_Maybe_Just__107066655(__eta_norm_0_unused_0_loop uint32) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__107066655:
for {
if false { continue Just__107066655 }
var __eta_norm_0_unused_0 uint32 = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(1326716170), UnsafePtr: nil}, true}
}
}

func Call_Data_Maybe_Just__508806911(__eta_norm_0_unused_0_loop uint32) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__508806911:
for {
if false { continue Just__508806911 }
var __eta_norm_0_unused_0 uint32 = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(3818857258), UnsafePtr: nil}, true}
}
}

func Call_Data_Maybe_Just__1678136480(__eta_norm_0_unused_0_loop uint32) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__1678136480:
for {
if false { continue Just__1678136480 }
var __eta_norm_0_unused_0 uint32 = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(20457557), UnsafePtr: nil}, true}
}
}

func Call_Data_Maybe_Just__2937552249(__eta_norm_0_unused_0_loop uint32) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__2937552249:
for {
if false { continue Just__2937552249 }
var __eta_norm_0_unused_0 uint32 = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(4227105004), UnsafePtr: nil}, true}
}
}

func Call_Data_Maybe_Just__3996651813(__eta_norm_0_0_loop *Constructor_Data_Date_Date) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__3996651813:
for {
if false { continue Just__3996651813 }
var __eta_norm_0_0 *Constructor_Data_Date_Date = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(__eta_norm_0_0)}, true}
}
}

func Call_Data_Maybe_Just__4227547111(__eta_norm_0_unused_0_loop uint32) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__4227547111:
for {
if false { continue Just__4227547111 }
var __eta_norm_0_unused_0 uint32 = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(1454898258), UnsafePtr: nil}, true}
}
}

func Call_Data_Maybe_Just__3629863589(__eta_norm_0_0_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__3629863589:
for {
if false { continue Just__3629863589 }
var __eta_norm_0_0 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_3123684004_1293498952(__eta_norm_0_0))}, true}
}
}

func Call_Data_Maybe_Just__2347572452(__eta_norm_0_unused_0_loop uint32) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__2347572452:
for {
if false { continue Just__2347572452 }
var __eta_norm_0_unused_0 uint32 = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, true}
}
}

func Call_Data_Maybe_Just__920868163(__eta_norm_0_unused_0_loop uint32) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__920868163:
for {
if false { continue Just__920868163 }
var __eta_norm_0_unused_0 uint32 = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, true}
}
}

func Call_Data_Maybe_Just__113341384(__eta_norm_0_unused_0_loop uint32) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__113341384:
for {
if false { continue Just__113341384 }
var __eta_norm_0_unused_0 uint32 = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, true}
}
}

func Call_Data_Maybe_Just__3120863845(__eta_norm_0_0_loop *Constructor_Data_Tuple_Tuple[int64, int64]) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__3120863845:
for {
if false { continue Just__3120863845 }
var __eta_norm_0_0 *Constructor_Data_Tuple_Tuple[int64, int64] = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_3363075976_138441832(__eta_norm_0_0))}, true}
}
}

func Call_Data_Maybe_Just__3356776037(__eta_norm_0_0_loop bool) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__3356776037:
for {
if false { continue Just__3356776037 }
var __eta_norm_0_0 bool = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Bool(__eta_norm_0_0), true}
}
}

func Call_Data_Maybe_Just__2827584037(__eta_norm_0_0_loop string) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__2827584037:
for {
if false { continue Just__2827584037 }
var __eta_norm_0_0 string = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Str(__eta_norm_0_0), true}
}
}

func Call_Data_Maybe_Just__983269893(__eta_norm_0_0_loop int64) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__983269893:
for {
if false { continue Just__983269893 }
var __eta_norm_0_0 int64 = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Int(__eta_norm_0_0), true}
}
}

func Call_Data_Maybe_Just__157193253(__eta_norm_0_0_loop float64) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__157193253:
for {
if false { continue Just__157193253 }
var __eta_norm_0_0 float64 = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Float(__eta_norm_0_0), true}
}
}

func Call_Data_Maybe_Just__436376869(__eta_norm_0_0_loop string) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__436376869:
for {
if false { continue Just__436376869 }
var __eta_norm_0_0 string = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Str(__eta_norm_0_0), true}
}
}

func Call_Data_Maybe_Just__4131531262(__eta_norm_0_unused_0_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__4131531262:
for {
if false { continue Just__4131531262 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return struct{V0 gopurs_runtime.Value; V1 bool}{Get_Data_Unit_unit(), true}
}
}

func Call_Data_Maybe_showMaybe(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_345859663_1386611502((&Constructor_Data_Show_Show[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 string
{
var __t_tag_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_1)
if (__t_tag_0 != nil) {
__t2 = (("(Just ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_1.UnsafePtr).V0).StrVal())) + (")")
goto end_branch_2
} else {

}
}
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_1)
if (__t_tag_1 == nil) {
__t2 = "Nothing"
goto end_branch_2
} else {

}
}
{
__t2 = func() string { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Str(__t2)
})})))}
}

func Call_Data_Maybe_semigroupMaybe(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_125234255_4179793454((&Constructor_Data_Semigroup_Semigroup[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_1)
if (__t_tag_0 == nil) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_2)
goto end_branch_5
} else {

}
}
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_2)
if (__t_tag_1 == nil) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_1)
goto end_branch_5
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_1)
var __t_and_4 bool = false
if (__t_tag_2 != nil) {

var __t_tag_3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_2)
__t_and_4 = (__t_tag_3 != nil)
}
if __t_and_4 {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_1.UnsafePtr).V0, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v1_2.UnsafePtr).V0), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_5
} else {

}
}
{
__t5 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t5)}
})})))}
}

func Call_Data_Maybe_optional(dictAlt_0_loop *Constructor_Control_Alt_Alt[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictAlt_0 *Constructor_Control_Alt_Alt[gopurs_runtime.Value] = dictAlt_0_loop
_ = dictAlt_0
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictAlt_0.V0), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func2(func(dictApplicative_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictAlt_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Maybe_Just(), a_3), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}))
})
}

func Call_Data_Maybe_monoidMaybe(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
// TAST (Let): semigroupMaybe1_1_0 shape=LitRecord bindingType=(ADT ["Data","Semigroup","Semigroup"] [(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])])
semigroupMaybe1_1_0 := (&Constructor_Data_Semigroup_Semigroup[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_1)
if (__t_tag_1 == nil) {
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_2)
goto end_branch_6
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_2)
if (__t_tag_2 == nil) {
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_1)
goto end_branch_6
} else {

}
}
{
var __t_tag_3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_1)
var __t_and_5 bool = false
if (__t_tag_3 != nil) {

var __t_tag_4 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_2)
__t_and_5 = (__t_tag_4 != nil)
}
if __t_and_5 {
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_1.UnsafePtr).V0, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v1_2.UnsafePtr).V0), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_6
} else {

}
}
{
__t6 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t6)}
})})
_ = semigroupMaybe1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_1089607855_1201789390((&Constructor_Data_Monoid_Monoid[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_125234255_4179793454(semigroupMaybe1_1_0))}
}), gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))})})))}
}

func Call_Data_Maybe_maybe_prime_(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2 == nil) {
__t0 = gopurs_runtime.Apply(v_0, Get_Data_Unit_unit())
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (v2_2).V0)
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

func Call_Data_Maybe_maybe(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2 == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (v2_2).V0)
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

func Call_Data_Maybe_maybe__4101309216(v_unused_0_loop uint32, v1_unused_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just[uint32]) uint32 {
maybe__4101309216:
for {
if false { continue maybe__4101309216 }
var v_unused_0 uint32 = v_unused_0_loop
_ = v_unused_0
var v1_unused_1 gopurs_runtime.Value = v1_unused_1_loop
_ = v1_unused_1
var v2_2 *Constructor_Data_Maybe_Just[uint32] = v2_2_loop
_ = v2_2
var __t0 uint32
{
if (v2_2 == nil) {
__t0 = 3889233761
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = uint32(gopurs_runtime.Value{Type: 9, IntVal: int64((v2_2).V0), UnsafePtr: nil}.IntVal)
goto end_branch_0
} else {

}
}
{
__t0 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}
}

func Call_Data_Maybe_maybe__3207254341(v_unused_0_loop uint32, v1_unused_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just[uint32]) uint32 {
maybe__3207254341:
for {
if false { continue maybe__3207254341 }
var v_unused_0 uint32 = v_unused_0_loop
_ = v_unused_0
var v1_unused_1 gopurs_runtime.Value = v1_unused_1_loop
_ = v1_unused_1
var v2_2 *Constructor_Data_Maybe_Just[uint32] = v2_2_loop
_ = v2_2
var __t0 uint32
{
if (v2_2 == nil) {
__t0 = 1908470532
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = uint32(gopurs_runtime.Value{Type: 9, IntVal: int64((v2_2).V0), UnsafePtr: nil}.IntVal)
goto end_branch_0
} else {

}
}
{
__t0 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}
}

func Call_Data_Maybe_maybe__539590268(v_0_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value], v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 bool} {
maybe__539590268:
for {
if false { continue maybe__539590268 }
var v_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (v2_2 == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(v1_1, (v2_2).V0))
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

func Call_Data_Maybe_maybe__17375519(v_0_loop bool, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just[uint32]) bool {
maybe__17375519:
for {
if false { continue maybe__17375519 }
var v_0 bool = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just[uint32] = v2_2_loop
_ = v2_2
var __t0 bool
{
if (v2_2 == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = (gopurs_runtime.Apply(v1_1, gopurs_runtime.Value{Type: 9, IntVal: int64((v2_2).V0), UnsafePtr: nil}).IntVal) != (0)
goto end_branch_0
} else {

}
}
{
__t0 = func() bool { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}
}

func Call_Data_Maybe_maybe__3303070975(v_0_loop bool, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just[int64]) bool {
maybe__3303070975:
for {
if false { continue maybe__3303070975 }
var v_0 bool = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just[int64] = v2_2_loop
_ = v2_2
var __t0 bool
{
if (v2_2 == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = (gopurs_runtime.Apply(v1_1, gopurs_runtime.Int((v2_2).V0)).IntVal) != (0)
goto end_branch_0
} else {

}
}
{
__t0 = func() bool { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}
}

func Call_Data_Maybe_maybe__1344312831(v_0_loop bool, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just[float64]) bool {
maybe__1344312831:
for {
if false { continue maybe__1344312831 }
var v_0 bool = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just[float64] = v2_2_loop
_ = v2_2
var __t0 bool
{
if (v2_2 == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = (gopurs_runtime.Apply(v1_1, gopurs_runtime.Float((v2_2).V0)).IntVal) != (0)
goto end_branch_0
} else {

}
}
{
__t0 = func() bool { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}
}

func Call_Data_Maybe_maybe__774193768(v_0_loop int64, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just[int64]) int64 {
maybe__774193768:
for {
if false { continue maybe__774193768 }
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just[int64] = v2_2_loop
_ = v2_2
var __t0 int64
{
if (v2_2 == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = gopurs_runtime.Apply(v1_1, gopurs_runtime.Int((v2_2).V0)).IntVal
goto end_branch_0
} else {

}
}
{
__t0 = func() int64 { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}
}

func Call_Data_Maybe_maybe__4030275671(v1_unused_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_Maybe_Just[int64], v2_2_loop gopurs_runtime.Value) int64 {
maybe__4030275671:
for {
if false { continue maybe__4030275671 }
var v1_unused_0 gopurs_runtime.Value = v1_unused_0_loop
_ = v1_unused_0
var v_1 *Constructor_Data_Maybe_Just[int64] = v_1_loop
_ = v_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t2 gopurs_runtime.Value
{
var __t_tag_0 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_Maybe_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_2))
if (__t_tag_0 == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_1170268447_3094389156(v_1))}
goto end_branch_2
} else {

}
}
{
var __t_tag_1 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_Maybe_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_2))
if (__t_tag_1 != nil) {
__t2 = gopurs_runtime.Int((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v2_2.UnsafePtr).V0.IntVal)
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Int(func() int64 { panic("Failed pattern match") }())
}
end_branch_2:
return __t2.IntVal
}
}

func Call_Data_Maybe_isNothing(v2_0_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) bool {
var v2_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (v2_0 == nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
if (v2_0 != nil) {
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

func Call_Data_Maybe_isNothing__2526948203(__eta_norm_0_0_loop *Constructor_Data_Maybe_Just[uint32]) bool {
isNothing__2526948203:
for {
if false { continue isNothing__2526948203 }
var __eta_norm_0_0 *Constructor_Data_Maybe_Just[uint32] = __eta_norm_0_0_loop
_ = __eta_norm_0_0
var __t0 gopurs_runtime.Value
{
if (__eta_norm_0_0 == nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
if (__eta_norm_0_0 != nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return (gopurs_runtime.Bool((__t0.IntVal) != (0)).IntVal) != (0)
}
}

func Call_Data_Maybe_isNothing__684798802(__eta_norm_0_0_loop *Constructor_Data_Maybe_Just[int64]) bool {
isNothing__684798802:
for {
if false { continue isNothing__684798802 }
var __eta_norm_0_0 *Constructor_Data_Maybe_Just[int64] = __eta_norm_0_0_loop
_ = __eta_norm_0_0
var __t0 gopurs_runtime.Value
{
if (__eta_norm_0_0 == nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
if (__eta_norm_0_0 != nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return (gopurs_runtime.Bool((__t0.IntVal) != (0)).IntVal) != (0)
}
}

func Call_Data_Maybe_isJust(v2_0_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) bool {
var v2_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (v2_0 == nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
if (v2_0 != nil) {
__t0 = gopurs_runtime.Bool(true)
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

func Call_Data_Maybe_isJust__684798802(__eta_norm_0_0_loop *Constructor_Data_Maybe_Just[int64]) bool {
isJust__684798802:
for {
if false { continue isJust__684798802 }
var __eta_norm_0_0 *Constructor_Data_Maybe_Just[int64] = __eta_norm_0_0_loop
_ = __eta_norm_0_0
var __t0 gopurs_runtime.Value
{
if (__eta_norm_0_0 == nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
if (__eta_norm_0_0 != nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return (gopurs_runtime.Bool((__t0.IntVal) != (0)).IntVal) != (0)
}
}

func Call_Data_Maybe_isJust__912756578(__eta_norm_0_0_loop *Constructor_Data_Maybe_Just[float64]) bool {
isJust__912756578:
for {
if false { continue isJust__912756578 }
var __eta_norm_0_0 *Constructor_Data_Maybe_Just[float64] = __eta_norm_0_0_loop
_ = __eta_norm_0_0
var __t0 gopurs_runtime.Value
{
if (__eta_norm_0_0 == nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
if (__eta_norm_0_0 != nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return (gopurs_runtime.Bool((__t0.IntVal) != (0)).IntVal) != (0)
}
}

func Call_Data_Maybe_fromMaybe_prime_(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return gopurs_runtime.Apply2(Get_Data_Maybe_maybe_prime_(), a_0, gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1))
}

func Call_Data_Maybe_fromMaybe(a_0_loop gopurs_runtime.Value, v2_1_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v2_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = v2_1_loop
_ = v2_1
var __t0 gopurs_runtime.Value
{
if (v2_1 == nil) {
__t0 = a_0
goto end_branch_0
} else {

}
}
{
if (v2_1 != nil) {
__t0 = (v2_1).V0
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

func Call_Data_Maybe_fromMaybe__2288822913(a_unused_0_loop uint32, __eta_norm_0_1_loop *Constructor_Data_Maybe_Just[uint32]) uint32 {
fromMaybe__2288822913:
for {
if false { continue fromMaybe__2288822913 }
var a_unused_0 uint32 = a_unused_0_loop
_ = a_unused_0
var __eta_norm_0_1 *Constructor_Data_Maybe_Just[uint32] = __eta_norm_0_1_loop
_ = __eta_norm_0_1
var __t0 gopurs_runtime.Value
{
if (__eta_norm_0_1 == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: int64(3889233761), UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if (__eta_norm_0_1 != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: int64((__eta_norm_0_1).V0), UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__t0.IntVal)), UnsafePtr: nil}.IntVal)
}
}

func Call_Data_Maybe_fromMaybe__1106780484(a_unused_0_loop uint32, __eta_norm_0_1_loop *Constructor_Data_Maybe_Just[uint32]) uint32 {
fromMaybe__1106780484:
for {
if false { continue fromMaybe__1106780484 }
var a_unused_0 uint32 = a_unused_0_loop
_ = a_unused_0
var __eta_norm_0_1 *Constructor_Data_Maybe_Just[uint32] = __eta_norm_0_1_loop
_ = __eta_norm_0_1
var __t0 gopurs_runtime.Value
{
if (__eta_norm_0_1 == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: int64(1908470532), UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if (__eta_norm_0_1 != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: int64((__eta_norm_0_1).V0), UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__t0.IntVal)), UnsafePtr: nil}.IntVal)
}
}

func Call_Data_Maybe_fromMaybe__521152297(a_0_loop int64, __eta_norm_0_1_loop *Constructor_Data_Maybe_Just[int64]) int64 {
fromMaybe__521152297:
for {
if false { continue fromMaybe__521152297 }
var a_0 int64 = a_0_loop
_ = a_0
var __eta_norm_0_1 *Constructor_Data_Maybe_Just[int64] = __eta_norm_0_1_loop
_ = __eta_norm_0_1
var __t0 gopurs_runtime.Value
{
if (__eta_norm_0_1 == nil) {
__t0 = gopurs_runtime.Int(a_0)
goto end_branch_0
} else {

}
}
{
if (__eta_norm_0_1 != nil) {
__t0 = gopurs_runtime.Int((__eta_norm_0_1).V0)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Int(__t0.IntVal).IntVal
}
}

func Call_Data_Maybe_fromJust(_dollar___unused_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
_ = _dollar___unused_0
var v_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = v_1_loop
_ = v_1
var __t0 gopurs_runtime.Value
{
if (v_1 != nil) {
__t0 = (v_1).V0
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

func Call_Data_Maybe_fromJust__617296645(v_0_loop *Constructor_Data_Maybe_Just[uint32]) uint32 {
fromJust__617296645:
for {
if false { continue fromJust__617296645 }
var v_0 *Constructor_Data_Maybe_Just[uint32] = v_0_loop
_ = v_0
var __t0 uint32
{
if (v_0 != nil) {
__t0 = (v_0).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_0:
return uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(__t0), UnsafePtr: nil}.IntVal)
}
}

func Call_Data_Maybe_fromJust__3553976581(v_0_loop *Constructor_Data_Maybe_Just[uint32]) uint32 {
fromJust__3553976581:
for {
if false { continue fromJust__3553976581 }
var v_0 *Constructor_Data_Maybe_Just[uint32] = v_0_loop
_ = v_0
var __t0 uint32
{
if (v_0 != nil) {
__t0 = (v_0).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_0:
return uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(__t0), UnsafePtr: nil}.IntVal)
}
}

func Call_Data_Maybe_fromJust__3582419621(v_0_loop *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]) *Constructor_Data_Date_Date {
fromJust__3582419621:
for {
if false { continue fromJust__3582419621 }
var v_0 *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date] = v_0_loop
_ = v_0
var __t0 *Constructor_Data_Date_Date
{
if (v_0 != nil) {
__t0 = (v_0).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() *Constructor_Data_Date_Date { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}
}

func Call_Data_Maybe_fromJust__1283891301(v_0_loop *Constructor_Data_Maybe_Just[*Constructor_Data_Time_Time]) *Constructor_Data_Time_Time {
fromJust__1283891301:
for {
if false { continue fromJust__1283891301 }
var v_0 *Constructor_Data_Maybe_Just[*Constructor_Data_Time_Time] = v_0_loop
_ = v_0
var __t0 *Constructor_Data_Time_Time
{
if (v_0 != nil) {
__t0 = (v_0).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() *Constructor_Data_Time_Time { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}
}

func Call_Data_Maybe_fromJust__3527346789(v_0_loop *Constructor_Data_Maybe_Just[[]string]) []string {
fromJust__3527346789:
for {
if false { continue fromJust__3527346789 }
var v_0 *Constructor_Data_Maybe_Just[[]string] = v_0_loop
_ = v_0
var __t0 []string
{
if (v_0 != nil) {
__t0 = (v_0).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() []string { panic("Failed pattern match") }()
}
end_branch_0:
return func() []string {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
					arr := __t0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Str(v) }
					return gopurs_runtime.Array(boxed)
				}().UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr { unboxed[i] = v.StrVal() }
					return unboxed
				}()
}
}

func Call_Data_Maybe_fromJust__3648713157(v_0_loop *Constructor_Data_Maybe_Just[[]int64]) []int64 {
fromJust__3648713157:
for {
if false { continue fromJust__3648713157 }
var v_0 *Constructor_Data_Maybe_Just[[]int64] = v_0_loop
_ = v_0
var __t0 []int64
{
if (v_0 != nil) {
__t0 = (v_0).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() []int64 { panic("Failed pattern match") }()
}
end_branch_0:
return func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
					arr := __t0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}().UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}()
}
}

func Call_Data_Maybe_fromJust__1645345861(v_0_loop *Constructor_Data_Maybe_Just[int64]) int64 {
fromJust__1645345861:
for {
if false { continue fromJust__1645345861 }
var v_0 *Constructor_Data_Maybe_Just[int64] = v_0_loop
_ = v_0
var __t0 int64
{
if (v_0 != nil) {
__t0 = (v_0).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() int64 { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Int(__t0).IntVal
}
}

func Call_Data_Maybe_fromJust__1901609285(v_0_loop *Constructor_Data_Maybe_Just[string]) string {
fromJust__1901609285:
for {
if false { continue fromJust__1901609285 }
var v_0 *Constructor_Data_Maybe_Just[string] = v_0_loop
_ = v_0
var __t0 string
{
if (v_0 != nil) {
__t0 = (v_0).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() string { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Str(__t0).StrVal()
}
}

func Call_Data_Maybe_eqMaybe(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_3508461103_3790796878((&Constructor_Data_Eq_Eq[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_Data_Maybe_ordMaybe(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_2155612431_4177771502((&Constructor_Data_Ord_Ord[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_3508461103_3790796878(eqMaybe1_1_0))}
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

func Call_Data_Maybe_boundedMaybe(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBounded_0, "Ord0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 shape=App(Other) bindingType=Any
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): eqMaybe1_2_2 shape=Let(LitRecord) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])])
eqMaybe1_2_2 := (&Constructor_Data_Eq_Eq[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(x_3 gopurs_runtime.Value, y_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 bool
{
var __t_tag_7 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](x_3)
if (__t_tag_7 == nil) {
var __t_tag_8 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](y_4)
__t9 = (__t_tag_8 == nil)
goto end_branch_9
} else {

}
}
{
var __t_tag_4 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](x_3)
var __t_and_6 bool = false
if (__t_tag_4 != nil) {

var __t_tag_5 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](y_4)
__t_and_6 = ((__t_tag_5 != nil)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_3, "eq"), (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(x_3.UnsafePtr).V0, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(y_4.UnsafePtr).V0).IntVal) != (0))
}
__t9 = __t_and_6
}
end_branch_9:
return gopurs_runtime.Bool(__t9)
})})
_ = eqMaybe1_2_2
// TAST (Let): ordMaybe1_1_0 shape=Let(Let(LitRecord)) bindingType=(ADT ["Data","Ord","Ord"] [(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])])
ordMaybe1_1_0 := (&Constructor_Data_Ord_Ord[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_3508461103_3790796878(eqMaybe1_2_2))}
}), gopurs_runtime.Func2(func(x_3 gopurs_runtime.Value, y_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t17 uint32
{
var __t_tag_10 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](x_3)
if (__t_tag_10 == nil) {
var __t12 uint32
{
var __t_tag_11 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](y_4)
if (__t_tag_11 == nil) {
__t12 = 902936544
goto end_branch_12
} else {

}
}
{
__t12 = 1527465420
}
end_branch_12:
__t17 = __t12
goto end_branch_17
} else {

}
}
{
var __t_tag_13 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](y_4)
if (__t_tag_13 == nil) {
__t17 = 380165415
goto end_branch_17
} else {

}
}
{
var __t_tag_14 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](x_3)
var __t_and_16 bool = false
if (__t_tag_14 != nil) {

var __t_tag_15 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](y_4)
__t_and_16 = (__t_tag_15 != nil)
}
if __t_and_16 {
__t17 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "compare"), (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(x_3.UnsafePtr).V0, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(y_4.UnsafePtr).V0).IntVal)
goto end_branch_17
} else {

}
}
{
__t17 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_17:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t17), UnsafePtr: nil}
})})
_ = ordMaybe1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_591333647_2094947566((&Constructor_Data_Bounded_Bounded[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_2155612431_4177771502(ordMaybe1_1_0))}
}), gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.RecordGet(dictBounded_0, "top"), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())})))}
}

func Call_Data_Maybe_semiringMaybe(dictSemiring_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemiring_0 gopurs_runtime.Value = dictSemiring_0_loop
_ = dictSemiring_0
return gopurs_runtime.Value{Type: 9, IntVal: 134961754, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_2047434479_2826095630((&Constructor_Data_Semiring_Semiring[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_1)
if (__t_tag_0 == nil) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_2)
goto end_branch_5
} else {

}
}
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_2)
if (__t_tag_1 == nil) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_1)
goto end_branch_5
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_1)
var __t_and_4 bool = false
if (__t_tag_2 != nil) {

var __t_tag_3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_2)
__t_and_4 = (__t_tag_3 != nil)
}
if __t_and_4 {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "add"), (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_1.UnsafePtr).V0, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v1_2.UnsafePtr).V0), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_5
} else {

}
}
{
__t5 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t5)}
}), gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_6 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](x_1)
if (__t_tag_6 != nil) {
var __t8 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_7 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](y_2)
if (__t_tag_7 != nil) {
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "mul"), (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(x_1.UnsafePtr).V0, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(y_2.UnsafePtr).V0), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_8:
__t9 = __t8
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_9:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t9)}
}), gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.RecordGet(dictSemiring_0, "one"), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()), gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))})})))}
}

func Rebox_Data_Maybe_1089607855_1201789390(in *Constructor_Data_Monoid_Monoid[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(in.V1)}
	return out
}

func Rebox_Data_Maybe_1170268447_3094389156(in *Constructor_Data_Maybe_Just[int64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Int(in.V0)
	return out
}

func Rebox_Data_Maybe_125234255_4179793454(in *Constructor_Data_Semigroup_Semigroup[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Maybe_1642601656_2568689657(in *Constructor_Control_Monad_Monad[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Monad_Monad[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Monad_Monad[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Maybe_1662389854_1766074591(in *Constructor_Data_Eq_Eq1[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Eq_Eq1[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq1[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Maybe_170825214_3985601471(in *Constructor_Data_Ord_Ord1[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Ord_Ord1[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord1[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Maybe_2047434479_2826095630(in *Constructor_Data_Semiring_Semiring[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Semiring_Semiring[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(in.V2)}
		out.V3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(in.V3)}
	return out
}

func Rebox_Data_Maybe_2155612431_4177771502(in *Constructor_Data_Ord_Ord[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Maybe_267397720_3706288089(in *Constructor_Control_Plus_Plus[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Plus_Plus[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Plus_Plus[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Maybe_3094389156_1170268447(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[int64]{}
		out.V0 = in.V0.IntVal
	return out
}

func Rebox_Data_Maybe_3123684004_1293498952(in *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Maybe_3363075976_138441832(in *Constructor_Data_Tuple_Tuple[int64, int64]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Int(in.V0)
		out.V1 = gopurs_runtime.Int(in.V1)
	return out
}

func Rebox_Data_Maybe_345859663_1386611502(in *Constructor_Data_Show_Show[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Maybe_3508461103_3790796878(in *Constructor_Data_Eq_Eq[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Maybe_3552963512_3741347833(in *Constructor_Control_Apply_Apply[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Apply_Apply[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Apply_Apply[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Maybe_3689823567_2812149806(in *Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Functor_Functor[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Functor_Functor[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Maybe_3818752775_2241343270(in *Constructor_Data_Functor_Invariant_Invariant[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Functor_Invariant_Invariant[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Functor_Invariant_Invariant[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Maybe_3887487416_2748095225(in *Constructor_Control_Bind_Bind[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Bind_Bind[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Bind_Bind[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Maybe_3904200120_3421983481(in *Constructor_Control_Alt_Alt[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Alt_Alt[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Alt_Alt[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Maybe_497217223_1323331594(in *Constructor_Data_Generic_Rep_Inl[uint32, gopurs_runtime.Value]) *Constructor_Data_Generic_Rep_Inl[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Generic_Rep_Inl[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V0), UnsafePtr: nil}
	return out
}

func Rebox_Data_Maybe_570732504_3290176857(in *Constructor_Control_Extend_Extend[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Extend_Extend[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Extend_Extend[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Maybe_591333647_2094947566(in *Constructor_Data_Bounded_Bounded[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(in.V1)}
		out.V2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(in.V2)}
	return out
}

func Rebox_Data_Maybe_649684152_1439734649(in *Constructor_Control_Applicative_Applicative[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Maybe_660678937_2687169876(in *Constructor_Data_Generic_Rep_Inr[uint32, gopurs_runtime.Value]) *Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Maybe_673668088_2307501113(in *Constructor_Control_Alternative_Alternative[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Maybe_732223057_2818661616(in *Constructor_Data_Generic_Rep_Generic[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}


