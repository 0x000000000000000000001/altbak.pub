<?php

namespace Data\Monoid;

// ALL IMPORTS: Data.Boolean, Data.Eq, Data.EuclideanRing, Data.Monoid, Data.Ord, Data.Ordering, Data.Semigroup, Data.Symbol, Data.Unit, Prim, Prim.Row, Prim.RowList, Record.Unsafe, Type.Proxy
// TO REQUIRE: Data.Boolean, Data.Eq, Data.EuclideanRing, Data.Monoid, Data.Ord, Data.Ordering, Data.Semigroup, Data.Symbol, Data.Unit, Record.Unsafe, Type.Proxy
require_once __DIR__ . '/../Data.Boolean/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.EuclideanRing/index.php';
require_once __DIR__ . '/../Data.Monoid/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ordering/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Symbol/index.php';
require_once __DIR__ . '/../Data.Unit/index.php';
require_once __DIR__ . '/../Record.Unsafe/index.php';
require_once __DIR__ . '/../Type.Proxy/index.php';

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




// Data_Monoid_lessThanOrEq
$GLOBALS['Data_Monoid_lessThanOrEq'] = (function() use (&$__fn) {
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

// Data_Monoid_monoidUnit
$GLOBALS['Data_Monoid_monoidUnit'] = (object)["mempty" => $GLOBALS['Data_Unit_unit'], "Semigroup0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Semigroup_semigroupUnit'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Monoid_monoidString
$GLOBALS['Data_Monoid_monoidString'] = (object)["mempty" => "", "Semigroup0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Semigroup_semigroupString'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Monoid_monoidRecordNil
$GLOBALS['Data_Monoid_monoidRecordNil'] = (object)["memptyRecord" => function($v_0) {
  $__num = \func_num_args();
  $__res = (object)[];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "SemigroupRecord0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Semigroup_semigroupRecordNil'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Monoid_monoidOrdering
$GLOBALS['Data_Monoid_monoidOrdering'] = (object)["mempty" => new \Data\Ordering\Data_Ordering_EQ(), "Semigroup0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Ordering_semigroupOrdering'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Monoid_monoidArray
$GLOBALS['Data_Monoid_monoidArray'] = (object)["mempty" => [], "Semigroup0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Semigroup_semigroupArray'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Monoid_memptyRecord
function majData_majMonoid_memptymajRecord($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMonoid_memptymajRecord';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'memptyRecord'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Monoid_memptyRecord'] = __NAMESPACE__ . '\\majData_majMonoid_memptymajRecord';

// Data_Monoid_monoidRecord
function majData_majMonoid_monoidmajRecord($_dollar__unused_0, $dictMonoidRecord_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMonoid_monoidmajRecord';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $semigroupRecord1_2_0 = (object)["append" => (((($dictMonoidRecord_1)->{'SemigroupRecord0'})(null))->{'appendRecord'})(new \Type\Proxy\Type_Proxy_Proxy())];
  $__res = (object)["mempty" => (($dictMonoidRecord_1)->{'memptyRecord'})(new \Type\Proxy\Type_Proxy_Proxy()), "Semigroup0" => function($_dollar__unused_3) use ($semigroupRecord1_2_0) {
  $__num = \func_num_args();
  $__res = $semigroupRecord1_2_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Monoid_monoidRecord'] = __NAMESPACE__ . '\\majData_majMonoid_monoidmajRecord';

// Data_Monoid_mempty
function majData_majMonoid_mempty($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMonoid_mempty';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'mempty'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Monoid_mempty'] = __NAMESPACE__ . '\\majData_majMonoid_mempty';

// Data_Monoid_monoidFn
function majData_majMonoid_monoidmajFn($dictMonoid_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMonoid_monoidmajFn';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $mempty1_1_0 = ($dictMonoid_0)->{'mempty'};
  $__local_var_2_1 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $semigroupFn_3_2 = (object)["append" => (function() use ($__local_var_2_1) {
  $__fn = function($f_3, $g_4 = null, $x_5 = null) use ($__local_var_2_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($__local_var_2_1)->{'append'})(($f_3)($x_5)))(($g_4)($x_5));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()];
  $__res = (object)["mempty" => function($v_4) use ($mempty1_1_0) {
  $__num = \func_num_args();
  $__res = $mempty1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Semigroup0" => function($_dollar__unused_4) use ($semigroupFn_3_2) {
  $__num = \func_num_args();
  $__res = $semigroupFn_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Monoid_monoidFn'] = __NAMESPACE__ . '\\majData_majMonoid_monoidmajFn';

// Data_Monoid_monoidRecordCons
function majData_majMonoid_monoidmajRecordmajCons($dictIsSymbol_0, $dictMonoid_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMonoid_monoidmajRecordmajCons';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $mempty1_2_0 = ($dictMonoid_1)->{'mempty'};
  $Semigroup0_3_1 = (($dictMonoid_1)->{'Semigroup0'})(null);
  $__res = (function() use ($Semigroup0_3_1, $dictIsSymbol_0, $mempty1_2_0) {
  $__fn = function($_dollar__unused_4, $dictMonoidRecord_5 = null) use ($Semigroup0_3_1, $dictIsSymbol_0, $mempty1_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_6_2 = (($dictMonoidRecord_5)->{'SemigroupRecord0'})(null);
  $semigroupRecordCons1_7_3 = (object)["appendRecord" => (function() use ($Semigroup0_3_1, $__local_var_6_2, $dictIsSymbol_0) {
  $__fn = function($v_7, $ra_8 = null, $rb_9 = null) use ($Semigroup0_3_1, $__local_var_6_2, $dictIsSymbol_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $key_10_3 = (($dictIsSymbol_0)->{'reflectSymbol'})(new \Type\Proxy\Type_Proxy_Proxy());
  $get_11_4 = ($GLOBALS['Record_Unsafe_unsafeGet'])($key_10_3);
  $__res = \Record\Unsafe\majRecord_majUnsafe_unsafemajSet($key_10_3, ((($Semigroup0_3_1)->{'append'})(($get_11_4)($ra_8)))(($get_11_4)($rb_9)), (((($__local_var_6_2)->{'appendRecord'})(new \Type\Proxy\Type_Proxy_Proxy()))($ra_8))($rb_9));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()];
  $__res = (object)["memptyRecord" => function($v_8) use ($dictIsSymbol_0, $dictMonoidRecord_5, $mempty1_2_0) {
  $__num = \func_num_args();
  $__res = \Record\Unsafe\majRecord_majUnsafe_unsafemajSet((($dictIsSymbol_0)->{'reflectSymbol'})(new \Type\Proxy\Type_Proxy_Proxy()), $mempty1_2_0, (($dictMonoidRecord_5)->{'memptyRecord'})(new \Type\Proxy\Type_Proxy_Proxy()));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "SemigroupRecord0" => function($_dollar__unused_8) use ($semigroupRecordCons1_7_3) {
  $__num = \func_num_args();
  $__res = $semigroupRecordCons1_7_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Monoid_monoidRecordCons'] = __NAMESPACE__ . '\\majData_majMonoid_monoidmajRecordmajCons';

// Data_Monoid_power
function majData_majMonoid_power($dictMonoid_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMonoid_power';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $mempty1_1_0 = ($dictMonoid_0)->{'mempty'};
  $__local_var_2_1 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $__res = function($x_3) use ($__local_var_2_1, $mempty1_1_0) {
  $__num = \func_num_args();
  $go__4_2 = null;
  $go__4_2 = function($p_5) use ($__local_var_2_1, &$go__4_2, $mempty1_1_0, $x_3) {
  $__num = \func_num_args();
  $__t4 = null;;
  if ((($GLOBALS['Data_Monoid_lessThanOrEq'])($p_5))(0)) {
$__t4 = $mempty1_1_0;
goto end_branch_4;;
};
  switch ($p_5) {
case 1:
$__t4 = $x_3;
goto end_branch_4;;
break;
default:
;
break;
};
  switch (((($GLOBALS['Data_EuclideanRing_euclideanRingInt'])->{'mod'})($p_5))(2)) {
case 0:
$x_prime_6_5 = ($go__4_2)(((($GLOBALS['Data_EuclideanRing_euclideanRingInt'])->{'div'})($p_5))(2));
$__t4 = ((($__local_var_2_1)->{'append'})($x_prime_6_5))($x_prime_6_5);
goto end_branch_4;;
break;
default:
;
break;
};
  $x_prime_6_3 = ($go__4_2)(((($GLOBALS['Data_EuclideanRing_euclideanRingInt'])->{'div'})($p_5))(2));
  $__t4 = ((($__local_var_2_1)->{'append'})($x_prime_6_3))(((($__local_var_2_1)->{'append'})($x_prime_6_3))($x_3));
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = $go__4_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Monoid_power'] = __NAMESPACE__ . '\\majData_majMonoid_power';

// Data_Monoid_guard
function majData_majMonoid_guard($dictMonoid_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMonoid_guard';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $mempty1_1_0 = ($dictMonoid_0)->{'mempty'};
  $__res = (function() use ($mempty1_1_0) {
  $__fn = function($v_2, $v1_3 = null) use ($mempty1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t1 = null;;
  if ($v_2) {
$__t1 = $v1_3;
goto end_branch_1;;
};
  $__t1 = $mempty1_1_0;
  end_branch_1:;
  $__res = $__t1;
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
$GLOBALS['Data_Monoid_guard'] = __NAMESPACE__ . '\\majData_majMonoid_guard';

