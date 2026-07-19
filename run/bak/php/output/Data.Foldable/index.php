<?php

namespace Data\Foldable;

// ALL IMPORTS: Control.Alt, Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Plus, Control.Semigroupoid, Data.Const, Data.Either, Data.Eq, Data.Foldable, Data.Function, Data.Functor.App, Data.Functor.Compose, Data.Functor.Coproduct, Data.Functor.Product, Data.HeytingAlgebra, Data.Identity, Data.Maybe, Data.Maybe.First, Data.Maybe.Last, Data.Monoid, Data.Monoid.Additive, Data.Monoid.Conj, Data.Monoid.Disj, Data.Monoid.Dual, Data.Monoid.Endo, Data.Monoid.Multiplicative, Data.Newtype, Data.Ord, Data.Ordering, Data.Semigroup, Data.Semiring, Data.Tuple, Data.Unit, Prelude, Prim
// TO REQUIRE: Control.Alt, Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Plus, Control.Semigroupoid, Data.Const, Data.Either, Data.Eq, Data.Foldable, Data.Function, Data.Functor.App, Data.Functor.Compose, Data.Functor.Coproduct, Data.Functor.Product, Data.HeytingAlgebra, Data.Identity, Data.Maybe, Data.Maybe.First, Data.Maybe.Last, Data.Monoid, Data.Monoid.Additive, Data.Monoid.Conj, Data.Monoid.Disj, Data.Monoid.Dual, Data.Monoid.Endo, Data.Monoid.Multiplicative, Data.Newtype, Data.Ord, Data.Ordering, Data.Semigroup, Data.Semiring, Data.Tuple, Data.Unit, Prelude
require_once __DIR__ . '/../Control.Alt/index.php';
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Category/index.php';
require_once __DIR__ . '/../Control.Plus/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Const/index.php';
require_once __DIR__ . '/../Data.Either/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Foldable/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor.App/index.php';
require_once __DIR__ . '/../Data.Functor.Compose/index.php';
require_once __DIR__ . '/../Data.Functor.Coproduct/index.php';
require_once __DIR__ . '/../Data.Functor.Product/index.php';
require_once __DIR__ . '/../Data.HeytingAlgebra/index.php';
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
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ordering/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Semiring/index.php';
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
\PhpursThunks::$thunks['Data_Foldable_monoidEndo'] = function() { $v = (($GLOBALS['Data_Monoid_Endo_monoidEndo'] ?? \PhpursThunks::eval('Data_Monoid_Endo_monoidEndo')))(($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn'))); return $v; };
\PhpursThunks::$thunks['Data_Foldable_Empty'] = function() { $v = ($GLOBALS['__phpurs_data0_Empty'] ??= new Phpurs_Data0("Empty")); return $v; };
\PhpursThunks::$thunks['Data_Foldable_Node'] = function() { $v = function($value0) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data1("Node", $value0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Foldable_Append'] = function() { $v = (function() {
  $__fn = function($value0, $value1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("Append", $value0, $value1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Foldable_semigroupFreeMonoidTree'] = function() { $v = (object)["append" => ($GLOBALS['Data_Foldable_Append'] ?? \PhpursThunks::eval('Data_Foldable_Append'))]; return $v; };
\PhpursThunks::$thunks['Data_Foldable_monoidFreeMonoidTree'] = function() { $v = (object)["mempty" => new Phpurs_Data0("Empty"), "Semigroup0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_Foldable_semigroupFreeMonoidTree'] ?? \PhpursThunks::eval('Data_Foldable_semigroupFreeMonoidTree'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Foldable_foldr'] = function() { $v = function($dict_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Foldable_foldr"), recVars=[];
  $__res = ($dict_0)->foldr;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Foldable_indexr'] = function() { $v = (function() {
  $__fn = function($dictFoldable_0, $idx_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Foldable_indexr"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(function($v_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($v_2)->elem;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($dictFoldable_0)->foldr)((function() use ($idx_1) {
  $__fn = function($a_2, $cursor_3 = null) use ($idx_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object(($cursor_3)->elem) && ((($cursor_3)->elem)->tag === "Just"))) {
$__t0 = $cursor_3;
} else {
if (((($GLOBALS['Data_Eq_eqIntImpl'] ?? \PhpursThunks::eval('Data_Eq_eqIntImpl')))(($cursor_3)->pos))($idx_1)) {
$__t0 = (object)["elem" => new Phpurs_Data1("Just", $a_2), "pos" => ($cursor_3)->pos];
} else {
$__t0 = (object)["pos" => ((($GLOBALS['Data_Semiring_intAdd'] ?? \PhpursThunks::eval('Data_Semiring_intAdd')))(($cursor_3)->pos))(1), "elem" => ($cursor_3)->elem];
};
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))((object)["elem" => new Phpurs_Data0("Nothing"), "pos" => 0]));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Foldable_null'] = function() { $v = function($dictFoldable_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Foldable_null"), recVars=[];
  $__res = ((($dictFoldable_0)->foldr)((function() {
  $__fn = function($v_1, $v1_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = false;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(true);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Foldable_oneOf'] = function() { $v = (function() {
  $__fn = function($dictFoldable_0, $dictPlus_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Foldable_oneOf"), recVars=[];
  $__res = ((($dictFoldable_0)->foldr)(((($dictPlus_1)->Alt0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->alt))(($dictPlus_1)->empty);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Foldable_oneOfMap'] = function() { $v = (function() {
  $__fn = function($dictFoldable_0, $dictPlus_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Foldable_oneOfMap"), recVars=[];
  $alt_2_0 = ((($dictPlus_1)->Alt0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->alt;
  $empty_3_1 = ($dictPlus_1)->empty;
  $__res = function($f_4) use ($alt_2_0, $dictFoldable_0, $empty_3_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($dictFoldable_0)->foldr)(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))($alt_2_0))($f_4)))($empty_3_1);
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
\PhpursThunks::$thunks['Data_Foldable_traverse_'] = function() { $v = function($dictApplicative_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Foldable_traverse_"), recVars=[];
  $applySecond_1_0 = (($GLOBALS['Control_Apply_applySecond'] ?? \PhpursThunks::eval('Control_Apply_applySecond')))((($dictApplicative_0)->Apply0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = (function() use ($applySecond_1_0, $dictApplicative_0) {
  $__fn = function($dictFoldable_2, $f_3 = null) use ($applySecond_1_0, $dictApplicative_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($dictFoldable_2)->foldr)(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))($applySecond_1_0))($f_3)))((($dictApplicative_0)->pure)(($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit'))));
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
\PhpursThunks::$thunks['Data_Foldable_for_'] = function() { $v = function($dictApplicative_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Foldable_for_"), recVars=[];
  $traverse_1_1_0 = (($GLOBALS['Data_Foldable_traverse_'] ?? \PhpursThunks::eval('Data_Foldable_traverse_')))($dictApplicative_0);
  $__res = function($dictFoldable_2) use ($traverse_1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__local_var_3_1 = ($traverse_1_1_0)($dictFoldable_2);
  $__res = (function() use ($__local_var_3_1) {
  $__fn = function($b_4, $a_5 = null) use ($__local_var_3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($__local_var_3_1)($a_5))($b_4);
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
}; return $v; };
\PhpursThunks::$thunks['Data_Foldable_sequence_'] = function() { $v = function($dictApplicative_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Foldable_sequence_"), recVars=[];
  $traverse_1_1_0 = (($GLOBALS['Data_Foldable_traverse_'] ?? \PhpursThunks::eval('Data_Foldable_traverse_')))($dictApplicative_0);
  $__res = function($dictFoldable_2) use ($traverse_1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($traverse_1_1_0)($dictFoldable_2))((($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))->identity);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Foldable_foldl'] = function() { $v = function($dict_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Foldable_foldl"), recVars=[];
  $__res = ($dict_0)->foldl;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Foldable_indexl'] = function() { $v = (function() {
  $__fn = function($dictFoldable_0, $idx_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Foldable_indexl"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(function($v_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($v_2)->elem;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($dictFoldable_0)->foldl)((function() use ($idx_1) {
  $__fn = function($cursor_2, $a_3 = null) use ($idx_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object(($cursor_2)->elem) && ((($cursor_2)->elem)->tag === "Just"))) {
$__t0 = $cursor_2;
} else {
if (((($GLOBALS['Data_Eq_eqIntImpl'] ?? \PhpursThunks::eval('Data_Eq_eqIntImpl')))(($cursor_2)->pos))($idx_1)) {
$__t0 = (object)["elem" => new Phpurs_Data1("Just", $a_3), "pos" => ($cursor_2)->pos];
} else {
$__t0 = (object)["pos" => ((($GLOBALS['Data_Semiring_intAdd'] ?? \PhpursThunks::eval('Data_Semiring_intAdd')))(($cursor_2)->pos))(1), "elem" => ($cursor_2)->elem];
};
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))((object)["elem" => new Phpurs_Data0("Nothing"), "pos" => 0]));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Foldable_intercalate'] = function() { $v = (function() {
  $__fn = function($dictFoldable_0, $dictMonoid_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Foldable_intercalate"), recVars=[];
  $__local_var_2_0 = (($dictMonoid_1)->Semigroup0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $mempty_3_1 = ($dictMonoid_1)->mempty;
  $__res = (function() use ($__local_var_2_0, $dictFoldable_0, $mempty_3_1) {
  $__fn = function($sep_4, $xs_5 = null) use ($__local_var_2_0, $dictFoldable_0, $mempty_3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((((($dictFoldable_0)->foldl)((function() use ($__local_var_2_0, $sep_4) {
  $__fn = function($v_6, $v1_7 = null) use ($__local_var_2_0, $sep_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if (($v_6)->init) {
$__t2 = (object)["init" => false, "acc" => $v1_7];
} else {
$__t2 = (object)["init" => false, "acc" => ((($__local_var_2_0)->append)(($v_6)->acc))(((($__local_var_2_0)->append)($sep_4))($v1_7))];
};
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))((object)["init" => true, "acc" => $mempty_3_1]))($xs_5))->acc;
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
\PhpursThunks::$thunks['Data_Foldable_length'] = function() { $v = (function() {
  $__fn = function($dictFoldable_0, $dictSemiring_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Foldable_length"), recVars=[];
  $one_2_0 = ($dictSemiring_1)->one;
  $__res = ((($dictFoldable_0)->foldl)((function() use ($dictSemiring_1, $one_2_0) {
  $__fn = function($c_3, $v_4 = null) use ($dictSemiring_1, $one_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($dictSemiring_1)->add)($one_2_0))($c_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(($dictSemiring_1)->zero);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Foldable_maximumBy'] = function() { $v = (function() {
  $__fn = function($dictFoldable_0, $cmp_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Foldable_maximumBy"), recVars=[];
  $__res = ((($dictFoldable_0)->foldl)((function() use ($cmp_1) {
  $__fn = function($v_2, $v1_3 = null) use ($cmp_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v_2) && (($v_2)->tag === "Nothing"))) {
$__t0 = new Phpurs_Data1("Just", $v1_3);
} else {
if ((is_object($v_2) && (($v_2)->tag === "Just"))) {
if ((is_object((($cmp_1)(($v_2)->value0))($v1_3)) && (((($cmp_1)(($v_2)->value0))($v1_3))->tag === "GT"))) {
$__t1 = ($v_2)->value0;
} else {
$__t1 = $v1_3;
};
$__t0 = new Phpurs_Data1("Just", $__t1);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(new Phpurs_Data0("Nothing"));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Foldable_maximum'] = function() { $v = (function() {
  $__fn = function($dictOrd_0, $dictFoldable_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Foldable_maximum"), recVars=[];
  $__res = ((($dictFoldable_1)->foldl)((function() use ($dictOrd_0) {
  $__fn = function($v_2, $v1_3 = null) use ($dictOrd_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v_2) && (($v_2)->tag === "Nothing"))) {
$__t0 = new Phpurs_Data1("Just", $v1_3);
} else {
if ((is_object($v_2) && (($v_2)->tag === "Just"))) {
if ((is_object(((($dictOrd_0)->compare)(($v_2)->value0))($v1_3)) && ((((($dictOrd_0)->compare)(($v_2)->value0))($v1_3))->tag === "GT"))) {
$__t1 = ($v_2)->value0;
} else {
$__t1 = $v1_3;
};
$__t0 = new Phpurs_Data1("Just", $__t1);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(new Phpurs_Data0("Nothing"));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Foldable_minimumBy'] = function() { $v = (function() {
  $__fn = function($dictFoldable_0, $cmp_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Foldable_minimumBy"), recVars=[];
  $__res = ((($dictFoldable_0)->foldl)((function() use ($cmp_1) {
  $__fn = function($v_2, $v1_3 = null) use ($cmp_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v_2) && (($v_2)->tag === "Nothing"))) {
$__t0 = new Phpurs_Data1("Just", $v1_3);
} else {
if ((is_object($v_2) && (($v_2)->tag === "Just"))) {
if ((is_object((($cmp_1)(($v_2)->value0))($v1_3)) && (((($cmp_1)(($v_2)->value0))($v1_3))->tag === "LT"))) {
$__t1 = ($v_2)->value0;
} else {
$__t1 = $v1_3;
};
$__t0 = new Phpurs_Data1("Just", $__t1);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(new Phpurs_Data0("Nothing"));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Foldable_minimum'] = function() { $v = (function() {
  $__fn = function($dictOrd_0, $dictFoldable_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Foldable_minimum"), recVars=[];
  $__res = ((($dictFoldable_1)->foldl)((function() use ($dictOrd_0) {
  $__fn = function($v_2, $v1_3 = null) use ($dictOrd_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v_2) && (($v_2)->tag === "Nothing"))) {
$__t0 = new Phpurs_Data1("Just", $v1_3);
} else {
if ((is_object($v_2) && (($v_2)->tag === "Just"))) {
if ((is_object(((($dictOrd_0)->compare)(($v_2)->value0))($v1_3)) && ((((($dictOrd_0)->compare)(($v_2)->value0))($v1_3))->tag === "LT"))) {
$__t1 = ($v_2)->value0;
} else {
$__t1 = $v1_3;
};
$__t0 = new Phpurs_Data1("Just", $__t1);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(new Phpurs_Data0("Nothing"));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Foldable_product'] = function() { $v = (function() {
  $__fn = function($dictFoldable_0, $dictSemiring_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Foldable_product"), recVars=[];
  $__res = ((($dictFoldable_0)->foldl)(($dictSemiring_1)->mul))(($dictSemiring_1)->one);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Foldable_sum'] = function() { $v = (function() {
  $__fn = function($dictFoldable_0, $dictSemiring_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Foldable_sum"), recVars=[];
  $__res = ((($dictFoldable_0)->foldl)(($dictSemiring_1)->add))(($dictSemiring_1)->zero);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Foldable_foldableTuple'] = function() { $v = (object)["foldr" => (function() {
  $__fn = function($f_0, $z_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($f_0)(($v_2)->value1))($z_1);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldl" => (function() {
  $__fn = function($f_0, $z_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($f_0)($z_1))(($v_2)->value1);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldMap" => (function() {
  $__fn = function($dictMonoid_0, $f_1 = null, $v_2 = null) use (&$__fn) {
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
})()]; return $v; };
\PhpursThunks::$thunks['Data_Foldable_foldableMultiplicative'] = function() { $v = (object)["foldr" => (function() {
  $__fn = function($f_0, $z_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($f_0)($v_2))($z_1);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldl" => (function() {
  $__fn = function($f_0, $z_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($f_0)($z_1))($v_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldMap" => (function() {
  $__fn = function($dictMonoid_0, $f_1 = null, $v_2 = null) use (&$__fn) {
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
})()]; return $v; };
\PhpursThunks::$thunks['Data_Foldable_foldableMaybe'] = function() { $v = (object)["foldr" => (function() {
  $__fn = function($v_0, $v1_1 = null, $v2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v2_2) && (($v2_2)->tag === "Nothing"))) {
$__t0 = $v1_1;
} else {
if ((is_object($v2_2) && (($v2_2)->tag === "Just"))) {
$__t0 = (($v_0)(($v2_2)->value0))($v1_1);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldl" => (function() {
  $__fn = function($v_0, $v1_1 = null, $v2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v2_2) && (($v2_2)->tag === "Nothing"))) {
$__t1 = $v1_1;
} else {
if ((is_object($v2_2) && (($v2_2)->tag === "Just"))) {
$__t1 = (($v_0)($v1_1))(($v2_2)->value0);
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
})(), "foldMap" => function($dictMonoid_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $mempty_1_2 = ($dictMonoid_0)->mempty;
  $__res = (function() use ($mempty_1_2) {
  $__fn = function($v_2, $v1_3 = null) use ($mempty_1_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v1_3) && (($v1_3)->tag === "Nothing"))) {
$__t3 = $mempty_1_2;
} else {
if ((is_object($v1_3) && (($v1_3)->tag === "Just"))) {
$__t3 = ($v_2)(($v1_3)->value0);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
};
};
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Foldable_foldableIdentity'] = function() { $v = (object)["foldr" => (function() {
  $__fn = function($f_0, $z_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($f_0)($v_2))($z_1);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldl" => (function() {
  $__fn = function($f_0, $z_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($f_0)($z_1))($v_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldMap" => (function() {
  $__fn = function($dictMonoid_0, $f_1 = null, $v_2 = null) use (&$__fn) {
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
})()]; return $v; };
\PhpursThunks::$thunks['Data_Foldable_foldableEither'] = function() { $v = (object)["foldr" => (function() {
  $__fn = function($v_0, $v1_1 = null, $v2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v2_2) && (($v2_2)->tag === "Left"))) {
$__t0 = $v1_1;
} else {
if ((is_object($v2_2) && (($v2_2)->tag === "Right"))) {
$__t0 = (($v_0)(($v2_2)->value0))($v1_1);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldl" => (function() {
  $__fn = function($v_0, $v1_1 = null, $v2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v2_2) && (($v2_2)->tag === "Left"))) {
$__t1 = $v1_1;
} else {
if ((is_object($v2_2) && (($v2_2)->tag === "Right"))) {
$__t1 = (($v_0)($v1_1))(($v2_2)->value0);
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
})(), "foldMap" => function($dictMonoid_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $mempty_1_2 = ($dictMonoid_0)->mempty;
  $__res = (function() use ($mempty_1_2) {
  $__fn = function($v_2, $v1_3 = null) use ($mempty_1_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v1_3) && (($v1_3)->tag === "Left"))) {
$__t3 = $mempty_1_2;
} else {
if ((is_object($v1_3) && (($v1_3)->tag === "Right"))) {
$__t3 = ($v_2)(($v1_3)->value0);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
};
};
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Foldable_foldableDual'] = function() { $v = (object)["foldr" => (function() {
  $__fn = function($f_0, $z_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($f_0)($v_2))($z_1);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldl" => (function() {
  $__fn = function($f_0, $z_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($f_0)($z_1))($v_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldMap" => (function() {
  $__fn = function($dictMonoid_0, $f_1 = null, $v_2 = null) use (&$__fn) {
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
})()]; return $v; };
\PhpursThunks::$thunks['Data_Foldable_foldableDisj'] = function() { $v = (object)["foldr" => (function() {
  $__fn = function($f_0, $z_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($f_0)($v_2))($z_1);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldl" => (function() {
  $__fn = function($f_0, $z_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($f_0)($z_1))($v_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldMap" => (function() {
  $__fn = function($dictMonoid_0, $f_1 = null, $v_2 = null) use (&$__fn) {
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
})()]; return $v; };
\PhpursThunks::$thunks['Data_Foldable_foldableConst'] = function() { $v = (object)["foldr" => (function() {
  $__fn = function($v_0, $z_1 = null, $v1_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $z_1;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldl" => (function() {
  $__fn = function($v_0, $z_1 = null, $v1_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $z_1;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldMap" => function($dictMonoid_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $mempty_1_0 = ($dictMonoid_0)->mempty;
  $__res = (function() use ($mempty_1_0) {
  $__fn = function($v_2, $v1_3 = null) use ($mempty_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $mempty_1_0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Foldable_foldableConj'] = function() { $v = (object)["foldr" => (function() {
  $__fn = function($f_0, $z_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($f_0)($v_2))($z_1);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldl" => (function() {
  $__fn = function($f_0, $z_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($f_0)($z_1))($v_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldMap" => (function() {
  $__fn = function($dictMonoid_0, $f_1 = null, $v_2 = null) use (&$__fn) {
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
})()]; return $v; };
\PhpursThunks::$thunks['Data_Foldable_foldableAdditive'] = function() { $v = (object)["foldr" => (function() {
  $__fn = function($f_0, $z_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($f_0)($v_2))($z_1);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldl" => (function() {
  $__fn = function($f_0, $z_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($f_0)($z_1))($v_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldMap" => (function() {
  $__fn = function($dictMonoid_0, $f_1 = null, $v_2 = null) use (&$__fn) {
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
})()]; return $v; };
\PhpursThunks::$thunks['Data_Foldable_foldMapDefaultR'] = function() { $v = (function() {
  $__fn = function($dictFoldable_0, $dictMonoid_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Foldable_foldMapDefaultR"), recVars=[];
  $mempty_2_0 = ($dictMonoid_1)->mempty;
  $__res = function($f_3) use ($dictFoldable_0, $dictMonoid_1, $mempty_2_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($dictFoldable_0)->foldr)((function() use ($dictMonoid_1, $f_3) {
  $__fn = function($x_4, $acc_5 = null) use ($dictMonoid_1, $f_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((((($dictMonoid_1)->Semigroup0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->append)(($f_3)($x_4)))($acc_5);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($mempty_2_0);
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
\PhpursThunks::$thunks['Data_Foldable_foldableArray'] = function() { $v = (object)["foldr" => ($GLOBALS['Data_Foldable_foldrArray'] ?? \PhpursThunks::eval('Data_Foldable_foldrArray')), "foldl" => ($GLOBALS['Data_Foldable_foldlArray'] ?? \PhpursThunks::eval('Data_Foldable_foldlArray')), "foldMap" => function($dictMonoid_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Foldable_foldableArray"];
  $__res = ((($GLOBALS['Data_Foldable_foldMapDefaultR'] ?? \PhpursThunks::eval('Data_Foldable_foldMapDefaultR')))(($GLOBALS['Data_Foldable_foldableArray'] ?? \PhpursThunks::eval('Data_Foldable_foldableArray'))))($dictMonoid_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Foldable_foldableFreeMonoidTree'] = function() { $v = (object)["foldl" => function($fn_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Foldable_foldableFreeMonoidTree"];
  $go_1_0 = null;
  $go_1_0 = (function() use ($fn_0, &$go_1_0) {
  $__fn = function($acc_2, $lhs_3 = null, $rhs_4 = null) use ($fn_0, &$go_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_1_0"), recVars=["Data_Foldable_foldableFreeMonoidTree","go_1_0"];
  while (true) {
if ((is_object($lhs_3) && (($lhs_3)->tag === "Node"))) {
$__tco_2 = (($fn_0)($acc_2))(($lhs_3)->value0);
$__tco_3 = $rhs_4;
$__tco_4 = new Phpurs_Data0("Empty");
$acc_2 = $__tco_2;
$lhs_3 = $__tco_3;
$rhs_4 = $__tco_4;
continue ;
$__t1 = null;
} else {
if ((is_object($lhs_3) && (($lhs_3)->tag === "Append"))) {
if ((is_object(($lhs_3)->value1) && ((($lhs_3)->value1)->tag === "Empty"))) {
$__tco_9 = $acc_2;
$__tco_10 = ($lhs_3)->value0;
$__tco_11 = $rhs_4;
$acc_2 = $__tco_9;
$lhs_3 = $__tco_10;
$rhs_4 = $__tco_11;
continue ;
$__t8 = null;
} else {
if ((is_object($rhs_4) && (($rhs_4)->tag === "Empty"))) {
$__tco_12 = $acc_2;
$__tco_13 = ($lhs_3)->value0;
$__tco_14 = ($lhs_3)->value1;
$acc_2 = $__tco_12;
$lhs_3 = $__tco_13;
$rhs_4 = $__tco_14;
continue ;
$__t8 = null;
} else {
$__tco_5 = $acc_2;
$__tco_6 = ($lhs_3)->value0;
$__tco_7 = new Phpurs_Data2("Append", ($lhs_3)->value1, $rhs_4);
$acc_2 = $__tco_5;
$lhs_3 = $__tco_6;
$rhs_4 = $__tco_7;
continue ;
$__t8 = null;
};
};
$__t1 = $__t8;
} else {
if ((is_object($lhs_3) && (($lhs_3)->tag === "Empty"))) {
if ((is_object($rhs_4) && (($rhs_4)->tag === "Empty"))) {
$__t18 = $acc_2;
} else {
$__tco_15 = $acc_2;
$__tco_16 = $rhs_4;
$__tco_17 = new Phpurs_Data0("Empty");
$acc_2 = $__tco_15;
$lhs_3 = $__tco_16;
$rhs_4 = $__tco_17;
continue ;
$__t18 = null;
};
$__t1 = $__t18;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
};
$__res = $__t1;
goto __end;;
};
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  $__res = (function() use (&$go_1_0) {
  $__fn = function($a_2, $b_3 = null) use (&$go_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Foldable_foldableFreeMonoidTree","go_1_0"];
  $__res = ((($go_1_0)($a_2))($b_3))(new Phpurs_Data0("Empty"));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldr" => function($fn_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Foldable_foldableFreeMonoidTree"];
  $go_1_19 = null;
  $go_1_19 = (function() use ($fn_0, &$go_1_19) {
  $__fn = function($acc_2, $lhs_3 = null, $rhs_4 = null) use ($fn_0, &$go_1_19, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_1_19"), recVars=["Data_Foldable_foldableFreeMonoidTree","go_1_19"];
  while (true) {
if ((is_object($rhs_4) && (($rhs_4)->tag === "Node"))) {
$__tco_21 = (($fn_0)(($rhs_4)->value0))($acc_2);
$__tco_22 = new Phpurs_Data0("Empty");
$__tco_23 = $lhs_3;
$acc_2 = $__tco_21;
$lhs_3 = $__tco_22;
$rhs_4 = $__tco_23;
continue ;
$__t20 = null;
} else {
if ((is_object($rhs_4) && (($rhs_4)->tag === "Append"))) {
if ((is_object(($rhs_4)->value0) && ((($rhs_4)->value0)->tag === "Empty"))) {
$__tco_28 = $acc_2;
$__tco_29 = $lhs_3;
$__tco_30 = ($rhs_4)->value1;
$acc_2 = $__tco_28;
$lhs_3 = $__tco_29;
$rhs_4 = $__tco_30;
continue ;
$__t27 = null;
} else {
if ((is_object($lhs_3) && (($lhs_3)->tag === "Empty"))) {
$__tco_31 = $acc_2;
$__tco_32 = ($rhs_4)->value0;
$__tco_33 = ($rhs_4)->value1;
$acc_2 = $__tco_31;
$lhs_3 = $__tco_32;
$rhs_4 = $__tco_33;
continue ;
$__t27 = null;
} else {
$__tco_24 = $acc_2;
$__tco_25 = new Phpurs_Data2("Append", $lhs_3, ($rhs_4)->value0);
$__tco_26 = ($rhs_4)->value1;
$acc_2 = $__tco_24;
$lhs_3 = $__tco_25;
$rhs_4 = $__tco_26;
continue ;
$__t27 = null;
};
};
$__t20 = $__t27;
} else {
if ((is_object($rhs_4) && (($rhs_4)->tag === "Empty"))) {
if ((is_object($lhs_3) && (($lhs_3)->tag === "Empty"))) {
$__t37 = $acc_2;
} else {
$__tco_34 = $acc_2;
$__tco_35 = new Phpurs_Data0("Empty");
$__tco_36 = $lhs_3;
$acc_2 = $__tco_34;
$lhs_3 = $__tco_35;
$rhs_4 = $__tco_36;
continue ;
$__t37 = null;
};
$__t20 = $__t37;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t20 = null;
};
};
};
$__res = $__t20;
goto __end;;
};
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  $__res = (function() use (&$go_1_19) {
  $__fn = function($a_2, $b_3 = null) use (&$go_1_19, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Foldable_foldableFreeMonoidTree","go_1_19"];
  $__res = ((($go_1_19)($a_2))(new Phpurs_Data0("Empty")))($b_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldMap" => function($dictMonoid_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Foldable_foldableFreeMonoidTree"];
  $__res = ((($GLOBALS['Data_Foldable_foldMapDefaultR'] ?? \PhpursThunks::eval('Data_Foldable_foldMapDefaultR')))(($GLOBALS['Data_Foldable_foldableFreeMonoidTree'] ?? \PhpursThunks::eval('Data_Foldable_foldableFreeMonoidTree'))))($dictMonoid_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Foldable_foldMapDefaultL'] = function() { $v = (function() {
  $__fn = function($dictFoldable_0, $dictMonoid_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Foldable_foldMapDefaultL"), recVars=[];
  $mempty_2_0 = ($dictMonoid_1)->mempty;
  $__res = function($f_3) use ($dictFoldable_0, $dictMonoid_1, $mempty_2_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($dictFoldable_0)->foldl)((function() use ($dictMonoid_1, $f_3) {
  $__fn = function($acc_4, $x_5 = null) use ($dictMonoid_1, $f_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((((($dictMonoid_1)->Semigroup0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->append)($acc_4))(($f_3)($x_5));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($mempty_2_0);
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
\PhpursThunks::$thunks['Data_Foldable_foldMap'] = function() { $v = function($dict_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Foldable_foldMap"), recVars=[];
  $__res = ($dict_0)->foldMap;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Foldable_foldableApp'] = function() { $v = function($dictFoldable_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Foldable_foldableApp"), recVars=[];
  $__res = (object)["foldr" => (function() use ($dictFoldable_0) {
  $__fn = function($f_1, $i_2 = null, $v_3 = null) use ($dictFoldable_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (((($dictFoldable_0)->foldr)($f_1))($i_2))($v_3);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldl" => (function() use ($dictFoldable_0) {
  $__fn = function($f_1, $i_2 = null, $v_3 = null) use ($dictFoldable_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (((($dictFoldable_0)->foldl)($f_1))($i_2))($v_3);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldMap" => function($dictMonoid_1) use ($dictFoldable_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($dictFoldable_0)->foldMap)($dictMonoid_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Foldable_foldableCompose'] = function() { $v = (function() {
  $__fn = function($dictFoldable_0, $dictFoldable1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Foldable_foldableCompose"), recVars=[];
  $__res = (object)["foldr" => (function() use ($dictFoldable1_1, $dictFoldable_0) {
  $__fn = function($f_2, $i_3 = null, $v_4 = null) use ($dictFoldable1_1, $dictFoldable_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__local_var_5_0 = (($dictFoldable1_1)->foldr)($f_2);
  $__res = (((($dictFoldable_0)->foldr)((function() use ($__local_var_5_0) {
  $__fn = function($b_6, $a_7 = null) use ($__local_var_5_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($__local_var_5_0)($a_7))($b_6);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($i_3))($v_4);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldl" => (function() use ($dictFoldable1_1, $dictFoldable_0) {
  $__fn = function($f_2, $i_3 = null, $v_4 = null) use ($dictFoldable1_1, $dictFoldable_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (((($dictFoldable_0)->foldl)((($dictFoldable1_1)->foldl)($f_2)))($i_3))($v_4);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldMap" => function($dictMonoid_2) use ($dictFoldable1_1, $dictFoldable_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $foldMap4_3_1 = (($dictFoldable_0)->foldMap)($dictMonoid_2);
  $foldMap5_4_2 = (($dictFoldable1_1)->foldMap)($dictMonoid_2);
  $__res = (function() use ($foldMap4_3_1, $foldMap5_4_2) {
  $__fn = function($f_5, $v_6 = null) use ($foldMap4_3_1, $foldMap5_4_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($foldMap4_3_1)(($foldMap5_4_2)($f_5)))($v_6);
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
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Foldable_foldableCoproduct'] = function() { $v = (function() {
  $__fn = function($dictFoldable_0, $dictFoldable1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Foldable_foldableCoproduct"), recVars=[];
  $__res = (object)["foldr" => (function() use ($dictFoldable1_1, $dictFoldable_0) {
  $__fn = function($f_2, $z_3 = null) use ($dictFoldable1_1, $dictFoldable_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__local_var_4_0 = ((($dictFoldable_0)->foldr)($f_2))($z_3);
  $__local_var_5_1 = ((($dictFoldable1_1)->foldr)($f_2))($z_3);
  $__res = function($v2_6) use ($__local_var_4_0, $__local_var_5_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v2_6) && (($v2_6)->tag === "Left"))) {
$__t2 = ($__local_var_4_0)(($v2_6)->value0);
} else {
if ((is_object($v2_6) && (($v2_6)->tag === "Right"))) {
$__t2 = ($__local_var_5_1)(($v2_6)->value0);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
};
};
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "foldl" => (function() use ($dictFoldable1_1, $dictFoldable_0) {
  $__fn = function($f_2, $z_3 = null) use ($dictFoldable1_1, $dictFoldable_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__local_var_4_3 = ((($dictFoldable_0)->foldl)($f_2))($z_3);
  $__local_var_5_4 = ((($dictFoldable1_1)->foldl)($f_2))($z_3);
  $__res = function($v2_6) use ($__local_var_4_3, $__local_var_5_4) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v2_6) && (($v2_6)->tag === "Left"))) {
$__t5 = ($__local_var_4_3)(($v2_6)->value0);
} else {
if ((is_object($v2_6) && (($v2_6)->tag === "Right"))) {
$__t5 = ($__local_var_5_4)(($v2_6)->value0);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t5 = null;
};
};
  $__res = $__t5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "foldMap" => function($dictMonoid_2) use ($dictFoldable1_1, $dictFoldable_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $foldMap4_3_6 = (($dictFoldable_0)->foldMap)($dictMonoid_2);
  $foldMap5_4_7 = (($dictFoldable1_1)->foldMap)($dictMonoid_2);
  $__res = function($f_5) use ($foldMap4_3_6, $foldMap5_4_7) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__local_var_6_8 = ($foldMap4_3_6)($f_5);
  $__local_var_7_9 = ($foldMap5_4_7)($f_5);
  $__res = function($v2_8) use ($__local_var_6_8, $__local_var_7_9) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v2_8) && (($v2_8)->tag === "Left"))) {
$__t10 = ($__local_var_6_8)(($v2_8)->value0);
} else {
if ((is_object($v2_8) && (($v2_8)->tag === "Right"))) {
$__t10 = ($__local_var_7_9)(($v2_8)->value0);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t10 = null;
};
};
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
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Foldable_foldableFirst'] = function() { $v = (object)["foldr" => (function() {
  $__fn = function($f_0, $z_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v_2) && (($v_2)->tag === "Nothing"))) {
$__t0 = $z_1;
} else {
if ((is_object($v_2) && (($v_2)->tag === "Just"))) {
$__t0 = (($f_0)(($v_2)->value0))($z_1);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldl" => (function() {
  $__fn = function($f_0, $z_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v_2) && (($v_2)->tag === "Nothing"))) {
$__t1 = $z_1;
} else {
if ((is_object($v_2) && (($v_2)->tag === "Just"))) {
$__t1 = (($f_0)($z_1))(($v_2)->value0);
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
})(), "foldMap" => function($dictMonoid_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $mempty_1_2 = ($dictMonoid_0)->mempty;
  $__res = (function() use ($mempty_1_2) {
  $__fn = function($v_2, $v1_3 = null) use ($mempty_1_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v1_3) && (($v1_3)->tag === "Nothing"))) {
$__t3 = $mempty_1_2;
} else {
if ((is_object($v1_3) && (($v1_3)->tag === "Just"))) {
$__t3 = ($v_2)(($v1_3)->value0);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
};
};
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Foldable_foldableLast'] = function() { $v = (object)["foldr" => (function() {
  $__fn = function($f_0, $z_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v_2) && (($v_2)->tag === "Nothing"))) {
$__t0 = $z_1;
} else {
if ((is_object($v_2) && (($v_2)->tag === "Just"))) {
$__t0 = (($f_0)(($v_2)->value0))($z_1);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldl" => (function() {
  $__fn = function($f_0, $z_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v_2) && (($v_2)->tag === "Nothing"))) {
$__t1 = $z_1;
} else {
if ((is_object($v_2) && (($v_2)->tag === "Just"))) {
$__t1 = (($f_0)($z_1))(($v_2)->value0);
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
})(), "foldMap" => function($dictMonoid_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $mempty_1_2 = ($dictMonoid_0)->mempty;
  $__res = (function() use ($mempty_1_2) {
  $__fn = function($v_2, $v1_3 = null) use ($mempty_1_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v1_3) && (($v1_3)->tag === "Nothing"))) {
$__t3 = $mempty_1_2;
} else {
if ((is_object($v1_3) && (($v1_3)->tag === "Just"))) {
$__t3 = ($v_2)(($v1_3)->value0);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
};
};
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Foldable_foldableProduct'] = function() { $v = (function() {
  $__fn = function($dictFoldable_0, $dictFoldable1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Foldable_foldableProduct"), recVars=[];
  $__res = (object)["foldr" => (function() use ($dictFoldable1_1, $dictFoldable_0) {
  $__fn = function($f_2, $z_3 = null, $v_4 = null) use ($dictFoldable1_1, $dictFoldable_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (((($dictFoldable_0)->foldr)($f_2))((((($dictFoldable1_1)->foldr)($f_2))($z_3))(($v_4)->value1)))(($v_4)->value0);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldl" => (function() use ($dictFoldable1_1, $dictFoldable_0) {
  $__fn = function($f_2, $z_3 = null, $v_4 = null) use ($dictFoldable1_1, $dictFoldable_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (((($dictFoldable1_1)->foldl)($f_2))((((($dictFoldable_0)->foldl)($f_2))($z_3))(($v_4)->value0)))(($v_4)->value1);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldMap" => function($dictMonoid_2) use ($dictFoldable1_1, $dictFoldable_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $foldMap4_3_0 = (($dictFoldable_0)->foldMap)($dictMonoid_2);
  $foldMap5_4_1 = (($dictFoldable1_1)->foldMap)($dictMonoid_2);
  $__res = (function() use ($dictMonoid_2, $foldMap4_3_0, $foldMap5_4_1) {
  $__fn = function($f_5, $v_6 = null) use ($dictMonoid_2, $foldMap4_3_0, $foldMap5_4_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((((($dictMonoid_2)->Semigroup0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->append)((($foldMap4_3_0)($f_5))(($v_6)->value0)))((($foldMap5_4_1)($f_5))(($v_6)->value1));
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
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Foldable_foldlDefault'] = function() { $v = function($dictFoldable_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Foldable_foldlDefault"), recVars=[];
  $foldMap2_1_0 = (($dictFoldable_0)->foldMap)(($GLOBALS['Data_Foldable_monoidFreeMonoidTree'] ?? \PhpursThunks::eval('Data_Foldable_monoidFreeMonoidTree')));
  $__res = (function() use ($foldMap2_1_0) {
  $__fn = function($c_2, $u_3 = null, $xs_4 = null) use ($foldMap2_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((((($GLOBALS['Data_Foldable_foldableFreeMonoidTree'] ?? \PhpursThunks::eval('Data_Foldable_foldableFreeMonoidTree')))->foldl)($c_2))($u_3))((($foldMap2_1_0)(($GLOBALS['Data_Foldable_Node'] ?? \PhpursThunks::eval('Data_Foldable_Node'))))($xs_4));
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
\PhpursThunks::$thunks['Data_Foldable_foldrDefault'] = function() { $v = function($dictFoldable_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Foldable_foldrDefault"), recVars=[];
  $foldMap2_1_0 = (($dictFoldable_0)->foldMap)(($GLOBALS['Data_Foldable_monoidFreeMonoidTree'] ?? \PhpursThunks::eval('Data_Foldable_monoidFreeMonoidTree')));
  $__res = (function() use ($foldMap2_1_0) {
  $__fn = function($c_2, $u_3 = null, $xs_4 = null) use ($foldMap2_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((((($GLOBALS['Data_Foldable_foldableFreeMonoidTree'] ?? \PhpursThunks::eval('Data_Foldable_foldableFreeMonoidTree')))->foldr)($c_2))($u_3))((($foldMap2_1_0)(($GLOBALS['Data_Foldable_Node'] ?? \PhpursThunks::eval('Data_Foldable_Node'))))($xs_4));
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
\PhpursThunks::$thunks['Data_Foldable_lookup'] = function() { $v = function($dictFoldable_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Foldable_lookup"), recVars=[];
  $foldMap2_1_0 = (($dictFoldable_0)->foldMap)(($GLOBALS['Data_Maybe_First_monoidFirst'] ?? \PhpursThunks::eval('Data_Maybe_First_monoidFirst')));
  $__res = (function() use ($foldMap2_1_0) {
  $__fn = function($dictEq_2, $a_3 = null) use ($foldMap2_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce'))))(($foldMap2_1_0)(function($v_4) use ($a_3, $dictEq_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if (((($dictEq_2)->eq)($a_3))(($v_4)->value0)) {
$__t1 = new Phpurs_Data1("Just", ($v_4)->value1);
} else {
$__t1 = new Phpurs_Data0("Nothing");
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
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
\PhpursThunks::$thunks['Data_Foldable_surroundMap'] = function() { $v = function($dictFoldable_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Foldable_surroundMap"), recVars=[];
  $foldMap2_1_0 = (($dictFoldable_0)->foldMap)(($GLOBALS['Data_Foldable_monoidEndo'] ?? \PhpursThunks::eval('Data_Foldable_monoidEndo')));
  $__res = (function() use ($foldMap2_1_0) {
  $__fn = function($dictSemigroup_2, $d_3 = null, $t_4 = null, $f_5 = null) use ($foldMap2_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce')))((($foldMap2_1_0)((function() use ($d_3, $dictSemigroup_2, $t_4) {
  $__fn = function($a_6, $m_7 = null) use ($d_3, $dictSemigroup_2, $t_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($dictSemigroup_2)->append)($d_3))(((($dictSemigroup_2)->append)(($t_4)($a_6)))($m_7));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($f_5)))($d_3);
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Foldable_surround'] = function() { $v = function($dictFoldable_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Foldable_surround"), recVars=[];
  $surroundMap1_1_0 = (($GLOBALS['Data_Foldable_surroundMap'] ?? \PhpursThunks::eval('Data_Foldable_surroundMap')))($dictFoldable_0);
  $__res = function($dictSemigroup_2) use ($surroundMap1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $surroundMap2_3_1 = ($surroundMap1_1_0)($dictSemigroup_2);
  $__res = function($d_4) use ($surroundMap2_3_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($surroundMap2_3_1)($d_4))((($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))->identity);
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
}; return $v; };
\PhpursThunks::$thunks['Data_Foldable_foldM'] = function() { $v = (function() {
  $__fn = function($dictFoldable_0, $dictMonad_1 = null, $f_2 = null, $b0_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Foldable_foldM"), recVars=[];
  $__res = ((($dictFoldable_0)->foldl)((function() use ($dictMonad_1, $f_2) {
  $__fn = function($b_4, $a_5 = null) use ($dictMonad_1, $f_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((((($dictMonad_1)->Bind1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->bind)($b_4))(function($a_6) use ($a_5, $f_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($f_2)($a_6))($a_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))((((($dictMonad_1)->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->pure)($b0_3));
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Foldable_fold'] = function() { $v = (function() {
  $__fn = function($dictFoldable_0, $dictMonoid_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Foldable_fold"), recVars=[];
  $__res = ((($dictFoldable_0)->foldMap)($dictMonoid_1))((($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))->identity);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Foldable_findMap'] = function() { $v = (function() {
  $__fn = function($dictFoldable_0, $p_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Foldable_findMap"), recVars=[];
  $__res = ((($dictFoldable_0)->foldl)((function() use ($p_1) {
  $__fn = function($v_2, $v1_3 = null) use ($p_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v_2) && (($v_2)->tag === "Nothing"))) {
$__t0 = ($p_1)($v1_3);
} else {
$__t0 = $v_2;
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(new Phpurs_Data0("Nothing"));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Foldable_find'] = function() { $v = (function() {
  $__fn = function($dictFoldable_0, $p_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Foldable_find"), recVars=[];
  $__res = ((($dictFoldable_0)->foldl)((function() use ($p_1) {
  $__fn = function($v_2, $v1_3 = null) use ($p_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if (((is_object($v_2) && (($v_2)->tag === "Nothing")) && ($p_1)($v1_3))) {
$__t0 = new Phpurs_Data1("Just", $v1_3);
} else {
$__t0 = $v_2;
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(new Phpurs_Data0("Nothing"));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Foldable_any'] = function() { $v = (function() {
  $__fn = function($dictFoldable_0, $dictHeytingAlgebra_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Foldable_any"), recVars=[];
  $semigroupDisj1_2_0 = (object)["append" => (function() use ($dictHeytingAlgebra_1) {
  $__fn = function($v_2, $v1_3 = null) use ($dictHeytingAlgebra_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($dictHeytingAlgebra_1)->disj)($v_2))($v1_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $__res = (($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce')))((($dictFoldable_0)->foldMap)((object)["mempty" => ($dictHeytingAlgebra_1)->ff, "Semigroup0" => function($dollar__unused_3) use ($semigroupDisj1_2_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $semigroupDisj1_2_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Foldable_elem'] = function() { $v = function($dictFoldable_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Foldable_elem"), recVars=[];
  $any1_1_0 = ((($GLOBALS['Data_Foldable_any'] ?? \PhpursThunks::eval('Data_Foldable_any')))($dictFoldable_0))(($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'] ?? \PhpursThunks::eval('Data_HeytingAlgebra_heytingAlgebraBoolean')));
  $__res = function($dictEq_2) use ($any1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))($any1_1_0))(($dictEq_2)->eq);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Foldable_notElem'] = function() { $v = function($dictFoldable_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Foldable_notElem"), recVars=[];
  $any1_1_0 = ((($GLOBALS['Data_Foldable_any'] ?? \PhpursThunks::eval('Data_Foldable_any')))($dictFoldable_0))(($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'] ?? \PhpursThunks::eval('Data_HeytingAlgebra_heytingAlgebraBoolean')));
  $__res = function($dictEq_2) use ($any1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $elem2_3_1 = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))($any1_1_0))(($dictEq_2)->eq);
  $__res = function($x_4) use ($elem2_3_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_HeytingAlgebra_boolNot'] ?? \PhpursThunks::eval('Data_HeytingAlgebra_boolNot'))))(($elem2_3_1)($x_4));
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
}; return $v; };
\PhpursThunks::$thunks['Data_Foldable_or'] = function() { $v = (function() {
  $__fn = function($dictFoldable_0, $dictHeytingAlgebra_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Foldable_or"), recVars=[];
  $__res = (((($GLOBALS['Data_Foldable_any'] ?? \PhpursThunks::eval('Data_Foldable_any')))($dictFoldable_0))($dictHeytingAlgebra_1))((($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))->identity);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Foldable_all'] = function() { $v = (function() {
  $__fn = function($dictFoldable_0, $dictHeytingAlgebra_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Foldable_all"), recVars=[];
  $semigroupConj1_2_0 = (object)["append" => (function() use ($dictHeytingAlgebra_1) {
  $__fn = function($v_2, $v1_3 = null) use ($dictHeytingAlgebra_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($dictHeytingAlgebra_1)->conj)($v_2))($v1_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $__res = (($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce')))((($dictFoldable_0)->foldMap)((object)["mempty" => ($dictHeytingAlgebra_1)->tt, "Semigroup0" => function($dollar__unused_3) use ($semigroupConj1_2_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $semigroupConj1_2_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Foldable_and'] = function() { $v = (function() {
  $__fn = function($dictFoldable_0, $dictHeytingAlgebra_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Foldable_and"), recVars=[];
  $__res = (((($GLOBALS['Data_Foldable_all'] ?? \PhpursThunks::eval('Data_Foldable_all')))($dictFoldable_0))($dictHeytingAlgebra_1))((($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))->identity);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };
$ffi_Data_Foldable = \call_user_func(function() {
  $exports = [];
$foldrArray = function($f, $init = null, $xs = null) use (&$foldrArray) {
    if (\func_num_args() < 3) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$foldrArray) {
            return $foldrArray(...\array_merge($__args, $more));
        };
    }
    
    $acc = $init;
    for ($i = \count($xs) - 1; $i >= 0; $i--) {
        $f1 = $f($xs[$i]);
        $acc = $f1($acc);
    }
    return $acc;
};
$exports['foldrArray'] = $foldrArray;

$foldlArray = function($f, $init = null, $xs = null) use (&$foldlArray) {
    if (\func_num_args() < 3) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$foldlArray) {
            return $foldlArray(...\array_merge($__args, $more));
        };
    }
    
    $acc = $init;
    for ($i = 0, $len = \count($xs); $i < $len; $i++) {
        $f1 = $f($acc);
        $acc = $f1($xs[$i]);
    }
    return $acc;
};
$exports['foldlArray'] = $foldlArray;

return $exports;
  return $exports;
});
\PhpursThunks::$thunks['Data_Foldable_foldlArray'] = function() use (&$ffi_Data_Foldable) { return $ffi_Data_Foldable['foldlArray']; };
\PhpursThunks::$thunks['Data_Foldable_foldrArray'] = function() use (&$ffi_Data_Foldable) { return $ffi_Data_Foldable['foldrArray']; };






























































