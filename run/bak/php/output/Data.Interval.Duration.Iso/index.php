<?php

namespace Data\Interval\Duration\Iso;

// ALL IMPORTS: Control.Applicative, Control.Plus, Control.Semigroupoid, Data.Either, Data.Eq, Data.Foldable, Data.Function, Data.HeytingAlgebra, Data.Interval.Duration, Data.Interval.Duration.Iso, Data.List, Data.List.NonEmpty, Data.List.Types, Data.Map, Data.Map.Internal, Data.Maybe, Data.Monoid, Data.Monoid.Additive, Data.Newtype, Data.Number, Data.Ord, Data.Ordering, Data.Semigroup, Data.Semiring, Data.Show, Data.Tuple, Prelude, Prim
// TO REQUIRE: Control.Applicative, Control.Plus, Control.Semigroupoid, Data.Either, Data.Eq, Data.Foldable, Data.Function, Data.HeytingAlgebra, Data.Interval.Duration, Data.Interval.Duration.Iso, Data.List, Data.List.NonEmpty, Data.List.Types, Data.Map, Data.Map.Internal, Data.Maybe, Data.Monoid, Data.Monoid.Additive, Data.Newtype, Data.Number, Data.Ord, Data.Ordering, Data.Semigroup, Data.Semiring, Data.Show, Data.Tuple, Prelude
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Plus/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Either/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Foldable/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.HeytingAlgebra/index.php';
require_once __DIR__ . '/../Data.Interval.Duration/index.php';
require_once __DIR__ . '/../Data.Interval.Duration.Iso/index.php';
require_once __DIR__ . '/../Data.List/index.php';
require_once __DIR__ . '/../Data.List.NonEmpty/index.php';
require_once __DIR__ . '/../Data.List.Types/index.php';
require_once __DIR__ . '/../Data.Map/index.php';
require_once __DIR__ . '/../Data.Map.Internal/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Monoid/index.php';
require_once __DIR__ . '/../Data.Monoid.Additive/index.php';
require_once __DIR__ . '/../Data.Newtype/index.php';
require_once __DIR__ . '/../Data.Number/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ordering/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Semiring/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
require_once __DIR__ . '/../Data.Tuple/index.php';
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


final class Data_Interval_Duration_Iso_IsEmpty { public function __construct() {} }
final class Data_Interval_Duration_Iso_InvalidWeekComponentUsage { public function __construct() {} }
final class Data_Interval_Duration_Iso_ContainsNegativeValue { public function __construct(public mixed $value0) {} }
final class Data_Interval_Duration_Iso_InvalidFractionalUse { public function __construct(public mixed $value0) {} }

