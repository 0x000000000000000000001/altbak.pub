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
\PhpursThunks::$thunks['Data_DateTime_DateTime'] = function() { $v = (function() {
  $__fn = function($value0 = null, $value1 = null) use (&$__fn) {
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
\PhpursThunks::$thunks['Data_DateTime_toRecord'] = function() { $v = function($v_0 = null) {
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
  $__res = (object)["year" => (($v_0)->{'value0'})->{'value0'}, "month" => $__t0, "day" => (($v_0)->{'value0'})->{'value2'}, "hour" => (($v_0)->{'value1'})->{'value0'}, "minute" => (($v_0)->{'value1'})->{'value1'}, "second" => (($v_0)->{'value1'})->{'value2'}, "millisecond" => (($v_0)->{'value1'})->{'value3'}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_DateTime_time'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ($v_0)->{'value1'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_DateTime_showDateTime'] = function() { $v = (object)["show" => function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))("(DateTime "))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))(((($GLOBALS['Data_Date_showDate'] ?? \PhpursThunks::eval('Data_Date_showDate')))->{'show'})(($v_0)->{'value0'})))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))(" "))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))(((($GLOBALS['Data_Time_showTime'] ?? \PhpursThunks::eval('Data_Time_showTime')))->{'show'})(($v_0)->{'value1'})))(")"))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_DateTime_modifyTimeF'] = function() { $v = (function() {
  $__fn = function($dictFunctor_0 = null, $f_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictFunctor_0)->{'map'})((($GLOBALS['Data_DateTime_DateTime'] ?? \PhpursThunks::eval('Data_DateTime_DateTime')))(($v_2)->{'value0'})))(($f_1)(($v_2)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_DateTime_modifyTime'] = function() { $v = (function() {
  $__fn = function($f_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("DateTime", ($v_1)->{'value0'}, ($f_0)(($v_1)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_DateTime_modifyDateF'] = function() { $v = (function() {
  $__fn = function($dictFunctor_0 = null, $f_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__local_var_3_0 = ($v_2)->{'value1'};
  $__res = ((($dictFunctor_0)->{'map'})(function($a_4 = null) use ($__local_var_3_0) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("DateTime", $a_4, $__local_var_3_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($f_1)(($v_2)->{'value0'}));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_DateTime_modifyDate'] = function() { $v = (function() {
  $__fn = function($f_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("DateTime", ($f_0)(($v_1)->{'value0'}), ($v_1)->{'value1'});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_DateTime_eqDateTime'] = function() { $v = (object)["eq" => (function() {
  $__fn = function($x_0 = null, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_HeytingAlgebra_boolConj'] ?? \PhpursThunks::eval('Data_HeytingAlgebra_boolConj')))((((($GLOBALS['Data_Date_eqDate'] ?? \PhpursThunks::eval('Data_Date_eqDate')))->{'eq'})(($x_0)->{'value0'}))(($y_1)->{'value0'})))((((($GLOBALS['Data_Time_eqTime'] ?? \PhpursThunks::eval('Data_Time_eqTime')))->{'eq'})(($x_0)->{'value1'}))(($y_1)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Data_DateTime_ordDateTime'] = function() { $v = (object)["compare" => (function() {
  $__fn = function($x_0 = null, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v_2_0 = (((($GLOBALS['Data_Date_ordDate'] ?? \PhpursThunks::eval('Data_Date_ordDate')))->{'compare'})(($x_0)->{'value0'}))(($y_1)->{'value0'});
  $__t1 = null;;
  if ((is_object($v_2_0) && (($v_2_0)->{'tag'} === "LT"))) {
$__t1 = new Phpurs_Data0("LT");
goto end_branch_1;;
};
  if ((is_object($v_2_0) && (($v_2_0)->{'tag'} === "GT"))) {
$__t1 = new Phpurs_Data0("GT");
goto end_branch_1;;
};
  $__t1 = (((($GLOBALS['Data_Time_ordTime'] ?? \PhpursThunks::eval('Data_Time_ordTime')))->{'compare'})(($x_0)->{'value1'}))(($y_1)->{'value1'});
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_DateTime_eqDateTime'] ?? \PhpursThunks::eval('Data_DateTime_eqDateTime'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_DateTime_diff'] = function() { $v = (function() {
  $__fn = function($dictDuration_0 = null, $dt1_1 = null, $dt2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($dictDuration_0)->{'toDuration'})((((($GLOBALS['Data_Function_Uncurried_runFn2'] ?? \PhpursThunks::eval('Data_Function_Uncurried_runFn2')))(($GLOBALS['Data_DateTime_calcDiff'] ?? \PhpursThunks::eval('Data_DateTime_calcDiff'))))((($GLOBALS['Data_DateTime_toRecord'] ?? \PhpursThunks::eval('Data_DateTime_toRecord')))($dt1_1)))((($GLOBALS['Data_DateTime_toRecord'] ?? \PhpursThunks::eval('Data_DateTime_toRecord')))($dt2_2)));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_DateTime_date'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ($v_0)->{'value0'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_DateTime_boundedDateTime'] = function() { $v = (object)["bottom" => new Phpurs_Data2("DateTime", new Phpurs_Data3("Date", (($GLOBALS['Data_Date_Component_boundedYear'] ?? \PhpursThunks::eval('Data_Date_Component_boundedYear')))->{'bottom'}, new Phpurs_Data0("January"), 1), new Phpurs_Data4("Time", 0, 0, 0, 0)), "top" => new Phpurs_Data2("DateTime", new Phpurs_Data3("Date", 275759, new Phpurs_Data0("December"), 31), new Phpurs_Data4("Time", 23, 59, 59, 999)), "Ord0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_DateTime_ordDateTime'] ?? \PhpursThunks::eval('Data_DateTime_ordDateTime'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_DateTime_adjust'] = function() { $v = (function() {
  $__fn = function($dictDuration_0 = null, $d_1 = null, $dt_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__local_var_3_0 = ((((($GLOBALS['Data_DateTime_adjustImpl'] ?? \PhpursThunks::eval('Data_DateTime_adjustImpl')))(($GLOBALS['Data_Maybe_Just'] ?? \PhpursThunks::eval('Data_Maybe_Just'))))(new Phpurs_Data0("Nothing")))((($dictDuration_0)->{'fromDuration'})($d_1)))((($GLOBALS['Data_DateTime_toRecord'] ?? \PhpursThunks::eval('Data_DateTime_toRecord')))($dt_2));
  $__t1 = null;;
  if ((is_object($__local_var_3_0) && (($__local_var_3_0)->{'tag'} === "Just"))) {
$__local_var_4_2 = ((($GLOBALS['Data_Date_Component_boundedEnumYear'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumYear')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'year'});
$__t3 = null;;
if ((is_object($__local_var_4_2) && (($__local_var_4_2)->{'tag'} === "Just"))) {
$__t3 = new Phpurs_Data1("Just", (($GLOBALS['Data_Date_exactDate'] ?? \PhpursThunks::eval('Data_Date_exactDate')))(($__local_var_4_2)->{'value0'}));
goto end_branch_3;;
};
$__t3 = new Phpurs_Data0("Nothing");
end_branch_3:;
$__local_var_5_3 = $__t3;
$__t5 = null;;
switch ((($__local_var_3_0)->{'value0'})->{'month'}) {
case 1:
$__t6 = null;;
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->{'tag'} === "Just"))) {
$__local_var_6_7 = ((($GLOBALS['Data_Date_Component_boundedEnumDay'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumDay')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'day'});
$__t8 = null;;
if ((is_object($__local_var_6_7) && (($__local_var_6_7)->{'tag'} === "Just"))) {
$__local_var_7_9 = ((($__local_var_5_3)->{'value0'})(new Phpurs_Data0("January")))(($__local_var_6_7)->{'value0'});
$__t10 = null;;
if ((is_object($__local_var_7_9) && (($__local_var_7_9)->{'tag'} === "Just"))) {
$__local_var_8_11 = ((($GLOBALS['Data_Time_Component_boundedEnumHour'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumHour')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'hour'});
$__t12 = null;;
if ((is_object($__local_var_8_11) && (($__local_var_8_11)->{'tag'} === "Just"))) {
$__local_var_9_13 = ((($GLOBALS['Data_Time_Component_boundedEnumMinute'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMinute')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'minute'});
$__t14 = null;;
if ((is_object($__local_var_9_13) && (($__local_var_9_13)->{'tag'} === "Just"))) {
$__local_var_10_15 = ((($GLOBALS['Data_Time_Component_boundedEnumSecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumSecond')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'second'});
$__t16 = null;;
if ((is_object($__local_var_10_15) && (($__local_var_10_15)->{'tag'} === "Just"))) {
$__local_var_11_17 = ((($GLOBALS['Data_Time_Component_boundedEnumMillisecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMillisecond')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'millisecond'});
$__t18 = null;;
if ((is_object($__local_var_11_17) && (($__local_var_11_17)->{'tag'} === "Just"))) {
$__t18 = new Phpurs_Data1("Just", new Phpurs_Data2("DateTime", ($__local_var_7_9)->{'value0'}, new Phpurs_Data4("Time", ($__local_var_8_11)->{'value0'}, ($__local_var_9_13)->{'value0'}, ($__local_var_10_15)->{'value0'}, ($__local_var_11_17)->{'value0'})));
goto end_branch_18;;
};
$__t18 = new Phpurs_Data0("Nothing");
end_branch_18:;
$__t16 = $__t18;
goto end_branch_16;;
};
$__t16 = new Phpurs_Data0("Nothing");
end_branch_16:;
$__t14 = $__t16;
goto end_branch_14;;
};
$__t14 = new Phpurs_Data0("Nothing");
end_branch_14:;
$__t12 = $__t14;
goto end_branch_12;;
};
$__t12 = new Phpurs_Data0("Nothing");
end_branch_12:;
$__t10 = $__t12;
goto end_branch_10;;
};
$__t10 = new Phpurs_Data0("Nothing");
end_branch_10:;
$__t8 = $__t10;
goto end_branch_8;;
};
$__t8 = new Phpurs_Data0("Nothing");
end_branch_8:;
$__t6 = $__t8;
goto end_branch_6;;
};
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->{'tag'} === "Nothing"))) {
$__t6 = new Phpurs_Data0("Nothing");
goto end_branch_6;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t6 = null;
end_branch_6:;
$__t5 = $__t6;
goto end_branch_5;;
break;
default:
;
break;
};
switch ((($__local_var_3_0)->{'value0'})->{'month'}) {
case 2:
$__t19 = null;;
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->{'tag'} === "Just"))) {
$__local_var_6_20 = ((($GLOBALS['Data_Date_Component_boundedEnumDay'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumDay')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'day'});
$__t21 = null;;
if ((is_object($__local_var_6_20) && (($__local_var_6_20)->{'tag'} === "Just"))) {
$__local_var_7_22 = ((($__local_var_5_3)->{'value0'})(new Phpurs_Data0("February")))(($__local_var_6_20)->{'value0'});
$__t23 = null;;
if ((is_object($__local_var_7_22) && (($__local_var_7_22)->{'tag'} === "Just"))) {
$__local_var_8_24 = ((($GLOBALS['Data_Time_Component_boundedEnumHour'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumHour')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'hour'});
$__t25 = null;;
if ((is_object($__local_var_8_24) && (($__local_var_8_24)->{'tag'} === "Just"))) {
$__local_var_9_26 = ((($GLOBALS['Data_Time_Component_boundedEnumMinute'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMinute')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'minute'});
$__t27 = null;;
if ((is_object($__local_var_9_26) && (($__local_var_9_26)->{'tag'} === "Just"))) {
$__local_var_10_28 = ((($GLOBALS['Data_Time_Component_boundedEnumSecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumSecond')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'second'});
$__t29 = null;;
if ((is_object($__local_var_10_28) && (($__local_var_10_28)->{'tag'} === "Just"))) {
$__local_var_11_30 = ((($GLOBALS['Data_Time_Component_boundedEnumMillisecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMillisecond')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'millisecond'});
$__t31 = null;;
if ((is_object($__local_var_11_30) && (($__local_var_11_30)->{'tag'} === "Just"))) {
$__t31 = new Phpurs_Data1("Just", new Phpurs_Data2("DateTime", ($__local_var_7_22)->{'value0'}, new Phpurs_Data4("Time", ($__local_var_8_24)->{'value0'}, ($__local_var_9_26)->{'value0'}, ($__local_var_10_28)->{'value0'}, ($__local_var_11_30)->{'value0'})));
goto end_branch_31;;
};
$__t31 = new Phpurs_Data0("Nothing");
end_branch_31:;
$__t29 = $__t31;
goto end_branch_29;;
};
$__t29 = new Phpurs_Data0("Nothing");
end_branch_29:;
$__t27 = $__t29;
goto end_branch_27;;
};
$__t27 = new Phpurs_Data0("Nothing");
end_branch_27:;
$__t25 = $__t27;
goto end_branch_25;;
};
$__t25 = new Phpurs_Data0("Nothing");
end_branch_25:;
$__t23 = $__t25;
goto end_branch_23;;
};
$__t23 = new Phpurs_Data0("Nothing");
end_branch_23:;
$__t21 = $__t23;
goto end_branch_21;;
};
$__t21 = new Phpurs_Data0("Nothing");
end_branch_21:;
$__t19 = $__t21;
goto end_branch_19;;
};
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->{'tag'} === "Nothing"))) {
$__t19 = new Phpurs_Data0("Nothing");
goto end_branch_19;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t19 = null;
end_branch_19:;
$__t5 = $__t19;
goto end_branch_5;;
break;
default:
;
break;
};
switch ((($__local_var_3_0)->{'value0'})->{'month'}) {
case 3:
$__t32 = null;;
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->{'tag'} === "Just"))) {
$__local_var_6_33 = ((($GLOBALS['Data_Date_Component_boundedEnumDay'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumDay')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'day'});
$__t34 = null;;
if ((is_object($__local_var_6_33) && (($__local_var_6_33)->{'tag'} === "Just"))) {
$__local_var_7_35 = ((($__local_var_5_3)->{'value0'})(new Phpurs_Data0("March")))(($__local_var_6_33)->{'value0'});
$__t36 = null;;
if ((is_object($__local_var_7_35) && (($__local_var_7_35)->{'tag'} === "Just"))) {
$__local_var_8_37 = ((($GLOBALS['Data_Time_Component_boundedEnumHour'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumHour')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'hour'});
$__t38 = null;;
if ((is_object($__local_var_8_37) && (($__local_var_8_37)->{'tag'} === "Just"))) {
$__local_var_9_39 = ((($GLOBALS['Data_Time_Component_boundedEnumMinute'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMinute')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'minute'});
$__t40 = null;;
if ((is_object($__local_var_9_39) && (($__local_var_9_39)->{'tag'} === "Just"))) {
$__local_var_10_41 = ((($GLOBALS['Data_Time_Component_boundedEnumSecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumSecond')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'second'});
$__t42 = null;;
if ((is_object($__local_var_10_41) && (($__local_var_10_41)->{'tag'} === "Just"))) {
$__local_var_11_43 = ((($GLOBALS['Data_Time_Component_boundedEnumMillisecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMillisecond')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'millisecond'});
$__t44 = null;;
if ((is_object($__local_var_11_43) && (($__local_var_11_43)->{'tag'} === "Just"))) {
$__t44 = new Phpurs_Data1("Just", new Phpurs_Data2("DateTime", ($__local_var_7_35)->{'value0'}, new Phpurs_Data4("Time", ($__local_var_8_37)->{'value0'}, ($__local_var_9_39)->{'value0'}, ($__local_var_10_41)->{'value0'}, ($__local_var_11_43)->{'value0'})));
goto end_branch_44;;
};
$__t44 = new Phpurs_Data0("Nothing");
end_branch_44:;
$__t42 = $__t44;
goto end_branch_42;;
};
$__t42 = new Phpurs_Data0("Nothing");
end_branch_42:;
$__t40 = $__t42;
goto end_branch_40;;
};
$__t40 = new Phpurs_Data0("Nothing");
end_branch_40:;
$__t38 = $__t40;
goto end_branch_38;;
};
$__t38 = new Phpurs_Data0("Nothing");
end_branch_38:;
$__t36 = $__t38;
goto end_branch_36;;
};
$__t36 = new Phpurs_Data0("Nothing");
end_branch_36:;
$__t34 = $__t36;
goto end_branch_34;;
};
$__t34 = new Phpurs_Data0("Nothing");
end_branch_34:;
$__t32 = $__t34;
goto end_branch_32;;
};
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->{'tag'} === "Nothing"))) {
$__t32 = new Phpurs_Data0("Nothing");
goto end_branch_32;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t32 = null;
end_branch_32:;
$__t5 = $__t32;
goto end_branch_5;;
break;
default:
;
break;
};
switch ((($__local_var_3_0)->{'value0'})->{'month'}) {
case 4:
$__t45 = null;;
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->{'tag'} === "Just"))) {
$__local_var_6_46 = ((($GLOBALS['Data_Date_Component_boundedEnumDay'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumDay')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'day'});
$__t47 = null;;
if ((is_object($__local_var_6_46) && (($__local_var_6_46)->{'tag'} === "Just"))) {
$__local_var_7_48 = ((($__local_var_5_3)->{'value0'})(new Phpurs_Data0("April")))(($__local_var_6_46)->{'value0'});
$__t49 = null;;
if ((is_object($__local_var_7_48) && (($__local_var_7_48)->{'tag'} === "Just"))) {
$__local_var_8_50 = ((($GLOBALS['Data_Time_Component_boundedEnumHour'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumHour')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'hour'});
$__t51 = null;;
if ((is_object($__local_var_8_50) && (($__local_var_8_50)->{'tag'} === "Just"))) {
$__local_var_9_52 = ((($GLOBALS['Data_Time_Component_boundedEnumMinute'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMinute')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'minute'});
$__t53 = null;;
if ((is_object($__local_var_9_52) && (($__local_var_9_52)->{'tag'} === "Just"))) {
$__local_var_10_54 = ((($GLOBALS['Data_Time_Component_boundedEnumSecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumSecond')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'second'});
$__t55 = null;;
if ((is_object($__local_var_10_54) && (($__local_var_10_54)->{'tag'} === "Just"))) {
$__local_var_11_56 = ((($GLOBALS['Data_Time_Component_boundedEnumMillisecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMillisecond')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'millisecond'});
$__t57 = null;;
if ((is_object($__local_var_11_56) && (($__local_var_11_56)->{'tag'} === "Just"))) {
$__t57 = new Phpurs_Data1("Just", new Phpurs_Data2("DateTime", ($__local_var_7_48)->{'value0'}, new Phpurs_Data4("Time", ($__local_var_8_50)->{'value0'}, ($__local_var_9_52)->{'value0'}, ($__local_var_10_54)->{'value0'}, ($__local_var_11_56)->{'value0'})));
goto end_branch_57;;
};
$__t57 = new Phpurs_Data0("Nothing");
end_branch_57:;
$__t55 = $__t57;
goto end_branch_55;;
};
$__t55 = new Phpurs_Data0("Nothing");
end_branch_55:;
$__t53 = $__t55;
goto end_branch_53;;
};
$__t53 = new Phpurs_Data0("Nothing");
end_branch_53:;
$__t51 = $__t53;
goto end_branch_51;;
};
$__t51 = new Phpurs_Data0("Nothing");
end_branch_51:;
$__t49 = $__t51;
goto end_branch_49;;
};
$__t49 = new Phpurs_Data0("Nothing");
end_branch_49:;
$__t47 = $__t49;
goto end_branch_47;;
};
$__t47 = new Phpurs_Data0("Nothing");
end_branch_47:;
$__t45 = $__t47;
goto end_branch_45;;
};
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->{'tag'} === "Nothing"))) {
$__t45 = new Phpurs_Data0("Nothing");
goto end_branch_45;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t45 = null;
end_branch_45:;
$__t5 = $__t45;
goto end_branch_5;;
break;
default:
;
break;
};
switch ((($__local_var_3_0)->{'value0'})->{'month'}) {
case 5:
$__t58 = null;;
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->{'tag'} === "Just"))) {
$__local_var_6_59 = ((($GLOBALS['Data_Date_Component_boundedEnumDay'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumDay')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'day'});
$__t60 = null;;
if ((is_object($__local_var_6_59) && (($__local_var_6_59)->{'tag'} === "Just"))) {
$__local_var_7_61 = ((($__local_var_5_3)->{'value0'})(new Phpurs_Data0("May")))(($__local_var_6_59)->{'value0'});
$__t62 = null;;
if ((is_object($__local_var_7_61) && (($__local_var_7_61)->{'tag'} === "Just"))) {
$__local_var_8_63 = ((($GLOBALS['Data_Time_Component_boundedEnumHour'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumHour')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'hour'});
$__t64 = null;;
if ((is_object($__local_var_8_63) && (($__local_var_8_63)->{'tag'} === "Just"))) {
$__local_var_9_65 = ((($GLOBALS['Data_Time_Component_boundedEnumMinute'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMinute')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'minute'});
$__t66 = null;;
if ((is_object($__local_var_9_65) && (($__local_var_9_65)->{'tag'} === "Just"))) {
$__local_var_10_67 = ((($GLOBALS['Data_Time_Component_boundedEnumSecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumSecond')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'second'});
$__t68 = null;;
if ((is_object($__local_var_10_67) && (($__local_var_10_67)->{'tag'} === "Just"))) {
$__local_var_11_69 = ((($GLOBALS['Data_Time_Component_boundedEnumMillisecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMillisecond')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'millisecond'});
$__t70 = null;;
if ((is_object($__local_var_11_69) && (($__local_var_11_69)->{'tag'} === "Just"))) {
$__t70 = new Phpurs_Data1("Just", new Phpurs_Data2("DateTime", ($__local_var_7_61)->{'value0'}, new Phpurs_Data4("Time", ($__local_var_8_63)->{'value0'}, ($__local_var_9_65)->{'value0'}, ($__local_var_10_67)->{'value0'}, ($__local_var_11_69)->{'value0'})));
goto end_branch_70;;
};
$__t70 = new Phpurs_Data0("Nothing");
end_branch_70:;
$__t68 = $__t70;
goto end_branch_68;;
};
$__t68 = new Phpurs_Data0("Nothing");
end_branch_68:;
$__t66 = $__t68;
goto end_branch_66;;
};
$__t66 = new Phpurs_Data0("Nothing");
end_branch_66:;
$__t64 = $__t66;
goto end_branch_64;;
};
$__t64 = new Phpurs_Data0("Nothing");
end_branch_64:;
$__t62 = $__t64;
goto end_branch_62;;
};
$__t62 = new Phpurs_Data0("Nothing");
end_branch_62:;
$__t60 = $__t62;
goto end_branch_60;;
};
$__t60 = new Phpurs_Data0("Nothing");
end_branch_60:;
$__t58 = $__t60;
goto end_branch_58;;
};
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->{'tag'} === "Nothing"))) {
$__t58 = new Phpurs_Data0("Nothing");
goto end_branch_58;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t58 = null;
end_branch_58:;
$__t5 = $__t58;
goto end_branch_5;;
break;
default:
;
break;
};
switch ((($__local_var_3_0)->{'value0'})->{'month'}) {
case 6:
$__t71 = null;;
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->{'tag'} === "Just"))) {
$__local_var_6_72 = ((($GLOBALS['Data_Date_Component_boundedEnumDay'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumDay')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'day'});
$__t73 = null;;
if ((is_object($__local_var_6_72) && (($__local_var_6_72)->{'tag'} === "Just"))) {
$__local_var_7_74 = ((($__local_var_5_3)->{'value0'})(new Phpurs_Data0("June")))(($__local_var_6_72)->{'value0'});
$__t75 = null;;
if ((is_object($__local_var_7_74) && (($__local_var_7_74)->{'tag'} === "Just"))) {
$__local_var_8_76 = ((($GLOBALS['Data_Time_Component_boundedEnumHour'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumHour')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'hour'});
$__t77 = null;;
if ((is_object($__local_var_8_76) && (($__local_var_8_76)->{'tag'} === "Just"))) {
$__local_var_9_78 = ((($GLOBALS['Data_Time_Component_boundedEnumMinute'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMinute')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'minute'});
$__t79 = null;;
if ((is_object($__local_var_9_78) && (($__local_var_9_78)->{'tag'} === "Just"))) {
$__local_var_10_80 = ((($GLOBALS['Data_Time_Component_boundedEnumSecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumSecond')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'second'});
$__t81 = null;;
if ((is_object($__local_var_10_80) && (($__local_var_10_80)->{'tag'} === "Just"))) {
$__local_var_11_82 = ((($GLOBALS['Data_Time_Component_boundedEnumMillisecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMillisecond')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'millisecond'});
$__t83 = null;;
if ((is_object($__local_var_11_82) && (($__local_var_11_82)->{'tag'} === "Just"))) {
$__t83 = new Phpurs_Data1("Just", new Phpurs_Data2("DateTime", ($__local_var_7_74)->{'value0'}, new Phpurs_Data4("Time", ($__local_var_8_76)->{'value0'}, ($__local_var_9_78)->{'value0'}, ($__local_var_10_80)->{'value0'}, ($__local_var_11_82)->{'value0'})));
goto end_branch_83;;
};
$__t83 = new Phpurs_Data0("Nothing");
end_branch_83:;
$__t81 = $__t83;
goto end_branch_81;;
};
$__t81 = new Phpurs_Data0("Nothing");
end_branch_81:;
$__t79 = $__t81;
goto end_branch_79;;
};
$__t79 = new Phpurs_Data0("Nothing");
end_branch_79:;
$__t77 = $__t79;
goto end_branch_77;;
};
$__t77 = new Phpurs_Data0("Nothing");
end_branch_77:;
$__t75 = $__t77;
goto end_branch_75;;
};
$__t75 = new Phpurs_Data0("Nothing");
end_branch_75:;
$__t73 = $__t75;
goto end_branch_73;;
};
$__t73 = new Phpurs_Data0("Nothing");
end_branch_73:;
$__t71 = $__t73;
goto end_branch_71;;
};
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->{'tag'} === "Nothing"))) {
$__t71 = new Phpurs_Data0("Nothing");
goto end_branch_71;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t71 = null;
end_branch_71:;
$__t5 = $__t71;
goto end_branch_5;;
break;
default:
;
break;
};
switch ((($__local_var_3_0)->{'value0'})->{'month'}) {
case 7:
$__t84 = null;;
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->{'tag'} === "Just"))) {
$__local_var_6_85 = ((($GLOBALS['Data_Date_Component_boundedEnumDay'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumDay')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'day'});
$__t86 = null;;
if ((is_object($__local_var_6_85) && (($__local_var_6_85)->{'tag'} === "Just"))) {
$__local_var_7_87 = ((($__local_var_5_3)->{'value0'})(new Phpurs_Data0("July")))(($__local_var_6_85)->{'value0'});
$__t88 = null;;
if ((is_object($__local_var_7_87) && (($__local_var_7_87)->{'tag'} === "Just"))) {
$__local_var_8_89 = ((($GLOBALS['Data_Time_Component_boundedEnumHour'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumHour')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'hour'});
$__t90 = null;;
if ((is_object($__local_var_8_89) && (($__local_var_8_89)->{'tag'} === "Just"))) {
$__local_var_9_91 = ((($GLOBALS['Data_Time_Component_boundedEnumMinute'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMinute')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'minute'});
$__t92 = null;;
if ((is_object($__local_var_9_91) && (($__local_var_9_91)->{'tag'} === "Just"))) {
$__local_var_10_93 = ((($GLOBALS['Data_Time_Component_boundedEnumSecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumSecond')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'second'});
$__t94 = null;;
if ((is_object($__local_var_10_93) && (($__local_var_10_93)->{'tag'} === "Just"))) {
$__local_var_11_95 = ((($GLOBALS['Data_Time_Component_boundedEnumMillisecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMillisecond')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'millisecond'});
$__t96 = null;;
if ((is_object($__local_var_11_95) && (($__local_var_11_95)->{'tag'} === "Just"))) {
$__t96 = new Phpurs_Data1("Just", new Phpurs_Data2("DateTime", ($__local_var_7_87)->{'value0'}, new Phpurs_Data4("Time", ($__local_var_8_89)->{'value0'}, ($__local_var_9_91)->{'value0'}, ($__local_var_10_93)->{'value0'}, ($__local_var_11_95)->{'value0'})));
goto end_branch_96;;
};
$__t96 = new Phpurs_Data0("Nothing");
end_branch_96:;
$__t94 = $__t96;
goto end_branch_94;;
};
$__t94 = new Phpurs_Data0("Nothing");
end_branch_94:;
$__t92 = $__t94;
goto end_branch_92;;
};
$__t92 = new Phpurs_Data0("Nothing");
end_branch_92:;
$__t90 = $__t92;
goto end_branch_90;;
};
$__t90 = new Phpurs_Data0("Nothing");
end_branch_90:;
$__t88 = $__t90;
goto end_branch_88;;
};
$__t88 = new Phpurs_Data0("Nothing");
end_branch_88:;
$__t86 = $__t88;
goto end_branch_86;;
};
$__t86 = new Phpurs_Data0("Nothing");
end_branch_86:;
$__t84 = $__t86;
goto end_branch_84;;
};
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->{'tag'} === "Nothing"))) {
$__t84 = new Phpurs_Data0("Nothing");
goto end_branch_84;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t84 = null;
end_branch_84:;
$__t5 = $__t84;
goto end_branch_5;;
break;
default:
;
break;
};
switch ((($__local_var_3_0)->{'value0'})->{'month'}) {
case 8:
$__t97 = null;;
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->{'tag'} === "Just"))) {
$__local_var_6_98 = ((($GLOBALS['Data_Date_Component_boundedEnumDay'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumDay')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'day'});
$__t99 = null;;
if ((is_object($__local_var_6_98) && (($__local_var_6_98)->{'tag'} === "Just"))) {
$__local_var_7_100 = ((($__local_var_5_3)->{'value0'})(new Phpurs_Data0("August")))(($__local_var_6_98)->{'value0'});
$__t101 = null;;
if ((is_object($__local_var_7_100) && (($__local_var_7_100)->{'tag'} === "Just"))) {
$__local_var_8_102 = ((($GLOBALS['Data_Time_Component_boundedEnumHour'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumHour')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'hour'});
$__t103 = null;;
if ((is_object($__local_var_8_102) && (($__local_var_8_102)->{'tag'} === "Just"))) {
$__local_var_9_104 = ((($GLOBALS['Data_Time_Component_boundedEnumMinute'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMinute')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'minute'});
$__t105 = null;;
if ((is_object($__local_var_9_104) && (($__local_var_9_104)->{'tag'} === "Just"))) {
$__local_var_10_106 = ((($GLOBALS['Data_Time_Component_boundedEnumSecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumSecond')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'second'});
$__t107 = null;;
if ((is_object($__local_var_10_106) && (($__local_var_10_106)->{'tag'} === "Just"))) {
$__local_var_11_108 = ((($GLOBALS['Data_Time_Component_boundedEnumMillisecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMillisecond')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'millisecond'});
$__t109 = null;;
if ((is_object($__local_var_11_108) && (($__local_var_11_108)->{'tag'} === "Just"))) {
$__t109 = new Phpurs_Data1("Just", new Phpurs_Data2("DateTime", ($__local_var_7_100)->{'value0'}, new Phpurs_Data4("Time", ($__local_var_8_102)->{'value0'}, ($__local_var_9_104)->{'value0'}, ($__local_var_10_106)->{'value0'}, ($__local_var_11_108)->{'value0'})));
goto end_branch_109;;
};
$__t109 = new Phpurs_Data0("Nothing");
end_branch_109:;
$__t107 = $__t109;
goto end_branch_107;;
};
$__t107 = new Phpurs_Data0("Nothing");
end_branch_107:;
$__t105 = $__t107;
goto end_branch_105;;
};
$__t105 = new Phpurs_Data0("Nothing");
end_branch_105:;
$__t103 = $__t105;
goto end_branch_103;;
};
$__t103 = new Phpurs_Data0("Nothing");
end_branch_103:;
$__t101 = $__t103;
goto end_branch_101;;
};
$__t101 = new Phpurs_Data0("Nothing");
end_branch_101:;
$__t99 = $__t101;
goto end_branch_99;;
};
$__t99 = new Phpurs_Data0("Nothing");
end_branch_99:;
$__t97 = $__t99;
goto end_branch_97;;
};
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->{'tag'} === "Nothing"))) {
$__t97 = new Phpurs_Data0("Nothing");
goto end_branch_97;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t97 = null;
end_branch_97:;
$__t5 = $__t97;
goto end_branch_5;;
break;
default:
;
break;
};
switch ((($__local_var_3_0)->{'value0'})->{'month'}) {
case 9:
$__t110 = null;;
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->{'tag'} === "Just"))) {
$__local_var_6_111 = ((($GLOBALS['Data_Date_Component_boundedEnumDay'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumDay')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'day'});
$__t112 = null;;
if ((is_object($__local_var_6_111) && (($__local_var_6_111)->{'tag'} === "Just"))) {
$__local_var_7_113 = ((($__local_var_5_3)->{'value0'})(new Phpurs_Data0("September")))(($__local_var_6_111)->{'value0'});
$__t114 = null;;
if ((is_object($__local_var_7_113) && (($__local_var_7_113)->{'tag'} === "Just"))) {
$__local_var_8_115 = ((($GLOBALS['Data_Time_Component_boundedEnumHour'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumHour')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'hour'});
$__t116 = null;;
if ((is_object($__local_var_8_115) && (($__local_var_8_115)->{'tag'} === "Just"))) {
$__local_var_9_117 = ((($GLOBALS['Data_Time_Component_boundedEnumMinute'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMinute')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'minute'});
$__t118 = null;;
if ((is_object($__local_var_9_117) && (($__local_var_9_117)->{'tag'} === "Just"))) {
$__local_var_10_119 = ((($GLOBALS['Data_Time_Component_boundedEnumSecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumSecond')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'second'});
$__t120 = null;;
if ((is_object($__local_var_10_119) && (($__local_var_10_119)->{'tag'} === "Just"))) {
$__local_var_11_121 = ((($GLOBALS['Data_Time_Component_boundedEnumMillisecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMillisecond')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'millisecond'});
$__t122 = null;;
if ((is_object($__local_var_11_121) && (($__local_var_11_121)->{'tag'} === "Just"))) {
$__t122 = new Phpurs_Data1("Just", new Phpurs_Data2("DateTime", ($__local_var_7_113)->{'value0'}, new Phpurs_Data4("Time", ($__local_var_8_115)->{'value0'}, ($__local_var_9_117)->{'value0'}, ($__local_var_10_119)->{'value0'}, ($__local_var_11_121)->{'value0'})));
goto end_branch_122;;
};
$__t122 = new Phpurs_Data0("Nothing");
end_branch_122:;
$__t120 = $__t122;
goto end_branch_120;;
};
$__t120 = new Phpurs_Data0("Nothing");
end_branch_120:;
$__t118 = $__t120;
goto end_branch_118;;
};
$__t118 = new Phpurs_Data0("Nothing");
end_branch_118:;
$__t116 = $__t118;
goto end_branch_116;;
};
$__t116 = new Phpurs_Data0("Nothing");
end_branch_116:;
$__t114 = $__t116;
goto end_branch_114;;
};
$__t114 = new Phpurs_Data0("Nothing");
end_branch_114:;
$__t112 = $__t114;
goto end_branch_112;;
};
$__t112 = new Phpurs_Data0("Nothing");
end_branch_112:;
$__t110 = $__t112;
goto end_branch_110;;
};
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->{'tag'} === "Nothing"))) {
$__t110 = new Phpurs_Data0("Nothing");
goto end_branch_110;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t110 = null;
end_branch_110:;
$__t5 = $__t110;
goto end_branch_5;;
break;
default:
;
break;
};
switch ((($__local_var_3_0)->{'value0'})->{'month'}) {
case 10:
$__t123 = null;;
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->{'tag'} === "Just"))) {
$__local_var_6_124 = ((($GLOBALS['Data_Date_Component_boundedEnumDay'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumDay')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'day'});
$__t125 = null;;
if ((is_object($__local_var_6_124) && (($__local_var_6_124)->{'tag'} === "Just"))) {
$__local_var_7_126 = ((($__local_var_5_3)->{'value0'})(new Phpurs_Data0("October")))(($__local_var_6_124)->{'value0'});
$__t127 = null;;
if ((is_object($__local_var_7_126) && (($__local_var_7_126)->{'tag'} === "Just"))) {
$__local_var_8_128 = ((($GLOBALS['Data_Time_Component_boundedEnumHour'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumHour')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'hour'});
$__t129 = null;;
if ((is_object($__local_var_8_128) && (($__local_var_8_128)->{'tag'} === "Just"))) {
$__local_var_9_130 = ((($GLOBALS['Data_Time_Component_boundedEnumMinute'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMinute')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'minute'});
$__t131 = null;;
if ((is_object($__local_var_9_130) && (($__local_var_9_130)->{'tag'} === "Just"))) {
$__local_var_10_132 = ((($GLOBALS['Data_Time_Component_boundedEnumSecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumSecond')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'second'});
$__t133 = null;;
if ((is_object($__local_var_10_132) && (($__local_var_10_132)->{'tag'} === "Just"))) {
$__local_var_11_134 = ((($GLOBALS['Data_Time_Component_boundedEnumMillisecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMillisecond')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'millisecond'});
$__t135 = null;;
if ((is_object($__local_var_11_134) && (($__local_var_11_134)->{'tag'} === "Just"))) {
$__t135 = new Phpurs_Data1("Just", new Phpurs_Data2("DateTime", ($__local_var_7_126)->{'value0'}, new Phpurs_Data4("Time", ($__local_var_8_128)->{'value0'}, ($__local_var_9_130)->{'value0'}, ($__local_var_10_132)->{'value0'}, ($__local_var_11_134)->{'value0'})));
goto end_branch_135;;
};
$__t135 = new Phpurs_Data0("Nothing");
end_branch_135:;
$__t133 = $__t135;
goto end_branch_133;;
};
$__t133 = new Phpurs_Data0("Nothing");
end_branch_133:;
$__t131 = $__t133;
goto end_branch_131;;
};
$__t131 = new Phpurs_Data0("Nothing");
end_branch_131:;
$__t129 = $__t131;
goto end_branch_129;;
};
$__t129 = new Phpurs_Data0("Nothing");
end_branch_129:;
$__t127 = $__t129;
goto end_branch_127;;
};
$__t127 = new Phpurs_Data0("Nothing");
end_branch_127:;
$__t125 = $__t127;
goto end_branch_125;;
};
$__t125 = new Phpurs_Data0("Nothing");
end_branch_125:;
$__t123 = $__t125;
goto end_branch_123;;
};
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->{'tag'} === "Nothing"))) {
$__t123 = new Phpurs_Data0("Nothing");
goto end_branch_123;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t123 = null;
end_branch_123:;
$__t5 = $__t123;
goto end_branch_5;;
break;
default:
;
break;
};
switch ((($__local_var_3_0)->{'value0'})->{'month'}) {
case 11:
$__t136 = null;;
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->{'tag'} === "Just"))) {
$__local_var_6_137 = ((($GLOBALS['Data_Date_Component_boundedEnumDay'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumDay')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'day'});
$__t138 = null;;
if ((is_object($__local_var_6_137) && (($__local_var_6_137)->{'tag'} === "Just"))) {
$__local_var_7_139 = ((($__local_var_5_3)->{'value0'})(new Phpurs_Data0("November")))(($__local_var_6_137)->{'value0'});
$__t140 = null;;
if ((is_object($__local_var_7_139) && (($__local_var_7_139)->{'tag'} === "Just"))) {
$__local_var_8_141 = ((($GLOBALS['Data_Time_Component_boundedEnumHour'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumHour')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'hour'});
$__t142 = null;;
if ((is_object($__local_var_8_141) && (($__local_var_8_141)->{'tag'} === "Just"))) {
$__local_var_9_143 = ((($GLOBALS['Data_Time_Component_boundedEnumMinute'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMinute')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'minute'});
$__t144 = null;;
if ((is_object($__local_var_9_143) && (($__local_var_9_143)->{'tag'} === "Just"))) {
$__local_var_10_145 = ((($GLOBALS['Data_Time_Component_boundedEnumSecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumSecond')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'second'});
$__t146 = null;;
if ((is_object($__local_var_10_145) && (($__local_var_10_145)->{'tag'} === "Just"))) {
$__local_var_11_147 = ((($GLOBALS['Data_Time_Component_boundedEnumMillisecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMillisecond')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'millisecond'});
$__t148 = null;;
if ((is_object($__local_var_11_147) && (($__local_var_11_147)->{'tag'} === "Just"))) {
$__t148 = new Phpurs_Data1("Just", new Phpurs_Data2("DateTime", ($__local_var_7_139)->{'value0'}, new Phpurs_Data4("Time", ($__local_var_8_141)->{'value0'}, ($__local_var_9_143)->{'value0'}, ($__local_var_10_145)->{'value0'}, ($__local_var_11_147)->{'value0'})));
goto end_branch_148;;
};
$__t148 = new Phpurs_Data0("Nothing");
end_branch_148:;
$__t146 = $__t148;
goto end_branch_146;;
};
$__t146 = new Phpurs_Data0("Nothing");
end_branch_146:;
$__t144 = $__t146;
goto end_branch_144;;
};
$__t144 = new Phpurs_Data0("Nothing");
end_branch_144:;
$__t142 = $__t144;
goto end_branch_142;;
};
$__t142 = new Phpurs_Data0("Nothing");
end_branch_142:;
$__t140 = $__t142;
goto end_branch_140;;
};
$__t140 = new Phpurs_Data0("Nothing");
end_branch_140:;
$__t138 = $__t140;
goto end_branch_138;;
};
$__t138 = new Phpurs_Data0("Nothing");
end_branch_138:;
$__t136 = $__t138;
goto end_branch_136;;
};
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->{'tag'} === "Nothing"))) {
$__t136 = new Phpurs_Data0("Nothing");
goto end_branch_136;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t136 = null;
end_branch_136:;
$__t5 = $__t136;
goto end_branch_5;;
break;
default:
;
break;
};
switch ((($__local_var_3_0)->{'value0'})->{'month'}) {
case 12:
$__t149 = null;;
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->{'tag'} === "Just"))) {
$__local_var_6_150 = ((($GLOBALS['Data_Date_Component_boundedEnumDay'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumDay')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'day'});
$__t151 = null;;
if ((is_object($__local_var_6_150) && (($__local_var_6_150)->{'tag'} === "Just"))) {
$__local_var_7_152 = ((($__local_var_5_3)->{'value0'})(new Phpurs_Data0("December")))(($__local_var_6_150)->{'value0'});
$__t153 = null;;
if ((is_object($__local_var_7_152) && (($__local_var_7_152)->{'tag'} === "Just"))) {
$__local_var_8_154 = ((($GLOBALS['Data_Time_Component_boundedEnumHour'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumHour')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'hour'});
$__t155 = null;;
if ((is_object($__local_var_8_154) && (($__local_var_8_154)->{'tag'} === "Just"))) {
$__local_var_9_156 = ((($GLOBALS['Data_Time_Component_boundedEnumMinute'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMinute')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'minute'});
$__t157 = null;;
if ((is_object($__local_var_9_156) && (($__local_var_9_156)->{'tag'} === "Just"))) {
$__local_var_10_158 = ((($GLOBALS['Data_Time_Component_boundedEnumSecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumSecond')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'second'});
$__t159 = null;;
if ((is_object($__local_var_10_158) && (($__local_var_10_158)->{'tag'} === "Just"))) {
$__local_var_11_160 = ((($GLOBALS['Data_Time_Component_boundedEnumMillisecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMillisecond')))->{'toEnum'})((($__local_var_3_0)->{'value0'})->{'millisecond'});
$__t161 = null;;
if ((is_object($__local_var_11_160) && (($__local_var_11_160)->{'tag'} === "Just"))) {
$__t161 = new Phpurs_Data1("Just", new Phpurs_Data2("DateTime", ($__local_var_7_152)->{'value0'}, new Phpurs_Data4("Time", ($__local_var_8_154)->{'value0'}, ($__local_var_9_156)->{'value0'}, ($__local_var_10_158)->{'value0'}, ($__local_var_11_160)->{'value0'})));
goto end_branch_161;;
};
$__t161 = new Phpurs_Data0("Nothing");
end_branch_161:;
$__t159 = $__t161;
goto end_branch_159;;
};
$__t159 = new Phpurs_Data0("Nothing");
end_branch_159:;
$__t157 = $__t159;
goto end_branch_157;;
};
$__t157 = new Phpurs_Data0("Nothing");
end_branch_157:;
$__t155 = $__t157;
goto end_branch_155;;
};
$__t155 = new Phpurs_Data0("Nothing");
end_branch_155:;
$__t153 = $__t155;
goto end_branch_153;;
};
$__t153 = new Phpurs_Data0("Nothing");
end_branch_153:;
$__t151 = $__t153;
goto end_branch_151;;
};
$__t151 = new Phpurs_Data0("Nothing");
end_branch_151:;
$__t149 = $__t151;
goto end_branch_149;;
};
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->{'tag'} === "Nothing"))) {
$__t149 = new Phpurs_Data0("Nothing");
goto end_branch_149;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t149 = null;
end_branch_149:;
$__t5 = $__t149;
goto end_branch_5;;
break;
default:
;
break;
};
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->{'tag'} === "Just"))) {
$__t5 = new Phpurs_Data0("Nothing");
goto end_branch_5;;
};
if ((is_object($__local_var_5_3) && (($__local_var_5_3)->{'tag'} === "Nothing"))) {
$__t5 = new Phpurs_Data0("Nothing");
goto end_branch_5;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t5 = null;
end_branch_5:;
$__t1 = $__t5;
goto end_branch_1;;
};
  if ((is_object($__local_var_3_0) && (($__local_var_3_0)->{'tag'} === "Nothing"))) {
$__t1 = new Phpurs_Data0("Nothing");
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };
$ffi_Data_DateTime = \call_user_func(function() {
  $exports = [];
$createUTC = function($y, $mo, $d, $h, $m, $s, $ms) {
    $dt = new \DateTime('now', new \DateTimeZone('UTC'));
    $dt->setDate($y, $mo + 1, $d);
    $dt->setTime($h, $m, $s, $ms * 1000);
    return (float)$dt->getTimestamp() * 1000 + (int)$dt->format('v');
};

$calcDiff = function($rec1, $rec2 = null) use (&$calcDiff) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$calcDiff) {

            return $calcDiff(...\array_merge($__args, $more));
        };
    }

    $msUTC1 = $createUTC($rec1->year, $rec1->month - 1, $rec1->day, $rec1->hour, $rec1->minute, $rec1->second, $rec1->millisecond);
    $msUTC2 = $createUTC($rec2->year, $rec2->month - 1, $rec2->day, $rec2->hour, $rec2->minute, $rec2->second, $rec2->millisecond);
    return $msUTC1 - $msUTC2;
};

$adjustImpl = function($just, $nothing = null, $offset = null, $rec = null) use (&$adjustImpl) {
    if (\func_num_args() < 4) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$adjustImpl) {

            return $adjustImpl(...\array_merge($__args, $more));
        };
    }

    $msUTC = $createUTC($rec->year, $rec->month - 1, $rec->day, $rec->hour, $rec->minute, $rec->second, $rec->millisecond);
    $targetMs = $msUTC + $offset;
    
    $seconds = floor($targetMs / 1000);
    $ms = $targetMs - ($seconds * 1000);
    
    try {
        $dt = new \DateTime("@" . $seconds, new \DateTimeZone('UTC'));
        return $just((object)[
            'year' => (int)$dt->format('Y'),
            'month' => (int)$dt->format('n'),
            'day' => (int)$dt->format('j'),
            'hour' => (int)$dt->format('G'),
            'minute' => (int)$dt->format('i'),
            'second' => (int)$dt->format('s'),
            'millisecond' => (int)$ms
        ]);
    } catch (\Exception $e) {
        return $nothing;
    }
};

$exports['createUTC'] = $createUTC;
$exports['calcDiff'] = $calcDiff;
$exports['adjustImpl'] = $adjustImpl;
return $exports;
  return $exports;
});
\PhpursThunks::$thunks['Data_DateTime_adjustImpl'] = function() use (&$ffi_Data_DateTime) { return $ffi_Data_DateTime['adjustImpl']; };
\PhpursThunks::$thunks['Data_DateTime_calcDiff'] = function() use (&$ffi_Data_DateTime) { return $ffi_Data_DateTime['calcDiff']; };
















