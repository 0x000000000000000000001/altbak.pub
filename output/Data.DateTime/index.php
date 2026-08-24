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
if (!\function_exists(__NAMESPACE__ . '\\phpurs_execute_effect')) {
  function phpurs_execute_effect($val) {
    if (\is_callable($val)) {
      return $val($GLOBALS['Data_Unit_unit']);
    }
    return $val;
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

$calcDiff = function($rec1, $rec2) use (&$calcDiff, &$createUTC) {

    $msUTC1 = $createUTC($rec1->year, $rec1->month - 1, $rec1->day, $rec1->hour, $rec1->minute, $rec1->second, $rec1->millisecond);
    $msUTC2 = $createUTC($rec2->year, $rec2->month - 1, $rec2->day, $rec2->hour, $rec2->minute, $rec2->second, $rec2->millisecond);
    return $msUTC1 - $msUTC2;
};

$adjustImpl = function($just, $nothing, $offset, $rec) use (&$adjustImpl, &$createUTC) {

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
function majData_majDatemajTime_adjustmajImpl($v0, $v1 = null, $v2 = null, $v3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majDatemajTime_adjustmajImpl';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  global $ffi_Data_DateTime;
  $f = (\array_key_exists('adjustImpl', $ffi_Data_DateTime) ? $ffi_Data_DateTime['adjustImpl'] : new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2, $v3);
}
$GLOBALS['Data_DateTime_adjustImpl'] = __NAMESPACE__ . '\\majData_majDatemajTime_adjustmajImpl';

$GLOBALS['Data_DateTime_calcDiff'] = (\array_key_exists('calcDiff', $ffi_Data_DateTime) ? $ffi_Data_DateTime['calcDiff'] : new class { public function __invoke(...$args) { return $this; } });


final class Data_DateTime_DateTime { public $tag = 'DateTime'; public function __construct(public  $value0, public  $value1) {} }

// Data_DateTime_DateTime
$GLOBALS['Data_DateTime_DateTime'] = (function() {
  $__fn = function($value0, $value1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\DateTime\Data_DateTime_DateTime($value0, $value1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_DateTime_toRecord
function majData_majDatemajTime_tomajRecord($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majDatemajTime_tomajRecord';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__t0 = null;;
  if ((($v_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_January) {
$__t0 = 1;
goto end_branch_0;;
};
  if ((($v_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_February) {
$__t0 = 2;
goto end_branch_0;;
};
  if ((($v_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_March) {
$__t0 = 3;
goto end_branch_0;;
};
  if ((($v_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_April) {
$__t0 = 4;
goto end_branch_0;;
};
  if ((($v_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_May) {
$__t0 = 5;
goto end_branch_0;;
};
  if ((($v_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_June) {
$__t0 = 6;
goto end_branch_0;;
};
  if ((($v_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_July) {
$__t0 = 7;
goto end_branch_0;;
};
  if ((($v_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_August) {
$__t0 = 8;
goto end_branch_0;;
};
  if ((($v_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_September) {
$__t0 = 9;
goto end_branch_0;;
};
  if ((($v_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_October) {
$__t0 = 10;
goto end_branch_0;;
};
  if ((($v_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_November) {
$__t0 = 11;
goto end_branch_0;;
};
  if ((($v_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_December) {
$__t0 = 12;
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = (object)["year" => (($v_0)->{'value0'})->{'value0'}, "month" => $__t0, "day" => (($v_0)->{'value0'})->{'value2'}, "hour" => (($v_0)->{'value1'})->{'value0'}, "minute" => (($v_0)->{'value1'})->{'value1'}, "second" => (($v_0)->{'value1'})->{'value2'}, "millisecond" => (($v_0)->{'value1'})->{'value3'}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_DateTime_toRecord'] = __NAMESPACE__ . '\\majData_majDatemajTime_tomajRecord';

// Data_DateTime_time
function majData_majDatemajTime_time($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majDatemajTime_time';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($v_0)->{'value1'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_DateTime_time'] = __NAMESPACE__ . '\\majData_majDatemajTime_time';

// Data_DateTime_showDateTime
$GLOBALS['Data_DateTime_showDateTime'] = (object)["show" => function($v_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((($v_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_January) {
$__t0 = "January";
goto end_branch_0;;
};
  if ((($v_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_February) {
$__t0 = "February";
goto end_branch_0;;
};
  if ((($v_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_March) {
$__t0 = "March";
goto end_branch_0;;
};
  if ((($v_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_April) {
$__t0 = "April";
goto end_branch_0;;
};
  if ((($v_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_May) {
$__t0 = "May";
goto end_branch_0;;
};
  if ((($v_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_June) {
$__t0 = "June";
goto end_branch_0;;
};
  if ((($v_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_July) {
$__t0 = "July";
goto end_branch_0;;
};
  if ((($v_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_August) {
$__t0 = "August";
goto end_branch_0;;
};
  if ((($v_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_September) {
$__t0 = "September";
goto end_branch_0;;
};
  if ((($v_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_October) {
$__t0 = "October";
goto end_branch_0;;
};
  if ((($v_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_November) {
$__t0 = "November";
goto end_branch_0;;
};
  if ((($v_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_December) {
$__t0 = "December";
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = (((((((((((((("(DateTime (Date (Year " . \Data\Show\majData_majShow_showmajIntmajImpl((($v_0)->{'value0'})->{'value0'})) . ") ") . $__t0) . " (Day ") . \Data\Show\majData_majShow_showmajIntmajImpl((($v_0)->{'value0'})->{'value2'})) . ")) (Time (Hour ") . \Data\Show\majData_majShow_showmajIntmajImpl((($v_0)->{'value1'})->{'value0'})) . ") (Minute ") . \Data\Show\majData_majShow_showmajIntmajImpl((($v_0)->{'value1'})->{'value1'})) . ") (Second ") . \Data\Show\majData_majShow_showmajIntmajImpl((($v_0)->{'value1'})->{'value2'})) . ") (Millisecond ") . \Data\Show\majData_majShow_showmajIntmajImpl((($v_0)->{'value1'})->{'value3'})) . ")))");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_DateTime_modifyTimeF
function majData_majDatemajTime_modifymajTimemajF($dictFunctor_0, $f_1 = null, $v_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majDatemajTime_modifymajTimemajF';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictFunctor_0)->{'map'})(($GLOBALS['Data_DateTime_DateTime'])(($v_2)->{'value0'})))(($f_1)(($v_2)->{'value1'}));
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_DateTime_modifyTimeF'] = __NAMESPACE__ . '\\majData_majDatemajTime_modifymajTimemajF';

// Data_DateTime_modifyTime
function majData_majDatemajTime_modifymajTime($f_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majDatemajTime_modifymajTime';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\DateTime\Data_DateTime_DateTime(($v_1)->{'value0'}, ($f_0)(($v_1)->{'value1'}));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_DateTime_modifyTime'] = __NAMESPACE__ . '\\majData_majDatemajTime_modifymajTime';

// Data_DateTime_modifyDateF
function majData_majDatemajTime_modifymajDatemajF($dictFunctor_0, $f_1 = null, $v_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majDatemajTime_modifymajDatemajF';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictFunctor_0)->{'map'})(function($a_3) use ($v_2) {
  $__num = \func_num_args();
  $__res = new \Data\DateTime\Data_DateTime_DateTime($a_3, ($v_2)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($f_1)(($v_2)->{'value0'}));
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_DateTime_modifyDateF'] = __NAMESPACE__ . '\\majData_majDatemajTime_modifymajDatemajF';

// Data_DateTime_modifyDate
function majData_majDatemajTime_modifymajDate($f_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majDatemajTime_modifymajDate';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\DateTime\Data_DateTime_DateTime(($f_0)(($v_1)->{'value0'}), ($v_1)->{'value1'});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_DateTime_modifyDate'] = __NAMESPACE__ . '\\majData_majDatemajTime_modifymajDate';

// Data_DateTime_eqDateTime
$GLOBALS['Data_DateTime_eqDateTime'] = (object)["eq" => function($x_0) {
  $__num = \func_num_args();
  $__res = function($y_1) use ($x_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((($x_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_January) {
$__t0 = (($y_1)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_January;
goto end_branch_0;;
};
  if ((($x_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_February) {
$__t0 = (($y_1)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_February;
goto end_branch_0;;
};
  if ((($x_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_March) {
$__t0 = (($y_1)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_March;
goto end_branch_0;;
};
  if ((($x_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_April) {
$__t0 = (($y_1)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_April;
goto end_branch_0;;
};
  if ((($x_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_May) {
$__t0 = (($y_1)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_May;
goto end_branch_0;;
};
  if ((($x_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_June) {
$__t0 = (($y_1)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_June;
goto end_branch_0;;
};
  if ((($x_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_July) {
$__t0 = (($y_1)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_July;
goto end_branch_0;;
};
  if ((($x_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_August) {
$__t0 = (($y_1)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_August;
goto end_branch_0;;
};
  if ((($x_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_September) {
$__t0 = (($y_1)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_September;
goto end_branch_0;;
};
  if ((($x_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_October) {
$__t0 = (($y_1)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_October;
goto end_branch_0;;
};
  if ((($x_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_November) {
$__t0 = (($y_1)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_November;
goto end_branch_0;;
};
  $__t0 = ((($x_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_December && (($y_1)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_December);
  end_branch_0:;
  $__res = (((((($x_0)->{'value0'})->{'value0'} === (($y_1)->{'value0'})->{'value0'}) && $__t0) && ((($x_0)->{'value0'})->{'value2'} === (($y_1)->{'value0'})->{'value2'})) && (((((($x_0)->{'value1'})->{'value0'} === (($y_1)->{'value1'})->{'value0'}) && ((($x_0)->{'value1'})->{'value1'} === (($y_1)->{'value1'})->{'value1'})) && ((($x_0)->{'value1'})->{'value2'} === (($y_1)->{'value1'})->{'value2'})) && ((($x_0)->{'value1'})->{'value3'} === (($y_1)->{'value1'})->{'value3'})));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_DateTime_ordDateTime
$GLOBALS['Data_DateTime_ordDateTime'] = (object)["compare" => function($x_0) {
  $__num = \func_num_args();
  $__res = function($y_1) use ($x_0) {
  $__num = \func_num_args();
  $v_2_0 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), (($x_0)->{'value0'})->{'value0'}, (($y_1)->{'value0'})->{'value0'});
  $__t13 = null;;
  if ($v_2_0 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t13 = new \Data\Ordering\Data_Ordering_LT();
goto end_branch_13;;
};
  if ($v_2_0 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t13 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_13;;
};
  $__t1 = null;;
  if ((($x_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_January) {
$__t2 = null;;
if ((($y_1)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_January) {
$__t2 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), (($x_0)->{'value0'})->{'value2'}, (($y_1)->{'value0'})->{'value2'});
goto end_branch_2;;
};
$__t2 = new \Data\Ordering\Data_Ordering_LT();
end_branch_2:;
$__t1 = $__t2;
goto end_branch_1;;
};
  if ((($y_1)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_January) {
$__t1 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_1;;
};
  if ((($x_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_February) {
$__t3 = null;;
if ((($y_1)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_February) {
$__t3 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), (($x_0)->{'value0'})->{'value2'}, (($y_1)->{'value0'})->{'value2'});
goto end_branch_3;;
};
$__t3 = new \Data\Ordering\Data_Ordering_LT();
end_branch_3:;
$__t1 = $__t3;
goto end_branch_1;;
};
  if ((($y_1)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_February) {
$__t1 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_1;;
};
  if ((($x_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_March) {
$__t4 = null;;
if ((($y_1)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_March) {
$__t4 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), (($x_0)->{'value0'})->{'value2'}, (($y_1)->{'value0'})->{'value2'});
goto end_branch_4;;
};
$__t4 = new \Data\Ordering\Data_Ordering_LT();
end_branch_4:;
$__t1 = $__t4;
goto end_branch_1;;
};
  if ((($y_1)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_March) {
$__t1 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_1;;
};
  if ((($x_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_April) {
$__t5 = null;;
if ((($y_1)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_April) {
$__t5 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), (($x_0)->{'value0'})->{'value2'}, (($y_1)->{'value0'})->{'value2'});
goto end_branch_5;;
};
$__t5 = new \Data\Ordering\Data_Ordering_LT();
end_branch_5:;
$__t1 = $__t5;
goto end_branch_1;;
};
  if ((($y_1)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_April) {
$__t1 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_1;;
};
  if ((($x_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_May) {
$__t6 = null;;
if ((($y_1)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_May) {
$__t6 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), (($x_0)->{'value0'})->{'value2'}, (($y_1)->{'value0'})->{'value2'});
goto end_branch_6;;
};
$__t6 = new \Data\Ordering\Data_Ordering_LT();
end_branch_6:;
$__t1 = $__t6;
goto end_branch_1;;
};
  if ((($y_1)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_May) {
$__t1 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_1;;
};
  if ((($x_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_June) {
$__t7 = null;;
if ((($y_1)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_June) {
$__t7 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), (($x_0)->{'value0'})->{'value2'}, (($y_1)->{'value0'})->{'value2'});
goto end_branch_7;;
};
$__t7 = new \Data\Ordering\Data_Ordering_LT();
end_branch_7:;
$__t1 = $__t7;
goto end_branch_1;;
};
  if ((($y_1)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_June) {
$__t1 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_1;;
};
  if ((($x_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_July) {
$__t8 = null;;
if ((($y_1)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_July) {
$__t8 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), (($x_0)->{'value0'})->{'value2'}, (($y_1)->{'value0'})->{'value2'});
goto end_branch_8;;
};
$__t8 = new \Data\Ordering\Data_Ordering_LT();
end_branch_8:;
$__t1 = $__t8;
goto end_branch_1;;
};
  if ((($y_1)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_July) {
$__t1 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_1;;
};
  if ((($x_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_August) {
$__t9 = null;;
if ((($y_1)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_August) {
$__t9 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), (($x_0)->{'value0'})->{'value2'}, (($y_1)->{'value0'})->{'value2'});
goto end_branch_9;;
};
$__t9 = new \Data\Ordering\Data_Ordering_LT();
end_branch_9:;
$__t1 = $__t9;
goto end_branch_1;;
};
  if ((($y_1)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_August) {
$__t1 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_1;;
};
  if ((($x_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_September) {
$__t10 = null;;
if ((($y_1)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_September) {
$__t10 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), (($x_0)->{'value0'})->{'value2'}, (($y_1)->{'value0'})->{'value2'});
goto end_branch_10;;
};
$__t10 = new \Data\Ordering\Data_Ordering_LT();
end_branch_10:;
$__t1 = $__t10;
goto end_branch_1;;
};
  if ((($y_1)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_September) {
$__t1 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_1;;
};
  if ((($x_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_October) {
$__t11 = null;;
if ((($y_1)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_October) {
$__t11 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), (($x_0)->{'value0'})->{'value2'}, (($y_1)->{'value0'})->{'value2'});
goto end_branch_11;;
};
$__t11 = new \Data\Ordering\Data_Ordering_LT();
end_branch_11:;
$__t1 = $__t11;
goto end_branch_1;;
};
  if ((($y_1)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_October) {
$__t1 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_1;;
};
  if ((($x_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_November) {
$__t12 = null;;
if ((($y_1)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_November) {
$__t12 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), (($x_0)->{'value0'})->{'value2'}, (($y_1)->{'value0'})->{'value2'});
goto end_branch_12;;
};
$__t12 = new \Data\Ordering\Data_Ordering_LT();
end_branch_12:;
$__t1 = $__t12;
goto end_branch_1;;
};
  if ((($y_1)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_November) {
$__t1 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_1;;
};
  if (((($x_0)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_December && (($y_1)->{'value0'})->{'value1'} instanceof \Data\Date\Component\Data_Date_Component_December)) {
$__t1 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), (($x_0)->{'value0'})->{'value2'}, (($y_1)->{'value0'})->{'value2'});
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__t13 = $__t1;
  end_branch_13:;
  $v_2_0 = $__t13;
  $__t21 = null;;
  if ($v_2_0 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t21 = new \Data\Ordering\Data_Ordering_LT();
goto end_branch_21;;
};
  if ($v_2_0 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t21 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_21;;
};
  $v_3_15 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), (($x_0)->{'value1'})->{'value0'}, (($y_1)->{'value1'})->{'value0'});
  $__t20 = null;;
  if ($v_3_15 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t20 = new \Data\Ordering\Data_Ordering_LT();
goto end_branch_20;;
};
  if ($v_3_15 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t20 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_20;;
};
  $v1_4_16 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), (($x_0)->{'value1'})->{'value1'}, (($y_1)->{'value1'})->{'value1'});
  $__t19 = null;;
  if ($v1_4_16 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t19 = new \Data\Ordering\Data_Ordering_LT();
goto end_branch_19;;
};
  if ($v1_4_16 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t19 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_19;;
};
  $v2_5_17 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), (($x_0)->{'value1'})->{'value2'}, (($y_1)->{'value1'})->{'value2'});
  $__t18 = null;;
  if ($v2_5_17 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t18 = new \Data\Ordering\Data_Ordering_LT();
goto end_branch_18;;
};
  if ($v2_5_17 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t18 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_18;;
};
  $__t18 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), (($x_0)->{'value1'})->{'value3'}, (($y_1)->{'value1'})->{'value3'});
  end_branch_18:;
  $__t19 = $__t18;
  end_branch_19:;
  $__t20 = $__t19;
  end_branch_20:;
  $__t21 = $__t20;
  end_branch_21:;
  $__res = $__t21;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_DateTime_eqDateTime'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_DateTime_diff
function majData_majDatemajTime_diff($dictDuration_0, $dt1_1 = null, $dt2_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majDatemajTime_diff';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($dictDuration_0)->{'toDuration'})(($GLOBALS['Data_DateTime_calcDiff'])(\Data\DateTime\majData_majDatemajTime_tomajRecord($dt1_1), \Data\DateTime\majData_majDatemajTime_tomajRecord($dt2_2)));
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_DateTime_diff'] = __NAMESPACE__ . '\\majData_majDatemajTime_diff';

// Data_DateTime_date
function majData_majDatemajTime_date($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majDatemajTime_date';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($v_0)->{'value0'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_DateTime_date'] = __NAMESPACE__ . '\\majData_majDatemajTime_date';

// Data_DateTime_boundedDateTime
$GLOBALS['Data_DateTime_boundedDateTime'] = (object)["bottom" => new \Data\DateTime\Data_DateTime_DateTime(new \Data\Date\Data_Date_Date(-271820, new \Data\Date\Component\Data_Date_Component_January(), 1), new \Data\Time\Data_Time_Time(0, 0, 0, 0)), "top" => new \Data\DateTime\Data_DateTime_DateTime(new \Data\Date\Data_Date_Date(275759, new \Data\Date\Component\Data_Date_Component_December(), 31), new \Data\Time\Data_Time_Time(23, 59, 59, 999)), "Ord0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_DateTime_ordDateTime'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_DateTime_adjust
function majData_majDatemajTime_adjust($dictDuration_0, $d_1 = null, $dt_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majDatemajTime_adjust';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__local_var_3_0 = \Data\DateTime\majData_majDatemajTime_adjustmajImpl($GLOBALS['Data_Maybe_Just'], new \Data\Maybe\Data_Maybe_Nothing(), (($dictDuration_0)->{'fromDuration'})($d_1), \Data\DateTime\majData_majDatemajTime_tomajRecord($dt_2));
  $__t1 = null;;
  if ($__local_var_3_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t2 = null;;
if ((((($__local_var_3_0)->{'value0'})->{'year'} >= -271820) && ((($__local_var_3_0)->{'value0'})->{'year'} <= 275759))) {
$__t2 = new \Data\Maybe\Data_Maybe_Just(($GLOBALS['Data_Date_exactDate'])((($__local_var_3_0)->{'value0'})->{'year'}));
goto end_branch_2;;
};
$__t2 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_2:;
$__local_var_4_2 = $__t2;
$__t5 = null;;
switch ((($__local_var_3_0)->{'value0'})->{'month'}) {
case 1:
$__t6 = null;;
if ($__local_var_4_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t6 = new \Data\Maybe\Data_Maybe_Just((($__local_var_4_2)->{'value0'})(new \Data\Date\Component\Data_Date_Component_January()));
goto end_branch_6;;
};
if ($__local_var_4_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t6 = new \Data\Maybe\Data_Maybe_Nothing();
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
$__t7 = null;;
if ($__local_var_4_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t7 = new \Data\Maybe\Data_Maybe_Just((($__local_var_4_2)->{'value0'})(new \Data\Date\Component\Data_Date_Component_February()));
goto end_branch_7;;
};
if ($__local_var_4_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t7 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_7;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t7 = null;
end_branch_7:;
$__t5 = $__t7;
goto end_branch_5;;
break;
default:
;
break;
};
switch ((($__local_var_3_0)->{'value0'})->{'month'}) {
case 3:
$__t8 = null;;
if ($__local_var_4_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t8 = new \Data\Maybe\Data_Maybe_Just((($__local_var_4_2)->{'value0'})(new \Data\Date\Component\Data_Date_Component_March()));
goto end_branch_8;;
};
if ($__local_var_4_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t8 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_8;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t8 = null;
end_branch_8:;
$__t5 = $__t8;
goto end_branch_5;;
break;
default:
;
break;
};
switch ((($__local_var_3_0)->{'value0'})->{'month'}) {
case 4:
$__t9 = null;;
if ($__local_var_4_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t9 = new \Data\Maybe\Data_Maybe_Just((($__local_var_4_2)->{'value0'})(new \Data\Date\Component\Data_Date_Component_April()));
goto end_branch_9;;
};
if ($__local_var_4_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t9 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_9;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t9 = null;
end_branch_9:;
$__t5 = $__t9;
goto end_branch_5;;
break;
default:
;
break;
};
switch ((($__local_var_3_0)->{'value0'})->{'month'}) {
case 5:
$__t10 = null;;
if ($__local_var_4_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t10 = new \Data\Maybe\Data_Maybe_Just((($__local_var_4_2)->{'value0'})(new \Data\Date\Component\Data_Date_Component_May()));
goto end_branch_10;;
};
if ($__local_var_4_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t10 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_10;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t10 = null;
end_branch_10:;
$__t5 = $__t10;
goto end_branch_5;;
break;
default:
;
break;
};
switch ((($__local_var_3_0)->{'value0'})->{'month'}) {
case 6:
$__t11 = null;;
if ($__local_var_4_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t11 = new \Data\Maybe\Data_Maybe_Just((($__local_var_4_2)->{'value0'})(new \Data\Date\Component\Data_Date_Component_June()));
goto end_branch_11;;
};
if ($__local_var_4_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t11 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_11;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t11 = null;
end_branch_11:;
$__t5 = $__t11;
goto end_branch_5;;
break;
default:
;
break;
};
switch ((($__local_var_3_0)->{'value0'})->{'month'}) {
case 7:
$__t12 = null;;
if ($__local_var_4_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t12 = new \Data\Maybe\Data_Maybe_Just((($__local_var_4_2)->{'value0'})(new \Data\Date\Component\Data_Date_Component_July()));
goto end_branch_12;;
};
if ($__local_var_4_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t12 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_12;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t12 = null;
end_branch_12:;
$__t5 = $__t12;
goto end_branch_5;;
break;
default:
;
break;
};
switch ((($__local_var_3_0)->{'value0'})->{'month'}) {
case 8:
$__t13 = null;;
if ($__local_var_4_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t13 = new \Data\Maybe\Data_Maybe_Just((($__local_var_4_2)->{'value0'})(new \Data\Date\Component\Data_Date_Component_August()));
goto end_branch_13;;
};
if ($__local_var_4_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t13 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_13;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t13 = null;
end_branch_13:;
$__t5 = $__t13;
goto end_branch_5;;
break;
default:
;
break;
};
switch ((($__local_var_3_0)->{'value0'})->{'month'}) {
case 9:
$__t14 = null;;
if ($__local_var_4_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t14 = new \Data\Maybe\Data_Maybe_Just((($__local_var_4_2)->{'value0'})(new \Data\Date\Component\Data_Date_Component_September()));
goto end_branch_14;;
};
if ($__local_var_4_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t14 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_14;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t14 = null;
end_branch_14:;
$__t5 = $__t14;
goto end_branch_5;;
break;
default:
;
break;
};
switch ((($__local_var_3_0)->{'value0'})->{'month'}) {
case 10:
$__t15 = null;;
if ($__local_var_4_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t15 = new \Data\Maybe\Data_Maybe_Just((($__local_var_4_2)->{'value0'})(new \Data\Date\Component\Data_Date_Component_October()));
goto end_branch_15;;
};
if ($__local_var_4_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t15 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_15;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t15 = null;
end_branch_15:;
$__t5 = $__t15;
goto end_branch_5;;
break;
default:
;
break;
};
switch ((($__local_var_3_0)->{'value0'})->{'month'}) {
case 11:
$__t16 = null;;
if ($__local_var_4_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t16 = new \Data\Maybe\Data_Maybe_Just((($__local_var_4_2)->{'value0'})(new \Data\Date\Component\Data_Date_Component_November()));
goto end_branch_16;;
};
if ($__local_var_4_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t16 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_16;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t16 = null;
end_branch_16:;
$__t5 = $__t16;
goto end_branch_5;;
break;
default:
;
break;
};
switch ((($__local_var_3_0)->{'value0'})->{'month'}) {
case 12:
$__t17 = null;;
if ($__local_var_4_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t17 = new \Data\Maybe\Data_Maybe_Just((($__local_var_4_2)->{'value0'})(new \Data\Date\Component\Data_Date_Component_December()));
goto end_branch_17;;
};
if ($__local_var_4_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t17 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_17;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t17 = null;
end_branch_17:;
$__t5 = $__t17;
goto end_branch_5;;
break;
default:
;
break;
};
$__t4 = null;;
if ($__local_var_4_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t4 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_4;;
};
if ($__local_var_4_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t4 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_4;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t4 = null;
end_branch_4:;
$__t5 = $__t4;
end_branch_5:;
$__local_var_4_2 = $__t5;
$__t20 = null;;
if ((((($__local_var_3_0)->{'value0'})->{'day'} >= 1) && ((($__local_var_3_0)->{'value0'})->{'day'} <= 31))) {
$__t21 = null;;
if ($__local_var_4_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t21 = new \Data\Maybe\Data_Maybe_Just((($__local_var_4_2)->{'value0'})((($__local_var_3_0)->{'value0'})->{'day'}));
goto end_branch_21;;
};
if ($__local_var_4_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t21 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_21;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t21 = null;
end_branch_21:;
$__t20 = $__t21;
goto end_branch_20;;
};
$__t19 = null;;
if ($__local_var_4_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t19 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_19;;
};
if ($__local_var_4_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t19 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_19;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t19 = null;
end_branch_19:;
$__t20 = $__t19;
end_branch_20:;
$__local_var_4_2 = $__t20;
$__t23 = null;;
if ($__local_var_4_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t23 = ($__local_var_4_2)->{'value0'};
goto end_branch_23;;
};
if ($__local_var_4_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t23 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_23;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t23 = null;
end_branch_23:;
$__local_var_4_2 = $__t23;
$__t25 = null;;
if ($__local_var_4_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t25 = new \Data\Maybe\Data_Maybe_Just(($GLOBALS['Data_DateTime_DateTime'])(($__local_var_4_2)->{'value0'}));
goto end_branch_25;;
};
$__t25 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_25:;
$__local_var_4_2 = $__t25;
$__t28 = null;;
if ((((($__local_var_3_0)->{'value0'})->{'hour'} >= 0) && ((($__local_var_3_0)->{'value0'})->{'hour'} <= 23))) {
$__t29 = null;;
if ($__local_var_4_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t30 = null;;
if ((function() use ($__local_var_3_0, &$__fn) {
$__t31 = null;;
if ((((($__local_var_3_0)->{'value0'})->{'millisecond'} >= 0) && ((($__local_var_3_0)->{'value0'})->{'millisecond'} <= 999))) {
$__t32 = null;;
if (((((($__local_var_3_0)->{'value0'})->{'second'} >= 0) && ((($__local_var_3_0)->{'value0'})->{'second'} <= 59)) && (((($__local_var_3_0)->{'value0'})->{'minute'} >= 0) && ((($__local_var_3_0)->{'value0'})->{'minute'} <= 59)))) {
$__t32 = true;
goto end_branch_32;;
};
if ((function() use ($__local_var_3_0, &$__fn) {
$__t33 = null;;
if ((((($__local_var_3_0)->{'value0'})->{'second'} >= 0) && ((($__local_var_3_0)->{'value0'})->{'second'} <= 59))) {
$__t33 = ( ! (((($__local_var_3_0)->{'value0'})->{'minute'} >= 0) && ((($__local_var_3_0)->{'value0'})->{'minute'} <= 59)));
goto end_branch_33;;
};
$__t33 = true;
end_branch_33:;
return $__t33;
})()) {
$__t32 = false;
goto end_branch_32;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t32 = null;
end_branch_32:;
$__t31 = $__t32;
goto end_branch_31;;
};
if (((((($__local_var_3_0)->{'value0'})->{'second'} >= 0) && ((($__local_var_3_0)->{'value0'})->{'second'} <= 59)) && (((($__local_var_3_0)->{'value0'})->{'minute'} >= 0) && ((($__local_var_3_0)->{'value0'})->{'minute'} <= 59)))) {
$__t31 = false;
goto end_branch_31;;
};
if ((function() use ($__local_var_3_0, &$__fn) {
$__t34 = null;;
if ((((($__local_var_3_0)->{'value0'})->{'second'} >= 0) && ((($__local_var_3_0)->{'value0'})->{'second'} <= 59))) {
$__t34 = ( ! (((($__local_var_3_0)->{'value0'})->{'minute'} >= 0) && ((($__local_var_3_0)->{'value0'})->{'minute'} <= 59)));
goto end_branch_34;;
};
$__t34 = true;
end_branch_34:;
return $__t34;
})()) {
$__t31 = false;
goto end_branch_31;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t31 = null;
end_branch_31:;
return $__t31;
})()) {
$__t35 = null;;
if ((((($__local_var_3_0)->{'value0'})->{'millisecond'} >= 0) && ((($__local_var_3_0)->{'value0'})->{'millisecond'} <= 999))) {
$__t36 = null;;
if (((((($__local_var_3_0)->{'value0'})->{'second'} >= 0) && ((($__local_var_3_0)->{'value0'})->{'second'} <= 59)) && (((($__local_var_3_0)->{'value0'})->{'minute'} >= 0) && ((($__local_var_3_0)->{'value0'})->{'minute'} <= 59)))) {
$__t37 = null;;
if ((((($__local_var_3_0)->{'value0'})->{'second'} >= 0) && ((($__local_var_3_0)->{'value0'})->{'second'} <= 59))) {
$__t38 = null;;
if ((((($__local_var_3_0)->{'value0'})->{'minute'} >= 0) && ((($__local_var_3_0)->{'value0'})->{'minute'} <= 59))) {
$__t39 = null;;
if ((((($__local_var_3_0)->{'value0'})->{'minute'} >= 0) && ((($__local_var_3_0)->{'value0'})->{'minute'} <= 59))) {
$__t39 = new \Data\Time\Data_Time_Time((($__local_var_3_0)->{'value0'})->{'hour'}, (($__local_var_3_0)->{'value0'})->{'minute'}, (($__local_var_3_0)->{'value0'})->{'second'}, (($__local_var_3_0)->{'value0'})->{'millisecond'});
goto end_branch_39;;
};
$__t39 = (((new \Data\Maybe\Data_Maybe_Nothing())->{'value0'})((($__local_var_3_0)->{'value0'})->{'second'}))((($__local_var_3_0)->{'value0'})->{'millisecond'});
end_branch_39:;
$__t38 = $__t39;
goto end_branch_38;;
};
$__t38 = ((new \Data\Maybe\Data_Maybe_Nothing())->{'value0'})((($__local_var_3_0)->{'value0'})->{'millisecond'});
end_branch_38:;
$__t37 = $__t38;
goto end_branch_37;;
};
if ((((($__local_var_3_0)->{'value0'})->{'minute'} >= 0) && ((($__local_var_3_0)->{'value0'})->{'minute'} <= 59))) {
$__t37 = ((new \Data\Maybe\Data_Maybe_Nothing())->{'value0'})((($__local_var_3_0)->{'value0'})->{'millisecond'});
goto end_branch_37;;
};
$__t37 = ((new \Data\Maybe\Data_Maybe_Nothing())->{'value0'})((($__local_var_3_0)->{'value0'})->{'millisecond'});
end_branch_37:;
$__t36 = $__t37;
goto end_branch_36;;
};
if ((function() use ($__local_var_3_0, &$__fn) {
$__t40 = null;;
if ((((($__local_var_3_0)->{'value0'})->{'second'} >= 0) && ((($__local_var_3_0)->{'value0'})->{'second'} <= 59))) {
$__t40 = ( ! (((($__local_var_3_0)->{'value0'})->{'minute'} >= 0) && ((($__local_var_3_0)->{'value0'})->{'minute'} <= 59)));
goto end_branch_40;;
};
$__t40 = true;
end_branch_40:;
return $__t40;
})()) {
$__t36 = (new \Data\Maybe\Data_Maybe_Nothing())->{'value0'};
goto end_branch_36;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t36 = null;
end_branch_36:;
$__t35 = $__t36;
goto end_branch_35;;
};
if (((((($__local_var_3_0)->{'value0'})->{'second'} >= 0) && ((($__local_var_3_0)->{'value0'})->{'second'} <= 59)) && (((($__local_var_3_0)->{'value0'})->{'minute'} >= 0) && ((($__local_var_3_0)->{'value0'})->{'minute'} <= 59)))) {
$__t35 = (new \Data\Maybe\Data_Maybe_Nothing())->{'value0'};
goto end_branch_35;;
};
if ((function() use ($__local_var_3_0, &$__fn) {
$__t41 = null;;
if ((((($__local_var_3_0)->{'value0'})->{'second'} >= 0) && ((($__local_var_3_0)->{'value0'})->{'second'} <= 59))) {
$__t41 = ( ! (((($__local_var_3_0)->{'value0'})->{'minute'} >= 0) && ((($__local_var_3_0)->{'value0'})->{'minute'} <= 59)));
goto end_branch_41;;
};
$__t41 = true;
end_branch_41:;
return $__t41;
})()) {
$__t35 = (new \Data\Maybe\Data_Maybe_Nothing())->{'value0'};
goto end_branch_35;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t35 = null;
end_branch_35:;
$__t30 = new \Data\Maybe\Data_Maybe_Just((($__local_var_4_2)->{'value0'})($__t35));
goto end_branch_30;;
};
$__t30 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_30:;
$__t29 = $__t30;
goto end_branch_29;;
};
if ($__local_var_4_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t29 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_29;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t29 = null;
end_branch_29:;
$__t28 = $__t29;
goto end_branch_28;;
};
$__t27 = null;;
if ($__local_var_4_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t27 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_27;;
};
if ($__local_var_4_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t27 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_27;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t27 = null;
end_branch_27:;
$__t28 = $__t27;
end_branch_28:;
$__t1 = $__t28;
goto end_branch_1;;
};
  if ($__local_var_3_0 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t1 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_DateTime_adjust'] = __NAMESPACE__ . '\\majData_majDatemajTime_adjust';

