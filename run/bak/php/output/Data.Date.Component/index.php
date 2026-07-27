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

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };


final class Data_Date_Component_Monday { public $tag = 'Monday'; public function __construct() {} }
final class Data_Date_Component_Tuesday { public $tag = 'Tuesday'; public function __construct() {} }
final class Data_Date_Component_Wednesday { public $tag = 'Wednesday'; public function __construct() {} }
final class Data_Date_Component_Thursday { public $tag = 'Thursday'; public function __construct() {} }
final class Data_Date_Component_Friday { public $tag = 'Friday'; public function __construct() {} }
final class Data_Date_Component_Saturday { public $tag = 'Saturday'; public function __construct() {} }
final class Data_Date_Component_Sunday { public $tag = 'Sunday'; public function __construct() {} }
final class Data_Date_Component_January { public $tag = 'January'; public function __construct() {} }
final class Data_Date_Component_February { public $tag = 'February'; public function __construct() {} }
final class Data_Date_Component_March { public $tag = 'March'; public function __construct() {} }
final class Data_Date_Component_April { public $tag = 'April'; public function __construct() {} }
final class Data_Date_Component_May { public $tag = 'May'; public function __construct() {} }
final class Data_Date_Component_June { public $tag = 'June'; public function __construct() {} }
final class Data_Date_Component_July { public $tag = 'July'; public function __construct() {} }
final class Data_Date_Component_August { public $tag = 'August'; public function __construct() {} }
final class Data_Date_Component_September { public $tag = 'September'; public function __construct() {} }
final class Data_Date_Component_October { public $tag = 'October'; public function __construct() {} }
final class Data_Date_Component_November { public $tag = 'November'; public function __construct() {} }
final class Data_Date_Component_December { public $tag = 'December'; public function __construct() {} }

