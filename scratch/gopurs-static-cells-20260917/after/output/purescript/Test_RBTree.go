package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Test_RBTree_R gopurs_runtime.Value
var once_Test_RBTree_R sync.Once
func Get_Test_RBTree_R() gopurs_runtime.Value {
	once_Test_RBTree_R.Do(func() {
		cache_Test_RBTree_R = gopurs_runtime.Value{Type: 9, IntVal: int64(3668501016), UnsafePtr: nil}
	})
	return cache_Test_RBTree_R
}

var cache_Test_RBTree_B gopurs_runtime.Value
var once_Test_RBTree_B sync.Once
func Get_Test_RBTree_B() gopurs_runtime.Value {
	once_Test_RBTree_B.Do(func() {
		cache_Test_RBTree_B = gopurs_runtime.Value{Type: 9, IntVal: int64(1583507464), UnsafePtr: nil}
	})
	return cache_Test_RBTree_B
}

var cache_Test_RBTree_E gopurs_runtime.Value
var once_Test_RBTree_E sync.Once
func Get_Test_RBTree_E() gopurs_runtime.Value {
	once_Test_RBTree_E.Do(func() {
		cache_Test_RBTree_E = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_Test_RBTree_T)(nil))}
	})
	return cache_Test_RBTree_E
}

var cache_Test_RBTree_T gopurs_runtime.Value
var once_Test_RBTree_T sync.Once
func Get_Test_RBTree_T() gopurs_runtime.Value {
	once_Test_RBTree_T.Do(func() {
		cache_Test_RBTree_T = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((&Constructor_Test_RBTree_T{1, uint32(value0.IntVal), gopurs_runtime.CoerceToStruct[Constructor_Test_RBTree_T](value1), value2.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Test_RBTree_T](value3)}))}
})
})
})
})
	})
	return cache_Test_RBTree_T
}

var cache_Test_RBTree_max gopurs_runtime.Value
var once_Test_RBTree_max sync.Once
func Get_Test_RBTree_max() gopurs_runtime.Value {
	once_Test_RBTree_max.Do(func() {
		cache_Test_RBTree_max = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Test_RBTree_max(x_0_box.IntVal, y_1_box.IntVal))
})
	})
	return cache_Test_RBTree_max
}

var cache_Test_RBTree_makeBlack gopurs_runtime.Value
var once_Test_RBTree_makeBlack sync.Once
func Get_Test_RBTree_makeBlack() gopurs_runtime.Value {
	once_Test_RBTree_makeBlack.Do(func() {
		cache_Test_RBTree_makeBlack = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(Call_Test_RBTree_makeBlack(gopurs_runtime.CoerceToStruct[Constructor_Test_RBTree_T](v_0_box)))}
})
	})
	return cache_Test_RBTree_makeBlack
}

var cache_Test_RBTree_describe gopurs_runtime.Value
var once_Test_RBTree_describe sync.Once
func Get_Test_RBTree_describe() gopurs_runtime.Value {
	once_Test_RBTree_describe.Do(func() {
		cache_Test_RBTree_describe = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Red-Black Tree (100k Worst-Case Insertions):"))
	})
	return cache_Test_RBTree_describe
}

var cache_Test_RBTree_depth gopurs_runtime.Value
var once_Test_RBTree_depth sync.Once
func Get_Test_RBTree_depth() gopurs_runtime.Value {
	once_Test_RBTree_depth.Do(func() {
		cache_Test_RBTree_depth = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Test_RBTree_depth(gopurs_runtime.CoerceToStruct[Constructor_Test_RBTree_T](v_0_box)))
})
	})
	return cache_Test_RBTree_depth
}

var cache_Test_RBTree_balance gopurs_runtime.Value
var once_Test_RBTree_balance sync.Once
func Get_Test_RBTree_balance() gopurs_runtime.Value {
	once_Test_RBTree_balance.Do(func() {
		cache_Test_RBTree_balance = gopurs_runtime.Func4(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value, v3_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(Call_Test_RBTree_balance(uint32(v_0_box.IntVal), gopurs_runtime.CoerceToStruct[Constructor_Test_RBTree_T](v1_1_box), v2_2_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Test_RBTree_T](v3_3_box)))}
})
	})
	return cache_Test_RBTree_balance
}

var cache_Test_RBTree_ins gopurs_runtime.Value
var once_Test_RBTree_ins sync.Once
func Get_Test_RBTree_ins() gopurs_runtime.Value {
	once_Test_RBTree_ins.Do(func() {
		cache_Test_RBTree_ins = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(Call_Test_RBTree_ins(v_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Test_RBTree_T](v1_1_box)))}
})
	})
	return cache_Test_RBTree_ins
}

var cache_Test_RBTree_insert gopurs_runtime.Value
var once_Test_RBTree_insert sync.Once
func Get_Test_RBTree_insert() gopurs_runtime.Value {
	once_Test_RBTree_insert.Do(func() {
		cache_Test_RBTree_insert = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(Call_Test_RBTree_insert(x_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Test_RBTree_T](s_1_box)))}
})
	})
	return cache_Test_RBTree_insert
}

var cache_Test_RBTree_buildTree gopurs_runtime.Value
var once_Test_RBTree_buildTree sync.Once
func Get_Test_RBTree_buildTree() gopurs_runtime.Value {
	once_Test_RBTree_buildTree.Do(func() {
		cache_Test_RBTree_buildTree = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(Call_Test_RBTree_buildTree(v_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Test_RBTree_T](v1_1_box)))}
})
	})
	return cache_Test_RBTree_buildTree
}

var cache_Test_RBTree_act gopurs_runtime.Value
var once_Test_RBTree_act sync.Once
func Get_Test_RBTree_act() gopurs_runtime.Value {
	once_Test_RBTree_act.Do(func() {
		cache_Test_RBTree_act = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
__local_var_0_0 := gopurs_runtime.Apply(Get_Bench_opaque(), gopurs_runtime.Int(int64(100000)))
_ = __local_var_0_0
__local_var_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
_ = __local_var_1_1
return gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(Call_Test_RBTree_depth(Call_Test_RBTree___gopurs_owned_buildTree_0(__local_var_1_1.IntVal, (*Constructor_Test_RBTree_T)(nil))))).StrVal())
})
	})
	return cache_Test_RBTree_act
}

func Call_Test_RBTree___gopurs_owned_balance_0_consume(__arg0 uint32, __arg1 *Constructor_Test_RBTree_T, __arg2 int64, __arg3 *Constructor_Test_RBTree_T, __donor *Constructor_Test_RBTree_T) *Constructor_Test_RBTree_T {
__owned_loop:
for {
if false { continue __owned_loop }
if (__arg0) == (1583507464) {
if (__arg1) != (nil) {
if (__arg1.V0) == (3668501016) {
if (__arg1.V1) != (nil) {
if (__arg1.V1.V0) == (3668501016) {
__let_scalar_40 := int64(__arg1.V1.V2)
_ = __let_scalar_40
__let_scalar_41 := int64(__arg1.V2)
_ = __let_scalar_41
__let_scalar_42 := int64(__arg2)
_ = __let_scalar_42
__scalar_43 := uint32(3668501016)
_ = __scalar_43
__scalar_44 := uint32(1583507464)
_ = __scalar_44
__read_45 := __arg1.V1.V1
_ = __read_45
__scalar_46 := int64(__let_scalar_40)
_ = __scalar_46
__read_47 := __arg1.V1.V3
_ = __read_47
__scalar_48 := int64(__let_scalar_41)
_ = __scalar_48
__scalar_49 := uint32(1583507464)
_ = __scalar_49
__read_50 := __arg1.V3
_ = __read_50
__scalar_51 := int64(__let_scalar_42)
_ = __scalar_51
__read_52 := __arg3
_ = __read_52
__donor_slot_55 := __donor
_ = __donor_slot_55
__dead_53 := __arg1
_ = __dead_53
__dead_54 := __arg1.V1
_ = __dead_54
__cell_56 := __dead_53
__cell_56.Rc = 1
__cell_56.V0 = __scalar_44
__cell_56.V1 = __read_45
__cell_56.V2 = __scalar_46
__cell_56.V3 = __read_47
__cell_57 := __dead_54
__cell_57.Rc = 1
__cell_57.V0 = __scalar_49
__cell_57.V1 = __read_50
__cell_57.V2 = __scalar_51
__cell_57.V3 = __read_52
__cell_58 := __donor_slot_55
if __cell_58 == nil {
__cell_58 = new(Constructor_Test_RBTree_T)
}
__cell_58.Rc = 1
__cell_58.V0 = __scalar_43
__cell_58.V1 = __cell_56
__cell_58.V2 = __scalar_48
__cell_58.V3 = __cell_57
return __cell_58
} else {
if (__arg1.V3) != (nil) {
if (__arg1.V3.V0) == (3668501016) {
__let_scalar_67 := int64(__arg1.V2)
_ = __let_scalar_67
__let_scalar_68 := int64(__arg1.V3.V2)
_ = __let_scalar_68
__let_scalar_69 := int64(__arg2)
_ = __let_scalar_69
__scalar_70 := uint32(3668501016)
_ = __scalar_70
__scalar_71 := uint32(1583507464)
_ = __scalar_71
__read_72 := __arg1.V1
_ = __read_72
__scalar_73 := int64(__let_scalar_67)
_ = __scalar_73
__read_74 := __arg1.V3.V1
_ = __read_74
__scalar_75 := int64(__let_scalar_68)
_ = __scalar_75
__scalar_76 := uint32(1583507464)
_ = __scalar_76
__read_77 := __arg1.V3.V3
_ = __read_77
__scalar_78 := int64(__let_scalar_69)
_ = __scalar_78
__read_79 := __arg3
_ = __read_79
__donor_slot_82 := __donor
_ = __donor_slot_82
__dead_80 := __arg1
_ = __dead_80
__dead_81 := __arg1.V3
_ = __dead_81
__cell_83 := __dead_80
__cell_83.Rc = 1
__cell_83.V0 = __scalar_71
__cell_83.V1 = __read_72
__cell_83.V2 = __scalar_73
__cell_83.V3 = __read_74
__cell_84 := __dead_81
__cell_84.Rc = 1
__cell_84.V0 = __scalar_76
__cell_84.V1 = __read_77
__cell_84.V2 = __scalar_78
__cell_84.V3 = __read_79
__cell_85 := __donor_slot_82
if __cell_85 == nil {
__cell_85 = new(Constructor_Test_RBTree_T)
}
__cell_85.Rc = 1
__cell_85.V0 = __scalar_70
__cell_85.V1 = __cell_83
__cell_85.V2 = __scalar_75
__cell_85.V3 = __cell_84
return __cell_85
} else {
if ((__arg3) != (nil)) && ((__arg3.V0) == (3668501016)) {
if (__arg3.V1) != (nil) {
if (__arg3.V1.V0) == (3668501016) {
__let_scalar_102 := int64(__arg2)
_ = __let_scalar_102
__let_scalar_103 := int64(__arg3.V1.V2)
_ = __let_scalar_103
__let_scalar_104 := int64(__arg3.V2)
_ = __let_scalar_104
__scalar_105 := uint32(3668501016)
_ = __scalar_105
__scalar_106 := uint32(1583507464)
_ = __scalar_106
__read_107 := __arg1
_ = __read_107
__scalar_108 := int64(__let_scalar_102)
_ = __scalar_108
__read_109 := __arg3.V1.V1
_ = __read_109
__scalar_110 := int64(__let_scalar_103)
_ = __scalar_110
__scalar_111 := uint32(1583507464)
_ = __scalar_111
__read_112 := __arg3.V1.V3
_ = __read_112
__scalar_113 := int64(__let_scalar_104)
_ = __scalar_113
__read_114 := __arg3.V3
_ = __read_114
__donor_slot_117 := __donor
_ = __donor_slot_117
__dead_115 := __arg3
_ = __dead_115
__dead_116 := __arg3.V1
_ = __dead_116
__cell_118 := __dead_115
__cell_118.Rc = 1
__cell_118.V0 = __scalar_106
__cell_118.V1 = __read_107
__cell_118.V2 = __scalar_108
__cell_118.V3 = __read_109
__cell_119 := __dead_116
__cell_119.Rc = 1
__cell_119.V0 = __scalar_111
__cell_119.V1 = __read_112
__cell_119.V2 = __scalar_113
__cell_119.V3 = __read_114
__cell_120 := __donor_slot_117
if __cell_120 == nil {
__cell_120 = new(Constructor_Test_RBTree_T)
}
__cell_120.Rc = 1
__cell_120.V0 = __scalar_105
__cell_120.V1 = __cell_118
__cell_120.V2 = __scalar_110
__cell_120.V3 = __cell_119
return __cell_120
} else {
if ((__arg3.V3) != (nil)) && ((__arg3.V3.V0) == (3668501016)) {
__let_scalar_121 := int64(__arg2)
_ = __let_scalar_121
__let_scalar_122 := int64(__arg3.V2)
_ = __let_scalar_122
__let_scalar_123 := int64(__arg3.V3.V2)
_ = __let_scalar_123
__scalar_124 := uint32(3668501016)
_ = __scalar_124
__scalar_125 := uint32(1583507464)
_ = __scalar_125
__read_126 := __arg1
_ = __read_126
__scalar_127 := int64(__let_scalar_121)
_ = __scalar_127
__read_128 := __arg3.V1
_ = __read_128
__scalar_129 := int64(__let_scalar_122)
_ = __scalar_129
__scalar_130 := uint32(1583507464)
_ = __scalar_130
__read_131 := __arg3.V3.V1
_ = __read_131
__scalar_132 := int64(__let_scalar_123)
_ = __scalar_132
__read_133 := __arg3.V3.V3
_ = __read_133
__donor_slot_136 := __donor
_ = __donor_slot_136
__dead_134 := __arg3
_ = __dead_134
__dead_135 := __arg3.V3
_ = __dead_135
__cell_137 := __dead_134
__cell_137.Rc = 1
__cell_137.V0 = __scalar_125
__cell_137.V1 = __read_126
__cell_137.V2 = __scalar_127
__cell_137.V3 = __read_128
__cell_138 := __dead_135
__cell_138.Rc = 1
__cell_138.V0 = __scalar_130
__cell_138.V1 = __read_131
__cell_138.V2 = __scalar_132
__cell_138.V3 = __read_133
__cell_139 := __donor_slot_136
if __cell_139 == nil {
__cell_139 = new(Constructor_Test_RBTree_T)
}
__cell_139.Rc = 1
__cell_139.V0 = __scalar_124
__cell_139.V1 = __cell_137
__cell_139.V2 = __scalar_129
__cell_139.V3 = __cell_138
return __cell_139
} else {
__let_scalar_94 := uint32(__arg0)
_ = __let_scalar_94
__let_scalar_95 := int64(__arg2)
_ = __let_scalar_95
__scalar_96 := uint32(__let_scalar_94)
_ = __scalar_96
__read_97 := __arg1
_ = __read_97
__scalar_98 := int64(__let_scalar_95)
_ = __scalar_98
__read_99 := __arg3
_ = __read_99
__donor_slot_100 := __donor
_ = __donor_slot_100
var __cell_101 *Constructor_Test_RBTree_T
if (__donor_slot_100) != (nil) {
__cell_101 = __donor_slot_100
__donor_slot_100 = nil
} else {

}
if (__cell_101) == (nil) {
__cell_101 = new(Constructor_Test_RBTree_T)
} else {

}
__cell_101.Rc = 1
__cell_101.V0 = __scalar_96
__cell_101.V1 = __read_97
__cell_101.V2 = __scalar_98
__cell_101.V3 = __read_99
return __cell_101
}
}
} else {
if ((__arg3.V3) != (nil)) && ((__arg3.V3.V0) == (3668501016)) {
__let_scalar_140 := int64(__arg2)
_ = __let_scalar_140
__let_scalar_141 := int64(__arg3.V2)
_ = __let_scalar_141
__let_scalar_142 := int64(__arg3.V3.V2)
_ = __let_scalar_142
__scalar_143 := uint32(3668501016)
_ = __scalar_143
__scalar_144 := uint32(1583507464)
_ = __scalar_144
__read_145 := __arg1
_ = __read_145
__scalar_146 := int64(__let_scalar_140)
_ = __scalar_146
__read_147 := __arg3.V1
_ = __read_147
__scalar_148 := int64(__let_scalar_141)
_ = __scalar_148
__scalar_149 := uint32(1583507464)
_ = __scalar_149
__read_150 := __arg3.V3.V1
_ = __read_150
__scalar_151 := int64(__let_scalar_142)
_ = __scalar_151
__read_152 := __arg3.V3.V3
_ = __read_152
__donor_slot_155 := __donor
_ = __donor_slot_155
__dead_153 := __arg3
_ = __dead_153
__dead_154 := __arg3.V3
_ = __dead_154
__cell_156 := __dead_153
__cell_156.Rc = 1
__cell_156.V0 = __scalar_144
__cell_156.V1 = __read_145
__cell_156.V2 = __scalar_146
__cell_156.V3 = __read_147
__cell_157 := __dead_154
__cell_157.Rc = 1
__cell_157.V0 = __scalar_149
__cell_157.V1 = __read_150
__cell_157.V2 = __scalar_151
__cell_157.V3 = __read_152
__cell_158 := __donor_slot_155
if __cell_158 == nil {
__cell_158 = new(Constructor_Test_RBTree_T)
}
__cell_158.Rc = 1
__cell_158.V0 = __scalar_143
__cell_158.V1 = __cell_156
__cell_158.V2 = __scalar_148
__cell_158.V3 = __cell_157
return __cell_158
} else {
__let_scalar_86 := uint32(__arg0)
_ = __let_scalar_86
__let_scalar_87 := int64(__arg2)
_ = __let_scalar_87
__scalar_88 := uint32(__let_scalar_86)
_ = __scalar_88
__read_89 := __arg1
_ = __read_89
__scalar_90 := int64(__let_scalar_87)
_ = __scalar_90
__read_91 := __arg3
_ = __read_91
__donor_slot_92 := __donor
_ = __donor_slot_92
var __cell_93 *Constructor_Test_RBTree_T
if (__donor_slot_92) != (nil) {
__cell_93 = __donor_slot_92
__donor_slot_92 = nil
} else {

}
if (__cell_93) == (nil) {
__cell_93 = new(Constructor_Test_RBTree_T)
} else {

}
__cell_93.Rc = 1
__cell_93.V0 = __scalar_88
__cell_93.V1 = __read_89
__cell_93.V2 = __scalar_90
__cell_93.V3 = __read_91
return __cell_93
}
}
} else {
__let_scalar_59 := uint32(__arg0)
_ = __let_scalar_59
__let_scalar_60 := int64(__arg2)
_ = __let_scalar_60
__scalar_61 := uint32(__let_scalar_59)
_ = __scalar_61
__read_62 := __arg1
_ = __read_62
__scalar_63 := int64(__let_scalar_60)
_ = __scalar_63
__read_64 := __arg3
_ = __read_64
__donor_slot_65 := __donor
_ = __donor_slot_65
var __cell_66 *Constructor_Test_RBTree_T
if (__donor_slot_65) != (nil) {
__cell_66 = __donor_slot_65
__donor_slot_65 = nil
} else {

}
if (__cell_66) == (nil) {
__cell_66 = new(Constructor_Test_RBTree_T)
} else {

}
__cell_66.Rc = 1
__cell_66.V0 = __scalar_61
__cell_66.V1 = __read_62
__cell_66.V2 = __scalar_63
__cell_66.V3 = __read_64
return __cell_66
}
}
} else {
if ((__arg3) != (nil)) && ((__arg3.V0) == (3668501016)) {
if (__arg3.V1) != (nil) {
if (__arg3.V1.V0) == (3668501016) {
__let_scalar_175 := int64(__arg2)
_ = __let_scalar_175
__let_scalar_176 := int64(__arg3.V1.V2)
_ = __let_scalar_176
__let_scalar_177 := int64(__arg3.V2)
_ = __let_scalar_177
__scalar_178 := uint32(3668501016)
_ = __scalar_178
__scalar_179 := uint32(1583507464)
_ = __scalar_179
__read_180 := __arg1
_ = __read_180
__scalar_181 := int64(__let_scalar_175)
_ = __scalar_181
__read_182 := __arg3.V1.V1
_ = __read_182
__scalar_183 := int64(__let_scalar_176)
_ = __scalar_183
__scalar_184 := uint32(1583507464)
_ = __scalar_184
__read_185 := __arg3.V1.V3
_ = __read_185
__scalar_186 := int64(__let_scalar_177)
_ = __scalar_186
__read_187 := __arg3.V3
_ = __read_187
__donor_slot_190 := __donor
_ = __donor_slot_190
__dead_188 := __arg3
_ = __dead_188
__dead_189 := __arg3.V1
_ = __dead_189
__cell_191 := __dead_188
__cell_191.Rc = 1
__cell_191.V0 = __scalar_179
__cell_191.V1 = __read_180
__cell_191.V2 = __scalar_181
__cell_191.V3 = __read_182
__cell_192 := __dead_189
__cell_192.Rc = 1
__cell_192.V0 = __scalar_184
__cell_192.V1 = __read_185
__cell_192.V2 = __scalar_186
__cell_192.V3 = __read_187
__cell_193 := __donor_slot_190
if __cell_193 == nil {
__cell_193 = new(Constructor_Test_RBTree_T)
}
__cell_193.Rc = 1
__cell_193.V0 = __scalar_178
__cell_193.V1 = __cell_191
__cell_193.V2 = __scalar_183
__cell_193.V3 = __cell_192
return __cell_193
} else {
if ((__arg3.V3) != (nil)) && ((__arg3.V3.V0) == (3668501016)) {
__let_scalar_194 := int64(__arg2)
_ = __let_scalar_194
__let_scalar_195 := int64(__arg3.V2)
_ = __let_scalar_195
__let_scalar_196 := int64(__arg3.V3.V2)
_ = __let_scalar_196
__scalar_197 := uint32(3668501016)
_ = __scalar_197
__scalar_198 := uint32(1583507464)
_ = __scalar_198
__read_199 := __arg1
_ = __read_199
__scalar_200 := int64(__let_scalar_194)
_ = __scalar_200
__read_201 := __arg3.V1
_ = __read_201
__scalar_202 := int64(__let_scalar_195)
_ = __scalar_202
__scalar_203 := uint32(1583507464)
_ = __scalar_203
__read_204 := __arg3.V3.V1
_ = __read_204
__scalar_205 := int64(__let_scalar_196)
_ = __scalar_205
__read_206 := __arg3.V3.V3
_ = __read_206
__donor_slot_209 := __donor
_ = __donor_slot_209
__dead_207 := __arg3
_ = __dead_207
__dead_208 := __arg3.V3
_ = __dead_208
__cell_210 := __dead_207
__cell_210.Rc = 1
__cell_210.V0 = __scalar_198
__cell_210.V1 = __read_199
__cell_210.V2 = __scalar_200
__cell_210.V3 = __read_201
__cell_211 := __dead_208
__cell_211.Rc = 1
__cell_211.V0 = __scalar_203
__cell_211.V1 = __read_204
__cell_211.V2 = __scalar_205
__cell_211.V3 = __read_206
__cell_212 := __donor_slot_209
if __cell_212 == nil {
__cell_212 = new(Constructor_Test_RBTree_T)
}
__cell_212.Rc = 1
__cell_212.V0 = __scalar_197
__cell_212.V1 = __cell_210
__cell_212.V2 = __scalar_202
__cell_212.V3 = __cell_211
return __cell_212
} else {
__let_scalar_167 := uint32(__arg0)
_ = __let_scalar_167
__let_scalar_168 := int64(__arg2)
_ = __let_scalar_168
__scalar_169 := uint32(__let_scalar_167)
_ = __scalar_169
__read_170 := __arg1
_ = __read_170
__scalar_171 := int64(__let_scalar_168)
_ = __scalar_171
__read_172 := __arg3
_ = __read_172
__donor_slot_173 := __donor
_ = __donor_slot_173
var __cell_174 *Constructor_Test_RBTree_T
if (__donor_slot_173) != (nil) {
__cell_174 = __donor_slot_173
__donor_slot_173 = nil
} else {

}
if (__cell_174) == (nil) {
__cell_174 = new(Constructor_Test_RBTree_T)
} else {

}
__cell_174.Rc = 1
__cell_174.V0 = __scalar_169
__cell_174.V1 = __read_170
__cell_174.V2 = __scalar_171
__cell_174.V3 = __read_172
return __cell_174
}
}
} else {
if ((__arg3.V3) != (nil)) && ((__arg3.V3.V0) == (3668501016)) {
__let_scalar_213 := int64(__arg2)
_ = __let_scalar_213
__let_scalar_214 := int64(__arg3.V2)
_ = __let_scalar_214
__let_scalar_215 := int64(__arg3.V3.V2)
_ = __let_scalar_215
__scalar_216 := uint32(3668501016)
_ = __scalar_216
__scalar_217 := uint32(1583507464)
_ = __scalar_217
__read_218 := __arg1
_ = __read_218
__scalar_219 := int64(__let_scalar_213)
_ = __scalar_219
__read_220 := __arg3.V1
_ = __read_220
__scalar_221 := int64(__let_scalar_214)
_ = __scalar_221
__scalar_222 := uint32(1583507464)
_ = __scalar_222
__read_223 := __arg3.V3.V1
_ = __read_223
__scalar_224 := int64(__let_scalar_215)
_ = __scalar_224
__read_225 := __arg3.V3.V3
_ = __read_225
__donor_slot_228 := __donor
_ = __donor_slot_228
__dead_226 := __arg3
_ = __dead_226
__dead_227 := __arg3.V3
_ = __dead_227
__cell_229 := __dead_226
__cell_229.Rc = 1
__cell_229.V0 = __scalar_217
__cell_229.V1 = __read_218
__cell_229.V2 = __scalar_219
__cell_229.V3 = __read_220
__cell_230 := __dead_227
__cell_230.Rc = 1
__cell_230.V0 = __scalar_222
__cell_230.V1 = __read_223
__cell_230.V2 = __scalar_224
__cell_230.V3 = __read_225
__cell_231 := __donor_slot_228
if __cell_231 == nil {
__cell_231 = new(Constructor_Test_RBTree_T)
}
__cell_231.Rc = 1
__cell_231.V0 = __scalar_216
__cell_231.V1 = __cell_229
__cell_231.V2 = __scalar_221
__cell_231.V3 = __cell_230
return __cell_231
} else {
__let_scalar_159 := uint32(__arg0)
_ = __let_scalar_159
__let_scalar_160 := int64(__arg2)
_ = __let_scalar_160
__scalar_161 := uint32(__let_scalar_159)
_ = __scalar_161
__read_162 := __arg1
_ = __read_162
__scalar_163 := int64(__let_scalar_160)
_ = __scalar_163
__read_164 := __arg3
_ = __read_164
__donor_slot_165 := __donor
_ = __donor_slot_165
var __cell_166 *Constructor_Test_RBTree_T
if (__donor_slot_165) != (nil) {
__cell_166 = __donor_slot_165
__donor_slot_165 = nil
} else {

}
if (__cell_166) == (nil) {
__cell_166 = new(Constructor_Test_RBTree_T)
} else {

}
__cell_166.Rc = 1
__cell_166.V0 = __scalar_161
__cell_166.V1 = __read_162
__cell_166.V2 = __scalar_163
__cell_166.V3 = __read_164
return __cell_166
}
}
} else {
__let_scalar_32 := uint32(__arg0)
_ = __let_scalar_32
__let_scalar_33 := int64(__arg2)
_ = __let_scalar_33
__scalar_34 := uint32(__let_scalar_32)
_ = __scalar_34
__read_35 := __arg1
_ = __read_35
__scalar_36 := int64(__let_scalar_33)
_ = __scalar_36
__read_37 := __arg3
_ = __read_37
__donor_slot_38 := __donor
_ = __donor_slot_38
var __cell_39 *Constructor_Test_RBTree_T
if (__donor_slot_38) != (nil) {
__cell_39 = __donor_slot_38
__donor_slot_38 = nil
} else {

}
if (__cell_39) == (nil) {
__cell_39 = new(Constructor_Test_RBTree_T)
} else {

}
__cell_39.Rc = 1
__cell_39.V0 = __scalar_34
__cell_39.V1 = __read_35
__cell_39.V2 = __scalar_36
__cell_39.V3 = __read_37
return __cell_39
}
}
}
} else {
if (__arg1.V3) != (nil) {
if (__arg1.V3.V0) == (3668501016) {
__let_scalar_240 := int64(__arg1.V2)
_ = __let_scalar_240
__let_scalar_241 := int64(__arg1.V3.V2)
_ = __let_scalar_241
__let_scalar_242 := int64(__arg2)
_ = __let_scalar_242
__scalar_243 := uint32(3668501016)
_ = __scalar_243
__scalar_244 := uint32(1583507464)
_ = __scalar_244
__read_245 := __arg1.V1
_ = __read_245
__scalar_246 := int64(__let_scalar_240)
_ = __scalar_246
__read_247 := __arg1.V3.V1
_ = __read_247
__scalar_248 := int64(__let_scalar_241)
_ = __scalar_248
__scalar_249 := uint32(1583507464)
_ = __scalar_249
__read_250 := __arg1.V3.V3
_ = __read_250
__scalar_251 := int64(__let_scalar_242)
_ = __scalar_251
__read_252 := __arg3
_ = __read_252
__donor_slot_255 := __donor
_ = __donor_slot_255
__dead_253 := __arg1
_ = __dead_253
__dead_254 := __arg1.V3
_ = __dead_254
__cell_256 := __dead_253
__cell_256.Rc = 1
__cell_256.V0 = __scalar_244
__cell_256.V1 = __read_245
__cell_256.V2 = __scalar_246
__cell_256.V3 = __read_247
__cell_257 := __dead_254
__cell_257.Rc = 1
__cell_257.V0 = __scalar_249
__cell_257.V1 = __read_250
__cell_257.V2 = __scalar_251
__cell_257.V3 = __read_252
__cell_258 := __donor_slot_255
if __cell_258 == nil {
__cell_258 = new(Constructor_Test_RBTree_T)
}
__cell_258.Rc = 1
__cell_258.V0 = __scalar_243
__cell_258.V1 = __cell_256
__cell_258.V2 = __scalar_248
__cell_258.V3 = __cell_257
return __cell_258
} else {
if ((__arg3) != (nil)) && ((__arg3.V0) == (3668501016)) {
if (__arg3.V1) != (nil) {
if (__arg3.V1.V0) == (3668501016) {
__let_scalar_275 := int64(__arg2)
_ = __let_scalar_275
__let_scalar_276 := int64(__arg3.V1.V2)
_ = __let_scalar_276
__let_scalar_277 := int64(__arg3.V2)
_ = __let_scalar_277
__scalar_278 := uint32(3668501016)
_ = __scalar_278
__scalar_279 := uint32(1583507464)
_ = __scalar_279
__read_280 := __arg1
_ = __read_280
__scalar_281 := int64(__let_scalar_275)
_ = __scalar_281
__read_282 := __arg3.V1.V1
_ = __read_282
__scalar_283 := int64(__let_scalar_276)
_ = __scalar_283
__scalar_284 := uint32(1583507464)
_ = __scalar_284
__read_285 := __arg3.V1.V3
_ = __read_285
__scalar_286 := int64(__let_scalar_277)
_ = __scalar_286
__read_287 := __arg3.V3
_ = __read_287
__donor_slot_290 := __donor
_ = __donor_slot_290
__dead_288 := __arg3
_ = __dead_288
__dead_289 := __arg3.V1
_ = __dead_289
__cell_291 := __dead_288
__cell_291.Rc = 1
__cell_291.V0 = __scalar_279
__cell_291.V1 = __read_280
__cell_291.V2 = __scalar_281
__cell_291.V3 = __read_282
__cell_292 := __dead_289
__cell_292.Rc = 1
__cell_292.V0 = __scalar_284
__cell_292.V1 = __read_285
__cell_292.V2 = __scalar_286
__cell_292.V3 = __read_287
__cell_293 := __donor_slot_290
if __cell_293 == nil {
__cell_293 = new(Constructor_Test_RBTree_T)
}
__cell_293.Rc = 1
__cell_293.V0 = __scalar_278
__cell_293.V1 = __cell_291
__cell_293.V2 = __scalar_283
__cell_293.V3 = __cell_292
return __cell_293
} else {
if ((__arg3.V3) != (nil)) && ((__arg3.V3.V0) == (3668501016)) {
__let_scalar_294 := int64(__arg2)
_ = __let_scalar_294
__let_scalar_295 := int64(__arg3.V2)
_ = __let_scalar_295
__let_scalar_296 := int64(__arg3.V3.V2)
_ = __let_scalar_296
__scalar_297 := uint32(3668501016)
_ = __scalar_297
__scalar_298 := uint32(1583507464)
_ = __scalar_298
__read_299 := __arg1
_ = __read_299
__scalar_300 := int64(__let_scalar_294)
_ = __scalar_300
__read_301 := __arg3.V1
_ = __read_301
__scalar_302 := int64(__let_scalar_295)
_ = __scalar_302
__scalar_303 := uint32(1583507464)
_ = __scalar_303
__read_304 := __arg3.V3.V1
_ = __read_304
__scalar_305 := int64(__let_scalar_296)
_ = __scalar_305
__read_306 := __arg3.V3.V3
_ = __read_306
__donor_slot_309 := __donor
_ = __donor_slot_309
__dead_307 := __arg3
_ = __dead_307
__dead_308 := __arg3.V3
_ = __dead_308
__cell_310 := __dead_307
__cell_310.Rc = 1
__cell_310.V0 = __scalar_298
__cell_310.V1 = __read_299
__cell_310.V2 = __scalar_300
__cell_310.V3 = __read_301
__cell_311 := __dead_308
__cell_311.Rc = 1
__cell_311.V0 = __scalar_303
__cell_311.V1 = __read_304
__cell_311.V2 = __scalar_305
__cell_311.V3 = __read_306
__cell_312 := __donor_slot_309
if __cell_312 == nil {
__cell_312 = new(Constructor_Test_RBTree_T)
}
__cell_312.Rc = 1
__cell_312.V0 = __scalar_297
__cell_312.V1 = __cell_310
__cell_312.V2 = __scalar_302
__cell_312.V3 = __cell_311
return __cell_312
} else {
__let_scalar_267 := uint32(__arg0)
_ = __let_scalar_267
__let_scalar_268 := int64(__arg2)
_ = __let_scalar_268
__scalar_269 := uint32(__let_scalar_267)
_ = __scalar_269
__read_270 := __arg1
_ = __read_270
__scalar_271 := int64(__let_scalar_268)
_ = __scalar_271
__read_272 := __arg3
_ = __read_272
__donor_slot_273 := __donor
_ = __donor_slot_273
var __cell_274 *Constructor_Test_RBTree_T
if (__donor_slot_273) != (nil) {
__cell_274 = __donor_slot_273
__donor_slot_273 = nil
} else {

}
if (__cell_274) == (nil) {
__cell_274 = new(Constructor_Test_RBTree_T)
} else {

}
__cell_274.Rc = 1
__cell_274.V0 = __scalar_269
__cell_274.V1 = __read_270
__cell_274.V2 = __scalar_271
__cell_274.V3 = __read_272
return __cell_274
}
}
} else {
if ((__arg3.V3) != (nil)) && ((__arg3.V3.V0) == (3668501016)) {
__let_scalar_313 := int64(__arg2)
_ = __let_scalar_313
__let_scalar_314 := int64(__arg3.V2)
_ = __let_scalar_314
__let_scalar_315 := int64(__arg3.V3.V2)
_ = __let_scalar_315
__scalar_316 := uint32(3668501016)
_ = __scalar_316
__scalar_317 := uint32(1583507464)
_ = __scalar_317
__read_318 := __arg1
_ = __read_318
__scalar_319 := int64(__let_scalar_313)
_ = __scalar_319
__read_320 := __arg3.V1
_ = __read_320
__scalar_321 := int64(__let_scalar_314)
_ = __scalar_321
__scalar_322 := uint32(1583507464)
_ = __scalar_322
__read_323 := __arg3.V3.V1
_ = __read_323
__scalar_324 := int64(__let_scalar_315)
_ = __scalar_324
__read_325 := __arg3.V3.V3
_ = __read_325
__donor_slot_328 := __donor
_ = __donor_slot_328
__dead_326 := __arg3
_ = __dead_326
__dead_327 := __arg3.V3
_ = __dead_327
__cell_329 := __dead_326
__cell_329.Rc = 1
__cell_329.V0 = __scalar_317
__cell_329.V1 = __read_318
__cell_329.V2 = __scalar_319
__cell_329.V3 = __read_320
__cell_330 := __dead_327
__cell_330.Rc = 1
__cell_330.V0 = __scalar_322
__cell_330.V1 = __read_323
__cell_330.V2 = __scalar_324
__cell_330.V3 = __read_325
__cell_331 := __donor_slot_328
if __cell_331 == nil {
__cell_331 = new(Constructor_Test_RBTree_T)
}
__cell_331.Rc = 1
__cell_331.V0 = __scalar_316
__cell_331.V1 = __cell_329
__cell_331.V2 = __scalar_321
__cell_331.V3 = __cell_330
return __cell_331
} else {
__let_scalar_259 := uint32(__arg0)
_ = __let_scalar_259
__let_scalar_260 := int64(__arg2)
_ = __let_scalar_260
__scalar_261 := uint32(__let_scalar_259)
_ = __scalar_261
__read_262 := __arg1
_ = __read_262
__scalar_263 := int64(__let_scalar_260)
_ = __scalar_263
__read_264 := __arg3
_ = __read_264
__donor_slot_265 := __donor
_ = __donor_slot_265
var __cell_266 *Constructor_Test_RBTree_T
if (__donor_slot_265) != (nil) {
__cell_266 = __donor_slot_265
__donor_slot_265 = nil
} else {

}
if (__cell_266) == (nil) {
__cell_266 = new(Constructor_Test_RBTree_T)
} else {

}
__cell_266.Rc = 1
__cell_266.V0 = __scalar_261
__cell_266.V1 = __read_262
__cell_266.V2 = __scalar_263
__cell_266.V3 = __read_264
return __cell_266
}
}
} else {
__let_scalar_232 := uint32(__arg0)
_ = __let_scalar_232
__let_scalar_233 := int64(__arg2)
_ = __let_scalar_233
__scalar_234 := uint32(__let_scalar_232)
_ = __scalar_234
__read_235 := __arg1
_ = __read_235
__scalar_236 := int64(__let_scalar_233)
_ = __scalar_236
__read_237 := __arg3
_ = __read_237
__donor_slot_238 := __donor
_ = __donor_slot_238
var __cell_239 *Constructor_Test_RBTree_T
if (__donor_slot_238) != (nil) {
__cell_239 = __donor_slot_238
__donor_slot_238 = nil
} else {

}
if (__cell_239) == (nil) {
__cell_239 = new(Constructor_Test_RBTree_T)
} else {

}
__cell_239.Rc = 1
__cell_239.V0 = __scalar_234
__cell_239.V1 = __read_235
__cell_239.V2 = __scalar_236
__cell_239.V3 = __read_237
return __cell_239
}
}
} else {
if ((__arg3) != (nil)) && ((__arg3.V0) == (3668501016)) {
if (__arg3.V1) != (nil) {
if (__arg3.V1.V0) == (3668501016) {
__let_scalar_348 := int64(__arg2)
_ = __let_scalar_348
__let_scalar_349 := int64(__arg3.V1.V2)
_ = __let_scalar_349
__let_scalar_350 := int64(__arg3.V2)
_ = __let_scalar_350
__scalar_351 := uint32(3668501016)
_ = __scalar_351
__scalar_352 := uint32(1583507464)
_ = __scalar_352
__read_353 := __arg1
_ = __read_353
__scalar_354 := int64(__let_scalar_348)
_ = __scalar_354
__read_355 := __arg3.V1.V1
_ = __read_355
__scalar_356 := int64(__let_scalar_349)
_ = __scalar_356
__scalar_357 := uint32(1583507464)
_ = __scalar_357
__read_358 := __arg3.V1.V3
_ = __read_358
__scalar_359 := int64(__let_scalar_350)
_ = __scalar_359
__read_360 := __arg3.V3
_ = __read_360
__donor_slot_363 := __donor
_ = __donor_slot_363
__dead_361 := __arg3
_ = __dead_361
__dead_362 := __arg3.V1
_ = __dead_362
__cell_364 := __dead_361
__cell_364.Rc = 1
__cell_364.V0 = __scalar_352
__cell_364.V1 = __read_353
__cell_364.V2 = __scalar_354
__cell_364.V3 = __read_355
__cell_365 := __dead_362
__cell_365.Rc = 1
__cell_365.V0 = __scalar_357
__cell_365.V1 = __read_358
__cell_365.V2 = __scalar_359
__cell_365.V3 = __read_360
__cell_366 := __donor_slot_363
if __cell_366 == nil {
__cell_366 = new(Constructor_Test_RBTree_T)
}
__cell_366.Rc = 1
__cell_366.V0 = __scalar_351
__cell_366.V1 = __cell_364
__cell_366.V2 = __scalar_356
__cell_366.V3 = __cell_365
return __cell_366
} else {
if ((__arg3.V3) != (nil)) && ((__arg3.V3.V0) == (3668501016)) {
__let_scalar_367 := int64(__arg2)
_ = __let_scalar_367
__let_scalar_368 := int64(__arg3.V2)
_ = __let_scalar_368
__let_scalar_369 := int64(__arg3.V3.V2)
_ = __let_scalar_369
__scalar_370 := uint32(3668501016)
_ = __scalar_370
__scalar_371 := uint32(1583507464)
_ = __scalar_371
__read_372 := __arg1
_ = __read_372
__scalar_373 := int64(__let_scalar_367)
_ = __scalar_373
__read_374 := __arg3.V1
_ = __read_374
__scalar_375 := int64(__let_scalar_368)
_ = __scalar_375
__scalar_376 := uint32(1583507464)
_ = __scalar_376
__read_377 := __arg3.V3.V1
_ = __read_377
__scalar_378 := int64(__let_scalar_369)
_ = __scalar_378
__read_379 := __arg3.V3.V3
_ = __read_379
__donor_slot_382 := __donor
_ = __donor_slot_382
__dead_380 := __arg3
_ = __dead_380
__dead_381 := __arg3.V3
_ = __dead_381
__cell_383 := __dead_380
__cell_383.Rc = 1
__cell_383.V0 = __scalar_371
__cell_383.V1 = __read_372
__cell_383.V2 = __scalar_373
__cell_383.V3 = __read_374
__cell_384 := __dead_381
__cell_384.Rc = 1
__cell_384.V0 = __scalar_376
__cell_384.V1 = __read_377
__cell_384.V2 = __scalar_378
__cell_384.V3 = __read_379
__cell_385 := __donor_slot_382
if __cell_385 == nil {
__cell_385 = new(Constructor_Test_RBTree_T)
}
__cell_385.Rc = 1
__cell_385.V0 = __scalar_370
__cell_385.V1 = __cell_383
__cell_385.V2 = __scalar_375
__cell_385.V3 = __cell_384
return __cell_385
} else {
__let_scalar_340 := uint32(__arg0)
_ = __let_scalar_340
__let_scalar_341 := int64(__arg2)
_ = __let_scalar_341
__scalar_342 := uint32(__let_scalar_340)
_ = __scalar_342
__read_343 := __arg1
_ = __read_343
__scalar_344 := int64(__let_scalar_341)
_ = __scalar_344
__read_345 := __arg3
_ = __read_345
__donor_slot_346 := __donor
_ = __donor_slot_346
var __cell_347 *Constructor_Test_RBTree_T
if (__donor_slot_346) != (nil) {
__cell_347 = __donor_slot_346
__donor_slot_346 = nil
} else {

}
if (__cell_347) == (nil) {
__cell_347 = new(Constructor_Test_RBTree_T)
} else {

}
__cell_347.Rc = 1
__cell_347.V0 = __scalar_342
__cell_347.V1 = __read_343
__cell_347.V2 = __scalar_344
__cell_347.V3 = __read_345
return __cell_347
}
}
} else {
if ((__arg3.V3) != (nil)) && ((__arg3.V3.V0) == (3668501016)) {
__let_scalar_386 := int64(__arg2)
_ = __let_scalar_386
__let_scalar_387 := int64(__arg3.V2)
_ = __let_scalar_387
__let_scalar_388 := int64(__arg3.V3.V2)
_ = __let_scalar_388
__scalar_389 := uint32(3668501016)
_ = __scalar_389
__scalar_390 := uint32(1583507464)
_ = __scalar_390
__read_391 := __arg1
_ = __read_391
__scalar_392 := int64(__let_scalar_386)
_ = __scalar_392
__read_393 := __arg3.V1
_ = __read_393
__scalar_394 := int64(__let_scalar_387)
_ = __scalar_394
__scalar_395 := uint32(1583507464)
_ = __scalar_395
__read_396 := __arg3.V3.V1
_ = __read_396
__scalar_397 := int64(__let_scalar_388)
_ = __scalar_397
__read_398 := __arg3.V3.V3
_ = __read_398
__donor_slot_401 := __donor
_ = __donor_slot_401
__dead_399 := __arg3
_ = __dead_399
__dead_400 := __arg3.V3
_ = __dead_400
__cell_402 := __dead_399
__cell_402.Rc = 1
__cell_402.V0 = __scalar_390
__cell_402.V1 = __read_391
__cell_402.V2 = __scalar_392
__cell_402.V3 = __read_393
__cell_403 := __dead_400
__cell_403.Rc = 1
__cell_403.V0 = __scalar_395
__cell_403.V1 = __read_396
__cell_403.V2 = __scalar_397
__cell_403.V3 = __read_398
__cell_404 := __donor_slot_401
if __cell_404 == nil {
__cell_404 = new(Constructor_Test_RBTree_T)
}
__cell_404.Rc = 1
__cell_404.V0 = __scalar_389
__cell_404.V1 = __cell_402
__cell_404.V2 = __scalar_394
__cell_404.V3 = __cell_403
return __cell_404
} else {
__let_scalar_332 := uint32(__arg0)
_ = __let_scalar_332
__let_scalar_333 := int64(__arg2)
_ = __let_scalar_333
__scalar_334 := uint32(__let_scalar_332)
_ = __scalar_334
__read_335 := __arg1
_ = __read_335
__scalar_336 := int64(__let_scalar_333)
_ = __scalar_336
__read_337 := __arg3
_ = __read_337
__donor_slot_338 := __donor
_ = __donor_slot_338
var __cell_339 *Constructor_Test_RBTree_T
if (__donor_slot_338) != (nil) {
__cell_339 = __donor_slot_338
__donor_slot_338 = nil
} else {

}
if (__cell_339) == (nil) {
__cell_339 = new(Constructor_Test_RBTree_T)
} else {

}
__cell_339.Rc = 1
__cell_339.V0 = __scalar_334
__cell_339.V1 = __read_335
__cell_339.V2 = __scalar_336
__cell_339.V3 = __read_337
return __cell_339
}
}
} else {
__let_scalar_24 := uint32(__arg0)
_ = __let_scalar_24
__let_scalar_25 := int64(__arg2)
_ = __let_scalar_25
__scalar_26 := uint32(__let_scalar_24)
_ = __scalar_26
__read_27 := __arg1
_ = __read_27
__scalar_28 := int64(__let_scalar_25)
_ = __scalar_28
__read_29 := __arg3
_ = __read_29
__donor_slot_30 := __donor
_ = __donor_slot_30
var __cell_31 *Constructor_Test_RBTree_T
if (__donor_slot_30) != (nil) {
__cell_31 = __donor_slot_30
__donor_slot_30 = nil
} else {

}
if (__cell_31) == (nil) {
__cell_31 = new(Constructor_Test_RBTree_T)
} else {

}
__cell_31.Rc = 1
__cell_31.V0 = __scalar_26
__cell_31.V1 = __read_27
__cell_31.V2 = __scalar_28
__cell_31.V3 = __read_29
return __cell_31
}
}
}
} else {
if ((__arg3) != (nil)) && ((__arg3.V0) == (3668501016)) {
if (__arg3.V1) != (nil) {
if (__arg3.V1.V0) == (3668501016) {
__let_scalar_421 := int64(__arg2)
_ = __let_scalar_421
__let_scalar_422 := int64(__arg3.V1.V2)
_ = __let_scalar_422
__let_scalar_423 := int64(__arg3.V2)
_ = __let_scalar_423
__scalar_424 := uint32(3668501016)
_ = __scalar_424
__scalar_425 := uint32(1583507464)
_ = __scalar_425
__read_426 := __arg1
_ = __read_426
__scalar_427 := int64(__let_scalar_421)
_ = __scalar_427
__read_428 := __arg3.V1.V1
_ = __read_428
__scalar_429 := int64(__let_scalar_422)
_ = __scalar_429
__scalar_430 := uint32(1583507464)
_ = __scalar_430
__read_431 := __arg3.V1.V3
_ = __read_431
__scalar_432 := int64(__let_scalar_423)
_ = __scalar_432
__read_433 := __arg3.V3
_ = __read_433
__donor_slot_436 := __donor
_ = __donor_slot_436
__dead_434 := __arg3
_ = __dead_434
__dead_435 := __arg3.V1
_ = __dead_435
__cell_437 := __dead_434
__cell_437.Rc = 1
__cell_437.V0 = __scalar_425
__cell_437.V1 = __read_426
__cell_437.V2 = __scalar_427
__cell_437.V3 = __read_428
__cell_438 := __dead_435
__cell_438.Rc = 1
__cell_438.V0 = __scalar_430
__cell_438.V1 = __read_431
__cell_438.V2 = __scalar_432
__cell_438.V3 = __read_433
__cell_439 := __donor_slot_436
if __cell_439 == nil {
__cell_439 = new(Constructor_Test_RBTree_T)
}
__cell_439.Rc = 1
__cell_439.V0 = __scalar_424
__cell_439.V1 = __cell_437
__cell_439.V2 = __scalar_429
__cell_439.V3 = __cell_438
return __cell_439
} else {
if ((__arg3.V3) != (nil)) && ((__arg3.V3.V0) == (3668501016)) {
__let_scalar_440 := int64(__arg2)
_ = __let_scalar_440
__let_scalar_441 := int64(__arg3.V2)
_ = __let_scalar_441
__let_scalar_442 := int64(__arg3.V3.V2)
_ = __let_scalar_442
__scalar_443 := uint32(3668501016)
_ = __scalar_443
__scalar_444 := uint32(1583507464)
_ = __scalar_444
__read_445 := __arg1
_ = __read_445
__scalar_446 := int64(__let_scalar_440)
_ = __scalar_446
__read_447 := __arg3.V1
_ = __read_447
__scalar_448 := int64(__let_scalar_441)
_ = __scalar_448
__scalar_449 := uint32(1583507464)
_ = __scalar_449
__read_450 := __arg3.V3.V1
_ = __read_450
__scalar_451 := int64(__let_scalar_442)
_ = __scalar_451
__read_452 := __arg3.V3.V3
_ = __read_452
__donor_slot_455 := __donor
_ = __donor_slot_455
__dead_453 := __arg3
_ = __dead_453
__dead_454 := __arg3.V3
_ = __dead_454
__cell_456 := __dead_453
__cell_456.Rc = 1
__cell_456.V0 = __scalar_444
__cell_456.V1 = __read_445
__cell_456.V2 = __scalar_446
__cell_456.V3 = __read_447
__cell_457 := __dead_454
__cell_457.Rc = 1
__cell_457.V0 = __scalar_449
__cell_457.V1 = __read_450
__cell_457.V2 = __scalar_451
__cell_457.V3 = __read_452
__cell_458 := __donor_slot_455
if __cell_458 == nil {
__cell_458 = new(Constructor_Test_RBTree_T)
}
__cell_458.Rc = 1
__cell_458.V0 = __scalar_443
__cell_458.V1 = __cell_456
__cell_458.V2 = __scalar_448
__cell_458.V3 = __cell_457
return __cell_458
} else {
__let_scalar_413 := uint32(__arg0)
_ = __let_scalar_413
__let_scalar_414 := int64(__arg2)
_ = __let_scalar_414
__scalar_415 := uint32(__let_scalar_413)
_ = __scalar_415
__read_416 := __arg1
_ = __read_416
__scalar_417 := int64(__let_scalar_414)
_ = __scalar_417
__read_418 := __arg3
_ = __read_418
__donor_slot_419 := __donor
_ = __donor_slot_419
var __cell_420 *Constructor_Test_RBTree_T
if (__donor_slot_419) != (nil) {
__cell_420 = __donor_slot_419
__donor_slot_419 = nil
} else {

}
if (__cell_420) == (nil) {
__cell_420 = new(Constructor_Test_RBTree_T)
} else {

}
__cell_420.Rc = 1
__cell_420.V0 = __scalar_415
__cell_420.V1 = __read_416
__cell_420.V2 = __scalar_417
__cell_420.V3 = __read_418
return __cell_420
}
}
} else {
if ((__arg3.V3) != (nil)) && ((__arg3.V3.V0) == (3668501016)) {
__let_scalar_459 := int64(__arg2)
_ = __let_scalar_459
__let_scalar_460 := int64(__arg3.V2)
_ = __let_scalar_460
__let_scalar_461 := int64(__arg3.V3.V2)
_ = __let_scalar_461
__scalar_462 := uint32(3668501016)
_ = __scalar_462
__scalar_463 := uint32(1583507464)
_ = __scalar_463
__read_464 := __arg1
_ = __read_464
__scalar_465 := int64(__let_scalar_459)
_ = __scalar_465
__read_466 := __arg3.V1
_ = __read_466
__scalar_467 := int64(__let_scalar_460)
_ = __scalar_467
__scalar_468 := uint32(1583507464)
_ = __scalar_468
__read_469 := __arg3.V3.V1
_ = __read_469
__scalar_470 := int64(__let_scalar_461)
_ = __scalar_470
__read_471 := __arg3.V3.V3
_ = __read_471
__donor_slot_474 := __donor
_ = __donor_slot_474
__dead_472 := __arg3
_ = __dead_472
__dead_473 := __arg3.V3
_ = __dead_473
__cell_475 := __dead_472
__cell_475.Rc = 1
__cell_475.V0 = __scalar_463
__cell_475.V1 = __read_464
__cell_475.V2 = __scalar_465
__cell_475.V3 = __read_466
__cell_476 := __dead_473
__cell_476.Rc = 1
__cell_476.V0 = __scalar_468
__cell_476.V1 = __read_469
__cell_476.V2 = __scalar_470
__cell_476.V3 = __read_471
__cell_477 := __donor_slot_474
if __cell_477 == nil {
__cell_477 = new(Constructor_Test_RBTree_T)
}
__cell_477.Rc = 1
__cell_477.V0 = __scalar_462
__cell_477.V1 = __cell_475
__cell_477.V2 = __scalar_467
__cell_477.V3 = __cell_476
return __cell_477
} else {
__let_scalar_405 := uint32(__arg0)
_ = __let_scalar_405
__let_scalar_406 := int64(__arg2)
_ = __let_scalar_406
__scalar_407 := uint32(__let_scalar_405)
_ = __scalar_407
__read_408 := __arg1
_ = __read_408
__scalar_409 := int64(__let_scalar_406)
_ = __scalar_409
__read_410 := __arg3
_ = __read_410
__donor_slot_411 := __donor
_ = __donor_slot_411
var __cell_412 *Constructor_Test_RBTree_T
if (__donor_slot_411) != (nil) {
__cell_412 = __donor_slot_411
__donor_slot_411 = nil
} else {

}
if (__cell_412) == (nil) {
__cell_412 = new(Constructor_Test_RBTree_T)
} else {

}
__cell_412.Rc = 1
__cell_412.V0 = __scalar_407
__cell_412.V1 = __read_408
__cell_412.V2 = __scalar_409
__cell_412.V3 = __read_410
return __cell_412
}
}
} else {
__let_scalar_16 := uint32(__arg0)
_ = __let_scalar_16
__let_scalar_17 := int64(__arg2)
_ = __let_scalar_17
__scalar_18 := uint32(__let_scalar_16)
_ = __scalar_18
__read_19 := __arg1
_ = __read_19
__scalar_20 := int64(__let_scalar_17)
_ = __scalar_20
__read_21 := __arg3
_ = __read_21
__donor_slot_22 := __donor
_ = __donor_slot_22
var __cell_23 *Constructor_Test_RBTree_T
if (__donor_slot_22) != (nil) {
__cell_23 = __donor_slot_22
__donor_slot_22 = nil
} else {

}
if (__cell_23) == (nil) {
__cell_23 = new(Constructor_Test_RBTree_T)
} else {

}
__cell_23.Rc = 1
__cell_23.V0 = __scalar_18
__cell_23.V1 = __read_19
__cell_23.V2 = __scalar_20
__cell_23.V3 = __read_21
return __cell_23
}
}
} else {
if ((__arg3) != (nil)) && ((__arg3.V0) == (3668501016)) {
if (__arg3.V1) != (nil) {
if (__arg3.V1.V0) == (3668501016) {
__let_scalar_494 := int64(__arg2)
_ = __let_scalar_494
__let_scalar_495 := int64(__arg3.V1.V2)
_ = __let_scalar_495
__let_scalar_496 := int64(__arg3.V2)
_ = __let_scalar_496
__scalar_497 := uint32(3668501016)
_ = __scalar_497
__scalar_498 := uint32(1583507464)
_ = __scalar_498
__read_499 := __arg1
_ = __read_499
__scalar_500 := int64(__let_scalar_494)
_ = __scalar_500
__read_501 := __arg3.V1.V1
_ = __read_501
__scalar_502 := int64(__let_scalar_495)
_ = __scalar_502
__scalar_503 := uint32(1583507464)
_ = __scalar_503
__read_504 := __arg3.V1.V3
_ = __read_504
__scalar_505 := int64(__let_scalar_496)
_ = __scalar_505
__read_506 := __arg3.V3
_ = __read_506
__donor_slot_509 := __donor
_ = __donor_slot_509
__dead_507 := __arg3
_ = __dead_507
__dead_508 := __arg3.V1
_ = __dead_508
__cell_510 := __dead_507
__cell_510.Rc = 1
__cell_510.V0 = __scalar_498
__cell_510.V1 = __read_499
__cell_510.V2 = __scalar_500
__cell_510.V3 = __read_501
__cell_511 := __dead_508
__cell_511.Rc = 1
__cell_511.V0 = __scalar_503
__cell_511.V1 = __read_504
__cell_511.V2 = __scalar_505
__cell_511.V3 = __read_506
__cell_512 := __donor_slot_509
if __cell_512 == nil {
__cell_512 = new(Constructor_Test_RBTree_T)
}
__cell_512.Rc = 1
__cell_512.V0 = __scalar_497
__cell_512.V1 = __cell_510
__cell_512.V2 = __scalar_502
__cell_512.V3 = __cell_511
return __cell_512
} else {
if ((__arg3.V3) != (nil)) && ((__arg3.V3.V0) == (3668501016)) {
__let_scalar_513 := int64(__arg2)
_ = __let_scalar_513
__let_scalar_514 := int64(__arg3.V2)
_ = __let_scalar_514
__let_scalar_515 := int64(__arg3.V3.V2)
_ = __let_scalar_515
__scalar_516 := uint32(3668501016)
_ = __scalar_516
__scalar_517 := uint32(1583507464)
_ = __scalar_517
__read_518 := __arg1
_ = __read_518
__scalar_519 := int64(__let_scalar_513)
_ = __scalar_519
__read_520 := __arg3.V1
_ = __read_520
__scalar_521 := int64(__let_scalar_514)
_ = __scalar_521
__scalar_522 := uint32(1583507464)
_ = __scalar_522
__read_523 := __arg3.V3.V1
_ = __read_523
__scalar_524 := int64(__let_scalar_515)
_ = __scalar_524
__read_525 := __arg3.V3.V3
_ = __read_525
__donor_slot_528 := __donor
_ = __donor_slot_528
__dead_526 := __arg3
_ = __dead_526
__dead_527 := __arg3.V3
_ = __dead_527
__cell_529 := __dead_526
__cell_529.Rc = 1
__cell_529.V0 = __scalar_517
__cell_529.V1 = __read_518
__cell_529.V2 = __scalar_519
__cell_529.V3 = __read_520
__cell_530 := __dead_527
__cell_530.Rc = 1
__cell_530.V0 = __scalar_522
__cell_530.V1 = __read_523
__cell_530.V2 = __scalar_524
__cell_530.V3 = __read_525
__cell_531 := __donor_slot_528
if __cell_531 == nil {
__cell_531 = new(Constructor_Test_RBTree_T)
}
__cell_531.Rc = 1
__cell_531.V0 = __scalar_516
__cell_531.V1 = __cell_529
__cell_531.V2 = __scalar_521
__cell_531.V3 = __cell_530
return __cell_531
} else {
__let_scalar_486 := uint32(__arg0)
_ = __let_scalar_486
__let_scalar_487 := int64(__arg2)
_ = __let_scalar_487
__scalar_488 := uint32(__let_scalar_486)
_ = __scalar_488
__read_489 := __arg1
_ = __read_489
__scalar_490 := int64(__let_scalar_487)
_ = __scalar_490
__read_491 := __arg3
_ = __read_491
__donor_slot_492 := __donor
_ = __donor_slot_492
var __cell_493 *Constructor_Test_RBTree_T
if (__donor_slot_492) != (nil) {
__cell_493 = __donor_slot_492
__donor_slot_492 = nil
} else {

}
if (__cell_493) == (nil) {
__cell_493 = new(Constructor_Test_RBTree_T)
} else {

}
__cell_493.Rc = 1
__cell_493.V0 = __scalar_488
__cell_493.V1 = __read_489
__cell_493.V2 = __scalar_490
__cell_493.V3 = __read_491
return __cell_493
}
}
} else {
if ((__arg3.V3) != (nil)) && ((__arg3.V3.V0) == (3668501016)) {
__let_scalar_532 := int64(__arg2)
_ = __let_scalar_532
__let_scalar_533 := int64(__arg3.V2)
_ = __let_scalar_533
__let_scalar_534 := int64(__arg3.V3.V2)
_ = __let_scalar_534
__scalar_535 := uint32(3668501016)
_ = __scalar_535
__scalar_536 := uint32(1583507464)
_ = __scalar_536
__read_537 := __arg1
_ = __read_537
__scalar_538 := int64(__let_scalar_532)
_ = __scalar_538
__read_539 := __arg3.V1
_ = __read_539
__scalar_540 := int64(__let_scalar_533)
_ = __scalar_540
__scalar_541 := uint32(1583507464)
_ = __scalar_541
__read_542 := __arg3.V3.V1
_ = __read_542
__scalar_543 := int64(__let_scalar_534)
_ = __scalar_543
__read_544 := __arg3.V3.V3
_ = __read_544
__donor_slot_547 := __donor
_ = __donor_slot_547
__dead_545 := __arg3
_ = __dead_545
__dead_546 := __arg3.V3
_ = __dead_546
__cell_548 := __dead_545
__cell_548.Rc = 1
__cell_548.V0 = __scalar_536
__cell_548.V1 = __read_537
__cell_548.V2 = __scalar_538
__cell_548.V3 = __read_539
__cell_549 := __dead_546
__cell_549.Rc = 1
__cell_549.V0 = __scalar_541
__cell_549.V1 = __read_542
__cell_549.V2 = __scalar_543
__cell_549.V3 = __read_544
__cell_550 := __donor_slot_547
if __cell_550 == nil {
__cell_550 = new(Constructor_Test_RBTree_T)
}
__cell_550.Rc = 1
__cell_550.V0 = __scalar_535
__cell_550.V1 = __cell_548
__cell_550.V2 = __scalar_540
__cell_550.V3 = __cell_549
return __cell_550
} else {
__let_scalar_478 := uint32(__arg0)
_ = __let_scalar_478
__let_scalar_479 := int64(__arg2)
_ = __let_scalar_479
__scalar_480 := uint32(__let_scalar_478)
_ = __scalar_480
__read_481 := __arg1
_ = __read_481
__scalar_482 := int64(__let_scalar_479)
_ = __scalar_482
__read_483 := __arg3
_ = __read_483
__donor_slot_484 := __donor
_ = __donor_slot_484
var __cell_485 *Constructor_Test_RBTree_T
if (__donor_slot_484) != (nil) {
__cell_485 = __donor_slot_484
__donor_slot_484 = nil
} else {

}
if (__cell_485) == (nil) {
__cell_485 = new(Constructor_Test_RBTree_T)
} else {

}
__cell_485.Rc = 1
__cell_485.V0 = __scalar_480
__cell_485.V1 = __read_481
__cell_485.V2 = __scalar_482
__cell_485.V3 = __read_483
return __cell_485
}
}
} else {
__let_scalar_8 := uint32(__arg0)
_ = __let_scalar_8
__let_scalar_9 := int64(__arg2)
_ = __let_scalar_9
__scalar_10 := uint32(__let_scalar_8)
_ = __scalar_10
__read_11 := __arg1
_ = __read_11
__scalar_12 := int64(__let_scalar_9)
_ = __scalar_12
__read_13 := __arg3
_ = __read_13
__donor_slot_14 := __donor
_ = __donor_slot_14
var __cell_15 *Constructor_Test_RBTree_T
if (__donor_slot_14) != (nil) {
__cell_15 = __donor_slot_14
__donor_slot_14 = nil
} else {

}
if (__cell_15) == (nil) {
__cell_15 = new(Constructor_Test_RBTree_T)
} else {

}
__cell_15.Rc = 1
__cell_15.V0 = __scalar_10
__cell_15.V1 = __read_11
__cell_15.V2 = __scalar_12
__cell_15.V3 = __read_13
return __cell_15
}
}
} else {
__let_scalar_0 := uint32(__arg0)
_ = __let_scalar_0
__let_scalar_1 := int64(__arg2)
_ = __let_scalar_1
__scalar_2 := uint32(__let_scalar_0)
_ = __scalar_2
__read_3 := __arg1
_ = __read_3
__scalar_4 := int64(__let_scalar_1)
_ = __scalar_4
__read_5 := __arg3
_ = __read_5
__donor_slot_6 := __donor
_ = __donor_slot_6
var __cell_7 *Constructor_Test_RBTree_T
if (__donor_slot_6) != (nil) {
__cell_7 = __donor_slot_6
__donor_slot_6 = nil
} else {

}
if (__cell_7) == (nil) {
__cell_7 = new(Constructor_Test_RBTree_T)
} else {

}
__cell_7.Rc = 1
__cell_7.V0 = __scalar_2
__cell_7.V1 = __read_3
__cell_7.V2 = __scalar_4
__cell_7.V3 = __read_5
return __cell_7
}
}
}

