<?php

namespace Test\StateMonad;

// ALL IMPORTS: Bench, Control.Bind, Data.Function, Data.Ring, Data.Semiring, Data.Show, Data.Unit, Effect, Effect.Console, Prelude, Prim, Test.StateMonad
// TO REQUIRE: Bench, Control.Bind, Data.Function, Data.Ring, Data.Semiring, Data.Show, Data.Unit, Effect, Effect.Console, Prelude, Test.StateMonad
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
require_once __DIR__ . '/../Test.StateMonad/index.php';

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




// Test_StateMonad_State
function majTest_majStatemajMonad_majState($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majTest_majStatemajMonad_majState';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Test_StateMonad_State'] = __NAMESPACE__ . '\\majTest_majStatemajMonad_majState';

// Test_StateMonad_runState
function majTest_majStatemajMonad_runmajState($v_0, $s_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majTest_majStatemajMonad_runmajState';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($v_0)($s_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Test_StateMonad_runState'] = __NAMESPACE__ . '\\majTest_majStatemajMonad_runmajState';

// Test_StateMonad_put
function majTest_majStatemajMonad_put($s_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majTest_majStatemajMonad_put';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ["val" => $GLOBALS['Data_Unit_unit'], "state" => $s_0];
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Test_StateMonad_put'] = __NAMESPACE__ . '\\majTest_majStatemajMonad_put';

// Test_StateMonad_pureState
function majTest_majStatemajMonad_puremajState($a_0, $s_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majTest_majStatemajMonad_puremajState';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ["val" => $a_0, "state" => $s_1];
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Test_StateMonad_pureState'] = __NAMESPACE__ . '\\majTest_majStatemajMonad_puremajState';

// Test_StateMonad_get
function majTest_majStatemajMonad_get($s_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majTest_majStatemajMonad_get';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ["val" => $s_0, "state" => $s_0];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Test_StateMonad_get'] = __NAMESPACE__ . '\\majTest_majStatemajMonad_get';

// Test_StateMonad_describe
$GLOBALS['Test_StateMonad_describe'] = ($GLOBALS['Effect_Console_log'])("State Monad (1.2k Binds, 60 Stack Depth):");

// Test_StateMonad_bindState
function majTest_majStatemajMonad_bindmajState($v_0, $g_1 = null, $s_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majTest_majStatemajMonad_bindmajState';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $r1_3_0 = ($v_0)($s_2);
  $__res = ($g_1)(($r1_3_0)['val'], ($r1_3_0)['state']);
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Test_StateMonad_bindState'] = __NAMESPACE__ . '\\majTest_majStatemajMonad_bindmajState';

// Test_StateMonad_modify
function majTest_majStatemajMonad_modify($f_0, $s_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majTest_majStatemajMonad_modify';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ["val" => $GLOBALS['Data_Unit_unit'], "state" => ($f_0)($s_1)];
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Test_StateMonad_modify'] = __NAMESPACE__ . '\\majTest_majStatemajMonad_modify';

// Test_StateMonad_chainModifications
function majTest_majStatemajMonad_chainmajModifications($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majTest_majStatemajMonad_chainmajModifications';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__tco_var_Test_StateMonad_chainModifications_v_0 = $v_0;
  tco_loop_Test_StateMonad_chainModifications:;
  $v_0 = $__tco_var_Test_StateMonad_chainModifications_v_0;
  $__res = match ($v_0) { 0 => function($s_1 = null) {
  $__num = \func_num_args();
  $__res = ["val" => $GLOBALS['Data_Unit_unit'], "state" => $s_1];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, default => function($s_1 = null) use ($v_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Test_StateMonad_chainModifications'])(($v_0 - 1), ($s_1 + 1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
} };
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Test_StateMonad_chainModifications'] = __NAMESPACE__ . '\\majTest_majStatemajMonad_chainmajModifications';

// Test_StateMonad_runManyTimes
function majTest_majStatemajMonad_runmajManymajTimes($v_0, $v1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majTest_majStatemajMonad_runmajManymajTimes';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_Test_StateMonad_runManyTimes_v_0 = $v_0;
  $__tco_var_Test_StateMonad_runManyTimes_v1_1 = $v1_1;
  tco_loop_Test_StateMonad_runManyTimes:;
  $v_0 = $__tco_var_Test_StateMonad_runManyTimes_v_0;
  $v1_1 = $__tco_var_Test_StateMonad_runManyTimes_v1_1;
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
  $__tco_1 = ($v1_1 + (($GLOBALS['Test_StateMonad_chainModifications'])(60, 0))['state']);
  $__tco_var_Test_StateMonad_runManyTimes_v_0 = $__tco_0;
  $__tco_var_Test_StateMonad_runManyTimes_v1_1 = $__tco_1;
  goto tco_loop_Test_StateMonad_runManyTimes;;
  $__t2 = null;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Test_StateMonad_runManyTimes'] = __NAMESPACE__ . '\\majTest_majStatemajMonad_runmajManymajTimes';

// Test_StateMonad_act
$GLOBALS['Test_StateMonad_act'] = (($GLOBALS['Effect_bindEffect'])['bind'])(($GLOBALS['Bench_opaque'])(20), function($dummy_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Effect_Console_log'])((($GLOBALS['Data_Show_showInt'])['show'])(($GLOBALS['Test_StateMonad_runManyTimes'])($dummy_0, 0)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});

