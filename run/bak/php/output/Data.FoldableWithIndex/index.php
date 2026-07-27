<?php

namespace Data\FoldableWithIndex;

// ALL IMPORTS: Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Semigroupoid, Data.Const, Data.Either, Data.Foldable, Data.FoldableWithIndex, Data.Function, Data.Functor.App, Data.Functor.Compose, Data.Functor.Coproduct, Data.Functor.Product, Data.FunctorWithIndex, Data.Identity, Data.Maybe, Data.Maybe.First, Data.Maybe.Last, Data.Monoid, Data.Monoid.Additive, Data.Monoid.Conj, Data.Monoid.Disj, Data.Monoid.Dual, Data.Monoid.Endo, Data.Monoid.Multiplicative, Data.Newtype, Data.Semigroup, Data.Tuple, Data.Unit, Prelude, Prim
// TO REQUIRE: Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Semigroupoid, Data.Const, Data.Either, Data.Foldable, Data.FoldableWithIndex, Data.Function, Data.Functor.App, Data.Functor.Compose, Data.Functor.Coproduct, Data.Functor.Product, Data.FunctorWithIndex, Data.Identity, Data.Maybe, Data.Maybe.First, Data.Maybe.Last, Data.Monoid, Data.Monoid.Additive, Data.Monoid.Conj, Data.Monoid.Disj, Data.Monoid.Dual, Data.Monoid.Endo, Data.Monoid.Multiplicative, Data.Newtype, Data.Semigroup, Data.Tuple, Data.Unit, Prelude
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Category/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Const/index.php';
require_once __DIR__ . '/../Data.Either/index.php';
require_once __DIR__ . '/../Data.Foldable/index.php';
require_once __DIR__ . '/../Data.FoldableWithIndex/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor.App/index.php';
require_once __DIR__ . '/../Data.Functor.Compose/index.php';
require_once __DIR__ . '/../Data.Functor.Coproduct/index.php';
require_once __DIR__ . '/../Data.Functor.Product/index.php';
require_once __DIR__ . '/../Data.FunctorWithIndex/index.php';
require_once __DIR__ . '/../Data.Identity/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Maybe.First/index.php';
require_once __DIR__ . '/../Data.Maybe.Last/index.php';
require_once __DIR__ . '/../Data.Monoid/index.php';
require_once __DIR__ . '/../Data.Monoid.Additive/index.php';
require_once __DIR__ . '/../Data.Monoid.Conj/index.php';
require_once __DIR__ . '/../Data.Monoid.Disj/index.php';
require_once __DIR__ . '/../Data.Monoid.Dual/index.php';
require_once __DIR__ . '/../Data.Monoid.Endo/index.php';
require_once __DIR__ . '/../Data.Monoid.Multiplicative/index.php';
require_once __DIR__ . '/../Data.Newtype/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
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




