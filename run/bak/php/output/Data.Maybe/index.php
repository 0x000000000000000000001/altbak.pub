<?php

namespace Data\Maybe;

// ALL IMPORTS: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Extend, Control.Monad, Control.Plus, Data.Bounded, Data.Eq, Data.Function, Data.Functor, Data.Functor.Invariant, Data.Generic.Rep, Data.Maybe, Data.Monoid, Data.Ord, Data.Ordering, Data.Semigroup, Data.Semiring, Data.Show, Data.Unit, Prelude, Prim
// TO REQUIRE: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Extend, Control.Monad, Control.Plus, Data.Bounded, Data.Eq, Data.Function, Data.Functor, Data.Functor.Invariant, Data.Generic.Rep, Data.Maybe, Data.Monoid, Data.Ord, Data.Ordering, Data.Semigroup, Data.Semiring, Data.Show, Data.Unit, Prelude
require_once __DIR__ . '/../Control.Alt/index.php';
require_once __DIR__ . '/../Control.Alternative/index.php';
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Category/index.php';
require_once __DIR__ . '/../Control.Extend/index.php';
require_once __DIR__ . '/../Control.Monad/index.php';
require_once __DIR__ . '/../Control.Plus/index.php';
require_once __DIR__ . '/../Data.Bounded/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Functor.Invariant/index.php';
require_once __DIR__ . '/../Data.Generic.Rep/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Monoid/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ordering/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Semiring/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
require_once __DIR__ . '/../Data.Unit/index.php';
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


final class Data_Maybe_Nothing { public $tag = 'Nothing'; public function __construct() {} }
final class Data_Maybe_Just { public $tag = 'Just'; public function __construct(public  $value0) {} }

// Data_Maybe_identity
function majData_majMaybe_identity($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMaybe_identity';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Maybe_identity'] = __NAMESPACE__ . '\\majData_majMaybe_identity';

// Data_Maybe_Nothing
$GLOBALS['Data_Maybe_Nothing'] = ($GLOBALS['__phpurs_data0_Nothing'] ??= new \Data\Maybe\Data_Maybe_Nothing());

