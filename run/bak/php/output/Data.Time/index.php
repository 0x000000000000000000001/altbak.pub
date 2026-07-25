<?php

namespace Data\Time;

// ALL IMPORTS: Control.Apply, Data.Bounded, Data.Enum, Data.Eq, Data.EuclideanRing, Data.Function, Data.Functor, Data.HeytingAlgebra, Data.Int, Data.Maybe, Data.Newtype, Data.Number, Data.Ord, Data.Ordering, Data.Ring, Data.Semigroup, Data.Semiring, Data.Show, Data.Time, Data.Time.Component, Data.Time.Duration, Data.Tuple, Partial.Unsafe, Prelude, Prim
// TO REQUIRE: Control.Apply, Data.Bounded, Data.Enum, Data.Eq, Data.EuclideanRing, Data.Function, Data.Functor, Data.HeytingAlgebra, Data.Int, Data.Maybe, Data.Newtype, Data.Number, Data.Ord, Data.Ordering, Data.Ring, Data.Semigroup, Data.Semiring, Data.Show, Data.Time, Data.Time.Component, Data.Time.Duration, Data.Tuple, Partial.Unsafe, Prelude
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Data.Bounded/index.php';
require_once __DIR__ . '/../Data.Enum/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.EuclideanRing/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.HeytingAlgebra/index.php';
require_once __DIR__ . '/../Data.Int/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Newtype/index.php';
require_once __DIR__ . '/../Data.Number/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ordering/index.php';
require_once __DIR__ . '/../Data.Ring/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Semiring/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
require_once __DIR__ . '/../Data.Time/index.php';
require_once __DIR__ . '/../Data.Time.Component/index.php';
require_once __DIR__ . '/../Data.Time.Duration/index.php';
require_once __DIR__ . '/../Data.Tuple/index.php';
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
\PhpursThunks::$thunks['Data_Time_fromJust'] = function() { $v = function($v_0 = null) {
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
\PhpursThunks::$thunks['Data_Time_negateDuration'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))->{'identity'}))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce')))(($GLOBALS['Data_Time_Duration_negate'] ?? \PhpursThunks::eval('Data_Time_Duration_negate')))))((($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))->{'identity'})); return $v; };
\PhpursThunks::$thunks['Data_Time_Time'] = function() { $v = (function() {
  $__fn = function($value0 = null, $value1 = null, $value2 = null, $value3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = new Phpurs_Data4("Time", $value0, $value1, $value2, $value3);
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Time_showTime'] = function() { $v = (object)["show" => function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))("(Time "))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))("(Hour "))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))((($GLOBALS['Data_Show_showIntImpl'] ?? \PhpursThunks::eval('Data_Show_showIntImpl')))(($v_0)->{'value0'})))(")"))))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))(" "))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))("(Minute "))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))((($GLOBALS['Data_Show_showIntImpl'] ?? \PhpursThunks::eval('Data_Show_showIntImpl')))(($v_0)->{'value1'})))(")"))))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))(" "))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))("(Second "))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))((($GLOBALS['Data_Show_showIntImpl'] ?? \PhpursThunks::eval('Data_Show_showIntImpl')))(($v_0)->{'value2'})))(")"))))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))(" "))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))("(Millisecond "))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))((($GLOBALS['Data_Show_showIntImpl'] ?? \PhpursThunks::eval('Data_Show_showIntImpl')))(($v_0)->{'value3'})))(")"))))(")"))))))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Time_setSecond'] = function() { $v = (function() {
  $__fn = function($s_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data4("Time", ($v_1)->{'value0'}, ($v_1)->{'value1'}, $s_0, ($v_1)->{'value3'});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Time_setMinute'] = function() { $v = (function() {
  $__fn = function($m_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data4("Time", ($v_1)->{'value0'}, $m_0, ($v_1)->{'value2'}, ($v_1)->{'value3'});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Time_setMillisecond'] = function() { $v = (function() {
  $__fn = function($ms_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data4("Time", ($v_1)->{'value0'}, ($v_1)->{'value1'}, ($v_1)->{'value2'}, $ms_0);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Time_setHour'] = function() { $v = (function() {
  $__fn = function($h_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data4("Time", $h_0, ($v_1)->{'value1'}, ($v_1)->{'value2'}, ($v_1)->{'value3'});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Time_second'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ($v_0)->{'value2'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Time_minute'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ($v_0)->{'value1'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Time_millisecond'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ($v_0)->{'value3'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Time_millisToTime'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $hours_1_0 = (($GLOBALS['Data_Number_floor'] ?? \PhpursThunks::eval('Data_Number_floor')))(((($GLOBALS['Data_EuclideanRing_numDiv'] ?? \PhpursThunks::eval('Data_EuclideanRing_numDiv')))($v_0))(3600000.0));
  $minutes_2_1 = (($GLOBALS['Data_Number_floor'] ?? \PhpursThunks::eval('Data_Number_floor')))(((($GLOBALS['Data_EuclideanRing_numDiv'] ?? \PhpursThunks::eval('Data_EuclideanRing_numDiv')))(((($GLOBALS['Data_Ring_numSub'] ?? \PhpursThunks::eval('Data_Ring_numSub')))($v_0))(((($GLOBALS['Data_Semiring_numMul'] ?? \PhpursThunks::eval('Data_Semiring_numMul')))($hours_1_0))(3600000.0))))(60000.0));
  $seconds_3_2 = (($GLOBALS['Data_Number_floor'] ?? \PhpursThunks::eval('Data_Number_floor')))(((($GLOBALS['Data_EuclideanRing_numDiv'] ?? \PhpursThunks::eval('Data_EuclideanRing_numDiv')))(((($GLOBALS['Data_Ring_numSub'] ?? \PhpursThunks::eval('Data_Ring_numSub')))($v_0))(((($GLOBALS['Data_Semiring_numAdd'] ?? \PhpursThunks::eval('Data_Semiring_numAdd')))(((($GLOBALS['Data_Semiring_numMul'] ?? \PhpursThunks::eval('Data_Semiring_numMul')))($hours_1_0))(3600000.0)))(((($GLOBALS['Data_Semiring_numMul'] ?? \PhpursThunks::eval('Data_Semiring_numMul')))($minutes_2_1))(60000.0)))))(1000.0));
  $__local_var_4_3 = ((($GLOBALS['Data_Time_Component_boundedEnumHour'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumHour')))->{'toEnum'})((($GLOBALS['Data_Int_floor'] ?? \PhpursThunks::eval('Data_Int_floor')))($hours_1_0));
  $__t4 = null;;
  if ((is_object($__local_var_4_3) && (($__local_var_4_3)->{'tag'} === "Just"))) {
$__local_var_5_5 = ((($GLOBALS['Data_Time_Component_boundedEnumMinute'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMinute')))->{'toEnum'})((($GLOBALS['Data_Int_floor'] ?? \PhpursThunks::eval('Data_Int_floor')))($minutes_2_1));
$__t6 = null;;
if ((is_object($__local_var_5_5) && (($__local_var_5_5)->{'tag'} === "Just"))) {
$__local_var_6_7 = ((($GLOBALS['Data_Time_Component_boundedEnumSecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumSecond')))->{'toEnum'})((($GLOBALS['Data_Int_floor'] ?? \PhpursThunks::eval('Data_Int_floor')))($seconds_3_2));
$__t8 = null;;
if ((is_object($__local_var_6_7) && (($__local_var_6_7)->{'tag'} === "Just"))) {
$__local_var_7_9 = ((($GLOBALS['Data_Time_Component_boundedEnumMillisecond'] ?? \PhpursThunks::eval('Data_Time_Component_boundedEnumMillisecond')))->{'toEnum'})((($GLOBALS['Data_Int_floor'] ?? \PhpursThunks::eval('Data_Int_floor')))(((($GLOBALS['Data_Ring_numSub'] ?? \PhpursThunks::eval('Data_Ring_numSub')))($v_0))(((($GLOBALS['Data_Semiring_numAdd'] ?? \PhpursThunks::eval('Data_Semiring_numAdd')))(((($GLOBALS['Data_Semiring_numAdd'] ?? \PhpursThunks::eval('Data_Semiring_numAdd')))(((($GLOBALS['Data_Semiring_numMul'] ?? \PhpursThunks::eval('Data_Semiring_numMul')))($hours_1_0))(3600000.0)))(((($GLOBALS['Data_Semiring_numMul'] ?? \PhpursThunks::eval('Data_Semiring_numMul')))($minutes_2_1))(60000.0))))(((($GLOBALS['Data_Semiring_numMul'] ?? \PhpursThunks::eval('Data_Semiring_numMul')))($seconds_3_2))(1000.0)))));
$__t10 = null;;
if ((is_object($__local_var_7_9) && (($__local_var_7_9)->{'tag'} === "Just"))) {
$__t10 = new Phpurs_Data1("Just", new Phpurs_Data4("Time", ($__local_var_4_3)->{'value0'}, ($__local_var_5_5)->{'value0'}, ($__local_var_6_7)->{'value0'}, ($__local_var_7_9)->{'value0'}));
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
$__t6 = new Phpurs_Data0("Nothing");
end_branch_6:;
$__t4 = $__t6;
goto end_branch_4;;
};
  $__t4 = new Phpurs_Data0("Nothing");
  end_branch_4:;
  $__res = ((($GLOBALS['Partial_Unsafe__unsafePartial'] ?? \PhpursThunks::eval('Partial_Unsafe__unsafePartial')))(function($dollar__unused_4 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Time_fromJust'] ?? \PhpursThunks::eval('Data_Time_fromJust'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($__t4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Time_hour'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ($v_0)->{'value0'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Time_timeToMillis'] = function() { $v = function($t_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Semiring_numAdd'] ?? \PhpursThunks::eval('Data_Semiring_numAdd')))(((($GLOBALS['Data_Semiring_numAdd'] ?? \PhpursThunks::eval('Data_Semiring_numAdd')))(((($GLOBALS['Data_Semiring_numAdd'] ?? \PhpursThunks::eval('Data_Semiring_numAdd')))(((($GLOBALS['Data_Semiring_numMul'] ?? \PhpursThunks::eval('Data_Semiring_numMul')))(3600000.0))((($GLOBALS['Data_Int_toNumber'] ?? \PhpursThunks::eval('Data_Int_toNumber')))(($t_0)->{'value0'}))))(((($GLOBALS['Data_Semiring_numMul'] ?? \PhpursThunks::eval('Data_Semiring_numMul')))(60000.0))((($GLOBALS['Data_Int_toNumber'] ?? \PhpursThunks::eval('Data_Int_toNumber')))(($t_0)->{'value1'})))))(((($GLOBALS['Data_Semiring_numMul'] ?? \PhpursThunks::eval('Data_Semiring_numMul')))(1000.0))((($GLOBALS['Data_Int_toNumber'] ?? \PhpursThunks::eval('Data_Int_toNumber')))(($t_0)->{'value2'})))))((($GLOBALS['Data_Int_toNumber'] ?? \PhpursThunks::eval('Data_Int_toNumber')))(($t_0)->{'value3'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Time_eqTime'] = function() { $v = (object)["eq" => (function() {
  $__fn = function($x_0 = null, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_HeytingAlgebra_boolConj'] ?? \PhpursThunks::eval('Data_HeytingAlgebra_boolConj')))(((($GLOBALS['Data_HeytingAlgebra_boolConj'] ?? \PhpursThunks::eval('Data_HeytingAlgebra_boolConj')))(((($GLOBALS['Data_HeytingAlgebra_boolConj'] ?? \PhpursThunks::eval('Data_HeytingAlgebra_boolConj')))(((($GLOBALS['Data_Eq_eqIntImpl'] ?? \PhpursThunks::eval('Data_Eq_eqIntImpl')))(($x_0)->{'value0'}))(($y_1)->{'value0'})))(((($GLOBALS['Data_Eq_eqIntImpl'] ?? \PhpursThunks::eval('Data_Eq_eqIntImpl')))(($x_0)->{'value1'}))(($y_1)->{'value1'}))))(((($GLOBALS['Data_Eq_eqIntImpl'] ?? \PhpursThunks::eval('Data_Eq_eqIntImpl')))(($x_0)->{'value2'}))(($y_1)->{'value2'}))))(((($GLOBALS['Data_Eq_eqIntImpl'] ?? \PhpursThunks::eval('Data_Eq_eqIntImpl')))(($x_0)->{'value3'}))(($y_1)->{'value3'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Data_Time_ordTime'] = function() { $v = (object)["compare" => (function() {
  $__fn = function($x_0 = null, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v_2_0 = (((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->{'compare'})(($x_0)->{'value0'}))(($y_1)->{'value0'});
  $__t5 = null;;
  if ((is_object($v_2_0) && (($v_2_0)->{'tag'} === "LT"))) {
$__t5 = new Phpurs_Data0("LT");
goto end_branch_5;;
};
  if ((is_object($v_2_0) && (($v_2_0)->{'tag'} === "GT"))) {
$__t5 = new Phpurs_Data0("GT");
goto end_branch_5;;
};
  $v1_3_1 = (((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->{'compare'})(($x_0)->{'value1'}))(($y_1)->{'value1'});
  $__t4 = null;;
  if ((is_object($v1_3_1) && (($v1_3_1)->{'tag'} === "LT"))) {
$__t4 = new Phpurs_Data0("LT");
goto end_branch_4;;
};
  if ((is_object($v1_3_1) && (($v1_3_1)->{'tag'} === "GT"))) {
$__t4 = new Phpurs_Data0("GT");
goto end_branch_4;;
};
  $v2_4_2 = (((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->{'compare'})(($x_0)->{'value2'}))(($y_1)->{'value2'});
  $__t3 = null;;
  if ((is_object($v2_4_2) && (($v2_4_2)->{'tag'} === "LT"))) {
$__t3 = new Phpurs_Data0("LT");
goto end_branch_3;;
};
  if ((is_object($v2_4_2) && (($v2_4_2)->{'tag'} === "GT"))) {
$__t3 = new Phpurs_Data0("GT");
goto end_branch_3;;
};
  $__t3 = (((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->{'compare'})(($x_0)->{'value3'}))(($y_1)->{'value3'});
  end_branch_3:;
  $__t4 = $__t3;
  end_branch_4:;
  $__t5 = $__t4;
  end_branch_5:;
  $__res = $__t5;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Time_eqTime'] ?? \PhpursThunks::eval('Data_Time_eqTime'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Time_diff'] = function() { $v = (function() {
  $__fn = function($dictDuration_0 = null, $t1_1 = null, $t2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($dictDuration_0)->{'toDuration'})(((($GLOBALS['Data_Semiring_numAdd'] ?? \PhpursThunks::eval('Data_Semiring_numAdd')))((($GLOBALS['Data_Time_timeToMillis'] ?? \PhpursThunks::eval('Data_Time_timeToMillis')))($t1_1)))((($GLOBALS['Data_Time_negateDuration'] ?? \PhpursThunks::eval('Data_Time_negateDuration')))((($GLOBALS['Data_Time_timeToMillis'] ?? \PhpursThunks::eval('Data_Time_timeToMillis')))($t2_2))));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Time_boundedTime'] = function() { $v = (object)["bottom" => new Phpurs_Data4("Time", 0, 0, 0, 0), "top" => new Phpurs_Data4("Time", 23, 59, 59, 999), "Ord0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Time_ordTime'] ?? \PhpursThunks::eval('Data_Time_ordTime'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Time_maxTime'] = function() { $v = (($GLOBALS['Data_Time_timeToMillis'] ?? \PhpursThunks::eval('Data_Time_timeToMillis')))(new Phpurs_Data4("Time", 23, 59, 59, 999)); return $v; };
\PhpursThunks::$thunks['Data_Time_minTime'] = function() { $v = (($GLOBALS['Data_Time_timeToMillis'] ?? \PhpursThunks::eval('Data_Time_timeToMillis')))(new Phpurs_Data4("Time", 0, 0, 0, 0)); return $v; };
\PhpursThunks::$thunks['Data_Time_adjust'] = function() { $v = (function() {
  $__fn = function($dictDuration_0 = null, $d_1 = null, $t_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $d__prime___3_0 = (($dictDuration_0)->{'fromDuration'})($d_1);
  $wholeDays_4_1 = (($GLOBALS['Data_Number_floor'] ?? \PhpursThunks::eval('Data_Number_floor')))(((($GLOBALS['Data_EuclideanRing_numDiv'] ?? \PhpursThunks::eval('Data_EuclideanRing_numDiv')))((($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce')))($d__prime___3_0)))(86400000.0));
  $msAdjusted_5_2 = ((($GLOBALS['Data_Semiring_numAdd'] ?? \PhpursThunks::eval('Data_Semiring_numAdd')))((($GLOBALS['Data_Time_timeToMillis'] ?? \PhpursThunks::eval('Data_Time_timeToMillis')))($t_2)))(((($GLOBALS['Data_Semiring_numAdd'] ?? \PhpursThunks::eval('Data_Semiring_numAdd')))($d__prime___3_0))((($GLOBALS['Data_Time_negateDuration'] ?? \PhpursThunks::eval('Data_Time_negateDuration')))(((($GLOBALS['Data_Time_Duration_durationDays'] ?? \PhpursThunks::eval('Data_Time_Duration_durationDays')))->{'fromDuration'})($wholeDays_4_1))));
  $__t3 = null;;
  if ((is_object((((($GLOBALS['Data_Ord_ordNumber'] ?? \PhpursThunks::eval('Data_Ord_ordNumber')))->{'compare'})($msAdjusted_5_2))(($GLOBALS['Data_Time_maxTime'] ?? \PhpursThunks::eval('Data_Time_maxTime')))) && (((((($GLOBALS['Data_Ord_ordNumber'] ?? \PhpursThunks::eval('Data_Ord_ordNumber')))->{'compare'})($msAdjusted_5_2))(($GLOBALS['Data_Time_maxTime'] ?? \PhpursThunks::eval('Data_Time_maxTime'))))->{'tag'} === "GT"))) {
$__t3 = 1.0;
goto end_branch_3;;
};
  if ((is_object((((($GLOBALS['Data_Ord_ordNumber'] ?? \PhpursThunks::eval('Data_Ord_ordNumber')))->{'compare'})($msAdjusted_5_2))(($GLOBALS['Data_Time_minTime'] ?? \PhpursThunks::eval('Data_Time_minTime')))) && (((((($GLOBALS['Data_Ord_ordNumber'] ?? \PhpursThunks::eval('Data_Ord_ordNumber')))->{'compare'})($msAdjusted_5_2))(($GLOBALS['Data_Time_minTime'] ?? \PhpursThunks::eval('Data_Time_minTime'))))->{'tag'} === "LT"))) {
$__t3 = ((($GLOBALS['Data_Ring_numSub'] ?? \PhpursThunks::eval('Data_Ring_numSub')))(0.0))(1.0);
goto end_branch_3;;
};
  $__t3 = 0.0;
  end_branch_3:;
  $wrap_6_3 = $__t3;
  $__res = new Phpurs_Data2("Tuple", ((($GLOBALS['Data_Semiring_numAdd'] ?? \PhpursThunks::eval('Data_Semiring_numAdd')))($wholeDays_4_1))($wrap_6_3), (($GLOBALS['Data_Time_millisToTime'] ?? \PhpursThunks::eval('Data_Time_millisToTime')))(((($GLOBALS['Data_Semiring_numAdd'] ?? \PhpursThunks::eval('Data_Semiring_numAdd')))($msAdjusted_5_2))(((($GLOBALS['Data_Semiring_numMul'] ?? \PhpursThunks::eval('Data_Semiring_numMul')))(86400000.0))(((($GLOBALS['Data_Ring_numSub'] ?? \PhpursThunks::eval('Data_Ring_numSub')))(0.0))($wrap_6_3)))));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };























