<?php

namespace Data\Functor\Coproduct\Nested;

// ALL IMPORTS: Data.Const, Data.Either, Data.Functor.Coproduct, Data.Newtype, Data.Void, Prelude, Prim
// TO REQUIRE: Data.Const, Data.Either, Data.Functor.Coproduct, Data.Newtype, Data.Void, Prelude
require_once __DIR__ . '/../Data.Const/index.php';
require_once __DIR__ . '/../Data.Either/index.php';
require_once __DIR__ . '/../Data.Functor.Coproduct/index.php';
require_once __DIR__ . '/../Data.Newtype/index.php';
require_once __DIR__ . '/../Data.Void/index.php';
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

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };




// Data_Functor_Coproduct_Nested_in9
function majData_majFunctor_majCoproduct_majNested_in9($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majCoproduct_majNested_in9';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = new \Data\Either\Data_Either_Right(new \Data\Either\Data_Either_Right(new \Data\Either\Data_Either_Right(new \Data\Either\Data_Either_Right(new \Data\Either\Data_Either_Right(new \Data\Either\Data_Either_Right(new \Data\Either\Data_Either_Right(new \Data\Either\Data_Either_Right(new \Data\Either\Data_Either_Left($v_0)))))))));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Functor_Coproduct_Nested_in9'] = __NAMESPACE__ . '\\majData_majFunctor_majCoproduct_majNested_in9';

// Data_Functor_Coproduct_Nested_in8
function majData_majFunctor_majCoproduct_majNested_in8($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majCoproduct_majNested_in8';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = new \Data\Either\Data_Either_Right(new \Data\Either\Data_Either_Right(new \Data\Either\Data_Either_Right(new \Data\Either\Data_Either_Right(new \Data\Either\Data_Either_Right(new \Data\Either\Data_Either_Right(new \Data\Either\Data_Either_Right(new \Data\Either\Data_Either_Left($v_0))))))));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Functor_Coproduct_Nested_in8'] = __NAMESPACE__ . '\\majData_majFunctor_majCoproduct_majNested_in8';

// Data_Functor_Coproduct_Nested_in7
function majData_majFunctor_majCoproduct_majNested_in7($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majCoproduct_majNested_in7';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = new \Data\Either\Data_Either_Right(new \Data\Either\Data_Either_Right(new \Data\Either\Data_Either_Right(new \Data\Either\Data_Either_Right(new \Data\Either\Data_Either_Right(new \Data\Either\Data_Either_Right(new \Data\Either\Data_Either_Left($v_0)))))));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Functor_Coproduct_Nested_in7'] = __NAMESPACE__ . '\\majData_majFunctor_majCoproduct_majNested_in7';

// Data_Functor_Coproduct_Nested_in6
function majData_majFunctor_majCoproduct_majNested_in6($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majCoproduct_majNested_in6';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = new \Data\Either\Data_Either_Right(new \Data\Either\Data_Either_Right(new \Data\Either\Data_Either_Right(new \Data\Either\Data_Either_Right(new \Data\Either\Data_Either_Right(new \Data\Either\Data_Either_Left($v_0))))));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Functor_Coproduct_Nested_in6'] = __NAMESPACE__ . '\\majData_majFunctor_majCoproduct_majNested_in6';

// Data_Functor_Coproduct_Nested_in5
function majData_majFunctor_majCoproduct_majNested_in5($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majCoproduct_majNested_in5';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = new \Data\Either\Data_Either_Right(new \Data\Either\Data_Either_Right(new \Data\Either\Data_Either_Right(new \Data\Either\Data_Either_Right(new \Data\Either\Data_Either_Left($v_0)))));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Functor_Coproduct_Nested_in5'] = __NAMESPACE__ . '\\majData_majFunctor_majCoproduct_majNested_in5';

// Data_Functor_Coproduct_Nested_in4
function majData_majFunctor_majCoproduct_majNested_in4($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majCoproduct_majNested_in4';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = new \Data\Either\Data_Either_Right(new \Data\Either\Data_Either_Right(new \Data\Either\Data_Either_Right(new \Data\Either\Data_Either_Left($v_0))));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Functor_Coproduct_Nested_in4'] = __NAMESPACE__ . '\\majData_majFunctor_majCoproduct_majNested_in4';

// Data_Functor_Coproduct_Nested_in3
function majData_majFunctor_majCoproduct_majNested_in3($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majCoproduct_majNested_in3';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = new \Data\Either\Data_Either_Right(new \Data\Either\Data_Either_Right(new \Data\Either\Data_Either_Left($v_0)));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Functor_Coproduct_Nested_in3'] = __NAMESPACE__ . '\\majData_majFunctor_majCoproduct_majNested_in3';

// Data_Functor_Coproduct_Nested_in2
function majData_majFunctor_majCoproduct_majNested_in2($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majCoproduct_majNested_in2';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = new \Data\Either\Data_Either_Right(new \Data\Either\Data_Either_Left($v_0));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Functor_Coproduct_Nested_in2'] = __NAMESPACE__ . '\\majData_majFunctor_majCoproduct_majNested_in2';

