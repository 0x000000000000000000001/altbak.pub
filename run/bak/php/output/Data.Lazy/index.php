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

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };
function majData_majLazy_defer($v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majLazy_defer';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $f = new class { public function __invoke(...$args) { return $this; } };
  return $f($v0);
}
$GLOBALS['Data_Lazy_defer'] = __NAMESPACE__ . '\\majData_majLazy_defer';

function majData_majLazy_force($v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majLazy_force';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $f = new class { public function __invoke(...$args) { return $this; } };
  return $f($v0);
}
$GLOBALS['Data_Lazy_force'] = __NAMESPACE__ . '\\majData_majLazy_force';





// Data_Lazy_showLazy
function majData_majLazy_showmajLazy($dictShow_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majLazy_showmajLazy';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["show" => function($x_1) use ($dictShow_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})("(defer \\_ -> "))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})((($dictShow_0)->{'show'})(\Data\Lazy\majData_majLazy_force($x_1))))(")"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Lazy_showLazy'] = __NAMESPACE__ . '\\majData_majLazy_showmajLazy';

// Data_Lazy_semiringLazy
function majData_majLazy_semiringmajLazy($dictSemiring_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majLazy_semiringmajLazy';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $zero_1_0 = ($dictSemiring_0)->{'zero'};
  $one_2_1 = ($dictSemiring_0)->{'one'};
  $__res = (object)["add" => function($a_3) use ($dictSemiring_0) {
  $__num = \func_num_args();
  $__res = function($b_4) use ($a_3, $dictSemiring_0) {
  $__num = \func_num_args();
  $__res = \Data\Lazy\majData_majLazy_defer(function($v_5) use ($a_3, $b_4, $dictSemiring_0) {
  $__num = \func_num_args();
  $__res = ((($dictSemiring_0)->{'add'})(\Data\Lazy\majData_majLazy_force($a_3)))(\Data\Lazy\majData_majLazy_force($b_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "zero" => \Data\Lazy\majData_majLazy_defer(function($v_3) use ($zero_1_0) {
  $__num = \func_num_args();
  $__res = $zero_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), "mul" => function($a_3) use ($dictSemiring_0) {
  $__num = \func_num_args();
  $__res = function($b_4) use ($a_3, $dictSemiring_0) {
  $__num = \func_num_args();
  $__res = \Data\Lazy\majData_majLazy_defer(function($v_5) use ($a_3, $b_4, $dictSemiring_0) {
  $__num = \func_num_args();
  $__res = ((($dictSemiring_0)->{'mul'})(\Data\Lazy\majData_majLazy_force($a_3)))(\Data\Lazy\majData_majLazy_force($b_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "one" => \Data\Lazy\majData_majLazy_defer(function($v_3) use ($one_2_1) {
  $__num = \func_num_args();
  $__res = $one_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Lazy_semiringLazy'] = __NAMESPACE__ . '\\majData_majLazy_semiringmajLazy';

// Data_Lazy_semigroupLazy
function majData_majLazy_semigroupmajLazy($dictSemigroup_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majLazy_semigroupmajLazy';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["append" => function($a_1) use ($dictSemigroup_0) {
  $__num = \func_num_args();
  $__res = function($b_2) use ($a_1, $dictSemigroup_0) {
  $__num = \func_num_args();
  $__res = \Data\Lazy\majData_majLazy_defer(function($v_3) use ($a_1, $b_2, $dictSemigroup_0) {
  $__num = \func_num_args();
  $__res = ((($dictSemigroup_0)->{'append'})(\Data\Lazy\majData_majLazy_force($a_1)))(\Data\Lazy\majData_majLazy_force($b_2));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Lazy_semigroupLazy'] = __NAMESPACE__ . '\\majData_majLazy_semigroupmajLazy';

// Data_Lazy_ringLazy
function majData_majLazy_ringmajLazy($dictRing_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majLazy_ringmajLazy';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $semiringLazy1_1_0 = ($GLOBALS['Data_Lazy_semiringLazy'])((($dictRing_0)->{'Semiring0'})(null));
  $__res = (object)["sub" => function($a_2) use ($dictRing_0) {
  $__num = \func_num_args();
  $__res = function($b_3) use ($a_2, $dictRing_0) {
  $__num = \func_num_args();
  $__res = \Data\Lazy\majData_majLazy_defer(function($v_4) use ($a_2, $b_3, $dictRing_0) {
  $__num = \func_num_args();
  $__res = ((($dictRing_0)->{'sub'})(\Data\Lazy\majData_majLazy_force($a_2)))(\Data\Lazy\majData_majLazy_force($b_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Semiring0" => function($_dollar__unused_2) use ($semiringLazy1_1_0) {
  $__num = \func_num_args();
  $__res = $semiringLazy1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Lazy_ringLazy'] = __NAMESPACE__ . '\\majData_majLazy_ringmajLazy';

// Data_Lazy_monoidLazy
function majData_majLazy_monoidmajLazy($dictMonoid_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majLazy_monoidmajLazy';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $mempty_1_0 = ($dictMonoid_0)->{'mempty'};
  $semigroupLazy1_2_1 = ($GLOBALS['Data_Lazy_semigroupLazy'])((($dictMonoid_0)->{'Semigroup0'})(null));
  $__res = (object)["mempty" => \Data\Lazy\majData_majLazy_defer(function($v_3) use ($mempty_1_0) {
  $__num = \func_num_args();
  $__res = $mempty_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), "Semigroup0" => function($_dollar__unused_3) use ($semigroupLazy1_2_1) {
  $__num = \func_num_args();
  $__res = $semigroupLazy1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Lazy_monoidLazy'] = __NAMESPACE__ . '\\majData_majLazy_monoidmajLazy';

// Data_Lazy_lazyLazy
$GLOBALS['Data_Lazy_lazyLazy'] = (object)["defer" => function($f_0) {
  $__num = \func_num_args();
  $__res = \Data\Lazy\majData_majLazy_defer(function($v_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = \Data\Lazy\majData_majLazy_force(($f_0)($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Lazy_functorLazy
$GLOBALS['Data_Lazy_functorLazy'] = (object)["map" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($l_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = \Data\Lazy\majData_majLazy_defer(function($v_2) use ($f_0, $l_1) {
  $__num = \func_num_args();
  $__res = ($f_0)(\Data\Lazy\majData_majLazy_force($l_1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Lazy_functorWithIndexLazy
$GLOBALS['Data_Lazy_functorWithIndexLazy'] = (object)["mapWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Lazy_functorLazy'])->{'map'})(($f_0)($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Lazy_functorLazy'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Lazy_invariantLazy
$GLOBALS['Data_Lazy_invariantLazy'] = (object)["imap" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($v_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Lazy_functorLazy'])->{'map'})($f_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Lazy_foldableLazy
$GLOBALS['Data_Lazy_foldableLazy'] = (object)["foldr" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($z_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = function($l_2) use ($f_0, $z_1) {
  $__num = \func_num_args();
  $__res = (($f_0)(\Data\Lazy\majData_majLazy_force($l_2)))($z_1);
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
}, "foldl" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($z_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = function($l_2) use ($f_0, $z_1) {
  $__num = \func_num_args();
  $__res = (($f_0)($z_1))(\Data\Lazy\majData_majLazy_force($l_2));
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
}, "foldMap" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $__res = function($f_1) {
  $__num = \func_num_args();
  $__res = function($l_2) use ($f_1) {
  $__num = \func_num_args();
  $__res = ($f_1)(\Data\Lazy\majData_majLazy_force($l_2));
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
}];

// Data_Lazy_foldableWithIndexLazy
$GLOBALS['Data_Lazy_foldableWithIndexLazy'] = (object)["foldrWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Lazy_foldableLazy'])->{'foldr'})(($f_0)($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldlWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Lazy_foldableLazy'])->{'foldl'})(($f_0)($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldMapWithIndex" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $foldMap1_1_0 = (($GLOBALS['Data_Lazy_foldableLazy'])->{'foldMap'})($dictMonoid_0);
  $__res = function($f_2) use ($foldMap1_1_0) {
  $__num = \func_num_args();
  $__res = ($foldMap1_1_0)(($f_2)($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Lazy_foldableLazy'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Lazy_traversableLazy
$GLOBALS['Data_Lazy_traversableLazy'] = (object)["traverse" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $__res = function($f_1) use ($dictApplicative_0) {
  $__num = \func_num_args();
  $__res = function($l_2) use ($dictApplicative_0, $f_1) {
  $__num = \func_num_args();
  $__res = ((((((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Lazy_defer']))($GLOBALS['Data_Function_const'])))(($f_1)(\Data\Lazy\majData_majLazy_force($l_2)));
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
}, "sequence" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $__res = function($l_1) use ($dictApplicative_0) {
  $__num = \func_num_args();
  $__res = ((((((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Lazy_defer']))($GLOBALS['Data_Function_const'])))(\Data\Lazy\majData_majLazy_force($l_1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Lazy_functorLazy'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Lazy_foldableLazy'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Lazy_traversableWithIndexLazy
$GLOBALS['Data_Lazy_traversableWithIndexLazy'] = (object)["traverseWithIndex" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $traverse1_1_0 = (($GLOBALS['Data_Lazy_traversableLazy'])->{'traverse'})($dictApplicative_0);
  $__res = function($f_2) use ($traverse1_1_0) {
  $__num = \func_num_args();
  $__res = ($traverse1_1_0)(($f_2)($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FunctorWithIndex0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Lazy_functorWithIndexLazy'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Lazy_foldableWithIndexLazy'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Lazy_traversableLazy'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Lazy_foldable1Lazy
$GLOBALS['Data_Lazy_foldable1Lazy'] = (object)["foldMap1" => function($dictSemigroup_0) {
  $__num = \func_num_args();
  $__res = function($f_1) {
  $__num = \func_num_args();
  $__res = function($l_2) use ($f_1) {
  $__num = \func_num_args();
  $__res = ($f_1)(\Data\Lazy\majData_majLazy_force($l_2));
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
}, "foldr1" => function($v_0) {
  $__num = \func_num_args();
  $__res = function($l_1) {
  $__num = \func_num_args();
  $__res = \Data\Lazy\majData_majLazy_force($l_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldl1" => function($v_0) {
  $__num = \func_num_args();
  $__res = function($l_1) {
  $__num = \func_num_args();
  $__res = \Data\Lazy\majData_majLazy_force($l_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Lazy_foldableLazy'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Lazy_traversable1Lazy
$GLOBALS['Data_Lazy_traversable1Lazy'] = (object)["traverse1" => function($dictApply_0) {
  $__num = \func_num_args();
  $__res = function($f_1) use ($dictApply_0) {
  $__num = \func_num_args();
  $__res = function($l_2) use ($dictApply_0, $f_1) {
  $__num = \func_num_args();
  $__res = ((((($dictApply_0)->{'Functor0'})(null))->{'map'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Lazy_defer']))($GLOBALS['Data_Function_const'])))(($f_1)(\Data\Lazy\majData_majLazy_force($l_2)));
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
}, "sequence1" => function($dictApply_0) {
  $__num = \func_num_args();
  $__res = function($l_1) use ($dictApply_0) {
  $__num = \func_num_args();
  $__res = ((((($dictApply_0)->{'Functor0'})(null))->{'map'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Lazy_defer']))($GLOBALS['Data_Function_const'])))(\Data\Lazy\majData_majLazy_force($l_1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable10" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Lazy_foldable1Lazy'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Lazy_traversableLazy'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Lazy_extendLazy
$GLOBALS['Data_Lazy_extendLazy'] = (object)["extend" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($x_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = \Data\Lazy\majData_majLazy_defer(function($v_2) use ($f_0, $x_1) {
  $__num = \func_num_args();
  $__res = ($f_0)($x_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Lazy_functorLazy'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Lazy_eqLazy
function majData_majLazy_eqmajLazy($dictEq_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majLazy_eqmajLazy';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["eq" => function($x_1) use ($dictEq_0) {
  $__num = \func_num_args();
  $__res = function($y_2) use ($dictEq_0, $x_1) {
  $__num = \func_num_args();
  $__res = ((($dictEq_0)->{'eq'})(\Data\Lazy\majData_majLazy_force($x_1)))(\Data\Lazy\majData_majLazy_force($y_2));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Lazy_eqLazy'] = __NAMESPACE__ . '\\majData_majLazy_eqmajLazy';

// Data_Lazy_ordLazy
function majData_majLazy_ordmajLazy($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majLazy_ordmajLazy';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $eqLazy1_1_0 = ($GLOBALS['Data_Lazy_eqLazy'])((($dictOrd_0)->{'Eq0'})(null));
  $__res = (object)["compare" => function($x_2) use ($dictOrd_0) {
  $__num = \func_num_args();
  $__res = function($y_3) use ($dictOrd_0, $x_2) {
  $__num = \func_num_args();
  $__res = ((($dictOrd_0)->{'compare'})(\Data\Lazy\majData_majLazy_force($x_2)))(\Data\Lazy\majData_majLazy_force($y_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq0" => function($_dollar__unused_2) use ($eqLazy1_1_0) {
  $__num = \func_num_args();
  $__res = $eqLazy1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Lazy_ordLazy'] = __NAMESPACE__ . '\\majData_majLazy_ordmajLazy';

// Data_Lazy_eq1Lazy
$GLOBALS['Data_Lazy_eq1Lazy'] = (object)["eq1" => function($dictEq_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Lazy_eqLazy'])($dictEq_0))->{'eq'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Lazy_ord1Lazy
$GLOBALS['Data_Lazy_ord1Lazy'] = (object)["compare1" => function($dictOrd_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Lazy_ordLazy'])($dictOrd_0))->{'compare'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq10" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Lazy_eq1Lazy'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Lazy_comonadLazy
$GLOBALS['Data_Lazy_comonadLazy'] = (object)["extract" => $GLOBALS['Data_Lazy_force'], "Extend0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Lazy_extendLazy'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Lazy_commutativeRingLazy
function majData_majLazy_commutativemajRingmajLazy($dictCommutativeRing_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majLazy_commutativemajRingmajLazy';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $ringLazy1_1_0 = ($GLOBALS['Data_Lazy_ringLazy'])((($dictCommutativeRing_0)->{'Ring0'})(null));
  $__res = (object)["Ring0" => function($_dollar__unused_2) use ($ringLazy1_1_0) {
  $__num = \func_num_args();
  $__res = $ringLazy1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Lazy_commutativeRingLazy'] = __NAMESPACE__ . '\\majData_majLazy_commutativemajRingmajLazy';

// Data_Lazy_euclideanRingLazy
function majData_majLazy_euclideanmajRingmajLazy($dictEuclideanRing_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majLazy_euclideanmajRingmajLazy';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $ringLazy1_1_0 = ($GLOBALS['Data_Lazy_ringLazy'])((((($dictEuclideanRing_0)->{'CommutativeRing0'})(null))->{'Ring0'})(null));
  $commutativeRingLazy1_1_0 = (object)["Ring0" => function($_dollar__unused_2) use ($ringLazy1_1_0) {
  $__num = \func_num_args();
  $__res = $ringLazy1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["degree" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(($dictEuclideanRing_0)->{'degree'}))($GLOBALS['Data_Lazy_force']), "div" => function($a_2) use ($dictEuclideanRing_0) {
  $__num = \func_num_args();
  $__res = function($b_3) use ($a_2, $dictEuclideanRing_0) {
  $__num = \func_num_args();
  $__res = \Data\Lazy\majData_majLazy_defer(function($v_4) use ($a_2, $b_3, $dictEuclideanRing_0) {
  $__num = \func_num_args();
  $__res = ((($dictEuclideanRing_0)->{'div'})(\Data\Lazy\majData_majLazy_force($a_2)))(\Data\Lazy\majData_majLazy_force($b_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "mod" => function($a_2) use ($dictEuclideanRing_0) {
  $__num = \func_num_args();
  $__res = function($b_3) use ($a_2, $dictEuclideanRing_0) {
  $__num = \func_num_args();
  $__res = \Data\Lazy\majData_majLazy_defer(function($v_4) use ($a_2, $b_3, $dictEuclideanRing_0) {
  $__num = \func_num_args();
  $__res = ((($dictEuclideanRing_0)->{'mod'})(\Data\Lazy\majData_majLazy_force($a_2)))(\Data\Lazy\majData_majLazy_force($b_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "CommutativeRing0" => function($_dollar__unused_2) use ($commutativeRingLazy1_1_0) {
  $__num = \func_num_args();
  $__res = $commutativeRingLazy1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Lazy_euclideanRingLazy'] = __NAMESPACE__ . '\\majData_majLazy_euclideanmajRingmajLazy';

// Data_Lazy_boundedLazy
function majData_majLazy_boundedmajLazy($dictBounded_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majLazy_boundedmajLazy';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $top_1_0 = ($dictBounded_0)->{'top'};
  $bottom_2_1 = ($dictBounded_0)->{'bottom'};
  $ordLazy1_3_2 = ($GLOBALS['Data_Lazy_ordLazy'])((($dictBounded_0)->{'Ord0'})(null));
  $__res = (object)["top" => \Data\Lazy\majData_majLazy_defer(function($v_4) use ($top_1_0) {
  $__num = \func_num_args();
  $__res = $top_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), "bottom" => \Data\Lazy\majData_majLazy_defer(function($v_4) use ($bottom_2_1) {
  $__num = \func_num_args();
  $__res = $bottom_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), "Ord0" => function($_dollar__unused_4) use ($ordLazy1_3_2) {
  $__num = \func_num_args();
  $__res = $ordLazy1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Lazy_boundedLazy'] = __NAMESPACE__ . '\\majData_majLazy_boundedmajLazy';

// Data_Lazy_applyLazy
$GLOBALS['Data_Lazy_applyLazy'] = (object)["apply" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($x_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = \Data\Lazy\majData_majLazy_defer(function($v_2) use ($f_0, $x_1) {
  $__num = \func_num_args();
  $__res = (\Data\Lazy\majData_majLazy_force($f_0))(\Data\Lazy\majData_majLazy_force($x_1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Lazy_functorLazy'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Lazy_bindLazy
$GLOBALS['Data_Lazy_bindLazy'] = (object)["bind" => function($l_0) {
  $__num = \func_num_args();
  $__res = function($f_1) use ($l_0) {
  $__num = \func_num_args();
  $__res = \Data\Lazy\majData_majLazy_defer(function($v_2) use ($f_1, $l_0) {
  $__num = \func_num_args();
  $__res = \Data\Lazy\majData_majLazy_force(($f_1)(\Data\Lazy\majData_majLazy_force($l_0)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Lazy_applyLazy'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Lazy_heytingAlgebraLazy
function majData_majLazy_heytingmajAlgebramajLazy($dictHeytingAlgebra_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majLazy_heytingmajAlgebramajLazy';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $ff_1_0 = ($dictHeytingAlgebra_0)->{'ff'};
  $tt_2_1 = ($dictHeytingAlgebra_0)->{'tt'};
  $implies_3_2 = ($dictHeytingAlgebra_0)->{'implies'};
  $conj_4_3 = ($dictHeytingAlgebra_0)->{'conj'};
  $disj_5_4 = ($dictHeytingAlgebra_0)->{'disj'};
  $not_6_5 = ($dictHeytingAlgebra_0)->{'not'};
  $__res = (object)["ff" => \Data\Lazy\majData_majLazy_defer(function($v_7) use ($ff_1_0) {
  $__num = \func_num_args();
  $__res = $ff_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), "tt" => \Data\Lazy\majData_majLazy_defer(function($v_7) use ($tt_2_1) {
  $__num = \func_num_args();
  $__res = $tt_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), "implies" => function($a_7) use ($implies_3_2) {
  $__num = \func_num_args();
  $__res = function($b_8) use ($a_7, $implies_3_2) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Lazy_applyLazy'])->{'apply'})(((($GLOBALS['Data_Lazy_functorLazy'])->{'map'})($implies_3_2))($a_7)))($b_8);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "conj" => function($a_7) use ($conj_4_3) {
  $__num = \func_num_args();
  $__res = function($b_8) use ($a_7, $conj_4_3) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Lazy_applyLazy'])->{'apply'})(((($GLOBALS['Data_Lazy_functorLazy'])->{'map'})($conj_4_3))($a_7)))($b_8);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "disj" => function($a_7) use ($disj_5_4) {
  $__num = \func_num_args();
  $__res = function($b_8) use ($a_7, $disj_5_4) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Lazy_applyLazy'])->{'apply'})(((($GLOBALS['Data_Lazy_functorLazy'])->{'map'})($disj_5_4))($a_7)))($b_8);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "not" => function($a_7) use ($not_6_5) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Lazy_functorLazy'])->{'map'})($not_6_5))($a_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Lazy_heytingAlgebraLazy'] = __NAMESPACE__ . '\\majData_majLazy_heytingmajAlgebramajLazy';

// Data_Lazy_booleanAlgebraLazy
function majData_majLazy_booleanmajAlgebramajLazy($dictBooleanAlgebra_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majLazy_booleanmajAlgebramajLazy';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $heytingAlgebraLazy1_1_0 = ($GLOBALS['Data_Lazy_heytingAlgebraLazy'])((($dictBooleanAlgebra_0)->{'HeytingAlgebra0'})(null));
  $__res = (object)["HeytingAlgebra0" => function($_dollar__unused_2) use ($heytingAlgebraLazy1_1_0) {
  $__num = \func_num_args();
  $__res = $heytingAlgebraLazy1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Lazy_booleanAlgebraLazy'] = __NAMESPACE__ . '\\majData_majLazy_booleanmajAlgebramajLazy';

// Data_Lazy_applicativeLazy
$GLOBALS['Data_Lazy_applicativeLazy'] = (object)["pure" => function($a_0) {
  $__num = \func_num_args();
  $__res = \Data\Lazy\majData_majLazy_defer(function($v_1) use ($a_0) {
  $__num = \func_num_args();
  $__res = $a_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Lazy_applyLazy'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Lazy_monadLazy
$GLOBALS['Data_Lazy_monadLazy'] = (object)["Applicative0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Lazy_applicativeLazy'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Lazy_bindLazy'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

