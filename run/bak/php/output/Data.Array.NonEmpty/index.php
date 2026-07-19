<?php

namespace Data\Array\NonEmpty;

// ALL IMPORTS: Control.Alternative, Control.Bind, Control.Lazy, Control.Monad.Rec.Class, Control.Semigroupoid, Data.Array, Data.Array.NonEmpty, Data.Array.NonEmpty.Internal, Data.Bifunctor, Data.Boolean, Data.Eq, Data.Foldable, Data.Function, Data.Functor, Data.Maybe, Data.NonEmpty, Data.Ord, Data.Ring, Data.Semigroup, Data.Semigroup.Foldable, Data.Semiring, Data.Tuple, Data.Unfoldable, Data.Unfoldable1, Partial.Unsafe, Prelude, Prim, Safe.Coerce, Unsafe.Coerce
// TO REQUIRE: Control.Alternative, Control.Bind, Control.Lazy, Control.Monad.Rec.Class, Control.Semigroupoid, Data.Array, Data.Array.NonEmpty, Data.Array.NonEmpty.Internal, Data.Bifunctor, Data.Boolean, Data.Eq, Data.Foldable, Data.Function, Data.Functor, Data.Maybe, Data.NonEmpty, Data.Ord, Data.Ring, Data.Semigroup, Data.Semigroup.Foldable, Data.Semiring, Data.Tuple, Data.Unfoldable, Data.Unfoldable1, Partial.Unsafe, Prelude, Safe.Coerce, Unsafe.Coerce
require_once __DIR__ . '/../Control.Alternative/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Lazy/index.php';
require_once __DIR__ . '/../Control.Monad.Rec.Class/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Array/index.php';
require_once __DIR__ . '/../Data.Array.NonEmpty/index.php';
require_once __DIR__ . '/../Data.Array.NonEmpty.Internal/index.php';
require_once __DIR__ . '/../Data.Bifunctor/index.php';
require_once __DIR__ . '/../Data.Boolean/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Foldable/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.NonEmpty/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ring/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Semigroup.Foldable/index.php';
require_once __DIR__ . '/../Data.Semiring/index.php';
require_once __DIR__ . '/../Data.Tuple/index.php';
require_once __DIR__ . '/../Data.Unfoldable/index.php';
require_once __DIR__ . '/../Data.Unfoldable1/index.php';
require_once __DIR__ . '/../Partial.Unsafe/index.php';
require_once __DIR__ . '/../Prelude/index.php';
require_once __DIR__ . '/../Safe.Coerce/index.php';
require_once __DIR__ . '/../Unsafe.Coerce/index.php';

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
\PhpursThunks::$thunks['Data_Array_NonEmpty_max'] = function() { $v = (function() {
  $__fn = function($x_0, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_max"), recVars=[];
  $v_2_0 = (((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($x_0))($y_1);
  if ((is_object($v_2_0) && (($v_2_0)->tag === "LT"))) {
$__t1 = $y_1;
} else {
if ((is_object($v_2_0) && (($v_2_0)->tag === "EQ"))) {
$__t1 = $x_0;
} else {
if ((is_object($v_2_0) && (($v_2_0)->tag === "GT"))) {
$__t1 = $x_0;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_fromJust'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_fromJust"), recVars=[];
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
\PhpursThunks::$thunks['Data_Array_NonEmpty_unsafeIndex1'] = function() { $v = (function() {
  $__fn = function($__local_var_0, $__local_var_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_unsafeIndex1"), recVars=[];
  $__res = ($__local_var_0)[0];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_transpose'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_transpose'] ?? \PhpursThunks::eval('Data_Array_transpose'))))(($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce')))); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_toArray'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_toArray"), recVars=[];
  $__res = $v_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_unionBy__prime__'] = function() { $v = (function() {
  $__fn = function($eq_0, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_unionBy'"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_Internal_NonEmptyArray'))))(((($GLOBALS['Data_Array_unionBy'] ?? \PhpursThunks::eval('Data_Array_unionBy')))($eq_0))($xs_1));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_union__prime__'] = function() { $v = function($dictEq_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_union'"), recVars=[];
  $__res = (($GLOBALS['Data_Array_NonEmpty_unionBy__prime__'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_unionBy__prime__')))(($dictEq_0)->eq);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_unionBy'] = function() { $v = (function() {
  $__fn = function($eq_0, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_unionBy"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_Internal_NonEmptyArray'))))(((($GLOBALS['Data_Array_unionBy'] ?? \PhpursThunks::eval('Data_Array_unionBy')))($eq_0))($xs_1))))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray')));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_union'] = function() { $v = function($dictEq_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_union"), recVars=[];
  $__res = (($GLOBALS['Data_Array_NonEmpty_unionBy'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_unionBy')))(($dictEq_0)->eq);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_unzip'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = new Phpurs_Data2("Tuple", ($v_0)->value0, ($v_0)->value1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_unzip'] ?? \PhpursThunks::eval('Data_Array_unzip'))))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray')))); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_updateAt'] = function() { $v = (function() {
  $__fn = function($i_0, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_updateAt"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(((($GLOBALS['Data_Array_updateAt'] ?? \PhpursThunks::eval('Data_Array_updateAt')))($i_0))($x_1)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray'))));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_zip'] = function() { $v = (function() {
  $__fn = function($xs_0, $ys_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_zip"), recVars=[];
  $__res = (($GLOBALS['Data_Array_zipWithImpl'] ?? \PhpursThunks::eval('Data_Array_zipWithImpl')))(($GLOBALS['Data_Tuple_Tuple'] ?? \PhpursThunks::eval('Data_Tuple_Tuple')), $xs_0, $ys_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_zipWith'] = function() { $v = (function() {
  $__fn = function($f_0, $xs_1 = null, $ys_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_zipWith"), recVars=[];
  $__res = (($GLOBALS['Data_Array_zipWithImpl'] ?? \PhpursThunks::eval('Data_Array_zipWithImpl')))($f_0, $xs_1, $ys_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_zipWithA'] = function() { $v = function($dictApplicative_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_zipWithA"), recVars=[];
  $__res = (($GLOBALS['Data_Array_zipWithA'] ?? \PhpursThunks::eval('Data_Array_zipWithA')))($dictApplicative_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_splitAt'] = function() { $v = (function() {
  $__fn = function($i_0, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_splitAt"), recVars=[];
  $__res = ((($GLOBALS['Data_Array_splitAt'] ?? \PhpursThunks::eval('Data_Array_splitAt')))($i_0))($xs_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_some'] = function() { $v = (function() {
  $__fn = function($dictAlternative_0, $dictLazy_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_some"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce'))))(((($GLOBALS['Data_Array_some'] ?? \PhpursThunks::eval('Data_Array_some')))($dictAlternative_0))($dictLazy_1));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_snoc__prime__'] = function() { $v = (function() {
  $__fn = function($xs_0, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_snoc'"), recVars=[];
  $__res = (($GLOBALS['Control_Monad_ST_Internal_run'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_run')))(((($GLOBALS['Data_Array_ST_withArray'] ?? \PhpursThunks::eval('Data_Array_ST_withArray')))((($GLOBALS['Data_Array_ST_push'] ?? \PhpursThunks::eval('Data_Array_ST_push')))($x_1)))($xs_0));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_snoc'] = function() { $v = (function() {
  $__fn = function($xs_0, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_snoc"), recVars=[];
  $__res = (($GLOBALS['Control_Monad_ST_Internal_run'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_run')))(((($GLOBALS['Data_Array_ST_withArray'] ?? \PhpursThunks::eval('Data_Array_ST_withArray')))((($GLOBALS['Data_Array_ST_push'] ?? \PhpursThunks::eval('Data_Array_ST_push')))($x_1)))($xs_0));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_singleton'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_Internal_NonEmptyArray'))))(($GLOBALS['Data_Array_singleton'] ?? \PhpursThunks::eval('Data_Array_singleton'))); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_replicate'] = function() { $v = (function() {
  $__fn = function($i_0, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_replicate"), recVars=[];
  $__res = (($GLOBALS['Data_Array_replicateImpl'] ?? \PhpursThunks::eval('Data_Array_replicateImpl')))(((($GLOBALS['Data_Array_NonEmpty_max'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_max')))(1))($i_0), $x_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_range'] = function() { $v = (function() {
  $__fn = function($x_0, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_range"), recVars=[];
  $__res = (($GLOBALS['Data_Array_rangeImpl'] ?? \PhpursThunks::eval('Data_Array_rangeImpl')))($x_0, $y_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_prependArray'] = function() { $v = (function() {
  $__fn = function($xs_0, $ys_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_prependArray"), recVars=[];
  $__res = ((($GLOBALS['Data_Semigroup_concatArray'] ?? \PhpursThunks::eval('Data_Semigroup_concatArray')))($xs_0))($ys_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_modifyAt'] = function() { $v = (function() {
  $__fn = function($i_0, $f_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_modifyAt"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(((($GLOBALS['Data_Array_modifyAt'] ?? \PhpursThunks::eval('Data_Array_modifyAt')))($i_0))($f_1)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray'))));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_intersectBy__prime__'] = function() { $v = (function() {
  $__fn = function($eq_0, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_intersectBy'"), recVars=[];
  $__res = ((($GLOBALS['Data_Array_intersectBy'] ?? \PhpursThunks::eval('Data_Array_intersectBy')))($eq_0))($xs_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_intersectBy'] = function() { $v = (function() {
  $__fn = function($eq_0, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_intersectBy"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(((($GLOBALS['Data_Array_NonEmpty_intersectBy__prime__'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_intersectBy__prime__')))($eq_0))($xs_1)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray')));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_intersect__prime__'] = function() { $v = function($dictEq_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_intersect'"), recVars=[];
  $__res = (($GLOBALS['Data_Array_NonEmpty_intersectBy__prime__'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_intersectBy__prime__')))(($dictEq_0)->eq);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_intersect'] = function() { $v = function($dictEq_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_intersect"), recVars=[];
  $__res = (($GLOBALS['Data_Array_NonEmpty_intersectBy'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_intersectBy')))(($dictEq_0)->eq);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_intercalate'] = function() { $v = function($dictSemigroup_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_intercalate"), recVars=[];
  $__local_var_1_0 = ((($GLOBALS['Data_Semigroup_Foldable_intercalateMap'] ?? \PhpursThunks::eval('Data_Semigroup_Foldable_intercalateMap')))(($GLOBALS['Data_Array_NonEmpty_Internal_foldable1NonEmptyArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_Internal_foldable1NonEmptyArray'))))($dictSemigroup_0);
  $__res = function($a_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($__local_var_1_0)($a_2))((($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))->identity);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_insertAt'] = function() { $v = (function() {
  $__fn = function($i_0, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_insertAt"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(((($GLOBALS['Data_Array_insertAt'] ?? \PhpursThunks::eval('Data_Array_insertAt')))($i_0))($x_1)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray'))));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_fromFoldable1'] = function() { $v = function($dictFoldable1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_fromFoldable1"), recVars=[];
  $__local_var_1_0 = ((($dictFoldable1_0)->Foldable0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->foldr;
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_Internal_NonEmptyArray'))))(function($__local_var_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Data_Array_fromFoldableImpl'] ?? \PhpursThunks::eval('Data_Array_fromFoldableImpl')))($__local_var_1_0, $__local_var_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_fromArray'] = function() { $v = function($xs_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_fromArray"), recVars=[];
  if ((count($xs_0) > 0)) {
$__t0 = new Phpurs_Data1("Just", $xs_0);
} else {
$__t0 = new Phpurs_Data0("Nothing");
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_fromFoldable'] = function() { $v = function($dictFoldable_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_fromFoldable"), recVars=[];
  $__local_var_1_0 = ($dictFoldable_0)->foldr;
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_NonEmpty_fromArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_fromArray'))))(function($__local_var_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Data_Array_fromFoldableImpl'] ?? \PhpursThunks::eval('Data_Array_fromFoldableImpl')))($__local_var_1_0, $__local_var_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_transpose__prime__'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_NonEmpty_fromArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_fromArray'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_transpose'] ?? \PhpursThunks::eval('Data_Array_transpose'))))(($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce')))); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_foldr1'] = function() { $v = (($GLOBALS['Data_Array_NonEmpty_Internal_foldable1NonEmptyArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_Internal_foldable1NonEmptyArray')))->foldr1; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_foldl1'] = function() { $v = (($GLOBALS['Data_Array_NonEmpty_Internal_foldable1NonEmptyArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_Internal_foldable1NonEmptyArray')))->foldl1; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_foldMap1'] = function() { $v = function($dictSemigroup_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_foldMap1"), recVars=[];
  $__res = (((($GLOBALS['Data_Semigroup_Foldable_foldMap1DefaultL'] ?? \PhpursThunks::eval('Data_Semigroup_Foldable_foldMap1DefaultL')))(($GLOBALS['Data_Array_NonEmpty_Internal_foldable1NonEmptyArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_Internal_foldable1NonEmptyArray'))))(($GLOBALS['Data_Functor_functorArray'] ?? \PhpursThunks::eval('Data_Functor_functorArray'))))($dictSemigroup_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_fold1'] = function() { $v = function($dictSemigroup_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_fold1"), recVars=[];
  $__res = ((((($GLOBALS['Data_Semigroup_Foldable_foldMap1DefaultL'] ?? \PhpursThunks::eval('Data_Semigroup_Foldable_foldMap1DefaultL')))(($GLOBALS['Data_Array_NonEmpty_Internal_foldable1NonEmptyArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_Internal_foldable1NonEmptyArray'))))(($GLOBALS['Data_Functor_functorArray'] ?? \PhpursThunks::eval('Data_Functor_functorArray'))))($dictSemigroup_0))((($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))->identity);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_difference__prime__'] = function() { $v = function($dictEq_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_difference'"), recVars=[];
  $__res = (($GLOBALS['Data_Foldable_foldrArray'] ?? \PhpursThunks::eval('Data_Foldable_foldrArray')))((($GLOBALS['Data_Array_delete'] ?? \PhpursThunks::eval('Data_Array_delete')))($dictEq_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_cons__prime__'] = function() { $v = (function() {
  $__fn = function($x_0, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_cons'"), recVars=[];
  $__res = ((($GLOBALS['Data_Semigroup_concatArray'] ?? \PhpursThunks::eval('Data_Semigroup_concatArray')))([$x_0]))($xs_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_fromNonEmpty'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_fromNonEmpty"), recVars=[];
  $__res = ((($GLOBALS['Data_Semigroup_concatArray'] ?? \PhpursThunks::eval('Data_Semigroup_concatArray')))([($v_0)->value0]))(($v_0)->value1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_concatMap'] = function() { $v = (function() {
  $__fn = function($b_0, $a_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_concatMap"), recVars=[];
  $__res = ((($GLOBALS['Control_Bind_arrayBind'] ?? \PhpursThunks::eval('Control_Bind_arrayBind')))($a_1))($b_0);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_concat'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_Internal_NonEmptyArray'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_concat'] ?? \PhpursThunks::eval('Data_Array_concat'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray'))))((($GLOBALS['Data_Functor_arrayMap'] ?? \PhpursThunks::eval('Data_Functor_arrayMap')))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray')))))); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_appendArray'] = function() { $v = (function() {
  $__fn = function($xs_0, $ys_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_appendArray"), recVars=[];
  $__res = ((($GLOBALS['Data_Semigroup_concatArray'] ?? \PhpursThunks::eval('Data_Semigroup_concatArray')))($xs_0))($ys_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_alterAt'] = function() { $v = (function() {
  $__fn = function($i_0, $f_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_alterAt"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(((($GLOBALS['Data_Array_alterAt'] ?? \PhpursThunks::eval('Data_Array_alterAt')))($i_0))($f_1)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray')));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_head'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_NonEmpty_fromJust'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_fromJust'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_head'] ?? \PhpursThunks::eval('Data_Array_head'))))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray')))); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_init'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_NonEmpty_fromJust'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_fromJust'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_init'] ?? \PhpursThunks::eval('Data_Array_init'))))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray')))); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_last'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_NonEmpty_fromJust'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_fromJust'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_last'] ?? \PhpursThunks::eval('Data_Array_last'))))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray')))); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_tail'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_NonEmpty_fromJust'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_fromJust'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_tail'] ?? \PhpursThunks::eval('Data_Array_tail'))))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray')))); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_uncons'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_NonEmpty_fromJust'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_fromJust'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_uncons'] ?? \PhpursThunks::eval('Data_Array_uncons'))))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray')))); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_toNonEmpty'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = new Phpurs_Data2("NonEmpty", ($v_0)->head, ($v_0)->tail);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($GLOBALS['Data_Array_NonEmpty_uncons'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_uncons'))); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_unsnoc'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_NonEmpty_fromJust'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_fromJust'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_unsnoc'] ?? \PhpursThunks::eval('Data_Array_unsnoc'))))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray')))); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_all'] = function() { $v = function($p_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_all"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_Array_all'] ?? \PhpursThunks::eval('Data_Array_all')))($p_0)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_any'] = function() { $v = function($p_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_any"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_Array_any'] ?? \PhpursThunks::eval('Data_Array_any')))($p_0)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_catMaybes'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_Array_mapMaybe'] ?? \PhpursThunks::eval('Data_Array_mapMaybe')))((($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))->identity)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray'))); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_delete'] = function() { $v = (function() {
  $__fn = function($dictEq_0, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_delete"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(((($GLOBALS['Data_Array_delete'] ?? \PhpursThunks::eval('Data_Array_delete')))($dictEq_0))($x_1)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray')));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_deleteAt'] = function() { $v = function($i_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_deleteAt"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_Array_deleteAt'] ?? \PhpursThunks::eval('Data_Array_deleteAt')))($i_0)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_deleteBy'] = function() { $v = (function() {
  $__fn = function($f_0, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_deleteBy"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(((($GLOBALS['Data_Array_deleteBy'] ?? \PhpursThunks::eval('Data_Array_deleteBy')))($f_0))($x_1)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray')));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_difference'] = function() { $v = function($dictEq_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_difference"), recVars=[];
  $difference__prime__1_1_0 = (($GLOBALS['Data_Foldable_foldrArray'] ?? \PhpursThunks::eval('Data_Foldable_foldrArray')))((($GLOBALS['Data_Array_delete'] ?? \PhpursThunks::eval('Data_Array_delete')))($dictEq_0));
  $__res = function($xs_2) use ($difference__prime__1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($difference__prime__1_1_0)($xs_2)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_drop'] = function() { $v = function($i_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_drop"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_Array_drop'] ?? \PhpursThunks::eval('Data_Array_drop')))($i_0)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_dropEnd'] = function() { $v = function($i_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_dropEnd"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_Array_dropEnd'] ?? \PhpursThunks::eval('Data_Array_dropEnd')))($i_0)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_dropWhile'] = function() { $v = function($f_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_dropWhile"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_Array_dropWhile'] ?? \PhpursThunks::eval('Data_Array_dropWhile')))($f_0)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_elem'] = function() { $v = (function() {
  $__fn = function($dictEq_0, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_elem"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(((($GLOBALS['Data_Array_elem'] ?? \PhpursThunks::eval('Data_Array_elem')))($dictEq_0))($x_1)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray')));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_elemIndex'] = function() { $v = (function() {
  $__fn = function($dictEq_0, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_elemIndex"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(((($GLOBALS['Data_Array_elemIndex'] ?? \PhpursThunks::eval('Data_Array_elemIndex')))($dictEq_0))($x_1)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray')));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_elemLastIndex'] = function() { $v = (function() {
  $__fn = function($dictEq_0, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_elemLastIndex"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(((($GLOBALS['Data_Array_elemLastIndex'] ?? \PhpursThunks::eval('Data_Array_elemLastIndex')))($dictEq_0))($x_1)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray')));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_filter'] = function() { $v = function($f_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_filter"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_Array_filter'] ?? \PhpursThunks::eval('Data_Array_filter')))($f_0)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_filterA'] = function() { $v = function($dictApplicative_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_filterA"), recVars=[];
  $filterA1_1_0 = (($GLOBALS['Data_Array_filterA'] ?? \PhpursThunks::eval('Data_Array_filterA')))($dictApplicative_0);
  $__res = function($f_2) use ($filterA1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($filterA1_1_0)($f_2)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_find'] = function() { $v = function($p_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_find"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_Array_find'] ?? \PhpursThunks::eval('Data_Array_find')))($p_0)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_findIndex'] = function() { $v = function($p_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_findIndex"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_Array_findIndex'] ?? \PhpursThunks::eval('Data_Array_findIndex')))($p_0)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_findLastIndex'] = function() { $v = function($x_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_findLastIndex"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_Array_findLastIndex'] ?? \PhpursThunks::eval('Data_Array_findLastIndex')))($x_0)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_findMap'] = function() { $v = function($p_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_findMap"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_Array_findMap'] ?? \PhpursThunks::eval('Data_Array_findMap')))($p_0)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_foldM'] = function() { $v = (function() {
  $__fn = function($dictMonad_0, $f_1 = null, $acc_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_foldM"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((((($GLOBALS['Data_Array_foldM'] ?? \PhpursThunks::eval('Data_Array_foldM')))($dictMonad_0))($f_1))($acc_2)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray')));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_foldRecM'] = function() { $v = function($dictMonadRec_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_foldRecM"), recVars=[];
  $foldRecM1_1_0 = (($GLOBALS['Data_Array_foldRecM'] ?? \PhpursThunks::eval('Data_Array_foldRecM')))($dictMonadRec_0);
  $__res = (function() use ($foldRecM1_1_0) {
  $__fn = function($f_2, $acc_3 = null) use ($foldRecM1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($foldRecM1_1_0)($f_2))($acc_3)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray')));
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
\PhpursThunks::$thunks['Data_Array_NonEmpty_index'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_index'] ?? \PhpursThunks::eval('Data_Array_index'))))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray'))); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_length'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_length'] ?? \PhpursThunks::eval('Data_Array_length'))))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray'))); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_mapMaybe'] = function() { $v = function($f_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_mapMaybe"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_Array_mapMaybe'] ?? \PhpursThunks::eval('Data_Array_mapMaybe')))($f_0)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_notElem'] = function() { $v = (function() {
  $__fn = function($dictEq_0, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_notElem"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(((($GLOBALS['Data_Array_notElem'] ?? \PhpursThunks::eval('Data_Array_notElem')))($dictEq_0))($x_1)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray')));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_partition'] = function() { $v = function($f_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_partition"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_Array_partition'] ?? \PhpursThunks::eval('Data_Array_partition')))($f_0)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_slice'] = function() { $v = (function() {
  $__fn = function($start_0, $end_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_slice"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(((($GLOBALS['Data_Array_slice'] ?? \PhpursThunks::eval('Data_Array_slice')))($start_0))($end_1)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray')));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_span'] = function() { $v = function($f_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_span"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_Array_span'] ?? \PhpursThunks::eval('Data_Array_span')))($f_0)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_take'] = function() { $v = function($i_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_take"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_Array_take'] ?? \PhpursThunks::eval('Data_Array_take')))($i_0)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_takeEnd'] = function() { $v = function($i_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_takeEnd"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_Array_takeEnd'] ?? \PhpursThunks::eval('Data_Array_takeEnd')))($i_0)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_takeWhile'] = function() { $v = function($f_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_takeWhile"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_Array_takeWhile'] ?? \PhpursThunks::eval('Data_Array_takeWhile')))($f_0)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_toUnfoldable'] = function() { $v = function($dictUnfoldable_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_toUnfoldable"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(function($xs_1) use ($dictUnfoldable_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $len_2_0 = count($xs_1);
  $__res = ((($dictUnfoldable_0)->unfoldr)(function($i_3) use ($len_2_0, $xs_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if (($i_3 < $len_2_0)) {
$__t1 = new Phpurs_Data1("Just", new Phpurs_Data2("Tuple", ($xs_1)[0], ($i_3 + 1)));
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_cons'] = function() { $v = function($x_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_cons"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_Internal_NonEmptyArray'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_Array_cons'] ?? \PhpursThunks::eval('Data_Array_cons')))($x_0)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray'))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_group'] = function() { $v = function($dictEq_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_group"), recVars=[];
  $eq2_1_0 = ($dictEq_0)->eq;
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_Internal_NonEmptyArray'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(function($xs_2) use ($eq2_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Data_Array_groupBy'] ?? \PhpursThunks::eval('Data_Array_groupBy')))($eq2_1_0))($xs_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray'))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_groupAllBy'] = function() { $v = function($op_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_groupAllBy"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_Internal_NonEmptyArray'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_Array_groupAllBy'] ?? \PhpursThunks::eval('Data_Array_groupAllBy')))($op_0)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray'))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_groupAll'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_groupAll"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_Internal_NonEmptyArray'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_Array_groupAllBy'] ?? \PhpursThunks::eval('Data_Array_groupAllBy')))(($dictOrd_0)->compare)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray'))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_groupBy'] = function() { $v = function($op_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_groupBy"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_Internal_NonEmptyArray'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_Array_groupBy'] ?? \PhpursThunks::eval('Data_Array_groupBy')))($op_0)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray'))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_insert'] = function() { $v = (function() {
  $__fn = function($dictOrd_0, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_insert"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_Internal_NonEmptyArray'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(((($GLOBALS['Data_Array_insert'] ?? \PhpursThunks::eval('Data_Array_insert')))($dictOrd_0))($x_1)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray'))));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_insertBy'] = function() { $v = (function() {
  $__fn = function($f_0, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_insertBy"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_Internal_NonEmptyArray'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(((($GLOBALS['Data_Array_insertBy'] ?? \PhpursThunks::eval('Data_Array_insertBy')))($f_0))($x_1)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray'))));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_intersperse'] = function() { $v = function($x_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_intersperse"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_Internal_NonEmptyArray'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_Array_intersperse'] ?? \PhpursThunks::eval('Data_Array_intersperse')))($x_0)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray'))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_mapWithIndex'] = function() { $v = function($f_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_mapWithIndex"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_Internal_NonEmptyArray'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_FunctorWithIndex_mapWithIndexArray'] ?? \PhpursThunks::eval('Data_FunctorWithIndex_mapWithIndexArray')))($f_0)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray'))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_modifyAtIndices'] = function() { $v = function($dictFoldable_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_modifyAtIndices"), recVars=[];
  $modifyAtIndices1_1_0 = (($GLOBALS['Data_Array_modifyAtIndices'] ?? \PhpursThunks::eval('Data_Array_modifyAtIndices')))($dictFoldable_0);
  $__res = (function() use ($modifyAtIndices1_1_0) {
  $__fn = function($is_2, $f_3 = null) use ($modifyAtIndices1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_Internal_NonEmptyArray'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($modifyAtIndices1_1_0)($is_2))($f_3)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray'))));
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
\PhpursThunks::$thunks['Data_Array_NonEmpty_nub'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_nub"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_Internal_NonEmptyArray'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_Array_nub'] ?? \PhpursThunks::eval('Data_Array_nub')))($dictOrd_0)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray'))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_nubBy'] = function() { $v = function($f_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_nubBy"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_Internal_NonEmptyArray'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_Array_nubBy'] ?? \PhpursThunks::eval('Data_Array_nubBy')))($f_0)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray'))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_nubByEq'] = function() { $v = function($f_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_nubByEq"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_Internal_NonEmptyArray'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_Array_nubByEq'] ?? \PhpursThunks::eval('Data_Array_nubByEq')))($f_0)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray'))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_nubEq'] = function() { $v = function($dictEq_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_nubEq"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_Internal_NonEmptyArray'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_Array_nubEq'] ?? \PhpursThunks::eval('Data_Array_nubEq')))($dictEq_0)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray'))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_reverse'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_Internal_NonEmptyArray'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_reverse'] ?? \PhpursThunks::eval('Data_Array_reverse'))))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray')))); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_scanl'] = function() { $v = (function() {
  $__fn = function($f_0, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_scanl"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_Internal_NonEmptyArray'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(((($GLOBALS['Data_Array_scanl'] ?? \PhpursThunks::eval('Data_Array_scanl')))($f_0))($x_1)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray'))));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_scanr'] = function() { $v = (function() {
  $__fn = function($f_0, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_scanr"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_Internal_NonEmptyArray'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(((($GLOBALS['Data_Array_scanr'] ?? \PhpursThunks::eval('Data_Array_scanr')))($f_0))($x_1)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray'))));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_sort'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_sort"), recVars=[];
  $compare_1_0 = ($dictOrd_0)->compare;
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_Internal_NonEmptyArray'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(function($xs_2) use ($compare_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Data_Array_sortBy'] ?? \PhpursThunks::eval('Data_Array_sortBy')))($compare_1_0))($xs_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray'))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_sortBy'] = function() { $v = function($f_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_sortBy"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_Internal_NonEmptyArray'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_Array_sortBy'] ?? \PhpursThunks::eval('Data_Array_sortBy')))($f_0)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray'))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_sortWith'] = function() { $v = (function() {
  $__fn = function($dictOrd_0, $f_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_sortWith"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_Internal_NonEmptyArray'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(((($GLOBALS['Data_Array_sortWith'] ?? \PhpursThunks::eval('Data_Array_sortWith')))($dictOrd_0))($f_1)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray'))));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_updateAtIndices'] = function() { $v = function($dictFoldable_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_updateAtIndices"), recVars=[];
  $updateAtIndices1_1_0 = (($GLOBALS['Data_Array_updateAtIndices'] ?? \PhpursThunks::eval('Data_Array_updateAtIndices')))($dictFoldable_0);
  $__res = function($pairs_2) use ($updateAtIndices1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_Internal_NonEmptyArray'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($updateAtIndices1_1_0)($pairs_2)))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray'))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_unsafeIndex'] = function() { $v = function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_unsafeIndex"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_NonEmpty_unsafeIndex1'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_unsafeIndex1'))))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_unsafeIndex2'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_NonEmpty_unsafeIndex1'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_unsafeIndex1'))))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray'))); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_toUnfoldable1'] = function() { $v = (function() {
  $__fn = function($dictUnfoldable1_0, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_NonEmpty_toUnfoldable1"), recVars=[];
  $len_2_0 = (($GLOBALS['Data_Array_NonEmpty_length'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_length')))($xs_1);
  $__res = ((($dictUnfoldable1_0)->unfoldr1)(function($i_3) use ($len_2_0, $xs_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if (($i_3 < ($len_2_0 - 1))) {
$__t1 = new Phpurs_Data1("Just", ($i_3 + 1));
} else {
$__t1 = new Phpurs_Data0("Nothing");
};
  $__res = new Phpurs_Data2("Tuple", ((($GLOBALS['Data_Array_NonEmpty_unsafeIndex2'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_unsafeIndex2')))($xs_1))($i_3), $__t1);
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
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };













































































