// Data_Functor_Coproduct_Nested_in10
function majData_majFunctor_majCoproduct_majNested_in10($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majCoproduct_majNested_in10';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = new \Data\Either\Data_Either_Right(new \Data\Either\Data_Either_Right(new \Data\Either\Data_Either_Right(new \Data\Either\Data_Either_Right(new \Data\Either\Data_Either_Right(new \Data\Either\Data_Either_Right(new \Data\Either\Data_Either_Right(new \Data\Either\Data_Either_Right(new \Data\Either\Data_Either_Right(new \Data\Either\Data_Either_Left($v_0))))))))));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Functor_Coproduct_Nested_in10'] = __NAMESPACE__ . '\\majData_majFunctor_majCoproduct_majNested_in10';

// Data_Functor_Coproduct_Nested_in1_closure
$GLOBALS['Data_Functor_Coproduct_Nested_in1_closure'] = $GLOBALS['Data_Functor_Coproduct_left'];

// Data_Functor_Coproduct_Nested_in1
function majData_majFunctor_majCoproduct_majNested_in1($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majCoproduct_majNested_in1';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Functor_Coproduct_Nested_in1_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Functor_Coproduct_Nested_in1'] = __NAMESPACE__ . '\\majData_majFunctor_majCoproduct_majNested_in1';

