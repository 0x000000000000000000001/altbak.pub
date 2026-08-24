<?php

namespace Data\Time\Component\Gen;

// ALL IMPORTS: Control.Monad.Gen, Data.Enum.Gen, Data.Time.Component, Prim
// TO REQUIRE: Control.Monad.Gen, Data.Enum.Gen, Data.Time.Component
require_once __DIR__ . '/../Control.Monad.Gen/index.php';
require_once __DIR__ . '/../Data.Enum.Gen/index.php';
require_once __DIR__ . '/../Data.Time.Component/index.php';

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




// Data_Time_Component_Gen_genSecond
function majData_majTime_majComponent_majGen_genmajSecond($dictMonadGen_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTime_majComponent_majGen_genmajSecond';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Applicative0_1_0 = (((($dictMonadGen_0)->{'Monad0'})(null))->{'Applicative0'})(null);
  $v_2_1 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_2) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((($n_2 >= 0) && ($n_2 <= 59))) {
$__t1 = new \Data\Maybe\Data_Maybe_Just($n_2);
goto end_branch_1;;
};
  $__t1 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_2) {
  $__num = \func_num_args();
  $__res = ($v_2 + 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_2) {
  $__num = \func_num_args();
  $__res = $v_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), 0);
  $__t3 = null;;
  if ($v_2_1 instanceof \Data\Maybe\Data_Maybe_Just) {
$Monad0_3_4 = (($dictMonadGen_0)->{'Monad0'})(null);
$pure_4_5 = ((($Monad0_3_4)->{'Applicative0'})(null))->{'pure'};
$foldableNonEmpty1_5_6 = (object)["foldMap" => function($dictMonoid_5) {
  $__num = \func_num_args();
  $Semigroup0_6_6 = (($dictMonoid_5)->{'Semigroup0'})(null);
  $__res = function($f_7) use ($Semigroup0_6_6, $dictMonoid_5) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($Semigroup0_6_6, $dictMonoid_5, $f_7) {
  $__num = \func_num_args();
  $Semigroup0_9_7 = (($dictMonoid_5)->{'Semigroup0'})(null);
  $__res = ((($Semigroup0_6_6)->{'append'})(($f_7)(($v_8)->{'value0'})))(\Data\Foldable\majData_majFoldable_foldrmajArray(function($x_10) use ($Semigroup0_9_7, $f_7) {
  $__num = \func_num_args();
  $__res = function($acc_11) use ($Semigroup0_9_7, $f_7, $x_10) {
  $__num = \func_num_args();
  $__res = ((($Semigroup0_9_7)->{'append'})(($f_7)($x_10)))($acc_11);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($dictMonoid_5)->{'mempty'}, ($v_8)->{'value1'}));
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
}, "foldl" => function($f_5) {
  $__num = \func_num_args();
  $__res = function($b_6) use ($f_5) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($b_6, $f_5) {
  $__num = \func_num_args();
  $__res = \Data\Foldable\majData_majFoldable_foldlmajArray($f_5, (($f_5)($b_6))(($v_7)->{'value0'}), ($v_7)->{'value1'});
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
}, "foldr" => function($f_5) {
  $__num = \func_num_args();
  $__res = function($b_6) use ($f_5) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($b_6, $f_5) {
  $__num = \func_num_args();
  $__res = (($f_5)(($v_7)->{'value0'}))(\Data\Foldable\majData_majFoldable_foldrmajArray($f_5, $b_6, ($v_7)->{'value1'}));
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
$__local_var_5_6 = (object)["foldMap1" => function($dictSemigroup_6) {
  $__num = \func_num_args();
  $__res = function($f_7) use ($dictSemigroup_6) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($dictSemigroup_6, $f_7) {
  $__num = \func_num_args();
  $__res = \Data\Foldable\majData_majFoldable_foldlmajArray(function($s_9) use ($dictSemigroup_6, $f_7) {
  $__num = \func_num_args();
  $__res = function($a1_10) use ($dictSemigroup_6, $f_7, $s_9) {
  $__num = \func_num_args();
  $__res = ((($dictSemigroup_6)->{'append'})($s_9))(($f_7)($a1_10));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($f_7)(($v_8)->{'value0'}), ($v_8)->{'value1'});
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
}, "foldr1" => function($f_6) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($f_6) {
  $__num = \func_num_args();
  $__local_var_8_9 = ($f_6)(($v_7)->{'value0'});
  $__local_var_9_10 = \Data\Foldable\majData_majFoldable_foldrmajArray(function($a1_9) use ($f_6) {
  $__num = \func_num_args();
  $__local_var_10_10 = ($f_6)($a1_9);
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Maybe_Just']))(function($v2_11) use ($__local_var_10_10, $a1_9) {
  $__num = \func_num_args();
  $__t11 = null;;
  if ($v2_11 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t11 = $a1_9;
goto end_branch_11;;
};
  if ($v2_11 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t11 = ($__local_var_10_10)(($v2_11)->{'value0'});
goto end_branch_11;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t11 = null;
  end_branch_11:;
  $__res = $__t11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, new \Data\Maybe\Data_Maybe_Nothing(), ($v_7)->{'value1'});
  $__t13 = null;;
  if ($__local_var_9_10 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t13 = ($v_7)->{'value0'};
goto end_branch_13;;
};
  if ($__local_var_9_10 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t13 = ($__local_var_8_9)(($__local_var_9_10)->{'value0'});
goto end_branch_13;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t13 = null;
  end_branch_13:;
  $__res = $__t13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldl1" => function($f_6) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($f_6) {
  $__num = \func_num_args();
  $__res = \Data\Foldable\majData_majFoldable_foldlmajArray($f_6, ($v_7)->{'value0'}, ($v_7)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable0" => function($_dollar___unused_6) use ($foldableNonEmpty1_5_6) {
  $__num = \func_num_args();
  $__res = $foldableNonEmpty1_5_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
$__t20 = null;;
switch (($v_2_1)->{'value0'}) {
case 59:
$__t20 = \Data\Unfoldable1\majData_majUnfoldable1_unfoldr1majArraymajImpl($GLOBALS['Data_Maybe_isNothing'], function($v_6) {
  $__num = \func_num_args();
  $__t21 = null;;
  if ($v_6 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t21 = ($v_6)->{'value0'};
goto end_branch_21;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t21 = null;
  end_branch_21:;
  $__res = $__t21;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $GLOBALS['Data_Tuple_fst'], $GLOBALS['Data_Tuple_snd'], function($i_6) use ($v_2_1) {
  $__num = \func_num_args();
  $__t22 = null;;
  if (($i_6 <= 0)) {
$__t22 = new \Data\Tuple\Data_Tuple_Tuple(($v_2_1)->{'value0'}, new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_22;;
};
  $__t22 = new \Data\Tuple\Data_Tuple_Tuple(($v_2_1)->{'value0'}, new \Data\Maybe\Data_Maybe_Just(($i_6 - 1)));
  end_branch_22:;
  $__res = $__t22;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, 0);
goto end_branch_20;;
break;
default:
;
break;
};
if ((($v_2_1)->{'value0'} < 59)) {
$__t20 = \Data\Unfoldable1\majData_majUnfoldable1_unfoldr1majArraymajImpl($GLOBALS['Data_Maybe_isNothing'], function($v_6) {
  $__num = \func_num_args();
  $__t23 = null;;
  if ($v_6 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t23 = ($v_6)->{'value0'};
goto end_branch_23;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t23 = null;
  end_branch_23:;
  $__res = $__t23;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $GLOBALS['Data_Tuple_fst'], $GLOBALS['Data_Tuple_snd'], function($a_6) {
  $__num = \func_num_args();
  $__local_var_7_24 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_7) {
  $__num = \func_num_args();
  $__t24 = null;;
  if ((($n_7 >= 0) && ($n_7 <= 59))) {
$__t24 = new \Data\Maybe\Data_Maybe_Just($n_7);
goto end_branch_24;;
};
  $__t24 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_24:;
  $__res = $__t24;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_7) {
  $__num = \func_num_args();
  $__res = ($v_7 + 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_7) {
  $__num = \func_num_args();
  $__res = $v_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), $a_6);
  $__t26 = null;;
  if ($__local_var_7_24 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t27 = null;;
if ((($__local_var_7_24)->{'value0'} <= 59)) {
$__t27 = new \Data\Maybe\Data_Maybe_Just(($__local_var_7_24)->{'value0'});
goto end_branch_27;;
};
$__t27 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_27:;
$__t26 = $__t27;
goto end_branch_26;;
};
  if ($__local_var_7_24 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t26 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_26;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t26 = null;
  end_branch_26:;
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_6, $__t26);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($v_2_1)->{'value0'});
goto end_branch_20;;
};
$__t20 = \Data\Unfoldable1\majData_majUnfoldable1_unfoldr1majArraymajImpl($GLOBALS['Data_Maybe_isNothing'], function($v_6) {
  $__num = \func_num_args();
  $__t15 = null;;
  if ($v_6 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t15 = ($v_6)->{'value0'};
goto end_branch_15;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t15 = null;
  end_branch_15:;
  $__res = $__t15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $GLOBALS['Data_Tuple_fst'], $GLOBALS['Data_Tuple_snd'], function($a_6) {
  $__num = \func_num_args();
  $__local_var_7_16 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_7) {
  $__num = \func_num_args();
  $__t16 = null;;
  if ((($n_7 >= 0) && ($n_7 <= 59))) {
$__t16 = new \Data\Maybe\Data_Maybe_Just($n_7);
goto end_branch_16;;
};
  $__t16 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_16:;
  $__res = $__t16;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_7) {
  $__num = \func_num_args();
  $__res = ($v_7 - 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_7) {
  $__num = \func_num_args();
  $__res = $v_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), $a_6);
  $__t18 = null;;
  if ($__local_var_7_16 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t19 = null;;
if ((($__local_var_7_16)->{'value0'} >= 59)) {
$__t19 = new \Data\Maybe\Data_Maybe_Just(($__local_var_7_16)->{'value0'});
goto end_branch_19;;
};
$__t19 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_19:;
$__t18 = $__t19;
goto end_branch_18;;
};
  if ($__local_var_7_16 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t18 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_18;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t18 = null;
  end_branch_18:;
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_6, $__t18);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($v_2_1)->{'value0'});
end_branch_20:;
$__local_var_6_15 = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(0, $__t20);
$__t3 = ((((($Monad0_3_4)->{'Bind1'})(null))->{'bind'})(((($dictMonadGen_0)->{'chooseInt'})(0))(((((((($__local_var_5_6)->{'Foldable0'})(null))->{'foldl'})(function($c_7) {
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
}))(0))($__local_var_6_15) - 1))))(function($n_7) use ($__local_var_5_6, $__local_var_6_15, $pure_4_5) {
  $__num = \func_num_args();
  $go__go_8_29 = null;
  $go__go_8_29 = (function() use ($__local_var_5_6, $__local_var_6_15, &$go__go_8_29) {
  $__fn = function(int $v_9, $v1_10 = null) use ($__local_var_5_6, $__local_var_6_15, &$go__go_8_29, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_8_29_29_v_9 = $v_9;
  $__tco_var_go__go_8_29_29_v1_10 = $v1_10;
  tco_loop_go__go_8_29_29:;
  $v_9 = $__tco_var_go__go_8_29_29_v_9;
  $v1_10 = $__tco_var_go__go_8_29_29_v1_10;
  $__t29 = null;;
  if ($v1_10 instanceof \Control\Monad\Gen\Control_Monad_Gen_Cons) {
$__t32 = null;;
if (($v1_10)->{'value1'} instanceof \Control\Monad\Gen\Control_Monad_Gen_Nil) {
$__t32 = ($v1_10)->{'value0'};
goto end_branch_32;;
};
if (($v_9 <= 0)) {
$__t32 = ($v1_10)->{'value0'};
goto end_branch_32;;
};
$__tco_30 = ($v_9 - 1);
$__tco_31 = ($v1_10)->{'value1'};
$__tco_var_go__go_8_29_29_v_9 = $__tco_30;
$__tco_var_go__go_8_29_29_v1_10 = $__tco_31;
goto tco_loop_go__go_8_29_29;;
$__t32 = null;
end_branch_32:;
$__t29 = $__t32;
goto end_branch_29;;
};
  if ($v1_10 instanceof \Control\Monad\Gen\Control_Monad_Gen_Nil) {
$__t29 = (((($__local_var_5_6)->{'foldMap1'})($GLOBALS['Data_Semigroup_Last_semigroupLast']))(function($x_11) {
  $__num = \func_num_args();
  $__res = $x_11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($__local_var_6_15);
goto end_branch_29;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t29 = null;
  end_branch_29:;
  $__res = $__t29;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ($pure_4_5)((($go__go_8_29)($n_7))((((((($__local_var_5_6)->{'Foldable0'})(null))->{'foldr'})($GLOBALS['Control_Monad_Gen_Cons']))(new \Control\Monad\Gen\Control_Monad_Gen_Nil()))($__local_var_6_15)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
goto end_branch_3;;
};
  if ($v_2_1 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t3 = (($Applicative0_1_0)->{'pure'})(0);
goto end_branch_3;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t3 = null;
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Time_Component_Gen_genSecond'] = __NAMESPACE__ . '\\majData_majTime_majComponent_majGen_genmajSecond';

// Data_Time_Component_Gen_genMinute
function majData_majTime_majComponent_majGen_genmajMinute($dictMonadGen_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTime_majComponent_majGen_genmajMinute';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Applicative0_1_0 = (((($dictMonadGen_0)->{'Monad0'})(null))->{'Applicative0'})(null);
  $v_2_1 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_2) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((($n_2 >= 0) && ($n_2 <= 59))) {
$__t1 = new \Data\Maybe\Data_Maybe_Just($n_2);
goto end_branch_1;;
};
  $__t1 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_2) {
  $__num = \func_num_args();
  $__res = ($v_2 + 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_2) {
  $__num = \func_num_args();
  $__res = $v_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), 0);
  $__t3 = null;;
  if ($v_2_1 instanceof \Data\Maybe\Data_Maybe_Just) {
$Monad0_3_4 = (($dictMonadGen_0)->{'Monad0'})(null);
$pure_4_5 = ((($Monad0_3_4)->{'Applicative0'})(null))->{'pure'};
$foldableNonEmpty1_5_6 = (object)["foldMap" => function($dictMonoid_5) {
  $__num = \func_num_args();
  $Semigroup0_6_6 = (($dictMonoid_5)->{'Semigroup0'})(null);
  $__res = function($f_7) use ($Semigroup0_6_6, $dictMonoid_5) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($Semigroup0_6_6, $dictMonoid_5, $f_7) {
  $__num = \func_num_args();
  $Semigroup0_9_7 = (($dictMonoid_5)->{'Semigroup0'})(null);
  $__res = ((($Semigroup0_6_6)->{'append'})(($f_7)(($v_8)->{'value0'})))(\Data\Foldable\majData_majFoldable_foldrmajArray(function($x_10) use ($Semigroup0_9_7, $f_7) {
  $__num = \func_num_args();
  $__res = function($acc_11) use ($Semigroup0_9_7, $f_7, $x_10) {
  $__num = \func_num_args();
  $__res = ((($Semigroup0_9_7)->{'append'})(($f_7)($x_10)))($acc_11);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($dictMonoid_5)->{'mempty'}, ($v_8)->{'value1'}));
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
}, "foldl" => function($f_5) {
  $__num = \func_num_args();
  $__res = function($b_6) use ($f_5) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($b_6, $f_5) {
  $__num = \func_num_args();
  $__res = \Data\Foldable\majData_majFoldable_foldlmajArray($f_5, (($f_5)($b_6))(($v_7)->{'value0'}), ($v_7)->{'value1'});
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
}, "foldr" => function($f_5) {
  $__num = \func_num_args();
  $__res = function($b_6) use ($f_5) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($b_6, $f_5) {
  $__num = \func_num_args();
  $__res = (($f_5)(($v_7)->{'value0'}))(\Data\Foldable\majData_majFoldable_foldrmajArray($f_5, $b_6, ($v_7)->{'value1'}));
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
$__local_var_5_6 = (object)["foldMap1" => function($dictSemigroup_6) {
  $__num = \func_num_args();
  $__res = function($f_7) use ($dictSemigroup_6) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($dictSemigroup_6, $f_7) {
  $__num = \func_num_args();
  $__res = \Data\Foldable\majData_majFoldable_foldlmajArray(function($s_9) use ($dictSemigroup_6, $f_7) {
  $__num = \func_num_args();
  $__res = function($a1_10) use ($dictSemigroup_6, $f_7, $s_9) {
  $__num = \func_num_args();
  $__res = ((($dictSemigroup_6)->{'append'})($s_9))(($f_7)($a1_10));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($f_7)(($v_8)->{'value0'}), ($v_8)->{'value1'});
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
}, "foldr1" => function($f_6) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($f_6) {
  $__num = \func_num_args();
  $__local_var_8_9 = ($f_6)(($v_7)->{'value0'});
  $__local_var_9_10 = \Data\Foldable\majData_majFoldable_foldrmajArray(function($a1_9) use ($f_6) {
  $__num = \func_num_args();
  $__local_var_10_10 = ($f_6)($a1_9);
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Maybe_Just']))(function($v2_11) use ($__local_var_10_10, $a1_9) {
  $__num = \func_num_args();
  $__t11 = null;;
  if ($v2_11 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t11 = $a1_9;
goto end_branch_11;;
};
  if ($v2_11 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t11 = ($__local_var_10_10)(($v2_11)->{'value0'});
goto end_branch_11;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t11 = null;
  end_branch_11:;
  $__res = $__t11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, new \Data\Maybe\Data_Maybe_Nothing(), ($v_7)->{'value1'});
  $__t13 = null;;
  if ($__local_var_9_10 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t13 = ($v_7)->{'value0'};
goto end_branch_13;;
};
  if ($__local_var_9_10 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t13 = ($__local_var_8_9)(($__local_var_9_10)->{'value0'});
goto end_branch_13;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t13 = null;
  end_branch_13:;
  $__res = $__t13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldl1" => function($f_6) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($f_6) {
  $__num = \func_num_args();
  $__res = \Data\Foldable\majData_majFoldable_foldlmajArray($f_6, ($v_7)->{'value0'}, ($v_7)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable0" => function($_dollar___unused_6) use ($foldableNonEmpty1_5_6) {
  $__num = \func_num_args();
  $__res = $foldableNonEmpty1_5_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
$__t20 = null;;
switch (($v_2_1)->{'value0'}) {
case 59:
$__t20 = \Data\Unfoldable1\majData_majUnfoldable1_unfoldr1majArraymajImpl($GLOBALS['Data_Maybe_isNothing'], function($v_6) {
  $__num = \func_num_args();
  $__t21 = null;;
  if ($v_6 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t21 = ($v_6)->{'value0'};
goto end_branch_21;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t21 = null;
  end_branch_21:;
  $__res = $__t21;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $GLOBALS['Data_Tuple_fst'], $GLOBALS['Data_Tuple_snd'], function($i_6) use ($v_2_1) {
  $__num = \func_num_args();
  $__t22 = null;;
  if (($i_6 <= 0)) {
$__t22 = new \Data\Tuple\Data_Tuple_Tuple(($v_2_1)->{'value0'}, new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_22;;
};
  $__t22 = new \Data\Tuple\Data_Tuple_Tuple(($v_2_1)->{'value0'}, new \Data\Maybe\Data_Maybe_Just(($i_6 - 1)));
  end_branch_22:;
  $__res = $__t22;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, 0);
goto end_branch_20;;
break;
default:
;
break;
};
if ((($v_2_1)->{'value0'} < 59)) {
$__t20 = \Data\Unfoldable1\majData_majUnfoldable1_unfoldr1majArraymajImpl($GLOBALS['Data_Maybe_isNothing'], function($v_6) {
  $__num = \func_num_args();
  $__t23 = null;;
  if ($v_6 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t23 = ($v_6)->{'value0'};
goto end_branch_23;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t23 = null;
  end_branch_23:;
  $__res = $__t23;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $GLOBALS['Data_Tuple_fst'], $GLOBALS['Data_Tuple_snd'], function($a_6) {
  $__num = \func_num_args();
  $__local_var_7_24 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_7) {
  $__num = \func_num_args();
  $__t24 = null;;
  if ((($n_7 >= 0) && ($n_7 <= 59))) {
$__t24 = new \Data\Maybe\Data_Maybe_Just($n_7);
goto end_branch_24;;
};
  $__t24 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_24:;
  $__res = $__t24;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_7) {
  $__num = \func_num_args();
  $__res = ($v_7 + 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_7) {
  $__num = \func_num_args();
  $__res = $v_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), $a_6);
  $__t26 = null;;
  if ($__local_var_7_24 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t27 = null;;
if ((($__local_var_7_24)->{'value0'} <= 59)) {
$__t27 = new \Data\Maybe\Data_Maybe_Just(($__local_var_7_24)->{'value0'});
goto end_branch_27;;
};
$__t27 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_27:;
$__t26 = $__t27;
goto end_branch_26;;
};
  if ($__local_var_7_24 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t26 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_26;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t26 = null;
  end_branch_26:;
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_6, $__t26);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($v_2_1)->{'value0'});
goto end_branch_20;;
};
$__t20 = \Data\Unfoldable1\majData_majUnfoldable1_unfoldr1majArraymajImpl($GLOBALS['Data_Maybe_isNothing'], function($v_6) {
  $__num = \func_num_args();
  $__t15 = null;;
  if ($v_6 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t15 = ($v_6)->{'value0'};
goto end_branch_15;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t15 = null;
  end_branch_15:;
  $__res = $__t15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $GLOBALS['Data_Tuple_fst'], $GLOBALS['Data_Tuple_snd'], function($a_6) {
  $__num = \func_num_args();
  $__local_var_7_16 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_7) {
  $__num = \func_num_args();
  $__t16 = null;;
  if ((($n_7 >= 0) && ($n_7 <= 59))) {
$__t16 = new \Data\Maybe\Data_Maybe_Just($n_7);
goto end_branch_16;;
};
  $__t16 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_16:;
  $__res = $__t16;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_7) {
  $__num = \func_num_args();
  $__res = ($v_7 - 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_7) {
  $__num = \func_num_args();
  $__res = $v_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), $a_6);
  $__t18 = null;;
  if ($__local_var_7_16 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t19 = null;;
if ((($__local_var_7_16)->{'value0'} >= 59)) {
$__t19 = new \Data\Maybe\Data_Maybe_Just(($__local_var_7_16)->{'value0'});
goto end_branch_19;;
};
$__t19 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_19:;
$__t18 = $__t19;
goto end_branch_18;;
};
  if ($__local_var_7_16 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t18 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_18;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t18 = null;
  end_branch_18:;
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_6, $__t18);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($v_2_1)->{'value0'});
end_branch_20:;
$__local_var_6_15 = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(0, $__t20);
$__t3 = ((((($Monad0_3_4)->{'Bind1'})(null))->{'bind'})(((($dictMonadGen_0)->{'chooseInt'})(0))(((((((($__local_var_5_6)->{'Foldable0'})(null))->{'foldl'})(function($c_7) {
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
}))(0))($__local_var_6_15) - 1))))(function($n_7) use ($__local_var_5_6, $__local_var_6_15, $pure_4_5) {
  $__num = \func_num_args();
  $go__go_8_29 = null;
  $go__go_8_29 = (function() use ($__local_var_5_6, $__local_var_6_15, &$go__go_8_29) {
  $__fn = function(int $v_9, $v1_10 = null) use ($__local_var_5_6, $__local_var_6_15, &$go__go_8_29, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_8_29_29_v_9 = $v_9;
  $__tco_var_go__go_8_29_29_v1_10 = $v1_10;
  tco_loop_go__go_8_29_29:;
  $v_9 = $__tco_var_go__go_8_29_29_v_9;
  $v1_10 = $__tco_var_go__go_8_29_29_v1_10;
  $__t29 = null;;
  if ($v1_10 instanceof \Control\Monad\Gen\Control_Monad_Gen_Cons) {
$__t32 = null;;
if (($v1_10)->{'value1'} instanceof \Control\Monad\Gen\Control_Monad_Gen_Nil) {
$__t32 = ($v1_10)->{'value0'};
goto end_branch_32;;
};
if (($v_9 <= 0)) {
$__t32 = ($v1_10)->{'value0'};
goto end_branch_32;;
};
$__tco_30 = ($v_9 - 1);
$__tco_31 = ($v1_10)->{'value1'};
$__tco_var_go__go_8_29_29_v_9 = $__tco_30;
$__tco_var_go__go_8_29_29_v1_10 = $__tco_31;
goto tco_loop_go__go_8_29_29;;
$__t32 = null;
end_branch_32:;
$__t29 = $__t32;
goto end_branch_29;;
};
  if ($v1_10 instanceof \Control\Monad\Gen\Control_Monad_Gen_Nil) {
$__t29 = (((($__local_var_5_6)->{'foldMap1'})($GLOBALS['Data_Semigroup_Last_semigroupLast']))(function($x_11) {
  $__num = \func_num_args();
  $__res = $x_11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($__local_var_6_15);
goto end_branch_29;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t29 = null;
  end_branch_29:;
  $__res = $__t29;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ($pure_4_5)((($go__go_8_29)($n_7))((((((($__local_var_5_6)->{'Foldable0'})(null))->{'foldr'})($GLOBALS['Control_Monad_Gen_Cons']))(new \Control\Monad\Gen\Control_Monad_Gen_Nil()))($__local_var_6_15)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
goto end_branch_3;;
};
  if ($v_2_1 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t3 = (($Applicative0_1_0)->{'pure'})(0);
goto end_branch_3;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t3 = null;
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Time_Component_Gen_genMinute'] = __NAMESPACE__ . '\\majData_majTime_majComponent_majGen_genmajMinute';

// Data_Time_Component_Gen_genMillisecond
function majData_majTime_majComponent_majGen_genmajMillisecond($dictMonadGen_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTime_majComponent_majGen_genmajMillisecond';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Applicative0_1_0 = (((($dictMonadGen_0)->{'Monad0'})(null))->{'Applicative0'})(null);
  $v_2_1 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_2) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((($n_2 >= 0) && ($n_2 <= 999))) {
$__t1 = new \Data\Maybe\Data_Maybe_Just($n_2);
goto end_branch_1;;
};
  $__t1 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_2) {
  $__num = \func_num_args();
  $__res = ($v_2 + 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_2) {
  $__num = \func_num_args();
  $__res = $v_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), 0);
  $__t3 = null;;
  if ($v_2_1 instanceof \Data\Maybe\Data_Maybe_Just) {
$Monad0_3_4 = (($dictMonadGen_0)->{'Monad0'})(null);
$pure_4_5 = ((($Monad0_3_4)->{'Applicative0'})(null))->{'pure'};
$foldableNonEmpty1_5_6 = (object)["foldMap" => function($dictMonoid_5) {
  $__num = \func_num_args();
  $Semigroup0_6_6 = (($dictMonoid_5)->{'Semigroup0'})(null);
  $__res = function($f_7) use ($Semigroup0_6_6, $dictMonoid_5) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($Semigroup0_6_6, $dictMonoid_5, $f_7) {
  $__num = \func_num_args();
  $Semigroup0_9_7 = (($dictMonoid_5)->{'Semigroup0'})(null);
  $__res = ((($Semigroup0_6_6)->{'append'})(($f_7)(($v_8)->{'value0'})))(\Data\Foldable\majData_majFoldable_foldrmajArray(function($x_10) use ($Semigroup0_9_7, $f_7) {
  $__num = \func_num_args();
  $__res = function($acc_11) use ($Semigroup0_9_7, $f_7, $x_10) {
  $__num = \func_num_args();
  $__res = ((($Semigroup0_9_7)->{'append'})(($f_7)($x_10)))($acc_11);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($dictMonoid_5)->{'mempty'}, ($v_8)->{'value1'}));
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
}, "foldl" => function($f_5) {
  $__num = \func_num_args();
  $__res = function($b_6) use ($f_5) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($b_6, $f_5) {
  $__num = \func_num_args();
  $__res = \Data\Foldable\majData_majFoldable_foldlmajArray($f_5, (($f_5)($b_6))(($v_7)->{'value0'}), ($v_7)->{'value1'});
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
}, "foldr" => function($f_5) {
  $__num = \func_num_args();
  $__res = function($b_6) use ($f_5) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($b_6, $f_5) {
  $__num = \func_num_args();
  $__res = (($f_5)(($v_7)->{'value0'}))(\Data\Foldable\majData_majFoldable_foldrmajArray($f_5, $b_6, ($v_7)->{'value1'}));
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
$__local_var_5_6 = (object)["foldMap1" => function($dictSemigroup_6) {
  $__num = \func_num_args();
  $__res = function($f_7) use ($dictSemigroup_6) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($dictSemigroup_6, $f_7) {
  $__num = \func_num_args();
  $__res = \Data\Foldable\majData_majFoldable_foldlmajArray(function($s_9) use ($dictSemigroup_6, $f_7) {
  $__num = \func_num_args();
  $__res = function($a1_10) use ($dictSemigroup_6, $f_7, $s_9) {
  $__num = \func_num_args();
  $__res = ((($dictSemigroup_6)->{'append'})($s_9))(($f_7)($a1_10));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($f_7)(($v_8)->{'value0'}), ($v_8)->{'value1'});
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
}, "foldr1" => function($f_6) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($f_6) {
  $__num = \func_num_args();
  $__local_var_8_9 = ($f_6)(($v_7)->{'value0'});
  $__local_var_9_10 = \Data\Foldable\majData_majFoldable_foldrmajArray(function($a1_9) use ($f_6) {
  $__num = \func_num_args();
  $__local_var_10_10 = ($f_6)($a1_9);
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Maybe_Just']))(function($v2_11) use ($__local_var_10_10, $a1_9) {
  $__num = \func_num_args();
  $__t11 = null;;
  if ($v2_11 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t11 = $a1_9;
goto end_branch_11;;
};
  if ($v2_11 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t11 = ($__local_var_10_10)(($v2_11)->{'value0'});
goto end_branch_11;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t11 = null;
  end_branch_11:;
  $__res = $__t11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, new \Data\Maybe\Data_Maybe_Nothing(), ($v_7)->{'value1'});
  $__t13 = null;;
  if ($__local_var_9_10 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t13 = ($v_7)->{'value0'};
goto end_branch_13;;
};
  if ($__local_var_9_10 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t13 = ($__local_var_8_9)(($__local_var_9_10)->{'value0'});
goto end_branch_13;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t13 = null;
  end_branch_13:;
  $__res = $__t13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldl1" => function($f_6) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($f_6) {
  $__num = \func_num_args();
  $__res = \Data\Foldable\majData_majFoldable_foldlmajArray($f_6, ($v_7)->{'value0'}, ($v_7)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable0" => function($_dollar___unused_6) use ($foldableNonEmpty1_5_6) {
  $__num = \func_num_args();
  $__res = $foldableNonEmpty1_5_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
$__t20 = null;;
switch (($v_2_1)->{'value0'}) {
case 999:
$__t20 = \Data\Unfoldable1\majData_majUnfoldable1_unfoldr1majArraymajImpl($GLOBALS['Data_Maybe_isNothing'], function($v_6) {
  $__num = \func_num_args();
  $__t21 = null;;
  if ($v_6 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t21 = ($v_6)->{'value0'};
goto end_branch_21;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t21 = null;
  end_branch_21:;
  $__res = $__t21;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $GLOBALS['Data_Tuple_fst'], $GLOBALS['Data_Tuple_snd'], function($i_6) use ($v_2_1) {
  $__num = \func_num_args();
  $__t22 = null;;
  if (($i_6 <= 0)) {
$__t22 = new \Data\Tuple\Data_Tuple_Tuple(($v_2_1)->{'value0'}, new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_22;;
};
  $__t22 = new \Data\Tuple\Data_Tuple_Tuple(($v_2_1)->{'value0'}, new \Data\Maybe\Data_Maybe_Just(($i_6 - 1)));
  end_branch_22:;
  $__res = $__t22;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, 0);
goto end_branch_20;;
break;
default:
;
break;
};
if ((($v_2_1)->{'value0'} < 999)) {
$__t20 = \Data\Unfoldable1\majData_majUnfoldable1_unfoldr1majArraymajImpl($GLOBALS['Data_Maybe_isNothing'], function($v_6) {
  $__num = \func_num_args();
  $__t23 = null;;
  if ($v_6 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t23 = ($v_6)->{'value0'};
goto end_branch_23;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t23 = null;
  end_branch_23:;
  $__res = $__t23;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $GLOBALS['Data_Tuple_fst'], $GLOBALS['Data_Tuple_snd'], function($a_6) {
  $__num = \func_num_args();
  $__local_var_7_24 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_7) {
  $__num = \func_num_args();
  $__t24 = null;;
  if ((($n_7 >= 0) && ($n_7 <= 999))) {
$__t24 = new \Data\Maybe\Data_Maybe_Just($n_7);
goto end_branch_24;;
};
  $__t24 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_24:;
  $__res = $__t24;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_7) {
  $__num = \func_num_args();
  $__res = ($v_7 + 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_7) {
  $__num = \func_num_args();
  $__res = $v_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), $a_6);
  $__t26 = null;;
  if ($__local_var_7_24 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t27 = null;;
if ((($__local_var_7_24)->{'value0'} <= 999)) {
$__t27 = new \Data\Maybe\Data_Maybe_Just(($__local_var_7_24)->{'value0'});
goto end_branch_27;;
};
$__t27 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_27:;
$__t26 = $__t27;
goto end_branch_26;;
};
  if ($__local_var_7_24 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t26 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_26;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t26 = null;
  end_branch_26:;
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_6, $__t26);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($v_2_1)->{'value0'});
goto end_branch_20;;
};
$__t20 = \Data\Unfoldable1\majData_majUnfoldable1_unfoldr1majArraymajImpl($GLOBALS['Data_Maybe_isNothing'], function($v_6) {
  $__num = \func_num_args();
  $__t15 = null;;
  if ($v_6 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t15 = ($v_6)->{'value0'};
goto end_branch_15;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t15 = null;
  end_branch_15:;
  $__res = $__t15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $GLOBALS['Data_Tuple_fst'], $GLOBALS['Data_Tuple_snd'], function($a_6) {
  $__num = \func_num_args();
  $__local_var_7_16 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_7) {
  $__num = \func_num_args();
  $__t16 = null;;
  if ((($n_7 >= 0) && ($n_7 <= 999))) {
$__t16 = new \Data\Maybe\Data_Maybe_Just($n_7);
goto end_branch_16;;
};
  $__t16 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_16:;
  $__res = $__t16;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_7) {
  $__num = \func_num_args();
  $__res = ($v_7 - 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_7) {
  $__num = \func_num_args();
  $__res = $v_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), $a_6);
  $__t18 = null;;
  if ($__local_var_7_16 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t19 = null;;
if ((($__local_var_7_16)->{'value0'} >= 999)) {
$__t19 = new \Data\Maybe\Data_Maybe_Just(($__local_var_7_16)->{'value0'});
goto end_branch_19;;
};
$__t19 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_19:;
$__t18 = $__t19;
goto end_branch_18;;
};
  if ($__local_var_7_16 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t18 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_18;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t18 = null;
  end_branch_18:;
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_6, $__t18);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($v_2_1)->{'value0'});
end_branch_20:;
$__local_var_6_15 = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(0, $__t20);
$__t3 = ((((($Monad0_3_4)->{'Bind1'})(null))->{'bind'})(((($dictMonadGen_0)->{'chooseInt'})(0))(((((((($__local_var_5_6)->{'Foldable0'})(null))->{'foldl'})(function($c_7) {
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
}))(0))($__local_var_6_15) - 1))))(function($n_7) use ($__local_var_5_6, $__local_var_6_15, $pure_4_5) {
  $__num = \func_num_args();
  $go__go_8_29 = null;
  $go__go_8_29 = (function() use ($__local_var_5_6, $__local_var_6_15, &$go__go_8_29) {
  $__fn = function(int $v_9, $v1_10 = null) use ($__local_var_5_6, $__local_var_6_15, &$go__go_8_29, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_8_29_29_v_9 = $v_9;
  $__tco_var_go__go_8_29_29_v1_10 = $v1_10;
  tco_loop_go__go_8_29_29:;
  $v_9 = $__tco_var_go__go_8_29_29_v_9;
  $v1_10 = $__tco_var_go__go_8_29_29_v1_10;
  $__t29 = null;;
  if ($v1_10 instanceof \Control\Monad\Gen\Control_Monad_Gen_Cons) {
$__t32 = null;;
if (($v1_10)->{'value1'} instanceof \Control\Monad\Gen\Control_Monad_Gen_Nil) {
$__t32 = ($v1_10)->{'value0'};
goto end_branch_32;;
};
if (($v_9 <= 0)) {
$__t32 = ($v1_10)->{'value0'};
goto end_branch_32;;
};
$__tco_30 = ($v_9 - 1);
$__tco_31 = ($v1_10)->{'value1'};
$__tco_var_go__go_8_29_29_v_9 = $__tco_30;
$__tco_var_go__go_8_29_29_v1_10 = $__tco_31;
goto tco_loop_go__go_8_29_29;;
$__t32 = null;
end_branch_32:;
$__t29 = $__t32;
goto end_branch_29;;
};
  if ($v1_10 instanceof \Control\Monad\Gen\Control_Monad_Gen_Nil) {
$__t29 = (((($__local_var_5_6)->{'foldMap1'})($GLOBALS['Data_Semigroup_Last_semigroupLast']))(function($x_11) {
  $__num = \func_num_args();
  $__res = $x_11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($__local_var_6_15);
goto end_branch_29;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t29 = null;
  end_branch_29:;
  $__res = $__t29;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ($pure_4_5)((($go__go_8_29)($n_7))((((((($__local_var_5_6)->{'Foldable0'})(null))->{'foldr'})($GLOBALS['Control_Monad_Gen_Cons']))(new \Control\Monad\Gen\Control_Monad_Gen_Nil()))($__local_var_6_15)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
goto end_branch_3;;
};
  if ($v_2_1 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t3 = (($Applicative0_1_0)->{'pure'})(0);
goto end_branch_3;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t3 = null;
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Time_Component_Gen_genMillisecond'] = __NAMESPACE__ . '\\majData_majTime_majComponent_majGen_genmajMillisecond';

// Data_Time_Component_Gen_genHour
function majData_majTime_majComponent_majGen_genmajHour($dictMonadGen_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTime_majComponent_majGen_genmajHour';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Applicative0_1_0 = (((($dictMonadGen_0)->{'Monad0'})(null))->{'Applicative0'})(null);
  $v_2_1 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_2) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((($n_2 >= 0) && ($n_2 <= 23))) {
$__t1 = new \Data\Maybe\Data_Maybe_Just($n_2);
goto end_branch_1;;
};
  $__t1 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_2) {
  $__num = \func_num_args();
  $__res = ($v_2 + 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_2) {
  $__num = \func_num_args();
  $__res = $v_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), 0);
  $__t3 = null;;
  if ($v_2_1 instanceof \Data\Maybe\Data_Maybe_Just) {
$Monad0_3_4 = (($dictMonadGen_0)->{'Monad0'})(null);
$pure_4_5 = ((($Monad0_3_4)->{'Applicative0'})(null))->{'pure'};
$foldableNonEmpty1_5_6 = (object)["foldMap" => function($dictMonoid_5) {
  $__num = \func_num_args();
  $Semigroup0_6_6 = (($dictMonoid_5)->{'Semigroup0'})(null);
  $__res = function($f_7) use ($Semigroup0_6_6, $dictMonoid_5) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($Semigroup0_6_6, $dictMonoid_5, $f_7) {
  $__num = \func_num_args();
  $Semigroup0_9_7 = (($dictMonoid_5)->{'Semigroup0'})(null);
  $__res = ((($Semigroup0_6_6)->{'append'})(($f_7)(($v_8)->{'value0'})))(\Data\Foldable\majData_majFoldable_foldrmajArray(function($x_10) use ($Semigroup0_9_7, $f_7) {
  $__num = \func_num_args();
  $__res = function($acc_11) use ($Semigroup0_9_7, $f_7, $x_10) {
  $__num = \func_num_args();
  $__res = ((($Semigroup0_9_7)->{'append'})(($f_7)($x_10)))($acc_11);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($dictMonoid_5)->{'mempty'}, ($v_8)->{'value1'}));
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
}, "foldl" => function($f_5) {
  $__num = \func_num_args();
  $__res = function($b_6) use ($f_5) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($b_6, $f_5) {
  $__num = \func_num_args();
  $__res = \Data\Foldable\majData_majFoldable_foldlmajArray($f_5, (($f_5)($b_6))(($v_7)->{'value0'}), ($v_7)->{'value1'});
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
}, "foldr" => function($f_5) {
  $__num = \func_num_args();
  $__res = function($b_6) use ($f_5) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($b_6, $f_5) {
  $__num = \func_num_args();
  $__res = (($f_5)(($v_7)->{'value0'}))(\Data\Foldable\majData_majFoldable_foldrmajArray($f_5, $b_6, ($v_7)->{'value1'}));
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
$__local_var_5_6 = (object)["foldMap1" => function($dictSemigroup_6) {
  $__num = \func_num_args();
  $__res = function($f_7) use ($dictSemigroup_6) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($dictSemigroup_6, $f_7) {
  $__num = \func_num_args();
  $__res = \Data\Foldable\majData_majFoldable_foldlmajArray(function($s_9) use ($dictSemigroup_6, $f_7) {
  $__num = \func_num_args();
  $__res = function($a1_10) use ($dictSemigroup_6, $f_7, $s_9) {
  $__num = \func_num_args();
  $__res = ((($dictSemigroup_6)->{'append'})($s_9))(($f_7)($a1_10));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($f_7)(($v_8)->{'value0'}), ($v_8)->{'value1'});
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
}, "foldr1" => function($f_6) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($f_6) {
  $__num = \func_num_args();
  $__local_var_8_9 = ($f_6)(($v_7)->{'value0'});
  $__local_var_9_10 = \Data\Foldable\majData_majFoldable_foldrmajArray(function($a1_9) use ($f_6) {
  $__num = \func_num_args();
  $__local_var_10_10 = ($f_6)($a1_9);
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Maybe_Just']))(function($v2_11) use ($__local_var_10_10, $a1_9) {
  $__num = \func_num_args();
  $__t11 = null;;
  if ($v2_11 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t11 = $a1_9;
goto end_branch_11;;
};
  if ($v2_11 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t11 = ($__local_var_10_10)(($v2_11)->{'value0'});
goto end_branch_11;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t11 = null;
  end_branch_11:;
  $__res = $__t11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, new \Data\Maybe\Data_Maybe_Nothing(), ($v_7)->{'value1'});
  $__t13 = null;;
  if ($__local_var_9_10 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t13 = ($v_7)->{'value0'};
goto end_branch_13;;
};
  if ($__local_var_9_10 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t13 = ($__local_var_8_9)(($__local_var_9_10)->{'value0'});
goto end_branch_13;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t13 = null;
  end_branch_13:;
  $__res = $__t13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldl1" => function($f_6) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($f_6) {
  $__num = \func_num_args();
  $__res = \Data\Foldable\majData_majFoldable_foldlmajArray($f_6, ($v_7)->{'value0'}, ($v_7)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable0" => function($_dollar___unused_6) use ($foldableNonEmpty1_5_6) {
  $__num = \func_num_args();
  $__res = $foldableNonEmpty1_5_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
$__t20 = null;;
switch (($v_2_1)->{'value0'}) {
case 23:
$__t20 = \Data\Unfoldable1\majData_majUnfoldable1_unfoldr1majArraymajImpl($GLOBALS['Data_Maybe_isNothing'], function($v_6) {
  $__num = \func_num_args();
  $__t21 = null;;
  if ($v_6 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t21 = ($v_6)->{'value0'};
goto end_branch_21;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t21 = null;
  end_branch_21:;
  $__res = $__t21;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $GLOBALS['Data_Tuple_fst'], $GLOBALS['Data_Tuple_snd'], function($i_6) use ($v_2_1) {
  $__num = \func_num_args();
  $__t22 = null;;
  if (($i_6 <= 0)) {
$__t22 = new \Data\Tuple\Data_Tuple_Tuple(($v_2_1)->{'value0'}, new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_22;;
};
  $__t22 = new \Data\Tuple\Data_Tuple_Tuple(($v_2_1)->{'value0'}, new \Data\Maybe\Data_Maybe_Just(($i_6 - 1)));
  end_branch_22:;
  $__res = $__t22;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, 0);
goto end_branch_20;;
break;
default:
;
break;
};
if ((($v_2_1)->{'value0'} < 23)) {
$__t20 = \Data\Unfoldable1\majData_majUnfoldable1_unfoldr1majArraymajImpl($GLOBALS['Data_Maybe_isNothing'], function($v_6) {
  $__num = \func_num_args();
  $__t23 = null;;
  if ($v_6 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t23 = ($v_6)->{'value0'};
goto end_branch_23;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t23 = null;
  end_branch_23:;
  $__res = $__t23;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $GLOBALS['Data_Tuple_fst'], $GLOBALS['Data_Tuple_snd'], function($a_6) {
  $__num = \func_num_args();
  $__local_var_7_24 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_7) {
  $__num = \func_num_args();
  $__t24 = null;;
  if ((($n_7 >= 0) && ($n_7 <= 23))) {
$__t24 = new \Data\Maybe\Data_Maybe_Just($n_7);
goto end_branch_24;;
};
  $__t24 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_24:;
  $__res = $__t24;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_7) {
  $__num = \func_num_args();
  $__res = ($v_7 + 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_7) {
  $__num = \func_num_args();
  $__res = $v_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), $a_6);
  $__t26 = null;;
  if ($__local_var_7_24 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t27 = null;;
if ((($__local_var_7_24)->{'value0'} <= 23)) {
$__t27 = new \Data\Maybe\Data_Maybe_Just(($__local_var_7_24)->{'value0'});
goto end_branch_27;;
};
$__t27 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_27:;
$__t26 = $__t27;
goto end_branch_26;;
};
  if ($__local_var_7_24 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t26 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_26;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t26 = null;
  end_branch_26:;
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_6, $__t26);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($v_2_1)->{'value0'});
goto end_branch_20;;
};
$__t20 = \Data\Unfoldable1\majData_majUnfoldable1_unfoldr1majArraymajImpl($GLOBALS['Data_Maybe_isNothing'], function($v_6) {
  $__num = \func_num_args();
  $__t15 = null;;
  if ($v_6 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t15 = ($v_6)->{'value0'};
goto end_branch_15;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t15 = null;
  end_branch_15:;
  $__res = $__t15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $GLOBALS['Data_Tuple_fst'], $GLOBALS['Data_Tuple_snd'], function($a_6) {
  $__num = \func_num_args();
  $__local_var_7_16 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_7) {
  $__num = \func_num_args();
  $__t16 = null;;
  if ((($n_7 >= 0) && ($n_7 <= 23))) {
$__t16 = new \Data\Maybe\Data_Maybe_Just($n_7);
goto end_branch_16;;
};
  $__t16 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_16:;
  $__res = $__t16;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_7) {
  $__num = \func_num_args();
  $__res = ($v_7 - 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_7) {
  $__num = \func_num_args();
  $__res = $v_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), $a_6);
  $__t18 = null;;
  if ($__local_var_7_16 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t19 = null;;
if ((($__local_var_7_16)->{'value0'} >= 23)) {
$__t19 = new \Data\Maybe\Data_Maybe_Just(($__local_var_7_16)->{'value0'});
goto end_branch_19;;
};
$__t19 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_19:;
$__t18 = $__t19;
goto end_branch_18;;
};
  if ($__local_var_7_16 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t18 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_18;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t18 = null;
  end_branch_18:;
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_6, $__t18);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($v_2_1)->{'value0'});
end_branch_20:;
$__local_var_6_15 = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(0, $__t20);
$__t3 = ((((($Monad0_3_4)->{'Bind1'})(null))->{'bind'})(((($dictMonadGen_0)->{'chooseInt'})(0))(((((((($__local_var_5_6)->{'Foldable0'})(null))->{'foldl'})(function($c_7) {
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
}))(0))($__local_var_6_15) - 1))))(function($n_7) use ($__local_var_5_6, $__local_var_6_15, $pure_4_5) {
  $__num = \func_num_args();
  $go__go_8_29 = null;
  $go__go_8_29 = (function() use ($__local_var_5_6, $__local_var_6_15, &$go__go_8_29) {
  $__fn = function(int $v_9, $v1_10 = null) use ($__local_var_5_6, $__local_var_6_15, &$go__go_8_29, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_8_29_29_v_9 = $v_9;
  $__tco_var_go__go_8_29_29_v1_10 = $v1_10;
  tco_loop_go__go_8_29_29:;
  $v_9 = $__tco_var_go__go_8_29_29_v_9;
  $v1_10 = $__tco_var_go__go_8_29_29_v1_10;
  $__t29 = null;;
  if ($v1_10 instanceof \Control\Monad\Gen\Control_Monad_Gen_Cons) {
$__t32 = null;;
if (($v1_10)->{'value1'} instanceof \Control\Monad\Gen\Control_Monad_Gen_Nil) {
$__t32 = ($v1_10)->{'value0'};
goto end_branch_32;;
};
if (($v_9 <= 0)) {
$__t32 = ($v1_10)->{'value0'};
goto end_branch_32;;
};
$__tco_30 = ($v_9 - 1);
$__tco_31 = ($v1_10)->{'value1'};
$__tco_var_go__go_8_29_29_v_9 = $__tco_30;
$__tco_var_go__go_8_29_29_v1_10 = $__tco_31;
goto tco_loop_go__go_8_29_29;;
$__t32 = null;
end_branch_32:;
$__t29 = $__t32;
goto end_branch_29;;
};
  if ($v1_10 instanceof \Control\Monad\Gen\Control_Monad_Gen_Nil) {
$__t29 = (((($__local_var_5_6)->{'foldMap1'})($GLOBALS['Data_Semigroup_Last_semigroupLast']))(function($x_11) {
  $__num = \func_num_args();
  $__res = $x_11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($__local_var_6_15);
goto end_branch_29;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t29 = null;
  end_branch_29:;
  $__res = $__t29;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ($pure_4_5)((($go__go_8_29)($n_7))((((((($__local_var_5_6)->{'Foldable0'})(null))->{'foldr'})($GLOBALS['Control_Monad_Gen_Cons']))(new \Control\Monad\Gen\Control_Monad_Gen_Nil()))($__local_var_6_15)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
goto end_branch_3;;
};
  if ($v_2_1 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t3 = (($Applicative0_1_0)->{'pure'})(0);
goto end_branch_3;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t3 = null;
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Time_Component_Gen_genHour'] = __NAMESPACE__ . '\\majData_majTime_majComponent_majGen_genmajHour';

