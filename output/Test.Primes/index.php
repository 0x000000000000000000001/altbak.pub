<?php

namespace Test\Primes;

// ALL IMPORTS: Bench, Control.Bind, Data.Eq, Data.EuclideanRing, Data.Function, Data.Ord, Data.Ring, Data.Semiring, Data.Show, Effect, Effect.Console, Prelude, Prim, Test.Primes
// TO REQUIRE: Bench, Control.Bind, Data.Eq, Data.EuclideanRing, Data.Function, Data.Ord, Data.Ring, Data.Semiring, Data.Show, Effect, Effect.Console, Prelude, Test.Primes
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
require_once __DIR__ . '/../Test.Primes/index.php';

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


final class Test_Primes_Nil { public $tag = 'Nil'; public function __construct() {} }
final class Test_Primes_Cons { public $tag = 'Cons'; public function __construct(public  $value0, public  $value1) {} }

// Test_Primes_Nil
$GLOBALS['Test_Primes_Nil'] = ($GLOBALS['__phpurs_data0_Nil'] ??= new \Test\Primes\Test_Primes_Nil());

// Test_Primes_Cons
$GLOBALS['Test_Primes_Cons'] = (function() {
  $__fn = function($value0, $value1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Test\Primes\Test_Primes_Cons($value0, $value1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Test_Primes_sumList
function majTest_majPrimes_summajList($lst_0): int|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majTest_majPrimes_summajList';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $go__go_1_0 = null;
  $go__go_1_0 = (function() use (&$go__go_1_0) {
  $__fn = function($v_2, $v1_3 = null) use (&$go__go_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_1_0_0_v_2 = $v_2;
  $__tco_var_go__go_1_0_0_v1_3 = $v1_3;
  tco_loop_go__go_1_0_0:;
  $v_2 = $__tco_var_go__go_1_0_0_v_2;
  $v1_3 = $__tco_var_go__go_1_0_0_v1_3;
  $__t0 = null;;
  if ($v_2 instanceof \Test\Primes\Test_Primes_Nil) {
$__t0 = $v1_3;
goto end_branch_0;;
};
  if ($v_2 instanceof \Test\Primes\Test_Primes_Cons) {
$__tco_1 = ($v_2)->{'value1'};
$__tco_2 = ($v1_3 + ($v_2)->{'value0'});
$__tco_var_go__go_1_0_0_v_2 = $__tco_1;
$__tco_var_go__go_1_0_0_v1_3 = $__tco_2;
goto tco_loop_go__go_1_0_0;;
$__t0 = null;
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
  $__res = (($go__go_1_0)($lst_0))(0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Test_Primes_sumList'] = __NAMESPACE__ . '\\majTest_majPrimes_summajList';

// Test_Primes_reverse
function majTest_majPrimes_reverse($lst_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majTest_majPrimes_reverse';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $go__go_1_0 = null;
  $go__go_1_0 = (function() use (&$go__go_1_0) {
  $__fn = function($v_2, $v1_3 = null) use (&$go__go_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_1_0_0_v_2 = $v_2;
  $__tco_var_go__go_1_0_0_v1_3 = $v1_3;
  tco_loop_go__go_1_0_0:;
  $v_2 = $__tco_var_go__go_1_0_0_v_2;
  $v1_3 = $__tco_var_go__go_1_0_0_v1_3;
  $__t0 = null;;
  if ($v_2 instanceof \Test\Primes\Test_Primes_Nil) {
$__t0 = $v1_3;
goto end_branch_0;;
};
  if ($v_2 instanceof \Test\Primes\Test_Primes_Cons) {
$__tco_1 = ($v_2)->{'value1'};
$__tco_2 = new \Test\Primes\Test_Primes_Cons(($v_2)->{'value0'}, $v1_3);
$__tco_var_go__go_1_0_0_v_2 = $__tco_1;
$__tco_var_go__go_1_0_0_v1_3 = $__tco_2;
goto tco_loop_go__go_1_0_0;;
$__t0 = null;
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
  $__res = (($go__go_1_0)($lst_0))(new \Test\Primes\Test_Primes_Nil());
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Test_Primes_reverse'] = __NAMESPACE__ . '\\majTest_majPrimes_reverse';

// Test_Primes_range
function majTest_majPrimes_range(int $start_0, $end_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majTest_majPrimes_range';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__go_2_0 = null;
  $go__go_2_0 = (function() use (&$go__go_2_0, $start_0) {
  $__fn = function(int $curr_3, $acc_4 = null) use (&$go__go_2_0, $start_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_2_0_0_curr_3 = $curr_3;
  $__tco_var_go__go_2_0_0_acc_4 = $acc_4;
  tco_loop_go__go_2_0_0:;
  $curr_3 = $__tco_var_go__go_2_0_0_curr_3;
  $acc_4 = $__tco_var_go__go_2_0_0_acc_4;
  $__t2 = null;;
  if (($curr_3 < $start_0)) {
$__t2 = $acc_4;
goto end_branch_2;;
};
  $__tco_0 = ($curr_3 - 1);
  $__tco_1 = new \Test\Primes\Test_Primes_Cons($curr_3, $acc_4);
  $__tco_var_go__go_2_0_0_curr_3 = $__tco_0;
  $__tco_var_go__go_2_0_0_acc_4 = $__tco_1;
  goto tco_loop_go__go_2_0_0;;
  $__t2 = null;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = (($go__go_2_0)($end_1))(new \Test\Primes\Test_Primes_Nil());
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Test_Primes_range'] = __NAMESPACE__ . '\\majTest_majPrimes_range';

// Test_Primes_filter
function majTest_majPrimes_filter($p_0, $lst_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majTest_majPrimes_filter';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__go_2_0 = null;
  $go__go_2_0 = (function() use (&$go__go_2_0, $p_0) {
  $__fn = function($v_3, $v1_4 = null) use (&$go__go_2_0, $p_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_2_0_0_v_3 = $v_3;
  $__tco_var_go__go_2_0_0_v1_4 = $v1_4;
  tco_loop_go__go_2_0_0:;
  $v_3 = $__tco_var_go__go_2_0_0_v_3;
  $v1_4 = $__tco_var_go__go_2_0_0_v1_4;
  $__t0 = null;;
  if ($v_3 instanceof \Test\Primes\Test_Primes_Nil) {
$go__go_5_1 = null;
$go__go_5_1 = (function() use (&$__tco_var_go__go_2_0_0_v_3, &$__tco_var_go__go_2_0_0_v1_4, &$go__go_5_1) {
  $__fn = function($v_6, $v1_7 = null) use (&$__tco_var_go__go_2_0_0_v_3, &$__tco_var_go__go_2_0_0_v1_4, &$go__go_5_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_5_1_1_v_6 = $v_6;
  $__tco_var_go__go_5_1_1_v1_7 = $v1_7;
  tco_loop_go__go_5_1_1:;
  $v_6 = $__tco_var_go__go_5_1_1_v_6;
  $v1_7 = $__tco_var_go__go_5_1_1_v1_7;
  $__t1 = null;;
  if ($v_6 instanceof \Test\Primes\Test_Primes_Nil) {
$__t1 = $v1_7;
goto end_branch_1;;
};
  if ($v_6 instanceof \Test\Primes\Test_Primes_Cons) {
$__tco_2 = ($v_6)->{'value1'};
$__tco_3 = new \Test\Primes\Test_Primes_Cons(($v_6)->{'value0'}, $v1_7);
$__tco_var_go__go_5_1_1_v_6 = $__tco_2;
$__tco_var_go__go_5_1_1_v1_7 = $__tco_3;
goto tco_loop_go__go_5_1_1;;
$__t1 = null;
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
})();
$__t0 = (($go__go_5_1)($v1_4))(new \Test\Primes\Test_Primes_Nil());
goto end_branch_0;;
};
  if ($v_3 instanceof \Test\Primes\Test_Primes_Cons) {
$__t4 = null;;
if (($p_0)(($v_3)->{'value0'})) {
$__tco_5 = ($v_3)->{'value1'};
$__tco_6 = new \Test\Primes\Test_Primes_Cons(($v_3)->{'value0'}, $v1_4);
$__tco_var_go__go_2_0_0_v_3 = $__tco_5;
$__tco_var_go__go_2_0_0_v1_4 = $__tco_6;
goto tco_loop_go__go_2_0_0;;
$__t4 = null;
goto end_branch_4;;
};
$__tco_2 = ($v_3)->{'value1'};
$__tco_3 = $v1_4;
$__tco_var_go__go_2_0_0_v_3 = $__tco_2;
$__tco_var_go__go_2_0_0_v1_4 = $__tco_3;
goto tco_loop_go__go_2_0_0;;
$__t4 = null;
end_branch_4:;
$__t0 = $__t4;
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
  $__res = (($go__go_2_0)($lst_1))(new \Test\Primes\Test_Primes_Nil());
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Test_Primes_filter'] = __NAMESPACE__ . '\\majTest_majPrimes_filter';

// Test_Primes_sieve
function majTest_majPrimes_sieve($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majTest_majPrimes_sieve';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__tco_var_Test_Primes_sieve_v_0 = $v_0;
  tco_loop_Test_Primes_sieve:;
  $v_0 = $__tco_var_Test_Primes_sieve_v_0;
  $__t0 = null;;
  if ($v_0 instanceof \Test\Primes\Test_Primes_Nil) {
$__t0 = new \Test\Primes\Test_Primes_Nil();
goto end_branch_0;;
};
  if ($v_0 instanceof \Test\Primes\Test_Primes_Cons) {
$__local_var_1_1 = ($v_0)->{'value0'};
$go__go_2_2 = null;
$go__go_2_2 = (function() use ($__local_var_1_1, &$go__go_2_2) {
  $__fn = function($v_3, $v1_4 = null) use ($__local_var_1_1, &$go__go_2_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_2_2_2_v_3 = $v_3;
  $__tco_var_go__go_2_2_2_v1_4 = $v1_4;
  tco_loop_go__go_2_2_2:;
  $v_3 = $__tco_var_go__go_2_2_2_v_3;
  $v1_4 = $__tco_var_go__go_2_2_2_v1_4;
  $__t2 = null;;
  if ($v_3 instanceof \Test\Primes\Test_Primes_Nil) {
$go__go_5_3 = null;
$go__go_5_3 = (function() use (&$__tco_var_go__go_2_2_2_v_3, &$__tco_var_go__go_2_2_2_v1_4, &$go__go_5_3) {
  $__fn = function($v_6, $v1_7 = null) use (&$__tco_var_go__go_2_2_2_v_3, &$__tco_var_go__go_2_2_2_v1_4, &$go__go_5_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_5_3_3_v_6 = $v_6;
  $__tco_var_go__go_5_3_3_v1_7 = $v1_7;
  tco_loop_go__go_5_3_3:;
  $v_6 = $__tco_var_go__go_5_3_3_v_6;
  $v1_7 = $__tco_var_go__go_5_3_3_v1_7;
  $__t3 = null;;
  if ($v_6 instanceof \Test\Primes\Test_Primes_Nil) {
$__t3 = $v1_7;
goto end_branch_3;;
};
  if ($v_6 instanceof \Test\Primes\Test_Primes_Cons) {
$__tco_4 = ($v_6)->{'value1'};
$__tco_5 = new \Test\Primes\Test_Primes_Cons(($v_6)->{'value0'}, $v1_7);
$__tco_var_go__go_5_3_3_v_6 = $__tco_4;
$__tco_var_go__go_5_3_3_v1_7 = $__tco_5;
goto tco_loop_go__go_5_3_3;;
$__t3 = null;
goto end_branch_3;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t3 = null;
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
$__t2 = (($go__go_5_3)($v1_4))(new \Test\Primes\Test_Primes_Nil());
goto end_branch_2;;
};
  if ($v_3 instanceof \Test\Primes\Test_Primes_Cons) {
$__t6 = null;;
if (( ! (\Data\EuclideanRing\majData_majEuclideanmajRing_intmajMod(($v_3)->{'value0'}, $__local_var_1_1) === 0))) {
$__tco_7 = ($v_3)->{'value1'};
$__tco_8 = new \Test\Primes\Test_Primes_Cons(($v_3)->{'value0'}, $v1_4);
$__tco_var_go__go_2_2_2_v_3 = $__tco_7;
$__tco_var_go__go_2_2_2_v1_4 = $__tco_8;
goto tco_loop_go__go_2_2_2;;
$__t6 = null;
goto end_branch_6;;
};
$__tco_4 = ($v_3)->{'value1'};
$__tco_5 = $v1_4;
$__tco_var_go__go_2_2_2_v_3 = $__tco_4;
$__tco_var_go__go_2_2_2_v1_4 = $__tco_5;
goto tco_loop_go__go_2_2_2;;
$__t6 = null;
end_branch_6:;
$__t2 = $__t6;
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
$__t0 = new \Test\Primes\Test_Primes_Cons($__local_var_1_1, \Test\Primes\majTest_majPrimes_sieve((($go__go_2_2)(($v_0)->{'value1'}))(new \Test\Primes\Test_Primes_Nil())));
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
$GLOBALS['Test_Primes_sieve'] = __NAMESPACE__ . '\\majTest_majPrimes_sieve';

// Test_Primes_describe
$GLOBALS['Test_Primes_describe'] = \Effect\Console\majEffect_majConsole_log("Prime Sieve (sum primes up to 500):");

// Test_Primes_act
$GLOBALS['Test_Primes_act'] = (function() use (&$__fn) {
$__local_var_0_0 = \Bench\majBench_opaque(500);
return function() use ($__local_var_0_0, &$__fn) {
$dummy_1_1 = phpurs_execute_effect($__local_var_0_0);
$go__go_2_2 = null;
$go__go_2_2 = (function() use (&$go__go_2_2) {
  $__fn = function($v_3, $v1_4 = null) use (&$go__go_2_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_2_2_2_v_3 = $v_3;
  $__tco_var_go__go_2_2_2_v1_4 = $v1_4;
  tco_loop_go__go_2_2_2:;
  $v_3 = $__tco_var_go__go_2_2_2_v_3;
  $v1_4 = $__tco_var_go__go_2_2_2_v1_4;
  $__t2 = null;;
  if ($v_3 instanceof \Test\Primes\Test_Primes_Nil) {
$__t2 = $v1_4;
goto end_branch_2;;
};
  if ($v_3 instanceof \Test\Primes\Test_Primes_Cons) {
$__tco_3 = ($v_3)->{'value1'};
$__tco_4 = ($v1_4 + ($v_3)->{'value0'});
$__tco_var_go__go_2_2_2_v_3 = $__tco_3;
$__tco_var_go__go_2_2_2_v1_4 = $__tco_4;
goto tco_loop_go__go_2_2_2;;
$__t2 = null;
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
$go__go_3_3 = null;
$go__go_3_3 = (function() use (&$go__go_3_3) {
  $__fn = function(int $curr_4, $acc_5 = null) use (&$go__go_3_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_3_3_3_curr_4 = $curr_4;
  $__tco_var_go__go_3_3_3_acc_5 = $acc_5;
  tco_loop_go__go_3_3_3:;
  $curr_4 = $__tco_var_go__go_3_3_3_curr_4;
  $acc_5 = $__tco_var_go__go_3_3_3_acc_5;
  $__t5 = null;;
  if (($curr_4 < 2)) {
$__t5 = $acc_5;
goto end_branch_5;;
};
  $__tco_3 = ($curr_4 - 1);
  $__tco_4 = new \Test\Primes\Test_Primes_Cons($curr_4, $acc_5);
  $__tco_var_go__go_3_3_3_curr_4 = $__tco_3;
  $__tco_var_go__go_3_3_3_acc_5 = $__tco_4;
  goto tco_loop_go__go_3_3_3;;
  $__t5 = null;
  end_branch_5:;
  $__res = $__t5;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
return phpurs_execute_effect(\Effect\Console\majEffect_majConsole_log(\Data\Show\majData_majShow_showmajIntmajImpl((($go__go_2_2)(\Test\Primes\majTest_majPrimes_sieve((($go__go_3_3)($dummy_1_1))(new \Test\Primes\Test_Primes_Nil()))))(0))));
};
})();

