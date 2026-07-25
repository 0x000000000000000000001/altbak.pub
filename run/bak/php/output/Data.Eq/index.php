<?php

namespace Data\Eq;

// ALL IMPORTS: Data.Eq, Data.HeytingAlgebra, Data.Symbol, Data.Unit, Data.Void, Prim, Prim.Row, Prim.RowList, Record.Unsafe, Type.Proxy
// TO REQUIRE: Data.Eq, Data.HeytingAlgebra, Data.Symbol, Data.Unit, Data.Void, Record.Unsafe, Type.Proxy
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.HeytingAlgebra/index.php';
require_once __DIR__ . '/../Data.Symbol/index.php';
require_once __DIR__ . '/../Data.Unit/index.php';
require_once __DIR__ . '/../Data.Void/index.php';
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
\PhpursThunks::$thunks['Data_Eq_eqVoid'] = function() { $v = (object)["eq" => (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = true;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Data_Eq_eqUnit'] = function() { $v = (object)["eq" => (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = true;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Data_Eq_eqString'] = function() { $v = (object)["eq" => ($GLOBALS['Data_Eq_eqStringImpl'] ?? \PhpursThunks::eval('Data_Eq_eqStringImpl'))]; return $v; };
\PhpursThunks::$thunks['Data_Eq_eqRowNil'] = function() { $v = (object)["eqRecord" => (function() {
  $__fn = function($v_0 = null, $v1_1 = null, $v2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = true;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Data_Eq_eqRecord'] = function() { $v = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)->{'eqRecord'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Eq_eqRec'] = function() { $v = (function() {
  $__fn = function($dollar__unused_0 = null, $dictEqRecord_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (object)["eq" => (($dictEqRecord_1)->{'eqRecord'})(new Phpurs_Data0("Proxy"))];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Eq_eqProxy'] = function() { $v = (object)["eq" => (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = true;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Data_Eq_eqNumber'] = function() { $v = (object)["eq" => ($GLOBALS['Data_Eq_eqNumberImpl'] ?? \PhpursThunks::eval('Data_Eq_eqNumberImpl'))]; return $v; };
\PhpursThunks::$thunks['Data_Eq_eqInt'] = function() { $v = (object)["eq" => ($GLOBALS['Data_Eq_eqIntImpl'] ?? \PhpursThunks::eval('Data_Eq_eqIntImpl'))]; return $v; };
\PhpursThunks::$thunks['Data_Eq_eqChar'] = function() { $v = (object)["eq" => ($GLOBALS['Data_Eq_eqCharImpl'] ?? \PhpursThunks::eval('Data_Eq_eqCharImpl'))]; return $v; };
\PhpursThunks::$thunks['Data_Eq_eqBoolean'] = function() { $v = (object)["eq" => ($GLOBALS['Data_Eq_eqBooleanImpl'] ?? \PhpursThunks::eval('Data_Eq_eqBooleanImpl'))]; return $v; };
\PhpursThunks::$thunks['Data_Eq_eq1'] = function() { $v = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)->{'eq1'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Eq_eq'] = function() { $v = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)->{'eq'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Eq_eqArray'] = function() { $v = function($dictEq_0 = null) {
  $__num = \func_num_args();
  $__res = (object)["eq" => (($GLOBALS['Data_Eq_eqArrayImpl'] ?? \PhpursThunks::eval('Data_Eq_eqArrayImpl')))(($dictEq_0)->{'eq'})];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Eq_eq1Array'] = function() { $v = (object)["eq1" => function($dictEq_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Eq_eqArrayImpl'] ?? \PhpursThunks::eval('Data_Eq_eqArrayImpl')))(($dictEq_0)->{'eq'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Eq_eqRowCons'] = function() { $v = (function() {
  $__fn = function($dictEqRecord_0 = null, $dollar__unused_1 = null, $dictIsSymbol_2 = null, $dictEq_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = (object)["eqRecord" => (function() use ($dictEqRecord_0, $dictEq_3, $dictIsSymbol_2) {
  $__fn = function($v_4 = null, $ra_5 = null, $rb_6 = null) use ($dictEqRecord_0, $dictEq_3, $dictIsSymbol_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $get_7_0 = (($GLOBALS['Record_Unsafe_unsafeGet'] ?? \PhpursThunks::eval('Record_Unsafe_unsafeGet')))((($dictIsSymbol_2)->{'reflectSymbol'})(new Phpurs_Data0("Proxy")));
  $__res = ((($GLOBALS['Data_HeytingAlgebra_boolConj'] ?? \PhpursThunks::eval('Data_HeytingAlgebra_boolConj')))(((($dictEq_3)->{'eq'})(($get_7_0)($ra_5)))(($get_7_0)($rb_6))))((((($dictEqRecord_0)->{'eqRecord'})(new Phpurs_Data0("Proxy")))($ra_5))($rb_6));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Eq_notEq'] = function() { $v = (function() {
  $__fn = function($dictEq_0 = null, $x_1 = null, $y_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($GLOBALS['Data_Eq_eqBooleanImpl'] ?? \PhpursThunks::eval('Data_Eq_eqBooleanImpl')))(((($dictEq_0)->{'eq'})($x_1))($y_2)))(false);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Eq_notEq1'] = function() { $v = (function() {
  $__fn = function($dictEq1_0 = null, $dictEq_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $eq12_2_0 = (($dictEq1_0)->{'eq1'})($dictEq_1);
  $__res = (function() use ($eq12_2_0) {
  $__fn = function($x_3 = null, $y_4 = null) use ($eq12_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_Eq_eqBooleanImpl'] ?? \PhpursThunks::eval('Data_Eq_eqBooleanImpl')))((($eq12_2_0)($x_3))($y_4)))(false);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
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
$ffi_Data_Eq = \call_user_func(function() {
  $exports = [];
$eqIntImpl = function($a, $b = null) use (&$eqIntImpl) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$eqIntImpl) {

            return $eqIntImpl(...\array_merge($__args, $more));
        };
    }
    return $a === $b;
};
$eqStringImpl = $eqIntImpl;
$eqNumberImpl = $eqIntImpl;
$eqCharImpl = $eqIntImpl;
$eqBooleanImpl = $eqIntImpl;

$eqArrayImpl = function($f, $xs = null, $ys = null) use (&$eqArrayImpl) {
    if (\func_num_args() < 3) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$eqArrayImpl) {
            return $eqArrayImpl(...\array_merge($__args, $more));
        };
    }
    if (\count($xs) !== \count($ys)) return false;
    for ($i = 0; $i < \count($xs); $i++) {
        if (!$f($xs[$i], $ys[$i])) return false;
    }
    return true;
};

$exports['eqIntImpl'] = $eqIntImpl;
$exports['eqStringImpl'] = $eqStringImpl;
$exports['eqNumberImpl'] = $eqNumberImpl;
$exports['eqCharImpl'] = $eqCharImpl;
$exports['eqBooleanImpl'] = $eqBooleanImpl;
$exports['eqArrayImpl'] = $eqArrayImpl;
return $exports;
  return $exports;
});
\PhpursThunks::$thunks['Data_Eq_eqArrayImpl'] = function() use (&$ffi_Data_Eq) { return $ffi_Data_Eq['eqArrayImpl']; };
\PhpursThunks::$thunks['Data_Eq_eqBooleanImpl'] = function() use (&$ffi_Data_Eq) { return $ffi_Data_Eq['eqBooleanImpl']; };
\PhpursThunks::$thunks['Data_Eq_eqCharImpl'] = function() use (&$ffi_Data_Eq) { return $ffi_Data_Eq['eqCharImpl']; };
\PhpursThunks::$thunks['Data_Eq_eqIntImpl'] = function() use (&$ffi_Data_Eq) { return $ffi_Data_Eq['eqIntImpl']; };
\PhpursThunks::$thunks['Data_Eq_eqNumberImpl'] = function() use (&$ffi_Data_Eq) { return $ffi_Data_Eq['eqNumberImpl']; };
\PhpursThunks::$thunks['Data_Eq_eqStringImpl'] = function() use (&$ffi_Data_Eq) { return $ffi_Data_Eq['eqStringImpl']; };




















