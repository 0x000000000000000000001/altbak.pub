<?php

namespace Effect\Uncurried;

// ALL IMPORTS: Data.Monoid, Data.Semigroup, Effect, Effect.Uncurried, Prim
// TO REQUIRE: Data.Monoid, Data.Semigroup, Effect, Effect.Uncurried
require_once __DIR__ . '/../Data.Monoid/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Effect/index.php';
require_once __DIR__ . '/../Effect.Uncurried/index.php';

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
$ffi_Effect_Uncurried = \call_user_func(function() {
  $exports = [];
$exports = [];

for ($i = 1; $i <= 10; $i++) {
    $exports["mkEffectFn$i"] = function($fn) use ($i) {
        return function(...$args) use ($fn) {
            $curr = $fn;
            foreach ($args as $arg) {
                $curr = $curr($arg);
            }
            return $curr();
        };
    };

    $exports["runEffectFn$i"] = function(...$args) use ($i) {
        $expectedArgs = $i + 1; // fn + $i arguments
        
        $curry = function($collectedArgs) use (&$curry, $expectedArgs) {
            if (\count($collectedArgs) >= $expectedArgs) {
                $fn = $collectedArgs[0];
                $actualArgs = \array_slice($collectedArgs, 1, $expectedArgs - 1);
                $res = function() use ($fn, $actualArgs) {
                    return $fn(...$actualArgs);
                };
                if (\count($collectedArgs) > $expectedArgs) {
                    $extra = \array_slice($collectedArgs, $expectedArgs);
                    return $res(...$extra);
                }
                return $res;
            }
            return function(...$more) use (&$curry, $collectedArgs) {
                return $curry(\array_merge($collectedArgs, $more));
            };
        };
        
        return $curry($args);
    };
}

return $exports;
  return $exports;
});
$GLOBALS['Effect_Uncurried_mkEffectFn1'] = $ffi_Effect_Uncurried['mkEffectFn1'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Effect_Uncurried_mkEffectFn10'] = $ffi_Effect_Uncurried['mkEffectFn10'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Effect_Uncurried_mkEffectFn2'] = $ffi_Effect_Uncurried['mkEffectFn2'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Effect_Uncurried_mkEffectFn3'] = $ffi_Effect_Uncurried['mkEffectFn3'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Effect_Uncurried_mkEffectFn4'] = $ffi_Effect_Uncurried['mkEffectFn4'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Effect_Uncurried_mkEffectFn5'] = $ffi_Effect_Uncurried['mkEffectFn5'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Effect_Uncurried_mkEffectFn6'] = $ffi_Effect_Uncurried['mkEffectFn6'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Effect_Uncurried_mkEffectFn7'] = $ffi_Effect_Uncurried['mkEffectFn7'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Effect_Uncurried_mkEffectFn8'] = $ffi_Effect_Uncurried['mkEffectFn8'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Effect_Uncurried_mkEffectFn9'] = $ffi_Effect_Uncurried['mkEffectFn9'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Effect_Uncurried_runEffectFn1'] = $ffi_Effect_Uncurried['runEffectFn1'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Effect_Uncurried_runEffectFn10'] = $ffi_Effect_Uncurried['runEffectFn10'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Effect_Uncurried_runEffectFn2'] = $ffi_Effect_Uncurried['runEffectFn2'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Effect_Uncurried_runEffectFn3'] = $ffi_Effect_Uncurried['runEffectFn3'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Effect_Uncurried_runEffectFn4'] = $ffi_Effect_Uncurried['runEffectFn4'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Effect_Uncurried_runEffectFn5'] = $ffi_Effect_Uncurried['runEffectFn5'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Effect_Uncurried_runEffectFn6'] = $ffi_Effect_Uncurried['runEffectFn6'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Effect_Uncurried_runEffectFn7'] = $ffi_Effect_Uncurried['runEffectFn7'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Effect_Uncurried_runEffectFn8'] = $ffi_Effect_Uncurried['runEffectFn8'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Effect_Uncurried_runEffectFn9'] = $ffi_Effect_Uncurried['runEffectFn9'] ?? new class { public function __invoke(...$args) { return $this; } };




// Effect_Uncurried_semigroupEffectFn9
function majEffect_majUncurried_semigroupmajEffectmajFn9($dictSemigroup_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majUncurried_semigroupmajEffectmajFn9';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ["append" => (function() use ($dictSemigroup_0) {
  $__fn = function($f1_1 = null, $f2_2 = null) use ($dictSemigroup_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Effect_Uncurried_mkEffectFn9'])((function() use ($dictSemigroup_0, $f1_1, $f2_2) {
  $__fn = function($a_3 = null, $b_4 = null, $c_5 = null, $d_6 = null, $e_7 = null, $f_8 = null, $g_9 = null, $h_10 = null, $i_11 = null) use ($dictSemigroup_0, $f1_1, $f2_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 9) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 9);
  }
  $__res = ($GLOBALS['Effect_lift2'])(($dictSemigroup_0)['append'], ($GLOBALS['Effect_Uncurried_runEffectFn9'])($f1_1, $a_3, $b_4, $c_5, $d_6, $e_7, $f_8, $g_9, $h_10, $i_11), ($GLOBALS['Effect_Uncurried_runEffectFn9'])($f2_2, $a_3, $b_4, $c_5, $d_6, $e_7, $f_8, $g_9, $h_10, $i_11));
  goto __end;;
  __end:
  return $__num > 9 ? $__res(...\array_slice(\func_get_args(), 9)) : $__res;
  };
  return $__fn;
})());
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Effect_Uncurried_semigroupEffectFn9'] = __NAMESPACE__ . '\\majEffect_majUncurried_semigroupmajEffectmajFn9';

