<?php

namespace Data\Time\Gen;

// ALL IMPORTS: Control.Apply, Control.Monad.Gen, Data.Functor, Data.Time, Data.Time.Component.Gen, Prelude, Prim
// TO REQUIRE: Control.Apply, Control.Monad.Gen, Data.Functor, Data.Time, Data.Time.Component.Gen, Prelude
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Monad.Gen/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Time/index.php';
require_once __DIR__ . '/../Data.Time.Component.Gen/index.php';
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




// Data_Time_Gen_genTime
function majData_majTime_majGen_genmajTime($dictMonadGen_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTime_majGen_genmajTime';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Bind1_1_0 = (((($dictMonadGen_0)->{'Monad0'})(null))->{'Bind1'})(null);
  $Apply0_2_1 = (($Bind1_1_0)->{'Apply0'})(null);
  $Applicative0_3_2 = (((($dictMonadGen_0)->{'Monad0'})(null))->{'Applicative0'})(null);
  $v_4_3 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_4) {
  $__num = \func_num_args();
  $__t3 = null;;
  if ((($n_4 >= 0) && ($n_4 <= 23))) {
$__t3 = new \Data\Maybe\Data_Maybe_Just($n_4);
goto end_branch_3;;
};
  $__t3 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_4) {
  $__num = \func_num_args();
  $__res = ($v_4 + 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_4) {
  $__num = \func_num_args();
  $__res = $v_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), 0);
  $__t5 = null;;
  if ($v_4_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$Monad0_5_6 = (($dictMonadGen_0)->{'Monad0'})(null);
$pure_6_7 = ((($Monad0_5_6)->{'Applicative0'})(null))->{'pure'};
$foldableNonEmpty1_7_8 = (object)["foldMap" => function($dictMonoid_7) {
  $__num = \func_num_args();
  $Semigroup0_8_8 = (($dictMonoid_7)->{'Semigroup0'})(null);
  $__res = function($f_9) use ($Semigroup0_8_8, $dictMonoid_7) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($Semigroup0_8_8, $dictMonoid_7, $f_9) {
  $__num = \func_num_args();
  $Semigroup0_11_9 = (($dictMonoid_7)->{'Semigroup0'})(null);
  $__res = ((($Semigroup0_8_8)->{'append'})(($f_9)(($v_10)->{'value0'})))(\Data\Foldable\majData_majFoldable_foldrmajArray(function($x_12) use ($Semigroup0_11_9, $f_9) {
  $__num = \func_num_args();
  $__res = function($acc_13) use ($Semigroup0_11_9, $f_9, $x_12) {
  $__num = \func_num_args();
  $__res = ((($Semigroup0_11_9)->{'append'})(($f_9)($x_12)))($acc_13);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($dictMonoid_7)->{'mempty'}, ($v_10)->{'value1'}));
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
}, "foldl" => function($f_7) {
  $__num = \func_num_args();
  $__res = function($b_8) use ($f_7) {
  $__num = \func_num_args();
  $__res = function($v_9) use ($b_8, $f_7) {
  $__num = \func_num_args();
  $__res = \Data\Foldable\majData_majFoldable_foldlmajArray($f_7, (($f_7)($b_8))(($v_9)->{'value0'}), ($v_9)->{'value1'});
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
}, "foldr" => function($f_7) {
  $__num = \func_num_args();
  $__res = function($b_8) use ($f_7) {
  $__num = \func_num_args();
  $__res = function($v_9) use ($b_8, $f_7) {
  $__num = \func_num_args();
  $__res = (($f_7)(($v_9)->{'value0'}))(\Data\Foldable\majData_majFoldable_foldrmajArray($f_7, $b_8, ($v_9)->{'value1'}));
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
$__local_var_7_8 = (object)["foldMap1" => function($dictSemigroup_8) {
  $__num = \func_num_args();
  $__res = function($f_9) use ($dictSemigroup_8) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($dictSemigroup_8, $f_9) {
  $__num = \func_num_args();
  $__res = \Data\Foldable\majData_majFoldable_foldlmajArray(function($s_11) use ($dictSemigroup_8, $f_9) {
  $__num = \func_num_args();
  $__res = function($a1_12) use ($dictSemigroup_8, $f_9, $s_11) {
  $__num = \func_num_args();
  $__res = ((($dictSemigroup_8)->{'append'})($s_11))(($f_9)($a1_12));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($f_9)(($v_10)->{'value0'}), ($v_10)->{'value1'});
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
}, "foldr1" => function($f_8) {
  $__num = \func_num_args();
  $__res = function($v_9) use ($f_8) {
  $__num = \func_num_args();
  $__local_var_10_11 = ($f_8)(($v_9)->{'value0'});
  $__local_var_11_12 = \Data\Foldable\majData_majFoldable_foldrmajArray(function($a1_11) use ($f_8) {
  $__num = \func_num_args();
  $__local_var_12_12 = ($f_8)($a1_11);
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Maybe_Just']))(function($v2_13) use ($__local_var_12_12, $a1_11) {
  $__num = \func_num_args();
  $__t13 = null;;
  if ($v2_13 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t13 = $a1_11;
goto end_branch_13;;
};
  if ($v2_13 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t13 = ($__local_var_12_12)(($v2_13)->{'value0'});
goto end_branch_13;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t13 = null;
  end_branch_13:;
  $__res = $__t13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, new \Data\Maybe\Data_Maybe_Nothing(), ($v_9)->{'value1'});
  $__t15 = null;;
  if ($__local_var_11_12 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t15 = ($v_9)->{'value0'};
goto end_branch_15;;
};
  if ($__local_var_11_12 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t15 = ($__local_var_10_11)(($__local_var_11_12)->{'value0'});
goto end_branch_15;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t15 = null;
  end_branch_15:;
  $__res = $__t15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldl1" => function($f_8) {
  $__num = \func_num_args();
  $__res = function($v_9) use ($f_8) {
  $__num = \func_num_args();
  $__res = \Data\Foldable\majData_majFoldable_foldlmajArray($f_8, ($v_9)->{'value0'}, ($v_9)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable0" => function($_dollar___unused_8) use ($foldableNonEmpty1_7_8) {
  $__num = \func_num_args();
  $__res = $foldableNonEmpty1_7_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
$__t22 = null;;
switch (($v_4_3)->{'value0'}) {
case 23:
$__t22 = \Data\Unfoldable1\majData_majUnfoldable1_unfoldr1majArraymajImpl($GLOBALS['Data_Maybe_isNothing'], function($v_8) {
  $__num = \func_num_args();
  $__t23 = null;;
  if ($v_8 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t23 = ($v_8)->{'value0'};
goto end_branch_23;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t23 = null;
  end_branch_23:;
  $__res = $__t23;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $GLOBALS['Data_Tuple_fst'], $GLOBALS['Data_Tuple_snd'], function($i_8) use ($v_4_3) {
  $__num = \func_num_args();
  $__t24 = null;;
  if (($i_8 <= 0)) {
$__t24 = new \Data\Tuple\Data_Tuple_Tuple(($v_4_3)->{'value0'}, new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_24;;
};
  $__t24 = new \Data\Tuple\Data_Tuple_Tuple(($v_4_3)->{'value0'}, new \Data\Maybe\Data_Maybe_Just(($i_8 - 1)));
  end_branch_24:;
  $__res = $__t24;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, 0);
goto end_branch_22;;
break;
default:
;
break;
};
if ((($v_4_3)->{'value0'} < 23)) {
$__t22 = \Data\Unfoldable1\majData_majUnfoldable1_unfoldr1majArraymajImpl($GLOBALS['Data_Maybe_isNothing'], function($v_8) {
  $__num = \func_num_args();
  $__t25 = null;;
  if ($v_8 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t25 = ($v_8)->{'value0'};
goto end_branch_25;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t25 = null;
  end_branch_25:;
  $__res = $__t25;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $GLOBALS['Data_Tuple_fst'], $GLOBALS['Data_Tuple_snd'], function($a_8) {
  $__num = \func_num_args();
  $__local_var_9_26 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_9) {
  $__num = \func_num_args();
  $__t26 = null;;
  if ((($n_9 >= 0) && ($n_9 <= 23))) {
$__t26 = new \Data\Maybe\Data_Maybe_Just($n_9);
goto end_branch_26;;
};
  $__t26 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_26:;
  $__res = $__t26;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_9) {
  $__num = \func_num_args();
  $__res = ($v_9 + 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_9) {
  $__num = \func_num_args();
  $__res = $v_9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), $a_8);
  $__t28 = null;;
  if ($__local_var_9_26 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t29 = null;;
if ((($__local_var_9_26)->{'value0'} <= 23)) {
$__t29 = new \Data\Maybe\Data_Maybe_Just(($__local_var_9_26)->{'value0'});
goto end_branch_29;;
};
$__t29 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_29:;
$__t28 = $__t29;
goto end_branch_28;;
};
  if ($__local_var_9_26 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t28 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_28;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t28 = null;
  end_branch_28:;
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_8, $__t28);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($v_4_3)->{'value0'});
goto end_branch_22;;
};
$__t22 = \Data\Unfoldable1\majData_majUnfoldable1_unfoldr1majArraymajImpl($GLOBALS['Data_Maybe_isNothing'], function($v_8) {
  $__num = \func_num_args();
  $__t17 = null;;
  if ($v_8 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t17 = ($v_8)->{'value0'};
goto end_branch_17;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t17 = null;
  end_branch_17:;
  $__res = $__t17;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $GLOBALS['Data_Tuple_fst'], $GLOBALS['Data_Tuple_snd'], function($a_8) {
  $__num = \func_num_args();
  $__local_var_9_18 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_9) {
  $__num = \func_num_args();
  $__t18 = null;;
  if ((($n_9 >= 0) && ($n_9 <= 23))) {
$__t18 = new \Data\Maybe\Data_Maybe_Just($n_9);
goto end_branch_18;;
};
  $__t18 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_18:;
  $__res = $__t18;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_9) {
  $__num = \func_num_args();
  $__res = ($v_9 - 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_9) {
  $__num = \func_num_args();
  $__res = $v_9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), $a_8);
  $__t20 = null;;
  if ($__local_var_9_18 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t21 = null;;
if ((($__local_var_9_18)->{'value0'} >= 23)) {
$__t21 = new \Data\Maybe\Data_Maybe_Just(($__local_var_9_18)->{'value0'});
goto end_branch_21;;
};
$__t21 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_21:;
$__t20 = $__t21;
goto end_branch_20;;
};
  if ($__local_var_9_18 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t20 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_20;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t20 = null;
  end_branch_20:;
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_8, $__t20);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($v_4_3)->{'value0'});
end_branch_22:;
$__local_var_8_17 = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(0, $__t22);
$__t5 = ((((($Monad0_5_6)->{'Bind1'})(null))->{'bind'})(((($dictMonadGen_0)->{'chooseInt'})(0))(((((((($__local_var_7_8)->{'Foldable0'})(null))->{'foldl'})(function($c_9) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($c_9) {
  $__num = \func_num_args();
  $__res = (1 + $c_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(0))($__local_var_8_17) - 1))))(function($n_9) use ($__local_var_7_8, $__local_var_8_17, $pure_6_7) {
  $__num = \func_num_args();
  $go__go_10_31 = null;
  $go__go_10_31 = (function() use ($__local_var_7_8, $__local_var_8_17, &$go__go_10_31) {
  $__fn = function(int $v_11, $v1_12 = null) use ($__local_var_7_8, $__local_var_8_17, &$go__go_10_31, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_10_31_31_v_11 = $v_11;
  $__tco_var_go__go_10_31_31_v1_12 = $v1_12;
  tco_loop_go__go_10_31_31:;
  $v_11 = $__tco_var_go__go_10_31_31_v_11;
  $v1_12 = $__tco_var_go__go_10_31_31_v1_12;
  $__t31 = null;;
  if ($v1_12 instanceof \Control\Monad\Gen\Control_Monad_Gen_Cons) {
$__t34 = null;;
if (($v1_12)->{'value1'} instanceof \Control\Monad\Gen\Control_Monad_Gen_Nil) {
$__t34 = ($v1_12)->{'value0'};
goto end_branch_34;;
};
if (($v_11 <= 0)) {
$__t34 = ($v1_12)->{'value0'};
goto end_branch_34;;
};
$__tco_32 = ($v_11 - 1);
$__tco_33 = ($v1_12)->{'value1'};
$__tco_var_go__go_10_31_31_v_11 = $__tco_32;
$__tco_var_go__go_10_31_31_v1_12 = $__tco_33;
goto tco_loop_go__go_10_31_31;;
$__t34 = null;
end_branch_34:;
$__t31 = $__t34;
goto end_branch_31;;
};
  if ($v1_12 instanceof \Control\Monad\Gen\Control_Monad_Gen_Nil) {
$__t31 = (((($__local_var_7_8)->{'foldMap1'})($GLOBALS['Data_Semigroup_Last_semigroupLast']))(function($x_13) {
  $__num = \func_num_args();
  $__res = $x_13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($__local_var_8_17);
goto end_branch_31;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t31 = null;
  end_branch_31:;
  $__res = $__t31;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ($pure_6_7)((($go__go_10_31)($n_9))((((((($__local_var_7_8)->{'Foldable0'})(null))->{'foldr'})($GLOBALS['Control_Monad_Gen_Cons']))(new \Control\Monad\Gen\Control_Monad_Gen_Nil()))($__local_var_8_17)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
goto end_branch_5;;
};
  if ($v_4_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t5 = (($Applicative0_3_2)->{'pure'})(0);
goto end_branch_5;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t5 = null;
  end_branch_5:;
  $Applicative0_3_32 = (((($dictMonadGen_0)->{'Monad0'})(null))->{'Applicative0'})(null);
  $v_4_33 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_4) {
  $__num = \func_num_args();
  $__t33 = null;;
  if ((($n_4 >= 0) && ($n_4 <= 59))) {
$__t33 = new \Data\Maybe\Data_Maybe_Just($n_4);
goto end_branch_33;;
};
  $__t33 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_33:;
  $__res = $__t33;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_4) {
  $__num = \func_num_args();
  $__res = ($v_4 + 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_4) {
  $__num = \func_num_args();
  $__res = $v_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), 0);
  $__t35 = null;;
  if ($v_4_33 instanceof \Data\Maybe\Data_Maybe_Just) {
$Monad0_5_36 = (($dictMonadGen_0)->{'Monad0'})(null);
$pure_6_37 = ((($Monad0_5_36)->{'Applicative0'})(null))->{'pure'};
$foldableNonEmpty1_7_38 = (object)["foldMap" => function($dictMonoid_7) {
  $__num = \func_num_args();
  $Semigroup0_8_38 = (($dictMonoid_7)->{'Semigroup0'})(null);
  $__res = function($f_9) use ($Semigroup0_8_38, $dictMonoid_7) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($Semigroup0_8_38, $dictMonoid_7, $f_9) {
  $__num = \func_num_args();
  $Semigroup0_11_39 = (($dictMonoid_7)->{'Semigroup0'})(null);
  $__res = ((($Semigroup0_8_38)->{'append'})(($f_9)(($v_10)->{'value0'})))(\Data\Foldable\majData_majFoldable_foldrmajArray(function($x_12) use ($Semigroup0_11_39, $f_9) {
  $__num = \func_num_args();
  $__res = function($acc_13) use ($Semigroup0_11_39, $f_9, $x_12) {
  $__num = \func_num_args();
  $__res = ((($Semigroup0_11_39)->{'append'})(($f_9)($x_12)))($acc_13);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($dictMonoid_7)->{'mempty'}, ($v_10)->{'value1'}));
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
}, "foldl" => function($f_7) {
  $__num = \func_num_args();
  $__res = function($b_8) use ($f_7) {
  $__num = \func_num_args();
  $__res = function($v_9) use ($b_8, $f_7) {
  $__num = \func_num_args();
  $__res = \Data\Foldable\majData_majFoldable_foldlmajArray($f_7, (($f_7)($b_8))(($v_9)->{'value0'}), ($v_9)->{'value1'});
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
}, "foldr" => function($f_7) {
  $__num = \func_num_args();
  $__res = function($b_8) use ($f_7) {
  $__num = \func_num_args();
  $__res = function($v_9) use ($b_8, $f_7) {
  $__num = \func_num_args();
  $__res = (($f_7)(($v_9)->{'value0'}))(\Data\Foldable\majData_majFoldable_foldrmajArray($f_7, $b_8, ($v_9)->{'value1'}));
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
$__local_var_7_38 = (object)["foldMap1" => function($dictSemigroup_8) {
  $__num = \func_num_args();
  $__res = function($f_9) use ($dictSemigroup_8) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($dictSemigroup_8, $f_9) {
  $__num = \func_num_args();
  $__res = \Data\Foldable\majData_majFoldable_foldlmajArray(function($s_11) use ($dictSemigroup_8, $f_9) {
  $__num = \func_num_args();
  $__res = function($a1_12) use ($dictSemigroup_8, $f_9, $s_11) {
  $__num = \func_num_args();
  $__res = ((($dictSemigroup_8)->{'append'})($s_11))(($f_9)($a1_12));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($f_9)(($v_10)->{'value0'}), ($v_10)->{'value1'});
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
}, "foldr1" => function($f_8) {
  $__num = \func_num_args();
  $__res = function($v_9) use ($f_8) {
  $__num = \func_num_args();
  $__local_var_10_41 = ($f_8)(($v_9)->{'value0'});
  $__local_var_11_42 = \Data\Foldable\majData_majFoldable_foldrmajArray(function($a1_11) use ($f_8) {
  $__num = \func_num_args();
  $__local_var_12_42 = ($f_8)($a1_11);
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Maybe_Just']))(function($v2_13) use ($__local_var_12_42, $a1_11) {
  $__num = \func_num_args();
  $__t43 = null;;
  if ($v2_13 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t43 = $a1_11;
goto end_branch_43;;
};
  if ($v2_13 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t43 = ($__local_var_12_42)(($v2_13)->{'value0'});
goto end_branch_43;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t43 = null;
  end_branch_43:;
  $__res = $__t43;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, new \Data\Maybe\Data_Maybe_Nothing(), ($v_9)->{'value1'});
  $__t45 = null;;
  if ($__local_var_11_42 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t45 = ($v_9)->{'value0'};
goto end_branch_45;;
};
  if ($__local_var_11_42 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t45 = ($__local_var_10_41)(($__local_var_11_42)->{'value0'});
goto end_branch_45;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t45 = null;
  end_branch_45:;
  $__res = $__t45;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldl1" => function($f_8) {
  $__num = \func_num_args();
  $__res = function($v_9) use ($f_8) {
  $__num = \func_num_args();
  $__res = \Data\Foldable\majData_majFoldable_foldlmajArray($f_8, ($v_9)->{'value0'}, ($v_9)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable0" => function($_dollar___unused_8) use ($foldableNonEmpty1_7_38) {
  $__num = \func_num_args();
  $__res = $foldableNonEmpty1_7_38;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
$__t52 = null;;
switch (($v_4_33)->{'value0'}) {
case 59:
$__t52 = \Data\Unfoldable1\majData_majUnfoldable1_unfoldr1majArraymajImpl($GLOBALS['Data_Maybe_isNothing'], function($v_8) {
  $__num = \func_num_args();
  $__t53 = null;;
  if ($v_8 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t53 = ($v_8)->{'value0'};
goto end_branch_53;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t53 = null;
  end_branch_53:;
  $__res = $__t53;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $GLOBALS['Data_Tuple_fst'], $GLOBALS['Data_Tuple_snd'], function($i_8) use ($v_4_33) {
  $__num = \func_num_args();
  $__t54 = null;;
  if (($i_8 <= 0)) {
$__t54 = new \Data\Tuple\Data_Tuple_Tuple(($v_4_33)->{'value0'}, new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_54;;
};
  $__t54 = new \Data\Tuple\Data_Tuple_Tuple(($v_4_33)->{'value0'}, new \Data\Maybe\Data_Maybe_Just(($i_8 - 1)));
  end_branch_54:;
  $__res = $__t54;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, 0);
goto end_branch_52;;
break;
default:
;
break;
};
if ((($v_4_33)->{'value0'} < 59)) {
$__t52 = \Data\Unfoldable1\majData_majUnfoldable1_unfoldr1majArraymajImpl($GLOBALS['Data_Maybe_isNothing'], function($v_8) {
  $__num = \func_num_args();
  $__t55 = null;;
  if ($v_8 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t55 = ($v_8)->{'value0'};
goto end_branch_55;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t55 = null;
  end_branch_55:;
  $__res = $__t55;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $GLOBALS['Data_Tuple_fst'], $GLOBALS['Data_Tuple_snd'], function($a_8) {
  $__num = \func_num_args();
  $__local_var_9_56 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_9) {
  $__num = \func_num_args();
  $__t56 = null;;
  if ((($n_9 >= 0) && ($n_9 <= 59))) {
$__t56 = new \Data\Maybe\Data_Maybe_Just($n_9);
goto end_branch_56;;
};
  $__t56 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_56:;
  $__res = $__t56;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_9) {
  $__num = \func_num_args();
  $__res = ($v_9 + 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_9) {
  $__num = \func_num_args();
  $__res = $v_9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), $a_8);
  $__t58 = null;;
  if ($__local_var_9_56 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t59 = null;;
if ((($__local_var_9_56)->{'value0'} <= 59)) {
$__t59 = new \Data\Maybe\Data_Maybe_Just(($__local_var_9_56)->{'value0'});
goto end_branch_59;;
};
$__t59 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_59:;
$__t58 = $__t59;
goto end_branch_58;;
};
  if ($__local_var_9_56 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t58 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_58;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t58 = null;
  end_branch_58:;
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_8, $__t58);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($v_4_33)->{'value0'});
goto end_branch_52;;
};
$__t52 = \Data\Unfoldable1\majData_majUnfoldable1_unfoldr1majArraymajImpl($GLOBALS['Data_Maybe_isNothing'], function($v_8) {
  $__num = \func_num_args();
  $__t47 = null;;
  if ($v_8 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t47 = ($v_8)->{'value0'};
goto end_branch_47;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t47 = null;
  end_branch_47:;
  $__res = $__t47;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $GLOBALS['Data_Tuple_fst'], $GLOBALS['Data_Tuple_snd'], function($a_8) {
  $__num = \func_num_args();
  $__local_var_9_48 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_9) {
  $__num = \func_num_args();
  $__t48 = null;;
  if ((($n_9 >= 0) && ($n_9 <= 59))) {
$__t48 = new \Data\Maybe\Data_Maybe_Just($n_9);
goto end_branch_48;;
};
  $__t48 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_48:;
  $__res = $__t48;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_9) {
  $__num = \func_num_args();
  $__res = ($v_9 - 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_9) {
  $__num = \func_num_args();
  $__res = $v_9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), $a_8);
  $__t50 = null;;
  if ($__local_var_9_48 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t51 = null;;
if ((($__local_var_9_48)->{'value0'} >= 59)) {
$__t51 = new \Data\Maybe\Data_Maybe_Just(($__local_var_9_48)->{'value0'});
goto end_branch_51;;
};
$__t51 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_51:;
$__t50 = $__t51;
goto end_branch_50;;
};
  if ($__local_var_9_48 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t50 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_50;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t50 = null;
  end_branch_50:;
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_8, $__t50);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($v_4_33)->{'value0'});
end_branch_52:;
$__local_var_8_47 = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(0, $__t52);
$__t35 = ((((($Monad0_5_36)->{'Bind1'})(null))->{'bind'})(((($dictMonadGen_0)->{'chooseInt'})(0))(((((((($__local_var_7_38)->{'Foldable0'})(null))->{'foldl'})(function($c_9) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($c_9) {
  $__num = \func_num_args();
  $__res = (1 + $c_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(0))($__local_var_8_47) - 1))))(function($n_9) use ($__local_var_7_38, $__local_var_8_47, $pure_6_37) {
  $__num = \func_num_args();
  $go__go_10_61 = null;
  $go__go_10_61 = (function() use ($__local_var_7_38, $__local_var_8_47, &$go__go_10_61) {
  $__fn = function(int $v_11, $v1_12 = null) use ($__local_var_7_38, $__local_var_8_47, &$go__go_10_61, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_10_61_61_v_11 = $v_11;
  $__tco_var_go__go_10_61_61_v1_12 = $v1_12;
  tco_loop_go__go_10_61_61:;
  $v_11 = $__tco_var_go__go_10_61_61_v_11;
  $v1_12 = $__tco_var_go__go_10_61_61_v1_12;
  $__t61 = null;;
  if ($v1_12 instanceof \Control\Monad\Gen\Control_Monad_Gen_Cons) {
$__t64 = null;;
if (($v1_12)->{'value1'} instanceof \Control\Monad\Gen\Control_Monad_Gen_Nil) {
$__t64 = ($v1_12)->{'value0'};
goto end_branch_64;;
};
if (($v_11 <= 0)) {
$__t64 = ($v1_12)->{'value0'};
goto end_branch_64;;
};
$__tco_62 = ($v_11 - 1);
$__tco_63 = ($v1_12)->{'value1'};
$__tco_var_go__go_10_61_61_v_11 = $__tco_62;
$__tco_var_go__go_10_61_61_v1_12 = $__tco_63;
goto tco_loop_go__go_10_61_61;;
$__t64 = null;
end_branch_64:;
$__t61 = $__t64;
goto end_branch_61;;
};
  if ($v1_12 instanceof \Control\Monad\Gen\Control_Monad_Gen_Nil) {
$__t61 = (((($__local_var_7_38)->{'foldMap1'})($GLOBALS['Data_Semigroup_Last_semigroupLast']))(function($x_13) {
  $__num = \func_num_args();
  $__res = $x_13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($__local_var_8_47);
goto end_branch_61;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t61 = null;
  end_branch_61:;
  $__res = $__t61;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ($pure_6_37)((($go__go_10_61)($n_9))((((((($__local_var_7_38)->{'Foldable0'})(null))->{'foldr'})($GLOBALS['Control_Monad_Gen_Cons']))(new \Control\Monad\Gen\Control_Monad_Gen_Nil()))($__local_var_8_47)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
goto end_branch_35;;
};
  if ($v_4_33 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t35 = (($Applicative0_3_32)->{'pure'})(0);
goto end_branch_35;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t35 = null;
  end_branch_35:;
  $Applicative0_3_62 = (((($dictMonadGen_0)->{'Monad0'})(null))->{'Applicative0'})(null);
  $v_4_63 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_4) {
  $__num = \func_num_args();
  $__t63 = null;;
  if ((($n_4 >= 0) && ($n_4 <= 59))) {
$__t63 = new \Data\Maybe\Data_Maybe_Just($n_4);
goto end_branch_63;;
};
  $__t63 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_63:;
  $__res = $__t63;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_4) {
  $__num = \func_num_args();
  $__res = ($v_4 + 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_4) {
  $__num = \func_num_args();
  $__res = $v_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), 0);
  $__t65 = null;;
  if ($v_4_63 instanceof \Data\Maybe\Data_Maybe_Just) {
$Monad0_5_66 = (($dictMonadGen_0)->{'Monad0'})(null);
$pure_6_67 = ((($Monad0_5_66)->{'Applicative0'})(null))->{'pure'};
$foldableNonEmpty1_7_68 = (object)["foldMap" => function($dictMonoid_7) {
  $__num = \func_num_args();
  $Semigroup0_8_68 = (($dictMonoid_7)->{'Semigroup0'})(null);
  $__res = function($f_9) use ($Semigroup0_8_68, $dictMonoid_7) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($Semigroup0_8_68, $dictMonoid_7, $f_9) {
  $__num = \func_num_args();
  $Semigroup0_11_69 = (($dictMonoid_7)->{'Semigroup0'})(null);
  $__res = ((($Semigroup0_8_68)->{'append'})(($f_9)(($v_10)->{'value0'})))(\Data\Foldable\majData_majFoldable_foldrmajArray(function($x_12) use ($Semigroup0_11_69, $f_9) {
  $__num = \func_num_args();
  $__res = function($acc_13) use ($Semigroup0_11_69, $f_9, $x_12) {
  $__num = \func_num_args();
  $__res = ((($Semigroup0_11_69)->{'append'})(($f_9)($x_12)))($acc_13);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($dictMonoid_7)->{'mempty'}, ($v_10)->{'value1'}));
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
}, "foldl" => function($f_7) {
  $__num = \func_num_args();
  $__res = function($b_8) use ($f_7) {
  $__num = \func_num_args();
  $__res = function($v_9) use ($b_8, $f_7) {
  $__num = \func_num_args();
  $__res = \Data\Foldable\majData_majFoldable_foldlmajArray($f_7, (($f_7)($b_8))(($v_9)->{'value0'}), ($v_9)->{'value1'});
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
}, "foldr" => function($f_7) {
  $__num = \func_num_args();
  $__res = function($b_8) use ($f_7) {
  $__num = \func_num_args();
  $__res = function($v_9) use ($b_8, $f_7) {
  $__num = \func_num_args();
  $__res = (($f_7)(($v_9)->{'value0'}))(\Data\Foldable\majData_majFoldable_foldrmajArray($f_7, $b_8, ($v_9)->{'value1'}));
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
$__local_var_7_68 = (object)["foldMap1" => function($dictSemigroup_8) {
  $__num = \func_num_args();
  $__res = function($f_9) use ($dictSemigroup_8) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($dictSemigroup_8, $f_9) {
  $__num = \func_num_args();
  $__res = \Data\Foldable\majData_majFoldable_foldlmajArray(function($s_11) use ($dictSemigroup_8, $f_9) {
  $__num = \func_num_args();
  $__res = function($a1_12) use ($dictSemigroup_8, $f_9, $s_11) {
  $__num = \func_num_args();
  $__res = ((($dictSemigroup_8)->{'append'})($s_11))(($f_9)($a1_12));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($f_9)(($v_10)->{'value0'}), ($v_10)->{'value1'});
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
}, "foldr1" => function($f_8) {
  $__num = \func_num_args();
  $__res = function($v_9) use ($f_8) {
  $__num = \func_num_args();
  $__local_var_10_71 = ($f_8)(($v_9)->{'value0'});
  $__local_var_11_72 = \Data\Foldable\majData_majFoldable_foldrmajArray(function($a1_11) use ($f_8) {
  $__num = \func_num_args();
  $__local_var_12_72 = ($f_8)($a1_11);
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Maybe_Just']))(function($v2_13) use ($__local_var_12_72, $a1_11) {
  $__num = \func_num_args();
  $__t73 = null;;
  if ($v2_13 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t73 = $a1_11;
goto end_branch_73;;
};
  if ($v2_13 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t73 = ($__local_var_12_72)(($v2_13)->{'value0'});
goto end_branch_73;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t73 = null;
  end_branch_73:;
  $__res = $__t73;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, new \Data\Maybe\Data_Maybe_Nothing(), ($v_9)->{'value1'});
  $__t75 = null;;
  if ($__local_var_11_72 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t75 = ($v_9)->{'value0'};
goto end_branch_75;;
};
  if ($__local_var_11_72 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t75 = ($__local_var_10_71)(($__local_var_11_72)->{'value0'});
goto end_branch_75;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t75 = null;
  end_branch_75:;
  $__res = $__t75;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldl1" => function($f_8) {
  $__num = \func_num_args();
  $__res = function($v_9) use ($f_8) {
  $__num = \func_num_args();
  $__res = \Data\Foldable\majData_majFoldable_foldlmajArray($f_8, ($v_9)->{'value0'}, ($v_9)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable0" => function($_dollar___unused_8) use ($foldableNonEmpty1_7_68) {
  $__num = \func_num_args();
  $__res = $foldableNonEmpty1_7_68;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
$__t82 = null;;
switch (($v_4_63)->{'value0'}) {
case 59:
$__t82 = \Data\Unfoldable1\majData_majUnfoldable1_unfoldr1majArraymajImpl($GLOBALS['Data_Maybe_isNothing'], function($v_8) {
  $__num = \func_num_args();
  $__t83 = null;;
  if ($v_8 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t83 = ($v_8)->{'value0'};
goto end_branch_83;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t83 = null;
  end_branch_83:;
  $__res = $__t83;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $GLOBALS['Data_Tuple_fst'], $GLOBALS['Data_Tuple_snd'], function($i_8) use ($v_4_63) {
  $__num = \func_num_args();
  $__t84 = null;;
  if (($i_8 <= 0)) {
$__t84 = new \Data\Tuple\Data_Tuple_Tuple(($v_4_63)->{'value0'}, new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_84;;
};
  $__t84 = new \Data\Tuple\Data_Tuple_Tuple(($v_4_63)->{'value0'}, new \Data\Maybe\Data_Maybe_Just(($i_8 - 1)));
  end_branch_84:;
  $__res = $__t84;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, 0);
goto end_branch_82;;
break;
default:
;
break;
};
if ((($v_4_63)->{'value0'} < 59)) {
$__t82 = \Data\Unfoldable1\majData_majUnfoldable1_unfoldr1majArraymajImpl($GLOBALS['Data_Maybe_isNothing'], function($v_8) {
  $__num = \func_num_args();
  $__t85 = null;;
  if ($v_8 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t85 = ($v_8)->{'value0'};
goto end_branch_85;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t85 = null;
  end_branch_85:;
  $__res = $__t85;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $GLOBALS['Data_Tuple_fst'], $GLOBALS['Data_Tuple_snd'], function($a_8) {
  $__num = \func_num_args();
  $__local_var_9_86 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_9) {
  $__num = \func_num_args();
  $__t86 = null;;
  if ((($n_9 >= 0) && ($n_9 <= 59))) {
$__t86 = new \Data\Maybe\Data_Maybe_Just($n_9);
goto end_branch_86;;
};
  $__t86 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_86:;
  $__res = $__t86;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_9) {
  $__num = \func_num_args();
  $__res = ($v_9 + 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_9) {
  $__num = \func_num_args();
  $__res = $v_9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), $a_8);
  $__t88 = null;;
  if ($__local_var_9_86 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t89 = null;;
if ((($__local_var_9_86)->{'value0'} <= 59)) {
$__t89 = new \Data\Maybe\Data_Maybe_Just(($__local_var_9_86)->{'value0'});
goto end_branch_89;;
};
$__t89 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_89:;
$__t88 = $__t89;
goto end_branch_88;;
};
  if ($__local_var_9_86 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t88 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_88;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t88 = null;
  end_branch_88:;
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_8, $__t88);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($v_4_63)->{'value0'});
goto end_branch_82;;
};
$__t82 = \Data\Unfoldable1\majData_majUnfoldable1_unfoldr1majArraymajImpl($GLOBALS['Data_Maybe_isNothing'], function($v_8) {
  $__num = \func_num_args();
  $__t77 = null;;
  if ($v_8 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t77 = ($v_8)->{'value0'};
goto end_branch_77;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t77 = null;
  end_branch_77:;
  $__res = $__t77;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $GLOBALS['Data_Tuple_fst'], $GLOBALS['Data_Tuple_snd'], function($a_8) {
  $__num = \func_num_args();
  $__local_var_9_78 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_9) {
  $__num = \func_num_args();
  $__t78 = null;;
  if ((($n_9 >= 0) && ($n_9 <= 59))) {
$__t78 = new \Data\Maybe\Data_Maybe_Just($n_9);
goto end_branch_78;;
};
  $__t78 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_78:;
  $__res = $__t78;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_9) {
  $__num = \func_num_args();
  $__res = ($v_9 - 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_9) {
  $__num = \func_num_args();
  $__res = $v_9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), $a_8);
  $__t80 = null;;
  if ($__local_var_9_78 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t81 = null;;
if ((($__local_var_9_78)->{'value0'} >= 59)) {
$__t81 = new \Data\Maybe\Data_Maybe_Just(($__local_var_9_78)->{'value0'});
goto end_branch_81;;
};
$__t81 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_81:;
$__t80 = $__t81;
goto end_branch_80;;
};
  if ($__local_var_9_78 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t80 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_80;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t80 = null;
  end_branch_80:;
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_8, $__t80);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($v_4_63)->{'value0'});
end_branch_82:;
$__local_var_8_77 = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(0, $__t82);
$__t65 = ((((($Monad0_5_66)->{'Bind1'})(null))->{'bind'})(((($dictMonadGen_0)->{'chooseInt'})(0))(((((((($__local_var_7_68)->{'Foldable0'})(null))->{'foldl'})(function($c_9) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($c_9) {
  $__num = \func_num_args();
  $__res = (1 + $c_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(0))($__local_var_8_77) - 1))))(function($n_9) use ($__local_var_7_68, $__local_var_8_77, $pure_6_67) {
  $__num = \func_num_args();
  $go__go_10_91 = null;
  $go__go_10_91 = (function() use ($__local_var_7_68, $__local_var_8_77, &$go__go_10_91) {
  $__fn = function(int $v_11, $v1_12 = null) use ($__local_var_7_68, $__local_var_8_77, &$go__go_10_91, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_10_91_91_v_11 = $v_11;
  $__tco_var_go__go_10_91_91_v1_12 = $v1_12;
  tco_loop_go__go_10_91_91:;
  $v_11 = $__tco_var_go__go_10_91_91_v_11;
  $v1_12 = $__tco_var_go__go_10_91_91_v1_12;
  $__t91 = null;;
  if ($v1_12 instanceof \Control\Monad\Gen\Control_Monad_Gen_Cons) {
$__t94 = null;;
if (($v1_12)->{'value1'} instanceof \Control\Monad\Gen\Control_Monad_Gen_Nil) {
$__t94 = ($v1_12)->{'value0'};
goto end_branch_94;;
};
if (($v_11 <= 0)) {
$__t94 = ($v1_12)->{'value0'};
goto end_branch_94;;
};
$__tco_92 = ($v_11 - 1);
$__tco_93 = ($v1_12)->{'value1'};
$__tco_var_go__go_10_91_91_v_11 = $__tco_92;
$__tco_var_go__go_10_91_91_v1_12 = $__tco_93;
goto tco_loop_go__go_10_91_91;;
$__t94 = null;
end_branch_94:;
$__t91 = $__t94;
goto end_branch_91;;
};
  if ($v1_12 instanceof \Control\Monad\Gen\Control_Monad_Gen_Nil) {
$__t91 = (((($__local_var_7_68)->{'foldMap1'})($GLOBALS['Data_Semigroup_Last_semigroupLast']))(function($x_13) {
  $__num = \func_num_args();
  $__res = $x_13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($__local_var_8_77);
goto end_branch_91;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t91 = null;
  end_branch_91:;
  $__res = $__t91;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ($pure_6_67)((($go__go_10_91)($n_9))((((((($__local_var_7_68)->{'Foldable0'})(null))->{'foldr'})($GLOBALS['Control_Monad_Gen_Cons']))(new \Control\Monad\Gen\Control_Monad_Gen_Nil()))($__local_var_8_77)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
goto end_branch_65;;
};
  if ($v_4_63 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t65 = (($Applicative0_3_62)->{'pure'})(0);
goto end_branch_65;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t65 = null;
  end_branch_65:;
  $Applicative0_3_92 = (((($dictMonadGen_0)->{'Monad0'})(null))->{'Applicative0'})(null);
  $v_4_93 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_4) {
  $__num = \func_num_args();
  $__t93 = null;;
  if ((($n_4 >= 0) && ($n_4 <= 999))) {
$__t93 = new \Data\Maybe\Data_Maybe_Just($n_4);
goto end_branch_93;;
};
  $__t93 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_93:;
  $__res = $__t93;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_4) {
  $__num = \func_num_args();
  $__res = ($v_4 + 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_4) {
  $__num = \func_num_args();
  $__res = $v_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), 0);
  $__t95 = null;;
  if ($v_4_93 instanceof \Data\Maybe\Data_Maybe_Just) {
$Monad0_5_96 = (($dictMonadGen_0)->{'Monad0'})(null);
$pure_6_97 = ((($Monad0_5_96)->{'Applicative0'})(null))->{'pure'};
$foldableNonEmpty1_7_98 = (object)["foldMap" => function($dictMonoid_7) {
  $__num = \func_num_args();
  $Semigroup0_8_98 = (($dictMonoid_7)->{'Semigroup0'})(null);
  $__res = function($f_9) use ($Semigroup0_8_98, $dictMonoid_7) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($Semigroup0_8_98, $dictMonoid_7, $f_9) {
  $__num = \func_num_args();
  $Semigroup0_11_99 = (($dictMonoid_7)->{'Semigroup0'})(null);
  $__res = ((($Semigroup0_8_98)->{'append'})(($f_9)(($v_10)->{'value0'})))(\Data\Foldable\majData_majFoldable_foldrmajArray(function($x_12) use ($Semigroup0_11_99, $f_9) {
  $__num = \func_num_args();
  $__res = function($acc_13) use ($Semigroup0_11_99, $f_9, $x_12) {
  $__num = \func_num_args();
  $__res = ((($Semigroup0_11_99)->{'append'})(($f_9)($x_12)))($acc_13);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($dictMonoid_7)->{'mempty'}, ($v_10)->{'value1'}));
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
}, "foldl" => function($f_7) {
  $__num = \func_num_args();
  $__res = function($b_8) use ($f_7) {
  $__num = \func_num_args();
  $__res = function($v_9) use ($b_8, $f_7) {
  $__num = \func_num_args();
  $__res = \Data\Foldable\majData_majFoldable_foldlmajArray($f_7, (($f_7)($b_8))(($v_9)->{'value0'}), ($v_9)->{'value1'});
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
}, "foldr" => function($f_7) {
  $__num = \func_num_args();
  $__res = function($b_8) use ($f_7) {
  $__num = \func_num_args();
  $__res = function($v_9) use ($b_8, $f_7) {
  $__num = \func_num_args();
  $__res = (($f_7)(($v_9)->{'value0'}))(\Data\Foldable\majData_majFoldable_foldrmajArray($f_7, $b_8, ($v_9)->{'value1'}));
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
$__local_var_7_98 = (object)["foldMap1" => function($dictSemigroup_8) {
  $__num = \func_num_args();
  $__res = function($f_9) use ($dictSemigroup_8) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($dictSemigroup_8, $f_9) {
  $__num = \func_num_args();
  $__res = \Data\Foldable\majData_majFoldable_foldlmajArray(function($s_11) use ($dictSemigroup_8, $f_9) {
  $__num = \func_num_args();
  $__res = function($a1_12) use ($dictSemigroup_8, $f_9, $s_11) {
  $__num = \func_num_args();
  $__res = ((($dictSemigroup_8)->{'append'})($s_11))(($f_9)($a1_12));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($f_9)(($v_10)->{'value0'}), ($v_10)->{'value1'});
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
}, "foldr1" => function($f_8) {
  $__num = \func_num_args();
  $__res = function($v_9) use ($f_8) {
  $__num = \func_num_args();
  $__local_var_10_101 = ($f_8)(($v_9)->{'value0'});
  $__local_var_11_102 = \Data\Foldable\majData_majFoldable_foldrmajArray(function($a1_11) use ($f_8) {
  $__num = \func_num_args();
  $__local_var_12_102 = ($f_8)($a1_11);
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Maybe_Just']))(function($v2_13) use ($__local_var_12_102, $a1_11) {
  $__num = \func_num_args();
  $__t103 = null;;
  if ($v2_13 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t103 = $a1_11;
goto end_branch_103;;
};
  if ($v2_13 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t103 = ($__local_var_12_102)(($v2_13)->{'value0'});
goto end_branch_103;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t103 = null;
  end_branch_103:;
  $__res = $__t103;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, new \Data\Maybe\Data_Maybe_Nothing(), ($v_9)->{'value1'});
  $__t105 = null;;
  if ($__local_var_11_102 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t105 = ($v_9)->{'value0'};
goto end_branch_105;;
};
  if ($__local_var_11_102 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t105 = ($__local_var_10_101)(($__local_var_11_102)->{'value0'});
goto end_branch_105;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t105 = null;
  end_branch_105:;
  $__res = $__t105;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldl1" => function($f_8) {
  $__num = \func_num_args();
  $__res = function($v_9) use ($f_8) {
  $__num = \func_num_args();
  $__res = \Data\Foldable\majData_majFoldable_foldlmajArray($f_8, ($v_9)->{'value0'}, ($v_9)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable0" => function($_dollar___unused_8) use ($foldableNonEmpty1_7_98) {
  $__num = \func_num_args();
  $__res = $foldableNonEmpty1_7_98;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
$__t112 = null;;
switch (($v_4_93)->{'value0'}) {
case 999:
$__t112 = \Data\Unfoldable1\majData_majUnfoldable1_unfoldr1majArraymajImpl($GLOBALS['Data_Maybe_isNothing'], function($v_8) {
  $__num = \func_num_args();
  $__t113 = null;;
  if ($v_8 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t113 = ($v_8)->{'value0'};
goto end_branch_113;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t113 = null;
  end_branch_113:;
  $__res = $__t113;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $GLOBALS['Data_Tuple_fst'], $GLOBALS['Data_Tuple_snd'], function($i_8) use ($v_4_93) {
  $__num = \func_num_args();
  $__t114 = null;;
  if (($i_8 <= 0)) {
$__t114 = new \Data\Tuple\Data_Tuple_Tuple(($v_4_93)->{'value0'}, new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_114;;
};
  $__t114 = new \Data\Tuple\Data_Tuple_Tuple(($v_4_93)->{'value0'}, new \Data\Maybe\Data_Maybe_Just(($i_8 - 1)));
  end_branch_114:;
  $__res = $__t114;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, 0);
goto end_branch_112;;
break;
default:
;
break;
};
if ((($v_4_93)->{'value0'} < 999)) {
$__t112 = \Data\Unfoldable1\majData_majUnfoldable1_unfoldr1majArraymajImpl($GLOBALS['Data_Maybe_isNothing'], function($v_8) {
  $__num = \func_num_args();
  $__t115 = null;;
  if ($v_8 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t115 = ($v_8)->{'value0'};
goto end_branch_115;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t115 = null;
  end_branch_115:;
  $__res = $__t115;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $GLOBALS['Data_Tuple_fst'], $GLOBALS['Data_Tuple_snd'], function($a_8) {
  $__num = \func_num_args();
  $__local_var_9_116 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_9) {
  $__num = \func_num_args();
  $__t116 = null;;
  if ((($n_9 >= 0) && ($n_9 <= 999))) {
$__t116 = new \Data\Maybe\Data_Maybe_Just($n_9);
goto end_branch_116;;
};
  $__t116 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_116:;
  $__res = $__t116;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_9) {
  $__num = \func_num_args();
  $__res = ($v_9 + 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_9) {
  $__num = \func_num_args();
  $__res = $v_9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), $a_8);
  $__t118 = null;;
  if ($__local_var_9_116 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t119 = null;;
if ((($__local_var_9_116)->{'value0'} <= 999)) {
$__t119 = new \Data\Maybe\Data_Maybe_Just(($__local_var_9_116)->{'value0'});
goto end_branch_119;;
};
$__t119 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_119:;
$__t118 = $__t119;
goto end_branch_118;;
};
  if ($__local_var_9_116 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t118 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_118;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t118 = null;
  end_branch_118:;
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_8, $__t118);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($v_4_93)->{'value0'});
goto end_branch_112;;
};
$__t112 = \Data\Unfoldable1\majData_majUnfoldable1_unfoldr1majArraymajImpl($GLOBALS['Data_Maybe_isNothing'], function($v_8) {
  $__num = \func_num_args();
  $__t107 = null;;
  if ($v_8 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t107 = ($v_8)->{'value0'};
goto end_branch_107;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t107 = null;
  end_branch_107:;
  $__res = $__t107;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $GLOBALS['Data_Tuple_fst'], $GLOBALS['Data_Tuple_snd'], function($a_8) {
  $__num = \func_num_args();
  $__local_var_9_108 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_9) {
  $__num = \func_num_args();
  $__t108 = null;;
  if ((($n_9 >= 0) && ($n_9 <= 999))) {
$__t108 = new \Data\Maybe\Data_Maybe_Just($n_9);
goto end_branch_108;;
};
  $__t108 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_108:;
  $__res = $__t108;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_9) {
  $__num = \func_num_args();
  $__res = ($v_9 - 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_9) {
  $__num = \func_num_args();
  $__res = $v_9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), $a_8);
  $__t110 = null;;
  if ($__local_var_9_108 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t111 = null;;
if ((($__local_var_9_108)->{'value0'} >= 999)) {
$__t111 = new \Data\Maybe\Data_Maybe_Just(($__local_var_9_108)->{'value0'});
goto end_branch_111;;
};
$__t111 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_111:;
$__t110 = $__t111;
goto end_branch_110;;
};
  if ($__local_var_9_108 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t110 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_110;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t110 = null;
  end_branch_110:;
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_8, $__t110);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($v_4_93)->{'value0'});
end_branch_112:;
$__local_var_8_107 = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(0, $__t112);
$__t95 = ((((($Monad0_5_96)->{'Bind1'})(null))->{'bind'})(((($dictMonadGen_0)->{'chooseInt'})(0))(((((((($__local_var_7_98)->{'Foldable0'})(null))->{'foldl'})(function($c_9) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($c_9) {
  $__num = \func_num_args();
  $__res = (1 + $c_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(0))($__local_var_8_107) - 1))))(function($n_9) use ($__local_var_7_98, $__local_var_8_107, $pure_6_97) {
  $__num = \func_num_args();
  $go__go_10_121 = null;
  $go__go_10_121 = (function() use ($__local_var_7_98, $__local_var_8_107, &$go__go_10_121) {
  $__fn = function(int $v_11, $v1_12 = null) use ($__local_var_7_98, $__local_var_8_107, &$go__go_10_121, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_10_121_121_v_11 = $v_11;
  $__tco_var_go__go_10_121_121_v1_12 = $v1_12;
  tco_loop_go__go_10_121_121:;
  $v_11 = $__tco_var_go__go_10_121_121_v_11;
  $v1_12 = $__tco_var_go__go_10_121_121_v1_12;
  $__t121 = null;;
  if ($v1_12 instanceof \Control\Monad\Gen\Control_Monad_Gen_Cons) {
$__t124 = null;;
if (($v1_12)->{'value1'} instanceof \Control\Monad\Gen\Control_Monad_Gen_Nil) {
$__t124 = ($v1_12)->{'value0'};
goto end_branch_124;;
};
if (($v_11 <= 0)) {
$__t124 = ($v1_12)->{'value0'};
goto end_branch_124;;
};
$__tco_122 = ($v_11 - 1);
$__tco_123 = ($v1_12)->{'value1'};
$__tco_var_go__go_10_121_121_v_11 = $__tco_122;
$__tco_var_go__go_10_121_121_v1_12 = $__tco_123;
goto tco_loop_go__go_10_121_121;;
$__t124 = null;
end_branch_124:;
$__t121 = $__t124;
goto end_branch_121;;
};
  if ($v1_12 instanceof \Control\Monad\Gen\Control_Monad_Gen_Nil) {
$__t121 = (((($__local_var_7_98)->{'foldMap1'})($GLOBALS['Data_Semigroup_Last_semigroupLast']))(function($x_13) {
  $__num = \func_num_args();
  $__res = $x_13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($__local_var_8_107);
goto end_branch_121;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t121 = null;
  end_branch_121:;
  $__res = $__t121;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ($pure_6_97)((($go__go_10_121)($n_9))((((((($__local_var_7_98)->{'Foldable0'})(null))->{'foldr'})($GLOBALS['Control_Monad_Gen_Cons']))(new \Control\Monad\Gen\Control_Monad_Gen_Nil()))($__local_var_8_107)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
goto end_branch_95;;
};
  if ($v_4_93 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t95 = (($Applicative0_3_92)->{'pure'})(0);
goto end_branch_95;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t95 = null;
  end_branch_95:;
  $__res = ((($Apply0_2_1)->{'apply'})(((($Apply0_2_1)->{'apply'})(((($Apply0_2_1)->{'apply'})(((((((($Bind1_1_0)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Time_Time']))($__t5)))($__t35)))($__t65)))($__t95);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Time_Gen_genTime'] = __NAMESPACE__ . '\\majData_majTime_majGen_genmajTime';

