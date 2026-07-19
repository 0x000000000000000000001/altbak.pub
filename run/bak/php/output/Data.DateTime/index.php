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
  $__res = (((("(DateTime " . ((($GLOBALS['Data_Date_showDate'] ?? \PhpursThunks::eval('Data_Date_showDate')))->show)(($v_0)->value0)) . " ") . ((($GLOBALS['Data_Time_showTime'] ?? \PhpursThunks::eval('Data_Time_showTime')))->show)(($v_0)->value1)) . ")");
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
  $__res = ((((($GLOBALS['Data_Date_eqDate'] ?? \PhpursThunks::eval('Data_Date_eqDate')))->eq)(($x_0)->value0))(($y_1)->value0) && (((((($x_0)->value1)->value0 === (($y_1)->value1)->value0) && ((($x_0)->value1)->value1 === (($y_1)->value1)->value1)) && ((($x_0)->value1)->value2 === (($y_1)->value1)->value2)) && ((($x_0)->value1)->value3 === (($y_1)->value1)->value3)));
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
  $__res = (($dictDuration_0)->toDuration)((($GLOBALS['Data_DateTime_calcDiff'] ?? \PhpursThunks::eval('Data_DateTime_calcDiff')))((($GLOBALS['Data_DateTime_toRecord'] ?? \PhpursThunks::eval('Data_DateTime_toRecord')))($dt1_1), (($GLOBALS['Data_DateTime_toRecord'] ?? \PhpursThunks::eval('Data_DateTime_toRecord')))($dt2_2)));
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
\PhpursThunks::$thunks['Data_DateTime_boundedDateTime'] = function() { $v = (object)["bottom" => new Phpurs_Data2("DateTime", new Phpurs_Data3("Date", -271820, new Phpurs_Data0("January"), 1), new Phpurs_Data4("Time", 0, 0, 0, 0)), "top" => new Phpurs_Data2("DateTime", new Phpurs_Data3("Date", 275759, new Phpurs_Data0("December"), 31), new Phpurs_Data4("Time", 23, 59, 59, 999)), "Ord0" => function($dollar__unused_0) {
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
if ((((($__local_var_3_0)->value0)->year >= -271820) && ((($__local_var_3_0)->value0)->year <= 275759))) {
$__t2 = new Phpurs_Data1("Just", (($__local_var_3_0)->value0)->year);
} else {
$__t2 = new Phpurs_Data0("Nothing");
};
$__local_var_4_2 = $__t2;
if ((is_object($__local_var_4_2) && (($__local_var_4_2)->tag === "Just"))) {
$__t4 = new Phpurs_Data1("Just", (($GLOBALS['Data_Date_exactDate'] ?? \PhpursThunks::eval('Data_Date_exactDate')))(($__local_var_4_2)->value0));
} else {
$__t4 = new Phpurs_Data0("Nothing");
};
$__local_var_5_4 = $__t4;
switch ((($__local_var_3_0)->value0)->month) {
case 1:
$__t6 = new Phpurs_Data1("Just", new Phpurs_Data0("January"));
break;
case 2:
$__t6 = new Phpurs_Data1("Just", new Phpurs_Data0("February"));
break;
case 3:
$__t6 = new Phpurs_Data1("Just", new Phpurs_Data0("March"));
break;
case 4:
$__t6 = new Phpurs_Data1("Just", new Phpurs_Data0("April"));
break;
case 5:
$__t6 = new Phpurs_Data1("Just", new Phpurs_Data0("May"));
break;
case 6:
$__t6 = new Phpurs_Data1("Just", new Phpurs_Data0("June"));
break;
case 7:
$__t6 = new Phpurs_Data1("Just", new Phpurs_Data0("July"));
break;
case 8:
$__t6 = new Phpurs_Data1("Just", new Phpurs_Data0("August"));
break;
case 9:
$__t6 = new Phpurs_Data1("Just", new Phpurs_Data0("September"));
break;
case 10:
$__t6 = new Phpurs_Data1("Just", new Phpurs_Data0("October"));
break;
case 11:
$__t6 = new Phpurs_Data1("Just", new Phpurs_Data0("November"));
break;
case 12:
$__t6 = new Phpurs_Data1("Just", new Phpurs_Data0("December"));
break;
default:
$__t6 = new Phpurs_Data0("Nothing");
break;
};
$__local_var_6_6 = $__t6;
if ((is_object($__local_var_5_4) && (($__local_var_5_4)->tag === "Just"))) {
if ((is_object($__local_var_6_6) && (($__local_var_6_6)->tag === "Just"))) {
if ((((($__local_var_3_0)->value0)->day >= 1) && ((($__local_var_3_0)->value0)->day <= 31))) {
$__t10 = new Phpurs_Data1("Just", ((($__local_var_5_4)->value0)(($__local_var_6_6)->value0))((($__local_var_3_0)->value0)->day));
} else {
$__t10 = new Phpurs_Data0("Nothing");
};
$__t9 = $__t10;
} else {
if ((((($__local_var_3_0)->value0)->day >= 1) && ((($__local_var_3_0)->value0)->day <= 31))) {
$__t9 = new Phpurs_Data0("Nothing");
} else {
$__t9 = new Phpurs_Data0("Nothing");
};
};
$__t8 = $__t9;
} else {
if ((is_object($__local_var_5_4) && (($__local_var_5_4)->tag === "Nothing"))) {
if ((((($__local_var_3_0)->value0)->day >= 1) && ((($__local_var_3_0)->value0)->day <= 31))) {
$__t11 = new Phpurs_Data0("Nothing");
} else {
$__t11 = new Phpurs_Data0("Nothing");
};
$__t8 = $__t11;
} else {
if ((((($__local_var_3_0)->value0)->day >= 1) && ((($__local_var_3_0)->value0)->day <= 31))) {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t8 = null;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t8 = null;
};
};
};
$__local_var_7_8 = $__t8;
if ((is_object($__local_var_7_8) && (($__local_var_7_8)->tag === "Just"))) {
$__t13 = ($__local_var_7_8)->value0;
} else {
if ((is_object($__local_var_7_8) && (($__local_var_7_8)->tag === "Nothing"))) {
$__t13 = new Phpurs_Data0("Nothing");
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t13 = null;
};
};
$__local_var_8_13 = $__t13;
if ((is_object($__local_var_8_13) && (($__local_var_8_13)->tag === "Just"))) {
$__t15 = new Phpurs_Data1("Just", (($GLOBALS['Data_DateTime_DateTime'] ?? \PhpursThunks::eval('Data_DateTime_DateTime')))(($__local_var_8_13)->value0));
} else {
$__t15 = new Phpurs_Data0("Nothing");
};
$__local_var_9_15 = $__t15;
if ((((($__local_var_3_0)->value0)->hour >= 0) && ((($__local_var_3_0)->value0)->hour <= 23))) {
if ((((($__local_var_3_0)->value0)->minute >= 0) && ((($__local_var_3_0)->value0)->minute <= 59))) {
if ((((($__local_var_3_0)->value0)->second >= 0) && ((($__local_var_3_0)->value0)->second <= 59))) {
if ((((($__local_var_3_0)->value0)->millisecond >= 0) && ((($__local_var_3_0)->value0)->millisecond <= 999))) {
if ((is_object($__local_var_9_15) && (($__local_var_9_15)->tag === "Just"))) {
$__t21 = new Phpurs_Data1("Just", (($__local_var_9_15)->value0)(new Phpurs_Data4("Time", (($__local_var_3_0)->value0)->hour, (($__local_var_3_0)->value0)->minute, (($__local_var_3_0)->value0)->second, (($__local_var_3_0)->value0)->millisecond)));
} else {
if ((is_object($__local_var_9_15) && (($__local_var_9_15)->tag === "Nothing"))) {
$__t21 = new Phpurs_Data0("Nothing");
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t21 = null;
};
};
$__t20 = $__t21;
} else {
if ((is_object($__local_var_9_15) && (($__local_var_9_15)->tag === "Just"))) {
$__t20 = new Phpurs_Data0("Nothing");
} else {
if ((is_object($__local_var_9_15) && (($__local_var_9_15)->tag === "Nothing"))) {
$__t20 = new Phpurs_Data0("Nothing");
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t20 = null;
};
};
};
$__t19 = $__t20;
} else {
if ((((($__local_var_3_0)->value0)->millisecond >= 0) && ((($__local_var_3_0)->value0)->millisecond <= 999))) {
if ((is_object($__local_var_9_15) && (($__local_var_9_15)->tag === "Just"))) {
$__t22 = new Phpurs_Data0("Nothing");
} else {
if ((is_object($__local_var_9_15) && (($__local_var_9_15)->tag === "Nothing"))) {
$__t22 = new Phpurs_Data0("Nothing");
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t22 = null;
};
};
$__t19 = $__t22;
} else {
if ((is_object($__local_var_9_15) && (($__local_var_9_15)->tag === "Just"))) {
$__t19 = new Phpurs_Data0("Nothing");
} else {
if ((is_object($__local_var_9_15) && (($__local_var_9_15)->tag === "Nothing"))) {
$__t19 = new Phpurs_Data0("Nothing");
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t19 = null;
};
};
};
};
$__t18 = $__t19;
} else {
if ((((($__local_var_3_0)->value0)->second >= 0) && ((($__local_var_3_0)->value0)->second <= 59))) {
if ((((($__local_var_3_0)->value0)->millisecond >= 0) && ((($__local_var_3_0)->value0)->millisecond <= 999))) {
if ((is_object($__local_var_9_15) && (($__local_var_9_15)->tag === "Just"))) {
$__t24 = new Phpurs_Data0("Nothing");
} else {
if ((is_object($__local_var_9_15) && (($__local_var_9_15)->tag === "Nothing"))) {
$__t24 = new Phpurs_Data0("Nothing");
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t24 = null;
};
};
$__t23 = $__t24;
} else {
if ((is_object($__local_var_9_15) && (($__local_var_9_15)->tag === "Just"))) {
$__t23 = new Phpurs_Data0("Nothing");
} else {
if ((is_object($__local_var_9_15) && (($__local_var_9_15)->tag === "Nothing"))) {
$__t23 = new Phpurs_Data0("Nothing");
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t23 = null;
};
};
};
$__t18 = $__t23;
} else {
if ((((($__local_var_3_0)->value0)->millisecond >= 0) && ((($__local_var_3_0)->value0)->millisecond <= 999))) {
if ((is_object($__local_var_9_15) && (($__local_var_9_15)->tag === "Just"))) {
$__t25 = new Phpurs_Data0("Nothing");
} else {
if ((is_object($__local_var_9_15) && (($__local_var_9_15)->tag === "Nothing"))) {
$__t25 = new Phpurs_Data0("Nothing");
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t25 = null;
};
};
$__t18 = $__t25;
} else {
if ((is_object($__local_var_9_15) && (($__local_var_9_15)->tag === "Just"))) {
$__t18 = new Phpurs_Data0("Nothing");
} else {
if ((is_object($__local_var_9_15) && (($__local_var_9_15)->tag === "Nothing"))) {
$__t18 = new Phpurs_Data0("Nothing");
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t18 = null;
};
};
};
};
};
$__t17 = $__t18;
} else {
if ((((($__local_var_3_0)->value0)->millisecond >= 0) && ((($__local_var_3_0)->value0)->millisecond <= 999))) {
if ((is_object($__local_var_9_15) && (($__local_var_9_15)->tag === "Just"))) {
$__t26 = new Phpurs_Data0("Nothing");
} else {
if ((is_object($__local_var_9_15) && (($__local_var_9_15)->tag === "Nothing"))) {
$__t26 = new Phpurs_Data0("Nothing");
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t26 = null;
};
};
$__t17 = $__t26;
} else {
if ((is_object($__local_var_9_15) && (($__local_var_9_15)->tag === "Just"))) {
$__t17 = new Phpurs_Data0("Nothing");
} else {
if ((is_object($__local_var_9_15) && (($__local_var_9_15)->tag === "Nothing"))) {
$__t17 = new Phpurs_Data0("Nothing");
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t17 = null;
};
};
};
};
$__t1 = $__t17;
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
















