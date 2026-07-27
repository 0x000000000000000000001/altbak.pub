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

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };




// Data_TraversableWithIndex_traverseWithIndexDefault
function majData_majTraversablemajWithmajIndex_traversemajWithmajIndexmajDefault($dictTraversableWithIndex_0, $dictApplicative_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTraversablemajWithmajIndex_traversemajWithmajIndexmajDefault';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $sequence1_2_0 = (((($dictTraversableWithIndex_0)->{'Traversable2'})(null))->{'sequence'})($dictApplicative_1);
  $__res = function($f_3) use ($dictTraversableWithIndex_0, $sequence1_2_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($sequence1_2_0))((((($dictTraversableWithIndex_0)->{'FunctorWithIndex0'})(null))->{'mapWithIndex'})($f_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_TraversableWithIndex_traverseWithIndexDefault'] = __NAMESPACE__ . '\\majData_majTraversablemajWithmajIndex_traversemajWithmajIndexmajDefault';

// Data_TraversableWithIndex_traverseWithIndex
function majData_majTraversablemajWithmajIndex_traversemajWithmajIndex($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTraversablemajWithmajIndex_traversemajWithmajIndex';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'traverseWithIndex'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_TraversableWithIndex_traverseWithIndex'] = __NAMESPACE__ . '\\majData_majTraversablemajWithmajIndex_traversemajWithmajIndex';

// Data_TraversableWithIndex_traverseDefault
function majData_majTraversablemajWithmajIndex_traversemajDefault($dictTraversableWithIndex_0, $dictApplicative_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTraversablemajWithmajIndex_traversemajDefault';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $traverseWithIndex2_2_0 = (($dictTraversableWithIndex_0)->{'traverseWithIndex'})($dictApplicative_1);
  $__res = function($f_3) use ($traverseWithIndex2_2_0) {
  $__num = \func_num_args();
  $__res = ($traverseWithIndex2_2_0)(function($v_4) use ($f_3) {
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
$GLOBALS['Data_TraversableWithIndex_traverseDefault'] = __NAMESPACE__ . '\\majData_majTraversablemajWithmajIndex_traversemajDefault';

// Data_TraversableWithIndex_traversableWithIndexTuple
$GLOBALS['Data_TraversableWithIndex_traversableWithIndexTuple'] = (object)["traverseWithIndex" => (function() {
  $__fn = function($dictApplicative_0, $f_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})(($GLOBALS['Data_Tuple_Tuple'])(($v_2)->{'value0'})))((($f_1)($GLOBALS['Data_Unit_unit']))(($v_2)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "FunctorWithIndex0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FunctorWithIndex_functorWithIndexTuple'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FoldableWithIndex_foldableWithIndexTuple'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Traversable_traversableTuple'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_TraversableWithIndex_traversableWithIndexProduct
function majData_majTraversablemajWithmajIndex_traversablemajWithmajIndexmajProduct($dictTraversableWithIndex_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTraversablemajWithmajIndex_traversablemajWithmajIndexmajProduct';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $functorWithIndexProduct_1_0 = ($GLOBALS['Data_FunctorWithIndex_functorWithIndexProduct'])((($dictTraversableWithIndex_0)->{'FunctorWithIndex0'})(null));
  $foldableWithIndexProduct_2_1 = ($GLOBALS['Data_FoldableWithIndex_foldableWithIndexProduct'])((($dictTraversableWithIndex_0)->{'FoldableWithIndex1'})(null));
  $traversableProduct_3_2 = ($GLOBALS['Data_Traversable_traversableProduct'])((($dictTraversableWithIndex_0)->{'Traversable2'})(null));
  $__res = function($dictTraversableWithIndex1_4) use ($dictTraversableWithIndex_0, $foldableWithIndexProduct_2_1, $functorWithIndexProduct_1_0, $traversableProduct_3_2) {
  $__num = \func_num_args();
  $functorWithIndexProduct1_5_3 = ($functorWithIndexProduct_1_0)((($dictTraversableWithIndex1_4)->{'FunctorWithIndex0'})(null));
  $foldableWithIndexProduct1_6_4 = ($foldableWithIndexProduct_2_1)((($dictTraversableWithIndex1_4)->{'FoldableWithIndex1'})(null));
  $traversableProduct1_7_5 = ($traversableProduct_3_2)((($dictTraversableWithIndex1_4)->{'Traversable2'})(null));
  $__res = (object)["traverseWithIndex" => function($dictApplicative_8) use ($dictTraversableWithIndex1_4, $dictTraversableWithIndex_0) {
  $__num = \func_num_args();
  $__local_var_9_6 = (($dictApplicative_8)->{'Apply0'})(null);
  $traverseWithIndex3_10_7 = (($dictTraversableWithIndex_0)->{'traverseWithIndex'})($dictApplicative_8);
  $traverseWithIndex4_11_8 = (($dictTraversableWithIndex1_4)->{'traverseWithIndex'})($dictApplicative_8);
  $__res = (function() use ($__local_var_9_6, $traverseWithIndex3_10_7, $traverseWithIndex4_11_8) {
  $__fn = function($f_12, $v_13 = null) use ($__local_var_9_6, $traverseWithIndex3_10_7, $traverseWithIndex4_11_8, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_9_6)->{'apply'})(((((($__local_var_9_6)->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Functor_Product_product']))((($traverseWithIndex3_10_7)((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_12))($GLOBALS['Data_Either_Left'])))(($v_13)->{'value0'}))))((($traverseWithIndex4_11_8)((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_12))($GLOBALS['Data_Either_Right'])))(($v_13)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FunctorWithIndex0" => function($_dollar__unused_8) use ($functorWithIndexProduct1_5_3) {
  $__num = \func_num_args();
  $__res = $functorWithIndexProduct1_5_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($_dollar__unused_8) use ($foldableWithIndexProduct1_6_4) {
  $__num = \func_num_args();
  $__res = $foldableWithIndexProduct1_6_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($_dollar__unused_8) use ($traversableProduct1_7_5) {
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_TraversableWithIndex_traversableWithIndexProduct'] = __NAMESPACE__ . '\\majData_majTraversablemajWithmajIndex_traversablemajWithmajIndexmajProduct';

// Data_TraversableWithIndex_traversableWithIndexMultiplicative
$GLOBALS['Data_TraversableWithIndex_traversableWithIndexMultiplicative'] = (object)["traverseWithIndex" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $traverse8_1_0 = (($GLOBALS['Data_Traversable_traversableMultiplicative'])->{'traverse'})($dictApplicative_0);
  $__res = function($f_2) use ($traverse8_1_0) {
  $__num = \func_num_args();
  $__res = ($traverse8_1_0)(($f_2)($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FunctorWithIndex0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FunctorWithIndex_functorWithIndexMultiplicative'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FoldableWithIndex_foldableWithIndexMultiplicative'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Traversable_traversableMultiplicative'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_TraversableWithIndex_traversableWithIndexMaybe
$GLOBALS['Data_TraversableWithIndex_traversableWithIndexMaybe'] = (object)["traverseWithIndex" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $traverse8_1_0 = (($GLOBALS['Data_Traversable_traversableMaybe'])->{'traverse'})($dictApplicative_0);
  $__res = function($f_2) use ($traverse8_1_0) {
  $__num = \func_num_args();
  $__res = ($traverse8_1_0)(($f_2)($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FunctorWithIndex0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FunctorWithIndex_functorWithIndexMaybe'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FoldableWithIndex_foldableWithIndexMaybe'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Traversable_traversableMaybe'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_TraversableWithIndex_traversableWithIndexLast
$GLOBALS['Data_TraversableWithIndex_traversableWithIndexLast'] = (object)["traverseWithIndex" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $traverse8_1_0 = (($GLOBALS['Data_Traversable_traversableLast'])->{'traverse'})($dictApplicative_0);
  $__res = function($f_2) use ($traverse8_1_0) {
  $__num = \func_num_args();
  $__res = ($traverse8_1_0)(($f_2)($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FunctorWithIndex0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FunctorWithIndex_functorWithIndexLast'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FoldableWithIndex_foldableWithIndexLast'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Traversable_traversableLast'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_TraversableWithIndex_traversableWithIndexIdentity
$GLOBALS['Data_TraversableWithIndex_traversableWithIndexIdentity'] = (object)["traverseWithIndex" => (function() {
  $__fn = function($dictApplicative_0, $f_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Identity_Identity']))((($f_1)($GLOBALS['Data_Unit_unit']))($v_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "FunctorWithIndex0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FunctorWithIndex_functorWithIndexIdentity'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FoldableWithIndex_foldableWithIndexIdentity'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Traversable_traversableIdentity'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_TraversableWithIndex_traversableWithIndexFirst
$GLOBALS['Data_TraversableWithIndex_traversableWithIndexFirst'] = (object)["traverseWithIndex" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $traverse8_1_0 = (($GLOBALS['Data_Traversable_traversableFirst'])->{'traverse'})($dictApplicative_0);
  $__res = function($f_2) use ($traverse8_1_0) {
  $__num = \func_num_args();
  $__res = ($traverse8_1_0)(($f_2)($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FunctorWithIndex0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FunctorWithIndex_functorWithIndexFirst'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FoldableWithIndex_foldableWithIndexFirst'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Traversable_traversableFirst'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_TraversableWithIndex_traversableWithIndexEither
$GLOBALS['Data_TraversableWithIndex_traversableWithIndexEither'] = (object)["traverseWithIndex" => (function() {
  $__fn = function($dictApplicative_0, $v_1 = null, $v1_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if ($v1_2 instanceof \Data\Either\Data_Either_Left) {
$__t0 = (($dictApplicative_0)->{'pure'})(new \Data\Either\Data_Either_Left(($v1_2)->{'value0'}));
goto end_branch_0;;
};
  if ($v1_2 instanceof \Data\Either\Data_Either_Right) {
$__t0 = ((((((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Either_Right']))((($v_1)($GLOBALS['Data_Unit_unit']))(($v1_2)->{'value0'}));
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
})(), "FunctorWithIndex0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FunctorWithIndex_functorWithIndexEither'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FoldableWithIndex_foldableWithIndexEither'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Traversable_traversableEither'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_TraversableWithIndex_traversableWithIndexDual
$GLOBALS['Data_TraversableWithIndex_traversableWithIndexDual'] = (object)["traverseWithIndex" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $traverse8_1_0 = (($GLOBALS['Data_Traversable_traversableDual'])->{'traverse'})($dictApplicative_0);
  $__res = function($f_2) use ($traverse8_1_0) {
  $__num = \func_num_args();
  $__res = ($traverse8_1_0)(($f_2)($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FunctorWithIndex0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FunctorWithIndex_functorWithIndexDual'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FoldableWithIndex_foldableWithIndexDual'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Traversable_traversableDual'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_TraversableWithIndex_traversableWithIndexDisj
$GLOBALS['Data_TraversableWithIndex_traversableWithIndexDisj'] = (object)["traverseWithIndex" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $traverse8_1_0 = (($GLOBALS['Data_Traversable_traversableDisj'])->{'traverse'})($dictApplicative_0);
  $__res = function($f_2) use ($traverse8_1_0) {
  $__num = \func_num_args();
  $__res = ($traverse8_1_0)(($f_2)($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FunctorWithIndex0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FunctorWithIndex_functorWithIndexDisj'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FoldableWithIndex_foldableWithIndexDisj'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Traversable_traversableDisj'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_TraversableWithIndex_traversableWithIndexCoproduct
function majData_majTraversablemajWithmajIndex_traversablemajWithmajIndexmajCoproduct($dictTraversableWithIndex_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTraversablemajWithmajIndex_traversablemajWithmajIndexmajCoproduct';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $functorWithIndexCoproduct_1_0 = ($GLOBALS['Data_FunctorWithIndex_functorWithIndexCoproduct'])((($dictTraversableWithIndex_0)->{'FunctorWithIndex0'})(null));
  $foldableWithIndexCoproduct_2_1 = ($GLOBALS['Data_FoldableWithIndex_foldableWithIndexCoproduct'])((($dictTraversableWithIndex_0)->{'FoldableWithIndex1'})(null));
  $traversableCoproduct_3_2 = ($GLOBALS['Data_Traversable_traversableCoproduct'])((($dictTraversableWithIndex_0)->{'Traversable2'})(null));
  $__res = function($dictTraversableWithIndex1_4) use ($dictTraversableWithIndex_0, $foldableWithIndexCoproduct_2_1, $functorWithIndexCoproduct_1_0, $traversableCoproduct_3_2) {
  $__num = \func_num_args();
  $functorWithIndexCoproduct1_5_3 = ($functorWithIndexCoproduct_1_0)((($dictTraversableWithIndex1_4)->{'FunctorWithIndex0'})(null));
  $foldableWithIndexCoproduct1_6_4 = ($foldableWithIndexCoproduct_2_1)((($dictTraversableWithIndex1_4)->{'FoldableWithIndex1'})(null));
  $traversableCoproduct1_7_5 = ($traversableCoproduct_3_2)((($dictTraversableWithIndex1_4)->{'Traversable2'})(null));
  $__res = (object)["traverseWithIndex" => function($dictApplicative_8) use ($dictTraversableWithIndex1_4, $dictTraversableWithIndex_0) {
  $__num = \func_num_args();
  $__local_var_9_6 = (((($dictApplicative_8)->{'Apply0'})(null))->{'Functor0'})(null);
  $traverseWithIndex3_10_7 = (($dictTraversableWithIndex_0)->{'traverseWithIndex'})($dictApplicative_8);
  $traverseWithIndex4_11_8 = (($dictTraversableWithIndex1_4)->{'traverseWithIndex'})($dictApplicative_8);
  $__res = function($f_12) use ($__local_var_9_6, $traverseWithIndex3_10_7, $traverseWithIndex4_11_8) {
  $__num = \func_num_args();
  $__local_var_13_9 = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($__local_var_9_6)->{'map'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Functor_Coproduct_Coproduct']))($GLOBALS['Data_Either_Left']))))(($traverseWithIndex3_10_7)((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_12))($GLOBALS['Data_Either_Left'])));
  $__local_var_14_10 = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($__local_var_9_6)->{'map'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Functor_Coproduct_Coproduct']))($GLOBALS['Data_Either_Right']))))(($traverseWithIndex4_11_8)((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_12))($GLOBALS['Data_Either_Right'])));
  $__res = function($v2_15) use ($__local_var_13_9, $__local_var_14_10) {
  $__num = \func_num_args();
  $__t11 = null;;
  if ($v2_15 instanceof \Data\Either\Data_Either_Left) {
$__t11 = ($__local_var_13_9)(($v2_15)->{'value0'});
goto end_branch_11;;
};
  if ($v2_15 instanceof \Data\Either\Data_Either_Right) {
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
}, "FunctorWithIndex0" => function($_dollar__unused_8) use ($functorWithIndexCoproduct1_5_3) {
  $__num = \func_num_args();
  $__res = $functorWithIndexCoproduct1_5_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($_dollar__unused_8) use ($foldableWithIndexCoproduct1_6_4) {
  $__num = \func_num_args();
  $__res = $foldableWithIndexCoproduct1_6_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($_dollar__unused_8) use ($traversableCoproduct1_7_5) {
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_TraversableWithIndex_traversableWithIndexCoproduct'] = __NAMESPACE__ . '\\majData_majTraversablemajWithmajIndex_traversablemajWithmajIndexmajCoproduct';

// Data_TraversableWithIndex_traversableWithIndexConst
$GLOBALS['Data_TraversableWithIndex_traversableWithIndexConst'] = (object)["traverseWithIndex" => (function() {
  $__fn = function($dictApplicative_0, $v_1 = null, $v1_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($dictApplicative_0)->{'pure'})($v1_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "FunctorWithIndex0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FunctorWithIndex_functorWithIndexConst'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FoldableWithIndex_foldableWithIndexConst'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Traversable_traversableConst'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_TraversableWithIndex_traversableWithIndexConj
$GLOBALS['Data_TraversableWithIndex_traversableWithIndexConj'] = (object)["traverseWithIndex" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $traverse8_1_0 = (($GLOBALS['Data_Traversable_traversableConj'])->{'traverse'})($dictApplicative_0);
  $__res = function($f_2) use ($traverse8_1_0) {
  $__num = \func_num_args();
  $__res = ($traverse8_1_0)(($f_2)($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FunctorWithIndex0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FunctorWithIndex_functorWithIndexConj'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FoldableWithIndex_foldableWithIndexConj'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Traversable_traversableConj'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_TraversableWithIndex_traversableWithIndexCompose
function majData_majTraversablemajWithmajIndex_traversablemajWithmajIndexmajCompose($dictTraversableWithIndex_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTraversablemajWithmajIndex_traversablemajWithmajIndexmajCompose';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $functorWithIndexCompose_1_0 = ($GLOBALS['Data_FunctorWithIndex_functorWithIndexCompose'])((($dictTraversableWithIndex_0)->{'FunctorWithIndex0'})(null));
  $foldableWithIndexCompose_2_1 = ($GLOBALS['Data_FoldableWithIndex_foldableWithIndexCompose'])((($dictTraversableWithIndex_0)->{'FoldableWithIndex1'})(null));
  $traversableCompose_3_2 = ($GLOBALS['Data_Traversable_traversableCompose'])((($dictTraversableWithIndex_0)->{'Traversable2'})(null));
  $__res = function($dictTraversableWithIndex1_4) use ($dictTraversableWithIndex_0, $foldableWithIndexCompose_2_1, $functorWithIndexCompose_1_0, $traversableCompose_3_2) {
  $__num = \func_num_args();
  $functorWithIndexCompose1_5_3 = ($functorWithIndexCompose_1_0)((($dictTraversableWithIndex1_4)->{'FunctorWithIndex0'})(null));
  $foldableWithIndexCompose1_6_4 = ($foldableWithIndexCompose_2_1)((($dictTraversableWithIndex1_4)->{'FoldableWithIndex1'})(null));
  $traversableCompose1_7_5 = ($traversableCompose_3_2)((($dictTraversableWithIndex1_4)->{'Traversable2'})(null));
  $__res = (object)["traverseWithIndex" => function($dictApplicative_8) use ($dictTraversableWithIndex1_4, $dictTraversableWithIndex_0) {
  $__num = \func_num_args();
  $traverseWithIndex3_9_6 = (($dictTraversableWithIndex_0)->{'traverseWithIndex'})($dictApplicative_8);
  $traverseWithIndex4_10_7 = (($dictTraversableWithIndex1_4)->{'traverseWithIndex'})($dictApplicative_8);
  $__res = (function() use ($dictApplicative_8, $traverseWithIndex3_9_6, $traverseWithIndex4_10_7) {
  $__fn = function($f_11, $v_12 = null) use ($dictApplicative_8, $traverseWithIndex3_9_6, $traverseWithIndex4_10_7, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((((($dictApplicative_8)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Functor_Compose_Compose']))((($traverseWithIndex3_9_6)((($GLOBALS['Control_Semigroupoid_composeImpl'])($traverseWithIndex4_10_7))((function() use ($f_11) {
  $__fn = function($a_13, $b_14 = null) use ($f_11, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($f_11)(new \Data\Tuple\Data_Tuple_Tuple($a_13, $b_14));
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
}, "FunctorWithIndex0" => function($_dollar__unused_8) use ($functorWithIndexCompose1_5_3) {
  $__num = \func_num_args();
  $__res = $functorWithIndexCompose1_5_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($_dollar__unused_8) use ($foldableWithIndexCompose1_6_4) {
  $__num = \func_num_args();
  $__res = $foldableWithIndexCompose1_6_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($_dollar__unused_8) use ($traversableCompose1_7_5) {
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_TraversableWithIndex_traversableWithIndexCompose'] = __NAMESPACE__ . '\\majData_majTraversablemajWithmajIndex_traversablemajWithmajIndexmajCompose';

// Data_TraversableWithIndex_traversableWithIndexArray
$GLOBALS['Data_TraversableWithIndex_traversableWithIndexArray'] = (object)["traverseWithIndex" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_TraversableWithIndex_traverseWithIndexDefault'])($GLOBALS['Data_TraversableWithIndex_traversableWithIndexArray']))($dictApplicative_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FunctorWithIndex0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FunctorWithIndex_functorWithIndexArray'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FoldableWithIndex_foldableWithIndexArray'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Traversable_traversableArray'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_TraversableWithIndex_traversableWithIndexApp
function majData_majTraversablemajWithmajIndex_traversablemajWithmajIndexmajApp($dictTraversableWithIndex_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTraversablemajWithmajIndex_traversablemajWithmajIndexmajApp';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictTraversableWithIndex_0)->{'FunctorWithIndex0'})(null);
  $__local_var_2_1 = (($__local_var_1_0)->{'Functor0'})(null);
  $functorWithIndexApp_3_2 = (object)["mapWithIndex" => (function() use ($__local_var_1_0) {
  $__fn = function($f_3, $v_4 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_1_0)->{'mapWithIndex'})($f_3))($v_4);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_3) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = $__local_var_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_4_3 = (($dictTraversableWithIndex_0)->{'FoldableWithIndex1'})(null);
  $__local_var_5_4 = (($__local_var_4_3)->{'Foldable0'})(null);
  $foldableApp_6_5 = (object)["foldr" => (function() use ($__local_var_5_4) {
  $__fn = function($f_6, $i_7 = null, $v_8 = null) use ($__local_var_5_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($__local_var_5_4)->{'foldr'})($f_6))($i_7))($v_8);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldl" => (function() use ($__local_var_5_4) {
  $__fn = function($f_6, $i_7 = null, $v_8 = null) use ($__local_var_5_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($__local_var_5_4)->{'foldl'})($f_6))($i_7))($v_8);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldMap" => function($dictMonoid_6) use ($__local_var_5_4) {
  $__num = \func_num_args();
  $__res = (($__local_var_5_4)->{'foldMap'})($dictMonoid_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $foldableWithIndexApp_6_5 = (object)["foldrWithIndex" => (function() use ($__local_var_4_3) {
  $__fn = function($f_7, $z_8 = null, $v_9 = null) use ($__local_var_4_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($__local_var_4_3)->{'foldrWithIndex'})($f_7))($z_8))($v_9);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldlWithIndex" => (function() use ($__local_var_4_3) {
  $__fn = function($f_7, $z_8 = null, $v_9 = null) use ($__local_var_4_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($__local_var_4_3)->{'foldlWithIndex'})($f_7))($z_8))($v_9);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldMapWithIndex" => function($dictMonoid_7) use ($__local_var_4_3) {
  $__num = \func_num_args();
  $__res = (($__local_var_4_3)->{'foldMapWithIndex'})($dictMonoid_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable0" => function($_dollar__unused_7) use ($foldableApp_6_5) {
  $__num = \func_num_args();
  $__res = $foldableApp_6_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $traversableApp_7_7 = \Data\Traversable\majData_majTraversable_traversablemajApp((($dictTraversableWithIndex_0)->{'Traversable2'})(null));
  $__res = (object)["traverseWithIndex" => function($dictApplicative_8) use ($dictTraversableWithIndex_0) {
  $__num = \func_num_args();
  $traverseWithIndex2_9_8 = (($dictTraversableWithIndex_0)->{'traverseWithIndex'})($dictApplicative_8);
  $__res = (function() use ($dictApplicative_8, $traverseWithIndex2_9_8) {
  $__fn = function($f_10, $v_11 = null) use ($dictApplicative_8, $traverseWithIndex2_9_8, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((((($dictApplicative_8)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Functor_App_App']))((($traverseWithIndex2_9_8)($f_10))($v_11));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FunctorWithIndex0" => function($_dollar__unused_8) use ($functorWithIndexApp_3_2) {
  $__num = \func_num_args();
  $__res = $functorWithIndexApp_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($_dollar__unused_8) use ($foldableWithIndexApp_6_5) {
  $__num = \func_num_args();
  $__res = $foldableWithIndexApp_6_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($_dollar__unused_8) use ($traversableApp_7_7) {
  $__num = \func_num_args();
  $__res = $traversableApp_7_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_TraversableWithIndex_traversableWithIndexApp'] = __NAMESPACE__ . '\\majData_majTraversablemajWithmajIndex_traversablemajWithmajIndexmajApp';

// Data_TraversableWithIndex_traversableWithIndexAdditive
$GLOBALS['Data_TraversableWithIndex_traversableWithIndexAdditive'] = (object)["traverseWithIndex" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $traverse8_1_0 = (($GLOBALS['Data_Traversable_traversableAdditive'])->{'traverse'})($dictApplicative_0);
  $__res = function($f_2) use ($traverse8_1_0) {
  $__num = \func_num_args();
  $__res = ($traverse8_1_0)(($f_2)($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FunctorWithIndex0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FunctorWithIndex_functorWithIndexAdditive'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FoldableWithIndex_foldableWithIndexAdditive'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Traversable_traversableAdditive'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_TraversableWithIndex_mapAccumRWithIndex
function majData_majTraversablemajWithmajIndex_mapmajAccummajRmajWithmajIndex($dictTraversableWithIndex_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTraversablemajWithmajIndex_mapmajAccummajRmajWithmajIndex';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $traverseWithIndex1_1_0 = (($dictTraversableWithIndex_0)->{'traverseWithIndex'})($GLOBALS['Data_Traversable_Accum_Internal_applicativeStateR']);
  $__res = (function() use ($traverseWithIndex1_1_0) {
  $__fn = function($f_2, $s0_3 = null, $xs_4 = null) use ($traverseWithIndex1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($traverseWithIndex1_1_0)((function() use ($f_2) {
  $__fn = function($i_5, $a_6 = null, $s_7 = null) use ($f_2, &$__fn) {
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_TraversableWithIndex_mapAccumRWithIndex'] = __NAMESPACE__ . '\\majData_majTraversablemajWithmajIndex_mapmajAccummajRmajWithmajIndex';

// Data_TraversableWithIndex_scanrWithIndex
function majData_majTraversablemajWithmajIndex_scanrmajWithmajIndex($dictTraversableWithIndex_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTraversablemajWithmajIndex_scanrmajWithmajIndex';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $mapAccumRWithIndex1_1_0 = ($GLOBALS['Data_TraversableWithIndex_mapAccumRWithIndex'])($dictTraversableWithIndex_0);
  $__res = (function() use ($mapAccumRWithIndex1_1_0) {
  $__fn = function($f_2, $b0_3 = null, $xs_4 = null) use ($mapAccumRWithIndex1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($mapAccumRWithIndex1_1_0)((function() use ($f_2) {
  $__fn = function($i_5, $b_6 = null, $a_7 = null) use ($f_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $b_prime_8_1 = ((($f_2)($i_5))($a_7))($b_6);
  $__res = (object)["accum" => $b_prime_8_1, "value" => $b_prime_8_1];
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()))($b0_3))($xs_4))->{'value'};
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
$GLOBALS['Data_TraversableWithIndex_scanrWithIndex'] = __NAMESPACE__ . '\\majData_majTraversablemajWithmajIndex_scanrmajWithmajIndex';

// Data_TraversableWithIndex_mapAccumLWithIndex
function majData_majTraversablemajWithmajIndex_mapmajAccummajLmajWithmajIndex($dictTraversableWithIndex_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTraversablemajWithmajIndex_mapmajAccummajLmajWithmajIndex';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $traverseWithIndex1_1_0 = (($dictTraversableWithIndex_0)->{'traverseWithIndex'})($GLOBALS['Data_Traversable_Accum_Internal_applicativeStateL']);
  $__res = (function() use ($traverseWithIndex1_1_0) {
  $__fn = function($f_2, $s0_3 = null, $xs_4 = null) use ($traverseWithIndex1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($traverseWithIndex1_1_0)((function() use ($f_2) {
  $__fn = function($i_5, $a_6 = null, $s_7 = null) use ($f_2, &$__fn) {
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_TraversableWithIndex_mapAccumLWithIndex'] = __NAMESPACE__ . '\\majData_majTraversablemajWithmajIndex_mapmajAccummajLmajWithmajIndex';

// Data_TraversableWithIndex_scanlWithIndex
function majData_majTraversablemajWithmajIndex_scanlmajWithmajIndex($dictTraversableWithIndex_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTraversablemajWithmajIndex_scanlmajWithmajIndex';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $mapAccumLWithIndex1_1_0 = ($GLOBALS['Data_TraversableWithIndex_mapAccumLWithIndex'])($dictTraversableWithIndex_0);
  $__res = (function() use ($mapAccumLWithIndex1_1_0) {
  $__fn = function($f_2, $b0_3 = null, $xs_4 = null) use ($mapAccumLWithIndex1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($mapAccumLWithIndex1_1_0)((function() use ($f_2) {
  $__fn = function($i_5, $b_6 = null, $a_7 = null) use ($f_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $b_prime_8_1 = ((($f_2)($i_5))($b_6))($a_7);
  $__res = (object)["accum" => $b_prime_8_1, "value" => $b_prime_8_1];
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()))($b0_3))($xs_4))->{'value'};
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
$GLOBALS['Data_TraversableWithIndex_scanlWithIndex'] = __NAMESPACE__ . '\\majData_majTraversablemajWithmajIndex_scanlmajWithmajIndex';

// Data_TraversableWithIndex_forWithIndex
function majData_majTraversablemajWithmajIndex_formajWithmajIndex($dictApplicative_0, $dictTraversableWithIndex_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTraversablemajWithmajIndex_formajWithmajIndex';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_2_0 = (($dictTraversableWithIndex_1)->{'traverseWithIndex'})($dictApplicative_0);
  $__res = (function() use ($__local_var_2_0) {
  $__fn = function($b_3, $a_4 = null) use ($__local_var_2_0, &$__fn) {
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
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_TraversableWithIndex_forWithIndex'] = __NAMESPACE__ . '\\majData_majTraversablemajWithmajIndex_formajWithmajIndex';

