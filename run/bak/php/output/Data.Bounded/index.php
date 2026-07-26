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
$GLOBALS['Data_Bounded_bottomChar'] = $ffi_Data_Bounded['bottomChar'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_Bounded_bottomInt'] = $ffi_Data_Bounded['bottomInt'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_Bounded_bottomNumber'] = $ffi_Data_Bounded['bottomNumber'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_Bounded_topChar'] = $ffi_Data_Bounded['topChar'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_Bounded_topInt'] = $ffi_Data_Bounded['topInt'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_Bounded_topNumber'] = $ffi_Data_Bounded['topNumber'] ?? new class { public function __invoke(...$args) { return $this; } };


// Data_Bounded_topRecord
$GLOBALS['Data_Bounded_topRecord'] = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)['topRecord'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Bounded_top
$GLOBALS['Data_Bounded_top'] = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)['top'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Bounded_boundedUnit
$GLOBALS['Data_Bounded_boundedUnit'] = ["top" => $GLOBALS['Data_Unit_unit'], "bottom" => $GLOBALS['Data_Unit_unit'], "Ord0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Ord_ordUnit'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Bounded_boundedRecordNil
$GLOBALS['Data_Bounded_boundedRecordNil'] = ["topRecord" => (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = [];
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
  $__res = [];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "OrdRecord0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Ord_ordRecordNil'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Bounded_boundedProxy
$GLOBALS['Data_Bounded_boundedProxy'] = ["bottom" => new Phpurs_Data0("Proxy"), "top" => new Phpurs_Data0("Proxy"), "Ord0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Ord_ordProxy'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Bounded_boundedOrdering
$GLOBALS['Data_Bounded_boundedOrdering'] = ["top" => new Phpurs_Data0("GT"), "bottom" => new Phpurs_Data0("LT"), "Ord0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Ord_ordOrdering'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Bounded_boundedNumber
$GLOBALS['Data_Bounded_boundedNumber'] = ["top" => $GLOBALS['Data_Bounded_topNumber'], "bottom" => $GLOBALS['Data_Bounded_bottomNumber'], "Ord0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Ord_ordNumber'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Bounded_boundedInt
$GLOBALS['Data_Bounded_boundedInt'] = ["top" => $GLOBALS['Data_Bounded_topInt'], "bottom" => $GLOBALS['Data_Bounded_bottomInt'], "Ord0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Ord_ordInt'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Bounded_boundedChar
$GLOBALS['Data_Bounded_boundedChar'] = ["top" => $GLOBALS['Data_Bounded_topChar'], "bottom" => $GLOBALS['Data_Bounded_bottomChar'], "Ord0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Ord_ordChar'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Bounded_boundedBoolean
$GLOBALS['Data_Bounded_boundedBoolean'] = ["top" => true, "bottom" => false, "Ord0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Ord_ordBoolean'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Bounded_bottomRecord
$GLOBALS['Data_Bounded_bottomRecord'] = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)['bottomRecord'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Bounded_boundedRecord
$GLOBALS['Data_Bounded_boundedRecord'] = (function() {
  $__fn = function($dollar__unused_0 = null, $dictBoundedRecord_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $ordRecord1_2_0 = (($GLOBALS['Data_Ord_ordRecord'])($GLOBALS['Prim_undefined']))((($dictBoundedRecord_1)['OrdRecord0'])($GLOBALS['Prim_undefined']));
  $__res = ["top" => ((($dictBoundedRecord_1)['topRecord'])(new Phpurs_Data0("Proxy")))(new Phpurs_Data0("Proxy")), "bottom" => ((($dictBoundedRecord_1)['bottomRecord'])(new Phpurs_Data0("Proxy")))(new Phpurs_Data0("Proxy")), "Ord0" => function($dollar__unused_3 = null) use ($ordRecord1_2_0) {
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
})();

// Data_Bounded_bottom
$GLOBALS['Data_Bounded_bottom'] = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)['bottom'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Bounded_boundedRecordCons
$GLOBALS['Data_Bounded_boundedRecordCons'] = (function() {
  $__fn = function($dictIsSymbol_0 = null, $dictBounded_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $top1_2_0 = ($dictBounded_1)['top'];
  $bottom1_3_1 = ($dictBounded_1)['bottom'];
  $Ord0_4_2 = (($dictBounded_1)['Ord0'])($GLOBALS['Prim_undefined']);
  $__res = (function() use ($Ord0_4_2, $bottom1_3_1, $dictIsSymbol_0, $top1_2_0) {
  $__fn = function($dollar__unused_5 = null, $dollar__unused_6 = null, $dictBoundedRecord_7 = null) use ($Ord0_4_2, $bottom1_3_1, $dictIsSymbol_0, $top1_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $ordRecordCons_8_3 = (((($GLOBALS['Data_Ord_ordRecordCons'])((($dictBoundedRecord_7)['OrdRecord0'])($GLOBALS['Prim_undefined'])))($GLOBALS['Prim_undefined']))($dictIsSymbol_0))($Ord0_4_2);
  $__res = ["topRecord" => (function() use ($dictBoundedRecord_7, $dictIsSymbol_0, $top1_2_0) {
  $__fn = function($v_9 = null, $rowProxy_10 = null) use ($dictBoundedRecord_7, $dictIsSymbol_0, $top1_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Record_Unsafe_unsafeSet'])((($dictIsSymbol_0)['reflectSymbol'])(new Phpurs_Data0("Proxy"))))($top1_2_0))(((($dictBoundedRecord_7)['topRecord'])(new Phpurs_Data0("Proxy")))($rowProxy_10));
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
  $__res = ((($GLOBALS['Record_Unsafe_unsafeSet'])((($dictIsSymbol_0)['reflectSymbol'])(new Phpurs_Data0("Proxy"))))($bottom1_3_1))(((($dictBoundedRecord_7)['bottomRecord'])(new Phpurs_Data0("Proxy")))($rowProxy_10));
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
})();

