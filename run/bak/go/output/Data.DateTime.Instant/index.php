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
\PhpursThunks::$thunks['Data_DateTime_Instant_negateDuration'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))['identity']))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Time_Duration_negate'] ?? \PhpursThunks::eval('Data_Time_Duration_negate'))))((($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))['identity'])); return $v; };
\PhpursThunks::$thunks['Data_DateTime_Instant_unInstant'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = $v_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_DateTime_Instant_toDateTime'] = function() { $v = (($GLOBALS['Data_DateTime_Instant_toDateTimeImpl'] ?? \PhpursThunks::eval('Data_DateTime_Instant_toDateTimeImpl')))((function() {
  $__fn = function($y_0 = null, $mo_1 = null, $d_2 = null, $h_3 = null, $mi_4 = null, $s_5 = null, $ms_6 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 7) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 7);
  }
  $__t0 = null;;
  switch ($mo_1) {
case 1:
$__t0 = new Phpurs_Data0("January");
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($mo_1) {
case 2:
$__t0 = new Phpurs_Data0("February");
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($mo_1) {
case 3:
$__t0 = new Phpurs_Data0("March");
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($mo_1) {
case 4:
$__t0 = new Phpurs_Data0("April");
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($mo_1) {
case 5:
$__t0 = new Phpurs_Data0("May");
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($mo_1) {
case 6:
$__t0 = new Phpurs_Data0("June");
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($mo_1) {
case 7:
$__t0 = new Phpurs_Data0("July");
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($mo_1) {
case 8:
$__t0 = new Phpurs_Data0("August");
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($mo_1) {
case 9:
$__t0 = new Phpurs_Data0("September");
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($mo_1) {
case 10:
$__t0 = new Phpurs_Data0("October");
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($mo_1) {
case 11:
$__t0 = new Phpurs_Data0("November");
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($mo_1) {
case 12:
$__t0 = new Phpurs_Data0("December");
goto end_branch_0;;
break;
default:
;
break;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = new Phpurs_Data2("DateTime", (((($GLOBALS['Data_Date_canonicalDate'] ?? \PhpursThunks::eval('Data_Date_canonicalDate')))($y_0))($__t0))($d_2), new Phpurs_Data4("Time", $h_3, $mi_4, $s_5, $ms_6));
  goto __end;;
  __end:
  return $__num > 7 ? $__res(...\array_slice(\func_get_args(), 7)) : $__res;
  };
  return $__fn;
})()); return $v; };
\PhpursThunks::$thunks['Data_DateTime_Instant_showInstant'] = function() { $v = ["show" => function($v_0 = null) {
  $__num = \func_num_args();
  $__res = (("(Instant (Milliseconds " . (($GLOBALS['Data_Show_showNumberImpl'] ?? \PhpursThunks::eval('Data_Show_showNumberImpl')))($v_0)) . "))");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_DateTime_Instant_ordDateTime'] = function() { $v = ($GLOBALS['Data_Ord_ordNumber'] ?? \PhpursThunks::eval('Data_Ord_ordNumber')); return $v; };
\PhpursThunks::$thunks['Data_DateTime_Instant_instant'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((($v_0 >= -8639977881600000.0) && ($v_0 <= 8639977881599999.0))) {
$__t0 = new Phpurs_Data1("Just", $v_0);
goto end_branch_0;;
};
  $__t0 = new Phpurs_Data0("Nothing");
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_DateTime_Instant_fromDateTime'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object((($v_0)->{'value0'})->{'value1'}) && (((($v_0)->{'value0'})->{'value1'})->{'tag'} === "January"))) {
$__t0 = 1;
goto end_branch_0;;
};
  if ((is_object((($v_0)->{'value0'})->{'value1'}) && (((($v_0)->{'value0'})->{'value1'})->{'tag'} === "February"))) {
$__t0 = 2;
goto end_branch_0;;
};
  if ((is_object((($v_0)->{'value0'})->{'value1'}) && (((($v_0)->{'value0'})->{'value1'})->{'tag'} === "March"))) {
$__t0 = 3;
goto end_branch_0;;
};
  if ((is_object((($v_0)->{'value0'})->{'value1'}) && (((($v_0)->{'value0'})->{'value1'})->{'tag'} === "April"))) {
$__t0 = 4;
goto end_branch_0;;
};
  if ((is_object((($v_0)->{'value0'})->{'value1'}) && (((($v_0)->{'value0'})->{'value1'})->{'tag'} === "May"))) {
$__t0 = 5;
goto end_branch_0;;
};
  if ((is_object((($v_0)->{'value0'})->{'value1'}) && (((($v_0)->{'value0'})->{'value1'})->{'tag'} === "June"))) {
$__t0 = 6;
goto end_branch_0;;
};
  if ((is_object((($v_0)->{'value0'})->{'value1'}) && (((($v_0)->{'value0'})->{'value1'})->{'tag'} === "July"))) {
$__t0 = 7;
goto end_branch_0;;
};
  if ((is_object((($v_0)->{'value0'})->{'value1'}) && (((($v_0)->{'value0'})->{'value1'})->{'tag'} === "August"))) {
$__t0 = 8;
goto end_branch_0;;
};
  if ((is_object((($v_0)->{'value0'})->{'value1'}) && (((($v_0)->{'value0'})->{'value1'})->{'tag'} === "September"))) {
$__t0 = 9;
goto end_branch_0;;
};
  if ((is_object((($v_0)->{'value0'})->{'value1'}) && (((($v_0)->{'value0'})->{'value1'})->{'tag'} === "October"))) {
$__t0 = 10;
goto end_branch_0;;
};
  if ((is_object((($v_0)->{'value0'})->{'value1'}) && (((($v_0)->{'value0'})->{'value1'})->{'tag'} === "November"))) {
$__t0 = 11;
goto end_branch_0;;
};
  if ((is_object((($v_0)->{'value0'})->{'value1'}) && (((($v_0)->{'value0'})->{'value1'})->{'tag'} === "December"))) {
$__t0 = 12;
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = (($GLOBALS['Data_DateTime_Instant_fromDateTimeImpl'] ?? \PhpursThunks::eval('Data_DateTime_Instant_fromDateTimeImpl')))((($v_0)->{'value0'})->{'value0'}, $__t0, (($v_0)->{'value0'})->{'value2'}, (($v_0)->{'value1'})->{'value0'}, (($v_0)->{'value1'})->{'value1'}, (($v_0)->{'value1'})->{'value2'}, (($v_0)->{'value1'})->{'value3'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_DateTime_Instant_fromDate'] = function() { $v = function($d_0 = null) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object(($d_0)->{'value1'}) && ((($d_0)->{'value1'})->{'tag'} === "January"))) {
$__t0 = 1;
goto end_branch_0;;
};
  if ((is_object(($d_0)->{'value1'}) && ((($d_0)->{'value1'})->{'tag'} === "February"))) {
$__t0 = 2;
goto end_branch_0;;
};
  if ((is_object(($d_0)->{'value1'}) && ((($d_0)->{'value1'})->{'tag'} === "March"))) {
$__t0 = 3;
goto end_branch_0;;
};
  if ((is_object(($d_0)->{'value1'}) && ((($d_0)->{'value1'})->{'tag'} === "April"))) {
$__t0 = 4;
goto end_branch_0;;
};
  if ((is_object(($d_0)->{'value1'}) && ((($d_0)->{'value1'})->{'tag'} === "May"))) {
$__t0 = 5;
goto end_branch_0;;
};
  if ((is_object(($d_0)->{'value1'}) && ((($d_0)->{'value1'})->{'tag'} === "June"))) {
$__t0 = 6;
goto end_branch_0;;
};
  if ((is_object(($d_0)->{'value1'}) && ((($d_0)->{'value1'})->{'tag'} === "July"))) {
$__t0 = 7;
goto end_branch_0;;
};
  if ((is_object(($d_0)->{'value1'}) && ((($d_0)->{'value1'})->{'tag'} === "August"))) {
$__t0 = 8;
goto end_branch_0;;
};
  if ((is_object(($d_0)->{'value1'}) && ((($d_0)->{'value1'})->{'tag'} === "September"))) {
$__t0 = 9;
goto end_branch_0;;
};
  if ((is_object(($d_0)->{'value1'}) && ((($d_0)->{'value1'})->{'tag'} === "October"))) {
$__t0 = 10;
goto end_branch_0;;
};
  if ((is_object(($d_0)->{'value1'}) && ((($d_0)->{'value1'})->{'tag'} === "November"))) {
$__t0 = 11;
goto end_branch_0;;
};
  if ((is_object(($d_0)->{'value1'}) && ((($d_0)->{'value1'})->{'tag'} === "December"))) {
$__t0 = 12;
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = (($GLOBALS['Data_DateTime_Instant_fromDateTimeImpl'] ?? \PhpursThunks::eval('Data_DateTime_Instant_fromDateTimeImpl')))(($d_0)->{'value0'}, $__t0, ($d_0)->{'value2'}, 0, 0, 0, 0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_DateTime_Instant_eqDateTime'] = function() { $v = ($GLOBALS['Data_Eq_eqNumber'] ?? \PhpursThunks::eval('Data_Eq_eqNumber')); return $v; };
\PhpursThunks::$thunks['Data_DateTime_Instant_diff'] = function() { $v = (function() {
  $__fn = function($dictDuration_0 = null, $dt1_1 = null, $dt2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($dictDuration_0)['toDuration'])(($dt1_1 + (($GLOBALS['Data_DateTime_Instant_negateDuration'] ?? \PhpursThunks::eval('Data_DateTime_Instant_negateDuration')))($dt2_2)));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_DateTime_Instant_boundedInstant'] = function() { $v = ["bottom" => -8639977881600000.0, "top" => 8639977881599999.0, "Ord0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Ord_ordNumber'] ?? \PhpursThunks::eval('Data_Ord_ordNumber'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };
$ffi_Data_DateTime_Instant = \call_user_func(function() {
  $exports = [];
$fromDateTimeImpl = function($y, $mo = null, $d = null, $h = null, $mi = null, $s = null, $ms = null) use (&$fromDateTimeImpl) {
    if (\func_num_args() < 7) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$fromDateTimeImpl) {

            return $fromDateTimeImpl(...\array_merge($__args, $more));
        };
    }
    $dt = new \DateTime('now', new \DateTimeZone('UTC'));
    $dt->setDate($y, $mo, $d);
    $dt->setTime($h, $mi, $s, $ms * 1000);
    return (float)$dt->getTimestamp() * 1000 + (int)$dt->format('v');
};

$toDateTimeImpl = function($ctor, $instant = null) use (&$toDateTimeImpl) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$toDateTimeImpl) {

            return $toDateTimeImpl(...\array_merge($__args, $more));
        };
    }
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
\PhpursThunks::$thunks['Data_DateTime_Instant_fromDateTimeImpl'] = function() use (&$ffi_Data_DateTime_Instant) { return $ffi_Data_DateTime_Instant['fromDateTimeImpl']; };
\PhpursThunks::$thunks['Data_DateTime_Instant_toDateTimeImpl'] = function() use (&$ffi_Data_DateTime_Instant) { return $ffi_Data_DateTime_Instant['toDateTimeImpl']; };













