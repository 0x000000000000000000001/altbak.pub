<?php

namespace Data\Array\ST;

// ALL IMPORTS: Control.Applicative, Control.Bind, Control.Monad.ST, Control.Monad.ST.Internal, Control.Monad.ST.Uncurried, Data.Array.ST, Data.Maybe, Data.Ord, Data.Ordering, Data.Ring, Prelude, Prim
// TO REQUIRE: Control.Applicative, Control.Bind, Control.Monad.ST, Control.Monad.ST.Internal, Control.Monad.ST.Uncurried, Data.Array.ST, Data.Maybe, Data.Ord, Data.Ordering, Data.Ring, Prelude
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Monad.ST/index.php';
require_once __DIR__ . '/../Control.Monad.ST.Internal/index.php';
require_once __DIR__ . '/../Control.Monad.ST.Uncurried/index.php';
require_once __DIR__ . '/../Data.Array.ST/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ordering/index.php';
require_once __DIR__ . '/../Data.Ring/index.php';
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
$ffi_Data_Array_ST = \call_user_func(function() {
  $exports = [];
$new = function() use (&$new) {
    return (object)["value" => []];
};

$peekImpl = function($just, $nothing = null, $i = null, $xs = null) use (&$peekImpl) {
    if (\func_num_args() < 4) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$peekImpl) {

            return $peekImpl(...\array_merge($__args, $more));
        };
    }
    return ($i >= 0 && $i < \count($xs->value)) ? $just($xs->value[$i]) : $nothing;
};

$pokeImpl = function($i, $a = null, $xs = null) use (&$pokeImpl) {
    if (\func_num_args() < 3) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$pokeImpl) {

            return $pokeImpl(...\array_merge($__args, $more));
        };
    }
    if ($i >= 0 && $i < \count($xs->value)) {
        $xs->value[$i] = $a;
        return true;
    }
    return false;
};

$lengthImpl = function($xs) use (&$lengthImpl) {
    return \count($xs->value);
};

$popImpl = function($just, $nothing = null, $xs = null) use (&$popImpl) {
    if (\func_num_args() < 3) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$popImpl) {

            return $popImpl(...\array_merge($__args, $more));
        };
    }
    return \count($xs->value) > 0 ? $just(array_pop($xs->value)) : $nothing;
};

$pushAllImpl = function($as, $xs = null) use (&$pushAllImpl) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$pushAllImpl) {

            return $pushAllImpl(...\array_merge($__args, $more));
        };
    }
    foreach ($as as $a) {
        $xs->value[] = $a;
    }
    return \count($xs->value);
};

$shiftImpl = function($just, $nothing = null, $xs = null) use (&$shiftImpl) {
    if (\func_num_args() < 3) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$shiftImpl) {

            return $shiftImpl(...\array_merge($__args, $more));
        };
    }
    return \count($xs->value) > 0 ? $just(array_shift($xs->value)) : $nothing;
};

$unshiftAllImpl = function($as, $xs = null) use (&$unshiftAllImpl) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$unshiftAllImpl) {

            return $unshiftAllImpl(...\array_merge($__args, $more));
        };
    }
    array_unshift($xs->value, ...$as);
    return \count($xs->value);
};

$spliceImpl = function($i, $howMany = null, $bs = null, $xs = null) use (&$spliceImpl) {
    if (\func_num_args() < 4) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$spliceImpl) {

            return $spliceImpl(...\array_merge($__args, $more));
        };
    }
    return array_splice($xs->value, $i, $howMany, $bs);
};

$unsafeFreezeImpl = function($xs) use (&$unsafeFreezeImpl) {
    return $xs->value;
};

$unsafeThawImpl = function($xs) use (&$unsafeThawImpl) {
    return (object)["value" => $xs];
};

$freezeImpl = function($xs) use (&$freezeImpl) {
    return $xs->value;
};

$thawImpl = function($xs) use (&$thawImpl) {
    return (object)["value" => $xs];
};

$cloneImpl = function($xs) use (&$cloneImpl) {
    return (object)["value" => $xs->value];
};

