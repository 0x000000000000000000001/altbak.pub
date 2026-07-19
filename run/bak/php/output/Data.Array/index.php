<?php

namespace Data\Array;

// ALL IMPORTS: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Lazy, Control.Monad.Rec.Class, Control.Monad.ST, Control.Monad.ST.Internal, Control.Semigroupoid, Data.Array, Data.Array.NonEmpty.Internal, Data.Array.ST, Data.Array.ST.Iterator, Data.Boolean, Data.Eq, Data.Foldable, Data.Function, Data.Function.Uncurried, Data.Functor, Data.FunctorWithIndex, Data.HeytingAlgebra, Data.Maybe, Data.Ord, Data.Ordering, Data.Ring, Data.Semigroup, Data.Semiring, Data.Traversable, Data.Tuple, Data.Unfoldable, Partial.Unsafe, Prelude, Prim
// TO REQUIRE: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Lazy, Control.Monad.Rec.Class, Control.Monad.ST, Control.Monad.ST.Internal, Control.Semigroupoid, Data.Array, Data.Array.NonEmpty.Internal, Data.Array.ST, Data.Array.ST.Iterator, Data.Boolean, Data.Eq, Data.Foldable, Data.Function, Data.Function.Uncurried, Data.Functor, Data.FunctorWithIndex, Data.HeytingAlgebra, Data.Maybe, Data.Ord, Data.Ordering, Data.Ring, Data.Semigroup, Data.Semiring, Data.Traversable, Data.Tuple, Data.Unfoldable, Partial.Unsafe, Prelude
require_once __DIR__ . '/../Control.Alt/index.php';
require_once __DIR__ . '/../Control.Alternative/index.php';
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Category/index.php';
require_once __DIR__ . '/../Control.Lazy/index.php';
require_once __DIR__ . '/../Control.Monad.Rec.Class/index.php';
require_once __DIR__ . '/../Control.Monad.ST/index.php';
require_once __DIR__ . '/../Control.Monad.ST.Internal/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Array/index.php';
require_once __DIR__ . '/../Data.Array.NonEmpty.Internal/index.php';
require_once __DIR__ . '/../Data.Array.ST/index.php';
require_once __DIR__ . '/../Data.Array.ST.Iterator/index.php';
require_once __DIR__ . '/../Data.Boolean/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Foldable/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Function.Uncurried/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.FunctorWithIndex/index.php';
require_once __DIR__ . '/../Data.HeytingAlgebra/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ordering/index.php';
require_once __DIR__ . '/../Data.Ring/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Semiring/index.php';
require_once __DIR__ . '/../Data.Traversable/index.php';
require_once __DIR__ . '/../Data.Tuple/index.php';
require_once __DIR__ . '/../Data.Unfoldable/index.php';
require_once __DIR__ . '/../Partial.Unsafe/index.php';
require_once __DIR__ . '/../Prelude/index.php';

