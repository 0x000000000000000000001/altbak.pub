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
\PhpursThunks::$thunks['Data_Date_fromJust'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Date_fromJust"), recVars=[];
  if ((is_object($v_0) && (($v_0)->tag === "Just"))) {
$__t0 = ($v_0)->value0;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Date_greaterThan'] = function() { $v = (function() use (&$__fn) {
$__local_var_0_0 = (($GLOBALS['Data_Maybe_ordMaybe'] ?? \PhpursThunks::eval('Data_Maybe_ordMaybe')))(($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')));
return (function() use ($__local_var_0_0) {
  $__fn = function($a1_1, $a2_2 = null) use ($__local_var_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (is_object(((($__local_var_0_0)->compare)($a1_1))($a2_2)) && ((((($__local_var_0_0)->compare)($a1_1))($a2_2))->tag === "GT"));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
})(); return $v; };
\PhpursThunks::$thunks['Data_Date_Date'] = function() { $v = (function() {
  $__fn = function($value0, $value1 = null, $value2 = null) use (&$__fn) {
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
})(); return $v; };
\PhpursThunks::$thunks['Data_Date_year'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Date_year"), recVars=[];
  $__res = ($v_0)->value0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Date_weekday'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Date_weekday"), recVars=[];
  if ((is_object(($v_0)->value1) && ((($v_0)->value1)->tag === "January"))) {
$__t0 = 1;
} else {
if ((is_object(($v_0)->value1) && ((($v_0)->value1)->tag === "February"))) {
$__t0 = 2;
} else {
if ((is_object(($v_0)->value1) && ((($v_0)->value1)->tag === "March"))) {
$__t0 = 3;
} else {
if ((is_object(($v_0)->value1) && ((($v_0)->value1)->tag === "April"))) {
$__t0 = 4;
} else {
if ((is_object(($v_0)->value1) && ((($v_0)->value1)->tag === "May"))) {
$__t0 = 5;
} else {
if ((is_object(($v_0)->value1) && ((($v_0)->value1)->tag === "June"))) {
$__t0 = 6;
} else {
if ((is_object(($v_0)->value1) && ((($v_0)->value1)->tag === "July"))) {
$__t0 = 7;
} else {
if ((is_object(($v_0)->value1) && ((($v_0)->value1)->tag === "August"))) {
$__t0 = 8;
} else {
if ((is_object(($v_0)->value1) && ((($v_0)->value1)->tag === "September"))) {
$__t0 = 9;
} else {
if ((is_object(($v_0)->value1) && ((($v_0)->value1)->tag === "October"))) {
$__t0 = 10;
} else {
if ((is_object(($v_0)->value1) && ((($v_0)->value1)->tag === "November"))) {
$__t0 = 11;
} else {
if ((is_object(($v_0)->value1) && ((($v_0)->value1)->tag === "December"))) {
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
  $n_1_0 = (($GLOBALS['Data_Date_calcWeekday'] ?? \PhpursThunks::eval('Data_Date_calcWeekday')))(($v_0)->value0, $__t0, ($v_0)->value2);
  switch ($n_1_0) {
case 0:
$__t2 = new Phpurs_Data0("Sunday");
break;
case 1:
$__t2 = new Phpurs_Data0("Monday");
break;
case 2:
$__t2 = new Phpurs_Data0("Tuesday");
break;
case 3:
$__t2 = new Phpurs_Data0("Wednesday");
break;
case 4:
$__t2 = new Phpurs_Data0("Thursday");
break;
case 5:
$__t2 = new Phpurs_Data0("Friday");
break;
case 6:
$__t2 = new Phpurs_Data0("Saturday");
break;
case 7:
$__t2 = new Phpurs_Data0("Sunday");
break;
default:
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
break;
};
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Date_showDate'] = function() { $v = (object)["show" => function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object(($v_0)->value1) && ((($v_0)->value1)->tag === "January"))) {
$__t0 = "January";
} else {
if ((is_object(($v_0)->value1) && ((($v_0)->value1)->tag === "February"))) {
$__t0 = "February";
} else {
if ((is_object(($v_0)->value1) && ((($v_0)->value1)->tag === "March"))) {
$__t0 = "March";
} else {
if ((is_object(($v_0)->value1) && ((($v_0)->value1)->tag === "April"))) {
$__t0 = "April";
} else {
if ((is_object(($v_0)->value1) && ((($v_0)->value1)->tag === "May"))) {
$__t0 = "May";
} else {
if ((is_object(($v_0)->value1) && ((($v_0)->value1)->tag === "June"))) {
$__t0 = "June";
} else {
if ((is_object(($v_0)->value1) && ((($v_0)->value1)->tag === "July"))) {
$__t0 = "July";
} else {
if ((is_object(($v_0)->value1) && ((($v_0)->value1)->tag === "August"))) {
$__t0 = "August";
} else {
if ((is_object(($v_0)->value1) && ((($v_0)->value1)->tag === "September"))) {
$__t0 = "September";
} else {
if ((is_object(($v_0)->value1) && ((($v_0)->value1)->tag === "October"))) {
$__t0 = "October";
} else {
if ((is_object(($v_0)->value1) && ((($v_0)->value1)->tag === "November"))) {
$__t0 = "November";
} else {
if ((is_object(($v_0)->value1) && ((($v_0)->value1)->tag === "December"))) {
$__t0 = "December";
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
  $__res = (((((("(Date (Year " . (($GLOBALS['Data_Show_showIntImpl'] ?? \PhpursThunks::eval('Data_Show_showIntImpl')))(($v_0)->value0)) . ") ") . $__t0) . " (Day ") . (($GLOBALS['Data_Show_showIntImpl'] ?? \PhpursThunks::eval('Data_Show_showIntImpl')))(($v_0)->value2)) . "))");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Date_month'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Date_month"), recVars=[];
  $__res = ($v_0)->value1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Date_isLeapYear'] = function() { $v = function($y_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Date_isLeapYear"), recVars=[];
  $__res = ((((($GLOBALS['Data_EuclideanRing_intMod'] ?? \PhpursThunks::eval('Data_EuclideanRing_intMod')))($y_0))(4) === 0) && ((((($GLOBALS['Data_EuclideanRing_intMod'] ?? \PhpursThunks::eval('Data_EuclideanRing_intMod')))($y_0))(400) === 0) || (((($GLOBALS['Data_EuclideanRing_intMod'] ?? \PhpursThunks::eval('Data_EuclideanRing_intMod')))($y_0))(100) !== 0)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Date_lastDayOfMonth'] = function() { $v = (function() {
  $__fn = function($y_0, $m_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Date_lastDayOfMonth"), recVars=[];
  $unsafeDay_2_0 = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Date_fromJust'] ?? \PhpursThunks::eval('Data_Date_fromJust'))))((($GLOBALS['Data_Date_Component_boundedEnumDay'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumDay')))->toEnum);
  if ((is_object($m_1) && (($m_1)->tag === "January"))) {
$__t1 = ($unsafeDay_2_0)(31);
} else {
if ((is_object($m_1) && (($m_1)->tag === "February"))) {
if ((($GLOBALS['Data_Date_isLeapYear'] ?? \PhpursThunks::eval('Data_Date_isLeapYear')))($y_0)) {
$__t2 = ($unsafeDay_2_0)(29);
} else {
$__t2 = ($unsafeDay_2_0)(28);
};
$__t1 = $__t2;
} else {
if ((is_object($m_1) && (($m_1)->tag === "March"))) {
$__t1 = ($unsafeDay_2_0)(31);
} else {
if ((is_object($m_1) && (($m_1)->tag === "April"))) {
$__t1 = ($unsafeDay_2_0)(30);
} else {
if ((is_object($m_1) && (($m_1)->tag === "May"))) {
$__t1 = ($unsafeDay_2_0)(31);
} else {
if ((is_object($m_1) && (($m_1)->tag === "June"))) {
$__t1 = ($unsafeDay_2_0)(30);
} else {
if ((is_object($m_1) && (($m_1)->tag === "July"))) {
$__t1 = ($unsafeDay_2_0)(31);
} else {
if ((is_object($m_1) && (($m_1)->tag === "August"))) {
$__t1 = ($unsafeDay_2_0)(31);
} else {
if ((is_object($m_1) && (($m_1)->tag === "September"))) {
$__t1 = ($unsafeDay_2_0)(30);
} else {
if ((is_object($m_1) && (($m_1)->tag === "October"))) {
$__t1 = ($unsafeDay_2_0)(31);
} else {
if ((is_object($m_1) && (($m_1)->tag === "November"))) {
$__t1 = ($unsafeDay_2_0)(30);
} else {
if ((is_object($m_1) && (($m_1)->tag === "December"))) {
$__t1 = ($unsafeDay_2_0)(31);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
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
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Date_eqDate'] = function() { $v = (object)["eq" => (function() {
  $__fn = function($x_0, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object(($x_0)->value1) && ((($x_0)->value1)->tag === "January"))) {
$__t0 = (is_object(($y_1)->value1) && ((($y_1)->value1)->tag === "January"));
} else {
if ((is_object(($x_0)->value1) && ((($x_0)->value1)->tag === "February"))) {
$__t0 = (is_object(($y_1)->value1) && ((($y_1)->value1)->tag === "February"));
} else {
if ((is_object(($x_0)->value1) && ((($x_0)->value1)->tag === "March"))) {
$__t0 = (is_object(($y_1)->value1) && ((($y_1)->value1)->tag === "March"));
} else {
if ((is_object(($x_0)->value1) && ((($x_0)->value1)->tag === "April"))) {
$__t0 = (is_object(($y_1)->value1) && ((($y_1)->value1)->tag === "April"));
} else {
if ((is_object(($x_0)->value1) && ((($x_0)->value1)->tag === "May"))) {
$__t0 = (is_object(($y_1)->value1) && ((($y_1)->value1)->tag === "May"));
} else {
if ((is_object(($x_0)->value1) && ((($x_0)->value1)->tag === "June"))) {
$__t0 = (is_object(($y_1)->value1) && ((($y_1)->value1)->tag === "June"));
} else {
if ((is_object(($x_0)->value1) && ((($x_0)->value1)->tag === "July"))) {
$__t0 = (is_object(($y_1)->value1) && ((($y_1)->value1)->tag === "July"));
} else {
if ((is_object(($x_0)->value1) && ((($x_0)->value1)->tag === "August"))) {
$__t0 = (is_object(($y_1)->value1) && ((($y_1)->value1)->tag === "August"));
} else {
if ((is_object(($x_0)->value1) && ((($x_0)->value1)->tag === "September"))) {
$__t0 = (is_object(($y_1)->value1) && ((($y_1)->value1)->tag === "September"));
} else {
if ((is_object(($x_0)->value1) && ((($x_0)->value1)->tag === "October"))) {
$__t0 = (is_object(($y_1)->value1) && ((($y_1)->value1)->tag === "October"));
} else {
if ((is_object(($x_0)->value1) && ((($x_0)->value1)->tag === "November"))) {
$__t0 = (is_object(($y_1)->value1) && ((($y_1)->value1)->tag === "November"));
} else {
$__t0 = ((is_object(($x_0)->value1) && ((($x_0)->value1)->tag === "December")) && (is_object(($y_1)->value1) && ((($y_1)->value1)->tag === "December")));
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
  $__res = (((($x_0)->value0 === ($y_1)->value0) && $__t0) && (($x_0)->value2 === ($y_1)->value2));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Data_Date_ordDate'] = function() { $v = (object)["compare" => (function() {
  $__fn = function($x_0, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $v_2_0 = (((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)(($x_0)->value0))(($y_1)->value0);
  if ((is_object($v_2_0) && (($v_2_0)->tag === "LT"))) {
$__t3 = new Phpurs_Data0("LT");
} else {
if ((is_object($v_2_0) && (($v_2_0)->tag === "GT"))) {
$__t3 = new Phpurs_Data0("GT");
} else {
$v1_3_1 = (((($GLOBALS['Data_Date_Component_ordMonth'] ?? \PhpursThunks::eval('Data_Date_Component_ordMonth')))->compare)(($x_0)->value1))(($y_1)->value1);
if ((is_object($v1_3_1) && (($v1_3_1)->tag === "LT"))) {
$__t2 = new Phpurs_Data0("LT");
} else {
if ((is_object($v1_3_1) && (($v1_3_1)->tag === "GT"))) {
$__t2 = new Phpurs_Data0("GT");
} else {
$__t2 = (((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)(($x_0)->value2))(($y_1)->value2);
};
};
$__t3 = $__t2;
};
};
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_Date_eqDate'] ?? \PhpursThunks::eval('Data_Date_eqDate'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Date_enumDate'] = function() { $v = (object)["succ" => function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $sm_1_0 = ((($GLOBALS['Data_Date_Component_enumMonth'] ?? \PhpursThunks::eval('Data_Date_Component_enumMonth')))->succ)(($v_0)->value1);
  $v1_2_1 = ((($GLOBALS['Data_Date_Component_enumDay'] ?? \PhpursThunks::eval('Data_Date_Component_enumDay')))->succ)(($v_0)->value2);
  if (((($GLOBALS['Data_Date_greaterThan'] ?? \PhpursThunks::eval('Data_Date_greaterThan')))($v1_2_1))(new Phpurs_Data1("Just", ((($GLOBALS['Data_Date_lastDayOfMonth'] ?? \PhpursThunks::eval('Data_Date_lastDayOfMonth')))(($v_0)->value0))(($v_0)->value1)))) {
if ((function() use ($sm_1_0, &$__fn) {
if ((is_object($sm_1_0) && (($sm_1_0)->tag === "Nothing"))) {
$__t15 = true;
} else {
if ((is_object($sm_1_0) && (($sm_1_0)->tag === "Just"))) {
$__t15 = false;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t15 = null;
};
};
return $__t15;
})()) {
$__t14 = ((($GLOBALS['Data_Date_Component_enumYear'] ?? \PhpursThunks::eval('Data_Date_Component_enumYear')))->succ)(($v_0)->value0);
} else {
$__t14 = new Phpurs_Data1("Just", ($v_0)->value0);
};
$__local_var_3_14 = $__t14;
if ((is_object($__local_var_3_14) && (($__local_var_3_14)->tag === "Just"))) {
if ((is_object($sm_1_0) && (($sm_1_0)->tag === "Nothing"))) {
$__t18 = new Phpurs_Data0("January");
} else {
if ((is_object($sm_1_0) && (($sm_1_0)->tag === "Just"))) {
$__t18 = ($sm_1_0)->value0;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t18 = null;
};
};
$__t17 = new Phpurs_Data1("Just", new Phpurs_Data3("Date", ($__local_var_3_14)->value0, $__t18, 1));
} else {
$__t17 = new Phpurs_Data0("Nothing");
};
$__t13 = $__t17;
} else {
if ((function() use ($sm_1_0, $v1_2_1, &$__fn) {
if ((is_object($v1_2_1) && (($v1_2_1)->tag === "Nothing"))) {
$__t3 = true;
} else {
if ((is_object($v1_2_1) && (($v1_2_1)->tag === "Just"))) {
$__t3 = false;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
};
};
if ((is_object($sm_1_0) && (($sm_1_0)->tag === "Nothing"))) {
$__t4 = true;
} else {
if ((is_object($sm_1_0) && (($sm_1_0)->tag === "Just"))) {
$__t4 = false;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t4 = null;
};
};
return ($__t3 && $__t4);
})()) {
$__t2 = ((($GLOBALS['Data_Date_Component_enumYear'] ?? \PhpursThunks::eval('Data_Date_Component_enumYear')))->succ)(($v_0)->value0);
} else {
$__t2 = new Phpurs_Data1("Just", ($v_0)->value0);
};
$__local_var_3_2 = $__t2;
if ((is_object($__local_var_3_2) && (($__local_var_3_2)->tag === "Just"))) {
if ((function() use ($v1_2_1, &$__fn) {
if ((is_object($v1_2_1) && (($v1_2_1)->tag === "Nothing"))) {
$__t8 = true;
} else {
if ((is_object($v1_2_1) && (($v1_2_1)->tag === "Just"))) {
$__t8 = false;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t8 = null;
};
};
return $__t8;
})()) {
if ((is_object($sm_1_0) && (($sm_1_0)->tag === "Nothing"))) {
$__t9 = new Phpurs_Data0("January");
} else {
if ((is_object($sm_1_0) && (($sm_1_0)->tag === "Just"))) {
$__t9 = ($sm_1_0)->value0;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t9 = null;
};
};
$__t7 = $__t9;
} else {
$__t7 = ($v_0)->value1;
};
$__local_var_4_7 = ((($GLOBALS['Data_Date_Date'] ?? \PhpursThunks::eval('Data_Date_Date')))(($__local_var_3_2)->value0))($__t7);
if ((function() use ($v1_2_1, &$__fn) {
if ((is_object($v1_2_1) && (($v1_2_1)->tag === "Nothing"))) {
$__t12 = true;
} else {
if ((is_object($v1_2_1) && (($v1_2_1)->tag === "Just"))) {
$__t12 = false;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t12 = null;
};
};
return $__t12;
})()) {
$__t11 = new Phpurs_Data1("Just", ($__local_var_4_7)(1));
} else {
if ((is_object($v1_2_1) && (($v1_2_1)->tag === "Just"))) {
$__t11 = new Phpurs_Data1("Just", ($__local_var_4_7)(($v1_2_1)->value0));
} else {
$__t11 = new Phpurs_Data0("Nothing");
};
};
$__t6 = $__t11;
} else {
$__t6 = new Phpurs_Data0("Nothing");
};
$__t13 = $__t6;
};
  $__res = $__t13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "pred" => function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $pm_1_19 = ((($GLOBALS['Data_Date_Component_enumMonth'] ?? \PhpursThunks::eval('Data_Date_Component_enumMonth')))->pred)(($v_0)->value1);
  $pd_2_20 = ((($GLOBALS['Data_Date_Component_enumDay'] ?? \PhpursThunks::eval('Data_Date_Component_enumDay')))->pred)(($v_0)->value2);
  if ((function() use ($pd_2_20, &$__fn) {
if ((is_object($pd_2_20) && (($pd_2_20)->tag === "Nothing"))) {
$__t22 = true;
} else {
if ((is_object($pd_2_20) && (($pd_2_20)->tag === "Just"))) {
$__t22 = false;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t22 = null;
};
};
return $__t22;
})()) {
if ((is_object($pm_1_19) && (($pm_1_19)->tag === "Nothing"))) {
$__t23 = new Phpurs_Data0("December");
} else {
if ((is_object($pm_1_19) && (($pm_1_19)->tag === "Just"))) {
$__t23 = ($pm_1_19)->value0;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t23 = null;
};
};
$__t21 = $__t23;
} else {
$__t21 = ($v_0)->value1;
};
  $m__prime___3_21 = $__t21;
  $l_4_25 = ((($GLOBALS['Data_Date_lastDayOfMonth'] ?? \PhpursThunks::eval('Data_Date_lastDayOfMonth')))(($v_0)->value0))($m__prime___3_21);
  if ((function() use ($pd_2_20, $pm_1_19, &$__fn) {
if ((is_object($pd_2_20) && (($pd_2_20)->tag === "Nothing"))) {
$__t27 = true;
} else {
if ((is_object($pd_2_20) && (($pd_2_20)->tag === "Just"))) {
$__t27 = false;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t27 = null;
};
};
if ((is_object($pm_1_19) && (($pm_1_19)->tag === "Nothing"))) {
$__t28 = true;
} else {
if ((is_object($pm_1_19) && (($pm_1_19)->tag === "Just"))) {
$__t28 = false;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t28 = null;
};
};
return ($__t27 && $__t28);
})()) {
$__t26 = ((($GLOBALS['Data_Date_Component_enumYear'] ?? \PhpursThunks::eval('Data_Date_Component_enumYear')))->pred)(($v_0)->value0);
} else {
$__t26 = new Phpurs_Data1("Just", ($v_0)->value0);
};
  $__local_var_5_26 = $__t26;
  if ((is_object($__local_var_5_26) && (($__local_var_5_26)->tag === "Just"))) {
$__local_var_6_31 = ((($GLOBALS['Data_Date_Date'] ?? \PhpursThunks::eval('Data_Date_Date')))(($__local_var_5_26)->value0))($m__prime___3_21);
if ((function() use ($pd_2_20, &$__fn) {
if ((is_object($pd_2_20) && (($pd_2_20)->tag === "Nothing"))) {
$__t33 = true;
} else {
if ((is_object($pd_2_20) && (($pd_2_20)->tag === "Just"))) {
$__t33 = false;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t33 = null;
};
};
return $__t33;
})()) {
$__t32 = new Phpurs_Data1("Just", ($__local_var_6_31)($l_4_25));
} else {
if ((is_object($pd_2_20) && (($pd_2_20)->tag === "Just"))) {
$__t32 = new Phpurs_Data1("Just", ($__local_var_6_31)(($pd_2_20)->value0));
} else {
$__t32 = new Phpurs_Data0("Nothing");
};
};
$__t30 = $__t32;
} else {
$__t30 = new Phpurs_Data0("Nothing");
};
  $__res = $__t30;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Ord0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_Date_ordDate'] ?? \PhpursThunks::eval('Data_Date_ordDate'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Date_diff'] = function() { $v = (function() {
  $__fn = function($dictDuration_0, $v_1 = null, $v1_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Date_diff"), recVars=[];
  if ((is_object(($v_1)->value1) && ((($v_1)->value1)->tag === "January"))) {
$__t0 = 1;
} else {
if ((is_object(($v_1)->value1) && ((($v_1)->value1)->tag === "February"))) {
$__t0 = 2;
} else {
if ((is_object(($v_1)->value1) && ((($v_1)->value1)->tag === "March"))) {
$__t0 = 3;
} else {
if ((is_object(($v_1)->value1) && ((($v_1)->value1)->tag === "April"))) {
$__t0 = 4;
} else {
if ((is_object(($v_1)->value1) && ((($v_1)->value1)->tag === "May"))) {
$__t0 = 5;
} else {
if ((is_object(($v_1)->value1) && ((($v_1)->value1)->tag === "June"))) {
$__t0 = 6;
} else {
if ((is_object(($v_1)->value1) && ((($v_1)->value1)->tag === "July"))) {
$__t0 = 7;
} else {
if ((is_object(($v_1)->value1) && ((($v_1)->value1)->tag === "August"))) {
$__t0 = 8;
} else {
if ((is_object(($v_1)->value1) && ((($v_1)->value1)->tag === "September"))) {
$__t0 = 9;
} else {
if ((is_object(($v_1)->value1) && ((($v_1)->value1)->tag === "October"))) {
$__t0 = 10;
} else {
if ((is_object(($v_1)->value1) && ((($v_1)->value1)->tag === "November"))) {
$__t0 = 11;
} else {
if ((is_object(($v_1)->value1) && ((($v_1)->value1)->tag === "December"))) {
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
  if ((is_object(($v1_2)->value1) && ((($v1_2)->value1)->tag === "January"))) {
$__t1 = 1;
} else {
if ((is_object(($v1_2)->value1) && ((($v1_2)->value1)->tag === "February"))) {
$__t1 = 2;
} else {
if ((is_object(($v1_2)->value1) && ((($v1_2)->value1)->tag === "March"))) {
$__t1 = 3;
} else {
if ((is_object(($v1_2)->value1) && ((($v1_2)->value1)->tag === "April"))) {
$__t1 = 4;
} else {
if ((is_object(($v1_2)->value1) && ((($v1_2)->value1)->tag === "May"))) {
$__t1 = 5;
} else {
if ((is_object(($v1_2)->value1) && ((($v1_2)->value1)->tag === "June"))) {
$__t1 = 6;
} else {
if ((is_object(($v1_2)->value1) && ((($v1_2)->value1)->tag === "July"))) {
$__t1 = 7;
} else {
if ((is_object(($v1_2)->value1) && ((($v1_2)->value1)->tag === "August"))) {
$__t1 = 8;
} else {
if ((is_object(($v1_2)->value1) && ((($v1_2)->value1)->tag === "September"))) {
$__t1 = 9;
} else {
if ((is_object(($v1_2)->value1) && ((($v1_2)->value1)->tag === "October"))) {
$__t1 = 10;
} else {
if ((is_object(($v1_2)->value1) && ((($v1_2)->value1)->tag === "November"))) {
$__t1 = 11;
} else {
if ((is_object(($v1_2)->value1) && ((($v1_2)->value1)->tag === "December"))) {
$__t1 = 12;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
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
  $__res = (($dictDuration_0)->toDuration)((($GLOBALS['Data_Date_calcDiff'] ?? \PhpursThunks::eval('Data_Date_calcDiff')))(($v_1)->value0, $__t0, ($v_1)->value2, ($v1_2)->value0, $__t1, ($v1_2)->value2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Date_day'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Date_day"), recVars=[];
  $__res = ($v_0)->value2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Date_canonicalDate'] = function() { $v = (function() {
  $__fn = function($y_0, $m_1 = null, $d_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Date_canonicalDate"), recVars=[];
  if ((is_object($m_1) && (($m_1)->tag === "January"))) {
$__t1 = 1;
} else {
if ((is_object($m_1) && (($m_1)->tag === "February"))) {
$__t1 = 2;
} else {
if ((is_object($m_1) && (($m_1)->tag === "March"))) {
$__t1 = 3;
} else {
if ((is_object($m_1) && (($m_1)->tag === "April"))) {
$__t1 = 4;
} else {
if ((is_object($m_1) && (($m_1)->tag === "May"))) {
$__t1 = 5;
} else {
if ((is_object($m_1) && (($m_1)->tag === "June"))) {
$__t1 = 6;
} else {
if ((is_object($m_1) && (($m_1)->tag === "July"))) {
$__t1 = 7;
} else {
if ((is_object($m_1) && (($m_1)->tag === "August"))) {
$__t1 = 8;
} else {
if ((is_object($m_1) && (($m_1)->tag === "September"))) {
$__t1 = 9;
} else {
if ((is_object($m_1) && (($m_1)->tag === "October"))) {
$__t1 = 10;
} else {
if ((is_object($m_1) && (($m_1)->tag === "November"))) {
$__t1 = 11;
} else {
if ((is_object($m_1) && (($m_1)->tag === "December"))) {
$__t1 = 12;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
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
  $__res = (($GLOBALS['Data_Date_canonicalDateImpl'] ?? \PhpursThunks::eval('Data_Date_canonicalDateImpl')))((function() {
  $__fn = function($y__prime___3, $m__prime___4 = null, $d__prime___5 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  switch ($m__prime___4) {
case 1:
$__t0 = new Phpurs_Data0("January");
break;
case 2:
$__t0 = new Phpurs_Data0("February");
break;
case 3:
$__t0 = new Phpurs_Data0("March");
break;
case 4:
$__t0 = new Phpurs_Data0("April");
break;
case 5:
$__t0 = new Phpurs_Data0("May");
break;
case 6:
$__t0 = new Phpurs_Data0("June");
break;
case 7:
$__t0 = new Phpurs_Data0("July");
break;
case 8:
$__t0 = new Phpurs_Data0("August");
break;
case 9:
$__t0 = new Phpurs_Data0("September");
break;
case 10:
$__t0 = new Phpurs_Data0("October");
break;
case 11:
$__t0 = new Phpurs_Data0("November");
break;
case 12:
$__t0 = new Phpurs_Data0("December");
break;
default:
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
break;
};
  $__res = new Phpurs_Data3("Date", $y__prime___3, $__t0, $d__prime___5);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), $y_0, $__t1, $d_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Date_exactDate'] = function() { $v = (function() {
  $__fn = function($y_0, $m_1 = null, $d_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Date_exactDate"), recVars=[];
  if ((((($GLOBALS['Data_Date_eqDate'] ?? \PhpursThunks::eval('Data_Date_eqDate')))->eq)((((($GLOBALS['Data_Date_canonicalDate'] ?? \PhpursThunks::eval('Data_Date_canonicalDate')))($y_0))($m_1))($d_2)))(new Phpurs_Data3("Date", $y_0, $m_1, $d_2))) {
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data3("Date", $y_0, $m_1, $d_2));
} else {
$__t0 = new Phpurs_Data0("Nothing");
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Date_boundedDate'] = function() { $v = (object)["bottom" => new Phpurs_Data3("Date", -271820, new Phpurs_Data0("January"), 1), "top" => new Phpurs_Data3("Date", 275759, new Phpurs_Data0("December"), 31), "Ord0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_Date_ordDate'] ?? \PhpursThunks::eval('Data_Date_ordDate'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Date_adjust'] = function() { $v = (function() {
  $__fn = function($v_0, $date_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Date_adjust"), recVars=[];
  $adj_2_0 = null;
  $adj_2_0 = (function() use (&$adj_2_0) {
  $__fn = function($v1_3, $v2_4 = null) use (&$adj_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "adj_2_0"), recVars=["adj_2_0"];
  while (true) {
switch ($v1_3) {
case 0:
$__t15 = new Phpurs_Data1("Just", $v2_4);
break;
default:
$j_5_1 = ($v1_3 + ($v2_4)->value2);
$low_6_2 = ($j_5_1 < 1);
if ($low_6_2) {
$__local_var_7_4 = ((($GLOBALS['Data_Date_Component_enumMonth'] ?? \PhpursThunks::eval('Data_Date_Component_enumMonth')))->pred)(($v2_4)->value1);
if ((is_object($__local_var_7_4) && (($__local_var_7_4)->tag === "Nothing"))) {
$__t5 = new Phpurs_Data0("December");
} else {
if ((is_object($__local_var_7_4) && (($__local_var_7_4)->tag === "Just"))) {
$__t5 = ($__local_var_7_4)->value0;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t5 = null;
};
};
$__t3 = $__t5;
} else {
$__t3 = ($v2_4)->value1;
};
$l_7_3 = ((($GLOBALS['Data_Date_lastDayOfMonth'] ?? \PhpursThunks::eval('Data_Date_lastDayOfMonth')))(($v2_4)->value0))($__t3);
$hi_8_7 = ($j_5_1 > $l_7_3);
if ($low_6_2) {
$__t8 = $j_5_1;
} else {
if ($hi_8_7) {
$__t8 = (($j_5_1 - $l_7_3) - 1);
} else {
$__t8 = 0;
};
};
$__local_var_9_8 = ($adj_2_0)($__t8);
if ($low_6_2) {
$__t12 = ((($GLOBALS['Data_Date_enumDate'] ?? \PhpursThunks::eval('Data_Date_enumDate')))->pred)(new Phpurs_Data3("Date", ($v2_4)->value0, ($v2_4)->value1, 1));
} else {
if ($hi_8_7) {
$__t12 = ((($GLOBALS['Data_Date_enumDate'] ?? \PhpursThunks::eval('Data_Date_enumDate')))->succ)(new Phpurs_Data3("Date", ($v2_4)->value0, ($v2_4)->value1, $l_7_3));
} else {
$__local_var_10_10 = ((($GLOBALS['Data_Date_Date'] ?? \PhpursThunks::eval('Data_Date_Date')))(($v2_4)->value0))(($v2_4)->value1);
if ((($j_5_1 >= 1) && ($j_5_1 <= 31))) {
$__t11 = new Phpurs_Data1("Just", ($__local_var_10_10)($j_5_1));
} else {
$__t11 = new Phpurs_Data0("Nothing");
};
$__t12 = $__t11;
};
};
$__local_var_10_10 = $__t12;
if ((is_object($__local_var_10_10) && (($__local_var_10_10)->tag === "Just"))) {
$__t14 = ($__local_var_9_8)(($__local_var_10_10)->value0);
} else {
if ((is_object($__local_var_10_10) && (($__local_var_10_10)->tag === "Nothing"))) {
$__t14 = new Phpurs_Data0("Nothing");
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t14 = null;
};
};
$__t15 = $__t14;
break;
};
$__res = $__t15;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__local_var_3_16 = (($GLOBALS['Data_Int_fromNumber'] ?? \PhpursThunks::eval('Data_Int_fromNumber')))($v_0);
  if ((is_object($__local_var_3_16) && (($__local_var_3_16)->tag === "Just"))) {
$__t17 = (($adj_2_0)(($__local_var_3_16)->value0))($date_1);
} else {
if ((is_object($__local_var_3_16) && (($__local_var_3_16)->tag === "Nothing"))) {
$__t17 = new Phpurs_Data0("Nothing");
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t17 = null;
};
};
  $__res = $__t17;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };




















