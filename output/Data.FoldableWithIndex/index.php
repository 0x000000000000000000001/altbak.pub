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
if (!\function_exists(__NAMESPACE__ . '\\phpurs_execute_effect')) {
  function phpurs_execute_effect($val) {
    if (\is_callable($val)) {
      return $val($GLOBALS['Data_Unit_unit']);
    }
    return $val;
  }
}

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };




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
  $Functor0_2_1 = (($__local_var_1_0)->{'Functor0'})(null);
  $applySecond_1_0 = function($a_3) use ($Functor0_2_1, $__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($b_4) use ($Functor0_2_1, $__local_var_1_0, $a_3) {
  $__num = \func_num_args();
  $__res = ((($__local_var_1_0)->{'apply'})(((($Functor0_2_1)->{'map'})(function($v_5) {
  $__num = \func_num_args();
  $__res = function($x_6) {
  $__num = \func_num_args();
  $__res = $x_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($a_3)))($b_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = function($dictFoldableWithIndex_2) use ($applySecond_1_0, $dictApplicative_0) {
  $__num = \func_num_args();
  $__res = function($f_3) use ($applySecond_1_0, $dictApplicative_0, $dictFoldableWithIndex_2) {
  $__num = \func_num_args();
  $__res = ((($dictFoldableWithIndex_2)->{'foldrWithIndex'})(function($i_4) use ($applySecond_1_0, $f_3) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($applySecond_1_0))(($f_3)($i_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($dictApplicative_0)->{'pure'})($GLOBALS['Data_Unit_unit']));
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
$GLOBALS['Data_FoldableWithIndex_traverseWithIndex_'] = __NAMESPACE__ . '\\majData_majFoldablemajWithmajIndex_traversemajWithmajIndex_';

// Data_FoldableWithIndex_forWithIndex_
function majData_majFoldablemajWithmajIndex_formajWithmajIndex_($dictApplicative_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldablemajWithmajIndex_formajWithmajIndex_';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictApplicative_0)->{'Apply0'})(null);
  $Functor0_2_1 = (($__local_var_1_0)->{'Functor0'})(null);
  $applySecond_1_0 = function($a_3) use ($Functor0_2_1, $__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($b_4) use ($Functor0_2_1, $__local_var_1_0, $a_3) {
  $__num = \func_num_args();
  $__res = ((($__local_var_1_0)->{'apply'})(((($Functor0_2_1)->{'map'})(function($v_5) {
  $__num = \func_num_args();
  $__res = function($x_6) {
  $__num = \func_num_args();
  $__res = $x_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($a_3)))($b_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = function($dictFoldableWithIndex_2) use ($applySecond_1_0, $dictApplicative_0) {
  $__num = \func_num_args();
  $__res = function($b_3) use ($applySecond_1_0, $dictApplicative_0, $dictFoldableWithIndex_2) {
  $__num = \func_num_args();
  $__res = function($a_4) use ($applySecond_1_0, $b_3, $dictApplicative_0, $dictFoldableWithIndex_2) {
  $__num = \func_num_args();
  $__res = (((($dictFoldableWithIndex_2)->{'foldrWithIndex'})(function($i_5) use ($a_4, $applySecond_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($applySecond_1_0))(($a_4)($i_5));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($dictApplicative_0)->{'pure'})($GLOBALS['Data_Unit_unit'])))($b_3);
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
$GLOBALS['Data_FoldableWithIndex_foldableWithIndexTuple'] = (object)["foldrWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($z_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($f_0, $z_1) {
  $__num = \func_num_args();
  $__res = ((($f_0)($GLOBALS['Data_Unit_unit']))(($v_2)->{'value1'}))($z_1);
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
}, "foldlWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($z_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($f_0, $z_1) {
  $__num = \func_num_args();
  $__res = ((($f_0)($GLOBALS['Data_Unit_unit']))($z_1))(($v_2)->{'value1'});
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
}, "foldMapWithIndex" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $__res = function($f_1) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($f_1) {
  $__num = \func_num_args();
  $__res = (($f_1)($GLOBALS['Data_Unit_unit']))(($v_2)->{'value1'});
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
}, "Foldable0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableTuple'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_FoldableWithIndex_foldableWithIndexMultiplicative
$GLOBALS['Data_FoldableWithIndex_foldableWithIndexMultiplicative'] = (object)["foldrWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__local_var_1_0 = ($f_0)($GLOBALS['Data_Unit_unit']);
  $__res = function($z_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($__local_var_1_0, $z_2) {
  $__num = \func_num_args();
  $__res = (($__local_var_1_0)($v_3))($z_2);
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
}, "foldlWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__local_var_1_1 = ($f_0)($GLOBALS['Data_Unit_unit']);
  $__res = function($z_2) use ($__local_var_1_1) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($__local_var_1_1, $z_2) {
  $__num = \func_num_args();
  $__res = (($__local_var_1_1)($z_2))($v_3);
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
}, "foldMapWithIndex" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $__res = function($f_1) {
  $__num = \func_num_args();
  $__local_var_2_2 = ($f_1)($GLOBALS['Data_Unit_unit']);
  $__res = function($v_3) use ($__local_var_2_2) {
  $__num = \func_num_args();
  $__res = ($__local_var_2_2)($v_3);
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
}, "Foldable0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableMultiplicative'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_FoldableWithIndex_foldableWithIndexMaybe
$GLOBALS['Data_FoldableWithIndex_foldableWithIndexMaybe'] = (object)["foldrWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__local_var_1_0 = ($f_0)($GLOBALS['Data_Unit_unit']);
  $__res = function($v1_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v2_3) use ($__local_var_1_0, $v1_2) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v2_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t1 = $v1_2;
goto end_branch_1;;
};
  if ($v2_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = (($__local_var_1_0)(($v2_3)->{'value0'}))($v1_2);
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
}, "foldlWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__local_var_1_2 = ($f_0)($GLOBALS['Data_Unit_unit']);
  $__res = function($v1_2) use ($__local_var_1_2) {
  $__num = \func_num_args();
  $__res = function($v2_3) use ($__local_var_1_2, $v1_2) {
  $__num = \func_num_args();
  $__t3 = null;;
  if ($v2_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t3 = $v1_2;
goto end_branch_3;;
};
  if ($v2_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t3 = (($__local_var_1_2)($v1_2))(($v2_3)->{'value0'});
goto end_branch_3;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t3 = null;
  end_branch_3:;
  $__res = $__t3;
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
}, "foldMapWithIndex" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $mempty_1_4 = ($dictMonoid_0)->{'mempty'};
  $foldMap8_1_4 = function($v_2) use ($mempty_1_4) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($mempty_1_4, $v_2) {
  $__num = \func_num_args();
  $__t5 = null;;
  if ($v1_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t5 = $mempty_1_4;
goto end_branch_5;;
};
  if ($v1_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t5 = ($v_2)(($v1_3)->{'value0'});
goto end_branch_5;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t5 = null;
  end_branch_5:;
  $__res = $__t5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = function($f_2) use ($foldMap8_1_4) {
  $__num = \func_num_args();
  $__res = ($foldMap8_1_4)(($f_2)($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableMaybe'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_FoldableWithIndex_foldableWithIndexLast
$GLOBALS['Data_FoldableWithIndex_foldableWithIndexLast'] = (object)["foldrWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__local_var_1_0 = ($f_0)($GLOBALS['Data_Unit_unit']);
  $__res = function($z_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($__local_var_1_0, $z_2) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t1 = $z_2;
goto end_branch_1;;
};
  if ($v_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = (($__local_var_1_0)(($v_3)->{'value0'}))($z_2);
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
}, "foldlWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__local_var_1_2 = ($f_0)($GLOBALS['Data_Unit_unit']);
  $__res = function($z_2) use ($__local_var_1_2) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($__local_var_1_2, $z_2) {
  $__num = \func_num_args();
  $__t3 = null;;
  if ($v_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t3 = $z_2;
goto end_branch_3;;
};
  if ($v_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t3 = (($__local_var_1_2)($z_2))(($v_3)->{'value0'});
goto end_branch_3;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t3 = null;
  end_branch_3:;
  $__res = $__t3;
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
}, "foldMapWithIndex" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $__res = function($f_1) use ($dictMonoid_0) {
  $__num = \func_num_args();
  $__local_var_2_4 = ($f_1)($GLOBALS['Data_Unit_unit']);
  $__res = function($v_3) use ($__local_var_2_4, $dictMonoid_0) {
  $__num = \func_num_args();
  $__t5 = null;;
  if ($v_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t5 = ($dictMonoid_0)->{'mempty'};
goto end_branch_5;;
};
  if ($v_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t5 = ($__local_var_2_4)(($v_3)->{'value0'});
goto end_branch_5;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t5 = null;
  end_branch_5:;
  $__res = $__t5;
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
}, "Foldable0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableLast'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_FoldableWithIndex_foldableWithIndexIdentity
$GLOBALS['Data_FoldableWithIndex_foldableWithIndexIdentity'] = (object)["foldrWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($z_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($f_0, $z_1) {
  $__num = \func_num_args();
  $__res = ((($f_0)($GLOBALS['Data_Unit_unit']))($v_2))($z_1);
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
}, "foldlWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($z_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($f_0, $z_1) {
  $__num = \func_num_args();
  $__res = ((($f_0)($GLOBALS['Data_Unit_unit']))($z_1))($v_2);
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
}, "foldMapWithIndex" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $__res = function($f_1) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($f_1) {
  $__num = \func_num_args();
  $__res = (($f_1)($GLOBALS['Data_Unit_unit']))($v_2);
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
}, "Foldable0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableIdentity'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_FoldableWithIndex_foldableWithIndexFirst
$GLOBALS['Data_FoldableWithIndex_foldableWithIndexFirst'] = (object)["foldrWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__local_var_1_0 = ($f_0)($GLOBALS['Data_Unit_unit']);
  $__res = function($z_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($__local_var_1_0, $z_2) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t1 = $z_2;
goto end_branch_1;;
};
  if ($v_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = (($__local_var_1_0)(($v_3)->{'value0'}))($z_2);
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
}, "foldlWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__local_var_1_2 = ($f_0)($GLOBALS['Data_Unit_unit']);
  $__res = function($z_2) use ($__local_var_1_2) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($__local_var_1_2, $z_2) {
  $__num = \func_num_args();
  $__t3 = null;;
  if ($v_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t3 = $z_2;
goto end_branch_3;;
};
  if ($v_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t3 = (($__local_var_1_2)($z_2))(($v_3)->{'value0'});
goto end_branch_3;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t3 = null;
  end_branch_3:;
  $__res = $__t3;
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
}, "foldMapWithIndex" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $__res = function($f_1) use ($dictMonoid_0) {
  $__num = \func_num_args();
  $__local_var_2_4 = ($f_1)($GLOBALS['Data_Unit_unit']);
  $__res = function($v_3) use ($__local_var_2_4, $dictMonoid_0) {
  $__num = \func_num_args();
  $__t5 = null;;
  if ($v_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t5 = ($dictMonoid_0)->{'mempty'};
goto end_branch_5;;
};
  if ($v_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t5 = ($__local_var_2_4)(($v_3)->{'value0'});
goto end_branch_5;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t5 = null;
  end_branch_5:;
  $__res = $__t5;
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
}, "Foldable0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableFirst'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_FoldableWithIndex_foldableWithIndexEither
$GLOBALS['Data_FoldableWithIndex_foldableWithIndexEither'] = (object)["foldrWithIndex" => function($v_0) {
  $__num = \func_num_args();
  $__res = function($v1_1) use ($v_0) {
  $__num = \func_num_args();
  $__res = function($v2_2) use ($v1_1, $v_0) {
  $__num = \func_num_args();
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldlWithIndex" => function($v_0) {
  $__num = \func_num_args();
  $__res = function($v1_1) use ($v_0) {
  $__num = \func_num_args();
  $__res = function($v2_2) use ($v1_1, $v_0) {
  $__num = \func_num_args();
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldMapWithIndex" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $mempty_1_2 = ($dictMonoid_0)->{'mempty'};
  $__res = function($v_2) use ($mempty_1_2) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($mempty_1_2, $v_2) {
  $__num = \func_num_args();
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableEither'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_FoldableWithIndex_foldableWithIndexDual
$GLOBALS['Data_FoldableWithIndex_foldableWithIndexDual'] = (object)["foldrWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__local_var_1_0 = ($f_0)($GLOBALS['Data_Unit_unit']);
  $__res = function($z_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($__local_var_1_0, $z_2) {
  $__num = \func_num_args();
  $__res = (($__local_var_1_0)($v_3))($z_2);
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
}, "foldlWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__local_var_1_1 = ($f_0)($GLOBALS['Data_Unit_unit']);
  $__res = function($z_2) use ($__local_var_1_1) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($__local_var_1_1, $z_2) {
  $__num = \func_num_args();
  $__res = (($__local_var_1_1)($z_2))($v_3);
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
}, "foldMapWithIndex" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $__res = function($f_1) {
  $__num = \func_num_args();
  $__local_var_2_2 = ($f_1)($GLOBALS['Data_Unit_unit']);
  $__res = function($v_3) use ($__local_var_2_2) {
  $__num = \func_num_args();
  $__res = ($__local_var_2_2)($v_3);
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
}, "Foldable0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableDual'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_FoldableWithIndex_foldableWithIndexDisj
$GLOBALS['Data_FoldableWithIndex_foldableWithIndexDisj'] = (object)["foldrWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__local_var_1_0 = ($f_0)($GLOBALS['Data_Unit_unit']);
  $__res = function($z_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($__local_var_1_0, $z_2) {
  $__num = \func_num_args();
  $__res = (($__local_var_1_0)($v_3))($z_2);
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
}, "foldlWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__local_var_1_1 = ($f_0)($GLOBALS['Data_Unit_unit']);
  $__res = function($z_2) use ($__local_var_1_1) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($__local_var_1_1, $z_2) {
  $__num = \func_num_args();
  $__res = (($__local_var_1_1)($z_2))($v_3);
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
}, "foldMapWithIndex" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $__res = function($f_1) {
  $__num = \func_num_args();
  $__local_var_2_2 = ($f_1)($GLOBALS['Data_Unit_unit']);
  $__res = function($v_3) use ($__local_var_2_2) {
  $__num = \func_num_args();
  $__res = ($__local_var_2_2)($v_3);
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
}, "Foldable0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableDisj'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_FoldableWithIndex_foldableWithIndexConst
$GLOBALS['Data_FoldableWithIndex_foldableWithIndexConst'] = (object)["foldrWithIndex" => function($v_0) {
  $__num = \func_num_args();
  $__res = function($z_1) {
  $__num = \func_num_args();
  $__res = function($v1_2) use ($z_1) {
  $__num = \func_num_args();
  $__res = $z_1;
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
}, "foldlWithIndex" => function($v_0) {
  $__num = \func_num_args();
  $__res = function($z_1) {
  $__num = \func_num_args();
  $__res = function($v1_2) use ($z_1) {
  $__num = \func_num_args();
  $__res = $z_1;
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
}, "foldMapWithIndex" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $mempty_1_0 = ($dictMonoid_0)->{'mempty'};
  $__res = function($v_2) use ($mempty_1_0) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($mempty_1_0) {
  $__num = \func_num_args();
  $__res = $mempty_1_0;
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
}, "Foldable0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableConst'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_FoldableWithIndex_foldableWithIndexConj
$GLOBALS['Data_FoldableWithIndex_foldableWithIndexConj'] = (object)["foldrWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__local_var_1_0 = ($f_0)($GLOBALS['Data_Unit_unit']);
  $__res = function($z_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($__local_var_1_0, $z_2) {
  $__num = \func_num_args();
  $__res = (($__local_var_1_0)($v_3))($z_2);
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
}, "foldlWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__local_var_1_1 = ($f_0)($GLOBALS['Data_Unit_unit']);
  $__res = function($z_2) use ($__local_var_1_1) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($__local_var_1_1, $z_2) {
  $__num = \func_num_args();
  $__res = (($__local_var_1_1)($z_2))($v_3);
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
}, "foldMapWithIndex" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $__res = function($f_1) {
  $__num = \func_num_args();
  $__local_var_2_2 = ($f_1)($GLOBALS['Data_Unit_unit']);
  $__res = function($v_3) use ($__local_var_2_2) {
  $__num = \func_num_args();
  $__res = ($__local_var_2_2)($v_3);
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
}, "Foldable0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableConj'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_FoldableWithIndex_foldableWithIndexAdditive
$GLOBALS['Data_FoldableWithIndex_foldableWithIndexAdditive'] = (object)["foldrWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__local_var_1_0 = ($f_0)($GLOBALS['Data_Unit_unit']);
  $__res = function($z_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($__local_var_1_0, $z_2) {
  $__num = \func_num_args();
  $__res = (($__local_var_1_0)($v_3))($z_2);
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
}, "foldlWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__local_var_1_1 = ($f_0)($GLOBALS['Data_Unit_unit']);
  $__res = function($z_2) use ($__local_var_1_1) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($__local_var_1_1, $z_2) {
  $__num = \func_num_args();
  $__res = (($__local_var_1_1)($z_2))($v_3);
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
}, "foldMapWithIndex" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $__res = function($f_1) {
  $__num = \func_num_args();
  $__local_var_2_2 = ($f_1)($GLOBALS['Data_Unit_unit']);
  $__res = function($v_3) use ($__local_var_2_2) {
  $__num = \func_num_args();
  $__res = ($__local_var_2_2)($v_3);
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
}, "Foldable0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableAdditive'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_FoldableWithIndex_foldWithIndexM
function majData_majFoldablemajWithmajIndex_foldmajWithmajIndexmajM($dictFoldableWithIndex_0, $dictMonad_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldablemajWithmajIndex_foldmajWithmajIndexmajM';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $Bind1_2_0 = (($dictMonad_1)->{'Bind1'})(null);
  $Applicative0_3_1 = (($dictMonad_1)->{'Applicative0'})(null);
  $__res = function($f_4) use ($Applicative0_3_1, $Bind1_2_0, $dictFoldableWithIndex_0) {
  $__num = \func_num_args();
  $__res = function($a0_5) use ($Applicative0_3_1, $Bind1_2_0, $dictFoldableWithIndex_0, $f_4) {
  $__num = \func_num_args();
  $__res = ((($dictFoldableWithIndex_0)->{'foldlWithIndex'})(function($i_6) use ($Bind1_2_0, $f_4) {
  $__num = \func_num_args();
  $__res = function($ma_7) use ($Bind1_2_0, $f_4, $i_6) {
  $__num = \func_num_args();
  $__res = function($b_8) use ($Bind1_2_0, $f_4, $i_6, $ma_7) {
  $__num = \func_num_args();
  $__local_var_9_2 = ($f_4)($i_6);
  $__res = ((($Bind1_2_0)->{'bind'})($ma_7))(function($a_10) use ($__local_var_9_2, $b_8) {
  $__num = \func_num_args();
  $__res = (($__local_var_9_2)($a_10))($b_8);
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($Applicative0_3_1)->{'pure'})($a0_5));
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
$GLOBALS['Data_FoldableWithIndex_foldWithIndexM'] = __NAMESPACE__ . '\\majData_majFoldablemajWithmajIndex_foldmajWithmajIndexmajM';

// Data_FoldableWithIndex_foldMapWithIndexDefaultR
function majData_majFoldablemajWithmajIndex_foldmajMapmajWithmajIndexmajDefaultmajR($dictFoldableWithIndex_0, $dictMonoid_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldablemajWithmajIndex_foldmajMapmajWithmajIndexmajDefaultmajR';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $Semigroup0_2_0 = (($dictMonoid_1)->{'Semigroup0'})(null);
  $mempty_3_1 = ($dictMonoid_1)->{'mempty'};
  $__res = function($f_4) use ($Semigroup0_2_0, $dictFoldableWithIndex_0, $mempty_3_1) {
  $__num = \func_num_args();
  $__res = ((($dictFoldableWithIndex_0)->{'foldrWithIndex'})(function($i_5) use ($Semigroup0_2_0, $f_4) {
  $__num = \func_num_args();
  $__res = function($x_6) use ($Semigroup0_2_0, $f_4, $i_5) {
  $__num = \func_num_args();
  $__res = function($acc_7) use ($Semigroup0_2_0, $f_4, $i_5, $x_6) {
  $__num = \func_num_args();
  $__res = ((($Semigroup0_2_0)->{'append'})((($f_4)($i_5))($x_6)))($acc_7);
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
}))($mempty_3_1);
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
$GLOBALS['Data_FoldableWithIndex_foldableWithIndexArray'] = (object)["foldrWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($z_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Foldable_foldrArray'])(function($v_2) use ($f_0) {
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
}))($z_1)))(($GLOBALS['Data_FunctorWithIndex_mapWithIndexArray'])($GLOBALS['Data_Tuple_Tuple']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldlWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($z_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Foldable_foldlArray'])(function($y_2) use ($f_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($f_0, $y_2) {
  $__num = \func_num_args();
  $__res = ((($f_0)(($v_3)->{'value0'}))($y_2))(($v_3)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($z_1)))(($GLOBALS['Data_FunctorWithIndex_mapWithIndexArray'])($GLOBALS['Data_Tuple_Tuple']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldMapWithIndex" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $Semigroup0_1_2 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $mempty_2_3 = ($dictMonoid_0)->{'mempty'};
  $__res = function($f_3) use ($Semigroup0_1_2, $mempty_2_3) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_FoldableWithIndex_foldableWithIndexArray'])->{'foldrWithIndex'})(function($i_4) use ($Semigroup0_1_2, $f_3) {
  $__num = \func_num_args();
  $__res = function($x_5) use ($Semigroup0_1_2, $f_3, $i_4) {
  $__num = \func_num_args();
  $__res = function($acc_6) use ($Semigroup0_1_2, $f_3, $i_4, $x_5) {
  $__num = \func_num_args();
  $__res = ((($Semigroup0_1_2)->{'append'})((($f_3)($i_4))($x_5)))($acc_6);
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
}))($mempty_2_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable0" => function($_dollar___unused_0) {
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
  $Semigroup0_2_0 = (($dictMonoid_1)->{'Semigroup0'})(null);
  $mempty_3_1 = ($dictMonoid_1)->{'mempty'};
  $__res = function($f_4) use ($Semigroup0_2_0, $dictFoldableWithIndex_0, $mempty_3_1) {
  $__num = \func_num_args();
  $__res = ((($dictFoldableWithIndex_0)->{'foldlWithIndex'})(function($i_5) use ($Semigroup0_2_0, $f_4) {
  $__num = \func_num_args();
  $__res = function($acc_6) use ($Semigroup0_2_0, $f_4, $i_5) {
  $__num = \func_num_args();
  $__res = function($x_7) use ($Semigroup0_2_0, $acc_6, $f_4, $i_5) {
  $__num = \func_num_args();
  $__res = ((($Semigroup0_2_0)->{'append'})($acc_6))((($f_4)($i_5))($x_7));
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
}))($mempty_3_1);
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
  $foldableApp_1_0 = (object)["foldr" => function($f_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($i_3) use ($__local_var_1_0, $f_2) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($__local_var_1_0, $f_2, $i_3) {
  $__num = \func_num_args();
  $__res = (((($__local_var_1_0)->{'foldr'})($f_2))($i_3))($v_4);
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
}, "foldl" => function($f_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($i_3) use ($__local_var_1_0, $f_2) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($__local_var_1_0, $f_2, $i_3) {
  $__num = \func_num_args();
  $__res = (((($__local_var_1_0)->{'foldl'})($f_2))($i_3))($v_4);
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
}, "foldMap" => function($dictMonoid_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($f_3) use ($__local_var_1_0, $dictMonoid_2) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($__local_var_1_0, $dictMonoid_2, $f_3) {
  $__num = \func_num_args();
  $__res = (((($__local_var_1_0)->{'foldMap'})($dictMonoid_2))($f_3))($v_4);
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
  $__res = (object)["foldrWithIndex" => function($f_2) use ($dictFoldableWithIndex_0) {
  $__num = \func_num_args();
  $__res = function($z_3) use ($dictFoldableWithIndex_0, $f_2) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($dictFoldableWithIndex_0, $f_2, $z_3) {
  $__num = \func_num_args();
  $__res = (((($dictFoldableWithIndex_0)->{'foldrWithIndex'})($f_2))($z_3))($v_4);
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
}, "foldlWithIndex" => function($f_2) use ($dictFoldableWithIndex_0) {
  $__num = \func_num_args();
  $__res = function($z_3) use ($dictFoldableWithIndex_0, $f_2) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($dictFoldableWithIndex_0, $f_2, $z_3) {
  $__num = \func_num_args();
  $__res = (((($dictFoldableWithIndex_0)->{'foldlWithIndex'})($f_2))($z_3))($v_4);
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
}, "foldMapWithIndex" => function($dictMonoid_2) use ($dictFoldableWithIndex_0) {
  $__num = \func_num_args();
  $__res = function($f_3) use ($dictFoldableWithIndex_0, $dictMonoid_2) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($dictFoldableWithIndex_0, $dictMonoid_2, $f_3) {
  $__num = \func_num_args();
  $__res = (((($dictFoldableWithIndex_0)->{'foldMapWithIndex'})($dictMonoid_2))($f_3))($v_4);
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
}, "Foldable0" => function($_dollar___unused_2) use ($foldableApp_1_0) {
  $__num = \func_num_args();
  $__res = $foldableApp_1_0;
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
  $foldlWithIndex1_3_1 = ($dictFoldableWithIndex1_2)->{'foldlWithIndex'};
  $__local_var_4_2 = (($dictFoldableWithIndex1_2)->{'Foldable0'})(null);
  $foldableCompose1_4_2 = (object)["foldr" => function($f_5) use ($__local_var_1_0, $__local_var_4_2) {
  $__num = \func_num_args();
  $__res = function($i_6) use ($__local_var_1_0, $__local_var_4_2, $f_5) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($__local_var_1_0, $__local_var_4_2, $f_5, $i_6) {
  $__num = \func_num_args();
  $__local_var_8_3 = (($__local_var_4_2)->{'foldr'})($f_5);
  $__res = (((($__local_var_1_0)->{'foldr'})(function($b_9) use ($__local_var_8_3) {
  $__num = \func_num_args();
  $__res = function($a_10) use ($__local_var_8_3, $b_9) {
  $__num = \func_num_args();
  $__res = (($__local_var_8_3)($a_10))($b_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($i_6))($v_7);
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
}, "foldl" => function($f_5) use ($__local_var_1_0, $__local_var_4_2) {
  $__num = \func_num_args();
  $__res = function($i_6) use ($__local_var_1_0, $__local_var_4_2, $f_5) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($__local_var_1_0, $__local_var_4_2, $f_5, $i_6) {
  $__num = \func_num_args();
  $__res = (((($__local_var_1_0)->{'foldl'})((($__local_var_4_2)->{'foldl'})($f_5)))($i_6))($v_7);
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
}, "foldMap" => function($dictMonoid_5) use ($__local_var_1_0, $__local_var_4_2) {
  $__num = \func_num_args();
  $__res = function($f_6) use ($__local_var_1_0, $__local_var_4_2, $dictMonoid_5) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($__local_var_1_0, $__local_var_4_2, $dictMonoid_5, $f_6) {
  $__num = \func_num_args();
  $__res = (((($__local_var_1_0)->{'foldMap'})($dictMonoid_5))(((($__local_var_4_2)->{'foldMap'})($dictMonoid_5))($f_6)))($v_7);
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
  $__res = (object)["foldrWithIndex" => function($f_5) use ($dictFoldableWithIndex1_2, $dictFoldableWithIndex_0) {
  $__num = \func_num_args();
  $__res = function($i_6) use ($dictFoldableWithIndex1_2, $dictFoldableWithIndex_0, $f_5) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($dictFoldableWithIndex1_2, $dictFoldableWithIndex_0, $f_5, $i_6) {
  $__num = \func_num_args();
  $__res = (((($dictFoldableWithIndex_0)->{'foldrWithIndex'})(function($a_8) use ($dictFoldableWithIndex1_2, $f_5) {
  $__num = \func_num_args();
  $__local_var_9_5 = (($dictFoldableWithIndex1_2)->{'foldrWithIndex'})(function($b_9) use ($a_8, $f_5) {
  $__num = \func_num_args();
  $__res = ($f_5)(new \Data\Tuple\Data_Tuple_Tuple($a_8, $b_9));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($b_10) use ($__local_var_9_5) {
  $__num = \func_num_args();
  $__res = function($a_11) use ($__local_var_9_5, $b_10) {
  $__num = \func_num_args();
  $__res = (($__local_var_9_5)($a_11))($b_10);
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
}))($i_6))($v_7);
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
}, "foldlWithIndex" => function($f_5) use ($dictFoldableWithIndex_0, $foldlWithIndex1_3_1) {
  $__num = \func_num_args();
  $__res = function($i_6) use ($dictFoldableWithIndex_0, $f_5, $foldlWithIndex1_3_1) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($dictFoldableWithIndex_0, $f_5, $foldlWithIndex1_3_1, $i_6) {
  $__num = \func_num_args();
  $__res = (((($dictFoldableWithIndex_0)->{'foldlWithIndex'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($foldlWithIndex1_3_1))(function($a_8) use ($f_5) {
  $__num = \func_num_args();
  $__res = function($b_9) use ($a_8, $f_5) {
  $__num = \func_num_args();
  $__res = ($f_5)(new \Data\Tuple\Data_Tuple_Tuple($a_8, $b_9));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))($i_6))($v_7);
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
}, "foldMapWithIndex" => function($dictMonoid_5) use ($dictFoldableWithIndex1_2, $dictFoldableWithIndex_0) {
  $__num = \func_num_args();
  $foldMapWithIndex2_6_6 = (($dictFoldableWithIndex1_2)->{'foldMapWithIndex'})($dictMonoid_5);
  $__res = function($f_7) use ($dictFoldableWithIndex_0, $dictMonoid_5, $foldMapWithIndex2_6_6) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($dictFoldableWithIndex_0, $dictMonoid_5, $f_7, $foldMapWithIndex2_6_6) {
  $__num = \func_num_args();
  $__res = (((($dictFoldableWithIndex_0)->{'foldMapWithIndex'})($dictMonoid_5))((($GLOBALS['Control_Semigroupoid_composeImpl'])($foldMapWithIndex2_6_6))(function($a_9) use ($f_7) {
  $__num = \func_num_args();
  $__res = function($b_10) use ($a_9, $f_7) {
  $__num = \func_num_args();
  $__res = ($f_7)(new \Data\Tuple\Data_Tuple_Tuple($a_9, $b_10));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))($v_8);
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
}, "Foldable0" => function($_dollar___unused_5) use ($foldableCompose1_4_2) {
  $__num = \func_num_args();
  $__res = $foldableCompose1_4_2;
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
  $__local_var_1_0 = (($dictFoldableWithIndex_0)->{'Foldable0'})(null);
  $__res = function($dictFoldableWithIndex1_2) use ($__local_var_1_0, $dictFoldableWithIndex_0) {
  $__num = \func_num_args();
  $__local_var_3_1 = (($dictFoldableWithIndex1_2)->{'Foldable0'})(null);
  $foldableCoproduct1_3_1 = (object)["foldr" => function($f_4) use ($__local_var_1_0, $__local_var_3_1) {
  $__num = \func_num_args();
  $__res = function($z_5) use ($__local_var_1_0, $__local_var_3_1, $f_4) {
  $__num = \func_num_args();
  $__local_var_6_2 = ((($__local_var_1_0)->{'foldr'})($f_4))($z_5);
  $__local_var_7_3 = ((($__local_var_3_1)->{'foldr'})($f_4))($z_5);
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldl" => function($f_4) use ($__local_var_1_0, $__local_var_3_1) {
  $__num = \func_num_args();
  $__res = function($z_5) use ($__local_var_1_0, $__local_var_3_1, $f_4) {
  $__num = \func_num_args();
  $__local_var_6_5 = ((($__local_var_1_0)->{'foldl'})($f_4))($z_5);
  $__local_var_7_6 = ((($__local_var_3_1)->{'foldl'})($f_4))($z_5);
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldMap" => function($dictMonoid_4) use ($__local_var_1_0, $__local_var_3_1) {
  $__num = \func_num_args();
  $__res = function($f_5) use ($__local_var_1_0, $__local_var_3_1, $dictMonoid_4) {
  $__num = \func_num_args();
  $__local_var_6_8 = ((($__local_var_1_0)->{'foldMap'})($dictMonoid_4))($f_5);
  $__local_var_7_9 = ((($__local_var_3_1)->{'foldMap'})($dictMonoid_4))($f_5);
  $__res = function($v2_8) use ($__local_var_6_8, $__local_var_7_9) {
  $__num = \func_num_args();
  $__t10 = null;;
  if ($v2_8 instanceof \Data\Either\Data_Either_Left) {
$__t10 = ($__local_var_6_8)(($v2_8)->{'value0'});
goto end_branch_10;;
};
  if ($v2_8 instanceof \Data\Either\Data_Either_Right) {
$__t10 = ($__local_var_7_9)(($v2_8)->{'value0'});
goto end_branch_10;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t10 = null;
  end_branch_10:;
  $__res = $__t10;
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
  $__res = (object)["foldrWithIndex" => function($f_4) use ($dictFoldableWithIndex1_2, $dictFoldableWithIndex_0) {
  $__num = \func_num_args();
  $__res = function($z_5) use ($dictFoldableWithIndex1_2, $dictFoldableWithIndex_0, $f_4) {
  $__num = \func_num_args();
  $__local_var_6_12 = ((($dictFoldableWithIndex_0)->{'foldrWithIndex'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_4))($GLOBALS['Data_Either_Left'])))($z_5);
  $__local_var_7_13 = ((($dictFoldableWithIndex1_2)->{'foldrWithIndex'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_4))($GLOBALS['Data_Either_Right'])))($z_5);
  $__res = function($v2_8) use ($__local_var_6_12, $__local_var_7_13) {
  $__num = \func_num_args();
  $__t14 = null;;
  if ($v2_8 instanceof \Data\Either\Data_Either_Left) {
$__t14 = ($__local_var_6_12)(($v2_8)->{'value0'});
goto end_branch_14;;
};
  if ($v2_8 instanceof \Data\Either\Data_Either_Right) {
$__t14 = ($__local_var_7_13)(($v2_8)->{'value0'});
goto end_branch_14;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t14 = null;
  end_branch_14:;
  $__res = $__t14;
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
}, "foldlWithIndex" => function($f_4) use ($dictFoldableWithIndex1_2, $dictFoldableWithIndex_0) {
  $__num = \func_num_args();
  $__res = function($z_5) use ($dictFoldableWithIndex1_2, $dictFoldableWithIndex_0, $f_4) {
  $__num = \func_num_args();
  $__local_var_6_15 = ((($dictFoldableWithIndex_0)->{'foldlWithIndex'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_4))($GLOBALS['Data_Either_Left'])))($z_5);
  $__local_var_7_16 = ((($dictFoldableWithIndex1_2)->{'foldlWithIndex'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_4))($GLOBALS['Data_Either_Right'])))($z_5);
  $__res = function($v2_8) use ($__local_var_6_15, $__local_var_7_16) {
  $__num = \func_num_args();
  $__t17 = null;;
  if ($v2_8 instanceof \Data\Either\Data_Either_Left) {
$__t17 = ($__local_var_6_15)(($v2_8)->{'value0'});
goto end_branch_17;;
};
  if ($v2_8 instanceof \Data\Either\Data_Either_Right) {
$__t17 = ($__local_var_7_16)(($v2_8)->{'value0'});
goto end_branch_17;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t17 = null;
  end_branch_17:;
  $__res = $__t17;
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
}, "foldMapWithIndex" => function($dictMonoid_4) use ($dictFoldableWithIndex1_2, $dictFoldableWithIndex_0) {
  $__num = \func_num_args();
  $__res = function($f_5) use ($dictFoldableWithIndex1_2, $dictFoldableWithIndex_0, $dictMonoid_4) {
  $__num = \func_num_args();
  $__local_var_6_18 = ((($dictFoldableWithIndex_0)->{'foldMapWithIndex'})($dictMonoid_4))((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_5))($GLOBALS['Data_Either_Left']));
  $__local_var_7_19 = ((($dictFoldableWithIndex1_2)->{'foldMapWithIndex'})($dictMonoid_4))((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_5))($GLOBALS['Data_Either_Right']));
  $__res = function($v2_8) use ($__local_var_6_18, $__local_var_7_19) {
  $__num = \func_num_args();
  $__t20 = null;;
  if ($v2_8 instanceof \Data\Either\Data_Either_Left) {
$__t20 = ($__local_var_6_18)(($v2_8)->{'value0'});
goto end_branch_20;;
};
  if ($v2_8 instanceof \Data\Either\Data_Either_Right) {
$__t20 = ($__local_var_7_19)(($v2_8)->{'value0'});
goto end_branch_20;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t20 = null;
  end_branch_20:;
  $__res = $__t20;
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
}, "Foldable0" => function($_dollar___unused_4) use ($foldableCoproduct1_3_1) {
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
  $__local_var_1_0 = (($dictFoldableWithIndex_0)->{'Foldable0'})(null);
  $__res = function($dictFoldableWithIndex1_2) use ($__local_var_1_0, $dictFoldableWithIndex_0) {
  $__num = \func_num_args();
  $__local_var_3_1 = (($dictFoldableWithIndex1_2)->{'Foldable0'})(null);
  $foldableProduct1_3_1 = (object)["foldr" => function($f_4) use ($__local_var_1_0, $__local_var_3_1) {
  $__num = \func_num_args();
  $__res = function($z_5) use ($__local_var_1_0, $__local_var_3_1, $f_4) {
  $__num = \func_num_args();
  $__res = function($v_6) use ($__local_var_1_0, $__local_var_3_1, $f_4, $z_5) {
  $__num = \func_num_args();
  $__res = (((($__local_var_1_0)->{'foldr'})($f_4))((((($__local_var_3_1)->{'foldr'})($f_4))($z_5))(($v_6)->{'value1'})))(($v_6)->{'value0'});
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
}, "foldl" => function($f_4) use ($__local_var_1_0, $__local_var_3_1) {
  $__num = \func_num_args();
  $__res = function($z_5) use ($__local_var_1_0, $__local_var_3_1, $f_4) {
  $__num = \func_num_args();
  $__res = function($v_6) use ($__local_var_1_0, $__local_var_3_1, $f_4, $z_5) {
  $__num = \func_num_args();
  $__res = (((($__local_var_3_1)->{'foldl'})($f_4))((((($__local_var_1_0)->{'foldl'})($f_4))($z_5))(($v_6)->{'value0'})))(($v_6)->{'value1'});
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
}, "foldMap" => function($dictMonoid_4) use ($__local_var_1_0, $__local_var_3_1) {
  $__num = \func_num_args();
  $Semigroup0_5_2 = (($dictMonoid_4)->{'Semigroup0'})(null);
  $__res = function($f_6) use ($Semigroup0_5_2, $__local_var_1_0, $__local_var_3_1, $dictMonoid_4) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($Semigroup0_5_2, $__local_var_1_0, $__local_var_3_1, $dictMonoid_4, $f_6) {
  $__num = \func_num_args();
  $__res = ((($Semigroup0_5_2)->{'append'})((((($__local_var_1_0)->{'foldMap'})($dictMonoid_4))($f_6))(($v_7)->{'value0'})))((((($__local_var_3_1)->{'foldMap'})($dictMonoid_4))($f_6))(($v_7)->{'value1'}));
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
  $__res = (object)["foldrWithIndex" => function($f_4) use ($dictFoldableWithIndex1_2, $dictFoldableWithIndex_0) {
  $__num = \func_num_args();
  $__res = function($z_5) use ($dictFoldableWithIndex1_2, $dictFoldableWithIndex_0, $f_4) {
  $__num = \func_num_args();
  $__res = function($v_6) use ($dictFoldableWithIndex1_2, $dictFoldableWithIndex_0, $f_4, $z_5) {
  $__num = \func_num_args();
  $__res = (((($dictFoldableWithIndex_0)->{'foldrWithIndex'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_4))($GLOBALS['Data_Either_Left'])))((((($dictFoldableWithIndex1_2)->{'foldrWithIndex'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_4))($GLOBALS['Data_Either_Right'])))($z_5))(($v_6)->{'value1'})))(($v_6)->{'value0'});
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
}, "foldlWithIndex" => function($f_4) use ($dictFoldableWithIndex1_2, $dictFoldableWithIndex_0) {
  $__num = \func_num_args();
  $__res = function($z_5) use ($dictFoldableWithIndex1_2, $dictFoldableWithIndex_0, $f_4) {
  $__num = \func_num_args();
  $__res = function($v_6) use ($dictFoldableWithIndex1_2, $dictFoldableWithIndex_0, $f_4, $z_5) {
  $__num = \func_num_args();
  $__res = (((($dictFoldableWithIndex1_2)->{'foldlWithIndex'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_4))($GLOBALS['Data_Either_Right'])))((((($dictFoldableWithIndex_0)->{'foldlWithIndex'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_4))($GLOBALS['Data_Either_Left'])))($z_5))(($v_6)->{'value0'})))(($v_6)->{'value1'});
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
}, "foldMapWithIndex" => function($dictMonoid_4) use ($dictFoldableWithIndex1_2, $dictFoldableWithIndex_0) {
  $__num = \func_num_args();
  $Semigroup0_5_4 = (($dictMonoid_4)->{'Semigroup0'})(null);
  $__res = function($f_6) use ($Semigroup0_5_4, $dictFoldableWithIndex1_2, $dictFoldableWithIndex_0, $dictMonoid_4) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($Semigroup0_5_4, $dictFoldableWithIndex1_2, $dictFoldableWithIndex_0, $dictMonoid_4, $f_6) {
  $__num = \func_num_args();
  $__res = ((($Semigroup0_5_4)->{'append'})((((($dictFoldableWithIndex_0)->{'foldMapWithIndex'})($dictMonoid_4))((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_6))($GLOBALS['Data_Either_Left'])))(($v_7)->{'value0'})))((((($dictFoldableWithIndex1_2)->{'foldMapWithIndex'})($dictMonoid_4))((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_6))($GLOBALS['Data_Either_Right'])))(($v_7)->{'value1'}));
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
}, "Foldable0" => function($_dollar___unused_4) use ($foldableProduct1_3_1) {
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
function majData_majFoldablemajWithmajIndex_foldlmajWithmajIndexmajDefault($dictFoldableWithIndex_0, $c_1 = null, $u_2 = null, $xs_3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldablemajWithmajIndex_foldlmajWithmajIndexmajDefault';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $semigroupEndo1_4_0 = (object)["append" => function($v_4) {
  $__num = \func_num_args();
  $__res = function($v1_5) use ($v_4) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($v_4))($v1_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_5_1 = (object)["mempty" => function($x_5) {
  $__num = \func_num_args();
  $__res = $x_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Semigroup0" => function($_dollar___unused_5) use ($semigroupEndo1_4_0) {
  $__num = \func_num_args();
  $__res = $semigroupEndo1_4_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_6_2 = (($__local_var_5_1)->{'Semigroup0'})(null);
  $semigroupDual1_6_2 = (object)["append" => function($v_7) use ($__local_var_6_2) {
  $__num = \func_num_args();
  $__res = function($v1_8) use ($__local_var_6_2, $v_7) {
  $__num = \func_num_args();
  $__res = ((($__local_var_6_2)->{'append'})($v1_8))($v_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ((((($dictFoldableWithIndex_0)->{'foldMapWithIndex'})((object)["mempty" => ($__local_var_5_1)->{'mempty'}, "Semigroup0" => function($_dollar___unused_7) use ($semigroupDual1_6_2) {
  $__num = \func_num_args();
  $__res = $semigroupDual1_6_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]))(function($i_4) use ($c_1) {
  $__num = \func_num_args();
  $__local_var_5_4 = ($c_1)($i_4);
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_5) {
  $__num = \func_num_args();
  $__res = $x_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_5) {
  $__num = \func_num_args();
  $__res = $x_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($b_6) use ($__local_var_5_4) {
  $__num = \func_num_args();
  $__res = function($a_7) use ($__local_var_5_4, $b_6) {
  $__num = \func_num_args();
  $__res = (($__local_var_5_4)($a_7))($b_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($xs_3))($u_2);
  goto __end;;
  __end:
  return 4 < $__num ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
}
$GLOBALS['Data_FoldableWithIndex_foldlWithIndexDefault'] = __NAMESPACE__ . '\\majData_majFoldablemajWithmajIndex_foldlmajWithmajIndexmajDefault';

// Data_FoldableWithIndex_foldrWithIndexDefault
function majData_majFoldablemajWithmajIndex_foldrmajWithmajIndexmajDefault($dictFoldableWithIndex_0, $c_1 = null, $u_2 = null, $xs_3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldablemajWithmajIndex_foldrmajWithmajIndexmajDefault';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $semigroupEndo1_4_0 = (object)["append" => function($v_4) {
  $__num = \func_num_args();
  $__res = function($v1_5) use ($v_4) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($v_4))($v1_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ((((($dictFoldableWithIndex_0)->{'foldMapWithIndex'})((object)["mempty" => function($x_5) {
  $__num = \func_num_args();
  $__res = $x_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Semigroup0" => function($_dollar___unused_5) use ($semigroupEndo1_4_0) {
  $__num = \func_num_args();
  $__res = $semigroupEndo1_4_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]))(function($i_4) use ($c_1) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_5) {
  $__num = \func_num_args();
  $__res = $x_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($c_1)($i_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($xs_3))($u_2);
  goto __end;;
  __end:
  return 4 < $__num ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
}
$GLOBALS['Data_FoldableWithIndex_foldrWithIndexDefault'] = __NAMESPACE__ . '\\majData_majFoldablemajWithmajIndex_foldrmajWithmajIndexmajDefault';

// Data_FoldableWithIndex_surroundMapWithIndex
function majData_majFoldablemajWithmajIndex_surroundmajMapmajWithmajIndex($dictFoldableWithIndex_0, $dictSemigroup_1 = null, $d_2 = null, $t_3 = null, $f_4 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldablemajWithmajIndex_surroundmajMapmajWithmajIndex';
  if ($__num < 5) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 5);
  }
  $semigroupEndo1_5_0 = (object)["append" => function($v_5) {
  $__num = \func_num_args();
  $__res = function($v1_6) use ($v_5) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($v_5))($v1_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ((((($dictFoldableWithIndex_0)->{'foldMapWithIndex'})((object)["mempty" => function($x_6) {
  $__num = \func_num_args();
  $__res = $x_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Semigroup0" => function($_dollar___unused_6) use ($semigroupEndo1_5_0) {
  $__num = \func_num_args();
  $__res = $semigroupEndo1_5_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]))(function($i_5) use ($d_2, $dictSemigroup_1, $t_3) {
  $__num = \func_num_args();
  $__res = function($a_6) use ($d_2, $dictSemigroup_1, $i_5, $t_3) {
  $__num = \func_num_args();
  $__res = function($m_7) use ($a_6, $d_2, $dictSemigroup_1, $i_5, $t_3) {
  $__num = \func_num_args();
  $__res = ((($dictSemigroup_1)->{'append'})($d_2))(((($dictSemigroup_1)->{'append'})((($t_3)($i_5))($a_6)))($m_7));
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
}))($f_4))($d_2);
  goto __end;;
  __end:
  return 5 < $__num ? $__res(...\array_slice(\func_get_args(), 5)) : $__res;
}
$GLOBALS['Data_FoldableWithIndex_surroundMapWithIndex'] = __NAMESPACE__ . '\\majData_majFoldablemajWithmajIndex_surroundmajMapmajWithmajIndex';

// Data_FoldableWithIndex_foldMapDefault
function majData_majFoldablemajWithmajIndex_foldmajMapmajDefault($dictFoldableWithIndex_0, $dictMonoid_1 = null, $f_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldablemajWithmajIndex_foldmajMapmajDefault';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictFoldableWithIndex_0)->{'foldMapWithIndex'})($dictMonoid_1))(function($v_3) use ($f_2) {
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
$GLOBALS['Data_FoldableWithIndex_foldMapDefault'] = __NAMESPACE__ . '\\majData_majFoldablemajWithmajIndex_foldmajMapmajDefault';

// Data_FoldableWithIndex_findWithIndex
function majData_majFoldablemajWithmajIndex_findmajWithmajIndex($dictFoldableWithIndex_0, $p_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldablemajWithmajIndex_findmajWithmajIndex';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictFoldableWithIndex_0)->{'foldlWithIndex'})(function($v_2) use ($p_1) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($p_1, $v_2) {
  $__num = \func_num_args();
  $__res = function($v2_4) use ($p_1, $v1_3, $v_2) {
  $__num = \func_num_args();
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(new \Data\Maybe\Data_Maybe_Nothing());
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
  $__res = ((($dictFoldableWithIndex_0)->{'foldlWithIndex'})(function($v_2) use ($f_1) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($f_1, $v_2) {
  $__num = \func_num_args();
  $__res = function($v2_4) use ($f_1, $v1_3, $v_2) {
  $__num = \func_num_args();
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(new \Data\Maybe\Data_Maybe_Nothing());
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
  $semigroupDisj1_2_0 = (object)["append" => function($v_2) use ($dictHeytingAlgebra_1) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($dictHeytingAlgebra_1, $v_2) {
  $__num = \func_num_args();
  $__res = ((($dictHeytingAlgebra_1)->{'disj'})($v_2))($v1_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $monoidDisj_2_0 = (object)["mempty" => ($dictHeytingAlgebra_1)->{'ff'}, "Semigroup0" => function($_dollar___unused_3) use ($semigroupDisj1_2_0) {
  $__num = \func_num_args();
  $__res = $semigroupDisj1_2_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = function($t_3) use ($dictFoldableWithIndex_0, $monoidDisj_2_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Unsafe_Coerce_unsafeCoerce']))(((($dictFoldableWithIndex_0)->{'foldMapWithIndex'})($monoidDisj_2_0))(function($i_4) use ($t_3) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_5) {
  $__num = \func_num_args();
  $__res = $x_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($t_3)($i_4));
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
  $semigroupConj1_2_0 = (object)["append" => function($v_2) use ($dictHeytingAlgebra_1) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($dictHeytingAlgebra_1, $v_2) {
  $__num = \func_num_args();
  $__res = ((($dictHeytingAlgebra_1)->{'conj'})($v_2))($v1_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $monoidConj_2_0 = (object)["mempty" => ($dictHeytingAlgebra_1)->{'tt'}, "Semigroup0" => function($_dollar___unused_3) use ($semigroupConj1_2_0) {
  $__num = \func_num_args();
  $__res = $semigroupConj1_2_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = function($t_3) use ($dictFoldableWithIndex_0, $monoidConj_2_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Unsafe_Coerce_unsafeCoerce']))(((($dictFoldableWithIndex_0)->{'foldMapWithIndex'})($monoidConj_2_0))(function($i_4) use ($t_3) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_5) {
  $__num = \func_num_args();
  $__res = $x_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($t_3)($i_4));
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

