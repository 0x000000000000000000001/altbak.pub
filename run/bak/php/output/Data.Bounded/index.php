<?php

namespace Data\Bounded;

// ALL IMPORTS: Data.Bounded, Data.Ord, Data.Ordering, Data.Symbol, Data.Unit, Prim, Prim.Row, Prim.RowList, Record.Unsafe, Type.Proxy
// TO REQUIRE: Data.Bounded, Data.Ord, Data.Ordering, Data.Symbol, Data.Unit, Record.Unsafe, Type.Proxy
require_once __DIR__ . '/../Data.Bounded/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ordering/index.php';
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
\PhpursThunks::$thunks['Data_Bounded_topRecord'] = function() { $v = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)->{'topRecord'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Bounded_top'] = function() { $v = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)->{'top'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Bounded_boundedUnit'] = function() { $v = (object)["top" => ($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit')), "bottom" => ($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit')), "Ord0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Ord_ordUnit'] ?? \PhpursThunks::eval('Data_Ord_ordUnit'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Bounded_boundedRecordNil'] = function() { $v = (object)["topRecord" => (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (object)[];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "bottomRecord" => (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (object)[];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "OrdRecord0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Ord_ordRecordNil'] ?? \PhpursThunks::eval('Data_Ord_ordRecordNil'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Bounded_boundedProxy'] = function() { $v = (object)["bottom" => new Phpurs_Data0("Proxy"), "top" => new Phpurs_Data0("Proxy"), "Ord0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Ord_ordProxy'] ?? \PhpursThunks::eval('Data_Ord_ordProxy'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Bounded_boundedOrdering'] = function() { $v = (object)["top" => new Phpurs_Data0("GT"), "bottom" => new Phpurs_Data0("LT"), "Ord0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Ord_ordOrdering'] ?? \PhpursThunks::eval('Data_Ord_ordOrdering'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Bounded_boundedNumber'] = function() { $v = (object)["top" => ($GLOBALS['Data_Bounded_topNumber'] ?? \PhpursThunks::eval('Data_Bounded_topNumber')), "bottom" => ($GLOBALS['Data_Bounded_bottomNumber'] ?? \PhpursThunks::eval('Data_Bounded_bottomNumber')), "Ord0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Ord_ordNumber'] ?? \PhpursThunks::eval('Data_Ord_ordNumber'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Bounded_boundedInt'] = function() { $v = (object)["top" => ($GLOBALS['Data_Bounded_topInt'] ?? \PhpursThunks::eval('Data_Bounded_topInt')), "bottom" => ($GLOBALS['Data_Bounded_bottomInt'] ?? \PhpursThunks::eval('Data_Bounded_bottomInt')), "Ord0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Bounded_boundedChar'] = function() { $v = (object)["top" => ($GLOBALS['Data_Bounded_topChar'] ?? \PhpursThunks::eval('Data_Bounded_topChar')), "bottom" => ($GLOBALS['Data_Bounded_bottomChar'] ?? \PhpursThunks::eval('Data_Bounded_bottomChar')), "Ord0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Ord_ordChar'] ?? \PhpursThunks::eval('Data_Ord_ordChar'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Bounded_boundedBoolean'] = function() { $v = (object)["top" => true, "bottom" => false, "Ord0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Ord_ordBoolean'] ?? \PhpursThunks::eval('Data_Ord_ordBoolean'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Bounded_bottomRecord'] = function() { $v = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)->{'bottomRecord'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Bounded_boundedRecord'] = function() { $v = (function() {
  $__fn = function($dollar__unused_0 = null, $dictBoundedRecord_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $ordRecord1_2_0 = ((($GLOBALS['Data_Ord_ordRecord'] ?? \PhpursThunks::eval('Data_Ord_ordRecord')))(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))((($dictBoundedRecord_1)->{'OrdRecord0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = (object)["top" => ((($dictBoundedRecord_1)->{'topRecord'})(new Phpurs_Data0("Proxy")))(new Phpurs_Data0("Proxy")), "bottom" => ((($dictBoundedRecord_1)->{'bottomRecord'})(new Phpurs_Data0("Proxy")))(new Phpurs_Data0("Proxy")), "Ord0" => function($dollar__unused_3 = null) use ($ordRecord1_2_0) {
  $__num = \func_num_args();
  $__res = $ordRecord1_2_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Bounded_bottom'] = function() { $v = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)->{'bottom'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Bounded_boundedRecordCons'] = function() { $v = (function() {
  $__fn = function($dictIsSymbol_0 = null, $dictBounded_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $top1_2_0 = ($dictBounded_1)->{'top'};
  $bottom1_3_1 = ($dictBounded_1)->{'bottom'};
  $Ord0_4_2 = (($dictBounded_1)->{'Ord0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $__res = (function() use ($Ord0_4_2, $bottom1_3_1, $dictIsSymbol_0, $top1_2_0) {
  $__fn = function($dollar__unused_5 = null, $dollar__unused_6 = null, $dictBoundedRecord_7 = null) use ($Ord0_4_2, $bottom1_3_1, $dictIsSymbol_0, $top1_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $ordRecordCons_8_3 = ((((($GLOBALS['Data_Ord_ordRecordCons'] ?? \PhpursThunks::eval('Data_Ord_ordRecordCons')))((($dictBoundedRecord_7)->{'OrdRecord0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')))))(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))($dictIsSymbol_0))($Ord0_4_2);
  $__res = (object)["topRecord" => (function() use ($dictBoundedRecord_7, $dictIsSymbol_0, $top1_2_0) {
  $__fn = function($v_9 = null, $rowProxy_10 = null) use ($dictBoundedRecord_7, $dictIsSymbol_0, $top1_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($GLOBALS['Record_Unsafe_unsafeSet'] ?? \PhpursThunks::eval('Record_Unsafe_unsafeSet')))((($dictIsSymbol_0)->{'reflectSymbol'})(new Phpurs_Data0("Proxy"))))($top1_2_0))(((($dictBoundedRecord_7)->{'topRecord'})(new Phpurs_Data0("Proxy")))($rowProxy_10));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "bottomRecord" => (function() use ($bottom1_3_1, $dictBoundedRecord_7, $dictIsSymbol_0) {
  $__fn = function($v_9 = null, $rowProxy_10 = null) use ($bottom1_3_1, $dictBoundedRecord_7, $dictIsSymbol_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($GLOBALS['Record_Unsafe_unsafeSet'] ?? \PhpursThunks::eval('Record_Unsafe_unsafeSet')))((($dictIsSymbol_0)->{'reflectSymbol'})(new Phpurs_Data0("Proxy"))))($bottom1_3_1))(((($dictBoundedRecord_7)->{'bottomRecord'})(new Phpurs_Data0("Proxy")))($rowProxy_10));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "OrdRecord0" => function($dollar__unused_9 = null) use ($ordRecordCons_8_3) {
  $__num = \func_num_args();
  $__res = $ordRecordCons_8_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };
$ffi_Data_Bounded = \call_user_func(function() {
  $exports = [];
$topInt = 2147483647;
$bottomInt = -2147483648;
$topChar = "\u{10FFFF}";
$bottomChar = "\u{0000}";
$topNumber = INF;
$bottomNumber = -INF;

$exports['topInt'] = $topInt;
$exports['bottomInt'] = $bottomInt;
$exports['topChar'] = $topChar;
$exports['bottomChar'] = $bottomChar;
$exports['topNumber'] = $topNumber;
$exports['bottomNumber'] = $bottomNumber;
return $exports;
  return $exports;
});
\PhpursThunks::$thunks['Data_Bounded_bottomChar'] = function() use (&$ffi_Data_Bounded) { return $ffi_Data_Bounded['bottomChar']; };
\PhpursThunks::$thunks['Data_Bounded_bottomInt'] = function() use (&$ffi_Data_Bounded) { return $ffi_Data_Bounded['bottomInt']; };
\PhpursThunks::$thunks['Data_Bounded_bottomNumber'] = function() use (&$ffi_Data_Bounded) { return $ffi_Data_Bounded['bottomNumber']; };
\PhpursThunks::$thunks['Data_Bounded_topChar'] = function() use (&$ffi_Data_Bounded) { return $ffi_Data_Bounded['topChar']; };
\PhpursThunks::$thunks['Data_Bounded_topInt'] = function() use (&$ffi_Data_Bounded) { return $ffi_Data_Bounded['topInt']; };
\PhpursThunks::$thunks['Data_Bounded_topNumber'] = function() use (&$ffi_Data_Bounded) { return $ffi_Data_Bounded['topNumber']; };
















