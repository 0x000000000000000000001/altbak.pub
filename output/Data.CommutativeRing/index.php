<?php

namespace Data\CommutativeRing;

// ALL IMPORTS: Data.CommutativeRing, Data.Ring, Data.Semiring, Data.Symbol, Data.Unit, Prim, Prim.Row, Prim.RowList, Type.Proxy
// TO REQUIRE: Data.CommutativeRing, Data.Ring, Data.Semiring, Data.Symbol, Data.Unit, Type.Proxy
require_once __DIR__ . '/../Data.CommutativeRing/index.php';
require_once __DIR__ . '/../Data.Ring/index.php';
require_once __DIR__ . '/../Data.Semiring/index.php';
require_once __DIR__ . '/../Data.Symbol/index.php';
require_once __DIR__ . '/../Data.Unit/index.php';
require_once __DIR__ . '/../Type.Proxy/index.php';

if (!class_exists(__NAMESPACE__ . '\\Phpurs_Data0')) {
  class Phpurs_Data0 { public $tag; public function __construct($t) { $this->tag = $t; } }
  class Phpurs_Data1 { public $tag; public $value0; public function __construct($t, $value0) { $this->tag = $t; $this->value0 = $value0; } }
  class Phpurs_Data2 { public $tag; public $value0, $value1; public function __construct($t, $value0, $value1) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; } }
  class Phpurs_Data3 { public $tag; public $value0, $value1, $value2; public function __construct($t, $value0, $value1, $value2) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; } }
  class Phpurs_Data4 { public $tag; public $value0, $value1, $value2, $value3; public function __construct($t, $value0, $value1, $value2, $value3) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; $this->value3 = $value3; } }
  class Phpurs_Data5 { public $tag; public $value0, $value1, $value2, $value3, $value4; public function __construct($t, $value0, $value1, $value2, $value3, $value4) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; $this->value3 = $value3; $this->value4 = $value4; } }
  class Phpurs_Data6 { public $tag; public $value0, $value1, $value2, $value3, $value4, $value5; public function __construct($t, $value0, $value1, $value2, $value3, $value4, $value5) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; $this->value3 = $value3; $this->value4 = $value4; $this->value5 = $value5; } }
  class Phpurs_Data7 { public $tag; public $value0, $value1, $value2, $value3, $value4, $value5, $value6; public function __construct($t, $value0, $value1, $value2, $value3, $value4, $value5, $value6) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; $this->value3 = $value3; $this->value4 = $value4; $this->value5 = $value5; $this->value6 = $value6; } }
  class Phpurs_Data8 { public $tag; public $value0, $value1, $value2, $value3, $value4, $value5, $value6, $value7; public function __construct($t, $value0, $value1, $value2, $value3, $value4, $value5, $value6, $value7) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; $this->value3 = $value3; $this->value4 = $value4; $this->value5 = $value5; $this->value6 = $value6; $this->value7 = $value7; } }
  class Phpurs_Data9 { public $tag; public $value0, $value1, $value2, $value3, $value4, $value5, $value6, $value7, $value8; public function __construct($t, $value0, $value1, $value2, $value3, $value4, $value5, $value6, $value7, $value8) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; $this->value3 = $value3; $this->value4 = $value4; $this->value5 = $value5; $this->value6 = $value6; $this->value7 = $value7; $this->value8 = $value8; } }
  class Phpurs_Data10 { public $tag; public $value0, $value1, $value2, $value3, $value4, $value5, $value6, $value7, $value8, $value9; public function __construct($t, $value0, $value1, $value2, $value3, $value4, $value5, $value6, $value7, $value8, $value9) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; $this->value3 = $value3; $this->value4 = $value4; $this->value5 = $value5; $this->value6 = $value6; $this->value7 = $value7; $this->value8 = $value8; $this->value9 = $value9; } }
  class Phpurs_Data11 { public $tag; public $value0, $value1, $value2, $value3, $value4, $value5, $value6, $value7, $value8, $value9, $value10; public function __construct($t, $value0, $value1, $value2, $value3, $value4, $value5, $value6, $value7, $value8, $value9, $value10) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; $this->value3 = $value3; $this->value4 = $value4; $this->value5 = $value5; $this->value6 = $value6; $this->value7 = $value7; $this->value8 = $value8; $this->value9 = $value9; $this->value10 = $value10; } }
  class Phpurs_Data12 { public $tag; public $value0, $value1, $value2, $value3, $value4, $value5, $value6, $value7, $value8, $value9, $value10, $value11; public function __construct($t, $value0, $value1, $value2, $value3, $value4, $value5, $value6, $value7, $value8, $value9, $value10, $value11) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; $this->value3 = $value3; $this->value4 = $value4; $this->value5 = $value5; $this->value6 = $value6; $this->value7 = $value7; $this->value8 = $value8; $this->value9 = $value9; $this->value10 = $value10; $this->value11 = $value11; } }
}
if (!\function_exists(__NAMESPACE__ . '\\phpurs_curry_fallback')) {
  function phpurs_curry_fallback($fn, $args, $expected) {
    $missing = $expected - \count($args);
    if ($missing === 1) {
      return function($a) use ($fn, $args, $expected) {
        $num = \func_num_args();
        if ($num > 1) {
          $merged = \array_merge($args, \func_get_args());
          $res = $fn(...\array_slice($merged, 0, $expected));
          return $res(...\array_slice($merged, $expected));
        }
        $args[] = $a;
        return $fn(...$args);
      };
    }
    if ($missing === 2) {
      return function($a, $b = null) use ($fn, $args, $expected) {
        $num = \func_num_args();
        if ($num === 1) { $args[] = $a; return phpurs_curry_fallback($fn, $args, $expected); }
        if ($num > 2) {
          $merged = \array_merge($args, \func_get_args());
          $res = $fn(...\array_slice($merged, 0, $expected));
          return $res(...\array_slice($merged, $expected));
        }
        $args[] = $a; $args[] = $b;
        return $fn(...$args);
      };
    }
    if ($missing === 3) {
      return function($a, $b = null, $c = null) use ($fn, $args, $expected) {
        $num = \func_num_args();
        if ($num === 1) { $args[] = $a; return phpurs_curry_fallback($fn, $args, $expected); }
        if ($num === 2) { $args[] = $a; $args[] = $b; return phpurs_curry_fallback($fn, $args, $expected); }
        if ($num > 3) {
          $merged = \array_merge($args, \func_get_args());
          $res = $fn(...\array_slice($merged, 0, $expected));
          return $res(...\array_slice($merged, $expected));
        }
        $args[] = $a; $args[] = $b; $args[] = $c;
        return $fn(...$args);
      };
    }
    if ($missing === 4) {
      return function($a, $b = null, $c = null, $d = null) use ($fn, $args, $expected) {
        $num = \func_num_args();
        if ($num === 1) { $args[] = $a; return phpurs_curry_fallback($fn, $args, $expected); }
        if ($num === 2) { $args[] = $a; $args[] = $b; return phpurs_curry_fallback($fn, $args, $expected); }
        if ($num === 3) { $args[] = $a; $args[] = $b; $args[] = $c; return phpurs_curry_fallback($fn, $args, $expected); }
        if ($num > 4) {
          $merged = \array_merge($args, \func_get_args());
          $res = $fn(...\array_slice($merged, 0, $expected));
          return $res(...\array_slice($merged, $expected));
        }
        $args[] = $a; $args[] = $b; $args[] = $c; $args[] = $d;
        return $fn(...$args);
      };
    }
    return function(...$more) use ($fn, $args, $expected) {
      $merged = \array_merge($args, $more);
      if (\count($merged) >= $expected) {
        $res = $fn(...\array_slice($merged, 0, $expected));
        if (\count($merged) > $expected) {
          return $res(...\array_slice($merged, $expected));
        }
        return $res;
      }
      return phpurs_curry_fallback($fn, $merged, $expected);
    };
  }
}
if (!\function_exists(__NAMESPACE__ . '\\phpurs_execute_effect')) {
  function phpurs_execute_effect($val) {
    if (\is_callable($val)) {
      return $val($GLOBALS['Data_Unit_unit']);
    }
    return $val;
  }
}

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };




// Data_CommutativeRing_commutativeRingUnit
$GLOBALS['Data_CommutativeRing_commutativeRingUnit'] = (object)["Ring0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Ring_ringUnit'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_CommutativeRing_commutativeRingRecordNil
$GLOBALS['Data_CommutativeRing_commutativeRingRecordNil'] = (object)["RingRecord0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Ring_ringRecordNil'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_CommutativeRing_commutativeRingRecordCons
function majData_majCommutativemajRing_commutativemajRingmajRecordmajCons($dictIsSymbol_0, $_dollar___unused_1 = null, $dictCommutativeRingRecord_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majCommutativemajRing_commutativemajRingmajRecordmajCons';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__local_var_3_0 = (($dictCommutativeRingRecord_2)->{'RingRecord0'})(null);
  $__local_var_4_1 = (($__local_var_3_0)->{'SemiringRecord0'})(null);
  $__res = function($dictCommutativeRing_5) use ($__local_var_3_0, $__local_var_4_1, $dictIsSymbol_0) {
  $__num = \func_num_args();
  $__local_var_6_2 = (($dictCommutativeRing_5)->{'Ring0'})(null);
  $__local_var_7_3 = (($__local_var_6_2)->{'Semiring0'})(null);
  $one1_8_4 = ($__local_var_7_3)->{'one'};
  $zero1_9_5 = ($__local_var_7_3)->{'zero'};
  $semiringRecordCons2_7_3 = (object)["addRecord" => function($v_10) use ($__local_var_4_1, $__local_var_7_3, $dictIsSymbol_0) {
  $__num = \func_num_args();
  $__res = function($ra_11) use ($__local_var_4_1, $__local_var_7_3, $dictIsSymbol_0) {
  $__num = \func_num_args();
  $__res = function($rb_12) use ($__local_var_4_1, $__local_var_7_3, $dictIsSymbol_0, $ra_11) {
  $__num = \func_num_args();
  $key_13_6 = (($dictIsSymbol_0)->{'reflectSymbol'})(new \Type\Proxy\Type_Proxy_Proxy());
  $get_14_7 = ($GLOBALS['Record_Unsafe_unsafeGet'])($key_13_6);
  $__res = \Record\Unsafe\majRecord_majUnsafe_unsafemajSet($key_13_6, ((($__local_var_7_3)->{'add'})(($get_14_7)($ra_11)))(($get_14_7)($rb_12)), (((($__local_var_4_1)->{'addRecord'})(new \Type\Proxy\Type_Proxy_Proxy()))($ra_11))($rb_12));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "mulRecord" => function($v_10) use ($__local_var_4_1, $__local_var_7_3, $dictIsSymbol_0) {
  $__num = \func_num_args();
  $__res = function($ra_11) use ($__local_var_4_1, $__local_var_7_3, $dictIsSymbol_0) {
  $__num = \func_num_args();
  $__res = function($rb_12) use ($__local_var_4_1, $__local_var_7_3, $dictIsSymbol_0, $ra_11) {
  $__num = \func_num_args();
  $key_13_8 = (($dictIsSymbol_0)->{'reflectSymbol'})(new \Type\Proxy\Type_Proxy_Proxy());
  $get_14_9 = ($GLOBALS['Record_Unsafe_unsafeGet'])($key_13_8);
  $__res = \Record\Unsafe\majRecord_majUnsafe_unsafemajSet($key_13_8, ((($__local_var_7_3)->{'mul'})(($get_14_9)($ra_11)))(($get_14_9)($rb_12)), (((($__local_var_4_1)->{'mulRecord'})(new \Type\Proxy\Type_Proxy_Proxy()))($ra_11))($rb_12));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "oneRecord" => function($v_10) use ($__local_var_4_1, $dictIsSymbol_0, $one1_8_4) {
  $__num = \func_num_args();
  $__res = function($v1_11) use ($__local_var_4_1, $dictIsSymbol_0, $one1_8_4) {
  $__num = \func_num_args();
  $__res = \Record\Unsafe\majRecord_majUnsafe_unsafemajSet((($dictIsSymbol_0)->{'reflectSymbol'})(new \Type\Proxy\Type_Proxy_Proxy()), $one1_8_4, ((($__local_var_4_1)->{'oneRecord'})(new \Type\Proxy\Type_Proxy_Proxy()))(new \Type\Proxy\Type_Proxy_Proxy()));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "zeroRecord" => function($v_10) use ($__local_var_4_1, $dictIsSymbol_0, $zero1_9_5) {
  $__num = \func_num_args();
  $__res = function($v1_11) use ($__local_var_4_1, $dictIsSymbol_0, $zero1_9_5) {
  $__num = \func_num_args();
  $__res = \Record\Unsafe\majRecord_majUnsafe_unsafemajSet((($dictIsSymbol_0)->{'reflectSymbol'})(new \Type\Proxy\Type_Proxy_Proxy()), $zero1_9_5, ((($__local_var_4_1)->{'zeroRecord'})(new \Type\Proxy\Type_Proxy_Proxy()))(new \Type\Proxy\Type_Proxy_Proxy()));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $ringRecordCons2_6_2 = (object)["subRecord" => function($v_8) use ($__local_var_3_0, $__local_var_6_2, $dictIsSymbol_0) {
  $__num = \func_num_args();
  $__res = function($ra_9) use ($__local_var_3_0, $__local_var_6_2, $dictIsSymbol_0) {
  $__num = \func_num_args();
  $__res = function($rb_10) use ($__local_var_3_0, $__local_var_6_2, $dictIsSymbol_0, $ra_9) {
  $__num = \func_num_args();
  $key_11_11 = (($dictIsSymbol_0)->{'reflectSymbol'})(new \Type\Proxy\Type_Proxy_Proxy());
  $get_12_12 = ($GLOBALS['Record_Unsafe_unsafeGet'])($key_11_11);
  $__res = \Record\Unsafe\majRecord_majUnsafe_unsafemajSet($key_11_11, ((($__local_var_6_2)->{'sub'})(($get_12_12)($ra_9)))(($get_12_12)($rb_10)), (((($__local_var_3_0)->{'subRecord'})(new \Type\Proxy\Type_Proxy_Proxy()))($ra_9))($rb_10));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "SemiringRecord0" => function($_dollar___unused_8) use ($semiringRecordCons2_7_3) {
  $__num = \func_num_args();
  $__res = $semiringRecordCons2_7_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["RingRecord0" => function($_dollar___unused_7) use ($ringRecordCons2_6_2) {
  $__num = \func_num_args();
  $__res = $ringRecordCons2_6_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_CommutativeRing_commutativeRingRecordCons'] = __NAMESPACE__ . '\\majData_majCommutativemajRing_commutativemajRingmajRecordmajCons';

// Data_CommutativeRing_commutativeRingRecord
function majData_majCommutativemajRing_commutativemajRingmajRecord($_dollar___unused_0, $dictCommutativeRingRecord_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majCommutativemajRing_commutativemajRingmajRecord';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_2_0 = (($dictCommutativeRingRecord_1)->{'RingRecord0'})(null);
  $__local_var_3_1 = (($__local_var_2_0)->{'SemiringRecord0'})(null);
  $semiringRecord1_3_1 = (object)["add" => (($__local_var_3_1)->{'addRecord'})(new \Type\Proxy\Type_Proxy_Proxy()), "mul" => (($__local_var_3_1)->{'mulRecord'})(new \Type\Proxy\Type_Proxy_Proxy()), "one" => ((($__local_var_3_1)->{'oneRecord'})(new \Type\Proxy\Type_Proxy_Proxy()))(new \Type\Proxy\Type_Proxy_Proxy()), "zero" => ((($__local_var_3_1)->{'zeroRecord'})(new \Type\Proxy\Type_Proxy_Proxy()))(new \Type\Proxy\Type_Proxy_Proxy())];
  $ringRecord1_2_0 = (object)["sub" => (($__local_var_2_0)->{'subRecord'})(new \Type\Proxy\Type_Proxy_Proxy()), "Semiring0" => function($_dollar___unused_4) use ($semiringRecord1_3_1) {
  $__num = \func_num_args();
  $__res = $semiringRecord1_3_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["Ring0" => function($_dollar___unused_3) use ($ringRecord1_2_0) {
  $__num = \func_num_args();
  $__res = $ringRecord1_2_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_CommutativeRing_commutativeRingRecord'] = __NAMESPACE__ . '\\majData_majCommutativemajRing_commutativemajRingmajRecord';

// Data_CommutativeRing_commutativeRingProxy
$GLOBALS['Data_CommutativeRing_commutativeRingProxy'] = (object)["Ring0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Ring_ringProxy'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_CommutativeRing_commutativeRingNumber
$GLOBALS['Data_CommutativeRing_commutativeRingNumber'] = (object)["Ring0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Ring_ringNumber'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_CommutativeRing_commutativeRingInt
$GLOBALS['Data_CommutativeRing_commutativeRingInt'] = (object)["Ring0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Ring_ringInt'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_CommutativeRing_commutativeRingFn
function majData_majCommutativemajRing_commutativemajRingmajFn($dictCommutativeRing_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majCommutativemajRing_commutativemajRingmajFn';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictCommutativeRing_0)->{'Ring0'})(null);
  $__local_var_2_1 = (($__local_var_1_0)->{'Semiring0'})(null);
  $semiringFn_2_1 = (object)["add" => function($f_3) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = function($g_4) use ($__local_var_2_1, $f_3) {
  $__num = \func_num_args();
  $__res = function($x_5) use ($__local_var_2_1, $f_3, $g_4) {
  $__num = \func_num_args();
  $__res = ((($__local_var_2_1)->{'add'})(($f_3)($x_5)))(($g_4)($x_5));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "zero" => function($v_3) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = ($__local_var_2_1)->{'zero'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "mul" => function($f_3) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = function($g_4) use ($__local_var_2_1, $f_3) {
  $__num = \func_num_args();
  $__res = function($x_5) use ($__local_var_2_1, $f_3, $g_4) {
  $__num = \func_num_args();
  $__res = ((($__local_var_2_1)->{'mul'})(($f_3)($x_5)))(($g_4)($x_5));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "one" => function($v_3) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = ($__local_var_2_1)->{'one'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $ringFn_1_0 = (object)["sub" => function($f_3) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($g_4) use ($__local_var_1_0, $f_3) {
  $__num = \func_num_args();
  $__res = function($x_5) use ($__local_var_1_0, $f_3, $g_4) {
  $__num = \func_num_args();
  $__res = ((($__local_var_1_0)->{'sub'})(($f_3)($x_5)))(($g_4)($x_5));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Semiring0" => function($_dollar___unused_3) use ($semiringFn_2_1) {
  $__num = \func_num_args();
  $__res = $semiringFn_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["Ring0" => function($_dollar___unused_2) use ($ringFn_1_0) {
  $__num = \func_num_args();
  $__res = $ringFn_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_CommutativeRing_commutativeRingFn'] = __NAMESPACE__ . '\\majData_majCommutativemajRing_commutativemajRingmajFn';