// Data_Maybe_Just
$GLOBALS['Data_Maybe_Just'] = function($value0) {
  $__num = \func_num_args();
  $__res = new \Data\Maybe\Data_Maybe_Just($value0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Maybe_showMaybe
function majData_majMaybe_showmajMaybe($dictShow_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMaybe_showmajMaybe';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["show" => function($v_1) use ($dictShow_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v_1 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t0 = ((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})("(Just "))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})((($dictShow_0)->{'show'})(($v_1)->{'value0'})))(")"));
goto end_branch_0;;
};
  if ($v_1 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t0 = "Nothing";
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Maybe_showMaybe'] = __NAMESPACE__ . '\\majData_majMaybe_showmajMaybe';

// Data_Maybe_semigroupMaybe
function majData_majMaybe_semigroupmajMaybe($dictSemigroup_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMaybe_semigroupmajMaybe';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["append" => (function() use ($dictSemigroup_0) {
  $__fn = function($v_1, $v1_2 = null) use ($dictSemigroup_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ($v_1 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t0 = $v1_2;
goto end_branch_0;;
};
  if ($v1_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t0 = $v_1;
goto end_branch_0;;
};
  if (($v_1 instanceof \Data\Maybe\Data_Maybe_Just && $v1_2 instanceof \Data\Maybe\Data_Maybe_Just)) {
$__t0 = new \Data\Maybe\Data_Maybe_Just(((($dictSemigroup_0)->{'append'})(($v_1)->{'value0'}))(($v1_2)->{'value0'}));
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Maybe_semigroupMaybe'] = __NAMESPACE__ . '\\majData_majMaybe_semigroupmajMaybe';

// Data_Maybe_optional
function majData_majMaybe_optional($dictAlt_0, $dictApplicative_1 = null, $a_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMaybe_optional';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictAlt_0)->{'alt'})(((((($dictAlt_0)->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Maybe_Just']))($a_2)))((($dictApplicative_1)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing()));
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Maybe_optional'] = __NAMESPACE__ . '\\majData_majMaybe_optional';

// Data_Maybe_monoidMaybe
function majData_majMaybe_monoidmajMaybe($dictSemigroup_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMaybe_monoidmajMaybe';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $semigroupMaybe1_1_0 = (object)["append" => (function() use ($dictSemigroup_0) {
  $__fn = function($v_1, $v1_2 = null) use ($dictSemigroup_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ($v_1 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t0 = $v1_2;
goto end_branch_0;;
};
  if ($v1_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t0 = $v_1;
goto end_branch_0;;
};
  if (($v_1 instanceof \Data\Maybe\Data_Maybe_Just && $v1_2 instanceof \Data\Maybe\Data_Maybe_Just)) {
$__t0 = new \Data\Maybe\Data_Maybe_Just(((($dictSemigroup_0)->{'append'})(($v_1)->{'value0'}))(($v1_2)->{'value0'}));
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $__res = (object)["mempty" => new \Data\Maybe\Data_Maybe_Nothing(), "Semigroup0" => function($_dollar__unused_2) use ($semigroupMaybe1_1_0) {
  $__num = \func_num_args();
  $__res = $semigroupMaybe1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Maybe_monoidMaybe'] = __NAMESPACE__ . '\\majData_majMaybe_monoidmajMaybe';

// Data_Maybe_maybe'
function majData_majMaybe_maybe__prime__($v_0, $v1_1 = null, $v2_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMaybe_maybe__prime__';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if ($v2_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t0 = ($v_0)($GLOBALS['Data_Unit_unit']);
goto end_branch_0;;
};
  if ($v2_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t0 = ($v1_1)(($v2_2)->{'value0'});
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Maybe_maybe__prime__'] = __NAMESPACE__ . '\\majData_majMaybe_maybe__prime__';

// Data_Maybe_maybe
function majData_majMaybe_maybe($v_0, $v1_1 = null, $v2_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMaybe_maybe';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if ($v2_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t0 = $v_0;
goto end_branch_0;;
};
  if ($v2_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t0 = ($v1_1)(($v2_2)->{'value0'});
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Maybe_maybe'] = __NAMESPACE__ . '\\majData_majMaybe_maybe';

// Data_Maybe_isNothing
function majData_majMaybe_ismajNothing($v2_0): bool|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMaybe_ismajNothing';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__t0 = null;;
  if ($v2_0 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t0 = true;
goto end_branch_0;;
};
  if ($v2_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t0 = false;
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
$GLOBALS['Data_Maybe_isNothing'] = __NAMESPACE__ . '\\majData_majMaybe_ismajNothing';

// Data_Maybe_isJust
function majData_majMaybe_ismajJust($v2_0): bool|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMaybe_ismajJust';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__t0 = null;;
  if ($v2_0 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t0 = false;
goto end_branch_0;;
};
  if ($v2_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t0 = true;
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
$GLOBALS['Data_Maybe_isJust'] = __NAMESPACE__ . '\\majData_majMaybe_ismajJust';

// Data_Maybe_genericMaybe
$GLOBALS['Data_Maybe_genericMaybe'] = (object)["to" => function($x_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($x_0 instanceof \Data\Generic\Rep\Data_Generic_Rep_Inl) {
$__t0 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_0;;
};
  if ($x_0 instanceof \Data\Generic\Rep\Data_Generic_Rep_Inr) {
$__t0 = new \Data\Maybe\Data_Maybe_Just(($x_0)->{'value0'});
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "from" => function($x_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($x_0 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t1 = new \Data\Generic\Rep\Data_Generic_Rep_Inl(new \Data\Generic\Rep\Data_Generic_Rep_NoArguments());
goto end_branch_1;;
};
  if ($x_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = new \Data\Generic\Rep\Data_Generic_Rep_Inr(($x_0)->{'value0'});
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Maybe_functorMaybe
$GLOBALS['Data_Maybe_functorMaybe'] = (object)["map" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ($v1_1 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t0 = new \Data\Maybe\Data_Maybe_Just(($v_0)(($v1_1)->{'value0'}));
goto end_branch_0;;
};
  $__t0 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];

// Data_Maybe_invariantMaybe
$GLOBALS['Data_Maybe_invariantMaybe'] = (object)["imap" => (function() {
  $__fn = function($f_0, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})($f_0);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];

// Data_Maybe_fromMaybe'
function majData_majMaybe_frommajMaybe__prime__($a_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMaybe_frommajMaybe__prime__';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Data_Maybe_maybe__prime__'])($a_0))($GLOBALS['Data_Maybe_identity']);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Maybe_fromMaybe__prime__'] = __NAMESPACE__ . '\\majData_majMaybe_frommajMaybe__prime__';

// Data_Maybe_fromMaybe
function majData_majMaybe_frommajMaybe($a_0, $v2_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMaybe_frommajMaybe';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ($v2_1 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t0 = $a_0;
goto end_branch_0;;
};
  if ($v2_1 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t0 = ($v2_1)->{'value0'};
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Maybe_fromMaybe'] = __NAMESPACE__ . '\\majData_majMaybe_frommajMaybe';

// Data_Maybe_fromJust
function majData_majMaybe_frommajJust($_dollar__unused_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMaybe_frommajJust';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ($v_1 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t0 = ($v_1)->{'value0'};
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Maybe_fromJust'] = __NAMESPACE__ . '\\majData_majMaybe_frommajJust';

// Data_Maybe_extendMaybe
$GLOBALS['Data_Maybe_extendMaybe'] = (object)["extend" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ($v1_1 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t0 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_0;;
};
  $__t0 = new \Data\Maybe\Data_Maybe_Just(($v_0)($v1_1));
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Maybe_functorMaybe'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Maybe_eqMaybe
function majData_majMaybe_eqmajMaybe($dictEq_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMaybe_eqmajMaybe';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["eq" => (function() use ($dictEq_0) {
  $__fn = function($x_1, $y_2 = null) use ($dictEq_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ($x_1 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t0 = $y_2 instanceof \Data\Maybe\Data_Maybe_Nothing;
goto end_branch_0;;
};
  $__t0 = ($x_1 instanceof \Data\Maybe\Data_Maybe_Just && ($y_2 instanceof \Data\Maybe\Data_Maybe_Just && ((($dictEq_0)->{'eq'})(($x_1)->{'value0'}))(($y_2)->{'value0'})));
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Maybe_eqMaybe'] = __NAMESPACE__ . '\\majData_majMaybe_eqmajMaybe';

// Data_Maybe_ordMaybe
function majData_majMaybe_ordmajMaybe($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMaybe_ordmajMaybe';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictOrd_0)->{'Eq0'})(null);
  $eqMaybe1_2_1 = (object)["eq" => (function() use ($__local_var_1_0) {
  $__fn = function($x_2, $y_3 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t1 = null;;
  if ($x_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t1 = $y_3 instanceof \Data\Maybe\Data_Maybe_Nothing;
goto end_branch_1;;
};
  $__t1 = ($x_2 instanceof \Data\Maybe\Data_Maybe_Just && ($y_3 instanceof \Data\Maybe\Data_Maybe_Just && ((($__local_var_1_0)->{'eq'})(($x_2)->{'value0'}))(($y_3)->{'value0'})));
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $__res = (object)["compare" => (function() use ($dictOrd_0) {
  $__fn = function($x_3, $y_4 = null) use ($dictOrd_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t3 = null;;
  if ($x_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t4 = null;;
if ($y_4 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t4 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_4;;
};
$__t4 = new \Data\Ordering\Data_Ordering_LT();
end_branch_4:;
$__t3 = $__t4;
goto end_branch_3;;
};
  if ($y_4 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t3 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_3;;
};
  if (($x_3 instanceof \Data\Maybe\Data_Maybe_Just && $y_4 instanceof \Data\Maybe\Data_Maybe_Just)) {
$__t3 = ((($dictOrd_0)->{'compare'})(($x_3)->{'value0'}))(($y_4)->{'value0'});
goto end_branch_3;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t3 = null;
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($_dollar__unused_3) use ($eqMaybe1_2_1) {
  $__num = \func_num_args();
  $__res = $eqMaybe1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Maybe_ordMaybe'] = __NAMESPACE__ . '\\majData_majMaybe_ordmajMaybe';

// Data_Maybe_eq1Maybe
$GLOBALS['Data_Maybe_eq1Maybe'] = (object)["eq1" => (function() {
  $__fn = function($dictEq_0, $x_1 = null, $y_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if ($x_1 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t0 = $y_2 instanceof \Data\Maybe\Data_Maybe_Nothing;
goto end_branch_0;;
};
  $__t0 = ($x_1 instanceof \Data\Maybe\Data_Maybe_Just && ($y_2 instanceof \Data\Maybe\Data_Maybe_Just && ((($dictEq_0)->{'eq'})(($x_1)->{'value0'}))(($y_2)->{'value0'})));
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()];

// Data_Maybe_ord1Maybe
$GLOBALS['Data_Maybe_ord1Maybe'] = (object)["compare1" => (function() {
  $__fn = function($dictOrd_0, $x_1 = null, $y_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if ($x_1 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t1 = null;;
if ($y_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t1 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_1;;
};
$__t1 = new \Data\Ordering\Data_Ordering_LT();
end_branch_1:;
$__t0 = $__t1;
goto end_branch_0;;
};
  if ($y_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t0 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_0;;
};
  if (($x_1 instanceof \Data\Maybe\Data_Maybe_Just && $y_2 instanceof \Data\Maybe\Data_Maybe_Just)) {
$__t0 = ((($dictOrd_0)->{'compare'})(($x_1)->{'value0'}))(($y_2)->{'value0'});
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "Eq10" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Maybe_eq1Maybe'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Maybe_boundedMaybe
function majData_majMaybe_boundedmajMaybe($dictBounded_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMaybe_boundedmajMaybe';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictBounded_0)->{'Ord0'})(null);
  $__local_var_2_1 = (($__local_var_1_0)->{'Eq0'})(null);
  $eqMaybe1_3_2 = (object)["eq" => (function() use ($__local_var_2_1) {
  $__fn = function($x_3, $y_4 = null) use ($__local_var_2_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t2 = null;;
  if ($x_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t2 = $y_4 instanceof \Data\Maybe\Data_Maybe_Nothing;
goto end_branch_2;;
};
  $__t2 = ($x_3 instanceof \Data\Maybe\Data_Maybe_Just && ($y_4 instanceof \Data\Maybe\Data_Maybe_Just && ((($__local_var_2_1)->{'eq'})(($x_3)->{'value0'}))(($y_4)->{'value0'})));
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $ordMaybe1_3_2 = (object)["compare" => (function() use ($__local_var_1_0) {
  $__fn = function($x_4, $y_5 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t4 = null;;
  if ($x_4 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t5 = null;;
if ($y_5 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t5 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_5;;
};
$__t5 = new \Data\Ordering\Data_Ordering_LT();
end_branch_5:;
$__t4 = $__t5;
goto end_branch_4;;
};
  if ($y_5 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t4 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_4;;
};
  if (($x_4 instanceof \Data\Maybe\Data_Maybe_Just && $y_5 instanceof \Data\Maybe\Data_Maybe_Just)) {
$__t4 = ((($__local_var_1_0)->{'compare'})(($x_4)->{'value0'}))(($y_5)->{'value0'});
goto end_branch_4;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t4 = null;
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($_dollar__unused_4) use ($eqMaybe1_3_2) {
  $__num = \func_num_args();
  $__res = $eqMaybe1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["top" => new \Data\Maybe\Data_Maybe_Just(($dictBounded_0)->{'top'}), "bottom" => new \Data\Maybe\Data_Maybe_Nothing(), "Ord0" => function($_dollar__unused_4) use ($ordMaybe1_3_2) {
  $__num = \func_num_args();
  $__res = $ordMaybe1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Maybe_boundedMaybe'] = __NAMESPACE__ . '\\majData_majMaybe_boundedmajMaybe';

// Data_Maybe_applyMaybe
$GLOBALS['Data_Maybe_applyMaybe'] = (object)["apply" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ($v_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t0 = ((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})(($v_0)->{'value0'}))($v1_1);
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t0 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Maybe_functorMaybe'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Maybe_bindMaybe
$GLOBALS['Data_Maybe_bindMaybe'] = (object)["bind" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ($v_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t0 = ($v1_1)(($v_0)->{'value0'});
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t0 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Apply0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Maybe_applyMaybe'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Maybe_semiringMaybe
function majData_majMaybe_semiringmajMaybe($dictSemiring_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMaybe_semiringmajMaybe';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $mul_1_0 = ($dictSemiring_0)->{'mul'};
  $__res = (object)["zero" => new \Data\Maybe\Data_Maybe_Nothing(), "one" => new \Data\Maybe\Data_Maybe_Just(($dictSemiring_0)->{'one'}), "add" => (function() use ($dictSemiring_0) {
  $__fn = function($v_2, $v1_3 = null) use ($dictSemiring_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t1 = null;;
  if ($v_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t1 = $v1_3;
goto end_branch_1;;
};
  if ($v1_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t1 = $v_2;
goto end_branch_1;;
};
  if (($v_2 instanceof \Data\Maybe\Data_Maybe_Just && $v1_3 instanceof \Data\Maybe\Data_Maybe_Just)) {
$__t1 = new \Data\Maybe\Data_Maybe_Just(((($dictSemiring_0)->{'add'})(($v_2)->{'value0'}))(($v1_3)->{'value0'}));
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
})(), "mul" => (function() use ($mul_1_0) {
  $__fn = function($x_2, $y_3 = null) use ($mul_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_Maybe_applyMaybe'])->{'apply'})(((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})($mul_1_0))($x_2)))($y_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Maybe_semiringMaybe'] = __NAMESPACE__ . '\\majData_majMaybe_semiringmajMaybe';

// Data_Maybe_applicativeMaybe
$GLOBALS['Data_Maybe_applicativeMaybe'] = (object)["pure" => $GLOBALS['Data_Maybe_Just'], "Apply0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Maybe_applyMaybe'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Maybe_monadMaybe
$GLOBALS['Data_Maybe_monadMaybe'] = (object)["Applicative0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Maybe_applicativeMaybe'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Maybe_bindMaybe'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Maybe_altMaybe
$GLOBALS['Data_Maybe_altMaybe'] = (object)["alt" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ($v_0 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t0 = $v1_1;
goto end_branch_0;;
};
  $__t0 = $v_0;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Maybe_functorMaybe'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Maybe_plusMaybe
$GLOBALS['Data_Maybe_plusMaybe'] = (object)["empty" => new \Data\Maybe\Data_Maybe_Nothing(), "Alt0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Maybe_altMaybe'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Maybe_alternativeMaybe
$GLOBALS['Data_Maybe_alternativeMaybe'] = (object)["Applicative0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Maybe_applicativeMaybe'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Plus1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Maybe_plusMaybe'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

