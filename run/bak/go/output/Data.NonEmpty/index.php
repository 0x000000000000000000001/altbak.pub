<?php

namespace Data\NonEmpty;

// ALL IMPORTS: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Plus, Control.Semigroupoid, Data.Eq, Data.Foldable, Data.FoldableWithIndex, Data.Function, Data.Functor, Data.FunctorWithIndex, Data.HeytingAlgebra, Data.Maybe, Data.NonEmpty, Data.Ord, Data.Ordering, Data.Semigroup, Data.Semigroup.Foldable, Data.Show, Data.Traversable, Data.TraversableWithIndex, Data.Tuple, Data.Unfoldable, Data.Unfoldable1, Prelude, Prim
// TO REQUIRE: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Plus, Control.Semigroupoid, Data.Eq, Data.Foldable, Data.FoldableWithIndex, Data.Function, Data.Functor, Data.FunctorWithIndex, Data.HeytingAlgebra, Data.Maybe, Data.NonEmpty, Data.Ord, Data.Ordering, Data.Semigroup, Data.Semigroup.Foldable, Data.Show, Data.Traversable, Data.TraversableWithIndex, Data.Tuple, Data.Unfoldable, Data.Unfoldable1, Prelude
require_once __DIR__ . '/../Control.Alt/index.php';
require_once __DIR__ . '/../Control.Alternative/index.php';
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Plus/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Foldable/index.php';
require_once __DIR__ . '/../Data.FoldableWithIndex/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.FunctorWithIndex/index.php';
require_once __DIR__ . '/../Data.HeytingAlgebra/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.NonEmpty/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ordering/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Semigroup.Foldable/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
require_once __DIR__ . '/../Data.Traversable/index.php';
require_once __DIR__ . '/../Data.TraversableWithIndex/index.php';
require_once __DIR__ . '/../Data.Tuple/index.php';
require_once __DIR__ . '/../Data.Unfoldable/index.php';
require_once __DIR__ . '/../Data.Unfoldable1/index.php';
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
\PhpursThunks::$thunks['Data_NonEmpty_NonEmpty'] = function() { $v = (function() {
  $__fn = function($value0 = null, $value1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("NonEmpty", $value0, $value1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_NonEmpty_unfoldable1NonEmpty'] = function() { $v = function($dictUnfoldable_0 = null) {
  $__num = \func_num_args();
  $__res = ["unfoldr1" => (function() use ($dictUnfoldable_0) {
  $__fn = function($f_1 = null, $b_2 = null) use ($dictUnfoldable_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_3_0 = ($f_1)($b_2);
  $__res = new Phpurs_Data2("NonEmpty", ($__local_var_3_0)->{'value0'}, ((($dictUnfoldable_0)['unfoldr'])(function($v1_4 = null) use ($f_1) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((is_object($v1_4) && (($v1_4)->{'tag'} === "Just"))) {
$__t1 = new Phpurs_Data1("Just", ($f_1)(($v1_4)->{'value0'}));
goto end_branch_1;;
};
  $__t1 = new Phpurs_Data0("Nothing");
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($__local_var_3_0)->{'value1'}));
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
\PhpursThunks::$thunks['Data_NonEmpty_tail'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ($v_0)->{'value1'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_NonEmpty_singleton'] = function() { $v = function($dictPlus_0 = null) {
  $__num = \func_num_args();
  $empty_1_0 = ($dictPlus_0)['empty'];
  $__res = function($a_2 = null) use ($empty_1_0) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("NonEmpty", $a_2, $empty_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_NonEmpty_showNonEmpty'] = function() { $v = (function() {
  $__fn = function($dictShow_0 = null, $dictShow1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ["show" => function($v_2 = null) use ($dictShow1_1, $dictShow_0) {
  $__num = \func_num_args();
  $__res = (((("(NonEmpty " . (($dictShow_0)['show'])(($v_2)->{'value0'})) . " ") . (($dictShow1_1)['show'])(($v_2)->{'value1'})) . ")");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_NonEmpty_semigroupNonEmpty'] = function() { $v = (function() {
  $__fn = function($dictApplicative_0 = null, $dictSemigroup_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ["append" => (function() use ($dictApplicative_0, $dictSemigroup_1) {
  $__fn = function($v_2 = null, $v1_3 = null) use ($dictApplicative_0, $dictSemigroup_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("NonEmpty", ($v_2)->{'value0'}, ((($dictSemigroup_1)['append'])(($v_2)->{'value1'}))(((($dictSemigroup_1)['append'])((($dictApplicative_0)['pure'])(($v1_3)->{'value0'})))(($v1_3)->{'value1'})));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_NonEmpty_oneOf'] = function() { $v = (function() {
  $__fn = function($dictAlternative_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((((($dictAlternative_0)['Plus1'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Alt0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['alt'])((((($dictAlternative_0)['Applicative0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['pure'])(($v_1)->{'value0'})))(($v_1)->{'value1'});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_NonEmpty_head'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ($v_0)->{'value0'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_NonEmpty_functorNonEmpty'] = function() { $v = function($dictFunctor_0 = null) {
  $__num = \func_num_args();
  $__res = ["map" => (function() use ($dictFunctor_0) {
  $__fn = function($f_1 = null, $m_2 = null) use ($dictFunctor_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("NonEmpty", ($f_1)(($m_2)->{'value0'}), ((($dictFunctor_0)['map'])($f_1))(($m_2)->{'value1'}));
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
\PhpursThunks::$thunks['Data_NonEmpty_functorWithIndex'] = function() { $v = function($dictFunctorWithIndex_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($dictFunctorWithIndex_0)['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $functorNonEmpty1_2_1 = ["map" => (function() use ($__local_var_1_0) {
  $__fn = function($f_2 = null, $m_3 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("NonEmpty", ($f_2)(($m_3)->{'value0'}), ((($__local_var_1_0)['map'])($f_2))(($m_3)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $__res = ["mapWithIndex" => (function() use ($dictFunctorWithIndex_0) {
  $__fn = function($f_3 = null, $v_4 = null) use ($dictFunctorWithIndex_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("NonEmpty", (($f_3)(new Phpurs_Data0("Nothing")))(($v_4)->{'value0'}), ((($dictFunctorWithIndex_0)['mapWithIndex'])(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))($f_3))(($GLOBALS['Data_Maybe_Just'] ?? \PhpursThunks::eval('Data_Maybe_Just')))))(($v_4)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($dollar__unused_3 = null) use ($functorNonEmpty1_2_1) {
  $__num = \func_num_args();
  $__res = $functorNonEmpty1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_NonEmpty_fromNonEmpty'] = function() { $v = (function() {
  $__fn = function($f_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($f_0)(($v_1)->{'value0'}))(($v_1)->{'value1'});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_NonEmpty_foldableNonEmpty'] = function() { $v = function($dictFoldable_0 = null) {
  $__num = \func_num_args();
  $__res = ["foldMap" => function($dictMonoid_1 = null) use ($dictFoldable_0) {
  $__num = \func_num_args();
  $foldMap1_2_0 = (($dictFoldable_0)['foldMap'])($dictMonoid_1);
  $__res = (function() use ($dictMonoid_1, $foldMap1_2_0) {
  $__fn = function($f_3 = null, $v_4 = null) use ($dictMonoid_1, $foldMap1_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((($dictMonoid_1)['Semigroup0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['append'])(($f_3)(($v_4)->{'value0'})))((($foldMap1_2_0)($f_3))(($v_4)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldl" => (function() use ($dictFoldable_0) {
  $__fn = function($f_1 = null, $b_2 = null, $v_3 = null) use ($dictFoldable_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($dictFoldable_0)['foldl'])($f_1))((($f_1)($b_2))(($v_3)->{'value0'})))(($v_3)->{'value1'});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldr" => (function() use ($dictFoldable_0) {
  $__fn = function($f_1 = null, $b_2 = null, $v_3 = null) use ($dictFoldable_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($f_1)(($v_3)->{'value0'}))((((($dictFoldable_0)['foldr'])($f_1))($b_2))(($v_3)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_NonEmpty_foldableWithIndexNonEmpty'] = function() { $v = function($dictFoldableWithIndex_0 = null) {
  $__num = \func_num_args();
  $foldableNonEmpty1_1_0 = (($GLOBALS['Data_NonEmpty_foldableNonEmpty'] ?? \PhpursThunks::eval('Data_NonEmpty_foldableNonEmpty')))((($dictFoldableWithIndex_0)['Foldable0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = ["foldMapWithIndex" => function($dictMonoid_2 = null) use ($dictFoldableWithIndex_0) {
  $__num = \func_num_args();
  $foldMapWithIndex1_3_1 = (($dictFoldableWithIndex_0)['foldMapWithIndex'])($dictMonoid_2);
  $__res = (function() use ($dictMonoid_2, $foldMapWithIndex1_3_1) {
  $__fn = function($f_4 = null, $v_5 = null) use ($dictMonoid_2, $foldMapWithIndex1_3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((($dictMonoid_2)['Semigroup0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['append'])((($f_4)(new Phpurs_Data0("Nothing")))(($v_5)->{'value0'})))((($foldMapWithIndex1_3_1)(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))($f_4))(($GLOBALS['Data_Maybe_Just'] ?? \PhpursThunks::eval('Data_Maybe_Just')))))(($v_5)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldlWithIndex" => (function() use ($dictFoldableWithIndex_0) {
  $__fn = function($f_2 = null, $b_3 = null, $v_4 = null) use ($dictFoldableWithIndex_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($dictFoldableWithIndex_0)['foldlWithIndex'])(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))($f_2))(($GLOBALS['Data_Maybe_Just'] ?? \PhpursThunks::eval('Data_Maybe_Just')))))(((($f_2)(new Phpurs_Data0("Nothing")))($b_3))(($v_4)->{'value0'})))(($v_4)->{'value1'});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldrWithIndex" => (function() use ($dictFoldableWithIndex_0) {
  $__fn = function($f_2 = null, $b_3 = null, $v_4 = null) use ($dictFoldableWithIndex_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($f_2)(new Phpurs_Data0("Nothing")))(($v_4)->{'value0'}))((((($dictFoldableWithIndex_0)['foldrWithIndex'])(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))($f_2))(($GLOBALS['Data_Maybe_Just'] ?? \PhpursThunks::eval('Data_Maybe_Just')))))($b_3))(($v_4)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "Foldable0" => function($dollar__unused_2 = null) use ($foldableNonEmpty1_1_0) {
  $__num = \func_num_args();
  $__res = $foldableNonEmpty1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_NonEmpty_traversableNonEmpty'] = function() { $v = function($dictTraversable_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($dictTraversable_0)['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $functorNonEmpty1_2_1 = ["map" => (function() use ($__local_var_1_0) {
  $__fn = function($f_2 = null, $m_3 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("NonEmpty", ($f_2)(($m_3)->{'value0'}), ((($__local_var_1_0)['map'])($f_2))(($m_3)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $foldableNonEmpty1_3_2 = (($GLOBALS['Data_NonEmpty_foldableNonEmpty'] ?? \PhpursThunks::eval('Data_NonEmpty_foldableNonEmpty')))((($dictTraversable_0)['Foldable1'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = ["sequence" => function($dictApplicative_4 = null) use ($dictTraversable_0) {
  $__num = \func_num_args();
  $Apply0_5_3 = (($dictApplicative_4)['Apply0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $sequence1_6_4 = (($dictTraversable_0)['sequence'])($dictApplicative_4);
  $__res = function($v_7 = null) use ($Apply0_5_3, $sequence1_6_4) {
  $__num = \func_num_args();
  $__res = ((($Apply0_5_3)['apply'])(((((($Apply0_5_3)['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['map'])(($GLOBALS['Data_NonEmpty_NonEmpty'] ?? \PhpursThunks::eval('Data_NonEmpty_NonEmpty'))))(($v_7)->{'value0'})))(($sequence1_6_4)(($v_7)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "traverse" => function($dictApplicative_4 = null) use ($dictTraversable_0) {
  $__num = \func_num_args();
  $Apply0_5_5 = (($dictApplicative_4)['Apply0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $traverse1_6_6 = (($dictTraversable_0)['traverse'])($dictApplicative_4);
  $__res = (function() use ($Apply0_5_5, $traverse1_6_6) {
  $__fn = function($f_7 = null, $v_8 = null) use ($Apply0_5_5, $traverse1_6_6, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($Apply0_5_5)['apply'])(((((($Apply0_5_5)['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['map'])(($GLOBALS['Data_NonEmpty_NonEmpty'] ?? \PhpursThunks::eval('Data_NonEmpty_NonEmpty'))))(($f_7)(($v_8)->{'value0'}))))((($traverse1_6_6)($f_7))(($v_8)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($dollar__unused_4 = null) use ($functorNonEmpty1_2_1) {
  $__num = \func_num_args();
  $__res = $functorNonEmpty1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($dollar__unused_4 = null) use ($foldableNonEmpty1_3_2) {
  $__num = \func_num_args();
  $__res = $foldableNonEmpty1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_NonEmpty_traversableWithIndexNonEmpty'] = function() { $v = function($dictTraversableWithIndex_0 = null) {
  $__num = \func_num_args();
  $functorWithIndex1_1_0 = (($GLOBALS['Data_NonEmpty_functorWithIndex'] ?? \PhpursThunks::eval('Data_NonEmpty_functorWithIndex')))((($dictTraversableWithIndex_0)['FunctorWithIndex0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $foldableWithIndexNonEmpty1_2_1 = (($GLOBALS['Data_NonEmpty_foldableWithIndexNonEmpty'] ?? \PhpursThunks::eval('Data_NonEmpty_foldableWithIndexNonEmpty')))((($dictTraversableWithIndex_0)['FoldableWithIndex1'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $traversableNonEmpty1_3_2 = (($GLOBALS['Data_NonEmpty_traversableNonEmpty'] ?? \PhpursThunks::eval('Data_NonEmpty_traversableNonEmpty')))((($dictTraversableWithIndex_0)['Traversable2'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = ["traverseWithIndex" => function($dictApplicative_4 = null) use ($dictTraversableWithIndex_0) {
  $__num = \func_num_args();
  $Apply0_5_3 = (($dictApplicative_4)['Apply0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $traverseWithIndex1_6_4 = (($dictTraversableWithIndex_0)['traverseWithIndex'])($dictApplicative_4);
  $__res = (function() use ($Apply0_5_3, $traverseWithIndex1_6_4) {
  $__fn = function($f_7 = null, $v_8 = null) use ($Apply0_5_3, $traverseWithIndex1_6_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($Apply0_5_3)['apply'])(((((($Apply0_5_3)['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['map'])(($GLOBALS['Data_NonEmpty_NonEmpty'] ?? \PhpursThunks::eval('Data_NonEmpty_NonEmpty'))))((($f_7)(new Phpurs_Data0("Nothing")))(($v_8)->{'value0'}))))((($traverseWithIndex1_6_4)(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))($f_7))(($GLOBALS['Data_Maybe_Just'] ?? \PhpursThunks::eval('Data_Maybe_Just')))))(($v_8)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FunctorWithIndex0" => function($dollar__unused_4 = null) use ($functorWithIndex1_1_0) {
  $__num = \func_num_args();
  $__res = $functorWithIndex1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($dollar__unused_4 = null) use ($foldableWithIndexNonEmpty1_2_1) {
  $__num = \func_num_args();
  $__res = $foldableWithIndexNonEmpty1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($dollar__unused_4 = null) use ($traversableNonEmpty1_3_2) {
  $__num = \func_num_args();
  $__res = $traversableNonEmpty1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_NonEmpty_foldable1NonEmpty'] = function() { $v = function($dictFoldable_0 = null) {
  $__num = \func_num_args();
  $foldableNonEmpty1_1_0 = (($GLOBALS['Data_NonEmpty_foldableNonEmpty'] ?? \PhpursThunks::eval('Data_NonEmpty_foldableNonEmpty')))($dictFoldable_0);
  $__res = ["foldMap1" => (function() use ($dictFoldable_0) {
  $__fn = function($dictSemigroup_2 = null, $f_3 = null, $v_4 = null) use ($dictFoldable_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($dictFoldable_0)['foldl'])((function() use ($dictSemigroup_2, $f_3) {
  $__fn = function($s_5 = null, $a1_6 = null) use ($dictSemigroup_2, $f_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictSemigroup_2)['append'])($s_5))(($f_3)($a1_6));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(($f_3)(($v_4)->{'value0'})))(($v_4)->{'value1'});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldr1" => (function() use ($dictFoldable_0) {
  $__fn = function($f_2 = null, $v_3 = null) use ($dictFoldable_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_4_1 = ($f_2)(($v_3)->{'value0'});
  $__local_var_5_2 = (((($dictFoldable_0)['foldr'])(function($a1_5 = null) use ($f_2) {
  $__num = \func_num_args();
  $__local_var_6_2 = ($f_2)($a1_5);
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Maybe_Just'] ?? \PhpursThunks::eval('Data_Maybe_Just'))))(function($v2_7 = null) use ($__local_var_6_2, $a1_5) {
  $__num = \func_num_args();
  $__t3 = null;;
  if ((is_object($v2_7) && (($v2_7)->{'tag'} === "Nothing"))) {
$__t3 = $a1_5;
goto end_branch_3;;
};
  if ((is_object($v2_7) && (($v2_7)->{'tag'} === "Just"))) {
$__t3 = ($__local_var_6_2)(($v2_7)->{'value0'});
goto end_branch_3;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t3 = null;
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(new Phpurs_Data0("Nothing")))(($v_3)->{'value1'});
  $__t5 = null;;
  if ((is_object($__local_var_5_2) && (($__local_var_5_2)->{'tag'} === "Nothing"))) {
$__t5 = ($v_3)->{'value0'};
goto end_branch_5;;
};
  if ((is_object($__local_var_5_2) && (($__local_var_5_2)->{'tag'} === "Just"))) {
$__t5 = ($__local_var_4_1)(($__local_var_5_2)->{'value0'});
goto end_branch_5;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t5 = null;
  end_branch_5:;
  $__res = $__t5;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "foldl1" => (function() use ($dictFoldable_0) {
  $__fn = function($f_2 = null, $v_3 = null) use ($dictFoldable_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($dictFoldable_0)['foldl'])($f_2))(($v_3)->{'value0'}))(($v_3)->{'value1'});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Foldable0" => function($dollar__unused_2 = null) use ($foldableNonEmpty1_1_0) {
  $__num = \func_num_args();
  $__res = $foldableNonEmpty1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_NonEmpty_foldl1'] = function() { $v = function($dictFoldable_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_NonEmpty_foldable1NonEmpty'] ?? \PhpursThunks::eval('Data_NonEmpty_foldable1NonEmpty')))($dictFoldable_0))['foldl1'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_NonEmpty_eqNonEmpty'] = function() { $v = (function() {
  $__fn = function($dictEq1_0 = null, $dictEq_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $eq11_2_0 = (($dictEq1_0)['eq1'])($dictEq_1);
  $__res = ["eq" => (function() use ($dictEq_1, $eq11_2_0) {
  $__fn = function($x_3 = null, $y_4 = null) use ($dictEq_1, $eq11_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($dictEq_1)['eq'])(($x_3)->{'value0'}))(($y_4)->{'value0'}) && (($eq11_2_0)(($x_3)->{'value1'}))(($y_4)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_NonEmpty_ordNonEmpty'] = function() { $v = function($dictOrd1_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($dictOrd1_0)['Eq10'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $__res = function($dictOrd_2 = null) use ($__local_var_1_0, $dictOrd1_0) {
  $__num = \func_num_args();
  $compare11_3_1 = (($dictOrd1_0)['compare1'])($dictOrd_2);
  $__local_var_4_2 = (($dictOrd_2)['Eq0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $eq11_5_3 = (($__local_var_1_0)['eq1'])($__local_var_4_2);
  $eqNonEmpty2_6_4 = ["eq" => (function() use ($__local_var_4_2, $eq11_5_3) {
  $__fn = function($x_6 = null, $y_7 = null) use ($__local_var_4_2, $eq11_5_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($__local_var_4_2)['eq'])(($x_6)->{'value0'}))(($y_7)->{'value0'}) && (($eq11_5_3)(($x_6)->{'value1'}))(($y_7)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $__res = ["compare" => (function() use ($compare11_3_1, $dictOrd_2) {
  $__fn = function($x_7 = null, $y_8 = null) use ($compare11_3_1, $dictOrd_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v_9_5 = ((($dictOrd_2)['compare'])(($x_7)->{'value0'}))(($y_8)->{'value0'});
  $__t6 = null;;
  if ((is_object($v_9_5) && (($v_9_5)->{'tag'} === "LT"))) {
$__t6 = new Phpurs_Data0("LT");
goto end_branch_6;;
};
  if ((is_object($v_9_5) && (($v_9_5)->{'tag'} === "GT"))) {
$__t6 = new Phpurs_Data0("GT");
goto end_branch_6;;
};
  $__t6 = (($compare11_3_1)(($x_7)->{'value1'}))(($y_8)->{'value1'});
  end_branch_6:;
  $__res = $__t6;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($dollar__unused_7 = null) use ($eqNonEmpty2_6_4) {
  $__num = \func_num_args();
  $__res = $eqNonEmpty2_6_4;
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
\PhpursThunks::$thunks['Data_NonEmpty_eq1NonEmpty'] = function() { $v = function($dictEq1_0 = null) {
  $__num = \func_num_args();
  $__res = ["eq1" => function($dictEq_1 = null) use ($dictEq1_0) {
  $__num = \func_num_args();
  $eq11_2_0 = (($dictEq1_0)['eq1'])($dictEq_1);
  $__res = (function() use ($dictEq_1, $eq11_2_0) {
  $__fn = function($x_3 = null, $y_4 = null) use ($dictEq_1, $eq11_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($dictEq_1)['eq'])(($x_3)->{'value0'}))(($y_4)->{'value0'}) && (($eq11_2_0)(($x_3)->{'value1'}))(($y_4)->{'value1'}));
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_NonEmpty_ord1NonEmpty'] = function() { $v = function($dictOrd1_0 = null) {
  $__num = \func_num_args();
  $ordNonEmpty1_1_0 = (($GLOBALS['Data_NonEmpty_ordNonEmpty'] ?? \PhpursThunks::eval('Data_NonEmpty_ordNonEmpty')))($dictOrd1_0);
  $__local_var_2_1 = (($dictOrd1_0)['Eq10'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $eq1NonEmpty1_3_2 = ["eq1" => function($dictEq_3 = null) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $eq11_4_2 = (($__local_var_2_1)['eq1'])($dictEq_3);
  $__res = (function() use ($dictEq_3, $eq11_4_2) {
  $__fn = function($x_5 = null, $y_6 = null) use ($dictEq_3, $eq11_4_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($dictEq_3)['eq'])(($x_5)->{'value0'}))(($y_6)->{'value0'}) && (($eq11_4_2)(($x_5)->{'value1'}))(($y_6)->{'value1'}));
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
  $__res = ["compare1" => function($dictOrd_4 = null) use ($ordNonEmpty1_1_0) {
  $__num = \func_num_args();
  $__res = (($ordNonEmpty1_1_0)($dictOrd_4))['compare'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq10" => function($dollar__unused_4 = null) use ($eq1NonEmpty1_3_2) {
  $__num = \func_num_args();
  $__res = $eq1NonEmpty1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };























