<?php

namespace Control\Monad\Cont\Trans;

// ALL IMPORTS: Control.Applicative, Control.Apply, Control.Bind, Control.Monad, Control.Monad.Cont.Class, Control.Monad.Cont.Trans, Control.Monad.Reader.Class, Control.Monad.ST.Class, Control.Monad.State.Class, Control.Monad.Trans.Class, Control.Semigroupoid, Data.Function, Data.Functor, Data.Monoid, Data.Newtype, Data.Semigroup, Effect.Class, Prelude, Prim
// TO REQUIRE: Control.Applicative, Control.Apply, Control.Bind, Control.Monad, Control.Monad.Cont.Class, Control.Monad.Cont.Trans, Control.Monad.Reader.Class, Control.Monad.ST.Class, Control.Monad.State.Class, Control.Monad.Trans.Class, Control.Semigroupoid, Data.Function, Data.Functor, Data.Monoid, Data.Newtype, Data.Semigroup, Effect.Class, Prelude
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Monad/index.php';
require_once __DIR__ . '/../Control.Monad.Cont.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Cont.Trans/index.php';
require_once __DIR__ . '/../Control.Monad.Reader.Class/index.php';
require_once __DIR__ . '/../Control.Monad.ST.Class/index.php';
require_once __DIR__ . '/../Control.Monad.State.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Trans.Class/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
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