$sortByImpl = function($compare, $fromOrdering = null, $xs = null) use (&$sortByImpl) {
    if (\func_num_args() < 3) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$sortByImpl) {

            return $sortByImpl(...\array_merge($__args, $more));
        };
    }
    \usort($xs->value, function($a, $b) use ($compare, $fromOrdering) {
        return $fromOrdering($compare($a)($b));
    });
    return $xs;
};

$toAssocArrayImpl = function($xs) use (&$toAssocArrayImpl) {
    $n = \count($xs->value);
    $as = [];
    for ($i = 0; $i < $n; $i++) {
        $as[] = (object)["value" => $xs->value[$i], "index" => $i];
    }
    return $as;
};

$pushImpl = function($a, $xs = null) use (&$pushImpl) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$pushImpl) {

            return $pushImpl(...\array_merge($__args, $more));
        };
    }
    $xs->value[] = $a;
    return \count($xs->value);
};

$exports['new'] = $new;
$exports['peekImpl'] = $peekImpl;
$exports['pokeImpl'] = $pokeImpl;
$exports['lengthImpl'] = $lengthImpl;
$exports['popImpl'] = $popImpl;
$exports['pushAllImpl'] = $pushAllImpl;
$exports['shiftImpl'] = $shiftImpl;
$exports['unshiftAllImpl'] = $unshiftAllImpl;
$exports['spliceImpl'] = $spliceImpl;
$exports['unsafeFreezeImpl'] = $unsafeFreezeImpl;
$exports['unsafeThawImpl'] = $unsafeThawImpl;
$exports['freezeImpl'] = $freezeImpl;
$exports['thawImpl'] = $thawImpl;
$exports['cloneImpl'] = $cloneImpl;
$exports['sortByImpl'] = $sortByImpl;
$exports['toAssocArrayImpl'] = $toAssocArrayImpl;
$exports['pushImpl'] = $pushImpl;
return $exports;
  return $exports;
});
$GLOBALS['Data_Array_ST_cloneImpl'] = ($ffi_Data_Array_ST['cloneImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array_ST_freezeImpl'] = ($ffi_Data_Array_ST['freezeImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array_ST_lengthImpl'] = ($ffi_Data_Array_ST['lengthImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array_ST_new'] = ($ffi_Data_Array_ST['new'] ?? new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array_ST_peekImpl'] = ($ffi_Data_Array_ST['peekImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array_ST_pokeImpl'] = ($ffi_Data_Array_ST['pokeImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array_ST_popImpl'] = ($ffi_Data_Array_ST['popImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array_ST_pushAllImpl'] = ($ffi_Data_Array_ST['pushAllImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array_ST_pushImpl'] = ($ffi_Data_Array_ST['pushImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array_ST_shiftImpl'] = ($ffi_Data_Array_ST['shiftImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array_ST_sortByImpl'] = ($ffi_Data_Array_ST['sortByImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array_ST_spliceImpl'] = ($ffi_Data_Array_ST['spliceImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array_ST_thawImpl'] = ($ffi_Data_Array_ST['thawImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array_ST_toAssocArrayImpl'] = ($ffi_Data_Array_ST['toAssocArrayImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array_ST_unsafeFreezeImpl'] = ($ffi_Data_Array_ST['unsafeFreezeImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array_ST_unsafeThawImpl'] = ($ffi_Data_Array_ST['unsafeThawImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array_ST_unshiftAllImpl'] = ($ffi_Data_Array_ST['unshiftAllImpl'] ?? new class { public function __invoke(...$args) { return $this; } });




// Data_Array_ST_unshiftAll_closure
$GLOBALS['Data_Array_ST_unshiftAll_closure'] = ($GLOBALS['Control_Monad_ST_Uncurried_runSTFn2'])($GLOBALS['Data_Array_ST_unshiftAllImpl']);

// Data_Array_ST_unshiftAll
function majData_majArray_majSmajT_unshiftmajAll($v_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majSmajT_unshiftmajAll';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Array_ST_unshiftAll_closure'])($v_0, $v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_ST_unshiftAll'] = __NAMESPACE__ . '\\majData_majArray_majSmajT_unshiftmajAll';

// Data_Array_ST_unshift
function majData_majArray_majSmajT_unshift($a_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majSmajT_unshift';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Monad_ST_Uncurried_runSTFn2'])($GLOBALS['Data_Array_ST_unshiftAllImpl']))([$a_0]);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_ST_unshift'] = __NAMESPACE__ . '\\majData_majArray_majSmajT_unshift';

// Data_Array_ST_unsafeThaw_closure
$GLOBALS['Data_Array_ST_unsafeThaw_closure'] = ($GLOBALS['Control_Monad_ST_Uncurried_runSTFn1'])($GLOBALS['Data_Array_ST_unsafeThawImpl']);

// Data_Array_ST_unsafeThaw
function majData_majArray_majSmajT_unsafemajThaw($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majSmajT_unsafemajThaw';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Array_ST_unsafeThaw_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_ST_unsafeThaw'] = __NAMESPACE__ . '\\majData_majArray_majSmajT_unsafemajThaw';

// Data_Array_ST_unsafeFreeze_closure
$GLOBALS['Data_Array_ST_unsafeFreeze_closure'] = ($GLOBALS['Control_Monad_ST_Uncurried_runSTFn1'])($GLOBALS['Data_Array_ST_unsafeFreezeImpl']);

// Data_Array_ST_unsafeFreeze
function majData_majArray_majSmajT_unsafemajFreeze($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majSmajT_unsafemajFreeze';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Array_ST_unsafeFreeze_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_ST_unsafeFreeze'] = __NAMESPACE__ . '\\majData_majArray_majSmajT_unsafemajFreeze';

// Data_Array_ST_toAssocArray_closure
$GLOBALS['Data_Array_ST_toAssocArray_closure'] = ($GLOBALS['Control_Monad_ST_Uncurried_runSTFn1'])($GLOBALS['Data_Array_ST_toAssocArrayImpl']);

// Data_Array_ST_toAssocArray
function majData_majArray_majSmajT_tomajAssocmajArray($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majSmajT_tomajAssocmajArray';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Array_ST_toAssocArray_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_ST_toAssocArray'] = __NAMESPACE__ . '\\majData_majArray_majSmajT_tomajAssocmajArray';

// Data_Array_ST_thaw_closure
$GLOBALS['Data_Array_ST_thaw_closure'] = ($GLOBALS['Control_Monad_ST_Uncurried_runSTFn1'])($GLOBALS['Data_Array_ST_thawImpl']);

// Data_Array_ST_thaw
function majData_majArray_majSmajT_thaw($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majSmajT_thaw';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Array_ST_thaw_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_ST_thaw'] = __NAMESPACE__ . '\\majData_majArray_majSmajT_thaw';

// Data_Array_ST_withArray
function majData_majArray_majSmajT_withmajArray($f_0, $xs_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majSmajT_withmajArray';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bindST'])->{'bind'})(\Control\Monad\ST\Uncurried\majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn1($GLOBALS['Data_Array_ST_thawImpl'], $xs_1)))(function($result_2) use ($f_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bindST'])->{'bind'})(($f_0)($result_2)))(function($_dollar__unused_3) use ($result_2) {
  $__num = \func_num_args();
  $__res = \Control\Monad\ST\Uncurried\majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn1($GLOBALS['Data_Array_ST_unsafeFreezeImpl'], $result_2);
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
$GLOBALS['Data_Array_ST_withArray'] = __NAMESPACE__ . '\\majData_majArray_majSmajT_withmajArray';

// Data_Array_ST_splice_closure
$GLOBALS['Data_Array_ST_splice_closure'] = ($GLOBALS['Control_Monad_ST_Uncurried_runSTFn4'])($GLOBALS['Data_Array_ST_spliceImpl']);

// Data_Array_ST_splice
function majData_majArray_majSmajT_splice(int $v_0, $v_1 = null, $v_2 = null, $v_3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majSmajT_splice';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = ($GLOBALS['Data_Array_ST_splice_closure'])($v_0, $v_1, $v_2, $v_3);
  goto __end;;
  __end:
  return 4 < $__num ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
}
$GLOBALS['Data_Array_ST_splice'] = __NAMESPACE__ . '\\majData_majArray_majSmajT_splice';

// Data_Array_ST_sortBy
function majData_majArray_majSmajT_sortmajBy($comp_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majSmajT_sortmajBy';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ((($GLOBALS['Control_Monad_ST_Uncurried_runSTFn3'])($GLOBALS['Data_Array_ST_sortByImpl']))($comp_0))(function($v_1) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v_1 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t0 = 1;
goto end_branch_0;;
};
  if ($v_1 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t0 = 0;
goto end_branch_0;;
};
  if ($v_1 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t0 = -1;
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_ST_sortBy'] = __NAMESPACE__ . '\\majData_majArray_majSmajT_sortmajBy';

// Data_Array_ST_sortWith
function majData_majArray_majSmajT_sortmajWith($dictOrd_0, $f_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majSmajT_sortmajWith';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Array_ST_sortBy'])((function() use ($dictOrd_0, $f_1) {
  $__fn = function($x_2, $y_3 = null) use ($dictOrd_0, $f_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictOrd_0)->{'compare'})(($f_1)($x_2)))(($f_1)($y_3));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})());
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_ST_sortWith'] = __NAMESPACE__ . '\\majData_majArray_majSmajT_sortmajWith';

// Data_Array_ST_sort
function majData_majArray_majSmajT_sort($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majSmajT_sort';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Array_ST_sortBy'])(($dictOrd_0)->{'compare'});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_ST_sort'] = __NAMESPACE__ . '\\majData_majArray_majSmajT_sort';

// Data_Array_ST_shift_closure
$GLOBALS['Data_Array_ST_shift_closure'] = ((($GLOBALS['Control_Monad_ST_Uncurried_runSTFn3'])($GLOBALS['Data_Array_ST_shiftImpl']))($GLOBALS['Data_Maybe_Just']))(new \Data\Maybe\Data_Maybe_Nothing());

// Data_Array_ST_shift
function majData_majArray_majSmajT_shift($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majSmajT_shift';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Array_ST_shift_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_ST_shift'] = __NAMESPACE__ . '\\majData_majArray_majSmajT_shift';

// Data_Array_ST_run
function majData_majArray_majSmajT_run($st_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majSmajT_run';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = \Control\Monad\ST\Internal\majControl_majMonad_majSmajT_majInternal_run(((($GLOBALS['Control_Monad_ST_Internal_bindST'])->{'bind'})($st_0))($GLOBALS['Data_Array_ST_unsafeFreeze']));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_ST_run'] = __NAMESPACE__ . '\\majData_majArray_majSmajT_run';

