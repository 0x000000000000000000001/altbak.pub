<?php

namespace Control\Monad\Gen\Common;

// ALL IMPORTS: Control.Applicative, Control.Apply, Control.Bind, Control.Monad.Gen, Control.Monad.Gen.Class, Control.Monad.Gen.Common, Control.Monad.Rec.Class, Control.Semigroupoid, Data.Either, Data.Functor, Data.Identity, Data.Maybe, Data.NonEmpty, Data.Ord, Data.Ring, Data.Tuple, Data.Unfoldable, Prelude, Prim
// TO REQUIRE: Control.Applicative, Control.Apply, Control.Bind, Control.Monad.Gen, Control.Monad.Gen.Class, Control.Monad.Gen.Common, Control.Monad.Rec.Class, Control.Semigroupoid, Data.Either, Data.Functor, Data.Identity, Data.Maybe, Data.NonEmpty, Data.Ord, Data.Ring, Data.Tuple, Data.Unfoldable, Prelude
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Monad.Gen/index.php';
require_once __DIR__ . '/../Control.Monad.Gen.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Gen.Common/index.php';
require_once __DIR__ . '/../Control.Monad.Rec.Class/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Either/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Identity/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.NonEmpty/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ring/index.php';
require_once __DIR__ . '/../Data.Tuple/index.php';
require_once __DIR__ . '/../Data.Unfoldable/index.php';
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




