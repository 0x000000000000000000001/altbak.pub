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




// Data_Set_NonEmpty_unionSet
function majData_majSet_majNonmajEmpty_unionmajSet($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_majNonmajEmpty_unionmajSet';
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
$GLOBALS['Data_Set_NonEmpty_unionSet'] = __NAMESPACE__ . '\\majData_majSet_majNonmajEmpty_unionmajSet';

// Data_Set_NonEmpty_toUnfoldable1
function majData_majSet_majNonmajEmpty_tomajUnfoldable1($dictUnfoldable1_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_majNonmajEmpty_tomajUnfoldable1';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $stepNext_1_0 = ($GLOBALS['Data_Map_Internal_stepWith'])($GLOBALS['Data_Map_Internal_iterMapL'], (function() {
  $__fn = function($k_1 = null, $v_2 = null, $next_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = new \Data\Maybe\Data_Maybe_Just(new \Data\Tuple\Data_Tuple_Tuple($k_1, $next_3));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), function($v_1 = null) {
  $__num = \func_num_args();
  $__res = new \Data\Maybe\Data_Maybe_Nothing();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = ($GLOBALS['Control_Semigroupoid_composeImpl'])((($dictUnfoldable1_0)['unfoldr1'])(function($v_2 = null) use ($stepNext_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Tuple_functorTuple'])['map'])($stepNext_1_0, $v_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), ($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Map_Internal_stepWith'])($GLOBALS['Data_Map_Internal_iterMapL'], (function() {
  $__fn = function($k_2 = null, $v_3 = null, $next_4 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple($k_2, $next_4);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), function($v_2 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Partial__crashWith'])("toUnfoldable1: impossible");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), ($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Map_Internal_toMapIter'], ($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Set_toMap'], $GLOBALS['Unsafe_Coerce_unsafeCoerce']))));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Set_NonEmpty_toUnfoldable1'] = __NAMESPACE__ . '\\majData_majSet_majNonmajEmpty_tomajUnfoldable1';

// Data_Set_NonEmpty_toUnfoldable11
$GLOBALS['Data_Set_NonEmpty_toUnfoldable11'] = ($GLOBALS['Data_Set_NonEmpty_toUnfoldable1'])($GLOBALS['Data_Unfoldable1_unfoldable1Array']);

// Data_Set_NonEmpty_toUnfoldable12_closure
$GLOBALS['Data_Set_NonEmpty_toUnfoldable12_closure'] = ($GLOBALS['Data_Set_NonEmpty_toUnfoldable1'])($GLOBALS['Data_List_Types_unfoldable1NonEmptyList']);

// Data_Set_NonEmpty_toUnfoldable12
function majData_majSet_majNonmajEmpty_tomajUnfoldable12($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_majNonmajEmpty_tomajUnfoldable12';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Set_NonEmpty_toUnfoldable12_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Set_NonEmpty_toUnfoldable12'] = __NAMESPACE__ . '\\majData_majSet_majNonmajEmpty_tomajUnfoldable12';

// Data_Set_NonEmpty_toUnfoldable
function majData_majSet_majNonmajEmpty_tomajUnfoldable($dictUnfoldable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_majNonmajEmpty_tomajUnfoldable';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_List_toUnfoldable'])($dictUnfoldable_0), $GLOBALS['Data_Set_toList']);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Set_NonEmpty_toUnfoldable'] = __NAMESPACE__ . '\\majData_majSet_majNonmajEmpty_tomajUnfoldable';

// Data_Set_NonEmpty_toSet
function majData_majSet_majNonmajEmpty_tomajSet($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_majNonmajEmpty_tomajSet';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $v_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Set_NonEmpty_toSet'] = __NAMESPACE__ . '\\majData_majSet_majNonmajEmpty_tomajSet';

// Data_Set_NonEmpty_subset
function majData_majSet_majNonmajEmpty_subset($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_majNonmajEmpty_subset';
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
$GLOBALS['Data_Set_NonEmpty_subset'] = __NAMESPACE__ . '\\majData_majSet_majNonmajEmpty_subset';

// Data_Set_NonEmpty_size_closure
$GLOBALS['Data_Set_NonEmpty_size_closure'] = $GLOBALS['Data_Map_Internal_size'];

// Data_Set_NonEmpty_size
function majData_majSet_majNonmajEmpty_size($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_majNonmajEmpty_size';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Set_NonEmpty_size_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Set_NonEmpty_size'] = __NAMESPACE__ . '\\majData_majSet_majNonmajEmpty_size';

// Data_Set_NonEmpty_singleton_closure
$GLOBALS['Data_Set_NonEmpty_singleton_closure'] = $GLOBALS['Data_Set_singleton'];

// Data_Set_NonEmpty_singleton
function majData_majSet_majNonmajEmpty_singleton($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_majNonmajEmpty_singleton';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Set_NonEmpty_singleton_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Set_NonEmpty_singleton'] = __NAMESPACE__ . '\\majData_majSet_majNonmajEmpty_singleton';

// Data_Set_NonEmpty_showNonEmptySet
function majData_majSet_majNonmajEmpty_showmajNonmajEmptymajSet($dictShow_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_majNonmajEmpty_showmajNonmajEmptymajSet';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ["show" => function($s_1 = null) use ($dictShow_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Semigroup_semigroupString'])['append'])("(fromFoldable1 ", (($GLOBALS['Data_Semigroup_semigroupString'])['append'])(((($GLOBALS['Data_Array_NonEmpty_Internal_showNonEmptyArray'])($dictShow_0))['show'])(($GLOBALS['Data_Set_NonEmpty_toUnfoldable11'])($s_1)), ")"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Set_NonEmpty_showNonEmptySet'] = __NAMESPACE__ . '\\majData_majSet_majNonmajEmpty_showmajNonmajEmptymajSet';

// Data_Set_NonEmpty_semigroupNonEmptySet
function majData_majSet_majNonmajEmpty_semigroupmajNonmajEmptymajSet($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_majNonmajEmpty_semigroupmajNonmajEmptymajSet';
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
$GLOBALS['Data_Set_NonEmpty_semigroupNonEmptySet'] = __NAMESPACE__ . '\\majData_majSet_majNonmajEmpty_semigroupmajNonmajEmptymajSet';

// Data_Set_NonEmpty_properSubset
function majData_majSet_majNonmajEmpty_propermajSubset($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_majNonmajEmpty_propermajSubset';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Set_properSubset'])($dictOrd_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Set_NonEmpty_properSubset'] = __NAMESPACE__ . '\\majData_majSet_majNonmajEmpty_propermajSubset';

// Data_Set_NonEmpty_ordNonEmptySet
function majData_majSet_majNonmajEmpty_ordmajNonmajEmptymajSet($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_majNonmajEmpty_ordmajNonmajEmptymajSet';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Set_ordSet'])($dictOrd_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Set_NonEmpty_ordNonEmptySet'] = __NAMESPACE__ . '\\majData_majSet_majNonmajEmpty_ordmajNonmajEmptymajSet';

// Data_Set_NonEmpty_ord1NonEmptySet
$GLOBALS['Data_Set_NonEmpty_ord1NonEmptySet'] = $GLOBALS['Data_Set_ord1Set'];

// Data_Set_NonEmpty_min
function majData_majSet_majNonmajEmpty_min($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_majNonmajEmpty_min';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($GLOBALS['Data_Maybe_functorMaybe'])['map'])(function($v1_1 = null) {
  $__num = \func_num_args();
  $__res = ($v1_1)['key'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($GLOBALS['Data_Map_Internal_findMin'])($v_0));
  $__t1 = null;;
  if ($__local_var_1_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = ($__local_var_1_0)->{'value0'};
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Set_NonEmpty_min'] = __NAMESPACE__ . '\\majData_majSet_majNonmajEmpty_min';

// Data_Set_NonEmpty_member
function majData_majSet_majNonmajEmpty_member($dictOrd_0, $k_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_majNonmajEmpty_member';
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
$GLOBALS['Data_Set_NonEmpty_member'] = __NAMESPACE__ . '\\majData_majSet_majNonmajEmpty_member';

// Data_Set_NonEmpty_max
function majData_majSet_majNonmajEmpty_max($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_majNonmajEmpty_max';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($GLOBALS['Data_Maybe_functorMaybe'])['map'])(function($v1_1 = null) {
  $__num = \func_num_args();
  $__res = ($v1_1)['key'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($GLOBALS['Data_Map_Internal_findMax'])($v_0));
  $__t1 = null;;
  if ($__local_var_1_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = ($__local_var_1_0)->{'value0'};
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Set_NonEmpty_max'] = __NAMESPACE__ . '\\majData_majSet_majNonmajEmpty_max';

// Data_Set_NonEmpty_mapMaybe
function majData_majSet_majNonmajEmpty_mapmajMaybe($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_majNonmajEmpty_mapmajMaybe';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Set_mapMaybe'])($dictOrd_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Set_NonEmpty_mapMaybe'] = __NAMESPACE__ . '\\majData_majSet_majNonmajEmpty_mapmajMaybe';

// Data_Set_NonEmpty_map
function majData_majSet_majNonmajEmpty_map($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_majNonmajEmpty_map';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Set_map'])($dictOrd_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Set_NonEmpty_map'] = __NAMESPACE__ . '\\majData_majSet_majNonmajEmpty_map';

// Data_Set_NonEmpty_insert
function majData_majSet_majNonmajEmpty_insert($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_majNonmajEmpty_insert';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Set_insert'])($dictOrd_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Set_NonEmpty_insert'] = __NAMESPACE__ . '\\majData_majSet_majNonmajEmpty_insert';

// Data_Set_NonEmpty_fromSet
function majData_majSet_majNonmajEmpty_frommajSet($s_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_majNonmajEmpty_frommajSet';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__t0 = null;;
  if ($s_0 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t0 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_0;;
};
  $__t0 = new \Data\Maybe\Data_Maybe_Just($s_0);
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Set_NonEmpty_fromSet'] = __NAMESPACE__ . '\\majData_majSet_majNonmajEmpty_frommajSet';

// Data_Set_NonEmpty_intersection
function majData_majSet_majNonmajEmpty_intersection($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_majNonmajEmpty_intersection';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $compare_1_0 = ($dictOrd_0)['compare'];
  $__res = (function() use ($compare_1_0) {
  $__fn = function($v_2 = null, $v1_3 = null) use ($compare_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_4_1 = ($GLOBALS['Data_Map_Internal_unsafeIntersectionWith'])($compare_1_0, $GLOBALS['Data_Function_const'], $v_2, $v1_3);
  $__t2 = null;;
  if ($__local_var_4_1 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t2 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_2;;
};
  $__t2 = new \Data\Maybe\Data_Maybe_Just($__local_var_4_1);
  end_branch_2:;
  $__res = $__t2;
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
$GLOBALS['Data_Set_NonEmpty_intersection'] = __NAMESPACE__ . '\\majData_majSet_majNonmajEmpty_intersection';

// Data_Set_NonEmpty_fromFoldable1
function majData_majSet_majNonmajEmpty_frommajFoldable1($dictFoldable1_0, $dictOrd_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_majNonmajEmpty_frommajFoldable1';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $compare_2_0 = ($dictOrd_1)['compare'];
  $__res = (($dictFoldable1_0)['foldMap1'])(["append" => (function() use ($compare_2_0) {
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
})()], $GLOBALS['Data_Set_singleton']);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Set_NonEmpty_fromFoldable1'] = __NAMESPACE__ . '\\majData_majSet_majNonmajEmpty_frommajFoldable1';

// Data_Set_NonEmpty_fromFoldable
function majData_majSet_majNonmajEmpty_frommajFoldable($dictFoldable_0, $dictOrd_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_majNonmajEmpty_frommajFoldable';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Set_NonEmpty_fromSet'], (($dictFoldable_0)['foldl'])((function() use ($dictOrd_1) {
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
})(), new \Data\Map\Internal\Data_Map_Internal_Leaf()));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Set_NonEmpty_fromFoldable'] = __NAMESPACE__ . '\\majData_majSet_majNonmajEmpty_frommajFoldable';

// Data_Set_NonEmpty_foldableNonEmptySet
$GLOBALS['Data_Set_NonEmpty_foldableNonEmptySet'] = $GLOBALS['Data_Set_foldableSet'];

// Data_Set_NonEmpty_foldable1NonEmptySet
$GLOBALS['Data_Set_NonEmpty_foldable1NonEmptySet'] = ["foldMap1" => function($dictSemigroup_0 = null) {
  $__num = \func_num_args();
  $foldMap11_1_0 = (($GLOBALS['Data_List_Types_foldable1NonEmptyList'])['foldMap1'])($dictSemigroup_0);
  $__res = function($f_2 = null) use ($foldMap11_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Semigroupoid_composeImpl'])(($foldMap11_1_0)($f_2), $GLOBALS['Data_Set_NonEmpty_toUnfoldable12']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldr1" => function($f_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_List_Types_foldable1NonEmptyList'])['foldr1'])($f_0), $GLOBALS['Data_Set_NonEmpty_toUnfoldable12']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldl1" => function($f_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_List_Types_foldable1NonEmptyList'])['foldl1'])($f_0), $GLOBALS['Data_Set_NonEmpty_toUnfoldable12']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable0" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Set_foldableSet'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Set_NonEmpty_filter
function majData_majSet_majNonmajEmpty_filter($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_majNonmajEmpty_filter';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Set_filter'])($dictOrd_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Set_NonEmpty_filter'] = __NAMESPACE__ . '\\majData_majSet_majNonmajEmpty_filter';

// Data_Set_NonEmpty_eqNonEmptySet
function majData_majSet_majNonmajEmpty_eqmajNonmajEmptymajSet($dictEq_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_majNonmajEmpty_eqmajNonmajEmptymajSet';
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
$GLOBALS['Data_Set_NonEmpty_eqNonEmptySet'] = __NAMESPACE__ . '\\majData_majSet_majNonmajEmpty_eqmajNonmajEmptymajSet';

// Data_Set_NonEmpty_eq1NonEmptySet
$GLOBALS['Data_Set_NonEmpty_eq1NonEmptySet'] = $GLOBALS['Data_Set_eq1Set'];

// Data_Set_NonEmpty_difference
function majData_majSet_majNonmajEmpty_difference($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_majNonmajEmpty_difference';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $compare_1_0 = ($dictOrd_0)['compare'];
  $__res = (function() use ($compare_1_0) {
  $__fn = function($v_2 = null, $v1_3 = null) use ($compare_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_4_1 = ($GLOBALS['Data_Map_Internal_unsafeDifference'])($compare_1_0, $v_2, $v1_3);
  $__t2 = null;;
  if ($__local_var_4_1 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t2 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_2;;
};
  $__t2 = new \Data\Maybe\Data_Maybe_Just($__local_var_4_1);
  end_branch_2:;
  $__res = $__t2;
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
$GLOBALS['Data_Set_NonEmpty_difference'] = __NAMESPACE__ . '\\majData_majSet_majNonmajEmpty_difference';

// Data_Set_NonEmpty_delete
function majData_majSet_majNonmajEmpty_delete($dictOrd_0, $a_1 = null, $v_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_majNonmajEmpty_delete';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__local_var_3_0 = ($GLOBALS['Data_Map_Internal_delete'])($dictOrd_0, $a_1, $v_2);
  $__t1 = null;;
  if ($__local_var_3_0 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t1 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_1;;
};
  $__t1 = new \Data\Maybe\Data_Maybe_Just($__local_var_3_0);
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Set_NonEmpty_delete'] = __NAMESPACE__ . '\\majData_majSet_majNonmajEmpty_delete';

// Data_Set_NonEmpty_cons
function majData_majSet_majNonmajEmpty_cons($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSet_majNonmajEmpty_cons';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Set_insert'])($dictOrd_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Set_NonEmpty_cons'] = __NAMESPACE__ . '\\majData_majSet_majNonmajEmpty_cons';

