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

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };
$ffi_Data_Date = \call_user_func(function() {
  $exports = [];
$createDate = function($y, $m, $d) {
    $dt = new \DateTime('now', new \DateTimeZone('UTC'));
    $dt->setDate($y, $m + 1, $d);
    $dt->setTime(0, 0, 0, 0);
    return $dt;
};

$canonicalDateImpl = function($ctor, $y = null, $m = null, $d = null) use (&$canonicalDateImpl, $createDate) {
    if (\func_num_args() < 4) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$canonicalDateImpl) {

            return $canonicalDateImpl(...\array_merge($__args, $more));
        };
    }

    $date = $createDate($y, $m - 1, $d);
    return $ctor
        ((int)$date->format('Y'))
        ((int)$date->format('n'))
        ((int)$date->format('j'));
};

$calcWeekday = function($y, $m = null, $d = null) use (&$calcWeekday, $createDate) {
    if (\func_num_args() < 3) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$calcWeekday) {

            return $calcWeekday(...\array_merge($__args, $more));
        };
    }

    $date = $createDate($y, $m - 1, $d);
    return (int)$date->format('w'); // 0 (for Sunday) through 6 (for Saturday)
};

$calcDiff = function($y1, $m1 = null, $d1 = null, $y2 = null, $m2 = null, $d2 = null) use (&$calcDiff, $createDate) {
    if (\func_num_args() < 6) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$calcDiff) {

            return $calcDiff(...\array_merge($__args, $more));
        };
    }

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
$GLOBALS['Data_Date_calcDiff'] = $ffi_Data_Date['calcDiff'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_Date_calcWeekday'] = $ffi_Data_Date['calcWeekday'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_Date_canonicalDateImpl'] = $ffi_Data_Date['canonicalDateImpl'] ?? new class { public function __invoke(...$args) { return $this; } };


// Data_Date_fromJust
$GLOBALS['Data_Date_fromJust'] = function($v_0 = null) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Just"))) {
$__t0 = ($v_0)->{'value0'};
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

