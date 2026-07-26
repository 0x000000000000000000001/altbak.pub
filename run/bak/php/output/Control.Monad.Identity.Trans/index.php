<?php

namespace Control\Monad\Identity\Trans;

// ALL IMPORTS: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Comonad, Control.Extend, Control.Monad, Control.Monad.Cont.Class, Control.Monad.Error.Class, Control.Monad.Identity.Trans, Control.Monad.Reader.Class, Control.Monad.Rec.Class, Control.Monad.ST.Class, Control.Monad.State.Class, Control.Monad.Trans.Class, Control.Monad.Writer.Class, Control.MonadPlus, Control.Plus, Control.Semigroupoid, Data.Eq, Data.Foldable, Data.Functor, Data.Newtype, Data.Ord, Data.Traversable, Effect.Class, Prelude, Prim
// TO REQUIRE: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Comonad, Control.Extend, Control.Monad, Control.Monad.Cont.Class, Control.Monad.Error.Class, Control.Monad.Identity.Trans, Control.Monad.Reader.Class, Control.Monad.Rec.Class, Control.Monad.ST.Class, Control.Monad.State.Class, Control.Monad.Trans.Class, Control.Monad.Writer.Class, Control.MonadPlus, Control.Plus, Control.Semigroupoid, Data.Eq, Data.Foldable, Data.Functor, Data.Newtype, Data.Ord, Data.Traversable, Effect.Class, Prelude
require_once __DIR__ . '/../Control.Alt/index.php';
require_once __DIR__ . '/../Control.Alternative/index.php';
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Comonad/index.php';
require_once __DIR__ . '/../Control.Extend/index.php';
require_once __DIR__ . '/../Control.Monad/index.php';
require_once __DIR__ . '/../Control.Monad.Cont.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Error.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Identity.Trans/index.php';
require_once __DIR__ . '/../Control.Monad.Reader.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Rec.Class/index.php';
require_once __DIR__ . '/../Control.Monad.ST.Class/index.php';
require_once __DIR__ . '/../Control.Monad.State.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Trans.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Writer.Class/index.php';
require_once __DIR__ . '/../Control.MonadPlus/index.php';
require_once __DIR__ . '/../Control.Plus/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Foldable/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Newtype/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Traversable/index.php';
require_once __DIR__ . '/../Effect.Class/index.php';
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


// Control_Monad_Identity_Trans_IdentityT
$GLOBALS['Control_Monad_Identity_Trans_IdentityT'] = function($x_0 = null) {
  $__num = \func_num_args();
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Identity_Trans_monadSTIdentityT
$GLOBALS['Control_Monad_Identity_Trans_monadSTIdentityT'] = function($dictMonadST_0 = null) {
  $__num = \func_num_args();
  $__res = $dictMonadST_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Identity_Trans_traversableIdentityT
$GLOBALS['Control_Monad_Identity_Trans_traversableIdentityT'] = function($dictTraversable_0 = null) {
  $__num = \func_num_args();
  $__res = $dictTraversable_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Identity_Trans_runIdentityT
$GLOBALS['Control_Monad_Identity_Trans_runIdentityT'] = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = $v_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Identity_Trans_plusIdentityT
$GLOBALS['Control_Monad_Identity_Trans_plusIdentityT'] = function($dictPlus_0 = null) {
  $__num = \func_num_args();
  $__res = $dictPlus_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Identity_Trans_newtypeIdentityT
$GLOBALS['Control_Monad_Identity_Trans_newtypeIdentityT'] = ["Coercible0" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Control_Monad_Identity_Trans_monadWriterIdentityT
$GLOBALS['Control_Monad_Identity_Trans_monadWriterIdentityT'] = function($dictMonadWriter_0 = null) {
  $__num = \func_num_args();
  $__res = $dictMonadWriter_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Identity_Trans_monadTransIdentityT
$GLOBALS['Control_Monad_Identity_Trans_monadTransIdentityT'] = ["lift" => function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Control_Monad_Identity_Trans_IdentityT'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Control_Monad_Identity_Trans_monadThrowIdentityT
$GLOBALS['Control_Monad_Identity_Trans_monadThrowIdentityT'] = function($dictMonadThrow_0 = null) {
  $__num = \func_num_args();
  $__res = $dictMonadThrow_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Identity_Trans_monadTellIdentityT
$GLOBALS['Control_Monad_Identity_Trans_monadTellIdentityT'] = function($dictMonadTell_0 = null) {
  $__num = \func_num_args();
  $__res = $dictMonadTell_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Identity_Trans_monadStateIdentityT
$GLOBALS['Control_Monad_Identity_Trans_monadStateIdentityT'] = function($dictMonadState_0 = null) {
  $__num = \func_num_args();
  $__res = $dictMonadState_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Identity_Trans_monadRecIdentityT
$GLOBALS['Control_Monad_Identity_Trans_monadRecIdentityT'] = function($dictMonadRec_0 = null) {
  $__num = \func_num_args();
  $__res = $dictMonadRec_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Identity_Trans_monadReaderIdentityT
$GLOBALS['Control_Monad_Identity_Trans_monadReaderIdentityT'] = function($dictMonadReader_0 = null) {
  $__num = \func_num_args();
  $__res = $dictMonadReader_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Identity_Trans_monadPlusIdentityT
$GLOBALS['Control_Monad_Identity_Trans_monadPlusIdentityT'] = function($dictMonadPlus_0 = null) {
  $__num = \func_num_args();
  $__res = $dictMonadPlus_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Identity_Trans_monadIdentityT
$GLOBALS['Control_Monad_Identity_Trans_monadIdentityT'] = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $__res = $dictMonad_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Identity_Trans_monadErrorIdentityT
$GLOBALS['Control_Monad_Identity_Trans_monadErrorIdentityT'] = function($dictMonadError_0 = null) {
  $__num = \func_num_args();
  $__res = $dictMonadError_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Identity_Trans_monadEffectIdentityT
$GLOBALS['Control_Monad_Identity_Trans_monadEffectIdentityT'] = function($dictMonadEffect_0 = null) {
  $__num = \func_num_args();
  $__res = $dictMonadEffect_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Identity_Trans_monadContIdentityT
$GLOBALS['Control_Monad_Identity_Trans_monadContIdentityT'] = function($dictMonadCont_0 = null) {
  $__num = \func_num_args();
  $__res = $dictMonadCont_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Identity_Trans_monadAskIdentityT
$GLOBALS['Control_Monad_Identity_Trans_monadAskIdentityT'] = function($dictMonadAsk_0 = null) {
  $__num = \func_num_args();
  $__res = $dictMonadAsk_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Identity_Trans_mapIdentityT
$GLOBALS['Control_Monad_Identity_Trans_mapIdentityT'] = (function() {
  $__fn = function($f_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($f_0)($v_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Control_Monad_Identity_Trans_functorIdentityT
$GLOBALS['Control_Monad_Identity_Trans_functorIdentityT'] = function($dictFunctor_0 = null) {
  $__num = \func_num_args();
  $__res = $dictFunctor_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Identity_Trans_foldableIdentityT
$GLOBALS['Control_Monad_Identity_Trans_foldableIdentityT'] = function($dictFoldable_0 = null) {
  $__num = \func_num_args();
  $__res = $dictFoldable_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Identity_Trans_extendIdentityI
$GLOBALS['Control_Monad_Identity_Trans_extendIdentityI'] = function($dictExtend_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($dictExtend_0)['Functor0'])(null);
  $__res = ["extend" => (function() use ($dictExtend_0) {
  $__fn = function($f_2 = null, $v_3 = null) use ($dictExtend_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictExtend_0)['extend'])(((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])($f_2))($GLOBALS['Control_Monad_Identity_Trans_IdentityT'])))($v_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_2 = null) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = $__local_var_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Identity_Trans_eqIdentityT
$GLOBALS['Control_Monad_Identity_Trans_eqIdentityT'] = (function() {
  $__fn = function($dictEq1_0 = null, $dictEq_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $eq11_2_0 = (($dictEq1_0)['eq1'])($dictEq_1);
  $__res = ["eq" => (function() use ($eq11_2_0) {
  $__fn = function($x_3 = null, $y_4 = null) use ($eq11_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($eq11_2_0)($x_3))($y_4);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Control_Monad_Identity_Trans_ordIdentityT
$GLOBALS['Control_Monad_Identity_Trans_ordIdentityT'] = function($dictOrd1_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($dictOrd1_0)['Eq10'])(null);
  $__res = function($dictOrd_2 = null) use ($__local_var_1_0, $dictOrd1_0) {
  $__num = \func_num_args();
  $compare11_3_1 = (($dictOrd1_0)['compare1'])($dictOrd_2);
  $eq11_4_2 = (($__local_var_1_0)['eq1'])((($dictOrd_2)['Eq0'])(null));
  $eqIdentityT2_5_3 = ["eq" => (function() use ($eq11_4_2) {
  $__fn = function($x_5 = null, $y_6 = null) use ($eq11_4_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($eq11_4_2)($x_5))($y_6);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $__res = ["compare" => (function() use ($compare11_3_1) {
  $__fn = function($x_6 = null, $y_7 = null) use ($compare11_3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($compare11_3_1)($x_6))($y_7);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($_dollar__unused_6 = null) use ($eqIdentityT2_5_3) {
  $__num = \func_num_args();
  $__res = $eqIdentityT2_5_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Identity_Trans_eq1IdentityT
$GLOBALS['Control_Monad_Identity_Trans_eq1IdentityT'] = function($dictEq1_0 = null) {
  $__num = \func_num_args();
  $__res = ["eq1" => function($dictEq_1 = null) use ($dictEq1_0) {
  $__num = \func_num_args();
  $__res = (($dictEq1_0)['eq1'])($dictEq_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Identity_Trans_ord1IdentityT
$GLOBALS['Control_Monad_Identity_Trans_ord1IdentityT'] = function($dictOrd1_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($dictOrd1_0)['Eq10'])(null);
  $__local_var_2_1 = (($dictOrd1_0)['Eq10'])(null);
  $eq1IdentityT1_3_2 = ["eq1" => function($dictEq_3 = null) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = (($__local_var_2_1)['eq1'])($dictEq_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ["compare1" => function($dictOrd_4 = null) use ($__local_var_1_0, $dictOrd1_0) {
  $__num = \func_num_args();
  $compare11_5_3 = (($dictOrd1_0)['compare1'])($dictOrd_4);
  $eq11_6_4 = (($__local_var_1_0)['eq1'])((($dictOrd_4)['Eq0'])(null));
  $eqIdentityT2_7_5 = ["eq" => (function() use ($eq11_6_4) {
  $__fn = function($x_7 = null, $y_8 = null) use ($eq11_6_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($eq11_6_4)($x_7))($y_8);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $__res = (["compare" => (function() use ($compare11_5_3) {
  $__fn = function($x_8 = null, $y_9 = null) use ($compare11_5_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($compare11_5_3)($x_8))($y_9);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($_dollar__unused_8 = null) use ($eqIdentityT2_7_5) {
  $__num = \func_num_args();
  $__res = $eqIdentityT2_7_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}])['compare'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq10" => function($_dollar__unused_4 = null) use ($eq1IdentityT1_3_2) {
  $__num = \func_num_args();
  $__res = $eq1IdentityT1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Identity_Trans_comonadIdentityT
$GLOBALS['Control_Monad_Identity_Trans_comonadIdentityT'] = function($dictComonad_0 = null) {
  $__num = \func_num_args();
  $extendIdentityI1_1_0 = ($GLOBALS['Control_Monad_Identity_Trans_extendIdentityI'])((($dictComonad_0)['Extend0'])(null));
  $__res = ["extract" => ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])(($dictComonad_0)['extract']))($GLOBALS['Control_Monad_Identity_Trans_runIdentityT']), "Extend0" => function($_dollar__unused_2 = null) use ($extendIdentityI1_1_0) {
  $__num = \func_num_args();
  $__res = $extendIdentityI1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Identity_Trans_bindIdentityT
$GLOBALS['Control_Monad_Identity_Trans_bindIdentityT'] = function($dictBind_0 = null) {
  $__num = \func_num_args();
  $__res = $dictBind_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Identity_Trans_applyIdentityT
$GLOBALS['Control_Monad_Identity_Trans_applyIdentityT'] = function($dictApply_0 = null) {
  $__num = \func_num_args();
  $__res = $dictApply_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Identity_Trans_applicativeIdentityT
$GLOBALS['Control_Monad_Identity_Trans_applicativeIdentityT'] = function($dictApplicative_0 = null) {
  $__num = \func_num_args();
  $__res = $dictApplicative_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Identity_Trans_alternativeIdentityT
$GLOBALS['Control_Monad_Identity_Trans_alternativeIdentityT'] = function($dictAlternative_0 = null) {
  $__num = \func_num_args();
  $__res = $dictAlternative_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Identity_Trans_altIdentityT
$GLOBALS['Control_Monad_Identity_Trans_altIdentityT'] = function($dictAlt_0 = null) {
  $__num = \func_num_args();
  $__res = $dictAlt_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

