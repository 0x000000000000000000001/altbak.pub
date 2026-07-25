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
\PhpursThunks::$thunks['Data_Functor_Coproduct_Nested_in9'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data1("Right", new Phpurs_Data1("Right", new Phpurs_Data1("Right", new Phpurs_Data1("Right", new Phpurs_Data1("Right", new Phpurs_Data1("Right", new Phpurs_Data1("Right", new Phpurs_Data1("Right", new Phpurs_Data1("Left", $v_0)))))))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Functor_Coproduct_Nested_in8'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data1("Right", new Phpurs_Data1("Right", new Phpurs_Data1("Right", new Phpurs_Data1("Right", new Phpurs_Data1("Right", new Phpurs_Data1("Right", new Phpurs_Data1("Right", new Phpurs_Data1("Left", $v_0))))))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Functor_Coproduct_Nested_in7'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data1("Right", new Phpurs_Data1("Right", new Phpurs_Data1("Right", new Phpurs_Data1("Right", new Phpurs_Data1("Right", new Phpurs_Data1("Right", new Phpurs_Data1("Left", $v_0)))))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Functor_Coproduct_Nested_in6'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data1("Right", new Phpurs_Data1("Right", new Phpurs_Data1("Right", new Phpurs_Data1("Right", new Phpurs_Data1("Right", new Phpurs_Data1("Left", $v_0))))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Functor_Coproduct_Nested_in5'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data1("Right", new Phpurs_Data1("Right", new Phpurs_Data1("Right", new Phpurs_Data1("Right", new Phpurs_Data1("Left", $v_0)))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Functor_Coproduct_Nested_in4'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data1("Right", new Phpurs_Data1("Right", new Phpurs_Data1("Right", new Phpurs_Data1("Left", $v_0))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Functor_Coproduct_Nested_in3'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data1("Right", new Phpurs_Data1("Right", new Phpurs_Data1("Left", $v_0)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Functor_Coproduct_Nested_in2'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data1("Right", new Phpurs_Data1("Left", $v_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Functor_Coproduct_Nested_in10'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data1("Right", new Phpurs_Data1("Right", new Phpurs_Data1("Right", new Phpurs_Data1("Right", new Phpurs_Data1("Right", new Phpurs_Data1("Right", new Phpurs_Data1("Right", new Phpurs_Data1("Right", new Phpurs_Data1("Right", new Phpurs_Data1("Left", $v_0))))))))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Functor_Coproduct_Nested_in1'] = function() { $v = ($GLOBALS['Data_Functor_Coproduct_left'] ?? \PhpursThunks::eval('Data_Functor_Coproduct_left')); return $v; };
\PhpursThunks::$thunks['Data_Functor_Coproduct_Nested_coproduct9'] = function() { $v = (function() {
  $__fn = function($a_0 = null, $b_1 = null, $c_2 = null, $d_3 = null, $e_4 = null, $f_5 = null, $g_6 = null, $h_7 = null, $i_8 = null, $y_9 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 10) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 10);
  }
  $__t0 = null;;
  if ((is_object($y_9) && (($y_9)->{'tag'} === "Left"))) {
$__t0 = ($a_0)(($y_9)->{'value0'});
goto end_branch_0;;
};
  if ((is_object($y_9) && (($y_9)->{'tag'} === "Right"))) {
$__t1 = null;;
if ((is_object(($y_9)->{'value0'}) && ((($y_9)->{'value0'})->{'tag'} === "Left"))) {
$__t1 = ($b_1)((($y_9)->{'value0'})->{'value0'});
goto end_branch_1;;
};
if ((is_object(($y_9)->{'value0'}) && ((($y_9)->{'value0'})->{'tag'} === "Right"))) {
$__t2 = null;;
if ((is_object((($y_9)->{'value0'})->{'value0'}) && (((($y_9)->{'value0'})->{'value0'})->{'tag'} === "Left"))) {
$__t2 = ($c_2)(((($y_9)->{'value0'})->{'value0'})->{'value0'});
goto end_branch_2;;
};
if ((is_object((($y_9)->{'value0'})->{'value0'}) && (((($y_9)->{'value0'})->{'value0'})->{'tag'} === "Right"))) {
$__t3 = null;;
if ((is_object(((($y_9)->{'value0'})->{'value0'})->{'value0'}) && ((((($y_9)->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Left"))) {
$__t3 = ($d_3)((((($y_9)->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_3;;
};
if ((is_object(((($y_9)->{'value0'})->{'value0'})->{'value0'}) && ((((($y_9)->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Right"))) {
$__t4 = null;;
if ((is_object((((($y_9)->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && (((((($y_9)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Left"))) {
$__t4 = ($e_4)(((((($y_9)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_4;;
};
if ((is_object((((($y_9)->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && (((((($y_9)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Right"))) {
$__t5 = null;;
if ((is_object(((((($y_9)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && ((((((($y_9)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Left"))) {
$__t5 = ($f_5)((((((($y_9)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_5;;
};
if ((is_object(((((($y_9)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && ((((((($y_9)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Right"))) {
$__t6 = null;;
if ((is_object((((((($y_9)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && (((((((($y_9)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Left"))) {
$__t6 = ($g_6)(((((((($y_9)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_6;;
};
if ((is_object((((((($y_9)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && (((((((($y_9)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Right"))) {
$__t7 = null;;
if ((is_object(((((((($y_9)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && ((((((((($y_9)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Left"))) {
$__t7 = ($h_7)((((((((($y_9)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_7;;
};
if ((is_object(((((((($y_9)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && ((((((((($y_9)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Right"))) {
$__t8 = null;;
if ((is_object((((((((($y_9)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && (((((((((($y_9)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Left"))) {
$__t8 = ($i_8)(((((((((($y_9)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_8;;
};
if ((is_object((((((((($y_9)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && (((((((((($y_9)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Right"))) {
$spin_10_9 = null;
$spin_10_9 = function($v_11 = null) use (&$spin_10_9) {
  $__num = \func_num_args();
  $__tco_var_spin_10_9_9_v_11 = $v_11;
  $__tco_done_spin_10_9_9 = false;
  $__tco_res_spin_10_9_9 = null;
  $__tco_loop_spin_10_9_9 = function($v_11 = null) use (&$__tco_done_spin_10_9_9, &$__tco_var_spin_10_9_9_v_11, &$spin_10_9) {
  $__num = \func_num_args();
  $__tco_done_spin_10_9_9 = true;
  $v_11 = $__tco_var_spin_10_9_9_v_11;
  $__tco_9 = $v_11;
  $__tco_var_spin_10_9_9_v_11 = $__tco_9;
  $__tco_done_spin_10_9_9 = false;
  $__res = null;
  goto __end;;
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  while (($__tco_done_spin_10_9_9 === false)) {
$__tco_res_spin_10_9_9 = ($__tco_loop_spin_10_9_9)($__tco_var_spin_10_9_9_v_11);
};
  $__res = $__tco_res_spin_10_9_9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
$__t8 = ($spin_10_9)((($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce')))(((((((((($y_9)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}));
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
  return $__num > 10 ? $__res(...\array_slice(\func_get_args(), 10)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Functor_Coproduct_Nested_coproduct8'] = function() { $v = (function() {
  $__fn = function($a_0 = null, $b_1 = null, $c_2 = null, $d_3 = null, $e_4 = null, $f_5 = null, $g_6 = null, $h_7 = null, $y_8 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 9) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 9);
  }
  $__t0 = null;;
  if ((is_object($y_8) && (($y_8)->{'tag'} === "Left"))) {
$__t0 = ($a_0)(($y_8)->{'value0'});
goto end_branch_0;;
};
  if ((is_object($y_8) && (($y_8)->{'tag'} === "Right"))) {
$__t1 = null;;
if ((is_object(($y_8)->{'value0'}) && ((($y_8)->{'value0'})->{'tag'} === "Left"))) {
$__t1 = ($b_1)((($y_8)->{'value0'})->{'value0'});
goto end_branch_1;;
};
if ((is_object(($y_8)->{'value0'}) && ((($y_8)->{'value0'})->{'tag'} === "Right"))) {
$__t2 = null;;
if ((is_object((($y_8)->{'value0'})->{'value0'}) && (((($y_8)->{'value0'})->{'value0'})->{'tag'} === "Left"))) {
$__t2 = ($c_2)(((($y_8)->{'value0'})->{'value0'})->{'value0'});
goto end_branch_2;;
};
if ((is_object((($y_8)->{'value0'})->{'value0'}) && (((($y_8)->{'value0'})->{'value0'})->{'tag'} === "Right"))) {
$__t3 = null;;
if ((is_object(((($y_8)->{'value0'})->{'value0'})->{'value0'}) && ((((($y_8)->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Left"))) {
$__t3 = ($d_3)((((($y_8)->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_3;;
};
if ((is_object(((($y_8)->{'value0'})->{'value0'})->{'value0'}) && ((((($y_8)->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Right"))) {
$__t4 = null;;
if ((is_object((((($y_8)->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && (((((($y_8)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Left"))) {
$__t4 = ($e_4)(((((($y_8)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_4;;
};
if ((is_object((((($y_8)->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && (((((($y_8)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Right"))) {
$__t5 = null;;
if ((is_object(((((($y_8)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && ((((((($y_8)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Left"))) {
$__t5 = ($f_5)((((((($y_8)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_5;;
};
if ((is_object(((((($y_8)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && ((((((($y_8)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Right"))) {
$__t6 = null;;
if ((is_object((((((($y_8)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && (((((((($y_8)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Left"))) {
$__t6 = ($g_6)(((((((($y_8)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_6;;
};
if ((is_object((((((($y_8)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && (((((((($y_8)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Right"))) {
$__t7 = null;;
if ((is_object(((((((($y_8)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && ((((((((($y_8)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Left"))) {
$__t7 = ($h_7)((((((((($y_8)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_7;;
};
if ((is_object(((((((($y_8)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && ((((((((($y_8)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Right"))) {
$spin_9_8 = null;
$spin_9_8 = function($v_10 = null) use (&$spin_9_8) {
  $__num = \func_num_args();
  $__tco_var_spin_9_8_8_v_10 = $v_10;
  $__tco_done_spin_9_8_8 = false;
  $__tco_res_spin_9_8_8 = null;
  $__tco_loop_spin_9_8_8 = function($v_10 = null) use (&$__tco_done_spin_9_8_8, &$__tco_var_spin_9_8_8_v_10, &$spin_9_8) {
  $__num = \func_num_args();
  $__tco_done_spin_9_8_8 = true;
  $v_10 = $__tco_var_spin_9_8_8_v_10;
  $__tco_8 = $v_10;
  $__tco_var_spin_9_8_8_v_10 = $__tco_8;
  $__tco_done_spin_9_8_8 = false;
  $__res = null;
  goto __end;;
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  while (($__tco_done_spin_9_8_8 === false)) {
$__tco_res_spin_9_8_8 = ($__tco_loop_spin_9_8_8)($__tco_var_spin_9_8_8_v_10);
};
  $__res = $__tco_res_spin_9_8_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
$__t7 = ($spin_9_8)((($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce')))((((((((($y_8)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}));
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
  return $__num > 9 ? $__res(...\array_slice(\func_get_args(), 9)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Functor_Coproduct_Nested_coproduct7'] = function() { $v = (function() {
  $__fn = function($a_0 = null, $b_1 = null, $c_2 = null, $d_3 = null, $e_4 = null, $f_5 = null, $g_6 = null, $y_7 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 8) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 8);
  }
  $__t0 = null;;
  if ((is_object($y_7) && (($y_7)->{'tag'} === "Left"))) {
$__t0 = ($a_0)(($y_7)->{'value0'});
goto end_branch_0;;
};
  if ((is_object($y_7) && (($y_7)->{'tag'} === "Right"))) {
$__t1 = null;;
if ((is_object(($y_7)->{'value0'}) && ((($y_7)->{'value0'})->{'tag'} === "Left"))) {
$__t1 = ($b_1)((($y_7)->{'value0'})->{'value0'});
goto end_branch_1;;
};
if ((is_object(($y_7)->{'value0'}) && ((($y_7)->{'value0'})->{'tag'} === "Right"))) {
$__t2 = null;;
if ((is_object((($y_7)->{'value0'})->{'value0'}) && (((($y_7)->{'value0'})->{'value0'})->{'tag'} === "Left"))) {
$__t2 = ($c_2)(((($y_7)->{'value0'})->{'value0'})->{'value0'});
goto end_branch_2;;
};
if ((is_object((($y_7)->{'value0'})->{'value0'}) && (((($y_7)->{'value0'})->{'value0'})->{'tag'} === "Right"))) {
$__t3 = null;;
if ((is_object(((($y_7)->{'value0'})->{'value0'})->{'value0'}) && ((((($y_7)->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Left"))) {
$__t3 = ($d_3)((((($y_7)->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_3;;
};
if ((is_object(((($y_7)->{'value0'})->{'value0'})->{'value0'}) && ((((($y_7)->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Right"))) {
$__t4 = null;;
if ((is_object((((($y_7)->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && (((((($y_7)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Left"))) {
$__t4 = ($e_4)(((((($y_7)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_4;;
};
if ((is_object((((($y_7)->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && (((((($y_7)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Right"))) {
$__t5 = null;;
if ((is_object(((((($y_7)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && ((((((($y_7)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Left"))) {
$__t5 = ($f_5)((((((($y_7)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_5;;
};
if ((is_object(((((($y_7)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && ((((((($y_7)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Right"))) {
$__t6 = null;;
if ((is_object((((((($y_7)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && (((((((($y_7)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Left"))) {
$__t6 = ($g_6)(((((((($y_7)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_6;;
};
if ((is_object((((((($y_7)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && (((((((($y_7)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Right"))) {
$spin_8_7 = null;
$spin_8_7 = function($v_9 = null) use (&$spin_8_7) {
  $__num = \func_num_args();
  $__tco_var_spin_8_7_7_v_9 = $v_9;
  $__tco_done_spin_8_7_7 = false;
  $__tco_res_spin_8_7_7 = null;
  $__tco_loop_spin_8_7_7 = function($v_9 = null) use (&$__tco_done_spin_8_7_7, &$__tco_var_spin_8_7_7_v_9, &$spin_8_7) {
  $__num = \func_num_args();
  $__tco_done_spin_8_7_7 = true;
  $v_9 = $__tco_var_spin_8_7_7_v_9;
  $__tco_7 = $v_9;
  $__tco_var_spin_8_7_7_v_9 = $__tco_7;
  $__tco_done_spin_8_7_7 = false;
  $__res = null;
  goto __end;;
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  while (($__tco_done_spin_8_7_7 === false)) {
$__tco_res_spin_8_7_7 = ($__tco_loop_spin_8_7_7)($__tco_var_spin_8_7_7_v_9);
};
  $__res = $__tco_res_spin_8_7_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
$__t6 = ($spin_8_7)((($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce')))(((((((($y_7)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}));
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
  return $__num > 8 ? $__res(...\array_slice(\func_get_args(), 8)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Functor_Coproduct_Nested_coproduct6'] = function() { $v = (function() {
  $__fn = function($a_0 = null, $b_1 = null, $c_2 = null, $d_3 = null, $e_4 = null, $f_5 = null, $y_6 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 7) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 7);
  }
  $__t0 = null;;
  if ((is_object($y_6) && (($y_6)->{'tag'} === "Left"))) {
$__t0 = ($a_0)(($y_6)->{'value0'});
goto end_branch_0;;
};
  if ((is_object($y_6) && (($y_6)->{'tag'} === "Right"))) {
$__t1 = null;;
if ((is_object(($y_6)->{'value0'}) && ((($y_6)->{'value0'})->{'tag'} === "Left"))) {
$__t1 = ($b_1)((($y_6)->{'value0'})->{'value0'});
goto end_branch_1;;
};
if ((is_object(($y_6)->{'value0'}) && ((($y_6)->{'value0'})->{'tag'} === "Right"))) {
$__t2 = null;;
if ((is_object((($y_6)->{'value0'})->{'value0'}) && (((($y_6)->{'value0'})->{'value0'})->{'tag'} === "Left"))) {
$__t2 = ($c_2)(((($y_6)->{'value0'})->{'value0'})->{'value0'});
goto end_branch_2;;
};
if ((is_object((($y_6)->{'value0'})->{'value0'}) && (((($y_6)->{'value0'})->{'value0'})->{'tag'} === "Right"))) {
$__t3 = null;;
if ((is_object(((($y_6)->{'value0'})->{'value0'})->{'value0'}) && ((((($y_6)->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Left"))) {
$__t3 = ($d_3)((((($y_6)->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_3;;
};
if ((is_object(((($y_6)->{'value0'})->{'value0'})->{'value0'}) && ((((($y_6)->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Right"))) {
$__t4 = null;;
if ((is_object((((($y_6)->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && (((((($y_6)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Left"))) {
$__t4 = ($e_4)(((((($y_6)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_4;;
};
if ((is_object((((($y_6)->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && (((((($y_6)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Right"))) {
$__t5 = null;;
if ((is_object(((((($y_6)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && ((((((($y_6)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Left"))) {
$__t5 = ($f_5)((((((($y_6)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_5;;
};
if ((is_object(((((($y_6)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && ((((((($y_6)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Right"))) {
$spin_7_6 = null;
$spin_7_6 = function($v_8 = null) use (&$spin_7_6) {
  $__num = \func_num_args();
  $__tco_var_spin_7_6_6_v_8 = $v_8;
  $__tco_done_spin_7_6_6 = false;
  $__tco_res_spin_7_6_6 = null;
  $__tco_loop_spin_7_6_6 = function($v_8 = null) use (&$__tco_done_spin_7_6_6, &$__tco_var_spin_7_6_6_v_8, &$spin_7_6) {
  $__num = \func_num_args();
  $__tco_done_spin_7_6_6 = true;
  $v_8 = $__tco_var_spin_7_6_6_v_8;
  $__tco_6 = $v_8;
  $__tco_var_spin_7_6_6_v_8 = $__tco_6;
  $__tco_done_spin_7_6_6 = false;
  $__res = null;
  goto __end;;
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  while (($__tco_done_spin_7_6_6 === false)) {
$__tco_res_spin_7_6_6 = ($__tco_loop_spin_7_6_6)($__tco_var_spin_7_6_6_v_8);
};
  $__res = $__tco_res_spin_7_6_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
$__t5 = ($spin_7_6)((($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce')))((((((($y_6)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}));
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
  return $__num > 7 ? $__res(...\array_slice(\func_get_args(), 7)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Functor_Coproduct_Nested_coproduct5'] = function() { $v = (function() {
  $__fn = function($a_0 = null, $b_1 = null, $c_2 = null, $d_3 = null, $e_4 = null, $y_5 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 6) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 6);
  }
  $__t0 = null;;
  if ((is_object($y_5) && (($y_5)->{'tag'} === "Left"))) {
$__t0 = ($a_0)(($y_5)->{'value0'});
goto end_branch_0;;
};
  if ((is_object($y_5) && (($y_5)->{'tag'} === "Right"))) {
$__t1 = null;;
if ((is_object(($y_5)->{'value0'}) && ((($y_5)->{'value0'})->{'tag'} === "Left"))) {
$__t1 = ($b_1)((($y_5)->{'value0'})->{'value0'});
goto end_branch_1;;
};
if ((is_object(($y_5)->{'value0'}) && ((($y_5)->{'value0'})->{'tag'} === "Right"))) {
$__t2 = null;;
if ((is_object((($y_5)->{'value0'})->{'value0'}) && (((($y_5)->{'value0'})->{'value0'})->{'tag'} === "Left"))) {
$__t2 = ($c_2)(((($y_5)->{'value0'})->{'value0'})->{'value0'});
goto end_branch_2;;
};
if ((is_object((($y_5)->{'value0'})->{'value0'}) && (((($y_5)->{'value0'})->{'value0'})->{'tag'} === "Right"))) {
$__t3 = null;;
if ((is_object(((($y_5)->{'value0'})->{'value0'})->{'value0'}) && ((((($y_5)->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Left"))) {
$__t3 = ($d_3)((((($y_5)->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_3;;
};
if ((is_object(((($y_5)->{'value0'})->{'value0'})->{'value0'}) && ((((($y_5)->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Right"))) {
$__t4 = null;;
if ((is_object((((($y_5)->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && (((((($y_5)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Left"))) {
$__t4 = ($e_4)(((((($y_5)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_4;;
};
if ((is_object((((($y_5)->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && (((((($y_5)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Right"))) {
$spin_6_5 = null;
$spin_6_5 = function($v_7 = null) use (&$spin_6_5) {
  $__num = \func_num_args();
  $__tco_var_spin_6_5_5_v_7 = $v_7;
  $__tco_done_spin_6_5_5 = false;
  $__tco_res_spin_6_5_5 = null;
  $__tco_loop_spin_6_5_5 = function($v_7 = null) use (&$__tco_done_spin_6_5_5, &$__tco_var_spin_6_5_5_v_7, &$spin_6_5) {
  $__num = \func_num_args();
  $__tco_done_spin_6_5_5 = true;
  $v_7 = $__tco_var_spin_6_5_5_v_7;
  $__tco_5 = $v_7;
  $__tco_var_spin_6_5_5_v_7 = $__tco_5;
  $__tco_done_spin_6_5_5 = false;
  $__res = null;
  goto __end;;
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  while (($__tco_done_spin_6_5_5 === false)) {
$__tco_res_spin_6_5_5 = ($__tco_loop_spin_6_5_5)($__tco_var_spin_6_5_5_v_7);
};
  $__res = $__tco_res_spin_6_5_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
$__t4 = ($spin_6_5)((($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce')))(((((($y_5)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}));
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
  return $__num > 6 ? $__res(...\array_slice(\func_get_args(), 6)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Functor_Coproduct_Nested_coproduct4'] = function() { $v = (function() {
  $__fn = function($a_0 = null, $b_1 = null, $c_2 = null, $d_3 = null, $y_4 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 5) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 5);
  }
  $__t0 = null;;
  if ((is_object($y_4) && (($y_4)->{'tag'} === "Left"))) {
$__t0 = ($a_0)(($y_4)->{'value0'});
goto end_branch_0;;
};
  if ((is_object($y_4) && (($y_4)->{'tag'} === "Right"))) {
$__t1 = null;;
if ((is_object(($y_4)->{'value0'}) && ((($y_4)->{'value0'})->{'tag'} === "Left"))) {
$__t1 = ($b_1)((($y_4)->{'value0'})->{'value0'});
goto end_branch_1;;
};
if ((is_object(($y_4)->{'value0'}) && ((($y_4)->{'value0'})->{'tag'} === "Right"))) {
$__t2 = null;;
if ((is_object((($y_4)->{'value0'})->{'value0'}) && (((($y_4)->{'value0'})->{'value0'})->{'tag'} === "Left"))) {
$__t2 = ($c_2)(((($y_4)->{'value0'})->{'value0'})->{'value0'});
goto end_branch_2;;
};
if ((is_object((($y_4)->{'value0'})->{'value0'}) && (((($y_4)->{'value0'})->{'value0'})->{'tag'} === "Right"))) {
$__t3 = null;;
if ((is_object(((($y_4)->{'value0'})->{'value0'})->{'value0'}) && ((((($y_4)->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Left"))) {
$__t3 = ($d_3)((((($y_4)->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_3;;
};
if ((is_object(((($y_4)->{'value0'})->{'value0'})->{'value0'}) && ((((($y_4)->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Right"))) {
$spin_5_4 = null;
$spin_5_4 = function($v_6 = null) use (&$spin_5_4) {
  $__num = \func_num_args();
  $__tco_var_spin_5_4_4_v_6 = $v_6;
  $__tco_done_spin_5_4_4 = false;
  $__tco_res_spin_5_4_4 = null;
  $__tco_loop_spin_5_4_4 = function($v_6 = null) use (&$__tco_done_spin_5_4_4, &$__tco_var_spin_5_4_4_v_6, &$spin_5_4) {
  $__num = \func_num_args();
  $__tco_done_spin_5_4_4 = true;
  $v_6 = $__tco_var_spin_5_4_4_v_6;
  $__tco_4 = $v_6;
  $__tco_var_spin_5_4_4_v_6 = $__tco_4;
  $__tco_done_spin_5_4_4 = false;
  $__res = null;
  goto __end;;
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  while (($__tco_done_spin_5_4_4 === false)) {
$__tco_res_spin_5_4_4 = ($__tco_loop_spin_5_4_4)($__tco_var_spin_5_4_4_v_6);
};
  $__res = $__tco_res_spin_5_4_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
$__t3 = ($spin_5_4)((($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce')))((((($y_4)->{'value0'})->{'value0'})->{'value0'})->{'value0'}));
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
  return $__num > 5 ? $__res(...\array_slice(\func_get_args(), 5)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Functor_Coproduct_Nested_coproduct3'] = function() { $v = (function() {
  $__fn = function($a_0 = null, $b_1 = null, $c_2 = null, $y_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__t0 = null;;
  if ((is_object($y_3) && (($y_3)->{'tag'} === "Left"))) {
$__t0 = ($a_0)(($y_3)->{'value0'});
goto end_branch_0;;
};
  if ((is_object($y_3) && (($y_3)->{'tag'} === "Right"))) {
$__t1 = null;;
if ((is_object(($y_3)->{'value0'}) && ((($y_3)->{'value0'})->{'tag'} === "Left"))) {
$__t1 = ($b_1)((($y_3)->{'value0'})->{'value0'});
goto end_branch_1;;
};
if ((is_object(($y_3)->{'value0'}) && ((($y_3)->{'value0'})->{'tag'} === "Right"))) {
$__t2 = null;;
if ((is_object((($y_3)->{'value0'})->{'value0'}) && (((($y_3)->{'value0'})->{'value0'})->{'tag'} === "Left"))) {
$__t2 = ($c_2)(((($y_3)->{'value0'})->{'value0'})->{'value0'});
goto end_branch_2;;
};
if ((is_object((($y_3)->{'value0'})->{'value0'}) && (((($y_3)->{'value0'})->{'value0'})->{'tag'} === "Right"))) {
$spin_4_3 = null;
$spin_4_3 = function($v_5 = null) use (&$spin_4_3) {
  $__num = \func_num_args();
  $__tco_var_spin_4_3_3_v_5 = $v_5;
  $__tco_done_spin_4_3_3 = false;
  $__tco_res_spin_4_3_3 = null;
  $__tco_loop_spin_4_3_3 = function($v_5 = null) use (&$__tco_done_spin_4_3_3, &$__tco_var_spin_4_3_3_v_5, &$spin_4_3) {
  $__num = \func_num_args();
  $__tco_done_spin_4_3_3 = true;
  $v_5 = $__tco_var_spin_4_3_3_v_5;
  $__tco_3 = $v_5;
  $__tco_var_spin_4_3_3_v_5 = $__tco_3;
  $__tco_done_spin_4_3_3 = false;
  $__res = null;
  goto __end;;
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  while (($__tco_done_spin_4_3_3 === false)) {
$__tco_res_spin_4_3_3 = ($__tco_loop_spin_4_3_3)($__tco_var_spin_4_3_3_v_5);
};
  $__res = $__tco_res_spin_4_3_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
$__t2 = ($spin_4_3)((($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce')))(((($y_3)->{'value0'})->{'value0'})->{'value0'}));
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
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Functor_Coproduct_Nested_coproduct2'] = function() { $v = (function() {
  $__fn = function($a_0 = null, $b_1 = null, $y_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if ((is_object($y_2) && (($y_2)->{'tag'} === "Left"))) {
$__t0 = ($a_0)(($y_2)->{'value0'});
goto end_branch_0;;
};
  if ((is_object($y_2) && (($y_2)->{'tag'} === "Right"))) {
$__t1 = null;;
if ((is_object(($y_2)->{'value0'}) && ((($y_2)->{'value0'})->{'tag'} === "Left"))) {
$__t1 = ($b_1)((($y_2)->{'value0'})->{'value0'});
goto end_branch_1;;
};
if ((is_object(($y_2)->{'value0'}) && ((($y_2)->{'value0'})->{'tag'} === "Right"))) {
$spin_3_2 = null;
$spin_3_2 = function($v_4 = null) use (&$spin_3_2) {
  $__num = \func_num_args();
  $__tco_var_spin_3_2_2_v_4 = $v_4;
  $__tco_done_spin_3_2_2 = false;
  $__tco_res_spin_3_2_2 = null;
  $__tco_loop_spin_3_2_2 = function($v_4 = null) use (&$__tco_done_spin_3_2_2, &$__tco_var_spin_3_2_2_v_4, &$spin_3_2) {
  $__num = \func_num_args();
  $__tco_done_spin_3_2_2 = true;
  $v_4 = $__tco_var_spin_3_2_2_v_4;
  $__tco_2 = $v_4;
  $__tco_var_spin_3_2_2_v_4 = $__tco_2;
  $__tco_done_spin_3_2_2 = false;
  $__res = null;
  goto __end;;
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  while (($__tco_done_spin_3_2_2 === false)) {
$__tco_res_spin_3_2_2 = ($__tco_loop_spin_3_2_2)($__tco_var_spin_3_2_2_v_4);
};
  $__res = $__tco_res_spin_3_2_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
$__t1 = ($spin_3_2)((($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce')))((($y_2)->{'value0'})->{'value0'}));
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
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Functor_Coproduct_Nested_coproduct10'] = function() { $v = (function() {
  $__fn = function($a_0 = null, $b_1 = null, $c_2 = null, $d_3 = null, $e_4 = null, $f_5 = null, $g_6 = null, $h_7 = null, $i_8 = null, $j_9 = null, $y_10 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 11) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 11);
  }
  $__t0 = null;;
  if ((is_object($y_10) && (($y_10)->{'tag'} === "Left"))) {
$__t0 = ($a_0)(($y_10)->{'value0'});
goto end_branch_0;;
};
  if ((is_object($y_10) && (($y_10)->{'tag'} === "Right"))) {
$__t1 = null;;
if ((is_object(($y_10)->{'value0'}) && ((($y_10)->{'value0'})->{'tag'} === "Left"))) {
$__t1 = ($b_1)((($y_10)->{'value0'})->{'value0'});
goto end_branch_1;;
};
if ((is_object(($y_10)->{'value0'}) && ((($y_10)->{'value0'})->{'tag'} === "Right"))) {
$__t2 = null;;
if ((is_object((($y_10)->{'value0'})->{'value0'}) && (((($y_10)->{'value0'})->{'value0'})->{'tag'} === "Left"))) {
$__t2 = ($c_2)(((($y_10)->{'value0'})->{'value0'})->{'value0'});
goto end_branch_2;;
};
if ((is_object((($y_10)->{'value0'})->{'value0'}) && (((($y_10)->{'value0'})->{'value0'})->{'tag'} === "Right"))) {
$__t3 = null;;
if ((is_object(((($y_10)->{'value0'})->{'value0'})->{'value0'}) && ((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Left"))) {
$__t3 = ($d_3)((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_3;;
};
if ((is_object(((($y_10)->{'value0'})->{'value0'})->{'value0'}) && ((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Right"))) {
$__t4 = null;;
if ((is_object((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && (((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Left"))) {
$__t4 = ($e_4)(((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_4;;
};
if ((is_object((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && (((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Right"))) {
$__t5 = null;;
if ((is_object(((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && ((((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Left"))) {
$__t5 = ($f_5)((((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_5;;
};
if ((is_object(((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && ((((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Right"))) {
$__t6 = null;;
if ((is_object((((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && (((((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Left"))) {
$__t6 = ($g_6)(((((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_6;;
};
if ((is_object((((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && (((((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Right"))) {
$__t7 = null;;
if ((is_object(((((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && ((((((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Left"))) {
$__t7 = ($h_7)((((((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_7;;
};
if ((is_object(((((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && ((((((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Right"))) {
$__t8 = null;;
if ((is_object((((((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && (((((((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Left"))) {
$__t8 = ($i_8)(((((((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_8;;
};
if ((is_object((((((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && (((((((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Right"))) {
$__t9 = null;;
if ((is_object(((((((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && ((((((((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Left"))) {
$__t9 = ($j_9)((((((((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_9;;
};
if ((is_object(((((((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && ((((((((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Right"))) {
$spin_11_10 = null;
$spin_11_10 = function($v_12 = null) use (&$spin_11_10) {
  $__num = \func_num_args();
  $__tco_var_spin_11_10_10_v_12 = $v_12;
  $__tco_done_spin_11_10_10 = false;
  $__tco_res_spin_11_10_10 = null;
  $__tco_loop_spin_11_10_10 = function($v_12 = null) use (&$__tco_done_spin_11_10_10, &$__tco_var_spin_11_10_10_v_12, &$spin_11_10) {
  $__num = \func_num_args();
  $__tco_done_spin_11_10_10 = true;
  $v_12 = $__tco_var_spin_11_10_10_v_12;
  $__tco_10 = $v_12;
  $__tco_var_spin_11_10_10_v_12 = $__tco_10;
  $__tco_done_spin_11_10_10 = false;
  $__res = null;
  goto __end;;
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  while (($__tco_done_spin_11_10_10 === false)) {
$__tco_res_spin_11_10_10 = ($__tco_loop_spin_11_10_10)($__tco_var_spin_11_10_10_v_12);
};
  $__res = $__tco_res_spin_11_10_10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
$__t9 = ($spin_11_10)((($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce')))((((((((((($y_10)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}));
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
  return $__num > 11 ? $__res(...\array_slice(\func_get_args(), 11)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Functor_Coproduct_Nested_coproduct1'] = function() { $v = function($y_0 = null) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object($y_0) && (($y_0)->{'tag'} === "Left"))) {
$__t0 = ($y_0)->{'value0'};
goto end_branch_0;;
};
  if ((is_object($y_0) && (($y_0)->{'tag'} === "Right"))) {
$spin_1_1 = null;
$spin_1_1 = function($v_2 = null) use (&$spin_1_1) {
  $__num = \func_num_args();
  $__tco_var_spin_1_1_1_v_2 = $v_2;
  $__tco_done_spin_1_1_1 = false;
  $__tco_res_spin_1_1_1 = null;
  $__tco_loop_spin_1_1_1 = function($v_2 = null) use (&$__tco_done_spin_1_1_1, &$__tco_var_spin_1_1_1_v_2, &$spin_1_1) {
  $__num = \func_num_args();
  $__tco_done_spin_1_1_1 = true;
  $v_2 = $__tco_var_spin_1_1_1_v_2;
  $__tco_1 = $v_2;
  $__tco_var_spin_1_1_1_v_2 = $__tco_1;
  $__tco_done_spin_1_1_1 = false;
  $__res = null;
  goto __end;;
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  while (($__tco_done_spin_1_1_1 === false)) {
$__tco_res_spin_1_1_1 = ($__tco_loop_spin_1_1_1)($__tco_var_spin_1_1_1_v_2);
};
  $__res = $__tco_res_spin_1_1_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
$__t0 = ($spin_1_1)((($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce')))(($y_0)->{'value0'}));
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Functor_Coproduct_Nested_at9'] = function() { $v = (function() {
  $__fn = function($b_0 = null, $f_1 = null, $y_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if (((is_object($y_2) && (($y_2)->{'tag'} === "Right")) && ((is_object(($y_2)->{'value0'}) && ((($y_2)->{'value0'})->{'tag'} === "Right")) && ((is_object((($y_2)->{'value0'})->{'value0'}) && (((($y_2)->{'value0'})->{'value0'})->{'tag'} === "Right")) && ((is_object(((($y_2)->{'value0'})->{'value0'})->{'value0'}) && ((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Right")) && ((is_object((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && (((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Right")) && ((is_object(((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && ((((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Right")) && ((is_object((((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && (((((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Right")) && ((is_object(((((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && ((((((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Right")) && (is_object((((((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && (((((((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Left"))))))))))) {
$__t0 = ($f_1)(((((((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_0;;
};
  $__t0 = $b_0;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Functor_Coproduct_Nested_at8'] = function() { $v = (function() {
  $__fn = function($b_0 = null, $f_1 = null, $y_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if (((is_object($y_2) && (($y_2)->{'tag'} === "Right")) && ((is_object(($y_2)->{'value0'}) && ((($y_2)->{'value0'})->{'tag'} === "Right")) && ((is_object((($y_2)->{'value0'})->{'value0'}) && (((($y_2)->{'value0'})->{'value0'})->{'tag'} === "Right")) && ((is_object(((($y_2)->{'value0'})->{'value0'})->{'value0'}) && ((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Right")) && ((is_object((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && (((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Right")) && ((is_object(((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && ((((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Right")) && ((is_object((((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && (((((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Right")) && (is_object(((((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && ((((((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Left")))))))))) {
$__t0 = ($f_1)((((((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_0;;
};
  $__t0 = $b_0;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Functor_Coproduct_Nested_at7'] = function() { $v = (function() {
  $__fn = function($b_0 = null, $f_1 = null, $y_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if (((is_object($y_2) && (($y_2)->{'tag'} === "Right")) && ((is_object(($y_2)->{'value0'}) && ((($y_2)->{'value0'})->{'tag'} === "Right")) && ((is_object((($y_2)->{'value0'})->{'value0'}) && (((($y_2)->{'value0'})->{'value0'})->{'tag'} === "Right")) && ((is_object(((($y_2)->{'value0'})->{'value0'})->{'value0'}) && ((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Right")) && ((is_object((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && (((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Right")) && ((is_object(((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && ((((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Right")) && (is_object((((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && (((((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Left"))))))))) {
$__t0 = ($f_1)(((((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_0;;
};
  $__t0 = $b_0;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Functor_Coproduct_Nested_at6'] = function() { $v = (function() {
  $__fn = function($b_0 = null, $f_1 = null, $y_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if (((is_object($y_2) && (($y_2)->{'tag'} === "Right")) && ((is_object(($y_2)->{'value0'}) && ((($y_2)->{'value0'})->{'tag'} === "Right")) && ((is_object((($y_2)->{'value0'})->{'value0'}) && (((($y_2)->{'value0'})->{'value0'})->{'tag'} === "Right")) && ((is_object(((($y_2)->{'value0'})->{'value0'})->{'value0'}) && ((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Right")) && ((is_object((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && (((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Right")) && (is_object(((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && ((((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Left")))))))) {
$__t0 = ($f_1)((((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_0;;
};
  $__t0 = $b_0;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Functor_Coproduct_Nested_at5'] = function() { $v = (function() {
  $__fn = function($b_0 = null, $f_1 = null, $y_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if (((is_object($y_2) && (($y_2)->{'tag'} === "Right")) && ((is_object(($y_2)->{'value0'}) && ((($y_2)->{'value0'})->{'tag'} === "Right")) && ((is_object((($y_2)->{'value0'})->{'value0'}) && (((($y_2)->{'value0'})->{'value0'})->{'tag'} === "Right")) && ((is_object(((($y_2)->{'value0'})->{'value0'})->{'value0'}) && ((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Right")) && (is_object((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && (((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Left"))))))) {
$__t0 = ($f_1)(((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_0;;
};
  $__t0 = $b_0;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Functor_Coproduct_Nested_at4'] = function() { $v = (function() {
  $__fn = function($b_0 = null, $f_1 = null, $y_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if (((is_object($y_2) && (($y_2)->{'tag'} === "Right")) && ((is_object(($y_2)->{'value0'}) && ((($y_2)->{'value0'})->{'tag'} === "Right")) && ((is_object((($y_2)->{'value0'})->{'value0'}) && (((($y_2)->{'value0'})->{'value0'})->{'tag'} === "Right")) && (is_object(((($y_2)->{'value0'})->{'value0'})->{'value0'}) && ((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Left")))))) {
$__t0 = ($f_1)((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_0;;
};
  $__t0 = $b_0;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Functor_Coproduct_Nested_at3'] = function() { $v = (function() {
  $__fn = function($b_0 = null, $f_1 = null, $y_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if (((is_object($y_2) && (($y_2)->{'tag'} === "Right")) && ((is_object(($y_2)->{'value0'}) && ((($y_2)->{'value0'})->{'tag'} === "Right")) && (is_object((($y_2)->{'value0'})->{'value0'}) && (((($y_2)->{'value0'})->{'value0'})->{'tag'} === "Left"))))) {
$__t0 = ($f_1)(((($y_2)->{'value0'})->{'value0'})->{'value0'});
goto end_branch_0;;
};
  $__t0 = $b_0;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Functor_Coproduct_Nested_at2'] = function() { $v = (function() {
  $__fn = function($b_0 = null, $f_1 = null, $y_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if (((is_object($y_2) && (($y_2)->{'tag'} === "Right")) && (is_object(($y_2)->{'value0'}) && ((($y_2)->{'value0'})->{'tag'} === "Left")))) {
$__t0 = ($f_1)((($y_2)->{'value0'})->{'value0'});
goto end_branch_0;;
};
  $__t0 = $b_0;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Functor_Coproduct_Nested_at10'] = function() { $v = (function() {
  $__fn = function($b_0 = null, $f_1 = null, $y_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if (((is_object($y_2) && (($y_2)->{'tag'} === "Right")) && ((is_object(($y_2)->{'value0'}) && ((($y_2)->{'value0'})->{'tag'} === "Right")) && ((is_object((($y_2)->{'value0'})->{'value0'}) && (((($y_2)->{'value0'})->{'value0'})->{'tag'} === "Right")) && ((is_object(((($y_2)->{'value0'})->{'value0'})->{'value0'}) && ((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Right")) && ((is_object((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && (((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Right")) && ((is_object(((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && ((((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Right")) && ((is_object((((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && (((((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Right")) && ((is_object(((((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && ((((((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Right")) && ((is_object((((((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && (((((((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Right")) && (is_object(((((((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'}) && ((((((((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'tag'} === "Left")))))))))))) {
$__t0 = ($f_1)((((((((((($y_2)->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'})->{'value0'});
goto end_branch_0;;
};
  $__t0 = $b_0;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Functor_Coproduct_Nested_at1'] = function() { $v = (function() {
  $__fn = function($b_0 = null, $f_1 = null, $y_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if ((is_object($y_2) && (($y_2)->{'tag'} === "Left"))) {
$__t0 = ($f_1)(($y_2)->{'value0'});
goto end_branch_0;;
};
  $__t0 = $b_0;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };
































