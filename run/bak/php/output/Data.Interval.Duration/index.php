<?php

namespace Data\Interval\Duration;

// ALL IMPORTS: Control.Semigroupoid, Data.Eq, Data.EuclideanRing, Data.Interval.Duration, Data.Map, Data.Map.Internal, Data.Monoid, Data.Newtype, Data.Ord, Data.Ordering, Data.Semigroup, Data.Semiring, Data.Show, Prelude, Prim
// TO REQUIRE: Control.Semigroupoid, Data.Eq, Data.EuclideanRing, Data.Interval.Duration, Data.Map, Data.Map.Internal, Data.Monoid, Data.Newtype, Data.Ord, Data.Ordering, Data.Semigroup, Data.Semiring, Data.Show, Prelude
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.EuclideanRing/index.php';
require_once __DIR__ . '/../Data.Interval.Duration/index.php';
require_once __DIR__ . '/../Data.Map/index.php';
require_once __DIR__ . '/../Data.Map.Internal/index.php';
require_once __DIR__ . '/../Data.Monoid/index.php';
require_once __DIR__ . '/../Data.Newtype/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ordering/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Semiring/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
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


final class Data_Interval_Duration_Second { public $tag = 'Second'; public function __construct() {} }
final class Data_Interval_Duration_Minute { public $tag = 'Minute'; public function __construct() {} }
final class Data_Interval_Duration_Hour { public $tag = 'Hour'; public function __construct() {} }
final class Data_Interval_Duration_Day { public $tag = 'Day'; public function __construct() {} }
final class Data_Interval_Duration_Week { public $tag = 'Week'; public function __construct() {} }
final class Data_Interval_Duration_Month { public $tag = 'Month'; public function __construct() {} }
final class Data_Interval_Duration_Year { public $tag = 'Year'; public function __construct() {} }

// Data_Interval_Duration_Second
$GLOBALS['Data_Interval_Duration_Second'] = ($GLOBALS['__phpurs_data0_Second'] ??= new \Data\Interval\Duration\Data_Interval_Duration_Second());

// Data_Interval_Duration_Minute
$GLOBALS['Data_Interval_Duration_Minute'] = ($GLOBALS['__phpurs_data0_Minute'] ??= new \Data\Interval\Duration\Data_Interval_Duration_Minute());

// Data_Interval_Duration_Hour
$GLOBALS['Data_Interval_Duration_Hour'] = ($GLOBALS['__phpurs_data0_Hour'] ??= new \Data\Interval\Duration\Data_Interval_Duration_Hour());

// Data_Interval_Duration_Day
$GLOBALS['Data_Interval_Duration_Day'] = ($GLOBALS['__phpurs_data0_Day'] ??= new \Data\Interval\Duration\Data_Interval_Duration_Day());

// Data_Interval_Duration_Week
$GLOBALS['Data_Interval_Duration_Week'] = ($GLOBALS['__phpurs_data0_Week'] ??= new \Data\Interval\Duration\Data_Interval_Duration_Week());

// Data_Interval_Duration_Month
$GLOBALS['Data_Interval_Duration_Month'] = ($GLOBALS['__phpurs_data0_Month'] ??= new \Data\Interval\Duration\Data_Interval_Duration_Month());

// Data_Interval_Duration_Year
$GLOBALS['Data_Interval_Duration_Year'] = ($GLOBALS['__phpurs_data0_Year'] ??= new \Data\Interval\Duration\Data_Interval_Duration_Year());

// Data_Interval_Duration_Duration
function majData_majInterval_majDuration_majDuration($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majInterval_majDuration_majDuration';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Interval_Duration_Duration'] = __NAMESPACE__ . '\\majData_majInterval_majDuration_majDuration';

