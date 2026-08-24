<?php

namespace AppFFI;

// ALL IMPORTS: Bench, Control.Bind, Data.Function, Data.Functor, Effect, Prelude, Prim, Test.AckermannFFI, Test.AckermannFFICheatcode, Test.ArrayOpsFFI, Test.ArrayOpsFFICheatcode, Test.AstTreeFFI, Test.AstTreeFFICheatcode, Test.ChurchFFI, Test.ChurchFFICheatcode, Test.FibFFI, Test.FibFFICheatcode, Test.LazyEvaluationFFI, Test.LazyEvaluationFFICheatcode, Test.ListOpsFFI, Test.ListOpsFFICheatcode, Test.PolymorphismFFI, Test.PolymorphismFFICheatcode, Test.PrimesFFI, Test.PrimesFFICheatcode, Test.RBTreeFFI, Test.RBTreeFFICheatcode, Test.RecordsFFI, Test.RecordsFFICheatcode, Test.RowToListFFI, Test.RowToListFFICheatcode, Test.StateMonadFFI, Test.StateMonadFFICheatcode, Test.TCOFFI, Test.TCOFFICheatcode
// TO REQUIRE: Bench, Control.Bind, Data.Function, Data.Functor, Effect, Prelude, Test.AckermannFFI, Test.AckermannFFICheatcode, Test.ArrayOpsFFI, Test.ArrayOpsFFICheatcode, Test.AstTreeFFI, Test.AstTreeFFICheatcode, Test.ChurchFFI, Test.ChurchFFICheatcode, Test.FibFFI, Test.FibFFICheatcode, Test.LazyEvaluationFFI, Test.LazyEvaluationFFICheatcode, Test.ListOpsFFI, Test.ListOpsFFICheatcode, Test.PolymorphismFFI, Test.PolymorphismFFICheatcode, Test.PrimesFFI, Test.PrimesFFICheatcode, Test.RBTreeFFI, Test.RBTreeFFICheatcode, Test.RecordsFFI, Test.RecordsFFICheatcode, Test.RowToListFFI, Test.RowToListFFICheatcode, Test.StateMonadFFI, Test.StateMonadFFICheatcode, Test.TCOFFI, Test.TCOFFICheatcode
require_once __DIR__ . '/../Bench/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Effect/index.php';
require_once __DIR__ . '/../Prelude/index.php';
require_once __DIR__ . '/../Test.AckermannFFI/index.php';
require_once __DIR__ . '/../Test.AckermannFFICheatcode/index.php';
require_once __DIR__ . '/../Test.ArrayOpsFFI/index.php';
require_once __DIR__ . '/../Test.ArrayOpsFFICheatcode/index.php';
require_once __DIR__ . '/../Test.AstTreeFFI/index.php';
require_once __DIR__ . '/../Test.AstTreeFFICheatcode/index.php';
require_once __DIR__ . '/../Test.ChurchFFI/index.php';
require_once __DIR__ . '/../Test.ChurchFFICheatcode/index.php';
require_once __DIR__ . '/../Test.FibFFI/index.php';
require_once __DIR__ . '/../Test.FibFFICheatcode/index.php';
require_once __DIR__ . '/../Test.LazyEvaluationFFI/index.php';
require_once __DIR__ . '/../Test.LazyEvaluationFFICheatcode/index.php';
require_once __DIR__ . '/../Test.ListOpsFFI/index.php';
require_once __DIR__ . '/../Test.ListOpsFFICheatcode/index.php';
require_once __DIR__ . '/../Test.PolymorphismFFI/index.php';
require_once __DIR__ . '/../Test.PolymorphismFFICheatcode/index.php';
require_once __DIR__ . '/../Test.PrimesFFI/index.php';
require_once __DIR__ . '/../Test.PrimesFFICheatcode/index.php';
require_once __DIR__ . '/../Test.RBTreeFFI/index.php';
require_once __DIR__ . '/../Test.RBTreeFFICheatcode/index.php';
require_once __DIR__ . '/../Test.RecordsFFI/index.php';
require_once __DIR__ . '/../Test.RecordsFFICheatcode/index.php';
require_once __DIR__ . '/../Test.RowToListFFI/index.php';
require_once __DIR__ . '/../Test.RowToListFFICheatcode/index.php';
require_once __DIR__ . '/../Test.StateMonadFFI/index.php';
require_once __DIR__ . '/../Test.StateMonadFFICheatcode/index.php';
require_once __DIR__ . '/../Test.TCOFFI/index.php';
require_once __DIR__ . '/../Test.TCOFFICheatcode/index.php';

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




