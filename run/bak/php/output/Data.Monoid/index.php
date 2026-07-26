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


// Data_Monoid_monoidUnit
$GLOBALS['Data_Monoid_monoidUnit'] = ["mempty" => $GLOBALS['Data_Unit_unit'], "Semigroup0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Semigroup_semigroupUnit'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Monoid_monoidString
$GLOBALS['Data_Monoid_monoidString'] = ["mempty" => "", "Semigroup0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Semigroup_semigroupString'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Monoid_monoidRecordNil
$GLOBALS['Data_Monoid_monoidRecordNil'] = ["memptyRecord" => function($v_0 = null) {
  $__num = \func_num_args();
  $__res = [];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "SemigroupRecord0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Semigroup_semigroupRecordNil'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Monoid_monoidOrdering
$GLOBALS['Data_Monoid_monoidOrdering'] = ["mempty" => new Phpurs_Data0("EQ"), "Semigroup0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Ordering_semigroupOrdering'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Monoid_monoidArray
$GLOBALS['Data_Monoid_monoidArray'] = ["mempty" => [], "Semigroup0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Semigroup_semigroupArray'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Monoid_memptyRecord
$GLOBALS['Data_Monoid_memptyRecord'] = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)['memptyRecord'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Monoid_monoidRecord
$GLOBALS['Data_Monoid_monoidRecord'] = (function() {
  $__fn = function($dollar__unused_0 = null, $dictMonoidRecord_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $semigroupRecord1_2_0 = ["append" => (((($dictMonoidRecord_1)['SemigroupRecord0'])($GLOBALS['Prim_undefined']))['appendRecord'])(new Phpurs_Data0("Proxy"))];
  $__res = ["mempty" => (($dictMonoidRecord_1)['memptyRecord'])(new Phpurs_Data0("Proxy")), "Semigroup0" => function($dollar__unused_3 = null) use ($semigroupRecord1_2_0) {
  $__num = \func_num_args();
  $__res = $semigroupRecord1_2_0;
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

// Data_Monoid_mempty
$GLOBALS['Data_Monoid_mempty'] = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)['mempty'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Monoid_monoidFn
$GLOBALS['Data_Monoid_monoidFn'] = function($dictMonoid_0 = null) {
  $__num = \func_num_args();
  $mempty1_1_0 = ($dictMonoid_0)['mempty'];
  $__local_var_2_1 = (($dictMonoid_0)['Semigroup0'])($GLOBALS['Prim_undefined']);
  $semigroupFn_3_2 = ["append" => (function() use ($__local_var_2_1) {
  $__fn = function($f_3 = null, $g_4 = null, $x_5 = null) use ($__local_var_2_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($__local_var_2_1)['append'])(($f_3)($x_5)))(($g_4)($x_5));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()];
  $__res = ["mempty" => function($v_4 = null) use ($mempty1_1_0) {
  $__num = \func_num_args();
  $__res = $mempty1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Semigroup0" => function($dollar__unused_4 = null) use ($semigroupFn_3_2) {
  $__num = \func_num_args();
  $__res = $semigroupFn_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Monoid_monoidRecordCons
$GLOBALS['Data_Monoid_monoidRecordCons'] = (function() {
  $__fn = function($dictIsSymbol_0 = null, $dictMonoid_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $mempty1_2_0 = ($dictMonoid_1)['mempty'];
  $Semigroup0_3_1 = (($dictMonoid_1)['Semigroup0'])($GLOBALS['Prim_undefined']);
  $__res = (function() use ($Semigroup0_3_1, $dictIsSymbol_0, $mempty1_2_0) {
  $__fn = function($dollar__unused_4 = null, $dictMonoidRecord_5 = null) use ($Semigroup0_3_1, $dictIsSymbol_0, $mempty1_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $semigroupRecordCons1_6_2 = (((($GLOBALS['Data_Semigroup_semigroupRecordCons'])($dictIsSymbol_0))($GLOBALS['Prim_undefined']))((($dictMonoidRecord_5)['SemigroupRecord0'])($GLOBALS['Prim_undefined'])))($Semigroup0_3_1);
  $__res = ["memptyRecord" => function($v_7 = null) use ($dictIsSymbol_0, $dictMonoidRecord_5, $mempty1_2_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Record_Unsafe_unsafeSet'])((($dictIsSymbol_0)['reflectSymbol'])(new Phpurs_Data0("Proxy"))))($mempty1_2_0))((($dictMonoidRecord_5)['memptyRecord'])(new Phpurs_Data0("Proxy")));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "SemigroupRecord0" => function($dollar__unused_7 = null) use ($semigroupRecordCons1_6_2) {
  $__num = \func_num_args();
  $__res = $semigroupRecordCons1_6_2;
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
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Monoid_power
$GLOBALS['Data_Monoid_power'] = function($dictMonoid_0 = null) {
  $__num = \func_num_args();
  $mempty1_1_0 = ($dictMonoid_0)['mempty'];
  $__local_var_2_1 = (($dictMonoid_0)['Semigroup0'])($GLOBALS['Prim_undefined']);
  $__res = function($x_3 = null) use ($__local_var_2_1, $mempty1_1_0) {
  $__num = \func_num_args();
  $go_4_2 = null;
  $go_4_2 = function($p_5 = null) use ($__local_var_2_1, &$go_4_2, $mempty1_1_0, $x_3) {
  $__num = \func_num_args();
  $__t4 = null;;
  if (($p_5 <= 0)) {
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
  switch ((($GLOBALS['Data_EuclideanRing_intMod'])($p_5))(2)) {
case 0:
$x__prime___6_5 = ($go_4_2)(($p_5 / 2));
$__t4 = ((($__local_var_2_1)['append'])($x__prime___6_5))($x__prime___6_5);
goto end_branch_4;;
break;
default:
;
break;
};
  $x__prime___6_3 = ($go_4_2)(($p_5 / 2));
  $__t4 = ((($__local_var_2_1)['append'])($x__prime___6_3))(((($__local_var_2_1)['append'])($x__prime___6_3))($x_3));
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = $go_4_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Monoid_guard
$GLOBALS['Data_Monoid_guard'] = function($dictMonoid_0 = null) {
  $__num = \func_num_args();
  $mempty1_1_0 = ($dictMonoid_0)['mempty'];
  $__res = (function() use ($mempty1_1_0) {
  $__fn = function($v_2 = null, $v1_3 = null) use ($mempty1_1_0, &$__fn) {
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

