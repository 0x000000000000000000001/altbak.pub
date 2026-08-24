<?php

namespace Data\Decidable;

// ALL IMPORTS: Control.Category, Data.Comparison, Data.Decidable, Data.Decide, Data.Divisible, Data.Equivalence, Data.Op, Data.Predicate, Data.Void, Prelude, Prim
// TO REQUIRE: Control.Category, Data.Comparison, Data.Decidable, Data.Decide, Data.Divisible, Data.Equivalence, Data.Op, Data.Predicate, Data.Void, Prelude
require_once __DIR__ . '/../Control.Category/index.php';
require_once __DIR__ . '/../Data.Comparison/index.php';
require_once __DIR__ . '/../Data.Decidable/index.php';
require_once __DIR__ . '/../Data.Decide/index.php';
require_once __DIR__ . '/../Data.Divisible/index.php';
require_once __DIR__ . '/../Data.Equivalence/index.php';
require_once __DIR__ . '/../Data.Op/index.php';
require_once __DIR__ . '/../Data.Predicate/index.php';
require_once __DIR__ . '/../Data.Void/index.php';
require_once __DIR__ . '/../Prelude/index.php';

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




// Data_Decidable_lose
function majData_majDecidable_lose($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majDecidable_lose';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'lose'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Decidable_lose'] = __NAMESPACE__ . '\\majData_majDecidable_lose';