// Data_Interval_Duration_showDurationComponent
$GLOBALS['Data_Interval_Duration_showDurationComponent'] = (object)["show" => function($v_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v_0 instanceof \Data\Interval\Duration\Data_Interval_Duration_Minute) {
$__t0 = "Minute";
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Interval\Duration\Data_Interval_Duration_Second) {
$__t0 = "Second";
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Interval\Duration\Data_Interval_Duration_Hour) {
$__t0 = "Hour";
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Interval\Duration\Data_Interval_Duration_Day) {
$__t0 = "Day";
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Interval\Duration\Data_Interval_Duration_Week) {
$__t0 = "Week";
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Interval\Duration\Data_Interval_Duration_Month) {
$__t0 = "Month";
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Interval\Duration\Data_Interval_Duration_Year) {
$__t0 = "Year";
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

// Data_Interval_Duration_show
$GLOBALS['Data_Interval_Duration_show'] = ((($GLOBALS['Data_Map_Internal_showMap'])($GLOBALS['Data_Interval_Duration_showDurationComponent']))($GLOBALS['Data_Show_showNumber']))->{'show'};

// Data_Interval_Duration_showDuration
$GLOBALS['Data_Interval_Duration_showDuration'] = (object)["show" => function($v_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})("(Duration "))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})(($GLOBALS['Data_Interval_Duration_show'])($v_0)))(")"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Interval_Duration_newtypeDuration
$GLOBALS['Data_Interval_Duration_newtypeDuration'] = (object)["Coercible0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Interval_Duration_eqDurationComponent
$GLOBALS['Data_Interval_Duration_eqDurationComponent'] = (object)["eq" => function($x_0) {
  $__num = \func_num_args();
  $__res = function($y_1) use ($x_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($x_0 instanceof \Data\Interval\Duration\Data_Interval_Duration_Second) {
$__t1 = null;;
if ($y_1 instanceof \Data\Interval\Duration\Data_Interval_Duration_Second) {
$__t1 = true;
goto end_branch_1;;
};
$__t1 = false;
end_branch_1:;
$__t0 = $__t1;
goto end_branch_0;;
};
  if ($x_0 instanceof \Data\Interval\Duration\Data_Interval_Duration_Minute) {
$__t2 = null;;
if ($y_1 instanceof \Data\Interval\Duration\Data_Interval_Duration_Minute) {
$__t2 = true;
goto end_branch_2;;
};
$__t2 = false;
end_branch_2:;
$__t0 = $__t2;
goto end_branch_0;;
};
  if ($x_0 instanceof \Data\Interval\Duration\Data_Interval_Duration_Hour) {
$__t3 = null;;
if ($y_1 instanceof \Data\Interval\Duration\Data_Interval_Duration_Hour) {
$__t3 = true;
goto end_branch_3;;
};
$__t3 = false;
end_branch_3:;
$__t0 = $__t3;
goto end_branch_0;;
};
  if ($x_0 instanceof \Data\Interval\Duration\Data_Interval_Duration_Day) {
$__t4 = null;;
if ($y_1 instanceof \Data\Interval\Duration\Data_Interval_Duration_Day) {
$__t4 = true;
goto end_branch_4;;
};
$__t4 = false;
end_branch_4:;
$__t0 = $__t4;
goto end_branch_0;;
};
  if ($x_0 instanceof \Data\Interval\Duration\Data_Interval_Duration_Week) {
$__t5 = null;;
if ($y_1 instanceof \Data\Interval\Duration\Data_Interval_Duration_Week) {
$__t5 = true;
goto end_branch_5;;
};
$__t5 = false;
end_branch_5:;
$__t0 = $__t5;
goto end_branch_0;;
};
  if ($x_0 instanceof \Data\Interval\Duration\Data_Interval_Duration_Month) {
$__t6 = null;;
if ($y_1 instanceof \Data\Interval\Duration\Data_Interval_Duration_Month) {
$__t6 = true;
goto end_branch_6;;
};
$__t6 = false;
end_branch_6:;
$__t0 = $__t6;
goto end_branch_0;;
};
  if (($x_0 instanceof \Data\Interval\Duration\Data_Interval_Duration_Year && $y_1 instanceof \Data\Interval\Duration\Data_Interval_Duration_Year)) {
$__t0 = true;
goto end_branch_0;;
};
  $__t0 = false;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Interval_Duration_eq
$GLOBALS['Data_Interval_Duration_eq'] = ((($GLOBALS['Data_Map_Internal_eqMap'])($GLOBALS['Data_Interval_Duration_eqDurationComponent']))($GLOBALS['Data_Eq_eqNumber']))->{'eq'};

// Data_Interval_Duration_ordDurationComponent
$GLOBALS['Data_Interval_Duration_ordDurationComponent'] = (object)["compare" => function($x_0) {
  $__num = \func_num_args();
  $__res = function($y_1) use ($x_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($x_0 instanceof \Data\Interval\Duration\Data_Interval_Duration_Second) {
$__t1 = null;;
if ($y_1 instanceof \Data\Interval\Duration\Data_Interval_Duration_Second) {
$__t1 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_1;;
};
$__t1 = new \Data\Ordering\Data_Ordering_LT();
end_branch_1:;
$__t0 = $__t1;
goto end_branch_0;;
};
  if ($y_1 instanceof \Data\Interval\Duration\Data_Interval_Duration_Second) {
$__t0 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_0;;
};
  if ($x_0 instanceof \Data\Interval\Duration\Data_Interval_Duration_Minute) {
$__t2 = null;;
if ($y_1 instanceof \Data\Interval\Duration\Data_Interval_Duration_Minute) {
$__t2 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_2;;
};
$__t2 = new \Data\Ordering\Data_Ordering_LT();
end_branch_2:;
$__t0 = $__t2;
goto end_branch_0;;
};
  if ($y_1 instanceof \Data\Interval\Duration\Data_Interval_Duration_Minute) {
$__t0 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_0;;
};
  if ($x_0 instanceof \Data\Interval\Duration\Data_Interval_Duration_Hour) {
$__t3 = null;;
if ($y_1 instanceof \Data\Interval\Duration\Data_Interval_Duration_Hour) {
$__t3 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_3;;
};
$__t3 = new \Data\Ordering\Data_Ordering_LT();
end_branch_3:;
$__t0 = $__t3;
goto end_branch_0;;
};
  if ($y_1 instanceof \Data\Interval\Duration\Data_Interval_Duration_Hour) {
$__t0 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_0;;
};
  if ($x_0 instanceof \Data\Interval\Duration\Data_Interval_Duration_Day) {
$__t4 = null;;
if ($y_1 instanceof \Data\Interval\Duration\Data_Interval_Duration_Day) {
$__t4 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_4;;
};
$__t4 = new \Data\Ordering\Data_Ordering_LT();
end_branch_4:;
$__t0 = $__t4;
goto end_branch_0;;
};
  if ($y_1 instanceof \Data\Interval\Duration\Data_Interval_Duration_Day) {
$__t0 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_0;;
};
  if ($x_0 instanceof \Data\Interval\Duration\Data_Interval_Duration_Week) {
$__t5 = null;;
if ($y_1 instanceof \Data\Interval\Duration\Data_Interval_Duration_Week) {
$__t5 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_5;;
};
$__t5 = new \Data\Ordering\Data_Ordering_LT();
end_branch_5:;
$__t0 = $__t5;
goto end_branch_0;;
};
  if ($y_1 instanceof \Data\Interval\Duration\Data_Interval_Duration_Week) {
$__t0 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_0;;
};
  if ($x_0 instanceof \Data\Interval\Duration\Data_Interval_Duration_Month) {
$__t6 = null;;
if ($y_1 instanceof \Data\Interval\Duration\Data_Interval_Duration_Month) {
$__t6 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_6;;
};
$__t6 = new \Data\Ordering\Data_Ordering_LT();
end_branch_6:;
$__t0 = $__t6;
goto end_branch_0;;
};
  if ($y_1 instanceof \Data\Interval\Duration\Data_Interval_Duration_Month) {
$__t0 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_0;;
};
  if (($x_0 instanceof \Data\Interval\Duration\Data_Interval_Duration_Year && $y_1 instanceof \Data\Interval\Duration\Data_Interval_Duration_Year)) {
$__t0 = new \Data\Ordering\Data_Ordering_EQ();
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
}, "Eq0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Interval_Duration_eqDurationComponent'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Interval_Duration_compare
$GLOBALS['Data_Interval_Duration_compare'] = ((($GLOBALS['Data_Map_Internal_ordMap'])($GLOBALS['Data_Interval_Duration_ordDurationComponent']))($GLOBALS['Data_Ord_ordNumber']))->{'compare'};

// Data_Interval_Duration_semigroupDuration
$GLOBALS['Data_Interval_Duration_semigroupDuration'] = (object)["append" => function($v_0) {
  $__num = \func_num_args();
  $__res = function($v1_1) use ($v_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Map_Internal_unsafeUnionWith'])(($GLOBALS['Data_Interval_Duration_ordDurationComponent'])->{'compare'}, $GLOBALS['Data_Semiring_numAdd'], $v_0, $v1_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Interval_Duration_monoidDuration
$GLOBALS['Data_Interval_Duration_monoidDuration'] = (object)["mempty" => new \Data\Map\Internal\Data_Map_Internal_Leaf(), "Semigroup0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Interval_Duration_semigroupDuration'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Interval_Duration_eqDuration
$GLOBALS['Data_Interval_Duration_eqDuration'] = (object)["eq" => function($x_0) {
  $__num = \func_num_args();
  $__res = function($y_1) use ($x_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Interval_Duration_eq'])($x_0))($y_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Interval_Duration_ordDuration
$GLOBALS['Data_Interval_Duration_ordDuration'] = (object)["compare" => function($x_0) {
  $__num = \func_num_args();
  $__res = function($y_1) use ($x_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Interval_Duration_compare'])($x_0))($y_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Interval_Duration_eqDuration'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Interval_Duration_hour
function majData_majInterval_majDuration_hour(float $v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majInterval_majDuration_hour';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = new \Data\Map\Internal\Data_Map_Internal_Node(1, 1, new \Data\Interval\Duration\Data_Interval_Duration_Hour(), $v_0, new \Data\Map\Internal\Data_Map_Internal_Leaf(), new \Data\Map\Internal\Data_Map_Internal_Leaf());
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Interval_Duration_hour'] = __NAMESPACE__ . '\\majData_majInterval_majDuration_hour';

// Data_Interval_Duration_millisecond_closure
$GLOBALS['Data_Interval_Duration_millisecond_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_0) {
  $__num = \func_num_args();
  $__res = new \Data\Map\Internal\Data_Map_Internal_Node(1, 1, new \Data\Interval\Duration\Data_Interval_Duration_Second(), $v_0, new \Data\Map\Internal\Data_Map_Internal_Leaf(), new \Data\Map\Internal\Data_Map_Internal_Leaf());
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_EuclideanRing_euclideanRingNumber'])->{'div'})($v_0))(1000.0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});

// Data_Interval_Duration_millisecond
function majData_majInterval_majDuration_millisecond(float $v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majInterval_majDuration_millisecond';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Interval_Duration_millisecond_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Interval_Duration_millisecond'] = __NAMESPACE__ . '\\majData_majInterval_majDuration_millisecond';

// Data_Interval_Duration_minute
function majData_majInterval_majDuration_minute(float $v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majInterval_majDuration_minute';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = new \Data\Map\Internal\Data_Map_Internal_Node(1, 1, new \Data\Interval\Duration\Data_Interval_Duration_Minute(), $v_0, new \Data\Map\Internal\Data_Map_Internal_Leaf(), new \Data\Map\Internal\Data_Map_Internal_Leaf());
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Interval_Duration_minute'] = __NAMESPACE__ . '\\majData_majInterval_majDuration_minute';

// Data_Interval_Duration_month
function majData_majInterval_majDuration_month(float $v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majInterval_majDuration_month';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = new \Data\Map\Internal\Data_Map_Internal_Node(1, 1, new \Data\Interval\Duration\Data_Interval_Duration_Month(), $v_0, new \Data\Map\Internal\Data_Map_Internal_Leaf(), new \Data\Map\Internal\Data_Map_Internal_Leaf());
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Interval_Duration_month'] = __NAMESPACE__ . '\\majData_majInterval_majDuration_month';

// Data_Interval_Duration_second
function majData_majInterval_majDuration_second(float $v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majInterval_majDuration_second';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = new \Data\Map\Internal\Data_Map_Internal_Node(1, 1, new \Data\Interval\Duration\Data_Interval_Duration_Second(), $v_0, new \Data\Map\Internal\Data_Map_Internal_Leaf(), new \Data\Map\Internal\Data_Map_Internal_Leaf());
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Interval_Duration_second'] = __NAMESPACE__ . '\\majData_majInterval_majDuration_second';

// Data_Interval_Duration_week
function majData_majInterval_majDuration_week(float $v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majInterval_majDuration_week';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = new \Data\Map\Internal\Data_Map_Internal_Node(1, 1, new \Data\Interval\Duration\Data_Interval_Duration_Week(), $v_0, new \Data\Map\Internal\Data_Map_Internal_Leaf(), new \Data\Map\Internal\Data_Map_Internal_Leaf());
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Interval_Duration_week'] = __NAMESPACE__ . '\\majData_majInterval_majDuration_week';

// Data_Interval_Duration_year
function majData_majInterval_majDuration_year(float $v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majInterval_majDuration_year';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = new \Data\Map\Internal\Data_Map_Internal_Node(1, 1, new \Data\Interval\Duration\Data_Interval_Duration_Year(), $v_0, new \Data\Map\Internal\Data_Map_Internal_Leaf(), new \Data\Map\Internal\Data_Map_Internal_Leaf());
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Interval_Duration_year'] = __NAMESPACE__ . '\\majData_majInterval_majDuration_year';

// Data_Interval_Duration_day
function majData_majInterval_majDuration_day(float $v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majInterval_majDuration_day';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = new \Data\Map\Internal\Data_Map_Internal_Node(1, 1, new \Data\Interval\Duration\Data_Interval_Duration_Day(), $v_0, new \Data\Map\Internal\Data_Map_Internal_Leaf(), new \Data\Map\Internal\Data_Map_Internal_Leaf());
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Interval_Duration_day'] = __NAMESPACE__ . '\\majData_majInterval_majDuration_day';