func Call_Test_RBTree___gopurs_owned_balance_0(__arg0 uint32, __arg1 *Constructor_Test_RBTree_T, __arg2 int64, __arg3 *Constructor_Test_RBTree_T) *Constructor_Test_RBTree_T {
return Call_Test_RBTree___gopurs_owned_balance_0_consume(__arg0, __arg1, __arg2, __arg3, nil)
}

func Call_Test_RBTree___gopurs_owned_buildTree_0_consume(__arg0 int64, __arg1 *Constructor_Test_RBTree_T, __donor *Constructor_Test_RBTree_T) *Constructor_Test_RBTree_T {
__owned_loop:
for {
if false { continue __owned_loop }
if (__arg0) == (int64(0)) {
__read_6 := __arg1
_ = __read_6
__donor_slot_7 := __donor
_ = __donor_slot_7
return __read_6
} else {
__scalar_0 := int64((__arg0) - (int64(1)))
_ = __scalar_0
__scalar_1 := int64(__arg0)
_ = __scalar_1
__read_2 := __arg1
_ = __read_2
__donor_slot_3 := __donor
_ = __donor_slot_3
__result_4 := Call_Test_RBTree___gopurs_owned_insert_0_consume(__scalar_1, __read_2, nil)
_ = __result_4
var __cell_5 *Constructor_Test_RBTree_T
if (__donor_slot_3) != (nil) {
__cell_5 = __donor_slot_3
__donor_slot_3 = nil
} else {

}
__arg0 = __scalar_0
__arg1 = __result_4
__donor = __cell_5
continue __owned_loop
}
}
}

