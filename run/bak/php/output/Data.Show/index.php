<?php

namespace Data\Show;

// ALL IMPORTS: Data.Semigroup, Data.Show, Data.Symbol, Data.Unit, Data.Void, Prim, Prim.Row, Prim.RowList, Record.Unsafe, Type.Proxy
// TO REQUIRE: Data.Semigroup, Data.Show, Data.Symbol, Data.Unit, Data.Void, Record.Unsafe, Type.Proxy
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
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

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };
$ffi_Data_Show = \call_user_func(function() {
  $exports = [];
$showIntImpl = function($i) use (&$showIntImpl) { return (string)$i; };
$showStringImpl = function($s) use (&$showStringImpl) { return json_encode($s); };
$showNumberImpl = function($n) use (&$showNumberImpl) {
    $str = (string)$n;
    if (strpos($str, '.') === false && strpos($str, 'e') === false && strpos($str, 'E') === false && !is_nan($n)) {
        return $str . '.0';
    }
    return $str;
};
$showCharImpl = function($c) use (&$showCharImpl) {
    return "'" . $c . "'";
};
$showArrayImpl = function($f, $xs = null) use (&$showArrayImpl) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$showArrayImpl) {

            return $showArrayImpl(...\array_merge($__args, $more));
        };
    }
    return "[" . implode(",", array_map($f, $xs)) . "]";
};

