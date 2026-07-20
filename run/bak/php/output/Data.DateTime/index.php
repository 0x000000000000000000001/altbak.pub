<?php

namespace Data\DateTime;

// ALL IMPORTS: Control.Apply, Control.Bind, Data.Bounded, Data.Date, Data.Date.Component, Data.DateTime, Data.Enum, Data.Eq, Data.Function, Data.Function.Uncurried, Data.Functor, Data.HeytingAlgebra, Data.Maybe, Data.Ord, Data.Ordering, Data.Semigroup, Data.Show, Data.Time, Data.Time.Component, Data.Time.Duration, Prelude, Prim
// TO REQUIRE: Control.Apply, Control.Bind, Data.Bounded, Data.Date, Data.Date.Component, Data.DateTime, Data.Enum, Data.Eq, Data.Function, Data.Function.Uncurried, Data.Functor, Data.HeytingAlgebra, Data.Maybe, Data.Ord, Data.Ordering, Data.Semigroup, Data.Show, Data.Time, Data.Time.Component, Data.Time.Duration, Prelude
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Data.Bounded/index.php';
require_once __DIR__ . '/../Data.Date/index.php';
require_once __DIR__ . '/../Data.Date.Component/index.php';
require_once __DIR__ . '/../Data.DateTime/index.php';
require_once __DIR__ . '/../Data.Enum/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Function.Uncurried/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.HeytingAlgebra/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ordering/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
require_once __DIR__ . '/../Data.Time/index.php';
require_once __DIR__ . '/../Data.Time.Component/index.php';
require_once __DIR__ . '/../Data.Time.Duration/index.php';
require_once __DIR__ . '/../Prelude/index.php';

if (!class_exists(__NAMESPACE__ . '\\Phpurs_Data0')) {
  class Phpurs_Data0 { public $tag; public function __construct($t) { $this->tag = $t; } }
  class Phpurs_Data1 { public $tag; public $value0; public function __construct($t, $value0) { $this->tag = $t; $this->value0 = $value0; } }
  class Phpurs_Data2 { public $tag; public $value0, $value1; public function __construct($t, $value0, $value1) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; } }
  class Phpurs_Data3 { public $tag; public $value0, $value1, $value2; public function __construct($t, $value0, $value1, $value2) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; } }
  class Phpurs_Data4 { public $tag; public $value0, $value1, $value2, $value3; public function __construct($t, $value0, $value1, $value2, $value3) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; $this->value3 = $value3; } }
  class Phpurs_Data5 { public $tag; public $value0, $value1, $value2, $value3, $value4; public function __construct($t, $value0, $value1, $value2, $value3, $value4) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; $this->value3 = $value3; $this->value4 = $value4; } }
  class Phpurs_Data6 { public $tag; public $value0, $value1, $value2, $value3, $value4, $value5; public function __construct($t, $value0, $value1, $value2, $value3, $value4, $value5) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; $this->value3 = $value3; $this->value4 = $value4; $this->value5 = $value5; } }
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
\PhpursThunks::$thunks['Data_DateTime_DateTime'] = function() { $v = (function() {
  $__fn = function($value0, $value1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("DateTime", $value0, $value1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_DateTime_toRecord'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_DateTime_toRecord"), recVars=[];
  if ((is_object((($v_0)->value0)->value1) && (((($v_0)->value0)->value1)->tag === "January"))) {
$__t0 = 1;
} else {
if ((is_object((($v_0)->value0)->value1) && (((($v_0)->value0)->value1)->tag === "February"))) {
$__t0 = 2;
} else {
if ((is_object((($v_0)->value0)->value1) && (((($v_0)->value0)->value1)->tag === "March"))) {
$__t0 = 3;
} else {
if ((is_object((($v_0)->value0)->value1) && (((($v_0)->value0)->value1)->tag === "April"))) {
$__t0 = 4;
} else {
if ((is_object((($v_0)->value0)->value1) && (((($v_0)->value0)->value1)->tag === "May"))) {
$__t0 = 5;
} else {
if ((is_object((($v_0)->value0)->value1) && (((($v_0)->value0)->value1)->tag === "June"))) {
$__t0 = 6;
} else {
if ((is_object((($v_0)->value0)->value1) && (((($v_0)->value0)->value1)->tag === "July"))) {
$__t0 = 7;
} else {
if ((is_object((($v_0)->value0)->value1) && (((($v_0)->value0)->value1)->tag === "August"))) {
$__t0 = 8;
} else {
if ((is_object((($v_0)->value0)->value1) && (((($v_0)->value0)->value1)->tag === "September"))) {
$__t0 = 9;
} else {
if ((is_object((($v_0)->value0)->value1) && (((($v_0)->value0)->value1)->tag === "October"))) {
$__t0 = 10;
} else {
if ((is_object((($v_0)->value0)->value1) && (((($v_0)->value0)->value1)->tag === "November"))) {
$__t0 = 11;
} else {
if ((is_object((($v_0)->value0)->value1) && (((($v_0)->value0)->value1)->tag === "December"))) {
$__t0 = 12;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
};
};
};
};
};
};
};
};
};
};
  $__res = (object)["year" => (($v_0)->value0)->value0, "month" => $__t0, "day" => (($v_0)->value0)->value2, "hour" => (($v_0)->value1)->value0, "minute" => (($v_0)->value1)->value1, "second" => (($v_0)->value1)->value2, "millisecond" => (($v_0)->value1)->value3];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_DateTime_time'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_DateTime_time"), recVars=[];
  $__res = ($v_0)->value1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_DateTime_showDateTime'] = function() { $v = (object)["show" => function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))("(DateTime "))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))(((($GLOBALS['Data_Date_showDate'] ?? \PhpursThunks::eval('Data_Date_showDate')))->show)(($v_0)->value0)))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))(" "))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))(((($GLOBALS['Data_Time_showTime'] ?? \PhpursThunks::eval('Data_Time_showTime')))->show)(($v_0)->value1)))(")"))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_DateTime_modifyTimeF'] = function() { $v = (function() {
  $__fn = function($dictFunctor_0, $f_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_DateTime_modifyTimeF"), recVars=[];
  $__res = ((($dictFunctor_0)->map)((($GLOBALS['Data_DateTime_DateTime'] ?? \PhpursThunks::eval('Data_DateTime_DateTime')))(($v_2)->value0)))(($f_1)(($v_2)->value1));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_DateTime_modifyTime'] = function() { $v = (function() {
  $__fn = function($f_0, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_DateTime_modifyTime"), recVars=[];
  $__res = new Phpurs_Data2("DateTime", ($v_1)->value0, ($f_0)(($v_1)->value1));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_DateTime_modifyDateF'] = function() { $v = (function() {
  $__fn = function($dictFunctor_0, $f_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_DateTime_modifyDateF"), recVars=[];
  $__local_var_3_0 = ($v_2)->value1;
  $__res = ((($dictFunctor_0)->map)(function($a_4) use ($__local_var_3_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = new Phpurs_Data2("DateTime", $a_4, $__local_var_3_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($f_1)(($v_2)->value0));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_DateTime_modifyDate'] = function() { $v = (function() {
  $__fn = function($f_0, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_DateTime_modifyDate"), recVars=[];
  $__res = new Phpurs_Data2("DateTime", ($f_0)(($v_1)->value0), ($v_1)->value1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_DateTime_eqDateTime'] = function() { $v = (object)["eq" => (function() {
  $__fn = function($x_0, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Data_HeytingAlgebra_boolConj'] ?? \PhpursThunks::eval('Data_HeytingAlgebra_boolConj')))((((($GLOBALS['Data_Date_eqDate'] ?? \PhpursThunks::eval('Data_Date_eqDate')))->eq)(($x_0)->value0))(($y_1)->value0)))((((($GLOBALS['Data_Time_eqTime'] ?? \PhpursThunks::eval('Data_Time_eqTime')))->eq)(($x_0)->value1))(($y_1)->value1));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Data_DateTime_ordDateTime'] = function() { $v = (object)["compare" => (function() {
  $__fn = function($x_0, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $v_2_0 = (((($GLOBALS['Data_Date_ordDate'] ?? \PhpursThunks::eval('Data_Date_ordDate')))->compare)(($x_0)->value0))(($y_1)->value0);
  if ((is_object($v_2_0) && (($v_2_0)->tag === "LT"))) {
$__t1 = new Phpurs_Data0("LT");
} else {
if ((is_object($v_2_0) && (($v_2_0)->tag === "GT"))) {
$__t1 = new Phpurs_Data0("GT");
} else {
$__t1 = (((($GLOBALS['Data_Time_ordTime'] ?? \PhpursThunks::eval('Data_Time_ordTime')))->compare)(($x_0)->value1))(($y_1)->value1);
};
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_DateTime_eqDateTime'] ?? \PhpursThunks::eval('Data_DateTime_eqDateTime'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_DateTime_diff'] = function() { $v = (function() {
  $__fn = function($dictDuration_0, $dt1_1 = null, $dt2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_DateTime_diff"), recVars=[];
  $__res = (($dictDuration_0)->toDuration)((((($GLOBALS['Data_Function_Uncurried_runFn2'] ?? \PhpursThunks::eval('Data_Function_Uncurried_runFn2')))(($GLOBALS['Data_DateTime_calcDiff'] ?? \PhpursThunks::eval('Data_DateTime_calcDiff'))))((($GLOBALS['Data_DateTime_toRecord'] ?? \PhpursThunks::eval('Data_DateTime_toRecord')))($dt1_1)))((($GLOBALS['Data_DateTime_toRecord'] ?? \PhpursThunks::eval('Data_DateTime_toRecord')))($dt2_2)));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_DateTime_date'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_DateTime_date"), recVars=[];
  $__res = ($v_0)->value0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_DateTime_boundedDateTime'] = function() { $v = (object)["bottom" => new Phpurs_Data2("DateTime", new Phpurs_Data3("Date", (($GLOBALS['Data_Date_Component_boundedYear'] ?? \PhpursThunks::eval('Data_Date_Component_boundedYear')))->bottom, new Phpurs_Data0("January"), 1), new Phpurs_Data4("Time", 0, 0, 0, 0)), "top" => new Phpurs_Data2("DateTime", new Phpurs_Data3("Date", 275759, new Phpurs_Data0("December"), 31), new Phpurs_Data4("Time", 23, 59, 59, 999)), "Ord0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_DateTime_ordDateTime'] ?? \PhpursThunks::eval('Data_DateTime_ordDateTime'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_DateTime_adjust'] = function() { $v = (function() {
  $__fn = function($dictDuration_0, $d_1 = null, $dt_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_DateTime_adjust"), recVars=[];
  $__local_var_3_0 = ((((($GLOBALS['Data_DateTime_adjustImpl'] ?? \PhpursThunks::eval('Data_DateTime_adjustImpl')))(($GLOBALS['Data_Maybe_Just'] ?? \PhpursThunks::eval('Data_Maybe_Just'))))(new Phpurs_Data0("Nothing")))((($dictDuration_0)->fromDuration)($d_1)))((($GLOBALS['Data_DateTime_toRecord'] ?? \PhpursThunks::eval('Data_DateTime_toRecord')))($dt_2));
  if ((is_object($__local_var_3_0) && (($__local_var_3_0)->tag === "Just"))) {
$__local_var_4_2 = ((($GLOBALS['Data_Date_Component_boundedEnumYear'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumYear')))->toEnum)((($__local_var_3_0)->value0)->year);
if ((is_object($__local_var_4_2) && (($__local_var_4_2)->tag === "Just"))) {
$__t3 = new Phpurs_Data1("Just", (($GLOBALS['Data_Date_exactDate'] ?? \PhpursThunks::eval('Data_Date_exactDate')))(($__local_var_4_2)->value0));
} else {
$__t3 = new Phpurs_Data0("Nothing");
};
$__local_var_5_3 = $__t3;
if (((($__local_var_3_0)->value0)->month === 1)) {
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->tag === "Just"))) {
$__local_var_6_7 = ((($GLOBALS['Data_Date_Component_boundedEnumDay'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumDay')))->toEnum)((($__local_var_3_0)->value0)->day);
if ((is_object($__local_var_6_7) && (($__local_var_6_7)->tag === "Just"))) {
$__local_var_7_9 = ((($__local_var_5_3)->value0)(new Phpurs_Data0("January")))(($__local_var_6_7)->value0);
if ((is_object($__local_var_7_9) && (($__local_var_7_9)->tag === "Just"))) {
$__local_var_8_11 = ((($GLOBALS['Data_Time_Component_boundedEnumHour'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumHour')))->toEnum)((($__local_var_3_0)->value0)->hour);
if ((is_object($__local_var_8_11) && (($__local_var_8_11)->tag === "Just"))) {
$__local_var_9_13 = ((($GLOBALS['Data_Time_Component_boundedEnumMinute'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMinute')))->toEnum)((($__local_var_3_0)->value0)->minute);
if ((is_object($__local_var_9_13) && (($__local_var_9_13)->tag === "Just"))) {
$__local_var_10_15 = ((($GLOBALS['Data_Time_Component_boundedEnumSecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumSecond')))->toEnum)((($__local_var_3_0)->value0)->second);
if ((is_object($__local_var_10_15) && (($__local_var_10_15)->tag === "Just"))) {
$__local_var_11_17 = ((($GLOBALS['Data_Time_Component_boundedEnumMillisecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMillisecond')))->toEnum)((($__local_var_3_0)->value0)->millisecond);
if ((is_object($__local_var_11_17) && (($__local_var_11_17)->tag === "Just"))) {
$__t18 = new Phpurs_Data1("Just", new Phpurs_Data2("DateTime", ($__local_var_7_9)->value0, new Phpurs_Data4("Time", ($__local_var_8_11)->value0, ($__local_var_9_13)->value0, ($__local_var_10_15)->value0, ($__local_var_11_17)->value0)));
} else {
$__t18 = new Phpurs_Data0("Nothing");
};
$__t16 = $__t18;
} else {
$__t16 = new Phpurs_Data0("Nothing");
};
$__t14 = $__t16;
} else {
$__t14 = new Phpurs_Data0("Nothing");
};
$__t12 = $__t14;
} else {
$__t12 = new Phpurs_Data0("Nothing");
};
$__t10 = $__t12;
} else {
$__t10 = new Phpurs_Data0("Nothing");
};
$__t8 = $__t10;
} else {
$__t8 = new Phpurs_Data0("Nothing");
};
$__t6 = $__t8;
} else {
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->tag === "Nothing"))) {
$__t6 = new Phpurs_Data0("Nothing");
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t6 = null;
};
};
$__t5 = $__t6;
} else {
if (((($__local_var_3_0)->value0)->month === 2)) {
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->tag === "Just"))) {
$__local_var_6_20 = ((($GLOBALS['Data_Date_Component_boundedEnumDay'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumDay')))->toEnum)((($__local_var_3_0)->value0)->day);
if ((is_object($__local_var_6_20) && (($__local_var_6_20)->tag === "Just"))) {
$__local_var_7_22 = ((($__local_var_5_3)->value0)(new Phpurs_Data0("February")))(($__local_var_6_20)->value0);
if ((is_object($__local_var_7_22) && (($__local_var_7_22)->tag === "Just"))) {
$__local_var_8_24 = ((($GLOBALS['Data_Time_Component_boundedEnumHour'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumHour')))->toEnum)((($__local_var_3_0)->value0)->hour);
if ((is_object($__local_var_8_24) && (($__local_var_8_24)->tag === "Just"))) {
$__local_var_9_26 = ((($GLOBALS['Data_Time_Component_boundedEnumMinute'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMinute')))->toEnum)((($__local_var_3_0)->value0)->minute);
if ((is_object($__local_var_9_26) && (($__local_var_9_26)->tag === "Just"))) {
$__local_var_10_28 = ((($GLOBALS['Data_Time_Component_boundedEnumSecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumSecond')))->toEnum)((($__local_var_3_0)->value0)->second);
if ((is_object($__local_var_10_28) && (($__local_var_10_28)->tag === "Just"))) {
$__local_var_11_30 = ((($GLOBALS['Data_Time_Component_boundedEnumMillisecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMillisecond')))->toEnum)((($__local_var_3_0)->value0)->millisecond);
if ((is_object($__local_var_11_30) && (($__local_var_11_30)->tag === "Just"))) {
$__t31 = new Phpurs_Data1("Just", new Phpurs_Data2("DateTime", ($__local_var_7_22)->value0, new Phpurs_Data4("Time", ($__local_var_8_24)->value0, ($__local_var_9_26)->value0, ($__local_var_10_28)->value0, ($__local_var_11_30)->value0)));
} else {
$__t31 = new Phpurs_Data0("Nothing");
};
$__t29 = $__t31;
} else {
$__t29 = new Phpurs_Data0("Nothing");
};
$__t27 = $__t29;
} else {
$__t27 = new Phpurs_Data0("Nothing");
};
$__t25 = $__t27;
} else {
$__t25 = new Phpurs_Data0("Nothing");
};
$__t23 = $__t25;
} else {
$__t23 = new Phpurs_Data0("Nothing");
};
$__t21 = $__t23;
} else {
$__t21 = new Phpurs_Data0("Nothing");
};
$__t19 = $__t21;
} else {
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->tag === "Nothing"))) {
$__t19 = new Phpurs_Data0("Nothing");
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t19 = null;
};
};
$__t5 = $__t19;
} else {
if (((($__local_var_3_0)->value0)->month === 3)) {
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->tag === "Just"))) {
$__local_var_6_33 = ((($GLOBALS['Data_Date_Component_boundedEnumDay'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumDay')))->toEnum)((($__local_var_3_0)->value0)->day);
if ((is_object($__local_var_6_33) && (($__local_var_6_33)->tag === "Just"))) {
$__local_var_7_35 = ((($__local_var_5_3)->value0)(new Phpurs_Data0("March")))(($__local_var_6_33)->value0);
if ((is_object($__local_var_7_35) && (($__local_var_7_35)->tag === "Just"))) {
$__local_var_8_37 = ((($GLOBALS['Data_Time_Component_boundedEnumHour'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumHour')))->toEnum)((($__local_var_3_0)->value0)->hour);
if ((is_object($__local_var_8_37) && (($__local_var_8_37)->tag === "Just"))) {
$__local_var_9_39 = ((($GLOBALS['Data_Time_Component_boundedEnumMinute'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMinute')))->toEnum)((($__local_var_3_0)->value0)->minute);
if ((is_object($__local_var_9_39) && (($__local_var_9_39)->tag === "Just"))) {
$__local_var_10_41 = ((($GLOBALS['Data_Time_Component_boundedEnumSecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumSecond')))->toEnum)((($__local_var_3_0)->value0)->second);
if ((is_object($__local_var_10_41) && (($__local_var_10_41)->tag === "Just"))) {
$__local_var_11_43 = ((($GLOBALS['Data_Time_Component_boundedEnumMillisecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMillisecond')))->toEnum)((($__local_var_3_0)->value0)->millisecond);
if ((is_object($__local_var_11_43) && (($__local_var_11_43)->tag === "Just"))) {
$__t44 = new Phpurs_Data1("Just", new Phpurs_Data2("DateTime", ($__local_var_7_35)->value0, new Phpurs_Data4("Time", ($__local_var_8_37)->value0, ($__local_var_9_39)->value0, ($__local_var_10_41)->value0, ($__local_var_11_43)->value0)));
} else {
$__t44 = new Phpurs_Data0("Nothing");
};
$__t42 = $__t44;
} else {
$__t42 = new Phpurs_Data0("Nothing");
};
$__t40 = $__t42;
} else {
$__t40 = new Phpurs_Data0("Nothing");
};
$__t38 = $__t40;
} else {
$__t38 = new Phpurs_Data0("Nothing");
};
$__t36 = $__t38;
} else {
$__t36 = new Phpurs_Data0("Nothing");
};
$__t34 = $__t36;
} else {
$__t34 = new Phpurs_Data0("Nothing");
};
$__t32 = $__t34;
} else {
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->tag === "Nothing"))) {
$__t32 = new Phpurs_Data0("Nothing");
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t32 = null;
};
};
$__t5 = $__t32;
} else {
if (((($__local_var_3_0)->value0)->month === 4)) {
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->tag === "Just"))) {
$__local_var_6_46 = ((($GLOBALS['Data_Date_Component_boundedEnumDay'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumDay')))->toEnum)((($__local_var_3_0)->value0)->day);
if ((is_object($__local_var_6_46) && (($__local_var_6_46)->tag === "Just"))) {
$__local_var_7_48 = ((($__local_var_5_3)->value0)(new Phpurs_Data0("April")))(($__local_var_6_46)->value0);
if ((is_object($__local_var_7_48) && (($__local_var_7_48)->tag === "Just"))) {
$__local_var_8_50 = ((($GLOBALS['Data_Time_Component_boundedEnumHour'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumHour')))->toEnum)((($__local_var_3_0)->value0)->hour);
if ((is_object($__local_var_8_50) && (($__local_var_8_50)->tag === "Just"))) {
$__local_var_9_52 = ((($GLOBALS['Data_Time_Component_boundedEnumMinute'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMinute')))->toEnum)((($__local_var_3_0)->value0)->minute);
if ((is_object($__local_var_9_52) && (($__local_var_9_52)->tag === "Just"))) {
$__local_var_10_54 = ((($GLOBALS['Data_Time_Component_boundedEnumSecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumSecond')))->toEnum)((($__local_var_3_0)->value0)->second);
if ((is_object($__local_var_10_54) && (($__local_var_10_54)->tag === "Just"))) {
$__local_var_11_56 = ((($GLOBALS['Data_Time_Component_boundedEnumMillisecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMillisecond')))->toEnum)((($__local_var_3_0)->value0)->millisecond);
if ((is_object($__local_var_11_56) && (($__local_var_11_56)->tag === "Just"))) {
$__t57 = new Phpurs_Data1("Just", new Phpurs_Data2("DateTime", ($__local_var_7_48)->value0, new Phpurs_Data4("Time", ($__local_var_8_50)->value0, ($__local_var_9_52)->value0, ($__local_var_10_54)->value0, ($__local_var_11_56)->value0)));
} else {
$__t57 = new Phpurs_Data0("Nothing");
};
$__t55 = $__t57;
} else {
$__t55 = new Phpurs_Data0("Nothing");
};
$__t53 = $__t55;
} else {
$__t53 = new Phpurs_Data0("Nothing");
};
$__t51 = $__t53;
} else {
$__t51 = new Phpurs_Data0("Nothing");
};
$__t49 = $__t51;
} else {
$__t49 = new Phpurs_Data0("Nothing");
};
$__t47 = $__t49;
} else {
$__t47 = new Phpurs_Data0("Nothing");
};
$__t45 = $__t47;
} else {
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->tag === "Nothing"))) {
$__t45 = new Phpurs_Data0("Nothing");
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t45 = null;
};
};
$__t5 = $__t45;
} else {
if (((($__local_var_3_0)->value0)->month === 5)) {
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->tag === "Just"))) {
$__local_var_6_59 = ((($GLOBALS['Data_Date_Component_boundedEnumDay'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumDay')))->toEnum)((($__local_var_3_0)->value0)->day);
if ((is_object($__local_var_6_59) && (($__local_var_6_59)->tag === "Just"))) {
$__local_var_7_61 = ((($__local_var_5_3)->value0)(new Phpurs_Data0("May")))(($__local_var_6_59)->value0);
if ((is_object($__local_var_7_61) && (($__local_var_7_61)->tag === "Just"))) {
$__local_var_8_63 = ((($GLOBALS['Data_Time_Component_boundedEnumHour'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumHour')))->toEnum)((($__local_var_3_0)->value0)->hour);
if ((is_object($__local_var_8_63) && (($__local_var_8_63)->tag === "Just"))) {
$__local_var_9_65 = ((($GLOBALS['Data_Time_Component_boundedEnumMinute'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMinute')))->toEnum)((($__local_var_3_0)->value0)->minute);
if ((is_object($__local_var_9_65) && (($__local_var_9_65)->tag === "Just"))) {
$__local_var_10_67 = ((($GLOBALS['Data_Time_Component_boundedEnumSecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumSecond')))->toEnum)((($__local_var_3_0)->value0)->second);
if ((is_object($__local_var_10_67) && (($__local_var_10_67)->tag === "Just"))) {
$__local_var_11_69 = ((($GLOBALS['Data_Time_Component_boundedEnumMillisecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMillisecond')))->toEnum)((($__local_var_3_0)->value0)->millisecond);
if ((is_object($__local_var_11_69) && (($__local_var_11_69)->tag === "Just"))) {
$__t70 = new Phpurs_Data1("Just", new Phpurs_Data2("DateTime", ($__local_var_7_61)->value0, new Phpurs_Data4("Time", ($__local_var_8_63)->value0, ($__local_var_9_65)->value0, ($__local_var_10_67)->value0, ($__local_var_11_69)->value0)));
} else {
$__t70 = new Phpurs_Data0("Nothing");
};
$__t68 = $__t70;
} else {
$__t68 = new Phpurs_Data0("Nothing");
};
$__t66 = $__t68;
} else {
$__t66 = new Phpurs_Data0("Nothing");
};
$__t64 = $__t66;
} else {
$__t64 = new Phpurs_Data0("Nothing");
};
$__t62 = $__t64;
} else {
$__t62 = new Phpurs_Data0("Nothing");
};
$__t60 = $__t62;
} else {
$__t60 = new Phpurs_Data0("Nothing");
};
$__t58 = $__t60;
} else {
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->tag === "Nothing"))) {
$__t58 = new Phpurs_Data0("Nothing");
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t58 = null;
};
};
$__t5 = $__t58;
} else {
if (((($__local_var_3_0)->value0)->month === 6)) {
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->tag === "Just"))) {
$__local_var_6_72 = ((($GLOBALS['Data_Date_Component_boundedEnumDay'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumDay')))->toEnum)((($__local_var_3_0)->value0)->day);
if ((is_object($__local_var_6_72) && (($__local_var_6_72)->tag === "Just"))) {
$__local_var_7_74 = ((($__local_var_5_3)->value0)(new Phpurs_Data0("June")))(($__local_var_6_72)->value0);
if ((is_object($__local_var_7_74) && (($__local_var_7_74)->tag === "Just"))) {
$__local_var_8_76 = ((($GLOBALS['Data_Time_Component_boundedEnumHour'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumHour')))->toEnum)((($__local_var_3_0)->value0)->hour);
if ((is_object($__local_var_8_76) && (($__local_var_8_76)->tag === "Just"))) {
$__local_var_9_78 = ((($GLOBALS['Data_Time_Component_boundedEnumMinute'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMinute')))->toEnum)((($__local_var_3_0)->value0)->minute);
if ((is_object($__local_var_9_78) && (($__local_var_9_78)->tag === "Just"))) {
$__local_var_10_80 = ((($GLOBALS['Data_Time_Component_boundedEnumSecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumSecond')))->toEnum)((($__local_var_3_0)->value0)->second);
if ((is_object($__local_var_10_80) && (($__local_var_10_80)->tag === "Just"))) {
$__local_var_11_82 = ((($GLOBALS['Data_Time_Component_boundedEnumMillisecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMillisecond')))->toEnum)((($__local_var_3_0)->value0)->millisecond);
if ((is_object($__local_var_11_82) && (($__local_var_11_82)->tag === "Just"))) {
$__t83 = new Phpurs_Data1("Just", new Phpurs_Data2("DateTime", ($__local_var_7_74)->value0, new Phpurs_Data4("Time", ($__local_var_8_76)->value0, ($__local_var_9_78)->value0, ($__local_var_10_80)->value0, ($__local_var_11_82)->value0)));
} else {
$__t83 = new Phpurs_Data0("Nothing");
};
$__t81 = $__t83;
} else {
$__t81 = new Phpurs_Data0("Nothing");
};
$__t79 = $__t81;
} else {
$__t79 = new Phpurs_Data0("Nothing");
};
$__t77 = $__t79;
} else {
$__t77 = new Phpurs_Data0("Nothing");
};
$__t75 = $__t77;
} else {
$__t75 = new Phpurs_Data0("Nothing");
};
$__t73 = $__t75;
} else {
$__t73 = new Phpurs_Data0("Nothing");
};
$__t71 = $__t73;
} else {
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->tag === "Nothing"))) {
$__t71 = new Phpurs_Data0("Nothing");
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t71 = null;
};
};
$__t5 = $__t71;
} else {
if (((($__local_var_3_0)->value0)->month === 7)) {
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->tag === "Just"))) {
$__local_var_6_85 = ((($GLOBALS['Data_Date_Component_boundedEnumDay'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumDay')))->toEnum)((($__local_var_3_0)->value0)->day);
if ((is_object($__local_var_6_85) && (($__local_var_6_85)->tag === "Just"))) {
$__local_var_7_87 = ((($__local_var_5_3)->value0)(new Phpurs_Data0("July")))(($__local_var_6_85)->value0);
if ((is_object($__local_var_7_87) && (($__local_var_7_87)->tag === "Just"))) {
$__local_var_8_89 = ((($GLOBALS['Data_Time_Component_boundedEnumHour'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumHour')))->toEnum)((($__local_var_3_0)->value0)->hour);
if ((is_object($__local_var_8_89) && (($__local_var_8_89)->tag === "Just"))) {
$__local_var_9_91 = ((($GLOBALS['Data_Time_Component_boundedEnumMinute'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMinute')))->toEnum)((($__local_var_3_0)->value0)->minute);
if ((is_object($__local_var_9_91) && (($__local_var_9_91)->tag === "Just"))) {
$__local_var_10_93 = ((($GLOBALS['Data_Time_Component_boundedEnumSecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumSecond')))->toEnum)((($__local_var_3_0)->value0)->second);
if ((is_object($__local_var_10_93) && (($__local_var_10_93)->tag === "Just"))) {
$__local_var_11_95 = ((($GLOBALS['Data_Time_Component_boundedEnumMillisecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMillisecond')))->toEnum)((($__local_var_3_0)->value0)->millisecond);
if ((is_object($__local_var_11_95) && (($__local_var_11_95)->tag === "Just"))) {
$__t96 = new Phpurs_Data1("Just", new Phpurs_Data2("DateTime", ($__local_var_7_87)->value0, new Phpurs_Data4("Time", ($__local_var_8_89)->value0, ($__local_var_9_91)->value0, ($__local_var_10_93)->value0, ($__local_var_11_95)->value0)));
} else {
$__t96 = new Phpurs_Data0("Nothing");
};
$__t94 = $__t96;
} else {
$__t94 = new Phpurs_Data0("Nothing");
};
$__t92 = $__t94;
} else {
$__t92 = new Phpurs_Data0("Nothing");
};
$__t90 = $__t92;
} else {
$__t90 = new Phpurs_Data0("Nothing");
};
$__t88 = $__t90;
} else {
$__t88 = new Phpurs_Data0("Nothing");
};
$__t86 = $__t88;
} else {
$__t86 = new Phpurs_Data0("Nothing");
};
$__t84 = $__t86;
} else {
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->tag === "Nothing"))) {
$__t84 = new Phpurs_Data0("Nothing");
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t84 = null;
};
};
$__t5 = $__t84;
} else {
if (((($__local_var_3_0)->value0)->month === 8)) {
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->tag === "Just"))) {
$__local_var_6_98 = ((($GLOBALS['Data_Date_Component_boundedEnumDay'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumDay')))->toEnum)((($__local_var_3_0)->value0)->day);
if ((is_object($__local_var_6_98) && (($__local_var_6_98)->tag === "Just"))) {
$__local_var_7_100 = ((($__local_var_5_3)->value0)(new Phpurs_Data0("August")))(($__local_var_6_98)->value0);
if ((is_object($__local_var_7_100) && (($__local_var_7_100)->tag === "Just"))) {
$__local_var_8_102 = ((($GLOBALS['Data_Time_Component_boundedEnumHour'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumHour')))->toEnum)((($__local_var_3_0)->value0)->hour);
if ((is_object($__local_var_8_102) && (($__local_var_8_102)->tag === "Just"))) {
$__local_var_9_104 = ((($GLOBALS['Data_Time_Component_boundedEnumMinute'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMinute')))->toEnum)((($__local_var_3_0)->value0)->minute);
if ((is_object($__local_var_9_104) && (($__local_var_9_104)->tag === "Just"))) {
$__local_var_10_106 = ((($GLOBALS['Data_Time_Component_boundedEnumSecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumSecond')))->toEnum)((($__local_var_3_0)->value0)->second);
if ((is_object($__local_var_10_106) && (($__local_var_10_106)->tag === "Just"))) {
$__local_var_11_108 = ((($GLOBALS['Data_Time_Component_boundedEnumMillisecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMillisecond')))->toEnum)((($__local_var_3_0)->value0)->millisecond);
if ((is_object($__local_var_11_108) && (($__local_var_11_108)->tag === "Just"))) {
$__t109 = new Phpurs_Data1("Just", new Phpurs_Data2("DateTime", ($__local_var_7_100)->value0, new Phpurs_Data4("Time", ($__local_var_8_102)->value0, ($__local_var_9_104)->value0, ($__local_var_10_106)->value0, ($__local_var_11_108)->value0)));
} else {
$__t109 = new Phpurs_Data0("Nothing");
};
$__t107 = $__t109;
} else {
$__t107 = new Phpurs_Data0("Nothing");
};
$__t105 = $__t107;
} else {
$__t105 = new Phpurs_Data0("Nothing");
};
$__t103 = $__t105;
} else {
$__t103 = new Phpurs_Data0("Nothing");
};
$__t101 = $__t103;
} else {
$__t101 = new Phpurs_Data0("Nothing");
};
$__t99 = $__t101;
} else {
$__t99 = new Phpurs_Data0("Nothing");
};
$__t97 = $__t99;
} else {
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->tag === "Nothing"))) {
$__t97 = new Phpurs_Data0("Nothing");
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t97 = null;
};
};
$__t5 = $__t97;
} else {
if (((($__local_var_3_0)->value0)->month === 9)) {
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->tag === "Just"))) {
$__local_var_6_111 = ((($GLOBALS['Data_Date_Component_boundedEnumDay'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumDay')))->toEnum)((($__local_var_3_0)->value0)->day);
if ((is_object($__local_var_6_111) && (($__local_var_6_111)->tag === "Just"))) {
$__local_var_7_113 = ((($__local_var_5_3)->value0)(new Phpurs_Data0("September")))(($__local_var_6_111)->value0);
if ((is_object($__local_var_7_113) && (($__local_var_7_113)->tag === "Just"))) {
$__local_var_8_115 = ((($GLOBALS['Data_Time_Component_boundedEnumHour'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumHour')))->toEnum)((($__local_var_3_0)->value0)->hour);
if ((is_object($__local_var_8_115) && (($__local_var_8_115)->tag === "Just"))) {
$__local_var_9_117 = ((($GLOBALS['Data_Time_Component_boundedEnumMinute'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMinute')))->toEnum)((($__local_var_3_0)->value0)->minute);
if ((is_object($__local_var_9_117) && (($__local_var_9_117)->tag === "Just"))) {
$__local_var_10_119 = ((($GLOBALS['Data_Time_Component_boundedEnumSecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumSecond')))->toEnum)((($__local_var_3_0)->value0)->second);
if ((is_object($__local_var_10_119) && (($__local_var_10_119)->tag === "Just"))) {
$__local_var_11_121 = ((($GLOBALS['Data_Time_Component_boundedEnumMillisecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMillisecond')))->toEnum)((($__local_var_3_0)->value0)->millisecond);
if ((is_object($__local_var_11_121) && (($__local_var_11_121)->tag === "Just"))) {
$__t122 = new Phpurs_Data1("Just", new Phpurs_Data2("DateTime", ($__local_var_7_113)->value0, new Phpurs_Data4("Time", ($__local_var_8_115)->value0, ($__local_var_9_117)->value0, ($__local_var_10_119)->value0, ($__local_var_11_121)->value0)));
} else {
$__t122 = new Phpurs_Data0("Nothing");
};
$__t120 = $__t122;
} else {
$__t120 = new Phpurs_Data0("Nothing");
};
$__t118 = $__t120;
} else {
$__t118 = new Phpurs_Data0("Nothing");
};
$__t116 = $__t118;
} else {
$__t116 = new Phpurs_Data0("Nothing");
};
$__t114 = $__t116;
} else {
$__t114 = new Phpurs_Data0("Nothing");
};
$__t112 = $__t114;
} else {
$__t112 = new Phpurs_Data0("Nothing");
};
$__t110 = $__t112;
} else {
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->tag === "Nothing"))) {
$__t110 = new Phpurs_Data0("Nothing");
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t110 = null;
};
};
$__t5 = $__t110;
} else {
if (((($__local_var_3_0)->value0)->month === 10)) {
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->tag === "Just"))) {
$__local_var_6_124 = ((($GLOBALS['Data_Date_Component_boundedEnumDay'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumDay')))->toEnum)((($__local_var_3_0)->value0)->day);
if ((is_object($__local_var_6_124) && (($__local_var_6_124)->tag === "Just"))) {
$__local_var_7_126 = ((($__local_var_5_3)->value0)(new Phpurs_Data0("October")))(($__local_var_6_124)->value0);
if ((is_object($__local_var_7_126) && (($__local_var_7_126)->tag === "Just"))) {
$__local_var_8_128 = ((($GLOBALS['Data_Time_Component_boundedEnumHour'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumHour')))->toEnum)((($__local_var_3_0)->value0)->hour);
if ((is_object($__local_var_8_128) && (($__local_var_8_128)->tag === "Just"))) {
$__local_var_9_130 = ((($GLOBALS['Data_Time_Component_boundedEnumMinute'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMinute')))->toEnum)((($__local_var_3_0)->value0)->minute);
if ((is_object($__local_var_9_130) && (($__local_var_9_130)->tag === "Just"))) {
$__local_var_10_132 = ((($GLOBALS['Data_Time_Component_boundedEnumSecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumSecond')))->toEnum)((($__local_var_3_0)->value0)->second);
if ((is_object($__local_var_10_132) && (($__local_var_10_132)->tag === "Just"))) {
$__local_var_11_134 = ((($GLOBALS['Data_Time_Component_boundedEnumMillisecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMillisecond')))->toEnum)((($__local_var_3_0)->value0)->millisecond);
if ((is_object($__local_var_11_134) && (($__local_var_11_134)->tag === "Just"))) {
$__t135 = new Phpurs_Data1("Just", new Phpurs_Data2("DateTime", ($__local_var_7_126)->value0, new Phpurs_Data4("Time", ($__local_var_8_128)->value0, ($__local_var_9_130)->value0, ($__local_var_10_132)->value0, ($__local_var_11_134)->value0)));
} else {
$__t135 = new Phpurs_Data0("Nothing");
};
$__t133 = $__t135;
} else {
$__t133 = new Phpurs_Data0("Nothing");
};
$__t131 = $__t133;
} else {
$__t131 = new Phpurs_Data0("Nothing");
};
$__t129 = $__t131;
} else {
$__t129 = new Phpurs_Data0("Nothing");
};
$__t127 = $__t129;
} else {
$__t127 = new Phpurs_Data0("Nothing");
};
$__t125 = $__t127;
} else {
$__t125 = new Phpurs_Data0("Nothing");
};
$__t123 = $__t125;
} else {
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->tag === "Nothing"))) {
$__t123 = new Phpurs_Data0("Nothing");
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t123 = null;
};
};
$__t5 = $__t123;
} else {
if (((($__local_var_3_0)->value0)->month === 11)) {
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->tag === "Just"))) {
$__local_var_6_137 = ((($GLOBALS['Data_Date_Component_boundedEnumDay'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumDay')))->toEnum)((($__local_var_3_0)->value0)->day);
if ((is_object($__local_var_6_137) && (($__local_var_6_137)->tag === "Just"))) {
$__local_var_7_139 = ((($__local_var_5_3)->value0)(new Phpurs_Data0("November")))(($__local_var_6_137)->value0);
if ((is_object($__local_var_7_139) && (($__local_var_7_139)->tag === "Just"))) {
$__local_var_8_141 = ((($GLOBALS['Data_Time_Component_boundedEnumHour'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumHour')))->toEnum)((($__local_var_3_0)->value0)->hour);
if ((is_object($__local_var_8_141) && (($__local_var_8_141)->tag === "Just"))) {
$__local_var_9_143 = ((($GLOBALS['Data_Time_Component_boundedEnumMinute'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMinute')))->toEnum)((($__local_var_3_0)->value0)->minute);
if ((is_object($__local_var_9_143) && (($__local_var_9_143)->tag === "Just"))) {
$__local_var_10_145 = ((($GLOBALS['Data_Time_Component_boundedEnumSecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumSecond')))->toEnum)((($__local_var_3_0)->value0)->second);
if ((is_object($__local_var_10_145) && (($__local_var_10_145)->tag === "Just"))) {
$__local_var_11_147 = ((($GLOBALS['Data_Time_Component_boundedEnumMillisecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMillisecond')))->toEnum)((($__local_var_3_0)->value0)->millisecond);
if ((is_object($__local_var_11_147) && (($__local_var_11_147)->tag === "Just"))) {
$__t148 = new Phpurs_Data1("Just", new Phpurs_Data2("DateTime", ($__local_var_7_139)->value0, new Phpurs_Data4("Time", ($__local_var_8_141)->value0, ($__local_var_9_143)->value0, ($__local_var_10_145)->value0, ($__local_var_11_147)->value0)));
} else {
$__t148 = new Phpurs_Data0("Nothing");
};
$__t146 = $__t148;
} else {
$__t146 = new Phpurs_Data0("Nothing");
};
$__t144 = $__t146;
} else {
$__t144 = new Phpurs_Data0("Nothing");
};
$__t142 = $__t144;
} else {
$__t142 = new Phpurs_Data0("Nothing");
};
$__t140 = $__t142;
} else {
$__t140 = new Phpurs_Data0("Nothing");
};
$__t138 = $__t140;
} else {
$__t138 = new Phpurs_Data0("Nothing");
};
$__t136 = $__t138;
} else {
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->tag === "Nothing"))) {
$__t136 = new Phpurs_Data0("Nothing");
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t136 = null;
};
};
$__t5 = $__t136;
} else {
if (((($__local_var_3_0)->value0)->month === 12)) {
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->tag === "Just"))) {
$__local_var_6_150 = ((($GLOBALS['Data_Date_Component_boundedEnumDay'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumDay')))->toEnum)((($__local_var_3_0)->value0)->day);
if ((is_object($__local_var_6_150) && (($__local_var_6_150)->tag === "Just"))) {
$__local_var_7_152 = ((($__local_var_5_3)->value0)(new Phpurs_Data0("December")))(($__local_var_6_150)->value0);
if ((is_object($__local_var_7_152) && (($__local_var_7_152)->tag === "Just"))) {
$__local_var_8_154 = ((($GLOBALS['Data_Time_Component_boundedEnumHour'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumHour')))->toEnum)((($__local_var_3_0)->value0)->hour);
if ((is_object($__local_var_8_154) && (($__local_var_8_154)->tag === "Just"))) {
$__local_var_9_156 = ((($GLOBALS['Data_Time_Component_boundedEnumMinute'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMinute')))->toEnum)((($__local_var_3_0)->value0)->minute);
if ((is_object($__local_var_9_156) && (($__local_var_9_156)->tag === "Just"))) {
$__local_var_10_158 = ((($GLOBALS['Data_Time_Component_boundedEnumSecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumSecond')))->toEnum)((($__local_var_3_0)->value0)->second);
if ((is_object($__local_var_10_158) && (($__local_var_10_158)->tag === "Just"))) {
$__local_var_11_160 = ((($GLOBALS['Data_Time_Component_boundedEnumMillisecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMillisecond')))->toEnum)((($__local_var_3_0)->value0)->millisecond);
if ((is_object($__local_var_11_160) && (($__local_var_11_160)->tag === "Just"))) {
$__t161 = new Phpurs_Data1("Just", new Phpurs_Data2("DateTime", ($__local_var_7_152)->value0, new Phpurs_Data4("Time", ($__local_var_8_154)->value0, ($__local_var_9_156)->value0, ($__local_var_10_158)->value0, ($__local_var_11_160)->value0)));
} else {
$__t161 = new Phpurs_Data0("Nothing");
};
$__t159 = $__t161;
} else {
$__t159 = new Phpurs_Data0("Nothing");
};
$__t157 = $__t159;
} else {
$__t157 = new Phpurs_Data0("Nothing");
};
$__t155 = $__t157;
} else {
$__t155 = new Phpurs_Data0("Nothing");
};
$__t153 = $__t155;
} else {
$__t153 = new Phpurs_Data0("Nothing");
};
$__t151 = $__t153;
} else {
$__t151 = new Phpurs_Data0("Nothing");
};
$__t149 = $__t151;
} else {
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->tag === "Nothing"))) {
$__t149 = new Phpurs_Data0("Nothing");
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t149 = null;
};
};
$__t5 = $__t149;
} else {
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->tag === "Just"))) {
$__t5 = new Phpurs_Data0("Nothing");
} else {
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->tag === "Nothing"))) {
$__t5 = new Phpurs_Data0("Nothing");
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t5 = null;
};
};
};
};
};
};
};
};
};
};
};
};
};
};
$__t1 = $__t5;
} else {
if ((is_object($__local_var_3_0) && (($__local_var_3_0)->tag === "Nothing"))) {
$__t1 = new Phpurs_Data0("Nothing");
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };
















