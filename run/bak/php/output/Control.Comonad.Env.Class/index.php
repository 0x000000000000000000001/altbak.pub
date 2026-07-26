<?php

namespace Control\Comonad\Env\Class;

// ALL IMPORTS: Control.Comonad, Control.Comonad.Env.Class, Control.Comonad.Env.Trans, Control.Comonad.Store, Control.Comonad.Store.Trans, Control.Comonad.Traced.Trans, Control.Comonad.Trans.Class, Control.Semigroupoid, Data.Tuple, Prelude, Prim
// TO REQUIRE: Control.Comonad, Control.Comonad.Env.Class, Control.Comonad.Env.Trans, Control.Comonad.Store, Control.Comonad.Store.Trans, Control.Comonad.Traced.Trans, Control.Comonad.Trans.Class, Control.Semigroupoid, Data.Tuple, Prelude
require_once __DIR__ . '/../Control.Comonad/index.php';
require_once __DIR__ . '/../Control.Comonad.Env.Class/index.php';
require_once __DIR__ . '/../Control.Comonad.Env.Trans/index.php';
require_once __DIR__ . '/../Control.Comonad.Store/index.php';
require_once __DIR__ . '/../Control.Comonad.Store.Trans/index.php';
require_once __DIR__ . '/../Control.Comonad.Traced.Trans/index.php';
require_once __DIR__ . '/../Control.Comonad.Trans.Class/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Tuple/index.php';
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