// Data_Date_Date
$GLOBALS['Data_Date_Date'] = (function() {
  $__fn = function($value0 = null, $value1 = null, $value2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = new Phpurs_Data3("Date", $value0, $value1, $value2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_Date_year
$GLOBALS['Data_Date_year'] = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ($v_0)->{'value0'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Date_weekday
$GLOBALS['Data_Date_weekday'] = function($v_0 = null) {
  $__num = \func_num_args();
  $n_1_0 = ($GLOBALS['Data_Date_calcWeekday'])(($v_0)->{'value0'}, (($GLOBALS['Data_Date_Component_boundedEnumMonth'])['fromEnum'])(($v_0)->{'value1'}), ($v_0)->{'value2'});
  $__t3 = null;;
  if (((($GLOBALS['Data_Eq_eqInt'])['eq'])($n_1_0))(0)) {
$__local_var_2_4 = (($GLOBALS['Data_Date_Component_boundedEnumWeekday'])['toEnum'])(7);
$__t5 = null;;
if ((is_object($__local_var_2_4) && (($__local_var_2_4)->{'tag'} === "Just"))) {
$__t5 = ($__local_var_2_4)->{'value0'};
goto end_branch_5;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t5 = null;
end_branch_5:;
$__t3 = $__t5;
goto end_branch_3;;
};
  $__local_var_2_1 = (($GLOBALS['Data_Date_Component_boundedEnumWeekday'])['toEnum'])($n_1_0);
  $__t2 = null;;
  if ((is_object($__local_var_2_1) && (($__local_var_2_1)->{'tag'} === "Just"))) {
$__t2 = ($__local_var_2_1)->{'value0'};
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__t3 = $__t2;
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Date_showDate
$GLOBALS['Data_Date_showDate'] = ["show" => function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Semigroup_semigroupString'])['append'])("(Date "))(((($GLOBALS['Data_Semigroup_semigroupString'])['append'])((($GLOBALS['Data_Date_Component_showYear'])['show'])(($v_0)->{'value0'})))(((($GLOBALS['Data_Semigroup_semigroupString'])['append'])(" "))(((($GLOBALS['Data_Semigroup_semigroupString'])['append'])((($GLOBALS['Data_Date_Component_showMonth'])['show'])(($v_0)->{'value1'})))(((($GLOBALS['Data_Semigroup_semigroupString'])['append'])(" "))(((($GLOBALS['Data_Semigroup_semigroupString'])['append'])((($GLOBALS['Data_Date_Component_showDay'])['show'])(($v_0)->{'value2'})))(")"))))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Date_month
$GLOBALS['Data_Date_month'] = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ($v_0)->{'value1'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Date_isLeapYear
$GLOBALS['Data_Date_isLeapYear'] = function($y_0 = null) {
  $__num = \func_num_args();
  $y_prime_1_0 = (($GLOBALS['Data_Date_Component_boundedEnumYear'])['fromEnum'])($y_0);
  $__res = ((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])['conj'])(((($GLOBALS['Data_Eq_eqInt'])['eq'])(((($GLOBALS['Data_EuclideanRing_euclideanRingInt'])['mod'])($y_prime_1_0))(4)))(0)))(((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])['disj'])(((($GLOBALS['Data_Eq_eqInt'])['eq'])(((($GLOBALS['Data_EuclideanRing_euclideanRingInt'])['mod'])($y_prime_1_0))(400)))(0)))((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])['not'])(((($GLOBALS['Data_Eq_eqInt'])['eq'])(((($GLOBALS['Data_EuclideanRing_euclideanRingInt'])['mod'])($y_prime_1_0))(100)))(0))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Date_lastDayOfMonth
$GLOBALS['Data_Date_lastDayOfMonth'] = (function() {
  $__fn = function($y_0 = null, $m_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $unsafeDay_2_0 = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])($GLOBALS['Data_Date_fromJust']))(($GLOBALS['Data_Date_Component_boundedEnumDay'])['toEnum']);
  $__t1 = null;;
  if ((is_object($m_1) && (($m_1)->{'tag'} === "January"))) {
$__t1 = ($unsafeDay_2_0)(31);
goto end_branch_1;;
};
  if ((is_object($m_1) && (($m_1)->{'tag'} === "February"))) {
$__t2 = null;;
if (($GLOBALS['Data_Date_isLeapYear'])($y_0)) {
$__t2 = ($unsafeDay_2_0)(29);
goto end_branch_2;;
};
$__t2 = ($unsafeDay_2_0)(28);
end_branch_2:;
$__t1 = $__t2;
goto end_branch_1;;
};
  if ((is_object($m_1) && (($m_1)->{'tag'} === "March"))) {
$__t1 = ($unsafeDay_2_0)(31);
goto end_branch_1;;
};
  if ((is_object($m_1) && (($m_1)->{'tag'} === "April"))) {
$__t1 = ($unsafeDay_2_0)(30);
goto end_branch_1;;
};
  if ((is_object($m_1) && (($m_1)->{'tag'} === "May"))) {
$__t1 = ($unsafeDay_2_0)(31);
goto end_branch_1;;
};
  if ((is_object($m_1) && (($m_1)->{'tag'} === "June"))) {
$__t1 = ($unsafeDay_2_0)(30);
goto end_branch_1;;
};
  if ((is_object($m_1) && (($m_1)->{'tag'} === "July"))) {
$__t1 = ($unsafeDay_2_0)(31);
goto end_branch_1;;
};
  if ((is_object($m_1) && (($m_1)->{'tag'} === "August"))) {
$__t1 = ($unsafeDay_2_0)(31);
goto end_branch_1;;
};
  if ((is_object($m_1) && (($m_1)->{'tag'} === "September"))) {
$__t1 = ($unsafeDay_2_0)(30);
goto end_branch_1;;
};
  if ((is_object($m_1) && (($m_1)->{'tag'} === "October"))) {
$__t1 = ($unsafeDay_2_0)(31);
goto end_branch_1;;
};
  if ((is_object($m_1) && (($m_1)->{'tag'} === "November"))) {
$__t1 = ($unsafeDay_2_0)(30);
goto end_branch_1;;
};
  if ((is_object($m_1) && (($m_1)->{'tag'} === "December"))) {
$__t1 = ($unsafeDay_2_0)(31);
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

// Data_Date_eqDate
$GLOBALS['Data_Date_eqDate'] = ["eq" => (function() {
  $__fn = function($x_0 = null, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])['conj'])(((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])['conj'])(((($GLOBALS['Data_Eq_eqInt'])['eq'])(($x_0)->{'value0'}))(($y_1)->{'value0'})))(((($GLOBALS['Data_Date_Component_eqMonth'])['eq'])(($x_0)->{'value1'}))(($y_1)->{'value1'}))))(((($GLOBALS['Data_Eq_eqInt'])['eq'])(($x_0)->{'value2'}))(($y_1)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];

// Data_Date_ordDate
$GLOBALS['Data_Date_ordDate'] = ["compare" => (function() {
  $__fn = function($x_0 = null, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v_2_0 = ((($GLOBALS['Data_Ord_ordInt'])['compare'])(($x_0)->{'value0'}))(($y_1)->{'value0'});
  $__t3 = null;;
  if ((is_object($v_2_0) && (($v_2_0)->{'tag'} === "LT"))) {
$__t3 = new Phpurs_Data0("LT");
goto end_branch_3;;
};
  if ((is_object($v_2_0) && (($v_2_0)->{'tag'} === "GT"))) {
$__t3 = new Phpurs_Data0("GT");
goto end_branch_3;;
};
  $v1_3_1 = ((($GLOBALS['Data_Date_Component_ordMonth'])['compare'])(($x_0)->{'value1'}))(($y_1)->{'value1'});
  $__t2 = null;;
  if ((is_object($v1_3_1) && (($v1_3_1)->{'tag'} === "LT"))) {
$__t2 = new Phpurs_Data0("LT");
goto end_branch_2;;
};
  if ((is_object($v1_3_1) && (($v1_3_1)->{'tag'} === "GT"))) {
$__t2 = new Phpurs_Data0("GT");
goto end_branch_2;;
};
  $__t2 = ((($GLOBALS['Data_Ord_ordInt'])['compare'])(($x_0)->{'value2'}))(($y_1)->{'value2'});
  end_branch_2:;
  $__t3 = $__t2;
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Date_eqDate'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Date_enumDate
$GLOBALS['Data_Date_enumDate'] = ["succ" => function($v_0 = null) {
  $__num = \func_num_args();
  $sm_1_0 = (($GLOBALS['Data_Date_Component_enumMonth'])['succ'])(($v_0)->{'value1'});
  $v1_2_1 = (($GLOBALS['Data_Date_Component_enumDay'])['succ'])(($v_0)->{'value2'});
  $__t2 = null;;
  if ((function() use ($v1_2_1, $v_0, &$__fn) {
$__local_var_3_3 = (($GLOBALS['Data_Date_lastDayOfMonth'])(($v_0)->{'value0'}))(($v_0)->{'value1'});
$__t4 = null;;
if ((is_object($v1_2_1) && (($v1_2_1)->{'tag'} === "Nothing"))) {
$__t4 = false;
goto end_branch_4;;
};
if ((is_object($v1_2_1) && (($v1_2_1)->{'tag'} === "Just"))) {
$__t4 = (($v1_2_1)->{'value0'} > $__local_var_3_3);
goto end_branch_4;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t4 = null;
end_branch_4:;
return $__t4;
})()) {
$__t2 = new Phpurs_Data0("Nothing");
goto end_branch_2;;
};
  $__t2 = $v1_2_1;
  end_branch_2:;
  $sd_3_2 = $__t2;
  $__t6 = null;;
  if ((function() use ($sd_3_2, $sm_1_0, &$__fn) {
$__t7 = null;;
if ((is_object($sd_3_2) && (($sd_3_2)->{'tag'} === "Nothing"))) {
$__t7 = true;
goto end_branch_7;;
};
if ((is_object($sd_3_2) && (($sd_3_2)->{'tag'} === "Just"))) {
$__t7 = false;
goto end_branch_7;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t7 = null;
end_branch_7:;
$__t8 = null;;
if ((is_object($sm_1_0) && (($sm_1_0)->{'tag'} === "Nothing"))) {
$__t8 = true;
goto end_branch_8;;
};
if ((is_object($sm_1_0) && (($sm_1_0)->{'tag'} === "Just"))) {
$__t8 = false;
goto end_branch_8;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t8 = null;
end_branch_8:;
return ((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])['conj'])($__t7))($__t8);
})()) {
$__t6 = (($GLOBALS['Data_Date_Component_enumYear'])['succ'])(($v_0)->{'value0'});
goto end_branch_6;;
};
  $__t6 = new Phpurs_Data1("Just", ($v_0)->{'value0'});
  end_branch_6:;
  $__t9 = null;;
  if ((function() use ($sd_3_2, &$__fn) {
$__t10 = null;;
if ((is_object($sd_3_2) && (($sd_3_2)->{'tag'} === "Nothing"))) {
$__t10 = true;
goto end_branch_10;;
};
if ((is_object($sd_3_2) && (($sd_3_2)->{'tag'} === "Just"))) {
$__t10 = false;
goto end_branch_10;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t10 = null;
end_branch_10:;
return $__t10;
})()) {
$__t11 = null;;
if ((is_object($sm_1_0) && (($sm_1_0)->{'tag'} === "Nothing"))) {
$__t11 = new Phpurs_Data0("January");
goto end_branch_11;;
};
if ((is_object($sm_1_0) && (($sm_1_0)->{'tag'} === "Just"))) {
$__t11 = (($GLOBALS['Control_Category_categoryFn'])['identity'])(($sm_1_0)->{'value0'});
goto end_branch_11;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t11 = null;
end_branch_11:;
$__t9 = $__t11;
goto end_branch_9;;
};
  $__t9 = ($v_0)->{'value1'};
  end_branch_9:;
  $__t12 = null;;
  if ((function() use ($sd_3_2, &$__fn) {
$__t13 = null;;
if ((is_object($sd_3_2) && (($sd_3_2)->{'tag'} === "Nothing"))) {
$__t13 = true;
goto end_branch_13;;
};
if ((is_object($sd_3_2) && (($sd_3_2)->{'tag'} === "Just"))) {
$__t13 = false;
goto end_branch_13;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t13 = null;
end_branch_13:;
return $__t13;
})()) {
$__t12 = (($GLOBALS['Data_Date_Component_boundedEnumDay'])['toEnum'])(1);
goto end_branch_12;;
};
  $__t12 = $sd_3_2;
  end_branch_12:;
  $__res = ((($GLOBALS['Data_Maybe_applyMaybe'])['apply'])(((($GLOBALS['Data_Maybe_applyMaybe'])['apply'])(((($GLOBALS['Data_Maybe_functorMaybe'])['map'])($GLOBALS['Data_Date_Date']))($__t6)))((($GLOBALS['Data_Maybe_applicativeMaybe'])['pure'])($__t9))))($__t12);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "pred" => function($v_0 = null) {
  $__num = \func_num_args();
  $pm_1_14 = (($GLOBALS['Data_Date_Component_enumMonth'])['pred'])(($v_0)->{'value1'});
  $pd_2_15 = (($GLOBALS['Data_Date_Component_enumDay'])['pred'])(($v_0)->{'value2'});
  $__t16 = null;;
  if ((function() use ($pd_2_15, &$__fn) {
$__t17 = null;;
if ((is_object($pd_2_15) && (($pd_2_15)->{'tag'} === "Nothing"))) {
$__t17 = true;
goto end_branch_17;;
};
if ((is_object($pd_2_15) && (($pd_2_15)->{'tag'} === "Just"))) {
$__t17 = false;
goto end_branch_17;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t17 = null;
end_branch_17:;
return $__t17;
})()) {
$__t18 = null;;
if ((is_object($pm_1_14) && (($pm_1_14)->{'tag'} === "Nothing"))) {
$__t18 = new Phpurs_Data0("December");
goto end_branch_18;;
};
if ((is_object($pm_1_14) && (($pm_1_14)->{'tag'} === "Just"))) {
$__t18 = (($GLOBALS['Control_Category_categoryFn'])['identity'])(($pm_1_14)->{'value0'});
goto end_branch_18;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t18 = null;
end_branch_18:;
$__t16 = $__t18;
goto end_branch_16;;
};
  $__t16 = ($v_0)->{'value1'};
  end_branch_16:;
  $m_prime_3_16 = $__t16;
  $l_4_20 = (($GLOBALS['Data_Date_lastDayOfMonth'])(($v_0)->{'value0'}))($m_prime_3_16);
  $__t21 = null;;
  if ((function() use ($pd_2_15, $pm_1_14, &$__fn) {
$__t22 = null;;
if ((is_object($pd_2_15) && (($pd_2_15)->{'tag'} === "Nothing"))) {
$__t22 = true;
goto end_branch_22;;
};
if ((is_object($pd_2_15) && (($pd_2_15)->{'tag'} === "Just"))) {
$__t22 = false;
goto end_branch_22;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t22 = null;
end_branch_22:;
$__t23 = null;;
if ((is_object($pm_1_14) && (($pm_1_14)->{'tag'} === "Nothing"))) {
$__t23 = true;
goto end_branch_23;;
};
if ((is_object($pm_1_14) && (($pm_1_14)->{'tag'} === "Just"))) {
$__t23 = false;
goto end_branch_23;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t23 = null;
end_branch_23:;
return ((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])['conj'])($__t22))($__t23);
})()) {
$__t21 = (($GLOBALS['Data_Date_Component_enumYear'])['pred'])(($v_0)->{'value0'});
goto end_branch_21;;
};
  $__t21 = new Phpurs_Data1("Just", ($v_0)->{'value0'});
  end_branch_21:;
  $__t24 = null;;
  if ((function() use ($pd_2_15, &$__fn) {
$__t25 = null;;
if ((is_object($pd_2_15) && (($pd_2_15)->{'tag'} === "Nothing"))) {
$__t25 = true;
goto end_branch_25;;
};
if ((is_object($pd_2_15) && (($pd_2_15)->{'tag'} === "Just"))) {
$__t25 = false;
goto end_branch_25;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t25 = null;
end_branch_25:;
return $__t25;
})()) {
$__t24 = new Phpurs_Data1("Just", $l_4_20);
goto end_branch_24;;
};
  $__t24 = $pd_2_15;
  end_branch_24:;
  $__res = ((($GLOBALS['Data_Maybe_applyMaybe'])['apply'])(((($GLOBALS['Data_Maybe_applyMaybe'])['apply'])(((($GLOBALS['Data_Maybe_functorMaybe'])['map'])($GLOBALS['Data_Date_Date']))($__t21)))((($GLOBALS['Data_Maybe_applicativeMaybe'])['pure'])($m_prime_3_16))))($__t24);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Ord0" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Date_ordDate'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Date_diff
$GLOBALS['Data_Date_diff'] = (function() {
  $__fn = function($dictDuration_0 = null, $v_1 = null, $v1_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($dictDuration_0)['toDuration'])(($GLOBALS['Data_Date_calcDiff'])(($v_1)->{'value0'}, (($GLOBALS['Data_Date_Component_boundedEnumMonth'])['fromEnum'])(($v_1)->{'value1'}), ($v_1)->{'value2'}, ($v1_2)->{'value0'}, (($GLOBALS['Data_Date_Component_boundedEnumMonth'])['fromEnum'])(($v1_2)->{'value1'}), ($v1_2)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_Date_day
$GLOBALS['Data_Date_day'] = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ($v_0)->{'value2'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Date_canonicalDate
$GLOBALS['Data_Date_canonicalDate'] = (function() {
  $__fn = function($y_0 = null, $m_1 = null, $d_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($GLOBALS['Data_Date_canonicalDateImpl'])((function() {
  $__fn = function($y_prime_3 = null, $m_prime_4 = null, $d_prime_5 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__local_var_6_0 = (($GLOBALS['Data_Date_Component_boundedEnumMonth'])['toEnum'])($m_prime_4);
  $__t1 = null;;
  if ((is_object($__local_var_6_0) && (($__local_var_6_0)->{'tag'} === "Just"))) {
$__t1 = ($__local_var_6_0)->{'value0'};
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = new Phpurs_Data3("Date", $y_prime_3, $__t1, $d_prime_5);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), $y_0, (($GLOBALS['Data_Date_Component_boundedEnumMonth'])['fromEnum'])($m_1), $d_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_Date_exactDate
$GLOBALS['Data_Date_exactDate'] = (function() {
  $__fn = function($y_0 = null, $m_1 = null, $d_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if (((($GLOBALS['Data_Date_eqDate'])['eq'])(((($GLOBALS['Data_Date_canonicalDate'])($y_0))($m_1))($d_2)))(new Phpurs_Data3("Date", $y_0, $m_1, $d_2))) {
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data3("Date", $y_0, $m_1, $d_2));
goto end_branch_0;;
};
  $__t0 = new Phpurs_Data0("Nothing");
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_Date_boundedDate
$GLOBALS['Data_Date_boundedDate'] = ["bottom" => new Phpurs_Data3("Date", ($GLOBALS['Data_Date_Component_boundedYear'])['bottom'], ($GLOBALS['Data_Date_Component_boundedMonth'])['bottom'], ($GLOBALS['Data_Date_Component_boundedDay'])['bottom']), "top" => new Phpurs_Data3("Date", ($GLOBALS['Data_Date_Component_boundedYear'])['top'], ($GLOBALS['Data_Date_Component_boundedMonth'])['top'], ($GLOBALS['Data_Date_Component_boundedDay'])['top']), "Ord0" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Date_ordDate'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Date_adjust
$GLOBALS['Data_Date_adjust'] = (function() {
  $__fn = function($v_0 = null, $date_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $adj_2_0 = null;
  $adj_2_0 = (function() use (&$adj_2_0) {
  $__fn = function($v1_3 = null, $v2_4 = null) use (&$adj_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t10 = null;;
  switch ($v1_3) {
case 0:
$__t10 = new Phpurs_Data1("Just", $v2_4);
goto end_branch_10;;
break;
default:
;
break;
};
  $j_5_1 = ((($GLOBALS['Data_Semiring_semiringInt'])['add'])($v1_3))((($GLOBALS['Data_Date_Component_boundedEnumDay'])['fromEnum'])(($v2_4)->{'value2'}));
  $low_6_2 = ($j_5_1 < 1);
  $__t3 = null;;
  if ($low_6_2) {
$__local_var_7_4 = (($GLOBALS['Data_Date_Component_enumMonth'])['pred'])(($v2_4)->{'value1'});
$__t5 = null;;
if ((is_object($__local_var_7_4) && (($__local_var_7_4)->{'tag'} === "Nothing"))) {
$__t5 = new Phpurs_Data0("December");
goto end_branch_5;;
};
if ((is_object($__local_var_7_4) && (($__local_var_7_4)->{'tag'} === "Just"))) {
$__t5 = (($GLOBALS['Control_Category_categoryFn'])['identity'])(($__local_var_7_4)->{'value0'});
goto end_branch_5;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t5 = null;
end_branch_5:;
$__t3 = $__t5;
goto end_branch_3;;
};
  $__t3 = ($v2_4)->{'value1'};
  end_branch_3:;
  $l_7_3 = (($GLOBALS['Data_Date_lastDayOfMonth'])(($v2_4)->{'value0'}))($__t3);
  $hi_8_7 = ($j_5_1 > (($GLOBALS['Data_Date_Component_boundedEnumDay'])['fromEnum'])($l_7_3));
  $__t8 = null;;
  if ($low_6_2) {
$__t8 = ((($GLOBALS['Data_Maybe_bindMaybe'])['bind'])(((($GLOBALS['Data_Maybe_functorMaybe'])['map'])((($GLOBALS['Data_Date_Date'])(($v2_4)->{'value0'}))(($v2_4)->{'value1'})))((($GLOBALS['Data_Date_Component_boundedEnumDay'])['toEnum'])(1))))(($GLOBALS['Data_Date_enumDate'])['pred']);
goto end_branch_8;;
};
  if ($hi_8_7) {
$__t8 = (($GLOBALS['Data_Date_enumDate'])['succ'])(new Phpurs_Data3("Date", ($v2_4)->{'value0'}, ($v2_4)->{'value1'}, $l_7_3));
goto end_branch_8;;
};
  $__t8 = ((($GLOBALS['Data_Maybe_functorMaybe'])['map'])((($GLOBALS['Data_Date_Date'])(($v2_4)->{'value0'}))(($v2_4)->{'value1'})))((($GLOBALS['Data_Date_Component_boundedEnumDay'])['toEnum'])($j_5_1));
  end_branch_8:;
  $__t9 = null;;
  if ($low_6_2) {
$__t9 = $j_5_1;
goto end_branch_9;;
};
  if ($hi_8_7) {
$__t9 = ((($GLOBALS['Data_Ring_ringInt'])['sub'])(((($GLOBALS['Data_Ring_ringInt'])['sub'])($j_5_1))((($GLOBALS['Data_Date_Component_boundedEnumDay'])['fromEnum'])($l_7_3))))(1);
goto end_branch_9;;
};
  $__t9 = 0;
  end_branch_9:;
  $__t10 = ((($GLOBALS['Data_Maybe_bindMaybe'])['bind'])($__t8))(($adj_2_0)($__t9));
  end_branch_10:;
  $__res = $__t10;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ((($GLOBALS['Data_Maybe_bindMaybe'])['bind'])(($GLOBALS['Data_Int_fromNumber'])($v_0)))(function($a_3 = null) use (&$adj_2_0, $date_1) {
  $__num = \func_num_args();
  $__res = (($adj_2_0)($a_3))($date_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

