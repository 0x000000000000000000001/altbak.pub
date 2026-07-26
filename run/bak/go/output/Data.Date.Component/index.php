<?php

namespace Data\Date\Component;

// ALL IMPORTS: Control.Semigroupoid, Data.Boolean, Data.Bounded, Data.Date.Component, Data.Enum, Data.Eq, Data.HeytingAlgebra, Data.Maybe, Data.Ord, Data.Ordering, Data.Ring, Data.Semigroup, Data.Semiring, Data.Show, Prelude, Prim
// TO REQUIRE: Control.Semigroupoid, Data.Boolean, Data.Bounded, Data.Date.Component, Data.Enum, Data.Eq, Data.HeytingAlgebra, Data.Maybe, Data.Ord, Data.Ordering, Data.Ring, Data.Semigroup, Data.Semiring, Data.Show, Prelude
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Boolean/index.php';
require_once __DIR__ . '/../Data.Bounded/index.php';
require_once __DIR__ . '/../Data.Date.Component/index.php';
require_once __DIR__ . '/../Data.Enum/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.HeytingAlgebra/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ordering/index.php';
require_once __DIR__ . '/../Data.Ring/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Semiring/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
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
\PhpursThunks::$thunks['Data_Date_Component_Monday'] = function() { $v = ($GLOBALS['__phpurs_data0_Monday'] ??= new Phpurs_Data0("Monday")); return $v; };
\PhpursThunks::$thunks['Data_Date_Component_Tuesday'] = function() { $v = ($GLOBALS['__phpurs_data0_Tuesday'] ??= new Phpurs_Data0("Tuesday")); return $v; };
\PhpursThunks::$thunks['Data_Date_Component_Wednesday'] = function() { $v = ($GLOBALS['__phpurs_data0_Wednesday'] ??= new Phpurs_Data0("Wednesday")); return $v; };
\PhpursThunks::$thunks['Data_Date_Component_Thursday'] = function() { $v = ($GLOBALS['__phpurs_data0_Thursday'] ??= new Phpurs_Data0("Thursday")); return $v; };
\PhpursThunks::$thunks['Data_Date_Component_Friday'] = function() { $v = ($GLOBALS['__phpurs_data0_Friday'] ??= new Phpurs_Data0("Friday")); return $v; };
\PhpursThunks::$thunks['Data_Date_Component_Saturday'] = function() { $v = ($GLOBALS['__phpurs_data0_Saturday'] ??= new Phpurs_Data0("Saturday")); return $v; };
\PhpursThunks::$thunks['Data_Date_Component_Sunday'] = function() { $v = ($GLOBALS['__phpurs_data0_Sunday'] ??= new Phpurs_Data0("Sunday")); return $v; };
\PhpursThunks::$thunks['Data_Date_Component_January'] = function() { $v = ($GLOBALS['__phpurs_data0_January'] ??= new Phpurs_Data0("January")); return $v; };
\PhpursThunks::$thunks['Data_Date_Component_February'] = function() { $v = ($GLOBALS['__phpurs_data0_February'] ??= new Phpurs_Data0("February")); return $v; };
\PhpursThunks::$thunks['Data_Date_Component_March'] = function() { $v = ($GLOBALS['__phpurs_data0_March'] ??= new Phpurs_Data0("March")); return $v; };
\PhpursThunks::$thunks['Data_Date_Component_April'] = function() { $v = ($GLOBALS['__phpurs_data0_April'] ??= new Phpurs_Data0("April")); return $v; };
\PhpursThunks::$thunks['Data_Date_Component_May'] = function() { $v = ($GLOBALS['__phpurs_data0_May'] ??= new Phpurs_Data0("May")); return $v; };
\PhpursThunks::$thunks['Data_Date_Component_June'] = function() { $v = ($GLOBALS['__phpurs_data0_June'] ??= new Phpurs_Data0("June")); return $v; };
\PhpursThunks::$thunks['Data_Date_Component_July'] = function() { $v = ($GLOBALS['__phpurs_data0_July'] ??= new Phpurs_Data0("July")); return $v; };
\PhpursThunks::$thunks['Data_Date_Component_August'] = function() { $v = ($GLOBALS['__phpurs_data0_August'] ??= new Phpurs_Data0("August")); return $v; };
\PhpursThunks::$thunks['Data_Date_Component_September'] = function() { $v = ($GLOBALS['__phpurs_data0_September'] ??= new Phpurs_Data0("September")); return $v; };
\PhpursThunks::$thunks['Data_Date_Component_October'] = function() { $v = ($GLOBALS['__phpurs_data0_October'] ??= new Phpurs_Data0("October")); return $v; };
\PhpursThunks::$thunks['Data_Date_Component_November'] = function() { $v = ($GLOBALS['__phpurs_data0_November'] ??= new Phpurs_Data0("November")); return $v; };
\PhpursThunks::$thunks['Data_Date_Component_December'] = function() { $v = ($GLOBALS['__phpurs_data0_December'] ??= new Phpurs_Data0("December")); return $v; };
\PhpursThunks::$thunks['Data_Date_Component_showYear'] = function() { $v = ["show" => function($v_0 = null) {
  $__num = \func_num_args();
  $__res = (("(Year " . (($GLOBALS['Data_Show_showIntImpl'] ?? \PhpursThunks::eval('Data_Show_showIntImpl')))($v_0)) . ")");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Date_Component_showWeekday'] = function() { $v = ["show" => function($v_0 = null) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Monday"))) {
$__t0 = "Monday";
goto end_branch_0;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Tuesday"))) {
$__t0 = "Tuesday";
goto end_branch_0;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Wednesday"))) {
$__t0 = "Wednesday";
goto end_branch_0;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Thursday"))) {
$__t0 = "Thursday";
goto end_branch_0;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Friday"))) {
$__t0 = "Friday";
goto end_branch_0;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Saturday"))) {
$__t0 = "Saturday";
goto end_branch_0;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Sunday"))) {
$__t0 = "Sunday";
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Date_Component_showMonth'] = function() { $v = ["show" => function($v_0 = null) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object($v_0) && (($v_0)->{'tag'} === "January"))) {
$__t0 = "January";
goto end_branch_0;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "February"))) {
$__t0 = "February";
goto end_branch_0;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "March"))) {
$__t0 = "March";
goto end_branch_0;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "April"))) {
$__t0 = "April";
goto end_branch_0;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "May"))) {
$__t0 = "May";
goto end_branch_0;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "June"))) {
$__t0 = "June";
goto end_branch_0;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "July"))) {
$__t0 = "July";
goto end_branch_0;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "August"))) {
$__t0 = "August";
goto end_branch_0;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "September"))) {
$__t0 = "September";
goto end_branch_0;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "October"))) {
$__t0 = "October";
goto end_branch_0;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "November"))) {
$__t0 = "November";
goto end_branch_0;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "December"))) {
$__t0 = "December";
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Date_Component_showDay'] = function() { $v = ["show" => function($v_0 = null) {
  $__num = \func_num_args();
  $__res = (("(Day " . (($GLOBALS['Data_Show_showIntImpl'] ?? \PhpursThunks::eval('Data_Show_showIntImpl')))($v_0)) . ")");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Date_Component_ordYear'] = function() { $v = ($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')); return $v; };
