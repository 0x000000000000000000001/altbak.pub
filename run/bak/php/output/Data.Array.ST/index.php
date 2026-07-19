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
\PhpursThunks::$thunks['Data_Array_ST_unshiftAll'] = function() { $v = (($GLOBALS['Control_Monad_ST_Uncurried_runSTFn2'] ?? \PhpursThunks::eval('Control_Monad_ST_Uncurried_runSTFn2')))(($GLOBALS['Data_Array_ST_unshiftAllImpl'] ?? \PhpursThunks::eval('Data_Array_ST_unshiftAllImpl'))); return $v; };
\PhpursThunks::$thunks['Data_Array_ST_unshift'] = function() { $v = function($a_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_ST_unshift"), recVars=[];
  $__res = ((($GLOBALS['Control_Monad_ST_Uncurried_runSTFn2'] ?? \PhpursThunks::eval('Control_Monad_ST_Uncurried_runSTFn2')))(($GLOBALS['Data_Array_ST_unshiftAllImpl'] ?? \PhpursThunks::eval('Data_Array_ST_unshiftAllImpl'))))([$a_0]);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_ST_unsafeThaw'] = function() { $v = (($GLOBALS['Control_Monad_ST_Uncurried_runSTFn1'] ?? \PhpursThunks::eval('Control_Monad_ST_Uncurried_runSTFn1')))(($GLOBALS['Data_Array_ST_unsafeThawImpl'] ?? \PhpursThunks::eval('Data_Array_ST_unsafeThawImpl'))); return $v; };
\PhpursThunks::$thunks['Data_Array_ST_unsafeFreeze'] = function() { $v = (($GLOBALS['Control_Monad_ST_Uncurried_runSTFn1'] ?? \PhpursThunks::eval('Control_Monad_ST_Uncurried_runSTFn1')))(($GLOBALS['Data_Array_ST_unsafeFreezeImpl'] ?? \PhpursThunks::eval('Data_Array_ST_unsafeFreezeImpl'))); return $v; };
\PhpursThunks::$thunks['Data_Array_ST_toAssocArray'] = function() { $v = (($GLOBALS['Control_Monad_ST_Uncurried_runSTFn1'] ?? \PhpursThunks::eval('Control_Monad_ST_Uncurried_runSTFn1')))(($GLOBALS['Data_Array_ST_toAssocArrayImpl'] ?? \PhpursThunks::eval('Data_Array_ST_toAssocArrayImpl'))); return $v; };
\PhpursThunks::$thunks['Data_Array_ST_thaw'] = function() { $v = (($GLOBALS['Control_Monad_ST_Uncurried_runSTFn1'] ?? \PhpursThunks::eval('Control_Monad_ST_Uncurried_runSTFn1')))(($GLOBALS['Data_Array_ST_thawImpl'] ?? \PhpursThunks::eval('Data_Array_ST_thawImpl'))); return $v; };
\PhpursThunks::$thunks['Data_Array_ST_withArray'] = function() { $v = (function() {
  $__fn = function($f_0, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_ST_withArray"), recVars=[];
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bind_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_bind_')))((($GLOBALS['Data_Array_ST_thaw'] ?? \PhpursThunks::eval('Data_Array_ST_thaw')))($xs_1)))(function($result_2) use ($f_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bind_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_bind_')))(($f_0)($result_2)))(function($dollar__unused_3) use ($result_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Data_Array_ST_unsafeFreeze'] ?? \PhpursThunks::eval('Data_Array_ST_unsafeFreeze')))($result_2);
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
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_ST_splice'] = function() { $v = (($GLOBALS['Control_Monad_ST_Uncurried_runSTFn4'] ?? \PhpursThunks::eval('Control_Monad_ST_Uncurried_runSTFn4')))(($GLOBALS['Data_Array_ST_spliceImpl'] ?? \PhpursThunks::eval('Data_Array_ST_spliceImpl'))); return $v; };
\PhpursThunks::$thunks['Data_Array_ST_sortBy'] = function() { $v = function($comp_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_ST_sortBy"), recVars=[];
  $__res = (((($GLOBALS['Control_Monad_ST_Uncurried_runSTFn3'] ?? \PhpursThunks::eval('Control_Monad_ST_Uncurried_runSTFn3')))(($GLOBALS['Data_Array_ST_sortByImpl'] ?? \PhpursThunks::eval('Data_Array_ST_sortByImpl'))))($comp_0))(function($v_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v_1) && (($v_1)->tag === "GT"))) {
$__t0 = 1;
} else {
if ((is_object($v_1) && (($v_1)->tag === "EQ"))) {
$__t0 = 0;
} else {
if ((is_object($v_1) && (($v_1)->tag === "LT"))) {
$__t0 = ((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))(0))(1);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_ST_sortWith'] = function() { $v = (function() {
  $__fn = function($dictOrd_0, $f_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_ST_sortWith"), recVars=[];
  $__res = (($GLOBALS['Data_Array_ST_sortBy'] ?? \PhpursThunks::eval('Data_Array_ST_sortBy')))((function() use ($dictOrd_0, $f_1) {
  $__fn = function($x_2, $y_3 = null) use ($dictOrd_0, $f_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($dictOrd_0)->compare)(($f_1)($x_2)))(($f_1)($y_3));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})());
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Array_ST_sort'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_ST_sort"), recVars=[];
  $__res = (($GLOBALS['Data_Array_ST_sortBy'] ?? \PhpursThunks::eval('Data_Array_ST_sortBy')))(($dictOrd_0)->compare);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_ST_shift'] = function() { $v = (((($GLOBALS['Control_Monad_ST_Uncurried_runSTFn3'] ?? \PhpursThunks::eval('Control_Monad_ST_Uncurried_runSTFn3')))(($GLOBALS['Data_Array_ST_shiftImpl'] ?? \PhpursThunks::eval('Data_Array_ST_shiftImpl'))))(($GLOBALS['Data_Maybe_Just'] ?? \PhpursThunks::eval('Data_Maybe_Just'))))(new Phpurs_Data0("Nothing")); return $v; };
\PhpursThunks::$thunks['Data_Array_ST_run'] = function() { $v = function($st_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_ST_run"), recVars=[];
  $__res = (($GLOBALS['Control_Monad_ST_Internal_run'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_run')))(((($GLOBALS['Control_Monad_ST_Internal_bind_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_bind_')))($st_0))(($GLOBALS['Data_Array_ST_unsafeFreeze'] ?? \PhpursThunks::eval('Data_Array_ST_unsafeFreeze'))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_ST_pushAll'] = function() { $v = (($GLOBALS['Control_Monad_ST_Uncurried_runSTFn2'] ?? \PhpursThunks::eval('Control_Monad_ST_Uncurried_runSTFn2')))(($GLOBALS['Data_Array_ST_pushAllImpl'] ?? \PhpursThunks::eval('Data_Array_ST_pushAllImpl'))); return $v; };
\PhpursThunks::$thunks['Data_Array_ST_push'] = function() { $v = (($GLOBALS['Control_Monad_ST_Uncurried_runSTFn2'] ?? \PhpursThunks::eval('Control_Monad_ST_Uncurried_runSTFn2')))(($GLOBALS['Data_Array_ST_pushImpl'] ?? \PhpursThunks::eval('Data_Array_ST_pushImpl'))); return $v; };
\PhpursThunks::$thunks['Data_Array_ST_pop'] = function() { $v = (((($GLOBALS['Control_Monad_ST_Uncurried_runSTFn3'] ?? \PhpursThunks::eval('Control_Monad_ST_Uncurried_runSTFn3')))(($GLOBALS['Data_Array_ST_popImpl'] ?? \PhpursThunks::eval('Data_Array_ST_popImpl'))))(($GLOBALS['Data_Maybe_Just'] ?? \PhpursThunks::eval('Data_Maybe_Just'))))(new Phpurs_Data0("Nothing")); return $v; };
\PhpursThunks::$thunks['Data_Array_ST_poke'] = function() { $v = (($GLOBALS['Control_Monad_ST_Uncurried_runSTFn3'] ?? \PhpursThunks::eval('Control_Monad_ST_Uncurried_runSTFn3')))(($GLOBALS['Data_Array_ST_pokeImpl'] ?? \PhpursThunks::eval('Data_Array_ST_pokeImpl'))); return $v; };
\PhpursThunks::$thunks['Data_Array_ST_peek'] = function() { $v = (((($GLOBALS['Control_Monad_ST_Uncurried_runSTFn4'] ?? \PhpursThunks::eval('Control_Monad_ST_Uncurried_runSTFn4')))(($GLOBALS['Data_Array_ST_peekImpl'] ?? \PhpursThunks::eval('Data_Array_ST_peekImpl'))))(($GLOBALS['Data_Maybe_Just'] ?? \PhpursThunks::eval('Data_Maybe_Just'))))(new Phpurs_Data0("Nothing")); return $v; };
\PhpursThunks::$thunks['Data_Array_ST_modify'] = function() { $v = (function() {
  $__fn = function($i_0, $f_1 = null, $xs_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Array_ST_modify"), recVars=[];
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bind_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_bind_')))(((($GLOBALS['Data_Array_ST_peek'] ?? \PhpursThunks::eval('Data_Array_ST_peek')))($i_0))($xs_2)))(function($entry_3) use ($f_1, $i_0, $xs_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($entry_3) && (($entry_3)->tag === "Just"))) {
$__t0 = (((($GLOBALS['Data_Array_ST_poke'] ?? \PhpursThunks::eval('Data_Array_ST_poke')))($i_0))(($f_1)(($entry_3)->value0)))($xs_2);
} else {
if ((is_object($entry_3) && (($entry_3)->tag === "Nothing"))) {
$__t0 = (($GLOBALS['Control_Monad_ST_Internal_pure_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_pure_')))(false);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
  $__res = $__t0;
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
\PhpursThunks::$thunks['Data_Array_ST_length'] = function() { $v = (($GLOBALS['Control_Monad_ST_Uncurried_runSTFn1'] ?? \PhpursThunks::eval('Control_Monad_ST_Uncurried_runSTFn1')))(($GLOBALS['Data_Array_ST_lengthImpl'] ?? \PhpursThunks::eval('Data_Array_ST_lengthImpl'))); return $v; };
\PhpursThunks::$thunks['Data_Array_ST_freeze'] = function() { $v = (($GLOBALS['Control_Monad_ST_Uncurried_runSTFn1'] ?? \PhpursThunks::eval('Control_Monad_ST_Uncurried_runSTFn1')))(($GLOBALS['Data_Array_ST_freezeImpl'] ?? \PhpursThunks::eval('Data_Array_ST_freezeImpl'))); return $v; };
\PhpursThunks::$thunks['Data_Array_ST_clone'] = function() { $v = (($GLOBALS['Control_Monad_ST_Uncurried_runSTFn1'] ?? \PhpursThunks::eval('Control_Monad_ST_Uncurried_runSTFn1')))(($GLOBALS['Data_Array_ST_cloneImpl'] ?? \PhpursThunks::eval('Data_Array_ST_cloneImpl'))); return $v; };
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
\PhpursThunks::$thunks['Data_Array_ST_cloneImpl'] = function() use (&$ffi_Data_Array_ST) { return $ffi_Data_Array_ST['cloneImpl']; };
\PhpursThunks::$thunks['Data_Array_ST_freezeImpl'] = function() use (&$ffi_Data_Array_ST) { return $ffi_Data_Array_ST['freezeImpl']; };
\PhpursThunks::$thunks['Data_Array_ST_lengthImpl'] = function() use (&$ffi_Data_Array_ST) { return $ffi_Data_Array_ST['lengthImpl']; };
\PhpursThunks::$thunks['Data_Array_ST_new'] = function() use (&$ffi_Data_Array_ST) { return $ffi_Data_Array_ST['new']; };
\PhpursThunks::$thunks['Data_Array_ST_peekImpl'] = function() use (&$ffi_Data_Array_ST) { return $ffi_Data_Array_ST['peekImpl']; };
\PhpursThunks::$thunks['Data_Array_ST_pokeImpl'] = function() use (&$ffi_Data_Array_ST) { return $ffi_Data_Array_ST['pokeImpl']; };
\PhpursThunks::$thunks['Data_Array_ST_popImpl'] = function() use (&$ffi_Data_Array_ST) { return $ffi_Data_Array_ST['popImpl']; };
\PhpursThunks::$thunks['Data_Array_ST_pushAllImpl'] = function() use (&$ffi_Data_Array_ST) { return $ffi_Data_Array_ST['pushAllImpl']; };
\PhpursThunks::$thunks['Data_Array_ST_pushImpl'] = function() use (&$ffi_Data_Array_ST) { return $ffi_Data_Array_ST['pushImpl']; };
\PhpursThunks::$thunks['Data_Array_ST_shiftImpl'] = function() use (&$ffi_Data_Array_ST) { return $ffi_Data_Array_ST['shiftImpl']; };
\PhpursThunks::$thunks['Data_Array_ST_sortByImpl'] = function() use (&$ffi_Data_Array_ST) { return $ffi_Data_Array_ST['sortByImpl']; };
\PhpursThunks::$thunks['Data_Array_ST_spliceImpl'] = function() use (&$ffi_Data_Array_ST) { return $ffi_Data_Array_ST['spliceImpl']; };
\PhpursThunks::$thunks['Data_Array_ST_thawImpl'] = function() use (&$ffi_Data_Array_ST) { return $ffi_Data_Array_ST['thawImpl']; };
\PhpursThunks::$thunks['Data_Array_ST_toAssocArrayImpl'] = function() use (&$ffi_Data_Array_ST) { return $ffi_Data_Array_ST['toAssocArrayImpl']; };
\PhpursThunks::$thunks['Data_Array_ST_unsafeFreezeImpl'] = function() use (&$ffi_Data_Array_ST) { return $ffi_Data_Array_ST['unsafeFreezeImpl']; };
\PhpursThunks::$thunks['Data_Array_ST_unsafeThawImpl'] = function() use (&$ffi_Data_Array_ST) { return $ffi_Data_Array_ST['unsafeThawImpl']; };
\PhpursThunks::$thunks['Data_Array_ST_unshiftAllImpl'] = function() use (&$ffi_Data_Array_ST) { return $ffi_Data_Array_ST['unshiftAllImpl']; };
























