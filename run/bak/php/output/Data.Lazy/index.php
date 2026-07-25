<?php

namespace Data\Lazy;

// ALL IMPORTS: Control.Applicative, Control.Apply, Control.Bind, Control.Comonad, Control.Extend, Control.Lazy, Control.Monad, Control.Semigroupoid, Data.BooleanAlgebra, Data.Bounded, Data.CommutativeRing, Data.Eq, Data.EuclideanRing, Data.Foldable, Data.FoldableWithIndex, Data.Function, Data.Functor, Data.Functor.Invariant, Data.FunctorWithIndex, Data.HeytingAlgebra, Data.Lazy, Data.Monoid, Data.Ord, Data.Ring, Data.Semigroup, Data.Semigroup.Foldable, Data.Semigroup.Traversable, Data.Semiring, Data.Show, Data.Traversable, Data.TraversableWithIndex, Data.Unit, Prelude, Prim
// TO REQUIRE: Control.Applicative, Control.Apply, Control.Bind, Control.Comonad, Control.Extend, Control.Lazy, Control.Monad, Control.Semigroupoid, Data.BooleanAlgebra, Data.Bounded, Data.CommutativeRing, Data.Eq, Data.EuclideanRing, Data.Foldable, Data.FoldableWithIndex, Data.Function, Data.Functor, Data.Functor.Invariant, Data.FunctorWithIndex, Data.HeytingAlgebra, Data.Lazy, Data.Monoid, Data.Ord, Data.Ring, Data.Semigroup, Data.Semigroup.Foldable, Data.Semigroup.Traversable, Data.Semiring, Data.Show, Data.Traversable, Data.TraversableWithIndex, Data.Unit, Prelude
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
require_once __DIR__ . '/../Data.EuclideanRing/index.php';
require_once __DIR__ . '/../Data.Foldable/index.php';
require_once __DIR__ . '/../Data.FoldableWithIndex/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Functor.Invariant/index.php';
require_once __DIR__ . '/../Data.FunctorWithIndex/index.php';
require_once __DIR__ . '/../Data.HeytingAlgebra/index.php';
require_once __DIR__ . '/../Data.Lazy/index.php';
require_once __DIR__ . '/../Data.Monoid/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ring/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Semigroup.Foldable/index.php';
require_once __DIR__ . '/../Data.Semigroup.Traversable/index.php';
require_once __DIR__ . '/../Data.Semiring/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
require_once __DIR__ . '/../Data.Traversable/index.php';
require_once __DIR__ . '/../Data.TraversableWithIndex/index.php';
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
\PhpursThunks::$thunks['Data_Lazy_showLazy'] = function() { $v = function($dictShow_0 = null) {
  $__num = \func_num_args();
  $__res = (object)["show" => function($x_1 = null) use ($dictShow_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))("(defer \\_ -> "))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))((($dictShow_0)->{'show'})((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($x_1))))(")"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Lazy_semiringLazy'] = function() { $v = function($dictSemiring_0 = null) {
  $__num = \func_num_args();
  $zero_1_0 = ($dictSemiring_0)->{'zero'};
  $one_2_1 = ($dictSemiring_0)->{'one'};
  $__res = (object)["add" => (function() use ($dictSemiring_0) {
  $__fn = function($a_3 = null, $b_4 = null) use ($dictSemiring_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_5 = null) use ($a_3, $b_4, $dictSemiring_0) {
  $__num = \func_num_args();
  $__res = ((($dictSemiring_0)->{'add'})((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($a_3)))((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($b_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "zero" => (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_3 = null) use ($zero_1_0) {
  $__num = \func_num_args();
  $__res = $zero_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), "mul" => (function() use ($dictSemiring_0) {
  $__fn = function($a_3 = null, $b_4 = null) use ($dictSemiring_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_5 = null) use ($a_3, $b_4, $dictSemiring_0) {
  $__num = \func_num_args();
  $__res = ((($dictSemiring_0)->{'mul'})((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($a_3)))((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($b_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "one" => (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_3 = null) use ($one_2_1) {
  $__num = \func_num_args();
  $__res = $one_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Lazy_semigroupLazy'] = function() { $v = function($dictSemigroup_0 = null) {
  $__num = \func_num_args();
  $__res = (object)["append" => (function() use ($dictSemigroup_0) {
  $__fn = function($a_1 = null, $b_2 = null) use ($dictSemigroup_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_3 = null) use ($a_1, $b_2, $dictSemigroup_0) {
  $__num = \func_num_args();
  $__res = ((($dictSemigroup_0)->{'append'})((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($a_1)))((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($b_2));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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
\PhpursThunks::$thunks['Data_Lazy_ringLazy'] = function() { $v = function($dictRing_0 = null) {
  $__num = \func_num_args();
  $semiringLazy1_1_0 = (($GLOBALS['Data_Lazy_semiringLazy'] ?? \PhpursThunks::eval('Data_Lazy_semiringLazy')))((($dictRing_0)->{'Semiring0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = (object)["sub" => (function() use ($dictRing_0) {
  $__fn = function($a_2 = null, $b_3 = null) use ($dictRing_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_4 = null) use ($a_2, $b_3, $dictRing_0) {
  $__num = \func_num_args();
  $__res = ((($dictRing_0)->{'sub'})((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($a_2)))((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($b_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Semiring0" => function($dollar__unused_2 = null) use ($semiringLazy1_1_0) {
  $__num = \func_num_args();
  $__res = $semiringLazy1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Lazy_monoidLazy'] = function() { $v = function($dictMonoid_0 = null) {
  $__num = \func_num_args();
  $mempty_1_0 = ($dictMonoid_0)->{'mempty'};
  $semigroupLazy1_2_1 = (($GLOBALS['Data_Lazy_semigroupLazy'] ?? \PhpursThunks::eval('Data_Lazy_semigroupLazy')))((($dictMonoid_0)->{'Semigroup0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = (object)["mempty" => (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_3 = null) use ($mempty_1_0) {
  $__num = \func_num_args();
  $__res = $mempty_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), "Semigroup0" => function($dollar__unused_3 = null) use ($semigroupLazy1_2_1) {
  $__num = \func_num_args();
  $__res = $semigroupLazy1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Lazy_lazyLazy'] = function() { $v = (object)["defer" => function($f_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_1 = null) use ($f_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))(($f_0)(($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit'))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Lazy_functorLazy'] = function() { $v = (object)["map" => (function() {
  $__fn = function($f_0 = null, $l_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_2 = null) use ($f_0, $l_1) {
  $__num = \func_num_args();
  $__res = ($f_0)((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($l_1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Data_Lazy_functorWithIndexLazy'] = function() { $v = (object)["mapWithIndex" => function($f_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Lazy_functorLazy'] ?? \PhpursThunks::eval('Data_Lazy_functorLazy')))->{'map'})(($f_0)(($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit'))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Lazy_functorLazy'] ?? \PhpursThunks::eval('Data_Lazy_functorLazy'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Lazy_invariantLazy'] = function() { $v = (object)["imap" => (function() {
  $__fn = function($f_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_Lazy_functorLazy'] ?? \PhpursThunks::eval('Data_Lazy_functorLazy')))->{'map'})($f_0);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Data_Lazy_foldableLazy'] = function() { $v = (object)["foldr" => (function() {
  $__fn = function($f_0 = null, $z_1 = null, $l_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($f_0)((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($l_2)))($z_1);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldl" => (function() {
  $__fn = function($f_0 = null, $z_1 = null, $l_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($f_0)($z_1))((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($l_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldMap" => (function() {
  $__fn = function($dictMonoid_0 = null, $f_1 = null, $l_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($f_1)((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($l_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Data_Lazy_foldableWithIndexLazy'] = function() { $v = (object)["foldrWithIndex" => function($f_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Lazy_foldableLazy'] ?? \PhpursThunks::eval('Data_Lazy_foldableLazy')))->{'foldr'})(($f_0)(($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit'))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldlWithIndex" => function($f_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Lazy_foldableLazy'] ?? \PhpursThunks::eval('Data_Lazy_foldableLazy')))->{'foldl'})(($f_0)(($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit'))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldMapWithIndex" => (function() {
  $__fn = function($dictMonoid_0 = null, $f_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($GLOBALS['Data_Lazy_foldableLazy'] ?? \PhpursThunks::eval('Data_Lazy_foldableLazy')))->{'foldMap'})($dictMonoid_0))(($f_1)(($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit'))));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Foldable0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Lazy_foldableLazy'] ?? \PhpursThunks::eval('Data_Lazy_foldableLazy'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Lazy_traversableLazy'] = function() { $v = (object)["traverse" => (function() {
  $__fn = function($dictApplicative_0 = null, $f_1 = null, $l_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((((($dictApplicative_0)->{'Apply0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->{'Functor0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->{'map'})(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer'))))(($GLOBALS['Data_Function_const'] ?? \PhpursThunks::eval('Data_Function_const')))))(($f_1)((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($l_2)));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "sequence" => (function() {
  $__fn = function($dictApplicative_0 = null, $l_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((((($dictApplicative_0)->{'Apply0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->{'Functor0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->{'map'})(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer'))))(($GLOBALS['Data_Function_const'] ?? \PhpursThunks::eval('Data_Function_const')))))((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($l_1));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Lazy_functorLazy'] ?? \PhpursThunks::eval('Data_Lazy_functorLazy'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Lazy_foldableLazy'] ?? \PhpursThunks::eval('Data_Lazy_foldableLazy'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Lazy_traversableWithIndexLazy'] = function() { $v = (object)["traverseWithIndex" => (function() {
  $__fn = function($dictApplicative_0 = null, $f_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($GLOBALS['Data_Lazy_traversableLazy'] ?? \PhpursThunks::eval('Data_Lazy_traversableLazy')))->{'traverse'})($dictApplicative_0))(($f_1)(($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit'))));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "FunctorWithIndex0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Lazy_functorWithIndexLazy'] ?? \PhpursThunks::eval('Data_Lazy_functorWithIndexLazy'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Lazy_foldableWithIndexLazy'] ?? \PhpursThunks::eval('Data_Lazy_foldableWithIndexLazy'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Lazy_traversableLazy'] ?? \PhpursThunks::eval('Data_Lazy_traversableLazy'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Lazy_foldable1Lazy'] = function() { $v = (object)["foldMap1" => (function() {
  $__fn = function($dictSemigroup_0 = null, $f_1 = null, $l_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($f_1)((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($l_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldr1" => (function() {
  $__fn = function($v_0 = null, $l_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($l_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "foldl1" => (function() {
  $__fn = function($v_0 = null, $l_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($l_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Foldable0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Lazy_foldableLazy'] ?? \PhpursThunks::eval('Data_Lazy_foldableLazy'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Lazy_traversable1Lazy'] = function() { $v = (object)["traverse1" => (function() {
  $__fn = function($dictApply_0 = null, $f_1 = null, $l_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((($dictApply_0)->{'Functor0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->{'map'})(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer'))))(($GLOBALS['Data_Function_const'] ?? \PhpursThunks::eval('Data_Function_const')))))(($f_1)((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($l_2)));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "sequence1" => (function() {
  $__fn = function($dictApply_0 = null, $l_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((($dictApply_0)->{'Functor0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->{'map'})(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer'))))(($GLOBALS['Data_Function_const'] ?? \PhpursThunks::eval('Data_Function_const')))))((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($l_1));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Foldable10" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Lazy_foldable1Lazy'] ?? \PhpursThunks::eval('Data_Lazy_foldable1Lazy'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Lazy_traversableLazy'] ?? \PhpursThunks::eval('Data_Lazy_traversableLazy'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Lazy_extendLazy'] = function() { $v = (object)["extend" => (function() {
  $__fn = function($f_0 = null, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_2 = null) use ($f_0, $x_1) {
  $__num = \func_num_args();
  $__res = ($f_0)($x_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Lazy_functorLazy'] ?? \PhpursThunks::eval('Data_Lazy_functorLazy'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Lazy_eqLazy'] = function() { $v = function($dictEq_0 = null) {
  $__num = \func_num_args();
  $__res = (object)["eq" => (function() use ($dictEq_0) {
  $__fn = function($x_1 = null, $y_2 = null) use ($dictEq_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictEq_0)->{'eq'})((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($x_1)))((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($y_2));
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
\PhpursThunks::$thunks['Data_Lazy_ordLazy'] = function() { $v = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($dictOrd_0)->{'Eq0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $eqLazy1_2_1 = (object)["eq" => (function() use ($__local_var_1_0) {
  $__fn = function($x_2 = null, $y_3 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_1_0)->{'eq'})((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($x_2)))((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($y_3));
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
  $__res = ((($dictOrd_0)->{'compare'})((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($x_3)))((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($y_4));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($dollar__unused_3 = null) use ($eqLazy1_2_1) {
  $__num = \func_num_args();
  $__res = $eqLazy1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Lazy_eq1Lazy'] = function() { $v = (object)["eq1" => (function() {
  $__fn = function($dictEq_0 = null, $x_1 = null, $y_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictEq_0)->{'eq'})((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($x_1)))((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($y_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Data_Lazy_ord1Lazy'] = function() { $v = (object)["compare1" => function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Lazy_ordLazy'] ?? \PhpursThunks::eval('Data_Lazy_ordLazy')))($dictOrd_0))->{'compare'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq10" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Lazy_eq1Lazy'] ?? \PhpursThunks::eval('Data_Lazy_eq1Lazy'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Lazy_comonadLazy'] = function() { $v = (object)["extract" => ($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')), "Extend0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Lazy_extendLazy'] ?? \PhpursThunks::eval('Data_Lazy_extendLazy'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Lazy_commutativeRingLazy'] = function() { $v = function($dictCommutativeRing_0 = null) {
  $__num = \func_num_args();
  $ringLazy1_1_0 = (($GLOBALS['Data_Lazy_ringLazy'] ?? \PhpursThunks::eval('Data_Lazy_ringLazy')))((($dictCommutativeRing_0)->{'Ring0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = (object)["Ring0" => function($dollar__unused_2 = null) use ($ringLazy1_1_0) {
  $__num = \func_num_args();
  $__res = $ringLazy1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Lazy_euclideanRingLazy'] = function() { $v = function($dictEuclideanRing_0 = null) {
  $__num = \func_num_args();
  $ringLazy1_1_0 = (($GLOBALS['Data_Lazy_ringLazy'] ?? \PhpursThunks::eval('Data_Lazy_ringLazy')))((((($dictEuclideanRing_0)->{'CommutativeRing0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->{'Ring0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = (object)["degree" => ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($dictEuclideanRing_0)->{'degree'}))(($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force'))), "div" => (function() use ($dictEuclideanRing_0) {
  $__fn = function($a_2 = null, $b_3 = null) use ($dictEuclideanRing_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_4 = null) use ($a_2, $b_3, $dictEuclideanRing_0) {
  $__num = \func_num_args();
  $__res = ((($dictEuclideanRing_0)->{'div'})((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($a_2)))((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($b_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "mod" => (function() use ($dictEuclideanRing_0) {
  $__fn = function($a_2 = null, $b_3 = null) use ($dictEuclideanRing_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_4 = null) use ($a_2, $b_3, $dictEuclideanRing_0) {
  $__num = \func_num_args();
  $__res = ((($dictEuclideanRing_0)->{'mod'})((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($a_2)))((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($b_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "CommutativeRing0" => function($dollar__unused_2 = null) use ($ringLazy1_1_0) {
  $__num = \func_num_args();
  $__res = (object)["Ring0" => function($dollar__unused_3 = null) use ($ringLazy1_1_0) {
  $__num = \func_num_args();
  $__res = $ringLazy1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Lazy_boundedLazy'] = function() { $v = function($dictBounded_0 = null) {
  $__num = \func_num_args();
  $top_1_0 = ($dictBounded_0)->{'top'};
  $bottom_2_1 = ($dictBounded_0)->{'bottom'};
  $ordLazy1_3_2 = (($GLOBALS['Data_Lazy_ordLazy'] ?? \PhpursThunks::eval('Data_Lazy_ordLazy')))((($dictBounded_0)->{'Ord0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = (object)["top" => (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_4 = null) use ($top_1_0) {
  $__num = \func_num_args();
  $__res = $top_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), "bottom" => (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_4 = null) use ($bottom_2_1) {
  $__num = \func_num_args();
  $__res = $bottom_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), "Ord0" => function($dollar__unused_4 = null) use ($ordLazy1_3_2) {
  $__num = \func_num_args();
  $__res = $ordLazy1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Lazy_applyLazy'] = function() { $v = (object)["apply" => (function() {
  $__fn = function($f_0 = null, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_2 = null) use ($f_0, $x_1) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($f_0))((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($x_1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Lazy_functorLazy'] ?? \PhpursThunks::eval('Data_Lazy_functorLazy'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Lazy_bindLazy'] = function() { $v = (object)["bind" => (function() {
  $__fn = function($l_0 = null, $f_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_2 = null) use ($f_1, $l_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))(($f_1)((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($l_0)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Apply0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Lazy_applyLazy'] ?? \PhpursThunks::eval('Data_Lazy_applyLazy'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Lazy_heytingAlgebraLazy'] = function() { $v = function($dictHeytingAlgebra_0 = null) {
  $__num = \func_num_args();
  $ff_1_0 = ($dictHeytingAlgebra_0)->{'ff'};
  $tt_2_1 = ($dictHeytingAlgebra_0)->{'tt'};
  $__res = (object)["ff" => (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_3 = null) use ($ff_1_0) {
  $__num = \func_num_args();
  $__res = $ff_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), "tt" => (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_3 = null) use ($tt_2_1) {
  $__num = \func_num_args();
  $__res = $tt_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), "implies" => (function() use ($dictHeytingAlgebra_0) {
  $__fn = function($a_3 = null, $b_4 = null) use ($dictHeytingAlgebra_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_5_2 = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_5 = null) use ($a_3, $dictHeytingAlgebra_0) {
  $__num = \func_num_args();
  $__res = (($dictHeytingAlgebra_0)->{'implies'})((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($a_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_6 = null) use ($__local_var_5_2, $b_4) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($__local_var_5_2))((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($b_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "conj" => (function() use ($dictHeytingAlgebra_0) {
  $__fn = function($a_3 = null, $b_4 = null) use ($dictHeytingAlgebra_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_5_3 = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_5 = null) use ($a_3, $dictHeytingAlgebra_0) {
  $__num = \func_num_args();
  $__res = (($dictHeytingAlgebra_0)->{'conj'})((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($a_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_6 = null) use ($__local_var_5_3, $b_4) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($__local_var_5_3))((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($b_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "disj" => (function() use ($dictHeytingAlgebra_0) {
  $__fn = function($a_3 = null, $b_4 = null) use ($dictHeytingAlgebra_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_5_4 = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_5 = null) use ($a_3, $dictHeytingAlgebra_0) {
  $__num = \func_num_args();
  $__res = (($dictHeytingAlgebra_0)->{'disj'})((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($a_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_6 = null) use ($__local_var_5_4, $b_4) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($__local_var_5_4))((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($b_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "not" => function($a_3 = null) use ($dictHeytingAlgebra_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_4 = null) use ($a_3, $dictHeytingAlgebra_0) {
  $__num = \func_num_args();
  $__res = (($dictHeytingAlgebra_0)->{'not'})((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($a_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Lazy_booleanAlgebraLazy'] = function() { $v = function($dictBooleanAlgebra_0 = null) {
  $__num = \func_num_args();
  $heytingAlgebraLazy1_1_0 = (($GLOBALS['Data_Lazy_heytingAlgebraLazy'] ?? \PhpursThunks::eval('Data_Lazy_heytingAlgebraLazy')))((($dictBooleanAlgebra_0)->{'HeytingAlgebra0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = (object)["HeytingAlgebra0" => function($dollar__unused_2 = null) use ($heytingAlgebraLazy1_1_0) {
  $__num = \func_num_args();
  $__res = $heytingAlgebraLazy1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Lazy_applicativeLazy'] = function() { $v = (object)["pure" => function($a_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_1 = null) use ($a_0) {
  $__num = \func_num_args();
  $__res = $a_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Lazy_applyLazy'] ?? \PhpursThunks::eval('Data_Lazy_applyLazy'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Lazy_monadLazy'] = function() { $v = (object)["Applicative0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Lazy_applicativeLazy'] ?? \PhpursThunks::eval('Data_Lazy_applicativeLazy'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Lazy_bindLazy'] ?? \PhpursThunks::eval('Data_Lazy_bindLazy'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };
































