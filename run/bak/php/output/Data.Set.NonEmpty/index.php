<?php

namespace Data\Set\NonEmpty;

// ALL IMPORTS: Control.Semigroupoid, Data.Array.NonEmpty, Data.Array.NonEmpty.Internal, Data.Eq, Data.Foldable, Data.Function.Uncurried, Data.Functor, Data.List.NonEmpty, Data.List.Types, Data.Map.Internal, Data.Maybe, Data.Ord, Data.Semigroup, Data.Semigroup.Foldable, Data.Set, Data.Set.NonEmpty, Data.Show, Data.Tuple, Data.Unfoldable, Data.Unfoldable1, Partial.Unsafe, Prelude, Prim, Safe.Coerce
// TO REQUIRE: Control.Semigroupoid, Data.Array.NonEmpty, Data.Array.NonEmpty.Internal, Data.Eq, Data.Foldable, Data.Function.Uncurried, Data.Functor, Data.List.NonEmpty, Data.List.Types, Data.Map.Internal, Data.Maybe, Data.Ord, Data.Semigroup, Data.Semigroup.Foldable, Data.Set, Data.Set.NonEmpty, Data.Show, Data.Tuple, Data.Unfoldable, Data.Unfoldable1, Partial.Unsafe, Prelude, Safe.Coerce
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Array.NonEmpty/index.php';
require_once __DIR__ . '/../Data.Array.NonEmpty.Internal/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Foldable/index.php';
require_once __DIR__ . '/../Data.Function.Uncurried/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.List.NonEmpty/index.php';
require_once __DIR__ . '/../Data.List.Types/index.php';
require_once __DIR__ . '/../Data.Map.Internal/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Semigroup.Foldable/index.php';
require_once __DIR__ . '/../Data.Set/index.php';
require_once __DIR__ . '/../Data.Set.NonEmpty/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
require_once __DIR__ . '/../Data.Tuple/index.php';
require_once __DIR__ . '/../Data.Unfoldable/index.php';
require_once __DIR__ . '/../Data.Unfoldable1/index.php';
require_once __DIR__ . '/../Partial.Unsafe/index.php';
require_once __DIR__ . '/../Prelude/index.php';
require_once __DIR__ . '/../Safe.Coerce/index.php';

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
\PhpursThunks::$thunks['Data_Set_NonEmpty_unionSet'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Set_NonEmpty_unionSet"), recVars=[];
  $compare_1_0 = ($dictOrd_0)->compare;
  $__res = (function() use ($compare_1_0) {
  $__fn = function($m1_2, $m2_3 = null) use ($compare_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Data_Map_Internal_unsafeUnionWith'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeUnionWith')))($compare_1_0, ($GLOBALS['Data_Function_const'] ?? \PhpursThunks::eval('Data_Function_const')), $m1_2, $m2_3);
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
\PhpursThunks::$thunks['Data_Set_NonEmpty_toUnfoldable1'] = function() { $v = function($dictUnfoldable1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Set_NonEmpty_toUnfoldable1"), recVars=[];
  $stepNext_1_0 = (((($GLOBALS['Data_Map_Internal_stepWith'] ?? \PhpursThunks::eval('Data_Map_Internal_stepWith')))(($GLOBALS['Data_Map_Internal_iterMapL'] ?? \PhpursThunks::eval('Data_Map_Internal_iterMapL'))))((function() {
  $__fn = function($k_1, $v_2 = null, $next_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = new Phpurs_Data1("Just", new Phpurs_Data2("Tuple", $k_1, $next_3));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()))(function($v_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = new Phpurs_Data0("Nothing");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($dictUnfoldable1_0)->unfoldr1)(function($v_2) use ($stepNext_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = new Phpurs_Data2("Tuple", ($v_2)->value0, ($stepNext_1_0)(($v_2)->value1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((((($GLOBALS['Data_Map_Internal_stepWith'] ?? \PhpursThunks::eval('Data_Map_Internal_stepWith')))(($GLOBALS['Data_Map_Internal_iterMapL'] ?? \PhpursThunks::eval('Data_Map_Internal_iterMapL'))))((function() {
  $__fn = function($k_2, $v_3 = null, $next_4 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = new Phpurs_Data2("Tuple", $k_2, $next_4);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()))(function($v_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Partial__crashWith'] ?? \PhpursThunks::eval('Partial__crashWith')))("toUnfoldable1: impossible");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Map_Internal_toMapIter'] ?? \PhpursThunks::eval('Data_Map_Internal_toMapIter'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Set_toMap'] ?? \PhpursThunks::eval('Data_Set_toMap'))))(($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce'))))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Set_NonEmpty_toUnfoldable11'] = function() { $v = (($GLOBALS['Data_Set_NonEmpty_toUnfoldable1'] ?? \PhpursThunks::eval('Data_Set_NonEmpty_toUnfoldable1')))(($GLOBALS['Data_Unfoldable1_unfoldable1Array'] ?? \PhpursThunks::eval('Data_Unfoldable1_unfoldable1Array'))); return $v; };
\PhpursThunks::$thunks['Data_Set_NonEmpty_toUnfoldable12'] = function() { $v = (($GLOBALS['Data_Set_NonEmpty_toUnfoldable1'] ?? \PhpursThunks::eval('Data_Set_NonEmpty_toUnfoldable1')))(($GLOBALS['Data_List_Types_unfoldable1NonEmptyList'] ?? \PhpursThunks::eval('Data_List_Types_unfoldable1NonEmptyList'))); return $v; };
\PhpursThunks::$thunks['Data_Set_NonEmpty_toUnfoldable'] = function() { $v = function($dictUnfoldable_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Set_NonEmpty_toUnfoldable"), recVars=[];
  $__res = (($GLOBALS['Data_Set_toUnfoldable'] ?? \PhpursThunks::eval('Data_Set_toUnfoldable')))($dictUnfoldable_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Set_NonEmpty_toSet'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Set_NonEmpty_toSet"), recVars=[];
  $__res = $v_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Set_NonEmpty_subset'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Set_NonEmpty_subset"), recVars=[];
  $compare_1_0 = ($dictOrd_0)->compare;
  $__res = (function() use ($compare_1_0) {
  $__fn = function($s1_2, $s2_3 = null) use ($compare_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (is_object((($GLOBALS['Data_Map_Internal_unsafeDifference'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeDifference')))($compare_1_0, $s1_2, $s2_3)) && (((($GLOBALS['Data_Map_Internal_unsafeDifference'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeDifference')))($compare_1_0, $s1_2, $s2_3))->tag === "Leaf"));
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
\PhpursThunks::$thunks['Data_Set_NonEmpty_size'] = function() { $v = ($GLOBALS['Data_Map_Internal_size'] ?? \PhpursThunks::eval('Data_Map_Internal_size')); return $v; };
\PhpursThunks::$thunks['Data_Set_NonEmpty_singleton'] = function() { $v = ($GLOBALS['Data_Set_singleton'] ?? \PhpursThunks::eval('Data_Set_singleton')); return $v; };
\PhpursThunks::$thunks['Data_Set_NonEmpty_showNonEmptySet'] = function() { $v = function($dictShow_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Set_NonEmpty_showNonEmptySet"), recVars=[];
  $__res = (object)["show" => function($s_1) use ($dictShow_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (("(fromFoldable1 (NonEmptyArray " . ((($GLOBALS['Data_Show_showArrayImpl'] ?? \PhpursThunks::eval('Data_Show_showArrayImpl')))(($dictShow_0)->show))((($GLOBALS['Data_Set_NonEmpty_toUnfoldable11'] ?? \PhpursThunks::eval('Data_Set_NonEmpty_toUnfoldable11')))($s_1))) . "))");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Set_NonEmpty_semigroupNonEmptySet'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Set_NonEmpty_semigroupNonEmptySet"), recVars=[];
  $compare_1_0 = ($dictOrd_0)->compare;
  $__res = (object)["append" => (function() use ($compare_1_0) {
  $__fn = function($m1_2, $m2_3 = null) use ($compare_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Data_Map_Internal_unsafeUnionWith'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeUnionWith')))($compare_1_0, ($GLOBALS['Data_Function_const'] ?? \PhpursThunks::eval('Data_Function_const')), $m1_2, $m2_3);
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
\PhpursThunks::$thunks['Data_Set_NonEmpty_properSubset'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Set_NonEmpty_properSubset"), recVars=[];
  $__res = (($GLOBALS['Data_Set_properSubset'] ?? \PhpursThunks::eval('Data_Set_properSubset')))($dictOrd_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Set_NonEmpty_ordNonEmptySet'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Set_NonEmpty_ordNonEmptySet"), recVars=[];
  $__res = (($GLOBALS['Data_Set_ordSet'] ?? \PhpursThunks::eval('Data_Set_ordSet')))($dictOrd_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Set_NonEmpty_ord1NonEmptySet'] = function() { $v = ($GLOBALS['Data_Set_ord1Set'] ?? \PhpursThunks::eval('Data_Set_ord1Set')); return $v; };
\PhpursThunks::$thunks['Data_Set_NonEmpty_min'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Set_NonEmpty_min"), recVars=[];
  $__local_var_1_0 = (($GLOBALS['Data_Map_Internal_findMin'] ?? \PhpursThunks::eval('Data_Map_Internal_findMin')))($v_0);
  if ((is_object($__local_var_1_0) && (($__local_var_1_0)->tag === "Just"))) {
$__t1 = (($__local_var_1_0)->value0)->key;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Set_NonEmpty_member'] = function() { $v = (function() {
  $__fn = function($dictOrd_0, $k_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Set_NonEmpty_member"), recVars=[];
  $go_2_0 = null;
  $go_2_0 = function($v_3) use ($dictOrd_0, &$go_2_0, $k_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_2_0"), recVars=["go_2_0"];
  while (true) {
if ((is_object($v_3) && (($v_3)->tag === "Leaf"))) {
$__t1 = false;
} else {
if ((is_object($v_3) && (($v_3)->tag === "Node"))) {
$v1_4_2 = ((($dictOrd_0)->compare)($k_1))(($v_3)->value2);
if ((is_object($v1_4_2) && (($v1_4_2)->tag === "LT"))) {
$__tco_4 = ($v_3)->value4;
$v_3 = $__tco_4;
continue ;
$__t3 = null;
} else {
if ((is_object($v1_4_2) && (($v1_4_2)->tag === "GT"))) {
$__tco_5 = ($v_3)->value5;
$v_3 = $__tco_5;
continue ;
$__t3 = null;
} else {
if ((is_object($v1_4_2) && (($v1_4_2)->tag === "EQ"))) {
$__t3 = true;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
};
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
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = $go_2_0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Set_NonEmpty_max'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Set_NonEmpty_max"), recVars=[];
  $__local_var_1_0 = (($GLOBALS['Data_Map_Internal_findMax'] ?? \PhpursThunks::eval('Data_Map_Internal_findMax')))($v_0);
  if ((is_object($__local_var_1_0) && (($__local_var_1_0)->tag === "Just"))) {
$__t1 = (($__local_var_1_0)->value0)->key;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Set_NonEmpty_mapMaybe'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Set_NonEmpty_mapMaybe"), recVars=[];
  $__res = (($GLOBALS['Data_Set_mapMaybe'] ?? \PhpursThunks::eval('Data_Set_mapMaybe')))($dictOrd_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Set_NonEmpty_map'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Set_NonEmpty_map"), recVars=[];
  $__res = (($GLOBALS['Data_Set_map'] ?? \PhpursThunks::eval('Data_Set_map')))($dictOrd_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Set_NonEmpty_insert'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Set_NonEmpty_insert"), recVars=[];
  $__res = (($GLOBALS['Data_Set_insert'] ?? \PhpursThunks::eval('Data_Set_insert')))($dictOrd_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Set_NonEmpty_fromSet'] = function() { $v = function($s_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Set_NonEmpty_fromSet"), recVars=[];
  if ((is_object($s_0) && (($s_0)->tag === "Leaf"))) {
$__t0 = new Phpurs_Data0("Nothing");
} else {
$__t0 = new Phpurs_Data1("Just", $s_0);
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Set_NonEmpty_intersection'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Set_NonEmpty_intersection"), recVars=[];
  $compare_1_0 = ($dictOrd_0)->compare;
  $__res = (function() use ($compare_1_0) {
  $__fn = function($v_2, $v1_3 = null) use ($compare_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__local_var_4_1 = (($GLOBALS['Data_Map_Internal_unsafeIntersectionWith'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeIntersectionWith')))($compare_1_0, ($GLOBALS['Data_Function_const'] ?? \PhpursThunks::eval('Data_Function_const')), $v_2, $v1_3);
  if ((is_object($__local_var_4_1) && (($__local_var_4_1)->tag === "Leaf"))) {
$__t2 = new Phpurs_Data0("Nothing");
} else {
$__t2 = new Phpurs_Data1("Just", $__local_var_4_1);
};
  $__res = $__t2;
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
\PhpursThunks::$thunks['Data_Set_NonEmpty_fromFoldable1'] = function() { $v = (function() {
  $__fn = function($dictFoldable1_0, $dictOrd_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Set_NonEmpty_fromFoldable1"), recVars=[];
  $compare_2_0 = ($dictOrd_1)->compare;
  $__res = ((($dictFoldable1_0)->foldMap1)((object)["append" => (function() use ($compare_2_0) {
  $__fn = function($m1_3, $m2_4 = null) use ($compare_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Data_Map_Internal_unsafeUnionWith'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeUnionWith')))($compare_2_0, ($GLOBALS['Data_Function_const'] ?? \PhpursThunks::eval('Data_Function_const')), $m1_3, $m2_4);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()]))(($GLOBALS['Data_Set_singleton'] ?? \PhpursThunks::eval('Data_Set_singleton')));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Set_NonEmpty_fromFoldable'] = function() { $v = (function() {
  $__fn = function($dictFoldable_0, $dictOrd_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Set_NonEmpty_fromFoldable"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Set_NonEmpty_fromSet'] ?? \PhpursThunks::eval('Data_Set_NonEmpty_fromSet'))))(((($dictFoldable_0)->foldl)((function() use ($dictOrd_1) {
  $__fn = function($m_2, $a_3 = null) use ($dictOrd_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((((($GLOBALS['Data_Map_Internal_insert'] ?? \PhpursThunks::eval('Data_Map_Internal_insert')))($dictOrd_1))($a_3))(($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit'))))($m_2);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(new Phpurs_Data0("Leaf")));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Set_NonEmpty_foldableNonEmptySet'] = function() { $v = ($GLOBALS['Data_Set_foldableSet'] ?? \PhpursThunks::eval('Data_Set_foldableSet')); return $v; };
\PhpursThunks::$thunks['Data_Set_NonEmpty_foldable1NonEmptySet'] = function() { $v = (object)["foldMap1" => function($dictSemigroup_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $foldMap11_1_0 = ((($GLOBALS['Data_List_Types_foldable1NonEmptyList'] ?? \PhpursThunks::eval('Data_List_Types_foldable1NonEmptyList')))->foldMap1)($dictSemigroup_0);
  $__res = function($f_2) use ($foldMap11_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($foldMap11_1_0)($f_2)))(($GLOBALS['Data_Set_NonEmpty_toUnfoldable12'] ?? \PhpursThunks::eval('Data_Set_NonEmpty_toUnfoldable12')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldr1" => function($f_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(((($GLOBALS['Data_List_Types_foldable1NonEmptyList'] ?? \PhpursThunks::eval('Data_List_Types_foldable1NonEmptyList')))->foldr1)($f_0)))(($GLOBALS['Data_Set_NonEmpty_toUnfoldable12'] ?? \PhpursThunks::eval('Data_Set_NonEmpty_toUnfoldable12')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldl1" => function($f_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(((($GLOBALS['Data_List_Types_foldable1NonEmptyList'] ?? \PhpursThunks::eval('Data_List_Types_foldable1NonEmptyList')))->foldl1)($f_0)))(($GLOBALS['Data_Set_NonEmpty_toUnfoldable12'] ?? \PhpursThunks::eval('Data_Set_NonEmpty_toUnfoldable12')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_Set_foldableSet'] ?? \PhpursThunks::eval('Data_Set_foldableSet'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Set_NonEmpty_filter'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Set_NonEmpty_filter"), recVars=[];
  $__res = (($GLOBALS['Data_Set_filter'] ?? \PhpursThunks::eval('Data_Set_filter')))($dictOrd_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Set_NonEmpty_eqNonEmptySet'] = function() { $v = function($dictEq_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Set_NonEmpty_eqNonEmptySet"), recVars=[];
  $__res = (object)["eq" => (function() use ($dictEq_0) {
  $__fn = function($v_1, $v1_2 = null) use ($dictEq_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (((((($GLOBALS['Data_Map_Internal_eqMap'] ?? \PhpursThunks::eval('Data_Map_Internal_eqMap')))($dictEq_0))(($GLOBALS['Data_Eq_eqUnit'] ?? \PhpursThunks::eval('Data_Eq_eqUnit'))))->eq)($v_1))($v1_2);
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
\PhpursThunks::$thunks['Data_Set_NonEmpty_eq1NonEmptySet'] = function() { $v = ($GLOBALS['Data_Set_eq1Set'] ?? \PhpursThunks::eval('Data_Set_eq1Set')); return $v; };
\PhpursThunks::$thunks['Data_Set_NonEmpty_difference'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Set_NonEmpty_difference"), recVars=[];
  $compare_1_0 = ($dictOrd_0)->compare;
  $__res = (function() use ($compare_1_0) {
  $__fn = function($v_2, $v1_3 = null) use ($compare_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__local_var_4_1 = (($GLOBALS['Data_Map_Internal_unsafeDifference'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeDifference')))($compare_1_0, $v_2, $v1_3);
  if ((is_object($__local_var_4_1) && (($__local_var_4_1)->tag === "Leaf"))) {
$__t2 = new Phpurs_Data0("Nothing");
} else {
$__t2 = new Phpurs_Data1("Just", $__local_var_4_1);
};
  $__res = $__t2;
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
\PhpursThunks::$thunks['Data_Set_NonEmpty_delete'] = function() { $v = (function() {
  $__fn = function($dictOrd_0, $a_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Set_NonEmpty_delete"), recVars=[];
  $__local_var_3_0 = (((($GLOBALS['Data_Map_Internal_delete'] ?? \PhpursThunks::eval('Data_Map_Internal_delete')))($dictOrd_0))($a_1))($v_2);
  if ((is_object($__local_var_3_0) && (($__local_var_3_0)->tag === "Leaf"))) {
$__t1 = new Phpurs_Data0("Nothing");
} else {
$__t1 = new Phpurs_Data1("Just", $__local_var_3_0);
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Set_NonEmpty_cons'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Set_NonEmpty_cons"), recVars=[];
  $__res = (($GLOBALS['Data_Set_insert'] ?? \PhpursThunks::eval('Data_Set_insert')))($dictOrd_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };


































