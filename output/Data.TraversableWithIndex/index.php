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
if (!\function_exists(__NAMESPACE__ . '\\phpurs_execute_effect')) {
  function phpurs_execute_effect($val) {
    if (\is_callable($val)) {
      return $val($GLOBALS['Data_Unit_unit']);
    }
    return $val;
  }
}

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };




// Data_TraversableWithIndex_traverseWithIndexDefault
function majData_majTraversablemajWithmajIndex_traversemajWithmajIndexmajDefault($dictTraversableWithIndex_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTraversablemajWithmajIndex_traversemajWithmajIndexmajDefault';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $FunctorWithIndex0_1_0 = (($dictTraversableWithIndex_0)->{'FunctorWithIndex0'})(null);
  $__res = function($dictApplicative_2) use ($FunctorWithIndex0_1_0, $dictTraversableWithIndex_0) {
  $__num = \func_num_args();
  $sequence1_3_1 = (((($dictTraversableWithIndex_0)->{'Traversable2'})(null))->{'sequence'})($dictApplicative_2);
  $__res = function($f_4) use ($FunctorWithIndex0_1_0, $sequence1_3_1) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($sequence1_3_1))((($FunctorWithIndex0_1_0)->{'mapWithIndex'})($f_4));
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
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
function majData_majTraversablemajWithmajIndex_traversemajDefault($dictTraversableWithIndex_0, $dictApplicative_1 = null, $f_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTraversablemajWithmajIndex_traversemajDefault';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictTraversableWithIndex_0)->{'traverseWithIndex'})($dictApplicative_1))(function($v_3) use ($f_2) {
  $__num = \func_num_args();
  $__res = $f_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_TraversableWithIndex_traverseDefault'] = __NAMESPACE__ . '\\majData_majTraversablemajWithmajIndex_traversemajDefault';

// Data_TraversableWithIndex_traversableWithIndexTuple
$GLOBALS['Data_TraversableWithIndex_traversableWithIndexTuple'] = (object)["traverseWithIndex" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $Functor0_1_0 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($f_2) use ($Functor0_1_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($Functor0_1_0, $f_2) {
  $__num = \func_num_args();
  $__res = ((($Functor0_1_0)->{'map'})(($GLOBALS['Data_Tuple_Tuple'])(($v_3)->{'value0'})))((($f_2)($GLOBALS['Data_Unit_unit']))(($v_3)->{'value1'}));
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
}, "FunctorWithIndex0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FunctorWithIndex_functorWithIndexTuple'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FoldableWithIndex_foldableWithIndexTuple'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($_dollar___unused_0) {
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
  $__local_var_1_0 = (($dictTraversableWithIndex_0)->{'FunctorWithIndex0'})(null);
  $__local_var_2_1 = (($__local_var_1_0)->{'Functor0'})(null);
  $__local_var_3_2 = (($dictTraversableWithIndex_0)->{'FoldableWithIndex1'})(null);
  $__local_var_4_3 = (($__local_var_3_2)->{'Foldable0'})(null);
  $__local_var_5_4 = (($dictTraversableWithIndex_0)->{'Traversable2'})(null);
  $__local_var_6_5 = (($__local_var_5_4)->{'Functor0'})(null);
  $__local_var_7_6 = (($__local_var_5_4)->{'Foldable1'})(null);
  $__res = function($dictTraversableWithIndex1_8) use ($__local_var_1_0, $__local_var_2_1, $__local_var_3_2, $__local_var_4_3, $__local_var_5_4, $__local_var_6_5, $__local_var_7_6, $dictTraversableWithIndex_0) {
  $__num = \func_num_args();
  $__local_var_9_7 = (($dictTraversableWithIndex1_8)->{'FunctorWithIndex0'})(null);
  $__local_var_10_8 = (($__local_var_9_7)->{'Functor0'})(null);
  $functorProduct1_10_8 = (object)["map" => function($f_11) use ($__local_var_10_8, $__local_var_2_1) {
  $__num = \func_num_args();
  $__res = function($v_12) use ($__local_var_10_8, $__local_var_2_1, $f_11) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(((($__local_var_2_1)->{'map'})($f_11))(($v_12)->{'value0'}), ((($__local_var_10_8)->{'map'})($f_11))(($v_12)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $functorWithIndexProduct1_9_7 = (object)["mapWithIndex" => function($f_11) use ($__local_var_1_0, $__local_var_9_7) {
  $__num = \func_num_args();
  $__res = function($v_12) use ($__local_var_1_0, $__local_var_9_7, $f_11) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(((($__local_var_1_0)->{'mapWithIndex'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_11))($GLOBALS['Data_Either_Left'])))(($v_12)->{'value0'}), ((($__local_var_9_7)->{'mapWithIndex'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_11))($GLOBALS['Data_Either_Right'])))(($v_12)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_11) use ($functorProduct1_10_8) {
  $__num = \func_num_args();
  $__res = $functorProduct1_10_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_10_11 = (($dictTraversableWithIndex1_8)->{'FoldableWithIndex1'})(null);
  $__local_var_11_12 = (($__local_var_10_11)->{'Foldable0'})(null);
  $foldableProduct1_11_12 = (object)["foldr" => function($f_12) use ($__local_var_11_12, $__local_var_4_3) {
  $__num = \func_num_args();
  $__res = function($z_13) use ($__local_var_11_12, $__local_var_4_3, $f_12) {
  $__num = \func_num_args();
  $__res = function($v_14) use ($__local_var_11_12, $__local_var_4_3, $f_12, $z_13) {
  $__num = \func_num_args();
  $__res = (((($__local_var_4_3)->{'foldr'})($f_12))((((($__local_var_11_12)->{'foldr'})($f_12))($z_13))(($v_14)->{'value1'})))(($v_14)->{'value0'});
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
}, "foldl" => function($f_12) use ($__local_var_11_12, $__local_var_4_3) {
  $__num = \func_num_args();
  $__res = function($z_13) use ($__local_var_11_12, $__local_var_4_3, $f_12) {
  $__num = \func_num_args();
  $__res = function($v_14) use ($__local_var_11_12, $__local_var_4_3, $f_12, $z_13) {
  $__num = \func_num_args();
  $__res = (((($__local_var_11_12)->{'foldl'})($f_12))((((($__local_var_4_3)->{'foldl'})($f_12))($z_13))(($v_14)->{'value0'})))(($v_14)->{'value1'});
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
}, "foldMap" => function($dictMonoid_12) use ($__local_var_11_12, $__local_var_4_3) {
  $__num = \func_num_args();
  $Semigroup0_13_13 = (($dictMonoid_12)->{'Semigroup0'})(null);
  $__res = function($f_14) use ($Semigroup0_13_13, $__local_var_11_12, $__local_var_4_3, $dictMonoid_12) {
  $__num = \func_num_args();
  $__res = function($v_15) use ($Semigroup0_13_13, $__local_var_11_12, $__local_var_4_3, $dictMonoid_12, $f_14) {
  $__num = \func_num_args();
  $__res = ((($Semigroup0_13_13)->{'append'})((((($__local_var_4_3)->{'foldMap'})($dictMonoid_12))($f_14))(($v_15)->{'value0'})))((((($__local_var_11_12)->{'foldMap'})($dictMonoid_12))($f_14))(($v_15)->{'value1'}));
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
  $foldableWithIndexProduct1_10_11 = (object)["foldrWithIndex" => function($f_12) use ($__local_var_10_11, $__local_var_3_2) {
  $__num = \func_num_args();
  $__res = function($z_13) use ($__local_var_10_11, $__local_var_3_2, $f_12) {
  $__num = \func_num_args();
  $__res = function($v_14) use ($__local_var_10_11, $__local_var_3_2, $f_12, $z_13) {
  $__num = \func_num_args();
  $__res = (((($__local_var_3_2)->{'foldrWithIndex'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_12))($GLOBALS['Data_Either_Left'])))((((($__local_var_10_11)->{'foldrWithIndex'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_12))($GLOBALS['Data_Either_Right'])))($z_13))(($v_14)->{'value1'})))(($v_14)->{'value0'});
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
}, "foldlWithIndex" => function($f_12) use ($__local_var_10_11, $__local_var_3_2) {
  $__num = \func_num_args();
  $__res = function($z_13) use ($__local_var_10_11, $__local_var_3_2, $f_12) {
  $__num = \func_num_args();
  $__res = function($v_14) use ($__local_var_10_11, $__local_var_3_2, $f_12, $z_13) {
  $__num = \func_num_args();
  $__res = (((($__local_var_10_11)->{'foldlWithIndex'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_12))($GLOBALS['Data_Either_Right'])))((((($__local_var_3_2)->{'foldlWithIndex'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_12))($GLOBALS['Data_Either_Left'])))($z_13))(($v_14)->{'value0'})))(($v_14)->{'value1'});
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
}, "foldMapWithIndex" => function($dictMonoid_12) use ($__local_var_10_11, $__local_var_3_2) {
  $__num = \func_num_args();
  $Semigroup0_13_15 = (($dictMonoid_12)->{'Semigroup0'})(null);
  $__res = function($f_14) use ($Semigroup0_13_15, $__local_var_10_11, $__local_var_3_2, $dictMonoid_12) {
  $__num = \func_num_args();
  $__res = function($v_15) use ($Semigroup0_13_15, $__local_var_10_11, $__local_var_3_2, $dictMonoid_12, $f_14) {
  $__num = \func_num_args();
  $__res = ((($Semigroup0_13_15)->{'append'})((((($__local_var_3_2)->{'foldMapWithIndex'})($dictMonoid_12))((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_14))($GLOBALS['Data_Either_Left'])))(($v_15)->{'value0'})))((((($__local_var_10_11)->{'foldMapWithIndex'})($dictMonoid_12))((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_14))($GLOBALS['Data_Either_Right'])))(($v_15)->{'value1'}));
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
}, "Foldable0" => function($_dollar___unused_12) use ($foldableProduct1_11_12) {
  $__num = \func_num_args();
  $__res = $foldableProduct1_11_12;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_11_17 = (($dictTraversableWithIndex1_8)->{'Traversable2'})(null);
  $__local_var_12_18 = (($__local_var_11_17)->{'Functor0'})(null);
  $functorProduct1_12_18 = (object)["map" => function($f_13) use ($__local_var_12_18, $__local_var_6_5) {
  $__num = \func_num_args();
  $__res = function($v_14) use ($__local_var_12_18, $__local_var_6_5, $f_13) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(((($__local_var_6_5)->{'map'})($f_13))(($v_14)->{'value0'}), ((($__local_var_12_18)->{'map'})($f_13))(($v_14)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_13_20 = (($__local_var_11_17)->{'Foldable1'})(null);
  $foldableProduct1_13_20 = (object)["foldr" => function($f_14) use ($__local_var_13_20, $__local_var_7_6) {
  $__num = \func_num_args();
  $__res = function($z_15) use ($__local_var_13_20, $__local_var_7_6, $f_14) {
  $__num = \func_num_args();
  $__res = function($v_16) use ($__local_var_13_20, $__local_var_7_6, $f_14, $z_15) {
  $__num = \func_num_args();
  $__res = (((($__local_var_7_6)->{'foldr'})($f_14))((((($__local_var_13_20)->{'foldr'})($f_14))($z_15))(($v_16)->{'value1'})))(($v_16)->{'value0'});
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
}, "foldl" => function($f_14) use ($__local_var_13_20, $__local_var_7_6) {
  $__num = \func_num_args();
  $__res = function($z_15) use ($__local_var_13_20, $__local_var_7_6, $f_14) {
  $__num = \func_num_args();
  $__res = function($v_16) use ($__local_var_13_20, $__local_var_7_6, $f_14, $z_15) {
  $__num = \func_num_args();
  $__res = (((($__local_var_13_20)->{'foldl'})($f_14))((((($__local_var_7_6)->{'foldl'})($f_14))($z_15))(($v_16)->{'value0'})))(($v_16)->{'value1'});
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
}, "foldMap" => function($dictMonoid_14) use ($__local_var_13_20, $__local_var_7_6) {
  $__num = \func_num_args();
  $Semigroup0_15_21 = (($dictMonoid_14)->{'Semigroup0'})(null);
  $__res = function($f_16) use ($Semigroup0_15_21, $__local_var_13_20, $__local_var_7_6, $dictMonoid_14) {
  $__num = \func_num_args();
  $__res = function($v_17) use ($Semigroup0_15_21, $__local_var_13_20, $__local_var_7_6, $dictMonoid_14, $f_16) {
  $__num = \func_num_args();
  $__res = ((($Semigroup0_15_21)->{'append'})((((($__local_var_7_6)->{'foldMap'})($dictMonoid_14))($f_16))(($v_17)->{'value0'})))((((($__local_var_13_20)->{'foldMap'})($dictMonoid_14))($f_16))(($v_17)->{'value1'}));
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
  $traversableProduct1_11_17 = (object)["traverse" => function($dictApplicative_14) use ($__local_var_11_17, $__local_var_5_4) {
  $__num = \func_num_args();
  $Apply0_15_23 = (($dictApplicative_14)->{'Apply0'})(null);
  $__res = function($f_16) use ($Apply0_15_23, $__local_var_11_17, $__local_var_5_4, $dictApplicative_14) {
  $__num = \func_num_args();
  $__res = function($v_17) use ($Apply0_15_23, $__local_var_11_17, $__local_var_5_4, $dictApplicative_14, $f_16) {
  $__num = \func_num_args();
  $__res = ((($Apply0_15_23)->{'apply'})(((((($Apply0_15_23)->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Functor_Product_product']))((((($__local_var_5_4)->{'traverse'})($dictApplicative_14))($f_16))(($v_17)->{'value0'}))))((((($__local_var_11_17)->{'traverse'})($dictApplicative_14))($f_16))(($v_17)->{'value1'}));
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
}, "sequence" => function($dictApplicative_14) use ($__local_var_11_17, $__local_var_5_4) {
  $__num = \func_num_args();
  $Apply0_15_24 = (($dictApplicative_14)->{'Apply0'})(null);
  $__res = function($v_16) use ($Apply0_15_24, $__local_var_11_17, $__local_var_5_4, $dictApplicative_14) {
  $__num = \func_num_args();
  $__res = ((($Apply0_15_24)->{'apply'})(((((($Apply0_15_24)->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Functor_Product_product']))(((($__local_var_5_4)->{'sequence'})($dictApplicative_14))(($v_16)->{'value0'}))))(((($__local_var_11_17)->{'sequence'})($dictApplicative_14))(($v_16)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_14) use ($functorProduct1_12_18) {
  $__num = \func_num_args();
  $__res = $functorProduct1_12_18;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar___unused_14) use ($foldableProduct1_13_20) {
  $__num = \func_num_args();
  $__res = $foldableProduct1_13_20;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["traverseWithIndex" => function($dictApplicative_12) use ($dictTraversableWithIndex1_8, $dictTraversableWithIndex_0) {
  $__num = \func_num_args();
  $Apply0_13_26 = (($dictApplicative_12)->{'Apply0'})(null);
  $__res = function($f_14) use ($Apply0_13_26, $dictApplicative_12, $dictTraversableWithIndex1_8, $dictTraversableWithIndex_0) {
  $__num = \func_num_args();
  $__res = function($v_15) use ($Apply0_13_26, $dictApplicative_12, $dictTraversableWithIndex1_8, $dictTraversableWithIndex_0, $f_14) {
  $__num = \func_num_args();
  $__res = ((($Apply0_13_26)->{'apply'})(((((($Apply0_13_26)->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Functor_Product_product']))((((($dictTraversableWithIndex_0)->{'traverseWithIndex'})($dictApplicative_12))((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_14))($GLOBALS['Data_Either_Left'])))(($v_15)->{'value0'}))))((((($dictTraversableWithIndex1_8)->{'traverseWithIndex'})($dictApplicative_12))((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_14))($GLOBALS['Data_Either_Right'])))(($v_15)->{'value1'}));
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
}, "FunctorWithIndex0" => function($_dollar___unused_12) use ($functorWithIndexProduct1_9_7) {
  $__num = \func_num_args();
  $__res = $functorWithIndexProduct1_9_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($_dollar___unused_12) use ($foldableWithIndexProduct1_10_11) {
  $__num = \func_num_args();
  $__res = $foldableWithIndexProduct1_10_11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($_dollar___unused_12) use ($traversableProduct1_11_17) {
  $__num = \func_num_args();
  $__res = $traversableProduct1_11_17;
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
  $Functor0_1_0 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $traverse8_1_0 = function($f_2) use ($Functor0_1_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($Functor0_1_0, $f_2) {
  $__num = \func_num_args();
  $__res = ((($Functor0_1_0)->{'map'})(function($x_4) {
  $__num = \func_num_args();
  $__res = $x_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($f_2)($v_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
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
}, "FunctorWithIndex0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FunctorWithIndex_functorWithIndexMultiplicative'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FoldableWithIndex_foldableWithIndexMultiplicative'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Traversable_traversableMultiplicative'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_TraversableWithIndex_traversableWithIndexMaybe
$GLOBALS['Data_TraversableWithIndex_traversableWithIndexMaybe'] = (object)["traverseWithIndex" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $Functor0_1_0 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $traverse8_1_0 = function($v_2) use ($Functor0_1_0, $dictApplicative_0) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($Functor0_1_0, $dictApplicative_0, $v_2) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v1_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t1 = (($dictApplicative_0)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_1;;
};
  if ($v1_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = ((($Functor0_1_0)->{'map'})($GLOBALS['Data_Maybe_Just']))(($v_2)(($v1_3)->{'value0'}));
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
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
}, "FunctorWithIndex0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FunctorWithIndex_functorWithIndexMaybe'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FoldableWithIndex_foldableWithIndexMaybe'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Traversable_traversableMaybe'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_TraversableWithIndex_traversableWithIndexLast
$GLOBALS['Data_TraversableWithIndex_traversableWithIndexLast'] = (object)["traverseWithIndex" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $Functor0_1_0 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $traverse8_1_0 = function($f_2) use ($Functor0_1_0, $dictApplicative_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($Functor0_1_0, $dictApplicative_0, $f_2) {
  $__num = \func_num_args();
  $Functor0_4_1 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__t2 = null;;
  if ($v_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t2 = (($dictApplicative_0)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_2;;
};
  if ($v_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t2 = ((($Functor0_4_1)->{'map'})($GLOBALS['Data_Maybe_Just']))(($f_2)(($v_3)->{'value0'}));
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = ((($Functor0_1_0)->{'map'})(function($x_4) {
  $__num = \func_num_args();
  $__res = $x_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($__t2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
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
}, "FunctorWithIndex0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FunctorWithIndex_functorWithIndexLast'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FoldableWithIndex_foldableWithIndexLast'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Traversable_traversableLast'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_TraversableWithIndex_traversableWithIndexIdentity
$GLOBALS['Data_TraversableWithIndex_traversableWithIndexIdentity'] = (object)["traverseWithIndex" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $Functor0_1_0 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($f_2) use ($Functor0_1_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($Functor0_1_0, $f_2) {
  $__num = \func_num_args();
  $__res = ((($Functor0_1_0)->{'map'})(function($x_4) {
  $__num = \func_num_args();
  $__res = $x_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($f_2)($GLOBALS['Data_Unit_unit']))($v_3));
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
}, "FunctorWithIndex0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FunctorWithIndex_functorWithIndexIdentity'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FoldableWithIndex_foldableWithIndexIdentity'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Traversable_traversableIdentity'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_TraversableWithIndex_traversableWithIndexFirst
$GLOBALS['Data_TraversableWithIndex_traversableWithIndexFirst'] = (object)["traverseWithIndex" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $Functor0_1_0 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $traverse8_1_0 = function($f_2) use ($Functor0_1_0, $dictApplicative_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($Functor0_1_0, $dictApplicative_0, $f_2) {
  $__num = \func_num_args();
  $Functor0_4_1 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__t2 = null;;
  if ($v_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t2 = (($dictApplicative_0)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_2;;
};
  if ($v_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t2 = ((($Functor0_4_1)->{'map'})($GLOBALS['Data_Maybe_Just']))(($f_2)(($v_3)->{'value0'}));
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = ((($Functor0_1_0)->{'map'})(function($x_4) {
  $__num = \func_num_args();
  $__res = $x_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($__t2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
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
}, "FunctorWithIndex0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FunctorWithIndex_functorWithIndexFirst'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FoldableWithIndex_foldableWithIndexFirst'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Traversable_traversableFirst'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_TraversableWithIndex_traversableWithIndexEither
$GLOBALS['Data_TraversableWithIndex_traversableWithIndexEither'] = (object)["traverseWithIndex" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $Functor0_1_0 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($v_2) use ($Functor0_1_0, $dictApplicative_0) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($Functor0_1_0, $dictApplicative_0, $v_2) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v1_3 instanceof \Data\Either\Data_Either_Left) {
$__t1 = (($dictApplicative_0)->{'pure'})(new \Data\Either\Data_Either_Left(($v1_3)->{'value0'}));
goto end_branch_1;;
};
  if ($v1_3 instanceof \Data\Either\Data_Either_Right) {
$__t1 = ((($Functor0_1_0)->{'map'})($GLOBALS['Data_Either_Right']))((($v_2)($GLOBALS['Data_Unit_unit']))(($v1_3)->{'value0'}));
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
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
}, "FunctorWithIndex0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FunctorWithIndex_functorWithIndexEither'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FoldableWithIndex_foldableWithIndexEither'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Traversable_traversableEither'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_TraversableWithIndex_traversableWithIndexDual
$GLOBALS['Data_TraversableWithIndex_traversableWithIndexDual'] = (object)["traverseWithIndex" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $Functor0_1_0 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $traverse8_1_0 = function($f_2) use ($Functor0_1_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($Functor0_1_0, $f_2) {
  $__num = \func_num_args();
  $__res = ((($Functor0_1_0)->{'map'})(function($x_4) {
  $__num = \func_num_args();
  $__res = $x_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($f_2)($v_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
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
}, "FunctorWithIndex0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FunctorWithIndex_functorWithIndexDual'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FoldableWithIndex_foldableWithIndexDual'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Traversable_traversableDual'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_TraversableWithIndex_traversableWithIndexDisj
$GLOBALS['Data_TraversableWithIndex_traversableWithIndexDisj'] = (object)["traverseWithIndex" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $Functor0_1_0 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $traverse8_1_0 = function($f_2) use ($Functor0_1_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($Functor0_1_0, $f_2) {
  $__num = \func_num_args();
  $__res = ((($Functor0_1_0)->{'map'})(function($x_4) {
  $__num = \func_num_args();
  $__res = $x_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($f_2)($v_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
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
}, "FunctorWithIndex0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FunctorWithIndex_functorWithIndexDisj'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FoldableWithIndex_foldableWithIndexDisj'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($_dollar___unused_0) {
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
  $__local_var_1_0 = (($dictTraversableWithIndex_0)->{'FunctorWithIndex0'})(null);
  $__local_var_2_1 = (($__local_var_1_0)->{'Functor0'})(null);
  $__local_var_3_2 = (($dictTraversableWithIndex_0)->{'FoldableWithIndex1'})(null);
  $__local_var_4_3 = (($__local_var_3_2)->{'Foldable0'})(null);
  $__local_var_5_4 = (($dictTraversableWithIndex_0)->{'Traversable2'})(null);
  $__local_var_6_5 = (($__local_var_5_4)->{'Functor0'})(null);
  $__local_var_7_6 = (($__local_var_5_4)->{'Foldable1'})(null);
  $__res = function($dictTraversableWithIndex1_8) use ($__local_var_1_0, $__local_var_2_1, $__local_var_3_2, $__local_var_4_3, $__local_var_5_4, $__local_var_6_5, $__local_var_7_6, $dictTraversableWithIndex_0) {
  $__num = \func_num_args();
  $__local_var_9_7 = (($dictTraversableWithIndex1_8)->{'FunctorWithIndex0'})(null);
  $__local_var_10_8 = (($__local_var_9_7)->{'Functor0'})(null);
  $functorCoproduct1_10_8 = (object)["map" => function($f_11) use ($__local_var_10_8, $__local_var_2_1) {
  $__num = \func_num_args();
  $__res = function($v_12) use ($__local_var_10_8, $__local_var_2_1, $f_11) {
  $__num = \func_num_args();
  $__local_var_13_9 = (($__local_var_2_1)->{'map'})($f_11);
  $__local_var_14_10 = (($__local_var_10_8)->{'map'})($f_11);
  $__t11 = null;;
  if ($v_12 instanceof \Data\Either\Data_Either_Left) {
$__t11 = new \Data\Either\Data_Either_Left(($__local_var_13_9)(($v_12)->{'value0'}));
goto end_branch_11;;
};
  if ($v_12 instanceof \Data\Either\Data_Either_Right) {
$__t11 = new \Data\Either\Data_Either_Right(($__local_var_14_10)(($v_12)->{'value0'}));
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
}];
  $functorWithIndexCoproduct1_9_7 = (object)["mapWithIndex" => function($f_11) use ($__local_var_1_0, $__local_var_9_7) {
  $__num = \func_num_args();
  $__res = function($v_12) use ($__local_var_1_0, $__local_var_9_7, $f_11) {
  $__num = \func_num_args();
  $__local_var_13_13 = (($__local_var_1_0)->{'mapWithIndex'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_11))($GLOBALS['Data_Either_Left']));
  $__local_var_14_14 = (($__local_var_9_7)->{'mapWithIndex'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_11))($GLOBALS['Data_Either_Right']));
  $__t15 = null;;
  if ($v_12 instanceof \Data\Either\Data_Either_Left) {
$__t15 = new \Data\Either\Data_Either_Left(($__local_var_13_13)(($v_12)->{'value0'}));
goto end_branch_15;;
};
  if ($v_12 instanceof \Data\Either\Data_Either_Right) {
$__t15 = new \Data\Either\Data_Either_Right(($__local_var_14_14)(($v_12)->{'value0'}));
goto end_branch_15;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t15 = null;
  end_branch_15:;
  $__res = $__t15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_11) use ($functorCoproduct1_10_8) {
  $__num = \func_num_args();
  $__res = $functorCoproduct1_10_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_10_17 = (($dictTraversableWithIndex1_8)->{'FoldableWithIndex1'})(null);
  $__local_var_11_18 = (($__local_var_10_17)->{'Foldable0'})(null);
  $foldableCoproduct1_11_18 = (object)["foldr" => function($f_12) use ($__local_var_11_18, $__local_var_4_3) {
  $__num = \func_num_args();
  $__res = function($z_13) use ($__local_var_11_18, $__local_var_4_3, $f_12) {
  $__num = \func_num_args();
  $__local_var_14_19 = ((($__local_var_4_3)->{'foldr'})($f_12))($z_13);
  $__local_var_15_20 = ((($__local_var_11_18)->{'foldr'})($f_12))($z_13);
  $__res = function($v2_16) use ($__local_var_14_19, $__local_var_15_20) {
  $__num = \func_num_args();
  $__t21 = null;;
  if ($v2_16 instanceof \Data\Either\Data_Either_Left) {
$__t21 = ($__local_var_14_19)(($v2_16)->{'value0'});
goto end_branch_21;;
};
  if ($v2_16 instanceof \Data\Either\Data_Either_Right) {
$__t21 = ($__local_var_15_20)(($v2_16)->{'value0'});
goto end_branch_21;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t21 = null;
  end_branch_21:;
  $__res = $__t21;
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
}, "foldl" => function($f_12) use ($__local_var_11_18, $__local_var_4_3) {
  $__num = \func_num_args();
  $__res = function($z_13) use ($__local_var_11_18, $__local_var_4_3, $f_12) {
  $__num = \func_num_args();
  $__local_var_14_22 = ((($__local_var_4_3)->{'foldl'})($f_12))($z_13);
  $__local_var_15_23 = ((($__local_var_11_18)->{'foldl'})($f_12))($z_13);
  $__res = function($v2_16) use ($__local_var_14_22, $__local_var_15_23) {
  $__num = \func_num_args();
  $__t24 = null;;
  if ($v2_16 instanceof \Data\Either\Data_Either_Left) {
$__t24 = ($__local_var_14_22)(($v2_16)->{'value0'});
goto end_branch_24;;
};
  if ($v2_16 instanceof \Data\Either\Data_Either_Right) {
$__t24 = ($__local_var_15_23)(($v2_16)->{'value0'});
goto end_branch_24;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t24 = null;
  end_branch_24:;
  $__res = $__t24;
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
}, "foldMap" => function($dictMonoid_12) use ($__local_var_11_18, $__local_var_4_3) {
  $__num = \func_num_args();
  $__res = function($f_13) use ($__local_var_11_18, $__local_var_4_3, $dictMonoid_12) {
  $__num = \func_num_args();
  $__local_var_14_25 = ((($__local_var_4_3)->{'foldMap'})($dictMonoid_12))($f_13);
  $__local_var_15_26 = ((($__local_var_11_18)->{'foldMap'})($dictMonoid_12))($f_13);
  $__res = function($v2_16) use ($__local_var_14_25, $__local_var_15_26) {
  $__num = \func_num_args();
  $__t27 = null;;
  if ($v2_16 instanceof \Data\Either\Data_Either_Left) {
$__t27 = ($__local_var_14_25)(($v2_16)->{'value0'});
goto end_branch_27;;
};
  if ($v2_16 instanceof \Data\Either\Data_Either_Right) {
$__t27 = ($__local_var_15_26)(($v2_16)->{'value0'});
goto end_branch_27;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t27 = null;
  end_branch_27:;
  $__res = $__t27;
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
  $foldableWithIndexCoproduct1_10_17 = (object)["foldrWithIndex" => function($f_12) use ($__local_var_10_17, $__local_var_3_2) {
  $__num = \func_num_args();
  $__res = function($z_13) use ($__local_var_10_17, $__local_var_3_2, $f_12) {
  $__num = \func_num_args();
  $__local_var_14_29 = ((($__local_var_3_2)->{'foldrWithIndex'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_12))($GLOBALS['Data_Either_Left'])))($z_13);
  $__local_var_15_30 = ((($__local_var_10_17)->{'foldrWithIndex'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_12))($GLOBALS['Data_Either_Right'])))($z_13);
  $__res = function($v2_16) use ($__local_var_14_29, $__local_var_15_30) {
  $__num = \func_num_args();
  $__t31 = null;;
  if ($v2_16 instanceof \Data\Either\Data_Either_Left) {
$__t31 = ($__local_var_14_29)(($v2_16)->{'value0'});
goto end_branch_31;;
};
  if ($v2_16 instanceof \Data\Either\Data_Either_Right) {
$__t31 = ($__local_var_15_30)(($v2_16)->{'value0'});
goto end_branch_31;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t31 = null;
  end_branch_31:;
  $__res = $__t31;
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
}, "foldlWithIndex" => function($f_12) use ($__local_var_10_17, $__local_var_3_2) {
  $__num = \func_num_args();
  $__res = function($z_13) use ($__local_var_10_17, $__local_var_3_2, $f_12) {
  $__num = \func_num_args();
  $__local_var_14_32 = ((($__local_var_3_2)->{'foldlWithIndex'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_12))($GLOBALS['Data_Either_Left'])))($z_13);
  $__local_var_15_33 = ((($__local_var_10_17)->{'foldlWithIndex'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_12))($GLOBALS['Data_Either_Right'])))($z_13);
  $__res = function($v2_16) use ($__local_var_14_32, $__local_var_15_33) {
  $__num = \func_num_args();
  $__t34 = null;;
  if ($v2_16 instanceof \Data\Either\Data_Either_Left) {
$__t34 = ($__local_var_14_32)(($v2_16)->{'value0'});
goto end_branch_34;;
};
  if ($v2_16 instanceof \Data\Either\Data_Either_Right) {
$__t34 = ($__local_var_15_33)(($v2_16)->{'value0'});
goto end_branch_34;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t34 = null;
  end_branch_34:;
  $__res = $__t34;
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
}, "foldMapWithIndex" => function($dictMonoid_12) use ($__local_var_10_17, $__local_var_3_2) {
  $__num = \func_num_args();
  $__res = function($f_13) use ($__local_var_10_17, $__local_var_3_2, $dictMonoid_12) {
  $__num = \func_num_args();
  $__local_var_14_35 = ((($__local_var_3_2)->{'foldMapWithIndex'})($dictMonoid_12))((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_13))($GLOBALS['Data_Either_Left']));
  $__local_var_15_36 = ((($__local_var_10_17)->{'foldMapWithIndex'})($dictMonoid_12))((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_13))($GLOBALS['Data_Either_Right']));
  $__res = function($v2_16) use ($__local_var_14_35, $__local_var_15_36) {
  $__num = \func_num_args();
  $__t37 = null;;
  if ($v2_16 instanceof \Data\Either\Data_Either_Left) {
$__t37 = ($__local_var_14_35)(($v2_16)->{'value0'});
goto end_branch_37;;
};
  if ($v2_16 instanceof \Data\Either\Data_Either_Right) {
$__t37 = ($__local_var_15_36)(($v2_16)->{'value0'});
goto end_branch_37;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t37 = null;
  end_branch_37:;
  $__res = $__t37;
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
}, "Foldable0" => function($_dollar___unused_12) use ($foldableCoproduct1_11_18) {
  $__num = \func_num_args();
  $__res = $foldableCoproduct1_11_18;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_11_39 = (($dictTraversableWithIndex1_8)->{'Traversable2'})(null);
  $__local_var_12_40 = (($__local_var_11_39)->{'Functor0'})(null);
  $functorCoproduct1_12_40 = (object)["map" => function($f_13) use ($__local_var_12_40, $__local_var_6_5) {
  $__num = \func_num_args();
  $__res = function($v_14) use ($__local_var_12_40, $__local_var_6_5, $f_13) {
  $__num = \func_num_args();
  $__local_var_15_41 = (($__local_var_6_5)->{'map'})($f_13);
  $__local_var_16_42 = (($__local_var_12_40)->{'map'})($f_13);
  $__t43 = null;;
  if ($v_14 instanceof \Data\Either\Data_Either_Left) {
$__t43 = new \Data\Either\Data_Either_Left(($__local_var_15_41)(($v_14)->{'value0'}));
goto end_branch_43;;
};
  if ($v_14 instanceof \Data\Either\Data_Either_Right) {
$__t43 = new \Data\Either\Data_Either_Right(($__local_var_16_42)(($v_14)->{'value0'}));
goto end_branch_43;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t43 = null;
  end_branch_43:;
  $__res = $__t43;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_13_45 = (($__local_var_11_39)->{'Foldable1'})(null);
  $foldableCoproduct1_13_45 = (object)["foldr" => function($f_14) use ($__local_var_13_45, $__local_var_7_6) {
  $__num = \func_num_args();
  $__res = function($z_15) use ($__local_var_13_45, $__local_var_7_6, $f_14) {
  $__num = \func_num_args();
  $__local_var_16_46 = ((($__local_var_7_6)->{'foldr'})($f_14))($z_15);
  $__local_var_17_47 = ((($__local_var_13_45)->{'foldr'})($f_14))($z_15);
  $__res = function($v2_18) use ($__local_var_16_46, $__local_var_17_47) {
  $__num = \func_num_args();
  $__t48 = null;;
  if ($v2_18 instanceof \Data\Either\Data_Either_Left) {
$__t48 = ($__local_var_16_46)(($v2_18)->{'value0'});
goto end_branch_48;;
};
  if ($v2_18 instanceof \Data\Either\Data_Either_Right) {
$__t48 = ($__local_var_17_47)(($v2_18)->{'value0'});
goto end_branch_48;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t48 = null;
  end_branch_48:;
  $__res = $__t48;
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
}, "foldl" => function($f_14) use ($__local_var_13_45, $__local_var_7_6) {
  $__num = \func_num_args();
  $__res = function($z_15) use ($__local_var_13_45, $__local_var_7_6, $f_14) {
  $__num = \func_num_args();
  $__local_var_16_49 = ((($__local_var_7_6)->{'foldl'})($f_14))($z_15);
  $__local_var_17_50 = ((($__local_var_13_45)->{'foldl'})($f_14))($z_15);
  $__res = function($v2_18) use ($__local_var_16_49, $__local_var_17_50) {
  $__num = \func_num_args();
  $__t51 = null;;
  if ($v2_18 instanceof \Data\Either\Data_Either_Left) {
$__t51 = ($__local_var_16_49)(($v2_18)->{'value0'});
goto end_branch_51;;
};
  if ($v2_18 instanceof \Data\Either\Data_Either_Right) {
$__t51 = ($__local_var_17_50)(($v2_18)->{'value0'});
goto end_branch_51;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t51 = null;
  end_branch_51:;
  $__res = $__t51;
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
}, "foldMap" => function($dictMonoid_14) use ($__local_var_13_45, $__local_var_7_6) {
  $__num = \func_num_args();
  $__res = function($f_15) use ($__local_var_13_45, $__local_var_7_6, $dictMonoid_14) {
  $__num = \func_num_args();
  $__local_var_16_52 = ((($__local_var_7_6)->{'foldMap'})($dictMonoid_14))($f_15);
  $__local_var_17_53 = ((($__local_var_13_45)->{'foldMap'})($dictMonoid_14))($f_15);
  $__res = function($v2_18) use ($__local_var_16_52, $__local_var_17_53) {
  $__num = \func_num_args();
  $__t54 = null;;
  if ($v2_18 instanceof \Data\Either\Data_Either_Left) {
$__t54 = ($__local_var_16_52)(($v2_18)->{'value0'});
goto end_branch_54;;
};
  if ($v2_18 instanceof \Data\Either\Data_Either_Right) {
$__t54 = ($__local_var_17_53)(($v2_18)->{'value0'});
goto end_branch_54;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t54 = null;
  end_branch_54:;
  $__res = $__t54;
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
  $traversableCoproduct1_11_39 = (object)["traverse" => function($dictApplicative_14) use ($__local_var_11_39, $__local_var_5_4) {
  $__num = \func_num_args();
  $Functor0_15_56 = (((($dictApplicative_14)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($f_16) use ($Functor0_15_56, $__local_var_11_39, $__local_var_5_4, $dictApplicative_14) {
  $__num = \func_num_args();
  $__local_var_17_57 = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($Functor0_15_56)->{'map'})((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_17) {
  $__num = \func_num_args();
  $__res = $x_17;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Either_Left']))))(((($__local_var_5_4)->{'traverse'})($dictApplicative_14))($f_16));
  $__local_var_18_58 = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($Functor0_15_56)->{'map'})((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_18) {
  $__num = \func_num_args();
  $__res = $x_18;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Either_Right']))))(((($__local_var_11_39)->{'traverse'})($dictApplicative_14))($f_16));
  $__res = function($v2_19) use ($__local_var_17_57, $__local_var_18_58) {
  $__num = \func_num_args();
  $__t59 = null;;
  if ($v2_19 instanceof \Data\Either\Data_Either_Left) {
$__t59 = ($__local_var_17_57)(($v2_19)->{'value0'});
goto end_branch_59;;
};
  if ($v2_19 instanceof \Data\Either\Data_Either_Right) {
$__t59 = ($__local_var_18_58)(($v2_19)->{'value0'});
goto end_branch_59;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t59 = null;
  end_branch_59:;
  $__res = $__t59;
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
}, "sequence" => function($dictApplicative_14) use ($__local_var_11_39, $__local_var_5_4) {
  $__num = \func_num_args();
  $Functor0_15_60 = (((($dictApplicative_14)->{'Apply0'})(null))->{'Functor0'})(null);
  $__local_var_16_61 = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($Functor0_15_60)->{'map'})((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_16) {
  $__num = \func_num_args();
  $__res = $x_16;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Either_Left']))))((($__local_var_5_4)->{'sequence'})($dictApplicative_14));
  $__local_var_17_62 = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($Functor0_15_60)->{'map'})((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_17) {
  $__num = \func_num_args();
  $__res = $x_17;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Either_Right']))))((($__local_var_11_39)->{'sequence'})($dictApplicative_14));
  $__res = function($v2_18) use ($__local_var_16_61, $__local_var_17_62) {
  $__num = \func_num_args();
  $__t63 = null;;
  if ($v2_18 instanceof \Data\Either\Data_Either_Left) {
$__t63 = ($__local_var_16_61)(($v2_18)->{'value0'});
goto end_branch_63;;
};
  if ($v2_18 instanceof \Data\Either\Data_Either_Right) {
$__t63 = ($__local_var_17_62)(($v2_18)->{'value0'});
goto end_branch_63;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t63 = null;
  end_branch_63:;
  $__res = $__t63;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_14) use ($functorCoproduct1_12_40) {
  $__num = \func_num_args();
  $__res = $functorCoproduct1_12_40;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar___unused_14) use ($foldableCoproduct1_13_45) {
  $__num = \func_num_args();
  $__res = $foldableCoproduct1_13_45;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["traverseWithIndex" => function($dictApplicative_12) use ($dictTraversableWithIndex1_8, $dictTraversableWithIndex_0) {
  $__num = \func_num_args();
  $Functor0_13_65 = (((($dictApplicative_12)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($f_14) use ($Functor0_13_65, $dictApplicative_12, $dictTraversableWithIndex1_8, $dictTraversableWithIndex_0) {
  $__num = \func_num_args();
  $__local_var_15_66 = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($Functor0_13_65)->{'map'})((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_15) {
  $__num = \func_num_args();
  $__res = $x_15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Either_Left']))))(((($dictTraversableWithIndex_0)->{'traverseWithIndex'})($dictApplicative_12))((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_14))($GLOBALS['Data_Either_Left'])));
  $__local_var_16_67 = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($Functor0_13_65)->{'map'})((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_16) {
  $__num = \func_num_args();
  $__res = $x_16;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Either_Right']))))(((($dictTraversableWithIndex1_8)->{'traverseWithIndex'})($dictApplicative_12))((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_14))($GLOBALS['Data_Either_Right'])));
  $__res = function($v2_17) use ($__local_var_15_66, $__local_var_16_67) {
  $__num = \func_num_args();
  $__t68 = null;;
  if ($v2_17 instanceof \Data\Either\Data_Either_Left) {
$__t68 = ($__local_var_15_66)(($v2_17)->{'value0'});
goto end_branch_68;;
};
  if ($v2_17 instanceof \Data\Either\Data_Either_Right) {
$__t68 = ($__local_var_16_67)(($v2_17)->{'value0'});
goto end_branch_68;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t68 = null;
  end_branch_68:;
  $__res = $__t68;
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
}, "FunctorWithIndex0" => function($_dollar___unused_12) use ($functorWithIndexCoproduct1_9_7) {
  $__num = \func_num_args();
  $__res = $functorWithIndexCoproduct1_9_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($_dollar___unused_12) use ($foldableWithIndexCoproduct1_10_17) {
  $__num = \func_num_args();
  $__res = $foldableWithIndexCoproduct1_10_17;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($_dollar___unused_12) use ($traversableCoproduct1_11_39) {
  $__num = \func_num_args();
  $__res = $traversableCoproduct1_11_39;
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
$GLOBALS['Data_TraversableWithIndex_traversableWithIndexConst'] = (object)["traverseWithIndex" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $__res = function($v_1) use ($dictApplicative_0) {
  $__num = \func_num_args();
  $__res = function($v1_2) use ($dictApplicative_0) {
  $__num = \func_num_args();
  $__res = (($dictApplicative_0)->{'pure'})($v1_2);
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
}, "FunctorWithIndex0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FunctorWithIndex_functorWithIndexConst'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FoldableWithIndex_foldableWithIndexConst'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Traversable_traversableConst'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_TraversableWithIndex_traversableWithIndexConj
$GLOBALS['Data_TraversableWithIndex_traversableWithIndexConj'] = (object)["traverseWithIndex" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $Functor0_1_0 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $traverse8_1_0 = function($f_2) use ($Functor0_1_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($Functor0_1_0, $f_2) {
  $__num = \func_num_args();
  $__res = ((($Functor0_1_0)->{'map'})(function($x_4) {
  $__num = \func_num_args();
  $__res = $x_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($f_2)($v_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
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
}, "FunctorWithIndex0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FunctorWithIndex_functorWithIndexConj'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FoldableWithIndex_foldableWithIndexConj'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($_dollar___unused_0) {
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
  $__local_var_1_0 = (($dictTraversableWithIndex_0)->{'FunctorWithIndex0'})(null);
  $__local_var_2_1 = (($__local_var_1_0)->{'Functor0'})(null);
  $__local_var_3_2 = (($dictTraversableWithIndex_0)->{'FoldableWithIndex1'})(null);
  $__local_var_4_3 = (($__local_var_3_2)->{'Foldable0'})(null);
  $__local_var_5_4 = (($dictTraversableWithIndex_0)->{'Traversable2'})(null);
  $__local_var_6_5 = (($__local_var_5_4)->{'Functor0'})(null);
  $__local_var_7_6 = (($__local_var_5_4)->{'Foldable1'})(null);
  $__res = function($dictTraversableWithIndex1_8) use ($__local_var_1_0, $__local_var_2_1, $__local_var_3_2, $__local_var_4_3, $__local_var_5_4, $__local_var_6_5, $__local_var_7_6, $dictTraversableWithIndex_0) {
  $__num = \func_num_args();
  $__local_var_9_7 = (($dictTraversableWithIndex1_8)->{'FunctorWithIndex0'})(null);
  $mapWithIndex1_10_8 = ($__local_var_9_7)->{'mapWithIndex'};
  $__local_var_11_9 = (($__local_var_9_7)->{'Functor0'})(null);
  $functorCompose1_11_9 = (object)["map" => function($f_12) use ($__local_var_11_9, $__local_var_2_1) {
  $__num = \func_num_args();
  $__res = function($v_13) use ($__local_var_11_9, $__local_var_2_1, $f_12) {
  $__num = \func_num_args();
  $__res = ((($__local_var_2_1)->{'map'})((($__local_var_11_9)->{'map'})($f_12)))($v_13);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $functorWithIndexCompose1_9_7 = (object)["mapWithIndex" => function($f_12) use ($__local_var_1_0, $mapWithIndex1_10_8) {
  $__num = \func_num_args();
  $__res = function($v_13) use ($__local_var_1_0, $f_12, $mapWithIndex1_10_8) {
  $__num = \func_num_args();
  $__res = ((($__local_var_1_0)->{'mapWithIndex'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($mapWithIndex1_10_8))(function($a_14) use ($f_12) {
  $__num = \func_num_args();
  $__res = function($b_15) use ($a_14, $f_12) {
  $__num = \func_num_args();
  $__res = ($f_12)(new \Data\Tuple\Data_Tuple_Tuple($a_14, $b_15));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))($v_13);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_12) use ($functorCompose1_11_9) {
  $__num = \func_num_args();
  $__res = $functorCompose1_11_9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_10_12 = (($dictTraversableWithIndex1_8)->{'FoldableWithIndex1'})(null);
  $foldlWithIndex1_11_13 = ($__local_var_10_12)->{'foldlWithIndex'};
  $__local_var_12_14 = (($__local_var_10_12)->{'Foldable0'})(null);
  $foldableCompose1_12_14 = (object)["foldr" => function($f_13) use ($__local_var_12_14, $__local_var_4_3) {
  $__num = \func_num_args();
  $__res = function($i_14) use ($__local_var_12_14, $__local_var_4_3, $f_13) {
  $__num = \func_num_args();
  $__res = function($v_15) use ($__local_var_12_14, $__local_var_4_3, $f_13, $i_14) {
  $__num = \func_num_args();
  $__local_var_16_15 = (($__local_var_12_14)->{'foldr'})($f_13);
  $__res = (((($__local_var_4_3)->{'foldr'})(function($b_17) use ($__local_var_16_15) {
  $__num = \func_num_args();
  $__res = function($a_18) use ($__local_var_16_15, $b_17) {
  $__num = \func_num_args();
  $__res = (($__local_var_16_15)($a_18))($b_17);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($i_14))($v_15);
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
}, "foldl" => function($f_13) use ($__local_var_12_14, $__local_var_4_3) {
  $__num = \func_num_args();
  $__res = function($i_14) use ($__local_var_12_14, $__local_var_4_3, $f_13) {
  $__num = \func_num_args();
  $__res = function($v_15) use ($__local_var_12_14, $__local_var_4_3, $f_13, $i_14) {
  $__num = \func_num_args();
  $__res = (((($__local_var_4_3)->{'foldl'})((($__local_var_12_14)->{'foldl'})($f_13)))($i_14))($v_15);
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
}, "foldMap" => function($dictMonoid_13) use ($__local_var_12_14, $__local_var_4_3) {
  $__num = \func_num_args();
  $__res = function($f_14) use ($__local_var_12_14, $__local_var_4_3, $dictMonoid_13) {
  $__num = \func_num_args();
  $__res = function($v_15) use ($__local_var_12_14, $__local_var_4_3, $dictMonoid_13, $f_14) {
  $__num = \func_num_args();
  $__res = (((($__local_var_4_3)->{'foldMap'})($dictMonoid_13))(((($__local_var_12_14)->{'foldMap'})($dictMonoid_13))($f_14)))($v_15);
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
  $foldableWithIndexCompose1_10_12 = (object)["foldrWithIndex" => function($f_13) use ($__local_var_10_12, $__local_var_3_2) {
  $__num = \func_num_args();
  $__res = function($i_14) use ($__local_var_10_12, $__local_var_3_2, $f_13) {
  $__num = \func_num_args();
  $__res = function($v_15) use ($__local_var_10_12, $__local_var_3_2, $f_13, $i_14) {
  $__num = \func_num_args();
  $__res = (((($__local_var_3_2)->{'foldrWithIndex'})(function($a_16) use ($__local_var_10_12, $f_13) {
  $__num = \func_num_args();
  $__local_var_17_17 = (($__local_var_10_12)->{'foldrWithIndex'})(function($b_17) use ($a_16, $f_13) {
  $__num = \func_num_args();
  $__res = ($f_13)(new \Data\Tuple\Data_Tuple_Tuple($a_16, $b_17));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($b_18) use ($__local_var_17_17) {
  $__num = \func_num_args();
  $__res = function($a_19) use ($__local_var_17_17, $b_18) {
  $__num = \func_num_args();
  $__res = (($__local_var_17_17)($a_19))($b_18);
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
}))($i_14))($v_15);
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
}, "foldlWithIndex" => function($f_13) use ($__local_var_3_2, $foldlWithIndex1_11_13) {
  $__num = \func_num_args();
  $__res = function($i_14) use ($__local_var_3_2, $f_13, $foldlWithIndex1_11_13) {
  $__num = \func_num_args();
  $__res = function($v_15) use ($__local_var_3_2, $f_13, $foldlWithIndex1_11_13, $i_14) {
  $__num = \func_num_args();
  $__res = (((($__local_var_3_2)->{'foldlWithIndex'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($foldlWithIndex1_11_13))(function($a_16) use ($f_13) {
  $__num = \func_num_args();
  $__res = function($b_17) use ($a_16, $f_13) {
  $__num = \func_num_args();
  $__res = ($f_13)(new \Data\Tuple\Data_Tuple_Tuple($a_16, $b_17));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))($i_14))($v_15);
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
}, "foldMapWithIndex" => function($dictMonoid_13) use ($__local_var_10_12, $__local_var_3_2) {
  $__num = \func_num_args();
  $foldMapWithIndex2_14_18 = (($__local_var_10_12)->{'foldMapWithIndex'})($dictMonoid_13);
  $__res = function($f_15) use ($__local_var_3_2, $dictMonoid_13, $foldMapWithIndex2_14_18) {
  $__num = \func_num_args();
  $__res = function($v_16) use ($__local_var_3_2, $dictMonoid_13, $f_15, $foldMapWithIndex2_14_18) {
  $__num = \func_num_args();
  $__res = (((($__local_var_3_2)->{'foldMapWithIndex'})($dictMonoid_13))((($GLOBALS['Control_Semigroupoid_composeImpl'])($foldMapWithIndex2_14_18))(function($a_17) use ($f_15) {
  $__num = \func_num_args();
  $__res = function($b_18) use ($a_17, $f_15) {
  $__num = \func_num_args();
  $__res = ($f_15)(new \Data\Tuple\Data_Tuple_Tuple($a_17, $b_18));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))($v_16);
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
}, "Foldable0" => function($_dollar___unused_13) use ($foldableCompose1_12_14) {
  $__num = \func_num_args();
  $__res = $foldableCompose1_12_14;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_11_20 = (($dictTraversableWithIndex1_8)->{'Traversable2'})(null);
  $__local_var_12_21 = (($__local_var_11_20)->{'Functor0'})(null);
  $functorCompose1_12_21 = (object)["map" => function($f_13) use ($__local_var_12_21, $__local_var_6_5) {
  $__num = \func_num_args();
  $__res = function($v_14) use ($__local_var_12_21, $__local_var_6_5, $f_13) {
  $__num = \func_num_args();
  $__res = ((($__local_var_6_5)->{'map'})((($__local_var_12_21)->{'map'})($f_13)))($v_14);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_13_23 = (($__local_var_11_20)->{'Foldable1'})(null);
  $foldableCompose1_13_23 = (object)["foldr" => function($f_14) use ($__local_var_13_23, $__local_var_7_6) {
  $__num = \func_num_args();
  $__res = function($i_15) use ($__local_var_13_23, $__local_var_7_6, $f_14) {
  $__num = \func_num_args();
  $__res = function($v_16) use ($__local_var_13_23, $__local_var_7_6, $f_14, $i_15) {
  $__num = \func_num_args();
  $__local_var_17_24 = (($__local_var_13_23)->{'foldr'})($f_14);
  $__res = (((($__local_var_7_6)->{'foldr'})(function($b_18) use ($__local_var_17_24) {
  $__num = \func_num_args();
  $__res = function($a_19) use ($__local_var_17_24, $b_18) {
  $__num = \func_num_args();
  $__res = (($__local_var_17_24)($a_19))($b_18);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($i_15))($v_16);
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
}, "foldl" => function($f_14) use ($__local_var_13_23, $__local_var_7_6) {
  $__num = \func_num_args();
  $__res = function($i_15) use ($__local_var_13_23, $__local_var_7_6, $f_14) {
  $__num = \func_num_args();
  $__res = function($v_16) use ($__local_var_13_23, $__local_var_7_6, $f_14, $i_15) {
  $__num = \func_num_args();
  $__res = (((($__local_var_7_6)->{'foldl'})((($__local_var_13_23)->{'foldl'})($f_14)))($i_15))($v_16);
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
}, "foldMap" => function($dictMonoid_14) use ($__local_var_13_23, $__local_var_7_6) {
  $__num = \func_num_args();
  $__res = function($f_15) use ($__local_var_13_23, $__local_var_7_6, $dictMonoid_14) {
  $__num = \func_num_args();
  $__res = function($v_16) use ($__local_var_13_23, $__local_var_7_6, $dictMonoid_14, $f_15) {
  $__num = \func_num_args();
  $__res = (((($__local_var_7_6)->{'foldMap'})($dictMonoid_14))(((($__local_var_13_23)->{'foldMap'})($dictMonoid_14))($f_15)))($v_16);
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
  $traversableCompose1_11_20 = (object)["traverse" => function($dictApplicative_14) use ($__local_var_11_20, $__local_var_5_4) {
  $__num = \func_num_args();
  $Functor0_15_26 = (((($dictApplicative_14)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($f_16) use ($Functor0_15_26, $__local_var_11_20, $__local_var_5_4, $dictApplicative_14) {
  $__num = \func_num_args();
  $__res = function($v_17) use ($Functor0_15_26, $__local_var_11_20, $__local_var_5_4, $dictApplicative_14, $f_16) {
  $__num = \func_num_args();
  $__res = ((($Functor0_15_26)->{'map'})(function($x_18) {
  $__num = \func_num_args();
  $__res = $x_18;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((((($__local_var_5_4)->{'traverse'})($dictApplicative_14))(((($__local_var_11_20)->{'traverse'})($dictApplicative_14))($f_16)))($v_17));
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
}, "sequence" => function($dictApplicative_14) use ($__local_var_11_20, $__local_var_5_4) {
  $__num = \func_num_args();
  $__res = ((((($GLOBALS['Data_Traversable_traversableCompose'])($__local_var_5_4))($__local_var_11_20))->{'traverse'})($dictApplicative_14))(function($x_15) {
  $__num = \func_num_args();
  $__res = $x_15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_14) use ($functorCompose1_12_21) {
  $__num = \func_num_args();
  $__res = $functorCompose1_12_21;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar___unused_14) use ($foldableCompose1_13_23) {
  $__num = \func_num_args();
  $__res = $foldableCompose1_13_23;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["traverseWithIndex" => function($dictApplicative_12) use ($dictTraversableWithIndex1_8, $dictTraversableWithIndex_0) {
  $__num = \func_num_args();
  $Functor0_13_28 = (((($dictApplicative_12)->{'Apply0'})(null))->{'Functor0'})(null);
  $traverseWithIndex2_14_29 = (($dictTraversableWithIndex1_8)->{'traverseWithIndex'})($dictApplicative_12);
  $__res = function($f_15) use ($Functor0_13_28, $dictApplicative_12, $dictTraversableWithIndex_0, $traverseWithIndex2_14_29) {
  $__num = \func_num_args();
  $__res = function($v_16) use ($Functor0_13_28, $dictApplicative_12, $dictTraversableWithIndex_0, $f_15, $traverseWithIndex2_14_29) {
  $__num = \func_num_args();
  $__res = ((($Functor0_13_28)->{'map'})(function($x_17) {
  $__num = \func_num_args();
  $__res = $x_17;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((((($dictTraversableWithIndex_0)->{'traverseWithIndex'})($dictApplicative_12))((($GLOBALS['Control_Semigroupoid_composeImpl'])($traverseWithIndex2_14_29))(function($a_17) use ($f_15) {
  $__num = \func_num_args();
  $__res = function($b_18) use ($a_17, $f_15) {
  $__num = \func_num_args();
  $__res = ($f_15)(new \Data\Tuple\Data_Tuple_Tuple($a_17, $b_18));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))($v_16));
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
}, "FunctorWithIndex0" => function($_dollar___unused_12) use ($functorWithIndexCompose1_9_7) {
  $__num = \func_num_args();
  $__res = $functorWithIndexCompose1_9_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($_dollar___unused_12) use ($foldableWithIndexCompose1_10_12) {
  $__num = \func_num_args();
  $__res = $foldableWithIndexCompose1_10_12;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($_dollar___unused_12) use ($traversableCompose1_11_20) {
  $__num = \func_num_args();
  $__res = $traversableCompose1_11_20;
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
  $FunctorWithIndex0_1_0 = (($GLOBALS['Data_TraversableWithIndex_traversableWithIndexArray'])->{'FunctorWithIndex0'})(null);
  $sequence1_2_1 = (((($GLOBALS['Data_TraversableWithIndex_traversableWithIndexArray'])->{'Traversable2'})(null))->{'sequence'})($dictApplicative_0);
  $__res = function($f_3) use ($FunctorWithIndex0_1_0, $sequence1_2_1) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($sequence1_2_1))((($FunctorWithIndex0_1_0)->{'mapWithIndex'})($f_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FunctorWithIndex0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FunctorWithIndex_functorWithIndexArray'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FoldableWithIndex_foldableWithIndexArray'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($_dollar___unused_0) {
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
  $functorApp_2_1 = (($__local_var_1_0)->{'Functor0'})(null);
  $functorWithIndexApp_1_0 = (object)["mapWithIndex" => function($f_3) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($__local_var_1_0, $f_3) {
  $__num = \func_num_args();
  $__res = ((($__local_var_1_0)->{'mapWithIndex'})($f_3))($v_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_3) use ($functorApp_2_1) {
  $__num = \func_num_args();
  $__res = $functorApp_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_2_3 = (($dictTraversableWithIndex_0)->{'FoldableWithIndex1'})(null);
  $__local_var_3_4 = (($__local_var_2_3)->{'Foldable0'})(null);
  $foldableApp_3_4 = (object)["foldr" => function($f_4) use ($__local_var_3_4) {
  $__num = \func_num_args();
  $__res = function($i_5) use ($__local_var_3_4, $f_4) {
  $__num = \func_num_args();
  $__res = function($v_6) use ($__local_var_3_4, $f_4, $i_5) {
  $__num = \func_num_args();
  $__res = (((($__local_var_3_4)->{'foldr'})($f_4))($i_5))($v_6);
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
}, "foldl" => function($f_4) use ($__local_var_3_4) {
  $__num = \func_num_args();
  $__res = function($i_5) use ($__local_var_3_4, $f_4) {
  $__num = \func_num_args();
  $__res = function($v_6) use ($__local_var_3_4, $f_4, $i_5) {
  $__num = \func_num_args();
  $__res = (((($__local_var_3_4)->{'foldl'})($f_4))($i_5))($v_6);
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
}, "foldMap" => function($dictMonoid_4) use ($__local_var_3_4) {
  $__num = \func_num_args();
  $__res = function($f_5) use ($__local_var_3_4, $dictMonoid_4) {
  $__num = \func_num_args();
  $__res = function($v_6) use ($__local_var_3_4, $dictMonoid_4, $f_5) {
  $__num = \func_num_args();
  $__res = (((($__local_var_3_4)->{'foldMap'})($dictMonoid_4))($f_5))($v_6);
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
  $foldableWithIndexApp_2_3 = (object)["foldrWithIndex" => function($f_4) use ($__local_var_2_3) {
  $__num = \func_num_args();
  $__res = function($z_5) use ($__local_var_2_3, $f_4) {
  $__num = \func_num_args();
  $__res = function($v_6) use ($__local_var_2_3, $f_4, $z_5) {
  $__num = \func_num_args();
  $__res = (((($__local_var_2_3)->{'foldrWithIndex'})($f_4))($z_5))($v_6);
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
}, "foldlWithIndex" => function($f_4) use ($__local_var_2_3) {
  $__num = \func_num_args();
  $__res = function($z_5) use ($__local_var_2_3, $f_4) {
  $__num = \func_num_args();
  $__res = function($v_6) use ($__local_var_2_3, $f_4, $z_5) {
  $__num = \func_num_args();
  $__res = (((($__local_var_2_3)->{'foldlWithIndex'})($f_4))($z_5))($v_6);
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
}, "foldMapWithIndex" => function($dictMonoid_4) use ($__local_var_2_3) {
  $__num = \func_num_args();
  $__res = function($f_5) use ($__local_var_2_3, $dictMonoid_4) {
  $__num = \func_num_args();
  $__res = function($v_6) use ($__local_var_2_3, $dictMonoid_4, $f_5) {
  $__num = \func_num_args();
  $__res = (((($__local_var_2_3)->{'foldMapWithIndex'})($dictMonoid_4))($f_5))($v_6);
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
}, "Foldable0" => function($_dollar___unused_4) use ($foldableApp_3_4) {
  $__num = \func_num_args();
  $__res = $foldableApp_3_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_3_7 = (($dictTraversableWithIndex_0)->{'Traversable2'})(null);
  $functorApp_4_8 = (($__local_var_3_7)->{'Functor0'})(null);
  $__local_var_5_9 = (($__local_var_3_7)->{'Foldable1'})(null);
  $foldableApp_5_9 = (object)["foldr" => function($f_6) use ($__local_var_5_9) {
  $__num = \func_num_args();
  $__res = function($i_7) use ($__local_var_5_9, $f_6) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($__local_var_5_9, $f_6, $i_7) {
  $__num = \func_num_args();
  $__res = (((($__local_var_5_9)->{'foldr'})($f_6))($i_7))($v_8);
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
}, "foldl" => function($f_6) use ($__local_var_5_9) {
  $__num = \func_num_args();
  $__res = function($i_7) use ($__local_var_5_9, $f_6) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($__local_var_5_9, $f_6, $i_7) {
  $__num = \func_num_args();
  $__res = (((($__local_var_5_9)->{'foldl'})($f_6))($i_7))($v_8);
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
}, "foldMap" => function($dictMonoid_6) use ($__local_var_5_9) {
  $__num = \func_num_args();
  $__res = function($f_7) use ($__local_var_5_9, $dictMonoid_6) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($__local_var_5_9, $dictMonoid_6, $f_7) {
  $__num = \func_num_args();
  $__res = (((($__local_var_5_9)->{'foldMap'})($dictMonoid_6))($f_7))($v_8);
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
  $traversableApp_3_7 = (object)["traverse" => function($dictApplicative_6) use ($__local_var_3_7) {
  $__num = \func_num_args();
  $Functor0_7_11 = (((($dictApplicative_6)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($f_8) use ($Functor0_7_11, $__local_var_3_7, $dictApplicative_6) {
  $__num = \func_num_args();
  $__res = function($v_9) use ($Functor0_7_11, $__local_var_3_7, $dictApplicative_6, $f_8) {
  $__num = \func_num_args();
  $__res = ((($Functor0_7_11)->{'map'})(function($x_10) {
  $__num = \func_num_args();
  $__res = $x_10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((((($__local_var_3_7)->{'traverse'})($dictApplicative_6))($f_8))($v_9));
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
}, "sequence" => function($dictApplicative_6) use ($__local_var_3_7) {
  $__num = \func_num_args();
  $Functor0_7_12 = (((($dictApplicative_6)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($v_8) use ($Functor0_7_12, $__local_var_3_7, $dictApplicative_6) {
  $__num = \func_num_args();
  $__res = ((($Functor0_7_12)->{'map'})(function($x_9) {
  $__num = \func_num_args();
  $__res = $x_9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($__local_var_3_7)->{'sequence'})($dictApplicative_6))($v_8));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_6) use ($functorApp_4_8) {
  $__num = \func_num_args();
  $__res = $functorApp_4_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar___unused_6) use ($foldableApp_5_9) {
  $__num = \func_num_args();
  $__res = $foldableApp_5_9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["traverseWithIndex" => function($dictApplicative_4) use ($dictTraversableWithIndex_0) {
  $__num = \func_num_args();
  $Functor0_5_14 = (((($dictApplicative_4)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($f_6) use ($Functor0_5_14, $dictApplicative_4, $dictTraversableWithIndex_0) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($Functor0_5_14, $dictApplicative_4, $dictTraversableWithIndex_0, $f_6) {
  $__num = \func_num_args();
  $__res = ((($Functor0_5_14)->{'map'})(function($x_8) {
  $__num = \func_num_args();
  $__res = $x_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((((($dictTraversableWithIndex_0)->{'traverseWithIndex'})($dictApplicative_4))($f_6))($v_7));
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
}, "FunctorWithIndex0" => function($_dollar___unused_4) use ($functorWithIndexApp_1_0) {
  $__num = \func_num_args();
  $__res = $functorWithIndexApp_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($_dollar___unused_4) use ($foldableWithIndexApp_2_3) {
  $__num = \func_num_args();
  $__res = $foldableWithIndexApp_2_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($_dollar___unused_4) use ($traversableApp_3_7) {
  $__num = \func_num_args();
  $__res = $traversableApp_3_7;
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
  $Functor0_1_0 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $traverse8_1_0 = function($f_2) use ($Functor0_1_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($Functor0_1_0, $f_2) {
  $__num = \func_num_args();
  $__res = ((($Functor0_1_0)->{'map'})(function($x_4) {
  $__num = \func_num_args();
  $__res = $x_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($f_2)($v_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
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
}, "FunctorWithIndex0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FunctorWithIndex_functorWithIndexAdditive'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_FoldableWithIndex_foldableWithIndexAdditive'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Traversable_traversableAdditive'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_TraversableWithIndex_mapAccumRWithIndex
function majData_majTraversablemajWithmajIndex_mapmajAccummajRmajWithmajIndex($dictTraversableWithIndex_0, $f_1 = null, $s0_2 = null, $xs_3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTraversablemajWithmajIndex_mapmajAccummajRmajWithmajIndex';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = ((((($dictTraversableWithIndex_0)->{'traverseWithIndex'})($GLOBALS['Data_Traversable_Accum_Internal_applicativeStateR']))(function($i_4) use ($f_1) {
  $__num = \func_num_args();
  $__res = function($a_5) use ($f_1, $i_4) {
  $__num = \func_num_args();
  $__res = function($s_6) use ($a_5, $f_1, $i_4) {
  $__num = \func_num_args();
  $__res = ((($f_1)($i_4))($s_6))($a_5);
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
}))($xs_3))($s0_2);
  goto __end;;
  __end:
  return 4 < $__num ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
}
$GLOBALS['Data_TraversableWithIndex_mapAccumRWithIndex'] = __NAMESPACE__ . '\\majData_majTraversablemajWithmajIndex_mapmajAccummajRmajWithmajIndex';

// Data_TraversableWithIndex_scanrWithIndex
function majData_majTraversablemajWithmajIndex_scanrmajWithmajIndex($dictTraversableWithIndex_0, $f_1 = null, $b0_2 = null, $xs_3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTraversablemajWithmajIndex_scanrmajWithmajIndex';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = (((((($dictTraversableWithIndex_0)->{'traverseWithIndex'})($GLOBALS['Data_Traversable_Accum_Internal_applicativeStateR']))(function($i_4) use ($f_1) {
  $__num = \func_num_args();
  $__res = function($a_5) use ($f_1, $i_4) {
  $__num = \func_num_args();
  $__res = function($s_6) use ($a_5, $f_1, $i_4) {
  $__num = \func_num_args();
  $b_prime__7_0 = ((($f_1)($i_4))($a_5))($s_6);
  $__res = (object)["accum" => $b_prime__7_0, "value" => $b_prime__7_0];
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
}))($xs_3))($b0_2))->{'value'};
  goto __end;;
  __end:
  return 4 < $__num ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
}
$GLOBALS['Data_TraversableWithIndex_scanrWithIndex'] = __NAMESPACE__ . '\\majData_majTraversablemajWithmajIndex_scanrmajWithmajIndex';

// Data_TraversableWithIndex_mapAccumLWithIndex
function majData_majTraversablemajWithmajIndex_mapmajAccummajLmajWithmajIndex($dictTraversableWithIndex_0, $f_1 = null, $s0_2 = null, $xs_3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTraversablemajWithmajIndex_mapmajAccummajLmajWithmajIndex';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = ((((($dictTraversableWithIndex_0)->{'traverseWithIndex'})($GLOBALS['Data_Traversable_Accum_Internal_applicativeStateL']))(function($i_4) use ($f_1) {
  $__num = \func_num_args();
  $__res = function($a_5) use ($f_1, $i_4) {
  $__num = \func_num_args();
  $__res = function($s_6) use ($a_5, $f_1, $i_4) {
  $__num = \func_num_args();
  $__res = ((($f_1)($i_4))($s_6))($a_5);
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
}))($xs_3))($s0_2);
  goto __end;;
  __end:
  return 4 < $__num ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
}
$GLOBALS['Data_TraversableWithIndex_mapAccumLWithIndex'] = __NAMESPACE__ . '\\majData_majTraversablemajWithmajIndex_mapmajAccummajLmajWithmajIndex';

// Data_TraversableWithIndex_scanlWithIndex
function majData_majTraversablemajWithmajIndex_scanlmajWithmajIndex($dictTraversableWithIndex_0, $f_1 = null, $b0_2 = null, $xs_3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTraversablemajWithmajIndex_scanlmajWithmajIndex';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = (((((($dictTraversableWithIndex_0)->{'traverseWithIndex'})($GLOBALS['Data_Traversable_Accum_Internal_applicativeStateL']))(function($i_4) use ($f_1) {
  $__num = \func_num_args();
  $__res = function($a_5) use ($f_1, $i_4) {
  $__num = \func_num_args();
  $__res = function($s_6) use ($a_5, $f_1, $i_4) {
  $__num = \func_num_args();
  $b_prime__7_0 = ((($f_1)($i_4))($s_6))($a_5);
  $__res = (object)["accum" => $b_prime__7_0, "value" => $b_prime__7_0];
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
}))($xs_3))($b0_2))->{'value'};
  goto __end;;
  __end:
  return 4 < $__num ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
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
  $__res = function($b_3) use ($__local_var_2_0) {
  $__num = \func_num_args();
  $__res = function($a_4) use ($__local_var_2_0, $b_3) {
  $__num = \func_num_args();
  $__res = (($__local_var_2_0)($a_4))($b_3);
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
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_TraversableWithIndex_forWithIndex'] = __NAMESPACE__ . '\\majData_majTraversablemajWithmajIndex_formajWithmajIndex';