func Call_Test_RBTree___gopurs_owned_buildTree_0(__arg0 int64, __arg1 *Constructor_Test_RBTree_T) *Constructor_Test_RBTree_T {
return Call_Test_RBTree___gopurs_owned_buildTree_0_consume(__arg0, __arg1, nil)
}

func Call_Test_RBTree___gopurs_owned_ins_0_consume(__arg0 int64, __arg1 *Constructor_Test_RBTree_T, __donor *Constructor_Test_RBTree_T) *Constructor_Test_RBTree_T {
__owned_loop:
for {
if false { continue __owned_loop }
if (__arg1) == (nil) {
__scalar_0 := uint32(3668501016)
_ = __scalar_0
__scalar_1 := int64(__arg0)
_ = __scalar_1
__donor_slot_3 := __donor
_ = __donor_slot_3
__dead_2 := __arg1
_ = __dead_2
var __cell_4 *Constructor_Test_RBTree_T
if (__donor_slot_3) != (nil) {
__cell_4 = __donor_slot_3
__donor_slot_3 = nil
} else {
if (__dead_2) != (nil) {
__cell_4 = __dead_2
__dead_2 = nil
} else {

}
}
if (__cell_4) == (nil) {
__cell_4 = new(Constructor_Test_RBTree_T)
} else {

}
__cell_4.Rc = 1
__cell_4.V0 = __scalar_0
__cell_4.V1 = nil
__cell_4.V2 = __scalar_1
__cell_4.V3 = nil
return __cell_4
} else {
if (__arg1) != (nil) {
if (__arg0) < (__arg1.V2) {
__scalar_22 := uint32(__arg1.V0)
_ = __scalar_22
__scalar_23 := int64(__arg0)
_ = __scalar_23
__read_24 := __arg1.V1
_ = __read_24
__scalar_25 := int64(__arg1.V2)
_ = __scalar_25
__read_26 := __arg1.V3
_ = __read_26
__donor_slot_28 := __donor
_ = __donor_slot_28
__dead_27 := __arg1
_ = __dead_27
__result_29 := Call_Test_RBTree___gopurs_owned_ins_0_consume(__scalar_23, __read_24, nil)
_ = __result_29
var __cell_30 *Constructor_Test_RBTree_T
if (__donor_slot_28) != (nil) {
__cell_30 = __donor_slot_28
__donor_slot_28 = nil
} else {
if (__dead_27) != (nil) {
__cell_30 = __dead_27
__dead_27 = nil
} else {

}
}
__result_31 := Call_Test_RBTree___gopurs_owned_balance_0_consume(__scalar_22, __result_29, __scalar_25, __read_26, __cell_30)
_ = __result_31
return __result_31
} else {
if (__arg0) > (__arg1.V2) {
__scalar_12 := uint32(__arg1.V0)
_ = __scalar_12
__read_13 := __arg1.V1
_ = __read_13
__scalar_14 := int64(__arg1.V2)
_ = __scalar_14
__scalar_15 := int64(__arg0)
_ = __scalar_15
__read_16 := __arg1.V3
_ = __read_16
__donor_slot_18 := __donor
_ = __donor_slot_18
__dead_17 := __arg1
_ = __dead_17
__result_19 := Call_Test_RBTree___gopurs_owned_ins_0_consume(__scalar_15, __read_16, nil)
_ = __result_19
var __cell_20 *Constructor_Test_RBTree_T
if (__donor_slot_18) != (nil) {
__cell_20 = __donor_slot_18
__donor_slot_18 = nil
} else {
if (__dead_17) != (nil) {
__cell_20 = __dead_17
__dead_17 = nil
} else {

}
}
__result_21 := Call_Test_RBTree___gopurs_owned_balance_0_consume(__scalar_12, __read_13, __scalar_14, __result_19, __cell_20)
_ = __result_21
return __result_21
} else {
__scalar_5 := uint32(__arg1.V0)
_ = __scalar_5
__read_6 := __arg1.V1
_ = __read_6
__scalar_7 := int64(__arg1.V2)
_ = __scalar_7
__read_8 := __arg1.V3
_ = __read_8
__donor_slot_10 := __donor
_ = __donor_slot_10
__dead_9 := __arg1
_ = __dead_9
var __cell_11 *Constructor_Test_RBTree_T
if (__donor_slot_10) != (nil) {
__cell_11 = __donor_slot_10
__donor_slot_10 = nil
} else {
if (__dead_9) != (nil) {
__cell_11 = __dead_9
__dead_9 = nil
} else {

}
}
if (__cell_11) == (nil) {
__cell_11 = new(Constructor_Test_RBTree_T)
} else {

}
__cell_11.Rc = 1
__cell_11.V0 = __scalar_5
__cell_11.V1 = __read_6
__cell_11.V2 = __scalar_7
__cell_11.V3 = __read_8
return __cell_11
}
}
} else {
panic("Failed pattern match")
}
}
}
}

