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

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };


// Data_Set_Set
$GLOBALS['Data_Set_Set'] = function($x_0 = null) {
  $__num = \func_num_args();
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Set_union
$GLOBALS['Data_Set_union'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $compare_1_0 = ($dictOrd_0)['compare'];
  $__res = (function() use ($compare_1_0) {
  $__fn = function($m1_2 = null, $m2_3 = null) use ($compare_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Map_Internal_unsafeUnionWith'])($compare_1_0, $GLOBALS['Data_Function_const'], $m1_2, $m2_3);
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

// Data_Set_toggle
$GLOBALS['Data_Set_toggle'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $alter_1_0 = ($GLOBALS['Data_Map_Internal_alter'])($dictOrd_0);
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
$__t1 = new Phpurs_Data1("Just", $GLOBALS['Data_Unit_unit']);
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
};

// Data_Set_toMap
$GLOBALS['Data_Set_toMap'] = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = $v_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Set_toList
$GLOBALS['Data_Set_toList'] = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Map_Internal_keys'])($v_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Set_toUnfoldable
$GLOBALS['Data_Set_toUnfoldable'] = function($dictUnfoldable_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])(($GLOBALS['Data_List_toUnfoldable'])($dictUnfoldable_0)))($GLOBALS['Data_Set_toList']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Set_toUnfoldable1
$GLOBALS['Data_Set_toUnfoldable1'] = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])(($GLOBALS['Data_List_toUnfoldable'])($GLOBALS['Data_Unfoldable_unfoldableArray'])))($GLOBALS['Data_Set_toList']);

// Data_Set_size
$GLOBALS['Data_Set_size'] = $GLOBALS['Data_Map_Internal_size'];

// Data_Set_singleton
$GLOBALS['Data_Set_singleton'] = function($a_0 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data6("Node", 1, 1, $a_0, $GLOBALS['Data_Unit_unit'], new Phpurs_Data0("Leaf"), new Phpurs_Data0("Leaf"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Set_showSet
$GLOBALS['Data_Set_showSet'] = function($dictShow_0 = null) {
  $__num = \func_num_args();
  $__res = ["show" => function($s_1 = null) use ($dictShow_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Semigroup_semigroupString'])['append'])("(fromFoldable "))(((($GLOBALS['Data_Semigroup_semigroupString'])['append'])((($GLOBALS['Data_Show_showArrayImpl'])(($dictShow_0)['show']))(($GLOBALS['Data_Set_toUnfoldable1'])($s_1))))(")"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Set_semigroupSet
$GLOBALS['Data_Set_semigroupSet'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $compare_1_0 = ($dictOrd_0)['compare'];
  $__res = ["append" => (function() use ($compare_1_0) {
  $__fn = function($m1_2 = null, $m2_3 = null) use ($compare_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Map_Internal_unsafeUnionWith'])($compare_1_0, $GLOBALS['Data_Function_const'], $m1_2, $m2_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Set_member
$GLOBALS['Data_Set_member'] = (function() {
  $__fn = function($dictOrd_0 = null, $k_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__2_0 = null;
  $go__2_0 = function($v_3 = null) use ($dictOrd_0, &$go__2_0, $k_1) {
  $__num = \func_num_args();
  $__tco_var_go__2_0_0_v_3 = $v_3;
  tco_loop_go__2_0_0:;
  $v_3 = $__tco_var_go__2_0_0_v_3;
  $__t0 = null;;
  if ((is_object($v_3) && (($v_3)->{'tag'} === "Leaf"))) {
$__t0 = false;
goto end_branch_0;;
};
  if ((is_object($v_3) && (($v_3)->{'tag'} === "Node"))) {
$v1_4_1 = ((($dictOrd_0)['compare'])($k_1))(($v_3)->{'value2'});
$__t2 = null;;
if ((is_object($v1_4_1) && (($v1_4_1)->{'tag'} === "LT"))) {
$__tco_3 = ($v_3)->{'value4'};
$__tco_var_go__2_0_0_v_3 = $__tco_3;
goto tco_loop_go__2_0_0;;
$__t2 = null;
goto end_branch_2;;
};
if ((is_object($v1_4_1) && (($v1_4_1)->{'tag'} === "GT"))) {
$__tco_4 = ($v_3)->{'value5'};
$__tco_var_go__2_0_0_v_3 = $__tco_4;
goto tco_loop_go__2_0_0;;
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
  $__res = $go__2_0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Set_isEmpty
$GLOBALS['Data_Set_isEmpty'] = $GLOBALS['Data_Map_Internal_isEmpty'];

// Data_Set_intersection
$GLOBALS['Data_Set_intersection'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $compare_1_0 = ($dictOrd_0)['compare'];
  $__res = (function() use ($compare_1_0) {
  $__fn = function($m1_2 = null, $m2_3 = null) use ($compare_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Map_Internal_unsafeIntersectionWith'])($compare_1_0, $GLOBALS['Data_Function_const'], $m1_2, $m2_3);
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

// Data_Set_insert
$GLOBALS['Data_Set_insert'] = (function() {
  $__fn = function($dictOrd_0 = null, $a_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($GLOBALS['Data_Map_Internal_insert'])($dictOrd_0))($a_1))($GLOBALS['Data_Unit_unit']))($v_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_Set_fromMap
$GLOBALS['Data_Set_fromMap'] = $GLOBALS['Data_Set_Set'];

// Data_Set_foldableSet
$GLOBALS['Data_Set_foldableSet'] = ["foldMap" => function($dictMonoid_0 = null) {
  $__num = \func_num_args();
  $foldMap1_1_0 = (($GLOBALS['Data_List_Types_foldableList'])['foldMap'])($dictMonoid_0);
  $__res = function($f_2 = null) use ($foldMap1_1_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])(($foldMap1_1_0)($f_2)))($GLOBALS['Data_Set_toList']);
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
  $__res = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])(((($GLOBALS['Data_List_Types_foldableList'])['foldl'])($f_0))($x_1)))($GLOBALS['Data_Set_toList']);
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
  $__res = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])(((($GLOBALS['Data_List_Types_foldableList'])['foldr'])($f_0))($x_1)))($GLOBALS['Data_Set_toList']);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];

// Data_Set_findMin
$GLOBALS['Data_Set_findMin'] = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Maybe_functorMaybe'])['map'])(function($v1_1 = null) {
  $__num = \func_num_args();
  $__res = ($v1_1)['key'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($GLOBALS['Data_Map_Internal_findMin'])($v_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Set_findMax
$GLOBALS['Data_Set_findMax'] = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Maybe_functorMaybe'])['map'])(function($v1_1 = null) {
  $__num = \func_num_args();
  $__res = ($v1_1)['key'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($GLOBALS['Data_Map_Internal_findMax'])($v_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Set_filter
$GLOBALS['Data_Set_filter'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Map_Internal_filterKeys'])($dictOrd_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Set_eqSet
$GLOBALS['Data_Set_eqSet'] = function($dictEq_0 = null) {
  $__num = \func_num_args();
  $__res = ["eq" => (function() use ($dictEq_0) {
  $__fn = function($v_1 = null, $v1_2 = null) use ($dictEq_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((($GLOBALS['Data_Map_Internal_eqMap'])($dictEq_0))($GLOBALS['Data_Eq_eqUnit']))['eq'])($v_1))($v1_2);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Set_ordSet
$GLOBALS['Data_Set_ordSet'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($dictOrd_0)['Eq0'])(null);
  $eqSet1_2_1 = ["eq" => (function() use ($__local_var_1_0) {
  $__fn = function($v_2 = null, $v1_3 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((($GLOBALS['Data_Map_Internal_eqMap'])($__local_var_1_0))($GLOBALS['Data_Eq_eqUnit']))['eq'])($v_2))($v1_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $__res = ["compare" => (function() use ($dictOrd_0) {
  $__fn = function($s1_3 = null, $s2_4 = null) use ($dictOrd_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($GLOBALS['Data_List_Types_ordList'])($dictOrd_0))['compare'])(($GLOBALS['Data_Map_Internal_keys'])($s1_3)))(($GLOBALS['Data_Map_Internal_keys'])($s2_4));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($_dollar__unused_3 = null) use ($eqSet1_2_1) {
  $__num = \func_num_args();
  $__res = $eqSet1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Set_eq1Set
$GLOBALS['Data_Set_eq1Set'] = ["eq1" => (function() {
  $__fn = function($dictEq_0 = null, $v_1 = null, $v1_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((($GLOBALS['Data_Map_Internal_eqMap'])($dictEq_0))($GLOBALS['Data_Eq_eqUnit']))['eq'])($v_1))($v1_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()];

// Data_Set_ord1Set
$GLOBALS['Data_Set_ord1Set'] = ["compare1" => function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Set_ordSet'])($dictOrd_0))['compare'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq10" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Set_eq1Set'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Set_empty
$GLOBALS['Data_Set_empty'] = new Phpurs_Data0("Leaf");

// Data_Set_fromFoldable
$GLOBALS['Data_Set_fromFoldable'] = (function() {
  $__fn = function($dictFoldable_0 = null, $dictOrd_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictFoldable_0)['foldl'])((function() use ($dictOrd_1) {
  $__fn = function($m_2 = null, $a_3 = null) use ($dictOrd_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($GLOBALS['Data_Map_Internal_insert'])($dictOrd_1))($a_3))($GLOBALS['Data_Unit_unit']))($m_2);
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
})();

// Data_Set_map
$GLOBALS['Data_Set_map'] = (function() {
  $__fn = function($dictOrd_0 = null, $f_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_Set_foldableSet'])['foldl'])((function() use ($dictOrd_0, $f_1) {
  $__fn = function($m_2 = null, $a_3 = null) use ($dictOrd_0, $f_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($GLOBALS['Data_Map_Internal_insert'])($dictOrd_0))(($f_1)($a_3)))($GLOBALS['Data_Unit_unit']))($m_2);
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
})();

// Data_Set_mapMaybe
$GLOBALS['Data_Set_mapMaybe'] = (function() {
  $__fn = function($dictOrd_0 = null, $f_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_Set_foldableSet'])['foldr'])((function() use ($dictOrd_0, $f_1) {
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
$__t1 = (((($GLOBALS['Data_Map_Internal_insert'])($dictOrd_0))(($__local_var_4_0)->{'value0'}))($GLOBALS['Data_Unit_unit']))($acc_3);
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
})()))(new Phpurs_Data0("Leaf"));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Set_monoidSet
$GLOBALS['Data_Set_monoidSet'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $compare_1_0 = ($dictOrd_0)['compare'];
  $semigroupSet1_1_0 = ["append" => (function() use ($compare_1_0) {
  $__fn = function($m1_2 = null, $m2_3 = null) use ($compare_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Map_Internal_unsafeUnionWith'])($compare_1_0, $GLOBALS['Data_Function_const'], $m1_2, $m2_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $__res = ["mempty" => new Phpurs_Data0("Leaf"), "Semigroup0" => function($_dollar__unused_2 = null) use ($semigroupSet1_1_0) {
  $__num = \func_num_args();
  $__res = $semigroupSet1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Set_unions
$GLOBALS['Data_Set_unions'] = (function() {
  $__fn = function($dictFoldable_0 = null, $dictOrd_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $compare_2_0 = ($dictOrd_1)['compare'];
  $__res = ((($dictFoldable_0)['foldl'])((function() use ($compare_2_0) {
  $__fn = function($m1_3 = null, $m2_4 = null) use ($compare_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Map_Internal_unsafeUnionWith'])($compare_2_0, $GLOBALS['Data_Function_const'], $m1_3, $m2_4);
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
})();

// Data_Set_difference
$GLOBALS['Data_Set_difference'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $compare_1_0 = ($dictOrd_0)['compare'];
  $__res = (function() use ($compare_1_0) {
  $__fn = function($m1_2 = null, $m2_3 = null) use ($compare_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Map_Internal_unsafeDifference'])($compare_1_0, $m1_2, $m2_3);
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

// Data_Set_subset
$GLOBALS['Data_Set_subset'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $compare_1_0 = ($dictOrd_0)['compare'];
  $__res = (function() use ($compare_1_0) {
  $__fn = function($s1_2 = null, $s2_3 = null) use ($compare_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (is_object(($GLOBALS['Data_Map_Internal_unsafeDifference'])($compare_1_0, $s1_2, $s2_3)) && ((($GLOBALS['Data_Map_Internal_unsafeDifference'])($compare_1_0, $s1_2, $s2_3))->{'tag'} === "Leaf"));
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

// Data_Set_properSubset
$GLOBALS['Data_Set_properSubset'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $compare_1_0 = ($dictOrd_0)['compare'];
  $__res = (function() use ($compare_1_0) {
  $__fn = function($s1_2 = null, $s2_3 = null) use ($compare_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t1 = null;;
  if ((is_object($s1_2) && (($s1_2)->{'tag'} === "Leaf"))) {
$__t1 = 0;
goto end_branch_1;;
};
  if ((is_object($s1_2) && (($s1_2)->{'tag'} === "Node"))) {
$__t1 = ($s1_2)->{'value1'};
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__t2 = null;;
  if ((is_object($s2_3) && (($s2_3)->{'tag'} === "Leaf"))) {
$__t2 = 0;
goto end_branch_2;;
};
  if ((is_object($s2_3) && (($s2_3)->{'tag'} === "Node"))) {
$__t2 = ($s2_3)->{'value1'};
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = ((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])['conj'])(((($GLOBALS['Data_Eq_eqBoolean'])['eq'])(((($GLOBALS['Data_Eq_eqInt'])['eq'])($__t1))($__t2)))(false)))((is_object(($GLOBALS['Data_Map_Internal_unsafeDifference'])($compare_1_0, $s1_2, $s2_3)) && ((($GLOBALS['Data_Map_Internal_unsafeDifference'])($compare_1_0, $s1_2, $s2_3))->{'tag'} === "Leaf")));
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

// Data_Set_delete
$GLOBALS['Data_Set_delete'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Map_Internal_delete'])($dictOrd_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Set_checkValid
$GLOBALS['Data_Set_checkValid'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Map_Internal_checkValid'])($dictOrd_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Set_catMaybes
$GLOBALS['Data_Set_catMaybes'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Set_mapMaybe'])($dictOrd_0))(($GLOBALS['Control_Category_categoryFn'])['identity']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

