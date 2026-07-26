<?php

namespace Data\Number\Format;

// ALL IMPORTS: Control.Semigroupoid, Data.Number.Format, Data.Ord, Prelude, Prim
// TO REQUIRE: Control.Semigroupoid, Data.Number.Format, Data.Ord, Prelude
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Number.Format/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
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
$GLOBALS['Data_Number_Format_toExponentialNative'] = new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_Number_Format_toFixedNative'] = new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_Number_Format_toPrecisionNative'] = new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_Number_Format_toString'] = new class { public function __invoke(...$args) { return $this; } };


final class Data_Number_Format_Precision { public function __construct(public int $value0) {} }
final class Data_Number_Format_Fixed { public function __construct(public int $value0) {} }
final class Data_Number_Format_Exponential { public function __construct(public int $value0) {} }

// Data_Number_Format_clamp
$GLOBALS['Data_Number_Format_clamp'] = (function() use (&$__fn) {
$__local_var_0_0 = ($GLOBALS['Data_Ord_ordIntImpl'])(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT());
return (function() use ($__local_var_0_0) {
  $__fn = function($low_1 = null, $hi_2 = null, $x_3 = null) use ($__local_var_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $v_4_1 = ($__local_var_0_0)($low_1, $x_3);
  $__t2 = null;;
  if ($v_4_1 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t2 = $x_3;
goto end_branch_2;;
};
  if ($v_4_1 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t2 = $low_1;
goto end_branch_2;;
};
  if ($v_4_1 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t2 = $low_1;
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__local_var_5_2 = $__t2;
  $v_6_4 = ($__local_var_0_0)($hi_2, $__local_var_5_2);
  $__t5 = null;;
  if ($v_6_4 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t5 = $hi_2;
goto end_branch_5;;
};
  if ($v_6_4 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t5 = $hi_2;
goto end_branch_5;;
};
  if ($v_6_4 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t5 = $__local_var_5_2;
goto end_branch_5;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t5 = null;
  end_branch_5:;
  $__res = $__t5;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
})();

// Data_Number_Format_Precision
$GLOBALS['Data_Number_Format_Precision'] = function($value0 = null) {
  $__num = \func_num_args();
  $__res = new \Data\Number\Format\Data_Number_Format_Precision($value0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Number_Format_Fixed
$GLOBALS['Data_Number_Format_Fixed'] = function($value0 = null) {
  $__num = \func_num_args();
  $__res = new \Data\Number\Format\Data_Number_Format_Fixed($value0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Number_Format_Exponential
$GLOBALS['Data_Number_Format_Exponential'] = function($value0 = null) {
  $__num = \func_num_args();
  $__res = new \Data\Number\Format\Data_Number_Format_Exponential($value0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Number_Format_toStringWith
function majData_majNumber_majFormat_tomajStringmajWith($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majNumber_majFormat_tomajStringmajWith';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__t0 = null;;
  if ($v_0 instanceof \Data\Number\Format\Data_Number_Format_Precision) {
$__t0 = ($GLOBALS['Data_Number_Format_toPrecisionNative'])(($v_0)->{'value0'});
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Number\Format\Data_Number_Format_Fixed) {
$__t0 = ($GLOBALS['Data_Number_Format_toFixedNative'])(($v_0)->{'value0'});
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Number\Format\Data_Number_Format_Exponential) {
$__t0 = ($GLOBALS['Data_Number_Format_toExponentialNative'])(($v_0)->{'value0'});
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Number_Format_toStringWith'] = __NAMESPACE__ . '\\majData_majNumber_majFormat_tomajStringmajWith';

// Data_Number_Format_precision_closure
$GLOBALS['Data_Number_Format_precision_closure'] = ($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Number_Format_Precision'], ($GLOBALS['Data_Number_Format_clamp'])(1, 21));

// Data_Number_Format_precision
function majData_majNumber_majFormat_precision($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majNumber_majFormat_precision';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Number_Format_precision_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Number_Format_precision'] = __NAMESPACE__ . '\\majData_majNumber_majFormat_precision';

// Data_Number_Format_fixed_closure
$GLOBALS['Data_Number_Format_fixed_closure'] = ($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Number_Format_Fixed'], ($GLOBALS['Data_Number_Format_clamp'])(0, 20));

// Data_Number_Format_fixed
function majData_majNumber_majFormat_fixed($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majNumber_majFormat_fixed';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Number_Format_fixed_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Number_Format_fixed'] = __NAMESPACE__ . '\\majData_majNumber_majFormat_fixed';

// Data_Number_Format_exponential_closure
$GLOBALS['Data_Number_Format_exponential_closure'] = ($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Number_Format_Exponential'], ($GLOBALS['Data_Number_Format_clamp'])(0, 20));

// Data_Number_Format_exponential
function majData_majNumber_majFormat_exponential($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majNumber_majFormat_exponential';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Number_Format_exponential_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Number_Format_exponential'] = __NAMESPACE__ . '\\majData_majNumber_majFormat_exponential';

