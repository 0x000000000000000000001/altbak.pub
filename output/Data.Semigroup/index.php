<?php

namespace Data\Semigroup;

// ALL IMPORTS: Data.Semigroup, Data.Symbol, Data.Unit, Data.Void, Prim, Prim.Row, Prim.RowList, Record.Unsafe, Type.Proxy
// TO REQUIRE: Data.Semigroup, Data.Symbol, Data.Unit, Data.Void, Record.Unsafe, Type.Proxy
require_once __DIR__ . '/../Data.Semigroup/index.php';
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
if (!\function_exists(__NAMESPACE__ . '\\phpurs_execute_effect')) {
  function phpurs_execute_effect($val) {
    if (\is_callable($val)) {
      return $val($GLOBALS['Data_Unit_unit']);
    }
    return $val;
  }
}

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };
$ffi_Data_Semigroup = \call_user_func(function() {
  $exports = [];
$concatString = function($x, $y) use (&$concatString) {
    return $x . $y;
};
$concatArray = function($x, $y) use (&$concatArray) {
    return \array_merge($x, $y);
};

$exports['concatString'] = $concatString;
$exports['concatArray'] = $concatArray;
return $exports;
  return $exports;
});
function majData_majSemigroup_concatmajArray($v0, $v1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majSemigroup_concatmajArray';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Data_Semigroup;
  $f = (\array_key_exists('concatArray', $ffi_Data_Semigroup) ? $ffi_Data_Semigroup['concatArray'] : new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Data_Semigroup_concatArray'] = __NAMESPACE__ . '\\majData_majSemigroup_concatmajArray';

function majData_majSemigroup_concatmajString(string $v0, $v1 = null): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majSemigroup_concatmajString';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Data_Semigroup;
  $f = (\array_key_exists('concatString', $ffi_Data_Semigroup) ? $ffi_Data_Semigroup['concatString'] : new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Data_Semigroup_concatString'] = __NAMESPACE__ . '\\majData_majSemigroup_concatmajString';





// Data_Semigroup_semigroupVoid
$GLOBALS['Data_Semigroup_semigroupVoid'] = (object)["append" => function($v_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Void_absurd'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Semigroup_semigroupUnit
$GLOBALS['Data_Semigroup_semigroupUnit'] = (object)["append" => function($v_0) {
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
}];

// Data_Semigroup_semigroupString
$GLOBALS['Data_Semigroup_semigroupString'] = (object)["append" => $GLOBALS['Data_Semigroup_concatString']];

// Data_Semigroup_semigroupRecordNil
$GLOBALS['Data_Semigroup_semigroupRecordNil'] = (object)["appendRecord" => function($v_0) {
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
}];

// Data_Semigroup_semigroupProxy
$GLOBALS['Data_Semigroup_semigroupProxy'] = (object)["append" => function($v_0) {
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
}];

// Data_Semigroup_semigroupArray
$GLOBALS['Data_Semigroup_semigroupArray'] = (object)["append" => $GLOBALS['Data_Semigroup_concatArray']];

// Data_Semigroup_appendRecord
function majData_majSemigroup_appendmajRecord($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSemigroup_appendmajRecord';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'appendRecord'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Semigroup_appendRecord'] = __NAMESPACE__ . '\\majData_majSemigroup_appendmajRecord';

// Data_Semigroup_semigroupRecord
function majData_majSemigroup_semigroupmajRecord($_dollar___unused_0, $dictSemigroupRecord_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSemigroup_semigroupmajRecord';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (object)["append" => (($dictSemigroupRecord_1)->{'appendRecord'})(new \Type\Proxy\Type_Proxy_Proxy())];
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Semigroup_semigroupRecord'] = __NAMESPACE__ . '\\majData_majSemigroup_semigroupmajRecord';

// Data_Semigroup_append
function majData_majSemigroup_append($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSemigroup_append';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'append'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Semigroup_append'] = __NAMESPACE__ . '\\majData_majSemigroup_append';

// Data_Semigroup_semigroupFn
function majData_majSemigroup_semigroupmajFn($dictSemigroup_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSemigroup_semigroupmajFn';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["append" => function($f_1) use ($dictSemigroup_0) {
  $__num = \func_num_args();
  $__res = function($g_2) use ($dictSemigroup_0, $f_1) {
  $__num = \func_num_args();
  $__res = function($x_3) use ($dictSemigroup_0, $f_1, $g_2) {
  $__num = \func_num_args();
  $__res = ((($dictSemigroup_0)->{'append'})(($f_1)($x_3)))(($g_2)($x_3));
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
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Semigroup_semigroupFn'] = __NAMESPACE__ . '\\majData_majSemigroup_semigroupmajFn';

// Data_Semigroup_semigroupRecordCons
function majData_majSemigroup_semigroupmajRecordmajCons($dictIsSymbol_0, $_dollar___unused_1 = null, $dictSemigroupRecord_2 = null, $dictSemigroup_3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSemigroup_semigroupmajRecordmajCons';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = (object)["appendRecord" => function($v_4) use ($dictIsSymbol_0, $dictSemigroupRecord_2, $dictSemigroup_3) {
  $__num = \func_num_args();
  $__res = function($ra_5) use ($dictIsSymbol_0, $dictSemigroupRecord_2, $dictSemigroup_3) {
  $__num = \func_num_args();
  $__res = function($rb_6) use ($dictIsSymbol_0, $dictSemigroupRecord_2, $dictSemigroup_3, $ra_5) {
  $__num = \func_num_args();
  $key_7_0 = (($dictIsSymbol_0)->{'reflectSymbol'})(new \Type\Proxy\Type_Proxy_Proxy());
  $get_8_1 = ($GLOBALS['Record_Unsafe_unsafeGet'])($key_7_0);
  $__res = \Record\Unsafe\majRecord_majUnsafe_unsafemajSet($key_7_0, ((($dictSemigroup_3)->{'append'})(($get_8_1)($ra_5)))(($get_8_1)($rb_6)), (((($dictSemigroupRecord_2)->{'appendRecord'})(new \Type\Proxy\Type_Proxy_Proxy()))($ra_5))($rb_6));
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
}];
  goto __end;;
  __end:
  return 4 < $__num ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
}
$GLOBALS['Data_Semigroup_semigroupRecordCons'] = __NAMESPACE__ . '\\majData_majSemigroup_semigroupmajRecordmajCons';