// Effect_Uncurried_semigroupEffectFn8
function majEffect_majUncurried_semigroupmajEffectmajFn8($dictSemigroup_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majUncurried_semigroupmajEffectmajFn8';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ["append" => (function() use ($dictSemigroup_0) {
  $__fn = function($f1_1 = null, $f2_2 = null) use ($dictSemigroup_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Effect_Uncurried_mkEffectFn8'])((function() use ($dictSemigroup_0, $f1_1, $f2_2) {
  $__fn = function($a_3 = null, $b_4 = null, $c_5 = null, $d_6 = null, $e_7 = null, $f_8 = null, $g_9 = null, $h_10 = null) use ($dictSemigroup_0, $f1_1, $f2_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 8) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 8);
  }
  $__res = ($GLOBALS['Effect_lift2'])(($dictSemigroup_0)['append'], ($GLOBALS['Effect_Uncurried_runEffectFn8'])($f1_1, $a_3, $b_4, $c_5, $d_6, $e_7, $f_8, $g_9, $h_10), ($GLOBALS['Effect_Uncurried_runEffectFn8'])($f2_2, $a_3, $b_4, $c_5, $d_6, $e_7, $f_8, $g_9, $h_10));
  goto __end;;
  __end:
  return $__num > 8 ? $__res(...\array_slice(\func_get_args(), 8)) : $__res;
  };
  return $__fn;
})());
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Effect_Uncurried_semigroupEffectFn8'] = __NAMESPACE__ . '\\majEffect_majUncurried_semigroupmajEffectmajFn8';

// Effect_Uncurried_semigroupEffectFn7
function majEffect_majUncurried_semigroupmajEffectmajFn7($dictSemigroup_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majUncurried_semigroupmajEffectmajFn7';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ["append" => (function() use ($dictSemigroup_0) {
  $__fn = function($f1_1 = null, $f2_2 = null) use ($dictSemigroup_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Effect_Uncurried_mkEffectFn7'])((function() use ($dictSemigroup_0, $f1_1, $f2_2) {
  $__fn = function($a_3 = null, $b_4 = null, $c_5 = null, $d_6 = null, $e_7 = null, $f_8 = null, $g_9 = null) use ($dictSemigroup_0, $f1_1, $f2_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 7) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 7);
  }
  $__res = ($GLOBALS['Effect_lift2'])(($dictSemigroup_0)['append'], ($GLOBALS['Effect_Uncurried_runEffectFn7'])($f1_1, $a_3, $b_4, $c_5, $d_6, $e_7, $f_8, $g_9), ($GLOBALS['Effect_Uncurried_runEffectFn7'])($f2_2, $a_3, $b_4, $c_5, $d_6, $e_7, $f_8, $g_9));
  goto __end;;
  __end:
  return $__num > 7 ? $__res(...\array_slice(\func_get_args(), 7)) : $__res;
  };
  return $__fn;
})());
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Effect_Uncurried_semigroupEffectFn7'] = __NAMESPACE__ . '\\majEffect_majUncurried_semigroupmajEffectmajFn7';