// Data_Date_Component_greaterThanOrEq
$GLOBALS['Data_Date_Component_greaterThanOrEq'] = (function() use (&$__fn) {
$__local_var_0_0 = ((($GLOBALS['Data_Ord_ordIntImpl'])(new \Data\Ordering\Data_Ordering_LT()))(new \Data\Ordering\Data_Ordering_EQ()))(new \Data\Ordering\Data_Ordering_GT());
return (function() use ($__local_var_0_0) {
  $__fn = function($a1_1, $a2_2 = null) use ($__local_var_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ( ! (($__local_var_0_0)($a1_1))($a2_2) instanceof \Data\Ordering\Data_Ordering_LT);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
})();

// Data_Date_Component_lessThanOrEq
$GLOBALS['Data_Date_Component_lessThanOrEq'] = (function() use (&$__fn) {
$__local_var_0_0 = ((($GLOBALS['Data_Ord_ordIntImpl'])(new \Data\Ordering\Data_Ordering_LT()))(new \Data\Ordering\Data_Ordering_EQ()))(new \Data\Ordering\Data_Ordering_GT());
return (function() use ($__local_var_0_0) {
  $__fn = function($a1_1, $a2_2 = null) use ($__local_var_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ( ! (($__local_var_0_0)($a1_1))($a2_2) instanceof \Data\Ordering\Data_Ordering_GT);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
})();

// Data_Date_Component_Monday
$GLOBALS['Data_Date_Component_Monday'] = ($GLOBALS['__phpurs_data0_Monday'] ??= new \Data\Date\Component\Data_Date_Component_Monday());

// Data_Date_Component_Tuesday
$GLOBALS['Data_Date_Component_Tuesday'] = ($GLOBALS['__phpurs_data0_Tuesday'] ??= new \Data\Date\Component\Data_Date_Component_Tuesday());

// Data_Date_Component_Wednesday
$GLOBALS['Data_Date_Component_Wednesday'] = ($GLOBALS['__phpurs_data0_Wednesday'] ??= new \Data\Date\Component\Data_Date_Component_Wednesday());

// Data_Date_Component_Thursday
$GLOBALS['Data_Date_Component_Thursday'] = ($GLOBALS['__phpurs_data0_Thursday'] ??= new \Data\Date\Component\Data_Date_Component_Thursday());

// Data_Date_Component_Friday
$GLOBALS['Data_Date_Component_Friday'] = ($GLOBALS['__phpurs_data0_Friday'] ??= new \Data\Date\Component\Data_Date_Component_Friday());

// Data_Date_Component_Saturday
$GLOBALS['Data_Date_Component_Saturday'] = ($GLOBALS['__phpurs_data0_Saturday'] ??= new \Data\Date\Component\Data_Date_Component_Saturday());

// Data_Date_Component_Sunday
$GLOBALS['Data_Date_Component_Sunday'] = ($GLOBALS['__phpurs_data0_Sunday'] ??= new \Data\Date\Component\Data_Date_Component_Sunday());

// Data_Date_Component_January
$GLOBALS['Data_Date_Component_January'] = ($GLOBALS['__phpurs_data0_January'] ??= new \Data\Date\Component\Data_Date_Component_January());

// Data_Date_Component_February
$GLOBALS['Data_Date_Component_February'] = ($GLOBALS['__phpurs_data0_February'] ??= new \Data\Date\Component\Data_Date_Component_February());

// Data_Date_Component_March
$GLOBALS['Data_Date_Component_March'] = ($GLOBALS['__phpurs_data0_March'] ??= new \Data\Date\Component\Data_Date_Component_March());

// Data_Date_Component_April
$GLOBALS['Data_Date_Component_April'] = ($GLOBALS['__phpurs_data0_April'] ??= new \Data\Date\Component\Data_Date_Component_April());

// Data_Date_Component_May
$GLOBALS['Data_Date_Component_May'] = ($GLOBALS['__phpurs_data0_May'] ??= new \Data\Date\Component\Data_Date_Component_May());

// Data_Date_Component_June
$GLOBALS['Data_Date_Component_June'] = ($GLOBALS['__phpurs_data0_June'] ??= new \Data\Date\Component\Data_Date_Component_June());

// Data_Date_Component_July
$GLOBALS['Data_Date_Component_July'] = ($GLOBALS['__phpurs_data0_July'] ??= new \Data\Date\Component\Data_Date_Component_July());

// Data_Date_Component_August
$GLOBALS['Data_Date_Component_August'] = ($GLOBALS['__phpurs_data0_August'] ??= new \Data\Date\Component\Data_Date_Component_August());

// Data_Date_Component_September
$GLOBALS['Data_Date_Component_September'] = ($GLOBALS['__phpurs_data0_September'] ??= new \Data\Date\Component\Data_Date_Component_September());

// Data_Date_Component_October
$GLOBALS['Data_Date_Component_October'] = ($GLOBALS['__phpurs_data0_October'] ??= new \Data\Date\Component\Data_Date_Component_October());

// Data_Date_Component_November
$GLOBALS['Data_Date_Component_November'] = ($GLOBALS['__phpurs_data0_November'] ??= new \Data\Date\Component\Data_Date_Component_November());

// Data_Date_Component_December
$GLOBALS['Data_Date_Component_December'] = ($GLOBALS['__phpurs_data0_December'] ??= new \Data\Date\Component\Data_Date_Component_December());

// Data_Date_Component_showYear
$GLOBALS['Data_Date_Component_showYear'] = (object)["show" => function($v_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})("(Year "))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})((($GLOBALS['Data_Show_showInt'])->{'show'})($v_0)))(")"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Date_Component_showWeekday
$GLOBALS['Data_Date_Component_showWeekday'] = (object)["show" => function($v_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v_0 instanceof \Data\Date\Component\Data_Date_Component_Monday) {
$__t0 = "Monday";
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Date\Component\Data_Date_Component_Tuesday) {
$__t0 = "Tuesday";
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Date\Component\Data_Date_Component_Wednesday) {
$__t0 = "Wednesday";
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Date\Component\Data_Date_Component_Thursday) {
$__t0 = "Thursday";
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Date\Component\Data_Date_Component_Friday) {
$__t0 = "Friday";
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Date\Component\Data_Date_Component_Saturday) {
$__t0 = "Saturday";
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Date\Component\Data_Date_Component_Sunday) {
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
}];

// Data_Date_Component_showMonth
$GLOBALS['Data_Date_Component_showMonth'] = (object)["show" => function($v_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v_0 instanceof \Data\Date\Component\Data_Date_Component_January) {
$__t0 = "January";
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Date\Component\Data_Date_Component_February) {
$__t0 = "February";
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Date\Component\Data_Date_Component_March) {
$__t0 = "March";
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Date\Component\Data_Date_Component_April) {
$__t0 = "April";
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Date\Component\Data_Date_Component_May) {
$__t0 = "May";
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Date\Component\Data_Date_Component_June) {
$__t0 = "June";
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Date\Component\Data_Date_Component_July) {
$__t0 = "July";
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Date\Component\Data_Date_Component_August) {
$__t0 = "August";
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Date\Component\Data_Date_Component_September) {
$__t0 = "September";
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Date\Component\Data_Date_Component_October) {
$__t0 = "October";
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Date\Component\Data_Date_Component_November) {
$__t0 = "November";
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Date\Component\Data_Date_Component_December) {
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
}];

// Data_Date_Component_showDay
$GLOBALS['Data_Date_Component_showDay'] = (object)["show" => function($v_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})("(Day "))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})((($GLOBALS['Data_Show_showInt'])->{'show'})($v_0)))(")"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Date_Component_ordYear
$GLOBALS['Data_Date_Component_ordYear'] = (object)["compare" => ((($GLOBALS['Data_Ord_ordIntImpl'])(new \Data\Ordering\Data_Ordering_LT()))(new \Data\Ordering\Data_Ordering_EQ()))(new \Data\Ordering\Data_Ordering_GT()), "Eq0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = (object)["eq" => $GLOBALS['Data_Eq_eqIntImpl']];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Date_Component_ordDay
$GLOBALS['Data_Date_Component_ordDay'] = (object)["compare" => ((($GLOBALS['Data_Ord_ordIntImpl'])(new \Data\Ordering\Data_Ordering_LT()))(new \Data\Ordering\Data_Ordering_EQ()))(new \Data\Ordering\Data_Ordering_GT()), "Eq0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = (object)["eq" => $GLOBALS['Data_Eq_eqIntImpl']];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Date_Component_eqYear
$GLOBALS['Data_Date_Component_eqYear'] = (object)["eq" => $GLOBALS['Data_Eq_eqIntImpl']];

// Data_Date_Component_eqWeekday
$GLOBALS['Data_Date_Component_eqWeekday'] = (object)["eq" => (function() {
  $__fn = function($x_0, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ($x_0 instanceof \Data\Date\Component\Data_Date_Component_Monday) {
$__t0 = $y_1 instanceof \Data\Date\Component\Data_Date_Component_Monday;
goto end_branch_0;;
};
  if ($x_0 instanceof \Data\Date\Component\Data_Date_Component_Tuesday) {
$__t0 = $y_1 instanceof \Data\Date\Component\Data_Date_Component_Tuesday;
goto end_branch_0;;
};
  if ($x_0 instanceof \Data\Date\Component\Data_Date_Component_Wednesday) {
$__t0 = $y_1 instanceof \Data\Date\Component\Data_Date_Component_Wednesday;
goto end_branch_0;;
};
  if ($x_0 instanceof \Data\Date\Component\Data_Date_Component_Thursday) {
$__t0 = $y_1 instanceof \Data\Date\Component\Data_Date_Component_Thursday;
goto end_branch_0;;
};
  if ($x_0 instanceof \Data\Date\Component\Data_Date_Component_Friday) {
$__t0 = $y_1 instanceof \Data\Date\Component\Data_Date_Component_Friday;
goto end_branch_0;;
};
  if ($x_0 instanceof \Data\Date\Component\Data_Date_Component_Saturday) {
$__t0 = $y_1 instanceof \Data\Date\Component\Data_Date_Component_Saturday;
goto end_branch_0;;
};
  $__t0 = ($x_0 instanceof \Data\Date\Component\Data_Date_Component_Sunday && $y_1 instanceof \Data\Date\Component\Data_Date_Component_Sunday);
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];

// Data_Date_Component_ordWeekday
$GLOBALS['Data_Date_Component_ordWeekday'] = (object)["compare" => (function() {
  $__fn = function($x_0, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ($x_0 instanceof \Data\Date\Component\Data_Date_Component_Monday) {
$__t1 = null;;
if ($y_1 instanceof \Data\Date\Component\Data_Date_Component_Monday) {
$__t1 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_1;;
};
$__t1 = new \Data\Ordering\Data_Ordering_LT();
end_branch_1:;
$__t0 = $__t1;
goto end_branch_0;;
};
  if ($y_1 instanceof \Data\Date\Component\Data_Date_Component_Monday) {
$__t0 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_0;;
};
  if ($x_0 instanceof \Data\Date\Component\Data_Date_Component_Tuesday) {
$__t2 = null;;
if ($y_1 instanceof \Data\Date\Component\Data_Date_Component_Tuesday) {
$__t2 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_2;;
};
$__t2 = new \Data\Ordering\Data_Ordering_LT();
end_branch_2:;
$__t0 = $__t2;
goto end_branch_0;;
};
  if ($y_1 instanceof \Data\Date\Component\Data_Date_Component_Tuesday) {
$__t0 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_0;;
};
  if ($x_0 instanceof \Data\Date\Component\Data_Date_Component_Wednesday) {
$__t3 = null;;
if ($y_1 instanceof \Data\Date\Component\Data_Date_Component_Wednesday) {
$__t3 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_3;;
};
$__t3 = new \Data\Ordering\Data_Ordering_LT();
end_branch_3:;
$__t0 = $__t3;
goto end_branch_0;;
};
  if ($y_1 instanceof \Data\Date\Component\Data_Date_Component_Wednesday) {
$__t0 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_0;;
};
  if ($x_0 instanceof \Data\Date\Component\Data_Date_Component_Thursday) {
$__t4 = null;;
if ($y_1 instanceof \Data\Date\Component\Data_Date_Component_Thursday) {
$__t4 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_4;;
};
$__t4 = new \Data\Ordering\Data_Ordering_LT();
end_branch_4:;
$__t0 = $__t4;
goto end_branch_0;;
};
  if ($y_1 instanceof \Data\Date\Component\Data_Date_Component_Thursday) {
$__t0 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_0;;
};
  if ($x_0 instanceof \Data\Date\Component\Data_Date_Component_Friday) {
$__t5 = null;;
if ($y_1 instanceof \Data\Date\Component\Data_Date_Component_Friday) {
$__t5 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_5;;
};
$__t5 = new \Data\Ordering\Data_Ordering_LT();
end_branch_5:;
$__t0 = $__t5;
goto end_branch_0;;
};
  if ($y_1 instanceof \Data\Date\Component\Data_Date_Component_Friday) {
$__t0 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_0;;
};
  if ($x_0 instanceof \Data\Date\Component\Data_Date_Component_Saturday) {
$__t6 = null;;
if ($y_1 instanceof \Data\Date\Component\Data_Date_Component_Saturday) {
$__t6 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_6;;
};
$__t6 = new \Data\Ordering\Data_Ordering_LT();
end_branch_6:;
$__t0 = $__t6;
goto end_branch_0;;
};
  if ($y_1 instanceof \Data\Date\Component\Data_Date_Component_Saturday) {
$__t0 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_0;;
};
  if (($x_0 instanceof \Data\Date\Component\Data_Date_Component_Sunday && $y_1 instanceof \Data\Date\Component\Data_Date_Component_Sunday)) {
$__t0 = new \Data\Ordering\Data_Ordering_EQ();
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
})(), "Eq0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Date_Component_eqWeekday'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Date_Component_eqMonth
$GLOBALS['Data_Date_Component_eqMonth'] = (object)["eq" => (function() {
  $__fn = function($x_0, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ($x_0 instanceof \Data\Date\Component\Data_Date_Component_January) {
$__t0 = $y_1 instanceof \Data\Date\Component\Data_Date_Component_January;
goto end_branch_0;;
};
  if ($x_0 instanceof \Data\Date\Component\Data_Date_Component_February) {
$__t0 = $y_1 instanceof \Data\Date\Component\Data_Date_Component_February;
goto end_branch_0;;
};
  if ($x_0 instanceof \Data\Date\Component\Data_Date_Component_March) {
$__t0 = $y_1 instanceof \Data\Date\Component\Data_Date_Component_March;
goto end_branch_0;;
};
  if ($x_0 instanceof \Data\Date\Component\Data_Date_Component_April) {
$__t0 = $y_1 instanceof \Data\Date\Component\Data_Date_Component_April;
goto end_branch_0;;
};
  if ($x_0 instanceof \Data\Date\Component\Data_Date_Component_May) {
$__t0 = $y_1 instanceof \Data\Date\Component\Data_Date_Component_May;
goto end_branch_0;;
};
  if ($x_0 instanceof \Data\Date\Component\Data_Date_Component_June) {
$__t0 = $y_1 instanceof \Data\Date\Component\Data_Date_Component_June;
goto end_branch_0;;
};
  if ($x_0 instanceof \Data\Date\Component\Data_Date_Component_July) {
$__t0 = $y_1 instanceof \Data\Date\Component\Data_Date_Component_July;
goto end_branch_0;;
};
  if ($x_0 instanceof \Data\Date\Component\Data_Date_Component_August) {
$__t0 = $y_1 instanceof \Data\Date\Component\Data_Date_Component_August;
goto end_branch_0;;
};
  if ($x_0 instanceof \Data\Date\Component\Data_Date_Component_September) {
$__t0 = $y_1 instanceof \Data\Date\Component\Data_Date_Component_September;
goto end_branch_0;;
};
  if ($x_0 instanceof \Data\Date\Component\Data_Date_Component_October) {
$__t0 = $y_1 instanceof \Data\Date\Component\Data_Date_Component_October;
goto end_branch_0;;
};
  if ($x_0 instanceof \Data\Date\Component\Data_Date_Component_November) {
$__t0 = $y_1 instanceof \Data\Date\Component\Data_Date_Component_November;
goto end_branch_0;;
};
  $__t0 = ($x_0 instanceof \Data\Date\Component\Data_Date_Component_December && $y_1 instanceof \Data\Date\Component\Data_Date_Component_December);
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];

// Data_Date_Component_ordMonth
$GLOBALS['Data_Date_Component_ordMonth'] = (object)["compare" => (function() {
  $__fn = function($x_0, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ($x_0 instanceof \Data\Date\Component\Data_Date_Component_January) {
$__t1 = null;;
if ($y_1 instanceof \Data\Date\Component\Data_Date_Component_January) {
$__t1 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_1;;
};
$__t1 = new \Data\Ordering\Data_Ordering_LT();
end_branch_1:;
$__t0 = $__t1;
goto end_branch_0;;
};
  if ($y_1 instanceof \Data\Date\Component\Data_Date_Component_January) {
$__t0 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_0;;
};
  if ($x_0 instanceof \Data\Date\Component\Data_Date_Component_February) {
$__t2 = null;;
if ($y_1 instanceof \Data\Date\Component\Data_Date_Component_February) {
$__t2 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_2;;
};
$__t2 = new \Data\Ordering\Data_Ordering_LT();
end_branch_2:;
$__t0 = $__t2;
goto end_branch_0;;
};
  if ($y_1 instanceof \Data\Date\Component\Data_Date_Component_February) {
$__t0 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_0;;
};
  if ($x_0 instanceof \Data\Date\Component\Data_Date_Component_March) {
$__t3 = null;;
if ($y_1 instanceof \Data\Date\Component\Data_Date_Component_March) {
$__t3 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_3;;
};
$__t3 = new \Data\Ordering\Data_Ordering_LT();
end_branch_3:;
$__t0 = $__t3;
goto end_branch_0;;
};
  if ($y_1 instanceof \Data\Date\Component\Data_Date_Component_March) {
$__t0 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_0;;
};
  if ($x_0 instanceof \Data\Date\Component\Data_Date_Component_April) {
$__t4 = null;;
if ($y_1 instanceof \Data\Date\Component\Data_Date_Component_April) {
$__t4 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_4;;
};
$__t4 = new \Data\Ordering\Data_Ordering_LT();
end_branch_4:;
$__t0 = $__t4;
goto end_branch_0;;
};
  if ($y_1 instanceof \Data\Date\Component\Data_Date_Component_April) {
$__t0 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_0;;
};
  if ($x_0 instanceof \Data\Date\Component\Data_Date_Component_May) {
$__t5 = null;;
if ($y_1 instanceof \Data\Date\Component\Data_Date_Component_May) {
$__t5 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_5;;
};
$__t5 = new \Data\Ordering\Data_Ordering_LT();
end_branch_5:;
$__t0 = $__t5;
goto end_branch_0;;
};
  if ($y_1 instanceof \Data\Date\Component\Data_Date_Component_May) {
$__t0 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_0;;
};
  if ($x_0 instanceof \Data\Date\Component\Data_Date_Component_June) {
$__t6 = null;;
if ($y_1 instanceof \Data\Date\Component\Data_Date_Component_June) {
$__t6 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_6;;
};
$__t6 = new \Data\Ordering\Data_Ordering_LT();
end_branch_6:;
$__t0 = $__t6;
goto end_branch_0;;
};
  if ($y_1 instanceof \Data\Date\Component\Data_Date_Component_June) {
$__t0 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_0;;
};
  if ($x_0 instanceof \Data\Date\Component\Data_Date_Component_July) {
$__t7 = null;;
if ($y_1 instanceof \Data\Date\Component\Data_Date_Component_July) {
$__t7 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_7;;
};
$__t7 = new \Data\Ordering\Data_Ordering_LT();
end_branch_7:;
$__t0 = $__t7;
goto end_branch_0;;
};
  if ($y_1 instanceof \Data\Date\Component\Data_Date_Component_July) {
$__t0 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_0;;
};
  if ($x_0 instanceof \Data\Date\Component\Data_Date_Component_August) {
$__t8 = null;;
if ($y_1 instanceof \Data\Date\Component\Data_Date_Component_August) {
$__t8 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_8;;
};
$__t8 = new \Data\Ordering\Data_Ordering_LT();
end_branch_8:;
$__t0 = $__t8;
goto end_branch_0;;
};
  if ($y_1 instanceof \Data\Date\Component\Data_Date_Component_August) {
$__t0 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_0;;
};
  if ($x_0 instanceof \Data\Date\Component\Data_Date_Component_September) {
$__t9 = null;;
if ($y_1 instanceof \Data\Date\Component\Data_Date_Component_September) {
$__t9 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_9;;
};
$__t9 = new \Data\Ordering\Data_Ordering_LT();
end_branch_9:;
$__t0 = $__t9;
goto end_branch_0;;
};
  if ($y_1 instanceof \Data\Date\Component\Data_Date_Component_September) {
$__t0 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_0;;
};
  if ($x_0 instanceof \Data\Date\Component\Data_Date_Component_October) {
$__t10 = null;;
if ($y_1 instanceof \Data\Date\Component\Data_Date_Component_October) {
$__t10 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_10;;
};
$__t10 = new \Data\Ordering\Data_Ordering_LT();
end_branch_10:;
$__t0 = $__t10;
goto end_branch_0;;
};
  if ($y_1 instanceof \Data\Date\Component\Data_Date_Component_October) {
$__t0 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_0;;
};
  if ($x_0 instanceof \Data\Date\Component\Data_Date_Component_November) {
$__t11 = null;;
if ($y_1 instanceof \Data\Date\Component\Data_Date_Component_November) {
$__t11 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_11;;
};
$__t11 = new \Data\Ordering\Data_Ordering_LT();
end_branch_11:;
$__t0 = $__t11;
goto end_branch_0;;
};
  if ($y_1 instanceof \Data\Date\Component\Data_Date_Component_November) {
$__t0 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_0;;
};
  if (($x_0 instanceof \Data\Date\Component\Data_Date_Component_December && $y_1 instanceof \Data\Date\Component\Data_Date_Component_December)) {
$__t0 = new \Data\Ordering\Data_Ordering_EQ();
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
})(), "Eq0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Date_Component_eqMonth'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Date_Component_eqDay
$GLOBALS['Data_Date_Component_eqDay'] = (object)["eq" => $GLOBALS['Data_Eq_eqIntImpl']];

// Data_Date_Component_boundedYear
$GLOBALS['Data_Date_Component_boundedYear'] = (object)["bottom" => -271820, "top" => 275759, "Ord0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Date_Component_ordYear'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Date_Component_boundedWeekday
$GLOBALS['Data_Date_Component_boundedWeekday'] = (object)["bottom" => new \Data\Date\Component\Data_Date_Component_Monday(), "top" => new \Data\Date\Component\Data_Date_Component_Sunday(), "Ord0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Date_Component_ordWeekday'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Date_Component_boundedMonth
$GLOBALS['Data_Date_Component_boundedMonth'] = (object)["bottom" => new \Data\Date\Component\Data_Date_Component_January(), "top" => new \Data\Date\Component\Data_Date_Component_December(), "Ord0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Date_Component_ordMonth'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Date_Component_boundedEnumYear
$GLOBALS['Data_Date_Component_boundedEnumYear'] = (object)["cardinality" => 547580, "toEnum" => function($n_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if (((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'conj'})((($GLOBALS['Data_Date_Component_greaterThanOrEq'])($n_0))(-271820)))((($GLOBALS['Data_Date_Component_lessThanOrEq'])($n_0))(275759))) {
$__t0 = new \Data\Maybe\Data_Maybe_Just($n_0);
goto end_branch_0;;
};
  $__t0 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "fromEnum" => function($v_0) {
  $__num = \func_num_args();
  $__res = $v_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bounded0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Date_Component_boundedYear'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Enum1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Date_Component_enumYear'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Date_Component_enumYear
$GLOBALS['Data_Date_Component_enumYear'] = (object)["succ" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Date_Component_boundedEnumYear'])->{'toEnum'}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_0) {
  $__num = \func_num_args();
  $__res = ($v_0 + 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($GLOBALS['Data_Date_Component_boundedEnumYear'])->{'fromEnum'})), "pred" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Date_Component_boundedEnumYear'])->{'toEnum'}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_0) {
  $__num = \func_num_args();
  $__res = ($v_0 - 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($GLOBALS['Data_Date_Component_boundedEnumYear'])->{'fromEnum'})), "Ord0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Date_Component_ordYear'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Date_Component_boundedEnumWeekday
$GLOBALS['Data_Date_Component_boundedEnumWeekday'] = (object)["cardinality" => 7, "toEnum" => function($v_0) {
  $__num = \func_num_args();
  $__res = match ($v_0) { 1 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_Monday()), 2 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_Tuesday()), 3 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_Wednesday()), 4 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_Thursday()), 5 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_Friday()), 6 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_Saturday()), 7 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_Sunday()), default => new \Data\Maybe\Data_Maybe_Nothing() };
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "fromEnum" => function($v_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v_0 instanceof \Data\Date\Component\Data_Date_Component_Monday) {
$__t1 = 1;
goto end_branch_1;;
};
  if ($v_0 instanceof \Data\Date\Component\Data_Date_Component_Tuesday) {
$__t1 = 2;
goto end_branch_1;;
};
  if ($v_0 instanceof \Data\Date\Component\Data_Date_Component_Wednesday) {
$__t1 = 3;
goto end_branch_1;;
};
  if ($v_0 instanceof \Data\Date\Component\Data_Date_Component_Thursday) {
$__t1 = 4;
goto end_branch_1;;
};
  if ($v_0 instanceof \Data\Date\Component\Data_Date_Component_Friday) {
$__t1 = 5;
goto end_branch_1;;
};
  if ($v_0 instanceof \Data\Date\Component\Data_Date_Component_Saturday) {
$__t1 = 6;
goto end_branch_1;;
};
  if ($v_0 instanceof \Data\Date\Component\Data_Date_Component_Sunday) {
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
}, "Bounded0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Date_Component_boundedWeekday'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Enum1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Date_Component_enumWeekday'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Date_Component_enumWeekday
$GLOBALS['Data_Date_Component_enumWeekday'] = (object)["succ" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Date_Component_boundedEnumWeekday'])->{'toEnum'}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_0) {
  $__num = \func_num_args();
  $__res = ($v_0 + 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($GLOBALS['Data_Date_Component_boundedEnumWeekday'])->{'fromEnum'})), "pred" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Date_Component_boundedEnumWeekday'])->{'toEnum'}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_0) {
  $__num = \func_num_args();
  $__res = ($v_0 - 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($GLOBALS['Data_Date_Component_boundedEnumWeekday'])->{'fromEnum'})), "Ord0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Date_Component_ordWeekday'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Date_Component_boundedEnumMonth
$GLOBALS['Data_Date_Component_boundedEnumMonth'] = (object)["cardinality" => 12, "toEnum" => function($v_0) {
  $__num = \func_num_args();
  $__res = match ($v_0) { 1 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_January()), 2 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_February()), 3 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_March()), 4 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_April()), 5 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_May()), 6 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_June()), 7 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_July()), 8 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_August()), 9 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_September()), 10 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_October()), 11 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_November()), 12 => new \Data\Maybe\Data_Maybe_Just(new \Data\Date\Component\Data_Date_Component_December()), default => new \Data\Maybe\Data_Maybe_Nothing() };
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "fromEnum" => function($v_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v_0 instanceof \Data\Date\Component\Data_Date_Component_January) {
$__t1 = 1;
goto end_branch_1;;
};
  if ($v_0 instanceof \Data\Date\Component\Data_Date_Component_February) {
$__t1 = 2;
goto end_branch_1;;
};
  if ($v_0 instanceof \Data\Date\Component\Data_Date_Component_March) {
$__t1 = 3;
goto end_branch_1;;
};
  if ($v_0 instanceof \Data\Date\Component\Data_Date_Component_April) {
$__t1 = 4;
goto end_branch_1;;
};
  if ($v_0 instanceof \Data\Date\Component\Data_Date_Component_May) {
$__t1 = 5;
goto end_branch_1;;
};
  if ($v_0 instanceof \Data\Date\Component\Data_Date_Component_June) {
$__t1 = 6;
goto end_branch_1;;
};
  if ($v_0 instanceof \Data\Date\Component\Data_Date_Component_July) {
$__t1 = 7;
goto end_branch_1;;
};
  if ($v_0 instanceof \Data\Date\Component\Data_Date_Component_August) {
$__t1 = 8;
goto end_branch_1;;
};
  if ($v_0 instanceof \Data\Date\Component\Data_Date_Component_September) {
$__t1 = 9;
goto end_branch_1;;
};
  if ($v_0 instanceof \Data\Date\Component\Data_Date_Component_October) {
$__t1 = 10;
goto end_branch_1;;
};
  if ($v_0 instanceof \Data\Date\Component\Data_Date_Component_November) {
$__t1 = 11;
goto end_branch_1;;
};
  if ($v_0 instanceof \Data\Date\Component\Data_Date_Component_December) {
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
}, "Bounded0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Date_Component_boundedMonth'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Enum1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Date_Component_enumMonth'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Date_Component_enumMonth
$GLOBALS['Data_Date_Component_enumMonth'] = (object)["succ" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Date_Component_boundedEnumMonth'])->{'toEnum'}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_0) {
  $__num = \func_num_args();
  $__res = ($v_0 + 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($GLOBALS['Data_Date_Component_boundedEnumMonth'])->{'fromEnum'})), "pred" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Date_Component_boundedEnumMonth'])->{'toEnum'}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_0) {
  $__num = \func_num_args();
  $__res = ($v_0 - 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($GLOBALS['Data_Date_Component_boundedEnumMonth'])->{'fromEnum'})), "Ord0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Date_Component_ordMonth'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Date_Component_boundedDay
$GLOBALS['Data_Date_Component_boundedDay'] = (object)["bottom" => 1, "top" => 31, "Ord0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Date_Component_ordDay'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Date_Component_boundedEnumDay
$GLOBALS['Data_Date_Component_boundedEnumDay'] = (object)["cardinality" => 31, "toEnum" => function($n_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if (((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'conj'})((($GLOBALS['Data_Date_Component_greaterThanOrEq'])($n_0))(1)))((($GLOBALS['Data_Date_Component_lessThanOrEq'])($n_0))(31))) {
$__t0 = new \Data\Maybe\Data_Maybe_Just($n_0);
goto end_branch_0;;
};
  $__t0 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "fromEnum" => function($v_0) {
  $__num = \func_num_args();
  $__res = $v_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bounded0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Date_Component_boundedDay'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Enum1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Date_Component_enumDay'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Date_Component_enumDay
$GLOBALS['Data_Date_Component_enumDay'] = (object)["succ" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Date_Component_boundedEnumDay'])->{'toEnum'}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_0) {
  $__num = \func_num_args();
  $__res = ($v_0 + 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($GLOBALS['Data_Date_Component_boundedEnumDay'])->{'fromEnum'})), "pred" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Date_Component_boundedEnumDay'])->{'toEnum'}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_0) {
  $__num = \func_num_args();
  $__res = ($v_0 - 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($GLOBALS['Data_Date_Component_boundedEnumDay'])->{'fromEnum'})), "Ord0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Date_Component_ordDay'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

