<?php

namespace Data\Set;

// ALL IMPORTS: Control.Category, Control.Semigroupoid, Data.Eq, Data.Foldable, Data.Function, Data.Functor, Data.HeytingAlgebra, Data.List, Data.List.Types, Data.Map.Internal, Data.Maybe, Data.Monoid, Data.Ord, Data.Semigroup, Data.Set, Data.Show, Data.Unfoldable, Data.Unit, Prelude, Prim, Safe.Coerce
// TO REQUIRE: Control.Category, Control.Semigroupoid, Data.Eq, Data.Foldable, Data.Function, Data.Functor, Data.HeytingAlgebra, Data.List, Data.List.Types, Data.Map.Internal, Data.Maybe, Data.Monoid, Data.Ord, Data.Semigroup, Data.Set, Data.Show, Data.Unfoldable, Data.Unit, Prelude, Safe.Coerce
require_once __DIR__ . '/../Control.Category/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Foldable/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.HeytingAlgebra/index.php';
require_once __DIR__ . '/../Data.List/index.php';
require_once __DIR__ . '/../Data.List.Types/index.php';
require_once __DIR__ . '/../Data.Map.Internal/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Monoid/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Set/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
require_once __DIR__ . '/../Data.Unfoldable/index.php';
require_once __DIR__ . '/../Data.Unit/index.php';
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
\PhpursThunks::$thunks['Data_Set_Set'] = function() { $v = function($x_0 = null) {
  $__num = \func_num_args();
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Set_union'] = function() { $v = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $compare_1_0 = ($dictOrd_0)->{'compare'};
  $__res = (($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce')))((function() use ($compare_1_0) {
  $__fn = function($m1_2 = null, $m2_3 = null) use ($compare_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((((($GLOBALS['Data_Function_Uncurried_runFn4'] ?? \PhpursThunks::eval('Data_Function_Uncurried_runFn4')))(($GLOBALS['Data_Map_Internal_unsafeUnionWith'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeUnionWith'))))($compare_1_0))(($GLOBALS['Data_Function_const'] ?? \PhpursThunks::eval('Data_Function_const'))))($m1_2))($m2_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})());
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Set_toggle'] = function() { $v = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $alter_1_0 = (($GLOBALS['Data_Map_Internal_alter'] ?? \PhpursThunks::eval('Data_Map_Internal_alter')))($dictOrd_0);
  $__res = (function() use ($alter_1_0) {
  $__fn = function($a_2 = null, $v_3 = null) use ($alter_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($alter_1_0)(function($v2_4 = null) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((is_object($v2_4) && (($v2_4)->{'tag'} === "Nothing"))) {
$__t1 = new Phpurs_Data1("Just", ($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit')));
goto end_branch_1;;
};
  if ((is_object($v2_4) && (($v2_4)->{'tag'} === "Just"))) {
$__t1 = new Phpurs_Data0("Nothing");
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($a_2))($v_3);
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
\PhpursThunks::$thunks['Data_Set_toMap'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = $v_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Set_toList'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Map_Internal_keys'] ?? \PhpursThunks::eval('Data_Map_Internal_keys')))($v_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Set_toUnfoldable'] = function() { $v = function($dictUnfoldable_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($dictUnfoldable_0)->{'unfoldr'})(function($xs_1 = null) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object($xs_1) && (($xs_1)->{'tag'} === "Nil"))) {
$__t0 = new Phpurs_Data0("Nothing");
goto end_branch_0;;
};
  if ((is_object($xs_1) && (($xs_1)->{'tag'} === "Cons"))) {
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data2("Tuple", ($xs_1)->{'value0'}, ($xs_1)->{'value1'}));
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(($GLOBALS['Data_Set_toList'] ?? \PhpursThunks::eval('Data_Set_toList')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Set_toUnfoldable1'] = function() { $v = (($GLOBALS['Data_Set_toUnfoldable'] ?? \PhpursThunks::eval('Data_Set_toUnfoldable')))(($GLOBALS['Data_Unfoldable_unfoldableArray'] ?? \PhpursThunks::eval('Data_Unfoldable_unfoldableArray'))); return $v; };
\PhpursThunks::$thunks['Data_Set_size'] = function() { $v = (($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce')))(($GLOBALS['Data_Map_Internal_size'] ?? \PhpursThunks::eval('Data_Map_Internal_size'))); return $v; };
\PhpursThunks::$thunks['Data_Set_singleton'] = function() { $v = function($a_0 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data6("Node", 1, 1, $a_0, ($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit')), new Phpurs_Data0("Leaf"), new Phpurs_Data0("Leaf"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Set_showSet'] = function() { $v = function($dictShow_0 = null) {
  $__num = \func_num_args();
  $__res = (object)["show" => function($s_1 = null) use ($dictShow_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))("(fromFoldable "))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))(((($GLOBALS['Data_Show_showArrayImpl'] ?? \PhpursThunks::eval('Data_Show_showArrayImpl')))(($dictShow_0)->{'show'}))((($GLOBALS['Data_Set_toUnfoldable1'] ?? \PhpursThunks::eval('Data_Set_toUnfoldable1')))($s_1))))(")"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Set_semigroupSet'] = function() { $v = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $__res = (object)["append" => (($GLOBALS['Data_Set_union'] ?? \PhpursThunks::eval('Data_Set_union')))($dictOrd_0)];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Set_member'] = function() { $v = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce')))(function($k_1 = null) use ($dictOrd_0) {
  $__num = \func_num_args();
  $go_2_0 = null;
  $go_2_0 = function($v_3 = null) use ($dictOrd_0, &$go_2_0, $k_1) {
  $__num = \func_num_args();
  $__tco_var_go_2_0_0_v_3 = $v_3;
  $__tco_done_go_2_0_0 = false;
  $__tco_res_go_2_0_0 = null;
  $__tco_loop_go_2_0_0 = function($v_3 = null) use (&$__tco_done_go_2_0_0, &$__tco_var_go_2_0_0_v_3, $dictOrd_0, &$go_2_0, $k_1) {
  $__num = \func_num_args();
  $__tco_done_go_2_0_0 = true;
  $v_3 = $__tco_var_go_2_0_0_v_3;
  $__t0 = null;;
  if ((is_object($v_3) && (($v_3)->{'tag'} === "Leaf"))) {
$__t0 = false;
goto end_branch_0;;
};
  if ((is_object($v_3) && (($v_3)->{'tag'} === "Node"))) {
$v1_4_1 = ((($dictOrd_0)->{'compare'})($k_1))(($v_3)->{'value2'});
$__t2 = null;;
if ((is_object($v1_4_1) && (($v1_4_1)->{'tag'} === "LT"))) {
$__tco_3 = ($v_3)->{'value4'};
$__tco_var_go_2_0_0_v_3 = $__tco_3;
$__tco_done_go_2_0_0 = false;
$__res = null;
goto __end;;
$__t2 = null;
goto end_branch_2;;
};
if ((is_object($v1_4_1) && (($v1_4_1)->{'tag'} === "GT"))) {
$__tco_4 = ($v_3)->{'value5'};
$__tco_var_go_2_0_0_v_3 = $__tco_4;
$__tco_done_go_2_0_0 = false;
$__res = null;
goto __end;;
$__t2 = null;
goto end_branch_2;;
};
if ((is_object($v1_4_1) && (($v1_4_1)->{'tag'} === "EQ"))) {
$__t2 = true;
goto end_branch_2;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
end_branch_2:;
$__t0 = $__t2;
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
  while (($__tco_done_go_2_0_0 === false)) {
$__tco_res_go_2_0_0 = ($__tco_loop_go_2_0_0)($__tco_var_go_2_0_0_v_3);
};
  $__res = $__tco_res_go_2_0_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = $go_2_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Set_isEmpty'] = function() { $v = (($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce')))(($GLOBALS['Data_Map_Internal_isEmpty'] ?? \PhpursThunks::eval('Data_Map_Internal_isEmpty'))); return $v; };
\PhpursThunks::$thunks['Data_Set_intersection'] = function() { $v = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $compare_1_0 = ($dictOrd_0)->{'compare'};
  $__res = (($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce')))((function() use ($compare_1_0) {
  $__fn = function($m1_2 = null, $m2_3 = null) use ($compare_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((((($GLOBALS['Data_Function_Uncurried_runFn4'] ?? \PhpursThunks::eval('Data_Function_Uncurried_runFn4')))(($GLOBALS['Data_Map_Internal_unsafeIntersectionWith'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeIntersectionWith'))))($compare_1_0))(($GLOBALS['Data_Function_const'] ?? \PhpursThunks::eval('Data_Function_const'))))($m1_2))($m2_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})());
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Set_insert'] = function() { $v = (function() {
  $__fn = function($dictOrd_0 = null, $a_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((($GLOBALS['Data_Map_Internal_insert'] ?? \PhpursThunks::eval('Data_Map_Internal_insert')))($dictOrd_0))($a_1))(($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit'))))($v_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Set_fromMap'] = function() { $v = ($GLOBALS['Data_Set_Set'] ?? \PhpursThunks::eval('Data_Set_Set')); return $v; };
\PhpursThunks::$thunks['Data_Set_foldableSet'] = function() { $v = (object)["foldMap" => function($dictMonoid_0 = null) {
  $__num = \func_num_args();
  $foldMap1_1_0 = ((($GLOBALS['Data_List_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Types_foldableList')))->{'foldMap'})($dictMonoid_0);
  $__res = function($f_2 = null) use ($foldMap1_1_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($foldMap1_1_0)($f_2)))(($GLOBALS['Data_Set_toList'] ?? \PhpursThunks::eval('Data_Set_toList')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldl" => (function() {
  $__fn = function($f_0 = null, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go_2_1 = null;
  $go_2_1 = (function() use ($f_0, &$go_2_1) {
  $__fn = function($b_3 = null, $v_4 = null) use ($f_0, &$go_2_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go_2_1_1_b_3 = $b_3;
  $__tco_var_go_2_1_1_v_4 = $v_4;
  $__tco_done_go_2_1_1 = false;
  $__tco_res_go_2_1_1 = null;
  $__tco_loop_go_2_1_1 = (function() use (&$__tco_done_go_2_1_1, &$__tco_var_go_2_1_1_b_3, &$__tco_var_go_2_1_1_v_4, $f_0, &$go_2_1) {
  $__fn = function($b_3 = null, $v_4 = null) use (&$__tco_done_go_2_1_1, &$__tco_var_go_2_1_1_b_3, &$__tco_var_go_2_1_1_v_4, $f_0, &$go_2_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_done_go_2_1_1 = true;
  $b_3 = $__tco_var_go_2_1_1_b_3;
  $v_4 = $__tco_var_go_2_1_1_v_4;
  $__t1 = null;;
  if ((is_object($v_4) && (($v_4)->{'tag'} === "Nil"))) {
$__t1 = $b_3;
goto end_branch_1;;
};
  if ((is_object($v_4) && (($v_4)->{'tag'} === "Cons"))) {
$__tco_2 = (($f_0)($b_3))(($v_4)->{'value0'});
$__tco_3 = ($v_4)->{'value1'};
$__tco_var_go_2_1_1_b_3 = $__tco_2;
$__tco_var_go_2_1_1_v_4 = $__tco_3;
$__tco_done_go_2_1_1 = false;
$__res = null;
goto __end;;
$__t1 = null;
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
  while (($__tco_done_go_2_1_1 === false)) {
$__tco_res_go_2_1_1 = ($__tco_loop_go_2_1_1)($__tco_var_go_2_1_1_b_3, $__tco_var_go_2_1_1_v_4);
};
  $__res = $__tco_res_go_2_1_1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($go_2_1)($x_1)))(($GLOBALS['Data_Set_toList'] ?? \PhpursThunks::eval('Data_Set_toList')));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "foldr" => (function() {
  $__fn = function($f_0 = null, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((((($GLOBALS['Data_List_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Types_foldableList')))->{'foldr'})($f_0))($x_1)))(($GLOBALS['Data_Set_toList'] ?? \PhpursThunks::eval('Data_Set_toList')));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Data_Set_findMin'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($GLOBALS['Data_Map_Internal_findMin'] ?? \PhpursThunks::eval('Data_Map_Internal_findMin')))($v_0);
  $__t1 = null;;
  if ((is_object($__local_var_1_0) && (($__local_var_1_0)->{'tag'} === "Just"))) {
$__t1 = new Phpurs_Data1("Just", (($__local_var_1_0)->{'value0'})->{'key'});
goto end_branch_1;;
};
  $__t1 = new Phpurs_Data0("Nothing");
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Set_findMax'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($GLOBALS['Data_Map_Internal_findMax'] ?? \PhpursThunks::eval('Data_Map_Internal_findMax')))($v_0);
  $__t1 = null;;
  if ((is_object($__local_var_1_0) && (($__local_var_1_0)->{'tag'} === "Just"))) {
$__t1 = new Phpurs_Data1("Just", (($__local_var_1_0)->{'value0'})->{'key'});
goto end_branch_1;;
};
  $__t1 = new Phpurs_Data0("Nothing");
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Set_filter'] = function() { $v = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce')))((($GLOBALS['Data_Map_Internal_filterKeys'] ?? \PhpursThunks::eval('Data_Map_Internal_filterKeys')))($dictOrd_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Set_eqSet'] = function() { $v = function($dictEq_0 = null) {
  $__num = \func_num_args();
  $__res = (object)["eq" => (function() use ($dictEq_0) {
  $__fn = function($v_1 = null, $v1_2 = null) use ($dictEq_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((((($GLOBALS['Data_Map_Internal_eqMap'] ?? \PhpursThunks::eval('Data_Map_Internal_eqMap')))($dictEq_0))(($GLOBALS['Data_Eq_eqUnit'] ?? \PhpursThunks::eval('Data_Eq_eqUnit'))))->{'eq'})($v_1))($v1_2);
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
\PhpursThunks::$thunks['Data_Set_ordSet'] = function() { $v = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($dictOrd_0)->{'Eq0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $eqSet1_2_1 = (object)["eq" => (function() use ($__local_var_1_0) {
  $__fn = function($v_2 = null, $v1_3 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((((($GLOBALS['Data_Map_Internal_eqMap'] ?? \PhpursThunks::eval('Data_Map_Internal_eqMap')))($__local_var_1_0))(($GLOBALS['Data_Eq_eqUnit'] ?? \PhpursThunks::eval('Data_Eq_eqUnit'))))->{'eq'})($v_2))($v1_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $__res = (object)["compare" => (function() use ($dictOrd_0) {
  $__fn = function($s1_3 = null, $s2_4 = null) use ($dictOrd_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((($GLOBALS['Data_List_Types_ordList'] ?? \PhpursThunks::eval('Data_List_Types_ordList')))($dictOrd_0))->{'compare'})((($GLOBALS['Data_Map_Internal_keys'] ?? \PhpursThunks::eval('Data_Map_Internal_keys')))($s1_3)))((($GLOBALS['Data_Map_Internal_keys'] ?? \PhpursThunks::eval('Data_Map_Internal_keys')))($s2_4));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($dollar__unused_3 = null) use ($eqSet1_2_1) {
  $__num = \func_num_args();
  $__res = $eqSet1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Set_eq1Set'] = function() { $v = (object)["eq1" => (function() {
  $__fn = function($dictEq_0 = null, $v_1 = null, $v1_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((((($GLOBALS['Data_Map_Internal_eqMap'] ?? \PhpursThunks::eval('Data_Map_Internal_eqMap')))($dictEq_0))(($GLOBALS['Data_Eq_eqUnit'] ?? \PhpursThunks::eval('Data_Eq_eqUnit'))))->{'eq'})($v_1))($v1_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Data_Set_ord1Set'] = function() { $v = (object)["compare1" => function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Set_ordSet'] ?? \PhpursThunks::eval('Data_Set_ordSet')))($dictOrd_0))->{'compare'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq10" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Set_eq1Set'] ?? \PhpursThunks::eval('Data_Set_eq1Set'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Set_empty'] = function() { $v = new Phpurs_Data0("Leaf"); return $v; };
\PhpursThunks::$thunks['Data_Set_fromFoldable'] = function() { $v = (function() {
  $__fn = function($dictFoldable_0 = null, $dictOrd_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictFoldable_0)->{'foldl'})((function() use ($dictOrd_1) {
  $__fn = function($m_2 = null, $a_3 = null) use ($dictOrd_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((($GLOBALS['Data_Map_Internal_insert'] ?? \PhpursThunks::eval('Data_Map_Internal_insert')))($dictOrd_1))($a_3))(($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit'))))($m_2);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(new Phpurs_Data0("Leaf"));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Set_map'] = function() { $v = (function() {
  $__fn = function($dictOrd_0 = null, $f_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($GLOBALS['Data_Set_foldableSet'] ?? \PhpursThunks::eval('Data_Set_foldableSet')))->{'foldl'})((function() use ($dictOrd_0, $f_1) {
  $__fn = function($m_2 = null, $a_3 = null) use ($dictOrd_0, $f_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((($GLOBALS['Data_Map_Internal_insert'] ?? \PhpursThunks::eval('Data_Map_Internal_insert')))($dictOrd_0))(($f_1)($a_3)))(($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit'))))($m_2);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(new Phpurs_Data0("Leaf"));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Set_mapMaybe'] = function() { $v = (function() {
  $__fn = function($dictOrd_0 = null, $f_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((((($GLOBALS['Data_List_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Types_foldableList')))->{'foldr'})((function() use ($dictOrd_0, $f_1) {
  $__fn = function($a_2 = null, $acc_3 = null) use ($dictOrd_0, $f_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_4_0 = ($f_1)($a_2);
  $__t1 = null;;
  if ((is_object($__local_var_4_0) && (($__local_var_4_0)->{'tag'} === "Nothing"))) {
$__t1 = $acc_3;
goto end_branch_1;;
};
  if ((is_object($__local_var_4_0) && (($__local_var_4_0)->{'tag'} === "Just"))) {
$__t1 = ((((($GLOBALS['Data_Map_Internal_insert'] ?? \PhpursThunks::eval('Data_Map_Internal_insert')))($dictOrd_0))(($__local_var_4_0)->{'value0'}))(($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit'))))($acc_3);
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
})()))(new Phpurs_Data0("Leaf"))))(($GLOBALS['Data_Set_toList'] ?? \PhpursThunks::eval('Data_Set_toList')));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Set_monoidSet'] = function() { $v = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $semigroupSet1_1_0 = (object)["append" => (($GLOBALS['Data_Set_union'] ?? \PhpursThunks::eval('Data_Set_union')))($dictOrd_0)];
  $__res = (object)["mempty" => new Phpurs_Data0("Leaf"), "Semigroup0" => function($dollar__unused_2 = null) use ($semigroupSet1_1_0) {
  $__num = \func_num_args();
  $__res = $semigroupSet1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Set_unions'] = function() { $v = (function() {
  $__fn = function($dictFoldable_0 = null, $dictOrd_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictFoldable_0)->{'foldl'})((($GLOBALS['Data_Set_union'] ?? \PhpursThunks::eval('Data_Set_union')))($dictOrd_1)))(new Phpurs_Data0("Leaf"));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Set_difference'] = function() { $v = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $compare_1_0 = ($dictOrd_0)->{'compare'};
  $__res = (($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce')))((function() use ($compare_1_0) {
  $__fn = function($m1_2 = null, $m2_3 = null) use ($compare_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((($GLOBALS['Data_Function_Uncurried_runFn3'] ?? \PhpursThunks::eval('Data_Function_Uncurried_runFn3')))(($GLOBALS['Data_Map_Internal_unsafeDifference'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeDifference'))))($compare_1_0))($m1_2))($m2_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})());
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Set_subset'] = function() { $v = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $compare_1_0 = ($dictOrd_0)->{'compare'};
  $difference1_1_0 = (($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce')))((function() use ($compare_1_0) {
  $__fn = function($m1_2 = null, $m2_3 = null) use ($compare_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((($GLOBALS['Data_Function_Uncurried_runFn3'] ?? \PhpursThunks::eval('Data_Function_Uncurried_runFn3')))(($GLOBALS['Data_Map_Internal_unsafeDifference'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeDifference'))))($compare_1_0))($m1_2))($m2_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})());
  $__res = (function() use ($difference1_1_0) {
  $__fn = function($s1_2 = null, $s2_3 = null) use ($difference1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_Set_isEmpty'] ?? \PhpursThunks::eval('Data_Set_isEmpty')))((($difference1_1_0)($s1_2))($s2_3));
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
\PhpursThunks::$thunks['Data_Set_properSubset'] = function() { $v = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $subset1_1_0 = (($GLOBALS['Data_Set_subset'] ?? \PhpursThunks::eval('Data_Set_subset')))($dictOrd_0);
  $__res = (function() use ($subset1_1_0) {
  $__fn = function($s1_2 = null, $s2_3 = null) use ($subset1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_HeytingAlgebra_boolConj'] ?? \PhpursThunks::eval('Data_HeytingAlgebra_boolConj')))(((($GLOBALS['Data_Eq_eqBooleanImpl'] ?? \PhpursThunks::eval('Data_Eq_eqBooleanImpl')))(((($GLOBALS['Data_Eq_eqIntImpl'] ?? \PhpursThunks::eval('Data_Eq_eqIntImpl')))((($GLOBALS['Data_Set_size'] ?? \PhpursThunks::eval('Data_Set_size')))($s1_2)))((($GLOBALS['Data_Set_size'] ?? \PhpursThunks::eval('Data_Set_size')))($s2_3))))(false)))((($subset1_1_0)($s1_2))($s2_3));
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
\PhpursThunks::$thunks['Data_Set_delete'] = function() { $v = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce')))((($GLOBALS['Data_Map_Internal_delete'] ?? \PhpursThunks::eval('Data_Map_Internal_delete')))($dictOrd_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Set_checkValid'] = function() { $v = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce')))((($GLOBALS['Data_Map_Internal_checkValid'] ?? \PhpursThunks::eval('Data_Map_Internal_checkValid')))($dictOrd_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Set_catMaybes'] = function() { $v = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Set_mapMaybe'] ?? \PhpursThunks::eval('Data_Set_mapMaybe')))($dictOrd_0))((($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))->{'identity'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };






































