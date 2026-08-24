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
if (!\function_exists(__NAMESPACE__ . '\\phpurs_execute_effect')) {
  function phpurs_execute_effect($val) {
    if (\is_callable($val)) {
      return $val($GLOBALS['Data_Unit_unit']);
    }
    return $val;
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
$GLOBALS['Data_DateTime_Instant_fromDateTimeImpl'] = (\array_key_exists('fromDateTimeImpl', $ffi_Data_DateTime_Instant) ? $ffi_Data_DateTime_Instant['fromDateTimeImpl'] : new class { public function __invoke(...$args) { return $this; } });
function majData_majDatemajTime_majInstant_tomajDatemajTimemajImpl($v0, $v1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majDatemajTime_majInstant_tomajDatemajTimemajImpl';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Data_DateTime_Instant;
  $f = (\array_key_exists('toDateTimeImpl', $ffi_Data_DateTime_Instant) ? $ffi_Data_DateTime_Instant['toDateTimeImpl'] : new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Data_DateTime_Instant_toDateTimeImpl'] = __NAMESPACE__ . '\\majData_majDatemajTime_majInstant_tomajDatemajTimemajImpl';





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
$GLOBALS['Data_DateTime_Instant_toDateTime_closure'] = ($GLOBALS['Data_DateTime_Instant_toDateTimeImpl'])(function($y_0) {
  $__num = \func_num_args();
  $__res = function($mo_1) use ($y_0) {
  $__num = \func_num_args();
  $__res = function($d_2) use ($mo_1, $y_0) {
  $__num = \func_num_args();
  $__res = function($h_3) use ($d_2, $mo_1, $y_0) {
  $__num = \func_num_args();
  $__res = function($mi_4) use ($d_2, $h_3, $mo_1, $y_0) {
  $__num = \func_num_args();
  $__res = function($s_5) use ($d_2, $h_3, $mi_4, $mo_1, $y_0) {
  $__num = \func_num_args();
  $__res = function($ms_6) use ($d_2, $h_3, $mi_4, $mo_1, $s_5, $y_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  switch ($mo_1) {
case 1:
$__t0 = new \Data\Date\Component\Data_Date_Component_January();
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($mo_1) {
case 2:
$__t0 = new \Data\Date\Component\Data_Date_Component_February();
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($mo_1) {
case 3:
$__t0 = new \Data\Date\Component\Data_Date_Component_March();
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($mo_1) {
case 4:
$__t0 = new \Data\Date\Component\Data_Date_Component_April();
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($mo_1) {
case 5:
$__t0 = new \Data\Date\Component\Data_Date_Component_May();
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($mo_1) {
case 6:
$__t0 = new \Data\Date\Component\Data_Date_Component_June();
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($mo_1) {
case 7:
$__t0 = new \Data\Date\Component\Data_Date_Component_July();
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($mo_1) {
case 8:
$__t0 = new \Data\Date\Component\Data_Date_Component_August();
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($mo_1) {
case 9:
$__t0 = new \Data\Date\Component\Data_Date_Component_September();
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($mo_1) {
case 10:
$__t0 = new \Data\Date\Component\Data_Date_Component_October();
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($mo_1) {
case 11:
$__t0 = new \Data\Date\Component\Data_Date_Component_November();
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($mo_1) {
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
  $__res = new \Data\DateTime\Data_DateTime_DateTime(\Data\Date\majData_majDate_canonicalmajDate($y_0, $__t0, $d_2), new \Data\Time\Data_Time_Time($h_3, $mi_4, $s_5, $ms_6));
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
};
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});

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
  $__res = (("(Instant (Milliseconds " . \Data\Show\majData_majShow_showmajNumbermajImpl($v_0)) . "))");
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
  if ((( ! \Data\Ord\majData_majOrd_ordmajNumbermajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), $v_0, -8639977881600000.0) instanceof \Data\Ordering\Data_Ordering_LT) && ( ! \Data\Ord\majData_majOrd_ordmajNumbermajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), $v_0, 8639977881599999.0) instanceof \Data\Ordering\Data_Ordering_GT))) {
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
  $__t0 = null;;
  if ((($v_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_January) {
$__t0 = 1;
goto end_branch_0;;
};
  if ((($v_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_February) {
$__t0 = 2;
goto end_branch_0;;
};
  if ((($v_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_March) {
$__t0 = 3;
goto end_branch_0;;
};
  if ((($v_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_April) {
$__t0 = 4;
goto end_branch_0;;
};
  if ((($v_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_May) {
$__t0 = 5;
goto end_branch_0;;
};
  if ((($v_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_June) {
$__t0 = 6;
goto end_branch_0;;
};
  if ((($v_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_July) {
$__t0 = 7;
goto end_branch_0;;
};
  if ((($v_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_August) {
$__t0 = 8;
goto end_branch_0;;
};
  if ((($v_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_September) {
$__t0 = 9;
goto end_branch_0;;
};
  if ((($v_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_October) {
$__t0 = 10;
goto end_branch_0;;
};
  if ((($v_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_November) {
$__t0 = 11;
goto end_branch_0;;
};
  if ((($v_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_December) {
$__t0 = 12;
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = ($GLOBALS['Data_DateTime_Instant_fromDateTimeImpl'])((($v_0)->{'value0'})->{'value0'}, $__t0, (($v_0)->{'value0'})->{'value2'}, (($v_0)->{'value1'})->{'value0'}, (($v_0)->{'value1'})->{'value1'}, (($v_0)->{'value1'})->{'value2'}, (($v_0)->{'value1'})->{'value3'});
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
  $__t0 = null;;
  if (($d_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_January) {
$__t0 = 1;
goto end_branch_0;;
};
  if (($d_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_February) {
$__t0 = 2;
goto end_branch_0;;
};
  if (($d_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_March) {
$__t0 = 3;
goto end_branch_0;;
};
  if (($d_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_April) {
$__t0 = 4;
goto end_branch_0;;
};
  if (($d_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_May) {
$__t0 = 5;
goto end_branch_0;;
};
  if (($d_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_June) {
$__t0 = 6;
goto end_branch_0;;
};
  if (($d_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_July) {
$__t0 = 7;
goto end_branch_0;;
};
  if (($d_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_August) {
$__t0 = 8;
goto end_branch_0;;
};
  if (($d_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_September) {
$__t0 = 9;
goto end_branch_0;;
};
  if (($d_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_October) {
$__t0 = 10;
goto end_branch_0;;
};
  if (($d_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_November) {
$__t0 = 11;
goto end_branch_0;;
};
  if (($d_0)->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_December) {
$__t0 = 12;
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = ($GLOBALS['Data_DateTime_Instant_fromDateTimeImpl'])(($d_0)->{'value0'}, $__t0, ($d_0)->{'value2'}, 0, 0, 0, 0);
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
  $__res = (($dictDuration_0)->{'toDuration'})(($dt1_1 + \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($x_3) {
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
}), $dt2_2)));
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_DateTime_Instant_diff'] = __NAMESPACE__ . '\\majData_majDatemajTime_majInstant_diff';

// Data_DateTime_Instant_boundedInstant
$GLOBALS['Data_DateTime_Instant_boundedInstant'] = (object)["bottom" => -8639977881600000.0, "top" => 8639977881599999.0, "Ord0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Ord_ordNumber'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

