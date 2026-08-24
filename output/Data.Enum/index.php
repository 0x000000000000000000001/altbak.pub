<?php

namespace Data\Enum;

// ALL IMPORTS: Control.Alternative, Control.Apply, Control.Bind, Control.MonadPlus, Control.Semigroupoid, Data.Boolean, Data.Bounded, Data.Either, Data.Enum, Data.Eq, Data.Function, Data.Functor, Data.HeytingAlgebra, Data.Maybe, Data.Newtype, Data.Ord, Data.Ordering, Data.Ring, Data.Semigroup, Data.Semiring, Data.Show, Data.Tuple, Data.Unfoldable, Data.Unfoldable1, Data.Unit, Partial.Unsafe, Prelude, Prim
// TO REQUIRE: Control.Alternative, Control.Apply, Control.Bind, Control.MonadPlus, Control.Semigroupoid, Data.Boolean, Data.Bounded, Data.Either, Data.Enum, Data.Eq, Data.Function, Data.Functor, Data.HeytingAlgebra, Data.Maybe, Data.Newtype, Data.Ord, Data.Ordering, Data.Ring, Data.Semigroup, Data.Semiring, Data.Show, Data.Tuple, Data.Unfoldable, Data.Unfoldable1, Data.Unit, Partial.Unsafe, Prelude
require_once __DIR__ . '/../Control.Alternative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.MonadPlus/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Boolean/index.php';
require_once __DIR__ . '/../Data.Bounded/index.php';
require_once __DIR__ . '/../Data.Either/index.php';
require_once __DIR__ . '/../Data.Enum/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.HeytingAlgebra/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Newtype/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ordering/index.php';
require_once __DIR__ . '/../Data.Ring/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Semiring/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
require_once __DIR__ . '/../Data.Tuple/index.php';
require_once __DIR__ . '/../Data.Unfoldable/index.php';
require_once __DIR__ . '/../Data.Unfoldable1/index.php';
require_once __DIR__ . '/../Data.Unit/index.php';
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
if (!\function_exists(__NAMESPACE__ . '\\phpurs_execute_effect')) {
  function phpurs_execute_effect($val) {
    if (\is_callable($val)) {
      return $val($GLOBALS['Data_Unit_unit']);
    }
    return $val;
  }
}

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };
$ffi_Data_Enum = \call_user_func(function() {
  $exports = [];
$toCharCode = function($c) {
    if (\function_exists('mb_ord')) {
        return \mb_ord($c, "UTF-8");
    }
    // Very rudimentary fallback
    return \ord($c);
};

$fromCharCode = function($c) {
    if (\function_exists('mb_chr')) {
        return \mb_chr($c, "UTF-8");
    }
    return \chr($c);
};

$exports['toCharCode'] = $toCharCode;
$exports['fromCharCode'] = $fromCharCode;
return $exports;
  return $exports;
});
function majData_majEnum_frommajCharmajCode(int $v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majEnum_frommajCharmajCode';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Data_Enum;
  $f = (\array_key_exists('fromCharCode', $ffi_Data_Enum) ? $ffi_Data_Enum['fromCharCode'] : new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Data_Enum_fromCharCode'] = __NAMESPACE__ . '\\majData_majEnum_frommajCharmajCode';

function majData_majEnum_tomajCharmajCode($v0): int|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majEnum_tomajCharmajCode';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Data_Enum;
  $f = (\array_key_exists('toCharCode', $ffi_Data_Enum) ? $ffi_Data_Enum['toCharCode'] : new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Data_Enum_toCharCode'] = __NAMESPACE__ . '\\majData_majEnum_tomajCharmajCode';





// Data_Enum_Cardinality
function majData_majEnum_majCardinality($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEnum_majCardinality';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Enum_Cardinality'] = __NAMESPACE__ . '\\majData_majEnum_majCardinality';

// Data_Enum_toEnum
function majData_majEnum_tomajEnum($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEnum_tomajEnum';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'toEnum'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Enum_toEnum'] = __NAMESPACE__ . '\\majData_majEnum_tomajEnum';

// Data_Enum_succ
function majData_majEnum_succ($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEnum_succ';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'succ'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Enum_succ'] = __NAMESPACE__ . '\\majData_majEnum_succ';

// Data_Enum_upFromIncluding
function majData_majEnum_upmajFrommajIncluding($dictEnum_0, $dictUnfoldable1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEnum_upmajFrommajIncluding';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($dictUnfoldable1_1)->{'unfoldr1'})(function($x_2) use ($dictEnum_0) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple($x_2, (($dictEnum_0)->{'succ'})($x_2));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Enum_upFromIncluding'] = __NAMESPACE__ . '\\majData_majEnum_upmajFrommajIncluding';

// Data_Enum_showCardinality
$GLOBALS['Data_Enum_showCardinality'] = (object)["show" => function($v_0) {
  $__num = \func_num_args();
  $__res = (("(Cardinality " . \Data\Show\majData_majShow_showmajIntmajImpl($v_0)) . ")");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Enum_pred
function majData_majEnum_pred($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEnum_pred';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'pred'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Enum_pred'] = __NAMESPACE__ . '\\majData_majEnum_pred';

// Data_Enum_ordCardinality
$GLOBALS['Data_Enum_ordCardinality'] = $GLOBALS['Data_Ord_ordInt'];

// Data_Enum_newtypeCardinality
$GLOBALS['Data_Enum_newtypeCardinality'] = (object)["Coercible0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Enum_fromEnum
function majData_majEnum_frommajEnum($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEnum_frommajEnum';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'fromEnum'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Enum_fromEnum'] = __NAMESPACE__ . '\\majData_majEnum_frommajEnum';

// Data_Enum_toEnumWithDefaults
function majData_majEnum_tomajEnummajWithmajDefaults($dictBoundedEnum_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEnum_tomajEnummajWithmajDefaults';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $bottom2_1_0 = ((($dictBoundedEnum_0)->{'Bounded0'})(null))->{'bottom'};
  $__res = function($low_2) use ($bottom2_1_0, $dictBoundedEnum_0) {
  $__num = \func_num_args();
  $__res = function($high_3) use ($bottom2_1_0, $dictBoundedEnum_0, $low_2) {
  $__num = \func_num_args();
  $__res = function($x_4) use ($bottom2_1_0, $dictBoundedEnum_0, $high_3, $low_2) {
  $__num = \func_num_args();
  $v_5_1 = (($dictBoundedEnum_0)->{'toEnum'})($x_4);
  $__t2 = null;;
  if ($v_5_1 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t2 = ($v_5_1)->{'value0'};
goto end_branch_2;;
};
  if ($v_5_1 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t3 = null;;
if (($x_4 < (($dictBoundedEnum_0)->{'fromEnum'})($bottom2_1_0))) {
$__t3 = $low_2;
goto end_branch_3;;
};
$__t3 = $high_3;
end_branch_3:;
$__t2 = $__t3;
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Enum_toEnumWithDefaults'] = __NAMESPACE__ . '\\majData_majEnum_tomajEnummajWithmajDefaults';

// Data_Enum_eqCardinality
$GLOBALS['Data_Enum_eqCardinality'] = $GLOBALS['Data_Eq_eqInt'];

// Data_Enum_enumUnit
$GLOBALS['Data_Enum_enumUnit'] = (object)["succ" => function($v_0) {
  $__num = \func_num_args();
  $__res = new \Data\Maybe\Data_Maybe_Nothing();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "pred" => function($v_0) {
  $__num = \func_num_args();
  $__res = new \Data\Maybe\Data_Maybe_Nothing();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Ord0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Ord_ordUnit'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Enum_enumTuple
function majData_majEnum_enummajTuple($dictEnum_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEnum_enummajTuple';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictEnum_0)->{'Ord0'})(null);
  $__local_var_2_1 = (($__local_var_1_0)->{'Eq0'})(null);
  $__res = function($dictBoundedEnum_3) use ($__local_var_1_0, $__local_var_2_1, $dictEnum_0) {
  $__num = \func_num_args();
  $Bounded0_4_2 = (($dictBoundedEnum_3)->{'Bounded0'})(null);
  $Enum1_5_3 = (($dictBoundedEnum_3)->{'Enum1'})(null);
  $__local_var_6_4 = (((($dictBoundedEnum_3)->{'Enum1'})(null))->{'Ord0'})(null);
  $__local_var_7_5 = (($__local_var_6_4)->{'Eq0'})(null);
  $eqTuple2_7_5 = (object)["eq" => function($x_8) use ($__local_var_2_1, $__local_var_7_5) {
  $__num = \func_num_args();
  $__res = function($y_9) use ($__local_var_2_1, $__local_var_7_5, $x_8) {
  $__num = \func_num_args();
  $__res = (((($__local_var_2_1)->{'eq'})(($x_8)->{'value0'}))(($y_9)->{'value0'}) && ((($__local_var_7_5)->{'eq'})(($x_8)->{'value1'}))(($y_9)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $ordTuple1_6_4 = (object)["compare" => function($x_8) use ($__local_var_1_0, $__local_var_6_4) {
  $__num = \func_num_args();
  $__res = function($y_9) use ($__local_var_1_0, $__local_var_6_4, $x_8) {
  $__num = \func_num_args();
  $v_10_7 = ((($__local_var_1_0)->{'compare'})(($x_8)->{'value0'}))(($y_9)->{'value0'});
  $__t8 = null;;
  if ($v_10_7 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t8 = new \Data\Ordering\Data_Ordering_LT();
goto end_branch_8;;
};
  if ($v_10_7 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t8 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_8;;
};
  $__t8 = ((($__local_var_6_4)->{'compare'})(($x_8)->{'value1'}))(($y_9)->{'value1'});
  end_branch_8:;
  $__res = $__t8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq0" => function($_dollar___unused_8) use ($eqTuple2_7_5) {
  $__num = \func_num_args();
  $__res = $eqTuple2_7_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["succ" => function($v_7) use ($Bounded0_4_2, $Enum1_5_3, $dictEnum_0) {
  $__num = \func_num_args();
  $__local_var_8_10 = ($Bounded0_4_2)->{'bottom'};
  $__local_var_8_10 = function($a_9) use ($__local_var_8_10) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_9, $__local_var_8_10);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__local_var_9_12 = (($dictEnum_0)->{'succ'})(($v_7)->{'value0'});
  $__t13 = null;;
  if ($__local_var_9_12 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t13 = new \Data\Maybe\Data_Maybe_Just(($__local_var_8_10)(($__local_var_9_12)->{'value0'}));
goto end_branch_13;;
};
  $__t13 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_13:;
  $__local_var_8_10 = $__t13;
  $__local_var_9_15 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Maybe_Just']))(($GLOBALS['Data_Tuple_Tuple'])(($v_7)->{'value0'}));
  $__local_var_10_16 = (($Enum1_5_3)->{'succ'})(($v_7)->{'value1'});
  $__t17 = null;;
  if ($__local_var_10_16 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t17 = $__local_var_8_10;
goto end_branch_17;;
};
  if ($__local_var_10_16 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t17 = ($__local_var_9_15)(($__local_var_10_16)->{'value0'});
goto end_branch_17;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t17 = null;
  end_branch_17:;
  $__res = $__t17;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "pred" => function($v_7) use ($Bounded0_4_2, $Enum1_5_3, $dictEnum_0) {
  $__num = \func_num_args();
  $__local_var_8_18 = ($Bounded0_4_2)->{'top'};
  $__local_var_8_18 = function($a_9) use ($__local_var_8_18) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_9, $__local_var_8_18);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__local_var_9_20 = (($dictEnum_0)->{'pred'})(($v_7)->{'value0'});
  $__t21 = null;;
  if ($__local_var_9_20 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t21 = new \Data\Maybe\Data_Maybe_Just(($__local_var_8_18)(($__local_var_9_20)->{'value0'}));
goto end_branch_21;;
};
  $__t21 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_21:;
  $__local_var_8_18 = $__t21;
  $__local_var_9_23 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Maybe_Just']))(($GLOBALS['Data_Tuple_Tuple'])(($v_7)->{'value0'}));
  $__local_var_10_24 = (($Enum1_5_3)->{'pred'})(($v_7)->{'value1'});
  $__t25 = null;;
  if ($__local_var_10_24 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t25 = $__local_var_8_18;
goto end_branch_25;;
};
  if ($__local_var_10_24 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t25 = ($__local_var_9_23)(($__local_var_10_24)->{'value0'});
goto end_branch_25;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t25 = null;
  end_branch_25:;
  $__res = $__t25;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Ord0" => function($_dollar___unused_7) use ($ordTuple1_6_4) {
  $__num = \func_num_args();
  $__res = $ordTuple1_6_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Enum_enumTuple'] = __NAMESPACE__ . '\\majData_majEnum_enummajTuple';

// Data_Enum_enumOrdering
$GLOBALS['Data_Enum_enumOrdering'] = (object)["succ" => function($v_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v_0 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t0 = new \Data\Maybe\Data_Maybe_Just(new \Data\Ordering\Data_Ordering_EQ());
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t0 = new \Data\Maybe\Data_Maybe_Just(new \Data\Ordering\Data_Ordering_GT());
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t0 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "pred" => function($v_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v_0 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t1 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_1;;
};
  if ($v_0 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t1 = new \Data\Maybe\Data_Maybe_Just(new \Data\Ordering\Data_Ordering_LT());
goto end_branch_1;;
};
  if ($v_0 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t1 = new \Data\Maybe\Data_Maybe_Just(new \Data\Ordering\Data_Ordering_EQ());
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Ord0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Ord_ordOrdering'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Enum_enumMaybe
function majData_majEnum_enummajMaybe($dictBoundedEnum_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEnum_enummajMaybe';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Bounded0_1_0 = (($dictBoundedEnum_0)->{'Bounded0'})(null);
  $Enum1_2_1 = (($dictBoundedEnum_0)->{'Enum1'})(null);
  $__local_var_3_2 = (((($dictBoundedEnum_0)->{'Enum1'})(null))->{'Ord0'})(null);
  $__local_var_4_3 = (($__local_var_3_2)->{'Eq0'})(null);
  $eqMaybe1_4_3 = (object)["eq" => function($x_5) use ($__local_var_4_3) {
  $__num = \func_num_args();
  $__res = function($y_6) use ($__local_var_4_3, $x_5) {
  $__num = \func_num_args();
  $__t4 = null;;
  if ($x_5 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t4 = $y_6 instanceof \Data\Maybe\Data_Maybe_Nothing;
goto end_branch_4;;
};
  $__t4 = ($x_5 instanceof \Data\Maybe\Data_Maybe_Just && ($y_6 instanceof \Data\Maybe\Data_Maybe_Just && ((($__local_var_4_3)->{'eq'})(($x_5)->{'value0'}))(($y_6)->{'value0'})));
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $ordMaybe_3_2 = (object)["compare" => function($x_5) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__res = function($y_6) use ($__local_var_3_2, $x_5) {
  $__num = \func_num_args();
  $__t6 = null;;
  if ($x_5 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t7 = null;;
if ($y_6 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t7 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_7;;
};
$__t7 = new \Data\Ordering\Data_Ordering_LT();
end_branch_7:;
$__t6 = $__t7;
goto end_branch_6;;
};
  if ($y_6 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t6 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_6;;
};
  if (($x_5 instanceof \Data\Maybe\Data_Maybe_Just && $y_6 instanceof \Data\Maybe\Data_Maybe_Just)) {
$__t6 = ((($__local_var_3_2)->{'compare'})(($x_5)->{'value0'}))(($y_6)->{'value0'});
goto end_branch_6;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t6 = null;
  end_branch_6:;
  $__res = $__t6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq0" => function($_dollar___unused_5) use ($eqMaybe1_4_3) {
  $__num = \func_num_args();
  $__res = $eqMaybe1_4_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["succ" => function($v_4) use ($Bounded0_1_0, $Enum1_2_1) {
  $__num = \func_num_args();
  $__t9 = null;;
  if ($v_4 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t9 = new \Data\Maybe\Data_Maybe_Just(new \Data\Maybe\Data_Maybe_Just(($Bounded0_1_0)->{'bottom'}));
goto end_branch_9;;
};
  if ($v_4 instanceof \Data\Maybe\Data_Maybe_Just) {
$__local_var_5_10 = (($Enum1_2_1)->{'succ'})(($v_4)->{'value0'});
$__t11 = null;;
if ($__local_var_5_10 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t11 = new \Data\Maybe\Data_Maybe_Just(new \Data\Maybe\Data_Maybe_Just(($__local_var_5_10)->{'value0'}));
goto end_branch_11;;
};
$__t11 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_11:;
$__t9 = $__t11;
goto end_branch_9;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t9 = null;
  end_branch_9:;
  $__res = $__t9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "pred" => function($v_4) use ($Enum1_2_1) {
  $__num = \func_num_args();
  $__t12 = null;;
  if ($v_4 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t12 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_12;;
};
  if ($v_4 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t12 = new \Data\Maybe\Data_Maybe_Just((($Enum1_2_1)->{'pred'})(($v_4)->{'value0'}));
goto end_branch_12;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t12 = null;
  end_branch_12:;
  $__res = $__t12;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Ord0" => function($_dollar___unused_4) use ($ordMaybe_3_2) {
  $__num = \func_num_args();
  $__res = $ordMaybe_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Enum_enumMaybe'] = __NAMESPACE__ . '\\majData_majEnum_enummajMaybe';

// Data_Enum_enumInt
$GLOBALS['Data_Enum_enumInt'] = (object)["succ" => function($n_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if (($n_0 < $GLOBALS['Data_Bounded_topInt'])) {
$__t0 = new \Data\Maybe\Data_Maybe_Just(($n_0 + 1));
goto end_branch_0;;
};
  $__t0 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "pred" => function($n_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if (($n_0 > $GLOBALS['Data_Bounded_bottomInt'])) {
$__t1 = new \Data\Maybe\Data_Maybe_Just(($n_0 - 1));
goto end_branch_1;;
};
  $__t1 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Ord0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Ord_ordInt'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Enum_enumFromTo
function majData_majEnum_enummajFrommajTo($dictEnum_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEnum_enummajFrommajTo';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Ord0_1_0 = (($dictEnum_0)->{'Ord0'})(null);
  $Eq0_2_1 = (($Ord0_1_0)->{'Eq0'})(null);
  $Ord01_3_2 = (($dictEnum_0)->{'Ord0'})(null);
  $__res = function($dictUnfoldable1_4) use ($Eq0_2_1, $Ord01_3_2, $Ord0_1_0, $dictEnum_0) {
  $__num = \func_num_args();
  $__res = function($v_5) use ($Eq0_2_1, $Ord01_3_2, $Ord0_1_0, $dictEnum_0, $dictUnfoldable1_4) {
  $__num = \func_num_args();
  $__res = function($v1_6) use ($Eq0_2_1, $Ord01_3_2, $Ord0_1_0, $dictEnum_0, $dictUnfoldable1_4, $v_5) {
  $__num = \func_num_args();
  $__t6 = null;;
  if (((($Eq0_2_1)->{'eq'})($v_5))($v1_6)) {
$__t6 = ((($dictUnfoldable1_4)->{'unfoldr1'})(function($i_7) use ($v_5) {
  $__num = \func_num_args();
  $__t7 = null;;
  if (($i_7 <= 0)) {
$__t7 = new \Data\Tuple\Data_Tuple_Tuple($v_5, new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_7;;
};
  $__t7 = new \Data\Tuple\Data_Tuple_Tuple($v_5, new \Data\Maybe\Data_Maybe_Just(($i_7 - 1)));
  end_branch_7:;
  $__res = $__t7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(0);
goto end_branch_6;;
};
  if (((($Ord01_3_2)->{'compare'})($v_5))($v1_6) instanceof \Data\Ordering\Data_Ordering_LT) {
$__t6 = ((($dictUnfoldable1_4)->{'unfoldr1'})(function($a_7) use ($Ord0_1_0, $dictEnum_0, $v1_6) {
  $__num = \func_num_args();
  $__local_var_8_8 = (($dictEnum_0)->{'succ'})($a_7);
  $__t9 = null;;
  if ($__local_var_8_8 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t10 = null;;
if (( ! ((($Ord0_1_0)->{'compare'})(($__local_var_8_8)->{'value0'}))($v1_6) instanceof \Data\Ordering\Data_Ordering_GT)) {
$__t10 = new \Data\Maybe\Data_Maybe_Just(($__local_var_8_8)->{'value0'});
goto end_branch_10;;
};
$__t10 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_10:;
$__t9 = $__t10;
goto end_branch_9;;
};
  if ($__local_var_8_8 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t9 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_9;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t9 = null;
  end_branch_9:;
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_7, $__t9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_5);
goto end_branch_6;;
};
  $__t6 = ((($dictUnfoldable1_4)->{'unfoldr1'})(function($a_7) use ($Ord0_1_0, $dictEnum_0, $v1_6) {
  $__num = \func_num_args();
  $__local_var_8_3 = (($dictEnum_0)->{'pred'})($a_7);
  $__t4 = null;;
  if ($__local_var_8_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t5 = null;;
if (( ! ((($Ord0_1_0)->{'compare'})(($__local_var_8_3)->{'value0'}))($v1_6) instanceof \Data\Ordering\Data_Ordering_LT)) {
$__t5 = new \Data\Maybe\Data_Maybe_Just(($__local_var_8_3)->{'value0'});
goto end_branch_5;;
};
$__t5 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_5:;
$__t4 = $__t5;
goto end_branch_4;;
};
  if ($__local_var_8_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t4 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_4;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t4 = null;
  end_branch_4:;
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_7, $__t4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_5);
  end_branch_6:;
  $__res = $__t6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Enum_enumFromTo'] = __NAMESPACE__ . '\\majData_majEnum_enummajFrommajTo';

// Data_Enum_enumFromThenTo
function majData_majEnum_enummajFrommajThenmajTo($dictUnfoldable_0, $dictFunctor_1 = null, $dictBoundedEnum_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEnum_enummajFrommajThenmajTo';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $toEnum1_3_0 = ($dictBoundedEnum_2)->{'toEnum'};
  $__res = function($a_4) use ($dictBoundedEnum_2, $dictFunctor_1, $dictUnfoldable_0, $toEnum1_3_0) {
  $__num = \func_num_args();
  $__res = function($b_5) use ($a_4, $dictBoundedEnum_2, $dictFunctor_1, $dictUnfoldable_0, $toEnum1_3_0) {
  $__num = \func_num_args();
  $__res = function($c_6) use ($a_4, $b_5, $dictBoundedEnum_2, $dictFunctor_1, $dictUnfoldable_0, $toEnum1_3_0) {
  $__num = \func_num_args();
  $a_prime__7_1 = (($dictBoundedEnum_2)->{'fromEnum'})($a_4);
  $__local_var_8_3 = ((($dictBoundedEnum_2)->{'fromEnum'})($b_5) - $a_prime__7_1);
  $__local_var_9_4 = (($dictBoundedEnum_2)->{'fromEnum'})($c_6);
  $__res = ((($dictFunctor_1)->{'map'})((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_8) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ($v_8 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t2 = ($v_8)->{'value0'};
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($toEnum1_3_0)))(((($dictUnfoldable_0)->{'unfoldr'})(function($e_10) use ($__local_var_8_3, $__local_var_9_4) {
  $__num = \func_num_args();
  $__t5 = null;;
  if (($e_10 <= $__local_var_9_4)) {
$__t5 = new \Data\Maybe\Data_Maybe_Just(new \Data\Tuple\Data_Tuple_Tuple($e_10, ($e_10 + $__local_var_8_3)));
goto end_branch_5;;
};
  $__t5 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_5:;
  $__res = $__t5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($a_prime__7_1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Enum_enumFromThenTo'] = __NAMESPACE__ . '\\majData_majEnum_enummajFrommajThenmajTo';

// Data_Enum_enumEither
function majData_majEnum_enummajEither($dictBoundedEnum_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEnum_enummajEither';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Enum1_1_0 = (($dictBoundedEnum_0)->{'Enum1'})(null);
  $Bounded0_2_1 = (($dictBoundedEnum_0)->{'Bounded0'})(null);
  $__local_var_3_2 = (((($dictBoundedEnum_0)->{'Enum1'})(null))->{'Ord0'})(null);
  $__local_var_4_3 = (($__local_var_3_2)->{'Eq0'})(null);
  $__res = function($dictBoundedEnum1_5) use ($Bounded0_2_1, $Enum1_1_0, $__local_var_3_2, $__local_var_4_3) {
  $__num = \func_num_args();
  $Bounded01_6_4 = (($dictBoundedEnum1_5)->{'Bounded0'})(null);
  $Enum11_7_5 = (($dictBoundedEnum1_5)->{'Enum1'})(null);
  $__local_var_8_6 = (((($dictBoundedEnum1_5)->{'Enum1'})(null))->{'Ord0'})(null);
  $__local_var_9_7 = (($__local_var_8_6)->{'Eq0'})(null);
  $eqEither2_9_7 = (object)["eq" => function($x_10) use ($__local_var_4_3, $__local_var_9_7) {
  $__num = \func_num_args();
  $__res = function($y_11) use ($__local_var_4_3, $__local_var_9_7, $x_10) {
  $__num = \func_num_args();
  $__t8 = null;;
  if ($x_10 instanceof \Data\Either\Data_Either_Left) {
$__t8 = ($y_11 instanceof \Data\Either\Data_Either_Left && ((($__local_var_4_3)->{'eq'})(($x_10)->{'value0'}))(($y_11)->{'value0'}));
goto end_branch_8;;
};
  $__t8 = ($x_10 instanceof \Data\Either\Data_Either_Right && ($y_11 instanceof \Data\Either\Data_Either_Right && ((($__local_var_9_7)->{'eq'})(($x_10)->{'value0'}))(($y_11)->{'value0'})));
  end_branch_8:;
  $__res = $__t8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $ordEither1_8_6 = (object)["compare" => function($x_10) use ($__local_var_3_2, $__local_var_8_6) {
  $__num = \func_num_args();
  $__res = function($y_11) use ($__local_var_3_2, $__local_var_8_6, $x_10) {
  $__num = \func_num_args();
  $__t10 = null;;
  if ($x_10 instanceof \Data\Either\Data_Either_Left) {
$__t11 = null;;
if ($y_11 instanceof \Data\Either\Data_Either_Left) {
$__t11 = ((($__local_var_3_2)->{'compare'})(($x_10)->{'value0'}))(($y_11)->{'value0'});
goto end_branch_11;;
};
$__t11 = new \Data\Ordering\Data_Ordering_LT();
end_branch_11:;
$__t10 = $__t11;
goto end_branch_10;;
};
  if ($y_11 instanceof \Data\Either\Data_Either_Left) {
$__t10 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_10;;
};
  if (($x_10 instanceof \Data\Either\Data_Either_Right && $y_11 instanceof \Data\Either\Data_Either_Right)) {
$__t10 = ((($__local_var_8_6)->{'compare'})(($x_10)->{'value0'}))(($y_11)->{'value0'});
goto end_branch_10;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t10 = null;
  end_branch_10:;
  $__res = $__t10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq0" => function($_dollar___unused_10) use ($eqEither2_9_7) {
  $__num = \func_num_args();
  $__res = $eqEither2_9_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["succ" => function($v_9) use ($Bounded01_6_4, $Enum11_7_5, $Enum1_1_0) {
  $__num = \func_num_args();
  $__t13 = null;;
  if ($v_9 instanceof \Data\Either\Data_Either_Left) {
$__local_var_10_14 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Maybe_Just']))($GLOBALS['Data_Either_Left']);
$__local_var_11_15 = (($Enum1_1_0)->{'succ'})(($v_9)->{'value0'});
$__t16 = null;;
if ($__local_var_11_15 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t16 = new \Data\Maybe\Data_Maybe_Just(new \Data\Either\Data_Either_Right(($Bounded01_6_4)->{'bottom'}));
goto end_branch_16;;
};
if ($__local_var_11_15 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t16 = ($__local_var_10_14)(($__local_var_11_15)->{'value0'});
goto end_branch_16;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t16 = null;
end_branch_16:;
$__t13 = $__t16;
goto end_branch_13;;
};
  if ($v_9 instanceof \Data\Either\Data_Either_Right) {
$__local_var_10_17 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Maybe_Just']))($GLOBALS['Data_Either_Right']);
$__local_var_11_18 = (($Enum11_7_5)->{'succ'})(($v_9)->{'value0'});
$__t19 = null;;
if ($__local_var_11_18 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t19 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_19;;
};
if ($__local_var_11_18 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t19 = ($__local_var_10_17)(($__local_var_11_18)->{'value0'});
goto end_branch_19;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t19 = null;
end_branch_19:;
$__t13 = $__t19;
goto end_branch_13;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t13 = null;
  end_branch_13:;
  $__res = $__t13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "pred" => function($v_9) use ($Bounded0_2_1, $Enum11_7_5, $Enum1_1_0) {
  $__num = \func_num_args();
  $__t20 = null;;
  if ($v_9 instanceof \Data\Either\Data_Either_Left) {
$__local_var_10_21 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Maybe_Just']))($GLOBALS['Data_Either_Left']);
$__local_var_11_22 = (($Enum1_1_0)->{'pred'})(($v_9)->{'value0'});
$__t23 = null;;
if ($__local_var_11_22 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t23 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_23;;
};
if ($__local_var_11_22 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t23 = ($__local_var_10_21)(($__local_var_11_22)->{'value0'});
goto end_branch_23;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t23 = null;
end_branch_23:;
$__t20 = $__t23;
goto end_branch_20;;
};
  if ($v_9 instanceof \Data\Either\Data_Either_Right) {
$__local_var_10_24 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Maybe_Just']))($GLOBALS['Data_Either_Right']);
$__local_var_11_25 = (($Enum11_7_5)->{'pred'})(($v_9)->{'value0'});
$__t26 = null;;
if ($__local_var_11_25 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t26 = new \Data\Maybe\Data_Maybe_Just(new \Data\Either\Data_Either_Left(($Bounded0_2_1)->{'top'}));
goto end_branch_26;;
};
if ($__local_var_11_25 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t26 = ($__local_var_10_24)(($__local_var_11_25)->{'value0'});
goto end_branch_26;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t26 = null;
end_branch_26:;
$__t20 = $__t26;
goto end_branch_20;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t20 = null;
  end_branch_20:;
  $__res = $__t20;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Ord0" => function($_dollar___unused_9) use ($ordEither1_8_6) {
  $__num = \func_num_args();
  $__res = $ordEither1_8_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Enum_enumEither'] = __NAMESPACE__ . '\\majData_majEnum_enummajEither';

// Data_Enum_enumBoolean
$GLOBALS['Data_Enum_enumBoolean'] = (object)["succ" => function($v_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if (( ! $v_0)) {
$__t0 = new \Data\Maybe\Data_Maybe_Just(true);
goto end_branch_0;;
};
  $__t0 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "pred" => function($v_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v_0) {
$__t1 = new \Data\Maybe\Data_Maybe_Just(false);
goto end_branch_1;;
};
  $__t1 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Ord0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Ord_ordBoolean'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Enum_downFromIncluding
function majData_majEnum_downmajFrommajIncluding($dictEnum_0, $dictUnfoldable1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEnum_downmajFrommajIncluding';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($dictUnfoldable1_1)->{'unfoldr1'})(function($x_2) use ($dictEnum_0) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple($x_2, (($dictEnum_0)->{'pred'})($x_2));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Enum_downFromIncluding'] = __NAMESPACE__ . '\\majData_majEnum_downmajFrommajIncluding';

// Data_Enum_downFrom
function majData_majEnum_downmajFrom($dictEnum_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEnum_downmajFrom';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $pred1_1_0 = ($dictEnum_0)->{'pred'};
  $__res = function($dictUnfoldable_2) use ($pred1_1_0) {
  $__num = \func_num_args();
  $__res = (($dictUnfoldable_2)->{'unfoldr'})((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v1_3) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v1_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = new \Data\Maybe\Data_Maybe_Just(new \Data\Tuple\Data_Tuple_Tuple(($v1_3)->{'value0'}, ($v1_3)->{'value0'}));
goto end_branch_1;;
};
  $__t1 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($pred1_1_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Enum_downFrom'] = __NAMESPACE__ . '\\majData_majEnum_downmajFrom';

// Data_Enum_upFrom
function majData_majEnum_upmajFrom($dictEnum_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEnum_upmajFrom';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $succ1_1_0 = ($dictEnum_0)->{'succ'};
  $__res = function($dictUnfoldable_2) use ($succ1_1_0) {
  $__num = \func_num_args();
  $__res = (($dictUnfoldable_2)->{'unfoldr'})((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v1_3) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v1_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = new \Data\Maybe\Data_Maybe_Just(new \Data\Tuple\Data_Tuple_Tuple(($v1_3)->{'value0'}, ($v1_3)->{'value0'}));
goto end_branch_1;;
};
  $__t1 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($succ1_1_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Enum_upFrom'] = __NAMESPACE__ . '\\majData_majEnum_upmajFrom';

// Data_Enum_defaultToEnum
function majData_majEnum_defaultmajTomajEnum($dictBounded_0, $dictEnum_1 = null, $i_prime__2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEnum_defaultmajTomajEnum';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $go__go_3_0 = null;
  $go__go_3_0 = (function() use ($dictEnum_1, &$go__go_3_0) {
  $__fn = function(int $i_4, $x_5 = null) use ($dictEnum_1, &$go__go_3_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_3_0_0_i_4 = $i_4;
  $__tco_var_go__go_3_0_0_x_5 = $x_5;
  tco_loop_go__go_3_0_0:;
  $i_4 = $__tco_var_go__go_3_0_0_i_4;
  $x_5 = $__tco_var_go__go_3_0_0_x_5;
  $__t4 = null;;
  switch ($i_4) {
case 0:
$__t4 = new \Data\Maybe\Data_Maybe_Just($x_5);
goto end_branch_4;;
break;
default:
;
break;
};
  $v_6_0 = (($dictEnum_1)->{'succ'})($x_5);
  $__t1 = null;;
  if ($v_6_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__tco_2 = ($i_4 - 1);
$__tco_3 = ($v_6_0)->{'value0'};
$__tco_var_go__go_3_0_0_i_4 = $__tco_2;
$__tco_var_go__go_3_0_0_x_5 = $__tco_3;
goto tco_loop_go__go_3_0_0;;
$__t1 = null;
goto end_branch_1;;
};
  if ($v_6_0 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t1 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__t4 = $__t1;
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__t1 = null;;
  if (($i_prime__2 < 0)) {
$__t1 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_1;;
};
  $__t1 = (($go__go_3_0)($i_prime__2))(($dictBounded_0)->{'bottom'});
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Enum_defaultToEnum'] = __NAMESPACE__ . '\\majData_majEnum_defaultmajTomajEnum';

// Data_Enum_defaultSucc
function majData_majEnum_defaultmajSucc($toEnum_prime__0, $fromEnum_prime__1 = null, $a_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEnum_defaultmajSucc';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($toEnum_prime__0)((($fromEnum_prime__1)($a_2) + 1));
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Enum_defaultSucc'] = __NAMESPACE__ . '\\majData_majEnum_defaultmajSucc';

// Data_Enum_defaultPred
function majData_majEnum_defaultmajPred($toEnum_prime__0, $fromEnum_prime__1 = null, $a_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEnum_defaultmajPred';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($toEnum_prime__0)((($fromEnum_prime__1)($a_2) - 1));
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Enum_defaultPred'] = __NAMESPACE__ . '\\majData_majEnum_defaultmajPred';

// Data_Enum_defaultFromEnum
function majData_majEnum_defaultmajFrommajEnum($dictEnum_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEnum_defaultmajFrommajEnum';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $go__go_1_0 = null;
  $go__go_1_0 = (function() use ($dictEnum_0, &$go__go_1_0) {
  $__fn = function(int $i_2, $x_3 = null) use ($dictEnum_0, &$go__go_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_1_0_0_i_2 = $i_2;
  $__tco_var_go__go_1_0_0_x_3 = $x_3;
  tco_loop_go__go_1_0_0:;
  $i_2 = $__tco_var_go__go_1_0_0_i_2;
  $x_3 = $__tco_var_go__go_1_0_0_x_3;
  $v_4_0 = (($dictEnum_0)->{'pred'})($x_3);
  $__t1 = null;;
  if ($v_4_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__tco_2 = ($i_2 + 1);
$__tco_3 = ($v_4_0)->{'value0'};
$__tco_var_go__go_1_0_0_i_2 = $__tco_2;
$__tco_var_go__go_1_0_0_x_3 = $__tco_3;
goto tco_loop_go__go_1_0_0;;
$__t1 = null;
goto end_branch_1;;
};
  if ($v_4_0 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t1 = $i_2;
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
})();
  $__res = ($go__go_1_0)(0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Enum_defaultFromEnum'] = __NAMESPACE__ . '\\majData_majEnum_defaultmajFrommajEnum';

// Data_Enum_defaultCardinality
function majData_majEnum_defaultmajCardinality($dictBounded_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEnum_defaultmajCardinality';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $bottom2_1_0 = ($dictBounded_0)->{'bottom'};
  $__res = function($dictEnum_2) use ($bottom2_1_0) {
  $__num = \func_num_args();
  $go__go_3_1 = null;
  $go__go_3_1 = (function() use ($dictEnum_2, &$go__go_3_1) {
  $__fn = function(int $i_4, $x_5 = null) use ($dictEnum_2, &$go__go_3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_3_1_1_i_4 = $i_4;
  $__tco_var_go__go_3_1_1_x_5 = $x_5;
  tco_loop_go__go_3_1_1:;
  $i_4 = $__tco_var_go__go_3_1_1_i_4;
  $x_5 = $__tco_var_go__go_3_1_1_x_5;
  $v_6_1 = (($dictEnum_2)->{'succ'})($x_5);
  $__t2 = null;;
  if ($v_6_1 instanceof \Data\Maybe\Data_Maybe_Just) {
$__tco_3 = ($i_4 + 1);
$__tco_4 = ($v_6_1)->{'value0'};
$__tco_var_go__go_3_1_1_i_4 = $__tco_3;
$__tco_var_go__go_3_1_1_x_5 = $__tco_4;
goto tco_loop_go__go_3_1_1;;
$__t2 = null;
goto end_branch_2;;
};
  if ($v_6_1 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t2 = $i_4;
goto end_branch_2;;
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
})();
  $__res = (($go__go_3_1)(1))($bottom2_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Enum_defaultCardinality'] = __NAMESPACE__ . '\\majData_majEnum_defaultmajCardinality';

// Data_Enum_charToEnum
function majData_majEnum_charmajTomajEnum(int $v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEnum_charmajTomajEnum';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__t0 = null;;
  if ((($v_0 >= \Data\Enum\majData_majEnum_tomajCharmajCode($GLOBALS['Data_Bounded_bottomChar'])) && ($v_0 <= \Data\Enum\majData_majEnum_tomajCharmajCode($GLOBALS['Data_Bounded_topChar'])))) {
$__t0 = new \Data\Maybe\Data_Maybe_Just(\Data\Enum\majData_majEnum_frommajCharmajCode($v_0));
goto end_branch_0;;
};
  $__t0 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Enum_charToEnum'] = __NAMESPACE__ . '\\majData_majEnum_charmajTomajEnum';

// Data_Enum_enumChar
$GLOBALS['Data_Enum_enumChar'] = (object)["succ" => function($a_0) {
  $__num = \func_num_args();
  $__res = \Data\Enum\majData_majEnum_charmajTomajEnum((\Data\Enum\majData_majEnum_tomajCharmajCode($a_0) + 1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "pred" => function($a_0) {
  $__num = \func_num_args();
  $__res = \Data\Enum\majData_majEnum_charmajTomajEnum((\Data\Enum\majData_majEnum_tomajCharmajCode($a_0) - 1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Ord0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Ord_ordChar'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Enum_cardinality
function majData_majEnum_cardinality($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEnum_cardinality';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'cardinality'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Enum_cardinality'] = __NAMESPACE__ . '\\majData_majEnum_cardinality';

// Data_Enum_boundedEnumUnit
$GLOBALS['Data_Enum_boundedEnumUnit'] = (object)["cardinality" => 1, "toEnum" => function($v_0) {
  $__num = \func_num_args();
  $__res = match ($v_0) { 0 => new \Data\Maybe\Data_Maybe_Just($GLOBALS['Data_Unit_unit']), default => new \Data\Maybe\Data_Maybe_Nothing() };
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "fromEnum" => function($v_0) {
  $__num = \func_num_args();
  $__res = 0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bounded0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Bounded_boundedUnit'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Enum1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Enum_enumUnit'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Enum_boundedEnumOrdering
$GLOBALS['Data_Enum_boundedEnumOrdering'] = (object)["cardinality" => 3, "toEnum" => function($v_0) {
  $__num = \func_num_args();
  $__res = match ($v_0) { 0 => new \Data\Maybe\Data_Maybe_Just(new \Data\Ordering\Data_Ordering_LT()), 1 => new \Data\Maybe\Data_Maybe_Just(new \Data\Ordering\Data_Ordering_EQ()), 2 => new \Data\Maybe\Data_Maybe_Just(new \Data\Ordering\Data_Ordering_GT()), default => new \Data\Maybe\Data_Maybe_Nothing() };
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "fromEnum" => function($v_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v_0 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t1 = 0;
goto end_branch_1;;
};
  if ($v_0 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t1 = 1;
goto end_branch_1;;
};
  if ($v_0 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t1 = 2;
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bounded0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Bounded_boundedOrdering'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Enum1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Enum_enumOrdering'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Enum_boundedEnumChar
$GLOBALS['Data_Enum_boundedEnumChar'] = (object)["cardinality" => (\Data\Enum\majData_majEnum_tomajCharmajCode($GLOBALS['Data_Bounded_topChar']) - \Data\Enum\majData_majEnum_tomajCharmajCode($GLOBALS['Data_Bounded_bottomChar'])), "toEnum" => $GLOBALS['Data_Enum_charToEnum'], "fromEnum" => $GLOBALS['Data_Enum_toCharCode'], "Bounded0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Bounded_boundedChar'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Enum1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Enum_enumChar'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Enum_boundedEnumBoolean
$GLOBALS['Data_Enum_boundedEnumBoolean'] = (object)["cardinality" => 2, "toEnum" => function($v_0) {
  $__num = \func_num_args();
  $__res = match ($v_0) { 0 => new \Data\Maybe\Data_Maybe_Just(false), 1 => new \Data\Maybe\Data_Maybe_Just(true), default => new \Data\Maybe\Data_Maybe_Nothing() };
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "fromEnum" => function($v_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if (( ! $v_0)) {
$__t1 = 0;
goto end_branch_1;;
};
  if ($v_0) {
$__t1 = 1;
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bounded0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Bounded_boundedBoolean'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Enum1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Enum_enumBoolean'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

