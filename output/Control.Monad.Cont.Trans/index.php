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
if (!\function_exists(__NAMESPACE__ . '\\phpurs_execute_effect')) {
  function phpurs_execute_effect($val) {
    if (\is_callable($val)) {
      return $val($GLOBALS['Data_Unit_unit']);
    }
    return $val;
  }
}

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };




// Control_Monad_Cont_Trans_ContT
function majControl_majMonad_majCont_majTrans_majContmajT($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majCont_majTrans_majContmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Cont_Trans_ContT'] = __NAMESPACE__ . '\\majControl_majMonad_majCont_majTrans_majContmajT';

// Control_Monad_Cont_Trans_withContT
function majControl_majMonad_majCont_majTrans_withmajContmajT($f_0, $v_1 = null, $k_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majCont_majTrans_withmajContmajT';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($v_1)(($f_0)($k_2));
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Control_Monad_Cont_Trans_withContT'] = __NAMESPACE__ . '\\majControl_majMonad_majCont_majTrans_withmajContmajT';

// Control_Monad_Cont_Trans_runContT
function majControl_majMonad_majCont_majTrans_runmajContmajT($v_0, $k_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majCont_majTrans_runmajContmajT';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($v_0)($k_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Control_Monad_Cont_Trans_runContT'] = __NAMESPACE__ . '\\majControl_majMonad_majCont_majTrans_runmajContmajT';

// Control_Monad_Cont_Trans_newtypeContT
$GLOBALS['Control_Monad_Cont_Trans_newtypeContT'] = (object)["Coercible0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Control_Monad_Cont_Trans_monadTransContT
$GLOBALS['Control_Monad_Cont_Trans_monadTransContT'] = (object)["lift" => function($dictMonad_0) {
  $__num = \func_num_args();
  $Bind1_1_0 = (($dictMonad_0)->{'Bind1'})(null);
  $__res = function($m_2) use ($Bind1_1_0) {
  $__num = \func_num_args();
  $__res = function($k_3) use ($Bind1_1_0, $m_2) {
  $__num = \func_num_args();
  $__res = ((($Bind1_1_0)->{'bind'})($m_2))($k_3);
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
}];

// Control_Monad_Cont_Trans_mapContT
function majControl_majMonad_majCont_majTrans_mapmajContmajT($f_0, $v_1 = null, $k_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majCont_majTrans_mapmajContmajT';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($f_0)(($v_1)($k_2));
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Control_Monad_Cont_Trans_mapContT'] = __NAMESPACE__ . '\\majControl_majMonad_majCont_majTrans_mapmajContmajT';

// Control_Monad_Cont_Trans_functorContT
function majControl_majMonad_majCont_majTrans_functormajContmajT($dictFunctor_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majCont_majTrans_functormajContmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["map" => function($f_1) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($f_1) {
  $__num = \func_num_args();
  $__res = function($k_3) use ($f_1, $v_2) {
  $__num = \func_num_args();
  $__res = ($v_2)(function($a_4) use ($f_1, $k_3) {
  $__num = \func_num_args();
  $__res = ($k_3)(($f_1)($a_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Cont_Trans_functorContT'] = __NAMESPACE__ . '\\majControl_majMonad_majCont_majTrans_functormajContmajT';

// Control_Monad_Cont_Trans_applyContT
function majControl_majMonad_majCont_majTrans_applymajContmajT($dictApply_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majCont_majTrans_applymajContmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $functorContT1_1_0 = (object)["map" => function($f_1) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($f_1) {
  $__num = \func_num_args();
  $__res = function($k_3) use ($f_1, $v_2) {
  $__num = \func_num_args();
  $__res = ($v_2)(function($a_4) use ($f_1, $k_3) {
  $__num = \func_num_args();
  $__res = ($k_3)(($f_1)($a_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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
}];
  $__res = (object)["apply" => function($v_2) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($v_2) {
  $__num = \func_num_args();
  $__res = function($k_4) use ($v1_3, $v_2) {
  $__num = \func_num_args();
  $__res = ($v_2)(function($g_5) use ($k_4, $v1_3) {
  $__num = \func_num_args();
  $__res = ($v1_3)(function($a_6) use ($g_5, $k_4) {
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_2) use ($functorContT1_1_0) {
  $__num = \func_num_args();
  $__res = $functorContT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Cont_Trans_applyContT'] = __NAMESPACE__ . '\\majControl_majMonad_majCont_majTrans_applymajContmajT';

// Control_Monad_Cont_Trans_bindContT
function majControl_majMonad_majCont_majTrans_bindmajContmajT($dictBind_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majCont_majTrans_bindmajContmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $functorContT1_1_0 = (object)["map" => function($f_1) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($f_1) {
  $__num = \func_num_args();
  $__res = function($k_3) use ($f_1, $v_2) {
  $__num = \func_num_args();
  $__res = ($v_2)(function($a_4) use ($f_1, $k_3) {
  $__num = \func_num_args();
  $__res = ($k_3)(($f_1)($a_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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
}];
  $applyContT1_1_0 = (object)["apply" => function($v_2) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($v_2) {
  $__num = \func_num_args();
  $__res = function($k_4) use ($v1_3, $v_2) {
  $__num = \func_num_args();
  $__res = ($v_2)(function($g_5) use ($k_4, $v1_3) {
  $__num = \func_num_args();
  $__res = ($v1_3)(function($a_6) use ($g_5, $k_4) {
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_2) use ($functorContT1_1_0) {
  $__num = \func_num_args();
  $__res = $functorContT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["bind" => function($v_2) {
  $__num = \func_num_args();
  $__res = function($k_3) use ($v_2) {
  $__num = \func_num_args();
  $__res = function($k_prime__4) use ($k_3, $v_2) {
  $__num = \func_num_args();
  $__res = ($v_2)(function($a_5) use ($k_3, $k_prime__4) {
  $__num = \func_num_args();
  $__res = (($k_3)($a_5))($k_prime__4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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
}, "Apply0" => function($_dollar___unused_2) use ($applyContT1_1_0) {
  $__num = \func_num_args();
  $__res = $applyContT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Cont_Trans_bindContT'] = __NAMESPACE__ . '\\majControl_majMonad_majCont_majTrans_bindmajContmajT';

// Control_Monad_Cont_Trans_semigroupContT
function majControl_majMonad_majCont_majTrans_semigroupmajContmajT($dictApply_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majCont_majTrans_semigroupmajContmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $functorContT1_1_0 = (object)["map" => function($f_1) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($f_1) {
  $__num = \func_num_args();
  $__res = function($k_3) use ($f_1, $v_2) {
  $__num = \func_num_args();
  $__res = ($v_2)(function($a_4) use ($f_1, $k_3) {
  $__num = \func_num_args();
  $__res = ($k_3)(($f_1)($a_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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
}];
  $applyContT1_1_0 = (object)["apply" => function($v_2) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($v_2) {
  $__num = \func_num_args();
  $__res = function($k_4) use ($v1_3, $v_2) {
  $__num = \func_num_args();
  $__res = ($v_2)(function($g_5) use ($k_4, $v1_3) {
  $__num = \func_num_args();
  $__res = ($v1_3)(function($a_6) use ($g_5, $k_4) {
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_2) use ($functorContT1_1_0) {
  $__num = \func_num_args();
  $__res = $functorContT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = function($dictSemigroup_2) use ($applyContT1_1_0) {
  $__num = \func_num_args();
  $Functor0_3_2 = (($applyContT1_1_0)->{'Functor0'})(null);
  $__local_var_4_3 = ($dictSemigroup_2)->{'append'};
  $__res = (object)["append" => function($a_5) use ($Functor0_3_2, $__local_var_4_3, $applyContT1_1_0) {
  $__num = \func_num_args();
  $__res = function($b_6) use ($Functor0_3_2, $__local_var_4_3, $a_5, $applyContT1_1_0) {
  $__num = \func_num_args();
  $__res = ((($applyContT1_1_0)->{'apply'})(((($Functor0_3_2)->{'map'})($__local_var_4_3))($a_5)))($b_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Cont_Trans_semigroupContT'] = __NAMESPACE__ . '\\majControl_majMonad_majCont_majTrans_semigroupmajContmajT';

// Control_Monad_Cont_Trans_applicativeContT
function majControl_majMonad_majCont_majTrans_applicativemajContmajT($dictApplicative_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majCont_majTrans_applicativemajContmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $functorContT1_1_0 = (object)["map" => function($f_1) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($f_1) {
  $__num = \func_num_args();
  $__res = function($k_3) use ($f_1, $v_2) {
  $__num = \func_num_args();
  $__res = ($v_2)(function($a_4) use ($f_1, $k_3) {
  $__num = \func_num_args();
  $__res = ($k_3)(($f_1)($a_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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
}];
  $applyContT1_1_0 = (object)["apply" => function($v_2) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($v_2) {
  $__num = \func_num_args();
  $__res = function($k_4) use ($v1_3, $v_2) {
  $__num = \func_num_args();
  $__res = ($v_2)(function($g_5) use ($k_4, $v1_3) {
  $__num = \func_num_args();
  $__res = ($v1_3)(function($a_6) use ($g_5, $k_4) {
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_2) use ($functorContT1_1_0) {
  $__num = \func_num_args();
  $__res = $functorContT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["pure" => function($a_2) {
  $__num = \func_num_args();
  $__res = function($k_3) use ($a_2) {
  $__num = \func_num_args();
  $__res = ($k_3)($a_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_2) use ($applyContT1_1_0) {
  $__num = \func_num_args();
  $__res = $applyContT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Cont_Trans_applicativeContT'] = __NAMESPACE__ . '\\majControl_majMonad_majCont_majTrans_applicativemajContmajT';

// Control_Monad_Cont_Trans_monadContT
function majControl_majMonad_majCont_majTrans_monadmajContmajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majCont_majTrans_monadmajContmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $functorContT1_1_0 = (object)["map" => function($f_1) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($f_1) {
  $__num = \func_num_args();
  $__res = function($k_3) use ($f_1, $v_2) {
  $__num = \func_num_args();
  $__res = ($v_2)(function($a_4) use ($f_1, $k_3) {
  $__num = \func_num_args();
  $__res = ($k_3)(($f_1)($a_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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
}];
  $applyContT1_1_0 = (object)["apply" => function($v_2) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($v_2) {
  $__num = \func_num_args();
  $__res = function($k_4) use ($v1_3, $v_2) {
  $__num = \func_num_args();
  $__res = ($v_2)(function($g_5) use ($k_4, $v1_3) {
  $__num = \func_num_args();
  $__res = ($v1_3)(function($a_6) use ($g_5, $k_4) {
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_2) use ($functorContT1_1_0) {
  $__num = \func_num_args();
  $__res = $functorContT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeContT1_1_0 = (object)["pure" => function($a_2) {
  $__num = \func_num_args();
  $__res = function($k_3) use ($a_2) {
  $__num = \func_num_args();
  $__res = ($k_3)($a_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_2) use ($applyContT1_1_0) {
  $__num = \func_num_args();
  $__res = $applyContT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $functorContT1_2_3 = (object)["map" => function($f_2) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($f_2) {
  $__num = \func_num_args();
  $__res = function($k_4) use ($f_2, $v_3) {
  $__num = \func_num_args();
  $__res = ($v_3)(function($a_5) use ($f_2, $k_4) {
  $__num = \func_num_args();
  $__res = ($k_4)(($f_2)($a_5));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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
}];
  $applyContT1_2_3 = (object)["apply" => function($v_3) {
  $__num = \func_num_args();
  $__res = function($v1_4) use ($v_3) {
  $__num = \func_num_args();
  $__res = function($k_5) use ($v1_4, $v_3) {
  $__num = \func_num_args();
  $__res = ($v_3)(function($g_6) use ($k_5, $v1_4) {
  $__num = \func_num_args();
  $__res = ($v1_4)(function($a_7) use ($g_6, $k_5) {
  $__num = \func_num_args();
  $__res = ($k_5)(($g_6)($a_7));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_3) use ($functorContT1_2_3) {
  $__num = \func_num_args();
  $__res = $functorContT1_2_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $bindContT1_2_3 = (object)["bind" => function($v_3) {
  $__num = \func_num_args();
  $__res = function($k_4) use ($v_3) {
  $__num = \func_num_args();
  $__res = function($k_prime__5) use ($k_4, $v_3) {
  $__num = \func_num_args();
  $__res = ($v_3)(function($a_6) use ($k_4, $k_prime__5) {
  $__num = \func_num_args();
  $__res = (($k_4)($a_6))($k_prime__5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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
}, "Apply0" => function($_dollar___unused_3) use ($applyContT1_2_3) {
  $__num = \func_num_args();
  $__res = $applyContT1_2_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["Applicative0" => function($_dollar___unused_3) use ($applicativeContT1_1_0) {
  $__num = \func_num_args();
  $__res = $applicativeContT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_3) use ($bindContT1_2_3) {
  $__num = \func_num_args();
  $__res = $bindContT1_2_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Cont_Trans_monadContT'] = __NAMESPACE__ . '\\majControl_majMonad_majCont_majTrans_monadmajContmajT';

// Control_Monad_Cont_Trans_monadAskContT
function majControl_majMonad_majCont_majTrans_monadmajAskmajContmajT($dictMonadAsk_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majCont_majTrans_monadmajAskmajContmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $functorContT1_1_0 = (object)["map" => function($f_1) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($f_1) {
  $__num = \func_num_args();
  $__res = function($k_3) use ($f_1, $v_2) {
  $__num = \func_num_args();
  $__res = ($v_2)(function($a_4) use ($f_1, $k_3) {
  $__num = \func_num_args();
  $__res = ($k_3)(($f_1)($a_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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
}];
  $applyContT1_1_0 = (object)["apply" => function($v_2) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($v_2) {
  $__num = \func_num_args();
  $__res = function($k_4) use ($v1_3, $v_2) {
  $__num = \func_num_args();
  $__res = ($v_2)(function($g_5) use ($k_4, $v1_3) {
  $__num = \func_num_args();
  $__res = ($v1_3)(function($a_6) use ($g_5, $k_4) {
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_2) use ($functorContT1_1_0) {
  $__num = \func_num_args();
  $__res = $functorContT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeContT1_1_0 = (object)["pure" => function($a_2) {
  $__num = \func_num_args();
  $__res = function($k_3) use ($a_2) {
  $__num = \func_num_args();
  $__res = ($k_3)($a_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_2) use ($applyContT1_1_0) {
  $__num = \func_num_args();
  $__res = $applyContT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $functorContT1_2_3 = (object)["map" => function($f_2) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($f_2) {
  $__num = \func_num_args();
  $__res = function($k_4) use ($f_2, $v_3) {
  $__num = \func_num_args();
  $__res = ($v_3)(function($a_5) use ($f_2, $k_4) {
  $__num = \func_num_args();
  $__res = ($k_4)(($f_2)($a_5));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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
}];
  $applyContT1_2_3 = (object)["apply" => function($v_3) {
  $__num = \func_num_args();
  $__res = function($v1_4) use ($v_3) {
  $__num = \func_num_args();
  $__res = function($k_5) use ($v1_4, $v_3) {
  $__num = \func_num_args();
  $__res = ($v_3)(function($g_6) use ($k_5, $v1_4) {
  $__num = \func_num_args();
  $__res = ($v1_4)(function($a_7) use ($g_6, $k_5) {
  $__num = \func_num_args();
  $__res = ($k_5)(($g_6)($a_7));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_3) use ($functorContT1_2_3) {
  $__num = \func_num_args();
  $__res = $functorContT1_2_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $bindContT1_2_3 = (object)["bind" => function($v_3) {
  $__num = \func_num_args();
  $__res = function($k_4) use ($v_3) {
  $__num = \func_num_args();
  $__res = function($k_prime__5) use ($k_4, $v_3) {
  $__num = \func_num_args();
  $__res = ($v_3)(function($a_6) use ($k_4, $k_prime__5) {
  $__num = \func_num_args();
  $__res = (($k_4)($a_6))($k_prime__5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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
}, "Apply0" => function($_dollar___unused_3) use ($applyContT1_2_3) {
  $__num = \func_num_args();
  $__res = $applyContT1_2_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $monadContT1_1_0 = (object)["Applicative0" => function($_dollar___unused_3) use ($applicativeContT1_1_0) {
  $__num = \func_num_args();
  $__res = $applicativeContT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_3) use ($bindContT1_2_3) {
  $__num = \func_num_args();
  $__res = $bindContT1_2_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_2_7 = (((($dictMonadAsk_0)->{'Monad0'})(null))->{'Bind1'})(null);
  $__local_var_3_8 = ($dictMonadAsk_0)->{'ask'};
  $__res = (object)["ask" => function($k_4) use ($Bind1_2_7, $__local_var_3_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_2_7)->{'bind'})($__local_var_3_8))($k_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad0" => function($_dollar___unused_2) use ($monadContT1_1_0) {
  $__num = \func_num_args();
  $__res = $monadContT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Cont_Trans_monadAskContT'] = __NAMESPACE__ . '\\majControl_majMonad_majCont_majTrans_monadmajAskmajContmajT';

// Control_Monad_Cont_Trans_monadReaderContT
function majControl_majMonad_majCont_majTrans_monadmajReadermajContmajT($dictMonadReader_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majCont_majTrans_monadmajReadermajContmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $MonadAsk0_1_0 = (($dictMonadReader_0)->{'MonadAsk0'})(null);
  $Bind1_2_1 = (((($MonadAsk0_1_0)->{'Monad0'})(null))->{'Bind1'})(null);
  $ask_3_2 = ($MonadAsk0_1_0)->{'ask'};
  $functorContT1_4_3 = (object)["map" => function($f_4) {
  $__num = \func_num_args();
  $__res = function($v_5) use ($f_4) {
  $__num = \func_num_args();
  $__res = function($k_6) use ($f_4, $v_5) {
  $__num = \func_num_args();
  $__res = ($v_5)(function($a_7) use ($f_4, $k_6) {
  $__num = \func_num_args();
  $__res = ($k_6)(($f_4)($a_7));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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
}];
  $applyContT1_4_3 = (object)["apply" => function($v_5) {
  $__num = \func_num_args();
  $__res = function($v1_6) use ($v_5) {
  $__num = \func_num_args();
  $__res = function($k_7) use ($v1_6, $v_5) {
  $__num = \func_num_args();
  $__res = ($v_5)(function($g_8) use ($k_7, $v1_6) {
  $__num = \func_num_args();
  $__res = ($v1_6)(function($a_9) use ($g_8, $k_7) {
  $__num = \func_num_args();
  $__res = ($k_7)(($g_8)($a_9));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_5) use ($functorContT1_4_3) {
  $__num = \func_num_args();
  $__res = $functorContT1_4_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeContT1_4_3 = (object)["pure" => function($a_5) {
  $__num = \func_num_args();
  $__res = function($k_6) use ($a_5) {
  $__num = \func_num_args();
  $__res = ($k_6)($a_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_5) use ($applyContT1_4_3) {
  $__num = \func_num_args();
  $__res = $applyContT1_4_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $functorContT1_5_6 = (object)["map" => function($f_5) {
  $__num = \func_num_args();
  $__res = function($v_6) use ($f_5) {
  $__num = \func_num_args();
  $__res = function($k_7) use ($f_5, $v_6) {
  $__num = \func_num_args();
  $__res = ($v_6)(function($a_8) use ($f_5, $k_7) {
  $__num = \func_num_args();
  $__res = ($k_7)(($f_5)($a_8));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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
}];
  $applyContT1_5_6 = (object)["apply" => function($v_6) {
  $__num = \func_num_args();
  $__res = function($v1_7) use ($v_6) {
  $__num = \func_num_args();
  $__res = function($k_8) use ($v1_7, $v_6) {
  $__num = \func_num_args();
  $__res = ($v_6)(function($g_9) use ($k_8, $v1_7) {
  $__num = \func_num_args();
  $__res = ($v1_7)(function($a_10) use ($g_9, $k_8) {
  $__num = \func_num_args();
  $__res = ($k_8)(($g_9)($a_10));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_6) use ($functorContT1_5_6) {
  $__num = \func_num_args();
  $__res = $functorContT1_5_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $bindContT1_5_6 = (object)["bind" => function($v_6) {
  $__num = \func_num_args();
  $__res = function($k_7) use ($v_6) {
  $__num = \func_num_args();
  $__res = function($k_prime__8) use ($k_7, $v_6) {
  $__num = \func_num_args();
  $__res = ($v_6)(function($a_9) use ($k_7, $k_prime__8) {
  $__num = \func_num_args();
  $__res = (($k_7)($a_9))($k_prime__8);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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
}, "Apply0" => function($_dollar___unused_6) use ($applyContT1_5_6) {
  $__num = \func_num_args();
  $__res = $applyContT1_5_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $monadContT1_4_3 = (object)["Applicative0" => function($_dollar___unused_6) use ($applicativeContT1_4_3) {
  $__num = \func_num_args();
  $__res = $applicativeContT1_4_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_6) use ($bindContT1_5_6) {
  $__num = \func_num_args();
  $__res = $bindContT1_5_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_5_10 = (((($MonadAsk0_1_0)->{'Monad0'})(null))->{'Bind1'})(null);
  $__local_var_6_11 = ($MonadAsk0_1_0)->{'ask'};
  $monadAskContT1_4_3 = (object)["ask" => function($k_7) use ($Bind1_5_10, $__local_var_6_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_5_10)->{'bind'})($__local_var_6_11))($k_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad0" => function($_dollar___unused_5) use ($monadContT1_4_3) {
  $__num = \func_num_args();
  $__res = $monadContT1_4_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["local" => function($f_5) use ($Bind1_2_1, $ask_3_2, $dictMonadReader_0) {
  $__num = \func_num_args();
  $__res = function($v_6) use ($Bind1_2_1, $ask_3_2, $dictMonadReader_0, $f_5) {
  $__num = \func_num_args();
  $__res = function($k_7) use ($Bind1_2_1, $ask_3_2, $dictMonadReader_0, $f_5, $v_6) {
  $__num = \func_num_args();
  $__res = ((($Bind1_2_1)->{'bind'})($ask_3_2))(function($r_8) use ($dictMonadReader_0, $f_5, $k_7, $v_6) {
  $__num = \func_num_args();
  $__res = ((($dictMonadReader_0)->{'local'})($f_5))(($v_6)((($GLOBALS['Control_Semigroupoid_composeImpl'])((($dictMonadReader_0)->{'local'})(function($v_9) use ($r_8) {
  $__num = \func_num_args();
  $__res = $r_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))($k_7)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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
}, "MonadAsk0" => function($_dollar___unused_5) use ($monadAskContT1_4_3) {
  $__num = \func_num_args();
  $__res = $monadAskContT1_4_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Cont_Trans_monadReaderContT'] = __NAMESPACE__ . '\\majControl_majMonad_majCont_majTrans_monadmajReadermajContmajT';

// Control_Monad_Cont_Trans_monadContContT
function majControl_majMonad_majCont_majTrans_monadmajContmajContmajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majCont_majTrans_monadmajContmajContmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $functorContT1_1_0 = (object)["map" => function($f_1) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($f_1) {
  $__num = \func_num_args();
  $__res = function($k_3) use ($f_1, $v_2) {
  $__num = \func_num_args();
  $__res = ($v_2)(function($a_4) use ($f_1, $k_3) {
  $__num = \func_num_args();
  $__res = ($k_3)(($f_1)($a_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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
}];
  $applyContT1_1_0 = (object)["apply" => function($v_2) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($v_2) {
  $__num = \func_num_args();
  $__res = function($k_4) use ($v1_3, $v_2) {
  $__num = \func_num_args();
  $__res = ($v_2)(function($g_5) use ($k_4, $v1_3) {
  $__num = \func_num_args();
  $__res = ($v1_3)(function($a_6) use ($g_5, $k_4) {
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_2) use ($functorContT1_1_0) {
  $__num = \func_num_args();
  $__res = $functorContT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeContT1_1_0 = (object)["pure" => function($a_2) {
  $__num = \func_num_args();
  $__res = function($k_3) use ($a_2) {
  $__num = \func_num_args();
  $__res = ($k_3)($a_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_2) use ($applyContT1_1_0) {
  $__num = \func_num_args();
  $__res = $applyContT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $functorContT1_2_3 = (object)["map" => function($f_2) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($f_2) {
  $__num = \func_num_args();
  $__res = function($k_4) use ($f_2, $v_3) {
  $__num = \func_num_args();
  $__res = ($v_3)(function($a_5) use ($f_2, $k_4) {
  $__num = \func_num_args();
  $__res = ($k_4)(($f_2)($a_5));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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
}];
  $applyContT1_2_3 = (object)["apply" => function($v_3) {
  $__num = \func_num_args();
  $__res = function($v1_4) use ($v_3) {
  $__num = \func_num_args();
  $__res = function($k_5) use ($v1_4, $v_3) {
  $__num = \func_num_args();
  $__res = ($v_3)(function($g_6) use ($k_5, $v1_4) {
  $__num = \func_num_args();
  $__res = ($v1_4)(function($a_7) use ($g_6, $k_5) {
  $__num = \func_num_args();
  $__res = ($k_5)(($g_6)($a_7));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_3) use ($functorContT1_2_3) {
  $__num = \func_num_args();
  $__res = $functorContT1_2_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $bindContT1_2_3 = (object)["bind" => function($v_3) {
  $__num = \func_num_args();
  $__res = function($k_4) use ($v_3) {
  $__num = \func_num_args();
  $__res = function($k_prime__5) use ($k_4, $v_3) {
  $__num = \func_num_args();
  $__res = ($v_3)(function($a_6) use ($k_4, $k_prime__5) {
  $__num = \func_num_args();
  $__res = (($k_4)($a_6))($k_prime__5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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
}, "Apply0" => function($_dollar___unused_3) use ($applyContT1_2_3) {
  $__num = \func_num_args();
  $__res = $applyContT1_2_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $monadContT1_1_0 = (object)["Applicative0" => function($_dollar___unused_3) use ($applicativeContT1_1_0) {
  $__num = \func_num_args();
  $__res = $applicativeContT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_3) use ($bindContT1_2_3) {
  $__num = \func_num_args();
  $__res = $bindContT1_2_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["callCC" => function($f_2) {
  $__num = \func_num_args();
  $__res = function($k_3) use ($f_2) {
  $__num = \func_num_args();
  $__res = (($f_2)(function($a_4) use ($k_3) {
  $__num = \func_num_args();
  $__res = function($v1_5) use ($a_4, $k_3) {
  $__num = \func_num_args();
  $__res = ($k_3)($a_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($k_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad0" => function($_dollar___unused_2) use ($monadContT1_1_0) {
  $__num = \func_num_args();
  $__res = $monadContT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Cont_Trans_monadContContT'] = __NAMESPACE__ . '\\majControl_majMonad_majCont_majTrans_monadmajContmajContmajT';

// Control_Monad_Cont_Trans_monadEffectContT
function majControl_majMonad_majCont_majTrans_monadmajEffectmajContmajT($dictMonadEffect_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majCont_majTrans_monadmajEffectmajContmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $functorContT1_1_0 = (object)["map" => function($f_1) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($f_1) {
  $__num = \func_num_args();
  $__res = function($k_3) use ($f_1, $v_2) {
  $__num = \func_num_args();
  $__res = ($v_2)(function($a_4) use ($f_1, $k_3) {
  $__num = \func_num_args();
  $__res = ($k_3)(($f_1)($a_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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
}];
  $applyContT1_1_0 = (object)["apply" => function($v_2) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($v_2) {
  $__num = \func_num_args();
  $__res = function($k_4) use ($v1_3, $v_2) {
  $__num = \func_num_args();
  $__res = ($v_2)(function($g_5) use ($k_4, $v1_3) {
  $__num = \func_num_args();
  $__res = ($v1_3)(function($a_6) use ($g_5, $k_4) {
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_2) use ($functorContT1_1_0) {
  $__num = \func_num_args();
  $__res = $functorContT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeContT1_1_0 = (object)["pure" => function($a_2) {
  $__num = \func_num_args();
  $__res = function($k_3) use ($a_2) {
  $__num = \func_num_args();
  $__res = ($k_3)($a_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_2) use ($applyContT1_1_0) {
  $__num = \func_num_args();
  $__res = $applyContT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $functorContT1_2_3 = (object)["map" => function($f_2) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($f_2) {
  $__num = \func_num_args();
  $__res = function($k_4) use ($f_2, $v_3) {
  $__num = \func_num_args();
  $__res = ($v_3)(function($a_5) use ($f_2, $k_4) {
  $__num = \func_num_args();
  $__res = ($k_4)(($f_2)($a_5));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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
}];
  $applyContT1_2_3 = (object)["apply" => function($v_3) {
  $__num = \func_num_args();
  $__res = function($v1_4) use ($v_3) {
  $__num = \func_num_args();
  $__res = function($k_5) use ($v1_4, $v_3) {
  $__num = \func_num_args();
  $__res = ($v_3)(function($g_6) use ($k_5, $v1_4) {
  $__num = \func_num_args();
  $__res = ($v1_4)(function($a_7) use ($g_6, $k_5) {
  $__num = \func_num_args();
  $__res = ($k_5)(($g_6)($a_7));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_3) use ($functorContT1_2_3) {
  $__num = \func_num_args();
  $__res = $functorContT1_2_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $bindContT1_2_3 = (object)["bind" => function($v_3) {
  $__num = \func_num_args();
  $__res = function($k_4) use ($v_3) {
  $__num = \func_num_args();
  $__res = function($k_prime__5) use ($k_4, $v_3) {
  $__num = \func_num_args();
  $__res = ($v_3)(function($a_6) use ($k_4, $k_prime__5) {
  $__num = \func_num_args();
  $__res = (($k_4)($a_6))($k_prime__5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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
}, "Apply0" => function($_dollar___unused_3) use ($applyContT1_2_3) {
  $__num = \func_num_args();
  $__res = $applyContT1_2_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $monadContT1_1_0 = (object)["Applicative0" => function($_dollar___unused_3) use ($applicativeContT1_1_0) {
  $__num = \func_num_args();
  $__res = $applicativeContT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_3) use ($bindContT1_2_3) {
  $__num = \func_num_args();
  $__res = $bindContT1_2_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_2_7 = (((($dictMonadEffect_0)->{'Monad0'})(null))->{'Bind1'})(null);
  $__res = (object)["liftEffect" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($m_3) use ($Bind1_2_7) {
  $__num = \func_num_args();
  $__res = function($k_4) use ($Bind1_2_7, $m_3) {
  $__num = \func_num_args();
  $__res = ((($Bind1_2_7)->{'bind'})($m_3))($k_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($dictMonadEffect_0)->{'liftEffect'}), "Monad0" => function($_dollar___unused_2) use ($monadContT1_1_0) {
  $__num = \func_num_args();
  $__res = $monadContT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Cont_Trans_monadEffectContT'] = __NAMESPACE__ . '\\majControl_majMonad_majCont_majTrans_monadmajEffectmajContmajT';

// Control_Monad_Cont_Trans_monadStateContT
function majControl_majMonad_majCont_majTrans_monadmajStatemajContmajT($dictMonadState_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majCont_majTrans_monadmajStatemajContmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $functorContT1_1_0 = (object)["map" => function($f_1) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($f_1) {
  $__num = \func_num_args();
  $__res = function($k_3) use ($f_1, $v_2) {
  $__num = \func_num_args();
  $__res = ($v_2)(function($a_4) use ($f_1, $k_3) {
  $__num = \func_num_args();
  $__res = ($k_3)(($f_1)($a_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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
}];
  $applyContT1_1_0 = (object)["apply" => function($v_2) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($v_2) {
  $__num = \func_num_args();
  $__res = function($k_4) use ($v1_3, $v_2) {
  $__num = \func_num_args();
  $__res = ($v_2)(function($g_5) use ($k_4, $v1_3) {
  $__num = \func_num_args();
  $__res = ($v1_3)(function($a_6) use ($g_5, $k_4) {
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_2) use ($functorContT1_1_0) {
  $__num = \func_num_args();
  $__res = $functorContT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeContT1_1_0 = (object)["pure" => function($a_2) {
  $__num = \func_num_args();
  $__res = function($k_3) use ($a_2) {
  $__num = \func_num_args();
  $__res = ($k_3)($a_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_2) use ($applyContT1_1_0) {
  $__num = \func_num_args();
  $__res = $applyContT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $functorContT1_2_3 = (object)["map" => function($f_2) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($f_2) {
  $__num = \func_num_args();
  $__res = function($k_4) use ($f_2, $v_3) {
  $__num = \func_num_args();
  $__res = ($v_3)(function($a_5) use ($f_2, $k_4) {
  $__num = \func_num_args();
  $__res = ($k_4)(($f_2)($a_5));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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
}];
  $applyContT1_2_3 = (object)["apply" => function($v_3) {
  $__num = \func_num_args();
  $__res = function($v1_4) use ($v_3) {
  $__num = \func_num_args();
  $__res = function($k_5) use ($v1_4, $v_3) {
  $__num = \func_num_args();
  $__res = ($v_3)(function($g_6) use ($k_5, $v1_4) {
  $__num = \func_num_args();
  $__res = ($v1_4)(function($a_7) use ($g_6, $k_5) {
  $__num = \func_num_args();
  $__res = ($k_5)(($g_6)($a_7));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_3) use ($functorContT1_2_3) {
  $__num = \func_num_args();
  $__res = $functorContT1_2_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $bindContT1_2_3 = (object)["bind" => function($v_3) {
  $__num = \func_num_args();
  $__res = function($k_4) use ($v_3) {
  $__num = \func_num_args();
  $__res = function($k_prime__5) use ($k_4, $v_3) {
  $__num = \func_num_args();
  $__res = ($v_3)(function($a_6) use ($k_4, $k_prime__5) {
  $__num = \func_num_args();
  $__res = (($k_4)($a_6))($k_prime__5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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
}, "Apply0" => function($_dollar___unused_3) use ($applyContT1_2_3) {
  $__num = \func_num_args();
  $__res = $applyContT1_2_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $monadContT1_1_0 = (object)["Applicative0" => function($_dollar___unused_3) use ($applicativeContT1_1_0) {
  $__num = \func_num_args();
  $__res = $applicativeContT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_3) use ($bindContT1_2_3) {
  $__num = \func_num_args();
  $__res = $bindContT1_2_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_2_7 = (((($dictMonadState_0)->{'Monad0'})(null))->{'Bind1'})(null);
  $__res = (object)["state" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($m_3) use ($Bind1_2_7) {
  $__num = \func_num_args();
  $__res = function($k_4) use ($Bind1_2_7, $m_3) {
  $__num = \func_num_args();
  $__res = ((($Bind1_2_7)->{'bind'})($m_3))($k_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($dictMonadState_0)->{'state'}), "Monad0" => function($_dollar___unused_2) use ($monadContT1_1_0) {
  $__num = \func_num_args();
  $__res = $monadContT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Cont_Trans_monadStateContT'] = __NAMESPACE__ . '\\majControl_majMonad_majCont_majTrans_monadmajStatemajContmajT';

// Control_Monad_Cont_Trans_monadSTContT
function majControl_majMonad_majCont_majTrans_monadmajSmajTmajContmajT($dictMonadST_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majCont_majTrans_monadmajSmajTmajContmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $functorContT1_1_0 = (object)["map" => function($f_1) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($f_1) {
  $__num = \func_num_args();
  $__res = function($k_3) use ($f_1, $v_2) {
  $__num = \func_num_args();
  $__res = ($v_2)(function($a_4) use ($f_1, $k_3) {
  $__num = \func_num_args();
  $__res = ($k_3)(($f_1)($a_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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
}];
  $applyContT1_1_0 = (object)["apply" => function($v_2) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($v_2) {
  $__num = \func_num_args();
  $__res = function($k_4) use ($v1_3, $v_2) {
  $__num = \func_num_args();
  $__res = ($v_2)(function($g_5) use ($k_4, $v1_3) {
  $__num = \func_num_args();
  $__res = ($v1_3)(function($a_6) use ($g_5, $k_4) {
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_2) use ($functorContT1_1_0) {
  $__num = \func_num_args();
  $__res = $functorContT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeContT1_1_0 = (object)["pure" => function($a_2) {
  $__num = \func_num_args();
  $__res = function($k_3) use ($a_2) {
  $__num = \func_num_args();
  $__res = ($k_3)($a_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_2) use ($applyContT1_1_0) {
  $__num = \func_num_args();
  $__res = $applyContT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $functorContT1_2_3 = (object)["map" => function($f_2) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($f_2) {
  $__num = \func_num_args();
  $__res = function($k_4) use ($f_2, $v_3) {
  $__num = \func_num_args();
  $__res = ($v_3)(function($a_5) use ($f_2, $k_4) {
  $__num = \func_num_args();
  $__res = ($k_4)(($f_2)($a_5));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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
}];
  $applyContT1_2_3 = (object)["apply" => function($v_3) {
  $__num = \func_num_args();
  $__res = function($v1_4) use ($v_3) {
  $__num = \func_num_args();
  $__res = function($k_5) use ($v1_4, $v_3) {
  $__num = \func_num_args();
  $__res = ($v_3)(function($g_6) use ($k_5, $v1_4) {
  $__num = \func_num_args();
  $__res = ($v1_4)(function($a_7) use ($g_6, $k_5) {
  $__num = \func_num_args();
  $__res = ($k_5)(($g_6)($a_7));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_3) use ($functorContT1_2_3) {
  $__num = \func_num_args();
  $__res = $functorContT1_2_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $bindContT1_2_3 = (object)["bind" => function($v_3) {
  $__num = \func_num_args();
  $__res = function($k_4) use ($v_3) {
  $__num = \func_num_args();
  $__res = function($k_prime__5) use ($k_4, $v_3) {
  $__num = \func_num_args();
  $__res = ($v_3)(function($a_6) use ($k_4, $k_prime__5) {
  $__num = \func_num_args();
  $__res = (($k_4)($a_6))($k_prime__5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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
}, "Apply0" => function($_dollar___unused_3) use ($applyContT1_2_3) {
  $__num = \func_num_args();
  $__res = $applyContT1_2_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $monadContT1_1_0 = (object)["Applicative0" => function($_dollar___unused_3) use ($applicativeContT1_1_0) {
  $__num = \func_num_args();
  $__res = $applicativeContT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_3) use ($bindContT1_2_3) {
  $__num = \func_num_args();
  $__res = $bindContT1_2_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_2_7 = (((($dictMonadST_0)->{'Monad0'})(null))->{'Bind1'})(null);
  $__res = (object)["liftST" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($m_3) use ($Bind1_2_7) {
  $__num = \func_num_args();
  $__res = function($k_4) use ($Bind1_2_7, $m_3) {
  $__num = \func_num_args();
  $__res = ((($Bind1_2_7)->{'bind'})($m_3))($k_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($dictMonadST_0)->{'liftST'}), "Monad0" => function($_dollar___unused_2) use ($monadContT1_1_0) {
  $__num = \func_num_args();
  $__res = $monadContT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Cont_Trans_monadSTContT'] = __NAMESPACE__ . '\\majControl_majMonad_majCont_majTrans_monadmajSmajTmajContmajT';

// Control_Monad_Cont_Trans_monoidContT
function majControl_majMonad_majCont_majTrans_monoidmajContmajT($dictApplicative_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majCont_majTrans_monoidmajContmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $functorContT1_1_0 = (object)["map" => function($f_1) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($f_1) {
  $__num = \func_num_args();
  $__res = function($k_3) use ($f_1, $v_2) {
  $__num = \func_num_args();
  $__res = ($v_2)(function($a_4) use ($f_1, $k_3) {
  $__num = \func_num_args();
  $__res = ($k_3)(($f_1)($a_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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
}];
  $applyContT1_1_0 = (object)["apply" => function($v_2) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($v_2) {
  $__num = \func_num_args();
  $__res = function($k_4) use ($v1_3, $v_2) {
  $__num = \func_num_args();
  $__res = ($v_2)(function($g_5) use ($k_4, $v1_3) {
  $__num = \func_num_args();
  $__res = ($v1_3)(function($a_6) use ($g_5, $k_4) {
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_2) use ($functorContT1_1_0) {
  $__num = \func_num_args();
  $__res = $functorContT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeContT1_1_0 = (object)["pure" => function($a_2) {
  $__num = \func_num_args();
  $__res = function($k_3) use ($a_2) {
  $__num = \func_num_args();
  $__res = ($k_3)($a_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_2) use ($applyContT1_1_0) {
  $__num = \func_num_args();
  $__res = $applyContT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $functorContT1_2_3 = (object)["map" => function($f_2) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($f_2) {
  $__num = \func_num_args();
  $__res = function($k_4) use ($f_2, $v_3) {
  $__num = \func_num_args();
  $__res = ($v_3)(function($a_5) use ($f_2, $k_4) {
  $__num = \func_num_args();
  $__res = ($k_4)(($f_2)($a_5));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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
}];
  $applyContT1_2_3 = (object)["apply" => function($v_3) {
  $__num = \func_num_args();
  $__res = function($v1_4) use ($v_3) {
  $__num = \func_num_args();
  $__res = function($k_5) use ($v1_4, $v_3) {
  $__num = \func_num_args();
  $__res = ($v_3)(function($g_6) use ($k_5, $v1_4) {
  $__num = \func_num_args();
  $__res = ($v1_4)(function($a_7) use ($g_6, $k_5) {
  $__num = \func_num_args();
  $__res = ($k_5)(($g_6)($a_7));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_3) use ($functorContT1_2_3) {
  $__num = \func_num_args();
  $__res = $functorContT1_2_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = function($dictMonoid_3) use ($applicativeContT1_1_0, $applyContT1_2_3) {
  $__num = \func_num_args();
  $Functor0_4_5 = (($applyContT1_2_3)->{'Functor0'})(null);
  $__local_var_5_6 = ((($dictMonoid_3)->{'Semigroup0'})(null))->{'append'};
  $semigroupContT2_4_5 = (object)["append" => function($a_6) use ($Functor0_4_5, $__local_var_5_6, $applyContT1_2_3) {
  $__num = \func_num_args();
  $__res = function($b_7) use ($Functor0_4_5, $__local_var_5_6, $a_6, $applyContT1_2_3) {
  $__num = \func_num_args();
  $__res = ((($applyContT1_2_3)->{'apply'})(((($Functor0_4_5)->{'map'})($__local_var_5_6))($a_6)))($b_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["mempty" => (($applicativeContT1_1_0)->{'pure'})(($dictMonoid_3)->{'mempty'}), "Semigroup0" => function($_dollar___unused_5) use ($semigroupContT2_4_5) {
  $__num = \func_num_args();
  $__res = $semigroupContT2_4_5;
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Cont_Trans_monoidContT'] = __NAMESPACE__ . '\\majControl_majMonad_majCont_majTrans_monoidmajContmajT';

