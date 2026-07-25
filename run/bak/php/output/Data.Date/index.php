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
\PhpursThunks::$thunks['Data_Date_fromJust'] = function() { $v = function($v_0 = null) {
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
}; return $v; };
\PhpursThunks::$thunks['Data_Date_greaterThan'] = function() { $v = (function() use (&$__fn) {
$__local_var_0_0 = (($GLOBALS['Data_Maybe_ordMaybe'] ?? \PhpursThunks::eval('Data_Maybe_ordMaybe')))(($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')));
return (function() use ($__local_var_0_0) {
  $__fn = function($a1_1 = null, $a2_2 = null) use ($__local_var_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (is_object(((($__local_var_0_0)->{'compare'})($a1_1))($a2_2)) && ((((($__local_var_0_0)->{'compare'})($a1_1))($a2_2))->{'tag'} === "GT"));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
})(); return $v; };
\PhpursThunks::$thunks['Data_Date_Date'] = function() { $v = (function() {
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
})(); return $v; };
\PhpursThunks::$thunks['Data_Date_year'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ($v_0)->{'value0'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Date_weekday'] = function() { $v = (($GLOBALS['Partial_Unsafe__unsafePartial'] ?? \PhpursThunks::eval('Partial_Unsafe__unsafePartial')))((function() {
  $__fn = function($dollar__unused_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ((is_object(($v_1)->{'value1'}) && ((($v_1)->{'value1'})->{'tag'} === "January"))) {
$__t0 = 1;
goto end_branch_0;;
};
  if ((is_object(($v_1)->{'value1'}) && ((($v_1)->{'value1'})->{'tag'} === "February"))) {
$__t0 = 2;
goto end_branch_0;;
};
  if ((is_object(($v_1)->{'value1'}) && ((($v_1)->{'value1'})->{'tag'} === "March"))) {
$__t0 = 3;
goto end_branch_0;;
};
  if ((is_object(($v_1)->{'value1'}) && ((($v_1)->{'value1'})->{'tag'} === "April"))) {
$__t0 = 4;
goto end_branch_0;;
};
  if ((is_object(($v_1)->{'value1'}) && ((($v_1)->{'value1'})->{'tag'} === "May"))) {
$__t0 = 5;
goto end_branch_0;;
};
  if ((is_object(($v_1)->{'value1'}) && ((($v_1)->{'value1'})->{'tag'} === "June"))) {
$__t0 = 6;
goto end_branch_0;;
};
  if ((is_object(($v_1)->{'value1'}) && ((($v_1)->{'value1'})->{'tag'} === "July"))) {
$__t0 = 7;
goto end_branch_0;;
};
  if ((is_object(($v_1)->{'value1'}) && ((($v_1)->{'value1'})->{'tag'} === "August"))) {
$__t0 = 8;
goto end_branch_0;;
};
  if ((is_object(($v_1)->{'value1'}) && ((($v_1)->{'value1'})->{'tag'} === "September"))) {
$__t0 = 9;
goto end_branch_0;;
};
  if ((is_object(($v_1)->{'value1'}) && ((($v_1)->{'value1'})->{'tag'} === "October"))) {
$__t0 = 10;
goto end_branch_0;;
};
  if ((is_object(($v_1)->{'value1'}) && ((($v_1)->{'value1'})->{'tag'} === "November"))) {
$__t0 = 11;
goto end_branch_0;;
};
  if ((is_object(($v_1)->{'value1'}) && ((($v_1)->{'value1'})->{'tag'} === "December"))) {
$__t0 = 12;
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $n_2_0 = ((((($GLOBALS['Data_Function_Uncurried_runFn3'] ?? \PhpursThunks::eval('Data_Function_Uncurried_runFn3')))(($GLOBALS['Data_Date_calcWeekday'] ?? \PhpursThunks::eval('Data_Date_calcWeekday'))))(($v_1)->{'value0'}))($__t0))(($v_1)->{'value2'});
  $__t2 = null;;
  if (((($GLOBALS['Data_Eq_eqIntImpl'] ?? \PhpursThunks::eval('Data_Eq_eqIntImpl')))($n_2_0))(0)) {
$__t2 = new Phpurs_Data0("Sunday");
goto end_branch_2;;
};
  switch ($n_2_0) {
case 1:
$__t2 = new Phpurs_Data0("Monday");
goto end_branch_2;;
break;
default:
;
break;
};
  switch ($n_2_0) {
case 2:
$__t2 = new Phpurs_Data0("Tuesday");
goto end_branch_2;;
break;
default:
;
break;
};
  switch ($n_2_0) {
case 3:
$__t2 = new Phpurs_Data0("Wednesday");
goto end_branch_2;;
break;
default:
;
break;
};
  switch ($n_2_0) {
case 4:
$__t2 = new Phpurs_Data0("Thursday");
goto end_branch_2;;
break;
default:
;
break;
};
  switch ($n_2_0) {
case 5:
$__t2 = new Phpurs_Data0("Friday");
goto end_branch_2;;
break;
default:
;
break;
};
  switch ($n_2_0) {
case 6:
$__t2 = new Phpurs_Data0("Saturday");
goto end_branch_2;;
break;
default:
;
break;
};
  switch ($n_2_0) {
case 7:
$__t2 = new Phpurs_Data0("Sunday");
goto end_branch_2;;
break;
default:
;
break;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()); return $v; };
\PhpursThunks::$thunks['Data_Date_showDate'] = function() { $v = (object)["show" => function($v_0 = null) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object(($v_0)->{'value1'}) && ((($v_0)->{'value1'})->{'tag'} === "January"))) {
$__t0 = "January";
goto end_branch_0;;
};
  if ((is_object(($v_0)->{'value1'}) && ((($v_0)->{'value1'})->{'tag'} === "February"))) {
$__t0 = "February";
goto end_branch_0;;
};
  if ((is_object(($v_0)->{'value1'}) && ((($v_0)->{'value1'})->{'tag'} === "March"))) {
$__t0 = "March";
goto end_branch_0;;
};
  if ((is_object(($v_0)->{'value1'}) && ((($v_0)->{'value1'})->{'tag'} === "April"))) {
$__t0 = "April";
goto end_branch_0;;
};
  if ((is_object(($v_0)->{'value1'}) && ((($v_0)->{'value1'})->{'tag'} === "May"))) {
$__t0 = "May";
goto end_branch_0;;
};
  if ((is_object(($v_0)->{'value1'}) && ((($v_0)->{'value1'})->{'tag'} === "June"))) {
$__t0 = "June";
goto end_branch_0;;
};
  if ((is_object(($v_0)->{'value1'}) && ((($v_0)->{'value1'})->{'tag'} === "July"))) {
$__t0 = "July";
goto end_branch_0;;
};
  if ((is_object(($v_0)->{'value1'}) && ((($v_0)->{'value1'})->{'tag'} === "August"))) {
$__t0 = "August";
goto end_branch_0;;
};
  if ((is_object(($v_0)->{'value1'}) && ((($v_0)->{'value1'})->{'tag'} === "September"))) {
$__t0 = "September";
goto end_branch_0;;
};
  if ((is_object(($v_0)->{'value1'}) && ((($v_0)->{'value1'})->{'tag'} === "October"))) {
$__t0 = "October";
goto end_branch_0;;
};
  if ((is_object(($v_0)->{'value1'}) && ((($v_0)->{'value1'})->{'tag'} === "November"))) {
$__t0 = "November";
goto end_branch_0;;
};
  if ((is_object(($v_0)->{'value1'}) && ((($v_0)->{'value1'})->{'tag'} === "December"))) {
$__t0 = "December";
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = ((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))("(Date "))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))("(Year "))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))((($GLOBALS['Data_Show_showIntImpl'] ?? \PhpursThunks::eval('Data_Show_showIntImpl')))(($v_0)->{'value0'})))(")"))))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))(" "))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))($__t0))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))(" "))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))("(Day "))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))((($GLOBALS['Data_Show_showIntImpl'] ?? \PhpursThunks::eval('Data_Show_showIntImpl')))(($v_0)->{'value2'})))(")"))))(")"))))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Date_month'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ($v_0)->{'value1'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Date_isLeapYear'] = function() { $v = function($y_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_HeytingAlgebra_boolConj'] ?? \PhpursThunks::eval('Data_HeytingAlgebra_boolConj')))(((($GLOBALS['Data_Eq_eqIntImpl'] ?? \PhpursThunks::eval('Data_Eq_eqIntImpl')))(((($GLOBALS['Data_EuclideanRing_intMod'] ?? \PhpursThunks::eval('Data_EuclideanRing_intMod')))($y_0))(4)))(0)))(((($GLOBALS['Data_HeytingAlgebra_boolDisj'] ?? \PhpursThunks::eval('Data_HeytingAlgebra_boolDisj')))(((($GLOBALS['Data_Eq_eqIntImpl'] ?? \PhpursThunks::eval('Data_Eq_eqIntImpl')))(((($GLOBALS['Data_EuclideanRing_intMod'] ?? \PhpursThunks::eval('Data_EuclideanRing_intMod')))($y_0))(400)))(0)))((($GLOBALS['Data_HeytingAlgebra_boolNot'] ?? \PhpursThunks::eval('Data_HeytingAlgebra_boolNot')))(((($GLOBALS['Data_Eq_eqIntImpl'] ?? \PhpursThunks::eval('Data_Eq_eqIntImpl')))(((($GLOBALS['Data_EuclideanRing_intMod'] ?? \PhpursThunks::eval('Data_EuclideanRing_intMod')))($y_0))(100)))(0))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Date_lastDayOfMonth'] = function() { $v = (function() {
  $__fn = function($y_0 = null, $m_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $unsafeDay_2_0 = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Partial_Unsafe__unsafePartial'] ?? \PhpursThunks::eval('Partial_Unsafe__unsafePartial')))(function($dollar__unused_2 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Date_fromJust'] ?? \PhpursThunks::eval('Data_Date_fromJust'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))((($GLOBALS['Data_Date_Component_boundedEnumDay'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumDay')))->{'toEnum'});
  $__t1 = null;;
  if ((is_object($m_1) && (($m_1)->{'tag'} === "January"))) {
$__t1 = ($unsafeDay_2_0)(31);
goto end_branch_1;;
};
  if ((is_object($m_1) && (($m_1)->{'tag'} === "February"))) {
$__t2 = null;;
if ((($GLOBALS['Data_Date_isLeapYear'] ?? \PhpursThunks::eval('Data_Date_isLeapYear')))($y_0)) {
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
})(); return $v; };
\PhpursThunks::$thunks['Data_Date_eqDate'] = function() { $v = (object)["eq" => (function() {
  $__fn = function($x_0 = null, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ((is_object(($x_0)->{'value1'}) && ((($x_0)->{'value1'})->{'tag'} === "January"))) {
$__t0 = (is_object(($y_1)->{'value1'}) && ((($y_1)->{'value1'})->{'tag'} === "January"));
goto end_branch_0;;
};
  if ((is_object(($x_0)->{'value1'}) && ((($x_0)->{'value1'})->{'tag'} === "February"))) {
$__t0 = (is_object(($y_1)->{'value1'}) && ((($y_1)->{'value1'})->{'tag'} === "February"));
goto end_branch_0;;
};
  if ((is_object(($x_0)->{'value1'}) && ((($x_0)->{'value1'})->{'tag'} === "March"))) {
$__t0 = (is_object(($y_1)->{'value1'}) && ((($y_1)->{'value1'})->{'tag'} === "March"));
goto end_branch_0;;
};
  if ((is_object(($x_0)->{'value1'}) && ((($x_0)->{'value1'})->{'tag'} === "April"))) {
$__t0 = (is_object(($y_1)->{'value1'}) && ((($y_1)->{'value1'})->{'tag'} === "April"));
goto end_branch_0;;
};
  if ((is_object(($x_0)->{'value1'}) && ((($x_0)->{'value1'})->{'tag'} === "May"))) {
$__t0 = (is_object(($y_1)->{'value1'}) && ((($y_1)->{'value1'})->{'tag'} === "May"));
goto end_branch_0;;
};
  if ((is_object(($x_0)->{'value1'}) && ((($x_0)->{'value1'})->{'tag'} === "June"))) {
$__t0 = (is_object(($y_1)->{'value1'}) && ((($y_1)->{'value1'})->{'tag'} === "June"));
goto end_branch_0;;
};
  if ((is_object(($x_0)->{'value1'}) && ((($x_0)->{'value1'})->{'tag'} === "July"))) {
$__t0 = (is_object(($y_1)->{'value1'}) && ((($y_1)->{'value1'})->{'tag'} === "July"));
goto end_branch_0;;
};
  if ((is_object(($x_0)->{'value1'}) && ((($x_0)->{'value1'})->{'tag'} === "August"))) {
$__t0 = (is_object(($y_1)->{'value1'}) && ((($y_1)->{'value1'})->{'tag'} === "August"));
goto end_branch_0;;
};
  if ((is_object(($x_0)->{'value1'}) && ((($x_0)->{'value1'})->{'tag'} === "September"))) {
$__t0 = (is_object(($y_1)->{'value1'}) && ((($y_1)->{'value1'})->{'tag'} === "September"));
goto end_branch_0;;
};
  if ((is_object(($x_0)->{'value1'}) && ((($x_0)->{'value1'})->{'tag'} === "October"))) {
$__t0 = (is_object(($y_1)->{'value1'}) && ((($y_1)->{'value1'})->{'tag'} === "October"));
goto end_branch_0;;
};
  if ((is_object(($x_0)->{'value1'}) && ((($x_0)->{'value1'})->{'tag'} === "November"))) {
$__t0 = (is_object(($y_1)->{'value1'}) && ((($y_1)->{'value1'})->{'tag'} === "November"));
goto end_branch_0;;
};
  $__t0 = ((is_object(($x_0)->{'value1'}) && ((($x_0)->{'value1'})->{'tag'} === "December")) && (is_object(($y_1)->{'value1'}) && ((($y_1)->{'value1'})->{'tag'} === "December")));
  end_branch_0:;
  $__res = ((($GLOBALS['Data_HeytingAlgebra_boolConj'] ?? \PhpursThunks::eval('Data_HeytingAlgebra_boolConj')))(((($GLOBALS['Data_HeytingAlgebra_boolConj'] ?? \PhpursThunks::eval('Data_HeytingAlgebra_boolConj')))(((($GLOBALS['Data_Eq_eqIntImpl'] ?? \PhpursThunks::eval('Data_Eq_eqIntImpl')))(($x_0)->{'value0'}))(($y_1)->{'value0'})))($__t0)))(((($GLOBALS['Data_Eq_eqIntImpl'] ?? \PhpursThunks::eval('Data_Eq_eqIntImpl')))(($x_0)->{'value2'}))(($y_1)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Data_Date_ordDate'] = function() { $v = (object)["compare" => (function() {
  $__fn = function($x_0 = null, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v_2_0 = (((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->{'compare'})(($x_0)->{'value0'}))(($y_1)->{'value0'});
  $__t3 = null;;
  if ((is_object($v_2_0) && (($v_2_0)->{'tag'} === "LT"))) {
$__t3 = new Phpurs_Data0("LT");
goto end_branch_3;;
};
  if ((is_object($v_2_0) && (($v_2_0)->{'tag'} === "GT"))) {
$__t3 = new Phpurs_Data0("GT");
goto end_branch_3;;
};
  $v1_3_1 = (((($GLOBALS['Data_Date_Component_ordMonth'] ?? \PhpursThunks::eval('Data_Date_Component_ordMonth')))->{'compare'})(($x_0)->{'value1'}))(($y_1)->{'value1'});
  $__t2 = null;;
  if ((is_object($v1_3_1) && (($v1_3_1)->{'tag'} === "LT"))) {
$__t2 = new Phpurs_Data0("LT");
goto end_branch_2;;
};
  if ((is_object($v1_3_1) && (($v1_3_1)->{'tag'} === "GT"))) {
$__t2 = new Phpurs_Data0("GT");
goto end_branch_2;;
};
  $__t2 = (((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->{'compare'})(($x_0)->{'value2'}))(($y_1)->{'value2'});
  end_branch_2:;
  $__t3 = $__t2;
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Date_eqDate'] ?? \PhpursThunks::eval('Data_Date_eqDate'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Date_enumDate'] = function() { $v = (object)["succ" => function($v_0 = null) {
  $__num = \func_num_args();
  $sm_1_0 = ((($GLOBALS['Data_Date_Component_enumMonth'] ?? \PhpursThunks::eval('Data_Date_Component_enumMonth')))->{'succ'})(($v_0)->{'value1'});
  $v1_2_1 = ((($GLOBALS['Data_Date_Component_enumDay'] ?? \PhpursThunks::eval('Data_Date_Component_enumDay')))->{'succ'})(($v_0)->{'value2'});
  $__t2 = null;;
  if (((($GLOBALS['Data_Date_greaterThan'] ?? \PhpursThunks::eval('Data_Date_greaterThan')))($v1_2_1))(new Phpurs_Data1("Just", ((($GLOBALS['Data_Date_lastDayOfMonth'] ?? \PhpursThunks::eval('Data_Date_lastDayOfMonth')))(($v_0)->{'value0'}))(($v_0)->{'value1'})))) {
$__t2 = new Phpurs_Data0("Nothing");
goto end_branch_2;;
};
  $__t2 = $v1_2_1;
  end_branch_2:;
  $sd_3_2 = $__t2;
  $__t4 = null;;
  if ((function() use ($sd_3_2, $sm_1_0, &$__fn) {
$__t5 = null;;
if ((is_object($sd_3_2) && (($sd_3_2)->{'tag'} === "Nothing"))) {
$__t5 = true;
goto end_branch_5;;
};
if ((is_object($sd_3_2) && (($sd_3_2)->{'tag'} === "Just"))) {
$__t5 = false;
goto end_branch_5;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t5 = null;
end_branch_5:;
$__t6 = null;;
if ((is_object($sm_1_0) && (($sm_1_0)->{'tag'} === "Nothing"))) {
$__t6 = true;
goto end_branch_6;;
};
if ((is_object($sm_1_0) && (($sm_1_0)->{'tag'} === "Just"))) {
$__t6 = false;
goto end_branch_6;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t6 = null;
end_branch_6:;
return ((($GLOBALS['Data_HeytingAlgebra_boolConj'] ?? \PhpursThunks::eval('Data_HeytingAlgebra_boolConj')))($__t5))($__t6);
})()) {
$__t4 = ((($GLOBALS['Data_Date_Component_enumYear'] ?? \PhpursThunks::eval('Data_Date_Component_enumYear')))->{'succ'})(($v_0)->{'value0'});
goto end_branch_4;;
};
  $__t4 = new Phpurs_Data1("Just", ($v_0)->{'value0'});
  end_branch_4:;
  $__local_var_4_4 = $__t4;
  $__t8 = null;;
  if ((is_object($__local_var_4_4) && (($__local_var_4_4)->{'tag'} === "Just"))) {
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
$__t11 = ($sm_1_0)->{'value0'};
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
$__local_var_5_9 = ((($GLOBALS['Data_Date_Date'] ?? \PhpursThunks::eval('Data_Date_Date')))(($__local_var_4_4)->{'value0'}))($__t9);
$__t13 = null;;
if ((function() use ($sd_3_2, &$__fn) {
$__t14 = null;;
if ((is_object($sd_3_2) && (($sd_3_2)->{'tag'} === "Nothing"))) {
$__t14 = true;
goto end_branch_14;;
};
if ((is_object($sd_3_2) && (($sd_3_2)->{'tag'} === "Just"))) {
$__t14 = false;
goto end_branch_14;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t14 = null;
end_branch_14:;
return $__t14;
})()) {
$__t13 = ((($GLOBALS['Data_Date_Component_boundedEnumDay'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumDay')))->{'toEnum'})(1);
goto end_branch_13;;
};
$__t13 = $sd_3_2;
end_branch_13:;
$__local_var_6_13 = $__t13;
$__t16 = null;;
if ((is_object($__local_var_6_13) && (($__local_var_6_13)->{'tag'} === "Just"))) {
$__t16 = new Phpurs_Data1("Just", ($__local_var_5_9)(($__local_var_6_13)->{'value0'}));
goto end_branch_16;;
};
$__t16 = new Phpurs_Data0("Nothing");
end_branch_16:;
$__t8 = $__t16;
goto end_branch_8;;
};
  $__t8 = new Phpurs_Data0("Nothing");
  end_branch_8:;
  $__res = $__t8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "pred" => function($v_0 = null) {
  $__num = \func_num_args();
  $pm_1_17 = ((($GLOBALS['Data_Date_Component_enumMonth'] ?? \PhpursThunks::eval('Data_Date_Component_enumMonth')))->{'pred'})(($v_0)->{'value1'});
  $pd_2_18 = ((($GLOBALS['Data_Date_Component_enumDay'] ?? \PhpursThunks::eval('Data_Date_Component_enumDay')))->{'pred'})(($v_0)->{'value2'});
  $__t19 = null;;
  if ((function() use ($pd_2_18, &$__fn) {
$__t20 = null;;
if ((is_object($pd_2_18) && (($pd_2_18)->{'tag'} === "Nothing"))) {
$__t20 = true;
goto end_branch_20;;
};
if ((is_object($pd_2_18) && (($pd_2_18)->{'tag'} === "Just"))) {
$__t20 = false;
goto end_branch_20;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t20 = null;
end_branch_20:;
return $__t20;
})()) {
$__t21 = null;;
if ((is_object($pm_1_17) && (($pm_1_17)->{'tag'} === "Nothing"))) {
$__t21 = new Phpurs_Data0("December");
goto end_branch_21;;
};
if ((is_object($pm_1_17) && (($pm_1_17)->{'tag'} === "Just"))) {
$__t21 = ($pm_1_17)->{'value0'};
goto end_branch_21;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t21 = null;
end_branch_21:;
$__t19 = $__t21;
goto end_branch_19;;
};
  $__t19 = ($v_0)->{'value1'};
  end_branch_19:;
  $m__prime___3_19 = $__t19;
  $l_4_23 = ((($GLOBALS['Data_Date_lastDayOfMonth'] ?? \PhpursThunks::eval('Data_Date_lastDayOfMonth')))(($v_0)->{'value0'}))($m__prime___3_19);
  $__t24 = null;;
  if ((function() use ($pd_2_18, $pm_1_17, &$__fn) {
$__t25 = null;;
if ((is_object($pd_2_18) && (($pd_2_18)->{'tag'} === "Nothing"))) {
$__t25 = true;
goto end_branch_25;;
};
if ((is_object($pd_2_18) && (($pd_2_18)->{'tag'} === "Just"))) {
$__t25 = false;
goto end_branch_25;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t25 = null;
end_branch_25:;
$__t26 = null;;
if ((is_object($pm_1_17) && (($pm_1_17)->{'tag'} === "Nothing"))) {
$__t26 = true;
goto end_branch_26;;
};
if ((is_object($pm_1_17) && (($pm_1_17)->{'tag'} === "Just"))) {
$__t26 = false;
goto end_branch_26;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t26 = null;
end_branch_26:;
return ((($GLOBALS['Data_HeytingAlgebra_boolConj'] ?? \PhpursThunks::eval('Data_HeytingAlgebra_boolConj')))($__t25))($__t26);
})()) {
$__t24 = ((($GLOBALS['Data_Date_Component_enumYear'] ?? \PhpursThunks::eval('Data_Date_Component_enumYear')))->{'pred'})(($v_0)->{'value0'});
goto end_branch_24;;
};
  $__t24 = new Phpurs_Data1("Just", ($v_0)->{'value0'});
  end_branch_24:;
  $__local_var_5_24 = $__t24;
  $__t28 = null;;
  if ((is_object($__local_var_5_24) && (($__local_var_5_24)->{'tag'} === "Just"))) {
$__local_var_6_29 = ((($GLOBALS['Data_Date_Date'] ?? \PhpursThunks::eval('Data_Date_Date')))(($__local_var_5_24)->{'value0'}))($m__prime___3_19);
$__t30 = null;;
if ((function() use ($pd_2_18, &$__fn) {
$__t31 = null;;
if ((is_object($pd_2_18) && (($pd_2_18)->{'tag'} === "Nothing"))) {
$__t31 = true;
goto end_branch_31;;
};
if ((is_object($pd_2_18) && (($pd_2_18)->{'tag'} === "Just"))) {
$__t31 = false;
goto end_branch_31;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t31 = null;
end_branch_31:;
return $__t31;
})()) {
$__t30 = new Phpurs_Data1("Just", ($__local_var_6_29)($l_4_23));
goto end_branch_30;;
};
if ((is_object($pd_2_18) && (($pd_2_18)->{'tag'} === "Just"))) {
$__t30 = new Phpurs_Data1("Just", ($__local_var_6_29)(($pd_2_18)->{'value0'}));
goto end_branch_30;;
};
$__t30 = new Phpurs_Data0("Nothing");
end_branch_30:;
$__t28 = $__t30;
goto end_branch_28;;
};
  $__t28 = new Phpurs_Data0("Nothing");
  end_branch_28:;
  $__res = $__t28;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Ord0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Date_ordDate'] ?? \PhpursThunks::eval('Data_Date_ordDate'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Date_diff'] = function() { $v = (function() {
  $__fn = function($dictDuration_0 = null, $v_1 = null, $v1_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if ((is_object(($v_1)->{'value1'}) && ((($v_1)->{'value1'})->{'tag'} === "January"))) {
$__t0 = 1;
goto end_branch_0;;
};
  if ((is_object(($v_1)->{'value1'}) && ((($v_1)->{'value1'})->{'tag'} === "February"))) {
$__t0 = 2;
goto end_branch_0;;
};
  if ((is_object(($v_1)->{'value1'}) && ((($v_1)->{'value1'})->{'tag'} === "March"))) {
$__t0 = 3;
goto end_branch_0;;
};
  if ((is_object(($v_1)->{'value1'}) && ((($v_1)->{'value1'})->{'tag'} === "April"))) {
$__t0 = 4;
goto end_branch_0;;
};
  if ((is_object(($v_1)->{'value1'}) && ((($v_1)->{'value1'})->{'tag'} === "May"))) {
$__t0 = 5;
goto end_branch_0;;
};
  if ((is_object(($v_1)->{'value1'}) && ((($v_1)->{'value1'})->{'tag'} === "June"))) {
$__t0 = 6;
goto end_branch_0;;
};
  if ((is_object(($v_1)->{'value1'}) && ((($v_1)->{'value1'})->{'tag'} === "July"))) {
$__t0 = 7;
goto end_branch_0;;
};
  if ((is_object(($v_1)->{'value1'}) && ((($v_1)->{'value1'})->{'tag'} === "August"))) {
$__t0 = 8;
goto end_branch_0;;
};
  if ((is_object(($v_1)->{'value1'}) && ((($v_1)->{'value1'})->{'tag'} === "September"))) {
$__t0 = 9;
goto end_branch_0;;
};
  if ((is_object(($v_1)->{'value1'}) && ((($v_1)->{'value1'})->{'tag'} === "October"))) {
$__t0 = 10;
goto end_branch_0;;
};
  if ((is_object(($v_1)->{'value1'}) && ((($v_1)->{'value1'})->{'tag'} === "November"))) {
$__t0 = 11;
goto end_branch_0;;
};
  if ((is_object(($v_1)->{'value1'}) && ((($v_1)->{'value1'})->{'tag'} === "December"))) {
$__t0 = 12;
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__t1 = null;;
  if ((is_object(($v1_2)->{'value1'}) && ((($v1_2)->{'value1'})->{'tag'} === "January"))) {
$__t1 = 1;
goto end_branch_1;;
};
  if ((is_object(($v1_2)->{'value1'}) && ((($v1_2)->{'value1'})->{'tag'} === "February"))) {
$__t1 = 2;
goto end_branch_1;;
};
  if ((is_object(($v1_2)->{'value1'}) && ((($v1_2)->{'value1'})->{'tag'} === "March"))) {
$__t1 = 3;
goto end_branch_1;;
};
  if ((is_object(($v1_2)->{'value1'}) && ((($v1_2)->{'value1'})->{'tag'} === "April"))) {
$__t1 = 4;
goto end_branch_1;;
};
  if ((is_object(($v1_2)->{'value1'}) && ((($v1_2)->{'value1'})->{'tag'} === "May"))) {
$__t1 = 5;
goto end_branch_1;;
};
  if ((is_object(($v1_2)->{'value1'}) && ((($v1_2)->{'value1'})->{'tag'} === "June"))) {
$__t1 = 6;
goto end_branch_1;;
};
  if ((is_object(($v1_2)->{'value1'}) && ((($v1_2)->{'value1'})->{'tag'} === "July"))) {
$__t1 = 7;
goto end_branch_1;;
};
  if ((is_object(($v1_2)->{'value1'}) && ((($v1_2)->{'value1'})->{'tag'} === "August"))) {
$__t1 = 8;
goto end_branch_1;;
};
  if ((is_object(($v1_2)->{'value1'}) && ((($v1_2)->{'value1'})->{'tag'} === "September"))) {
$__t1 = 9;
goto end_branch_1;;
};
  if ((is_object(($v1_2)->{'value1'}) && ((($v1_2)->{'value1'})->{'tag'} === "October"))) {
$__t1 = 10;
goto end_branch_1;;
};
  if ((is_object(($v1_2)->{'value1'}) && ((($v1_2)->{'value1'})->{'tag'} === "November"))) {
$__t1 = 11;
goto end_branch_1;;
};
  if ((is_object(($v1_2)->{'value1'}) && ((($v1_2)->{'value1'})->{'tag'} === "December"))) {
$__t1 = 12;
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = (($dictDuration_0)->{'toDuration'})((((((((($GLOBALS['Data_Function_Uncurried_runFn6'] ?? \PhpursThunks::eval('Data_Function_Uncurried_runFn6')))(($GLOBALS['Data_Date_calcDiff'] ?? \PhpursThunks::eval('Data_Date_calcDiff'))))(($v_1)->{'value0'}))($__t0))(($v_1)->{'value2'}))(($v1_2)->{'value0'}))($__t1))(($v1_2)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Date_day'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ($v_0)->{'value2'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Date_canonicalDate'] = function() { $v = (function() {
  $__fn = function($y_0 = null, $m_1 = null, $d_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t1 = null;;
  if ((is_object($m_1) && (($m_1)->{'tag'} === "January"))) {
$__t1 = 1;
goto end_branch_1;;
};
  if ((is_object($m_1) && (($m_1)->{'tag'} === "February"))) {
$__t1 = 2;
goto end_branch_1;;
};
  if ((is_object($m_1) && (($m_1)->{'tag'} === "March"))) {
$__t1 = 3;
goto end_branch_1;;
};
  if ((is_object($m_1) && (($m_1)->{'tag'} === "April"))) {
$__t1 = 4;
goto end_branch_1;;
};
  if ((is_object($m_1) && (($m_1)->{'tag'} === "May"))) {
$__t1 = 5;
goto end_branch_1;;
};
  if ((is_object($m_1) && (($m_1)->{'tag'} === "June"))) {
$__t1 = 6;
goto end_branch_1;;
};
  if ((is_object($m_1) && (($m_1)->{'tag'} === "July"))) {
$__t1 = 7;
goto end_branch_1;;
};
  if ((is_object($m_1) && (($m_1)->{'tag'} === "August"))) {
$__t1 = 8;
goto end_branch_1;;
};
  if ((is_object($m_1) && (($m_1)->{'tag'} === "September"))) {
$__t1 = 9;
goto end_branch_1;;
};
  if ((is_object($m_1) && (($m_1)->{'tag'} === "October"))) {
$__t1 = 10;
goto end_branch_1;;
};
  if ((is_object($m_1) && (($m_1)->{'tag'} === "November"))) {
$__t1 = 11;
goto end_branch_1;;
};
  if ((is_object($m_1) && (($m_1)->{'tag'} === "December"))) {
$__t1 = 12;
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = (((((($GLOBALS['Data_Function_Uncurried_runFn4'] ?? \PhpursThunks::eval('Data_Function_Uncurried_runFn4')))(($GLOBALS['Data_Date_canonicalDateImpl'] ?? \PhpursThunks::eval('Data_Date_canonicalDateImpl'))))((($GLOBALS['Partial_Unsafe__unsafePartial'] ?? \PhpursThunks::eval('Partial_Unsafe__unsafePartial')))((function() {
  $__fn = function($dollar__unused_3 = null, $y__prime___4 = null, $m__prime___5 = null, $d__prime___6 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__t0 = null;;
  switch ($m__prime___5) {
case 1:
$__t0 = new Phpurs_Data0("January");
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($m__prime___5) {
case 2:
$__t0 = new Phpurs_Data0("February");
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($m__prime___5) {
case 3:
$__t0 = new Phpurs_Data0("March");
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($m__prime___5) {
case 4:
$__t0 = new Phpurs_Data0("April");
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($m__prime___5) {
case 5:
$__t0 = new Phpurs_Data0("May");
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($m__prime___5) {
case 6:
$__t0 = new Phpurs_Data0("June");
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($m__prime___5) {
case 7:
$__t0 = new Phpurs_Data0("July");
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($m__prime___5) {
case 8:
$__t0 = new Phpurs_Data0("August");
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($m__prime___5) {
case 9:
$__t0 = new Phpurs_Data0("September");
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($m__prime___5) {
case 10:
$__t0 = new Phpurs_Data0("October");
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($m__prime___5) {
case 11:
$__t0 = new Phpurs_Data0("November");
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($m__prime___5) {
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
  $__res = new Phpurs_Data3("Date", $y__prime___4, $__t0, $d__prime___6);
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})())))($y_0))($__t1))($d_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Date_exactDate'] = function() { $v = (function() {
  $__fn = function($y_0 = null, $m_1 = null, $d_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if ((((($GLOBALS['Data_Date_eqDate'] ?? \PhpursThunks::eval('Data_Date_eqDate')))->{'eq'})((((($GLOBALS['Data_Date_canonicalDate'] ?? \PhpursThunks::eval('Data_Date_canonicalDate')))($y_0))($m_1))($d_2)))(new Phpurs_Data3("Date", $y_0, $m_1, $d_2))) {
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
})(); return $v; };
\PhpursThunks::$thunks['Data_Date_boundedDate'] = function() { $v = (object)["bottom" => new Phpurs_Data3("Date", (($GLOBALS['Data_Date_Component_boundedYear'] ?? \PhpursThunks::eval('Data_Date_Component_boundedYear')))->{'bottom'}, new Phpurs_Data0("January"), 1), "top" => new Phpurs_Data3("Date", 275759, new Phpurs_Data0("December"), 31), "Ord0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Date_ordDate'] ?? \PhpursThunks::eval('Data_Date_ordDate'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Date_adjust'] = function() { $v = (function() {
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
  $__t19 = null;;
  switch ($v1_3) {
case 0:
$__t19 = new Phpurs_Data1("Just", $v2_4);
goto end_branch_19;;
break;
default:
;
break;
};
  $j_5_1 = ((($GLOBALS['Data_Semiring_intAdd'] ?? \PhpursThunks::eval('Data_Semiring_intAdd')))($v1_3))(($v2_4)->{'value2'});
  $low_6_2 = (is_object((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->{'compare'})($j_5_1))(1)) && (((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->{'compare'})($j_5_1))(1))->{'tag'} === "LT"));
  $__t3 = null;;
  if ($low_6_2) {
$__local_var_7_4 = ((($GLOBALS['Data_Date_Component_enumMonth'] ?? \PhpursThunks::eval('Data_Date_Component_enumMonth')))->{'pred'})(($v2_4)->{'value1'});
$__t5 = null;;
if ((is_object($__local_var_7_4) && (($__local_var_7_4)->{'tag'} === "Nothing"))) {
$__t5 = new Phpurs_Data0("December");
goto end_branch_5;;
};
if ((is_object($__local_var_7_4) && (($__local_var_7_4)->{'tag'} === "Just"))) {
$__t5 = ($__local_var_7_4)->{'value0'};
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
  $l_7_3 = ((($GLOBALS['Data_Date_lastDayOfMonth'] ?? \PhpursThunks::eval('Data_Date_lastDayOfMonth')))(($v2_4)->{'value0'}))($__t3);
  $hi_8_7 = (is_object((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->{'compare'})($j_5_1))($l_7_3)) && (((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->{'compare'})($j_5_1))($l_7_3))->{'tag'} === "GT"));
  $__t8 = null;;
  if ($low_6_2) {
$__t8 = $j_5_1;
goto end_branch_8;;
};
  if ($hi_8_7) {
$__t8 = ((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))(((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))($j_5_1))($l_7_3)))(1);
goto end_branch_8;;
};
  $__t8 = 0;
  end_branch_8:;
  $__local_var_9_8 = ($adj_2_0)($__t8);
  $__t13 = null;;
  if ($low_6_2) {
$__local_var_10_14 = ((($GLOBALS['Data_Date_Date'] ?? \PhpursThunks::eval('Data_Date_Date')))(($v2_4)->{'value0'}))(($v2_4)->{'value1'});
$__local_var_11_15 = ((($GLOBALS['Data_Date_Component_boundedEnumDay'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumDay')))->{'toEnum'})(1);
$__t16 = null;;
if ((is_object($__local_var_11_15) && (($__local_var_11_15)->{'tag'} === "Just"))) {
$__t16 = ((($GLOBALS['Data_Date_enumDate'] ?? \PhpursThunks::eval('Data_Date_enumDate')))->{'pred'})(($__local_var_10_14)(($__local_var_11_15)->{'value0'}));
goto end_branch_16;;
};
$__t16 = new Phpurs_Data0("Nothing");
end_branch_16:;
$__t13 = $__t16;
goto end_branch_13;;
};
  if ($hi_8_7) {
$__t13 = ((($GLOBALS['Data_Date_enumDate'] ?? \PhpursThunks::eval('Data_Date_enumDate')))->{'succ'})(new Phpurs_Data3("Date", ($v2_4)->{'value0'}, ($v2_4)->{'value1'}, $l_7_3));
goto end_branch_13;;
};
  $__local_var_10_10 = ((($GLOBALS['Data_Date_Date'] ?? \PhpursThunks::eval('Data_Date_Date')))(($v2_4)->{'value0'}))(($v2_4)->{'value1'});
  $__local_var_11_11 = ((($GLOBALS['Data_Date_Component_boundedEnumDay'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumDay')))->{'toEnum'})($j_5_1);
  $__t12 = null;;
  if ((is_object($__local_var_11_11) && (($__local_var_11_11)->{'tag'} === "Just"))) {
$__t12 = new Phpurs_Data1("Just", ($__local_var_10_10)(($__local_var_11_11)->{'value0'}));
goto end_branch_12;;
};
  $__t12 = new Phpurs_Data0("Nothing");
  end_branch_12:;
  $__t13 = $__t12;
  end_branch_13:;
  $__local_var_10_10 = $__t13;
  $__t18 = null;;
  if ((is_object($__local_var_10_10) && (($__local_var_10_10)->{'tag'} === "Just"))) {
$__t18 = ($__local_var_9_8)(($__local_var_10_10)->{'value0'});
goto end_branch_18;;
};
  if ((is_object($__local_var_10_10) && (($__local_var_10_10)->{'tag'} === "Nothing"))) {
$__t18 = new Phpurs_Data0("Nothing");
goto end_branch_18;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t18 = null;
  end_branch_18:;
  $__t19 = $__t18;
  end_branch_19:;
  $__res = $__t19;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__local_var_3_20 = (($GLOBALS['Data_Int_fromNumber'] ?? \PhpursThunks::eval('Data_Int_fromNumber')))($v_0);
  $__t21 = null;;
  if ((is_object($__local_var_3_20) && (($__local_var_3_20)->{'tag'} === "Just"))) {
$__t21 = (($adj_2_0)(($__local_var_3_20)->{'value0'}))($date_1);
goto end_branch_21;;
};
  if ((is_object($__local_var_3_20) && (($__local_var_3_20)->{'tag'} === "Nothing"))) {
$__t21 = new Phpurs_Data0("Nothing");
goto end_branch_21;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t21 = null;
  end_branch_21:;
  $__res = $__t21;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
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
\PhpursThunks::$thunks['Data_Date_calcDiff'] = function() use (&$ffi_Data_Date) { return $ffi_Data_Date['calcDiff']; };
\PhpursThunks::$thunks['Data_Date_calcWeekday'] = function() use (&$ffi_Data_Date) { return $ffi_Data_Date['calcWeekday']; };
\PhpursThunks::$thunks['Data_Date_canonicalDateImpl'] = function() use (&$ffi_Data_Date) { return $ffi_Data_Date['canonicalDateImpl']; };




