// AppFFI_main
$GLOBALS['AppFFI_main'] = (function() use (&$__fn) {
$__local_var_0_0 = \Bench\majBench_runmajBench($GLOBALS['Test_AstTreeFFI_describe'], $GLOBALS['Test_AstTreeFFI_act']);
return function() use ($__local_var_0_0, &$__fn) {
$a_prime__1_1 = phpurs_execute_effect($__local_var_0_0);
$_dollar___unused_1_1 = phpurs_execute_effect(phpurs_execute_effect($GLOBALS['Data_Unit_unit']));
$a_prime__2_3 = phpurs_execute_effect(\Bench\majBench_runmajBench($GLOBALS['Test_FibFFI_describe'], $GLOBALS['Test_FibFFI_act']));
$_dollar___unused_2_3 = phpurs_execute_effect(phpurs_execute_effect($GLOBALS['Data_Unit_unit']));
$a_prime__3_5 = phpurs_execute_effect(\Bench\majBench_runmajBench($GLOBALS['Test_ListOpsFFI_describe'], $GLOBALS['Test_ListOpsFFI_act']));
$_dollar___unused_3_5 = phpurs_execute_effect(phpurs_execute_effect($GLOBALS['Data_Unit_unit']));
$a_prime__4_7 = phpurs_execute_effect(\Bench\majBench_runmajBench($GLOBALS['Test_TCOFFI_describe'], $GLOBALS['Test_TCOFFI_act']));
$_dollar___unused_4_7 = phpurs_execute_effect(phpurs_execute_effect($GLOBALS['Data_Unit_unit']));
$a_prime__5_9 = phpurs_execute_effect(\Bench\majBench_runmajBench($GLOBALS['Test_RecordsFFI_describe'], $GLOBALS['Test_RecordsFFI_act']));
$_dollar___unused_5_9 = phpurs_execute_effect(phpurs_execute_effect($GLOBALS['Data_Unit_unit']));
$a_prime__6_11 = phpurs_execute_effect(\Bench\majBench_runmajBench($GLOBALS['Test_AckermannFFI_describe'], $GLOBALS['Test_AckermannFFI_act']));
$_dollar___unused_6_11 = phpurs_execute_effect(phpurs_execute_effect($GLOBALS['Data_Unit_unit']));
$a_prime__7_13 = phpurs_execute_effect(\Bench\majBench_runmajBench($GLOBALS['Test_ChurchFFI_describe'], $GLOBALS['Test_ChurchFFI_act']));
$_dollar___unused_7_13 = phpurs_execute_effect(phpurs_execute_effect($GLOBALS['Data_Unit_unit']));
$a_prime__8_15 = phpurs_execute_effect(\Bench\majBench_runmajBench($GLOBALS['Test_PrimesFFI_describe'], $GLOBALS['Test_PrimesFFI_act']));
$_dollar___unused_8_15 = phpurs_execute_effect(phpurs_execute_effect($GLOBALS['Data_Unit_unit']));
$a_prime__9_17 = phpurs_execute_effect(\Bench\majBench_runmajBench($GLOBALS['Test_RBTreeFFI_describe'], $GLOBALS['Test_RBTreeFFI_act']));
$_dollar___unused_9_17 = phpurs_execute_effect(phpurs_execute_effect($GLOBALS['Data_Unit_unit']));
$a_prime__10_19 = phpurs_execute_effect(\Bench\majBench_runmajBench($GLOBALS['Test_PolymorphismFFI_describe'], $GLOBALS['Test_PolymorphismFFI_act']));
$_dollar___unused_10_19 = phpurs_execute_effect(phpurs_execute_effect($GLOBALS['Data_Unit_unit']));
$a_prime__11_21 = phpurs_execute_effect(\Bench\majBench_runmajBench($GLOBALS['Test_StateMonadFFI_describe'], $GLOBALS['Test_StateMonadFFI_act']));
$_dollar___unused_11_21 = phpurs_execute_effect(phpurs_execute_effect($GLOBALS['Data_Unit_unit']));
$a_prime__12_23 = phpurs_execute_effect(\Bench\majBench_runmajBench($GLOBALS['Test_LazyEvaluationFFI_describe'], $GLOBALS['Test_LazyEvaluationFFI_act']));
$_dollar___unused_12_23 = phpurs_execute_effect(phpurs_execute_effect($GLOBALS['Data_Unit_unit']));
$a_prime__13_25 = phpurs_execute_effect(\Bench\majBench_runmajBench($GLOBALS['Test_ArrayOpsFFI_describe'], $GLOBALS['Test_ArrayOpsFFI_act']));
$_dollar___unused_13_25 = phpurs_execute_effect(phpurs_execute_effect($GLOBALS['Data_Unit_unit']));
$a_prime__14_27 = phpurs_execute_effect(\Bench\majBench_runmajBench($GLOBALS['Test_RowToListFFI_describe'], $GLOBALS['Test_RowToListFFI_act']));
$_dollar___unused_14_27 = phpurs_execute_effect(phpurs_execute_effect($GLOBALS['Data_Unit_unit']));
$a_prime__15_29 = phpurs_execute_effect(\Bench\majBench_runmajBench($GLOBALS['Test_AstTreeFFICheatcode_describe'], $GLOBALS['Test_AstTreeFFICheatcode_act']));
$_dollar___unused_15_29 = phpurs_execute_effect(phpurs_execute_effect($GLOBALS['Data_Unit_unit']));
$a_prime__16_31 = phpurs_execute_effect(\Bench\majBench_runmajBench($GLOBALS['Test_FibFFICheatcode_describe'], $GLOBALS['Test_FibFFICheatcode_act']));
$_dollar___unused_16_31 = phpurs_execute_effect(phpurs_execute_effect($GLOBALS['Data_Unit_unit']));
$a_prime__17_33 = phpurs_execute_effect(\Bench\majBench_runmajBench($GLOBALS['Test_ListOpsFFICheatcode_describe'], $GLOBALS['Test_ListOpsFFICheatcode_act']));
$_dollar___unused_17_33 = phpurs_execute_effect(phpurs_execute_effect($GLOBALS['Data_Unit_unit']));
$a_prime__18_35 = phpurs_execute_effect(\Bench\majBench_runmajBench($GLOBALS['Test_TCOFFICheatcode_describe'], $GLOBALS['Test_TCOFFICheatcode_act']));
$_dollar___unused_18_35 = phpurs_execute_effect(phpurs_execute_effect($GLOBALS['Data_Unit_unit']));
$a_prime__19_37 = phpurs_execute_effect(\Bench\majBench_runmajBench($GLOBALS['Test_RecordsFFICheatcode_describe'], $GLOBALS['Test_RecordsFFICheatcode_act']));
$_dollar___unused_19_37 = phpurs_execute_effect(phpurs_execute_effect($GLOBALS['Data_Unit_unit']));
$a_prime__20_39 = phpurs_execute_effect(\Bench\majBench_runmajBench($GLOBALS['Test_AckermannFFICheatcode_describe'], $GLOBALS['Test_AckermannFFICheatcode_act']));
$_dollar___unused_20_39 = phpurs_execute_effect(phpurs_execute_effect($GLOBALS['Data_Unit_unit']));
$a_prime__21_41 = phpurs_execute_effect(\Bench\majBench_runmajBench($GLOBALS['Test_ChurchFFICheatcode_describe'], $GLOBALS['Test_ChurchFFICheatcode_act']));
$_dollar___unused_21_41 = phpurs_execute_effect(phpurs_execute_effect($GLOBALS['Data_Unit_unit']));
$a_prime__22_43 = phpurs_execute_effect(\Bench\majBench_runmajBench($GLOBALS['Test_PrimesFFICheatcode_describe'], $GLOBALS['Test_PrimesFFICheatcode_act']));
$_dollar___unused_22_43 = phpurs_execute_effect(phpurs_execute_effect($GLOBALS['Data_Unit_unit']));
$a_prime__23_45 = phpurs_execute_effect(\Bench\majBench_runmajBench($GLOBALS['Test_RBTreeFFICheatcode_describe'], $GLOBALS['Test_RBTreeFFICheatcode_act']));
$_dollar___unused_23_45 = phpurs_execute_effect(phpurs_execute_effect($GLOBALS['Data_Unit_unit']));
$a_prime__24_47 = phpurs_execute_effect(\Bench\majBench_runmajBench($GLOBALS['Test_PolymorphismFFICheatcode_describe'], $GLOBALS['Test_PolymorphismFFICheatcode_act']));
$_dollar___unused_24_47 = phpurs_execute_effect(phpurs_execute_effect($GLOBALS['Data_Unit_unit']));
$a_prime__25_49 = phpurs_execute_effect(\Bench\majBench_runmajBench($GLOBALS['Test_StateMonadFFICheatcode_describe'], $GLOBALS['Test_StateMonadFFICheatcode_act']));
$_dollar___unused_25_49 = phpurs_execute_effect(phpurs_execute_effect($GLOBALS['Data_Unit_unit']));
$a_prime__26_51 = phpurs_execute_effect(\Bench\majBench_runmajBench($GLOBALS['Test_LazyEvaluationFFICheatcode_describe'], $GLOBALS['Test_LazyEvaluationFFICheatcode_act']));
$_dollar___unused_26_51 = phpurs_execute_effect(phpurs_execute_effect($GLOBALS['Data_Unit_unit']));
$a_prime__27_53 = phpurs_execute_effect(\Bench\majBench_runmajBench($GLOBALS['Test_ArrayOpsFFICheatcode_describe'], $GLOBALS['Test_ArrayOpsFFICheatcode_act']));
$_dollar___unused_27_53 = phpurs_execute_effect(phpurs_execute_effect($GLOBALS['Data_Unit_unit']));
$a_prime__28_55 = phpurs_execute_effect(\Bench\majBench_runmajBench($GLOBALS['Test_RowToListFFICheatcode_describe'], $GLOBALS['Test_RowToListFFICheatcode_act']));
return phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect($GLOBALS['Data_Unit_unit']))))))))))))))))))))))))))));
};
})();

