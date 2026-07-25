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
\PhpursThunks::$thunks['Data_Decidable_lose'] = function() { $v = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)->{'lose'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Decidable_lost'] = function() { $v = function($dictDecidable_0 = null) {
  $__num = \func_num_args();
  $__res = (($dictDecidable_0)->{'lose'})((($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))->{'identity'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Decidable_decidablePredicate'] = function() { $v = (object)["lose" => (function() {
  $__fn = function($f_0 = null, $a_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $spin_2_0 = null;
  $spin_2_0 = function($v_3 = null) use (&$spin_2_0) {
  $__num = \func_num_args();
  $__tco_var_spin_2_0_0_v_3 = $v_3;
  $__tco_done_spin_2_0_0 = false;
  $__tco_res_spin_2_0_0 = null;
  $__tco_loop_spin_2_0_0 = function($v_3 = null) use (&$__tco_done_spin_2_0_0, &$__tco_var_spin_2_0_0_v_3, &$spin_2_0) {
  $__num = \func_num_args();
  $__tco_done_spin_2_0_0 = true;
  $v_3 = $__tco_var_spin_2_0_0_v_3;
  $__tco_0 = $v_3;
  $__tco_var_spin_2_0_0_v_3 = $__tco_0;
  $__tco_done_spin_2_0_0 = false;
  $__res = null;
  goto __end;;
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  while (($__tco_done_spin_2_0_0 === false)) {
$__tco_res_spin_2_0_0 = ($__tco_loop_spin_2_0_0)($__tco_var_spin_2_0_0_v_3);
};
  $__res = $__tco_res_spin_2_0_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = ($spin_2_0)(($f_0)($a_1));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Decide0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Decide_choosePredicate'] ?? \PhpursThunks::eval('Data_Decide_choosePredicate'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Divisible1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Divisible_divisiblePredicate'] ?? \PhpursThunks::eval('Data_Divisible_divisiblePredicate'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Decidable_decidableOp'] = function() { $v = function($dictMonoid_0 = null) {
  $__num = \func_num_args();
  $chooseOp_1_0 = (($GLOBALS['Data_Decide_chooseOp'] ?? \PhpursThunks::eval('Data_Decide_chooseOp')))((($dictMonoid_0)->{'Semigroup0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $divisibleOp_2_1 = (($GLOBALS['Data_Divisible_divisibleOp'] ?? \PhpursThunks::eval('Data_Divisible_divisibleOp')))($dictMonoid_0);
  $__res = (object)["lose" => (function() {
  $__fn = function($f_3 = null, $a_4 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $spin_5_2 = null;
  $spin_5_2 = function($v_6 = null) use (&$spin_5_2) {
  $__num = \func_num_args();
  $__tco_var_spin_5_2_2_v_6 = $v_6;
  $__tco_done_spin_5_2_2 = false;
  $__tco_res_spin_5_2_2 = null;
  $__tco_loop_spin_5_2_2 = function($v_6 = null) use (&$__tco_done_spin_5_2_2, &$__tco_var_spin_5_2_2_v_6, &$spin_5_2) {
  $__num = \func_num_args();
  $__tco_done_spin_5_2_2 = true;
  $v_6 = $__tco_var_spin_5_2_2_v_6;
  $__tco_2 = $v_6;
  $__tco_var_spin_5_2_2_v_6 = $__tco_2;
  $__tco_done_spin_5_2_2 = false;
  $__res = null;
  goto __end;;
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  while (($__tco_done_spin_5_2_2 === false)) {
$__tco_res_spin_5_2_2 = ($__tco_loop_spin_5_2_2)($__tco_var_spin_5_2_2_v_6);
};
  $__res = $__tco_res_spin_5_2_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = ($spin_5_2)(($f_3)($a_4));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Decide0" => function($dollar__unused_3 = null) use ($chooseOp_1_0) {
  $__num = \func_num_args();
  $__res = $chooseOp_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Divisible1" => function($dollar__unused_3 = null) use ($divisibleOp_2_1) {
  $__num = \func_num_args();
  $__res = $divisibleOp_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Decidable_decidableEquivalence'] = function() { $v = (object)["lose" => (function() {
  $__fn = function($f_0 = null, $a_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $spin_2_0 = null;
  $spin_2_0 = function($v_3 = null) use (&$spin_2_0) {
  $__num = \func_num_args();
  $__tco_var_spin_2_0_0_v_3 = $v_3;
  $__tco_done_spin_2_0_0 = false;
  $__tco_res_spin_2_0_0 = null;
  $__tco_loop_spin_2_0_0 = function($v_3 = null) use (&$__tco_done_spin_2_0_0, &$__tco_var_spin_2_0_0_v_3, &$spin_2_0) {
  $__num = \func_num_args();
  $__tco_done_spin_2_0_0 = true;
  $v_3 = $__tco_var_spin_2_0_0_v_3;
  $__tco_0 = $v_3;
  $__tco_var_spin_2_0_0_v_3 = $__tco_0;
  $__tco_done_spin_2_0_0 = false;
  $__res = null;
  goto __end;;
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  while (($__tco_done_spin_2_0_0 === false)) {
$__tco_res_spin_2_0_0 = ($__tco_loop_spin_2_0_0)($__tco_var_spin_2_0_0_v_3);
};
  $__res = $__tco_res_spin_2_0_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = ($spin_2_0)(($f_0)($a_1));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Decide0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Decide_chooseEquivalence'] ?? \PhpursThunks::eval('Data_Decide_chooseEquivalence'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Divisible1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Divisible_divisibleEquivalence'] ?? \PhpursThunks::eval('Data_Divisible_divisibleEquivalence'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Decidable_decidableComparison'] = function() { $v = (object)["lose" => (function() {
  $__fn = function($f_0 = null, $a_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $spin_3_0 = null;
  $spin_3_0 = function($v_4 = null) use (&$spin_3_0) {
  $__num = \func_num_args();
  $__tco_var_spin_3_0_0_v_4 = $v_4;
  $__tco_done_spin_3_0_0 = false;
  $__tco_res_spin_3_0_0 = null;
  $__tco_loop_spin_3_0_0 = function($v_4 = null) use (&$__tco_done_spin_3_0_0, &$__tco_var_spin_3_0_0_v_4, &$spin_3_0) {
  $__num = \func_num_args();
  $__tco_done_spin_3_0_0 = true;
  $v_4 = $__tco_var_spin_3_0_0_v_4;
  $__tco_0 = $v_4;
  $__tco_var_spin_3_0_0_v_4 = $__tco_0;
  $__tco_done_spin_3_0_0 = false;
  $__res = null;
  goto __end;;
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  while (($__tco_done_spin_3_0_0 === false)) {
$__tco_res_spin_3_0_0 = ($__tco_loop_spin_3_0_0)($__tco_var_spin_3_0_0_v_4);
};
  $__res = $__tco_res_spin_3_0_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = ($spin_3_0)(($f_0)($a_1));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "Decide0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Decide_chooseComparison'] ?? \PhpursThunks::eval('Data_Decide_chooseComparison'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Divisible1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Divisible_divisibleComparison'] ?? \PhpursThunks::eval('Data_Divisible_divisibleComparison'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };








