<?php

namespace Data\DateTime\Instant;

// ALL IMPORTS: Data.Boolean, Data.Bounded, Data.Date, Data.Date.Component, Data.DateTime, Data.DateTime.Instant, Data.Enum, Data.Eq, Data.Function.Uncurried, Data.HeytingAlgebra, Data.Maybe, Data.Ord, Data.Ring, Data.Semigroup, Data.Show, Data.Time, Data.Time.Component, Data.Time.Duration, Partial.Unsafe, Prelude, Prim
// TO REQUIRE: Data.Boolean, Data.Bounded, Data.Date, Data.Date.Component, Data.DateTime, Data.DateTime.Instant, Data.Enum, Data.Eq, Data.Function.Uncurried, Data.HeytingAlgebra, Data.Maybe, Data.Ord, Data.Ring, Data.Semigroup, Data.Show, Data.Time, Data.Time.Component, Data.Time.Duration, Partial.Unsafe, Prelude
require_once __DIR__ . '/../Data.Boolean/index.php';
require_once __DIR__ . '/../Data.Bounded/index.php';
require_once __DIR__ . '/../Data.Date/index.php';
require_once __DIR__ . '/../Data.Date.Component/index.php';
require_once __DIR__ . '/../Data.DateTime/index.php';
require_once __DIR__ . '/../Data.DateTime.Instant/index.php';
require_once __DIR__ . '/../Data.Enum/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Function.Uncurried/index.php';
require_once __DIR__ . '/../Data.HeytingAlgebra/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ring/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
require_once __DIR__ . '/../Data.Time/index.php';
require_once __DIR__ . '/../Data.Time.Component/index.php';
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
$ffi_Data_DateTime_Instant = \call_user_func(function() {
  $exports = [];
$fromDateTimeImpl = function($y, $mo, $d, $h, $mi, $s, $ms) use (&$fromDateTimeImpl) {
    $dt = new \DateTime('now', new \DateTimeZone('UTC'));
    $dt->setDate($y, $mo, $d);
    $dt->setTime($h, $mi, $s, $ms * 1000);
    return (float)$dt->getTimestamp() * 1000 + (int)$dt->format('v');
};

$toDateTimeImpl = function($ctor, $instant) use (&$toDateTimeImpl) {
    $seconds = floor($instant / 1000);
    $ms = $instant - ($seconds * 1000);
    $dt = new \DateTime("@" . $seconds, new \DateTimeZone('UTC'));
    
    return $ctor
        ((int)$dt->format('Y'))
        ((int)$dt->format('n'))
        ((int)$dt->format('j'))
        ((int)$dt->format('G'))
        ((int)$dt->format('i'))
        ((int)$dt->format('s'))
        ((int)$ms);
};

$exports['fromDateTimeImpl'] = $fromDateTimeImpl;
$exports['toDateTimeImpl'] = $toDateTimeImpl;
return $exports;
  return $exports;
});
$GLOBALS['Data_DateTime_Instant_fromDateTimeImpl'] = ($ffi_Data_DateTime_Instant['fromDateTimeImpl'] ?? \phpurs_ffi_fallback($ffi_Data_DateTime_Instant, 'fromDateTimeImpl'));
function majData_majDatemajTime_majInstant_tomajDatemajTimemajImpl($v0, $v1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majDatemajTime_majInstant_tomajDatemajTimemajImpl';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Data_DateTime_Instant;
  $f = ($ffi_Data_DateTime_Instant['toDateTimeImpl'] ?? \phpurs_ffi_fallback($ffi_Data_DateTime_Instant, 'toDateTimeImpl'));
  return $f($v0, $v1);
}
$GLOBALS['Data_DateTime_Instant_toDateTimeImpl'] = __NAMESPACE__ . '\\majData_majDatemajTime_majInstant_tomajDatemajTimemajImpl';





// Data_DateTime_Instant_negate
$GLOBALS['Data_DateTime_Instant_negate'] = (function() use (&$__fn) {
$zero_0_0 = ((($GLOBALS['Data_Ring_ringNumber'])->{'Semiring0'})(null))->{'zero'};
return function($a_1) use ($zero_0_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Ring_ringNumber'])->{'sub'})($zero_0_0))($a_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
})();

// Data_DateTime_Instant_negateDuration
$GLOBALS['Data_DateTime_Instant_negateDuration'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Time_Duration_durationMilliseconds'])->{'toDuration'}))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Time_Duration_negate']))(($GLOBALS['Data_Time_Duration_durationMilliseconds'])->{'fromDuration'}));

// Data_DateTime_Instant_unInstant
function majData_majDatemajTime_majInstant_unmajInstant(float $v_0): float|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majDatemajTime_majInstant_unmajInstant';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $v_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_DateTime_Instant_unInstant'] = __NAMESPACE__ . '\\majData_majDatemajTime_majInstant_unmajInstant';

// Data_DateTime_Instant_toDateTime_closure
$GLOBALS['Data_DateTime_Instant_toDateTime_closure'] = ($GLOBALS['Data_DateTime_Instant_toDateTimeImpl'])((function() {
  $__fn = function($y_0, $mo_1 = null, $d_2 = null, $h_3 = null, $mi_4 = null, $s_5 = null, $ms_6 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 7) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 7);
  }
  $__local_var_7_0 = (($GLOBALS['Data_Date_Component_boundedEnumMonth'])->{'toEnum'})($mo_1);
  $__t1 = null;;
  if ($__local_var_7_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = ($__local_var_7_0)->{'value0'};
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = new \Data\DateTime\Data_DateTime_DateTime(\Data\Date\majData_majDate_canonicalmajDate($y_0, $__t1, $d_2), new \Data\Time\Data_Time_Time($h_3, $mi_4, $s_5, $ms_6));
  goto __end;;
  __end:
  return $__num > 7 ? $__res(...\array_slice(\func_get_args(), 7)) : $__res;
  };
  return $__fn;
})());

// Data_DateTime_Instant_toDateTime
function majData_majDatemajTime_majInstant_tomajDatemajTime(float $v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majDatemajTime_majInstant_tomajDatemajTime';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_DateTime_Instant_toDateTime_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_DateTime_Instant_toDateTime'] = __NAMESPACE__ . '\\majData_majDatemajTime_majInstant_tomajDatemajTime';

// Data_DateTime_Instant_showInstant
$GLOBALS['Data_DateTime_Instant_showInstant'] = (object)["show" => function($v_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})("(Instant "))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})((($GLOBALS['Data_Time_Duration_showMilliseconds'])->{'show'})($v_0)))(")"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_DateTime_Instant_ordDateTime
$GLOBALS['Data_DateTime_Instant_ordDateTime'] = $GLOBALS['Data_Ord_ordNumber'];

// Data_DateTime_Instant_instant
function majData_majDatemajTime_majInstant_instant(float $v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majDatemajTime_majInstant_instant';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__t0 = null;;
  if (((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'conj'})(($v_0 >= ($GLOBALS['Data_DateTime_Instant_negate'])(8639977881600000.0))))(($v_0 <= 8639977881599999.0))) {
$__t0 = new \Data\Maybe\Data_Maybe_Just($v_0);
goto end_branch_0;;
};
  $__t0 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_DateTime_Instant_instant'] = __NAMESPACE__ . '\\majData_majDatemajTime_majInstant_instant';

// Data_DateTime_Instant_fromDateTime
function majData_majDatemajTime_majInstant_frommajDatemajTime($v_0): float|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majDatemajTime_majInstant_frommajDatemajTime';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_DateTime_Instant_fromDateTimeImpl'])((($v_0)->{'value0'})->{'value0'}, (($GLOBALS['Data_Date_Component_boundedEnumMonth'])->{'fromEnum'})((($v_0)->{'value0'})->{'value1'}), (($v_0)->{'value0'})->{'value2'}, (($v_0)->{'value1'})->{'value0'}, (($v_0)->{'value1'})->{'value1'}, (($v_0)->{'value1'})->{'value2'}, (($v_0)->{'value1'})->{'value3'});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_DateTime_Instant_fromDateTime'] = __NAMESPACE__ . '\\majData_majDatemajTime_majInstant_frommajDatemajTime';

// Data_DateTime_Instant_fromDate
function majData_majDatemajTime_majInstant_frommajDate($d_0): float|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majDatemajTime_majInstant_frommajDate';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_DateTime_Instant_fromDateTimeImpl'])(($d_0)->{'value0'}, (($GLOBALS['Data_Date_Component_boundedEnumMonth'])->{'fromEnum'})(($d_0)->{'value1'}), ($d_0)->{'value2'}, ($GLOBALS['Data_Time_Component_boundedHour'])->{'bottom'}, ($GLOBALS['Data_Time_Component_boundedMinute'])->{'bottom'}, ($GLOBALS['Data_Time_Component_boundedSecond'])->{'bottom'}, ($GLOBALS['Data_Time_Component_boundedMillisecond'])->{'bottom'});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_DateTime_Instant_fromDate'] = __NAMESPACE__ . '\\majData_majDatemajTime_majInstant_frommajDate';

// Data_DateTime_Instant_eqDateTime
$GLOBALS['Data_DateTime_Instant_eqDateTime'] = $GLOBALS['Data_Eq_eqNumber'];

// Data_DateTime_Instant_diff
function majData_majDatemajTime_majInstant_diff($dictDuration_0, $dt1_1 = null, $dt2_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majDatemajTime_majInstant_diff';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($dictDuration_0)->{'toDuration'})(((($GLOBALS['Data_Time_Duration_semigroupMilliseconds'])->{'append'})($dt1_1))(($GLOBALS['Data_DateTime_Instant_negateDuration'])($dt2_2)));
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_DateTime_Instant_diff'] = __NAMESPACE__ . '\\majData_majDatemajTime_majInstant_diff';

// Data_DateTime_Instant_boundedInstant
$GLOBALS['Data_DateTime_Instant_boundedInstant'] = (object)["bottom" => ($GLOBALS['Data_DateTime_Instant_negate'])(8639977881600000.0), "top" => 8639977881599999.0, "Ord0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Ord_ordNumber'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

