<?php

namespace Test\Primes;

// ALL IMPORTS: Data.Eq, Data.EuclideanRing, Data.Function, Data.Ord, Data.Ring, Data.Semiring, Data.Show, Effect, Effect.Console, Prelude, Prim, Test.Primes
// TO REQUIRE: Data.Eq, Data.EuclideanRing, Data.Function, Data.Ord, Data.Ring, Data.Semiring, Data.Show, Effect, Effect.Console, Prelude, Test.Primes
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
\PhpursThunks::$thunks['Test_Primes_Nil'] = function() { $v = ($GLOBALS['__phpurs_data0_Nil'] ??= new Phpurs_Data0("Nil")); return $v; };
\PhpursThunks::$thunks['Test_Primes_Cons'] = function() { $v = (function() {
  $__fn = function($value0, $value1 = null) use (&$__fn) {
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
\PhpursThunks::$thunks['Test_Primes_sumList'] = function() { $v = function($lst_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Test_Primes_sumList"), recVars=[];
  $go_1_0 = null;
  $go_1_0 = (function() use (&$go_1_0) {
  $__fn = function($v_2, $v1_3 = null) use (&$go_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_1_0"), recVars=["go_1_0"];
  while (true) {
if ((is_object($v_2) && (($v_2)->tag === "Nil"))) {
$__t1 = $v1_3;
} else {
if ((is_object($v_2) && (($v_2)->tag === "Cons"))) {
$__tco_2 = ($v_2)->value1;
$__tco_3 = ((($GLOBALS['Data_Semiring_intAdd'] ?? \PhpursThunks::eval('Data_Semiring_intAdd')))($v1_3))(($v_2)->value0);
$v_2 = $__tco_2;
$v1_3 = $__tco_3;
continue ;
$__t1 = null;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
$__res = $__t1;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = (($go_1_0)($lst_0))(0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Test_Primes_reverse'] = function() { $v = function($lst_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Test_Primes_reverse"), recVars=[];
  $go_1_0 = null;
  $go_1_0 = (function() use (&$go_1_0) {
  $__fn = function($v_2, $v1_3 = null) use (&$go_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_1_0"), recVars=["go_1_0"];
  while (true) {
if ((is_object($v_2) && (($v_2)->tag === "Nil"))) {
$__t1 = $v1_3;
} else {
if ((is_object($v_2) && (($v_2)->tag === "Cons"))) {
$__tco_2 = ($v_2)->value1;
$__tco_3 = new Phpurs_Data2("Cons", ($v_2)->value0, $v1_3);
$v_2 = $__tco_2;
$v1_3 = $__tco_3;
continue ;
$__t1 = null;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
$__res = $__t1;
goto __end;;
};
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
\PhpursThunks::$thunks['Test_Primes_range'] = function() { $v = (function() {
  $__fn = function($start_0, $end_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Test_Primes_range"), recVars=[];
  $go_2_0 = null;
  $go_2_0 = (function() use (&$go_2_0, $start_0) {
  $__fn = function($curr_3, $acc_4 = null) use (&$go_2_0, $start_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_2_0"), recVars=["go_2_0"];
  while (true) {
if ((is_object((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($curr_3))($start_0)) && (((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($curr_3))($start_0))->tag === "LT"))) {
$__t3 = $acc_4;
} else {
$__tco_1 = ((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))($curr_3))(1);
$__tco_2 = new Phpurs_Data2("Cons", $curr_3, $acc_4);
$curr_3 = $__tco_1;
$acc_4 = $__tco_2;
continue ;
$__t3 = null;
};
$__res = $__t3;
goto __end;;
};
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
\PhpursThunks::$thunks['Test_Primes_filter'] = function() { $v = (function() {
  $__fn = function($p_0, $lst_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Test_Primes_filter"), recVars=[];
  $go_2_0 = null;
  $go_2_0 = (function() use (&$go_2_0, $p_0) {
  $__fn = function($v_3, $v1_4 = null) use (&$go_2_0, $p_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_2_0"), recVars=["go_2_0"];
  while (true) {
if ((is_object($v_3) && (($v_3)->tag === "Nil"))) {
$go_5_2 = null;
$go_5_2 = (function() use (&$go_5_2) {
  $__fn = function($v_6, $v1_7 = null) use (&$go_5_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_5_2"), recVars=["go_2_0","go_5_2"];
  while (true) {
if ((is_object($v_6) && (($v_6)->tag === "Nil"))) {
$__t3 = $v1_7;
} else {
if ((is_object($v_6) && (($v_6)->tag === "Cons"))) {
$__tco_4 = ($v_6)->value1;
$__tco_5 = new Phpurs_Data2("Cons", ($v_6)->value0, $v1_7);
$v_6 = $__tco_4;
$v1_7 = $__tco_5;
continue ;
$__t3 = null;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
};
};
$__res = $__t3;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
$__t1 = (($go_5_2)($v1_4))(new Phpurs_Data0("Nil"));
} else {
if ((is_object($v_3) && (($v_3)->tag === "Cons"))) {
if (($p_0)(($v_3)->value0)) {
$__tco_9 = ($v_3)->value1;
$__tco_10 = new Phpurs_Data2("Cons", ($v_3)->value0, $v1_4);
$v_3 = $__tco_9;
$v1_4 = $__tco_10;
continue ;
$__t8 = null;
} else {
$__tco_6 = ($v_3)->value1;
$__tco_7 = $v1_4;
$v_3 = $__tco_6;
$v1_4 = $__tco_7;
continue ;
$__t8 = null;
};
$__t1 = $__t8;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
$__res = $__t1;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = (($go_2_0)($lst_1))(new Phpurs_Data0("Nil"));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Test_Primes_sieve'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Test_Primes_sieve"), recVars=["Test_Primes_sieve"];
  while (true) {
if ((is_object($v_0) && (($v_0)->tag === "Nil"))) {
$__t0 = new Phpurs_Data0("Nil");
} else {
if ((is_object($v_0) && (($v_0)->tag === "Cons"))) {
$__local_var_1_1 = ($v_0)->value0;
$go_2_2 = null;
$go_2_2 = (function() use ($__local_var_1_1, &$go_2_2) {
  $__fn = function($v_3, $v1_4 = null) use ($__local_var_1_1, &$go_2_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_2_2"), recVars=["Test_Primes_sieve","go_2_2"];
  while (true) {
if ((is_object($v_3) && (($v_3)->tag === "Nil"))) {
$go_5_4 = null;
$go_5_4 = (function() use (&$go_5_4) {
  $__fn = function($v_6, $v1_7 = null) use (&$go_5_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_5_4"), recVars=["Test_Primes_sieve","go_2_2","go_5_4"];
  while (true) {
if ((is_object($v_6) && (($v_6)->tag === "Nil"))) {
$__t5 = $v1_7;
} else {
if ((is_object($v_6) && (($v_6)->tag === "Cons"))) {
$__tco_6 = ($v_6)->value1;
$__tco_7 = new Phpurs_Data2("Cons", ($v_6)->value0, $v1_7);
$v_6 = $__tco_6;
$v1_7 = $__tco_7;
continue ;
$__t5 = null;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t5 = null;
};
};
$__res = $__t5;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
$__t3 = (($go_5_4)($v1_4))(new Phpurs_Data0("Nil"));
} else {
if ((is_object($v_3) && (($v_3)->tag === "Cons"))) {
if (((($GLOBALS['Data_Eq_eqBooleanImpl'] ?? \PhpursThunks::eval('Data_Eq_eqBooleanImpl')))(((($GLOBALS['Data_Eq_eqIntImpl'] ?? \PhpursThunks::eval('Data_Eq_eqIntImpl')))(((($GLOBALS['Data_EuclideanRing_intMod'] ?? \PhpursThunks::eval('Data_EuclideanRing_intMod')))(($v_3)->value0))($__local_var_1_1)))(0)))(false)) {
$__tco_11 = ($v_3)->value1;
$__tco_12 = new Phpurs_Data2("Cons", ($v_3)->value0, $v1_4);
$v_3 = $__tco_11;
$v1_4 = $__tco_12;
continue ;
$__t10 = null;
} else {
$__tco_8 = ($v_3)->value1;
$__tco_9 = $v1_4;
$v_3 = $__tco_8;
$v1_4 = $__tco_9;
continue ;
$__t10 = null;
};
$__t3 = $__t10;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
};
};
$__res = $__t3;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
$__t0 = new Phpurs_Data2("Cons", $__local_var_1_1, (($GLOBALS['Test_Primes_sieve'] ?? \PhpursThunks::eval('Test_Primes_sieve')))((($go_2_2)(($v_0)->value1))(new Phpurs_Data0("Nil"))));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
$__res = $__t0;
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Test_Primes_describe'] = function() { $v = (($GLOBALS['Effect_Console_log'] ?? \PhpursThunks::eval('Effect_Console_log')))("Prime Sieve (sum primes up to 500):"); return $v; };
\PhpursThunks::$thunks['Test_Primes_act'] = function() { $v = (($GLOBALS['Effect_Console_log'] ?? \PhpursThunks::eval('Effect_Console_log')))((($GLOBALS['Data_Show_showIntImpl'] ?? \PhpursThunks::eval('Data_Show_showIntImpl')))((($GLOBALS['Test_Primes_sumList'] ?? \PhpursThunks::eval('Test_Primes_sumList')))((($GLOBALS['Test_Primes_sieve'] ?? \PhpursThunks::eval('Test_Primes_sieve')))(((($GLOBALS['Test_Primes_range'] ?? \PhpursThunks::eval('Test_Primes_range')))(2))(500))))); return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };











