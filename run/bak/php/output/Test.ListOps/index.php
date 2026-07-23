<?php

namespace Test\ListOps;

// ALL IMPORTS: Bench, Control.Bind, Data.Eq, Data.EuclideanRing, Data.Function, Data.Ord, Data.Ring, Data.Semiring, Data.Show, Effect, Effect.Console, Prelude, Prim, Test.ListOps
// TO REQUIRE: Bench, Control.Bind, Data.Eq, Data.EuclideanRing, Data.Function, Data.Ord, Data.Ring, Data.Semiring, Data.Show, Effect, Effect.Console, Prelude, Test.ListOps
require_once __DIR__ . '/../Bench/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.EuclideanRing/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ring/index.php';
require_once __DIR__ . '/../Data.Semiring/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
require_once __DIR__ . '/../Effect/index.php';
require_once __DIR__ . '/../Effect.Console/index.php';
require_once __DIR__ . '/../Prelude/index.php';
require_once __DIR__ . '/../Test.ListOps/index.php';

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
\PhpursThunks::$thunks['Test_ListOps_Nil'] = function() { $v = ($GLOBALS['__phpurs_data0_Nil'] ??= new Phpurs_Data0("Nil")); return $v; };
\PhpursThunks::$thunks['Test_ListOps_Cons'] = function() { $v = (function() {
  $__fn = function($value0 = null, $value1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("Cons", $value0, $value1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Test_ListOps_range'] = function() { $v = (function() {
  $__fn = function($start_0 = null, $end_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go_2_0 = null;
  $go_2_0 = (function() use (&$go_2_0, $start_0) {
  $__fn = function($curr_3 = null, $acc_4 = null) use (&$go_2_0, $start_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go_2_0_0_curr_3 = $curr_3;
  $__tco_var_go_2_0_0_acc_4 = $acc_4;
  $__tco_done_go_2_0_0 = false;
  $__tco_res_go_2_0_0 = null;
  $__tco_loop_go_2_0_0 = (function() use (&$__tco_done_go_2_0_0, &$__tco_var_go_2_0_0_curr_3, &$__tco_var_go_2_0_0_acc_4, &$go_2_0, $start_0) {
  $__fn = function($curr_3 = null, $acc_4 = null) use (&$__tco_done_go_2_0_0, &$__tco_var_go_2_0_0_curr_3, &$__tco_var_go_2_0_0_acc_4, &$go_2_0, $start_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_done_go_2_0_0 = true;
  $curr_3 = $__tco_var_go_2_0_0_curr_3;
  $acc_4 = $__tco_var_go_2_0_0_acc_4;
  $__t2 = null;;
  if ((is_object((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->{'compare'})($curr_3))($start_0)) && (((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->{'compare'})($curr_3))($start_0))->{'tag'} === "LT"))) {
$__t2 = $acc_4;
goto end_branch_2;;
};
  $__tco_0 = ((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))($curr_3))(1);
  $__tco_1 = new Phpurs_Data2("Cons", $curr_3, $acc_4);
  $__tco_var_go_2_0_0_curr_3 = $__tco_0;
  $__tco_var_go_2_0_0_acc_4 = $__tco_1;
  $__tco_done_go_2_0_0 = false;
  $__res = null;
  goto __end;;
  $__t2 = null;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  while (($__tco_done_go_2_0_0 === false)) {
$__tco_res_go_2_0_0 = ($__tco_loop_go_2_0_0)($__tco_var_go_2_0_0_curr_3, $__tco_var_go_2_0_0_acc_4);
};
  $__res = $__tco_res_go_2_0_0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = (($go_2_0)($end_1))(new Phpurs_Data0("Nil"));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Test_ListOps_foldl'] = function() { $v = (function() {
  $__fn = function($v_0 = null, $v1_1 = null, $v2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_Test_ListOps_foldl_v_0 = $v_0;
  $__tco_var_Test_ListOps_foldl_v1_1 = $v1_1;
  $__tco_var_Test_ListOps_foldl_v2_2 = $v2_2;
  $__tco_done_Test_ListOps_foldl = false;
  $__tco_res_Test_ListOps_foldl = null;
  $__tco_loop_Test_ListOps_foldl = (function() use (&$__tco_done_Test_ListOps_foldl, &$__tco_var_Test_ListOps_foldl_v_0, &$__tco_var_Test_ListOps_foldl_v1_1, &$__tco_var_Test_ListOps_foldl_v2_2) {
  $__fn = function($v_0 = null, $v1_1 = null, $v2_2 = null) use (&$__tco_done_Test_ListOps_foldl, &$__tco_var_Test_ListOps_foldl_v_0, &$__tco_var_Test_ListOps_foldl_v1_1, &$__tco_var_Test_ListOps_foldl_v2_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_done_Test_ListOps_foldl = true;
  $v_0 = $__tco_var_Test_ListOps_foldl_v_0;
  $v1_1 = $__tco_var_Test_ListOps_foldl_v1_1;
  $v2_2 = $__tco_var_Test_ListOps_foldl_v2_2;
  $__t0 = null;;
  if ((is_object($v2_2) && (($v2_2)->{'tag'} === "Nil"))) {
$__t0 = $v1_1;
goto end_branch_0;;
};
  if ((is_object($v2_2) && (($v2_2)->{'tag'} === "Cons"))) {
$__tco_1 = $v_0;
$__tco_2 = (($v_0)($v1_1))(($v2_2)->{'value0'});
$__tco_3 = ($v2_2)->{'value1'};
$__tco_var_Test_ListOps_foldl_v_0 = $__tco_1;
$__tco_var_Test_ListOps_foldl_v1_1 = $__tco_2;
$__tco_var_Test_ListOps_foldl_v2_2 = $__tco_3;
$__tco_done_Test_ListOps_foldl = false;
$__res = null;
goto __end;;
$__t0 = null;
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  while (($__tco_done_Test_ListOps_foldl === false)) {
$__tco_res_Test_ListOps_foldl = ($__tco_loop_Test_ListOps_foldl)($__tco_var_Test_ListOps_foldl_v_0, $__tco_var_Test_ListOps_foldl_v1_1, $__tco_var_Test_ListOps_foldl_v2_2);
};
  $__res = $__tco_res_Test_ListOps_foldl;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Test_ListOps_filterEvens'] = function() { $v = function($lst_0 = null) {
  $__num = \func_num_args();
  $go_1_0 = null;
  $go_1_0 = (function() use (&$go_1_0) {
  $__fn = function($v_2 = null, $v1_3 = null) use (&$go_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go_1_0_0_v_2 = $v_2;
  $__tco_var_go_1_0_0_v1_3 = $v1_3;
  $__tco_done_go_1_0_0 = false;
  $__tco_res_go_1_0_0 = null;
  $__tco_loop_go_1_0_0 = (function() use (&$__tco_done_go_1_0_0, &$__tco_var_go_1_0_0_v_2, &$__tco_var_go_1_0_0_v1_3, &$go_1_0) {
  $__fn = function($v_2 = null, $v1_3 = null) use (&$__tco_done_go_1_0_0, &$__tco_var_go_1_0_0_v_2, &$__tco_var_go_1_0_0_v1_3, &$go_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_done_go_1_0_0 = true;
  $v_2 = $__tco_var_go_1_0_0_v_2;
  $v1_3 = $__tco_var_go_1_0_0_v1_3;
  $__t0 = null;;
  if ((is_object($v_2) && (($v_2)->{'tag'} === "Nil"))) {
$__t0 = $v1_3;
goto end_branch_0;;
};
  if ((is_object($v_2) && (($v_2)->{'tag'} === "Cons"))) {
$__t3 = null;;
if (((($GLOBALS['Data_Eq_eqIntImpl'] ?? \PhpursThunks::eval('Data_Eq_eqIntImpl')))(((($GLOBALS['Data_EuclideanRing_intMod'] ?? \PhpursThunks::eval('Data_EuclideanRing_intMod')))(($v_2)->{'value0'}))(2)))(0)) {
$__tco_4 = ($v_2)->{'value1'};
$__tco_5 = new Phpurs_Data2("Cons", ($v_2)->{'value0'}, $v1_3);
$__tco_var_go_1_0_0_v_2 = $__tco_4;
$__tco_var_go_1_0_0_v1_3 = $__tco_5;
$__tco_done_go_1_0_0 = false;
$__res = null;
goto __end;;
$__t3 = null;
goto end_branch_3;;
};
$__tco_1 = ($v_2)->{'value1'};
$__tco_2 = $v1_3;
$__tco_var_go_1_0_0_v_2 = $__tco_1;
$__tco_var_go_1_0_0_v1_3 = $__tco_2;
$__tco_done_go_1_0_0 = false;
$__res = null;
goto __end;;
$__t3 = null;
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
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  while (($__tco_done_go_1_0_0 === false)) {
$__tco_res_go_1_0_0 = ($__tco_loop_go_1_0_0)($__tco_var_go_1_0_0_v_2, $__tco_var_go_1_0_0_v1_3);
};
  $__res = $__tco_res_go_1_0_0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = (($go_1_0)($lst_0))(new Phpurs_Data0("Nil"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Test_ListOps_sumEvens'] = function() { $v = function($n_0 = null) {
  $__num = \func_num_args();
  $__res = (((($GLOBALS['Test_ListOps_foldl'] ?? \PhpursThunks::eval('Test_ListOps_foldl')))(($GLOBALS['Data_Semiring_intAdd'] ?? \PhpursThunks::eval('Data_Semiring_intAdd'))))(0))((($GLOBALS['Test_ListOps_filterEvens'] ?? \PhpursThunks::eval('Test_ListOps_filterEvens')))(((($GLOBALS['Test_ListOps_range'] ?? \PhpursThunks::eval('Test_ListOps_range')))(1))($n_0)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Test_ListOps_describe'] = function() { $v = (($GLOBALS['Effect_Console_log'] ?? \PhpursThunks::eval('Effect_Console_log')))("List Processing (900 elements):"); return $v; };
\PhpursThunks::$thunks['Test_ListOps_act'] = function() { $v = ((($GLOBALS['Effect_bindE'] ?? \PhpursThunks::eval('Effect_bindE')))((($GLOBALS['Bench_opaque'] ?? \PhpursThunks::eval('Bench_opaque')))(900)))(function($dummy_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Effect_Console_log'] ?? \PhpursThunks::eval('Effect_Console_log')))((($GLOBALS['Data_Show_showIntImpl'] ?? \PhpursThunks::eval('Data_Show_showIntImpl')))((((($GLOBALS['Test_ListOps_foldl'] ?? \PhpursThunks::eval('Test_ListOps_foldl')))(($GLOBALS['Data_Semiring_intAdd'] ?? \PhpursThunks::eval('Data_Semiring_intAdd'))))(0))((($GLOBALS['Test_ListOps_filterEvens'] ?? \PhpursThunks::eval('Test_ListOps_filterEvens')))(((($GLOBALS['Test_ListOps_range'] ?? \PhpursThunks::eval('Test_ListOps_range')))(1))($dummy_0)))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}); return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };










