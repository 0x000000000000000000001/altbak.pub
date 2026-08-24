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
if (!\function_exists(__NAMESPACE__ . '\\phpurs_execute_effect')) {
  function phpurs_execute_effect($val) {
    if (\is_callable($val)) {
      return $val($GLOBALS['Data_Unit_unit']);
    }
    return $val;
  }
}

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };
function majData_majNumber_majFormat_tomajExponentialmajNative(int $v0, $v1 = null): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majNumber_majFormat_tomajExponentialmajNative';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $f = new class { public function __invoke(...$args) { return $this; } };
  return $f($v0, $v1);
}
$GLOBALS['Data_Number_Format_toExponentialNative'] = __NAMESPACE__ . '\\majData_majNumber_majFormat_tomajExponentialmajNative';

function majData_majNumber_majFormat_tomajFixedmajNative(int $v0, $v1 = null): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majNumber_majFormat_tomajFixedmajNative';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $f = new class { public function __invoke(...$args) { return $this; } };
  return $f($v0, $v1);
}
$GLOBALS['Data_Number_Format_toFixedNative'] = __NAMESPACE__ . '\\majData_majNumber_majFormat_tomajFixedmajNative';

function majData_majNumber_majFormat_tomajPrecisionmajNative(int $v0, $v1 = null): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majNumber_majFormat_tomajPrecisionmajNative';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $f = new class { public function __invoke(...$args) { return $this; } };
  return $f($v0, $v1);
}
$GLOBALS['Data_Number_Format_toPrecisionNative'] = __NAMESPACE__ . '\\majData_majNumber_majFormat_tomajPrecisionmajNative';

function majData_majNumber_majFormat_tomajString(float $v0): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majNumber_majFormat_tomajString';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $f = new class { public function __invoke(...$args) { return $this; } };
  return $f($v0);
}
$GLOBALS['Data_Number_Format_toString'] = __NAMESPACE__ . '\\majData_majNumber_majFormat_tomajString';



final class Data_Number_Format_Precision { public $tag = 'Precision'; public function __construct(public int $value0) {} }
final class Data_Number_Format_Fixed { public $tag = 'Fixed'; public function __construct(public int $value0) {} }
final class Data_Number_Format_Exponential { public $tag = 'Exponential'; public function __construct(public int $value0) {} }

// Data_Number_Format_Precision
$GLOBALS['Data_Number_Format_Precision'] = function($value0) {
  $__num = \func_num_args();
  $__res = new \Data\Number\Format\Data_Number_Format_Precision($value0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Number_Format_Fixed
$GLOBALS['Data_Number_Format_Fixed'] = function($value0) {
  $__num = \func_num_args();
  $__res = new \Data\Number\Format\Data_Number_Format_Fixed($value0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Number_Format_Exponential
$GLOBALS['Data_Number_Format_Exponential'] = function($value0) {
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
$GLOBALS['Data_Number_Format_precision_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Number_Format_Precision']))(function($x_0) {
  $__num = \func_num_args();
  $v_1_0 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), 1, $x_0);
  $__t1 = null;;
  if ($v_1_0 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t1 = $x_0;
goto end_branch_1;;
};
  if ($v_1_0 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t1 = 1;
goto end_branch_1;;
};
  if ($v_1_0 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t1 = 1;
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__local_var_1_0 = $__t1;
  $v_2_3 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), 21, $__local_var_1_0);
  $__t4 = null;;
  if ($v_2_3 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t4 = 21;
goto end_branch_4;;
};
  if ($v_2_3 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t4 = 21;
goto end_branch_4;;
};
  if ($v_2_3 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t4 = $__local_var_1_0;
goto end_branch_4;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t4 = null;
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});

// Data_Number_Format_precision
function majData_majNumber_majFormat_precision(int $v_0) {
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
$GLOBALS['Data_Number_Format_fixed_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Number_Format_Fixed']))(function($x_0) {
  $__num = \func_num_args();
  $v_1_0 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), 0, $x_0);
  $__t1 = null;;
  if ($v_1_0 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t1 = $x_0;
goto end_branch_1;;
};
  if ($v_1_0 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t1 = 0;
goto end_branch_1;;
};
  if ($v_1_0 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t1 = 0;
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__local_var_1_0 = $__t1;
  $v_2_3 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), 20, $__local_var_1_0);
  $__t4 = null;;
  if ($v_2_3 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t4 = 20;
goto end_branch_4;;
};
  if ($v_2_3 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t4 = 20;
goto end_branch_4;;
};
  if ($v_2_3 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t4 = $__local_var_1_0;
goto end_branch_4;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t4 = null;
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});

// Data_Number_Format_fixed
function majData_majNumber_majFormat_fixed(int $v_0) {
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
$GLOBALS['Data_Number_Format_exponential_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Number_Format_Exponential']))(function($x_0) {
  $__num = \func_num_args();
  $v_1_0 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), 0, $x_0);
  $__t1 = null;;
  if ($v_1_0 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t1 = $x_0;
goto end_branch_1;;
};
  if ($v_1_0 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t1 = 0;
goto end_branch_1;;
};
  if ($v_1_0 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t1 = 0;
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__local_var_1_0 = $__t1;
  $v_2_3 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), 20, $__local_var_1_0);
  $__t4 = null;;
  if ($v_2_3 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t4 = 20;
goto end_branch_4;;
};
  if ($v_2_3 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t4 = 20;
goto end_branch_4;;
};
  if ($v_2_3 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t4 = $__local_var_1_0;
goto end_branch_4;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t4 = null;
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});

// Data_Number_Format_exponential
function majData_majNumber_majFormat_exponential(int $v_0) {
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

