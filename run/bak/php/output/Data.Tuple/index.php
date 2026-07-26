<?php

namespace Data\Tuple;

// ALL IMPORTS: Control.Applicative, Control.Apply, Control.Bind, Control.Comonad, Control.Extend, Control.Lazy, Control.Monad, Control.Semigroupoid, Data.BooleanAlgebra, Data.Bounded, Data.CommutativeRing, Data.Eq, Data.Function, Data.Functor, Data.Functor.Invariant, Data.Generic.Rep, Data.HeytingAlgebra, Data.Monoid, Data.Ord, Data.Ordering, Data.Ring, Data.Semigroup, Data.Semiring, Data.Show, Data.Tuple, Data.Unit, Prelude, Prim
// TO REQUIRE: Control.Applicative, Control.Apply, Control.Bind, Control.Comonad, Control.Extend, Control.Lazy, Control.Monad, Control.Semigroupoid, Data.BooleanAlgebra, Data.Bounded, Data.CommutativeRing, Data.Eq, Data.Function, Data.Functor, Data.Functor.Invariant, Data.Generic.Rep, Data.HeytingAlgebra, Data.Monoid, Data.Ord, Data.Ordering, Data.Ring, Data.Semigroup, Data.Semiring, Data.Show, Data.Tuple, Data.Unit, Prelude
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Comonad/index.php';
require_once __DIR__ . '/../Control.Extend/index.php';
require_once __DIR__ . '/../Control.Lazy/index.php';
require_once __DIR__ . '/../Control.Monad/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.BooleanAlgebra/index.php';
require_once __DIR__ . '/../Data.Bounded/index.php';
require_once __DIR__ . '/../Data.CommutativeRing/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Functor.Invariant/index.php';
require_once __DIR__ . '/../Data.Generic.Rep/index.php';
require_once __DIR__ . '/../Data.HeytingAlgebra/index.php';
require_once __DIR__ . '/../Data.Monoid/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ordering/index.php';
require_once __DIR__ . '/../Data.Ring/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Semiring/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
require_once __DIR__ . '/../Data.Tuple/index.php';
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


