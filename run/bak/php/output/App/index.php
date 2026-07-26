<?php

namespace App;

// ALL IMPORTS: Bench, Control.Bind, Data.EuclideanRing, Data.Function, Data.Semigroup, Data.Semiring, Effect, Effect.Console, Prelude, Prim, Test.Ackermann, Test.ArrayOps, Test.AstTree, Test.Church, Test.Fib, Test.LazyEvaluation, Test.ListOps, Test.Polymorphism, Test.Primes, Test.RBTree, Test.Records, Test.StateMonad, Test.TCO
// TO REQUIRE: Bench, Control.Bind, Data.EuclideanRing, Data.Function, Data.Semigroup, Data.Semiring, Effect, Effect.Console, Prelude, Test.Ackermann, Test.ArrayOps, Test.AstTree, Test.Church, Test.Fib, Test.LazyEvaluation, Test.ListOps, Test.Polymorphism, Test.Primes, Test.RBTree, Test.Records, Test.StateMonad, Test.TCO
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

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };


// App_main
$GLOBALS['App_main'] = ((($GLOBALS['Effect_bindEffect'])['bind'])((($GLOBALS['Bench_runBench'])($GLOBALS['Test_AstTree_describe']))($GLOBALS['Test_AstTree_act'])))(function($t1_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Effect_bindEffect'])['bind'])((($GLOBALS['Bench_runBench'])($GLOBALS['Test_Fib_describe']))($GLOBALS['Test_Fib_act'])))(function($t2_1 = null) use ($t1_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Effect_bindEffect'])['bind'])((($GLOBALS['Bench_runBench'])($GLOBALS['Test_ListOps_describe']))($GLOBALS['Test_ListOps_act'])))(function($t3_2 = null) use ($t1_0, $t2_1) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Effect_bindEffect'])['bind'])((($GLOBALS['Bench_runBench'])($GLOBALS['Test_TCO_describe']))($GLOBALS['Test_TCO_act'])))(function($t4_3 = null) use ($t1_0, $t2_1, $t3_2) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Effect_bindEffect'])['bind'])((($GLOBALS['Bench_runBench'])($GLOBALS['Test_Records_describe']))($GLOBALS['Test_Records_act'])))(function($t5_4 = null) use ($t1_0, $t2_1, $t3_2, $t4_3) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Effect_bindEffect'])['bind'])((($GLOBALS['Bench_runBench'])($GLOBALS['Test_Ackermann_describe']))($GLOBALS['Test_Ackermann_act'])))(function($t6_5 = null) use ($t1_0, $t2_1, $t3_2, $t4_3, $t5_4) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Effect_bindEffect'])['bind'])((($GLOBALS['Bench_runBench'])($GLOBALS['Test_Church_describe']))($GLOBALS['Test_Church_act'])))(function($t7_6 = null) use ($t1_0, $t2_1, $t3_2, $t4_3, $t5_4, $t6_5) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Effect_bindEffect'])['bind'])((($GLOBALS['Bench_runBench'])($GLOBALS['Test_Primes_describe']))($GLOBALS['Test_Primes_act'])))(function($t8_7 = null) use ($t1_0, $t2_1, $t3_2, $t4_3, $t5_4, $t6_5, $t7_6) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Effect_bindEffect'])['bind'])((($GLOBALS['Bench_runBench'])($GLOBALS['Test_RBTree_describe']))($GLOBALS['Test_RBTree_act'])))(function($t9_8 = null) use ($t1_0, $t2_1, $t3_2, $t4_3, $t5_4, $t6_5, $t7_6, $t8_7) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Effect_bindEffect'])['bind'])((($GLOBALS['Bench_runBench'])($GLOBALS['Test_Polymorphism_describe']))($GLOBALS['Test_Polymorphism_act'])))(function($t10_9 = null) use ($t1_0, $t2_1, $t3_2, $t4_3, $t5_4, $t6_5, $t7_6, $t8_7, $t9_8) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Effect_bindEffect'])['bind'])((($GLOBALS['Bench_runBench'])($GLOBALS['Test_StateMonad_describe']))($GLOBALS['Test_StateMonad_act'])))(function($t11_10 = null) use ($t10_9, $t1_0, $t2_1, $t3_2, $t4_3, $t5_4, $t6_5, $t7_6, $t8_7, $t9_8) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Effect_bindEffect'])['bind'])((($GLOBALS['Bench_runBench'])($GLOBALS['Test_LazyEvaluation_describe']))($GLOBALS['Test_LazyEvaluation_act'])))(function($t12_11 = null) use ($t10_9, $t11_10, $t1_0, $t2_1, $t3_2, $t4_3, $t5_4, $t6_5, $t7_6, $t8_7, $t9_8) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Effect_bindEffect'])['bind'])((($GLOBALS['Bench_runBench'])($GLOBALS['Test_ArrayOps_describe']))($GLOBALS['Test_ArrayOps_act'])))(function($t13_12 = null) use ($t10_9, $t11_10, $t12_11, $t1_0, $t2_1, $t3_2, $t4_3, $t5_4, $t6_5, $t7_6, $t8_7, $t9_8) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Effect_Console_log'])(((($GLOBALS['Data_Semigroup_semigroupString'])['append'])("Total exec time: "))(((($GLOBALS['Data_Semigroup_semigroupString'])['append'])(($GLOBALS['Bench_formatNumber'])(((($GLOBALS['Data_Semiring_semiringNumber'])['add'])(((($GLOBALS['Data_Semiring_semiringNumber'])['add'])(((($GLOBALS['Data_Semiring_semiringNumber'])['add'])(((($GLOBALS['Data_Semiring_semiringNumber'])['add'])(((($GLOBALS['Data_Semiring_semiringNumber'])['add'])(((($GLOBALS['Data_Semiring_semiringNumber'])['add'])(((($GLOBALS['Data_Semiring_semiringNumber'])['add'])(((($GLOBALS['Data_Semiring_semiringNumber'])['add'])(((($GLOBALS['Data_Semiring_semiringNumber'])['add'])(((($GLOBALS['Data_Semiring_semiringNumber'])['add'])(((($GLOBALS['Data_Semiring_semiringNumber'])['add'])(((($GLOBALS['Data_Semiring_semiringNumber'])['add'])(((($GLOBALS['Data_EuclideanRing_euclideanRingNumber'])['div'])($t1_0))(1000.0)))(((($GLOBALS['Data_EuclideanRing_euclideanRingNumber'])['div'])($t2_1))(1000.0))))(((($GLOBALS['Data_EuclideanRing_euclideanRingNumber'])['div'])($t3_2))(1000.0))))(((($GLOBALS['Data_EuclideanRing_euclideanRingNumber'])['div'])($t4_3))(1000.0))))(((($GLOBALS['Data_EuclideanRing_euclideanRingNumber'])['div'])($t5_4))(1000.0))))(((($GLOBALS['Data_EuclideanRing_euclideanRingNumber'])['div'])($t6_5))(1000.0))))(((($GLOBALS['Data_EuclideanRing_euclideanRingNumber'])['div'])($t7_6))(1000.0))))(((($GLOBALS['Data_EuclideanRing_euclideanRingNumber'])['div'])($t8_7))(1000.0))))(((($GLOBALS['Data_EuclideanRing_euclideanRingNumber'])['div'])($t9_8))(1000.0))))(((($GLOBALS['Data_EuclideanRing_euclideanRingNumber'])['div'])($t10_9))(1000.0))))(((($GLOBALS['Data_EuclideanRing_euclideanRingNumber'])['div'])($t11_10))(1000.0))))(((($GLOBALS['Data_EuclideanRing_euclideanRingNumber'])['div'])($t12_11))(1000.0))))(((($GLOBALS['Data_EuclideanRing_euclideanRingNumber'])['div'])($t13_12))(1000.0)))))(" ms
")));
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

