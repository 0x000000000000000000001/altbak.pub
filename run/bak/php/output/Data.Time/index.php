<?php

namespace Data\Time;

// ALL IMPORTS: Control.Apply, Data.Bounded, Data.Enum, Data.Eq, Data.EuclideanRing, Data.Function, Data.Functor, Data.HeytingAlgebra, Data.Int, Data.Maybe, Data.Newtype, Data.Number, Data.Ord, Data.Ordering, Data.Ring, Data.Semigroup, Data.Semiring, Data.Show, Data.Time, Data.Time.Component, Data.Time.Duration, Data.Tuple, Partial.Unsafe, Prelude, Prim
// TO REQUIRE: Control.Apply, Data.Bounded, Data.Enum, Data.Eq, Data.EuclideanRing, Data.Function, Data.Functor, Data.HeytingAlgebra, Data.Int, Data.Maybe, Data.Newtype, Data.Number, Data.Ord, Data.Ordering, Data.Ring, Data.Semigroup, Data.Semiring, Data.Show, Data.Time, Data.Time.Component, Data.Time.Duration, Data.Tuple, Partial.Unsafe, Prelude
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Data.Bounded/index.php';
require_once __DIR__ . '/../Data.Enum/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.EuclideanRing/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.HeytingAlgebra/index.php';
require_once __DIR__ . '/../Data.Int/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Newtype/index.php';
require_once __DIR__ . '/../Data.Number/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ordering/index.php';
require_once __DIR__ . '/../Data.Ring/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Semiring/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
require_once __DIR__ . '/../Data.Time/index.php';
require_once __DIR__ . '/../Data.Time.Component/index.php';
require_once __DIR__ . '/../Data.Time.Duration/index.php';
require_once __DIR__ . '/../Data.Tuple/index.php';
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


final class Data_Time_Time { public $tag = 'Time'; public function __construct(public int $value0, public int $value1, public int $value2, public int $value3) {} }

// Data_Time_negateDuration
$GLOBALS['Data_Time_negateDuration'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Time_Duration_durationMilliseconds'])->{'toDuration'}))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Time_Duration_negate']))(($GLOBALS['Data_Time_Duration_durationMilliseconds'])->{'fromDuration'}));

// Data_Time_negate
$GLOBALS['Data_Time_negate'] = (function() use (&$__fn) {
$zero_0_0 = ((($GLOBALS['Data_Ring_ringNumber'])->{'Semiring0'})(null))->{'zero'};
return function($a_1) use ($zero_0_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Ring_ringNumber'])->{'sub'})($zero_0_0))($a_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
})();

// Data_Time_Time
$GLOBALS['Data_Time_Time'] = (function() {
  $__fn = function($value0, $value1 = null, $value2 = null, $value3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = new \Data\Time\Data_Time_Time($value0, $value1, $value2, $value3);
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})();

// Data_Time_showTime
$GLOBALS['Data_Time_showTime'] = (object)["show" => function($v_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})("(Time "))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})((($GLOBALS['Data_Time_Component_showHour'])->{'show'})(($v_0)->{'value0'})))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})(" "))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})((($GLOBALS['Data_Time_Component_showMinute'])->{'show'})(($v_0)->{'value1'})))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})(" "))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})((($GLOBALS['Data_Time_Component_showSecond'])->{'show'})(($v_0)->{'value2'})))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})(" "))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})((($GLOBALS['Data_Time_Component_showMillisecond'])->{'show'})(($v_0)->{'value3'})))(")"))))))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Time_setSecond
function majData_majTime_setmajSecond(int $s_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTime_setmajSecond';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Time\Data_Time_Time(($v_1)->{'value0'}, ($v_1)->{'value1'}, $s_0, ($v_1)->{'value3'});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Time_setSecond'] = __NAMESPACE__ . '\\majData_majTime_setmajSecond';

// Data_Time_setMinute
function majData_majTime_setmajMinute(int $m_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTime_setmajMinute';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Time\Data_Time_Time(($v_1)->{'value0'}, $m_0, ($v_1)->{'value2'}, ($v_1)->{'value3'});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Time_setMinute'] = __NAMESPACE__ . '\\majData_majTime_setmajMinute';