// Data_FoldableWithIndex_monoidEndo
$GLOBALS['Data_FoldableWithIndex_monoidEndo'] = (function() use (&$__fn) {
$__local_var_0_0 = (($GLOBALS['Control_Category_categoryFn'])->{'Semigroupoid0'})(null);
$semigroupEndo1_1_1 = (object)["append" => (function() use ($__local_var_0_0) {
  $__fn = function($v_1, $v1_2 = null) use ($__local_var_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_0_0)->{'compose'})($v_1))($v1_2);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
return (object)["mempty" => function($x_2) {
  $__num = \func_num_args();
  $__res = $x_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Semigroup0" => function($_dollar__unused_2) use ($semigroupEndo1_1_1) {
  $__num = \func_num_args();
  $__res = $semigroupEndo1_1_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
})();

// Data_FoldableWithIndex_monoidDual
$GLOBALS['Data_FoldableWithIndex_monoidDual'] = (function() use (&$__fn) {
$__local_var_0_0 = (($GLOBALS['Data_FoldableWithIndex_monoidEndo'])->{'Semigroup0'})(null);
$semigroupDual1_1_1 = (object)["append" => (function() use ($__local_var_0_0) {
  $__fn = function($v_1, $v1_2 = null) use ($__local_var_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_0_0)->{'append'})($v1_2))($v_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
return (object)["mempty" => ($GLOBALS['Data_FoldableWithIndex_monoidEndo'])->{'mempty'}, "Semigroup0" => function($_dollar__unused_2) use ($semigroupDual1_1_1) {
  $__num = \func_num_args();
  $__res = $semigroupDual1_1_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
})();

// Data_FoldableWithIndex_foldrWithIndex
function majData_majFoldablemajWithmajIndex_foldrmajWithmajIndex($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldablemajWithmajIndex_foldrmajWithmajIndex';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'foldrWithIndex'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_FoldableWithIndex_foldrWithIndex'] = __NAMESPACE__ . '\\majData_majFoldablemajWithmajIndex_foldrmajWithmajIndex';

// Data_FoldableWithIndex_traverseWithIndex_
function majData_majFoldablemajWithmajIndex_traversemajWithmajIndex_($dictApplicative_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldablemajWithmajIndex_traversemajWithmajIndex_';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictApplicative_0)->{'Apply0'})(null);
  $__res = (function() use ($__local_var_1_0, $dictApplicative_0) {
  $__fn = function($dictFoldableWithIndex_2, $f_3 = null) use ($__local_var_1_0, $dictApplicative_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictFoldableWithIndex_2)->{'foldrWithIndex'})(function($i_4) use ($__local_var_1_0, $f_3) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((function() use ($__local_var_1_0) {
  $__fn = function($a_5, $b_6 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_1_0)->{'apply'})(((((($__local_var_1_0)->{'Functor0'})(null))->{'map'})(function($v_7) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Control_Apply_identity'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($a_5)))($b_6);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(($f_3)($i_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($dictApplicative_0)->{'pure'})($GLOBALS['Data_Unit_unit']));
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
$GLOBALS['Data_FoldableWithIndex_traverseWithIndex_'] = __NAMESPACE__ . '\\majData_majFoldablemajWithmajIndex_traversemajWithmajIndex_';

// Data_FoldableWithIndex_forWithIndex_
function majData_majFoldablemajWithmajIndex_formajWithmajIndex_($dictApplicative_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldablemajWithmajIndex_formajWithmajIndex_';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $traverseWithIndex_1_1_0 = ($GLOBALS['Data_FoldableWithIndex_traverseWithIndex_'])($dictApplicative_0);
  $__res = function($dictFoldableWithIndex_2) use ($traverseWithIndex_1_1_0) {
  $__num = \func_num_args();
  $__local_var_3_1 = ($traverseWithIndex_1_1_0)($dictFoldableWithIndex_2);
  $__res = (function() use ($__local_var_3_1) {
  $__fn = function($b_4, $a_5 = null) use ($__local_var_3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($__local_var_3_1)($a_5))($b_4);
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
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_FoldableWithIndex_forWithIndex_'] = __NAMESPACE__ . '\\majData_majFoldablemajWithmajIndex_formajWithmajIndex_';

// Data_FoldableWithIndex_foldrDefault
function majData_majFoldablemajWithmajIndex_foldrmajDefault($dictFoldableWithIndex_0, $f_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldablemajWithmajIndex_foldrmajDefault';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($dictFoldableWithIndex_0)->{'foldrWithIndex'})(function($v_2) use ($f_1) {
  $__num = \func_num_args();
  $__res = $f_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_FoldableWithIndex_foldrDefault'] = __NAMESPACE__ . '\\majData_majFoldablemajWithmajIndex_foldrmajDefault';

// Data_FoldableWithIndex_foldlWithIndex
function majData_majFoldablemajWithmajIndex_foldlmajWithmajIndex($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldablemajWithmajIndex_foldlmajWithmajIndex';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'foldlWithIndex'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_FoldableWithIndex_foldlWithIndex'] = __NAMESPACE__ . '\\majData_majFoldablemajWithmajIndex_foldlmajWithmajIndex';

// Data_FoldableWithIndex_foldlDefault
function majData_majFoldablemajWithmajIndex_foldlmajDefault($dictFoldableWithIndex_0, $f_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldablemajWithmajIndex_foldlmajDefault';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($dictFoldableWithIndex_0)->{'foldlWithIndex'})(function($v_2) use ($f_1) {
  $__num = \func_num_args();
  $__res = $f_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_FoldableWithIndex_foldlDefault'] = __NAMESPACE__ . '\\majData_majFoldablemajWithmajIndex_foldlmajDefault';

// Data_FoldableWithIndex_foldableWithIndexTuple
$GLOBALS['Data_FoldableWithIndex_foldableWithIndexTuple'] = (object)["foldrWithIndex" => (function() {
  $__fn = function($f_0, $z_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($f_0)($GLOBALS['Data_Unit_unit']))(($v_2)->{'value1'}))($z_1);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldlWithIndex" => (function() {
  $__fn = function($f_0, $z_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($f_0)($GLOBALS['Data_Unit_unit']))($z_1))(($v_2)->{'value1'});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldMapWithIndex" => (function() {
  $__fn = function($dictMonoid_0, $f_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($f_1)($GLOBALS['Data_Unit_unit']))(($v_2)->{'value1'});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "Foldable0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableTuple'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_FoldableWithIndex_foldableWithIndexMultiplicative
$GLOBALS['Data_FoldableWithIndex_foldableWithIndexMultiplicative'] = (object)["foldrWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Foldable_foldableMultiplicative'])->{'foldr'})(($f_0)($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldlWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Foldable_foldableMultiplicative'])->{'foldl'})(($f_0)($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldMapWithIndex" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $foldMap8_1_0 = (($GLOBALS['Data_Foldable_foldableMultiplicative'])->{'foldMap'})($dictMonoid_0);
  $__res = function($f_2) use ($foldMap8_1_0) {
  $__num = \func_num_args();
  $__res = ($foldMap8_1_0)(($f_2)($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableMultiplicative'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_FoldableWithIndex_foldableWithIndexMaybe
$GLOBALS['Data_FoldableWithIndex_foldableWithIndexMaybe'] = (object)["foldrWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Foldable_foldableMaybe'])->{'foldr'})(($f_0)($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldlWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Foldable_foldableMaybe'])->{'foldl'})(($f_0)($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldMapWithIndex" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $foldMap8_1_0 = (($GLOBALS['Data_Foldable_foldableMaybe'])->{'foldMap'})($dictMonoid_0);
  $__res = function($f_2) use ($foldMap8_1_0) {
  $__num = \func_num_args();
  $__res = ($foldMap8_1_0)(($f_2)($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableMaybe'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_FoldableWithIndex_foldableWithIndexLast
$GLOBALS['Data_FoldableWithIndex_foldableWithIndexLast'] = (object)["foldrWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Foldable_foldableLast'])->{'foldr'})(($f_0)($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldlWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Foldable_foldableLast'])->{'foldl'})(($f_0)($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldMapWithIndex" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $foldMap8_1_0 = (($GLOBALS['Data_Foldable_foldableLast'])->{'foldMap'})($dictMonoid_0);
  $__res = function($f_2) use ($foldMap8_1_0) {
  $__num = \func_num_args();
  $__res = ($foldMap8_1_0)(($f_2)($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableLast'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_FoldableWithIndex_foldableWithIndexIdentity
$GLOBALS['Data_FoldableWithIndex_foldableWithIndexIdentity'] = (object)["foldrWithIndex" => (function() {
  $__fn = function($f_0, $z_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($f_0)($GLOBALS['Data_Unit_unit']))($v_2))($z_1);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldlWithIndex" => (function() {
  $__fn = function($f_0, $z_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($f_0)($GLOBALS['Data_Unit_unit']))($z_1))($v_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldMapWithIndex" => (function() {
  $__fn = function($dictMonoid_0, $f_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($f_1)($GLOBALS['Data_Unit_unit']))($v_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "Foldable0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableIdentity'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_FoldableWithIndex_foldableWithIndexFirst
$GLOBALS['Data_FoldableWithIndex_foldableWithIndexFirst'] = (object)["foldrWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Foldable_foldableFirst'])->{'foldr'})(($f_0)($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldlWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Foldable_foldableFirst'])->{'foldl'})(($f_0)($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldMapWithIndex" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $foldMap8_1_0 = (($GLOBALS['Data_Foldable_foldableFirst'])->{'foldMap'})($dictMonoid_0);
  $__res = function($f_2) use ($foldMap8_1_0) {
  $__num = \func_num_args();
  $__res = ($foldMap8_1_0)(($f_2)($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableFirst'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_FoldableWithIndex_foldableWithIndexEither
$GLOBALS['Data_FoldableWithIndex_foldableWithIndexEither'] = (object)["foldrWithIndex" => (function() {
  $__fn = function($v_0, $v1_1 = null, $v2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if ($v2_2 instanceof \Data\Either\Data_Either_Left) {
$__t0 = $v1_1;
goto end_branch_0;;
};
  if ($v2_2 instanceof \Data\Either\Data_Either_Right) {
$__t0 = ((($v_0)($GLOBALS['Data_Unit_unit']))(($v2_2)->{'value0'}))($v1_1);
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
})(), "foldlWithIndex" => (function() {
  $__fn = function($v_0, $v1_1 = null, $v2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t1 = null;;
  if ($v2_2 instanceof \Data\Either\Data_Either_Left) {
$__t1 = $v1_1;
goto end_branch_1;;
};
  if ($v2_2 instanceof \Data\Either\Data_Either_Right) {
$__t1 = ((($v_0)($GLOBALS['Data_Unit_unit']))($v1_1))(($v2_2)->{'value0'});
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldMapWithIndex" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $mempty_1_2 = ($dictMonoid_0)->{'mempty'};
  $__res = (function() use ($mempty_1_2) {
  $__fn = function($v_2, $v1_3 = null) use ($mempty_1_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t3 = null;;
  if ($v1_3 instanceof \Data\Either\Data_Either_Left) {
$__t3 = $mempty_1_2;
goto end_branch_3;;
};
  if ($v1_3 instanceof \Data\Either\Data_Either_Right) {
$__t3 = (($v_2)($GLOBALS['Data_Unit_unit']))(($v1_3)->{'value0'});
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
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableEither'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_FoldableWithIndex_foldableWithIndexDual
$GLOBALS['Data_FoldableWithIndex_foldableWithIndexDual'] = (object)["foldrWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Foldable_foldableDual'])->{'foldr'})(($f_0)($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldlWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Foldable_foldableDual'])->{'foldl'})(($f_0)($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldMapWithIndex" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $foldMap8_1_0 = (($GLOBALS['Data_Foldable_foldableDual'])->{'foldMap'})($dictMonoid_0);
  $__res = function($f_2) use ($foldMap8_1_0) {
  $__num = \func_num_args();
  $__res = ($foldMap8_1_0)(($f_2)($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableDual'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_FoldableWithIndex_foldableWithIndexDisj
$GLOBALS['Data_FoldableWithIndex_foldableWithIndexDisj'] = (object)["foldrWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Foldable_foldableDisj'])->{'foldr'})(($f_0)($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldlWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Foldable_foldableDisj'])->{'foldl'})(($f_0)($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldMapWithIndex" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $foldMap8_1_0 = (($GLOBALS['Data_Foldable_foldableDisj'])->{'foldMap'})($dictMonoid_0);
  $__res = function($f_2) use ($foldMap8_1_0) {
  $__num = \func_num_args();
  $__res = ($foldMap8_1_0)(($f_2)($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableDisj'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_FoldableWithIndex_foldableWithIndexConst
$GLOBALS['Data_FoldableWithIndex_foldableWithIndexConst'] = (object)["foldrWithIndex" => (function() {
  $__fn = function($v_0, $z_1 = null, $v1_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = $z_1;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldlWithIndex" => (function() {
  $__fn = function($v_0, $z_1 = null, $v1_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = $z_1;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldMapWithIndex" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $mempty_1_0 = ($dictMonoid_0)->{'mempty'};
  $__res = (function() use ($mempty_1_0) {
  $__fn = function($v_2, $v1_3 = null) use ($mempty_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = $mempty_1_0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableConst'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_FoldableWithIndex_foldableWithIndexConj
$GLOBALS['Data_FoldableWithIndex_foldableWithIndexConj'] = (object)["foldrWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Foldable_foldableConj'])->{'foldr'})(($f_0)($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldlWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Foldable_foldableConj'])->{'foldl'})(($f_0)($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldMapWithIndex" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $foldMap8_1_0 = (($GLOBALS['Data_Foldable_foldableConj'])->{'foldMap'})($dictMonoid_0);
  $__res = function($f_2) use ($foldMap8_1_0) {
  $__num = \func_num_args();
  $__res = ($foldMap8_1_0)(($f_2)($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableConj'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_FoldableWithIndex_foldableWithIndexAdditive
$GLOBALS['Data_FoldableWithIndex_foldableWithIndexAdditive'] = (object)["foldrWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Foldable_foldableAdditive'])->{'foldr'})(($f_0)($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldlWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Foldable_foldableAdditive'])->{'foldl'})(($f_0)($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldMapWithIndex" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $foldMap8_1_0 = (($GLOBALS['Data_Foldable_foldableAdditive'])->{'foldMap'})($dictMonoid_0);
  $__res = function($f_2) use ($foldMap8_1_0) {
  $__num = \func_num_args();
  $__res = ($foldMap8_1_0)(($f_2)($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableAdditive'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_FoldableWithIndex_foldWithIndexM
function majData_majFoldablemajWithmajIndex_foldmajWithmajIndexmajM($dictFoldableWithIndex_0, $dictMonad_1 = null, $f_2 = null, $a0_3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldablemajWithmajIndex_foldmajWithmajIndexmajM';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = ((($dictFoldableWithIndex_0)->{'foldlWithIndex'})((function() use ($dictMonad_1, $f_2) {
  $__fn = function($i_4, $ma_5 = null, $b_6 = null) use ($dictMonad_1, $f_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__local_var_7_0 = ($f_2)($i_4);
  $__res = ((((($dictMonad_1)->{'Bind1'})(null))->{'bind'})($ma_5))(function($a_8) use ($__local_var_7_0, $b_6) {
  $__num = \func_num_args();
  $__res = (($__local_var_7_0)($a_8))($b_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()))((((($dictMonad_1)->{'Applicative0'})(null))->{'pure'})($a0_3));
  goto __end;;
  __end:
  return 4 < $__num ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
}
$GLOBALS['Data_FoldableWithIndex_foldWithIndexM'] = __NAMESPACE__ . '\\majData_majFoldablemajWithmajIndex_foldmajWithmajIndexmajM';

// Data_FoldableWithIndex_foldMapWithIndexDefaultR
function majData_majFoldablemajWithmajIndex_foldmajMapmajWithmajIndexmajDefaultmajR($dictFoldableWithIndex_0, $dictMonoid_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldablemajWithmajIndex_foldmajMapmajWithmajIndexmajDefaultmajR';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $mempty_2_0 = ($dictMonoid_1)->{'mempty'};
  $__res = function($f_3) use ($dictFoldableWithIndex_0, $dictMonoid_1, $mempty_2_0) {
  $__num = \func_num_args();
  $__res = ((($dictFoldableWithIndex_0)->{'foldrWithIndex'})((function() use ($dictMonoid_1, $f_3) {
  $__fn = function($i_4, $x_5 = null, $acc_6 = null) use ($dictMonoid_1, $f_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((($dictMonoid_1)->{'Semigroup0'})(null))->{'append'})((($f_3)($i_4))($x_5)))($acc_6);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()))($mempty_2_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_FoldableWithIndex_foldMapWithIndexDefaultR'] = __NAMESPACE__ . '\\majData_majFoldablemajWithmajIndex_foldmajMapmajWithmajIndexmajDefaultmajR';

// Data_FoldableWithIndex_foldableWithIndexArray
$GLOBALS['Data_FoldableWithIndex_foldableWithIndexArray'] = (object)["foldrWithIndex" => (function() {
  $__fn = function($f_0, $z_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(((($GLOBALS['Data_Foldable_foldableArray'])->{'foldr'})(function($v_2) use ($f_0) {
  $__num = \func_num_args();
  $__local_var_3_0 = ($v_2)->{'value0'};
  $__local_var_4_1 = ($v_2)->{'value1'};
  $__res = function($y_5) use ($__local_var_3_0, $__local_var_4_1, $f_0) {
  $__num = \func_num_args();
  $__res = ((($f_0)($__local_var_3_0))($__local_var_4_1))($y_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($z_1)))((($GLOBALS['Data_FunctorWithIndex_functorWithIndexArray'])->{'mapWithIndex'})($GLOBALS['Data_Tuple_Tuple']));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "foldlWithIndex" => (function() {
  $__fn = function($f_0, $z_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(((($GLOBALS['Data_Foldable_foldableArray'])->{'foldl'})((function() use ($f_0) {
  $__fn = function($y_2, $v_3 = null) use ($f_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($f_0)(($v_3)->{'value0'}))($y_2))(($v_3)->{'value1'});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($z_1)))((($GLOBALS['Data_FunctorWithIndex_functorWithIndexArray'])->{'mapWithIndex'})($GLOBALS['Data_Tuple_Tuple']));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "foldMapWithIndex" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $mempty_1_2 = ($dictMonoid_0)->{'mempty'};
  $__res = function($f_2) use ($dictMonoid_0, $mempty_1_2) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_FoldableWithIndex_foldableWithIndexArray'])->{'foldrWithIndex'})((function() use ($dictMonoid_0, $f_2) {
  $__fn = function($i_3, $x_4 = null, $acc_5 = null) use ($dictMonoid_0, $f_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((($dictMonoid_0)->{'Semigroup0'})(null))->{'append'})((($f_2)($i_3))($x_4)))($acc_5);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()))($mempty_1_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableArray'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_FoldableWithIndex_foldMapWithIndexDefaultL
function majData_majFoldablemajWithmajIndex_foldmajMapmajWithmajIndexmajDefaultmajL($dictFoldableWithIndex_0, $dictMonoid_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldablemajWithmajIndex_foldmajMapmajWithmajIndexmajDefaultmajL';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $mempty_2_0 = ($dictMonoid_1)->{'mempty'};
  $__res = function($f_3) use ($dictFoldableWithIndex_0, $dictMonoid_1, $mempty_2_0) {
  $__num = \func_num_args();
  $__res = ((($dictFoldableWithIndex_0)->{'foldlWithIndex'})((function() use ($dictMonoid_1, $f_3) {
  $__fn = function($i_4, $acc_5 = null, $x_6 = null) use ($dictMonoid_1, $f_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((($dictMonoid_1)->{'Semigroup0'})(null))->{'append'})($acc_5))((($f_3)($i_4))($x_6));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()))($mempty_2_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_FoldableWithIndex_foldMapWithIndexDefaultL'] = __NAMESPACE__ . '\\majData_majFoldablemajWithmajIndex_foldmajMapmajWithmajIndexmajDefaultmajL';

// Data_FoldableWithIndex_foldMapWithIndex
function majData_majFoldablemajWithmajIndex_foldmajMapmajWithmajIndex($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldablemajWithmajIndex_foldmajMapmajWithmajIndex';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'foldMapWithIndex'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_FoldableWithIndex_foldMapWithIndex'] = __NAMESPACE__ . '\\majData_majFoldablemajWithmajIndex_foldmajMapmajWithmajIndex';

// Data_FoldableWithIndex_foldableWithIndexApp
function majData_majFoldablemajWithmajIndex_foldablemajWithmajIndexmajApp($dictFoldableWithIndex_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldablemajWithmajIndex_foldablemajWithmajIndexmajApp';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictFoldableWithIndex_0)->{'Foldable0'})(null);
  $foldableApp_2_1 = (object)["foldr" => (function() use ($__local_var_1_0) {
  $__fn = function($f_2, $i_3 = null, $v_4 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($__local_var_1_0)->{'foldr'})($f_2))($i_3))($v_4);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldl" => (function() use ($__local_var_1_0) {
  $__fn = function($f_2, $i_3 = null, $v_4 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($__local_var_1_0)->{'foldl'})($f_2))($i_3))($v_4);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldMap" => function($dictMonoid_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = (($__local_var_1_0)->{'foldMap'})($dictMonoid_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["foldrWithIndex" => (function() use ($dictFoldableWithIndex_0) {
  $__fn = function($f_3, $z_4 = null, $v_5 = null) use ($dictFoldableWithIndex_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($dictFoldableWithIndex_0)->{'foldrWithIndex'})($f_3))($z_4))($v_5);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldlWithIndex" => (function() use ($dictFoldableWithIndex_0) {
  $__fn = function($f_3, $z_4 = null, $v_5 = null) use ($dictFoldableWithIndex_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($dictFoldableWithIndex_0)->{'foldlWithIndex'})($f_3))($z_4))($v_5);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldMapWithIndex" => function($dictMonoid_3) use ($dictFoldableWithIndex_0) {
  $__num = \func_num_args();
  $__res = (($dictFoldableWithIndex_0)->{'foldMapWithIndex'})($dictMonoid_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable0" => function($_dollar__unused_3) use ($foldableApp_2_1) {
  $__num = \func_num_args();
  $__res = $foldableApp_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_FoldableWithIndex_foldableWithIndexApp'] = __NAMESPACE__ . '\\majData_majFoldablemajWithmajIndex_foldablemajWithmajIndexmajApp';

// Data_FoldableWithIndex_foldableWithIndexCompose
function majData_majFoldablemajWithmajIndex_foldablemajWithmajIndexmajCompose($dictFoldableWithIndex_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldablemajWithmajIndex_foldablemajWithmajIndexmajCompose';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictFoldableWithIndex_0)->{'Foldable0'})(null);
  $__res = function($dictFoldableWithIndex1_2) use ($__local_var_1_0, $dictFoldableWithIndex_0) {
  $__num = \func_num_args();
  $foldlWithIndex2_3_1 = ($dictFoldableWithIndex1_2)->{'foldlWithIndex'};
  $__local_var_4_2 = (($dictFoldableWithIndex1_2)->{'Foldable0'})(null);
  $foldableCompose1_5_3 = (object)["foldr" => (function() use ($__local_var_1_0, $__local_var_4_2) {
  $__fn = function($f_5, $i_6 = null, $v_7 = null) use ($__local_var_1_0, $__local_var_4_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__local_var_8_3 = (($__local_var_4_2)->{'foldr'})($f_5);
  $__res = (((($__local_var_1_0)->{'foldr'})((function() use ($__local_var_8_3) {
  $__fn = function($b_9, $a_10 = null) use ($__local_var_8_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($__local_var_8_3)($a_10))($b_9);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($i_6))($v_7);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldl" => (function() use ($__local_var_1_0, $__local_var_4_2) {
  $__fn = function($f_5, $i_6 = null, $v_7 = null) use ($__local_var_1_0, $__local_var_4_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($__local_var_1_0)->{'foldl'})((($__local_var_4_2)->{'foldl'})($f_5)))($i_6))($v_7);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldMap" => function($dictMonoid_5) use ($__local_var_1_0, $__local_var_4_2) {
  $__num = \func_num_args();
  $foldMap4_6_4 = (($__local_var_1_0)->{'foldMap'})($dictMonoid_5);
  $foldMap5_7_5 = (($__local_var_4_2)->{'foldMap'})($dictMonoid_5);
  $__res = (function() use ($foldMap4_6_4, $foldMap5_7_5) {
  $__fn = function($f_8, $v_9 = null) use ($foldMap4_6_4, $foldMap5_7_5, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($foldMap4_6_4)(($foldMap5_7_5)($f_8)))($v_9);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["foldrWithIndex" => (function() use ($dictFoldableWithIndex1_2, $dictFoldableWithIndex_0) {
  $__fn = function($f_6, $i_7 = null, $v_8 = null) use ($dictFoldableWithIndex1_2, $dictFoldableWithIndex_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($dictFoldableWithIndex_0)->{'foldrWithIndex'})(function($a_9) use ($dictFoldableWithIndex1_2, $f_6) {
  $__num = \func_num_args();
  $__local_var_10_7 = (($dictFoldableWithIndex1_2)->{'foldrWithIndex'})(function($b_10) use ($a_9, $f_6) {
  $__num = \func_num_args();
  $__res = ($f_6)(new \Data\Tuple\Data_Tuple_Tuple($a_9, $b_10));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = (function() use ($__local_var_10_7) {
  $__fn = function($b_11, $a_12 = null) use ($__local_var_10_7, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($__local_var_10_7)($a_12))($b_11);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($i_7))($v_8);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldlWithIndex" => (function() use ($dictFoldableWithIndex_0, $foldlWithIndex2_3_1) {
  $__fn = function($f_6, $i_7 = null, $v_8 = null) use ($dictFoldableWithIndex_0, $foldlWithIndex2_3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($dictFoldableWithIndex_0)->{'foldlWithIndex'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($foldlWithIndex2_3_1))((function() use ($f_6) {
  $__fn = function($a_9, $b_10 = null) use ($f_6, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($f_6)(new \Data\Tuple\Data_Tuple_Tuple($a_9, $b_10));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})())))($i_7))($v_8);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldMapWithIndex" => function($dictMonoid_6) use ($dictFoldableWithIndex1_2, $dictFoldableWithIndex_0) {
  $__num = \func_num_args();
  $foldMapWithIndex3_7_8 = (($dictFoldableWithIndex_0)->{'foldMapWithIndex'})($dictMonoid_6);
  $foldMapWithIndex4_8_9 = (($dictFoldableWithIndex1_2)->{'foldMapWithIndex'})($dictMonoid_6);
  $__res = (function() use ($foldMapWithIndex3_7_8, $foldMapWithIndex4_8_9) {
  $__fn = function($f_9, $v_10 = null) use ($foldMapWithIndex3_7_8, $foldMapWithIndex4_8_9, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($foldMapWithIndex3_7_8)((($GLOBALS['Control_Semigroupoid_composeImpl'])($foldMapWithIndex4_8_9))((function() use ($f_9) {
  $__fn = function($a_11, $b_12 = null) use ($f_9, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($f_9)(new \Data\Tuple\Data_Tuple_Tuple($a_11, $b_12));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})())))($v_10);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable0" => function($_dollar__unused_6) use ($foldableCompose1_5_3) {
  $__num = \func_num_args();
  $__res = $foldableCompose1_5_3;
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
$GLOBALS['Data_FoldableWithIndex_foldableWithIndexCompose'] = __NAMESPACE__ . '\\majData_majFoldablemajWithmajIndex_foldablemajWithmajIndexmajCompose';

// Data_FoldableWithIndex_foldableWithIndexCoproduct
function majData_majFoldablemajWithmajIndex_foldablemajWithmajIndexmajCoproduct($dictFoldableWithIndex_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldablemajWithmajIndex_foldablemajWithmajIndexmajCoproduct';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $foldableCoproduct_1_0 = ($GLOBALS['Data_Foldable_foldableCoproduct'])((($dictFoldableWithIndex_0)->{'Foldable0'})(null));
  $__res = function($dictFoldableWithIndex1_2) use ($dictFoldableWithIndex_0, $foldableCoproduct_1_0) {
  $__num = \func_num_args();
  $foldableCoproduct1_3_1 = ($foldableCoproduct_1_0)((($dictFoldableWithIndex1_2)->{'Foldable0'})(null));
  $__res = (object)["foldrWithIndex" => (function() use ($dictFoldableWithIndex1_2, $dictFoldableWithIndex_0) {
  $__fn = function($f_4, $z_5 = null) use ($dictFoldableWithIndex1_2, $dictFoldableWithIndex_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_6_2 = ((($dictFoldableWithIndex_0)->{'foldrWithIndex'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_4))($GLOBALS['Data_Either_Left'])))($z_5);
  $__local_var_7_3 = ((($dictFoldableWithIndex1_2)->{'foldrWithIndex'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_4))($GLOBALS['Data_Either_Right'])))($z_5);
  $__res = function($v2_8) use ($__local_var_6_2, $__local_var_7_3) {
  $__num = \func_num_args();
  $__t4 = null;;
  if ($v2_8 instanceof \Data\Either\Data_Either_Left) {
$__t4 = ($__local_var_6_2)(($v2_8)->{'value0'});
goto end_branch_4;;
};
  if ($v2_8 instanceof \Data\Either\Data_Either_Right) {
$__t4 = ($__local_var_7_3)(($v2_8)->{'value0'});
goto end_branch_4;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t4 = null;
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "foldlWithIndex" => (function() use ($dictFoldableWithIndex1_2, $dictFoldableWithIndex_0) {
  $__fn = function($f_4, $z_5 = null) use ($dictFoldableWithIndex1_2, $dictFoldableWithIndex_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_6_5 = ((($dictFoldableWithIndex_0)->{'foldlWithIndex'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_4))($GLOBALS['Data_Either_Left'])))($z_5);
  $__local_var_7_6 = ((($dictFoldableWithIndex1_2)->{'foldlWithIndex'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_4))($GLOBALS['Data_Either_Right'])))($z_5);
  $__res = function($v2_8) use ($__local_var_6_5, $__local_var_7_6) {
  $__num = \func_num_args();
  $__t7 = null;;
  if ($v2_8 instanceof \Data\Either\Data_Either_Left) {
$__t7 = ($__local_var_6_5)(($v2_8)->{'value0'});
goto end_branch_7;;
};
  if ($v2_8 instanceof \Data\Either\Data_Either_Right) {
$__t7 = ($__local_var_7_6)(($v2_8)->{'value0'});
goto end_branch_7;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t7 = null;
  end_branch_7:;
  $__res = $__t7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "foldMapWithIndex" => function($dictMonoid_4) use ($dictFoldableWithIndex1_2, $dictFoldableWithIndex_0) {
  $__num = \func_num_args();
  $foldMapWithIndex3_5_8 = (($dictFoldableWithIndex_0)->{'foldMapWithIndex'})($dictMonoid_4);
  $foldMapWithIndex4_6_9 = (($dictFoldableWithIndex1_2)->{'foldMapWithIndex'})($dictMonoid_4);
  $__res = function($f_7) use ($foldMapWithIndex3_5_8, $foldMapWithIndex4_6_9) {
  $__num = \func_num_args();
  $__local_var_8_10 = ($foldMapWithIndex3_5_8)((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_7))($GLOBALS['Data_Either_Left']));
  $__local_var_9_11 = ($foldMapWithIndex4_6_9)((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_7))($GLOBALS['Data_Either_Right']));
  $__res = function($v2_10) use ($__local_var_8_10, $__local_var_9_11) {
  $__num = \func_num_args();
  $__t12 = null;;
  if ($v2_10 instanceof \Data\Either\Data_Either_Left) {
$__t12 = ($__local_var_8_10)(($v2_10)->{'value0'});
goto end_branch_12;;
};
  if ($v2_10 instanceof \Data\Either\Data_Either_Right) {
$__t12 = ($__local_var_9_11)(($v2_10)->{'value0'});
goto end_branch_12;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t12 = null;
  end_branch_12:;
  $__res = $__t12;
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
}, "Foldable0" => function($_dollar__unused_4) use ($foldableCoproduct1_3_1) {
  $__num = \func_num_args();
  $__res = $foldableCoproduct1_3_1;
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
$GLOBALS['Data_FoldableWithIndex_foldableWithIndexCoproduct'] = __NAMESPACE__ . '\\majData_majFoldablemajWithmajIndex_foldablemajWithmajIndexmajCoproduct';

// Data_FoldableWithIndex_foldableWithIndexProduct
function majData_majFoldablemajWithmajIndex_foldablemajWithmajIndexmajProduct($dictFoldableWithIndex_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldablemajWithmajIndex_foldablemajWithmajIndexmajProduct';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $foldableProduct_1_0 = ($GLOBALS['Data_Foldable_foldableProduct'])((($dictFoldableWithIndex_0)->{'Foldable0'})(null));
  $__res = function($dictFoldableWithIndex1_2) use ($dictFoldableWithIndex_0, $foldableProduct_1_0) {
  $__num = \func_num_args();
  $foldableProduct1_3_1 = ($foldableProduct_1_0)((($dictFoldableWithIndex1_2)->{'Foldable0'})(null));
  $__res = (object)["foldrWithIndex" => (function() use ($dictFoldableWithIndex1_2, $dictFoldableWithIndex_0) {
  $__fn = function($f_4, $z_5 = null, $v_6 = null) use ($dictFoldableWithIndex1_2, $dictFoldableWithIndex_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($dictFoldableWithIndex_0)->{'foldrWithIndex'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_4))($GLOBALS['Data_Either_Left'])))((((($dictFoldableWithIndex1_2)->{'foldrWithIndex'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_4))($GLOBALS['Data_Either_Right'])))($z_5))(($v_6)->{'value1'})))(($v_6)->{'value0'});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldlWithIndex" => (function() use ($dictFoldableWithIndex1_2, $dictFoldableWithIndex_0) {
  $__fn = function($f_4, $z_5 = null, $v_6 = null) use ($dictFoldableWithIndex1_2, $dictFoldableWithIndex_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($dictFoldableWithIndex1_2)->{'foldlWithIndex'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_4))($GLOBALS['Data_Either_Right'])))((((($dictFoldableWithIndex_0)->{'foldlWithIndex'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_4))($GLOBALS['Data_Either_Left'])))($z_5))(($v_6)->{'value0'})))(($v_6)->{'value1'});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldMapWithIndex" => function($dictMonoid_4) use ($dictFoldableWithIndex1_2, $dictFoldableWithIndex_0) {
  $__num = \func_num_args();
  $foldMapWithIndex3_5_2 = (($dictFoldableWithIndex_0)->{'foldMapWithIndex'})($dictMonoid_4);
  $foldMapWithIndex4_6_3 = (($dictFoldableWithIndex1_2)->{'foldMapWithIndex'})($dictMonoid_4);
  $__res = (function() use ($dictMonoid_4, $foldMapWithIndex3_5_2, $foldMapWithIndex4_6_3) {
  $__fn = function($f_7, $v_8 = null) use ($dictMonoid_4, $foldMapWithIndex3_5_2, $foldMapWithIndex4_6_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((($dictMonoid_4)->{'Semigroup0'})(null))->{'append'})((($foldMapWithIndex3_5_2)((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_7))($GLOBALS['Data_Either_Left'])))(($v_8)->{'value0'})))((($foldMapWithIndex4_6_3)((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_7))($GLOBALS['Data_Either_Right'])))(($v_8)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable0" => function($_dollar__unused_4) use ($foldableProduct1_3_1) {
  $__num = \func_num_args();
  $__res = $foldableProduct1_3_1;
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
$GLOBALS['Data_FoldableWithIndex_foldableWithIndexProduct'] = __NAMESPACE__ . '\\majData_majFoldablemajWithmajIndex_foldablemajWithmajIndexmajProduct';

// Data_FoldableWithIndex_foldlWithIndexDefault
function majData_majFoldablemajWithmajIndex_foldlmajWithmajIndexmajDefault($dictFoldableWithIndex_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldablemajWithmajIndex_foldlmajWithmajIndexmajDefault';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $foldMapWithIndex1_1_0 = (($dictFoldableWithIndex_0)->{'foldMapWithIndex'})($GLOBALS['Data_FoldableWithIndex_monoidDual']);
  $__res = (function() use ($foldMapWithIndex1_1_0) {
  $__fn = function($c_2, $u_3 = null, $xs_4 = null) use ($foldMapWithIndex1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($foldMapWithIndex1_1_0)(function($i_5) use ($c_2) {
  $__num = \func_num_args();
  $__local_var_6_1 = ($c_2)($i_5);
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Monoid_Dual_Dual']))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Monoid_Endo_Endo']))((function() use ($__local_var_6_1) {
  $__fn = function($b_7, $a_8 = null) use ($__local_var_6_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($__local_var_6_1)($a_8))($b_7);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($xs_4))($u_3);
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
$GLOBALS['Data_FoldableWithIndex_foldlWithIndexDefault'] = __NAMESPACE__ . '\\majData_majFoldablemajWithmajIndex_foldlmajWithmajIndexmajDefault';

// Data_FoldableWithIndex_foldrWithIndexDefault
function majData_majFoldablemajWithmajIndex_foldrmajWithmajIndexmajDefault($dictFoldableWithIndex_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldablemajWithmajIndex_foldrmajWithmajIndexmajDefault';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $foldMapWithIndex1_1_0 = (($dictFoldableWithIndex_0)->{'foldMapWithIndex'})($GLOBALS['Data_FoldableWithIndex_monoidEndo']);
  $__res = (function() use ($foldMapWithIndex1_1_0) {
  $__fn = function($c_2, $u_3 = null, $xs_4 = null) use ($foldMapWithIndex1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($foldMapWithIndex1_1_0)(function($i_5) use ($c_2) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Monoid_Endo_Endo']))(($c_2)($i_5));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($xs_4))($u_3);
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
$GLOBALS['Data_FoldableWithIndex_foldrWithIndexDefault'] = __NAMESPACE__ . '\\majData_majFoldablemajWithmajIndex_foldrmajWithmajIndexmajDefault';

// Data_FoldableWithIndex_surroundMapWithIndex
function majData_majFoldablemajWithmajIndex_surroundmajMapmajWithmajIndex($dictFoldableWithIndex_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldablemajWithmajIndex_surroundmajMapmajWithmajIndex';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $foldMapWithIndex1_1_0 = (($dictFoldableWithIndex_0)->{'foldMapWithIndex'})($GLOBALS['Data_FoldableWithIndex_monoidEndo']);
  $__res = (function() use ($foldMapWithIndex1_1_0) {
  $__fn = function($dictSemigroup_2, $d_3 = null, $t_4 = null, $f_5 = null) use ($foldMapWithIndex1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = ((($foldMapWithIndex1_1_0)((function() use ($d_3, $dictSemigroup_2, $t_4) {
  $__fn = function($i_6, $a_7 = null, $m_8 = null) use ($d_3, $dictSemigroup_2, $t_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictSemigroup_2)->{'append'})($d_3))(((($dictSemigroup_2)->{'append'})((($t_4)($i_6))($a_7)))($m_8));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()))($f_5))($d_3);
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_FoldableWithIndex_surroundMapWithIndex'] = __NAMESPACE__ . '\\majData_majFoldablemajWithmajIndex_surroundmajMapmajWithmajIndex';

// Data_FoldableWithIndex_foldMapDefault
function majData_majFoldablemajWithmajIndex_foldmajMapmajDefault($dictFoldableWithIndex_0, $dictMonoid_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldablemajWithmajIndex_foldmajMapmajDefault';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $foldMapWithIndex2_2_0 = (($dictFoldableWithIndex_0)->{'foldMapWithIndex'})($dictMonoid_1);
  $__res = function($f_3) use ($foldMapWithIndex2_2_0) {
  $__num = \func_num_args();
  $__res = ($foldMapWithIndex2_2_0)(function($v_4) use ($f_3) {
  $__num = \func_num_args();
  $__res = $f_3;
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
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_FoldableWithIndex_foldMapDefault'] = __NAMESPACE__ . '\\majData_majFoldablemajWithmajIndex_foldmajMapmajDefault';

// Data_FoldableWithIndex_findWithIndex
function majData_majFoldablemajWithmajIndex_findmajWithmajIndex($dictFoldableWithIndex_0, $p_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldablemajWithmajIndex_findmajWithmajIndex';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictFoldableWithIndex_0)->{'foldlWithIndex'})((function() use ($p_1) {
  $__fn = function($v_2, $v1_3 = null, $v2_4 = null) use ($p_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if (($v1_3 instanceof \Data\Maybe\Data_Maybe_Nothing && (($p_1)($v_2))($v2_4))) {
$__t0 = new \Data\Maybe\Data_Maybe_Just((object)["index" => $v_2, "value" => $v2_4]);
goto end_branch_0;;
};
  $__t0 = $v1_3;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()))(new \Data\Maybe\Data_Maybe_Nothing());
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_FoldableWithIndex_findWithIndex'] = __NAMESPACE__ . '\\majData_majFoldablemajWithmajIndex_findmajWithmajIndex';

// Data_FoldableWithIndex_findMapWithIndex
function majData_majFoldablemajWithmajIndex_findmajMapmajWithmajIndex($dictFoldableWithIndex_0, $f_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldablemajWithmajIndex_findmajMapmajWithmajIndex';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictFoldableWithIndex_0)->{'foldlWithIndex'})((function() use ($f_1) {
  $__fn = function($v_2, $v1_3 = null, $v2_4 = null) use ($f_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if ($v1_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t0 = (($f_1)($v_2))($v2_4);
goto end_branch_0;;
};
  $__t0 = $v1_3;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()))(new \Data\Maybe\Data_Maybe_Nothing());
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_FoldableWithIndex_findMapWithIndex'] = __NAMESPACE__ . '\\majData_majFoldablemajWithmajIndex_findmajMapmajWithmajIndex';

// Data_FoldableWithIndex_anyWithIndex
function majData_majFoldablemajWithmajIndex_anymajWithmajIndex($dictFoldableWithIndex_0, $dictHeytingAlgebra_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldablemajWithmajIndex_anymajWithmajIndex';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $semigroupDisj1_2_0 = (object)["append" => (function() use ($dictHeytingAlgebra_1) {
  $__fn = function($v_2, $v1_3 = null) use ($dictHeytingAlgebra_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictHeytingAlgebra_1)->{'disj'})($v_2))($v1_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $foldMapWithIndex2_2_0 = (($dictFoldableWithIndex_0)->{'foldMapWithIndex'})((object)["mempty" => ($dictHeytingAlgebra_1)->{'ff'}, "Semigroup0" => function($_dollar__unused_3) use ($semigroupDisj1_2_0) {
  $__num = \func_num_args();
  $__res = $semigroupDisj1_2_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]);
  $__res = function($t_3) use ($foldMapWithIndex2_2_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Unsafe_Coerce_unsafeCoerce']))(($foldMapWithIndex2_2_0)(function($i_4) use ($t_3) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Monoid_Disj_Disj']))(($t_3)($i_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_FoldableWithIndex_anyWithIndex'] = __NAMESPACE__ . '\\majData_majFoldablemajWithmajIndex_anymajWithmajIndex';

// Data_FoldableWithIndex_allWithIndex
function majData_majFoldablemajWithmajIndex_allmajWithmajIndex($dictFoldableWithIndex_0, $dictHeytingAlgebra_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldablemajWithmajIndex_allmajWithmajIndex';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $semigroupConj1_2_0 = (object)["append" => (function() use ($dictHeytingAlgebra_1) {
  $__fn = function($v_2, $v1_3 = null) use ($dictHeytingAlgebra_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictHeytingAlgebra_1)->{'conj'})($v_2))($v1_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $foldMapWithIndex2_2_0 = (($dictFoldableWithIndex_0)->{'foldMapWithIndex'})((object)["mempty" => ($dictHeytingAlgebra_1)->{'tt'}, "Semigroup0" => function($_dollar__unused_3) use ($semigroupConj1_2_0) {
  $__num = \func_num_args();
  $__res = $semigroupConj1_2_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]);
  $__res = function($t_3) use ($foldMapWithIndex2_2_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Unsafe_Coerce_unsafeCoerce']))(($foldMapWithIndex2_2_0)(function($i_4) use ($t_3) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Monoid_Conj_Conj']))(($t_3)($i_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_FoldableWithIndex_allWithIndex'] = __NAMESPACE__ . '\\majData_majFoldablemajWithmajIndex_allmajWithmajIndex';

