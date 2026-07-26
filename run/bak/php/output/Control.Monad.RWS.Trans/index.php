<?php

namespace Control\Monad\RWS\Trans;

// ALL IMPORTS: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Lazy, Control.Monad, Control.Monad.Error.Class, Control.Monad.RWS.Trans, Control.Monad.Reader.Class, Control.Monad.Rec.Class, Control.Monad.ST.Class, Control.Monad.State.Class, Control.Monad.Trans.Class, Control.Monad.Writer.Class, Control.Plus, Control.Semigroupoid, Data.Function, Data.Functor, Data.Monoid, Data.Newtype, Data.Semigroup, Data.Tuple, Data.Unit, Effect.Class, Prelude, Prim
// TO REQUIRE: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Lazy, Control.Monad, Control.Monad.Error.Class, Control.Monad.RWS.Trans, Control.Monad.Reader.Class, Control.Monad.Rec.Class, Control.Monad.ST.Class, Control.Monad.State.Class, Control.Monad.Trans.Class, Control.Monad.Writer.Class, Control.Plus, Control.Semigroupoid, Data.Function, Data.Functor, Data.Monoid, Data.Newtype, Data.Semigroup, Data.Tuple, Data.Unit, Effect.Class, Prelude
require_once __DIR__ . '/../Control.Alt/index.php';
require_once __DIR__ . '/../Control.Alternative/index.php';
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Lazy/index.php';
require_once __DIR__ . '/../Control.Monad/index.php';
require_once __DIR__ . '/../Control.Monad.Error.Class/index.php';
require_once __DIR__ . '/../Control.Monad.RWS.Trans/index.php';
require_once __DIR__ . '/../Control.Monad.Reader.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Rec.Class/index.php';
require_once __DIR__ . '/../Control.Monad.ST.Class/index.php';
require_once __DIR__ . '/../Control.Monad.State.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Trans.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Writer.Class/index.php';
require_once __DIR__ . '/../Control.Plus/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Monoid/index.php';
require_once __DIR__ . '/../Data.Newtype/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Tuple/index.php';
require_once __DIR__ . '/../Data.Unit/index.php';
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