if (!class_exists(__NAMESPACE__ . '\\Phpurs_Data0')) {
  class Phpurs_Data0 { public $tag; public function __construct($t) { $this->tag = $t; } }
  class Phpurs_Data1 { public $tag; public $value0; public function __construct($t, $value0) { $this->tag = $t; $this->value0 = $value0; } }
  class Phpurs_Data2 { public $tag; public $value0, $value1; public function __construct($t, $value0, $value1) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; } }
  class Phpurs_Data3 { public $tag; public $value0, $value1, $value2; public function __construct($t, $value0, $value1, $value2) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; } }
  class Phpurs_Data4 { public $tag; public $value0, $value1, $value2, $value3; public function __construct($t, $value0, $value1, $value2, $value3) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; $this->value3 = $value3; } }
  class Phpurs_Data5 { public $tag; public $value0, $value1, $value2, $value3, $value4; public function __construct($t, $value0, $value1, $value2, $value3, $value4) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; $this->value3 = $value3; $this->value4 = $value4; } }
  class Phpurs_Data6 { public $tag; public $value0, $value1, $value2, $value3, $value4, $value5; public function __construct($t, $value0, $value1, $value2, $value3, $value4, $value5) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; $this->value3 = $value3; $this->value4 = $value4; $this->value5 = $value5; } }
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
\PhpursThunks::$thunks['Data_Array_traverse_'] = function() { $v = (($GLOBALS['Data_Foldable_traverse_'] ?? \PhpursThunks::eval('Data_Foldable_traverse_')))(($GLOBALS['Control_Monad_ST_Internal_applicativeST'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_applicativeST'))); return $v; };
\PhpursThunks::$thunks['Data_Array_void'] = function() { $v = (($GLOBALS['Control_Monad_ST_Internal_map_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_map_')))(function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}); return $v; };
\PhpursThunks::$thunks['Data_Array_fromJust'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_fromJust"), recVars=[];
  if ((is_object($v_0) && (($v_0)->tag === "Just"))) {
$__t0 = ($v_0)->value0;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_zipWith'] = function() { $v = (($GLOBALS['Data_Function_Uncurried_runFn3'] ?? \PhpursThunks::eval('Data_Function_Uncurried_runFn3')))(($GLOBALS['Data_Array_zipWithImpl'] ?? \PhpursThunks::eval('Data_Array_zipWithImpl'))); return $v; };
\PhpursThunks::$thunks['Data_Array_zipWithA'] = function() { $v = function($dictApplicative_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_zipWithA"), recVars=[];
  $sequence1_1_0 = (((($GLOBALS['Data_Traversable_traversableArray'] ?? \PhpursThunks::eval('Data_Traversable_traversableArray')))->traverse)($dictApplicative_0))((($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))->identity);
  $__res = (function() use ($sequence1_1_0) {
  $__fn = function($f_2, $xs_3 = null, $ys_4 = null) use ($sequence1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($sequence1_1_0)((((($GLOBALS['Data_Array_zipWith'] ?? \PhpursThunks::eval('Data_Array_zipWith')))($f_2))($xs_3))($ys_4));
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
\PhpursThunks::$thunks['Data_Array_zip'] = function() { $v = (($GLOBALS['Data_Array_zipWith'] ?? \PhpursThunks::eval('Data_Array_zipWith')))(($GLOBALS['Data_Tuple_Tuple'] ?? \PhpursThunks::eval('Data_Tuple_Tuple'))); return $v; };
\PhpursThunks::$thunks['Data_Array_updateAtIndices'] = function() { $v = function($dictFoldable_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_updateAtIndices"), recVars=[];
  $traverse_1_1_0 = (($GLOBALS['Data_Array_traverse_'] ?? \PhpursThunks::eval('Data_Array_traverse_')))($dictFoldable_0);
  $__res = (function() use ($traverse_1_1_0) {
  $__fn = function($us_2, $xs_3 = null) use ($traverse_1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Control_Monad_ST_Internal_run'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_run')))(((($GLOBALS['Data_Array_ST_withArray'] ?? \PhpursThunks::eval('Data_Array_ST_withArray')))(function($res_4) use ($traverse_1_1_0, $us_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($traverse_1_1_0)(function($v_5) use ($res_4) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (((($GLOBALS['Data_Array_ST_poke'] ?? \PhpursThunks::eval('Data_Array_ST_poke')))(($v_5)->value0))(($v_5)->value1))($res_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($us_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($xs_3));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_updateAt'] = function() { $v = (((($GLOBALS['Data_Function_Uncurried_runFn5'] ?? \PhpursThunks::eval('Data_Function_Uncurried_runFn5')))(($GLOBALS['Data_Array__updateAt'] ?? \PhpursThunks::eval('Data_Array__updateAt'))))(($GLOBALS['Data_Maybe_Just'] ?? \PhpursThunks::eval('Data_Maybe_Just'))))(new Phpurs_Data0("Nothing")); return $v; };
\PhpursThunks::$thunks['Data_Array_unsafeIndex'] = function() { $v = function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_unsafeIndex"), recVars=[];
  $__res = (($GLOBALS['Data_Function_Uncurried_runFn2'] ?? \PhpursThunks::eval('Data_Function_Uncurried_runFn2')))(($GLOBALS['Data_Array_unsafeIndexImpl'] ?? \PhpursThunks::eval('Data_Array_unsafeIndexImpl')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_unsafeIndex1'] = function() { $v = (($GLOBALS['Data_Function_Uncurried_runFn2'] ?? \PhpursThunks::eval('Data_Function_Uncurried_runFn2')))(($GLOBALS['Data_Array_unsafeIndexImpl'] ?? \PhpursThunks::eval('Data_Array_unsafeIndexImpl'))); return $v; };
\PhpursThunks::$thunks['Data_Array_uncons'] = function() { $v = (((($GLOBALS['Data_Function_Uncurried_runFn3'] ?? \PhpursThunks::eval('Data_Function_Uncurried_runFn3')))(($GLOBALS['Data_Array_unconsImpl'] ?? \PhpursThunks::eval('Data_Array_unconsImpl'))))(function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = new Phpurs_Data0("Nothing");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((function() {
  $__fn = function($x_0, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = new Phpurs_Data1("Just", (object)["head" => $x_0, "tail" => $xs_1]);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()); return $v; };
\PhpursThunks::$thunks['Data_Array_toUnfoldable'] = function() { $v = (function() {
  $__fn = function($dictUnfoldable_0, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_toUnfoldable"), recVars=[];
  $len_2_0 = (($GLOBALS['Data_Array_length'] ?? \PhpursThunks::eval('Data_Array_length')))($xs_1);
  $__res = ((($dictUnfoldable_0)->unfoldr)(function($i_3) use ($len_2_0, $xs_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($i_3))($len_2_0)) && (((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($i_3))($len_2_0))->tag === "LT"))) {
$__t1 = new Phpurs_Data1("Just", new Phpurs_Data2("Tuple", (($GLOBALS['Partial_Unsafe__unsafePartial'] ?? \PhpursThunks::eval('Partial_Unsafe__unsafePartial')))(function($dollar__unused_4) use ($i_3, $xs_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Data_Array_unsafeIndex1'] ?? \PhpursThunks::eval('Data_Array_unsafeIndex1')))($xs_1))($i_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), ((($GLOBALS['Data_Semiring_intAdd'] ?? \PhpursThunks::eval('Data_Semiring_intAdd')))($i_3))(1)));
} else {
$__t1 = new Phpurs_Data0("Nothing");
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(0);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_tail'] = function() { $v = (((($GLOBALS['Data_Function_Uncurried_runFn3'] ?? \PhpursThunks::eval('Data_Function_Uncurried_runFn3')))(($GLOBALS['Data_Array_unconsImpl'] ?? \PhpursThunks::eval('Data_Array_unconsImpl'))))(function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = new Phpurs_Data0("Nothing");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((function() {
  $__fn = function($v_0, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = new Phpurs_Data1("Just", $xs_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()); return $v; };
\PhpursThunks::$thunks['Data_Array_sortBy'] = function() { $v = function($comp_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_sortBy"), recVars=[];
  $__res = (((($GLOBALS['Data_Function_Uncurried_runFn3'] ?? \PhpursThunks::eval('Data_Function_Uncurried_runFn3')))(($GLOBALS['Data_Array_sortByImpl'] ?? \PhpursThunks::eval('Data_Array_sortByImpl'))))($comp_0))(function($v_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v_1) && (($v_1)->tag === "GT"))) {
$__t0 = 1;
} else {
if ((is_object($v_1) && (($v_1)->tag === "EQ"))) {
$__t0 = 0;
} else {
if ((is_object($v_1) && (($v_1)->tag === "LT"))) {
$__t0 = ((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))(0))(1);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_sortWith'] = function() { $v = (function() {
  $__fn = function($dictOrd_0, $f_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_sortWith"), recVars=[];
  $__res = (($GLOBALS['Data_Array_sortBy'] ?? \PhpursThunks::eval('Data_Array_sortBy')))((function() use ($dictOrd_0, $f_1) {
  $__fn = function($x_2, $y_3 = null) use ($dictOrd_0, $f_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($dictOrd_0)->compare)(($f_1)($x_2)))(($f_1)($y_3));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})());
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_sort'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_sort"), recVars=[];
  $compare_1_0 = ($dictOrd_0)->compare;
  $__res = function($xs_2) use ($compare_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Data_Array_sortBy'] ?? \PhpursThunks::eval('Data_Array_sortBy')))($compare_1_0))($xs_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_snoc'] = function() { $v = (function() {
  $__fn = function($xs_0, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_snoc"), recVars=[];
  $__res = (($GLOBALS['Control_Monad_ST_Internal_run'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_run')))(((($GLOBALS['Data_Array_ST_withArray'] ?? \PhpursThunks::eval('Data_Array_ST_withArray')))((($GLOBALS['Data_Array_ST_push'] ?? \PhpursThunks::eval('Data_Array_ST_push')))($x_1)))($xs_0));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_slice'] = function() { $v = (($GLOBALS['Data_Function_Uncurried_runFn3'] ?? \PhpursThunks::eval('Data_Function_Uncurried_runFn3')))(($GLOBALS['Data_Array_sliceImpl'] ?? \PhpursThunks::eval('Data_Array_sliceImpl'))); return $v; };
\PhpursThunks::$thunks['Data_Array_splitAt'] = function() { $v = (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_splitAt"), recVars=[];
  if (( ! (is_object((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($v_0))(0)) && (((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($v_0))(0))->tag === "GT")))) {
$__t0 = (object)["before" => [], "after" => $v1_1];
} else {
$__t0 = (object)["before" => (((($GLOBALS['Data_Array_slice'] ?? \PhpursThunks::eval('Data_Array_slice')))(0))($v_0))($v1_1), "after" => (((($GLOBALS['Data_Array_slice'] ?? \PhpursThunks::eval('Data_Array_slice')))($v_0))((($GLOBALS['Data_Array_length'] ?? \PhpursThunks::eval('Data_Array_length')))($v1_1)))($v1_1)];
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_take'] = function() { $v = (function() {
  $__fn = function($n_0, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_take"), recVars=[];
  if ((is_object((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($n_0))(1)) && (((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($n_0))(1))->tag === "LT"))) {
$__t0 = [];
} else {
$__t0 = (((($GLOBALS['Data_Array_slice'] ?? \PhpursThunks::eval('Data_Array_slice')))(0))($n_0))($xs_1);
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_singleton'] = function() { $v = function($a_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_singleton"), recVars=[];
  $__res = [$a_0];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_scanr'] = function() { $v = (($GLOBALS['Data_Function_Uncurried_runFn3'] ?? \PhpursThunks::eval('Data_Function_Uncurried_runFn3')))(($GLOBALS['Data_Array_scanrImpl'] ?? \PhpursThunks::eval('Data_Array_scanrImpl'))); return $v; };
\PhpursThunks::$thunks['Data_Array_scanl'] = function() { $v = (($GLOBALS['Data_Function_Uncurried_runFn3'] ?? \PhpursThunks::eval('Data_Function_Uncurried_runFn3')))(($GLOBALS['Data_Array_scanlImpl'] ?? \PhpursThunks::eval('Data_Array_scanlImpl'))); return $v; };
\PhpursThunks::$thunks['Data_Array_replicate'] = function() { $v = (($GLOBALS['Data_Function_Uncurried_runFn2'] ?? \PhpursThunks::eval('Data_Function_Uncurried_runFn2')))(($GLOBALS['Data_Array_replicateImpl'] ?? \PhpursThunks::eval('Data_Array_replicateImpl'))); return $v; };
\PhpursThunks::$thunks['Data_Array_range'] = function() { $v = (($GLOBALS['Data_Function_Uncurried_runFn2'] ?? \PhpursThunks::eval('Data_Function_Uncurried_runFn2')))(($GLOBALS['Data_Array_rangeImpl'] ?? \PhpursThunks::eval('Data_Array_rangeImpl'))); return $v; };
\PhpursThunks::$thunks['Data_Array_partition'] = function() { $v = (($GLOBALS['Data_Function_Uncurried_runFn2'] ?? \PhpursThunks::eval('Data_Function_Uncurried_runFn2')))(($GLOBALS['Data_Array_partitionImpl'] ?? \PhpursThunks::eval('Data_Array_partitionImpl'))); return $v; };
\PhpursThunks::$thunks['Data_Array_null'] = function() { $v = function($xs_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_null"), recVars=[];
  $__res = ((($GLOBALS['Data_Eq_eqIntImpl'] ?? \PhpursThunks::eval('Data_Eq_eqIntImpl')))((($GLOBALS['Data_Array_length'] ?? \PhpursThunks::eval('Data_Array_length')))($xs_0)))(0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_modifyAtIndices'] = function() { $v = function($dictFoldable_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_modifyAtIndices"), recVars=[];
  $traverse_1_1_0 = (($GLOBALS['Data_Array_traverse_'] ?? \PhpursThunks::eval('Data_Array_traverse_')))($dictFoldable_0);
  $__res = (function() use ($traverse_1_1_0) {
  $__fn = function($is_2, $f_3 = null, $xs_4 = null) use ($traverse_1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Control_Monad_ST_Internal_run'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_run')))(((($GLOBALS['Data_Array_ST_withArray'] ?? \PhpursThunks::eval('Data_Array_ST_withArray')))(function($res_5) use ($f_3, $is_2, $traverse_1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($traverse_1_1_0)(function($i_6) use ($f_3, $res_5) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (((($GLOBALS['Data_Array_ST_modify'] ?? \PhpursThunks::eval('Data_Array_ST_modify')))($i_6))($f_3))($res_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($is_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($xs_4));
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
\PhpursThunks::$thunks['Data_Array_mapWithIndex'] = function() { $v = ($GLOBALS['Data_FunctorWithIndex_mapWithIndexArray'] ?? \PhpursThunks::eval('Data_FunctorWithIndex_mapWithIndexArray')); return $v; };
\PhpursThunks::$thunks['Data_Array_intersperse'] = function() { $v = (function() {
  $__fn = function($a_0, $arr_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_intersperse"), recVars=[];
  $v_2_0 = (($GLOBALS['Data_Array_length'] ?? \PhpursThunks::eval('Data_Array_length')))($arr_1);
  if ((is_object((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($v_2_0))(2)) && (((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($v_2_0))(2))->tag === "LT"))) {
$__t1 = $arr_1;
} else {
$__t1 = (($GLOBALS['Control_Monad_ST_Internal_run'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_run')))(((($GLOBALS['Control_Monad_ST_Internal_bind_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_bind_')))(((($GLOBALS['Control_Monad_ST_Internal_bind_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_bind_')))(($GLOBALS['Data_Array_ST_new'] ?? \PhpursThunks::eval('Data_Array_ST_new'))))(function($out_3) use ($a_0, $arr_1, $v_2_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bind_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_bind_')))(((($GLOBALS['Data_Array_ST_push'] ?? \PhpursThunks::eval('Data_Array_ST_push')))((($GLOBALS['Partial_Unsafe__unsafePartial'] ?? \PhpursThunks::eval('Partial_Unsafe__unsafePartial')))(function($dollar__unused_4) use ($arr_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Data_Array_unsafeIndex1'] ?? \PhpursThunks::eval('Data_Array_unsafeIndex1')))($arr_1))(0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))($out_3)))(function($dollar__unused_4) use ($a_0, $arr_1, $out_3, $v_2_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bind_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_bind_')))((((($GLOBALS['Control_Monad_ST_Internal_for'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_for')))(1))($v_2_0))(function($idx_5) use ($a_0, $arr_1, $out_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bind_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_bind_')))(((($GLOBALS['Data_Array_ST_push'] ?? \PhpursThunks::eval('Data_Array_ST_push')))($a_0))($out_3)))(function($dollar__unused_6) use ($arr_1, $idx_5, $out_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Data_Array_void'] ?? \PhpursThunks::eval('Data_Array_void')))(((($GLOBALS['Data_Array_ST_push'] ?? \PhpursThunks::eval('Data_Array_ST_push')))((($GLOBALS['Partial_Unsafe__unsafePartial'] ?? \PhpursThunks::eval('Partial_Unsafe__unsafePartial')))(function($dollar__unused_7) use ($arr_1, $idx_5) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Data_Array_unsafeIndex1'] ?? \PhpursThunks::eval('Data_Array_unsafeIndex1')))($arr_1))($idx_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))($out_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(function($dollar__unused_5) use ($out_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Control_Monad_ST_Internal_pure_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_pure_')))($out_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(($GLOBALS['Data_Array_ST_unsafeFreeze'] ?? \PhpursThunks::eval('Data_Array_ST_unsafeFreeze'))));
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_intercalate'] = function() { $v = function($dictMonoid_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_intercalate"), recVars=[];
  $__res = ((($GLOBALS['Data_Foldable_intercalate'] ?? \PhpursThunks::eval('Data_Foldable_intercalate')))(($GLOBALS['Data_Foldable_foldableArray'] ?? \PhpursThunks::eval('Data_Foldable_foldableArray'))))($dictMonoid_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_insertAt'] = function() { $v = (((($GLOBALS['Data_Function_Uncurried_runFn5'] ?? \PhpursThunks::eval('Data_Function_Uncurried_runFn5')))(($GLOBALS['Data_Array__insertAt'] ?? \PhpursThunks::eval('Data_Array__insertAt'))))(($GLOBALS['Data_Maybe_Just'] ?? \PhpursThunks::eval('Data_Maybe_Just'))))(new Phpurs_Data0("Nothing")); return $v; };
\PhpursThunks::$thunks['Data_Array_init'] = function() { $v = function($xs_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_init"), recVars=[];
  if (((($GLOBALS['Data_Eq_eqIntImpl'] ?? \PhpursThunks::eval('Data_Eq_eqIntImpl')))((($GLOBALS['Data_Array_length'] ?? \PhpursThunks::eval('Data_Array_length')))($xs_0)))(0)) {
$__t0 = new Phpurs_Data0("Nothing");
} else {
$__t0 = new Phpurs_Data1("Just", (((($GLOBALS['Data_Array_slice'] ?? \PhpursThunks::eval('Data_Array_slice')))(0))(((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))((($GLOBALS['Data_Array_length'] ?? \PhpursThunks::eval('Data_Array_length')))($xs_0)))(1)))($xs_0));
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_index'] = function() { $v = (((($GLOBALS['Data_Function_Uncurried_runFn4'] ?? \PhpursThunks::eval('Data_Function_Uncurried_runFn4')))(($GLOBALS['Data_Array_indexImpl'] ?? \PhpursThunks::eval('Data_Array_indexImpl'))))(($GLOBALS['Data_Maybe_Just'] ?? \PhpursThunks::eval('Data_Maybe_Just'))))(new Phpurs_Data0("Nothing")); return $v; };
\PhpursThunks::$thunks['Data_Array_last'] = function() { $v = function($xs_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_last"), recVars=[];
  $__res = ((($GLOBALS['Data_Array_index'] ?? \PhpursThunks::eval('Data_Array_index')))($xs_0))(((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))((($GLOBALS['Data_Array_length'] ?? \PhpursThunks::eval('Data_Array_length')))($xs_0)))(1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_unsnoc'] = function() { $v = function($xs_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_unsnoc"), recVars=[];
  $__local_var_1_0 = (($GLOBALS['Data_Array_init'] ?? \PhpursThunks::eval('Data_Array_init')))($xs_0);
  if ((is_object($__local_var_1_0) && (($__local_var_1_0)->tag === "Just"))) {
$__local_var_2_2 = ((($GLOBALS['Data_Array_index'] ?? \PhpursThunks::eval('Data_Array_index')))($xs_0))(((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))((($GLOBALS['Data_Array_length'] ?? \PhpursThunks::eval('Data_Array_length')))($xs_0)))(1));
if ((is_object($__local_var_2_2) && (($__local_var_2_2)->tag === "Just"))) {
$__t3 = new Phpurs_Data1("Just", (object)["init" => ($__local_var_1_0)->value0, "last" => ($__local_var_2_2)->value0]);
} else {
$__t3 = new Phpurs_Data0("Nothing");
};
$__t1 = $__t3;
} else {
$__t1 = new Phpurs_Data0("Nothing");
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_modifyAt'] = function() { $v = (function() {
  $__fn = function($i_0, $f_1 = null, $xs_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_modifyAt"), recVars=[];
  $__local_var_3_0 = ((($GLOBALS['Data_Array_index'] ?? \PhpursThunks::eval('Data_Array_index')))($xs_2))($i_0);
  if ((is_object($__local_var_3_0) && (($__local_var_3_0)->tag === "Nothing"))) {
$__t1 = new Phpurs_Data0("Nothing");
} else {
if ((is_object($__local_var_3_0) && (($__local_var_3_0)->tag === "Just"))) {
$__t1 = (((($GLOBALS['Data_Array_updateAt'] ?? \PhpursThunks::eval('Data_Array_updateAt')))($i_0))(($f_1)(($__local_var_3_0)->value0)))($xs_2);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_span'] = function() { $v = (function() {
  $__fn = function($p_0, $arr_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_span"), recVars=[];
  $go_2_0 = null;
  $go_2_0 = function($i_3) use ($arr_1, &$go_2_0, $p_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_2_0"), recVars=["go_2_0"];
  while (true) {
$v_4_1 = ((($GLOBALS['Data_Array_index'] ?? \PhpursThunks::eval('Data_Array_index')))($arr_1))($i_3);
if ((is_object($v_4_1) && (($v_4_1)->tag === "Just"))) {
if (($p_0)(($v_4_1)->value0)) {
$__tco_4 = ((($GLOBALS['Data_Semiring_intAdd'] ?? \PhpursThunks::eval('Data_Semiring_intAdd')))($i_3))(1);
$i_3 = $__tco_4;
continue ;
$__t3 = null;
} else {
$__t3 = new Phpurs_Data1("Just", $i_3);
};
$__t2 = $__t3;
} else {
if ((is_object($v_4_1) && (($v_4_1)->tag === "Nothing"))) {
$__t2 = new Phpurs_Data0("Nothing");
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
};
};
$__res = $__t2;
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $breakIndex_3_5 = ($go_2_0)(0);
  if ((is_object($breakIndex_3_5) && (($breakIndex_3_5)->tag === "Just"))) {
switch (($breakIndex_3_5)->value0) {
case 0:
$__t7 = (object)["init" => [], "rest" => $arr_1];
break;
default:
$__t7 = (object)["init" => (((($GLOBALS['Data_Array_slice'] ?? \PhpursThunks::eval('Data_Array_slice')))(0))(($breakIndex_3_5)->value0))($arr_1), "rest" => (((($GLOBALS['Data_Array_slice'] ?? \PhpursThunks::eval('Data_Array_slice')))(($breakIndex_3_5)->value0))((($GLOBALS['Data_Array_length'] ?? \PhpursThunks::eval('Data_Array_length')))($arr_1)))($arr_1)];
break;
};
$__t6 = $__t7;
} else {
if ((is_object($breakIndex_3_5) && (($breakIndex_3_5)->tag === "Nothing"))) {
$__t6 = (object)["init" => $arr_1, "rest" => []];
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t6 = null;
};
};
  $__res = $__t6;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_takeWhile'] = function() { $v = (function() {
  $__fn = function($p_0, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_takeWhile"), recVars=[];
  $__res = (((($GLOBALS['Data_Array_span'] ?? \PhpursThunks::eval('Data_Array_span')))($p_0))($xs_1))->init;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_unzip'] = function() { $v = function($xs_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_unzip"), recVars=[];
  $__res = (($GLOBALS['Control_Monad_ST_Internal_run'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_run')))(((($GLOBALS['Control_Monad_ST_Internal_bind_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_bind_')))(($GLOBALS['Data_Array_ST_new'] ?? \PhpursThunks::eval('Data_Array_ST_new'))))(function($fsts_1) use ($xs_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bind_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_bind_')))(($GLOBALS['Data_Array_ST_new'] ?? \PhpursThunks::eval('Data_Array_ST_new'))))(function($snds_2) use ($fsts_1, $xs_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bind_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_bind_')))(((($GLOBALS['Control_Monad_ST_Internal_map_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_map_')))((($GLOBALS['Data_Array_ST_Iterator_Iterator'] ?? \PhpursThunks::eval('Data_Array_ST_Iterator_Iterator')))(function($v_3) use ($xs_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Data_Array_index'] ?? \PhpursThunks::eval('Data_Array_index')))($xs_0))($v_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))((($GLOBALS['Control_Monad_ST_Internal_new'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_new')))(0))))(function($iter_3) use ($fsts_1, $snds_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bind_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_bind_')))(((($GLOBALS['Data_Array_ST_Iterator_iterate'] ?? \PhpursThunks::eval('Data_Array_ST_Iterator_iterate')))($iter_3))(function($v_4) use ($fsts_1, $snds_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__local_var_5_0 = ($v_4)->value1;
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bind_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_bind_')))((($GLOBALS['Data_Array_void'] ?? \PhpursThunks::eval('Data_Array_void')))(((($GLOBALS['Data_Array_ST_push'] ?? \PhpursThunks::eval('Data_Array_ST_push')))(($v_4)->value0))($fsts_1))))(function($dollar__unused_6) use ($__local_var_5_0, $snds_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Data_Array_void'] ?? \PhpursThunks::eval('Data_Array_void')))(((($GLOBALS['Data_Array_ST_push'] ?? \PhpursThunks::eval('Data_Array_ST_push')))($__local_var_5_0))($snds_2));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(function($dollar__unused_4) use ($fsts_1, $snds_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bind_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_bind_')))((($GLOBALS['Data_Array_ST_unsafeFreeze'] ?? \PhpursThunks::eval('Data_Array_ST_unsafeFreeze')))($fsts_1)))(function($fsts__prime___5) use ($snds_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bind_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_bind_')))((($GLOBALS['Data_Array_ST_unsafeFreeze'] ?? \PhpursThunks::eval('Data_Array_ST_unsafeFreeze')))($snds_2)))(function($snds__prime___6) use ($fsts__prime___5) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Control_Monad_ST_Internal_pure_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_pure_')))(new Phpurs_Data2("Tuple", $fsts__prime___5, $snds__prime___6));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_head'] = function() { $v = function($xs_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_head"), recVars=[];
  $__res = ((($GLOBALS['Data_Array_index'] ?? \PhpursThunks::eval('Data_Array_index')))($xs_0))(0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_nubBy'] = function() { $v = (function() {
  $__fn = function($comp_0, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_nubBy"), recVars=[];
  $indexedAndSorted_2_0 = ((($GLOBALS['Data_Array_sortBy'] ?? \PhpursThunks::eval('Data_Array_sortBy')))((function() use ($comp_0) {
  $__fn = function($x_2, $y_3 = null) use ($comp_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($comp_0)(($x_2)->value1))(($y_3)->value1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(((($GLOBALS['Data_FunctorWithIndex_mapWithIndexArray'] ?? \PhpursThunks::eval('Data_FunctorWithIndex_mapWithIndexArray')))(($GLOBALS['Data_Tuple_Tuple'] ?? \PhpursThunks::eval('Data_Tuple_Tuple'))))($xs_1));
  $v_3_1 = ((($GLOBALS['Data_Array_index'] ?? \PhpursThunks::eval('Data_Array_index')))($indexedAndSorted_2_0))(0);
  if ((is_object($v_3_1) && (($v_3_1)->tag === "Nothing"))) {
$__t2 = [];
} else {
if ((is_object($v_3_1) && (($v_3_1)->tag === "Just"))) {
$__t2 = ((($GLOBALS['Data_Functor_arrayMap'] ?? \PhpursThunks::eval('Data_Functor_arrayMap')))(($GLOBALS['Data_Tuple_snd'] ?? \PhpursThunks::eval('Data_Tuple_snd'))))((((($GLOBALS['Data_Array_sortWith'] ?? \PhpursThunks::eval('Data_Array_sortWith')))(($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt'))))(($GLOBALS['Data_Tuple_fst'] ?? \PhpursThunks::eval('Data_Tuple_fst'))))((($GLOBALS['Control_Monad_ST_Internal_run'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_run')))(((($GLOBALS['Control_Monad_ST_Internal_bind_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_bind_')))((($GLOBALS['Data_Array_ST_unsafeThaw'] ?? \PhpursThunks::eval('Data_Array_ST_unsafeThaw')))([($v_3_1)->value0])))(function($result_4) use ($comp_0, $indexedAndSorted_2_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bind_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_bind_')))(((($GLOBALS['Control_Monad_ST_Internal_foreach'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_foreach')))($indexedAndSorted_2_0))(function($v1_5) use ($comp_0, $result_4) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__local_var_6_3 = ($v1_5)->value1;
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bind_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_bind_')))(((($GLOBALS['Control_Monad_ST_Internal_map_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_map_')))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Tuple_snd'] ?? \PhpursThunks::eval('Data_Tuple_snd'))))((($GLOBALS['Partial_Unsafe__unsafePartial'] ?? \PhpursThunks::eval('Partial_Unsafe__unsafePartial')))(function($dollar__unused_7) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_fromJust'] ?? \PhpursThunks::eval('Data_Array_fromJust'))))(($GLOBALS['Data_Array_last'] ?? \PhpursThunks::eval('Data_Array_last')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))))((($GLOBALS['Data_Array_ST_unsafeFreeze'] ?? \PhpursThunks::eval('Data_Array_ST_unsafeFreeze')))($result_4))))(function($lst_7) use ($__local_var_6_3, $comp_0, $result_4, $v1_5) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__local_var_8_4 = (($GLOBALS['Data_Array_void'] ?? \PhpursThunks::eval('Data_Array_void')))(((($GLOBALS['Data_Array_ST_push'] ?? \PhpursThunks::eval('Data_Array_ST_push')))($v1_5))($result_4));
  if (((($GLOBALS['Data_Eq_eqBooleanImpl'] ?? \PhpursThunks::eval('Data_Eq_eqBooleanImpl')))((is_object((($comp_0)($lst_7))($__local_var_6_3)) && (((($comp_0)($lst_7))($__local_var_6_3))->tag === "EQ"))))(false)) {
$__t5 = $__local_var_8_4;
} else {
$__t5 = (($GLOBALS['Control_Monad_ST_Internal_pure_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_pure_')))(($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit')));
};
  $__res = $__t5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(function($dollar__unused_5) use ($result_4) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Data_Array_ST_unsafeFreeze'] ?? \PhpursThunks::eval('Data_Array_ST_unsafeFreeze')))($result_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
};
};
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_nub'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_nub"), recVars=[];
  $__res = (($GLOBALS['Data_Array_nubBy'] ?? \PhpursThunks::eval('Data_Array_nubBy')))(($dictOrd_0)->compare);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_groupBy'] = function() { $v = (function() {
  $__fn = function($op_0, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_groupBy"), recVars=[];
  $__res = (($GLOBALS['Control_Monad_ST_Internal_run'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_run')))(((($GLOBALS['Control_Monad_ST_Internal_bind_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_bind_')))(($GLOBALS['Data_Array_ST_new'] ?? \PhpursThunks::eval('Data_Array_ST_new'))))(function($result_2) use ($op_0, $xs_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bind_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_bind_')))(((($GLOBALS['Control_Monad_ST_Internal_map_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_map_')))((($GLOBALS['Data_Array_ST_Iterator_Iterator'] ?? \PhpursThunks::eval('Data_Array_ST_Iterator_Iterator')))(function($v_3) use ($xs_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Data_Array_index'] ?? \PhpursThunks::eval('Data_Array_index')))($xs_1))($v_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))((($GLOBALS['Control_Monad_ST_Internal_new'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_new')))(0))))(function($iter_3) use ($op_0, $result_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bind_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_bind_')))(((($GLOBALS['Data_Array_ST_Iterator_iterate'] ?? \PhpursThunks::eval('Data_Array_ST_Iterator_iterate')))($iter_3))(function($x_4) use ($iter_3, $op_0, $result_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Data_Array_void'] ?? \PhpursThunks::eval('Data_Array_void')))(((($GLOBALS['Control_Monad_ST_Internal_bind_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_bind_')))(($GLOBALS['Data_Array_ST_new'] ?? \PhpursThunks::eval('Data_Array_ST_new'))))(function($sub1_5) use ($iter_3, $op_0, $result_2, $x_4) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bind_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_bind_')))(((($GLOBALS['Data_Array_ST_push'] ?? \PhpursThunks::eval('Data_Array_ST_push')))($x_4))($sub1_5)))(function($dollar__unused_6) use ($iter_3, $op_0, $result_2, $sub1_5, $x_4) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bind_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_bind_')))((((($GLOBALS['Data_Array_ST_Iterator_pushWhile'] ?? \PhpursThunks::eval('Data_Array_ST_Iterator_pushWhile')))(($op_0)($x_4)))($iter_3))($sub1_5)))(function($dollar__unused_7) use ($result_2, $sub1_5) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bind_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_bind_')))((($GLOBALS['Data_Array_ST_unsafeFreeze'] ?? \PhpursThunks::eval('Data_Array_ST_unsafeFreeze')))($sub1_5)))(function($grp_8) use ($result_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Data_Array_ST_push'] ?? \PhpursThunks::eval('Data_Array_ST_push')))($grp_8))($result_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(function($dollar__unused_4) use ($result_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Data_Array_ST_unsafeFreeze'] ?? \PhpursThunks::eval('Data_Array_ST_unsafeFreeze')))($result_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_groupAllBy'] = function() { $v = function($cmp_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_groupAllBy"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_Array_groupBy'] ?? \PhpursThunks::eval('Data_Array_groupBy')))((function() use ($cmp_0) {
  $__fn = function($x_1, $y_2 = null) use ($cmp_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (is_object((($cmp_0)($x_1))($y_2)) && (((($cmp_0)($x_1))($y_2))->tag === "EQ"));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})())))((($GLOBALS['Data_Array_sortBy'] ?? \PhpursThunks::eval('Data_Array_sortBy')))($cmp_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_groupAll'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_groupAll"), recVars=[];
  $__res = (($GLOBALS['Data_Array_groupAllBy'] ?? \PhpursThunks::eval('Data_Array_groupAllBy')))(($dictOrd_0)->compare);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_group'] = function() { $v = function($dictEq_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_group"), recVars=[];
  $eq2_1_0 = ($dictEq_0)->eq;
  $__res = function($xs_2) use ($eq2_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Data_Array_groupBy'] ?? \PhpursThunks::eval('Data_Array_groupBy')))($eq2_1_0))($xs_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_fromFoldable'] = function() { $v = function($dictFoldable_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_fromFoldable"), recVars=[];
  $__res = ((($GLOBALS['Data_Function_Uncurried_runFn2'] ?? \PhpursThunks::eval('Data_Function_Uncurried_runFn2')))(($GLOBALS['Data_Array_fromFoldableImpl'] ?? \PhpursThunks::eval('Data_Array_fromFoldableImpl'))))(($dictFoldable_0)->foldr);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_foldr'] = function() { $v = ($GLOBALS['Data_Foldable_foldrArray'] ?? \PhpursThunks::eval('Data_Foldable_foldrArray')); return $v; };
\PhpursThunks::$thunks['Data_Array_foldl'] = function() { $v = ($GLOBALS['Data_Foldable_foldlArray'] ?? \PhpursThunks::eval('Data_Foldable_foldlArray')); return $v; };
\PhpursThunks::$thunks['Data_Array_transpose'] = function() { $v = function($xs_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_transpose"), recVars=[];
  $go_1_0 = null;
  $go_1_0 = (function() use (&$go_1_0, $xs_0) {
  $__fn = function($idx_2, $allArrays_3 = null) use (&$go_1_0, $xs_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_1_0"), recVars=["go_1_0"];
  while (true) {
$v_4_1 = (((($GLOBALS['Data_Foldable_foldlArray'] ?? \PhpursThunks::eval('Data_Foldable_foldlArray')))((function() use ($idx_2) {
  $__fn = function($acc_4, $nextArr_5 = null) use ($idx_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["go_1_0"];
  $__local_var_6_1 = ((($GLOBALS['Data_Array_index'] ?? \PhpursThunks::eval('Data_Array_index')))($nextArr_5))($idx_2);
  if ((is_object($__local_var_6_1) && (($__local_var_6_1)->tag === "Nothing"))) {
$__t2 = $acc_4;
} else {
if ((is_object($__local_var_6_1) && (($__local_var_6_1)->tag === "Just"))) {
if ((is_object($acc_4) && (($acc_4)->tag === "Nothing"))) {
$__t3 = [($__local_var_6_1)->value0];
} else {
if ((is_object($acc_4) && (($acc_4)->tag === "Just"))) {
$__t3 = (($GLOBALS['Control_Monad_ST_Internal_run'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_run')))(((($GLOBALS['Data_Array_ST_withArray'] ?? \PhpursThunks::eval('Data_Array_ST_withArray')))((($GLOBALS['Data_Array_ST_push'] ?? \PhpursThunks::eval('Data_Array_ST_push')))(($__local_var_6_1)->value0)))(($acc_4)->value0));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
};
};
$__t2 = new Phpurs_Data1("Just", $__t3);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
};
};
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(new Phpurs_Data0("Nothing")))($xs_0);
if ((is_object($v_4_1) && (($v_4_1)->tag === "Nothing"))) {
$__t5 = $allArrays_3;
} else {
if ((is_object($v_4_1) && (($v_4_1)->tag === "Just"))) {
$__tco_6 = ((($GLOBALS['Data_Semiring_intAdd'] ?? \PhpursThunks::eval('Data_Semiring_intAdd')))($idx_2))(1);
$__tco_7 = (($GLOBALS['Control_Monad_ST_Internal_run'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_run')))(((($GLOBALS['Data_Array_ST_withArray'] ?? \PhpursThunks::eval('Data_Array_ST_withArray')))((($GLOBALS['Data_Array_ST_push'] ?? \PhpursThunks::eval('Data_Array_ST_push')))(($v_4_1)->value0)))($allArrays_3));
$idx_2 = $__tco_6;
$allArrays_3 = $__tco_7;
continue ;
$__t5 = null;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t5 = null;
};
};
$__res = $__t5;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = (($go_1_0)(0))([]);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_foldRecM'] = function() { $v = function($dictMonadRec_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_foldRecM"), recVars=[];
  $Monad0_1_0 = (($dictMonadRec_0)->Monad0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $__local_var_2_1 = (($Monad0_1_0)->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $__res = (function() use ($Monad0_1_0, $__local_var_2_1, $dictMonadRec_0) {
  $__fn = function($f_3, $b_4 = null, $array_5 = null) use ($Monad0_1_0, $__local_var_2_1, $dictMonadRec_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($dictMonadRec_0)->tailRecM)(function($o_6) use ($Monad0_1_0, $__local_var_2_1, $array_5, $f_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if (( ! (is_object((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)(($o_6)->b))((($GLOBALS['Data_Array_length'] ?? \PhpursThunks::eval('Data_Array_length')))($array_5))) && (((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)(($o_6)->b))((($GLOBALS['Data_Array_length'] ?? \PhpursThunks::eval('Data_Array_length')))($array_5)))->tag === "LT")))) {
$__t2 = (($__local_var_2_1)->pure)(new Phpurs_Data1("Done", ($o_6)->a));
} else {
$__t2 = ((((($Monad0_1_0)->Bind1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->bind)((($f_3)(($o_6)->a))((($GLOBALS['Partial_Unsafe__unsafePartial'] ?? \PhpursThunks::eval('Partial_Unsafe__unsafePartial')))(function($dollar__unused_7) use ($array_5, $o_6) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Data_Array_unsafeIndex1'] ?? \PhpursThunks::eval('Data_Array_unsafeIndex1')))($array_5))(($o_6)->b);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))))(function($res__prime___7) use ($__local_var_2_1, $o_6) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($__local_var_2_1)->pure)(new Phpurs_Data1("Loop", (object)["a" => $res__prime___7, "b" => ((($GLOBALS['Data_Semiring_intAdd'] ?? \PhpursThunks::eval('Data_Semiring_intAdd')))(($o_6)->b))(1)]));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
};
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((object)["a" => $b_4, "b" => 0]);
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
\PhpursThunks::$thunks['Data_Array_foldMap'] = function() { $v = function($dictMonoid_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_foldMap"), recVars=[];
  $__res = ((($GLOBALS['Data_Foldable_foldMapDefaultR'] ?? \PhpursThunks::eval('Data_Foldable_foldMapDefaultR')))(($GLOBALS['Data_Foldable_foldableArray'] ?? \PhpursThunks::eval('Data_Foldable_foldableArray'))))($dictMonoid_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_foldM'] = function() { $v = (function() {
  $__fn = function($dictMonad_0, $f_1 = null, $b_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_foldM"), recVars=["Data_Array_foldM"];
  while (true) {
$__res = (((($GLOBALS['Data_Function_Uncurried_runFn3'] ?? \PhpursThunks::eval('Data_Function_Uncurried_runFn3')))(($GLOBALS['Data_Array_unconsImpl'] ?? \PhpursThunks::eval('Data_Array_unconsImpl'))))(function($v_3) use ($b_2, $dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Array_foldM"];
  $__res = (((($dictMonad_0)->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->pure)($b_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((function() use ($b_2, $dictMonad_0, $f_1) {
  $__fn = function($a_3, $as_4 = null) use ($b_2, $dictMonad_0, $f_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Array_foldM"];
  $__res = ((((($dictMonad_0)->Bind1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->bind)((($f_1)($b_2))($a_3)))(function($b__prime___5) use ($as_4, $dictMonad_0, $f_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Array_foldM"];
  $__res = ((((($GLOBALS['Data_Array_foldM'] ?? \PhpursThunks::eval('Data_Array_foldM')))($dictMonad_0))($f_1))($b__prime___5))($as_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})());
goto __end;;
};
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_fold'] = function() { $v = function($dictMonoid_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_fold"), recVars=[];
  $__res = (((($GLOBALS['Data_Foldable_foldMapDefaultR'] ?? \PhpursThunks::eval('Data_Foldable_foldMapDefaultR')))(($GLOBALS['Data_Foldable_foldableArray'] ?? \PhpursThunks::eval('Data_Foldable_foldableArray'))))($dictMonoid_0))((($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))->identity);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_findMap'] = function() { $v = (((($GLOBALS['Data_Function_Uncurried_runFn4'] ?? \PhpursThunks::eval('Data_Function_Uncurried_runFn4')))(($GLOBALS['Data_Array_findMapImpl'] ?? \PhpursThunks::eval('Data_Array_findMapImpl'))))(new Phpurs_Data0("Nothing")))(($GLOBALS['Data_Maybe_isJust'] ?? \PhpursThunks::eval('Data_Maybe_isJust'))); return $v; };
\PhpursThunks::$thunks['Data_Array_findLastIndex'] = function() { $v = (((($GLOBALS['Data_Function_Uncurried_runFn4'] ?? \PhpursThunks::eval('Data_Function_Uncurried_runFn4')))(($GLOBALS['Data_Array_findLastIndexImpl'] ?? \PhpursThunks::eval('Data_Array_findLastIndexImpl'))))(($GLOBALS['Data_Maybe_Just'] ?? \PhpursThunks::eval('Data_Maybe_Just'))))(new Phpurs_Data0("Nothing")); return $v; };
\PhpursThunks::$thunks['Data_Array_insertBy'] = function() { $v = (function() {
  $__fn = function($cmp_0, $x_1 = null, $ys_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_insertBy"), recVars=[];
  $__local_var_3_0 = ((($GLOBALS['Data_Array_findLastIndex'] ?? \PhpursThunks::eval('Data_Array_findLastIndex')))(function($y_3) use ($cmp_0, $x_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (is_object((($cmp_0)($x_1))($y_3)) && (((($cmp_0)($x_1))($y_3))->tag === "GT"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($ys_2);
  if ((is_object($__local_var_3_0) && (($__local_var_3_0)->tag === "Nothing"))) {
$__t1 = 0;
} else {
if ((is_object($__local_var_3_0) && (($__local_var_3_0)->tag === "Just"))) {
$__t1 = ((($GLOBALS['Data_Semiring_intAdd'] ?? \PhpursThunks::eval('Data_Semiring_intAdd')))(($__local_var_3_0)->value0))(1);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
  $i_4_1 = $__t1;
  $__res = (($GLOBALS['Partial_Unsafe__unsafePartial'] ?? \PhpursThunks::eval('Partial_Unsafe__unsafePartial')))(function($dollar__unused_5) use ($i_4_1, $x_1, $ys_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__local_var_6_3 = (((($GLOBALS['Data_Array_insertAt'] ?? \PhpursThunks::eval('Data_Array_insertAt')))($i_4_1))($x_1))($ys_2);
  if ((is_object($__local_var_6_3) && (($__local_var_6_3)->tag === "Just"))) {
$__t4 = ($__local_var_6_3)->value0;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t4 = null;
};
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_insert'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_insert"), recVars=[];
  $__res = (($GLOBALS['Data_Array_insertBy'] ?? \PhpursThunks::eval('Data_Array_insertBy')))(($dictOrd_0)->compare);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_findIndex'] = function() { $v = (((($GLOBALS['Data_Function_Uncurried_runFn4'] ?? \PhpursThunks::eval('Data_Function_Uncurried_runFn4')))(($GLOBALS['Data_Array_findIndexImpl'] ?? \PhpursThunks::eval('Data_Array_findIndexImpl'))))(($GLOBALS['Data_Maybe_Just'] ?? \PhpursThunks::eval('Data_Maybe_Just'))))(new Phpurs_Data0("Nothing")); return $v; };
\PhpursThunks::$thunks['Data_Array_find'] = function() { $v = (function() {
  $__fn = function($f_0, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_find"), recVars=[];
  $__local_var_2_0 = (($GLOBALS['Partial_Unsafe__unsafePartial'] ?? \PhpursThunks::eval('Partial_Unsafe__unsafePartial')))(function($dollar__unused_2) use ($xs_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Data_Array_unsafeIndex1'] ?? \PhpursThunks::eval('Data_Array_unsafeIndex1')))($xs_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__local_var_3_1 = ((($GLOBALS['Data_Array_findIndex'] ?? \PhpursThunks::eval('Data_Array_findIndex')))($f_0))($xs_1);
  if ((is_object($__local_var_3_1) && (($__local_var_3_1)->tag === "Just"))) {
$__t2 = new Phpurs_Data1("Just", ($__local_var_2_0)(($__local_var_3_1)->value0));
} else {
$__t2 = new Phpurs_Data0("Nothing");
};
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_filter'] = function() { $v = (($GLOBALS['Data_Function_Uncurried_runFn2'] ?? \PhpursThunks::eval('Data_Function_Uncurried_runFn2')))(($GLOBALS['Data_Array_filterImpl'] ?? \PhpursThunks::eval('Data_Array_filterImpl'))); return $v; };
\PhpursThunks::$thunks['Data_Array_intersectBy'] = function() { $v = (function() {
  $__fn = function($eq2_0, $xs_1 = null, $ys_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_intersectBy"), recVars=[];
  $__res = ((($GLOBALS['Data_Array_filter'] ?? \PhpursThunks::eval('Data_Array_filter')))(function($x_3) use ($eq2_0, $ys_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__local_var_4_0 = ((($GLOBALS['Data_Array_findIndex'] ?? \PhpursThunks::eval('Data_Array_findIndex')))(($eq2_0)($x_3)))($ys_2);
  if ((is_object($__local_var_4_0) && (($__local_var_4_0)->tag === "Nothing"))) {
$__t1 = false;
} else {
if ((is_object($__local_var_4_0) && (($__local_var_4_0)->tag === "Just"))) {
$__t1 = true;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($xs_1);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_intersect'] = function() { $v = function($dictEq_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_intersect"), recVars=[];
  $__res = (($GLOBALS['Data_Array_intersectBy'] ?? \PhpursThunks::eval('Data_Array_intersectBy')))(($dictEq_0)->eq);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_elemLastIndex'] = function() { $v = (function() {
  $__fn = function($dictEq_0, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_elemLastIndex"), recVars=[];
  $__res = (($GLOBALS['Data_Array_findLastIndex'] ?? \PhpursThunks::eval('Data_Array_findLastIndex')))(function($v_2) use ($dictEq_0, $x_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($dictEq_0)->eq)($v_2))($x_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_elemIndex'] = function() { $v = (function() {
  $__fn = function($dictEq_0, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_elemIndex"), recVars=[];
  $__res = (($GLOBALS['Data_Array_findIndex'] ?? \PhpursThunks::eval('Data_Array_findIndex')))(function($v_2) use ($dictEq_0, $x_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($dictEq_0)->eq)($v_2))($x_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_notElem'] = function() { $v = (function() {
  $__fn = function($dictEq_0, $a_1 = null, $arr_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_notElem"), recVars=[];
  $__local_var_3_0 = ((($GLOBALS['Data_Array_findIndex'] ?? \PhpursThunks::eval('Data_Array_findIndex')))(function($v_3) use ($a_1, $dictEq_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($dictEq_0)->eq)($v_3))($a_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($arr_2);
  if ((is_object($__local_var_3_0) && (($__local_var_3_0)->tag === "Nothing"))) {
$__t1 = true;
} else {
if ((is_object($__local_var_3_0) && (($__local_var_3_0)->tag === "Just"))) {
$__t1 = false;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_elem'] = function() { $v = (function() {
  $__fn = function($dictEq_0, $a_1 = null, $arr_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_elem"), recVars=[];
  $__local_var_3_0 = ((($GLOBALS['Data_Array_findIndex'] ?? \PhpursThunks::eval('Data_Array_findIndex')))(function($v_3) use ($a_1, $dictEq_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($dictEq_0)->eq)($v_3))($a_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($arr_2);
  if ((is_object($__local_var_3_0) && (($__local_var_3_0)->tag === "Nothing"))) {
$__t1 = false;
} else {
if ((is_object($__local_var_3_0) && (($__local_var_3_0)->tag === "Just"))) {
$__t1 = true;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_dropWhile'] = function() { $v = (function() {
  $__fn = function($p_0, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_dropWhile"), recVars=[];
  $__res = (((($GLOBALS['Data_Array_span'] ?? \PhpursThunks::eval('Data_Array_span')))($p_0))($xs_1))->rest;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_dropEnd'] = function() { $v = (function() {
  $__fn = function($n_0, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_dropEnd"), recVars=[];
  $__local_var_2_0 = ((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))((($GLOBALS['Data_Array_length'] ?? \PhpursThunks::eval('Data_Array_length')))($xs_1)))($n_0);
  if ((is_object((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($__local_var_2_0))(1)) && (((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($__local_var_2_0))(1))->tag === "LT"))) {
$__t1 = [];
} else {
$__t1 = (((($GLOBALS['Data_Array_slice'] ?? \PhpursThunks::eval('Data_Array_slice')))(0))($__local_var_2_0))($xs_1);
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_drop'] = function() { $v = (function() {
  $__fn = function($n_0, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_drop"), recVars=[];
  if ((is_object((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($n_0))(1)) && (((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($n_0))(1))->tag === "LT"))) {
$__t0 = $xs_1;
} else {
$__t0 = (((($GLOBALS['Data_Array_slice'] ?? \PhpursThunks::eval('Data_Array_slice')))($n_0))((($GLOBALS['Data_Array_length'] ?? \PhpursThunks::eval('Data_Array_length')))($xs_1)))($xs_1);
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_takeEnd'] = function() { $v = (function() {
  $__fn = function($n_0, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_takeEnd"), recVars=[];
  $__res = ((($GLOBALS['Data_Array_drop'] ?? \PhpursThunks::eval('Data_Array_drop')))(((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))((($GLOBALS['Data_Array_length'] ?? \PhpursThunks::eval('Data_Array_length')))($xs_1)))($n_0)))($xs_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_deleteAt'] = function() { $v = (((($GLOBALS['Data_Function_Uncurried_runFn4'] ?? \PhpursThunks::eval('Data_Function_Uncurried_runFn4')))(($GLOBALS['Data_Array__deleteAt'] ?? \PhpursThunks::eval('Data_Array__deleteAt'))))(($GLOBALS['Data_Maybe_Just'] ?? \PhpursThunks::eval('Data_Maybe_Just'))))(new Phpurs_Data0("Nothing")); return $v; };
\PhpursThunks::$thunks['Data_Array_deleteBy'] = function() { $v = (function() {
  $__fn = function($v_0, $v1_1 = null, $v2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_deleteBy"), recVars=[];
  switch (count($v2_2)) {
case 0:
$__t5 = [];
break;
default:
$__local_var_3_0 = ((($GLOBALS['Data_Array_findIndex'] ?? \PhpursThunks::eval('Data_Array_findIndex')))(($v_0)($v1_1)))($v2_2);
if ((is_object($__local_var_3_0) && (($__local_var_3_0)->tag === "Nothing"))) {
$__t1 = $v2_2;
} else {
if ((is_object($__local_var_3_0) && (($__local_var_3_0)->tag === "Just"))) {
$__local_var_4_2 = ($__local_var_3_0)->value0;
$__t1 = (($GLOBALS['Partial_Unsafe__unsafePartial'] ?? \PhpursThunks::eval('Partial_Unsafe__unsafePartial')))(function($dollar__unused_5) use ($__local_var_4_2, $v2_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__local_var_6_3 = ((($GLOBALS['Data_Array_deleteAt'] ?? \PhpursThunks::eval('Data_Array_deleteAt')))($__local_var_4_2))($v2_2);
  if ((is_object($__local_var_6_3) && (($__local_var_6_3)->tag === "Just"))) {
$__t4 = ($__local_var_6_3)->value0;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t4 = null;
};
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
$__t5 = $__t1;
break;
};
  $__res = $__t5;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_delete'] = function() { $v = function($dictEq_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_delete"), recVars=[];
  $__res = (($GLOBALS['Data_Array_deleteBy'] ?? \PhpursThunks::eval('Data_Array_deleteBy')))(($dictEq_0)->eq);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_difference'] = function() { $v = function($dictEq_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_difference"), recVars=[];
  $__res = (($GLOBALS['Data_Foldable_foldrArray'] ?? \PhpursThunks::eval('Data_Foldable_foldrArray')))((($GLOBALS['Data_Array_delete'] ?? \PhpursThunks::eval('Data_Array_delete')))($dictEq_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_cons'] = function() { $v = (function() {
  $__fn = function($x_0, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_cons"), recVars=[];
  $__res = ((($GLOBALS['Data_Semigroup_concatArray'] ?? \PhpursThunks::eval('Data_Semigroup_concatArray')))([$x_0]))($xs_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_some'] = function() { $v = (function() {
  $__fn = function($dictAlternative_0, $dictLazy_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_some"), recVars=["Data_Array_some","Data_Array_many"];
  while (true) {
$__res = ((((((($dictAlternative_0)->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Apply0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->apply)(((((((((($dictAlternative_0)->Plus1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Alt0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Functor0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->map)(($GLOBALS['Data_Array_cons'] ?? \PhpursThunks::eval('Data_Array_cons'))))($v_2)))((($dictLazy_1)->defer)(function($v1_3) use ($dictAlternative_0, $dictLazy_1, $v_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Array_some","Data_Array_many"];
  $__res = (((($GLOBALS['Data_Array_many'] ?? \PhpursThunks::eval('Data_Array_many')))($dictAlternative_0))($dictLazy_1))($v_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
goto __end;;
};
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_many'] = function() { $v = (function() {
  $__fn = function($dictAlternative_0, $dictLazy_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_many"), recVars=["Data_Array_some","Data_Array_many"];
  while (true) {
$__res = ((((((($dictAlternative_0)->Plus1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Alt0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->alt)((((($GLOBALS['Data_Array_some'] ?? \PhpursThunks::eval('Data_Array_some')))($dictAlternative_0))($dictLazy_1))($v_2)))((((($dictAlternative_0)->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->pure)([]));
goto __end;;
};
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_concatMap'] = function() { $v = (function() {
  $__fn = function($b_0, $a_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_concatMap"), recVars=[];
  $__res = ((($GLOBALS['Control_Bind_arrayBind'] ?? \PhpursThunks::eval('Control_Bind_arrayBind')))($a_1))($b_0);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_mapMaybe'] = function() { $v = function($f_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_mapMaybe"), recVars=[];
  $__res = (($GLOBALS['Data_Array_concatMap'] ?? \PhpursThunks::eval('Data_Array_concatMap')))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(function($v2_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v2_1) && (($v2_1)->tag === "Nothing"))) {
$__t0 = [];
} else {
if ((is_object($v2_1) && (($v2_1)->tag === "Just"))) {
$__t0 = [($v2_1)->value0];
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($f_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_filterA'] = function() { $v = function($dictApplicative_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_filterA"), recVars=[];
  $traverse1_1_0 = ((($GLOBALS['Data_Traversable_traversableArray'] ?? \PhpursThunks::eval('Data_Traversable_traversableArray')))->traverse)($dictApplicative_0);
  $__local_var_2_1 = (((($dictApplicative_0)->Apply0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Functor0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $__res = function($p_3) use ($__local_var_2_1, $traverse1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($__local_var_2_1)->map)((($GLOBALS['Data_Array_mapMaybe'] ?? \PhpursThunks::eval('Data_Array_mapMaybe')))(function($v_4) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if (($v_4)->value1) {
$__t2 = new Phpurs_Data1("Just", ($v_4)->value0);
} else {
$__t2 = new Phpurs_Data0("Nothing");
};
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))))(($traverse1_1_0)(function($x_4) use ($__local_var_2_1, $p_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($__local_var_2_1)->map)((($GLOBALS['Data_Tuple_Tuple'] ?? \PhpursThunks::eval('Data_Tuple_Tuple')))($x_4)))(($p_3)($x_4));
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_catMaybes'] = function() { $v = (($GLOBALS['Data_Array_mapMaybe'] ?? \PhpursThunks::eval('Data_Array_mapMaybe')))((($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))->identity); return $v; };
\PhpursThunks::$thunks['Data_Array_any'] = function() { $v = (($GLOBALS['Data_Function_Uncurried_runFn2'] ?? \PhpursThunks::eval('Data_Function_Uncurried_runFn2')))(($GLOBALS['Data_Array_anyImpl'] ?? \PhpursThunks::eval('Data_Array_anyImpl'))); return $v; };
\PhpursThunks::$thunks['Data_Array_nubByEq'] = function() { $v = (function() {
  $__fn = function($eq2_0, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_nubByEq"), recVars=[];
  $__res = (($GLOBALS['Control_Monad_ST_Internal_run'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_run')))(((($GLOBALS['Control_Monad_ST_Internal_bind_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_bind_')))(($GLOBALS['Data_Array_ST_new'] ?? \PhpursThunks::eval('Data_Array_ST_new'))))(function($arr_2) use ($eq2_0, $xs_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bind_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_bind_')))(((($GLOBALS['Control_Monad_ST_Internal_foreach'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_foreach')))($xs_1))(function($x_3) use ($arr_2, $eq2_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bind_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_bind_')))(((($GLOBALS['Control_Monad_ST_Internal_map_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_map_')))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_HeytingAlgebra_boolNot'] ?? \PhpursThunks::eval('Data_HeytingAlgebra_boolNot'))))((($GLOBALS['Data_Array_any'] ?? \PhpursThunks::eval('Data_Array_any')))(function($v_4) use ($eq2_0, $x_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($eq2_0)($v_4))($x_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))))((($GLOBALS['Data_Array_ST_unsafeFreeze'] ?? \PhpursThunks::eval('Data_Array_ST_unsafeFreeze')))($arr_2))))(function($e_4) use ($arr_2, $x_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__local_var_5_0 = (($GLOBALS['Data_Array_void'] ?? \PhpursThunks::eval('Data_Array_void')))(((($GLOBALS['Data_Array_ST_push'] ?? \PhpursThunks::eval('Data_Array_ST_push')))($x_3))($arr_2));
  if ($e_4) {
$__t1 = $__local_var_5_0;
} else {
$__t1 = (($GLOBALS['Control_Monad_ST_Internal_pure_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_pure_')))(($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit')));
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(function($dollar__unused_3) use ($arr_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Data_Array_ST_unsafeFreeze'] ?? \PhpursThunks::eval('Data_Array_ST_unsafeFreeze')))($arr_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_nubEq'] = function() { $v = function($dictEq_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_nubEq"), recVars=[];
  $__res = (($GLOBALS['Data_Array_nubByEq'] ?? \PhpursThunks::eval('Data_Array_nubByEq')))(($dictEq_0)->eq);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_unionBy'] = function() { $v = (function() {
  $__fn = function($eq2_0, $xs_1 = null, $ys_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_unionBy"), recVars=[];
  $__res = ((($GLOBALS['Data_Semigroup_concatArray'] ?? \PhpursThunks::eval('Data_Semigroup_concatArray')))($xs_1))((((($GLOBALS['Data_Foldable_foldlArray'] ?? \PhpursThunks::eval('Data_Foldable_foldlArray')))((function() use ($eq2_0) {
  $__fn = function($b_3, $a_4 = null) use ($eq2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (((($GLOBALS['Data_Array_deleteBy'] ?? \PhpursThunks::eval('Data_Array_deleteBy')))($eq2_0))($a_4))($b_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(((($GLOBALS['Data_Array_nubByEq'] ?? \PhpursThunks::eval('Data_Array_nubByEq')))($eq2_0))($ys_2)))($xs_1));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_union'] = function() { $v = function($dictEq_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_union"), recVars=[];
  $__res = (($GLOBALS['Data_Array_unionBy'] ?? \PhpursThunks::eval('Data_Array_unionBy')))(($dictEq_0)->eq);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_alterAt'] = function() { $v = (function() {
  $__fn = function($i_0, $f_1 = null, $xs_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_alterAt"), recVars=[];
  $__local_var_3_0 = ((($GLOBALS['Data_Array_index'] ?? \PhpursThunks::eval('Data_Array_index')))($xs_2))($i_0);
  if ((is_object($__local_var_3_0) && (($__local_var_3_0)->tag === "Nothing"))) {
$__t1 = new Phpurs_Data0("Nothing");
} else {
if ((is_object($__local_var_3_0) && (($__local_var_3_0)->tag === "Just"))) {
$v_4_2 = ($f_1)(($__local_var_3_0)->value0);
if ((is_object($v_4_2) && (($v_4_2)->tag === "Nothing"))) {
$__t3 = ((($GLOBALS['Data_Array_deleteAt'] ?? \PhpursThunks::eval('Data_Array_deleteAt')))($i_0))($xs_2);
} else {
if ((is_object($v_4_2) && (($v_4_2)->tag === "Just"))) {
$__t3 = (((($GLOBALS['Data_Array_updateAt'] ?? \PhpursThunks::eval('Data_Array_updateAt')))($i_0))(($v_4_2)->value0))($xs_2);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
};
};
$__t1 = $__t3;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_all'] = function() { $v = (($GLOBALS['Data_Function_Uncurried_runFn2'] ?? \PhpursThunks::eval('Data_Function_Uncurried_runFn2')))(($GLOBALS['Data_Array_allImpl'] ?? \PhpursThunks::eval('Data_Array_allImpl'))); return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };
$ffi_Data_Array = \call_user_func(function() {
  $exports = [];
$rangeImpl = function($start, $end = null) use (&$rangeImpl) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$rangeImpl) {

            return $rangeImpl(...\array_merge($__args, $more));
        };
    }
    $step = $start > $end ? -1 : 1;
    $result = [];
    $i = $start;
    while ($i !== $end) {
        $result[] = $i;
        $i += $step;
    }
    $result[] = $i;
    return $result;
};

$replicateImpl = function($count, $value = null) use (&$replicateImpl) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$replicateImpl) {

            return $replicateImpl(...\array_merge($__args, $more));
        };
    }
    if ($count < 1) return [];
    return array_fill(0, $count, $value);
};

$fromFoldableImpl = function($foldr, $xs = null) use (&$fromFoldableImpl) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$fromFoldableImpl) {

            return $fromFoldableImpl(...\array_merge($__args, $more));
        };
    }
    
    $emptyList = new \stdClass();
    $curryCons = function($head, $tail = null) {
        if (\func_num_args() < 2) {
            return function($tail) use ($head) {
                $obj = new \stdClass();
                $obj->head = $head;
                $obj->tail = $tail;
                return $obj;
            };
        }
        $obj = new \stdClass();
        $obj->head = $head;
        $obj->tail = $tail;
        return $obj;
    };
    
    $listToArray = function($list) use ($emptyList) {
        $result = [];
        $xs = $list;
        while ($xs !== $emptyList && $xs !== null) {
            $result[] = $xs->head;
            $xs = $xs->tail;
        }
        return $result;
    };
    
    return $listToArray($foldr($curryCons)($emptyList)($xs));
};

$length = function($xs) use (&$length) {
    return \count($xs);
};

$unconsImpl = function($empty, $next = null, $xs = null) use (&$unconsImpl) {
    if (\func_num_args() < 3) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$unconsImpl) {

            return $unconsImpl(...\array_merge($__args, $more));
        };
    }
    if (\count($xs) === 0) return $empty((object)[]);
    return $next($xs[0])(\array_slice($xs, 1));
};

$indexImpl = function($just, $nothing = null, $xs = null, $i = null) use (&$indexImpl) {
    if (\func_num_args() < 4) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$indexImpl) {

            return $indexImpl(...\array_merge($__args, $more));
        };
    }
    return ($i < 0 || $i >= \count($xs)) ? $nothing : $just($xs[$i]);
};

$findMapImpl = function($nothing, $isJust = null, $f = null, $xs = null) use (&$findMapImpl) {
    if (\func_num_args() < 4) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$findMapImpl) {

            return $findMapImpl(...\array_merge($__args, $more));
        };
    }
    foreach ($xs as $x) {
        $result = $f($x);
        if ($isJust($result)) return $result;
    }
    return $nothing;
};

$findIndexImpl = function($just, $nothing = null, $f = null, $xs = null) use (&$findIndexImpl) {
    if (\func_num_args() < 4) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$findIndexImpl) {

            return $findIndexImpl(...\array_merge($__args, $more));
        };
    }
    foreach ($xs as $i => $x) {
        if ($f($x)) return $just($i);
    }
    return $nothing;
};

$findLastIndexImpl = function($just, $nothing = null, $f = null, $xs = null) use (&$findLastIndexImpl) {
    if (\func_num_args() < 4) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$findLastIndexImpl) {

            return $findLastIndexImpl(...\array_merge($__args, $more));
        };
    }
    for ($i = \count($xs) - 1; $i >= 0; $i--) {
        if ($f($xs[$i])) return $just($i);
    }
    return $nothing;
};

$_insertAt = function($just, $nothing = null, $i = null, $a = null, $l = null) use (&$_insertAt) {
    if (\func_num_args() < 5) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$_insertAt) {

            return $_insertAt(...\array_merge($__args, $more));
        };
    }
    if ($i < 0 || $i > \count($l)) return $nothing;
    $l1 = $l;
    array_splice($l1, $i, 0, [$a]);
    return $just($l1);
};

$_deleteAt = function($just, $nothing = null, $i = null, $l = null) use (&$_deleteAt) {
    if (\func_num_args() < 4) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$_deleteAt) {

            return $_deleteAt(...\array_merge($__args, $more));
        };
    }
    if ($i < 0 || $i >= \count($l)) return $nothing;
    $l1 = $l;
    array_splice($l1, $i, 1);
    return $just($l1);
};

$_updateAt = function($just, $nothing = null, $i = null, $a = null, $l = null) use (&$_updateAt) {
    if (\func_num_args() < 5) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$_updateAt) {

            return $_updateAt(...\array_merge($__args, $more));
        };
    }
    if ($i < 0 || $i >= \count($l)) return $nothing;
    $l1 = $l;
    $l1[$i] = $a;
    return $just($l1);
};

$reverse = function($l) use (&$reverse) {
    return array_reverse($l);
};

$concat = function($xss) use (&$concat) {
    if (\count($xss) === 0) return [];
    return \array_merge(...$xss);
};

$filterImpl = function($f, $xs = null) use (&$filterImpl) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$filterImpl) {

            return $filterImpl(...\array_merge($__args, $more));
        };
    }
    $res = [];
    foreach ($xs as $x) {
        if ($f($x)) $res[] = $x;
    }
    return $res;
};

$partitionImpl = function($f, $xs = null) use (&$partitionImpl) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$partitionImpl) {

            return $partitionImpl(...\array_merge($__args, $more));
        };
    }
    $yes = [];
    $no = [];
    foreach ($xs as $x) {
        if ($f($x)) $yes[] = $x;
        else $no[] = $x;
    }
    return (object)["yes" => $yes, "no" => $no];
};

$scanlImpl = function($f, $b = null, $xs = null) use (&$scanlImpl) {
    if (\func_num_args() < 3) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$scanlImpl) {

            return $scanlImpl(...\array_merge($__args, $more));
        };
    }
    $acc = $b;
    $out = [];
    foreach ($xs as $x) {
        $acc = $f($acc)($x);
        $out[] = $acc;
    }
    return $out;
};

$scanrImpl = function($f, $b = null, $xs = null) use (&$scanrImpl) {
    if (\func_num_args() < 3) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$scanrImpl) {

            return $scanrImpl(...\array_merge($__args, $more));
        };
    }
    $len = \count($xs);
    $acc = $b;
    $out = array_fill(0, $len, null);
    for ($i = $len - 1; $i >= 0; $i--) {
        $acc = $f($xs[$i])($acc);
        $out[$i] = $acc;
    }
    return $out;
};

$sortByImpl = function($compare, $fromOrdering = null, $xs = null) use (&$sortByImpl) {
    if (\func_num_args() < 3) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$sortByImpl) {

            return $sortByImpl(...\array_merge($__args, $more));
        };
    }
    $out = $xs;
    \usort($out, function($a, $b) use ($compare, $fromOrdering) {
        return $fromOrdering($compare($a)($b));
    });
    return $out;
};

$sliceImpl = function($s, $e = null, $l = null) use (&$sliceImpl) {
    if (\func_num_args() < 3) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$sliceImpl) {

            return $sliceImpl(...\array_merge($__args, $more));
        };
    }
    return \array_slice($l, $s, $e - $s);
};

$zipWithImpl = function($f, $xs = null, $ys = null) use (&$zipWithImpl) {
    if (\func_num_args() < 3) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$zipWithImpl) {

            return $zipWithImpl(...\array_merge($__args, $more));
        };
    }
    $l = \min(\count($xs), \count($ys));
    $result = [];
    for ($i = 0; $i < $l; $i++) {
        $result[] = $f($xs[$i])($ys[$i]);
    }
    return $result;
};

$anyImpl = function($p, $xs = null) use (&$anyImpl) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$anyImpl) {

            return $anyImpl(...\array_merge($__args, $more));
        };
    }
    foreach ($xs as $x) {
        if ($p($x)) return true;
    }
    return false;
};

$allImpl = function($p, $xs = null) use (&$allImpl) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$allImpl) {

            return $allImpl(...\array_merge($__args, $more));
        };
    }
    foreach ($xs as $x) {
        if (!$p($x)) return false;
    }
    return true;
};

$unsafeIndexImpl = function($xs, $n = null) use (&$unsafeIndexImpl) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$unsafeIndexImpl) {

            return $unsafeIndexImpl(...\array_merge($__args, $more));
        };
    }
    return $xs[$n];
};

$exports['rangeImpl'] = $rangeImpl;
$exports['replicateImpl'] = $replicateImpl;
$exports['fromFoldableImpl'] = $fromFoldableImpl;
$exports['length'] = $length;
$exports['unconsImpl'] = $unconsImpl;
$exports['indexImpl'] = $indexImpl;
$exports['findMapImpl'] = $findMapImpl;
$exports['findIndexImpl'] = $findIndexImpl;
$exports['findLastIndexImpl'] = $findLastIndexImpl;
$exports['_insertAt'] = $_insertAt;
$exports['_deleteAt'] = $_deleteAt;
$exports['_updateAt'] = $_updateAt;
$exports['reverse'] = $reverse;
$exports['concat'] = $concat;
$exports['filterImpl'] = $filterImpl;
$exports['partitionImpl'] = $partitionImpl;
$exports['scanlImpl'] = $scanlImpl;
$exports['scanrImpl'] = $scanrImpl;
$exports['sortByImpl'] = $sortByImpl;
$exports['sliceImpl'] = $sliceImpl;
$exports['zipWithImpl'] = $zipWithImpl;
$exports['anyImpl'] = $anyImpl;
$exports['allImpl'] = $allImpl;
$exports['unsafeIndexImpl'] = $unsafeIndexImpl;
return $exports;
  return $exports;
});
\PhpursThunks::$thunks['Data_Array__deleteAt'] = function() use (&$ffi_Data_Array) { return $ffi_Data_Array['_deleteAt']; };
\PhpursThunks::$thunks['Data_Array__insertAt'] = function() use (&$ffi_Data_Array) { return $ffi_Data_Array['_insertAt']; };
\PhpursThunks::$thunks['Data_Array__updateAt'] = function() use (&$ffi_Data_Array) { return $ffi_Data_Array['_updateAt']; };
\PhpursThunks::$thunks['Data_Array_allImpl'] = function() use (&$ffi_Data_Array) { return $ffi_Data_Array['allImpl']; };
\PhpursThunks::$thunks['Data_Array_anyImpl'] = function() use (&$ffi_Data_Array) { return $ffi_Data_Array['anyImpl']; };
\PhpursThunks::$thunks['Data_Array_concat'] = function() use (&$ffi_Data_Array) { return $ffi_Data_Array['concat']; };
\PhpursThunks::$thunks['Data_Array_filterImpl'] = function() use (&$ffi_Data_Array) { return $ffi_Data_Array['filterImpl']; };
\PhpursThunks::$thunks['Data_Array_findIndexImpl'] = function() use (&$ffi_Data_Array) { return $ffi_Data_Array['findIndexImpl']; };
\PhpursThunks::$thunks['Data_Array_findLastIndexImpl'] = function() use (&$ffi_Data_Array) { return $ffi_Data_Array['findLastIndexImpl']; };
\PhpursThunks::$thunks['Data_Array_findMapImpl'] = function() use (&$ffi_Data_Array) { return $ffi_Data_Array['findMapImpl']; };
\PhpursThunks::$thunks['Data_Array_fromFoldableImpl'] = function() use (&$ffi_Data_Array) { return $ffi_Data_Array['fromFoldableImpl']; };
\PhpursThunks::$thunks['Data_Array_indexImpl'] = function() use (&$ffi_Data_Array) { return $ffi_Data_Array['indexImpl']; };
\PhpursThunks::$thunks['Data_Array_length'] = function() use (&$ffi_Data_Array) { return $ffi_Data_Array['length']; };
\PhpursThunks::$thunks['Data_Array_partitionImpl'] = function() use (&$ffi_Data_Array) { return $ffi_Data_Array['partitionImpl']; };
\PhpursThunks::$thunks['Data_Array_rangeImpl'] = function() use (&$ffi_Data_Array) { return $ffi_Data_Array['rangeImpl']; };
\PhpursThunks::$thunks['Data_Array_replicateImpl'] = function() use (&$ffi_Data_Array) { return $ffi_Data_Array['replicateImpl']; };
\PhpursThunks::$thunks['Data_Array_reverse'] = function() use (&$ffi_Data_Array) { return $ffi_Data_Array['reverse']; };
\PhpursThunks::$thunks['Data_Array_scanlImpl'] = function() use (&$ffi_Data_Array) { return $ffi_Data_Array['scanlImpl']; };
\PhpursThunks::$thunks['Data_Array_scanrImpl'] = function() use (&$ffi_Data_Array) { return $ffi_Data_Array['scanrImpl']; };
\PhpursThunks::$thunks['Data_Array_sliceImpl'] = function() use (&$ffi_Data_Array) { return $ffi_Data_Array['sliceImpl']; };
\PhpursThunks::$thunks['Data_Array_sortByImpl'] = function() use (&$ffi_Data_Array) { return $ffi_Data_Array['sortByImpl']; };
\PhpursThunks::$thunks['Data_Array_unconsImpl'] = function() use (&$ffi_Data_Array) { return $ffi_Data_Array['unconsImpl']; };
\PhpursThunks::$thunks['Data_Array_unsafeIndexImpl'] = function() use (&$ffi_Data_Array) { return $ffi_Data_Array['unsafeIndexImpl']; };
\PhpursThunks::$thunks['Data_Array_zipWithImpl'] = function() use (&$ffi_Data_Array) { return $ffi_Data_Array['zipWithImpl']; };




























































