func Call_Test_RBTree___gopurs_owned_ins_0(__arg0 int64, __arg1 *Constructor_Test_RBTree_T) *Constructor_Test_RBTree_T {
return Call_Test_RBTree___gopurs_owned_ins_0_consume(__arg0, __arg1, nil)
}

func Call_Test_RBTree___gopurs_owned_insert_0_consume(__arg0 int64, __arg1 *Constructor_Test_RBTree_T, __donor *Constructor_Test_RBTree_T) *Constructor_Test_RBTree_T {
__owned_loop:
for {
if false { continue __owned_loop }
__scalar_0 := int64(__arg0)
_ = __scalar_0
__read_1 := __arg1
_ = __read_1
__donor_slot_2 := __donor
_ = __donor_slot_2
var __cell_3 *Constructor_Test_RBTree_T
if (__donor_slot_2) != (nil) {
__cell_3 = __donor_slot_2
__donor_slot_2 = nil
} else {

}
__result_4 := Call_Test_RBTree___gopurs_owned_ins_0_consume(__scalar_0, __read_1, __cell_3)
_ = __result_4
var __let_tree_5 *Constructor_Test_RBTree_T
__let_tree_5 = __result_4
_ = __let_tree_5
__donor = nil
if (__let_tree_5) != (nil) {
__scalar_6 := uint32(1583507464)
_ = __scalar_6
__read_7 := __let_tree_5.V1
_ = __read_7
__scalar_8 := int64(__let_tree_5.V2)
_ = __scalar_8
__read_9 := __let_tree_5.V3
_ = __read_9
__donor_slot_11 := __donor
_ = __donor_slot_11
__dead_10 := __let_tree_5
_ = __dead_10
var __cell_12 *Constructor_Test_RBTree_T
if (__donor_slot_11) != (nil) {
__cell_12 = __donor_slot_11
__donor_slot_11 = nil
} else {
if (__dead_10) != (nil) {
__cell_12 = __dead_10
__dead_10 = nil
} else {

}
}
if (__cell_12) == (nil) {
__cell_12 = new(Constructor_Test_RBTree_T)
} else {

}
__cell_12.Rc = 1
__cell_12.V0 = __scalar_6
__cell_12.V1 = __read_7
__cell_12.V2 = __scalar_8
__cell_12.V3 = __read_9
return __cell_12
} else {
if (__let_tree_5) == (nil) {
__donor_slot_14 := __donor
_ = __donor_slot_14
__dead_13 := __let_tree_5
_ = __dead_13
return nil
} else {
panic("Failed pattern match")
}
}
}
}

func Call_Test_RBTree___gopurs_owned_insert_0(__arg0 int64, __arg1 *Constructor_Test_RBTree_T) *Constructor_Test_RBTree_T {
return Call_Test_RBTree___gopurs_owned_insert_0_consume(__arg0, __arg1, nil)
}

func Call_Test_RBTree___gopurs_owned_makeBlack_0_consume(__arg0 *Constructor_Test_RBTree_T, __donor *Constructor_Test_RBTree_T) *Constructor_Test_RBTree_T {
__owned_loop:
for {
if false { continue __owned_loop }
if (__arg0) != (nil) {
__scalar_0 := uint32(1583507464)
_ = __scalar_0
__read_1 := __arg0.V1
_ = __read_1
__scalar_2 := int64(__arg0.V2)
_ = __scalar_2
__read_3 := __arg0.V3
_ = __read_3
__donor_slot_5 := __donor
_ = __donor_slot_5
__dead_4 := __arg0
_ = __dead_4
var __cell_6 *Constructor_Test_RBTree_T
if (__donor_slot_5) != (nil) {
__cell_6 = __donor_slot_5
__donor_slot_5 = nil
} else {
if (__dead_4) != (nil) {
__cell_6 = __dead_4
__dead_4 = nil
} else {

}
}
if (__cell_6) == (nil) {
__cell_6 = new(Constructor_Test_RBTree_T)
} else {

}
__cell_6.Rc = 1
__cell_6.V0 = __scalar_0
__cell_6.V1 = __read_1
__cell_6.V2 = __scalar_2
__cell_6.V3 = __read_3
return __cell_6
} else {
if (__arg0) == (nil) {
__donor_slot_8 := __donor
_ = __donor_slot_8
__dead_7 := __arg0
_ = __dead_7
return nil
} else {
panic("Failed pattern match")
}
}
}
}

func Call_Test_RBTree___gopurs_owned_makeBlack_0(__arg0 *Constructor_Test_RBTree_T) *Constructor_Test_RBTree_T {
return Call_Test_RBTree___gopurs_owned_makeBlack_0_consume(__arg0, nil)
}

type Constructor_Test_RBTree_R struct {
	Rc uint32
}


type Constructor_Test_RBTree_B struct {
	Rc uint32
}


type Constructor_Test_RBTree_E struct {
	Rc uint32
}


type Constructor_Test_RBTree_T struct {
	Rc uint32
	V0 uint32
	V1 *Constructor_Test_RBTree_T
	V2 int64
	V3 *Constructor_Test_RBTree_T
}


func Call_Test_RBTree_max(x_0_loop int64, y_1_loop int64) int64 {
var x_0 int64 = x_0_loop
_ = x_0
var y_1 int64 = y_1_loop
_ = y_1
var __t0 int64
{
if (x_0) > (y_1) {
__t0 = x_0
goto end_branch_0
} else {

}
}
{
__t0 = y_1
}
end_branch_0:
return __t0
}

