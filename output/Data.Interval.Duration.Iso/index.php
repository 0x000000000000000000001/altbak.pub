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
if (!\function_exists(__NAMESPACE__ . '\\phpurs_execute_effect')) {
  function phpurs_execute_effect($val) {
    if (\is_callable($val)) {
      return $val($GLOBALS['Data_Unit_unit']);
    }
    return $val;
  }
}

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };


final class Data_Interval_Duration_Iso_IsEmpty { public $tag = 'IsEmpty'; public function __construct() {} }
final class Data_Interval_Duration_Iso_InvalidWeekComponentUsage { public $tag = 'InvalidWeekComponentUsage'; public function __construct() {} }
final class Data_Interval_Duration_Iso_ContainsNegativeValue { public $tag = 'ContainsNegativeValue'; public function __construct(public  $value0) {} }
final class Data_Interval_Duration_Iso_InvalidFractionalUse { public $tag = 'InvalidFractionalUse'; public function __construct(public  $value0) {} }

// Data_Interval_Duration_Iso_IsEmpty
$GLOBALS['Data_Interval_Duration_Iso_IsEmpty'] = ($GLOBALS['__phpurs_data0_IsEmpty'] ??= new \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_IsEmpty());

// Data_Interval_Duration_Iso_InvalidWeekComponentUsage
$GLOBALS['Data_Interval_Duration_Iso_InvalidWeekComponentUsage'] = ($GLOBALS['__phpurs_data0_InvalidWeekComponentUsage'] ??= new \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_InvalidWeekComponentUsage());

