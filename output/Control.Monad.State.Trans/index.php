<?php

namespace Control\Monad\State\Trans;

// ALL IMPORTS: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Lazy, Control.Monad, Control.Monad.Cont.Class, Control.Monad.Error.Class, Control.Monad.Reader.Class, Control.Monad.Rec.Class, Control.Monad.ST.Class, Control.Monad.State.Class, Control.Monad.State.Trans, Control.Monad.Trans.Class, Control.Monad.Writer.Class, Control.MonadPlus, Control.Plus, Control.Semigroupoid, Data.Function, Data.Functor, Data.Monoid, Data.Newtype, Data.Semigroup, Data.Tuple, Data.Unit, Effect.Class, Prelude, Prim
// TO REQUIRE: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Lazy, Control.Monad, Control.Monad.Cont.Class, Control.Monad.Error.Class, Control.Monad.Reader.Class, Control.Monad.Rec.Class, Control.Monad.ST.Class, Control.Monad.State.Class, Control.Monad.State.Trans, Control.Monad.Trans.Class, Control.Monad.Writer.Class, Control.MonadPlus, Control.Plus, Control.Semigroupoid, Data.Function, Data.Functor, Data.Monoid, Data.Newtype, Data.Semigroup, Data.Tuple, Data.Unit, Effect.Class, Prelude
require_once __DIR__ . '/../Control.Alt/index.php';
require_once __DIR__ . '/../Control.Alternative/index.php';
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Lazy/index.php';
require_once __DIR__ . '/../Control.Monad/index.php';
require_once __DIR__ . '/../Control.Monad.Cont.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Error.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Reader.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Rec.Class/index.php';
require_once __DIR__ . '/../Control.Monad.ST.Class/index.php';
require_once __DIR__ . '/../Control.Monad.State.Class/index.php';
require_once __DIR__ . '/../Control.Monad.State.Trans/index.php';
require_once __DIR__ . '/../Control.Monad.Trans.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Writer.Class/index.php';
require_once __DIR__ . '/../Control.MonadPlus/index.php';
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




