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




// Data_Set_identity
function majData_majSet_identity($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_identity';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Set_identity'] = __NAMESPACE__ . '\\majData_majSet_identity';

// Data_Set_Set
function majData_majSet_majSet($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_majSet';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Set_Set'] = __NAMESPACE__ . '\\majData_majSet_majSet';

// Data_Set_union
function majData_majSet_union($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_union';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Set_union'] = __NAMESPACE__ . '\\majData_majSet_union';

// Data_Set_toggle
function majData_majSet_toggle($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_toggle';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $alter_1_0 = ($GLOBALS['Data_Map_Internal_alter'])($dictOrd_0);
  $__res = (function() use ($alter_1_0) {
  $__fn = function($a_2 = null, $v_3 = null) use ($alter_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($alter_1_0)(function($v2_4 = null) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v2_4 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t1 = new \Data\Maybe\Data_Maybe_Just($GLOBALS['Data_Unit_unit']);
goto end_branch_1;;
};
  if ($v2_4 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $a_2, $v_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Set_toggle'] = __NAMESPACE__ . '\\majData_majSet_toggle';

// Data_Set_toMap
function majData_majSet_tomajMap($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_tomajMap';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $v_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Set_toMap'] = __NAMESPACE__ . '\\majData_majSet_tomajMap';

// Data_Set_toList
function majData_majSet_tomajList($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_tomajList';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Map_Internal_keys'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Set_toList'] = __NAMESPACE__ . '\\majData_majSet_tomajList';

// Data_Set_toUnfoldable
function majData_majSet_tomajUnfoldable($dictUnfoldable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_tomajUnfoldable';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_List_toUnfoldable'])($dictUnfoldable_0), $GLOBALS['Data_Set_toList']);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Set_toUnfoldable'] = __NAMESPACE__ . '\\majData_majSet_tomajUnfoldable';

// Data_Set_toUnfoldable1
$GLOBALS['Data_Set_toUnfoldable1'] = ($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_List_toUnfoldable'])($GLOBALS['Data_Unfoldable_unfoldableArray']), $GLOBALS['Data_Set_toList']);

// Data_Set_size_closure
$GLOBALS['Data_Set_size_closure'] = $GLOBALS['Data_Map_Internal_size'];

// Data_Set_size
function majData_majSet_size($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_size';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Set_size_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Set_size'] = __NAMESPACE__ . '\\majData_majSet_size';

// Data_Set_singleton
function majData_majSet_singleton($a_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_singleton';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = new \Data\Map\Internal\Data_Map_Internal_Node(1, 1, $a_0, $GLOBALS['Data_Unit_unit'], new \Data\Map\Internal\Data_Map_Internal_Leaf(), new \Data\Map\Internal\Data_Map_Internal_Leaf());
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Set_singleton'] = __NAMESPACE__ . '\\majData_majSet_singleton';

// Data_Set_showSet
function majData_majSet_showmajSet($dictShow_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_showmajSet';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ["show" => function($s_1 = null) use ($dictShow_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Semigroup_semigroupString'])['append'])("(fromFoldable ", (($GLOBALS['Data_Semigroup_semigroupString'])['append'])(($GLOBALS['Data_Show_showArrayImpl'])(($dictShow_0)['show'], ($GLOBALS['Data_Set_toUnfoldable1'])($s_1)), ")"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Set_showSet'] = __NAMESPACE__ . '\\majData_majSet_showmajSet';

// Data_Set_semigroupSet
function majData_majSet_semigroupmajSet($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_semigroupmajSet';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Set_semigroupSet'] = __NAMESPACE__ . '\\majData_majSet_semigroupmajSet';

// Data_Set_member
function majData_majSet_member($dictOrd_0, $k_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_member';
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
  if ($v_3 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t0 = false;
goto end_branch_0;;
};
  if ($v_3 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$v1_4_1 = (($dictOrd_0)['compare'])($k_1, ($v_3)->{'value2'});
$__t2 = null;;
if ($v1_4_1 instanceof \Data\Ordering\Data_Ordering_LT) {
$__tco_3 = ($v_3)->{'value4'};
$__tco_var_go__2_0_0_v_3 = $__tco_3;
goto tco_loop_go__2_0_0;;
$__t2 = null;
goto end_branch_2;;
};
if ($v1_4_1 instanceof \Data\Ordering\Data_Ordering_GT) {
$__tco_4 = ($v_3)->{'value5'};
$__tco_var_go__2_0_0_v_3 = $__tco_4;
goto tco_loop_go__2_0_0;;
$__t2 = null;
goto end_branch_2;;
};
if ($v1_4_1 instanceof \Data\Ordering\Data_Ordering_EQ) {
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
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Set_member'] = __NAMESPACE__ . '\\majData_majSet_member';

// Data_Set_isEmpty_closure
$GLOBALS['Data_Set_isEmpty_closure'] = $GLOBALS['Data_Map_Internal_isEmpty'];

// Data_Set_isEmpty
function majData_majSet_ismajEmpty($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_ismajEmpty';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Set_isEmpty_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Set_isEmpty'] = __NAMESPACE__ . '\\majData_majSet_ismajEmpty';

// Data_Set_intersection
function majData_majSet_intersection($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_intersection';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Set_intersection'] = __NAMESPACE__ . '\\majData_majSet_intersection';

// Data_Set_insert
function majData_majSet_insert($dictOrd_0, $a_1 = null, $v_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_insert';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($GLOBALS['Data_Map_Internal_insert'])($dictOrd_0, $a_1, $GLOBALS['Data_Unit_unit'], $v_2);
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Set_insert'] = __NAMESPACE__ . '\\majData_majSet_insert';

// Data_Set_fromMap_closure
$GLOBALS['Data_Set_fromMap_closure'] = $GLOBALS['Data_Set_Set'];

// Data_Set_fromMap
function majData_majSet_frommajMap($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_frommajMap';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Set_fromMap_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Set_fromMap'] = __NAMESPACE__ . '\\majData_majSet_frommajMap';

// Data_Set_foldableSet
$GLOBALS['Data_Set_foldableSet'] = ["foldMap" => function($dictMonoid_0 = null) {
  $__num = \func_num_args();
  $foldMap1_1_0 = (($GLOBALS['Data_List_Types_foldableList'])['foldMap'])($dictMonoid_0);
  $__res = function($f_2 = null) use ($foldMap1_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Semigroupoid_composeImpl'])(($foldMap1_1_0)($f_2), $GLOBALS['Data_Set_toList']);
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
  $__res = ($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_List_Types_foldableList'])['foldl'])($f_0, $x_1), $GLOBALS['Data_Set_toList']);
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
  $__res = ($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_List_Types_foldableList'])['foldr'])($f_0, $x_1), $GLOBALS['Data_Set_toList']);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];

// Data_Set_findMin
function majData_majSet_findmajMin($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_findmajMin';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Data_Maybe_functorMaybe'])['map'])(function($v1_1 = null) {
  $__num = \func_num_args();
  $__res = ($v1_1)['key'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($GLOBALS['Data_Map_Internal_findMin'])($v_0));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Set_findMin'] = __NAMESPACE__ . '\\majData_majSet_findmajMin';

// Data_Set_findMax
function majData_majSet_findmajMax($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_findmajMax';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Data_Maybe_functorMaybe'])['map'])(function($v1_1 = null) {
  $__num = \func_num_args();
  $__res = ($v1_1)['key'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($GLOBALS['Data_Map_Internal_findMax'])($v_0));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Set_findMax'] = __NAMESPACE__ . '\\majData_majSet_findmajMax';

// Data_Set_filter
function majData_majSet_filter($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_filter';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Map_Internal_filterKeys'])($dictOrd_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Set_filter'] = __NAMESPACE__ . '\\majData_majSet_filter';

// Data_Set_eqSet
function majData_majSet_eqmajSet($dictEq_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_eqmajSet';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ["eq" => (function() use ($dictEq_0) {
  $__fn = function($v_1 = null, $v1_2 = null) use ($dictEq_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_Map_Internal_eqMap'])($dictEq_0, $GLOBALS['Data_Eq_eqUnit']))['eq'])($v_1, $v1_2);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Set_eqSet'] = __NAMESPACE__ . '\\majData_majSet_eqmajSet';

// Data_Set_ordSet
function majData_majSet_ordmajSet($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_ordmajSet';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictOrd_0)['Eq0'])(null);
  $eqSet1_2_1 = ["eq" => (function() use ($__local_var_1_0) {
  $__fn = function($v_2 = null, $v1_3 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_Map_Internal_eqMap'])($__local_var_1_0, $GLOBALS['Data_Eq_eqUnit']))['eq'])($v_2, $v1_3);
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
  $__res = ((($GLOBALS['Data_List_Types_ordList'])($dictOrd_0))['compare'])(($GLOBALS['Data_Map_Internal_keys'])($s1_3), ($GLOBALS['Data_Map_Internal_keys'])($s2_4));
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Set_ordSet'] = __NAMESPACE__ . '\\majData_majSet_ordmajSet';

// Data_Set_eq1Set
$GLOBALS['Data_Set_eq1Set'] = ["eq1" => (function() {
  $__fn = function($dictEq_0 = null, $v_1 = null, $v1_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($GLOBALS['Data_Map_Internal_eqMap'])($dictEq_0, $GLOBALS['Data_Eq_eqUnit']))['eq'])($v_1, $v1_2);
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
$GLOBALS['Data_Set_empty'] = new \Data\Map\Internal\Data_Map_Internal_Leaf();

// Data_Set_fromFoldable
function majData_majSet_frommajFoldable($dictFoldable_0, $dictOrd_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_frommajFoldable';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($dictFoldable_0)['foldl'])((function() use ($dictOrd_1) {
  $__fn = function($m_2 = null, $a_3 = null) use ($dictOrd_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Map_Internal_insert'])($dictOrd_1, $a_3, $GLOBALS['Data_Unit_unit'], $m_2);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), new \Data\Map\Internal\Data_Map_Internal_Leaf());
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Set_fromFoldable'] = __NAMESPACE__ . '\\majData_majSet_frommajFoldable';

// Data_Set_map
function majData_majSet_map($dictOrd_0, $f_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_map';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_Set_foldableSet'])['foldl'])((function() use ($dictOrd_0, $f_1) {
  $__fn = function($m_2 = null, $a_3 = null) use ($dictOrd_0, $f_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Map_Internal_insert'])($dictOrd_0, ($f_1)($a_3), $GLOBALS['Data_Unit_unit'], $m_2);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), new \Data\Map\Internal\Data_Map_Internal_Leaf());
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Set_map'] = __NAMESPACE__ . '\\majData_majSet_map';

// Data_Set_mapMaybe
function majData_majSet_mapmajMaybe($dictOrd_0, $f_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_mapmajMaybe';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_Set_foldableSet'])['foldr'])((function() use ($dictOrd_0, $f_1) {
  $__fn = function($a_2 = null, $acc_3 = null) use ($dictOrd_0, $f_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_4_0 = ($f_1)($a_2);
  $__t1 = null;;
  if ($__local_var_4_0 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t1 = $acc_3;
goto end_branch_1;;
};
  if ($__local_var_4_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = ($GLOBALS['Data_Map_Internal_insert'])($dictOrd_0, ($__local_var_4_0)->{'value0'}, $GLOBALS['Data_Unit_unit'], $acc_3);
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
})(), new \Data\Map\Internal\Data_Map_Internal_Leaf());
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Set_mapMaybe'] = __NAMESPACE__ . '\\majData_majSet_mapmajMaybe';

// Data_Set_monoidSet
function majData_majSet_monoidmajSet($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_monoidmajSet';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
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
  $__res = ["mempty" => new \Data\Map\Internal\Data_Map_Internal_Leaf(), "Semigroup0" => function($_dollar__unused_2 = null) use ($semigroupSet1_1_0) {
  $__num = \func_num_args();
  $__res = $semigroupSet1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Set_monoidSet'] = __NAMESPACE__ . '\\majData_majSet_monoidmajSet';

// Data_Set_unions
function majData_majSet_unions($dictFoldable_0, $dictOrd_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_unions';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $compare_2_0 = ($dictOrd_1)['compare'];
  $__res = (($dictFoldable_0)['foldl'])((function() use ($compare_2_0) {
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
})(), new \Data\Map\Internal\Data_Map_Internal_Leaf());
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Set_unions'] = __NAMESPACE__ . '\\majData_majSet_unions';

// Data_Set_difference
function majData_majSet_difference($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_difference';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Set_difference'] = __NAMESPACE__ . '\\majData_majSet_difference';

// Data_Set_subset
function majData_majSet_subset($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_subset';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $compare_1_0 = ($dictOrd_0)['compare'];
  $__res = (function() use ($compare_1_0) {
  $__fn = function($s1_2 = null, $s2_3 = null) use ($compare_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Map_Internal_unsafeDifference'])($compare_1_0, $s1_2, $s2_3) instanceof \Data\Map\Internal\Data_Map_Internal_Leaf;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Set_subset'] = __NAMESPACE__ . '\\majData_majSet_subset';

// Data_Set_properSubset
function majData_majSet_propermajSubset($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_propermajSubset';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $compare_1_0 = ($dictOrd_0)['compare'];
  $__res = (function() use ($compare_1_0) {
  $__fn = function($s1_2 = null, $s2_3 = null) use ($compare_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t1 = null;;
  if ($s1_2 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t1 = 0;
goto end_branch_1;;
};
  if ($s1_2 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t1 = ($s1_2)->{'value1'};
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__t2 = null;;
  if ($s2_3 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t2 = 0;
goto end_branch_2;;
};
  if ($s2_3 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t2 = ($s2_3)->{'value1'};
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = (($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])['conj'])((($GLOBALS['Data_Eq_eqBoolean'])['eq'])(($__t1 === $__t2), false), ($GLOBALS['Data_Map_Internal_unsafeDifference'])($compare_1_0, $s1_2, $s2_3) instanceof \Data\Map\Internal\Data_Map_Internal_Leaf);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Set_properSubset'] = __NAMESPACE__ . '\\majData_majSet_propermajSubset';

// Data_Set_delete
function majData_majSet_delete($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_delete';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Map_Internal_delete'])($dictOrd_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Set_delete'] = __NAMESPACE__ . '\\majData_majSet_delete';

// Data_Set_checkValid
function majData_majSet_checkmajValid($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_checkmajValid';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Map_Internal_checkValid'])($dictOrd_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Set_checkValid'] = __NAMESPACE__ . '\\majData_majSet_checkmajValid';

// Data_Set_catMaybes
function majData_majSet_catmajMaybes($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_catmajMaybes';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Set_mapMaybe'])($dictOrd_0, $GLOBALS['Data_Set_identity']);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Set_catMaybes'] = __NAMESPACE__ . '\\majData_majSet_catmajMaybes';