// Data_Functor_Coproduct_Nested_coproduct9
function majData_majFunctor_majCoproduct_majNested_coproduct9($a_0, $b_1 = null, $c_2 = null, $d_3 = null, $e_4 = null, $f_5 = null, $g_6 = null, $h_7 = null, $i_8 = null, $y_9 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majCoproduct_majNested_coproduct9';
  if ($__num < 10) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 10);
  }
  $__t0 = null;;
  if ($y_9 instanceof \Data\Either\Data_Either_Left) {
$__t0 = ($a_0)(($y_9)->{'value0'});
goto end_branch_0;;
};
  if ($y_9 instanceof \Data\Either\Data_Either_Right) {
$__t1 = null;;
if (($y_9)->{'value0'} instanceof \Data\Either\Data_Either_Left) {
$__t1 = ($b_1)((($y_9)->{'value0'})->{'value0'});
goto end_branch_1;;
};
if (($y_9)->{'value0'} instanceof \Data\Either\Data_Either_Right) {
$__t2 = null;;
if ((($y_9)->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Left) {
$__t2 = ($c_2)(((($y_9)->{'value0'})->{'value0'})->{'value0'});
goto end_branch_2;;
};
if ((($y_9)->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right) {
$__t3 = null;;
if (((($y_9)->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Left) {
$__t3 = ($d_3)((((($y_9)->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_3;;
};
if (((($y_9)->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right) {
$__t4 = null;;
if ((((($y_9)->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Left) {
$__t4 = ($e_4)(((((($y_9)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_4;;
};
if ((((($y_9)->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right) {
$__t5 = null;;
if (((((($y_9)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Left) {
$__t5 = ($f_5)((((((($y_9)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_5;;
};
if (((((($y_9)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right) {
$__t6 = null;;
if ((((((($y_9)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Left) {
$__t6 = ($g_6)(((((((($y_9)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_6;;
};
if ((((((($y_9)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right) {
$__t7 = null;;
if (((((((($y_9)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Left) {
$__t7 = ($h_7)((((((((($y_9)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_7;;
};
if (((((((($y_9)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right) {
$__t8 = null;;
if ((((((((($y_9)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Left) {
$__t8 = ($i_8)(((((((((($y_9)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_8;;
};
if ((((((((($y_9)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right) {
$spin_10_9 = null;
$spin_10_9 = function($v_11) use (&$spin_10_9) {
  $__num = \func_num_args();
  $__tco_var_spin_10_9_9_v_11 = $v_11;
  tco_loop_spin_10_9_9:;
  $v_11 = $__tco_var_spin_10_9_9_v_11;
  $__tco_9 = $v_11;
  $__tco_var_spin_10_9_9_v_11 = $__tco_9;
  goto tco_loop_spin_10_9_9;;
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
$__t8 = ($spin_10_9)(((((((((($y_9)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_8;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t8 = null;
end_branch_8:;
$__t7 = $__t8;
goto end_branch_7;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t7 = null;
end_branch_7:;
$__t6 = $__t7;
goto end_branch_6;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t6 = null;
end_branch_6:;
$__t5 = $__t6;
goto end_branch_5;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t5 = null;
end_branch_5:;
$__t4 = $__t5;
goto end_branch_4;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t4 = null;
end_branch_4:;
$__t3 = $__t4;
goto end_branch_3;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
end_branch_3:;
$__t2 = $__t3;
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
$__t0 = $__t1;
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 10 < $__num ? $__res(...\array_slice(\func_get_args(), 10)) : $__res;
}
$GLOBALS['Data_Functor_Coproduct_Nested_coproduct9'] = __NAMESPACE__ . '\\majData_majFunctor_majCoproduct_majNested_coproduct9';

// Data_Functor_Coproduct_Nested_coproduct8
function majData_majFunctor_majCoproduct_majNested_coproduct8($a_0, $b_1 = null, $c_2 = null, $d_3 = null, $e_4 = null, $f_5 = null, $g_6 = null, $h_7 = null, $y_8 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majCoproduct_majNested_coproduct8';
  if ($__num < 9) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 9);
  }
  $__t0 = null;;
  if ($y_8 instanceof \Data\Either\Data_Either_Left) {
$__t0 = ($a_0)(($y_8)->{'value0'});
goto end_branch_0;;
};
  if ($y_8 instanceof \Data\Either\Data_Either_Right) {
$__t1 = null;;
if (($y_8)->{'value0'} instanceof \Data\Either\Data_Either_Left) {
$__t1 = ($b_1)((($y_8)->{'value0'})->{'value0'});
goto end_branch_1;;
};
if (($y_8)->{'value0'} instanceof \Data\Either\Data_Either_Right) {
$__t2 = null;;
if ((($y_8)->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Left) {
$__t2 = ($c_2)(((($y_8)->{'value0'})->{'value0'})->{'value0'});
goto end_branch_2;;
};
if ((($y_8)->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right) {
$__t3 = null;;
if (((($y_8)->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Left) {
$__t3 = ($d_3)((((($y_8)->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_3;;
};
if (((($y_8)->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right) {
$__t4 = null;;
if ((((($y_8)->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Left) {
$__t4 = ($e_4)(((((($y_8)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_4;;
};
if ((((($y_8)->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right) {
$__t5 = null;;
if (((((($y_8)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Left) {
$__t5 = ($f_5)((((((($y_8)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_5;;
};
if (((((($y_8)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right) {
$__t6 = null;;
if ((((((($y_8)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Left) {
$__t6 = ($g_6)(((((((($y_8)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_6;;
};
if ((((((($y_8)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right) {
$__t7 = null;;
if (((((((($y_8)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Left) {
$__t7 = ($h_7)((((((((($y_8)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_7;;
};
if (((((((($y_8)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right) {
$spin_9_8 = null;
$spin_9_8 = function($v_10) use (&$spin_9_8) {
  $__num = \func_num_args();
  $__tco_var_spin_9_8_8_v_10 = $v_10;
  tco_loop_spin_9_8_8:;
  $v_10 = $__tco_var_spin_9_8_8_v_10;
  $__tco_8 = $v_10;
  $__tco_var_spin_9_8_8_v_10 = $__tco_8;
  goto tco_loop_spin_9_8_8;;
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
$__t7 = ($spin_9_8)((((((((($y_8)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_7;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t7 = null;
end_branch_7:;
$__t6 = $__t7;
goto end_branch_6;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t6 = null;
end_branch_6:;
$__t5 = $__t6;
goto end_branch_5;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t5 = null;
end_branch_5:;
$__t4 = $__t5;
goto end_branch_4;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t4 = null;
end_branch_4:;
$__t3 = $__t4;
goto end_branch_3;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
end_branch_3:;
$__t2 = $__t3;
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
$__t0 = $__t1;
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 9 < $__num ? $__res(...\array_slice(\func_get_args(), 9)) : $__res;
}
$GLOBALS['Data_Functor_Coproduct_Nested_coproduct8'] = __NAMESPACE__ . '\\majData_majFunctor_majCoproduct_majNested_coproduct8';

// Data_Functor_Coproduct_Nested_coproduct7
function majData_majFunctor_majCoproduct_majNested_coproduct7($a_0, $b_1 = null, $c_2 = null, $d_3 = null, $e_4 = null, $f_5 = null, $g_6 = null, $y_7 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majCoproduct_majNested_coproduct7';
  if ($__num < 8) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 8);
  }
  $__t0 = null;;
  if ($y_7 instanceof \Data\Either\Data_Either_Left) {
$__t0 = ($a_0)(($y_7)->{'value0'});
goto end_branch_0;;
};
  if ($y_7 instanceof \Data\Either\Data_Either_Right) {
$__t1 = null;;
if (($y_7)->{'value0'} instanceof \Data\Either\Data_Either_Left) {
$__t1 = ($b_1)((($y_7)->{'value0'})->{'value0'});
goto end_branch_1;;
};
if (($y_7)->{'value0'} instanceof \Data\Either\Data_Either_Right) {
$__t2 = null;;
if ((($y_7)->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Left) {
$__t2 = ($c_2)(((($y_7)->{'value0'})->{'value0'})->{'value0'});
goto end_branch_2;;
};
if ((($y_7)->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right) {
$__t3 = null;;
if (((($y_7)->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Left) {
$__t3 = ($d_3)((((($y_7)->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_3;;
};
if (((($y_7)->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right) {
$__t4 = null;;
if ((((($y_7)->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Left) {
$__t4 = ($e_4)(((((($y_7)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_4;;
};
if ((((($y_7)->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right) {
$__t5 = null;;
if (((((($y_7)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Left) {
$__t5 = ($f_5)((((((($y_7)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_5;;
};
if (((((($y_7)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right) {
$__t6 = null;;
if ((((((($y_7)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Left) {
$__t6 = ($g_6)(((((((($y_7)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_6;;
};
if ((((((($y_7)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right) {
$spin_8_7 = null;
$spin_8_7 = function($v_9) use (&$spin_8_7) {
  $__num = \func_num_args();
  $__tco_var_spin_8_7_7_v_9 = $v_9;
  tco_loop_spin_8_7_7:;
  $v_9 = $__tco_var_spin_8_7_7_v_9;
  $__tco_7 = $v_9;
  $__tco_var_spin_8_7_7_v_9 = $__tco_7;
  goto tco_loop_spin_8_7_7;;
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
$__t6 = ($spin_8_7)(((((((($y_7)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_6;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t6 = null;
end_branch_6:;
$__t5 = $__t6;
goto end_branch_5;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t5 = null;
end_branch_5:;
$__t4 = $__t5;
goto end_branch_4;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t4 = null;
end_branch_4:;
$__t3 = $__t4;
goto end_branch_3;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
end_branch_3:;
$__t2 = $__t3;
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
$__t0 = $__t1;
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 8 < $__num ? $__res(...\array_slice(\func_get_args(), 8)) : $__res;
}
$GLOBALS['Data_Functor_Coproduct_Nested_coproduct7'] = __NAMESPACE__ . '\\majData_majFunctor_majCoproduct_majNested_coproduct7';

// Data_Functor_Coproduct_Nested_coproduct6
function majData_majFunctor_majCoproduct_majNested_coproduct6($a_0, $b_1 = null, $c_2 = null, $d_3 = null, $e_4 = null, $f_5 = null, $y_6 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majCoproduct_majNested_coproduct6';
  if ($__num < 7) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 7);
  }
  $__t0 = null;;
  if ($y_6 instanceof \Data\Either\Data_Either_Left) {
$__t0 = ($a_0)(($y_6)->{'value0'});
goto end_branch_0;;
};
  if ($y_6 instanceof \Data\Either\Data_Either_Right) {
$__t1 = null;;
if (($y_6)->{'value0'} instanceof \Data\Either\Data_Either_Left) {
$__t1 = ($b_1)((($y_6)->{'value0'})->{'value0'});
goto end_branch_1;;
};
if (($y_6)->{'value0'} instanceof \Data\Either\Data_Either_Right) {
$__t2 = null;;
if ((($y_6)->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Left) {
$__t2 = ($c_2)(((($y_6)->{'value0'})->{'value0'})->{'value0'});
goto end_branch_2;;
};
if ((($y_6)->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right) {
$__t3 = null;;
if (((($y_6)->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Left) {
$__t3 = ($d_3)((((($y_6)->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_3;;
};
if (((($y_6)->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right) {
$__t4 = null;;
if ((((($y_6)->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Left) {
$__t4 = ($e_4)(((((($y_6)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_4;;
};
if ((((($y_6)->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right) {
$__t5 = null;;
if (((((($y_6)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Left) {
$__t5 = ($f_5)((((((($y_6)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_5;;
};
if (((((($y_6)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right) {
$spin_7_6 = null;
$spin_7_6 = function($v_8) use (&$spin_7_6) {
  $__num = \func_num_args();
  $__tco_var_spin_7_6_6_v_8 = $v_8;
  tco_loop_spin_7_6_6:;
  $v_8 = $__tco_var_spin_7_6_6_v_8;
  $__tco_6 = $v_8;
  $__tco_var_spin_7_6_6_v_8 = $__tco_6;
  goto tco_loop_spin_7_6_6;;
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
$__t5 = ($spin_7_6)((((((($y_6)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_5;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t5 = null;
end_branch_5:;
$__t4 = $__t5;
goto end_branch_4;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t4 = null;
end_branch_4:;
$__t3 = $__t4;
goto end_branch_3;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
end_branch_3:;
$__t2 = $__t3;
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
$__t0 = $__t1;
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 7 < $__num ? $__res(...\array_slice(\func_get_args(), 7)) : $__res;
}
$GLOBALS['Data_Functor_Coproduct_Nested_coproduct6'] = __NAMESPACE__ . '\\majData_majFunctor_majCoproduct_majNested_coproduct6';

// Data_Functor_Coproduct_Nested_coproduct5
function majData_majFunctor_majCoproduct_majNested_coproduct5($a_0, $b_1 = null, $c_2 = null, $d_3 = null, $e_4 = null, $y_5 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majCoproduct_majNested_coproduct5';
  if ($__num < 6) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 6);
  }
  $__t0 = null;;
  if ($y_5 instanceof \Data\Either\Data_Either_Left) {
$__t0 = ($a_0)(($y_5)->{'value0'});
goto end_branch_0;;
};
  if ($y_5 instanceof \Data\Either\Data_Either_Right) {
$__t1 = null;;
if (($y_5)->{'value0'} instanceof \Data\Either\Data_Either_Left) {
$__t1 = ($b_1)((($y_5)->{'value0'})->{'value0'});
goto end_branch_1;;
};
if (($y_5)->{'value0'} instanceof \Data\Either\Data_Either_Right) {
$__t2 = null;;
if ((($y_5)->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Left) {
$__t2 = ($c_2)(((($y_5)->{'value0'})->{'value0'})->{'value0'});
goto end_branch_2;;
};
if ((($y_5)->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right) {
$__t3 = null;;
if (((($y_5)->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Left) {
$__t3 = ($d_3)((((($y_5)->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_3;;
};
if (((($y_5)->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right) {
$__t4 = null;;
if ((((($y_5)->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Left) {
$__t4 = ($e_4)(((((($y_5)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_4;;
};
if ((((($y_5)->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right) {
$spin_6_5 = null;
$spin_6_5 = function($v_7) use (&$spin_6_5) {
  $__num = \func_num_args();
  $__tco_var_spin_6_5_5_v_7 = $v_7;
  tco_loop_spin_6_5_5:;
  $v_7 = $__tco_var_spin_6_5_5_v_7;
  $__tco_5 = $v_7;
  $__tco_var_spin_6_5_5_v_7 = $__tco_5;
  goto tco_loop_spin_6_5_5;;
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
$__t4 = ($spin_6_5)(((((($y_5)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_4;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t4 = null;
end_branch_4:;
$__t3 = $__t4;
goto end_branch_3;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
end_branch_3:;
$__t2 = $__t3;
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
$__t0 = $__t1;
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 6 < $__num ? $__res(...\array_slice(\func_get_args(), 6)) : $__res;
}
$GLOBALS['Data_Functor_Coproduct_Nested_coproduct5'] = __NAMESPACE__ . '\\majData_majFunctor_majCoproduct_majNested_coproduct5';

// Data_Functor_Coproduct_Nested_coproduct4
function majData_majFunctor_majCoproduct_majNested_coproduct4($a_0, $b_1 = null, $c_2 = null, $d_3 = null, $y_4 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majCoproduct_majNested_coproduct4';
  if ($__num < 5) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 5);
  }
  $__t0 = null;;
  if ($y_4 instanceof \Data\Either\Data_Either_Left) {
$__t0 = ($a_0)(($y_4)->{'value0'});
goto end_branch_0;;
};
  if ($y_4 instanceof \Data\Either\Data_Either_Right) {
$__t1 = null;;
if (($y_4)->{'value0'} instanceof \Data\Either\Data_Either_Left) {
$__t1 = ($b_1)((($y_4)->{'value0'})->{'value0'});
goto end_branch_1;;
};
if (($y_4)->{'value0'} instanceof \Data\Either\Data_Either_Right) {
$__t2 = null;;
if ((($y_4)->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Left) {
$__t2 = ($c_2)(((($y_4)->{'value0'})->{'value0'})->{'value0'});
goto end_branch_2;;
};
if ((($y_4)->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right) {
$__t3 = null;;
if (((($y_4)->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Left) {
$__t3 = ($d_3)((((($y_4)->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_3;;
};
if (((($y_4)->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right) {
$spin_5_4 = null;
$spin_5_4 = function($v_6) use (&$spin_5_4) {
  $__num = \func_num_args();
  $__tco_var_spin_5_4_4_v_6 = $v_6;
  tco_loop_spin_5_4_4:;
  $v_6 = $__tco_var_spin_5_4_4_v_6;
  $__tco_4 = $v_6;
  $__tco_var_spin_5_4_4_v_6 = $__tco_4;
  goto tco_loop_spin_5_4_4;;
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
$__t3 = ($spin_5_4)((((($y_4)->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_3;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
end_branch_3:;
$__t2 = $__t3;
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
$__t0 = $__t1;
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 5 < $__num ? $__res(...\array_slice(\func_get_args(), 5)) : $__res;
}
$GLOBALS['Data_Functor_Coproduct_Nested_coproduct4'] = __NAMESPACE__ . '\\majData_majFunctor_majCoproduct_majNested_coproduct4';

// Data_Functor_Coproduct_Nested_coproduct3
function majData_majFunctor_majCoproduct_majNested_coproduct3($a_0, $b_1 = null, $c_2 = null, $y_3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majCoproduct_majNested_coproduct3';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__t0 = null;;
  if ($y_3 instanceof \Data\Either\Data_Either_Left) {
$__t0 = ($a_0)(($y_3)->{'value0'});
goto end_branch_0;;
};
  if ($y_3 instanceof \Data\Either\Data_Either_Right) {
$__t1 = null;;
if (($y_3)->{'value0'} instanceof \Data\Either\Data_Either_Left) {
$__t1 = ($b_1)((($y_3)->{'value0'})->{'value0'});
goto end_branch_1;;
};
if (($y_3)->{'value0'} instanceof \Data\Either\Data_Either_Right) {
$__t2 = null;;
if ((($y_3)->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Left) {
$__t2 = ($c_2)(((($y_3)->{'value0'})->{'value0'})->{'value0'});
goto end_branch_2;;
};
if ((($y_3)->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right) {
$spin_4_3 = null;
$spin_4_3 = function($v_5) use (&$spin_4_3) {
  $__num = \func_num_args();
  $__tco_var_spin_4_3_3_v_5 = $v_5;
  tco_loop_spin_4_3_3:;
  $v_5 = $__tco_var_spin_4_3_3_v_5;
  $__tco_3 = $v_5;
  $__tco_var_spin_4_3_3_v_5 = $__tco_3;
  goto tco_loop_spin_4_3_3;;
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
$__t2 = ($spin_4_3)(((($y_3)->{'value0'})->{'value0'})->{'value0'});
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
$__t0 = $__t1;
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 4 < $__num ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
}
$GLOBALS['Data_Functor_Coproduct_Nested_coproduct3'] = __NAMESPACE__ . '\\majData_majFunctor_majCoproduct_majNested_coproduct3';

// Data_Functor_Coproduct_Nested_coproduct2
function majData_majFunctor_majCoproduct_majNested_coproduct2($a_0, $b_1 = null, $y_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majCoproduct_majNested_coproduct2';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if ($y_2 instanceof \Data\Either\Data_Either_Left) {
$__t0 = ($a_0)(($y_2)->{'value0'});
goto end_branch_0;;
};
  if ($y_2 instanceof \Data\Either\Data_Either_Right) {
$__t1 = null;;
if (($y_2)->{'value0'} instanceof \Data\Either\Data_Either_Left) {
$__t1 = ($b_1)((($y_2)->{'value0'})->{'value0'});
goto end_branch_1;;
};
if (($y_2)->{'value0'} instanceof \Data\Either\Data_Either_Right) {
$spin_3_2 = null;
$spin_3_2 = function($v_4) use (&$spin_3_2) {
  $__num = \func_num_args();
  $__tco_var_spin_3_2_2_v_4 = $v_4;
  tco_loop_spin_3_2_2:;
  $v_4 = $__tco_var_spin_3_2_2_v_4;
  $__tco_2 = $v_4;
  $__tco_var_spin_3_2_2_v_4 = $__tco_2;
  goto tco_loop_spin_3_2_2;;
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
$__t1 = ($spin_3_2)((($y_2)->{'value0'})->{'value0'});
goto end_branch_1;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
end_branch_1:;
$__t0 = $__t1;
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Functor_Coproduct_Nested_coproduct2'] = __NAMESPACE__ . '\\majData_majFunctor_majCoproduct_majNested_coproduct2';

// Data_Functor_Coproduct_Nested_coproduct10
function majData_majFunctor_majCoproduct_majNested_coproduct10($a_0, $b_1 = null, $c_2 = null, $d_3 = null, $e_4 = null, $f_5 = null, $g_6 = null, $h_7 = null, $i_8 = null, $j_9 = null, $y_10 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majCoproduct_majNested_coproduct10';
  if ($__num < 11) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 11);
  }
  $__t0 = null;;
  if ($y_10 instanceof \Data\Either\Data_Either_Left) {
$__t0 = ($a_0)(($y_10)->{'value0'});
goto end_branch_0;;
};
  if ($y_10 instanceof \Data\Either\Data_Either_Right) {
$__t1 = null;;
if (($y_10)->{'value0'} instanceof \Data\Either\Data_Either_Left) {
$__t1 = ($b_1)((($y_10)->{'value0'})->{'value0'});
goto end_branch_1;;
};
if (($y_10)->{'value0'} instanceof \Data\Either\Data_Either_Right) {
$__t2 = null;;
if ((($y_10)->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Left) {
$__t2 = ($c_2)(((($y_10)->{'value0'})->{'value0'})->{'value0'});
goto end_branch_2;;
};
if ((($y_10)->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right) {
$__t3 = null;;
if (((($y_10)->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Left) {
$__t3 = ($d_3)((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_3;;
};
if (((($y_10)->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right) {
$__t4 = null;;
if ((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Left) {
$__t4 = ($e_4)(((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_4;;
};
if ((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right) {
$__t5 = null;;
if (((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Left) {
$__t5 = ($f_5)((((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_5;;
};
if (((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right) {
$__t6 = null;;
if ((((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Left) {
$__t6 = ($g_6)(((((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_6;;
};
if ((((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right) {
$__t7 = null;;
if (((((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Left) {
$__t7 = ($h_7)((((((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_7;;
};
if (((((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right) {
$__t8 = null;;
if ((((((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Left) {
$__t8 = ($i_8)(((((((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_8;;
};
if ((((((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right) {
$__t9 = null;;
if (((((((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Left) {
$__t9 = ($j_9)((((((((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_9;;
};
if (((((((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right) {
$spin_11_10 = null;
$spin_11_10 = function($v_12) use (&$spin_11_10) {
  $__num = \func_num_args();
  $__tco_var_spin_11_10_10_v_12 = $v_12;
  tco_loop_spin_11_10_10:;
  $v_12 = $__tco_var_spin_11_10_10_v_12;
  $__tco_10 = $v_12;
  $__tco_var_spin_11_10_10_v_12 = $__tco_10;
  goto tco_loop_spin_11_10_10;;
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
$__t9 = ($spin_11_10)((((((((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_9;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t9 = null;
end_branch_9:;
$__t8 = $__t9;
goto end_branch_8;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t8 = null;
end_branch_8:;
$__t7 = $__t8;
goto end_branch_7;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t7 = null;
end_branch_7:;
$__t6 = $__t7;
goto end_branch_6;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t6 = null;
end_branch_6:;
$__t5 = $__t6;
goto end_branch_5;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t5 = null;
end_branch_5:;
$__t4 = $__t5;
goto end_branch_4;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t4 = null;
end_branch_4:;
$__t3 = $__t4;
goto end_branch_3;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
end_branch_3:;
$__t2 = $__t3;
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
$__t0 = $__t1;
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 11 < $__num ? $__res(...\array_slice(\func_get_args(), 11)) : $__res;
}
$GLOBALS['Data_Functor_Coproduct_Nested_coproduct10'] = __NAMESPACE__ . '\\majData_majFunctor_majCoproduct_majNested_coproduct10';

// Data_Functor_Coproduct_Nested_coproduct1
function majData_majFunctor_majCoproduct_majNested_coproduct1($y_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majCoproduct_majNested_coproduct1';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__t0 = null;;
  if ($y_0 instanceof \Data\Either\Data_Either_Left) {
$__t0 = ($y_0)->{'value0'};
goto end_branch_0;;
};
  if ($y_0 instanceof \Data\Either\Data_Either_Right) {
$spin_1_1 = null;
$spin_1_1 = function($v_2) use (&$spin_1_1) {
  $__num = \func_num_args();
  $__tco_var_spin_1_1_1_v_2 = $v_2;
  tco_loop_spin_1_1_1:;
  $v_2 = $__tco_var_spin_1_1_1_v_2;
  $__tco_1 = $v_2;
  $__tco_var_spin_1_1_1_v_2 = $__tco_1;
  goto tco_loop_spin_1_1_1;;
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
$__t0 = ($spin_1_1)(($y_0)->{'value0'});
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Functor_Coproduct_Nested_coproduct1'] = __NAMESPACE__ . '\\majData_majFunctor_majCoproduct_majNested_coproduct1';

// Data_Functor_Coproduct_Nested_at9
function majData_majFunctor_majCoproduct_majNested_at9($b_0, $f_1 = null, $y_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majCoproduct_majNested_at9';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if (($y_2 instanceof \Data\Either\Data_Either_Right && (($y_2)->{'value0'} instanceof \Data\Either\Data_Either_Right && ((($y_2)->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right && (((($y_2)->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right && ((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right && (((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right && ((((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right && (((((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right && (((((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Left))))))))) {
$__t0 = ($f_1)(((((((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_0;;
};
  $__t0 = $b_0;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Functor_Coproduct_Nested_at9'] = __NAMESPACE__ . '\\majData_majFunctor_majCoproduct_majNested_at9';

// Data_Functor_Coproduct_Nested_at8
function majData_majFunctor_majCoproduct_majNested_at8($b_0, $f_1 = null, $y_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majCoproduct_majNested_at8';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if (($y_2 instanceof \Data\Either\Data_Either_Right && (($y_2)->{'value0'} instanceof \Data\Either\Data_Either_Right && ((($y_2)->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right && (((($y_2)->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right && ((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right && (((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right && ((((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right && ((((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Left)))))))) {
$__t0 = ($f_1)((((((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_0;;
};
  $__t0 = $b_0;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Functor_Coproduct_Nested_at8'] = __NAMESPACE__ . '\\majData_majFunctor_majCoproduct_majNested_at8';

// Data_Functor_Coproduct_Nested_at7
function majData_majFunctor_majCoproduct_majNested_at7($b_0, $f_1 = null, $y_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majCoproduct_majNested_at7';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if (($y_2 instanceof \Data\Either\Data_Either_Right && (($y_2)->{'value0'} instanceof \Data\Either\Data_Either_Right && ((($y_2)->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right && (((($y_2)->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right && ((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right && (((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right && (((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Left))))))) {
$__t0 = ($f_1)(((((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_0;;
};
  $__t0 = $b_0;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Functor_Coproduct_Nested_at7'] = __NAMESPACE__ . '\\majData_majFunctor_majCoproduct_majNested_at7';

// Data_Functor_Coproduct_Nested_at6
function majData_majFunctor_majCoproduct_majNested_at6($b_0, $f_1 = null, $y_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majCoproduct_majNested_at6';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if (($y_2 instanceof \Data\Either\Data_Either_Right && (($y_2)->{'value0'} instanceof \Data\Either\Data_Either_Right && ((($y_2)->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right && (((($y_2)->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right && ((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right && ((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Left)))))) {
$__t0 = ($f_1)((((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_0;;
};
  $__t0 = $b_0;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Functor_Coproduct_Nested_at6'] = __NAMESPACE__ . '\\majData_majFunctor_majCoproduct_majNested_at6';

// Data_Functor_Coproduct_Nested_at5
function majData_majFunctor_majCoproduct_majNested_at5($b_0, $f_1 = null, $y_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majCoproduct_majNested_at5';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if (($y_2 instanceof \Data\Either\Data_Either_Right && (($y_2)->{'value0'} instanceof \Data\Either\Data_Either_Right && ((($y_2)->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right && (((($y_2)->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right && (((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Left))))) {
$__t0 = ($f_1)(((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_0;;
};
  $__t0 = $b_0;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Functor_Coproduct_Nested_at5'] = __NAMESPACE__ . '\\majData_majFunctor_majCoproduct_majNested_at5';

// Data_Functor_Coproduct_Nested_at4
function majData_majFunctor_majCoproduct_majNested_at4($b_0, $f_1 = null, $y_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majCoproduct_majNested_at4';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if (($y_2 instanceof \Data\Either\Data_Either_Right && (($y_2)->{'value0'} instanceof \Data\Either\Data_Either_Right && ((($y_2)->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right && ((($y_2)->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Left)))) {
$__t0 = ($f_1)((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_0;;
};
  $__t0 = $b_0;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Functor_Coproduct_Nested_at4'] = __NAMESPACE__ . '\\majData_majFunctor_majCoproduct_majNested_at4';

// Data_Functor_Coproduct_Nested_at3
function majData_majFunctor_majCoproduct_majNested_at3($b_0, $f_1 = null, $y_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majCoproduct_majNested_at3';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if (($y_2 instanceof \Data\Either\Data_Either_Right && (($y_2)->{'value0'} instanceof \Data\Either\Data_Either_Right && (($y_2)->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Left))) {
$__t0 = ($f_1)(((($y_2)->{'value0'})->{'value0'})->{'value0'});
goto end_branch_0;;
};
  $__t0 = $b_0;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Functor_Coproduct_Nested_at3'] = __NAMESPACE__ . '\\majData_majFunctor_majCoproduct_majNested_at3';

// Data_Functor_Coproduct_Nested_at2
function majData_majFunctor_majCoproduct_majNested_at2($b_0, $f_1 = null, $y_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majCoproduct_majNested_at2';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if (($y_2 instanceof \Data\Either\Data_Either_Right && ($y_2)->{'value0'} instanceof \Data\Either\Data_Either_Left)) {
$__t0 = ($f_1)((($y_2)->{'value0'})->{'value0'});
goto end_branch_0;;
};
  $__t0 = $b_0;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Functor_Coproduct_Nested_at2'] = __NAMESPACE__ . '\\majData_majFunctor_majCoproduct_majNested_at2';

// Data_Functor_Coproduct_Nested_at10
function majData_majFunctor_majCoproduct_majNested_at10($b_0, $f_1 = null, $y_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majCoproduct_majNested_at10';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if (($y_2 instanceof \Data\Either\Data_Either_Right && (($y_2)->{'value0'} instanceof \Data\Either\Data_Either_Right && ((($y_2)->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right && (((($y_2)->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right && ((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right && (((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right && ((((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right && (((((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right && ((((((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Right && ((((((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'} instanceof \Data\Either\Data_Either_Left)))))))))) {
$__t0 = ($f_1)((((((((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_0;;
};
  $__t0 = $b_0;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Functor_Coproduct_Nested_at10'] = __NAMESPACE__ . '\\majData_majFunctor_majCoproduct_majNested_at10';

// Data_Functor_Coproduct_Nested_at1
function majData_majFunctor_majCoproduct_majNested_at1($b_0, $f_1 = null, $y_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majCoproduct_majNested_at1';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if ($y_2 instanceof \Data\Either\Data_Either_Left) {
$__t0 = ($f_1)(($y_2)->{'value0'});
goto end_branch_0;;
};
  $__t0 = $b_0;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Functor_Coproduct_Nested_at1'] = __NAMESPACE__ . '\\majData_majFunctor_majCoproduct_majNested_at1';

