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
\PhpursThunks::$thunks['Data_Array_ST_Iterator_void'] = function() { $v = (($GLOBALS['Control_Monad_ST_Internal_map_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_map_')))(function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}); return $v; };
\PhpursThunks::$thunks['Data_Array_ST_Iterator_Iterator'] = function() { $v = (function() {
  $__fn = function($value0 = null, $value1 = null) use (&$__fn) {
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
\PhpursThunks::$thunks['Data_Array_ST_Iterator_peek'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bind_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_bind_')))((($GLOBALS['Control_Monad_ST_Internal_read'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_read')))(($v_0)->{'value1'})))(function($i_1 = null) use ($v_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_ST_Internal_pure_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_pure_')))((($v_0)->{'value0'})($i_1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_ST_Iterator_next'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = ($v_0)->{'value1'};
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bind_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_bind_')))((($GLOBALS['Control_Monad_ST_Internal_read'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_read')))($__local_var_1_0)))(function($i_2 = null) use ($__local_var_1_0, $v_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bind_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_bind_')))(((($GLOBALS['Control_Monad_ST_Internal_modifyImpl'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_modifyImpl')))(function($s_3 = null) {
  $__num = \func_num_args();
  $s__prime___4_1 = ((($GLOBALS['Data_Semiring_intAdd'] ?? \PhpursThunks::eval('Data_Semiring_intAdd')))($s_3))(1);
  $__res = (object)["state" => $s__prime___4_1, "value" => $s__prime___4_1];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($__local_var_1_0)))(function($dollar__unused_3 = null) use ($i_2, $v_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_ST_Internal_pure_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_pure_')))((($v_0)->{'value0'})($i_2));
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
  $__fn = function($p_0 = null, $iter_1 = null, $array_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bind_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_bind_')))((($GLOBALS['Control_Monad_ST_Internal_new'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_new')))(false)))(function($break_3 = null) use ($array_2, $iter_1, $p_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_while'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_while')))(((($GLOBALS['Control_Monad_ST_Internal_map_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_map_')))(($GLOBALS['Data_HeytingAlgebra_boolNot'] ?? \PhpursThunks::eval('Data_HeytingAlgebra_boolNot'))))((($GLOBALS['Control_Monad_ST_Internal_read'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_read')))($break_3))))(((($GLOBALS['Control_Monad_ST_Internal_bind_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_bind_')))((($GLOBALS['Data_Array_ST_Iterator_peek'] ?? \PhpursThunks::eval('Data_Array_ST_Iterator_peek')))($iter_1)))(function($mx_4 = null) use ($array_2, $break_3, $iter_1, $p_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if (((is_object($mx_4) && (($mx_4)->{'tag'} === "Just")) && ($p_0)(($mx_4)->{'value0'}))) {
$__t0 = ((($GLOBALS['Control_Monad_ST_Internal_bind_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_bind_')))(((($GLOBALS['Data_Array_ST_push'] ?? \PhpursThunks::eval('Data_Array_ST_push')))(($mx_4)->{'value0'}))($array_2)))(function($dollar__unused_5 = null) use ($iter_1) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Array_ST_Iterator_void'] ?? \PhpursThunks::eval('Data_Array_ST_Iterator_void')))((($GLOBALS['Data_Array_ST_Iterator_next'] ?? \PhpursThunks::eval('Data_Array_ST_Iterator_next')))($iter_1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
goto end_branch_0;;
};
  $__t0 = (($GLOBALS['Data_Array_ST_Iterator_void'] ?? \PhpursThunks::eval('Data_Array_ST_Iterator_void')))(((($GLOBALS['Control_Monad_ST_Internal_write'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_write')))(true))($break_3));
  end_branch_0:;
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
\PhpursThunks::$thunks['Data_Array_ST_Iterator_pushAll'] = function() { $v = (($GLOBALS['Data_Array_ST_Iterator_pushWhile'] ?? \PhpursThunks::eval('Data_Array_ST_Iterator_pushWhile')))(function($v_0 = null) {
  $__num = \func_num_args();
  $__res = true;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}); return $v; };
\PhpursThunks::$thunks['Data_Array_ST_Iterator_iterator'] = function() { $v = function($f_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_map_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_map_')))((($GLOBALS['Data_Array_ST_Iterator_Iterator'] ?? \PhpursThunks::eval('Data_Array_ST_Iterator_Iterator')))($f_0)))((($GLOBALS['Control_Monad_ST_Internal_new'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_new')))(0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_ST_Iterator_iterate'] = function() { $v = (function() {
  $__fn = function($iter_0 = null, $f_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bind_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_bind_')))((($GLOBALS['Control_Monad_ST_Internal_new'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_new')))(false)))(function($break_2 = null) use ($f_1, $iter_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_while'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_while')))(((($GLOBALS['Control_Monad_ST_Internal_map_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_map_')))(($GLOBALS['Data_HeytingAlgebra_boolNot'] ?? \PhpursThunks::eval('Data_HeytingAlgebra_boolNot'))))((($GLOBALS['Control_Monad_ST_Internal_read'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_read')))($break_2))))(((($GLOBALS['Control_Monad_ST_Internal_bind_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_bind_')))((($GLOBALS['Data_Array_ST_Iterator_next'] ?? \PhpursThunks::eval('Data_Array_ST_Iterator_next')))($iter_0)))(function($mx_3 = null) use ($break_2, $f_1) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object($mx_3) && (($mx_3)->{'tag'} === "Just"))) {
$__t0 = ($f_1)(($mx_3)->{'value0'});
goto end_branch_0;;
};
  if ((is_object($mx_3) && (($mx_3)->{'tag'} === "Nothing"))) {
$__t0 = (($GLOBALS['Data_Array_ST_Iterator_void'] ?? \PhpursThunks::eval('Data_Array_ST_Iterator_void')))(((($GLOBALS['Control_Monad_ST_Internal_write'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_write')))(true))($break_2));
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
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











