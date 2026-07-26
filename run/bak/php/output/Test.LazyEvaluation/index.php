<?php

namespace Test\LazyEvaluation;

// ALL IMPORTS: Bench, Control.Bind, Data.Function, Data.Ring, Data.Semiring, Data.Show, Data.Unit, Effect, Effect.Console, Prelude, Prim, Test.LazyEvaluation
// TO REQUIRE: Bench, Control.Bind, Data.Function, Data.Ring, Data.Semiring, Data.Show, Data.Unit, Effect, Effect.Console, Prelude, Test.LazyEvaluation
require_once __DIR__ . '/../Bench/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Ring/index.php';
require_once __DIR__ . '/../Data.Semiring/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
require_once __DIR__ . '/../Data.Unit/index.php';
require_once __DIR__ . '/../Effect/index.php';
require_once __DIR__ . '/../Effect.Console/index.php';
require_once __DIR__ . '/../Prelude/index.php';
require_once __DIR__ . '/../Test.LazyEvaluation/index.php';

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




// Test_LazyEvaluation_Lazy
function majTest_majLazymajEvaluation_majLazy($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majTest_majLazymajEvaluation_majLazy';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Test_LazyEvaluation_Lazy'] = __NAMESPACE__ . '\\majTest_majLazymajEvaluation_majLazy';

// Test_LazyEvaluation_force
function majTest_majLazymajEvaluation_force($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majTest_majLazymajEvaluation_force';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($v_0)($GLOBALS['Data_Unit_unit']);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Test_LazyEvaluation_force'] = __NAMESPACE__ . '\\majTest_majLazymajEvaluation_force';

// Test_LazyEvaluation_describe
$GLOBALS['Test_LazyEvaluation_describe'] = ($GLOBALS['Effect_Console_log'])("Lazy Evaluation (1M Thunks Forced, 1k Depth):");

// Test_LazyEvaluation_defer_closure
$GLOBALS['Test_LazyEvaluation_defer_closure'] = $GLOBALS['Test_LazyEvaluation_Lazy'];

// Test_LazyEvaluation_defer
function majTest_majLazymajEvaluation_defer($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majTest_majLazymajEvaluation_defer';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Test_LazyEvaluation_defer_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Test_LazyEvaluation_defer'] = __NAMESPACE__ . '\\majTest_majLazymajEvaluation_defer';

// Test_LazyEvaluation_buildThunks
function majTest_majLazymajEvaluation_buildmajThunks($v_0, $v1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majTest_majLazymajEvaluation_buildmajThunks';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_Test_LazyEvaluation_buildThunks_v_0 = $v_0;
  $__tco_var_Test_LazyEvaluation_buildThunks_v1_1 = $v1_1;
  tco_loop_Test_LazyEvaluation_buildThunks:;
  $v_0 = $__tco_var_Test_LazyEvaluation_buildThunks_v_0;
  $v1_1 = $__tco_var_Test_LazyEvaluation_buildThunks_v1_1;
  $__t2 = null;;
  switch ($v_0) {
case 0:
$__t2 = $v1_1;
goto end_branch_2;;
break;
default:
;
break;
};
  $__tco_0 = ($v_0 - 1);
  $__tco_1 = function($v2_2 = null) use ($v1_1) {
  $__num = \func_num_args();
  $__res = (($v1_1)($GLOBALS['Data_Unit_unit']) + 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__tco_var_Test_LazyEvaluation_buildThunks_v_0 = $__tco_0;
  $__tco_var_Test_LazyEvaluation_buildThunks_v1_1 = $__tco_1;
  goto tco_loop_Test_LazyEvaluation_buildThunks;;
  $__t2 = null;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Test_LazyEvaluation_buildThunks'] = __NAMESPACE__ . '\\majTest_majLazymajEvaluation_buildmajThunks';

// Test_LazyEvaluation_runManyTimes
function majTest_majLazymajEvaluation_runmajManymajTimes($v_0, $v1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majTest_majLazymajEvaluation_runmajManymajTimes';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_Test_LazyEvaluation_runManyTimes_v_0 = $v_0;
  $__tco_var_Test_LazyEvaluation_runManyTimes_v1_1 = $v1_1;
  tco_loop_Test_LazyEvaluation_runManyTimes:;
  $v_0 = $__tco_var_Test_LazyEvaluation_runManyTimes_v_0;
  $v1_1 = $__tco_var_Test_LazyEvaluation_runManyTimes_v1_1;
  $__t2 = null;;
  switch ($v_0) {
case 0:
$__t2 = $v1_1;
goto end_branch_2;;
break;
default:
;
break;
};
  $__tco_0 = ($v_0 - 1);
  $__tco_1 = ($v1_1 + ($GLOBALS['Test_LazyEvaluation_buildThunks'])(1000, function($v2_2 = null) {
  $__num = \func_num_args();
  $__res = 0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $GLOBALS['Data_Unit_unit']));
  $__tco_var_Test_LazyEvaluation_runManyTimes_v_0 = $__tco_0;
  $__tco_var_Test_LazyEvaluation_runManyTimes_v1_1 = $__tco_1;
  goto tco_loop_Test_LazyEvaluation_runManyTimes;;
  $__t2 = null;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Test_LazyEvaluation_runManyTimes'] = __NAMESPACE__ . '\\majTest_majLazymajEvaluation_runmajManymajTimes';

// Test_LazyEvaluation_act
$GLOBALS['Test_LazyEvaluation_act'] = (($GLOBALS['Effect_bindEffect'])['bind'])(($GLOBALS['Bench_opaque'])(1000), function($dummy_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Effect_Console_log'])((($GLOBALS['Data_Show_showInt'])['show'])(($GLOBALS['Test_LazyEvaluation_runManyTimes'])($dummy_0, 0)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});

