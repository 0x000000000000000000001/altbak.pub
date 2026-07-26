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
\PhpursThunks::$thunks['Data_Interval_Duration_Second'] = function() { $v = ($GLOBALS['__phpurs_data0_Second'] ??= new Phpurs_Data0("Second")); return $v; };
\PhpursThunks::$thunks['Data_Interval_Duration_Minute'] = function() { $v = ($GLOBALS['__phpurs_data0_Minute'] ??= new Phpurs_Data0("Minute")); return $v; };
\PhpursThunks::$thunks['Data_Interval_Duration_Hour'] = function() { $v = ($GLOBALS['__phpurs_data0_Hour'] ??= new Phpurs_Data0("Hour")); return $v; };
\PhpursThunks::$thunks['Data_Interval_Duration_Day'] = function() { $v = ($GLOBALS['__phpurs_data0_Day'] ??= new Phpurs_Data0("Day")); return $v; };
\PhpursThunks::$thunks['Data_Interval_Duration_Week'] = function() { $v = ($GLOBALS['__phpurs_data0_Week'] ??= new Phpurs_Data0("Week")); return $v; };
\PhpursThunks::$thunks['Data_Interval_Duration_Month'] = function() { $v = ($GLOBALS['__phpurs_data0_Month'] ??= new Phpurs_Data0("Month")); return $v; };
\PhpursThunks::$thunks['Data_Interval_Duration_Year'] = function() { $v = ($GLOBALS['__phpurs_data0_Year'] ??= new Phpurs_Data0("Year")); return $v; };
\PhpursThunks::$thunks['Data_Interval_Duration_Duration'] = function() { $v = function($x_0 = null) {
  $__num = \func_num_args();
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Interval_Duration_showDurationComponent'] = function() { $v = (object)["show" => function($v_0 = null) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Minute"))) {
$__t0 = "Minute";
goto end_branch_0;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Second"))) {
$__t0 = "Second";
goto end_branch_0;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Hour"))) {
$__t0 = "Hour";
goto end_branch_0;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Day"))) {
$__t0 = "Day";
goto end_branch_0;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Week"))) {
$__t0 = "Week";
goto end_branch_0;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Month"))) {
$__t0 = "Month";
goto end_branch_0;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Year"))) {
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
}]; return $v; };
\PhpursThunks::$thunks['Data_Interval_Duration_show'] = function() { $v = (((($GLOBALS['Data_Map_Internal_showMap'] ?? \PhpursThunks::eval('Data_Map_Internal_showMap')))(($GLOBALS['Data_Interval_Duration_showDurationComponent'] ?? \PhpursThunks::eval('Data_Interval_Duration_showDurationComponent'))))(($GLOBALS['Data_Show_showNumber'] ?? \PhpursThunks::eval('Data_Show_showNumber'))))->{'show'}; return $v; };
\PhpursThunks::$thunks['Data_Interval_Duration_showDuration'] = function() { $v = (object)["show" => function($v_0 = null) {
  $__num = \func_num_args();
  $__res = (("(Duration " . (($GLOBALS['Data_Interval_Duration_show'] ?? \PhpursThunks::eval('Data_Interval_Duration_show')))($v_0)) . ")");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Interval_Duration_newtypeDuration'] = function() { $v = (object)["Coercible0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Interval_Duration_eqDurationComponent'] = function() { $v = (object)["eq" => (function() {
  $__fn = function($x_0 = null, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ((is_object($x_0) && (($x_0)->{'tag'} === "Second"))) {
$__t0 = (is_object($y_1) && (($y_1)->{'tag'} === "Second"));
goto end_branch_0;;
};
  if ((is_object($x_0) && (($x_0)->{'tag'} === "Minute"))) {
$__t0 = (is_object($y_1) && (($y_1)->{'tag'} === "Minute"));
goto end_branch_0;;
};
  if ((is_object($x_0) && (($x_0)->{'tag'} === "Hour"))) {
$__t0 = (is_object($y_1) && (($y_1)->{'tag'} === "Hour"));
goto end_branch_0;;
};
  if ((is_object($x_0) && (($x_0)->{'tag'} === "Day"))) {
$__t0 = (is_object($y_1) && (($y_1)->{'tag'} === "Day"));
goto end_branch_0;;
};
  if ((is_object($x_0) && (($x_0)->{'tag'} === "Week"))) {
$__t0 = (is_object($y_1) && (($y_1)->{'tag'} === "Week"));
goto end_branch_0;;
};
  if ((is_object($x_0) && (($x_0)->{'tag'} === "Month"))) {
$__t0 = (is_object($y_1) && (($y_1)->{'tag'} === "Month"));
goto end_branch_0;;
};
  $__t0 = ((is_object($x_0) && (($x_0)->{'tag'} === "Year")) && (is_object($y_1) && (($y_1)->{'tag'} === "Year")));
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Data_Interval_Duration_eq'] = function() { $v = (((($GLOBALS['Data_Map_Internal_eqMap'] ?? \PhpursThunks::eval('Data_Map_Internal_eqMap')))(($GLOBALS['Data_Interval_Duration_eqDurationComponent'] ?? \PhpursThunks::eval('Data_Interval_Duration_eqDurationComponent'))))(($GLOBALS['Data_Eq_eqNumber'] ?? \PhpursThunks::eval('Data_Eq_eqNumber'))))->{'eq'}; return $v; };
\PhpursThunks::$thunks['Data_Interval_Duration_ordDurationComponent'] = function() { $v = (object)["compare" => (function() {
  $__fn = function($x_0 = null, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ((is_object($x_0) && (($x_0)->{'tag'} === "Second"))) {
$__t1 = null;;
if ((is_object($y_1) && (($y_1)->{'tag'} === "Second"))) {
$__t1 = new Phpurs_Data0("EQ");
goto end_branch_1;;
};
$__t1 = new Phpurs_Data0("LT");
end_branch_1:;
$__t0 = $__t1;
goto end_branch_0;;
};
  if ((is_object($y_1) && (($y_1)->{'tag'} === "Second"))) {
$__t0 = new Phpurs_Data0("GT");
goto end_branch_0;;
};
  if ((is_object($x_0) && (($x_0)->{'tag'} === "Minute"))) {
$__t2 = null;;
if ((is_object($y_1) && (($y_1)->{'tag'} === "Minute"))) {
$__t2 = new Phpurs_Data0("EQ");
goto end_branch_2;;
};
$__t2 = new Phpurs_Data0("LT");
end_branch_2:;
$__t0 = $__t2;
goto end_branch_0;;
};
  if ((is_object($y_1) && (($y_1)->{'tag'} === "Minute"))) {
$__t0 = new Phpurs_Data0("GT");
goto end_branch_0;;
};
  if ((is_object($x_0) && (($x_0)->{'tag'} === "Hour"))) {
$__t3 = null;;
if ((is_object($y_1) && (($y_1)->{'tag'} === "Hour"))) {
$__t3 = new Phpurs_Data0("EQ");
goto end_branch_3;;
};
$__t3 = new Phpurs_Data0("LT");
end_branch_3:;
$__t0 = $__t3;
goto end_branch_0;;
};
  if ((is_object($y_1) && (($y_1)->{'tag'} === "Hour"))) {
$__t0 = new Phpurs_Data0("GT");
goto end_branch_0;;
};
  if ((is_object($x_0) && (($x_0)->{'tag'} === "Day"))) {
$__t4 = null;;
if ((is_object($y_1) && (($y_1)->{'tag'} === "Day"))) {
$__t4 = new Phpurs_Data0("EQ");
goto end_branch_4;;
};
$__t4 = new Phpurs_Data0("LT");
end_branch_4:;
$__t0 = $__t4;
goto end_branch_0;;
};
  if ((is_object($y_1) && (($y_1)->{'tag'} === "Day"))) {
$__t0 = new Phpurs_Data0("GT");
goto end_branch_0;;
};
  if ((is_object($x_0) && (($x_0)->{'tag'} === "Week"))) {
$__t5 = null;;
if ((is_object($y_1) && (($y_1)->{'tag'} === "Week"))) {
$__t5 = new Phpurs_Data0("EQ");
goto end_branch_5;;
};
$__t5 = new Phpurs_Data0("LT");
end_branch_5:;
$__t0 = $__t5;
goto end_branch_0;;
};
  if ((is_object($y_1) && (($y_1)->{'tag'} === "Week"))) {
$__t0 = new Phpurs_Data0("GT");
goto end_branch_0;;
};
  if ((is_object($x_0) && (($x_0)->{'tag'} === "Month"))) {
$__t6 = null;;
if ((is_object($y_1) && (($y_1)->{'tag'} === "Month"))) {
$__t6 = new Phpurs_Data0("EQ");
goto end_branch_6;;
};
$__t6 = new Phpurs_Data0("LT");
end_branch_6:;
$__t0 = $__t6;
goto end_branch_0;;
};
  if ((is_object($y_1) && (($y_1)->{'tag'} === "Month"))) {
$__t0 = new Phpurs_Data0("GT");
goto end_branch_0;;
};
  if (((is_object($x_0) && (($x_0)->{'tag'} === "Year")) && (is_object($y_1) && (($y_1)->{'tag'} === "Year")))) {
$__t0 = new Phpurs_Data0("EQ");
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
})(), "Eq0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Interval_Duration_eqDurationComponent'] ?? \PhpursThunks::eval('Data_Interval_Duration_eqDurationComponent'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Interval_Duration_compare'] = function() { $v = (((($GLOBALS['Data_Map_Internal_ordMap'] ?? \PhpursThunks::eval('Data_Map_Internal_ordMap')))(($GLOBALS['Data_Interval_Duration_ordDurationComponent'] ?? \PhpursThunks::eval('Data_Interval_Duration_ordDurationComponent'))))(($GLOBALS['Data_Ord_ordNumber'] ?? \PhpursThunks::eval('Data_Ord_ordNumber'))))->{'compare'}; return $v; };
\PhpursThunks::$thunks['Data_Interval_Duration_semigroupDuration'] = function() { $v = (object)["append" => (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_Map_Internal_unsafeUnionWith'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeUnionWith')))((($GLOBALS['Data_Interval_Duration_ordDurationComponent'] ?? \PhpursThunks::eval('Data_Interval_Duration_ordDurationComponent')))->{'compare'}, ($GLOBALS['Data_Semiring_numAdd'] ?? \PhpursThunks::eval('Data_Semiring_numAdd')), $v_0, $v1_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Data_Interval_Duration_monoidDuration'] = function() { $v = (object)["mempty" => new Phpurs_Data0("Leaf"), "Semigroup0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Interval_Duration_semigroupDuration'] ?? \PhpursThunks::eval('Data_Interval_Duration_semigroupDuration'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Interval_Duration_eqDuration'] = function() { $v = (object)["eq" => (function() {
  $__fn = function($x_0 = null, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_Interval_Duration_eq'] ?? \PhpursThunks::eval('Data_Interval_Duration_eq')))($x_0))($y_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Data_Interval_Duration_ordDuration'] = function() { $v = (object)["compare" => (function() {
  $__fn = function($x_0 = null, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_Interval_Duration_compare'] ?? \PhpursThunks::eval('Data_Interval_Duration_compare')))($x_0))($y_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Interval_Duration_eqDuration'] ?? \PhpursThunks::eval('Data_Interval_Duration_eqDuration'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Interval_Duration_hour'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data6("Node", 1, 1, new Phpurs_Data0("Hour"), $v_0, new Phpurs_Data0("Leaf"), new Phpurs_Data0("Leaf"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Interval_Duration_millisecond'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(function($v_0 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data6("Node", 1, 1, new Phpurs_Data0("Second"), $v_0, new Phpurs_Data0("Leaf"), new Phpurs_Data0("Leaf"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ($v_0 / 1000.0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}); return $v; };
\PhpursThunks::$thunks['Data_Interval_Duration_minute'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data6("Node", 1, 1, new Phpurs_Data0("Minute"), $v_0, new Phpurs_Data0("Leaf"), new Phpurs_Data0("Leaf"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Interval_Duration_month'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data6("Node", 1, 1, new Phpurs_Data0("Month"), $v_0, new Phpurs_Data0("Leaf"), new Phpurs_Data0("Leaf"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Interval_Duration_second'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data6("Node", 1, 1, new Phpurs_Data0("Second"), $v_0, new Phpurs_Data0("Leaf"), new Phpurs_Data0("Leaf"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Interval_Duration_week'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data6("Node", 1, 1, new Phpurs_Data0("Week"), $v_0, new Phpurs_Data0("Leaf"), new Phpurs_Data0("Leaf"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Interval_Duration_year'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data6("Node", 1, 1, new Phpurs_Data0("Year"), $v_0, new Phpurs_Data0("Leaf"), new Phpurs_Data0("Leaf"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Interval_Duration_day'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data6("Node", 1, 1, new Phpurs_Data0("Day"), $v_0, new Phpurs_Data0("Leaf"), new Phpurs_Data0("Leaf"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };






























