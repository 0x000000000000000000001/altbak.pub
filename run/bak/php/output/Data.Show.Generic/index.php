<?php

namespace Data\Show\Generic;

// ALL IMPORTS: Data.Generic.Rep, Data.Semigroup, Data.Show, Data.Show.Generic, Data.Symbol, Prelude, Prim, Type.Proxy
// TO REQUIRE: Data.Generic.Rep, Data.Semigroup, Data.Show, Data.Show.Generic, Data.Symbol, Prelude, Type.Proxy
require_once __DIR__ . '/../Data.Generic.Rep/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
require_once __DIR__ . '/../Data.Show.Generic/index.php';
require_once __DIR__ . '/../Data.Symbol/index.php';
require_once __DIR__ . '/../Prelude/index.php';
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
$ffi_Data_Show_Generic = \call_user_func(function() {
  $exports = [];
$intercalate = function ($separator, $xs = null) use (&$intercalate) {
    if (func_num_args() < 2) {
        $__args = func_get_args();
        return function (...$more) use ($__args, &$intercalate) {
            return $intercalate(...array_merge($__args, $more));
        };
    }
    return implode($separator, $xs);
};

$exports['intercalate'] = $intercalate;
return $exports;
  return $exports;
});
function majData_majShow_majGeneric_intercalate(string $v0, $v1 = null): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majShow_majGeneric_intercalate';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Data_Show_Generic;
  $f = ($ffi_Data_Show_Generic['intercalate'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Data_Show_Generic_intercalate'] = __NAMESPACE__ . '\\majData_majShow_majGeneric_intercalate';





// Data_Show_Generic_genericShowArgsNoArguments
$GLOBALS['Data_Show_Generic_genericShowArgsNoArguments'] = (object)["genericShowArgs" => function($v_0) {
  $__num = \func_num_args();
  $__res = [];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Show_Generic_genericShowArgsArgument
function majData_majShow_majGeneric_genericmajShowmajArgsmajArgument($dictShow_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majShow_majGeneric_genericmajShowmajArgsmajArgument';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["genericShowArgs" => function($v_1) use ($dictShow_0) {
  $__num = \func_num_args();
  $__res = [(($dictShow_0)->{'show'})($v_1)];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Show_Generic_genericShowArgsArgument'] = __NAMESPACE__ . '\\majData_majShow_majGeneric_genericmajShowmajArgsmajArgument';

// Data_Show_Generic_genericShowArgs
function majData_majShow_majGeneric_genericmajShowmajArgs($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majShow_majGeneric_genericmajShowmajArgs';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'genericShowArgs'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Show_Generic_genericShowArgs'] = __NAMESPACE__ . '\\majData_majShow_majGeneric_genericmajShowmajArgs';

// Data_Show_Generic_genericShowArgsProduct
function majData_majShow_majGeneric_genericmajShowmajArgsmajProduct($dictGenericShowArgs_0, $dictGenericShowArgs1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majShow_majGeneric_genericmajShowmajArgsmajProduct';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (object)["genericShowArgs" => function($v_2) use ($dictGenericShowArgs1_1, $dictGenericShowArgs_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Semigroup_semigroupArray'])->{'append'})((($dictGenericShowArgs_0)->{'genericShowArgs'})(($v_2)->{'value0'})))((($dictGenericShowArgs1_1)->{'genericShowArgs'})(($v_2)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Show_Generic_genericShowArgsProduct'] = __NAMESPACE__ . '\\majData_majShow_majGeneric_genericmajShowmajArgsmajProduct';

// Data_Show_Generic_genericShowConstructor
function majData_majShow_majGeneric_genericmajShowmajConstructor($dictGenericShowArgs_0, $dictIsSymbol_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majShow_majGeneric_genericmajShowmajConstructor';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (object)["genericShow__prime__" => function($v_2) use ($dictGenericShowArgs_0, $dictIsSymbol_1) {
  $__num = \func_num_args();
  $ctor_3_0 = (($dictIsSymbol_1)->{'reflectSymbol'})(new \Type\Proxy\Type_Proxy_Proxy());
  $v1_4_1 = (($dictGenericShowArgs_0)->{'genericShowArgs'})($v_2);
  $__res = match (count($v1_4_1)) { 0 => $ctor_3_0, default => ((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})("("))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})(\Data\Show\Generic\majData_majShow_majGeneric_intercalate(" ", ((($GLOBALS['Data_Semigroup_semigroupArray'])->{'append'})([$ctor_3_0]))($v1_4_1))))(")")) };
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Show_Generic_genericShowConstructor'] = __NAMESPACE__ . '\\majData_majShow_majGeneric_genericmajShowmajConstructor';

// Data_Show_Generic_genericShow'
function majData_majShow_majGeneric_genericmajShow__prime__($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majShow_majGeneric_genericmajShow__prime__';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'genericShow__prime__'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Show_Generic_genericShow__prime__'] = __NAMESPACE__ . '\\majData_majShow_majGeneric_genericmajShow__prime__';

// Data_Show_Generic_genericShowNoConstructors
$GLOBALS['Data_Show_Generic_genericShowNoConstructors'] = (object)["genericShow__prime__" => function($a_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Show_Generic_genericShowNoConstructors'])->{'genericShow__prime__'})($a_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Show_Generic_genericShowSum
function majData_majShow_majGeneric_genericmajShowmajSum($dictGenericShow_0, $dictGenericShow1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majShow_majGeneric_genericmajShowmajSum';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (object)["genericShow__prime__" => function($v_2) use ($dictGenericShow1_1, $dictGenericShow_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v_2 instanceof \Data\Generic\Rep\Data_Generic_Rep_Inl) {
$__t0 = (($dictGenericShow_0)->{'genericShow__prime__'})(($v_2)->{'value0'});
goto end_branch_0;;
};
  if ($v_2 instanceof \Data\Generic\Rep\Data_Generic_Rep_Inr) {
$__t0 = (($dictGenericShow1_1)->{'genericShow__prime__'})(($v_2)->{'value0'});
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Show_Generic_genericShowSum'] = __NAMESPACE__ . '\\majData_majShow_majGeneric_genericmajShowmajSum';

// Data_Show_Generic_genericShow
function majData_majShow_majGeneric_genericmajShow($dictGeneric_0, $dictGenericShow_1 = null, $x_2 = null): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majShow_majGeneric_genericmajShow';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($dictGenericShow_1)->{'genericShow__prime__'})((($dictGeneric_0)->{'from'})($x_2));
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Show_Generic_genericShow'] = __NAMESPACE__ . '\\majData_majShow_majGeneric_genericmajShow';

