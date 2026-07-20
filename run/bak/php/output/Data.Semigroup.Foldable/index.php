<?php

namespace Data\Semigroup\Foldable;

// ALL IMPORTS: Control.Apply, Control.Category, Control.Semigroupoid, Data.Eq, Data.Foldable, Data.Function, Data.Functor, Data.Identity, Data.Monoid.Dual, Data.Monoid.Multiplicative, Data.Newtype, Data.Ord.Max, Data.Ord.Min, Data.Ordering, Data.Semigroup, Data.Semigroup.Foldable, Data.Tuple, Data.Unit, Prelude, Prim
// TO REQUIRE: Control.Apply, Control.Category, Control.Semigroupoid, Data.Eq, Data.Foldable, Data.Function, Data.Functor, Data.Identity, Data.Monoid.Dual, Data.Monoid.Multiplicative, Data.Newtype, Data.Ord.Max, Data.Ord.Min, Data.Ordering, Data.Semigroup, Data.Semigroup.Foldable, Data.Tuple, Data.Unit, Prelude
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Category/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Foldable/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Identity/index.php';
require_once __DIR__ . '/../Data.Monoid.Dual/index.php';
require_once __DIR__ . '/../Data.Monoid.Multiplicative/index.php';
require_once __DIR__ . '/../Data.Newtype/index.php';
require_once __DIR__ . '/../Data.Ord.Max/index.php';
require_once __DIR__ . '/../Data.Ord.Min/index.php';
require_once __DIR__ . '/../Data.Ordering/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Semigroup.Foldable/index.php';
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
\PhpursThunks::$thunks['Data_Semigroup_Foldable_JoinWith'] = function() { $v = function($x_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Semigroup_Foldable_JoinWith"), recVars=[];
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Semigroup_Foldable_FoldRight1'] = function() { $v = (function() {
  $__fn = function($value0, $value1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("FoldRight1", $value0, $value1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Semigroup_Foldable_Act'] = function() { $v = function($x_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Semigroup_Foldable_Act"), recVars=[];
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Semigroup_Foldable_runFoldRight1'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Semigroup_Foldable_runFoldRight1"), recVars=[];
  $__res = (($v_0)->value0)(($v_0)->value1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Semigroup_Foldable_mkFoldRight1'] = function() { $v = (($GLOBALS['Data_Semigroup_Foldable_FoldRight1'] ?? \PhpursThunks::eval('Data_Semigroup_Foldable_FoldRight1')))(($GLOBALS['Data_Function_const'] ?? \PhpursThunks::eval('Data_Function_const'))); return $v; };
\PhpursThunks::$thunks['Data_Semigroup_Foldable_foldr1'] = function() { $v = function($dict_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Semigroup_Foldable_foldr1"), recVars=[];
  $__res = ($dict_0)->foldr1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Semigroup_Foldable_foldl1'] = function() { $v = function($dict_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Semigroup_Foldable_foldl1"), recVars=[];
  $__res = ($dict_0)->foldl1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Semigroup_Foldable_maximumBy'] = function() { $v = (function() {
  $__fn = function($dictFoldable1_0, $cmp_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Semigroup_Foldable_maximumBy"), recVars=[];
  $__res = (($dictFoldable1_0)->foldl1)((function() use ($cmp_1) {
  $__fn = function($x_2, $y_3 = null) use ($cmp_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object((($cmp_1)($x_2))($y_3)) && (((($cmp_1)($x_2))($y_3))->tag === "GT"))) {
$__t0 = $x_2;
} else {
$__t0 = $y_3;
};
  $__res = $__t0;
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
\PhpursThunks::$thunks['Data_Semigroup_Foldable_minimumBy'] = function() { $v = (function() {
  $__fn = function($dictFoldable1_0, $cmp_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Semigroup_Foldable_minimumBy"), recVars=[];
  $__res = (($dictFoldable1_0)->foldl1)((function() use ($cmp_1) {
  $__fn = function($x_2, $y_3 = null) use ($cmp_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object((($cmp_1)($x_2))($y_3)) && (((($cmp_1)($x_2))($y_3))->tag === "LT"))) {
$__t0 = $x_2;
} else {
$__t0 = $y_3;
};
  $__res = $__t0;
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
\PhpursThunks::$thunks['Data_Semigroup_Foldable_foldableTuple'] = function() { $v = (object)["foldMap1" => (function() {
  $__fn = function($dictSemigroup_0, $f_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($f_1)(($v_2)->value1);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldr1" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($v1_1)->value1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "foldl1" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($v1_1)->value1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Foldable0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_Foldable_foldableTuple'] ?? \PhpursThunks::eval('Data_Foldable_foldableTuple'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Semigroup_Foldable_foldableMultiplicative'] = function() { $v = (object)["foldr1" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $v1_1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "foldl1" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $v1_1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "foldMap1" => (function() {
  $__fn = function($dictSemigroup_0, $f_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($f_1)($v_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "Foldable0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_Foldable_foldableMultiplicative'] ?? \PhpursThunks::eval('Data_Foldable_foldableMultiplicative'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Semigroup_Foldable_foldableIdentity'] = function() { $v = (object)["foldMap1" => (function() {
  $__fn = function($dictSemigroup_0, $f_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($f_1)($v_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldl1" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $v1_1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "foldr1" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $v1_1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Foldable0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_Foldable_foldableIdentity'] ?? \PhpursThunks::eval('Data_Foldable_foldableIdentity'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Semigroup_Foldable_foldableDual'] = function() { $v = (object)["foldr1" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $v1_1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "foldl1" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $v1_1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "foldMap1" => (function() {
  $__fn = function($dictSemigroup_0, $f_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($f_1)($v_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "Foldable0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_Foldable_foldableDual'] ?? \PhpursThunks::eval('Data_Foldable_foldableDual'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Semigroup_Foldable_foldRight1Semigroup'] = function() { $v = (object)["append" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__local_var_2_0 = ($v_0)->value1;
  $__res = new Phpurs_Data2("FoldRight1", (function() use ($__local_var_2_0, $v1_1, $v_0) {
  $__fn = function($a_3, $f_4 = null) use ($__local_var_2_0, $v1_1, $v_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($v_0)->value0)((($f_4)($__local_var_2_0))(((($v1_1)->value0)($a_3))($f_4))))($f_4);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), ($v1_1)->value1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Data_Semigroup_Foldable_semigroupDual'] = function() { $v = (object)["append" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__local_var_2_0 = ($v1_1)->value1;
  $__res = new Phpurs_Data2("FoldRight1", (function() use ($__local_var_2_0, $v1_1, $v_0) {
  $__fn = function($a_3, $f_4 = null) use ($__local_var_2_0, $v1_1, $v_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($v1_1)->value0)((($f_4)($__local_var_2_0))(((($v_0)->value0)($a_3))($f_4))))($f_4);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), ($v_0)->value1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Data_Semigroup_Foldable_foldMap1DefaultR'] = function() { $v = (function() {
  $__fn = function($dictFoldable1_0, $dictFunctor_1 = null, $dictSemigroup_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Semigroup_Foldable_foldMap1DefaultR"), recVars=[];
  $append_3_0 = ($dictSemigroup_2)->append;
  $__res = function($f_4) use ($append_3_0, $dictFoldable1_0, $dictFunctor_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($dictFoldable1_0)->foldr1)($append_3_0)))((($dictFunctor_1)->map)($f_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Semigroup_Foldable_foldMap1DefaultL'] = function() { $v = (function() {
  $__fn = function($dictFoldable1_0, $dictFunctor_1 = null, $dictSemigroup_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Semigroup_Foldable_foldMap1DefaultL"), recVars=[];
  $append_3_0 = ($dictSemigroup_2)->append;
  $__res = function($f_4) use ($append_3_0, $dictFoldable1_0, $dictFunctor_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($dictFoldable1_0)->foldl1)($append_3_0)))((($dictFunctor_1)->map)($f_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Semigroup_Foldable_foldMap1'] = function() { $v = function($dict_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Semigroup_Foldable_foldMap1"), recVars=[];
  $__res = ($dict_0)->foldMap1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Semigroup_Foldable_foldl1Default'] = function() { $v = function($dictFoldable1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Semigroup_Foldable_foldl1Default"), recVars=[];
  $__local_var_1_0 = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Semigroup_Foldable_runFoldRight1'] ?? \PhpursThunks::eval('Data_Semigroup_Foldable_runFoldRight1'))))(((($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce')))((($dictFoldable1_0)->foldMap1)(($GLOBALS['Data_Semigroup_Foldable_semigroupDual'] ?? \PhpursThunks::eval('Data_Semigroup_Foldable_semigroupDual')))))(($GLOBALS['Data_Semigroup_Foldable_mkFoldRight1'] ?? \PhpursThunks::eval('Data_Semigroup_Foldable_mkFoldRight1'))));
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((function() use ($__local_var_1_0) {
  $__fn = function($b_2, $a_3 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($__local_var_1_0)($a_3))($b_2);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(($GLOBALS['Data_Function_flip'] ?? \PhpursThunks::eval('Data_Function_flip')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Semigroup_Foldable_foldr1Default'] = function() { $v = function($dictFoldable1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Semigroup_Foldable_foldr1Default"), recVars=[];
  $__local_var_1_0 = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Semigroup_Foldable_runFoldRight1'] ?? \PhpursThunks::eval('Data_Semigroup_Foldable_runFoldRight1'))))(((($dictFoldable1_0)->foldMap1)(($GLOBALS['Data_Semigroup_Foldable_foldRight1Semigroup'] ?? \PhpursThunks::eval('Data_Semigroup_Foldable_foldRight1Semigroup'))))(($GLOBALS['Data_Semigroup_Foldable_mkFoldRight1'] ?? \PhpursThunks::eval('Data_Semigroup_Foldable_mkFoldRight1'))));
  $__res = (function() use ($__local_var_1_0) {
  $__fn = function($b_2, $a_3 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($__local_var_1_0)($a_3))($b_2);
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
\PhpursThunks::$thunks['Data_Semigroup_Foldable_intercalateMap'] = function() { $v = (function() {
  $__fn = function($dictFoldable1_0, $dictSemigroup_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Semigroup_Foldable_intercalateMap"), recVars=[];
  $foldMap12_2_0 = (($dictFoldable1_0)->foldMap1)((object)["append" => (function() use ($dictSemigroup_1) {
  $__fn = function($v_2, $v1_3 = null, $j_4 = null) use ($dictSemigroup_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($dictSemigroup_1)->append)(($v_2)($j_4)))(((($dictSemigroup_1)->append)($j_4))(($v1_3)($j_4)));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()]);
  $__res = (function() use ($foldMap12_2_0) {
  $__fn = function($j_3, $f_4 = null, $foldable_5 = null) use ($foldMap12_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($foldMap12_2_0)(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Semigroup_Foldable_JoinWith'] ?? \PhpursThunks::eval('Data_Semigroup_Foldable_JoinWith'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Function_const'] ?? \PhpursThunks::eval('Data_Function_const'))))($f_4))))($foldable_5))($j_3);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Semigroup_Foldable_intercalate'] = function() { $v = (function() {
  $__fn = function($dictFoldable1_0, $dictSemigroup_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Semigroup_Foldable_intercalate"), recVars=[];
  $__local_var_2_0 = ((($GLOBALS['Data_Semigroup_Foldable_intercalateMap'] ?? \PhpursThunks::eval('Data_Semigroup_Foldable_intercalateMap')))($dictFoldable1_0))($dictSemigroup_1);
  $__res = function($a_3) use ($__local_var_2_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($__local_var_2_0)($a_3))((($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))->identity);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Semigroup_Foldable_maximum'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Semigroup_Foldable_maximum"), recVars=[];
  $semigroupMax_1_0 = (object)["append" => (function() use ($dictOrd_0) {
  $__fn = function($v_1, $v1_2 = null) use ($dictOrd_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $v_3_0 = ((($dictOrd_0)->compare)($v_1))($v1_2);
  if ((is_object($v_3_0) && (($v_3_0)->tag === "LT"))) {
$__t1 = $v1_2;
} else {
if ((is_object($v_3_0) && (($v_3_0)->tag === "EQ"))) {
$__t1 = $v_1;
} else {
if ((is_object($v_3_0) && (($v_3_0)->tag === "GT"))) {
$__t1 = $v_1;
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
})()];
  $__res = function($dictFoldable1_2) use ($semigroupMax_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce')))(((($dictFoldable1_2)->foldMap1)($semigroupMax_1_0))(($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce'))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Semigroup_Foldable_minimum'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Semigroup_Foldable_minimum"), recVars=[];
  $semigroupMin_1_0 = (object)["append" => (function() use ($dictOrd_0) {
  $__fn = function($v_1, $v1_2 = null) use ($dictOrd_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $v_3_0 = ((($dictOrd_0)->compare)($v_1))($v1_2);
  if ((is_object($v_3_0) && (($v_3_0)->tag === "LT"))) {
$__t1 = $v_1;
} else {
if ((is_object($v_3_0) && (($v_3_0)->tag === "EQ"))) {
$__t1 = $v_1;
} else {
if ((is_object($v_3_0) && (($v_3_0)->tag === "GT"))) {
$__t1 = $v1_2;
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
})()];
  $__res = function($dictFoldable1_2) use ($semigroupMin_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce')))(((($dictFoldable1_2)->foldMap1)($semigroupMin_1_0))(($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce'))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Semigroup_Foldable_traverse1_'] = function() { $v = (function() {
  $__fn = function($dictFoldable1_0, $dictApply_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Semigroup_Foldable_traverse1_"), recVars=[];
  $__local_var_2_0 = (($dictApply_1)->Functor0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $foldMap12_3_1 = (($dictFoldable1_0)->foldMap1)((object)["append" => (function() use ($dictApply_1) {
  $__fn = function($v_3, $v1_4 = null) use ($dictApply_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (((($GLOBALS['Control_Apply_applySecond'] ?? \PhpursThunks::eval('Control_Apply_applySecond')))($dictApply_1))($v_3))($v1_4);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()]);
  $__res = (function() use ($__local_var_2_0, $foldMap12_3_1) {
  $__fn = function($f_4, $t_5 = null) use ($__local_var_2_0, $foldMap12_3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($__local_var_2_0)->map)(function($v_6) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($foldMap12_3_1)(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Semigroup_Foldable_Act'] ?? \PhpursThunks::eval('Data_Semigroup_Foldable_Act'))))($f_4)))($t_5));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Semigroup_Foldable_for1_'] = function() { $v = (function() {
  $__fn = function($dictFoldable1_0, $dictApply_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Semigroup_Foldable_for1_"), recVars=[];
  $__local_var_2_0 = ((($GLOBALS['Data_Semigroup_Foldable_traverse1_'] ?? \PhpursThunks::eval('Data_Semigroup_Foldable_traverse1_')))($dictFoldable1_0))($dictApply_1);
  $__res = (function() use ($__local_var_2_0) {
  $__fn = function($b_3, $a_4 = null) use ($__local_var_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($__local_var_2_0)($a_4))($b_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Semigroup_Foldable_sequence1_'] = function() { $v = (function() {
  $__fn = function($dictFoldable1_0, $dictApply_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Semigroup_Foldable_sequence1_"), recVars=[];
  $__res = (((($GLOBALS['Data_Semigroup_Foldable_traverse1_'] ?? \PhpursThunks::eval('Data_Semigroup_Foldable_traverse1_')))($dictFoldable1_0))($dictApply_1))((($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))->identity);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Semigroup_Foldable_fold1'] = function() { $v = (function() {
  $__fn = function($dictFoldable1_0, $dictSemigroup_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Semigroup_Foldable_fold1"), recVars=[];
  $__res = ((($dictFoldable1_0)->foldMap1)($dictSemigroup_1))((($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))->identity);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };






























