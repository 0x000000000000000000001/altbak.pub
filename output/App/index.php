<?php

namespace App;

// ALL IMPORTS: Bench, Control.Bind, Data.EuclideanRing, Data.Function, Data.Semigroup, Data.Semiring, Effect, Effect.Console, Prelude, Prim, Test.Ackermann, Test.ArrayOps, Test.AstTree, Test.Church, Test.Fib, Test.LazyEvaluation, Test.ListOps, Test.Polymorphism, Test.Primes, Test.RBTree, Test.Records, Test.RowToList, Test.StateMonad, Test.TCO
// TO REQUIRE: Bench, Control.Bind, Data.EuclideanRing, Data.Function, Data.Semigroup, Data.Semiring, Effect, Effect.Console, Prelude, Test.Ackermann, Test.ArrayOps, Test.AstTree, Test.Church, Test.Fib, Test.LazyEvaluation, Test.ListOps, Test.Polymorphism, Test.Primes, Test.RBTree, Test.Records, Test.RowToList, Test.StateMonad, Test.TCO
require_once __DIR__ . '/../Bench/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Data.EuclideanRing/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Semiring/index.php';
require_once __DIR__ . '/../Effect/index.php';
require_once __DIR__ . '/../Effect.Console/index.php';
require_once __DIR__ . '/../Prelude/index.php';
require_once __DIR__ . '/../Test.Ackermann/index.php';
require_once __DIR__ . '/../Test.ArrayOps/index.php';
require_once __DIR__ . '/../Test.AstTree/index.php';
require_once __DIR__ . '/../Test.Church/index.php';
require_once __DIR__ . '/../Test.Fib/index.php';
require_once __DIR__ . '/../Test.LazyEvaluation/index.php';
require_once __DIR__ . '/../Test.ListOps/index.php';
require_once __DIR__ . '/../Test.Polymorphism/index.php';
require_once __DIR__ . '/../Test.Primes/index.php';
require_once __DIR__ . '/../Test.RBTree/index.php';
require_once __DIR__ . '/../Test.Records/index.php';
require_once __DIR__ . '/../Test.RowToList/index.php';
require_once __DIR__ . '/../Test.StateMonad/index.php';
require_once __DIR__ . '/../Test.TCO/index.php';

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




// App_main
$GLOBALS['App_main'] = (function() use (&$__fn) {
$__local_var_0_0 = \Bench\majBench_runmajBench($GLOBALS['Test_AstTree_describe'], $GLOBALS['Test_AstTree_act']);
return function() use ($__local_var_0_0, &$__fn) {
$t1_1_1 = phpurs_execute_effect($__local_var_0_0);
$t2_2_2 = phpurs_execute_effect(\Bench\majBench_runmajBench($GLOBALS['Test_Fib_describe'], $GLOBALS['Test_Fib_act']));
$t3_3_3 = phpurs_execute_effect(\Bench\majBench_runmajBench($GLOBALS['Test_ListOps_describe'], $GLOBALS['Test_ListOps_act']));
$t4_4_4 = phpurs_execute_effect(\Bench\majBench_runmajBench($GLOBALS['Test_TCO_describe'], $GLOBALS['Test_TCO_act']));
$t5_5_5 = phpurs_execute_effect(\Bench\majBench_runmajBench($GLOBALS['Test_Records_describe'], $GLOBALS['Test_Records_act']));
$t6_6_6 = phpurs_execute_effect(\Bench\majBench_runmajBench($GLOBALS['Test_Ackermann_describe'], $GLOBALS['Test_Ackermann_act']));
$t7_7_7 = phpurs_execute_effect(\Bench\majBench_runmajBench($GLOBALS['Test_Church_describe'], $GLOBALS['Test_Church_act']));
$t8_8_8 = phpurs_execute_effect(\Bench\majBench_runmajBench($GLOBALS['Test_Primes_describe'], $GLOBALS['Test_Primes_act']));
$t9_9_9 = phpurs_execute_effect(\Bench\majBench_runmajBench($GLOBALS['Test_RBTree_describe'], $GLOBALS['Test_RBTree_act']));
$t10_10_10 = phpurs_execute_effect(\Bench\majBench_runmajBench($GLOBALS['Test_Polymorphism_describe'], $GLOBALS['Test_Polymorphism_act']));
$t11_11_11 = phpurs_execute_effect(\Bench\majBench_runmajBench($GLOBALS['Test_StateMonad_describe'], $GLOBALS['Test_StateMonad_act']));
$t12_12_12 = phpurs_execute_effect(\Bench\majBench_runmajBench($GLOBALS['Test_LazyEvaluation_describe'], $GLOBALS['Test_LazyEvaluation_act']));
$t13_13_13 = phpurs_execute_effect(\Bench\majBench_runmajBench($GLOBALS['Test_ArrayOps_describe'], $GLOBALS['Test_ArrayOps_act']));
$t14_14_14 = phpurs_execute_effect(\Bench\majBench_runmajBench($GLOBALS['Test_RowToList_describe'], $GLOBALS['Test_RowToList_act']));
return phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(\Effect\Console\majEffect_majConsole_log((("
==================================================

Total exec time: " . \Bench\majBench_formatmajNumber((((((((((((((($t1_1_1 / 1000.0) + ($t2_2_2 / 1000.0)) + ($t3_3_3 / 1000.0)) + ($t4_4_4 / 1000.0)) + ($t5_5_5 / 1000.0)) + ($t6_6_6 / 1000.0)) + ($t7_7_7 / 1000.0)) + ($t8_8_8 / 1000.0)) + ($t9_9_9 / 1000.0)) + ($t10_10_10 / 1000.0)) + ($t11_11_11 / 1000.0)) + ($t12_12_12 / 1000.0)) + ($t13_13_13 / 1000.0)) + ($t14_14_14 / 1000.0)))) . " ms
"))))))))))))))));
};
})();