func Call_Test_RBTree_makeBlack(v_0_loop *Constructor_Test_RBTree_T) *Constructor_Test_RBTree_T {
var v_0 *Constructor_Test_RBTree_T = v_0_loop
_ = v_0
var __t1 *Constructor_Test_RBTree_T
{
if (v_0 != nil) {
var __reuse_0 *Constructor_Test_RBTree_T
if ((v_0) != (nil)) && (((v_0).V0) == (1583507464)) {
__reuse_0 = v_0
} else {
__reuse_0 = (&Constructor_Test_RBTree_T{1, 1583507464, (v_0).V1, (v_0).V2, (v_0).V3})
}
__t1 = __reuse_0
goto end_branch_1
} else {

}
}
{
if (v_0 == nil) {
__t1 = (*Constructor_Test_RBTree_T)(nil)
goto end_branch_1
} else {

}
}
{
__t1 = func() *Constructor_Test_RBTree_T { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}

func Call_Test_RBTree_depth(v_0_loop *Constructor_Test_RBTree_T) int64 {
depth:
for {
if false { continue depth }
var v_0 *Constructor_Test_RBTree_T = v_0_loop
_ = v_0
var __t3 int64
{
if (v_0 == nil) {
__t3 = int64(0)
goto end_branch_3
} else {

}
}
{
if (v_0 != nil) {
// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
__local_var_1_0 := Call_Test_RBTree_depth((v_0).V1)
_ = __local_var_1_0
// TAST (Let): __local_var_2_1 shape=App(Var) bindingType=Int
__local_var_2_1 := Call_Test_RBTree_depth((v_0).V3)
_ = __local_var_2_1
var __t2 int64
{
if (__local_var_1_0) > (__local_var_2_1) {
__t2 = __local_var_1_0
goto end_branch_2
} else {

}
}
{
__t2 = __local_var_2_1
}
end_branch_2:
__t3 = (int64(1)) + (__t2)
goto end_branch_3
} else {

}
}
{
__t3 = func() int64 { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}
}

func Call_Test_RBTree_balance(v_0_loop uint32, v1_1_loop *Constructor_Test_RBTree_T, v2_2_loop int64, v3_3_loop *Constructor_Test_RBTree_T) *Constructor_Test_RBTree_T {
var v_0 uint32 = v_0_loop
_ = v_0
var v1_1 *Constructor_Test_RBTree_T = v1_1_loop
_ = v1_1
var v2_2 int64 = v2_2_loop
_ = v2_2
var v3_3 *Constructor_Test_RBTree_T = v3_3_loop
_ = v3_3
var __t308 *Constructor_Test_RBTree_T
{
if (v_0 == 1583507464) {
var __t307 *Constructor_Test_RBTree_T
{
if (v1_1 != nil) {
var __t265 *Constructor_Test_RBTree_T
{
var __t_tag_12 uint32 = (v1_1).V0
_ = __t_tag_12
if (uint32(__t_tag_12) == 3668501016) {
var __t223 *Constructor_Test_RBTree_T
{
var __t_tag_17 *Constructor_Test_RBTree_T = (v1_1).V1
_ = __t_tag_17
if (__t_tag_17 != nil) {
var __t126 *Constructor_Test_RBTree_T
{
var __t_tag_22 uint32 = ((v1_1).V1).V0
_ = __t_tag_22
if (uint32(__t_tag_22) == 3668501016) {
// TAST (Let): __local_var_4_23 shape=Other bindingType=Any
__local_var_4_23 := ((v1_1).V1).V1
_ = __local_var_4_23
// TAST (Let): __local_var_5_24 shape=Other bindingType=Any
__local_var_5_24 := ((v1_1).V1).V3
_ = __local_var_5_24
// TAST (Let): __local_var_6_25 shape=Other bindingType=Any
__local_var_6_25 := (v1_1).V3
_ = __local_var_6_25
// TAST (Let): __local_var_7_26 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_7_26 := v3_3
_ = __local_var_7_26
// TAST (Let): __local_var_8_27 shape=Other bindingType=Any
__local_var_8_27 := ((v1_1).V1).V2
_ = __local_var_8_27
// TAST (Let): __local_var_9_28 shape=Other bindingType=Any
__local_var_9_28 := (v1_1).V2
_ = __local_var_9_28
// TAST (Let): __local_var_10_29 shape=Other bindingType=Int
__local_var_10_29 := v2_2
_ = __local_var_10_29
__t126 = (&Constructor_Test_RBTree_T{1, 3668501016, (&Constructor_Test_RBTree_T{1, 1583507464, __local_var_4_23, __local_var_8_27, __local_var_5_24}), __local_var_9_28, (&Constructor_Test_RBTree_T{1, 1583507464, __local_var_6_25, __local_var_10_29, __local_var_7_26})})
goto end_branch_126
} else {

}
}
{
var __t_tag_30 *Constructor_Test_RBTree_T = (v1_1).V3
_ = __t_tag_30
if (__t_tag_30 != nil) {
var __t84 *Constructor_Test_RBTree_T
{
var __t_tag_35 uint32 = ((v1_1).V3).V0
_ = __t_tag_35
if (uint32(__t_tag_35) == 3668501016) {
// TAST (Let): __local_var_4_36 shape=Other bindingType=Any
__local_var_4_36 := (v1_1).V1
_ = __local_var_4_36
// TAST (Let): __local_var_5_37 shape=Other bindingType=Any
__local_var_5_37 := ((v1_1).V3).V1
_ = __local_var_5_37
// TAST (Let): __local_var_6_38 shape=Other bindingType=Any
__local_var_6_38 := ((v1_1).V3).V3
_ = __local_var_6_38
// TAST (Let): __local_var_7_39 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_7_39 := v3_3
_ = __local_var_7_39
// TAST (Let): __local_var_8_40 shape=Other bindingType=Any
__local_var_8_40 := (v1_1).V2
_ = __local_var_8_40
// TAST (Let): __local_var_9_41 shape=Other bindingType=Any
__local_var_9_41 := ((v1_1).V3).V2
_ = __local_var_9_41
// TAST (Let): __local_var_10_42 shape=Other bindingType=Int
__local_var_10_42 := v2_2
_ = __local_var_10_42
__t84 = (&Constructor_Test_RBTree_T{1, 3668501016, (&Constructor_Test_RBTree_T{1, 1583507464, __local_var_4_36, __local_var_8_40, __local_var_5_37}), __local_var_9_41, (&Constructor_Test_RBTree_T{1, 1583507464, __local_var_6_38, __local_var_10_42, __local_var_7_39})})
goto end_branch_84
} else {

}
}
{
var __t_and_44 bool = false
if (v3_3 != nil) {

var __t_tag_43 uint32 = (v3_3).V0
_ = __t_tag_43
__t_and_44 = (uint32(__t_tag_43) == 3668501016)
}
if __t_and_44 {
var __t83 *Constructor_Test_RBTree_T
{
var __t_tag_49 *Constructor_Test_RBTree_T = (v3_3).V1
_ = __t_tag_49
if (__t_tag_49 != nil) {
var __t72 *Constructor_Test_RBTree_T
{
var __t_tag_54 uint32 = ((v3_3).V1).V0
_ = __t_tag_54
if (uint32(__t_tag_54) == 3668501016) {
// TAST (Let): __local_var_4_55 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_4_55 := v1_1
_ = __local_var_4_55
// TAST (Let): __local_var_5_56 shape=Other bindingType=Any
__local_var_5_56 := ((v3_3).V1).V1
_ = __local_var_5_56
// TAST (Let): __local_var_6_57 shape=Other bindingType=Any
__local_var_6_57 := ((v3_3).V1).V3
_ = __local_var_6_57
// TAST (Let): __local_var_7_58 shape=Other bindingType=Any
__local_var_7_58 := (v3_3).V3
_ = __local_var_7_58
// TAST (Let): __local_var_8_59 shape=Other bindingType=Int
__local_var_8_59 := v2_2
_ = __local_var_8_59
// TAST (Let): __local_var_9_60 shape=Other bindingType=Any
__local_var_9_60 := ((v3_3).V1).V2
_ = __local_var_9_60
// TAST (Let): __local_var_10_61 shape=Other bindingType=Any
__local_var_10_61 := (v3_3).V2
_ = __local_var_10_61
__t72 = (&Constructor_Test_RBTree_T{1, 3668501016, (&Constructor_Test_RBTree_T{1, 1583507464, __local_var_4_55, __local_var_8_59, __local_var_5_56}), __local_var_9_60, (&Constructor_Test_RBTree_T{1, 1583507464, __local_var_6_57, __local_var_10_61, __local_var_7_58})})
goto end_branch_72
} else {

}
}
{
var __t_tag_62 *Constructor_Test_RBTree_T = (v3_3).V3
_ = __t_tag_62
var __t_and_64 bool = false
if (__t_tag_62 != nil) {

var __t_tag_63 uint32 = ((v3_3).V3).V0
_ = __t_tag_63
__t_and_64 = (uint32(__t_tag_63) == 3668501016)
}
if __t_and_64 {
// TAST (Let): __local_var_4_65 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_4_65 := v1_1
_ = __local_var_4_65
// TAST (Let): __local_var_5_66 shape=Other bindingType=Any
__local_var_5_66 := (v3_3).V1
_ = __local_var_5_66
// TAST (Let): __local_var_6_67 shape=Other bindingType=Any
__local_var_6_67 := ((v3_3).V3).V1
_ = __local_var_6_67
// TAST (Let): __local_var_7_68 shape=Other bindingType=Any
__local_var_7_68 := ((v3_3).V3).V3
_ = __local_var_7_68
// TAST (Let): __local_var_8_69 shape=Other bindingType=Int
__local_var_8_69 := v2_2
_ = __local_var_8_69
// TAST (Let): __local_var_9_70 shape=Other bindingType=Any
__local_var_9_70 := (v3_3).V2
_ = __local_var_9_70
// TAST (Let): __local_var_10_71 shape=Other bindingType=Any
__local_var_10_71 := ((v3_3).V3).V2
_ = __local_var_10_71
__t72 = (&Constructor_Test_RBTree_T{1, 3668501016, (&Constructor_Test_RBTree_T{1, 1583507464, __local_var_4_65, __local_var_8_69, __local_var_5_66}), __local_var_9_70, (&Constructor_Test_RBTree_T{1, 1583507464, __local_var_6_67, __local_var_10_71, __local_var_7_68})})
goto end_branch_72
} else {

}
}
{
// TAST (Let): __local_var_4_50 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_4_50 := v1_1
_ = __local_var_4_50
// TAST (Let): __local_var_5_51 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_5_51 := v3_3
_ = __local_var_5_51
// TAST (Let): __local_var_6_52 shape=Other bindingType=(ADT ["Test","RBTree","Color"] [])
__local_var_6_52 := v_0
_ = __local_var_6_52
// TAST (Let): __local_var_7_53 shape=Other bindingType=Int
__local_var_7_53 := v2_2
_ = __local_var_7_53
__t72 = (&Constructor_Test_RBTree_T{1, __local_var_6_52, __local_var_4_50, __local_var_7_53, __local_var_5_51})
}
end_branch_72:
__t83 = __t72
goto end_branch_83
} else {

}
}
{
var __t_tag_73 *Constructor_Test_RBTree_T = (v3_3).V3
_ = __t_tag_73
var __t_and_75 bool = false
if (__t_tag_73 != nil) {

var __t_tag_74 uint32 = ((v3_3).V3).V0
_ = __t_tag_74
__t_and_75 = (uint32(__t_tag_74) == 3668501016)
}
if __t_and_75 {
// TAST (Let): __local_var_4_76 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_4_76 := v1_1
_ = __local_var_4_76
// TAST (Let): __local_var_5_77 shape=Other bindingType=Any
__local_var_5_77 := (v3_3).V1
_ = __local_var_5_77
// TAST (Let): __local_var_6_78 shape=Other bindingType=Any
__local_var_6_78 := ((v3_3).V3).V1
_ = __local_var_6_78
// TAST (Let): __local_var_7_79 shape=Other bindingType=Any
__local_var_7_79 := ((v3_3).V3).V3
_ = __local_var_7_79
// TAST (Let): __local_var_8_80 shape=Other bindingType=Int
__local_var_8_80 := v2_2
_ = __local_var_8_80
// TAST (Let): __local_var_9_81 shape=Other bindingType=Any
__local_var_9_81 := (v3_3).V2
_ = __local_var_9_81
// TAST (Let): __local_var_10_82 shape=Other bindingType=Any
__local_var_10_82 := ((v3_3).V3).V2
_ = __local_var_10_82
__t83 = (&Constructor_Test_RBTree_T{1, 3668501016, (&Constructor_Test_RBTree_T{1, 1583507464, __local_var_4_76, __local_var_8_80, __local_var_5_77}), __local_var_9_81, (&Constructor_Test_RBTree_T{1, 1583507464, __local_var_6_78, __local_var_10_82, __local_var_7_79})})
goto end_branch_83
} else {

}
}
{
// TAST (Let): __local_var_4_45 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_4_45 := v1_1
_ = __local_var_4_45
// TAST (Let): __local_var_5_46 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_5_46 := v3_3
_ = __local_var_5_46
// TAST (Let): __local_var_6_47 shape=Other bindingType=(ADT ["Test","RBTree","Color"] [])
__local_var_6_47 := v_0
_ = __local_var_6_47
// TAST (Let): __local_var_7_48 shape=Other bindingType=Int
__local_var_7_48 := v2_2
_ = __local_var_7_48
__t83 = (&Constructor_Test_RBTree_T{1, __local_var_6_47, __local_var_4_45, __local_var_7_48, __local_var_5_46})
}
end_branch_83:
__t84 = __t83
goto end_branch_84
} else {

}
}
{
// TAST (Let): __local_var_4_31 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_4_31 := v1_1
_ = __local_var_4_31
// TAST (Let): __local_var_5_32 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_5_32 := v3_3
_ = __local_var_5_32
// TAST (Let): __local_var_6_33 shape=Other bindingType=(ADT ["Test","RBTree","Color"] [])
__local_var_6_33 := v_0
_ = __local_var_6_33
// TAST (Let): __local_var_7_34 shape=Other bindingType=Int
__local_var_7_34 := v2_2
_ = __local_var_7_34
__t84 = (&Constructor_Test_RBTree_T{1, __local_var_6_33, __local_var_4_31, __local_var_7_34, __local_var_5_32})
}
end_branch_84:
__t126 = __t84
goto end_branch_126
} else {

}
}
{
var __t_and_86 bool = false
if (v3_3 != nil) {

var __t_tag_85 uint32 = (v3_3).V0
_ = __t_tag_85
__t_and_86 = (uint32(__t_tag_85) == 3668501016)
}
if __t_and_86 {
var __t125 *Constructor_Test_RBTree_T
{
var __t_tag_91 *Constructor_Test_RBTree_T = (v3_3).V1
_ = __t_tag_91
if (__t_tag_91 != nil) {
var __t114 *Constructor_Test_RBTree_T
{
var __t_tag_96 uint32 = ((v3_3).V1).V0
_ = __t_tag_96
if (uint32(__t_tag_96) == 3668501016) {
// TAST (Let): __local_var_4_97 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_4_97 := v1_1
_ = __local_var_4_97
// TAST (Let): __local_var_5_98 shape=Other bindingType=Any
__local_var_5_98 := ((v3_3).V1).V1
_ = __local_var_5_98
// TAST (Let): __local_var_6_99 shape=Other bindingType=Any
__local_var_6_99 := ((v3_3).V1).V3
_ = __local_var_6_99
// TAST (Let): __local_var_7_100 shape=Other bindingType=Any
__local_var_7_100 := (v3_3).V3
_ = __local_var_7_100
// TAST (Let): __local_var_8_101 shape=Other bindingType=Int
__local_var_8_101 := v2_2
_ = __local_var_8_101
// TAST (Let): __local_var_9_102 shape=Other bindingType=Any
__local_var_9_102 := ((v3_3).V1).V2
_ = __local_var_9_102
// TAST (Let): __local_var_10_103 shape=Other bindingType=Any
__local_var_10_103 := (v3_3).V2
_ = __local_var_10_103
__t114 = (&Constructor_Test_RBTree_T{1, 3668501016, (&Constructor_Test_RBTree_T{1, 1583507464, __local_var_4_97, __local_var_8_101, __local_var_5_98}), __local_var_9_102, (&Constructor_Test_RBTree_T{1, 1583507464, __local_var_6_99, __local_var_10_103, __local_var_7_100})})
goto end_branch_114
} else {

}
}
{
var __t_tag_104 *Constructor_Test_RBTree_T = (v3_3).V3
_ = __t_tag_104
var __t_and_106 bool = false
if (__t_tag_104 != nil) {

var __t_tag_105 uint32 = ((v3_3).V3).V0
_ = __t_tag_105
__t_and_106 = (uint32(__t_tag_105) == 3668501016)
}
if __t_and_106 {
// TAST (Let): __local_var_4_107 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_4_107 := v1_1
_ = __local_var_4_107
// TAST (Let): __local_var_5_108 shape=Other bindingType=Any
__local_var_5_108 := (v3_3).V1
_ = __local_var_5_108
// TAST (Let): __local_var_6_109 shape=Other bindingType=Any
__local_var_6_109 := ((v3_3).V3).V1
_ = __local_var_6_109
// TAST (Let): __local_var_7_110 shape=Other bindingType=Any
__local_var_7_110 := ((v3_3).V3).V3
_ = __local_var_7_110
// TAST (Let): __local_var_8_111 shape=Other bindingType=Int
__local_var_8_111 := v2_2
_ = __local_var_8_111
// TAST (Let): __local_var_9_112 shape=Other bindingType=Any
__local_var_9_112 := (v3_3).V2
_ = __local_var_9_112
// TAST (Let): __local_var_10_113 shape=Other bindingType=Any
__local_var_10_113 := ((v3_3).V3).V2
_ = __local_var_10_113
__t114 = (&Constructor_Test_RBTree_T{1, 3668501016, (&Constructor_Test_RBTree_T{1, 1583507464, __local_var_4_107, __local_var_8_111, __local_var_5_108}), __local_var_9_112, (&Constructor_Test_RBTree_T{1, 1583507464, __local_var_6_109, __local_var_10_113, __local_var_7_110})})
goto end_branch_114
} else {

}
}
{
// TAST (Let): __local_var_4_92 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_4_92 := v1_1
_ = __local_var_4_92
// TAST (Let): __local_var_5_93 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_5_93 := v3_3
_ = __local_var_5_93
// TAST (Let): __local_var_6_94 shape=Other bindingType=(ADT ["Test","RBTree","Color"] [])
__local_var_6_94 := v_0
_ = __local_var_6_94
// TAST (Let): __local_var_7_95 shape=Other bindingType=Int
__local_var_7_95 := v2_2
_ = __local_var_7_95
__t114 = (&Constructor_Test_RBTree_T{1, __local_var_6_94, __local_var_4_92, __local_var_7_95, __local_var_5_93})
}
end_branch_114:
__t125 = __t114
goto end_branch_125
} else {

}
}
{
var __t_tag_115 *Constructor_Test_RBTree_T = (v3_3).V3
_ = __t_tag_115
var __t_and_117 bool = false
if (__t_tag_115 != nil) {

var __t_tag_116 uint32 = ((v3_3).V3).V0
_ = __t_tag_116
__t_and_117 = (uint32(__t_tag_116) == 3668501016)
}
if __t_and_117 {
// TAST (Let): __local_var_4_118 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_4_118 := v1_1
_ = __local_var_4_118
// TAST (Let): __local_var_5_119 shape=Other bindingType=Any
__local_var_5_119 := (v3_3).V1
_ = __local_var_5_119
// TAST (Let): __local_var_6_120 shape=Other bindingType=Any
__local_var_6_120 := ((v3_3).V3).V1
_ = __local_var_6_120
// TAST (Let): __local_var_7_121 shape=Other bindingType=Any
__local_var_7_121 := ((v3_3).V3).V3
_ = __local_var_7_121
// TAST (Let): __local_var_8_122 shape=Other bindingType=Int
__local_var_8_122 := v2_2
_ = __local_var_8_122
// TAST (Let): __local_var_9_123 shape=Other bindingType=Any
__local_var_9_123 := (v3_3).V2
_ = __local_var_9_123
// TAST (Let): __local_var_10_124 shape=Other bindingType=Any
__local_var_10_124 := ((v3_3).V3).V2
_ = __local_var_10_124
__t125 = (&Constructor_Test_RBTree_T{1, 3668501016, (&Constructor_Test_RBTree_T{1, 1583507464, __local_var_4_118, __local_var_8_122, __local_var_5_119}), __local_var_9_123, (&Constructor_Test_RBTree_T{1, 1583507464, __local_var_6_120, __local_var_10_124, __local_var_7_121})})
goto end_branch_125
} else {

}
}
{
// TAST (Let): __local_var_4_87 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_4_87 := v1_1
_ = __local_var_4_87
// TAST (Let): __local_var_5_88 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_5_88 := v3_3
_ = __local_var_5_88
// TAST (Let): __local_var_6_89 shape=Other bindingType=(ADT ["Test","RBTree","Color"] [])
__local_var_6_89 := v_0
_ = __local_var_6_89
// TAST (Let): __local_var_7_90 shape=Other bindingType=Int
__local_var_7_90 := v2_2
_ = __local_var_7_90
__t125 = (&Constructor_Test_RBTree_T{1, __local_var_6_89, __local_var_4_87, __local_var_7_90, __local_var_5_88})
}
end_branch_125:
__t126 = __t125
goto end_branch_126
} else {

}
}
{
// TAST (Let): __local_var_4_18 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_4_18 := v1_1
_ = __local_var_4_18
// TAST (Let): __local_var_5_19 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_5_19 := v3_3
_ = __local_var_5_19
// TAST (Let): __local_var_6_20 shape=Other bindingType=(ADT ["Test","RBTree","Color"] [])
__local_var_6_20 := v_0
_ = __local_var_6_20
// TAST (Let): __local_var_7_21 shape=Other bindingType=Int
__local_var_7_21 := v2_2
_ = __local_var_7_21
__t126 = (&Constructor_Test_RBTree_T{1, __local_var_6_20, __local_var_4_18, __local_var_7_21, __local_var_5_19})
}
end_branch_126:
__t223 = __t126
goto end_branch_223
} else {

}
}
{
var __t_tag_127 *Constructor_Test_RBTree_T = (v1_1).V3
_ = __t_tag_127
if (__t_tag_127 != nil) {
var __t181 *Constructor_Test_RBTree_T
{
var __t_tag_132 uint32 = ((v1_1).V3).V0
_ = __t_tag_132
if (uint32(__t_tag_132) == 3668501016) {
// TAST (Let): __local_var_4_133 shape=Other bindingType=Any
__local_var_4_133 := (v1_1).V1
_ = __local_var_4_133
// TAST (Let): __local_var_5_134 shape=Other bindingType=Any
__local_var_5_134 := ((v1_1).V3).V1
_ = __local_var_5_134
// TAST (Let): __local_var_6_135 shape=Other bindingType=Any
__local_var_6_135 := ((v1_1).V3).V3
_ = __local_var_6_135
// TAST (Let): __local_var_7_136 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_7_136 := v3_3
_ = __local_var_7_136
// TAST (Let): __local_var_8_137 shape=Other bindingType=Any
__local_var_8_137 := (v1_1).V2
_ = __local_var_8_137
// TAST (Let): __local_var_9_138 shape=Other bindingType=Any
__local_var_9_138 := ((v1_1).V3).V2
_ = __local_var_9_138
// TAST (Let): __local_var_10_139 shape=Other bindingType=Int
__local_var_10_139 := v2_2
_ = __local_var_10_139
__t181 = (&Constructor_Test_RBTree_T{1, 3668501016, (&Constructor_Test_RBTree_T{1, 1583507464, __local_var_4_133, __local_var_8_137, __local_var_5_134}), __local_var_9_138, (&Constructor_Test_RBTree_T{1, 1583507464, __local_var_6_135, __local_var_10_139, __local_var_7_136})})
goto end_branch_181
} else {

}
}
{
var __t_and_141 bool = false
if (v3_3 != nil) {

var __t_tag_140 uint32 = (v3_3).V0
_ = __t_tag_140
__t_and_141 = (uint32(__t_tag_140) == 3668501016)
}
if __t_and_141 {
var __t180 *Constructor_Test_RBTree_T
{
var __t_tag_146 *Constructor_Test_RBTree_T = (v3_3).V1
_ = __t_tag_146
if (__t_tag_146 != nil) {
var __t169 *Constructor_Test_RBTree_T
{
var __t_tag_151 uint32 = ((v3_3).V1).V0
_ = __t_tag_151
if (uint32(__t_tag_151) == 3668501016) {
// TAST (Let): __local_var_4_152 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_4_152 := v1_1
_ = __local_var_4_152
// TAST (Let): __local_var_5_153 shape=Other bindingType=Any
__local_var_5_153 := ((v3_3).V1).V1
_ = __local_var_5_153
// TAST (Let): __local_var_6_154 shape=Other bindingType=Any
__local_var_6_154 := ((v3_3).V1).V3
_ = __local_var_6_154
// TAST (Let): __local_var_7_155 shape=Other bindingType=Any
__local_var_7_155 := (v3_3).V3
_ = __local_var_7_155
// TAST (Let): __local_var_8_156 shape=Other bindingType=Int
__local_var_8_156 := v2_2
_ = __local_var_8_156
// TAST (Let): __local_var_9_157 shape=Other bindingType=Any
__local_var_9_157 := ((v3_3).V1).V2
_ = __local_var_9_157
// TAST (Let): __local_var_10_158 shape=Other bindingType=Any
__local_var_10_158 := (v3_3).V2
_ = __local_var_10_158
__t169 = (&Constructor_Test_RBTree_T{1, 3668501016, (&Constructor_Test_RBTree_T{1, 1583507464, __local_var_4_152, __local_var_8_156, __local_var_5_153}), __local_var_9_157, (&Constructor_Test_RBTree_T{1, 1583507464, __local_var_6_154, __local_var_10_158, __local_var_7_155})})
goto end_branch_169
} else {

}
}
{
var __t_tag_159 *Constructor_Test_RBTree_T = (v3_3).V3
_ = __t_tag_159
var __t_and_161 bool = false
if (__t_tag_159 != nil) {

var __t_tag_160 uint32 = ((v3_3).V3).V0
_ = __t_tag_160
__t_and_161 = (uint32(__t_tag_160) == 3668501016)
}
if __t_and_161 {
// TAST (Let): __local_var_4_162 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_4_162 := v1_1
_ = __local_var_4_162
// TAST (Let): __local_var_5_163 shape=Other bindingType=Any
__local_var_5_163 := (v3_3).V1
_ = __local_var_5_163
// TAST (Let): __local_var_6_164 shape=Other bindingType=Any
__local_var_6_164 := ((v3_3).V3).V1
_ = __local_var_6_164
// TAST (Let): __local_var_7_165 shape=Other bindingType=Any
__local_var_7_165 := ((v3_3).V3).V3
_ = __local_var_7_165
// TAST (Let): __local_var_8_166 shape=Other bindingType=Int
__local_var_8_166 := v2_2
_ = __local_var_8_166
// TAST (Let): __local_var_9_167 shape=Other bindingType=Any
__local_var_9_167 := (v3_3).V2
_ = __local_var_9_167
// TAST (Let): __local_var_10_168 shape=Other bindingType=Any
__local_var_10_168 := ((v3_3).V3).V2
_ = __local_var_10_168
__t169 = (&Constructor_Test_RBTree_T{1, 3668501016, (&Constructor_Test_RBTree_T{1, 1583507464, __local_var_4_162, __local_var_8_166, __local_var_5_163}), __local_var_9_167, (&Constructor_Test_RBTree_T{1, 1583507464, __local_var_6_164, __local_var_10_168, __local_var_7_165})})
goto end_branch_169
} else {

}
}
{
// TAST (Let): __local_var_4_147 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_4_147 := v1_1
_ = __local_var_4_147
// TAST (Let): __local_var_5_148 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_5_148 := v3_3
_ = __local_var_5_148
// TAST (Let): __local_var_6_149 shape=Other bindingType=(ADT ["Test","RBTree","Color"] [])
__local_var_6_149 := v_0
_ = __local_var_6_149
// TAST (Let): __local_var_7_150 shape=Other bindingType=Int
__local_var_7_150 := v2_2
_ = __local_var_7_150
__t169 = (&Constructor_Test_RBTree_T{1, __local_var_6_149, __local_var_4_147, __local_var_7_150, __local_var_5_148})
}
end_branch_169:
__t180 = __t169
goto end_branch_180
} else {

}
}
{
var __t_tag_170 *Constructor_Test_RBTree_T = (v3_3).V3
_ = __t_tag_170
var __t_and_172 bool = false
if (__t_tag_170 != nil) {

var __t_tag_171 uint32 = ((v3_3).V3).V0
_ = __t_tag_171
__t_and_172 = (uint32(__t_tag_171) == 3668501016)
}
if __t_and_172 {
// TAST (Let): __local_var_4_173 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_4_173 := v1_1
_ = __local_var_4_173
// TAST (Let): __local_var_5_174 shape=Other bindingType=Any
__local_var_5_174 := (v3_3).V1
_ = __local_var_5_174
// TAST (Let): __local_var_6_175 shape=Other bindingType=Any
__local_var_6_175 := ((v3_3).V3).V1
_ = __local_var_6_175
// TAST (Let): __local_var_7_176 shape=Other bindingType=Any
__local_var_7_176 := ((v3_3).V3).V3
_ = __local_var_7_176
// TAST (Let): __local_var_8_177 shape=Other bindingType=Int
__local_var_8_177 := v2_2
_ = __local_var_8_177
// TAST (Let): __local_var_9_178 shape=Other bindingType=Any
__local_var_9_178 := (v3_3).V2
_ = __local_var_9_178
// TAST (Let): __local_var_10_179 shape=Other bindingType=Any
__local_var_10_179 := ((v3_3).V3).V2
_ = __local_var_10_179
__t180 = (&Constructor_Test_RBTree_T{1, 3668501016, (&Constructor_Test_RBTree_T{1, 1583507464, __local_var_4_173, __local_var_8_177, __local_var_5_174}), __local_var_9_178, (&Constructor_Test_RBTree_T{1, 1583507464, __local_var_6_175, __local_var_10_179, __local_var_7_176})})
goto end_branch_180
} else {

}
}
{
// TAST (Let): __local_var_4_142 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_4_142 := v1_1
_ = __local_var_4_142
// TAST (Let): __local_var_5_143 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_5_143 := v3_3
_ = __local_var_5_143
// TAST (Let): __local_var_6_144 shape=Other bindingType=(ADT ["Test","RBTree","Color"] [])
__local_var_6_144 := v_0
_ = __local_var_6_144
// TAST (Let): __local_var_7_145 shape=Other bindingType=Int
__local_var_7_145 := v2_2
_ = __local_var_7_145
__t180 = (&Constructor_Test_RBTree_T{1, __local_var_6_144, __local_var_4_142, __local_var_7_145, __local_var_5_143})
}
end_branch_180:
__t181 = __t180
goto end_branch_181
} else {

}
}
{
// TAST (Let): __local_var_4_128 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_4_128 := v1_1
_ = __local_var_4_128
// TAST (Let): __local_var_5_129 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_5_129 := v3_3
_ = __local_var_5_129
// TAST (Let): __local_var_6_130 shape=Other bindingType=(ADT ["Test","RBTree","Color"] [])
__local_var_6_130 := v_0
_ = __local_var_6_130
// TAST (Let): __local_var_7_131 shape=Other bindingType=Int
__local_var_7_131 := v2_2
_ = __local_var_7_131
__t181 = (&Constructor_Test_RBTree_T{1, __local_var_6_130, __local_var_4_128, __local_var_7_131, __local_var_5_129})
}
end_branch_181:
__t223 = __t181
goto end_branch_223
} else {

}
}
{
var __t_and_183 bool = false
if (v3_3 != nil) {

var __t_tag_182 uint32 = (v3_3).V0
_ = __t_tag_182
__t_and_183 = (uint32(__t_tag_182) == 3668501016)
}
if __t_and_183 {
var __t222 *Constructor_Test_RBTree_T
{
var __t_tag_188 *Constructor_Test_RBTree_T = (v3_3).V1
_ = __t_tag_188
if (__t_tag_188 != nil) {
var __t211 *Constructor_Test_RBTree_T
{
var __t_tag_193 uint32 = ((v3_3).V1).V0
_ = __t_tag_193
if (uint32(__t_tag_193) == 3668501016) {
// TAST (Let): __local_var_4_194 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_4_194 := v1_1
_ = __local_var_4_194
// TAST (Let): __local_var_5_195 shape=Other bindingType=Any
__local_var_5_195 := ((v3_3).V1).V1
_ = __local_var_5_195
// TAST (Let): __local_var_6_196 shape=Other bindingType=Any
__local_var_6_196 := ((v3_3).V1).V3
_ = __local_var_6_196
// TAST (Let): __local_var_7_197 shape=Other bindingType=Any
__local_var_7_197 := (v3_3).V3
_ = __local_var_7_197
// TAST (Let): __local_var_8_198 shape=Other bindingType=Int
__local_var_8_198 := v2_2
_ = __local_var_8_198
// TAST (Let): __local_var_9_199 shape=Other bindingType=Any
__local_var_9_199 := ((v3_3).V1).V2
_ = __local_var_9_199
// TAST (Let): __local_var_10_200 shape=Other bindingType=Any
__local_var_10_200 := (v3_3).V2
_ = __local_var_10_200
__t211 = (&Constructor_Test_RBTree_T{1, 3668501016, (&Constructor_Test_RBTree_T{1, 1583507464, __local_var_4_194, __local_var_8_198, __local_var_5_195}), __local_var_9_199, (&Constructor_Test_RBTree_T{1, 1583507464, __local_var_6_196, __local_var_10_200, __local_var_7_197})})
goto end_branch_211
} else {

}
}
{
var __t_tag_201 *Constructor_Test_RBTree_T = (v3_3).V3
_ = __t_tag_201
var __t_and_203 bool = false
if (__t_tag_201 != nil) {

var __t_tag_202 uint32 = ((v3_3).V3).V0
_ = __t_tag_202
__t_and_203 = (uint32(__t_tag_202) == 3668501016)
}
if __t_and_203 {
// TAST (Let): __local_var_4_204 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_4_204 := v1_1
_ = __local_var_4_204
// TAST (Let): __local_var_5_205 shape=Other bindingType=Any
__local_var_5_205 := (v3_3).V1
_ = __local_var_5_205
// TAST (Let): __local_var_6_206 shape=Other bindingType=Any
__local_var_6_206 := ((v3_3).V3).V1
_ = __local_var_6_206
// TAST (Let): __local_var_7_207 shape=Other bindingType=Any
__local_var_7_207 := ((v3_3).V3).V3
_ = __local_var_7_207
// TAST (Let): __local_var_8_208 shape=Other bindingType=Int
__local_var_8_208 := v2_2
_ = __local_var_8_208
// TAST (Let): __local_var_9_209 shape=Other bindingType=Any
__local_var_9_209 := (v3_3).V2
_ = __local_var_9_209
// TAST (Let): __local_var_10_210 shape=Other bindingType=Any
__local_var_10_210 := ((v3_3).V3).V2
_ = __local_var_10_210
__t211 = (&Constructor_Test_RBTree_T{1, 3668501016, (&Constructor_Test_RBTree_T{1, 1583507464, __local_var_4_204, __local_var_8_208, __local_var_5_205}), __local_var_9_209, (&Constructor_Test_RBTree_T{1, 1583507464, __local_var_6_206, __local_var_10_210, __local_var_7_207})})
goto end_branch_211
} else {

}
}
{
// TAST (Let): __local_var_4_189 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_4_189 := v1_1
_ = __local_var_4_189
// TAST (Let): __local_var_5_190 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_5_190 := v3_3
_ = __local_var_5_190
// TAST (Let): __local_var_6_191 shape=Other bindingType=(ADT ["Test","RBTree","Color"] [])
__local_var_6_191 := v_0
_ = __local_var_6_191
// TAST (Let): __local_var_7_192 shape=Other bindingType=Int
__local_var_7_192 := v2_2
_ = __local_var_7_192
__t211 = (&Constructor_Test_RBTree_T{1, __local_var_6_191, __local_var_4_189, __local_var_7_192, __local_var_5_190})
}
end_branch_211:
__t222 = __t211
goto end_branch_222
} else {

}
}
{
var __t_tag_212 *Constructor_Test_RBTree_T = (v3_3).V3
_ = __t_tag_212
var __t_and_214 bool = false
if (__t_tag_212 != nil) {

var __t_tag_213 uint32 = ((v3_3).V3).V0
_ = __t_tag_213
__t_and_214 = (uint32(__t_tag_213) == 3668501016)
}
if __t_and_214 {
// TAST (Let): __local_var_4_215 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_4_215 := v1_1
_ = __local_var_4_215
// TAST (Let): __local_var_5_216 shape=Other bindingType=Any
__local_var_5_216 := (v3_3).V1
_ = __local_var_5_216
// TAST (Let): __local_var_6_217 shape=Other bindingType=Any
__local_var_6_217 := ((v3_3).V3).V1
_ = __local_var_6_217
// TAST (Let): __local_var_7_218 shape=Other bindingType=Any
__local_var_7_218 := ((v3_3).V3).V3
_ = __local_var_7_218
// TAST (Let): __local_var_8_219 shape=Other bindingType=Int
__local_var_8_219 := v2_2
_ = __local_var_8_219
// TAST (Let): __local_var_9_220 shape=Other bindingType=Any
__local_var_9_220 := (v3_3).V2
_ = __local_var_9_220
// TAST (Let): __local_var_10_221 shape=Other bindingType=Any
__local_var_10_221 := ((v3_3).V3).V2
_ = __local_var_10_221
__t222 = (&Constructor_Test_RBTree_T{1, 3668501016, (&Constructor_Test_RBTree_T{1, 1583507464, __local_var_4_215, __local_var_8_219, __local_var_5_216}), __local_var_9_220, (&Constructor_Test_RBTree_T{1, 1583507464, __local_var_6_217, __local_var_10_221, __local_var_7_218})})
goto end_branch_222
} else {

}
}
{
// TAST (Let): __local_var_4_184 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_4_184 := v1_1
_ = __local_var_4_184
// TAST (Let): __local_var_5_185 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_5_185 := v3_3
_ = __local_var_5_185
// TAST (Let): __local_var_6_186 shape=Other bindingType=(ADT ["Test","RBTree","Color"] [])
__local_var_6_186 := v_0
_ = __local_var_6_186
// TAST (Let): __local_var_7_187 shape=Other bindingType=Int
__local_var_7_187 := v2_2
_ = __local_var_7_187
__t222 = (&Constructor_Test_RBTree_T{1, __local_var_6_186, __local_var_4_184, __local_var_7_187, __local_var_5_185})
}
end_branch_222:
__t223 = __t222
goto end_branch_223
} else {

}
}
{
// TAST (Let): __local_var_4_13 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_4_13 := v1_1
_ = __local_var_4_13
// TAST (Let): __local_var_5_14 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_5_14 := v3_3
_ = __local_var_5_14
// TAST (Let): __local_var_6_15 shape=Other bindingType=(ADT ["Test","RBTree","Color"] [])
__local_var_6_15 := v_0
_ = __local_var_6_15
// TAST (Let): __local_var_7_16 shape=Other bindingType=Int
__local_var_7_16 := v2_2
_ = __local_var_7_16
__t223 = (&Constructor_Test_RBTree_T{1, __local_var_6_15, __local_var_4_13, __local_var_7_16, __local_var_5_14})
}
end_branch_223:
__t265 = __t223
goto end_branch_265
} else {

}
}
{
var __t_and_225 bool = false
if (v3_3 != nil) {

var __t_tag_224 uint32 = (v3_3).V0
_ = __t_tag_224
__t_and_225 = (uint32(__t_tag_224) == 3668501016)
}
if __t_and_225 {
var __t264 *Constructor_Test_RBTree_T
{
var __t_tag_230 *Constructor_Test_RBTree_T = (v3_3).V1
_ = __t_tag_230
if (__t_tag_230 != nil) {
var __t253 *Constructor_Test_RBTree_T
{
var __t_tag_235 uint32 = ((v3_3).V1).V0
_ = __t_tag_235
if (uint32(__t_tag_235) == 3668501016) {
// TAST (Let): __local_var_4_236 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_4_236 := v1_1
_ = __local_var_4_236
// TAST (Let): __local_var_5_237 shape=Other bindingType=Any
__local_var_5_237 := ((v3_3).V1).V1
_ = __local_var_5_237
// TAST (Let): __local_var_6_238 shape=Other bindingType=Any
__local_var_6_238 := ((v3_3).V1).V3
_ = __local_var_6_238
// TAST (Let): __local_var_7_239 shape=Other bindingType=Any
__local_var_7_239 := (v3_3).V3
_ = __local_var_7_239
// TAST (Let): __local_var_8_240 shape=Other bindingType=Int
__local_var_8_240 := v2_2
_ = __local_var_8_240
// TAST (Let): __local_var_9_241 shape=Other bindingType=Any
__local_var_9_241 := ((v3_3).V1).V2
_ = __local_var_9_241
// TAST (Let): __local_var_10_242 shape=Other bindingType=Any
__local_var_10_242 := (v3_3).V2
_ = __local_var_10_242
__t253 = (&Constructor_Test_RBTree_T{1, 3668501016, (&Constructor_Test_RBTree_T{1, 1583507464, __local_var_4_236, __local_var_8_240, __local_var_5_237}), __local_var_9_241, (&Constructor_Test_RBTree_T{1, 1583507464, __local_var_6_238, __local_var_10_242, __local_var_7_239})})
goto end_branch_253
} else {

}
}
{
var __t_tag_243 *Constructor_Test_RBTree_T = (v3_3).V3
_ = __t_tag_243
var __t_and_245 bool = false
if (__t_tag_243 != nil) {

var __t_tag_244 uint32 = ((v3_3).V3).V0
_ = __t_tag_244
__t_and_245 = (uint32(__t_tag_244) == 3668501016)
}
if __t_and_245 {
// TAST (Let): __local_var_4_246 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_4_246 := v1_1
_ = __local_var_4_246
// TAST (Let): __local_var_5_247 shape=Other bindingType=Any
__local_var_5_247 := (v3_3).V1
_ = __local_var_5_247
// TAST (Let): __local_var_6_248 shape=Other bindingType=Any
__local_var_6_248 := ((v3_3).V3).V1
_ = __local_var_6_248
// TAST (Let): __local_var_7_249 shape=Other bindingType=Any
__local_var_7_249 := ((v3_3).V3).V3
_ = __local_var_7_249
// TAST (Let): __local_var_8_250 shape=Other bindingType=Int
__local_var_8_250 := v2_2
_ = __local_var_8_250
// TAST (Let): __local_var_9_251 shape=Other bindingType=Any
__local_var_9_251 := (v3_3).V2
_ = __local_var_9_251
// TAST (Let): __local_var_10_252 shape=Other bindingType=Any
__local_var_10_252 := ((v3_3).V3).V2
_ = __local_var_10_252
__t253 = (&Constructor_Test_RBTree_T{1, 3668501016, (&Constructor_Test_RBTree_T{1, 1583507464, __local_var_4_246, __local_var_8_250, __local_var_5_247}), __local_var_9_251, (&Constructor_Test_RBTree_T{1, 1583507464, __local_var_6_248, __local_var_10_252, __local_var_7_249})})
goto end_branch_253
} else {

}
}
{
// TAST (Let): __local_var_4_231 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_4_231 := v1_1
_ = __local_var_4_231
// TAST (Let): __local_var_5_232 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_5_232 := v3_3
_ = __local_var_5_232
// TAST (Let): __local_var_6_233 shape=Other bindingType=(ADT ["Test","RBTree","Color"] [])
__local_var_6_233 := v_0
_ = __local_var_6_233
// TAST (Let): __local_var_7_234 shape=Other bindingType=Int
__local_var_7_234 := v2_2
_ = __local_var_7_234
__t253 = (&Constructor_Test_RBTree_T{1, __local_var_6_233, __local_var_4_231, __local_var_7_234, __local_var_5_232})
}
end_branch_253:
__t264 = __t253
goto end_branch_264
} else {

}
}
{
var __t_tag_254 *Constructor_Test_RBTree_T = (v3_3).V3
_ = __t_tag_254
var __t_and_256 bool = false
if (__t_tag_254 != nil) {

var __t_tag_255 uint32 = ((v3_3).V3).V0
_ = __t_tag_255
__t_and_256 = (uint32(__t_tag_255) == 3668501016)
}
if __t_and_256 {
// TAST (Let): __local_var_4_257 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_4_257 := v1_1
_ = __local_var_4_257
// TAST (Let): __local_var_5_258 shape=Other bindingType=Any
__local_var_5_258 := (v3_3).V1
_ = __local_var_5_258
// TAST (Let): __local_var_6_259 shape=Other bindingType=Any
__local_var_6_259 := ((v3_3).V3).V1
_ = __local_var_6_259
// TAST (Let): __local_var_7_260 shape=Other bindingType=Any
__local_var_7_260 := ((v3_3).V3).V3
_ = __local_var_7_260
// TAST (Let): __local_var_8_261 shape=Other bindingType=Int
__local_var_8_261 := v2_2
_ = __local_var_8_261
// TAST (Let): __local_var_9_262 shape=Other bindingType=Any
__local_var_9_262 := (v3_3).V2
_ = __local_var_9_262
// TAST (Let): __local_var_10_263 shape=Other bindingType=Any
__local_var_10_263 := ((v3_3).V3).V2
_ = __local_var_10_263
__t264 = (&Constructor_Test_RBTree_T{1, 3668501016, (&Constructor_Test_RBTree_T{1, 1583507464, __local_var_4_257, __local_var_8_261, __local_var_5_258}), __local_var_9_262, (&Constructor_Test_RBTree_T{1, 1583507464, __local_var_6_259, __local_var_10_263, __local_var_7_260})})
goto end_branch_264
} else {

}
}
{
// TAST (Let): __local_var_4_226 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_4_226 := v1_1
_ = __local_var_4_226
// TAST (Let): __local_var_5_227 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_5_227 := v3_3
_ = __local_var_5_227
// TAST (Let): __local_var_6_228 shape=Other bindingType=(ADT ["Test","RBTree","Color"] [])
__local_var_6_228 := v_0
_ = __local_var_6_228
// TAST (Let): __local_var_7_229 shape=Other bindingType=Int
__local_var_7_229 := v2_2
_ = __local_var_7_229
__t264 = (&Constructor_Test_RBTree_T{1, __local_var_6_228, __local_var_4_226, __local_var_7_229, __local_var_5_227})
}
end_branch_264:
__t265 = __t264
goto end_branch_265
} else {

}
}
{
// TAST (Let): __local_var_4_8 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_4_8 := v1_1
_ = __local_var_4_8
// TAST (Let): __local_var_5_9 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_5_9 := v3_3
_ = __local_var_5_9
// TAST (Let): __local_var_6_10 shape=Other bindingType=(ADT ["Test","RBTree","Color"] [])
__local_var_6_10 := v_0
_ = __local_var_6_10
// TAST (Let): __local_var_7_11 shape=Other bindingType=Int
__local_var_7_11 := v2_2
_ = __local_var_7_11
__t265 = (&Constructor_Test_RBTree_T{1, __local_var_6_10, __local_var_4_8, __local_var_7_11, __local_var_5_9})
}
end_branch_265:
__t307 = __t265
goto end_branch_307
} else {

}
}
{
var __t_and_267 bool = false
if (v3_3 != nil) {

var __t_tag_266 uint32 = (v3_3).V0
_ = __t_tag_266
__t_and_267 = (uint32(__t_tag_266) == 3668501016)
}
if __t_and_267 {
var __t306 *Constructor_Test_RBTree_T
{
var __t_tag_272 *Constructor_Test_RBTree_T = (v3_3).V1
_ = __t_tag_272
if (__t_tag_272 != nil) {
var __t295 *Constructor_Test_RBTree_T
{
var __t_tag_277 uint32 = ((v3_3).V1).V0
_ = __t_tag_277
if (uint32(__t_tag_277) == 3668501016) {
// TAST (Let): __local_var_4_278 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_4_278 := v1_1
_ = __local_var_4_278
// TAST (Let): __local_var_5_279 shape=Other bindingType=Any
__local_var_5_279 := ((v3_3).V1).V1
_ = __local_var_5_279
// TAST (Let): __local_var_6_280 shape=Other bindingType=Any
__local_var_6_280 := ((v3_3).V1).V3
_ = __local_var_6_280
// TAST (Let): __local_var_7_281 shape=Other bindingType=Any
__local_var_7_281 := (v3_3).V3
_ = __local_var_7_281
// TAST (Let): __local_var_8_282 shape=Other bindingType=Int
__local_var_8_282 := v2_2
_ = __local_var_8_282
// TAST (Let): __local_var_9_283 shape=Other bindingType=Any
__local_var_9_283 := ((v3_3).V1).V2
_ = __local_var_9_283
// TAST (Let): __local_var_10_284 shape=Other bindingType=Any
__local_var_10_284 := (v3_3).V2
_ = __local_var_10_284
__t295 = (&Constructor_Test_RBTree_T{1, 3668501016, (&Constructor_Test_RBTree_T{1, 1583507464, __local_var_4_278, __local_var_8_282, __local_var_5_279}), __local_var_9_283, (&Constructor_Test_RBTree_T{1, 1583507464, __local_var_6_280, __local_var_10_284, __local_var_7_281})})
goto end_branch_295
} else {

}
}
{
var __t_tag_285 *Constructor_Test_RBTree_T = (v3_3).V3
_ = __t_tag_285
var __t_and_287 bool = false
if (__t_tag_285 != nil) {

var __t_tag_286 uint32 = ((v3_3).V3).V0
_ = __t_tag_286
__t_and_287 = (uint32(__t_tag_286) == 3668501016)
}
if __t_and_287 {
// TAST (Let): __local_var_4_288 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_4_288 := v1_1
_ = __local_var_4_288
// TAST (Let): __local_var_5_289 shape=Other bindingType=Any
__local_var_5_289 := (v3_3).V1
_ = __local_var_5_289
// TAST (Let): __local_var_6_290 shape=Other bindingType=Any
__local_var_6_290 := ((v3_3).V3).V1
_ = __local_var_6_290
// TAST (Let): __local_var_7_291 shape=Other bindingType=Any
__local_var_7_291 := ((v3_3).V3).V3
_ = __local_var_7_291
// TAST (Let): __local_var_8_292 shape=Other bindingType=Int
__local_var_8_292 := v2_2
_ = __local_var_8_292
// TAST (Let): __local_var_9_293 shape=Other bindingType=Any
__local_var_9_293 := (v3_3).V2
_ = __local_var_9_293
// TAST (Let): __local_var_10_294 shape=Other bindingType=Any
__local_var_10_294 := ((v3_3).V3).V2
_ = __local_var_10_294
__t295 = (&Constructor_Test_RBTree_T{1, 3668501016, (&Constructor_Test_RBTree_T{1, 1583507464, __local_var_4_288, __local_var_8_292, __local_var_5_289}), __local_var_9_293, (&Constructor_Test_RBTree_T{1, 1583507464, __local_var_6_290, __local_var_10_294, __local_var_7_291})})
goto end_branch_295
} else {

}
}
{
// TAST (Let): __local_var_4_273 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_4_273 := v1_1
_ = __local_var_4_273
// TAST (Let): __local_var_5_274 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_5_274 := v3_3
_ = __local_var_5_274
// TAST (Let): __local_var_6_275 shape=Other bindingType=(ADT ["Test","RBTree","Color"] [])
__local_var_6_275 := v_0
_ = __local_var_6_275
// TAST (Let): __local_var_7_276 shape=Other bindingType=Int
__local_var_7_276 := v2_2
_ = __local_var_7_276
__t295 = (&Constructor_Test_RBTree_T{1, __local_var_6_275, __local_var_4_273, __local_var_7_276, __local_var_5_274})
}
end_branch_295:
__t306 = __t295
goto end_branch_306
} else {

}
}
{
var __t_tag_296 *Constructor_Test_RBTree_T = (v3_3).V3
_ = __t_tag_296
var __t_and_298 bool = false
if (__t_tag_296 != nil) {

var __t_tag_297 uint32 = ((v3_3).V3).V0
_ = __t_tag_297
__t_and_298 = (uint32(__t_tag_297) == 3668501016)
}
if __t_and_298 {
// TAST (Let): __local_var_4_299 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_4_299 := v1_1
_ = __local_var_4_299
// TAST (Let): __local_var_5_300 shape=Other bindingType=Any
__local_var_5_300 := (v3_3).V1
_ = __local_var_5_300
// TAST (Let): __local_var_6_301 shape=Other bindingType=Any
__local_var_6_301 := ((v3_3).V3).V1
_ = __local_var_6_301
// TAST (Let): __local_var_7_302 shape=Other bindingType=Any
__local_var_7_302 := ((v3_3).V3).V3
_ = __local_var_7_302
// TAST (Let): __local_var_8_303 shape=Other bindingType=Int
__local_var_8_303 := v2_2
_ = __local_var_8_303
// TAST (Let): __local_var_9_304 shape=Other bindingType=Any
__local_var_9_304 := (v3_3).V2
_ = __local_var_9_304
// TAST (Let): __local_var_10_305 shape=Other bindingType=Any
__local_var_10_305 := ((v3_3).V3).V2
_ = __local_var_10_305
__t306 = (&Constructor_Test_RBTree_T{1, 3668501016, (&Constructor_Test_RBTree_T{1, 1583507464, __local_var_4_299, __local_var_8_303, __local_var_5_300}), __local_var_9_304, (&Constructor_Test_RBTree_T{1, 1583507464, __local_var_6_301, __local_var_10_305, __local_var_7_302})})
goto end_branch_306
} else {

}
}
{
// TAST (Let): __local_var_4_268 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_4_268 := v1_1
_ = __local_var_4_268
// TAST (Let): __local_var_5_269 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_5_269 := v3_3
_ = __local_var_5_269
// TAST (Let): __local_var_6_270 shape=Other bindingType=(ADT ["Test","RBTree","Color"] [])
__local_var_6_270 := v_0
_ = __local_var_6_270
// TAST (Let): __local_var_7_271 shape=Other bindingType=Int
__local_var_7_271 := v2_2
_ = __local_var_7_271
__t306 = (&Constructor_Test_RBTree_T{1, __local_var_6_270, __local_var_4_268, __local_var_7_271, __local_var_5_269})
}
end_branch_306:
__t307 = __t306
goto end_branch_307
} else {

}
}
{
// TAST (Let): __local_var_4_4 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_4_4 := v1_1
_ = __local_var_4_4
// TAST (Let): __local_var_5_5 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_5_5 := v3_3
_ = __local_var_5_5
// TAST (Let): __local_var_6_6 shape=Other bindingType=(ADT ["Test","RBTree","Color"] [])
__local_var_6_6 := v_0
_ = __local_var_6_6
// TAST (Let): __local_var_7_7 shape=Other bindingType=Int
__local_var_7_7 := v2_2
_ = __local_var_7_7
__t307 = (&Constructor_Test_RBTree_T{1, __local_var_6_6, __local_var_4_4, __local_var_7_7, __local_var_5_5})
}
end_branch_307:
__t308 = __t307
goto end_branch_308
} else {

}
}
{
// TAST (Let): __local_var_4_0 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_4_0 := v1_1
_ = __local_var_4_0
// TAST (Let): __local_var_5_1 shape=Other bindingType=(ADT ["Test","RBTree","Tree"] [])
__local_var_5_1 := v3_3
_ = __local_var_5_1
// TAST (Let): __local_var_6_2 shape=Other bindingType=(ADT ["Test","RBTree","Color"] [])
__local_var_6_2 := v_0
_ = __local_var_6_2
// TAST (Let): __local_var_7_3 shape=Other bindingType=Int
__local_var_7_3 := v2_2
_ = __local_var_7_3
__t308 = (&Constructor_Test_RBTree_T{1, __local_var_6_2, __local_var_4_0, __local_var_7_3, __local_var_5_1})
}
end_branch_308:
return __t308
}

func Call_Test_RBTree_ins(v_0_loop int64, v1_1_loop *Constructor_Test_RBTree_T) *Constructor_Test_RBTree_T {
ins:
for {
if false { continue ins }
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 *Constructor_Test_RBTree_T = v1_1_loop
_ = v1_1
var __t2 *Constructor_Test_RBTree_T
{
if (v1_1 == nil) {
__t2 = (&Constructor_Test_RBTree_T{1, 3668501016, (*Constructor_Test_RBTree_T)(nil), v_0, (*Constructor_Test_RBTree_T)(nil)})
goto end_branch_2
} else {

}
}
{
if (v1_1 != nil) {
var __t1 *Constructor_Test_RBTree_T
{
if (v_0) < ((v1_1).V2) {
__t1 = Call_Test_RBTree_balance((v1_1).V0, Call_Test_RBTree_ins(v_0, (v1_1).V1), (v1_1).V2, (v1_1).V3)
goto end_branch_1
} else {

}
}
{
var __t0 *Constructor_Test_RBTree_T
{
if (v_0) > ((v1_1).V2) {
__t0 = Call_Test_RBTree_balance((v1_1).V0, (v1_1).V1, (v1_1).V2, Call_Test_RBTree_ins(v_0, (v1_1).V3))
goto end_branch_0
} else {

}
}
{
__t0 = (&Constructor_Test_RBTree_T{1, (v1_1).V0, (v1_1).V1, (v1_1).V2, (v1_1).V3})
}
end_branch_0:
__t1 = __t0
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = func() *Constructor_Test_RBTree_T { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}
}

func Call_Test_RBTree_insert(x_0_loop int64, s_1_loop *Constructor_Test_RBTree_T) *Constructor_Test_RBTree_T {
var x_0 int64 = x_0_loop
_ = x_0
var s_1 *Constructor_Test_RBTree_T = s_1_loop
_ = s_1
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
__local_var_2_0 := Call_Test_RBTree_ins(x_0, s_1)
_ = __local_var_2_0
var __t2 *Constructor_Test_RBTree_T
{
if (__local_var_2_0 != nil) {
var __reuse_1 *Constructor_Test_RBTree_T
if ((__local_var_2_0) != (nil)) && (((__local_var_2_0).V0) == (1583507464)) {
__reuse_1 = __local_var_2_0
} else {
__reuse_1 = (&Constructor_Test_RBTree_T{1, 1583507464, (__local_var_2_0).V1, (__local_var_2_0).V2, (__local_var_2_0).V3})
}
__t2 = __reuse_1
goto end_branch_2
} else {

}
}
{
if (__local_var_2_0 == nil) {
__t2 = (*Constructor_Test_RBTree_T)(nil)
goto end_branch_2
} else {

}
}
{
__t2 = func() *Constructor_Test_RBTree_T { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}

func Call_Test_RBTree_buildTree(v_0_loop int64, v1_1_loop *Constructor_Test_RBTree_T) *Constructor_Test_RBTree_T {
buildTree:
for {
if false { continue buildTree }
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 *Constructor_Test_RBTree_T = v1_1_loop
_ = v1_1
var __t0 *Constructor_Test_RBTree_T
{
if (v_0) == (int64(0)) {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
v_0_loop = (v_0) - (int64(1))
v1_1_loop = Call_Test_RBTree_insert(v_0, v1_1)
continue buildTree
__t0 = func() *Constructor_Test_RBTree_T { panic("unreachable") }()
}
end_branch_0:
return __t0
}
}


