<?php

namespace Data\Map;

// ALL IMPORTS: Control.Alt, Control.Apply, Control.Bind, Control.Plus, Control.Semigroupoid, Data.Eq, Data.Foldable, Data.FoldableWithIndex, Data.Functor, Data.FunctorWithIndex, Data.Map, Data.Map.Internal, Data.Monoid, Data.Newtype, Data.Ord, Data.Semigroup, Data.Set, Data.Show, Data.Traversable, Data.TraversableWithIndex, Prelude, Prim
// TO REQUIRE: Control.Alt, Control.Apply, Control.Bind, Control.Plus, Control.Semigroupoid, Data.Eq, Data.Foldable, Data.FoldableWithIndex, Data.Functor, Data.FunctorWithIndex, Data.Map, Data.Map.Internal, Data.Monoid, Data.Newtype, Data.Ord, Data.Semigroup, Data.Set, Data.Show, Data.Traversable, Data.TraversableWithIndex, Prelude
require_once __DIR__ . '/../Control.Alt/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Plus/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Foldable/index.php';
require_once __DIR__ . '/../Data.FoldableWithIndex/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.FunctorWithIndex/index.php';
require_once __DIR__ . '/../Data.Map/index.php';
require_once __DIR__ . '/../Data.Map.Internal/index.php';
require_once __DIR__ . '/../Data.Monoid/index.php';
require_once __DIR__ . '/../Data.Newtype/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Set/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
require_once __DIR__ . '/../Data.Traversable/index.php';
require_once __DIR__ . '/../Data.TraversableWithIndex/index.php';
require_once __DIR__ . '/../Prelude/index.php';

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




// Data_Map_SemigroupMap
function majData_majMap_majSemigroupmajMap($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majSemigroupmajMap';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Map_SemigroupMap'] = __NAMESPACE__ . '\\majData_majMap_majSemigroupmajMap';

// Data_Map_traversableWithIndexSemigroupMap
$GLOBALS['Data_Map_traversableWithIndexSemigroupMap'] = $GLOBALS['Data_Map_Internal_traversableWithIndexMap'];

// Data_Map_traversableSemigroupMap
$GLOBALS['Data_Map_traversableSemigroupMap'] = $GLOBALS['Data_Map_Internal_traversableMap'];

// Data_Map_showSemigroupMap
function majData_majMap_showmajSemigroupmajMap($dictShow_0, $dictShow1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_showmajSemigroupmajMap';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $showArray_2_0 = (object)["show" => ($GLOBALS['Data_Show_showArrayImpl'])(function($v_2) use ($dictShow1_1, $dictShow_0) {
  $__num = \func_num_args();
  $__res = (((("(Tuple " . (($dictShow_0)->{'show'})(($v_2)->{'value0'})) . " ") . (($dictShow1_1)->{'show'})(($v_2)->{'value1'})) . ")");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})];
  $__res = (object)["show" => function($as_3) use ($showArray_2_0) {
  $__num = \func_num_args();
  $__res = (("(fromFoldable " . (($showArray_2_0)->{'show'})(\Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(((((($GLOBALS['Data_Unfoldable_unfoldrArrayImpl'])($GLOBALS['Data_Maybe_isNothing']))(function($v_4) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v_4 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = ($v_4)->{'value0'};
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Tuple_fst']))($GLOBALS['Data_Tuple_snd']))($GLOBALS['Data_Map_Internal_stepUnfoldr']), $GLOBALS['Data_Map_Internal_toMapIter'], $as_3))) . ")");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Map_showSemigroupMap'] = __NAMESPACE__ . '\\majData_majMap_showmajSemigroupmajMap';

// Data_Map_semigroupSemigroupMap
function majData_majMap_semigroupmajSemigroupmajMap($dictOrd_0, $dictSemigroup_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_semigroupmajSemigroupmajMap';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $append_2_0 = ($dictSemigroup_1)->{'append'};
  $__res = (object)["append" => function($v_3) use ($append_2_0, $dictOrd_0) {
  $__num = \func_num_args();
  $__res = function($v1_4) use ($append_2_0, $dictOrd_0, $v_3) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Map_Internal_unsafeUnionWith'])(($dictOrd_0)->{'compare'}, $append_2_0, $v_3, $v1_4);
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
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Map_semigroupSemigroupMap'] = __NAMESPACE__ . '\\majData_majMap_semigroupmajSemigroupmajMap';

// Data_Map_plusSemigroupMap
function majData_majMap_plusmajSemigroupmajMap($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_plusmajSemigroupmajMap';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $compare_1_0 = ($dictOrd_0)->{'compare'};
  $altMap1_1_0 = (object)["alt" => function($m1_2) use ($compare_1_0) {
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
}, "Functor0" => function($_dollar___unused_1) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Map_Internal_functorMap'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["empty" => new \Data\Map\Internal\Data_Map_Internal_Leaf(), "Alt0" => function($_dollar___unused_2) use ($altMap1_1_0) {
  $__num = \func_num_args();
  $__res = $altMap1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Map_plusSemigroupMap'] = __NAMESPACE__ . '\\majData_majMap_plusmajSemigroupmajMap';

// Data_Map_ordSemigroupMap
function majData_majMap_ordmajSemigroupmajMap($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_ordmajSemigroupmajMap';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictOrd_0)->{'Eq0'})(null);
  $__local_var_2_1 = (($dictOrd_0)->{'Eq0'})(null);
  $__res = function($dictOrd1_3) use ($__local_var_1_0, $__local_var_2_1, $dictOrd_0) {
  $__num = \func_num_args();
  $__local_var_4_2 = (($dictOrd1_3)->{'Eq0'})(null);
  $go__go_5_3 = null;
  $go__go_5_3 = function($a_6) use ($__local_var_1_0, $__local_var_4_2, &$go__go_5_3) {
  $__num = \func_num_args();
  $__res = function($b_7) use ($__local_var_1_0, $__local_var_4_2, $a_6, &$go__go_5_3) {
  $__num = \func_num_args();
  $v_8_4 = \Data\Map\Internal\majData_majMap_majInternal_stepmajAsc($a_6);
  $__t5 = null;;
  if ($v_8_4 instanceof \Data\Map\Internal\Data_Map_Internal_IterNext) {
$v2_9_6 = \Data\Map\Internal\majData_majMap_majInternal_stepmajAsc($b_7);
$__t5 = ($v2_9_6 instanceof \Data\Map\Internal\Data_Map_Internal_IterNext && ((((($__local_var_1_0)->{'eq'})(($v_8_4)->{'value0'}))(($v2_9_6)->{'value0'}) && ((($__local_var_4_2)->{'eq'})(($v_8_4)->{'value1'}))(($v2_9_6)->{'value1'})) && (($go__go_5_3)(($v_8_4)->{'value2'}))(($v2_9_6)->{'value2'})));
goto end_branch_5;;
};
  if ($v_8_4 instanceof \Data\Map\Internal\Data_Map_Internal_IterDone) {
$__t5 = true;
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
};
  $eqMapIter2_4_2 = (object)["eq" => $go__go_5_3];
  $go__go_5_8 = null;
  $go__go_5_8 = (function() use ($dictOrd1_3, $dictOrd_0, &$go__go_5_8) {
  $__fn = function($a_6, $b_7 = null) use ($dictOrd1_3, $dictOrd_0, &$go__go_5_8, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_5_8_8_a_6 = $a_6;
  $__tco_var_go__go_5_8_8_b_7 = $b_7;
  tco_loop_go__go_5_8_8:;
  $a_6 = $__tco_var_go__go_5_8_8_a_6;
  $b_7 = $__tco_var_go__go_5_8_8_b_7;
  $v_8_8 = \Data\Map\Internal\majData_majMap_majInternal_stepmajAsc($b_7);
  $v1_9_9 = \Data\Map\Internal\majData_majMap_majInternal_stepmajAsc($a_6);
  $__t10 = null;;
  if ($v1_9_9 instanceof \Data\Map\Internal\Data_Map_Internal_IterNext) {
$__t11 = null;;
if ($v_8_8 instanceof \Data\Map\Internal\Data_Map_Internal_IterNext) {
$v3_10_12 = ((($dictOrd_0)->{'compare'})(($v1_9_9)->{'value0'}))(($v_8_8)->{'value0'});
$__t13 = null;;
if ($v3_10_12 instanceof \Data\Ordering\Data_Ordering_EQ) {
$v4_11_14 = ((($dictOrd1_3)->{'compare'})(($v1_9_9)->{'value1'}))(($v_8_8)->{'value1'});
$__t15 = null;;
if ($v4_11_14 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__tco_16 = ($v1_9_9)->{'value2'};
$__tco_17 = ($v_8_8)->{'value2'};
$__tco_var_go__go_5_8_8_a_6 = $__tco_16;
$__tco_var_go__go_5_8_8_b_7 = $__tco_17;
goto tco_loop_go__go_5_8_8;;
$__t15 = null;
goto end_branch_15;;
};
$__t15 = $v4_11_14;
end_branch_15:;
$__t13 = $__t15;
goto end_branch_13;;
};
$__t13 = $v3_10_12;
end_branch_13:;
$__t11 = $__t13;
goto end_branch_11;;
};
if ($v_8_8 instanceof \Data\Map\Internal\Data_Map_Internal_IterDone) {
$__t11 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_11;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t11 = null;
end_branch_11:;
$__t10 = $__t11;
goto end_branch_10;;
};
  if ($v1_9_9 instanceof \Data\Map\Internal\Data_Map_Internal_IterDone) {
$__t18 = null;;
if ($v_8_8 instanceof \Data\Map\Internal\Data_Map_Internal_IterDone) {
$__t18 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_18;;
};
$__t18 = new \Data\Ordering\Data_Ordering_LT();
end_branch_18:;
$__t10 = $__t18;
goto end_branch_10;;
};
  if ($v_8_8 instanceof \Data\Map\Internal\Data_Map_Internal_IterDone) {
$__t10 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_10;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t10 = null;
  end_branch_10:;
  $__res = $__t10;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $ordMapIter2_4_2 = (object)["compare" => $go__go_5_8, "Eq0" => function($_dollar___unused_5) use ($eqMapIter2_4_2) {
  $__num = \func_num_args();
  $__res = $eqMapIter2_4_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_5_10 = (($dictOrd1_3)->{'Eq0'})(null);
  $go__go_6_11 = null;
  $go__go_6_11 = function($a_7) use ($__local_var_2_1, $__local_var_5_10, &$go__go_6_11) {
  $__num = \func_num_args();
  $__res = function($b_8) use ($__local_var_2_1, $__local_var_5_10, $a_7, &$go__go_6_11) {
  $__num = \func_num_args();
  $v_9_12 = \Data\Map\Internal\majData_majMap_majInternal_stepmajAsc($a_7);
  $__t13 = null;;
  if ($v_9_12 instanceof \Data\Map\Internal\Data_Map_Internal_IterNext) {
$v2_10_14 = \Data\Map\Internal\majData_majMap_majInternal_stepmajAsc($b_8);
$__t13 = ($v2_10_14 instanceof \Data\Map\Internal\Data_Map_Internal_IterNext && ((((($__local_var_2_1)->{'eq'})(($v_9_12)->{'value0'}))(($v2_10_14)->{'value0'}) && ((($__local_var_5_10)->{'eq'})(($v_9_12)->{'value1'}))(($v2_10_14)->{'value1'})) && (($go__go_6_11)(($v_9_12)->{'value2'}))(($v2_10_14)->{'value2'})));
goto end_branch_13;;
};
  if ($v_9_12 instanceof \Data\Map\Internal\Data_Map_Internal_IterDone) {
$__t13 = true;
goto end_branch_13;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t13 = null;
  end_branch_13:;
  $__res = $__t13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $eqMapIter2_6_11 = (object)["eq" => $go__go_6_11];
  $eqMap2_5_10 = (object)["eq" => function($xs_7) use ($eqMapIter2_6_11) {
  $__num = \func_num_args();
  $__res = function($ys_8) use ($eqMapIter2_6_11, $xs_7) {
  $__num = \func_num_args();
  $__t16 = null;;
  if ($xs_7 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t16 = $ys_8 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf;
goto end_branch_16;;
};
  if ($xs_7 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t16 = ($ys_8 instanceof \Data\Map\Internal\Data_Map_Internal_Node && ((($xs_7)->{'value1'} === ($ys_8)->{'value1'}) && ((($eqMapIter2_6_11)->{'eq'})(new \Data\Map\Internal\Data_Map_Internal_IterNode($xs_7, new \Data\Map\Internal\Data_Map_Internal_IterLeaf())))(new \Data\Map\Internal\Data_Map_Internal_IterNode($ys_8, new \Data\Map\Internal\Data_Map_Internal_IterLeaf()))));
goto end_branch_16;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t16 = null;
  end_branch_16:;
  $__res = $__t16;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["compare" => function($xs_6) use ($ordMapIter2_4_2) {
  $__num = \func_num_args();
  $__res = function($ys_7) use ($ordMapIter2_4_2, $xs_6) {
  $__num = \func_num_args();
  $__t19 = null;;
  if ($xs_6 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t20 = null;;
if ($ys_7 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t20 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_20;;
};
$__t20 = new \Data\Ordering\Data_Ordering_LT();
end_branch_20:;
$__t19 = $__t20;
goto end_branch_19;;
};
  $__t18 = null;;
  if ($ys_7 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t18 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_18;;
};
  $__t18 = ((($ordMapIter2_4_2)->{'compare'})(new \Data\Map\Internal\Data_Map_Internal_IterNode($xs_6, new \Data\Map\Internal\Data_Map_Internal_IterLeaf())))(new \Data\Map\Internal\Data_Map_Internal_IterNode($ys_7, new \Data\Map\Internal\Data_Map_Internal_IterLeaf()));
  end_branch_18:;
  $__t19 = $__t18;
  end_branch_19:;
  $__res = $__t19;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq0" => function($_dollar___unused_6) use ($eqMap2_5_10) {
  $__num = \func_num_args();
  $__res = $eqMap2_5_10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Map_ordSemigroupMap'] = __NAMESPACE__ . '\\majData_majMap_ordmajSemigroupmajMap';

// Data_Map_ord1SemigroupMap
function majData_majMap_ord1majSemigroupmajMap($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_ord1majSemigroupmajMap';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictOrd_0)->{'Eq0'})(null);
  $__local_var_2_1 = (($dictOrd_0)->{'Eq0'})(null);
  $eq1Map1_2_1 = (object)["eq1" => function($dictEq1_3) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $go__go_4_2 = null;
  $go__go_4_2 = function($a_5) use ($__local_var_2_1, $dictEq1_3, &$go__go_4_2) {
  $__num = \func_num_args();
  $__res = function($b_6) use ($__local_var_2_1, $a_5, $dictEq1_3, &$go__go_4_2) {
  $__num = \func_num_args();
  $v_7_3 = \Data\Map\Internal\majData_majMap_majInternal_stepmajAsc($a_5);
  $__t4 = null;;
  if ($v_7_3 instanceof \Data\Map\Internal\Data_Map_Internal_IterNext) {
$v2_8_5 = \Data\Map\Internal\majData_majMap_majInternal_stepmajAsc($b_6);
$__t4 = ($v2_8_5 instanceof \Data\Map\Internal\Data_Map_Internal_IterNext && ((((($__local_var_2_1)->{'eq'})(($v_7_3)->{'value0'}))(($v2_8_5)->{'value0'}) && ((($dictEq1_3)->{'eq'})(($v_7_3)->{'value1'}))(($v2_8_5)->{'value1'})) && (($go__go_4_2)(($v_7_3)->{'value2'}))(($v2_8_5)->{'value2'})));
goto end_branch_4;;
};
  if ($v_7_3 instanceof \Data\Map\Internal\Data_Map_Internal_IterDone) {
$__t4 = true;
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $eqMapIter2_4_2 = (object)["eq" => $go__go_4_2];
  $__res = function($xs_5) use ($eqMapIter2_4_2) {
  $__num = \func_num_args();
  $__res = function($ys_6) use ($eqMapIter2_4_2, $xs_5) {
  $__num = \func_num_args();
  $__t7 = null;;
  if ($xs_5 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t7 = $ys_6 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf;
goto end_branch_7;;
};
  if ($xs_5 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t7 = ($ys_6 instanceof \Data\Map\Internal\Data_Map_Internal_Node && ((($xs_5)->{'value1'} === ($ys_6)->{'value1'}) && ((($eqMapIter2_4_2)->{'eq'})(new \Data\Map\Internal\Data_Map_Internal_IterNode($xs_5, new \Data\Map\Internal\Data_Map_Internal_IterLeaf())))(new \Data\Map\Internal\Data_Map_Internal_IterNode($ys_6, new \Data\Map\Internal\Data_Map_Internal_IterLeaf()))));
goto end_branch_7;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t7 = null;
  end_branch_7:;
  $__res = $__t7;
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
  $__res = (object)["compare1" => function($dictOrd1_3) use ($__local_var_1_0, $dictOrd_0) {
  $__num = \func_num_args();
  $__local_var_4_9 = (($dictOrd1_3)->{'Eq0'})(null);
  $go__go_5_10 = null;
  $go__go_5_10 = function($a_6) use ($__local_var_1_0, $__local_var_4_9, &$go__go_5_10) {
  $__num = \func_num_args();
  $__res = function($b_7) use ($__local_var_1_0, $__local_var_4_9, $a_6, &$go__go_5_10) {
  $__num = \func_num_args();
  $v_8_11 = \Data\Map\Internal\majData_majMap_majInternal_stepmajAsc($a_6);
  $__t12 = null;;
  if ($v_8_11 instanceof \Data\Map\Internal\Data_Map_Internal_IterNext) {
$v2_9_13 = \Data\Map\Internal\majData_majMap_majInternal_stepmajAsc($b_7);
$__t12 = ($v2_9_13 instanceof \Data\Map\Internal\Data_Map_Internal_IterNext && ((((($__local_var_1_0)->{'eq'})(($v_8_11)->{'value0'}))(($v2_9_13)->{'value0'}) && ((($__local_var_4_9)->{'eq'})(($v_8_11)->{'value1'}))(($v2_9_13)->{'value1'})) && (($go__go_5_10)(($v_8_11)->{'value2'}))(($v2_9_13)->{'value2'})));
goto end_branch_12;;
};
  if ($v_8_11 instanceof \Data\Map\Internal\Data_Map_Internal_IterDone) {
$__t12 = true;
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
};
  $eqMapIter2_4_9 = (object)["eq" => $go__go_5_10];
  $go__go_5_15 = null;
  $go__go_5_15 = (function() use ($dictOrd1_3, $dictOrd_0, &$go__go_5_15) {
  $__fn = function($a_6, $b_7 = null) use ($dictOrd1_3, $dictOrd_0, &$go__go_5_15, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_5_15_15_a_6 = $a_6;
  $__tco_var_go__go_5_15_15_b_7 = $b_7;
  tco_loop_go__go_5_15_15:;
  $a_6 = $__tco_var_go__go_5_15_15_a_6;
  $b_7 = $__tco_var_go__go_5_15_15_b_7;
  $v_8_15 = \Data\Map\Internal\majData_majMap_majInternal_stepmajAsc($b_7);
  $v1_9_16 = \Data\Map\Internal\majData_majMap_majInternal_stepmajAsc($a_6);
  $__t17 = null;;
  if ($v1_9_16 instanceof \Data\Map\Internal\Data_Map_Internal_IterNext) {
$__t18 = null;;
if ($v_8_15 instanceof \Data\Map\Internal\Data_Map_Internal_IterNext) {
$v3_10_19 = ((($dictOrd_0)->{'compare'})(($v1_9_16)->{'value0'}))(($v_8_15)->{'value0'});
$__t20 = null;;
if ($v3_10_19 instanceof \Data\Ordering\Data_Ordering_EQ) {
$v4_11_21 = ((($dictOrd1_3)->{'compare'})(($v1_9_16)->{'value1'}))(($v_8_15)->{'value1'});
$__t22 = null;;
if ($v4_11_21 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__tco_23 = ($v1_9_16)->{'value2'};
$__tco_24 = ($v_8_15)->{'value2'};
$__tco_var_go__go_5_15_15_a_6 = $__tco_23;
$__tco_var_go__go_5_15_15_b_7 = $__tco_24;
goto tco_loop_go__go_5_15_15;;
$__t22 = null;
goto end_branch_22;;
};
$__t22 = $v4_11_21;
end_branch_22:;
$__t20 = $__t22;
goto end_branch_20;;
};
$__t20 = $v3_10_19;
end_branch_20:;
$__t18 = $__t20;
goto end_branch_18;;
};
if ($v_8_15 instanceof \Data\Map\Internal\Data_Map_Internal_IterDone) {
$__t18 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_18;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t18 = null;
end_branch_18:;
$__t17 = $__t18;
goto end_branch_17;;
};
  if ($v1_9_16 instanceof \Data\Map\Internal\Data_Map_Internal_IterDone) {
$__t25 = null;;
if ($v_8_15 instanceof \Data\Map\Internal\Data_Map_Internal_IterDone) {
$__t25 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_25;;
};
$__t25 = new \Data\Ordering\Data_Ordering_LT();
end_branch_25:;
$__t17 = $__t25;
goto end_branch_17;;
};
  if ($v_8_15 instanceof \Data\Map\Internal\Data_Map_Internal_IterDone) {
$__t17 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_17;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t17 = null;
  end_branch_17:;
  $__res = $__t17;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $ordMapIter2_4_9 = (object)["compare" => $go__go_5_15, "Eq0" => function($_dollar___unused_5) use ($eqMapIter2_4_9) {
  $__num = \func_num_args();
  $__res = $eqMapIter2_4_9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = function($xs_5) use ($ordMapIter2_4_9) {
  $__num = \func_num_args();
  $__res = function($ys_6) use ($ordMapIter2_4_9, $xs_5) {
  $__num = \func_num_args();
  $__t18 = null;;
  if ($xs_5 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t19 = null;;
if ($ys_6 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t19 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_19;;
};
$__t19 = new \Data\Ordering\Data_Ordering_LT();
end_branch_19:;
$__t18 = $__t19;
goto end_branch_18;;
};
  $__t17 = null;;
  if ($ys_6 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t17 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_17;;
};
  $__t17 = ((($ordMapIter2_4_9)->{'compare'})(new \Data\Map\Internal\Data_Map_Internal_IterNode($xs_5, new \Data\Map\Internal\Data_Map_Internal_IterLeaf())))(new \Data\Map\Internal\Data_Map_Internal_IterNode($ys_6, new \Data\Map\Internal\Data_Map_Internal_IterLeaf()));
  end_branch_17:;
  $__t18 = $__t17;
  end_branch_18:;
  $__res = $__t18;
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
}, "Eq10" => function($_dollar___unused_3) use ($eq1Map1_2_1) {
  $__num = \func_num_args();
  $__res = $eq1Map1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Map_ord1SemigroupMap'] = __NAMESPACE__ . '\\majData_majMap_ord1majSemigroupmajMap';

// Data_Map_newtypeSemigroupMap
$GLOBALS['Data_Map_newtypeSemigroupMap'] = (object)["Coercible0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Map_monoidSemigroupMap
function majData_majMap_monoidmajSemigroupmajMap($dictOrd_0, $dictSemigroup_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_monoidmajSemigroupmajMap';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $append_2_0 = ($dictSemigroup_1)->{'append'};
  $semigroupSemigroupMap2_2_0 = (object)["append" => function($v_3) use ($append_2_0, $dictOrd_0) {
  $__num = \func_num_args();
  $__res = function($v1_4) use ($append_2_0, $dictOrd_0, $v_3) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Map_Internal_unsafeUnionWith'])(($dictOrd_0)->{'compare'}, $append_2_0, $v_3, $v1_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["mempty" => new \Data\Map\Internal\Data_Map_Internal_Leaf(), "Semigroup0" => function($_dollar___unused_3) use ($semigroupSemigroupMap2_2_0) {
  $__num = \func_num_args();
  $__res = $semigroupSemigroupMap2_2_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Map_monoidSemigroupMap'] = __NAMESPACE__ . '\\majData_majMap_monoidmajSemigroupmajMap';

// Data_Map_keys_closure
$GLOBALS['Data_Map_keys_closure'] = (function() use (&$__fn) {
$go__go_0_0 = null;
$go__go_0_0 = function($v_1) use (&$go__go_0_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v_1 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t1 = new \Data\Map\Internal\Data_Map_Internal_Leaf();
goto end_branch_1;;
};
  if ($v_1 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t1 = new \Data\Map\Internal\Data_Map_Internal_Node(($v_1)->{'value0'}, ($v_1)->{'value1'}, ($v_1)->{'value2'}, $GLOBALS['Data_Unit_unit'], ($go__go_0_0)(($v_1)->{'value4'}), ($go__go_0_0)(($v_1)->{'value5'}));
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
return (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Set_fromMap']))($go__go_0_0);
})();

// Data_Map_keys
function majData_majMap_keys($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_keys';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Map_keys_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Map_keys'] = __NAMESPACE__ . '\\majData_majMap_keys';

// Data_Map_functorWithIndexSemigroupMap
$GLOBALS['Data_Map_functorWithIndexSemigroupMap'] = $GLOBALS['Data_Map_Internal_functorWithIndexMap'];

// Data_Map_functorSemigroupMap
$GLOBALS['Data_Map_functorSemigroupMap'] = $GLOBALS['Data_Map_Internal_functorMap'];

// Data_Map_foldableWithIndexSemigroupMap
$GLOBALS['Data_Map_foldableWithIndexSemigroupMap'] = $GLOBALS['Data_Map_Internal_foldableWithIndexMap'];

// Data_Map_foldableSemigroupMap
$GLOBALS['Data_Map_foldableSemigroupMap'] = $GLOBALS['Data_Map_Internal_foldableMap'];

// Data_Map_eqSemigroupMap
function majData_majMap_eqmajSemigroupmajMap($dictEq_0, $dictEq1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_eqmajSemigroupmajMap';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__go_2_0 = null;
  $go__go_2_0 = function($a_3) use ($dictEq1_1, $dictEq_0, &$go__go_2_0) {
  $__num = \func_num_args();
  $__res = function($b_4) use ($a_3, $dictEq1_1, $dictEq_0, &$go__go_2_0) {
  $__num = \func_num_args();
  $v_5_1 = \Data\Map\Internal\majData_majMap_majInternal_stepmajAsc($a_3);
  $__t2 = null;;
  if ($v_5_1 instanceof \Data\Map\Internal\Data_Map_Internal_IterNext) {
$v2_6_3 = \Data\Map\Internal\majData_majMap_majInternal_stepmajAsc($b_4);
$__t2 = ($v2_6_3 instanceof \Data\Map\Internal\Data_Map_Internal_IterNext && ((((($dictEq_0)->{'eq'})(($v_5_1)->{'value0'}))(($v2_6_3)->{'value0'}) && ((($dictEq1_1)->{'eq'})(($v_5_1)->{'value1'}))(($v2_6_3)->{'value1'})) && (($go__go_2_0)(($v_5_1)->{'value2'}))(($v2_6_3)->{'value2'})));
goto end_branch_2;;
};
  if ($v_5_1 instanceof \Data\Map\Internal\Data_Map_Internal_IterDone) {
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
  $eqMapIter2_2_0 = (object)["eq" => $go__go_2_0];
  $__res = (object)["eq" => function($xs_3) use ($eqMapIter2_2_0) {
  $__num = \func_num_args();
  $__res = function($ys_4) use ($eqMapIter2_2_0, $xs_3) {
  $__num = \func_num_args();
  $__t5 = null;;
  if ($xs_3 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t5 = $ys_4 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf;
goto end_branch_5;;
};
  if ($xs_3 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t5 = ($ys_4 instanceof \Data\Map\Internal\Data_Map_Internal_Node && ((($xs_3)->{'value1'} === ($ys_4)->{'value1'}) && ((($eqMapIter2_2_0)->{'eq'})(new \Data\Map\Internal\Data_Map_Internal_IterNode($xs_3, new \Data\Map\Internal\Data_Map_Internal_IterLeaf())))(new \Data\Map\Internal\Data_Map_Internal_IterNode($ys_4, new \Data\Map\Internal\Data_Map_Internal_IterLeaf()))));
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
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Map_eqSemigroupMap'] = __NAMESPACE__ . '\\majData_majMap_eqmajSemigroupmajMap';

// Data_Map_eq1SemigroupMap
function majData_majMap_eq1majSemigroupmajMap($dictEq_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_eq1majSemigroupmajMap';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["eq1" => function($dictEq1_1) use ($dictEq_0) {
  $__num = \func_num_args();
  $go__go_2_0 = null;
  $go__go_2_0 = function($a_3) use ($dictEq1_1, $dictEq_0, &$go__go_2_0) {
  $__num = \func_num_args();
  $__res = function($b_4) use ($a_3, $dictEq1_1, $dictEq_0, &$go__go_2_0) {
  $__num = \func_num_args();
  $v_5_1 = \Data\Map\Internal\majData_majMap_majInternal_stepmajAsc($a_3);
  $__t2 = null;;
  if ($v_5_1 instanceof \Data\Map\Internal\Data_Map_Internal_IterNext) {
$v2_6_3 = \Data\Map\Internal\majData_majMap_majInternal_stepmajAsc($b_4);
$__t2 = ($v2_6_3 instanceof \Data\Map\Internal\Data_Map_Internal_IterNext && ((((($dictEq_0)->{'eq'})(($v_5_1)->{'value0'}))(($v2_6_3)->{'value0'}) && ((($dictEq1_1)->{'eq'})(($v_5_1)->{'value1'}))(($v2_6_3)->{'value1'})) && (($go__go_2_0)(($v_5_1)->{'value2'}))(($v2_6_3)->{'value2'})));
goto end_branch_2;;
};
  if ($v_5_1 instanceof \Data\Map\Internal\Data_Map_Internal_IterDone) {
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
  $eqMapIter2_2_0 = (object)["eq" => $go__go_2_0];
  $__res = function($xs_3) use ($eqMapIter2_2_0) {
  $__num = \func_num_args();
  $__res = function($ys_4) use ($eqMapIter2_2_0, $xs_3) {
  $__num = \func_num_args();
  $__t5 = null;;
  if ($xs_3 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t5 = $ys_4 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf;
goto end_branch_5;;
};
  if ($xs_3 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t5 = ($ys_4 instanceof \Data\Map\Internal\Data_Map_Internal_Node && ((($xs_3)->{'value1'} === ($ys_4)->{'value1'}) && ((($eqMapIter2_2_0)->{'eq'})(new \Data\Map\Internal\Data_Map_Internal_IterNode($xs_3, new \Data\Map\Internal\Data_Map_Internal_IterLeaf())))(new \Data\Map\Internal\Data_Map_Internal_IterNode($ys_4, new \Data\Map\Internal\Data_Map_Internal_IterLeaf()))));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Map_eq1SemigroupMap'] = __NAMESPACE__ . '\\majData_majMap_eq1majSemigroupmajMap';

// Data_Map_bindSemigroupMap
function majData_majMap_bindmajSemigroupmajMap($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_bindmajSemigroupmajMap';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $compare_1_0 = ($dictOrd_0)->{'compare'};
  $applyMap1_1_0 = (object)["apply" => function($m1_2) use ($compare_1_0) {
  $__num = \func_num_args();
  $__res = function($m2_3) use ($compare_1_0, $m1_2) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Map_Internal_unsafeIntersectionWith'])($compare_1_0, function($x_4) {
  $__num = \func_num_args();
  $__res = $x_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $m1_2, $m2_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_1) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Map_Internal_functorMap'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["bind" => function($m_2) use ($dictOrd_0) {
  $__num = \func_num_args();
  $__res = function($f_3) use ($dictOrd_0, $m_2) {
  $__num = \func_num_args();
  $go__go_4_2 = null;
  $go__go_4_2 = function($v_5) use ($dictOrd_0, $f_3, &$go__go_4_2) {
  $__num = \func_num_args();
  $__t3 = null;;
  if ($v_5 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t3 = new \Data\Map\Internal\Data_Map_Internal_Leaf();
goto end_branch_3;;
};
  if ($v_5 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__local_var_6_4 = ($v_5)->{'value2'};
$go__go_7_5 = null;
$go__go_7_5 = function($v_8) use ($__local_var_6_4, $dictOrd_0, &$go__go_7_5) {
  $__num = \func_num_args();
  $__tco_var_go__go_7_5_5_v_8 = $v_8;
  tco_loop_go__go_7_5_5:;
  $v_8 = $__tco_var_go__go_7_5_5_v_8;
  $__t5 = null;;
  if ($v_8 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t5 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_5;;
};
  if ($v_8 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$v1_9_6 = ((($dictOrd_0)->{'compare'})($__local_var_6_4))(($v_8)->{'value2'});
$__t7 = null;;
if ($v1_9_6 instanceof \Data\Ordering\Data_Ordering_LT) {
$__tco_8 = ($v_8)->{'value4'};
$__tco_var_go__go_7_5_5_v_8 = $__tco_8;
goto tco_loop_go__go_7_5_5;;
$__t7 = null;
goto end_branch_7;;
};
if ($v1_9_6 instanceof \Data\Ordering\Data_Ordering_GT) {
$__tco_9 = ($v_8)->{'value5'};
$__tco_var_go__go_7_5_5_v_8 = $__tco_9;
goto tco_loop_go__go_7_5_5;;
$__t7 = null;
goto end_branch_7;;
};
if ($v1_9_6 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t7 = new \Data\Maybe\Data_Maybe_Just(($v_8)->{'value3'});
goto end_branch_7;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t7 = null;
end_branch_7:;
$__t5 = $__t7;
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
$v2_6_4 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl($go__go_7_5, $f_3, ($v_5)->{'value3'});
$__t7 = null;;
if ($v2_6_4 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t7 = ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($v_5)->{'value2'}, ($v2_6_4)->{'value0'}, ($go__go_4_2)(($v_5)->{'value4'}), ($go__go_4_2)(($v_5)->{'value5'}));
goto end_branch_7;;
};
if ($v2_6_4 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t7 = ($GLOBALS['Data_Map_Internal_unsafeJoinNodes'])(($go__go_4_2)(($v_5)->{'value4'}), ($go__go_4_2)(($v_5)->{'value5'}));
goto end_branch_7;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t7 = null;
end_branch_7:;
$__t3 = $__t7;
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
  $__res = ($go__go_4_2)($m_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_2) use ($applyMap1_1_0) {
  $__num = \func_num_args();
  $__res = $applyMap1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Map_bindSemigroupMap'] = __NAMESPACE__ . '\\majData_majMap_bindmajSemigroupmajMap';

// Data_Map_applySemigroupMap
function majData_majMap_applymajSemigroupmajMap($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_applymajSemigroupmajMap';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $compare_1_0 = ($dictOrd_0)->{'compare'};
  $__res = (object)["apply" => function($m1_2) use ($compare_1_0) {
  $__num = \func_num_args();
  $__res = function($m2_3) use ($compare_1_0, $m1_2) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Map_Internal_unsafeIntersectionWith'])($compare_1_0, function($x_4) {
  $__num = \func_num_args();
  $__res = $x_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $m1_2, $m2_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_1) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Map_Internal_functorMap'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Map_applySemigroupMap'] = __NAMESPACE__ . '\\majData_majMap_applymajSemigroupmajMap';

// Data_Map_altSemigroupMap
function majData_majMap_altmajSemigroupmajMap($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_altmajSemigroupmajMap';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $compare_1_0 = ($dictOrd_0)->{'compare'};
  $__res = (object)["alt" => function($m1_2) use ($compare_1_0) {
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
}, "Functor0" => function($_dollar___unused_1) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Map_Internal_functorMap'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Map_altSemigroupMap'] = __NAMESPACE__ . '\\majData_majMap_altmajSemigroupmajMap';

