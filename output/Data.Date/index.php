<?php

namespace Data\Date;

// ALL IMPORTS: Control.Applicative, Control.Apply, Control.Bind, Control.Semigroupoid, Data.Boolean, Data.Bounded, Data.Date, Data.Date.Component, Data.Enum, Data.Eq, Data.EuclideanRing, Data.Function, Data.Function.Uncurried, Data.Functor, Data.HeytingAlgebra, Data.Int, Data.Maybe, Data.Ord, Data.Ordering, Data.Ring, Data.Semigroup, Data.Semiring, Data.Show, Data.Time.Duration, Partial.Unsafe, Prelude, Prim
// TO REQUIRE: Control.Applicative, Control.Apply, Control.Bind, Control.Semigroupoid, Data.Boolean, Data.Bounded, Data.Date, Data.Date.Component, Data.Enum, Data.Eq, Data.EuclideanRing, Data.Function, Data.Function.Uncurried, Data.Functor, Data.HeytingAlgebra, Data.Int, Data.Maybe, Data.Ord, Data.Ordering, Data.Ring, Data.Semigroup, Data.Semiring, Data.Show, Data.Time.Duration, Partial.Unsafe, Prelude
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Boolean/index.php';
require_once __DIR__ . '/../Data.Bounded/index.php';
require_once __DIR__ . '/../Data.Date/index.php';
require_once __DIR__ . '/../Data.Date.Component/index.php';
require_once __DIR__ . '/../Data.Enum/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.EuclideanRing/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Function.Uncurried/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.HeytingAlgebra/index.php';
require_once __DIR__ . '/../Data.Int/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ordering/index.php';
require_once __DIR__ . '/../Data.Ring/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Semiring/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
require_once __DIR__ . '/../Data.Time.Duration/index.php';
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
$ffi_Data_Date = \call_user_func(function() {
  $exports = [];
$createDate = function($y, $m, $d) {
    $dt = new \DateTime('now', new \DateTimeZone('UTC'));
    $dt->setDate($y, $m + 1, $d);
    $dt->setTime(0, 0, 0, 0);
    return $dt;
};

$canonicalDateImpl = function($ctor, $y, $m, $d) use (&$canonicalDateImpl, $createDate) {

    $date = $createDate($y, $m - 1, $d);
    return $ctor
        ((int)$date->format('Y'))
        ((int)$date->format('n'))
        ((int)$date->format('j'));
};

$calcWeekday = function($y, $m, $d) use (&$calcWeekday, $createDate) {

    $date = $createDate($y, $m - 1, $d);
    return (int)$date->format('w'); // 0 (for Sunday) through 6 (for Saturday)
};

$calcDiff = function($y1, $m1, $d1, $y2, $m2, $d2) use (&$calcDiff, $createDate) {

    $dt1 = $createDate($y1, $m1 - 1, $d1);
    $dt2 = $createDate($y2, $m2 - 1, $d2);
    // returns diff in milliseconds
    return ($dt1->getTimestamp() - $dt2->getTimestamp()) * 1000;
};

$exports['createDate'] = $createDate;
$exports['canonicalDateImpl'] = $canonicalDateImpl;
$exports['calcWeekday'] = $calcWeekday;
$exports['calcDiff'] = $calcDiff;
return $exports;
  return $exports;
});
$GLOBALS['Data_Date_calcDiff'] = (\array_key_exists('calcDiff', $ffi_Data_Date) ? $ffi_Data_Date['calcDiff'] : new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Date_calcWeekday'] = (\array_key_exists('calcWeekday', $ffi_Data_Date) ? $ffi_Data_Date['calcWeekday'] : new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Date_canonicalDateImpl'] = (\array_key_exists('canonicalDateImpl', $ffi_Data_Date) ? $ffi_Data_Date['canonicalDateImpl'] : new class { public function __invoke(...$args) { return $this; } });


final class Data_Date_Date { public $tag = 'Date'; public function __construct(public int $value0, public  $value1, public int $value2) {} }

// Data_Date_Date
$GLOBALS['Data_Date_Date'] = (function() {
  $__fn = function($value0, $value1 = null, $value2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = new \Data\Date\Data_Date_Date($value0, $value1, $value2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_Date_year
function majData_majDate_year($v_0): int|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majDate_year';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($v_0)->{'value0'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Date_year'] = __NAMESPACE__ . '\\majData_majDate_year';

// Data_Date_weekday
function majData_majDate_weekday($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majDate_weekday';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__t0 = null;;
  if (($v_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_January) {
$__t0 = 1;
goto end_branch_0;;
};
  if (($v_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_February) {
$__t0 = 2;
goto end_branch_0;;
};
  if (($v_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_March) {
$__t0 = 3;
goto end_branch_0;;
};
  if (($v_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_April) {
$__t0 = 4;
goto end_branch_0;;
};
  if (($v_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_May) {
$__t0 = 5;
goto end_branch_0;;
};
  if (($v_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_June) {
$__t0 = 6;
goto end_branch_0;;
};
  if (($v_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_July) {
$__t0 = 7;
goto end_branch_0;;
};
  if (($v_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_August) {
$__t0 = 8;
goto end_branch_0;;
};
  if (($v_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_September) {
$__t0 = 9;
goto end_branch_0;;
};
  if (($v_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_October) {
$__t0 = 10;
goto end_branch_0;;
};
  if (($v_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_November) {
$__t0 = 11;
goto end_branch_0;;
};
  if (($v_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_December) {
$__t0 = 12;
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $n_1_0 = ($GLOBALS['Data_Date_calcWeekday'])(($v_0)->{'value0'}, $__t0, ($v_0)->{'value2'});
  $__t3 = null;;
  switch ($n_1_0) {
case 0:
$__t3 = new \Data\Date\Component\Data_Date_Component_Sunday();
goto end_branch_3;;
break;
default:
;
break;
};
  $__t2 = null;;
  switch ($n_1_0) {
case 1:
$__t2 = new \Data\Date\Component\Data_Date_Component_Monday();
goto end_branch_2;;
break;
default:
;
break;
};
  switch ($n_1_0) {
case 2:
$__t2 = new \Data\Date\Component\Data_Date_Component_Tuesday();
goto end_branch_2;;
break;
default:
;
break;
};
  switch ($n_1_0) {
case 3:
$__t2 = new \Data\Date\Component\Data_Date_Component_Wednesday();
goto end_branch_2;;
break;
default:
;
break;
};
  switch ($n_1_0) {
case 4:
$__t2 = new \Data\Date\Component\Data_Date_Component_Thursday();
goto end_branch_2;;
break;
default:
;
break;
};
  switch ($n_1_0) {
case 5:
$__t2 = new \Data\Date\Component\Data_Date_Component_Friday();
goto end_branch_2;;
break;
default:
;
break;
};
  switch ($n_1_0) {
case 6:
$__t2 = new \Data\Date\Component\Data_Date_Component_Saturday();
goto end_branch_2;;
break;
default:
;
break;
};
  switch ($n_1_0) {
case 7:
$__t2 = new \Data\Date\Component\Data_Date_Component_Sunday();
goto end_branch_2;;
break;
default:
;
break;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__t3 = $__t2;
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Date_weekday'] = __NAMESPACE__ . '\\majData_majDate_weekday';

// Data_Date_showDate
$GLOBALS['Data_Date_showDate'] = (object)["show" => function($v_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if (($v_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_January) {
$__t0 = "January";
goto end_branch_0;;
};
  if (($v_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_February) {
$__t0 = "February";
goto end_branch_0;;
};
  if (($v_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_March) {
$__t0 = "March";
goto end_branch_0;;
};
  if (($v_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_April) {
$__t0 = "April";
goto end_branch_0;;
};
  if (($v_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_May) {
$__t0 = "May";
goto end_branch_0;;
};
  if (($v_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_June) {
$__t0 = "June";
goto end_branch_0;;
};
  if (($v_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_July) {
$__t0 = "July";
goto end_branch_0;;
};
  if (($v_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_August) {
$__t0 = "August";
goto end_branch_0;;
};
  if (($v_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_September) {
$__t0 = "September";
goto end_branch_0;;
};
  if (($v_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_October) {
$__t0 = "October";
goto end_branch_0;;
};
  if (($v_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_November) {
$__t0 = "November";
goto end_branch_0;;
};
  if (($v_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_December) {
$__t0 = "December";
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = (((((("(Date (Year " . \Data\Show\majData_majShow_showmajIntmajImpl(($v_0)->{'value0'})) . ") ") . $__t0) . " (Day ") . \Data\Show\majData_majShow_showmajIntmajImpl(($v_0)->{'value2'})) . "))");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Date_month
function majData_majDate_month($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majDate_month';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($v_0)->{'value1'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Date_month'] = __NAMESPACE__ . '\\majData_majDate_month';

// Data_Date_isLeapYear
function majData_majDate_ismajLeapmajYear(int $y_0): bool|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majDate_ismajLeapmajYear';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ((\Data\EuclideanRing\majData_majEuclideanmajRing_intmajMod($y_0, 4) === 0) && ((\Data\EuclideanRing\majData_majEuclideanmajRing_intmajMod($y_0, 400) === 0) || ( ! (\Data\EuclideanRing\majData_majEuclideanmajRing_intmajMod($y_0, 100) === 0))));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Date_isLeapYear'] = __NAMESPACE__ . '\\majData_majDate_ismajLeapmajYear';

// Data_Date_lastDayOfMonth
function majData_majDate_lastmajDaymajOfmajMonth(int $y_0, $m_1 = null): int|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majDate_lastmajDaymajOfmajMonth';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $unsafeDay_2_0 = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_2) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t0 = ($v_2)->{'value0'};
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($n_2) {
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
});
  $__t3 = null;;
  if ($m_1 instanceof \Data\Date\Component\Data_Date_Component_January) {
$__t3 = ($unsafeDay_2_0)(31);
goto end_branch_3;;
};
  if ($m_1 instanceof \Data\Date\Component\Data_Date_Component_February) {
$__t4 = null;;
if (\Data\Date\majData_majDate_ismajLeapmajYear($y_0)) {
$__t4 = ($unsafeDay_2_0)(29);
goto end_branch_4;;
};
$__t4 = ($unsafeDay_2_0)(28);
end_branch_4:;
$__t3 = $__t4;
goto end_branch_3;;
};
  if ($m_1 instanceof \Data\Date\Component\Data_Date_Component_March) {
$__t3 = ($unsafeDay_2_0)(31);
goto end_branch_3;;
};
  if ($m_1 instanceof \Data\Date\Component\Data_Date_Component_April) {
$__t3 = ($unsafeDay_2_0)(30);
goto end_branch_3;;
};
  if ($m_1 instanceof \Data\Date\Component\Data_Date_Component_May) {
$__t3 = ($unsafeDay_2_0)(31);
goto end_branch_3;;
};
  if ($m_1 instanceof \Data\Date\Component\Data_Date_Component_June) {
$__t3 = ($unsafeDay_2_0)(30);
goto end_branch_3;;
};
  if ($m_1 instanceof \Data\Date\Component\Data_Date_Component_July) {
$__t3 = ($unsafeDay_2_0)(31);
goto end_branch_3;;
};
  if ($m_1 instanceof \Data\Date\Component\Data_Date_Component_August) {
$__t3 = ($unsafeDay_2_0)(31);
goto end_branch_3;;
};
  if ($m_1 instanceof \Data\Date\Component\Data_Date_Component_September) {
$__t3 = ($unsafeDay_2_0)(30);
goto end_branch_3;;
};
  if ($m_1 instanceof \Data\Date\Component\Data_Date_Component_October) {
$__t3 = ($unsafeDay_2_0)(31);
goto end_branch_3;;
};
  if ($m_1 instanceof \Data\Date\Component\Data_Date_Component_November) {
$__t3 = ($unsafeDay_2_0)(30);
goto end_branch_3;;
};
  if ($m_1 instanceof \Data\Date\Component\Data_Date_Component_December) {
$__t3 = ($unsafeDay_2_0)(31);
goto end_branch_3;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t3 = null;
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Date_lastDayOfMonth'] = __NAMESPACE__ . '\\majData_majDate_lastmajDaymajOfmajMonth';

// Data_Date_eqDate
$GLOBALS['Data_Date_eqDate'] = (object)["eq" => function($x_0) {
  $__num = \func_num_args();
  $__res = function($y_1) use ($x_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if (($x_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_January) {
$__t0 = ($y_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_January;
goto end_branch_0;;
};
  if (($x_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_February) {
$__t0 = ($y_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_February;
goto end_branch_0;;
};
  if (($x_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_March) {
$__t0 = ($y_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_March;
goto end_branch_0;;
};
  if (($x_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_April) {
$__t0 = ($y_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_April;
goto end_branch_0;;
};
  if (($x_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_May) {
$__t0 = ($y_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_May;
goto end_branch_0;;
};
  if (($x_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_June) {
$__t0 = ($y_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_June;
goto end_branch_0;;
};
  if (($x_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_July) {
$__t0 = ($y_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_July;
goto end_branch_0;;
};
  if (($x_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_August) {
$__t0 = ($y_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_August;
goto end_branch_0;;
};
  if (($x_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_September) {
$__t0 = ($y_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_September;
goto end_branch_0;;
};
  if (($x_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_October) {
$__t0 = ($y_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_October;
goto end_branch_0;;
};
  if (($x_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_November) {
$__t0 = ($y_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_November;
goto end_branch_0;;
};
  $__t0 = (($x_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_December && ($y_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_December);
  end_branch_0:;
  $__res = (((($x_0)->{'value0'} === ($y_1)->{'value0'}) && $__t0) && (($x_0)->{'value2'} === ($y_1)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Date_ordDate
$GLOBALS['Data_Date_ordDate'] = (object)["compare" => function($x_0) {
  $__num = \func_num_args();
  $__res = function($y_1) use ($x_0) {
  $__num = \func_num_args();
  $v_2_0 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), ($x_0)->{'value0'}, ($y_1)->{'value0'});
  $__t13 = null;;
  if ($v_2_0 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t13 = new \Data\Ordering\Data_Ordering_LT();
goto end_branch_13;;
};
  if ($v_2_0 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t13 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_13;;
};
  $__t1 = null;;
  if (($x_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_January) {
$__t2 = null;;
if (($y_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_January) {
$__t2 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), ($x_0)->{'value2'}, ($y_1)->{'value2'});
goto end_branch_2;;
};
$__t2 = new \Data\Ordering\Data_Ordering_LT();
end_branch_2:;
$__t1 = $__t2;
goto end_branch_1;;
};
  if (($y_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_January) {
$__t1 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_1;;
};
  if (($x_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_February) {
$__t3 = null;;
if (($y_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_February) {
$__t3 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), ($x_0)->{'value2'}, ($y_1)->{'value2'});
goto end_branch_3;;
};
$__t3 = new \Data\Ordering\Data_Ordering_LT();
end_branch_3:;
$__t1 = $__t3;
goto end_branch_1;;
};
  if (($y_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_February) {
$__t1 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_1;;
};
  if (($x_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_March) {
$__t4 = null;;
if (($y_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_March) {
$__t4 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), ($x_0)->{'value2'}, ($y_1)->{'value2'});
goto end_branch_4;;
};
$__t4 = new \Data\Ordering\Data_Ordering_LT();
end_branch_4:;
$__t1 = $__t4;
goto end_branch_1;;
};
  if (($y_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_March) {
$__t1 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_1;;
};
  if (($x_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_April) {
$__t5 = null;;
if (($y_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_April) {
$__t5 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), ($x_0)->{'value2'}, ($y_1)->{'value2'});
goto end_branch_5;;
};
$__t5 = new \Data\Ordering\Data_Ordering_LT();
end_branch_5:;
$__t1 = $__t5;
goto end_branch_1;;
};
  if (($y_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_April) {
$__t1 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_1;;
};
  if (($x_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_May) {
$__t6 = null;;
if (($y_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_May) {
$__t6 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), ($x_0)->{'value2'}, ($y_1)->{'value2'});
goto end_branch_6;;
};
$__t6 = new \Data\Ordering\Data_Ordering_LT();
end_branch_6:;
$__t1 = $__t6;
goto end_branch_1;;
};
  if (($y_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_May) {
$__t1 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_1;;
};
  if (($x_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_June) {
$__t7 = null;;
if (($y_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_June) {
$__t7 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), ($x_0)->{'value2'}, ($y_1)->{'value2'});
goto end_branch_7;;
};
$__t7 = new \Data\Ordering\Data_Ordering_LT();
end_branch_7:;
$__t1 = $__t7;
goto end_branch_1;;
};
  if (($y_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_June) {
$__t1 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_1;;
};
  if (($x_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_July) {
$__t8 = null;;
if (($y_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_July) {
$__t8 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), ($x_0)->{'value2'}, ($y_1)->{'value2'});
goto end_branch_8;;
};
$__t8 = new \Data\Ordering\Data_Ordering_LT();
end_branch_8:;
$__t1 = $__t8;
goto end_branch_1;;
};
  if (($y_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_July) {
$__t1 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_1;;
};
  if (($x_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_August) {
$__t9 = null;;
if (($y_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_August) {
$__t9 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), ($x_0)->{'value2'}, ($y_1)->{'value2'});
goto end_branch_9;;
};
$__t9 = new \Data\Ordering\Data_Ordering_LT();
end_branch_9:;
$__t1 = $__t9;
goto end_branch_1;;
};
  if (($y_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_August) {
$__t1 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_1;;
};
  if (($x_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_September) {
$__t10 = null;;
if (($y_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_September) {
$__t10 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), ($x_0)->{'value2'}, ($y_1)->{'value2'});
goto end_branch_10;;
};
$__t10 = new \Data\Ordering\Data_Ordering_LT();
end_branch_10:;
$__t1 = $__t10;
goto end_branch_1;;
};
  if (($y_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_September) {
$__t1 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_1;;
};
  if (($x_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_October) {
$__t11 = null;;
if (($y_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_October) {
$__t11 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), ($x_0)->{'value2'}, ($y_1)->{'value2'});
goto end_branch_11;;
};
$__t11 = new \Data\Ordering\Data_Ordering_LT();
end_branch_11:;
$__t1 = $__t11;
goto end_branch_1;;
};
  if (($y_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_October) {
$__t1 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_1;;
};
  if (($x_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_November) {
$__t12 = null;;
if (($y_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_November) {
$__t12 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), ($x_0)->{'value2'}, ($y_1)->{'value2'});
goto end_branch_12;;
};
$__t12 = new \Data\Ordering\Data_Ordering_LT();
end_branch_12:;
$__t1 = $__t12;
goto end_branch_1;;
};
  if (($y_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_November) {
$__t1 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_1;;
};
  if ((($x_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_December && ($y_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_December)) {
$__t1 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), ($x_0)->{'value2'}, ($y_1)->{'value2'});
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__t13 = $__t1;
  end_branch_13:;
  $__res = $__t13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Date_eqDate'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Date_enumDate
$GLOBALS['Data_Date_enumDate'] = (object)["succ" => function($v_0) {
  $__num = \func_num_args();
  $sm_1_0 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($v_1) {
  $__num = \func_num_args();
  $__res = match ($v_1) { 1 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_January()), 2 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_February()), 3 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_March()), 4 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_April()), 5 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_May()), 6 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_June()), 7 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_July()), 8 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_August()), 9 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_September()), 10 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_October()), 11 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_November()), 12 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_December()), default => new \Data\Maybe\Data_Maybe_Nothing() };
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_1) {
  $__num = \func_num_args();
  $__res = ($v_1 + 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_1) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v_1 instanceof \Data\Date\Component\Data_Date_Component_January) {
$__t1 = 1;
goto end_branch_1;;
};
  if ($v_1 instanceof \Data\Date\Component\Data_Date_Component_February) {
$__t1 = 2;
goto end_branch_1;;
};
  if ($v_1 instanceof \Data\Date\Component\Data_Date_Component_March) {
$__t1 = 3;
goto end_branch_1;;
};
  if ($v_1 instanceof \Data\Date\Component\Data_Date_Component_April) {
$__t1 = 4;
goto end_branch_1;;
};
  if ($v_1 instanceof \Data\Date\Component\Data_Date_Component_May) {
$__t1 = 5;
goto end_branch_1;;
};
  if ($v_1 instanceof \Data\Date\Component\Data_Date_Component_June) {
$__t1 = 6;
goto end_branch_1;;
};
  if ($v_1 instanceof \Data\Date\Component\Data_Date_Component_July) {
$__t1 = 7;
goto end_branch_1;;
};
  if ($v_1 instanceof \Data\Date\Component\Data_Date_Component_August) {
$__t1 = 8;
goto end_branch_1;;
};
  if ($v_1 instanceof \Data\Date\Component\Data_Date_Component_September) {
$__t1 = 9;
goto end_branch_1;;
};
  if ($v_1 instanceof \Data\Date\Component\Data_Date_Component_October) {
$__t1 = 10;
goto end_branch_1;;
};
  if ($v_1 instanceof \Data\Date\Component\Data_Date_Component_November) {
$__t1 = 11;
goto end_branch_1;;
};
  if ($v_1 instanceof \Data\Date\Component\Data_Date_Component_December) {
$__t1 = 12;
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), ($v_0)->{'value1'});
  $v1_2_3 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_2) {
  $__num = \func_num_args();
  $__t3 = null;;
  if ((($n_2 >= 1) && ($n_2 <= 31))) {
$__t3 = new \Data\Maybe\Data_Maybe_Just($n_2);
goto end_branch_3;;
};
  $__t3 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_3:;
  $__res = $__t3;
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
}), ($v_0)->{'value2'});
  $__t23 = null;;
  if ((function() use ($v1_2_3, $v_0, &$__fn) {
$__local_var_3_24 = new \Data\Maybe\Data_Maybe_Just(\Data\Date\majData_majDate_lastmajDaymajOfmajMonth(($v_0)->{'value0'}, ($v_0)->{'value1'}));
$__t25 = null;;
if ($v1_2_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t25 = false;
goto end_branch_25;;
};
if ($v1_2_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t25 = (($v1_2_3)->{'value0'} > ($__local_var_3_24)->{'value0'});
goto end_branch_25;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t25 = null;
end_branch_25:;
return $__t25;
})()) {
$__t26 = null;;
if ((function() use ($sm_1_0, &$__fn) {
$__t27 = null;;
if ($sm_1_0 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t27 = true;
goto end_branch_27;;
};
if ($sm_1_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t27 = false;
goto end_branch_27;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t27 = null;
end_branch_27:;
return $__t27;
})()) {
$__t26 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_3) {
  $__num = \func_num_args();
  $__t28 = null;;
  if ((($n_3 >= -271820) && ($n_3 <= 275759))) {
$__t28 = new \Data\Maybe\Data_Maybe_Just($n_3);
goto end_branch_28;;
};
  $__t28 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_28:;
  $__res = $__t28;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_3) {
  $__num = \func_num_args();
  $__res = ($v_3 + 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_3) {
  $__num = \func_num_args();
  $__res = $v_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), ($v_0)->{'value0'});
goto end_branch_26;;
};
$__t26 = new \Data\Maybe\Data_Maybe_Just(($v_0)->{'value0'});
end_branch_26:;
$__local_var_3_26 = $__t26;
$__t30 = null;;
if ($__local_var_3_26 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t30 = new \Data\Maybe\Data_Maybe_Just(($GLOBALS['Data_Date_Date'])(($__local_var_3_26)->{'value0'}));
goto end_branch_30;;
};
$__t30 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_30:;
$__local_var_3_26 = $__t30;
$__t32 = null;;
if ($sm_1_0 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t32 = new \Data\Date\Component\Data_Date_Component_January();
goto end_branch_32;;
};
if ($sm_1_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t32 = ($sm_1_0)->{'value0'};
goto end_branch_32;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t32 = null;
end_branch_32:;
$__local_var_4_32 = new \Data\Maybe\Data_Maybe_Just($__t32);
$__t34 = null;;
if ($__local_var_3_26 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t34 = new \Data\Maybe\Data_Maybe_Just((($__local_var_3_26)->{'value0'})(($__local_var_4_32)->{'value0'}));
goto end_branch_34;;
};
if ($__local_var_3_26 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t34 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_34;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t34 = null;
end_branch_34:;
$__local_var_3_26 = $__t34;
$__t36 = null;;
if ($__local_var_3_26 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t36 = new \Data\Maybe\Data_Maybe_Just((($__local_var_3_26)->{'value0'})(1));
goto end_branch_36;;
};
if ($__local_var_3_26 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t36 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_36;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t36 = null;
end_branch_36:;
$__t23 = $__t36;
goto end_branch_23;;
};
  $__t5 = null;;
  if ((function() use ($sm_1_0, $v1_2_3, &$__fn) {
$__t6 = null;;
if ($v1_2_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t6 = true;
goto end_branch_6;;
};
if ($v1_2_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t6 = false;
goto end_branch_6;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t6 = null;
end_branch_6:;
$__t7 = null;;
if ($sm_1_0 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t7 = true;
goto end_branch_7;;
};
if ($sm_1_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t7 = false;
goto end_branch_7;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t7 = null;
end_branch_7:;
return ($__t6 && $__t7);
})()) {
$__t5 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_3) {
  $__num = \func_num_args();
  $__t8 = null;;
  if ((($n_3 >= -271820) && ($n_3 <= 275759))) {
$__t8 = new \Data\Maybe\Data_Maybe_Just($n_3);
goto end_branch_8;;
};
  $__t8 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_8:;
  $__res = $__t8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_3) {
  $__num = \func_num_args();
  $__res = ($v_3 + 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_3) {
  $__num = \func_num_args();
  $__res = $v_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), ($v_0)->{'value0'});
goto end_branch_5;;
};
  $__t5 = new \Data\Maybe\Data_Maybe_Just(($v_0)->{'value0'});
  end_branch_5:;
  $__local_var_3_5 = $__t5;
  $__t10 = null;;
  if ($__local_var_3_5 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t10 = new \Data\Maybe\Data_Maybe_Just(($GLOBALS['Data_Date_Date'])(($__local_var_3_5)->{'value0'}));
goto end_branch_10;;
};
  $__t10 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_10:;
  $__local_var_3_5 = $__t10;
  $__t12 = null;;
  if ((function() use ($v1_2_3, &$__fn) {
$__t13 = null;;
if ($v1_2_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t13 = true;
goto end_branch_13;;
};
if ($v1_2_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t13 = false;
goto end_branch_13;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t13 = null;
end_branch_13:;
return $__t13;
})()) {
$__t14 = null;;
if ($sm_1_0 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t14 = new \Data\Date\Component\Data_Date_Component_January();
goto end_branch_14;;
};
if ($sm_1_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t14 = ($sm_1_0)->{'value0'};
goto end_branch_14;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t14 = null;
end_branch_14:;
$__t12 = $__t14;
goto end_branch_12;;
};
  $__t12 = ($v_0)->{'value1'};
  end_branch_12:;
  $__local_var_4_12 = new \Data\Maybe\Data_Maybe_Just($__t12);
  $__t16 = null;;
  if ($__local_var_3_5 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t16 = new \Data\Maybe\Data_Maybe_Just((($__local_var_3_5)->{'value0'})(($__local_var_4_12)->{'value0'}));
goto end_branch_16;;
};
  if ($__local_var_3_5 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t16 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_16;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t16 = null;
  end_branch_16:;
  $__local_var_3_5 = $__t16;
  $__t20 = null;;
  if ((function() use ($v1_2_3, &$__fn) {
$__t21 = null;;
if ($v1_2_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t21 = true;
goto end_branch_21;;
};
if ($v1_2_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t21 = false;
goto end_branch_21;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t21 = null;
end_branch_21:;
return $__t21;
})()) {
$__t22 = null;;
if ($__local_var_3_5 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t22 = new \Data\Maybe\Data_Maybe_Just((($__local_var_3_5)->{'value0'})(1));
goto end_branch_22;;
};
if ($__local_var_3_5 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t22 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_22;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t22 = null;
end_branch_22:;
$__t20 = $__t22;
goto end_branch_20;;
};
  $__t18 = null;;
  if ($__local_var_3_5 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t19 = null;;
if ($v1_2_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t19 = new \Data\Maybe\Data_Maybe_Just((($__local_var_3_5)->{'value0'})(($v1_2_3)->{'value0'}));
goto end_branch_19;;
};
$__t19 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_19:;
$__t18 = $__t19;
goto end_branch_18;;
};
  if ($__local_var_3_5 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t18 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_18;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t18 = null;
  end_branch_18:;
  $__t20 = $__t18;
  end_branch_20:;
  $__t23 = $__t20;
  end_branch_23:;
  $__res = $__t23;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "pred" => function($v_0) {
  $__num = \func_num_args();
  $pm_1_37 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($v_1) {
  $__num = \func_num_args();
  $__res = match ($v_1) { 1 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_January()), 2 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_February()), 3 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_March()), 4 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_April()), 5 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_May()), 6 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_June()), 7 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_July()), 8 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_August()), 9 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_September()), 10 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_October()), 11 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_November()), 12 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_December()), default => new \Data\Maybe\Data_Maybe_Nothing() };
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_1) {
  $__num = \func_num_args();
  $__res = ($v_1 - 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_1) {
  $__num = \func_num_args();
  $__t38 = null;;
  if ($v_1 instanceof \Data\Date\Component\Data_Date_Component_January) {
$__t38 = 1;
goto end_branch_38;;
};
  if ($v_1 instanceof \Data\Date\Component\Data_Date_Component_February) {
$__t38 = 2;
goto end_branch_38;;
};
  if ($v_1 instanceof \Data\Date\Component\Data_Date_Component_March) {
$__t38 = 3;
goto end_branch_38;;
};
  if ($v_1 instanceof \Data\Date\Component\Data_Date_Component_April) {
$__t38 = 4;
goto end_branch_38;;
};
  if ($v_1 instanceof \Data\Date\Component\Data_Date_Component_May) {
$__t38 = 5;
goto end_branch_38;;
};
  if ($v_1 instanceof \Data\Date\Component\Data_Date_Component_June) {
$__t38 = 6;
goto end_branch_38;;
};
  if ($v_1 instanceof \Data\Date\Component\Data_Date_Component_July) {
$__t38 = 7;
goto end_branch_38;;
};
  if ($v_1 instanceof \Data\Date\Component\Data_Date_Component_August) {
$__t38 = 8;
goto end_branch_38;;
};
  if ($v_1 instanceof \Data\Date\Component\Data_Date_Component_September) {
$__t38 = 9;
goto end_branch_38;;
};
  if ($v_1 instanceof \Data\Date\Component\Data_Date_Component_October) {
$__t38 = 10;
goto end_branch_38;;
};
  if ($v_1 instanceof \Data\Date\Component\Data_Date_Component_November) {
$__t38 = 11;
goto end_branch_38;;
};
  if ($v_1 instanceof \Data\Date\Component\Data_Date_Component_December) {
$__t38 = 12;
goto end_branch_38;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t38 = null;
  end_branch_38:;
  $__res = $__t38;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), ($v_0)->{'value1'});
  $pd_2_40 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_2) {
  $__num = \func_num_args();
  $__t40 = null;;
  if ((($n_2 >= 1) && ($n_2 <= 31))) {
$__t40 = new \Data\Maybe\Data_Maybe_Just($n_2);
goto end_branch_40;;
};
  $__t40 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_40:;
  $__res = $__t40;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_2) {
  $__num = \func_num_args();
  $__res = ($v_2 - 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_2) {
  $__num = \func_num_args();
  $__res = $v_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), ($v_0)->{'value2'});
  $__t42 = null;;
  if ((function() use ($pd_2_40, &$__fn) {
$__t43 = null;;
if ($pd_2_40 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t43 = true;
goto end_branch_43;;
};
if ($pd_2_40 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t43 = false;
goto end_branch_43;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t43 = null;
end_branch_43:;
return $__t43;
})()) {
$__t44 = null;;
if ($pm_1_37 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t44 = new \Data\Date\Component\Data_Date_Component_December();
goto end_branch_44;;
};
if ($pm_1_37 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t44 = ($pm_1_37)->{'value0'};
goto end_branch_44;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t44 = null;
end_branch_44:;
$__t42 = $__t44;
goto end_branch_42;;
};
  $__t42 = ($v_0)->{'value1'};
  end_branch_42:;
  $m_prime__3_42 = $__t42;
  $l_4_46 = \Data\Date\majData_majDate_lastmajDaymajOfmajMonth(($v_0)->{'value0'}, $m_prime__3_42);
  $__t47 = null;;
  if ((function() use ($pd_2_40, $pm_1_37, &$__fn) {
$__t48 = null;;
if ($pd_2_40 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t48 = true;
goto end_branch_48;;
};
if ($pd_2_40 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t48 = false;
goto end_branch_48;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t48 = null;
end_branch_48:;
$__t49 = null;;
if ($pm_1_37 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t49 = true;
goto end_branch_49;;
};
if ($pm_1_37 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t49 = false;
goto end_branch_49;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t49 = null;
end_branch_49:;
return ($__t48 && $__t49);
})()) {
$__t47 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_5) {
  $__num = \func_num_args();
  $__t50 = null;;
  if ((($n_5 >= -271820) && ($n_5 <= 275759))) {
$__t50 = new \Data\Maybe\Data_Maybe_Just($n_5);
goto end_branch_50;;
};
  $__t50 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_50:;
  $__res = $__t50;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_5) {
  $__num = \func_num_args();
  $__res = ($v_5 - 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_5) {
  $__num = \func_num_args();
  $__res = $v_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), ($v_0)->{'value0'});
goto end_branch_47;;
};
  $__t47 = new \Data\Maybe\Data_Maybe_Just(($v_0)->{'value0'});
  end_branch_47:;
  $__local_var_5_47 = $__t47;
  $__t52 = null;;
  if ($__local_var_5_47 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t52 = new \Data\Maybe\Data_Maybe_Just(($GLOBALS['Data_Date_Date'])(($__local_var_5_47)->{'value0'}));
goto end_branch_52;;
};
  $__t52 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_52:;
  $__local_var_5_47 = $__t52;
  $__t54 = null;;
  if ($__local_var_5_47 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t54 = new \Data\Maybe\Data_Maybe_Just((($__local_var_5_47)->{'value0'})($m_prime__3_42));
goto end_branch_54;;
};
  if ($__local_var_5_47 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t54 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_54;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t54 = null;
  end_branch_54:;
  $__local_var_5_47 = $__t54;
  $__t58 = null;;
  if ((function() use ($pd_2_40, &$__fn) {
$__t59 = null;;
if ($pd_2_40 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t59 = true;
goto end_branch_59;;
};
if ($pd_2_40 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t59 = false;
goto end_branch_59;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t59 = null;
end_branch_59:;
return $__t59;
})()) {
$__t60 = null;;
if ($__local_var_5_47 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t60 = new \Data\Maybe\Data_Maybe_Just((($__local_var_5_47)->{'value0'})($l_4_46));
goto end_branch_60;;
};
if ($__local_var_5_47 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t60 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_60;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t60 = null;
end_branch_60:;
$__t58 = $__t60;
goto end_branch_58;;
};
  $__t56 = null;;
  if ($__local_var_5_47 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t57 = null;;
if ($pd_2_40 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t57 = new \Data\Maybe\Data_Maybe_Just((($__local_var_5_47)->{'value0'})(($pd_2_40)->{'value0'}));
goto end_branch_57;;
};
$__t57 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_57:;
$__t56 = $__t57;
goto end_branch_56;;
};
  if ($__local_var_5_47 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t56 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_56;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t56 = null;
  end_branch_56:;
  $__t58 = $__t56;
  end_branch_58:;
  $__res = $__t58;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Ord0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Date_ordDate'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Date_diff
function majData_majDate_diff($dictDuration_0, $v_1 = null, $v1_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majDate_diff';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if (($v_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_January) {
$__t0 = 1;
goto end_branch_0;;
};
  if (($v_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_February) {
$__t0 = 2;
goto end_branch_0;;
};
  if (($v_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_March) {
$__t0 = 3;
goto end_branch_0;;
};
  if (($v_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_April) {
$__t0 = 4;
goto end_branch_0;;
};
  if (($v_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_May) {
$__t0 = 5;
goto end_branch_0;;
};
  if (($v_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_June) {
$__t0 = 6;
goto end_branch_0;;
};
  if (($v_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_July) {
$__t0 = 7;
goto end_branch_0;;
};
  if (($v_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_August) {
$__t0 = 8;
goto end_branch_0;;
};
  if (($v_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_September) {
$__t0 = 9;
goto end_branch_0;;
};
  if (($v_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_October) {
$__t0 = 10;
goto end_branch_0;;
};
  if (($v_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_November) {
$__t0 = 11;
goto end_branch_0;;
};
  if (($v_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_December) {
$__t0 = 12;
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__t1 = null;;
  if (($v1_2)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_January) {
$__t1 = 1;
goto end_branch_1;;
};
  if (($v1_2)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_February) {
$__t1 = 2;
goto end_branch_1;;
};
  if (($v1_2)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_March) {
$__t1 = 3;
goto end_branch_1;;
};
  if (($v1_2)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_April) {
$__t1 = 4;
goto end_branch_1;;
};
  if (($v1_2)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_May) {
$__t1 = 5;
goto end_branch_1;;
};
  if (($v1_2)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_June) {
$__t1 = 6;
goto end_branch_1;;
};
  if (($v1_2)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_July) {
$__t1 = 7;
goto end_branch_1;;
};
  if (($v1_2)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_August) {
$__t1 = 8;
goto end_branch_1;;
};
  if (($v1_2)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_September) {
$__t1 = 9;
goto end_branch_1;;
};
  if (($v1_2)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_October) {
$__t1 = 10;
goto end_branch_1;;
};
  if (($v1_2)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_November) {
$__t1 = 11;
goto end_branch_1;;
};
  if (($v1_2)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_December) {
$__t1 = 12;
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = (($dictDuration_0)->{'toDuration'})(($GLOBALS['Data_Date_calcDiff'])(($v_1)->{'value0'}, $__t0, ($v_1)->{'value2'}, ($v1_2)->{'value0'}, $__t1, ($v1_2)->{'value2'}));
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Date_diff'] = __NAMESPACE__ . '\\majData_majDate_diff';

// Data_Date_day
function majData_majDate_day($v_0): int|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majDate_day';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($v_0)->{'value2'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Date_day'] = __NAMESPACE__ . '\\majData_majDate_day';

// Data_Date_canonicalDate
function majData_majDate_canonicalmajDate(int $y_0, $m_1 = null, $d_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majDate_canonicalmajDate';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t1 = null;;
  if ($m_1 instanceof \Data\Date\Component\Data_Date_Component_January) {
$__t1 = 1;
goto end_branch_1;;
};
  if ($m_1 instanceof \Data\Date\Component\Data_Date_Component_February) {
$__t1 = 2;
goto end_branch_1;;
};
  if ($m_1 instanceof \Data\Date\Component\Data_Date_Component_March) {
$__t1 = 3;
goto end_branch_1;;
};
  if ($m_1 instanceof \Data\Date\Component\Data_Date_Component_April) {
$__t1 = 4;
goto end_branch_1;;
};
  if ($m_1 instanceof \Data\Date\Component\Data_Date_Component_May) {
$__t1 = 5;
goto end_branch_1;;
};
  if ($m_1 instanceof \Data\Date\Component\Data_Date_Component_June) {
$__t1 = 6;
goto end_branch_1;;
};
  if ($m_1 instanceof \Data\Date\Component\Data_Date_Component_July) {
$__t1 = 7;
goto end_branch_1;;
};
  if ($m_1 instanceof \Data\Date\Component\Data_Date_Component_August) {
$__t1 = 8;
goto end_branch_1;;
};
  if ($m_1 instanceof \Data\Date\Component\Data_Date_Component_September) {
$__t1 = 9;
goto end_branch_1;;
};
  if ($m_1 instanceof \Data\Date\Component\Data_Date_Component_October) {
$__t1 = 10;
goto end_branch_1;;
};
  if ($m_1 instanceof \Data\Date\Component\Data_Date_Component_November) {
$__t1 = 11;
goto end_branch_1;;
};
  if ($m_1 instanceof \Data\Date\Component\Data_Date_Component_December) {
$__t1 = 12;
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = ($GLOBALS['Data_Date_canonicalDateImpl'])(function($y_prime__3) {
  $__num = \func_num_args();
  $__res = function($m_prime__4) use ($y_prime__3) {
  $__num = \func_num_args();
  $__res = function($d_prime__5) use ($m_prime__4, $y_prime__3) {
  $__num = \func_num_args();
  $__t0 = null;;
  switch ($m_prime__4) {
case 1:
$__t0 = new \Data\Date\Component\Data_Date_Component_January();
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($m_prime__4) {
case 2:
$__t0 = new \Data\Date\Component\Data_Date_Component_February();
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($m_prime__4) {
case 3:
$__t0 = new \Data\Date\Component\Data_Date_Component_March();
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($m_prime__4) {
case 4:
$__t0 = new \Data\Date\Component\Data_Date_Component_April();
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($m_prime__4) {
case 5:
$__t0 = new \Data\Date\Component\Data_Date_Component_May();
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($m_prime__4) {
case 6:
$__t0 = new \Data\Date\Component\Data_Date_Component_June();
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($m_prime__4) {
case 7:
$__t0 = new \Data\Date\Component\Data_Date_Component_July();
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($m_prime__4) {
case 8:
$__t0 = new \Data\Date\Component\Data_Date_Component_August();
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($m_prime__4) {
case 9:
$__t0 = new \Data\Date\Component\Data_Date_Component_September();
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($m_prime__4) {
case 10:
$__t0 = new \Data\Date\Component\Data_Date_Component_October();
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($m_prime__4) {
case 11:
$__t0 = new \Data\Date\Component\Data_Date_Component_November();
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($m_prime__4) {
case 12:
$__t0 = new \Data\Date\Component\Data_Date_Component_December();
goto end_branch_0;;
break;
default:
;
break;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = new \Data\Date\Data_Date_Date($y_prime__3, $__t0, $d_prime__5);
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
}, $y_0, $__t1, $d_2);
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Date_canonicalDate'] = __NAMESPACE__ . '\\majData_majDate_canonicalmajDate';

// Data_Date_exactDate
function majData_majDate_exactmajDate(int $y_0, $m_1 = null, $d_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majDate_exactmajDate';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if ((function() use ($d_2, $m_1, $y_0, &$__fn) {
$__local_var_3_1 = \Data\Date\majData_majDate_canonicalmajDate($y_0, $m_1, $d_2);
$__t2 = null;;
if (($__local_var_3_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_January) {
$__t2 = $m_1 instanceof \Data\Date\Component\Data_Date_Component_January;
goto end_branch_2;;
};
if (($__local_var_3_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_February) {
$__t2 = $m_1 instanceof \Data\Date\Component\Data_Date_Component_February;
goto end_branch_2;;
};
if (($__local_var_3_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_March) {
$__t2 = $m_1 instanceof \Data\Date\Component\Data_Date_Component_March;
goto end_branch_2;;
};
if (($__local_var_3_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_April) {
$__t2 = $m_1 instanceof \Data\Date\Component\Data_Date_Component_April;
goto end_branch_2;;
};
if (($__local_var_3_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_May) {
$__t2 = $m_1 instanceof \Data\Date\Component\Data_Date_Component_May;
goto end_branch_2;;
};
if (($__local_var_3_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_June) {
$__t2 = $m_1 instanceof \Data\Date\Component\Data_Date_Component_June;
goto end_branch_2;;
};
if (($__local_var_3_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_July) {
$__t2 = $m_1 instanceof \Data\Date\Component\Data_Date_Component_July;
goto end_branch_2;;
};
if (($__local_var_3_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_August) {
$__t2 = $m_1 instanceof \Data\Date\Component\Data_Date_Component_August;
goto end_branch_2;;
};
if (($__local_var_3_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_September) {
$__t2 = $m_1 instanceof \Data\Date\Component\Data_Date_Component_September;
goto end_branch_2;;
};
if (($__local_var_3_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_October) {
$__t2 = $m_1 instanceof \Data\Date\Component\Data_Date_Component_October;
goto end_branch_2;;
};
if (($__local_var_3_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_November) {
$__t2 = $m_1 instanceof \Data\Date\Component\Data_Date_Component_November;
goto end_branch_2;;
};
$__t2 = (($__local_var_3_1)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_December && $m_1 instanceof \Data\Date\Component\Data_Date_Component_December);
end_branch_2:;
return (((($__local_var_3_1)->{'value0'} === $y_0) && $__t2) && (($__local_var_3_1)->{'value2'} === $d_2));
})()) {
$__t0 = new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Data_Date_Date($y_0, $m_1, $d_2));
goto end_branch_0;;
};
  $__t0 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Date_exactDate'] = __NAMESPACE__ . '\\majData_majDate_exactmajDate';

// Data_Date_boundedDate
$GLOBALS['Data_Date_boundedDate'] = (object)["bottom" => new \Data\Date\Data_Date_Date(-271820, new \Data\Date\Component\Data_Date_Component_January(), 1), "top" => new \Data\Date\Data_Date_Date(275759, new \Data\Date\Component\Data_Date_Component_December(), 31), "Ord0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Date_ordDate'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Date_adjust
function majData_majDate_adjust(float $v_0, $date_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majDate_adjust';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $adj_2_0 = null;
  $adj_2_0 = function($v1_3) use (&$adj_2_0) {
  $__num = \func_num_args();
  $__res = function($v2_4) use (&$adj_2_0, $v1_3) {
  $__num = \func_num_args();
  $__t211 = null;;
  switch ($v1_3) {
case 0:
$__t211 = new \Data\Maybe\Data_Maybe_Just($v2_4);
goto end_branch_211;;
break;
default:
;
break;
};
  $j_5_1 = ($v1_3 + ($v2_4)->{'value2'});
  $low_6_2 = ($j_5_1 < 1);
  $__t3 = null;;
  if ($low_6_2) {
$__local_var_7_4 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($v_7) {
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
  $__t5 = null;;
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_January) {
$__t5 = 1;
goto end_branch_5;;
};
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_February) {
$__t5 = 2;
goto end_branch_5;;
};
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_March) {
$__t5 = 3;
goto end_branch_5;;
};
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_April) {
$__t5 = 4;
goto end_branch_5;;
};
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_May) {
$__t5 = 5;
goto end_branch_5;;
};
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_June) {
$__t5 = 6;
goto end_branch_5;;
};
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_July) {
$__t5 = 7;
goto end_branch_5;;
};
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_August) {
$__t5 = 8;
goto end_branch_5;;
};
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_September) {
$__t5 = 9;
goto end_branch_5;;
};
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_October) {
$__t5 = 10;
goto end_branch_5;;
};
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_November) {
$__t5 = 11;
goto end_branch_5;;
};
  if ($v_7 instanceof \Data\Date\Component\Data_Date_Component_December) {
$__t5 = 12;
goto end_branch_5;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t5 = null;
  end_branch_5:;
  $__res = $__t5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), ($v2_4)->{'value1'});
$__t7 = null;;
if ($__local_var_7_4 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t7 = new \Data\Date\Component\Data_Date_Component_December();
goto end_branch_7;;
};
if ($__local_var_7_4 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t7 = ($__local_var_7_4)->{'value0'};
goto end_branch_7;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t7 = null;
end_branch_7:;
$__t3 = $__t7;
goto end_branch_3;;
};
  $__t3 = ($v2_4)->{'value1'};
  end_branch_3:;
  $l_7_3 = \Data\Date\majData_majDate_lastmajDaymajOfmajMonth(($v2_4)->{'value0'}, $__t3);
  $hi_8_9 = ($j_5_1 > $l_7_3);
  $__t10 = null;;
  if ($low_6_2) {
$__t10 = $j_5_1;
goto end_branch_10;;
};
  if ($hi_8_9) {
$__t10 = (($j_5_1 - $l_7_3) - 1);
goto end_branch_10;;
};
  $__t10 = 0;
  end_branch_10:;
  $__local_var_9_10 = ($adj_2_0)($__t10);
  $__t19 = null;;
  if ($low_6_2) {
$__t20 = null;;
if ((function() use ($v2_4, &$__fn) {
$__local_var_10_21 = new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Data_Date_Date(($v2_4)->{'value0'}, ($v2_4)->{'value1'}, 1));
$pm_11_22 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($v_11) {
  $__num = \func_num_args();
  $__res = match ($v_11) { 1 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_January()), 2 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_February()), 3 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_March()), 4 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_April()), 5 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_May()), 6 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_June()), 7 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_July()), 8 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_August()), 9 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_September()), 10 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_October()), 11 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_November()), 12 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_December()), default => new \Data\Maybe\Data_Maybe_Nothing() };
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_11) {
  $__num = \func_num_args();
  $__res = ($v_11 - 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_11) {
  $__num = \func_num_args();
  $__t23 = null;;
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_January) {
$__t23 = 1;
goto end_branch_23;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_February) {
$__t23 = 2;
goto end_branch_23;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_March) {
$__t23 = 3;
goto end_branch_23;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_April) {
$__t23 = 4;
goto end_branch_23;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_May) {
$__t23 = 5;
goto end_branch_23;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_June) {
$__t23 = 6;
goto end_branch_23;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_July) {
$__t23 = 7;
goto end_branch_23;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_August) {
$__t23 = 8;
goto end_branch_23;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_September) {
$__t23 = 9;
goto end_branch_23;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_October) {
$__t23 = 10;
goto end_branch_23;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_November) {
$__t23 = 11;
goto end_branch_23;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_December) {
$__t23 = 12;
goto end_branch_23;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t23 = null;
  end_branch_23:;
  $__res = $__t23;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), (($__local_var_10_21)->{'value0'})->{'value1'});
$pd_12_25 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_12) {
  $__num = \func_num_args();
  $__t25 = null;;
  if ((($n_12 >= 1) && ($n_12 <= 31))) {
$__t25 = new \Data\Maybe\Data_Maybe_Just($n_12);
goto end_branch_25;;
};
  $__t25 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_25:;
  $__res = $__t25;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_12) {
  $__num = \func_num_args();
  $__res = ($v_12 - 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_12) {
  $__num = \func_num_args();
  $__res = $v_12;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), (($__local_var_10_21)->{'value0'})->{'value2'});
$__t27 = null;;
if ((function() use ($pd_12_25, &$__fn) {
$__t28 = null;;
if ($pd_12_25 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t28 = true;
goto end_branch_28;;
};
if ($pd_12_25 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t28 = false;
goto end_branch_28;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t28 = null;
end_branch_28:;
return $__t28;
})()) {
$__t29 = null;;
if ($pm_11_22 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t29 = new \Data\Date\Component\Data_Date_Component_December();
goto end_branch_29;;
};
if ($pm_11_22 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t29 = ($pm_11_22)->{'value0'};
goto end_branch_29;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t29 = null;
end_branch_29:;
$__t27 = $__t29;
goto end_branch_27;;
};
$__t27 = (($__local_var_10_21)->{'value0'})->{'value1'};
end_branch_27:;
$m_prime__13_27 = $__t27;
$l_14_31 = \Data\Date\majData_majDate_lastmajDaymajOfmajMonth((($__local_var_10_21)->{'value0'})->{'value0'}, $m_prime__13_27);
$__t32 = null;;
if ((function() use ($pd_12_25, $pm_11_22, &$__fn) {
$__t33 = null;;
if ($pd_12_25 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t33 = true;
goto end_branch_33;;
};
if ($pd_12_25 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t33 = false;
goto end_branch_33;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t33 = null;
end_branch_33:;
$__t34 = null;;
if ($pm_11_22 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t34 = true;
goto end_branch_34;;
};
if ($pm_11_22 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t34 = false;
goto end_branch_34;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t34 = null;
end_branch_34:;
return ($__t33 && $__t34);
})()) {
$__t32 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_15) {
  $__num = \func_num_args();
  $__t35 = null;;
  if ((($n_15 >= -271820) && ($n_15 <= 275759))) {
$__t35 = new \Data\Maybe\Data_Maybe_Just($n_15);
goto end_branch_35;;
};
  $__t35 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_35:;
  $__res = $__t35;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_15) {
  $__num = \func_num_args();
  $__res = ($v_15 - 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_15) {
  $__num = \func_num_args();
  $__res = $v_15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), (($__local_var_10_21)->{'value0'})->{'value0'});
goto end_branch_32;;
};
$__t32 = new \Data\Maybe\Data_Maybe_Just((($__local_var_10_21)->{'value0'})->{'value0'});
end_branch_32:;
$__local_var_15_32 = $__t32;
$__t37 = null;;
if ($__local_var_15_32 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t37 = new \Data\Maybe\Data_Maybe_Just(($GLOBALS['Data_Date_Date'])(($__local_var_15_32)->{'value0'}));
goto end_branch_37;;
};
$__t37 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_37:;
$__local_var_15_32 = $__t37;
$__t39 = null;;
if ($__local_var_15_32 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t39 = new \Data\Maybe\Data_Maybe_Just((($__local_var_15_32)->{'value0'})($m_prime__13_27));
goto end_branch_39;;
};
if ($__local_var_15_32 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t39 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_39;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t39 = null;
end_branch_39:;
$__local_var_15_32 = $__t39;
$__t43 = null;;
if ((function() use ($pd_12_25, &$__fn) {
$__t44 = null;;
if ($pd_12_25 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t44 = true;
goto end_branch_44;;
};
if ($pd_12_25 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t44 = false;
goto end_branch_44;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t44 = null;
end_branch_44:;
return $__t44;
})()) {
$__t45 = null;;
if ($__local_var_15_32 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t45 = new \Data\Maybe\Data_Maybe_Just((($__local_var_15_32)->{'value0'})($l_14_31));
goto end_branch_45;;
};
if ($__local_var_15_32 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t45 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_45;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t45 = null;
end_branch_45:;
$__t43 = $__t45;
goto end_branch_43;;
};
$__t41 = null;;
if ($__local_var_15_32 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t42 = null;;
if ($pd_12_25 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t42 = new \Data\Maybe\Data_Maybe_Just((($__local_var_15_32)->{'value0'})(($pd_12_25)->{'value0'}));
goto end_branch_42;;
};
$__t42 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_42:;
$__t41 = $__t42;
goto end_branch_41;;
};
if ($__local_var_15_32 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t41 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_41;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t41 = null;
end_branch_41:;
$__t43 = $__t41;
end_branch_43:;
return $__t43 instanceof \Data\Maybe\Data_Maybe_Just;
})()) {
$__local_var_10_46 = new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Data_Date_Date(($v2_4)->{'value0'}, ($v2_4)->{'value1'}, 1));
$pm_11_47 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($v_11) {
  $__num = \func_num_args();
  $__res = match ($v_11) { 1 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_January()), 2 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_February()), 3 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_March()), 4 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_April()), 5 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_May()), 6 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_June()), 7 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_July()), 8 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_August()), 9 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_September()), 10 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_October()), 11 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_November()), 12 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_December()), default => new \Data\Maybe\Data_Maybe_Nothing() };
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_11) {
  $__num = \func_num_args();
  $__res = ($v_11 - 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_11) {
  $__num = \func_num_args();
  $__t48 = null;;
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_January) {
$__t48 = 1;
goto end_branch_48;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_February) {
$__t48 = 2;
goto end_branch_48;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_March) {
$__t48 = 3;
goto end_branch_48;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_April) {
$__t48 = 4;
goto end_branch_48;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_May) {
$__t48 = 5;
goto end_branch_48;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_June) {
$__t48 = 6;
goto end_branch_48;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_July) {
$__t48 = 7;
goto end_branch_48;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_August) {
$__t48 = 8;
goto end_branch_48;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_September) {
$__t48 = 9;
goto end_branch_48;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_October) {
$__t48 = 10;
goto end_branch_48;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_November) {
$__t48 = 11;
goto end_branch_48;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_December) {
$__t48 = 12;
goto end_branch_48;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t48 = null;
  end_branch_48:;
  $__res = $__t48;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), (($__local_var_10_46)->{'value0'})->{'value1'});
$pd_12_50 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_12) {
  $__num = \func_num_args();
  $__t50 = null;;
  if ((($n_12 >= 1) && ($n_12 <= 31))) {
$__t50 = new \Data\Maybe\Data_Maybe_Just($n_12);
goto end_branch_50;;
};
  $__t50 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_50:;
  $__res = $__t50;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_12) {
  $__num = \func_num_args();
  $__res = ($v_12 - 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_12) {
  $__num = \func_num_args();
  $__res = $v_12;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), (($__local_var_10_46)->{'value0'})->{'value2'});
$__t52 = null;;
if ((function() use ($pd_12_50, &$__fn) {
$__t53 = null;;
if ($pd_12_50 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t53 = true;
goto end_branch_53;;
};
if ($pd_12_50 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t53 = false;
goto end_branch_53;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t53 = null;
end_branch_53:;
return $__t53;
})()) {
$__t54 = null;;
if ($pm_11_47 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t54 = new \Data\Date\Component\Data_Date_Component_December();
goto end_branch_54;;
};
if ($pm_11_47 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t54 = ($pm_11_47)->{'value0'};
goto end_branch_54;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t54 = null;
end_branch_54:;
$__t52 = $__t54;
goto end_branch_52;;
};
$__t52 = (($__local_var_10_46)->{'value0'})->{'value1'};
end_branch_52:;
$m_prime__13_52 = $__t52;
$l_14_56 = \Data\Date\majData_majDate_lastmajDaymajOfmajMonth((($__local_var_10_46)->{'value0'})->{'value0'}, $m_prime__13_52);
$__t57 = null;;
if ((function() use ($pd_12_50, $pm_11_47, &$__fn) {
$__t58 = null;;
if ($pd_12_50 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t58 = true;
goto end_branch_58;;
};
if ($pd_12_50 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t58 = false;
goto end_branch_58;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t58 = null;
end_branch_58:;
$__t59 = null;;
if ($pm_11_47 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t59 = true;
goto end_branch_59;;
};
if ($pm_11_47 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t59 = false;
goto end_branch_59;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t59 = null;
end_branch_59:;
return ($__t58 && $__t59);
})()) {
$__t57 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_15) {
  $__num = \func_num_args();
  $__t60 = null;;
  if ((($n_15 >= -271820) && ($n_15 <= 275759))) {
$__t60 = new \Data\Maybe\Data_Maybe_Just($n_15);
goto end_branch_60;;
};
  $__t60 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_60:;
  $__res = $__t60;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_15) {
  $__num = \func_num_args();
  $__res = ($v_15 - 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_15) {
  $__num = \func_num_args();
  $__res = $v_15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), (($__local_var_10_46)->{'value0'})->{'value0'});
goto end_branch_57;;
};
$__t57 = new \Data\Maybe\Data_Maybe_Just((($__local_var_10_46)->{'value0'})->{'value0'});
end_branch_57:;
$__local_var_15_57 = $__t57;
$__t62 = null;;
if ($__local_var_15_57 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t62 = new \Data\Maybe\Data_Maybe_Just(($GLOBALS['Data_Date_Date'])(($__local_var_15_57)->{'value0'}));
goto end_branch_62;;
};
$__t62 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_62:;
$__local_var_15_57 = $__t62;
$__t64 = null;;
if ($__local_var_15_57 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t64 = new \Data\Maybe\Data_Maybe_Just((($__local_var_15_57)->{'value0'})($m_prime__13_52));
goto end_branch_64;;
};
if ($__local_var_15_57 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t64 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_64;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t64 = null;
end_branch_64:;
$__local_var_15_57 = $__t64;
$__t68 = null;;
if ((function() use ($pd_12_50, &$__fn) {
$__t69 = null;;
if ($pd_12_50 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t69 = true;
goto end_branch_69;;
};
if ($pd_12_50 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t69 = false;
goto end_branch_69;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t69 = null;
end_branch_69:;
return $__t69;
})()) {
$__t70 = null;;
if ($__local_var_15_57 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t70 = new \Data\Maybe\Data_Maybe_Just((($__local_var_15_57)->{'value0'})($l_14_56));
goto end_branch_70;;
};
if ($__local_var_15_57 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t70 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_70;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t70 = null;
end_branch_70:;
$__t68 = $__t70;
goto end_branch_68;;
};
$__t66 = null;;
if ($__local_var_15_57 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t67 = null;;
if ($pd_12_50 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t67 = new \Data\Maybe\Data_Maybe_Just((($__local_var_15_57)->{'value0'})(($pd_12_50)->{'value0'}));
goto end_branch_67;;
};
$__t67 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_67:;
$__t66 = $__t67;
goto end_branch_66;;
};
if ($__local_var_15_57 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t66 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_66;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t66 = null;
end_branch_66:;
$__t68 = $__t66;
end_branch_68:;
$__t20 = ($__local_var_9_10)(($__t68)->{'value0'});
goto end_branch_20;;
};
if ((function() use ($v2_4, &$__fn) {
$__local_var_10_71 = new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Data_Date_Date(($v2_4)->{'value0'}, ($v2_4)->{'value1'}, 1));
$pm_11_72 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($v_11) {
  $__num = \func_num_args();
  $__res = match ($v_11) { 1 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_January()), 2 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_February()), 3 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_March()), 4 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_April()), 5 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_May()), 6 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_June()), 7 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_July()), 8 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_August()), 9 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_September()), 10 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_October()), 11 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_November()), 12 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_December()), default => new \Data\Maybe\Data_Maybe_Nothing() };
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_11) {
  $__num = \func_num_args();
  $__res = ($v_11 - 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_11) {
  $__num = \func_num_args();
  $__t73 = null;;
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_January) {
$__t73 = 1;
goto end_branch_73;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_February) {
$__t73 = 2;
goto end_branch_73;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_March) {
$__t73 = 3;
goto end_branch_73;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_April) {
$__t73 = 4;
goto end_branch_73;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_May) {
$__t73 = 5;
goto end_branch_73;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_June) {
$__t73 = 6;
goto end_branch_73;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_July) {
$__t73 = 7;
goto end_branch_73;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_August) {
$__t73 = 8;
goto end_branch_73;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_September) {
$__t73 = 9;
goto end_branch_73;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_October) {
$__t73 = 10;
goto end_branch_73;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_November) {
$__t73 = 11;
goto end_branch_73;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_December) {
$__t73 = 12;
goto end_branch_73;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t73 = null;
  end_branch_73:;
  $__res = $__t73;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), (($__local_var_10_71)->{'value0'})->{'value1'});
$pd_12_75 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_12) {
  $__num = \func_num_args();
  $__t75 = null;;
  if ((($n_12 >= 1) && ($n_12 <= 31))) {
$__t75 = new \Data\Maybe\Data_Maybe_Just($n_12);
goto end_branch_75;;
};
  $__t75 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_75:;
  $__res = $__t75;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_12) {
  $__num = \func_num_args();
  $__res = ($v_12 - 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_12) {
  $__num = \func_num_args();
  $__res = $v_12;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), (($__local_var_10_71)->{'value0'})->{'value2'});
$__t77 = null;;
if ((function() use ($pd_12_75, &$__fn) {
$__t78 = null;;
if ($pd_12_75 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t78 = true;
goto end_branch_78;;
};
if ($pd_12_75 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t78 = false;
goto end_branch_78;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t78 = null;
end_branch_78:;
return $__t78;
})()) {
$__t79 = null;;
if ($pm_11_72 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t79 = new \Data\Date\Component\Data_Date_Component_December();
goto end_branch_79;;
};
if ($pm_11_72 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t79 = ($pm_11_72)->{'value0'};
goto end_branch_79;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t79 = null;
end_branch_79:;
$__t77 = $__t79;
goto end_branch_77;;
};
$__t77 = (($__local_var_10_71)->{'value0'})->{'value1'};
end_branch_77:;
$m_prime__13_77 = $__t77;
$l_14_81 = \Data\Date\majData_majDate_lastmajDaymajOfmajMonth((($__local_var_10_71)->{'value0'})->{'value0'}, $m_prime__13_77);
$__t82 = null;;
if ((function() use ($pd_12_75, $pm_11_72, &$__fn) {
$__t83 = null;;
if ($pd_12_75 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t83 = true;
goto end_branch_83;;
};
if ($pd_12_75 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t83 = false;
goto end_branch_83;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t83 = null;
end_branch_83:;
$__t84 = null;;
if ($pm_11_72 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t84 = true;
goto end_branch_84;;
};
if ($pm_11_72 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t84 = false;
goto end_branch_84;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t84 = null;
end_branch_84:;
return ($__t83 && $__t84);
})()) {
$__t82 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_15) {
  $__num = \func_num_args();
  $__t85 = null;;
  if ((($n_15 >= -271820) && ($n_15 <= 275759))) {
$__t85 = new \Data\Maybe\Data_Maybe_Just($n_15);
goto end_branch_85;;
};
  $__t85 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_85:;
  $__res = $__t85;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_15) {
  $__num = \func_num_args();
  $__res = ($v_15 - 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_15) {
  $__num = \func_num_args();
  $__res = $v_15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), (($__local_var_10_71)->{'value0'})->{'value0'});
goto end_branch_82;;
};
$__t82 = new \Data\Maybe\Data_Maybe_Just((($__local_var_10_71)->{'value0'})->{'value0'});
end_branch_82:;
$__local_var_15_82 = $__t82;
$__t87 = null;;
if ($__local_var_15_82 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t87 = new \Data\Maybe\Data_Maybe_Just(($GLOBALS['Data_Date_Date'])(($__local_var_15_82)->{'value0'}));
goto end_branch_87;;
};
$__t87 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_87:;
$__local_var_15_82 = $__t87;
$__t89 = null;;
if ($__local_var_15_82 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t89 = new \Data\Maybe\Data_Maybe_Just((($__local_var_15_82)->{'value0'})($m_prime__13_77));
goto end_branch_89;;
};
if ($__local_var_15_82 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t89 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_89;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t89 = null;
end_branch_89:;
$__local_var_15_82 = $__t89;
$__t93 = null;;
if ((function() use ($pd_12_75, &$__fn) {
$__t94 = null;;
if ($pd_12_75 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t94 = true;
goto end_branch_94;;
};
if ($pd_12_75 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t94 = false;
goto end_branch_94;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t94 = null;
end_branch_94:;
return $__t94;
})()) {
$__t95 = null;;
if ($__local_var_15_82 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t95 = new \Data\Maybe\Data_Maybe_Just((($__local_var_15_82)->{'value0'})($l_14_81));
goto end_branch_95;;
};
if ($__local_var_15_82 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t95 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_95;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t95 = null;
end_branch_95:;
$__t93 = $__t95;
goto end_branch_93;;
};
$__t91 = null;;
if ($__local_var_15_82 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t92 = null;;
if ($pd_12_75 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t92 = new \Data\Maybe\Data_Maybe_Just((($__local_var_15_82)->{'value0'})(($pd_12_75)->{'value0'}));
goto end_branch_92;;
};
$__t92 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_92:;
$__t91 = $__t92;
goto end_branch_91;;
};
if ($__local_var_15_82 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t91 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_91;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t91 = null;
end_branch_91:;
$__t93 = $__t91;
end_branch_93:;
return $__t93 instanceof \Data\Maybe\Data_Maybe_Nothing;
})()) {
$__t20 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_20;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t20 = null;
end_branch_20:;
$__t19 = $__t20;
goto end_branch_19;;
};
  if ($hi_8_9) {
$__t96 = null;;
if ((function() use ($l_7_3, $v2_4, &$__fn) {
$__local_var_10_97 = new \Data\Date\Data_Date_Date(($v2_4)->{'value0'}, ($v2_4)->{'value1'}, $l_7_3);
$sm_11_98 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($v_11) {
  $__num = \func_num_args();
  $__res = match ($v_11) { 1 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_January()), 2 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_February()), 3 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_March()), 4 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_April()), 5 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_May()), 6 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_June()), 7 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_July()), 8 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_August()), 9 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_September()), 10 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_October()), 11 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_November()), 12 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_December()), default => new \Data\Maybe\Data_Maybe_Nothing() };
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_11) {
  $__num = \func_num_args();
  $__res = ($v_11 + 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_11) {
  $__num = \func_num_args();
  $__t99 = null;;
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_January) {
$__t99 = 1;
goto end_branch_99;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_February) {
$__t99 = 2;
goto end_branch_99;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_March) {
$__t99 = 3;
goto end_branch_99;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_April) {
$__t99 = 4;
goto end_branch_99;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_May) {
$__t99 = 5;
goto end_branch_99;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_June) {
$__t99 = 6;
goto end_branch_99;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_July) {
$__t99 = 7;
goto end_branch_99;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_August) {
$__t99 = 8;
goto end_branch_99;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_September) {
$__t99 = 9;
goto end_branch_99;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_October) {
$__t99 = 10;
goto end_branch_99;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_November) {
$__t99 = 11;
goto end_branch_99;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_December) {
$__t99 = 12;
goto end_branch_99;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t99 = null;
  end_branch_99:;
  $__res = $__t99;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), ($__local_var_10_97)->{'value1'});
$v1_12_101 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_12) {
  $__num = \func_num_args();
  $__t101 = null;;
  if ((($n_12 >= 1) && ($n_12 <= 31))) {
$__t101 = new \Data\Maybe\Data_Maybe_Just($n_12);
goto end_branch_101;;
};
  $__t101 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_101:;
  $__res = $__t101;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_12) {
  $__num = \func_num_args();
  $__res = ($v_12 + 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_12) {
  $__num = \func_num_args();
  $__res = $v_12;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), ($__local_var_10_97)->{'value2'});
$__t121 = null;;
if ((function() use ($__local_var_10_97, $v1_12_101, &$__fn) {
$__local_var_13_122 = new \Data\Maybe\Data_Maybe_Just(\Data\Date\majData_majDate_lastmajDaymajOfmajMonth(($__local_var_10_97)->{'value0'}, ($__local_var_10_97)->{'value1'}));
$__t123 = null;;
if ($v1_12_101 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t123 = false;
goto end_branch_123;;
};
if ($v1_12_101 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t123 = (($v1_12_101)->{'value0'} > ($__local_var_13_122)->{'value0'});
goto end_branch_123;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t123 = null;
end_branch_123:;
return $__t123;
})()) {
$__t124 = null;;
if ((function() use ($sm_11_98, &$__fn) {
$__t125 = null;;
if ($sm_11_98 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t125 = true;
goto end_branch_125;;
};
if ($sm_11_98 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t125 = false;
goto end_branch_125;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t125 = null;
end_branch_125:;
return $__t125;
})()) {
$__t124 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_13) {
  $__num = \func_num_args();
  $__t126 = null;;
  if ((($n_13 >= -271820) && ($n_13 <= 275759))) {
$__t126 = new \Data\Maybe\Data_Maybe_Just($n_13);
goto end_branch_126;;
};
  $__t126 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_126:;
  $__res = $__t126;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_13) {
  $__num = \func_num_args();
  $__res = ($v_13 + 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_13) {
  $__num = \func_num_args();
  $__res = $v_13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), ($__local_var_10_97)->{'value0'});
goto end_branch_124;;
};
$__t124 = new \Data\Maybe\Data_Maybe_Just(($__local_var_10_97)->{'value0'});
end_branch_124:;
$__local_var_13_124 = $__t124;
$__t128 = null;;
if ($__local_var_13_124 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t128 = new \Data\Maybe\Data_Maybe_Just(($GLOBALS['Data_Date_Date'])(($__local_var_13_124)->{'value0'}));
goto end_branch_128;;
};
$__t128 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_128:;
$__local_var_13_124 = $__t128;
$__t130 = null;;
if ($sm_11_98 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t130 = new \Data\Date\Component\Data_Date_Component_January();
goto end_branch_130;;
};
if ($sm_11_98 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t130 = ($sm_11_98)->{'value0'};
goto end_branch_130;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t130 = null;
end_branch_130:;
$__local_var_14_130 = new \Data\Maybe\Data_Maybe_Just($__t130);
$__t132 = null;;
if ($__local_var_13_124 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t132 = new \Data\Maybe\Data_Maybe_Just((($__local_var_13_124)->{'value0'})(($__local_var_14_130)->{'value0'}));
goto end_branch_132;;
};
if ($__local_var_13_124 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t132 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_132;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t132 = null;
end_branch_132:;
$__local_var_13_124 = $__t132;
$__t134 = null;;
if ($__local_var_13_124 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t134 = new \Data\Maybe\Data_Maybe_Just((($__local_var_13_124)->{'value0'})(1));
goto end_branch_134;;
};
if ($__local_var_13_124 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t134 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_134;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t134 = null;
end_branch_134:;
$__t121 = $__t134;
goto end_branch_121;;
};
$__t103 = null;;
if ((function() use ($sm_11_98, $v1_12_101, &$__fn) {
$__t104 = null;;
if ($v1_12_101 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t104 = true;
goto end_branch_104;;
};
if ($v1_12_101 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t104 = false;
goto end_branch_104;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t104 = null;
end_branch_104:;
$__t105 = null;;
if ($sm_11_98 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t105 = true;
goto end_branch_105;;
};
if ($sm_11_98 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t105 = false;
goto end_branch_105;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t105 = null;
end_branch_105:;
return ($__t104 && $__t105);
})()) {
$__t103 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_13) {
  $__num = \func_num_args();
  $__t106 = null;;
  if ((($n_13 >= -271820) && ($n_13 <= 275759))) {
$__t106 = new \Data\Maybe\Data_Maybe_Just($n_13);
goto end_branch_106;;
};
  $__t106 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_106:;
  $__res = $__t106;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_13) {
  $__num = \func_num_args();
  $__res = ($v_13 + 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_13) {
  $__num = \func_num_args();
  $__res = $v_13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), ($__local_var_10_97)->{'value0'});
goto end_branch_103;;
};
$__t103 = new \Data\Maybe\Data_Maybe_Just(($__local_var_10_97)->{'value0'});
end_branch_103:;
$__local_var_13_103 = $__t103;
$__t108 = null;;
if ($__local_var_13_103 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t108 = new \Data\Maybe\Data_Maybe_Just(($GLOBALS['Data_Date_Date'])(($__local_var_13_103)->{'value0'}));
goto end_branch_108;;
};
$__t108 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_108:;
$__local_var_13_103 = $__t108;
$__t110 = null;;
if ((function() use ($v1_12_101, &$__fn) {
$__t111 = null;;
if ($v1_12_101 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t111 = true;
goto end_branch_111;;
};
if ($v1_12_101 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t111 = false;
goto end_branch_111;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t111 = null;
end_branch_111:;
return $__t111;
})()) {
$__t112 = null;;
if ($sm_11_98 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t112 = new \Data\Date\Component\Data_Date_Component_January();
goto end_branch_112;;
};
if ($sm_11_98 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t112 = ($sm_11_98)->{'value0'};
goto end_branch_112;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t112 = null;
end_branch_112:;
$__t110 = $__t112;
goto end_branch_110;;
};
$__t110 = ($__local_var_10_97)->{'value1'};
end_branch_110:;
$__local_var_14_110 = new \Data\Maybe\Data_Maybe_Just($__t110);
$__t114 = null;;
if ($__local_var_13_103 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t114 = new \Data\Maybe\Data_Maybe_Just((($__local_var_13_103)->{'value0'})(($__local_var_14_110)->{'value0'}));
goto end_branch_114;;
};
if ($__local_var_13_103 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t114 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_114;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t114 = null;
end_branch_114:;
$__local_var_13_103 = $__t114;
$__t118 = null;;
if ((function() use ($v1_12_101, &$__fn) {
$__t119 = null;;
if ($v1_12_101 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t119 = true;
goto end_branch_119;;
};
if ($v1_12_101 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t119 = false;
goto end_branch_119;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t119 = null;
end_branch_119:;
return $__t119;
})()) {
$__t120 = null;;
if ($__local_var_13_103 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t120 = new \Data\Maybe\Data_Maybe_Just((($__local_var_13_103)->{'value0'})(1));
goto end_branch_120;;
};
if ($__local_var_13_103 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t120 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_120;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t120 = null;
end_branch_120:;
$__t118 = $__t120;
goto end_branch_118;;
};
$__t116 = null;;
if ($__local_var_13_103 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t117 = null;;
if ($v1_12_101 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t117 = new \Data\Maybe\Data_Maybe_Just((($__local_var_13_103)->{'value0'})(($v1_12_101)->{'value0'}));
goto end_branch_117;;
};
$__t117 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_117:;
$__t116 = $__t117;
goto end_branch_116;;
};
if ($__local_var_13_103 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t116 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_116;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t116 = null;
end_branch_116:;
$__t118 = $__t116;
end_branch_118:;
$__t121 = $__t118;
end_branch_121:;
return $__t121 instanceof \Data\Maybe\Data_Maybe_Just;
})()) {
$__local_var_10_135 = new \Data\Date\Data_Date_Date(($v2_4)->{'value0'}, ($v2_4)->{'value1'}, $l_7_3);
$sm_11_136 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($v_11) {
  $__num = \func_num_args();
  $__res = match ($v_11) { 1 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_January()), 2 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_February()), 3 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_March()), 4 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_April()), 5 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_May()), 6 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_June()), 7 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_July()), 8 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_August()), 9 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_September()), 10 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_October()), 11 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_November()), 12 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_December()), default => new \Data\Maybe\Data_Maybe_Nothing() };
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_11) {
  $__num = \func_num_args();
  $__res = ($v_11 + 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_11) {
  $__num = \func_num_args();
  $__t137 = null;;
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_January) {
$__t137 = 1;
goto end_branch_137;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_February) {
$__t137 = 2;
goto end_branch_137;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_March) {
$__t137 = 3;
goto end_branch_137;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_April) {
$__t137 = 4;
goto end_branch_137;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_May) {
$__t137 = 5;
goto end_branch_137;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_June) {
$__t137 = 6;
goto end_branch_137;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_July) {
$__t137 = 7;
goto end_branch_137;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_August) {
$__t137 = 8;
goto end_branch_137;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_September) {
$__t137 = 9;
goto end_branch_137;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_October) {
$__t137 = 10;
goto end_branch_137;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_November) {
$__t137 = 11;
goto end_branch_137;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_December) {
$__t137 = 12;
goto end_branch_137;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t137 = null;
  end_branch_137:;
  $__res = $__t137;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), ($__local_var_10_135)->{'value1'});
$v1_12_139 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_12) {
  $__num = \func_num_args();
  $__t139 = null;;
  if ((($n_12 >= 1) && ($n_12 <= 31))) {
$__t139 = new \Data\Maybe\Data_Maybe_Just($n_12);
goto end_branch_139;;
};
  $__t139 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_139:;
  $__res = $__t139;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_12) {
  $__num = \func_num_args();
  $__res = ($v_12 + 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_12) {
  $__num = \func_num_args();
  $__res = $v_12;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), ($__local_var_10_135)->{'value2'});
$__t159 = null;;
if ((function() use ($__local_var_10_135, $v1_12_139, &$__fn) {
$__local_var_13_160 = new \Data\Maybe\Data_Maybe_Just(\Data\Date\majData_majDate_lastmajDaymajOfmajMonth(($__local_var_10_135)->{'value0'}, ($__local_var_10_135)->{'value1'}));
$__t161 = null;;
if ($v1_12_139 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t161 = false;
goto end_branch_161;;
};
if ($v1_12_139 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t161 = (($v1_12_139)->{'value0'} > ($__local_var_13_160)->{'value0'});
goto end_branch_161;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t161 = null;
end_branch_161:;
return $__t161;
})()) {
$__t162 = null;;
if ((function() use ($sm_11_136, &$__fn) {
$__t163 = null;;
if ($sm_11_136 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t163 = true;
goto end_branch_163;;
};
if ($sm_11_136 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t163 = false;
goto end_branch_163;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t163 = null;
end_branch_163:;
return $__t163;
})()) {
$__t162 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_13) {
  $__num = \func_num_args();
  $__t164 = null;;
  if ((($n_13 >= -271820) && ($n_13 <= 275759))) {
$__t164 = new \Data\Maybe\Data_Maybe_Just($n_13);
goto end_branch_164;;
};
  $__t164 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_164:;
  $__res = $__t164;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_13) {
  $__num = \func_num_args();
  $__res = ($v_13 + 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_13) {
  $__num = \func_num_args();
  $__res = $v_13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), ($__local_var_10_135)->{'value0'});
goto end_branch_162;;
};
$__t162 = new \Data\Maybe\Data_Maybe_Just(($__local_var_10_135)->{'value0'});
end_branch_162:;
$__local_var_13_162 = $__t162;
$__t166 = null;;
if ($__local_var_13_162 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t166 = new \Data\Maybe\Data_Maybe_Just(($GLOBALS['Data_Date_Date'])(($__local_var_13_162)->{'value0'}));
goto end_branch_166;;
};
$__t166 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_166:;
$__local_var_13_162 = $__t166;
$__t168 = null;;
if ($sm_11_136 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t168 = new \Data\Date\Component\Data_Date_Component_January();
goto end_branch_168;;
};
if ($sm_11_136 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t168 = ($sm_11_136)->{'value0'};
goto end_branch_168;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t168 = null;
end_branch_168:;
$__local_var_14_168 = new \Data\Maybe\Data_Maybe_Just($__t168);
$__t170 = null;;
if ($__local_var_13_162 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t170 = new \Data\Maybe\Data_Maybe_Just((($__local_var_13_162)->{'value0'})(($__local_var_14_168)->{'value0'}));
goto end_branch_170;;
};
if ($__local_var_13_162 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t170 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_170;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t170 = null;
end_branch_170:;
$__local_var_13_162 = $__t170;
$__t172 = null;;
if ($__local_var_13_162 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t172 = new \Data\Maybe\Data_Maybe_Just((($__local_var_13_162)->{'value0'})(1));
goto end_branch_172;;
};
if ($__local_var_13_162 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t172 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_172;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t172 = null;
end_branch_172:;
$__t159 = $__t172;
goto end_branch_159;;
};
$__t141 = null;;
if ((function() use ($sm_11_136, $v1_12_139, &$__fn) {
$__t142 = null;;
if ($v1_12_139 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t142 = true;
goto end_branch_142;;
};
if ($v1_12_139 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t142 = false;
goto end_branch_142;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t142 = null;
end_branch_142:;
$__t143 = null;;
if ($sm_11_136 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t143 = true;
goto end_branch_143;;
};
if ($sm_11_136 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t143 = false;
goto end_branch_143;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t143 = null;
end_branch_143:;
return ($__t142 && $__t143);
})()) {
$__t141 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_13) {
  $__num = \func_num_args();
  $__t144 = null;;
  if ((($n_13 >= -271820) && ($n_13 <= 275759))) {
$__t144 = new \Data\Maybe\Data_Maybe_Just($n_13);
goto end_branch_144;;
};
  $__t144 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_144:;
  $__res = $__t144;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_13) {
  $__num = \func_num_args();
  $__res = ($v_13 + 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_13) {
  $__num = \func_num_args();
  $__res = $v_13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), ($__local_var_10_135)->{'value0'});
goto end_branch_141;;
};
$__t141 = new \Data\Maybe\Data_Maybe_Just(($__local_var_10_135)->{'value0'});
end_branch_141:;
$__local_var_13_141 = $__t141;
$__t146 = null;;
if ($__local_var_13_141 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t146 = new \Data\Maybe\Data_Maybe_Just(($GLOBALS['Data_Date_Date'])(($__local_var_13_141)->{'value0'}));
goto end_branch_146;;
};
$__t146 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_146:;
$__local_var_13_141 = $__t146;
$__t148 = null;;
if ((function() use ($v1_12_139, &$__fn) {
$__t149 = null;;
if ($v1_12_139 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t149 = true;
goto end_branch_149;;
};
if ($v1_12_139 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t149 = false;
goto end_branch_149;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t149 = null;
end_branch_149:;
return $__t149;
})()) {
$__t150 = null;;
if ($sm_11_136 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t150 = new \Data\Date\Component\Data_Date_Component_January();
goto end_branch_150;;
};
if ($sm_11_136 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t150 = ($sm_11_136)->{'value0'};
goto end_branch_150;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t150 = null;
end_branch_150:;
$__t148 = $__t150;
goto end_branch_148;;
};
$__t148 = ($__local_var_10_135)->{'value1'};
end_branch_148:;
$__local_var_14_148 = new \Data\Maybe\Data_Maybe_Just($__t148);
$__t152 = null;;
if ($__local_var_13_141 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t152 = new \Data\Maybe\Data_Maybe_Just((($__local_var_13_141)->{'value0'})(($__local_var_14_148)->{'value0'}));
goto end_branch_152;;
};
if ($__local_var_13_141 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t152 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_152;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t152 = null;
end_branch_152:;
$__local_var_13_141 = $__t152;
$__t156 = null;;
if ((function() use ($v1_12_139, &$__fn) {
$__t157 = null;;
if ($v1_12_139 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t157 = true;
goto end_branch_157;;
};
if ($v1_12_139 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t157 = false;
goto end_branch_157;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t157 = null;
end_branch_157:;
return $__t157;
})()) {
$__t158 = null;;
if ($__local_var_13_141 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t158 = new \Data\Maybe\Data_Maybe_Just((($__local_var_13_141)->{'value0'})(1));
goto end_branch_158;;
};
if ($__local_var_13_141 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t158 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_158;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t158 = null;
end_branch_158:;
$__t156 = $__t158;
goto end_branch_156;;
};
$__t154 = null;;
if ($__local_var_13_141 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t155 = null;;
if ($v1_12_139 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t155 = new \Data\Maybe\Data_Maybe_Just((($__local_var_13_141)->{'value0'})(($v1_12_139)->{'value0'}));
goto end_branch_155;;
};
$__t155 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_155:;
$__t154 = $__t155;
goto end_branch_154;;
};
if ($__local_var_13_141 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t154 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_154;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t154 = null;
end_branch_154:;
$__t156 = $__t154;
end_branch_156:;
$__t159 = $__t156;
end_branch_159:;
$__t96 = ($__local_var_9_10)(($__t159)->{'value0'});
goto end_branch_96;;
};
if ((function() use ($l_7_3, $v2_4, &$__fn) {
$__local_var_10_173 = new \Data\Date\Data_Date_Date(($v2_4)->{'value0'}, ($v2_4)->{'value1'}, $l_7_3);
$sm_11_174 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($v_11) {
  $__num = \func_num_args();
  $__res = match ($v_11) { 1 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_January()), 2 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_February()), 3 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_March()), 4 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_April()), 5 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_May()), 6 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_June()), 7 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_July()), 8 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_August()), 9 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_September()), 10 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_October()), 11 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_November()), 12 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_December()), default => new \Data\Maybe\Data_Maybe_Nothing() };
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_11) {
  $__num = \func_num_args();
  $__res = ($v_11 + 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_11) {
  $__num = \func_num_args();
  $__t175 = null;;
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_January) {
$__t175 = 1;
goto end_branch_175;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_February) {
$__t175 = 2;
goto end_branch_175;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_March) {
$__t175 = 3;
goto end_branch_175;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_April) {
$__t175 = 4;
goto end_branch_175;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_May) {
$__t175 = 5;
goto end_branch_175;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_June) {
$__t175 = 6;
goto end_branch_175;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_July) {
$__t175 = 7;
goto end_branch_175;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_August) {
$__t175 = 8;
goto end_branch_175;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_September) {
$__t175 = 9;
goto end_branch_175;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_October) {
$__t175 = 10;
goto end_branch_175;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_November) {
$__t175 = 11;
goto end_branch_175;;
};
  if ($v_11 instanceof \Data\Date\Component\Data_Date_Component_December) {
$__t175 = 12;
goto end_branch_175;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t175 = null;
  end_branch_175:;
  $__res = $__t175;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), ($__local_var_10_173)->{'value1'});
$v1_12_177 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_12) {
  $__num = \func_num_args();
  $__t177 = null;;
  if ((($n_12 >= 1) && ($n_12 <= 31))) {
$__t177 = new \Data\Maybe\Data_Maybe_Just($n_12);
goto end_branch_177;;
};
  $__t177 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_177:;
  $__res = $__t177;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_12) {
  $__num = \func_num_args();
  $__res = ($v_12 + 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_12) {
  $__num = \func_num_args();
  $__res = $v_12;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), ($__local_var_10_173)->{'value2'});
$__t197 = null;;
if ((function() use ($__local_var_10_173, $v1_12_177, &$__fn) {
$__local_var_13_198 = new \Data\Maybe\Data_Maybe_Just(\Data\Date\majData_majDate_lastmajDaymajOfmajMonth(($__local_var_10_173)->{'value0'}, ($__local_var_10_173)->{'value1'}));
$__t199 = null;;
if ($v1_12_177 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t199 = false;
goto end_branch_199;;
};
if ($v1_12_177 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t199 = (($v1_12_177)->{'value0'} > ($__local_var_13_198)->{'value0'});
goto end_branch_199;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t199 = null;
end_branch_199:;
return $__t199;
})()) {
$__t200 = null;;
if ((function() use ($sm_11_174, &$__fn) {
$__t201 = null;;
if ($sm_11_174 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t201 = true;
goto end_branch_201;;
};
if ($sm_11_174 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t201 = false;
goto end_branch_201;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t201 = null;
end_branch_201:;
return $__t201;
})()) {
$__t200 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_13) {
  $__num = \func_num_args();
  $__t202 = null;;
  if ((($n_13 >= -271820) && ($n_13 <= 275759))) {
$__t202 = new \Data\Maybe\Data_Maybe_Just($n_13);
goto end_branch_202;;
};
  $__t202 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_202:;
  $__res = $__t202;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_13) {
  $__num = \func_num_args();
  $__res = ($v_13 + 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_13) {
  $__num = \func_num_args();
  $__res = $v_13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), ($__local_var_10_173)->{'value0'});
goto end_branch_200;;
};
$__t200 = new \Data\Maybe\Data_Maybe_Just(($__local_var_10_173)->{'value0'});
end_branch_200:;
$__local_var_13_200 = $__t200;
$__t204 = null;;
if ($__local_var_13_200 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t204 = new \Data\Maybe\Data_Maybe_Just(($GLOBALS['Data_Date_Date'])(($__local_var_13_200)->{'value0'}));
goto end_branch_204;;
};
$__t204 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_204:;
$__local_var_13_200 = $__t204;
$__t206 = null;;
if ($sm_11_174 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t206 = new \Data\Date\Component\Data_Date_Component_January();
goto end_branch_206;;
};
if ($sm_11_174 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t206 = ($sm_11_174)->{'value0'};
goto end_branch_206;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t206 = null;
end_branch_206:;
$__local_var_14_206 = new \Data\Maybe\Data_Maybe_Just($__t206);
$__t208 = null;;
if ($__local_var_13_200 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t208 = new \Data\Maybe\Data_Maybe_Just((($__local_var_13_200)->{'value0'})(($__local_var_14_206)->{'value0'}));
goto end_branch_208;;
};
if ($__local_var_13_200 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t208 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_208;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t208 = null;
end_branch_208:;
$__local_var_13_200 = $__t208;
$__t210 = null;;
if ($__local_var_13_200 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t210 = new \Data\Maybe\Data_Maybe_Just((($__local_var_13_200)->{'value0'})(1));
goto end_branch_210;;
};
if ($__local_var_13_200 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t210 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_210;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t210 = null;
end_branch_210:;
$__t197 = $__t210;
goto end_branch_197;;
};
$__t179 = null;;
if ((function() use ($sm_11_174, $v1_12_177, &$__fn) {
$__t180 = null;;
if ($v1_12_177 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t180 = true;
goto end_branch_180;;
};
if ($v1_12_177 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t180 = false;
goto end_branch_180;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t180 = null;
end_branch_180:;
$__t181 = null;;
if ($sm_11_174 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t181 = true;
goto end_branch_181;;
};
if ($sm_11_174 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t181 = false;
goto end_branch_181;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t181 = null;
end_branch_181:;
return ($__t180 && $__t181);
})()) {
$__t179 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($n_13) {
  $__num = \func_num_args();
  $__t182 = null;;
  if ((($n_13 >= -271820) && ($n_13 <= 275759))) {
$__t182 = new \Data\Maybe\Data_Maybe_Just($n_13);
goto end_branch_182;;
};
  $__t182 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_182:;
  $__res = $__t182;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_13) {
  $__num = \func_num_args();
  $__res = ($v_13 + 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_13) {
  $__num = \func_num_args();
  $__res = $v_13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), ($__local_var_10_173)->{'value0'});
goto end_branch_179;;
};
$__t179 = new \Data\Maybe\Data_Maybe_Just(($__local_var_10_173)->{'value0'});
end_branch_179:;
$__local_var_13_179 = $__t179;
$__t184 = null;;
if ($__local_var_13_179 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t184 = new \Data\Maybe\Data_Maybe_Just(($GLOBALS['Data_Date_Date'])(($__local_var_13_179)->{'value0'}));
goto end_branch_184;;
};
$__t184 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_184:;
$__local_var_13_179 = $__t184;
$__t186 = null;;
if ((function() use ($v1_12_177, &$__fn) {
$__t187 = null;;
if ($v1_12_177 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t187 = true;
goto end_branch_187;;
};
if ($v1_12_177 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t187 = false;
goto end_branch_187;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t187 = null;
end_branch_187:;
return $__t187;
})()) {
$__t188 = null;;
if ($sm_11_174 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t188 = new \Data\Date\Component\Data_Date_Component_January();
goto end_branch_188;;
};
if ($sm_11_174 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t188 = ($sm_11_174)->{'value0'};
goto end_branch_188;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t188 = null;
end_branch_188:;
$__t186 = $__t188;
goto end_branch_186;;
};
$__t186 = ($__local_var_10_173)->{'value1'};
end_branch_186:;
$__local_var_14_186 = new \Data\Maybe\Data_Maybe_Just($__t186);
$__t190 = null;;
if ($__local_var_13_179 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t190 = new \Data\Maybe\Data_Maybe_Just((($__local_var_13_179)->{'value0'})(($__local_var_14_186)->{'value0'}));
goto end_branch_190;;
};
if ($__local_var_13_179 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t190 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_190;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t190 = null;
end_branch_190:;
$__local_var_13_179 = $__t190;
$__t194 = null;;
if ((function() use ($v1_12_177, &$__fn) {
$__t195 = null;;
if ($v1_12_177 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t195 = true;
goto end_branch_195;;
};
if ($v1_12_177 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t195 = false;
goto end_branch_195;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t195 = null;
end_branch_195:;
return $__t195;
})()) {
$__t196 = null;;
if ($__local_var_13_179 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t196 = new \Data\Maybe\Data_Maybe_Just((($__local_var_13_179)->{'value0'})(1));
goto end_branch_196;;
};
if ($__local_var_13_179 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t196 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_196;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t196 = null;
end_branch_196:;
$__t194 = $__t196;
goto end_branch_194;;
};
$__t192 = null;;
if ($__local_var_13_179 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t193 = null;;
if ($v1_12_177 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t193 = new \Data\Maybe\Data_Maybe_Just((($__local_var_13_179)->{'value0'})(($v1_12_177)->{'value0'}));
goto end_branch_193;;
};
$__t193 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_193:;
$__t192 = $__t193;
goto end_branch_192;;
};
if ($__local_var_13_179 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t192 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_192;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t192 = null;
end_branch_192:;
$__t194 = $__t192;
end_branch_194:;
$__t197 = $__t194;
end_branch_197:;
return $__t197 instanceof \Data\Maybe\Data_Maybe_Nothing;
})()) {
$__t96 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_96;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t96 = null;
end_branch_96:;
$__t19 = $__t96;
goto end_branch_19;;
};
  $__t12 = null;;
  if ((function() use ($j_5_1, $v2_4, &$__fn) {
$__local_var_10_13 = (($GLOBALS['Data_Date_Date'])(($v2_4)->{'value0'}))(($v2_4)->{'value1'});
$__t14 = null;;
if ((($j_5_1 >= 1) && ($j_5_1 <= 31))) {
$__t14 = new \Data\Maybe\Data_Maybe_Just(($__local_var_10_13)($j_5_1));
goto end_branch_14;;
};
$__t14 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_14:;
return $__t14 instanceof \Data\Maybe\Data_Maybe_Just;
})()) {
$__local_var_10_15 = (($GLOBALS['Data_Date_Date'])(($v2_4)->{'value0'}))(($v2_4)->{'value1'});
$__t16 = null;;
if ((($j_5_1 >= 1) && ($j_5_1 <= 31))) {
$__t16 = ($__local_var_10_15)($j_5_1);
goto end_branch_16;;
};
$__t16 = (new \Data\Maybe\Data_Maybe_Nothing())->{'value0'};
end_branch_16:;
$__t12 = ($__local_var_9_10)($__t16);
goto end_branch_12;;
};
  if ((function() use ($j_5_1, $v2_4, &$__fn) {
$__local_var_10_17 = (($GLOBALS['Data_Date_Date'])(($v2_4)->{'value0'}))(($v2_4)->{'value1'});
$__t18 = null;;
if ((($j_5_1 >= 1) && ($j_5_1 <= 31))) {
$__t18 = new \Data\Maybe\Data_Maybe_Just(($__local_var_10_17)($j_5_1));
goto end_branch_18;;
};
$__t18 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_18:;
return $__t18 instanceof \Data\Maybe\Data_Maybe_Nothing;
})()) {
$__t12 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_12;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t12 = null;
  end_branch_12:;
  $__t19 = $__t12;
  end_branch_19:;
  $__t211 = $__t19;
  end_branch_211:;
  $__res = $__t211;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__local_var_3_212 = \Data\Int\majData_majInt_frommajNumber($v_0);
  $__t213 = null;;
  if ($__local_var_3_212 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t213 = (($adj_2_0)(($__local_var_3_212)->{'value0'}))($date_1);
goto end_branch_213;;
};
  if ($__local_var_3_212 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t213 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_213;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t213 = null;
  end_branch_213:;
  $__res = $__t213;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Date_adjust'] = __NAMESPACE__ . '\\majData_majDate_adjust';