// Control_Monad_Gen_Common_genTuple
function majControl_majMonad_majGen_majCommon_genmajTuple($dictApply_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majGen_majCommon_genmajTuple';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Functor0_1_0 = (($dictApply_0)->{'Functor0'})(null);
  $__res = function($a_2) use ($Functor0_1_0, $dictApply_0) {
  $__num = \func_num_args();
  $__res = function($b_3) use ($Functor0_1_0, $a_2, $dictApply_0) {
  $__num = \func_num_args();
  $__res = ((($dictApply_0)->{'apply'})(((($Functor0_1_0)->{'map'})($GLOBALS['Data_Tuple_Tuple']))($a_2)))($b_3);
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
$GLOBALS['Control_Monad_Gen_Common_genTuple'] = __NAMESPACE__ . '\\majControl_majMonad_majGen_majCommon_genmajTuple';

// Control_Monad_Gen_Common_genNonEmpty
function majControl_majMonad_majGen_majCommon_genmajNonmajEmpty($dictMonadRec_0, $dictMonadGen_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majGen_majCommon_genmajNonmajEmpty';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $Bind1_2_0 = (((($dictMonadGen_1)->{'Monad0'})(null))->{'Bind1'})(null);
  $Apply0_3_1 = (($Bind1_2_0)->{'Apply0'})(null);
  $Functor0_4_2 = (((($Bind1_2_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($dictUnfoldable_5) use ($Apply0_3_1, $Functor0_4_2, $dictMonadGen_1, $dictMonadRec_0) {
  $__num = \func_num_args();
  $__res = function($gen_6) use ($Apply0_3_1, $Functor0_4_2, $dictMonadGen_1, $dictMonadRec_0, $dictUnfoldable_5) {
  $__num = \func_num_args();
  $Monad0_7_5 = (($dictMonadGen_1)->{'Monad0'})(null);
  $pure_8_6 = ((($Monad0_7_5)->{'Applicative0'})(null))->{'pure'};
  $Bind1_9_7 = (($Monad0_7_5)->{'Bind1'})(null);
  $__res = ((($Apply0_3_1)->{'apply'})(((($Functor0_4_2)->{'map'})($GLOBALS['Data_NonEmpty_NonEmpty']))($gen_6)))(((($dictMonadGen_1)->{'resize'})((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($y_7) {
  $__num = \func_num_args();
  $v_8_3 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), 0, $y_7);
  $__t4 = null;;
  if ($v_8_3 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t4 = $y_7;
goto end_branch_4;;
};
  if ($v_8_3 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t4 = 0;
goto end_branch_4;;
};
  if ($v_8_3 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t4 = 0;
goto end_branch_4;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t4 = null;
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_7) {
  $__num = \func_num_args();
  $__res = ($v_7 - 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(((((((((($Monad0_7_5)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})((($dictUnfoldable_5)->{'unfoldr'})(function($v_10) {
  $__num = \func_num_args();
  $__t8 = null;;
  if ($v_10 instanceof \Control\Monad\Gen\Control_Monad_Gen_Nil) {
$__t8 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_8;;
};
  if ($v_10 instanceof \Control\Monad\Gen\Control_Monad_Gen_Cons) {
$__t8 = new \Data\Maybe\Data_Maybe_Just(new \Data\Tuple\Data_Tuple_Tuple(($v_10)->{'value0'}, ($v_10)->{'value1'}));
goto end_branch_8;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t8 = null;
  end_branch_8:;
  $__res = $__t8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))((($dictMonadGen_1)->{'sized'})((($GLOBALS['Control_Semigroupoid_composeImpl'])((($dictMonadRec_0)->{'tailRecM'})(function($v_10) use ($Bind1_9_7, $gen_6, $pure_8_6) {
  $__num = \func_num_args();
  $__t11 = null;;
  if ((($v_10)->{'value1'} <= 0)) {
$__t11 = ($pure_8_6)(new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Done(($v_10)->{'value0'}));
goto end_branch_11;;
};
  $__local_var_11_9 = ($v_10)->{'value0'};
  $__local_var_12_10 = ($v_10)->{'value1'};
  $__t11 = ((($Bind1_9_7)->{'bind'})($gen_6))(function($x_13) use ($__local_var_11_9, $__local_var_12_10, $pure_8_6) {
  $__num = \func_num_args();
  $__res = ($pure_8_6)(new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Loop(new \Data\Tuple\Data_Tuple_Tuple(new \Control\Monad\Gen\Control_Monad_Gen_Cons($x_13, $__local_var_11_9), ($__local_var_12_10 - 1))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  end_branch_11:;
  $__res = $__t11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(($GLOBALS['Data_Tuple_Tuple'])(new \Control\Monad\Gen\Control_Monad_Gen_Nil()))))));
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
$GLOBALS['Control_Monad_Gen_Common_genNonEmpty'] = __NAMESPACE__ . '\\majControl_majMonad_majGen_majCommon_genmajNonmajEmpty';

// Control_Monad_Gen_Common_genMaybe'
function majControl_majMonad_majGen_majCommon_genmajMaybe__prime__($dictMonadGen_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majGen_majCommon_genmajMaybe__prime__';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Monad0_1_0 = (($dictMonadGen_0)->{'Monad0'})(null);
  $Bind1_2_1 = (($Monad0_1_0)->{'Bind1'})(null);
  $Functor0_3_2 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $Applicative0_4_3 = (($Monad0_1_0)->{'Applicative0'})(null);
  $__res = function($bias_5) use ($Applicative0_4_3, $Bind1_2_1, $Functor0_3_2, $dictMonadGen_0) {
  $__num = \func_num_args();
  $__res = function($gen_6) use ($Applicative0_4_3, $Bind1_2_1, $Functor0_3_2, $bias_5, $dictMonadGen_0) {
  $__num = \func_num_args();
  $__res = ((($Bind1_2_1)->{'bind'})(((($dictMonadGen_0)->{'chooseFloat'})(0.0))(1.0)))(function($n_7) use ($Applicative0_4_3, $Functor0_3_2, $bias_5, $gen_6) {
  $__num = \func_num_args();
  $__t4 = null;;
  if (\Data\Ord\majData_majOrd_ordmajNumbermajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), $n_7, $bias_5) instanceof \Data\Ordering\Data_Ordering_LT) {
$__t4 = ((($Functor0_3_2)->{'map'})($GLOBALS['Data_Maybe_Just']))($gen_6);
goto end_branch_4;;
};
  $__t4 = (($Applicative0_4_3)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
  end_branch_4:;
  $__res = $__t4;
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
$GLOBALS['Control_Monad_Gen_Common_genMaybe__prime__'] = __NAMESPACE__ . '\\majControl_majMonad_majGen_majCommon_genmajMaybe__prime__';

// Control_Monad_Gen_Common_genMaybe
function majControl_majMonad_majGen_majCommon_genmajMaybe($dictMonadGen_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majGen_majCommon_genmajMaybe';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Monad0_1_0 = (($dictMonadGen_0)->{'Monad0'})(null);
  $Bind1_2_1 = (($Monad0_1_0)->{'Bind1'})(null);
  $Functor0_3_2 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $Applicative0_4_3 = (($Monad0_1_0)->{'Applicative0'})(null);
  $__res = function($gen_5) use ($Applicative0_4_3, $Bind1_2_1, $Functor0_3_2, $dictMonadGen_0) {
  $__num = \func_num_args();
  $__res = ((($Bind1_2_1)->{'bind'})(((($dictMonadGen_0)->{'chooseFloat'})(0.0))(1.0)))(function($n_6) use ($Applicative0_4_3, $Functor0_3_2, $gen_5) {
  $__num = \func_num_args();
  $__t4 = null;;
  if (\Data\Ord\majData_majOrd_ordmajNumbermajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), $n_6, 0.75) instanceof \Data\Ordering\Data_Ordering_LT) {
$__t4 = ((($Functor0_3_2)->{'map'})($GLOBALS['Data_Maybe_Just']))($gen_5);
goto end_branch_4;;
};
  $__t4 = (($Applicative0_4_3)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
  end_branch_4:;
  $__res = $__t4;
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Gen_Common_genMaybe'] = __NAMESPACE__ . '\\majControl_majMonad_majGen_majCommon_genmajMaybe';

// Control_Monad_Gen_Common_genIdentity
function majControl_majMonad_majGen_majCommon_genmajIdentity($dictFunctor_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majGen_majCommon_genmajIdentity';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($dictFunctor_0)->{'map'})(function($x_1) {
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
$GLOBALS['Control_Monad_Gen_Common_genIdentity'] = __NAMESPACE__ . '\\majControl_majMonad_majGen_majCommon_genmajIdentity';

// Control_Monad_Gen_Common_genEither'
function majControl_majMonad_majGen_majCommon_genmajEither__prime__($dictMonadGen_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majGen_majCommon_genmajEither__prime__';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Monad0_1_0 = (($dictMonadGen_0)->{'Monad0'})(null);
  $Bind1_2_1 = (($Monad0_1_0)->{'Bind1'})(null);
  $Functor0_3_2 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($bias_4) use ($Bind1_2_1, $Functor0_3_2, $dictMonadGen_0) {
  $__num = \func_num_args();
  $__res = function($genA_5) use ($Bind1_2_1, $Functor0_3_2, $bias_4, $dictMonadGen_0) {
  $__num = \func_num_args();
  $__res = function($genB_6) use ($Bind1_2_1, $Functor0_3_2, $bias_4, $dictMonadGen_0, $genA_5) {
  $__num = \func_num_args();
  $__res = ((($Bind1_2_1)->{'bind'})(((($dictMonadGen_0)->{'chooseFloat'})(0.0))(1.0)))(function($n_7) use ($Functor0_3_2, $bias_4, $genA_5, $genB_6) {
  $__num = \func_num_args();
  $__t3 = null;;
  if (\Data\Ord\majData_majOrd_ordmajNumbermajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), $n_7, $bias_4) instanceof \Data\Ordering\Data_Ordering_LT) {
$__t3 = ((($Functor0_3_2)->{'map'})($GLOBALS['Data_Either_Left']))($genA_5);
goto end_branch_3;;
};
  $__t3 = ((($Functor0_3_2)->{'map'})($GLOBALS['Data_Either_Right']))($genB_6);
  end_branch_3:;
  $__res = $__t3;
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
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Gen_Common_genEither__prime__'] = __NAMESPACE__ . '\\majControl_majMonad_majGen_majCommon_genmajEither__prime__';

// Control_Monad_Gen_Common_genEither
function majControl_majMonad_majGen_majCommon_genmajEither($dictMonadGen_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majGen_majCommon_genmajEither';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Monad0_1_0 = (($dictMonadGen_0)->{'Monad0'})(null);
  $Bind1_2_1 = (($Monad0_1_0)->{'Bind1'})(null);
  $Functor0_3_2 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($genA_4) use ($Bind1_2_1, $Functor0_3_2, $dictMonadGen_0) {
  $__num = \func_num_args();
  $__res = function($genB_5) use ($Bind1_2_1, $Functor0_3_2, $dictMonadGen_0, $genA_4) {
  $__num = \func_num_args();
  $__res = ((($Bind1_2_1)->{'bind'})(((($dictMonadGen_0)->{'chooseFloat'})(0.0))(1.0)))(function($n_6) use ($Functor0_3_2, $genA_4, $genB_5) {
  $__num = \func_num_args();
  $__t3 = null;;
  if (\Data\Ord\majData_majOrd_ordmajNumbermajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), $n_6, 0.5) instanceof \Data\Ordering\Data_Ordering_LT) {
$__t3 = ((($Functor0_3_2)->{'map'})($GLOBALS['Data_Either_Left']))($genA_4);
goto end_branch_3;;
};
  $__t3 = ((($Functor0_3_2)->{'map'})($GLOBALS['Data_Either_Right']))($genB_5);
  end_branch_3:;
  $__res = $__t3;
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
$GLOBALS['Control_Monad_Gen_Common_genEither'] = __NAMESPACE__ . '\\majControl_majMonad_majGen_majCommon_genmajEither';