// Effect_Uncurried_semigroupEffectFn6
function majEffect_majUncurried_semigroupmajEffectmajFn6($dictSemigroup_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majUncurried_semigroupmajEffectmajFn6';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ["append" => (function() use ($dictSemigroup_0) {
  $__fn = function($f1_1 = null, $f2_2 = null) use ($dictSemigroup_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Effect_Uncurried_mkEffectFn6'])((function() use ($dictSemigroup_0, $f1_1, $f2_2) {
  $__fn = function($a_3 = null, $b_4 = null, $c_5 = null, $d_6 = null, $e_7 = null, $f_8 = null) use ($dictSemigroup_0, $f1_1, $f2_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 6) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 6);
  }
  $__res = ($GLOBALS['Effect_lift2'])(($dictSemigroup_0)['append'], ($GLOBALS['Effect_Uncurried_runEffectFn6'])($f1_1, $a_3, $b_4, $c_5, $d_6, $e_7, $f_8), ($GLOBALS['Effect_Uncurried_runEffectFn6'])($f2_2, $a_3, $b_4, $c_5, $d_6, $e_7, $f_8));
  goto __end;;
  __end:
  return $__num > 6 ? $__res(...\array_slice(\func_get_args(), 6)) : $__res;
  };
  return $__fn;
})());
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Effect_Uncurried_semigroupEffectFn6'] = __NAMESPACE__ . '\\majEffect_majUncurried_semigroupmajEffectmajFn6';

// Effect_Uncurried_semigroupEffectFn5
function majEffect_majUncurried_semigroupmajEffectmajFn5($dictSemigroup_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majUncurried_semigroupmajEffectmajFn5';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ["append" => (function() use ($dictSemigroup_0) {
  $__fn = function($f1_1 = null, $f2_2 = null) use ($dictSemigroup_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Effect_Uncurried_mkEffectFn5'])((function() use ($dictSemigroup_0, $f1_1, $f2_2) {
  $__fn = function($a_3 = null, $b_4 = null, $c_5 = null, $d_6 = null, $e_7 = null) use ($dictSemigroup_0, $f1_1, $f2_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 5) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 5);
  }
  $__res = ($GLOBALS['Effect_lift2'])(($dictSemigroup_0)['append'], ($GLOBALS['Effect_Uncurried_runEffectFn5'])($f1_1, $a_3, $b_4, $c_5, $d_6, $e_7), ($GLOBALS['Effect_Uncurried_runEffectFn5'])($f2_2, $a_3, $b_4, $c_5, $d_6, $e_7));
  goto __end;;
  __end:
  return $__num > 5 ? $__res(...\array_slice(\func_get_args(), 5)) : $__res;
  };
  return $__fn;
})());
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Effect_Uncurried_semigroupEffectFn5'] = __NAMESPACE__ . '\\majEffect_majUncurried_semigroupmajEffectmajFn5';

// Effect_Uncurried_semigroupEffectFn4
function majEffect_majUncurried_semigroupmajEffectmajFn4($dictSemigroup_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majUncurried_semigroupmajEffectmajFn4';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ["append" => (function() use ($dictSemigroup_0) {
  $__fn = function($f1_1 = null, $f2_2 = null) use ($dictSemigroup_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Effect_Uncurried_mkEffectFn4'])((function() use ($dictSemigroup_0, $f1_1, $f2_2) {
  $__fn = function($a_3 = null, $b_4 = null, $c_5 = null, $d_6 = null) use ($dictSemigroup_0, $f1_1, $f2_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = ($GLOBALS['Effect_lift2'])(($dictSemigroup_0)['append'], ($GLOBALS['Effect_Uncurried_runEffectFn4'])($f1_1, $a_3, $b_4, $c_5, $d_6), ($GLOBALS['Effect_Uncurried_runEffectFn4'])($f2_2, $a_3, $b_4, $c_5, $d_6));
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})());
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Effect_Uncurried_semigroupEffectFn4'] = __NAMESPACE__ . '\\majEffect_majUncurried_semigroupmajEffectmajFn4';

