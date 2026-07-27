<?php

namespace Data\FunctorWithIndex;

// ALL IMPORTS: Control.Semigroupoid, Data.Bifunctor, Data.Const, Data.Either, Data.Function, Data.Functor, Data.Functor.App, Data.Functor.Compose, Data.Functor.Coproduct, Data.Functor.Product, Data.FunctorWithIndex, Data.Identity, Data.Maybe, Data.Maybe.First, Data.Maybe.Last, Data.Monoid.Additive, Data.Monoid.Conj, Data.Monoid.Disj, Data.Monoid.Dual, Data.Monoid.Multiplicative, Data.Tuple, Data.Unit, Prelude, Prim
// TO REQUIRE: Control.Semigroupoid, Data.Bifunctor, Data.Const, Data.Either, Data.Function, Data.Functor, Data.Functor.App, Data.Functor.Compose, Data.Functor.Coproduct, Data.Functor.Product, Data.FunctorWithIndex, Data.Identity, Data.Maybe, Data.Maybe.First, Data.Maybe.Last, Data.Monoid.Additive, Data.Monoid.Conj, Data.Monoid.Disj, Data.Monoid.Dual, Data.Monoid.Multiplicative, Data.Tuple, Data.Unit, Prelude
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Bifunctor/index.php';
require_once __DIR__ . '/../Data.Const/index.php';
require_once __DIR__ . '/../Data.Either/index.php';
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
$ffi_Data_FunctorWithIndex = \call_user_func(function() {
  $exports = [];
$mapWithIndexArray = function($f, $xs = null) use (&$mapWithIndexArray) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$mapWithIndexArray) {
            return $mapWithIndexArray(...\array_merge($__args, $more));
        };
    }
    
    $len = \count($xs);
    $result = array_fill(0, $len, null);
    for ($i = 0; $i < $len; $i++) {
        $f1 = $f($i);
        $result[$i] = $f1($xs[$i]);
    }
    return $result;
};

$exports['mapWithIndexArray'] = $mapWithIndexArray;

return $exports;
  return $exports;
});
function majData_majFunctormajWithmajIndex_mapmajWithmajIndexmajArray($v0, $v1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majFunctormajWithmajIndex_mapmajWithmajIndexmajArray';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Data_FunctorWithIndex;
  $f = ($ffi_Data_FunctorWithIndex['mapWithIndexArray'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Data_FunctorWithIndex_mapWithIndexArray'] = __NAMESPACE__ . '\\majData_majFunctormajWithmajIndex_mapmajWithmajIndexmajArray';





// Data_FunctorWithIndex_mapWithIndex
function majData_majFunctormajWithmajIndex_mapmajWithmajIndex($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctormajWithmajIndex_mapmajWithmajIndex';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'mapWithIndex'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_FunctorWithIndex_mapWithIndex'] = __NAMESPACE__ . '\\majData_majFunctormajWithmajIndex_mapmajWithmajIndex';

// Data_FunctorWithIndex_mapDefault
function majData_majFunctormajWithmajIndex_mapmajDefault($dictFunctorWithIndex_0, $f_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctormajWithmajIndex_mapmajDefault';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($dictFunctorWithIndex_0)->{'mapWithIndex'})(function($v_2) use ($f_1) {
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
$GLOBALS['Data_FunctorWithIndex_mapDefault'] = __NAMESPACE__ . '\\majData_majFunctormajWithmajIndex_mapmajDefault';

// Data_FunctorWithIndex_functorWithIndexTuple
$GLOBALS['Data_FunctorWithIndex_functorWithIndexTuple'] = (object)["mapWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Tuple_functorTuple'])->{'map'})(($f_0)($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Tuple_functorTuple'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_FunctorWithIndex_functorWithIndexProduct
function majData_majFunctormajWithmajIndex_functormajWithmajIndexmajProduct($dictFunctorWithIndex_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctormajWithmajIndex_functormajWithmajIndexmajProduct';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $functorProduct_1_0 = ($GLOBALS['Data_Functor_Product_functorProduct'])((($dictFunctorWithIndex_0)->{'Functor0'})(null));
  $__res = function($dictFunctorWithIndex1_2) use ($dictFunctorWithIndex_0, $functorProduct_1_0) {
  $__num = \func_num_args();
  $functorProduct1_3_1 = ($functorProduct_1_0)((($dictFunctorWithIndex1_2)->{'Functor0'})(null));
  $__res = (object)["mapWithIndex" => (function() use ($dictFunctorWithIndex1_2, $dictFunctorWithIndex_0) {
  $__fn = function($f_4, $v_5 = null) use ($dictFunctorWithIndex1_2, $dictFunctorWithIndex_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($GLOBALS['Data_Bifunctor_bifunctorTuple'])->{'bimap'})((($dictFunctorWithIndex_0)->{'mapWithIndex'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_4))($GLOBALS['Data_Either_Left']))))((($dictFunctorWithIndex1_2)->{'mapWithIndex'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_4))($GLOBALS['Data_Either_Right']))))($v_5);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_4) use ($functorProduct1_3_1) {
  $__num = \func_num_args();
  $__res = $functorProduct1_3_1;
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
$GLOBALS['Data_FunctorWithIndex_functorWithIndexProduct'] = __NAMESPACE__ . '\\majData_majFunctormajWithmajIndex_functormajWithmajIndexmajProduct';

// Data_FunctorWithIndex_functorWithIndexMultiplicative
$GLOBALS['Data_FunctorWithIndex_functorWithIndexMultiplicative'] = (object)["mapWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Monoid_Multiplicative_functorMultiplicative'])->{'map'})(($f_0)($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Monoid_Multiplicative_functorMultiplicative'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_FunctorWithIndex_functorWithIndexMaybe
$GLOBALS['Data_FunctorWithIndex_functorWithIndexMaybe'] = (object)["mapWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})(($f_0)($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Maybe_functorMaybe'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_FunctorWithIndex_functorWithIndexLast
$GLOBALS['Data_FunctorWithIndex_functorWithIndexLast'] = (object)["mapWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})(($f_0)($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Maybe_functorMaybe'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_FunctorWithIndex_functorWithIndexIdentity
$GLOBALS['Data_FunctorWithIndex_functorWithIndexIdentity'] = (object)["mapWithIndex" => (function() {
  $__fn = function($f_0, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($f_0)($GLOBALS['Data_Unit_unit']))($v_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Identity_functorIdentity'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_FunctorWithIndex_functorWithIndexFirst
$GLOBALS['Data_FunctorWithIndex_functorWithIndexFirst'] = (object)["mapWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})(($f_0)($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Maybe_functorMaybe'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_FunctorWithIndex_functorWithIndexEither
$GLOBALS['Data_FunctorWithIndex_functorWithIndexEither'] = (object)["mapWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Either_functorEither'])->{'map'})(($f_0)($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Either_functorEither'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_FunctorWithIndex_functorWithIndexDual
$GLOBALS['Data_FunctorWithIndex_functorWithIndexDual'] = (object)["mapWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Monoid_Dual_functorDual'])->{'map'})(($f_0)($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Monoid_Dual_functorDual'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_FunctorWithIndex_functorWithIndexDisj
$GLOBALS['Data_FunctorWithIndex_functorWithIndexDisj'] = (object)["mapWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Monoid_Disj_functorDisj'])->{'map'})(($f_0)($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Monoid_Disj_functorDisj'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_FunctorWithIndex_functorWithIndexCoproduct
function majData_majFunctormajWithmajIndex_functormajWithmajIndexmajCoproduct($dictFunctorWithIndex_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctormajWithmajIndex_functormajWithmajIndexmajCoproduct';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $functorCoproduct_1_0 = ($GLOBALS['Data_Functor_Coproduct_functorCoproduct'])((($dictFunctorWithIndex_0)->{'Functor0'})(null));
  $__res = function($dictFunctorWithIndex1_2) use ($dictFunctorWithIndex_0, $functorCoproduct_1_0) {
  $__num = \func_num_args();
  $functorCoproduct1_3_1 = ($functorCoproduct_1_0)((($dictFunctorWithIndex1_2)->{'Functor0'})(null));
  $__res = (object)["mapWithIndex" => (function() use ($dictFunctorWithIndex1_2, $dictFunctorWithIndex_0) {
  $__fn = function($f_4, $v_5 = null) use ($dictFunctorWithIndex1_2, $dictFunctorWithIndex_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($GLOBALS['Data_Bifunctor_bifunctorEither'])->{'bimap'})((($dictFunctorWithIndex_0)->{'mapWithIndex'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_4))($GLOBALS['Data_Either_Left']))))((($dictFunctorWithIndex1_2)->{'mapWithIndex'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_4))($GLOBALS['Data_Either_Right']))))($v_5);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_4) use ($functorCoproduct1_3_1) {
  $__num = \func_num_args();
  $__res = $functorCoproduct1_3_1;
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
$GLOBALS['Data_FunctorWithIndex_functorWithIndexCoproduct'] = __NAMESPACE__ . '\\majData_majFunctormajWithmajIndex_functormajWithmajIndexmajCoproduct';

// Data_FunctorWithIndex_functorWithIndexConst
$GLOBALS['Data_FunctorWithIndex_functorWithIndexConst'] = (object)["mapWithIndex" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = $v1_1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Const_functorConst'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_FunctorWithIndex_functorWithIndexConj
$GLOBALS['Data_FunctorWithIndex_functorWithIndexConj'] = (object)["mapWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Monoid_Conj_functorConj'])->{'map'})(($f_0)($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Monoid_Conj_functorConj'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_FunctorWithIndex_functorWithIndexCompose
function majData_majFunctormajWithmajIndex_functormajWithmajIndexmajCompose($dictFunctorWithIndex_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctormajWithmajIndex_functormajWithmajIndexmajCompose';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictFunctorWithIndex_0)->{'Functor0'})(null);
  $__res = function($dictFunctorWithIndex1_2) use ($__local_var_1_0, $dictFunctorWithIndex_0) {
  $__num = \func_num_args();
  $mapWithIndex2_3_1 = ($dictFunctorWithIndex1_2)->{'mapWithIndex'};
  $__local_var_4_2 = (($dictFunctorWithIndex1_2)->{'Functor0'})(null);
  $functorCompose1_5_3 = (object)["map" => (function() use ($__local_var_1_0, $__local_var_4_2) {
  $__fn = function($f_5, $v_6 = null) use ($__local_var_1_0, $__local_var_4_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_1_0)->{'map'})((($__local_var_4_2)->{'map'})($f_5)))($v_6);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $__res = (object)["mapWithIndex" => (function() use ($dictFunctorWithIndex_0, $mapWithIndex2_3_1) {
  $__fn = function($f_6, $v_7 = null) use ($dictFunctorWithIndex_0, $mapWithIndex2_3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictFunctorWithIndex_0)->{'mapWithIndex'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($mapWithIndex2_3_1))((function() use ($f_6) {
  $__fn = function($a_8, $b_9 = null) use ($f_6, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($f_6)(new \Data\Tuple\Data_Tuple_Tuple($a_8, $b_9));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})())))($v_7);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_6) use ($functorCompose1_5_3) {
  $__num = \func_num_args();
  $__res = $functorCompose1_5_3;
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
$GLOBALS['Data_FunctorWithIndex_functorWithIndexCompose'] = __NAMESPACE__ . '\\majData_majFunctormajWithmajIndex_functormajWithmajIndexmajCompose';

// Data_FunctorWithIndex_functorWithIndexArray
$GLOBALS['Data_FunctorWithIndex_functorWithIndexArray'] = (object)["mapWithIndex" => $GLOBALS['Data_FunctorWithIndex_mapWithIndexArray'], "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Functor_functorArray'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_FunctorWithIndex_functorWithIndexApp
function majData_majFunctormajWithmajIndex_functormajWithmajIndexmajApp($dictFunctorWithIndex_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctormajWithmajIndex_functormajWithmajIndexmajApp';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictFunctorWithIndex_0)->{'Functor0'})(null);
  $__res = (object)["mapWithIndex" => (function() use ($dictFunctorWithIndex_0) {
  $__fn = function($f_2, $v_3 = null) use ($dictFunctorWithIndex_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictFunctorWithIndex_0)->{'mapWithIndex'})($f_2))($v_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = $__local_var_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_FunctorWithIndex_functorWithIndexApp'] = __NAMESPACE__ . '\\majData_majFunctormajWithmajIndex_functormajWithmajIndexmajApp';

// Data_FunctorWithIndex_functorWithIndexAdditive
$GLOBALS['Data_FunctorWithIndex_functorWithIndexAdditive'] = (object)["mapWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Monoid_Additive_functorAdditive'])->{'map'})(($f_0)($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Monoid_Additive_functorAdditive'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

