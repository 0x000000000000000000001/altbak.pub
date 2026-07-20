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
\PhpursThunks::$thunks['Data_Date_Component_showYear'] = function() { $v = (object)["show" => function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))("(Year "))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))((($GLOBALS['Data_Show_showIntImpl'] ?? \PhpursThunks::eval('Data_Show_showIntImpl')))($v_0)))(")"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Date_Component_showWeekday'] = function() { $v = (object)["show" => function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v_0) && (($v_0)->tag === "Monday"))) {
$__t0 = "Monday";
} else {
if ((is_object($v_0) && (($v_0)->tag === "Tuesday"))) {
$__t0 = "Tuesday";
} else {
if ((is_object($v_0) && (($v_0)->tag === "Wednesday"))) {
$__t0 = "Wednesday";
} else {
if ((is_object($v_0) && (($v_0)->tag === "Thursday"))) {
$__t0 = "Thursday";
} else {
if ((is_object($v_0) && (($v_0)->tag === "Friday"))) {
$__t0 = "Friday";
} else {
if ((is_object($v_0) && (($v_0)->tag === "Saturday"))) {
$__t0 = "Saturday";
} else {
if ((is_object($v_0) && (($v_0)->tag === "Sunday"))) {
$__t0 = "Sunday";
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
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Date_Component_showMonth'] = function() { $v = (object)["show" => function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v_0) && (($v_0)->tag === "January"))) {
$__t0 = "January";
} else {
if ((is_object($v_0) && (($v_0)->tag === "February"))) {
$__t0 = "February";
} else {
if ((is_object($v_0) && (($v_0)->tag === "March"))) {
$__t0 = "March";
} else {
if ((is_object($v_0) && (($v_0)->tag === "April"))) {
$__t0 = "April";
} else {
if ((is_object($v_0) && (($v_0)->tag === "May"))) {
$__t0 = "May";
} else {
if ((is_object($v_0) && (($v_0)->tag === "June"))) {
$__t0 = "June";
} else {
if ((is_object($v_0) && (($v_0)->tag === "July"))) {
$__t0 = "July";
} else {
if ((is_object($v_0) && (($v_0)->tag === "August"))) {
$__t0 = "August";
} else {
if ((is_object($v_0) && (($v_0)->tag === "September"))) {
$__t0 = "September";
} else {
if ((is_object($v_0) && (($v_0)->tag === "October"))) {
$__t0 = "October";
} else {
if ((is_object($v_0) && (($v_0)->tag === "November"))) {
$__t0 = "November";
} else {
if ((is_object($v_0) && (($v_0)->tag === "December"))) {
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
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Date_Component_showDay'] = function() { $v = (object)["show" => function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))("(Day "))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))((($GLOBALS['Data_Show_showIntImpl'] ?? \PhpursThunks::eval('Data_Show_showIntImpl')))($v_0)))(")"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Date_Component_ordYear'] = function() { $v = ($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')); return $v; };
\PhpursThunks::$thunks['Data_Date_Component_ordDay'] = function() { $v = ($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')); return $v; };
\PhpursThunks::$thunks['Data_Date_Component_eqYear'] = function() { $v = ($GLOBALS['Data_Eq_eqInt'] ?? \PhpursThunks::eval('Data_Eq_eqInt')); return $v; };
\PhpursThunks::$thunks['Data_Date_Component_eqWeekday'] = function() { $v = (object)["eq" => (function() {
  $__fn = function($x_0, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($x_0) && (($x_0)->tag === "Monday"))) {
$__t0 = (is_object($y_1) && (($y_1)->tag === "Monday"));
} else {
if ((is_object($x_0) && (($x_0)->tag === "Tuesday"))) {
$__t0 = (is_object($y_1) && (($y_1)->tag === "Tuesday"));
} else {
if ((is_object($x_0) && (($x_0)->tag === "Wednesday"))) {
$__t0 = (is_object($y_1) && (($y_1)->tag === "Wednesday"));
} else {
if ((is_object($x_0) && (($x_0)->tag === "Thursday"))) {
$__t0 = (is_object($y_1) && (($y_1)->tag === "Thursday"));
} else {
if ((is_object($x_0) && (($x_0)->tag === "Friday"))) {
$__t0 = (is_object($y_1) && (($y_1)->tag === "Friday"));
} else {
if ((is_object($x_0) && (($x_0)->tag === "Saturday"))) {
$__t0 = (is_object($y_1) && (($y_1)->tag === "Saturday"));
} else {
$__t0 = ((is_object($x_0) && (($x_0)->tag === "Sunday")) && (is_object($y_1) && (($y_1)->tag === "Sunday")));
};
};
};
};
};
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Data_Date_Component_ordWeekday'] = function() { $v = (object)["compare" => (function() {
  $__fn = function($x_0, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($x_0) && (($x_0)->tag === "Monday"))) {
if ((is_object($y_1) && (($y_1)->tag === "Monday"))) {
$__t1 = new Phpurs_Data0("EQ");
} else {
$__t1 = new Phpurs_Data0("LT");
};
$__t0 = $__t1;
} else {
if ((is_object($y_1) && (($y_1)->tag === "Monday"))) {
$__t0 = new Phpurs_Data0("GT");
} else {
if ((is_object($x_0) && (($x_0)->tag === "Tuesday"))) {
if ((is_object($y_1) && (($y_1)->tag === "Tuesday"))) {
$__t2 = new Phpurs_Data0("EQ");
} else {
$__t2 = new Phpurs_Data0("LT");
};
$__t0 = $__t2;
} else {
if ((is_object($y_1) && (($y_1)->tag === "Tuesday"))) {
$__t0 = new Phpurs_Data0("GT");
} else {
if ((is_object($x_0) && (($x_0)->tag === "Wednesday"))) {
if ((is_object($y_1) && (($y_1)->tag === "Wednesday"))) {
$__t3 = new Phpurs_Data0("EQ");
} else {
$__t3 = new Phpurs_Data0("LT");
};
$__t0 = $__t3;
} else {
if ((is_object($y_1) && (($y_1)->tag === "Wednesday"))) {
$__t0 = new Phpurs_Data0("GT");
} else {
if ((is_object($x_0) && (($x_0)->tag === "Thursday"))) {
if ((is_object($y_1) && (($y_1)->tag === "Thursday"))) {
$__t4 = new Phpurs_Data0("EQ");
} else {
$__t4 = new Phpurs_Data0("LT");
};
$__t0 = $__t4;
} else {
if ((is_object($y_1) && (($y_1)->tag === "Thursday"))) {
$__t0 = new Phpurs_Data0("GT");
} else {
if ((is_object($x_0) && (($x_0)->tag === "Friday"))) {
if ((is_object($y_1) && (($y_1)->tag === "Friday"))) {
$__t5 = new Phpurs_Data0("EQ");
} else {
$__t5 = new Phpurs_Data0("LT");
};
$__t0 = $__t5;
} else {
if ((is_object($y_1) && (($y_1)->tag === "Friday"))) {
$__t0 = new Phpurs_Data0("GT");
} else {
if ((is_object($x_0) && (($x_0)->tag === "Saturday"))) {
if ((is_object($y_1) && (($y_1)->tag === "Saturday"))) {
$__t6 = new Phpurs_Data0("EQ");
} else {
$__t6 = new Phpurs_Data0("LT");
};
$__t0 = $__t6;
} else {
if ((is_object($y_1) && (($y_1)->tag === "Saturday"))) {
$__t0 = new Phpurs_Data0("GT");
} else {
if (((is_object($x_0) && (($x_0)->tag === "Sunday")) && (is_object($y_1) && (($y_1)->tag === "Sunday")))) {
$__t0 = new Phpurs_Data0("EQ");
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
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_Date_Component_eqWeekday'] ?? \PhpursThunks::eval('Data_Date_Component_eqWeekday'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Date_Component_eqMonth'] = function() { $v = (object)["eq" => (function() {
  $__fn = function($x_0, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($x_0) && (($x_0)->tag === "January"))) {
$__t0 = (is_object($y_1) && (($y_1)->tag === "January"));
} else {
if ((is_object($x_0) && (($x_0)->tag === "February"))) {
$__t0 = (is_object($y_1) && (($y_1)->tag === "February"));
} else {
if ((is_object($x_0) && (($x_0)->tag === "March"))) {
$__t0 = (is_object($y_1) && (($y_1)->tag === "March"));
} else {
if ((is_object($x_0) && (($x_0)->tag === "April"))) {
$__t0 = (is_object($y_1) && (($y_1)->tag === "April"));
} else {
if ((is_object($x_0) && (($x_0)->tag === "May"))) {
$__t0 = (is_object($y_1) && (($y_1)->tag === "May"));
} else {
if ((is_object($x_0) && (($x_0)->tag === "June"))) {
$__t0 = (is_object($y_1) && (($y_1)->tag === "June"));
} else {
if ((is_object($x_0) && (($x_0)->tag === "July"))) {
$__t0 = (is_object($y_1) && (($y_1)->tag === "July"));
} else {
if ((is_object($x_0) && (($x_0)->tag === "August"))) {
$__t0 = (is_object($y_1) && (($y_1)->tag === "August"));
} else {
if ((is_object($x_0) && (($x_0)->tag === "September"))) {
$__t0 = (is_object($y_1) && (($y_1)->tag === "September"));
} else {
if ((is_object($x_0) && (($x_0)->tag === "October"))) {
$__t0 = (is_object($y_1) && (($y_1)->tag === "October"));
} else {
if ((is_object($x_0) && (($x_0)->tag === "November"))) {
$__t0 = (is_object($y_1) && (($y_1)->tag === "November"));
} else {
$__t0 = ((is_object($x_0) && (($x_0)->tag === "December")) && (is_object($y_1) && (($y_1)->tag === "December")));
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
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Data_Date_Component_ordMonth'] = function() { $v = (object)["compare" => (function() {
  $__fn = function($x_0, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($x_0) && (($x_0)->tag === "January"))) {
if ((is_object($y_1) && (($y_1)->tag === "January"))) {
$__t1 = new Phpurs_Data0("EQ");
} else {
$__t1 = new Phpurs_Data0("LT");
};
$__t0 = $__t1;
} else {
if ((is_object($y_1) && (($y_1)->tag === "January"))) {
$__t0 = new Phpurs_Data0("GT");
} else {
if ((is_object($x_0) && (($x_0)->tag === "February"))) {
if ((is_object($y_1) && (($y_1)->tag === "February"))) {
$__t2 = new Phpurs_Data0("EQ");
} else {
$__t2 = new Phpurs_Data0("LT");
};
$__t0 = $__t2;
} else {
if ((is_object($y_1) && (($y_1)->tag === "February"))) {
$__t0 = new Phpurs_Data0("GT");
} else {
if ((is_object($x_0) && (($x_0)->tag === "March"))) {
if ((is_object($y_1) && (($y_1)->tag === "March"))) {
$__t3 = new Phpurs_Data0("EQ");
} else {
$__t3 = new Phpurs_Data0("LT");
};
$__t0 = $__t3;
} else {
if ((is_object($y_1) && (($y_1)->tag === "March"))) {
$__t0 = new Phpurs_Data0("GT");
} else {
if ((is_object($x_0) && (($x_0)->tag === "April"))) {
if ((is_object($y_1) && (($y_1)->tag === "April"))) {
$__t4 = new Phpurs_Data0("EQ");
} else {
$__t4 = new Phpurs_Data0("LT");
};
$__t0 = $__t4;
} else {
if ((is_object($y_1) && (($y_1)->tag === "April"))) {
$__t0 = new Phpurs_Data0("GT");
} else {
if ((is_object($x_0) && (($x_0)->tag === "May"))) {
if ((is_object($y_1) && (($y_1)->tag === "May"))) {
$__t5 = new Phpurs_Data0("EQ");
} else {
$__t5 = new Phpurs_Data0("LT");
};
$__t0 = $__t5;
} else {
if ((is_object($y_1) && (($y_1)->tag === "May"))) {
$__t0 = new Phpurs_Data0("GT");
} else {
if ((is_object($x_0) && (($x_0)->tag === "June"))) {
if ((is_object($y_1) && (($y_1)->tag === "June"))) {
$__t6 = new Phpurs_Data0("EQ");
} else {
$__t6 = new Phpurs_Data0("LT");
};
$__t0 = $__t6;
} else {
if ((is_object($y_1) && (($y_1)->tag === "June"))) {
$__t0 = new Phpurs_Data0("GT");
} else {
if ((is_object($x_0) && (($x_0)->tag === "July"))) {
if ((is_object($y_1) && (($y_1)->tag === "July"))) {
$__t7 = new Phpurs_Data0("EQ");
} else {
$__t7 = new Phpurs_Data0("LT");
};
$__t0 = $__t7;
} else {
if ((is_object($y_1) && (($y_1)->tag === "July"))) {
$__t0 = new Phpurs_Data0("GT");
} else {
if ((is_object($x_0) && (($x_0)->tag === "August"))) {
if ((is_object($y_1) && (($y_1)->tag === "August"))) {
$__t8 = new Phpurs_Data0("EQ");
} else {
$__t8 = new Phpurs_Data0("LT");
};
$__t0 = $__t8;
} else {
if ((is_object($y_1) && (($y_1)->tag === "August"))) {
$__t0 = new Phpurs_Data0("GT");
} else {
if ((is_object($x_0) && (($x_0)->tag === "September"))) {
if ((is_object($y_1) && (($y_1)->tag === "September"))) {
$__t9 = new Phpurs_Data0("EQ");
} else {
$__t9 = new Phpurs_Data0("LT");
};
$__t0 = $__t9;
} else {
if ((is_object($y_1) && (($y_1)->tag === "September"))) {
$__t0 = new Phpurs_Data0("GT");
} else {
if ((is_object($x_0) && (($x_0)->tag === "October"))) {
if ((is_object($y_1) && (($y_1)->tag === "October"))) {
$__t10 = new Phpurs_Data0("EQ");
} else {
$__t10 = new Phpurs_Data0("LT");
};
$__t0 = $__t10;
} else {
if ((is_object($y_1) && (($y_1)->tag === "October"))) {
$__t0 = new Phpurs_Data0("GT");
} else {
if ((is_object($x_0) && (($x_0)->tag === "November"))) {
if ((is_object($y_1) && (($y_1)->tag === "November"))) {
$__t11 = new Phpurs_Data0("EQ");
} else {
$__t11 = new Phpurs_Data0("LT");
};
$__t0 = $__t11;
} else {
if ((is_object($y_1) && (($y_1)->tag === "November"))) {
$__t0 = new Phpurs_Data0("GT");
} else {
if (((is_object($x_0) && (($x_0)->tag === "December")) && (is_object($y_1) && (($y_1)->tag === "December")))) {
$__t0 = new Phpurs_Data0("EQ");
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
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_Date_Component_eqMonth'] ?? \PhpursThunks::eval('Data_Date_Component_eqMonth'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Date_Component_eqDay'] = function() { $v = ($GLOBALS['Data_Eq_eqInt'] ?? \PhpursThunks::eval('Data_Eq_eqInt')); return $v; };
\PhpursThunks::$thunks['Data_Date_Component_boundedYear'] = function() { $v = (object)["bottom" => ((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))(0))(271820), "top" => 275759, "Ord0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Date_Component_boundedWeekday'] = function() { $v = (object)["bottom" => new Phpurs_Data0("Monday"), "top" => new Phpurs_Data0("Sunday"), "Ord0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_Date_Component_ordWeekday'] ?? \PhpursThunks::eval('Data_Date_Component_ordWeekday'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Date_Component_boundedMonth'] = function() { $v = (object)["bottom" => new Phpurs_Data0("January"), "top" => new Phpurs_Data0("December"), "Ord0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_Date_Component_ordMonth'] ?? \PhpursThunks::eval('Data_Date_Component_ordMonth'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Date_Component_boundedEnumYear'] = function() { $v = (object)["cardinality" => 547580, "toEnum" => function($n_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Date_Component_boundedEnumYear","Data_Date_Component_enumYear"];
  if (((($GLOBALS['Data_HeytingAlgebra_boolConj'] ?? \PhpursThunks::eval('Data_HeytingAlgebra_boolConj')))(( ! (is_object((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($n_0))(((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))(0))(271820))) && (((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($n_0))(((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))(0))(271820)))->tag === "LT")))))(( ! (is_object((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($n_0))(275759)) && (((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($n_0))(275759))->tag === "GT"))))) {
$__t0 = new Phpurs_Data1("Just", $n_0);
} else {
$__t0 = new Phpurs_Data0("Nothing");
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "fromEnum" => function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Date_Component_boundedEnumYear","Data_Date_Component_enumYear"];
  $__res = $v_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bounded0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Date_Component_boundedEnumYear","Data_Date_Component_enumYear"];
  $__res = ($GLOBALS['Data_Date_Component_boundedYear'] ?? \PhpursThunks::eval('Data_Date_Component_boundedYear'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Enum1" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Date_Component_boundedEnumYear","Data_Date_Component_enumYear"];
  $__res = ($GLOBALS['Data_Date_Component_enumYear'] ?? \PhpursThunks::eval('Data_Date_Component_enumYear'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Date_Component_enumYear'] = function() { $v = (object)["succ" => ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_Date_Component_boundedEnumYear'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumYear')))->toEnum))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Date_Component_boundedEnumYear","Data_Date_Component_enumYear"];
  $__res = ((($GLOBALS['Data_Semiring_intAdd'] ?? \PhpursThunks::eval('Data_Semiring_intAdd')))($v_0))(1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Data_Date_Component_boundedEnumYear'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumYear')))->fromEnum)), "pred" => ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_Date_Component_boundedEnumYear'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumYear')))->toEnum))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Date_Component_boundedEnumYear","Data_Date_Component_enumYear"];
  $__res = ((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))($v_0))(1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Data_Date_Component_boundedEnumYear'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumYear')))->fromEnum)), "Ord0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Date_Component_boundedEnumYear","Data_Date_Component_enumYear"];
  $__res = ($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Date_Component_boundedEnumWeekday'] = function() { $v = (object)["cardinality" => 7, "toEnum" => function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Date_Component_boundedEnumWeekday","Data_Date_Component_enumWeekday"];
  switch ($v_0) {
case 1:
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data0("Monday"));
break;
case 2:
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data0("Tuesday"));
break;
case 3:
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data0("Wednesday"));
break;
case 4:
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data0("Thursday"));
break;
case 5:
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data0("Friday"));
break;
case 6:
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data0("Saturday"));
break;
case 7:
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data0("Sunday"));
break;
default:
$__t0 = new Phpurs_Data0("Nothing");
break;
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "fromEnum" => function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Date_Component_boundedEnumWeekday","Data_Date_Component_enumWeekday"];
  if ((is_object($v_0) && (($v_0)->tag === "Monday"))) {
$__t1 = 1;
} else {
if ((is_object($v_0) && (($v_0)->tag === "Tuesday"))) {
$__t1 = 2;
} else {
if ((is_object($v_0) && (($v_0)->tag === "Wednesday"))) {
$__t1 = 3;
} else {
if ((is_object($v_0) && (($v_0)->tag === "Thursday"))) {
$__t1 = 4;
} else {
if ((is_object($v_0) && (($v_0)->tag === "Friday"))) {
$__t1 = 5;
} else {
if ((is_object($v_0) && (($v_0)->tag === "Saturday"))) {
$__t1 = 6;
} else {
if ((is_object($v_0) && (($v_0)->tag === "Sunday"))) {
$__t1 = 7;
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
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bounded0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Date_Component_boundedEnumWeekday","Data_Date_Component_enumWeekday"];
  $__res = ($GLOBALS['Data_Date_Component_boundedWeekday'] ?? \PhpursThunks::eval('Data_Date_Component_boundedWeekday'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Enum1" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Date_Component_boundedEnumWeekday","Data_Date_Component_enumWeekday"];
  $__res = ($GLOBALS['Data_Date_Component_enumWeekday'] ?? \PhpursThunks::eval('Data_Date_Component_enumWeekday'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Date_Component_enumWeekday'] = function() { $v = (object)["succ" => ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_Date_Component_boundedEnumWeekday'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumWeekday')))->toEnum))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Date_Component_boundedEnumWeekday","Data_Date_Component_enumWeekday"];
  $__res = ((($GLOBALS['Data_Semiring_intAdd'] ?? \PhpursThunks::eval('Data_Semiring_intAdd')))($v_0))(1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Data_Date_Component_boundedEnumWeekday'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumWeekday')))->fromEnum)), "pred" => ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_Date_Component_boundedEnumWeekday'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumWeekday')))->toEnum))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Date_Component_boundedEnumWeekday","Data_Date_Component_enumWeekday"];
  $__res = ((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))($v_0))(1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Data_Date_Component_boundedEnumWeekday'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumWeekday')))->fromEnum)), "Ord0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Date_Component_boundedEnumWeekday","Data_Date_Component_enumWeekday"];
  $__res = ($GLOBALS['Data_Date_Component_ordWeekday'] ?? \PhpursThunks::eval('Data_Date_Component_ordWeekday'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Date_Component_boundedEnumMonth'] = function() { $v = (object)["cardinality" => 12, "toEnum" => function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Date_Component_boundedEnumMonth","Data_Date_Component_enumMonth"];
  switch ($v_0) {
case 1:
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data0("January"));
break;
case 2:
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data0("February"));
break;
case 3:
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data0("March"));
break;
case 4:
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data0("April"));
break;
case 5:
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data0("May"));
break;
case 6:
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data0("June"));
break;
case 7:
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data0("July"));
break;
case 8:
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data0("August"));
break;
case 9:
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data0("September"));
break;
case 10:
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data0("October"));
break;
case 11:
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data0("November"));
break;
case 12:
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data0("December"));
break;
default:
$__t0 = new Phpurs_Data0("Nothing");
break;
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "fromEnum" => function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Date_Component_boundedEnumMonth","Data_Date_Component_enumMonth"];
  if ((is_object($v_0) && (($v_0)->tag === "January"))) {
$__t1 = 1;
} else {
if ((is_object($v_0) && (($v_0)->tag === "February"))) {
$__t1 = 2;
} else {
if ((is_object($v_0) && (($v_0)->tag === "March"))) {
$__t1 = 3;
} else {
if ((is_object($v_0) && (($v_0)->tag === "April"))) {
$__t1 = 4;
} else {
if ((is_object($v_0) && (($v_0)->tag === "May"))) {
$__t1 = 5;
} else {
if ((is_object($v_0) && (($v_0)->tag === "June"))) {
$__t1 = 6;
} else {
if ((is_object($v_0) && (($v_0)->tag === "July"))) {
$__t1 = 7;
} else {
if ((is_object($v_0) && (($v_0)->tag === "August"))) {
$__t1 = 8;
} else {
if ((is_object($v_0) && (($v_0)->tag === "September"))) {
$__t1 = 9;
} else {
if ((is_object($v_0) && (($v_0)->tag === "October"))) {
$__t1 = 10;
} else {
if ((is_object($v_0) && (($v_0)->tag === "November"))) {
$__t1 = 11;
} else {
if ((is_object($v_0) && (($v_0)->tag === "December"))) {
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
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bounded0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Date_Component_boundedEnumMonth","Data_Date_Component_enumMonth"];
  $__res = ($GLOBALS['Data_Date_Component_boundedMonth'] ?? \PhpursThunks::eval('Data_Date_Component_boundedMonth'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Enum1" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Date_Component_boundedEnumMonth","Data_Date_Component_enumMonth"];
  $__res = ($GLOBALS['Data_Date_Component_enumMonth'] ?? \PhpursThunks::eval('Data_Date_Component_enumMonth'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Date_Component_enumMonth'] = function() { $v = (object)["succ" => ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_Date_Component_boundedEnumMonth'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumMonth')))->toEnum))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Date_Component_boundedEnumMonth","Data_Date_Component_enumMonth"];
  $__res = ((($GLOBALS['Data_Semiring_intAdd'] ?? \PhpursThunks::eval('Data_Semiring_intAdd')))($v_0))(1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Data_Date_Component_boundedEnumMonth'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumMonth')))->fromEnum)), "pred" => ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_Date_Component_boundedEnumMonth'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumMonth')))->toEnum))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Date_Component_boundedEnumMonth","Data_Date_Component_enumMonth"];
  $__res = ((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))($v_0))(1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Data_Date_Component_boundedEnumMonth'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumMonth')))->fromEnum)), "Ord0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Date_Component_boundedEnumMonth","Data_Date_Component_enumMonth"];
  $__res = ($GLOBALS['Data_Date_Component_ordMonth'] ?? \PhpursThunks::eval('Data_Date_Component_ordMonth'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Date_Component_boundedDay'] = function() { $v = (object)["bottom" => 1, "top" => 31, "Ord0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Date_Component_boundedEnumDay'] = function() { $v = (object)["cardinality" => 31, "toEnum" => function($n_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Date_Component_boundedEnumDay","Data_Date_Component_enumDay"];
  if (((($GLOBALS['Data_HeytingAlgebra_boolConj'] ?? \PhpursThunks::eval('Data_HeytingAlgebra_boolConj')))(( ! (is_object((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($n_0))(1)) && (((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($n_0))(1))->tag === "LT")))))(( ! (is_object((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($n_0))(31)) && (((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($n_0))(31))->tag === "GT"))))) {
$__t0 = new Phpurs_Data1("Just", $n_0);
} else {
$__t0 = new Phpurs_Data0("Nothing");
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "fromEnum" => function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Date_Component_boundedEnumDay","Data_Date_Component_enumDay"];
  $__res = $v_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bounded0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Date_Component_boundedEnumDay","Data_Date_Component_enumDay"];
  $__res = ($GLOBALS['Data_Date_Component_boundedDay'] ?? \PhpursThunks::eval('Data_Date_Component_boundedDay'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Enum1" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Date_Component_boundedEnumDay","Data_Date_Component_enumDay"];
  $__res = ($GLOBALS['Data_Date_Component_enumDay'] ?? \PhpursThunks::eval('Data_Date_Component_enumDay'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Date_Component_enumDay'] = function() { $v = (object)["succ" => ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_Date_Component_boundedEnumDay'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumDay')))->toEnum))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Date_Component_boundedEnumDay","Data_Date_Component_enumDay"];
  $__res = ((($GLOBALS['Data_Semiring_intAdd'] ?? \PhpursThunks::eval('Data_Semiring_intAdd')))($v_0))(1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Data_Date_Component_boundedEnumDay'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumDay')))->fromEnum)), "pred" => ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_Date_Component_boundedEnumDay'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumDay')))->toEnum))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Date_Component_boundedEnumDay","Data_Date_Component_enumDay"];
  $__res = ((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))($v_0))(1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Data_Date_Component_boundedEnumDay'] ?? \PhpursThunks::eval('Data_Date_Component_boundedEnumDay')))->fromEnum)), "Ord0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Date_Component_boundedEnumDay","Data_Date_Component_enumDay"];
  $__res = ($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };













































