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
$GLOBALS['Data_DateTime_adjustImpl'] = $ffi_Data_DateTime['adjustImpl'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_DateTime_calcDiff'] = $ffi_Data_DateTime['calcDiff'] ?? new class { public function __invoke(...$args) { return $this; } };


// Data_DateTime_DateTime
$GLOBALS['Data_DateTime_DateTime'] = (function() {
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
})();

// Data_DateTime_toRecord
$GLOBALS['Data_DateTime_toRecord'] = function($v_0 = null) {
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
  $__res = ["year" => (($v_0)->{'value0'})->{'value0'}, "month" => $__t0, "day" => (($v_0)->{'value0'})->{'value2'}, "hour" => (($v_0)->{'value1'})->{'value0'}, "minute" => (($v_0)->{'value1'})->{'value1'}, "second" => (($v_0)->{'value1'})->{'value2'}, "millisecond" => (($v_0)->{'value1'})->{'value3'}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_DateTime_time
$GLOBALS['Data_DateTime_time'] = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ($v_0)->{'value1'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_DateTime_showDateTime
$GLOBALS['Data_DateTime_showDateTime'] = ["show" => function($v_0 = null) {
  $__num = \func_num_args();
  $__res = (((("(DateTime " . (($GLOBALS['Data_Date_showDate'])['show'])(($v_0)->{'value0'})) . " ") . (($GLOBALS['Data_Time_showTime'])['show'])(($v_0)->{'value1'})) . ")");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_DateTime_modifyTimeF
$GLOBALS['Data_DateTime_modifyTimeF'] = (function() {
  $__fn = function($dictFunctor_0 = null, $f_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictFunctor_0)['map'])(($GLOBALS['Data_DateTime_DateTime'])(($v_2)->{'value0'})))(($f_1)(($v_2)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_DateTime_modifyTime
$GLOBALS['Data_DateTime_modifyTime'] = (function() {
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
})();

// Data_DateTime_modifyDateF
$GLOBALS['Data_DateTime_modifyDateF'] = (function() {
  $__fn = function($dictFunctor_0 = null, $f_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__local_var_3_0 = ($v_2)->{'value1'};
  $__res = ((($dictFunctor_0)['map'])(function($a_4 = null) use ($__local_var_3_0) {
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
})();

// Data_DateTime_modifyDate
$GLOBALS['Data_DateTime_modifyDate'] = (function() {
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
})();

// Data_DateTime_eqDateTime
$GLOBALS['Data_DateTime_eqDateTime'] = ["eq" => (function() {
  $__fn = function($x_0 = null, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($GLOBALS['Data_Date_eqDate'])['eq'])(($x_0)->{'value0'}))(($y_1)->{'value0'}) && (((((($x_0)->{'value1'})->{'value0'} === (($y_1)->{'value1'})->{'value0'}) && ((($x_0)->{'value1'})->{'value1'} === (($y_1)->{'value1'})->{'value1'})) && ((($x_0)->{'value1'})->{'value2'} === (($y_1)->{'value1'})->{'value2'})) && ((($x_0)->{'value1'})->{'value3'} === (($y_1)->{'value1'})->{'value3'})));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];

// Data_DateTime_ordDateTime
$GLOBALS['Data_DateTime_ordDateTime'] = ["compare" => (function() {
  $__fn = function($x_0 = null, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v_2_0 = ((($GLOBALS['Data_Date_ordDate'])['compare'])(($x_0)->{'value0'}))(($y_1)->{'value0'});
  $__t1 = null;;
  if ((is_object($v_2_0) && (($v_2_0)->{'tag'} === "LT"))) {
$__t1 = new Phpurs_Data0("LT");
goto end_branch_1;;
};
  if ((is_object($v_2_0) && (($v_2_0)->{'tag'} === "GT"))) {
$__t1 = new Phpurs_Data0("GT");
goto end_branch_1;;
};
  $__t1 = ((($GLOBALS['Data_Time_ordTime'])['compare'])(($x_0)->{'value1'}))(($y_1)->{'value1'});
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_DateTime_eqDateTime'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_DateTime_diff
$GLOBALS['Data_DateTime_diff'] = (function() {
  $__fn = function($dictDuration_0 = null, $dt1_1 = null, $dt2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($dictDuration_0)['toDuration'])(($GLOBALS['Data_DateTime_calcDiff'])(($GLOBALS['Data_DateTime_toRecord'])($dt1_1), ($GLOBALS['Data_DateTime_toRecord'])($dt2_2)));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_DateTime_date
$GLOBALS['Data_DateTime_date'] = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ($v_0)->{'value0'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_DateTime_boundedDateTime
$GLOBALS['Data_DateTime_boundedDateTime'] = ["bottom" => new Phpurs_Data2("DateTime", new Phpurs_Data3("Date", -271820, new Phpurs_Data0("January"), 1), new Phpurs_Data4("Time", 0, 0, 0, 0)), "top" => new Phpurs_Data2("DateTime", new Phpurs_Data3("Date", 275759, new Phpurs_Data0("December"), 31), new Phpurs_Data4("Time", 23, 59, 59, 999)), "Ord0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_DateTime_ordDateTime'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_DateTime_adjust
$GLOBALS['Data_DateTime_adjust'] = (function() {
  $__fn = function($dictDuration_0 = null, $d_1 = null, $dt_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__local_var_3_0 = (((($GLOBALS['Data_DateTime_adjustImpl'])($GLOBALS['Data_Maybe_Just']))(new Phpurs_Data0("Nothing")))((($dictDuration_0)['fromDuration'])($d_1)))(($GLOBALS['Data_DateTime_toRecord'])($dt_2));
  $__t1 = null;;
  if ((is_object($__local_var_3_0) && (($__local_var_3_0)->{'tag'} === "Just"))) {
$__t2 = null;;
if ((((($__local_var_3_0)->{'value0'})['year'] >= -271820) && ((($__local_var_3_0)->{'value0'})['year'] <= 275759))) {
$__t2 = new Phpurs_Data1("Just", (($__local_var_3_0)->{'value0'})['year']);
goto end_branch_2;;
};
$__t2 = new Phpurs_Data0("Nothing");
end_branch_2:;
$__local_var_4_2 = $__t2;
$__t4 = null;;
if ((is_object($__local_var_4_2) && (($__local_var_4_2)->{'tag'} === "Just"))) {
$__t4 = new Phpurs_Data1("Just", ($GLOBALS['Data_Date_exactDate'])(($__local_var_4_2)->{'value0'}));
goto end_branch_4;;
};
$__t4 = new Phpurs_Data0("Nothing");
end_branch_4:;
$__local_var_5_4 = $__t4;
$__local_var_6_6 = match ((($__local_var_3_0)->{'value0'})['month']) { 1 => new Phpurs_Data1("Just", new Phpurs_Data0("January")), 2 => new Phpurs_Data1("Just", new Phpurs_Data0("February")), 3 => new Phpurs_Data1("Just", new Phpurs_Data0("March")), 4 => new Phpurs_Data1("Just", new Phpurs_Data0("April")), 5 => new Phpurs_Data1("Just", new Phpurs_Data0("May")), 6 => new Phpurs_Data1("Just", new Phpurs_Data0("June")), 7 => new Phpurs_Data1("Just", new Phpurs_Data0("July")), 8 => new Phpurs_Data1("Just", new Phpurs_Data0("August")), 9 => new Phpurs_Data1("Just", new Phpurs_Data0("September")), 10 => new Phpurs_Data1("Just", new Phpurs_Data0("October")), 11 => new Phpurs_Data1("Just", new Phpurs_Data0("November")), 12 => new Phpurs_Data1("Just", new Phpurs_Data0("December")), default => new Phpurs_Data0("Nothing") };
$__t8 = null;;
if ((is_object($__local_var_5_4) && (($__local_var_5_4)->{'tag'} === "Just"))) {
$__t9 = null;;
if ((is_object($__local_var_6_6) && (($__local_var_6_6)->{'tag'} === "Just"))) {
$__t10 = null;;
if ((((($__local_var_3_0)->{'value0'})['day'] >= 1) && ((($__local_var_3_0)->{'value0'})['day'] <= 31))) {
$__t10 = new Phpurs_Data1("Just", ((($__local_var_5_4)->{'value0'})(($__local_var_6_6)->{'value0'}))((($__local_var_3_0)->{'value0'})['day']));
goto end_branch_10;;
};
$__t10 = new Phpurs_Data0("Nothing");
end_branch_10:;
$__t9 = $__t10;
goto end_branch_9;;
};
if ((((($__local_var_3_0)->{'value0'})['day'] >= 1) && ((($__local_var_3_0)->{'value0'})['day'] <= 31))) {
$__t9 = new Phpurs_Data0("Nothing");
goto end_branch_9;;
};
$__t9 = new Phpurs_Data0("Nothing");
end_branch_9:;
$__t8 = $__t9;
goto end_branch_8;;
};
if ((is_object($__local_var_5_4) && (($__local_var_5_4)->{'tag'} === "Nothing"))) {
$__t11 = null;;
if ((((($__local_var_3_0)->{'value0'})['day'] >= 1) && ((($__local_var_3_0)->{'value0'})['day'] <= 31))) {
$__t11 = new Phpurs_Data0("Nothing");
goto end_branch_11;;
};
$__t11 = new Phpurs_Data0("Nothing");
end_branch_11:;
$__t8 = $__t11;
goto end_branch_8;;
};
if ((((($__local_var_3_0)->{'value0'})['day'] >= 1) && ((($__local_var_3_0)->{'value0'})['day'] <= 31))) {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t8 = null;
goto end_branch_8;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t8 = null;
end_branch_8:;
$__local_var_7_8 = $__t8;
$__t13 = null;;
if ((is_object($__local_var_7_8) && (($__local_var_7_8)->{'tag'} === "Just"))) {
$__t13 = ($__local_var_7_8)->{'value0'};
goto end_branch_13;;
};
if ((is_object($__local_var_7_8) && (($__local_var_7_8)->{'tag'} === "Nothing"))) {
$__t13 = new Phpurs_Data0("Nothing");
goto end_branch_13;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t13 = null;
end_branch_13:;
$__local_var_8_13 = $__t13;
$__t15 = null;;
if ((is_object($__local_var_8_13) && (($__local_var_8_13)->{'tag'} === "Just"))) {
$__t15 = new Phpurs_Data1("Just", ($GLOBALS['Data_DateTime_DateTime'])(($__local_var_8_13)->{'value0'}));
goto end_branch_15;;
};
$__t15 = new Phpurs_Data0("Nothing");
end_branch_15:;
$__local_var_9_15 = $__t15;
$__t17 = null;;
if ((((($__local_var_3_0)->{'value0'})['hour'] >= 0) && ((($__local_var_3_0)->{'value0'})['hour'] <= 23))) {
$__t18 = null;;
if ((((($__local_var_3_0)->{'value0'})['minute'] >= 0) && ((($__local_var_3_0)->{'value0'})['minute'] <= 59))) {
$__t19 = null;;
if ((((($__local_var_3_0)->{'value0'})['second'] >= 0) && ((($__local_var_3_0)->{'value0'})['second'] <= 59))) {
$__t20 = null;;
if ((((($__local_var_3_0)->{'value0'})['millisecond'] >= 0) && ((($__local_var_3_0)->{'value0'})['millisecond'] <= 999))) {
$__t21 = null;;
if ((is_object($__local_var_9_15) && (($__local_var_9_15)->{'tag'} === "Just"))) {
$__t21 = new Phpurs_Data1("Just", (($__local_var_9_15)->{'value0'})(new Phpurs_Data4("Time", (($__local_var_3_0)->{'value0'})['hour'], (($__local_var_3_0)->{'value0'})['minute'], (($__local_var_3_0)->{'value0'})['second'], (($__local_var_3_0)->{'value0'})['millisecond'])));
goto end_branch_21;;
};
if ((is_object($__local_var_9_15) && (($__local_var_9_15)->{'tag'} === "Nothing"))) {
$__t21 = new Phpurs_Data0("Nothing");
goto end_branch_21;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t21 = null;
end_branch_21:;
$__t20 = $__t21;
goto end_branch_20;;
};
if ((is_object($__local_var_9_15) && (($__local_var_9_15)->{'tag'} === "Just"))) {
$__t20 = new Phpurs_Data0("Nothing");
goto end_branch_20;;
};
if ((is_object($__local_var_9_15) && (($__local_var_9_15)->{'tag'} === "Nothing"))) {
$__t20 = new Phpurs_Data0("Nothing");
goto end_branch_20;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t20 = null;
end_branch_20:;
$__t19 = $__t20;
goto end_branch_19;;
};
if ((((($__local_var_3_0)->{'value0'})['millisecond'] >= 0) && ((($__local_var_3_0)->{'value0'})['millisecond'] <= 999))) {
$__t22 = null;;
if ((is_object($__local_var_9_15) && (($__local_var_9_15)->{'tag'} === "Just"))) {
$__t22 = new Phpurs_Data0("Nothing");
goto end_branch_22;;
};
if ((is_object($__local_var_9_15) && (($__local_var_9_15)->{'tag'} === "Nothing"))) {
$__t22 = new Phpurs_Data0("Nothing");
goto end_branch_22;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t22 = null;
end_branch_22:;
$__t19 = $__t22;
goto end_branch_19;;
};
if ((is_object($__local_var_9_15) && (($__local_var_9_15)->{'tag'} === "Just"))) {
$__t19 = new Phpurs_Data0("Nothing");
goto end_branch_19;;
};
if ((is_object($__local_var_9_15) && (($__local_var_9_15)->{'tag'} === "Nothing"))) {
$__t19 = new Phpurs_Data0("Nothing");
goto end_branch_19;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t19 = null;
end_branch_19:;
$__t18 = $__t19;
goto end_branch_18;;
};
if ((((($__local_var_3_0)->{'value0'})['second'] >= 0) && ((($__local_var_3_0)->{'value0'})['second'] <= 59))) {
$__t23 = null;;
if ((((($__local_var_3_0)->{'value0'})['millisecond'] >= 0) && ((($__local_var_3_0)->{'value0'})['millisecond'] <= 999))) {
$__t24 = null;;
if ((is_object($__local_var_9_15) && (($__local_var_9_15)->{'tag'} === "Just"))) {
$__t24 = new Phpurs_Data0("Nothing");
goto end_branch_24;;
};
if ((is_object($__local_var_9_15) && (($__local_var_9_15)->{'tag'} === "Nothing"))) {
$__t24 = new Phpurs_Data0("Nothing");
goto end_branch_24;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t24 = null;
end_branch_24:;
$__t23 = $__t24;
goto end_branch_23;;
};
if ((is_object($__local_var_9_15) && (($__local_var_9_15)->{'tag'} === "Just"))) {
$__t23 = new Phpurs_Data0("Nothing");
goto end_branch_23;;
};
if ((is_object($__local_var_9_15) && (($__local_var_9_15)->{'tag'} === "Nothing"))) {
$__t23 = new Phpurs_Data0("Nothing");
goto end_branch_23;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t23 = null;
end_branch_23:;
$__t18 = $__t23;
goto end_branch_18;;
};
if ((((($__local_var_3_0)->{'value0'})['millisecond'] >= 0) && ((($__local_var_3_0)->{'value0'})['millisecond'] <= 999))) {
$__t25 = null;;
if ((is_object($__local_var_9_15) && (($__local_var_9_15)->{'tag'} === "Just"))) {
$__t25 = new Phpurs_Data0("Nothing");
goto end_branch_25;;
};
if ((is_object($__local_var_9_15) && (($__local_var_9_15)->{'tag'} === "Nothing"))) {
$__t25 = new Phpurs_Data0("Nothing");
goto end_branch_25;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t25 = null;
end_branch_25:;
$__t18 = $__t25;
goto end_branch_18;;
};
if ((is_object($__local_var_9_15) && (($__local_var_9_15)->{'tag'} === "Just"))) {
$__t18 = new Phpurs_Data0("Nothing");
goto end_branch_18;;
};
if ((is_object($__local_var_9_15) && (($__local_var_9_15)->{'tag'} === "Nothing"))) {
$__t18 = new Phpurs_Data0("Nothing");
goto end_branch_18;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t18 = null;
end_branch_18:;
$__t17 = $__t18;
goto end_branch_17;;
};
if ((((($__local_var_3_0)->{'value0'})['millisecond'] >= 0) && ((($__local_var_3_0)->{'value0'})['millisecond'] <= 999))) {
$__t26 = null;;
if ((is_object($__local_var_9_15) && (($__local_var_9_15)->{'tag'} === "Just"))) {
$__t26 = new Phpurs_Data0("Nothing");
goto end_branch_26;;
};
if ((is_object($__local_var_9_15) && (($__local_var_9_15)->{'tag'} === "Nothing"))) {
$__t26 = new Phpurs_Data0("Nothing");
goto end_branch_26;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t26 = null;
end_branch_26:;
$__t17 = $__t26;
goto end_branch_17;;
};
if ((is_object($__local_var_9_15) && (($__local_var_9_15)->{'tag'} === "Just"))) {
$__t17 = new Phpurs_Data0("Nothing");
goto end_branch_17;;
};
if ((is_object($__local_var_9_15) && (($__local_var_9_15)->{'tag'} === "Nothing"))) {
$__t17 = new Phpurs_Data0("Nothing");
goto end_branch_17;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t17 = null;
end_branch_17:;
$__t1 = $__t17;
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
})();