// Data_Interval_Duration_Iso_lookup
function majData_majInterval_majDuration_majIso_lookup($k_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majInterval_majDuration_majIso_lookup';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $go__1_0 = null;
  $go__1_0 = function($v_2 = null) use (&$go__1_0, $k_0) {
  $__num = \func_num_args();
  $__tco_var_go__1_0_0_v_2 = $v_2;
  tco_loop_go__1_0_0:;
  $v_2 = $__tco_var_go__1_0_0_v_2;
  $__t0 = null;;
  if ($v_2 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t0 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_0;;
};
  if ($v_2 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$v1_3_1 = (($GLOBALS['Data_Interval_Duration_ordDurationComponent'])['compare'])($k_0, ($v_2)->{'value2'});
$__t2 = null;;
if ($v1_3_1 instanceof \Data\Ordering\Data_Ordering_LT) {
$__tco_3 = ($v_2)->{'value4'};
$__tco_var_go__1_0_0_v_2 = $__tco_3;
goto tco_loop_go__1_0_0;;
$__t2 = null;
goto end_branch_2;;
};
if ($v1_3_1 instanceof \Data\Ordering\Data_Ordering_GT) {
$__tco_4 = ($v_2)->{'value5'};
$__tco_var_go__1_0_0_v_2 = $__tco_4;
goto tco_loop_go__1_0_0;;
$__t2 = null;
goto end_branch_2;;
};
if ($v1_3_1 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t2 = new \Data\Maybe\Data_Maybe_Just(($v_2)->{'value3'});
goto end_branch_2;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
end_branch_2:;
$__t0 = $__t2;
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = $go__1_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Interval_Duration_Iso_lookup'] = __NAMESPACE__ . '\\majData_majInterval_majDuration_majIso_lookup';

// Data_Interval_Duration_Iso_greaterThan
$GLOBALS['Data_Interval_Duration_Iso_greaterThan'] = (function() use (&$__fn) {
$__local_var_0_0 = ($GLOBALS['Data_Ord_ordIntImpl'])(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT());
return (function() use ($__local_var_0_0) {
  $__fn = function($a1_1 = null, $a2_2 = null) use ($__local_var_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($__local_var_0_0)($a1_1, $a2_2) instanceof \Data\Ordering\Data_Ordering_GT;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
})();

// Data_Interval_Duration_Iso_foldMap1_closure
$GLOBALS['Data_Interval_Duration_Iso_foldMap1_closure'] = (($GLOBALS['Data_List_Types_foldableList'])['foldMap'])($GLOBALS['Data_List_Types_monoidList']);

// Data_Interval_Duration_Iso_foldMap1
function majData_majInterval_majDuration_majIso_foldmajMap1($v_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majInterval_majDuration_majIso_foldmajMap1';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Interval_Duration_Iso_foldMap1_closure'])($v_0, $v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Interval_Duration_Iso_foldMap1'] = __NAMESPACE__ . '\\majData_majInterval_majDuration_majIso_foldmajMap1';

// Data_Interval_Duration_Iso_foldMap2
$GLOBALS['Data_Interval_Duration_Iso_foldMap2'] = (function() use (&$__fn) {
$semigroupAdditive1_0_0 = ["append" => (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_Semiring_semiringNumber'])['add'])($v_0, $v1_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
return (($GLOBALS['Data_List_Types_foldableList'])['foldMap'])(["mempty" => ($GLOBALS['Data_Semiring_semiringNumber'])['zero'], "Semigroup0" => function($_dollar__unused_1 = null) use ($semigroupAdditive1_0_0) {
  $__num = \func_num_args();
  $__res = $semigroupAdditive1_0_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]);
})();

// Data_Interval_Duration_Iso_fold
$GLOBALS['Data_Interval_Duration_Iso_fold'] = (function() use (&$__fn) {
$__local_var_0_0 = (($GLOBALS['Data_List_Types_monoidList'])['Semigroup0'])(null);
$semigroupFn_1_1 = ["append" => (function() use ($__local_var_0_0) {
  $__fn = function($f_1 = null, $g_2 = null, $x_3 = null) use ($__local_var_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($__local_var_0_0)['append'])(($f_1)($x_3), ($g_2)($x_3));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()];
return (($GLOBALS['Data_Foldable_foldableArray'])['foldMap'])(["mempty" => function($v_2 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_List_Types_monoidList'])['mempty'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Semigroup0" => function($_dollar__unused_2 = null) use ($semigroupFn_1_1) {
  $__num = \func_num_args();
  $__res = $semigroupFn_1_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}], $GLOBALS['Data_Foldable_identity']);
})();

// Data_Interval_Duration_Iso_toUnfoldable
$GLOBALS['Data_Interval_Duration_Iso_toUnfoldable'] = ($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_List_Types_unfoldableList'])['unfoldr'])($GLOBALS['Data_Map_Internal_stepUnfoldr']), $GLOBALS['Data_Map_Internal_toMapIter']);

// Data_Interval_Duration_Iso_IsEmpty
$GLOBALS['Data_Interval_Duration_Iso_IsEmpty'] = ($GLOBALS['__phpurs_data0_IsEmpty'] ??= new \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_IsEmpty());

// Data_Interval_Duration_Iso_InvalidWeekComponentUsage
$GLOBALS['Data_Interval_Duration_Iso_InvalidWeekComponentUsage'] = ($GLOBALS['__phpurs_data0_InvalidWeekComponentUsage'] ??= new \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_InvalidWeekComponentUsage());

// Data_Interval_Duration_Iso_ContainsNegativeValue
$GLOBALS['Data_Interval_Duration_Iso_ContainsNegativeValue'] = function($value0 = null) {
  $__num = \func_num_args();
  $__res = new \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_ContainsNegativeValue($value0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Interval_Duration_Iso_InvalidFractionalUse
$GLOBALS['Data_Interval_Duration_Iso_InvalidFractionalUse'] = function($value0 = null) {
  $__num = \func_num_args();
  $__res = new \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_InvalidFractionalUse($value0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Interval_Duration_Iso_unIsoDuration
function majData_majInterval_majDuration_majIso_unmajIsomajDuration($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majInterval_majDuration_majIso_unmajIsomajDuration';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $v_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Interval_Duration_Iso_unIsoDuration'] = __NAMESPACE__ . '\\majData_majInterval_majDuration_majIso_unmajIsomajDuration';

// Data_Interval_Duration_Iso_showIsoDuration
$GLOBALS['Data_Interval_Duration_Iso_showIsoDuration'] = ["show" => function($v_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Semigroup_semigroupString'])['append'])("(IsoDuration ", (($GLOBALS['Data_Semigroup_semigroupString'])['append'])((($GLOBALS['Data_Interval_Duration_showDuration'])['show'])($v_0), ")"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Interval_Duration_Iso_showError
$GLOBALS['Data_Interval_Duration_Iso_showError'] = ["show" => function($v_0 = null) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v_0 instanceof \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_IsEmpty) {
$__t0 = "(IsEmpty)";
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_InvalidWeekComponentUsage) {
$__t0 = "(InvalidWeekComponentUsage)";
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_ContainsNegativeValue) {
$__t0 = (($GLOBALS['Data_Semigroup_semigroupString'])['append'])("(ContainsNegativeValue ", (($GLOBALS['Data_Semigroup_semigroupString'])['append'])((($GLOBALS['Data_Interval_Duration_showDurationComponent'])['show'])(($v_0)->{'value0'}), ")"));
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_InvalidFractionalUse) {
$__t0 = (($GLOBALS['Data_Semigroup_semigroupString'])['append'])("(InvalidFractionalUse ", (($GLOBALS['Data_Semigroup_semigroupString'])['append'])((($GLOBALS['Data_Interval_Duration_showDurationComponent'])['show'])(($v_0)->{'value0'}), ")"));
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Interval_Duration_Iso_prettyError
function majData_majInterval_majDuration_majIso_prettymajError($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majInterval_majDuration_majIso_prettymajError';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__t0 = null;;
  if ($v_0 instanceof \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_IsEmpty) {
$__t0 = "Duration is empty (has no components)";
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_InvalidWeekComponentUsage) {
$__t0 = "Week component of Duration is used with other components";
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_ContainsNegativeValue) {
$__t0 = (($GLOBALS['Data_Semigroup_semigroupString'])['append'])("Component `", (($GLOBALS['Data_Semigroup_semigroupString'])['append'])((($GLOBALS['Data_Interval_Duration_showDurationComponent'])['show'])(($v_0)->{'value0'}), "` contains negative value"));
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_InvalidFractionalUse) {
$__t0 = (($GLOBALS['Data_Semigroup_semigroupString'])['append'])("Invalid usage of Fractional value at component `", (($GLOBALS['Data_Semigroup_semigroupString'])['append'])((($GLOBALS['Data_Interval_Duration_showDurationComponent'])['show'])(($v_0)->{'value0'}), "`"));
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Interval_Duration_Iso_prettyError'] = __NAMESPACE__ . '\\majData_majInterval_majDuration_majIso_prettymajError';

// Data_Interval_Duration_Iso_eqIsoDuration
$GLOBALS['Data_Interval_Duration_Iso_eqIsoDuration'] = ["eq" => (function() {
  $__fn = function($x_0 = null, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_Interval_Duration_eqDuration'])['eq'])($x_0, $y_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];

// Data_Interval_Duration_Iso_ordIsoDuration
$GLOBALS['Data_Interval_Duration_Iso_ordIsoDuration'] = ["compare" => (function() {
  $__fn = function($x_0 = null, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_Interval_Duration_ordDuration'])['compare'])($x_0, $y_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Interval_Duration_Iso_eqIsoDuration'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Interval_Duration_Iso_eqError
$GLOBALS['Data_Interval_Duration_Iso_eqError'] = ["eq" => (function() {
  $__fn = function($x_0 = null, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ($x_0 instanceof \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_IsEmpty) {
$__t0 = $y_1 instanceof \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_IsEmpty;
goto end_branch_0;;
};
  if ($x_0 instanceof \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_InvalidWeekComponentUsage) {
$__t0 = $y_1 instanceof \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_InvalidWeekComponentUsage;
goto end_branch_0;;
};
  if ($x_0 instanceof \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_ContainsNegativeValue) {
$__t0 = ($y_1 instanceof \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_ContainsNegativeValue && (($GLOBALS['Data_Interval_Duration_eqDurationComponent'])['eq'])(($x_0)->{'value0'}, ($y_1)->{'value0'}));
goto end_branch_0;;
};
  $__t0 = ($x_0 instanceof \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_InvalidFractionalUse && ($y_1 instanceof \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_InvalidFractionalUse && (($GLOBALS['Data_Interval_Duration_eqDurationComponent'])['eq'])(($x_0)->{'value0'}, ($y_1)->{'value0'})));
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];

// Data_Interval_Duration_Iso_ordError
$GLOBALS['Data_Interval_Duration_Iso_ordError'] = ["compare" => (function() {
  $__fn = function($x_0 = null, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ($x_0 instanceof \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_IsEmpty) {
$__t1 = null;;
if ($y_1 instanceof \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_IsEmpty) {
$__t1 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_1;;
};
$__t1 = new \Data\Ordering\Data_Ordering_LT();
end_branch_1:;
$__t0 = $__t1;
goto end_branch_0;;
};
  if ($y_1 instanceof \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_IsEmpty) {
$__t0 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_0;;
};
  if ($x_0 instanceof \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_InvalidWeekComponentUsage) {
$__t2 = null;;
if ($y_1 instanceof \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_InvalidWeekComponentUsage) {
$__t2 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_2;;
};
$__t2 = new \Data\Ordering\Data_Ordering_LT();
end_branch_2:;
$__t0 = $__t2;
goto end_branch_0;;
};
  if ($y_1 instanceof \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_InvalidWeekComponentUsage) {
$__t0 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_0;;
};
  if ($x_0 instanceof \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_ContainsNegativeValue) {
$__t3 = null;;
if ($y_1 instanceof \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_ContainsNegativeValue) {
$__t3 = (($GLOBALS['Data_Interval_Duration_ordDurationComponent'])['compare'])(($x_0)->{'value0'}, ($y_1)->{'value0'});
goto end_branch_3;;
};
$__t3 = new \Data\Ordering\Data_Ordering_LT();
end_branch_3:;
$__t0 = $__t3;
goto end_branch_0;;
};
  if ($y_1 instanceof \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_ContainsNegativeValue) {
$__t0 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_0;;
};
  if (($x_0 instanceof \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_InvalidFractionalUse && $y_1 instanceof \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_InvalidFractionalUse)) {
$__t0 = (($GLOBALS['Data_Interval_Duration_ordDurationComponent'])['compare'])(($x_0)->{'value0'}, ($y_1)->{'value0'});
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Interval_Duration_Iso_eqError'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Interval_Duration_Iso_checkWeekUsage
function majData_majInterval_majDuration_majIso_checkmajWeekmajUsage($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majInterval_majDuration_majIso_checkmajWeekmajUsage';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__t0 = null;;
  if ((function() use ($v_0, &$__fn) {
$__local_var_1_1 = ($GLOBALS['Data_Interval_Duration_Iso_lookup'])(new \Data\Interval\Duration\Data_Interval_Duration_Week(), ($v_0)['asMap']);
$__t2 = null;;
if ($__local_var_1_1 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t2 = false;
goto end_branch_2;;
};
if ($__local_var_1_1 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t2 = true;
goto end_branch_2;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
end_branch_2:;
$__t3 = null;;
if (($v_0)['asMap'] instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t3 = 0;
goto end_branch_3;;
};
if (($v_0)['asMap'] instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t3 = (($v_0)['asMap'])->{'value1'};
goto end_branch_3;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
end_branch_3:;
return (($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])['conj'])($__t2, ($GLOBALS['Data_Interval_Duration_Iso_greaterThan'])($__t3, 1));
})()) {
$__t0 = (($GLOBALS['Data_List_Types_applicativeList'])['pure'])(new \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_InvalidWeekComponentUsage());
goto end_branch_0;;
};
  $__t0 = ($GLOBALS['Data_List_Types_plusList'])['empty'];
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Interval_Duration_Iso_checkWeekUsage'] = __NAMESPACE__ . '\\majData_majInterval_majDuration_majIso_checkmajWeekmajUsage';

// Data_Interval_Duration_Iso_checkNegativeValues
function majData_majInterval_majDuration_majIso_checkmajNegativemajValues($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majInterval_majDuration_majIso_checkmajNegativemajValues';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Interval_Duration_Iso_foldMap1'])(function($v1_1 = null) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((($v1_1)->{'value1'} >= 0.0)) {
$__t0 = ($GLOBALS['Data_List_Types_plusList'])['empty'];
goto end_branch_0;;
};
  $__t0 = (($GLOBALS['Data_List_Types_applicativeList'])['pure'])(new \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_ContainsNegativeValue(($v1_1)->{'value0'}));
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($v_0)['asList']);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Interval_Duration_Iso_checkNegativeValues'] = __NAMESPACE__ . '\\majData_majInterval_majDuration_majIso_checkmajNegativemajValues';

// Data_Interval_Duration_Iso_checkFractionalUse
function majData_majInterval_majDuration_majIso_checkmajFractionalmajUse($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majInterval_majDuration_majIso_checkmajFractionalmajUse';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = ($GLOBALS['Data_List_span'])(($GLOBALS['Control_Semigroupoid_composeImpl'])(function($a_1 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])['not'])((($GLOBALS['Data_Eq_eqBoolean'])['eq'])((($GLOBALS['Data_Eq_eqNumber'])['eq'])(($GLOBALS['Data_Number_floor'])($a_1), $a_1), false));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $GLOBALS['Data_Tuple_snd']), ($v_0)['asList']);
  $__t1 = null;;
  if ((($__local_var_1_0)['rest'] instanceof \Data\List\Types\Data_List_Types_Cons && (($GLOBALS['Data_Interval_Duration_Iso_foldMap2'])(($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Monoid_Additive_Additive'], $GLOBALS['Data_Number_abs']), $GLOBALS['Data_Tuple_snd']), (($__local_var_1_0)['rest'])->{'value1'}) > 0.0))) {
$__t1 = (($GLOBALS['Data_List_Types_applicativeList'])['pure'])(new \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_InvalidFractionalUse(((($__local_var_1_0)['rest'])->{'value0'})->{'value0'}));
goto end_branch_1;;
};
  $__t1 = ($GLOBALS['Data_List_Types_plusList'])['empty'];
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Interval_Duration_Iso_checkFractionalUse'] = __NAMESPACE__ . '\\majData_majInterval_majDuration_majIso_checkmajFractionalmajUse';

// Data_Interval_Duration_Iso_checkEmptiness
function majData_majInterval_majDuration_majIso_checkmajEmptiness($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majInterval_majDuration_majIso_checkmajEmptiness';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__t0 = null;;
  if (($v_0)['asList'] instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = (($GLOBALS['Data_List_Types_applicativeList'])['pure'])(new \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_IsEmpty());
goto end_branch_0;;
};
  $__t0 = ($GLOBALS['Data_List_Types_plusList'])['empty'];
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Interval_Duration_Iso_checkEmptiness'] = __NAMESPACE__ . '\\majData_majInterval_majDuration_majIso_checkmajEmptiness';

// Data_Interval_Duration_Iso_checkValidIsoDuration
function majData_majInterval_majDuration_majIso_checkmajValidmajIsomajDuration($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majInterval_majDuration_majIso_checkmajValidmajIsomajDuration';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $go__1_0 = null;
  $go__1_0 = (function() use (&$go__1_0) {
  $__fn = function($v_2 = null, $v1_3 = null) use (&$go__1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__1_0_0_v_2 = $v_2;
  $__tco_var_go__1_0_0_v1_3 = $v1_3;
  tco_loop_go__1_0_0:;
  $v_2 = $__tco_var_go__1_0_0_v_2;
  $v1_3 = $__tco_var_go__1_0_0_v1_3;
  $__t0 = null;;
  if ($v1_3 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = $v_2;
goto end_branch_0;;
};
  if ($v1_3 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_1 = new \Data\List\Types\Data_List_Types_Cons(($v1_3)->{'value0'}, $v_2);
$__tco_2 = ($v1_3)->{'value1'};
$__tco_var_go__1_0_0_v_2 = $__tco_1;
$__tco_var_go__1_0_0_v1_3 = $__tco_2;
goto tco_loop_go__1_0_0;;
$__t0 = null;
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ($GLOBALS['Data_Interval_Duration_Iso_fold'])([$GLOBALS['Data_Interval_Duration_Iso_checkWeekUsage'], $GLOBALS['Data_Interval_Duration_Iso_checkEmptiness'], $GLOBALS['Data_Interval_Duration_Iso_checkFractionalUse'], $GLOBALS['Data_Interval_Duration_Iso_checkNegativeValues']], ["asList" => ($go__1_0)(new \Data\List\Types\Data_List_Types_Nil(), ($GLOBALS['Data_Interval_Duration_Iso_toUnfoldable'])($v_0)), "asMap" => $v_0]);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Interval_Duration_Iso_checkValidIsoDuration'] = __NAMESPACE__ . '\\majData_majInterval_majDuration_majIso_checkmajValidmajIsomajDuration';

// Data_Interval_Duration_Iso_mkIsoDuration
function majData_majInterval_majDuration_majIso_mkmajIsomajDuration($d_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majInterval_majDuration_majIso_mkmajIsomajDuration';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = ($GLOBALS['Data_Interval_Duration_Iso_checkValidIsoDuration'])($d_0);
  $__t1 = null;;
  if ($__local_var_1_0 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = new \Data\Either\Data_Either_Right($d_0);
goto end_branch_1;;
};
  if ($__local_var_1_0 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__t1 = new \Data\Either\Data_Either_Left(new \Data\NonEmpty\Data_NonEmpty_NonEmpty(($__local_var_1_0)->{'value0'}, ($__local_var_1_0)->{'value1'}));
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Interval_Duration_Iso_mkIsoDuration'] = __NAMESPACE__ . '\\majData_majInterval_majDuration_majIso_mkmajIsomajDuration';

