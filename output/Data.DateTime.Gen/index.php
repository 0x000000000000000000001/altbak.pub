<?php

namespace Data\DateTime\Gen;

// ALL IMPORTS: Control.Apply, Control.Monad.Gen, Data.Date.Gen, Data.DateTime, Data.Functor, Data.Time.Gen, Prelude, Prim
// TO REQUIRE: Control.Apply, Control.Monad.Gen, Data.Date.Gen, Data.DateTime, Data.Functor, Data.Time.Gen, Prelude
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Monad.Gen/index.php';
require_once __DIR__ . '/../Data.Date.Gen/index.php';
require_once __DIR__ . '/../Data.DateTime/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Time.Gen/index.php';
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




// Data_DateTime_Gen_genDateTime
function majData_majDatemajTime_majGen_genmajDatemajTime($dictMonadGen_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majDatemajTime_majGen_genmajDatemajTime';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Bind1_1_0 = (((($dictMonadGen_0)->{'Monad0'})(null))->{'Bind1'})(null);
  $Monad0_2_1 = (($dictMonadGen_0)->{'Monad0'})(null);
  $Bind1_3_2 = (($Monad0_2_1)->{'Bind1'})(null);
  $Functor0_4_3 = (((((($Monad0_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $pure_5_4 = ((($Monad0_2_1)->{'Applicative0'})(null))->{'pure'};
  $Bind1_2_12 = (((($dictMonadGen_0)->{'Monad0'})(null))->{'Bind1'})(null);
  $Apply0_3_13 = (($Bind1_2_12)->{'Apply0'})(null);
  $Applicative0_4_14 = (((($dictMonadGen_0)->{'Monad0'})(null))->{'Applicative0'})(null);
  $v_5_15 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_5) {
  $__num = \func_num_args();
  $__t15 = null;;
  if ((($n_5 >= 0) && ($n_5 <= 23))) {
$__t15 = new \Data\Maybe\Data_Maybe_Just($n_5);
goto end_branch_15;;
};
  $__t15 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_15:;
  $__res = $__t15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_5) {
  $__num = \func_num_args();
  $__res = ($v_5 + 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_5) {
  $__num = \func_num_args();
  $__res = $v_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), 0);
  $__t17 = null;;
  if ($v_5_15 instanceof \Data\Maybe\Data_Maybe_Just) {
$Monad0_6_18 = (($dictMonadGen_0)->{'Monad0'})(null);
$pure_7_19 = ((($Monad0_6_18)->{'Applicative0'})(null))->{'pure'};
$foldableNonEmpty1_8_20 = (object)["foldMap" => function($dictMonoid_8) {
  $__num = \func_num_args();
  $Semigroup0_9_20 = (($dictMonoid_8)->{'Semigroup0'})(null);
  $__res = function($f_10) use ($Semigroup0_9_20, $dictMonoid_8) {
  $__num = \func_num_args();
  $__res = function($v_11) use ($Semigroup0_9_20, $dictMonoid_8, $f_10) {
  $__num = \func_num_args();
  $Semigroup0_12_21 = (($dictMonoid_8)->{'Semigroup0'})(null);
  $__res = ((($Semigroup0_9_20)->{'append'})(($f_10)(($v_11)->{'value0'})))(\Data\Foldable\majData_majFoldable_foldrmajArray(function($x_13) use ($Semigroup0_12_21, $f_10) {
  $__num = \func_num_args();
  $__res = function($acc_14) use ($Semigroup0_12_21, $f_10, $x_13) {
  $__num = \func_num_args();
  $__res = ((($Semigroup0_12_21)->{'append'})(($f_10)($x_13)))($acc_14);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($dictMonoid_8)->{'mempty'}, ($v_11)->{'value1'}));
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
}, "foldl" => function($f_8) {
  $__num = \func_num_args();
  $__res = function($b_9) use ($f_8) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($b_9, $f_8) {
  $__num = \func_num_args();
  $__res = \Data\Foldable\majData_majFoldable_foldlmajArray($f_8, (($f_8)($b_9))(($v_10)->{'value0'}), ($v_10)->{'value1'});
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
}, "foldr" => function($f_8) {
  $__num = \func_num_args();
  $__res = function($b_9) use ($f_8) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($b_9, $f_8) {
  $__num = \func_num_args();
  $__res = (($f_8)(($v_10)->{'value0'}))(\Data\Foldable\majData_majFoldable_foldrmajArray($f_8, $b_9, ($v_10)->{'value1'}));
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
$__local_var_8_20 = (object)["foldMap1" => function($dictSemigroup_9) {
  $__num = \func_num_args();
  $__res = function($f_10) use ($dictSemigroup_9) {
  $__num = \func_num_args();
  $__res = function($v_11) use ($dictSemigroup_9, $f_10) {
  $__num = \func_num_args();
  $__res = \Data\Foldable\majData_majFoldable_foldlmajArray(function($s_12) use ($dictSemigroup_9, $f_10) {
  $__num = \func_num_args();
  $__res = function($a1_13) use ($dictSemigroup_9, $f_10, $s_12) {
  $__num = \func_num_args();
  $__res = ((($dictSemigroup_9)->{'append'})($s_12))(($f_10)($a1_13));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($f_10)(($v_11)->{'value0'}), ($v_11)->{'value1'});
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
}, "foldr1" => function($f_9) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($f_9) {
  $__num = \func_num_args();
  $__local_var_11_23 = ($f_9)(($v_10)->{'value0'});
  $__local_var_12_24 = \Data\Foldable\majData_majFoldable_foldrmajArray(function($a1_12) use ($f_9) {
  $__num = \func_num_args();
  $__local_var_13_24 = ($f_9)($a1_12);
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Maybe_Just']))(function($v2_14) use ($__local_var_13_24, $a1_12) {
  $__num = \func_num_args();
  $__t25 = null;;
  if ($v2_14 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t25 = $a1_12;
goto end_branch_25;;
};
  if ($v2_14 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t25 = ($__local_var_13_24)(($v2_14)->{'value0'});
goto end_branch_25;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t25 = null;
  end_branch_25:;
  $__res = $__t25;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, new \Data\Maybe\Data_Maybe_Nothing(), ($v_10)->{'value1'});
  $__t27 = null;;
  if ($__local_var_12_24 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t27 = ($v_10)->{'value0'};
goto end_branch_27;;
};
  if ($__local_var_12_24 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t27 = ($__local_var_11_23)(($__local_var_12_24)->{'value0'});
goto end_branch_27;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t27 = null;
  end_branch_27:;
  $__res = $__t27;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldl1" => function($f_9) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($f_9) {
  $__num = \func_num_args();
  $__res = \Data\Foldable\majData_majFoldable_foldlmajArray($f_9, ($v_10)->{'value0'}, ($v_10)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable0" => function($_dollar___unused_9) use ($foldableNonEmpty1_8_20) {
  $__num = \func_num_args();
  $__res = $foldableNonEmpty1_8_20;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
$__t34 = null;;
switch (($v_5_15)->{'value0'}) {
case 23:
$__t34 = \Data\Unfoldable1\majData_majUnfoldable1_unfoldr1majArraymajImpl($GLOBALS['Data_Maybe_isNothing'], function($v_9) {
  $__num = \func_num_args();
  $__t35 = null;;
  if ($v_9 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t35 = ($v_9)->{'value0'};
goto end_branch_35;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t35 = null;
  end_branch_35:;
  $__res = $__t35;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $GLOBALS['Data_Tuple_fst'], $GLOBALS['Data_Tuple_snd'], function($i_9) use ($v_5_15) {
  $__num = \func_num_args();
  $__t36 = null;;
  if (($i_9 <= 0)) {
$__t36 = new \Data\Tuple\Data_Tuple_Tuple(($v_5_15)->{'value0'}, new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_36;;
};
  $__t36 = new \Data\Tuple\Data_Tuple_Tuple(($v_5_15)->{'value0'}, new \Data\Maybe\Data_Maybe_Just(($i_9 - 1)));
  end_branch_36:;
  $__res = $__t36;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, 0);
goto end_branch_34;;
break;
default:
;
break;
};
if ((($v_5_15)->{'value0'} < 23)) {
$__t34 = \Data\Unfoldable1\majData_majUnfoldable1_unfoldr1majArraymajImpl($GLOBALS['Data_Maybe_isNothing'], function($v_9) {
  $__num = \func_num_args();
  $__t37 = null;;
  if ($v_9 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t37 = ($v_9)->{'value0'};
goto end_branch_37;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t37 = null;
  end_branch_37:;
  $__res = $__t37;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $GLOBALS['Data_Tuple_fst'], $GLOBALS['Data_Tuple_snd'], function($a_9) {
  $__num = \func_num_args();
  $__local_var_10_38 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_10) {
  $__num = \func_num_args();
  $__t38 = null;;
  if ((($n_10 >= 0) && ($n_10 <= 23))) {
$__t38 = new \Data\Maybe\Data_Maybe_Just($n_10);
goto end_branch_38;;
};
  $__t38 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_38:;
  $__res = $__t38;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_10) {
  $__num = \func_num_args();
  $__res = ($v_10 + 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_10) {
  $__num = \func_num_args();
  $__res = $v_10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), $a_9);
  $__t40 = null;;
  if ($__local_var_10_38 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t41 = null;;
if ((($__local_var_10_38)->{'value0'} <= 23)) {
$__t41 = new \Data\Maybe\Data_Maybe_Just(($__local_var_10_38)->{'value0'});
goto end_branch_41;;
};
$__t41 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_41:;
$__t40 = $__t41;
goto end_branch_40;;
};
  if ($__local_var_10_38 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t40 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_40;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t40 = null;
  end_branch_40:;
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_9, $__t40);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($v_5_15)->{'value0'});
goto end_branch_34;;
};
$__t34 = \Data\Unfoldable1\majData_majUnfoldable1_unfoldr1majArraymajImpl($GLOBALS['Data_Maybe_isNothing'], function($v_9) {
  $__num = \func_num_args();
  $__t29 = null;;
  if ($v_9 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t29 = ($v_9)->{'value0'};
goto end_branch_29;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t29 = null;
  end_branch_29:;
  $__res = $__t29;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $GLOBALS['Data_Tuple_fst'], $GLOBALS['Data_Tuple_snd'], function($a_9) {
  $__num = \func_num_args();
  $__local_var_10_30 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_10) {
  $__num = \func_num_args();
  $__t30 = null;;
  if ((($n_10 >= 0) && ($n_10 <= 23))) {
$__t30 = new \Data\Maybe\Data_Maybe_Just($n_10);
goto end_branch_30;;
};
  $__t30 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_30:;
  $__res = $__t30;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_10) {
  $__num = \func_num_args();
  $__res = ($v_10 - 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_10) {
  $__num = \func_num_args();
  $__res = $v_10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), $a_9);
  $__t32 = null;;
  if ($__local_var_10_30 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t33 = null;;
if ((($__local_var_10_30)->{'value0'} >= 23)) {
$__t33 = new \Data\Maybe\Data_Maybe_Just(($__local_var_10_30)->{'value0'});
goto end_branch_33;;
};
$__t33 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_33:;
$__t32 = $__t33;
goto end_branch_32;;
};
  if ($__local_var_10_30 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t32 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_32;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t32 = null;
  end_branch_32:;
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_9, $__t32);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($v_5_15)->{'value0'});
end_branch_34:;
$__local_var_9_29 = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(0, $__t34);
$__t17 = ((((($Monad0_6_18)->{'Bind1'})(null))->{'bind'})(((($dictMonadGen_0)->{'chooseInt'})(0))(((((((($__local_var_8_20)->{'Foldable0'})(null))->{'foldl'})(function($c_10) {
  $__num = \func_num_args();
  $__res = function($v_11) use ($c_10) {
  $__num = \func_num_args();
  $__res = (1 + $c_10);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(0))($__local_var_9_29) - 1))))(function($n_10) use ($__local_var_8_20, $__local_var_9_29, $pure_7_19) {
  $__num = \func_num_args();
  $go__go_11_43 = null;
  $go__go_11_43 = (function() use ($__local_var_8_20, $__local_var_9_29, &$go__go_11_43) {
  $__fn = function(int $v_12, $v1_13 = null) use ($__local_var_8_20, $__local_var_9_29, &$go__go_11_43, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_11_43_43_v_12 = $v_12;
  $__tco_var_go__go_11_43_43_v1_13 = $v1_13;
  tco_loop_go__go_11_43_43:;
  $v_12 = $__tco_var_go__go_11_43_43_v_12;
  $v1_13 = $__tco_var_go__go_11_43_43_v1_13;
  $__t43 = null;;
  if ($v1_13 instanceof \Control\Monad\Gen\Control_Monad_Gen_Cons) {
$__t46 = null;;
if (($v1_13)->{'value1'} instanceof \Control\Monad\Gen\Control_Monad_Gen_Nil) {
$__t46 = ($v1_13)->{'value0'};
goto end_branch_46;;
};
if (($v_12 <= 0)) {
$__t46 = ($v1_13)->{'value0'};
goto end_branch_46;;
};
$__tco_44 = ($v_12 - 1);
$__tco_45 = ($v1_13)->{'value1'};
$__tco_var_go__go_11_43_43_v_12 = $__tco_44;
$__tco_var_go__go_11_43_43_v1_13 = $__tco_45;
goto tco_loop_go__go_11_43_43;;
$__t46 = null;
end_branch_46:;
$__t43 = $__t46;
goto end_branch_43;;
};
  if ($v1_13 instanceof \Control\Monad\Gen\Control_Monad_Gen_Nil) {
$__t43 = (((($__local_var_8_20)->{'foldMap1'})($GLOBALS['Data_Semigroup_Last_semigroupLast']))(function($x_14) {
  $__num = \func_num_args();
  $__res = $x_14;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($__local_var_9_29);
goto end_branch_43;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t43 = null;
  end_branch_43:;
  $__res = $__t43;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ($pure_7_19)((($go__go_11_43)($n_10))((((((($__local_var_8_20)->{'Foldable0'})(null))->{'foldr'})($GLOBALS['Control_Monad_Gen_Cons']))(new \Control\Monad\Gen\Control_Monad_Gen_Nil()))($__local_var_9_29)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
goto end_branch_17;;
};
  if ($v_5_15 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t17 = (($Applicative0_4_14)->{'pure'})(0);
goto end_branch_17;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t17 = null;
  end_branch_17:;
  $Applicative0_4_44 = (((($dictMonadGen_0)->{'Monad0'})(null))->{'Applicative0'})(null);
  $v_5_45 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_5) {
  $__num = \func_num_args();
  $__t45 = null;;
  if ((($n_5 >= 0) && ($n_5 <= 59))) {
$__t45 = new \Data\Maybe\Data_Maybe_Just($n_5);
goto end_branch_45;;
};
  $__t45 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_45:;
  $__res = $__t45;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_5) {
  $__num = \func_num_args();
  $__res = ($v_5 + 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_5) {
  $__num = \func_num_args();
  $__res = $v_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), 0);
  $__t47 = null;;
  if ($v_5_45 instanceof \Data\Maybe\Data_Maybe_Just) {
$Monad0_6_48 = (($dictMonadGen_0)->{'Monad0'})(null);
$pure_7_49 = ((($Monad0_6_48)->{'Applicative0'})(null))->{'pure'};
$foldableNonEmpty1_8_50 = (object)["foldMap" => function($dictMonoid_8) {
  $__num = \func_num_args();
  $Semigroup0_9_50 = (($dictMonoid_8)->{'Semigroup0'})(null);
  $__res = function($f_10) use ($Semigroup0_9_50, $dictMonoid_8) {
  $__num = \func_num_args();
  $__res = function($v_11) use ($Semigroup0_9_50, $dictMonoid_8, $f_10) {
  $__num = \func_num_args();
  $Semigroup0_12_51 = (($dictMonoid_8)->{'Semigroup0'})(null);
  $__res = ((($Semigroup0_9_50)->{'append'})(($f_10)(($v_11)->{'value0'})))(\Data\Foldable\majData_majFoldable_foldrmajArray(function($x_13) use ($Semigroup0_12_51, $f_10) {
  $__num = \func_num_args();
  $__res = function($acc_14) use ($Semigroup0_12_51, $f_10, $x_13) {
  $__num = \func_num_args();
  $__res = ((($Semigroup0_12_51)->{'append'})(($f_10)($x_13)))($acc_14);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($dictMonoid_8)->{'mempty'}, ($v_11)->{'value1'}));
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
}, "foldl" => function($f_8) {
  $__num = \func_num_args();
  $__res = function($b_9) use ($f_8) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($b_9, $f_8) {
  $__num = \func_num_args();
  $__res = \Data\Foldable\majData_majFoldable_foldlmajArray($f_8, (($f_8)($b_9))(($v_10)->{'value0'}), ($v_10)->{'value1'});
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
}, "foldr" => function($f_8) {
  $__num = \func_num_args();
  $__res = function($b_9) use ($f_8) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($b_9, $f_8) {
  $__num = \func_num_args();
  $__res = (($f_8)(($v_10)->{'value0'}))(\Data\Foldable\majData_majFoldable_foldrmajArray($f_8, $b_9, ($v_10)->{'value1'}));
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
$__local_var_8_50 = (object)["foldMap1" => function($dictSemigroup_9) {
  $__num = \func_num_args();
  $__res = function($f_10) use ($dictSemigroup_9) {
  $__num = \func_num_args();
  $__res = function($v_11) use ($dictSemigroup_9, $f_10) {
  $__num = \func_num_args();
  $__res = \Data\Foldable\majData_majFoldable_foldlmajArray(function($s_12) use ($dictSemigroup_9, $f_10) {
  $__num = \func_num_args();
  $__res = function($a1_13) use ($dictSemigroup_9, $f_10, $s_12) {
  $__num = \func_num_args();
  $__res = ((($dictSemigroup_9)->{'append'})($s_12))(($f_10)($a1_13));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($f_10)(($v_11)->{'value0'}), ($v_11)->{'value1'});
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
}, "foldr1" => function($f_9) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($f_9) {
  $__num = \func_num_args();
  $__local_var_11_53 = ($f_9)(($v_10)->{'value0'});
  $__local_var_12_54 = \Data\Foldable\majData_majFoldable_foldrmajArray(function($a1_12) use ($f_9) {
  $__num = \func_num_args();
  $__local_var_13_54 = ($f_9)($a1_12);
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Maybe_Just']))(function($v2_14) use ($__local_var_13_54, $a1_12) {
  $__num = \func_num_args();
  $__t55 = null;;
  if ($v2_14 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t55 = $a1_12;
goto end_branch_55;;
};
  if ($v2_14 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t55 = ($__local_var_13_54)(($v2_14)->{'value0'});
goto end_branch_55;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t55 = null;
  end_branch_55:;
  $__res = $__t55;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, new \Data\Maybe\Data_Maybe_Nothing(), ($v_10)->{'value1'});
  $__t57 = null;;
  if ($__local_var_12_54 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t57 = ($v_10)->{'value0'};
goto end_branch_57;;
};
  if ($__local_var_12_54 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t57 = ($__local_var_11_53)(($__local_var_12_54)->{'value0'});
goto end_branch_57;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t57 = null;
  end_branch_57:;
  $__res = $__t57;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldl1" => function($f_9) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($f_9) {
  $__num = \func_num_args();
  $__res = \Data\Foldable\majData_majFoldable_foldlmajArray($f_9, ($v_10)->{'value0'}, ($v_10)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable0" => function($_dollar___unused_9) use ($foldableNonEmpty1_8_50) {
  $__num = \func_num_args();
  $__res = $foldableNonEmpty1_8_50;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
$__t64 = null;;
switch (($v_5_45)->{'value0'}) {
case 59:
$__t64 = \Data\Unfoldable1\majData_majUnfoldable1_unfoldr1majArraymajImpl($GLOBALS['Data_Maybe_isNothing'], function($v_9) {
  $__num = \func_num_args();
  $__t65 = null;;
  if ($v_9 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t65 = ($v_9)->{'value0'};
goto end_branch_65;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t65 = null;
  end_branch_65:;
  $__res = $__t65;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $GLOBALS['Data_Tuple_fst'], $GLOBALS['Data_Tuple_snd'], function($i_9) use ($v_5_45) {
  $__num = \func_num_args();
  $__t66 = null;;
  if (($i_9 <= 0)) {
$__t66 = new \Data\Tuple\Data_Tuple_Tuple(($v_5_45)->{'value0'}, new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_66;;
};
  $__t66 = new \Data\Tuple\Data_Tuple_Tuple(($v_5_45)->{'value0'}, new \Data\Maybe\Data_Maybe_Just(($i_9 - 1)));
  end_branch_66:;
  $__res = $__t66;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, 0);
goto end_branch_64;;
break;
default:
;
break;
};
if ((($v_5_45)->{'value0'} < 59)) {
$__t64 = \Data\Unfoldable1\majData_majUnfoldable1_unfoldr1majArraymajImpl($GLOBALS['Data_Maybe_isNothing'], function($v_9) {
  $__num = \func_num_args();
  $__t67 = null;;
  if ($v_9 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t67 = ($v_9)->{'value0'};
goto end_branch_67;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t67 = null;
  end_branch_67:;
  $__res = $__t67;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $GLOBALS['Data_Tuple_fst'], $GLOBALS['Data_Tuple_snd'], function($a_9) {
  $__num = \func_num_args();
  $__local_var_10_68 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_10) {
  $__num = \func_num_args();
  $__t68 = null;;
  if ((($n_10 >= 0) && ($n_10 <= 59))) {
$__t68 = new \Data\Maybe\Data_Maybe_Just($n_10);
goto end_branch_68;;
};
  $__t68 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_68:;
  $__res = $__t68;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_10) {
  $__num = \func_num_args();
  $__res = ($v_10 + 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_10) {
  $__num = \func_num_args();
  $__res = $v_10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), $a_9);
  $__t70 = null;;
  if ($__local_var_10_68 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t71 = null;;
if ((($__local_var_10_68)->{'value0'} <= 59)) {
$__t71 = new \Data\Maybe\Data_Maybe_Just(($__local_var_10_68)->{'value0'});
goto end_branch_71;;
};
$__t71 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_71:;
$__t70 = $__t71;
goto end_branch_70;;
};
  if ($__local_var_10_68 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t70 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_70;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t70 = null;
  end_branch_70:;
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_9, $__t70);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($v_5_45)->{'value0'});
goto end_branch_64;;
};
$__t64 = \Data\Unfoldable1\majData_majUnfoldable1_unfoldr1majArraymajImpl($GLOBALS['Data_Maybe_isNothing'], function($v_9) {
  $__num = \func_num_args();
  $__t59 = null;;
  if ($v_9 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t59 = ($v_9)->{'value0'};
goto end_branch_59;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t59 = null;
  end_branch_59:;
  $__res = $__t59;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $GLOBALS['Data_Tuple_fst'], $GLOBALS['Data_Tuple_snd'], function($a_9) {
  $__num = \func_num_args();
  $__local_var_10_60 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_10) {
  $__num = \func_num_args();
  $__t60 = null;;
  if ((($n_10 >= 0) && ($n_10 <= 59))) {
$__t60 = new \Data\Maybe\Data_Maybe_Just($n_10);
goto end_branch_60;;
};
  $__t60 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_60:;
  $__res = $__t60;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_10) {
  $__num = \func_num_args();
  $__res = ($v_10 - 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_10) {
  $__num = \func_num_args();
  $__res = $v_10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), $a_9);
  $__t62 = null;;
  if ($__local_var_10_60 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t63 = null;;
if ((($__local_var_10_60)->{'value0'} >= 59)) {
$__t63 = new \Data\Maybe\Data_Maybe_Just(($__local_var_10_60)->{'value0'});
goto end_branch_63;;
};
$__t63 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_63:;
$__t62 = $__t63;
goto end_branch_62;;
};
  if ($__local_var_10_60 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t62 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_62;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t62 = null;
  end_branch_62:;
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_9, $__t62);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($v_5_45)->{'value0'});
end_branch_64:;
$__local_var_9_59 = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(0, $__t64);
$__t47 = ((((($Monad0_6_48)->{'Bind1'})(null))->{'bind'})(((($dictMonadGen_0)->{'chooseInt'})(0))(((((((($__local_var_8_50)->{'Foldable0'})(null))->{'foldl'})(function($c_10) {
  $__num = \func_num_args();
  $__res = function($v_11) use ($c_10) {
  $__num = \func_num_args();
  $__res = (1 + $c_10);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(0))($__local_var_9_59) - 1))))(function($n_10) use ($__local_var_8_50, $__local_var_9_59, $pure_7_49) {
  $__num = \func_num_args();
  $go__go_11_73 = null;
  $go__go_11_73 = (function() use ($__local_var_8_50, $__local_var_9_59, &$go__go_11_73) {
  $__fn = function(int $v_12, $v1_13 = null) use ($__local_var_8_50, $__local_var_9_59, &$go__go_11_73, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_11_73_73_v_12 = $v_12;
  $__tco_var_go__go_11_73_73_v1_13 = $v1_13;
  tco_loop_go__go_11_73_73:;
  $v_12 = $__tco_var_go__go_11_73_73_v_12;
  $v1_13 = $__tco_var_go__go_11_73_73_v1_13;
  $__t73 = null;;
  if ($v1_13 instanceof \Control\Monad\Gen\Control_Monad_Gen_Cons) {
$__t76 = null;;
if (($v1_13)->{'value1'} instanceof \Control\Monad\Gen\Control_Monad_Gen_Nil) {
$__t76 = ($v1_13)->{'value0'};
goto end_branch_76;;
};
if (($v_12 <= 0)) {
$__t76 = ($v1_13)->{'value0'};
goto end_branch_76;;
};
$__tco_74 = ($v_12 - 1);
$__tco_75 = ($v1_13)->{'value1'};
$__tco_var_go__go_11_73_73_v_12 = $__tco_74;
$__tco_var_go__go_11_73_73_v1_13 = $__tco_75;
goto tco_loop_go__go_11_73_73;;
$__t76 = null;
end_branch_76:;
$__t73 = $__t76;
goto end_branch_73;;
};
  if ($v1_13 instanceof \Control\Monad\Gen\Control_Monad_Gen_Nil) {
$__t73 = (((($__local_var_8_50)->{'foldMap1'})($GLOBALS['Data_Semigroup_Last_semigroupLast']))(function($x_14) {
  $__num = \func_num_args();
  $__res = $x_14;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($__local_var_9_59);
goto end_branch_73;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t73 = null;
  end_branch_73:;
  $__res = $__t73;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ($pure_7_49)((($go__go_11_73)($n_10))((((((($__local_var_8_50)->{'Foldable0'})(null))->{'foldr'})($GLOBALS['Control_Monad_Gen_Cons']))(new \Control\Monad\Gen\Control_Monad_Gen_Nil()))($__local_var_9_59)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
goto end_branch_47;;
};
  if ($v_5_45 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t47 = (($Applicative0_4_44)->{'pure'})(0);
goto end_branch_47;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t47 = null;
  end_branch_47:;
  $Applicative0_4_74 = (((($dictMonadGen_0)->{'Monad0'})(null))->{'Applicative0'})(null);
  $v_5_75 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_5) {
  $__num = \func_num_args();
  $__t75 = null;;
  if ((($n_5 >= 0) && ($n_5 <= 59))) {
$__t75 = new \Data\Maybe\Data_Maybe_Just($n_5);
goto end_branch_75;;
};
  $__t75 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_75:;
  $__res = $__t75;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_5) {
  $__num = \func_num_args();
  $__res = ($v_5 + 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_5) {
  $__num = \func_num_args();
  $__res = $v_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), 0);
  $__t77 = null;;
  if ($v_5_75 instanceof \Data\Maybe\Data_Maybe_Just) {
$Monad0_6_78 = (($dictMonadGen_0)->{'Monad0'})(null);
$pure_7_79 = ((($Monad0_6_78)->{'Applicative0'})(null))->{'pure'};
$foldableNonEmpty1_8_80 = (object)["foldMap" => function($dictMonoid_8) {
  $__num = \func_num_args();
  $Semigroup0_9_80 = (($dictMonoid_8)->{'Semigroup0'})(null);
  $__res = function($f_10) use ($Semigroup0_9_80, $dictMonoid_8) {
  $__num = \func_num_args();
  $__res = function($v_11) use ($Semigroup0_9_80, $dictMonoid_8, $f_10) {
  $__num = \func_num_args();
  $Semigroup0_12_81 = (($dictMonoid_8)->{'Semigroup0'})(null);
  $__res = ((($Semigroup0_9_80)->{'append'})(($f_10)(($v_11)->{'value0'})))(\Data\Foldable\majData_majFoldable_foldrmajArray(function($x_13) use ($Semigroup0_12_81, $f_10) {
  $__num = \func_num_args();
  $__res = function($acc_14) use ($Semigroup0_12_81, $f_10, $x_13) {
  $__num = \func_num_args();
  $__res = ((($Semigroup0_12_81)->{'append'})(($f_10)($x_13)))($acc_14);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($dictMonoid_8)->{'mempty'}, ($v_11)->{'value1'}));
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
}, "foldl" => function($f_8) {
  $__num = \func_num_args();
  $__res = function($b_9) use ($f_8) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($b_9, $f_8) {
  $__num = \func_num_args();
  $__res = \Data\Foldable\majData_majFoldable_foldlmajArray($f_8, (($f_8)($b_9))(($v_10)->{'value0'}), ($v_10)->{'value1'});
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
}, "foldr" => function($f_8) {
  $__num = \func_num_args();
  $__res = function($b_9) use ($f_8) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($b_9, $f_8) {
  $__num = \func_num_args();
  $__res = (($f_8)(($v_10)->{'value0'}))(\Data\Foldable\majData_majFoldable_foldrmajArray($f_8, $b_9, ($v_10)->{'value1'}));
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
$__local_var_8_80 = (object)["foldMap1" => function($dictSemigroup_9) {
  $__num = \func_num_args();
  $__res = function($f_10) use ($dictSemigroup_9) {
  $__num = \func_num_args();
  $__res = function($v_11) use ($dictSemigroup_9, $f_10) {
  $__num = \func_num_args();
  $__res = \Data\Foldable\majData_majFoldable_foldlmajArray(function($s_12) use ($dictSemigroup_9, $f_10) {
  $__num = \func_num_args();
  $__res = function($a1_13) use ($dictSemigroup_9, $f_10, $s_12) {
  $__num = \func_num_args();
  $__res = ((($dictSemigroup_9)->{'append'})($s_12))(($f_10)($a1_13));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($f_10)(($v_11)->{'value0'}), ($v_11)->{'value1'});
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
}, "foldr1" => function($f_9) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($f_9) {
  $__num = \func_num_args();
  $__local_var_11_83 = ($f_9)(($v_10)->{'value0'});
  $__local_var_12_84 = \Data\Foldable\majData_majFoldable_foldrmajArray(function($a1_12) use ($f_9) {
  $__num = \func_num_args();
  $__local_var_13_84 = ($f_9)($a1_12);
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Maybe_Just']))(function($v2_14) use ($__local_var_13_84, $a1_12) {
  $__num = \func_num_args();
  $__t85 = null;;
  if ($v2_14 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t85 = $a1_12;
goto end_branch_85;;
};
  if ($v2_14 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t85 = ($__local_var_13_84)(($v2_14)->{'value0'});
goto end_branch_85;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t85 = null;
  end_branch_85:;
  $__res = $__t85;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, new \Data\Maybe\Data_Maybe_Nothing(), ($v_10)->{'value1'});
  $__t87 = null;;
  if ($__local_var_12_84 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t87 = ($v_10)->{'value0'};
goto end_branch_87;;
};
  if ($__local_var_12_84 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t87 = ($__local_var_11_83)(($__local_var_12_84)->{'value0'});
goto end_branch_87;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t87 = null;
  end_branch_87:;
  $__res = $__t87;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldl1" => function($f_9) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($f_9) {
  $__num = \func_num_args();
  $__res = \Data\Foldable\majData_majFoldable_foldlmajArray($f_9, ($v_10)->{'value0'}, ($v_10)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable0" => function($_dollar___unused_9) use ($foldableNonEmpty1_8_80) {
  $__num = \func_num_args();
  $__res = $foldableNonEmpty1_8_80;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
$__t94 = null;;
switch (($v_5_75)->{'value0'}) {
case 59:
$__t94 = \Data\Unfoldable1\majData_majUnfoldable1_unfoldr1majArraymajImpl($GLOBALS['Data_Maybe_isNothing'], function($v_9) {
  $__num = \func_num_args();
  $__t95 = null;;
  if ($v_9 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t95 = ($v_9)->{'value0'};
goto end_branch_95;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t95 = null;
  end_branch_95:;
  $__res = $__t95;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $GLOBALS['Data_Tuple_fst'], $GLOBALS['Data_Tuple_snd'], function($i_9) use ($v_5_75) {
  $__num = \func_num_args();
  $__t96 = null;;
  if (($i_9 <= 0)) {
$__t96 = new \Data\Tuple\Data_Tuple_Tuple(($v_5_75)->{'value0'}, new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_96;;
};
  $__t96 = new \Data\Tuple\Data_Tuple_Tuple(($v_5_75)->{'value0'}, new \Data\Maybe\Data_Maybe_Just(($i_9 - 1)));
  end_branch_96:;
  $__res = $__t96;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, 0);
goto end_branch_94;;
break;
default:
;
break;
};
if ((($v_5_75)->{'value0'} < 59)) {
$__t94 = \Data\Unfoldable1\majData_majUnfoldable1_unfoldr1majArraymajImpl($GLOBALS['Data_Maybe_isNothing'], function($v_9) {
  $__num = \func_num_args();
  $__t97 = null;;
  if ($v_9 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t97 = ($v_9)->{'value0'};
goto end_branch_97;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t97 = null;
  end_branch_97:;
  $__res = $__t97;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $GLOBALS['Data_Tuple_fst'], $GLOBALS['Data_Tuple_snd'], function($a_9) {
  $__num = \func_num_args();
  $__local_var_10_98 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_10) {
  $__num = \func_num_args();
  $__t98 = null;;
  if ((($n_10 >= 0) && ($n_10 <= 59))) {
$__t98 = new \Data\Maybe\Data_Maybe_Just($n_10);
goto end_branch_98;;
};
  $__t98 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_98:;
  $__res = $__t98;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_10) {
  $__num = \func_num_args();
  $__res = ($v_10 + 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_10) {
  $__num = \func_num_args();
  $__res = $v_10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), $a_9);
  $__t100 = null;;
  if ($__local_var_10_98 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t101 = null;;
if ((($__local_var_10_98)->{'value0'} <= 59)) {
$__t101 = new \Data\Maybe\Data_Maybe_Just(($__local_var_10_98)->{'value0'});
goto end_branch_101;;
};
$__t101 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_101:;
$__t100 = $__t101;
goto end_branch_100;;
};
  if ($__local_var_10_98 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t100 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_100;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t100 = null;
  end_branch_100:;
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_9, $__t100);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($v_5_75)->{'value0'});
goto end_branch_94;;
};
$__t94 = \Data\Unfoldable1\majData_majUnfoldable1_unfoldr1majArraymajImpl($GLOBALS['Data_Maybe_isNothing'], function($v_9) {
  $__num = \func_num_args();
  $__t89 = null;;
  if ($v_9 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t89 = ($v_9)->{'value0'};
goto end_branch_89;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t89 = null;
  end_branch_89:;
  $__res = $__t89;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $GLOBALS['Data_Tuple_fst'], $GLOBALS['Data_Tuple_snd'], function($a_9) {
  $__num = \func_num_args();
  $__local_var_10_90 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_10) {
  $__num = \func_num_args();
  $__t90 = null;;
  if ((($n_10 >= 0) && ($n_10 <= 59))) {
$__t90 = new \Data\Maybe\Data_Maybe_Just($n_10);
goto end_branch_90;;
};
  $__t90 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_90:;
  $__res = $__t90;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_10) {
  $__num = \func_num_args();
  $__res = ($v_10 - 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_10) {
  $__num = \func_num_args();
  $__res = $v_10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), $a_9);
  $__t92 = null;;
  if ($__local_var_10_90 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t93 = null;;
if ((($__local_var_10_90)->{'value0'} >= 59)) {
$__t93 = new \Data\Maybe\Data_Maybe_Just(($__local_var_10_90)->{'value0'});
goto end_branch_93;;
};
$__t93 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_93:;
$__t92 = $__t93;
goto end_branch_92;;
};
  if ($__local_var_10_90 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t92 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_92;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t92 = null;
  end_branch_92:;
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_9, $__t92);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($v_5_75)->{'value0'});
end_branch_94:;
$__local_var_9_89 = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(0, $__t94);
$__t77 = ((((($Monad0_6_78)->{'Bind1'})(null))->{'bind'})(((($dictMonadGen_0)->{'chooseInt'})(0))(((((((($__local_var_8_80)->{'Foldable0'})(null))->{'foldl'})(function($c_10) {
  $__num = \func_num_args();
  $__res = function($v_11) use ($c_10) {
  $__num = \func_num_args();
  $__res = (1 + $c_10);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(0))($__local_var_9_89) - 1))))(function($n_10) use ($__local_var_8_80, $__local_var_9_89, $pure_7_79) {
  $__num = \func_num_args();
  $go__go_11_103 = null;
  $go__go_11_103 = (function() use ($__local_var_8_80, $__local_var_9_89, &$go__go_11_103) {
  $__fn = function(int $v_12, $v1_13 = null) use ($__local_var_8_80, $__local_var_9_89, &$go__go_11_103, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_11_103_103_v_12 = $v_12;
  $__tco_var_go__go_11_103_103_v1_13 = $v1_13;
  tco_loop_go__go_11_103_103:;
  $v_12 = $__tco_var_go__go_11_103_103_v_12;
  $v1_13 = $__tco_var_go__go_11_103_103_v1_13;
  $__t103 = null;;
  if ($v1_13 instanceof \Control\Monad\Gen\Control_Monad_Gen_Cons) {
$__t106 = null;;
if (($v1_13)->{'value1'} instanceof \Control\Monad\Gen\Control_Monad_Gen_Nil) {
$__t106 = ($v1_13)->{'value0'};
goto end_branch_106;;
};
if (($v_12 <= 0)) {
$__t106 = ($v1_13)->{'value0'};
goto end_branch_106;;
};
$__tco_104 = ($v_12 - 1);
$__tco_105 = ($v1_13)->{'value1'};
$__tco_var_go__go_11_103_103_v_12 = $__tco_104;
$__tco_var_go__go_11_103_103_v1_13 = $__tco_105;
goto tco_loop_go__go_11_103_103;;
$__t106 = null;
end_branch_106:;
$__t103 = $__t106;
goto end_branch_103;;
};
  if ($v1_13 instanceof \Control\Monad\Gen\Control_Monad_Gen_Nil) {
$__t103 = (((($__local_var_8_80)->{'foldMap1'})($GLOBALS['Data_Semigroup_Last_semigroupLast']))(function($x_14) {
  $__num = \func_num_args();
  $__res = $x_14;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($__local_var_9_89);
goto end_branch_103;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t103 = null;
  end_branch_103:;
  $__res = $__t103;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ($pure_7_79)((($go__go_11_103)($n_10))((((((($__local_var_8_80)->{'Foldable0'})(null))->{'foldr'})($GLOBALS['Control_Monad_Gen_Cons']))(new \Control\Monad\Gen\Control_Monad_Gen_Nil()))($__local_var_9_89)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
goto end_branch_77;;
};
  if ($v_5_75 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t77 = (($Applicative0_4_74)->{'pure'})(0);
goto end_branch_77;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t77 = null;
  end_branch_77:;
  $Applicative0_4_104 = (((($dictMonadGen_0)->{'Monad0'})(null))->{'Applicative0'})(null);
  $v_5_105 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_5) {
  $__num = \func_num_args();
  $__t105 = null;;
  if ((($n_5 >= 0) && ($n_5 <= 999))) {
$__t105 = new \Data\Maybe\Data_Maybe_Just($n_5);
goto end_branch_105;;
};
  $__t105 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_105:;
  $__res = $__t105;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_5) {
  $__num = \func_num_args();
  $__res = ($v_5 + 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_5) {
  $__num = \func_num_args();
  $__res = $v_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), 0);
  $__t107 = null;;
  if ($v_5_105 instanceof \Data\Maybe\Data_Maybe_Just) {
$Monad0_6_108 = (($dictMonadGen_0)->{'Monad0'})(null);
$pure_7_109 = ((($Monad0_6_108)->{'Applicative0'})(null))->{'pure'};
$foldableNonEmpty1_8_110 = (object)["foldMap" => function($dictMonoid_8) {
  $__num = \func_num_args();
  $Semigroup0_9_110 = (($dictMonoid_8)->{'Semigroup0'})(null);
  $__res = function($f_10) use ($Semigroup0_9_110, $dictMonoid_8) {
  $__num = \func_num_args();
  $__res = function($v_11) use ($Semigroup0_9_110, $dictMonoid_8, $f_10) {
  $__num = \func_num_args();
  $Semigroup0_12_111 = (($dictMonoid_8)->{'Semigroup0'})(null);
  $__res = ((($Semigroup0_9_110)->{'append'})(($f_10)(($v_11)->{'value0'})))(\Data\Foldable\majData_majFoldable_foldrmajArray(function($x_13) use ($Semigroup0_12_111, $f_10) {
  $__num = \func_num_args();
  $__res = function($acc_14) use ($Semigroup0_12_111, $f_10, $x_13) {
  $__num = \func_num_args();
  $__res = ((($Semigroup0_12_111)->{'append'})(($f_10)($x_13)))($acc_14);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($dictMonoid_8)->{'mempty'}, ($v_11)->{'value1'}));
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
}, "foldl" => function($f_8) {
  $__num = \func_num_args();
  $__res = function($b_9) use ($f_8) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($b_9, $f_8) {
  $__num = \func_num_args();
  $__res = \Data\Foldable\majData_majFoldable_foldlmajArray($f_8, (($f_8)($b_9))(($v_10)->{'value0'}), ($v_10)->{'value1'});
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
}, "foldr" => function($f_8) {
  $__num = \func_num_args();
  $__res = function($b_9) use ($f_8) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($b_9, $f_8) {
  $__num = \func_num_args();
  $__res = (($f_8)(($v_10)->{'value0'}))(\Data\Foldable\majData_majFoldable_foldrmajArray($f_8, $b_9, ($v_10)->{'value1'}));
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
$__local_var_8_110 = (object)["foldMap1" => function($dictSemigroup_9) {
  $__num = \func_num_args();
  $__res = function($f_10) use ($dictSemigroup_9) {
  $__num = \func_num_args();
  $__res = function($v_11) use ($dictSemigroup_9, $f_10) {
  $__num = \func_num_args();
  $__res = \Data\Foldable\majData_majFoldable_foldlmajArray(function($s_12) use ($dictSemigroup_9, $f_10) {
  $__num = \func_num_args();
  $__res = function($a1_13) use ($dictSemigroup_9, $f_10, $s_12) {
  $__num = \func_num_args();
  $__res = ((($dictSemigroup_9)->{'append'})($s_12))(($f_10)($a1_13));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($f_10)(($v_11)->{'value0'}), ($v_11)->{'value1'});
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
}, "foldr1" => function($f_9) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($f_9) {
  $__num = \func_num_args();
  $__local_var_11_113 = ($f_9)(($v_10)->{'value0'});
  $__local_var_12_114 = \Data\Foldable\majData_majFoldable_foldrmajArray(function($a1_12) use ($f_9) {
  $__num = \func_num_args();
  $__local_var_13_114 = ($f_9)($a1_12);
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Maybe_Just']))(function($v2_14) use ($__local_var_13_114, $a1_12) {
  $__num = \func_num_args();
  $__t115 = null;;
  if ($v2_14 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t115 = $a1_12;
goto end_branch_115;;
};
  if ($v2_14 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t115 = ($__local_var_13_114)(($v2_14)->{'value0'});
goto end_branch_115;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t115 = null;
  end_branch_115:;
  $__res = $__t115;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, new \Data\Maybe\Data_Maybe_Nothing(), ($v_10)->{'value1'});
  $__t117 = null;;
  if ($__local_var_12_114 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t117 = ($v_10)->{'value0'};
goto end_branch_117;;
};
  if ($__local_var_12_114 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t117 = ($__local_var_11_113)(($__local_var_12_114)->{'value0'});
goto end_branch_117;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t117 = null;
  end_branch_117:;
  $__res = $__t117;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldl1" => function($f_9) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($f_9) {
  $__num = \func_num_args();
  $__res = \Data\Foldable\majData_majFoldable_foldlmajArray($f_9, ($v_10)->{'value0'}, ($v_10)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable0" => function($_dollar___unused_9) use ($foldableNonEmpty1_8_110) {
  $__num = \func_num_args();
  $__res = $foldableNonEmpty1_8_110;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
$__t124 = null;;
switch (($v_5_105)->{'value0'}) {
case 999:
$__t124 = \Data\Unfoldable1\majData_majUnfoldable1_unfoldr1majArraymajImpl($GLOBALS['Data_Maybe_isNothing'], function($v_9) {
  $__num = \func_num_args();
  $__t125 = null;;
  if ($v_9 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t125 = ($v_9)->{'value0'};
goto end_branch_125;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t125 = null;
  end_branch_125:;
  $__res = $__t125;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $GLOBALS['Data_Tuple_fst'], $GLOBALS['Data_Tuple_snd'], function($i_9) use ($v_5_105) {
  $__num = \func_num_args();
  $__t126 = null;;
  if (($i_9 <= 0)) {
$__t126 = new \Data\Tuple\Data_Tuple_Tuple(($v_5_105)->{'value0'}, new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_126;;
};
  $__t126 = new \Data\Tuple\Data_Tuple_Tuple(($v_5_105)->{'value0'}, new \Data\Maybe\Data_Maybe_Just(($i_9 - 1)));
  end_branch_126:;
  $__res = $__t126;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, 0);
goto end_branch_124;;
break;
default:
;
break;
};
if ((($v_5_105)->{'value0'} < 999)) {
$__t124 = \Data\Unfoldable1\majData_majUnfoldable1_unfoldr1majArraymajImpl($GLOBALS['Data_Maybe_isNothing'], function($v_9) {
  $__num = \func_num_args();
  $__t127 = null;;
  if ($v_9 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t127 = ($v_9)->{'value0'};
goto end_branch_127;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t127 = null;
  end_branch_127:;
  $__res = $__t127;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $GLOBALS['Data_Tuple_fst'], $GLOBALS['Data_Tuple_snd'], function($a_9) {
  $__num = \func_num_args();
  $__local_var_10_128 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_10) {
  $__num = \func_num_args();
  $__t128 = null;;
  if ((($n_10 >= 0) && ($n_10 <= 999))) {
$__t128 = new \Data\Maybe\Data_Maybe_Just($n_10);
goto end_branch_128;;
};
  $__t128 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_128:;
  $__res = $__t128;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_10) {
  $__num = \func_num_args();
  $__res = ($v_10 + 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_10) {
  $__num = \func_num_args();
  $__res = $v_10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), $a_9);
  $__t130 = null;;
  if ($__local_var_10_128 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t131 = null;;
if ((($__local_var_10_128)->{'value0'} <= 999)) {
$__t131 = new \Data\Maybe\Data_Maybe_Just(($__local_var_10_128)->{'value0'});
goto end_branch_131;;
};
$__t131 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_131:;
$__t130 = $__t131;
goto end_branch_130;;
};
  if ($__local_var_10_128 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t130 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_130;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t130 = null;
  end_branch_130:;
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_9, $__t130);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($v_5_105)->{'value0'});
goto end_branch_124;;
};
$__t124 = \Data\Unfoldable1\majData_majUnfoldable1_unfoldr1majArraymajImpl($GLOBALS['Data_Maybe_isNothing'], function($v_9) {
  $__num = \func_num_args();
  $__t119 = null;;
  if ($v_9 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t119 = ($v_9)->{'value0'};
goto end_branch_119;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t119 = null;
  end_branch_119:;
  $__res = $__t119;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $GLOBALS['Data_Tuple_fst'], $GLOBALS['Data_Tuple_snd'], function($a_9) {
  $__num = \func_num_args();
  $__local_var_10_120 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_10) {
  $__num = \func_num_args();
  $__t120 = null;;
  if ((($n_10 >= 0) && ($n_10 <= 999))) {
$__t120 = new \Data\Maybe\Data_Maybe_Just($n_10);
goto end_branch_120;;
};
  $__t120 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_120:;
  $__res = $__t120;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_10) {
  $__num = \func_num_args();
  $__res = ($v_10 - 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_10) {
  $__num = \func_num_args();
  $__res = $v_10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), $a_9);
  $__t122 = null;;
  if ($__local_var_10_120 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t123 = null;;
if ((($__local_var_10_120)->{'value0'} >= 999)) {
$__t123 = new \Data\Maybe\Data_Maybe_Just(($__local_var_10_120)->{'value0'});
goto end_branch_123;;
};
$__t123 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_123:;
$__t122 = $__t123;
goto end_branch_122;;
};
  if ($__local_var_10_120 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t122 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_122;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t122 = null;
  end_branch_122:;
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_9, $__t122);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($v_5_105)->{'value0'});
end_branch_124:;
$__local_var_9_119 = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(0, $__t124);
$__t107 = ((((($Monad0_6_108)->{'Bind1'})(null))->{'bind'})(((($dictMonadGen_0)->{'chooseInt'})(0))(((((((($__local_var_8_110)->{'Foldable0'})(null))->{'foldl'})(function($c_10) {
  $__num = \func_num_args();
  $__res = function($v_11) use ($c_10) {
  $__num = \func_num_args();
  $__res = (1 + $c_10);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(0))($__local_var_9_119) - 1))))(function($n_10) use ($__local_var_8_110, $__local_var_9_119, $pure_7_109) {
  $__num = \func_num_args();
  $go__go_11_133 = null;
  $go__go_11_133 = (function() use ($__local_var_8_110, $__local_var_9_119, &$go__go_11_133) {
  $__fn = function(int $v_12, $v1_13 = null) use ($__local_var_8_110, $__local_var_9_119, &$go__go_11_133, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_11_133_133_v_12 = $v_12;
  $__tco_var_go__go_11_133_133_v1_13 = $v1_13;
  tco_loop_go__go_11_133_133:;
  $v_12 = $__tco_var_go__go_11_133_133_v_12;
  $v1_13 = $__tco_var_go__go_11_133_133_v1_13;
  $__t133 = null;;
  if ($v1_13 instanceof \Control\Monad\Gen\Control_Monad_Gen_Cons) {
$__t136 = null;;
if (($v1_13)->{'value1'} instanceof \Control\Monad\Gen\Control_Monad_Gen_Nil) {
$__t136 = ($v1_13)->{'value0'};
goto end_branch_136;;
};
if (($v_12 <= 0)) {
$__t136 = ($v1_13)->{'value0'};
goto end_branch_136;;
};
$__tco_134 = ($v_12 - 1);
$__tco_135 = ($v1_13)->{'value1'};
$__tco_var_go__go_11_133_133_v_12 = $__tco_134;
$__tco_var_go__go_11_133_133_v1_13 = $__tco_135;
goto tco_loop_go__go_11_133_133;;
$__t136 = null;
end_branch_136:;
$__t133 = $__t136;
goto end_branch_133;;
};
  if ($v1_13 instanceof \Control\Monad\Gen\Control_Monad_Gen_Nil) {
$__t133 = (((($__local_var_8_110)->{'foldMap1'})($GLOBALS['Data_Semigroup_Last_semigroupLast']))(function($x_14) {
  $__num = \func_num_args();
  $__res = $x_14;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($__local_var_9_119);
goto end_branch_133;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t133 = null;
  end_branch_133:;
  $__res = $__t133;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ($pure_7_109)((($go__go_11_133)($n_10))((((((($__local_var_8_110)->{'Foldable0'})(null))->{'foldr'})($GLOBALS['Control_Monad_Gen_Cons']))(new \Control\Monad\Gen\Control_Monad_Gen_Nil()))($__local_var_9_119)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
goto end_branch_107;;
};
  if ($v_5_105 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t107 = (($Applicative0_4_104)->{'pure'})(0);
goto end_branch_107;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t107 = null;
  end_branch_107:;
  $__res = ((((($Bind1_1_0)->{'Apply0'})(null))->{'apply'})(((((((($Bind1_1_0)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_DateTime_DateTime']))(((($Bind1_3_2)->{'bind'})(((((((((((($dictMonadGen_0)->{'Monad0'})(null))->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_6) {
  $__num = \func_num_args();
  $__t5 = null;;
  if ($v_6 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t5 = ($v_6)->{'value0'};
goto end_branch_5;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t5 = null;
  end_branch_5:;
  $__res = $__t5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($n_6) {
  $__num = \func_num_args();
  $__t6 = null;;
  if ((($n_6 >= -271820) && ($n_6 <= 275759))) {
$__t6 = new \Data\Maybe\Data_Maybe_Just($n_6);
goto end_branch_6;;
};
  $__t6 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_6:;
  $__res = $__t6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(((($dictMonadGen_0)->{'chooseInt'})(1900))(2100))))(function($year_6) use ($Bind1_3_2, $Functor0_4_3, $dictMonadGen_0, $pure_5_4) {
  $__num = \func_num_args();
  $__t7 = null;;
  if (\Data\Date\majData_majDate_ismajLeapmajYear($year_6)) {
$__t7 = 365;
goto end_branch_7;;
};
  $__t7 = 364;
  end_branch_7:;
  $__res = ((($Bind1_3_2)->{'bind'})(((($Functor0_4_3)->{'map'})((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_7) {
  $__num = \func_num_args();
  $__res = $x_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Int_toNumber'])))(((($dictMonadGen_0)->{'chooseInt'})(0))($__t7))))(function($days_7) use ($pure_5_4, $year_6) {
  $__num = \func_num_args();
  $__local_var_8_8 = \Data\Date\majData_majDate_exactmajDate($year_6, new \Data\Date\Component\Data_Date_Component_January(), 1);
  $__t9 = null;;
  if ($__local_var_8_8 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t9 = \Data\Date\majData_majDate_adjust($days_7, ($__local_var_8_8)->{'value0'});
goto end_branch_9;;
};
  if ($__local_var_8_8 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t9 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_9;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t9 = null;
  end_branch_9:;
  $__local_var_8_8 = $__t9;
  $__t11 = null;;
  if ($__local_var_8_8 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t11 = ($__local_var_8_8)->{'value0'};
goto end_branch_11;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t11 = null;
  end_branch_11:;
  $__res = ($pure_5_4)($__t11);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))))(((($Apply0_3_13)->{'apply'})(((($Apply0_3_13)->{'apply'})(((($Apply0_3_13)->{'apply'})(((((((($Bind1_2_12)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Time_Time']))($__t17)))($__t47)))($__t77)))($__t107));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_DateTime_Gen_genDateTime'] = __NAMESPACE__ . '\\majData_majDatemajTime_majGen_genmajDatemajTime';

