<?php

namespace Control\Monad\Reader\Trans;

// ALL IMPORTS: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Monad, Control.Monad.Cont.Class, Control.Monad.Error.Class, Control.Monad.Reader.Class, Control.Monad.Reader.Trans, Control.Monad.Rec.Class, Control.Monad.ST.Class, Control.Monad.State.Class, Control.Monad.Trans.Class, Control.Monad.Writer.Class, Control.MonadPlus, Control.Plus, Control.Semigroupoid, Data.Distributive, Data.Function, Data.Functor, Data.Monoid, Data.Newtype, Data.Semigroup, Effect.Class, Prelude, Prim
// TO REQUIRE: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Monad, Control.Monad.Cont.Class, Control.Monad.Error.Class, Control.Monad.Reader.Class, Control.Monad.Reader.Trans, Control.Monad.Rec.Class, Control.Monad.ST.Class, Control.Monad.State.Class, Control.Monad.Trans.Class, Control.Monad.Writer.Class, Control.MonadPlus, Control.Plus, Control.Semigroupoid, Data.Distributive, Data.Function, Data.Functor, Data.Monoid, Data.Newtype, Data.Semigroup, Effect.Class, Prelude
require_once __DIR__ . '/../Control.Alt/index.php';
require_once __DIR__ . '/../Control.Alternative/index.php';
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Monad/index.php';
require_once __DIR__ . '/../Control.Monad.Cont.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Error.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Reader.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Reader.Trans/index.php';
require_once __DIR__ . '/../Control.Monad.Rec.Class/index.php';
require_once __DIR__ . '/../Control.Monad.ST.Class/index.php';
require_once __DIR__ . '/../Control.Monad.State.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Trans.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Writer.Class/index.php';
require_once __DIR__ . '/../Control.MonadPlus/index.php';
require_once __DIR__ . '/../Control.Plus/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Distributive/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Monoid/index.php';
require_once __DIR__ . '/../Data.Newtype/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
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