$exports['showIntImpl'] = $showIntImpl;
$exports['showStringImpl'] = $showStringImpl;
$exports['showNumberImpl'] = $showNumberImpl;
$exports['showCharImpl'] = $showCharImpl;
$exports['showArrayImpl'] = $showArrayImpl;
return $exports;
  return $exports;
});
function majData_majShow_showmajArraymajImpl($v0, $v1 = null): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majShow_showmajArraymajImpl';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Data_Show;
  $f = ($ffi_Data_Show['showArrayImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Data_Show_showArrayImpl'] = __NAMESPACE__ . '\\majData_majShow_showmajArraymajImpl';

function majData_majShow_showmajCharmajImpl($v0): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majShow_showmajCharmajImpl';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Data_Show;
  $f = ($ffi_Data_Show['showCharImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Data_Show_showCharImpl'] = __NAMESPACE__ . '\\majData_majShow_showmajCharmajImpl';

function majData_majShow_showmajIntmajImpl(int $v0): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majShow_showmajIntmajImpl';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Data_Show;
  $f = ($ffi_Data_Show['showIntImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Data_Show_showIntImpl'] = __NAMESPACE__ . '\\majData_majShow_showmajIntmajImpl';

function majData_majShow_showmajNumbermajImpl(float $v0): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majShow_showmajNumbermajImpl';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Data_Show;
  $f = ($ffi_Data_Show['showNumberImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Data_Show_showNumberImpl'] = __NAMESPACE__ . '\\majData_majShow_showmajNumbermajImpl';

function majData_majShow_showmajStringmajImpl(string $v0): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majShow_showmajStringmajImpl';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Data_Show;
  $f = ($ffi_Data_Show['showStringImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Data_Show_showStringImpl'] = __NAMESPACE__ . '\\majData_majShow_showmajStringmajImpl';





// Data_Show_showVoid
$GLOBALS['Data_Show_showVoid'] = (object)["show" => $GLOBALS['Data_Void_absurd']];

// Data_Show_showUnit
$GLOBALS['Data_Show_showUnit'] = (object)["show" => function($v_0) {
  $__num = \func_num_args();
  $__res = "unit";
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Show_showString
$GLOBALS['Data_Show_showString'] = (object)["show" => $GLOBALS['Data_Show_showStringImpl']];

// Data_Show_showRecordFieldsNil
$GLOBALS['Data_Show_showRecordFieldsNil'] = (object)["showRecordFields" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = "";
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];

// Data_Show_showRecordFields
function majData_majShow_showmajRecordmajFields($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majShow_showmajRecordmajFields';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'showRecordFields'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Show_showRecordFields'] = __NAMESPACE__ . '\\majData_majShow_showmajRecordmajFields';

// Data_Show_showRecord
function majData_majShow_showmajRecord($_dollar__unused_0, $_dollar__unused_1 = null, $dictShowRecordFields_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majShow_showmajRecord';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (object)["show" => function($record_3) use ($dictShowRecordFields_2) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})("{"))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})(((($dictShowRecordFields_2)->{'showRecordFields'})(new \Type\Proxy\Type_Proxy_Proxy()))($record_3)))("}"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Show_showRecord'] = __NAMESPACE__ . '\\majData_majShow_showmajRecord';

// Data_Show_showProxy
$GLOBALS['Data_Show_showProxy'] = (object)["show" => function($v_0) {
  $__num = \func_num_args();
  $__res = "Proxy";
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Show_showNumber
$GLOBALS['Data_Show_showNumber'] = (object)["show" => $GLOBALS['Data_Show_showNumberImpl']];

// Data_Show_showInt
$GLOBALS['Data_Show_showInt'] = (object)["show" => $GLOBALS['Data_Show_showIntImpl']];

// Data_Show_showChar
$GLOBALS['Data_Show_showChar'] = (object)["show" => $GLOBALS['Data_Show_showCharImpl']];

// Data_Show_showBoolean
$GLOBALS['Data_Show_showBoolean'] = (object)["show" => function($v_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v_0) {
$__t0 = "true";
goto end_branch_0;;
};
  $__t0 = "false";
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Show_show
function majData_majShow_show($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majShow_show';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'show'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Show_show'] = __NAMESPACE__ . '\\majData_majShow_show';

// Data_Show_showArray
function majData_majShow_showmajArray($dictShow_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majShow_showmajArray';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["show" => ($GLOBALS['Data_Show_showArrayImpl'])(($dictShow_0)->{'show'})];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Show_showArray'] = __NAMESPACE__ . '\\majData_majShow_showmajArray';

// Data_Show_showRecordFieldsCons
function majData_majShow_showmajRecordmajFieldsmajCons($dictIsSymbol_0, $dictShowRecordFields_1 = null, $dictShow_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majShow_showmajRecordmajFieldsmajCons';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (object)["showRecordFields" => (function() use ($dictIsSymbol_0, $dictShowRecordFields_1, $dictShow_2) {
  $__fn = function($v_3, $record_4 = null) use ($dictIsSymbol_0, $dictShowRecordFields_1, $dictShow_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $key_5_0 = (($dictIsSymbol_0)->{'reflectSymbol'})(new \Type\Proxy\Type_Proxy_Proxy());
  $__res = ((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})(" "))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})($key_5_0))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})(": "))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})((($dictShow_2)->{'show'})(\Record\Unsafe\majRecord_majUnsafe_unsafemajGet($key_5_0, $record_4))))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})(","))(((($dictShowRecordFields_1)->{'showRecordFields'})(new \Type\Proxy\Type_Proxy_Proxy()))($record_4))))));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Show_showRecordFieldsCons'] = __NAMESPACE__ . '\\majData_majShow_showmajRecordmajFieldsmajCons';

// Data_Show_showRecordFieldsConsNil
function majData_majShow_showmajRecordmajFieldsmajConsmajNil($dictIsSymbol_0, $dictShow_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majShow_showmajRecordmajFieldsmajConsmajNil';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (object)["showRecordFields" => (function() use ($dictIsSymbol_0, $dictShow_1) {
  $__fn = function($v_2, $record_3 = null) use ($dictIsSymbol_0, $dictShow_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $key_4_0 = (($dictIsSymbol_0)->{'reflectSymbol'})(new \Type\Proxy\Type_Proxy_Proxy());
  $__res = ((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})(" "))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})($key_4_0))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})(": "))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})((($dictShow_1)->{'show'})(\Record\Unsafe\majRecord_majUnsafe_unsafemajGet($key_4_0, $record_3))))(" "))));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Show_showRecordFieldsConsNil'] = __NAMESPACE__ . '\\majData_majShow_showmajRecordmajFieldsmajConsmajNil';

