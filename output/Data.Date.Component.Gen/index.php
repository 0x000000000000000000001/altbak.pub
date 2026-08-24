<?php

namespace Data\Date\Component\Gen;

// ALL IMPORTS: Control.Monad.Gen, Control.Monad.Gen.Class, Control.Semigroupoid, Data.Date.Component, Data.Enum, Data.Enum.Gen, Data.Functor, Data.Maybe, Partial.Unsafe, Prelude, Prim
// TO REQUIRE: Control.Monad.Gen, Control.Monad.Gen.Class, Control.Semigroupoid, Data.Date.Component, Data.Enum, Data.Enum.Gen, Data.Functor, Data.Maybe, Partial.Unsafe, Prelude
require_once __DIR__ . '/../Control.Monad.Gen/index.php';
require_once __DIR__ . '/../Control.Monad.Gen.Class/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Date.Component/index.php';
require_once __DIR__ . '/../Data.Enum/index.php';
require_once __DIR__ . '/../Data.Enum.Gen/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Partial.Unsafe/index.php';
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




// Data_Date_Component_Gen_genYear
function majData_majDate_majComponent_majGen_genmajYear($dictMonadGen_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majDate_majComponent_majGen_genmajYear';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ((((((((((($dictMonadGen_0)->{'Monad0'})(null))->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_1) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v_1 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t0 = ($v_1)->{'value0'};
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($n_1) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((($n_1 >= -271820) && ($n_1 <= 275759))) {
$__t1 = new \Data\Maybe\Data_Maybe_Just($n_1);
goto end_branch_1;;
};
  $__t1 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(((($dictMonadGen_0)->{'chooseInt'})(1900))(2100));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Date_Component_Gen_genYear'] = __NAMESPACE__ . '\\majData_majDate_majComponent_majGen_genmajYear';

// Data_Date_Component_Gen_genWeekday
function majData_majDate_majComponent_majGen_genmajWeekday($dictMonadGen_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majDate_majComponent_majGen_genmajWeekday';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Applicative0_1_0 = (((($dictMonadGen_0)->{'Monad0'})(null))->{'Applicative0'})(null);
  $v_2_1 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($v_2) {
  $__num = \func_num_args();
  $__res = match ($v_2) { 1 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_Monday()), 2 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_Tuesday()), 3 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_Wednesday()), 4 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_Thursday()), 5 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_Friday()), 6 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_Saturday()), 7 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_Sunday()), default => new \Data\Maybe\Data_Maybe_Nothing() };
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
  $__t2 = null;;
  if ($v_2 instanceof \Data\Date\Component\Data_Date_Component_Monday) {
$__t2 = 1;
goto end_branch_2;;
};
  if ($v_2 instanceof \Data\Date\Component\Data_Date_Component_Tuesday) {
$__t2 = 2;
goto end_branch_2;;
};
  if ($v_2 instanceof \Data\Date\Component\Data_Date_Component_Wednesday) {
$__t2 = 3;
goto end_branch_2;;
};
  if ($v_2 instanceof \Data\Date\Component\Data_Date_Component_Thursday) {
$__t2 = 4;
goto end_branch_2;;
};
  if ($v_2 instanceof \Data\Date\Component\Data_Date_Component_Friday) {
$__t2 = 5;
goto end_branch_2;;
};
  if ($v_2 instanceof \Data\Date\Component\Data_Date_Component_Saturday) {
$__t2 = 6;
goto end_branch_2;;
};
  if ($v_2 instanceof \Data\Date\Component\Data_Date_Component_Sunday) {
$__t2 = 7;
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), new \Data\Date\Component\Data_Date_Component_Monday());
  $__t4 = null;;
  if ($v_2_1 instanceof \Data\Maybe\Data_Maybe_Just) {
$Monad0_3_5 = (($dictMonadGen_0)->{'Monad0'})(null);
$pure_4_6 = ((($Monad0_3_5)->{'Applicative0'})(null))->{'pure'};
$foldableNonEmpty1_5_7 = (object)["foldMap" => function($dictMonoid_5) {
  $__num = \func_num_args();
  $Semigroup0_6_7 = (($dictMonoid_5)->{'Semigroup0'})(null);
  $__res = function($f_7) use ($Semigroup0_6_7, $dictMonoid_5) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($Semigroup0_6_7, $dictMonoid_5, $f_7) {
  $__num = \func_num_args();
  $Semigroup0_9_8 = (($dictMonoid_5)->{'Semigroup0'})(null);
  $__res = ((($Semigroup0_6_7)->{'append'})(($f_7)(($v_8)->{'value0'})))(\Data\Foldable\majData_majFoldable_foldrmajArray(function($x_10) use ($Semigroup0_9_8, $f_7) {
  $__num = \func_num_args();
  $__res = function($acc_11) use ($Semigroup0_9_8, $f_7, $x_10) {
  $__num = \func_num_args();
  $__res = ((($Semigroup0_9_8)->{'append'})(($f_7)($x_10)))($acc_11);
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
$__local_var_5_7 = (object)["foldMap1" => function($dictSemigroup_6) {
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
  $__local_var_8_10 = ($f_6)(($v_7)->{'value0'});
  $__local_var_9_11 = \Data\Foldable\majData_majFoldable_foldrmajArray(function($a1_9) use ($f_6) {
  $__num = \func_num_args();
  $__local_var_10_11 = ($f_6)($a1_9);
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Maybe_Just']))(function($v2_11) use ($__local_var_10_11, $a1_9) {
  $__num = \func_num_args();
  $__t12 = null;;
  if ($v2_11 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t12 = $a1_9;
goto end_branch_12;;
};
  if ($v2_11 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t12 = ($__local_var_10_11)(($v2_11)->{'value0'});
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
}, new \Data\Maybe\Data_Maybe_Nothing(), ($v_7)->{'value1'});
  $__t14 = null;;
  if ($__local_var_9_11 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t14 = ($v_7)->{'value0'};
goto end_branch_14;;
};
  if ($__local_var_9_11 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t14 = ($__local_var_8_10)(($__local_var_9_11)->{'value0'});
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
}, "Foldable0" => function($_dollar___unused_6) use ($foldableNonEmpty1_5_7) {
  $__num = \func_num_args();
  $__res = $foldableNonEmpty1_5_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
$__t23 = null;;
if (($v_2_1)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_Sunday) {
$__t23 = \Data\Unfoldable1\majData_majUnfoldable1_unfoldr1majArraymajImpl($GLOBALS['Data_Maybe_isNothing'], function($v_6) {
  $__num = \func_num_args();
  $__t24 = null;;
  if ($v_6 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t24 = ($v_6)->{'value0'};
goto end_branch_24;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t24 = null;
  end_branch_24:;
  $__res = $__t24;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $GLOBALS['Data_Tuple_fst'], $GLOBALS['Data_Tuple_snd'], function($i_6) use ($v_2_1) {
  $__num = \func_num_args();
  $__t25 = null;;
  if (($i_6 <= 0)) {
$__t25 = new \Data\Tuple\Data_Tuple_Tuple(($v_2_1)->{'value0'}, new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_25;;
};
  $__t25 = new \Data\Tuple\Data_Tuple_Tuple(($v_2_1)->{'value0'}, new \Data\Maybe\Data_Maybe_Just(($i_6 - 1)));
  end_branch_25:;
  $__res = $__t25;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, 0);
goto end_branch_23;;
};
if ((function() use ($v_2_1, &$__fn) {
$__t26 = null;;
if (($v_2_1)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_Monday) {
$__t26 = true;
goto end_branch_26;;
};
if (($v_2_1)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_Tuesday) {
$__t26 = true;
goto end_branch_26;;
};
if (($v_2_1)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_Wednesday) {
$__t26 = true;
goto end_branch_26;;
};
if (($v_2_1)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_Thursday) {
$__t26 = true;
goto end_branch_26;;
};
if (($v_2_1)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_Friday) {
$__t26 = true;
goto end_branch_26;;
};
if (($v_2_1)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_Saturday) {
$__t26 = true;
goto end_branch_26;;
};
if (($v_2_1)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_Sunday) {
$__t26 = false;
goto end_branch_26;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t26 = null;
end_branch_26:;
return $__t26;
})()) {
$__t23 = \Data\Unfoldable1\majData_majUnfoldable1_unfoldr1majArraymajImpl($GLOBALS['Data_Maybe_isNothing'], function($v_6) {
  $__num = \func_num_args();
  $__t27 = null;;
  if ($v_6 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t27 = ($v_6)->{'value0'};
goto end_branch_27;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t27 = null;
  end_branch_27:;
  $__res = $__t27;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $GLOBALS['Data_Tuple_fst'], $GLOBALS['Data_Tuple_snd'], function($a_6) {
  $__num = \func_num_args();
  $__local_var_7_28 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($v_7) {
  $__num = \func_num_args();
  $__res = match ($v_7) { 1 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_Monday()), 2 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_Tuesday()), 3 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_Wednesday()), 4 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_Thursday()), 5 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_Friday()), 6 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_Saturday()), 7 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_Sunday()), default => new \Data\Maybe\Data_Maybe_Nothing() };
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
  $__t29 = null;;
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_Monday) {
$__t29 = 1;
goto end_branch_29;;
};
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_Tuesday) {
$__t29 = 2;
goto end_branch_29;;
};
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_Wednesday) {
$__t29 = 3;
goto end_branch_29;;
};
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_Thursday) {
$__t29 = 4;
goto end_branch_29;;
};
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_Friday) {
$__t29 = 5;
goto end_branch_29;;
};
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_Saturday) {
$__t29 = 6;
goto end_branch_29;;
};
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_Sunday) {
$__t29 = 7;
goto end_branch_29;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t29 = null;
  end_branch_29:;
  $__res = $__t29;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), $a_6);
  $__t31 = null;;
  if ($__local_var_7_28 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t32 = null;;
if ((function() use ($__local_var_7_28, &$__fn) {
$__t33 = null;;
if (($__local_var_7_28)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_Monday) {
$__t33 = true;
goto end_branch_33;;
};
if (($__local_var_7_28)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_Tuesday) {
$__t33 = true;
goto end_branch_33;;
};
if (($__local_var_7_28)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_Wednesday) {
$__t33 = true;
goto end_branch_33;;
};
if (($__local_var_7_28)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_Thursday) {
$__t33 = true;
goto end_branch_33;;
};
if (($__local_var_7_28)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_Friday) {
$__t33 = true;
goto end_branch_33;;
};
if (($__local_var_7_28)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_Saturday) {
$__t33 = true;
goto end_branch_33;;
};
if (($__local_var_7_28)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_Sunday) {
$__t33 = true;
goto end_branch_33;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t33 = null;
end_branch_33:;
return $__t33;
})()) {
$__t32 = new \Data\Maybe\Data_Maybe_Just(($__local_var_7_28)->{'value0'});
goto end_branch_32;;
};
$__t32 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_32:;
$__t31 = $__t32;
goto end_branch_31;;
};
  if ($__local_var_7_28 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t31 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_31;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t31 = null;
  end_branch_31:;
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_6, $__t31);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($v_2_1)->{'value0'});
goto end_branch_23;;
};
$__t23 = \Data\Unfoldable1\majData_majUnfoldable1_unfoldr1majArraymajImpl($GLOBALS['Data_Maybe_isNothing'], function($v_6) {
  $__num = \func_num_args();
  $__t16 = null;;
  if ($v_6 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t16 = ($v_6)->{'value0'};
goto end_branch_16;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t16 = null;
  end_branch_16:;
  $__res = $__t16;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $GLOBALS['Data_Tuple_fst'], $GLOBALS['Data_Tuple_snd'], function($a_6) {
  $__num = \func_num_args();
  $__local_var_7_17 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($v_7) {
  $__num = \func_num_args();
  $__res = match ($v_7) { 1 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_Monday()), 2 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_Tuesday()), 3 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_Wednesday()), 4 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_Thursday()), 5 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_Friday()), 6 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_Saturday()), 7 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_Sunday()), default => new \Data\Maybe\Data_Maybe_Nothing() };
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
  $__t18 = null;;
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_Monday) {
$__t18 = 1;
goto end_branch_18;;
};
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_Tuesday) {
$__t18 = 2;
goto end_branch_18;;
};
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_Wednesday) {
$__t18 = 3;
goto end_branch_18;;
};
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_Thursday) {
$__t18 = 4;
goto end_branch_18;;
};
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_Friday) {
$__t18 = 5;
goto end_branch_18;;
};
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_Saturday) {
$__t18 = 6;
goto end_branch_18;;
};
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_Sunday) {
$__t18 = 7;
goto end_branch_18;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t18 = null;
  end_branch_18:;
  $__res = $__t18;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), $a_6);
  $__t20 = null;;
  if ($__local_var_7_17 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t21 = null;;
if ((function() use ($__local_var_7_17, &$__fn) {
$__t22 = null;;
if (($__local_var_7_17)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_Monday) {
$__t22 = false;
goto end_branch_22;;
};
if (($__local_var_7_17)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_Tuesday) {
$__t22 = false;
goto end_branch_22;;
};
if (($__local_var_7_17)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_Wednesday) {
$__t22 = false;
goto end_branch_22;;
};
if (($__local_var_7_17)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_Thursday) {
$__t22 = false;
goto end_branch_22;;
};
if (($__local_var_7_17)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_Friday) {
$__t22 = false;
goto end_branch_22;;
};
if (($__local_var_7_17)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_Saturday) {
$__t22 = false;
goto end_branch_22;;
};
if (($__local_var_7_17)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_Sunday) {
$__t22 = true;
goto end_branch_22;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t22 = null;
end_branch_22:;
return $__t22;
})()) {
$__t21 = new \Data\Maybe\Data_Maybe_Just(($__local_var_7_17)->{'value0'});
goto end_branch_21;;
};
$__t21 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_21:;
$__t20 = $__t21;
goto end_branch_20;;
};
  if ($__local_var_7_17 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t20 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_20;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t20 = null;
  end_branch_20:;
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_6, $__t20);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($v_2_1)->{'value0'});
end_branch_23:;
$__local_var_6_16 = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(new \Data\Date\Component\Data_Date_Component_Monday(), $__t23);
$__t4 = ((((($Monad0_3_5)->{'Bind1'})(null))->{'bind'})(((($dictMonadGen_0)->{'chooseInt'})(0))(((((((($__local_var_5_7)->{'Foldable0'})(null))->{'foldl'})(function($c_7) {
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
}))(0))($__local_var_6_16) - 1))))(function($n_7) use ($__local_var_5_7, $__local_var_6_16, $pure_4_6) {
  $__num = \func_num_args();
  $go__go_8_35 = null;
  $go__go_8_35 = (function() use ($__local_var_5_7, $__local_var_6_16, &$go__go_8_35) {
  $__fn = function(int $v_9, $v1_10 = null) use ($__local_var_5_7, $__local_var_6_16, &$go__go_8_35, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_8_35_35_v_9 = $v_9;
  $__tco_var_go__go_8_35_35_v1_10 = $v1_10;
  tco_loop_go__go_8_35_35:;
  $v_9 = $__tco_var_go__go_8_35_35_v_9;
  $v1_10 = $__tco_var_go__go_8_35_35_v1_10;
  $__t35 = null;;
  if ($v1_10 instanceof \Control\Monad\Gen\Control_Monad_Gen_Cons) {
$__t38 = null;;
if (($v1_10)->{'value1'} instanceof \Control\Monad\Gen\Control_Monad_Gen_Nil) {
$__t38 = ($v1_10)->{'value0'};
goto end_branch_38;;
};
if (($v_9 <= 0)) {
$__t38 = ($v1_10)->{'value0'};
goto end_branch_38;;
};
$__tco_36 = ($v_9 - 1);
$__tco_37 = ($v1_10)->{'value1'};
$__tco_var_go__go_8_35_35_v_9 = $__tco_36;
$__tco_var_go__go_8_35_35_v1_10 = $__tco_37;
goto tco_loop_go__go_8_35_35;;
$__t38 = null;
end_branch_38:;
$__t35 = $__t38;
goto end_branch_35;;
};
  if ($v1_10 instanceof \Control\Monad\Gen\Control_Monad_Gen_Nil) {
$__t35 = (((($__local_var_5_7)->{'foldMap1'})($GLOBALS['Data_Semigroup_Last_semigroupLast']))(function($x_11) {
  $__num = \func_num_args();
  $__res = $x_11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($__local_var_6_16);
goto end_branch_35;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t35 = null;
  end_branch_35:;
  $__res = $__t35;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ($pure_4_6)((($go__go_8_35)($n_7))((((((($__local_var_5_7)->{'Foldable0'})(null))->{'foldr'})($GLOBALS['Control_Monad_Gen_Cons']))(new \Control\Monad\Gen\Control_Monad_Gen_Nil()))($__local_var_6_16)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
goto end_branch_4;;
};
  if ($v_2_1 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t4 = (($Applicative0_1_0)->{'pure'})(new \Data\Date\Component\Data_Date_Component_Monday());
goto end_branch_4;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t4 = null;
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Date_Component_Gen_genWeekday'] = __NAMESPACE__ . '\\majData_majDate_majComponent_majGen_genmajWeekday';

// Data_Date_Component_Gen_genMonth
function majData_majDate_majComponent_majGen_genmajMonth($dictMonadGen_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majDate_majComponent_majGen_genmajMonth';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Applicative0_1_0 = (((($dictMonadGen_0)->{'Monad0'})(null))->{'Applicative0'})(null);
  $v_2_1 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($v_2) {
  $__num = \func_num_args();
  $__res = match ($v_2) { 1 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_January()), 2 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_February()), 3 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_March()), 4 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_April()), 5 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_May()), 6 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_June()), 7 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_July()), 8 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_August()), 9 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_September()), 10 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_October()), 11 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_November()), 12 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_December()), default => new \Data\Maybe\Data_Maybe_Nothing() };
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
  $__t2 = null;;
  if ($v_2 instanceof \Data\Date\Component\Data_Date_Component_January) {
$__t2 = 1;
goto end_branch_2;;
};
  if ($v_2 instanceof \Data\Date\Component\Data_Date_Component_February) {
$__t2 = 2;
goto end_branch_2;;
};
  if ($v_2 instanceof \Data\Date\Component\Data_Date_Component_March) {
$__t2 = 3;
goto end_branch_2;;
};
  if ($v_2 instanceof \Data\Date\Component\Data_Date_Component_April) {
$__t2 = 4;
goto end_branch_2;;
};
  if ($v_2 instanceof \Data\Date\Component\Data_Date_Component_May) {
$__t2 = 5;
goto end_branch_2;;
};
  if ($v_2 instanceof \Data\Date\Component\Data_Date_Component_June) {
$__t2 = 6;
goto end_branch_2;;
};
  if ($v_2 instanceof \Data\Date\Component\Data_Date_Component_July) {
$__t2 = 7;
goto end_branch_2;;
};
  if ($v_2 instanceof \Data\Date\Component\Data_Date_Component_August) {
$__t2 = 8;
goto end_branch_2;;
};
  if ($v_2 instanceof \Data\Date\Component\Data_Date_Component_September) {
$__t2 = 9;
goto end_branch_2;;
};
  if ($v_2 instanceof \Data\Date\Component\Data_Date_Component_October) {
$__t2 = 10;
goto end_branch_2;;
};
  if ($v_2 instanceof \Data\Date\Component\Data_Date_Component_November) {
$__t2 = 11;
goto end_branch_2;;
};
  if ($v_2 instanceof \Data\Date\Component\Data_Date_Component_December) {
$__t2 = 12;
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), new \Data\Date\Component\Data_Date_Component_January());
  $__t4 = null;;
  if ($v_2_1 instanceof \Data\Maybe\Data_Maybe_Just) {
$Monad0_3_5 = (($dictMonadGen_0)->{'Monad0'})(null);
$pure_4_6 = ((($Monad0_3_5)->{'Applicative0'})(null))->{'pure'};
$foldableNonEmpty1_5_7 = (object)["foldMap" => function($dictMonoid_5) {
  $__num = \func_num_args();
  $Semigroup0_6_7 = (($dictMonoid_5)->{'Semigroup0'})(null);
  $__res = function($f_7) use ($Semigroup0_6_7, $dictMonoid_5) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($Semigroup0_6_7, $dictMonoid_5, $f_7) {
  $__num = \func_num_args();
  $Semigroup0_9_8 = (($dictMonoid_5)->{'Semigroup0'})(null);
  $__res = ((($Semigroup0_6_7)->{'append'})(($f_7)(($v_8)->{'value0'})))(\Data\Foldable\majData_majFoldable_foldrmajArray(function($x_10) use ($Semigroup0_9_8, $f_7) {
  $__num = \func_num_args();
  $__res = function($acc_11) use ($Semigroup0_9_8, $f_7, $x_10) {
  $__num = \func_num_args();
  $__res = ((($Semigroup0_9_8)->{'append'})(($f_7)($x_10)))($acc_11);
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
$__local_var_5_7 = (object)["foldMap1" => function($dictSemigroup_6) {
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
  $__local_var_8_10 = ($f_6)(($v_7)->{'value0'});
  $__local_var_9_11 = \Data\Foldable\majData_majFoldable_foldrmajArray(function($a1_9) use ($f_6) {
  $__num = \func_num_args();
  $__local_var_10_11 = ($f_6)($a1_9);
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Maybe_Just']))(function($v2_11) use ($__local_var_10_11, $a1_9) {
  $__num = \func_num_args();
  $__t12 = null;;
  if ($v2_11 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t12 = $a1_9;
goto end_branch_12;;
};
  if ($v2_11 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t12 = ($__local_var_10_11)(($v2_11)->{'value0'});
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
}, new \Data\Maybe\Data_Maybe_Nothing(), ($v_7)->{'value1'});
  $__t14 = null;;
  if ($__local_var_9_11 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t14 = ($v_7)->{'value0'};
goto end_branch_14;;
};
  if ($__local_var_9_11 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t14 = ($__local_var_8_10)(($__local_var_9_11)->{'value0'});
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
}, "Foldable0" => function($_dollar___unused_6) use ($foldableNonEmpty1_5_7) {
  $__num = \func_num_args();
  $__res = $foldableNonEmpty1_5_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
$__t23 = null;;
if (($v_2_1)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_December) {
$__t23 = \Data\Unfoldable1\majData_majUnfoldable1_unfoldr1majArraymajImpl($GLOBALS['Data_Maybe_isNothing'], function($v_6) {
  $__num = \func_num_args();
  $__t24 = null;;
  if ($v_6 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t24 = ($v_6)->{'value0'};
goto end_branch_24;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t24 = null;
  end_branch_24:;
  $__res = $__t24;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $GLOBALS['Data_Tuple_fst'], $GLOBALS['Data_Tuple_snd'], function($i_6) use ($v_2_1) {
  $__num = \func_num_args();
  $__t25 = null;;
  if (($i_6 <= 0)) {
$__t25 = new \Data\Tuple\Data_Tuple_Tuple(($v_2_1)->{'value0'}, new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_25;;
};
  $__t25 = new \Data\Tuple\Data_Tuple_Tuple(($v_2_1)->{'value0'}, new \Data\Maybe\Data_Maybe_Just(($i_6 - 1)));
  end_branch_25:;
  $__res = $__t25;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, 0);
goto end_branch_23;;
};
if ((function() use ($v_2_1, &$__fn) {
$__t26 = null;;
if (($v_2_1)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_January) {
$__t26 = true;
goto end_branch_26;;
};
if (($v_2_1)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_February) {
$__t26 = true;
goto end_branch_26;;
};
if (($v_2_1)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_March) {
$__t26 = true;
goto end_branch_26;;
};
if (($v_2_1)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_April) {
$__t26 = true;
goto end_branch_26;;
};
if (($v_2_1)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_May) {
$__t26 = true;
goto end_branch_26;;
};
if (($v_2_1)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_June) {
$__t26 = true;
goto end_branch_26;;
};
if (($v_2_1)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_July) {
$__t26 = true;
goto end_branch_26;;
};
if (($v_2_1)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_August) {
$__t26 = true;
goto end_branch_26;;
};
if (($v_2_1)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_September) {
$__t26 = true;
goto end_branch_26;;
};
if (($v_2_1)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_October) {
$__t26 = true;
goto end_branch_26;;
};
if (($v_2_1)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_November) {
$__t26 = true;
goto end_branch_26;;
};
if (($v_2_1)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_December) {
$__t26 = false;
goto end_branch_26;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t26 = null;
end_branch_26:;
return $__t26;
})()) {
$__t23 = \Data\Unfoldable1\majData_majUnfoldable1_unfoldr1majArraymajImpl($GLOBALS['Data_Maybe_isNothing'], function($v_6) {
  $__num = \func_num_args();
  $__t27 = null;;
  if ($v_6 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t27 = ($v_6)->{'value0'};
goto end_branch_27;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t27 = null;
  end_branch_27:;
  $__res = $__t27;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $GLOBALS['Data_Tuple_fst'], $GLOBALS['Data_Tuple_snd'], function($a_6) {
  $__num = \func_num_args();
  $__local_var_7_28 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($v_7) {
  $__num = \func_num_args();
  $__res = match ($v_7) { 1 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_January()), 2 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_February()), 3 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_March()), 4 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_April()), 5 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_May()), 6 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_June()), 7 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_July()), 8 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_August()), 9 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_September()), 10 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_October()), 11 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_November()), 12 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_December()), default => new \Data\Maybe\Data_Maybe_Nothing() };
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
  $__t29 = null;;
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_January) {
$__t29 = 1;
goto end_branch_29;;
};
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_February) {
$__t29 = 2;
goto end_branch_29;;
};
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_March) {
$__t29 = 3;
goto end_branch_29;;
};
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_April) {
$__t29 = 4;
goto end_branch_29;;
};
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_May) {
$__t29 = 5;
goto end_branch_29;;
};
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_June) {
$__t29 = 6;
goto end_branch_29;;
};
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_July) {
$__t29 = 7;
goto end_branch_29;;
};
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_August) {
$__t29 = 8;
goto end_branch_29;;
};
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_September) {
$__t29 = 9;
goto end_branch_29;;
};
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_October) {
$__t29 = 10;
goto end_branch_29;;
};
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_November) {
$__t29 = 11;
goto end_branch_29;;
};
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_December) {
$__t29 = 12;
goto end_branch_29;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t29 = null;
  end_branch_29:;
  $__res = $__t29;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), $a_6);
  $__t31 = null;;
  if ($__local_var_7_28 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t32 = null;;
if ((function() use ($__local_var_7_28, &$__fn) {
$__t33 = null;;
if (($__local_var_7_28)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_January) {
$__t33 = true;
goto end_branch_33;;
};
if (($__local_var_7_28)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_February) {
$__t33 = true;
goto end_branch_33;;
};
if (($__local_var_7_28)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_March) {
$__t33 = true;
goto end_branch_33;;
};
if (($__local_var_7_28)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_April) {
$__t33 = true;
goto end_branch_33;;
};
if (($__local_var_7_28)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_May) {
$__t33 = true;
goto end_branch_33;;
};
if (($__local_var_7_28)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_June) {
$__t33 = true;
goto end_branch_33;;
};
if (($__local_var_7_28)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_July) {
$__t33 = true;
goto end_branch_33;;
};
if (($__local_var_7_28)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_August) {
$__t33 = true;
goto end_branch_33;;
};
if (($__local_var_7_28)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_September) {
$__t33 = true;
goto end_branch_33;;
};
if (($__local_var_7_28)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_October) {
$__t33 = true;
goto end_branch_33;;
};
if (($__local_var_7_28)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_November) {
$__t33 = true;
goto end_branch_33;;
};
if (($__local_var_7_28)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_December) {
$__t33 = true;
goto end_branch_33;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t33 = null;
end_branch_33:;
return $__t33;
})()) {
$__t32 = new \Data\Maybe\Data_Maybe_Just(($__local_var_7_28)->{'value0'});
goto end_branch_32;;
};
$__t32 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_32:;
$__t31 = $__t32;
goto end_branch_31;;
};
  if ($__local_var_7_28 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t31 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_31;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t31 = null;
  end_branch_31:;
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_6, $__t31);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($v_2_1)->{'value0'});
goto end_branch_23;;
};
$__t23 = \Data\Unfoldable1\majData_majUnfoldable1_unfoldr1majArraymajImpl($GLOBALS['Data_Maybe_isNothing'], function($v_6) {
  $__num = \func_num_args();
  $__t16 = null;;
  if ($v_6 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t16 = ($v_6)->{'value0'};
goto end_branch_16;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t16 = null;
  end_branch_16:;
  $__res = $__t16;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $GLOBALS['Data_Tuple_fst'], $GLOBALS['Data_Tuple_snd'], function($a_6) {
  $__num = \func_num_args();
  $__local_var_7_17 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($v_7) {
  $__num = \func_num_args();
  $__res = match ($v_7) { 1 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_January()), 2 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_February()), 3 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_March()), 4 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_April()), 5 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_May()), 6 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_June()), 7 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_July()), 8 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_August()), 9 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_September()), 10 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_October()), 11 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_November()), 12 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_December()), default => new \Data\Maybe\Data_Maybe_Nothing() };
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
  $__t18 = null;;
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_January) {
$__t18 = 1;
goto end_branch_18;;
};
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_February) {
$__t18 = 2;
goto end_branch_18;;
};
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_March) {
$__t18 = 3;
goto end_branch_18;;
};
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_April) {
$__t18 = 4;
goto end_branch_18;;
};
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_May) {
$__t18 = 5;
goto end_branch_18;;
};
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_June) {
$__t18 = 6;
goto end_branch_18;;
};
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_July) {
$__t18 = 7;
goto end_branch_18;;
};
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_August) {
$__t18 = 8;
goto end_branch_18;;
};
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_September) {
$__t18 = 9;
goto end_branch_18;;
};
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_October) {
$__t18 = 10;
goto end_branch_18;;
};
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_November) {
$__t18 = 11;
goto end_branch_18;;
};
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_December) {
$__t18 = 12;
goto end_branch_18;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t18 = null;
  end_branch_18:;
  $__res = $__t18;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), $a_6);
  $__t20 = null;;
  if ($__local_var_7_17 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t21 = null;;
if ((function() use ($__local_var_7_17, &$__fn) {
$__t22 = null;;
if (($__local_var_7_17)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_January) {
$__t22 = false;
goto end_branch_22;;
};
if (($__local_var_7_17)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_February) {
$__t22 = false;
goto end_branch_22;;
};
if (($__local_var_7_17)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_March) {
$__t22 = false;
goto end_branch_22;;
};
if (($__local_var_7_17)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_April) {
$__t22 = false;
goto end_branch_22;;
};
if (($__local_var_7_17)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_May) {
$__t22 = false;
goto end_branch_22;;
};
if (($__local_var_7_17)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_June) {
$__t22 = false;
goto end_branch_22;;
};
if (($__local_var_7_17)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_July) {
$__t22 = false;
goto end_branch_22;;
};
if (($__local_var_7_17)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_August) {
$__t22 = false;
goto end_branch_22;;
};
if (($__local_var_7_17)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_September) {
$__t22 = false;
goto end_branch_22;;
};
if (($__local_var_7_17)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_October) {
$__t22 = false;
goto end_branch_22;;
};
if (($__local_var_7_17)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_November) {
$__t22 = false;
goto end_branch_22;;
};
if (($__local_var_7_17)->{'value0'} instanceof \Data\Date\Component\Data_Date_Component_December) {
$__t22 = true;
goto end_branch_22;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t22 = null;
end_branch_22:;
return $__t22;
})()) {
$__t21 = new \Data\Maybe\Data_Maybe_Just(($__local_var_7_17)->{'value0'});
goto end_branch_21;;
};
$__t21 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_21:;
$__t20 = $__t21;
goto end_branch_20;;
};
  if ($__local_var_7_17 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t20 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_20;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t20 = null;
  end_branch_20:;
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_6, $__t20);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($v_2_1)->{'value0'});
end_branch_23:;
$__local_var_6_16 = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(new \Data\Date\Component\Data_Date_Component_January(), $__t23);
$__t4 = ((((($Monad0_3_5)->{'Bind1'})(null))->{'bind'})(((($dictMonadGen_0)->{'chooseInt'})(0))(((((((($__local_var_5_7)->{'Foldable0'})(null))->{'foldl'})(function($c_7) {
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
}))(0))($__local_var_6_16) - 1))))(function($n_7) use ($__local_var_5_7, $__local_var_6_16, $pure_4_6) {
  $__num = \func_num_args();
  $go__go_8_35 = null;
  $go__go_8_35 = (function() use ($__local_var_5_7, $__local_var_6_16, &$go__go_8_35) {
  $__fn = function(int $v_9, $v1_10 = null) use ($__local_var_5_7, $__local_var_6_16, &$go__go_8_35, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_8_35_35_v_9 = $v_9;
  $__tco_var_go__go_8_35_35_v1_10 = $v1_10;
  tco_loop_go__go_8_35_35:;
  $v_9 = $__tco_var_go__go_8_35_35_v_9;
  $v1_10 = $__tco_var_go__go_8_35_35_v1_10;
  $__t35 = null;;
  if ($v1_10 instanceof \Control\Monad\Gen\Control_Monad_Gen_Cons) {
$__t38 = null;;
if (($v1_10)->{'value1'} instanceof \Control\Monad\Gen\Control_Monad_Gen_Nil) {
$__t38 = ($v1_10)->{'value0'};
goto end_branch_38;;
};
if (($v_9 <= 0)) {
$__t38 = ($v1_10)->{'value0'};
goto end_branch_38;;
};
$__tco_36 = ($v_9 - 1);
$__tco_37 = ($v1_10)->{'value1'};
$__tco_var_go__go_8_35_35_v_9 = $__tco_36;
$__tco_var_go__go_8_35_35_v1_10 = $__tco_37;
goto tco_loop_go__go_8_35_35;;
$__t38 = null;
end_branch_38:;
$__t35 = $__t38;
goto end_branch_35;;
};
  if ($v1_10 instanceof \Control\Monad\Gen\Control_Monad_Gen_Nil) {
$__t35 = (((($__local_var_5_7)->{'foldMap1'})($GLOBALS['Data_Semigroup_Last_semigroupLast']))(function($x_11) {
  $__num = \func_num_args();
  $__res = $x_11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($__local_var_6_16);
goto end_branch_35;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t35 = null;
  end_branch_35:;
  $__res = $__t35;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ($pure_4_6)((($go__go_8_35)($n_7))((((((($__local_var_5_7)->{'Foldable0'})(null))->{'foldr'})($GLOBALS['Control_Monad_Gen_Cons']))(new \Control\Monad\Gen\Control_Monad_Gen_Nil()))($__local_var_6_16)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
goto end_branch_4;;
};
  if ($v_2_1 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t4 = (($Applicative0_1_0)->{'pure'})(new \Data\Date\Component\Data_Date_Component_January());
goto end_branch_4;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t4 = null;
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Date_Component_Gen_genMonth'] = __NAMESPACE__ . '\\majData_majDate_majComponent_majGen_genmajMonth';

// Data_Date_Component_Gen_genDay
function majData_majDate_majComponent_majGen_genmajDay($dictMonadGen_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majDate_majComponent_majGen_genmajDay';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Applicative0_1_0 = (((($dictMonadGen_0)->{'Monad0'})(null))->{'Applicative0'})(null);
  $v_2_1 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_2) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((($n_2 >= 1) && ($n_2 <= 31))) {
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
}), 1);
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
case 31:
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
if ((($v_2_1)->{'value0'} < 31)) {
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
  if ((($n_7 >= 1) && ($n_7 <= 31))) {
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
if ((($__local_var_7_24)->{'value0'} <= 31)) {
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
  if ((($n_7 >= 1) && ($n_7 <= 31))) {
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
if ((($__local_var_7_16)->{'value0'} >= 31)) {
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
$__local_var_6_15 = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(1, $__t20);
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
$__t3 = (($Applicative0_1_0)->{'pure'})(1);
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
$GLOBALS['Data_Date_Component_Gen_genDay'] = __NAMESPACE__ . '\\majData_majDate_majComponent_majGen_genmajDay';

