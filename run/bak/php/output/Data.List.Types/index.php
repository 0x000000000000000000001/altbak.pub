<?php

namespace Data\List\Types;

// ALL IMPORTS: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Comonad, Control.Extend, Control.Monad, Control.MonadPlus, Control.Plus, Control.Semigroupoid, Data.Eq, Data.Foldable, Data.FoldableWithIndex, Data.Function, Data.Functor, Data.FunctorWithIndex, Data.HeytingAlgebra, Data.List.Types, Data.Maybe, Data.Monoid, Data.Newtype, Data.NonEmpty, Data.Ord, Data.Ordering, Data.Ring, Data.Semigroup, Data.Semigroup.Foldable, Data.Semigroup.Traversable, Data.Semiring, Data.Show, Data.Traversable, Data.TraversableWithIndex, Data.Tuple, Data.Unfoldable, Data.Unfoldable1, Prelude, Prim
// TO REQUIRE: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Comonad, Control.Extend, Control.Monad, Control.MonadPlus, Control.Plus, Control.Semigroupoid, Data.Eq, Data.Foldable, Data.FoldableWithIndex, Data.Function, Data.Functor, Data.FunctorWithIndex, Data.HeytingAlgebra, Data.List.Types, Data.Maybe, Data.Monoid, Data.Newtype, Data.NonEmpty, Data.Ord, Data.Ordering, Data.Ring, Data.Semigroup, Data.Semigroup.Foldable, Data.Semigroup.Traversable, Data.Semiring, Data.Show, Data.Traversable, Data.TraversableWithIndex, Data.Tuple, Data.Unfoldable, Data.Unfoldable1, Prelude
require_once __DIR__ . '/../Control.Alt/index.php';
require_once __DIR__ . '/../Control.Alternative/index.php';
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Category/index.php';
require_once __DIR__ . '/../Control.Comonad/index.php';
require_once __DIR__ . '/../Control.Extend/index.php';
require_once __DIR__ . '/../Control.Monad/index.php';
require_once __DIR__ . '/../Control.MonadPlus/index.php';
require_once __DIR__ . '/../Control.Plus/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Foldable/index.php';
require_once __DIR__ . '/../Data.FoldableWithIndex/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.FunctorWithIndex/index.php';
require_once __DIR__ . '/../Data.HeytingAlgebra/index.php';
require_once __DIR__ . '/../Data.List.Types/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Monoid/index.php';
require_once __DIR__ . '/../Data.Newtype/index.php';
require_once __DIR__ . '/../Data.NonEmpty/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ordering/index.php';
require_once __DIR__ . '/../Data.Ring/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Semigroup.Foldable/index.php';
require_once __DIR__ . '/../Data.Semigroup.Traversable/index.php';
require_once __DIR__ . '/../Data.Semiring/index.php';
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
\PhpursThunks::$thunks['Data_List_Types_Nil'] = function() { $v = ($GLOBALS['__phpurs_data0_Nil'] ??= new Phpurs_Data0("Nil")); return $v; };
\PhpursThunks::$thunks['Data_List_Types_Cons'] = function() { $v = (function() {
  $__fn = function($value0, $value1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("Cons", $value0, $value1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_Types_NonEmptyList'] = function() { $v = function($x_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Types_NonEmptyList"), recVars=[];
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_Types_toList'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Types_toList"), recVars=[];
  $__res = new Phpurs_Data2("Cons", ($v_0)->value0, ($v_0)->value1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_Types_newtypeNonEmptyList'] = function() { $v = (object)["Coercible0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Types_nelCons'] = function() { $v = (function() {
  $__fn = function($a_0, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Types_nelCons"), recVars=[];
  $__res = new Phpurs_Data2("NonEmpty", $a_0, new Phpurs_Data2("Cons", ($v_1)->value0, ($v_1)->value1));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_Types_listMap'] = function() { $v = function($f_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Types_listMap"), recVars=[];
  $chunkedRevMap_1_0 = null;
  $chunkedRevMap_1_0 = (function() use (&$chunkedRevMap_1_0, $f_0) {
  $__fn = function($v_2, $v1_3 = null) use (&$chunkedRevMap_1_0, $f_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "chunkedRevMap_1_0"), recVars=["chunkedRevMap_1_0"];
  while (true) {
if (((is_object($v1_3) && (($v1_3)->tag === "Cons")) && ((is_object(($v1_3)->value1) && ((($v1_3)->value1)->tag === "Cons")) && (is_object((($v1_3)->value1)->value1) && (((($v1_3)->value1)->value1)->tag === "Cons"))))) {
$__tco_9 = new Phpurs_Data2("Cons", $v1_3, $v_2);
$__tco_10 = ((($v1_3)->value1)->value1)->value1;
$v_2 = $__tco_9;
$v1_3 = $__tco_10;
continue ;
$__t8 = null;
} else {
$reverseUnrolledMap_4_1 = null;
$reverseUnrolledMap_4_1 = (function() use ($f_0, &$reverseUnrolledMap_4_1) {
  $__fn = function($v2_5, $v3_6 = null) use ($f_0, &$reverseUnrolledMap_4_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "reverseUnrolledMap_4_1"), recVars=["chunkedRevMap_1_0","reverseUnrolledMap_4_1"];
  while (true) {
if (((is_object($v2_5) && (($v2_5)->tag === "Cons")) && ((is_object(($v2_5)->value0) && ((($v2_5)->value0)->tag === "Cons")) && ((is_object((($v2_5)->value0)->value1) && (((($v2_5)->value0)->value1)->tag === "Cons")) && (is_object(((($v2_5)->value0)->value1)->value1) && ((((($v2_5)->value0)->value1)->value1)->tag === "Cons")))))) {
$__tco_3 = ($v2_5)->value1;
$__tco_4 = new Phpurs_Data2("Cons", ($f_0)((($v2_5)->value0)->value0), new Phpurs_Data2("Cons", ($f_0)(((($v2_5)->value0)->value1)->value0), new Phpurs_Data2("Cons", ($f_0)((((($v2_5)->value0)->value1)->value1)->value0), $v3_6)));
$v2_5 = $__tco_3;
$v3_6 = $__tco_4;
continue ;
$__t2 = null;
} else {
$__t2 = $v3_6;
};
$__res = $__t2;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
if ((is_object($v1_3) && (($v1_3)->tag === "Cons"))) {
if ((is_object(($v1_3)->value1) && ((($v1_3)->value1)->tag === "Cons"))) {
if ((is_object((($v1_3)->value1)->value1) && (((($v1_3)->value1)->value1)->tag === "Nil"))) {
$__t7 = new Phpurs_Data2("Cons", ($f_0)(($v1_3)->value0), new Phpurs_Data2("Cons", ($f_0)((($v1_3)->value1)->value0), new Phpurs_Data0("Nil")));
} else {
$__t7 = new Phpurs_Data0("Nil");
};
$__t6 = $__t7;
} else {
if ((is_object(($v1_3)->value1) && ((($v1_3)->value1)->tag === "Nil"))) {
$__t6 = new Phpurs_Data2("Cons", ($f_0)(($v1_3)->value0), new Phpurs_Data0("Nil"));
} else {
$__t6 = new Phpurs_Data0("Nil");
};
};
$__t5 = $__t6;
} else {
$__t5 = new Phpurs_Data0("Nil");
};
$__t8 = (($reverseUnrolledMap_4_1)($v_2))($__t5);
};
$__res = $__t8;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ($chunkedRevMap_1_0)(new Phpurs_Data0("Nil"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_Types_functorList'] = function() { $v = (object)["map" => ($GLOBALS['Data_List_Types_listMap'] ?? \PhpursThunks::eval('Data_List_Types_listMap'))]; return $v; };
\PhpursThunks::$thunks['Data_List_Types_functorNonEmptyList'] = function() { $v = (object)["map" => (function() {
  $__fn = function($f_0, $m_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = new Phpurs_Data2("NonEmpty", ($f_0)(($m_1)->value0), ((($GLOBALS['Data_List_Types_listMap'] ?? \PhpursThunks::eval('Data_List_Types_listMap')))($f_0))(($m_1)->value1));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Data_List_Types_foldableList'] = function() { $v = (object)["foldr" => (function() {
  $__fn = function($f_0, $b_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Types_foldableList"];
  $go_2_0 = null;
  $go_2_0 = (function() use (&$go_2_0) {
  $__fn = function($v_3, $v1_4 = null) use (&$go_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_2_0"), recVars=["Data_List_Types_foldableList","go_2_0"];
  while (true) {
if ((is_object($v1_4) && (($v1_4)->tag === "Nil"))) {
$__t1 = $v_3;
} else {
if ((is_object($v1_4) && (($v1_4)->tag === "Cons"))) {
$__tco_2 = new Phpurs_Data2("Cons", ($v1_4)->value0, $v_3);
$__tco_3 = ($v1_4)->value1;
$v_3 = $__tco_2;
$v1_4 = $__tco_3;
continue ;
$__t1 = null;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
$__res = $__t1;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((((($GLOBALS['Data_List_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Types_foldableList')))->foldl)((function() use ($f_0) {
  $__fn = function($b_2, $a_3 = null) use ($f_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Types_foldableList"];
  $__res = (($f_0)($a_3))($b_2);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($b_1)))(($go_2_0)(new Phpurs_Data0("Nil")));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "foldl" => function($f_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Types_foldableList"];
  $go_1_4 = null;
  $go_1_4 = (function() use ($f_0, &$go_1_4) {
  $__fn = function($b_2, $v_3 = null) use ($f_0, &$go_1_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_1_4"), recVars=["Data_List_Types_foldableList","go_1_4"];
  while (true) {
if ((is_object($v_3) && (($v_3)->tag === "Nil"))) {
$__t5 = $b_2;
} else {
if ((is_object($v_3) && (($v_3)->tag === "Cons"))) {
$__tco_6 = (($f_0)($b_2))(($v_3)->value0);
$__tco_7 = ($v_3)->value1;
$b_2 = $__tco_6;
$v_3 = $__tco_7;
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
  $__res = $go_1_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldMap" => function($dictMonoid_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Types_foldableList"];
  $mempty_1_8 = ($dictMonoid_0)->mempty;
  $__res = function($f_2) use ($dictMonoid_0, $mempty_1_8) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Types_foldableList"];
  $__res = (((($GLOBALS['Data_List_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Types_foldableList')))->foldl)(function($acc_3) use ($dictMonoid_0, $f_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Types_foldableList"];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((((($dictMonoid_0)->Semigroup0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->append)($acc_3)))($f_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($mempty_1_8);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Types_intercalate'] = function() { $v = ((($GLOBALS['Data_Foldable_intercalate'] ?? \PhpursThunks::eval('Data_Foldable_intercalate')))(($GLOBALS['Data_List_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Types_foldableList'))))(($GLOBALS['Data_Monoid_monoidString'] ?? \PhpursThunks::eval('Data_Monoid_monoidString'))); return $v; };
\PhpursThunks::$thunks['Data_List_Types_foldableNonEmptyList'] = function() { $v = (($GLOBALS['Data_NonEmpty_foldableNonEmpty'] ?? \PhpursThunks::eval('Data_NonEmpty_foldableNonEmpty')))(($GLOBALS['Data_List_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Types_foldableList'))); return $v; };
\PhpursThunks::$thunks['Data_List_Types_foldableWithIndexList'] = function() { $v = (object)["foldrWithIndex" => (function() {
  $__fn = function($f_0, $b_1 = null, $xs_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Types_foldableWithIndexList"];
  $go_3_0 = null;
  $go_3_0 = (function() use (&$go_3_0) {
  $__fn = function($b_4, $v_5 = null) use (&$go_3_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_3_0"), recVars=["Data_List_Types_foldableWithIndexList","go_3_0"];
  while (true) {
if ((is_object($v_5) && (($v_5)->tag === "Nil"))) {
$__t1 = $b_4;
} else {
if ((is_object($v_5) && (($v_5)->tag === "Cons"))) {
$__tco_2 = new Phpurs_Data2("Tuple", (($b_4)->value0 + 1), new Phpurs_Data2("Cons", ($v_5)->value0, ($b_4)->value1));
$__tco_3 = ($v_5)->value1;
$b_4 = $__tco_2;
$v_5 = $__tco_3;
continue ;
$__t1 = null;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
$__res = $__t1;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $v_4_4 = (($go_3_0)(new Phpurs_Data2("Tuple", 0, new Phpurs_Data0("Nil"))))($xs_2);
  $go_5_5 = null;
  $go_5_5 = (function() use ($f_0, &$go_5_5) {
  $__fn = function($b_6, $v_7 = null) use ($f_0, &$go_5_5, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_5_5"), recVars=["Data_List_Types_foldableWithIndexList","go_3_0","go_5_5"];
  while (true) {
if ((is_object($v_7) && (($v_7)->tag === "Nil"))) {
$__t6 = $b_6;
} else {
if ((is_object($v_7) && (($v_7)->tag === "Cons"))) {
$__tco_7 = new Phpurs_Data2("Tuple", (($b_6)->value0 - 1), ((($f_0)((($b_6)->value0 - 1)))(($v_7)->value0))(($b_6)->value1));
$__tco_8 = ($v_7)->value1;
$b_6 = $__tco_7;
$v_7 = $__tco_8;
continue ;
$__t6 = null;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t6 = null;
};
};
$__res = $__t6;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ((($go_5_5)(new Phpurs_Data2("Tuple", ($v_4_4)->value0, $b_1)))(($v_4_4)->value1))->value1;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldlWithIndex" => (function() {
  $__fn = function($f_0, $acc_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Types_foldableWithIndexList"];
  $go_2_9 = null;
  $go_2_9 = (function() use ($f_0, &$go_2_9) {
  $__fn = function($b_3, $v_4 = null) use ($f_0, &$go_2_9, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_2_9"), recVars=["Data_List_Types_foldableWithIndexList","go_2_9"];
  while (true) {
if ((is_object($v_4) && (($v_4)->tag === "Nil"))) {
$__t10 = $b_3;
} else {
if ((is_object($v_4) && (($v_4)->tag === "Cons"))) {
$__tco_11 = new Phpurs_Data2("Tuple", (($b_3)->value0 + 1), ((($f_0)(($b_3)->value0))(($b_3)->value1))(($v_4)->value0));
$__tco_12 = ($v_4)->value1;
$b_3 = $__tco_11;
$v_4 = $__tco_12;
continue ;
$__t10 = null;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t10 = null;
};
};
$__res = $__t10;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Tuple_snd'] ?? \PhpursThunks::eval('Data_Tuple_snd'))))(($go_2_9)(new Phpurs_Data2("Tuple", 0, $acc_1)));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "foldMapWithIndex" => function($dictMonoid_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Types_foldableWithIndexList"];
  $mempty_1_13 = ($dictMonoid_0)->mempty;
  $__res = function($f_2) use ($dictMonoid_0, $mempty_1_13) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Types_foldableWithIndexList"];
  $__res = (((($GLOBALS['Data_List_Types_foldableWithIndexList'] ?? \PhpursThunks::eval('Data_List_Types_foldableWithIndexList')))->foldlWithIndex)((function() use ($dictMonoid_0, $f_2) {
  $__fn = function($i_3, $acc_4 = null) use ($dictMonoid_0, $f_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Types_foldableWithIndexList"];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((((($dictMonoid_0)->Semigroup0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->append)($acc_4)))(($f_2)($i_3));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($mempty_1_13);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Types_foldableWithIndexList"];
  $__res = ($GLOBALS['Data_List_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Types_foldableList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Types_foldableWithIndexNonEmpty'] = function() { $v = (($GLOBALS['Data_NonEmpty_foldableWithIndexNonEmpty'] ?? \PhpursThunks::eval('Data_NonEmpty_foldableWithIndexNonEmpty')))(($GLOBALS['Data_List_Types_foldableWithIndexList'] ?? \PhpursThunks::eval('Data_List_Types_foldableWithIndexList'))); return $v; };
\PhpursThunks::$thunks['Data_List_Types_foldableWithIndexNonEmptyList'] = function() { $v = (object)["foldMapWithIndex" => function($dictMonoid_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $foldMapWithIndex1_1_0 = ((($GLOBALS['Data_List_Types_foldableWithIndexNonEmpty'] ?? \PhpursThunks::eval('Data_List_Types_foldableWithIndexNonEmpty')))->foldMapWithIndex)($dictMonoid_0);
  $__res = (function() use ($foldMapWithIndex1_1_0) {
  $__fn = function($f_2, $v_3 = null) use ($foldMapWithIndex1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($foldMapWithIndex1_1_0)(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))($f_2))(function($v2_4) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v2_4) && (($v2_4)->tag === "Nothing"))) {
$__t1 = 0;
} else {
if ((is_object($v2_4) && (($v2_4)->tag === "Just"))) {
$__t1 = (1 + ($v2_4)->value0);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))($v_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldlWithIndex" => (function() {
  $__fn = function($f_0, $b_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((((($GLOBALS['Data_List_Types_foldableWithIndexNonEmpty'] ?? \PhpursThunks::eval('Data_List_Types_foldableWithIndexNonEmpty')))->foldlWithIndex)(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))($f_0))(function($v2_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v2_3) && (($v2_3)->tag === "Nothing"))) {
$__t2 = 0;
} else {
if ((is_object($v2_3) && (($v2_3)->tag === "Just"))) {
$__t2 = (1 + ($v2_3)->value0);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
};
};
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))($b_1))($v_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldrWithIndex" => (function() {
  $__fn = function($f_0, $b_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((((($GLOBALS['Data_List_Types_foldableWithIndexNonEmpty'] ?? \PhpursThunks::eval('Data_List_Types_foldableWithIndexNonEmpty')))->foldrWithIndex)(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))($f_0))(function($v2_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v2_3) && (($v2_3)->tag === "Nothing"))) {
$__t3 = 0;
} else {
if ((is_object($v2_3) && (($v2_3)->tag === "Just"))) {
$__t3 = (1 + ($v2_3)->value0);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
};
};
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))($b_1))($v_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "Foldable0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_List_Types_foldableNonEmptyList'] ?? \PhpursThunks::eval('Data_List_Types_foldableNonEmptyList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Types_functorWithIndexList'] = function() { $v = (object)["mapWithIndex" => function($f_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (((($GLOBALS['Data_List_Types_foldableWithIndexList'] ?? \PhpursThunks::eval('Data_List_Types_foldableWithIndexList')))->foldrWithIndex)((function() use ($f_0) {
  $__fn = function($i_1, $x_2 = null, $acc_3 = null) use ($f_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = new Phpurs_Data2("Cons", (($f_0)($i_1))($x_2), $acc_3);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()))(new Phpurs_Data0("Nil"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_List_Types_functorList'] ?? \PhpursThunks::eval('Data_List_Types_functorList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Types_mapWithIndex'] = function() { $v = ((($GLOBALS['Data_NonEmpty_functorWithIndex'] ?? \PhpursThunks::eval('Data_NonEmpty_functorWithIndex')))(($GLOBALS['Data_List_Types_functorWithIndexList'] ?? \PhpursThunks::eval('Data_List_Types_functorWithIndexList'))))->mapWithIndex; return $v; };
\PhpursThunks::$thunks['Data_List_Types_functorWithIndexNonEmptyList'] = function() { $v = (object)["mapWithIndex" => (function() {
  $__fn = function($fn_0, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Data_List_Types_mapWithIndex'] ?? \PhpursThunks::eval('Data_List_Types_mapWithIndex')))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))($fn_0))(function($v2_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v2_2) && (($v2_2)->tag === "Nothing"))) {
$__t0 = 0;
} else {
if ((is_object($v2_2) && (($v2_2)->tag === "Just"))) {
$__t0 = (1 + ($v2_2)->value0);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))($v_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_List_Types_functorNonEmptyList'] ?? \PhpursThunks::eval('Data_List_Types_functorNonEmptyList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Types_semigroupList'] = function() { $v = (object)["append" => (function() {
  $__fn = function($xs_0, $ys_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((((($GLOBALS['Data_List_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Types_foldableList')))->foldr)(($GLOBALS['Data_List_Types_Cons'] ?? \PhpursThunks::eval('Data_List_Types_Cons'))))($ys_1))($xs_0);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Data_List_Types_monoidList'] = function() { $v = (object)["mempty" => new Phpurs_Data0("Nil"), "Semigroup0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_List_Types_semigroupList'] ?? \PhpursThunks::eval('Data_List_Types_semigroupList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Types_semigroupNonEmptyList'] = function() { $v = (object)["append" => (function() {
  $__fn = function($v_0, $as__prime___1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = new Phpurs_Data2("NonEmpty", ($v_0)->value0, ((((($GLOBALS['Data_List_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Types_foldableList')))->foldr)(($GLOBALS['Data_List_Types_Cons'] ?? \PhpursThunks::eval('Data_List_Types_Cons'))))(new Phpurs_Data2("Cons", ($as__prime___1)->value0, ($as__prime___1)->value1)))(($v_0)->value1));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Data_List_Types_showList'] = function() { $v = function($dictShow_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Types_showList"), recVars=[];
  $show_1_0 = ($dictShow_0)->show;
  $__res = (object)["show" => function($v_2) use ($show_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v_2) && (($v_2)->tag === "Nil"))) {
$__t1 = "Nil";
} else {
$__t1 = (("(" . ((($GLOBALS['Data_List_Types_intercalate'] ?? \PhpursThunks::eval('Data_List_Types_intercalate')))(" : "))(((($GLOBALS['Data_List_Types_listMap'] ?? \PhpursThunks::eval('Data_List_Types_listMap')))($show_1_0))($v_2))) . " : Nil)");
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_Types_showNonEmptyList'] = function() { $v = function($dictShow_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Types_showNonEmptyList"), recVars=[];
  $__local_var_1_0 = (($GLOBALS['Data_List_Types_showList'] ?? \PhpursThunks::eval('Data_List_Types_showList')))($dictShow_0);
  $__res = (object)["show" => function($v_2) use ($__local_var_1_0, $dictShow_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (((("(NonEmptyList (NonEmpty " . (($dictShow_0)->show)(($v_2)->value0)) . " ") . (($__local_var_1_0)->show)(($v_2)->value1)) . "))");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_Types_traversableList'] = function() { $v = (object)["traverse" => function($dictApplicative_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Types_traversableList"];
  $Apply0_1_0 = (($dictApplicative_0)->Apply0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $__res = function($f_2) use ($Apply0_1_0, $dictApplicative_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Types_traversableList"];
  $go_3_1 = null;
  $go_3_1 = (function() use (&$go_3_1) {
  $__fn = function($b_4, $v_5 = null) use (&$go_3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_3_1"), recVars=["Data_List_Types_traversableList","go_3_1"];
  while (true) {
if ((is_object($v_5) && (($v_5)->tag === "Nil"))) {
$__t2 = $b_4;
} else {
if ((is_object($v_5) && (($v_5)->tag === "Cons"))) {
$__tco_3 = new Phpurs_Data2("Cons", ($v_5)->value0, $b_4);
$__tco_4 = ($v_5)->value1;
$b_4 = $__tco_3;
$v_5 = $__tco_4;
continue ;
$__t2 = null;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
};
};
$__res = $__t2;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $go_3_5 = null;
  $go_3_5 = (function() use ($Apply0_1_0, $f_2, &$go_3_5) {
  $__fn = function($b_4, $v_5 = null) use ($Apply0_1_0, $f_2, &$go_3_5, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_3_5"), recVars=["Data_List_Types_traversableList","go_3_5"];
  while (true) {
if ((is_object($v_5) && (($v_5)->tag === "Nil"))) {
$__t6 = $b_4;
} else {
if ((is_object($v_5) && (($v_5)->tag === "Cons"))) {
$__tco_7 = (((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((((($GLOBALS['Control_Apply_lift2'] ?? \PhpursThunks::eval('Control_Apply_lift2')))($Apply0_1_0))((function() {
  $__fn = function($b_6, $a_7 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Types_traversableList","go_3_5"];
  $__res = new Phpurs_Data2("Cons", $a_7, $b_6);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($b_4)))($f_2))(($v_5)->value0);
$__tco_8 = ($v_5)->value1;
$b_4 = $__tco_7;
$v_5 = $__tco_8;
continue ;
$__t6 = null;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t6 = null;
};
};
$__res = $__t6;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((((($Apply0_1_0)->Functor0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->map)(($go_3_1)(new Phpurs_Data0("Nil")))))(($go_3_5)((($dictApplicative_0)->pure)(new Phpurs_Data0("Nil"))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "sequence" => function($dictApplicative_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Types_traversableList"];
  $__res = (((($GLOBALS['Data_List_Types_traversableList'] ?? \PhpursThunks::eval('Data_List_Types_traversableList')))->traverse)($dictApplicative_0))((($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))->identity);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Types_traversableList"];
  $__res = ($GLOBALS['Data_List_Types_functorList'] ?? \PhpursThunks::eval('Data_List_Types_functorList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Types_traversableList"];
  $__res = ($GLOBALS['Data_List_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Types_foldableList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Types_traversableNonEmptyList'] = function() { $v = (($GLOBALS['Data_NonEmpty_traversableNonEmpty'] ?? \PhpursThunks::eval('Data_NonEmpty_traversableNonEmpty')))(($GLOBALS['Data_List_Types_traversableList'] ?? \PhpursThunks::eval('Data_List_Types_traversableList'))); return $v; };
\PhpursThunks::$thunks['Data_List_Types_traversableWithIndexList'] = function() { $v = (object)["traverseWithIndex" => function($dictApplicative_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $Apply0_1_0 = (($dictApplicative_0)->Apply0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $__res = function($f_2) use ($Apply0_1_0, $dictApplicative_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $go_3_1 = null;
  $go_3_1 = (function() use (&$go_3_1) {
  $__fn = function($b_4, $v_5 = null) use (&$go_3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_3_1"), recVars=["go_3_1"];
  while (true) {
if ((is_object($v_5) && (($v_5)->tag === "Nil"))) {
$__t2 = $b_4;
} else {
if ((is_object($v_5) && (($v_5)->tag === "Cons"))) {
$__tco_3 = new Phpurs_Data2("Cons", ($v_5)->value0, $b_4);
$__tco_4 = ($v_5)->value1;
$b_4 = $__tco_3;
$v_5 = $__tco_4;
continue ;
$__t2 = null;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
};
};
$__res = $__t2;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((((($Apply0_1_0)->Functor0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->map)(($go_3_1)(new Phpurs_Data0("Nil")))))((((($GLOBALS['Data_List_Types_foldableWithIndexList'] ?? \PhpursThunks::eval('Data_List_Types_foldableWithIndexList')))->foldlWithIndex)((function() use ($Apply0_1_0, $f_2) {
  $__fn = function($i_3, $acc_4 = null) use ($Apply0_1_0, $f_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((((($GLOBALS['Control_Apply_lift2'] ?? \PhpursThunks::eval('Control_Apply_lift2')))($Apply0_1_0))((function() {
  $__fn = function($b_5, $a_6 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = new Phpurs_Data2("Cons", $a_6, $b_5);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($acc_4)))(($f_2)($i_3));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))((($dictApplicative_0)->pure)(new Phpurs_Data0("Nil"))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FunctorWithIndex0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_List_Types_functorWithIndexList'] ?? \PhpursThunks::eval('Data_List_Types_functorWithIndexList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_List_Types_foldableWithIndexList'] ?? \PhpursThunks::eval('Data_List_Types_foldableWithIndexList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_List_Types_traversableList'] ?? \PhpursThunks::eval('Data_List_Types_traversableList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Types_traverseWithIndex'] = function() { $v = ((($GLOBALS['Data_NonEmpty_traversableWithIndexNonEmpty'] ?? \PhpursThunks::eval('Data_NonEmpty_traversableWithIndexNonEmpty')))(($GLOBALS['Data_List_Types_traversableWithIndexList'] ?? \PhpursThunks::eval('Data_List_Types_traversableWithIndexList'))))->traverseWithIndex; return $v; };
\PhpursThunks::$thunks['Data_List_Types_traversableWithIndexNonEmptyList'] = function() { $v = (object)["traverseWithIndex" => function($dictApplicative_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $traverseWithIndex1_1_0 = (($GLOBALS['Data_List_Types_traverseWithIndex'] ?? \PhpursThunks::eval('Data_List_Types_traverseWithIndex')))($dictApplicative_0);
  $__res = (function() use ($dictApplicative_0, $traverseWithIndex1_1_0) {
  $__fn = function($f_2, $v_3 = null) use ($dictApplicative_0, $traverseWithIndex1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((((((($dictApplicative_0)->Apply0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Functor0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->map)(($GLOBALS['Data_List_Types_NonEmptyList'] ?? \PhpursThunks::eval('Data_List_Types_NonEmptyList'))))((($traverseWithIndex1_1_0)(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))($f_2))(function($v2_4) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v2_4) && (($v2_4)->tag === "Nothing"))) {
$__t1 = 0;
} else {
if ((is_object($v2_4) && (($v2_4)->tag === "Just"))) {
$__t1 = (1 + ($v2_4)->value0);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))($v_3));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FunctorWithIndex0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_List_Types_functorWithIndexNonEmptyList'] ?? \PhpursThunks::eval('Data_List_Types_functorWithIndexNonEmptyList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_List_Types_foldableWithIndexNonEmptyList'] ?? \PhpursThunks::eval('Data_List_Types_foldableWithIndexNonEmptyList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_List_Types_traversableNonEmptyList'] ?? \PhpursThunks::eval('Data_List_Types_traversableNonEmptyList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Types_unfoldable1List'] = function() { $v = (object)["unfoldr1" => (function() {
  $__fn = function($f_0, $b_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $go_2_0 = null;
  $go_2_0 = (function() use ($f_0, &$go_2_0) {
  $__fn = function($source_3, $memo_4 = null) use ($f_0, &$go_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_2_0"), recVars=["go_2_0"];
  while (true) {
$v_5_1 = ($f_0)($source_3);
if ((is_object(($v_5_1)->value1) && ((($v_5_1)->value1)->tag === "Just"))) {
$__tco_3 = (($v_5_1)->value1)->value0;
$__tco_4 = new Phpurs_Data2("Cons", ($v_5_1)->value0, $memo_4);
$source_3 = $__tco_3;
$memo_4 = $__tco_4;
continue ;
$__t2 = null;
} else {
if ((is_object(($v_5_1)->value1) && ((($v_5_1)->value1)->tag === "Nothing"))) {
$go_6_5 = null;
$go_6_5 = (function() use (&$go_6_5) {
  $__fn = function($b_7, $v_8 = null) use (&$go_6_5, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_6_5"), recVars=["go_2_0","go_6_5"];
  while (true) {
if ((is_object($v_8) && (($v_8)->tag === "Nil"))) {
$__t6 = $b_7;
} else {
if ((is_object($v_8) && (($v_8)->tag === "Cons"))) {
$__tco_7 = new Phpurs_Data2("Cons", ($v_8)->value0, $b_7);
$__tco_8 = ($v_8)->value1;
$b_7 = $__tco_7;
$v_8 = $__tco_8;
continue ;
$__t6 = null;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t6 = null;
};
};
$__res = $__t6;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
$__t2 = (($go_6_5)(new Phpurs_Data0("Nil")))(new Phpurs_Data2("Cons", ($v_5_1)->value0, $memo_4));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
};
};
$__res = $__t2;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = (($go_2_0)($b_1))(new Phpurs_Data0("Nil"));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Data_List_Types_unfoldableList'] = function() { $v = (object)["unfoldr" => (function() {
  $__fn = function($f_0, $b_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $go_2_0 = null;
  $go_2_0 = (function() use ($f_0, &$go_2_0) {
  $__fn = function($source_3, $memo_4 = null) use ($f_0, &$go_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_2_0"), recVars=["go_2_0"];
  while (true) {
$v_5_1 = ($f_0)($source_3);
if ((is_object($v_5_1) && (($v_5_1)->tag === "Nothing"))) {
$go_6_3 = null;
$go_6_3 = (function() use (&$go_6_3) {
  $__fn = function($b_7, $v_8 = null) use (&$go_6_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_6_3"), recVars=["go_2_0","go_6_3"];
  while (true) {
if ((is_object($v_8) && (($v_8)->tag === "Nil"))) {
$__t4 = $b_7;
} else {
if ((is_object($v_8) && (($v_8)->tag === "Cons"))) {
$__tco_5 = new Phpurs_Data2("Cons", ($v_8)->value0, $b_7);
$__tco_6 = ($v_8)->value1;
$b_7 = $__tco_5;
$v_8 = $__tco_6;
continue ;
$__t4 = null;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t4 = null;
};
};
$__res = $__t4;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
$__t2 = (($go_6_3)(new Phpurs_Data0("Nil")))($memo_4);
} else {
if ((is_object($v_5_1) && (($v_5_1)->tag === "Just"))) {
$__tco_7 = (($v_5_1)->value0)->value1;
$__tco_8 = new Phpurs_Data2("Cons", (($v_5_1)->value0)->value0, $memo_4);
$source_3 = $__tco_7;
$memo_4 = $__tco_8;
continue ;
$__t2 = null;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
};
};
$__res = $__t2;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = (($go_2_0)($b_1))(new Phpurs_Data0("Nil"));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Unfoldable10" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_List_Types_unfoldable1List'] ?? \PhpursThunks::eval('Data_List_Types_unfoldable1List'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Types_unfoldable1NonEmptyList'] = function() { $v = (object)["unfoldr1" => (function() {
  $__fn = function($f_0, $b_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__local_var_2_0 = ($f_0)($b_1);
  $go_3_1 = null;
  $go_3_1 = (function() use ($f_0, &$go_3_1) {
  $__fn = function($source_4, $memo_5 = null) use ($f_0, &$go_3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_3_1"), recVars=["go_3_1"];
  while (true) {
if ((is_object($source_4) && (($source_4)->tag === "Just"))) {
$__tco_7 = (($f_0)(($source_4)->value0))->value1;
$__tco_8 = new Phpurs_Data2("Cons", (($f_0)(($source_4)->value0))->value0, $memo_5);
$source_4 = $__tco_7;
$memo_5 = $__tco_8;
continue ;
$__t6 = null;
} else {
$go_6_2 = null;
$go_6_2 = (function() use (&$go_6_2) {
  $__fn = function($b_7, $v_8 = null) use (&$go_6_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_6_2"), recVars=["go_3_1","go_6_2"];
  while (true) {
if ((is_object($v_8) && (($v_8)->tag === "Nil"))) {
$__t3 = $b_7;
} else {
if ((is_object($v_8) && (($v_8)->tag === "Cons"))) {
$__tco_4 = new Phpurs_Data2("Cons", ($v_8)->value0, $b_7);
$__tco_5 = ($v_8)->value1;
$b_7 = $__tco_4;
$v_8 = $__tco_5;
continue ;
$__t3 = null;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
};
};
$__res = $__t3;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
$__t6 = (($go_6_2)(new Phpurs_Data0("Nil")))($memo_5);
};
$__res = $__t6;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = new Phpurs_Data2("NonEmpty", ($__local_var_2_0)->value0, (($go_3_1)(($__local_var_2_0)->value1))(new Phpurs_Data0("Nil")));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Data_List_Types_foldable1NonEmptyList'] = function() { $v = (($GLOBALS['Data_NonEmpty_foldable1NonEmpty'] ?? \PhpursThunks::eval('Data_NonEmpty_foldable1NonEmpty')))(($GLOBALS['Data_List_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Types_foldableList'))); return $v; };
\PhpursThunks::$thunks['Data_List_Types_extendNonEmptyList'] = function() { $v = (object)["extend" => (function() {
  $__fn = function($f_0, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = new Phpurs_Data2("NonEmpty", ($f_0)($v_1), (((((($GLOBALS['Data_List_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Types_foldableList')))->foldr)((function() use ($f_0) {
  $__fn = function($a_2, $v1_3 = null) use ($f_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (object)["val" => new Phpurs_Data2("Cons", ($f_0)(new Phpurs_Data2("NonEmpty", $a_2, ($v1_3)->acc)), ($v1_3)->val), "acc" => new Phpurs_Data2("Cons", $a_2, ($v1_3)->acc)];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))((object)["val" => new Phpurs_Data0("Nil"), "acc" => new Phpurs_Data0("Nil")]))(($v_1)->value1))->val);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_List_Types_functorNonEmptyList'] ?? \PhpursThunks::eval('Data_List_Types_functorNonEmptyList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Types_extendList'] = function() { $v = (object)["extend" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v1_1) && (($v1_1)->tag === "Nil"))) {
$__t0 = new Phpurs_Data0("Nil");
} else {
if ((is_object($v1_1) && (($v1_1)->tag === "Cons"))) {
$__t0 = new Phpurs_Data2("Cons", ($v_0)($v1_1), (((((($GLOBALS['Data_List_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Types_foldableList')))->foldr)((function() use ($v_0) {
  $__fn = function($a__prime___2, $v2_3 = null) use ($v_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (object)["val" => new Phpurs_Data2("Cons", ($v_0)(new Phpurs_Data2("Cons", $a__prime___2, ($v2_3)->acc)), ($v2_3)->val), "acc" => new Phpurs_Data2("Cons", $a__prime___2, ($v2_3)->acc)];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))((object)["val" => new Phpurs_Data0("Nil"), "acc" => new Phpurs_Data0("Nil")]))(($v1_1)->value1))->val);
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
})(), "Functor0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_List_Types_functorList'] ?? \PhpursThunks::eval('Data_List_Types_functorList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Types_eq1List'] = function() { $v = (object)["eq1" => (function() {
  $__fn = function($dictEq_0, $xs_1 = null, $ys_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $go_3_0 = null;
  $go_3_0 = (function() use ($dictEq_0, &$go_3_0) {
  $__fn = function($v_4, $v1_5 = null, $v2_6 = null) use ($dictEq_0, &$go_3_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_3_0"), recVars=["go_3_0"];
  while (true) {
if (( ! $v2_6)) {
$__t1 = false;
} else {
if ((is_object($v_4) && (($v_4)->tag === "Nil"))) {
$__t1 = ((is_object($v1_5) && (($v1_5)->tag === "Nil")) && $v2_6);
} else {
$__t1 = ((is_object($v_4) && (($v_4)->tag === "Cons")) && ((is_object($v1_5) && (($v1_5)->tag === "Cons")) && ((($go_3_0)(($v_4)->value1))(($v1_5)->value1))(($v2_6 && ((($dictEq_0)->eq)(($v1_5)->value0))(($v_4)->value0)))));
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
  $__res = ((($go_3_0)($xs_1))($ys_2))(true);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Data_List_Types_eq1NonEmptyList'] = function() { $v = (object)["eq1" => (function() {
  $__fn = function($dictEq_0, $x_1 = null, $y_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $go_3_0 = null;
  $go_3_0 = (function() use ($dictEq_0, &$go_3_0) {
  $__fn = function($v_4, $v1_5 = null, $v2_6 = null) use ($dictEq_0, &$go_3_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_3_0"), recVars=["go_3_0"];
  while (true) {
if (( ! $v2_6)) {
$__t1 = false;
} else {
if ((is_object($v_4) && (($v_4)->tag === "Nil"))) {
$__t1 = ((is_object($v1_5) && (($v1_5)->tag === "Nil")) && $v2_6);
} else {
$__t1 = ((is_object($v_4) && (($v_4)->tag === "Cons")) && ((is_object($v1_5) && (($v1_5)->tag === "Cons")) && ((($go_3_0)(($v_4)->value1))(($v1_5)->value1))(($v2_6 && ((($dictEq_0)->eq)(($v1_5)->value0))(($v_4)->value0)))));
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
  $__res = (((($dictEq_0)->eq)(($x_1)->value0))(($y_2)->value0) && ((($go_3_0)(($x_1)->value1))(($y_2)->value1))(true));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Data_List_Types_eqList'] = function() { $v = function($dictEq_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Types_eqList"), recVars=[];
  $__res = (object)["eq" => (function() use ($dictEq_0) {
  $__fn = function($xs_1, $ys_2 = null) use ($dictEq_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $go_3_0 = null;
  $go_3_0 = (function() use ($dictEq_0, &$go_3_0) {
  $__fn = function($v_4, $v1_5 = null, $v2_6 = null) use ($dictEq_0, &$go_3_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_3_0"), recVars=["go_3_0"];
  while (true) {
if (( ! $v2_6)) {
$__t1 = false;
} else {
if ((is_object($v_4) && (($v_4)->tag === "Nil"))) {
$__t1 = ((is_object($v1_5) && (($v1_5)->tag === "Nil")) && $v2_6);
} else {
$__t1 = ((is_object($v_4) && (($v_4)->tag === "Cons")) && ((is_object($v1_5) && (($v1_5)->tag === "Cons")) && ((($go_3_0)(($v_4)->value1))(($v1_5)->value1))(($v2_6 && ((($dictEq_0)->eq)(($v1_5)->value0))(($v_4)->value0)))));
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
  $__res = ((($go_3_0)($xs_1))($ys_2))(true);
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
\PhpursThunks::$thunks['Data_List_Types_eqNonEmptyList'] = function() { $v = function($dictEq_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Types_eqNonEmptyList"), recVars=[];
  $__res = (object)["eq" => (function() use ($dictEq_0) {
  $__fn = function($x_1, $y_2 = null) use ($dictEq_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $go_3_0 = null;
  $go_3_0 = (function() use ($dictEq_0, &$go_3_0) {
  $__fn = function($v_4, $v1_5 = null, $v2_6 = null) use ($dictEq_0, &$go_3_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_3_0"), recVars=["go_3_0"];
  while (true) {
if (( ! $v2_6)) {
$__t1 = false;
} else {
if ((is_object($v_4) && (($v_4)->tag === "Nil"))) {
$__t1 = ((is_object($v1_5) && (($v1_5)->tag === "Nil")) && $v2_6);
} else {
$__t1 = ((is_object($v_4) && (($v_4)->tag === "Cons")) && ((is_object($v1_5) && (($v1_5)->tag === "Cons")) && ((($go_3_0)(($v_4)->value1))(($v1_5)->value1))(($v2_6 && ((($dictEq_0)->eq)(($v1_5)->value0))(($v_4)->value0)))));
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
  $__res = (((($dictEq_0)->eq)(($x_1)->value0))(($y_2)->value0) && ((($go_3_0)(($x_1)->value1))(($y_2)->value1))(true));
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
\PhpursThunks::$thunks['Data_List_Types_ord1List'] = function() { $v = (object)["compare1" => (function() {
  $__fn = function($dictOrd_0, $xs_1 = null, $ys_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $go_3_0 = null;
  $go_3_0 = (function() use ($dictOrd_0, &$go_3_0) {
  $__fn = function($v_4, $v1_5 = null) use ($dictOrd_0, &$go_3_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_3_0"), recVars=["go_3_0"];
  while (true) {
if ((is_object($v_4) && (($v_4)->tag === "Nil"))) {
if ((is_object($v1_5) && (($v1_5)->tag === "Nil"))) {
$__t2 = new Phpurs_Data0("EQ");
} else {
$__t2 = new Phpurs_Data0("LT");
};
$__t1 = $__t2;
} else {
if ((is_object($v1_5) && (($v1_5)->tag === "Nil"))) {
$__t1 = new Phpurs_Data0("GT");
} else {
if (((is_object($v_4) && (($v_4)->tag === "Cons")) && (is_object($v1_5) && (($v1_5)->tag === "Cons")))) {
$v2_6_3 = ((($dictOrd_0)->compare)(($v_4)->value0))(($v1_5)->value0);
if ((is_object($v2_6_3) && (($v2_6_3)->tag === "EQ"))) {
$__tco_5 = ($v_4)->value1;
$__tco_6 = ($v1_5)->value1;
$v_4 = $__tco_5;
$v1_5 = $__tco_6;
continue ;
$__t4 = null;
} else {
$__t4 = $v2_6_3;
};
$__t1 = $__t4;
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
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = (($go_3_0)($xs_1))($ys_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "Eq10" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_List_Types_eq1List'] ?? \PhpursThunks::eval('Data_List_Types_eq1List'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Types_ordNonEmpty'] = function() { $v = (($GLOBALS['Data_NonEmpty_ordNonEmpty'] ?? \PhpursThunks::eval('Data_NonEmpty_ordNonEmpty')))(($GLOBALS['Data_List_Types_ord1List'] ?? \PhpursThunks::eval('Data_List_Types_ord1List'))); return $v; };
\PhpursThunks::$thunks['Data_List_Types_ord1NonEmptyList'] = function() { $v = (($GLOBALS['Data_NonEmpty_ord1NonEmpty'] ?? \PhpursThunks::eval('Data_NonEmpty_ord1NonEmpty')))(($GLOBALS['Data_List_Types_ord1List'] ?? \PhpursThunks::eval('Data_List_Types_ord1List'))); return $v; };
\PhpursThunks::$thunks['Data_List_Types_ordList'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Types_ordList"), recVars=[];
  $__local_var_1_0 = (($dictOrd_0)->Eq0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $eqList1_2_1 = (object)["eq" => (function() use ($__local_var_1_0) {
  $__fn = function($xs_2, $ys_3 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $go_4_1 = null;
  $go_4_1 = (function() use ($__local_var_1_0, &$go_4_1) {
  $__fn = function($v_5, $v1_6 = null, $v2_7 = null) use ($__local_var_1_0, &$go_4_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_4_1"), recVars=["go_4_1"];
  while (true) {
if (( ! $v2_7)) {
$__t2 = false;
} else {
if ((is_object($v_5) && (($v_5)->tag === "Nil"))) {
$__t2 = ((is_object($v1_6) && (($v1_6)->tag === "Nil")) && $v2_7);
} else {
$__t2 = ((is_object($v_5) && (($v_5)->tag === "Cons")) && ((is_object($v1_6) && (($v1_6)->tag === "Cons")) && ((($go_4_1)(($v_5)->value1))(($v1_6)->value1))(($v2_7 && ((($__local_var_1_0)->eq)(($v1_6)->value0))(($v_5)->value0)))));
};
};
$__res = $__t2;
goto __end;;
};
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  $__res = ((($go_4_1)($xs_2))($ys_3))(true);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $__res = (object)["compare" => (function() use ($dictOrd_0) {
  $__fn = function($xs_3, $ys_4 = null) use ($dictOrd_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $go_5_4 = null;
  $go_5_4 = (function() use ($dictOrd_0, &$go_5_4) {
  $__fn = function($v_6, $v1_7 = null) use ($dictOrd_0, &$go_5_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_5_4"), recVars=["go_5_4"];
  while (true) {
if ((is_object($v_6) && (($v_6)->tag === "Nil"))) {
if ((is_object($v1_7) && (($v1_7)->tag === "Nil"))) {
$__t6 = new Phpurs_Data0("EQ");
} else {
$__t6 = new Phpurs_Data0("LT");
};
$__t5 = $__t6;
} else {
if ((is_object($v1_7) && (($v1_7)->tag === "Nil"))) {
$__t5 = new Phpurs_Data0("GT");
} else {
if (((is_object($v_6) && (($v_6)->tag === "Cons")) && (is_object($v1_7) && (($v1_7)->tag === "Cons")))) {
$v2_8_7 = ((($dictOrd_0)->compare)(($v_6)->value0))(($v1_7)->value0);
if ((is_object($v2_8_7) && (($v2_8_7)->tag === "EQ"))) {
$__tco_9 = ($v_6)->value1;
$__tco_10 = ($v1_7)->value1;
$v_6 = $__tco_9;
$v1_7 = $__tco_10;
continue ;
$__t8 = null;
} else {
$__t8 = $v2_8_7;
};
$__t5 = $__t8;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t5 = null;
};
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
  $__res = (($go_5_4)($xs_3))($ys_4);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($dollar__unused_3) use ($eqList1_2_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $eqList1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_Types_ordNonEmptyList'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Types_ordNonEmptyList"), recVars=[];
  $__res = (($GLOBALS['Data_List_Types_ordNonEmpty'] ?? \PhpursThunks::eval('Data_List_Types_ordNonEmpty')))($dictOrd_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_Types_comonadNonEmptyList'] = function() { $v = (object)["extract" => function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($v_0)->value0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Extend0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_List_Types_extendNonEmptyList'] ?? \PhpursThunks::eval('Data_List_Types_extendNonEmptyList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Types_applyList'] = function() { $v = (object)["apply" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Types_applyList"];
  if ((is_object($v_0) && (($v_0)->tag === "Nil"))) {
$__t0 = new Phpurs_Data0("Nil");
} else {
if ((is_object($v_0) && (($v_0)->tag === "Cons"))) {
$__t0 = ((((($GLOBALS['Data_List_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Types_foldableList')))->foldr)(($GLOBALS['Data_List_Types_Cons'] ?? \PhpursThunks::eval('Data_List_Types_Cons'))))((((($GLOBALS['Data_List_Types_applyList'] ?? \PhpursThunks::eval('Data_List_Types_applyList')))->apply)(($v_0)->value1))($v1_1)))(((($GLOBALS['Data_List_Types_listMap'] ?? \PhpursThunks::eval('Data_List_Types_listMap')))(($v_0)->value0))($v1_1));
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
})(), "Functor0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Types_applyList"];
  $__res = ($GLOBALS['Data_List_Types_functorList'] ?? \PhpursThunks::eval('Data_List_Types_functorList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Types_applyNonEmptyList'] = function() { $v = (object)["apply" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = new Phpurs_Data2("NonEmpty", (($v_0)->value0)(($v1_1)->value0), ((((($GLOBALS['Data_List_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Types_foldableList')))->foldr)(($GLOBALS['Data_List_Types_Cons'] ?? \PhpursThunks::eval('Data_List_Types_Cons'))))((((($GLOBALS['Data_List_Types_applyList'] ?? \PhpursThunks::eval('Data_List_Types_applyList')))->apply)(new Phpurs_Data2("Cons", ($v_0)->value0, ($v_0)->value1)))(($v1_1)->value1)))((((($GLOBALS['Data_List_Types_applyList'] ?? \PhpursThunks::eval('Data_List_Types_applyList')))->apply)(($v_0)->value1))(new Phpurs_Data2("Cons", ($v1_1)->value0, new Phpurs_Data0("Nil")))));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_List_Types_functorNonEmptyList'] ?? \PhpursThunks::eval('Data_List_Types_functorNonEmptyList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Types_bindList'] = function() { $v = (object)["bind" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Types_bindList"];
  if ((is_object($v_0) && (($v_0)->tag === "Nil"))) {
$__t0 = new Phpurs_Data0("Nil");
} else {
if ((is_object($v_0) && (($v_0)->tag === "Cons"))) {
$__t0 = ((((($GLOBALS['Data_List_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Types_foldableList')))->foldr)(($GLOBALS['Data_List_Types_Cons'] ?? \PhpursThunks::eval('Data_List_Types_Cons'))))((((($GLOBALS['Data_List_Types_bindList'] ?? \PhpursThunks::eval('Data_List_Types_bindList')))->bind)(($v_0)->value1))($v1_1)))(($v1_1)(($v_0)->value0));
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
})(), "Apply0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Types_bindList"];
  $__res = ($GLOBALS['Data_List_Types_applyList'] ?? \PhpursThunks::eval('Data_List_Types_applyList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Types_bindNonEmptyList'] = function() { $v = (object)["bind" => (function() {
  $__fn = function($v_0, $f_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $v1_2_0 = ($f_1)(($v_0)->value0);
  $__res = new Phpurs_Data2("NonEmpty", ($v1_2_0)->value0, ((((($GLOBALS['Data_List_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Types_foldableList')))->foldr)(($GLOBALS['Data_List_Types_Cons'] ?? \PhpursThunks::eval('Data_List_Types_Cons'))))((((($GLOBALS['Data_List_Types_bindList'] ?? \PhpursThunks::eval('Data_List_Types_bindList')))->bind)(($v_0)->value1))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_List_Types_toList'] ?? \PhpursThunks::eval('Data_List_Types_toList'))))($f_1))))(($v1_2_0)->value1));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Apply0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_List_Types_applyNonEmptyList'] ?? \PhpursThunks::eval('Data_List_Types_applyNonEmptyList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Types_applicativeList'] = function() { $v = (object)["pure" => function($a_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = new Phpurs_Data2("Cons", $a_0, new Phpurs_Data0("Nil"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_List_Types_applyList'] ?? \PhpursThunks::eval('Data_List_Types_applyList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Types_monadList'] = function() { $v = (object)["Applicative0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_List_Types_applicativeList'] ?? \PhpursThunks::eval('Data_List_Types_applicativeList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_List_Types_bindList'] ?? \PhpursThunks::eval('Data_List_Types_bindList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Types_altNonEmptyList'] = function() { $v = (object)["alt" => (($GLOBALS['Data_List_Types_semigroupNonEmptyList'] ?? \PhpursThunks::eval('Data_List_Types_semigroupNonEmptyList')))->append, "Functor0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_List_Types_functorNonEmptyList'] ?? \PhpursThunks::eval('Data_List_Types_functorNonEmptyList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Types_altList'] = function() { $v = (object)["alt" => (($GLOBALS['Data_List_Types_semigroupList'] ?? \PhpursThunks::eval('Data_List_Types_semigroupList')))->append, "Functor0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_List_Types_functorList'] ?? \PhpursThunks::eval('Data_List_Types_functorList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Types_plusList'] = function() { $v = (object)["empty" => new Phpurs_Data0("Nil"), "Alt0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_List_Types_altList'] ?? \PhpursThunks::eval('Data_List_Types_altList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Types_alternativeList'] = function() { $v = (object)["Applicative0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_List_Types_applicativeList'] ?? \PhpursThunks::eval('Data_List_Types_applicativeList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Plus1" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_List_Types_plusList'] ?? \PhpursThunks::eval('Data_List_Types_plusList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Types_monadPlusList'] = function() { $v = (object)["Monad0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_List_Types_monadList'] ?? \PhpursThunks::eval('Data_List_Types_monadList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Alternative1" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_List_Types_alternativeList'] ?? \PhpursThunks::eval('Data_List_Types_alternativeList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Types_applicativeNonEmptyList'] = function() { $v = (object)["pure" => ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_List_Types_NonEmptyList'] ?? \PhpursThunks::eval('Data_List_Types_NonEmptyList'))))(function($a_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = new Phpurs_Data2("NonEmpty", $a_0, new Phpurs_Data0("Nil"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), "Apply0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_List_Types_applyNonEmptyList'] ?? \PhpursThunks::eval('Data_List_Types_applyNonEmptyList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Types_monadNonEmptyList'] = function() { $v = (object)["Applicative0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_List_Types_applicativeNonEmptyList'] ?? \PhpursThunks::eval('Data_List_Types_applicativeNonEmptyList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_List_Types_bindNonEmptyList'] ?? \PhpursThunks::eval('Data_List_Types_bindNonEmptyList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Types_traversable1NonEmptyList'] = function() { $v = (object)["traverse1" => function($dictApply_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Types_traversable1NonEmptyList"];
  $Functor0_1_0 = (($dictApply_0)->Functor0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $__res = (function() use ($Functor0_1_0, $dictApply_0) {
  $__fn = function($f_2, $v_3 = null) use ($Functor0_1_0, $dictApply_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Types_traversable1NonEmptyList"];
  $go_4_5 = null;
  $go_4_5 = (function() use ($dictApply_0, $f_2, &$go_4_5) {
  $__fn = function($b_5, $v_6 = null) use ($dictApply_0, $f_2, &$go_4_5, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_4_5"), recVars=["Data_List_Types_traversable1NonEmptyList","go_4_5"];
  while (true) {
if ((is_object($v_6) && (($v_6)->tag === "Nil"))) {
$__t6 = $b_5;
} else {
if ((is_object($v_6) && (($v_6)->tag === "Cons"))) {
$__tco_7 = (((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((((($GLOBALS['Control_Apply_lift2'] ?? \PhpursThunks::eval('Control_Apply_lift2')))($dictApply_0))((function() {
  $__fn = function($b_7, $a_8 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Types_traversable1NonEmptyList","go_4_5"];
  $__res = new Phpurs_Data2("NonEmpty", $a_8, new Phpurs_Data2("Cons", ($b_7)->value0, ($b_7)->value1));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($b_5)))($f_2))(($v_6)->value0);
$__tco_8 = ($v_6)->value1;
$b_5 = $__tco_7;
$v_6 = $__tco_8;
continue ;
$__t6 = null;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t6 = null;
};
};
$__res = $__t6;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ((($Functor0_1_0)->map)(function($v1_4) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Types_traversable1NonEmptyList"];
  $go_5_1 = null;
  $go_5_1 = (function() use (&$go_5_1) {
  $__fn = function($b_6, $v_7 = null) use (&$go_5_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_5_1"), recVars=["Data_List_Types_traversable1NonEmptyList","go_5_1"];
  while (true) {
if ((is_object($v_7) && (($v_7)->tag === "Nil"))) {
$__t2 = $b_6;
} else {
if ((is_object($v_7) && (($v_7)->tag === "Cons"))) {
$__tco_3 = new Phpurs_Data2("NonEmpty", ($v_7)->value0, new Phpurs_Data2("Cons", ($b_6)->value0, ($b_6)->value1));
$__tco_4 = ($v_7)->value1;
$b_6 = $__tco_3;
$v_7 = $__tco_4;
continue ;
$__t2 = null;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
};
};
$__res = $__t2;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = (($go_5_1)(((($GLOBALS['Data_List_Types_applicativeNonEmptyList'] ?? \PhpursThunks::eval('Data_List_Types_applicativeNonEmptyList')))->pure)(($v1_4)->value0)))(($v1_4)->value1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($go_4_5)(((($Functor0_1_0)->map)((($GLOBALS['Data_List_Types_applicativeNonEmptyList'] ?? \PhpursThunks::eval('Data_List_Types_applicativeNonEmptyList')))->pure))(($f_2)(($v_3)->value0))))(($v_3)->value1));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "sequence1" => function($dictApply_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Types_traversable1NonEmptyList"];
  $__res = (((($GLOBALS['Data_List_Types_traversable1NonEmptyList'] ?? \PhpursThunks::eval('Data_List_Types_traversable1NonEmptyList')))->traverse1)($dictApply_0))((($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))->identity);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable10" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Types_traversable1NonEmptyList"];
  $__res = ($GLOBALS['Data_List_Types_foldable1NonEmptyList'] ?? \PhpursThunks::eval('Data_List_Types_foldable1NonEmptyList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable1" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Types_traversable1NonEmptyList"];
  $__res = ($GLOBALS['Data_List_Types_traversableNonEmptyList'] ?? \PhpursThunks::eval('Data_List_Types_traversableNonEmptyList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };




























































