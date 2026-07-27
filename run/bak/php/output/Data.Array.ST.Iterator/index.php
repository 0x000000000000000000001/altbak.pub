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

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };


final class Data_Array_ST_Iterator_Iterator { public $tag = 'Iterator'; public function __construct(public  $value0, public  $value1) {} }

// Data_Array_ST_Iterator_void_closure
$GLOBALS['Data_Array_ST_Iterator_void_closure'] = (($GLOBALS['Control_Monad_ST_Internal_functorST'])->{'map'})(function($v_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Unit_unit'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});

// Data_Array_ST_Iterator_void
function majData_majArray_majSmajT_majIterator_void($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majSmajT_majIterator_void';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Array_ST_Iterator_void_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_ST_Iterator_void'] = __NAMESPACE__ . '\\majData_majArray_majSmajT_majIterator_void';

// Data_Array_ST_Iterator_Iterator
$GLOBALS['Data_Array_ST_Iterator_Iterator'] = (function() {
  $__fn = function($value0, $value1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Array\ST\Iterator\Data_Array_ST_Iterator_Iterator($value0, $value1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_ST_Iterator_peek
function majData_majArray_majSmajT_majIterator_peek($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majSmajT_majIterator_peek';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bindST'])->{'bind'})(\Control\Monad\ST\Internal\majControl_majMonad_majSmajT_majInternal_read(($v_0)->{'value1'})))(function($i_1) use ($v_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_ST_Internal_applicativeST'])->{'pure'})((($v_0)->{'value0'})($i_1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_ST_Iterator_peek'] = __NAMESPACE__ . '\\majData_majArray_majSmajT_majIterator_peek';

// Data_Array_ST_Iterator_next
function majData_majArray_majSmajT_majIterator_next($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majSmajT_majIterator_next';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = ($v_0)->{'value1'};
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bindST'])->{'bind'})(\Control\Monad\ST\Internal\majControl_majMonad_majSmajT_majInternal_read($__local_var_1_0)))(function($i_2) use ($__local_var_1_0, $v_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bindST'])->{'bind'})(\Control\Monad\ST\Internal\majControl_majMonad_majSmajT_majInternal_modifymajImpl(function($s_3) {
  $__num = \func_num_args();
  $s_prime_4_1 = ($s_3 + 1);
  $__res = (object)["state" => $s_prime_4_1, "value" => $s_prime_4_1];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $__local_var_1_0)))(function($_dollar__unused_3) use ($i_2, $v_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_ST_Internal_applicativeST'])->{'pure'})((($v_0)->{'value0'})($i_2));
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_ST_Iterator_next'] = __NAMESPACE__ . '\\majData_majArray_majSmajT_majIterator_next';

// Data_Array_ST_Iterator_pushWhile
function majData_majArray_majSmajT_majIterator_pushmajWhile($p_0, $iter_1 = null, $array_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majSmajT_majIterator_pushmajWhile';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bindST'])->{'bind'})(\Control\Monad\ST\Internal\majControl_majMonad_majSmajT_majInternal_new(false)))(function($break__3) use ($array_2, $iter_1, $p_0) {
  $__num = \func_num_args();
  $__res = \Control\Monad\ST\Internal\majControl_majMonad_majSmajT_majInternal_while(((($GLOBALS['Control_Monad_ST_Internal_functorST'])->{'map'})(($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'not'}))(\Control\Monad\ST\Internal\majControl_majMonad_majSmajT_majInternal_read($break__3)), ((($GLOBALS['Control_Monad_ST_Internal_bindST'])->{'bind'})(\Data\Array\ST\Iterator\majData_majArray_majSmajT_majIterator_peek($iter_1)))(function($mx_4) use ($array_2, $break__3, $iter_1, $p_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if (($mx_4 instanceof \Data\Maybe\Data_Maybe_Just && ($p_0)(($mx_4)->{'value0'}))) {
$__t0 = ((($GLOBALS['Control_Monad_ST_Internal_bindST'])->{'bind'})(\Control\Monad\ST\Uncurried\majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn2($GLOBALS['Data_Array_ST_pushImpl'], ($mx_4)->{'value0'}, $array_2)))(function($_dollar__unused_5) use ($iter_1) {
  $__num = \func_num_args();
  $__res = \Data\Array\ST\Iterator\majData_majArray_majSmajT_majIterator_void(\Data\Array\ST\Iterator\majData_majArray_majSmajT_majIterator_next($iter_1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
goto end_branch_0;;
};
  $__t0 = \Data\Array\ST\Iterator\majData_majArray_majSmajT_majIterator_void(\Control\Monad\ST\Internal\majControl_majMonad_majSmajT_majInternal_write(true, $break__3));
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
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Array_ST_Iterator_pushWhile'] = __NAMESPACE__ . '\\majData_majArray_majSmajT_majIterator_pushmajWhile';

// Data_Array_ST_Iterator_pushAll_closure
$GLOBALS['Data_Array_ST_Iterator_pushAll_closure'] = ($GLOBALS['Data_Array_ST_Iterator_pushWhile'])(function($v_0) {
  $__num = \func_num_args();
  $__res = true;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});

// Data_Array_ST_Iterator_pushAll
function majData_majArray_majSmajT_majIterator_pushmajAll($v_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majSmajT_majIterator_pushmajAll';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Array_ST_Iterator_pushAll_closure'])($v_0, $v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_ST_Iterator_pushAll'] = __NAMESPACE__ . '\\majData_majArray_majSmajT_majIterator_pushmajAll';

// Data_Array_ST_Iterator_iterator
function majData_majArray_majSmajT_majIterator_iterator($f_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majSmajT_majIterator_iterator';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_functorST'])->{'map'})(($GLOBALS['Data_Array_ST_Iterator_Iterator'])($f_0)))(\Control\Monad\ST\Internal\majControl_majMonad_majSmajT_majInternal_new(0));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_ST_Iterator_iterator'] = __NAMESPACE__ . '\\majData_majArray_majSmajT_majIterator_iterator';

// Data_Array_ST_Iterator_iterate
function majData_majArray_majSmajT_majIterator_iterate($iter_0, $f_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majSmajT_majIterator_iterate';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bindST'])->{'bind'})(\Control\Monad\ST\Internal\majControl_majMonad_majSmajT_majInternal_new(false)))(function($break__2) use ($f_1, $iter_0) {
  $__num = \func_num_args();
  $__res = \Control\Monad\ST\Internal\majControl_majMonad_majSmajT_majInternal_while(((($GLOBALS['Control_Monad_ST_Internal_functorST'])->{'map'})(($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'not'}))(\Control\Monad\ST\Internal\majControl_majMonad_majSmajT_majInternal_read($break__2)), ((($GLOBALS['Control_Monad_ST_Internal_bindST'])->{'bind'})(\Data\Array\ST\Iterator\majData_majArray_majSmajT_majIterator_next($iter_0)))(function($mx_3) use ($break__2, $f_1) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($mx_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t0 = ($f_1)(($mx_3)->{'value0'});
goto end_branch_0;;
};
  if ($mx_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t0 = \Data\Array\ST\Iterator\majData_majArray_majSmajT_majIterator_void(\Control\Monad\ST\Internal\majControl_majMonad_majSmajT_majInternal_write(true, $break__2));
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
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_ST_Iterator_iterate'] = __NAMESPACE__ . '\\majData_majArray_majSmajT_majIterator_iterate';

// Data_Array_ST_Iterator_exhausted_closure
$GLOBALS['Data_Array_ST_Iterator_exhausted_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Control_Monad_ST_Internal_functorST'])->{'map'})($GLOBALS['Data_Maybe_isNothing'])))($GLOBALS['Data_Array_ST_Iterator_peek']);

// Data_Array_ST_Iterator_exhausted
function majData_majArray_majSmajT_majIterator_exhausted($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majSmajT_majIterator_exhausted';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Array_ST_Iterator_exhausted_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_ST_Iterator_exhausted'] = __NAMESPACE__ . '\\majData_majArray_majSmajT_majIterator_exhausted';