// Control_Monad_RWS_Trans_RWSResult
$GLOBALS['Control_Monad_RWS_Trans_RWSResult'] = (function() {
  $__fn = function($value0 = null, $value1 = null, $value2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = new Phpurs_Data3("RWSResult", $value0, $value1, $value2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Control_Monad_RWS_Trans_RWST
$GLOBALS['Control_Monad_RWS_Trans_RWST'] = function($x_0 = null) {
  $__num = \func_num_args();
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_RWS_Trans_withRWST
$GLOBALS['Control_Monad_RWS_Trans_withRWST'] = (function() {
  $__fn = function($f_0 = null, $m_1 = null, $r_2 = null, $s_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__local_var_4_0 = (($f_0)($r_2))($s_3);
  $__res = (($m_1)(($__local_var_4_0)->{'value0'}))(($__local_var_4_0)->{'value1'});
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})();

// Control_Monad_RWS_Trans_runRWST
$GLOBALS['Control_Monad_RWS_Trans_runRWST'] = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = $v_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_RWS_Trans_newtypeRWST
$GLOBALS['Control_Monad_RWS_Trans_newtypeRWST'] = ["Coercible0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Prim_undefined'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Control_Monad_RWS_Trans_monadTransRWST
$GLOBALS['Control_Monad_RWS_Trans_monadTransRWST'] = function($dictMonoid_0 = null) {
  $__num = \func_num_args();
  $mempty_1_0 = ($dictMonoid_0)['mempty'];
  $__res = ["lift" => (function() use ($mempty_1_0) {
  $__fn = function($dictMonad_2 = null, $m_3 = null, $v_4 = null, $s_5 = null) use ($mempty_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = ((((($dictMonad_2)['Bind1'])($GLOBALS['Prim_undefined']))['bind'])($m_3))(function($a_6 = null) use ($dictMonad_2, $mempty_1_0, $s_5) {
  $__num = \func_num_args();
  $__res = (((($dictMonad_2)['Applicative0'])($GLOBALS['Prim_undefined']))['pure'])(new Phpurs_Data3("RWSResult", $s_5, $a_6, $mempty_1_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_RWS_Trans_mapRWST
$GLOBALS['Control_Monad_RWS_Trans_mapRWST'] = (function() {
  $__fn = function($f_0 = null, $v_1 = null, $r_2 = null, $s_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = ($f_0)((($v_1)($r_2))($s_3));
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})();

// Control_Monad_RWS_Trans_lazyRWST
$GLOBALS['Control_Monad_RWS_Trans_lazyRWST'] = ["defer" => (function() {
  $__fn = function($f_0 = null, $r_1 = null, $s_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($f_0)($GLOBALS['Data_Unit_unit']))($r_1))($s_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()];

// Control_Monad_RWS_Trans_functorRWST
$GLOBALS['Control_Monad_RWS_Trans_functorRWST'] = function($dictFunctor_0 = null) {
  $__num = \func_num_args();
  $__res = ["map" => (function() use ($dictFunctor_0) {
  $__fn = function($f_1 = null, $v_2 = null, $r_3 = null, $s_4 = null) use ($dictFunctor_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = ((($dictFunctor_0)['map'])(function($v1_5 = null) use ($f_1) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data3("RWSResult", ($v1_5)->{'value0'}, ($f_1)(($v1_5)->{'value1'}), ($v1_5)->{'value2'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v_2)($r_3))($s_4));
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_RWS_Trans_execRWST
$GLOBALS['Control_Monad_RWS_Trans_execRWST'] = (function() {
  $__fn = function($dictMonad_0 = null, $v_1 = null, $r_2 = null, $s_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = ((((($dictMonad_0)['Bind1'])($GLOBALS['Prim_undefined']))['bind'])((($v_1)($r_2))($s_3)))(function($v1_4 = null) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = (((($dictMonad_0)['Applicative0'])($GLOBALS['Prim_undefined']))['pure'])(new Phpurs_Data2("Tuple", ($v1_4)->{'value0'}, ($v1_4)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})();

// Control_Monad_RWS_Trans_evalRWST
$GLOBALS['Control_Monad_RWS_Trans_evalRWST'] = (function() {
  $__fn = function($dictMonad_0 = null, $v_1 = null, $r_2 = null, $s_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = ((((($dictMonad_0)['Bind1'])($GLOBALS['Prim_undefined']))['bind'])((($v_1)($r_2))($s_3)))(function($v1_4 = null) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = (((($dictMonad_0)['Applicative0'])($GLOBALS['Prim_undefined']))['pure'])(new Phpurs_Data2("Tuple", ($v1_4)->{'value1'}, ($v1_4)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})();

// Control_Monad_RWS_Trans_applyRWST
$GLOBALS['Control_Monad_RWS_Trans_applyRWST'] = function($dictBind_0 = null) {
  $__num = \func_num_args();
  $Functor0_1_0 = (((($dictBind_0)['Apply0'])($GLOBALS['Prim_undefined']))['Functor0'])($GLOBALS['Prim_undefined']);
  $functorRWST1_2_1 = ["map" => (function() use ($Functor0_1_0) {
  $__fn = function($f_2 = null, $v_3 = null, $r_4 = null, $s_5 = null) use ($Functor0_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = ((($Functor0_1_0)['map'])(function($v1_6 = null) use ($f_2) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data3("RWSResult", ($v1_6)->{'value0'}, ($f_2)(($v1_6)->{'value1'}), ($v1_6)->{'value2'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v_3)($r_4))($s_5));
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})()];
  $__res = function($dictMonoid_3 = null) use ($Functor0_1_0, $dictBind_0, $functorRWST1_2_1) {
  $__num = \func_num_args();
  $__res = ["apply" => (function() use ($Functor0_1_0, $dictBind_0, $dictMonoid_3) {
  $__fn = function($v_4 = null, $v1_5 = null, $r_6 = null, $s_7 = null) use ($Functor0_1_0, $dictBind_0, $dictMonoid_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = ((($dictBind_0)['bind'])((($v_4)($r_6))($s_7)))(function($v2_8 = null) use ($Functor0_1_0, $dictMonoid_3, $r_6, $v1_5) {
  $__num = \func_num_args();
  $__local_var_9_2 = ($v2_8)->{'value2'};
  $__res = ((($Functor0_1_0)['map'])(function($v3_10 = null) use ($__local_var_9_2, $dictMonoid_3, $v2_8) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data3("RWSResult", ($v3_10)->{'value0'}, (($v2_8)->{'value1'})(($v3_10)->{'value1'}), ((((($dictMonoid_3)['Semigroup0'])($GLOBALS['Prim_undefined']))['append'])($__local_var_9_2))(($v3_10)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v1_5)($r_6))(($v2_8)->{'value0'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($dollar__unused_4 = null) use ($functorRWST1_2_1) {
  $__num = \func_num_args();
  $__res = $functorRWST1_2_1;
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

// Control_Monad_RWS_Trans_bindRWST
$GLOBALS['Control_Monad_RWS_Trans_bindRWST'] = function($dictBind_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (((($dictBind_0)['Apply0'])($GLOBALS['Prim_undefined']))['Functor0'])($GLOBALS['Prim_undefined']);
  $applyRWST1_2_1 = ($GLOBALS['Control_Monad_RWS_Trans_applyRWST'])($dictBind_0);
  $__res = function($dictMonoid_3 = null) use ($__local_var_1_0, $applyRWST1_2_1, $dictBind_0) {
  $__num = \func_num_args();
  $applyRWST2_4_2 = ($applyRWST1_2_1)($dictMonoid_3);
  $__res = ["bind" => (function() use ($__local_var_1_0, $dictBind_0, $dictMonoid_3) {
  $__fn = function($v_5 = null, $f_6 = null, $r_7 = null, $s_8 = null) use ($__local_var_1_0, $dictBind_0, $dictMonoid_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = ((($dictBind_0)['bind'])((($v_5)($r_7))($s_8)))(function($v1_9 = null) use ($__local_var_1_0, $dictMonoid_3, $f_6, $r_7) {
  $__num = \func_num_args();
  $__local_var_10_3 = ($v1_9)->{'value2'};
  $__res = ((($__local_var_1_0)['map'])(function($v3_11 = null) use ($__local_var_10_3, $dictMonoid_3) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data3("RWSResult", ($v3_11)->{'value0'}, ($v3_11)->{'value1'}, ((((($dictMonoid_3)['Semigroup0'])($GLOBALS['Prim_undefined']))['append'])($__local_var_10_3))(($v3_11)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($f_6)(($v1_9)->{'value1'}))($r_7))(($v1_9)->{'value0'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(), "Apply0" => function($dollar__unused_5 = null) use ($applyRWST2_4_2) {
  $__num = \func_num_args();
  $__res = $applyRWST2_4_2;
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

// Control_Monad_RWS_Trans_semigroupRWST
$GLOBALS['Control_Monad_RWS_Trans_semigroupRWST'] = function($dictBind_0 = null) {
  $__num = \func_num_args();
  $applyRWST1_1_0 = ($GLOBALS['Control_Monad_RWS_Trans_applyRWST'])($dictBind_0);
  $__res = function($dictMonoid_2 = null) use ($applyRWST1_1_0) {
  $__num = \func_num_args();
  $lift2_3_1 = ($GLOBALS['Control_Apply_lift2'])(($applyRWST1_1_0)($dictMonoid_2));
  $__res = function($dictSemigroup_4 = null) use ($lift2_3_1) {
  $__num = \func_num_args();
  $__res = ["append" => ($lift2_3_1)(($dictSemigroup_4)['append'])];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_RWS_Trans_applicativeRWST
$GLOBALS['Control_Monad_RWS_Trans_applicativeRWST'] = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $applyRWST1_1_0 = ($GLOBALS['Control_Monad_RWS_Trans_applyRWST'])((($dictMonad_0)['Bind1'])($GLOBALS['Prim_undefined']));
  $__res = function($dictMonoid_2 = null) use ($applyRWST1_1_0, $dictMonad_0) {
  $__num = \func_num_args();
  $mempty_3_1 = ($dictMonoid_2)['mempty'];
  $applyRWST2_4_2 = ($applyRWST1_1_0)($dictMonoid_2);
  $__res = ["pure" => (function() use ($dictMonad_0, $mempty_3_1) {
  $__fn = function($a_5 = null, $v_6 = null, $s_7 = null) use ($dictMonad_0, $mempty_3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($dictMonad_0)['Applicative0'])($GLOBALS['Prim_undefined']))['pure'])(new Phpurs_Data3("RWSResult", $s_7, $a_5, $mempty_3_1));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "Apply0" => function($dollar__unused_5 = null) use ($applyRWST2_4_2) {
  $__num = \func_num_args();
  $__res = $applyRWST2_4_2;
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

// Control_Monad_RWS_Trans_monadRWST
$GLOBALS['Control_Monad_RWS_Trans_monadRWST'] = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $applicativeRWST1_1_0 = ($GLOBALS['Control_Monad_RWS_Trans_applicativeRWST'])($dictMonad_0);
  $bindRWST1_2_1 = ($GLOBALS['Control_Monad_RWS_Trans_bindRWST'])((($dictMonad_0)['Bind1'])($GLOBALS['Prim_undefined']));
  $__res = function($dictMonoid_3 = null) use ($applicativeRWST1_1_0, $bindRWST1_2_1) {
  $__num = \func_num_args();
  $applicativeRWST2_4_2 = ($applicativeRWST1_1_0)($dictMonoid_3);
  $bindRWST2_5_3 = ($bindRWST1_2_1)($dictMonoid_3);
  $__res = ["Applicative0" => function($dollar__unused_6 = null) use ($applicativeRWST2_4_2) {
  $__num = \func_num_args();
  $__res = $applicativeRWST2_4_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($dollar__unused_6 = null) use ($bindRWST2_5_3) {
  $__num = \func_num_args();
  $__res = $bindRWST2_5_3;
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

// Control_Monad_RWS_Trans_monadAskRWST
$GLOBALS['Control_Monad_RWS_Trans_monadAskRWST'] = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $monadRWST1_1_0 = ($GLOBALS['Control_Monad_RWS_Trans_monadRWST'])($dictMonad_0);
  $__res = function($dictMonoid_2 = null) use ($dictMonad_0, $monadRWST1_1_0) {
  $__num = \func_num_args();
  $mempty_3_1 = ($dictMonoid_2)['mempty'];
  $monadRWST2_4_2 = ($monadRWST1_1_0)($dictMonoid_2);
  $__res = ["ask" => (function() use ($dictMonad_0, $mempty_3_1) {
  $__fn = function($r_5 = null, $s_6 = null) use ($dictMonad_0, $mempty_3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($dictMonad_0)['Applicative0'])($GLOBALS['Prim_undefined']))['pure'])(new Phpurs_Data3("RWSResult", $s_6, $r_5, $mempty_3_1));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Monad0" => function($dollar__unused_5 = null) use ($monadRWST2_4_2) {
  $__num = \func_num_args();
  $__res = $monadRWST2_4_2;
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

// Control_Monad_RWS_Trans_monadReaderRWST
$GLOBALS['Control_Monad_RWS_Trans_monadReaderRWST'] = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $monadAskRWST1_1_0 = ($GLOBALS['Control_Monad_RWS_Trans_monadAskRWST'])($dictMonad_0);
  $__res = function($dictMonoid_2 = null) use ($monadAskRWST1_1_0) {
  $__num = \func_num_args();
  $monadAskRWST2_3_1 = ($monadAskRWST1_1_0)($dictMonoid_2);
  $__res = ["local" => (function() {
  $__fn = function($f_4 = null, $m_5 = null, $r_6 = null, $s_7 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = (($m_5)(($f_4)($r_6)))($s_7);
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(), "MonadAsk0" => function($dollar__unused_4 = null) use ($monadAskRWST2_3_1) {
  $__num = \func_num_args();
  $__res = $monadAskRWST2_3_1;
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

// Control_Monad_RWS_Trans_monadEffectRWS
$GLOBALS['Control_Monad_RWS_Trans_monadEffectRWS'] = (function() {
  $__fn = function($dictMonoid_0 = null, $dictMonadEffect_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $Monad0_2_0 = (($dictMonadEffect_1)['Monad0'])($GLOBALS['Prim_undefined']);
  $monadRWST1_3_1 = (($GLOBALS['Control_Monad_RWS_Trans_monadRWST'])($Monad0_2_0))($dictMonoid_0);
  $__res = ["liftEffect" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(((($GLOBALS['Control_Monad_RWS_Trans_monadTransRWST'])($dictMonoid_0))['lift'])($Monad0_2_0)))(($dictMonadEffect_1)['liftEffect']), "Monad0" => function($dollar__unused_4 = null) use ($monadRWST1_3_1) {
  $__num = \func_num_args();
  $__res = $monadRWST1_3_1;
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

// Control_Monad_RWS_Trans_monadRecRWST
$GLOBALS['Control_Monad_RWS_Trans_monadRecRWST'] = function($dictMonadRec_0 = null) {
  $__num = \func_num_args();
  $Monad0_1_0 = (($dictMonadRec_0)['Monad0'])($GLOBALS['Prim_undefined']);
  $monadRWST1_2_1 = ($GLOBALS['Control_Monad_RWS_Trans_monadRWST'])($Monad0_1_0);
  $__res = function($dictMonoid_3 = null) use ($Monad0_1_0, $dictMonadRec_0, $monadRWST1_2_1) {
  $__num = \func_num_args();
  $__local_var_4_2 = (($dictMonoid_3)['Semigroup0'])($GLOBALS['Prim_undefined']);
  $mempty_5_3 = ($dictMonoid_3)['mempty'];
  $monadRWST2_6_4 = ($monadRWST1_2_1)($dictMonoid_3);
  $__res = ["tailRecM" => (function() use ($Monad0_1_0, $__local_var_4_2, $dictMonadRec_0, $mempty_5_3) {
  $__fn = function($k_7 = null, $a_8 = null, $r_9 = null, $s_10 = null) use ($Monad0_1_0, $__local_var_4_2, $dictMonadRec_0, $mempty_5_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = ((($dictMonadRec_0)['tailRecM'])(function($v_11 = null) use ($Monad0_1_0, $__local_var_4_2, $k_7, $r_9) {
  $__num = \func_num_args();
  $__local_var_12_5 = ($v_11)->{'value2'};
  $__res = ((((($Monad0_1_0)['Bind1'])($GLOBALS['Prim_undefined']))['bind'])(((($k_7)(($v_11)->{'value1'}))($r_9))(($v_11)->{'value0'})))(function($v2_13 = null) use ($Monad0_1_0, $__local_var_12_5, $__local_var_4_2) {
  $__num = \func_num_args();
  $__t6 = null;;
  if ((is_object(($v2_13)->{'value1'}) && ((($v2_13)->{'value1'})->{'tag'} === "Loop"))) {
$__t6 = new Phpurs_Data1("Loop", new Phpurs_Data3("RWSResult", ($v2_13)->{'value0'}, (($v2_13)->{'value1'})->{'value0'}, ((($__local_var_4_2)['append'])($__local_var_12_5))(($v2_13)->{'value2'})));
goto end_branch_6;;
};
  if ((is_object(($v2_13)->{'value1'}) && ((($v2_13)->{'value1'})->{'tag'} === "Done"))) {
$__t6 = new Phpurs_Data1("Done", new Phpurs_Data3("RWSResult", ($v2_13)->{'value0'}, (($v2_13)->{'value1'})->{'value0'}, ((($__local_var_4_2)['append'])($__local_var_12_5))(($v2_13)->{'value2'})));
goto end_branch_6;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t6 = null;
  end_branch_6:;
  $__res = (((($Monad0_1_0)['Applicative0'])($GLOBALS['Prim_undefined']))['pure'])($__t6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(new Phpurs_Data3("RWSResult", $s_10, $a_8, $mempty_5_3));
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(), "Monad0" => function($dollar__unused_7 = null) use ($monadRWST2_6_4) {
  $__num = \func_num_args();
  $__res = $monadRWST2_6_4;
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

// Control_Monad_RWS_Trans_monadStateRWST
$GLOBALS['Control_Monad_RWS_Trans_monadStateRWST'] = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $monadRWST1_1_0 = ($GLOBALS['Control_Monad_RWS_Trans_monadRWST'])($dictMonad_0);
  $__res = function($dictMonoid_2 = null) use ($dictMonad_0, $monadRWST1_1_0) {
  $__num = \func_num_args();
  $mempty_3_1 = ($dictMonoid_2)['mempty'];
  $monadRWST2_4_2 = ($monadRWST1_1_0)($dictMonoid_2);
  $__res = ["state" => (function() use ($dictMonad_0, $mempty_3_1) {
  $__fn = function($f_5 = null, $v_6 = null, $s_7 = null) use ($dictMonad_0, $mempty_3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $v1_8_3 = ($f_5)($s_7);
  $__res = (((($dictMonad_0)['Applicative0'])($GLOBALS['Prim_undefined']))['pure'])(new Phpurs_Data3("RWSResult", ($v1_8_3)->{'value1'}, ($v1_8_3)->{'value0'}, $mempty_3_1));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "Monad0" => function($dollar__unused_5 = null) use ($monadRWST2_4_2) {
  $__num = \func_num_args();
  $__res = $monadRWST2_4_2;
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

// Control_Monad_RWS_Trans_monadTellRWST
$GLOBALS['Control_Monad_RWS_Trans_monadTellRWST'] = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $monadRWST1_1_0 = ($GLOBALS['Control_Monad_RWS_Trans_monadRWST'])($dictMonad_0);
  $__res = function($dictMonoid_2 = null) use ($dictMonad_0, $monadRWST1_1_0) {
  $__num = \func_num_args();
  $Semigroup0_3_1 = (($dictMonoid_2)['Semigroup0'])($GLOBALS['Prim_undefined']);
  $monadRWST2_4_2 = ($monadRWST1_1_0)($dictMonoid_2);
  $__res = ["tell" => (function() use ($dictMonad_0) {
  $__fn = function($w_5 = null, $v_6 = null, $s_7 = null) use ($dictMonad_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($dictMonad_0)['Applicative0'])($GLOBALS['Prim_undefined']))['pure'])(new Phpurs_Data3("RWSResult", $s_7, $GLOBALS['Data_Unit_unit'], $w_5));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "Semigroup0" => function($dollar__unused_5 = null) use ($Semigroup0_3_1) {
  $__num = \func_num_args();
  $__res = $Semigroup0_3_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad1" => function($dollar__unused_5 = null) use ($monadRWST2_4_2) {
  $__num = \func_num_args();
  $__res = $monadRWST2_4_2;
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

// Control_Monad_RWS_Trans_monadWriterRWST
$GLOBALS['Control_Monad_RWS_Trans_monadWriterRWST'] = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($dictMonad_0)['Bind1'])($GLOBALS['Prim_undefined']);
  $__local_var_2_1 = (($dictMonad_0)['Applicative0'])($GLOBALS['Prim_undefined']);
  $monadTellRWST1_3_2 = ($GLOBALS['Control_Monad_RWS_Trans_monadTellRWST'])($dictMonad_0);
  $__res = function($dictMonoid_4 = null) use ($__local_var_1_0, $__local_var_2_1, $monadTellRWST1_3_2) {
  $__num = \func_num_args();
  $monadTellRWST2_5_3 = ($monadTellRWST1_3_2)($dictMonoid_4);
  $__res = ["listen" => (function() use ($__local_var_1_0, $__local_var_2_1) {
  $__fn = function($m_6 = null, $r_7 = null, $s_8 = null) use ($__local_var_1_0, $__local_var_2_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($__local_var_1_0)['bind'])((($m_6)($r_7))($s_8)))(function($v_9 = null) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = (($__local_var_2_1)['pure'])(new Phpurs_Data3("RWSResult", ($v_9)->{'value0'}, new Phpurs_Data2("Tuple", ($v_9)->{'value1'}, ($v_9)->{'value2'}), ($v_9)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "pass" => (function() use ($__local_var_1_0, $__local_var_2_1) {
  $__fn = function($m_6 = null, $r_7 = null, $s_8 = null) use ($__local_var_1_0, $__local_var_2_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($__local_var_1_0)['bind'])((($m_6)($r_7))($s_8)))(function($v_9 = null) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = (($__local_var_2_1)['pure'])(new Phpurs_Data3("RWSResult", ($v_9)->{'value0'}, (($v_9)->{'value1'})->{'value0'}, ((($v_9)->{'value1'})->{'value1'})(($v_9)->{'value2'})));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "Monoid0" => function($dollar__unused_6 = null) use ($dictMonoid_4) {
  $__num = \func_num_args();
  $__res = $dictMonoid_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "MonadTell1" => function($dollar__unused_6 = null) use ($monadTellRWST2_5_3) {
  $__num = \func_num_args();
  $__res = $monadTellRWST2_5_3;
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

// Control_Monad_RWS_Trans_monadThrowRWST
$GLOBALS['Control_Monad_RWS_Trans_monadThrowRWST'] = function($dictMonadThrow_0 = null) {
  $__num = \func_num_args();
  $Monad0_1_0 = (($dictMonadThrow_0)['Monad0'])($GLOBALS['Prim_undefined']);
  $monadRWST1_2_1 = ($GLOBALS['Control_Monad_RWS_Trans_monadRWST'])($Monad0_1_0);
  $__res = function($dictMonoid_3 = null) use ($Monad0_1_0, $dictMonadThrow_0, $monadRWST1_2_1) {
  $__num = \func_num_args();
  $lift_4_2 = ((($GLOBALS['Control_Monad_RWS_Trans_monadTransRWST'])($dictMonoid_3))['lift'])($Monad0_1_0);
  $monadRWST2_5_3 = ($monadRWST1_2_1)($dictMonoid_3);
  $__res = ["throwError" => function($e_6 = null) use ($dictMonadThrow_0, $lift_4_2) {
  $__num = \func_num_args();
  $__res = ($lift_4_2)((($dictMonadThrow_0)['throwError'])($e_6));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad0" => function($dollar__unused_6 = null) use ($monadRWST2_5_3) {
  $__num = \func_num_args();
  $__res = $monadRWST2_5_3;
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

// Control_Monad_RWS_Trans_monadErrorRWST
$GLOBALS['Control_Monad_RWS_Trans_monadErrorRWST'] = function($dictMonadError_0 = null) {
  $__num = \func_num_args();
  $monadThrowRWST1_1_0 = ($GLOBALS['Control_Monad_RWS_Trans_monadThrowRWST'])((($dictMonadError_0)['MonadThrow0'])($GLOBALS['Prim_undefined']));
  $__res = function($dictMonoid_2 = null) use ($dictMonadError_0, $monadThrowRWST1_1_0) {
  $__num = \func_num_args();
  $monadThrowRWST2_3_1 = ($monadThrowRWST1_1_0)($dictMonoid_2);
  $__res = ["catchError" => (function() use ($dictMonadError_0) {
  $__fn = function($m_4 = null, $h_5 = null, $r_6 = null, $s_7 = null) use ($dictMonadError_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = ((($dictMonadError_0)['catchError'])((($m_4)($r_6))($s_7)))(function($e_8 = null) use ($h_5, $r_6, $s_7) {
  $__num = \func_num_args();
  $__res = ((($h_5)($e_8))($r_6))($s_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(), "MonadThrow0" => function($dollar__unused_4 = null) use ($monadThrowRWST2_3_1) {
  $__num = \func_num_args();
  $__res = $monadThrowRWST2_3_1;
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

// Control_Monad_RWS_Trans_monadSTRWST
$GLOBALS['Control_Monad_RWS_Trans_monadSTRWST'] = (function() {
  $__fn = function($dictMonoid_0 = null, $dictMonadST_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $Monad0_2_0 = (($dictMonadST_1)['Monad0'])($GLOBALS['Prim_undefined']);
  $monadRWST1_3_1 = (($GLOBALS['Control_Monad_RWS_Trans_monadRWST'])($Monad0_2_0))($dictMonoid_0);
  $__res = ["liftST" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(((($GLOBALS['Control_Monad_RWS_Trans_monadTransRWST'])($dictMonoid_0))['lift'])($Monad0_2_0)))(($dictMonadST_1)['liftST']), "Monad0" => function($dollar__unused_4 = null) use ($monadRWST1_3_1) {
  $__num = \func_num_args();
  $__res = $monadRWST1_3_1;
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

// Control_Monad_RWS_Trans_monoidRWST
$GLOBALS['Control_Monad_RWS_Trans_monoidRWST'] = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $applicativeRWST1_1_0 = ($GLOBALS['Control_Monad_RWS_Trans_applicativeRWST'])($dictMonad_0);
  $semigroupRWST1_2_1 = ($GLOBALS['Control_Monad_RWS_Trans_semigroupRWST'])((($dictMonad_0)['Bind1'])($GLOBALS['Prim_undefined']));
  $__res = function($dictMonoid_3 = null) use ($applicativeRWST1_1_0, $semigroupRWST1_2_1) {
  $__num = \func_num_args();
  $semigroupRWST2_4_2 = ($semigroupRWST1_2_1)($dictMonoid_3);
  $__res = function($dictMonoid1_5 = null) use ($applicativeRWST1_1_0, $dictMonoid_3, $semigroupRWST2_4_2) {
  $__num = \func_num_args();
  $semigroupRWST3_6_3 = ($semigroupRWST2_4_2)((($dictMonoid1_5)['Semigroup0'])($GLOBALS['Prim_undefined']));
  $__res = ["mempty" => ((($applicativeRWST1_1_0)($dictMonoid_3))['pure'])(($dictMonoid1_5)['mempty']), "Semigroup0" => function($dollar__unused_7 = null) use ($semigroupRWST3_6_3) {
  $__num = \func_num_args();
  $__res = $semigroupRWST3_6_3;
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_RWS_Trans_altRWST
$GLOBALS['Control_Monad_RWS_Trans_altRWST'] = function($dictAlt_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($dictAlt_0)['Functor0'])($GLOBALS['Prim_undefined']);
  $functorRWST1_2_1 = ["map" => (function() use ($__local_var_1_0) {
  $__fn = function($f_2 = null, $v_3 = null, $r_4 = null, $s_5 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = ((($__local_var_1_0)['map'])(function($v1_6 = null) use ($f_2) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data3("RWSResult", ($v1_6)->{'value0'}, ($f_2)(($v1_6)->{'value1'}), ($v1_6)->{'value2'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v_3)($r_4))($s_5));
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})()];
  $__res = ["alt" => (function() use ($dictAlt_0) {
  $__fn = function($v_3 = null, $v1_4 = null, $r_5 = null, $s_6 = null) use ($dictAlt_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = ((($dictAlt_0)['alt'])((($v_3)($r_5))($s_6)))((($v1_4)($r_5))($s_6));
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($dollar__unused_3 = null) use ($functorRWST1_2_1) {
  $__num = \func_num_args();
  $__res = $functorRWST1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_RWS_Trans_plusRWST
$GLOBALS['Control_Monad_RWS_Trans_plusRWST'] = function($dictPlus_0 = null) {
  $__num = \func_num_args();
  $empty_1_0 = ($dictPlus_0)['empty'];
  $altRWST1_2_1 = ($GLOBALS['Control_Monad_RWS_Trans_altRWST'])((($dictPlus_0)['Alt0'])($GLOBALS['Prim_undefined']));
  $__res = ["empty" => (function() use ($empty_1_0) {
  $__fn = function($v_3 = null, $v1_4 = null) use ($empty_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = $empty_1_0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Alt0" => function($dollar__unused_3 = null) use ($altRWST1_2_1) {
  $__num = \func_num_args();
  $__res = $altRWST1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_RWS_Trans_alternativeRWST
$GLOBALS['Control_Monad_RWS_Trans_alternativeRWST'] = (function() {
  $__fn = function($dictMonoid_0 = null, $dictAlternative_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $plusRWST1_2_0 = ($GLOBALS['Control_Monad_RWS_Trans_plusRWST'])((($dictAlternative_1)['Plus1'])($GLOBALS['Prim_undefined']));
  $__res = function($dictMonad_3 = null) use ($dictMonoid_0, $plusRWST1_2_0) {
  $__num = \func_num_args();
  $applicativeRWST1_4_1 = (($GLOBALS['Control_Monad_RWS_Trans_applicativeRWST'])($dictMonad_3))($dictMonoid_0);
  $__res = ["Applicative0" => function($dollar__unused_5 = null) use ($applicativeRWST1_4_1) {
  $__num = \func_num_args();
  $__res = $applicativeRWST1_4_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Plus1" => function($dollar__unused_5 = null) use ($plusRWST1_2_0) {
  $__num = \func_num_args();
  $__res = $plusRWST1_2_0;
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
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

