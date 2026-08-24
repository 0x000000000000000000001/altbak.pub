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
if (!\function_exists(__NAMESPACE__ . '\\phpurs_execute_effect')) {
  function phpurs_execute_effect($val) {
    if (\is_callable($val)) {
      return $val($GLOBALS['Data_Unit_unit']);
    }
    return $val;
  }
}

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };


final class Control_Monad_RWS_Trans_RWSResult { public $tag = 'RWSResult'; public function __construct(public  $value0, public  $value1, public  $value2) {} }

// Control_Monad_RWS_Trans_RWSResult
$GLOBALS['Control_Monad_RWS_Trans_RWSResult'] = (function() {
  $__fn = function($value0, $value1 = null, $value2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult($value0, $value1, $value2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Control_Monad_RWS_Trans_RWST
function majControl_majMonad_majRmajWmajS_majTrans_majRmajWmajSmajT($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majRmajWmajS_majTrans_majRmajWmajSmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_RWS_Trans_RWST'] = __NAMESPACE__ . '\\majControl_majMonad_majRmajWmajS_majTrans_majRmajWmajSmajT';

// Control_Monad_RWS_Trans_withRWST
function majControl_majMonad_majRmajWmajS_majTrans_withmajRmajWmajSmajT($f_0, $m_1 = null, $r_2 = null, $s_3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majRmajWmajS_majTrans_withmajRmajWmajSmajT';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__local_var_4_0 = (($f_0)($r_2))($s_3);
  $__res = (($m_1)(($__local_var_4_0)->{'value0'}))(($__local_var_4_0)->{'value1'});
  goto __end;;
  __end:
  return 4 < $__num ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
}
$GLOBALS['Control_Monad_RWS_Trans_withRWST'] = __NAMESPACE__ . '\\majControl_majMonad_majRmajWmajS_majTrans_withmajRmajWmajSmajT';

// Control_Monad_RWS_Trans_runRWST
function majControl_majMonad_majRmajWmajS_majTrans_runmajRmajWmajSmajT($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majRmajWmajS_majTrans_runmajRmajWmajSmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $v_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_RWS_Trans_runRWST'] = __NAMESPACE__ . '\\majControl_majMonad_majRmajWmajS_majTrans_runmajRmajWmajSmajT';

// Control_Monad_RWS_Trans_newtypeRWST
$GLOBALS['Control_Monad_RWS_Trans_newtypeRWST'] = (object)["Coercible0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Control_Monad_RWS_Trans_monadTransRWST
function majControl_majMonad_majRmajWmajS_majTrans_monadmajTransmajRmajWmajSmajT($dictMonoid_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majRmajWmajS_majTrans_monadmajTransmajRmajWmajSmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["lift" => function($dictMonad_1) use ($dictMonoid_0) {
  $__num = \func_num_args();
  $Bind1_2_0 = (($dictMonad_1)->{'Bind1'})(null);
  $pure_3_1 = ((($dictMonad_1)->{'Applicative0'})(null))->{'pure'};
  $__res = function($m_4) use ($Bind1_2_0, $dictMonoid_0, $pure_3_1) {
  $__num = \func_num_args();
  $__res = function($v_5) use ($Bind1_2_0, $dictMonoid_0, $m_4, $pure_3_1) {
  $__num = \func_num_args();
  $__res = function($s_6) use ($Bind1_2_0, $dictMonoid_0, $m_4, $pure_3_1) {
  $__num = \func_num_args();
  $__res = ((($Bind1_2_0)->{'bind'})($m_4))(function($a_7) use ($dictMonoid_0, $pure_3_1, $s_6) {
  $__num = \func_num_args();
  $__res = ($pure_3_1)(new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult($s_6, $a_7, ($dictMonoid_0)->{'mempty'}));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_RWS_Trans_monadTransRWST'] = __NAMESPACE__ . '\\majControl_majMonad_majRmajWmajS_majTrans_monadmajTransmajRmajWmajSmajT';

// Control_Monad_RWS_Trans_mapRWST
function majControl_majMonad_majRmajWmajS_majTrans_mapmajRmajWmajSmajT($f_0, $v_1 = null, $r_2 = null, $s_3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majRmajWmajS_majTrans_mapmajRmajWmajSmajT';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = ($f_0)((($v_1)($r_2))($s_3));
  goto __end;;
  __end:
  return 4 < $__num ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
}
$GLOBALS['Control_Monad_RWS_Trans_mapRWST'] = __NAMESPACE__ . '\\majControl_majMonad_majRmajWmajS_majTrans_mapmajRmajWmajSmajT';

// Control_Monad_RWS_Trans_lazyRWST
$GLOBALS['Control_Monad_RWS_Trans_lazyRWST'] = (object)["defer" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($r_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = function($s_2) use ($f_0, $r_1) {
  $__num = \func_num_args();
  $__res = ((($f_0)($GLOBALS['Data_Unit_unit']))($r_1))($s_2);
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

// Control_Monad_RWS_Trans_functorRWST
function majControl_majMonad_majRmajWmajS_majTrans_functormajRmajWmajSmajT($dictFunctor_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majRmajWmajS_majTrans_functormajRmajWmajSmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["map" => function($f_1) use ($dictFunctor_0) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($dictFunctor_0, $f_1) {
  $__num = \func_num_args();
  $__res = function($r_3) use ($dictFunctor_0, $f_1, $v_2) {
  $__num = \func_num_args();
  $__res = function($s_4) use ($dictFunctor_0, $f_1, $r_3, $v_2) {
  $__num = \func_num_args();
  $__res = ((($dictFunctor_0)->{'map'})(function($v1_5) use ($f_1) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v1_5)->{'value0'}, ($f_1)(($v1_5)->{'value1'}), ($v1_5)->{'value2'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v_2)($r_3))($s_4));
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_RWS_Trans_functorRWST'] = __NAMESPACE__ . '\\majControl_majMonad_majRmajWmajS_majTrans_functormajRmajWmajSmajT';

// Control_Monad_RWS_Trans_execRWST
function majControl_majMonad_majRmajWmajS_majTrans_execmajRmajWmajSmajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majRmajWmajS_majTrans_execmajRmajWmajSmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Bind1_1_0 = (($dictMonad_0)->{'Bind1'})(null);
  $Applicative0_2_1 = (($dictMonad_0)->{'Applicative0'})(null);
  $__res = function($v_3) use ($Applicative0_2_1, $Bind1_1_0) {
  $__num = \func_num_args();
  $__res = function($r_4) use ($Applicative0_2_1, $Bind1_1_0, $v_3) {
  $__num = \func_num_args();
  $__res = function($s_5) use ($Applicative0_2_1, $Bind1_1_0, $r_4, $v_3) {
  $__num = \func_num_args();
  $__res = ((($Bind1_1_0)->{'bind'})((($v_3)($r_4))($s_5)))(function($v1_6) use ($Applicative0_2_1) {
  $__num = \func_num_args();
  $__res = (($Applicative0_2_1)->{'pure'})(new \Data\Tuple\Data_Tuple_Tuple(($v1_6)->{'value0'}, ($v1_6)->{'value2'}));
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
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_RWS_Trans_execRWST'] = __NAMESPACE__ . '\\majControl_majMonad_majRmajWmajS_majTrans_execmajRmajWmajSmajT';

// Control_Monad_RWS_Trans_evalRWST
function majControl_majMonad_majRmajWmajS_majTrans_evalmajRmajWmajSmajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majRmajWmajS_majTrans_evalmajRmajWmajSmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Bind1_1_0 = (($dictMonad_0)->{'Bind1'})(null);
  $Applicative0_2_1 = (($dictMonad_0)->{'Applicative0'})(null);
  $__res = function($v_3) use ($Applicative0_2_1, $Bind1_1_0) {
  $__num = \func_num_args();
  $__res = function($r_4) use ($Applicative0_2_1, $Bind1_1_0, $v_3) {
  $__num = \func_num_args();
  $__res = function($s_5) use ($Applicative0_2_1, $Bind1_1_0, $r_4, $v_3) {
  $__num = \func_num_args();
  $__res = ((($Bind1_1_0)->{'bind'})((($v_3)($r_4))($s_5)))(function($v1_6) use ($Applicative0_2_1) {
  $__num = \func_num_args();
  $__res = (($Applicative0_2_1)->{'pure'})(new \Data\Tuple\Data_Tuple_Tuple(($v1_6)->{'value1'}, ($v1_6)->{'value2'}));
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
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_RWS_Trans_evalRWST'] = __NAMESPACE__ . '\\majControl_majMonad_majRmajWmajS_majTrans_evalmajRmajWmajSmajT';

// Control_Monad_RWS_Trans_applyRWST
function majControl_majMonad_majRmajWmajS_majTrans_applymajRmajWmajSmajT($dictBind_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majRmajWmajS_majTrans_applymajRmajWmajSmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Apply0_1_0 = (($dictBind_0)->{'Apply0'})(null);
  $Functor0_2_1 = (($Apply0_1_0)->{'Functor0'})(null);
  $__local_var_3_2 = (($Apply0_1_0)->{'Functor0'})(null);
  $functorRWST1_3_2 = (object)["map" => function($f_4) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__res = function($v_5) use ($__local_var_3_2, $f_4) {
  $__num = \func_num_args();
  $__res = function($r_6) use ($__local_var_3_2, $f_4, $v_5) {
  $__num = \func_num_args();
  $__res = function($s_7) use ($__local_var_3_2, $f_4, $r_6, $v_5) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_2)->{'map'})(function($v1_8) use ($f_4) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v1_8)->{'value0'}, ($f_4)(($v1_8)->{'value1'}), ($v1_8)->{'value2'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v_5)($r_6))($s_7));
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = function($dictMonoid_4) use ($Functor0_2_1, $dictBind_0, $functorRWST1_3_2) {
  $__num = \func_num_args();
  $Semigroup0_5_4 = (($dictMonoid_4)->{'Semigroup0'})(null);
  $__res = (object)["apply" => function($v_6) use ($Functor0_2_1, $Semigroup0_5_4, $dictBind_0) {
  $__num = \func_num_args();
  $__res = function($v1_7) use ($Functor0_2_1, $Semigroup0_5_4, $dictBind_0, $v_6) {
  $__num = \func_num_args();
  $__res = function($r_8) use ($Functor0_2_1, $Semigroup0_5_4, $dictBind_0, $v1_7, $v_6) {
  $__num = \func_num_args();
  $__res = function($s_9) use ($Functor0_2_1, $Semigroup0_5_4, $dictBind_0, $r_8, $v1_7, $v_6) {
  $__num = \func_num_args();
  $__res = ((($dictBind_0)->{'bind'})((($v_6)($r_8))($s_9)))(function($v2_10) use ($Functor0_2_1, $Semigroup0_5_4, $r_8, $v1_7) {
  $__num = \func_num_args();
  $__local_var_11_5 = ($v2_10)->{'value2'};
  $__res = ((($Functor0_2_1)->{'map'})(function($v3_12) use ($Semigroup0_5_4, $__local_var_11_5, $v2_10) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v3_12)->{'value0'}, (($v2_10)->{'value1'})(($v3_12)->{'value1'}), ((($Semigroup0_5_4)->{'append'})($__local_var_11_5))(($v3_12)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v1_7)($r_8))(($v2_10)->{'value0'}));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_6) use ($functorRWST1_3_2) {
  $__num = \func_num_args();
  $__res = $functorRWST1_3_2;
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
$GLOBALS['Control_Monad_RWS_Trans_applyRWST'] = __NAMESPACE__ . '\\majControl_majMonad_majRmajWmajS_majTrans_applymajRmajWmajSmajT';

// Control_Monad_RWS_Trans_bindRWST
function majControl_majMonad_majRmajWmajS_majTrans_bindmajRmajWmajSmajT($dictBind_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majRmajWmajS_majTrans_bindmajRmajWmajSmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Functor0_1_0 = (((($dictBind_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $Apply0_2_1 = (($dictBind_0)->{'Apply0'})(null);
  $Functor0_3_2 = (($Apply0_2_1)->{'Functor0'})(null);
  $__local_var_4_3 = (($Apply0_2_1)->{'Functor0'})(null);
  $functorRWST1_4_3 = (object)["map" => function($f_5) use ($__local_var_4_3) {
  $__num = \func_num_args();
  $__res = function($v_6) use ($__local_var_4_3, $f_5) {
  $__num = \func_num_args();
  $__res = function($r_7) use ($__local_var_4_3, $f_5, $v_6) {
  $__num = \func_num_args();
  $__res = function($s_8) use ($__local_var_4_3, $f_5, $r_7, $v_6) {
  $__num = \func_num_args();
  $__res = ((($__local_var_4_3)->{'map'})(function($v1_9) use ($f_5) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v1_9)->{'value0'}, ($f_5)(($v1_9)->{'value1'}), ($v1_9)->{'value2'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v_6)($r_7))($s_8));
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = function($dictMonoid_5) use ($Functor0_1_0, $Functor0_3_2, $dictBind_0, $functorRWST1_4_3) {
  $__num = \func_num_args();
  $Semigroup0_6_5 = (($dictMonoid_5)->{'Semigroup0'})(null);
  $Semigroup0_7_6 = (($dictMonoid_5)->{'Semigroup0'})(null);
  $applyRWST2_7_6 = (object)["apply" => function($v_8) use ($Functor0_3_2, $Semigroup0_7_6, $dictBind_0) {
  $__num = \func_num_args();
  $__res = function($v1_9) use ($Functor0_3_2, $Semigroup0_7_6, $dictBind_0, $v_8) {
  $__num = \func_num_args();
  $__res = function($r_10) use ($Functor0_3_2, $Semigroup0_7_6, $dictBind_0, $v1_9, $v_8) {
  $__num = \func_num_args();
  $__res = function($s_11) use ($Functor0_3_2, $Semigroup0_7_6, $dictBind_0, $r_10, $v1_9, $v_8) {
  $__num = \func_num_args();
  $__res = ((($dictBind_0)->{'bind'})((($v_8)($r_10))($s_11)))(function($v2_12) use ($Functor0_3_2, $Semigroup0_7_6, $r_10, $v1_9) {
  $__num = \func_num_args();
  $__local_var_13_7 = ($v2_12)->{'value2'};
  $__res = ((($Functor0_3_2)->{'map'})(function($v3_14) use ($Semigroup0_7_6, $__local_var_13_7, $v2_12) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v3_14)->{'value0'}, (($v2_12)->{'value1'})(($v3_14)->{'value1'}), ((($Semigroup0_7_6)->{'append'})($__local_var_13_7))(($v3_14)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v1_9)($r_10))(($v2_12)->{'value0'}));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_8) use ($functorRWST1_4_3) {
  $__num = \func_num_args();
  $__res = $functorRWST1_4_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["bind" => function($v_8) use ($Functor0_1_0, $Semigroup0_6_5, $dictBind_0) {
  $__num = \func_num_args();
  $__res = function($f_9) use ($Functor0_1_0, $Semigroup0_6_5, $dictBind_0, $v_8) {
  $__num = \func_num_args();
  $__res = function($r_10) use ($Functor0_1_0, $Semigroup0_6_5, $dictBind_0, $f_9, $v_8) {
  $__num = \func_num_args();
  $__res = function($s_11) use ($Functor0_1_0, $Semigroup0_6_5, $dictBind_0, $f_9, $r_10, $v_8) {
  $__num = \func_num_args();
  $__res = ((($dictBind_0)->{'bind'})((($v_8)($r_10))($s_11)))(function($v1_12) use ($Functor0_1_0, $Semigroup0_6_5, $f_9, $r_10) {
  $__num = \func_num_args();
  $__local_var_13_9 = ($v1_12)->{'value2'};
  $__res = ((($Functor0_1_0)->{'map'})(function($v3_14) use ($Semigroup0_6_5, $__local_var_13_9) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v3_14)->{'value0'}, ($v3_14)->{'value1'}, ((($Semigroup0_6_5)->{'append'})($__local_var_13_9))(($v3_14)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($f_9)(($v1_12)->{'value1'}))($r_10))(($v1_12)->{'value0'}));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_8) use ($applyRWST2_7_6) {
  $__num = \func_num_args();
  $__res = $applyRWST2_7_6;
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
$GLOBALS['Control_Monad_RWS_Trans_bindRWST'] = __NAMESPACE__ . '\\majControl_majMonad_majRmajWmajS_majTrans_bindmajRmajWmajSmajT';

// Control_Monad_RWS_Trans_semigroupRWST
function majControl_majMonad_majRmajWmajS_majTrans_semigroupmajRmajWmajSmajT($dictBind_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majRmajWmajS_majTrans_semigroupmajRmajWmajSmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Apply0_1_0 = (($dictBind_0)->{'Apply0'})(null);
  $Functor0_2_1 = (($Apply0_1_0)->{'Functor0'})(null);
  $__local_var_3_2 = (($Apply0_1_0)->{'Functor0'})(null);
  $functorRWST1_3_2 = (object)["map" => function($f_4) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__res = function($v_5) use ($__local_var_3_2, $f_4) {
  $__num = \func_num_args();
  $__res = function($r_6) use ($__local_var_3_2, $f_4, $v_5) {
  $__num = \func_num_args();
  $__res = function($s_7) use ($__local_var_3_2, $f_4, $r_6, $v_5) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_2)->{'map'})(function($v1_8) use ($f_4) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v1_8)->{'value0'}, ($f_4)(($v1_8)->{'value1'}), ($v1_8)->{'value2'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v_5)($r_6))($s_7));
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyRWST1_3_2 = function($dictMonoid_4) use ($Functor0_2_1, $dictBind_0, $functorRWST1_3_2) {
  $__num = \func_num_args();
  $Semigroup0_5_4 = (($dictMonoid_4)->{'Semigroup0'})(null);
  $__res = (object)["apply" => function($v_6) use ($Functor0_2_1, $Semigroup0_5_4, $dictBind_0) {
  $__num = \func_num_args();
  $__res = function($v1_7) use ($Functor0_2_1, $Semigroup0_5_4, $dictBind_0, $v_6) {
  $__num = \func_num_args();
  $__res = function($r_8) use ($Functor0_2_1, $Semigroup0_5_4, $dictBind_0, $v1_7, $v_6) {
  $__num = \func_num_args();
  $__res = function($s_9) use ($Functor0_2_1, $Semigroup0_5_4, $dictBind_0, $r_8, $v1_7, $v_6) {
  $__num = \func_num_args();
  $__res = ((($dictBind_0)->{'bind'})((($v_6)($r_8))($s_9)))(function($v2_10) use ($Functor0_2_1, $Semigroup0_5_4, $r_8, $v1_7) {
  $__num = \func_num_args();
  $__local_var_11_5 = ($v2_10)->{'value2'};
  $__res = ((($Functor0_2_1)->{'map'})(function($v3_12) use ($Semigroup0_5_4, $__local_var_11_5, $v2_10) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v3_12)->{'value0'}, (($v2_10)->{'value1'})(($v3_12)->{'value1'}), ((($Semigroup0_5_4)->{'append'})($__local_var_11_5))(($v3_12)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v1_7)($r_8))(($v2_10)->{'value0'}));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_6) use ($functorRWST1_3_2) {
  $__num = \func_num_args();
  $__res = $functorRWST1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = function($dictMonoid_4) use ($applyRWST1_3_2) {
  $__num = \func_num_args();
  $applyRWST2_5_7 = ($applyRWST1_3_2)($dictMonoid_4);
  $__res = function($dictSemigroup_6) use ($applyRWST2_5_7) {
  $__num = \func_num_args();
  $Functor0_7_8 = (($applyRWST2_5_7)->{'Functor0'})(null);
  $__local_var_8_9 = ($dictSemigroup_6)->{'append'};
  $__res = (object)["append" => function($a_9) use ($Functor0_7_8, $__local_var_8_9, $applyRWST2_5_7) {
  $__num = \func_num_args();
  $__res = function($b_10) use ($Functor0_7_8, $__local_var_8_9, $a_9, $applyRWST2_5_7) {
  $__num = \func_num_args();
  $__res = ((($applyRWST2_5_7)->{'apply'})(((($Functor0_7_8)->{'map'})($__local_var_8_9))($a_9)))($b_10);
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_RWS_Trans_semigroupRWST'] = __NAMESPACE__ . '\\majControl_majMonad_majRmajWmajS_majTrans_semigroupmajRmajWmajSmajT';

// Control_Monad_RWS_Trans_applicativeRWST
function majControl_majMonad_majRmajWmajS_majTrans_applicativemajRmajWmajSmajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majRmajWmajS_majTrans_applicativemajRmajWmajSmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $pure_1_0 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $__local_var_2_1 = (($dictMonad_0)->{'Bind1'})(null);
  $Apply0_3_2 = (($__local_var_2_1)->{'Apply0'})(null);
  $Functor0_4_3 = (($Apply0_3_2)->{'Functor0'})(null);
  $__local_var_5_4 = (($Apply0_3_2)->{'Functor0'})(null);
  $functorRWST1_5_4 = (object)["map" => function($f_6) use ($__local_var_5_4) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($__local_var_5_4, $f_6) {
  $__num = \func_num_args();
  $__res = function($r_8) use ($__local_var_5_4, $f_6, $v_7) {
  $__num = \func_num_args();
  $__res = function($s_9) use ($__local_var_5_4, $f_6, $r_8, $v_7) {
  $__num = \func_num_args();
  $__res = ((($__local_var_5_4)->{'map'})(function($v1_10) use ($f_6) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v1_10)->{'value0'}, ($f_6)(($v1_10)->{'value1'}), ($v1_10)->{'value2'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v_7)($r_8))($s_9));
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyRWST1_4_3 = function($dictMonoid_6) use ($Functor0_4_3, $__local_var_2_1, $functorRWST1_5_4) {
  $__num = \func_num_args();
  $Semigroup0_7_6 = (($dictMonoid_6)->{'Semigroup0'})(null);
  $__res = (object)["apply" => function($v_8) use ($Functor0_4_3, $Semigroup0_7_6, $__local_var_2_1) {
  $__num = \func_num_args();
  $__res = function($v1_9) use ($Functor0_4_3, $Semigroup0_7_6, $__local_var_2_1, $v_8) {
  $__num = \func_num_args();
  $__res = function($r_10) use ($Functor0_4_3, $Semigroup0_7_6, $__local_var_2_1, $v1_9, $v_8) {
  $__num = \func_num_args();
  $__res = function($s_11) use ($Functor0_4_3, $Semigroup0_7_6, $__local_var_2_1, $r_10, $v1_9, $v_8) {
  $__num = \func_num_args();
  $__res = ((($__local_var_2_1)->{'bind'})((($v_8)($r_10))($s_11)))(function($v2_12) use ($Functor0_4_3, $Semigroup0_7_6, $r_10, $v1_9) {
  $__num = \func_num_args();
  $__local_var_13_7 = ($v2_12)->{'value2'};
  $__res = ((($Functor0_4_3)->{'map'})(function($v3_14) use ($Semigroup0_7_6, $__local_var_13_7, $v2_12) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v3_14)->{'value0'}, (($v2_12)->{'value1'})(($v3_14)->{'value1'}), ((($Semigroup0_7_6)->{'append'})($__local_var_13_7))(($v3_14)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v1_9)($r_10))(($v2_12)->{'value0'}));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_8) use ($functorRWST1_5_4) {
  $__num = \func_num_args();
  $__res = $functorRWST1_5_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = function($dictMonoid_5) use ($applyRWST1_4_3, $pure_1_0) {
  $__num = \func_num_args();
  $applyRWST2_6_9 = ($applyRWST1_4_3)($dictMonoid_5);
  $__res = (object)["pure" => function($a_7) use ($dictMonoid_5, $pure_1_0) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($a_7, $dictMonoid_5, $pure_1_0) {
  $__num = \func_num_args();
  $__res = function($s_9) use ($a_7, $dictMonoid_5, $pure_1_0) {
  $__num = \func_num_args();
  $__res = ($pure_1_0)(new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult($s_9, $a_7, ($dictMonoid_5)->{'mempty'}));
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
}, "Apply0" => function($_dollar___unused_7) use ($applyRWST2_6_9) {
  $__num = \func_num_args();
  $__res = $applyRWST2_6_9;
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
$GLOBALS['Control_Monad_RWS_Trans_applicativeRWST'] = __NAMESPACE__ . '\\majControl_majMonad_majRmajWmajS_majTrans_applicativemajRmajWmajSmajT';

// Control_Monad_RWS_Trans_monadRWST
function majControl_majMonad_majRmajWmajS_majTrans_monadmajRmajWmajSmajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majRmajWmajS_majTrans_monadmajRmajWmajSmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $pure_1_0 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $__local_var_2_1 = (($dictMonad_0)->{'Bind1'})(null);
  $Apply0_3_2 = (($__local_var_2_1)->{'Apply0'})(null);
  $Functor0_4_3 = (($Apply0_3_2)->{'Functor0'})(null);
  $__local_var_5_4 = (($Apply0_3_2)->{'Functor0'})(null);
  $functorRWST1_5_4 = (object)["map" => function($f_6) use ($__local_var_5_4) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($__local_var_5_4, $f_6) {
  $__num = \func_num_args();
  $__res = function($r_8) use ($__local_var_5_4, $f_6, $v_7) {
  $__num = \func_num_args();
  $__res = function($s_9) use ($__local_var_5_4, $f_6, $r_8, $v_7) {
  $__num = \func_num_args();
  $__res = ((($__local_var_5_4)->{'map'})(function($v1_10) use ($f_6) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v1_10)->{'value0'}, ($f_6)(($v1_10)->{'value1'}), ($v1_10)->{'value2'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v_7)($r_8))($s_9));
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeRWST1_5_4 = function($dictMonoid_6) use ($Functor0_4_3, $__local_var_2_1, $functorRWST1_5_4, $pure_1_0) {
  $__num = \func_num_args();
  $Semigroup0_7_6 = (($dictMonoid_6)->{'Semigroup0'})(null);
  $applyRWST2_7_6 = (object)["apply" => function($v_8) use ($Functor0_4_3, $Semigroup0_7_6, $__local_var_2_1) {
  $__num = \func_num_args();
  $__res = function($v1_9) use ($Functor0_4_3, $Semigroup0_7_6, $__local_var_2_1, $v_8) {
  $__num = \func_num_args();
  $__res = function($r_10) use ($Functor0_4_3, $Semigroup0_7_6, $__local_var_2_1, $v1_9, $v_8) {
  $__num = \func_num_args();
  $__res = function($s_11) use ($Functor0_4_3, $Semigroup0_7_6, $__local_var_2_1, $r_10, $v1_9, $v_8) {
  $__num = \func_num_args();
  $__res = ((($__local_var_2_1)->{'bind'})((($v_8)($r_10))($s_11)))(function($v2_12) use ($Functor0_4_3, $Semigroup0_7_6, $r_10, $v1_9) {
  $__num = \func_num_args();
  $__local_var_13_7 = ($v2_12)->{'value2'};
  $__res = ((($Functor0_4_3)->{'map'})(function($v3_14) use ($Semigroup0_7_6, $__local_var_13_7, $v2_12) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v3_14)->{'value0'}, (($v2_12)->{'value1'})(($v3_14)->{'value1'}), ((($Semigroup0_7_6)->{'append'})($__local_var_13_7))(($v3_14)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v1_9)($r_10))(($v2_12)->{'value0'}));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_8) use ($functorRWST1_5_4) {
  $__num = \func_num_args();
  $__res = $functorRWST1_5_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["pure" => function($a_8) use ($dictMonoid_6, $pure_1_0) {
  $__num = \func_num_args();
  $__res = function($v_9) use ($a_8, $dictMonoid_6, $pure_1_0) {
  $__num = \func_num_args();
  $__res = function($s_10) use ($a_8, $dictMonoid_6, $pure_1_0) {
  $__num = \func_num_args();
  $__res = ($pure_1_0)(new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult($s_10, $a_8, ($dictMonoid_6)->{'mempty'}));
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
}, "Apply0" => function($_dollar___unused_8) use ($applyRWST2_7_6) {
  $__num = \func_num_args();
  $__res = $applyRWST2_7_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__local_var_6_10 = (($dictMonad_0)->{'Bind1'})(null);
  $Functor0_7_11 = (((($__local_var_6_10)->{'Apply0'})(null))->{'Functor0'})(null);
  $Apply0_8_12 = (($__local_var_6_10)->{'Apply0'})(null);
  $Functor0_9_13 = (($Apply0_8_12)->{'Functor0'})(null);
  $__local_var_10_14 = (($Apply0_8_12)->{'Functor0'})(null);
  $functorRWST1_10_14 = (object)["map" => function($f_11) use ($__local_var_10_14) {
  $__num = \func_num_args();
  $__res = function($v_12) use ($__local_var_10_14, $f_11) {
  $__num = \func_num_args();
  $__res = function($r_13) use ($__local_var_10_14, $f_11, $v_12) {
  $__num = \func_num_args();
  $__res = function($s_14) use ($__local_var_10_14, $f_11, $r_13, $v_12) {
  $__num = \func_num_args();
  $__res = ((($__local_var_10_14)->{'map'})(function($v1_15) use ($f_11) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v1_15)->{'value0'}, ($f_11)(($v1_15)->{'value1'}), ($v1_15)->{'value2'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v_12)($r_13))($s_14));
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $bindRWST1_10_14 = function($dictMonoid_11) use ($Functor0_7_11, $Functor0_9_13, $__local_var_6_10, $functorRWST1_10_14) {
  $__num = \func_num_args();
  $Semigroup0_12_16 = (($dictMonoid_11)->{'Semigroup0'})(null);
  $Semigroup0_13_17 = (($dictMonoid_11)->{'Semigroup0'})(null);
  $applyRWST2_13_17 = (object)["apply" => function($v_14) use ($Functor0_9_13, $Semigroup0_13_17, $__local_var_6_10) {
  $__num = \func_num_args();
  $__res = function($v1_15) use ($Functor0_9_13, $Semigroup0_13_17, $__local_var_6_10, $v_14) {
  $__num = \func_num_args();
  $__res = function($r_16) use ($Functor0_9_13, $Semigroup0_13_17, $__local_var_6_10, $v1_15, $v_14) {
  $__num = \func_num_args();
  $__res = function($s_17) use ($Functor0_9_13, $Semigroup0_13_17, $__local_var_6_10, $r_16, $v1_15, $v_14) {
  $__num = \func_num_args();
  $__res = ((($__local_var_6_10)->{'bind'})((($v_14)($r_16))($s_17)))(function($v2_18) use ($Functor0_9_13, $Semigroup0_13_17, $r_16, $v1_15) {
  $__num = \func_num_args();
  $__local_var_19_18 = ($v2_18)->{'value2'};
  $__res = ((($Functor0_9_13)->{'map'})(function($v3_20) use ($Semigroup0_13_17, $__local_var_19_18, $v2_18) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v3_20)->{'value0'}, (($v2_18)->{'value1'})(($v3_20)->{'value1'}), ((($Semigroup0_13_17)->{'append'})($__local_var_19_18))(($v3_20)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v1_15)($r_16))(($v2_18)->{'value0'}));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_14) use ($functorRWST1_10_14) {
  $__num = \func_num_args();
  $__res = $functorRWST1_10_14;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["bind" => function($v_14) use ($Functor0_7_11, $Semigroup0_12_16, $__local_var_6_10) {
  $__num = \func_num_args();
  $__res = function($f_15) use ($Functor0_7_11, $Semigroup0_12_16, $__local_var_6_10, $v_14) {
  $__num = \func_num_args();
  $__res = function($r_16) use ($Functor0_7_11, $Semigroup0_12_16, $__local_var_6_10, $f_15, $v_14) {
  $__num = \func_num_args();
  $__res = function($s_17) use ($Functor0_7_11, $Semigroup0_12_16, $__local_var_6_10, $f_15, $r_16, $v_14) {
  $__num = \func_num_args();
  $__res = ((($__local_var_6_10)->{'bind'})((($v_14)($r_16))($s_17)))(function($v1_18) use ($Functor0_7_11, $Semigroup0_12_16, $f_15, $r_16) {
  $__num = \func_num_args();
  $__local_var_19_20 = ($v1_18)->{'value2'};
  $__res = ((($Functor0_7_11)->{'map'})(function($v3_20) use ($Semigroup0_12_16, $__local_var_19_20) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v3_20)->{'value0'}, ($v3_20)->{'value1'}, ((($Semigroup0_12_16)->{'append'})($__local_var_19_20))(($v3_20)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($f_15)(($v1_18)->{'value1'}))($r_16))(($v1_18)->{'value0'}));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_14) use ($applyRWST2_13_17) {
  $__num = \func_num_args();
  $__res = $applyRWST2_13_17;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = function($dictMonoid_11) use ($applicativeRWST1_5_4, $bindRWST1_10_14) {
  $__num = \func_num_args();
  $applicativeRWST2_12_22 = ($applicativeRWST1_5_4)($dictMonoid_11);
  $bindRWST2_13_23 = ($bindRWST1_10_14)($dictMonoid_11);
  $__res = (object)["Applicative0" => function($_dollar___unused_14) use ($applicativeRWST2_12_22) {
  $__num = \func_num_args();
  $__res = $applicativeRWST2_12_22;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_14) use ($bindRWST2_13_23) {
  $__num = \func_num_args();
  $__res = $bindRWST2_13_23;
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
$GLOBALS['Control_Monad_RWS_Trans_monadRWST'] = __NAMESPACE__ . '\\majControl_majMonad_majRmajWmajS_majTrans_monadmajRmajWmajSmajT';

// Control_Monad_RWS_Trans_monadAskRWST
function majControl_majMonad_majRmajWmajS_majTrans_monadmajAskmajRmajWmajSmajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majRmajWmajS_majTrans_monadmajAskmajRmajWmajSmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $pure_1_0 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $pure_2_1 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $__local_var_3_2 = (($dictMonad_0)->{'Bind1'})(null);
  $Apply0_4_3 = (($__local_var_3_2)->{'Apply0'})(null);
  $Functor0_5_4 = (($Apply0_4_3)->{'Functor0'})(null);
  $__local_var_6_5 = (($Apply0_4_3)->{'Functor0'})(null);
  $functorRWST1_6_5 = (object)["map" => function($f_7) use ($__local_var_6_5) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($__local_var_6_5, $f_7) {
  $__num = \func_num_args();
  $__res = function($r_9) use ($__local_var_6_5, $f_7, $v_8) {
  $__num = \func_num_args();
  $__res = function($s_10) use ($__local_var_6_5, $f_7, $r_9, $v_8) {
  $__num = \func_num_args();
  $__res = ((($__local_var_6_5)->{'map'})(function($v1_11) use ($f_7) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v1_11)->{'value0'}, ($f_7)(($v1_11)->{'value1'}), ($v1_11)->{'value2'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v_8)($r_9))($s_10));
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_7_7 = (($dictMonad_0)->{'Bind1'})(null);
  $Functor0_8_8 = (((($__local_var_7_7)->{'Apply0'})(null))->{'Functor0'})(null);
  $Apply0_9_9 = (($__local_var_7_7)->{'Apply0'})(null);
  $Functor0_10_10 = (($Apply0_9_9)->{'Functor0'})(null);
  $__local_var_11_11 = (($Apply0_9_9)->{'Functor0'})(null);
  $functorRWST1_11_11 = (object)["map" => function($f_12) use ($__local_var_11_11) {
  $__num = \func_num_args();
  $__res = function($v_13) use ($__local_var_11_11, $f_12) {
  $__num = \func_num_args();
  $__res = function($r_14) use ($__local_var_11_11, $f_12, $v_13) {
  $__num = \func_num_args();
  $__res = function($s_15) use ($__local_var_11_11, $f_12, $r_14, $v_13) {
  $__num = \func_num_args();
  $__res = ((($__local_var_11_11)->{'map'})(function($v1_16) use ($f_12) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v1_16)->{'value0'}, ($f_12)(($v1_16)->{'value1'}), ($v1_16)->{'value2'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v_13)($r_14))($s_15));
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $monadRWST1_5_4 = function($dictMonoid_12) use ($Functor0_10_10, $Functor0_5_4, $Functor0_8_8, $__local_var_3_2, $__local_var_7_7, $functorRWST1_11_11, $functorRWST1_6_5, $pure_2_1) {
  $__num = \func_num_args();
  $Semigroup0_13_13 = (($dictMonoid_12)->{'Semigroup0'})(null);
  $applyRWST2_13_13 = (object)["apply" => function($v_14) use ($Functor0_5_4, $Semigroup0_13_13, $__local_var_3_2) {
  $__num = \func_num_args();
  $__res = function($v1_15) use ($Functor0_5_4, $Semigroup0_13_13, $__local_var_3_2, $v_14) {
  $__num = \func_num_args();
  $__res = function($r_16) use ($Functor0_5_4, $Semigroup0_13_13, $__local_var_3_2, $v1_15, $v_14) {
  $__num = \func_num_args();
  $__res = function($s_17) use ($Functor0_5_4, $Semigroup0_13_13, $__local_var_3_2, $r_16, $v1_15, $v_14) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_2)->{'bind'})((($v_14)($r_16))($s_17)))(function($v2_18) use ($Functor0_5_4, $Semigroup0_13_13, $r_16, $v1_15) {
  $__num = \func_num_args();
  $__local_var_19_14 = ($v2_18)->{'value2'};
  $__res = ((($Functor0_5_4)->{'map'})(function($v3_20) use ($Semigroup0_13_13, $__local_var_19_14, $v2_18) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v3_20)->{'value0'}, (($v2_18)->{'value1'})(($v3_20)->{'value1'}), ((($Semigroup0_13_13)->{'append'})($__local_var_19_14))(($v3_20)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v1_15)($r_16))(($v2_18)->{'value0'}));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_14) use ($functorRWST1_6_5) {
  $__num = \func_num_args();
  $__res = $functorRWST1_6_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeRWST2_13_13 = (object)["pure" => function($a_14) use ($dictMonoid_12, $pure_2_1) {
  $__num = \func_num_args();
  $__res = function($v_15) use ($a_14, $dictMonoid_12, $pure_2_1) {
  $__num = \func_num_args();
  $__res = function($s_16) use ($a_14, $dictMonoid_12, $pure_2_1) {
  $__num = \func_num_args();
  $__res = ($pure_2_1)(new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult($s_16, $a_14, ($dictMonoid_12)->{'mempty'}));
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
}, "Apply0" => function($_dollar___unused_14) use ($applyRWST2_13_13) {
  $__num = \func_num_args();
  $__res = $applyRWST2_13_13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Semigroup0_14_17 = (($dictMonoid_12)->{'Semigroup0'})(null);
  $Semigroup0_15_18 = (($dictMonoid_12)->{'Semigroup0'})(null);
  $applyRWST2_15_18 = (object)["apply" => function($v_16) use ($Functor0_10_10, $Semigroup0_15_18, $__local_var_7_7) {
  $__num = \func_num_args();
  $__res = function($v1_17) use ($Functor0_10_10, $Semigroup0_15_18, $__local_var_7_7, $v_16) {
  $__num = \func_num_args();
  $__res = function($r_18) use ($Functor0_10_10, $Semigroup0_15_18, $__local_var_7_7, $v1_17, $v_16) {
  $__num = \func_num_args();
  $__res = function($s_19) use ($Functor0_10_10, $Semigroup0_15_18, $__local_var_7_7, $r_18, $v1_17, $v_16) {
  $__num = \func_num_args();
  $__res = ((($__local_var_7_7)->{'bind'})((($v_16)($r_18))($s_19)))(function($v2_20) use ($Functor0_10_10, $Semigroup0_15_18, $r_18, $v1_17) {
  $__num = \func_num_args();
  $__local_var_21_19 = ($v2_20)->{'value2'};
  $__res = ((($Functor0_10_10)->{'map'})(function($v3_22) use ($Semigroup0_15_18, $__local_var_21_19, $v2_20) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v3_22)->{'value0'}, (($v2_20)->{'value1'})(($v3_22)->{'value1'}), ((($Semigroup0_15_18)->{'append'})($__local_var_21_19))(($v3_22)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v1_17)($r_18))(($v2_20)->{'value0'}));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_16) use ($functorRWST1_11_11) {
  $__num = \func_num_args();
  $__res = $functorRWST1_11_11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $bindRWST2_14_17 = (object)["bind" => function($v_16) use ($Functor0_8_8, $Semigroup0_14_17, $__local_var_7_7) {
  $__num = \func_num_args();
  $__res = function($f_17) use ($Functor0_8_8, $Semigroup0_14_17, $__local_var_7_7, $v_16) {
  $__num = \func_num_args();
  $__res = function($r_18) use ($Functor0_8_8, $Semigroup0_14_17, $__local_var_7_7, $f_17, $v_16) {
  $__num = \func_num_args();
  $__res = function($s_19) use ($Functor0_8_8, $Semigroup0_14_17, $__local_var_7_7, $f_17, $r_18, $v_16) {
  $__num = \func_num_args();
  $__res = ((($__local_var_7_7)->{'bind'})((($v_16)($r_18))($s_19)))(function($v1_20) use ($Functor0_8_8, $Semigroup0_14_17, $f_17, $r_18) {
  $__num = \func_num_args();
  $__local_var_21_21 = ($v1_20)->{'value2'};
  $__res = ((($Functor0_8_8)->{'map'})(function($v3_22) use ($Semigroup0_14_17, $__local_var_21_21) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v3_22)->{'value0'}, ($v3_22)->{'value1'}, ((($Semigroup0_14_17)->{'append'})($__local_var_21_21))(($v3_22)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($f_17)(($v1_20)->{'value1'}))($r_18))(($v1_20)->{'value0'}));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_16) use ($applyRWST2_15_18) {
  $__num = \func_num_args();
  $__res = $applyRWST2_15_18;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["Applicative0" => function($_dollar___unused_15) use ($applicativeRWST2_13_13) {
  $__num = \func_num_args();
  $__res = $applicativeRWST2_13_13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_15) use ($bindRWST2_14_17) {
  $__num = \func_num_args();
  $__res = $bindRWST2_14_17;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = function($dictMonoid_6) use ($monadRWST1_5_4, $pure_1_0) {
  $__num = \func_num_args();
  $monadRWST2_7_24 = ($monadRWST1_5_4)($dictMonoid_6);
  $__res = (object)["ask" => function($r_8) use ($dictMonoid_6, $pure_1_0) {
  $__num = \func_num_args();
  $__res = function($s_9) use ($dictMonoid_6, $pure_1_0, $r_8) {
  $__num = \func_num_args();
  $__res = ($pure_1_0)(new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult($s_9, $r_8, ($dictMonoid_6)->{'mempty'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad0" => function($_dollar___unused_8) use ($monadRWST2_7_24) {
  $__num = \func_num_args();
  $__res = $monadRWST2_7_24;
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
$GLOBALS['Control_Monad_RWS_Trans_monadAskRWST'] = __NAMESPACE__ . '\\majControl_majMonad_majRmajWmajS_majTrans_monadmajAskmajRmajWmajSmajT';

// Control_Monad_RWS_Trans_monadReaderRWST
function majControl_majMonad_majRmajWmajS_majTrans_monadmajReadermajRmajWmajSmajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majRmajWmajS_majTrans_monadmajReadermajRmajWmajSmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $pure_1_0 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $pure_2_1 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $__local_var_3_2 = (($dictMonad_0)->{'Bind1'})(null);
  $Apply0_4_3 = (($__local_var_3_2)->{'Apply0'})(null);
  $Functor0_5_4 = (($Apply0_4_3)->{'Functor0'})(null);
  $__local_var_6_5 = (($Apply0_4_3)->{'Functor0'})(null);
  $functorRWST1_6_5 = (object)["map" => function($f_7) use ($__local_var_6_5) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($__local_var_6_5, $f_7) {
  $__num = \func_num_args();
  $__res = function($r_9) use ($__local_var_6_5, $f_7, $v_8) {
  $__num = \func_num_args();
  $__res = function($s_10) use ($__local_var_6_5, $f_7, $r_9, $v_8) {
  $__num = \func_num_args();
  $__res = ((($__local_var_6_5)->{'map'})(function($v1_11) use ($f_7) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v1_11)->{'value0'}, ($f_7)(($v1_11)->{'value1'}), ($v1_11)->{'value2'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v_8)($r_9))($s_10));
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_7_7 = (($dictMonad_0)->{'Bind1'})(null);
  $Functor0_8_8 = (((($__local_var_7_7)->{'Apply0'})(null))->{'Functor0'})(null);
  $Apply0_9_9 = (($__local_var_7_7)->{'Apply0'})(null);
  $Functor0_10_10 = (($Apply0_9_9)->{'Functor0'})(null);
  $__local_var_11_11 = (($Apply0_9_9)->{'Functor0'})(null);
  $functorRWST1_11_11 = (object)["map" => function($f_12) use ($__local_var_11_11) {
  $__num = \func_num_args();
  $__res = function($v_13) use ($__local_var_11_11, $f_12) {
  $__num = \func_num_args();
  $__res = function($r_14) use ($__local_var_11_11, $f_12, $v_13) {
  $__num = \func_num_args();
  $__res = function($s_15) use ($__local_var_11_11, $f_12, $r_14, $v_13) {
  $__num = \func_num_args();
  $__res = ((($__local_var_11_11)->{'map'})(function($v1_16) use ($f_12) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v1_16)->{'value0'}, ($f_12)(($v1_16)->{'value1'}), ($v1_16)->{'value2'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v_13)($r_14))($s_15));
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $monadRWST1_7_7 = function($dictMonoid_12) use ($Functor0_10_10, $Functor0_5_4, $Functor0_8_8, $__local_var_3_2, $__local_var_7_7, $functorRWST1_11_11, $functorRWST1_6_5, $pure_2_1) {
  $__num = \func_num_args();
  $Semigroup0_13_13 = (($dictMonoid_12)->{'Semigroup0'})(null);
  $applyRWST2_13_13 = (object)["apply" => function($v_14) use ($Functor0_5_4, $Semigroup0_13_13, $__local_var_3_2) {
  $__num = \func_num_args();
  $__res = function($v1_15) use ($Functor0_5_4, $Semigroup0_13_13, $__local_var_3_2, $v_14) {
  $__num = \func_num_args();
  $__res = function($r_16) use ($Functor0_5_4, $Semigroup0_13_13, $__local_var_3_2, $v1_15, $v_14) {
  $__num = \func_num_args();
  $__res = function($s_17) use ($Functor0_5_4, $Semigroup0_13_13, $__local_var_3_2, $r_16, $v1_15, $v_14) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_2)->{'bind'})((($v_14)($r_16))($s_17)))(function($v2_18) use ($Functor0_5_4, $Semigroup0_13_13, $r_16, $v1_15) {
  $__num = \func_num_args();
  $__local_var_19_14 = ($v2_18)->{'value2'};
  $__res = ((($Functor0_5_4)->{'map'})(function($v3_20) use ($Semigroup0_13_13, $__local_var_19_14, $v2_18) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v3_20)->{'value0'}, (($v2_18)->{'value1'})(($v3_20)->{'value1'}), ((($Semigroup0_13_13)->{'append'})($__local_var_19_14))(($v3_20)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v1_15)($r_16))(($v2_18)->{'value0'}));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_14) use ($functorRWST1_6_5) {
  $__num = \func_num_args();
  $__res = $functorRWST1_6_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeRWST2_13_13 = (object)["pure" => function($a_14) use ($dictMonoid_12, $pure_2_1) {
  $__num = \func_num_args();
  $__res = function($v_15) use ($a_14, $dictMonoid_12, $pure_2_1) {
  $__num = \func_num_args();
  $__res = function($s_16) use ($a_14, $dictMonoid_12, $pure_2_1) {
  $__num = \func_num_args();
  $__res = ($pure_2_1)(new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult($s_16, $a_14, ($dictMonoid_12)->{'mempty'}));
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
}, "Apply0" => function($_dollar___unused_14) use ($applyRWST2_13_13) {
  $__num = \func_num_args();
  $__res = $applyRWST2_13_13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Semigroup0_14_17 = (($dictMonoid_12)->{'Semigroup0'})(null);
  $Semigroup0_15_18 = (($dictMonoid_12)->{'Semigroup0'})(null);
  $applyRWST2_15_18 = (object)["apply" => function($v_16) use ($Functor0_10_10, $Semigroup0_15_18, $__local_var_7_7) {
  $__num = \func_num_args();
  $__res = function($v1_17) use ($Functor0_10_10, $Semigroup0_15_18, $__local_var_7_7, $v_16) {
  $__num = \func_num_args();
  $__res = function($r_18) use ($Functor0_10_10, $Semigroup0_15_18, $__local_var_7_7, $v1_17, $v_16) {
  $__num = \func_num_args();
  $__res = function($s_19) use ($Functor0_10_10, $Semigroup0_15_18, $__local_var_7_7, $r_18, $v1_17, $v_16) {
  $__num = \func_num_args();
  $__res = ((($__local_var_7_7)->{'bind'})((($v_16)($r_18))($s_19)))(function($v2_20) use ($Functor0_10_10, $Semigroup0_15_18, $r_18, $v1_17) {
  $__num = \func_num_args();
  $__local_var_21_19 = ($v2_20)->{'value2'};
  $__res = ((($Functor0_10_10)->{'map'})(function($v3_22) use ($Semigroup0_15_18, $__local_var_21_19, $v2_20) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v3_22)->{'value0'}, (($v2_20)->{'value1'})(($v3_22)->{'value1'}), ((($Semigroup0_15_18)->{'append'})($__local_var_21_19))(($v3_22)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v1_17)($r_18))(($v2_20)->{'value0'}));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_16) use ($functorRWST1_11_11) {
  $__num = \func_num_args();
  $__res = $functorRWST1_11_11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $bindRWST2_14_17 = (object)["bind" => function($v_16) use ($Functor0_8_8, $Semigroup0_14_17, $__local_var_7_7) {
  $__num = \func_num_args();
  $__res = function($f_17) use ($Functor0_8_8, $Semigroup0_14_17, $__local_var_7_7, $v_16) {
  $__num = \func_num_args();
  $__res = function($r_18) use ($Functor0_8_8, $Semigroup0_14_17, $__local_var_7_7, $f_17, $v_16) {
  $__num = \func_num_args();
  $__res = function($s_19) use ($Functor0_8_8, $Semigroup0_14_17, $__local_var_7_7, $f_17, $r_18, $v_16) {
  $__num = \func_num_args();
  $__res = ((($__local_var_7_7)->{'bind'})((($v_16)($r_18))($s_19)))(function($v1_20) use ($Functor0_8_8, $Semigroup0_14_17, $f_17, $r_18) {
  $__num = \func_num_args();
  $__local_var_21_21 = ($v1_20)->{'value2'};
  $__res = ((($Functor0_8_8)->{'map'})(function($v3_22) use ($Semigroup0_14_17, $__local_var_21_21) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v3_22)->{'value0'}, ($v3_22)->{'value1'}, ((($Semigroup0_14_17)->{'append'})($__local_var_21_21))(($v3_22)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($f_17)(($v1_20)->{'value1'}))($r_18))(($v1_20)->{'value0'}));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_16) use ($applyRWST2_15_18) {
  $__num = \func_num_args();
  $__res = $applyRWST2_15_18;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["Applicative0" => function($_dollar___unused_15) use ($applicativeRWST2_13_13) {
  $__num = \func_num_args();
  $__res = $applicativeRWST2_13_13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_15) use ($bindRWST2_14_17) {
  $__num = \func_num_args();
  $__res = $bindRWST2_14_17;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $monadAskRWST1_3_2 = function($dictMonoid_8) use ($monadRWST1_7_7, $pure_1_0) {
  $__num = \func_num_args();
  $monadRWST2_9_24 = ($monadRWST1_7_7)($dictMonoid_8);
  $__res = (object)["ask" => function($r_10) use ($dictMonoid_8, $pure_1_0) {
  $__num = \func_num_args();
  $__res = function($s_11) use ($dictMonoid_8, $pure_1_0, $r_10) {
  $__num = \func_num_args();
  $__res = ($pure_1_0)(new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult($s_11, $r_10, ($dictMonoid_8)->{'mempty'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad0" => function($_dollar___unused_10) use ($monadRWST2_9_24) {
  $__num = \func_num_args();
  $__res = $monadRWST2_9_24;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = function($dictMonoid_4) use ($monadAskRWST1_3_2) {
  $__num = \func_num_args();
  $monadAskRWST2_5_26 = ($monadAskRWST1_3_2)($dictMonoid_4);
  $__res = (object)["local" => function($f_6) {
  $__num = \func_num_args();
  $__res = function($m_7) use ($f_6) {
  $__num = \func_num_args();
  $__res = function($r_8) use ($f_6, $m_7) {
  $__num = \func_num_args();
  $__res = function($s_9) use ($f_6, $m_7, $r_8) {
  $__num = \func_num_args();
  $__res = (($m_7)(($f_6)($r_8)))($s_9);
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "MonadAsk0" => function($_dollar___unused_6) use ($monadAskRWST2_5_26) {
  $__num = \func_num_args();
  $__res = $monadAskRWST2_5_26;
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
$GLOBALS['Control_Monad_RWS_Trans_monadReaderRWST'] = __NAMESPACE__ . '\\majControl_majMonad_majRmajWmajS_majTrans_monadmajReadermajRmajWmajSmajT';

// Control_Monad_RWS_Trans_monadEffectRWS
function majControl_majMonad_majRmajWmajS_majTrans_monadmajEffectmajRmajWmajS($dictMonoid_0, $dictMonadEffect_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majRmajWmajS_majTrans_monadmajEffectmajRmajWmajS';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $Monad0_2_0 = (($dictMonadEffect_1)->{'Monad0'})(null);
  $pure_3_1 = ((($Monad0_2_0)->{'Applicative0'})(null))->{'pure'};
  $__local_var_4_2 = (($Monad0_2_0)->{'Bind1'})(null);
  $Apply0_5_3 = (($__local_var_4_2)->{'Apply0'})(null);
  $Functor0_6_4 = (($Apply0_5_3)->{'Functor0'})(null);
  $__local_var_7_5 = (($Apply0_5_3)->{'Functor0'})(null);
  $functorRWST1_7_5 = (object)["map" => function($f_8) use ($__local_var_7_5) {
  $__num = \func_num_args();
  $__res = function($v_9) use ($__local_var_7_5, $f_8) {
  $__num = \func_num_args();
  $__res = function($r_10) use ($__local_var_7_5, $f_8, $v_9) {
  $__num = \func_num_args();
  $__res = function($s_11) use ($__local_var_7_5, $f_8, $r_10, $v_9) {
  $__num = \func_num_args();
  $__res = ((($__local_var_7_5)->{'map'})(function($v1_12) use ($f_8) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v1_12)->{'value0'}, ($f_8)(($v1_12)->{'value1'}), ($v1_12)->{'value2'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v_9)($r_10))($s_11));
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_8_7 = (($Monad0_2_0)->{'Bind1'})(null);
  $Functor0_9_8 = (((($__local_var_8_7)->{'Apply0'})(null))->{'Functor0'})(null);
  $Apply0_10_9 = (($__local_var_8_7)->{'Apply0'})(null);
  $Functor0_11_10 = (($Apply0_10_9)->{'Functor0'})(null);
  $__local_var_12_11 = (($Apply0_10_9)->{'Functor0'})(null);
  $functorRWST1_12_11 = (object)["map" => function($f_13) use ($__local_var_12_11) {
  $__num = \func_num_args();
  $__res = function($v_14) use ($__local_var_12_11, $f_13) {
  $__num = \func_num_args();
  $__res = function($r_15) use ($__local_var_12_11, $f_13, $v_14) {
  $__num = \func_num_args();
  $__res = function($s_16) use ($__local_var_12_11, $f_13, $r_15, $v_14) {
  $__num = \func_num_args();
  $__res = ((($__local_var_12_11)->{'map'})(function($v1_17) use ($f_13) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v1_17)->{'value0'}, ($f_13)(($v1_17)->{'value1'}), ($v1_17)->{'value2'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v_14)($r_15))($s_16));
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Semigroup0_13_13 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $applyRWST2_13_13 = (object)["apply" => function($v_14) use ($Functor0_6_4, $Semigroup0_13_13, $__local_var_4_2) {
  $__num = \func_num_args();
  $__res = function($v1_15) use ($Functor0_6_4, $Semigroup0_13_13, $__local_var_4_2, $v_14) {
  $__num = \func_num_args();
  $__res = function($r_16) use ($Functor0_6_4, $Semigroup0_13_13, $__local_var_4_2, $v1_15, $v_14) {
  $__num = \func_num_args();
  $__res = function($s_17) use ($Functor0_6_4, $Semigroup0_13_13, $__local_var_4_2, $r_16, $v1_15, $v_14) {
  $__num = \func_num_args();
  $__res = ((($__local_var_4_2)->{'bind'})((($v_14)($r_16))($s_17)))(function($v2_18) use ($Functor0_6_4, $Semigroup0_13_13, $r_16, $v1_15) {
  $__num = \func_num_args();
  $__local_var_19_14 = ($v2_18)->{'value2'};
  $__res = ((($Functor0_6_4)->{'map'})(function($v3_20) use ($Semigroup0_13_13, $__local_var_19_14, $v2_18) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v3_20)->{'value0'}, (($v2_18)->{'value1'})(($v3_20)->{'value1'}), ((($Semigroup0_13_13)->{'append'})($__local_var_19_14))(($v3_20)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v1_15)($r_16))(($v2_18)->{'value0'}));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_14) use ($functorRWST1_7_5) {
  $__num = \func_num_args();
  $__res = $functorRWST1_7_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeRWST2_13_13 = (object)["pure" => function($a_14) use ($dictMonoid_0, $pure_3_1) {
  $__num = \func_num_args();
  $__res = function($v_15) use ($a_14, $dictMonoid_0, $pure_3_1) {
  $__num = \func_num_args();
  $__res = function($s_16) use ($a_14, $dictMonoid_0, $pure_3_1) {
  $__num = \func_num_args();
  $__res = ($pure_3_1)(new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult($s_16, $a_14, ($dictMonoid_0)->{'mempty'}));
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
}, "Apply0" => function($_dollar___unused_14) use ($applyRWST2_13_13) {
  $__num = \func_num_args();
  $__res = $applyRWST2_13_13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Semigroup0_14_17 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $Semigroup0_15_18 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $applyRWST2_15_18 = (object)["apply" => function($v_16) use ($Functor0_11_10, $Semigroup0_15_18, $__local_var_8_7) {
  $__num = \func_num_args();
  $__res = function($v1_17) use ($Functor0_11_10, $Semigroup0_15_18, $__local_var_8_7, $v_16) {
  $__num = \func_num_args();
  $__res = function($r_18) use ($Functor0_11_10, $Semigroup0_15_18, $__local_var_8_7, $v1_17, $v_16) {
  $__num = \func_num_args();
  $__res = function($s_19) use ($Functor0_11_10, $Semigroup0_15_18, $__local_var_8_7, $r_18, $v1_17, $v_16) {
  $__num = \func_num_args();
  $__res = ((($__local_var_8_7)->{'bind'})((($v_16)($r_18))($s_19)))(function($v2_20) use ($Functor0_11_10, $Semigroup0_15_18, $r_18, $v1_17) {
  $__num = \func_num_args();
  $__local_var_21_19 = ($v2_20)->{'value2'};
  $__res = ((($Functor0_11_10)->{'map'})(function($v3_22) use ($Semigroup0_15_18, $__local_var_21_19, $v2_20) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v3_22)->{'value0'}, (($v2_20)->{'value1'})(($v3_22)->{'value1'}), ((($Semigroup0_15_18)->{'append'})($__local_var_21_19))(($v3_22)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v1_17)($r_18))(($v2_20)->{'value0'}));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_16) use ($functorRWST1_12_11) {
  $__num = \func_num_args();
  $__res = $functorRWST1_12_11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $bindRWST2_14_17 = (object)["bind" => function($v_16) use ($Functor0_9_8, $Semigroup0_14_17, $__local_var_8_7) {
  $__num = \func_num_args();
  $__res = function($f_17) use ($Functor0_9_8, $Semigroup0_14_17, $__local_var_8_7, $v_16) {
  $__num = \func_num_args();
  $__res = function($r_18) use ($Functor0_9_8, $Semigroup0_14_17, $__local_var_8_7, $f_17, $v_16) {
  $__num = \func_num_args();
  $__res = function($s_19) use ($Functor0_9_8, $Semigroup0_14_17, $__local_var_8_7, $f_17, $r_18, $v_16) {
  $__num = \func_num_args();
  $__res = ((($__local_var_8_7)->{'bind'})((($v_16)($r_18))($s_19)))(function($v1_20) use ($Functor0_9_8, $Semigroup0_14_17, $f_17, $r_18) {
  $__num = \func_num_args();
  $__local_var_21_21 = ($v1_20)->{'value2'};
  $__res = ((($Functor0_9_8)->{'map'})(function($v3_22) use ($Semigroup0_14_17, $__local_var_21_21) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v3_22)->{'value0'}, ($v3_22)->{'value1'}, ((($Semigroup0_14_17)->{'append'})($__local_var_21_21))(($v3_22)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($f_17)(($v1_20)->{'value1'}))($r_18))(($v1_20)->{'value0'}));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_16) use ($applyRWST2_15_18) {
  $__num = \func_num_args();
  $__res = $applyRWST2_15_18;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $monadRWST1_3_1 = (object)["Applicative0" => function($_dollar___unused_15) use ($applicativeRWST2_13_13) {
  $__num = \func_num_args();
  $__res = $applicativeRWST2_13_13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_15) use ($bindRWST2_14_17) {
  $__num = \func_num_args();
  $__res = $bindRWST2_14_17;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_4_24 = (($Monad0_2_0)->{'Bind1'})(null);
  $pure_5_25 = ((($Monad0_2_0)->{'Applicative0'})(null))->{'pure'};
  $__res = (object)["liftEffect" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($m_6) use ($Bind1_4_24, $dictMonoid_0, $pure_5_25) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($Bind1_4_24, $dictMonoid_0, $m_6, $pure_5_25) {
  $__num = \func_num_args();
  $__res = function($s_8) use ($Bind1_4_24, $dictMonoid_0, $m_6, $pure_5_25) {
  $__num = \func_num_args();
  $__res = ((($Bind1_4_24)->{'bind'})($m_6))(function($a_9) use ($dictMonoid_0, $pure_5_25, $s_8) {
  $__num = \func_num_args();
  $__res = ($pure_5_25)(new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult($s_8, $a_9, ($dictMonoid_0)->{'mempty'}));
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
}))(($dictMonadEffect_1)->{'liftEffect'}), "Monad0" => function($_dollar___unused_4) use ($monadRWST1_3_1) {
  $__num = \func_num_args();
  $__res = $monadRWST1_3_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Control_Monad_RWS_Trans_monadEffectRWS'] = __NAMESPACE__ . '\\majControl_majMonad_majRmajWmajS_majTrans_monadmajEffectmajRmajWmajS';

// Control_Monad_RWS_Trans_monadRecRWST
function majControl_majMonad_majRmajWmajS_majTrans_monadmajRecmajRmajWmajSmajT($dictMonadRec_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majRmajWmajS_majTrans_monadmajRecmajRmajWmajSmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Monad0_1_0 = (($dictMonadRec_0)->{'Monad0'})(null);
  $Bind1_2_1 = (($Monad0_1_0)->{'Bind1'})(null);
  $Applicative0_3_2 = (($Monad0_1_0)->{'Applicative0'})(null);
  $pure_4_3 = ((($Monad0_1_0)->{'Applicative0'})(null))->{'pure'};
  $__local_var_5_4 = (($Monad0_1_0)->{'Bind1'})(null);
  $Apply0_6_5 = (($__local_var_5_4)->{'Apply0'})(null);
  $Functor0_7_6 = (($Apply0_6_5)->{'Functor0'})(null);
  $__local_var_8_7 = (($Apply0_6_5)->{'Functor0'})(null);
  $functorRWST1_8_7 = (object)["map" => function($f_9) use ($__local_var_8_7) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($__local_var_8_7, $f_9) {
  $__num = \func_num_args();
  $__res = function($r_11) use ($__local_var_8_7, $f_9, $v_10) {
  $__num = \func_num_args();
  $__res = function($s_12) use ($__local_var_8_7, $f_9, $r_11, $v_10) {
  $__num = \func_num_args();
  $__res = ((($__local_var_8_7)->{'map'})(function($v1_13) use ($f_9) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v1_13)->{'value0'}, ($f_9)(($v1_13)->{'value1'}), ($v1_13)->{'value2'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v_10)($r_11))($s_12));
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_9_9 = (($Monad0_1_0)->{'Bind1'})(null);
  $Functor0_10_10 = (((($__local_var_9_9)->{'Apply0'})(null))->{'Functor0'})(null);
  $Apply0_11_11 = (($__local_var_9_9)->{'Apply0'})(null);
  $Functor0_12_12 = (($Apply0_11_11)->{'Functor0'})(null);
  $__local_var_13_13 = (($Apply0_11_11)->{'Functor0'})(null);
  $functorRWST1_13_13 = (object)["map" => function($f_14) use ($__local_var_13_13) {
  $__num = \func_num_args();
  $__res = function($v_15) use ($__local_var_13_13, $f_14) {
  $__num = \func_num_args();
  $__res = function($r_16) use ($__local_var_13_13, $f_14, $v_15) {
  $__num = \func_num_args();
  $__res = function($s_17) use ($__local_var_13_13, $f_14, $r_16, $v_15) {
  $__num = \func_num_args();
  $__res = ((($__local_var_13_13)->{'map'})(function($v1_18) use ($f_14) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v1_18)->{'value0'}, ($f_14)(($v1_18)->{'value1'}), ($v1_18)->{'value2'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v_15)($r_16))($s_17));
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $monadRWST1_7_6 = function($dictMonoid_14) use ($Functor0_10_10, $Functor0_12_12, $Functor0_7_6, $__local_var_5_4, $__local_var_9_9, $functorRWST1_13_13, $functorRWST1_8_7, $pure_4_3) {
  $__num = \func_num_args();
  $Semigroup0_15_15 = (($dictMonoid_14)->{'Semigroup0'})(null);
  $applyRWST2_15_15 = (object)["apply" => function($v_16) use ($Functor0_7_6, $Semigroup0_15_15, $__local_var_5_4) {
  $__num = \func_num_args();
  $__res = function($v1_17) use ($Functor0_7_6, $Semigroup0_15_15, $__local_var_5_4, $v_16) {
  $__num = \func_num_args();
  $__res = function($r_18) use ($Functor0_7_6, $Semigroup0_15_15, $__local_var_5_4, $v1_17, $v_16) {
  $__num = \func_num_args();
  $__res = function($s_19) use ($Functor0_7_6, $Semigroup0_15_15, $__local_var_5_4, $r_18, $v1_17, $v_16) {
  $__num = \func_num_args();
  $__res = ((($__local_var_5_4)->{'bind'})((($v_16)($r_18))($s_19)))(function($v2_20) use ($Functor0_7_6, $Semigroup0_15_15, $r_18, $v1_17) {
  $__num = \func_num_args();
  $__local_var_21_16 = ($v2_20)->{'value2'};
  $__res = ((($Functor0_7_6)->{'map'})(function($v3_22) use ($Semigroup0_15_15, $__local_var_21_16, $v2_20) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v3_22)->{'value0'}, (($v2_20)->{'value1'})(($v3_22)->{'value1'}), ((($Semigroup0_15_15)->{'append'})($__local_var_21_16))(($v3_22)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v1_17)($r_18))(($v2_20)->{'value0'}));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_16) use ($functorRWST1_8_7) {
  $__num = \func_num_args();
  $__res = $functorRWST1_8_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeRWST2_15_15 = (object)["pure" => function($a_16) use ($dictMonoid_14, $pure_4_3) {
  $__num = \func_num_args();
  $__res = function($v_17) use ($a_16, $dictMonoid_14, $pure_4_3) {
  $__num = \func_num_args();
  $__res = function($s_18) use ($a_16, $dictMonoid_14, $pure_4_3) {
  $__num = \func_num_args();
  $__res = ($pure_4_3)(new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult($s_18, $a_16, ($dictMonoid_14)->{'mempty'}));
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
}, "Apply0" => function($_dollar___unused_16) use ($applyRWST2_15_15) {
  $__num = \func_num_args();
  $__res = $applyRWST2_15_15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Semigroup0_16_19 = (($dictMonoid_14)->{'Semigroup0'})(null);
  $Semigroup0_17_20 = (($dictMonoid_14)->{'Semigroup0'})(null);
  $applyRWST2_17_20 = (object)["apply" => function($v_18) use ($Functor0_12_12, $Semigroup0_17_20, $__local_var_9_9) {
  $__num = \func_num_args();
  $__res = function($v1_19) use ($Functor0_12_12, $Semigroup0_17_20, $__local_var_9_9, $v_18) {
  $__num = \func_num_args();
  $__res = function($r_20) use ($Functor0_12_12, $Semigroup0_17_20, $__local_var_9_9, $v1_19, $v_18) {
  $__num = \func_num_args();
  $__res = function($s_21) use ($Functor0_12_12, $Semigroup0_17_20, $__local_var_9_9, $r_20, $v1_19, $v_18) {
  $__num = \func_num_args();
  $__res = ((($__local_var_9_9)->{'bind'})((($v_18)($r_20))($s_21)))(function($v2_22) use ($Functor0_12_12, $Semigroup0_17_20, $r_20, $v1_19) {
  $__num = \func_num_args();
  $__local_var_23_21 = ($v2_22)->{'value2'};
  $__res = ((($Functor0_12_12)->{'map'})(function($v3_24) use ($Semigroup0_17_20, $__local_var_23_21, $v2_22) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v3_24)->{'value0'}, (($v2_22)->{'value1'})(($v3_24)->{'value1'}), ((($Semigroup0_17_20)->{'append'})($__local_var_23_21))(($v3_24)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v1_19)($r_20))(($v2_22)->{'value0'}));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_18) use ($functorRWST1_13_13) {
  $__num = \func_num_args();
  $__res = $functorRWST1_13_13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $bindRWST2_16_19 = (object)["bind" => function($v_18) use ($Functor0_10_10, $Semigroup0_16_19, $__local_var_9_9) {
  $__num = \func_num_args();
  $__res = function($f_19) use ($Functor0_10_10, $Semigroup0_16_19, $__local_var_9_9, $v_18) {
  $__num = \func_num_args();
  $__res = function($r_20) use ($Functor0_10_10, $Semigroup0_16_19, $__local_var_9_9, $f_19, $v_18) {
  $__num = \func_num_args();
  $__res = function($s_21) use ($Functor0_10_10, $Semigroup0_16_19, $__local_var_9_9, $f_19, $r_20, $v_18) {
  $__num = \func_num_args();
  $__res = ((($__local_var_9_9)->{'bind'})((($v_18)($r_20))($s_21)))(function($v1_22) use ($Functor0_10_10, $Semigroup0_16_19, $f_19, $r_20) {
  $__num = \func_num_args();
  $__local_var_23_23 = ($v1_22)->{'value2'};
  $__res = ((($Functor0_10_10)->{'map'})(function($v3_24) use ($Semigroup0_16_19, $__local_var_23_23) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v3_24)->{'value0'}, ($v3_24)->{'value1'}, ((($Semigroup0_16_19)->{'append'})($__local_var_23_23))(($v3_24)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($f_19)(($v1_22)->{'value1'}))($r_20))(($v1_22)->{'value0'}));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_18) use ($applyRWST2_17_20) {
  $__num = \func_num_args();
  $__res = $applyRWST2_17_20;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["Applicative0" => function($_dollar___unused_17) use ($applicativeRWST2_15_15) {
  $__num = \func_num_args();
  $__res = $applicativeRWST2_15_15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_17) use ($bindRWST2_16_19) {
  $__num = \func_num_args();
  $__res = $bindRWST2_16_19;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = function($dictMonoid_8) use ($Applicative0_3_2, $Bind1_2_1, $dictMonadRec_0, $monadRWST1_7_6) {
  $__num = \func_num_args();
  $Semigroup0_9_26 = (($dictMonoid_8)->{'Semigroup0'})(null);
  $monadRWST2_10_27 = ($monadRWST1_7_6)($dictMonoid_8);
  $__res = (object)["tailRecM" => function($k_11) use ($Applicative0_3_2, $Bind1_2_1, $Semigroup0_9_26, $dictMonadRec_0, $dictMonoid_8) {
  $__num = \func_num_args();
  $__res = function($a_12) use ($Applicative0_3_2, $Bind1_2_1, $Semigroup0_9_26, $dictMonadRec_0, $dictMonoid_8, $k_11) {
  $__num = \func_num_args();
  $__res = function($r_13) use ($Applicative0_3_2, $Bind1_2_1, $Semigroup0_9_26, $a_12, $dictMonadRec_0, $dictMonoid_8, $k_11) {
  $__num = \func_num_args();
  $__res = function($s_14) use ($Applicative0_3_2, $Bind1_2_1, $Semigroup0_9_26, $a_12, $dictMonadRec_0, $dictMonoid_8, $k_11, $r_13) {
  $__num = \func_num_args();
  $__res = ((($dictMonadRec_0)->{'tailRecM'})(function($v_15) use ($Applicative0_3_2, $Bind1_2_1, $Semigroup0_9_26, $k_11, $r_13) {
  $__num = \func_num_args();
  $__local_var_16_28 = ($v_15)->{'value2'};
  $__res = ((($Bind1_2_1)->{'bind'})(((($k_11)(($v_15)->{'value1'}))($r_13))(($v_15)->{'value0'})))(function($v2_17) use ($Applicative0_3_2, $Semigroup0_9_26, $__local_var_16_28) {
  $__num = \func_num_args();
  $__t29 = null;;
  if (($v2_17)->{'value1'} instanceof \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Loop) {
$__t29 = new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Loop(new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v2_17)->{'value0'}, (($v2_17)->{'value1'})->{'value0'}, ((($Semigroup0_9_26)->{'append'})($__local_var_16_28))(($v2_17)->{'value2'})));
goto end_branch_29;;
};
  if (($v2_17)->{'value1'} instanceof \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Done) {
$__t29 = new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Done(new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v2_17)->{'value0'}, (($v2_17)->{'value1'})->{'value0'}, ((($Semigroup0_9_26)->{'append'})($__local_var_16_28))(($v2_17)->{'value2'})));
goto end_branch_29;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t29 = null;
  end_branch_29:;
  $__res = (($Applicative0_3_2)->{'pure'})($__t29);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult($s_14, $a_12, ($dictMonoid_8)->{'mempty'}));
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad0" => function($_dollar___unused_11) use ($monadRWST2_10_27) {
  $__num = \func_num_args();
  $__res = $monadRWST2_10_27;
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
$GLOBALS['Control_Monad_RWS_Trans_monadRecRWST'] = __NAMESPACE__ . '\\majControl_majMonad_majRmajWmajS_majTrans_monadmajRecmajRmajWmajSmajT';

// Control_Monad_RWS_Trans_monadStateRWST
function majControl_majMonad_majRmajWmajS_majTrans_monadmajStatemajRmajWmajSmajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majRmajWmajS_majTrans_monadmajStatemajRmajWmajSmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $pure_1_0 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $pure_2_1 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $__local_var_3_2 = (($dictMonad_0)->{'Bind1'})(null);
  $Apply0_4_3 = (($__local_var_3_2)->{'Apply0'})(null);
  $Functor0_5_4 = (($Apply0_4_3)->{'Functor0'})(null);
  $__local_var_6_5 = (($Apply0_4_3)->{'Functor0'})(null);
  $functorRWST1_6_5 = (object)["map" => function($f_7) use ($__local_var_6_5) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($__local_var_6_5, $f_7) {
  $__num = \func_num_args();
  $__res = function($r_9) use ($__local_var_6_5, $f_7, $v_8) {
  $__num = \func_num_args();
  $__res = function($s_10) use ($__local_var_6_5, $f_7, $r_9, $v_8) {
  $__num = \func_num_args();
  $__res = ((($__local_var_6_5)->{'map'})(function($v1_11) use ($f_7) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v1_11)->{'value0'}, ($f_7)(($v1_11)->{'value1'}), ($v1_11)->{'value2'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v_8)($r_9))($s_10));
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_7_7 = (($dictMonad_0)->{'Bind1'})(null);
  $Functor0_8_8 = (((($__local_var_7_7)->{'Apply0'})(null))->{'Functor0'})(null);
  $Apply0_9_9 = (($__local_var_7_7)->{'Apply0'})(null);
  $Functor0_10_10 = (($Apply0_9_9)->{'Functor0'})(null);
  $__local_var_11_11 = (($Apply0_9_9)->{'Functor0'})(null);
  $functorRWST1_11_11 = (object)["map" => function($f_12) use ($__local_var_11_11) {
  $__num = \func_num_args();
  $__res = function($v_13) use ($__local_var_11_11, $f_12) {
  $__num = \func_num_args();
  $__res = function($r_14) use ($__local_var_11_11, $f_12, $v_13) {
  $__num = \func_num_args();
  $__res = function($s_15) use ($__local_var_11_11, $f_12, $r_14, $v_13) {
  $__num = \func_num_args();
  $__res = ((($__local_var_11_11)->{'map'})(function($v1_16) use ($f_12) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v1_16)->{'value0'}, ($f_12)(($v1_16)->{'value1'}), ($v1_16)->{'value2'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v_13)($r_14))($s_15));
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $monadRWST1_5_4 = function($dictMonoid_12) use ($Functor0_10_10, $Functor0_5_4, $Functor0_8_8, $__local_var_3_2, $__local_var_7_7, $functorRWST1_11_11, $functorRWST1_6_5, $pure_2_1) {
  $__num = \func_num_args();
  $Semigroup0_13_13 = (($dictMonoid_12)->{'Semigroup0'})(null);
  $applyRWST2_13_13 = (object)["apply" => function($v_14) use ($Functor0_5_4, $Semigroup0_13_13, $__local_var_3_2) {
  $__num = \func_num_args();
  $__res = function($v1_15) use ($Functor0_5_4, $Semigroup0_13_13, $__local_var_3_2, $v_14) {
  $__num = \func_num_args();
  $__res = function($r_16) use ($Functor0_5_4, $Semigroup0_13_13, $__local_var_3_2, $v1_15, $v_14) {
  $__num = \func_num_args();
  $__res = function($s_17) use ($Functor0_5_4, $Semigroup0_13_13, $__local_var_3_2, $r_16, $v1_15, $v_14) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_2)->{'bind'})((($v_14)($r_16))($s_17)))(function($v2_18) use ($Functor0_5_4, $Semigroup0_13_13, $r_16, $v1_15) {
  $__num = \func_num_args();
  $__local_var_19_14 = ($v2_18)->{'value2'};
  $__res = ((($Functor0_5_4)->{'map'})(function($v3_20) use ($Semigroup0_13_13, $__local_var_19_14, $v2_18) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v3_20)->{'value0'}, (($v2_18)->{'value1'})(($v3_20)->{'value1'}), ((($Semigroup0_13_13)->{'append'})($__local_var_19_14))(($v3_20)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v1_15)($r_16))(($v2_18)->{'value0'}));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_14) use ($functorRWST1_6_5) {
  $__num = \func_num_args();
  $__res = $functorRWST1_6_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeRWST2_13_13 = (object)["pure" => function($a_14) use ($dictMonoid_12, $pure_2_1) {
  $__num = \func_num_args();
  $__res = function($v_15) use ($a_14, $dictMonoid_12, $pure_2_1) {
  $__num = \func_num_args();
  $__res = function($s_16) use ($a_14, $dictMonoid_12, $pure_2_1) {
  $__num = \func_num_args();
  $__res = ($pure_2_1)(new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult($s_16, $a_14, ($dictMonoid_12)->{'mempty'}));
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
}, "Apply0" => function($_dollar___unused_14) use ($applyRWST2_13_13) {
  $__num = \func_num_args();
  $__res = $applyRWST2_13_13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Semigroup0_14_17 = (($dictMonoid_12)->{'Semigroup0'})(null);
  $Semigroup0_15_18 = (($dictMonoid_12)->{'Semigroup0'})(null);
  $applyRWST2_15_18 = (object)["apply" => function($v_16) use ($Functor0_10_10, $Semigroup0_15_18, $__local_var_7_7) {
  $__num = \func_num_args();
  $__res = function($v1_17) use ($Functor0_10_10, $Semigroup0_15_18, $__local_var_7_7, $v_16) {
  $__num = \func_num_args();
  $__res = function($r_18) use ($Functor0_10_10, $Semigroup0_15_18, $__local_var_7_7, $v1_17, $v_16) {
  $__num = \func_num_args();
  $__res = function($s_19) use ($Functor0_10_10, $Semigroup0_15_18, $__local_var_7_7, $r_18, $v1_17, $v_16) {
  $__num = \func_num_args();
  $__res = ((($__local_var_7_7)->{'bind'})((($v_16)($r_18))($s_19)))(function($v2_20) use ($Functor0_10_10, $Semigroup0_15_18, $r_18, $v1_17) {
  $__num = \func_num_args();
  $__local_var_21_19 = ($v2_20)->{'value2'};
  $__res = ((($Functor0_10_10)->{'map'})(function($v3_22) use ($Semigroup0_15_18, $__local_var_21_19, $v2_20) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v3_22)->{'value0'}, (($v2_20)->{'value1'})(($v3_22)->{'value1'}), ((($Semigroup0_15_18)->{'append'})($__local_var_21_19))(($v3_22)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v1_17)($r_18))(($v2_20)->{'value0'}));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_16) use ($functorRWST1_11_11) {
  $__num = \func_num_args();
  $__res = $functorRWST1_11_11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $bindRWST2_14_17 = (object)["bind" => function($v_16) use ($Functor0_8_8, $Semigroup0_14_17, $__local_var_7_7) {
  $__num = \func_num_args();
  $__res = function($f_17) use ($Functor0_8_8, $Semigroup0_14_17, $__local_var_7_7, $v_16) {
  $__num = \func_num_args();
  $__res = function($r_18) use ($Functor0_8_8, $Semigroup0_14_17, $__local_var_7_7, $f_17, $v_16) {
  $__num = \func_num_args();
  $__res = function($s_19) use ($Functor0_8_8, $Semigroup0_14_17, $__local_var_7_7, $f_17, $r_18, $v_16) {
  $__num = \func_num_args();
  $__res = ((($__local_var_7_7)->{'bind'})((($v_16)($r_18))($s_19)))(function($v1_20) use ($Functor0_8_8, $Semigroup0_14_17, $f_17, $r_18) {
  $__num = \func_num_args();
  $__local_var_21_21 = ($v1_20)->{'value2'};
  $__res = ((($Functor0_8_8)->{'map'})(function($v3_22) use ($Semigroup0_14_17, $__local_var_21_21) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v3_22)->{'value0'}, ($v3_22)->{'value1'}, ((($Semigroup0_14_17)->{'append'})($__local_var_21_21))(($v3_22)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($f_17)(($v1_20)->{'value1'}))($r_18))(($v1_20)->{'value0'}));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_16) use ($applyRWST2_15_18) {
  $__num = \func_num_args();
  $__res = $applyRWST2_15_18;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["Applicative0" => function($_dollar___unused_15) use ($applicativeRWST2_13_13) {
  $__num = \func_num_args();
  $__res = $applicativeRWST2_13_13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_15) use ($bindRWST2_14_17) {
  $__num = \func_num_args();
  $__res = $bindRWST2_14_17;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = function($dictMonoid_6) use ($monadRWST1_5_4, $pure_1_0) {
  $__num = \func_num_args();
  $monadRWST2_7_24 = ($monadRWST1_5_4)($dictMonoid_6);
  $__res = (object)["state" => function($f_8) use ($dictMonoid_6, $pure_1_0) {
  $__num = \func_num_args();
  $__res = function($v_9) use ($dictMonoid_6, $f_8, $pure_1_0) {
  $__num = \func_num_args();
  $__res = function($s_10) use ($dictMonoid_6, $f_8, $pure_1_0) {
  $__num = \func_num_args();
  $v1_11_25 = ($f_8)($s_10);
  $__res = ($pure_1_0)(new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v1_11_25)->{'value1'}, ($v1_11_25)->{'value0'}, ($dictMonoid_6)->{'mempty'}));
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
}, "Monad0" => function($_dollar___unused_8) use ($monadRWST2_7_24) {
  $__num = \func_num_args();
  $__res = $monadRWST2_7_24;
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
$GLOBALS['Control_Monad_RWS_Trans_monadStateRWST'] = __NAMESPACE__ . '\\majControl_majMonad_majRmajWmajS_majTrans_monadmajStatemajRmajWmajSmajT';

// Control_Monad_RWS_Trans_monadTellRWST
function majControl_majMonad_majRmajWmajS_majTrans_monadmajTellmajRmajWmajSmajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majRmajWmajS_majTrans_monadmajTellmajRmajWmajSmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $pure_1_0 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $pure_2_1 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $__local_var_3_2 = (($dictMonad_0)->{'Bind1'})(null);
  $Apply0_4_3 = (($__local_var_3_2)->{'Apply0'})(null);
  $Functor0_5_4 = (($Apply0_4_3)->{'Functor0'})(null);
  $__local_var_6_5 = (($Apply0_4_3)->{'Functor0'})(null);
  $functorRWST1_6_5 = (object)["map" => function($f_7) use ($__local_var_6_5) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($__local_var_6_5, $f_7) {
  $__num = \func_num_args();
  $__res = function($r_9) use ($__local_var_6_5, $f_7, $v_8) {
  $__num = \func_num_args();
  $__res = function($s_10) use ($__local_var_6_5, $f_7, $r_9, $v_8) {
  $__num = \func_num_args();
  $__res = ((($__local_var_6_5)->{'map'})(function($v1_11) use ($f_7) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v1_11)->{'value0'}, ($f_7)(($v1_11)->{'value1'}), ($v1_11)->{'value2'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v_8)($r_9))($s_10));
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_7_7 = (($dictMonad_0)->{'Bind1'})(null);
  $Functor0_8_8 = (((($__local_var_7_7)->{'Apply0'})(null))->{'Functor0'})(null);
  $Apply0_9_9 = (($__local_var_7_7)->{'Apply0'})(null);
  $Functor0_10_10 = (($Apply0_9_9)->{'Functor0'})(null);
  $__local_var_11_11 = (($Apply0_9_9)->{'Functor0'})(null);
  $functorRWST1_11_11 = (object)["map" => function($f_12) use ($__local_var_11_11) {
  $__num = \func_num_args();
  $__res = function($v_13) use ($__local_var_11_11, $f_12) {
  $__num = \func_num_args();
  $__res = function($r_14) use ($__local_var_11_11, $f_12, $v_13) {
  $__num = \func_num_args();
  $__res = function($s_15) use ($__local_var_11_11, $f_12, $r_14, $v_13) {
  $__num = \func_num_args();
  $__res = ((($__local_var_11_11)->{'map'})(function($v1_16) use ($f_12) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v1_16)->{'value0'}, ($f_12)(($v1_16)->{'value1'}), ($v1_16)->{'value2'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v_13)($r_14))($s_15));
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $monadRWST1_5_4 = function($dictMonoid_12) use ($Functor0_10_10, $Functor0_5_4, $Functor0_8_8, $__local_var_3_2, $__local_var_7_7, $functorRWST1_11_11, $functorRWST1_6_5, $pure_2_1) {
  $__num = \func_num_args();
  $Semigroup0_13_13 = (($dictMonoid_12)->{'Semigroup0'})(null);
  $applyRWST2_13_13 = (object)["apply" => function($v_14) use ($Functor0_5_4, $Semigroup0_13_13, $__local_var_3_2) {
  $__num = \func_num_args();
  $__res = function($v1_15) use ($Functor0_5_4, $Semigroup0_13_13, $__local_var_3_2, $v_14) {
  $__num = \func_num_args();
  $__res = function($r_16) use ($Functor0_5_4, $Semigroup0_13_13, $__local_var_3_2, $v1_15, $v_14) {
  $__num = \func_num_args();
  $__res = function($s_17) use ($Functor0_5_4, $Semigroup0_13_13, $__local_var_3_2, $r_16, $v1_15, $v_14) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_2)->{'bind'})((($v_14)($r_16))($s_17)))(function($v2_18) use ($Functor0_5_4, $Semigroup0_13_13, $r_16, $v1_15) {
  $__num = \func_num_args();
  $__local_var_19_14 = ($v2_18)->{'value2'};
  $__res = ((($Functor0_5_4)->{'map'})(function($v3_20) use ($Semigroup0_13_13, $__local_var_19_14, $v2_18) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v3_20)->{'value0'}, (($v2_18)->{'value1'})(($v3_20)->{'value1'}), ((($Semigroup0_13_13)->{'append'})($__local_var_19_14))(($v3_20)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v1_15)($r_16))(($v2_18)->{'value0'}));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_14) use ($functorRWST1_6_5) {
  $__num = \func_num_args();
  $__res = $functorRWST1_6_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeRWST2_13_13 = (object)["pure" => function($a_14) use ($dictMonoid_12, $pure_2_1) {
  $__num = \func_num_args();
  $__res = function($v_15) use ($a_14, $dictMonoid_12, $pure_2_1) {
  $__num = \func_num_args();
  $__res = function($s_16) use ($a_14, $dictMonoid_12, $pure_2_1) {
  $__num = \func_num_args();
  $__res = ($pure_2_1)(new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult($s_16, $a_14, ($dictMonoid_12)->{'mempty'}));
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
}, "Apply0" => function($_dollar___unused_14) use ($applyRWST2_13_13) {
  $__num = \func_num_args();
  $__res = $applyRWST2_13_13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Semigroup0_14_17 = (($dictMonoid_12)->{'Semigroup0'})(null);
  $Semigroup0_15_18 = (($dictMonoid_12)->{'Semigroup0'})(null);
  $applyRWST2_15_18 = (object)["apply" => function($v_16) use ($Functor0_10_10, $Semigroup0_15_18, $__local_var_7_7) {
  $__num = \func_num_args();
  $__res = function($v1_17) use ($Functor0_10_10, $Semigroup0_15_18, $__local_var_7_7, $v_16) {
  $__num = \func_num_args();
  $__res = function($r_18) use ($Functor0_10_10, $Semigroup0_15_18, $__local_var_7_7, $v1_17, $v_16) {
  $__num = \func_num_args();
  $__res = function($s_19) use ($Functor0_10_10, $Semigroup0_15_18, $__local_var_7_7, $r_18, $v1_17, $v_16) {
  $__num = \func_num_args();
  $__res = ((($__local_var_7_7)->{'bind'})((($v_16)($r_18))($s_19)))(function($v2_20) use ($Functor0_10_10, $Semigroup0_15_18, $r_18, $v1_17) {
  $__num = \func_num_args();
  $__local_var_21_19 = ($v2_20)->{'value2'};
  $__res = ((($Functor0_10_10)->{'map'})(function($v3_22) use ($Semigroup0_15_18, $__local_var_21_19, $v2_20) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v3_22)->{'value0'}, (($v2_20)->{'value1'})(($v3_22)->{'value1'}), ((($Semigroup0_15_18)->{'append'})($__local_var_21_19))(($v3_22)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v1_17)($r_18))(($v2_20)->{'value0'}));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_16) use ($functorRWST1_11_11) {
  $__num = \func_num_args();
  $__res = $functorRWST1_11_11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $bindRWST2_14_17 = (object)["bind" => function($v_16) use ($Functor0_8_8, $Semigroup0_14_17, $__local_var_7_7) {
  $__num = \func_num_args();
  $__res = function($f_17) use ($Functor0_8_8, $Semigroup0_14_17, $__local_var_7_7, $v_16) {
  $__num = \func_num_args();
  $__res = function($r_18) use ($Functor0_8_8, $Semigroup0_14_17, $__local_var_7_7, $f_17, $v_16) {
  $__num = \func_num_args();
  $__res = function($s_19) use ($Functor0_8_8, $Semigroup0_14_17, $__local_var_7_7, $f_17, $r_18, $v_16) {
  $__num = \func_num_args();
  $__res = ((($__local_var_7_7)->{'bind'})((($v_16)($r_18))($s_19)))(function($v1_20) use ($Functor0_8_8, $Semigroup0_14_17, $f_17, $r_18) {
  $__num = \func_num_args();
  $__local_var_21_21 = ($v1_20)->{'value2'};
  $__res = ((($Functor0_8_8)->{'map'})(function($v3_22) use ($Semigroup0_14_17, $__local_var_21_21) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v3_22)->{'value0'}, ($v3_22)->{'value1'}, ((($Semigroup0_14_17)->{'append'})($__local_var_21_21))(($v3_22)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($f_17)(($v1_20)->{'value1'}))($r_18))(($v1_20)->{'value0'}));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_16) use ($applyRWST2_15_18) {
  $__num = \func_num_args();
  $__res = $applyRWST2_15_18;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["Applicative0" => function($_dollar___unused_15) use ($applicativeRWST2_13_13) {
  $__num = \func_num_args();
  $__res = $applicativeRWST2_13_13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_15) use ($bindRWST2_14_17) {
  $__num = \func_num_args();
  $__res = $bindRWST2_14_17;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = function($dictMonoid_6) use ($monadRWST1_5_4, $pure_1_0) {
  $__num = \func_num_args();
  $Semigroup0_7_24 = (($dictMonoid_6)->{'Semigroup0'})(null);
  $monadRWST2_8_25 = ($monadRWST1_5_4)($dictMonoid_6);
  $__res = (object)["tell" => function($w_9) use ($pure_1_0) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($pure_1_0, $w_9) {
  $__num = \func_num_args();
  $__res = function($s_11) use ($pure_1_0, $w_9) {
  $__num = \func_num_args();
  $__res = ($pure_1_0)(new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult($s_11, $GLOBALS['Data_Unit_unit'], $w_9));
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
}, "Semigroup0" => function($_dollar___unused_9) use ($Semigroup0_7_24) {
  $__num = \func_num_args();
  $__res = $Semigroup0_7_24;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad1" => function($_dollar___unused_9) use ($monadRWST2_8_25) {
  $__num = \func_num_args();
  $__res = $monadRWST2_8_25;
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
$GLOBALS['Control_Monad_RWS_Trans_monadTellRWST'] = __NAMESPACE__ . '\\majControl_majMonad_majRmajWmajS_majTrans_monadmajTellmajRmajWmajSmajT';

// Control_Monad_RWS_Trans_monadWriterRWST
function majControl_majMonad_majRmajWmajS_majTrans_monadmajWritermajRmajWmajSmajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majRmajWmajS_majTrans_monadmajWritermajRmajWmajSmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Bind1_1_0 = (($dictMonad_0)->{'Bind1'})(null);
  $Applicative0_2_1 = (($dictMonad_0)->{'Applicative0'})(null);
  $pure_3_2 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $pure_4_3 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $__local_var_5_4 = (($dictMonad_0)->{'Bind1'})(null);
  $Apply0_6_5 = (($__local_var_5_4)->{'Apply0'})(null);
  $Functor0_7_6 = (($Apply0_6_5)->{'Functor0'})(null);
  $__local_var_8_7 = (($Apply0_6_5)->{'Functor0'})(null);
  $functorRWST1_8_7 = (object)["map" => function($f_9) use ($__local_var_8_7) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($__local_var_8_7, $f_9) {
  $__num = \func_num_args();
  $__res = function($r_11) use ($__local_var_8_7, $f_9, $v_10) {
  $__num = \func_num_args();
  $__res = function($s_12) use ($__local_var_8_7, $f_9, $r_11, $v_10) {
  $__num = \func_num_args();
  $__res = ((($__local_var_8_7)->{'map'})(function($v1_13) use ($f_9) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v1_13)->{'value0'}, ($f_9)(($v1_13)->{'value1'}), ($v1_13)->{'value2'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v_10)($r_11))($s_12));
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_9_9 = (($dictMonad_0)->{'Bind1'})(null);
  $Functor0_10_10 = (((($__local_var_9_9)->{'Apply0'})(null))->{'Functor0'})(null);
  $Apply0_11_11 = (($__local_var_9_9)->{'Apply0'})(null);
  $Functor0_12_12 = (($Apply0_11_11)->{'Functor0'})(null);
  $__local_var_13_13 = (($Apply0_11_11)->{'Functor0'})(null);
  $functorRWST1_13_13 = (object)["map" => function($f_14) use ($__local_var_13_13) {
  $__num = \func_num_args();
  $__res = function($v_15) use ($__local_var_13_13, $f_14) {
  $__num = \func_num_args();
  $__res = function($r_16) use ($__local_var_13_13, $f_14, $v_15) {
  $__num = \func_num_args();
  $__res = function($s_17) use ($__local_var_13_13, $f_14, $r_16, $v_15) {
  $__num = \func_num_args();
  $__res = ((($__local_var_13_13)->{'map'})(function($v1_18) use ($f_14) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v1_18)->{'value0'}, ($f_14)(($v1_18)->{'value1'}), ($v1_18)->{'value2'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v_15)($r_16))($s_17));
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $monadRWST1_9_9 = function($dictMonoid_14) use ($Functor0_10_10, $Functor0_12_12, $Functor0_7_6, $__local_var_5_4, $__local_var_9_9, $functorRWST1_13_13, $functorRWST1_8_7, $pure_4_3) {
  $__num = \func_num_args();
  $Semigroup0_15_15 = (($dictMonoid_14)->{'Semigroup0'})(null);
  $applyRWST2_15_15 = (object)["apply" => function($v_16) use ($Functor0_7_6, $Semigroup0_15_15, $__local_var_5_4) {
  $__num = \func_num_args();
  $__res = function($v1_17) use ($Functor0_7_6, $Semigroup0_15_15, $__local_var_5_4, $v_16) {
  $__num = \func_num_args();
  $__res = function($r_18) use ($Functor0_7_6, $Semigroup0_15_15, $__local_var_5_4, $v1_17, $v_16) {
  $__num = \func_num_args();
  $__res = function($s_19) use ($Functor0_7_6, $Semigroup0_15_15, $__local_var_5_4, $r_18, $v1_17, $v_16) {
  $__num = \func_num_args();
  $__res = ((($__local_var_5_4)->{'bind'})((($v_16)($r_18))($s_19)))(function($v2_20) use ($Functor0_7_6, $Semigroup0_15_15, $r_18, $v1_17) {
  $__num = \func_num_args();
  $__local_var_21_16 = ($v2_20)->{'value2'};
  $__res = ((($Functor0_7_6)->{'map'})(function($v3_22) use ($Semigroup0_15_15, $__local_var_21_16, $v2_20) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v3_22)->{'value0'}, (($v2_20)->{'value1'})(($v3_22)->{'value1'}), ((($Semigroup0_15_15)->{'append'})($__local_var_21_16))(($v3_22)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v1_17)($r_18))(($v2_20)->{'value0'}));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_16) use ($functorRWST1_8_7) {
  $__num = \func_num_args();
  $__res = $functorRWST1_8_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeRWST2_15_15 = (object)["pure" => function($a_16) use ($dictMonoid_14, $pure_4_3) {
  $__num = \func_num_args();
  $__res = function($v_17) use ($a_16, $dictMonoid_14, $pure_4_3) {
  $__num = \func_num_args();
  $__res = function($s_18) use ($a_16, $dictMonoid_14, $pure_4_3) {
  $__num = \func_num_args();
  $__res = ($pure_4_3)(new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult($s_18, $a_16, ($dictMonoid_14)->{'mempty'}));
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
}, "Apply0" => function($_dollar___unused_16) use ($applyRWST2_15_15) {
  $__num = \func_num_args();
  $__res = $applyRWST2_15_15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Semigroup0_16_19 = (($dictMonoid_14)->{'Semigroup0'})(null);
  $Semigroup0_17_20 = (($dictMonoid_14)->{'Semigroup0'})(null);
  $applyRWST2_17_20 = (object)["apply" => function($v_18) use ($Functor0_12_12, $Semigroup0_17_20, $__local_var_9_9) {
  $__num = \func_num_args();
  $__res = function($v1_19) use ($Functor0_12_12, $Semigroup0_17_20, $__local_var_9_9, $v_18) {
  $__num = \func_num_args();
  $__res = function($r_20) use ($Functor0_12_12, $Semigroup0_17_20, $__local_var_9_9, $v1_19, $v_18) {
  $__num = \func_num_args();
  $__res = function($s_21) use ($Functor0_12_12, $Semigroup0_17_20, $__local_var_9_9, $r_20, $v1_19, $v_18) {
  $__num = \func_num_args();
  $__res = ((($__local_var_9_9)->{'bind'})((($v_18)($r_20))($s_21)))(function($v2_22) use ($Functor0_12_12, $Semigroup0_17_20, $r_20, $v1_19) {
  $__num = \func_num_args();
  $__local_var_23_21 = ($v2_22)->{'value2'};
  $__res = ((($Functor0_12_12)->{'map'})(function($v3_24) use ($Semigroup0_17_20, $__local_var_23_21, $v2_22) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v3_24)->{'value0'}, (($v2_22)->{'value1'})(($v3_24)->{'value1'}), ((($Semigroup0_17_20)->{'append'})($__local_var_23_21))(($v3_24)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v1_19)($r_20))(($v2_22)->{'value0'}));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_18) use ($functorRWST1_13_13) {
  $__num = \func_num_args();
  $__res = $functorRWST1_13_13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $bindRWST2_16_19 = (object)["bind" => function($v_18) use ($Functor0_10_10, $Semigroup0_16_19, $__local_var_9_9) {
  $__num = \func_num_args();
  $__res = function($f_19) use ($Functor0_10_10, $Semigroup0_16_19, $__local_var_9_9, $v_18) {
  $__num = \func_num_args();
  $__res = function($r_20) use ($Functor0_10_10, $Semigroup0_16_19, $__local_var_9_9, $f_19, $v_18) {
  $__num = \func_num_args();
  $__res = function($s_21) use ($Functor0_10_10, $Semigroup0_16_19, $__local_var_9_9, $f_19, $r_20, $v_18) {
  $__num = \func_num_args();
  $__res = ((($__local_var_9_9)->{'bind'})((($v_18)($r_20))($s_21)))(function($v1_22) use ($Functor0_10_10, $Semigroup0_16_19, $f_19, $r_20) {
  $__num = \func_num_args();
  $__local_var_23_23 = ($v1_22)->{'value2'};
  $__res = ((($Functor0_10_10)->{'map'})(function($v3_24) use ($Semigroup0_16_19, $__local_var_23_23) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v3_24)->{'value0'}, ($v3_24)->{'value1'}, ((($Semigroup0_16_19)->{'append'})($__local_var_23_23))(($v3_24)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($f_19)(($v1_22)->{'value1'}))($r_20))(($v1_22)->{'value0'}));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_18) use ($applyRWST2_17_20) {
  $__num = \func_num_args();
  $__res = $applyRWST2_17_20;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["Applicative0" => function($_dollar___unused_17) use ($applicativeRWST2_15_15) {
  $__num = \func_num_args();
  $__res = $applicativeRWST2_15_15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_17) use ($bindRWST2_16_19) {
  $__num = \func_num_args();
  $__res = $bindRWST2_16_19;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $monadTellRWST1_5_4 = function($dictMonoid_10) use ($monadRWST1_9_9, $pure_3_2) {
  $__num = \func_num_args();
  $Semigroup0_11_26 = (($dictMonoid_10)->{'Semigroup0'})(null);
  $monadRWST2_12_27 = ($monadRWST1_9_9)($dictMonoid_10);
  $__res = (object)["tell" => function($w_13) use ($pure_3_2) {
  $__num = \func_num_args();
  $__res = function($v_14) use ($pure_3_2, $w_13) {
  $__num = \func_num_args();
  $__res = function($s_15) use ($pure_3_2, $w_13) {
  $__num = \func_num_args();
  $__res = ($pure_3_2)(new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult($s_15, $GLOBALS['Data_Unit_unit'], $w_13));
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
}, "Semigroup0" => function($_dollar___unused_13) use ($Semigroup0_11_26) {
  $__num = \func_num_args();
  $__res = $Semigroup0_11_26;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad1" => function($_dollar___unused_13) use ($monadRWST2_12_27) {
  $__num = \func_num_args();
  $__res = $monadRWST2_12_27;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = function($dictMonoid_6) use ($Applicative0_2_1, $Bind1_1_0, $monadTellRWST1_5_4) {
  $__num = \func_num_args();
  $monadTellRWST2_7_29 = ($monadTellRWST1_5_4)($dictMonoid_6);
  $__res = (object)["listen" => function($m_8) use ($Applicative0_2_1, $Bind1_1_0) {
  $__num = \func_num_args();
  $__res = function($r_9) use ($Applicative0_2_1, $Bind1_1_0, $m_8) {
  $__num = \func_num_args();
  $__res = function($s_10) use ($Applicative0_2_1, $Bind1_1_0, $m_8, $r_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_1_0)->{'bind'})((($m_8)($r_9))($s_10)))(function($v_11) use ($Applicative0_2_1) {
  $__num = \func_num_args();
  $__res = (($Applicative0_2_1)->{'pure'})(new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v_11)->{'value0'}, new \Data\Tuple\Data_Tuple_Tuple(($v_11)->{'value1'}, ($v_11)->{'value2'}), ($v_11)->{'value2'}));
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
}, "pass" => function($m_8) use ($Applicative0_2_1, $Bind1_1_0) {
  $__num = \func_num_args();
  $__res = function($r_9) use ($Applicative0_2_1, $Bind1_1_0, $m_8) {
  $__num = \func_num_args();
  $__res = function($s_10) use ($Applicative0_2_1, $Bind1_1_0, $m_8, $r_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_1_0)->{'bind'})((($m_8)($r_9))($s_10)))(function($v_11) use ($Applicative0_2_1) {
  $__num = \func_num_args();
  $__res = (($Applicative0_2_1)->{'pure'})(new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v_11)->{'value0'}, (($v_11)->{'value1'})->{'value0'}, ((($v_11)->{'value1'})->{'value1'})(($v_11)->{'value2'})));
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
}, "Monoid0" => function($_dollar___unused_8) use ($dictMonoid_6) {
  $__num = \func_num_args();
  $__res = $dictMonoid_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "MonadTell1" => function($_dollar___unused_8) use ($monadTellRWST2_7_29) {
  $__num = \func_num_args();
  $__res = $monadTellRWST2_7_29;
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
$GLOBALS['Control_Monad_RWS_Trans_monadWriterRWST'] = __NAMESPACE__ . '\\majControl_majMonad_majRmajWmajS_majTrans_monadmajWritermajRmajWmajSmajT';

// Control_Monad_RWS_Trans_monadThrowRWST
function majControl_majMonad_majRmajWmajS_majTrans_monadmajThrowmajRmajWmajSmajT($dictMonadThrow_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majRmajWmajS_majTrans_monadmajThrowmajRmajWmajSmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Monad0_1_0 = (($dictMonadThrow_0)->{'Monad0'})(null);
  $__local_var_2_1 = (($dictMonadThrow_0)->{'Monad0'})(null);
  $pure_3_2 = ((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'};
  $__local_var_4_3 = (($__local_var_2_1)->{'Bind1'})(null);
  $Apply0_5_4 = (($__local_var_4_3)->{'Apply0'})(null);
  $Functor0_6_5 = (($Apply0_5_4)->{'Functor0'})(null);
  $__local_var_7_6 = (($Apply0_5_4)->{'Functor0'})(null);
  $functorRWST1_7_6 = (object)["map" => function($f_8) use ($__local_var_7_6) {
  $__num = \func_num_args();
  $__res = function($v_9) use ($__local_var_7_6, $f_8) {
  $__num = \func_num_args();
  $__res = function($r_10) use ($__local_var_7_6, $f_8, $v_9) {
  $__num = \func_num_args();
  $__res = function($s_11) use ($__local_var_7_6, $f_8, $r_10, $v_9) {
  $__num = \func_num_args();
  $__res = ((($__local_var_7_6)->{'map'})(function($v1_12) use ($f_8) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v1_12)->{'value0'}, ($f_8)(($v1_12)->{'value1'}), ($v1_12)->{'value2'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v_9)($r_10))($s_11));
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_8_8 = (($__local_var_2_1)->{'Bind1'})(null);
  $Functor0_9_9 = (((($__local_var_8_8)->{'Apply0'})(null))->{'Functor0'})(null);
  $Apply0_10_10 = (($__local_var_8_8)->{'Apply0'})(null);
  $Functor0_11_11 = (($Apply0_10_10)->{'Functor0'})(null);
  $__local_var_12_12 = (($Apply0_10_10)->{'Functor0'})(null);
  $functorRWST1_12_12 = (object)["map" => function($f_13) use ($__local_var_12_12) {
  $__num = \func_num_args();
  $__res = function($v_14) use ($__local_var_12_12, $f_13) {
  $__num = \func_num_args();
  $__res = function($r_15) use ($__local_var_12_12, $f_13, $v_14) {
  $__num = \func_num_args();
  $__res = function($s_16) use ($__local_var_12_12, $f_13, $r_15, $v_14) {
  $__num = \func_num_args();
  $__res = ((($__local_var_12_12)->{'map'})(function($v1_17) use ($f_13) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v1_17)->{'value0'}, ($f_13)(($v1_17)->{'value1'}), ($v1_17)->{'value2'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v_14)($r_15))($s_16));
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $monadRWST1_5_4 = function($dictMonoid_13) use ($Functor0_11_11, $Functor0_6_5, $Functor0_9_9, $__local_var_4_3, $__local_var_8_8, $functorRWST1_12_12, $functorRWST1_7_6, $pure_3_2) {
  $__num = \func_num_args();
  $Semigroup0_14_14 = (($dictMonoid_13)->{'Semigroup0'})(null);
  $applyRWST2_14_14 = (object)["apply" => function($v_15) use ($Functor0_6_5, $Semigroup0_14_14, $__local_var_4_3) {
  $__num = \func_num_args();
  $__res = function($v1_16) use ($Functor0_6_5, $Semigroup0_14_14, $__local_var_4_3, $v_15) {
  $__num = \func_num_args();
  $__res = function($r_17) use ($Functor0_6_5, $Semigroup0_14_14, $__local_var_4_3, $v1_16, $v_15) {
  $__num = \func_num_args();
  $__res = function($s_18) use ($Functor0_6_5, $Semigroup0_14_14, $__local_var_4_3, $r_17, $v1_16, $v_15) {
  $__num = \func_num_args();
  $__res = ((($__local_var_4_3)->{'bind'})((($v_15)($r_17))($s_18)))(function($v2_19) use ($Functor0_6_5, $Semigroup0_14_14, $r_17, $v1_16) {
  $__num = \func_num_args();
  $__local_var_20_15 = ($v2_19)->{'value2'};
  $__res = ((($Functor0_6_5)->{'map'})(function($v3_21) use ($Semigroup0_14_14, $__local_var_20_15, $v2_19) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v3_21)->{'value0'}, (($v2_19)->{'value1'})(($v3_21)->{'value1'}), ((($Semigroup0_14_14)->{'append'})($__local_var_20_15))(($v3_21)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v1_16)($r_17))(($v2_19)->{'value0'}));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_15) use ($functorRWST1_7_6) {
  $__num = \func_num_args();
  $__res = $functorRWST1_7_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeRWST2_14_14 = (object)["pure" => function($a_15) use ($dictMonoid_13, $pure_3_2) {
  $__num = \func_num_args();
  $__res = function($v_16) use ($a_15, $dictMonoid_13, $pure_3_2) {
  $__num = \func_num_args();
  $__res = function($s_17) use ($a_15, $dictMonoid_13, $pure_3_2) {
  $__num = \func_num_args();
  $__res = ($pure_3_2)(new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult($s_17, $a_15, ($dictMonoid_13)->{'mempty'}));
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
}, "Apply0" => function($_dollar___unused_15) use ($applyRWST2_14_14) {
  $__num = \func_num_args();
  $__res = $applyRWST2_14_14;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Semigroup0_15_18 = (($dictMonoid_13)->{'Semigroup0'})(null);
  $Semigroup0_16_19 = (($dictMonoid_13)->{'Semigroup0'})(null);
  $applyRWST2_16_19 = (object)["apply" => function($v_17) use ($Functor0_11_11, $Semigroup0_16_19, $__local_var_8_8) {
  $__num = \func_num_args();
  $__res = function($v1_18) use ($Functor0_11_11, $Semigroup0_16_19, $__local_var_8_8, $v_17) {
  $__num = \func_num_args();
  $__res = function($r_19) use ($Functor0_11_11, $Semigroup0_16_19, $__local_var_8_8, $v1_18, $v_17) {
  $__num = \func_num_args();
  $__res = function($s_20) use ($Functor0_11_11, $Semigroup0_16_19, $__local_var_8_8, $r_19, $v1_18, $v_17) {
  $__num = \func_num_args();
  $__res = ((($__local_var_8_8)->{'bind'})((($v_17)($r_19))($s_20)))(function($v2_21) use ($Functor0_11_11, $Semigroup0_16_19, $r_19, $v1_18) {
  $__num = \func_num_args();
  $__local_var_22_20 = ($v2_21)->{'value2'};
  $__res = ((($Functor0_11_11)->{'map'})(function($v3_23) use ($Semigroup0_16_19, $__local_var_22_20, $v2_21) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v3_23)->{'value0'}, (($v2_21)->{'value1'})(($v3_23)->{'value1'}), ((($Semigroup0_16_19)->{'append'})($__local_var_22_20))(($v3_23)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v1_18)($r_19))(($v2_21)->{'value0'}));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_17) use ($functorRWST1_12_12) {
  $__num = \func_num_args();
  $__res = $functorRWST1_12_12;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $bindRWST2_15_18 = (object)["bind" => function($v_17) use ($Functor0_9_9, $Semigroup0_15_18, $__local_var_8_8) {
  $__num = \func_num_args();
  $__res = function($f_18) use ($Functor0_9_9, $Semigroup0_15_18, $__local_var_8_8, $v_17) {
  $__num = \func_num_args();
  $__res = function($r_19) use ($Functor0_9_9, $Semigroup0_15_18, $__local_var_8_8, $f_18, $v_17) {
  $__num = \func_num_args();
  $__res = function($s_20) use ($Functor0_9_9, $Semigroup0_15_18, $__local_var_8_8, $f_18, $r_19, $v_17) {
  $__num = \func_num_args();
  $__res = ((($__local_var_8_8)->{'bind'})((($v_17)($r_19))($s_20)))(function($v1_21) use ($Functor0_9_9, $Semigroup0_15_18, $f_18, $r_19) {
  $__num = \func_num_args();
  $__local_var_22_22 = ($v1_21)->{'value2'};
  $__res = ((($Functor0_9_9)->{'map'})(function($v3_23) use ($Semigroup0_15_18, $__local_var_22_22) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v3_23)->{'value0'}, ($v3_23)->{'value1'}, ((($Semigroup0_15_18)->{'append'})($__local_var_22_22))(($v3_23)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($f_18)(($v1_21)->{'value1'}))($r_19))(($v1_21)->{'value0'}));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_17) use ($applyRWST2_16_19) {
  $__num = \func_num_args();
  $__res = $applyRWST2_16_19;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["Applicative0" => function($_dollar___unused_16) use ($applicativeRWST2_14_14) {
  $__num = \func_num_args();
  $__res = $applicativeRWST2_14_14;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_16) use ($bindRWST2_15_18) {
  $__num = \func_num_args();
  $__res = $bindRWST2_15_18;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = function($dictMonoid_6) use ($Monad0_1_0, $dictMonadThrow_0, $monadRWST1_5_4) {
  $__num = \func_num_args();
  $monadTransRWST1_7_25 = (object)["lift" => function($dictMonad_7) use ($dictMonoid_6) {
  $__num = \func_num_args();
  $Bind1_8_25 = (($dictMonad_7)->{'Bind1'})(null);
  $pure_9_26 = ((($dictMonad_7)->{'Applicative0'})(null))->{'pure'};
  $__res = function($m_10) use ($Bind1_8_25, $dictMonoid_6, $pure_9_26) {
  $__num = \func_num_args();
  $__res = function($v_11) use ($Bind1_8_25, $dictMonoid_6, $m_10, $pure_9_26) {
  $__num = \func_num_args();
  $__res = function($s_12) use ($Bind1_8_25, $dictMonoid_6, $m_10, $pure_9_26) {
  $__num = \func_num_args();
  $__res = ((($Bind1_8_25)->{'bind'})($m_10))(function($a_13) use ($dictMonoid_6, $pure_9_26, $s_12) {
  $__num = \func_num_args();
  $__res = ($pure_9_26)(new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult($s_12, $a_13, ($dictMonoid_6)->{'mempty'}));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $monadRWST2_8_28 = ($monadRWST1_5_4)($dictMonoid_6);
  $__res = (object)["throwError" => function($e_9) use ($Monad0_1_0, $dictMonadThrow_0, $monadTransRWST1_7_25) {
  $__num = \func_num_args();
  $__res = ((($monadTransRWST1_7_25)->{'lift'})($Monad0_1_0))((($dictMonadThrow_0)->{'throwError'})($e_9));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad0" => function($_dollar___unused_9) use ($monadRWST2_8_28) {
  $__num = \func_num_args();
  $__res = $monadRWST2_8_28;
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
$GLOBALS['Control_Monad_RWS_Trans_monadThrowRWST'] = __NAMESPACE__ . '\\majControl_majMonad_majRmajWmajS_majTrans_monadmajThrowmajRmajWmajSmajT';

// Control_Monad_RWS_Trans_monadErrorRWST
function majControl_majMonad_majRmajWmajS_majTrans_monadmajErrormajRmajWmajSmajT($dictMonadError_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majRmajWmajS_majTrans_monadmajErrormajRmajWmajSmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictMonadError_0)->{'MonadThrow0'})(null);
  $Monad0_2_1 = (($__local_var_1_0)->{'Monad0'})(null);
  $__local_var_3_2 = (($__local_var_1_0)->{'Monad0'})(null);
  $pure_4_3 = ((($__local_var_3_2)->{'Applicative0'})(null))->{'pure'};
  $__local_var_5_4 = (($__local_var_3_2)->{'Bind1'})(null);
  $Apply0_6_5 = (($__local_var_5_4)->{'Apply0'})(null);
  $Functor0_7_6 = (($Apply0_6_5)->{'Functor0'})(null);
  $__local_var_8_7 = (($Apply0_6_5)->{'Functor0'})(null);
  $functorRWST1_8_7 = (object)["map" => function($f_9) use ($__local_var_8_7) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($__local_var_8_7, $f_9) {
  $__num = \func_num_args();
  $__res = function($r_11) use ($__local_var_8_7, $f_9, $v_10) {
  $__num = \func_num_args();
  $__res = function($s_12) use ($__local_var_8_7, $f_9, $r_11, $v_10) {
  $__num = \func_num_args();
  $__res = ((($__local_var_8_7)->{'map'})(function($v1_13) use ($f_9) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v1_13)->{'value0'}, ($f_9)(($v1_13)->{'value1'}), ($v1_13)->{'value2'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v_10)($r_11))($s_12));
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_9_9 = (($__local_var_3_2)->{'Bind1'})(null);
  $Functor0_10_10 = (((($__local_var_9_9)->{'Apply0'})(null))->{'Functor0'})(null);
  $Apply0_11_11 = (($__local_var_9_9)->{'Apply0'})(null);
  $Functor0_12_12 = (($Apply0_11_11)->{'Functor0'})(null);
  $__local_var_13_13 = (($Apply0_11_11)->{'Functor0'})(null);
  $functorRWST1_13_13 = (object)["map" => function($f_14) use ($__local_var_13_13) {
  $__num = \func_num_args();
  $__res = function($v_15) use ($__local_var_13_13, $f_14) {
  $__num = \func_num_args();
  $__res = function($r_16) use ($__local_var_13_13, $f_14, $v_15) {
  $__num = \func_num_args();
  $__res = function($s_17) use ($__local_var_13_13, $f_14, $r_16, $v_15) {
  $__num = \func_num_args();
  $__res = ((($__local_var_13_13)->{'map'})(function($v1_18) use ($f_14) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v1_18)->{'value0'}, ($f_14)(($v1_18)->{'value1'}), ($v1_18)->{'value2'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v_15)($r_16))($s_17));
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $monadRWST1_8_7 = function($dictMonoid_14) use ($Functor0_10_10, $Functor0_12_12, $Functor0_7_6, $__local_var_5_4, $__local_var_9_9, $functorRWST1_13_13, $functorRWST1_8_7, $pure_4_3) {
  $__num = \func_num_args();
  $Semigroup0_15_15 = (($dictMonoid_14)->{'Semigroup0'})(null);
  $applyRWST2_15_15 = (object)["apply" => function($v_16) use ($Functor0_7_6, $Semigroup0_15_15, $__local_var_5_4) {
  $__num = \func_num_args();
  $__res = function($v1_17) use ($Functor0_7_6, $Semigroup0_15_15, $__local_var_5_4, $v_16) {
  $__num = \func_num_args();
  $__res = function($r_18) use ($Functor0_7_6, $Semigroup0_15_15, $__local_var_5_4, $v1_17, $v_16) {
  $__num = \func_num_args();
  $__res = function($s_19) use ($Functor0_7_6, $Semigroup0_15_15, $__local_var_5_4, $r_18, $v1_17, $v_16) {
  $__num = \func_num_args();
  $__res = ((($__local_var_5_4)->{'bind'})((($v_16)($r_18))($s_19)))(function($v2_20) use ($Functor0_7_6, $Semigroup0_15_15, $r_18, $v1_17) {
  $__num = \func_num_args();
  $__local_var_21_16 = ($v2_20)->{'value2'};
  $__res = ((($Functor0_7_6)->{'map'})(function($v3_22) use ($Semigroup0_15_15, $__local_var_21_16, $v2_20) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v3_22)->{'value0'}, (($v2_20)->{'value1'})(($v3_22)->{'value1'}), ((($Semigroup0_15_15)->{'append'})($__local_var_21_16))(($v3_22)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v1_17)($r_18))(($v2_20)->{'value0'}));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_16) use ($functorRWST1_8_7) {
  $__num = \func_num_args();
  $__res = $functorRWST1_8_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeRWST2_15_15 = (object)["pure" => function($a_16) use ($dictMonoid_14, $pure_4_3) {
  $__num = \func_num_args();
  $__res = function($v_17) use ($a_16, $dictMonoid_14, $pure_4_3) {
  $__num = \func_num_args();
  $__res = function($s_18) use ($a_16, $dictMonoid_14, $pure_4_3) {
  $__num = \func_num_args();
  $__res = ($pure_4_3)(new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult($s_18, $a_16, ($dictMonoid_14)->{'mempty'}));
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
}, "Apply0" => function($_dollar___unused_16) use ($applyRWST2_15_15) {
  $__num = \func_num_args();
  $__res = $applyRWST2_15_15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Semigroup0_16_19 = (($dictMonoid_14)->{'Semigroup0'})(null);
  $Semigroup0_17_20 = (($dictMonoid_14)->{'Semigroup0'})(null);
  $applyRWST2_17_20 = (object)["apply" => function($v_18) use ($Functor0_12_12, $Semigroup0_17_20, $__local_var_9_9) {
  $__num = \func_num_args();
  $__res = function($v1_19) use ($Functor0_12_12, $Semigroup0_17_20, $__local_var_9_9, $v_18) {
  $__num = \func_num_args();
  $__res = function($r_20) use ($Functor0_12_12, $Semigroup0_17_20, $__local_var_9_9, $v1_19, $v_18) {
  $__num = \func_num_args();
  $__res = function($s_21) use ($Functor0_12_12, $Semigroup0_17_20, $__local_var_9_9, $r_20, $v1_19, $v_18) {
  $__num = \func_num_args();
  $__res = ((($__local_var_9_9)->{'bind'})((($v_18)($r_20))($s_21)))(function($v2_22) use ($Functor0_12_12, $Semigroup0_17_20, $r_20, $v1_19) {
  $__num = \func_num_args();
  $__local_var_23_21 = ($v2_22)->{'value2'};
  $__res = ((($Functor0_12_12)->{'map'})(function($v3_24) use ($Semigroup0_17_20, $__local_var_23_21, $v2_22) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v3_24)->{'value0'}, (($v2_22)->{'value1'})(($v3_24)->{'value1'}), ((($Semigroup0_17_20)->{'append'})($__local_var_23_21))(($v3_24)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v1_19)($r_20))(($v2_22)->{'value0'}));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_18) use ($functorRWST1_13_13) {
  $__num = \func_num_args();
  $__res = $functorRWST1_13_13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $bindRWST2_16_19 = (object)["bind" => function($v_18) use ($Functor0_10_10, $Semigroup0_16_19, $__local_var_9_9) {
  $__num = \func_num_args();
  $__res = function($f_19) use ($Functor0_10_10, $Semigroup0_16_19, $__local_var_9_9, $v_18) {
  $__num = \func_num_args();
  $__res = function($r_20) use ($Functor0_10_10, $Semigroup0_16_19, $__local_var_9_9, $f_19, $v_18) {
  $__num = \func_num_args();
  $__res = function($s_21) use ($Functor0_10_10, $Semigroup0_16_19, $__local_var_9_9, $f_19, $r_20, $v_18) {
  $__num = \func_num_args();
  $__res = ((($__local_var_9_9)->{'bind'})((($v_18)($r_20))($s_21)))(function($v1_22) use ($Functor0_10_10, $Semigroup0_16_19, $f_19, $r_20) {
  $__num = \func_num_args();
  $__local_var_23_23 = ($v1_22)->{'value2'};
  $__res = ((($Functor0_10_10)->{'map'})(function($v3_24) use ($Semigroup0_16_19, $__local_var_23_23) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v3_24)->{'value0'}, ($v3_24)->{'value1'}, ((($Semigroup0_16_19)->{'append'})($__local_var_23_23))(($v3_24)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($f_19)(($v1_22)->{'value1'}))($r_20))(($v1_22)->{'value0'}));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_18) use ($applyRWST2_17_20) {
  $__num = \func_num_args();
  $__res = $applyRWST2_17_20;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["Applicative0" => function($_dollar___unused_17) use ($applicativeRWST2_15_15) {
  $__num = \func_num_args();
  $__res = $applicativeRWST2_15_15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_17) use ($bindRWST2_16_19) {
  $__num = \func_num_args();
  $__res = $bindRWST2_16_19;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $monadThrowRWST1_3_2 = function($dictMonoid_9) use ($Monad0_2_1, $__local_var_1_0, $monadRWST1_8_7) {
  $__num = \func_num_args();
  $monadTransRWST1_10_26 = (object)["lift" => function($dictMonad_10) use ($dictMonoid_9) {
  $__num = \func_num_args();
  $Bind1_11_26 = (($dictMonad_10)->{'Bind1'})(null);
  $pure_12_27 = ((($dictMonad_10)->{'Applicative0'})(null))->{'pure'};
  $__res = function($m_13) use ($Bind1_11_26, $dictMonoid_9, $pure_12_27) {
  $__num = \func_num_args();
  $__res = function($v_14) use ($Bind1_11_26, $dictMonoid_9, $m_13, $pure_12_27) {
  $__num = \func_num_args();
  $__res = function($s_15) use ($Bind1_11_26, $dictMonoid_9, $m_13, $pure_12_27) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_26)->{'bind'})($m_13))(function($a_16) use ($dictMonoid_9, $pure_12_27, $s_15) {
  $__num = \func_num_args();
  $__res = ($pure_12_27)(new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult($s_15, $a_16, ($dictMonoid_9)->{'mempty'}));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $monadRWST2_11_29 = ($monadRWST1_8_7)($dictMonoid_9);
  $__res = (object)["throwError" => function($e_12) use ($Monad0_2_1, $__local_var_1_0, $monadTransRWST1_10_26) {
  $__num = \func_num_args();
  $__res = ((($monadTransRWST1_10_26)->{'lift'})($Monad0_2_1))((($__local_var_1_0)->{'throwError'})($e_12));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad0" => function($_dollar___unused_12) use ($monadRWST2_11_29) {
  $__num = \func_num_args();
  $__res = $monadRWST2_11_29;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = function($dictMonoid_4) use ($dictMonadError_0, $monadThrowRWST1_3_2) {
  $__num = \func_num_args();
  $monadThrowRWST2_5_31 = ($monadThrowRWST1_3_2)($dictMonoid_4);
  $__res = (object)["catchError" => function($m_6) use ($dictMonadError_0) {
  $__num = \func_num_args();
  $__res = function($h_7) use ($dictMonadError_0, $m_6) {
  $__num = \func_num_args();
  $__res = function($r_8) use ($dictMonadError_0, $h_7, $m_6) {
  $__num = \func_num_args();
  $__res = function($s_9) use ($dictMonadError_0, $h_7, $m_6, $r_8) {
  $__num = \func_num_args();
  $__res = ((($dictMonadError_0)->{'catchError'})((($m_6)($r_8))($s_9)))(function($e_10) use ($h_7, $r_8, $s_9) {
  $__num = \func_num_args();
  $__res = ((($h_7)($e_10))($r_8))($s_9);
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "MonadThrow0" => function($_dollar___unused_6) use ($monadThrowRWST2_5_31) {
  $__num = \func_num_args();
  $__res = $monadThrowRWST2_5_31;
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
$GLOBALS['Control_Monad_RWS_Trans_monadErrorRWST'] = __NAMESPACE__ . '\\majControl_majMonad_majRmajWmajS_majTrans_monadmajErrormajRmajWmajSmajT';

// Control_Monad_RWS_Trans_monadSTRWST
function majControl_majMonad_majRmajWmajS_majTrans_monadmajSmajTmajRmajWmajSmajT($dictMonoid_0, $dictMonadST_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majRmajWmajS_majTrans_monadmajSmajTmajRmajWmajSmajT';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $Monad0_2_0 = (($dictMonadST_1)->{'Monad0'})(null);
  $pure_3_1 = ((($Monad0_2_0)->{'Applicative0'})(null))->{'pure'};
  $__local_var_4_2 = (($Monad0_2_0)->{'Bind1'})(null);
  $Apply0_5_3 = (($__local_var_4_2)->{'Apply0'})(null);
  $Functor0_6_4 = (($Apply0_5_3)->{'Functor0'})(null);
  $__local_var_7_5 = (($Apply0_5_3)->{'Functor0'})(null);
  $functorRWST1_7_5 = (object)["map" => function($f_8) use ($__local_var_7_5) {
  $__num = \func_num_args();
  $__res = function($v_9) use ($__local_var_7_5, $f_8) {
  $__num = \func_num_args();
  $__res = function($r_10) use ($__local_var_7_5, $f_8, $v_9) {
  $__num = \func_num_args();
  $__res = function($s_11) use ($__local_var_7_5, $f_8, $r_10, $v_9) {
  $__num = \func_num_args();
  $__res = ((($__local_var_7_5)->{'map'})(function($v1_12) use ($f_8) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v1_12)->{'value0'}, ($f_8)(($v1_12)->{'value1'}), ($v1_12)->{'value2'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v_9)($r_10))($s_11));
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_8_7 = (($Monad0_2_0)->{'Bind1'})(null);
  $Functor0_9_8 = (((($__local_var_8_7)->{'Apply0'})(null))->{'Functor0'})(null);
  $Apply0_10_9 = (($__local_var_8_7)->{'Apply0'})(null);
  $Functor0_11_10 = (($Apply0_10_9)->{'Functor0'})(null);
  $__local_var_12_11 = (($Apply0_10_9)->{'Functor0'})(null);
  $functorRWST1_12_11 = (object)["map" => function($f_13) use ($__local_var_12_11) {
  $__num = \func_num_args();
  $__res = function($v_14) use ($__local_var_12_11, $f_13) {
  $__num = \func_num_args();
  $__res = function($r_15) use ($__local_var_12_11, $f_13, $v_14) {
  $__num = \func_num_args();
  $__res = function($s_16) use ($__local_var_12_11, $f_13, $r_15, $v_14) {
  $__num = \func_num_args();
  $__res = ((($__local_var_12_11)->{'map'})(function($v1_17) use ($f_13) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v1_17)->{'value0'}, ($f_13)(($v1_17)->{'value1'}), ($v1_17)->{'value2'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v_14)($r_15))($s_16));
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Semigroup0_13_13 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $applyRWST2_13_13 = (object)["apply" => function($v_14) use ($Functor0_6_4, $Semigroup0_13_13, $__local_var_4_2) {
  $__num = \func_num_args();
  $__res = function($v1_15) use ($Functor0_6_4, $Semigroup0_13_13, $__local_var_4_2, $v_14) {
  $__num = \func_num_args();
  $__res = function($r_16) use ($Functor0_6_4, $Semigroup0_13_13, $__local_var_4_2, $v1_15, $v_14) {
  $__num = \func_num_args();
  $__res = function($s_17) use ($Functor0_6_4, $Semigroup0_13_13, $__local_var_4_2, $r_16, $v1_15, $v_14) {
  $__num = \func_num_args();
  $__res = ((($__local_var_4_2)->{'bind'})((($v_14)($r_16))($s_17)))(function($v2_18) use ($Functor0_6_4, $Semigroup0_13_13, $r_16, $v1_15) {
  $__num = \func_num_args();
  $__local_var_19_14 = ($v2_18)->{'value2'};
  $__res = ((($Functor0_6_4)->{'map'})(function($v3_20) use ($Semigroup0_13_13, $__local_var_19_14, $v2_18) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v3_20)->{'value0'}, (($v2_18)->{'value1'})(($v3_20)->{'value1'}), ((($Semigroup0_13_13)->{'append'})($__local_var_19_14))(($v3_20)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v1_15)($r_16))(($v2_18)->{'value0'}));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_14) use ($functorRWST1_7_5) {
  $__num = \func_num_args();
  $__res = $functorRWST1_7_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeRWST2_13_13 = (object)["pure" => function($a_14) use ($dictMonoid_0, $pure_3_1) {
  $__num = \func_num_args();
  $__res = function($v_15) use ($a_14, $dictMonoid_0, $pure_3_1) {
  $__num = \func_num_args();
  $__res = function($s_16) use ($a_14, $dictMonoid_0, $pure_3_1) {
  $__num = \func_num_args();
  $__res = ($pure_3_1)(new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult($s_16, $a_14, ($dictMonoid_0)->{'mempty'}));
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
}, "Apply0" => function($_dollar___unused_14) use ($applyRWST2_13_13) {
  $__num = \func_num_args();
  $__res = $applyRWST2_13_13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Semigroup0_14_17 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $Semigroup0_15_18 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $applyRWST2_15_18 = (object)["apply" => function($v_16) use ($Functor0_11_10, $Semigroup0_15_18, $__local_var_8_7) {
  $__num = \func_num_args();
  $__res = function($v1_17) use ($Functor0_11_10, $Semigroup0_15_18, $__local_var_8_7, $v_16) {
  $__num = \func_num_args();
  $__res = function($r_18) use ($Functor0_11_10, $Semigroup0_15_18, $__local_var_8_7, $v1_17, $v_16) {
  $__num = \func_num_args();
  $__res = function($s_19) use ($Functor0_11_10, $Semigroup0_15_18, $__local_var_8_7, $r_18, $v1_17, $v_16) {
  $__num = \func_num_args();
  $__res = ((($__local_var_8_7)->{'bind'})((($v_16)($r_18))($s_19)))(function($v2_20) use ($Functor0_11_10, $Semigroup0_15_18, $r_18, $v1_17) {
  $__num = \func_num_args();
  $__local_var_21_19 = ($v2_20)->{'value2'};
  $__res = ((($Functor0_11_10)->{'map'})(function($v3_22) use ($Semigroup0_15_18, $__local_var_21_19, $v2_20) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v3_22)->{'value0'}, (($v2_20)->{'value1'})(($v3_22)->{'value1'}), ((($Semigroup0_15_18)->{'append'})($__local_var_21_19))(($v3_22)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v1_17)($r_18))(($v2_20)->{'value0'}));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_16) use ($functorRWST1_12_11) {
  $__num = \func_num_args();
  $__res = $functorRWST1_12_11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $bindRWST2_14_17 = (object)["bind" => function($v_16) use ($Functor0_9_8, $Semigroup0_14_17, $__local_var_8_7) {
  $__num = \func_num_args();
  $__res = function($f_17) use ($Functor0_9_8, $Semigroup0_14_17, $__local_var_8_7, $v_16) {
  $__num = \func_num_args();
  $__res = function($r_18) use ($Functor0_9_8, $Semigroup0_14_17, $__local_var_8_7, $f_17, $v_16) {
  $__num = \func_num_args();
  $__res = function($s_19) use ($Functor0_9_8, $Semigroup0_14_17, $__local_var_8_7, $f_17, $r_18, $v_16) {
  $__num = \func_num_args();
  $__res = ((($__local_var_8_7)->{'bind'})((($v_16)($r_18))($s_19)))(function($v1_20) use ($Functor0_9_8, $Semigroup0_14_17, $f_17, $r_18) {
  $__num = \func_num_args();
  $__local_var_21_21 = ($v1_20)->{'value2'};
  $__res = ((($Functor0_9_8)->{'map'})(function($v3_22) use ($Semigroup0_14_17, $__local_var_21_21) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v3_22)->{'value0'}, ($v3_22)->{'value1'}, ((($Semigroup0_14_17)->{'append'})($__local_var_21_21))(($v3_22)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($f_17)(($v1_20)->{'value1'}))($r_18))(($v1_20)->{'value0'}));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_16) use ($applyRWST2_15_18) {
  $__num = \func_num_args();
  $__res = $applyRWST2_15_18;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $monadRWST1_3_1 = (object)["Applicative0" => function($_dollar___unused_15) use ($applicativeRWST2_13_13) {
  $__num = \func_num_args();
  $__res = $applicativeRWST2_13_13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_15) use ($bindRWST2_14_17) {
  $__num = \func_num_args();
  $__res = $bindRWST2_14_17;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_4_24 = (($Monad0_2_0)->{'Bind1'})(null);
  $pure_5_25 = ((($Monad0_2_0)->{'Applicative0'})(null))->{'pure'};
  $__res = (object)["liftST" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($m_6) use ($Bind1_4_24, $dictMonoid_0, $pure_5_25) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($Bind1_4_24, $dictMonoid_0, $m_6, $pure_5_25) {
  $__num = \func_num_args();
  $__res = function($s_8) use ($Bind1_4_24, $dictMonoid_0, $m_6, $pure_5_25) {
  $__num = \func_num_args();
  $__res = ((($Bind1_4_24)->{'bind'})($m_6))(function($a_9) use ($dictMonoid_0, $pure_5_25, $s_8) {
  $__num = \func_num_args();
  $__res = ($pure_5_25)(new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult($s_8, $a_9, ($dictMonoid_0)->{'mempty'}));
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
}))(($dictMonadST_1)->{'liftST'}), "Monad0" => function($_dollar___unused_4) use ($monadRWST1_3_1) {
  $__num = \func_num_args();
  $__res = $monadRWST1_3_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Control_Monad_RWS_Trans_monadSTRWST'] = __NAMESPACE__ . '\\majControl_majMonad_majRmajWmajS_majTrans_monadmajSmajTmajRmajWmajSmajT';

// Control_Monad_RWS_Trans_monoidRWST
function majControl_majMonad_majRmajWmajS_majTrans_monoidmajRmajWmajSmajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majRmajWmajS_majTrans_monoidmajRmajWmajSmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $pure_1_0 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $__local_var_2_1 = (($dictMonad_0)->{'Bind1'})(null);
  $Apply0_3_2 = (($__local_var_2_1)->{'Apply0'})(null);
  $Functor0_4_3 = (($Apply0_3_2)->{'Functor0'})(null);
  $__local_var_5_4 = (($Apply0_3_2)->{'Functor0'})(null);
  $functorRWST1_5_4 = (object)["map" => function($f_6) use ($__local_var_5_4) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($__local_var_5_4, $f_6) {
  $__num = \func_num_args();
  $__res = function($r_8) use ($__local_var_5_4, $f_6, $v_7) {
  $__num = \func_num_args();
  $__res = function($s_9) use ($__local_var_5_4, $f_6, $r_8, $v_7) {
  $__num = \func_num_args();
  $__res = ((($__local_var_5_4)->{'map'})(function($v1_10) use ($f_6) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v1_10)->{'value0'}, ($f_6)(($v1_10)->{'value1'}), ($v1_10)->{'value2'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v_7)($r_8))($s_9));
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_6_6 = (($dictMonad_0)->{'Bind1'})(null);
  $Apply0_7_7 = (($__local_var_6_6)->{'Apply0'})(null);
  $Functor0_8_8 = (($Apply0_7_7)->{'Functor0'})(null);
  $__local_var_9_9 = (($Apply0_7_7)->{'Functor0'})(null);
  $functorRWST1_9_9 = (object)["map" => function($f_10) use ($__local_var_9_9) {
  $__num = \func_num_args();
  $__res = function($v_11) use ($__local_var_9_9, $f_10) {
  $__num = \func_num_args();
  $__res = function($r_12) use ($__local_var_9_9, $f_10, $v_11) {
  $__num = \func_num_args();
  $__res = function($s_13) use ($__local_var_9_9, $f_10, $r_12, $v_11) {
  $__num = \func_num_args();
  $__res = ((($__local_var_9_9)->{'map'})(function($v1_14) use ($f_10) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v1_14)->{'value0'}, ($f_10)(($v1_14)->{'value1'}), ($v1_14)->{'value2'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v_11)($r_12))($s_13));
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = function($dictMonoid_10) use ($Functor0_4_3, $Functor0_8_8, $__local_var_2_1, $__local_var_6_6, $functorRWST1_5_4, $functorRWST1_9_9, $pure_1_0) {
  $__num = \func_num_args();
  $Semigroup0_11_11 = (($dictMonoid_10)->{'Semigroup0'})(null);
  $applyRWST2_11_11 = (object)["apply" => function($v_12) use ($Functor0_4_3, $Semigroup0_11_11, $__local_var_2_1) {
  $__num = \func_num_args();
  $__res = function($v1_13) use ($Functor0_4_3, $Semigroup0_11_11, $__local_var_2_1, $v_12) {
  $__num = \func_num_args();
  $__res = function($r_14) use ($Functor0_4_3, $Semigroup0_11_11, $__local_var_2_1, $v1_13, $v_12) {
  $__num = \func_num_args();
  $__res = function($s_15) use ($Functor0_4_3, $Semigroup0_11_11, $__local_var_2_1, $r_14, $v1_13, $v_12) {
  $__num = \func_num_args();
  $__res = ((($__local_var_2_1)->{'bind'})((($v_12)($r_14))($s_15)))(function($v2_16) use ($Functor0_4_3, $Semigroup0_11_11, $r_14, $v1_13) {
  $__num = \func_num_args();
  $__local_var_17_12 = ($v2_16)->{'value2'};
  $__res = ((($Functor0_4_3)->{'map'})(function($v3_18) use ($Semigroup0_11_11, $__local_var_17_12, $v2_16) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v3_18)->{'value0'}, (($v2_16)->{'value1'})(($v3_18)->{'value1'}), ((($Semigroup0_11_11)->{'append'})($__local_var_17_12))(($v3_18)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v1_13)($r_14))(($v2_16)->{'value0'}));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_12) use ($functorRWST1_5_4) {
  $__num = \func_num_args();
  $__res = $functorRWST1_5_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeRWST2_11_11 = (object)["pure" => function($a_12) use ($dictMonoid_10, $pure_1_0) {
  $__num = \func_num_args();
  $__res = function($v_13) use ($a_12, $dictMonoid_10, $pure_1_0) {
  $__num = \func_num_args();
  $__res = function($s_14) use ($a_12, $dictMonoid_10, $pure_1_0) {
  $__num = \func_num_args();
  $__res = ($pure_1_0)(new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult($s_14, $a_12, ($dictMonoid_10)->{'mempty'}));
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
}, "Apply0" => function($_dollar___unused_12) use ($applyRWST2_11_11) {
  $__num = \func_num_args();
  $__res = $applyRWST2_11_11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Semigroup0_12_15 = (($dictMonoid_10)->{'Semigroup0'})(null);
  $applyRWST2_12_15 = (object)["apply" => function($v_13) use ($Functor0_8_8, $Semigroup0_12_15, $__local_var_6_6) {
  $__num = \func_num_args();
  $__res = function($v1_14) use ($Functor0_8_8, $Semigroup0_12_15, $__local_var_6_6, $v_13) {
  $__num = \func_num_args();
  $__res = function($r_15) use ($Functor0_8_8, $Semigroup0_12_15, $__local_var_6_6, $v1_14, $v_13) {
  $__num = \func_num_args();
  $__res = function($s_16) use ($Functor0_8_8, $Semigroup0_12_15, $__local_var_6_6, $r_15, $v1_14, $v_13) {
  $__num = \func_num_args();
  $__res = ((($__local_var_6_6)->{'bind'})((($v_13)($r_15))($s_16)))(function($v2_17) use ($Functor0_8_8, $Semigroup0_12_15, $r_15, $v1_14) {
  $__num = \func_num_args();
  $__local_var_18_16 = ($v2_17)->{'value2'};
  $__res = ((($Functor0_8_8)->{'map'})(function($v3_19) use ($Semigroup0_12_15, $__local_var_18_16, $v2_17) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v3_19)->{'value0'}, (($v2_17)->{'value1'})(($v3_19)->{'value1'}), ((($Semigroup0_12_15)->{'append'})($__local_var_18_16))(($v3_19)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v1_14)($r_15))(($v2_17)->{'value0'}));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_13) use ($functorRWST1_9_9) {
  $__num = \func_num_args();
  $__res = $functorRWST1_9_9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = function($dictMonoid1_13) use ($applicativeRWST2_11_11, $applyRWST2_12_15) {
  $__num = \func_num_args();
  $Functor0_14_18 = (($applyRWST2_12_15)->{'Functor0'})(null);
  $__local_var_15_19 = ((($dictMonoid1_13)->{'Semigroup0'})(null))->{'append'};
  $semigroupRWST3_14_18 = (object)["append" => function($a_16) use ($Functor0_14_18, $__local_var_15_19, $applyRWST2_12_15) {
  $__num = \func_num_args();
  $__res = function($b_17) use ($Functor0_14_18, $__local_var_15_19, $a_16, $applyRWST2_12_15) {
  $__num = \func_num_args();
  $__res = ((($applyRWST2_12_15)->{'apply'})(((($Functor0_14_18)->{'map'})($__local_var_15_19))($a_16)))($b_17);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["mempty" => (($applicativeRWST2_11_11)->{'pure'})(($dictMonoid1_13)->{'mempty'}), "Semigroup0" => function($_dollar___unused_15) use ($semigroupRWST3_14_18) {
  $__num = \func_num_args();
  $__res = $semigroupRWST3_14_18;
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_RWS_Trans_monoidRWST'] = __NAMESPACE__ . '\\majControl_majMonad_majRmajWmajS_majTrans_monoidmajRmajWmajSmajT';

// Control_Monad_RWS_Trans_altRWST
function majControl_majMonad_majRmajWmajS_majTrans_altmajRmajWmajSmajT($dictAlt_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majRmajWmajS_majTrans_altmajRmajWmajSmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictAlt_0)->{'Functor0'})(null);
  $functorRWST1_1_0 = (object)["map" => function($f_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($__local_var_1_0, $f_2) {
  $__num = \func_num_args();
  $__res = function($r_4) use ($__local_var_1_0, $f_2, $v_3) {
  $__num = \func_num_args();
  $__res = function($s_5) use ($__local_var_1_0, $f_2, $r_4, $v_3) {
  $__num = \func_num_args();
  $__res = ((($__local_var_1_0)->{'map'})(function($v1_6) use ($f_2) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v1_6)->{'value0'}, ($f_2)(($v1_6)->{'value1'}), ($v1_6)->{'value2'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v_3)($r_4))($s_5));
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["alt" => function($v_2) use ($dictAlt_0) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($dictAlt_0, $v_2) {
  $__num = \func_num_args();
  $__res = function($r_4) use ($dictAlt_0, $v1_3, $v_2) {
  $__num = \func_num_args();
  $__res = function($s_5) use ($dictAlt_0, $r_4, $v1_3, $v_2) {
  $__num = \func_num_args();
  $__res = ((($dictAlt_0)->{'alt'})((($v_2)($r_4))($s_5)))((($v1_3)($r_4))($s_5));
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_2) use ($functorRWST1_1_0) {
  $__num = \func_num_args();
  $__res = $functorRWST1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_RWS_Trans_altRWST'] = __NAMESPACE__ . '\\majControl_majMonad_majRmajWmajS_majTrans_altmajRmajWmajSmajT';

// Control_Monad_RWS_Trans_plusRWST
function majControl_majMonad_majRmajWmajS_majTrans_plusmajRmajWmajSmajT($dictPlus_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majRmajWmajS_majTrans_plusmajRmajWmajSmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $empty_1_0 = ($dictPlus_0)->{'empty'};
  $__local_var_2_1 = (($dictPlus_0)->{'Alt0'})(null);
  $__local_var_3_2 = (($__local_var_2_1)->{'Functor0'})(null);
  $functorRWST1_3_2 = (object)["map" => function($f_4) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__res = function($v_5) use ($__local_var_3_2, $f_4) {
  $__num = \func_num_args();
  $__res = function($r_6) use ($__local_var_3_2, $f_4, $v_5) {
  $__num = \func_num_args();
  $__res = function($s_7) use ($__local_var_3_2, $f_4, $r_6, $v_5) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_2)->{'map'})(function($v1_8) use ($f_4) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v1_8)->{'value0'}, ($f_4)(($v1_8)->{'value1'}), ($v1_8)->{'value2'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v_5)($r_6))($s_7));
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $altRWST1_2_1 = (object)["alt" => function($v_4) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = function($v1_5) use ($__local_var_2_1, $v_4) {
  $__num = \func_num_args();
  $__res = function($r_6) use ($__local_var_2_1, $v1_5, $v_4) {
  $__num = \func_num_args();
  $__res = function($s_7) use ($__local_var_2_1, $r_6, $v1_5, $v_4) {
  $__num = \func_num_args();
  $__res = ((($__local_var_2_1)->{'alt'})((($v_4)($r_6))($s_7)))((($v1_5)($r_6))($s_7));
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_4) use ($functorRWST1_3_2) {
  $__num = \func_num_args();
  $__res = $functorRWST1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["empty" => function($v_3) use ($empty_1_0) {
  $__num = \func_num_args();
  $__res = function($v1_4) use ($empty_1_0) {
  $__num = \func_num_args();
  $__res = $empty_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Alt0" => function($_dollar___unused_3) use ($altRWST1_2_1) {
  $__num = \func_num_args();
  $__res = $altRWST1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_RWS_Trans_plusRWST'] = __NAMESPACE__ . '\\majControl_majMonad_majRmajWmajS_majTrans_plusmajRmajWmajSmajT';

// Control_Monad_RWS_Trans_alternativeRWST
function majControl_majMonad_majRmajWmajS_majTrans_alternativemajRmajWmajSmajT($dictMonoid_0, $dictAlternative_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majRmajWmajS_majTrans_alternativemajRmajWmajSmajT';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_2_0 = (($dictAlternative_1)->{'Plus1'})(null);
  $empty_3_1 = ($__local_var_2_0)->{'empty'};
  $__local_var_4_2 = (($__local_var_2_0)->{'Alt0'})(null);
  $__local_var_5_3 = (($__local_var_4_2)->{'Functor0'})(null);
  $functorRWST1_5_3 = (object)["map" => function($f_6) use ($__local_var_5_3) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($__local_var_5_3, $f_6) {
  $__num = \func_num_args();
  $__res = function($r_8) use ($__local_var_5_3, $f_6, $v_7) {
  $__num = \func_num_args();
  $__res = function($s_9) use ($__local_var_5_3, $f_6, $r_8, $v_7) {
  $__num = \func_num_args();
  $__res = ((($__local_var_5_3)->{'map'})(function($v1_10) use ($f_6) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v1_10)->{'value0'}, ($f_6)(($v1_10)->{'value1'}), ($v1_10)->{'value2'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v_7)($r_8))($s_9));
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $altRWST1_4_2 = (object)["alt" => function($v_6) use ($__local_var_4_2) {
  $__num = \func_num_args();
  $__res = function($v1_7) use ($__local_var_4_2, $v_6) {
  $__num = \func_num_args();
  $__res = function($r_8) use ($__local_var_4_2, $v1_7, $v_6) {
  $__num = \func_num_args();
  $__res = function($s_9) use ($__local_var_4_2, $r_8, $v1_7, $v_6) {
  $__num = \func_num_args();
  $__res = ((($__local_var_4_2)->{'alt'})((($v_6)($r_8))($s_9)))((($v1_7)($r_8))($s_9));
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_6) use ($functorRWST1_5_3) {
  $__num = \func_num_args();
  $__res = $functorRWST1_5_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $plusRWST1_2_0 = (object)["empty" => function($v_5) use ($empty_3_1) {
  $__num = \func_num_args();
  $__res = function($v1_6) use ($empty_3_1) {
  $__num = \func_num_args();
  $__res = $empty_3_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Alt0" => function($_dollar___unused_5) use ($altRWST1_4_2) {
  $__num = \func_num_args();
  $__res = $altRWST1_4_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = function($dictMonad_3) use ($dictMonoid_0, $plusRWST1_2_0) {
  $__num = \func_num_args();
  $pure_4_7 = ((($dictMonad_3)->{'Applicative0'})(null))->{'pure'};
  $__local_var_5_8 = (($dictMonad_3)->{'Bind1'})(null);
  $Apply0_6_9 = (($__local_var_5_8)->{'Apply0'})(null);
  $Functor0_7_10 = (($Apply0_6_9)->{'Functor0'})(null);
  $__local_var_8_11 = (($Apply0_6_9)->{'Functor0'})(null);
  $functorRWST1_8_11 = (object)["map" => function($f_9) use ($__local_var_8_11) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($__local_var_8_11, $f_9) {
  $__num = \func_num_args();
  $__res = function($r_11) use ($__local_var_8_11, $f_9, $v_10) {
  $__num = \func_num_args();
  $__res = function($s_12) use ($__local_var_8_11, $f_9, $r_11, $v_10) {
  $__num = \func_num_args();
  $__res = ((($__local_var_8_11)->{'map'})(function($v1_13) use ($f_9) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v1_13)->{'value0'}, ($f_9)(($v1_13)->{'value1'}), ($v1_13)->{'value2'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v_10)($r_11))($s_12));
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Semigroup0_9_13 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $applyRWST2_9_13 = (object)["apply" => function($v_10) use ($Functor0_7_10, $Semigroup0_9_13, $__local_var_5_8) {
  $__num = \func_num_args();
  $__res = function($v1_11) use ($Functor0_7_10, $Semigroup0_9_13, $__local_var_5_8, $v_10) {
  $__num = \func_num_args();
  $__res = function($r_12) use ($Functor0_7_10, $Semigroup0_9_13, $__local_var_5_8, $v1_11, $v_10) {
  $__num = \func_num_args();
  $__res = function($s_13) use ($Functor0_7_10, $Semigroup0_9_13, $__local_var_5_8, $r_12, $v1_11, $v_10) {
  $__num = \func_num_args();
  $__res = ((($__local_var_5_8)->{'bind'})((($v_10)($r_12))($s_13)))(function($v2_14) use ($Functor0_7_10, $Semigroup0_9_13, $r_12, $v1_11) {
  $__num = \func_num_args();
  $__local_var_15_14 = ($v2_14)->{'value2'};
  $__res = ((($Functor0_7_10)->{'map'})(function($v3_16) use ($Semigroup0_9_13, $__local_var_15_14, $v2_14) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult(($v3_16)->{'value0'}, (($v2_14)->{'value1'})(($v3_16)->{'value1'}), ((($Semigroup0_9_13)->{'append'})($__local_var_15_14))(($v3_16)->{'value2'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v1_11)($r_12))(($v2_14)->{'value0'}));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_10) use ($functorRWST1_8_11) {
  $__num = \func_num_args();
  $__res = $functorRWST1_8_11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeRWST1_4_7 = (object)["pure" => function($a_10) use ($dictMonoid_0, $pure_4_7) {
  $__num = \func_num_args();
  $__res = function($v_11) use ($a_10, $dictMonoid_0, $pure_4_7) {
  $__num = \func_num_args();
  $__res = function($s_12) use ($a_10, $dictMonoid_0, $pure_4_7) {
  $__num = \func_num_args();
  $__res = ($pure_4_7)(new \Control\Monad\RWS\Trans\Control_Monad_RWS_Trans_RWSResult($s_12, $a_10, ($dictMonoid_0)->{'mempty'}));
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
}, "Apply0" => function($_dollar___unused_10) use ($applyRWST2_9_13) {
  $__num = \func_num_args();
  $__res = $applyRWST2_9_13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["Applicative0" => function($_dollar___unused_5) use ($applicativeRWST1_4_7) {
  $__num = \func_num_args();
  $__res = $applicativeRWST1_4_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Plus1" => function($_dollar___unused_5) use ($plusRWST1_2_0) {
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
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Control_Monad_RWS_Trans_alternativeRWST'] = __NAMESPACE__ . '\\majControl_majMonad_majRmajWmajS_majTrans_alternativemajRmajWmajSmajT';

