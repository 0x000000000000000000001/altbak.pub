<?php

namespace Test\RBTree;

// ALL IMPORTS: Bench, Control.Bind, Data.Function, Data.Ord, Data.Ring, Data.Semiring, Data.Show, Effect, Effect.Console, Prelude, Prim, Test.RBTree
// TO REQUIRE: Bench, Control.Bind, Data.Function, Data.Ord, Data.Ring, Data.Semiring, Data.Show, Effect, Effect.Console, Prelude, Test.RBTree
require_once __DIR__ . '/../Bench/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ring/index.php';
require_once __DIR__ . '/../Data.Semiring/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
require_once __DIR__ . '/../Effect/index.php';
require_once __DIR__ . '/../Effect.Console/index.php';
require_once __DIR__ . '/../Prelude/index.php';
require_once __DIR__ . '/../Test.RBTree/index.php';

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


final class Test_RBTree_R { public function __construct() {} }
final class Test_RBTree_B { public function __construct() {} }
final class Test_RBTree_E { public function __construct() {} }
final class Test_RBTree_T { public function __construct(public  $value0, public  $value1, public int $value2, public  $value3) {} }

// Test_RBTree_R
$GLOBALS['Test_RBTree_R'] = ($GLOBALS['__phpurs_data0_R'] ??= new \Test\RBTree\Test_RBTree_R());

// Test_RBTree_B
$GLOBALS['Test_RBTree_B'] = ($GLOBALS['__phpurs_data0_B'] ??= new \Test\RBTree\Test_RBTree_B());

// Test_RBTree_E
$GLOBALS['Test_RBTree_E'] = ($GLOBALS['__phpurs_data0_E'] ??= new \Test\RBTree\Test_RBTree_E());

// Test_RBTree_T
$GLOBALS['Test_RBTree_T'] = (function() {
  $__fn = function($value0, $value1 = null, $value2 = null, $value3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = new \Test\RBTree\Test_RBTree_T($value0, $value1, $value2, $value3);
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})();

// Test_RBTree_max
function majTest_majRmajBmajTree_max(int $x_0, $y_1 = null): int|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majTest_majRmajBmajTree_max';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if (($x_0 > $y_1)) {
$__t0 = $x_0;
goto end_branch_0;;
};
  $__t0 = $y_1;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Test_RBTree_max'] = __NAMESPACE__ . '\\majTest_majRmajBmajTree_max';

// Test_RBTree_makeBlack
function majTest_majRmajBmajTree_makemajBlack($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majTest_majRmajBmajTree_makemajBlack';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__t0 = null;;
  if ($v_0 instanceof \Test\RBTree\Test_RBTree_T) {
$__t0 = new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_B(), ($v_0)->{'value1'}, ($v_0)->{'value2'}, ($v_0)->{'value3'});
goto end_branch_0;;
};
  if ($v_0 instanceof \Test\RBTree\Test_RBTree_E) {
$__t0 = new \Test\RBTree\Test_RBTree_E();
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Test_RBTree_makeBlack'] = __NAMESPACE__ . '\\majTest_majRmajBmajTree_makemajBlack';

// Test_RBTree_describe
$GLOBALS['Test_RBTree_describe'] = ($GLOBALS['Effect_Console_log'])("Red-Black Tree (100k Worst-Case Insertions):");

// Test_RBTree_depth
function majTest_majRmajBmajTree_depth($v_0): int|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majTest_majRmajBmajTree_depth';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__tco_var_Test_RBTree_depth_v_0 = $v_0;
  tco_loop_Test_RBTree_depth:;
  $v_0 = $__tco_var_Test_RBTree_depth_v_0;
  $__t0 = null;;
  if ($v_0 instanceof \Test\RBTree\Test_RBTree_E) {
$__t0 = 0;
goto end_branch_0;;
};
  if ($v_0 instanceof \Test\RBTree\Test_RBTree_T) {
$__local_var_1_1 = ($GLOBALS['Test_RBTree_depth'])(($v_0)->{'value1'});
$__local_var_2_2 = ($GLOBALS['Test_RBTree_depth'])(($v_0)->{'value3'});
$__t3 = null;;
if (($__local_var_1_1 > $__local_var_2_2)) {
$__t3 = (1 + $__local_var_1_1);
goto end_branch_3;;
};
$__t3 = (1 + $__local_var_2_2);
end_branch_3:;
$__t0 = $__t3;
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Test_RBTree_depth'] = __NAMESPACE__ . '\\majTest_majRmajBmajTree_depth';

// Test_RBTree_balance
function majTest_majRmajBmajTree_balance($v_0, $v1_1 = null, $v2_2 = null, $v3_3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majTest_majRmajBmajTree_balance';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__t0 = null;;
  if ($v_0 instanceof \Test\RBTree\Test_RBTree_B) {
$__t1 = null;;
if ($v1_1 instanceof \Test\RBTree\Test_RBTree_T) {
$__t2 = null;;
if (($v1_1)->{'value0'} instanceof \Test\RBTree\Test_RBTree_R) {
$__t3 = null;;
if (($v1_1)->{'value1'} instanceof \Test\RBTree\Test_RBTree_T) {
$__t4 = null;;
if ((($v1_1)->{'value1'})->{'value0'} instanceof \Test\RBTree\Test_RBTree_R) {
$__t4 = new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_R(), new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_B(), (($v1_1)->{'value1'})->{'value1'}, (($v1_1)->{'value1'})->{'value2'}, (($v1_1)->{'value1'})->{'value3'}), ($v1_1)->{'value2'}, new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_B(), ($v1_1)->{'value3'}, $v2_2, $v3_3));
goto end_branch_4;;
};
if (($v1_1)->{'value3'} instanceof \Test\RBTree\Test_RBTree_T) {
$__t5 = null;;
if ((($v1_1)->{'value3'})->{'value0'} instanceof \Test\RBTree\Test_RBTree_R) {
$__t5 = new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_R(), new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_B(), ($v1_1)->{'value1'}, ($v1_1)->{'value2'}, (($v1_1)->{'value3'})->{'value1'}), (($v1_1)->{'value3'})->{'value2'}, new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_B(), (($v1_1)->{'value3'})->{'value3'}, $v2_2, $v3_3));
goto end_branch_5;;
};
if (($v3_3 instanceof \Test\RBTree\Test_RBTree_T && ($v3_3)->{'value0'} instanceof \Test\RBTree\Test_RBTree_R)) {
$__t6 = null;;
if (($v3_3)->{'value1'} instanceof \Test\RBTree\Test_RBTree_T) {
$__t7 = null;;
if ((($v3_3)->{'value1'})->{'value0'} instanceof \Test\RBTree\Test_RBTree_R) {
$__t7 = new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_R(), new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_B(), $v1_1, $v2_2, (($v3_3)->{'value1'})->{'value1'}), (($v3_3)->{'value1'})->{'value2'}, new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_B(), (($v3_3)->{'value1'})->{'value3'}, ($v3_3)->{'value2'}, ($v3_3)->{'value3'}));
goto end_branch_7;;
};
if ((($v3_3)->{'value3'} instanceof \Test\RBTree\Test_RBTree_T && (($v3_3)->{'value3'})->{'value0'} instanceof \Test\RBTree\Test_RBTree_R)) {
$__t7 = new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_R(), new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_B(), $v1_1, $v2_2, ($v3_3)->{'value1'}), ($v3_3)->{'value2'}, new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_B(), (($v3_3)->{'value3'})->{'value1'}, (($v3_3)->{'value3'})->{'value2'}, (($v3_3)->{'value3'})->{'value3'}));
goto end_branch_7;;
};
$__t7 = new \Test\RBTree\Test_RBTree_T($v_0, $v1_1, $v2_2, $v3_3);
end_branch_7:;
$__t6 = $__t7;
goto end_branch_6;;
};
if ((($v3_3)->{'value3'} instanceof \Test\RBTree\Test_RBTree_T && (($v3_3)->{'value3'})->{'value0'} instanceof \Test\RBTree\Test_RBTree_R)) {
$__t6 = new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_R(), new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_B(), $v1_1, $v2_2, ($v3_3)->{'value1'}), ($v3_3)->{'value2'}, new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_B(), (($v3_3)->{'value3'})->{'value1'}, (($v3_3)->{'value3'})->{'value2'}, (($v3_3)->{'value3'})->{'value3'}));
goto end_branch_6;;
};
$__t6 = new \Test\RBTree\Test_RBTree_T($v_0, $v1_1, $v2_2, $v3_3);
end_branch_6:;
$__t5 = $__t6;
goto end_branch_5;;
};
$__t5 = new \Test\RBTree\Test_RBTree_T($v_0, $v1_1, $v2_2, $v3_3);
end_branch_5:;
$__t4 = $__t5;
goto end_branch_4;;
};
if (($v3_3 instanceof \Test\RBTree\Test_RBTree_T && ($v3_3)->{'value0'} instanceof \Test\RBTree\Test_RBTree_R)) {
$__t8 = null;;
if (($v3_3)->{'value1'} instanceof \Test\RBTree\Test_RBTree_T) {
$__t9 = null;;
if ((($v3_3)->{'value1'})->{'value0'} instanceof \Test\RBTree\Test_RBTree_R) {
$__t9 = new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_R(), new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_B(), $v1_1, $v2_2, (($v3_3)->{'value1'})->{'value1'}), (($v3_3)->{'value1'})->{'value2'}, new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_B(), (($v3_3)->{'value1'})->{'value3'}, ($v3_3)->{'value2'}, ($v3_3)->{'value3'}));
goto end_branch_9;;
};
if ((($v3_3)->{'value3'} instanceof \Test\RBTree\Test_RBTree_T && (($v3_3)->{'value3'})->{'value0'} instanceof \Test\RBTree\Test_RBTree_R)) {
$__t9 = new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_R(), new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_B(), $v1_1, $v2_2, ($v3_3)->{'value1'}), ($v3_3)->{'value2'}, new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_B(), (($v3_3)->{'value3'})->{'value1'}, (($v3_3)->{'value3'})->{'value2'}, (($v3_3)->{'value3'})->{'value3'}));
goto end_branch_9;;
};
$__t9 = new \Test\RBTree\Test_RBTree_T($v_0, $v1_1, $v2_2, $v3_3);
end_branch_9:;
$__t8 = $__t9;
goto end_branch_8;;
};
if ((($v3_3)->{'value3'} instanceof \Test\RBTree\Test_RBTree_T && (($v3_3)->{'value3'})->{'value0'} instanceof \Test\RBTree\Test_RBTree_R)) {
$__t8 = new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_R(), new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_B(), $v1_1, $v2_2, ($v3_3)->{'value1'}), ($v3_3)->{'value2'}, new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_B(), (($v3_3)->{'value3'})->{'value1'}, (($v3_3)->{'value3'})->{'value2'}, (($v3_3)->{'value3'})->{'value3'}));
goto end_branch_8;;
};
$__t8 = new \Test\RBTree\Test_RBTree_T($v_0, $v1_1, $v2_2, $v3_3);
end_branch_8:;
$__t4 = $__t8;
goto end_branch_4;;
};
$__t4 = new \Test\RBTree\Test_RBTree_T($v_0, $v1_1, $v2_2, $v3_3);
end_branch_4:;
$__t3 = $__t4;
goto end_branch_3;;
};
if (($v1_1)->{'value3'} instanceof \Test\RBTree\Test_RBTree_T) {
$__t10 = null;;
if ((($v1_1)->{'value3'})->{'value0'} instanceof \Test\RBTree\Test_RBTree_R) {
$__t10 = new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_R(), new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_B(), ($v1_1)->{'value1'}, ($v1_1)->{'value2'}, (($v1_1)->{'value3'})->{'value1'}), (($v1_1)->{'value3'})->{'value2'}, new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_B(), (($v1_1)->{'value3'})->{'value3'}, $v2_2, $v3_3));
goto end_branch_10;;
};
if (($v3_3 instanceof \Test\RBTree\Test_RBTree_T && ($v3_3)->{'value0'} instanceof \Test\RBTree\Test_RBTree_R)) {
$__t11 = null;;
if (($v3_3)->{'value1'} instanceof \Test\RBTree\Test_RBTree_T) {
$__t12 = null;;
if ((($v3_3)->{'value1'})->{'value0'} instanceof \Test\RBTree\Test_RBTree_R) {
$__t12 = new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_R(), new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_B(), $v1_1, $v2_2, (($v3_3)->{'value1'})->{'value1'}), (($v3_3)->{'value1'})->{'value2'}, new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_B(), (($v3_3)->{'value1'})->{'value3'}, ($v3_3)->{'value2'}, ($v3_3)->{'value3'}));
goto end_branch_12;;
};
if ((($v3_3)->{'value3'} instanceof \Test\RBTree\Test_RBTree_T && (($v3_3)->{'value3'})->{'value0'} instanceof \Test\RBTree\Test_RBTree_R)) {
$__t12 = new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_R(), new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_B(), $v1_1, $v2_2, ($v3_3)->{'value1'}), ($v3_3)->{'value2'}, new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_B(), (($v3_3)->{'value3'})->{'value1'}, (($v3_3)->{'value3'})->{'value2'}, (($v3_3)->{'value3'})->{'value3'}));
goto end_branch_12;;
};
$__t12 = new \Test\RBTree\Test_RBTree_T($v_0, $v1_1, $v2_2, $v3_3);
end_branch_12:;
$__t11 = $__t12;
goto end_branch_11;;
};
if ((($v3_3)->{'value3'} instanceof \Test\RBTree\Test_RBTree_T && (($v3_3)->{'value3'})->{'value0'} instanceof \Test\RBTree\Test_RBTree_R)) {
$__t11 = new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_R(), new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_B(), $v1_1, $v2_2, ($v3_3)->{'value1'}), ($v3_3)->{'value2'}, new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_B(), (($v3_3)->{'value3'})->{'value1'}, (($v3_3)->{'value3'})->{'value2'}, (($v3_3)->{'value3'})->{'value3'}));
goto end_branch_11;;
};
$__t11 = new \Test\RBTree\Test_RBTree_T($v_0, $v1_1, $v2_2, $v3_3);
end_branch_11:;
$__t10 = $__t11;
goto end_branch_10;;
};
$__t10 = new \Test\RBTree\Test_RBTree_T($v_0, $v1_1, $v2_2, $v3_3);
end_branch_10:;
$__t3 = $__t10;
goto end_branch_3;;
};
if (($v3_3 instanceof \Test\RBTree\Test_RBTree_T && ($v3_3)->{'value0'} instanceof \Test\RBTree\Test_RBTree_R)) {
$__t13 = null;;
if (($v3_3)->{'value1'} instanceof \Test\RBTree\Test_RBTree_T) {
$__t14 = null;;
if ((($v3_3)->{'value1'})->{'value0'} instanceof \Test\RBTree\Test_RBTree_R) {
$__t14 = new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_R(), new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_B(), $v1_1, $v2_2, (($v3_3)->{'value1'})->{'value1'}), (($v3_3)->{'value1'})->{'value2'}, new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_B(), (($v3_3)->{'value1'})->{'value3'}, ($v3_3)->{'value2'}, ($v3_3)->{'value3'}));
goto end_branch_14;;
};
if ((($v3_3)->{'value3'} instanceof \Test\RBTree\Test_RBTree_T && (($v3_3)->{'value3'})->{'value0'} instanceof \Test\RBTree\Test_RBTree_R)) {
$__t14 = new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_R(), new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_B(), $v1_1, $v2_2, ($v3_3)->{'value1'}), ($v3_3)->{'value2'}, new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_B(), (($v3_3)->{'value3'})->{'value1'}, (($v3_3)->{'value3'})->{'value2'}, (($v3_3)->{'value3'})->{'value3'}));
goto end_branch_14;;
};
$__t14 = new \Test\RBTree\Test_RBTree_T($v_0, $v1_1, $v2_2, $v3_3);
end_branch_14:;
$__t13 = $__t14;
goto end_branch_13;;
};
if ((($v3_3)->{'value3'} instanceof \Test\RBTree\Test_RBTree_T && (($v3_3)->{'value3'})->{'value0'} instanceof \Test\RBTree\Test_RBTree_R)) {
$__t13 = new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_R(), new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_B(), $v1_1, $v2_2, ($v3_3)->{'value1'}), ($v3_3)->{'value2'}, new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_B(), (($v3_3)->{'value3'})->{'value1'}, (($v3_3)->{'value3'})->{'value2'}, (($v3_3)->{'value3'})->{'value3'}));
goto end_branch_13;;
};
$__t13 = new \Test\RBTree\Test_RBTree_T($v_0, $v1_1, $v2_2, $v3_3);
end_branch_13:;
$__t3 = $__t13;
goto end_branch_3;;
};
$__t3 = new \Test\RBTree\Test_RBTree_T($v_0, $v1_1, $v2_2, $v3_3);
end_branch_3:;
$__t2 = $__t3;
goto end_branch_2;;
};
if (($v3_3 instanceof \Test\RBTree\Test_RBTree_T && ($v3_3)->{'value0'} instanceof \Test\RBTree\Test_RBTree_R)) {
$__t15 = null;;
if (($v3_3)->{'value1'} instanceof \Test\RBTree\Test_RBTree_T) {
$__t16 = null;;
if ((($v3_3)->{'value1'})->{'value0'} instanceof \Test\RBTree\Test_RBTree_R) {
$__t16 = new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_R(), new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_B(), $v1_1, $v2_2, (($v3_3)->{'value1'})->{'value1'}), (($v3_3)->{'value1'})->{'value2'}, new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_B(), (($v3_3)->{'value1'})->{'value3'}, ($v3_3)->{'value2'}, ($v3_3)->{'value3'}));
goto end_branch_16;;
};
if ((($v3_3)->{'value3'} instanceof \Test\RBTree\Test_RBTree_T && (($v3_3)->{'value3'})->{'value0'} instanceof \Test\RBTree\Test_RBTree_R)) {
$__t16 = new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_R(), new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_B(), $v1_1, $v2_2, ($v3_3)->{'value1'}), ($v3_3)->{'value2'}, new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_B(), (($v3_3)->{'value3'})->{'value1'}, (($v3_3)->{'value3'})->{'value2'}, (($v3_3)->{'value3'})->{'value3'}));
goto end_branch_16;;
};
$__t16 = new \Test\RBTree\Test_RBTree_T($v_0, $v1_1, $v2_2, $v3_3);
end_branch_16:;
$__t15 = $__t16;
goto end_branch_15;;
};
if ((($v3_3)->{'value3'} instanceof \Test\RBTree\Test_RBTree_T && (($v3_3)->{'value3'})->{'value0'} instanceof \Test\RBTree\Test_RBTree_R)) {
$__t15 = new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_R(), new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_B(), $v1_1, $v2_2, ($v3_3)->{'value1'}), ($v3_3)->{'value2'}, new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_B(), (($v3_3)->{'value3'})->{'value1'}, (($v3_3)->{'value3'})->{'value2'}, (($v3_3)->{'value3'})->{'value3'}));
goto end_branch_15;;
};
$__t15 = new \Test\RBTree\Test_RBTree_T($v_0, $v1_1, $v2_2, $v3_3);
end_branch_15:;
$__t2 = $__t15;
goto end_branch_2;;
};
$__t2 = new \Test\RBTree\Test_RBTree_T($v_0, $v1_1, $v2_2, $v3_3);
end_branch_2:;
$__t1 = $__t2;
goto end_branch_1;;
};
if (($v3_3 instanceof \Test\RBTree\Test_RBTree_T && ($v3_3)->{'value0'} instanceof \Test\RBTree\Test_RBTree_R)) {
$__t17 = null;;
if (($v3_3)->{'value1'} instanceof \Test\RBTree\Test_RBTree_T) {
$__t18 = null;;
if ((($v3_3)->{'value1'})->{'value0'} instanceof \Test\RBTree\Test_RBTree_R) {
$__t18 = new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_R(), new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_B(), $v1_1, $v2_2, (($v3_3)->{'value1'})->{'value1'}), (($v3_3)->{'value1'})->{'value2'}, new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_B(), (($v3_3)->{'value1'})->{'value3'}, ($v3_3)->{'value2'}, ($v3_3)->{'value3'}));
goto end_branch_18;;
};
if ((($v3_3)->{'value3'} instanceof \Test\RBTree\Test_RBTree_T && (($v3_3)->{'value3'})->{'value0'} instanceof \Test\RBTree\Test_RBTree_R)) {
$__t18 = new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_R(), new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_B(), $v1_1, $v2_2, ($v3_3)->{'value1'}), ($v3_3)->{'value2'}, new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_B(), (($v3_3)->{'value3'})->{'value1'}, (($v3_3)->{'value3'})->{'value2'}, (($v3_3)->{'value3'})->{'value3'}));
goto end_branch_18;;
};
$__t18 = new \Test\RBTree\Test_RBTree_T($v_0, $v1_1, $v2_2, $v3_3);
end_branch_18:;
$__t17 = $__t18;
goto end_branch_17;;
};
if ((($v3_3)->{'value3'} instanceof \Test\RBTree\Test_RBTree_T && (($v3_3)->{'value3'})->{'value0'} instanceof \Test\RBTree\Test_RBTree_R)) {
$__t17 = new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_R(), new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_B(), $v1_1, $v2_2, ($v3_3)->{'value1'}), ($v3_3)->{'value2'}, new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_B(), (($v3_3)->{'value3'})->{'value1'}, (($v3_3)->{'value3'})->{'value2'}, (($v3_3)->{'value3'})->{'value3'}));
goto end_branch_17;;
};
$__t17 = new \Test\RBTree\Test_RBTree_T($v_0, $v1_1, $v2_2, $v3_3);
end_branch_17:;
$__t1 = $__t17;
goto end_branch_1;;
};
$__t1 = new \Test\RBTree\Test_RBTree_T($v_0, $v1_1, $v2_2, $v3_3);
end_branch_1:;
$__t0 = $__t1;
goto end_branch_0;;
};
  $__t0 = new \Test\RBTree\Test_RBTree_T($v_0, $v1_1, $v2_2, $v3_3);
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 4 < $__num ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
}
$GLOBALS['Test_RBTree_balance'] = __NAMESPACE__ . '\\majTest_majRmajBmajTree_balance';

