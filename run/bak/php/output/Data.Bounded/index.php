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
$GLOBALS['Data_Bounded_bottomChar'] = ($ffi_Data_Bounded['bottomChar'] ?? new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Bounded_bottomInt'] = ($ffi_Data_Bounded['bottomInt'] ?? new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Bounded_bottomNumber'] = ($ffi_Data_Bounded['bottomNumber'] ?? new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Bounded_topChar'] = ($ffi_Data_Bounded['topChar'] ?? new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Bounded_topInt'] = ($ffi_Data_Bounded['topInt'] ?? new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Bounded_topNumber'] = ($ffi_Data_Bounded['topNumber'] ?? new class { public function __invoke(...$args) { return $this; } });




// Data_Bounded_topRecord
function majData_majBounded_topmajRecord($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBounded_topmajRecord';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'topRecord'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Bounded_topRecord'] = __NAMESPACE__ . '\\majData_majBounded_topmajRecord';

// Data_Bounded_top
function majData_majBounded_top($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBounded_top';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'top'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Bounded_top'] = __NAMESPACE__ . '\\majData_majBounded_top';

// Data_Bounded_boundedUnit
$GLOBALS['Data_Bounded_boundedUnit'] = (object)["top" => $GLOBALS['Data_Unit_unit'], "bottom" => $GLOBALS['Data_Unit_unit'], "Ord0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Ord_ordUnit'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Bounded_boundedRecordNil
$GLOBALS['Data_Bounded_boundedRecordNil'] = (object)["topRecord" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
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
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
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
})(), "OrdRecord0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Ord_ordRecordNil'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Bounded_boundedProxy
$GLOBALS['Data_Bounded_boundedProxy'] = (object)["bottom" => new \Type\Proxy\Type_Proxy_Proxy(), "top" => new \Type\Proxy\Type_Proxy_Proxy(), "Ord0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Ord_ordProxy'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Bounded_boundedOrdering
$GLOBALS['Data_Bounded_boundedOrdering'] = (object)["top" => new \Data\Ordering\Data_Ordering_GT(), "bottom" => new \Data\Ordering\Data_Ordering_LT(), "Ord0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Ord_ordOrdering'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Bounded_boundedNumber
$GLOBALS['Data_Bounded_boundedNumber'] = (object)["top" => $GLOBALS['Data_Bounded_topNumber'], "bottom" => $GLOBALS['Data_Bounded_bottomNumber'], "Ord0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Ord_ordNumber'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Bounded_boundedInt
$GLOBALS['Data_Bounded_boundedInt'] = (object)["top" => $GLOBALS['Data_Bounded_topInt'], "bottom" => $GLOBALS['Data_Bounded_bottomInt'], "Ord0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = (object)["compare" => ((($GLOBALS['Data_Ord_ordIntImpl'])(new \Data\Ordering\Data_Ordering_LT()))(new \Data\Ordering\Data_Ordering_EQ()))(new \Data\Ordering\Data_Ordering_GT()), "Eq0" => function($_dollar__unused_1) {
  $__num = \func_num_args();
  $__res = (object)["eq" => $GLOBALS['Data_Eq_eqIntImpl']];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Bounded_boundedChar
$GLOBALS['Data_Bounded_boundedChar'] = (object)["top" => $GLOBALS['Data_Bounded_topChar'], "bottom" => $GLOBALS['Data_Bounded_bottomChar'], "Ord0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Ord_ordChar'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Bounded_boundedBoolean
$GLOBALS['Data_Bounded_boundedBoolean'] = (object)["top" => true, "bottom" => false, "Ord0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Ord_ordBoolean'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Bounded_bottomRecord
function majData_majBounded_bottommajRecord($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBounded_bottommajRecord';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'bottomRecord'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Bounded_bottomRecord'] = __NAMESPACE__ . '\\majData_majBounded_bottommajRecord';

// Data_Bounded_boundedRecord
function majData_majBounded_boundedmajRecord($_dollar__unused_0, $dictBoundedRecord_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBounded_boundedmajRecord';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_2_0 = (($dictBoundedRecord_1)->{'OrdRecord0'})(null);
  $eqRec1_3_1 = (object)["eq" => (((($__local_var_2_0)->{'EqRecord0'})(null))->{'eqRecord'})(new \Type\Proxy\Type_Proxy_Proxy())];
  $ordRecord1_4_2 = (object)["compare" => (($__local_var_2_0)->{'compareRecord'})(new \Type\Proxy\Type_Proxy_Proxy()), "Eq0" => function($_dollar__unused_4) use ($eqRec1_3_1) {
  $__num = \func_num_args();
  $__res = $eqRec1_3_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["top" => ((($dictBoundedRecord_1)->{'topRecord'})(new \Type\Proxy\Type_Proxy_Proxy()))(new \Type\Proxy\Type_Proxy_Proxy()), "bottom" => ((($dictBoundedRecord_1)->{'bottomRecord'})(new \Type\Proxy\Type_Proxy_Proxy()))(new \Type\Proxy\Type_Proxy_Proxy()), "Ord0" => function($_dollar__unused_5) use ($ordRecord1_4_2) {
  $__num = \func_num_args();
  $__res = $ordRecord1_4_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Bounded_boundedRecord'] = __NAMESPACE__ . '\\majData_majBounded_boundedmajRecord';

// Data_Bounded_bottom
function majData_majBounded_bottom($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBounded_bottom';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'bottom'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Bounded_bottom'] = __NAMESPACE__ . '\\majData_majBounded_bottom';

// Data_Bounded_boundedRecordCons
function majData_majBounded_boundedmajRecordmajCons($dictIsSymbol_0, $dictBounded_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBounded_boundedmajRecordmajCons';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $top1_2_0 = ($dictBounded_1)->{'top'};
  $bottom1_3_1 = ($dictBounded_1)->{'bottom'};
  $Ord0_4_2 = (($dictBounded_1)->{'Ord0'})(null);
  $__res = (function() use ($Ord0_4_2, $bottom1_3_1, $dictIsSymbol_0, $top1_2_0) {
  $__fn = function($_dollar__unused_5, $_dollar__unused_6 = null, $dictBoundedRecord_7 = null) use ($Ord0_4_2, $bottom1_3_1, $dictIsSymbol_0, $top1_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__local_var_8_3 = (($dictBoundedRecord_7)->{'OrdRecord0'})(null);
  $__local_var_9_4 = (($__local_var_8_3)->{'EqRecord0'})(null);
  $__local_var_10_5 = (($Ord0_4_2)->{'Eq0'})(null);
  $eqRowCons2_11_6 = (object)["eqRecord" => (function() use ($__local_var_10_5, $__local_var_9_4, $dictIsSymbol_0) {
  $__fn = function($v_11, $ra_12 = null, $rb_13 = null) use ($__local_var_10_5, $__local_var_9_4, $dictIsSymbol_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $get_14_6 = ($GLOBALS['Record_Unsafe_unsafeGet'])((($dictIsSymbol_0)->{'reflectSymbol'})(new \Type\Proxy\Type_Proxy_Proxy()));
  $__res = ((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'conj'})(((($__local_var_10_5)->{'eq'})(($get_14_6)($ra_12)))(($get_14_6)($rb_13))))((((($__local_var_9_4)->{'eqRecord'})(new \Type\Proxy\Type_Proxy_Proxy()))($ra_12))($rb_13));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()];
  $ordRecordCons_11_6 = (object)["compareRecord" => (function() use ($Ord0_4_2, $__local_var_8_3, $dictIsSymbol_0) {
  $__fn = function($v_12, $ra_13 = null, $rb_14 = null) use ($Ord0_4_2, $__local_var_8_3, $dictIsSymbol_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $key_15_8 = (($dictIsSymbol_0)->{'reflectSymbol'})(new \Type\Proxy\Type_Proxy_Proxy());
  $left_16_9 = ((($Ord0_4_2)->{'compare'})(\Record\Unsafe\majRecord_majUnsafe_unsafemajGet($key_15_8, $ra_13)))(\Record\Unsafe\majRecord_majUnsafe_unsafemajGet($key_15_8, $rb_14));
  $__t10 = null;;
  if (((($GLOBALS['Data_Eq_eqBoolean'])->{'eq'})(((($GLOBALS['Data_Ordering_eqOrdering'])->{'eq'})($left_16_9))(new \Data\Ordering\Data_Ordering_EQ())))(false)) {
$__t10 = $left_16_9;
goto end_branch_10;;
};
  $__t10 = (((($__local_var_8_3)->{'compareRecord'})(new \Type\Proxy\Type_Proxy_Proxy()))($ra_13))($rb_14);
  end_branch_10:;
  $__res = $__t10;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "EqRecord0" => function($_dollar__unused_12) use ($eqRowCons2_11_6) {
  $__num = \func_num_args();
  $__res = $eqRowCons2_11_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["topRecord" => (function() use ($dictBoundedRecord_7, $dictIsSymbol_0, $top1_2_0) {
  $__fn = function($v_12, $rowProxy_13 = null) use ($dictBoundedRecord_7, $dictIsSymbol_0, $top1_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Record\Unsafe\majRecord_majUnsafe_unsafemajSet((($dictIsSymbol_0)->{'reflectSymbol'})(new \Type\Proxy\Type_Proxy_Proxy()), $top1_2_0, ((($dictBoundedRecord_7)->{'topRecord'})(new \Type\Proxy\Type_Proxy_Proxy()))($rowProxy_13));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "bottomRecord" => (function() use ($bottom1_3_1, $dictBoundedRecord_7, $dictIsSymbol_0) {
  $__fn = function($v_12, $rowProxy_13 = null) use ($bottom1_3_1, $dictBoundedRecord_7, $dictIsSymbol_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Record\Unsafe\majRecord_majUnsafe_unsafemajSet((($dictIsSymbol_0)->{'reflectSymbol'})(new \Type\Proxy\Type_Proxy_Proxy()), $bottom1_3_1, ((($dictBoundedRecord_7)->{'bottomRecord'})(new \Type\Proxy\Type_Proxy_Proxy()))($rowProxy_13));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "OrdRecord0" => function($_dollar__unused_12) use ($ordRecordCons_11_6) {
  $__num = \func_num_args();
  $__res = $ordRecordCons_11_6;
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
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Bounded_boundedRecordCons'] = __NAMESPACE__ . '\\majData_majBounded_boundedmajRecordmajCons';

