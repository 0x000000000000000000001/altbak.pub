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
\PhpursThunks::$thunks['Data_Array_ST_Iterator_void'] = function() { $v = (($GLOBALS['Control_Monad_ST_Internal_map_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_map_')))(function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}); return $v; };
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
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bind_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_bind_')))((($GLOBALS['Control_Monad_ST_Internal_read'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_read')))(($v_0)->value1)))(function($i_1) use ($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Control_Monad_ST_Internal_pure_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_pure_')))((($v_0)->value0)($i_1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_ST_Iterator_next'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_ST_Iterator_next"), recVars=[];
  $__local_var_1_0 = ($v_0)->value1;
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bind_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_bind_')))((($GLOBALS['Control_Monad_ST_Internal_read'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_read')))($__local_var_1_0)))(function($i_2) use ($__local_var_1_0, $v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bind_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_bind_')))(((($GLOBALS['Control_Monad_ST_Internal_modifyImpl'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_modifyImpl')))(function($s_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $s__prime___4_1 = ((($GLOBALS['Data_Semiring_intAdd'] ?? \PhpursThunks::eval('Data_Semiring_intAdd')))($s_3))(1);
  $__res = (object)["state" => $s__prime___4_1, "value" => $s__prime___4_1];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($__local_var_1_0)))(function($dollar__unused_3) use ($i_2, $v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Control_Monad_ST_Internal_pure_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_pure_')))((($v_0)->value0)($i_2));
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
}; return $v; };
\PhpursThunks::$thunks['Data_Array_ST_Iterator_pushWhile'] = function() { $v = (function() {
  $__fn = function($p_0, $iter_1 = null, $array_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_ST_Iterator_pushWhile"), recVars=[];
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bind_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_bind_')))((($GLOBALS['Control_Monad_ST_Internal_new'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_new')))(false)))(function($break_3) use ($array_2, $iter_1, $p_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_while'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_while')))(((($GLOBALS['Control_Monad_ST_Internal_map_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_map_')))(($GLOBALS['Data_HeytingAlgebra_boolNot'] ?? \PhpursThunks::eval('Data_HeytingAlgebra_boolNot'))))((($GLOBALS['Control_Monad_ST_Internal_read'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_read')))($break_3))))(((($GLOBALS['Control_Monad_ST_Internal_bind_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_bind_')))((($GLOBALS['Data_Array_ST_Iterator_peek'] ?? \PhpursThunks::eval('Data_Array_ST_Iterator_peek')))($iter_1)))(function($mx_4) use ($array_2, $break_3, $iter_1, $p_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if (((is_object($mx_4) && (($mx_4)->tag === "Just")) && ($p_0)(($mx_4)->value0))) {
$__t0 = ((($GLOBALS['Control_Monad_ST_Internal_bind_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_bind_')))(((($GLOBALS['Data_Array_ST_push'] ?? \PhpursThunks::eval('Data_Array_ST_push')))(($mx_4)->value0))($array_2)))(function($dollar__unused_5) use ($iter_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Data_Array_ST_Iterator_void'] ?? \PhpursThunks::eval('Data_Array_ST_Iterator_void')))((($GLOBALS['Data_Array_ST_Iterator_next'] ?? \PhpursThunks::eval('Data_Array_ST_Iterator_next')))($iter_1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
} else {
$__t0 = (($GLOBALS['Data_Array_ST_Iterator_void'] ?? \PhpursThunks::eval('Data_Array_ST_Iterator_void')))(((($GLOBALS['Control_Monad_ST_Internal_write'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_write')))(true))($break_3));
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_map_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_map_')))((($GLOBALS['Data_Array_ST_Iterator_Iterator'] ?? \PhpursThunks::eval('Data_Array_ST_Iterator_Iterator')))($f_0)))((($GLOBALS['Control_Monad_ST_Internal_new'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_new')))(0));
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
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bind_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_bind_')))((($GLOBALS['Control_Monad_ST_Internal_new'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_new')))(false)))(function($break_2) use ($f_1, $iter_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_while'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_while')))(((($GLOBALS['Control_Monad_ST_Internal_map_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_map_')))(($GLOBALS['Data_HeytingAlgebra_boolNot'] ?? \PhpursThunks::eval('Data_HeytingAlgebra_boolNot'))))((($GLOBALS['Control_Monad_ST_Internal_read'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_read')))($break_2))))(((($GLOBALS['Control_Monad_ST_Internal_bind_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_bind_')))((($GLOBALS['Data_Array_ST_Iterator_next'] ?? \PhpursThunks::eval('Data_Array_ST_Iterator_next')))($iter_0)))(function($mx_3) use ($break_2, $f_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($mx_3) && (($mx_3)->tag === "Just"))) {
$__t0 = ($f_1)(($mx_3)->value0);
} else {
if ((is_object($mx_3) && (($mx_3)->tag === "Nothing"))) {
$__t0 = (($GLOBALS['Data_Array_ST_Iterator_void'] ?? \PhpursThunks::eval('Data_Array_ST_Iterator_void')))(((($GLOBALS['Control_Monad_ST_Internal_write'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_write')))(true))($break_2));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_ST_Iterator_exhausted'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Control_Monad_ST_Internal_map_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_map_')))(($GLOBALS['Data_Maybe_isNothing'] ?? \PhpursThunks::eval('Data_Maybe_isNothing')))))(($GLOBALS['Data_Array_ST_Iterator_peek'] ?? \PhpursThunks::eval('Data_Array_ST_Iterator_peek'))); return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };











