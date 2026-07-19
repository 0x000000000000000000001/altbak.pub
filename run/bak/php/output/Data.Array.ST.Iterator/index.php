<?php

namespace Data\Array\ST\Iterator;

// ALL IMPORTS: Control.Applicative, Control.Bind, Control.Monad.ST, Control.Monad.ST.Internal, Control.Monad.ST.Ref, Control.Semigroupoid, Data.Array.ST, Data.Array.ST.Iterator, Data.Function, Data.Functor, Data.HeytingAlgebra, Data.Maybe, Data.Semiring, Prelude, Prim
// TO REQUIRE: Control.Applicative, Control.Bind, Control.Monad.ST, Control.Monad.ST.Internal, Control.Monad.ST.Ref, Control.Semigroupoid, Data.Array.ST, Data.Array.ST.Iterator, Data.Function, Data.Functor, Data.HeytingAlgebra, Data.Maybe, Data.Semiring, Prelude
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Monad.ST/index.php';
require_once __DIR__ . '/../Control.Monad.ST.Internal/index.php';
require_once __DIR__ . '/../Control.Monad.ST.Ref/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Array.ST/index.php';
require_once __DIR__ . '/../Data.Array.ST.Iterator/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.HeytingAlgebra/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Semiring/index.php';
require_once __DIR__ . '/../Prelude/index.php';

if (!class_exists(__NAMESPACE__ . '\\Phpurs_Data0')) {
  class Phpurs_Data0 { public $tag; public function __construct($t) { $this->tag = $t; } }
  class Phpurs_Data1 { public $tag; public $value0; public function __construct($t, $value0) { $this->tag = $t; $this->value0 = $value0; } }
  class Phpurs_Data2 { public $tag; public $value0, $value1; public function __construct($t, $value0, $value1) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; } }
  class Phpurs_Data3 { public $tag; public $value0, $value1, $value2; public function __construct($t, $value0, $value1, $value2) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; } }
  class Phpurs_Data4 { public $tag; public $value0, $value1, $value2, $value3; public function __construct($t, $value0, $value1, $value2, $value3) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; $this->value3 = $value3; } }
  class Phpurs_Data5 { public $tag; public $value0, $value1, $value2, $value3, $value4; public function __construct($t, $value0, $value1, $value2, $value3, $value4) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; $this->value3 = $value3; $this->value4 = $value4; } }
  class Phpurs_Data6 { public $tag; public $value0, $value1, $value2, $value3, $value4, $value5; public function __construct($t, $value0, $value1, $value2, $value3, $value4, $value5) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; $this->value3 = $value3; $this->value4 = $value4; $this->value5 = $value5; } }
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
\PhpursThunks::$thunks['Data_Array_ST_Iterator_Iterator'] = function() { $v = (function() {
  $__fn = function($value0, $value1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("Iterator", $value0, $value1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_ST_Iterator_peek'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_ST_Iterator_peek"), recVars=[];
  $__local_var_1_0 = ($v_0)->value1;
  $i_2_1 = ("TODO_PrimEffect")();
  $__res = (($v_0)->value0)($i_2_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_ST_Iterator_next'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_ST_Iterator_next"), recVars=[];
  $__local_var_1_0 = ($v_0)->value1;
  $i_2_1 = ("TODO_PrimEffect")();
  $dollar__unused_3_2 = (((($GLOBALS['Control_Monad_ST_Internal_modifyImpl'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_modifyImpl')))(function($s_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $s__prime___4_2 = ($s_3 + 1);
  $__res = (object)["state" => $s__prime___4_2, "value" => $s__prime___4_2];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($__local_var_1_0))();
  $__res = (($v_0)->value0)($i_2_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_ST_Iterator_pushWhile'] = function() { $v = (function() {
  $__fn = function($p_0, $iter_1 = null, $array_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_ST_Iterator_pushWhile"), recVars=[];
  $break_3_0 = ("TODO_PrimEffect")();
  $__local_var_4_1 = ("TODO_PrimEffect")();
  $__local_var_4_2 = ($iter_1)->value1;
  $i_5_3 = ("TODO_PrimEffect")();
  $mx_6_4 = (($iter_1)->value0)($i_5_3);
  if (((is_object($mx_6_4) && (($mx_6_4)->tag === "Just")) && ($p_0)(($mx_6_4)->value0))) {
$dollar__unused_7_7 = (((($GLOBALS['Data_Array_ST_push'] ?? \PhpursThunks::eval('Data_Array_ST_push')))(($mx_6_4)->value0))($array_2))();
$__local_var_8_8 = ((($GLOBALS['Data_Array_ST_Iterator_next'] ?? \PhpursThunks::eval('Data_Array_ST_Iterator_next')))($iter_1))();
$__t6 = ($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit'));
} else {
$__local_var_7_5 = ("TODO_PrimEffect")();
$__t6 = ($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit'));
};
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_while'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_while')))(( ! $__local_var_4_1)))($__t6);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_ST_Iterator_pushAll'] = function() { $v = (($GLOBALS['Data_Array_ST_Iterator_pushWhile'] ?? \PhpursThunks::eval('Data_Array_ST_Iterator_pushWhile')))(function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = true;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}); return $v; };
\PhpursThunks::$thunks['Data_Array_ST_Iterator_iterator'] = function() { $v = function($f_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_ST_Iterator_iterator"), recVars=[];
  $__local_var_1_0 = (($GLOBALS['Data_Array_ST_Iterator_Iterator'] ?? \PhpursThunks::eval('Data_Array_ST_Iterator_Iterator')))($f_0);
  $__local_var_2_1 = ("TODO_PrimEffect")();
  $__res = ($__local_var_1_0)($__local_var_2_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_ST_Iterator_iterate'] = function() { $v = (function() {
  $__fn = function($iter_0, $f_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_ST_Iterator_iterate"), recVars=[];
  $break_2_0 = ("TODO_PrimEffect")();
  $__local_var_3_1 = ("TODO_PrimEffect")();
  $__local_var_3_2 = (($GLOBALS['Data_Array_ST_Iterator_next'] ?? \PhpursThunks::eval('Data_Array_ST_Iterator_next')))($iter_0);
  $mx_4_3 = ($__local_var_3_2)();
  if ((is_object($mx_4_3) && (($mx_4_3)->tag === "Just"))) {
$__t4 = ($f_1)(($mx_4_3)->value0);
} else {
if ((is_object($mx_4_3) && (($mx_4_3)->tag === "Nothing"))) {
$__local_var_5_5 = ("TODO_PrimEffect")();
$__t4 = ($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit'));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t4 = null;
};
};
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_while'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_while')))(( ! $__local_var_3_1)))($__t4);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_ST_Iterator_exhausted'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(function($__local_var_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__local_var_1_0 = ($__local_var_0)();
  if ((is_object($__local_var_1_0) && (($__local_var_1_0)->tag === "Nothing"))) {
$__t1 = true;
} else {
if ((is_object($__local_var_1_0) && (($__local_var_1_0)->tag === "Just"))) {
$__t1 = false;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($GLOBALS['Data_Array_ST_Iterator_peek'] ?? \PhpursThunks::eval('Data_Array_ST_Iterator_peek'))); return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };










