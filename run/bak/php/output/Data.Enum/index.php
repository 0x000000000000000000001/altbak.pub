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
  $f = ($ffi_Data_Enum['fromCharCode'] ?? new class { public function __invoke(...$args) { return $this; } });
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
  $f = ($ffi_Data_Enum['toCharCode'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Data_Enum_toCharCode'] = __NAMESPACE__ . '\\majData_majEnum_tomajCharmajCode';





// Data_Enum_lessThan
$GLOBALS['Data_Enum_lessThan'] = (function() use (&$__fn) {
$__local_var_0_0 = ((($GLOBALS['Data_Ord_ordIntImpl'])(new \Data\Ordering\Data_Ordering_LT()))(new \Data\Ordering\Data_Ordering_EQ()))(new \Data\Ordering\Data_Ordering_GT());
return (function() use ($__local_var_0_0) {
  $__fn = function($a1_1, $a2_2 = null) use ($__local_var_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($__local_var_0_0)($a1_1))($a2_2) instanceof \Data\Ordering\Data_Ordering_LT;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
})();

// Data_Enum_greaterThan
$GLOBALS['Data_Enum_greaterThan'] = (function() use (&$__fn) {
$__local_var_0_0 = ((($GLOBALS['Data_Ord_ordIntImpl'])(new \Data\Ordering\Data_Ordering_LT()))(new \Data\Ordering\Data_Ordering_EQ()))(new \Data\Ordering\Data_Ordering_GT());
return (function() use ($__local_var_0_0) {
  $__fn = function($a1_1, $a2_2 = null) use ($__local_var_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($__local_var_0_0)($a1_1))($a2_2) instanceof \Data\Ordering\Data_Ordering_GT;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
})();

// Data_Enum_guard
$GLOBALS['Data_Enum_guard'] = ($GLOBALS['Control_Alternative_guard'])($GLOBALS['Data_Maybe_alternativeMaybe']);

// Data_Enum_lessThanOrEq
$GLOBALS['Data_Enum_lessThanOrEq'] = (function() use (&$__fn) {
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

// Data_Enum_fromJust
function majData_majEnum_frommajJust($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEnum_frommajJust';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__t0 = null;;
  if ($v_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t0 = ($v_0)->{'value0'};
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
$GLOBALS['Data_Enum_fromJust'] = __NAMESPACE__ . '\\majData_majEnum_frommajJust';

// Data_Enum_greaterThanOrEq
$GLOBALS['Data_Enum_greaterThanOrEq'] = (function() use (&$__fn) {
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
  $__res = ((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})("(Cardinality "))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})((($GLOBALS['Data_Show_showInt'])->{'show'})($v_0)))(")"));
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
$GLOBALS['Data_Enum_ordCardinality'] = (object)["compare" => ((($GLOBALS['Data_Ord_ordIntImpl'])(new \Data\Ordering\Data_Ordering_LT()))(new \Data\Ordering\Data_Ordering_EQ()))(new \Data\Ordering\Data_Ordering_GT()), "Eq0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = (object)["eq" => $GLOBALS['Data_Eq_eqIntImpl']];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Enum_newtypeCardinality
$GLOBALS['Data_Enum_newtypeCardinality'] = (object)["Coercible0" => function($_dollar__unused_0) {
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
  $__res = (function() use ($bottom2_1_0, $dictBoundedEnum_0) {
  $__fn = function($low_2, $high_3 = null, $x_4 = null) use ($bottom2_1_0, $dictBoundedEnum_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $v_5_1 = (($dictBoundedEnum_0)->{'toEnum'})($x_4);
  $__t2 = null;;
  if ($v_5_1 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t2 = ($v_5_1)->{'value0'};
goto end_branch_2;;
};
  if ($v_5_1 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t3 = null;;
if ((($GLOBALS['Data_Enum_lessThan'])($x_4))((($dictBoundedEnum_0)->{'fromEnum'})($bottom2_1_0))) {
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
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Enum_toEnumWithDefaults'] = __NAMESPACE__ . '\\majData_majEnum_tomajEnummajWithmajDefaults';

// Data_Enum_eqCardinality
$GLOBALS['Data_Enum_eqCardinality'] = (object)["eq" => $GLOBALS['Data_Eq_eqIntImpl']];

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
}, "Ord0" => function($_dollar__unused_0) {
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
  $ordTuple_1_0 = ($GLOBALS['Data_Tuple_ordTuple'])((($dictEnum_0)->{'Ord0'})(null));
  $__res = function($dictBoundedEnum_2) use ($dictEnum_0, $ordTuple_1_0) {
  $__num = \func_num_args();
  $Bounded0_3_1 = (($dictBoundedEnum_2)->{'Bounded0'})(null);
  $bottom2_4_2 = ($Bounded0_3_1)->{'bottom'};
  $Enum1_5_3 = (($dictBoundedEnum_2)->{'Enum1'})(null);
  $top2_6_4 = ($Bounded0_3_1)->{'top'};
  $ordTuple1_7_5 = ($ordTuple_1_0)((($Enum1_5_3)->{'Ord0'})(null));
  $__res = (object)["succ" => function($v_8) use ($Enum1_5_3, $bottom2_4_2, $dictEnum_0) {
  $__num = \func_num_args();
  $__local_var_9_6 = ((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})(function($a_9) use ($bottom2_4_2) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_9, $bottom2_4_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($dictEnum_0)->{'succ'})(($v_8)->{'value0'}));
  $__local_var_10_7 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Maybe_Just']))(($GLOBALS['Data_Tuple_Tuple'])(($v_8)->{'value0'}));
  $__local_var_11_8 = (($Enum1_5_3)->{'succ'})(($v_8)->{'value1'});
  $__t9 = null;;
  if ($__local_var_11_8 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t9 = $__local_var_9_6;
goto end_branch_9;;
};
  if ($__local_var_11_8 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t9 = ($__local_var_10_7)(($__local_var_11_8)->{'value0'});
goto end_branch_9;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t9 = null;
  end_branch_9:;
  $__res = $__t9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "pred" => function($v_8) use ($Enum1_5_3, $dictEnum_0, $top2_6_4) {
  $__num = \func_num_args();
  $__local_var_9_10 = ((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})(function($a_9) use ($top2_6_4) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_9, $top2_6_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($dictEnum_0)->{'pred'})(($v_8)->{'value0'}));
  $__local_var_10_11 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Maybe_Just']))(($GLOBALS['Data_Tuple_Tuple'])(($v_8)->{'value0'}));
  $__local_var_11_12 = (($Enum1_5_3)->{'pred'})(($v_8)->{'value1'});
  $__t13 = null;;
  if ($__local_var_11_12 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t13 = $__local_var_9_10;
goto end_branch_13;;
};
  if ($__local_var_11_12 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t13 = ($__local_var_10_11)(($__local_var_11_12)->{'value0'});
goto end_branch_13;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t13 = null;
  end_branch_13:;
  $__res = $__t13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Ord0" => function($_dollar__unused_8) use ($ordTuple1_7_5) {
  $__num = \func_num_args();
  $__res = $ordTuple1_7_5;
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
}, "Ord0" => function($_dollar__unused_0) {
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
  $bottom2_1_0 = ((($dictBoundedEnum_0)->{'Bounded0'})(null))->{'bottom'};
  $Enum1_2_1 = (($dictBoundedEnum_0)->{'Enum1'})(null);
  $__local_var_3_2 = (($Enum1_2_1)->{'Ord0'})(null);
  $__local_var_4_3 = (($__local_var_3_2)->{'Eq0'})(null);
  $eqMaybe1_5_4 = (object)["eq" => (function() use ($__local_var_4_3) {
  $__fn = function($x_5, $y_6 = null) use ($__local_var_4_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
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
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $ordMaybe_5_4 = (object)["compare" => (function() use ($__local_var_3_2) {
  $__fn = function($x_6, $y_7 = null) use ($__local_var_3_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t6 = null;;
  if ($x_6 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t7 = null;;
if ($y_7 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t7 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_7;;
};
$__t7 = new \Data\Ordering\Data_Ordering_LT();
end_branch_7:;
$__t6 = $__t7;
goto end_branch_6;;
};
  if ($y_7 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t6 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_6;;
};
  if (($x_6 instanceof \Data\Maybe\Data_Maybe_Just && $y_7 instanceof \Data\Maybe\Data_Maybe_Just)) {
$__t6 = ((($__local_var_3_2)->{'compare'})(($x_6)->{'value0'}))(($y_7)->{'value0'});
goto end_branch_6;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t6 = null;
  end_branch_6:;
  $__res = $__t6;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($_dollar__unused_6) use ($eqMaybe1_5_4) {
  $__num = \func_num_args();
  $__res = $eqMaybe1_5_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["succ" => function($v_6) use ($Enum1_2_1, $bottom2_1_0) {
  $__num = \func_num_args();
  $__t9 = null;;
  if ($v_6 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t9 = new \Data\Maybe\Data_Maybe_Just(new \Data\Maybe\Data_Maybe_Just($bottom2_1_0));
goto end_branch_9;;
};
  if ($v_6 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t9 = ((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})($GLOBALS['Data_Maybe_Just']))((($Enum1_2_1)->{'succ'})(($v_6)->{'value0'}));
goto end_branch_9;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t9 = null;
  end_branch_9:;
  $__res = $__t9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "pred" => function($v_6) use ($Enum1_2_1) {
  $__num = \func_num_args();
  $__t10 = null;;
  if ($v_6 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t10 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_10;;
};
  if ($v_6 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t10 = new \Data\Maybe\Data_Maybe_Just((($Enum1_2_1)->{'pred'})(($v_6)->{'value0'}));
goto end_branch_10;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t10 = null;
  end_branch_10:;
  $__res = $__t10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Ord0" => function($_dollar__unused_6) use ($ordMaybe_5_4) {
  $__num = \func_num_args();
  $__res = $ordMaybe_5_4;
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
  if ((($GLOBALS['Data_Enum_lessThan'])($n_0))(($GLOBALS['Data_Bounded_boundedInt'])->{'top'})) {
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
  if ((($GLOBALS['Data_Enum_greaterThan'])($n_0))(($GLOBALS['Data_Bounded_boundedInt'])->{'bottom'})) {
$__t1 = new \Data\Maybe\Data_Maybe_Just(($n_0 - 1));
goto end_branch_1;;
};
  $__t1 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Ord0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = (object)["compare" => ((($GLOBALS['Data_Ord_ordIntImpl'])(new \Data\Ordering\Data_Ordering_LT()))(new \Data\Ordering\Data_Ordering_EQ()))(new \Data\Ordering\Data_Ordering_GT()), "Eq0" => function($_dollar__unused_1) {
  $__num = \func_num_args();
  $__res = (object)["eq" => $GLOBALS['Data_Eq_eqIntImpl']];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
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
  $__res = (function() use ($Ord0_1_0, $dictEnum_0) {
  $__fn = function($dictUnfoldable1_2, $v_3 = null, $v1_4 = null) use ($Ord0_1_0, $dictEnum_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t1 = null;;
  if (((((($Ord0_1_0)->{'Eq0'})(null))->{'eq'})($v_3))($v1_4)) {
$__t1 = \Data\Unfoldable1\majData_majUnfoldable1_replicate1($dictUnfoldable1_2, 1, $v_3);
goto end_branch_1;;
};
  if (((($Ord0_1_0)->{'compare'})($v_3))($v1_4) instanceof \Data\Ordering\Data_Ordering_LT) {
$__t1 = ((($dictUnfoldable1_2)->{'unfoldr1'})(function($a_5) use ($Ord0_1_0, $dictEnum_0, $v1_4) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_5, ((($GLOBALS['Data_Maybe_bindMaybe'])->{'bind'})((($dictEnum_0)->{'succ'})($a_5)))(function($a_prime_6) use ($Ord0_1_0, $v1_4) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})(function($v_7) use ($a_prime_6) {
  $__num = \func_num_args();
  $__res = $a_prime_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($GLOBALS['Data_Enum_guard'])(( ! ((($Ord0_1_0)->{'compare'})($a_prime_6))($v1_4) instanceof \Data\Ordering\Data_Ordering_GT)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_3);
goto end_branch_1;;
};
  $__t1 = ((($dictUnfoldable1_2)->{'unfoldr1'})(function($a_5) use ($Ord0_1_0, $dictEnum_0, $v1_4) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_5, ((($GLOBALS['Data_Maybe_bindMaybe'])->{'bind'})((($dictEnum_0)->{'pred'})($a_5)))(function($a_prime_6) use ($Ord0_1_0, $v1_4) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})(function($v_7) use ($a_prime_6) {
  $__num = \func_num_args();
  $__res = $a_prime_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($GLOBALS['Data_Enum_guard'])(( ! ((($Ord0_1_0)->{'compare'})($a_prime_6))($v1_4) instanceof \Data\Ordering\Data_Ordering_LT)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_3);
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
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
  $__res = (function() use ($dictBoundedEnum_2, $dictFunctor_1, $dictUnfoldable_0, $toEnum1_3_0) {
  $__fn = function($a_4, $b_5 = null, $c_6 = null) use ($dictBoundedEnum_2, $dictFunctor_1, $dictUnfoldable_0, $toEnum1_3_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $a_prime_7_1 = (($dictBoundedEnum_2)->{'fromEnum'})($a_4);
  $__local_var_8_2 = ((($dictBoundedEnum_2)->{'fromEnum'})($b_5) - $a_prime_7_1);
  $__local_var_9_3 = (($dictBoundedEnum_2)->{'fromEnum'})($c_6);
  $__res = ((($dictFunctor_1)->{'map'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Enum_fromJust']))($toEnum1_3_0)))(((($dictUnfoldable_0)->{'unfoldr'})(function($e_10) use ($__local_var_8_2, $__local_var_9_3) {
  $__num = \func_num_args();
  $__t4 = null;;
  if ((($GLOBALS['Data_Enum_lessThanOrEq'])($e_10))($__local_var_9_3)) {
$__t4 = new \Data\Maybe\Data_Maybe_Just(new \Data\Tuple\Data_Tuple_Tuple($e_10, ($e_10 + $__local_var_8_2)));
goto end_branch_4;;
};
  $__t4 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($a_prime_7_1));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
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
  $top2_2_1 = ((($dictBoundedEnum_0)->{'Bounded0'})(null))->{'top'};
  $ordEither_3_2 = ($GLOBALS['Data_Either_ordEither'])((($Enum1_1_0)->{'Ord0'})(null));
  $__res = function($dictBoundedEnum1_4) use ($Enum1_1_0, $ordEither_3_2, $top2_2_1) {
  $__num = \func_num_args();
  $bottom2_5_3 = ((($dictBoundedEnum1_4)->{'Bounded0'})(null))->{'bottom'};
  $Enum11_6_4 = (($dictBoundedEnum1_4)->{'Enum1'})(null);
  $ordEither1_7_5 = ($ordEither_3_2)((($Enum11_6_4)->{'Ord0'})(null));
  $__res = (object)["succ" => function($v_8) use ($Enum11_6_4, $Enum1_1_0, $bottom2_5_3) {
  $__num = \func_num_args();
  $__t6 = null;;
  if ($v_8 instanceof \Data\Either\Data_Either_Left) {
$__local_var_9_7 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Maybe_Just']))($GLOBALS['Data_Either_Left']);
$__local_var_10_8 = (($Enum1_1_0)->{'succ'})(($v_8)->{'value0'});
$__t9 = null;;
if ($__local_var_10_8 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t9 = new \Data\Maybe\Data_Maybe_Just(new \Data\Either\Data_Either_Right($bottom2_5_3));
goto end_branch_9;;
};
if ($__local_var_10_8 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t9 = ($__local_var_9_7)(($__local_var_10_8)->{'value0'});
goto end_branch_9;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t9 = null;
end_branch_9:;
$__t6 = $__t9;
goto end_branch_6;;
};
  if ($v_8 instanceof \Data\Either\Data_Either_Right) {
$__local_var_9_10 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Maybe_Just']))($GLOBALS['Data_Either_Right']);
$__local_var_10_11 = (($Enum11_6_4)->{'succ'})(($v_8)->{'value0'});
$__t12 = null;;
if ($__local_var_10_11 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t12 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_12;;
};
if ($__local_var_10_11 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t12 = ($__local_var_9_10)(($__local_var_10_11)->{'value0'});
goto end_branch_12;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t12 = null;
end_branch_12:;
$__t6 = $__t12;
goto end_branch_6;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t6 = null;
  end_branch_6:;
  $__res = $__t6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "pred" => function($v_8) use ($Enum11_6_4, $Enum1_1_0, $top2_2_1) {
  $__num = \func_num_args();
  $__t13 = null;;
  if ($v_8 instanceof \Data\Either\Data_Either_Left) {
$__local_var_9_14 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Maybe_Just']))($GLOBALS['Data_Either_Left']);
$__local_var_10_15 = (($Enum1_1_0)->{'pred'})(($v_8)->{'value0'});
$__t16 = null;;
if ($__local_var_10_15 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t16 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_16;;
};
if ($__local_var_10_15 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t16 = ($__local_var_9_14)(($__local_var_10_15)->{'value0'});
goto end_branch_16;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t16 = null;
end_branch_16:;
$__t13 = $__t16;
goto end_branch_13;;
};
  if ($v_8 instanceof \Data\Either\Data_Either_Right) {
$__local_var_9_17 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Maybe_Just']))($GLOBALS['Data_Either_Right']);
$__local_var_10_18 = (($Enum11_6_4)->{'pred'})(($v_8)->{'value0'});
$__t19 = null;;
if ($__local_var_10_18 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t19 = new \Data\Maybe\Data_Maybe_Just(new \Data\Either\Data_Either_Left($top2_2_1));
goto end_branch_19;;
};
if ($__local_var_10_18 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t19 = ($__local_var_9_17)(($__local_var_10_18)->{'value0'});
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
}, "Ord0" => function($_dollar__unused_8) use ($ordEither1_7_5) {
  $__num = \func_num_args();
  $__res = $ordEither1_7_5;
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
}, "Ord0" => function($_dollar__unused_0) {
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

// Data_Enum_diag
function majData_majEnum_diag($a_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEnum_diag';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_0, $a_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Enum_diag'] = __NAMESPACE__ . '\\majData_majEnum_diag';

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
  $__res = (($dictUnfoldable_2)->{'unfoldr'})((($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})($GLOBALS['Data_Enum_diag'])))($pred1_1_0));
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
  $__res = (($dictUnfoldable_2)->{'unfoldr'})((($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})($GLOBALS['Data_Enum_diag'])))($succ1_1_0));
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
function majData_majEnum_defaultmajTomajEnum($dictBounded_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEnum_defaultmajTomajEnum';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $bottom2_1_0 = ($dictBounded_0)->{'bottom'};
  $__res = (function() use ($bottom2_1_0) {
  $__fn = function($dictEnum_2, $i_prime_3 = null) use ($bottom2_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__4_1 = null;
  $go__4_1 = (function() use ($dictEnum_2, &$go__4_1) {
  $__fn = function($i_5, $x_6 = null) use ($dictEnum_2, &$go__4_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__4_1_1_i_5 = $i_5;
  $__tco_var_go__4_1_1_x_6 = $x_6;
  tco_loop_go__4_1_1:;
  $i_5 = $__tco_var_go__4_1_1_i_5;
  $x_6 = $__tco_var_go__4_1_1_x_6;
  $__t5 = null;;
  switch ($i_5) {
case 0:
$__t5 = new \Data\Maybe\Data_Maybe_Just($x_6);
goto end_branch_5;;
break;
default:
;
break;
};
  $v_7_1 = (($dictEnum_2)->{'succ'})($x_6);
  $__t2 = null;;
  if ($v_7_1 instanceof \Data\Maybe\Data_Maybe_Just) {
$__tco_3 = ($i_5 - 1);
$__tco_4 = ($v_7_1)->{'value0'};
$__tco_var_go__4_1_1_i_5 = $__tco_3;
$__tco_var_go__4_1_1_x_6 = $__tco_4;
goto tco_loop_go__4_1_1;;
$__t2 = null;
goto end_branch_2;;
};
  if ($v_7_1 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t2 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__t5 = $__t2;
  end_branch_5:;
  $__res = $__t5;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__t2 = null;;
  if ((($GLOBALS['Data_Enum_lessThan'])($i_prime_3))(0)) {
$__t2 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_2;;
};
  $__t2 = (($go__4_1)($i_prime_3))($bottom2_1_0);
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Enum_defaultToEnum'] = __NAMESPACE__ . '\\majData_majEnum_defaultmajTomajEnum';

// Data_Enum_defaultSucc
function majData_majEnum_defaultmajSucc($toEnum_prime_0, $fromEnum_prime_1 = null, $a_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEnum_defaultmajSucc';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($toEnum_prime_0)((($fromEnum_prime_1)($a_2) + 1));
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Enum_defaultSucc'] = __NAMESPACE__ . '\\majData_majEnum_defaultmajSucc';

// Data_Enum_defaultPred
function majData_majEnum_defaultmajPred($toEnum_prime_0, $fromEnum_prime_1 = null, $a_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEnum_defaultmajPred';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($toEnum_prime_0)((($fromEnum_prime_1)($a_2) - 1));
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
  $go__1_0 = null;
  $go__1_0 = (function() use ($dictEnum_0, &$go__1_0) {
  $__fn = function($i_2, $x_3 = null) use ($dictEnum_0, &$go__1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__1_0_0_i_2 = $i_2;
  $__tco_var_go__1_0_0_x_3 = $x_3;
  tco_loop_go__1_0_0:;
  $i_2 = $__tco_var_go__1_0_0_i_2;
  $x_3 = $__tco_var_go__1_0_0_x_3;
  $v_4_0 = (($dictEnum_0)->{'pred'})($x_3);
  $__t1 = null;;
  if ($v_4_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__tco_2 = ($i_2 + 1);
$__tco_3 = ($v_4_0)->{'value0'};
$__tco_var_go__1_0_0_i_2 = $__tco_2;
$__tco_var_go__1_0_0_x_3 = $__tco_3;
goto tco_loop_go__1_0_0;;
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
  $__res = ($go__1_0)(0);
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
  $go__3_1 = null;
  $go__3_1 = (function() use ($dictEnum_2, &$go__3_1) {
  $__fn = function($i_4, $x_5 = null) use ($dictEnum_2, &$go__3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__3_1_1_i_4 = $i_4;
  $__tco_var_go__3_1_1_x_5 = $x_5;
  tco_loop_go__3_1_1:;
  $i_4 = $__tco_var_go__3_1_1_i_4;
  $x_5 = $__tco_var_go__3_1_1_x_5;
  $v_6_1 = (($dictEnum_2)->{'succ'})($x_5);
  $__t2 = null;;
  if ($v_6_1 instanceof \Data\Maybe\Data_Maybe_Just) {
$__tco_3 = ($i_4 + 1);
$__tco_4 = ($v_6_1)->{'value0'};
$__tco_var_go__3_1_1_i_4 = $__tco_3;
$__tco_var_go__3_1_1_x_5 = $__tco_4;
goto tco_loop_go__3_1_1;;
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
  $__res = (($go__3_1)(1))($bottom2_1_0);
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
  if (((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'conj'})((($GLOBALS['Data_Enum_greaterThanOrEq'])($v_0))(\Data\Enum\majData_majEnum_tomajCharmajCode(($GLOBALS['Data_Bounded_boundedChar'])->{'bottom'}))))((($GLOBALS['Data_Enum_lessThanOrEq'])($v_0))(\Data\Enum\majData_majEnum_tomajCharmajCode(($GLOBALS['Data_Bounded_boundedChar'])->{'top'})))) {
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
}, "Ord0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Ord_ordChar'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Enum_cardinality
function majData_majEnum_cardinality($dict_0): int|\Closure {
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
}, "Bounded0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Bounded_boundedUnit'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Enum1" => function($_dollar__unused_0) {
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
}, "Bounded0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Bounded_boundedOrdering'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Enum1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Enum_enumOrdering'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Enum_boundedEnumChar
$GLOBALS['Data_Enum_boundedEnumChar'] = (object)["cardinality" => (\Data\Enum\majData_majEnum_tomajCharmajCode(($GLOBALS['Data_Bounded_boundedChar'])->{'top'}) - \Data\Enum\majData_majEnum_tomajCharmajCode(($GLOBALS['Data_Bounded_boundedChar'])->{'bottom'})), "toEnum" => $GLOBALS['Data_Enum_charToEnum'], "fromEnum" => $GLOBALS['Data_Enum_toCharCode'], "Bounded0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Bounded_boundedChar'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Enum1" => function($_dollar__unused_0) {
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
}, "Bounded0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Bounded_boundedBoolean'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Enum1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Enum_enumBoolean'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

