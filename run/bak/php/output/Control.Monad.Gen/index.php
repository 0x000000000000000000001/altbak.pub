<?php

namespace Control\Monad\Gen;

// ALL IMPORTS: Control.Applicative, Control.Bind, Control.Monad.Gen, Control.Monad.Gen.Class, Control.Monad.Rec.Class, Control.Semigroupoid, Data.Boolean, Data.Foldable, Data.Function, Data.Functor, Data.Maybe, Data.Monoid.Additive, Data.Newtype, Data.Ord, Data.Ring, Data.Semigroup, Data.Semigroup.Foldable, Data.Semigroup.Last, Data.Semiring, Data.Tuple, Data.Unfoldable, Data.Unit, Prelude, Prim
// TO REQUIRE: Control.Applicative, Control.Bind, Control.Monad.Gen, Control.Monad.Gen.Class, Control.Monad.Rec.Class, Control.Semigroupoid, Data.Boolean, Data.Foldable, Data.Function, Data.Functor, Data.Maybe, Data.Monoid.Additive, Data.Newtype, Data.Ord, Data.Ring, Data.Semigroup, Data.Semigroup.Foldable, Data.Semigroup.Last, Data.Semiring, Data.Tuple, Data.Unfoldable, Data.Unit, Prelude
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Monad.Gen/index.php';
require_once __DIR__ . '/../Control.Monad.Gen.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Rec.Class/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Boolean/index.php';
require_once __DIR__ . '/../Data.Foldable/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Monoid.Additive/index.php';
require_once __DIR__ . '/../Data.Newtype/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ring/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Semigroup.Foldable/index.php';
require_once __DIR__ . '/../Data.Semigroup.Last/index.php';
require_once __DIR__ . '/../Data.Semiring/index.php';
require_once __DIR__ . '/../Data.Tuple/index.php';
require_once __DIR__ . '/../Data.Unfoldable/index.php';
require_once __DIR__ . '/../Data.Unit/index.php';
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

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };


