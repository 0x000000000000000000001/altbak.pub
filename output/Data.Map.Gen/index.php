<?php

namespace Data\Map\Gen;

// ALL IMPORTS: Control.Apply, Control.Bind, Control.Monad.Gen, Control.Monad.Gen.Class, Control.Monad.Rec.Class, Data.Function, Data.Functor, Data.List, Data.List.Types, Data.Map, Data.Map.Internal, Data.Tuple, Prelude, Prim
// TO REQUIRE: Control.Apply, Control.Bind, Control.Monad.Gen, Control.Monad.Gen.Class, Control.Monad.Rec.Class, Data.Function, Data.Functor, Data.List, Data.List.Types, Data.Map, Data.Map.Internal, Data.Tuple, Prelude
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Monad.Gen/index.php';
require_once __DIR__ . '/../Control.Monad.Gen.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Rec.Class/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.List/index.php';
require_once __DIR__ . '/../Data.List.Types/index.php';
require_once __DIR__ . '/../Data.Map/index.php';
require_once __DIR__ . '/../Data.Map.Internal/index.php';
require_once __DIR__ . '/../Data.Tuple/index.php';
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




// Data_Map_Gen_genMap
function majData_majMap_majGen_genmajMap($dictMonadRec_0, $dictMonadGen_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majGen_genmajMap';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $Monad0_2_0 = (($dictMonadGen_1)->{'Monad0'})(null);
  $Bind1_3_1 = (($Monad0_2_0)->{'Bind1'})(null);
  $Bind11_4_2 = (($Monad0_2_0)->{'Bind1'})(null);
  $Functor0_5_3 = (((($Bind11_4_2)->{'Apply0'})(null))->{'Functor0'})(null);
  $Apply0_6_4 = (($Bind11_4_2)->{'Apply0'})(null);
  $__res = function($dictOrd_7) use ($Apply0_6_4, $Bind1_3_1, $Functor0_5_3, $dictMonadGen_1, $dictMonadRec_0) {
  $__num = \func_num_args();
  $go__go_8_5 = null;
  $go__go_8_5 = (function() use ($dictOrd_7, &$go__go_8_5) {
  $__fn = function($b_9, $v_10 = null) use ($dictOrd_7, &$go__go_8_5, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_8_5_5_b_9 = $b_9;
  $__tco_var_go__go_8_5_5_v_10 = $v_10;
  tco_loop_go__go_8_5_5:;
  $b_9 = $__tco_var_go__go_8_5_5_b_9;
  $v_10 = $__tco_var_go__go_8_5_5_v_10;
  $__t5 = null;;
  if ($v_10 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t5 = $b_9;
goto end_branch_5;;
};
  if ($v_10 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__local_var_11_6 = (($v_10)->{'value0'})->{'value0'};
$__local_var_12_7 = (($v_10)->{'value0'})->{'value1'};
$go__go_13_8 = null;
$go__go_13_8 = function($v1_14) use ($__local_var_11_6, $__local_var_12_7, $dictOrd_7, &$go__go_13_8) {
  $__num = \func_num_args();
  $__t9 = null;;
  if ($v1_14 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t9 = new \Data\Map\Internal\Data_Map_Internal_Node(1, 1, $__local_var_11_6, $__local_var_12_7, new \Data\Map\Internal\Data_Map_Internal_Leaf(), new \Data\Map\Internal\Data_Map_Internal_Leaf());
goto end_branch_9;;
};
  if ($v1_14 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$v2_15_10 = ((($dictOrd_7)->{'compare'})($__local_var_11_6))(($v1_14)->{'value2'});
$__t11 = null;;
if ($v2_15_10 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t11 = ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($v1_14)->{'value2'}, ($v1_14)->{'value3'}, ($go__go_13_8)(($v1_14)->{'value4'}), ($v1_14)->{'value5'});
goto end_branch_11;;
};
if ($v2_15_10 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t11 = ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($v1_14)->{'value2'}, ($v1_14)->{'value3'}, ($v1_14)->{'value4'}, ($go__go_13_8)(($v1_14)->{'value5'}));
goto end_branch_11;;
};
if ($v2_15_10 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t11 = new \Data\Map\Internal\Data_Map_Internal_Node(($v1_14)->{'value0'}, ($v1_14)->{'value1'}, $__local_var_11_6, $__local_var_12_7, ($v1_14)->{'value4'}, ($v1_14)->{'value5'});
goto end_branch_11;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t11 = null;
end_branch_11:;
$__t9 = $__t11;
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
$__tco_12 = ($go__go_13_8)($b_9);
$__tco_13 = ($v_10)->{'value1'};
$__tco_var_go__go_8_5_5_b_9 = $__tco_12;
$__tco_var_go__go_8_5_5_v_10 = $__tco_13;
goto tco_loop_go__go_8_5_5;;
$__t5 = null;
goto end_branch_5;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t5 = null;
  end_branch_5:;
  $__res = $__t5;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $fromFoldable_8_5 = ($go__go_8_5)(new \Data\Map\Internal\Data_Map_Internal_Leaf());
  $__res = function($genKey_9) use ($Apply0_6_4, $Bind1_3_1, $Functor0_5_3, $dictMonadGen_1, $dictMonadRec_0, $fromFoldable_8_5) {
  $__num = \func_num_args();
  $__res = function($genValue_10) use ($Apply0_6_4, $Bind1_3_1, $Functor0_5_3, $dictMonadGen_1, $dictMonadRec_0, $fromFoldable_8_5, $genKey_9) {
  $__num = \func_num_args();
  $__res = (($dictMonadGen_1)->{'sized'})(function($size_11) use ($Apply0_6_4, $Bind1_3_1, $Functor0_5_3, $dictMonadGen_1, $dictMonadRec_0, $fromFoldable_8_5, $genKey_9, $genValue_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_1)->{'bind'})(((($dictMonadGen_1)->{'chooseInt'})(0))($size_11)))(function($newSize_12) use ($Apply0_6_4, $Functor0_5_3, $dictMonadGen_1, $dictMonadRec_0, $fromFoldable_8_5, $genKey_9, $genValue_10) {
  $__num = \func_num_args();
  $Monad0_13_7 = (($dictMonadGen_1)->{'Monad0'})(null);
  $pure_14_8 = ((($Monad0_13_7)->{'Applicative0'})(null))->{'pure'};
  $Bind1_15_9 = (($Monad0_13_7)->{'Bind1'})(null);
  $__local_var_16_10 = ((($Apply0_6_4)->{'apply'})(((($Functor0_5_3)->{'map'})($GLOBALS['Data_Tuple_Tuple']))($genKey_9)))($genValue_10);
  $__res = ((($dictMonadGen_1)->{'resize'})(function($v_13) use ($newSize_12) {
  $__num = \func_num_args();
  $__res = $newSize_12;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($Functor0_5_3)->{'map'})($fromFoldable_8_5))(((((((((($Monad0_13_7)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})(function($b_17) {
  $__num = \func_num_args();
  $go__go_18_11 = null;
  $go__go_18_11 = (function() use (&$go__go_18_11) {
  $__fn = function($source_19, $memo_20 = null) use (&$go__go_18_11, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_18_11_11_source_19 = $source_19;
  $__tco_var_go__go_18_11_11_memo_20 = $memo_20;
  tco_loop_go__go_18_11_11:;
  $source_19 = $__tco_var_go__go_18_11_11_source_19;
  $memo_20 = $__tco_var_go__go_18_11_11_memo_20;
  $__t11 = null;;
  if ($source_19 instanceof \Control\Monad\Gen\Control_Monad_Gen_Nil) {
$go__go_21_12 = null;
$go__go_21_12 = (function() use (&$__tco_var_go__go_18_11_11_source_19, &$__tco_var_go__go_18_11_11_memo_20, &$go__go_21_12) {
  $__fn = function($b_22, $v_23 = null) use (&$__tco_var_go__go_18_11_11_source_19, &$__tco_var_go__go_18_11_11_memo_20, &$go__go_21_12, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_21_12_12_b_22 = $b_22;
  $__tco_var_go__go_21_12_12_v_23 = $v_23;
  tco_loop_go__go_21_12_12:;
  $b_22 = $__tco_var_go__go_21_12_12_b_22;
  $v_23 = $__tco_var_go__go_21_12_12_v_23;
  $__t12 = null;;
  if ($v_23 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t12 = $b_22;
goto end_branch_12;;
};
  if ($v_23 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_13 = new \Data\List\Types\Data_List_Types_Cons(($v_23)->{'value0'}, $b_22);
$__tco_14 = ($v_23)->{'value1'};
$__tco_var_go__go_21_12_12_b_22 = $__tco_13;
$__tco_var_go__go_21_12_12_v_23 = $__tco_14;
goto tco_loop_go__go_21_12_12;;
$__t12 = null;
goto end_branch_12;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t12 = null;
  end_branch_12:;
  $__res = $__t12;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
$__t11 = (($go__go_21_12)(new \Data\List\Types\Data_List_Types_Nil()))($memo_20);
goto end_branch_11;;
};
  if ($source_19 instanceof \Control\Monad\Gen\Control_Monad_Gen_Cons) {
$__tco_13 = ($source_19)->{'value1'};
$__tco_14 = new \Data\List\Types\Data_List_Types_Cons(($source_19)->{'value0'}, $memo_20);
$__tco_var_go__go_18_11_11_source_19 = $__tco_13;
$__tco_var_go__go_18_11_11_memo_20 = $__tco_14;
goto tco_loop_go__go_18_11_11;;
$__t11 = null;
goto end_branch_11;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t11 = null;
  end_branch_11:;
  $__res = $__t11;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = (($go__go_18_11)($b_17))(new \Data\List\Types\Data_List_Types_Nil());
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($dictMonadGen_1)->{'sized'})((($GLOBALS['Control_Semigroupoid_composeImpl'])((($dictMonadRec_0)->{'tailRecM'})(function($v_17) use ($Bind1_15_9, $__local_var_16_10, $pure_14_8) {
  $__num = \func_num_args();
  $__t14 = null;;
  if ((($v_17)->{'value1'} <= 0)) {
$__t14 = ($pure_14_8)(new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Done(($v_17)->{'value0'}));
goto end_branch_14;;
};
  $__local_var_18_12 = ($v_17)->{'value0'};
  $__local_var_19_13 = ($v_17)->{'value1'};
  $__t14 = ((($Bind1_15_9)->{'bind'})($__local_var_16_10))(function($x_20) use ($__local_var_18_12, $__local_var_19_13, $pure_14_8) {
  $__num = \func_num_args();
  $__res = ($pure_14_8)(new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Loop(new \Data\Tuple\Data_Tuple_Tuple(new \Control\Monad\Gen\Control_Monad_Gen_Cons($x_20, $__local_var_18_12), ($__local_var_19_13 - 1))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  end_branch_14:;
  $__res = $__t14;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(($GLOBALS['Data_Tuple_Tuple'])(new \Control\Monad\Gen\Control_Monad_Gen_Nil()))))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Map_Gen_genMap'] = __NAMESPACE__ . '\\majData_majMap_majGen_genmajMap';