\PhpursThunks::$thunks['Data_Date_Component_ordDay'] = function() { $v = ($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')); return $v; };
\PhpursThunks::$thunks['Data_Date_Component_eqYear'] = function() { $v = ($GLOBALS['Data_Eq_eqInt'] ?? \PhpursThunks::eval('Data_Eq_eqInt')); return $v; };
\PhpursThunks::$thunks['Data_Date_Component_eqWeekday'] = function() { $v = ["eq" => (function() {
  $__fn = function($x_0 = null, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ((is_object($x_0) && (($x_0)->{'tag'} === "Monday"))) {
$__t0 = (is_object($y_1) && (($y_1)->{'tag'} === "Monday"));
goto end_branch_0;;
};
  if ((is_object($x_0) && (($x_0)->{'tag'} === "Tuesday"))) {
$__t0 = (is_object($y_1) && (($y_1)->{'tag'} === "Tuesday"));
goto end_branch_0;;
};
  if ((is_object($x_0) && (($x_0)->{'tag'} === "Wednesday"))) {
$__t0 = (is_object($y_1) && (($y_1)->{'tag'} === "Wednesday"));
goto end_branch_0;;
};
  if ((is_object($x_0) && (($x_0)->{'tag'} === "Thursday"))) {
$__t0 = (is_object($y_1) && (($y_1)->{'tag'} === "Thursday"));
goto end_branch_0;;
};
  if ((is_object($x_0) && (($x_0)->{'tag'} === "Friday"))) {
$__t0 = (is_object($y_1) && (($y_1)->{'tag'} === "Friday"));
goto end_branch_0;;
};
  if ((is_object($x_0) && (($x_0)->{'tag'} === "Saturday"))) {
$__t0 = (is_object($y_1) && (($y_1)->{'tag'} === "Saturday"));
goto end_branch_0;;
};
  $__t0 = ((is_object($x_0) && (($x_0)->{'tag'} === "Sunday")) && (is_object($y_1) && (($y_1)->{'tag'} === "Sunday")));
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Data_Date_Component_ordWeekday'] = function() { $v = ["compare" => (function() {
  $__fn = function($x_0 = null, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ((is_object($x_0) && (($x_0)->{'tag'} === "Monday"))) {
$__t1 = null;;
if ((is_object($y_1) && (($y_1)->{'tag'} === "Monday"))) {
$__t1 = new Phpurs_Data0("EQ");
goto end_branch_1;;
};
$__t1 = new Phpurs_Data0("LT");
end_branch_1:;
$__t0 = $__t1;
goto end_branch_0;;
};
  if ((is_object($y_1) && (($y_1)->{'tag'} === "Monday"))) {
$__t0 = new Phpurs_Data0("GT");
goto end_branch_0;;
};
  if ((is_object($x_0) && (($x_0)->{'tag'} === "Tuesday"))) {
$__t2 = null;;
if ((is_object($y_1) && (($y_1)->{'tag'} === "Tuesday"))) {
$__t2 = new Phpurs_Data0("EQ");
goto end_branch_2;;
};
$__t2 = new Phpurs_Data0("LT");
end_branch_2:;
$__t0 = $__t2;
goto end_branch_0;;
};
  if ((is_object($y_1) && (($y_1)->{'tag'} === "Tuesday"))) {
$__t0 = new Phpurs_Data0("GT");
goto end_branch_0;;
};
  if ((is_object($x_0) && (($x_0)->{'tag'} === "Wednesday"))) {
$__t3 = null;;
if ((is_object($y_1) && (($y_1)->{'tag'} === "Wednesday"))) {
$__t3 = new Phpurs_Data0("EQ");
goto end_branch_3;;
};
$__t3 = new Phpurs_Data0("LT");
end_branch_3:;
$__t0 = $__t3;
goto end_branch_0;;
};
  if ((is_object($y_1) && (($y_1)->{'tag'} === "Wednesday"))) {
$__t0 = new Phpurs_Data0("GT");
goto end_branch_0;;
};
  if ((is_object($x_0) && (($x_0)->{'tag'} === "Thursday"))) {
$__t4 = null;;
if ((is_object($y_1) && (($y_1)->{'tag'} === "Thursday"))) {
$__t4 = new Phpurs_Data0("EQ");
goto end_branch_4;;
};
$__t4 = new Phpurs_Data0("LT");
end_branch_4:;
$__t0 = $__t4;
goto end_branch_0;;
};
  if ((is_object($y_1) && (($y_1)->{'tag'} === "Thursday"))) {
$__t0 = new Phpurs_Data0("GT");
goto end_branch_0;;
};
  if ((is_object($x_0) && (($x_0)->{'tag'} === "Friday"))) {
$__t5 = null;;
if ((is_object($y_1) && (($y_1)->{'tag'} === "Friday"))) {
$__t5 = new Phpurs_Data0("EQ");
goto end_branch_5;;
};
$__t5 = new Phpurs_Data0("LT");
end_branch_5:;
$__t0 = $__t5;
goto end_branch_0;;
};
  if ((is_object($y_1) && (($y_1)->{'tag'} === "Friday"))) {
$__t0 = new Phpurs_Data0("GT");
goto end_branch_0;;
};
  if ((is_object($x_0) && (($x_0)->{'tag'} === "Saturday"))) {
$__t6 = null;;
if ((is_object($y_1) && (($y_1)->{'tag'} === "Saturday"))) {
$__t6 = new Phpurs_Data0("EQ");
goto end_branch_6;;
};
$__t6 = new Phpurs_Data0("LT");
end_branch_6:;
$__t0 = $__t6;
goto end_branch_0;;
};
  if ((is_object($y_1) && (($y_1)->{'tag'} === "Saturday"))) {
$__t0 = new Phpurs_Data0("GT");
goto end_branch_0;;
};
  if (((is_object($x_0) && (($x_0)->{'tag'} === "Sunday")) && (is_object($y_1) && (($y_1)->{'tag'} === "Sunday")))) {
$__t0 = new Phpurs_Data0("EQ");
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Date_Component_eqWeekday'] ?? \PhpursThunks::eval('Data_Date_Component_eqWeekday'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Date_Component_eqMonth'] = function() { $v = ["eq" => (function() {
  $__fn = function($x_0 = null, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ((is_object($x_0) && (($x_0)->{'tag'} === "January"))) {
$__t0 = (is_object($y_1) && (($y_1)->{'tag'} === "January"));
goto end_branch_0;;
};
  if ((is_object($x_0) && (($x_0)->{'tag'} === "February"))) {
$__t0 = (is_object($y_1) && (($y_1)->{'tag'} === "February"));
goto end_branch_0;;
};
  if ((is_object($x_0) && (($x_0)->{'tag'} === "March"))) {
$__t0 = (is_object($y_1) && (($y_1)->{'tag'} === "March"));
goto end_branch_0;;
};
  if ((is_object($x_0) && (($x_0)->{'tag'} === "April"))) {
$__t0 = (is_object($y_1) && (($y_1)->{'tag'} === "April"));
goto end_branch_0;;
};
  if ((is_object($x_0) && (($x_0)->{'tag'} === "May"))) {
$__t0 = (is_object($y_1) && (($y_1)->{'tag'} === "May"));
goto end_branch_0;;
};
  if ((is_object($x_0) && (($x_0)->{'tag'} === "June"))) {
$__t0 = (is_object($y_1) && (($y_1)->{'tag'} === "June"));
goto end_branch_0;;
};
  if ((is_object($x_0) && (($x_0)->{'tag'} === "July"))) {
$__t0 = (is_object($y_1) && (($y_1)->{'tag'} === "July"));
goto end_branch_0;;
};
  if ((is_object($x_0) && (($x_0)->{'tag'} === "August"))) {
$__t0 = (is_object($y_1) && (($y_1)->{'tag'} === "August"));
goto end_branch_0;;
};
  if ((is_object($x_0) && (($x_0)->{'tag'} === "September"))) {
$__t0 = (is_object($y_1) && (($y_1)->{'tag'} === "September"));
goto end_branch_0;;
};
  if ((is_object($x_0) && (($x_0)->{'tag'} === "October"))) {
$__t0 = (is_object($y_1) && (($y_1)->{'tag'} === "October"));
goto end_branch_0;;
};
  if ((is_object($x_0) && (($x_0)->{'tag'} === "November"))) {
$__t0 = (is_object($y_1) && (($y_1)->{'tag'} === "November"));
goto end_branch_0;;
};
  $__t0 = ((is_object($x_0) && (($x_0)->{'tag'} === "December")) && (is_object($y_1) && (($y_1)->{'tag'} === "December")));
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Data_Date_Component_ordMonth'] = function() { $v = ["compare" => (function() {
  $__fn = function($x_0 = null, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ((is_object($x_0) && (($x_0)->{'tag'} === "January"))) {
$__t1 = null;;
if ((is_object($y_1) && (($y_1)->{'tag'} === "January"))) {
$__t1 = new Phpurs_Data0("EQ");
goto end_branch_1;;
};
$__t1 = new Phpurs_Data0("LT");
end_branch_1:;
$__t0 = $__t1;
goto end_branch_0;;
};
  if ((is_object($y_1) && (($y_1)->{'tag'} === "January"))) {
$__t0 = new Phpurs_Data0("GT");
goto end_branch_0;;
};
  if ((is_object($x_0) && (($x_0)->{'tag'} === "February"))) {
$__t2 = null;;
if ((is_object($y_1) && (($y_1)->{'tag'} === "February"))) {
$__t2 = new Phpurs_Data0("EQ");
goto end_branch_2;;
};
$__t2 = new Phpurs_Data0("LT");
end_branch_2:;
$__t0 = $__t2;
goto end_branch_0;;
};
  if ((is_object($y_1) && (($y_1)->{'tag'} === "February"))) {
$__t0 = new Phpurs_Data0("GT");
goto end_branch_0;;
};
  if ((is_object($x_0) && (($x_0)->{'tag'} === "March"))) {
$__t3 = null;;
if ((is_object($y_1) && (($y_1)->{'tag'} === "March"))) {
$__t3 = new Phpurs_Data0("EQ");
goto end_branch_3;;
};
$__t3 = new Phpurs_Data0("LT");
end_branch_3:;
$__t0 = $__t3;
goto end_branch_0;;
};
  if ((is_object($y_1) && (($y_1)->{'tag'} === "March"))) {
$__t0 = new Phpurs_Data0("GT");
goto end_branch_0;;
};
  if ((is_object($x_0) && (($x_0)->{'tag'} === "April"))) {
$__t4 = null;;
if ((is_object($y_1) && (($y_1)->{'tag'} === "April"))) {
$__t4 = new Phpurs_Data0("EQ");
goto end_branch_4;;
};
$__t4 = new Phpurs_Data0("LT");
end_branch_4:;
$__t0 = $__t4;
goto end_branch_0;;
};
  if ((is_object($y_1) && (($y_1)->{'tag'} === "April"))) {
$__t0 = new Phpurs_Data0("GT");
goto end_branch_0;;
};
  if ((is_object($x_0) && (($x_0)->{'tag'} === "May"))) {
$__t5 = null;;
if ((is_object($y_1) && (($y_1)->{'tag'} === "May"))) {
$__t5 = new Phpurs_Data0("EQ");
goto end_branch_5;;
};
$__t5 = new Phpurs_Data0("LT");
end_branch_5:;
$__t0 = $__t5;
goto end_branch_0;;
};
  if ((is_object($y_1) && (($y_1)->{'tag'} === "May"))) {
$__t0 = new Phpurs_Data0("GT");
goto end_branch_0;;
};
  if ((is_object($x_0) && (($x_0)->{'tag'} === "June"))) {
$__t6 = null;;
if ((is_object($y_1) && (($y_1)->{'tag'} === "June"))) {
$__t6 = new Phpurs_Data0("EQ");
goto end_branch_6;;
};
$__t6 = new Phpurs_Data0("LT");
end_branch_6:;
$__t0 = $__t6;
goto end_branch_0;;
};
  if ((is_object($y_1) && (($y_1)->{'tag'} === "June"))) {
$__t0 = new Phpurs_Data0("GT");
goto end_branch_0;;
};
  if ((is_object($x_0) && (($x_0)->{'tag'} === "July"))) {
$__t7 = null;;
if ((is_object($y_1) && (($y_1)->{'tag'} === "July"))) {
$__t7 = new Phpurs_Data0("EQ");
goto end_branch_7;;
};
$__t7 = new Phpurs_Data0("LT");
end_branch_7:;
$__t0 = $__t7;
goto end_branch_0;;
};
  if ((is_object($y_1) && (($y_1)->{'tag'} === "July"))) {
$__t0 = new Phpurs_Data0("GT");
goto end_branch_0;;
};
  if ((is_object($x_0) && (($x_0)->{'tag'} === "August"))) {
$__t8 = null;;
if ((is_object($y_1) && (($y_1)->{'tag'} === "August"))) {
$__t8 = new Phpurs_Data0("EQ");
goto end_branch_8;;
};
$__t8 = new Phpurs_Data0("LT");
end_branch_8:;
$__t0 = $__t8;
goto end_branch_0;;
};
  if ((is_object($y_1) && (($y_1)->{'tag'} === "August"))) {
$__t0 = new Phpurs_Data0("GT");
goto end_branch_0;;
};
  if ((is_object($x_0) && (($x_0)->{'tag'} === "September"))) {
$__t9 = null;;
if ((is_object($y_1) && (($y_1)->{'tag'} === "September"))) {
$__t9 = new Phpurs_Data0("EQ");
goto end_branch_9;;
};
$__t9 = new Phpurs_Data0("LT");
end_branch_9:;
$__t0 = $__t9;
goto end_branch_0;;
};
  if ((is_object($y_1) && (($y_1)->{'tag'} === "September"))) {
$__t0 = new Phpurs_Data0("GT");
goto end_branch_0;;
};
  if ((is_object($x_0) && (($x_0)->{'tag'} === "October"))) {
$__t10 = null;;
if ((is_object($y_1) && (($y_1)->{'tag'} === "October"))) {
$__t10 = new Phpurs_Data0("EQ");
goto end_branch_10;;
};
$__t10 = new Phpurs_Data0("LT");
end_branch_10:;
$__t0 = $__t10;
goto end_branch_0;;
};
  if ((is_object($y_1) && (($y_1)->{'tag'} === "October"))) {
$__t0 = new Phpurs_Data0("GT");
goto end_branch_0;;
};
  if ((is_object($x_0) && (($x_0)->{'tag'} === "November"))) {
$__t11 = null;;
if ((is_object($y_1) && (($y_1)->{'tag'} === "November"))) {
$__t11 = new Phpurs_Data0("EQ");
goto end_branch_11;;
};
$__t11 = new Phpurs_Data0("LT");
end_branch_11:;
$__t0 = $__t11;
goto end_branch_0;;
};
  if ((is_object($y_1) && (($y_1)->{'tag'} === "November"))) {
$__t0 = new Phpurs_Data0("GT");
goto end_branch_0;;
};
  if (((is_object($x_0) && (($x_0)->{'tag'} === "December")) && (is_object($y_1) && (($y_1)->{'tag'} === "December")))) {
$__t0 = new Phpurs_Data0("EQ");
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Date_Component_eqMonth'] ?? \PhpursThunks::eval('Data_Date_Component_eqMonth'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Date_Component_eqDay'] = function() { $v = ($GLOBALS['Data_Eq_eqInt'] ?? \PhpursThunks::eval('Data_Eq_eqInt')); return $v; };
\PhpursThunks::$thunks['Data_Date_Component_boundedYear'] = function() { $v = ["bottom" => -271820, "top" => 275759, "Ord0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Date_Component_boundedWeekday'] = function() { $v = ["bottom" => new Phpurs_Data0("Monday"), "top" => new Phpurs_Data0("Sunday"), "Ord0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Date_Component_ordWeekday'] ?? \PhpursThunks::eval('Data_Date_Component_ordWeekday'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Date_Component_boundedMonth'] = function() { $v = ["bottom" => new Phpurs_Data0("January"), "top" => new Phpurs_Data0("December"), "Ord0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Date_Component_ordMonth'] ?? \PhpursThunks::eval('Data_Date_Component_ordMonth'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Date_Component_boundedEnumYear'] = function() { $v = ["cardinality" => 547580, "toEnum" => function($n_0 = null) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((($n_0 >= -271820) && ($n_0 <= 275759))) {
$__t0 = new Phpurs_Data1("Just", $n_0);
goto end_branch_0;;
};
  $__t0 = new Phpurs_Data0("Nothing");
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "fromEnum" => function($v_0 = null) {
  $__num = \func_num_args();
  $__res = $v_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bounded0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Date_Component_boundedYear'] ?? \PhpursThunks::eval('Data_Date_Component_boundedYear'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Enum1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Date_Component_enumYear'] ?? \PhpursThunks::eval('Data_Date_Component_enumYear'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Date_Component_enumYear'] = function() { $v = ["succ" => ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_Date_Component_boundedEnumYear'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumYear')))['toEnum']))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ($v_0 + 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Data_Date_Component_boundedEnumYear'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumYear')))['fromEnum'])), "pred" => ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_Date_Component_boundedEnumYear'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumYear')))['toEnum']))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ($v_0 - 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Data_Date_Component_boundedEnumYear'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumYear')))['fromEnum'])), "Ord0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Date_Component_boundedEnumWeekday'] = function() { $v = ["cardinality" => 7, "toEnum" => function($v_0 = null) {
  $__num = \func_num_args();
  $__t0 = null;;
  switch ($v_0) {
case 1:
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data0("Monday"));
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($v_0) {
case 2:
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data0("Tuesday"));
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($v_0) {
case 3:
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data0("Wednesday"));
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($v_0) {
case 4:
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data0("Thursday"));
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($v_0) {
case 5:
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data0("Friday"));
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($v_0) {
case 6:
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data0("Saturday"));
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($v_0) {
case 7:
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data0("Sunday"));
goto end_branch_0;;
break;
default:
;
break;
};
  $__t0 = new Phpurs_Data0("Nothing");
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "fromEnum" => function($v_0 = null) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Monday"))) {
$__t1 = 1;
goto end_branch_1;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Tuesday"))) {
$__t1 = 2;
goto end_branch_1;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Wednesday"))) {
$__t1 = 3;
goto end_branch_1;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Thursday"))) {
$__t1 = 4;
goto end_branch_1;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Friday"))) {
$__t1 = 5;
goto end_branch_1;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Saturday"))) {
$__t1 = 6;
goto end_branch_1;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Sunday"))) {
$__t1 = 7;
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bounded0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Date_Component_boundedWeekday'] ?? \PhpursThunks::eval('Data_Date_Component_boundedWeekday'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Enum1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Date_Component_enumWeekday'] ?? \PhpursThunks::eval('Data_Date_Component_enumWeekday'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Date_Component_enumWeekday'] = function() { $v = ["succ" => ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_Date_Component_boundedEnumWeekday'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumWeekday')))['toEnum']))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ($v_0 + 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Data_Date_Component_boundedEnumWeekday'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumWeekday')))['fromEnum'])), "pred" => ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_Date_Component_boundedEnumWeekday'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumWeekday')))['toEnum']))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ($v_0 - 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Data_Date_Component_boundedEnumWeekday'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumWeekday')))['fromEnum'])), "Ord0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Date_Component_ordWeekday'] ?? \PhpursThunks::eval('Data_Date_Component_ordWeekday'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Date_Component_boundedEnumMonth'] = function() { $v = ["cardinality" => 12, "toEnum" => function($v_0 = null) {
  $__num = \func_num_args();
  $__t0 = null;;
  switch ($v_0) {
case 1:
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data0("January"));
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($v_0) {
case 2:
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data0("February"));
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($v_0) {
case 3:
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data0("March"));
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($v_0) {
case 4:
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data0("April"));
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($v_0) {
case 5:
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data0("May"));
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($v_0) {
case 6:
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data0("June"));
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($v_0) {
case 7:
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data0("July"));
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($v_0) {
case 8:
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data0("August"));
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($v_0) {
case 9:
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data0("September"));
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($v_0) {
case 10:
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data0("October"));
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($v_0) {
case 11:
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data0("November"));
goto end_branch_0;;
break;
default:
;
break;
};
  switch ($v_0) {
case 12:
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data0("December"));
goto end_branch_0;;
break;
default:
;
break;
};
  $__t0 = new Phpurs_Data0("Nothing");
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "fromEnum" => function($v_0 = null) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((is_object($v_0) && (($v_0)->{'tag'} === "January"))) {
$__t1 = 1;
goto end_branch_1;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "February"))) {
$__t1 = 2;
goto end_branch_1;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "March"))) {
$__t1 = 3;
goto end_branch_1;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "April"))) {
$__t1 = 4;
goto end_branch_1;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "May"))) {
$__t1 = 5;
goto end_branch_1;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "June"))) {
$__t1 = 6;
goto end_branch_1;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "July"))) {
$__t1 = 7;
goto end_branch_1;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "August"))) {
$__t1 = 8;
goto end_branch_1;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "September"))) {
$__t1 = 9;
goto end_branch_1;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "October"))) {
$__t1 = 10;
goto end_branch_1;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "November"))) {
$__t1 = 11;
goto end_branch_1;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "December"))) {
$__t1 = 12;
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bounded0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Date_Component_boundedMonth'] ?? \PhpursThunks::eval('Data_Date_Component_boundedMonth'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Enum1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Date_Component_enumMonth'] ?? \PhpursThunks::eval('Data_Date_Component_enumMonth'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Date_Component_enumMonth'] = function() { $v = ["succ" => ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_Date_Component_boundedEnumMonth'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumMonth')))['toEnum']))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ($v_0 + 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Data_Date_Component_boundedEnumMonth'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumMonth')))['fromEnum'])), "pred" => ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_Date_Component_boundedEnumMonth'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumMonth')))['toEnum']))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ($v_0 - 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Data_Date_Component_boundedEnumMonth'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumMonth')))['fromEnum'])), "Ord0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Date_Component_ordMonth'] ?? \PhpursThunks::eval('Data_Date_Component_ordMonth'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Date_Component_boundedDay'] = function() { $v = ["bottom" => 1, "top" => 31, "Ord0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Date_Component_boundedEnumDay'] = function() { $v = ["cardinality" => 31, "toEnum" => function($n_0 = null) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((($n_0 >= 1) && ($n_0 <= 31))) {
$__t0 = new Phpurs_Data1("Just", $n_0);
goto end_branch_0;;
};
  $__t0 = new Phpurs_Data0("Nothing");
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "fromEnum" => function($v_0 = null) {
  $__num = \func_num_args();
  $__res = $v_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bounded0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Date_Component_boundedDay'] ?? \PhpursThunks::eval('Data_Date_Component_boundedDay'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Enum1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Date_Component_enumDay'] ?? \PhpursThunks::eval('Data_Date_Component_enumDay'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Date_Component_enumDay'] = function() { $v = ["succ" => ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_Date_Component_boundedEnumDay'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumDay')))['toEnum']))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ($v_0 + 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Data_Date_Component_boundedEnumDay'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumDay')))['fromEnum'])), "pred" => ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_Date_Component_boundedEnumDay'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumDay')))['toEnum']))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ($v_0 - 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Data_Date_Component_boundedEnumDay'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumDay')))['fromEnum'])), "Ord0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };













































