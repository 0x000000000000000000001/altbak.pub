<?php

namespace Data\Ring;

// ALL IMPORTS: Data.Ring, Data.Semiring, Data.Symbol, Data.Unit, Prim, Prim.Row, Prim.RowList, Record.Unsafe, Type.Proxy
// TO REQUIRE: Data.Ring, Data.Semiring, Data.Symbol, Data.Unit, Record.Unsafe, Type.Proxy
require_once __DIR__ . '/../Data.Ring/index.php';
require_once __DIR__ . '/../Data.Semiring/index.php';
require_once __DIR__ . '/../Data.Symbol/index.php';
require_once __DIR__ . '/../Data.Unit/index.php';
require_once __DIR__ . '/../Record.Unsafe/index.php';
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
$ffi_Data_Ring = \call_user_func(function() {
  $exports = [];
$intSub = function($a, $b) use (&$intSub) {
    return (($a - $b) << 32) >> 32;
};
$numSub = function($a, $b) use (&$numSub) {
    return (float)($a - $b);
};

$exports['intSub'] = $intSub;
$exports['numSub'] = $numSub;
return $exports;
  return $exports;
});
function majData_majRing_intmajSub(int $v0, $v1 = null): int|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majRing_intmajSub';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Data_Ring;
  $f = (\array_key_exists('intSub', $ffi_Data_Ring) ? $ffi_Data_Ring['intSub'] : new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Data_Ring_intSub'] = __NAMESPACE__ . '\\majData_majRing_intmajSub';

function majData_majRing_nummajSub(float $v0, $v1 = null): float|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majRing_nummajSub';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Data_Ring;
  $f = (\array_key_exists('numSub', $ffi_Data_Ring) ? $ffi_Data_Ring['numSub'] : new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Data_Ring_numSub'] = __NAMESPACE__ . '\\majData_majRing_nummajSub';





// Data_Ring_subRecord
function majData_majRing_submajRecord($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majRing_submajRecord';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'subRecord'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Ring_subRecord'] = __NAMESPACE__ . '\\majData_majRing_submajRecord';

// Data_Ring_sub
function majData_majRing_sub($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majRing_sub';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'sub'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Ring_sub'] = __NAMESPACE__ . '\\majData_majRing_sub';

// Data_Ring_ringUnit
$GLOBALS['Data_Ring_ringUnit'] = (object)["sub" => function($v_0) {
  $__num = \func_num_args();
  $__res = function($v1_1) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Unit_unit'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Semiring0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Semiring_semiringUnit'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Ring_ringRecordNil
$GLOBALS['Data_Ring_ringRecordNil'] = (object)["subRecord" => function($v_0) {
  $__num = \func_num_args();
  $__res = function($v1_1) {
  $__num = \func_num_args();
  $__res = function($v2_2) {
  $__num = \func_num_args();
  $__res = (object)[];
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
}, "SemiringRecord0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Semiring_semiringRecordNil'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Ring_ringRecordCons
function majData_majRing_ringmajRecordmajCons($dictIsSymbol_0, $_dollar___unused_1 = null, $dictRingRecord_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majRing_ringmajRecordmajCons';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__local_var_3_0 = (($dictRingRecord_2)->{'SemiringRecord0'})(null);
  $__res = function($dictRing_4) use ($__local_var_3_0, $dictIsSymbol_0, $dictRingRecord_2) {
  $__num = \func_num_args();
  $__local_var_5_1 = (($dictRing_4)->{'Semiring0'})(null);
  $one1_6_2 = ($__local_var_5_1)->{'one'};
  $zero1_7_3 = ($__local_var_5_1)->{'zero'};
  $semiringRecordCons2_5_1 = (object)["addRecord" => function($v_8) use ($__local_var_3_0, $__local_var_5_1, $dictIsSymbol_0) {
  $__num = \func_num_args();
  $__res = function($ra_9) use ($__local_var_3_0, $__local_var_5_1, $dictIsSymbol_0) {
  $__num = \func_num_args();
  $__res = function($rb_10) use ($__local_var_3_0, $__local_var_5_1, $dictIsSymbol_0, $ra_9) {
  $__num = \func_num_args();
  $key_11_4 = (($dictIsSymbol_0)->{'reflectSymbol'})(new \Type\Proxy\Type_Proxy_Proxy());
  $get_12_5 = ($GLOBALS['Record_Unsafe_unsafeGet'])($key_11_4);
  $__res = \Record\Unsafe\majRecord_majUnsafe_unsafemajSet($key_11_4, ((($__local_var_5_1)->{'add'})(($get_12_5)($ra_9)))(($get_12_5)($rb_10)), (((($__local_var_3_0)->{'addRecord'})(new \Type\Proxy\Type_Proxy_Proxy()))($ra_9))($rb_10));
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
}, "mulRecord" => function($v_8) use ($__local_var_3_0, $__local_var_5_1, $dictIsSymbol_0) {
  $__num = \func_num_args();
  $__res = function($ra_9) use ($__local_var_3_0, $__local_var_5_1, $dictIsSymbol_0) {
  $__num = \func_num_args();
  $__res = function($rb_10) use ($__local_var_3_0, $__local_var_5_1, $dictIsSymbol_0, $ra_9) {
  $__num = \func_num_args();
  $key_11_6 = (($dictIsSymbol_0)->{'reflectSymbol'})(new \Type\Proxy\Type_Proxy_Proxy());
  $get_12_7 = ($GLOBALS['Record_Unsafe_unsafeGet'])($key_11_6);
  $__res = \Record\Unsafe\majRecord_majUnsafe_unsafemajSet($key_11_6, ((($__local_var_5_1)->{'mul'})(($get_12_7)($ra_9)))(($get_12_7)($rb_10)), (((($__local_var_3_0)->{'mulRecord'})(new \Type\Proxy\Type_Proxy_Proxy()))($ra_9))($rb_10));
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
}, "oneRecord" => function($v_8) use ($__local_var_3_0, $dictIsSymbol_0, $one1_6_2) {
  $__num = \func_num_args();
  $__res = function($v1_9) use ($__local_var_3_0, $dictIsSymbol_0, $one1_6_2) {
  $__num = \func_num_args();
  $__res = \Record\Unsafe\majRecord_majUnsafe_unsafemajSet((($dictIsSymbol_0)->{'reflectSymbol'})(new \Type\Proxy\Type_Proxy_Proxy()), $one1_6_2, ((($__local_var_3_0)->{'oneRecord'})(new \Type\Proxy\Type_Proxy_Proxy()))(new \Type\Proxy\Type_Proxy_Proxy()));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "zeroRecord" => function($v_8) use ($__local_var_3_0, $dictIsSymbol_0, $zero1_7_3) {
  $__num = \func_num_args();
  $__res = function($v1_9) use ($__local_var_3_0, $dictIsSymbol_0, $zero1_7_3) {
  $__num = \func_num_args();
  $__res = \Record\Unsafe\majRecord_majUnsafe_unsafemajSet((($dictIsSymbol_0)->{'reflectSymbol'})(new \Type\Proxy\Type_Proxy_Proxy()), $zero1_7_3, ((($__local_var_3_0)->{'zeroRecord'})(new \Type\Proxy\Type_Proxy_Proxy()))(new \Type\Proxy\Type_Proxy_Proxy()));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["subRecord" => function($v_6) use ($dictIsSymbol_0, $dictRingRecord_2, $dictRing_4) {
  $__num = \func_num_args();
  $__res = function($ra_7) use ($dictIsSymbol_0, $dictRingRecord_2, $dictRing_4) {
  $__num = \func_num_args();
  $__res = function($rb_8) use ($dictIsSymbol_0, $dictRingRecord_2, $dictRing_4, $ra_7) {
  $__num = \func_num_args();
  $key_9_9 = (($dictIsSymbol_0)->{'reflectSymbol'})(new \Type\Proxy\Type_Proxy_Proxy());
  $get_10_10 = ($GLOBALS['Record_Unsafe_unsafeGet'])($key_9_9);
  $__res = \Record\Unsafe\majRecord_majUnsafe_unsafemajSet($key_9_9, ((($dictRing_4)->{'sub'})(($get_10_10)($ra_7)))(($get_10_10)($rb_8)), (((($dictRingRecord_2)->{'subRecord'})(new \Type\Proxy\Type_Proxy_Proxy()))($ra_7))($rb_8));
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
}, "SemiringRecord0" => function($_dollar___unused_6) use ($semiringRecordCons2_5_1) {
  $__num = \func_num_args();
  $__res = $semiringRecordCons2_5_1;
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
$GLOBALS['Data_Ring_ringRecordCons'] = __NAMESPACE__ . '\\majData_majRing_ringmajRecordmajCons';

// Data_Ring_ringRecord
function majData_majRing_ringmajRecord($_dollar___unused_0, $dictRingRecord_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majRing_ringmajRecord';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_2_0 = (($dictRingRecord_1)->{'SemiringRecord0'})(null);
  $semiringRecord1_2_0 = (object)["add" => (($__local_var_2_0)->{'addRecord'})(new \Type\Proxy\Type_Proxy_Proxy()), "mul" => (($__local_var_2_0)->{'mulRecord'})(new \Type\Proxy\Type_Proxy_Proxy()), "one" => ((($__local_var_2_0)->{'oneRecord'})(new \Type\Proxy\Type_Proxy_Proxy()))(new \Type\Proxy\Type_Proxy_Proxy()), "zero" => ((($__local_var_2_0)->{'zeroRecord'})(new \Type\Proxy\Type_Proxy_Proxy()))(new \Type\Proxy\Type_Proxy_Proxy())];
  $__res = (object)["sub" => (($dictRingRecord_1)->{'subRecord'})(new \Type\Proxy\Type_Proxy_Proxy()), "Semiring0" => function($_dollar___unused_3) use ($semiringRecord1_2_0) {
  $__num = \func_num_args();
  $__res = $semiringRecord1_2_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Ring_ringRecord'] = __NAMESPACE__ . '\\majData_majRing_ringmajRecord';

// Data_Ring_ringProxy
$GLOBALS['Data_Ring_ringProxy'] = (object)["sub" => function($v_0) {
  $__num = \func_num_args();
  $__res = function($v1_1) {
  $__num = \func_num_args();
  $__res = new \Type\Proxy\Type_Proxy_Proxy();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Semiring0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Semiring_semiringProxy'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Ring_ringNumber
$GLOBALS['Data_Ring_ringNumber'] = (object)["sub" => $GLOBALS['Data_Ring_numSub'], "Semiring0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Semiring_semiringNumber'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Ring_ringInt
$GLOBALS['Data_Ring_ringInt'] = (object)["sub" => $GLOBALS['Data_Ring_intSub'], "Semiring0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Semiring_semiringInt'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Ring_ringFn
function majData_majRing_ringmajFn($dictRing_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majRing_ringmajFn';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictRing_0)->{'Semiring0'})(null);
  $semiringFn_1_0 = (object)["add" => function($f_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($g_3) use ($__local_var_1_0, $f_2) {
  $__num = \func_num_args();
  $__res = function($x_4) use ($__local_var_1_0, $f_2, $g_3) {
  $__num = \func_num_args();
  $__res = ((($__local_var_1_0)->{'add'})(($f_2)($x_4)))(($g_3)($x_4));
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
}, "zero" => function($v_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = ($__local_var_1_0)->{'zero'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "mul" => function($f_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($g_3) use ($__local_var_1_0, $f_2) {
  $__num = \func_num_args();
  $__res = function($x_4) use ($__local_var_1_0, $f_2, $g_3) {
  $__num = \func_num_args();
  $__res = ((($__local_var_1_0)->{'mul'})(($f_2)($x_4)))(($g_3)($x_4));
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
}, "one" => function($v_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = ($__local_var_1_0)->{'one'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["sub" => function($f_2) use ($dictRing_0) {
  $__num = \func_num_args();
  $__res = function($g_3) use ($dictRing_0, $f_2) {
  $__num = \func_num_args();
  $__res = function($x_4) use ($dictRing_0, $f_2, $g_3) {
  $__num = \func_num_args();
  $__res = ((($dictRing_0)->{'sub'})(($f_2)($x_4)))(($g_3)($x_4));
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
}, "Semiring0" => function($_dollar___unused_2) use ($semiringFn_1_0) {
  $__num = \func_num_args();
  $__res = $semiringFn_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Ring_ringFn'] = __NAMESPACE__ . '\\majData_majRing_ringmajFn';

// Data_Ring_negate
function majData_majRing_negate($dictRing_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majRing_negate';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Semiring0_1_0 = (($dictRing_0)->{'Semiring0'})(null);
  $__res = function($a_2) use ($Semiring0_1_0, $dictRing_0) {
  $__num = \func_num_args();
  $__res = ((($dictRing_0)->{'sub'})(($Semiring0_1_0)->{'zero'}))($a_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Ring_negate'] = __NAMESPACE__ . '\\majData_majRing_negate';