// Data_Time_setMillisecond
function majData_majTime_setmajMillisecond(int $ms_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTime_setmajMillisecond';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Time\Data_Time_Time(($v_1)->{'value0'}, ($v_1)->{'value1'}, ($v_1)->{'value2'}, $ms_0);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Time_setMillisecond'] = __NAMESPACE__ . '\\majData_majTime_setmajMillisecond';

// Data_Time_setHour
function majData_majTime_setmajHour(int $h_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTime_setmajHour';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Time\Data_Time_Time($h_0, ($v_1)->{'value1'}, ($v_1)->{'value2'}, ($v_1)->{'value3'});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Time_setHour'] = __NAMESPACE__ . '\\majData_majTime_setmajHour';

// Data_Time_second
function majData_majTime_second($v_0): int|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTime_second';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($v_0)->{'value2'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Time_second'] = __NAMESPACE__ . '\\majData_majTime_second';

// Data_Time_minute
function majData_majTime_minute($v_0): int|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTime_minute';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($v_0)->{'value1'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Time_minute'] = __NAMESPACE__ . '\\majData_majTime_minute';

// Data_Time_millisecond
function majData_majTime_millisecond($v_0): int|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTime_millisecond';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($v_0)->{'value3'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Time_millisecond'] = __NAMESPACE__ . '\\majData_majTime_millisecond';

// Data_Time_millisToTime
function majData_majTime_millismajTomajTime(float $v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTime_millismajTomajTime';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $hours_1_0 = \Data\Number\majData_majNumber_floor(((($GLOBALS['Data_EuclideanRing_euclideanRingNumber'])->{'div'})($v_0))(3600000.0));
  $minutes_2_1 = \Data\Number\majData_majNumber_floor(((($GLOBALS['Data_EuclideanRing_euclideanRingNumber'])->{'div'})(((($GLOBALS['Data_Ring_ringNumber'])->{'sub'})($v_0))(((($GLOBALS['Data_Semiring_semiringNumber'])->{'mul'})($hours_1_0))(3600000.0))))(60000.0));
  $seconds_3_2 = \Data\Number\majData_majNumber_floor(((($GLOBALS['Data_EuclideanRing_euclideanRingNumber'])->{'div'})(((($GLOBALS['Data_Ring_ringNumber'])->{'sub'})($v_0))(((($GLOBALS['Data_Semiring_semiringNumber'])->{'add'})(((($GLOBALS['Data_Semiring_semiringNumber'])->{'mul'})($hours_1_0))(3600000.0)))(((($GLOBALS['Data_Semiring_semiringNumber'])->{'mul'})($minutes_2_1))(60000.0)))))(1000.0));
  $__local_var_4_3 = ((($GLOBALS['Data_Maybe_applyMaybe'])->{'apply'})(((($GLOBALS['Data_Maybe_applyMaybe'])->{'apply'})(((($GLOBALS['Data_Maybe_applyMaybe'])->{'apply'})(((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})($GLOBALS['Data_Time_Time']))((($GLOBALS['Data_Time_Component_boundedEnumHour'])->{'toEnum'})(\Data\Int\majData_majInt_floor($hours_1_0)))))((($GLOBALS['Data_Time_Component_boundedEnumMinute'])->{'toEnum'})(\Data\Int\majData_majInt_floor($minutes_2_1)))))((($GLOBALS['Data_Time_Component_boundedEnumSecond'])->{'toEnum'})(\Data\Int\majData_majInt_floor($seconds_3_2)))))((($GLOBALS['Data_Time_Component_boundedEnumMillisecond'])->{'toEnum'})(\Data\Int\majData_majInt_floor(((($GLOBALS['Data_Ring_ringNumber'])->{'sub'})($v_0))(((($GLOBALS['Data_Semiring_semiringNumber'])->{'add'})(((($GLOBALS['Data_Semiring_semiringNumber'])->{'add'})(((($GLOBALS['Data_Semiring_semiringNumber'])->{'mul'})($hours_1_0))(3600000.0)))(((($GLOBALS['Data_Semiring_semiringNumber'])->{'mul'})($minutes_2_1))(60000.0))))(((($GLOBALS['Data_Semiring_semiringNumber'])->{'mul'})($seconds_3_2))(1000.0))))));
  $__t4 = null;;
  if ($__local_var_4_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t4 = ($__local_var_4_3)->{'value0'};
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
$GLOBALS['Data_Time_millisToTime'] = __NAMESPACE__ . '\\majData_majTime_millismajTomajTime';

// Data_Time_hour
function majData_majTime_hour($v_0): int|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTime_hour';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($v_0)->{'value0'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Time_hour'] = __NAMESPACE__ . '\\majData_majTime_hour';