// Effect_Uncurried_semigroupEffectFn3
function majEffect_majUncurried_semigroupmajEffectmajFn3($dictSemigroup_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majUncurried_semigroupmajEffectmajFn3';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ["append" => (function() use ($dictSemigroup_0) {
  $__fn = function($f1_1 = null, $f2_2 = null) use ($dictSemigroup_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Effect_Uncurried_mkEffectFn3'])((function() use ($dictSemigroup_0, $f1_1, $f2_2) {
  $__fn = function($a_3 = null, $b_4 = null, $c_5 = null) use ($dictSemigroup_0, $f1_1, $f2_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($GLOBALS['Effect_lift2'])(($dictSemigroup_0)['append'], ($GLOBALS['Effect_Uncurried_runEffectFn3'])($f1_1, $a_3, $b_4, $c_5), ($GLOBALS['Effect_Uncurried_runEffectFn3'])($f2_2, $a_3, $b_4, $c_5));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})());
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Effect_Uncurried_semigroupEffectFn3'] = __NAMESPACE__ . '\\majEffect_majUncurried_semigroupmajEffectmajFn3';

// Effect_Uncurried_semigroupEffectFn2
function majEffect_majUncurried_semigroupmajEffectmajFn2($dictSemigroup_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majUncurried_semigroupmajEffectmajFn2';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ["append" => (function() use ($dictSemigroup_0) {
  $__fn = function($f1_1 = null, $f2_2 = null) use ($dictSemigroup_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Effect_Uncurried_mkEffectFn2'])((function() use ($dictSemigroup_0, $f1_1, $f2_2) {
  $__fn = function($a_3 = null, $b_4 = null) use ($dictSemigroup_0, $f1_1, $f2_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Effect_lift2'])(($dictSemigroup_0)['append'], ($GLOBALS['Effect_Uncurried_runEffectFn2'])($f1_1, $a_3, $b_4), ($GLOBALS['Effect_Uncurried_runEffectFn2'])($f2_2, $a_3, $b_4));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})());
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Effect_Uncurried_semigroupEffectFn2'] = __NAMESPACE__ . '\\majEffect_majUncurried_semigroupmajEffectmajFn2';

// Effect_Uncurried_semigroupEffectFn10
function majEffect_majUncurried_semigroupmajEffectmajFn10($dictSemigroup_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majUncurried_semigroupmajEffectmajFn10';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ["append" => (function() use ($dictSemigroup_0) {
  $__fn = function($f1_1 = null, $f2_2 = null) use ($dictSemigroup_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Effect_Uncurried_mkEffectFn10'])((function() use ($dictSemigroup_0, $f1_1, $f2_2) {
  $__fn = function($a_3 = null, $b_4 = null, $c_5 = null, $d_6 = null, $e_7 = null, $f_8 = null, $g_9 = null, $h_10 = null, $i_11 = null, $j_12 = null) use ($dictSemigroup_0, $f1_1, $f2_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 10) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 10);
  }
  $__res = ($GLOBALS['Effect_lift2'])(($dictSemigroup_0)['append'], ($GLOBALS['Effect_Uncurried_runEffectFn10'])($f1_1, $a_3, $b_4, $c_5, $d_6, $e_7, $f_8, $g_9, $h_10, $i_11, $j_12), ($GLOBALS['Effect_Uncurried_runEffectFn10'])($f2_2, $a_3, $b_4, $c_5, $d_6, $e_7, $f_8, $g_9, $h_10, $i_11, $j_12));
  goto __end;;
  __end:
  return $__num > 10 ? $__res(...\array_slice(\func_get_args(), 10)) : $__res;
  };
  return $__fn;
})());
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Effect_Uncurried_semigroupEffectFn10'] = __NAMESPACE__ . '\\majEffect_majUncurried_semigroupmajEffectmajFn10';

