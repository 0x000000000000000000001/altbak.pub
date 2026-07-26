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

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };


// Test_Primes_lessThan
$GLOBALS['Test_Primes_lessThan'] = (function() use (&$__fn) {
$__local_var_0_0 = ((($GLOBALS['Data_Ord_ordIntImpl'])(new Phpurs_Data0("LT")))(new Phpurs_Data0("EQ")))(new Phpurs_Data0("GT"));
return (function() use ($__local_var_0_0) {
  $__fn = function($a1_1 = null, $a2_2 = null) use ($__local_var_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (is_object((($__local_var_0_0)($a1_1))($a2_2)) && (((($__local_var_0_0)($a1_1))($a2_2))->{'tag'} === "LT"));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
})();

// Test_Primes_Nil
$GLOBALS['Test_Primes_Nil'] = ($GLOBALS['__phpurs_data0_Nil'] ??= new Phpurs_Data0("Nil"));

// Test_Primes_Cons
$GLOBALS['Test_Primes_Cons'] = (function() {
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
})();

// Test_Primes_sumList
$GLOBALS['Test_Primes_sumList'] = function($lst_0 = null) {
  $__num = \func_num_args();
  $go__1_0 = null;
  $go__1_0 = (function() use (&$go__1_0) {
  $__fn = function($v_2 = null, $v1_3 = null) use (&$go__1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__1_0_0_v_2 = $v_2;
  $__tco_var_go__1_0_0_v1_3 = $v1_3;
  tco_loop_go__1_0_0:;
  $v_2 = $__tco_var_go__1_0_0_v_2;
  $v1_3 = $__tco_var_go__1_0_0_v1_3;
  $__t0 = null;;
  if ((is_object($v_2) && (($v_2)->{'tag'} === "Nil"))) {
$__t0 = $v1_3;
goto end_branch_0;;
};
  if ((is_object($v_2) && (($v_2)->{'tag'} === "Cons"))) {
$__tco_1 = ($v_2)->{'value1'};
$__tco_2 = ($v1_3 + ($v_2)->{'value0'});
$__tco_var_go__1_0_0_v_2 = $__tco_1;
$__tco_var_go__1_0_0_v1_3 = $__tco_2;
goto tco_loop_go__1_0_0;;
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
  $__res = (($go__1_0)($lst_0))(0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Test_Primes_reverse
$GLOBALS['Test_Primes_reverse'] = function($lst_0 = null) {
  $__num = \func_num_args();
  $go__1_0 = null;
  $go__1_0 = (function() use (&$go__1_0) {
  $__fn = function($v_2 = null, $v1_3 = null) use (&$go__1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__1_0_0_v_2 = $v_2;
  $__tco_var_go__1_0_0_v1_3 = $v1_3;
  tco_loop_go__1_0_0:;
  $v_2 = $__tco_var_go__1_0_0_v_2;
  $v1_3 = $__tco_var_go__1_0_0_v1_3;
  $__t0 = null;;
  if ((is_object($v_2) && (($v_2)->{'tag'} === "Nil"))) {
$__t0 = $v1_3;
goto end_branch_0;;
};
  if ((is_object($v_2) && (($v_2)->{'tag'} === "Cons"))) {
$__tco_1 = ($v_2)->{'value1'};
$__tco_2 = new Phpurs_Data2("Cons", ($v_2)->{'value0'}, $v1_3);
$__tco_var_go__1_0_0_v_2 = $__tco_1;
$__tco_var_go__1_0_0_v1_3 = $__tco_2;
goto tco_loop_go__1_0_0;;
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
  $__res = (($go__1_0)($lst_0))(new Phpurs_Data0("Nil"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Test_Primes_range
$GLOBALS['Test_Primes_range'] = (function() {
  $__fn = function($start_0 = null, $end_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__2_0 = null;
  $go__2_0 = (function() use (&$go__2_0, $start_0) {
  $__fn = function($curr_3 = null, $acc_4 = null) use (&$go__2_0, $start_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__2_0_0_curr_3 = $curr_3;
  $__tco_var_go__2_0_0_acc_4 = $acc_4;
  tco_loop_go__2_0_0:;
  $curr_3 = $__tco_var_go__2_0_0_curr_3;
  $acc_4 = $__tco_var_go__2_0_0_acc_4;
  $__t2 = null;;
  if ((($GLOBALS['Test_Primes_lessThan'])($curr_3))($start_0)) {
$__t2 = $acc_4;
goto end_branch_2;;
};
  $__tco_0 = ($curr_3 - 1);
  $__tco_1 = new Phpurs_Data2("Cons", $curr_3, $acc_4);
  $__tco_var_go__2_0_0_curr_3 = $__tco_0;
  $__tco_var_go__2_0_0_acc_4 = $__tco_1;
  goto tco_loop_go__2_0_0;;
  $__t2 = null;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = (($go__2_0)($end_1))(new Phpurs_Data0("Nil"));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Test_Primes_filter
$GLOBALS['Test_Primes_filter'] = (function() {
  $__fn = function($p_0 = null, $lst_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__2_0 = null;
  $go__2_0 = (function() use (&$go__2_0, $p_0) {
  $__fn = function($v_3 = null, $v1_4 = null) use (&$go__2_0, $p_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__2_0_0_v_3 = $v_3;
  $__tco_var_go__2_0_0_v1_4 = $v1_4;
  tco_loop_go__2_0_0:;
  $v_3 = $__tco_var_go__2_0_0_v_3;
  $v1_4 = $__tco_var_go__2_0_0_v1_4;
  $__t0 = null;;
  if ((is_object($v_3) && (($v_3)->{'tag'} === "Nil"))) {
$go__5_1 = null;
$go__5_1 = (function() use (&$__tco_var_go__2_0_0_v_3, &$__tco_var_go__2_0_0_v1_4, &$go__5_1) {
  $__fn = function($v_6 = null, $v1_7 = null) use (&$__tco_var_go__2_0_0_v_3, &$__tco_var_go__2_0_0_v1_4, &$go__5_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__5_1_1_v_6 = $v_6;
  $__tco_var_go__5_1_1_v1_7 = $v1_7;
  tco_loop_go__5_1_1:;
  $v_6 = $__tco_var_go__5_1_1_v_6;
  $v1_7 = $__tco_var_go__5_1_1_v1_7;
  $__t1 = null;;
  if ((is_object($v_6) && (($v_6)->{'tag'} === "Nil"))) {
$__t1 = $v1_7;
goto end_branch_1;;
};
  if ((is_object($v_6) && (($v_6)->{'tag'} === "Cons"))) {
$__tco_2 = ($v_6)->{'value1'};
$__tco_3 = new Phpurs_Data2("Cons", ($v_6)->{'value0'}, $v1_7);
$__tco_var_go__5_1_1_v_6 = $__tco_2;
$__tco_var_go__5_1_1_v1_7 = $__tco_3;
goto tco_loop_go__5_1_1;;
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
$__t0 = (($go__5_1)($v1_4))(new Phpurs_Data0("Nil"));
goto end_branch_0;;
};
  if ((is_object($v_3) && (($v_3)->{'tag'} === "Cons"))) {
$__t4 = null;;
if (($p_0)(($v_3)->{'value0'})) {
$__tco_5 = ($v_3)->{'value1'};
$__tco_6 = new Phpurs_Data2("Cons", ($v_3)->{'value0'}, $v1_4);
$__tco_var_go__2_0_0_v_3 = $__tco_5;
$__tco_var_go__2_0_0_v1_4 = $__tco_6;
goto tco_loop_go__2_0_0;;
$__t4 = null;
goto end_branch_4;;
};
$__tco_2 = ($v_3)->{'value1'};
$__tco_3 = $v1_4;
$__tco_var_go__2_0_0_v_3 = $__tco_2;
$__tco_var_go__2_0_0_v1_4 = $__tco_3;
goto tco_loop_go__2_0_0;;
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
  $__res = (($go__2_0)($lst_1))(new Phpurs_Data0("Nil"));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Test_Primes_sieve
$GLOBALS['Test_Primes_sieve'] = function($v_0 = null) {
  $__num = \func_num_args();
  $__tco_var_Test_Primes_sieve_v_0 = $v_0;
  tco_loop_Test_Primes_sieve:;
  $v_0 = $__tco_var_Test_Primes_sieve_v_0;
  $__t0 = null;;
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Nil"))) {
$__t0 = new Phpurs_Data0("Nil");
goto end_branch_0;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Cons"))) {
$__local_var_1_1 = ($v_0)->{'value0'};
$go__2_2 = null;
$go__2_2 = (function() use ($__local_var_1_1, &$go__2_2) {
  $__fn = function($v_3 = null, $v1_4 = null) use ($__local_var_1_1, &$go__2_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__2_2_2_v_3 = $v_3;
  $__tco_var_go__2_2_2_v1_4 = $v1_4;
  tco_loop_go__2_2_2:;
  $v_3 = $__tco_var_go__2_2_2_v_3;
  $v1_4 = $__tco_var_go__2_2_2_v1_4;
  $__t2 = null;;
  if ((is_object($v_3) && (($v_3)->{'tag'} === "Nil"))) {
$go__5_3 = null;
$go__5_3 = (function() use (&$__tco_var_go__2_2_2_v_3, &$__tco_var_go__2_2_2_v1_4, &$go__5_3) {
  $__fn = function($v_6 = null, $v1_7 = null) use (&$__tco_var_go__2_2_2_v_3, &$__tco_var_go__2_2_2_v1_4, &$go__5_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__5_3_3_v_6 = $v_6;
  $__tco_var_go__5_3_3_v1_7 = $v1_7;
  tco_loop_go__5_3_3:;
  $v_6 = $__tco_var_go__5_3_3_v_6;
  $v1_7 = $__tco_var_go__5_3_3_v1_7;
  $__t3 = null;;
  if ((is_object($v_6) && (($v_6)->{'tag'} === "Nil"))) {
$__t3 = $v1_7;
goto end_branch_3;;
};
  if ((is_object($v_6) && (($v_6)->{'tag'} === "Cons"))) {
$__tco_4 = ($v_6)->{'value1'};
$__tco_5 = new Phpurs_Data2("Cons", ($v_6)->{'value0'}, $v1_7);
$__tco_var_go__5_3_3_v_6 = $__tco_4;
$__tco_var_go__5_3_3_v1_7 = $__tco_5;
goto tco_loop_go__5_3_3;;
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
$__t2 = (($go__5_3)($v1_4))(new Phpurs_Data0("Nil"));
goto end_branch_2;;
};
  if ((is_object($v_3) && (($v_3)->{'tag'} === "Cons"))) {
$__t6 = null;;
if (((($GLOBALS['Data_Eq_eqBoolean'])['eq'])((((($GLOBALS['Data_EuclideanRing_euclideanRingInt'])['mod'])(($v_3)->{'value0'}))($__local_var_1_1) === 0)))(false)) {
$__tco_7 = ($v_3)->{'value1'};
$__tco_8 = new Phpurs_Data2("Cons", ($v_3)->{'value0'}, $v1_4);
$__tco_var_go__2_2_2_v_3 = $__tco_7;
$__tco_var_go__2_2_2_v1_4 = $__tco_8;
goto tco_loop_go__2_2_2;;
$__t6 = null;
goto end_branch_6;;
};
$__tco_4 = ($v_3)->{'value1'};
$__tco_5 = $v1_4;
$__tco_var_go__2_2_2_v_3 = $__tco_4;
$__tco_var_go__2_2_2_v1_4 = $__tco_5;
goto tco_loop_go__2_2_2;;
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
$__t0 = new Phpurs_Data2("Cons", $__local_var_1_1, ($GLOBALS['Test_Primes_sieve'])((($go__2_2)(($v_0)->{'value1'}))(new Phpurs_Data0("Nil"))));
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Test_Primes_describe
$GLOBALS['Test_Primes_describe'] = ($GLOBALS['Effect_Console_log'])("Prime Sieve (sum primes up to 500):");

// Test_Primes_act
$GLOBALS['Test_Primes_act'] = ((($GLOBALS['Effect_bindEffect'])['bind'])(($GLOBALS['Bench_opaque'])(500)))(function($dummy_0 = null) {
  $__num = \func_num_args();
  $go__1_0 = null;
  $go__1_0 = (function() use (&$go__1_0) {
  $__fn = function($v_2 = null, $v1_3 = null) use (&$go__1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__1_0_0_v_2 = $v_2;
  $__tco_var_go__1_0_0_v1_3 = $v1_3;
  tco_loop_go__1_0_0:;
  $v_2 = $__tco_var_go__1_0_0_v_2;
  $v1_3 = $__tco_var_go__1_0_0_v1_3;
  $__t0 = null;;
  if ((is_object($v_2) && (($v_2)->{'tag'} === "Nil"))) {
$__t0 = $v1_3;
goto end_branch_0;;
};
  if ((is_object($v_2) && (($v_2)->{'tag'} === "Cons"))) {
$__tco_1 = ($v_2)->{'value1'};
$__tco_2 = ($v1_3 + ($v_2)->{'value0'});
$__tco_var_go__1_0_0_v_2 = $__tco_1;
$__tco_var_go__1_0_0_v1_3 = $__tco_2;
goto tco_loop_go__1_0_0;;
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
  $__res = ($GLOBALS['Effect_Console_log'])((($GLOBALS['Data_Show_showInt'])['show'])((($go__1_0)(($GLOBALS['Test_Primes_sieve'])((($GLOBALS['Test_Primes_range'])(2))($dummy_0))))(0)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});

