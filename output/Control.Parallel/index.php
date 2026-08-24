<?php

namespace Control\Parallel;

// ALL IMPORTS: Control.Alternative, Control.Apply, Control.Category, Control.Parallel, Control.Parallel.Class, Control.Semigroupoid, Data.Foldable, Data.Traversable, Prelude, Prim
// TO REQUIRE: Control.Alternative, Control.Apply, Control.Category, Control.Parallel, Control.Parallel.Class, Control.Semigroupoid, Data.Foldable, Data.Traversable, Prelude
require_once __DIR__ . '/../Control.Alternative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Category/index.php';
require_once __DIR__ . '/../Control.Parallel/index.php';
require_once __DIR__ . '/../Control.Parallel.Class/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Foldable/index.php';
require_once __DIR__ . '/../Data.Traversable/index.php';
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




// Control_Parallel_parTraverse_
function majControl_majParallel_parmajTraverse_($dictParallel_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majParallel_parmajTraverse_';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $sequential_1_0 = ($dictParallel_0)->{'sequential'};
  $parallel_2_1 = ($dictParallel_0)->{'parallel'};
  $__res = function($dictApplicative_3) use ($parallel_2_1, $sequential_1_0) {
  $__num = \func_num_args();
  $__res = function($dictFoldable_4) use ($dictApplicative_3, $parallel_2_1, $sequential_1_0) {
  $__num = \func_num_args();
  $__res = function($f_5) use ($dictApplicative_3, $dictFoldable_4, $parallel_2_1, $sequential_1_0) {
  $__num = \func_num_args();
  $__local_var_6_2 = (($dictApplicative_3)->{'Apply0'})(null);
  $Functor0_7_3 = (($__local_var_6_2)->{'Functor0'})(null);
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($sequential_1_0))(((($dictFoldable_4)->{'foldr'})((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($a_8) use ($Functor0_7_3, $__local_var_6_2) {
  $__num = \func_num_args();
  $__res = function($b_9) use ($Functor0_7_3, $__local_var_6_2, $a_8) {
  $__num = \func_num_args();
  $__res = ((($__local_var_6_2)->{'apply'})(((($Functor0_7_3)->{'map'})(function($v_10) {
  $__num = \func_num_args();
  $__res = function($x_11) {
  $__num = \func_num_args();
  $__res = $x_11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($a_8)))($b_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])($parallel_2_1))($f_5))))((($dictApplicative_3)->{'pure'})($GLOBALS['Data_Unit_unit'])));
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
$GLOBALS['Control_Parallel_parTraverse_'] = __NAMESPACE__ . '\\majControl_majParallel_parmajTraverse_';

// Control_Parallel_parTraverse
function majControl_majParallel_parmajTraverse($dictParallel_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majParallel_parmajTraverse';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $sequential_1_0 = ($dictParallel_0)->{'sequential'};
  $parallel_2_1 = ($dictParallel_0)->{'parallel'};
  $__res = function($dictApplicative_3) use ($parallel_2_1, $sequential_1_0) {
  $__num = \func_num_args();
  $__res = function($dictTraversable_4) use ($dictApplicative_3, $parallel_2_1, $sequential_1_0) {
  $__num = \func_num_args();
  $__res = function($f_5) use ($dictApplicative_3, $dictTraversable_4, $parallel_2_1, $sequential_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($sequential_1_0))(((($dictTraversable_4)->{'traverse'})($dictApplicative_3))((($GLOBALS['Control_Semigroupoid_composeImpl'])($parallel_2_1))($f_5)));
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
$GLOBALS['Control_Parallel_parTraverse'] = __NAMESPACE__ . '\\majControl_majParallel_parmajTraverse';

// Control_Parallel_parSequence_
function majControl_majParallel_parmajSequence_($dictParallel_0, $dictApplicative_1 = null, $dictFoldable_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majParallel_parmajSequence_';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__local_var_3_0 = (($dictApplicative_1)->{'Apply0'})(null);
  $Functor0_4_1 = (($__local_var_3_0)->{'Functor0'})(null);
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($dictParallel_0)->{'sequential'}))(((($dictFoldable_2)->{'foldr'})((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($a_5) use ($Functor0_4_1, $__local_var_3_0) {
  $__num = \func_num_args();
  $__res = function($b_6) use ($Functor0_4_1, $__local_var_3_0, $a_5) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_0)->{'apply'})(((($Functor0_4_1)->{'map'})(function($v_7) {
  $__num = \func_num_args();
  $__res = function($x_8) {
  $__num = \func_num_args();
  $__res = $x_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($a_5)))($b_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(($dictParallel_0)->{'parallel'}))(function($x_3) {
  $__num = \func_num_args();
  $__res = $x_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))))((($dictApplicative_1)->{'pure'})($GLOBALS['Data_Unit_unit'])));
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Control_Parallel_parSequence_'] = __NAMESPACE__ . '\\majControl_majParallel_parmajSequence_';

// Control_Parallel_parSequence
function majControl_majParallel_parmajSequence($dictParallel_0, $dictApplicative_1 = null, $dictTraversable_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majParallel_parmajSequence';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($dictParallel_0)->{'sequential'}))(((($dictTraversable_2)->{'traverse'})($dictApplicative_1))((($GLOBALS['Control_Semigroupoid_composeImpl'])(($dictParallel_0)->{'parallel'}))(function($x_3) {
  $__num = \func_num_args();
  $__res = $x_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})));
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Control_Parallel_parSequence'] = __NAMESPACE__ . '\\majControl_majParallel_parmajSequence';

// Control_Parallel_parOneOfMap
function majControl_majParallel_parmajOnemajOfmajMap($dictParallel_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majParallel_parmajOnemajOfmajMap';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $sequential_1_0 = ($dictParallel_0)->{'sequential'};
  $parallel_2_1 = ($dictParallel_0)->{'parallel'};
  $__res = function($dictAlternative_3) use ($parallel_2_1, $sequential_1_0) {
  $__num = \func_num_args();
  $Plus1_4_2 = (($dictAlternative_3)->{'Plus1'})(null);
  $__res = function($dictFoldable_5) use ($Plus1_4_2, $parallel_2_1, $sequential_1_0) {
  $__num = \func_num_args();
  $__res = function($dictFunctor_6) use ($Plus1_4_2, $dictFoldable_5, $parallel_2_1, $sequential_1_0) {
  $__num = \func_num_args();
  $__res = function($f_7) use ($Plus1_4_2, $dictFoldable_5, $parallel_2_1, $sequential_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($sequential_1_0))(((($dictFoldable_5)->{'foldr'})((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($Plus1_4_2)->{'Alt0'})(null))->{'alt'}))((($GLOBALS['Control_Semigroupoid_composeImpl'])($parallel_2_1))($f_7))))(($Plus1_4_2)->{'empty'}));
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Parallel_parOneOfMap'] = __NAMESPACE__ . '\\majControl_majParallel_parmajOnemajOfmajMap';

// Control_Parallel_parOneOf
function majControl_majParallel_parmajOnemajOf($dictParallel_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majParallel_parmajOnemajOf';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $sequential_1_0 = ($dictParallel_0)->{'sequential'};
  $parallel_2_1 = ($dictParallel_0)->{'parallel'};
  $__res = function($dictAlternative_3) use ($parallel_2_1, $sequential_1_0) {
  $__num = \func_num_args();
  $Plus1_4_2 = (($dictAlternative_3)->{'Plus1'})(null);
  $__res = function($dictFoldable_5) use ($Plus1_4_2, $parallel_2_1, $sequential_1_0) {
  $__num = \func_num_args();
  $__res = function($dictFunctor_6) use ($Plus1_4_2, $dictFoldable_5, $parallel_2_1, $sequential_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($sequential_1_0))(((($dictFoldable_5)->{'foldr'})((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($Plus1_4_2)->{'Alt0'})(null))->{'alt'}))($parallel_2_1)))(($Plus1_4_2)->{'empty'}));
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
$GLOBALS['Control_Parallel_parOneOf'] = __NAMESPACE__ . '\\majControl_majParallel_parmajOnemajOf';

// Control_Parallel_parApply
function majControl_majParallel_parmajApply($dictParallel_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majParallel_parmajApply';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Apply1_1_0 = (($dictParallel_0)->{'Apply1'})(null);
  $__res = function($mf_2) use ($Apply1_1_0, $dictParallel_0) {
  $__num = \func_num_args();
  $__res = function($ma_3) use ($Apply1_1_0, $dictParallel_0, $mf_2) {
  $__num = \func_num_args();
  $__res = (($dictParallel_0)->{'sequential'})(((($Apply1_1_0)->{'apply'})((($dictParallel_0)->{'parallel'})($mf_2)))((($dictParallel_0)->{'parallel'})($ma_3)));
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
$GLOBALS['Control_Parallel_parApply'] = __NAMESPACE__ . '\\majControl_majParallel_parmajApply';