// Control_Monad_Reader_Trans_ReaderT
$GLOBALS['Control_Monad_Reader_Trans_ReaderT'] = function($x_0 = null) {
  $__num = \func_num_args();
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Reader_Trans_withReaderT
$GLOBALS['Control_Monad_Reader_Trans_withReaderT'] = (function() {
  $__fn = function($f_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($v_1))($f_0);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Control_Monad_Reader_Trans_runReaderT
$GLOBALS['Control_Monad_Reader_Trans_runReaderT'] = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = $v_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Reader_Trans_newtypeReaderT
$GLOBALS['Control_Monad_Reader_Trans_newtypeReaderT'] = ["Coercible0" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Control_Monad_Reader_Trans_monadTransReaderT
$GLOBALS['Control_Monad_Reader_Trans_monadTransReaderT'] = ["lift" => function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Reader_Trans_ReaderT']))($GLOBALS['Data_Function_const']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Control_Monad_Reader_Trans_mapReaderT
$GLOBALS['Control_Monad_Reader_Trans_mapReaderT'] = (function() {
  $__fn = function($f_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($f_0))($v_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Control_Monad_Reader_Trans_functorReaderT
$GLOBALS['Control_Monad_Reader_Trans_functorReaderT'] = function($dictFunctor_0 = null) {
  $__num = \func_num_args();
  $__res = ["map" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Reader_Trans_mapReaderT']))(($dictFunctor_0)['map'])];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Reader_Trans_distributiveReaderT
$GLOBALS['Control_Monad_Reader_Trans_distributiveReaderT'] = function($dictDistributive_0 = null) {
  $__num = \func_num_args();
  $__tco_var_Control_Monad_Reader_Trans_distributiveReaderT_dictDistributive_0 = $dictDistributive_0;
  tco_loop_Control_Monad_Reader_Trans_distributiveReaderT:;
  $dictDistributive_0 = $__tco_var_Control_Monad_Reader_Trans_distributiveReaderT_dictDistributive_0;
  $functorReaderT1_1_0 = ["map" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Reader_Trans_mapReaderT']))(((($dictDistributive_0)['Functor0'])(null))['map'])];
  $__res = ["distribute" => function($dictFunctor_2 = null) use ($dictDistributive_0) {
  $__num = \func_num_args();
  $collect1_3_1 = (($dictDistributive_0)['collect'])($dictFunctor_2);
  $__res = (function() use ($collect1_3_1) {
  $__fn = function($a_4 = null, $e_5 = null) use ($collect1_3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($collect1_3_1)(function($r_6 = null) use ($e_5) {
  $__num = \func_num_args();
  $__res = ($r_6)($e_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($a_4);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "collect" => (function() use ($dictDistributive_0) {
  $__fn = function($dictFunctor_2 = null, $f_3 = null) use ($dictDistributive_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(((($GLOBALS['Control_Monad_Reader_Trans_distributiveReaderT'])($dictDistributive_0))['distribute'])($dictFunctor_2)))((($dictFunctor_2)['map'])($f_3));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_2 = null) use ($functorReaderT1_1_0) {
  $__num = \func_num_args();
  $__res = $functorReaderT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Reader_Trans_applyReaderT
$GLOBALS['Control_Monad_Reader_Trans_applyReaderT'] = function($dictApply_0 = null) {
  $__num = \func_num_args();
  $functorReaderT1_1_0 = ["map" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Reader_Trans_mapReaderT']))(((($dictApply_0)['Functor0'])(null))['map'])];
  $__res = ["apply" => (function() use ($dictApply_0) {
  $__fn = function($v_2 = null, $v1_3 = null, $r_4 = null) use ($dictApply_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictApply_0)['apply'])(($v_2)($r_4)))(($v1_3)($r_4));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_2 = null) use ($functorReaderT1_1_0) {
  $__num = \func_num_args();
  $__res = $functorReaderT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Reader_Trans_bindReaderT
$GLOBALS['Control_Monad_Reader_Trans_bindReaderT'] = function($dictBind_0 = null) {
  $__num = \func_num_args();
  $applyReaderT1_1_0 = ($GLOBALS['Control_Monad_Reader_Trans_applyReaderT'])((($dictBind_0)['Apply0'])(null));
  $__res = ["bind" => (function() use ($dictBind_0) {
  $__fn = function($v_2 = null, $k_3 = null, $r_4 = null) use ($dictBind_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictBind_0)['bind'])(($v_2)($r_4)))(function($a_5 = null) use ($k_3, $r_4) {
  $__num = \func_num_args();
  $__res = (($k_3)($a_5))($r_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "Apply0" => function($_dollar__unused_2 = null) use ($applyReaderT1_1_0) {
  $__num = \func_num_args();
  $__res = $applyReaderT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Reader_Trans_semigroupReaderT
$GLOBALS['Control_Monad_Reader_Trans_semigroupReaderT'] = function($dictApply_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = ($GLOBALS['Control_Monad_Reader_Trans_applyReaderT'])($dictApply_0);
  $__res = function($dictSemigroup_2 = null) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__local_var_3_1 = ($dictSemigroup_2)['append'];
  $__res = ["append" => (function() use ($__local_var_1_0, $__local_var_3_1) {
  $__fn = function($a_4 = null, $b_5 = null) use ($__local_var_1_0, $__local_var_3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_1_0)['apply'])(((((($__local_var_1_0)['Functor0'])(null))['map'])($__local_var_3_1))($a_4)))($b_5);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Reader_Trans_applicativeReaderT
$GLOBALS['Control_Monad_Reader_Trans_applicativeReaderT'] = function($dictApplicative_0 = null) {
  $__num = \func_num_args();
  $applyReaderT1_1_0 = ($GLOBALS['Control_Monad_Reader_Trans_applyReaderT'])((($dictApplicative_0)['Apply0'])(null));
  $__res = ["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Reader_Trans_ReaderT']))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Function_const']))(($dictApplicative_0)['pure'])), "Apply0" => function($_dollar__unused_2 = null) use ($applyReaderT1_1_0) {
  $__num = \func_num_args();
  $__res = $applyReaderT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Reader_Trans_monadReaderT
$GLOBALS['Control_Monad_Reader_Trans_monadReaderT'] = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $applicativeReaderT1_1_0 = ($GLOBALS['Control_Monad_Reader_Trans_applicativeReaderT'])((($dictMonad_0)['Applicative0'])(null));
  $bindReaderT1_2_1 = ($GLOBALS['Control_Monad_Reader_Trans_bindReaderT'])((($dictMonad_0)['Bind1'])(null));
  $__res = ["Applicative0" => function($_dollar__unused_3 = null) use ($applicativeReaderT1_1_0) {
  $__num = \func_num_args();
  $__res = $applicativeReaderT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar__unused_3 = null) use ($bindReaderT1_2_1) {
  $__num = \func_num_args();
  $__res = $bindReaderT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Reader_Trans_monadAskReaderT
$GLOBALS['Control_Monad_Reader_Trans_monadAskReaderT'] = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $monadReaderT1_1_0 = ($GLOBALS['Control_Monad_Reader_Trans_monadReaderT'])($dictMonad_0);
  $__res = ["ask" => ((($dictMonad_0)['Applicative0'])(null))['pure'], "Monad0" => function($_dollar__unused_2 = null) use ($monadReaderT1_1_0) {
  $__num = \func_num_args();
  $__res = $monadReaderT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Reader_Trans_monadReaderReaderT
$GLOBALS['Control_Monad_Reader_Trans_monadReaderReaderT'] = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $monadReaderT1_1_0 = ($GLOBALS['Control_Monad_Reader_Trans_monadReaderT'])($dictMonad_0);
  $monadAskReaderT1_2_1 = ["ask" => ((($dictMonad_0)['Applicative0'])(null))['pure'], "Monad0" => function($_dollar__unused_2 = null) use ($monadReaderT1_1_0) {
  $__num = \func_num_args();
  $__res = $monadReaderT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ["local" => $GLOBALS['Control_Monad_Reader_Trans_withReaderT'], "MonadAsk0" => function($_dollar__unused_3 = null) use ($monadAskReaderT1_2_1) {
  $__num = \func_num_args();
  $__res = $monadAskReaderT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Reader_Trans_monadContReaderT
$GLOBALS['Control_Monad_Reader_Trans_monadContReaderT'] = function($dictMonadCont_0 = null) {
  $__num = \func_num_args();
  $monadReaderT1_1_0 = ($GLOBALS['Control_Monad_Reader_Trans_monadReaderT'])((($dictMonadCont_0)['Monad0'])(null));
  $__res = ["callCC" => (function() use ($dictMonadCont_0) {
  $__fn = function($f_2 = null, $r_3 = null) use ($dictMonadCont_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($dictMonadCont_0)['callCC'])(function($c_4 = null) use ($f_2, $r_3) {
  $__num = \func_num_args();
  $__res = (($f_2)((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Reader_Trans_ReaderT']))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Function_const']))($c_4))))($r_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Monad0" => function($_dollar__unused_2 = null) use ($monadReaderT1_1_0) {
  $__num = \func_num_args();
  $__res = $monadReaderT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Reader_Trans_monadEffectReader
$GLOBALS['Control_Monad_Reader_Trans_monadEffectReader'] = function($dictMonadEffect_0 = null) {
  $__num = \func_num_args();
  $Monad0_1_0 = (($dictMonadEffect_0)['Monad0'])(null);
  $monadReaderT1_2_1 = ($GLOBALS['Control_Monad_Reader_Trans_monadReaderT'])($Monad0_1_0);
  $__res = ["liftEffect" => (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Control_Monad_Reader_Trans_monadTransReaderT'])['lift'])($Monad0_1_0)))(($dictMonadEffect_0)['liftEffect']), "Monad0" => function($_dollar__unused_3 = null) use ($monadReaderT1_2_1) {
  $__num = \func_num_args();
  $__res = $monadReaderT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Reader_Trans_monadRecReaderT
$GLOBALS['Control_Monad_Reader_Trans_monadRecReaderT'] = function($dictMonadRec_0 = null) {
  $__num = \func_num_args();
  $Monad0_1_0 = (($dictMonadRec_0)['Monad0'])(null);
  $__local_var_2_1 = (($Monad0_1_0)['Bind1'])(null);
  $pure_3_2 = ((($Monad0_1_0)['Applicative0'])(null))['pure'];
  $monadReaderT1_4_3 = ($GLOBALS['Control_Monad_Reader_Trans_monadReaderT'])($Monad0_1_0);
  $__res = ["tailRecM" => (function() use ($__local_var_2_1, $dictMonadRec_0, $pure_3_2) {
  $__fn = function($k_5 = null, $a_6 = null, $r_7 = null) use ($__local_var_2_1, $dictMonadRec_0, $pure_3_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictMonadRec_0)['tailRecM'])(function($a_prime_8 = null) use ($__local_var_2_1, $k_5, $pure_3_2, $r_7) {
  $__num = \func_num_args();
  $__res = ((($__local_var_2_1)['bind'])((($k_5)($a_prime_8))($r_7)))($pure_3_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($a_6);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "Monad0" => function($_dollar__unused_5 = null) use ($monadReaderT1_4_3) {
  $__num = \func_num_args();
  $__res = $monadReaderT1_4_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Reader_Trans_monadStateReaderT
$GLOBALS['Control_Monad_Reader_Trans_monadStateReaderT'] = function($dictMonadState_0 = null) {
  $__num = \func_num_args();
  $Monad0_1_0 = (($dictMonadState_0)['Monad0'])(null);
  $monadReaderT1_2_1 = ($GLOBALS['Control_Monad_Reader_Trans_monadReaderT'])($Monad0_1_0);
  $__res = ["state" => (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Control_Monad_Reader_Trans_monadTransReaderT'])['lift'])($Monad0_1_0)))(($dictMonadState_0)['state']), "Monad0" => function($_dollar__unused_3 = null) use ($monadReaderT1_2_1) {
  $__num = \func_num_args();
  $__res = $monadReaderT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Reader_Trans_monadTellReaderT
$GLOBALS['Control_Monad_Reader_Trans_monadTellReaderT'] = function($dictMonadTell_0 = null) {
  $__num = \func_num_args();
  $Monad1_1_0 = (($dictMonadTell_0)['Monad1'])(null);
  $Semigroup0_2_1 = (($dictMonadTell_0)['Semigroup0'])(null);
  $monadReaderT1_3_2 = ($GLOBALS['Control_Monad_Reader_Trans_monadReaderT'])($Monad1_1_0);
  $__res = ["tell" => (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Control_Monad_Reader_Trans_monadTransReaderT'])['lift'])($Monad1_1_0)))(($dictMonadTell_0)['tell']), "Semigroup0" => function($_dollar__unused_4 = null) use ($Semigroup0_2_1) {
  $__num = \func_num_args();
  $__res = $Semigroup0_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad1" => function($_dollar__unused_4 = null) use ($monadReaderT1_3_2) {
  $__num = \func_num_args();
  $__res = $monadReaderT1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Reader_Trans_monadWriterReaderT
$GLOBALS['Control_Monad_Reader_Trans_monadWriterReaderT'] = function($dictMonadWriter_0 = null) {
  $__num = \func_num_args();
  $Monoid0_1_0 = (($dictMonadWriter_0)['Monoid0'])(null);
  $monadTellReaderT1_2_1 = ($GLOBALS['Control_Monad_Reader_Trans_monadTellReaderT'])((($dictMonadWriter_0)['MonadTell1'])(null));
  $__res = ["listen" => ($GLOBALS['Control_Monad_Reader_Trans_mapReaderT'])(($dictMonadWriter_0)['listen']), "pass" => ($GLOBALS['Control_Monad_Reader_Trans_mapReaderT'])(($dictMonadWriter_0)['pass']), "Monoid0" => function($_dollar__unused_3 = null) use ($Monoid0_1_0) {
  $__num = \func_num_args();
  $__res = $Monoid0_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "MonadTell1" => function($_dollar__unused_3 = null) use ($monadTellReaderT1_2_1) {
  $__num = \func_num_args();
  $__res = $monadTellReaderT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Reader_Trans_monadThrowReaderT
$GLOBALS['Control_Monad_Reader_Trans_monadThrowReaderT'] = function($dictMonadThrow_0 = null) {
  $__num = \func_num_args();
  $Monad0_1_0 = (($dictMonadThrow_0)['Monad0'])(null);
  $monadReaderT1_2_1 = ($GLOBALS['Control_Monad_Reader_Trans_monadReaderT'])($Monad0_1_0);
  $__res = ["throwError" => (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Control_Monad_Reader_Trans_monadTransReaderT'])['lift'])($Monad0_1_0)))(($dictMonadThrow_0)['throwError']), "Monad0" => function($_dollar__unused_3 = null) use ($monadReaderT1_2_1) {
  $__num = \func_num_args();
  $__res = $monadReaderT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Reader_Trans_monadErrorReaderT
$GLOBALS['Control_Monad_Reader_Trans_monadErrorReaderT'] = function($dictMonadError_0 = null) {
  $__num = \func_num_args();
  $monadThrowReaderT1_1_0 = ($GLOBALS['Control_Monad_Reader_Trans_monadThrowReaderT'])((($dictMonadError_0)['MonadThrow0'])(null));
  $__res = ["catchError" => (function() use ($dictMonadError_0) {
  $__fn = function($v_2 = null, $h_3 = null, $r_4 = null) use ($dictMonadError_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictMonadError_0)['catchError'])(($v_2)($r_4)))(function($e_5 = null) use ($h_3, $r_4) {
  $__num = \func_num_args();
  $__res = (($h_3)($e_5))($r_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "MonadThrow0" => function($_dollar__unused_2 = null) use ($monadThrowReaderT1_1_0) {
  $__num = \func_num_args();
  $__res = $monadThrowReaderT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Reader_Trans_monadSTReaderT
$GLOBALS['Control_Monad_Reader_Trans_monadSTReaderT'] = function($dictMonadST_0 = null) {
  $__num = \func_num_args();
  $Monad0_1_0 = (($dictMonadST_0)['Monad0'])(null);
  $monadReaderT1_2_1 = ($GLOBALS['Control_Monad_Reader_Trans_monadReaderT'])($Monad0_1_0);
  $__res = ["liftST" => (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Control_Monad_Reader_Trans_monadTransReaderT'])['lift'])($Monad0_1_0)))(($dictMonadST_0)['liftST']), "Monad0" => function($_dollar__unused_3 = null) use ($monadReaderT1_2_1) {
  $__num = \func_num_args();
  $__res = $monadReaderT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Reader_Trans_monoidReaderT
$GLOBALS['Control_Monad_Reader_Trans_monoidReaderT'] = function($dictApplicative_0 = null) {
  $__num = \func_num_args();
  $semigroupReaderT1_1_0 = ($GLOBALS['Control_Monad_Reader_Trans_semigroupReaderT'])((($dictApplicative_0)['Apply0'])(null));
  $__res = function($dictMonoid_2 = null) use ($dictApplicative_0, $semigroupReaderT1_1_0) {
  $__num = \func_num_args();
  $semigroupReaderT2_3_1 = ($semigroupReaderT1_1_0)((($dictMonoid_2)['Semigroup0'])(null));
  $__res = ["mempty" => ((($GLOBALS['Control_Monad_Reader_Trans_applicativeReaderT'])($dictApplicative_0))['pure'])(($dictMonoid_2)['mempty']), "Semigroup0" => function($_dollar__unused_4 = null) use ($semigroupReaderT2_3_1) {
  $__num = \func_num_args();
  $__res = $semigroupReaderT2_3_1;
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

// Control_Monad_Reader_Trans_altReaderT
$GLOBALS['Control_Monad_Reader_Trans_altReaderT'] = function($dictAlt_0 = null) {
  $__num = \func_num_args();
  $functorReaderT1_1_0 = ["map" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Reader_Trans_mapReaderT']))(((($dictAlt_0)['Functor0'])(null))['map'])];
  $__res = ["alt" => (function() use ($dictAlt_0) {
  $__fn = function($v_2 = null, $v1_3 = null, $r_4 = null) use ($dictAlt_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictAlt_0)['alt'])(($v_2)($r_4)))(($v1_3)($r_4));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_2 = null) use ($functorReaderT1_1_0) {
  $__num = \func_num_args();
  $__res = $functorReaderT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Reader_Trans_plusReaderT
$GLOBALS['Control_Monad_Reader_Trans_plusReaderT'] = function($dictPlus_0 = null) {
  $__num = \func_num_args();
  $altReaderT1_1_0 = ($GLOBALS['Control_Monad_Reader_Trans_altReaderT'])((($dictPlus_0)['Alt0'])(null));
  $__local_var_2_1 = ($dictPlus_0)['empty'];
  $__res = ["empty" => function($v_3 = null) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = $__local_var_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Alt0" => function($_dollar__unused_2 = null) use ($altReaderT1_1_0) {
  $__num = \func_num_args();
  $__res = $altReaderT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Reader_Trans_alternativeReaderT
$GLOBALS['Control_Monad_Reader_Trans_alternativeReaderT'] = function($dictAlternative_0 = null) {
  $__num = \func_num_args();
  $applicativeReaderT1_1_0 = ($GLOBALS['Control_Monad_Reader_Trans_applicativeReaderT'])((($dictAlternative_0)['Applicative0'])(null));
  $plusReaderT1_2_1 = ($GLOBALS['Control_Monad_Reader_Trans_plusReaderT'])((($dictAlternative_0)['Plus1'])(null));
  $__res = ["Applicative0" => function($_dollar__unused_3 = null) use ($applicativeReaderT1_1_0) {
  $__num = \func_num_args();
  $__res = $applicativeReaderT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Plus1" => function($_dollar__unused_3 = null) use ($plusReaderT1_2_1) {
  $__num = \func_num_args();
  $__res = $plusReaderT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Reader_Trans_monadPlusReaderT
$GLOBALS['Control_Monad_Reader_Trans_monadPlusReaderT'] = function($dictMonadPlus_0 = null) {
  $__num = \func_num_args();
  $monadReaderT1_1_0 = ($GLOBALS['Control_Monad_Reader_Trans_monadReaderT'])((($dictMonadPlus_0)['Monad0'])(null));
  $alternativeReaderT1_2_1 = ($GLOBALS['Control_Monad_Reader_Trans_alternativeReaderT'])((($dictMonadPlus_0)['Alternative1'])(null));
  $__res = ["Monad0" => function($_dollar__unused_3 = null) use ($monadReaderT1_1_0) {
  $__num = \func_num_args();
  $__res = $monadReaderT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Alternative1" => function($_dollar__unused_3 = null) use ($alternativeReaderT1_2_1) {
  $__num = \func_num_args();
  $__res = $alternativeReaderT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

