<?php

namespace Data\String\Gen;

// ALL IMPORTS: Control.Bind, Control.Monad.Gen, Control.Monad.Gen.Class, Control.Monad.Rec.Class, Data.Char.Gen, Data.Function, Data.Functor, Data.Ord, Data.String.CodeUnits, Data.String.Gen, Data.Unfoldable, Prelude, Prim
// TO REQUIRE: Control.Bind, Control.Monad.Gen, Control.Monad.Gen.Class, Control.Monad.Rec.Class, Data.Char.Gen, Data.Function, Data.Functor, Data.Ord, Data.String.CodeUnits, Data.String.Gen, Data.Unfoldable, Prelude
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Monad.Gen/index.php';
require_once __DIR__ . '/../Control.Monad.Gen.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Rec.Class/index.php';
require_once __DIR__ . '/../Data.Char.Gen/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.String.CodeUnits/index.php';
require_once __DIR__ . '/../Data.String.Gen/index.php';
require_once __DIR__ . '/../Data.Unfoldable/index.php';
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




// Data_String_Gen_genString
function majData_majString_majGen_genmajString($dictMonadRec_0, $dictMonadGen_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majGen_genmajString';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $Monad0_2_0 = (($dictMonadGen_1)->{'Monad0'})(null);
  $Bind1_3_1 = (($Monad0_2_0)->{'Bind1'})(null);
  $Functor0_4_2 = (((((($Monad0_2_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($genChar_5) use ($Bind1_3_1, $Functor0_4_2, $dictMonadGen_1, $dictMonadRec_0) {
  $__num = \func_num_args();
  $__res = (($dictMonadGen_1)->{'sized'})(function($size_6) use ($Bind1_3_1, $Functor0_4_2, $dictMonadGen_1, $dictMonadRec_0, $genChar_5) {
  $__num = \func_num_args();
  $v_7_3 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), 1, $size_6);
  $__t4 = null;;
  if ($v_7_3 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t4 = $size_6;
goto end_branch_4;;
};
  if ($v_7_3 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t4 = 1;
goto end_branch_4;;
};
  if ($v_7_3 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t4 = 1;
goto end_branch_4;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t4 = null;
  end_branch_4:;
  $__res = ((($Bind1_3_1)->{'bind'})(((($dictMonadGen_1)->{'chooseInt'})(1))($__t4)))(function($newSize_7) use ($Functor0_4_2, $dictMonadGen_1, $dictMonadRec_0, $genChar_5) {
  $__num = \func_num_args();
  $Monad0_8_5 = (($dictMonadGen_1)->{'Monad0'})(null);
  $pure_9_6 = ((($Monad0_8_5)->{'Applicative0'})(null))->{'pure'};
  $Bind1_10_7 = (($Monad0_8_5)->{'Bind1'})(null);
  $__res = ((($dictMonadGen_1)->{'resize'})(function($v_8) use ($newSize_7) {
  $__num = \func_num_args();
  $__res = $newSize_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($Functor0_4_2)->{'map'})($GLOBALS['Data_String_CodeUnits_fromCharArray']))(((((((((($Monad0_8_5)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})(((((($GLOBALS['Data_Unfoldable_unfoldrArrayImpl'])($GLOBALS['Data_Maybe_isNothing']))(function($v_11) {
  $__num = \func_num_args();
  $__t8 = null;;
  if ($v_11 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t8 = ($v_11)->{'value0'};
goto end_branch_8;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t8 = null;
  end_branch_8:;
  $__res = $__t8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Tuple_fst']))($GLOBALS['Data_Tuple_snd']))(function($v_11) {
  $__num = \func_num_args();
  $__t9 = null;;
  if ($v_11 instanceof \Control\Monad\Gen\Control_Monad_Gen_Nil) {
$__t9 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_9;;
};
  if ($v_11 instanceof \Control\Monad\Gen\Control_Monad_Gen_Cons) {
$__t9 = new \Data\Maybe\Data_Maybe_Just(new \Data\Tuple\Data_Tuple_Tuple(($v_11)->{'value0'}, ($v_11)->{'value1'}));
goto end_branch_9;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t9 = null;
  end_branch_9:;
  $__res = $__t9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))((($dictMonadGen_1)->{'sized'})((($GLOBALS['Control_Semigroupoid_composeImpl'])((($dictMonadRec_0)->{'tailRecM'})(function($v_11) use ($Bind1_10_7, $genChar_5, $pure_9_6) {
  $__num = \func_num_args();
  $__t12 = null;;
  if ((($v_11)->{'value1'} <= 0)) {
$__t12 = ($pure_9_6)(new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Done(($v_11)->{'value0'}));
goto end_branch_12;;
};
  $__local_var_12_10 = ($v_11)->{'value0'};
  $__local_var_13_11 = ($v_11)->{'value1'};
  $__t12 = ((($Bind1_10_7)->{'bind'})($genChar_5))(function($x_14) use ($__local_var_12_10, $__local_var_13_11, $pure_9_6) {
  $__num = \func_num_args();
  $__res = ($pure_9_6)(new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Loop(new \Data\Tuple\Data_Tuple_Tuple(new \Control\Monad\Gen\Control_Monad_Gen_Cons($x_14, $__local_var_12_10), ($__local_var_13_11 - 1))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  end_branch_12:;
  $__res = $__t12;
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
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_Gen_genString'] = __NAMESPACE__ . '\\majData_majString_majGen_genmajString';

// Data_String_Gen_genUnicodeString
function majData_majString_majGen_genmajUnicodemajString($dictMonadRec_0, $dictMonadGen_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majGen_genmajUnicodemajString';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $Monad0_2_0 = (($dictMonadGen_1)->{'Monad0'})(null);
  $Bind1_3_1 = (($Monad0_2_0)->{'Bind1'})(null);
  $Functor0_4_2 = (((((($Monad0_2_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $__local_var_5_3 = ((((((((((($dictMonadGen_1)->{'Monad0'})(null))->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})(function($x_5) {
  $__num = \func_num_args();
  $v_6_3 = \Data\Enum\majData_majEnum_charmajTomajEnum($x_5);
  $__t4 = null;;
  if ($v_6_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t4 = ($v_6_3)->{'value0'};
goto end_branch_4;;
};
  if ($v_6_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t5 = null;;
if (($x_5 < \Data\Enum\majData_majEnum_tomajCharmajCode($GLOBALS['Data_Bounded_bottomChar']))) {
$__t5 = $GLOBALS['Data_Bounded_bottomChar'];
goto end_branch_5;;
};
$__t5 = $GLOBALS['Data_Bounded_topChar'];
end_branch_5:;
$__t4 = $__t5;
goto end_branch_4;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t4 = null;
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($dictMonadGen_1)->{'chooseInt'})(0))(65536));
  $__res = (($dictMonadGen_1)->{'sized'})(function($size_6) use ($Bind1_3_1, $Functor0_4_2, $__local_var_5_3, $dictMonadGen_1, $dictMonadRec_0) {
  $__num = \func_num_args();
  $v_7_7 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), 1, $size_6);
  $__t8 = null;;
  if ($v_7_7 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t8 = $size_6;
goto end_branch_8;;
};
  if ($v_7_7 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t8 = 1;
goto end_branch_8;;
};
  if ($v_7_7 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t8 = 1;
goto end_branch_8;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t8 = null;
  end_branch_8:;
  $__res = ((($Bind1_3_1)->{'bind'})(((($dictMonadGen_1)->{'chooseInt'})(1))($__t8)))(function($newSize_7) use ($Functor0_4_2, $__local_var_5_3, $dictMonadGen_1, $dictMonadRec_0) {
  $__num = \func_num_args();
  $Monad0_8_9 = (($dictMonadGen_1)->{'Monad0'})(null);
  $pure_9_10 = ((($Monad0_8_9)->{'Applicative0'})(null))->{'pure'};
  $Bind1_10_11 = (($Monad0_8_9)->{'Bind1'})(null);
  $__res = ((($dictMonadGen_1)->{'resize'})(function($v_8) use ($newSize_7) {
  $__num = \func_num_args();
  $__res = $newSize_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($Functor0_4_2)->{'map'})($GLOBALS['Data_String_CodeUnits_fromCharArray']))(((((((((($Monad0_8_9)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})(((((($GLOBALS['Data_Unfoldable_unfoldrArrayImpl'])($GLOBALS['Data_Maybe_isNothing']))(function($v_11) {
  $__num = \func_num_args();
  $__t12 = null;;
  if ($v_11 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t12 = ($v_11)->{'value0'};
goto end_branch_12;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t12 = null;
  end_branch_12:;
  $__res = $__t12;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Tuple_fst']))($GLOBALS['Data_Tuple_snd']))(function($v_11) {
  $__num = \func_num_args();
  $__t13 = null;;
  if ($v_11 instanceof \Control\Monad\Gen\Control_Monad_Gen_Nil) {
$__t13 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_13;;
};
  if ($v_11 instanceof \Control\Monad\Gen\Control_Monad_Gen_Cons) {
$__t13 = new \Data\Maybe\Data_Maybe_Just(new \Data\Tuple\Data_Tuple_Tuple(($v_11)->{'value0'}, ($v_11)->{'value1'}));
goto end_branch_13;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t13 = null;
  end_branch_13:;
  $__res = $__t13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))((($dictMonadGen_1)->{'sized'})((($GLOBALS['Control_Semigroupoid_composeImpl'])((($dictMonadRec_0)->{'tailRecM'})(function($v_11) use ($Bind1_10_11, $__local_var_5_3, $pure_9_10) {
  $__num = \func_num_args();
  $__t16 = null;;
  if ((($v_11)->{'value1'} <= 0)) {
$__t16 = ($pure_9_10)(new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Done(($v_11)->{'value0'}));
goto end_branch_16;;
};
  $__local_var_12_14 = ($v_11)->{'value0'};
  $__local_var_13_15 = ($v_11)->{'value1'};
  $__t16 = ((($Bind1_10_11)->{'bind'})($__local_var_5_3))(function($x_14) use ($__local_var_12_14, $__local_var_13_15, $pure_9_10) {
  $__num = \func_num_args();
  $__res = ($pure_9_10)(new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Loop(new \Data\Tuple\Data_Tuple_Tuple(new \Control\Monad\Gen\Control_Monad_Gen_Cons($x_14, $__local_var_12_14), ($__local_var_13_15 - 1))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  end_branch_16:;
  $__res = $__t16;
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
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_Gen_genUnicodeString'] = __NAMESPACE__ . '\\majData_majString_majGen_genmajUnicodemajString';

// Data_String_Gen_genDigitString
function majData_majString_majGen_genmajDigitmajString($dictMonadRec_0, $dictMonadGen_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majGen_genmajDigitmajString';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $Monad0_2_0 = (($dictMonadGen_1)->{'Monad0'})(null);
  $Bind1_3_1 = (($Monad0_2_0)->{'Bind1'})(null);
  $Functor0_4_2 = (((((($Monad0_2_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $__local_var_5_3 = ((((((((((($dictMonadGen_1)->{'Monad0'})(null))->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})(function($x_5) {
  $__num = \func_num_args();
  $v_6_3 = \Data\Enum\majData_majEnum_charmajTomajEnum($x_5);
  $__t4 = null;;
  if ($v_6_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t4 = ($v_6_3)->{'value0'};
goto end_branch_4;;
};
  if ($v_6_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t5 = null;;
if (($x_5 < \Data\Enum\majData_majEnum_tomajCharmajCode($GLOBALS['Data_Bounded_bottomChar']))) {
$__t5 = $GLOBALS['Data_Bounded_bottomChar'];
goto end_branch_5;;
};
$__t5 = $GLOBALS['Data_Bounded_topChar'];
end_branch_5:;
$__t4 = $__t5;
goto end_branch_4;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t4 = null;
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($dictMonadGen_1)->{'chooseInt'})(48))(57));
  $__res = (($dictMonadGen_1)->{'sized'})(function($size_6) use ($Bind1_3_1, $Functor0_4_2, $__local_var_5_3, $dictMonadGen_1, $dictMonadRec_0) {
  $__num = \func_num_args();
  $v_7_7 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), 1, $size_6);
  $__t8 = null;;
  if ($v_7_7 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t8 = $size_6;
goto end_branch_8;;
};
  if ($v_7_7 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t8 = 1;
goto end_branch_8;;
};
  if ($v_7_7 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t8 = 1;
goto end_branch_8;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t8 = null;
  end_branch_8:;
  $__res = ((($Bind1_3_1)->{'bind'})(((($dictMonadGen_1)->{'chooseInt'})(1))($__t8)))(function($newSize_7) use ($Functor0_4_2, $__local_var_5_3, $dictMonadGen_1, $dictMonadRec_0) {
  $__num = \func_num_args();
  $Monad0_8_9 = (($dictMonadGen_1)->{'Monad0'})(null);
  $pure_9_10 = ((($Monad0_8_9)->{'Applicative0'})(null))->{'pure'};
  $Bind1_10_11 = (($Monad0_8_9)->{'Bind1'})(null);
  $__res = ((($dictMonadGen_1)->{'resize'})(function($v_8) use ($newSize_7) {
  $__num = \func_num_args();
  $__res = $newSize_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($Functor0_4_2)->{'map'})($GLOBALS['Data_String_CodeUnits_fromCharArray']))(((((((((($Monad0_8_9)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})(((((($GLOBALS['Data_Unfoldable_unfoldrArrayImpl'])($GLOBALS['Data_Maybe_isNothing']))(function($v_11) {
  $__num = \func_num_args();
  $__t12 = null;;
  if ($v_11 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t12 = ($v_11)->{'value0'};
goto end_branch_12;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t12 = null;
  end_branch_12:;
  $__res = $__t12;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Tuple_fst']))($GLOBALS['Data_Tuple_snd']))(function($v_11) {
  $__num = \func_num_args();
  $__t13 = null;;
  if ($v_11 instanceof \Control\Monad\Gen\Control_Monad_Gen_Nil) {
$__t13 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_13;;
};
  if ($v_11 instanceof \Control\Monad\Gen\Control_Monad_Gen_Cons) {
$__t13 = new \Data\Maybe\Data_Maybe_Just(new \Data\Tuple\Data_Tuple_Tuple(($v_11)->{'value0'}, ($v_11)->{'value1'}));
goto end_branch_13;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t13 = null;
  end_branch_13:;
  $__res = $__t13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))((($dictMonadGen_1)->{'sized'})((($GLOBALS['Control_Semigroupoid_composeImpl'])((($dictMonadRec_0)->{'tailRecM'})(function($v_11) use ($Bind1_10_11, $__local_var_5_3, $pure_9_10) {
  $__num = \func_num_args();
  $__t16 = null;;
  if ((($v_11)->{'value1'} <= 0)) {
$__t16 = ($pure_9_10)(new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Done(($v_11)->{'value0'}));
goto end_branch_16;;
};
  $__local_var_12_14 = ($v_11)->{'value0'};
  $__local_var_13_15 = ($v_11)->{'value1'};
  $__t16 = ((($Bind1_10_11)->{'bind'})($__local_var_5_3))(function($x_14) use ($__local_var_12_14, $__local_var_13_15, $pure_9_10) {
  $__num = \func_num_args();
  $__res = ($pure_9_10)(new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Loop(new \Data\Tuple\Data_Tuple_Tuple(new \Control\Monad\Gen\Control_Monad_Gen_Cons($x_14, $__local_var_12_14), ($__local_var_13_15 - 1))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  end_branch_16:;
  $__res = $__t16;
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
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_Gen_genDigitString'] = __NAMESPACE__ . '\\majData_majString_majGen_genmajDigitmajString';

// Data_String_Gen_genAsciiString'
function majData_majString_majGen_genmajAsciimajString__prime__($dictMonadRec_0, $dictMonadGen_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majGen_genmajAsciimajString__prime__';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $Monad0_2_0 = (($dictMonadGen_1)->{'Monad0'})(null);
  $Bind1_3_1 = (($Monad0_2_0)->{'Bind1'})(null);
  $Functor0_4_2 = (((((($Monad0_2_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $__local_var_5_3 = ((((((((((($dictMonadGen_1)->{'Monad0'})(null))->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})(function($x_5) {
  $__num = \func_num_args();
  $v_6_3 = \Data\Enum\majData_majEnum_charmajTomajEnum($x_5);
  $__t4 = null;;
  if ($v_6_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t4 = ($v_6_3)->{'value0'};
goto end_branch_4;;
};
  if ($v_6_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t5 = null;;
if (($x_5 < \Data\Enum\majData_majEnum_tomajCharmajCode($GLOBALS['Data_Bounded_bottomChar']))) {
$__t5 = $GLOBALS['Data_Bounded_bottomChar'];
goto end_branch_5;;
};
$__t5 = $GLOBALS['Data_Bounded_topChar'];
end_branch_5:;
$__t4 = $__t5;
goto end_branch_4;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t4 = null;
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($dictMonadGen_1)->{'chooseInt'})(0))(127));
  $__res = (($dictMonadGen_1)->{'sized'})(function($size_6) use ($Bind1_3_1, $Functor0_4_2, $__local_var_5_3, $dictMonadGen_1, $dictMonadRec_0) {
  $__num = \func_num_args();
  $v_7_7 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), 1, $size_6);
  $__t8 = null;;
  if ($v_7_7 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t8 = $size_6;
goto end_branch_8;;
};
  if ($v_7_7 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t8 = 1;
goto end_branch_8;;
};
  if ($v_7_7 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t8 = 1;
goto end_branch_8;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t8 = null;
  end_branch_8:;
  $__res = ((($Bind1_3_1)->{'bind'})(((($dictMonadGen_1)->{'chooseInt'})(1))($__t8)))(function($newSize_7) use ($Functor0_4_2, $__local_var_5_3, $dictMonadGen_1, $dictMonadRec_0) {
  $__num = \func_num_args();
  $Monad0_8_9 = (($dictMonadGen_1)->{'Monad0'})(null);
  $pure_9_10 = ((($Monad0_8_9)->{'Applicative0'})(null))->{'pure'};
  $Bind1_10_11 = (($Monad0_8_9)->{'Bind1'})(null);
  $__res = ((($dictMonadGen_1)->{'resize'})(function($v_8) use ($newSize_7) {
  $__num = \func_num_args();
  $__res = $newSize_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($Functor0_4_2)->{'map'})($GLOBALS['Data_String_CodeUnits_fromCharArray']))(((((((((($Monad0_8_9)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})(((((($GLOBALS['Data_Unfoldable_unfoldrArrayImpl'])($GLOBALS['Data_Maybe_isNothing']))(function($v_11) {
  $__num = \func_num_args();
  $__t12 = null;;
  if ($v_11 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t12 = ($v_11)->{'value0'};
goto end_branch_12;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t12 = null;
  end_branch_12:;
  $__res = $__t12;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Tuple_fst']))($GLOBALS['Data_Tuple_snd']))(function($v_11) {
  $__num = \func_num_args();
  $__t13 = null;;
  if ($v_11 instanceof \Control\Monad\Gen\Control_Monad_Gen_Nil) {
$__t13 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_13;;
};
  if ($v_11 instanceof \Control\Monad\Gen\Control_Monad_Gen_Cons) {
$__t13 = new \Data\Maybe\Data_Maybe_Just(new \Data\Tuple\Data_Tuple_Tuple(($v_11)->{'value0'}, ($v_11)->{'value1'}));
goto end_branch_13;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t13 = null;
  end_branch_13:;
  $__res = $__t13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))((($dictMonadGen_1)->{'sized'})((($GLOBALS['Control_Semigroupoid_composeImpl'])((($dictMonadRec_0)->{'tailRecM'})(function($v_11) use ($Bind1_10_11, $__local_var_5_3, $pure_9_10) {
  $__num = \func_num_args();
  $__t16 = null;;
  if ((($v_11)->{'value1'} <= 0)) {
$__t16 = ($pure_9_10)(new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Done(($v_11)->{'value0'}));
goto end_branch_16;;
};
  $__local_var_12_14 = ($v_11)->{'value0'};
  $__local_var_13_15 = ($v_11)->{'value1'};
  $__t16 = ((($Bind1_10_11)->{'bind'})($__local_var_5_3))(function($x_14) use ($__local_var_12_14, $__local_var_13_15, $pure_9_10) {
  $__num = \func_num_args();
  $__res = ($pure_9_10)(new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Loop(new \Data\Tuple\Data_Tuple_Tuple(new \Control\Monad\Gen\Control_Monad_Gen_Cons($x_14, $__local_var_12_14), ($__local_var_13_15 - 1))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  end_branch_16:;
  $__res = $__t16;
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
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_Gen_genAsciiString__prime__'] = __NAMESPACE__ . '\\majData_majString_majGen_genmajAsciimajString__prime__';

// Data_String_Gen_genAsciiString
function majData_majString_majGen_genmajAsciimajString($dictMonadRec_0, $dictMonadGen_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majGen_genmajAsciimajString';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $Monad0_2_0 = (($dictMonadGen_1)->{'Monad0'})(null);
  $Bind1_3_1 = (($Monad0_2_0)->{'Bind1'})(null);
  $Functor0_4_2 = (((((($Monad0_2_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $__local_var_5_3 = ((((((((((($dictMonadGen_1)->{'Monad0'})(null))->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})(function($x_5) {
  $__num = \func_num_args();
  $v_6_3 = \Data\Enum\majData_majEnum_charmajTomajEnum($x_5);
  $__t4 = null;;
  if ($v_6_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t4 = ($v_6_3)->{'value0'};
goto end_branch_4;;
};
  if ($v_6_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t5 = null;;
if (($x_5 < \Data\Enum\majData_majEnum_tomajCharmajCode($GLOBALS['Data_Bounded_bottomChar']))) {
$__t5 = $GLOBALS['Data_Bounded_bottomChar'];
goto end_branch_5;;
};
$__t5 = $GLOBALS['Data_Bounded_topChar'];
end_branch_5:;
$__t4 = $__t5;
goto end_branch_4;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t4 = null;
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($dictMonadGen_1)->{'chooseInt'})(32))(127));
  $__res = (($dictMonadGen_1)->{'sized'})(function($size_6) use ($Bind1_3_1, $Functor0_4_2, $__local_var_5_3, $dictMonadGen_1, $dictMonadRec_0) {
  $__num = \func_num_args();
  $v_7_7 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), 1, $size_6);
  $__t8 = null;;
  if ($v_7_7 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t8 = $size_6;
goto end_branch_8;;
};
  if ($v_7_7 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t8 = 1;
goto end_branch_8;;
};
  if ($v_7_7 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t8 = 1;
goto end_branch_8;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t8 = null;
  end_branch_8:;
  $__res = ((($Bind1_3_1)->{'bind'})(((($dictMonadGen_1)->{'chooseInt'})(1))($__t8)))(function($newSize_7) use ($Functor0_4_2, $__local_var_5_3, $dictMonadGen_1, $dictMonadRec_0) {
  $__num = \func_num_args();
  $Monad0_8_9 = (($dictMonadGen_1)->{'Monad0'})(null);
  $pure_9_10 = ((($Monad0_8_9)->{'Applicative0'})(null))->{'pure'};
  $Bind1_10_11 = (($Monad0_8_9)->{'Bind1'})(null);
  $__res = ((($dictMonadGen_1)->{'resize'})(function($v_8) use ($newSize_7) {
  $__num = \func_num_args();
  $__res = $newSize_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($Functor0_4_2)->{'map'})($GLOBALS['Data_String_CodeUnits_fromCharArray']))(((((((((($Monad0_8_9)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})(((((($GLOBALS['Data_Unfoldable_unfoldrArrayImpl'])($GLOBALS['Data_Maybe_isNothing']))(function($v_11) {
  $__num = \func_num_args();
  $__t12 = null;;
  if ($v_11 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t12 = ($v_11)->{'value0'};
goto end_branch_12;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t12 = null;
  end_branch_12:;
  $__res = $__t12;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Tuple_fst']))($GLOBALS['Data_Tuple_snd']))(function($v_11) {
  $__num = \func_num_args();
  $__t13 = null;;
  if ($v_11 instanceof \Control\Monad\Gen\Control_Monad_Gen_Nil) {
$__t13 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_13;;
};
  if ($v_11 instanceof \Control\Monad\Gen\Control_Monad_Gen_Cons) {
$__t13 = new \Data\Maybe\Data_Maybe_Just(new \Data\Tuple\Data_Tuple_Tuple(($v_11)->{'value0'}, ($v_11)->{'value1'}));
goto end_branch_13;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t13 = null;
  end_branch_13:;
  $__res = $__t13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))((($dictMonadGen_1)->{'sized'})((($GLOBALS['Control_Semigroupoid_composeImpl'])((($dictMonadRec_0)->{'tailRecM'})(function($v_11) use ($Bind1_10_11, $__local_var_5_3, $pure_9_10) {
  $__num = \func_num_args();
  $__t16 = null;;
  if ((($v_11)->{'value1'} <= 0)) {
$__t16 = ($pure_9_10)(new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Done(($v_11)->{'value0'}));
goto end_branch_16;;
};
  $__local_var_12_14 = ($v_11)->{'value0'};
  $__local_var_13_15 = ($v_11)->{'value1'};
  $__t16 = ((($Bind1_10_11)->{'bind'})($__local_var_5_3))(function($x_14) use ($__local_var_12_14, $__local_var_13_15, $pure_9_10) {
  $__num = \func_num_args();
  $__res = ($pure_9_10)(new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Loop(new \Data\Tuple\Data_Tuple_Tuple(new \Control\Monad\Gen\Control_Monad_Gen_Cons($x_14, $__local_var_12_14), ($__local_var_13_15 - 1))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  end_branch_16:;
  $__res = $__t16;
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
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_Gen_genAsciiString'] = __NAMESPACE__ . '\\majData_majString_majGen_genmajAsciimajString';

// Data_String_Gen_genAlphaUppercaseString
function majData_majString_majGen_genmajAlphamajUppercasemajString($dictMonadRec_0, $dictMonadGen_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majGen_genmajAlphamajUppercasemajString';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $Monad0_2_0 = (($dictMonadGen_1)->{'Monad0'})(null);
  $Bind1_3_1 = (($Monad0_2_0)->{'Bind1'})(null);
  $Functor0_4_2 = (((((($Monad0_2_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $__local_var_5_3 = ((((((((((($dictMonadGen_1)->{'Monad0'})(null))->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})(function($x_5) {
  $__num = \func_num_args();
  $v_6_3 = \Data\Enum\majData_majEnum_charmajTomajEnum($x_5);
  $__t4 = null;;
  if ($v_6_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t4 = ($v_6_3)->{'value0'};
goto end_branch_4;;
};
  if ($v_6_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t5 = null;;
if (($x_5 < \Data\Enum\majData_majEnum_tomajCharmajCode($GLOBALS['Data_Bounded_bottomChar']))) {
$__t5 = $GLOBALS['Data_Bounded_bottomChar'];
goto end_branch_5;;
};
$__t5 = $GLOBALS['Data_Bounded_topChar'];
end_branch_5:;
$__t4 = $__t5;
goto end_branch_4;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t4 = null;
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($dictMonadGen_1)->{'chooseInt'})(65))(90));
  $__res = (($dictMonadGen_1)->{'sized'})(function($size_6) use ($Bind1_3_1, $Functor0_4_2, $__local_var_5_3, $dictMonadGen_1, $dictMonadRec_0) {
  $__num = \func_num_args();
  $v_7_7 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), 1, $size_6);
  $__t8 = null;;
  if ($v_7_7 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t8 = $size_6;
goto end_branch_8;;
};
  if ($v_7_7 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t8 = 1;
goto end_branch_8;;
};
  if ($v_7_7 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t8 = 1;
goto end_branch_8;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t8 = null;
  end_branch_8:;
  $__res = ((($Bind1_3_1)->{'bind'})(((($dictMonadGen_1)->{'chooseInt'})(1))($__t8)))(function($newSize_7) use ($Functor0_4_2, $__local_var_5_3, $dictMonadGen_1, $dictMonadRec_0) {
  $__num = \func_num_args();
  $Monad0_8_9 = (($dictMonadGen_1)->{'Monad0'})(null);
  $pure_9_10 = ((($Monad0_8_9)->{'Applicative0'})(null))->{'pure'};
  $Bind1_10_11 = (($Monad0_8_9)->{'Bind1'})(null);
  $__res = ((($dictMonadGen_1)->{'resize'})(function($v_8) use ($newSize_7) {
  $__num = \func_num_args();
  $__res = $newSize_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($Functor0_4_2)->{'map'})($GLOBALS['Data_String_CodeUnits_fromCharArray']))(((((((((($Monad0_8_9)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})(((((($GLOBALS['Data_Unfoldable_unfoldrArrayImpl'])($GLOBALS['Data_Maybe_isNothing']))(function($v_11) {
  $__num = \func_num_args();
  $__t12 = null;;
  if ($v_11 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t12 = ($v_11)->{'value0'};
goto end_branch_12;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t12 = null;
  end_branch_12:;
  $__res = $__t12;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Tuple_fst']))($GLOBALS['Data_Tuple_snd']))(function($v_11) {
  $__num = \func_num_args();
  $__t13 = null;;
  if ($v_11 instanceof \Control\Monad\Gen\Control_Monad_Gen_Nil) {
$__t13 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_13;;
};
  if ($v_11 instanceof \Control\Monad\Gen\Control_Monad_Gen_Cons) {
$__t13 = new \Data\Maybe\Data_Maybe_Just(new \Data\Tuple\Data_Tuple_Tuple(($v_11)->{'value0'}, ($v_11)->{'value1'}));
goto end_branch_13;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t13 = null;
  end_branch_13:;
  $__res = $__t13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))((($dictMonadGen_1)->{'sized'})((($GLOBALS['Control_Semigroupoid_composeImpl'])((($dictMonadRec_0)->{'tailRecM'})(function($v_11) use ($Bind1_10_11, $__local_var_5_3, $pure_9_10) {
  $__num = \func_num_args();
  $__t16 = null;;
  if ((($v_11)->{'value1'} <= 0)) {
$__t16 = ($pure_9_10)(new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Done(($v_11)->{'value0'}));
goto end_branch_16;;
};
  $__local_var_12_14 = ($v_11)->{'value0'};
  $__local_var_13_15 = ($v_11)->{'value1'};
  $__t16 = ((($Bind1_10_11)->{'bind'})($__local_var_5_3))(function($x_14) use ($__local_var_12_14, $__local_var_13_15, $pure_9_10) {
  $__num = \func_num_args();
  $__res = ($pure_9_10)(new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Loop(new \Data\Tuple\Data_Tuple_Tuple(new \Control\Monad\Gen\Control_Monad_Gen_Cons($x_14, $__local_var_12_14), ($__local_var_13_15 - 1))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  end_branch_16:;
  $__res = $__t16;
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
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_Gen_genAlphaUppercaseString'] = __NAMESPACE__ . '\\majData_majString_majGen_genmajAlphamajUppercasemajString';

// Data_String_Gen_genAlphaString
function majData_majString_majGen_genmajAlphamajString($dictMonadRec_0, $dictMonadGen_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majGen_genmajAlphamajString';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $Monad0_2_0 = (($dictMonadGen_1)->{'Monad0'})(null);
  $Bind1_3_1 = (($Monad0_2_0)->{'Bind1'})(null);
  $Functor0_4_2 = (((((($Monad0_2_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $foldableNonEmpty1_5_3 = (object)["foldMap" => function($dictMonoid_5) {
  $__num = \func_num_args();
  $Semigroup0_6_3 = (($dictMonoid_5)->{'Semigroup0'})(null);
  $__res = function($f_7) use ($Semigroup0_6_3, $dictMonoid_5) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($Semigroup0_6_3, $dictMonoid_5, $f_7) {
  $__num = \func_num_args();
  $Semigroup0_9_4 = (($dictMonoid_5)->{'Semigroup0'})(null);
  $__res = ((($Semigroup0_6_3)->{'append'})(($f_7)(($v_8)->{'value0'})))(\Data\Foldable\majData_majFoldable_foldrmajArray(function($x_10) use ($Semigroup0_9_4, $f_7) {
  $__num = \func_num_args();
  $__res = function($acc_11) use ($Semigroup0_9_4, $f_7, $x_10) {
  $__num = \func_num_args();
  $__res = ((($Semigroup0_9_4)->{'append'})(($f_7)($x_10)))($acc_11);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($dictMonoid_5)->{'mempty'}, ($v_8)->{'value1'}));
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
}, "foldl" => function($f_5) {
  $__num = \func_num_args();
  $__res = function($b_6) use ($f_5) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($b_6, $f_5) {
  $__num = \func_num_args();
  $__res = \Data\Foldable\majData_majFoldable_foldlmajArray($f_5, (($f_5)($b_6))(($v_7)->{'value0'}), ($v_7)->{'value1'});
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
}, "foldr" => function($f_5) {
  $__num = \func_num_args();
  $__res = function($b_6) use ($f_5) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($b_6, $f_5) {
  $__num = \func_num_args();
  $__res = (($f_5)(($v_7)->{'value0'}))(\Data\Foldable\majData_majFoldable_foldrmajArray($f_5, $b_6, ($v_7)->{'value1'}));
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
  $__local_var_5_3 = (object)["foldMap1" => function($dictSemigroup_6) {
  $__num = \func_num_args();
  $__res = function($f_7) use ($dictSemigroup_6) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($dictSemigroup_6, $f_7) {
  $__num = \func_num_args();
  $__res = \Data\Foldable\majData_majFoldable_foldlmajArray(function($s_9) use ($dictSemigroup_6, $f_7) {
  $__num = \func_num_args();
  $__res = function($a1_10) use ($dictSemigroup_6, $f_7, $s_9) {
  $__num = \func_num_args();
  $__res = ((($dictSemigroup_6)->{'append'})($s_9))(($f_7)($a1_10));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($f_7)(($v_8)->{'value0'}), ($v_8)->{'value1'});
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
}, "foldr1" => function($f_6) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($f_6) {
  $__num = \func_num_args();
  $__local_var_8_6 = ($f_6)(($v_7)->{'value0'});
  $__local_var_9_7 = \Data\Foldable\majData_majFoldable_foldrmajArray(function($a1_9) use ($f_6) {
  $__num = \func_num_args();
  $__local_var_10_7 = ($f_6)($a1_9);
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Maybe_Just']))(function($v2_11) use ($__local_var_10_7, $a1_9) {
  $__num = \func_num_args();
  $__t8 = null;;
  if ($v2_11 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t8 = $a1_9;
goto end_branch_8;;
};
  if ($v2_11 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t8 = ($__local_var_10_7)(($v2_11)->{'value0'});
goto end_branch_8;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t8 = null;
  end_branch_8:;
  $__res = $__t8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, new \Data\Maybe\Data_Maybe_Nothing(), ($v_7)->{'value1'});
  $__t10 = null;;
  if ($__local_var_9_7 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t10 = ($v_7)->{'value0'};
goto end_branch_10;;
};
  if ($__local_var_9_7 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t10 = ($__local_var_8_6)(($__local_var_9_7)->{'value0'});
goto end_branch_10;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t10 = null;
  end_branch_10:;
  $__res = $__t10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldl1" => function($f_6) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($f_6) {
  $__num = \func_num_args();
  $__res = \Data\Foldable\majData_majFoldable_foldlmajArray($f_6, ($v_7)->{'value0'}, ($v_7)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable0" => function($_dollar___unused_6) use ($foldableNonEmpty1_5_3) {
  $__num = \func_num_args();
  $__res = $foldableNonEmpty1_5_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_6_12 = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(((((((((((($dictMonadGen_1)->{'Monad0'})(null))->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})(function($x_6) {
  $__num = \func_num_args();
  $v_7_12 = \Data\Enum\majData_majEnum_charmajTomajEnum($x_6);
  $__t13 = null;;
  if ($v_7_12 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t13 = ($v_7_12)->{'value0'};
goto end_branch_13;;
};
  if ($v_7_12 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t14 = null;;
if (($x_6 < \Data\Enum\majData_majEnum_tomajCharmajCode($GLOBALS['Data_Bounded_bottomChar']))) {
$__t14 = $GLOBALS['Data_Bounded_bottomChar'];
goto end_branch_14;;
};
$__t14 = $GLOBALS['Data_Bounded_topChar'];
end_branch_14:;
$__t13 = $__t14;
goto end_branch_13;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t13 = null;
  end_branch_13:;
  $__res = $__t13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($dictMonadGen_1)->{'chooseInt'})(97))(122)), [((((((((((($dictMonadGen_1)->{'Monad0'})(null))->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})(function($x_6) {
  $__num = \func_num_args();
  $v_7_15 = \Data\Enum\majData_majEnum_charmajTomajEnum($x_6);
  $__t16 = null;;
  if ($v_7_15 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t16 = ($v_7_15)->{'value0'};
goto end_branch_16;;
};
  if ($v_7_15 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t17 = null;;
if (($x_6 < \Data\Enum\majData_majEnum_tomajCharmajCode($GLOBALS['Data_Bounded_bottomChar']))) {
$__t17 = $GLOBALS['Data_Bounded_bottomChar'];
goto end_branch_17;;
};
$__t17 = $GLOBALS['Data_Bounded_topChar'];
end_branch_17:;
$__t16 = $__t17;
goto end_branch_16;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t16 = null;
  end_branch_16:;
  $__res = $__t16;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($dictMonadGen_1)->{'chooseInt'})(65))(90))]);
  $__local_var_5_3 = ((((((($dictMonadGen_1)->{'Monad0'})(null))->{'Bind1'})(null))->{'bind'})(((($dictMonadGen_1)->{'chooseInt'})(0))(((((((($__local_var_5_3)->{'Foldable0'})(null))->{'foldl'})(function($c_7) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($c_7) {
  $__num = \func_num_args();
  $__res = (1 + $c_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(0))($__local_var_6_12) - 1))))(function($n_7) use ($__local_var_5_3, $__local_var_6_12) {
  $__num = \func_num_args();
  $go__go_8_19 = null;
  $go__go_8_19 = (function() use ($__local_var_5_3, $__local_var_6_12, &$go__go_8_19) {
  $__fn = function(int $v_9, $v1_10 = null) use ($__local_var_5_3, $__local_var_6_12, &$go__go_8_19, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_8_19_19_v_9 = $v_9;
  $__tco_var_go__go_8_19_19_v1_10 = $v1_10;
  tco_loop_go__go_8_19_19:;
  $v_9 = $__tco_var_go__go_8_19_19_v_9;
  $v1_10 = $__tco_var_go__go_8_19_19_v1_10;
  $__t19 = null;;
  if ($v1_10 instanceof \Control\Monad\Gen\Control_Monad_Gen_Cons) {
$__t22 = null;;
if (($v1_10)->{'value1'} instanceof \Control\Monad\Gen\Control_Monad_Gen_Nil) {
$__t22 = ($v1_10)->{'value0'};
goto end_branch_22;;
};
if (($v_9 <= 0)) {
$__t22 = ($v1_10)->{'value0'};
goto end_branch_22;;
};
$__tco_20 = ($v_9 - 1);
$__tco_21 = ($v1_10)->{'value1'};
$__tco_var_go__go_8_19_19_v_9 = $__tco_20;
$__tco_var_go__go_8_19_19_v1_10 = $__tco_21;
goto tco_loop_go__go_8_19_19;;
$__t22 = null;
end_branch_22:;
$__t19 = $__t22;
goto end_branch_19;;
};
  if ($v1_10 instanceof \Control\Monad\Gen\Control_Monad_Gen_Nil) {
$__t19 = (((($__local_var_5_3)->{'foldMap1'})($GLOBALS['Data_Semigroup_Last_semigroupLast']))(function($x_11) {
  $__num = \func_num_args();
  $__res = $x_11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($__local_var_6_12);
goto end_branch_19;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t19 = null;
  end_branch_19:;
  $__res = $__t19;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = (($go__go_8_19)($n_7))((((((($__local_var_5_3)->{'Foldable0'})(null))->{'foldr'})($GLOBALS['Control_Monad_Gen_Cons']))(new \Control\Monad\Gen\Control_Monad_Gen_Nil()))($__local_var_6_12));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = (($dictMonadGen_1)->{'sized'})(function($size_6) use ($Bind1_3_1, $Functor0_4_2, $__local_var_5_3, $dictMonadGen_1, $dictMonadRec_0) {
  $__num = \func_num_args();
  $v_7_21 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), 1, $size_6);
  $__t22 = null;;
  if ($v_7_21 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t22 = $size_6;
goto end_branch_22;;
};
  if ($v_7_21 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t22 = 1;
goto end_branch_22;;
};
  if ($v_7_21 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t22 = 1;
goto end_branch_22;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t22 = null;
  end_branch_22:;
  $__res = ((($Bind1_3_1)->{'bind'})(((($dictMonadGen_1)->{'chooseInt'})(1))($__t22)))(function($newSize_7) use ($Functor0_4_2, $__local_var_5_3, $dictMonadGen_1, $dictMonadRec_0) {
  $__num = \func_num_args();
  $Monad0_8_23 = (($dictMonadGen_1)->{'Monad0'})(null);
  $pure_9_24 = ((($Monad0_8_23)->{'Applicative0'})(null))->{'pure'};
  $Bind1_10_25 = (($Monad0_8_23)->{'Bind1'})(null);
  $__res = ((($dictMonadGen_1)->{'resize'})(function($v_8) use ($newSize_7) {
  $__num = \func_num_args();
  $__res = $newSize_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($Functor0_4_2)->{'map'})($GLOBALS['Data_String_CodeUnits_fromCharArray']))(((((((((($Monad0_8_23)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})(((((($GLOBALS['Data_Unfoldable_unfoldrArrayImpl'])($GLOBALS['Data_Maybe_isNothing']))(function($v_11) {
  $__num = \func_num_args();
  $__t26 = null;;
  if ($v_11 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t26 = ($v_11)->{'value0'};
goto end_branch_26;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t26 = null;
  end_branch_26:;
  $__res = $__t26;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Tuple_fst']))($GLOBALS['Data_Tuple_snd']))(function($v_11) {
  $__num = \func_num_args();
  $__t27 = null;;
  if ($v_11 instanceof \Control\Monad\Gen\Control_Monad_Gen_Nil) {
$__t27 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_27;;
};
  if ($v_11 instanceof \Control\Monad\Gen\Control_Monad_Gen_Cons) {
$__t27 = new \Data\Maybe\Data_Maybe_Just(new \Data\Tuple\Data_Tuple_Tuple(($v_11)->{'value0'}, ($v_11)->{'value1'}));
goto end_branch_27;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t27 = null;
  end_branch_27:;
  $__res = $__t27;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))((($dictMonadGen_1)->{'sized'})((($GLOBALS['Control_Semigroupoid_composeImpl'])((($dictMonadRec_0)->{'tailRecM'})(function($v_11) use ($Bind1_10_25, $__local_var_5_3, $pure_9_24) {
  $__num = \func_num_args();
  $__t30 = null;;
  if ((($v_11)->{'value1'} <= 0)) {
$__t30 = ($pure_9_24)(new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Done(($v_11)->{'value0'}));
goto end_branch_30;;
};
  $__local_var_12_28 = ($v_11)->{'value0'};
  $__local_var_13_29 = ($v_11)->{'value1'};
  $__t30 = ((($Bind1_10_25)->{'bind'})($__local_var_5_3))(function($x_14) use ($__local_var_12_28, $__local_var_13_29, $pure_9_24) {
  $__num = \func_num_args();
  $__res = ($pure_9_24)(new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Loop(new \Data\Tuple\Data_Tuple_Tuple(new \Control\Monad\Gen\Control_Monad_Gen_Cons($x_14, $__local_var_12_28), ($__local_var_13_29 - 1))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  end_branch_30:;
  $__res = $__t30;
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
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_Gen_genAlphaString'] = __NAMESPACE__ . '\\majData_majString_majGen_genmajAlphamajString';

// Data_String_Gen_genAlphaLowercaseString
function majData_majString_majGen_genmajAlphamajLowercasemajString($dictMonadRec_0, $dictMonadGen_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majGen_genmajAlphamajLowercasemajString';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $Monad0_2_0 = (($dictMonadGen_1)->{'Monad0'})(null);
  $Bind1_3_1 = (($Monad0_2_0)->{'Bind1'})(null);
  $Functor0_4_2 = (((((($Monad0_2_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $__local_var_5_3 = ((((((((((($dictMonadGen_1)->{'Monad0'})(null))->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})(function($x_5) {
  $__num = \func_num_args();
  $v_6_3 = \Data\Enum\majData_majEnum_charmajTomajEnum($x_5);
  $__t4 = null;;
  if ($v_6_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t4 = ($v_6_3)->{'value0'};
goto end_branch_4;;
};
  if ($v_6_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t5 = null;;
if (($x_5 < \Data\Enum\majData_majEnum_tomajCharmajCode($GLOBALS['Data_Bounded_bottomChar']))) {
$__t5 = $GLOBALS['Data_Bounded_bottomChar'];
goto end_branch_5;;
};
$__t5 = $GLOBALS['Data_Bounded_topChar'];
end_branch_5:;
$__t4 = $__t5;
goto end_branch_4;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t4 = null;
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($dictMonadGen_1)->{'chooseInt'})(97))(122));
  $__res = (($dictMonadGen_1)->{'sized'})(function($size_6) use ($Bind1_3_1, $Functor0_4_2, $__local_var_5_3, $dictMonadGen_1, $dictMonadRec_0) {
  $__num = \func_num_args();
  $v_7_7 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), 1, $size_6);
  $__t8 = null;;
  if ($v_7_7 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t8 = $size_6;
goto end_branch_8;;
};
  if ($v_7_7 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t8 = 1;
goto end_branch_8;;
};
  if ($v_7_7 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t8 = 1;
goto end_branch_8;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t8 = null;
  end_branch_8:;
  $__res = ((($Bind1_3_1)->{'bind'})(((($dictMonadGen_1)->{'chooseInt'})(1))($__t8)))(function($newSize_7) use ($Functor0_4_2, $__local_var_5_3, $dictMonadGen_1, $dictMonadRec_0) {
  $__num = \func_num_args();
  $Monad0_8_9 = (($dictMonadGen_1)->{'Monad0'})(null);
  $pure_9_10 = ((($Monad0_8_9)->{'Applicative0'})(null))->{'pure'};
  $Bind1_10_11 = (($Monad0_8_9)->{'Bind1'})(null);
  $__res = ((($dictMonadGen_1)->{'resize'})(function($v_8) use ($newSize_7) {
  $__num = \func_num_args();
  $__res = $newSize_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($Functor0_4_2)->{'map'})($GLOBALS['Data_String_CodeUnits_fromCharArray']))(((((((((($Monad0_8_9)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})(((((($GLOBALS['Data_Unfoldable_unfoldrArrayImpl'])($GLOBALS['Data_Maybe_isNothing']))(function($v_11) {
  $__num = \func_num_args();
  $__t12 = null;;
  if ($v_11 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t12 = ($v_11)->{'value0'};
goto end_branch_12;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t12 = null;
  end_branch_12:;
  $__res = $__t12;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Tuple_fst']))($GLOBALS['Data_Tuple_snd']))(function($v_11) {
  $__num = \func_num_args();
  $__t13 = null;;
  if ($v_11 instanceof \Control\Monad\Gen\Control_Monad_Gen_Nil) {
$__t13 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_13;;
};
  if ($v_11 instanceof \Control\Monad\Gen\Control_Monad_Gen_Cons) {
$__t13 = new \Data\Maybe\Data_Maybe_Just(new \Data\Tuple\Data_Tuple_Tuple(($v_11)->{'value0'}, ($v_11)->{'value1'}));
goto end_branch_13;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t13 = null;
  end_branch_13:;
  $__res = $__t13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))((($dictMonadGen_1)->{'sized'})((($GLOBALS['Control_Semigroupoid_composeImpl'])((($dictMonadRec_0)->{'tailRecM'})(function($v_11) use ($Bind1_10_11, $__local_var_5_3, $pure_9_10) {
  $__num = \func_num_args();
  $__t16 = null;;
  if ((($v_11)->{'value1'} <= 0)) {
$__t16 = ($pure_9_10)(new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Done(($v_11)->{'value0'}));
goto end_branch_16;;
};
  $__local_var_12_14 = ($v_11)->{'value0'};
  $__local_var_13_15 = ($v_11)->{'value1'};
  $__t16 = ((($Bind1_10_11)->{'bind'})($__local_var_5_3))(function($x_14) use ($__local_var_12_14, $__local_var_13_15, $pure_9_10) {
  $__num = \func_num_args();
  $__res = ($pure_9_10)(new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Loop(new \Data\Tuple\Data_Tuple_Tuple(new \Control\Monad\Gen\Control_Monad_Gen_Cons($x_14, $__local_var_12_14), ($__local_var_13_15 - 1))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  end_branch_16:;
  $__res = $__t16;
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
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_Gen_genAlphaLowercaseString'] = __NAMESPACE__ . '\\majData_majString_majGen_genmajAlphamajLowercasemajString';