// Data_Tuple_Tuple
$GLOBALS['Data_Tuple_Tuple'] = (function() {
  $__fn = function($value0 = null, $value1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("Tuple", $value0, $value1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Tuple_uncurry
$GLOBALS['Data_Tuple_uncurry'] = (function() {
  $__fn = function($f_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($f_0)(($v_1)->{'value0'}))(($v_1)->{'value1'});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Tuple_swap
$GLOBALS['Data_Tuple_swap'] = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Tuple", ($v_0)->{'value1'}, ($v_0)->{'value0'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Tuple_snd
$GLOBALS['Data_Tuple_snd'] = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ($v_0)->{'value1'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Tuple_showTuple
$GLOBALS['Data_Tuple_showTuple'] = (function() {
  $__fn = function($dictShow_0 = null, $dictShow1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ["show" => function($v_2 = null) use ($dictShow1_1, $dictShow_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Semigroup_semigroupString'])['append'])("(Tuple "))(((($GLOBALS['Data_Semigroup_semigroupString'])['append'])((($dictShow_0)['show'])(($v_2)->{'value0'})))(((($GLOBALS['Data_Semigroup_semigroupString'])['append'])(" "))(((($GLOBALS['Data_Semigroup_semigroupString'])['append'])((($dictShow1_1)['show'])(($v_2)->{'value1'})))(")"))));
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

// Data_Tuple_semiringTuple
$GLOBALS['Data_Tuple_semiringTuple'] = function($dictSemiring_0 = null) {
  $__num = \func_num_args();
  $one_1_0 = ($dictSemiring_0)['one'];
  $zero_2_1 = ($dictSemiring_0)['zero'];
  $__res = function($dictSemiring1_3 = null) use ($dictSemiring_0, $one_1_0, $zero_2_1) {
  $__num = \func_num_args();
  $__res = ["add" => (function() use ($dictSemiring1_3, $dictSemiring_0) {
  $__fn = function($v_4 = null, $v1_5 = null) use ($dictSemiring1_3, $dictSemiring_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("Tuple", ((($dictSemiring_0)['add'])(($v_4)->{'value0'}))(($v1_5)->{'value0'}), ((($dictSemiring1_3)['add'])(($v_4)->{'value1'}))(($v1_5)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "one" => new Phpurs_Data2("Tuple", $one_1_0, ($dictSemiring1_3)['one']), "mul" => (function() use ($dictSemiring1_3, $dictSemiring_0) {
  $__fn = function($v_4 = null, $v1_5 = null) use ($dictSemiring1_3, $dictSemiring_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("Tuple", ((($dictSemiring_0)['mul'])(($v_4)->{'value0'}))(($v1_5)->{'value0'}), ((($dictSemiring1_3)['mul'])(($v_4)->{'value1'}))(($v1_5)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "zero" => new Phpurs_Data2("Tuple", $zero_2_1, ($dictSemiring1_3)['zero'])];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Tuple_semigroupoidTuple
$GLOBALS['Data_Tuple_semigroupoidTuple'] = ["compose" => (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("Tuple", ($v1_1)->{'value0'}, ($v_0)->{'value1'});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];

// Data_Tuple_semigroupTuple
$GLOBALS['Data_Tuple_semigroupTuple'] = (function() {
  $__fn = function($dictSemigroup_0 = null, $dictSemigroup1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ["append" => (function() use ($dictSemigroup1_1, $dictSemigroup_0) {
  $__fn = function($v_2 = null, $v1_3 = null) use ($dictSemigroup1_1, $dictSemigroup_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("Tuple", ((($dictSemigroup_0)['append'])(($v_2)->{'value0'}))(($v1_3)->{'value0'}), ((($dictSemigroup1_1)['append'])(($v_2)->{'value1'}))(($v1_3)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Tuple_ringTuple
$GLOBALS['Data_Tuple_ringTuple'] = function($dictRing_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($dictRing_0)['Semiring0'])(null);
  $one_2_1 = ($__local_var_1_0)['one'];
  $zero_3_2 = ($__local_var_1_0)['zero'];
  $semiringTuple1_3_2 = function($dictSemiring1_4 = null) use ($__local_var_1_0, $one_2_1, $zero_3_2) {
  $__num = \func_num_args();
  $__res = ["add" => (function() use ($__local_var_1_0, $dictSemiring1_4) {
  $__fn = function($v_5 = null, $v1_6 = null) use ($__local_var_1_0, $dictSemiring1_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("Tuple", ((($__local_var_1_0)['add'])(($v_5)->{'value0'}))(($v1_6)->{'value0'}), ((($dictSemiring1_4)['add'])(($v_5)->{'value1'}))(($v1_6)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "one" => new Phpurs_Data2("Tuple", $one_2_1, ($dictSemiring1_4)['one']), "mul" => (function() use ($__local_var_1_0, $dictSemiring1_4) {
  $__fn = function($v_5 = null, $v1_6 = null) use ($__local_var_1_0, $dictSemiring1_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("Tuple", ((($__local_var_1_0)['mul'])(($v_5)->{'value0'}))(($v1_6)->{'value0'}), ((($dictSemiring1_4)['mul'])(($v_5)->{'value1'}))(($v1_6)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "zero" => new Phpurs_Data2("Tuple", $zero_3_2, ($dictSemiring1_4)['zero'])];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = function($dictRing1_4 = null) use ($dictRing_0, $semiringTuple1_3_2) {
  $__num = \func_num_args();
  $semiringTuple2_5_4 = ($semiringTuple1_3_2)((($dictRing1_4)['Semiring0'])(null));
  $__res = ["sub" => (function() use ($dictRing1_4, $dictRing_0) {
  $__fn = function($v_6 = null, $v1_7 = null) use ($dictRing1_4, $dictRing_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("Tuple", ((($dictRing_0)['sub'])(($v_6)->{'value0'}))(($v1_7)->{'value0'}), ((($dictRing1_4)['sub'])(($v_6)->{'value1'}))(($v1_7)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Semiring0" => function($_dollar__unused_6 = null) use ($semiringTuple2_5_4) {
  $__num = \func_num_args();
  $__res = $semiringTuple2_5_4;
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Tuple_monoidTuple
$GLOBALS['Data_Tuple_monoidTuple'] = function($dictMonoid_0 = null) {
  $__num = \func_num_args();
  $mempty_1_0 = ($dictMonoid_0)['mempty'];
  $__local_var_2_1 = (($dictMonoid_0)['Semigroup0'])(null);
  $__res = function($dictMonoid1_3 = null) use ($__local_var_2_1, $mempty_1_0) {
  $__num = \func_num_args();
  $__local_var_4_2 = (($dictMonoid1_3)['Semigroup0'])(null);
  $semigroupTuple2_5_3 = ["append" => (function() use ($__local_var_2_1, $__local_var_4_2) {
  $__fn = function($v_5 = null, $v1_6 = null) use ($__local_var_2_1, $__local_var_4_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("Tuple", ((($__local_var_2_1)['append'])(($v_5)->{'value0'}))(($v1_6)->{'value0'}), ((($__local_var_4_2)['append'])(($v_5)->{'value1'}))(($v1_6)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $__res = ["mempty" => new Phpurs_Data2("Tuple", $mempty_1_0, ($dictMonoid1_3)['mempty']), "Semigroup0" => function($_dollar__unused_6 = null) use ($semigroupTuple2_5_3) {
  $__num = \func_num_args();
  $__res = $semigroupTuple2_5_3;
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Tuple_heytingAlgebraTuple
$GLOBALS['Data_Tuple_heytingAlgebraTuple'] = function($dictHeytingAlgebra_0 = null) {
  $__num = \func_num_args();
  $tt_1_0 = ($dictHeytingAlgebra_0)['tt'];
  $ff_2_1 = ($dictHeytingAlgebra_0)['ff'];
  $__res = function($dictHeytingAlgebra1_3 = null) use ($dictHeytingAlgebra_0, $ff_2_1, $tt_1_0) {
  $__num = \func_num_args();
  $__res = ["tt" => new Phpurs_Data2("Tuple", $tt_1_0, ($dictHeytingAlgebra1_3)['tt']), "ff" => new Phpurs_Data2("Tuple", $ff_2_1, ($dictHeytingAlgebra1_3)['ff']), "implies" => (function() use ($dictHeytingAlgebra1_3, $dictHeytingAlgebra_0) {
  $__fn = function($v_4 = null, $v1_5 = null) use ($dictHeytingAlgebra1_3, $dictHeytingAlgebra_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("Tuple", ((($dictHeytingAlgebra_0)['implies'])(($v_4)->{'value0'}))(($v1_5)->{'value0'}), ((($dictHeytingAlgebra1_3)['implies'])(($v_4)->{'value1'}))(($v1_5)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "conj" => (function() use ($dictHeytingAlgebra1_3, $dictHeytingAlgebra_0) {
  $__fn = function($v_4 = null, $v1_5 = null) use ($dictHeytingAlgebra1_3, $dictHeytingAlgebra_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("Tuple", ((($dictHeytingAlgebra_0)['conj'])(($v_4)->{'value0'}))(($v1_5)->{'value0'}), ((($dictHeytingAlgebra1_3)['conj'])(($v_4)->{'value1'}))(($v1_5)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "disj" => (function() use ($dictHeytingAlgebra1_3, $dictHeytingAlgebra_0) {
  $__fn = function($v_4 = null, $v1_5 = null) use ($dictHeytingAlgebra1_3, $dictHeytingAlgebra_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("Tuple", ((($dictHeytingAlgebra_0)['disj'])(($v_4)->{'value0'}))(($v1_5)->{'value0'}), ((($dictHeytingAlgebra1_3)['disj'])(($v_4)->{'value1'}))(($v1_5)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "not" => function($v_4 = null) use ($dictHeytingAlgebra1_3, $dictHeytingAlgebra_0) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Tuple", (($dictHeytingAlgebra_0)['not'])(($v_4)->{'value0'}), (($dictHeytingAlgebra1_3)['not'])(($v_4)->{'value1'}));
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Tuple_genericTuple
$GLOBALS['Data_Tuple_genericTuple'] = ["to" => function($x_0 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Tuple", ($x_0)->{'value0'}, ($x_0)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "from" => function($x_0 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Product", ($x_0)->{'value0'}, ($x_0)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Tuple_functorTuple
$GLOBALS['Data_Tuple_functorTuple'] = ["map" => (function() {
  $__fn = function($f_0 = null, $m_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("Tuple", ($m_1)->{'value0'}, ($f_0)(($m_1)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];

// Data_Tuple_invariantTuple
$GLOBALS['Data_Tuple_invariantTuple'] = ["imap" => (function() {
  $__fn = function($f_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_Tuple_functorTuple'])['map'])($f_0);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];

// Data_Tuple_fst
$GLOBALS['Data_Tuple_fst'] = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ($v_0)->{'value0'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Tuple_lazyTuple
$GLOBALS['Data_Tuple_lazyTuple'] = (function() {
  $__fn = function($dictLazy_0 = null, $dictLazy1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ["defer" => function($f_2 = null) use ($dictLazy1_1, $dictLazy_0) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Tuple", (($dictLazy_0)['defer'])(function($v_3 = null) use ($f_2) {
  $__num = \func_num_args();
  $__res = (($f_2)($GLOBALS['Data_Unit_unit']))->{'value0'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), (($dictLazy1_1)['defer'])(function($v_3 = null) use ($f_2) {
  $__num = \func_num_args();
  $__res = (($f_2)($GLOBALS['Data_Unit_unit']))->{'value1'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
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

// Data_Tuple_extendTuple
$GLOBALS['Data_Tuple_extendTuple'] = ["extend" => (function() {
  $__fn = function($f_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("Tuple", ($v_1)->{'value0'}, ($f_0)($v_1));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Tuple_functorTuple'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Tuple_eqTuple
$GLOBALS['Data_Tuple_eqTuple'] = (function() {
  $__fn = function($dictEq_0 = null, $dictEq1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ["eq" => (function() use ($dictEq1_1, $dictEq_0) {
  $__fn = function($x_2 = null, $y_3 = null) use ($dictEq1_1, $dictEq_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])['conj'])(((($dictEq_0)['eq'])(($x_2)->{'value0'}))(($y_3)->{'value0'})))(((($dictEq1_1)['eq'])(($x_2)->{'value1'}))(($y_3)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Tuple_ordTuple
$GLOBALS['Data_Tuple_ordTuple'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $eqTuple1_1_0 = ($GLOBALS['Data_Tuple_eqTuple'])((($dictOrd_0)['Eq0'])(null));
  $__res = function($dictOrd1_2 = null) use ($dictOrd_0, $eqTuple1_1_0) {
  $__num = \func_num_args();
  $eqTuple2_3_1 = ($eqTuple1_1_0)((($dictOrd1_2)['Eq0'])(null));
  $__res = ["compare" => (function() use ($dictOrd1_2, $dictOrd_0) {
  $__fn = function($x_4 = null, $y_5 = null) use ($dictOrd1_2, $dictOrd_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v_6_2 = ((($dictOrd_0)['compare'])(($x_4)->{'value0'}))(($y_5)->{'value0'});
  $__t3 = null;;
  if ((is_object($v_6_2) && (($v_6_2)->{'tag'} === "LT"))) {
$__t3 = new Phpurs_Data0("LT");
goto end_branch_3;;
};
  if ((is_object($v_6_2) && (($v_6_2)->{'tag'} === "GT"))) {
$__t3 = new Phpurs_Data0("GT");
goto end_branch_3;;
};
  $__t3 = ((($dictOrd1_2)['compare'])(($x_4)->{'value1'}))(($y_5)->{'value1'});
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($_dollar__unused_4 = null) use ($eqTuple2_3_1) {
  $__num = \func_num_args();
  $__res = $eqTuple2_3_1;
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Tuple_eq1Tuple
$GLOBALS['Data_Tuple_eq1Tuple'] = function($dictEq_0 = null) {
  $__num = \func_num_args();
  $__res = ["eq1" => function($dictEq1_1 = null) use ($dictEq_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Tuple_eqTuple'])($dictEq_0))($dictEq1_1))['eq'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Tuple_ord1Tuple
$GLOBALS['Data_Tuple_ord1Tuple'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $ordTuple1_1_0 = ($GLOBALS['Data_Tuple_ordTuple'])($dictOrd_0);
  $__local_var_2_1 = (($dictOrd_0)['Eq0'])(null);
  $eq1Tuple1_3_2 = ["eq1" => function($dictEq1_3 = null) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Tuple_eqTuple'])($__local_var_2_1))($dictEq1_3))['eq'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ["compare1" => function($dictOrd1_4 = null) use ($ordTuple1_1_0) {
  $__num = \func_num_args();
  $__res = (($ordTuple1_1_0)($dictOrd1_4))['compare'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq10" => function($_dollar__unused_4 = null) use ($eq1Tuple1_3_2) {
  $__num = \func_num_args();
  $__res = $eq1Tuple1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Tuple_curry
$GLOBALS['Data_Tuple_curry'] = (function() {
  $__fn = function($f_0 = null, $a_1 = null, $b_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($f_0)(new Phpurs_Data2("Tuple", $a_1, $b_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_Tuple_comonadTuple
$GLOBALS['Data_Tuple_comonadTuple'] = ["extract" => $GLOBALS['Data_Tuple_snd'], "Extend0" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Tuple_extendTuple'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Tuple_commutativeRingTuple
$GLOBALS['Data_Tuple_commutativeRingTuple'] = function($dictCommutativeRing_0 = null) {
  $__num = \func_num_args();
  $ringTuple1_1_0 = ($GLOBALS['Data_Tuple_ringTuple'])((($dictCommutativeRing_0)['Ring0'])(null));
  $__res = function($dictCommutativeRing1_2 = null) use ($ringTuple1_1_0) {
  $__num = \func_num_args();
  $ringTuple2_3_1 = ($ringTuple1_1_0)((($dictCommutativeRing1_2)['Ring0'])(null));
  $__res = ["Ring0" => function($_dollar__unused_4 = null) use ($ringTuple2_3_1) {
  $__num = \func_num_args();
  $__res = $ringTuple2_3_1;
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Tuple_boundedTuple
$GLOBALS['Data_Tuple_boundedTuple'] = function($dictBounded_0 = null) {
  $__num = \func_num_args();
  $top_1_0 = ($dictBounded_0)['top'];
  $bottom_2_1 = ($dictBounded_0)['bottom'];
  $ordTuple1_3_2 = ($GLOBALS['Data_Tuple_ordTuple'])((($dictBounded_0)['Ord0'])(null));
  $__res = function($dictBounded1_4 = null) use ($bottom_2_1, $ordTuple1_3_2, $top_1_0) {
  $__num = \func_num_args();
  $ordTuple2_5_3 = ($ordTuple1_3_2)((($dictBounded1_4)['Ord0'])(null));
  $__res = ["top" => new Phpurs_Data2("Tuple", $top_1_0, ($dictBounded1_4)['top']), "bottom" => new Phpurs_Data2("Tuple", $bottom_2_1, ($dictBounded1_4)['bottom']), "Ord0" => function($_dollar__unused_6 = null) use ($ordTuple2_5_3) {
  $__num = \func_num_args();
  $__res = $ordTuple2_5_3;
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Tuple_booleanAlgebraTuple
$GLOBALS['Data_Tuple_booleanAlgebraTuple'] = function($dictBooleanAlgebra_0 = null) {
  $__num = \func_num_args();
  $heytingAlgebraTuple1_1_0 = ($GLOBALS['Data_Tuple_heytingAlgebraTuple'])((($dictBooleanAlgebra_0)['HeytingAlgebra0'])(null));
  $__res = function($dictBooleanAlgebra1_2 = null) use ($heytingAlgebraTuple1_1_0) {
  $__num = \func_num_args();
  $heytingAlgebraTuple2_3_1 = ($heytingAlgebraTuple1_1_0)((($dictBooleanAlgebra1_2)['HeytingAlgebra0'])(null));
  $__res = ["HeytingAlgebra0" => function($_dollar__unused_4 = null) use ($heytingAlgebraTuple2_3_1) {
  $__num = \func_num_args();
  $__res = $heytingAlgebraTuple2_3_1;
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Tuple_applyTuple
$GLOBALS['Data_Tuple_applyTuple'] = function($dictSemigroup_0 = null) {
  $__num = \func_num_args();
  $__res = ["apply" => (function() use ($dictSemigroup_0) {
  $__fn = function($v_1 = null, $v1_2 = null) use ($dictSemigroup_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("Tuple", ((($dictSemigroup_0)['append'])(($v_1)->{'value0'}))(($v1_2)->{'value0'}), (($v_1)->{'value1'})(($v1_2)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_1 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Tuple_functorTuple'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Tuple_bindTuple
$GLOBALS['Data_Tuple_bindTuple'] = function($dictSemigroup_0 = null) {
  $__num = \func_num_args();
  $applyTuple1_1_0 = ($GLOBALS['Data_Tuple_applyTuple'])($dictSemigroup_0);
  $__res = ["bind" => (function() use ($dictSemigroup_0) {
  $__fn = function($v_2 = null, $f_3 = null) use ($dictSemigroup_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v1_4_1 = ($f_3)(($v_2)->{'value1'});
  $__res = new Phpurs_Data2("Tuple", ((($dictSemigroup_0)['append'])(($v_2)->{'value0'}))(($v1_4_1)->{'value0'}), ($v1_4_1)->{'value1'});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Apply0" => function($_dollar__unused_2 = null) use ($applyTuple1_1_0) {
  $__num = \func_num_args();
  $__res = $applyTuple1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Tuple_applicativeTuple
$GLOBALS['Data_Tuple_applicativeTuple'] = function($dictMonoid_0 = null) {
  $__num = \func_num_args();
  $applyTuple1_1_0 = ($GLOBALS['Data_Tuple_applyTuple'])((($dictMonoid_0)['Semigroup0'])(null));
  $__res = ["pure" => ($GLOBALS['Data_Tuple_Tuple'])(($dictMonoid_0)['mempty']), "Apply0" => function($_dollar__unused_2 = null) use ($applyTuple1_1_0) {
  $__num = \func_num_args();
  $__res = $applyTuple1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Tuple_monadTuple
$GLOBALS['Data_Tuple_monadTuple'] = function($dictMonoid_0 = null) {
  $__num = \func_num_args();
  $applicativeTuple1_1_0 = ($GLOBALS['Data_Tuple_applicativeTuple'])($dictMonoid_0);
  $bindTuple1_2_1 = ($GLOBALS['Data_Tuple_bindTuple'])((($dictMonoid_0)['Semigroup0'])(null));
  $__res = ["Applicative0" => function($_dollar__unused_3 = null) use ($applicativeTuple1_1_0) {
  $__num = \func_num_args();
  $__res = $applicativeTuple1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar__unused_3 = null) use ($bindTuple1_2_1) {
  $__num = \func_num_args();
  $__res = $bindTuple1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