// Control_Monad_State_Trans_StateT
function majControl_majMonad_majState_majTrans_majStatemajT($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majState_majTrans_majStatemajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_State_Trans_StateT'] = __NAMESPACE__ . '\\majControl_majMonad_majState_majTrans_majStatemajT';

// Control_Monad_State_Trans_withStateT
function majControl_majMonad_majState_majTrans_withmajStatemajT($f_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majState_majTrans_withmajStatemajT';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($v_1))($f_0);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Control_Monad_State_Trans_withStateT'] = __NAMESPACE__ . '\\majControl_majMonad_majState_majTrans_withmajStatemajT';

// Control_Monad_State_Trans_runStateT
function majControl_majMonad_majState_majTrans_runmajStatemajT($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majState_majTrans_runmajStatemajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $v_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_State_Trans_runStateT'] = __NAMESPACE__ . '\\majControl_majMonad_majState_majTrans_runmajStatemajT';

// Control_Monad_State_Trans_newtypeStateT
$GLOBALS['Control_Monad_State_Trans_newtypeStateT'] = (object)["Coercible0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Control_Monad_State_Trans_monadTransStateT
$GLOBALS['Control_Monad_State_Trans_monadTransStateT'] = (object)["lift" => function($dictMonad_0) {
  $__num = \func_num_args();
  $Bind1_1_0 = (($dictMonad_0)->{'Bind1'})(null);
  $pure_2_1 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $__res = function($m_3) use ($Bind1_1_0, $pure_2_1) {
  $__num = \func_num_args();
  $__res = function($s_4) use ($Bind1_1_0, $m_3, $pure_2_1) {
  $__num = \func_num_args();
  $__res = ((($Bind1_1_0)->{'bind'})($m_3))(function($x_5) use ($pure_2_1, $s_4) {
  $__num = \func_num_args();
  $__res = ($pure_2_1)(new \Data\Tuple\Data_Tuple_Tuple($x_5, $s_4));
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

// Control_Monad_State_Trans_mapStateT
function majControl_majMonad_majState_majTrans_mapmajStatemajT($f_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majState_majTrans_mapmajStatemajT';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($f_0))($v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Control_Monad_State_Trans_mapStateT'] = __NAMESPACE__ . '\\majControl_majMonad_majState_majTrans_mapmajStatemajT';

// Control_Monad_State_Trans_lazyStateT
$GLOBALS['Control_Monad_State_Trans_lazyStateT'] = (object)["defer" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($s_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = (($f_0)($GLOBALS['Data_Unit_unit']))($s_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Control_Monad_State_Trans_functorStateT
function majControl_majMonad_majState_majTrans_functormajStatemajT($dictFunctor_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majState_majTrans_functormajStatemajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["map" => function($f_1) use ($dictFunctor_0) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($dictFunctor_0, $f_1) {
  $__num = \func_num_args();
  $__res = function($s_3) use ($dictFunctor_0, $f_1, $v_2) {
  $__num = \func_num_args();
  $__res = ((($dictFunctor_0)->{'map'})(function($v1_4) use ($f_1) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_1)(($v1_4)->{'value0'}), ($v1_4)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_2)($s_3));
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
$GLOBALS['Control_Monad_State_Trans_functorStateT'] = __NAMESPACE__ . '\\majControl_majMonad_majState_majTrans_functormajStatemajT';

// Control_Monad_State_Trans_execStateT
function majControl_majMonad_majState_majTrans_execmajStatemajT($dictFunctor_0, $v_1 = null, $s_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majState_majTrans_execmajStatemajT';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictFunctor_0)->{'map'})($GLOBALS['Data_Tuple_snd']))(($v_1)($s_2));
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Control_Monad_State_Trans_execStateT'] = __NAMESPACE__ . '\\majControl_majMonad_majState_majTrans_execmajStatemajT';

// Control_Monad_State_Trans_evalStateT
function majControl_majMonad_majState_majTrans_evalmajStatemajT($dictFunctor_0, $v_1 = null, $s_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majState_majTrans_evalmajStatemajT';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictFunctor_0)->{'map'})($GLOBALS['Data_Tuple_fst']))(($v_1)($s_2));
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Control_Monad_State_Trans_evalStateT'] = __NAMESPACE__ . '\\majControl_majMonad_majState_majTrans_evalmajStatemajT';

// Control_Monad_State_Trans_monadStateT
function majControl_majMonad_majState_majTrans_monadmajStatemajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majState_majTrans_monadmajStatemajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["Applicative0" => function($_dollar___unused_1) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_1) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_bindStateT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_State_Trans_monadStateT'] = __NAMESPACE__ . '\\majControl_majMonad_majState_majTrans_monadmajStatemajT';

// Control_Monad_State_Trans_bindStateT
function majControl_majMonad_majState_majTrans_bindmajStatemajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majState_majTrans_bindmajStatemajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Bind1_1_0 = (($dictMonad_0)->{'Bind1'})(null);
  $__res = (object)["bind" => function($v_2) use ($Bind1_1_0) {
  $__num = \func_num_args();
  $__res = function($f_3) use ($Bind1_1_0, $v_2) {
  $__num = \func_num_args();
  $__res = function($s_4) use ($Bind1_1_0, $f_3, $v_2) {
  $__num = \func_num_args();
  $__res = ((($Bind1_1_0)->{'bind'})(($v_2)($s_4)))(function($v1_5) use ($f_3) {
  $__num = \func_num_args();
  $__res = (($f_3)(($v1_5)->{'value0'}))(($v1_5)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_2) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_State_Trans_bindStateT'] = __NAMESPACE__ . '\\majControl_majMonad_majState_majTrans_bindmajStatemajT';

// Control_Monad_State_Trans_applyStateT
function majControl_majMonad_majState_majTrans_applymajStatemajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majState_majTrans_applymajStatemajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_1_0 = (object)["map" => function($f_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($__local_var_1_0, $f_2) {
  $__num = \func_num_args();
  $__res = function($s_4) use ($__local_var_1_0, $f_2, $v_3) {
  $__num = \func_num_args();
  $__res = ((($__local_var_1_0)->{'map'})(function($v1_5) use ($f_2) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_2)(($v1_5)->{'value0'}), ($v1_5)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_3)($s_4));
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
  $Bind1_2_2 = (($dictMonad_0)->{'Bind1'})(null);
  $Bind1_2_2 = (object)["bind" => function($v_3) use ($Bind1_2_2) {
  $__num = \func_num_args();
  $__res = function($f_4) use ($Bind1_2_2, $v_3) {
  $__num = \func_num_args();
  $__res = function($s_5) use ($Bind1_2_2, $f_4, $v_3) {
  $__num = \func_num_args();
  $__res = ((($Bind1_2_2)->{'bind'})(($v_3)($s_5)))(function($v1_6) use ($f_4) {
  $__num = \func_num_args();
  $__res = (($f_4)(($v1_6)->{'value0'}))(($v1_6)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_3) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_3_4 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_4) use ($Applicative0_3_4, $Bind1_2_2) {
  $__num = \func_num_args();
  $__res = function($a_5) use ($Applicative0_3_4, $Bind1_2_2, $f_4) {
  $__num = \func_num_args();
  $__res = ((($Bind1_2_2)->{'bind'})($f_4))(function($f_prime__6) use ($Applicative0_3_4, $Bind1_2_2, $a_5) {
  $__num = \func_num_args();
  $__res = ((($Bind1_2_2)->{'bind'})($a_5))(function($a_prime__7) use ($Applicative0_3_4, $f_prime__6) {
  $__num = \func_num_args();
  $__res = (($Applicative0_3_4)->{'pure'})(($f_prime__6)($a_prime__7));
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
}, "Functor0" => function($_dollar___unused_2) use ($functorStateT1_1_0) {
  $__num = \func_num_args();
  $__res = $functorStateT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_State_Trans_applyStateT'] = __NAMESPACE__ . '\\majControl_majMonad_majState_majTrans_applymajStatemajT';

// Control_Monad_State_Trans_applicativeStateT
function majControl_majMonad_majState_majTrans_applicativemajStatemajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majState_majTrans_applicativemajStatemajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $pure_1_0 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $__res = (object)["pure" => function($a_2) use ($pure_1_0) {
  $__num = \func_num_args();
  $__res = function($s_3) use ($a_2, $pure_1_0) {
  $__num = \func_num_args();
  $__res = ($pure_1_0)(new \Data\Tuple\Data_Tuple_Tuple($a_2, $s_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_2) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_3_1 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_3_1 = (object)["map" => function($f_4) use ($__local_var_3_1) {
  $__num = \func_num_args();
  $__res = function($v_5) use ($__local_var_3_1, $f_4) {
  $__num = \func_num_args();
  $__res = function($s_6) use ($__local_var_3_1, $f_4, $v_5) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_1)->{'map'})(function($v1_7) use ($f_4) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_4)(($v1_7)->{'value0'}), ($v1_7)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_5)($s_6));
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
  $Bind1_4_3 = (($dictMonad_0)->{'Bind1'})(null);
  $Bind1_4_3 = (object)["bind" => function($v_5) use ($Bind1_4_3) {
  $__num = \func_num_args();
  $__res = function($f_6) use ($Bind1_4_3, $v_5) {
  $__num = \func_num_args();
  $__res = function($s_7) use ($Bind1_4_3, $f_6, $v_5) {
  $__num = \func_num_args();
  $__res = ((($Bind1_4_3)->{'bind'})(($v_5)($s_7)))(function($v1_8) use ($f_6) {
  $__num = \func_num_args();
  $__res = (($f_6)(($v1_8)->{'value0'}))(($v1_8)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_5) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_5_5 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_6) use ($Applicative0_5_5, $Bind1_4_3) {
  $__num = \func_num_args();
  $__res = function($a_7) use ($Applicative0_5_5, $Bind1_4_3, $f_6) {
  $__num = \func_num_args();
  $__res = ((($Bind1_4_3)->{'bind'})($f_6))(function($f_prime__8) use ($Applicative0_5_5, $Bind1_4_3, $a_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_4_3)->{'bind'})($a_7))(function($a_prime__9) use ($Applicative0_5_5, $f_prime__8) {
  $__num = \func_num_args();
  $__res = (($Applicative0_5_5)->{'pure'})(($f_prime__8)($a_prime__9));
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
}, "Functor0" => function($_dollar___unused_4) use ($functorStateT1_3_1) {
  $__num = \func_num_args();
  $__res = $functorStateT1_3_1;
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_State_Trans_applicativeStateT'] = __NAMESPACE__ . '\\majControl_majMonad_majState_majTrans_applicativemajStatemajT';

// Control_Monad_State_Trans_semigroupStateT
function majControl_majMonad_majState_majTrans_semigroupmajStatemajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majState_majTrans_semigroupmajStatemajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_1_0 = (object)["map" => function($f_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($__local_var_1_0, $f_2) {
  $__num = \func_num_args();
  $__res = function($s_4) use ($__local_var_1_0, $f_2, $v_3) {
  $__num = \func_num_args();
  $__res = ((($__local_var_1_0)->{'map'})(function($v1_5) use ($f_2) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_2)(($v1_5)->{'value0'}), ($v1_5)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_3)($s_4));
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
  $Bind1_2_2 = (($dictMonad_0)->{'Bind1'})(null);
  $Bind1_2_2 = (object)["bind" => function($v_3) use ($Bind1_2_2) {
  $__num = \func_num_args();
  $__res = function($f_4) use ($Bind1_2_2, $v_3) {
  $__num = \func_num_args();
  $__res = function($s_5) use ($Bind1_2_2, $f_4, $v_3) {
  $__num = \func_num_args();
  $__res = ((($Bind1_2_2)->{'bind'})(($v_3)($s_5)))(function($v1_6) use ($f_4) {
  $__num = \func_num_args();
  $__res = (($f_4)(($v1_6)->{'value0'}))(($v1_6)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_3) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $pure_3_4 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $Applicative0_3_4 = (object)["pure" => function($a_4) use ($pure_3_4) {
  $__num = \func_num_args();
  $__res = function($s_5) use ($a_4, $pure_3_4) {
  $__num = \func_num_args();
  $__res = ($pure_3_4)(new \Data\Tuple\Data_Tuple_Tuple($a_4, $s_5));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_4) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_5_5 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_5_5 = (object)["map" => function($f_6) use ($__local_var_5_5) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($__local_var_5_5, $f_6) {
  $__num = \func_num_args();
  $__res = function($s_8) use ($__local_var_5_5, $f_6, $v_7) {
  $__num = \func_num_args();
  $__res = ((($__local_var_5_5)->{'map'})(function($v1_9) use ($f_6) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_6)(($v1_9)->{'value0'}), ($v1_9)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_7)($s_8));
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
  $Bind1_6_7 = (($dictMonad_0)->{'Bind1'})(null);
  $Bind1_6_7 = (object)["bind" => function($v_7) use ($Bind1_6_7) {
  $__num = \func_num_args();
  $__res = function($f_8) use ($Bind1_6_7, $v_7) {
  $__num = \func_num_args();
  $__res = function($s_9) use ($Bind1_6_7, $f_8, $v_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_7)->{'bind'})(($v_7)($s_9)))(function($v1_10) use ($f_8) {
  $__num = \func_num_args();
  $__res = (($f_8)(($v1_10)->{'value0'}))(($v1_10)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_7) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_7_9 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_8) use ($Applicative0_7_9, $Bind1_6_7) {
  $__num = \func_num_args();
  $__res = function($a_9) use ($Applicative0_7_9, $Bind1_6_7, $f_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_7)->{'bind'})($f_8))(function($f_prime__10) use ($Applicative0_7_9, $Bind1_6_7, $a_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_7)->{'bind'})($a_9))(function($a_prime__11) use ($Applicative0_7_9, $f_prime__10) {
  $__num = \func_num_args();
  $__res = (($Applicative0_7_9)->{'pure'})(($f_prime__10)($a_prime__11));
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
}, "Functor0" => function($_dollar___unused_6) use ($functorStateT1_5_5) {
  $__num = \func_num_args();
  $__res = $functorStateT1_5_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyStateT1_1_0 = (object)["apply" => function($f_4) use ($Applicative0_3_4, $Bind1_2_2) {
  $__num = \func_num_args();
  $__res = function($a_5) use ($Applicative0_3_4, $Bind1_2_2, $f_4) {
  $__num = \func_num_args();
  $__res = ((($Bind1_2_2)->{'bind'})($f_4))(function($f_prime__6) use ($Applicative0_3_4, $Bind1_2_2, $a_5) {
  $__num = \func_num_args();
  $__res = ((($Bind1_2_2)->{'bind'})($a_5))(function($a_prime__7) use ($Applicative0_3_4, $f_prime__6) {
  $__num = \func_num_args();
  $__res = (($Applicative0_3_4)->{'pure'})(($f_prime__6)($a_prime__7));
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
}, "Functor0" => function($_dollar___unused_2) use ($functorStateT1_1_0) {
  $__num = \func_num_args();
  $__res = $functorStateT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = function($dictSemigroup_2) use ($applyStateT1_1_0) {
  $__num = \func_num_args();
  $Functor0_3_12 = (($applyStateT1_1_0)->{'Functor0'})(null);
  $__local_var_4_13 = ($dictSemigroup_2)->{'append'};
  $__res = (object)["append" => function($a_5) use ($Functor0_3_12, $__local_var_4_13, $applyStateT1_1_0) {
  $__num = \func_num_args();
  $__res = function($b_6) use ($Functor0_3_12, $__local_var_4_13, $a_5, $applyStateT1_1_0) {
  $__num = \func_num_args();
  $__res = ((($applyStateT1_1_0)->{'apply'})(((($Functor0_3_12)->{'map'})($__local_var_4_13))($a_5)))($b_6);
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
$GLOBALS['Control_Monad_State_Trans_semigroupStateT'] = __NAMESPACE__ . '\\majControl_majMonad_majState_majTrans_semigroupmajStatemajT';

// Control_Monad_State_Trans_monadAskStateT
function majControl_majMonad_majState_majTrans_monadmajAskmajStatemajT($dictMonadAsk_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majState_majTrans_monadmajAskmajStatemajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictMonadAsk_0)->{'Monad0'})(null);
  $monadStateT1_1_0 = (object)["Applicative0" => function($_dollar___unused_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $pure_3_1 = ((($__local_var_1_0)->{'Applicative0'})(null))->{'pure'};
  $__res = (object)["pure" => function($a_4) use ($pure_3_1) {
  $__num = \func_num_args();
  $__res = function($s_5) use ($a_4, $pure_3_1) {
  $__num = \func_num_args();
  $__res = ($pure_3_1)(new \Data\Tuple\Data_Tuple_Tuple($a_4, $s_5));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_4) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__local_var_5_2 = (((((($__local_var_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_5_2 = (object)["map" => function($f_6) use ($__local_var_5_2) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($__local_var_5_2, $f_6) {
  $__num = \func_num_args();
  $__res = function($s_8) use ($__local_var_5_2, $f_6, $v_7) {
  $__num = \func_num_args();
  $__res = ((($__local_var_5_2)->{'map'})(function($v1_9) use ($f_6) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_6)(($v1_9)->{'value0'}), ($v1_9)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_7)($s_8));
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
  $Bind1_6_4 = (($__local_var_1_0)->{'Bind1'})(null);
  $Bind1_6_4 = (object)["bind" => function($v_7) use ($Bind1_6_4) {
  $__num = \func_num_args();
  $__res = function($f_8) use ($Bind1_6_4, $v_7) {
  $__num = \func_num_args();
  $__res = function($s_9) use ($Bind1_6_4, $f_8, $v_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_4)->{'bind'})(($v_7)($s_9)))(function($v1_10) use ($f_8) {
  $__num = \func_num_args();
  $__res = (($f_8)(($v1_10)->{'value0'}))(($v1_10)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_7) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__local_var_8_5 = (((((($__local_var_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_8_5 = (object)["map" => function($f_9) use ($__local_var_8_5) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($__local_var_8_5, $f_9) {
  $__num = \func_num_args();
  $__res = function($s_11) use ($__local_var_8_5, $f_9, $v_10) {
  $__num = \func_num_args();
  $__res = ((($__local_var_8_5)->{'map'})(function($v1_12) use ($f_9) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_9)(($v1_12)->{'value0'}), ($v1_12)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_10)($s_11));
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
  $Bind1_9_7 = (($__local_var_1_0)->{'Bind1'})(null);
  $Bind1_9_7 = (object)["bind" => function($v_10) use ($Bind1_9_7) {
  $__num = \func_num_args();
  $__res = function($f_11) use ($Bind1_9_7, $v_10) {
  $__num = \func_num_args();
  $__res = function($s_12) use ($Bind1_9_7, $f_11, $v_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_7)->{'bind'})(($v_10)($s_12)))(function($v1_13) use ($f_11) {
  $__num = \func_num_args();
  $__res = (($f_11)(($v1_13)->{'value0'}))(($v1_13)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_10) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($__local_var_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_10_9 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($__local_var_1_0);
  $__res = (object)["apply" => function($f_11) use ($Applicative0_10_9, $Bind1_9_7) {
  $__num = \func_num_args();
  $__res = function($a_12) use ($Applicative0_10_9, $Bind1_9_7, $f_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_7)->{'bind'})($f_11))(function($f_prime__13) use ($Applicative0_10_9, $Bind1_9_7, $a_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_7)->{'bind'})($a_12))(function($a_prime__14) use ($Applicative0_10_9, $f_prime__13) {
  $__num = \func_num_args();
  $__res = (($Applicative0_10_9)->{'pure'})(($f_prime__13)($a_prime__14));
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
}, "Functor0" => function($_dollar___unused_9) use ($functorStateT1_8_5) {
  $__num = \func_num_args();
  $__res = $functorStateT1_8_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_7_11 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($__local_var_1_0);
  $__res = (object)["apply" => function($f_8) use ($Applicative0_7_11, $Bind1_6_4) {
  $__num = \func_num_args();
  $__res = function($a_9) use ($Applicative0_7_11, $Bind1_6_4, $f_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_4)->{'bind'})($f_8))(function($f_prime__10) use ($Applicative0_7_11, $Bind1_6_4, $a_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_4)->{'bind'})($a_9))(function($a_prime__11) use ($Applicative0_7_11, $f_prime__10) {
  $__num = \func_num_args();
  $__res = (($Applicative0_7_11)->{'pure'})(($f_prime__10)($a_prime__11));
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
}, "Functor0" => function($_dollar___unused_6) use ($functorStateT1_5_2) {
  $__num = \func_num_args();
  $__res = $functorStateT1_5_2;
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
}, "Bind1" => function($_dollar___unused_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $Bind1_3_12 = (($__local_var_1_0)->{'Bind1'})(null);
  $__res = (object)["bind" => function($v_4) use ($Bind1_3_12) {
  $__num = \func_num_args();
  $__res = function($f_5) use ($Bind1_3_12, $v_4) {
  $__num = \func_num_args();
  $__res = function($s_6) use ($Bind1_3_12, $f_5, $v_4) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_12)->{'bind'})(($v_4)($s_6)))(function($v1_7) use ($f_5) {
  $__num = \func_num_args();
  $__res = (($f_5)(($v1_7)->{'value0'}))(($v1_7)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_4) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__local_var_5_13 = (((((($__local_var_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_5_13 = (object)["map" => function($f_6) use ($__local_var_5_13) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($__local_var_5_13, $f_6) {
  $__num = \func_num_args();
  $__res = function($s_8) use ($__local_var_5_13, $f_6, $v_7) {
  $__num = \func_num_args();
  $__res = ((($__local_var_5_13)->{'map'})(function($v1_9) use ($f_6) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_6)(($v1_9)->{'value0'}), ($v1_9)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_7)($s_8));
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
  $Bind1_6_15 = (($__local_var_1_0)->{'Bind1'})(null);
  $Bind1_6_15 = (object)["bind" => function($v_7) use ($Bind1_6_15) {
  $__num = \func_num_args();
  $__res = function($f_8) use ($Bind1_6_15, $v_7) {
  $__num = \func_num_args();
  $__res = function($s_9) use ($Bind1_6_15, $f_8, $v_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_15)->{'bind'})(($v_7)($s_9)))(function($v1_10) use ($f_8) {
  $__num = \func_num_args();
  $__res = (($f_8)(($v1_10)->{'value0'}))(($v1_10)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_7) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($__local_var_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $pure_7_17 = ((($__local_var_1_0)->{'Applicative0'})(null))->{'pure'};
  $Applicative0_7_17 = (object)["pure" => function($a_8) use ($pure_7_17) {
  $__num = \func_num_args();
  $__res = function($s_9) use ($a_8, $pure_7_17) {
  $__num = \func_num_args();
  $__res = ($pure_7_17)(new \Data\Tuple\Data_Tuple_Tuple($a_8, $s_9));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_8) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__local_var_9_18 = (((((($__local_var_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_9_18 = (object)["map" => function($f_10) use ($__local_var_9_18) {
  $__num = \func_num_args();
  $__res = function($v_11) use ($__local_var_9_18, $f_10) {
  $__num = \func_num_args();
  $__res = function($s_12) use ($__local_var_9_18, $f_10, $v_11) {
  $__num = \func_num_args();
  $__res = ((($__local_var_9_18)->{'map'})(function($v1_13) use ($f_10) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_10)(($v1_13)->{'value0'}), ($v1_13)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_11)($s_12));
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
  $Bind1_10_20 = (($__local_var_1_0)->{'Bind1'})(null);
  $Bind1_10_20 = (object)["bind" => function($v_11) use ($Bind1_10_20) {
  $__num = \func_num_args();
  $__res = function($f_12) use ($Bind1_10_20, $v_11) {
  $__num = \func_num_args();
  $__res = function($s_13) use ($Bind1_10_20, $f_12, $v_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_20)->{'bind'})(($v_11)($s_13)))(function($v1_14) use ($f_12) {
  $__num = \func_num_args();
  $__res = (($f_12)(($v1_14)->{'value0'}))(($v1_14)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_11) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($__local_var_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_11_22 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($__local_var_1_0);
  $__res = (object)["apply" => function($f_12) use ($Applicative0_11_22, $Bind1_10_20) {
  $__num = \func_num_args();
  $__res = function($a_13) use ($Applicative0_11_22, $Bind1_10_20, $f_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_20)->{'bind'})($f_12))(function($f_prime__14) use ($Applicative0_11_22, $Bind1_10_20, $a_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_20)->{'bind'})($a_13))(function($a_prime__15) use ($Applicative0_11_22, $f_prime__14) {
  $__num = \func_num_args();
  $__res = (($Applicative0_11_22)->{'pure'})(($f_prime__14)($a_prime__15));
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
}, "Functor0" => function($_dollar___unused_10) use ($functorStateT1_9_18) {
  $__num = \func_num_args();
  $__res = $functorStateT1_9_18;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_8) use ($Applicative0_7_17, $Bind1_6_15) {
  $__num = \func_num_args();
  $__res = function($a_9) use ($Applicative0_7_17, $Bind1_6_15, $f_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_15)->{'bind'})($f_8))(function($f_prime__10) use ($Applicative0_7_17, $Bind1_6_15, $a_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_15)->{'bind'})($a_9))(function($a_prime__11) use ($Applicative0_7_17, $f_prime__10) {
  $__num = \func_num_args();
  $__res = (($Applicative0_7_17)->{'pure'})(($f_prime__10)($a_prime__11));
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
}, "Functor0" => function($_dollar___unused_6) use ($functorStateT1_5_13) {
  $__num = \func_num_args();
  $__res = $functorStateT1_5_13;
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
}];
  $__local_var_2_25 = (($dictMonadAsk_0)->{'Monad0'})(null);
  $Bind1_3_26 = (($__local_var_2_25)->{'Bind1'})(null);
  $pure_4_27 = ((($__local_var_2_25)->{'Applicative0'})(null))->{'pure'};
  $__local_var_5_28 = ($dictMonadAsk_0)->{'ask'};
  $__res = (object)["ask" => function($s_6) use ($Bind1_3_26, $__local_var_5_28, $pure_4_27) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_26)->{'bind'})($__local_var_5_28))(function($x_7) use ($pure_4_27, $s_6) {
  $__num = \func_num_args();
  $__res = ($pure_4_27)(new \Data\Tuple\Data_Tuple_Tuple($x_7, $s_6));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad0" => function($_dollar___unused_2) use ($monadStateT1_1_0) {
  $__num = \func_num_args();
  $__res = $monadStateT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_State_Trans_monadAskStateT'] = __NAMESPACE__ . '\\majControl_majMonad_majState_majTrans_monadmajAskmajStatemajT';

// Control_Monad_State_Trans_monadReaderStateT
function majControl_majMonad_majState_majTrans_monadmajReadermajStatemajT($dictMonadReader_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majState_majTrans_monadmajReadermajStatemajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictMonadReader_0)->{'MonadAsk0'})(null);
  $__local_var_2_1 = (($__local_var_1_0)->{'Monad0'})(null);
  $monadStateT1_2_1 = (object)["Applicative0" => function($_dollar___unused_3) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $pure_4_2 = ((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'};
  $__res = (object)["pure" => function($a_5) use ($pure_4_2) {
  $__num = \func_num_args();
  $__res = function($s_6) use ($a_5, $pure_4_2) {
  $__num = \func_num_args();
  $__res = ($pure_4_2)(new \Data\Tuple\Data_Tuple_Tuple($a_5, $s_6));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_5) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_6_3 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_6_3 = (object)["map" => function($f_7) use ($__local_var_6_3) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($__local_var_6_3, $f_7) {
  $__num = \func_num_args();
  $__res = function($s_9) use ($__local_var_6_3, $f_7, $v_8) {
  $__num = \func_num_args();
  $__res = ((($__local_var_6_3)->{'map'})(function($v1_10) use ($f_7) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_7)(($v1_10)->{'value0'}), ($v1_10)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_8)($s_9));
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
  $Bind1_7_5 = (($__local_var_2_1)->{'Bind1'})(null);
  $Bind1_7_5 = (object)["bind" => function($v_8) use ($Bind1_7_5) {
  $__num = \func_num_args();
  $__res = function($f_9) use ($Bind1_7_5, $v_8) {
  $__num = \func_num_args();
  $__res = function($s_10) use ($Bind1_7_5, $f_9, $v_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_5)->{'bind'})(($v_8)($s_10)))(function($v1_11) use ($f_9) {
  $__num = \func_num_args();
  $__res = (($f_9)(($v1_11)->{'value0'}))(($v1_11)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_8) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_9_6 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_9_6 = (object)["map" => function($f_10) use ($__local_var_9_6) {
  $__num = \func_num_args();
  $__res = function($v_11) use ($__local_var_9_6, $f_10) {
  $__num = \func_num_args();
  $__res = function($s_12) use ($__local_var_9_6, $f_10, $v_11) {
  $__num = \func_num_args();
  $__res = ((($__local_var_9_6)->{'map'})(function($v1_13) use ($f_10) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_10)(($v1_13)->{'value0'}), ($v1_13)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_11)($s_12));
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
  $Bind1_10_8 = (($__local_var_2_1)->{'Bind1'})(null);
  $Bind1_10_8 = (object)["bind" => function($v_11) use ($Bind1_10_8) {
  $__num = \func_num_args();
  $__res = function($f_12) use ($Bind1_10_8, $v_11) {
  $__num = \func_num_args();
  $__res = function($s_13) use ($Bind1_10_8, $f_12, $v_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_8)->{'bind'})(($v_11)($s_13)))(function($v1_14) use ($f_12) {
  $__num = \func_num_args();
  $__res = (($f_12)(($v1_14)->{'value0'}))(($v1_14)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_11) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_12_9 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_12_9 = (object)["map" => function($f_13) use ($__local_var_12_9) {
  $__num = \func_num_args();
  $__res = function($v_14) use ($__local_var_12_9, $f_13) {
  $__num = \func_num_args();
  $__res = function($s_15) use ($__local_var_12_9, $f_13, $v_14) {
  $__num = \func_num_args();
  $__res = ((($__local_var_12_9)->{'map'})(function($v1_16) use ($f_13) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_13)(($v1_16)->{'value0'}), ($v1_16)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_14)($s_15));
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
  $Bind1_13_11 = (($__local_var_2_1)->{'Bind1'})(null);
  $Bind1_13_11 = (object)["bind" => function($v_14) use ($Bind1_13_11) {
  $__num = \func_num_args();
  $__res = function($f_15) use ($Bind1_13_11, $v_14) {
  $__num = \func_num_args();
  $__res = function($s_16) use ($Bind1_13_11, $f_15, $v_14) {
  $__num = \func_num_args();
  $__res = ((($Bind1_13_11)->{'bind'})(($v_14)($s_16)))(function($v1_17) use ($f_15) {
  $__num = \func_num_args();
  $__res = (($f_15)(($v1_17)->{'value0'}))(($v1_17)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_14) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($__local_var_2_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $pure_14_13 = ((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'};
  $Applicative0_14_13 = (object)["pure" => function($a_15) use ($pure_14_13) {
  $__num = \func_num_args();
  $__res = function($s_16) use ($a_15, $pure_14_13) {
  $__num = \func_num_args();
  $__res = ($pure_14_13)(new \Data\Tuple\Data_Tuple_Tuple($a_15, $s_16));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_15) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_16_14 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_16_14 = (object)["map" => function($f_17) use ($__local_var_16_14) {
  $__num = \func_num_args();
  $__res = function($v_18) use ($__local_var_16_14, $f_17) {
  $__num = \func_num_args();
  $__res = function($s_19) use ($__local_var_16_14, $f_17, $v_18) {
  $__num = \func_num_args();
  $__res = ((($__local_var_16_14)->{'map'})(function($v1_20) use ($f_17) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_17)(($v1_20)->{'value0'}), ($v1_20)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_18)($s_19));
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
  $Bind1_17_16 = (($__local_var_2_1)->{'Bind1'})(null);
  $Bind1_17_16 = (object)["bind" => function($v_18) use ($Bind1_17_16) {
  $__num = \func_num_args();
  $__res = function($f_19) use ($Bind1_17_16, $v_18) {
  $__num = \func_num_args();
  $__res = function($s_20) use ($Bind1_17_16, $f_19, $v_18) {
  $__num = \func_num_args();
  $__res = ((($Bind1_17_16)->{'bind'})(($v_18)($s_20)))(function($v1_21) use ($f_19) {
  $__num = \func_num_args();
  $__res = (($f_19)(($v1_21)->{'value0'}))(($v1_21)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_18) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($__local_var_2_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_18_18 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($__local_var_2_1);
  $__res = (object)["apply" => function($f_19) use ($Applicative0_18_18, $Bind1_17_16) {
  $__num = \func_num_args();
  $__res = function($a_20) use ($Applicative0_18_18, $Bind1_17_16, $f_19) {
  $__num = \func_num_args();
  $__res = ((($Bind1_17_16)->{'bind'})($f_19))(function($f_prime__21) use ($Applicative0_18_18, $Bind1_17_16, $a_20) {
  $__num = \func_num_args();
  $__res = ((($Bind1_17_16)->{'bind'})($a_20))(function($a_prime__22) use ($Applicative0_18_18, $f_prime__21) {
  $__num = \func_num_args();
  $__res = (($Applicative0_18_18)->{'pure'})(($f_prime__21)($a_prime__22));
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
}, "Functor0" => function($_dollar___unused_17) use ($functorStateT1_16_14) {
  $__num = \func_num_args();
  $__res = $functorStateT1_16_14;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_15) use ($Applicative0_14_13, $Bind1_13_11) {
  $__num = \func_num_args();
  $__res = function($a_16) use ($Applicative0_14_13, $Bind1_13_11, $f_15) {
  $__num = \func_num_args();
  $__res = ((($Bind1_13_11)->{'bind'})($f_15))(function($f_prime__17) use ($Applicative0_14_13, $Bind1_13_11, $a_16) {
  $__num = \func_num_args();
  $__res = ((($Bind1_13_11)->{'bind'})($a_16))(function($a_prime__18) use ($Applicative0_14_13, $f_prime__17) {
  $__num = \func_num_args();
  $__res = (($Applicative0_14_13)->{'pure'})(($f_prime__17)($a_prime__18));
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
}, "Functor0" => function($_dollar___unused_13) use ($functorStateT1_12_9) {
  $__num = \func_num_args();
  $__res = $functorStateT1_12_9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $pure_11_21 = ((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'};
  $Applicative0_11_21 = (object)["pure" => function($a_12) use ($pure_11_21) {
  $__num = \func_num_args();
  $__res = function($s_13) use ($a_12, $pure_11_21) {
  $__num = \func_num_args();
  $__res = ($pure_11_21)(new \Data\Tuple\Data_Tuple_Tuple($a_12, $s_13));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_12) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_13_22 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_13_22 = (object)["map" => function($f_14) use ($__local_var_13_22) {
  $__num = \func_num_args();
  $__res = function($v_15) use ($__local_var_13_22, $f_14) {
  $__num = \func_num_args();
  $__res = function($s_16) use ($__local_var_13_22, $f_14, $v_15) {
  $__num = \func_num_args();
  $__res = ((($__local_var_13_22)->{'map'})(function($v1_17) use ($f_14) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_14)(($v1_17)->{'value0'}), ($v1_17)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_15)($s_16));
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
  $Bind1_14_24 = (($__local_var_2_1)->{'Bind1'})(null);
  $Bind1_14_24 = (object)["bind" => function($v_15) use ($Bind1_14_24) {
  $__num = \func_num_args();
  $__res = function($f_16) use ($Bind1_14_24, $v_15) {
  $__num = \func_num_args();
  $__res = function($s_17) use ($Bind1_14_24, $f_16, $v_15) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_24)->{'bind'})(($v_15)($s_17)))(function($v1_18) use ($f_16) {
  $__num = \func_num_args();
  $__res = (($f_16)(($v1_18)->{'value0'}))(($v1_18)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_15) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_16_25 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_16_25 = (object)["map" => function($f_17) use ($__local_var_16_25) {
  $__num = \func_num_args();
  $__res = function($v_18) use ($__local_var_16_25, $f_17) {
  $__num = \func_num_args();
  $__res = function($s_19) use ($__local_var_16_25, $f_17, $v_18) {
  $__num = \func_num_args();
  $__res = ((($__local_var_16_25)->{'map'})(function($v1_20) use ($f_17) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_17)(($v1_20)->{'value0'}), ($v1_20)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_18)($s_19));
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
  $Bind1_17_27 = (($__local_var_2_1)->{'Bind1'})(null);
  $Bind1_17_27 = (object)["bind" => function($v_18) use ($Bind1_17_27) {
  $__num = \func_num_args();
  $__res = function($f_19) use ($Bind1_17_27, $v_18) {
  $__num = \func_num_args();
  $__res = function($s_20) use ($Bind1_17_27, $f_19, $v_18) {
  $__num = \func_num_args();
  $__res = ((($Bind1_17_27)->{'bind'})(($v_18)($s_20)))(function($v1_21) use ($f_19) {
  $__num = \func_num_args();
  $__res = (($f_19)(($v1_21)->{'value0'}))(($v1_21)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_18) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($__local_var_2_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_18_29 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($__local_var_2_1);
  $__res = (object)["apply" => function($f_19) use ($Applicative0_18_29, $Bind1_17_27) {
  $__num = \func_num_args();
  $__res = function($a_20) use ($Applicative0_18_29, $Bind1_17_27, $f_19) {
  $__num = \func_num_args();
  $__res = ((($Bind1_17_27)->{'bind'})($f_19))(function($f_prime__21) use ($Applicative0_18_29, $Bind1_17_27, $a_20) {
  $__num = \func_num_args();
  $__res = ((($Bind1_17_27)->{'bind'})($a_20))(function($a_prime__22) use ($Applicative0_18_29, $f_prime__21) {
  $__num = \func_num_args();
  $__res = (($Applicative0_18_29)->{'pure'})(($f_prime__21)($a_prime__22));
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
}, "Functor0" => function($_dollar___unused_17) use ($functorStateT1_16_25) {
  $__num = \func_num_args();
  $__res = $functorStateT1_16_25;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_15_31 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($__local_var_2_1);
  $__res = (object)["apply" => function($f_16) use ($Applicative0_15_31, $Bind1_14_24) {
  $__num = \func_num_args();
  $__res = function($a_17) use ($Applicative0_15_31, $Bind1_14_24, $f_16) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_24)->{'bind'})($f_16))(function($f_prime__18) use ($Applicative0_15_31, $Bind1_14_24, $a_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_24)->{'bind'})($a_17))(function($a_prime__19) use ($Applicative0_15_31, $f_prime__18) {
  $__num = \func_num_args();
  $__res = (($Applicative0_15_31)->{'pure'})(($f_prime__18)($a_prime__19));
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
}, "Functor0" => function($_dollar___unused_14) use ($functorStateT1_13_22) {
  $__num = \func_num_args();
  $__res = $functorStateT1_13_22;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_12) use ($Applicative0_11_21, $Bind1_10_8) {
  $__num = \func_num_args();
  $__res = function($a_13) use ($Applicative0_11_21, $Bind1_10_8, $f_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_8)->{'bind'})($f_12))(function($f_prime__14) use ($Applicative0_11_21, $Bind1_10_8, $a_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_8)->{'bind'})($a_13))(function($a_prime__15) use ($Applicative0_11_21, $f_prime__14) {
  $__num = \func_num_args();
  $__res = (($Applicative0_11_21)->{'pure'})(($f_prime__14)($a_prime__15));
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
}, "Functor0" => function($_dollar___unused_10) use ($functorStateT1_9_6) {
  $__num = \func_num_args();
  $__res = $functorStateT1_9_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $pure_8_34 = ((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'};
  $Applicative0_8_34 = (object)["pure" => function($a_9) use ($pure_8_34) {
  $__num = \func_num_args();
  $__res = function($s_10) use ($a_9, $pure_8_34) {
  $__num = \func_num_args();
  $__res = ($pure_8_34)(new \Data\Tuple\Data_Tuple_Tuple($a_9, $s_10));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_9) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_10_35 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_10_35 = (object)["map" => function($f_11) use ($__local_var_10_35) {
  $__num = \func_num_args();
  $__res = function($v_12) use ($__local_var_10_35, $f_11) {
  $__num = \func_num_args();
  $__res = function($s_13) use ($__local_var_10_35, $f_11, $v_12) {
  $__num = \func_num_args();
  $__res = ((($__local_var_10_35)->{'map'})(function($v1_14) use ($f_11) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_11)(($v1_14)->{'value0'}), ($v1_14)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_12)($s_13));
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
  $Bind1_11_37 = (($__local_var_2_1)->{'Bind1'})(null);
  $Bind1_11_37 = (object)["bind" => function($v_12) use ($Bind1_11_37) {
  $__num = \func_num_args();
  $__res = function($f_13) use ($Bind1_11_37, $v_12) {
  $__num = \func_num_args();
  $__res = function($s_14) use ($Bind1_11_37, $f_13, $v_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_37)->{'bind'})(($v_12)($s_14)))(function($v1_15) use ($f_13) {
  $__num = \func_num_args();
  $__res = (($f_13)(($v1_15)->{'value0'}))(($v1_15)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_12) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_13_38 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_13_38 = (object)["map" => function($f_14) use ($__local_var_13_38) {
  $__num = \func_num_args();
  $__res = function($v_15) use ($__local_var_13_38, $f_14) {
  $__num = \func_num_args();
  $__res = function($s_16) use ($__local_var_13_38, $f_14, $v_15) {
  $__num = \func_num_args();
  $__res = ((($__local_var_13_38)->{'map'})(function($v1_17) use ($f_14) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_14)(($v1_17)->{'value0'}), ($v1_17)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_15)($s_16));
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
  $Bind1_14_40 = (($__local_var_2_1)->{'Bind1'})(null);
  $Bind1_14_40 = (object)["bind" => function($v_15) use ($Bind1_14_40) {
  $__num = \func_num_args();
  $__res = function($f_16) use ($Bind1_14_40, $v_15) {
  $__num = \func_num_args();
  $__res = function($s_17) use ($Bind1_14_40, $f_16, $v_15) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_40)->{'bind'})(($v_15)($s_17)))(function($v1_18) use ($f_16) {
  $__num = \func_num_args();
  $__res = (($f_16)(($v1_18)->{'value0'}))(($v1_18)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_15) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($__local_var_2_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_15_42 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($__local_var_2_1);
  $__res = (object)["apply" => function($f_16) use ($Applicative0_15_42, $Bind1_14_40) {
  $__num = \func_num_args();
  $__res = function($a_17) use ($Applicative0_15_42, $Bind1_14_40, $f_16) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_40)->{'bind'})($f_16))(function($f_prime__18) use ($Applicative0_15_42, $Bind1_14_40, $a_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_40)->{'bind'})($a_17))(function($a_prime__19) use ($Applicative0_15_42, $f_prime__18) {
  $__num = \func_num_args();
  $__res = (($Applicative0_15_42)->{'pure'})(($f_prime__18)($a_prime__19));
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
}, "Functor0" => function($_dollar___unused_14) use ($functorStateT1_13_38) {
  $__num = \func_num_args();
  $__res = $functorStateT1_13_38;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_12_44 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($__local_var_2_1);
  $__res = (object)["apply" => function($f_13) use ($Applicative0_12_44, $Bind1_11_37) {
  $__num = \func_num_args();
  $__res = function($a_14) use ($Applicative0_12_44, $Bind1_11_37, $f_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_37)->{'bind'})($f_13))(function($f_prime__15) use ($Applicative0_12_44, $Bind1_11_37, $a_14) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_37)->{'bind'})($a_14))(function($a_prime__16) use ($Applicative0_12_44, $f_prime__15) {
  $__num = \func_num_args();
  $__res = (($Applicative0_12_44)->{'pure'})(($f_prime__15)($a_prime__16));
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
}, "Functor0" => function($_dollar___unused_11) use ($functorStateT1_10_35) {
  $__num = \func_num_args();
  $__res = $functorStateT1_10_35;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_9) use ($Applicative0_8_34, $Bind1_7_5) {
  $__num = \func_num_args();
  $__res = function($a_10) use ($Applicative0_8_34, $Bind1_7_5, $f_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_5)->{'bind'})($f_9))(function($f_prime__11) use ($Applicative0_8_34, $Bind1_7_5, $a_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_5)->{'bind'})($a_10))(function($a_prime__12) use ($Applicative0_8_34, $f_prime__11) {
  $__num = \func_num_args();
  $__res = (($Applicative0_8_34)->{'pure'})(($f_prime__11)($a_prime__12));
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
}, "Functor0" => function($_dollar___unused_7) use ($functorStateT1_6_3) {
  $__num = \func_num_args();
  $__res = $functorStateT1_6_3;
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
}, "Bind1" => function($_dollar___unused_3) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $Bind1_4_46 = (($__local_var_2_1)->{'Bind1'})(null);
  $__res = (object)["bind" => function($v_5) use ($Bind1_4_46) {
  $__num = \func_num_args();
  $__res = function($f_6) use ($Bind1_4_46, $v_5) {
  $__num = \func_num_args();
  $__res = function($s_7) use ($Bind1_4_46, $f_6, $v_5) {
  $__num = \func_num_args();
  $__res = ((($Bind1_4_46)->{'bind'})(($v_5)($s_7)))(function($v1_8) use ($f_6) {
  $__num = \func_num_args();
  $__res = (($f_6)(($v1_8)->{'value0'}))(($v1_8)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_5) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_6_47 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_6_47 = (object)["map" => function($f_7) use ($__local_var_6_47) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($__local_var_6_47, $f_7) {
  $__num = \func_num_args();
  $__res = function($s_9) use ($__local_var_6_47, $f_7, $v_8) {
  $__num = \func_num_args();
  $__res = ((($__local_var_6_47)->{'map'})(function($v1_10) use ($f_7) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_7)(($v1_10)->{'value0'}), ($v1_10)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_8)($s_9));
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
  $Bind1_7_49 = (($__local_var_2_1)->{'Bind1'})(null);
  $Bind1_7_49 = (object)["bind" => function($v_8) use ($Bind1_7_49) {
  $__num = \func_num_args();
  $__res = function($f_9) use ($Bind1_7_49, $v_8) {
  $__num = \func_num_args();
  $__res = function($s_10) use ($Bind1_7_49, $f_9, $v_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_49)->{'bind'})(($v_8)($s_10)))(function($v1_11) use ($f_9) {
  $__num = \func_num_args();
  $__res = (($f_9)(($v1_11)->{'value0'}))(($v1_11)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_8) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_9_50 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_9_50 = (object)["map" => function($f_10) use ($__local_var_9_50) {
  $__num = \func_num_args();
  $__res = function($v_11) use ($__local_var_9_50, $f_10) {
  $__num = \func_num_args();
  $__res = function($s_12) use ($__local_var_9_50, $f_10, $v_11) {
  $__num = \func_num_args();
  $__res = ((($__local_var_9_50)->{'map'})(function($v1_13) use ($f_10) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_10)(($v1_13)->{'value0'}), ($v1_13)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_11)($s_12));
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
  $Bind1_10_52 = (($__local_var_2_1)->{'Bind1'})(null);
  $Bind1_10_52 = (object)["bind" => function($v_11) use ($Bind1_10_52) {
  $__num = \func_num_args();
  $__res = function($f_12) use ($Bind1_10_52, $v_11) {
  $__num = \func_num_args();
  $__res = function($s_13) use ($Bind1_10_52, $f_12, $v_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_52)->{'bind'})(($v_11)($s_13)))(function($v1_14) use ($f_12) {
  $__num = \func_num_args();
  $__res = (($f_12)(($v1_14)->{'value0'}))(($v1_14)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_11) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($__local_var_2_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $pure_11_54 = ((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'};
  $Applicative0_11_54 = (object)["pure" => function($a_12) use ($pure_11_54) {
  $__num = \func_num_args();
  $__res = function($s_13) use ($a_12, $pure_11_54) {
  $__num = \func_num_args();
  $__res = ($pure_11_54)(new \Data\Tuple\Data_Tuple_Tuple($a_12, $s_13));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_12) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_13_55 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_13_55 = (object)["map" => function($f_14) use ($__local_var_13_55) {
  $__num = \func_num_args();
  $__res = function($v_15) use ($__local_var_13_55, $f_14) {
  $__num = \func_num_args();
  $__res = function($s_16) use ($__local_var_13_55, $f_14, $v_15) {
  $__num = \func_num_args();
  $__res = ((($__local_var_13_55)->{'map'})(function($v1_17) use ($f_14) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_14)(($v1_17)->{'value0'}), ($v1_17)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_15)($s_16));
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
  $Bind1_14_57 = (($__local_var_2_1)->{'Bind1'})(null);
  $Bind1_14_57 = (object)["bind" => function($v_15) use ($Bind1_14_57) {
  $__num = \func_num_args();
  $__res = function($f_16) use ($Bind1_14_57, $v_15) {
  $__num = \func_num_args();
  $__res = function($s_17) use ($Bind1_14_57, $f_16, $v_15) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_57)->{'bind'})(($v_15)($s_17)))(function($v1_18) use ($f_16) {
  $__num = \func_num_args();
  $__res = (($f_16)(($v1_18)->{'value0'}))(($v1_18)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_15) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($__local_var_2_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_15_59 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($__local_var_2_1);
  $__res = (object)["apply" => function($f_16) use ($Applicative0_15_59, $Bind1_14_57) {
  $__num = \func_num_args();
  $__res = function($a_17) use ($Applicative0_15_59, $Bind1_14_57, $f_16) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_57)->{'bind'})($f_16))(function($f_prime__18) use ($Applicative0_15_59, $Bind1_14_57, $a_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_57)->{'bind'})($a_17))(function($a_prime__19) use ($Applicative0_15_59, $f_prime__18) {
  $__num = \func_num_args();
  $__res = (($Applicative0_15_59)->{'pure'})(($f_prime__18)($a_prime__19));
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
}, "Functor0" => function($_dollar___unused_14) use ($functorStateT1_13_55) {
  $__num = \func_num_args();
  $__res = $functorStateT1_13_55;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_12) use ($Applicative0_11_54, $Bind1_10_52) {
  $__num = \func_num_args();
  $__res = function($a_13) use ($Applicative0_11_54, $Bind1_10_52, $f_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_52)->{'bind'})($f_12))(function($f_prime__14) use ($Applicative0_11_54, $Bind1_10_52, $a_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_52)->{'bind'})($a_13))(function($a_prime__15) use ($Applicative0_11_54, $f_prime__14) {
  $__num = \func_num_args();
  $__res = (($Applicative0_11_54)->{'pure'})(($f_prime__14)($a_prime__15));
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
}, "Functor0" => function($_dollar___unused_10) use ($functorStateT1_9_50) {
  $__num = \func_num_args();
  $__res = $functorStateT1_9_50;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $pure_8_62 = ((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'};
  $Applicative0_8_62 = (object)["pure" => function($a_9) use ($pure_8_62) {
  $__num = \func_num_args();
  $__res = function($s_10) use ($a_9, $pure_8_62) {
  $__num = \func_num_args();
  $__res = ($pure_8_62)(new \Data\Tuple\Data_Tuple_Tuple($a_9, $s_10));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_9) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_10_63 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_10_63 = (object)["map" => function($f_11) use ($__local_var_10_63) {
  $__num = \func_num_args();
  $__res = function($v_12) use ($__local_var_10_63, $f_11) {
  $__num = \func_num_args();
  $__res = function($s_13) use ($__local_var_10_63, $f_11, $v_12) {
  $__num = \func_num_args();
  $__res = ((($__local_var_10_63)->{'map'})(function($v1_14) use ($f_11) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_11)(($v1_14)->{'value0'}), ($v1_14)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_12)($s_13));
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
  $Bind1_11_65 = (($__local_var_2_1)->{'Bind1'})(null);
  $Bind1_11_65 = (object)["bind" => function($v_12) use ($Bind1_11_65) {
  $__num = \func_num_args();
  $__res = function($f_13) use ($Bind1_11_65, $v_12) {
  $__num = \func_num_args();
  $__res = function($s_14) use ($Bind1_11_65, $f_13, $v_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_65)->{'bind'})(($v_12)($s_14)))(function($v1_15) use ($f_13) {
  $__num = \func_num_args();
  $__res = (($f_13)(($v1_15)->{'value0'}))(($v1_15)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_12) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_13_66 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_13_66 = (object)["map" => function($f_14) use ($__local_var_13_66) {
  $__num = \func_num_args();
  $__res = function($v_15) use ($__local_var_13_66, $f_14) {
  $__num = \func_num_args();
  $__res = function($s_16) use ($__local_var_13_66, $f_14, $v_15) {
  $__num = \func_num_args();
  $__res = ((($__local_var_13_66)->{'map'})(function($v1_17) use ($f_14) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_14)(($v1_17)->{'value0'}), ($v1_17)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_15)($s_16));
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
  $Bind1_14_68 = (($__local_var_2_1)->{'Bind1'})(null);
  $Bind1_14_68 = (object)["bind" => function($v_15) use ($Bind1_14_68) {
  $__num = \func_num_args();
  $__res = function($f_16) use ($Bind1_14_68, $v_15) {
  $__num = \func_num_args();
  $__res = function($s_17) use ($Bind1_14_68, $f_16, $v_15) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_68)->{'bind'})(($v_15)($s_17)))(function($v1_18) use ($f_16) {
  $__num = \func_num_args();
  $__res = (($f_16)(($v1_18)->{'value0'}))(($v1_18)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_15) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($__local_var_2_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $pure_15_70 = ((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'};
  $Applicative0_15_70 = (object)["pure" => function($a_16) use ($pure_15_70) {
  $__num = \func_num_args();
  $__res = function($s_17) use ($a_16, $pure_15_70) {
  $__num = \func_num_args();
  $__res = ($pure_15_70)(new \Data\Tuple\Data_Tuple_Tuple($a_16, $s_17));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_16) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_17_71 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_17_71 = (object)["map" => function($f_18) use ($__local_var_17_71) {
  $__num = \func_num_args();
  $__res = function($v_19) use ($__local_var_17_71, $f_18) {
  $__num = \func_num_args();
  $__res = function($s_20) use ($__local_var_17_71, $f_18, $v_19) {
  $__num = \func_num_args();
  $__res = ((($__local_var_17_71)->{'map'})(function($v1_21) use ($f_18) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_18)(($v1_21)->{'value0'}), ($v1_21)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_19)($s_20));
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
  $Bind1_18_73 = (($__local_var_2_1)->{'Bind1'})(null);
  $Bind1_18_73 = (object)["bind" => function($v_19) use ($Bind1_18_73) {
  $__num = \func_num_args();
  $__res = function($f_20) use ($Bind1_18_73, $v_19) {
  $__num = \func_num_args();
  $__res = function($s_21) use ($Bind1_18_73, $f_20, $v_19) {
  $__num = \func_num_args();
  $__res = ((($Bind1_18_73)->{'bind'})(($v_19)($s_21)))(function($v1_22) use ($f_20) {
  $__num = \func_num_args();
  $__res = (($f_20)(($v1_22)->{'value0'}))(($v1_22)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_19) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($__local_var_2_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_19_75 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($__local_var_2_1);
  $__res = (object)["apply" => function($f_20) use ($Applicative0_19_75, $Bind1_18_73) {
  $__num = \func_num_args();
  $__res = function($a_21) use ($Applicative0_19_75, $Bind1_18_73, $f_20) {
  $__num = \func_num_args();
  $__res = ((($Bind1_18_73)->{'bind'})($f_20))(function($f_prime__22) use ($Applicative0_19_75, $Bind1_18_73, $a_21) {
  $__num = \func_num_args();
  $__res = ((($Bind1_18_73)->{'bind'})($a_21))(function($a_prime__23) use ($Applicative0_19_75, $f_prime__22) {
  $__num = \func_num_args();
  $__res = (($Applicative0_19_75)->{'pure'})(($f_prime__22)($a_prime__23));
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
}, "Functor0" => function($_dollar___unused_18) use ($functorStateT1_17_71) {
  $__num = \func_num_args();
  $__res = $functorStateT1_17_71;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_16) use ($Applicative0_15_70, $Bind1_14_68) {
  $__num = \func_num_args();
  $__res = function($a_17) use ($Applicative0_15_70, $Bind1_14_68, $f_16) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_68)->{'bind'})($f_16))(function($f_prime__18) use ($Applicative0_15_70, $Bind1_14_68, $a_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_68)->{'bind'})($a_17))(function($a_prime__19) use ($Applicative0_15_70, $f_prime__18) {
  $__num = \func_num_args();
  $__res = (($Applicative0_15_70)->{'pure'})(($f_prime__18)($a_prime__19));
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
}, "Functor0" => function($_dollar___unused_14) use ($functorStateT1_13_66) {
  $__num = \func_num_args();
  $__res = $functorStateT1_13_66;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $pure_12_78 = ((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'};
  $Applicative0_12_78 = (object)["pure" => function($a_13) use ($pure_12_78) {
  $__num = \func_num_args();
  $__res = function($s_14) use ($a_13, $pure_12_78) {
  $__num = \func_num_args();
  $__res = ($pure_12_78)(new \Data\Tuple\Data_Tuple_Tuple($a_13, $s_14));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_13) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_14_79 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_14_79 = (object)["map" => function($f_15) use ($__local_var_14_79) {
  $__num = \func_num_args();
  $__res = function($v_16) use ($__local_var_14_79, $f_15) {
  $__num = \func_num_args();
  $__res = function($s_17) use ($__local_var_14_79, $f_15, $v_16) {
  $__num = \func_num_args();
  $__res = ((($__local_var_14_79)->{'map'})(function($v1_18) use ($f_15) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_15)(($v1_18)->{'value0'}), ($v1_18)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_16)($s_17));
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
  $Bind1_15_81 = (($__local_var_2_1)->{'Bind1'})(null);
  $Bind1_15_81 = (object)["bind" => function($v_16) use ($Bind1_15_81) {
  $__num = \func_num_args();
  $__res = function($f_17) use ($Bind1_15_81, $v_16) {
  $__num = \func_num_args();
  $__res = function($s_18) use ($Bind1_15_81, $f_17, $v_16) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_81)->{'bind'})(($v_16)($s_18)))(function($v1_19) use ($f_17) {
  $__num = \func_num_args();
  $__res = (($f_17)(($v1_19)->{'value0'}))(($v1_19)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_16) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_17_82 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_17_82 = (object)["map" => function($f_18) use ($__local_var_17_82) {
  $__num = \func_num_args();
  $__res = function($v_19) use ($__local_var_17_82, $f_18) {
  $__num = \func_num_args();
  $__res = function($s_20) use ($__local_var_17_82, $f_18, $v_19) {
  $__num = \func_num_args();
  $__res = ((($__local_var_17_82)->{'map'})(function($v1_21) use ($f_18) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_18)(($v1_21)->{'value0'}), ($v1_21)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_19)($s_20));
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
  $Bind1_18_84 = (($__local_var_2_1)->{'Bind1'})(null);
  $Bind1_18_84 = (object)["bind" => function($v_19) use ($Bind1_18_84) {
  $__num = \func_num_args();
  $__res = function($f_20) use ($Bind1_18_84, $v_19) {
  $__num = \func_num_args();
  $__res = function($s_21) use ($Bind1_18_84, $f_20, $v_19) {
  $__num = \func_num_args();
  $__res = ((($Bind1_18_84)->{'bind'})(($v_19)($s_21)))(function($v1_22) use ($f_20) {
  $__num = \func_num_args();
  $__res = (($f_20)(($v1_22)->{'value0'}))(($v1_22)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_19) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($__local_var_2_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_19_86 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($__local_var_2_1);
  $__res = (object)["apply" => function($f_20) use ($Applicative0_19_86, $Bind1_18_84) {
  $__num = \func_num_args();
  $__res = function($a_21) use ($Applicative0_19_86, $Bind1_18_84, $f_20) {
  $__num = \func_num_args();
  $__res = ((($Bind1_18_84)->{'bind'})($f_20))(function($f_prime__22) use ($Applicative0_19_86, $Bind1_18_84, $a_21) {
  $__num = \func_num_args();
  $__res = ((($Bind1_18_84)->{'bind'})($a_21))(function($a_prime__23) use ($Applicative0_19_86, $f_prime__22) {
  $__num = \func_num_args();
  $__res = (($Applicative0_19_86)->{'pure'})(($f_prime__22)($a_prime__23));
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
}, "Functor0" => function($_dollar___unused_18) use ($functorStateT1_17_82) {
  $__num = \func_num_args();
  $__res = $functorStateT1_17_82;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_16_88 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($__local_var_2_1);
  $__res = (object)["apply" => function($f_17) use ($Applicative0_16_88, $Bind1_15_81) {
  $__num = \func_num_args();
  $__res = function($a_18) use ($Applicative0_16_88, $Bind1_15_81, $f_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_81)->{'bind'})($f_17))(function($f_prime__19) use ($Applicative0_16_88, $Bind1_15_81, $a_18) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_81)->{'bind'})($a_18))(function($a_prime__20) use ($Applicative0_16_88, $f_prime__19) {
  $__num = \func_num_args();
  $__res = (($Applicative0_16_88)->{'pure'})(($f_prime__19)($a_prime__20));
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
}, "Functor0" => function($_dollar___unused_15) use ($functorStateT1_14_79) {
  $__num = \func_num_args();
  $__res = $functorStateT1_14_79;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_13) use ($Applicative0_12_78, $Bind1_11_65) {
  $__num = \func_num_args();
  $__res = function($a_14) use ($Applicative0_12_78, $Bind1_11_65, $f_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_65)->{'bind'})($f_13))(function($f_prime__15) use ($Applicative0_12_78, $Bind1_11_65, $a_14) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_65)->{'bind'})($a_14))(function($a_prime__16) use ($Applicative0_12_78, $f_prime__15) {
  $__num = \func_num_args();
  $__res = (($Applicative0_12_78)->{'pure'})(($f_prime__15)($a_prime__16));
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
}, "Functor0" => function($_dollar___unused_11) use ($functorStateT1_10_63) {
  $__num = \func_num_args();
  $__res = $functorStateT1_10_63;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_9) use ($Applicative0_8_62, $Bind1_7_49) {
  $__num = \func_num_args();
  $__res = function($a_10) use ($Applicative0_8_62, $Bind1_7_49, $f_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_49)->{'bind'})($f_9))(function($f_prime__11) use ($Applicative0_8_62, $Bind1_7_49, $a_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_49)->{'bind'})($a_10))(function($a_prime__12) use ($Applicative0_8_62, $f_prime__11) {
  $__num = \func_num_args();
  $__res = (($Applicative0_8_62)->{'pure'})(($f_prime__11)($a_prime__12));
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
}, "Functor0" => function($_dollar___unused_7) use ($functorStateT1_6_47) {
  $__num = \func_num_args();
  $__res = $functorStateT1_6_47;
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
}];
  $__local_var_3_92 = (($__local_var_1_0)->{'Monad0'})(null);
  $Bind1_4_93 = (($__local_var_3_92)->{'Bind1'})(null);
  $pure_5_94 = ((($__local_var_3_92)->{'Applicative0'})(null))->{'pure'};
  $__local_var_6_95 = ($__local_var_1_0)->{'ask'};
  $monadAskStateT1_1_0 = (object)["ask" => function($s_7) use ($Bind1_4_93, $__local_var_6_95, $pure_5_94) {
  $__num = \func_num_args();
  $__res = ((($Bind1_4_93)->{'bind'})($__local_var_6_95))(function($x_8) use ($pure_5_94, $s_7) {
  $__num = \func_num_args();
  $__res = ($pure_5_94)(new \Data\Tuple\Data_Tuple_Tuple($x_8, $s_7));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad0" => function($_dollar___unused_3) use ($monadStateT1_2_1) {
  $__num = \func_num_args();
  $__res = $monadStateT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["local" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_State_Trans_mapStateT']))(($dictMonadReader_0)->{'local'}), "MonadAsk0" => function($_dollar___unused_2) use ($monadAskStateT1_1_0) {
  $__num = \func_num_args();
  $__res = $monadAskStateT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_State_Trans_monadReaderStateT'] = __NAMESPACE__ . '\\majControl_majMonad_majState_majTrans_monadmajReadermajStatemajT';

// Control_Monad_State_Trans_monadContStateT
function majControl_majMonad_majState_majTrans_monadmajContmajStatemajT($dictMonadCont_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majState_majTrans_monadmajContmajStatemajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictMonadCont_0)->{'Monad0'})(null);
  $monadStateT1_1_0 = (object)["Applicative0" => function($_dollar___unused_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $pure_3_1 = ((($__local_var_1_0)->{'Applicative0'})(null))->{'pure'};
  $__res = (object)["pure" => function($a_4) use ($pure_3_1) {
  $__num = \func_num_args();
  $__res = function($s_5) use ($a_4, $pure_3_1) {
  $__num = \func_num_args();
  $__res = ($pure_3_1)(new \Data\Tuple\Data_Tuple_Tuple($a_4, $s_5));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_4) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__local_var_5_2 = (((((($__local_var_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_5_2 = (object)["map" => function($f_6) use ($__local_var_5_2) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($__local_var_5_2, $f_6) {
  $__num = \func_num_args();
  $__res = function($s_8) use ($__local_var_5_2, $f_6, $v_7) {
  $__num = \func_num_args();
  $__res = ((($__local_var_5_2)->{'map'})(function($v1_9) use ($f_6) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_6)(($v1_9)->{'value0'}), ($v1_9)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_7)($s_8));
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
  $Bind1_6_4 = (($__local_var_1_0)->{'Bind1'})(null);
  $Bind1_6_4 = (object)["bind" => function($v_7) use ($Bind1_6_4) {
  $__num = \func_num_args();
  $__res = function($f_8) use ($Bind1_6_4, $v_7) {
  $__num = \func_num_args();
  $__res = function($s_9) use ($Bind1_6_4, $f_8, $v_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_4)->{'bind'})(($v_7)($s_9)))(function($v1_10) use ($f_8) {
  $__num = \func_num_args();
  $__res = (($f_8)(($v1_10)->{'value0'}))(($v1_10)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_7) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__local_var_8_5 = (((((($__local_var_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_8_5 = (object)["map" => function($f_9) use ($__local_var_8_5) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($__local_var_8_5, $f_9) {
  $__num = \func_num_args();
  $__res = function($s_11) use ($__local_var_8_5, $f_9, $v_10) {
  $__num = \func_num_args();
  $__res = ((($__local_var_8_5)->{'map'})(function($v1_12) use ($f_9) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_9)(($v1_12)->{'value0'}), ($v1_12)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_10)($s_11));
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
  $Bind1_9_7 = (($__local_var_1_0)->{'Bind1'})(null);
  $Bind1_9_7 = (object)["bind" => function($v_10) use ($Bind1_9_7) {
  $__num = \func_num_args();
  $__res = function($f_11) use ($Bind1_9_7, $v_10) {
  $__num = \func_num_args();
  $__res = function($s_12) use ($Bind1_9_7, $f_11, $v_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_7)->{'bind'})(($v_10)($s_12)))(function($v1_13) use ($f_11) {
  $__num = \func_num_args();
  $__res = (($f_11)(($v1_13)->{'value0'}))(($v1_13)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_10) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($__local_var_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_10_9 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($__local_var_1_0);
  $__res = (object)["apply" => function($f_11) use ($Applicative0_10_9, $Bind1_9_7) {
  $__num = \func_num_args();
  $__res = function($a_12) use ($Applicative0_10_9, $Bind1_9_7, $f_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_7)->{'bind'})($f_11))(function($f_prime__13) use ($Applicative0_10_9, $Bind1_9_7, $a_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_7)->{'bind'})($a_12))(function($a_prime__14) use ($Applicative0_10_9, $f_prime__13) {
  $__num = \func_num_args();
  $__res = (($Applicative0_10_9)->{'pure'})(($f_prime__13)($a_prime__14));
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
}, "Functor0" => function($_dollar___unused_9) use ($functorStateT1_8_5) {
  $__num = \func_num_args();
  $__res = $functorStateT1_8_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_7_11 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($__local_var_1_0);
  $__res = (object)["apply" => function($f_8) use ($Applicative0_7_11, $Bind1_6_4) {
  $__num = \func_num_args();
  $__res = function($a_9) use ($Applicative0_7_11, $Bind1_6_4, $f_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_4)->{'bind'})($f_8))(function($f_prime__10) use ($Applicative0_7_11, $Bind1_6_4, $a_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_4)->{'bind'})($a_9))(function($a_prime__11) use ($Applicative0_7_11, $f_prime__10) {
  $__num = \func_num_args();
  $__res = (($Applicative0_7_11)->{'pure'})(($f_prime__10)($a_prime__11));
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
}, "Functor0" => function($_dollar___unused_6) use ($functorStateT1_5_2) {
  $__num = \func_num_args();
  $__res = $functorStateT1_5_2;
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
}, "Bind1" => function($_dollar___unused_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $Bind1_3_12 = (($__local_var_1_0)->{'Bind1'})(null);
  $__res = (object)["bind" => function($v_4) use ($Bind1_3_12) {
  $__num = \func_num_args();
  $__res = function($f_5) use ($Bind1_3_12, $v_4) {
  $__num = \func_num_args();
  $__res = function($s_6) use ($Bind1_3_12, $f_5, $v_4) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_12)->{'bind'})(($v_4)($s_6)))(function($v1_7) use ($f_5) {
  $__num = \func_num_args();
  $__res = (($f_5)(($v1_7)->{'value0'}))(($v1_7)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_4) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__local_var_5_13 = (((((($__local_var_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_5_13 = (object)["map" => function($f_6) use ($__local_var_5_13) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($__local_var_5_13, $f_6) {
  $__num = \func_num_args();
  $__res = function($s_8) use ($__local_var_5_13, $f_6, $v_7) {
  $__num = \func_num_args();
  $__res = ((($__local_var_5_13)->{'map'})(function($v1_9) use ($f_6) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_6)(($v1_9)->{'value0'}), ($v1_9)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_7)($s_8));
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
  $Bind1_6_15 = (($__local_var_1_0)->{'Bind1'})(null);
  $Bind1_6_15 = (object)["bind" => function($v_7) use ($Bind1_6_15) {
  $__num = \func_num_args();
  $__res = function($f_8) use ($Bind1_6_15, $v_7) {
  $__num = \func_num_args();
  $__res = function($s_9) use ($Bind1_6_15, $f_8, $v_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_15)->{'bind'})(($v_7)($s_9)))(function($v1_10) use ($f_8) {
  $__num = \func_num_args();
  $__res = (($f_8)(($v1_10)->{'value0'}))(($v1_10)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_7) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($__local_var_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $pure_7_17 = ((($__local_var_1_0)->{'Applicative0'})(null))->{'pure'};
  $Applicative0_7_17 = (object)["pure" => function($a_8) use ($pure_7_17) {
  $__num = \func_num_args();
  $__res = function($s_9) use ($a_8, $pure_7_17) {
  $__num = \func_num_args();
  $__res = ($pure_7_17)(new \Data\Tuple\Data_Tuple_Tuple($a_8, $s_9));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_8) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__local_var_9_18 = (((((($__local_var_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_9_18 = (object)["map" => function($f_10) use ($__local_var_9_18) {
  $__num = \func_num_args();
  $__res = function($v_11) use ($__local_var_9_18, $f_10) {
  $__num = \func_num_args();
  $__res = function($s_12) use ($__local_var_9_18, $f_10, $v_11) {
  $__num = \func_num_args();
  $__res = ((($__local_var_9_18)->{'map'})(function($v1_13) use ($f_10) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_10)(($v1_13)->{'value0'}), ($v1_13)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_11)($s_12));
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
  $Bind1_10_20 = (($__local_var_1_0)->{'Bind1'})(null);
  $Bind1_10_20 = (object)["bind" => function($v_11) use ($Bind1_10_20) {
  $__num = \func_num_args();
  $__res = function($f_12) use ($Bind1_10_20, $v_11) {
  $__num = \func_num_args();
  $__res = function($s_13) use ($Bind1_10_20, $f_12, $v_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_20)->{'bind'})(($v_11)($s_13)))(function($v1_14) use ($f_12) {
  $__num = \func_num_args();
  $__res = (($f_12)(($v1_14)->{'value0'}))(($v1_14)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_11) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($__local_var_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_11_22 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($__local_var_1_0);
  $__res = (object)["apply" => function($f_12) use ($Applicative0_11_22, $Bind1_10_20) {
  $__num = \func_num_args();
  $__res = function($a_13) use ($Applicative0_11_22, $Bind1_10_20, $f_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_20)->{'bind'})($f_12))(function($f_prime__14) use ($Applicative0_11_22, $Bind1_10_20, $a_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_20)->{'bind'})($a_13))(function($a_prime__15) use ($Applicative0_11_22, $f_prime__14) {
  $__num = \func_num_args();
  $__res = (($Applicative0_11_22)->{'pure'})(($f_prime__14)($a_prime__15));
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
}, "Functor0" => function($_dollar___unused_10) use ($functorStateT1_9_18) {
  $__num = \func_num_args();
  $__res = $functorStateT1_9_18;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_8) use ($Applicative0_7_17, $Bind1_6_15) {
  $__num = \func_num_args();
  $__res = function($a_9) use ($Applicative0_7_17, $Bind1_6_15, $f_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_15)->{'bind'})($f_8))(function($f_prime__10) use ($Applicative0_7_17, $Bind1_6_15, $a_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_15)->{'bind'})($a_9))(function($a_prime__11) use ($Applicative0_7_17, $f_prime__10) {
  $__num = \func_num_args();
  $__res = (($Applicative0_7_17)->{'pure'})(($f_prime__10)($a_prime__11));
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
}, "Functor0" => function($_dollar___unused_6) use ($functorStateT1_5_13) {
  $__num = \func_num_args();
  $__res = $functorStateT1_5_13;
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
}];
  $__res = (object)["callCC" => function($f_2) use ($dictMonadCont_0) {
  $__num = \func_num_args();
  $__res = function($s_3) use ($dictMonadCont_0, $f_2) {
  $__num = \func_num_args();
  $__res = (($dictMonadCont_0)->{'callCC'})(function($c_4) use ($f_2, $s_3) {
  $__num = \func_num_args();
  $__res = (($f_2)(function($a_5) use ($c_4) {
  $__num = \func_num_args();
  $__res = function($s_prime__6) use ($a_5, $c_4) {
  $__num = \func_num_args();
  $__res = ($c_4)(new \Data\Tuple\Data_Tuple_Tuple($a_5, $s_prime__6));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($s_3);
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
}, "Monad0" => function($_dollar___unused_2) use ($monadStateT1_1_0) {
  $__num = \func_num_args();
  $__res = $monadStateT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_State_Trans_monadContStateT'] = __NAMESPACE__ . '\\majControl_majMonad_majState_majTrans_monadmajContmajStatemajT';

// Control_Monad_State_Trans_monadEffectState
function majControl_majMonad_majState_majTrans_monadmajEffectmajState($dictMonadEffect_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majState_majTrans_monadmajEffectmajState';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Monad0_1_0 = (($dictMonadEffect_0)->{'Monad0'})(null);
  $monadStateT1_2_1 = (object)["Applicative0" => function($_dollar___unused_2) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $pure_3_1 = ((($Monad0_1_0)->{'Applicative0'})(null))->{'pure'};
  $__res = (object)["pure" => function($a_4) use ($pure_3_1) {
  $__num = \func_num_args();
  $__res = function($s_5) use ($a_4, $pure_3_1) {
  $__num = \func_num_args();
  $__res = ($pure_3_1)(new \Data\Tuple\Data_Tuple_Tuple($a_4, $s_5));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_4) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__local_var_5_2 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_5_2 = (object)["map" => function($f_6) use ($__local_var_5_2) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($__local_var_5_2, $f_6) {
  $__num = \func_num_args();
  $__res = function($s_8) use ($__local_var_5_2, $f_6, $v_7) {
  $__num = \func_num_args();
  $__res = ((($__local_var_5_2)->{'map'})(function($v1_9) use ($f_6) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_6)(($v1_9)->{'value0'}), ($v1_9)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_7)($s_8));
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
  $Bind1_6_4 = (($Monad0_1_0)->{'Bind1'})(null);
  $Bind1_6_4 = (object)["bind" => function($v_7) use ($Bind1_6_4) {
  $__num = \func_num_args();
  $__res = function($f_8) use ($Bind1_6_4, $v_7) {
  $__num = \func_num_args();
  $__res = function($s_9) use ($Bind1_6_4, $f_8, $v_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_4)->{'bind'})(($v_7)($s_9)))(function($v1_10) use ($f_8) {
  $__num = \func_num_args();
  $__res = (($f_8)(($v1_10)->{'value0'}))(($v1_10)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_7) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__local_var_8_5 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_8_5 = (object)["map" => function($f_9) use ($__local_var_8_5) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($__local_var_8_5, $f_9) {
  $__num = \func_num_args();
  $__res = function($s_11) use ($__local_var_8_5, $f_9, $v_10) {
  $__num = \func_num_args();
  $__res = ((($__local_var_8_5)->{'map'})(function($v1_12) use ($f_9) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_9)(($v1_12)->{'value0'}), ($v1_12)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_10)($s_11));
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
  $Bind1_9_7 = (($Monad0_1_0)->{'Bind1'})(null);
  $Bind1_9_7 = (object)["bind" => function($v_10) use ($Bind1_9_7) {
  $__num = \func_num_args();
  $__res = function($f_11) use ($Bind1_9_7, $v_10) {
  $__num = \func_num_args();
  $__res = function($s_12) use ($Bind1_9_7, $f_11, $v_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_7)->{'bind'})(($v_10)($s_12)))(function($v1_13) use ($f_11) {
  $__num = \func_num_args();
  $__res = (($f_11)(($v1_13)->{'value0'}))(($v1_13)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_10) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_10_9 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($Monad0_1_0);
  $__res = (object)["apply" => function($f_11) use ($Applicative0_10_9, $Bind1_9_7) {
  $__num = \func_num_args();
  $__res = function($a_12) use ($Applicative0_10_9, $Bind1_9_7, $f_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_7)->{'bind'})($f_11))(function($f_prime__13) use ($Applicative0_10_9, $Bind1_9_7, $a_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_7)->{'bind'})($a_12))(function($a_prime__14) use ($Applicative0_10_9, $f_prime__13) {
  $__num = \func_num_args();
  $__res = (($Applicative0_10_9)->{'pure'})(($f_prime__13)($a_prime__14));
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
}, "Functor0" => function($_dollar___unused_9) use ($functorStateT1_8_5) {
  $__num = \func_num_args();
  $__res = $functorStateT1_8_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_7_11 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($Monad0_1_0);
  $__res = (object)["apply" => function($f_8) use ($Applicative0_7_11, $Bind1_6_4) {
  $__num = \func_num_args();
  $__res = function($a_9) use ($Applicative0_7_11, $Bind1_6_4, $f_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_4)->{'bind'})($f_8))(function($f_prime__10) use ($Applicative0_7_11, $Bind1_6_4, $a_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_4)->{'bind'})($a_9))(function($a_prime__11) use ($Applicative0_7_11, $f_prime__10) {
  $__num = \func_num_args();
  $__res = (($Applicative0_7_11)->{'pure'})(($f_prime__10)($a_prime__11));
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
}, "Functor0" => function($_dollar___unused_6) use ($functorStateT1_5_2) {
  $__num = \func_num_args();
  $__res = $functorStateT1_5_2;
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
}, "Bind1" => function($_dollar___unused_2) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $Bind1_3_12 = (($Monad0_1_0)->{'Bind1'})(null);
  $__res = (object)["bind" => function($v_4) use ($Bind1_3_12) {
  $__num = \func_num_args();
  $__res = function($f_5) use ($Bind1_3_12, $v_4) {
  $__num = \func_num_args();
  $__res = function($s_6) use ($Bind1_3_12, $f_5, $v_4) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_12)->{'bind'})(($v_4)($s_6)))(function($v1_7) use ($f_5) {
  $__num = \func_num_args();
  $__res = (($f_5)(($v1_7)->{'value0'}))(($v1_7)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_4) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__local_var_5_13 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_5_13 = (object)["map" => function($f_6) use ($__local_var_5_13) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($__local_var_5_13, $f_6) {
  $__num = \func_num_args();
  $__res = function($s_8) use ($__local_var_5_13, $f_6, $v_7) {
  $__num = \func_num_args();
  $__res = ((($__local_var_5_13)->{'map'})(function($v1_9) use ($f_6) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_6)(($v1_9)->{'value0'}), ($v1_9)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_7)($s_8));
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
  $Bind1_6_15 = (($Monad0_1_0)->{'Bind1'})(null);
  $Bind1_6_15 = (object)["bind" => function($v_7) use ($Bind1_6_15) {
  $__num = \func_num_args();
  $__res = function($f_8) use ($Bind1_6_15, $v_7) {
  $__num = \func_num_args();
  $__res = function($s_9) use ($Bind1_6_15, $f_8, $v_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_15)->{'bind'})(($v_7)($s_9)))(function($v1_10) use ($f_8) {
  $__num = \func_num_args();
  $__res = (($f_8)(($v1_10)->{'value0'}))(($v1_10)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_7) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $pure_7_17 = ((($Monad0_1_0)->{'Applicative0'})(null))->{'pure'};
  $Applicative0_7_17 = (object)["pure" => function($a_8) use ($pure_7_17) {
  $__num = \func_num_args();
  $__res = function($s_9) use ($a_8, $pure_7_17) {
  $__num = \func_num_args();
  $__res = ($pure_7_17)(new \Data\Tuple\Data_Tuple_Tuple($a_8, $s_9));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_8) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__local_var_9_18 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_9_18 = (object)["map" => function($f_10) use ($__local_var_9_18) {
  $__num = \func_num_args();
  $__res = function($v_11) use ($__local_var_9_18, $f_10) {
  $__num = \func_num_args();
  $__res = function($s_12) use ($__local_var_9_18, $f_10, $v_11) {
  $__num = \func_num_args();
  $__res = ((($__local_var_9_18)->{'map'})(function($v1_13) use ($f_10) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_10)(($v1_13)->{'value0'}), ($v1_13)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_11)($s_12));
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
  $Bind1_10_20 = (($Monad0_1_0)->{'Bind1'})(null);
  $Bind1_10_20 = (object)["bind" => function($v_11) use ($Bind1_10_20) {
  $__num = \func_num_args();
  $__res = function($f_12) use ($Bind1_10_20, $v_11) {
  $__num = \func_num_args();
  $__res = function($s_13) use ($Bind1_10_20, $f_12, $v_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_20)->{'bind'})(($v_11)($s_13)))(function($v1_14) use ($f_12) {
  $__num = \func_num_args();
  $__res = (($f_12)(($v1_14)->{'value0'}))(($v1_14)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_11) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_11_22 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($Monad0_1_0);
  $__res = (object)["apply" => function($f_12) use ($Applicative0_11_22, $Bind1_10_20) {
  $__num = \func_num_args();
  $__res = function($a_13) use ($Applicative0_11_22, $Bind1_10_20, $f_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_20)->{'bind'})($f_12))(function($f_prime__14) use ($Applicative0_11_22, $Bind1_10_20, $a_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_20)->{'bind'})($a_13))(function($a_prime__15) use ($Applicative0_11_22, $f_prime__14) {
  $__num = \func_num_args();
  $__res = (($Applicative0_11_22)->{'pure'})(($f_prime__14)($a_prime__15));
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
}, "Functor0" => function($_dollar___unused_10) use ($functorStateT1_9_18) {
  $__num = \func_num_args();
  $__res = $functorStateT1_9_18;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_8) use ($Applicative0_7_17, $Bind1_6_15) {
  $__num = \func_num_args();
  $__res = function($a_9) use ($Applicative0_7_17, $Bind1_6_15, $f_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_15)->{'bind'})($f_8))(function($f_prime__10) use ($Applicative0_7_17, $Bind1_6_15, $a_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_15)->{'bind'})($a_9))(function($a_prime__11) use ($Applicative0_7_17, $f_prime__10) {
  $__num = \func_num_args();
  $__res = (($Applicative0_7_17)->{'pure'})(($f_prime__10)($a_prime__11));
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
}, "Functor0" => function($_dollar___unused_6) use ($functorStateT1_5_13) {
  $__num = \func_num_args();
  $__res = $functorStateT1_5_13;
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
}];
  $Bind1_3_25 = (($Monad0_1_0)->{'Bind1'})(null);
  $pure_4_26 = ((($Monad0_1_0)->{'Applicative0'})(null))->{'pure'};
  $__res = (object)["liftEffect" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($m_5) use ($Bind1_3_25, $pure_4_26) {
  $__num = \func_num_args();
  $__res = function($s_6) use ($Bind1_3_25, $m_5, $pure_4_26) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_25)->{'bind'})($m_5))(function($x_7) use ($pure_4_26, $s_6) {
  $__num = \func_num_args();
  $__res = ($pure_4_26)(new \Data\Tuple\Data_Tuple_Tuple($x_7, $s_6));
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
}))(($dictMonadEffect_0)->{'liftEffect'}), "Monad0" => function($_dollar___unused_3) use ($monadStateT1_2_1) {
  $__num = \func_num_args();
  $__res = $monadStateT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_State_Trans_monadEffectState'] = __NAMESPACE__ . '\\majControl_majMonad_majState_majTrans_monadmajEffectmajState';

// Control_Monad_State_Trans_monadRecStateT
function majControl_majMonad_majState_majTrans_monadmajRecmajStatemajT($dictMonadRec_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majState_majTrans_monadmajRecmajStatemajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Monad0_1_0 = (($dictMonadRec_0)->{'Monad0'})(null);
  $Bind1_2_1 = (($Monad0_1_0)->{'Bind1'})(null);
  $Applicative0_3_2 = (($Monad0_1_0)->{'Applicative0'})(null);
  $monadStateT1_4_3 = (object)["Applicative0" => function($_dollar___unused_4) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $pure_5_3 = ((($Monad0_1_0)->{'Applicative0'})(null))->{'pure'};
  $__res = (object)["pure" => function($a_6) use ($pure_5_3) {
  $__num = \func_num_args();
  $__res = function($s_7) use ($a_6, $pure_5_3) {
  $__num = \func_num_args();
  $__res = ($pure_5_3)(new \Data\Tuple\Data_Tuple_Tuple($a_6, $s_7));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_6) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__local_var_7_4 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_7_4 = (object)["map" => function($f_8) use ($__local_var_7_4) {
  $__num = \func_num_args();
  $__res = function($v_9) use ($__local_var_7_4, $f_8) {
  $__num = \func_num_args();
  $__res = function($s_10) use ($__local_var_7_4, $f_8, $v_9) {
  $__num = \func_num_args();
  $__res = ((($__local_var_7_4)->{'map'})(function($v1_11) use ($f_8) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_8)(($v1_11)->{'value0'}), ($v1_11)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_9)($s_10));
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
  $Bind1_8_6 = (($Monad0_1_0)->{'Bind1'})(null);
  $Bind1_8_6 = (object)["bind" => function($v_9) use ($Bind1_8_6) {
  $__num = \func_num_args();
  $__res = function($f_10) use ($Bind1_8_6, $v_9) {
  $__num = \func_num_args();
  $__res = function($s_11) use ($Bind1_8_6, $f_10, $v_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_8_6)->{'bind'})(($v_9)($s_11)))(function($v1_12) use ($f_10) {
  $__num = \func_num_args();
  $__res = (($f_10)(($v1_12)->{'value0'}))(($v1_12)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_9) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__local_var_10_7 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_10_7 = (object)["map" => function($f_11) use ($__local_var_10_7) {
  $__num = \func_num_args();
  $__res = function($v_12) use ($__local_var_10_7, $f_11) {
  $__num = \func_num_args();
  $__res = function($s_13) use ($__local_var_10_7, $f_11, $v_12) {
  $__num = \func_num_args();
  $__res = ((($__local_var_10_7)->{'map'})(function($v1_14) use ($f_11) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_11)(($v1_14)->{'value0'}), ($v1_14)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_12)($s_13));
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
  $Bind1_11_9 = (($Monad0_1_0)->{'Bind1'})(null);
  $Bind1_11_9 = (object)["bind" => function($v_12) use ($Bind1_11_9) {
  $__num = \func_num_args();
  $__res = function($f_13) use ($Bind1_11_9, $v_12) {
  $__num = \func_num_args();
  $__res = function($s_14) use ($Bind1_11_9, $f_13, $v_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_9)->{'bind'})(($v_12)($s_14)))(function($v1_15) use ($f_13) {
  $__num = \func_num_args();
  $__res = (($f_13)(($v1_15)->{'value0'}))(($v1_15)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_12) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_12_11 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($Monad0_1_0);
  $__res = (object)["apply" => function($f_13) use ($Applicative0_12_11, $Bind1_11_9) {
  $__num = \func_num_args();
  $__res = function($a_14) use ($Applicative0_12_11, $Bind1_11_9, $f_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_9)->{'bind'})($f_13))(function($f_prime__15) use ($Applicative0_12_11, $Bind1_11_9, $a_14) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_9)->{'bind'})($a_14))(function($a_prime__16) use ($Applicative0_12_11, $f_prime__15) {
  $__num = \func_num_args();
  $__res = (($Applicative0_12_11)->{'pure'})(($f_prime__15)($a_prime__16));
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
}, "Functor0" => function($_dollar___unused_11) use ($functorStateT1_10_7) {
  $__num = \func_num_args();
  $__res = $functorStateT1_10_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_9_13 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($Monad0_1_0);
  $__res = (object)["apply" => function($f_10) use ($Applicative0_9_13, $Bind1_8_6) {
  $__num = \func_num_args();
  $__res = function($a_11) use ($Applicative0_9_13, $Bind1_8_6, $f_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_8_6)->{'bind'})($f_10))(function($f_prime__12) use ($Applicative0_9_13, $Bind1_8_6, $a_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_8_6)->{'bind'})($a_11))(function($a_prime__13) use ($Applicative0_9_13, $f_prime__12) {
  $__num = \func_num_args();
  $__res = (($Applicative0_9_13)->{'pure'})(($f_prime__12)($a_prime__13));
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
}, "Functor0" => function($_dollar___unused_8) use ($functorStateT1_7_4) {
  $__num = \func_num_args();
  $__res = $functorStateT1_7_4;
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
}, "Bind1" => function($_dollar___unused_4) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $Bind1_5_14 = (($Monad0_1_0)->{'Bind1'})(null);
  $__res = (object)["bind" => function($v_6) use ($Bind1_5_14) {
  $__num = \func_num_args();
  $__res = function($f_7) use ($Bind1_5_14, $v_6) {
  $__num = \func_num_args();
  $__res = function($s_8) use ($Bind1_5_14, $f_7, $v_6) {
  $__num = \func_num_args();
  $__res = ((($Bind1_5_14)->{'bind'})(($v_6)($s_8)))(function($v1_9) use ($f_7) {
  $__num = \func_num_args();
  $__res = (($f_7)(($v1_9)->{'value0'}))(($v1_9)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_6) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__local_var_7_15 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_7_15 = (object)["map" => function($f_8) use ($__local_var_7_15) {
  $__num = \func_num_args();
  $__res = function($v_9) use ($__local_var_7_15, $f_8) {
  $__num = \func_num_args();
  $__res = function($s_10) use ($__local_var_7_15, $f_8, $v_9) {
  $__num = \func_num_args();
  $__res = ((($__local_var_7_15)->{'map'})(function($v1_11) use ($f_8) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_8)(($v1_11)->{'value0'}), ($v1_11)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_9)($s_10));
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
  $Bind1_8_17 = (($Monad0_1_0)->{'Bind1'})(null);
  $Bind1_8_17 = (object)["bind" => function($v_9) use ($Bind1_8_17) {
  $__num = \func_num_args();
  $__res = function($f_10) use ($Bind1_8_17, $v_9) {
  $__num = \func_num_args();
  $__res = function($s_11) use ($Bind1_8_17, $f_10, $v_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_8_17)->{'bind'})(($v_9)($s_11)))(function($v1_12) use ($f_10) {
  $__num = \func_num_args();
  $__res = (($f_10)(($v1_12)->{'value0'}))(($v1_12)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_9) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $pure_9_19 = ((($Monad0_1_0)->{'Applicative0'})(null))->{'pure'};
  $Applicative0_9_19 = (object)["pure" => function($a_10) use ($pure_9_19) {
  $__num = \func_num_args();
  $__res = function($s_11) use ($a_10, $pure_9_19) {
  $__num = \func_num_args();
  $__res = ($pure_9_19)(new \Data\Tuple\Data_Tuple_Tuple($a_10, $s_11));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_10) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__local_var_11_20 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_11_20 = (object)["map" => function($f_12) use ($__local_var_11_20) {
  $__num = \func_num_args();
  $__res = function($v_13) use ($__local_var_11_20, $f_12) {
  $__num = \func_num_args();
  $__res = function($s_14) use ($__local_var_11_20, $f_12, $v_13) {
  $__num = \func_num_args();
  $__res = ((($__local_var_11_20)->{'map'})(function($v1_15) use ($f_12) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_12)(($v1_15)->{'value0'}), ($v1_15)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_13)($s_14));
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
  $Bind1_12_22 = (($Monad0_1_0)->{'Bind1'})(null);
  $Bind1_12_22 = (object)["bind" => function($v_13) use ($Bind1_12_22) {
  $__num = \func_num_args();
  $__res = function($f_14) use ($Bind1_12_22, $v_13) {
  $__num = \func_num_args();
  $__res = function($s_15) use ($Bind1_12_22, $f_14, $v_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_12_22)->{'bind'})(($v_13)($s_15)))(function($v1_16) use ($f_14) {
  $__num = \func_num_args();
  $__res = (($f_14)(($v1_16)->{'value0'}))(($v1_16)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_13) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_13_24 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($Monad0_1_0);
  $__res = (object)["apply" => function($f_14) use ($Applicative0_13_24, $Bind1_12_22) {
  $__num = \func_num_args();
  $__res = function($a_15) use ($Applicative0_13_24, $Bind1_12_22, $f_14) {
  $__num = \func_num_args();
  $__res = ((($Bind1_12_22)->{'bind'})($f_14))(function($f_prime__16) use ($Applicative0_13_24, $Bind1_12_22, $a_15) {
  $__num = \func_num_args();
  $__res = ((($Bind1_12_22)->{'bind'})($a_15))(function($a_prime__17) use ($Applicative0_13_24, $f_prime__16) {
  $__num = \func_num_args();
  $__res = (($Applicative0_13_24)->{'pure'})(($f_prime__16)($a_prime__17));
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
}, "Functor0" => function($_dollar___unused_12) use ($functorStateT1_11_20) {
  $__num = \func_num_args();
  $__res = $functorStateT1_11_20;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_10) use ($Applicative0_9_19, $Bind1_8_17) {
  $__num = \func_num_args();
  $__res = function($a_11) use ($Applicative0_9_19, $Bind1_8_17, $f_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_8_17)->{'bind'})($f_10))(function($f_prime__12) use ($Applicative0_9_19, $Bind1_8_17, $a_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_8_17)->{'bind'})($a_11))(function($a_prime__13) use ($Applicative0_9_19, $f_prime__12) {
  $__num = \func_num_args();
  $__res = (($Applicative0_9_19)->{'pure'})(($f_prime__12)($a_prime__13));
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
}, "Functor0" => function($_dollar___unused_8) use ($functorStateT1_7_15) {
  $__num = \func_num_args();
  $__res = $functorStateT1_7_15;
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
}];
  $__res = (object)["tailRecM" => function($f_5) use ($Applicative0_3_2, $Bind1_2_1, $dictMonadRec_0) {
  $__num = \func_num_args();
  $__res = function($a_6) use ($Applicative0_3_2, $Bind1_2_1, $dictMonadRec_0, $f_5) {
  $__num = \func_num_args();
  $__res = function($s_7) use ($Applicative0_3_2, $Bind1_2_1, $a_6, $dictMonadRec_0, $f_5) {
  $__num = \func_num_args();
  $__res = ((($dictMonadRec_0)->{'tailRecM'})(function($v_8) use ($Applicative0_3_2, $Bind1_2_1, $f_5) {
  $__num = \func_num_args();
  $__res = ((($Bind1_2_1)->{'bind'})((($f_5)(($v_8)->{'value0'}))(($v_8)->{'value1'})))(function($v2_9) use ($Applicative0_3_2) {
  $__num = \func_num_args();
  $__t27 = null;;
  if (($v2_9)->{'value0'} instanceof \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Loop) {
$__t27 = new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Loop(new \Data\Tuple\Data_Tuple_Tuple((($v2_9)->{'value0'})->{'value0'}, ($v2_9)->{'value1'}));
goto end_branch_27;;
};
  if (($v2_9)->{'value0'} instanceof \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Done) {
$__t27 = new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Done(new \Data\Tuple\Data_Tuple_Tuple((($v2_9)->{'value0'})->{'value0'}, ($v2_9)->{'value1'}));
goto end_branch_27;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t27 = null;
  end_branch_27:;
  $__res = (($Applicative0_3_2)->{'pure'})($__t27);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(new \Data\Tuple\Data_Tuple_Tuple($a_6, $s_7));
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
}, "Monad0" => function($_dollar___unused_5) use ($monadStateT1_4_3) {
  $__num = \func_num_args();
  $__res = $monadStateT1_4_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_State_Trans_monadRecStateT'] = __NAMESPACE__ . '\\majControl_majMonad_majState_majTrans_monadmajRecmajStatemajT';

// Control_Monad_State_Trans_monadStateStateT
function majControl_majMonad_majState_majTrans_monadmajStatemajStatemajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majState_majTrans_monadmajStatemajStatemajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $pure_1_0 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $monadStateT1_2_1 = (object)["Applicative0" => function($_dollar___unused_2) use ($dictMonad_0) {
  $__num = \func_num_args();
  $pure_3_1 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $__res = (object)["pure" => function($a_4) use ($pure_3_1) {
  $__num = \func_num_args();
  $__res = function($s_5) use ($a_4, $pure_3_1) {
  $__num = \func_num_args();
  $__res = ($pure_3_1)(new \Data\Tuple\Data_Tuple_Tuple($a_4, $s_5));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_4) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_5_2 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_5_2 = (object)["map" => function($f_6) use ($__local_var_5_2) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($__local_var_5_2, $f_6) {
  $__num = \func_num_args();
  $__res = function($s_8) use ($__local_var_5_2, $f_6, $v_7) {
  $__num = \func_num_args();
  $__res = ((($__local_var_5_2)->{'map'})(function($v1_9) use ($f_6) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_6)(($v1_9)->{'value0'}), ($v1_9)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_7)($s_8));
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
  $Bind1_6_4 = (($dictMonad_0)->{'Bind1'})(null);
  $Bind1_6_4 = (object)["bind" => function($v_7) use ($Bind1_6_4) {
  $__num = \func_num_args();
  $__res = function($f_8) use ($Bind1_6_4, $v_7) {
  $__num = \func_num_args();
  $__res = function($s_9) use ($Bind1_6_4, $f_8, $v_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_4)->{'bind'})(($v_7)($s_9)))(function($v1_10) use ($f_8) {
  $__num = \func_num_args();
  $__res = (($f_8)(($v1_10)->{'value0'}))(($v1_10)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_7) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_8_5 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_8_5 = (object)["map" => function($f_9) use ($__local_var_8_5) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($__local_var_8_5, $f_9) {
  $__num = \func_num_args();
  $__res = function($s_11) use ($__local_var_8_5, $f_9, $v_10) {
  $__num = \func_num_args();
  $__res = ((($__local_var_8_5)->{'map'})(function($v1_12) use ($f_9) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_9)(($v1_12)->{'value0'}), ($v1_12)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_10)($s_11));
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
  $Bind1_9_7 = (($dictMonad_0)->{'Bind1'})(null);
  $Bind1_9_7 = (object)["bind" => function($v_10) use ($Bind1_9_7) {
  $__num = \func_num_args();
  $__res = function($f_11) use ($Bind1_9_7, $v_10) {
  $__num = \func_num_args();
  $__res = function($s_12) use ($Bind1_9_7, $f_11, $v_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_7)->{'bind'})(($v_10)($s_12)))(function($v1_13) use ($f_11) {
  $__num = \func_num_args();
  $__res = (($f_11)(($v1_13)->{'value0'}))(($v1_13)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_10) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_10_9 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_11) use ($Applicative0_10_9, $Bind1_9_7) {
  $__num = \func_num_args();
  $__res = function($a_12) use ($Applicative0_10_9, $Bind1_9_7, $f_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_7)->{'bind'})($f_11))(function($f_prime__13) use ($Applicative0_10_9, $Bind1_9_7, $a_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_7)->{'bind'})($a_12))(function($a_prime__14) use ($Applicative0_10_9, $f_prime__13) {
  $__num = \func_num_args();
  $__res = (($Applicative0_10_9)->{'pure'})(($f_prime__13)($a_prime__14));
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
}, "Functor0" => function($_dollar___unused_9) use ($functorStateT1_8_5) {
  $__num = \func_num_args();
  $__res = $functorStateT1_8_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_7_11 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_8) use ($Applicative0_7_11, $Bind1_6_4) {
  $__num = \func_num_args();
  $__res = function($a_9) use ($Applicative0_7_11, $Bind1_6_4, $f_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_4)->{'bind'})($f_8))(function($f_prime__10) use ($Applicative0_7_11, $Bind1_6_4, $a_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_4)->{'bind'})($a_9))(function($a_prime__11) use ($Applicative0_7_11, $f_prime__10) {
  $__num = \func_num_args();
  $__res = (($Applicative0_7_11)->{'pure'})(($f_prime__10)($a_prime__11));
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
}, "Functor0" => function($_dollar___unused_6) use ($functorStateT1_5_2) {
  $__num = \func_num_args();
  $__res = $functorStateT1_5_2;
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
}, "Bind1" => function($_dollar___unused_2) use ($dictMonad_0) {
  $__num = \func_num_args();
  $Bind1_3_12 = (($dictMonad_0)->{'Bind1'})(null);
  $__res = (object)["bind" => function($v_4) use ($Bind1_3_12) {
  $__num = \func_num_args();
  $__res = function($f_5) use ($Bind1_3_12, $v_4) {
  $__num = \func_num_args();
  $__res = function($s_6) use ($Bind1_3_12, $f_5, $v_4) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_12)->{'bind'})(($v_4)($s_6)))(function($v1_7) use ($f_5) {
  $__num = \func_num_args();
  $__res = (($f_5)(($v1_7)->{'value0'}))(($v1_7)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_4) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_5_13 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_5_13 = (object)["map" => function($f_6) use ($__local_var_5_13) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($__local_var_5_13, $f_6) {
  $__num = \func_num_args();
  $__res = function($s_8) use ($__local_var_5_13, $f_6, $v_7) {
  $__num = \func_num_args();
  $__res = ((($__local_var_5_13)->{'map'})(function($v1_9) use ($f_6) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_6)(($v1_9)->{'value0'}), ($v1_9)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_7)($s_8));
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
  $Bind1_6_15 = (($dictMonad_0)->{'Bind1'})(null);
  $Bind1_6_15 = (object)["bind" => function($v_7) use ($Bind1_6_15) {
  $__num = \func_num_args();
  $__res = function($f_8) use ($Bind1_6_15, $v_7) {
  $__num = \func_num_args();
  $__res = function($s_9) use ($Bind1_6_15, $f_8, $v_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_15)->{'bind'})(($v_7)($s_9)))(function($v1_10) use ($f_8) {
  $__num = \func_num_args();
  $__res = (($f_8)(($v1_10)->{'value0'}))(($v1_10)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_7) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $pure_7_17 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $Applicative0_7_17 = (object)["pure" => function($a_8) use ($pure_7_17) {
  $__num = \func_num_args();
  $__res = function($s_9) use ($a_8, $pure_7_17) {
  $__num = \func_num_args();
  $__res = ($pure_7_17)(new \Data\Tuple\Data_Tuple_Tuple($a_8, $s_9));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_8) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_9_18 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_9_18 = (object)["map" => function($f_10) use ($__local_var_9_18) {
  $__num = \func_num_args();
  $__res = function($v_11) use ($__local_var_9_18, $f_10) {
  $__num = \func_num_args();
  $__res = function($s_12) use ($__local_var_9_18, $f_10, $v_11) {
  $__num = \func_num_args();
  $__res = ((($__local_var_9_18)->{'map'})(function($v1_13) use ($f_10) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_10)(($v1_13)->{'value0'}), ($v1_13)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_11)($s_12));
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
  $Bind1_10_20 = (($dictMonad_0)->{'Bind1'})(null);
  $Bind1_10_20 = (object)["bind" => function($v_11) use ($Bind1_10_20) {
  $__num = \func_num_args();
  $__res = function($f_12) use ($Bind1_10_20, $v_11) {
  $__num = \func_num_args();
  $__res = function($s_13) use ($Bind1_10_20, $f_12, $v_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_20)->{'bind'})(($v_11)($s_13)))(function($v1_14) use ($f_12) {
  $__num = \func_num_args();
  $__res = (($f_12)(($v1_14)->{'value0'}))(($v1_14)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_11) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_11_22 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_12) use ($Applicative0_11_22, $Bind1_10_20) {
  $__num = \func_num_args();
  $__res = function($a_13) use ($Applicative0_11_22, $Bind1_10_20, $f_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_20)->{'bind'})($f_12))(function($f_prime__14) use ($Applicative0_11_22, $Bind1_10_20, $a_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_20)->{'bind'})($a_13))(function($a_prime__15) use ($Applicative0_11_22, $f_prime__14) {
  $__num = \func_num_args();
  $__res = (($Applicative0_11_22)->{'pure'})(($f_prime__14)($a_prime__15));
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
}, "Functor0" => function($_dollar___unused_10) use ($functorStateT1_9_18) {
  $__num = \func_num_args();
  $__res = $functorStateT1_9_18;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_8) use ($Applicative0_7_17, $Bind1_6_15) {
  $__num = \func_num_args();
  $__res = function($a_9) use ($Applicative0_7_17, $Bind1_6_15, $f_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_15)->{'bind'})($f_8))(function($f_prime__10) use ($Applicative0_7_17, $Bind1_6_15, $a_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_15)->{'bind'})($a_9))(function($a_prime__11) use ($Applicative0_7_17, $f_prime__10) {
  $__num = \func_num_args();
  $__res = (($Applicative0_7_17)->{'pure'})(($f_prime__10)($a_prime__11));
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
}, "Functor0" => function($_dollar___unused_6) use ($functorStateT1_5_13) {
  $__num = \func_num_args();
  $__res = $functorStateT1_5_13;
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
}];
  $__res = (object)["state" => function($f_3) use ($pure_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_1_0))($f_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad0" => function($_dollar___unused_3) use ($monadStateT1_2_1) {
  $__num = \func_num_args();
  $__res = $monadStateT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_State_Trans_monadStateStateT'] = __NAMESPACE__ . '\\majControl_majMonad_majState_majTrans_monadmajStatemajStatemajT';

// Control_Monad_State_Trans_monadTellStateT
function majControl_majMonad_majState_majTrans_monadmajTellmajStatemajT($dictMonadTell_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majState_majTrans_monadmajTellmajStatemajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Monad1_1_0 = (($dictMonadTell_0)->{'Monad1'})(null);
  $Semigroup0_2_1 = (($dictMonadTell_0)->{'Semigroup0'})(null);
  $monadStateT1_3_2 = (object)["Applicative0" => function($_dollar___unused_3) use ($Monad1_1_0) {
  $__num = \func_num_args();
  $pure_4_2 = ((($Monad1_1_0)->{'Applicative0'})(null))->{'pure'};
  $__res = (object)["pure" => function($a_5) use ($pure_4_2) {
  $__num = \func_num_args();
  $__res = function($s_6) use ($a_5, $pure_4_2) {
  $__num = \func_num_args();
  $__res = ($pure_4_2)(new \Data\Tuple\Data_Tuple_Tuple($a_5, $s_6));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_5) use ($Monad1_1_0) {
  $__num = \func_num_args();
  $__local_var_6_3 = (((((($Monad1_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_6_3 = (object)["map" => function($f_7) use ($__local_var_6_3) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($__local_var_6_3, $f_7) {
  $__num = \func_num_args();
  $__res = function($s_9) use ($__local_var_6_3, $f_7, $v_8) {
  $__num = \func_num_args();
  $__res = ((($__local_var_6_3)->{'map'})(function($v1_10) use ($f_7) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_7)(($v1_10)->{'value0'}), ($v1_10)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_8)($s_9));
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
  $Bind1_7_5 = (($Monad1_1_0)->{'Bind1'})(null);
  $Bind1_7_5 = (object)["bind" => function($v_8) use ($Bind1_7_5) {
  $__num = \func_num_args();
  $__res = function($f_9) use ($Bind1_7_5, $v_8) {
  $__num = \func_num_args();
  $__res = function($s_10) use ($Bind1_7_5, $f_9, $v_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_5)->{'bind'})(($v_8)($s_10)))(function($v1_11) use ($f_9) {
  $__num = \func_num_args();
  $__res = (($f_9)(($v1_11)->{'value0'}))(($v1_11)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_8) use ($Monad1_1_0) {
  $__num = \func_num_args();
  $__local_var_9_6 = (((((($Monad1_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_9_6 = (object)["map" => function($f_10) use ($__local_var_9_6) {
  $__num = \func_num_args();
  $__res = function($v_11) use ($__local_var_9_6, $f_10) {
  $__num = \func_num_args();
  $__res = function($s_12) use ($__local_var_9_6, $f_10, $v_11) {
  $__num = \func_num_args();
  $__res = ((($__local_var_9_6)->{'map'})(function($v1_13) use ($f_10) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_10)(($v1_13)->{'value0'}), ($v1_13)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_11)($s_12));
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
  $Bind1_10_8 = (($Monad1_1_0)->{'Bind1'})(null);
  $Bind1_10_8 = (object)["bind" => function($v_11) use ($Bind1_10_8) {
  $__num = \func_num_args();
  $__res = function($f_12) use ($Bind1_10_8, $v_11) {
  $__num = \func_num_args();
  $__res = function($s_13) use ($Bind1_10_8, $f_12, $v_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_8)->{'bind'})(($v_11)($s_13)))(function($v1_14) use ($f_12) {
  $__num = \func_num_args();
  $__res = (($f_12)(($v1_14)->{'value0'}))(($v1_14)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_11) use ($Monad1_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($Monad1_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_11_10 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($Monad1_1_0);
  $__res = (object)["apply" => function($f_12) use ($Applicative0_11_10, $Bind1_10_8) {
  $__num = \func_num_args();
  $__res = function($a_13) use ($Applicative0_11_10, $Bind1_10_8, $f_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_8)->{'bind'})($f_12))(function($f_prime__14) use ($Applicative0_11_10, $Bind1_10_8, $a_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_8)->{'bind'})($a_13))(function($a_prime__15) use ($Applicative0_11_10, $f_prime__14) {
  $__num = \func_num_args();
  $__res = (($Applicative0_11_10)->{'pure'})(($f_prime__14)($a_prime__15));
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
}, "Functor0" => function($_dollar___unused_10) use ($functorStateT1_9_6) {
  $__num = \func_num_args();
  $__res = $functorStateT1_9_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_8_12 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($Monad1_1_0);
  $__res = (object)["apply" => function($f_9) use ($Applicative0_8_12, $Bind1_7_5) {
  $__num = \func_num_args();
  $__res = function($a_10) use ($Applicative0_8_12, $Bind1_7_5, $f_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_5)->{'bind'})($f_9))(function($f_prime__11) use ($Applicative0_8_12, $Bind1_7_5, $a_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_5)->{'bind'})($a_10))(function($a_prime__12) use ($Applicative0_8_12, $f_prime__11) {
  $__num = \func_num_args();
  $__res = (($Applicative0_8_12)->{'pure'})(($f_prime__11)($a_prime__12));
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
}, "Functor0" => function($_dollar___unused_7) use ($functorStateT1_6_3) {
  $__num = \func_num_args();
  $__res = $functorStateT1_6_3;
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
}, "Bind1" => function($_dollar___unused_3) use ($Monad1_1_0) {
  $__num = \func_num_args();
  $Bind1_4_13 = (($Monad1_1_0)->{'Bind1'})(null);
  $__res = (object)["bind" => function($v_5) use ($Bind1_4_13) {
  $__num = \func_num_args();
  $__res = function($f_6) use ($Bind1_4_13, $v_5) {
  $__num = \func_num_args();
  $__res = function($s_7) use ($Bind1_4_13, $f_6, $v_5) {
  $__num = \func_num_args();
  $__res = ((($Bind1_4_13)->{'bind'})(($v_5)($s_7)))(function($v1_8) use ($f_6) {
  $__num = \func_num_args();
  $__res = (($f_6)(($v1_8)->{'value0'}))(($v1_8)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_5) use ($Monad1_1_0) {
  $__num = \func_num_args();
  $__local_var_6_14 = (((((($Monad1_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_6_14 = (object)["map" => function($f_7) use ($__local_var_6_14) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($__local_var_6_14, $f_7) {
  $__num = \func_num_args();
  $__res = function($s_9) use ($__local_var_6_14, $f_7, $v_8) {
  $__num = \func_num_args();
  $__res = ((($__local_var_6_14)->{'map'})(function($v1_10) use ($f_7) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_7)(($v1_10)->{'value0'}), ($v1_10)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_8)($s_9));
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
  $Bind1_7_16 = (($Monad1_1_0)->{'Bind1'})(null);
  $Bind1_7_16 = (object)["bind" => function($v_8) use ($Bind1_7_16) {
  $__num = \func_num_args();
  $__res = function($f_9) use ($Bind1_7_16, $v_8) {
  $__num = \func_num_args();
  $__res = function($s_10) use ($Bind1_7_16, $f_9, $v_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_16)->{'bind'})(($v_8)($s_10)))(function($v1_11) use ($f_9) {
  $__num = \func_num_args();
  $__res = (($f_9)(($v1_11)->{'value0'}))(($v1_11)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_8) use ($Monad1_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($Monad1_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $pure_8_18 = ((($Monad1_1_0)->{'Applicative0'})(null))->{'pure'};
  $Applicative0_8_18 = (object)["pure" => function($a_9) use ($pure_8_18) {
  $__num = \func_num_args();
  $__res = function($s_10) use ($a_9, $pure_8_18) {
  $__num = \func_num_args();
  $__res = ($pure_8_18)(new \Data\Tuple\Data_Tuple_Tuple($a_9, $s_10));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_9) use ($Monad1_1_0) {
  $__num = \func_num_args();
  $__local_var_10_19 = (((((($Monad1_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_10_19 = (object)["map" => function($f_11) use ($__local_var_10_19) {
  $__num = \func_num_args();
  $__res = function($v_12) use ($__local_var_10_19, $f_11) {
  $__num = \func_num_args();
  $__res = function($s_13) use ($__local_var_10_19, $f_11, $v_12) {
  $__num = \func_num_args();
  $__res = ((($__local_var_10_19)->{'map'})(function($v1_14) use ($f_11) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_11)(($v1_14)->{'value0'}), ($v1_14)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_12)($s_13));
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
  $Bind1_11_21 = (($Monad1_1_0)->{'Bind1'})(null);
  $Bind1_11_21 = (object)["bind" => function($v_12) use ($Bind1_11_21) {
  $__num = \func_num_args();
  $__res = function($f_13) use ($Bind1_11_21, $v_12) {
  $__num = \func_num_args();
  $__res = function($s_14) use ($Bind1_11_21, $f_13, $v_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_21)->{'bind'})(($v_12)($s_14)))(function($v1_15) use ($f_13) {
  $__num = \func_num_args();
  $__res = (($f_13)(($v1_15)->{'value0'}))(($v1_15)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_12) use ($Monad1_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($Monad1_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_12_23 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($Monad1_1_0);
  $__res = (object)["apply" => function($f_13) use ($Applicative0_12_23, $Bind1_11_21) {
  $__num = \func_num_args();
  $__res = function($a_14) use ($Applicative0_12_23, $Bind1_11_21, $f_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_21)->{'bind'})($f_13))(function($f_prime__15) use ($Applicative0_12_23, $Bind1_11_21, $a_14) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_21)->{'bind'})($a_14))(function($a_prime__16) use ($Applicative0_12_23, $f_prime__15) {
  $__num = \func_num_args();
  $__res = (($Applicative0_12_23)->{'pure'})(($f_prime__15)($a_prime__16));
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
}, "Functor0" => function($_dollar___unused_11) use ($functorStateT1_10_19) {
  $__num = \func_num_args();
  $__res = $functorStateT1_10_19;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_9) use ($Applicative0_8_18, $Bind1_7_16) {
  $__num = \func_num_args();
  $__res = function($a_10) use ($Applicative0_8_18, $Bind1_7_16, $f_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_16)->{'bind'})($f_9))(function($f_prime__11) use ($Applicative0_8_18, $Bind1_7_16, $a_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_16)->{'bind'})($a_10))(function($a_prime__12) use ($Applicative0_8_18, $f_prime__11) {
  $__num = \func_num_args();
  $__res = (($Applicative0_8_18)->{'pure'})(($f_prime__11)($a_prime__12));
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
}, "Functor0" => function($_dollar___unused_7) use ($functorStateT1_6_14) {
  $__num = \func_num_args();
  $__res = $functorStateT1_6_14;
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
}];
  $Bind1_4_26 = (($Monad1_1_0)->{'Bind1'})(null);
  $pure_5_27 = ((($Monad1_1_0)->{'Applicative0'})(null))->{'pure'};
  $__res = (object)["tell" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($m_6) use ($Bind1_4_26, $pure_5_27) {
  $__num = \func_num_args();
  $__res = function($s_7) use ($Bind1_4_26, $m_6, $pure_5_27) {
  $__num = \func_num_args();
  $__res = ((($Bind1_4_26)->{'bind'})($m_6))(function($x_8) use ($pure_5_27, $s_7) {
  $__num = \func_num_args();
  $__res = ($pure_5_27)(new \Data\Tuple\Data_Tuple_Tuple($x_8, $s_7));
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
}))(($dictMonadTell_0)->{'tell'}), "Semigroup0" => function($_dollar___unused_4) use ($Semigroup0_2_1) {
  $__num = \func_num_args();
  $__res = $Semigroup0_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad1" => function($_dollar___unused_4) use ($monadStateT1_3_2) {
  $__num = \func_num_args();
  $__res = $monadStateT1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_State_Trans_monadTellStateT'] = __NAMESPACE__ . '\\majControl_majMonad_majState_majTrans_monadmajTellmajStatemajT';

// Control_Monad_State_Trans_monadWriterStateT
function majControl_majMonad_majState_majTrans_monadmajWritermajStatemajT($dictMonadWriter_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majState_majTrans_monadmajWritermajStatemajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $MonadTell1_1_0 = (($dictMonadWriter_0)->{'MonadTell1'})(null);
  $Monad1_2_1 = (($MonadTell1_1_0)->{'Monad1'})(null);
  $Bind1_3_2 = (($Monad1_2_1)->{'Bind1'})(null);
  $Applicative0_4_3 = (($Monad1_2_1)->{'Applicative0'})(null);
  $Monoid0_5_4 = (($dictMonadWriter_0)->{'Monoid0'})(null);
  $Monad1_6_5 = (($MonadTell1_1_0)->{'Monad1'})(null);
  $Semigroup0_7_6 = (($MonadTell1_1_0)->{'Semigroup0'})(null);
  $monadStateT1_8_7 = (object)["Applicative0" => function($_dollar___unused_8) use ($Monad1_6_5) {
  $__num = \func_num_args();
  $pure_9_7 = ((($Monad1_6_5)->{'Applicative0'})(null))->{'pure'};
  $__res = (object)["pure" => function($a_10) use ($pure_9_7) {
  $__num = \func_num_args();
  $__res = function($s_11) use ($a_10, $pure_9_7) {
  $__num = \func_num_args();
  $__res = ($pure_9_7)(new \Data\Tuple\Data_Tuple_Tuple($a_10, $s_11));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_10) use ($Monad1_6_5) {
  $__num = \func_num_args();
  $__local_var_11_8 = (((((($Monad1_6_5)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_11_8 = (object)["map" => function($f_12) use ($__local_var_11_8) {
  $__num = \func_num_args();
  $__res = function($v_13) use ($__local_var_11_8, $f_12) {
  $__num = \func_num_args();
  $__res = function($s_14) use ($__local_var_11_8, $f_12, $v_13) {
  $__num = \func_num_args();
  $__res = ((($__local_var_11_8)->{'map'})(function($v1_15) use ($f_12) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_12)(($v1_15)->{'value0'}), ($v1_15)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_13)($s_14));
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
  $Bind1_12_10 = (($Monad1_6_5)->{'Bind1'})(null);
  $Bind1_12_10 = (object)["bind" => function($v_13) use ($Bind1_12_10) {
  $__num = \func_num_args();
  $__res = function($f_14) use ($Bind1_12_10, $v_13) {
  $__num = \func_num_args();
  $__res = function($s_15) use ($Bind1_12_10, $f_14, $v_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_12_10)->{'bind'})(($v_13)($s_15)))(function($v1_16) use ($f_14) {
  $__num = \func_num_args();
  $__res = (($f_14)(($v1_16)->{'value0'}))(($v1_16)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_13) use ($Monad1_6_5) {
  $__num = \func_num_args();
  $__local_var_14_11 = (((((($Monad1_6_5)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_14_11 = (object)["map" => function($f_15) use ($__local_var_14_11) {
  $__num = \func_num_args();
  $__res = function($v_16) use ($__local_var_14_11, $f_15) {
  $__num = \func_num_args();
  $__res = function($s_17) use ($__local_var_14_11, $f_15, $v_16) {
  $__num = \func_num_args();
  $__res = ((($__local_var_14_11)->{'map'})(function($v1_18) use ($f_15) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_15)(($v1_18)->{'value0'}), ($v1_18)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_16)($s_17));
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
  $Bind1_15_13 = (($Monad1_6_5)->{'Bind1'})(null);
  $Bind1_15_13 = (object)["bind" => function($v_16) use ($Bind1_15_13) {
  $__num = \func_num_args();
  $__res = function($f_17) use ($Bind1_15_13, $v_16) {
  $__num = \func_num_args();
  $__res = function($s_18) use ($Bind1_15_13, $f_17, $v_16) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_13)->{'bind'})(($v_16)($s_18)))(function($v1_19) use ($f_17) {
  $__num = \func_num_args();
  $__res = (($f_17)(($v1_19)->{'value0'}))(($v1_19)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_16) use ($Monad1_6_5) {
  $__num = \func_num_args();
  $__local_var_17_14 = (((((($Monad1_6_5)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_17_14 = (object)["map" => function($f_18) use ($__local_var_17_14) {
  $__num = \func_num_args();
  $__res = function($v_19) use ($__local_var_17_14, $f_18) {
  $__num = \func_num_args();
  $__res = function($s_20) use ($__local_var_17_14, $f_18, $v_19) {
  $__num = \func_num_args();
  $__res = ((($__local_var_17_14)->{'map'})(function($v1_21) use ($f_18) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_18)(($v1_21)->{'value0'}), ($v1_21)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_19)($s_20));
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
  $Bind1_18_16 = (($Monad1_6_5)->{'Bind1'})(null);
  $Bind1_18_16 = (object)["bind" => function($v_19) use ($Bind1_18_16) {
  $__num = \func_num_args();
  $__res = function($f_20) use ($Bind1_18_16, $v_19) {
  $__num = \func_num_args();
  $__res = function($s_21) use ($Bind1_18_16, $f_20, $v_19) {
  $__num = \func_num_args();
  $__res = ((($Bind1_18_16)->{'bind'})(($v_19)($s_21)))(function($v1_22) use ($f_20) {
  $__num = \func_num_args();
  $__res = (($f_20)(($v1_22)->{'value0'}))(($v1_22)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_19) use ($Monad1_6_5) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($Monad1_6_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $pure_19_18 = ((($Monad1_6_5)->{'Applicative0'})(null))->{'pure'};
  $Applicative0_19_18 = (object)["pure" => function($a_20) use ($pure_19_18) {
  $__num = \func_num_args();
  $__res = function($s_21) use ($a_20, $pure_19_18) {
  $__num = \func_num_args();
  $__res = ($pure_19_18)(new \Data\Tuple\Data_Tuple_Tuple($a_20, $s_21));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_20) use ($Monad1_6_5) {
  $__num = \func_num_args();
  $__local_var_21_19 = (((((($Monad1_6_5)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_21_19 = (object)["map" => function($f_22) use ($__local_var_21_19) {
  $__num = \func_num_args();
  $__res = function($v_23) use ($__local_var_21_19, $f_22) {
  $__num = \func_num_args();
  $__res = function($s_24) use ($__local_var_21_19, $f_22, $v_23) {
  $__num = \func_num_args();
  $__res = ((($__local_var_21_19)->{'map'})(function($v1_25) use ($f_22) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_22)(($v1_25)->{'value0'}), ($v1_25)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_23)($s_24));
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
  $Bind1_22_21 = (($Monad1_6_5)->{'Bind1'})(null);
  $Bind1_22_21 = (object)["bind" => function($v_23) use ($Bind1_22_21) {
  $__num = \func_num_args();
  $__res = function($f_24) use ($Bind1_22_21, $v_23) {
  $__num = \func_num_args();
  $__res = function($s_25) use ($Bind1_22_21, $f_24, $v_23) {
  $__num = \func_num_args();
  $__res = ((($Bind1_22_21)->{'bind'})(($v_23)($s_25)))(function($v1_26) use ($f_24) {
  $__num = \func_num_args();
  $__res = (($f_24)(($v1_26)->{'value0'}))(($v1_26)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_23) use ($Monad1_6_5) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($Monad1_6_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_23_23 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($Monad1_6_5);
  $__res = (object)["apply" => function($f_24) use ($Applicative0_23_23, $Bind1_22_21) {
  $__num = \func_num_args();
  $__res = function($a_25) use ($Applicative0_23_23, $Bind1_22_21, $f_24) {
  $__num = \func_num_args();
  $__res = ((($Bind1_22_21)->{'bind'})($f_24))(function($f_prime__26) use ($Applicative0_23_23, $Bind1_22_21, $a_25) {
  $__num = \func_num_args();
  $__res = ((($Bind1_22_21)->{'bind'})($a_25))(function($a_prime__27) use ($Applicative0_23_23, $f_prime__26) {
  $__num = \func_num_args();
  $__res = (($Applicative0_23_23)->{'pure'})(($f_prime__26)($a_prime__27));
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
}, "Functor0" => function($_dollar___unused_22) use ($functorStateT1_21_19) {
  $__num = \func_num_args();
  $__res = $functorStateT1_21_19;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_20) use ($Applicative0_19_18, $Bind1_18_16) {
  $__num = \func_num_args();
  $__res = function($a_21) use ($Applicative0_19_18, $Bind1_18_16, $f_20) {
  $__num = \func_num_args();
  $__res = ((($Bind1_18_16)->{'bind'})($f_20))(function($f_prime__22) use ($Applicative0_19_18, $Bind1_18_16, $a_21) {
  $__num = \func_num_args();
  $__res = ((($Bind1_18_16)->{'bind'})($a_21))(function($a_prime__23) use ($Applicative0_19_18, $f_prime__22) {
  $__num = \func_num_args();
  $__res = (($Applicative0_19_18)->{'pure'})(($f_prime__22)($a_prime__23));
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
}, "Functor0" => function($_dollar___unused_18) use ($functorStateT1_17_14) {
  $__num = \func_num_args();
  $__res = $functorStateT1_17_14;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $pure_16_26 = ((($Monad1_6_5)->{'Applicative0'})(null))->{'pure'};
  $Applicative0_16_26 = (object)["pure" => function($a_17) use ($pure_16_26) {
  $__num = \func_num_args();
  $__res = function($s_18) use ($a_17, $pure_16_26) {
  $__num = \func_num_args();
  $__res = ($pure_16_26)(new \Data\Tuple\Data_Tuple_Tuple($a_17, $s_18));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_17) use ($Monad1_6_5) {
  $__num = \func_num_args();
  $__local_var_18_27 = (((((($Monad1_6_5)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_18_27 = (object)["map" => function($f_19) use ($__local_var_18_27) {
  $__num = \func_num_args();
  $__res = function($v_20) use ($__local_var_18_27, $f_19) {
  $__num = \func_num_args();
  $__res = function($s_21) use ($__local_var_18_27, $f_19, $v_20) {
  $__num = \func_num_args();
  $__res = ((($__local_var_18_27)->{'map'})(function($v1_22) use ($f_19) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_19)(($v1_22)->{'value0'}), ($v1_22)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_20)($s_21));
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
  $Bind1_19_29 = (($Monad1_6_5)->{'Bind1'})(null);
  $Bind1_19_29 = (object)["bind" => function($v_20) use ($Bind1_19_29) {
  $__num = \func_num_args();
  $__res = function($f_21) use ($Bind1_19_29, $v_20) {
  $__num = \func_num_args();
  $__res = function($s_22) use ($Bind1_19_29, $f_21, $v_20) {
  $__num = \func_num_args();
  $__res = ((($Bind1_19_29)->{'bind'})(($v_20)($s_22)))(function($v1_23) use ($f_21) {
  $__num = \func_num_args();
  $__res = (($f_21)(($v1_23)->{'value0'}))(($v1_23)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_20) use ($Monad1_6_5) {
  $__num = \func_num_args();
  $__local_var_21_30 = (((((($Monad1_6_5)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_21_30 = (object)["map" => function($f_22) use ($__local_var_21_30) {
  $__num = \func_num_args();
  $__res = function($v_23) use ($__local_var_21_30, $f_22) {
  $__num = \func_num_args();
  $__res = function($s_24) use ($__local_var_21_30, $f_22, $v_23) {
  $__num = \func_num_args();
  $__res = ((($__local_var_21_30)->{'map'})(function($v1_25) use ($f_22) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_22)(($v1_25)->{'value0'}), ($v1_25)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_23)($s_24));
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
  $Bind1_22_32 = (($Monad1_6_5)->{'Bind1'})(null);
  $Bind1_22_32 = (object)["bind" => function($v_23) use ($Bind1_22_32) {
  $__num = \func_num_args();
  $__res = function($f_24) use ($Bind1_22_32, $v_23) {
  $__num = \func_num_args();
  $__res = function($s_25) use ($Bind1_22_32, $f_24, $v_23) {
  $__num = \func_num_args();
  $__res = ((($Bind1_22_32)->{'bind'})(($v_23)($s_25)))(function($v1_26) use ($f_24) {
  $__num = \func_num_args();
  $__res = (($f_24)(($v1_26)->{'value0'}))(($v1_26)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_23) use ($Monad1_6_5) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($Monad1_6_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_23_34 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($Monad1_6_5);
  $__res = (object)["apply" => function($f_24) use ($Applicative0_23_34, $Bind1_22_32) {
  $__num = \func_num_args();
  $__res = function($a_25) use ($Applicative0_23_34, $Bind1_22_32, $f_24) {
  $__num = \func_num_args();
  $__res = ((($Bind1_22_32)->{'bind'})($f_24))(function($f_prime__26) use ($Applicative0_23_34, $Bind1_22_32, $a_25) {
  $__num = \func_num_args();
  $__res = ((($Bind1_22_32)->{'bind'})($a_25))(function($a_prime__27) use ($Applicative0_23_34, $f_prime__26) {
  $__num = \func_num_args();
  $__res = (($Applicative0_23_34)->{'pure'})(($f_prime__26)($a_prime__27));
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
}, "Functor0" => function($_dollar___unused_22) use ($functorStateT1_21_30) {
  $__num = \func_num_args();
  $__res = $functorStateT1_21_30;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_20_36 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($Monad1_6_5);
  $__res = (object)["apply" => function($f_21) use ($Applicative0_20_36, $Bind1_19_29) {
  $__num = \func_num_args();
  $__res = function($a_22) use ($Applicative0_20_36, $Bind1_19_29, $f_21) {
  $__num = \func_num_args();
  $__res = ((($Bind1_19_29)->{'bind'})($f_21))(function($f_prime__23) use ($Applicative0_20_36, $Bind1_19_29, $a_22) {
  $__num = \func_num_args();
  $__res = ((($Bind1_19_29)->{'bind'})($a_22))(function($a_prime__24) use ($Applicative0_20_36, $f_prime__23) {
  $__num = \func_num_args();
  $__res = (($Applicative0_20_36)->{'pure'})(($f_prime__23)($a_prime__24));
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
}, "Functor0" => function($_dollar___unused_19) use ($functorStateT1_18_27) {
  $__num = \func_num_args();
  $__res = $functorStateT1_18_27;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_17) use ($Applicative0_16_26, $Bind1_15_13) {
  $__num = \func_num_args();
  $__res = function($a_18) use ($Applicative0_16_26, $Bind1_15_13, $f_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_13)->{'bind'})($f_17))(function($f_prime__19) use ($Applicative0_16_26, $Bind1_15_13, $a_18) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_13)->{'bind'})($a_18))(function($a_prime__20) use ($Applicative0_16_26, $f_prime__19) {
  $__num = \func_num_args();
  $__res = (($Applicative0_16_26)->{'pure'})(($f_prime__19)($a_prime__20));
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
}, "Functor0" => function($_dollar___unused_15) use ($functorStateT1_14_11) {
  $__num = \func_num_args();
  $__res = $functorStateT1_14_11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $pure_13_39 = ((($Monad1_6_5)->{'Applicative0'})(null))->{'pure'};
  $Applicative0_13_39 = (object)["pure" => function($a_14) use ($pure_13_39) {
  $__num = \func_num_args();
  $__res = function($s_15) use ($a_14, $pure_13_39) {
  $__num = \func_num_args();
  $__res = ($pure_13_39)(new \Data\Tuple\Data_Tuple_Tuple($a_14, $s_15));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_14) use ($Monad1_6_5) {
  $__num = \func_num_args();
  $__local_var_15_40 = (((((($Monad1_6_5)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_15_40 = (object)["map" => function($f_16) use ($__local_var_15_40) {
  $__num = \func_num_args();
  $__res = function($v_17) use ($__local_var_15_40, $f_16) {
  $__num = \func_num_args();
  $__res = function($s_18) use ($__local_var_15_40, $f_16, $v_17) {
  $__num = \func_num_args();
  $__res = ((($__local_var_15_40)->{'map'})(function($v1_19) use ($f_16) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_16)(($v1_19)->{'value0'}), ($v1_19)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_17)($s_18));
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
  $Bind1_16_42 = (($Monad1_6_5)->{'Bind1'})(null);
  $Bind1_16_42 = (object)["bind" => function($v_17) use ($Bind1_16_42) {
  $__num = \func_num_args();
  $__res = function($f_18) use ($Bind1_16_42, $v_17) {
  $__num = \func_num_args();
  $__res = function($s_19) use ($Bind1_16_42, $f_18, $v_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_16_42)->{'bind'})(($v_17)($s_19)))(function($v1_20) use ($f_18) {
  $__num = \func_num_args();
  $__res = (($f_18)(($v1_20)->{'value0'}))(($v1_20)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_17) use ($Monad1_6_5) {
  $__num = \func_num_args();
  $__local_var_18_43 = (((((($Monad1_6_5)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_18_43 = (object)["map" => function($f_19) use ($__local_var_18_43) {
  $__num = \func_num_args();
  $__res = function($v_20) use ($__local_var_18_43, $f_19) {
  $__num = \func_num_args();
  $__res = function($s_21) use ($__local_var_18_43, $f_19, $v_20) {
  $__num = \func_num_args();
  $__res = ((($__local_var_18_43)->{'map'})(function($v1_22) use ($f_19) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_19)(($v1_22)->{'value0'}), ($v1_22)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_20)($s_21));
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
  $Bind1_19_45 = (($Monad1_6_5)->{'Bind1'})(null);
  $Bind1_19_45 = (object)["bind" => function($v_20) use ($Bind1_19_45) {
  $__num = \func_num_args();
  $__res = function($f_21) use ($Bind1_19_45, $v_20) {
  $__num = \func_num_args();
  $__res = function($s_22) use ($Bind1_19_45, $f_21, $v_20) {
  $__num = \func_num_args();
  $__res = ((($Bind1_19_45)->{'bind'})(($v_20)($s_22)))(function($v1_23) use ($f_21) {
  $__num = \func_num_args();
  $__res = (($f_21)(($v1_23)->{'value0'}))(($v1_23)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_20) use ($Monad1_6_5) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($Monad1_6_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_20_47 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($Monad1_6_5);
  $__res = (object)["apply" => function($f_21) use ($Applicative0_20_47, $Bind1_19_45) {
  $__num = \func_num_args();
  $__res = function($a_22) use ($Applicative0_20_47, $Bind1_19_45, $f_21) {
  $__num = \func_num_args();
  $__res = ((($Bind1_19_45)->{'bind'})($f_21))(function($f_prime__23) use ($Applicative0_20_47, $Bind1_19_45, $a_22) {
  $__num = \func_num_args();
  $__res = ((($Bind1_19_45)->{'bind'})($a_22))(function($a_prime__24) use ($Applicative0_20_47, $f_prime__23) {
  $__num = \func_num_args();
  $__res = (($Applicative0_20_47)->{'pure'})(($f_prime__23)($a_prime__24));
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
}, "Functor0" => function($_dollar___unused_19) use ($functorStateT1_18_43) {
  $__num = \func_num_args();
  $__res = $functorStateT1_18_43;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_17_49 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($Monad1_6_5);
  $__res = (object)["apply" => function($f_18) use ($Applicative0_17_49, $Bind1_16_42) {
  $__num = \func_num_args();
  $__res = function($a_19) use ($Applicative0_17_49, $Bind1_16_42, $f_18) {
  $__num = \func_num_args();
  $__res = ((($Bind1_16_42)->{'bind'})($f_18))(function($f_prime__20) use ($Applicative0_17_49, $Bind1_16_42, $a_19) {
  $__num = \func_num_args();
  $__res = ((($Bind1_16_42)->{'bind'})($a_19))(function($a_prime__21) use ($Applicative0_17_49, $f_prime__20) {
  $__num = \func_num_args();
  $__res = (($Applicative0_17_49)->{'pure'})(($f_prime__20)($a_prime__21));
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
}, "Functor0" => function($_dollar___unused_16) use ($functorStateT1_15_40) {
  $__num = \func_num_args();
  $__res = $functorStateT1_15_40;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_14) use ($Applicative0_13_39, $Bind1_12_10) {
  $__num = \func_num_args();
  $__res = function($a_15) use ($Applicative0_13_39, $Bind1_12_10, $f_14) {
  $__num = \func_num_args();
  $__res = ((($Bind1_12_10)->{'bind'})($f_14))(function($f_prime__16) use ($Applicative0_13_39, $Bind1_12_10, $a_15) {
  $__num = \func_num_args();
  $__res = ((($Bind1_12_10)->{'bind'})($a_15))(function($a_prime__17) use ($Applicative0_13_39, $f_prime__16) {
  $__num = \func_num_args();
  $__res = (($Applicative0_13_39)->{'pure'})(($f_prime__16)($a_prime__17));
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
}, "Functor0" => function($_dollar___unused_12) use ($functorStateT1_11_8) {
  $__num = \func_num_args();
  $__res = $functorStateT1_11_8;
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
}, "Bind1" => function($_dollar___unused_8) use ($Monad1_6_5) {
  $__num = \func_num_args();
  $Bind1_9_51 = (($Monad1_6_5)->{'Bind1'})(null);
  $__res = (object)["bind" => function($v_10) use ($Bind1_9_51) {
  $__num = \func_num_args();
  $__res = function($f_11) use ($Bind1_9_51, $v_10) {
  $__num = \func_num_args();
  $__res = function($s_12) use ($Bind1_9_51, $f_11, $v_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_51)->{'bind'})(($v_10)($s_12)))(function($v1_13) use ($f_11) {
  $__num = \func_num_args();
  $__res = (($f_11)(($v1_13)->{'value0'}))(($v1_13)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_10) use ($Monad1_6_5) {
  $__num = \func_num_args();
  $__local_var_11_52 = (((((($Monad1_6_5)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_11_52 = (object)["map" => function($f_12) use ($__local_var_11_52) {
  $__num = \func_num_args();
  $__res = function($v_13) use ($__local_var_11_52, $f_12) {
  $__num = \func_num_args();
  $__res = function($s_14) use ($__local_var_11_52, $f_12, $v_13) {
  $__num = \func_num_args();
  $__res = ((($__local_var_11_52)->{'map'})(function($v1_15) use ($f_12) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_12)(($v1_15)->{'value0'}), ($v1_15)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_13)($s_14));
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
  $Bind1_12_54 = (($Monad1_6_5)->{'Bind1'})(null);
  $Bind1_12_54 = (object)["bind" => function($v_13) use ($Bind1_12_54) {
  $__num = \func_num_args();
  $__res = function($f_14) use ($Bind1_12_54, $v_13) {
  $__num = \func_num_args();
  $__res = function($s_15) use ($Bind1_12_54, $f_14, $v_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_12_54)->{'bind'})(($v_13)($s_15)))(function($v1_16) use ($f_14) {
  $__num = \func_num_args();
  $__res = (($f_14)(($v1_16)->{'value0'}))(($v1_16)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_13) use ($Monad1_6_5) {
  $__num = \func_num_args();
  $__local_var_14_55 = (((((($Monad1_6_5)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_14_55 = (object)["map" => function($f_15) use ($__local_var_14_55) {
  $__num = \func_num_args();
  $__res = function($v_16) use ($__local_var_14_55, $f_15) {
  $__num = \func_num_args();
  $__res = function($s_17) use ($__local_var_14_55, $f_15, $v_16) {
  $__num = \func_num_args();
  $__res = ((($__local_var_14_55)->{'map'})(function($v1_18) use ($f_15) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_15)(($v1_18)->{'value0'}), ($v1_18)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_16)($s_17));
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
  $Bind1_15_57 = (($Monad1_6_5)->{'Bind1'})(null);
  $Bind1_15_57 = (object)["bind" => function($v_16) use ($Bind1_15_57) {
  $__num = \func_num_args();
  $__res = function($f_17) use ($Bind1_15_57, $v_16) {
  $__num = \func_num_args();
  $__res = function($s_18) use ($Bind1_15_57, $f_17, $v_16) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_57)->{'bind'})(($v_16)($s_18)))(function($v1_19) use ($f_17) {
  $__num = \func_num_args();
  $__res = (($f_17)(($v1_19)->{'value0'}))(($v1_19)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_16) use ($Monad1_6_5) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($Monad1_6_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $pure_16_59 = ((($Monad1_6_5)->{'Applicative0'})(null))->{'pure'};
  $Applicative0_16_59 = (object)["pure" => function($a_17) use ($pure_16_59) {
  $__num = \func_num_args();
  $__res = function($s_18) use ($a_17, $pure_16_59) {
  $__num = \func_num_args();
  $__res = ($pure_16_59)(new \Data\Tuple\Data_Tuple_Tuple($a_17, $s_18));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_17) use ($Monad1_6_5) {
  $__num = \func_num_args();
  $__local_var_18_60 = (((((($Monad1_6_5)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_18_60 = (object)["map" => function($f_19) use ($__local_var_18_60) {
  $__num = \func_num_args();
  $__res = function($v_20) use ($__local_var_18_60, $f_19) {
  $__num = \func_num_args();
  $__res = function($s_21) use ($__local_var_18_60, $f_19, $v_20) {
  $__num = \func_num_args();
  $__res = ((($__local_var_18_60)->{'map'})(function($v1_22) use ($f_19) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_19)(($v1_22)->{'value0'}), ($v1_22)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_20)($s_21));
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
  $Bind1_19_62 = (($Monad1_6_5)->{'Bind1'})(null);
  $Bind1_19_62 = (object)["bind" => function($v_20) use ($Bind1_19_62) {
  $__num = \func_num_args();
  $__res = function($f_21) use ($Bind1_19_62, $v_20) {
  $__num = \func_num_args();
  $__res = function($s_22) use ($Bind1_19_62, $f_21, $v_20) {
  $__num = \func_num_args();
  $__res = ((($Bind1_19_62)->{'bind'})(($v_20)($s_22)))(function($v1_23) use ($f_21) {
  $__num = \func_num_args();
  $__res = (($f_21)(($v1_23)->{'value0'}))(($v1_23)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_20) use ($Monad1_6_5) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($Monad1_6_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_20_64 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($Monad1_6_5);
  $__res = (object)["apply" => function($f_21) use ($Applicative0_20_64, $Bind1_19_62) {
  $__num = \func_num_args();
  $__res = function($a_22) use ($Applicative0_20_64, $Bind1_19_62, $f_21) {
  $__num = \func_num_args();
  $__res = ((($Bind1_19_62)->{'bind'})($f_21))(function($f_prime__23) use ($Applicative0_20_64, $Bind1_19_62, $a_22) {
  $__num = \func_num_args();
  $__res = ((($Bind1_19_62)->{'bind'})($a_22))(function($a_prime__24) use ($Applicative0_20_64, $f_prime__23) {
  $__num = \func_num_args();
  $__res = (($Applicative0_20_64)->{'pure'})(($f_prime__23)($a_prime__24));
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
}, "Functor0" => function($_dollar___unused_19) use ($functorStateT1_18_60) {
  $__num = \func_num_args();
  $__res = $functorStateT1_18_60;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_17) use ($Applicative0_16_59, $Bind1_15_57) {
  $__num = \func_num_args();
  $__res = function($a_18) use ($Applicative0_16_59, $Bind1_15_57, $f_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_57)->{'bind'})($f_17))(function($f_prime__19) use ($Applicative0_16_59, $Bind1_15_57, $a_18) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_57)->{'bind'})($a_18))(function($a_prime__20) use ($Applicative0_16_59, $f_prime__19) {
  $__num = \func_num_args();
  $__res = (($Applicative0_16_59)->{'pure'})(($f_prime__19)($a_prime__20));
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
}, "Functor0" => function($_dollar___unused_15) use ($functorStateT1_14_55) {
  $__num = \func_num_args();
  $__res = $functorStateT1_14_55;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $pure_13_67 = ((($Monad1_6_5)->{'Applicative0'})(null))->{'pure'};
  $Applicative0_13_67 = (object)["pure" => function($a_14) use ($pure_13_67) {
  $__num = \func_num_args();
  $__res = function($s_15) use ($a_14, $pure_13_67) {
  $__num = \func_num_args();
  $__res = ($pure_13_67)(new \Data\Tuple\Data_Tuple_Tuple($a_14, $s_15));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_14) use ($Monad1_6_5) {
  $__num = \func_num_args();
  $__local_var_15_68 = (((((($Monad1_6_5)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_15_68 = (object)["map" => function($f_16) use ($__local_var_15_68) {
  $__num = \func_num_args();
  $__res = function($v_17) use ($__local_var_15_68, $f_16) {
  $__num = \func_num_args();
  $__res = function($s_18) use ($__local_var_15_68, $f_16, $v_17) {
  $__num = \func_num_args();
  $__res = ((($__local_var_15_68)->{'map'})(function($v1_19) use ($f_16) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_16)(($v1_19)->{'value0'}), ($v1_19)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_17)($s_18));
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
  $Bind1_16_70 = (($Monad1_6_5)->{'Bind1'})(null);
  $Bind1_16_70 = (object)["bind" => function($v_17) use ($Bind1_16_70) {
  $__num = \func_num_args();
  $__res = function($f_18) use ($Bind1_16_70, $v_17) {
  $__num = \func_num_args();
  $__res = function($s_19) use ($Bind1_16_70, $f_18, $v_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_16_70)->{'bind'})(($v_17)($s_19)))(function($v1_20) use ($f_18) {
  $__num = \func_num_args();
  $__res = (($f_18)(($v1_20)->{'value0'}))(($v1_20)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_17) use ($Monad1_6_5) {
  $__num = \func_num_args();
  $__local_var_18_71 = (((((($Monad1_6_5)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_18_71 = (object)["map" => function($f_19) use ($__local_var_18_71) {
  $__num = \func_num_args();
  $__res = function($v_20) use ($__local_var_18_71, $f_19) {
  $__num = \func_num_args();
  $__res = function($s_21) use ($__local_var_18_71, $f_19, $v_20) {
  $__num = \func_num_args();
  $__res = ((($__local_var_18_71)->{'map'})(function($v1_22) use ($f_19) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_19)(($v1_22)->{'value0'}), ($v1_22)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_20)($s_21));
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
  $Bind1_19_73 = (($Monad1_6_5)->{'Bind1'})(null);
  $Bind1_19_73 = (object)["bind" => function($v_20) use ($Bind1_19_73) {
  $__num = \func_num_args();
  $__res = function($f_21) use ($Bind1_19_73, $v_20) {
  $__num = \func_num_args();
  $__res = function($s_22) use ($Bind1_19_73, $f_21, $v_20) {
  $__num = \func_num_args();
  $__res = ((($Bind1_19_73)->{'bind'})(($v_20)($s_22)))(function($v1_23) use ($f_21) {
  $__num = \func_num_args();
  $__res = (($f_21)(($v1_23)->{'value0'}))(($v1_23)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_20) use ($Monad1_6_5) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($Monad1_6_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $pure_20_75 = ((($Monad1_6_5)->{'Applicative0'})(null))->{'pure'};
  $Applicative0_20_75 = (object)["pure" => function($a_21) use ($pure_20_75) {
  $__num = \func_num_args();
  $__res = function($s_22) use ($a_21, $pure_20_75) {
  $__num = \func_num_args();
  $__res = ($pure_20_75)(new \Data\Tuple\Data_Tuple_Tuple($a_21, $s_22));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_21) use ($Monad1_6_5) {
  $__num = \func_num_args();
  $__local_var_22_76 = (((((($Monad1_6_5)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_22_76 = (object)["map" => function($f_23) use ($__local_var_22_76) {
  $__num = \func_num_args();
  $__res = function($v_24) use ($__local_var_22_76, $f_23) {
  $__num = \func_num_args();
  $__res = function($s_25) use ($__local_var_22_76, $f_23, $v_24) {
  $__num = \func_num_args();
  $__res = ((($__local_var_22_76)->{'map'})(function($v1_26) use ($f_23) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_23)(($v1_26)->{'value0'}), ($v1_26)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_24)($s_25));
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
  $Bind1_23_78 = (($Monad1_6_5)->{'Bind1'})(null);
  $Bind1_23_78 = (object)["bind" => function($v_24) use ($Bind1_23_78) {
  $__num = \func_num_args();
  $__res = function($f_25) use ($Bind1_23_78, $v_24) {
  $__num = \func_num_args();
  $__res = function($s_26) use ($Bind1_23_78, $f_25, $v_24) {
  $__num = \func_num_args();
  $__res = ((($Bind1_23_78)->{'bind'})(($v_24)($s_26)))(function($v1_27) use ($f_25) {
  $__num = \func_num_args();
  $__res = (($f_25)(($v1_27)->{'value0'}))(($v1_27)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_24) use ($Monad1_6_5) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($Monad1_6_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_24_80 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($Monad1_6_5);
  $__res = (object)["apply" => function($f_25) use ($Applicative0_24_80, $Bind1_23_78) {
  $__num = \func_num_args();
  $__res = function($a_26) use ($Applicative0_24_80, $Bind1_23_78, $f_25) {
  $__num = \func_num_args();
  $__res = ((($Bind1_23_78)->{'bind'})($f_25))(function($f_prime__27) use ($Applicative0_24_80, $Bind1_23_78, $a_26) {
  $__num = \func_num_args();
  $__res = ((($Bind1_23_78)->{'bind'})($a_26))(function($a_prime__28) use ($Applicative0_24_80, $f_prime__27) {
  $__num = \func_num_args();
  $__res = (($Applicative0_24_80)->{'pure'})(($f_prime__27)($a_prime__28));
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
}, "Functor0" => function($_dollar___unused_23) use ($functorStateT1_22_76) {
  $__num = \func_num_args();
  $__res = $functorStateT1_22_76;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_21) use ($Applicative0_20_75, $Bind1_19_73) {
  $__num = \func_num_args();
  $__res = function($a_22) use ($Applicative0_20_75, $Bind1_19_73, $f_21) {
  $__num = \func_num_args();
  $__res = ((($Bind1_19_73)->{'bind'})($f_21))(function($f_prime__23) use ($Applicative0_20_75, $Bind1_19_73, $a_22) {
  $__num = \func_num_args();
  $__res = ((($Bind1_19_73)->{'bind'})($a_22))(function($a_prime__24) use ($Applicative0_20_75, $f_prime__23) {
  $__num = \func_num_args();
  $__res = (($Applicative0_20_75)->{'pure'})(($f_prime__23)($a_prime__24));
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
}, "Functor0" => function($_dollar___unused_19) use ($functorStateT1_18_71) {
  $__num = \func_num_args();
  $__res = $functorStateT1_18_71;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $pure_17_83 = ((($Monad1_6_5)->{'Applicative0'})(null))->{'pure'};
  $Applicative0_17_83 = (object)["pure" => function($a_18) use ($pure_17_83) {
  $__num = \func_num_args();
  $__res = function($s_19) use ($a_18, $pure_17_83) {
  $__num = \func_num_args();
  $__res = ($pure_17_83)(new \Data\Tuple\Data_Tuple_Tuple($a_18, $s_19));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_18) use ($Monad1_6_5) {
  $__num = \func_num_args();
  $__local_var_19_84 = (((((($Monad1_6_5)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_19_84 = (object)["map" => function($f_20) use ($__local_var_19_84) {
  $__num = \func_num_args();
  $__res = function($v_21) use ($__local_var_19_84, $f_20) {
  $__num = \func_num_args();
  $__res = function($s_22) use ($__local_var_19_84, $f_20, $v_21) {
  $__num = \func_num_args();
  $__res = ((($__local_var_19_84)->{'map'})(function($v1_23) use ($f_20) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_20)(($v1_23)->{'value0'}), ($v1_23)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_21)($s_22));
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
  $Bind1_20_86 = (($Monad1_6_5)->{'Bind1'})(null);
  $Bind1_20_86 = (object)["bind" => function($v_21) use ($Bind1_20_86) {
  $__num = \func_num_args();
  $__res = function($f_22) use ($Bind1_20_86, $v_21) {
  $__num = \func_num_args();
  $__res = function($s_23) use ($Bind1_20_86, $f_22, $v_21) {
  $__num = \func_num_args();
  $__res = ((($Bind1_20_86)->{'bind'})(($v_21)($s_23)))(function($v1_24) use ($f_22) {
  $__num = \func_num_args();
  $__res = (($f_22)(($v1_24)->{'value0'}))(($v1_24)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_21) use ($Monad1_6_5) {
  $__num = \func_num_args();
  $__local_var_22_87 = (((((($Monad1_6_5)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_22_87 = (object)["map" => function($f_23) use ($__local_var_22_87) {
  $__num = \func_num_args();
  $__res = function($v_24) use ($__local_var_22_87, $f_23) {
  $__num = \func_num_args();
  $__res = function($s_25) use ($__local_var_22_87, $f_23, $v_24) {
  $__num = \func_num_args();
  $__res = ((($__local_var_22_87)->{'map'})(function($v1_26) use ($f_23) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_23)(($v1_26)->{'value0'}), ($v1_26)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_24)($s_25));
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
  $Bind1_23_89 = (($Monad1_6_5)->{'Bind1'})(null);
  $Bind1_23_89 = (object)["bind" => function($v_24) use ($Bind1_23_89) {
  $__num = \func_num_args();
  $__res = function($f_25) use ($Bind1_23_89, $v_24) {
  $__num = \func_num_args();
  $__res = function($s_26) use ($Bind1_23_89, $f_25, $v_24) {
  $__num = \func_num_args();
  $__res = ((($Bind1_23_89)->{'bind'})(($v_24)($s_26)))(function($v1_27) use ($f_25) {
  $__num = \func_num_args();
  $__res = (($f_25)(($v1_27)->{'value0'}))(($v1_27)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_24) use ($Monad1_6_5) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($Monad1_6_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_24_91 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($Monad1_6_5);
  $__res = (object)["apply" => function($f_25) use ($Applicative0_24_91, $Bind1_23_89) {
  $__num = \func_num_args();
  $__res = function($a_26) use ($Applicative0_24_91, $Bind1_23_89, $f_25) {
  $__num = \func_num_args();
  $__res = ((($Bind1_23_89)->{'bind'})($f_25))(function($f_prime__27) use ($Applicative0_24_91, $Bind1_23_89, $a_26) {
  $__num = \func_num_args();
  $__res = ((($Bind1_23_89)->{'bind'})($a_26))(function($a_prime__28) use ($Applicative0_24_91, $f_prime__27) {
  $__num = \func_num_args();
  $__res = (($Applicative0_24_91)->{'pure'})(($f_prime__27)($a_prime__28));
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
}, "Functor0" => function($_dollar___unused_23) use ($functorStateT1_22_87) {
  $__num = \func_num_args();
  $__res = $functorStateT1_22_87;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_21_93 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($Monad1_6_5);
  $__res = (object)["apply" => function($f_22) use ($Applicative0_21_93, $Bind1_20_86) {
  $__num = \func_num_args();
  $__res = function($a_23) use ($Applicative0_21_93, $Bind1_20_86, $f_22) {
  $__num = \func_num_args();
  $__res = ((($Bind1_20_86)->{'bind'})($f_22))(function($f_prime__24) use ($Applicative0_21_93, $Bind1_20_86, $a_23) {
  $__num = \func_num_args();
  $__res = ((($Bind1_20_86)->{'bind'})($a_23))(function($a_prime__25) use ($Applicative0_21_93, $f_prime__24) {
  $__num = \func_num_args();
  $__res = (($Applicative0_21_93)->{'pure'})(($f_prime__24)($a_prime__25));
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
}, "Functor0" => function($_dollar___unused_20) use ($functorStateT1_19_84) {
  $__num = \func_num_args();
  $__res = $functorStateT1_19_84;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_18) use ($Applicative0_17_83, $Bind1_16_70) {
  $__num = \func_num_args();
  $__res = function($a_19) use ($Applicative0_17_83, $Bind1_16_70, $f_18) {
  $__num = \func_num_args();
  $__res = ((($Bind1_16_70)->{'bind'})($f_18))(function($f_prime__20) use ($Applicative0_17_83, $Bind1_16_70, $a_19) {
  $__num = \func_num_args();
  $__res = ((($Bind1_16_70)->{'bind'})($a_19))(function($a_prime__21) use ($Applicative0_17_83, $f_prime__20) {
  $__num = \func_num_args();
  $__res = (($Applicative0_17_83)->{'pure'})(($f_prime__20)($a_prime__21));
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
}, "Functor0" => function($_dollar___unused_16) use ($functorStateT1_15_68) {
  $__num = \func_num_args();
  $__res = $functorStateT1_15_68;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_14) use ($Applicative0_13_67, $Bind1_12_54) {
  $__num = \func_num_args();
  $__res = function($a_15) use ($Applicative0_13_67, $Bind1_12_54, $f_14) {
  $__num = \func_num_args();
  $__res = ((($Bind1_12_54)->{'bind'})($f_14))(function($f_prime__16) use ($Applicative0_13_67, $Bind1_12_54, $a_15) {
  $__num = \func_num_args();
  $__res = ((($Bind1_12_54)->{'bind'})($a_15))(function($a_prime__17) use ($Applicative0_13_67, $f_prime__16) {
  $__num = \func_num_args();
  $__res = (($Applicative0_13_67)->{'pure'})(($f_prime__16)($a_prime__17));
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
}, "Functor0" => function($_dollar___unused_12) use ($functorStateT1_11_52) {
  $__num = \func_num_args();
  $__res = $functorStateT1_11_52;
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
}];
  $Bind1_9_97 = (($Monad1_6_5)->{'Bind1'})(null);
  $pure_10_98 = ((($Monad1_6_5)->{'Applicative0'})(null))->{'pure'};
  $monadTellStateT1_6_5 = (object)["tell" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($m_11) use ($Bind1_9_97, $pure_10_98) {
  $__num = \func_num_args();
  $__res = function($s_12) use ($Bind1_9_97, $m_11, $pure_10_98) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_97)->{'bind'})($m_11))(function($x_13) use ($pure_10_98, $s_12) {
  $__num = \func_num_args();
  $__res = ($pure_10_98)(new \Data\Tuple\Data_Tuple_Tuple($x_13, $s_12));
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
}))(($MonadTell1_1_0)->{'tell'}), "Semigroup0" => function($_dollar___unused_9) use ($Semigroup0_7_6) {
  $__num = \func_num_args();
  $__res = $Semigroup0_7_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad1" => function($_dollar___unused_9) use ($monadStateT1_8_7) {
  $__num = \func_num_args();
  $__res = $monadStateT1_8_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["listen" => function($m_7) use ($Applicative0_4_3, $Bind1_3_2, $dictMonadWriter_0) {
  $__num = \func_num_args();
  $__res = function($s_8) use ($Applicative0_4_3, $Bind1_3_2, $dictMonadWriter_0, $m_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_2)->{'bind'})((($dictMonadWriter_0)->{'listen'})(($m_7)($s_8))))(function($v_9) use ($Applicative0_4_3) {
  $__num = \func_num_args();
  $__res = (($Applicative0_4_3)->{'pure'})(new \Data\Tuple\Data_Tuple_Tuple(new \Data\Tuple\Data_Tuple_Tuple((($v_9)->{'value0'})->{'value0'}, ($v_9)->{'value1'}), (($v_9)->{'value0'})->{'value1'}));
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
}, "pass" => function($m_7) use ($Applicative0_4_3, $Bind1_3_2, $dictMonadWriter_0) {
  $__num = \func_num_args();
  $__res = function($s_8) use ($Applicative0_4_3, $Bind1_3_2, $dictMonadWriter_0, $m_7) {
  $__num = \func_num_args();
  $__res = (($dictMonadWriter_0)->{'pass'})(((($Bind1_3_2)->{'bind'})(($m_7)($s_8)))(function($v_9) use ($Applicative0_4_3) {
  $__num = \func_num_args();
  $__res = (($Applicative0_4_3)->{'pure'})(new \Data\Tuple\Data_Tuple_Tuple(new \Data\Tuple\Data_Tuple_Tuple((($v_9)->{'value0'})->{'value0'}, ($v_9)->{'value1'}), (($v_9)->{'value0'})->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monoid0" => function($_dollar___unused_7) use ($Monoid0_5_4) {
  $__num = \func_num_args();
  $__res = $Monoid0_5_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "MonadTell1" => function($_dollar___unused_7) use ($monadTellStateT1_6_5) {
  $__num = \func_num_args();
  $__res = $monadTellStateT1_6_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_State_Trans_monadWriterStateT'] = __NAMESPACE__ . '\\majControl_majMonad_majState_majTrans_monadmajWritermajStatemajT';

// Control_Monad_State_Trans_monadThrowStateT
function majControl_majMonad_majState_majTrans_monadmajThrowmajStatemajT($dictMonadThrow_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majState_majTrans_monadmajThrowmajStatemajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Monad0_1_0 = (($dictMonadThrow_0)->{'Monad0'})(null);
  $__local_var_2_1 = (($dictMonadThrow_0)->{'Monad0'})(null);
  $monadStateT1_2_1 = (object)["Applicative0" => function($_dollar___unused_3) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $pure_4_2 = ((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'};
  $__res = (object)["pure" => function($a_5) use ($pure_4_2) {
  $__num = \func_num_args();
  $__res = function($s_6) use ($a_5, $pure_4_2) {
  $__num = \func_num_args();
  $__res = ($pure_4_2)(new \Data\Tuple\Data_Tuple_Tuple($a_5, $s_6));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_5) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_6_3 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_6_3 = (object)["map" => function($f_7) use ($__local_var_6_3) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($__local_var_6_3, $f_7) {
  $__num = \func_num_args();
  $__res = function($s_9) use ($__local_var_6_3, $f_7, $v_8) {
  $__num = \func_num_args();
  $__res = ((($__local_var_6_3)->{'map'})(function($v1_10) use ($f_7) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_7)(($v1_10)->{'value0'}), ($v1_10)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_8)($s_9));
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
  $Bind1_7_5 = (($__local_var_2_1)->{'Bind1'})(null);
  $Bind1_7_5 = (object)["bind" => function($v_8) use ($Bind1_7_5) {
  $__num = \func_num_args();
  $__res = function($f_9) use ($Bind1_7_5, $v_8) {
  $__num = \func_num_args();
  $__res = function($s_10) use ($Bind1_7_5, $f_9, $v_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_5)->{'bind'})(($v_8)($s_10)))(function($v1_11) use ($f_9) {
  $__num = \func_num_args();
  $__res = (($f_9)(($v1_11)->{'value0'}))(($v1_11)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_8) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_9_6 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_9_6 = (object)["map" => function($f_10) use ($__local_var_9_6) {
  $__num = \func_num_args();
  $__res = function($v_11) use ($__local_var_9_6, $f_10) {
  $__num = \func_num_args();
  $__res = function($s_12) use ($__local_var_9_6, $f_10, $v_11) {
  $__num = \func_num_args();
  $__res = ((($__local_var_9_6)->{'map'})(function($v1_13) use ($f_10) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_10)(($v1_13)->{'value0'}), ($v1_13)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_11)($s_12));
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
  $Bind1_10_8 = (($__local_var_2_1)->{'Bind1'})(null);
  $Bind1_10_8 = (object)["bind" => function($v_11) use ($Bind1_10_8) {
  $__num = \func_num_args();
  $__res = function($f_12) use ($Bind1_10_8, $v_11) {
  $__num = \func_num_args();
  $__res = function($s_13) use ($Bind1_10_8, $f_12, $v_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_8)->{'bind'})(($v_11)($s_13)))(function($v1_14) use ($f_12) {
  $__num = \func_num_args();
  $__res = (($f_12)(($v1_14)->{'value0'}))(($v1_14)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_11) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($__local_var_2_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_11_10 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($__local_var_2_1);
  $__res = (object)["apply" => function($f_12) use ($Applicative0_11_10, $Bind1_10_8) {
  $__num = \func_num_args();
  $__res = function($a_13) use ($Applicative0_11_10, $Bind1_10_8, $f_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_8)->{'bind'})($f_12))(function($f_prime__14) use ($Applicative0_11_10, $Bind1_10_8, $a_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_8)->{'bind'})($a_13))(function($a_prime__15) use ($Applicative0_11_10, $f_prime__14) {
  $__num = \func_num_args();
  $__res = (($Applicative0_11_10)->{'pure'})(($f_prime__14)($a_prime__15));
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
}, "Functor0" => function($_dollar___unused_10) use ($functorStateT1_9_6) {
  $__num = \func_num_args();
  $__res = $functorStateT1_9_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_8_12 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($__local_var_2_1);
  $__res = (object)["apply" => function($f_9) use ($Applicative0_8_12, $Bind1_7_5) {
  $__num = \func_num_args();
  $__res = function($a_10) use ($Applicative0_8_12, $Bind1_7_5, $f_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_5)->{'bind'})($f_9))(function($f_prime__11) use ($Applicative0_8_12, $Bind1_7_5, $a_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_5)->{'bind'})($a_10))(function($a_prime__12) use ($Applicative0_8_12, $f_prime__11) {
  $__num = \func_num_args();
  $__res = (($Applicative0_8_12)->{'pure'})(($f_prime__11)($a_prime__12));
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
}, "Functor0" => function($_dollar___unused_7) use ($functorStateT1_6_3) {
  $__num = \func_num_args();
  $__res = $functorStateT1_6_3;
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
}, "Bind1" => function($_dollar___unused_3) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $Bind1_4_13 = (($__local_var_2_1)->{'Bind1'})(null);
  $__res = (object)["bind" => function($v_5) use ($Bind1_4_13) {
  $__num = \func_num_args();
  $__res = function($f_6) use ($Bind1_4_13, $v_5) {
  $__num = \func_num_args();
  $__res = function($s_7) use ($Bind1_4_13, $f_6, $v_5) {
  $__num = \func_num_args();
  $__res = ((($Bind1_4_13)->{'bind'})(($v_5)($s_7)))(function($v1_8) use ($f_6) {
  $__num = \func_num_args();
  $__res = (($f_6)(($v1_8)->{'value0'}))(($v1_8)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_5) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_6_14 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_6_14 = (object)["map" => function($f_7) use ($__local_var_6_14) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($__local_var_6_14, $f_7) {
  $__num = \func_num_args();
  $__res = function($s_9) use ($__local_var_6_14, $f_7, $v_8) {
  $__num = \func_num_args();
  $__res = ((($__local_var_6_14)->{'map'})(function($v1_10) use ($f_7) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_7)(($v1_10)->{'value0'}), ($v1_10)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_8)($s_9));
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
  $Bind1_7_16 = (($__local_var_2_1)->{'Bind1'})(null);
  $Bind1_7_16 = (object)["bind" => function($v_8) use ($Bind1_7_16) {
  $__num = \func_num_args();
  $__res = function($f_9) use ($Bind1_7_16, $v_8) {
  $__num = \func_num_args();
  $__res = function($s_10) use ($Bind1_7_16, $f_9, $v_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_16)->{'bind'})(($v_8)($s_10)))(function($v1_11) use ($f_9) {
  $__num = \func_num_args();
  $__res = (($f_9)(($v1_11)->{'value0'}))(($v1_11)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_8) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($__local_var_2_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $pure_8_18 = ((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'};
  $Applicative0_8_18 = (object)["pure" => function($a_9) use ($pure_8_18) {
  $__num = \func_num_args();
  $__res = function($s_10) use ($a_9, $pure_8_18) {
  $__num = \func_num_args();
  $__res = ($pure_8_18)(new \Data\Tuple\Data_Tuple_Tuple($a_9, $s_10));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_9) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_10_19 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_10_19 = (object)["map" => function($f_11) use ($__local_var_10_19) {
  $__num = \func_num_args();
  $__res = function($v_12) use ($__local_var_10_19, $f_11) {
  $__num = \func_num_args();
  $__res = function($s_13) use ($__local_var_10_19, $f_11, $v_12) {
  $__num = \func_num_args();
  $__res = ((($__local_var_10_19)->{'map'})(function($v1_14) use ($f_11) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_11)(($v1_14)->{'value0'}), ($v1_14)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_12)($s_13));
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
  $Bind1_11_21 = (($__local_var_2_1)->{'Bind1'})(null);
  $Bind1_11_21 = (object)["bind" => function($v_12) use ($Bind1_11_21) {
  $__num = \func_num_args();
  $__res = function($f_13) use ($Bind1_11_21, $v_12) {
  $__num = \func_num_args();
  $__res = function($s_14) use ($Bind1_11_21, $f_13, $v_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_21)->{'bind'})(($v_12)($s_14)))(function($v1_15) use ($f_13) {
  $__num = \func_num_args();
  $__res = (($f_13)(($v1_15)->{'value0'}))(($v1_15)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_12) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($__local_var_2_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_12_23 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($__local_var_2_1);
  $__res = (object)["apply" => function($f_13) use ($Applicative0_12_23, $Bind1_11_21) {
  $__num = \func_num_args();
  $__res = function($a_14) use ($Applicative0_12_23, $Bind1_11_21, $f_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_21)->{'bind'})($f_13))(function($f_prime__15) use ($Applicative0_12_23, $Bind1_11_21, $a_14) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_21)->{'bind'})($a_14))(function($a_prime__16) use ($Applicative0_12_23, $f_prime__15) {
  $__num = \func_num_args();
  $__res = (($Applicative0_12_23)->{'pure'})(($f_prime__15)($a_prime__16));
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
}, "Functor0" => function($_dollar___unused_11) use ($functorStateT1_10_19) {
  $__num = \func_num_args();
  $__res = $functorStateT1_10_19;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_9) use ($Applicative0_8_18, $Bind1_7_16) {
  $__num = \func_num_args();
  $__res = function($a_10) use ($Applicative0_8_18, $Bind1_7_16, $f_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_16)->{'bind'})($f_9))(function($f_prime__11) use ($Applicative0_8_18, $Bind1_7_16, $a_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_16)->{'bind'})($a_10))(function($a_prime__12) use ($Applicative0_8_18, $f_prime__11) {
  $__num = \func_num_args();
  $__res = (($Applicative0_8_18)->{'pure'})(($f_prime__11)($a_prime__12));
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
}, "Functor0" => function($_dollar___unused_7) use ($functorStateT1_6_14) {
  $__num = \func_num_args();
  $__res = $functorStateT1_6_14;
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
}];
  $__res = (object)["throwError" => function($e_3) use ($Monad0_1_0, $dictMonadThrow_0) {
  $__num = \func_num_args();
  $Bind1_4_26 = (($Monad0_1_0)->{'Bind1'})(null);
  $pure_5_27 = ((($Monad0_1_0)->{'Applicative0'})(null))->{'pure'};
  $__local_var_6_28 = (($dictMonadThrow_0)->{'throwError'})($e_3);
  $__res = function($s_7) use ($Bind1_4_26, $__local_var_6_28, $pure_5_27) {
  $__num = \func_num_args();
  $__res = ((($Bind1_4_26)->{'bind'})($__local_var_6_28))(function($x_8) use ($pure_5_27, $s_7) {
  $__num = \func_num_args();
  $__res = ($pure_5_27)(new \Data\Tuple\Data_Tuple_Tuple($x_8, $s_7));
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
}, "Monad0" => function($_dollar___unused_3) use ($monadStateT1_2_1) {
  $__num = \func_num_args();
  $__res = $monadStateT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_State_Trans_monadThrowStateT'] = __NAMESPACE__ . '\\majControl_majMonad_majState_majTrans_monadmajThrowmajStatemajT';

// Control_Monad_State_Trans_monadErrorStateT
function majControl_majMonad_majState_majTrans_monadmajErrormajStatemajT($dictMonadError_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majState_majTrans_monadmajErrormajStatemajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictMonadError_0)->{'MonadThrow0'})(null);
  $Monad0_2_1 = (($__local_var_1_0)->{'Monad0'})(null);
  $__local_var_3_2 = (($__local_var_1_0)->{'Monad0'})(null);
  $monadStateT1_3_2 = (object)["Applicative0" => function($_dollar___unused_4) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $pure_5_3 = ((($__local_var_3_2)->{'Applicative0'})(null))->{'pure'};
  $__res = (object)["pure" => function($a_6) use ($pure_5_3) {
  $__num = \func_num_args();
  $__res = function($s_7) use ($a_6, $pure_5_3) {
  $__num = \func_num_args();
  $__res = ($pure_5_3)(new \Data\Tuple\Data_Tuple_Tuple($a_6, $s_7));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_6) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__local_var_7_4 = (((((($__local_var_3_2)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_7_4 = (object)["map" => function($f_8) use ($__local_var_7_4) {
  $__num = \func_num_args();
  $__res = function($v_9) use ($__local_var_7_4, $f_8) {
  $__num = \func_num_args();
  $__res = function($s_10) use ($__local_var_7_4, $f_8, $v_9) {
  $__num = \func_num_args();
  $__res = ((($__local_var_7_4)->{'map'})(function($v1_11) use ($f_8) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_8)(($v1_11)->{'value0'}), ($v1_11)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_9)($s_10));
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
  $Bind1_8_6 = (($__local_var_3_2)->{'Bind1'})(null);
  $Bind1_8_6 = (object)["bind" => function($v_9) use ($Bind1_8_6) {
  $__num = \func_num_args();
  $__res = function($f_10) use ($Bind1_8_6, $v_9) {
  $__num = \func_num_args();
  $__res = function($s_11) use ($Bind1_8_6, $f_10, $v_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_8_6)->{'bind'})(($v_9)($s_11)))(function($v1_12) use ($f_10) {
  $__num = \func_num_args();
  $__res = (($f_10)(($v1_12)->{'value0'}))(($v1_12)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_9) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__local_var_10_7 = (((((($__local_var_3_2)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_10_7 = (object)["map" => function($f_11) use ($__local_var_10_7) {
  $__num = \func_num_args();
  $__res = function($v_12) use ($__local_var_10_7, $f_11) {
  $__num = \func_num_args();
  $__res = function($s_13) use ($__local_var_10_7, $f_11, $v_12) {
  $__num = \func_num_args();
  $__res = ((($__local_var_10_7)->{'map'})(function($v1_14) use ($f_11) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_11)(($v1_14)->{'value0'}), ($v1_14)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_12)($s_13));
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
  $Bind1_11_9 = (($__local_var_3_2)->{'Bind1'})(null);
  $Bind1_11_9 = (object)["bind" => function($v_12) use ($Bind1_11_9) {
  $__num = \func_num_args();
  $__res = function($f_13) use ($Bind1_11_9, $v_12) {
  $__num = \func_num_args();
  $__res = function($s_14) use ($Bind1_11_9, $f_13, $v_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_9)->{'bind'})(($v_12)($s_14)))(function($v1_15) use ($f_13) {
  $__num = \func_num_args();
  $__res = (($f_13)(($v1_15)->{'value0'}))(($v1_15)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_12) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__local_var_13_10 = (((((($__local_var_3_2)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_13_10 = (object)["map" => function($f_14) use ($__local_var_13_10) {
  $__num = \func_num_args();
  $__res = function($v_15) use ($__local_var_13_10, $f_14) {
  $__num = \func_num_args();
  $__res = function($s_16) use ($__local_var_13_10, $f_14, $v_15) {
  $__num = \func_num_args();
  $__res = ((($__local_var_13_10)->{'map'})(function($v1_17) use ($f_14) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_14)(($v1_17)->{'value0'}), ($v1_17)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_15)($s_16));
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
  $Bind1_14_12 = (($__local_var_3_2)->{'Bind1'})(null);
  $Bind1_14_12 = (object)["bind" => function($v_15) use ($Bind1_14_12) {
  $__num = \func_num_args();
  $__res = function($f_16) use ($Bind1_14_12, $v_15) {
  $__num = \func_num_args();
  $__res = function($s_17) use ($Bind1_14_12, $f_16, $v_15) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_12)->{'bind'})(($v_15)($s_17)))(function($v1_18) use ($f_16) {
  $__num = \func_num_args();
  $__res = (($f_16)(($v1_18)->{'value0'}))(($v1_18)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_15) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($__local_var_3_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $pure_15_14 = ((($__local_var_3_2)->{'Applicative0'})(null))->{'pure'};
  $Applicative0_15_14 = (object)["pure" => function($a_16) use ($pure_15_14) {
  $__num = \func_num_args();
  $__res = function($s_17) use ($a_16, $pure_15_14) {
  $__num = \func_num_args();
  $__res = ($pure_15_14)(new \Data\Tuple\Data_Tuple_Tuple($a_16, $s_17));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_16) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__local_var_17_15 = (((((($__local_var_3_2)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_17_15 = (object)["map" => function($f_18) use ($__local_var_17_15) {
  $__num = \func_num_args();
  $__res = function($v_19) use ($__local_var_17_15, $f_18) {
  $__num = \func_num_args();
  $__res = function($s_20) use ($__local_var_17_15, $f_18, $v_19) {
  $__num = \func_num_args();
  $__res = ((($__local_var_17_15)->{'map'})(function($v1_21) use ($f_18) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_18)(($v1_21)->{'value0'}), ($v1_21)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_19)($s_20));
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
  $Bind1_18_17 = (($__local_var_3_2)->{'Bind1'})(null);
  $Bind1_18_17 = (object)["bind" => function($v_19) use ($Bind1_18_17) {
  $__num = \func_num_args();
  $__res = function($f_20) use ($Bind1_18_17, $v_19) {
  $__num = \func_num_args();
  $__res = function($s_21) use ($Bind1_18_17, $f_20, $v_19) {
  $__num = \func_num_args();
  $__res = ((($Bind1_18_17)->{'bind'})(($v_19)($s_21)))(function($v1_22) use ($f_20) {
  $__num = \func_num_args();
  $__res = (($f_20)(($v1_22)->{'value0'}))(($v1_22)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_19) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($__local_var_3_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_19_19 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($__local_var_3_2);
  $__res = (object)["apply" => function($f_20) use ($Applicative0_19_19, $Bind1_18_17) {
  $__num = \func_num_args();
  $__res = function($a_21) use ($Applicative0_19_19, $Bind1_18_17, $f_20) {
  $__num = \func_num_args();
  $__res = ((($Bind1_18_17)->{'bind'})($f_20))(function($f_prime__22) use ($Applicative0_19_19, $Bind1_18_17, $a_21) {
  $__num = \func_num_args();
  $__res = ((($Bind1_18_17)->{'bind'})($a_21))(function($a_prime__23) use ($Applicative0_19_19, $f_prime__22) {
  $__num = \func_num_args();
  $__res = (($Applicative0_19_19)->{'pure'})(($f_prime__22)($a_prime__23));
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
}, "Functor0" => function($_dollar___unused_18) use ($functorStateT1_17_15) {
  $__num = \func_num_args();
  $__res = $functorStateT1_17_15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_16) use ($Applicative0_15_14, $Bind1_14_12) {
  $__num = \func_num_args();
  $__res = function($a_17) use ($Applicative0_15_14, $Bind1_14_12, $f_16) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_12)->{'bind'})($f_16))(function($f_prime__18) use ($Applicative0_15_14, $Bind1_14_12, $a_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_12)->{'bind'})($a_17))(function($a_prime__19) use ($Applicative0_15_14, $f_prime__18) {
  $__num = \func_num_args();
  $__res = (($Applicative0_15_14)->{'pure'})(($f_prime__18)($a_prime__19));
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
}, "Functor0" => function($_dollar___unused_14) use ($functorStateT1_13_10) {
  $__num = \func_num_args();
  $__res = $functorStateT1_13_10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $pure_12_22 = ((($__local_var_3_2)->{'Applicative0'})(null))->{'pure'};
  $Applicative0_12_22 = (object)["pure" => function($a_13) use ($pure_12_22) {
  $__num = \func_num_args();
  $__res = function($s_14) use ($a_13, $pure_12_22) {
  $__num = \func_num_args();
  $__res = ($pure_12_22)(new \Data\Tuple\Data_Tuple_Tuple($a_13, $s_14));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_13) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__local_var_14_23 = (((((($__local_var_3_2)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_14_23 = (object)["map" => function($f_15) use ($__local_var_14_23) {
  $__num = \func_num_args();
  $__res = function($v_16) use ($__local_var_14_23, $f_15) {
  $__num = \func_num_args();
  $__res = function($s_17) use ($__local_var_14_23, $f_15, $v_16) {
  $__num = \func_num_args();
  $__res = ((($__local_var_14_23)->{'map'})(function($v1_18) use ($f_15) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_15)(($v1_18)->{'value0'}), ($v1_18)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_16)($s_17));
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
  $Bind1_15_25 = (($__local_var_3_2)->{'Bind1'})(null);
  $Bind1_15_25 = (object)["bind" => function($v_16) use ($Bind1_15_25) {
  $__num = \func_num_args();
  $__res = function($f_17) use ($Bind1_15_25, $v_16) {
  $__num = \func_num_args();
  $__res = function($s_18) use ($Bind1_15_25, $f_17, $v_16) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_25)->{'bind'})(($v_16)($s_18)))(function($v1_19) use ($f_17) {
  $__num = \func_num_args();
  $__res = (($f_17)(($v1_19)->{'value0'}))(($v1_19)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_16) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__local_var_17_26 = (((((($__local_var_3_2)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_17_26 = (object)["map" => function($f_18) use ($__local_var_17_26) {
  $__num = \func_num_args();
  $__res = function($v_19) use ($__local_var_17_26, $f_18) {
  $__num = \func_num_args();
  $__res = function($s_20) use ($__local_var_17_26, $f_18, $v_19) {
  $__num = \func_num_args();
  $__res = ((($__local_var_17_26)->{'map'})(function($v1_21) use ($f_18) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_18)(($v1_21)->{'value0'}), ($v1_21)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_19)($s_20));
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
  $Bind1_18_28 = (($__local_var_3_2)->{'Bind1'})(null);
  $Bind1_18_28 = (object)["bind" => function($v_19) use ($Bind1_18_28) {
  $__num = \func_num_args();
  $__res = function($f_20) use ($Bind1_18_28, $v_19) {
  $__num = \func_num_args();
  $__res = function($s_21) use ($Bind1_18_28, $f_20, $v_19) {
  $__num = \func_num_args();
  $__res = ((($Bind1_18_28)->{'bind'})(($v_19)($s_21)))(function($v1_22) use ($f_20) {
  $__num = \func_num_args();
  $__res = (($f_20)(($v1_22)->{'value0'}))(($v1_22)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_19) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($__local_var_3_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_19_30 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($__local_var_3_2);
  $__res = (object)["apply" => function($f_20) use ($Applicative0_19_30, $Bind1_18_28) {
  $__num = \func_num_args();
  $__res = function($a_21) use ($Applicative0_19_30, $Bind1_18_28, $f_20) {
  $__num = \func_num_args();
  $__res = ((($Bind1_18_28)->{'bind'})($f_20))(function($f_prime__22) use ($Applicative0_19_30, $Bind1_18_28, $a_21) {
  $__num = \func_num_args();
  $__res = ((($Bind1_18_28)->{'bind'})($a_21))(function($a_prime__23) use ($Applicative0_19_30, $f_prime__22) {
  $__num = \func_num_args();
  $__res = (($Applicative0_19_30)->{'pure'})(($f_prime__22)($a_prime__23));
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
}, "Functor0" => function($_dollar___unused_18) use ($functorStateT1_17_26) {
  $__num = \func_num_args();
  $__res = $functorStateT1_17_26;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_16_32 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($__local_var_3_2);
  $__res = (object)["apply" => function($f_17) use ($Applicative0_16_32, $Bind1_15_25) {
  $__num = \func_num_args();
  $__res = function($a_18) use ($Applicative0_16_32, $Bind1_15_25, $f_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_25)->{'bind'})($f_17))(function($f_prime__19) use ($Applicative0_16_32, $Bind1_15_25, $a_18) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_25)->{'bind'})($a_18))(function($a_prime__20) use ($Applicative0_16_32, $f_prime__19) {
  $__num = \func_num_args();
  $__res = (($Applicative0_16_32)->{'pure'})(($f_prime__19)($a_prime__20));
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
}, "Functor0" => function($_dollar___unused_15) use ($functorStateT1_14_23) {
  $__num = \func_num_args();
  $__res = $functorStateT1_14_23;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_13) use ($Applicative0_12_22, $Bind1_11_9) {
  $__num = \func_num_args();
  $__res = function($a_14) use ($Applicative0_12_22, $Bind1_11_9, $f_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_9)->{'bind'})($f_13))(function($f_prime__15) use ($Applicative0_12_22, $Bind1_11_9, $a_14) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_9)->{'bind'})($a_14))(function($a_prime__16) use ($Applicative0_12_22, $f_prime__15) {
  $__num = \func_num_args();
  $__res = (($Applicative0_12_22)->{'pure'})(($f_prime__15)($a_prime__16));
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
}, "Functor0" => function($_dollar___unused_11) use ($functorStateT1_10_7) {
  $__num = \func_num_args();
  $__res = $functorStateT1_10_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $pure_9_35 = ((($__local_var_3_2)->{'Applicative0'})(null))->{'pure'};
  $Applicative0_9_35 = (object)["pure" => function($a_10) use ($pure_9_35) {
  $__num = \func_num_args();
  $__res = function($s_11) use ($a_10, $pure_9_35) {
  $__num = \func_num_args();
  $__res = ($pure_9_35)(new \Data\Tuple\Data_Tuple_Tuple($a_10, $s_11));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_10) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__local_var_11_36 = (((((($__local_var_3_2)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_11_36 = (object)["map" => function($f_12) use ($__local_var_11_36) {
  $__num = \func_num_args();
  $__res = function($v_13) use ($__local_var_11_36, $f_12) {
  $__num = \func_num_args();
  $__res = function($s_14) use ($__local_var_11_36, $f_12, $v_13) {
  $__num = \func_num_args();
  $__res = ((($__local_var_11_36)->{'map'})(function($v1_15) use ($f_12) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_12)(($v1_15)->{'value0'}), ($v1_15)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_13)($s_14));
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
  $Bind1_12_38 = (($__local_var_3_2)->{'Bind1'})(null);
  $Bind1_12_38 = (object)["bind" => function($v_13) use ($Bind1_12_38) {
  $__num = \func_num_args();
  $__res = function($f_14) use ($Bind1_12_38, $v_13) {
  $__num = \func_num_args();
  $__res = function($s_15) use ($Bind1_12_38, $f_14, $v_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_12_38)->{'bind'})(($v_13)($s_15)))(function($v1_16) use ($f_14) {
  $__num = \func_num_args();
  $__res = (($f_14)(($v1_16)->{'value0'}))(($v1_16)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_13) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__local_var_14_39 = (((((($__local_var_3_2)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_14_39 = (object)["map" => function($f_15) use ($__local_var_14_39) {
  $__num = \func_num_args();
  $__res = function($v_16) use ($__local_var_14_39, $f_15) {
  $__num = \func_num_args();
  $__res = function($s_17) use ($__local_var_14_39, $f_15, $v_16) {
  $__num = \func_num_args();
  $__res = ((($__local_var_14_39)->{'map'})(function($v1_18) use ($f_15) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_15)(($v1_18)->{'value0'}), ($v1_18)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_16)($s_17));
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
  $Bind1_15_41 = (($__local_var_3_2)->{'Bind1'})(null);
  $Bind1_15_41 = (object)["bind" => function($v_16) use ($Bind1_15_41) {
  $__num = \func_num_args();
  $__res = function($f_17) use ($Bind1_15_41, $v_16) {
  $__num = \func_num_args();
  $__res = function($s_18) use ($Bind1_15_41, $f_17, $v_16) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_41)->{'bind'})(($v_16)($s_18)))(function($v1_19) use ($f_17) {
  $__num = \func_num_args();
  $__res = (($f_17)(($v1_19)->{'value0'}))(($v1_19)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_16) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($__local_var_3_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_16_43 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($__local_var_3_2);
  $__res = (object)["apply" => function($f_17) use ($Applicative0_16_43, $Bind1_15_41) {
  $__num = \func_num_args();
  $__res = function($a_18) use ($Applicative0_16_43, $Bind1_15_41, $f_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_41)->{'bind'})($f_17))(function($f_prime__19) use ($Applicative0_16_43, $Bind1_15_41, $a_18) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_41)->{'bind'})($a_18))(function($a_prime__20) use ($Applicative0_16_43, $f_prime__19) {
  $__num = \func_num_args();
  $__res = (($Applicative0_16_43)->{'pure'})(($f_prime__19)($a_prime__20));
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
}, "Functor0" => function($_dollar___unused_15) use ($functorStateT1_14_39) {
  $__num = \func_num_args();
  $__res = $functorStateT1_14_39;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_13_45 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($__local_var_3_2);
  $__res = (object)["apply" => function($f_14) use ($Applicative0_13_45, $Bind1_12_38) {
  $__num = \func_num_args();
  $__res = function($a_15) use ($Applicative0_13_45, $Bind1_12_38, $f_14) {
  $__num = \func_num_args();
  $__res = ((($Bind1_12_38)->{'bind'})($f_14))(function($f_prime__16) use ($Applicative0_13_45, $Bind1_12_38, $a_15) {
  $__num = \func_num_args();
  $__res = ((($Bind1_12_38)->{'bind'})($a_15))(function($a_prime__17) use ($Applicative0_13_45, $f_prime__16) {
  $__num = \func_num_args();
  $__res = (($Applicative0_13_45)->{'pure'})(($f_prime__16)($a_prime__17));
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
}, "Functor0" => function($_dollar___unused_12) use ($functorStateT1_11_36) {
  $__num = \func_num_args();
  $__res = $functorStateT1_11_36;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_10) use ($Applicative0_9_35, $Bind1_8_6) {
  $__num = \func_num_args();
  $__res = function($a_11) use ($Applicative0_9_35, $Bind1_8_6, $f_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_8_6)->{'bind'})($f_10))(function($f_prime__12) use ($Applicative0_9_35, $Bind1_8_6, $a_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_8_6)->{'bind'})($a_11))(function($a_prime__13) use ($Applicative0_9_35, $f_prime__12) {
  $__num = \func_num_args();
  $__res = (($Applicative0_9_35)->{'pure'})(($f_prime__12)($a_prime__13));
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
}, "Functor0" => function($_dollar___unused_8) use ($functorStateT1_7_4) {
  $__num = \func_num_args();
  $__res = $functorStateT1_7_4;
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
}, "Bind1" => function($_dollar___unused_4) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $Bind1_5_47 = (($__local_var_3_2)->{'Bind1'})(null);
  $__res = (object)["bind" => function($v_6) use ($Bind1_5_47) {
  $__num = \func_num_args();
  $__res = function($f_7) use ($Bind1_5_47, $v_6) {
  $__num = \func_num_args();
  $__res = function($s_8) use ($Bind1_5_47, $f_7, $v_6) {
  $__num = \func_num_args();
  $__res = ((($Bind1_5_47)->{'bind'})(($v_6)($s_8)))(function($v1_9) use ($f_7) {
  $__num = \func_num_args();
  $__res = (($f_7)(($v1_9)->{'value0'}))(($v1_9)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_6) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__local_var_7_48 = (((((($__local_var_3_2)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_7_48 = (object)["map" => function($f_8) use ($__local_var_7_48) {
  $__num = \func_num_args();
  $__res = function($v_9) use ($__local_var_7_48, $f_8) {
  $__num = \func_num_args();
  $__res = function($s_10) use ($__local_var_7_48, $f_8, $v_9) {
  $__num = \func_num_args();
  $__res = ((($__local_var_7_48)->{'map'})(function($v1_11) use ($f_8) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_8)(($v1_11)->{'value0'}), ($v1_11)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_9)($s_10));
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
  $Bind1_8_50 = (($__local_var_3_2)->{'Bind1'})(null);
  $Bind1_8_50 = (object)["bind" => function($v_9) use ($Bind1_8_50) {
  $__num = \func_num_args();
  $__res = function($f_10) use ($Bind1_8_50, $v_9) {
  $__num = \func_num_args();
  $__res = function($s_11) use ($Bind1_8_50, $f_10, $v_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_8_50)->{'bind'})(($v_9)($s_11)))(function($v1_12) use ($f_10) {
  $__num = \func_num_args();
  $__res = (($f_10)(($v1_12)->{'value0'}))(($v1_12)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_9) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__local_var_10_51 = (((((($__local_var_3_2)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_10_51 = (object)["map" => function($f_11) use ($__local_var_10_51) {
  $__num = \func_num_args();
  $__res = function($v_12) use ($__local_var_10_51, $f_11) {
  $__num = \func_num_args();
  $__res = function($s_13) use ($__local_var_10_51, $f_11, $v_12) {
  $__num = \func_num_args();
  $__res = ((($__local_var_10_51)->{'map'})(function($v1_14) use ($f_11) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_11)(($v1_14)->{'value0'}), ($v1_14)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_12)($s_13));
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
  $Bind1_11_53 = (($__local_var_3_2)->{'Bind1'})(null);
  $Bind1_11_53 = (object)["bind" => function($v_12) use ($Bind1_11_53) {
  $__num = \func_num_args();
  $__res = function($f_13) use ($Bind1_11_53, $v_12) {
  $__num = \func_num_args();
  $__res = function($s_14) use ($Bind1_11_53, $f_13, $v_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_53)->{'bind'})(($v_12)($s_14)))(function($v1_15) use ($f_13) {
  $__num = \func_num_args();
  $__res = (($f_13)(($v1_15)->{'value0'}))(($v1_15)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_12) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($__local_var_3_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $pure_12_55 = ((($__local_var_3_2)->{'Applicative0'})(null))->{'pure'};
  $Applicative0_12_55 = (object)["pure" => function($a_13) use ($pure_12_55) {
  $__num = \func_num_args();
  $__res = function($s_14) use ($a_13, $pure_12_55) {
  $__num = \func_num_args();
  $__res = ($pure_12_55)(new \Data\Tuple\Data_Tuple_Tuple($a_13, $s_14));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_13) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__local_var_14_56 = (((((($__local_var_3_2)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_14_56 = (object)["map" => function($f_15) use ($__local_var_14_56) {
  $__num = \func_num_args();
  $__res = function($v_16) use ($__local_var_14_56, $f_15) {
  $__num = \func_num_args();
  $__res = function($s_17) use ($__local_var_14_56, $f_15, $v_16) {
  $__num = \func_num_args();
  $__res = ((($__local_var_14_56)->{'map'})(function($v1_18) use ($f_15) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_15)(($v1_18)->{'value0'}), ($v1_18)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_16)($s_17));
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
  $Bind1_15_58 = (($__local_var_3_2)->{'Bind1'})(null);
  $Bind1_15_58 = (object)["bind" => function($v_16) use ($Bind1_15_58) {
  $__num = \func_num_args();
  $__res = function($f_17) use ($Bind1_15_58, $v_16) {
  $__num = \func_num_args();
  $__res = function($s_18) use ($Bind1_15_58, $f_17, $v_16) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_58)->{'bind'})(($v_16)($s_18)))(function($v1_19) use ($f_17) {
  $__num = \func_num_args();
  $__res = (($f_17)(($v1_19)->{'value0'}))(($v1_19)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_16) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($__local_var_3_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_16_60 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($__local_var_3_2);
  $__res = (object)["apply" => function($f_17) use ($Applicative0_16_60, $Bind1_15_58) {
  $__num = \func_num_args();
  $__res = function($a_18) use ($Applicative0_16_60, $Bind1_15_58, $f_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_58)->{'bind'})($f_17))(function($f_prime__19) use ($Applicative0_16_60, $Bind1_15_58, $a_18) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_58)->{'bind'})($a_18))(function($a_prime__20) use ($Applicative0_16_60, $f_prime__19) {
  $__num = \func_num_args();
  $__res = (($Applicative0_16_60)->{'pure'})(($f_prime__19)($a_prime__20));
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
}, "Functor0" => function($_dollar___unused_15) use ($functorStateT1_14_56) {
  $__num = \func_num_args();
  $__res = $functorStateT1_14_56;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_13) use ($Applicative0_12_55, $Bind1_11_53) {
  $__num = \func_num_args();
  $__res = function($a_14) use ($Applicative0_12_55, $Bind1_11_53, $f_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_53)->{'bind'})($f_13))(function($f_prime__15) use ($Applicative0_12_55, $Bind1_11_53, $a_14) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_53)->{'bind'})($a_14))(function($a_prime__16) use ($Applicative0_12_55, $f_prime__15) {
  $__num = \func_num_args();
  $__res = (($Applicative0_12_55)->{'pure'})(($f_prime__15)($a_prime__16));
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
}, "Functor0" => function($_dollar___unused_11) use ($functorStateT1_10_51) {
  $__num = \func_num_args();
  $__res = $functorStateT1_10_51;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $pure_9_63 = ((($__local_var_3_2)->{'Applicative0'})(null))->{'pure'};
  $Applicative0_9_63 = (object)["pure" => function($a_10) use ($pure_9_63) {
  $__num = \func_num_args();
  $__res = function($s_11) use ($a_10, $pure_9_63) {
  $__num = \func_num_args();
  $__res = ($pure_9_63)(new \Data\Tuple\Data_Tuple_Tuple($a_10, $s_11));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_10) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__local_var_11_64 = (((((($__local_var_3_2)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_11_64 = (object)["map" => function($f_12) use ($__local_var_11_64) {
  $__num = \func_num_args();
  $__res = function($v_13) use ($__local_var_11_64, $f_12) {
  $__num = \func_num_args();
  $__res = function($s_14) use ($__local_var_11_64, $f_12, $v_13) {
  $__num = \func_num_args();
  $__res = ((($__local_var_11_64)->{'map'})(function($v1_15) use ($f_12) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_12)(($v1_15)->{'value0'}), ($v1_15)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_13)($s_14));
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
  $Bind1_12_66 = (($__local_var_3_2)->{'Bind1'})(null);
  $Bind1_12_66 = (object)["bind" => function($v_13) use ($Bind1_12_66) {
  $__num = \func_num_args();
  $__res = function($f_14) use ($Bind1_12_66, $v_13) {
  $__num = \func_num_args();
  $__res = function($s_15) use ($Bind1_12_66, $f_14, $v_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_12_66)->{'bind'})(($v_13)($s_15)))(function($v1_16) use ($f_14) {
  $__num = \func_num_args();
  $__res = (($f_14)(($v1_16)->{'value0'}))(($v1_16)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_13) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__local_var_14_67 = (((((($__local_var_3_2)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_14_67 = (object)["map" => function($f_15) use ($__local_var_14_67) {
  $__num = \func_num_args();
  $__res = function($v_16) use ($__local_var_14_67, $f_15) {
  $__num = \func_num_args();
  $__res = function($s_17) use ($__local_var_14_67, $f_15, $v_16) {
  $__num = \func_num_args();
  $__res = ((($__local_var_14_67)->{'map'})(function($v1_18) use ($f_15) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_15)(($v1_18)->{'value0'}), ($v1_18)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_16)($s_17));
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
  $Bind1_15_69 = (($__local_var_3_2)->{'Bind1'})(null);
  $Bind1_15_69 = (object)["bind" => function($v_16) use ($Bind1_15_69) {
  $__num = \func_num_args();
  $__res = function($f_17) use ($Bind1_15_69, $v_16) {
  $__num = \func_num_args();
  $__res = function($s_18) use ($Bind1_15_69, $f_17, $v_16) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_69)->{'bind'})(($v_16)($s_18)))(function($v1_19) use ($f_17) {
  $__num = \func_num_args();
  $__res = (($f_17)(($v1_19)->{'value0'}))(($v1_19)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_16) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($__local_var_3_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $pure_16_71 = ((($__local_var_3_2)->{'Applicative0'})(null))->{'pure'};
  $Applicative0_16_71 = (object)["pure" => function($a_17) use ($pure_16_71) {
  $__num = \func_num_args();
  $__res = function($s_18) use ($a_17, $pure_16_71) {
  $__num = \func_num_args();
  $__res = ($pure_16_71)(new \Data\Tuple\Data_Tuple_Tuple($a_17, $s_18));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_17) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__local_var_18_72 = (((((($__local_var_3_2)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_18_72 = (object)["map" => function($f_19) use ($__local_var_18_72) {
  $__num = \func_num_args();
  $__res = function($v_20) use ($__local_var_18_72, $f_19) {
  $__num = \func_num_args();
  $__res = function($s_21) use ($__local_var_18_72, $f_19, $v_20) {
  $__num = \func_num_args();
  $__res = ((($__local_var_18_72)->{'map'})(function($v1_22) use ($f_19) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_19)(($v1_22)->{'value0'}), ($v1_22)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_20)($s_21));
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
  $Bind1_19_74 = (($__local_var_3_2)->{'Bind1'})(null);
  $Bind1_19_74 = (object)["bind" => function($v_20) use ($Bind1_19_74) {
  $__num = \func_num_args();
  $__res = function($f_21) use ($Bind1_19_74, $v_20) {
  $__num = \func_num_args();
  $__res = function($s_22) use ($Bind1_19_74, $f_21, $v_20) {
  $__num = \func_num_args();
  $__res = ((($Bind1_19_74)->{'bind'})(($v_20)($s_22)))(function($v1_23) use ($f_21) {
  $__num = \func_num_args();
  $__res = (($f_21)(($v1_23)->{'value0'}))(($v1_23)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_20) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($__local_var_3_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_20_76 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($__local_var_3_2);
  $__res = (object)["apply" => function($f_21) use ($Applicative0_20_76, $Bind1_19_74) {
  $__num = \func_num_args();
  $__res = function($a_22) use ($Applicative0_20_76, $Bind1_19_74, $f_21) {
  $__num = \func_num_args();
  $__res = ((($Bind1_19_74)->{'bind'})($f_21))(function($f_prime__23) use ($Applicative0_20_76, $Bind1_19_74, $a_22) {
  $__num = \func_num_args();
  $__res = ((($Bind1_19_74)->{'bind'})($a_22))(function($a_prime__24) use ($Applicative0_20_76, $f_prime__23) {
  $__num = \func_num_args();
  $__res = (($Applicative0_20_76)->{'pure'})(($f_prime__23)($a_prime__24));
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
}, "Functor0" => function($_dollar___unused_19) use ($functorStateT1_18_72) {
  $__num = \func_num_args();
  $__res = $functorStateT1_18_72;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_17) use ($Applicative0_16_71, $Bind1_15_69) {
  $__num = \func_num_args();
  $__res = function($a_18) use ($Applicative0_16_71, $Bind1_15_69, $f_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_69)->{'bind'})($f_17))(function($f_prime__19) use ($Applicative0_16_71, $Bind1_15_69, $a_18) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_69)->{'bind'})($a_18))(function($a_prime__20) use ($Applicative0_16_71, $f_prime__19) {
  $__num = \func_num_args();
  $__res = (($Applicative0_16_71)->{'pure'})(($f_prime__19)($a_prime__20));
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
}, "Functor0" => function($_dollar___unused_15) use ($functorStateT1_14_67) {
  $__num = \func_num_args();
  $__res = $functorStateT1_14_67;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $pure_13_79 = ((($__local_var_3_2)->{'Applicative0'})(null))->{'pure'};
  $Applicative0_13_79 = (object)["pure" => function($a_14) use ($pure_13_79) {
  $__num = \func_num_args();
  $__res = function($s_15) use ($a_14, $pure_13_79) {
  $__num = \func_num_args();
  $__res = ($pure_13_79)(new \Data\Tuple\Data_Tuple_Tuple($a_14, $s_15));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_14) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__local_var_15_80 = (((((($__local_var_3_2)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_15_80 = (object)["map" => function($f_16) use ($__local_var_15_80) {
  $__num = \func_num_args();
  $__res = function($v_17) use ($__local_var_15_80, $f_16) {
  $__num = \func_num_args();
  $__res = function($s_18) use ($__local_var_15_80, $f_16, $v_17) {
  $__num = \func_num_args();
  $__res = ((($__local_var_15_80)->{'map'})(function($v1_19) use ($f_16) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_16)(($v1_19)->{'value0'}), ($v1_19)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_17)($s_18));
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
  $Bind1_16_82 = (($__local_var_3_2)->{'Bind1'})(null);
  $Bind1_16_82 = (object)["bind" => function($v_17) use ($Bind1_16_82) {
  $__num = \func_num_args();
  $__res = function($f_18) use ($Bind1_16_82, $v_17) {
  $__num = \func_num_args();
  $__res = function($s_19) use ($Bind1_16_82, $f_18, $v_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_16_82)->{'bind'})(($v_17)($s_19)))(function($v1_20) use ($f_18) {
  $__num = \func_num_args();
  $__res = (($f_18)(($v1_20)->{'value0'}))(($v1_20)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_17) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__local_var_18_83 = (((((($__local_var_3_2)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_18_83 = (object)["map" => function($f_19) use ($__local_var_18_83) {
  $__num = \func_num_args();
  $__res = function($v_20) use ($__local_var_18_83, $f_19) {
  $__num = \func_num_args();
  $__res = function($s_21) use ($__local_var_18_83, $f_19, $v_20) {
  $__num = \func_num_args();
  $__res = ((($__local_var_18_83)->{'map'})(function($v1_22) use ($f_19) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_19)(($v1_22)->{'value0'}), ($v1_22)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_20)($s_21));
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
  $Bind1_19_85 = (($__local_var_3_2)->{'Bind1'})(null);
  $Bind1_19_85 = (object)["bind" => function($v_20) use ($Bind1_19_85) {
  $__num = \func_num_args();
  $__res = function($f_21) use ($Bind1_19_85, $v_20) {
  $__num = \func_num_args();
  $__res = function($s_22) use ($Bind1_19_85, $f_21, $v_20) {
  $__num = \func_num_args();
  $__res = ((($Bind1_19_85)->{'bind'})(($v_20)($s_22)))(function($v1_23) use ($f_21) {
  $__num = \func_num_args();
  $__res = (($f_21)(($v1_23)->{'value0'}))(($v1_23)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_20) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($__local_var_3_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_20_87 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($__local_var_3_2);
  $__res = (object)["apply" => function($f_21) use ($Applicative0_20_87, $Bind1_19_85) {
  $__num = \func_num_args();
  $__res = function($a_22) use ($Applicative0_20_87, $Bind1_19_85, $f_21) {
  $__num = \func_num_args();
  $__res = ((($Bind1_19_85)->{'bind'})($f_21))(function($f_prime__23) use ($Applicative0_20_87, $Bind1_19_85, $a_22) {
  $__num = \func_num_args();
  $__res = ((($Bind1_19_85)->{'bind'})($a_22))(function($a_prime__24) use ($Applicative0_20_87, $f_prime__23) {
  $__num = \func_num_args();
  $__res = (($Applicative0_20_87)->{'pure'})(($f_prime__23)($a_prime__24));
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
}, "Functor0" => function($_dollar___unused_19) use ($functorStateT1_18_83) {
  $__num = \func_num_args();
  $__res = $functorStateT1_18_83;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_17_89 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($__local_var_3_2);
  $__res = (object)["apply" => function($f_18) use ($Applicative0_17_89, $Bind1_16_82) {
  $__num = \func_num_args();
  $__res = function($a_19) use ($Applicative0_17_89, $Bind1_16_82, $f_18) {
  $__num = \func_num_args();
  $__res = ((($Bind1_16_82)->{'bind'})($f_18))(function($f_prime__20) use ($Applicative0_17_89, $Bind1_16_82, $a_19) {
  $__num = \func_num_args();
  $__res = ((($Bind1_16_82)->{'bind'})($a_19))(function($a_prime__21) use ($Applicative0_17_89, $f_prime__20) {
  $__num = \func_num_args();
  $__res = (($Applicative0_17_89)->{'pure'})(($f_prime__20)($a_prime__21));
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
}, "Functor0" => function($_dollar___unused_16) use ($functorStateT1_15_80) {
  $__num = \func_num_args();
  $__res = $functorStateT1_15_80;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_14) use ($Applicative0_13_79, $Bind1_12_66) {
  $__num = \func_num_args();
  $__res = function($a_15) use ($Applicative0_13_79, $Bind1_12_66, $f_14) {
  $__num = \func_num_args();
  $__res = ((($Bind1_12_66)->{'bind'})($f_14))(function($f_prime__16) use ($Applicative0_13_79, $Bind1_12_66, $a_15) {
  $__num = \func_num_args();
  $__res = ((($Bind1_12_66)->{'bind'})($a_15))(function($a_prime__17) use ($Applicative0_13_79, $f_prime__16) {
  $__num = \func_num_args();
  $__res = (($Applicative0_13_79)->{'pure'})(($f_prime__16)($a_prime__17));
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
}, "Functor0" => function($_dollar___unused_12) use ($functorStateT1_11_64) {
  $__num = \func_num_args();
  $__res = $functorStateT1_11_64;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_10) use ($Applicative0_9_63, $Bind1_8_50) {
  $__num = \func_num_args();
  $__res = function($a_11) use ($Applicative0_9_63, $Bind1_8_50, $f_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_8_50)->{'bind'})($f_10))(function($f_prime__12) use ($Applicative0_9_63, $Bind1_8_50, $a_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_8_50)->{'bind'})($a_11))(function($a_prime__13) use ($Applicative0_9_63, $f_prime__12) {
  $__num = \func_num_args();
  $__res = (($Applicative0_9_63)->{'pure'})(($f_prime__12)($a_prime__13));
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
}, "Functor0" => function($_dollar___unused_8) use ($functorStateT1_7_48) {
  $__num = \func_num_args();
  $__res = $functorStateT1_7_48;
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
}];
  $monadThrowStateT1_1_0 = (object)["throwError" => function($e_4) use ($Monad0_2_1, $__local_var_1_0) {
  $__num = \func_num_args();
  $Bind1_5_93 = (($Monad0_2_1)->{'Bind1'})(null);
  $pure_6_94 = ((($Monad0_2_1)->{'Applicative0'})(null))->{'pure'};
  $__local_var_7_95 = (($__local_var_1_0)->{'throwError'})($e_4);
  $__res = function($s_8) use ($Bind1_5_93, $__local_var_7_95, $pure_6_94) {
  $__num = \func_num_args();
  $__res = ((($Bind1_5_93)->{'bind'})($__local_var_7_95))(function($x_9) use ($pure_6_94, $s_8) {
  $__num = \func_num_args();
  $__res = ($pure_6_94)(new \Data\Tuple\Data_Tuple_Tuple($x_9, $s_8));
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
}, "Monad0" => function($_dollar___unused_4) use ($monadStateT1_3_2) {
  $__num = \func_num_args();
  $__res = $monadStateT1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["catchError" => function($v_2) use ($dictMonadError_0) {
  $__num = \func_num_args();
  $__res = function($h_3) use ($dictMonadError_0, $v_2) {
  $__num = \func_num_args();
  $__res = function($s_4) use ($dictMonadError_0, $h_3, $v_2) {
  $__num = \func_num_args();
  $__res = ((($dictMonadError_0)->{'catchError'})(($v_2)($s_4)))(function($e_5) use ($h_3, $s_4) {
  $__num = \func_num_args();
  $__res = (($h_3)($e_5))($s_4);
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
}, "MonadThrow0" => function($_dollar___unused_2) use ($monadThrowStateT1_1_0) {
  $__num = \func_num_args();
  $__res = $monadThrowStateT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_State_Trans_monadErrorStateT'] = __NAMESPACE__ . '\\majControl_majMonad_majState_majTrans_monadmajErrormajStatemajT';

// Control_Monad_State_Trans_monadSTStateT
function majControl_majMonad_majState_majTrans_monadmajSmajTmajStatemajT($dictMonadST_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majState_majTrans_monadmajSmajTmajStatemajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Monad0_1_0 = (($dictMonadST_0)->{'Monad0'})(null);
  $monadStateT1_2_1 = (object)["Applicative0" => function($_dollar___unused_2) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $pure_3_1 = ((($Monad0_1_0)->{'Applicative0'})(null))->{'pure'};
  $__res = (object)["pure" => function($a_4) use ($pure_3_1) {
  $__num = \func_num_args();
  $__res = function($s_5) use ($a_4, $pure_3_1) {
  $__num = \func_num_args();
  $__res = ($pure_3_1)(new \Data\Tuple\Data_Tuple_Tuple($a_4, $s_5));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_4) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__local_var_5_2 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_5_2 = (object)["map" => function($f_6) use ($__local_var_5_2) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($__local_var_5_2, $f_6) {
  $__num = \func_num_args();
  $__res = function($s_8) use ($__local_var_5_2, $f_6, $v_7) {
  $__num = \func_num_args();
  $__res = ((($__local_var_5_2)->{'map'})(function($v1_9) use ($f_6) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_6)(($v1_9)->{'value0'}), ($v1_9)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_7)($s_8));
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
  $Bind1_6_4 = (($Monad0_1_0)->{'Bind1'})(null);
  $Bind1_6_4 = (object)["bind" => function($v_7) use ($Bind1_6_4) {
  $__num = \func_num_args();
  $__res = function($f_8) use ($Bind1_6_4, $v_7) {
  $__num = \func_num_args();
  $__res = function($s_9) use ($Bind1_6_4, $f_8, $v_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_4)->{'bind'})(($v_7)($s_9)))(function($v1_10) use ($f_8) {
  $__num = \func_num_args();
  $__res = (($f_8)(($v1_10)->{'value0'}))(($v1_10)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_7) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__local_var_8_5 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_8_5 = (object)["map" => function($f_9) use ($__local_var_8_5) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($__local_var_8_5, $f_9) {
  $__num = \func_num_args();
  $__res = function($s_11) use ($__local_var_8_5, $f_9, $v_10) {
  $__num = \func_num_args();
  $__res = ((($__local_var_8_5)->{'map'})(function($v1_12) use ($f_9) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_9)(($v1_12)->{'value0'}), ($v1_12)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_10)($s_11));
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
  $Bind1_9_7 = (($Monad0_1_0)->{'Bind1'})(null);
  $Bind1_9_7 = (object)["bind" => function($v_10) use ($Bind1_9_7) {
  $__num = \func_num_args();
  $__res = function($f_11) use ($Bind1_9_7, $v_10) {
  $__num = \func_num_args();
  $__res = function($s_12) use ($Bind1_9_7, $f_11, $v_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_7)->{'bind'})(($v_10)($s_12)))(function($v1_13) use ($f_11) {
  $__num = \func_num_args();
  $__res = (($f_11)(($v1_13)->{'value0'}))(($v1_13)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_10) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_10_9 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($Monad0_1_0);
  $__res = (object)["apply" => function($f_11) use ($Applicative0_10_9, $Bind1_9_7) {
  $__num = \func_num_args();
  $__res = function($a_12) use ($Applicative0_10_9, $Bind1_9_7, $f_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_7)->{'bind'})($f_11))(function($f_prime__13) use ($Applicative0_10_9, $Bind1_9_7, $a_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_7)->{'bind'})($a_12))(function($a_prime__14) use ($Applicative0_10_9, $f_prime__13) {
  $__num = \func_num_args();
  $__res = (($Applicative0_10_9)->{'pure'})(($f_prime__13)($a_prime__14));
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
}, "Functor0" => function($_dollar___unused_9) use ($functorStateT1_8_5) {
  $__num = \func_num_args();
  $__res = $functorStateT1_8_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_7_11 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($Monad0_1_0);
  $__res = (object)["apply" => function($f_8) use ($Applicative0_7_11, $Bind1_6_4) {
  $__num = \func_num_args();
  $__res = function($a_9) use ($Applicative0_7_11, $Bind1_6_4, $f_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_4)->{'bind'})($f_8))(function($f_prime__10) use ($Applicative0_7_11, $Bind1_6_4, $a_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_4)->{'bind'})($a_9))(function($a_prime__11) use ($Applicative0_7_11, $f_prime__10) {
  $__num = \func_num_args();
  $__res = (($Applicative0_7_11)->{'pure'})(($f_prime__10)($a_prime__11));
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
}, "Functor0" => function($_dollar___unused_6) use ($functorStateT1_5_2) {
  $__num = \func_num_args();
  $__res = $functorStateT1_5_2;
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
}, "Bind1" => function($_dollar___unused_2) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $Bind1_3_12 = (($Monad0_1_0)->{'Bind1'})(null);
  $__res = (object)["bind" => function($v_4) use ($Bind1_3_12) {
  $__num = \func_num_args();
  $__res = function($f_5) use ($Bind1_3_12, $v_4) {
  $__num = \func_num_args();
  $__res = function($s_6) use ($Bind1_3_12, $f_5, $v_4) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_12)->{'bind'})(($v_4)($s_6)))(function($v1_7) use ($f_5) {
  $__num = \func_num_args();
  $__res = (($f_5)(($v1_7)->{'value0'}))(($v1_7)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_4) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__local_var_5_13 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_5_13 = (object)["map" => function($f_6) use ($__local_var_5_13) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($__local_var_5_13, $f_6) {
  $__num = \func_num_args();
  $__res = function($s_8) use ($__local_var_5_13, $f_6, $v_7) {
  $__num = \func_num_args();
  $__res = ((($__local_var_5_13)->{'map'})(function($v1_9) use ($f_6) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_6)(($v1_9)->{'value0'}), ($v1_9)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_7)($s_8));
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
  $Bind1_6_15 = (($Monad0_1_0)->{'Bind1'})(null);
  $Bind1_6_15 = (object)["bind" => function($v_7) use ($Bind1_6_15) {
  $__num = \func_num_args();
  $__res = function($f_8) use ($Bind1_6_15, $v_7) {
  $__num = \func_num_args();
  $__res = function($s_9) use ($Bind1_6_15, $f_8, $v_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_15)->{'bind'})(($v_7)($s_9)))(function($v1_10) use ($f_8) {
  $__num = \func_num_args();
  $__res = (($f_8)(($v1_10)->{'value0'}))(($v1_10)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_7) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $pure_7_17 = ((($Monad0_1_0)->{'Applicative0'})(null))->{'pure'};
  $Applicative0_7_17 = (object)["pure" => function($a_8) use ($pure_7_17) {
  $__num = \func_num_args();
  $__res = function($s_9) use ($a_8, $pure_7_17) {
  $__num = \func_num_args();
  $__res = ($pure_7_17)(new \Data\Tuple\Data_Tuple_Tuple($a_8, $s_9));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_8) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__local_var_9_18 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_9_18 = (object)["map" => function($f_10) use ($__local_var_9_18) {
  $__num = \func_num_args();
  $__res = function($v_11) use ($__local_var_9_18, $f_10) {
  $__num = \func_num_args();
  $__res = function($s_12) use ($__local_var_9_18, $f_10, $v_11) {
  $__num = \func_num_args();
  $__res = ((($__local_var_9_18)->{'map'})(function($v1_13) use ($f_10) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_10)(($v1_13)->{'value0'}), ($v1_13)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_11)($s_12));
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
  $Bind1_10_20 = (($Monad0_1_0)->{'Bind1'})(null);
  $Bind1_10_20 = (object)["bind" => function($v_11) use ($Bind1_10_20) {
  $__num = \func_num_args();
  $__res = function($f_12) use ($Bind1_10_20, $v_11) {
  $__num = \func_num_args();
  $__res = function($s_13) use ($Bind1_10_20, $f_12, $v_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_20)->{'bind'})(($v_11)($s_13)))(function($v1_14) use ($f_12) {
  $__num = \func_num_args();
  $__res = (($f_12)(($v1_14)->{'value0'}))(($v1_14)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_11) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_11_22 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($Monad0_1_0);
  $__res = (object)["apply" => function($f_12) use ($Applicative0_11_22, $Bind1_10_20) {
  $__num = \func_num_args();
  $__res = function($a_13) use ($Applicative0_11_22, $Bind1_10_20, $f_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_20)->{'bind'})($f_12))(function($f_prime__14) use ($Applicative0_11_22, $Bind1_10_20, $a_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_20)->{'bind'})($a_13))(function($a_prime__15) use ($Applicative0_11_22, $f_prime__14) {
  $__num = \func_num_args();
  $__res = (($Applicative0_11_22)->{'pure'})(($f_prime__14)($a_prime__15));
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
}, "Functor0" => function($_dollar___unused_10) use ($functorStateT1_9_18) {
  $__num = \func_num_args();
  $__res = $functorStateT1_9_18;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_8) use ($Applicative0_7_17, $Bind1_6_15) {
  $__num = \func_num_args();
  $__res = function($a_9) use ($Applicative0_7_17, $Bind1_6_15, $f_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_15)->{'bind'})($f_8))(function($f_prime__10) use ($Applicative0_7_17, $Bind1_6_15, $a_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_15)->{'bind'})($a_9))(function($a_prime__11) use ($Applicative0_7_17, $f_prime__10) {
  $__num = \func_num_args();
  $__res = (($Applicative0_7_17)->{'pure'})(($f_prime__10)($a_prime__11));
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
}, "Functor0" => function($_dollar___unused_6) use ($functorStateT1_5_13) {
  $__num = \func_num_args();
  $__res = $functorStateT1_5_13;
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
}];
  $Bind1_3_25 = (($Monad0_1_0)->{'Bind1'})(null);
  $pure_4_26 = ((($Monad0_1_0)->{'Applicative0'})(null))->{'pure'};
  $__res = (object)["liftST" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($m_5) use ($Bind1_3_25, $pure_4_26) {
  $__num = \func_num_args();
  $__res = function($s_6) use ($Bind1_3_25, $m_5, $pure_4_26) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_25)->{'bind'})($m_5))(function($x_7) use ($pure_4_26, $s_6) {
  $__num = \func_num_args();
  $__res = ($pure_4_26)(new \Data\Tuple\Data_Tuple_Tuple($x_7, $s_6));
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
}))(($dictMonadST_0)->{'liftST'}), "Monad0" => function($_dollar___unused_3) use ($monadStateT1_2_1) {
  $__num = \func_num_args();
  $__res = $monadStateT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_State_Trans_monadSTStateT'] = __NAMESPACE__ . '\\majControl_majMonad_majState_majTrans_monadmajSmajTmajStatemajT';

// Control_Monad_State_Trans_monoidStateT
function majControl_majMonad_majState_majTrans_monoidmajStatemajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majState_majTrans_monoidmajStatemajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $pure_1_0 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $applicativeStateT1_1_0 = (object)["pure" => function($a_2) use ($pure_1_0) {
  $__num = \func_num_args();
  $__res = function($s_3) use ($a_2, $pure_1_0) {
  $__num = \func_num_args();
  $__res = ($pure_1_0)(new \Data\Tuple\Data_Tuple_Tuple($a_2, $s_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_2) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_3_1 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_3_1 = (object)["map" => function($f_4) use ($__local_var_3_1) {
  $__num = \func_num_args();
  $__res = function($v_5) use ($__local_var_3_1, $f_4) {
  $__num = \func_num_args();
  $__res = function($s_6) use ($__local_var_3_1, $f_4, $v_5) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_1)->{'map'})(function($v1_7) use ($f_4) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_4)(($v1_7)->{'value0'}), ($v1_7)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_5)($s_6));
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
  $Bind1_4_3 = (($dictMonad_0)->{'Bind1'})(null);
  $Bind1_4_3 = (object)["bind" => function($v_5) use ($Bind1_4_3) {
  $__num = \func_num_args();
  $__res = function($f_6) use ($Bind1_4_3, $v_5) {
  $__num = \func_num_args();
  $__res = function($s_7) use ($Bind1_4_3, $f_6, $v_5) {
  $__num = \func_num_args();
  $__res = ((($Bind1_4_3)->{'bind'})(($v_5)($s_7)))(function($v1_8) use ($f_6) {
  $__num = \func_num_args();
  $__res = (($f_6)(($v1_8)->{'value0'}))(($v1_8)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_5) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_6_4 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_6_4 = (object)["map" => function($f_7) use ($__local_var_6_4) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($__local_var_6_4, $f_7) {
  $__num = \func_num_args();
  $__res = function($s_9) use ($__local_var_6_4, $f_7, $v_8) {
  $__num = \func_num_args();
  $__res = ((($__local_var_6_4)->{'map'})(function($v1_10) use ($f_7) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_7)(($v1_10)->{'value0'}), ($v1_10)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_8)($s_9));
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
  $Bind1_7_6 = (($dictMonad_0)->{'Bind1'})(null);
  $Bind1_7_6 = (object)["bind" => function($v_8) use ($Bind1_7_6) {
  $__num = \func_num_args();
  $__res = function($f_9) use ($Bind1_7_6, $v_8) {
  $__num = \func_num_args();
  $__res = function($s_10) use ($Bind1_7_6, $f_9, $v_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_6)->{'bind'})(($v_8)($s_10)))(function($v1_11) use ($f_9) {
  $__num = \func_num_args();
  $__res = (($f_9)(($v1_11)->{'value0'}))(($v1_11)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_8) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_8_8 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_9) use ($Applicative0_8_8, $Bind1_7_6) {
  $__num = \func_num_args();
  $__res = function($a_10) use ($Applicative0_8_8, $Bind1_7_6, $f_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_6)->{'bind'})($f_9))(function($f_prime__11) use ($Applicative0_8_8, $Bind1_7_6, $a_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_6)->{'bind'})($a_10))(function($a_prime__12) use ($Applicative0_8_8, $f_prime__11) {
  $__num = \func_num_args();
  $__res = (($Applicative0_8_8)->{'pure'})(($f_prime__11)($a_prime__12));
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
}, "Functor0" => function($_dollar___unused_7) use ($functorStateT1_6_4) {
  $__num = \func_num_args();
  $__res = $functorStateT1_6_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_5_10 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_6) use ($Applicative0_5_10, $Bind1_4_3) {
  $__num = \func_num_args();
  $__res = function($a_7) use ($Applicative0_5_10, $Bind1_4_3, $f_6) {
  $__num = \func_num_args();
  $__res = ((($Bind1_4_3)->{'bind'})($f_6))(function($f_prime__8) use ($Applicative0_5_10, $Bind1_4_3, $a_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_4_3)->{'bind'})($a_7))(function($a_prime__9) use ($Applicative0_5_10, $f_prime__8) {
  $__num = \func_num_args();
  $__res = (($Applicative0_5_10)->{'pure'})(($f_prime__8)($a_prime__9));
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
}, "Functor0" => function($_dollar___unused_4) use ($functorStateT1_3_1) {
  $__num = \func_num_args();
  $__res = $functorStateT1_3_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_2_12 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_2_12 = (object)["map" => function($f_3) use ($__local_var_2_12) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($__local_var_2_12, $f_3) {
  $__num = \func_num_args();
  $__res = function($s_5) use ($__local_var_2_12, $f_3, $v_4) {
  $__num = \func_num_args();
  $__res = ((($__local_var_2_12)->{'map'})(function($v1_6) use ($f_3) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_3)(($v1_6)->{'value0'}), ($v1_6)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_4)($s_5));
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
  $Bind1_3_14 = (($dictMonad_0)->{'Bind1'})(null);
  $Bind1_3_14 = (object)["bind" => function($v_4) use ($Bind1_3_14) {
  $__num = \func_num_args();
  $__res = function($f_5) use ($Bind1_3_14, $v_4) {
  $__num = \func_num_args();
  $__res = function($s_6) use ($Bind1_3_14, $f_5, $v_4) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_14)->{'bind'})(($v_4)($s_6)))(function($v1_7) use ($f_5) {
  $__num = \func_num_args();
  $__res = (($f_5)(($v1_7)->{'value0'}))(($v1_7)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_4) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_5_15 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_5_15 = (object)["map" => function($f_6) use ($__local_var_5_15) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($__local_var_5_15, $f_6) {
  $__num = \func_num_args();
  $__res = function($s_8) use ($__local_var_5_15, $f_6, $v_7) {
  $__num = \func_num_args();
  $__res = ((($__local_var_5_15)->{'map'})(function($v1_9) use ($f_6) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_6)(($v1_9)->{'value0'}), ($v1_9)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_7)($s_8));
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
  $Bind1_6_17 = (($dictMonad_0)->{'Bind1'})(null);
  $Bind1_6_17 = (object)["bind" => function($v_7) use ($Bind1_6_17) {
  $__num = \func_num_args();
  $__res = function($f_8) use ($Bind1_6_17, $v_7) {
  $__num = \func_num_args();
  $__res = function($s_9) use ($Bind1_6_17, $f_8, $v_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_17)->{'bind'})(($v_7)($s_9)))(function($v1_10) use ($f_8) {
  $__num = \func_num_args();
  $__res = (($f_8)(($v1_10)->{'value0'}))(($v1_10)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_7) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $pure_7_19 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $Applicative0_7_19 = (object)["pure" => function($a_8) use ($pure_7_19) {
  $__num = \func_num_args();
  $__res = function($s_9) use ($a_8, $pure_7_19) {
  $__num = \func_num_args();
  $__res = ($pure_7_19)(new \Data\Tuple\Data_Tuple_Tuple($a_8, $s_9));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_8) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_9_20 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_9_20 = (object)["map" => function($f_10) use ($__local_var_9_20) {
  $__num = \func_num_args();
  $__res = function($v_11) use ($__local_var_9_20, $f_10) {
  $__num = \func_num_args();
  $__res = function($s_12) use ($__local_var_9_20, $f_10, $v_11) {
  $__num = \func_num_args();
  $__res = ((($__local_var_9_20)->{'map'})(function($v1_13) use ($f_10) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_10)(($v1_13)->{'value0'}), ($v1_13)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_11)($s_12));
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
  $Bind1_10_22 = (($dictMonad_0)->{'Bind1'})(null);
  $Bind1_10_22 = (object)["bind" => function($v_11) use ($Bind1_10_22) {
  $__num = \func_num_args();
  $__res = function($f_12) use ($Bind1_10_22, $v_11) {
  $__num = \func_num_args();
  $__res = function($s_13) use ($Bind1_10_22, $f_12, $v_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_22)->{'bind'})(($v_11)($s_13)))(function($v1_14) use ($f_12) {
  $__num = \func_num_args();
  $__res = (($f_12)(($v1_14)->{'value0'}))(($v1_14)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_11) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_11_24 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_12) use ($Applicative0_11_24, $Bind1_10_22) {
  $__num = \func_num_args();
  $__res = function($a_13) use ($Applicative0_11_24, $Bind1_10_22, $f_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_22)->{'bind'})($f_12))(function($f_prime__14) use ($Applicative0_11_24, $Bind1_10_22, $a_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_22)->{'bind'})($a_13))(function($a_prime__15) use ($Applicative0_11_24, $f_prime__14) {
  $__num = \func_num_args();
  $__res = (($Applicative0_11_24)->{'pure'})(($f_prime__14)($a_prime__15));
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
}, "Functor0" => function($_dollar___unused_10) use ($functorStateT1_9_20) {
  $__num = \func_num_args();
  $__res = $functorStateT1_9_20;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_8) use ($Applicative0_7_19, $Bind1_6_17) {
  $__num = \func_num_args();
  $__res = function($a_9) use ($Applicative0_7_19, $Bind1_6_17, $f_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_17)->{'bind'})($f_8))(function($f_prime__10) use ($Applicative0_7_19, $Bind1_6_17, $a_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_17)->{'bind'})($a_9))(function($a_prime__11) use ($Applicative0_7_19, $f_prime__10) {
  $__num = \func_num_args();
  $__res = (($Applicative0_7_19)->{'pure'})(($f_prime__10)($a_prime__11));
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
}, "Functor0" => function($_dollar___unused_6) use ($functorStateT1_5_15) {
  $__num = \func_num_args();
  $__res = $functorStateT1_5_15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $pure_4_27 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $Applicative0_4_27 = (object)["pure" => function($a_5) use ($pure_4_27) {
  $__num = \func_num_args();
  $__res = function($s_6) use ($a_5, $pure_4_27) {
  $__num = \func_num_args();
  $__res = ($pure_4_27)(new \Data\Tuple\Data_Tuple_Tuple($a_5, $s_6));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_5) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_6_28 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_6_28 = (object)["map" => function($f_7) use ($__local_var_6_28) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($__local_var_6_28, $f_7) {
  $__num = \func_num_args();
  $__res = function($s_9) use ($__local_var_6_28, $f_7, $v_8) {
  $__num = \func_num_args();
  $__res = ((($__local_var_6_28)->{'map'})(function($v1_10) use ($f_7) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_7)(($v1_10)->{'value0'}), ($v1_10)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_8)($s_9));
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
  $Bind1_7_30 = (($dictMonad_0)->{'Bind1'})(null);
  $Bind1_7_30 = (object)["bind" => function($v_8) use ($Bind1_7_30) {
  $__num = \func_num_args();
  $__res = function($f_9) use ($Bind1_7_30, $v_8) {
  $__num = \func_num_args();
  $__res = function($s_10) use ($Bind1_7_30, $f_9, $v_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_30)->{'bind'})(($v_8)($s_10)))(function($v1_11) use ($f_9) {
  $__num = \func_num_args();
  $__res = (($f_9)(($v1_11)->{'value0'}))(($v1_11)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_8) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_9_31 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_9_31 = (object)["map" => function($f_10) use ($__local_var_9_31) {
  $__num = \func_num_args();
  $__res = function($v_11) use ($__local_var_9_31, $f_10) {
  $__num = \func_num_args();
  $__res = function($s_12) use ($__local_var_9_31, $f_10, $v_11) {
  $__num = \func_num_args();
  $__res = ((($__local_var_9_31)->{'map'})(function($v1_13) use ($f_10) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_10)(($v1_13)->{'value0'}), ($v1_13)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_11)($s_12));
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
  $Bind1_10_33 = (($dictMonad_0)->{'Bind1'})(null);
  $Bind1_10_33 = (object)["bind" => function($v_11) use ($Bind1_10_33) {
  $__num = \func_num_args();
  $__res = function($f_12) use ($Bind1_10_33, $v_11) {
  $__num = \func_num_args();
  $__res = function($s_13) use ($Bind1_10_33, $f_12, $v_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_33)->{'bind'})(($v_11)($s_13)))(function($v1_14) use ($f_12) {
  $__num = \func_num_args();
  $__res = (($f_12)(($v1_14)->{'value0'}))(($v1_14)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_11) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $pure_11_35 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $Applicative0_11_35 = (object)["pure" => function($a_12) use ($pure_11_35) {
  $__num = \func_num_args();
  $__res = function($s_13) use ($a_12, $pure_11_35) {
  $__num = \func_num_args();
  $__res = ($pure_11_35)(new \Data\Tuple\Data_Tuple_Tuple($a_12, $s_13));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_12) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_13_36 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_13_36 = (object)["map" => function($f_14) use ($__local_var_13_36) {
  $__num = \func_num_args();
  $__res = function($v_15) use ($__local_var_13_36, $f_14) {
  $__num = \func_num_args();
  $__res = function($s_16) use ($__local_var_13_36, $f_14, $v_15) {
  $__num = \func_num_args();
  $__res = ((($__local_var_13_36)->{'map'})(function($v1_17) use ($f_14) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_14)(($v1_17)->{'value0'}), ($v1_17)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_15)($s_16));
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
  $Bind1_14_38 = (($dictMonad_0)->{'Bind1'})(null);
  $Bind1_14_38 = (object)["bind" => function($v_15) use ($Bind1_14_38) {
  $__num = \func_num_args();
  $__res = function($f_16) use ($Bind1_14_38, $v_15) {
  $__num = \func_num_args();
  $__res = function($s_17) use ($Bind1_14_38, $f_16, $v_15) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_38)->{'bind'})(($v_15)($s_17)))(function($v1_18) use ($f_16) {
  $__num = \func_num_args();
  $__res = (($f_16)(($v1_18)->{'value0'}))(($v1_18)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_15) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_15_40 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_16) use ($Applicative0_15_40, $Bind1_14_38) {
  $__num = \func_num_args();
  $__res = function($a_17) use ($Applicative0_15_40, $Bind1_14_38, $f_16) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_38)->{'bind'})($f_16))(function($f_prime__18) use ($Applicative0_15_40, $Bind1_14_38, $a_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_38)->{'bind'})($a_17))(function($a_prime__19) use ($Applicative0_15_40, $f_prime__18) {
  $__num = \func_num_args();
  $__res = (($Applicative0_15_40)->{'pure'})(($f_prime__18)($a_prime__19));
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
}, "Functor0" => function($_dollar___unused_14) use ($functorStateT1_13_36) {
  $__num = \func_num_args();
  $__res = $functorStateT1_13_36;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_12) use ($Applicative0_11_35, $Bind1_10_33) {
  $__num = \func_num_args();
  $__res = function($a_13) use ($Applicative0_11_35, $Bind1_10_33, $f_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_33)->{'bind'})($f_12))(function($f_prime__14) use ($Applicative0_11_35, $Bind1_10_33, $a_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_33)->{'bind'})($a_13))(function($a_prime__15) use ($Applicative0_11_35, $f_prime__14) {
  $__num = \func_num_args();
  $__res = (($Applicative0_11_35)->{'pure'})(($f_prime__14)($a_prime__15));
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
}, "Functor0" => function($_dollar___unused_10) use ($functorStateT1_9_31) {
  $__num = \func_num_args();
  $__res = $functorStateT1_9_31;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $pure_8_43 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $Applicative0_8_43 = (object)["pure" => function($a_9) use ($pure_8_43) {
  $__num = \func_num_args();
  $__res = function($s_10) use ($a_9, $pure_8_43) {
  $__num = \func_num_args();
  $__res = ($pure_8_43)(new \Data\Tuple\Data_Tuple_Tuple($a_9, $s_10));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_9) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_10_44 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_10_44 = (object)["map" => function($f_11) use ($__local_var_10_44) {
  $__num = \func_num_args();
  $__res = function($v_12) use ($__local_var_10_44, $f_11) {
  $__num = \func_num_args();
  $__res = function($s_13) use ($__local_var_10_44, $f_11, $v_12) {
  $__num = \func_num_args();
  $__res = ((($__local_var_10_44)->{'map'})(function($v1_14) use ($f_11) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_11)(($v1_14)->{'value0'}), ($v1_14)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_12)($s_13));
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
  $Bind1_11_46 = (($dictMonad_0)->{'Bind1'})(null);
  $Bind1_11_46 = (object)["bind" => function($v_12) use ($Bind1_11_46) {
  $__num = \func_num_args();
  $__res = function($f_13) use ($Bind1_11_46, $v_12) {
  $__num = \func_num_args();
  $__res = function($s_14) use ($Bind1_11_46, $f_13, $v_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_46)->{'bind'})(($v_12)($s_14)))(function($v1_15) use ($f_13) {
  $__num = \func_num_args();
  $__res = (($f_13)(($v1_15)->{'value0'}))(($v1_15)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_12) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_13_47 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_13_47 = (object)["map" => function($f_14) use ($__local_var_13_47) {
  $__num = \func_num_args();
  $__res = function($v_15) use ($__local_var_13_47, $f_14) {
  $__num = \func_num_args();
  $__res = function($s_16) use ($__local_var_13_47, $f_14, $v_15) {
  $__num = \func_num_args();
  $__res = ((($__local_var_13_47)->{'map'})(function($v1_17) use ($f_14) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_14)(($v1_17)->{'value0'}), ($v1_17)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_15)($s_16));
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
  $Bind1_14_49 = (($dictMonad_0)->{'Bind1'})(null);
  $Bind1_14_49 = (object)["bind" => function($v_15) use ($Bind1_14_49) {
  $__num = \func_num_args();
  $__res = function($f_16) use ($Bind1_14_49, $v_15) {
  $__num = \func_num_args();
  $__res = function($s_17) use ($Bind1_14_49, $f_16, $v_15) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_49)->{'bind'})(($v_15)($s_17)))(function($v1_18) use ($f_16) {
  $__num = \func_num_args();
  $__res = (($f_16)(($v1_18)->{'value0'}))(($v1_18)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_15) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_15_51 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_16) use ($Applicative0_15_51, $Bind1_14_49) {
  $__num = \func_num_args();
  $__res = function($a_17) use ($Applicative0_15_51, $Bind1_14_49, $f_16) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_49)->{'bind'})($f_16))(function($f_prime__18) use ($Applicative0_15_51, $Bind1_14_49, $a_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_49)->{'bind'})($a_17))(function($a_prime__19) use ($Applicative0_15_51, $f_prime__18) {
  $__num = \func_num_args();
  $__res = (($Applicative0_15_51)->{'pure'})(($f_prime__18)($a_prime__19));
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
}, "Functor0" => function($_dollar___unused_14) use ($functorStateT1_13_47) {
  $__num = \func_num_args();
  $__res = $functorStateT1_13_47;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_12_53 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_13) use ($Applicative0_12_53, $Bind1_11_46) {
  $__num = \func_num_args();
  $__res = function($a_14) use ($Applicative0_12_53, $Bind1_11_46, $f_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_46)->{'bind'})($f_13))(function($f_prime__15) use ($Applicative0_12_53, $Bind1_11_46, $a_14) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_46)->{'bind'})($a_14))(function($a_prime__16) use ($Applicative0_12_53, $f_prime__15) {
  $__num = \func_num_args();
  $__res = (($Applicative0_12_53)->{'pure'})(($f_prime__15)($a_prime__16));
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
}, "Functor0" => function($_dollar___unused_11) use ($functorStateT1_10_44) {
  $__num = \func_num_args();
  $__res = $functorStateT1_10_44;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_9) use ($Applicative0_8_43, $Bind1_7_30) {
  $__num = \func_num_args();
  $__res = function($a_10) use ($Applicative0_8_43, $Bind1_7_30, $f_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_30)->{'bind'})($f_9))(function($f_prime__11) use ($Applicative0_8_43, $Bind1_7_30, $a_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_30)->{'bind'})($a_10))(function($a_prime__12) use ($Applicative0_8_43, $f_prime__11) {
  $__num = \func_num_args();
  $__res = (($Applicative0_8_43)->{'pure'})(($f_prime__11)($a_prime__12));
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
}, "Functor0" => function($_dollar___unused_7) use ($functorStateT1_6_28) {
  $__num = \func_num_args();
  $__res = $functorStateT1_6_28;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyStateT1_2_12 = (object)["apply" => function($f_5) use ($Applicative0_4_27, $Bind1_3_14) {
  $__num = \func_num_args();
  $__res = function($a_6) use ($Applicative0_4_27, $Bind1_3_14, $f_5) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_14)->{'bind'})($f_5))(function($f_prime__7) use ($Applicative0_4_27, $Bind1_3_14, $a_6) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_14)->{'bind'})($a_6))(function($a_prime__8) use ($Applicative0_4_27, $f_prime__7) {
  $__num = \func_num_args();
  $__res = (($Applicative0_4_27)->{'pure'})(($f_prime__7)($a_prime__8));
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
}, "Functor0" => function($_dollar___unused_3) use ($functorStateT1_2_12) {
  $__num = \func_num_args();
  $__res = $functorStateT1_2_12;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = function($dictMonoid_3) use ($applicativeStateT1_1_0, $applyStateT1_2_12) {
  $__num = \func_num_args();
  $Functor0_4_57 = (($applyStateT1_2_12)->{'Functor0'})(null);
  $__local_var_5_58 = ((($dictMonoid_3)->{'Semigroup0'})(null))->{'append'};
  $semigroupStateT2_4_57 = (object)["append" => function($a_6) use ($Functor0_4_57, $__local_var_5_58, $applyStateT1_2_12) {
  $__num = \func_num_args();
  $__res = function($b_7) use ($Functor0_4_57, $__local_var_5_58, $a_6, $applyStateT1_2_12) {
  $__num = \func_num_args();
  $__res = ((($applyStateT1_2_12)->{'apply'})(((($Functor0_4_57)->{'map'})($__local_var_5_58))($a_6)))($b_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["mempty" => (($applicativeStateT1_1_0)->{'pure'})(($dictMonoid_3)->{'mempty'}), "Semigroup0" => function($_dollar___unused_5) use ($semigroupStateT2_4_57) {
  $__num = \func_num_args();
  $__res = $semigroupStateT2_4_57;
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
$GLOBALS['Control_Monad_State_Trans_monoidStateT'] = __NAMESPACE__ . '\\majControl_majMonad_majState_majTrans_monoidmajStatemajT';

// Control_Monad_State_Trans_altStateT
function majControl_majMonad_majState_majTrans_altmajStatemajT($dictMonad_0, $dictAlt_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majState_majTrans_altmajStatemajT';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_2_0 = (($dictAlt_1)->{'Functor0'})(null);
  $functorStateT1_2_0 = (object)["map" => function($f_3) use ($__local_var_2_0) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($__local_var_2_0, $f_3) {
  $__num = \func_num_args();
  $__res = function($s_5) use ($__local_var_2_0, $f_3, $v_4) {
  $__num = \func_num_args();
  $__res = ((($__local_var_2_0)->{'map'})(function($v1_6) use ($f_3) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_3)(($v1_6)->{'value0'}), ($v1_6)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_4)($s_5));
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
  $__res = (object)["alt" => function($v_3) use ($dictAlt_1) {
  $__num = \func_num_args();
  $__res = function($v1_4) use ($dictAlt_1, $v_3) {
  $__num = \func_num_args();
  $__res = function($s_5) use ($dictAlt_1, $v1_4, $v_3) {
  $__num = \func_num_args();
  $__res = ((($dictAlt_1)->{'alt'})(($v_3)($s_5)))(($v1_4)($s_5));
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
}, "Functor0" => function($_dollar___unused_3) use ($functorStateT1_2_0) {
  $__num = \func_num_args();
  $__res = $functorStateT1_2_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Control_Monad_State_Trans_altStateT'] = __NAMESPACE__ . '\\majControl_majMonad_majState_majTrans_altmajStatemajT';

// Control_Monad_State_Trans_plusStateT
function majControl_majMonad_majState_majTrans_plusmajStatemajT($dictMonad_0, $dictPlus_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majState_majTrans_plusmajStatemajT';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $empty_2_0 = ($dictPlus_1)->{'empty'};
  $__local_var_3_1 = (($dictPlus_1)->{'Alt0'})(null);
  $__local_var_4_2 = (($__local_var_3_1)->{'Functor0'})(null);
  $functorStateT1_4_2 = (object)["map" => function($f_5) use ($__local_var_4_2) {
  $__num = \func_num_args();
  $__res = function($v_6) use ($__local_var_4_2, $f_5) {
  $__num = \func_num_args();
  $__res = function($s_7) use ($__local_var_4_2, $f_5, $v_6) {
  $__num = \func_num_args();
  $__res = ((($__local_var_4_2)->{'map'})(function($v1_8) use ($f_5) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_5)(($v1_8)->{'value0'}), ($v1_8)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_6)($s_7));
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
  $altStateT2_3_1 = (object)["alt" => function($v_5) use ($__local_var_3_1) {
  $__num = \func_num_args();
  $__res = function($v1_6) use ($__local_var_3_1, $v_5) {
  $__num = \func_num_args();
  $__res = function($s_7) use ($__local_var_3_1, $v1_6, $v_5) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_1)->{'alt'})(($v_5)($s_7)))(($v1_6)($s_7));
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
}, "Functor0" => function($_dollar___unused_5) use ($functorStateT1_4_2) {
  $__num = \func_num_args();
  $__res = $functorStateT1_4_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["empty" => function($v_4) use ($empty_2_0) {
  $__num = \func_num_args();
  $__res = $empty_2_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Alt0" => function($_dollar___unused_4) use ($altStateT2_3_1) {
  $__num = \func_num_args();
  $__res = $altStateT2_3_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Control_Monad_State_Trans_plusStateT'] = __NAMESPACE__ . '\\majControl_majMonad_majState_majTrans_plusmajStatemajT';

// Control_Monad_State_Trans_alternativeStateT
function majControl_majMonad_majState_majTrans_alternativemajStatemajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majState_majTrans_alternativemajStatemajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $pure_1_0 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $applicativeStateT1_1_0 = (object)["pure" => function($a_2) use ($pure_1_0) {
  $__num = \func_num_args();
  $__res = function($s_3) use ($a_2, $pure_1_0) {
  $__num = \func_num_args();
  $__res = ($pure_1_0)(new \Data\Tuple\Data_Tuple_Tuple($a_2, $s_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_2) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_3_1 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_3_1 = (object)["map" => function($f_4) use ($__local_var_3_1) {
  $__num = \func_num_args();
  $__res = function($v_5) use ($__local_var_3_1, $f_4) {
  $__num = \func_num_args();
  $__res = function($s_6) use ($__local_var_3_1, $f_4, $v_5) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_1)->{'map'})(function($v1_7) use ($f_4) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_4)(($v1_7)->{'value0'}), ($v1_7)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_5)($s_6));
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
  $Bind1_4_3 = (($dictMonad_0)->{'Bind1'})(null);
  $Bind1_4_3 = (object)["bind" => function($v_5) use ($Bind1_4_3) {
  $__num = \func_num_args();
  $__res = function($f_6) use ($Bind1_4_3, $v_5) {
  $__num = \func_num_args();
  $__res = function($s_7) use ($Bind1_4_3, $f_6, $v_5) {
  $__num = \func_num_args();
  $__res = ((($Bind1_4_3)->{'bind'})(($v_5)($s_7)))(function($v1_8) use ($f_6) {
  $__num = \func_num_args();
  $__res = (($f_6)(($v1_8)->{'value0'}))(($v1_8)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_5) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_6_4 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_6_4 = (object)["map" => function($f_7) use ($__local_var_6_4) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($__local_var_6_4, $f_7) {
  $__num = \func_num_args();
  $__res = function($s_9) use ($__local_var_6_4, $f_7, $v_8) {
  $__num = \func_num_args();
  $__res = ((($__local_var_6_4)->{'map'})(function($v1_10) use ($f_7) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_7)(($v1_10)->{'value0'}), ($v1_10)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_8)($s_9));
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
  $Bind1_7_6 = (($dictMonad_0)->{'Bind1'})(null);
  $Bind1_7_6 = (object)["bind" => function($v_8) use ($Bind1_7_6) {
  $__num = \func_num_args();
  $__res = function($f_9) use ($Bind1_7_6, $v_8) {
  $__num = \func_num_args();
  $__res = function($s_10) use ($Bind1_7_6, $f_9, $v_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_6)->{'bind'})(($v_8)($s_10)))(function($v1_11) use ($f_9) {
  $__num = \func_num_args();
  $__res = (($f_9)(($v1_11)->{'value0'}))(($v1_11)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_8) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_8_8 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_9) use ($Applicative0_8_8, $Bind1_7_6) {
  $__num = \func_num_args();
  $__res = function($a_10) use ($Applicative0_8_8, $Bind1_7_6, $f_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_6)->{'bind'})($f_9))(function($f_prime__11) use ($Applicative0_8_8, $Bind1_7_6, $a_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_6)->{'bind'})($a_10))(function($a_prime__12) use ($Applicative0_8_8, $f_prime__11) {
  $__num = \func_num_args();
  $__res = (($Applicative0_8_8)->{'pure'})(($f_prime__11)($a_prime__12));
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
}, "Functor0" => function($_dollar___unused_7) use ($functorStateT1_6_4) {
  $__num = \func_num_args();
  $__res = $functorStateT1_6_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_5_10 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_6) use ($Applicative0_5_10, $Bind1_4_3) {
  $__num = \func_num_args();
  $__res = function($a_7) use ($Applicative0_5_10, $Bind1_4_3, $f_6) {
  $__num = \func_num_args();
  $__res = ((($Bind1_4_3)->{'bind'})($f_6))(function($f_prime__8) use ($Applicative0_5_10, $Bind1_4_3, $a_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_4_3)->{'bind'})($a_7))(function($a_prime__9) use ($Applicative0_5_10, $f_prime__8) {
  $__num = \func_num_args();
  $__res = (($Applicative0_5_10)->{'pure'})(($f_prime__8)($a_prime__9));
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
}, "Functor0" => function($_dollar___unused_4) use ($functorStateT1_3_1) {
  $__num = \func_num_args();
  $__res = $functorStateT1_3_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = function($dictAlternative_2) use ($applicativeStateT1_1_0) {
  $__num = \func_num_args();
  $__local_var_3_12 = (($dictAlternative_2)->{'Plus1'})(null);
  $empty_4_13 = ($__local_var_3_12)->{'empty'};
  $__local_var_5_14 = (($__local_var_3_12)->{'Alt0'})(null);
  $__local_var_6_15 = (($__local_var_5_14)->{'Functor0'})(null);
  $functorStateT1_6_15 = (object)["map" => function($f_7) use ($__local_var_6_15) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($__local_var_6_15, $f_7) {
  $__num = \func_num_args();
  $__res = function($s_9) use ($__local_var_6_15, $f_7, $v_8) {
  $__num = \func_num_args();
  $__res = ((($__local_var_6_15)->{'map'})(function($v1_10) use ($f_7) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_7)(($v1_10)->{'value0'}), ($v1_10)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_8)($s_9));
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
  $altStateT2_5_14 = (object)["alt" => function($v_7) use ($__local_var_5_14) {
  $__num = \func_num_args();
  $__res = function($v1_8) use ($__local_var_5_14, $v_7) {
  $__num = \func_num_args();
  $__res = function($s_9) use ($__local_var_5_14, $v1_8, $v_7) {
  $__num = \func_num_args();
  $__res = ((($__local_var_5_14)->{'alt'})(($v_7)($s_9)))(($v1_8)($s_9));
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
}, "Functor0" => function($_dollar___unused_7) use ($functorStateT1_6_15) {
  $__num = \func_num_args();
  $__res = $functorStateT1_6_15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $plusStateT2_3_12 = (object)["empty" => function($v_6) use ($empty_4_13) {
  $__num = \func_num_args();
  $__res = $empty_4_13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Alt0" => function($_dollar___unused_6) use ($altStateT2_5_14) {
  $__num = \func_num_args();
  $__res = $altStateT2_5_14;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["Applicative0" => function($_dollar___unused_4) use ($applicativeStateT1_1_0) {
  $__num = \func_num_args();
  $__res = $applicativeStateT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Plus1" => function($_dollar___unused_4) use ($plusStateT2_3_12) {
  $__num = \func_num_args();
  $__res = $plusStateT2_3_12;
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
$GLOBALS['Control_Monad_State_Trans_alternativeStateT'] = __NAMESPACE__ . '\\majControl_majMonad_majState_majTrans_alternativemajStatemajT';

// Control_Monad_State_Trans_monadPlusStateT
function majControl_majMonad_majState_majTrans_monadmajPlusmajStatemajT($dictMonadPlus_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majState_majTrans_monadmajPlusmajStatemajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Monad0_1_0 = (($dictMonadPlus_0)->{'Monad0'})(null);
  $monadStateT1_2_1 = (object)["Applicative0" => function($_dollar___unused_2) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $pure_3_1 = ((($Monad0_1_0)->{'Applicative0'})(null))->{'pure'};
  $__res = (object)["pure" => function($a_4) use ($pure_3_1) {
  $__num = \func_num_args();
  $__res = function($s_5) use ($a_4, $pure_3_1) {
  $__num = \func_num_args();
  $__res = ($pure_3_1)(new \Data\Tuple\Data_Tuple_Tuple($a_4, $s_5));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_4) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__local_var_5_2 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_5_2 = (object)["map" => function($f_6) use ($__local_var_5_2) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($__local_var_5_2, $f_6) {
  $__num = \func_num_args();
  $__res = function($s_8) use ($__local_var_5_2, $f_6, $v_7) {
  $__num = \func_num_args();
  $__res = ((($__local_var_5_2)->{'map'})(function($v1_9) use ($f_6) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_6)(($v1_9)->{'value0'}), ($v1_9)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_7)($s_8));
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
  $Bind1_6_4 = (($Monad0_1_0)->{'Bind1'})(null);
  $Bind1_6_4 = (object)["bind" => function($v_7) use ($Bind1_6_4) {
  $__num = \func_num_args();
  $__res = function($f_8) use ($Bind1_6_4, $v_7) {
  $__num = \func_num_args();
  $__res = function($s_9) use ($Bind1_6_4, $f_8, $v_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_4)->{'bind'})(($v_7)($s_9)))(function($v1_10) use ($f_8) {
  $__num = \func_num_args();
  $__res = (($f_8)(($v1_10)->{'value0'}))(($v1_10)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_7) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__local_var_8_5 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_8_5 = (object)["map" => function($f_9) use ($__local_var_8_5) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($__local_var_8_5, $f_9) {
  $__num = \func_num_args();
  $__res = function($s_11) use ($__local_var_8_5, $f_9, $v_10) {
  $__num = \func_num_args();
  $__res = ((($__local_var_8_5)->{'map'})(function($v1_12) use ($f_9) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_9)(($v1_12)->{'value0'}), ($v1_12)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_10)($s_11));
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
  $Bind1_9_7 = (($Monad0_1_0)->{'Bind1'})(null);
  $Bind1_9_7 = (object)["bind" => function($v_10) use ($Bind1_9_7) {
  $__num = \func_num_args();
  $__res = function($f_11) use ($Bind1_9_7, $v_10) {
  $__num = \func_num_args();
  $__res = function($s_12) use ($Bind1_9_7, $f_11, $v_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_7)->{'bind'})(($v_10)($s_12)))(function($v1_13) use ($f_11) {
  $__num = \func_num_args();
  $__res = (($f_11)(($v1_13)->{'value0'}))(($v1_13)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_10) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_10_9 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($Monad0_1_0);
  $__res = (object)["apply" => function($f_11) use ($Applicative0_10_9, $Bind1_9_7) {
  $__num = \func_num_args();
  $__res = function($a_12) use ($Applicative0_10_9, $Bind1_9_7, $f_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_7)->{'bind'})($f_11))(function($f_prime__13) use ($Applicative0_10_9, $Bind1_9_7, $a_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_7)->{'bind'})($a_12))(function($a_prime__14) use ($Applicative0_10_9, $f_prime__13) {
  $__num = \func_num_args();
  $__res = (($Applicative0_10_9)->{'pure'})(($f_prime__13)($a_prime__14));
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
}, "Functor0" => function($_dollar___unused_9) use ($functorStateT1_8_5) {
  $__num = \func_num_args();
  $__res = $functorStateT1_8_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_7_11 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($Monad0_1_0);
  $__res = (object)["apply" => function($f_8) use ($Applicative0_7_11, $Bind1_6_4) {
  $__num = \func_num_args();
  $__res = function($a_9) use ($Applicative0_7_11, $Bind1_6_4, $f_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_4)->{'bind'})($f_8))(function($f_prime__10) use ($Applicative0_7_11, $Bind1_6_4, $a_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_4)->{'bind'})($a_9))(function($a_prime__11) use ($Applicative0_7_11, $f_prime__10) {
  $__num = \func_num_args();
  $__res = (($Applicative0_7_11)->{'pure'})(($f_prime__10)($a_prime__11));
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
}, "Functor0" => function($_dollar___unused_6) use ($functorStateT1_5_2) {
  $__num = \func_num_args();
  $__res = $functorStateT1_5_2;
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
}, "Bind1" => function($_dollar___unused_2) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $Bind1_3_12 = (($Monad0_1_0)->{'Bind1'})(null);
  $__res = (object)["bind" => function($v_4) use ($Bind1_3_12) {
  $__num = \func_num_args();
  $__res = function($f_5) use ($Bind1_3_12, $v_4) {
  $__num = \func_num_args();
  $__res = function($s_6) use ($Bind1_3_12, $f_5, $v_4) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_12)->{'bind'})(($v_4)($s_6)))(function($v1_7) use ($f_5) {
  $__num = \func_num_args();
  $__res = (($f_5)(($v1_7)->{'value0'}))(($v1_7)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_4) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__local_var_5_13 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_5_13 = (object)["map" => function($f_6) use ($__local_var_5_13) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($__local_var_5_13, $f_6) {
  $__num = \func_num_args();
  $__res = function($s_8) use ($__local_var_5_13, $f_6, $v_7) {
  $__num = \func_num_args();
  $__res = ((($__local_var_5_13)->{'map'})(function($v1_9) use ($f_6) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_6)(($v1_9)->{'value0'}), ($v1_9)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_7)($s_8));
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
  $Bind1_6_15 = (($Monad0_1_0)->{'Bind1'})(null);
  $Bind1_6_15 = (object)["bind" => function($v_7) use ($Bind1_6_15) {
  $__num = \func_num_args();
  $__res = function($f_8) use ($Bind1_6_15, $v_7) {
  $__num = \func_num_args();
  $__res = function($s_9) use ($Bind1_6_15, $f_8, $v_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_15)->{'bind'})(($v_7)($s_9)))(function($v1_10) use ($f_8) {
  $__num = \func_num_args();
  $__res = (($f_8)(($v1_10)->{'value0'}))(($v1_10)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_7) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $pure_7_17 = ((($Monad0_1_0)->{'Applicative0'})(null))->{'pure'};
  $Applicative0_7_17 = (object)["pure" => function($a_8) use ($pure_7_17) {
  $__num = \func_num_args();
  $__res = function($s_9) use ($a_8, $pure_7_17) {
  $__num = \func_num_args();
  $__res = ($pure_7_17)(new \Data\Tuple\Data_Tuple_Tuple($a_8, $s_9));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_8) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__local_var_9_18 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_9_18 = (object)["map" => function($f_10) use ($__local_var_9_18) {
  $__num = \func_num_args();
  $__res = function($v_11) use ($__local_var_9_18, $f_10) {
  $__num = \func_num_args();
  $__res = function($s_12) use ($__local_var_9_18, $f_10, $v_11) {
  $__num = \func_num_args();
  $__res = ((($__local_var_9_18)->{'map'})(function($v1_13) use ($f_10) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_10)(($v1_13)->{'value0'}), ($v1_13)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_11)($s_12));
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
  $Bind1_10_20 = (($Monad0_1_0)->{'Bind1'})(null);
  $Bind1_10_20 = (object)["bind" => function($v_11) use ($Bind1_10_20) {
  $__num = \func_num_args();
  $__res = function($f_12) use ($Bind1_10_20, $v_11) {
  $__num = \func_num_args();
  $__res = function($s_13) use ($Bind1_10_20, $f_12, $v_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_20)->{'bind'})(($v_11)($s_13)))(function($v1_14) use ($f_12) {
  $__num = \func_num_args();
  $__res = (($f_12)(($v1_14)->{'value0'}))(($v1_14)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_11) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_11_22 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($Monad0_1_0);
  $__res = (object)["apply" => function($f_12) use ($Applicative0_11_22, $Bind1_10_20) {
  $__num = \func_num_args();
  $__res = function($a_13) use ($Applicative0_11_22, $Bind1_10_20, $f_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_20)->{'bind'})($f_12))(function($f_prime__14) use ($Applicative0_11_22, $Bind1_10_20, $a_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_20)->{'bind'})($a_13))(function($a_prime__15) use ($Applicative0_11_22, $f_prime__14) {
  $__num = \func_num_args();
  $__res = (($Applicative0_11_22)->{'pure'})(($f_prime__14)($a_prime__15));
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
}, "Functor0" => function($_dollar___unused_10) use ($functorStateT1_9_18) {
  $__num = \func_num_args();
  $__res = $functorStateT1_9_18;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_8) use ($Applicative0_7_17, $Bind1_6_15) {
  $__num = \func_num_args();
  $__res = function($a_9) use ($Applicative0_7_17, $Bind1_6_15, $f_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_15)->{'bind'})($f_8))(function($f_prime__10) use ($Applicative0_7_17, $Bind1_6_15, $a_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_15)->{'bind'})($a_9))(function($a_prime__11) use ($Applicative0_7_17, $f_prime__10) {
  $__num = \func_num_args();
  $__res = (($Applicative0_7_17)->{'pure'})(($f_prime__10)($a_prime__11));
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
}, "Functor0" => function($_dollar___unused_6) use ($functorStateT1_5_13) {
  $__num = \func_num_args();
  $__res = $functorStateT1_5_13;
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
}];
  $pure_3_25 = ((($Monad0_1_0)->{'Applicative0'})(null))->{'pure'};
  $applicativeStateT1_3_25 = (object)["pure" => function($a_4) use ($pure_3_25) {
  $__num = \func_num_args();
  $__res = function($s_5) use ($a_4, $pure_3_25) {
  $__num = \func_num_args();
  $__res = ($pure_3_25)(new \Data\Tuple\Data_Tuple_Tuple($a_4, $s_5));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_4) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__local_var_5_26 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_5_26 = (object)["map" => function($f_6) use ($__local_var_5_26) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($__local_var_5_26, $f_6) {
  $__num = \func_num_args();
  $__res = function($s_8) use ($__local_var_5_26, $f_6, $v_7) {
  $__num = \func_num_args();
  $__res = ((($__local_var_5_26)->{'map'})(function($v1_9) use ($f_6) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_6)(($v1_9)->{'value0'}), ($v1_9)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_7)($s_8));
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
  $Bind1_6_28 = (($Monad0_1_0)->{'Bind1'})(null);
  $Bind1_6_28 = (object)["bind" => function($v_7) use ($Bind1_6_28) {
  $__num = \func_num_args();
  $__res = function($f_8) use ($Bind1_6_28, $v_7) {
  $__num = \func_num_args();
  $__res = function($s_9) use ($Bind1_6_28, $f_8, $v_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_28)->{'bind'})(($v_7)($s_9)))(function($v1_10) use ($f_8) {
  $__num = \func_num_args();
  $__res = (($f_8)(($v1_10)->{'value0'}))(($v1_10)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_7) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__local_var_8_29 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_8_29 = (object)["map" => function($f_9) use ($__local_var_8_29) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($__local_var_8_29, $f_9) {
  $__num = \func_num_args();
  $__res = function($s_11) use ($__local_var_8_29, $f_9, $v_10) {
  $__num = \func_num_args();
  $__res = ((($__local_var_8_29)->{'map'})(function($v1_12) use ($f_9) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_9)(($v1_12)->{'value0'}), ($v1_12)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_10)($s_11));
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
  $Bind1_9_31 = (($Monad0_1_0)->{'Bind1'})(null);
  $Bind1_9_31 = (object)["bind" => function($v_10) use ($Bind1_9_31) {
  $__num = \func_num_args();
  $__res = function($f_11) use ($Bind1_9_31, $v_10) {
  $__num = \func_num_args();
  $__res = function($s_12) use ($Bind1_9_31, $f_11, $v_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_31)->{'bind'})(($v_10)($s_12)))(function($v1_13) use ($f_11) {
  $__num = \func_num_args();
  $__res = (($f_11)(($v1_13)->{'value0'}))(($v1_13)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_10) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__local_var_11_32 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_11_32 = (object)["map" => function($f_12) use ($__local_var_11_32) {
  $__num = \func_num_args();
  $__res = function($v_13) use ($__local_var_11_32, $f_12) {
  $__num = \func_num_args();
  $__res = function($s_14) use ($__local_var_11_32, $f_12, $v_13) {
  $__num = \func_num_args();
  $__res = ((($__local_var_11_32)->{'map'})(function($v1_15) use ($f_12) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_12)(($v1_15)->{'value0'}), ($v1_15)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_13)($s_14));
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
  $Bind1_12_34 = (($Monad0_1_0)->{'Bind1'})(null);
  $Bind1_12_34 = (object)["bind" => function($v_13) use ($Bind1_12_34) {
  $__num = \func_num_args();
  $__res = function($f_14) use ($Bind1_12_34, $v_13) {
  $__num = \func_num_args();
  $__res = function($s_15) use ($Bind1_12_34, $f_14, $v_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_12_34)->{'bind'})(($v_13)($s_15)))(function($v1_16) use ($f_14) {
  $__num = \func_num_args();
  $__res = (($f_14)(($v1_16)->{'value0'}))(($v1_16)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_13) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $pure_13_36 = ((($Monad0_1_0)->{'Applicative0'})(null))->{'pure'};
  $Applicative0_13_36 = (object)["pure" => function($a_14) use ($pure_13_36) {
  $__num = \func_num_args();
  $__res = function($s_15) use ($a_14, $pure_13_36) {
  $__num = \func_num_args();
  $__res = ($pure_13_36)(new \Data\Tuple\Data_Tuple_Tuple($a_14, $s_15));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_14) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__local_var_15_37 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_15_37 = (object)["map" => function($f_16) use ($__local_var_15_37) {
  $__num = \func_num_args();
  $__res = function($v_17) use ($__local_var_15_37, $f_16) {
  $__num = \func_num_args();
  $__res = function($s_18) use ($__local_var_15_37, $f_16, $v_17) {
  $__num = \func_num_args();
  $__res = ((($__local_var_15_37)->{'map'})(function($v1_19) use ($f_16) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_16)(($v1_19)->{'value0'}), ($v1_19)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_17)($s_18));
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
  $Bind1_16_39 = (($Monad0_1_0)->{'Bind1'})(null);
  $Bind1_16_39 = (object)["bind" => function($v_17) use ($Bind1_16_39) {
  $__num = \func_num_args();
  $__res = function($f_18) use ($Bind1_16_39, $v_17) {
  $__num = \func_num_args();
  $__res = function($s_19) use ($Bind1_16_39, $f_18, $v_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_16_39)->{'bind'})(($v_17)($s_19)))(function($v1_20) use ($f_18) {
  $__num = \func_num_args();
  $__res = (($f_18)(($v1_20)->{'value0'}))(($v1_20)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_17) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_17_41 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($Monad0_1_0);
  $__res = (object)["apply" => function($f_18) use ($Applicative0_17_41, $Bind1_16_39) {
  $__num = \func_num_args();
  $__res = function($a_19) use ($Applicative0_17_41, $Bind1_16_39, $f_18) {
  $__num = \func_num_args();
  $__res = ((($Bind1_16_39)->{'bind'})($f_18))(function($f_prime__20) use ($Applicative0_17_41, $Bind1_16_39, $a_19) {
  $__num = \func_num_args();
  $__res = ((($Bind1_16_39)->{'bind'})($a_19))(function($a_prime__21) use ($Applicative0_17_41, $f_prime__20) {
  $__num = \func_num_args();
  $__res = (($Applicative0_17_41)->{'pure'})(($f_prime__20)($a_prime__21));
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
}, "Functor0" => function($_dollar___unused_16) use ($functorStateT1_15_37) {
  $__num = \func_num_args();
  $__res = $functorStateT1_15_37;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_14) use ($Applicative0_13_36, $Bind1_12_34) {
  $__num = \func_num_args();
  $__res = function($a_15) use ($Applicative0_13_36, $Bind1_12_34, $f_14) {
  $__num = \func_num_args();
  $__res = ((($Bind1_12_34)->{'bind'})($f_14))(function($f_prime__16) use ($Applicative0_13_36, $Bind1_12_34, $a_15) {
  $__num = \func_num_args();
  $__res = ((($Bind1_12_34)->{'bind'})($a_15))(function($a_prime__17) use ($Applicative0_13_36, $f_prime__16) {
  $__num = \func_num_args();
  $__res = (($Applicative0_13_36)->{'pure'})(($f_prime__16)($a_prime__17));
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
}, "Functor0" => function($_dollar___unused_12) use ($functorStateT1_11_32) {
  $__num = \func_num_args();
  $__res = $functorStateT1_11_32;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $pure_10_44 = ((($Monad0_1_0)->{'Applicative0'})(null))->{'pure'};
  $Applicative0_10_44 = (object)["pure" => function($a_11) use ($pure_10_44) {
  $__num = \func_num_args();
  $__res = function($s_12) use ($a_11, $pure_10_44) {
  $__num = \func_num_args();
  $__res = ($pure_10_44)(new \Data\Tuple\Data_Tuple_Tuple($a_11, $s_12));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_11) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__local_var_12_45 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_12_45 = (object)["map" => function($f_13) use ($__local_var_12_45) {
  $__num = \func_num_args();
  $__res = function($v_14) use ($__local_var_12_45, $f_13) {
  $__num = \func_num_args();
  $__res = function($s_15) use ($__local_var_12_45, $f_13, $v_14) {
  $__num = \func_num_args();
  $__res = ((($__local_var_12_45)->{'map'})(function($v1_16) use ($f_13) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_13)(($v1_16)->{'value0'}), ($v1_16)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_14)($s_15));
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
  $Bind1_13_47 = (($Monad0_1_0)->{'Bind1'})(null);
  $Bind1_13_47 = (object)["bind" => function($v_14) use ($Bind1_13_47) {
  $__num = \func_num_args();
  $__res = function($f_15) use ($Bind1_13_47, $v_14) {
  $__num = \func_num_args();
  $__res = function($s_16) use ($Bind1_13_47, $f_15, $v_14) {
  $__num = \func_num_args();
  $__res = ((($Bind1_13_47)->{'bind'})(($v_14)($s_16)))(function($v1_17) use ($f_15) {
  $__num = \func_num_args();
  $__res = (($f_15)(($v1_17)->{'value0'}))(($v1_17)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_14) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__local_var_15_48 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_15_48 = (object)["map" => function($f_16) use ($__local_var_15_48) {
  $__num = \func_num_args();
  $__res = function($v_17) use ($__local_var_15_48, $f_16) {
  $__num = \func_num_args();
  $__res = function($s_18) use ($__local_var_15_48, $f_16, $v_17) {
  $__num = \func_num_args();
  $__res = ((($__local_var_15_48)->{'map'})(function($v1_19) use ($f_16) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_16)(($v1_19)->{'value0'}), ($v1_19)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_17)($s_18));
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
  $Bind1_16_50 = (($Monad0_1_0)->{'Bind1'})(null);
  $Bind1_16_50 = (object)["bind" => function($v_17) use ($Bind1_16_50) {
  $__num = \func_num_args();
  $__res = function($f_18) use ($Bind1_16_50, $v_17) {
  $__num = \func_num_args();
  $__res = function($s_19) use ($Bind1_16_50, $f_18, $v_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_16_50)->{'bind'})(($v_17)($s_19)))(function($v1_20) use ($f_18) {
  $__num = \func_num_args();
  $__res = (($f_18)(($v1_20)->{'value0'}))(($v1_20)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_17) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_17_52 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($Monad0_1_0);
  $__res = (object)["apply" => function($f_18) use ($Applicative0_17_52, $Bind1_16_50) {
  $__num = \func_num_args();
  $__res = function($a_19) use ($Applicative0_17_52, $Bind1_16_50, $f_18) {
  $__num = \func_num_args();
  $__res = ((($Bind1_16_50)->{'bind'})($f_18))(function($f_prime__20) use ($Applicative0_17_52, $Bind1_16_50, $a_19) {
  $__num = \func_num_args();
  $__res = ((($Bind1_16_50)->{'bind'})($a_19))(function($a_prime__21) use ($Applicative0_17_52, $f_prime__20) {
  $__num = \func_num_args();
  $__res = (($Applicative0_17_52)->{'pure'})(($f_prime__20)($a_prime__21));
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
}, "Functor0" => function($_dollar___unused_16) use ($functorStateT1_15_48) {
  $__num = \func_num_args();
  $__res = $functorStateT1_15_48;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_14_54 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($Monad0_1_0);
  $__res = (object)["apply" => function($f_15) use ($Applicative0_14_54, $Bind1_13_47) {
  $__num = \func_num_args();
  $__res = function($a_16) use ($Applicative0_14_54, $Bind1_13_47, $f_15) {
  $__num = \func_num_args();
  $__res = ((($Bind1_13_47)->{'bind'})($f_15))(function($f_prime__17) use ($Applicative0_14_54, $Bind1_13_47, $a_16) {
  $__num = \func_num_args();
  $__res = ((($Bind1_13_47)->{'bind'})($a_16))(function($a_prime__18) use ($Applicative0_14_54, $f_prime__17) {
  $__num = \func_num_args();
  $__res = (($Applicative0_14_54)->{'pure'})(($f_prime__17)($a_prime__18));
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
}, "Functor0" => function($_dollar___unused_13) use ($functorStateT1_12_45) {
  $__num = \func_num_args();
  $__res = $functorStateT1_12_45;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_11) use ($Applicative0_10_44, $Bind1_9_31) {
  $__num = \func_num_args();
  $__res = function($a_12) use ($Applicative0_10_44, $Bind1_9_31, $f_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_31)->{'bind'})($f_11))(function($f_prime__13) use ($Applicative0_10_44, $Bind1_9_31, $a_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_31)->{'bind'})($a_12))(function($a_prime__14) use ($Applicative0_10_44, $f_prime__13) {
  $__num = \func_num_args();
  $__res = (($Applicative0_10_44)->{'pure'})(($f_prime__13)($a_prime__14));
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
}, "Functor0" => function($_dollar___unused_9) use ($functorStateT1_8_29) {
  $__num = \func_num_args();
  $__res = $functorStateT1_8_29;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $pure_7_57 = ((($Monad0_1_0)->{'Applicative0'})(null))->{'pure'};
  $Applicative0_7_57 = (object)["pure" => function($a_8) use ($pure_7_57) {
  $__num = \func_num_args();
  $__res = function($s_9) use ($a_8, $pure_7_57) {
  $__num = \func_num_args();
  $__res = ($pure_7_57)(new \Data\Tuple\Data_Tuple_Tuple($a_8, $s_9));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_8) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__local_var_9_58 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_9_58 = (object)["map" => function($f_10) use ($__local_var_9_58) {
  $__num = \func_num_args();
  $__res = function($v_11) use ($__local_var_9_58, $f_10) {
  $__num = \func_num_args();
  $__res = function($s_12) use ($__local_var_9_58, $f_10, $v_11) {
  $__num = \func_num_args();
  $__res = ((($__local_var_9_58)->{'map'})(function($v1_13) use ($f_10) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_10)(($v1_13)->{'value0'}), ($v1_13)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_11)($s_12));
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
  $Bind1_10_60 = (($Monad0_1_0)->{'Bind1'})(null);
  $Bind1_10_60 = (object)["bind" => function($v_11) use ($Bind1_10_60) {
  $__num = \func_num_args();
  $__res = function($f_12) use ($Bind1_10_60, $v_11) {
  $__num = \func_num_args();
  $__res = function($s_13) use ($Bind1_10_60, $f_12, $v_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_60)->{'bind'})(($v_11)($s_13)))(function($v1_14) use ($f_12) {
  $__num = \func_num_args();
  $__res = (($f_12)(($v1_14)->{'value0'}))(($v1_14)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_11) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__local_var_12_61 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorStateT1_12_61 = (object)["map" => function($f_13) use ($__local_var_12_61) {
  $__num = \func_num_args();
  $__res = function($v_14) use ($__local_var_12_61, $f_13) {
  $__num = \func_num_args();
  $__res = function($s_15) use ($__local_var_12_61, $f_13, $v_14) {
  $__num = \func_num_args();
  $__res = ((($__local_var_12_61)->{'map'})(function($v1_16) use ($f_13) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_13)(($v1_16)->{'value0'}), ($v1_16)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_14)($s_15));
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
  $Bind1_13_63 = (($Monad0_1_0)->{'Bind1'})(null);
  $Bind1_13_63 = (object)["bind" => function($v_14) use ($Bind1_13_63) {
  $__num = \func_num_args();
  $__res = function($f_15) use ($Bind1_13_63, $v_14) {
  $__num = \func_num_args();
  $__res = function($s_16) use ($Bind1_13_63, $f_15, $v_14) {
  $__num = \func_num_args();
  $__res = ((($Bind1_13_63)->{'bind'})(($v_14)($s_16)))(function($v1_17) use ($f_15) {
  $__num = \func_num_args();
  $__res = (($f_15)(($v1_17)->{'value0'}))(($v1_17)->{'value1'});
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
}, "Apply0" => function($_dollar___unused_14) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_State_Trans_applyStateT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_14_65 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($Monad0_1_0);
  $__res = (object)["apply" => function($f_15) use ($Applicative0_14_65, $Bind1_13_63) {
  $__num = \func_num_args();
  $__res = function($a_16) use ($Applicative0_14_65, $Bind1_13_63, $f_15) {
  $__num = \func_num_args();
  $__res = ((($Bind1_13_63)->{'bind'})($f_15))(function($f_prime__17) use ($Applicative0_14_65, $Bind1_13_63, $a_16) {
  $__num = \func_num_args();
  $__res = ((($Bind1_13_63)->{'bind'})($a_16))(function($a_prime__18) use ($Applicative0_14_65, $f_prime__17) {
  $__num = \func_num_args();
  $__res = (($Applicative0_14_65)->{'pure'})(($f_prime__17)($a_prime__18));
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
}, "Functor0" => function($_dollar___unused_13) use ($functorStateT1_12_61) {
  $__num = \func_num_args();
  $__res = $functorStateT1_12_61;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_11_67 = ($GLOBALS['Control_Monad_State_Trans_applicativeStateT'])($Monad0_1_0);
  $__res = (object)["apply" => function($f_12) use ($Applicative0_11_67, $Bind1_10_60) {
  $__num = \func_num_args();
  $__res = function($a_13) use ($Applicative0_11_67, $Bind1_10_60, $f_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_60)->{'bind'})($f_12))(function($f_prime__14) use ($Applicative0_11_67, $Bind1_10_60, $a_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_60)->{'bind'})($a_13))(function($a_prime__15) use ($Applicative0_11_67, $f_prime__14) {
  $__num = \func_num_args();
  $__res = (($Applicative0_11_67)->{'pure'})(($f_prime__14)($a_prime__15));
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
}, "Functor0" => function($_dollar___unused_10) use ($functorStateT1_9_58) {
  $__num = \func_num_args();
  $__res = $functorStateT1_9_58;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_8) use ($Applicative0_7_57, $Bind1_6_28) {
  $__num = \func_num_args();
  $__res = function($a_9) use ($Applicative0_7_57, $Bind1_6_28, $f_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_28)->{'bind'})($f_8))(function($f_prime__10) use ($Applicative0_7_57, $Bind1_6_28, $a_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_28)->{'bind'})($a_9))(function($a_prime__11) use ($Applicative0_7_57, $f_prime__10) {
  $__num = \func_num_args();
  $__res = (($Applicative0_7_57)->{'pure'})(($f_prime__10)($a_prime__11));
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
}, "Functor0" => function($_dollar___unused_6) use ($functorStateT1_5_26) {
  $__num = \func_num_args();
  $__res = $functorStateT1_5_26;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_4_70 = (((($dictMonadPlus_0)->{'Alternative1'})(null))->{'Plus1'})(null);
  $empty_5_71 = ($__local_var_4_70)->{'empty'};
  $__local_var_6_72 = (($__local_var_4_70)->{'Alt0'})(null);
  $__local_var_7_73 = (($__local_var_6_72)->{'Functor0'})(null);
  $functorStateT1_7_73 = (object)["map" => function($f_8) use ($__local_var_7_73) {
  $__num = \func_num_args();
  $__res = function($v_9) use ($__local_var_7_73, $f_8) {
  $__num = \func_num_args();
  $__res = function($s_10) use ($__local_var_7_73, $f_8, $v_9) {
  $__num = \func_num_args();
  $__res = ((($__local_var_7_73)->{'map'})(function($v1_11) use ($f_8) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_8)(($v1_11)->{'value0'}), ($v1_11)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_9)($s_10));
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
  $altStateT2_6_72 = (object)["alt" => function($v_8) use ($__local_var_6_72) {
  $__num = \func_num_args();
  $__res = function($v1_9) use ($__local_var_6_72, $v_8) {
  $__num = \func_num_args();
  $__res = function($s_10) use ($__local_var_6_72, $v1_9, $v_8) {
  $__num = \func_num_args();
  $__res = ((($__local_var_6_72)->{'alt'})(($v_8)($s_10)))(($v1_9)($s_10));
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
}, "Functor0" => function($_dollar___unused_8) use ($functorStateT1_7_73) {
  $__num = \func_num_args();
  $__res = $functorStateT1_7_73;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $plusStateT2_4_70 = (object)["empty" => function($v_7) use ($empty_5_71) {
  $__num = \func_num_args();
  $__res = $empty_5_71;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Alt0" => function($_dollar___unused_7) use ($altStateT2_6_72) {
  $__num = \func_num_args();
  $__res = $altStateT2_6_72;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $alternativeStateT1_3_25 = (object)["Applicative0" => function($_dollar___unused_5) use ($applicativeStateT1_3_25) {
  $__num = \func_num_args();
  $__res = $applicativeStateT1_3_25;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Plus1" => function($_dollar___unused_5) use ($plusStateT2_4_70) {
  $__num = \func_num_args();
  $__res = $plusStateT2_4_70;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["Monad0" => function($_dollar___unused_4) use ($monadStateT1_2_1) {
  $__num = \func_num_args();
  $__res = $monadStateT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Alternative1" => function($_dollar___unused_4) use ($alternativeStateT1_3_25) {
  $__num = \func_num_args();
  $__res = $alternativeStateT1_3_25;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_State_Trans_monadPlusStateT'] = __NAMESPACE__ . '\\majControl_majMonad_majState_majTrans_monadmajPlusmajStatemajT';

