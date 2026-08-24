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
if (!\function_exists(__NAMESPACE__ . '\\phpurs_execute_effect')) {
  function phpurs_execute_effect($val) {
    if (\is_callable($val)) {
      return $val($GLOBALS['Data_Unit_unit']);
    }
    return $val;
  }
}

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };


final class Control_Monad_Gen_Cons { public $tag = 'Cons'; public function __construct(public  $value0, public  $value1) {} }
final class Control_Monad_Gen_Nil { public $tag = 'Nil'; public function __construct() {} }

// Control_Monad_Gen_Cons
$GLOBALS['Control_Monad_Gen_Cons'] = (function() {
  $__fn = function($value0, $value1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Control\Monad\Gen\Control_Monad_Gen_Cons($value0, $value1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Control_Monad_Gen_Nil
$GLOBALS['Control_Monad_Gen_Nil'] = ($GLOBALS['__phpurs_data0_Nil'] ??= new \Control\Monad\Gen\Control_Monad_Gen_Nil());

// Control_Monad_Gen_unfoldable
function majControl_majMonad_majGen_unfoldable($dictMonadRec_0, $dictMonadGen_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majGen_unfoldable';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $Monad0_2_0 = (($dictMonadGen_1)->{'Monad0'})(null);
  $pure_3_1 = ((($Monad0_2_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_4_2 = (($Monad0_2_0)->{'Bind1'})(null);
  $Functor0_5_3 = (((((($Monad0_2_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($dictUnfoldable_6) use ($Bind1_4_2, $Functor0_5_3, $dictMonadGen_1, $dictMonadRec_0, $pure_3_1) {
  $__num = \func_num_args();
  $__res = function($gen_7) use ($Bind1_4_2, $Functor0_5_3, $dictMonadGen_1, $dictMonadRec_0, $dictUnfoldable_6, $pure_3_1) {
  $__num = \func_num_args();
  $__res = ((($Functor0_5_3)->{'map'})((($dictUnfoldable_6)->{'unfoldr'})(function($v_8) {
  $__num = \func_num_args();
  $__t4 = null;;
  if ($v_8 instanceof \Control\Monad\Gen\Control_Monad_Gen_Nil) {
$__t4 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_4;;
};
  if ($v_8 instanceof \Control\Monad\Gen\Control_Monad_Gen_Cons) {
$__t4 = new \Data\Maybe\Data_Maybe_Just(new \Data\Tuple\Data_Tuple_Tuple(($v_8)->{'value0'}, ($v_8)->{'value1'}));
goto end_branch_4;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t4 = null;
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))((($dictMonadGen_1)->{'sized'})((($GLOBALS['Control_Semigroupoid_composeImpl'])((($dictMonadRec_0)->{'tailRecM'})(function($v_8) use ($Bind1_4_2, $gen_7, $pure_3_1) {
  $__num = \func_num_args();
  $__t7 = null;;
  if ((($v_8)->{'value1'} <= 0)) {
$__t7 = ($pure_3_1)(new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Done(($v_8)->{'value0'}));
goto end_branch_7;;
};
  $__local_var_9_5 = ($v_8)->{'value0'};
  $__local_var_10_6 = ($v_8)->{'value1'};
  $__t7 = ((($Bind1_4_2)->{'bind'})($gen_7))(function($x_11) use ($__local_var_10_6, $__local_var_9_5, $pure_3_1) {
  $__num = \func_num_args();
  $__res = ($pure_3_1)(new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Loop(new \Data\Tuple\Data_Tuple_Tuple(new \Control\Monad\Gen\Control_Monad_Gen_Cons($x_11, $__local_var_9_5), ($__local_var_10_6 - 1))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  end_branch_7:;
  $__res = $__t7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(($GLOBALS['Data_Tuple_Tuple'])(new \Control\Monad\Gen\Control_Monad_Gen_Nil()))));
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
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Control_Monad_Gen_unfoldable'] = __NAMESPACE__ . '\\majControl_majMonad_majGen_unfoldable';

// Control_Monad_Gen_semigroupFreqSemigroup
$GLOBALS['Control_Monad_Gen_semigroupFreqSemigroup'] = (object)["append" => function($v_0) {
  $__num = \func_num_args();
  $__res = function($v1_1) use ($v_0) {
  $__num = \func_num_args();
  $__res = function($pos_2) use ($v1_1, $v_0) {
  $__num = \func_num_args();
  $v2_3_0 = ($v_0)($pos_2);
  $__t1 = null;;
  if (($v2_3_0)->{'value0'} instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = ($v1_1)((($v2_3_0)->{'value0'})->{'value0'});
goto end_branch_1;;
};
  $__t1 = $v2_3_0;
  end_branch_1:;
  $__res = $__t1;
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
}];

// Control_Monad_Gen_oneOf
function majControl_majMonad_majGen_onemajOf($dictMonadGen_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majGen_onemajOf';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Bind1_1_0 = (((($dictMonadGen_0)->{'Monad0'})(null))->{'Bind1'})(null);
  $__res = function($dictFoldable1_2) use ($Bind1_1_0, $dictMonadGen_0) {
  $__num = \func_num_args();
  $Foldable0_3_1 = (($dictFoldable1_2)->{'Foldable0'})(null);
  $__res = function($xs_4) use ($Bind1_1_0, $Foldable0_3_1, $dictFoldable1_2, $dictMonadGen_0) {
  $__num = \func_num_args();
  $__res = ((($Bind1_1_0)->{'bind'})(((($dictMonadGen_0)->{'chooseInt'})(0))(((((($Foldable0_3_1)->{'foldl'})(function($c_5) {
  $__num = \func_num_args();
  $__res = function($v_6) use ($c_5) {
  $__num = \func_num_args();
  $__res = (1 + $c_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(0))($xs_4) - 1))))(function($n_5) use ($dictFoldable1_2, $xs_4) {
  $__num = \func_num_args();
  $go__go_6_2 = null;
  $go__go_6_2 = (function() use ($dictFoldable1_2, &$go__go_6_2, $xs_4) {
  $__fn = function(int $v_7, $v1_8 = null) use ($dictFoldable1_2, &$go__go_6_2, $xs_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_6_2_2_v_7 = $v_7;
  $__tco_var_go__go_6_2_2_v1_8 = $v1_8;
  tco_loop_go__go_6_2_2:;
  $v_7 = $__tco_var_go__go_6_2_2_v_7;
  $v1_8 = $__tco_var_go__go_6_2_2_v1_8;
  $__t2 = null;;
  if ($v1_8 instanceof \Control\Monad\Gen\Control_Monad_Gen_Cons) {
$__t5 = null;;
if (($v1_8)->{'value1'} instanceof \Control\Monad\Gen\Control_Monad_Gen_Nil) {
$__t5 = ($v1_8)->{'value0'};
goto end_branch_5;;
};
if (($v_7 <= 0)) {
$__t5 = ($v1_8)->{'value0'};
goto end_branch_5;;
};
$__tco_3 = ($v_7 - 1);
$__tco_4 = ($v1_8)->{'value1'};
$__tco_var_go__go_6_2_2_v_7 = $__tco_3;
$__tco_var_go__go_6_2_2_v1_8 = $__tco_4;
goto tco_loop_go__go_6_2_2;;
$__t5 = null;
end_branch_5:;
$__t2 = $__t5;
goto end_branch_2;;
};
  if ($v1_8 instanceof \Control\Monad\Gen\Control_Monad_Gen_Nil) {
$__t2 = (((($dictFoldable1_2)->{'foldMap1'})($GLOBALS['Data_Semigroup_Last_semigroupLast']))(function($x_9) {
  $__num = \func_num_args();
  $__res = $x_9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($xs_4);
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
  $__res = (($go__go_6_2)($n_5))((((((($dictFoldable1_2)->{'Foldable0'})(null))->{'foldr'})($GLOBALS['Control_Monad_Gen_Cons']))(new \Control\Monad\Gen\Control_Monad_Gen_Nil()))($xs_4));
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Gen_oneOf'] = __NAMESPACE__ . '\\majControl_majMonad_majGen_onemajOf';

// Control_Monad_Gen_freqSemigroup
function majControl_majMonad_majGen_freqmajSemigroup($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majGen_freqmajSemigroup';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = ($v_0)->{'value0'};
  $__local_var_2_1 = ($v_0)->{'value1'};
  $__res = function($pos_3) use ($__local_var_1_0, $__local_var_2_1) {
  $__num = \func_num_args();
  $__t2 = null;;
  if (( ! \Data\Ord\majData_majOrd_ordmajNumbermajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), $pos_3, $__local_var_1_0) instanceof \Data\Ordering\Data_Ordering_LT)) {
$__t2 = new \Data\Tuple\Data_Tuple_Tuple(new \Data\Maybe\Data_Maybe_Just(($pos_3 - $__local_var_1_0)), $__local_var_2_1);
goto end_branch_2;;
};
  $__t2 = new \Data\Tuple\Data_Tuple_Tuple(new \Data\Maybe\Data_Maybe_Nothing(), $__local_var_2_1);
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Gen_freqSemigroup'] = __NAMESPACE__ . '\\majControl_majMonad_majGen_freqmajSemigroup';

// Control_Monad_Gen_frequency
function majControl_majMonad_majGen_frequency($dictMonadGen_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majGen_frequency';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Bind1_1_0 = (((($dictMonadGen_0)->{'Monad0'})(null))->{'Bind1'})(null);
  $__res = function($dictFoldable1_2) use ($Bind1_1_0, $dictMonadGen_0) {
  $__num = \func_num_args();
  $semigroupAdditive1_3_1 = (object)["append" => function($v_3) {
  $__num = \func_num_args();
  $__res = function($v1_4) use ($v_3) {
  $__num = \func_num_args();
  $__res = ($v_3 + $v1_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $foldMap_3_1 = (((($dictFoldable1_2)->{'Foldable0'})(null))->{'foldMap'})((object)["mempty" => 0.0, "Semigroup0" => function($_dollar___unused_4) use ($semigroupAdditive1_3_1) {
  $__num = \func_num_args();
  $__res = $semigroupAdditive1_3_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]);
  $__res = function($xs_4) use ($Bind1_1_0, $dictFoldable1_2, $dictMonadGen_0, $foldMap_3_1) {
  $__num = \func_num_args();
  $__res = ((($Bind1_1_0)->{'bind'})(((($dictMonadGen_0)->{'chooseFloat'})(0.0))((($foldMap_3_1)($GLOBALS['Data_Tuple_fst']))($xs_4))))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Tuple_snd']))((((($dictFoldable1_2)->{'foldMap1'})($GLOBALS['Control_Monad_Gen_semigroupFreqSemigroup']))($GLOBALS['Control_Monad_Gen_freqSemigroup']))($xs_4)));
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Gen_frequency'] = __NAMESPACE__ . '\\majControl_majMonad_majGen_frequency';

// Control_Monad_Gen_filtered
function majControl_majMonad_majGen_filtered($dictMonadRec_0, $dictMonadGen_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majGen_filtered';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $Functor0_2_0 = (((((((($dictMonadGen_1)->{'Monad0'})(null))->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($gen_3) use ($Functor0_2_0, $dictMonadRec_0) {
  $__num = \func_num_args();
  $__res = ((($dictMonadRec_0)->{'tailRecM'})(function($v_4) use ($Functor0_2_0, $gen_3) {
  $__num = \func_num_args();
  $__res = ((($Functor0_2_0)->{'map'})(function($a_5) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($a_5 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t1 = new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Loop($GLOBALS['Data_Unit_unit']);
goto end_branch_1;;
};
  if ($a_5 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Done(($a_5)->{'value0'});
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
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Control_Monad_Gen_filtered'] = __NAMESPACE__ . '\\majControl_majMonad_majGen_filtered';

// Control_Monad_Gen_suchThat
function majControl_majMonad_majGen_suchmajThat($dictMonadRec_0, $dictMonadGen_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majGen_suchmajThat';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $Functor0_2_0 = (((((((($dictMonadGen_1)->{'Monad0'})(null))->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $filtered2_2_0 = function($gen_3) use ($Functor0_2_0, $dictMonadRec_0) {
  $__num = \func_num_args();
  $__res = ((($dictMonadRec_0)->{'tailRecM'})(function($v_4) use ($Functor0_2_0, $gen_3) {
  $__num = \func_num_args();
  $__res = ((($Functor0_2_0)->{'map'})(function($a_5) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($a_5 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t1 = new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Loop($GLOBALS['Data_Unit_unit']);
goto end_branch_1;;
};
  if ($a_5 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Done(($a_5)->{'value0'});
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
  $Functor0_3_3 = (((((((($dictMonadGen_1)->{'Monad0'})(null))->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($gen_4) use ($Functor0_3_3, $filtered2_2_0) {
  $__num = \func_num_args();
  $__res = function($pred_5) use ($Functor0_3_3, $filtered2_2_0, $gen_4) {
  $__num = \func_num_args();
  $__res = ($filtered2_2_0)(((($Functor0_3_3)->{'map'})(function($a_6) use ($pred_5) {
  $__num = \func_num_args();
  $__t4 = null;;
  if (($pred_5)($a_6)) {
$__t4 = new \Data\Maybe\Data_Maybe_Just($a_6);
goto end_branch_4;;
};
  $__t4 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($gen_4));
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
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Control_Monad_Gen_suchThat'] = __NAMESPACE__ . '\\majControl_majMonad_majGen_suchmajThat';

// Control_Monad_Gen_elements
function majControl_majMonad_majGen_elements($dictMonadGen_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majGen_elements';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Monad0_1_0 = (($dictMonadGen_0)->{'Monad0'})(null);
  $Bind1_2_1 = (($Monad0_1_0)->{'Bind1'})(null);
  $pure_3_2 = ((($Monad0_1_0)->{'Applicative0'})(null))->{'pure'};
  $__res = function($dictFoldable1_4) use ($Bind1_2_1, $dictMonadGen_0, $pure_3_2) {
  $__num = \func_num_args();
  $Foldable0_5_3 = (($dictFoldable1_4)->{'Foldable0'})(null);
  $__res = function($xs_6) use ($Bind1_2_1, $Foldable0_5_3, $dictFoldable1_4, $dictMonadGen_0, $pure_3_2) {
  $__num = \func_num_args();
  $__res = ((($Bind1_2_1)->{'bind'})(((($dictMonadGen_0)->{'chooseInt'})(0))(((((($Foldable0_5_3)->{'foldl'})(function($c_7) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($c_7) {
  $__num = \func_num_args();
  $__res = (1 + $c_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(0))($xs_6) - 1))))(function($n_7) use ($dictFoldable1_4, $pure_3_2, $xs_6) {
  $__num = \func_num_args();
  $go__go_8_4 = null;
  $go__go_8_4 = (function() use ($dictFoldable1_4, &$go__go_8_4, $xs_6) {
  $__fn = function(int $v_9, $v1_10 = null) use ($dictFoldable1_4, &$go__go_8_4, $xs_6, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_8_4_4_v_9 = $v_9;
  $__tco_var_go__go_8_4_4_v1_10 = $v1_10;
  tco_loop_go__go_8_4_4:;
  $v_9 = $__tco_var_go__go_8_4_4_v_9;
  $v1_10 = $__tco_var_go__go_8_4_4_v1_10;
  $__t4 = null;;
  if ($v1_10 instanceof \Control\Monad\Gen\Control_Monad_Gen_Cons) {
$__t7 = null;;
if (($v1_10)->{'value1'} instanceof \Control\Monad\Gen\Control_Monad_Gen_Nil) {
$__t7 = ($v1_10)->{'value0'};
goto end_branch_7;;
};
if (($v_9 <= 0)) {
$__t7 = ($v1_10)->{'value0'};
goto end_branch_7;;
};
$__tco_5 = ($v_9 - 1);
$__tco_6 = ($v1_10)->{'value1'};
$__tco_var_go__go_8_4_4_v_9 = $__tco_5;
$__tco_var_go__go_8_4_4_v1_10 = $__tco_6;
goto tco_loop_go__go_8_4_4;;
$__t7 = null;
end_branch_7:;
$__t4 = $__t7;
goto end_branch_4;;
};
  if ($v1_10 instanceof \Control\Monad\Gen\Control_Monad_Gen_Nil) {
$__t4 = (((($dictFoldable1_4)->{'foldMap1'})($GLOBALS['Data_Semigroup_Last_semigroupLast']))(function($x_11) {
  $__num = \func_num_args();
  $__res = $x_11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($xs_6);
goto end_branch_4;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t4 = null;
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ($pure_3_2)((($go__go_8_4)($n_7))((((((($dictFoldable1_4)->{'Foldable0'})(null))->{'foldr'})($GLOBALS['Control_Monad_Gen_Cons']))(new \Control\Monad\Gen\Control_Monad_Gen_Nil()))($xs_6)));
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Gen_elements'] = __NAMESPACE__ . '\\majControl_majMonad_majGen_elements';

// Control_Monad_Gen_choose
function majControl_majMonad_majGen_choose($dictMonadGen_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majGen_choose';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Bind1_1_0 = (((($dictMonadGen_0)->{'Monad0'})(null))->{'Bind1'})(null);
  $chooseBool_2_1 = ($dictMonadGen_0)->{'chooseBool'};
  $__res = function($genA_3) use ($Bind1_1_0, $chooseBool_2_1) {
  $__num = \func_num_args();
  $__res = function($genB_4) use ($Bind1_1_0, $chooseBool_2_1, $genA_3) {
  $__num = \func_num_args();
  $__res = ((($Bind1_1_0)->{'bind'})($chooseBool_2_1))(function($v_5) use ($genA_3, $genB_4) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ($v_5) {
$__t2 = $genA_3;
goto end_branch_2;;
};
  $__t2 = $genB_4;
  end_branch_2:;
  $__res = $__t2;
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Gen_choose'] = __NAMESPACE__ . '\\majControl_majMonad_majGen_choose';