// Data_Time_timeToMillis
function majData_majTime_timemajTomajMillis($t_0): float|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTime_timemajTomajMillis';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ((($GLOBALS['Data_Semiring_semiringNumber'])->{'add'})(((($GLOBALS['Data_Semiring_semiringNumber'])->{'add'})(((($GLOBALS['Data_Semiring_semiringNumber'])->{'add'})(((($GLOBALS['Data_Semiring_semiringNumber'])->{'mul'})(3600000.0))(\Data\Int\majData_majInt_tomajNumber((($GLOBALS['Data_Time_Component_boundedEnumHour'])->{'fromEnum'})(($t_0)->{'value0'})))))(((($GLOBALS['Data_Semiring_semiringNumber'])->{'mul'})(60000.0))(\Data\Int\majData_majInt_tomajNumber((($GLOBALS['Data_Time_Component_boundedEnumMinute'])->{'fromEnum'})(($t_0)->{'value1'}))))))(((($GLOBALS['Data_Semiring_semiringNumber'])->{'mul'})(1000.0))(\Data\Int\majData_majInt_tomajNumber((($GLOBALS['Data_Time_Component_boundedEnumSecond'])->{'fromEnum'})(($t_0)->{'value2'}))))))(\Data\Int\majData_majInt_tomajNumber((($GLOBALS['Data_Time_Component_boundedEnumMillisecond'])->{'fromEnum'})(($t_0)->{'value3'})));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Time_timeToMillis'] = __NAMESPACE__ . '\\majData_majTime_timemajTomajMillis';

// Data_Time_eqTime
$GLOBALS['Data_Time_eqTime'] = (object)["eq" => (function() {
  $__fn = function($x_0, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'conj'})(((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'conj'})(((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'conj'})(((($GLOBALS['Data_Time_Component_eqHour'])->{'eq'})(($x_0)->{'value0'}))(($y_1)->{'value0'})))(((($GLOBALS['Data_Time_Component_eqMinute'])->{'eq'})(($x_0)->{'value1'}))(($y_1)->{'value1'}))))(((($GLOBALS['Data_Time_Component_eqSecond'])->{'eq'})(($x_0)->{'value2'}))(($y_1)->{'value2'}))))(((($GLOBALS['Data_Time_Component_eqMillisecond'])->{'eq'})(($x_0)->{'value3'}))(($y_1)->{'value3'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];

// Data_Time_ordTime
$GLOBALS['Data_Time_ordTime'] = (object)["compare" => (function() {
  $__fn = function($x_0, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v_2_0 = ((($GLOBALS['Data_Time_Component_ordHour'])->{'compare'})(($x_0)->{'value0'}))(($y_1)->{'value0'});
  $__t5 = null;;
  if ($v_2_0 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t5 = new \Data\Ordering\Data_Ordering_LT();
goto end_branch_5;;
};
  if ($v_2_0 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t5 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_5;;
};
  $v1_3_1 = ((($GLOBALS['Data_Time_Component_ordMinute'])->{'compare'})(($x_0)->{'value1'}))(($y_1)->{'value1'});
  $__t4 = null;;
  if ($v1_3_1 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t4 = new \Data\Ordering\Data_Ordering_LT();
goto end_branch_4;;
};
  if ($v1_3_1 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t4 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_4;;
};
  $v2_4_2 = ((($GLOBALS['Data_Time_Component_ordSecond'])->{'compare'})(($x_0)->{'value2'}))(($y_1)->{'value2'});
  $__t3 = null;;
  if ($v2_4_2 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t3 = new \Data\Ordering\Data_Ordering_LT();
goto end_branch_3;;
};
  if ($v2_4_2 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t3 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_3;;
};
  $__t3 = ((($GLOBALS['Data_Time_Component_ordMillisecond'])->{'compare'})(($x_0)->{'value3'}))(($y_1)->{'value3'});
  end_branch_3:;
  $__t4 = $__t3;
  end_branch_4:;
  $__t5 = $__t4;
  end_branch_5:;
  $__res = $__t5;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Time_eqTime'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Time_diff
function majData_majTime_diff($dictDuration_0, $t1_1 = null, $t2_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTime_diff';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($dictDuration_0)->{'toDuration'})(((($GLOBALS['Data_Time_Duration_semigroupMilliseconds'])->{'append'})(\Data\Time\majData_majTime_timemajTomajMillis($t1_1)))(($GLOBALS['Data_Time_negateDuration'])(\Data\Time\majData_majTime_timemajTomajMillis($t2_2))));
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Time_diff'] = __NAMESPACE__ . '\\majData_majTime_diff';

// Data_Time_boundedTime
$GLOBALS['Data_Time_boundedTime'] = (object)["bottom" => new \Data\Time\Data_Time_Time(($GLOBALS['Data_Time_Component_boundedHour'])->{'bottom'}, ($GLOBALS['Data_Time_Component_boundedMinute'])->{'bottom'}, ($GLOBALS['Data_Time_Component_boundedSecond'])->{'bottom'}, ($GLOBALS['Data_Time_Component_boundedMillisecond'])->{'bottom'}), "top" => new \Data\Time\Data_Time_Time(($GLOBALS['Data_Time_Component_boundedHour'])->{'top'}, ($GLOBALS['Data_Time_Component_boundedMinute'])->{'top'}, ($GLOBALS['Data_Time_Component_boundedSecond'])->{'top'}, ($GLOBALS['Data_Time_Component_boundedMillisecond'])->{'top'}), "Ord0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Time_ordTime'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Time_maxTime
$GLOBALS['Data_Time_maxTime'] = \Data\Time\majData_majTime_timemajTomajMillis(($GLOBALS['Data_Time_boundedTime'])->{'top'});

// Data_Time_minTime
$GLOBALS['Data_Time_minTime'] = \Data\Time\majData_majTime_timemajTomajMillis(($GLOBALS['Data_Time_boundedTime'])->{'bottom'});

// Data_Time_adjust
function majData_majTime_adjust($dictDuration_0, $d_1 = null, $t_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTime_adjust';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $d_prime_3_0 = (($dictDuration_0)->{'fromDuration'})($d_1);
  $wholeDays_4_1 = \Data\Number\majData_majNumber_floor(((($GLOBALS['Data_EuclideanRing_euclideanRingNumber'])->{'div'})($d_prime_3_0))(86400000.0));
  $msAdjusted_5_2 = ((($GLOBALS['Data_Time_Duration_semigroupMilliseconds'])->{'append'})(\Data\Time\majData_majTime_timemajTomajMillis($t_2)))(((($GLOBALS['Data_Time_Duration_semigroupMilliseconds'])->{'append'})($d_prime_3_0))(($GLOBALS['Data_Time_negateDuration'])((($GLOBALS['Data_Time_Duration_durationDays'])->{'fromDuration'})($wholeDays_4_1))));
  $__t3 = null;;
  if (($msAdjusted_5_2 > $GLOBALS['Data_Time_maxTime'])) {
$__t3 = 1.0;
goto end_branch_3;;
};
  if (($msAdjusted_5_2 < $GLOBALS['Data_Time_minTime'])) {
$__t3 = ($GLOBALS['Data_Time_negate'])(1.0);
goto end_branch_3;;
};
  $__t3 = 0.0;
  end_branch_3:;
  $wrap_6_3 = $__t3;
  $__res = new \Data\Tuple\Data_Tuple_Tuple(((($GLOBALS['Data_Time_Duration_semigroupDays'])->{'append'})($wholeDays_4_1))($wrap_6_3), \Data\Time\majData_majTime_millismajTomajTime(((($GLOBALS['Data_Time_Duration_semigroupMilliseconds'])->{'append'})($msAdjusted_5_2))(((($GLOBALS['Data_Semiring_semiringNumber'])->{'mul'})(86400000.0))(($GLOBALS['Data_Time_negate'])($wrap_6_3)))));
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Time_adjust'] = __NAMESPACE__ . '\\majData_majTime_adjust';

