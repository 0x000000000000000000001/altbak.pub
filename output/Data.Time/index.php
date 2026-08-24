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
if (!\function_exists(__NAMESPACE__ . '\\phpurs_execute_effect')) {
  function phpurs_execute_effect($val) {
    if (\is_callable($val)) {
      return $val($GLOBALS['Data_Unit_unit']);
    }
    return $val;
  }
}

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };


final class Data_Time_Time { public $tag = 'Time'; public function __construct(public int $value0, public int $value1, public int $value2, public int $value3) {} }

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
  $__res = (((((((("(Time (Hour " . \Data\Show\majData_majShow_showmajIntmajImpl(($v_0)->{'value0'})) . ") (Minute ") . \Data\Show\majData_majShow_showmajIntmajImpl(($v_0)->{'value1'})) . ") (Second ") . \Data\Show\majData_majShow_showmajIntmajImpl(($v_0)->{'value2'})) . ") (Millisecond ") . \Data\Show\majData_majShow_showmajIntmajImpl(($v_0)->{'value3'})) . "))");
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
  $hours_1_0 = \Data\Number\majData_majNumber_floor(($v_0 / 3600000.0));
  $minutes_2_1 = \Data\Number\majData_majNumber_floor((($v_0 - ($hours_1_0 * 3600000.0)) / 60000.0));
  $seconds_3_2 = \Data\Number\majData_majNumber_floor((($v_0 - (($hours_1_0 * 3600000.0) + ($minutes_2_1 * 60000.0))) / 1000.0));
  $__local_var_4_3 = \Data\Int\majData_majInt_floor($hours_1_0);
  $__t4 = null;;
  if ((($__local_var_4_3 >= 0) && ($__local_var_4_3 <= 23))) {
$__t4 = new \Data\Maybe\Data_Maybe_Just($__local_var_4_3);
goto end_branch_4;;
};
  $__t4 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_4:;
  $__local_var_4_3 = $__t4;
  $__t6 = null;;
  if ($__local_var_4_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t6 = new \Data\Maybe\Data_Maybe_Just(($GLOBALS['Data_Time_Time'])(($__local_var_4_3)->{'value0'}));
goto end_branch_6;;
};
  $__t6 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_6:;
  $__local_var_4_3 = $__t6;
  $__local_var_5_8 = \Data\Int\majData_majInt_floor($minutes_2_1);
  $__t9 = null;;
  if ((($__local_var_5_8 >= 0) && ($__local_var_5_8 <= 59))) {
$__t9 = new \Data\Maybe\Data_Maybe_Just($__local_var_5_8);
goto end_branch_9;;
};
  $__t9 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_9:;
  $__local_var_5_8 = $__t9;
  $__t11 = null;;
  if ($__local_var_4_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t12 = null;;
if ($__local_var_5_8 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t12 = new \Data\Maybe\Data_Maybe_Just((($__local_var_4_3)->{'value0'})(($__local_var_5_8)->{'value0'}));
goto end_branch_12;;
};
$__t12 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_12:;
$__t11 = $__t12;
goto end_branch_11;;
};
  if ($__local_var_4_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t11 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_11;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t11 = null;
  end_branch_11:;
  $__local_var_4_3 = $__t11;
  $__local_var_5_14 = \Data\Int\majData_majInt_floor($seconds_3_2);
  $__t15 = null;;
  if ((($__local_var_5_14 >= 0) && ($__local_var_5_14 <= 59))) {
$__t15 = new \Data\Maybe\Data_Maybe_Just($__local_var_5_14);
goto end_branch_15;;
};
  $__t15 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_15:;
  $__local_var_5_14 = $__t15;
  $__t17 = null;;
  if ($__local_var_4_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t18 = null;;
if ($__local_var_5_14 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t18 = new \Data\Maybe\Data_Maybe_Just((($__local_var_4_3)->{'value0'})(($__local_var_5_14)->{'value0'}));
goto end_branch_18;;
};
$__t18 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_18:;
$__t17 = $__t18;
goto end_branch_17;;
};
  if ($__local_var_4_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t17 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_17;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t17 = null;
  end_branch_17:;
  $__local_var_4_3 = $__t17;
  $__local_var_5_20 = \Data\Int\majData_majInt_floor(($v_0 - ((($hours_1_0 * 3600000.0) + ($minutes_2_1 * 60000.0)) + ($seconds_3_2 * 1000.0))));
  $__t21 = null;;
  if ((($__local_var_5_20 >= 0) && ($__local_var_5_20 <= 999))) {
$__t21 = new \Data\Maybe\Data_Maybe_Just($__local_var_5_20);
goto end_branch_21;;
};
  $__t21 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_21:;
  $__local_var_5_20 = $__t21;
  $__t23 = null;;
  if ($__local_var_4_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t24 = null;;
if ($__local_var_5_20 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t24 = new \Data\Maybe\Data_Maybe_Just((($__local_var_4_3)->{'value0'})(($__local_var_5_20)->{'value0'}));
goto end_branch_24;;
};
$__t24 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_24:;
$__t23 = $__t24;
goto end_branch_23;;
};
  if ($__local_var_4_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t23 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_23;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t23 = null;
  end_branch_23:;
  $__local_var_4_3 = $__t23;
  $__t26 = null;;
  if ($__local_var_4_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t26 = ($__local_var_4_3)->{'value0'};
goto end_branch_26;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t26 = null;
  end_branch_26:;
  $__res = $__t26;
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
  $__res = ((((3600000.0 * \Data\Int\majData_majInt_tomajNumber(($t_0)->{'value0'})) + (60000.0 * \Data\Int\majData_majInt_tomajNumber(($t_0)->{'value1'}))) + (1000.0 * \Data\Int\majData_majInt_tomajNumber(($t_0)->{'value2'}))) + \Data\Int\majData_majInt_tomajNumber(($t_0)->{'value3'}));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Time_timeToMillis'] = __NAMESPACE__ . '\\majData_majTime_timemajTomajMillis';

// Data_Time_eqTime
$GLOBALS['Data_Time_eqTime'] = (object)["eq" => function($x_0) {
  $__num = \func_num_args();
  $__res = function($y_1) use ($x_0) {
  $__num = \func_num_args();
  $__res = ((((($x_0)->{'value0'} === ($y_1)->{'value0'}) && (($x_0)->{'value1'} === ($y_1)->{'value1'})) && (($x_0)->{'value2'} === ($y_1)->{'value2'})) && (($x_0)->{'value3'} === ($y_1)->{'value3'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Time_ordTime
$GLOBALS['Data_Time_ordTime'] = (object)["compare" => function($x_0) {
  $__num = \func_num_args();
  $__res = function($y_1) use ($x_0) {
  $__num = \func_num_args();
  $v_2_0 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), ($x_0)->{'value0'}, ($y_1)->{'value0'});
  $__t5 = null;;
  if ($v_2_0 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t5 = new \Data\Ordering\Data_Ordering_LT();
goto end_branch_5;;
};
  if ($v_2_0 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t5 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_5;;
};
  $v1_3_1 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), ($x_0)->{'value1'}, ($y_1)->{'value1'});
  $__t4 = null;;
  if ($v1_3_1 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t4 = new \Data\Ordering\Data_Ordering_LT();
goto end_branch_4;;
};
  if ($v1_3_1 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t4 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_4;;
};
  $v2_4_2 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), ($x_0)->{'value2'}, ($y_1)->{'value2'});
  $__t3 = null;;
  if ($v2_4_2 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t3 = new \Data\Ordering\Data_Ordering_LT();
goto end_branch_3;;
};
  if ($v2_4_2 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t3 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_3;;
};
  $__t3 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), ($x_0)->{'value3'}, ($y_1)->{'value3'});
  end_branch_3:;
  $__t4 = $__t3;
  end_branch_4:;
  $__t5 = $__t4;
  end_branch_5:;
  $__res = $__t5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq0" => function($_dollar___unused_0) {
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
  $__res = (($dictDuration_0)->{'toDuration'})((\Data\Time\majData_majTime_timemajTomajMillis($t1_1) + \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($x_3) {
  $__num = \func_num_args();
  $__res = $x_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($a_3) {
  $__num = \func_num_args();
  $__res = ( - $a_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($x_3) {
  $__num = \func_num_args();
  $__res = $x_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), \Data\Time\majData_majTime_timemajTomajMillis($t2_2))));
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Time_diff'] = __NAMESPACE__ . '\\majData_majTime_diff';

// Data_Time_boundedTime
$GLOBALS['Data_Time_boundedTime'] = (object)["bottom" => new \Data\Time\Data_Time_Time(0, 0, 0, 0), "top" => new \Data\Time\Data_Time_Time(23, 59, 59, 999), "Ord0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Time_ordTime'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Time_maxTime
$GLOBALS['Data_Time_maxTime'] = \Data\Time\majData_majTime_timemajTomajMillis(new \Data\Time\Data_Time_Time(23, 59, 59, 999));

// Data_Time_minTime
$GLOBALS['Data_Time_minTime'] = \Data\Time\majData_majTime_timemajTomajMillis(new \Data\Time\Data_Time_Time(0, 0, 0, 0));

// Data_Time_adjust
function majData_majTime_adjust($dictDuration_0, $d_1 = null, $t_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTime_adjust';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $d_prime__3_0 = (($dictDuration_0)->{'fromDuration'})($d_1);
  $wholeDays_4_1 = \Data\Number\majData_majNumber_floor(($d_prime__3_0 / 86400000.0));
  $msAdjusted_5_2 = ((\Data\Time\majData_majTime_timemajTomajMillis($t_2) + $d_prime__3_0) + \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($x_5) {
  $__num = \func_num_args();
  $__res = $x_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($a_5) {
  $__num = \func_num_args();
  $__res = ( - $a_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($x_5) {
  $__num = \func_num_args();
  $__res = $x_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), ($wholeDays_4_1 * 86400000.0)));
  $__t4 = null;;
  if (\Data\Ord\majData_majOrd_ordmajNumbermajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), $msAdjusted_5_2, $GLOBALS['Data_Time_maxTime']) instanceof \Data\Ordering\Data_Ordering_GT) {
$__t4 = 1.0;
goto end_branch_4;;
};
  $__t3 = null;;
  if (\Data\Ord\majData_majOrd_ordmajNumbermajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), $msAdjusted_5_2, $GLOBALS['Data_Time_minTime']) instanceof \Data\Ordering\Data_Ordering_LT) {
$__t3 = -1.0;
goto end_branch_3;;
};
  $__t3 = 0.0;
  end_branch_3:;
  $__t4 = $__t3;
  end_branch_4:;
  $wrap_6_3 = $__t4;
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($wholeDays_4_1 + $wrap_6_3), \Data\Time\majData_majTime_millismajTomajTime(($msAdjusted_5_2 + (86400000.0 * ( - $wrap_6_3)))));
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Time_adjust'] = __NAMESPACE__ . '\\majData_majTime_adjust';