// Control_Monad_Cont_Trans_ContT
$GLOBALS['Control_Monad_Cont_Trans_ContT'] = function($x_0 = null) {
  $__num = \func_num_args();
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Cont_Trans_withContT
$GLOBALS['Control_Monad_Cont_Trans_withContT'] = (function() {
  $__fn = function($f_0 = null, $v_1 = null, $k_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($v_1)(($f_0)($k_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Control_Monad_Cont_Trans_runContT
$GLOBALS['Control_Monad_Cont_Trans_runContT'] = (function() {
  $__fn = function($v_0 = null, $k_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($v_0)($k_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Control_Monad_Cont_Trans_newtypeContT
$GLOBALS['Control_Monad_Cont_Trans_newtypeContT'] = ["Coercible0" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Control_Monad_Cont_Trans_monadTransContT
$GLOBALS['Control_Monad_Cont_Trans_monadTransContT'] = ["lift" => (function() {
  $__fn = function($dictMonad_0 = null, $m_1 = null, $k_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((($dictMonad_0)['Bind1'])(null))['bind'])($m_1))($k_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()];

// Control_Monad_Cont_Trans_mapContT
$GLOBALS['Control_Monad_Cont_Trans_mapContT'] = (function() {
  $__fn = function($f_0 = null, $v_1 = null, $k_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($f_0)(($v_1)($k_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Control_Monad_Cont_Trans_functorContT
$GLOBALS['Control_Monad_Cont_Trans_functorContT'] = function($dictFunctor_0 = null) {
  $__num = \func_num_args();
  $__res = ["map" => (function() {
  $__fn = function($f_1 = null, $v_2 = null, $k_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($v_2)(function($a_4 = null) use ($f_1, $k_3) {
  $__num = \func_num_args();
  $__res = ($k_3)(($f_1)($a_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Cont_Trans_applyContT
$GLOBALS['Control_Monad_Cont_Trans_applyContT'] = function($dictApply_0 = null) {
  $__num = \func_num_args();
  $functorContT1_1_0 = ["map" => (function() {
  $__fn = function($f_1 = null, $v_2 = null, $k_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($v_2)(function($a_4 = null) use ($f_1, $k_3) {
  $__num = \func_num_args();
  $__res = ($k_3)(($f_1)($a_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()];
  $__res = ["apply" => (function() {
  $__fn = function($v_2 = null, $v1_3 = null, $k_4 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($v_2)(function($g_5 = null) use ($k_4, $v1_3) {
  $__num = \func_num_args();
  $__res = ($v1_3)(function($a_6 = null) use ($g_5, $k_4) {
  $__num = \func_num_args();
  $__res = ($k_4)(($g_5)($a_6));
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
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_2 = null) use ($functorContT1_1_0) {
  $__num = \func_num_args();
  $__res = $functorContT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Cont_Trans_bindContT
$GLOBALS['Control_Monad_Cont_Trans_bindContT'] = function($dictBind_0 = null) {
  $__num = \func_num_args();
  $functorContT1_1_0 = ["map" => (function() {
  $__fn = function($f_1 = null, $v_2 = null, $k_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($v_2)(function($a_4 = null) use ($f_1, $k_3) {
  $__num = \func_num_args();
  $__res = ($k_3)(($f_1)($a_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()];
  $applyContT1_2_1 = ["apply" => (function() {
  $__fn = function($v_2 = null, $v1_3 = null, $k_4 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($v_2)(function($g_5 = null) use ($k_4, $v1_3) {
  $__num = \func_num_args();
  $__res = ($v1_3)(function($a_6 = null) use ($g_5, $k_4) {
  $__num = \func_num_args();
  $__res = ($k_4)(($g_5)($a_6));
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
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_2 = null) use ($functorContT1_1_0) {
  $__num = \func_num_args();
  $__res = $functorContT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ["bind" => (function() {
  $__fn = function($v_3 = null, $k_4 = null, $k_prime_5 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($v_3)(function($a_6 = null) use ($k_4, $k_prime_5) {
  $__num = \func_num_args();
  $__res = (($k_4)($a_6))($k_prime_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "Apply0" => function($_dollar__unused_3 = null) use ($applyContT1_2_1) {
  $__num = \func_num_args();
  $__res = $applyContT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Cont_Trans_semigroupContT
$GLOBALS['Control_Monad_Cont_Trans_semigroupContT'] = (function() {
  $__fn = function($dictApply_0 = null, $dictSemigroup_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ["append" => (function() use ($dictSemigroup_1) {
  $__fn = function($a_2 = null, $b_3 = null, $k_4 = null) use ($dictSemigroup_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($a_2)(function($a_5 = null) use ($b_3, $dictSemigroup_1, $k_4) {
  $__num = \func_num_args();
  $__local_var_6_0 = (($dictSemigroup_1)['append'])($a_5);
  $__res = ($b_3)(function($a_7 = null) use ($__local_var_6_0, $k_4) {
  $__num = \func_num_args();
  $__res = ($k_4)(($__local_var_6_0)($a_7));
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
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Control_Monad_Cont_Trans_applicativeContT
$GLOBALS['Control_Monad_Cont_Trans_applicativeContT'] = function($dictApplicative_0 = null) {
  $__num = \func_num_args();
  $functorContT1_1_0 = ["map" => (function() {
  $__fn = function($f_1 = null, $v_2 = null, $k_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($v_2)(function($a_4 = null) use ($f_1, $k_3) {
  $__num = \func_num_args();
  $__res = ($k_3)(($f_1)($a_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()];
  $applyContT1_2_1 = ["apply" => (function() {
  $__fn = function($v_2 = null, $v1_3 = null, $k_4 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($v_2)(function($g_5 = null) use ($k_4, $v1_3) {
  $__num = \func_num_args();
  $__res = ($v1_3)(function($a_6 = null) use ($g_5, $k_4) {
  $__num = \func_num_args();
  $__res = ($k_4)(($g_5)($a_6));
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
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_2 = null) use ($functorContT1_1_0) {
  $__num = \func_num_args();
  $__res = $functorContT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ["pure" => (function() {
  $__fn = function($a_3 = null, $k_4 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($k_4)($a_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Apply0" => function($_dollar__unused_3 = null) use ($applyContT1_2_1) {
  $__num = \func_num_args();
  $__res = $applyContT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Cont_Trans_monadContT
$GLOBALS['Control_Monad_Cont_Trans_monadContT'] = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $functorContT1_1_0 = ["map" => (function() {
  $__fn = function($f_1 = null, $v_2 = null, $k_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($v_2)(function($a_4 = null) use ($f_1, $k_3) {
  $__num = \func_num_args();
  $__res = ($k_3)(($f_1)($a_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()];
  $applyContT1_2_1 = ["apply" => (function() {
  $__fn = function($v_2 = null, $v1_3 = null, $k_4 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($v_2)(function($g_5 = null) use ($k_4, $v1_3) {
  $__num = \func_num_args();
  $__res = ($v1_3)(function($a_6 = null) use ($g_5, $k_4) {
  $__num = \func_num_args();
  $__res = ($k_4)(($g_5)($a_6));
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
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_2 = null) use ($functorContT1_1_0) {
  $__num = \func_num_args();
  $__res = $functorContT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeContT1_2_1 = ["pure" => (function() {
  $__fn = function($a_3 = null, $k_4 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($k_4)($a_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Apply0" => function($_dollar__unused_3 = null) use ($applyContT1_2_1) {
  $__num = \func_num_args();
  $__res = $applyContT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $functorContT1_3_3 = ["map" => (function() {
  $__fn = function($f_3 = null, $v_4 = null, $k_5 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($v_4)(function($a_6 = null) use ($f_3, $k_5) {
  $__num = \func_num_args();
  $__res = ($k_5)(($f_3)($a_6));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()];
  $applyContT1_4_4 = ["apply" => (function() {
  $__fn = function($v_4 = null, $v1_5 = null, $k_6 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($v_4)(function($g_7 = null) use ($k_6, $v1_5) {
  $__num = \func_num_args();
  $__res = ($v1_5)(function($a_8 = null) use ($g_7, $k_6) {
  $__num = \func_num_args();
  $__res = ($k_6)(($g_7)($a_8));
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
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_4 = null) use ($functorContT1_3_3) {
  $__num = \func_num_args();
  $__res = $functorContT1_3_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $bindContT1_4_4 = ["bind" => (function() {
  $__fn = function($v_5 = null, $k_6 = null, $k_prime_7 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($v_5)(function($a_8 = null) use ($k_6, $k_prime_7) {
  $__num = \func_num_args();
  $__res = (($k_6)($a_8))($k_prime_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "Apply0" => function($_dollar__unused_5 = null) use ($applyContT1_4_4) {
  $__num = \func_num_args();
  $__res = $applyContT1_4_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ["Applicative0" => function($_dollar__unused_5 = null) use ($applicativeContT1_2_1) {
  $__num = \func_num_args();
  $__res = $applicativeContT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar__unused_5 = null) use ($bindContT1_4_4) {
  $__num = \func_num_args();
  $__res = $bindContT1_4_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Cont_Trans_monadAskContT
$GLOBALS['Control_Monad_Cont_Trans_monadAskContT'] = function($dictMonadAsk_0 = null) {
  $__num = \func_num_args();
  $Monad0_1_0 = (($dictMonadAsk_0)['Monad0'])(null);
  $monadContT1_2_1 = ($GLOBALS['Control_Monad_Cont_Trans_monadContT'])($Monad0_1_0);
  $__res = ["ask" => ((($GLOBALS['Control_Monad_Cont_Trans_monadTransContT'])['lift'])($Monad0_1_0))(($dictMonadAsk_0)['ask']), "Monad0" => function($_dollar__unused_3 = null) use ($monadContT1_2_1) {
  $__num = \func_num_args();
  $__res = $monadContT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Cont_Trans_monadReaderContT
$GLOBALS['Control_Monad_Cont_Trans_monadReaderContT'] = function($dictMonadReader_0 = null) {
  $__num = \func_num_args();
  $MonadAsk0_1_0 = (($dictMonadReader_0)['MonadAsk0'])(null);
  $ask_2_1 = ($MonadAsk0_1_0)['ask'];
  $monadAskContT1_3_2 = ($GLOBALS['Control_Monad_Cont_Trans_monadAskContT'])($MonadAsk0_1_0);
  $__res = ["local" => (function() use ($MonadAsk0_1_0, $ask_2_1, $dictMonadReader_0) {
  $__fn = function($f_4 = null, $v_5 = null, $k_6 = null) use ($MonadAsk0_1_0, $ask_2_1, $dictMonadReader_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((((($MonadAsk0_1_0)['Monad0'])(null))['Bind1'])(null))['bind'])($ask_2_1))(function($r_7 = null) use ($dictMonadReader_0, $f_4, $k_6, $v_5) {
  $__num = \func_num_args();
  $__res = ((($dictMonadReader_0)['local'])($f_4))(($v_5)(((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])((($dictMonadReader_0)['local'])(function($v_8 = null) use ($r_7) {
  $__num = \func_num_args();
  $__res = $r_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))($k_6)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "MonadAsk0" => function($_dollar__unused_4 = null) use ($monadAskContT1_3_2) {
  $__num = \func_num_args();
  $__res = $monadAskContT1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Cont_Trans_monadContContT
$GLOBALS['Control_Monad_Cont_Trans_monadContContT'] = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $monadContT1_1_0 = ($GLOBALS['Control_Monad_Cont_Trans_monadContT'])($dictMonad_0);
  $__res = ["callCC" => (function() {
  $__fn = function($f_2 = null, $k_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($f_2)((function() use ($k_3) {
  $__fn = function($a_4 = null, $v1_5 = null) use ($k_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($k_3)($a_4);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($k_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Monad0" => function($_dollar__unused_2 = null) use ($monadContT1_1_0) {
  $__num = \func_num_args();
  $__res = $monadContT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Cont_Trans_monadEffectContT
$GLOBALS['Control_Monad_Cont_Trans_monadEffectContT'] = function($dictMonadEffect_0 = null) {
  $__num = \func_num_args();
  $Monad0_1_0 = (($dictMonadEffect_0)['Monad0'])(null);
  $monadContT1_2_1 = ($GLOBALS['Control_Monad_Cont_Trans_monadContT'])($Monad0_1_0);
  $__res = ["liftEffect" => ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])((($GLOBALS['Control_Monad_Cont_Trans_monadTransContT'])['lift'])($Monad0_1_0)))(($dictMonadEffect_0)['liftEffect']), "Monad0" => function($_dollar__unused_3 = null) use ($monadContT1_2_1) {
  $__num = \func_num_args();
  $__res = $monadContT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Cont_Trans_monadStateContT
$GLOBALS['Control_Monad_Cont_Trans_monadStateContT'] = function($dictMonadState_0 = null) {
  $__num = \func_num_args();
  $Monad0_1_0 = (($dictMonadState_0)['Monad0'])(null);
  $monadContT1_2_1 = ($GLOBALS['Control_Monad_Cont_Trans_monadContT'])($Monad0_1_0);
  $__res = ["state" => ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])((($GLOBALS['Control_Monad_Cont_Trans_monadTransContT'])['lift'])($Monad0_1_0)))(($dictMonadState_0)['state']), "Monad0" => function($_dollar__unused_3 = null) use ($monadContT1_2_1) {
  $__num = \func_num_args();
  $__res = $monadContT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Cont_Trans_monadSTContT
$GLOBALS['Control_Monad_Cont_Trans_monadSTContT'] = function($dictMonadST_0 = null) {
  $__num = \func_num_args();
  $Monad0_1_0 = (($dictMonadST_0)['Monad0'])(null);
  $monadContT1_2_1 = ($GLOBALS['Control_Monad_Cont_Trans_monadContT'])($Monad0_1_0);
  $__res = ["liftST" => ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])((($GLOBALS['Control_Monad_Cont_Trans_monadTransContT'])['lift'])($Monad0_1_0)))(($dictMonadST_0)['liftST']), "Monad0" => function($_dollar__unused_3 = null) use ($monadContT1_2_1) {
  $__num = \func_num_args();
  $__res = $monadContT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Cont_Trans_monoidContT
$GLOBALS['Control_Monad_Cont_Trans_monoidContT'] = (function() {
  $__fn = function($dictApplicative_0 = null, $dictMonoid_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_2_0 = (($dictMonoid_1)['Semigroup0'])(null);
  $semigroupContT2_3_1 = ["append" => (function() use ($__local_var_2_0) {
  $__fn = function($a_3 = null, $b_4 = null, $k_5 = null) use ($__local_var_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($a_3)(function($a_6 = null) use ($__local_var_2_0, $b_4, $k_5) {
  $__num = \func_num_args();
  $__local_var_7_1 = (($__local_var_2_0)['append'])($a_6);
  $__res = ($b_4)(function($a_8 = null) use ($__local_var_7_1, $k_5) {
  $__num = \func_num_args();
  $__res = ($k_5)(($__local_var_7_1)($a_8));
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
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()];
  $__local_var_4_3 = ($dictMonoid_1)['mempty'];
  $__res = ["mempty" => function($k_5 = null) use ($__local_var_4_3) {
  $__num = \func_num_args();
  $__res = ($k_5)($__local_var_4_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Semigroup0" => function($_dollar__unused_4 = null) use ($semigroupContT2_3_1) {
  $__num = \func_num_args();
  $__res = $semigroupContT2_3_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