// Test_RBTree_ins
function majTest_majRmajBmajTree_ins(int $v_0, $v1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majTest_majRmajBmajTree_ins';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_Test_RBTree_ins_v_0 = $v_0;
  $__tco_var_Test_RBTree_ins_v1_1 = $v1_1;
  tco_loop_Test_RBTree_ins:;
  $v_0 = $__tco_var_Test_RBTree_ins_v_0;
  $v1_1 = $__tco_var_Test_RBTree_ins_v1_1;
  $__t0 = null;;
  if ($v1_1 instanceof \Test\RBTree\Test_RBTree_E) {
$__t0 = new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_R(), new \Test\RBTree\Test_RBTree_E(), $v_0, new \Test\RBTree\Test_RBTree_E());
goto end_branch_0;;
};
  if ($v1_1 instanceof \Test\RBTree\Test_RBTree_T) {
$__t1 = null;;
if (($v_0 < ($v1_1)->{'value2'})) {
$__t1 = ($GLOBALS['Test_RBTree_balance'])(($v1_1)->{'value0'}, ($GLOBALS['Test_RBTree_ins'])($v_0, ($v1_1)->{'value1'}), ($v1_1)->{'value2'}, ($v1_1)->{'value3'});
goto end_branch_1;;
};
if (($v_0 > ($v1_1)->{'value2'})) {
$__t1 = ($GLOBALS['Test_RBTree_balance'])(($v1_1)->{'value0'}, ($v1_1)->{'value1'}, ($v1_1)->{'value2'}, ($GLOBALS['Test_RBTree_ins'])($v_0, ($v1_1)->{'value3'}));
goto end_branch_1;;
};
$__t1 = new \Test\RBTree\Test_RBTree_T(($v1_1)->{'value0'}, ($v1_1)->{'value1'}, ($v1_1)->{'value2'}, ($v1_1)->{'value3'});
end_branch_1:;
$__t0 = $__t1;
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Test_RBTree_ins'] = __NAMESPACE__ . '\\majTest_majRmajBmajTree_ins';

// Test_RBTree_insert
function majTest_majRmajBmajTree_insert(int $x_0, $s_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majTest_majRmajBmajTree_insert';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_2_0 = ($GLOBALS['Test_RBTree_ins'])($x_0, $s_1);
  $__t1 = null;;
  if ($__local_var_2_0 instanceof \Test\RBTree\Test_RBTree_T) {
$__t1 = new \Test\RBTree\Test_RBTree_T(new \Test\RBTree\Test_RBTree_B(), ($__local_var_2_0)->{'value1'}, ($__local_var_2_0)->{'value2'}, ($__local_var_2_0)->{'value3'});
goto end_branch_1;;
};
  if ($__local_var_2_0 instanceof \Test\RBTree\Test_RBTree_E) {
$__t1 = new \Test\RBTree\Test_RBTree_E();
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Test_RBTree_insert'] = __NAMESPACE__ . '\\majTest_majRmajBmajTree_insert';

// Test_RBTree_buildTree
function majTest_majRmajBmajTree_buildmajTree(int $v_0, $v1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majTest_majRmajBmajTree_buildmajTree';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_Test_RBTree_buildTree_v_0 = $v_0;
  $__tco_var_Test_RBTree_buildTree_v1_1 = $v1_1;
  tco_loop_Test_RBTree_buildTree:;
  $v_0 = $__tco_var_Test_RBTree_buildTree_v_0;
  $v1_1 = $__tco_var_Test_RBTree_buildTree_v1_1;
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
  $__tco_1 = ($GLOBALS['Test_RBTree_insert'])($v_0, $v1_1);
  $__tco_var_Test_RBTree_buildTree_v_0 = $__tco_0;
  $__tco_var_Test_RBTree_buildTree_v1_1 = $__tco_1;
  goto tco_loop_Test_RBTree_buildTree;;
  $__t2 = null;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Test_RBTree_buildTree'] = __NAMESPACE__ . '\\majTest_majRmajBmajTree_buildmajTree';

// Test_RBTree_act
$GLOBALS['Test_RBTree_act'] = (($GLOBALS['Effect_bindEffect'])['bind'])(($GLOBALS['Bench_opaque'])(100000), function($dummy_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Effect_Console_log'])((($GLOBALS['Data_Show_showInt'])['show'])(($GLOBALS['Test_RBTree_depth'])(($GLOBALS['Test_RBTree_buildTree'])($dummy_0, new \Test\RBTree\Test_RBTree_E()))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});

