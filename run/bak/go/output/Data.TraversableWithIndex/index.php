<?php

namespace Data\TraversableWithIndex;

// ALL IMPORTS: Control.Applicative, Control.Apply, Control.Semigroupoid, Data.Const, Data.Either, Data.FoldableWithIndex, Data.Function, Data.Functor, Data.Functor.App, Data.Functor.Compose, Data.Functor.Coproduct, Data.Functor.Product, Data.FunctorWithIndex, Data.Identity, Data.Maybe, Data.Maybe.First, Data.Maybe.Last, Data.Monoid.Additive, Data.Monoid.Conj, Data.Monoid.Disj, Data.Monoid.Dual, Data.Monoid.Multiplicative, Data.Traversable, Data.Traversable.Accum, Data.Traversable.Accum.Internal, Data.TraversableWithIndex, Data.Tuple, Data.Unit, Prelude, Prim
// TO REQUIRE: Control.Applicative, Control.Apply, Control.Semigroupoid, Data.Const, Data.Either, Data.FoldableWithIndex, Data.Function, Data.Functor, Data.Functor.App, Data.Functor.Compose, Data.Functor.Coproduct, Data.Functor.Product, Data.FunctorWithIndex, Data.Identity, Data.Maybe, Data.Maybe.First, Data.Maybe.Last, Data.Monoid.Additive, Data.Monoid.Conj, Data.Monoid.Disj, Data.Monoid.Dual, Data.Monoid.Multiplicative, Data.Traversable, Data.Traversable.Accum, Data.Traversable.Accum.Internal, Data.TraversableWithIndex, Data.Tuple, Data.Unit, Prelude
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Const/index.php';
require_once __DIR__ . '/../Data.Either/index.php';
require_once __DIR__ . '/../Data.FoldableWithIndex/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Functor.App/index.php';
require_once __DIR__ . '/../Data.Functor.Compose/index.php';
require_once __DIR__ . '/../Data.Functor.Coproduct/index.php';
require_once __DIR__ . '/../Data.Functor.Product/index.php';
require_once __DIR__ . '/../Data.FunctorWithIndex/index.php';
require_once __DIR__ . '/../Data.Identity/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Maybe.First/index.php';
require_once __DIR__ . '/../Data.Maybe.Last/index.php';
require_once __DIR__ . '/../Data.Monoid.Additive/index.php';
require_once __DIR__ . '/../Data.Monoid.Conj/index.php';
require_once __DIR__ . '/../Data.Monoid.Disj/index.php';
require_once __DIR__ . '/../Data.Monoid.Dual/index.php';
require_once __DIR__ . '/../Data.Monoid.Multiplicative/index.php';
require_once __DIR__ . '/../Data.Traversable/index.php';
require_once __DIR__ . '/../Data.Traversable.Accum/index.php';
require_once __DIR__ . '/../Data.Traversable.Accum.Internal/index.php';
require_once __DIR__ . '/../Data.TraversableWithIndex/index.php';
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
\PhpursThunks::$thunks['Data_TraversableWithIndex_traverseWithIndexDefault'] = function() { $v = (function() {
  $__fn = function($dictTraversableWithIndex_0 = null, $dictApplicative_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $sequence1_2_0 = (((($dictTraversableWithIndex_0)['Traversable2'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['sequence'])($dictApplicative_1);
  $__res = function($f_3 = null) use ($dictTraversableWithIndex_0, $sequence1_2_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))($sequence1_2_0))((((($dictTraversableWithIndex_0)['FunctorWithIndex0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['mapWithIndex'])($f_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_TraversableWithIndex_traverseWithIndex'] = function() { $v = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)['traverseWithIndex'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_TraversableWithIndex_traverseDefault'] = function() { $v = (function() {
  $__fn = function($dictTraversableWithIndex_0 = null, $dictApplicative_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $traverseWithIndex2_2_0 = (($dictTraversableWithIndex_0)['traverseWithIndex'])($dictApplicative_1);
  $__res = function($f_3 = null) use ($traverseWithIndex2_2_0) {
  $__num = \func_num_args();
  $__res = ($traverseWithIndex2_2_0)(function($v_4 = null) use ($f_3) {
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
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_TraversableWithIndex_traversableWithIndexTuple'] = function() { $v = ["traverseWithIndex" => (function() {
  $__fn = function($dictApplicative_0 = null, $f_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((((($dictApplicative_0)['Apply0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['map'])((($GLOBALS['Data_Tuple_Tuple'] ?? \PhpursThunks::eval('Data_Tuple_Tuple')))(($v_2)->{'value0'})))((($f_1)(($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit'))))(($v_2)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "FunctorWithIndex0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_FunctorWithIndex_functorWithIndexTuple'] ?? \PhpursThunks::eval('Data_FunctorWithIndex_functorWithIndexTuple'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_FoldableWithIndex_foldableWithIndexTuple'] ?? \PhpursThunks::eval('Data_FoldableWithIndex_foldableWithIndexTuple'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Traversable_traversableTuple'] ?? \PhpursThunks::eval('Data_Traversable_traversableTuple'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_TraversableWithIndex_traversableWithIndexProduct'] = function() { $v = function($dictTraversableWithIndex_0 = null) {
  $__num = \func_num_args();
  $functorWithIndexProduct_1_0 = (($GLOBALS['Data_FunctorWithIndex_functorWithIndexProduct'] ?? \PhpursThunks::eval('Data_FunctorWithIndex_functorWithIndexProduct')))((($dictTraversableWithIndex_0)['FunctorWithIndex0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $foldableWithIndexProduct_2_1 = (($GLOBALS['Data_FoldableWithIndex_foldableWithIndexProduct'] ?? \PhpursThunks::eval('Data_FoldableWithIndex_foldableWithIndexProduct')))((($dictTraversableWithIndex_0)['FoldableWithIndex1'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $traversableProduct_3_2 = (($GLOBALS['Data_Traversable_traversableProduct'] ?? \PhpursThunks::eval('Data_Traversable_traversableProduct')))((($dictTraversableWithIndex_0)['Traversable2'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = function($dictTraversableWithIndex1_4 = null) use ($dictTraversableWithIndex_0, $foldableWithIndexProduct_2_1, $functorWithIndexProduct_1_0, $traversableProduct_3_2) {
  $__num = \func_num_args();
  $functorWithIndexProduct1_5_3 = ($functorWithIndexProduct_1_0)((($dictTraversableWithIndex1_4)['FunctorWithIndex0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $foldableWithIndexProduct1_6_4 = ($foldableWithIndexProduct_2_1)((($dictTraversableWithIndex1_4)['FoldableWithIndex1'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $traversableProduct1_7_5 = ($traversableProduct_3_2)((($dictTraversableWithIndex1_4)['Traversable2'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = ["traverseWithIndex" => function($dictApplicative_8 = null) use ($dictTraversableWithIndex1_4, $dictTraversableWithIndex_0) {
  $__num = \func_num_args();
  $lift2_9_6 = (($GLOBALS['Control_Apply_lift2'] ?? \PhpursThunks::eval('Control_Apply_lift2')))((($dictApplicative_8)['Apply0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $traverseWithIndex3_10_7 = (($dictTraversableWithIndex_0)['traverseWithIndex'])($dictApplicative_8);
  $traverseWithIndex4_11_8 = (($dictTraversableWithIndex1_4)['traverseWithIndex'])($dictApplicative_8);
  $__res = (function() use ($lift2_9_6, $traverseWithIndex3_10_7, $traverseWithIndex4_11_8) {
  $__fn = function($f_12 = null, $v_13 = null) use ($lift2_9_6, $traverseWithIndex3_10_7, $traverseWithIndex4_11_8, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($lift2_9_6)(($GLOBALS['Data_Functor_Product_product'] ?? \PhpursThunks::eval('Data_Functor_Product_product'))))((($traverseWithIndex3_10_7)(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))($f_12))(($GLOBALS['Data_Either_Left'] ?? \PhpursThunks::eval('Data_Either_Left')))))(($v_13)->{'value0'})))((($traverseWithIndex4_11_8)(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))($f_12))(($GLOBALS['Data_Either_Right'] ?? \PhpursThunks::eval('Data_Either_Right')))))(($v_13)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FunctorWithIndex0" => function($dollar__unused_8 = null) use ($functorWithIndexProduct1_5_3) {
  $__num = \func_num_args();
  $__res = $functorWithIndexProduct1_5_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($dollar__unused_8 = null) use ($foldableWithIndexProduct1_6_4) {
  $__num = \func_num_args();
  $__res = $foldableWithIndexProduct1_6_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($dollar__unused_8 = null) use ($traversableProduct1_7_5) {
  $__num = \func_num_args();
  $__res = $traversableProduct1_7_5;
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
}; return $v; };
\PhpursThunks::$thunks['Data_TraversableWithIndex_traversableWithIndexMultiplicative'] = function() { $v = ["traverseWithIndex" => (function() {
  $__fn = function($dictApplicative_0 = null, $f_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($GLOBALS['Data_Traversable_traversableMultiplicative'] ?? \PhpursThunks::eval('Data_Traversable_traversableMultiplicative')))['traverse'])($dictApplicative_0))(($f_1)(($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit'))));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "FunctorWithIndex0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_FunctorWithIndex_functorWithIndexMultiplicative'] ?? \PhpursThunks::eval('Data_FunctorWithIndex_functorWithIndexMultiplicative'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_FoldableWithIndex_foldableWithIndexMultiplicative'] ?? \PhpursThunks::eval('Data_FoldableWithIndex_foldableWithIndexMultiplicative'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Traversable_traversableMultiplicative'] ?? \PhpursThunks::eval('Data_Traversable_traversableMultiplicative'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_TraversableWithIndex_traversableWithIndexMaybe'] = function() { $v = ["traverseWithIndex" => (function() {
  $__fn = function($dictApplicative_0 = null, $f_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($GLOBALS['Data_Traversable_traversableMaybe'] ?? \PhpursThunks::eval('Data_Traversable_traversableMaybe')))['traverse'])($dictApplicative_0))(($f_1)(($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit'))));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "FunctorWithIndex0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_FunctorWithIndex_functorWithIndexMaybe'] ?? \PhpursThunks::eval('Data_FunctorWithIndex_functorWithIndexMaybe'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_FoldableWithIndex_foldableWithIndexMaybe'] ?? \PhpursThunks::eval('Data_FoldableWithIndex_foldableWithIndexMaybe'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Traversable_traversableMaybe'] ?? \PhpursThunks::eval('Data_Traversable_traversableMaybe'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_TraversableWithIndex_traversableWithIndexLast'] = function() { $v = ["traverseWithIndex" => (function() {
  $__fn = function($dictApplicative_0 = null, $f_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($GLOBALS['Data_Traversable_traversableLast'] ?? \PhpursThunks::eval('Data_Traversable_traversableLast')))['traverse'])($dictApplicative_0))(($f_1)(($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit'))));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "FunctorWithIndex0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_FunctorWithIndex_functorWithIndexLast'] ?? \PhpursThunks::eval('Data_FunctorWithIndex_functorWithIndexLast'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_FoldableWithIndex_foldableWithIndexLast'] ?? \PhpursThunks::eval('Data_FoldableWithIndex_foldableWithIndexLast'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Traversable_traversableLast'] ?? \PhpursThunks::eval('Data_Traversable_traversableLast'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_TraversableWithIndex_traversableWithIndexIdentity'] = function() { $v = ["traverseWithIndex" => (function() {
  $__fn = function($dictApplicative_0 = null, $f_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((((($dictApplicative_0)['Apply0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['map'])(($GLOBALS['Data_Identity_Identity'] ?? \PhpursThunks::eval('Data_Identity_Identity'))))((($f_1)(($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit'))))($v_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "FunctorWithIndex0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_FunctorWithIndex_functorWithIndexIdentity'] ?? \PhpursThunks::eval('Data_FunctorWithIndex_functorWithIndexIdentity'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_FoldableWithIndex_foldableWithIndexIdentity'] ?? \PhpursThunks::eval('Data_FoldableWithIndex_foldableWithIndexIdentity'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Traversable_traversableIdentity'] ?? \PhpursThunks::eval('Data_Traversable_traversableIdentity'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_TraversableWithIndex_traversableWithIndexFirst'] = function() { $v = ["traverseWithIndex" => (function() {
  $__fn = function($dictApplicative_0 = null, $f_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($GLOBALS['Data_Traversable_traversableFirst'] ?? \PhpursThunks::eval('Data_Traversable_traversableFirst')))['traverse'])($dictApplicative_0))(($f_1)(($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit'))));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "FunctorWithIndex0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_FunctorWithIndex_functorWithIndexFirst'] ?? \PhpursThunks::eval('Data_FunctorWithIndex_functorWithIndexFirst'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_FoldableWithIndex_foldableWithIndexFirst'] ?? \PhpursThunks::eval('Data_FoldableWithIndex_foldableWithIndexFirst'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Traversable_traversableFirst'] ?? \PhpursThunks::eval('Data_Traversable_traversableFirst'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_TraversableWithIndex_traversableWithIndexEither'] = function() { $v = ["traverseWithIndex" => (function() {
  $__fn = function($dictApplicative_0 = null, $v_1 = null, $v1_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if ((is_object($v1_2) && (($v1_2)->{'tag'} === "Left"))) {
$__t0 = (($dictApplicative_0)['pure'])(new Phpurs_Data1("Left", ($v1_2)->{'value0'}));
goto end_branch_0;;
};
  if ((is_object($v1_2) && (($v1_2)->{'tag'} === "Right"))) {
$__t0 = ((((((($dictApplicative_0)['Apply0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['map'])(($GLOBALS['Data_Either_Right'] ?? \PhpursThunks::eval('Data_Either_Right'))))((($v_1)(($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit'))))(($v1_2)->{'value0'}));
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
})(), "FunctorWithIndex0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_FunctorWithIndex_functorWithIndexEither'] ?? \PhpursThunks::eval('Data_FunctorWithIndex_functorWithIndexEither'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_FoldableWithIndex_foldableWithIndexEither'] ?? \PhpursThunks::eval('Data_FoldableWithIndex_foldableWithIndexEither'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Traversable_traversableEither'] ?? \PhpursThunks::eval('Data_Traversable_traversableEither'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_TraversableWithIndex_traversableWithIndexDual'] = function() { $v = ["traverseWithIndex" => (function() {
  $__fn = function($dictApplicative_0 = null, $f_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($GLOBALS['Data_Traversable_traversableDual'] ?? \PhpursThunks::eval('Data_Traversable_traversableDual')))['traverse'])($dictApplicative_0))(($f_1)(($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit'))));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "FunctorWithIndex0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_FunctorWithIndex_functorWithIndexDual'] ?? \PhpursThunks::eval('Data_FunctorWithIndex_functorWithIndexDual'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_FoldableWithIndex_foldableWithIndexDual'] ?? \PhpursThunks::eval('Data_FoldableWithIndex_foldableWithIndexDual'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Traversable_traversableDual'] ?? \PhpursThunks::eval('Data_Traversable_traversableDual'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_TraversableWithIndex_traversableWithIndexDisj'] = function() { $v = ["traverseWithIndex" => (function() {
  $__fn = function($dictApplicative_0 = null, $f_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($GLOBALS['Data_Traversable_traversableDisj'] ?? \PhpursThunks::eval('Data_Traversable_traversableDisj')))['traverse'])($dictApplicative_0))(($f_1)(($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit'))));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "FunctorWithIndex0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_FunctorWithIndex_functorWithIndexDisj'] ?? \PhpursThunks::eval('Data_FunctorWithIndex_functorWithIndexDisj'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_FoldableWithIndex_foldableWithIndexDisj'] ?? \PhpursThunks::eval('Data_FoldableWithIndex_foldableWithIndexDisj'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Traversable_traversableDisj'] ?? \PhpursThunks::eval('Data_Traversable_traversableDisj'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_TraversableWithIndex_traversableWithIndexCoproduct'] = function() { $v = function($dictTraversableWithIndex_0 = null) {
  $__num = \func_num_args();
  $functorWithIndexCoproduct_1_0 = (($GLOBALS['Data_FunctorWithIndex_functorWithIndexCoproduct'] ?? \PhpursThunks::eval('Data_FunctorWithIndex_functorWithIndexCoproduct')))((($dictTraversableWithIndex_0)['FunctorWithIndex0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $foldableWithIndexCoproduct_2_1 = (($GLOBALS['Data_FoldableWithIndex_foldableWithIndexCoproduct'] ?? \PhpursThunks::eval('Data_FoldableWithIndex_foldableWithIndexCoproduct')))((($dictTraversableWithIndex_0)['FoldableWithIndex1'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $traversableCoproduct_3_2 = (($GLOBALS['Data_Traversable_traversableCoproduct'] ?? \PhpursThunks::eval('Data_Traversable_traversableCoproduct')))((($dictTraversableWithIndex_0)['Traversable2'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = function($dictTraversableWithIndex1_4 = null) use ($dictTraversableWithIndex_0, $foldableWithIndexCoproduct_2_1, $functorWithIndexCoproduct_1_0, $traversableCoproduct_3_2) {
  $__num = \func_num_args();
  $functorWithIndexCoproduct1_5_3 = ($functorWithIndexCoproduct_1_0)((($dictTraversableWithIndex1_4)['FunctorWithIndex0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $foldableWithIndexCoproduct1_6_4 = ($foldableWithIndexCoproduct_2_1)((($dictTraversableWithIndex1_4)['FoldableWithIndex1'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $traversableCoproduct1_7_5 = ($traversableCoproduct_3_2)((($dictTraversableWithIndex1_4)['Traversable2'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = ["traverseWithIndex" => function($dictApplicative_8 = null) use ($dictTraversableWithIndex1_4, $dictTraversableWithIndex_0) {
  $__num = \func_num_args();
  $__local_var_9_6 = (((($dictApplicative_8)['Apply0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $traverseWithIndex3_10_7 = (($dictTraversableWithIndex_0)['traverseWithIndex'])($dictApplicative_8);
  $traverseWithIndex4_11_8 = (($dictTraversableWithIndex1_4)['traverseWithIndex'])($dictApplicative_8);
  $__res = function($f_12 = null) use ($__local_var_9_6, $traverseWithIndex3_10_7, $traverseWithIndex4_11_8) {
  $__num = \func_num_args();
  $__local_var_13_9 = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($__local_var_9_6)['map'])(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Functor_Coproduct_Coproduct'] ?? \PhpursThunks::eval('Data_Functor_Coproduct_Coproduct'))))(($GLOBALS['Data_Either_Left'] ?? \PhpursThunks::eval('Data_Either_Left'))))))(($traverseWithIndex3_10_7)(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))($f_12))(($GLOBALS['Data_Either_Left'] ?? \PhpursThunks::eval('Data_Either_Left')))));
  $__local_var_14_10 = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($__local_var_9_6)['map'])(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Functor_Coproduct_Coproduct'] ?? \PhpursThunks::eval('Data_Functor_Coproduct_Coproduct'))))(($GLOBALS['Data_Either_Right'] ?? \PhpursThunks::eval('Data_Either_Right'))))))(($traverseWithIndex4_11_8)(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))($f_12))(($GLOBALS['Data_Either_Right'] ?? \PhpursThunks::eval('Data_Either_Right')))));
  $__res = function($v2_15 = null) use ($__local_var_13_9, $__local_var_14_10) {
  $__num = \func_num_args();
  $__t11 = null;;
  if ((is_object($v2_15) && (($v2_15)->{'tag'} === "Left"))) {
$__t11 = ($__local_var_13_9)(($v2_15)->{'value0'});
goto end_branch_11;;
};
  if ((is_object($v2_15) && (($v2_15)->{'tag'} === "Right"))) {
$__t11 = ($__local_var_14_10)(($v2_15)->{'value0'});
goto end_branch_11;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t11 = null;
  end_branch_11:;
  $__res = $__t11;
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
}, "FunctorWithIndex0" => function($dollar__unused_8 = null) use ($functorWithIndexCoproduct1_5_3) {
  $__num = \func_num_args();
  $__res = $functorWithIndexCoproduct1_5_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($dollar__unused_8 = null) use ($foldableWithIndexCoproduct1_6_4) {
  $__num = \func_num_args();
  $__res = $foldableWithIndexCoproduct1_6_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($dollar__unused_8 = null) use ($traversableCoproduct1_7_5) {
  $__num = \func_num_args();
  $__res = $traversableCoproduct1_7_5;
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
}; return $v; };
\PhpursThunks::$thunks['Data_TraversableWithIndex_traversableWithIndexConst'] = function() { $v = ["traverseWithIndex" => (function() {
  $__fn = function($dictApplicative_0 = null, $v_1 = null, $v1_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($dictApplicative_0)['pure'])($v1_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "FunctorWithIndex0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_FunctorWithIndex_functorWithIndexConst'] ?? \PhpursThunks::eval('Data_FunctorWithIndex_functorWithIndexConst'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_FoldableWithIndex_foldableWithIndexConst'] ?? \PhpursThunks::eval('Data_FoldableWithIndex_foldableWithIndexConst'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Traversable_traversableConst'] ?? \PhpursThunks::eval('Data_Traversable_traversableConst'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_TraversableWithIndex_traversableWithIndexConj'] = function() { $v = ["traverseWithIndex" => (function() {
  $__fn = function($dictApplicative_0 = null, $f_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($GLOBALS['Data_Traversable_traversableConj'] ?? \PhpursThunks::eval('Data_Traversable_traversableConj')))['traverse'])($dictApplicative_0))(($f_1)(($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit'))));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "FunctorWithIndex0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_FunctorWithIndex_functorWithIndexConj'] ?? \PhpursThunks::eval('Data_FunctorWithIndex_functorWithIndexConj'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_FoldableWithIndex_foldableWithIndexConj'] ?? \PhpursThunks::eval('Data_FoldableWithIndex_foldableWithIndexConj'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Traversable_traversableConj'] ?? \PhpursThunks::eval('Data_Traversable_traversableConj'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_TraversableWithIndex_traversableWithIndexCompose'] = function() { $v = function($dictTraversableWithIndex_0 = null) {
  $__num = \func_num_args();
  $functorWithIndexCompose_1_0 = (($GLOBALS['Data_FunctorWithIndex_functorWithIndexCompose'] ?? \PhpursThunks::eval('Data_FunctorWithIndex_functorWithIndexCompose')))((($dictTraversableWithIndex_0)['FunctorWithIndex0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $foldableWithIndexCompose_2_1 = (($GLOBALS['Data_FoldableWithIndex_foldableWithIndexCompose'] ?? \PhpursThunks::eval('Data_FoldableWithIndex_foldableWithIndexCompose')))((($dictTraversableWithIndex_0)['FoldableWithIndex1'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $traversableCompose_3_2 = (($GLOBALS['Data_Traversable_traversableCompose'] ?? \PhpursThunks::eval('Data_Traversable_traversableCompose')))((($dictTraversableWithIndex_0)['Traversable2'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = function($dictTraversableWithIndex1_4 = null) use ($dictTraversableWithIndex_0, $foldableWithIndexCompose_2_1, $functorWithIndexCompose_1_0, $traversableCompose_3_2) {
  $__num = \func_num_args();
  $functorWithIndexCompose1_5_3 = ($functorWithIndexCompose_1_0)((($dictTraversableWithIndex1_4)['FunctorWithIndex0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $foldableWithIndexCompose1_6_4 = ($foldableWithIndexCompose_2_1)((($dictTraversableWithIndex1_4)['FoldableWithIndex1'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $traversableCompose1_7_5 = ($traversableCompose_3_2)((($dictTraversableWithIndex1_4)['Traversable2'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = ["traverseWithIndex" => function($dictApplicative_8 = null) use ($dictTraversableWithIndex1_4, $dictTraversableWithIndex_0) {
  $__num = \func_num_args();
  $traverseWithIndex3_9_6 = (($dictTraversableWithIndex_0)['traverseWithIndex'])($dictApplicative_8);
  $traverseWithIndex4_10_7 = (($dictTraversableWithIndex1_4)['traverseWithIndex'])($dictApplicative_8);
  $__res = (function() use ($dictApplicative_8, $traverseWithIndex3_9_6, $traverseWithIndex4_10_7) {
  $__fn = function($f_11 = null, $v_12 = null) use ($dictApplicative_8, $traverseWithIndex3_9_6, $traverseWithIndex4_10_7, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((((($dictApplicative_8)['Apply0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['map'])(($GLOBALS['Data_Functor_Compose_Compose'] ?? \PhpursThunks::eval('Data_Functor_Compose_Compose'))))((($traverseWithIndex3_9_6)(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))($traverseWithIndex4_10_7))((function() use ($f_11) {
  $__fn = function($a_13 = null, $b_14 = null) use ($f_11, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($f_11)(new Phpurs_Data2("Tuple", $a_13, $b_14));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})())))($v_12));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FunctorWithIndex0" => function($dollar__unused_8 = null) use ($functorWithIndexCompose1_5_3) {
  $__num = \func_num_args();
  $__res = $functorWithIndexCompose1_5_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($dollar__unused_8 = null) use ($foldableWithIndexCompose1_6_4) {
  $__num = \func_num_args();
  $__res = $foldableWithIndexCompose1_6_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($dollar__unused_8 = null) use ($traversableCompose1_7_5) {
  $__num = \func_num_args();
  $__res = $traversableCompose1_7_5;
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
}; return $v; };
\PhpursThunks::$thunks['Data_TraversableWithIndex_traversableWithIndexArray'] = function() { $v = ["traverseWithIndex" => function($dictApplicative_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_TraversableWithIndex_traverseWithIndexDefault'] ?? \PhpursThunks::eval('Data_TraversableWithIndex_traverseWithIndexDefault')))(($GLOBALS['Data_TraversableWithIndex_traversableWithIndexArray'] ?? \PhpursThunks::eval('Data_TraversableWithIndex_traversableWithIndexArray'))))($dictApplicative_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FunctorWithIndex0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_FunctorWithIndex_functorWithIndexArray'] ?? \PhpursThunks::eval('Data_FunctorWithIndex_functorWithIndexArray'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_FoldableWithIndex_foldableWithIndexArray'] ?? \PhpursThunks::eval('Data_FoldableWithIndex_foldableWithIndexArray'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Traversable_traversableArray'] ?? \PhpursThunks::eval('Data_Traversable_traversableArray'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_TraversableWithIndex_traversableWithIndexApp'] = function() { $v = function($dictTraversableWithIndex_0 = null) {
  $__num = \func_num_args();
  $functorWithIndexApp_1_0 = (($GLOBALS['Data_FunctorWithIndex_functorWithIndexApp'] ?? \PhpursThunks::eval('Data_FunctorWithIndex_functorWithIndexApp')))((($dictTraversableWithIndex_0)['FunctorWithIndex0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $foldableWithIndexApp_2_1 = (($GLOBALS['Data_FoldableWithIndex_foldableWithIndexApp'] ?? \PhpursThunks::eval('Data_FoldableWithIndex_foldableWithIndexApp')))((($dictTraversableWithIndex_0)['FoldableWithIndex1'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $traversableApp_3_2 = (($GLOBALS['Data_Traversable_traversableApp'] ?? \PhpursThunks::eval('Data_Traversable_traversableApp')))((($dictTraversableWithIndex_0)['Traversable2'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = ["traverseWithIndex" => function($dictApplicative_4 = null) use ($dictTraversableWithIndex_0) {
  $__num = \func_num_args();
  $traverseWithIndex2_5_3 = (($dictTraversableWithIndex_0)['traverseWithIndex'])($dictApplicative_4);
  $__res = (function() use ($dictApplicative_4, $traverseWithIndex2_5_3) {
  $__fn = function($f_6 = null, $v_7 = null) use ($dictApplicative_4, $traverseWithIndex2_5_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((((($dictApplicative_4)['Apply0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['map'])(($GLOBALS['Data_Functor_App_App'] ?? \PhpursThunks::eval('Data_Functor_App_App'))))((($traverseWithIndex2_5_3)($f_6))($v_7));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FunctorWithIndex0" => function($dollar__unused_4 = null) use ($functorWithIndexApp_1_0) {
  $__num = \func_num_args();
  $__res = $functorWithIndexApp_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($dollar__unused_4 = null) use ($foldableWithIndexApp_2_1) {
  $__num = \func_num_args();
  $__res = $foldableWithIndexApp_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($dollar__unused_4 = null) use ($traversableApp_3_2) {
  $__num = \func_num_args();
  $__res = $traversableApp_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_TraversableWithIndex_traversableWithIndexAdditive'] = function() { $v = ["traverseWithIndex" => (function() {
  $__fn = function($dictApplicative_0 = null, $f_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($GLOBALS['Data_Traversable_traversableAdditive'] ?? \PhpursThunks::eval('Data_Traversable_traversableAdditive')))['traverse'])($dictApplicative_0))(($f_1)(($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit'))));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "FunctorWithIndex0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_FunctorWithIndex_functorWithIndexAdditive'] ?? \PhpursThunks::eval('Data_FunctorWithIndex_functorWithIndexAdditive'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_FoldableWithIndex_foldableWithIndexAdditive'] ?? \PhpursThunks::eval('Data_FoldableWithIndex_foldableWithIndexAdditive'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Traversable_traversableAdditive'] ?? \PhpursThunks::eval('Data_Traversable_traversableAdditive'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_TraversableWithIndex_mapAccumRWithIndex'] = function() { $v = function($dictTraversableWithIndex_0 = null) {
  $__num = \func_num_args();
  $traverseWithIndex1_1_0 = (($dictTraversableWithIndex_0)['traverseWithIndex'])(($GLOBALS['Data_Traversable_Accum_Internal_applicativeStateR'] ?? \PhpursThunks::eval('Data_Traversable_Accum_Internal_applicativeStateR')));
  $__res = (function() use ($traverseWithIndex1_1_0) {
  $__fn = function($f_2 = null, $s0_3 = null, $xs_4 = null) use ($traverseWithIndex1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($traverseWithIndex1_1_0)((function() use ($f_2) {
  $__fn = function($i_5 = null, $a_6 = null, $s_7 = null) use ($f_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($f_2)($i_5))($s_7))($a_6);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()))($xs_4))($s0_3);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_TraversableWithIndex_scanrWithIndex'] = function() { $v = function($dictTraversableWithIndex_0 = null) {
  $__num = \func_num_args();
  $mapAccumRWithIndex1_1_0 = (($GLOBALS['Data_TraversableWithIndex_mapAccumRWithIndex'] ?? \PhpursThunks::eval('Data_TraversableWithIndex_mapAccumRWithIndex')))($dictTraversableWithIndex_0);
  $__res = (function() use ($mapAccumRWithIndex1_1_0) {
  $__fn = function($f_2 = null, $b0_3 = null, $xs_4 = null) use ($mapAccumRWithIndex1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($mapAccumRWithIndex1_1_0)((function() use ($f_2) {
  $__fn = function($i_5 = null, $b_6 = null, $a_7 = null) use ($f_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $b__prime___8_1 = ((($f_2)($i_5))($a_7))($b_6);
  $__res = ["accum" => $b__prime___8_1, "value" => $b__prime___8_1];
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()))($b0_3))($xs_4))['value'];
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_TraversableWithIndex_mapAccumLWithIndex'] = function() { $v = function($dictTraversableWithIndex_0 = null) {
  $__num = \func_num_args();
  $traverseWithIndex1_1_0 = (($dictTraversableWithIndex_0)['traverseWithIndex'])(($GLOBALS['Data_Traversable_Accum_Internal_applicativeStateL'] ?? \PhpursThunks::eval('Data_Traversable_Accum_Internal_applicativeStateL')));
  $__res = (function() use ($traverseWithIndex1_1_0) {
  $__fn = function($f_2 = null, $s0_3 = null, $xs_4 = null) use ($traverseWithIndex1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($traverseWithIndex1_1_0)((function() use ($f_2) {
  $__fn = function($i_5 = null, $a_6 = null, $s_7 = null) use ($f_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($f_2)($i_5))($s_7))($a_6);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()))($xs_4))($s0_3);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_TraversableWithIndex_scanlWithIndex'] = function() { $v = function($dictTraversableWithIndex_0 = null) {
  $__num = \func_num_args();
  $mapAccumLWithIndex1_1_0 = (($GLOBALS['Data_TraversableWithIndex_mapAccumLWithIndex'] ?? \PhpursThunks::eval('Data_TraversableWithIndex_mapAccumLWithIndex')))($dictTraversableWithIndex_0);
  $__res = (function() use ($mapAccumLWithIndex1_1_0) {
  $__fn = function($f_2 = null, $b0_3 = null, $xs_4 = null) use ($mapAccumLWithIndex1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($mapAccumLWithIndex1_1_0)((function() use ($f_2) {
  $__fn = function($i_5 = null, $b_6 = null, $a_7 = null) use ($f_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $b__prime___8_1 = ((($f_2)($i_5))($b_6))($a_7);
  $__res = ["accum" => $b__prime___8_1, "value" => $b__prime___8_1];
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()))($b0_3))($xs_4))['value'];
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_TraversableWithIndex_forWithIndex'] = function() { $v = (function() {
  $__fn = function($dictApplicative_0 = null, $dictTraversableWithIndex_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_2_0 = (($dictTraversableWithIndex_1)['traverseWithIndex'])($dictApplicative_0);
  $__res = (function() use ($__local_var_2_0) {
  $__fn = function($b_3 = null, $a_4 = null) use ($__local_var_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($__local_var_2_0)($a_4))($b_3);
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
})(); return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };



