// Effect_Uncurried_semigroupEffectFn1
function majEffect_majUncurried_semigroupmajEffectmajFn1($dictSemigroup_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majUncurried_semigroupmajEffectmajFn1';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ["append" => (function() use ($dictSemigroup_0) {
  $__fn = function($f1_1 = null, $f2_2 = null) use ($dictSemigroup_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Effect_Uncurried_mkEffectFn1'])(function($a_3 = null) use ($dictSemigroup_0, $f1_1, $f2_2) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Effect_lift2'])(($dictSemigroup_0)['append'], ($GLOBALS['Effect_Uncurried_runEffectFn1'])($f1_1, $a_3), ($GLOBALS['Effect_Uncurried_runEffectFn1'])($f2_2, $a_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Effect_Uncurried_semigroupEffectFn1'] = __NAMESPACE__ . '\\majEffect_majUncurried_semigroupmajEffectmajFn1';

// Effect_Uncurried_monoidEffectFn9
function majEffect_majUncurried_monoidmajEffectmajFn9($dictMonoid_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majUncurried_monoidmajEffectmajFn9';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $mempty_1_0 = (($GLOBALS['Effect_monoidEffect'])($dictMonoid_0))['mempty'];
  $semigroupEffectFn91_2_1 = ($GLOBALS['Effect_Uncurried_semigroupEffectFn9'])((($dictMonoid_0)['Semigroup0'])(null));
  $__res = ["mempty" => ($GLOBALS['Effect_Uncurried_mkEffectFn9'])((function() use ($mempty_1_0) {
  $__fn = function($v_3 = null, $v1_4 = null, $v2_5 = null, $v3_6 = null, $v4_7 = null, $v5_8 = null, $v6_9 = null, $v7_10 = null, $v8_11 = null) use ($mempty_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 9) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 9);
  }
  $__res = $mempty_1_0;
  goto __end;;
  __end:
  return $__num > 9 ? $__res(...\array_slice(\func_get_args(), 9)) : $__res;
  };
  return $__fn;
})()), "Semigroup0" => function($_dollar__unused_3 = null) use ($semigroupEffectFn91_2_1) {
  $__num = \func_num_args();
  $__res = $semigroupEffectFn91_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Effect_Uncurried_monoidEffectFn9'] = __NAMESPACE__ . '\\majEffect_majUncurried_monoidmajEffectmajFn9';

// Effect_Uncurried_monoidEffectFn8
function majEffect_majUncurried_monoidmajEffectmajFn8($dictMonoid_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majUncurried_monoidmajEffectmajFn8';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $mempty_1_0 = (($GLOBALS['Effect_monoidEffect'])($dictMonoid_0))['mempty'];
  $semigroupEffectFn81_2_1 = ($GLOBALS['Effect_Uncurried_semigroupEffectFn8'])((($dictMonoid_0)['Semigroup0'])(null));
  $__res = ["mempty" => ($GLOBALS['Effect_Uncurried_mkEffectFn8'])((function() use ($mempty_1_0) {
  $__fn = function($v_3 = null, $v1_4 = null, $v2_5 = null, $v3_6 = null, $v4_7 = null, $v5_8 = null, $v6_9 = null, $v7_10 = null) use ($mempty_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 8) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 8);
  }
  $__res = $mempty_1_0;
  goto __end;;
  __end:
  return $__num > 8 ? $__res(...\array_slice(\func_get_args(), 8)) : $__res;
  };
  return $__fn;
})()), "Semigroup0" => function($_dollar__unused_3 = null) use ($semigroupEffectFn81_2_1) {
  $__num = \func_num_args();
  $__res = $semigroupEffectFn81_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Effect_Uncurried_monoidEffectFn8'] = __NAMESPACE__ . '\\majEffect_majUncurried_monoidmajEffectmajFn8';

// Effect_Uncurried_monoidEffectFn7
function majEffect_majUncurried_monoidmajEffectmajFn7($dictMonoid_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majUncurried_monoidmajEffectmajFn7';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $mempty_1_0 = (($GLOBALS['Effect_monoidEffect'])($dictMonoid_0))['mempty'];
  $semigroupEffectFn71_2_1 = ($GLOBALS['Effect_Uncurried_semigroupEffectFn7'])((($dictMonoid_0)['Semigroup0'])(null));
  $__res = ["mempty" => ($GLOBALS['Effect_Uncurried_mkEffectFn7'])((function() use ($mempty_1_0) {
  $__fn = function($v_3 = null, $v1_4 = null, $v2_5 = null, $v3_6 = null, $v4_7 = null, $v5_8 = null, $v6_9 = null) use ($mempty_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 7) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 7);
  }
  $__res = $mempty_1_0;
  goto __end;;
  __end:
  return $__num > 7 ? $__res(...\array_slice(\func_get_args(), 7)) : $__res;
  };
  return $__fn;
})()), "Semigroup0" => function($_dollar__unused_3 = null) use ($semigroupEffectFn71_2_1) {
  $__num = \func_num_args();
  $__res = $semigroupEffectFn71_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Effect_Uncurried_monoidEffectFn7'] = __NAMESPACE__ . '\\majEffect_majUncurried_monoidmajEffectmajFn7';

// Effect_Uncurried_monoidEffectFn6
function majEffect_majUncurried_monoidmajEffectmajFn6($dictMonoid_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majUncurried_monoidmajEffectmajFn6';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $mempty_1_0 = (($GLOBALS['Effect_monoidEffect'])($dictMonoid_0))['mempty'];
  $semigroupEffectFn61_2_1 = ($GLOBALS['Effect_Uncurried_semigroupEffectFn6'])((($dictMonoid_0)['Semigroup0'])(null));
  $__res = ["mempty" => ($GLOBALS['Effect_Uncurried_mkEffectFn6'])((function() use ($mempty_1_0) {
  $__fn = function($v_3 = null, $v1_4 = null, $v2_5 = null, $v3_6 = null, $v4_7 = null, $v5_8 = null) use ($mempty_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 6) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 6);
  }
  $__res = $mempty_1_0;
  goto __end;;
  __end:
  return $__num > 6 ? $__res(...\array_slice(\func_get_args(), 6)) : $__res;
  };
  return $__fn;
})()), "Semigroup0" => function($_dollar__unused_3 = null) use ($semigroupEffectFn61_2_1) {
  $__num = \func_num_args();
  $__res = $semigroupEffectFn61_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Effect_Uncurried_monoidEffectFn6'] = __NAMESPACE__ . '\\majEffect_majUncurried_monoidmajEffectmajFn6';

// Effect_Uncurried_monoidEffectFn5
function majEffect_majUncurried_monoidmajEffectmajFn5($dictMonoid_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majUncurried_monoidmajEffectmajFn5';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $mempty_1_0 = (($GLOBALS['Effect_monoidEffect'])($dictMonoid_0))['mempty'];
  $semigroupEffectFn51_2_1 = ($GLOBALS['Effect_Uncurried_semigroupEffectFn5'])((($dictMonoid_0)['Semigroup0'])(null));
  $__res = ["mempty" => ($GLOBALS['Effect_Uncurried_mkEffectFn5'])((function() use ($mempty_1_0) {
  $__fn = function($v_3 = null, $v1_4 = null, $v2_5 = null, $v3_6 = null, $v4_7 = null) use ($mempty_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 5) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 5);
  }
  $__res = $mempty_1_0;
  goto __end;;
  __end:
  return $__num > 5 ? $__res(...\array_slice(\func_get_args(), 5)) : $__res;
  };
  return $__fn;
})()), "Semigroup0" => function($_dollar__unused_3 = null) use ($semigroupEffectFn51_2_1) {
  $__num = \func_num_args();
  $__res = $semigroupEffectFn51_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Effect_Uncurried_monoidEffectFn5'] = __NAMESPACE__ . '\\majEffect_majUncurried_monoidmajEffectmajFn5';

// Effect_Uncurried_monoidEffectFn4
function majEffect_majUncurried_monoidmajEffectmajFn4($dictMonoid_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majUncurried_monoidmajEffectmajFn4';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $mempty_1_0 = (($GLOBALS['Effect_monoidEffect'])($dictMonoid_0))['mempty'];
  $semigroupEffectFn41_2_1 = ($GLOBALS['Effect_Uncurried_semigroupEffectFn4'])((($dictMonoid_0)['Semigroup0'])(null));
  $__res = ["mempty" => ($GLOBALS['Effect_Uncurried_mkEffectFn4'])((function() use ($mempty_1_0) {
  $__fn = function($v_3 = null, $v1_4 = null, $v2_5 = null, $v3_6 = null) use ($mempty_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = $mempty_1_0;
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})()), "Semigroup0" => function($_dollar__unused_3 = null) use ($semigroupEffectFn41_2_1) {
  $__num = \func_num_args();
  $__res = $semigroupEffectFn41_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Effect_Uncurried_monoidEffectFn4'] = __NAMESPACE__ . '\\majEffect_majUncurried_monoidmajEffectmajFn4';

// Effect_Uncurried_monoidEffectFn3
function majEffect_majUncurried_monoidmajEffectmajFn3($dictMonoid_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majUncurried_monoidmajEffectmajFn3';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $mempty_1_0 = (($GLOBALS['Effect_monoidEffect'])($dictMonoid_0))['mempty'];
  $semigroupEffectFn31_2_1 = ($GLOBALS['Effect_Uncurried_semigroupEffectFn3'])((($dictMonoid_0)['Semigroup0'])(null));
  $__res = ["mempty" => ($GLOBALS['Effect_Uncurried_mkEffectFn3'])((function() use ($mempty_1_0) {
  $__fn = function($v_3 = null, $v1_4 = null, $v2_5 = null) use ($mempty_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = $mempty_1_0;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()), "Semigroup0" => function($_dollar__unused_3 = null) use ($semigroupEffectFn31_2_1) {
  $__num = \func_num_args();
  $__res = $semigroupEffectFn31_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Effect_Uncurried_monoidEffectFn3'] = __NAMESPACE__ . '\\majEffect_majUncurried_monoidmajEffectmajFn3';

// Effect_Uncurried_monoidEffectFn2
function majEffect_majUncurried_monoidmajEffectmajFn2($dictMonoid_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majUncurried_monoidmajEffectmajFn2';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $mempty_1_0 = (($GLOBALS['Effect_monoidEffect'])($dictMonoid_0))['mempty'];
  $semigroupEffectFn21_2_1 = ($GLOBALS['Effect_Uncurried_semigroupEffectFn2'])((($dictMonoid_0)['Semigroup0'])(null));
  $__res = ["mempty" => ($GLOBALS['Effect_Uncurried_mkEffectFn2'])((function() use ($mempty_1_0) {
  $__fn = function($v_3 = null, $v1_4 = null) use ($mempty_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = $mempty_1_0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()), "Semigroup0" => function($_dollar__unused_3 = null) use ($semigroupEffectFn21_2_1) {
  $__num = \func_num_args();
  $__res = $semigroupEffectFn21_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Effect_Uncurried_monoidEffectFn2'] = __NAMESPACE__ . '\\majEffect_majUncurried_monoidmajEffectmajFn2';

// Effect_Uncurried_monoidEffectFn10
function majEffect_majUncurried_monoidmajEffectmajFn10($dictMonoid_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majUncurried_monoidmajEffectmajFn10';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $mempty_1_0 = (($GLOBALS['Effect_monoidEffect'])($dictMonoid_0))['mempty'];
  $semigroupEffectFn101_2_1 = ($GLOBALS['Effect_Uncurried_semigroupEffectFn10'])((($dictMonoid_0)['Semigroup0'])(null));
  $__res = ["mempty" => ($GLOBALS['Effect_Uncurried_mkEffectFn10'])((function() use ($mempty_1_0) {
  $__fn = function($v_3 = null, $v1_4 = null, $v2_5 = null, $v3_6 = null, $v4_7 = null, $v5_8 = null, $v6_9 = null, $v7_10 = null, $v8_11 = null, $v9_12 = null) use ($mempty_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 10) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 10);
  }
  $__res = $mempty_1_0;
  goto __end;;
  __end:
  return $__num > 10 ? $__res(...\array_slice(\func_get_args(), 10)) : $__res;
  };
  return $__fn;
})()), "Semigroup0" => function($_dollar__unused_3 = null) use ($semigroupEffectFn101_2_1) {
  $__num = \func_num_args();
  $__res = $semigroupEffectFn101_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Effect_Uncurried_monoidEffectFn10'] = __NAMESPACE__ . '\\majEffect_majUncurried_monoidmajEffectmajFn10';

// Effect_Uncurried_monoidEffectFn1
function majEffect_majUncurried_monoidmajEffectmajFn1($dictMonoid_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majUncurried_monoidmajEffectmajFn1';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $mempty_1_0 = (($GLOBALS['Effect_monoidEffect'])($dictMonoid_0))['mempty'];
  $semigroupEffectFn11_2_1 = ($GLOBALS['Effect_Uncurried_semigroupEffectFn1'])((($dictMonoid_0)['Semigroup0'])(null));
  $__res = ["mempty" => ($GLOBALS['Effect_Uncurried_mkEffectFn1'])(function($v_3 = null) use ($mempty_1_0) {
  $__num = \func_num_args();
  $__res = $mempty_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), "Semigroup0" => function($_dollar__unused_3 = null) use ($semigroupEffectFn11_2_1) {
  $__num = \func_num_args();
  $__res = $semigroupEffectFn11_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Effect_Uncurried_monoidEffectFn1'] = __NAMESPACE__ . '\\majEffect_majUncurried_monoidmajEffectmajFn1';

