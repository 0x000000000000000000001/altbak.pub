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


// Data_Array_NonEmpty_max
$GLOBALS['Data_Array_NonEmpty_max'] = (function() {
  $__fn = function($x_0 = null, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v_2_0 = ((($GLOBALS['Data_Ord_ordInt'])['compare'])($x_0))($y_1);
  $__t1 = null;;
  if ((is_object($v_2_0) && (($v_2_0)->{'tag'} === "LT"))) {
$__t1 = $y_1;
goto end_branch_1;;
};
  if ((is_object($v_2_0) && (($v_2_0)->{'tag'} === "EQ"))) {
$__t1 = $x_0;
goto end_branch_1;;
};
  if ((is_object($v_2_0) && (($v_2_0)->{'tag'} === "GT"))) {
$__t1 = $x_0;
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_NonEmpty_fromJust
$GLOBALS['Data_Array_NonEmpty_fromJust'] = function($v_0 = null) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Just"))) {
$__t0 = ($v_0)->{'value0'};
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

// Data_Array_NonEmpty_unsafeIndex1
$GLOBALS['Data_Array_NonEmpty_unsafeIndex1'] = (function() {
  $__fn = function($__local_var_0 = null, $__local_var_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($__local_var_0)[0];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_NonEmpty_transpose
$GLOBALS['Data_Array_NonEmpty_transpose'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Unsafe_Coerce_unsafeCoerce']))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_transpose']))($GLOBALS['Unsafe_Coerce_unsafeCoerce']));

// Data_Array_NonEmpty_toArray
$GLOBALS['Data_Array_NonEmpty_toArray'] = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = $v_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_NonEmpty_unionBy'
$GLOBALS['Data_Array_NonEmpty_unionBy__prime__'] = (function() {
  $__fn = function($eq_0 = null, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray']))((($GLOBALS['Data_Array_unionBy'])($eq_0))($xs_1));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_NonEmpty_union'
$GLOBALS['Data_Array_NonEmpty_union__prime__'] = function($dictEq_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Array_NonEmpty_unionBy__prime__'])(($dictEq_0)['eq']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_NonEmpty_unionBy
$GLOBALS['Data_Array_NonEmpty_unionBy'] = (function() {
  $__fn = function($eq_0 = null, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray']))((($GLOBALS['Data_Array_unionBy'])($eq_0))($xs_1))))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_NonEmpty_union
$GLOBALS['Data_Array_NonEmpty_union'] = function($dictEq_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Array_NonEmpty_unionBy'])(($dictEq_0)['eq']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_NonEmpty_unzip
$GLOBALS['Data_Array_NonEmpty_unzip'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_0 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Tuple", ($v_0)->{'value0'}, ($v_0)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_unzip']))($GLOBALS['Data_Array_NonEmpty_toArray']));

// Data_Array_NonEmpty_updateAt
$GLOBALS['Data_Array_NonEmpty_updateAt'] = (function() {
  $__fn = function($i_0 = null, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Unsafe_Coerce_unsafeCoerce']))((($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Array_updateAt'])($i_0))($x_1)))($GLOBALS['Data_Array_NonEmpty_toArray']));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_NonEmpty_zip
$GLOBALS['Data_Array_NonEmpty_zip'] = (function() {
  $__fn = function($xs_0 = null, $ys_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Array_zipWithImpl'])($GLOBALS['Data_Tuple_Tuple'], $xs_0, $ys_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_NonEmpty_zipWith
$GLOBALS['Data_Array_NonEmpty_zipWith'] = (function() {
  $__fn = function($f_0 = null, $xs_1 = null, $ys_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($GLOBALS['Data_Array_zipWithImpl'])($f_0, $xs_1, $ys_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_Array_NonEmpty_zipWithA
$GLOBALS['Data_Array_NonEmpty_zipWithA'] = function($dictApplicative_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Array_zipWithA'])($dictApplicative_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_NonEmpty_splitAt
$GLOBALS['Data_Array_NonEmpty_splitAt'] = (function() {
  $__fn = function($i_0 = null, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_Array_splitAt'])($i_0))($xs_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_NonEmpty_some
$GLOBALS['Data_Array_NonEmpty_some'] = (function() {
  $__fn = function($dictAlternative_0 = null, $dictLazy_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Unsafe_Coerce_unsafeCoerce']))((($GLOBALS['Data_Array_some'])($dictAlternative_0))($dictLazy_1));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_NonEmpty_snoc'
$GLOBALS['Data_Array_NonEmpty_snoc__prime__'] = (function() {
  $__fn = function($xs_0 = null, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Control_Monad_ST_Internal_run'])((($GLOBALS['Data_Array_ST_withArray'])(($GLOBALS['Data_Array_ST_push'])($x_1)))($xs_0));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_NonEmpty_snoc
$GLOBALS['Data_Array_NonEmpty_snoc'] = (function() {
  $__fn = function($xs_0 = null, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Control_Monad_ST_Internal_run'])((($GLOBALS['Data_Array_ST_withArray'])(($GLOBALS['Data_Array_ST_push'])($x_1)))($xs_0));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_NonEmpty_singleton
$GLOBALS['Data_Array_NonEmpty_singleton'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray']))($GLOBALS['Data_Array_singleton']);

// Data_Array_NonEmpty_replicate
$GLOBALS['Data_Array_NonEmpty_replicate'] = (function() {
  $__fn = function($i_0 = null, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Array_replicateImpl'])((($GLOBALS['Data_Array_NonEmpty_max'])(1))($i_0), $x_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_NonEmpty_range
$GLOBALS['Data_Array_NonEmpty_range'] = (function() {
  $__fn = function($x_0 = null, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Array_rangeImpl'])($x_0, $y_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_NonEmpty_prependArray
$GLOBALS['Data_Array_NonEmpty_prependArray'] = (function() {
  $__fn = function($xs_0 = null, $ys_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_Semigroup_concatArray'])($xs_0))($ys_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_NonEmpty_modifyAt
$GLOBALS['Data_Array_NonEmpty_modifyAt'] = (function() {
  $__fn = function($i_0 = null, $f_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Unsafe_Coerce_unsafeCoerce']))((($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Array_modifyAt'])($i_0))($f_1)))($GLOBALS['Data_Array_NonEmpty_toArray']));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_NonEmpty_intersectBy'
$GLOBALS['Data_Array_NonEmpty_intersectBy__prime__'] = (function() {
  $__fn = function($eq_0 = null, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_Array_intersectBy'])($eq_0))($xs_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_NonEmpty_intersectBy
$GLOBALS['Data_Array_NonEmpty_intersectBy'] = (function() {
  $__fn = function($eq_0 = null, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Array_NonEmpty_intersectBy__prime__'])($eq_0))($xs_1)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_NonEmpty_intersect'
$GLOBALS['Data_Array_NonEmpty_intersect__prime__'] = function($dictEq_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Array_NonEmpty_intersectBy__prime__'])(($dictEq_0)['eq']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_NonEmpty_intersect
$GLOBALS['Data_Array_NonEmpty_intersect'] = function($dictEq_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Array_NonEmpty_intersectBy'])(($dictEq_0)['eq']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_NonEmpty_intercalate
$GLOBALS['Data_Array_NonEmpty_intercalate'] = function($dictSemigroup_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($GLOBALS['Data_Semigroup_Foldable_intercalateMap'])($GLOBALS['Data_Array_NonEmpty_Internal_foldable1NonEmptyArray']))($dictSemigroup_0);
  $__res = function($a_2 = null) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = (($__local_var_1_0)($a_2))(($GLOBALS['Control_Category_categoryFn'])['identity']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_NonEmpty_insertAt
$GLOBALS['Data_Array_NonEmpty_insertAt'] = (function() {
  $__fn = function($i_0 = null, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Unsafe_Coerce_unsafeCoerce']))((($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Array_insertAt'])($i_0))($x_1)))($GLOBALS['Data_Array_NonEmpty_toArray']));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_NonEmpty_fromFoldable1
$GLOBALS['Data_Array_NonEmpty_fromFoldable1'] = function($dictFoldable1_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = ((($dictFoldable1_0)['Foldable0'])($GLOBALS['Prim_undefined']))['foldr'];
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray']))(function($__local_var_2 = null) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Array_fromFoldableImpl'])($__local_var_1_0, $__local_var_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_NonEmpty_fromArray
$GLOBALS['Data_Array_NonEmpty_fromArray'] = function($xs_0 = null) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((count($xs_0) > 0)) {
$__t0 = new Phpurs_Data1("Just", $xs_0);
goto end_branch_0;;
};
  $__t0 = new Phpurs_Data0("Nothing");
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_NonEmpty_fromFoldable
$GLOBALS['Data_Array_NonEmpty_fromFoldable'] = function($dictFoldable_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = ($dictFoldable_0)['foldr'];
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_fromArray']))(function($__local_var_2 = null) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Array_fromFoldableImpl'])($__local_var_1_0, $__local_var_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_NonEmpty_transpose'
$GLOBALS['Data_Array_NonEmpty_transpose__prime__'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_fromArray']))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_transpose']))($GLOBALS['Unsafe_Coerce_unsafeCoerce']));

// Data_Array_NonEmpty_foldr1
$GLOBALS['Data_Array_NonEmpty_foldr1'] = ($GLOBALS['Data_Array_NonEmpty_Internal_foldable1NonEmptyArray'])['foldr1'];

// Data_Array_NonEmpty_foldl1
$GLOBALS['Data_Array_NonEmpty_foldl1'] = ($GLOBALS['Data_Array_NonEmpty_Internal_foldable1NonEmptyArray'])['foldl1'];

// Data_Array_NonEmpty_foldMap1
$GLOBALS['Data_Array_NonEmpty_foldMap1'] = function($dictSemigroup_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Semigroup_Foldable_foldMap1DefaultL'])($GLOBALS['Data_Array_NonEmpty_Internal_foldable1NonEmptyArray']))($GLOBALS['Data_Functor_functorArray']))($dictSemigroup_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_NonEmpty_fold1
$GLOBALS['Data_Array_NonEmpty_fold1'] = function($dictSemigroup_0 = null) {
  $__num = \func_num_args();
  $__res = (((($GLOBALS['Data_Semigroup_Foldable_foldMap1DefaultL'])($GLOBALS['Data_Array_NonEmpty_Internal_foldable1NonEmptyArray']))($GLOBALS['Data_Functor_functorArray']))($dictSemigroup_0))(($GLOBALS['Control_Category_categoryFn'])['identity']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_NonEmpty_difference'
$GLOBALS['Data_Array_NonEmpty_difference__prime__'] = function($dictEq_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Foldable_foldrArray'])(($GLOBALS['Data_Array_delete'])($dictEq_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_NonEmpty_cons'
$GLOBALS['Data_Array_NonEmpty_cons__prime__'] = (function() {
  $__fn = function($x_0 = null, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_Semigroup_concatArray'])([$x_0]))($xs_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_NonEmpty_fromNonEmpty
$GLOBALS['Data_Array_NonEmpty_fromNonEmpty'] = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Semigroup_concatArray'])([($v_0)->{'value0'}]))(($v_0)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_NonEmpty_concatMap
$GLOBALS['Data_Array_NonEmpty_concatMap'] = (function() {
  $__fn = function($b_0 = null, $a_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Bind_arrayBind'])($a_1))($b_0);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_NonEmpty_concat
$GLOBALS['Data_Array_NonEmpty_concat'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray']))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_concat']))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_toArray']))(($GLOBALS['Data_Functor_arrayMap'])($GLOBALS['Data_Array_NonEmpty_toArray']))));

// Data_Array_NonEmpty_appendArray
$GLOBALS['Data_Array_NonEmpty_appendArray'] = (function() {
  $__fn = function($xs_0 = null, $ys_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_Semigroup_concatArray'])($xs_0))($ys_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_NonEmpty_alterAt
$GLOBALS['Data_Array_NonEmpty_alterAt'] = (function() {
  $__fn = function($i_0 = null, $f_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Array_alterAt'])($i_0))($f_1)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_NonEmpty_head
$GLOBALS['Data_Array_NonEmpty_head'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_fromJust']))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_head']))($GLOBALS['Data_Array_NonEmpty_toArray']));

// Data_Array_NonEmpty_init
$GLOBALS['Data_Array_NonEmpty_init'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_fromJust']))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_init']))($GLOBALS['Data_Array_NonEmpty_toArray']));

// Data_Array_NonEmpty_last
$GLOBALS['Data_Array_NonEmpty_last'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_fromJust']))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_last']))($GLOBALS['Data_Array_NonEmpty_toArray']));

// Data_Array_NonEmpty_tail
$GLOBALS['Data_Array_NonEmpty_tail'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_fromJust']))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_tail']))($GLOBALS['Data_Array_NonEmpty_toArray']));

// Data_Array_NonEmpty_uncons
$GLOBALS['Data_Array_NonEmpty_uncons'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_fromJust']))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_uncons']))($GLOBALS['Data_Array_NonEmpty_toArray']));

// Data_Array_NonEmpty_toNonEmpty
$GLOBALS['Data_Array_NonEmpty_toNonEmpty'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_0 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("NonEmpty", ($v_0)['head'], ($v_0)['tail']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Array_NonEmpty_uncons']);

// Data_Array_NonEmpty_unsnoc
$GLOBALS['Data_Array_NonEmpty_unsnoc'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_fromJust']))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_unsnoc']))($GLOBALS['Data_Array_NonEmpty_toArray']));

// Data_Array_NonEmpty_all
$GLOBALS['Data_Array_NonEmpty_all'] = function($p_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_all'])($p_0)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_NonEmpty_any
$GLOBALS['Data_Array_NonEmpty_any'] = function($p_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_any'])($p_0)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_NonEmpty_catMaybes
$GLOBALS['Data_Array_NonEmpty_catMaybes'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_mapMaybe'])(($GLOBALS['Control_Category_categoryFn'])['identity'])))($GLOBALS['Data_Array_NonEmpty_toArray']);

// Data_Array_NonEmpty_delete
$GLOBALS['Data_Array_NonEmpty_delete'] = (function() {
  $__fn = function($dictEq_0 = null, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Array_delete'])($dictEq_0))($x_1)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_NonEmpty_deleteAt
$GLOBALS['Data_Array_NonEmpty_deleteAt'] = function($i_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_deleteAt'])($i_0)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_NonEmpty_deleteBy
$GLOBALS['Data_Array_NonEmpty_deleteBy'] = (function() {
  $__fn = function($f_0 = null, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Array_deleteBy'])($f_0))($x_1)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_NonEmpty_difference
$GLOBALS['Data_Array_NonEmpty_difference'] = function($dictEq_0 = null) {
  $__num = \func_num_args();
  $difference__prime__1_1_0 = ($GLOBALS['Data_Foldable_foldrArray'])(($GLOBALS['Data_Array_delete'])($dictEq_0));
  $__res = function($xs_2 = null) use ($difference__prime__1_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($difference__prime__1_1_0)($xs_2)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_NonEmpty_drop
$GLOBALS['Data_Array_NonEmpty_drop'] = function($i_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_drop'])($i_0)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_NonEmpty_dropEnd
$GLOBALS['Data_Array_NonEmpty_dropEnd'] = function($i_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_dropEnd'])($i_0)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_NonEmpty_dropWhile
$GLOBALS['Data_Array_NonEmpty_dropWhile'] = function($f_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_dropWhile'])($f_0)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_NonEmpty_elem
$GLOBALS['Data_Array_NonEmpty_elem'] = (function() {
  $__fn = function($dictEq_0 = null, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Array_elem'])($dictEq_0))($x_1)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_NonEmpty_elemIndex
$GLOBALS['Data_Array_NonEmpty_elemIndex'] = (function() {
  $__fn = function($dictEq_0 = null, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Array_elemIndex'])($dictEq_0))($x_1)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_NonEmpty_elemLastIndex
$GLOBALS['Data_Array_NonEmpty_elemLastIndex'] = (function() {
  $__fn = function($dictEq_0 = null, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Array_elemLastIndex'])($dictEq_0))($x_1)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_NonEmpty_filter
$GLOBALS['Data_Array_NonEmpty_filter'] = function($f_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_filter'])($f_0)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_NonEmpty_filterA
$GLOBALS['Data_Array_NonEmpty_filterA'] = function($dictApplicative_0 = null) {
  $__num = \func_num_args();
  $filterA1_1_0 = ($GLOBALS['Data_Array_filterA'])($dictApplicative_0);
  $__res = function($f_2 = null) use ($filterA1_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($filterA1_1_0)($f_2)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_NonEmpty_find
$GLOBALS['Data_Array_NonEmpty_find'] = function($p_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_find'])($p_0)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_NonEmpty_findIndex
$GLOBALS['Data_Array_NonEmpty_findIndex'] = function($p_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_findIndex'])($p_0)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_NonEmpty_findLastIndex
$GLOBALS['Data_Array_NonEmpty_findLastIndex'] = function($x_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_findLastIndex'])($x_0)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_NonEmpty_findMap
$GLOBALS['Data_Array_NonEmpty_findMap'] = function($p_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_findMap'])($p_0)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_NonEmpty_foldM
$GLOBALS['Data_Array_NonEmpty_foldM'] = (function() {
  $__fn = function($dictMonad_0 = null, $f_1 = null, $acc_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(((($GLOBALS['Data_Array_foldM'])($dictMonad_0))($f_1))($acc_2)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_Array_NonEmpty_foldRecM
$GLOBALS['Data_Array_NonEmpty_foldRecM'] = function($dictMonadRec_0 = null) {
  $__num = \func_num_args();
  $foldRecM1_1_0 = ($GLOBALS['Data_Array_foldRecM'])($dictMonadRec_0);
  $__res = (function() use ($foldRecM1_1_0) {
  $__fn = function($f_2 = null, $acc_3 = null) use ($foldRecM1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($foldRecM1_1_0)($f_2))($acc_3)))($GLOBALS['Data_Array_NonEmpty_toArray']);
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

// Data_Array_NonEmpty_index
$GLOBALS['Data_Array_NonEmpty_index'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_index']))($GLOBALS['Data_Array_NonEmpty_toArray']);

// Data_Array_NonEmpty_length
$GLOBALS['Data_Array_NonEmpty_length'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_length']))($GLOBALS['Data_Array_NonEmpty_toArray']);

// Data_Array_NonEmpty_mapMaybe
$GLOBALS['Data_Array_NonEmpty_mapMaybe'] = function($f_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_mapMaybe'])($f_0)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_NonEmpty_notElem
$GLOBALS['Data_Array_NonEmpty_notElem'] = (function() {
  $__fn = function($dictEq_0 = null, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Array_notElem'])($dictEq_0))($x_1)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_NonEmpty_partition
$GLOBALS['Data_Array_NonEmpty_partition'] = function($f_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_partition'])($f_0)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_NonEmpty_slice
$GLOBALS['Data_Array_NonEmpty_slice'] = (function() {
  $__fn = function($start_0 = null, $end_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Array_slice'])($start_0))($end_1)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_NonEmpty_span
$GLOBALS['Data_Array_NonEmpty_span'] = function($f_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_span'])($f_0)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_NonEmpty_take
$GLOBALS['Data_Array_NonEmpty_take'] = function($i_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_take'])($i_0)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_NonEmpty_takeEnd
$GLOBALS['Data_Array_NonEmpty_takeEnd'] = function($i_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_takeEnd'])($i_0)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_NonEmpty_takeWhile
$GLOBALS['Data_Array_NonEmpty_takeWhile'] = function($f_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_takeWhile'])($f_0)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_NonEmpty_toUnfoldable
$GLOBALS['Data_Array_NonEmpty_toUnfoldable'] = function($dictUnfoldable_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($xs_1 = null) use ($dictUnfoldable_0) {
  $__num = \func_num_args();
  $len_2_0 = count($xs_1);
  $__res = ((($dictUnfoldable_0)['unfoldr'])(function($i_3 = null) use ($len_2_0, $xs_1) {
  $__num = \func_num_args();
  $__t1 = null;;
  if (($i_3 < $len_2_0)) {
$__t1 = new Phpurs_Data1("Just", new Phpurs_Data2("Tuple", ($xs_1)[0], ($i_3 + 1)));
goto end_branch_1;;
};
  $__t1 = new Phpurs_Data0("Nothing");
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_NonEmpty_cons
$GLOBALS['Data_Array_NonEmpty_cons'] = function($x_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray']))((($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_cons'])($x_0)))($GLOBALS['Data_Array_NonEmpty_toArray']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_NonEmpty_group
$GLOBALS['Data_Array_NonEmpty_group'] = function($dictEq_0 = null) {
  $__num = \func_num_args();
  $eq2_1_0 = ($dictEq_0)['eq'];
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray']))((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($xs_2 = null) use ($eq2_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Array_groupBy'])($eq2_1_0))($xs_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Array_NonEmpty_toArray']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_NonEmpty_groupAllBy
$GLOBALS['Data_Array_NonEmpty_groupAllBy'] = function($op_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray']))((($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_groupAllBy'])($op_0)))($GLOBALS['Data_Array_NonEmpty_toArray']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_NonEmpty_groupAll
$GLOBALS['Data_Array_NonEmpty_groupAll'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray']))((($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_groupAllBy'])(($dictOrd_0)['compare'])))($GLOBALS['Data_Array_NonEmpty_toArray']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_NonEmpty_groupBy
$GLOBALS['Data_Array_NonEmpty_groupBy'] = function($op_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray']))((($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_groupBy'])($op_0)))($GLOBALS['Data_Array_NonEmpty_toArray']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_NonEmpty_insert
$GLOBALS['Data_Array_NonEmpty_insert'] = (function() {
  $__fn = function($dictOrd_0 = null, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray']))((($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Array_insert'])($dictOrd_0))($x_1)))($GLOBALS['Data_Array_NonEmpty_toArray']));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_NonEmpty_insertBy
$GLOBALS['Data_Array_NonEmpty_insertBy'] = (function() {
  $__fn = function($f_0 = null, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray']))((($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Array_insertBy'])($f_0))($x_1)))($GLOBALS['Data_Array_NonEmpty_toArray']));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_NonEmpty_intersperse
$GLOBALS['Data_Array_NonEmpty_intersperse'] = function($x_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray']))((($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_intersperse'])($x_0)))($GLOBALS['Data_Array_NonEmpty_toArray']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_NonEmpty_mapWithIndex
$GLOBALS['Data_Array_NonEmpty_mapWithIndex'] = function($f_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray']))((($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_FunctorWithIndex_mapWithIndexArray'])($f_0)))($GLOBALS['Data_Array_NonEmpty_toArray']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_NonEmpty_modifyAtIndices
$GLOBALS['Data_Array_NonEmpty_modifyAtIndices'] = function($dictFoldable_0 = null) {
  $__num = \func_num_args();
  $modifyAtIndices1_1_0 = ($GLOBALS['Data_Array_modifyAtIndices'])($dictFoldable_0);
  $__res = (function() use ($modifyAtIndices1_1_0) {
  $__fn = function($is_2 = null, $f_3 = null) use ($modifyAtIndices1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray']))((($GLOBALS['Control_Semigroupoid_composeImpl'])((($modifyAtIndices1_1_0)($is_2))($f_3)))($GLOBALS['Data_Array_NonEmpty_toArray']));
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

// Data_Array_NonEmpty_nub
$GLOBALS['Data_Array_NonEmpty_nub'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray']))((($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_nub'])($dictOrd_0)))($GLOBALS['Data_Array_NonEmpty_toArray']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_NonEmpty_nubBy
$GLOBALS['Data_Array_NonEmpty_nubBy'] = function($f_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray']))((($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_nubBy'])($f_0)))($GLOBALS['Data_Array_NonEmpty_toArray']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_NonEmpty_nubByEq
$GLOBALS['Data_Array_NonEmpty_nubByEq'] = function($f_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray']))((($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_nubByEq'])($f_0)))($GLOBALS['Data_Array_NonEmpty_toArray']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_NonEmpty_nubEq
$GLOBALS['Data_Array_NonEmpty_nubEq'] = function($dictEq_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray']))((($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_nubEq'])($dictEq_0)))($GLOBALS['Data_Array_NonEmpty_toArray']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_NonEmpty_reverse
$GLOBALS['Data_Array_NonEmpty_reverse'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray']))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_reverse']))($GLOBALS['Data_Array_NonEmpty_toArray']));

// Data_Array_NonEmpty_scanl
$GLOBALS['Data_Array_NonEmpty_scanl'] = (function() {
  $__fn = function($f_0 = null, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray']))((($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Array_scanl'])($f_0))($x_1)))($GLOBALS['Data_Array_NonEmpty_toArray']));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_NonEmpty_scanr
$GLOBALS['Data_Array_NonEmpty_scanr'] = (function() {
  $__fn = function($f_0 = null, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray']))((($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Array_scanr'])($f_0))($x_1)))($GLOBALS['Data_Array_NonEmpty_toArray']));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_NonEmpty_sort
$GLOBALS['Data_Array_NonEmpty_sort'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $compare_1_0 = ($dictOrd_0)['compare'];
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray']))((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($xs_2 = null) use ($compare_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Array_sortBy'])($compare_1_0))($xs_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Array_NonEmpty_toArray']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_NonEmpty_sortBy
$GLOBALS['Data_Array_NonEmpty_sortBy'] = function($f_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray']))((($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_sortBy'])($f_0)))($GLOBALS['Data_Array_NonEmpty_toArray']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_NonEmpty_sortWith
$GLOBALS['Data_Array_NonEmpty_sortWith'] = (function() {
  $__fn = function($dictOrd_0 = null, $f_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray']))((($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Array_sortWith'])($dictOrd_0))($f_1)))($GLOBALS['Data_Array_NonEmpty_toArray']));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_NonEmpty_updateAtIndices
$GLOBALS['Data_Array_NonEmpty_updateAtIndices'] = function($dictFoldable_0 = null) {
  $__num = \func_num_args();
  $updateAtIndices1_1_0 = ($GLOBALS['Data_Array_updateAtIndices'])($dictFoldable_0);
  $__res = function($pairs_2 = null) use ($updateAtIndices1_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_Internal_NonEmptyArray']))((($GLOBALS['Control_Semigroupoid_composeImpl'])(($updateAtIndices1_1_0)($pairs_2)))($GLOBALS['Data_Array_NonEmpty_toArray']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_NonEmpty_unsafeIndex
$GLOBALS['Data_Array_NonEmpty_unsafeIndex'] = function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_unsafeIndex1']))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_NonEmpty_unsafeIndex2
$GLOBALS['Data_Array_NonEmpty_unsafeIndex2'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_unsafeIndex1']))($GLOBALS['Data_Array_NonEmpty_toArray']);

// Data_Array_NonEmpty_toUnfoldable1
$GLOBALS['Data_Array_NonEmpty_toUnfoldable1'] = (function() {
  $__fn = function($dictUnfoldable1_0 = null, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $len_2_0 = ($GLOBALS['Data_Array_NonEmpty_length'])($xs_1);
  $__res = ((($dictUnfoldable1_0)['unfoldr1'])(function($i_3 = null) use ($len_2_0, $xs_1) {
  $__num = \func_num_args();
  $__t1 = null;;
  if (($i_3 < ($len_2_0 - 1))) {
$__t1 = new Phpurs_Data1("Just", ($i_3 + 1));
goto end_branch_1;;
};
  $__t1 = new Phpurs_Data0("Nothing");
  end_branch_1:;
  $__res = new Phpurs_Data2("Tuple", (($GLOBALS['Data_Array_NonEmpty_unsafeIndex2'])($xs_1))($i_3), $__t1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(0);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

