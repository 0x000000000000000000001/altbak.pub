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

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };


// Control_Parallel_parTraverse_
$GLOBALS['Control_Parallel_parTraverse_'] = function($dictParallel_0 = null) {
  $__num = \func_num_args();
  $sequential_1_0 = ($dictParallel_0)['sequential'];
  $parallel_2_1 = ($dictParallel_0)['parallel'];
  $__res = function($dictApplicative_3 = null) use ($parallel_2_1, $sequential_1_0) {
  $__num = \func_num_args();
  $traverse__4_2 = ($GLOBALS['Data_Foldable_traverse_'])($dictApplicative_3);
  $__res = function($dictFoldable_5 = null) use ($parallel_2_1, $sequential_1_0, $traverse__4_2) {
  $__num = \func_num_args();
  $traverse_1_6_3 = ($traverse__4_2)($dictFoldable_5);
  $__res = function($f_7 = null) use ($parallel_2_1, $sequential_1_0, $traverse_1_6_3) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($sequential_1_0))(($traverse_1_6_3)((($GLOBALS['Control_Semigroupoid_composeImpl'])($parallel_2_1))($f_7)));
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

// Control_Parallel_parTraverse
$GLOBALS['Control_Parallel_parTraverse'] = function($dictParallel_0 = null) {
  $__num = \func_num_args();
  $sequential_1_0 = ($dictParallel_0)['sequential'];
  $parallel_2_1 = ($dictParallel_0)['parallel'];
  $__res = (function() use ($parallel_2_1, $sequential_1_0) {
  $__fn = function($dictApplicative_3 = null, $dictTraversable_4 = null) use ($parallel_2_1, $sequential_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $traverse_5_2 = (($dictTraversable_4)['traverse'])($dictApplicative_3);
  $__res = function($f_6 = null) use ($parallel_2_1, $sequential_1_0, $traverse_5_2) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($sequential_1_0))(($traverse_5_2)((($GLOBALS['Control_Semigroupoid_composeImpl'])($parallel_2_1))($f_6)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
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

// Control_Parallel_parSequence_
$GLOBALS['Control_Parallel_parSequence_'] = function($dictParallel_0 = null) {
  $__num = \func_num_args();
  $parTraverse_1_1_0 = ($GLOBALS['Control_Parallel_parTraverse_'])($dictParallel_0);
  $__res = function($dictApplicative_2 = null) use ($parTraverse_1_1_0) {
  $__num = \func_num_args();
  $parTraverse_2_3_1 = ($parTraverse_1_1_0)($dictApplicative_2);
  $__res = function($dictFoldable_4 = null) use ($parTraverse_2_3_1) {
  $__num = \func_num_args();
  $__res = (($parTraverse_2_3_1)($dictFoldable_4))(($GLOBALS['Control_Category_categoryFn'])['identity']);
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

// Control_Parallel_parSequence
$GLOBALS['Control_Parallel_parSequence'] = function($dictParallel_0 = null) {
  $__num = \func_num_args();
  $parTraverse1_1_0 = ($GLOBALS['Control_Parallel_parTraverse'])($dictParallel_0);
  $__res = function($dictApplicative_2 = null) use ($parTraverse1_1_0) {
  $__num = \func_num_args();
  $parTraverse2_3_1 = ($parTraverse1_1_0)($dictApplicative_2);
  $__res = function($dictTraversable_4 = null) use ($parTraverse2_3_1) {
  $__num = \func_num_args();
  $__res = (($parTraverse2_3_1)($dictTraversable_4))(($GLOBALS['Control_Category_categoryFn'])['identity']);
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

// Control_Parallel_parOneOfMap
$GLOBALS['Control_Parallel_parOneOfMap'] = function($dictParallel_0 = null) {
  $__num = \func_num_args();
  $sequential_1_0 = ($dictParallel_0)['sequential'];
  $parallel_2_1 = ($dictParallel_0)['parallel'];
  $__res = function($dictAlternative_3 = null) use ($parallel_2_1, $sequential_1_0) {
  $__num = \func_num_args();
  $Plus1_4_2 = (($dictAlternative_3)['Plus1'])($GLOBALS['Prim_undefined']);
  $__res = function($dictFoldable_5 = null) use ($Plus1_4_2, $parallel_2_1, $sequential_1_0) {
  $__num = \func_num_args();
  $oneOfMap_6_3 = (($GLOBALS['Data_Foldable_oneOfMap'])($dictFoldable_5))($Plus1_4_2);
  $__res = (function() use ($oneOfMap_6_3, $parallel_2_1, $sequential_1_0) {
  $__fn = function($dictFunctor_7 = null, $f_8 = null) use ($oneOfMap_6_3, $parallel_2_1, $sequential_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($sequential_1_0))(($oneOfMap_6_3)((($GLOBALS['Control_Semigroupoid_composeImpl'])($parallel_2_1))($f_8)));
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Parallel_parOneOf
$GLOBALS['Control_Parallel_parOneOf'] = function($dictParallel_0 = null) {
  $__num = \func_num_args();
  $sequential_1_0 = ($dictParallel_0)['sequential'];
  $parallel_2_1 = ($dictParallel_0)['parallel'];
  $__res = function($dictAlternative_3 = null) use ($parallel_2_1, $sequential_1_0) {
  $__num = \func_num_args();
  $Plus1_4_2 = (($dictAlternative_3)['Plus1'])($GLOBALS['Prim_undefined']);
  $__res = function($dictFoldable_5 = null) use ($Plus1_4_2, $parallel_2_1, $sequential_1_0) {
  $__num = \func_num_args();
  $oneOfMap_6_3 = (($GLOBALS['Data_Foldable_oneOfMap'])($dictFoldable_5))($Plus1_4_2);
  $__res = function($dictFunctor_7 = null) use ($oneOfMap_6_3, $parallel_2_1, $sequential_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($sequential_1_0))(($oneOfMap_6_3)($parallel_2_1));
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

// Control_Parallel_parApply
$GLOBALS['Control_Parallel_parApply'] = (function() {
  $__fn = function($dictParallel_0 = null, $mf_1 = null, $ma_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($dictParallel_0)['sequential'])(((((($dictParallel_0)['Apply1'])($GLOBALS['Prim_undefined']))['apply'])((($dictParallel_0)['parallel'])($mf_1)))((($dictParallel_0)['parallel'])($ma_2)));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