// Data_Array_ST_pushAll_closure
$GLOBALS['Data_Array_ST_pushAll_closure'] = ($GLOBALS['Control_Monad_ST_Uncurried_runSTFn2'])($GLOBALS['Data_Array_ST_pushAllImpl']);

// Data_Array_ST_pushAll
function majData_majArray_majSmajT_pushmajAll($v_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majSmajT_pushmajAll';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Array_ST_pushAll_closure'])($v_0, $v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_ST_pushAll'] = __NAMESPACE__ . '\\majData_majArray_majSmajT_pushmajAll';

// Data_Array_ST_push_closure
$GLOBALS['Data_Array_ST_push_closure'] = ($GLOBALS['Control_Monad_ST_Uncurried_runSTFn2'])($GLOBALS['Data_Array_ST_pushImpl']);

// Data_Array_ST_push
function majData_majArray_majSmajT_push($v_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majSmajT_push';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Array_ST_push_closure'])($v_0, $v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_ST_push'] = __NAMESPACE__ . '\\majData_majArray_majSmajT_push';

// Data_Array_ST_pop_closure
$GLOBALS['Data_Array_ST_pop_closure'] = ((($GLOBALS['Control_Monad_ST_Uncurried_runSTFn3'])($GLOBALS['Data_Array_ST_popImpl']))($GLOBALS['Data_Maybe_Just']))(new \Data\Maybe\Data_Maybe_Nothing());

// Data_Array_ST_pop
function majData_majArray_majSmajT_pop($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majSmajT_pop';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Array_ST_pop_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_ST_pop'] = __NAMESPACE__ . '\\majData_majArray_majSmajT_pop';

// Data_Array_ST_poke_closure
$GLOBALS['Data_Array_ST_poke_closure'] = ($GLOBALS['Control_Monad_ST_Uncurried_runSTFn3'])($GLOBALS['Data_Array_ST_pokeImpl']);

// Data_Array_ST_poke
function majData_majArray_majSmajT_poke(int $v_0, $v_1 = null, $v_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majSmajT_poke';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($GLOBALS['Data_Array_ST_poke_closure'])($v_0, $v_1, $v_2);
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Array_ST_poke'] = __NAMESPACE__ . '\\majData_majArray_majSmajT_poke';

// Data_Array_ST_peek_closure
$GLOBALS['Data_Array_ST_peek_closure'] = ((($GLOBALS['Control_Monad_ST_Uncurried_runSTFn4'])($GLOBALS['Data_Array_ST_peekImpl']))($GLOBALS['Data_Maybe_Just']))(new \Data\Maybe\Data_Maybe_Nothing());

// Data_Array_ST_peek
function majData_majArray_majSmajT_peek(int $v_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majSmajT_peek';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Array_ST_peek_closure'])($v_0, $v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_ST_peek'] = __NAMESPACE__ . '\\majData_majArray_majSmajT_peek';

// Data_Array_ST_modify
function majData_majArray_majSmajT_modify(int $i_0, $f_1 = null, $xs_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majSmajT_modify';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bindST'])->{'bind'})(\Control\Monad\ST\Uncurried\majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn4($GLOBALS['Data_Array_ST_peekImpl'], $GLOBALS['Data_Maybe_Just'], new \Data\Maybe\Data_Maybe_Nothing(), $i_0, $xs_2)))(function($entry_3) use ($f_1, $i_0, $xs_2) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($entry_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t0 = \Control\Monad\ST\Uncurried\majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn3($GLOBALS['Data_Array_ST_pokeImpl'], $i_0, ($f_1)(($entry_3)->{'value0'}), $xs_2);
goto end_branch_0;;
};
  if ($entry_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t0 = (($GLOBALS['Control_Monad_ST_Internal_applicativeST'])->{'pure'})(false);
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Array_ST_modify'] = __NAMESPACE__ . '\\majData_majArray_majSmajT_modify';

// Data_Array_ST_length_closure
$GLOBALS['Data_Array_ST_length_closure'] = ($GLOBALS['Control_Monad_ST_Uncurried_runSTFn1'])($GLOBALS['Data_Array_ST_lengthImpl']);

// Data_Array_ST_length
function majData_majArray_majSmajT_length($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majSmajT_length';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Array_ST_length_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_ST_length'] = __NAMESPACE__ . '\\majData_majArray_majSmajT_length';

// Data_Array_ST_freeze_closure
$GLOBALS['Data_Array_ST_freeze_closure'] = ($GLOBALS['Control_Monad_ST_Uncurried_runSTFn1'])($GLOBALS['Data_Array_ST_freezeImpl']);

// Data_Array_ST_freeze
function majData_majArray_majSmajT_freeze($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majSmajT_freeze';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Array_ST_freeze_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_ST_freeze'] = __NAMESPACE__ . '\\majData_majArray_majSmajT_freeze';

// Data_Array_ST_clone_closure
$GLOBALS['Data_Array_ST_clone_closure'] = ($GLOBALS['Control_Monad_ST_Uncurried_runSTFn1'])($GLOBALS['Data_Array_ST_cloneImpl']);

// Data_Array_ST_clone
function majData_majArray_majSmajT_clone($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majSmajT_clone';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Array_ST_clone_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_ST_clone'] = __NAMESPACE__ . '\\majData_majArray_majSmajT_clone';