// Data_Decidable_lost
function majData_majDecidable_lost($dictDecidable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majDecidable_lost';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($dictDecidable_0)->{'lose'})(function($x_1) {
  $__num = \func_num_args();
  $__res = $x_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Decidable_lost'] = __NAMESPACE__ . '\\majData_majDecidable_lost';

// Data_Decidable_decidablePredicate
$GLOBALS['Data_Decidable_decidablePredicate'] = (object)["lose" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($a_1) use ($f_0) {
  $__num = \func_num_args();
  $spin_2_0 = null;
  $spin_2_0 = function($v_3) use (&$spin_2_0) {
  $__num = \func_num_args();
  $__tco_var_spin_2_0_0_v_3 = $v_3;
  tco_loop_spin_2_0_0:;
  $v_3 = $__tco_var_spin_2_0_0_v_3;
  $__tco_0 = $v_3;
  $__tco_var_spin_2_0_0_v_3 = $__tco_0;
  goto tco_loop_spin_2_0_0;;
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = ($spin_2_0)(($f_0)($a_1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Decide0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Decide_choosePredicate'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Divisible1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Divisible_divisiblePredicate'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Decidable_decidableOp
function majData_majDecidable_decidablemajOp($dictMonoid_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majDecidable_decidablemajOp';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $divideOp_2_1 = (object)["divide" => function($f_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($__local_var_1_0, $f_2) {
  $__num = \func_num_args();
  $__res = function($v1_4) use ($__local_var_1_0, $f_2, $v_3) {
  $__num = \func_num_args();
  $__res = function($a_5) use ($__local_var_1_0, $f_2, $v1_4, $v_3) {
  $__num = \func_num_args();
  $v2_6_1 = ($f_2)($a_5);
  $__res = ((($__local_var_1_0)->{'append'})(($v_3)(($v2_6_1)->{'value0'})))(($v1_4)(($v2_6_1)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Contravariant0" => function($_dollar___unused_2) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Op_contravariantOp'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $chooseOp_1_0 = (object)["choose" => function($f_3) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($f_3) {
  $__num = \func_num_args();
  $__res = function($v1_5) use ($f_3, $v_4) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v2_6) use ($v1_5, $v_4) {
  $__num = \func_num_args();
  $__t3 = null;;
  if ($v2_6 instanceof \Data\Either\Data_Either_Left) {
$__t3 = ($v_4)(($v2_6)->{'value0'});
goto end_branch_3;;
};
  if ($v2_6 instanceof \Data\Either\Data_Either_Right) {
$__t3 = ($v1_5)(($v2_6)->{'value0'});
goto end_branch_3;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t3 = null;
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($f_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Divide0" => function($_dollar___unused_3) use ($divideOp_2_1) {
  $__num = \func_num_args();
  $__res = $divideOp_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_2_5 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $divideOp_2_5 = (object)["divide" => function($f_3) use ($__local_var_2_5) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($__local_var_2_5, $f_3) {
  $__num = \func_num_args();
  $__res = function($v1_5) use ($__local_var_2_5, $f_3, $v_4) {
  $__num = \func_num_args();
  $__res = function($a_6) use ($__local_var_2_5, $f_3, $v1_5, $v_4) {
  $__num = \func_num_args();
  $v2_7_6 = ($f_3)($a_6);
  $__res = ((($__local_var_2_5)->{'append'})(($v_4)(($v2_7_6)->{'value0'})))(($v1_5)(($v2_7_6)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Contravariant0" => function($_dollar___unused_3) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Op_contravariantOp'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_3_8 = ($dictMonoid_0)->{'mempty'};
  $divisibleOp_2_5 = (object)["conquer" => function($v_4) use ($__local_var_3_8) {
  $__num = \func_num_args();
  $__res = $__local_var_3_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Divide0" => function($_dollar___unused_3) use ($divideOp_2_5) {
  $__num = \func_num_args();
  $__res = $divideOp_2_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["lose" => function($f_3) {
  $__num = \func_num_args();
  $__res = function($a_4) use ($f_3) {
  $__num = \func_num_args();
  $spin_5_10 = null;
  $spin_5_10 = function($v_6) use (&$spin_5_10) {
  $__num = \func_num_args();
  $__tco_var_spin_5_10_10_v_6 = $v_6;
  tco_loop_spin_5_10_10:;
  $v_6 = $__tco_var_spin_5_10_10_v_6;
  $__tco_10 = $v_6;
  $__tco_var_spin_5_10_10_v_6 = $__tco_10;
  goto tco_loop_spin_5_10_10;;
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = ($spin_5_10)(($f_3)($a_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Decide0" => function($_dollar___unused_3) use ($chooseOp_1_0) {
  $__num = \func_num_args();
  $__res = $chooseOp_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Divisible1" => function($_dollar___unused_3) use ($divisibleOp_2_5) {
  $__num = \func_num_args();
  $__res = $divisibleOp_2_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Decidable_decidableOp'] = __NAMESPACE__ . '\\majData_majDecidable_decidablemajOp';

// Data_Decidable_decidableEquivalence
$GLOBALS['Data_Decidable_decidableEquivalence'] = (object)["lose" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($a_1) use ($f_0) {
  $__num = \func_num_args();
  $spin_2_0 = null;
  $spin_2_0 = function($v_3) use (&$spin_2_0) {
  $__num = \func_num_args();
  $__tco_var_spin_2_0_0_v_3 = $v_3;
  tco_loop_spin_2_0_0:;
  $v_3 = $__tco_var_spin_2_0_0_v_3;
  $__tco_0 = $v_3;
  $__tco_var_spin_2_0_0_v_3 = $__tco_0;
  goto tco_loop_spin_2_0_0;;
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = ($spin_2_0)(($f_0)($a_1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Decide0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Decide_chooseEquivalence'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Divisible1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Divisible_divisibleEquivalence'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Decidable_decidableComparison
$GLOBALS['Data_Decidable_decidableComparison'] = (object)["lose" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($a_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($a_1, $f_0) {
  $__num = \func_num_args();
  $spin_3_0 = null;
  $spin_3_0 = function($v_4) use (&$spin_3_0) {
  $__num = \func_num_args();
  $__tco_var_spin_3_0_0_v_4 = $v_4;
  tco_loop_spin_3_0_0:;
  $v_4 = $__tco_var_spin_3_0_0_v_4;
  $__tco_0 = $v_4;
  $__tco_var_spin_3_0_0_v_4 = $__tco_0;
  goto tco_loop_spin_3_0_0;;
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = ($spin_3_0)(($f_0)($a_1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Decide0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Decide_chooseComparison'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Divisible1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Divisible_divisibleComparison'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

