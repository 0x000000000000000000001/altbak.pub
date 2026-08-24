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
if (!\function_exists(__NAMESPACE__ . '\\phpurs_execute_effect')) {
  function phpurs_execute_effect($val) {
    if (\is_callable($val)) {
      return $val($GLOBALS['Data_Unit_unit']);
    }
    return $val;
  }
}

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };




// Data_Set_union
function majData_majSet_union($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_union';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $compare_1_0 = ($dictOrd_0)->{'compare'};
  $__res = function($m1_2) use ($compare_1_0) {
  $__num = \func_num_args();
  $__res = function($m2_3) use ($compare_1_0, $m1_2) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Map_Internal_unsafeUnionWith'])($compare_1_0, $GLOBALS['Data_Function_const'], $m1_2, $m2_3);
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Set_union'] = __NAMESPACE__ . '\\majData_majSet_union';

// Data_Set_toggle
function majData_majSet_toggle($dictOrd_0, $a_1 = null, $v_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_toggle';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $v_3_0 = ($GLOBALS['Data_Map_Internal_unsafeSplit'])(($dictOrd_0)->{'compare'}, $a_1, $v_2);
  $__t1 = null;;
  if (($v_3_0)->{'value0'} instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t1 = ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])($a_1, $GLOBALS['Data_Unit_unit'], ($v_3_0)->{'value1'}, ($v_3_0)->{'value2'});
goto end_branch_1;;
};
  if (($v_3_0)->{'value0'} instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = ($GLOBALS['Data_Map_Internal_unsafeJoinNodes'])(($v_3_0)->{'value1'}, ($v_3_0)->{'value2'});
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
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
  $go__go_1_0 = null;
  $go__go_1_0 = (function() use (&$go__go_1_0) {
  $__fn = function($m_prime__2, $__local_var_3 = null) use (&$go__go_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t1 = null;;
  if ($m_prime__2 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t1 = $__local_var_3;
goto end_branch_1;;
};
  if ($m_prime__2 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t1 = ($go__go_1_0)(($m_prime__2)->{'value4'}, new \Data\List\Types\Data_List_Types_Cons(($m_prime__2)->{'value2'}, ($go__go_1_0)(($m_prime__2)->{'value5'}, $__local_var_3)));
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
  $__res = ($go__go_1_0)($v_0, new \Data\List\Types\Data_List_Types_Nil());
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
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($dictUnfoldable_0)->{'unfoldr'})(function($xs_1) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($xs_1 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_0;;
};
  if ($xs_1 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__t0 = new \Data\Maybe\Data_Maybe_Just(new \Data\Tuple\Data_Tuple_Tuple(($xs_1)->{'value0'}, ($xs_1)->{'value1'}));
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))($GLOBALS['Data_Set_toList']);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Set_toUnfoldable'] = __NAMESPACE__ . '\\majData_majSet_tomajUnfoldable';

// Data_Set_size_closure
$GLOBALS['Data_Set_size_closure'] = $GLOBALS['Data_Map_Internal_size'];

// Data_Set_size
function majData_majSet_size($v_0): int|\Closure {
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
  $showArray_1_0 = (object)["show" => ($GLOBALS['Data_Show_showArrayImpl'])(($dictShow_0)->{'show'})];
  $__res = (object)["show" => function($s_2) use ($showArray_1_0) {
  $__num = \func_num_args();
  $__res = (("(fromFoldable " . (($showArray_1_0)->{'show'})(\Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(((((($GLOBALS['Data_Unfoldable_unfoldrArrayImpl'])($GLOBALS['Data_Maybe_isNothing']))(function($v_3) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = ($v_3)->{'value0'};
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Tuple_fst']))($GLOBALS['Data_Tuple_snd']))(function($xs_3) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ($xs_3 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t2 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_2;;
};
  if ($xs_3 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__t2 = new \Data\Maybe\Data_Maybe_Just(new \Data\Tuple\Data_Tuple_Tuple(($xs_3)->{'value0'}, ($xs_3)->{'value1'}));
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), $GLOBALS['Data_Set_toList'], $s_2))) . ")");
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
  $compare_1_0 = ($dictOrd_0)->{'compare'};
  $__res = (object)["append" => function($m1_2) use ($compare_1_0) {
  $__num = \func_num_args();
  $__res = function($m2_3) use ($compare_1_0, $m1_2) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Map_Internal_unsafeUnionWith'])($compare_1_0, $GLOBALS['Data_Function_const'], $m1_2, $m2_3);
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
  $go__go_2_0 = null;
  $go__go_2_0 = function($v_3) use ($dictOrd_0, &$go__go_2_0, $k_1) {
  $__num = \func_num_args();
  $__tco_var_go__go_2_0_0_v_3 = $v_3;
  tco_loop_go__go_2_0_0:;
  $v_3 = $__tco_var_go__go_2_0_0_v_3;
  $__t0 = null;;
  if ($v_3 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t0 = false;
goto end_branch_0;;
};
  if ($v_3 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$v1_4_1 = ((($dictOrd_0)->{'compare'})($k_1))(($v_3)->{'value2'});
$__t2 = null;;
if ($v1_4_1 instanceof \Data\Ordering\Data_Ordering_LT) {
$__tco_3 = ($v_3)->{'value4'};
$__tco_var_go__go_2_0_0_v_3 = $__tco_3;
goto tco_loop_go__go_2_0_0;;
$__t2 = null;
goto end_branch_2;;
};
if ($v1_4_1 instanceof \Data\Ordering\Data_Ordering_GT) {
$__tco_4 = ($v_3)->{'value5'};
$__tco_var_go__go_2_0_0_v_3 = $__tco_4;
goto tco_loop_go__go_2_0_0;;
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
  $__res = $go__go_2_0;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Set_member'] = __NAMESPACE__ . '\\majData_majSet_member';

// Data_Set_isEmpty_closure
$GLOBALS['Data_Set_isEmpty_closure'] = $GLOBALS['Data_Map_Internal_isEmpty'];

// Data_Set_isEmpty
function majData_majSet_ismajEmpty($v_0): bool|\Closure {
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
  $compare_1_0 = ($dictOrd_0)->{'compare'};
  $__res = function($m1_2) use ($compare_1_0) {
  $__num = \func_num_args();
  $__res = function($m2_3) use ($compare_1_0, $m1_2) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Map_Internal_unsafeIntersectionWith'])($compare_1_0, $GLOBALS['Data_Function_const'], $m1_2, $m2_3);
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
  $go__go_3_0 = null;
  $go__go_3_0 = function($v1_4) use ($a_1, $dictOrd_0, &$go__go_3_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v1_4 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t1 = new \Data\Map\Internal\Data_Map_Internal_Node(1, 1, $a_1, $GLOBALS['Data_Unit_unit'], new \Data\Map\Internal\Data_Map_Internal_Leaf(), new \Data\Map\Internal\Data_Map_Internal_Leaf());
goto end_branch_1;;
};
  if ($v1_4 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$v2_5_2 = ((($dictOrd_0)->{'compare'})($a_1))(($v1_4)->{'value2'});
$__t3 = null;;
if ($v2_5_2 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t3 = ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($v1_4)->{'value2'}, ($v1_4)->{'value3'}, ($go__go_3_0)(($v1_4)->{'value4'}), ($v1_4)->{'value5'});
goto end_branch_3;;
};
if ($v2_5_2 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t3 = ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($v1_4)->{'value2'}, ($v1_4)->{'value3'}, ($v1_4)->{'value4'}, ($go__go_3_0)(($v1_4)->{'value5'}));
goto end_branch_3;;
};
if ($v2_5_2 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t3 = new \Data\Map\Internal\Data_Map_Internal_Node(($v1_4)->{'value0'}, ($v1_4)->{'value1'}, $a_1, $GLOBALS['Data_Unit_unit'], ($v1_4)->{'value4'}, ($v1_4)->{'value5'});
goto end_branch_3;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
end_branch_3:;
$__t1 = $__t3;
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = ($go__go_3_0)($v_2);
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Set_insert'] = __NAMESPACE__ . '\\majData_majSet_insert';

// Data_Set_fromMap
function majData_majSet_frommajMap($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_frommajMap';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Set_fromMap'] = __NAMESPACE__ . '\\majData_majSet_frommajMap';

// Data_Set_foldableSet
$GLOBALS['Data_Set_foldableSet'] = (object)["foldMap" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $__res = function($f_1) use ($dictMonoid_0) {
  $__num = \func_num_args();
  $Semigroup0_2_0 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $go__go_3_1 = null;
  $go__go_3_1 = (function() use ($Semigroup0_2_0, $f_1, &$go__go_3_1) {
  $__fn = function($b_4, $v_5 = null) use ($Semigroup0_2_0, $f_1, &$go__go_3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_3_1_1_b_4 = $b_4;
  $__tco_var_go__go_3_1_1_v_5 = $v_5;
  tco_loop_go__go_3_1_1:;
  $b_4 = $__tco_var_go__go_3_1_1_b_4;
  $v_5 = $__tco_var_go__go_3_1_1_v_5;
  $__t1 = null;;
  if ($v_5 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = $b_4;
goto end_branch_1;;
};
  if ($v_5 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_2 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl((($Semigroup0_2_0)->{'append'})($b_4), $f_1, ($v_5)->{'value0'});
$__tco_3 = ($v_5)->{'value1'};
$__tco_var_go__go_3_1_1_b_4 = $__tco_2;
$__tco_var_go__go_3_1_1_v_5 = $__tco_3;
goto tco_loop_go__go_3_1_1;;
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
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($go__go_3_1)(($dictMonoid_0)->{'mempty'})))($GLOBALS['Data_Set_toList']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldl" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($x_1) use ($f_0) {
  $__num = \func_num_args();
  $go__go_2_2 = null;
  $go__go_2_2 = (function() use ($f_0, &$go__go_2_2) {
  $__fn = function($b_3, $v_4 = null) use ($f_0, &$go__go_2_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_2_2_2_b_3 = $b_3;
  $__tco_var_go__go_2_2_2_v_4 = $v_4;
  tco_loop_go__go_2_2_2:;
  $b_3 = $__tco_var_go__go_2_2_2_b_3;
  $v_4 = $__tco_var_go__go_2_2_2_v_4;
  $__t2 = null;;
  if ($v_4 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t2 = $b_3;
goto end_branch_2;;
};
  if ($v_4 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_3 = (($f_0)($b_3))(($v_4)->{'value0'});
$__tco_4 = ($v_4)->{'value1'};
$__tco_var_go__go_2_2_2_b_3 = $__tco_3;
$__tco_var_go__go_2_2_2_v_4 = $__tco_4;
goto tco_loop_go__go_2_2_2;;
$__t2 = null;
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($go__go_2_2)($x_1)))($GLOBALS['Data_Set_toList']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldr" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($x_1) use ($f_0) {
  $__num = \func_num_args();
  $go__go_2_3 = null;
  $go__go_2_3 = (function() use ($f_0, &$go__go_2_3) {
  $__fn = function($b_3, $v_4 = null) use ($f_0, &$go__go_2_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_2_3_3_b_3 = $b_3;
  $__tco_var_go__go_2_3_3_v_4 = $v_4;
  tco_loop_go__go_2_3_3:;
  $b_3 = $__tco_var_go__go_2_3_3_b_3;
  $v_4 = $__tco_var_go__go_2_3_3_v_4;
  $__t3 = null;;
  if ($v_4 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t3 = $b_3;
goto end_branch_3;;
};
  if ($v_4 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_4 = (($f_0)(($v_4)->{'value0'}))($b_3);
$__tco_5 = ($v_4)->{'value1'};
$__tco_var_go__go_2_3_3_b_3 = $__tco_4;
$__tco_var_go__go_2_3_3_v_4 = $__tco_5;
goto tco_loop_go__go_2_3_3;;
$__t3 = null;
goto end_branch_3;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t3 = null;
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $go__go_2_4 = null;
  $go__go_2_4 = (function() use (&$go__go_2_4) {
  $__fn = function($v_3, $v1_4 = null) use (&$go__go_2_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_2_4_4_v_3 = $v_3;
  $__tco_var_go__go_2_4_4_v1_4 = $v1_4;
  tco_loop_go__go_2_4_4:;
  $v_3 = $__tco_var_go__go_2_4_4_v_3;
  $v1_4 = $__tco_var_go__go_2_4_4_v1_4;
  $__t4 = null;;
  if ($v1_4 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t4 = $v_3;
goto end_branch_4;;
};
  if ($v1_4 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_5 = new \Data\List\Types\Data_List_Types_Cons(($v1_4)->{'value0'}, $v_3);
$__tco_6 = ($v1_4)->{'value1'};
$__tco_var_go__go_2_4_4_v_3 = $__tco_5;
$__tco_var_go__go_2_4_4_v1_4 = $__tco_6;
goto tco_loop_go__go_2_4_4;;
$__t4 = null;
goto end_branch_4;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t4 = null;
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Control_Semigroupoid_composeImpl'])(($go__go_2_3)($x_1)))(($go__go_2_4)(new \Data\List\Types\Data_List_Types_Nil()))))($GLOBALS['Data_Set_toList']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Set_findMin
function majData_majSet_findmajMin($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_findmajMin';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = \Data\Map\Internal\majData_majMap_majInternal_findmajMin($v_0);
  $__t1 = null;;
  if ($__local_var_1_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = new \Data\Maybe\Data_Maybe_Just((($__local_var_1_0)->{'value0'})->{'key'});
goto end_branch_1;;
};
  $__t1 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_1:;
  $__res = $__t1;
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
  $__local_var_1_0 = \Data\Map\Internal\majData_majMap_majInternal_findmajMax($v_0);
  $__t1 = null;;
  if ($__local_var_1_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = new \Data\Maybe\Data_Maybe_Just((($__local_var_1_0)->{'value0'})->{'key'});
goto end_branch_1;;
};
  $__t1 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Set_findMax'] = __NAMESPACE__ . '\\majData_majSet_findmajMax';

// Data_Set_filter
function majData_majSet_filter($dictOrd_0, $f_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_filter';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__go_2_0 = null;
  $go__go_2_0 = function($v_3) use ($f_1, &$go__go_2_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v_3 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t1 = new \Data\Map\Internal\Data_Map_Internal_Leaf();
goto end_branch_1;;
};
  if ($v_3 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t2 = null;;
if (($f_1)(($v_3)->{'value2'})) {
$__t2 = ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($v_3)->{'value2'}, ($v_3)->{'value3'}, ($go__go_2_0)(($v_3)->{'value4'}), ($go__go_2_0)(($v_3)->{'value5'}));
goto end_branch_2;;
};
$__t2 = ($GLOBALS['Data_Map_Internal_unsafeJoinNodes'])(($go__go_2_0)(($v_3)->{'value4'}), ($go__go_2_0)(($v_3)->{'value5'}));
end_branch_2:;
$__t1 = $__t2;
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = $go__go_2_0;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Set_filter'] = __NAMESPACE__ . '\\majData_majSet_filter';

// Data_Set_eqSet
function majData_majSet_eqmajSet($dictEq_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_eqmajSet';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $go__go_1_0 = null;
  $go__go_1_0 = function($a_2) use ($dictEq_0, &$go__go_1_0) {
  $__num = \func_num_args();
  $__res = function($b_3) use ($a_2, $dictEq_0, &$go__go_1_0) {
  $__num = \func_num_args();
  $v_4_1 = \Data\Map\Internal\majData_majMap_majInternal_stepmajAsc($a_2);
  $__t2 = null;;
  if ($v_4_1 instanceof \Data\Map\Internal\Data_Map_Internal_IterNext) {
$v2_5_3 = \Data\Map\Internal\majData_majMap_majInternal_stepmajAsc($b_3);
$__t2 = ($v2_5_3 instanceof \Data\Map\Internal\Data_Map_Internal_IterNext && (((($dictEq_0)->{'eq'})(($v_4_1)->{'value0'}))(($v2_5_3)->{'value0'}) && (($go__go_1_0)(($v_4_1)->{'value2'}))(($v2_5_3)->{'value2'})));
goto end_branch_2;;
};
  if ($v_4_1 instanceof \Data\Map\Internal\Data_Map_Internal_IterDone) {
$__t2 = true;
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $eqMapIter2_1_0 = (object)["eq" => $go__go_1_0];
  $eqMap_1_0 = (object)["eq" => function($xs_2) use ($eqMapIter2_1_0) {
  $__num = \func_num_args();
  $__res = function($ys_3) use ($eqMapIter2_1_0, $xs_2) {
  $__num = \func_num_args();
  $__t5 = null;;
  if ($xs_2 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t5 = $ys_3 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf;
goto end_branch_5;;
};
  if ($xs_2 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t5 = ($ys_3 instanceof \Data\Map\Internal\Data_Map_Internal_Node && ((($xs_2)->{'value1'} === ($ys_3)->{'value1'}) && ((($eqMapIter2_1_0)->{'eq'})(new \Data\Map\Internal\Data_Map_Internal_IterNode($xs_2, new \Data\Map\Internal\Data_Map_Internal_IterLeaf())))(new \Data\Map\Internal\Data_Map_Internal_IterNode($ys_3, new \Data\Map\Internal\Data_Map_Internal_IterLeaf()))));
goto end_branch_5;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t5 = null;
  end_branch_5:;
  $__res = $__t5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["eq" => function($v_2) use ($eqMap_1_0) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($eqMap_1_0, $v_2) {
  $__num = \func_num_args();
  $__res = ((($eqMap_1_0)->{'eq'})($v_2))($v1_3);
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
  $__local_var_1_0 = (($dictOrd_0)->{'Eq0'})(null);
  $eqList1_1_0 = (object)["eq" => function($xs_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($ys_3) use ($__local_var_1_0, $xs_2) {
  $__num = \func_num_args();
  $go__go_4_1 = null;
  $go__go_4_1 = function($v_5) use ($__local_var_1_0, &$go__go_4_1) {
  $__num = \func_num_args();
  $__res = function($v1_6) use ($__local_var_1_0, &$go__go_4_1, $v_5) {
  $__num = \func_num_args();
  $__res = function($v2_7) use ($__local_var_1_0, &$go__go_4_1, $v1_6, $v_5) {
  $__num = \func_num_args();
  $__t2 = null;;
  if (( ! $v2_7)) {
$__t2 = false;
goto end_branch_2;;
};
  if ($v_5 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t2 = ($v1_6 instanceof \Data\List\Types\Data_List_Types_Nil && $v2_7);
goto end_branch_2;;
};
  $__t2 = ($v_5 instanceof \Data\List\Types\Data_List_Types_Cons && ($v1_6 instanceof \Data\List\Types\Data_List_Types_Cons && ((($go__go_4_1)(($v_5)->{'value1'}))(($v1_6)->{'value1'}))(($v2_7 && ((($__local_var_1_0)->{'eq'})(($v1_6)->{'value0'}))(($v_5)->{'value0'})))));
  end_branch_2:;
  $__res = $__t2;
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
};
  $__res = ((($go__go_4_1)($xs_2))($ys_3))(true);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $ordList_1_0 = (object)["compare" => function($xs_2) use ($dictOrd_0) {
  $__num = \func_num_args();
  $__res = function($ys_3) use ($dictOrd_0, $xs_2) {
  $__num = \func_num_args();
  $go__go_4_4 = null;
  $go__go_4_4 = (function() use ($dictOrd_0, &$go__go_4_4) {
  $__fn = function($v_5, $v1_6 = null) use ($dictOrd_0, &$go__go_4_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_4_4_4_v_5 = $v_5;
  $__tco_var_go__go_4_4_4_v1_6 = $v1_6;
  tco_loop_go__go_4_4_4:;
  $v_5 = $__tco_var_go__go_4_4_4_v_5;
  $v1_6 = $__tco_var_go__go_4_4_4_v1_6;
  $__t4 = null;;
  if ($v_5 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t5 = null;;
if ($v1_6 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t5 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_5;;
};
$__t5 = new \Data\Ordering\Data_Ordering_LT();
end_branch_5:;
$__t4 = $__t5;
goto end_branch_4;;
};
  if ($v1_6 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t4 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_4;;
};
  if (($v_5 instanceof \Data\List\Types\Data_List_Types_Cons && $v1_6 instanceof \Data\List\Types\Data_List_Types_Cons)) {
$v2_7_6 = ((($dictOrd_0)->{'compare'})(($v_5)->{'value0'}))(($v1_6)->{'value0'});
$__t7 = null;;
if ($v2_7_6 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__tco_8 = ($v_5)->{'value1'};
$__tco_9 = ($v1_6)->{'value1'};
$__tco_var_go__go_4_4_4_v_5 = $__tco_8;
$__tco_var_go__go_4_4_4_v1_6 = $__tco_9;
goto tco_loop_go__go_4_4_4;;
$__t7 = null;
goto end_branch_7;;
};
$__t7 = $v2_7_6;
end_branch_7:;
$__t4 = $__t7;
goto end_branch_4;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t4 = null;
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = (($go__go_4_4)($xs_2))($ys_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq0" => function($_dollar___unused_2) use ($eqList1_1_0) {
  $__num = \func_num_args();
  $__res = $eqList1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_2_6 = (($dictOrd_0)->{'Eq0'})(null);
  $go__go_3_7 = null;
  $go__go_3_7 = function($a_4) use ($__local_var_2_6, &$go__go_3_7) {
  $__num = \func_num_args();
  $__res = function($b_5) use ($__local_var_2_6, $a_4, &$go__go_3_7) {
  $__num = \func_num_args();
  $v_6_8 = \Data\Map\Internal\majData_majMap_majInternal_stepmajAsc($a_4);
  $__t9 = null;;
  if ($v_6_8 instanceof \Data\Map\Internal\Data_Map_Internal_IterNext) {
$v2_7_10 = \Data\Map\Internal\majData_majMap_majInternal_stepmajAsc($b_5);
$__t9 = ($v2_7_10 instanceof \Data\Map\Internal\Data_Map_Internal_IterNext && (((($__local_var_2_6)->{'eq'})(($v_6_8)->{'value0'}))(($v2_7_10)->{'value0'}) && (($go__go_3_7)(($v_6_8)->{'value2'}))(($v2_7_10)->{'value2'})));
goto end_branch_9;;
};
  if ($v_6_8 instanceof \Data\Map\Internal\Data_Map_Internal_IterDone) {
$__t9 = true;
goto end_branch_9;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t9 = null;
  end_branch_9:;
  $__res = $__t9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $eqMapIter2_3_7 = (object)["eq" => $go__go_3_7];
  $eqMap_3_7 = (object)["eq" => function($xs_4) use ($eqMapIter2_3_7) {
  $__num = \func_num_args();
  $__res = function($ys_5) use ($eqMapIter2_3_7, $xs_4) {
  $__num = \func_num_args();
  $__t12 = null;;
  if ($xs_4 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t12 = $ys_5 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf;
goto end_branch_12;;
};
  if ($xs_4 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t12 = ($ys_5 instanceof \Data\Map\Internal\Data_Map_Internal_Node && ((($xs_4)->{'value1'} === ($ys_5)->{'value1'}) && ((($eqMapIter2_3_7)->{'eq'})(new \Data\Map\Internal\Data_Map_Internal_IterNode($xs_4, new \Data\Map\Internal\Data_Map_Internal_IterLeaf())))(new \Data\Map\Internal\Data_Map_Internal_IterNode($ys_5, new \Data\Map\Internal\Data_Map_Internal_IterLeaf()))));
goto end_branch_12;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t12 = null;
  end_branch_12:;
  $__res = $__t12;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $eqSet1_2_6 = (object)["eq" => function($v_4) use ($eqMap_3_7) {
  $__num = \func_num_args();
  $__res = function($v1_5) use ($eqMap_3_7, $v_4) {
  $__num = \func_num_args();
  $__res = ((($eqMap_3_7)->{'eq'})($v_4))($v1_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["compare" => function($s1_3) use ($ordList_1_0) {
  $__num = \func_num_args();
  $__res = function($s2_4) use ($ordList_1_0, $s1_3) {
  $__num = \func_num_args();
  $go__go_5_15 = null;
  $go__go_5_15 = (function() use (&$go__go_5_15) {
  $__fn = function($m_prime__6, $__local_var_7 = null) use (&$go__go_5_15, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t16 = null;;
  if ($m_prime__6 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t16 = $__local_var_7;
goto end_branch_16;;
};
  if ($m_prime__6 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t16 = ($go__go_5_15)(($m_prime__6)->{'value4'}, new \Data\List\Types\Data_List_Types_Cons(($m_prime__6)->{'value2'}, ($go__go_5_15)(($m_prime__6)->{'value5'}, $__local_var_7)));
goto end_branch_16;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t16 = null;
  end_branch_16:;
  $__res = $__t16;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $go__go_5_17 = null;
  $go__go_5_17 = (function() use (&$go__go_5_17) {
  $__fn = function($m_prime__6, $__local_var_7 = null) use (&$go__go_5_17, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t18 = null;;
  if ($m_prime__6 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t18 = $__local_var_7;
goto end_branch_18;;
};
  if ($m_prime__6 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t18 = ($go__go_5_17)(($m_prime__6)->{'value4'}, new \Data\List\Types\Data_List_Types_Cons(($m_prime__6)->{'value2'}, ($go__go_5_17)(($m_prime__6)->{'value5'}, $__local_var_7)));
goto end_branch_18;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t18 = null;
  end_branch_18:;
  $__res = $__t18;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ((($ordList_1_0)->{'compare'})(($go__go_5_15)($s1_3, new \Data\List\Types\Data_List_Types_Nil())))(($go__go_5_17)($s2_4, new \Data\List\Types\Data_List_Types_Nil()));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq0" => function($_dollar___unused_3) use ($eqSet1_2_6) {
  $__num = \func_num_args();
  $__res = $eqSet1_2_6;
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
$GLOBALS['Data_Set_eq1Set'] = (object)["eq1" => function($dictEq_0) {
  $__num = \func_num_args();
  $go__go_1_0 = null;
  $go__go_1_0 = function($a_2) use ($dictEq_0, &$go__go_1_0) {
  $__num = \func_num_args();
  $__res = function($b_3) use ($a_2, $dictEq_0, &$go__go_1_0) {
  $__num = \func_num_args();
  $v_4_1 = \Data\Map\Internal\majData_majMap_majInternal_stepmajAsc($a_2);
  $__t2 = null;;
  if ($v_4_1 instanceof \Data\Map\Internal\Data_Map_Internal_IterNext) {
$v2_5_3 = \Data\Map\Internal\majData_majMap_majInternal_stepmajAsc($b_3);
$__t2 = ($v2_5_3 instanceof \Data\Map\Internal\Data_Map_Internal_IterNext && (((($dictEq_0)->{'eq'})(($v_4_1)->{'value0'}))(($v2_5_3)->{'value0'}) && (($go__go_1_0)(($v_4_1)->{'value2'}))(($v2_5_3)->{'value2'})));
goto end_branch_2;;
};
  if ($v_4_1 instanceof \Data\Map\Internal\Data_Map_Internal_IterDone) {
$__t2 = true;
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $eqMapIter2_1_0 = (object)["eq" => $go__go_1_0];
  $eqMap_1_0 = (object)["eq" => function($xs_2) use ($eqMapIter2_1_0) {
  $__num = \func_num_args();
  $__res = function($ys_3) use ($eqMapIter2_1_0, $xs_2) {
  $__num = \func_num_args();
  $__t5 = null;;
  if ($xs_2 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t5 = $ys_3 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf;
goto end_branch_5;;
};
  if ($xs_2 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t5 = ($ys_3 instanceof \Data\Map\Internal\Data_Map_Internal_Node && ((($xs_2)->{'value1'} === ($ys_3)->{'value1'}) && ((($eqMapIter2_1_0)->{'eq'})(new \Data\Map\Internal\Data_Map_Internal_IterNode($xs_2, new \Data\Map\Internal\Data_Map_Internal_IterLeaf())))(new \Data\Map\Internal\Data_Map_Internal_IterNode($ys_3, new \Data\Map\Internal\Data_Map_Internal_IterLeaf()))));
goto end_branch_5;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t5 = null;
  end_branch_5:;
  $__res = $__t5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = function($v_2) use ($eqMap_1_0) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($eqMap_1_0, $v_2) {
  $__num = \func_num_args();
  $__res = ((($eqMap_1_0)->{'eq'})($v_2))($v1_3);
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

// Data_Set_ord1Set
$GLOBALS['Data_Set_ord1Set'] = (object)["compare1" => function($dictOrd_0) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($dictOrd_0)->{'Eq0'})(null);
  $eqList1_1_0 = (object)["eq" => function($xs_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($ys_3) use ($__local_var_1_0, $xs_2) {
  $__num = \func_num_args();
  $go__go_4_1 = null;
  $go__go_4_1 = function($v_5) use ($__local_var_1_0, &$go__go_4_1) {
  $__num = \func_num_args();
  $__res = function($v1_6) use ($__local_var_1_0, &$go__go_4_1, $v_5) {
  $__num = \func_num_args();
  $__res = function($v2_7) use ($__local_var_1_0, &$go__go_4_1, $v1_6, $v_5) {
  $__num = \func_num_args();
  $__t2 = null;;
  if (( ! $v2_7)) {
$__t2 = false;
goto end_branch_2;;
};
  if ($v_5 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t2 = ($v1_6 instanceof \Data\List\Types\Data_List_Types_Nil && $v2_7);
goto end_branch_2;;
};
  $__t2 = ($v_5 instanceof \Data\List\Types\Data_List_Types_Cons && ($v1_6 instanceof \Data\List\Types\Data_List_Types_Cons && ((($go__go_4_1)(($v_5)->{'value1'}))(($v1_6)->{'value1'}))(($v2_7 && ((($__local_var_1_0)->{'eq'})(($v1_6)->{'value0'}))(($v_5)->{'value0'})))));
  end_branch_2:;
  $__res = $__t2;
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
};
  $__res = ((($go__go_4_1)($xs_2))($ys_3))(true);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $ordList_1_0 = (object)["compare" => function($xs_2) use ($dictOrd_0) {
  $__num = \func_num_args();
  $__res = function($ys_3) use ($dictOrd_0, $xs_2) {
  $__num = \func_num_args();
  $go__go_4_4 = null;
  $go__go_4_4 = (function() use ($dictOrd_0, &$go__go_4_4) {
  $__fn = function($v_5, $v1_6 = null) use ($dictOrd_0, &$go__go_4_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_4_4_4_v_5 = $v_5;
  $__tco_var_go__go_4_4_4_v1_6 = $v1_6;
  tco_loop_go__go_4_4_4:;
  $v_5 = $__tco_var_go__go_4_4_4_v_5;
  $v1_6 = $__tco_var_go__go_4_4_4_v1_6;
  $__t4 = null;;
  if ($v_5 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t5 = null;;
if ($v1_6 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t5 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_5;;
};
$__t5 = new \Data\Ordering\Data_Ordering_LT();
end_branch_5:;
$__t4 = $__t5;
goto end_branch_4;;
};
  if ($v1_6 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t4 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_4;;
};
  if (($v_5 instanceof \Data\List\Types\Data_List_Types_Cons && $v1_6 instanceof \Data\List\Types\Data_List_Types_Cons)) {
$v2_7_6 = ((($dictOrd_0)->{'compare'})(($v_5)->{'value0'}))(($v1_6)->{'value0'});
$__t7 = null;;
if ($v2_7_6 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__tco_8 = ($v_5)->{'value1'};
$__tco_9 = ($v1_6)->{'value1'};
$__tco_var_go__go_4_4_4_v_5 = $__tco_8;
$__tco_var_go__go_4_4_4_v1_6 = $__tco_9;
goto tco_loop_go__go_4_4_4;;
$__t7 = null;
goto end_branch_7;;
};
$__t7 = $v2_7_6;
end_branch_7:;
$__t4 = $__t7;
goto end_branch_4;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t4 = null;
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = (($go__go_4_4)($xs_2))($ys_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq0" => function($_dollar___unused_2) use ($eqList1_1_0) {
  $__num = \func_num_args();
  $__res = $eqList1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = function($s1_2) use ($ordList_1_0) {
  $__num = \func_num_args();
  $__res = function($s2_3) use ($ordList_1_0, $s1_2) {
  $__num = \func_num_args();
  $go__go_4_6 = null;
  $go__go_4_6 = (function() use (&$go__go_4_6) {
  $__fn = function($m_prime__5, $__local_var_6 = null) use (&$go__go_4_6, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t7 = null;;
  if ($m_prime__5 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t7 = $__local_var_6;
goto end_branch_7;;
};
  if ($m_prime__5 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t7 = ($go__go_4_6)(($m_prime__5)->{'value4'}, new \Data\List\Types\Data_List_Types_Cons(($m_prime__5)->{'value2'}, ($go__go_4_6)(($m_prime__5)->{'value5'}, $__local_var_6)));
goto end_branch_7;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t7 = null;
  end_branch_7:;
  $__res = $__t7;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $go__go_4_8 = null;
  $go__go_4_8 = (function() use (&$go__go_4_8) {
  $__fn = function($m_prime__5, $__local_var_6 = null) use (&$go__go_4_8, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t9 = null;;
  if ($m_prime__5 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t9 = $__local_var_6;
goto end_branch_9;;
};
  if ($m_prime__5 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t9 = ($go__go_4_8)(($m_prime__5)->{'value4'}, new \Data\List\Types\Data_List_Types_Cons(($m_prime__5)->{'value2'}, ($go__go_4_8)(($m_prime__5)->{'value5'}, $__local_var_6)));
goto end_branch_9;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t9 = null;
  end_branch_9:;
  $__res = $__t9;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ((($ordList_1_0)->{'compare'})(($go__go_4_6)($s1_2, new \Data\List\Types\Data_List_Types_Nil())))(($go__go_4_8)($s2_3, new \Data\List\Types\Data_List_Types_Nil()));
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
}, "Eq10" => function($_dollar___unused_0) {
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
  $__res = ((($dictFoldable_0)->{'foldl'})(function($m_2) use ($dictOrd_1) {
  $__num = \func_num_args();
  $__res = function($a_3) use ($dictOrd_1, $m_2) {
  $__num = \func_num_args();
  $go__go_4_0 = null;
  $go__go_4_0 = function($v1_5) use ($a_3, $dictOrd_1, &$go__go_4_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v1_5 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t1 = new \Data\Map\Internal\Data_Map_Internal_Node(1, 1, $a_3, $GLOBALS['Data_Unit_unit'], new \Data\Map\Internal\Data_Map_Internal_Leaf(), new \Data\Map\Internal\Data_Map_Internal_Leaf());
goto end_branch_1;;
};
  if ($v1_5 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$v2_6_2 = ((($dictOrd_1)->{'compare'})($a_3))(($v1_5)->{'value2'});
$__t3 = null;;
if ($v2_6_2 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t3 = ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($v1_5)->{'value2'}, ($v1_5)->{'value3'}, ($go__go_4_0)(($v1_5)->{'value4'}), ($v1_5)->{'value5'});
goto end_branch_3;;
};
if ($v2_6_2 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t3 = ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($v1_5)->{'value2'}, ($v1_5)->{'value3'}, ($v1_5)->{'value4'}, ($go__go_4_0)(($v1_5)->{'value5'}));
goto end_branch_3;;
};
if ($v2_6_2 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t3 = new \Data\Map\Internal\Data_Map_Internal_Node(($v1_5)->{'value0'}, ($v1_5)->{'value1'}, $a_3, $GLOBALS['Data_Unit_unit'], ($v1_5)->{'value4'}, ($v1_5)->{'value5'});
goto end_branch_3;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
end_branch_3:;
$__t1 = $__t3;
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = ($go__go_4_0)($m_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(new \Data\Map\Internal\Data_Map_Internal_Leaf());
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
  $go__go_2_0 = null;
  $go__go_2_0 = (function() use ($dictOrd_0, $f_1, &$go__go_2_0) {
  $__fn = function($b_3, $v_4 = null) use ($dictOrd_0, $f_1, &$go__go_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_2_0_0_b_3 = $b_3;
  $__tco_var_go__go_2_0_0_v_4 = $v_4;
  tco_loop_go__go_2_0_0:;
  $b_3 = $__tco_var_go__go_2_0_0_b_3;
  $v_4 = $__tco_var_go__go_2_0_0_v_4;
  $__t0 = null;;
  if ($v_4 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = $b_3;
goto end_branch_0;;
};
  if ($v_4 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__local_var_5_1 = ($f_1)(($v_4)->{'value0'});
$go__go_6_2 = null;
$go__go_6_2 = function($v1_7) use ($__local_var_5_1, $dictOrd_0, &$go__go_6_2) {
  $__num = \func_num_args();
  $__t3 = null;;
  if ($v1_7 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t3 = new \Data\Map\Internal\Data_Map_Internal_Node(1, 1, $__local_var_5_1, $GLOBALS['Data_Unit_unit'], new \Data\Map\Internal\Data_Map_Internal_Leaf(), new \Data\Map\Internal\Data_Map_Internal_Leaf());
goto end_branch_3;;
};
  if ($v1_7 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$v2_8_4 = ((($dictOrd_0)->{'compare'})($__local_var_5_1))(($v1_7)->{'value2'});
$__t5 = null;;
if ($v2_8_4 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t5 = ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($v1_7)->{'value2'}, ($v1_7)->{'value3'}, ($go__go_6_2)(($v1_7)->{'value4'}), ($v1_7)->{'value5'});
goto end_branch_5;;
};
if ($v2_8_4 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t5 = ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($v1_7)->{'value2'}, ($v1_7)->{'value3'}, ($v1_7)->{'value4'}, ($go__go_6_2)(($v1_7)->{'value5'}));
goto end_branch_5;;
};
if ($v2_8_4 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t5 = new \Data\Map\Internal\Data_Map_Internal_Node(($v1_7)->{'value0'}, ($v1_7)->{'value1'}, $__local_var_5_1, $GLOBALS['Data_Unit_unit'], ($v1_7)->{'value4'}, ($v1_7)->{'value5'});
goto end_branch_5;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t5 = null;
end_branch_5:;
$__t3 = $__t5;
goto end_branch_3;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t3 = null;
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
$__tco_6 = ($go__go_6_2)($b_3);
$__tco_7 = ($v_4)->{'value1'};
$__tco_var_go__go_2_0_0_b_3 = $__tco_6;
$__tco_var_go__go_2_0_0_v_4 = $__tco_7;
goto tco_loop_go__go_2_0_0;;
$__t0 = null;
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($go__go_2_0)(new \Data\Map\Internal\Data_Map_Internal_Leaf())))($GLOBALS['Data_Set_toList']);
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
  $go__go_2_0 = null;
  $go__go_2_0 = (function() use ($dictOrd_0, $f_1, &$go__go_2_0) {
  $__fn = function($b_3, $v_4 = null) use ($dictOrd_0, $f_1, &$go__go_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_2_0_0_b_3 = $b_3;
  $__tco_var_go__go_2_0_0_v_4 = $v_4;
  tco_loop_go__go_2_0_0:;
  $b_3 = $__tco_var_go__go_2_0_0_b_3;
  $v_4 = $__tco_var_go__go_2_0_0_v_4;
  $__t0 = null;;
  if ($v_4 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = $b_3;
goto end_branch_0;;
};
  if ($v_4 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__local_var_5_1 = ($f_1)(($v_4)->{'value0'});
$__t2 = null;;
if ($__local_var_5_1 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t2 = $b_3;
goto end_branch_2;;
};
if ($__local_var_5_1 instanceof \Data\Maybe\Data_Maybe_Just) {
$go__go_6_3 = null;
$go__go_6_3 = function($v1_7) use ($__local_var_5_1, $dictOrd_0, &$go__go_6_3) {
  $__num = \func_num_args();
  $__t4 = null;;
  if ($v1_7 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t4 = new \Data\Map\Internal\Data_Map_Internal_Node(1, 1, ($__local_var_5_1)->{'value0'}, $GLOBALS['Data_Unit_unit'], new \Data\Map\Internal\Data_Map_Internal_Leaf(), new \Data\Map\Internal\Data_Map_Internal_Leaf());
goto end_branch_4;;
};
  if ($v1_7 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$v2_8_5 = ((($dictOrd_0)->{'compare'})(($__local_var_5_1)->{'value0'}))(($v1_7)->{'value2'});
$__t6 = null;;
if ($v2_8_5 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t6 = ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($v1_7)->{'value2'}, ($v1_7)->{'value3'}, ($go__go_6_3)(($v1_7)->{'value4'}), ($v1_7)->{'value5'});
goto end_branch_6;;
};
if ($v2_8_5 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t6 = ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($v1_7)->{'value2'}, ($v1_7)->{'value3'}, ($v1_7)->{'value4'}, ($go__go_6_3)(($v1_7)->{'value5'}));
goto end_branch_6;;
};
if ($v2_8_5 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t6 = new \Data\Map\Internal\Data_Map_Internal_Node(($v1_7)->{'value0'}, ($v1_7)->{'value1'}, ($__local_var_5_1)->{'value0'}, $GLOBALS['Data_Unit_unit'], ($v1_7)->{'value4'}, ($v1_7)->{'value5'});
goto end_branch_6;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t6 = null;
end_branch_6:;
$__t4 = $__t6;
goto end_branch_4;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t4 = null;
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
$__t2 = ($go__go_6_3)($b_3);
goto end_branch_2;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
end_branch_2:;
$__tco_7 = $__t2;
$__tco_8 = ($v_4)->{'value1'};
$__tco_var_go__go_2_0_0_b_3 = $__tco_7;
$__tco_var_go__go_2_0_0_v_4 = $__tco_8;
goto tco_loop_go__go_2_0_0;;
$__t0 = null;
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $go__go_2_1 = null;
  $go__go_2_1 = (function() use (&$go__go_2_1) {
  $__fn = function($v_3, $v1_4 = null) use (&$go__go_2_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_2_1_1_v_3 = $v_3;
  $__tco_var_go__go_2_1_1_v1_4 = $v1_4;
  tco_loop_go__go_2_1_1:;
  $v_3 = $__tco_var_go__go_2_1_1_v_3;
  $v1_4 = $__tco_var_go__go_2_1_1_v1_4;
  $__t1 = null;;
  if ($v1_4 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = $v_3;
goto end_branch_1;;
};
  if ($v1_4 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_2 = new \Data\List\Types\Data_List_Types_Cons(($v1_4)->{'value0'}, $v_3);
$__tco_3 = ($v1_4)->{'value1'};
$__tco_var_go__go_2_1_1_v_3 = $__tco_2;
$__tco_var_go__go_2_1_1_v1_4 = $__tco_3;
goto tco_loop_go__go_2_1_1;;
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
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Control_Semigroupoid_composeImpl'])(($go__go_2_0)(new \Data\Map\Internal\Data_Map_Internal_Leaf())))(($go__go_2_1)(new \Data\List\Types\Data_List_Types_Nil()))))($GLOBALS['Data_Set_toList']);
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
  $compare_1_0 = ($dictOrd_0)->{'compare'};
  $semigroupSet1_1_0 = (object)["append" => function($m1_2) use ($compare_1_0) {
  $__num = \func_num_args();
  $__res = function($m2_3) use ($compare_1_0, $m1_2) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Map_Internal_unsafeUnionWith'])($compare_1_0, $GLOBALS['Data_Function_const'], $m1_2, $m2_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["mempty" => new \Data\Map\Internal\Data_Map_Internal_Leaf(), "Semigroup0" => function($_dollar___unused_2) use ($semigroupSet1_1_0) {
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
  $compare_2_0 = ($dictOrd_1)->{'compare'};
  $__res = ((($dictFoldable_0)->{'foldl'})(function($m1_3) use ($compare_2_0) {
  $__num = \func_num_args();
  $__res = function($m2_4) use ($compare_2_0, $m1_3) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Map_Internal_unsafeUnionWith'])($compare_2_0, $GLOBALS['Data_Function_const'], $m1_3, $m2_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(new \Data\Map\Internal\Data_Map_Internal_Leaf());
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
  $compare_1_0 = ($dictOrd_0)->{'compare'};
  $__res = function($m1_2) use ($compare_1_0) {
  $__num = \func_num_args();
  $__res = function($m2_3) use ($compare_1_0, $m1_2) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Map_Internal_unsafeDifference'])($compare_1_0, $m1_2, $m2_3);
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Set_difference'] = __NAMESPACE__ . '\\majData_majSet_difference';

// Data_Set_subset
function majData_majSet_subset($dictOrd_0, $s1_1 = null, $s2_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_subset';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($GLOBALS['Data_Map_Internal_unsafeDifference'])(($dictOrd_0)->{'compare'}, $s1_1, $s2_2) instanceof \Data\Map\Internal\Data_Map_Internal_Leaf;
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Set_subset'] = __NAMESPACE__ . '\\majData_majSet_subset';

// Data_Set_properSubset
function majData_majSet_propermajSubset($dictOrd_0, $s1_1 = null, $s2_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_propermajSubset';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if ($s1_1 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t0 = 0;
goto end_branch_0;;
};
  if ($s1_1 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t0 = ($s1_1)->{'value1'};
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__t1 = null;;
  if ($s2_2 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t1 = 0;
goto end_branch_1;;
};
  if ($s2_2 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t1 = ($s2_2)->{'value1'};
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = (( ! ($__t0 === $__t1)) && ($GLOBALS['Data_Map_Internal_unsafeDifference'])(($dictOrd_0)->{'compare'}, $s1_1, $s2_2) instanceof \Data\Map\Internal\Data_Map_Internal_Leaf);
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Set_properSubset'] = __NAMESPACE__ . '\\majData_majSet_propermajSubset';

// Data_Set_delete
function majData_majSet_delete($dictOrd_0, $k_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_delete';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__go_2_0 = null;
  $go__go_2_0 = function($v_3) use ($dictOrd_0, &$go__go_2_0, $k_1) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v_3 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t1 = new \Data\Map\Internal\Data_Map_Internal_Leaf();
goto end_branch_1;;
};
  if ($v_3 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$v1_4_2 = ((($dictOrd_0)->{'compare'})($k_1))(($v_3)->{'value2'});
$__t3 = null;;
if ($v1_4_2 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t3 = ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($v_3)->{'value2'}, ($v_3)->{'value3'}, ($go__go_2_0)(($v_3)->{'value4'}), ($v_3)->{'value5'});
goto end_branch_3;;
};
if ($v1_4_2 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t3 = ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($v_3)->{'value2'}, ($v_3)->{'value3'}, ($v_3)->{'value4'}, ($go__go_2_0)(($v_3)->{'value5'}));
goto end_branch_3;;
};
if ($v1_4_2 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t3 = ($GLOBALS['Data_Map_Internal_unsafeJoinNodes'])(($v_3)->{'value4'}, ($v_3)->{'value5'});
goto end_branch_3;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
end_branch_3:;
$__t1 = $__t3;
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = $go__go_2_0;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Set_delete'] = __NAMESPACE__ . '\\majData_majSet_delete';

// Data_Set_checkValid
function majData_majSet_checkmajValid($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_checkmajValid';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $go__go_1_0 = null;
  $go__go_1_0 = function($v_2) use ($dictOrd_0, &$go__go_1_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v_2 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t1 = true;
goto end_branch_1;;
};
  if ($v_2 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t2 = null;;
if (($v_2)->{'value4'} instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t3 = null;;
if (($v_2)->{'value5'} instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t3 = true;
goto end_branch_3;;
};
if (($v_2)->{'value5'} instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t3 = ((($v_2)->{'value0'} === 2) && (((($v_2)->{'value5'})->{'value0'} === 1) && ((($v_2)->{'value1'} > (($v_2)->{'value5'})->{'value1'}) && (((($dictOrd_0)->{'compare'})((($v_2)->{'value5'})->{'value2'}))(($v_2)->{'value2'}) instanceof \Data\Ordering\Data_Ordering_GT && ($go__go_1_0)(($v_2)->{'value5'})))));
goto end_branch_3;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
end_branch_3:;
$__t2 = $__t3;
goto end_branch_2;;
};
if (($v_2)->{'value4'} instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t4 = null;;
if (($v_2)->{'value5'} instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t4 = ((($v_2)->{'value0'} === 2) && (((($v_2)->{'value4'})->{'value0'} === 1) && ((($v_2)->{'value1'} > (($v_2)->{'value4'})->{'value1'}) && (((($dictOrd_0)->{'compare'})((($v_2)->{'value4'})->{'value2'}))(($v_2)->{'value2'}) instanceof \Data\Ordering\Data_Ordering_LT && ($go__go_1_0)(($v_2)->{'value4'})))));
goto end_branch_4;;
};
if (($v_2)->{'value5'} instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__local_var_3_5 = ((($v_2)->{'value5'})->{'value0'} - (($v_2)->{'value4'})->{'value0'});
$__t6 = null;;
if (($__local_var_3_5 >= 0)) {
$__t6 = $__local_var_3_5;
goto end_branch_6;;
};
$__t6 = ( - $__local_var_3_5);
end_branch_6:;
$__t4 = ((($v_2)->{'value0'} > (($v_2)->{'value5'})->{'value0'}) && (((($dictOrd_0)->{'compare'})((($v_2)->{'value5'})->{'value2'}))(($v_2)->{'value2'}) instanceof \Data\Ordering\Data_Ordering_GT && ((($v_2)->{'value0'} > (($v_2)->{'value4'})->{'value0'}) && (((($dictOrd_0)->{'compare'})((($v_2)->{'value4'})->{'value2'}))(($v_2)->{'value2'}) instanceof \Data\Ordering\Data_Ordering_LT && (($__t6 < 2) && (((((($v_2)->{'value5'})->{'value1'} + (($v_2)->{'value4'})->{'value1'}) + 1) === ($v_2)->{'value1'}) && (($go__go_1_0)(($v_2)->{'value4'}) && ($go__go_1_0)(($v_2)->{'value5'}))))))));
goto end_branch_4;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t4 = null;
end_branch_4:;
$__t2 = $__t4;
goto end_branch_2;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
end_branch_2:;
$__t1 = $__t2;
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = $go__go_1_0;
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
  $go__go_1_0 = null;
  $go__go_1_0 = (function() use ($dictOrd_0, &$go__go_1_0) {
  $__fn = function($b_2, $v_3 = null) use ($dictOrd_0, &$go__go_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_1_0_0_b_2 = $b_2;
  $__tco_var_go__go_1_0_0_v_3 = $v_3;
  tco_loop_go__go_1_0_0:;
  $b_2 = $__tco_var_go__go_1_0_0_b_2;
  $v_3 = $__tco_var_go__go_1_0_0_v_3;
  $__t0 = null;;
  if ($v_3 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = $b_2;
goto end_branch_0;;
};
  if ($v_3 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__t1 = null;;
if (($v_3)->{'value0'} instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t1 = $b_2;
goto end_branch_1;;
};
if (($v_3)->{'value0'} instanceof \Data\Maybe\Data_Maybe_Just) {
$go__go_4_2 = null;
$go__go_4_2 = function($v1_5) use ($dictOrd_0, &$go__go_4_2, $v_3) {
  $__num = \func_num_args();
  $__t3 = null;;
  if ($v1_5 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t3 = new \Data\Map\Internal\Data_Map_Internal_Node(1, 1, (($v_3)->{'value0'})->{'value0'}, $GLOBALS['Data_Unit_unit'], new \Data\Map\Internal\Data_Map_Internal_Leaf(), new \Data\Map\Internal\Data_Map_Internal_Leaf());
goto end_branch_3;;
};
  if ($v1_5 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$v2_6_4 = ((($dictOrd_0)->{'compare'})((($v_3)->{'value0'})->{'value0'}))(($v1_5)->{'value2'});
$__t5 = null;;
if ($v2_6_4 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t5 = ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($v1_5)->{'value2'}, ($v1_5)->{'value3'}, ($go__go_4_2)(($v1_5)->{'value4'}), ($v1_5)->{'value5'});
goto end_branch_5;;
};
if ($v2_6_4 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t5 = ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($v1_5)->{'value2'}, ($v1_5)->{'value3'}, ($v1_5)->{'value4'}, ($go__go_4_2)(($v1_5)->{'value5'}));
goto end_branch_5;;
};
if ($v2_6_4 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t5 = new \Data\Map\Internal\Data_Map_Internal_Node(($v1_5)->{'value0'}, ($v1_5)->{'value1'}, (($v_3)->{'value0'})->{'value0'}, $GLOBALS['Data_Unit_unit'], ($v1_5)->{'value4'}, ($v1_5)->{'value5'});
goto end_branch_5;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t5 = null;
end_branch_5:;
$__t3 = $__t5;
goto end_branch_3;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t3 = null;
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
$__t1 = ($go__go_4_2)($b_2);
goto end_branch_1;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
end_branch_1:;
$__tco_6 = $__t1;
$__tco_7 = ($v_3)->{'value1'};
$__tco_var_go__go_1_0_0_b_2 = $__tco_6;
$__tco_var_go__go_1_0_0_v_3 = $__tco_7;
goto tco_loop_go__go_1_0_0;;
$__t0 = null;
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $go__go_1_1 = null;
  $go__go_1_1 = (function() use (&$go__go_1_1) {
  $__fn = function($v_2, $v1_3 = null) use (&$go__go_1_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_1_1_1_v_2 = $v_2;
  $__tco_var_go__go_1_1_1_v1_3 = $v1_3;
  tco_loop_go__go_1_1_1:;
  $v_2 = $__tco_var_go__go_1_1_1_v_2;
  $v1_3 = $__tco_var_go__go_1_1_1_v1_3;
  $__t1 = null;;
  if ($v1_3 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = $v_2;
goto end_branch_1;;
};
  if ($v1_3 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_2 = new \Data\List\Types\Data_List_Types_Cons(($v1_3)->{'value0'}, $v_2);
$__tco_3 = ($v1_3)->{'value1'};
$__tco_var_go__go_1_1_1_v_2 = $__tco_2;
$__tco_var_go__go_1_1_1_v1_3 = $__tco_3;
goto tco_loop_go__go_1_1_1;;
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
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Control_Semigroupoid_composeImpl'])(($go__go_1_0)(new \Data\Map\Internal\Data_Map_Internal_Leaf())))(($go__go_1_1)(new \Data\List\Types\Data_List_Types_Nil()))))($GLOBALS['Data_Set_toList']);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Set_catMaybes'] = __NAMESPACE__ . '\\majData_majSet_catmajMaybes';