// Data_Interval_Duration_Iso_ContainsNegativeValue
$GLOBALS['Data_Interval_Duration_Iso_ContainsNegativeValue'] = function($value0) {
  $__num = \func_num_args();
  $__res = new \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_ContainsNegativeValue($value0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Interval_Duration_Iso_InvalidFractionalUse
$GLOBALS['Data_Interval_Duration_Iso_InvalidFractionalUse'] = function($value0) {
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
$GLOBALS['Data_Interval_Duration_Iso_showIsoDuration'] = (object)["show" => function($v_0) {
  $__num = \func_num_args();
  $__res = (("(IsoDuration (Duration (fromFoldable " . \Data\Show\majData_majShow_showmajArraymajImpl(function($v_1) {
  $__num = \func_num_args();
  $__t0 = null;;
  if (($v_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Minute) {
$__t0 = "(Tuple Minute ";
goto end_branch_0;;
};
  if (($v_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Second) {
$__t0 = "(Tuple Second ";
goto end_branch_0;;
};
  if (($v_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Hour) {
$__t0 = "(Tuple Hour ";
goto end_branch_0;;
};
  if (($v_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Day) {
$__t0 = "(Tuple Day ";
goto end_branch_0;;
};
  if (($v_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Week) {
$__t0 = "(Tuple Week ";
goto end_branch_0;;
};
  if (($v_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Month) {
$__t0 = "(Tuple Month ";
goto end_branch_0;;
};
  if (($v_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Year) {
$__t0 = "(Tuple Year ";
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = (($__t0 . \Data\Show\majData_majShow_showmajNumbermajImpl(($v_1)->{'value1'})) . ")");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(((((($GLOBALS['Data_Unfoldable_unfoldrArrayImpl'])($GLOBALS['Data_Maybe_isNothing']))(function($v_1) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v_1 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = ($v_1)->{'value0'};
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Tuple_fst']))($GLOBALS['Data_Tuple_snd']))($GLOBALS['Data_Map_Internal_stepUnfoldr']), $GLOBALS['Data_Map_Internal_toMapIter'], $v_0))) . ")))");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Interval_Duration_Iso_showError
$GLOBALS['Data_Interval_Duration_Iso_showError'] = (object)["show" => function($v_0) {
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
$__t1 = null;;
if (($v_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Minute) {
$__t1 = "(ContainsNegativeValue Minute)";
goto end_branch_1;;
};
if (($v_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Second) {
$__t1 = "(ContainsNegativeValue Second)";
goto end_branch_1;;
};
if (($v_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Hour) {
$__t1 = "(ContainsNegativeValue Hour)";
goto end_branch_1;;
};
if (($v_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Day) {
$__t1 = "(ContainsNegativeValue Day)";
goto end_branch_1;;
};
if (($v_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Week) {
$__t1 = "(ContainsNegativeValue Week)";
goto end_branch_1;;
};
if (($v_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Month) {
$__t1 = "(ContainsNegativeValue Month)";
goto end_branch_1;;
};
if (($v_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Year) {
$__t1 = "(ContainsNegativeValue Year)";
goto end_branch_1;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
end_branch_1:;
$__t0 = $__t1;
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_InvalidFractionalUse) {
$__t2 = null;;
if (($v_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Minute) {
$__t2 = "(InvalidFractionalUse Minute)";
goto end_branch_2;;
};
if (($v_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Second) {
$__t2 = "(InvalidFractionalUse Second)";
goto end_branch_2;;
};
if (($v_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Hour) {
$__t2 = "(InvalidFractionalUse Hour)";
goto end_branch_2;;
};
if (($v_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Day) {
$__t2 = "(InvalidFractionalUse Day)";
goto end_branch_2;;
};
if (($v_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Week) {
$__t2 = "(InvalidFractionalUse Week)";
goto end_branch_2;;
};
if (($v_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Month) {
$__t2 = "(InvalidFractionalUse Month)";
goto end_branch_2;;
};
if (($v_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Year) {
$__t2 = "(InvalidFractionalUse Year)";
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
}];

// Data_Interval_Duration_Iso_prettyError
function majData_majInterval_majDuration_majIso_prettymajError($v_0): string|\Closure {
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
$__t1 = null;;
if (($v_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Minute) {
$__t1 = "Component `Minute` contains negative value";
goto end_branch_1;;
};
if (($v_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Second) {
$__t1 = "Component `Second` contains negative value";
goto end_branch_1;;
};
if (($v_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Hour) {
$__t1 = "Component `Hour` contains negative value";
goto end_branch_1;;
};
if (($v_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Day) {
$__t1 = "Component `Day` contains negative value";
goto end_branch_1;;
};
if (($v_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Week) {
$__t1 = "Component `Week` contains negative value";
goto end_branch_1;;
};
if (($v_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Month) {
$__t1 = "Component `Month` contains negative value";
goto end_branch_1;;
};
if (($v_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Year) {
$__t1 = "Component `Year` contains negative value";
goto end_branch_1;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
end_branch_1:;
$__t0 = $__t1;
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_InvalidFractionalUse) {
$__t2 = null;;
if (($v_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Minute) {
$__t2 = "Invalid usage of Fractional value at component `Minute`";
goto end_branch_2;;
};
if (($v_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Second) {
$__t2 = "Invalid usage of Fractional value at component `Second`";
goto end_branch_2;;
};
if (($v_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Hour) {
$__t2 = "Invalid usage of Fractional value at component `Hour`";
goto end_branch_2;;
};
if (($v_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Day) {
$__t2 = "Invalid usage of Fractional value at component `Day`";
goto end_branch_2;;
};
if (($v_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Week) {
$__t2 = "Invalid usage of Fractional value at component `Week`";
goto end_branch_2;;
};
if (($v_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Month) {
$__t2 = "Invalid usage of Fractional value at component `Month`";
goto end_branch_2;;
};
if (($v_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Year) {
$__t2 = "Invalid usage of Fractional value at component `Year`";
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Interval_Duration_Iso_prettyError'] = __NAMESPACE__ . '\\majData_majInterval_majDuration_majIso_prettymajError';

// Data_Interval_Duration_Iso_eqIsoDuration
$GLOBALS['Data_Interval_Duration_Iso_eqIsoDuration'] = (object)["eq" => function($x_0) {
  $__num = \func_num_args();
  $__res = function($y_1) use ($x_0) {
  $__num = \func_num_args();
  $go__go_2_0 = null;
  $go__go_2_0 = function($a_3) use (&$go__go_2_0) {
  $__num = \func_num_args();
  $__res = function($b_4) use ($a_3, &$go__go_2_0) {
  $__num = \func_num_args();
  $v_5_1 = \Data\Map\Internal\majData_majMap_majInternal_stepmajAsc($a_3);
  $__t2 = null;;
  if ($v_5_1 instanceof \Data\Map\Internal\Data_Map_Internal_IterNext) {
$v2_6_3 = \Data\Map\Internal\majData_majMap_majInternal_stepmajAsc($b_4);
$__t4 = null;;
if (($v_5_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Second) {
$__t4 = ($v2_6_3)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Second;
goto end_branch_4;;
};
if (($v_5_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Minute) {
$__t4 = ($v2_6_3)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Minute;
goto end_branch_4;;
};
if (($v_5_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Hour) {
$__t4 = ($v2_6_3)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Hour;
goto end_branch_4;;
};
if (($v_5_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Day) {
$__t4 = ($v2_6_3)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Day;
goto end_branch_4;;
};
if (($v_5_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Week) {
$__t4 = ($v2_6_3)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Week;
goto end_branch_4;;
};
if (($v_5_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Month) {
$__t4 = ($v2_6_3)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Month;
goto end_branch_4;;
};
$__t4 = (($v_5_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Year && ($v2_6_3)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Year);
end_branch_4:;
$__t2 = ($v2_6_3 instanceof \Data\Map\Internal\Data_Map_Internal_IterNext && (($__t4 && (($v_5_1)->{'value1'} === ($v2_6_3)->{'value1'})) && (($go__go_2_0)(($v_5_1)->{'value2'}))(($v2_6_3)->{'value2'})));
goto end_branch_2;;
};
  if ($v_5_1 instanceof \Data\Map\Internal\Data_Map_Internal_IterDone) {
$__t2 = true;
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
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
  $eqMapIter2_2_0 = (object)["eq" => $go__go_2_0];
  $__t6 = null;;
  if ($x_0 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t6 = $y_1 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf;
goto end_branch_6;;
};
  if ($x_0 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t6 = ($y_1 instanceof \Data\Map\Internal\Data_Map_Internal_Node && ((($x_0)->{'value1'} === ($y_1)->{'value1'}) && ((($eqMapIter2_2_0)->{'eq'})(new \Data\Map\Internal\Data_Map_Internal_IterNode($x_0, new \Data\Map\Internal\Data_Map_Internal_IterLeaf())))(new \Data\Map\Internal\Data_Map_Internal_IterNode($y_1, new \Data\Map\Internal\Data_Map_Internal_IterLeaf()))));
goto end_branch_6;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t6 = null;
  end_branch_6:;
  $__res = $__t6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Interval_Duration_Iso_ordIsoDuration
$GLOBALS['Data_Interval_Duration_Iso_ordIsoDuration'] = (object)["compare" => function($x_0) {
  $__num = \func_num_args();
  $__res = function($y_1) use ($x_0) {
  $__num = \func_num_args();
  $go__go_2_0 = null;
  $go__go_2_0 = function($a_3) use (&$go__go_2_0) {
  $__num = \func_num_args();
  $__res = function($b_4) use ($a_3, &$go__go_2_0) {
  $__num = \func_num_args();
  $v_5_1 = \Data\Map\Internal\majData_majMap_majInternal_stepmajAsc($a_3);
  $__t2 = null;;
  if ($v_5_1 instanceof \Data\Map\Internal\Data_Map_Internal_IterNext) {
$v2_6_3 = \Data\Map\Internal\majData_majMap_majInternal_stepmajAsc($b_4);
$__t4 = null;;
if (($v_5_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Second) {
$__t4 = ($v2_6_3)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Second;
goto end_branch_4;;
};
if (($v_5_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Minute) {
$__t4 = ($v2_6_3)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Minute;
goto end_branch_4;;
};
if (($v_5_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Hour) {
$__t4 = ($v2_6_3)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Hour;
goto end_branch_4;;
};
if (($v_5_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Day) {
$__t4 = ($v2_6_3)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Day;
goto end_branch_4;;
};
if (($v_5_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Week) {
$__t4 = ($v2_6_3)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Week;
goto end_branch_4;;
};
if (($v_5_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Month) {
$__t4 = ($v2_6_3)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Month;
goto end_branch_4;;
};
$__t4 = (($v_5_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Year && ($v2_6_3)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Year);
end_branch_4:;
$__t2 = ($v2_6_3 instanceof \Data\Map\Internal\Data_Map_Internal_IterNext && (($__t4 && (($v_5_1)->{'value1'} === ($v2_6_3)->{'value1'})) && (($go__go_2_0)(($v_5_1)->{'value2'}))(($v2_6_3)->{'value2'})));
goto end_branch_2;;
};
  if ($v_5_1 instanceof \Data\Map\Internal\Data_Map_Internal_IterDone) {
$__t2 = true;
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
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
  $eqMapIter2_2_0 = (object)["eq" => $go__go_2_0];
  $go__go_3_6 = null;
  $go__go_3_6 = (function() use (&$go__go_3_6) {
  $__fn = function($a_4, $b_5 = null) use (&$go__go_3_6, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_3_6_6_a_4 = $a_4;
  $__tco_var_go__go_3_6_6_b_5 = $b_5;
  tco_loop_go__go_3_6_6:;
  $a_4 = $__tco_var_go__go_3_6_6_a_4;
  $b_5 = $__tco_var_go__go_3_6_6_b_5;
  $v_6_6 = \Data\Map\Internal\majData_majMap_majInternal_stepmajAsc($b_5);
  $v1_7_7 = \Data\Map\Internal\majData_majMap_majInternal_stepmajAsc($a_4);
  $__t8 = null;;
  if ($v1_7_7 instanceof \Data\Map\Internal\Data_Map_Internal_IterNext) {
$__t9 = null;;
if ($v_6_6 instanceof \Data\Map\Internal\Data_Map_Internal_IterNext) {
$__t10 = null;;
if (($v1_7_7)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Second) {
$__t11 = null;;
if (($v_6_6)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Second) {
$__t11 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_11;;
};
$__t11 = new \Data\Ordering\Data_Ordering_LT();
end_branch_11:;
$__t10 = $__t11;
goto end_branch_10;;
};
if (($v_6_6)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Second) {
$__t10 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_10;;
};
if (($v1_7_7)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Minute) {
$__t12 = null;;
if (($v_6_6)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Minute) {
$__t12 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_12;;
};
$__t12 = new \Data\Ordering\Data_Ordering_LT();
end_branch_12:;
$__t10 = $__t12;
goto end_branch_10;;
};
if (($v_6_6)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Minute) {
$__t10 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_10;;
};
if (($v1_7_7)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Hour) {
$__t13 = null;;
if (($v_6_6)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Hour) {
$__t13 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_13;;
};
$__t13 = new \Data\Ordering\Data_Ordering_LT();
end_branch_13:;
$__t10 = $__t13;
goto end_branch_10;;
};
if (($v_6_6)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Hour) {
$__t10 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_10;;
};
if (($v1_7_7)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Day) {
$__t14 = null;;
if (($v_6_6)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Day) {
$__t14 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_14;;
};
$__t14 = new \Data\Ordering\Data_Ordering_LT();
end_branch_14:;
$__t10 = $__t14;
goto end_branch_10;;
};
if (($v_6_6)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Day) {
$__t10 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_10;;
};
if (($v1_7_7)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Week) {
$__t15 = null;;
if (($v_6_6)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Week) {
$__t15 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_15;;
};
$__t15 = new \Data\Ordering\Data_Ordering_LT();
end_branch_15:;
$__t10 = $__t15;
goto end_branch_10;;
};
if (($v_6_6)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Week) {
$__t10 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_10;;
};
if (($v1_7_7)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Month) {
$__t16 = null;;
if (($v_6_6)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Month) {
$__t16 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_16;;
};
$__t16 = new \Data\Ordering\Data_Ordering_LT();
end_branch_16:;
$__t10 = $__t16;
goto end_branch_10;;
};
if (($v_6_6)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Month) {
$__t10 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_10;;
};
if ((($v1_7_7)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Year && ($v_6_6)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Year)) {
$__t10 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_10;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t10 = null;
end_branch_10:;
$v3_8_10 = $__t10;
$__t18 = null;;
if ($v3_8_10 instanceof \Data\Ordering\Data_Ordering_EQ) {
$v4_9_19 = \Data\Ord\majData_majOrd_ordmajNumbermajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), ($v1_7_7)->{'value1'}, ($v_6_6)->{'value1'});
$__t20 = null;;
if ($v4_9_19 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__tco_21 = ($v1_7_7)->{'value2'};
$__tco_22 = ($v_6_6)->{'value2'};
$__tco_var_go__go_3_6_6_a_4 = $__tco_21;
$__tco_var_go__go_3_6_6_b_5 = $__tco_22;
goto tco_loop_go__go_3_6_6;;
$__t20 = null;
goto end_branch_20;;
};
$__t20 = $v4_9_19;
end_branch_20:;
$__t18 = $__t20;
goto end_branch_18;;
};
$__t18 = $v3_8_10;
end_branch_18:;
$__t9 = $__t18;
goto end_branch_9;;
};
if ($v_6_6 instanceof \Data\Map\Internal\Data_Map_Internal_IterDone) {
$__t9 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_9;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t9 = null;
end_branch_9:;
$__t8 = $__t9;
goto end_branch_8;;
};
  if ($v1_7_7 instanceof \Data\Map\Internal\Data_Map_Internal_IterDone) {
$__t23 = null;;
if ($v_6_6 instanceof \Data\Map\Internal\Data_Map_Internal_IterDone) {
$__t23 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_23;;
};
$__t23 = new \Data\Ordering\Data_Ordering_LT();
end_branch_23:;
$__t8 = $__t23;
goto end_branch_8;;
};
  if ($v_6_6 instanceof \Data\Map\Internal\Data_Map_Internal_IterDone) {
$__t8 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_8;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t8 = null;
  end_branch_8:;
  $__res = $__t8;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $ordMapIter2_2_0 = (object)["compare" => $go__go_3_6, "Eq0" => function($_dollar___unused_3) use ($eqMapIter2_2_0) {
  $__num = \func_num_args();
  $__res = $eqMapIter2_2_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__t9 = null;;
  if ($x_0 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t10 = null;;
if ($y_1 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t10 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_10;;
};
$__t10 = new \Data\Ordering\Data_Ordering_LT();
end_branch_10:;
$__t9 = $__t10;
goto end_branch_9;;
};
  $__t8 = null;;
  if ($y_1 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t8 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_8;;
};
  $__t8 = ((($ordMapIter2_2_0)->{'compare'})(new \Data\Map\Internal\Data_Map_Internal_IterNode($x_0, new \Data\Map\Internal\Data_Map_Internal_IterLeaf())))(new \Data\Map\Internal\Data_Map_Internal_IterNode($y_1, new \Data\Map\Internal\Data_Map_Internal_IterLeaf()));
  end_branch_8:;
  $__t9 = $__t8;
  end_branch_9:;
  $__res = $__t9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Interval_Duration_Iso_eqIsoDuration'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Interval_Duration_Iso_eqError
$GLOBALS['Data_Interval_Duration_Iso_eqError'] = (object)["eq" => function($x_0) {
  $__num = \func_num_args();
  $__res = function($y_1) use ($x_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($x_0 instanceof \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_IsEmpty) {
$__t1 = $y_1 instanceof \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_IsEmpty;
goto end_branch_1;;
};
  if ($x_0 instanceof \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_InvalidWeekComponentUsage) {
$__t1 = $y_1 instanceof \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_InvalidWeekComponentUsage;
goto end_branch_1;;
};
  if ($x_0 instanceof \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_ContainsNegativeValue) {
$__t2 = null;;
if (($x_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Second) {
$__t2 = ($y_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Second;
goto end_branch_2;;
};
if (($x_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Minute) {
$__t2 = ($y_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Minute;
goto end_branch_2;;
};
if (($x_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Hour) {
$__t2 = ($y_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Hour;
goto end_branch_2;;
};
if (($x_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Day) {
$__t2 = ($y_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Day;
goto end_branch_2;;
};
if (($x_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Week) {
$__t2 = ($y_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Week;
goto end_branch_2;;
};
if (($x_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Month) {
$__t2 = ($y_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Month;
goto end_branch_2;;
};
$__t2 = (($x_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Year && ($y_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Year);
end_branch_2:;
$__t1 = ($y_1 instanceof \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_ContainsNegativeValue && $__t2);
goto end_branch_1;;
};
  $__t0 = null;;
  if (($x_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Second) {
$__t0 = ($y_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Second;
goto end_branch_0;;
};
  if (($x_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Minute) {
$__t0 = ($y_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Minute;
goto end_branch_0;;
};
  if (($x_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Hour) {
$__t0 = ($y_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Hour;
goto end_branch_0;;
};
  if (($x_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Day) {
$__t0 = ($y_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Day;
goto end_branch_0;;
};
  if (($x_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Week) {
$__t0 = ($y_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Week;
goto end_branch_0;;
};
  if (($x_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Month) {
$__t0 = ($y_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Month;
goto end_branch_0;;
};
  $__t0 = (($x_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Year && ($y_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Year);
  end_branch_0:;
  $__t1 = ($x_0 instanceof \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_InvalidFractionalUse && ($y_1 instanceof \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_InvalidFractionalUse && $__t0));
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Interval_Duration_Iso_ordError
$GLOBALS['Data_Interval_Duration_Iso_ordError'] = (object)["compare" => function($x_0) {
  $__num = \func_num_args();
  $__res = function($y_1) use ($x_0) {
  $__num = \func_num_args();
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
$__t4 = null;;
if (($x_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Second) {
$__t5 = null;;
if (($y_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Second) {
$__t5 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_5;;
};
$__t5 = new \Data\Ordering\Data_Ordering_LT();
end_branch_5:;
$__t4 = $__t5;
goto end_branch_4;;
};
if (($y_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Second) {
$__t4 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_4;;
};
if (($x_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Minute) {
$__t6 = null;;
if (($y_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Minute) {
$__t6 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_6;;
};
$__t6 = new \Data\Ordering\Data_Ordering_LT();
end_branch_6:;
$__t4 = $__t6;
goto end_branch_4;;
};
if (($y_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Minute) {
$__t4 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_4;;
};
if (($x_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Hour) {
$__t7 = null;;
if (($y_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Hour) {
$__t7 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_7;;
};
$__t7 = new \Data\Ordering\Data_Ordering_LT();
end_branch_7:;
$__t4 = $__t7;
goto end_branch_4;;
};
if (($y_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Hour) {
$__t4 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_4;;
};
if (($x_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Day) {
$__t8 = null;;
if (($y_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Day) {
$__t8 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_8;;
};
$__t8 = new \Data\Ordering\Data_Ordering_LT();
end_branch_8:;
$__t4 = $__t8;
goto end_branch_4;;
};
if (($y_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Day) {
$__t4 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_4;;
};
if (($x_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Week) {
$__t9 = null;;
if (($y_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Week) {
$__t9 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_9;;
};
$__t9 = new \Data\Ordering\Data_Ordering_LT();
end_branch_9:;
$__t4 = $__t9;
goto end_branch_4;;
};
if (($y_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Week) {
$__t4 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_4;;
};
if (($x_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Month) {
$__t10 = null;;
if (($y_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Month) {
$__t10 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_10;;
};
$__t10 = new \Data\Ordering\Data_Ordering_LT();
end_branch_10:;
$__t4 = $__t10;
goto end_branch_4;;
};
if (($y_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Month) {
$__t4 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_4;;
};
if ((($x_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Year && ($y_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Year)) {
$__t4 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_4;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t4 = null;
end_branch_4:;
$__t3 = $__t4;
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
$__t11 = null;;
if (($x_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Second) {
$__t12 = null;;
if (($y_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Second) {
$__t12 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_12;;
};
$__t12 = new \Data\Ordering\Data_Ordering_LT();
end_branch_12:;
$__t11 = $__t12;
goto end_branch_11;;
};
if (($y_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Second) {
$__t11 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_11;;
};
if (($x_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Minute) {
$__t13 = null;;
if (($y_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Minute) {
$__t13 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_13;;
};
$__t13 = new \Data\Ordering\Data_Ordering_LT();
end_branch_13:;
$__t11 = $__t13;
goto end_branch_11;;
};
if (($y_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Minute) {
$__t11 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_11;;
};
if (($x_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Hour) {
$__t14 = null;;
if (($y_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Hour) {
$__t14 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_14;;
};
$__t14 = new \Data\Ordering\Data_Ordering_LT();
end_branch_14:;
$__t11 = $__t14;
goto end_branch_11;;
};
if (($y_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Hour) {
$__t11 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_11;;
};
if (($x_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Day) {
$__t15 = null;;
if (($y_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Day) {
$__t15 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_15;;
};
$__t15 = new \Data\Ordering\Data_Ordering_LT();
end_branch_15:;
$__t11 = $__t15;
goto end_branch_11;;
};
if (($y_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Day) {
$__t11 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_11;;
};
if (($x_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Week) {
$__t16 = null;;
if (($y_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Week) {
$__t16 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_16;;
};
$__t16 = new \Data\Ordering\Data_Ordering_LT();
end_branch_16:;
$__t11 = $__t16;
goto end_branch_11;;
};
if (($y_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Week) {
$__t11 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_11;;
};
if (($x_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Month) {
$__t17 = null;;
if (($y_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Month) {
$__t17 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_17;;
};
$__t17 = new \Data\Ordering\Data_Ordering_LT();
end_branch_17:;
$__t11 = $__t17;
goto end_branch_11;;
};
if (($y_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Month) {
$__t11 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_11;;
};
if ((($x_0)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Year && ($y_1)->{'value0'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Year)) {
$__t11 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_11;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t11 = null;
end_branch_11:;
$__t0 = $__t11;
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq0" => function($_dollar___unused_0) {
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
$go__go_1_1 = null;
$go__go_1_1 = function($v_2) use (&$go__go_1_1) {
  $__num = \func_num_args();
  $__tco_var_go__go_1_1_1_v_2 = $v_2;
  tco_loop_go__go_1_1_1:;
  $v_2 = $__tco_var_go__go_1_1_1_v_2;
  $__t1 = null;;
  if ($v_2 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t1 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_1;;
};
  if ($v_2 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t3 = null;;
if (($v_2)->{'value2'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Second) {
$__tco_4 = ($v_2)->{'value5'};
$__tco_var_go__go_1_1_1_v_2 = $__tco_4;
goto tco_loop_go__go_1_1_1;;
$__t3 = null;
goto end_branch_3;;
};
if (($v_2)->{'value2'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Minute) {
$__tco_5 = ($v_2)->{'value5'};
$__tco_var_go__go_1_1_1_v_2 = $__tco_5;
goto tco_loop_go__go_1_1_1;;
$__t3 = null;
goto end_branch_3;;
};
if (($v_2)->{'value2'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Hour) {
$__tco_6 = ($v_2)->{'value5'};
$__tco_var_go__go_1_1_1_v_2 = $__tco_6;
goto tco_loop_go__go_1_1_1;;
$__t3 = null;
goto end_branch_3;;
};
if (($v_2)->{'value2'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Day) {
$__tco_7 = ($v_2)->{'value5'};
$__tco_var_go__go_1_1_1_v_2 = $__tco_7;
goto tco_loop_go__go_1_1_1;;
$__t3 = null;
goto end_branch_3;;
};
if (($v_2)->{'value2'} instanceof \Data\Interval\Duration\Data_Interval_Duration_Week) {
$__t3 = new \Data\Maybe\Data_Maybe_Just(($v_2)->{'value3'});
goto end_branch_3;;
};
$__tco_2 = ($v_2)->{'value4'};
$__tco_var_go__go_1_1_1_v_2 = $__tco_2;
goto tco_loop_go__go_1_1_1;;
$__t3 = null;
end_branch_3:;
$__t1 = $__t3;
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
$__local_var_1_1 = ($go__go_1_1)(($v_0)->{'asMap'});
$__t3 = null;;
if ($__local_var_1_1 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t3 = false;
goto end_branch_3;;
};
if ($__local_var_1_1 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t3 = true;
goto end_branch_3;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
end_branch_3:;
$__t4 = null;;
if (($v_0)->{'asMap'} instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t4 = false;
goto end_branch_4;;
};
if (($v_0)->{'asMap'} instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t4 = ((($v_0)->{'asMap'})->{'value1'} > 1);
goto end_branch_4;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t4 = null;
end_branch_4:;
return ($__t3 && $__t4);
})()) {
$__t0 = new \Data\List\Types\Data_List_Types_Cons(new \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_InvalidWeekComponentUsage(), new \Data\List\Types\Data_List_Types_Nil());
goto end_branch_0;;
};
  $__t0 = new \Data\List\Types\Data_List_Types_Nil();
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
  $go__go_1_0 = null;
  $go__go_1_0 = (function() use (&$go__go_1_0) {
  $__fn = function($b_2, $v_3 = null) use (&$go__go_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_1_0_0_b_2 = $b_2;
  $__tco_var_go__go_1_0_0_v_3 = $v_3;
  tco_loop_go__go_1_0_0:;
  $b_2 = $__tco_var_go__go_1_0_0_b_2;
  $v_3 = $__tco_var_go__go_1_0_0_v_3;
  $__t0 = null;;
  if ($v_3 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = $b_2;
goto end_branch_0;;
};
  if ($v_3 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_4 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($ys_4) use ($b_2) {
  $__num = \func_num_args();
  $go__go_5_1 = null;
  $go__go_5_1 = (function() use (&$go__go_5_1) {
  $__fn = function($b_6, $v_7 = null) use (&$go__go_5_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_5_1_1_b_6 = $b_6;
  $__tco_var_go__go_5_1_1_v_7 = $v_7;
  tco_loop_go__go_5_1_1:;
  $b_6 = $__tco_var_go__go_5_1_1_b_6;
  $v_7 = $__tco_var_go__go_5_1_1_v_7;
  $__t1 = null;;
  if ($v_7 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = $b_6;
goto end_branch_1;;
};
  if ($v_7 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_2 = new \Data\List\Types\Data_List_Types_Cons(($v_7)->{'value0'}, $b_6);
$__tco_3 = ($v_7)->{'value1'};
$__tco_var_go__go_5_1_1_b_6 = $__tco_2;
$__tco_var_go__go_5_1_1_v_7 = $__tco_3;
goto tco_loop_go__go_5_1_1;;
$__t1 = null;
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
  $go__go_5_2 = null;
  $go__go_5_2 = (function() use (&$go__go_5_2) {
  $__fn = function($v_6, $v1_7 = null) use (&$go__go_5_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_5_2_2_v_6 = $v_6;
  $__tco_var_go__go_5_2_2_v1_7 = $v1_7;
  tco_loop_go__go_5_2_2:;
  $v_6 = $__tco_var_go__go_5_2_2_v_6;
  $v1_7 = $__tco_var_go__go_5_2_2_v1_7;
  $__t2 = null;;
  if ($v1_7 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t2 = $v_6;
goto end_branch_2;;
};
  if ($v1_7 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_3 = new \Data\List\Types\Data_List_Types_Cons(($v1_7)->{'value0'}, $v_6);
$__tco_4 = ($v1_7)->{'value1'};
$__tco_var_go__go_5_2_2_v_6 = $__tco_3;
$__tco_var_go__go_5_2_2_v1_7 = $__tco_4;
goto tco_loop_go__go_5_2_2;;
$__t2 = null;
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
  $__res = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(($go__go_5_1)($ys_4), ($go__go_5_2)(new \Data\List\Types\Data_List_Types_Nil()), $b_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, function($v1_4) {
  $__num = \func_num_args();
  $__t3 = null;;
  if (( ! \Data\Ord\majData_majOrd_ordmajNumbermajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), ($v1_4)->{'value1'}, 0.0) instanceof \Data\Ordering\Data_Ordering_LT)) {
$__t3 = new \Data\List\Types\Data_List_Types_Nil();
goto end_branch_3;;
};
  $__t3 = new \Data\List\Types\Data_List_Types_Cons(new \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_ContainsNegativeValue(($v1_4)->{'value0'}), new \Data\List\Types\Data_List_Types_Nil());
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($v_3)->{'value0'});
$__tco_5 = ($v_3)->{'value1'};
$__tco_var_go__go_1_0_0_b_2 = $__tco_4;
$__tco_var_go__go_1_0_0_v_3 = $__tco_5;
goto tco_loop_go__go_1_0_0;;
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
  $__res = (($go__go_1_0)(new \Data\List\Types\Data_List_Types_Nil()))(($v_0)->{'asList'});
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
  $v1_1_0 = (\Data\List\majData_majList_span((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($a_1) {
  $__num = \func_num_args();
  $__res = ( ! ( ! (\Data\Number\majData_majNumber_floor($a_1) === $a_1)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Tuple_snd']), ($v_0)->{'asList'}))->{'rest'};
  $__t1 = null;;
  if ((function() use ($v1_1_0, &$__fn) {
$semigroupAdditive1_2_2 = (object)["append" => function($v_2) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($v_2) {
  $__num = \func_num_args();
  $__res = ($v_2 + $v1_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
$__local_var_2_2 = (object)["mempty" => 0.0, "Semigroup0" => function($_dollar___unused_3) use ($semigroupAdditive1_2_2) {
  $__num = \func_num_args();
  $__res = $semigroupAdditive1_2_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
$Semigroup0_3_4 = (($__local_var_2_2)->{'Semigroup0'})(null);
$__local_var_4_5 = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_4) {
  $__num = \func_num_args();
  $__res = $x_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Number_abs'])))($GLOBALS['Data_Tuple_snd']);
$go__go_5_6 = null;
$go__go_5_6 = (function() use ($Semigroup0_3_4, $__local_var_4_5, &$go__go_5_6) {
  $__fn = function($b_6, $v_7 = null) use ($Semigroup0_3_4, $__local_var_4_5, &$go__go_5_6, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_5_6_6_b_6 = $b_6;
  $__tco_var_go__go_5_6_6_v_7 = $v_7;
  tco_loop_go__go_5_6_6:;
  $b_6 = $__tco_var_go__go_5_6_6_b_6;
  $v_7 = $__tco_var_go__go_5_6_6_v_7;
  $__t6 = null;;
  if ($v_7 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t6 = $b_6;
goto end_branch_6;;
};
  if ($v_7 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_7 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl((($Semigroup0_3_4)->{'append'})($b_6), $__local_var_4_5, ($v_7)->{'value0'});
$__tco_8 = ($v_7)->{'value1'};
$__tco_var_go__go_5_6_6_b_6 = $__tco_7;
$__tco_var_go__go_5_6_6_v_7 = $__tco_8;
goto tco_loop_go__go_5_6_6;;
$__t6 = null;
goto end_branch_6;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t6 = null;
  end_branch_6:;
  $__res = $__t6;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
return ($v1_1_0 instanceof \Data\List\Types\Data_List_Types_Cons && \Data\Ord\majData_majOrd_ordmajNumbermajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), (($go__go_5_6)(($__local_var_2_2)->{'mempty'}))(($v1_1_0)->{'value1'}), 0.0) instanceof \Data\Ordering\Data_Ordering_GT);
})()) {
$__t1 = new \Data\List\Types\Data_List_Types_Cons(new \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_InvalidFractionalUse((($v1_1_0)->{'value0'})->{'value0'}), new \Data\List\Types\Data_List_Types_Nil());
goto end_branch_1;;
};
  $__t1 = new \Data\List\Types\Data_List_Types_Nil();
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
  if (($v_0)->{'asList'} instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = new \Data\List\Types\Data_List_Types_Cons(new \Data\Interval\Duration\Iso\Data_Interval_Duration_Iso_IsEmpty(), new \Data\List\Types\Data_List_Types_Nil());
goto end_branch_0;;
};
  $__t0 = new \Data\List\Types\Data_List_Types_Nil();
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
  $semigroupFn_1_0 = (object)["append" => function($f_1) {
  $__num = \func_num_args();
  $__res = function($g_2) use ($f_1) {
  $__num = \func_num_args();
  $__res = function($x_3) use ($f_1, $g_2) {
  $__num = \func_num_args();
  $go__go_4_0 = null;
  $go__go_4_0 = (function() use (&$go__go_4_0) {
  $__fn = function($b_5, $v_6 = null) use (&$go__go_4_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_4_0_0_b_5 = $b_5;
  $__tco_var_go__go_4_0_0_v_6 = $v_6;
  tco_loop_go__go_4_0_0:;
  $b_5 = $__tco_var_go__go_4_0_0_b_5;
  $v_6 = $__tco_var_go__go_4_0_0_v_6;
  $__t0 = null;;
  if ($v_6 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = $b_5;
goto end_branch_0;;
};
  if ($v_6 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_1 = new \Data\List\Types\Data_List_Types_Cons(($v_6)->{'value0'}, $b_5);
$__tco_2 = ($v_6)->{'value1'};
$__tco_var_go__go_4_0_0_b_5 = $__tco_1;
$__tco_var_go__go_4_0_0_v_6 = $__tco_2;
goto tco_loop_go__go_4_0_0;;
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
  $go__go_4_1 = null;
  $go__go_4_1 = (function() use (&$go__go_4_1) {
  $__fn = function($v_5, $v1_6 = null) use (&$go__go_4_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_4_1_1_v_5 = $v_5;
  $__tco_var_go__go_4_1_1_v1_6 = $v1_6;
  tco_loop_go__go_4_1_1:;
  $v_5 = $__tco_var_go__go_4_1_1_v_5;
  $v1_6 = $__tco_var_go__go_4_1_1_v1_6;
  $__t1 = null;;
  if ($v1_6 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = $v_5;
goto end_branch_1;;
};
  if ($v1_6 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_2 = new \Data\List\Types\Data_List_Types_Cons(($v1_6)->{'value0'}, $v_5);
$__tco_3 = ($v1_6)->{'value1'};
$__tco_var_go__go_4_1_1_v_5 = $__tco_2;
$__tco_var_go__go_4_1_1_v1_6 = $__tco_3;
goto tco_loop_go__go_4_1_1;;
$__t1 = null;
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
  $__res = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(($go__go_4_0)(($g_2)($x_3)), ($go__go_4_1)(new \Data\List\Types\Data_List_Types_Nil()), ($f_1)($x_3));
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
  $__local_var_1_0 = (object)["mempty" => function($v_2) {
  $__num = \func_num_args();
  $__res = new \Data\List\Types\Data_List_Types_Nil();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Semigroup0" => function($_dollar___unused_2) use ($semigroupFn_1_0) {
  $__num = \func_num_args();
  $__res = $semigroupFn_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Semigroup0_2_4 = (($__local_var_1_0)->{'Semigroup0'})(null);
  $go__go_3_5 = null;
  $go__go_3_5 = (function() use (&$go__go_3_5) {
  $__fn = function($v_4, $v1_5 = null) use (&$go__go_3_5, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_3_5_5_v_4 = $v_4;
  $__tco_var_go__go_3_5_5_v1_5 = $v1_5;
  tco_loop_go__go_3_5_5:;
  $v_4 = $__tco_var_go__go_3_5_5_v_4;
  $v1_5 = $__tco_var_go__go_3_5_5_v1_5;
  $__t5 = null;;
  if ($v1_5 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t5 = $v_4;
goto end_branch_5;;
};
  if ($v1_5 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_6 = new \Data\List\Types\Data_List_Types_Cons(($v1_5)->{'value0'}, $v_4);
$__tco_7 = ($v1_5)->{'value1'};
$__tco_var_go__go_3_5_5_v_4 = $__tco_6;
$__tco_var_go__go_3_5_5_v1_5 = $__tco_7;
goto tco_loop_go__go_3_5_5;;
$__t5 = null;
goto end_branch_5;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t5 = null;
  end_branch_5:;
  $__res = $__t5;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = (\Data\Foldable\majData_majFoldable_foldrmajArray(function($x_3) use ($Semigroup0_2_4) {
  $__num = \func_num_args();
  $__res = function($acc_4) use ($Semigroup0_2_4, $x_3) {
  $__num = \func_num_args();
  $__res = ((($Semigroup0_2_4)->{'append'})($x_3))($acc_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($__local_var_1_0)->{'mempty'}, [$GLOBALS['Data_Interval_Duration_Iso_checkWeekUsage'], $GLOBALS['Data_Interval_Duration_Iso_checkEmptiness'], $GLOBALS['Data_Interval_Duration_Iso_checkFractionalUse'], $GLOBALS['Data_Interval_Duration_Iso_checkNegativeValues']]))((object)["asList" => (($go__go_3_5)(new \Data\List\Types\Data_List_Types_Nil()))(\Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($b_4) {
  $__num = \func_num_args();
  $go__go_5_6 = null;
  $go__go_5_6 = (function() use (&$go__go_5_6) {
  $__fn = function($source_6, $memo_7 = null) use (&$go__go_5_6, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_5_6_6_source_6 = $source_6;
  $__tco_var_go__go_5_6_6_memo_7 = $memo_7;
  tco_loop_go__go_5_6_6:;
  $source_6 = $__tco_var_go__go_5_6_6_source_6;
  $memo_7 = $__tco_var_go__go_5_6_6_memo_7;
  $v_8_6 = \Data\Map\Internal\majData_majMap_majInternal_stepmajUnfoldr($source_6);
  $__t7 = null;;
  if ($v_8_6 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$go__go_9_8 = null;
$go__go_9_8 = (function() use (&$__tco_var_go__go_5_6_6_source_6, &$__tco_var_go__go_5_6_6_memo_7, &$go__go_9_8) {
  $__fn = function($b_10, $v_11 = null) use (&$__tco_var_go__go_5_6_6_source_6, &$__tco_var_go__go_5_6_6_memo_7, &$go__go_9_8, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_9_8_8_b_10 = $b_10;
  $__tco_var_go__go_9_8_8_v_11 = $v_11;
  tco_loop_go__go_9_8_8:;
  $b_10 = $__tco_var_go__go_9_8_8_b_10;
  $v_11 = $__tco_var_go__go_9_8_8_v_11;
  $__t8 = null;;
  if ($v_11 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t8 = $b_10;
goto end_branch_8;;
};
  if ($v_11 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_9 = new \Data\List\Types\Data_List_Types_Cons(($v_11)->{'value0'}, $b_10);
$__tco_10 = ($v_11)->{'value1'};
$__tco_var_go__go_9_8_8_b_10 = $__tco_9;
$__tco_var_go__go_9_8_8_v_11 = $__tco_10;
goto tco_loop_go__go_9_8_8;;
$__t8 = null;
goto end_branch_8;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t8 = null;
  end_branch_8:;
  $__res = $__t8;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
$__t7 = (($go__go_9_8)(new \Data\List\Types\Data_List_Types_Nil()))($memo_7);
goto end_branch_7;;
};
  if ($v_8_6 instanceof \Data\Maybe\Data_Maybe_Just) {
$__tco_9 = (($v_8_6)->{'value0'})->{'value1'};
$__tco_10 = new \Data\List\Types\Data_List_Types_Cons((($v_8_6)->{'value0'})->{'value0'}, $memo_7);
$__tco_var_go__go_5_6_6_source_6 = $__tco_9;
$__tco_var_go__go_5_6_6_memo_7 = $__tco_10;
goto tco_loop_go__go_5_6_6;;
$__t7 = null;
goto end_branch_7;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t7 = null;
  end_branch_7:;
  $__res = $__t7;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = (($go__go_5_6)($b_4))(new \Data\List\Types\Data_List_Types_Nil());
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $GLOBALS['Data_Map_Internal_toMapIter'], $v_0)), "asMap" => $v_0]);
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
  $__local_var_1_0 = \Data\Interval\Duration\Iso\majData_majInterval_majDuration_majIso_checkmajValidmajIsomajDuration($d_0);
  $__t1 = null;;
  if ($__local_var_1_0 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_1;;
};
  if ($__local_var_1_0 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__t1 = new \Data\Maybe\Data_Maybe_Just(new \Data\NonEmpty\Data_NonEmpty_NonEmpty(($__local_var_1_0)->{'value0'}, ($__local_var_1_0)->{'value1'}));
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $v_1_0 = $__t1;
  $__t3 = null;;
  if ($v_1_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t3 = new \Data\Either\Data_Either_Left(($v_1_0)->{'value0'});
goto end_branch_3;;
};
  if ($v_1_0 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t3 = new \Data\Either\Data_Either_Right($d_0);
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
$GLOBALS['Data_Interval_Duration_Iso_mkIsoDuration'] = __NAMESPACE__ . '\\majData_majInterval_majDuration_majIso_mkmajIsomajDuration';

