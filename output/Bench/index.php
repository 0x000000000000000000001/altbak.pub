<?php

namespace Bench;

// ALL IMPORTS: Bench, Control.Applicative, Control.Bind, Data.Function, Data.Ring, Data.Semigroup, Effect, Effect.Aff, Effect.Class, Effect.Console, Prelude, Prim
// TO REQUIRE: Bench, Control.Applicative, Control.Bind, Data.Function, Data.Ring, Data.Semigroup, Effect, Effect.Aff, Effect.Class, Effect.Console, Prelude
require_once __DIR__ . '/../Bench/index.php';
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Ring/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Effect/index.php';
require_once __DIR__ . '/../Effect.Aff/index.php';
require_once __DIR__ . '/../Effect.Class/index.php';
require_once __DIR__ . '/../Effect.Console/index.php';
require_once __DIR__ . '/../Prelude/index.php';

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
$ffi_Bench = \call_user_func(function() {
  $exports = [];
$benchNow = function() {
    return microtime(true) * 1000000.0;
};

$opaque = function($a) {
    return function() use($a) {
        return $a;
    };
};

$formatNumber = function($n) {
    return number_format($n, 2, '.', '');
};

$exports['benchNow'] = $benchNow;
$exports['opaque'] = $opaque;
$exports['formatNumber'] = $formatNumber;

return $exports;
  return $exports;
});
$GLOBALS['Bench_benchNow'] = (\array_key_exists('benchNow', $ffi_Bench) ? $ffi_Bench['benchNow'] : new class { public function __invoke(...$args) { return $this; } });
function majBench_formatmajNumber(float $v0): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majBench_formatmajNumber';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Bench;
  $f = (\array_key_exists('formatNumber', $ffi_Bench) ? $ffi_Bench['formatNumber'] : new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Bench_formatNumber'] = __NAMESPACE__ . '\\majBench_formatmajNumber';

function majBench_opaque($v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majBench_opaque';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Bench;
  $f = (\array_key_exists('opaque', $ffi_Bench) ? $ffi_Bench['opaque'] : new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Bench_opaque'] = __NAMESPACE__ . '\\majBench_opaque';





// Bench_runBenchAff
function majBench_runmajBenchmajAff($describe_0, $act_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majBench_runmajBenchmajAff';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Effect\Aff\majEffect_majAff__bind(\Effect\Aff\majEffect_majAff__liftmajEffect(\Effect\Console\majEffect_majConsole_log("--------------------------------------------------

(Test)
")), function($_dollar___unused_2) use ($act_1, $describe_0) {
  $__num = \func_num_args();
  $__res = \Effect\Aff\majEffect_majAff__bind(\Effect\Aff\majEffect_majAff__liftmajEffect($describe_0), function($_dollar___unused_3) use ($act_1) {
  $__num = \func_num_args();
  $__res = \Effect\Aff\majEffect_majAff__bind(\Effect\Aff\majEffect_majAff__liftmajEffect(\Effect\Console\majEffect_majConsole_log("
(Output)
")), function($_dollar___unused_4) use ($act_1) {
  $__num = \func_num_args();
  $__res = \Effect\Aff\majEffect_majAff__bind(\Effect\Aff\majEffect_majAff__liftmajEffect($GLOBALS['Bench_benchNow']), function($t1_5) use ($act_1) {
  $__num = \func_num_args();
  $__res = \Effect\Aff\majEffect_majAff__bind($act_1, function($_dollar___unused_6) use ($t1_5) {
  $__num = \func_num_args();
  $__res = \Effect\Aff\majEffect_majAff__bind(\Effect\Aff\majEffect_majAff__liftmajEffect($GLOBALS['Bench_benchNow']), function($t2_7) use ($t1_5) {
  $__num = \func_num_args();
  $dt_8_0 = ($t2_7 - $t1_5);
  $__res = \Effect\Aff\majEffect_majAff__bind(\Effect\Aff\majEffect_majAff__liftmajEffect(\Effect\Console\majEffect_majConsole_log((("
(Execution time)

" . \Bench\majBench_formatmajNumber($dt_8_0)) . " μs
"))), function($_dollar___unused_9) use ($dt_8_0) {
  $__num = \func_num_args();
  $__res = \Effect\Aff\majEffect_majAff__pure($dt_8_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Bench_runBenchAff'] = __NAMESPACE__ . '\\majBench_runmajBenchmajAff';

// Bench_runBench
function majBench_runmajBench($describe_0, $act_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majBench_runmajBench';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_2_0 = \Effect\Console\majEffect_majConsole_log("--------------------------------------------------

(Test)
");
  $__res = function() use ($__local_var_2_0, $act_1, $describe_0, &$__fn) {
$_dollar___unused_3_1 = phpurs_execute_effect($__local_var_2_0);
$_dollar___unused_4_2 = phpurs_execute_effect($describe_0);
$_dollar___unused_5_3 = phpurs_execute_effect(\Effect\Console\majEffect_majConsole_log("
(Output)
"));
$t1_6_4 = phpurs_execute_effect($GLOBALS['Bench_benchNow']);
$_dollar___unused_7_5 = phpurs_execute_effect($act_1);
$t2_8_6 = phpurs_execute_effect($GLOBALS['Bench_benchNow']);
$dt_9_7 = ($t2_8_6 - $t1_6_4);
$_dollar___unused_10_8 = phpurs_execute_effect(\Effect\Console\majEffect_majConsole_log((("
(Execution time)

" . \Bench\majBench_formatmajNumber($dt_9_7)) . " μs
")));
return phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect($dt_9_7)))))));
};
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Bench_runBench'] = __NAMESPACE__ . '\\majBench_runmajBench';

