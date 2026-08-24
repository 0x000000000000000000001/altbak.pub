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
if (!\function_exists(__NAMESPACE__ . '\\phpurs_execute_effect')) {
  function phpurs_execute_effect($val) {
    if (\is_callable($val)) {
      return $val($GLOBALS['Data_Unit_unit']);
    }
    return $val;
  }
}

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };




// Control_Monad_Reader_Trans_ReaderT
function majControl_majMonad_majReader_majTrans_majReadermajT($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majReader_majTrans_majReadermajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Reader_Trans_ReaderT'] = __NAMESPACE__ . '\\majControl_majMonad_majReader_majTrans_majReadermajT';

// Control_Monad_Reader_Trans_withReaderT
function majControl_majMonad_majReader_majTrans_withmajReadermajT($f_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majReader_majTrans_withmajReadermajT';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($v_1))($f_0);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Control_Monad_Reader_Trans_withReaderT'] = __NAMESPACE__ . '\\majControl_majMonad_majReader_majTrans_withmajReadermajT';

// Control_Monad_Reader_Trans_runReaderT
function majControl_majMonad_majReader_majTrans_runmajReadermajT($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majReader_majTrans_runmajReadermajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $v_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Reader_Trans_runReaderT'] = __NAMESPACE__ . '\\majControl_majMonad_majReader_majTrans_runmajReadermajT';

// Control_Monad_Reader_Trans_newtypeReaderT
$GLOBALS['Control_Monad_Reader_Trans_newtypeReaderT'] = (object)["Coercible0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Control_Monad_Reader_Trans_monadTransReaderT
$GLOBALS['Control_Monad_Reader_Trans_monadTransReaderT'] = (object)["lift" => function($dictMonad_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_1) {
  $__num = \func_num_args();
  $__res = $x_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Function_const']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Control_Monad_Reader_Trans_mapReaderT
function majControl_majMonad_majReader_majTrans_mapmajReadermajT($f_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majReader_majTrans_mapmajReadermajT';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($f_0))($v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Control_Monad_Reader_Trans_mapReaderT'] = __NAMESPACE__ . '\\majControl_majMonad_majReader_majTrans_mapmajReadermajT';

// Control_Monad_Reader_Trans_functorReaderT
function majControl_majMonad_majReader_majTrans_functormajReadermajT($dictFunctor_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majReader_majTrans_functormajReadermajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["map" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Reader_Trans_mapReaderT']))(($dictFunctor_0)->{'map'})];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Reader_Trans_functorReaderT'] = __NAMESPACE__ . '\\majControl_majMonad_majReader_majTrans_functormajReadermajT';

// Control_Monad_Reader_Trans_distributiveReaderT
function majControl_majMonad_majReader_majTrans_distributivemajReadermajT($dictDistributive_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majReader_majTrans_distributivemajReadermajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__tco_var_Control_Monad_Reader_Trans_distributiveReaderT_dictDistributive_0 = $dictDistributive_0;
  tco_loop_Control_Monad_Reader_Trans_distributiveReaderT:;
  $dictDistributive_0 = $__tco_var_Control_Monad_Reader_Trans_distributiveReaderT_dictDistributive_0;
  $functorReaderT1_1_0 = (object)["map" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Reader_Trans_mapReaderT']))(((($dictDistributive_0)->{'Functor0'})(null))->{'map'})];
  $__res = (object)["distribute" => function($dictFunctor_2) use ($dictDistributive_0) {
  $__num = \func_num_args();
  $__res = function($a_3) use ($dictDistributive_0, $dictFunctor_2) {
  $__num = \func_num_args();
  $__res = function($e_4) use ($a_3, $dictDistributive_0, $dictFunctor_2) {
  $__num = \func_num_args();
  $__res = (((($dictDistributive_0)->{'collect'})($dictFunctor_2))(function($r_5) use ($e_4) {
  $__num = \func_num_args();
  $__res = ($r_5)($e_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($a_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "collect" => function($dictFunctor_2) use ($dictDistributive_0) {
  $__num = \func_num_args();
  $__res = function($f_3) use ($dictDistributive_0, $dictFunctor_2) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(((($GLOBALS['Control_Monad_Reader_Trans_distributiveReaderT'])($dictDistributive_0))->{'distribute'})($dictFunctor_2)))((($dictFunctor_2)->{'map'})($f_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_2) use ($functorReaderT1_1_0) {
  $__num = \func_num_args();
  $__res = $functorReaderT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Reader_Trans_distributiveReaderT'] = __NAMESPACE__ . '\\majControl_majMonad_majReader_majTrans_distributivemajReadermajT';

// Control_Monad_Reader_Trans_applyReaderT
function majControl_majMonad_majReader_majTrans_applymajReadermajT($dictApply_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majReader_majTrans_applymajReadermajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $functorReaderT1_1_0 = (object)["map" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Reader_Trans_mapReaderT']))(((($dictApply_0)->{'Functor0'})(null))->{'map'})];
  $__res = (object)["apply" => function($v_2) use ($dictApply_0) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($dictApply_0, $v_2) {
  $__num = \func_num_args();
  $__res = function($r_4) use ($dictApply_0, $v1_3, $v_2) {
  $__num = \func_num_args();
  $__res = ((($dictApply_0)->{'apply'})(($v_2)($r_4)))(($v1_3)($r_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_2) use ($functorReaderT1_1_0) {
  $__num = \func_num_args();
  $__res = $functorReaderT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Reader_Trans_applyReaderT'] = __NAMESPACE__ . '\\majControl_majMonad_majReader_majTrans_applymajReadermajT';

// Control_Monad_Reader_Trans_bindReaderT
function majControl_majMonad_majReader_majTrans_bindmajReadermajT($dictBind_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majReader_majTrans_bindmajReadermajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictBind_0)->{'Apply0'})(null);
  $functorReaderT1_2_1 = (object)["map" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Reader_Trans_mapReaderT']))(((($__local_var_1_0)->{'Functor0'})(null))->{'map'})];
  $applyReaderT1_1_0 = (object)["apply" => function($v_3) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v1_4) use ($__local_var_1_0, $v_3) {
  $__num = \func_num_args();
  $__res = function($r_5) use ($__local_var_1_0, $v1_4, $v_3) {
  $__num = \func_num_args();
  $__res = ((($__local_var_1_0)->{'apply'})(($v_3)($r_5)))(($v1_4)($r_5));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_3) use ($functorReaderT1_2_1) {
  $__num = \func_num_args();
  $__res = $functorReaderT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["bind" => function($v_2) use ($dictBind_0) {
  $__num = \func_num_args();
  $__res = function($k_3) use ($dictBind_0, $v_2) {
  $__num = \func_num_args();
  $__res = function($r_4) use ($dictBind_0, $k_3, $v_2) {
  $__num = \func_num_args();
  $__res = ((($dictBind_0)->{'bind'})(($v_2)($r_4)))(function($a_5) use ($k_3, $r_4) {
  $__num = \func_num_args();
  $__res = (($k_3)($a_5))($r_4);
  goto __end;;
  __end:
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
}, "Apply0" => function($_dollar___unused_2) use ($applyReaderT1_1_0) {
  $__num = \func_num_args();
  $__res = $applyReaderT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Reader_Trans_bindReaderT'] = __NAMESPACE__ . '\\majControl_majMonad_majReader_majTrans_bindmajReadermajT';

// Control_Monad_Reader_Trans_semigroupReaderT
function majControl_majMonad_majReader_majTrans_semigroupmajReadermajT($dictApply_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majReader_majTrans_semigroupmajReadermajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $functorReaderT1_1_0 = (object)["map" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Reader_Trans_mapReaderT']))(((($dictApply_0)->{'Functor0'})(null))->{'map'})];
  $applyReaderT1_1_0 = (object)["apply" => function($v_2) use ($dictApply_0) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($dictApply_0, $v_2) {
  $__num = \func_num_args();
  $__res = function($r_4) use ($dictApply_0, $v1_3, $v_2) {
  $__num = \func_num_args();
  $__res = ((($dictApply_0)->{'apply'})(($v_2)($r_4)))(($v1_3)($r_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_2) use ($functorReaderT1_1_0) {
  $__num = \func_num_args();
  $__res = $functorReaderT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = function($dictSemigroup_2) use ($applyReaderT1_1_0) {
  $__num = \func_num_args();
  $Functor0_3_2 = (($applyReaderT1_1_0)->{'Functor0'})(null);
  $__local_var_4_3 = ($dictSemigroup_2)->{'append'};
  $__res = (object)["append" => function($a_5) use ($Functor0_3_2, $__local_var_4_3, $applyReaderT1_1_0) {
  $__num = \func_num_args();
  $__res = function($b_6) use ($Functor0_3_2, $__local_var_4_3, $a_5, $applyReaderT1_1_0) {
  $__num = \func_num_args();
  $__res = ((($applyReaderT1_1_0)->{'apply'})(((($Functor0_3_2)->{'map'})($__local_var_4_3))($a_5)))($b_6);
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
$GLOBALS['Control_Monad_Reader_Trans_semigroupReaderT'] = __NAMESPACE__ . '\\majControl_majMonad_majReader_majTrans_semigroupmajReadermajT';

// Control_Monad_Reader_Trans_applicativeReaderT
function majControl_majMonad_majReader_majTrans_applicativemajReadermajT($dictApplicative_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majReader_majTrans_applicativemajReadermajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictApplicative_0)->{'Apply0'})(null);
  $functorReaderT1_2_1 = (object)["map" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Reader_Trans_mapReaderT']))(((($__local_var_1_0)->{'Functor0'})(null))->{'map'})];
  $applyReaderT1_1_0 = (object)["apply" => function($v_3) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v1_4) use ($__local_var_1_0, $v_3) {
  $__num = \func_num_args();
  $__res = function($r_5) use ($__local_var_1_0, $v1_4, $v_3) {
  $__num = \func_num_args();
  $__res = ((($__local_var_1_0)->{'apply'})(($v_3)($r_5)))(($v1_4)($r_5));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_3) use ($functorReaderT1_2_1) {
  $__num = \func_num_args();
  $__res = $functorReaderT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_2) {
  $__num = \func_num_args();
  $__res = $x_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Function_const']))(($dictApplicative_0)->{'pure'})), "Apply0" => function($_dollar___unused_2) use ($applyReaderT1_1_0) {
  $__num = \func_num_args();
  $__res = $applyReaderT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Reader_Trans_applicativeReaderT'] = __NAMESPACE__ . '\\majControl_majMonad_majReader_majTrans_applicativemajReadermajT';

// Control_Monad_Reader_Trans_monadReaderT
function majControl_majMonad_majReader_majTrans_monadmajReadermajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majReader_majTrans_monadmajReadermajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictMonad_0)->{'Applicative0'})(null);
  $__local_var_2_1 = (($__local_var_1_0)->{'Apply0'})(null);
  $functorReaderT1_3_2 = (object)["map" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Reader_Trans_mapReaderT']))(((($__local_var_2_1)->{'Functor0'})(null))->{'map'})];
  $applyReaderT1_2_1 = (object)["apply" => function($v_4) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = function($v1_5) use ($__local_var_2_1, $v_4) {
  $__num = \func_num_args();
  $__res = function($r_6) use ($__local_var_2_1, $v1_5, $v_4) {
  $__num = \func_num_args();
  $__res = ((($__local_var_2_1)->{'apply'})(($v_4)($r_6)))(($v1_5)($r_6));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_4) use ($functorReaderT1_3_2) {
  $__num = \func_num_args();
  $__res = $functorReaderT1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeReaderT1_1_0 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_3) {
  $__num = \func_num_args();
  $__res = $x_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Function_const']))(($__local_var_1_0)->{'pure'})), "Apply0" => function($_dollar___unused_3) use ($applyReaderT1_2_1) {
  $__num = \func_num_args();
  $__res = $applyReaderT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_2_5 = (($dictMonad_0)->{'Bind1'})(null);
  $__local_var_3_6 = (($__local_var_2_5)->{'Apply0'})(null);
  $functorReaderT1_4_7 = (object)["map" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Reader_Trans_mapReaderT']))(((($__local_var_3_6)->{'Functor0'})(null))->{'map'})];
  $applyReaderT1_3_6 = (object)["apply" => function($v_5) use ($__local_var_3_6) {
  $__num = \func_num_args();
  $__res = function($v1_6) use ($__local_var_3_6, $v_5) {
  $__num = \func_num_args();
  $__res = function($r_7) use ($__local_var_3_6, $v1_6, $v_5) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_6)->{'apply'})(($v_5)($r_7)))(($v1_6)($r_7));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_5) use ($functorReaderT1_4_7) {
  $__num = \func_num_args();
  $__res = $functorReaderT1_4_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $bindReaderT1_2_5 = (object)["bind" => function($v_4) use ($__local_var_2_5) {
  $__num = \func_num_args();
  $__res = function($k_5) use ($__local_var_2_5, $v_4) {
  $__num = \func_num_args();
  $__res = function($r_6) use ($__local_var_2_5, $k_5, $v_4) {
  $__num = \func_num_args();
  $__res = ((($__local_var_2_5)->{'bind'})(($v_4)($r_6)))(function($a_7) use ($k_5, $r_6) {
  $__num = \func_num_args();
  $__res = (($k_5)($a_7))($r_6);
  goto __end;;
  __end:
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
}, "Apply0" => function($_dollar___unused_4) use ($applyReaderT1_3_6) {
  $__num = \func_num_args();
  $__res = $applyReaderT1_3_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["Applicative0" => function($_dollar___unused_3) use ($applicativeReaderT1_1_0) {
  $__num = \func_num_args();
  $__res = $applicativeReaderT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_3) use ($bindReaderT1_2_5) {
  $__num = \func_num_args();
  $__res = $bindReaderT1_2_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Reader_Trans_monadReaderT'] = __NAMESPACE__ . '\\majControl_majMonad_majReader_majTrans_monadmajReadermajT';

// Control_Monad_Reader_Trans_monadAskReaderT
function majControl_majMonad_majReader_majTrans_monadmajAskmajReadermajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majReader_majTrans_monadmajAskmajReadermajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictMonad_0)->{'Applicative0'})(null);
  $__local_var_2_1 = (($__local_var_1_0)->{'Apply0'})(null);
  $functorReaderT1_3_2 = (object)["map" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Reader_Trans_mapReaderT']))(((($__local_var_2_1)->{'Functor0'})(null))->{'map'})];
  $applyReaderT1_2_1 = (object)["apply" => function($v_4) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = function($v1_5) use ($__local_var_2_1, $v_4) {
  $__num = \func_num_args();
  $__res = function($r_6) use ($__local_var_2_1, $v1_5, $v_4) {
  $__num = \func_num_args();
  $__res = ((($__local_var_2_1)->{'apply'})(($v_4)($r_6)))(($v1_5)($r_6));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_4) use ($functorReaderT1_3_2) {
  $__num = \func_num_args();
  $__res = $functorReaderT1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeReaderT1_1_0 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_3) {
  $__num = \func_num_args();
  $__res = $x_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Function_const']))(($__local_var_1_0)->{'pure'})), "Apply0" => function($_dollar___unused_3) use ($applyReaderT1_2_1) {
  $__num = \func_num_args();
  $__res = $applyReaderT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_2_5 = (($dictMonad_0)->{'Bind1'})(null);
  $__local_var_3_6 = (($__local_var_2_5)->{'Apply0'})(null);
  $functorReaderT1_4_7 = (object)["map" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Reader_Trans_mapReaderT']))(((($__local_var_3_6)->{'Functor0'})(null))->{'map'})];
  $applyReaderT1_3_6 = (object)["apply" => function($v_5) use ($__local_var_3_6) {
  $__num = \func_num_args();
  $__res = function($v1_6) use ($__local_var_3_6, $v_5) {
  $__num = \func_num_args();
  $__res = function($r_7) use ($__local_var_3_6, $v1_6, $v_5) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_6)->{'apply'})(($v_5)($r_7)))(($v1_6)($r_7));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_5) use ($functorReaderT1_4_7) {
  $__num = \func_num_args();
  $__res = $functorReaderT1_4_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $bindReaderT1_2_5 = (object)["bind" => function($v_4) use ($__local_var_2_5) {
  $__num = \func_num_args();
  $__res = function($k_5) use ($__local_var_2_5, $v_4) {
  $__num = \func_num_args();
  $__res = function($r_6) use ($__local_var_2_5, $k_5, $v_4) {
  $__num = \func_num_args();
  $__res = ((($__local_var_2_5)->{'bind'})(($v_4)($r_6)))(function($a_7) use ($k_5, $r_6) {
  $__num = \func_num_args();
  $__res = (($k_5)($a_7))($r_6);
  goto __end;;
  __end:
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
}, "Apply0" => function($_dollar___unused_4) use ($applyReaderT1_3_6) {
  $__num = \func_num_args();
  $__res = $applyReaderT1_3_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $monadReaderT1_1_0 = (object)["Applicative0" => function($_dollar___unused_3) use ($applicativeReaderT1_1_0) {
  $__num = \func_num_args();
  $__res = $applicativeReaderT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_3) use ($bindReaderT1_2_5) {
  $__num = \func_num_args();
  $__res = $bindReaderT1_2_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["ask" => ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'}, "Monad0" => function($_dollar___unused_2) use ($monadReaderT1_1_0) {
  $__num = \func_num_args();
  $__res = $monadReaderT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Reader_Trans_monadAskReaderT'] = __NAMESPACE__ . '\\majControl_majMonad_majReader_majTrans_monadmajAskmajReadermajT';

// Control_Monad_Reader_Trans_monadReaderReaderT
function majControl_majMonad_majReader_majTrans_monadmajReadermajReadermajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majReader_majTrans_monadmajReadermajReadermajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictMonad_0)->{'Applicative0'})(null);
  $__local_var_2_1 = (($__local_var_1_0)->{'Apply0'})(null);
  $functorReaderT1_3_2 = (object)["map" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Reader_Trans_mapReaderT']))(((($__local_var_2_1)->{'Functor0'})(null))->{'map'})];
  $applyReaderT1_2_1 = (object)["apply" => function($v_4) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = function($v1_5) use ($__local_var_2_1, $v_4) {
  $__num = \func_num_args();
  $__res = function($r_6) use ($__local_var_2_1, $v1_5, $v_4) {
  $__num = \func_num_args();
  $__res = ((($__local_var_2_1)->{'apply'})(($v_4)($r_6)))(($v1_5)($r_6));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_4) use ($functorReaderT1_3_2) {
  $__num = \func_num_args();
  $__res = $functorReaderT1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeReaderT1_1_0 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_3) {
  $__num = \func_num_args();
  $__res = $x_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Function_const']))(($__local_var_1_0)->{'pure'})), "Apply0" => function($_dollar___unused_3) use ($applyReaderT1_2_1) {
  $__num = \func_num_args();
  $__res = $applyReaderT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_2_5 = (($dictMonad_0)->{'Bind1'})(null);
  $__local_var_3_6 = (($__local_var_2_5)->{'Apply0'})(null);
  $functorReaderT1_4_7 = (object)["map" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Reader_Trans_mapReaderT']))(((($__local_var_3_6)->{'Functor0'})(null))->{'map'})];
  $applyReaderT1_3_6 = (object)["apply" => function($v_5) use ($__local_var_3_6) {
  $__num = \func_num_args();
  $__res = function($v1_6) use ($__local_var_3_6, $v_5) {
  $__num = \func_num_args();
  $__res = function($r_7) use ($__local_var_3_6, $v1_6, $v_5) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_6)->{'apply'})(($v_5)($r_7)))(($v1_6)($r_7));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_5) use ($functorReaderT1_4_7) {
  $__num = \func_num_args();
  $__res = $functorReaderT1_4_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $bindReaderT1_2_5 = (object)["bind" => function($v_4) use ($__local_var_2_5) {
  $__num = \func_num_args();
  $__res = function($k_5) use ($__local_var_2_5, $v_4) {
  $__num = \func_num_args();
  $__res = function($r_6) use ($__local_var_2_5, $k_5, $v_4) {
  $__num = \func_num_args();
  $__res = ((($__local_var_2_5)->{'bind'})(($v_4)($r_6)))(function($a_7) use ($k_5, $r_6) {
  $__num = \func_num_args();
  $__res = (($k_5)($a_7))($r_6);
  goto __end;;
  __end:
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
}, "Apply0" => function($_dollar___unused_4) use ($applyReaderT1_3_6) {
  $__num = \func_num_args();
  $__res = $applyReaderT1_3_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $monadReaderT1_1_0 = (object)["Applicative0" => function($_dollar___unused_3) use ($applicativeReaderT1_1_0) {
  $__num = \func_num_args();
  $__res = $applicativeReaderT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_3) use ($bindReaderT1_2_5) {
  $__num = \func_num_args();
  $__res = $bindReaderT1_2_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $monadAskReaderT1_1_0 = (object)["ask" => ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'}, "Monad0" => function($_dollar___unused_2) use ($monadReaderT1_1_0) {
  $__num = \func_num_args();
  $__res = $monadReaderT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["local" => $GLOBALS['Control_Monad_Reader_Trans_withReaderT'], "MonadAsk0" => function($_dollar___unused_2) use ($monadAskReaderT1_1_0) {
  $__num = \func_num_args();
  $__res = $monadAskReaderT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Reader_Trans_monadReaderReaderT'] = __NAMESPACE__ . '\\majControl_majMonad_majReader_majTrans_monadmajReadermajReadermajT';

// Control_Monad_Reader_Trans_monadContReaderT
function majControl_majMonad_majReader_majTrans_monadmajContmajReadermajT($dictMonadCont_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majReader_majTrans_monadmajContmajReadermajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictMonadCont_0)->{'Monad0'})(null);
  $__local_var_2_1 = (($__local_var_1_0)->{'Applicative0'})(null);
  $__local_var_3_2 = (($__local_var_2_1)->{'Apply0'})(null);
  $functorReaderT1_4_3 = (object)["map" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Reader_Trans_mapReaderT']))(((($__local_var_3_2)->{'Functor0'})(null))->{'map'})];
  $applyReaderT1_3_2 = (object)["apply" => function($v_5) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__res = function($v1_6) use ($__local_var_3_2, $v_5) {
  $__num = \func_num_args();
  $__res = function($r_7) use ($__local_var_3_2, $v1_6, $v_5) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_2)->{'apply'})(($v_5)($r_7)))(($v1_6)($r_7));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_5) use ($functorReaderT1_4_3) {
  $__num = \func_num_args();
  $__res = $functorReaderT1_4_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeReaderT1_2_1 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_4) {
  $__num = \func_num_args();
  $__res = $x_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Function_const']))(($__local_var_2_1)->{'pure'})), "Apply0" => function($_dollar___unused_4) use ($applyReaderT1_3_2) {
  $__num = \func_num_args();
  $__res = $applyReaderT1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_3_6 = (($__local_var_1_0)->{'Bind1'})(null);
  $__local_var_4_7 = (($__local_var_3_6)->{'Apply0'})(null);
  $functorReaderT1_5_8 = (object)["map" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Reader_Trans_mapReaderT']))(((($__local_var_4_7)->{'Functor0'})(null))->{'map'})];
  $applyReaderT1_4_7 = (object)["apply" => function($v_6) use ($__local_var_4_7) {
  $__num = \func_num_args();
  $__res = function($v1_7) use ($__local_var_4_7, $v_6) {
  $__num = \func_num_args();
  $__res = function($r_8) use ($__local_var_4_7, $v1_7, $v_6) {
  $__num = \func_num_args();
  $__res = ((($__local_var_4_7)->{'apply'})(($v_6)($r_8)))(($v1_7)($r_8));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_6) use ($functorReaderT1_5_8) {
  $__num = \func_num_args();
  $__res = $functorReaderT1_5_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $bindReaderT1_3_6 = (object)["bind" => function($v_5) use ($__local_var_3_6) {
  $__num = \func_num_args();
  $__res = function($k_6) use ($__local_var_3_6, $v_5) {
  $__num = \func_num_args();
  $__res = function($r_7) use ($__local_var_3_6, $k_6, $v_5) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_6)->{'bind'})(($v_5)($r_7)))(function($a_8) use ($k_6, $r_7) {
  $__num = \func_num_args();
  $__res = (($k_6)($a_8))($r_7);
  goto __end;;
  __end:
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
}, "Apply0" => function($_dollar___unused_5) use ($applyReaderT1_4_7) {
  $__num = \func_num_args();
  $__res = $applyReaderT1_4_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $monadReaderT1_1_0 = (object)["Applicative0" => function($_dollar___unused_4) use ($applicativeReaderT1_2_1) {
  $__num = \func_num_args();
  $__res = $applicativeReaderT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_4) use ($bindReaderT1_3_6) {
  $__num = \func_num_args();
  $__res = $bindReaderT1_3_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["callCC" => function($f_2) use ($dictMonadCont_0) {
  $__num = \func_num_args();
  $__res = function($r_3) use ($dictMonadCont_0, $f_2) {
  $__num = \func_num_args();
  $__res = (($dictMonadCont_0)->{'callCC'})(function($c_4) use ($f_2, $r_3) {
  $__num = \func_num_args();
  $__res = (($f_2)((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_5) {
  $__num = \func_num_args();
  $__res = $x_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Function_const']))($c_4))))($r_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad0" => function($_dollar___unused_2) use ($monadReaderT1_1_0) {
  $__num = \func_num_args();
  $__res = $monadReaderT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Reader_Trans_monadContReaderT'] = __NAMESPACE__ . '\\majControl_majMonad_majReader_majTrans_monadmajContmajReadermajT';

// Control_Monad_Reader_Trans_monadEffectReader
function majControl_majMonad_majReader_majTrans_monadmajEffectmajReader($dictMonadEffect_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majReader_majTrans_monadmajEffectmajReader';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Monad0_1_0 = (($dictMonadEffect_0)->{'Monad0'})(null);
  $__local_var_2_1 = (($Monad0_1_0)->{'Applicative0'})(null);
  $__local_var_3_2 = (($__local_var_2_1)->{'Apply0'})(null);
  $functorReaderT1_4_3 = (object)["map" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Reader_Trans_mapReaderT']))(((($__local_var_3_2)->{'Functor0'})(null))->{'map'})];
  $applyReaderT1_3_2 = (object)["apply" => function($v_5) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__res = function($v1_6) use ($__local_var_3_2, $v_5) {
  $__num = \func_num_args();
  $__res = function($r_7) use ($__local_var_3_2, $v1_6, $v_5) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_2)->{'apply'})(($v_5)($r_7)))(($v1_6)($r_7));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_5) use ($functorReaderT1_4_3) {
  $__num = \func_num_args();
  $__res = $functorReaderT1_4_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeReaderT1_2_1 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_4) {
  $__num = \func_num_args();
  $__res = $x_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Function_const']))(($__local_var_2_1)->{'pure'})), "Apply0" => function($_dollar___unused_4) use ($applyReaderT1_3_2) {
  $__num = \func_num_args();
  $__res = $applyReaderT1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_3_6 = (($Monad0_1_0)->{'Bind1'})(null);
  $__local_var_4_7 = (($__local_var_3_6)->{'Apply0'})(null);
  $functorReaderT1_5_8 = (object)["map" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Reader_Trans_mapReaderT']))(((($__local_var_4_7)->{'Functor0'})(null))->{'map'})];
  $applyReaderT1_4_7 = (object)["apply" => function($v_6) use ($__local_var_4_7) {
  $__num = \func_num_args();
  $__res = function($v1_7) use ($__local_var_4_7, $v_6) {
  $__num = \func_num_args();
  $__res = function($r_8) use ($__local_var_4_7, $v1_7, $v_6) {
  $__num = \func_num_args();
  $__res = ((($__local_var_4_7)->{'apply'})(($v_6)($r_8)))(($v1_7)($r_8));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_6) use ($functorReaderT1_5_8) {
  $__num = \func_num_args();
  $__res = $functorReaderT1_5_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $bindReaderT1_3_6 = (object)["bind" => function($v_5) use ($__local_var_3_6) {
  $__num = \func_num_args();
  $__res = function($k_6) use ($__local_var_3_6, $v_5) {
  $__num = \func_num_args();
  $__res = function($r_7) use ($__local_var_3_6, $k_6, $v_5) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_6)->{'bind'})(($v_5)($r_7)))(function($a_8) use ($k_6, $r_7) {
  $__num = \func_num_args();
  $__res = (($k_6)($a_8))($r_7);
  goto __end;;
  __end:
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
}, "Apply0" => function($_dollar___unused_5) use ($applyReaderT1_4_7) {
  $__num = \func_num_args();
  $__res = $applyReaderT1_4_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $monadReaderT1_2_1 = (object)["Applicative0" => function($_dollar___unused_4) use ($applicativeReaderT1_2_1) {
  $__num = \func_num_args();
  $__res = $applicativeReaderT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_4) use ($bindReaderT1_3_6) {
  $__num = \func_num_args();
  $__res = $bindReaderT1_3_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["liftEffect" => (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_3) {
  $__num = \func_num_args();
  $__res = $x_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Function_const'])))(($dictMonadEffect_0)->{'liftEffect'}), "Monad0" => function($_dollar___unused_3) use ($monadReaderT1_2_1) {
  $__num = \func_num_args();
  $__res = $monadReaderT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Reader_Trans_monadEffectReader'] = __NAMESPACE__ . '\\majControl_majMonad_majReader_majTrans_monadmajEffectmajReader';

// Control_Monad_Reader_Trans_monadRecReaderT
function majControl_majMonad_majReader_majTrans_monadmajRecmajReadermajT($dictMonadRec_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majReader_majTrans_monadmajRecmajReadermajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Monad0_1_0 = (($dictMonadRec_0)->{'Monad0'})(null);
  $Bind1_2_1 = (($Monad0_1_0)->{'Bind1'})(null);
  $pure_3_2 = ((($Monad0_1_0)->{'Applicative0'})(null))->{'pure'};
  $__local_var_4_3 = (($Monad0_1_0)->{'Applicative0'})(null);
  $__local_var_5_4 = (($__local_var_4_3)->{'Apply0'})(null);
  $functorReaderT1_6_5 = (object)["map" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Reader_Trans_mapReaderT']))(((($__local_var_5_4)->{'Functor0'})(null))->{'map'})];
  $applyReaderT1_5_4 = (object)["apply" => function($v_7) use ($__local_var_5_4) {
  $__num = \func_num_args();
  $__res = function($v1_8) use ($__local_var_5_4, $v_7) {
  $__num = \func_num_args();
  $__res = function($r_9) use ($__local_var_5_4, $v1_8, $v_7) {
  $__num = \func_num_args();
  $__res = ((($__local_var_5_4)->{'apply'})(($v_7)($r_9)))(($v1_8)($r_9));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_7) use ($functorReaderT1_6_5) {
  $__num = \func_num_args();
  $__res = $functorReaderT1_6_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeReaderT1_4_3 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_6) {
  $__num = \func_num_args();
  $__res = $x_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Function_const']))(($__local_var_4_3)->{'pure'})), "Apply0" => function($_dollar___unused_6) use ($applyReaderT1_5_4) {
  $__num = \func_num_args();
  $__res = $applyReaderT1_5_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_5_8 = (($Monad0_1_0)->{'Bind1'})(null);
  $__local_var_6_9 = (($__local_var_5_8)->{'Apply0'})(null);
  $functorReaderT1_7_10 = (object)["map" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Reader_Trans_mapReaderT']))(((($__local_var_6_9)->{'Functor0'})(null))->{'map'})];
  $applyReaderT1_6_9 = (object)["apply" => function($v_8) use ($__local_var_6_9) {
  $__num = \func_num_args();
  $__res = function($v1_9) use ($__local_var_6_9, $v_8) {
  $__num = \func_num_args();
  $__res = function($r_10) use ($__local_var_6_9, $v1_9, $v_8) {
  $__num = \func_num_args();
  $__res = ((($__local_var_6_9)->{'apply'})(($v_8)($r_10)))(($v1_9)($r_10));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_8) use ($functorReaderT1_7_10) {
  $__num = \func_num_args();
  $__res = $functorReaderT1_7_10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $bindReaderT1_5_8 = (object)["bind" => function($v_7) use ($__local_var_5_8) {
  $__num = \func_num_args();
  $__res = function($k_8) use ($__local_var_5_8, $v_7) {
  $__num = \func_num_args();
  $__res = function($r_9) use ($__local_var_5_8, $k_8, $v_7) {
  $__num = \func_num_args();
  $__res = ((($__local_var_5_8)->{'bind'})(($v_7)($r_9)))(function($a_10) use ($k_8, $r_9) {
  $__num = \func_num_args();
  $__res = (($k_8)($a_10))($r_9);
  goto __end;;
  __end:
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
}, "Apply0" => function($_dollar___unused_7) use ($applyReaderT1_6_9) {
  $__num = \func_num_args();
  $__res = $applyReaderT1_6_9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $monadReaderT1_4_3 = (object)["Applicative0" => function($_dollar___unused_6) use ($applicativeReaderT1_4_3) {
  $__num = \func_num_args();
  $__res = $applicativeReaderT1_4_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_6) use ($bindReaderT1_5_8) {
  $__num = \func_num_args();
  $__res = $bindReaderT1_5_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["tailRecM" => function($k_5) use ($Bind1_2_1, $dictMonadRec_0, $pure_3_2) {
  $__num = \func_num_args();
  $__res = function($a_6) use ($Bind1_2_1, $dictMonadRec_0, $k_5, $pure_3_2) {
  $__num = \func_num_args();
  $__res = function($r_7) use ($Bind1_2_1, $a_6, $dictMonadRec_0, $k_5, $pure_3_2) {
  $__num = \func_num_args();
  $__res = ((($dictMonadRec_0)->{'tailRecM'})(function($a_prime__8) use ($Bind1_2_1, $k_5, $pure_3_2, $r_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_2_1)->{'bind'})((($k_5)($a_prime__8))($r_7)))($pure_3_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($a_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad0" => function($_dollar___unused_5) use ($monadReaderT1_4_3) {
  $__num = \func_num_args();
  $__res = $monadReaderT1_4_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Reader_Trans_monadRecReaderT'] = __NAMESPACE__ . '\\majControl_majMonad_majReader_majTrans_monadmajRecmajReadermajT';

// Control_Monad_Reader_Trans_monadStateReaderT
function majControl_majMonad_majReader_majTrans_monadmajStatemajReadermajT($dictMonadState_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majReader_majTrans_monadmajStatemajReadermajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Monad0_1_0 = (($dictMonadState_0)->{'Monad0'})(null);
  $__local_var_2_1 = (($Monad0_1_0)->{'Applicative0'})(null);
  $__local_var_3_2 = (($__local_var_2_1)->{'Apply0'})(null);
  $functorReaderT1_4_3 = (object)["map" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Reader_Trans_mapReaderT']))(((($__local_var_3_2)->{'Functor0'})(null))->{'map'})];
  $applyReaderT1_3_2 = (object)["apply" => function($v_5) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__res = function($v1_6) use ($__local_var_3_2, $v_5) {
  $__num = \func_num_args();
  $__res = function($r_7) use ($__local_var_3_2, $v1_6, $v_5) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_2)->{'apply'})(($v_5)($r_7)))(($v1_6)($r_7));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_5) use ($functorReaderT1_4_3) {
  $__num = \func_num_args();
  $__res = $functorReaderT1_4_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeReaderT1_2_1 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_4) {
  $__num = \func_num_args();
  $__res = $x_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Function_const']))(($__local_var_2_1)->{'pure'})), "Apply0" => function($_dollar___unused_4) use ($applyReaderT1_3_2) {
  $__num = \func_num_args();
  $__res = $applyReaderT1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_3_6 = (($Monad0_1_0)->{'Bind1'})(null);
  $__local_var_4_7 = (($__local_var_3_6)->{'Apply0'})(null);
  $functorReaderT1_5_8 = (object)["map" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Reader_Trans_mapReaderT']))(((($__local_var_4_7)->{'Functor0'})(null))->{'map'})];
  $applyReaderT1_4_7 = (object)["apply" => function($v_6) use ($__local_var_4_7) {
  $__num = \func_num_args();
  $__res = function($v1_7) use ($__local_var_4_7, $v_6) {
  $__num = \func_num_args();
  $__res = function($r_8) use ($__local_var_4_7, $v1_7, $v_6) {
  $__num = \func_num_args();
  $__res = ((($__local_var_4_7)->{'apply'})(($v_6)($r_8)))(($v1_7)($r_8));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_6) use ($functorReaderT1_5_8) {
  $__num = \func_num_args();
  $__res = $functorReaderT1_5_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $bindReaderT1_3_6 = (object)["bind" => function($v_5) use ($__local_var_3_6) {
  $__num = \func_num_args();
  $__res = function($k_6) use ($__local_var_3_6, $v_5) {
  $__num = \func_num_args();
  $__res = function($r_7) use ($__local_var_3_6, $k_6, $v_5) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_6)->{'bind'})(($v_5)($r_7)))(function($a_8) use ($k_6, $r_7) {
  $__num = \func_num_args();
  $__res = (($k_6)($a_8))($r_7);
  goto __end;;
  __end:
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
}, "Apply0" => function($_dollar___unused_5) use ($applyReaderT1_4_7) {
  $__num = \func_num_args();
  $__res = $applyReaderT1_4_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $monadReaderT1_2_1 = (object)["Applicative0" => function($_dollar___unused_4) use ($applicativeReaderT1_2_1) {
  $__num = \func_num_args();
  $__res = $applicativeReaderT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_4) use ($bindReaderT1_3_6) {
  $__num = \func_num_args();
  $__res = $bindReaderT1_3_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["state" => (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_3) {
  $__num = \func_num_args();
  $__res = $x_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Function_const'])))(($dictMonadState_0)->{'state'}), "Monad0" => function($_dollar___unused_3) use ($monadReaderT1_2_1) {
  $__num = \func_num_args();
  $__res = $monadReaderT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Reader_Trans_monadStateReaderT'] = __NAMESPACE__ . '\\majControl_majMonad_majReader_majTrans_monadmajStatemajReadermajT';

// Control_Monad_Reader_Trans_monadTellReaderT
function majControl_majMonad_majReader_majTrans_monadmajTellmajReadermajT($dictMonadTell_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majReader_majTrans_monadmajTellmajReadermajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Monad1_1_0 = (($dictMonadTell_0)->{'Monad1'})(null);
  $Semigroup0_2_1 = (($dictMonadTell_0)->{'Semigroup0'})(null);
  $__local_var_3_2 = (($Monad1_1_0)->{'Applicative0'})(null);
  $__local_var_4_3 = (($__local_var_3_2)->{'Apply0'})(null);
  $functorReaderT1_5_4 = (object)["map" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Reader_Trans_mapReaderT']))(((($__local_var_4_3)->{'Functor0'})(null))->{'map'})];
  $applyReaderT1_4_3 = (object)["apply" => function($v_6) use ($__local_var_4_3) {
  $__num = \func_num_args();
  $__res = function($v1_7) use ($__local_var_4_3, $v_6) {
  $__num = \func_num_args();
  $__res = function($r_8) use ($__local_var_4_3, $v1_7, $v_6) {
  $__num = \func_num_args();
  $__res = ((($__local_var_4_3)->{'apply'})(($v_6)($r_8)))(($v1_7)($r_8));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_6) use ($functorReaderT1_5_4) {
  $__num = \func_num_args();
  $__res = $functorReaderT1_5_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeReaderT1_3_2 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_5) {
  $__num = \func_num_args();
  $__res = $x_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Function_const']))(($__local_var_3_2)->{'pure'})), "Apply0" => function($_dollar___unused_5) use ($applyReaderT1_4_3) {
  $__num = \func_num_args();
  $__res = $applyReaderT1_4_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_4_7 = (($Monad1_1_0)->{'Bind1'})(null);
  $__local_var_5_8 = (($__local_var_4_7)->{'Apply0'})(null);
  $functorReaderT1_6_9 = (object)["map" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Reader_Trans_mapReaderT']))(((($__local_var_5_8)->{'Functor0'})(null))->{'map'})];
  $applyReaderT1_5_8 = (object)["apply" => function($v_7) use ($__local_var_5_8) {
  $__num = \func_num_args();
  $__res = function($v1_8) use ($__local_var_5_8, $v_7) {
  $__num = \func_num_args();
  $__res = function($r_9) use ($__local_var_5_8, $v1_8, $v_7) {
  $__num = \func_num_args();
  $__res = ((($__local_var_5_8)->{'apply'})(($v_7)($r_9)))(($v1_8)($r_9));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_7) use ($functorReaderT1_6_9) {
  $__num = \func_num_args();
  $__res = $functorReaderT1_6_9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $bindReaderT1_4_7 = (object)["bind" => function($v_6) use ($__local_var_4_7) {
  $__num = \func_num_args();
  $__res = function($k_7) use ($__local_var_4_7, $v_6) {
  $__num = \func_num_args();
  $__res = function($r_8) use ($__local_var_4_7, $k_7, $v_6) {
  $__num = \func_num_args();
  $__res = ((($__local_var_4_7)->{'bind'})(($v_6)($r_8)))(function($a_9) use ($k_7, $r_8) {
  $__num = \func_num_args();
  $__res = (($k_7)($a_9))($r_8);
  goto __end;;
  __end:
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
}, "Apply0" => function($_dollar___unused_6) use ($applyReaderT1_5_8) {
  $__num = \func_num_args();
  $__res = $applyReaderT1_5_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $monadReaderT1_3_2 = (object)["Applicative0" => function($_dollar___unused_5) use ($applicativeReaderT1_3_2) {
  $__num = \func_num_args();
  $__res = $applicativeReaderT1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_5) use ($bindReaderT1_4_7) {
  $__num = \func_num_args();
  $__res = $bindReaderT1_4_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["tell" => (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_4) {
  $__num = \func_num_args();
  $__res = $x_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Function_const'])))(($dictMonadTell_0)->{'tell'}), "Semigroup0" => function($_dollar___unused_4) use ($Semigroup0_2_1) {
  $__num = \func_num_args();
  $__res = $Semigroup0_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad1" => function($_dollar___unused_4) use ($monadReaderT1_3_2) {
  $__num = \func_num_args();
  $__res = $monadReaderT1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Reader_Trans_monadTellReaderT'] = __NAMESPACE__ . '\\majControl_majMonad_majReader_majTrans_monadmajTellmajReadermajT';

// Control_Monad_Reader_Trans_monadWriterReaderT
function majControl_majMonad_majReader_majTrans_monadmajWritermajReadermajT($dictMonadWriter_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majReader_majTrans_monadmajWritermajReadermajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Monoid0_1_0 = (($dictMonadWriter_0)->{'Monoid0'})(null);
  $__local_var_2_1 = (($dictMonadWriter_0)->{'MonadTell1'})(null);
  $Monad1_3_2 = (($__local_var_2_1)->{'Monad1'})(null);
  $Semigroup0_4_3 = (($__local_var_2_1)->{'Semigroup0'})(null);
  $__local_var_5_4 = (($Monad1_3_2)->{'Applicative0'})(null);
  $__local_var_6_5 = (($__local_var_5_4)->{'Apply0'})(null);
  $functorReaderT1_7_6 = (object)["map" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Reader_Trans_mapReaderT']))(((($__local_var_6_5)->{'Functor0'})(null))->{'map'})];
  $applyReaderT1_6_5 = (object)["apply" => function($v_8) use ($__local_var_6_5) {
  $__num = \func_num_args();
  $__res = function($v1_9) use ($__local_var_6_5, $v_8) {
  $__num = \func_num_args();
  $__res = function($r_10) use ($__local_var_6_5, $v1_9, $v_8) {
  $__num = \func_num_args();
  $__res = ((($__local_var_6_5)->{'apply'})(($v_8)($r_10)))(($v1_9)($r_10));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_8) use ($functorReaderT1_7_6) {
  $__num = \func_num_args();
  $__res = $functorReaderT1_7_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeReaderT1_5_4 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_7) {
  $__num = \func_num_args();
  $__res = $x_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Function_const']))(($__local_var_5_4)->{'pure'})), "Apply0" => function($_dollar___unused_7) use ($applyReaderT1_6_5) {
  $__num = \func_num_args();
  $__res = $applyReaderT1_6_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_6_9 = (($Monad1_3_2)->{'Bind1'})(null);
  $__local_var_7_10 = (($__local_var_6_9)->{'Apply0'})(null);
  $functorReaderT1_8_11 = (object)["map" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Reader_Trans_mapReaderT']))(((($__local_var_7_10)->{'Functor0'})(null))->{'map'})];
  $applyReaderT1_7_10 = (object)["apply" => function($v_9) use ($__local_var_7_10) {
  $__num = \func_num_args();
  $__res = function($v1_10) use ($__local_var_7_10, $v_9) {
  $__num = \func_num_args();
  $__res = function($r_11) use ($__local_var_7_10, $v1_10, $v_9) {
  $__num = \func_num_args();
  $__res = ((($__local_var_7_10)->{'apply'})(($v_9)($r_11)))(($v1_10)($r_11));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_9) use ($functorReaderT1_8_11) {
  $__num = \func_num_args();
  $__res = $functorReaderT1_8_11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $bindReaderT1_6_9 = (object)["bind" => function($v_8) use ($__local_var_6_9) {
  $__num = \func_num_args();
  $__res = function($k_9) use ($__local_var_6_9, $v_8) {
  $__num = \func_num_args();
  $__res = function($r_10) use ($__local_var_6_9, $k_9, $v_8) {
  $__num = \func_num_args();
  $__res = ((($__local_var_6_9)->{'bind'})(($v_8)($r_10)))(function($a_11) use ($k_9, $r_10) {
  $__num = \func_num_args();
  $__res = (($k_9)($a_11))($r_10);
  goto __end;;
  __end:
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
}, "Apply0" => function($_dollar___unused_8) use ($applyReaderT1_7_10) {
  $__num = \func_num_args();
  $__res = $applyReaderT1_7_10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $monadReaderT1_5_4 = (object)["Applicative0" => function($_dollar___unused_7) use ($applicativeReaderT1_5_4) {
  $__num = \func_num_args();
  $__res = $applicativeReaderT1_5_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_7) use ($bindReaderT1_6_9) {
  $__num = \func_num_args();
  $__res = $bindReaderT1_6_9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $monadTellReaderT1_2_1 = (object)["tell" => (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_6) {
  $__num = \func_num_args();
  $__res = $x_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Function_const'])))(($__local_var_2_1)->{'tell'}), "Semigroup0" => function($_dollar___unused_6) use ($Semigroup0_4_3) {
  $__num = \func_num_args();
  $__res = $Semigroup0_4_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad1" => function($_dollar___unused_6) use ($monadReaderT1_5_4) {
  $__num = \func_num_args();
  $__res = $monadReaderT1_5_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["listen" => ($GLOBALS['Control_Monad_Reader_Trans_mapReaderT'])(($dictMonadWriter_0)->{'listen'}), "pass" => ($GLOBALS['Control_Monad_Reader_Trans_mapReaderT'])(($dictMonadWriter_0)->{'pass'}), "Monoid0" => function($_dollar___unused_3) use ($Monoid0_1_0) {
  $__num = \func_num_args();
  $__res = $Monoid0_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "MonadTell1" => function($_dollar___unused_3) use ($monadTellReaderT1_2_1) {
  $__num = \func_num_args();
  $__res = $monadTellReaderT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Reader_Trans_monadWriterReaderT'] = __NAMESPACE__ . '\\majControl_majMonad_majReader_majTrans_monadmajWritermajReadermajT';

// Control_Monad_Reader_Trans_monadThrowReaderT
function majControl_majMonad_majReader_majTrans_monadmajThrowmajReadermajT($dictMonadThrow_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majReader_majTrans_monadmajThrowmajReadermajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Monad0_1_0 = (($dictMonadThrow_0)->{'Monad0'})(null);
  $__local_var_2_1 = (($Monad0_1_0)->{'Applicative0'})(null);
  $__local_var_3_2 = (($__local_var_2_1)->{'Apply0'})(null);
  $functorReaderT1_4_3 = (object)["map" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Reader_Trans_mapReaderT']))(((($__local_var_3_2)->{'Functor0'})(null))->{'map'})];
  $applyReaderT1_3_2 = (object)["apply" => function($v_5) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__res = function($v1_6) use ($__local_var_3_2, $v_5) {
  $__num = \func_num_args();
  $__res = function($r_7) use ($__local_var_3_2, $v1_6, $v_5) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_2)->{'apply'})(($v_5)($r_7)))(($v1_6)($r_7));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_5) use ($functorReaderT1_4_3) {
  $__num = \func_num_args();
  $__res = $functorReaderT1_4_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeReaderT1_2_1 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_4) {
  $__num = \func_num_args();
  $__res = $x_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Function_const']))(($__local_var_2_1)->{'pure'})), "Apply0" => function($_dollar___unused_4) use ($applyReaderT1_3_2) {
  $__num = \func_num_args();
  $__res = $applyReaderT1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_3_6 = (($Monad0_1_0)->{'Bind1'})(null);
  $__local_var_4_7 = (($__local_var_3_6)->{'Apply0'})(null);
  $functorReaderT1_5_8 = (object)["map" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Reader_Trans_mapReaderT']))(((($__local_var_4_7)->{'Functor0'})(null))->{'map'})];
  $applyReaderT1_4_7 = (object)["apply" => function($v_6) use ($__local_var_4_7) {
  $__num = \func_num_args();
  $__res = function($v1_7) use ($__local_var_4_7, $v_6) {
  $__num = \func_num_args();
  $__res = function($r_8) use ($__local_var_4_7, $v1_7, $v_6) {
  $__num = \func_num_args();
  $__res = ((($__local_var_4_7)->{'apply'})(($v_6)($r_8)))(($v1_7)($r_8));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_6) use ($functorReaderT1_5_8) {
  $__num = \func_num_args();
  $__res = $functorReaderT1_5_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $bindReaderT1_3_6 = (object)["bind" => function($v_5) use ($__local_var_3_6) {
  $__num = \func_num_args();
  $__res = function($k_6) use ($__local_var_3_6, $v_5) {
  $__num = \func_num_args();
  $__res = function($r_7) use ($__local_var_3_6, $k_6, $v_5) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_6)->{'bind'})(($v_5)($r_7)))(function($a_8) use ($k_6, $r_7) {
  $__num = \func_num_args();
  $__res = (($k_6)($a_8))($r_7);
  goto __end;;
  __end:
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
}, "Apply0" => function($_dollar___unused_5) use ($applyReaderT1_4_7) {
  $__num = \func_num_args();
  $__res = $applyReaderT1_4_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $monadReaderT1_2_1 = (object)["Applicative0" => function($_dollar___unused_4) use ($applicativeReaderT1_2_1) {
  $__num = \func_num_args();
  $__res = $applicativeReaderT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_4) use ($bindReaderT1_3_6) {
  $__num = \func_num_args();
  $__res = $bindReaderT1_3_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["throwError" => (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_3) {
  $__num = \func_num_args();
  $__res = $x_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Function_const'])))(($dictMonadThrow_0)->{'throwError'}), "Monad0" => function($_dollar___unused_3) use ($monadReaderT1_2_1) {
  $__num = \func_num_args();
  $__res = $monadReaderT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Reader_Trans_monadThrowReaderT'] = __NAMESPACE__ . '\\majControl_majMonad_majReader_majTrans_monadmajThrowmajReadermajT';

// Control_Monad_Reader_Trans_monadErrorReaderT
function majControl_majMonad_majReader_majTrans_monadmajErrormajReadermajT($dictMonadError_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majReader_majTrans_monadmajErrormajReadermajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictMonadError_0)->{'MonadThrow0'})(null);
  $Monad0_2_1 = (($__local_var_1_0)->{'Monad0'})(null);
  $__local_var_3_2 = (($Monad0_2_1)->{'Applicative0'})(null);
  $__local_var_4_3 = (($__local_var_3_2)->{'Apply0'})(null);
  $functorReaderT1_5_4 = (object)["map" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Reader_Trans_mapReaderT']))(((($__local_var_4_3)->{'Functor0'})(null))->{'map'})];
  $applyReaderT1_4_3 = (object)["apply" => function($v_6) use ($__local_var_4_3) {
  $__num = \func_num_args();
  $__res = function($v1_7) use ($__local_var_4_3, $v_6) {
  $__num = \func_num_args();
  $__res = function($r_8) use ($__local_var_4_3, $v1_7, $v_6) {
  $__num = \func_num_args();
  $__res = ((($__local_var_4_3)->{'apply'})(($v_6)($r_8)))(($v1_7)($r_8));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_6) use ($functorReaderT1_5_4) {
  $__num = \func_num_args();
  $__res = $functorReaderT1_5_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeReaderT1_3_2 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_5) {
  $__num = \func_num_args();
  $__res = $x_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Function_const']))(($__local_var_3_2)->{'pure'})), "Apply0" => function($_dollar___unused_5) use ($applyReaderT1_4_3) {
  $__num = \func_num_args();
  $__res = $applyReaderT1_4_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_4_7 = (($Monad0_2_1)->{'Bind1'})(null);
  $__local_var_5_8 = (($__local_var_4_7)->{'Apply0'})(null);
  $functorReaderT1_6_9 = (object)["map" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Reader_Trans_mapReaderT']))(((($__local_var_5_8)->{'Functor0'})(null))->{'map'})];
  $applyReaderT1_5_8 = (object)["apply" => function($v_7) use ($__local_var_5_8) {
  $__num = \func_num_args();
  $__res = function($v1_8) use ($__local_var_5_8, $v_7) {
  $__num = \func_num_args();
  $__res = function($r_9) use ($__local_var_5_8, $v1_8, $v_7) {
  $__num = \func_num_args();
  $__res = ((($__local_var_5_8)->{'apply'})(($v_7)($r_9)))(($v1_8)($r_9));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_7) use ($functorReaderT1_6_9) {
  $__num = \func_num_args();
  $__res = $functorReaderT1_6_9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $bindReaderT1_4_7 = (object)["bind" => function($v_6) use ($__local_var_4_7) {
  $__num = \func_num_args();
  $__res = function($k_7) use ($__local_var_4_7, $v_6) {
  $__num = \func_num_args();
  $__res = function($r_8) use ($__local_var_4_7, $k_7, $v_6) {
  $__num = \func_num_args();
  $__res = ((($__local_var_4_7)->{'bind'})(($v_6)($r_8)))(function($a_9) use ($k_7, $r_8) {
  $__num = \func_num_args();
  $__res = (($k_7)($a_9))($r_8);
  goto __end;;
  __end:
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
}, "Apply0" => function($_dollar___unused_6) use ($applyReaderT1_5_8) {
  $__num = \func_num_args();
  $__res = $applyReaderT1_5_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $monadReaderT1_3_2 = (object)["Applicative0" => function($_dollar___unused_5) use ($applicativeReaderT1_3_2) {
  $__num = \func_num_args();
  $__res = $applicativeReaderT1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_5) use ($bindReaderT1_4_7) {
  $__num = \func_num_args();
  $__res = $bindReaderT1_4_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $monadThrowReaderT1_1_0 = (object)["throwError" => (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_4) {
  $__num = \func_num_args();
  $__res = $x_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Function_const'])))(($__local_var_1_0)->{'throwError'}), "Monad0" => function($_dollar___unused_4) use ($monadReaderT1_3_2) {
  $__num = \func_num_args();
  $__res = $monadReaderT1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["catchError" => function($v_2) use ($dictMonadError_0) {
  $__num = \func_num_args();
  $__res = function($h_3) use ($dictMonadError_0, $v_2) {
  $__num = \func_num_args();
  $__res = function($r_4) use ($dictMonadError_0, $h_3, $v_2) {
  $__num = \func_num_args();
  $__res = ((($dictMonadError_0)->{'catchError'})(($v_2)($r_4)))(function($e_5) use ($h_3, $r_4) {
  $__num = \func_num_args();
  $__res = (($h_3)($e_5))($r_4);
  goto __end;;
  __end:
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
}, "MonadThrow0" => function($_dollar___unused_2) use ($monadThrowReaderT1_1_0) {
  $__num = \func_num_args();
  $__res = $monadThrowReaderT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Reader_Trans_monadErrorReaderT'] = __NAMESPACE__ . '\\majControl_majMonad_majReader_majTrans_monadmajErrormajReadermajT';

// Control_Monad_Reader_Trans_monadSTReaderT
function majControl_majMonad_majReader_majTrans_monadmajSmajTmajReadermajT($dictMonadST_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majReader_majTrans_monadmajSmajTmajReadermajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Monad0_1_0 = (($dictMonadST_0)->{'Monad0'})(null);
  $__local_var_2_1 = (($Monad0_1_0)->{'Applicative0'})(null);
  $__local_var_3_2 = (($__local_var_2_1)->{'Apply0'})(null);
  $functorReaderT1_4_3 = (object)["map" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Reader_Trans_mapReaderT']))(((($__local_var_3_2)->{'Functor0'})(null))->{'map'})];
  $applyReaderT1_3_2 = (object)["apply" => function($v_5) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__res = function($v1_6) use ($__local_var_3_2, $v_5) {
  $__num = \func_num_args();
  $__res = function($r_7) use ($__local_var_3_2, $v1_6, $v_5) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_2)->{'apply'})(($v_5)($r_7)))(($v1_6)($r_7));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_5) use ($functorReaderT1_4_3) {
  $__num = \func_num_args();
  $__res = $functorReaderT1_4_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeReaderT1_2_1 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_4) {
  $__num = \func_num_args();
  $__res = $x_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Function_const']))(($__local_var_2_1)->{'pure'})), "Apply0" => function($_dollar___unused_4) use ($applyReaderT1_3_2) {
  $__num = \func_num_args();
  $__res = $applyReaderT1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_3_6 = (($Monad0_1_0)->{'Bind1'})(null);
  $__local_var_4_7 = (($__local_var_3_6)->{'Apply0'})(null);
  $functorReaderT1_5_8 = (object)["map" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Reader_Trans_mapReaderT']))(((($__local_var_4_7)->{'Functor0'})(null))->{'map'})];
  $applyReaderT1_4_7 = (object)["apply" => function($v_6) use ($__local_var_4_7) {
  $__num = \func_num_args();
  $__res = function($v1_7) use ($__local_var_4_7, $v_6) {
  $__num = \func_num_args();
  $__res = function($r_8) use ($__local_var_4_7, $v1_7, $v_6) {
  $__num = \func_num_args();
  $__res = ((($__local_var_4_7)->{'apply'})(($v_6)($r_8)))(($v1_7)($r_8));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_6) use ($functorReaderT1_5_8) {
  $__num = \func_num_args();
  $__res = $functorReaderT1_5_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $bindReaderT1_3_6 = (object)["bind" => function($v_5) use ($__local_var_3_6) {
  $__num = \func_num_args();
  $__res = function($k_6) use ($__local_var_3_6, $v_5) {
  $__num = \func_num_args();
  $__res = function($r_7) use ($__local_var_3_6, $k_6, $v_5) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_6)->{'bind'})(($v_5)($r_7)))(function($a_8) use ($k_6, $r_7) {
  $__num = \func_num_args();
  $__res = (($k_6)($a_8))($r_7);
  goto __end;;
  __end:
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
}, "Apply0" => function($_dollar___unused_5) use ($applyReaderT1_4_7) {
  $__num = \func_num_args();
  $__res = $applyReaderT1_4_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $monadReaderT1_2_1 = (object)["Applicative0" => function($_dollar___unused_4) use ($applicativeReaderT1_2_1) {
  $__num = \func_num_args();
  $__res = $applicativeReaderT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_4) use ($bindReaderT1_3_6) {
  $__num = \func_num_args();
  $__res = $bindReaderT1_3_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["liftST" => (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_3) {
  $__num = \func_num_args();
  $__res = $x_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Function_const'])))(($dictMonadST_0)->{'liftST'}), "Monad0" => function($_dollar___unused_3) use ($monadReaderT1_2_1) {
  $__num = \func_num_args();
  $__res = $monadReaderT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Reader_Trans_monadSTReaderT'] = __NAMESPACE__ . '\\majControl_majMonad_majReader_majTrans_monadmajSmajTmajReadermajT';

// Control_Monad_Reader_Trans_monoidReaderT
function majControl_majMonad_majReader_majTrans_monoidmajReadermajT($dictApplicative_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majReader_majTrans_monoidmajReadermajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictApplicative_0)->{'Apply0'})(null);
  $functorReaderT1_2_1 = (object)["map" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Reader_Trans_mapReaderT']))(((($__local_var_1_0)->{'Functor0'})(null))->{'map'})];
  $applyReaderT1_1_0 = (object)["apply" => function($v_3) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v1_4) use ($__local_var_1_0, $v_3) {
  $__num = \func_num_args();
  $__res = function($r_5) use ($__local_var_1_0, $v1_4, $v_3) {
  $__num = \func_num_args();
  $__res = ((($__local_var_1_0)->{'apply'})(($v_3)($r_5)))(($v1_4)($r_5));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_3) use ($functorReaderT1_2_1) {
  $__num = \func_num_args();
  $__res = $functorReaderT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeReaderT1_1_0 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_2) {
  $__num = \func_num_args();
  $__res = $x_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Function_const']))(($dictApplicative_0)->{'pure'})), "Apply0" => function($_dollar___unused_2) use ($applyReaderT1_1_0) {
  $__num = \func_num_args();
  $__res = $applyReaderT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_2_4 = (($dictApplicative_0)->{'Apply0'})(null);
  $functorReaderT1_3_5 = (object)["map" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Reader_Trans_mapReaderT']))(((($__local_var_2_4)->{'Functor0'})(null))->{'map'})];
  $applyReaderT1_3_5 = (object)["apply" => function($v_4) use ($__local_var_2_4) {
  $__num = \func_num_args();
  $__res = function($v1_5) use ($__local_var_2_4, $v_4) {
  $__num = \func_num_args();
  $__res = function($r_6) use ($__local_var_2_4, $v1_5, $v_4) {
  $__num = \func_num_args();
  $__res = ((($__local_var_2_4)->{'apply'})(($v_4)($r_6)))(($v1_5)($r_6));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_4) use ($functorReaderT1_3_5) {
  $__num = \func_num_args();
  $__res = $functorReaderT1_3_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = function($dictMonoid_4) use ($applicativeReaderT1_1_0, $applyReaderT1_3_5) {
  $__num = \func_num_args();
  $Functor0_5_7 = (($applyReaderT1_3_5)->{'Functor0'})(null);
  $__local_var_6_8 = ((($dictMonoid_4)->{'Semigroup0'})(null))->{'append'};
  $semigroupReaderT2_5_7 = (object)["append" => function($a_7) use ($Functor0_5_7, $__local_var_6_8, $applyReaderT1_3_5) {
  $__num = \func_num_args();
  $__res = function($b_8) use ($Functor0_5_7, $__local_var_6_8, $a_7, $applyReaderT1_3_5) {
  $__num = \func_num_args();
  $__res = ((($applyReaderT1_3_5)->{'apply'})(((($Functor0_5_7)->{'map'})($__local_var_6_8))($a_7)))($b_8);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["mempty" => (($applicativeReaderT1_1_0)->{'pure'})(($dictMonoid_4)->{'mempty'}), "Semigroup0" => function($_dollar___unused_6) use ($semigroupReaderT2_5_7) {
  $__num = \func_num_args();
  $__res = $semigroupReaderT2_5_7;
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
$GLOBALS['Control_Monad_Reader_Trans_monoidReaderT'] = __NAMESPACE__ . '\\majControl_majMonad_majReader_majTrans_monoidmajReadermajT';

// Control_Monad_Reader_Trans_altReaderT
function majControl_majMonad_majReader_majTrans_altmajReadermajT($dictAlt_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majReader_majTrans_altmajReadermajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $functorReaderT1_1_0 = (object)["map" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Reader_Trans_mapReaderT']))(((($dictAlt_0)->{'Functor0'})(null))->{'map'})];
  $__res = (object)["alt" => function($v_2) use ($dictAlt_0) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($dictAlt_0, $v_2) {
  $__num = \func_num_args();
  $__res = function($r_4) use ($dictAlt_0, $v1_3, $v_2) {
  $__num = \func_num_args();
  $__res = ((($dictAlt_0)->{'alt'})(($v_2)($r_4)))(($v1_3)($r_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_2) use ($functorReaderT1_1_0) {
  $__num = \func_num_args();
  $__res = $functorReaderT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Reader_Trans_altReaderT'] = __NAMESPACE__ . '\\majControl_majMonad_majReader_majTrans_altmajReadermajT';

// Control_Monad_Reader_Trans_plusReaderT
function majControl_majMonad_majReader_majTrans_plusmajReadermajT($dictPlus_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majReader_majTrans_plusmajReadermajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictPlus_0)->{'Alt0'})(null);
  $functorReaderT1_2_1 = (object)["map" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Reader_Trans_mapReaderT']))(((($__local_var_1_0)->{'Functor0'})(null))->{'map'})];
  $altReaderT1_1_0 = (object)["alt" => function($v_3) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v1_4) use ($__local_var_1_0, $v_3) {
  $__num = \func_num_args();
  $__res = function($r_5) use ($__local_var_1_0, $v1_4, $v_3) {
  $__num = \func_num_args();
  $__res = ((($__local_var_1_0)->{'alt'})(($v_3)($r_5)))(($v1_4)($r_5));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_3) use ($functorReaderT1_2_1) {
  $__num = \func_num_args();
  $__res = $functorReaderT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_2_3 = ($dictPlus_0)->{'empty'};
  $__res = (object)["empty" => function($v_3) use ($__local_var_2_3) {
  $__num = \func_num_args();
  $__res = $__local_var_2_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Alt0" => function($_dollar___unused_2) use ($altReaderT1_1_0) {
  $__num = \func_num_args();
  $__res = $altReaderT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Reader_Trans_plusReaderT'] = __NAMESPACE__ . '\\majControl_majMonad_majReader_majTrans_plusmajReadermajT';

// Control_Monad_Reader_Trans_alternativeReaderT
function majControl_majMonad_majReader_majTrans_alternativemajReadermajT($dictAlternative_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majReader_majTrans_alternativemajReadermajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictAlternative_0)->{'Applicative0'})(null);
  $__local_var_2_1 = (($__local_var_1_0)->{'Apply0'})(null);
  $functorReaderT1_3_2 = (object)["map" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Reader_Trans_mapReaderT']))(((($__local_var_2_1)->{'Functor0'})(null))->{'map'})];
  $applyReaderT1_2_1 = (object)["apply" => function($v_4) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = function($v1_5) use ($__local_var_2_1, $v_4) {
  $__num = \func_num_args();
  $__res = function($r_6) use ($__local_var_2_1, $v1_5, $v_4) {
  $__num = \func_num_args();
  $__res = ((($__local_var_2_1)->{'apply'})(($v_4)($r_6)))(($v1_5)($r_6));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_4) use ($functorReaderT1_3_2) {
  $__num = \func_num_args();
  $__res = $functorReaderT1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeReaderT1_1_0 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_3) {
  $__num = \func_num_args();
  $__res = $x_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Function_const']))(($__local_var_1_0)->{'pure'})), "Apply0" => function($_dollar___unused_3) use ($applyReaderT1_2_1) {
  $__num = \func_num_args();
  $__res = $applyReaderT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_2_5 = (($dictAlternative_0)->{'Plus1'})(null);
  $__local_var_3_6 = (($__local_var_2_5)->{'Alt0'})(null);
  $functorReaderT1_4_7 = (object)["map" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Reader_Trans_mapReaderT']))(((($__local_var_3_6)->{'Functor0'})(null))->{'map'})];
  $altReaderT1_3_6 = (object)["alt" => function($v_5) use ($__local_var_3_6) {
  $__num = \func_num_args();
  $__res = function($v1_6) use ($__local_var_3_6, $v_5) {
  $__num = \func_num_args();
  $__res = function($r_7) use ($__local_var_3_6, $v1_6, $v_5) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_6)->{'alt'})(($v_5)($r_7)))(($v1_6)($r_7));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_5) use ($functorReaderT1_4_7) {
  $__num = \func_num_args();
  $__res = $functorReaderT1_4_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_4_9 = ($__local_var_2_5)->{'empty'};
  $plusReaderT1_2_5 = (object)["empty" => function($v_5) use ($__local_var_4_9) {
  $__num = \func_num_args();
  $__res = $__local_var_4_9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Alt0" => function($_dollar___unused_4) use ($altReaderT1_3_6) {
  $__num = \func_num_args();
  $__res = $altReaderT1_3_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["Applicative0" => function($_dollar___unused_3) use ($applicativeReaderT1_1_0) {
  $__num = \func_num_args();
  $__res = $applicativeReaderT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Plus1" => function($_dollar___unused_3) use ($plusReaderT1_2_5) {
  $__num = \func_num_args();
  $__res = $plusReaderT1_2_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Reader_Trans_alternativeReaderT'] = __NAMESPACE__ . '\\majControl_majMonad_majReader_majTrans_alternativemajReadermajT';

// Control_Monad_Reader_Trans_monadPlusReaderT
function majControl_majMonad_majReader_majTrans_monadmajPlusmajReadermajT($dictMonadPlus_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majReader_majTrans_monadmajPlusmajReadermajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictMonadPlus_0)->{'Monad0'})(null);
  $__local_var_2_1 = (($__local_var_1_0)->{'Applicative0'})(null);
  $__local_var_3_2 = (($__local_var_2_1)->{'Apply0'})(null);
  $functorReaderT1_4_3 = (object)["map" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Reader_Trans_mapReaderT']))(((($__local_var_3_2)->{'Functor0'})(null))->{'map'})];
  $applyReaderT1_3_2 = (object)["apply" => function($v_5) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__res = function($v1_6) use ($__local_var_3_2, $v_5) {
  $__num = \func_num_args();
  $__res = function($r_7) use ($__local_var_3_2, $v1_6, $v_5) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_2)->{'apply'})(($v_5)($r_7)))(($v1_6)($r_7));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_5) use ($functorReaderT1_4_3) {
  $__num = \func_num_args();
  $__res = $functorReaderT1_4_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeReaderT1_2_1 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_4) {
  $__num = \func_num_args();
  $__res = $x_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Function_const']))(($__local_var_2_1)->{'pure'})), "Apply0" => function($_dollar___unused_4) use ($applyReaderT1_3_2) {
  $__num = \func_num_args();
  $__res = $applyReaderT1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_3_6 = (($__local_var_1_0)->{'Bind1'})(null);
  $__local_var_4_7 = (($__local_var_3_6)->{'Apply0'})(null);
  $functorReaderT1_5_8 = (object)["map" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Reader_Trans_mapReaderT']))(((($__local_var_4_7)->{'Functor0'})(null))->{'map'})];
  $applyReaderT1_4_7 = (object)["apply" => function($v_6) use ($__local_var_4_7) {
  $__num = \func_num_args();
  $__res = function($v1_7) use ($__local_var_4_7, $v_6) {
  $__num = \func_num_args();
  $__res = function($r_8) use ($__local_var_4_7, $v1_7, $v_6) {
  $__num = \func_num_args();
  $__res = ((($__local_var_4_7)->{'apply'})(($v_6)($r_8)))(($v1_7)($r_8));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_6) use ($functorReaderT1_5_8) {
  $__num = \func_num_args();
  $__res = $functorReaderT1_5_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $bindReaderT1_3_6 = (object)["bind" => function($v_5) use ($__local_var_3_6) {
  $__num = \func_num_args();
  $__res = function($k_6) use ($__local_var_3_6, $v_5) {
  $__num = \func_num_args();
  $__res = function($r_7) use ($__local_var_3_6, $k_6, $v_5) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_6)->{'bind'})(($v_5)($r_7)))(function($a_8) use ($k_6, $r_7) {
  $__num = \func_num_args();
  $__res = (($k_6)($a_8))($r_7);
  goto __end;;
  __end:
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
}, "Apply0" => function($_dollar___unused_5) use ($applyReaderT1_4_7) {
  $__num = \func_num_args();
  $__res = $applyReaderT1_4_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $monadReaderT1_1_0 = (object)["Applicative0" => function($_dollar___unused_4) use ($applicativeReaderT1_2_1) {
  $__num = \func_num_args();
  $__res = $applicativeReaderT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_4) use ($bindReaderT1_3_6) {
  $__num = \func_num_args();
  $__res = $bindReaderT1_3_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_2_12 = (($dictMonadPlus_0)->{'Alternative1'})(null);
  $__local_var_3_13 = (($__local_var_2_12)->{'Applicative0'})(null);
  $__local_var_4_14 = (($__local_var_3_13)->{'Apply0'})(null);
  $functorReaderT1_5_15 = (object)["map" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Reader_Trans_mapReaderT']))(((($__local_var_4_14)->{'Functor0'})(null))->{'map'})];
  $applyReaderT1_4_14 = (object)["apply" => function($v_6) use ($__local_var_4_14) {
  $__num = \func_num_args();
  $__res = function($v1_7) use ($__local_var_4_14, $v_6) {
  $__num = \func_num_args();
  $__res = function($r_8) use ($__local_var_4_14, $v1_7, $v_6) {
  $__num = \func_num_args();
  $__res = ((($__local_var_4_14)->{'apply'})(($v_6)($r_8)))(($v1_7)($r_8));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_6) use ($functorReaderT1_5_15) {
  $__num = \func_num_args();
  $__res = $functorReaderT1_5_15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeReaderT1_3_13 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_5) {
  $__num = \func_num_args();
  $__res = $x_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Function_const']))(($__local_var_3_13)->{'pure'})), "Apply0" => function($_dollar___unused_5) use ($applyReaderT1_4_14) {
  $__num = \func_num_args();
  $__res = $applyReaderT1_4_14;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_4_18 = (($__local_var_2_12)->{'Plus1'})(null);
  $__local_var_5_19 = (($__local_var_4_18)->{'Alt0'})(null);
  $functorReaderT1_6_20 = (object)["map" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Reader_Trans_mapReaderT']))(((($__local_var_5_19)->{'Functor0'})(null))->{'map'})];
  $altReaderT1_5_19 = (object)["alt" => function($v_7) use ($__local_var_5_19) {
  $__num = \func_num_args();
  $__res = function($v1_8) use ($__local_var_5_19, $v_7) {
  $__num = \func_num_args();
  $__res = function($r_9) use ($__local_var_5_19, $v1_8, $v_7) {
  $__num = \func_num_args();
  $__res = ((($__local_var_5_19)->{'alt'})(($v_7)($r_9)))(($v1_8)($r_9));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_7) use ($functorReaderT1_6_20) {
  $__num = \func_num_args();
  $__res = $functorReaderT1_6_20;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_6_22 = ($__local_var_4_18)->{'empty'};
  $plusReaderT1_4_18 = (object)["empty" => function($v_7) use ($__local_var_6_22) {
  $__num = \func_num_args();
  $__res = $__local_var_6_22;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Alt0" => function($_dollar___unused_6) use ($altReaderT1_5_19) {
  $__num = \func_num_args();
  $__res = $altReaderT1_5_19;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $alternativeReaderT1_2_12 = (object)["Applicative0" => function($_dollar___unused_5) use ($applicativeReaderT1_3_13) {
  $__num = \func_num_args();
  $__res = $applicativeReaderT1_3_13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Plus1" => function($_dollar___unused_5) use ($plusReaderT1_4_18) {
  $__num = \func_num_args();
  $__res = $plusReaderT1_4_18;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["Monad0" => function($_dollar___unused_3) use ($monadReaderT1_1_0) {
  $__num = \func_num_args();
  $__res = $monadReaderT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Alternative1" => function($_dollar___unused_3) use ($alternativeReaderT1_2_12) {
  $__num = \func_num_args();
  $__res = $alternativeReaderT1_2_12;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Reader_Trans_monadPlusReaderT'] = __NAMESPACE__ . '\\majControl_majMonad_majReader_majTrans_monadmajPlusmajReadermajT';