// Control_Monad_Gen_monoidAdditive
$GLOBALS['Control_Monad_Gen_monoidAdditive'] = (function() use (&$__fn) {
$semigroupAdditive1_0_0 = ["append" => (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_Semiring_semiringNumber'])['add'])($v_0))($v1_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
return ["mempty" => ($GLOBALS['Data_Semiring_semiringNumber'])['zero'], "Semigroup0" => function($_dollar__unused_1 = null) use ($semigroupAdditive1_0_0) {
  $__num = \func_num_args();
  $__res = $semigroupAdditive1_0_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
})();

// Control_Monad_Gen_Cons
$GLOBALS['Control_Monad_Gen_Cons'] = (function() {
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

// Control_Monad_Gen_Nil
$GLOBALS['Control_Monad_Gen_Nil'] = ($GLOBALS['__phpurs_data0_Nil'] ??= new Phpurs_Data0("Nil"));

// Control_Monad_Gen_unfoldable
$GLOBALS['Control_Monad_Gen_unfoldable'] = (function() {
  $__fn = function($dictMonadRec_0 = null, $dictMonadGen_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $Monad0_2_0 = (($dictMonadGen_1)['Monad0'])(null);
  $__local_var_3_1 = (($Monad0_2_0)['Applicative0'])(null);
  $Bind1_4_2 = (($Monad0_2_0)['Bind1'])(null);
  $__res = (function() use ($Bind1_4_2, $__local_var_3_1, $dictMonadGen_1, $dictMonadRec_0) {
  $__fn = function($dictUnfoldable_5 = null, $gen_6 = null) use ($Bind1_4_2, $__local_var_3_1, $dictMonadGen_1, $dictMonadRec_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((((($Bind1_4_2)['Apply0'])(null))['Functor0'])(null))['map'])((($dictUnfoldable_5)['unfoldr'])(function($v_7 = null) {
  $__num = \func_num_args();
  $__t3 = null;;
  if ((is_object($v_7) && (($v_7)->{'tag'} === "Nil"))) {
$__t3 = new Phpurs_Data0("Nothing");
goto end_branch_3;;
};
  if ((is_object($v_7) && (($v_7)->{'tag'} === "Cons"))) {
$__t3 = new Phpurs_Data1("Just", new Phpurs_Data2("Tuple", ($v_7)->{'value0'}, ($v_7)->{'value1'}));
goto end_branch_3;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t3 = null;
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))((($dictMonadGen_1)['sized'])(((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])((($dictMonadRec_0)['tailRecM'])(function($v_7 = null) use ($Bind1_4_2, $__local_var_3_1, $gen_6) {
  $__num = \func_num_args();
  $__t6 = null;;
  if ((($v_7)->{'value1'} <= 0)) {
$__t6 = (($__local_var_3_1)['pure'])(new Phpurs_Data1("Done", ($v_7)->{'value0'}));
goto end_branch_6;;
};
  $__local_var_8_4 = ($v_7)->{'value0'};
  $__local_var_9_5 = ($v_7)->{'value1'};
  $__t6 = ((($Bind1_4_2)['bind'])($gen_6))(function($x_10 = null) use ($__local_var_3_1, $__local_var_8_4, $__local_var_9_5) {
  $__num = \func_num_args();
  $__res = (($__local_var_3_1)['pure'])(new Phpurs_Data1("Loop", new Phpurs_Data2("Tuple", new Phpurs_Data2("Cons", $x_10, $__local_var_8_4), ((($GLOBALS['Data_Ring_ringInt'])['sub'])($__local_var_9_5))(1))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  end_branch_6:;
  $__res = $__t6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(($GLOBALS['Data_Tuple_Tuple'])(new Phpurs_Data0("Nil")))));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Control_Monad_Gen_semigroupFreqSemigroup
$GLOBALS['Control_Monad_Gen_semigroupFreqSemigroup'] = ["append" => (function() {
  $__fn = function($v_0 = null, $v1_1 = null, $pos_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $v2_3_0 = ($v_0)($pos_2);
  $__t1 = null;;
  if ((is_object(($v2_3_0)->{'value0'}) && ((($v2_3_0)->{'value0'})->{'tag'} === "Just"))) {
$__t1 = ($v1_1)((($v2_3_0)->{'value0'})->{'value0'});
goto end_branch_1;;
};
  $__t1 = $v2_3_0;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()];

// Control_Monad_Gen_fromIndex
$GLOBALS['Control_Monad_Gen_fromIndex'] = function($dictFoldable1_0 = null) {
  $__num = \func_num_args();
  $foldMap1_1_0 = (($dictFoldable1_0)['foldMap1'])($GLOBALS['Data_Semigroup_Last_semigroupLast']);
  $__res = (function() use ($dictFoldable1_0, $foldMap1_1_0) {
  $__fn = function($i_2 = null, $xs_3 = null) use ($dictFoldable1_0, $foldMap1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__4_1 = null;
  $go__4_1 = (function() use ($foldMap1_1_0, &$go__4_1, $xs_3) {
  $__fn = function($v_5 = null, $v1_6 = null) use ($foldMap1_1_0, &$go__4_1, $xs_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__4_1_1_v_5 = $v_5;
  $__tco_var_go__4_1_1_v1_6 = $v1_6;
  tco_loop_go__4_1_1:;
  $v_5 = $__tco_var_go__4_1_1_v_5;
  $v1_6 = $__tco_var_go__4_1_1_v1_6;
  $__t1 = null;;
  if ((is_object($v1_6) && (($v1_6)->{'tag'} === "Cons"))) {
$__t4 = null;;
if ((is_object(($v1_6)->{'value1'}) && ((($v1_6)->{'value1'})->{'tag'} === "Nil"))) {
$__t4 = ($v1_6)->{'value0'};
goto end_branch_4;;
};
if (($v_5 <= 0)) {
$__t4 = ($v1_6)->{'value0'};
goto end_branch_4;;
};
$__tco_2 = ((($GLOBALS['Data_Ring_ringInt'])['sub'])($v_5))(1);
$__tco_3 = ($v1_6)->{'value1'};
$__tco_var_go__4_1_1_v_5 = $__tco_2;
$__tco_var_go__4_1_1_v1_6 = $__tco_3;
goto tco_loop_go__4_1_1;;
$__t4 = null;
end_branch_4:;
$__t1 = $__t4;
goto end_branch_1;;
};
  if ((is_object($v1_6) && (($v1_6)->{'tag'} === "Nil"))) {
$__t1 = (($foldMap1_1_0)($GLOBALS['Data_Semigroup_Last_Last']))($xs_3);
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
  $__res = (($go__4_1)($i_2))((((((($dictFoldable1_0)['Foldable0'])(null))['foldr'])($GLOBALS['Control_Monad_Gen_Cons']))(new Phpurs_Data0("Nil")))($xs_3));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Gen_oneOf
$GLOBALS['Control_Monad_Gen_oneOf'] = (function() {
  $__fn = function($dictMonadGen_0 = null, $dictFoldable1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $length_2_0 = ((((($dictFoldable1_1)['Foldable0'])(null))['foldl'])((function() {
  $__fn = function($c_2 = null, $v_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_Semiring_semiringInt'])['add'])(($GLOBALS['Data_Semiring_semiringInt'])['one']))($c_2);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(($GLOBALS['Data_Semiring_semiringInt'])['zero']);
  $fromIndex1_3_1 = ($GLOBALS['Control_Monad_Gen_fromIndex'])($dictFoldable1_1);
  $__res = function($xs_4 = null) use ($dictMonadGen_0, $fromIndex1_3_1, $length_2_0) {
  $__num = \func_num_args();
  $__res = ((((((($dictMonadGen_0)['Monad0'])(null))['Bind1'])(null))['bind'])(((($dictMonadGen_0)['chooseInt'])(0))(((($GLOBALS['Data_Ring_ringInt'])['sub'])(($length_2_0)($xs_4)))(1))))(function($n_5 = null) use ($fromIndex1_3_1, $xs_4) {
  $__num = \func_num_args();
  $__res = (($fromIndex1_3_1)($n_5))($xs_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Control_Monad_Gen_freqSemigroup
$GLOBALS['Control_Monad_Gen_freqSemigroup'] = function($v_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = ($v_0)->{'value0'};
  $__local_var_2_1 = ($v_0)->{'value1'};
  $__res = function($pos_3 = null) use ($__local_var_1_0, $__local_var_2_1) {
  $__num = \func_num_args();
  $__t2 = null;;
  if (($pos_3 >= $__local_var_1_0)) {
$__t2 = new Phpurs_Data2("Tuple", new Phpurs_Data1("Just", ((($GLOBALS['Data_Ring_ringNumber'])['sub'])($pos_3))($__local_var_1_0)), $__local_var_2_1);
goto end_branch_2;;
};
  $__t2 = new Phpurs_Data2("Tuple", new Phpurs_Data0("Nothing"), $__local_var_2_1);
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Gen_frequency
$GLOBALS['Control_Monad_Gen_frequency'] = (function() {
  $__fn = function($dictMonadGen_0 = null, $dictFoldable1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $foldMap_2_0 = (((($dictFoldable1_1)['Foldable0'])(null))['foldMap'])($GLOBALS['Control_Monad_Gen_monoidAdditive']);
  $foldMap1_3_1 = (($dictFoldable1_1)['foldMap1'])($GLOBALS['Control_Monad_Gen_semigroupFreqSemigroup']);
  $__res = function($xs_4 = null) use ($dictMonadGen_0, $foldMap1_3_1, $foldMap_2_0) {
  $__num = \func_num_args();
  $__res = ((((((($dictMonadGen_0)['Monad0'])(null))['Bind1'])(null))['bind'])(((($dictMonadGen_0)['chooseFloat'])(0.0))((($foldMap_2_0)($GLOBALS['Data_Tuple_fst']))($xs_4))))(((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])($GLOBALS['Data_Tuple_snd']))((($foldMap1_3_1)($GLOBALS['Control_Monad_Gen_freqSemigroup']))($xs_4)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Control_Monad_Gen_filtered
$GLOBALS['Control_Monad_Gen_filtered'] = (function() {
  $__fn = function($dictMonadRec_0 = null, $dictMonadGen_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_2_0 = (((((((($dictMonadGen_1)['Monad0'])(null))['Bind1'])(null))['Apply0'])(null))['Functor0'])(null);
  $__res = function($gen_3 = null) use ($__local_var_2_0, $dictMonadRec_0) {
  $__num = \func_num_args();
  $__res = ((($dictMonadRec_0)['tailRecM'])(function($v_4 = null) use ($__local_var_2_0, $gen_3) {
  $__num = \func_num_args();
  $__res = ((($__local_var_2_0)['map'])(function($a_5 = null) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((is_object($a_5) && (($a_5)->{'tag'} === "Nothing"))) {
$__t1 = new Phpurs_Data1("Loop", $GLOBALS['Data_Unit_unit']);
goto end_branch_1;;
};
  if ((is_object($a_5) && (($a_5)->{'tag'} === "Just"))) {
$__t1 = new Phpurs_Data1("Done", ($a_5)->{'value0'});
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($gen_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Unit_unit']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Control_Monad_Gen_suchThat
$GLOBALS['Control_Monad_Gen_suchThat'] = (function() {
  $__fn = function($dictMonadRec_0 = null, $dictMonadGen_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $filtered2_2_0 = (($GLOBALS['Control_Monad_Gen_filtered'])($dictMonadRec_0))($dictMonadGen_1);
  $__local_var_3_1 = (((((((($dictMonadGen_1)['Monad0'])(null))['Bind1'])(null))['Apply0'])(null))['Functor0'])(null);
  $__res = (function() use ($__local_var_3_1, $filtered2_2_0) {
  $__fn = function($gen_4 = null, $pred_5 = null) use ($__local_var_3_1, $filtered2_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($filtered2_2_0)(((($__local_var_3_1)['map'])(function($a_6 = null) use ($pred_5) {
  $__num = \func_num_args();
  $__t2 = null;;
  if (($pred_5)($a_6)) {
$__t2 = new Phpurs_Data1("Just", $a_6);
goto end_branch_2;;
};
  $__t2 = new Phpurs_Data0("Nothing");
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($gen_4));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Control_Monad_Gen_elements
$GLOBALS['Control_Monad_Gen_elements'] = function($dictMonadGen_0 = null) {
  $__num = \func_num_args();
  $Monad0_1_0 = (($dictMonadGen_0)['Monad0'])(null);
  $__res = function($dictFoldable1_2 = null) use ($Monad0_1_0, $dictMonadGen_0) {
  $__num = \func_num_args();
  $length_3_1 = ((((($dictFoldable1_2)['Foldable0'])(null))['foldl'])((function() {
  $__fn = function($c_3 = null, $v_4 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_Semiring_semiringInt'])['add'])(($GLOBALS['Data_Semiring_semiringInt'])['one']))($c_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(($GLOBALS['Data_Semiring_semiringInt'])['zero']);
  $fromIndex1_4_2 = ($GLOBALS['Control_Monad_Gen_fromIndex'])($dictFoldable1_2);
  $__res = function($xs_5 = null) use ($Monad0_1_0, $dictMonadGen_0, $fromIndex1_4_2, $length_3_1) {
  $__num = \func_num_args();
  $__res = ((((($Monad0_1_0)['Bind1'])(null))['bind'])(((($dictMonadGen_0)['chooseInt'])(0))(((($GLOBALS['Data_Ring_ringInt'])['sub'])(($length_3_1)($xs_5)))(1))))(function($n_6 = null) use ($Monad0_1_0, $fromIndex1_4_2, $xs_5) {
  $__num = \func_num_args();
  $__res = (((($Monad0_1_0)['Applicative0'])(null))['pure'])((($fromIndex1_4_2)($n_6))($xs_5));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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

// Control_Monad_Gen_choose
$GLOBALS['Control_Monad_Gen_choose'] = function($dictMonadGen_0 = null) {
  $__num = \func_num_args();
  $chooseBool_1_0 = ($dictMonadGen_0)['chooseBool'];
  $__res = (function() use ($chooseBool_1_0, $dictMonadGen_0) {
  $__fn = function($genA_2 = null, $genB_3 = null) use ($chooseBool_1_0, $dictMonadGen_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((((($dictMonadGen_0)['Monad0'])(null))['Bind1'])(null))['bind'])($chooseBool_1_0))(function($v_4 = null) use ($genA_2, $genB_3) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v_4) {
$__t1 = $genA_2;
goto end_branch_1;;
};
  $__t1 = $genB_3;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

