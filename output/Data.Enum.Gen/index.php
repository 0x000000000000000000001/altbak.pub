<?php

namespace Data\Enum\Gen;

// ALL IMPORTS: Control.Applicative, Control.Monad.Gen, Data.Bounded, Data.Enum, Data.Foldable, Data.Maybe, Data.NonEmpty, Data.Unfoldable1, Prelude, Prim
// TO REQUIRE: Control.Applicative, Control.Monad.Gen, Data.Bounded, Data.Enum, Data.Foldable, Data.Maybe, Data.NonEmpty, Data.Unfoldable1, Prelude
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Monad.Gen/index.php';
require_once __DIR__ . '/../Data.Bounded/index.php';
require_once __DIR__ . '/../Data.Enum/index.php';
require_once __DIR__ . '/../Data.Foldable/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.NonEmpty/index.php';
require_once __DIR__ . '/../Data.Unfoldable1/index.php';
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




// Data_Enum_Gen_genBoundedEnum
function majData_majEnum_majGen_genmajBoundedmajEnum($dictMonadGen_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEnum_majGen_genmajBoundedmajEnum';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Applicative0_1_0 = (((($dictMonadGen_0)->{'Monad0'})(null))->{'Applicative0'})(null);
  $__res = function($dictBoundedEnum_2) use ($Applicative0_1_0, $dictMonadGen_0) {
  $__num = \func_num_args();
  $Enum1_3_1 = (($dictBoundedEnum_2)->{'Enum1'})(null);
  $Bounded0_4_2 = (($dictBoundedEnum_2)->{'Bounded0'})(null);
  $v_5_3 = (($Enum1_3_1)->{'succ'})(($Bounded0_4_2)->{'bottom'});
  $__t4 = null;;
  if ($v_5_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$Monad0_6_5 = (($dictMonadGen_0)->{'Monad0'})(null);
$pure_7_6 = ((($Monad0_6_5)->{'Applicative0'})(null))->{'pure'};
$foldableNonEmpty1_8_7 = (object)["foldMap" => function($dictMonoid_8) {
  $__num = \func_num_args();
  $Semigroup0_9_7 = (($dictMonoid_8)->{'Semigroup0'})(null);
  $__res = function($f_10) use ($Semigroup0_9_7, $dictMonoid_8) {
  $__num = \func_num_args();
  $__res = function($v_11) use ($Semigroup0_9_7, $dictMonoid_8, $f_10) {
  $__num = \func_num_args();
  $Semigroup0_12_8 = (($dictMonoid_8)->{'Semigroup0'})(null);
  $__res = ((($Semigroup0_9_7)->{'append'})(($f_10)(($v_11)->{'value0'})))(\Data\Foldable\majData_majFoldable_foldrmajArray(function($x_13) use ($Semigroup0_12_8, $f_10) {
  $__num = \func_num_args();
  $__res = function($acc_14) use ($Semigroup0_12_8, $f_10, $x_13) {
  $__num = \func_num_args();
  $__res = ((($Semigroup0_12_8)->{'append'})(($f_10)($x_13)))($acc_14);
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
$__local_var_8_7 = (object)["foldMap1" => function($dictSemigroup_9) {
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
  $__local_var_11_10 = ($f_9)(($v_10)->{'value0'});
  $__local_var_12_11 = \Data\Foldable\majData_majFoldable_foldrmajArray(function($a1_12) use ($f_9) {
  $__num = \func_num_args();
  $__local_var_13_11 = ($f_9)($a1_12);
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Maybe_Just']))(function($v2_14) use ($__local_var_13_11, $a1_12) {
  $__num = \func_num_args();
  $__t12 = null;;
  if ($v2_14 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t12 = $a1_12;
goto end_branch_12;;
};
  if ($v2_14 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t12 = ($__local_var_13_11)(($v2_14)->{'value0'});
goto end_branch_12;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t12 = null;
  end_branch_12:;
  $__res = $__t12;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, new \Data\Maybe\Data_Maybe_Nothing(), ($v_10)->{'value1'});
  $__t14 = null;;
  if ($__local_var_12_11 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t14 = ($v_10)->{'value0'};
goto end_branch_14;;
};
  if ($__local_var_12_11 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t14 = ($__local_var_11_10)(($__local_var_12_11)->{'value0'});
goto end_branch_14;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t14 = null;
  end_branch_14:;
  $__res = $__t14;
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
}, "Foldable0" => function($_dollar___unused_9) use ($foldableNonEmpty1_8_7) {
  $__num = \func_num_args();
  $__res = $foldableNonEmpty1_8_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
$Ord0_9_16 = (($Enum1_3_1)->{'Ord0'})(null);
$Ord01_10_17 = (($Enum1_3_1)->{'Ord0'})(null);
$__t22 = null;;
if (((((($Ord0_9_16)->{'Eq0'})(null))->{'eq'})(($v_5_3)->{'value0'}))(($Bounded0_4_2)->{'top'})) {
$__t22 = \Data\Unfoldable1\majData_majUnfoldable1_unfoldr1majArraymajImpl($GLOBALS['Data_Maybe_isNothing'], function($v_11) {
  $__num = \func_num_args();
  $__t23 = null;;
  if ($v_11 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t23 = ($v_11)->{'value0'};
goto end_branch_23;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t23 = null;
  end_branch_23:;
  $__res = $__t23;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $GLOBALS['Data_Tuple_fst'], $GLOBALS['Data_Tuple_snd'], function($i_11) use ($v_5_3) {
  $__num = \func_num_args();
  $__t24 = null;;
  if (($i_11 <= 0)) {
$__t24 = new \Data\Tuple\Data_Tuple_Tuple(($v_5_3)->{'value0'}, new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_24;;
};
  $__t24 = new \Data\Tuple\Data_Tuple_Tuple(($v_5_3)->{'value0'}, new \Data\Maybe\Data_Maybe_Just(($i_11 - 1)));
  end_branch_24:;
  $__res = $__t24;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, 0);
goto end_branch_22;;
};
if (((($Ord01_10_17)->{'compare'})(($v_5_3)->{'value0'}))(($Bounded0_4_2)->{'top'}) instanceof \Data\Ordering\Data_Ordering_LT) {
$__t22 = \Data\Unfoldable1\majData_majUnfoldable1_unfoldr1majArraymajImpl($GLOBALS['Data_Maybe_isNothing'], function($v_11) {
  $__num = \func_num_args();
  $__t25 = null;;
  if ($v_11 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t25 = ($v_11)->{'value0'};
goto end_branch_25;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t25 = null;
  end_branch_25:;
  $__res = $__t25;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $GLOBALS['Data_Tuple_fst'], $GLOBALS['Data_Tuple_snd'], function($a_11) use ($Bounded0_4_2, $Enum1_3_1, $Ord0_9_16) {
  $__num = \func_num_args();
  $__local_var_12_26 = (($Enum1_3_1)->{'succ'})($a_11);
  $__t27 = null;;
  if ($__local_var_12_26 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t28 = null;;
if (( ! ((($Ord0_9_16)->{'compare'})(($__local_var_12_26)->{'value0'}))(($Bounded0_4_2)->{'top'}) instanceof \Data\Ordering\Data_Ordering_GT)) {
$__t28 = new \Data\Maybe\Data_Maybe_Just(($__local_var_12_26)->{'value0'});
goto end_branch_28;;
};
$__t28 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_28:;
$__t27 = $__t28;
goto end_branch_27;;
};
  if ($__local_var_12_26 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t27 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_27;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t27 = null;
  end_branch_27:;
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_11, $__t27);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($v_5_3)->{'value0'});
goto end_branch_22;;
};
$__t22 = \Data\Unfoldable1\majData_majUnfoldable1_unfoldr1majArraymajImpl($GLOBALS['Data_Maybe_isNothing'], function($v_11) {
  $__num = \func_num_args();
  $__t18 = null;;
  if ($v_11 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t18 = ($v_11)->{'value0'};
goto end_branch_18;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t18 = null;
  end_branch_18:;
  $__res = $__t18;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $GLOBALS['Data_Tuple_fst'], $GLOBALS['Data_Tuple_snd'], function($a_11) use ($Bounded0_4_2, $Enum1_3_1, $Ord0_9_16) {
  $__num = \func_num_args();
  $__local_var_12_19 = (($Enum1_3_1)->{'pred'})($a_11);
  $__t20 = null;;
  if ($__local_var_12_19 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t21 = null;;
if (( ! ((($Ord0_9_16)->{'compare'})(($__local_var_12_19)->{'value0'}))(($Bounded0_4_2)->{'top'}) instanceof \Data\Ordering\Data_Ordering_LT)) {
$__t21 = new \Data\Maybe\Data_Maybe_Just(($__local_var_12_19)->{'value0'});
goto end_branch_21;;
};
$__t21 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_21:;
$__t20 = $__t21;
goto end_branch_20;;
};
  if ($__local_var_12_19 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t20 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_20;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t20 = null;
  end_branch_20:;
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_11, $__t20);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($v_5_3)->{'value0'});
end_branch_22:;
$__local_var_9_16 = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(($Bounded0_4_2)->{'bottom'}, $__t22);
$__t4 = ((((($Monad0_6_5)->{'Bind1'})(null))->{'bind'})(((($dictMonadGen_0)->{'chooseInt'})(0))(((((((($__local_var_8_7)->{'Foldable0'})(null))->{'foldl'})(function($c_10) {
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
}))(0))($__local_var_9_16) - 1))))(function($n_10) use ($__local_var_8_7, $__local_var_9_16, $pure_7_6) {
  $__num = \func_num_args();
  $go__go_11_30 = null;
  $go__go_11_30 = (function() use ($__local_var_8_7, $__local_var_9_16, &$go__go_11_30) {
  $__fn = function(int $v_12, $v1_13 = null) use ($__local_var_8_7, $__local_var_9_16, &$go__go_11_30, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_11_30_30_v_12 = $v_12;
  $__tco_var_go__go_11_30_30_v1_13 = $v1_13;
  tco_loop_go__go_11_30_30:;
  $v_12 = $__tco_var_go__go_11_30_30_v_12;
  $v1_13 = $__tco_var_go__go_11_30_30_v1_13;
  $__t30 = null;;
  if ($v1_13 instanceof \Control\Monad\Gen\Control_Monad_Gen_Cons) {
$__t33 = null;;
if (($v1_13)->{'value1'} instanceof \Control\Monad\Gen\Control_Monad_Gen_Nil) {
$__t33 = ($v1_13)->{'value0'};
goto end_branch_33;;
};
if (($v_12 <= 0)) {
$__t33 = ($v1_13)->{'value0'};
goto end_branch_33;;
};
$__tco_31 = ($v_12 - 1);
$__tco_32 = ($v1_13)->{'value1'};
$__tco_var_go__go_11_30_30_v_12 = $__tco_31;
$__tco_var_go__go_11_30_30_v1_13 = $__tco_32;
goto tco_loop_go__go_11_30_30;;
$__t33 = null;
end_branch_33:;
$__t30 = $__t33;
goto end_branch_30;;
};
  if ($v1_13 instanceof \Control\Monad\Gen\Control_Monad_Gen_Nil) {
$__t30 = (((($__local_var_8_7)->{'foldMap1'})($GLOBALS['Data_Semigroup_Last_semigroupLast']))(function($x_14) {
  $__num = \func_num_args();
  $__res = $x_14;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($__local_var_9_16);
goto end_branch_30;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t30 = null;
  end_branch_30:;
  $__res = $__t30;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ($pure_7_6)((($go__go_11_30)($n_10))((((((($__local_var_8_7)->{'Foldable0'})(null))->{'foldr'})($GLOBALS['Control_Monad_Gen_Cons']))(new \Control\Monad\Gen\Control_Monad_Gen_Nil()))($__local_var_9_16)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
goto end_branch_4;;
};
  if ($v_5_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t4 = (($Applicative0_1_0)->{'pure'})(($Bounded0_4_2)->{'bottom'});
goto end_branch_4;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t4 = null;
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Enum_Gen_genBoundedEnum'] = __NAMESPACE__ . '\\majData_majEnum_majGen_genmajBoundedmajEnum';

