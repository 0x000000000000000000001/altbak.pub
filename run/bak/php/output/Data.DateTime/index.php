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
  $__res = (object)["year" => (($GLOBALS['Data_Date_Component_boundedEnumYear'])->{'fromEnum'})((($v_0)->{'value0'})->{'value0'}), "month" => (($GLOBALS['Data_Date_Component_boundedEnumMonth'])->{'fromEnum'})((($v_0)->{'value0'})->{'value1'}), "day" => (($GLOBALS['Data_Date_Component_boundedEnumDay'])->{'fromEnum'})((($v_0)->{'value0'})->{'value2'}), "hour" => (($GLOBALS['Data_Time_Component_boundedEnumHour'])->{'fromEnum'})((($v_0)->{'value1'})->{'value0'}), "minute" => (($GLOBALS['Data_Time_Component_boundedEnumMinute'])->{'fromEnum'})((($v_0)->{'value1'})->{'value1'}), "second" => (($GLOBALS['Data_Time_Component_boundedEnumSecond'])->{'fromEnum'})((($v_0)->{'value1'})->{'value2'}), "millisecond" => (($GLOBALS['Data_Time_Component_boundedEnumMillisecond'])->{'fromEnum'})((($v_0)->{'value1'})->{'value3'})];
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
  $__res = ((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})("(DateTime "))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})((($GLOBALS['Data_Date_showDate'])->{'show'})(($v_0)->{'value0'})))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})(" "))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})((($GLOBALS['Data_Time_showTime'])->{'show'})(($v_0)->{'value1'})))(")"))));
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
  $__local_var_3_0 = ($v_2)->{'value1'};
  $__res = ((($dictFunctor_0)->{'map'})(function($a_4) use ($__local_var_3_0) {
  $__num = \func_num_args();
  $__res = new \Data\DateTime\Data_DateTime_DateTime($a_4, $__local_var_3_0);
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
  $__res = ((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'conj'})(((($GLOBALS['Data_Date_eqDate'])->{'eq'})(($x_0)->{'value0'}))(($y_1)->{'value0'})))(((($GLOBALS['Data_Time_eqTime'])->{'eq'})(($x_0)->{'value1'}))(($y_1)->{'value1'}));
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
  $v_2_0 = ((($GLOBALS['Data_Date_ordDate'])->{'compare'})(($x_0)->{'value0'}))(($y_1)->{'value0'});
  $__t1 = null;;
  if ($v_2_0 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t1 = new \Data\Ordering\Data_Ordering_LT();
goto end_branch_1;;
};
  if ($v_2_0 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t1 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_1;;
};
  $__t1 = ((($GLOBALS['Data_Time_ordTime'])->{'compare'})(($x_0)->{'value1'}))(($y_1)->{'value1'});
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq0" => function($_dollar__unused_0) {
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
$GLOBALS['Data_DateTime_boundedDateTime'] = (object)["bottom" => new \Data\DateTime\Data_DateTime_DateTime(($GLOBALS['Data_Date_boundedDate'])->{'bottom'}, ($GLOBALS['Data_Time_boundedTime'])->{'bottom'}), "top" => new \Data\DateTime\Data_DateTime_DateTime(($GLOBALS['Data_Date_boundedDate'])->{'top'}, ($GLOBALS['Data_Time_boundedTime'])->{'top'}), "Ord0" => function($_dollar__unused_0) {
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
  $__res = ((($GLOBALS['Data_Maybe_bindMaybe'])->{'bind'})(\Data\DateTime\majData_majDatemajTime_adjustmajImpl($GLOBALS['Data_Maybe_Just'], new \Data\Maybe\Data_Maybe_Nothing(), (($dictDuration_0)->{'fromDuration'})($d_1), \Data\DateTime\majData_majDatemajTime_tomajRecord($dt_2))))(function($rec_3) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Maybe_applyMaybe'])->{'apply'})(((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})($GLOBALS['Data_DateTime_DateTime']))(((($GLOBALS['Data_Maybe_bindMaybe'])->{'bind'})(((($GLOBALS['Data_Maybe_applyMaybe'])->{'apply'})(((($GLOBALS['Data_Maybe_applyMaybe'])->{'apply'})(((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})($GLOBALS['Data_Date_exactDate']))((($GLOBALS['Data_Date_Component_boundedEnumYear'])->{'toEnum'})(($rec_3)->{'year'}))))((($GLOBALS['Data_Date_Component_boundedEnumMonth'])->{'toEnum'})(($rec_3)->{'month'}))))((($GLOBALS['Data_Date_Component_boundedEnumDay'])->{'toEnum'})(($rec_3)->{'day'}))))($GLOBALS['Control_Bind_identity']))))(((($GLOBALS['Data_Maybe_applyMaybe'])->{'apply'})(((($GLOBALS['Data_Maybe_applyMaybe'])->{'apply'})(((($GLOBALS['Data_Maybe_applyMaybe'])->{'apply'})(((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})($GLOBALS['Data_Time_Time']))((($GLOBALS['Data_Time_Component_boundedEnumHour'])->{'toEnum'})(($rec_3)->{'hour'}))))((($GLOBALS['Data_Time_Component_boundedEnumMinute'])->{'toEnum'})(($rec_3)->{'minute'}))))((($GLOBALS['Data_Time_Component_boundedEnumSecond'])->{'toEnum'})(($rec_3)->{'second'}))))((($GLOBALS['Data_Time_Component_boundedEnumMillisecond'])->{'toEnum'})(($rec_3)->{'millisecond'})));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_DateTime_adjust'] = __NAMESPACE__ . '\\majData_majDatemajTime_adjust';

