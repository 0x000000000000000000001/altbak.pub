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
\PhpursThunks::$thunks['Test_RBTree_R'] = function() { $v = ($GLOBALS['__phpurs_data0_R'] ??= new Phpurs_Data0("R")); return $v; };
\PhpursThunks::$thunks['Test_RBTree_B'] = function() { $v = ($GLOBALS['__phpurs_data0_B'] ??= new Phpurs_Data0("B")); return $v; };
\PhpursThunks::$thunks['Test_RBTree_E'] = function() { $v = ($GLOBALS['__phpurs_data0_E'] ??= new Phpurs_Data0("E")); return $v; };
\PhpursThunks::$thunks['Test_RBTree_T'] = function() { $v = (function() {
  $__fn = function($value0 = null, $value1 = null, $value2 = null, $value3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = new Phpurs_Data4("T", $value0, $value1, $value2, $value3);
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Test_RBTree_max'] = function() { $v = (function() {
  $__fn = function($x_0 = null, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
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
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Test_RBTree_makeBlack'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object($v_0) && (($v_0)->{'tag'} === "T"))) {
$__t0 = new Phpurs_Data4("T", new Phpurs_Data0("B"), ($v_0)->{'value1'}, ($v_0)->{'value2'}, ($v_0)->{'value3'});
goto end_branch_0;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "E"))) {
$__t0 = new Phpurs_Data0("E");
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Test_RBTree_describe'] = function() { $v = (($GLOBALS['Effect_Console_log'] ?? \PhpursThunks::eval('Effect_Console_log')))("Red-Black Tree (100k Worst-Case Insertions):"); return $v; };
\PhpursThunks::$thunks['Test_RBTree_depth'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__tco_var_Test_RBTree_depth_v_0 = $v_0;
  tco_loop_Test_RBTree_depth:;
  $v_0 = $__tco_var_Test_RBTree_depth_v_0;
  $__t0 = null;;
  if ((is_object($v_0) && (($v_0)->{'tag'} === "E"))) {
$__t0 = 0;
goto end_branch_0;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "T"))) {
$__local_var_1_1 = (($GLOBALS['Test_RBTree_depth'] ?? \PhpursThunks::eval('Test_RBTree_depth')))(($v_0)->{'value1'});
$__local_var_2_2 = (($GLOBALS['Test_RBTree_depth'] ?? \PhpursThunks::eval('Test_RBTree_depth')))(($v_0)->{'value3'});
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Test_RBTree_balance'] = function() { $v = (function() {
  $__fn = function($v_0 = null, $v1_1 = null, $v2_2 = null, $v3_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__t0 = null;;
  if ((is_object($v_0) && (($v_0)->{'tag'} === "B"))) {
$__t1 = null;;
if ((is_object($v1_1) && (($v1_1)->{'tag'} === "T"))) {
$__t2 = null;;
if ((is_object(($v1_1)->{'value0'}) && ((($v1_1)->{'value0'})->{'tag'} === "R"))) {
$__t3 = null;;
if ((is_object(($v1_1)->{'value1'}) && ((($v1_1)->{'value1'})->{'tag'} === "T"))) {
$__t4 = null;;
if ((is_object((($v1_1)->{'value1'})->{'value0'}) && (((($v1_1)->{'value1'})->{'value0'})->{'tag'} === "R"))) {
$__t4 = new Phpurs_Data4("T", new Phpurs_Data0("R"), new Phpurs_Data4("T", new Phpurs_Data0("B"), (($v1_1)->{'value1'})->{'value1'}, (($v1_1)->{'value1'})->{'value2'}, (($v1_1)->{'value1'})->{'value3'}), ($v1_1)->{'value2'}, new Phpurs_Data4("T", new Phpurs_Data0("B"), ($v1_1)->{'value3'}, $v2_2, $v3_3));
goto end_branch_4;;
};
if ((is_object(($v1_1)->{'value3'}) && ((($v1_1)->{'value3'})->{'tag'} === "T"))) {
$__t5 = null;;
if ((is_object((($v1_1)->{'value3'})->{'value0'}) && (((($v1_1)->{'value3'})->{'value0'})->{'tag'} === "R"))) {
$__t5 = new Phpurs_Data4("T", new Phpurs_Data0("R"), new Phpurs_Data4("T", new Phpurs_Data0("B"), ($v1_1)->{'value1'}, ($v1_1)->{'value2'}, (($v1_1)->{'value3'})->{'value1'}), (($v1_1)->{'value3'})->{'value2'}, new Phpurs_Data4("T", new Phpurs_Data0("B"), (($v1_1)->{'value3'})->{'value3'}, $v2_2, $v3_3));
goto end_branch_5;;
};
if (((is_object($v3_3) && (($v3_3)->{'tag'} === "T")) && (is_object(($v3_3)->{'value0'}) && ((($v3_3)->{'value0'})->{'tag'} === "R")))) {
$__t6 = null;;
if ((is_object(($v3_3)->{'value1'}) && ((($v3_3)->{'value1'})->{'tag'} === "T"))) {
$__t7 = null;;
if ((is_object((($v3_3)->{'value1'})->{'value0'}) && (((($v3_3)->{'value1'})->{'value0'})->{'tag'} === "R"))) {
$__t7 = new Phpurs_Data4("T", new Phpurs_Data0("R"), new Phpurs_Data4("T", new Phpurs_Data0("B"), $v1_1, $v2_2, (($v3_3)->{'value1'})->{'value1'}), (($v3_3)->{'value1'})->{'value2'}, new Phpurs_Data4("T", new Phpurs_Data0("B"), (($v3_3)->{'value1'})->{'value3'}, ($v3_3)->{'value2'}, ($v3_3)->{'value3'}));
goto end_branch_7;;
};
if (((is_object(($v3_3)->{'value3'}) && ((($v3_3)->{'value3'})->{'tag'} === "T")) && (is_object((($v3_3)->{'value3'})->{'value0'}) && (((($v3_3)->{'value3'})->{'value0'})->{'tag'} === "R")))) {
$__t7 = new Phpurs_Data4("T", new Phpurs_Data0("R"), new Phpurs_Data4("T", new Phpurs_Data0("B"), $v1_1, $v2_2, ($v3_3)->{'value1'}), ($v3_3)->{'value2'}, new Phpurs_Data4("T", new Phpurs_Data0("B"), (($v3_3)->{'value3'})->{'value1'}, (($v3_3)->{'value3'})->{'value2'}, (($v3_3)->{'value3'})->{'value3'}));
goto end_branch_7;;
};
$__t7 = new Phpurs_Data4("T", $v_0, $v1_1, $v2_2, $v3_3);
end_branch_7:;
$__t6 = $__t7;
goto end_branch_6;;
};
if (((is_object(($v3_3)->{'value3'}) && ((($v3_3)->{'value3'})->{'tag'} === "T")) && (is_object((($v3_3)->{'value3'})->{'value0'}) && (((($v3_3)->{'value3'})->{'value0'})->{'tag'} === "R")))) {
$__t6 = new Phpurs_Data4("T", new Phpurs_Data0("R"), new Phpurs_Data4("T", new Phpurs_Data0("B"), $v1_1, $v2_2, ($v3_3)->{'value1'}), ($v3_3)->{'value2'}, new Phpurs_Data4("T", new Phpurs_Data0("B"), (($v3_3)->{'value3'})->{'value1'}, (($v3_3)->{'value3'})->{'value2'}, (($v3_3)->{'value3'})->{'value3'}));
goto end_branch_6;;
};
$__t6 = new Phpurs_Data4("T", $v_0, $v1_1, $v2_2, $v3_3);
end_branch_6:;
$__t5 = $__t6;
goto end_branch_5;;
};
$__t5 = new Phpurs_Data4("T", $v_0, $v1_1, $v2_2, $v3_3);
end_branch_5:;
$__t4 = $__t5;
goto end_branch_4;;
};
if (((is_object($v3_3) && (($v3_3)->{'tag'} === "T")) && (is_object(($v3_3)->{'value0'}) && ((($v3_3)->{'value0'})->{'tag'} === "R")))) {
$__t8 = null;;
if ((is_object(($v3_3)->{'value1'}) && ((($v3_3)->{'value1'})->{'tag'} === "T"))) {
$__t9 = null;;
if ((is_object((($v3_3)->{'value1'})->{'value0'}) && (((($v3_3)->{'value1'})->{'value0'})->{'tag'} === "R"))) {
$__t9 = new Phpurs_Data4("T", new Phpurs_Data0("R"), new Phpurs_Data4("T", new Phpurs_Data0("B"), $v1_1, $v2_2, (($v3_3)->{'value1'})->{'value1'}), (($v3_3)->{'value1'})->{'value2'}, new Phpurs_Data4("T", new Phpurs_Data0("B"), (($v3_3)->{'value1'})->{'value3'}, ($v3_3)->{'value2'}, ($v3_3)->{'value3'}));
goto end_branch_9;;
};
if (((is_object(($v3_3)->{'value3'}) && ((($v3_3)->{'value3'})->{'tag'} === "T")) && (is_object((($v3_3)->{'value3'})->{'value0'}) && (((($v3_3)->{'value3'})->{'value0'})->{'tag'} === "R")))) {
$__t9 = new Phpurs_Data4("T", new Phpurs_Data0("R"), new Phpurs_Data4("T", new Phpurs_Data0("B"), $v1_1, $v2_2, ($v3_3)->{'value1'}), ($v3_3)->{'value2'}, new Phpurs_Data4("T", new Phpurs_Data0("B"), (($v3_3)->{'value3'})->{'value1'}, (($v3_3)->{'value3'})->{'value2'}, (($v3_3)->{'value3'})->{'value3'}));
goto end_branch_9;;
};
$__t9 = new Phpurs_Data4("T", $v_0, $v1_1, $v2_2, $v3_3);
end_branch_9:;
$__t8 = $__t9;
goto end_branch_8;;
};
if (((is_object(($v3_3)->{'value3'}) && ((($v3_3)->{'value3'})->{'tag'} === "T")) && (is_object((($v3_3)->{'value3'})->{'value0'}) && (((($v3_3)->{'value3'})->{'value0'})->{'tag'} === "R")))) {
$__t8 = new Phpurs_Data4("T", new Phpurs_Data0("R"), new Phpurs_Data4("T", new Phpurs_Data0("B"), $v1_1, $v2_2, ($v3_3)->{'value1'}), ($v3_3)->{'value2'}, new Phpurs_Data4("T", new Phpurs_Data0("B"), (($v3_3)->{'value3'})->{'value1'}, (($v3_3)->{'value3'})->{'value2'}, (($v3_3)->{'value3'})->{'value3'}));
goto end_branch_8;;
};
$__t8 = new Phpurs_Data4("T", $v_0, $v1_1, $v2_2, $v3_3);
end_branch_8:;
$__t4 = $__t8;
goto end_branch_4;;
};
$__t4 = new Phpurs_Data4("T", $v_0, $v1_1, $v2_2, $v3_3);
end_branch_4:;
$__t3 = $__t4;
goto end_branch_3;;
};
if ((is_object(($v1_1)->{'value3'}) && ((($v1_1)->{'value3'})->{'tag'} === "T"))) {
$__t10 = null;;
if ((is_object((($v1_1)->{'value3'})->{'value0'}) && (((($v1_1)->{'value3'})->{'value0'})->{'tag'} === "R"))) {
$__t10 = new Phpurs_Data4("T", new Phpurs_Data0("R"), new Phpurs_Data4("T", new Phpurs_Data0("B"), ($v1_1)->{'value1'}, ($v1_1)->{'value2'}, (($v1_1)->{'value3'})->{'value1'}), (($v1_1)->{'value3'})->{'value2'}, new Phpurs_Data4("T", new Phpurs_Data0("B"), (($v1_1)->{'value3'})->{'value3'}, $v2_2, $v3_3));
goto end_branch_10;;
};
if (((is_object($v3_3) && (($v3_3)->{'tag'} === "T")) && (is_object(($v3_3)->{'value0'}) && ((($v3_3)->{'value0'})->{'tag'} === "R")))) {
$__t11 = null;;
if ((is_object(($v3_3)->{'value1'}) && ((($v3_3)->{'value1'})->{'tag'} === "T"))) {
$__t12 = null;;
if ((is_object((($v3_3)->{'value1'})->{'value0'}) && (((($v3_3)->{'value1'})->{'value0'})->{'tag'} === "R"))) {
$__t12 = new Phpurs_Data4("T", new Phpurs_Data0("R"), new Phpurs_Data4("T", new Phpurs_Data0("B"), $v1_1, $v2_2, (($v3_3)->{'value1'})->{'value1'}), (($v3_3)->{'value1'})->{'value2'}, new Phpurs_Data4("T", new Phpurs_Data0("B"), (($v3_3)->{'value1'})->{'value3'}, ($v3_3)->{'value2'}, ($v3_3)->{'value3'}));
goto end_branch_12;;
};
if (((is_object(($v3_3)->{'value3'}) && ((($v3_3)->{'value3'})->{'tag'} === "T")) && (is_object((($v3_3)->{'value3'})->{'value0'}) && (((($v3_3)->{'value3'})->{'value0'})->{'tag'} === "R")))) {
$__t12 = new Phpurs_Data4("T", new Phpurs_Data0("R"), new Phpurs_Data4("T", new Phpurs_Data0("B"), $v1_1, $v2_2, ($v3_3)->{'value1'}), ($v3_3)->{'value2'}, new Phpurs_Data4("T", new Phpurs_Data0("B"), (($v3_3)->{'value3'})->{'value1'}, (($v3_3)->{'value3'})->{'value2'}, (($v3_3)->{'value3'})->{'value3'}));
goto end_branch_12;;
};
$__t12 = new Phpurs_Data4("T", $v_0, $v1_1, $v2_2, $v3_3);
end_branch_12:;
$__t11 = $__t12;
goto end_branch_11;;
};
if (((is_object(($v3_3)->{'value3'}) && ((($v3_3)->{'value3'})->{'tag'} === "T")) && (is_object((($v3_3)->{'value3'})->{'value0'}) && (((($v3_3)->{'value3'})->{'value0'})->{'tag'} === "R")))) {
$__t11 = new Phpurs_Data4("T", new Phpurs_Data0("R"), new Phpurs_Data4("T", new Phpurs_Data0("B"), $v1_1, $v2_2, ($v3_3)->{'value1'}), ($v3_3)->{'value2'}, new Phpurs_Data4("T", new Phpurs_Data0("B"), (($v3_3)->{'value3'})->{'value1'}, (($v3_3)->{'value3'})->{'value2'}, (($v3_3)->{'value3'})->{'value3'}));
goto end_branch_11;;
};
$__t11 = new Phpurs_Data4("T", $v_0, $v1_1, $v2_2, $v3_3);
end_branch_11:;
$__t10 = $__t11;
goto end_branch_10;;
};
$__t10 = new Phpurs_Data4("T", $v_0, $v1_1, $v2_2, $v3_3);
end_branch_10:;
$__t3 = $__t10;
goto end_branch_3;;
};
if (((is_object($v3_3) && (($v3_3)->{'tag'} === "T")) && (is_object(($v3_3)->{'value0'}) && ((($v3_3)->{'value0'})->{'tag'} === "R")))) {
$__t13 = null;;
if ((is_object(($v3_3)->{'value1'}) && ((($v3_3)->{'value1'})->{'tag'} === "T"))) {
$__t14 = null;;
if ((is_object((($v3_3)->{'value1'})->{'value0'}) && (((($v3_3)->{'value1'})->{'value0'})->{'tag'} === "R"))) {
$__t14 = new Phpurs_Data4("T", new Phpurs_Data0("R"), new Phpurs_Data4("T", new Phpurs_Data0("B"), $v1_1, $v2_2, (($v3_3)->{'value1'})->{'value1'}), (($v3_3)->{'value1'})->{'value2'}, new Phpurs_Data4("T", new Phpurs_Data0("B"), (($v3_3)->{'value1'})->{'value3'}, ($v3_3)->{'value2'}, ($v3_3)->{'value3'}));
goto end_branch_14;;
};
if (((is_object(($v3_3)->{'value3'}) && ((($v3_3)->{'value3'})->{'tag'} === "T")) && (is_object((($v3_3)->{'value3'})->{'value0'}) && (((($v3_3)->{'value3'})->{'value0'})->{'tag'} === "R")))) {
$__t14 = new Phpurs_Data4("T", new Phpurs_Data0("R"), new Phpurs_Data4("T", new Phpurs_Data0("B"), $v1_1, $v2_2, ($v3_3)->{'value1'}), ($v3_3)->{'value2'}, new Phpurs_Data4("T", new Phpurs_Data0("B"), (($v3_3)->{'value3'})->{'value1'}, (($v3_3)->{'value3'})->{'value2'}, (($v3_3)->{'value3'})->{'value3'}));
goto end_branch_14;;
};
$__t14 = new Phpurs_Data4("T", $v_0, $v1_1, $v2_2, $v3_3);
end_branch_14:;
$__t13 = $__t14;
goto end_branch_13;;
};
if (((is_object(($v3_3)->{'value3'}) && ((($v3_3)->{'value3'})->{'tag'} === "T")) && (is_object((($v3_3)->{'value3'})->{'value0'}) && (((($v3_3)->{'value3'})->{'value0'})->{'tag'} === "R")))) {
$__t13 = new Phpurs_Data4("T", new Phpurs_Data0("R"), new Phpurs_Data4("T", new Phpurs_Data0("B"), $v1_1, $v2_2, ($v3_3)->{'value1'}), ($v3_3)->{'value2'}, new Phpurs_Data4("T", new Phpurs_Data0("B"), (($v3_3)->{'value3'})->{'value1'}, (($v3_3)->{'value3'})->{'value2'}, (($v3_3)->{'value3'})->{'value3'}));
goto end_branch_13;;
};
$__t13 = new Phpurs_Data4("T", $v_0, $v1_1, $v2_2, $v3_3);
end_branch_13:;
$__t3 = $__t13;
goto end_branch_3;;
};
$__t3 = new Phpurs_Data4("T", $v_0, $v1_1, $v2_2, $v3_3);
end_branch_3:;
$__t2 = $__t3;
goto end_branch_2;;
};
if (((is_object($v3_3) && (($v3_3)->{'tag'} === "T")) && (is_object(($v3_3)->{'value0'}) && ((($v3_3)->{'value0'})->{'tag'} === "R")))) {
$__t15 = null;;
if ((is_object(($v3_3)->{'value1'}) && ((($v3_3)->{'value1'})->{'tag'} === "T"))) {
$__t16 = null;;
if ((is_object((($v3_3)->{'value1'})->{'value0'}) && (((($v3_3)->{'value1'})->{'value0'})->{'tag'} === "R"))) {
$__t16 = new Phpurs_Data4("T", new Phpurs_Data0("R"), new Phpurs_Data4("T", new Phpurs_Data0("B"), $v1_1, $v2_2, (($v3_3)->{'value1'})->{'value1'}), (($v3_3)->{'value1'})->{'value2'}, new Phpurs_Data4("T", new Phpurs_Data0("B"), (($v3_3)->{'value1'})->{'value3'}, ($v3_3)->{'value2'}, ($v3_3)->{'value3'}));
goto end_branch_16;;
};
if (((is_object(($v3_3)->{'value3'}) && ((($v3_3)->{'value3'})->{'tag'} === "T")) && (is_object((($v3_3)->{'value3'})->{'value0'}) && (((($v3_3)->{'value3'})->{'value0'})->{'tag'} === "R")))) {
$__t16 = new Phpurs_Data4("T", new Phpurs_Data0("R"), new Phpurs_Data4("T", new Phpurs_Data0("B"), $v1_1, $v2_2, ($v3_3)->{'value1'}), ($v3_3)->{'value2'}, new Phpurs_Data4("T", new Phpurs_Data0("B"), (($v3_3)->{'value3'})->{'value1'}, (($v3_3)->{'value3'})->{'value2'}, (($v3_3)->{'value3'})->{'value3'}));
goto end_branch_16;;
};
$__t16 = new Phpurs_Data4("T", $v_0, $v1_1, $v2_2, $v3_3);
end_branch_16:;
$__t15 = $__t16;
goto end_branch_15;;
};
if (((is_object(($v3_3)->{'value3'}) && ((($v3_3)->{'value3'})->{'tag'} === "T")) && (is_object((($v3_3)->{'value3'})->{'value0'}) && (((($v3_3)->{'value3'})->{'value0'})->{'tag'} === "R")))) {
$__t15 = new Phpurs_Data4("T", new Phpurs_Data0("R"), new Phpurs_Data4("T", new Phpurs_Data0("B"), $v1_1, $v2_2, ($v3_3)->{'value1'}), ($v3_3)->{'value2'}, new Phpurs_Data4("T", new Phpurs_Data0("B"), (($v3_3)->{'value3'})->{'value1'}, (($v3_3)->{'value3'})->{'value2'}, (($v3_3)->{'value3'})->{'value3'}));
goto end_branch_15;;
};
$__t15 = new Phpurs_Data4("T", $v_0, $v1_1, $v2_2, $v3_3);
end_branch_15:;
$__t2 = $__t15;
goto end_branch_2;;
};
$__t2 = new Phpurs_Data4("T", $v_0, $v1_1, $v2_2, $v3_3);
end_branch_2:;
$__t1 = $__t2;
goto end_branch_1;;
};
if (((is_object($v3_3) && (($v3_3)->{'tag'} === "T")) && (is_object(($v3_3)->{'value0'}) && ((($v3_3)->{'value0'})->{'tag'} === "R")))) {
$__t17 = null;;
if ((is_object(($v3_3)->{'value1'}) && ((($v3_3)->{'value1'})->{'tag'} === "T"))) {
$__t18 = null;;
if ((is_object((($v3_3)->{'value1'})->{'value0'}) && (((($v3_3)->{'value1'})->{'value0'})->{'tag'} === "R"))) {
$__t18 = new Phpurs_Data4("T", new Phpurs_Data0("R"), new Phpurs_Data4("T", new Phpurs_Data0("B"), $v1_1, $v2_2, (($v3_3)->{'value1'})->{'value1'}), (($v3_3)->{'value1'})->{'value2'}, new Phpurs_Data4("T", new Phpurs_Data0("B"), (($v3_3)->{'value1'})->{'value3'}, ($v3_3)->{'value2'}, ($v3_3)->{'value3'}));
goto end_branch_18;;
};
if (((is_object(($v3_3)->{'value3'}) && ((($v3_3)->{'value3'})->{'tag'} === "T")) && (is_object((($v3_3)->{'value3'})->{'value0'}) && (((($v3_3)->{'value3'})->{'value0'})->{'tag'} === "R")))) {
$__t18 = new Phpurs_Data4("T", new Phpurs_Data0("R"), new Phpurs_Data4("T", new Phpurs_Data0("B"), $v1_1, $v2_2, ($v3_3)->{'value1'}), ($v3_3)->{'value2'}, new Phpurs_Data4("T", new Phpurs_Data0("B"), (($v3_3)->{'value3'})->{'value1'}, (($v3_3)->{'value3'})->{'value2'}, (($v3_3)->{'value3'})->{'value3'}));
goto end_branch_18;;
};
$__t18 = new Phpurs_Data4("T", $v_0, $v1_1, $v2_2, $v3_3);
end_branch_18:;
$__t17 = $__t18;
goto end_branch_17;;
};
if (((is_object(($v3_3)->{'value3'}) && ((($v3_3)->{'value3'})->{'tag'} === "T")) && (is_object((($v3_3)->{'value3'})->{'value0'}) && (((($v3_3)->{'value3'})->{'value0'})->{'tag'} === "R")))) {
$__t17 = new Phpurs_Data4("T", new Phpurs_Data0("R"), new Phpurs_Data4("T", new Phpurs_Data0("B"), $v1_1, $v2_2, ($v3_3)->{'value1'}), ($v3_3)->{'value2'}, new Phpurs_Data4("T", new Phpurs_Data0("B"), (($v3_3)->{'value3'})->{'value1'}, (($v3_3)->{'value3'})->{'value2'}, (($v3_3)->{'value3'})->{'value3'}));
goto end_branch_17;;
};
$__t17 = new Phpurs_Data4("T", $v_0, $v1_1, $v2_2, $v3_3);
end_branch_17:;
$__t1 = $__t17;
goto end_branch_1;;
};
$__t1 = new Phpurs_Data4("T", $v_0, $v1_1, $v2_2, $v3_3);
end_branch_1:;
$__t0 = $__t1;
goto end_branch_0;;
};
  $__t0 = new Phpurs_Data4("T", $v_0, $v1_1, $v2_2, $v3_3);
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Test_RBTree_ins'] = function() { $v = (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_Test_RBTree_ins_v_0 = $v_0;
  $__tco_var_Test_RBTree_ins_v1_1 = $v1_1;
  tco_loop_Test_RBTree_ins:;
  $v_0 = $__tco_var_Test_RBTree_ins_v_0;
  $v1_1 = $__tco_var_Test_RBTree_ins_v1_1;
  $__t0 = null;;
  if ((is_object($v1_1) && (($v1_1)->{'tag'} === "E"))) {
$__t0 = new Phpurs_Data4("T", new Phpurs_Data0("R"), new Phpurs_Data0("E"), $v_0, new Phpurs_Data0("E"));
goto end_branch_0;;
};
  if ((is_object($v1_1) && (($v1_1)->{'tag'} === "T"))) {
$__t1 = null;;
if (($v_0 < ($v1_1)->{'value2'})) {
$__t1 = ((((($GLOBALS['Test_RBTree_balance'] ?? \PhpursThunks::eval('Test_RBTree_balance')))(($v1_1)->{'value0'}))(((($GLOBALS['Test_RBTree_ins'] ?? \PhpursThunks::eval('Test_RBTree_ins')))($v_0))(($v1_1)->{'value1'})))(($v1_1)->{'value2'}))(($v1_1)->{'value3'});
goto end_branch_1;;
};
if (($v_0 > ($v1_1)->{'value2'})) {
$__t1 = ((((($GLOBALS['Test_RBTree_balance'] ?? \PhpursThunks::eval('Test_RBTree_balance')))(($v1_1)->{'value0'}))(($v1_1)->{'value1'}))(($v1_1)->{'value2'}))(((($GLOBALS['Test_RBTree_ins'] ?? \PhpursThunks::eval('Test_RBTree_ins')))($v_0))(($v1_1)->{'value3'}));
goto end_branch_1;;
};
$__t1 = new Phpurs_Data4("T", ($v1_1)->{'value0'}, ($v1_1)->{'value1'}, ($v1_1)->{'value2'}, ($v1_1)->{'value3'});
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
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Test_RBTree_insert'] = function() { $v = (function() {
  $__fn = function($x_0 = null, $s_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_2_0 = ((($GLOBALS['Test_RBTree_ins'] ?? \PhpursThunks::eval('Test_RBTree_ins')))($x_0))($s_1);
  $__t1 = null;;
  if ((is_object($__local_var_2_0) && (($__local_var_2_0)->{'tag'} === "T"))) {
$__t1 = new Phpurs_Data4("T", new Phpurs_Data0("B"), ($__local_var_2_0)->{'value1'}, ($__local_var_2_0)->{'value2'}, ($__local_var_2_0)->{'value3'});
goto end_branch_1;;
};
  if ((is_object($__local_var_2_0) && (($__local_var_2_0)->{'tag'} === "E"))) {
$__t1 = new Phpurs_Data0("E");
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Test_RBTree_buildTree'] = function() { $v = (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
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
  $__tco_1 = ((($GLOBALS['Test_RBTree_insert'] ?? \PhpursThunks::eval('Test_RBTree_insert')))($v_0))($v1_1);
  $__tco_var_Test_RBTree_buildTree_v_0 = $__tco_0;
  $__tco_var_Test_RBTree_buildTree_v1_1 = $__tco_1;
  goto tco_loop_Test_RBTree_buildTree;;
  $__t2 = null;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Test_RBTree_act'] = function() { $v = ((($GLOBALS['Effect_bindE'] ?? \PhpursThunks::eval('Effect_bindE')))((($GLOBALS['Bench_opaque'] ?? \PhpursThunks::eval('Bench_opaque')))(100000)))(function($dummy_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Effect_Console_log'] ?? \PhpursThunks::eval('Effect_Console_log')))((($GLOBALS['Data_Show_showIntImpl'] ?? \PhpursThunks::eval('Data_Show_showIntImpl')))((($GLOBALS['Test_RBTree_depth'] ?? \PhpursThunks::eval('Test_RBTree_depth')))(((($GLOBALS['Test_RBTree_buildTree'] ?? \PhpursThunks::eval('Test_RBTree_buildTree')))($dummy_0))(new Phpurs_Data0("E")))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}); return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };















