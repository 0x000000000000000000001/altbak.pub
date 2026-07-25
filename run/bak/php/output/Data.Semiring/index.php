<?php

namespace Data\Semiring;

// ALL IMPORTS: Data.Semiring, Data.Symbol, Data.Unit, Prim, Prim.Row, Prim.RowList, Record.Unsafe, Type.Proxy
// TO REQUIRE: Data.Semiring, Data.Symbol, Data.Unit, Record.Unsafe, Type.Proxy
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
\PhpursThunks::$thunks['Data_Semiring_zeroRecord'] = function() { $v = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)->{'zeroRecord'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Semiring_zero'] = function() { $v = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)->{'zero'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Semiring_semiringUnit'] = function() { $v = (object)["add" => (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit'));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "zero" => ($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit')), "mul" => (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit'));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "one" => ($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit'))]; return $v; };
\PhpursThunks::$thunks['Data_Semiring_semiringRecordNil'] = function() { $v = (object)["addRecord" => (function() {
  $__fn = function($v_0 = null, $v1_1 = null, $v2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (object)[];
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "mulRecord" => (function() {
  $__fn = function($v_0 = null, $v1_1 = null, $v2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (object)[];
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "oneRecord" => (function() {
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
})(), "zeroRecord" => (function() {
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
})()]; return $v; };
\PhpursThunks::$thunks['Data_Semiring_semiringProxy'] = function() { $v = (object)["add" => (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data0("Proxy");
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "mul" => (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data0("Proxy");
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "one" => new Phpurs_Data0("Proxy"), "zero" => new Phpurs_Data0("Proxy")]; return $v; };
\PhpursThunks::$thunks['Data_Semiring_semiringNumber'] = function() { $v = (object)["add" => ($GLOBALS['Data_Semiring_numAdd'] ?? \PhpursThunks::eval('Data_Semiring_numAdd')), "zero" => 0.0, "mul" => ($GLOBALS['Data_Semiring_numMul'] ?? \PhpursThunks::eval('Data_Semiring_numMul')), "one" => 1.0]; return $v; };
\PhpursThunks::$thunks['Data_Semiring_semiringInt'] = function() { $v = (object)["add" => ($GLOBALS['Data_Semiring_intAdd'] ?? \PhpursThunks::eval('Data_Semiring_intAdd')), "zero" => 0, "mul" => ($GLOBALS['Data_Semiring_intMul'] ?? \PhpursThunks::eval('Data_Semiring_intMul')), "one" => 1]; return $v; };
\PhpursThunks::$thunks['Data_Semiring_oneRecord'] = function() { $v = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)->{'oneRecord'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Semiring_one'] = function() { $v = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)->{'one'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Semiring_mulRecord'] = function() { $v = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)->{'mulRecord'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Semiring_mul'] = function() { $v = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)->{'mul'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Semiring_addRecord'] = function() { $v = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)->{'addRecord'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Semiring_semiringRecord'] = function() { $v = (function() {
  $__fn = function($dollar__unused_0 = null, $dictSemiringRecord_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (object)["add" => (($dictSemiringRecord_1)->{'addRecord'})(new Phpurs_Data0("Proxy")), "mul" => (($dictSemiringRecord_1)->{'mulRecord'})(new Phpurs_Data0("Proxy")), "one" => ((($dictSemiringRecord_1)->{'oneRecord'})(new Phpurs_Data0("Proxy")))(new Phpurs_Data0("Proxy")), "zero" => ((($dictSemiringRecord_1)->{'zeroRecord'})(new Phpurs_Data0("Proxy")))(new Phpurs_Data0("Proxy"))];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Semiring_add'] = function() { $v = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)->{'add'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Semiring_semiringFn'] = function() { $v = function($dictSemiring_0 = null) {
  $__num = \func_num_args();
  $zero1_1_0 = ($dictSemiring_0)->{'zero'};
  $one1_2_1 = ($dictSemiring_0)->{'one'};
  $__res = (object)["add" => (function() use ($dictSemiring_0) {
  $__fn = function($f_3 = null, $g_4 = null, $x_5 = null) use ($dictSemiring_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictSemiring_0)->{'add'})(($f_3)($x_5)))(($g_4)($x_5));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "zero" => function($v_3 = null) use ($zero1_1_0) {
  $__num = \func_num_args();
  $__res = $zero1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "mul" => (function() use ($dictSemiring_0) {
  $__fn = function($f_3 = null, $g_4 = null, $x_5 = null) use ($dictSemiring_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictSemiring_0)->{'mul'})(($f_3)($x_5)))(($g_4)($x_5));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "one" => function($v_3 = null) use ($one1_2_1) {
  $__num = \func_num_args();
  $__res = $one1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Semiring_semiringRecordCons'] = function() { $v = (function() {
  $__fn = function($dictIsSymbol_0 = null, $dollar__unused_1 = null, $dictSemiringRecord_2 = null, $dictSemiring_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $one1_4_0 = ($dictSemiring_3)->{'one'};
  $zero1_5_1 = ($dictSemiring_3)->{'zero'};
  $__res = (object)["addRecord" => (function() use ($dictIsSymbol_0, $dictSemiringRecord_2, $dictSemiring_3) {
  $__fn = function($v_6 = null, $ra_7 = null, $rb_8 = null) use ($dictIsSymbol_0, $dictSemiringRecord_2, $dictSemiring_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $key_9_2 = (($dictIsSymbol_0)->{'reflectSymbol'})(new Phpurs_Data0("Proxy"));
  $get_10_3 = (($GLOBALS['Record_Unsafe_unsafeGet'] ?? \PhpursThunks::eval('Record_Unsafe_unsafeGet')))($key_9_2);
  $__res = (((($GLOBALS['Record_Unsafe_unsafeSet'] ?? \PhpursThunks::eval('Record_Unsafe_unsafeSet')))($key_9_2))(((($dictSemiring_3)->{'add'})(($get_10_3)($ra_7)))(($get_10_3)($rb_8))))((((($dictSemiringRecord_2)->{'addRecord'})(new Phpurs_Data0("Proxy")))($ra_7))($rb_8));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "mulRecord" => (function() use ($dictIsSymbol_0, $dictSemiringRecord_2, $dictSemiring_3) {
  $__fn = function($v_6 = null, $ra_7 = null, $rb_8 = null) use ($dictIsSymbol_0, $dictSemiringRecord_2, $dictSemiring_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $key_9_4 = (($dictIsSymbol_0)->{'reflectSymbol'})(new Phpurs_Data0("Proxy"));
  $get_10_5 = (($GLOBALS['Record_Unsafe_unsafeGet'] ?? \PhpursThunks::eval('Record_Unsafe_unsafeGet')))($key_9_4);
  $__res = (((($GLOBALS['Record_Unsafe_unsafeSet'] ?? \PhpursThunks::eval('Record_Unsafe_unsafeSet')))($key_9_4))(((($dictSemiring_3)->{'mul'})(($get_10_5)($ra_7)))(($get_10_5)($rb_8))))((((($dictSemiringRecord_2)->{'mulRecord'})(new Phpurs_Data0("Proxy")))($ra_7))($rb_8));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "oneRecord" => (function() use ($dictIsSymbol_0, $dictSemiringRecord_2, $one1_4_0) {
  $__fn = function($v_6 = null, $v1_7 = null) use ($dictIsSymbol_0, $dictSemiringRecord_2, $one1_4_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($GLOBALS['Record_Unsafe_unsafeSet'] ?? \PhpursThunks::eval('Record_Unsafe_unsafeSet')))((($dictIsSymbol_0)->{'reflectSymbol'})(new Phpurs_Data0("Proxy"))))($one1_4_0))(((($dictSemiringRecord_2)->{'oneRecord'})(new Phpurs_Data0("Proxy")))(new Phpurs_Data0("Proxy")));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "zeroRecord" => (function() use ($dictIsSymbol_0, $dictSemiringRecord_2, $zero1_5_1) {
  $__fn = function($v_6 = null, $v1_7 = null) use ($dictIsSymbol_0, $dictSemiringRecord_2, $zero1_5_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($GLOBALS['Record_Unsafe_unsafeSet'] ?? \PhpursThunks::eval('Record_Unsafe_unsafeSet')))((($dictIsSymbol_0)->{'reflectSymbol'})(new Phpurs_Data0("Proxy"))))($zero1_5_1))(((($dictSemiringRecord_2)->{'zeroRecord'})(new Phpurs_Data0("Proxy")))(new Phpurs_Data0("Proxy")));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(); return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };
$ffi_Data_Semiring = \call_user_func(function() {
  $exports = [];
$intAdd = function($a, $b = null) use (&$intAdd) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$intAdd) {

            return $intAdd(...\array_merge($__args, $more));
        };
    }
    return $a + $b;
};
$intMul = function($a, $b = null) use (&$intMul) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$intMul) {

            return $intMul(...\array_merge($__args, $more));
        };
    }
    return $a * $b;
};
$numAdd = $intAdd;
$numMul = $intMul;

$exports['intAdd'] = $intAdd;
$exports['intMul'] = $intMul;
$exports['numAdd'] = $numAdd;
$exports['numMul'] = $numMul;
return $exports;
  return $exports;
});
\PhpursThunks::$thunks['Data_Semiring_intAdd'] = function() use (&$ffi_Data_Semiring) { return $ffi_Data_Semiring['intAdd']; };
\PhpursThunks::$thunks['Data_Semiring_intMul'] = function() use (&$ffi_Data_Semiring) { return $ffi_Data_Semiring['intMul']; };
\PhpursThunks::$thunks['Data_Semiring_numAdd'] = function() use (&$ffi_Data_Semiring) { return $ffi_Data_Semiring['numAdd']; };
\PhpursThunks::$thunks['Data_Semiring_numMul'] = function() use (&$ffi_Data_Semiring) { return $ffi_Data_Semiring['numMul']; };


