// Control_Comonad_Env_Class_local
$GLOBALS['Control_Comonad_Env_Class_local'] = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)['local'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Comonad_Env_Class_comonadAskTuple
$GLOBALS['Control_Comonad_Env_Class_comonadAskTuple'] = ["ask" => $GLOBALS['Data_Tuple_fst'], "Comonad0" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Tuple_comonadTuple'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Control_Comonad_Env_Class_comonadEnvTuple
$GLOBALS['Control_Comonad_Env_Class_comonadEnvTuple'] = ["local" => (function() {
  $__fn = function($f_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("Tuple", ($f_0)(($v_1)->{'value0'}), ($v_1)->{'value1'});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "ComonadAsk0" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Control_Comonad_Env_Class_comonadAskTuple'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Control_Comonad_Env_Class_comonadAskEnvT
$GLOBALS['Control_Comonad_Env_Class_comonadAskEnvT'] = function($dictComonad_0 = null) {
  $__num = \func_num_args();
  $comonadEnvT_1_0 = ($GLOBALS['Control_Comonad_Env_Trans_comonadEnvT'])($dictComonad_0);
  $__res = ["ask" => function($v_2 = null) {
  $__num = \func_num_args();
  $__res = ($v_2)->{'value0'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Comonad0" => function($_dollar__unused_2 = null) use ($comonadEnvT_1_0) {
  $__num = \func_num_args();
  $__res = $comonadEnvT_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Comonad_Env_Class_comonadEnvEnvT
$GLOBALS['Control_Comonad_Env_Class_comonadEnvEnvT'] = function($dictComonad_0 = null) {
  $__num = \func_num_args();
  $comonadEnvT_1_0 = ($GLOBALS['Control_Comonad_Env_Trans_comonadEnvT'])($dictComonad_0);
  $__res = ["local" => (function() {
  $__fn = function($f_2 = null, $v_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("Tuple", ($f_2)(($v_3)->{'value0'}), ($v_3)->{'value1'});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "ComonadAsk0" => function($_dollar__unused_2 = null) use ($comonadEnvT_1_0) {
  $__num = \func_num_args();
  $__res = ["ask" => function($v_3 = null) {
  $__num = \func_num_args();
  $__res = ($v_3)->{'value0'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Comonad0" => function($_dollar__unused_3 = null) use ($comonadEnvT_1_0) {
  $__num = \func_num_args();
  $__res = $comonadEnvT_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Comonad_Env_Class_ask
$GLOBALS['Control_Comonad_Env_Class_ask'] = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)['ask'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Comonad_Env_Class_asks
$GLOBALS['Control_Comonad_Env_Class_asks'] = (function() {
  $__fn = function($dictComonadAsk_0 = null, $f_1 = null, $x_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($f_1)((($dictComonadAsk_0)['ask'])($x_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Control_Comonad_Env_Class_comonadAskStoreT
$GLOBALS['Control_Comonad_Env_Class_comonadAskStoreT'] = function($dictComonadAsk_0 = null) {
  $__num = \func_num_args();
  $Comonad0_1_0 = (($dictComonadAsk_0)['Comonad0'])(null);
  $comonadStoreT_2_1 = ($GLOBALS['Control_Comonad_Store_Trans_comonadStoreT'])($Comonad0_1_0);
  $__res = ["ask" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(($dictComonadAsk_0)['ask']))((($GLOBALS['Control_Comonad_Store_Trans_comonadTransStoreT'])['lower'])($Comonad0_1_0)), "Comonad0" => function($_dollar__unused_3 = null) use ($comonadStoreT_2_1) {
  $__num = \func_num_args();
  $__res = $comonadStoreT_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Comonad_Env_Class_comonadEnvStoreT
$GLOBALS['Control_Comonad_Env_Class_comonadEnvStoreT'] = function($dictComonadEnv_0 = null) {
  $__num = \func_num_args();
  $comonadAskStoreT1_1_0 = ($GLOBALS['Control_Comonad_Env_Class_comonadAskStoreT'])((($dictComonadEnv_0)['ComonadAsk0'])(null));
  $__res = ["local" => (function() use ($dictComonadEnv_0) {
  $__fn = function($f_2 = null, $v_3 = null) use ($dictComonadEnv_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("Tuple", ((($dictComonadEnv_0)['local'])($f_2))(($v_3)->{'value0'}), ($v_3)->{'value1'});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "ComonadAsk0" => function($_dollar__unused_2 = null) use ($comonadAskStoreT1_1_0) {
  $__num = \func_num_args();
  $__res = $comonadAskStoreT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Comonad_Env_Class_comonadAskTracedT
$GLOBALS['Control_Comonad_Env_Class_comonadAskTracedT'] = function($dictComonadAsk_0 = null) {
  $__num = \func_num_args();
  $ask1_1_0 = ($dictComonadAsk_0)['ask'];
  $Comonad0_2_1 = (($dictComonadAsk_0)['Comonad0'])(null);
  $comonadTracedT_3_2 = ($GLOBALS['Control_Comonad_Traced_Trans_comonadTracedT'])($Comonad0_2_1);
  $__res = function($dictMonoid_4 = null) use ($Comonad0_2_1, $ask1_1_0, $comonadTracedT_3_2) {
  $__num = \func_num_args();
  $comonadTracedT1_5_3 = ($comonadTracedT_3_2)($dictMonoid_4);
  $mempty_6_4 = ($dictMonoid_4)['mempty'];
  $__res = ["ask" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($ask1_1_0))(function($v_7 = null) use ($Comonad0_2_1, $mempty_6_4) {
  $__num = \func_num_args();
  $__res = ((((((($Comonad0_2_1)['Extend0'])(null))['Functor0'])(null))['map'])(function($f_8 = null) use ($mempty_6_4) {
  $__num = \func_num_args();
  $__res = ($f_8)($mempty_6_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), "Comonad0" => function($_dollar__unused_6 = null) use ($comonadTracedT1_5_3) {
  $__num = \func_num_args();
  $__res = $comonadTracedT1_5_3;
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

// Control_Comonad_Env_Class_comonadEnvTracedT
$GLOBALS['Control_Comonad_Env_Class_comonadEnvTracedT'] = function($dictComonadEnv_0 = null) {
  $__num = \func_num_args();
  $comonadAskTracedT1_1_0 = ($GLOBALS['Control_Comonad_Env_Class_comonadAskTracedT'])((($dictComonadEnv_0)['ComonadAsk0'])(null));
  $__res = function($dictMonoid_2 = null) use ($comonadAskTracedT1_1_0, $dictComonadEnv_0) {
  $__num = \func_num_args();
  $comonadAskTracedT2_3_1 = ($comonadAskTracedT1_1_0)($dictMonoid_2);
  $__res = ["local" => (function() use ($dictComonadEnv_0) {
  $__fn = function($f_4 = null, $v_5 = null) use ($dictComonadEnv_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictComonadEnv_0)['local'])($f_4))($v_5);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "ComonadAsk0" => function($_dollar__unused_4 = null) use ($comonadAskTracedT2_3_1) {
  $__num = \func_num_args();
  $__res = $comonadAskTracedT2_3_1;
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

