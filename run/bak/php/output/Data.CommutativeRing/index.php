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

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };




// Data_CommutativeRing_commutativeRingUnit
$GLOBALS['Data_CommutativeRing_commutativeRingUnit'] = (object)["Ring0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Ring_ringUnit'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_CommutativeRing_commutativeRingRecordNil
$GLOBALS['Data_CommutativeRing_commutativeRingRecordNil'] = (object)["RingRecord0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Ring_ringRecordNil'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_CommutativeRing_commutativeRingRecordCons
function majData_majCommutativemajRing_commutativemajRingmajRecordmajCons($dictIsSymbol_0, $_dollar__unused_1 = null, $dictCommutativeRingRecord_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majCommutativemajRing_commutativemajRingmajRecordmajCons';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $ringRecordCons1_3_0 = ((($GLOBALS['Data_Ring_ringRecordCons'])($dictIsSymbol_0))(null))((($dictCommutativeRingRecord_2)->{'RingRecord0'})(null));
  $__res = function($dictCommutativeRing_4) use ($ringRecordCons1_3_0) {
  $__num = \func_num_args();
  $ringRecordCons2_5_1 = ($ringRecordCons1_3_0)((($dictCommutativeRing_4)->{'Ring0'})(null));
  $__res = (object)["RingRecord0" => function($_dollar__unused_6) use ($ringRecordCons2_5_1) {
  $__num = \func_num_args();
  $__res = $ringRecordCons2_5_1;
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
function majData_majCommutativemajRing_commutativemajRingmajRecord($_dollar__unused_0, $dictCommutativeRingRecord_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majCommutativemajRing_commutativemajRingmajRecord';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_2_0 = (($dictCommutativeRingRecord_1)->{'RingRecord0'})(null);
  $__local_var_3_1 = (($__local_var_2_0)->{'SemiringRecord0'})(null);
  $semiringRecord1_4_2 = (object)["add" => (($__local_var_3_1)->{'addRecord'})(new \Type\Proxy\Type_Proxy_Proxy()), "mul" => (($__local_var_3_1)->{'mulRecord'})(new \Type\Proxy\Type_Proxy_Proxy()), "one" => ((($__local_var_3_1)->{'oneRecord'})(new \Type\Proxy\Type_Proxy_Proxy()))(new \Type\Proxy\Type_Proxy_Proxy()), "zero" => ((($__local_var_3_1)->{'zeroRecord'})(new \Type\Proxy\Type_Proxy_Proxy()))(new \Type\Proxy\Type_Proxy_Proxy())];
  $ringRecord1_4_2 = (object)["sub" => (($__local_var_2_0)->{'subRecord'})(new \Type\Proxy\Type_Proxy_Proxy()), "Semiring0" => function($_dollar__unused_5) use ($semiringRecord1_4_2) {
  $__num = \func_num_args();
  $__res = $semiringRecord1_4_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["Ring0" => function($_dollar__unused_5) use ($ringRecord1_4_2) {
  $__num = \func_num_args();
  $__res = $ringRecord1_4_2;
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
$GLOBALS['Data_CommutativeRing_commutativeRingProxy'] = (object)["Ring0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Ring_ringProxy'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_CommutativeRing_commutativeRingNumber
$GLOBALS['Data_CommutativeRing_commutativeRingNumber'] = (object)["Ring0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Ring_ringNumber'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_CommutativeRing_commutativeRingInt
$GLOBALS['Data_CommutativeRing_commutativeRingInt'] = (object)["Ring0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = (object)["sub" => $GLOBALS['Data_Ring_intSub'], "Semiring0" => function($_dollar__unused_1) {
  $__num = \func_num_args();
  $__res = (object)["add" => $GLOBALS['Data_Semiring_intAdd'], "zero" => 0, "mul" => $GLOBALS['Data_Semiring_intMul'], "one" => 1];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
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
  $zero1_3_2 = ($__local_var_2_1)->{'zero'};
  $one1_4_3 = ($__local_var_2_1)->{'one'};
  $semiringFn_5_4 = (object)["add" => (function() use ($__local_var_2_1) {
  $__fn = function($f_5, $g_6 = null, $x_7 = null) use ($__local_var_2_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($__local_var_2_1)->{'add'})(($f_5)($x_7)))(($g_6)($x_7));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "zero" => function($v_5) use ($zero1_3_2) {
  $__num = \func_num_args();
  $__res = $zero1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "mul" => (function() use ($__local_var_2_1) {
  $__fn = function($f_5, $g_6 = null, $x_7 = null) use ($__local_var_2_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($__local_var_2_1)->{'mul'})(($f_5)($x_7)))(($g_6)($x_7));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "one" => function($v_5) use ($one1_4_3) {
  $__num = \func_num_args();
  $__res = $one1_4_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $ringFn_3_2 = (object)["sub" => (function() use ($__local_var_1_0) {
  $__fn = function($f_6, $g_7 = null, $x_8 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($__local_var_1_0)->{'sub'})(($f_6)($x_8)))(($g_7)($x_8));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "Semiring0" => function($_dollar__unused_6) use ($semiringFn_5_4) {
  $__num = \func_num_args();
  $__res = $semiringFn_5_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["Ring0" => function($_dollar__unused_4) use ($ringFn_3_2) {
  $__num = \func_num_args();
  $__res = $ringFn_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_CommutativeRing_commutativeRingFn'] = __NAMESPACE__ . '\\majData_majCommutativemajRing_commutativemajRingmajFn';

