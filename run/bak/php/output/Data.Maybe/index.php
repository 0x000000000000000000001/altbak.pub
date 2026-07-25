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
\PhpursThunks::$thunks['Data_Maybe_Nothing'] = function() { $v = ($GLOBALS['__phpurs_data0_Nothing'] ??= new Phpurs_Data0("Nothing")); return $v; };
\PhpursThunks::$thunks['Data_Maybe_Just'] = function() { $v = function($value0 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data1("Just", $value0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Maybe_showMaybe'] = function() { $v = function($dictShow_0 = null) {
  $__num = \func_num_args();
  $__res = (object)["show" => function($v_1 = null) use ($dictShow_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object($v_1) && (($v_1)->{'tag'} === "Just"))) {
$__t0 = ((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))("(Just "))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))((($dictShow_0)->{'show'})(($v_1)->{'value0'})))(")"));
goto end_branch_0;;
};
  if ((is_object($v_1) && (($v_1)->{'tag'} === "Nothing"))) {
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Maybe_semigroupMaybe'] = function() { $v = function($dictSemigroup_0 = null) {
  $__num = \func_num_args();
  $__res = (object)["append" => (function() use ($dictSemigroup_0) {
  $__fn = function($v_1 = null, $v1_2 = null) use ($dictSemigroup_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ((is_object($v_1) && (($v_1)->{'tag'} === "Nothing"))) {
$__t0 = $v1_2;
goto end_branch_0;;
};
  if ((is_object($v1_2) && (($v1_2)->{'tag'} === "Nothing"))) {
$__t0 = $v_1;
goto end_branch_0;;
};
  if (((is_object($v_1) && (($v_1)->{'tag'} === "Just")) && (is_object($v1_2) && (($v1_2)->{'tag'} === "Just")))) {
$__t0 = new Phpurs_Data1("Just", ((($dictSemigroup_0)->{'append'})(($v_1)->{'value0'}))(($v1_2)->{'value0'}));
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Maybe_optional'] = function() { $v = (function() {
  $__fn = function($dictAlt_0 = null, $dictApplicative_1 = null, $a_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictAlt_0)->{'alt'})(((((($dictAlt_0)->{'Functor0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->{'map'})(($GLOBALS['Data_Maybe_Just'] ?? \PhpursThunks::eval('Data_Maybe_Just'))))($a_2)))((($dictApplicative_1)->{'pure'})(new Phpurs_Data0("Nothing")));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Maybe_monoidMaybe'] = function() { $v = function($dictSemigroup_0 = null) {
  $__num = \func_num_args();
  $semigroupMaybe1_1_0 = (object)["append" => (function() use ($dictSemigroup_0) {
  $__fn = function($v_1 = null, $v1_2 = null) use ($dictSemigroup_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ((is_object($v_1) && (($v_1)->{'tag'} === "Nothing"))) {
$__t0 = $v1_2;
goto end_branch_0;;
};
  if ((is_object($v1_2) && (($v1_2)->{'tag'} === "Nothing"))) {
$__t0 = $v_1;
goto end_branch_0;;
};
  if (((is_object($v_1) && (($v_1)->{'tag'} === "Just")) && (is_object($v1_2) && (($v1_2)->{'tag'} === "Just")))) {
$__t0 = new Phpurs_Data1("Just", ((($dictSemigroup_0)->{'append'})(($v_1)->{'value0'}))(($v1_2)->{'value0'}));
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
  $__res = (object)["mempty" => new Phpurs_Data0("Nothing"), "Semigroup0" => function($dollar__unused_2 = null) use ($semigroupMaybe1_1_0) {
  $__num = \func_num_args();
  $__res = $semigroupMaybe1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Maybe_maybe__prime__'] = function() { $v = (function() {
  $__fn = function($v_0 = null, $v1_1 = null, $v2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if ((is_object($v2_2) && (($v2_2)->{'tag'} === "Nothing"))) {
$__t0 = ($v_0)(($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit')));
goto end_branch_0;;
};
  if ((is_object($v2_2) && (($v2_2)->{'tag'} === "Just"))) {
$__t0 = ($v1_1)(($v2_2)->{'value0'});
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
})(); return $v; };
\PhpursThunks::$thunks['Data_Maybe_maybe'] = function() { $v = (function() {
  $__fn = function($v_0 = null, $v1_1 = null, $v2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if ((is_object($v2_2) && (($v2_2)->{'tag'} === "Nothing"))) {
$__t0 = $v_0;
goto end_branch_0;;
};
  if ((is_object($v2_2) && (($v2_2)->{'tag'} === "Just"))) {
$__t0 = ($v1_1)(($v2_2)->{'value0'});
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
})(); return $v; };
\PhpursThunks::$thunks['Data_Maybe_isNothing'] = function() { $v = function($v2_0 = null) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object($v2_0) && (($v2_0)->{'tag'} === "Nothing"))) {
$__t0 = true;
goto end_branch_0;;
};
  if ((is_object($v2_0) && (($v2_0)->{'tag'} === "Just"))) {
$__t0 = false;
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
\PhpursThunks::$thunks['Data_Maybe_isJust'] = function() { $v = function($v2_0 = null) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object($v2_0) && (($v2_0)->{'tag'} === "Nothing"))) {
$__t0 = false;
goto end_branch_0;;
};
  if ((is_object($v2_0) && (($v2_0)->{'tag'} === "Just"))) {
$__t0 = true;
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
\PhpursThunks::$thunks['Data_Maybe_genericMaybe'] = function() { $v = (object)["to" => function($x_0 = null) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object($x_0) && (($x_0)->{'tag'} === "Inl"))) {
$__t0 = new Phpurs_Data0("Nothing");
goto end_branch_0;;
};
  if ((is_object($x_0) && (($x_0)->{'tag'} === "Inr"))) {
$__t0 = new Phpurs_Data1("Just", ($x_0)->{'value0'});
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "from" => function($x_0 = null) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((is_object($x_0) && (($x_0)->{'tag'} === "Nothing"))) {
$__t1 = new Phpurs_Data1("Inl", new Phpurs_Data0("NoArguments"));
goto end_branch_1;;
};
  if ((is_object($x_0) && (($x_0)->{'tag'} === "Just"))) {
$__t1 = new Phpurs_Data1("Inr", ($x_0)->{'value0'});
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Maybe_functorMaybe'] = function() { $v = (object)["map" => (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ((is_object($v1_1) && (($v1_1)->{'tag'} === "Just"))) {
$__t0 = new Phpurs_Data1("Just", ($v_0)(($v1_1)->{'value0'}));
goto end_branch_0;;
};
  $__t0 = new Phpurs_Data0("Nothing");
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Data_Maybe_invariantMaybe'] = function() { $v = (object)["imap" => (function() {
  $__fn = function($f_0 = null, $v_1 = null, $v1_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if ((is_object($v1_2) && (($v1_2)->{'tag'} === "Just"))) {
$__t0 = new Phpurs_Data1("Just", ($f_0)(($v1_2)->{'value0'}));
goto end_branch_0;;
};
  $__t0 = new Phpurs_Data0("Nothing");
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Data_Maybe_fromMaybe__prime__'] = function() { $v = function($a_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Maybe_maybe__prime__'] ?? \PhpursThunks::eval('Data_Maybe_maybe__prime__')))($a_0))((($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))->{'identity'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Maybe_fromMaybe'] = function() { $v = (function() {
  $__fn = function($a_0 = null, $v2_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ((is_object($v2_1) && (($v2_1)->{'tag'} === "Nothing"))) {
$__t0 = $a_0;
goto end_branch_0;;
};
  if ((is_object($v2_1) && (($v2_1)->{'tag'} === "Just"))) {
$__t0 = ($v2_1)->{'value0'};
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
})(); return $v; };
\PhpursThunks::$thunks['Data_Maybe_fromJust'] = function() { $v = (function() {
  $__fn = function($dollar__unused_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ((is_object($v_1) && (($v_1)->{'tag'} === "Just"))) {
$__t0 = ($v_1)->{'value0'};
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
})(); return $v; };
\PhpursThunks::$thunks['Data_Maybe_extendMaybe'] = function() { $v = (object)["extend" => (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ((is_object($v1_1) && (($v1_1)->{'tag'} === "Nothing"))) {
$__t0 = new Phpurs_Data0("Nothing");
goto end_branch_0;;
};
  $__t0 = new Phpurs_Data1("Just", ($v_0)($v1_1));
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Maybe_functorMaybe'] ?? \PhpursThunks::eval('Data_Maybe_functorMaybe'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Maybe_eqMaybe'] = function() { $v = function($dictEq_0 = null) {
  $__num = \func_num_args();
  $__res = (object)["eq" => (function() use ($dictEq_0) {
  $__fn = function($x_1 = null, $y_2 = null) use ($dictEq_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ((is_object($x_1) && (($x_1)->{'tag'} === "Nothing"))) {
$__t0 = (is_object($y_2) && (($y_2)->{'tag'} === "Nothing"));
goto end_branch_0;;
};
  $__t0 = ((is_object($x_1) && (($x_1)->{'tag'} === "Just")) && ((is_object($y_2) && (($y_2)->{'tag'} === "Just")) && ((($dictEq_0)->{'eq'})(($x_1)->{'value0'}))(($y_2)->{'value0'})));
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Maybe_ordMaybe'] = function() { $v = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($dictOrd_0)->{'Eq0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $eqMaybe1_2_1 = (object)["eq" => (function() use ($__local_var_1_0) {
  $__fn = function($x_2 = null, $y_3 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t1 = null;;
  if ((is_object($x_2) && (($x_2)->{'tag'} === "Nothing"))) {
$__t1 = (is_object($y_3) && (($y_3)->{'tag'} === "Nothing"));
goto end_branch_1;;
};
  $__t1 = ((is_object($x_2) && (($x_2)->{'tag'} === "Just")) && ((is_object($y_3) && (($y_3)->{'tag'} === "Just")) && ((($__local_var_1_0)->{'eq'})(($x_2)->{'value0'}))(($y_3)->{'value0'})));
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $__res = (object)["compare" => (function() use ($dictOrd_0) {
  $__fn = function($x_3 = null, $y_4 = null) use ($dictOrd_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t3 = null;;
  if ((is_object($x_3) && (($x_3)->{'tag'} === "Nothing"))) {
$__t4 = null;;
if ((is_object($y_4) && (($y_4)->{'tag'} === "Nothing"))) {
$__t4 = new Phpurs_Data0("EQ");
goto end_branch_4;;
};
$__t4 = new Phpurs_Data0("LT");
end_branch_4:;
$__t3 = $__t4;
goto end_branch_3;;
};
  if ((is_object($y_4) && (($y_4)->{'tag'} === "Nothing"))) {
$__t3 = new Phpurs_Data0("GT");
goto end_branch_3;;
};
  if (((is_object($x_3) && (($x_3)->{'tag'} === "Just")) && (is_object($y_4) && (($y_4)->{'tag'} === "Just")))) {
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
})(), "Eq0" => function($dollar__unused_3 = null) use ($eqMaybe1_2_1) {
  $__num = \func_num_args();
  $__res = $eqMaybe1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Maybe_eq1Maybe'] = function() { $v = (object)["eq1" => (function() {
  $__fn = function($dictEq_0 = null, $x_1 = null, $y_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if ((is_object($x_1) && (($x_1)->{'tag'} === "Nothing"))) {
$__t0 = (is_object($y_2) && (($y_2)->{'tag'} === "Nothing"));
goto end_branch_0;;
};
  $__t0 = ((is_object($x_1) && (($x_1)->{'tag'} === "Just")) && ((is_object($y_2) && (($y_2)->{'tag'} === "Just")) && ((($dictEq_0)->{'eq'})(($x_1)->{'value0'}))(($y_2)->{'value0'})));
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Data_Maybe_ord1Maybe'] = function() { $v = (object)["compare1" => function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Maybe_ordMaybe'] ?? \PhpursThunks::eval('Data_Maybe_ordMaybe')))($dictOrd_0))->{'compare'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq10" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Maybe_eq1Maybe'] ?? \PhpursThunks::eval('Data_Maybe_eq1Maybe'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Maybe_boundedMaybe'] = function() { $v = function($dictBounded_0 = null) {
  $__num = \func_num_args();
  $ordMaybe1_1_0 = (($GLOBALS['Data_Maybe_ordMaybe'] ?? \PhpursThunks::eval('Data_Maybe_ordMaybe')))((($dictBounded_0)->{'Ord0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = (object)["top" => new Phpurs_Data1("Just", ($dictBounded_0)->{'top'}), "bottom" => new Phpurs_Data0("Nothing"), "Ord0" => function($dollar__unused_2 = null) use ($ordMaybe1_1_0) {
  $__num = \func_num_args();
  $__res = $ordMaybe1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Maybe_applyMaybe'] = function() { $v = (object)["apply" => (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Just"))) {
$__t1 = null;;
if ((is_object($v1_1) && (($v1_1)->{'tag'} === "Just"))) {
$__t1 = new Phpurs_Data1("Just", (($v_0)->{'value0'})(($v1_1)->{'value0'}));
goto end_branch_1;;
};
$__t1 = new Phpurs_Data0("Nothing");
end_branch_1:;
$__t0 = $__t1;
goto end_branch_0;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Nothing"))) {
$__t0 = new Phpurs_Data0("Nothing");
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
})(), "Functor0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Maybe_functorMaybe'] ?? \PhpursThunks::eval('Data_Maybe_functorMaybe'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Maybe_bindMaybe'] = function() { $v = (object)["bind" => (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Just"))) {
$__t0 = ($v1_1)(($v_0)->{'value0'});
goto end_branch_0;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Nothing"))) {
$__t0 = new Phpurs_Data0("Nothing");
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
})(), "Apply0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Maybe_applyMaybe'] ?? \PhpursThunks::eval('Data_Maybe_applyMaybe'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Maybe_semiringMaybe'] = function() { $v = function($dictSemiring_0 = null) {
  $__num = \func_num_args();
  $__res = (object)["zero" => new Phpurs_Data0("Nothing"), "one" => new Phpurs_Data1("Just", ($dictSemiring_0)->{'one'}), "add" => (function() use ($dictSemiring_0) {
  $__fn = function($v_1 = null, $v1_2 = null) use ($dictSemiring_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ((is_object($v_1) && (($v_1)->{'tag'} === "Nothing"))) {
$__t0 = $v1_2;
goto end_branch_0;;
};
  if ((is_object($v1_2) && (($v1_2)->{'tag'} === "Nothing"))) {
$__t0 = $v_1;
goto end_branch_0;;
};
  if (((is_object($v_1) && (($v_1)->{'tag'} === "Just")) && (is_object($v1_2) && (($v1_2)->{'tag'} === "Just")))) {
$__t0 = new Phpurs_Data1("Just", ((($dictSemiring_0)->{'add'})(($v_1)->{'value0'}))(($v1_2)->{'value0'}));
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
})(), "mul" => (function() use ($dictSemiring_0) {
  $__fn = function($x_1 = null, $y_2 = null) use ($dictSemiring_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t1 = null;;
  if (((is_object($x_1) && (($x_1)->{'tag'} === "Just")) && (is_object($y_2) && (($y_2)->{'tag'} === "Just")))) {
$__t1 = new Phpurs_Data1("Just", ((($dictSemiring_0)->{'mul'})(($x_1)->{'value0'}))(($y_2)->{'value0'}));
goto end_branch_1;;
};
  $__t1 = new Phpurs_Data0("Nothing");
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Maybe_applicativeMaybe'] = function() { $v = (object)["pure" => ($GLOBALS['Data_Maybe_Just'] ?? \PhpursThunks::eval('Data_Maybe_Just')), "Apply0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Maybe_applyMaybe'] ?? \PhpursThunks::eval('Data_Maybe_applyMaybe'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Maybe_monadMaybe'] = function() { $v = (object)["Applicative0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Maybe_applicativeMaybe'] ?? \PhpursThunks::eval('Data_Maybe_applicativeMaybe'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Maybe_bindMaybe'] ?? \PhpursThunks::eval('Data_Maybe_bindMaybe'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Maybe_altMaybe'] = function() { $v = (object)["alt" => (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Nothing"))) {
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
})(), "Functor0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Maybe_functorMaybe'] ?? \PhpursThunks::eval('Data_Maybe_functorMaybe'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Maybe_plusMaybe'] = function() { $v = (object)["empty" => new Phpurs_Data0("Nothing"), "Alt0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Maybe_altMaybe'] ?? \PhpursThunks::eval('Data_Maybe_altMaybe'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Maybe_alternativeMaybe'] = function() { $v = (object)["Applicative0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Maybe_applicativeMaybe'] ?? \PhpursThunks::eval('Data_Maybe_applicativeMaybe'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Plus1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Maybe_plusMaybe'] ?? \PhpursThunks::eval('Data_Maybe_plusMaybe'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };
































