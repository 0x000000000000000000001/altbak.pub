<?php

namespace Control\Monad\Except\Trans;

// ALL IMPORTS: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Monad, Control.Monad.Cont.Class, Control.Monad.Error.Class, Control.Monad.Except.Trans, Control.Monad.Reader.Class, Control.Monad.Rec.Class, Control.Monad.ST.Class, Control.Monad.State.Class, Control.Monad.Trans.Class, Control.Monad.Writer.Class, Control.MonadPlus, Control.Plus, Control.Semigroupoid, Data.Either, Data.Function, Data.Functor, Data.Monoid, Data.Newtype, Data.Semigroup, Data.Tuple, Effect.Class, Prelude, Prim
// TO REQUIRE: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Monad, Control.Monad.Cont.Class, Control.Monad.Error.Class, Control.Monad.Except.Trans, Control.Monad.Reader.Class, Control.Monad.Rec.Class, Control.Monad.ST.Class, Control.Monad.State.Class, Control.Monad.Trans.Class, Control.Monad.Writer.Class, Control.MonadPlus, Control.Plus, Control.Semigroupoid, Data.Either, Data.Function, Data.Functor, Data.Monoid, Data.Newtype, Data.Semigroup, Data.Tuple, Effect.Class, Prelude
require_once __DIR__ . '/../Control.Alt/index.php';
require_once __DIR__ . '/../Control.Alternative/index.php';
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Category/index.php';
require_once __DIR__ . '/../Control.Monad/index.php';
require_once __DIR__ . '/../Control.Monad.Cont.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Error.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Except.Trans/index.php';
require_once __DIR__ . '/../Control.Monad.Reader.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Rec.Class/index.php';
require_once __DIR__ . '/../Control.Monad.ST.Class/index.php';
require_once __DIR__ . '/../Control.Monad.State.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Trans.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Writer.Class/index.php';
require_once __DIR__ . '/../Control.MonadPlus/index.php';
require_once __DIR__ . '/../Control.Plus/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Either/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Monoid/index.php';
require_once __DIR__ . '/../Data.Newtype/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Tuple/index.php';
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




// Control_Monad_Except_Trans_ExceptT
function majControl_majMonad_majExcept_majTrans_majExceptmajT($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majExcept_majTrans_majExceptmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Except_Trans_ExceptT'] = __NAMESPACE__ . '\\majControl_majMonad_majExcept_majTrans_majExceptmajT';

// Control_Monad_Except_Trans_withExceptT
function majControl_majMonad_majExcept_majTrans_withmajExceptmajT($dictFunctor_0, $f_1 = null, $v_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majExcept_majTrans_withmajExceptmajT';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictFunctor_0)->{'map'})(function($v2_3) use ($f_1) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v2_3 instanceof \Data\Either\Data_Either_Right) {
$__t0 = new \Data\Either\Data_Either_Right(($v2_3)->{'value0'});
goto end_branch_0;;
};
  if ($v2_3 instanceof \Data\Either\Data_Either_Left) {
$__t0 = new \Data\Either\Data_Either_Left(($f_1)(($v2_3)->{'value0'}));
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_2);
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Control_Monad_Except_Trans_withExceptT'] = __NAMESPACE__ . '\\majControl_majMonad_majExcept_majTrans_withmajExceptmajT';

// Control_Monad_Except_Trans_runExceptT
function majControl_majMonad_majExcept_majTrans_runmajExceptmajT($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majExcept_majTrans_runmajExceptmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $v_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Except_Trans_runExceptT'] = __NAMESPACE__ . '\\majControl_majMonad_majExcept_majTrans_runmajExceptmajT';

// Control_Monad_Except_Trans_newtypeExceptT
$GLOBALS['Control_Monad_Except_Trans_newtypeExceptT'] = (object)["Coercible0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Control_Monad_Except_Trans_monadTransExceptT
$GLOBALS['Control_Monad_Except_Trans_monadTransExceptT'] = (object)["lift" => function($dictMonad_0) {
  $__num = \func_num_args();
  $Bind1_1_0 = (($dictMonad_0)->{'Bind1'})(null);
  $pure_2_1 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $__res = function($m_3) use ($Bind1_1_0, $pure_2_1) {
  $__num = \func_num_args();
  $__res = ((($Bind1_1_0)->{'bind'})($m_3))(function($a_4) use ($pure_2_1) {
  $__num = \func_num_args();
  $__res = ($pure_2_1)(new \Data\Either\Data_Either_Right($a_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Control_Monad_Except_Trans_mapExceptT
function majControl_majMonad_majExcept_majTrans_mapmajExceptmajT($f_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majExcept_majTrans_mapmajExceptmajT';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($f_0)($v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Control_Monad_Except_Trans_mapExceptT'] = __NAMESPACE__ . '\\majControl_majMonad_majExcept_majTrans_mapmajExceptmajT';

// Control_Monad_Except_Trans_functorExceptT
function majControl_majMonad_majExcept_majTrans_functormajExceptmajT($dictFunctor_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majExcept_majTrans_functormajExceptmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["map" => function($f_1) use ($dictFunctor_0) {
  $__num = \func_num_args();
  $__local_var_2_0 = (($dictFunctor_0)->{'map'})(function($m_2) use ($f_1) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($m_2 instanceof \Data\Either\Data_Either_Left) {
$__t0 = new \Data\Either\Data_Either_Left(($m_2)->{'value0'});
goto end_branch_0;;
};
  if ($m_2 instanceof \Data\Either\Data_Either_Right) {
$__t0 = new \Data\Either\Data_Either_Right(($f_1)(($m_2)->{'value0'}));
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_3) use ($__local_var_2_0) {
  $__num = \func_num_args();
  $__res = ($__local_var_2_0)($v_3);
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
$GLOBALS['Control_Monad_Except_Trans_functorExceptT'] = __NAMESPACE__ . '\\majControl_majMonad_majExcept_majTrans_functormajExceptmajT';

// Control_Monad_Except_Trans_except
function majControl_majMonad_majExcept_majTrans_except($dictApplicative_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majExcept_majTrans_except';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_1) {
  $__num = \func_num_args();
  $__res = $x_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($dictApplicative_0)->{'pure'});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Except_Trans_except'] = __NAMESPACE__ . '\\majControl_majMonad_majExcept_majTrans_except';

// Control_Monad_Except_Trans_monadExceptT
function majControl_majMonad_majExcept_majTrans_monadmajExceptmajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majExcept_majTrans_monadmajExceptmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["Applicative0" => function($_dollar___unused_1) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_1) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_bindExceptT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Except_Trans_monadExceptT'] = __NAMESPACE__ . '\\majControl_majMonad_majExcept_majTrans_monadmajExceptmajT';

// Control_Monad_Except_Trans_bindExceptT
function majControl_majMonad_majExcept_majTrans_bindmajExceptmajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majExcept_majTrans_bindmajExceptmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Bind1_1_0 = (($dictMonad_0)->{'Bind1'})(null);
  $pure_2_1 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $__res = (object)["bind" => function($v_3) use ($Bind1_1_0, $pure_2_1) {
  $__num = \func_num_args();
  $__res = function($k_4) use ($Bind1_1_0, $pure_2_1, $v_3) {
  $__num = \func_num_args();
  $__local_var_5_2 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_2_1))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_1_0)->{'bind'})($v_3))(function($v2_6) use ($__local_var_5_2, $k_4) {
  $__num = \func_num_args();
  $__t3 = null;;
  if ($v2_6 instanceof \Data\Either\Data_Either_Left) {
$__t3 = ($__local_var_5_2)(($v2_6)->{'value0'});
goto end_branch_3;;
};
  if ($v2_6 instanceof \Data\Either\Data_Either_Right) {
$__t3 = ($k_4)(($v2_6)->{'value0'});
goto end_branch_3;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t3 = null;
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_3) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Except_Trans_bindExceptT'] = __NAMESPACE__ . '\\majControl_majMonad_majExcept_majTrans_bindmajExceptmajT';

// Control_Monad_Except_Trans_applyExceptT
function majControl_majMonad_majExcept_majTrans_applymajExceptmajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majExcept_majTrans_applymajExceptmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_1_0 = (object)["map" => function($f_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__local_var_3_1 = (($__local_var_1_0)->{'map'})(function($m_3) use ($f_2) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($m_3 instanceof \Data\Either\Data_Either_Left) {
$__t1 = new \Data\Either\Data_Either_Left(($m_3)->{'value0'});
goto end_branch_1;;
};
  if ($m_3 instanceof \Data\Either\Data_Either_Right) {
$__t1 = new \Data\Either\Data_Either_Right(($f_2)(($m_3)->{'value0'}));
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_4) use ($__local_var_3_1) {
  $__num = \func_num_args();
  $__res = ($__local_var_3_1)($v_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_2_4 = (($dictMonad_0)->{'Bind1'})(null);
  $pure_3_5 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_2_4 = (object)["bind" => function($v_4) use ($Bind1_2_4, $pure_3_5) {
  $__num = \func_num_args();
  $__res = function($k_5) use ($Bind1_2_4, $pure_3_5, $v_4) {
  $__num = \func_num_args();
  $__local_var_6_6 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_3_5))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_2_4)->{'bind'})($v_4))(function($v2_7) use ($__local_var_6_6, $k_5) {
  $__num = \func_num_args();
  $__t7 = null;;
  if ($v2_7 instanceof \Data\Either\Data_Either_Left) {
$__t7 = ($__local_var_6_6)(($v2_7)->{'value0'});
goto end_branch_7;;
};
  if ($v2_7 instanceof \Data\Either\Data_Either_Right) {
$__t7 = ($k_5)(($v2_7)->{'value0'});
goto end_branch_7;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t7 = null;
  end_branch_7:;
  $__res = $__t7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_4) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_3_9 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_4) use ($Applicative0_3_9, $Bind1_2_4) {
  $__num = \func_num_args();
  $__res = function($a_5) use ($Applicative0_3_9, $Bind1_2_4, $f_4) {
  $__num = \func_num_args();
  $__res = ((($Bind1_2_4)->{'bind'})($f_4))(function($f_prime__6) use ($Applicative0_3_9, $Bind1_2_4, $a_5) {
  $__num = \func_num_args();
  $__res = ((($Bind1_2_4)->{'bind'})($a_5))(function($a_prime__7) use ($Applicative0_3_9, $f_prime__6) {
  $__num = \func_num_args();
  $__res = (($Applicative0_3_9)->{'pure'})(($f_prime__6)($a_prime__7));
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
}, "Functor0" => function($_dollar___unused_2) use ($functorExceptT1_1_0) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Except_Trans_applyExceptT'] = __NAMESPACE__ . '\\majControl_majMonad_majExcept_majTrans_applymajExceptmajT';

// Control_Monad_Except_Trans_applicativeExceptT
function majControl_majMonad_majExcept_majTrans_applicativemajExceptmajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majExcept_majTrans_applicativemajExceptmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_1) {
  $__num = \func_num_args();
  $__res = $x_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_1) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_2_0 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_2_0 = (object)["map" => function($f_3) use ($__local_var_2_0) {
  $__num = \func_num_args();
  $__local_var_4_1 = (($__local_var_2_0)->{'map'})(function($m_4) use ($f_3) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($m_4 instanceof \Data\Either\Data_Either_Left) {
$__t1 = new \Data\Either\Data_Either_Left(($m_4)->{'value0'});
goto end_branch_1;;
};
  if ($m_4 instanceof \Data\Either\Data_Either_Right) {
$__t1 = new \Data\Either\Data_Either_Right(($f_3)(($m_4)->{'value0'}));
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_5) use ($__local_var_4_1) {
  $__num = \func_num_args();
  $__res = ($__local_var_4_1)($v_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_3_4 = (($dictMonad_0)->{'Bind1'})(null);
  $pure_4_5 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_3_4 = (object)["bind" => function($v_5) use ($Bind1_3_4, $pure_4_5) {
  $__num = \func_num_args();
  $__res = function($k_6) use ($Bind1_3_4, $pure_4_5, $v_5) {
  $__num = \func_num_args();
  $__local_var_7_6 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_4_5))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_3_4)->{'bind'})($v_5))(function($v2_8) use ($__local_var_7_6, $k_6) {
  $__num = \func_num_args();
  $__t7 = null;;
  if ($v2_8 instanceof \Data\Either\Data_Either_Left) {
$__t7 = ($__local_var_7_6)(($v2_8)->{'value0'});
goto end_branch_7;;
};
  if ($v2_8 instanceof \Data\Either\Data_Either_Right) {
$__t7 = ($k_6)(($v2_8)->{'value0'});
goto end_branch_7;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t7 = null;
  end_branch_7:;
  $__res = $__t7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_5) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_4_9 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_5) use ($Applicative0_4_9, $Bind1_3_4) {
  $__num = \func_num_args();
  $__res = function($a_6) use ($Applicative0_4_9, $Bind1_3_4, $f_5) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_4)->{'bind'})($f_5))(function($f_prime__7) use ($Applicative0_4_9, $Bind1_3_4, $a_6) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_4)->{'bind'})($a_6))(function($a_prime__8) use ($Applicative0_4_9, $f_prime__7) {
  $__num = \func_num_args();
  $__res = (($Applicative0_4_9)->{'pure'})(($f_prime__7)($a_prime__8));
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
}, "Functor0" => function($_dollar___unused_3) use ($functorExceptT1_2_0) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_2_0;
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
$GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'] = __NAMESPACE__ . '\\majControl_majMonad_majExcept_majTrans_applicativemajExceptmajT';

// Control_Monad_Except_Trans_semigroupExceptT
function majControl_majMonad_majExcept_majTrans_semigroupmajExceptmajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majExcept_majTrans_semigroupmajExceptmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_1_0 = (object)["map" => function($f_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__local_var_3_1 = (($__local_var_1_0)->{'map'})(function($m_3) use ($f_2) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($m_3 instanceof \Data\Either\Data_Either_Left) {
$__t1 = new \Data\Either\Data_Either_Left(($m_3)->{'value0'});
goto end_branch_1;;
};
  if ($m_3 instanceof \Data\Either\Data_Either_Right) {
$__t1 = new \Data\Either\Data_Either_Right(($f_2)(($m_3)->{'value0'}));
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_4) use ($__local_var_3_1) {
  $__num = \func_num_args();
  $__res = ($__local_var_3_1)($v_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_2_4 = (($dictMonad_0)->{'Bind1'})(null);
  $pure_3_5 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_2_4 = (object)["bind" => function($v_4) use ($Bind1_2_4, $pure_3_5) {
  $__num = \func_num_args();
  $__res = function($k_5) use ($Bind1_2_4, $pure_3_5, $v_4) {
  $__num = \func_num_args();
  $__local_var_6_6 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_3_5))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_2_4)->{'bind'})($v_4))(function($v2_7) use ($__local_var_6_6, $k_5) {
  $__num = \func_num_args();
  $__t7 = null;;
  if ($v2_7 instanceof \Data\Either\Data_Either_Left) {
$__t7 = ($__local_var_6_6)(($v2_7)->{'value0'});
goto end_branch_7;;
};
  if ($v2_7 instanceof \Data\Either\Data_Either_Right) {
$__t7 = ($k_5)(($v2_7)->{'value0'});
goto end_branch_7;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t7 = null;
  end_branch_7:;
  $__res = $__t7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_4) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_3_9 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_3) {
  $__num = \func_num_args();
  $__res = $x_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_3) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_4_9 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_4_9 = (object)["map" => function($f_5) use ($__local_var_4_9) {
  $__num = \func_num_args();
  $__local_var_6_10 = (($__local_var_4_9)->{'map'})(function($m_6) use ($f_5) {
  $__num = \func_num_args();
  $__t10 = null;;
  if ($m_6 instanceof \Data\Either\Data_Either_Left) {
$__t10 = new \Data\Either\Data_Either_Left(($m_6)->{'value0'});
goto end_branch_10;;
};
  if ($m_6 instanceof \Data\Either\Data_Either_Right) {
$__t10 = new \Data\Either\Data_Either_Right(($f_5)(($m_6)->{'value0'}));
goto end_branch_10;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t10 = null;
  end_branch_10:;
  $__res = $__t10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_7) use ($__local_var_6_10) {
  $__num = \func_num_args();
  $__res = ($__local_var_6_10)($v_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_5_13 = (($dictMonad_0)->{'Bind1'})(null);
  $pure_6_14 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_5_13 = (object)["bind" => function($v_7) use ($Bind1_5_13, $pure_6_14) {
  $__num = \func_num_args();
  $__res = function($k_8) use ($Bind1_5_13, $pure_6_14, $v_7) {
  $__num = \func_num_args();
  $__local_var_9_15 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_6_14))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_5_13)->{'bind'})($v_7))(function($v2_10) use ($__local_var_9_15, $k_8) {
  $__num = \func_num_args();
  $__t16 = null;;
  if ($v2_10 instanceof \Data\Either\Data_Either_Left) {
$__t16 = ($__local_var_9_15)(($v2_10)->{'value0'});
goto end_branch_16;;
};
  if ($v2_10 instanceof \Data\Either\Data_Either_Right) {
$__t16 = ($k_8)(($v2_10)->{'value0'});
goto end_branch_16;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t16 = null;
  end_branch_16:;
  $__res = $__t16;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_7) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_6_18 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_7) use ($Applicative0_6_18, $Bind1_5_13) {
  $__num = \func_num_args();
  $__res = function($a_8) use ($Applicative0_6_18, $Bind1_5_13, $f_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_5_13)->{'bind'})($f_7))(function($f_prime__9) use ($Applicative0_6_18, $Bind1_5_13, $a_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_5_13)->{'bind'})($a_8))(function($a_prime__10) use ($Applicative0_6_18, $f_prime__9) {
  $__num = \func_num_args();
  $__res = (($Applicative0_6_18)->{'pure'})(($f_prime__9)($a_prime__10));
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
}, "Functor0" => function($_dollar___unused_5) use ($functorExceptT1_4_9) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_4_9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyExceptT1_1_0 = (object)["apply" => function($f_4) use ($Applicative0_3_9, $Bind1_2_4) {
  $__num = \func_num_args();
  $__res = function($a_5) use ($Applicative0_3_9, $Bind1_2_4, $f_4) {
  $__num = \func_num_args();
  $__res = ((($Bind1_2_4)->{'bind'})($f_4))(function($f_prime__6) use ($Applicative0_3_9, $Bind1_2_4, $a_5) {
  $__num = \func_num_args();
  $__res = ((($Bind1_2_4)->{'bind'})($a_5))(function($a_prime__7) use ($Applicative0_3_9, $f_prime__6) {
  $__num = \func_num_args();
  $__res = (($Applicative0_3_9)->{'pure'})(($f_prime__6)($a_prime__7));
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
}, "Functor0" => function($_dollar___unused_2) use ($functorExceptT1_1_0) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = function($dictSemigroup_2) use ($applyExceptT1_1_0) {
  $__num = \func_num_args();
  $Functor0_3_21 = (($applyExceptT1_1_0)->{'Functor0'})(null);
  $__local_var_4_22 = ($dictSemigroup_2)->{'append'};
  $__res = (object)["append" => function($a_5) use ($Functor0_3_21, $__local_var_4_22, $applyExceptT1_1_0) {
  $__num = \func_num_args();
  $__res = function($b_6) use ($Functor0_3_21, $__local_var_4_22, $a_5, $applyExceptT1_1_0) {
  $__num = \func_num_args();
  $__res = ((($applyExceptT1_1_0)->{'apply'})(((($Functor0_3_21)->{'map'})($__local_var_4_22))($a_5)))($b_6);
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
$GLOBALS['Control_Monad_Except_Trans_semigroupExceptT'] = __NAMESPACE__ . '\\majControl_majMonad_majExcept_majTrans_semigroupmajExceptmajT';

// Control_Monad_Except_Trans_monadAskExceptT
function majControl_majMonad_majExcept_majTrans_monadmajAskmajExceptmajT($dictMonadAsk_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majExcept_majTrans_monadmajAskmajExceptmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictMonadAsk_0)->{'Monad0'})(null);
  $monadExceptT1_1_0 = (object)["Applicative0" => function($_dollar___unused_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_3) {
  $__num = \func_num_args();
  $__res = $x_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($__local_var_1_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_3) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__local_var_4_1 = (((((($__local_var_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_4_1 = (object)["map" => function($f_5) use ($__local_var_4_1) {
  $__num = \func_num_args();
  $__local_var_6_2 = (($__local_var_4_1)->{'map'})(function($m_6) use ($f_5) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ($m_6 instanceof \Data\Either\Data_Either_Left) {
$__t2 = new \Data\Either\Data_Either_Left(($m_6)->{'value0'});
goto end_branch_2;;
};
  if ($m_6 instanceof \Data\Either\Data_Either_Right) {
$__t2 = new \Data\Either\Data_Either_Right(($f_5)(($m_6)->{'value0'}));
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_7) use ($__local_var_6_2) {
  $__num = \func_num_args();
  $__res = ($__local_var_6_2)($v_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_5_5 = (($__local_var_1_0)->{'Bind1'})(null);
  $pure_6_6 = ((($__local_var_1_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_5_5 = (object)["bind" => function($v_7) use ($Bind1_5_5, $pure_6_6) {
  $__num = \func_num_args();
  $__res = function($k_8) use ($Bind1_5_5, $pure_6_6, $v_7) {
  $__num = \func_num_args();
  $__local_var_9_7 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_6_6))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_5_5)->{'bind'})($v_7))(function($v2_10) use ($__local_var_9_7, $k_8) {
  $__num = \func_num_args();
  $__t8 = null;;
  if ($v2_10 instanceof \Data\Either\Data_Either_Left) {
$__t8 = ($__local_var_9_7)(($v2_10)->{'value0'});
goto end_branch_8;;
};
  if ($v2_10 instanceof \Data\Either\Data_Either_Right) {
$__t8 = ($k_8)(($v2_10)->{'value0'});
goto end_branch_8;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t8 = null;
  end_branch_8:;
  $__res = $__t8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_7) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__local_var_8_9 = (((((($__local_var_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_8_9 = (object)["map" => function($f_9) use ($__local_var_8_9) {
  $__num = \func_num_args();
  $__local_var_10_10 = (($__local_var_8_9)->{'map'})(function($m_10) use ($f_9) {
  $__num = \func_num_args();
  $__t10 = null;;
  if ($m_10 instanceof \Data\Either\Data_Either_Left) {
$__t10 = new \Data\Either\Data_Either_Left(($m_10)->{'value0'});
goto end_branch_10;;
};
  if ($m_10 instanceof \Data\Either\Data_Either_Right) {
$__t10 = new \Data\Either\Data_Either_Right(($f_9)(($m_10)->{'value0'}));
goto end_branch_10;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t10 = null;
  end_branch_10:;
  $__res = $__t10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_11) use ($__local_var_10_10) {
  $__num = \func_num_args();
  $__res = ($__local_var_10_10)($v_11);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_9_13 = (($__local_var_1_0)->{'Bind1'})(null);
  $pure_10_14 = ((($__local_var_1_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_9_13 = (object)["bind" => function($v_11) use ($Bind1_9_13, $pure_10_14) {
  $__num = \func_num_args();
  $__res = function($k_12) use ($Bind1_9_13, $pure_10_14, $v_11) {
  $__num = \func_num_args();
  $__local_var_13_15 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_10_14))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_9_13)->{'bind'})($v_11))(function($v2_14) use ($__local_var_13_15, $k_12) {
  $__num = \func_num_args();
  $__t16 = null;;
  if ($v2_14 instanceof \Data\Either\Data_Either_Left) {
$__t16 = ($__local_var_13_15)(($v2_14)->{'value0'});
goto end_branch_16;;
};
  if ($v2_14 instanceof \Data\Either\Data_Either_Right) {
$__t16 = ($k_12)(($v2_14)->{'value0'});
goto end_branch_16;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t16 = null;
  end_branch_16:;
  $__res = $__t16;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_11) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($__local_var_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_10_18 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($__local_var_1_0);
  $__res = (object)["apply" => function($f_11) use ($Applicative0_10_18, $Bind1_9_13) {
  $__num = \func_num_args();
  $__res = function($a_12) use ($Applicative0_10_18, $Bind1_9_13, $f_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_13)->{'bind'})($f_11))(function($f_prime__13) use ($Applicative0_10_18, $Bind1_9_13, $a_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_13)->{'bind'})($a_12))(function($a_prime__14) use ($Applicative0_10_18, $f_prime__13) {
  $__num = \func_num_args();
  $__res = (($Applicative0_10_18)->{'pure'})(($f_prime__13)($a_prime__14));
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
}, "Functor0" => function($_dollar___unused_9) use ($functorExceptT1_8_9) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_8_9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_6_20 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($__local_var_1_0);
  $__res = (object)["apply" => function($f_7) use ($Applicative0_6_20, $Bind1_5_5) {
  $__num = \func_num_args();
  $__res = function($a_8) use ($Applicative0_6_20, $Bind1_5_5, $f_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_5_5)->{'bind'})($f_7))(function($f_prime__9) use ($Applicative0_6_20, $Bind1_5_5, $a_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_5_5)->{'bind'})($a_8))(function($a_prime__10) use ($Applicative0_6_20, $f_prime__9) {
  $__num = \func_num_args();
  $__res = (($Applicative0_6_20)->{'pure'})(($f_prime__9)($a_prime__10));
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
}, "Functor0" => function($_dollar___unused_5) use ($functorExceptT1_4_1) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_4_1;
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
  $Bind1_3_21 = (($__local_var_1_0)->{'Bind1'})(null);
  $pure_4_22 = ((($__local_var_1_0)->{'Applicative0'})(null))->{'pure'};
  $__res = (object)["bind" => function($v_5) use ($Bind1_3_21, $pure_4_22) {
  $__num = \func_num_args();
  $__res = function($k_6) use ($Bind1_3_21, $pure_4_22, $v_5) {
  $__num = \func_num_args();
  $__local_var_7_23 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_4_22))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_3_21)->{'bind'})($v_5))(function($v2_8) use ($__local_var_7_23, $k_6) {
  $__num = \func_num_args();
  $__t24 = null;;
  if ($v2_8 instanceof \Data\Either\Data_Either_Left) {
$__t24 = ($__local_var_7_23)(($v2_8)->{'value0'});
goto end_branch_24;;
};
  if ($v2_8 instanceof \Data\Either\Data_Either_Right) {
$__t24 = ($k_6)(($v2_8)->{'value0'});
goto end_branch_24;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t24 = null;
  end_branch_24:;
  $__res = $__t24;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_5) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__local_var_6_25 = (((((($__local_var_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_6_25 = (object)["map" => function($f_7) use ($__local_var_6_25) {
  $__num = \func_num_args();
  $__local_var_8_26 = (($__local_var_6_25)->{'map'})(function($m_8) use ($f_7) {
  $__num = \func_num_args();
  $__t26 = null;;
  if ($m_8 instanceof \Data\Either\Data_Either_Left) {
$__t26 = new \Data\Either\Data_Either_Left(($m_8)->{'value0'});
goto end_branch_26;;
};
  if ($m_8 instanceof \Data\Either\Data_Either_Right) {
$__t26 = new \Data\Either\Data_Either_Right(($f_7)(($m_8)->{'value0'}));
goto end_branch_26;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t26 = null;
  end_branch_26:;
  $__res = $__t26;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_9) use ($__local_var_8_26) {
  $__num = \func_num_args();
  $__res = ($__local_var_8_26)($v_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_7_29 = (($__local_var_1_0)->{'Bind1'})(null);
  $pure_8_30 = ((($__local_var_1_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_7_29 = (object)["bind" => function($v_9) use ($Bind1_7_29, $pure_8_30) {
  $__num = \func_num_args();
  $__res = function($k_10) use ($Bind1_7_29, $pure_8_30, $v_9) {
  $__num = \func_num_args();
  $__local_var_11_31 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_8_30))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_7_29)->{'bind'})($v_9))(function($v2_12) use ($__local_var_11_31, $k_10) {
  $__num = \func_num_args();
  $__t32 = null;;
  if ($v2_12 instanceof \Data\Either\Data_Either_Left) {
$__t32 = ($__local_var_11_31)(($v2_12)->{'value0'});
goto end_branch_32;;
};
  if ($v2_12 instanceof \Data\Either\Data_Either_Right) {
$__t32 = ($k_10)(($v2_12)->{'value0'});
goto end_branch_32;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t32 = null;
  end_branch_32:;
  $__res = $__t32;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_9) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($__local_var_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_8_34 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_8) {
  $__num = \func_num_args();
  $__res = $x_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($__local_var_1_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_8) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__local_var_9_34 = (((((($__local_var_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_9_34 = (object)["map" => function($f_10) use ($__local_var_9_34) {
  $__num = \func_num_args();
  $__local_var_11_35 = (($__local_var_9_34)->{'map'})(function($m_11) use ($f_10) {
  $__num = \func_num_args();
  $__t35 = null;;
  if ($m_11 instanceof \Data\Either\Data_Either_Left) {
$__t35 = new \Data\Either\Data_Either_Left(($m_11)->{'value0'});
goto end_branch_35;;
};
  if ($m_11 instanceof \Data\Either\Data_Either_Right) {
$__t35 = new \Data\Either\Data_Either_Right(($f_10)(($m_11)->{'value0'}));
goto end_branch_35;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t35 = null;
  end_branch_35:;
  $__res = $__t35;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_12) use ($__local_var_11_35) {
  $__num = \func_num_args();
  $__res = ($__local_var_11_35)($v_12);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_10_38 = (($__local_var_1_0)->{'Bind1'})(null);
  $pure_11_39 = ((($__local_var_1_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_10_38 = (object)["bind" => function($v_12) use ($Bind1_10_38, $pure_11_39) {
  $__num = \func_num_args();
  $__res = function($k_13) use ($Bind1_10_38, $pure_11_39, $v_12) {
  $__num = \func_num_args();
  $__local_var_14_40 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_11_39))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_10_38)->{'bind'})($v_12))(function($v2_15) use ($__local_var_14_40, $k_13) {
  $__num = \func_num_args();
  $__t41 = null;;
  if ($v2_15 instanceof \Data\Either\Data_Either_Left) {
$__t41 = ($__local_var_14_40)(($v2_15)->{'value0'});
goto end_branch_41;;
};
  if ($v2_15 instanceof \Data\Either\Data_Either_Right) {
$__t41 = ($k_13)(($v2_15)->{'value0'});
goto end_branch_41;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t41 = null;
  end_branch_41:;
  $__res = $__t41;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_12) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($__local_var_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_11_43 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($__local_var_1_0);
  $__res = (object)["apply" => function($f_12) use ($Applicative0_11_43, $Bind1_10_38) {
  $__num = \func_num_args();
  $__res = function($a_13) use ($Applicative0_11_43, $Bind1_10_38, $f_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_38)->{'bind'})($f_12))(function($f_prime__14) use ($Applicative0_11_43, $Bind1_10_38, $a_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_38)->{'bind'})($a_13))(function($a_prime__15) use ($Applicative0_11_43, $f_prime__14) {
  $__num = \func_num_args();
  $__res = (($Applicative0_11_43)->{'pure'})(($f_prime__14)($a_prime__15));
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
}, "Functor0" => function($_dollar___unused_10) use ($functorExceptT1_9_34) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_9_34;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_9) use ($Applicative0_8_34, $Bind1_7_29) {
  $__num = \func_num_args();
  $__res = function($a_10) use ($Applicative0_8_34, $Bind1_7_29, $f_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_29)->{'bind'})($f_9))(function($f_prime__11) use ($Applicative0_8_34, $Bind1_7_29, $a_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_29)->{'bind'})($a_10))(function($a_prime__12) use ($Applicative0_8_34, $f_prime__11) {
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
}, "Functor0" => function($_dollar___unused_7) use ($functorExceptT1_6_25) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_6_25;
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
  $__local_var_2_46 = (($dictMonadAsk_0)->{'Monad0'})(null);
  $pure_3_47 = ((($__local_var_2_46)->{'Applicative0'})(null))->{'pure'};
  $__res = (object)["ask" => ((((($__local_var_2_46)->{'Bind1'})(null))->{'bind'})(($dictMonadAsk_0)->{'ask'}))(function($a_4) use ($pure_3_47) {
  $__num = \func_num_args();
  $__res = ($pure_3_47)(new \Data\Either\Data_Either_Right($a_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), "Monad0" => function($_dollar___unused_2) use ($monadExceptT1_1_0) {
  $__num = \func_num_args();
  $__res = $monadExceptT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Except_Trans_monadAskExceptT'] = __NAMESPACE__ . '\\majControl_majMonad_majExcept_majTrans_monadmajAskmajExceptmajT';

// Control_Monad_Except_Trans_monadReaderExceptT
function majControl_majMonad_majExcept_majTrans_monadmajReadermajExceptmajT($dictMonadReader_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majExcept_majTrans_monadmajReadermajExceptmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictMonadReader_0)->{'MonadAsk0'})(null);
  $__local_var_2_1 = (($__local_var_1_0)->{'Monad0'})(null);
  $monadExceptT1_2_1 = (object)["Applicative0" => function($_dollar___unused_3) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_4) {
  $__num = \func_num_args();
  $__res = $x_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_4) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_5_2 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_5_2 = (object)["map" => function($f_6) use ($__local_var_5_2) {
  $__num = \func_num_args();
  $__local_var_7_3 = (($__local_var_5_2)->{'map'})(function($m_7) use ($f_6) {
  $__num = \func_num_args();
  $__t3 = null;;
  if ($m_7 instanceof \Data\Either\Data_Either_Left) {
$__t3 = new \Data\Either\Data_Either_Left(($m_7)->{'value0'});
goto end_branch_3;;
};
  if ($m_7 instanceof \Data\Either\Data_Either_Right) {
$__t3 = new \Data\Either\Data_Either_Right(($f_6)(($m_7)->{'value0'}));
goto end_branch_3;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t3 = null;
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_8) use ($__local_var_7_3) {
  $__num = \func_num_args();
  $__res = ($__local_var_7_3)($v_8);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_6_6 = (($__local_var_2_1)->{'Bind1'})(null);
  $pure_7_7 = ((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'};
  $Bind1_6_6 = (object)["bind" => function($v_8) use ($Bind1_6_6, $pure_7_7) {
  $__num = \func_num_args();
  $__res = function($k_9) use ($Bind1_6_6, $pure_7_7, $v_8) {
  $__num = \func_num_args();
  $__local_var_10_8 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_7_7))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_6_6)->{'bind'})($v_8))(function($v2_11) use ($__local_var_10_8, $k_9) {
  $__num = \func_num_args();
  $__t9 = null;;
  if ($v2_11 instanceof \Data\Either\Data_Either_Left) {
$__t9 = ($__local_var_10_8)(($v2_11)->{'value0'});
goto end_branch_9;;
};
  if ($v2_11 instanceof \Data\Either\Data_Either_Right) {
$__t9 = ($k_9)(($v2_11)->{'value0'});
goto end_branch_9;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t9 = null;
  end_branch_9:;
  $__res = $__t9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_8) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_9_10 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_9_10 = (object)["map" => function($f_10) use ($__local_var_9_10) {
  $__num = \func_num_args();
  $__local_var_11_11 = (($__local_var_9_10)->{'map'})(function($m_11) use ($f_10) {
  $__num = \func_num_args();
  $__t11 = null;;
  if ($m_11 instanceof \Data\Either\Data_Either_Left) {
$__t11 = new \Data\Either\Data_Either_Left(($m_11)->{'value0'});
goto end_branch_11;;
};
  if ($m_11 instanceof \Data\Either\Data_Either_Right) {
$__t11 = new \Data\Either\Data_Either_Right(($f_10)(($m_11)->{'value0'}));
goto end_branch_11;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t11 = null;
  end_branch_11:;
  $__res = $__t11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_12) use ($__local_var_11_11) {
  $__num = \func_num_args();
  $__res = ($__local_var_11_11)($v_12);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_10_14 = (($__local_var_2_1)->{'Bind1'})(null);
  $pure_11_15 = ((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'};
  $Bind1_10_14 = (object)["bind" => function($v_12) use ($Bind1_10_14, $pure_11_15) {
  $__num = \func_num_args();
  $__res = function($k_13) use ($Bind1_10_14, $pure_11_15, $v_12) {
  $__num = \func_num_args();
  $__local_var_14_16 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_11_15))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_10_14)->{'bind'})($v_12))(function($v2_15) use ($__local_var_14_16, $k_13) {
  $__num = \func_num_args();
  $__t17 = null;;
  if ($v2_15 instanceof \Data\Either\Data_Either_Left) {
$__t17 = ($__local_var_14_16)(($v2_15)->{'value0'});
goto end_branch_17;;
};
  if ($v2_15 instanceof \Data\Either\Data_Either_Right) {
$__t17 = ($k_13)(($v2_15)->{'value0'});
goto end_branch_17;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t17 = null;
  end_branch_17:;
  $__res = $__t17;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_12) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_13_18 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_13_18 = (object)["map" => function($f_14) use ($__local_var_13_18) {
  $__num = \func_num_args();
  $__local_var_15_19 = (($__local_var_13_18)->{'map'})(function($m_15) use ($f_14) {
  $__num = \func_num_args();
  $__t19 = null;;
  if ($m_15 instanceof \Data\Either\Data_Either_Left) {
$__t19 = new \Data\Either\Data_Either_Left(($m_15)->{'value0'});
goto end_branch_19;;
};
  if ($m_15 instanceof \Data\Either\Data_Either_Right) {
$__t19 = new \Data\Either\Data_Either_Right(($f_14)(($m_15)->{'value0'}));
goto end_branch_19;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t19 = null;
  end_branch_19:;
  $__res = $__t19;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_16) use ($__local_var_15_19) {
  $__num = \func_num_args();
  $__res = ($__local_var_15_19)($v_16);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_14_22 = (($__local_var_2_1)->{'Bind1'})(null);
  $pure_15_23 = ((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'};
  $Bind1_14_22 = (object)["bind" => function($v_16) use ($Bind1_14_22, $pure_15_23) {
  $__num = \func_num_args();
  $__res = function($k_17) use ($Bind1_14_22, $pure_15_23, $v_16) {
  $__num = \func_num_args();
  $__local_var_18_24 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_15_23))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_14_22)->{'bind'})($v_16))(function($v2_19) use ($__local_var_18_24, $k_17) {
  $__num = \func_num_args();
  $__t25 = null;;
  if ($v2_19 instanceof \Data\Either\Data_Either_Left) {
$__t25 = ($__local_var_18_24)(($v2_19)->{'value0'});
goto end_branch_25;;
};
  if ($v2_19 instanceof \Data\Either\Data_Either_Right) {
$__t25 = ($k_17)(($v2_19)->{'value0'});
goto end_branch_25;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t25 = null;
  end_branch_25:;
  $__res = $__t25;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_16) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($__local_var_2_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_15_27 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_15) {
  $__num = \func_num_args();
  $__res = $x_15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_15) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_16_27 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_16_27 = (object)["map" => function($f_17) use ($__local_var_16_27) {
  $__num = \func_num_args();
  $__local_var_18_28 = (($__local_var_16_27)->{'map'})(function($m_18) use ($f_17) {
  $__num = \func_num_args();
  $__t28 = null;;
  if ($m_18 instanceof \Data\Either\Data_Either_Left) {
$__t28 = new \Data\Either\Data_Either_Left(($m_18)->{'value0'});
goto end_branch_28;;
};
  if ($m_18 instanceof \Data\Either\Data_Either_Right) {
$__t28 = new \Data\Either\Data_Either_Right(($f_17)(($m_18)->{'value0'}));
goto end_branch_28;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t28 = null;
  end_branch_28:;
  $__res = $__t28;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_19) use ($__local_var_18_28) {
  $__num = \func_num_args();
  $__res = ($__local_var_18_28)($v_19);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_17_31 = (($__local_var_2_1)->{'Bind1'})(null);
  $pure_18_32 = ((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'};
  $Bind1_17_31 = (object)["bind" => function($v_19) use ($Bind1_17_31, $pure_18_32) {
  $__num = \func_num_args();
  $__res = function($k_20) use ($Bind1_17_31, $pure_18_32, $v_19) {
  $__num = \func_num_args();
  $__local_var_21_33 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_18_32))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_17_31)->{'bind'})($v_19))(function($v2_22) use ($__local_var_21_33, $k_20) {
  $__num = \func_num_args();
  $__t34 = null;;
  if ($v2_22 instanceof \Data\Either\Data_Either_Left) {
$__t34 = ($__local_var_21_33)(($v2_22)->{'value0'});
goto end_branch_34;;
};
  if ($v2_22 instanceof \Data\Either\Data_Either_Right) {
$__t34 = ($k_20)(($v2_22)->{'value0'});
goto end_branch_34;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t34 = null;
  end_branch_34:;
  $__res = $__t34;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_19) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($__local_var_2_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_18_36 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($__local_var_2_1);
  $__res = (object)["apply" => function($f_19) use ($Applicative0_18_36, $Bind1_17_31) {
  $__num = \func_num_args();
  $__res = function($a_20) use ($Applicative0_18_36, $Bind1_17_31, $f_19) {
  $__num = \func_num_args();
  $__res = ((($Bind1_17_31)->{'bind'})($f_19))(function($f_prime__21) use ($Applicative0_18_36, $Bind1_17_31, $a_20) {
  $__num = \func_num_args();
  $__res = ((($Bind1_17_31)->{'bind'})($a_20))(function($a_prime__22) use ($Applicative0_18_36, $f_prime__21) {
  $__num = \func_num_args();
  $__res = (($Applicative0_18_36)->{'pure'})(($f_prime__21)($a_prime__22));
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
}, "Functor0" => function($_dollar___unused_17) use ($functorExceptT1_16_27) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_16_27;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_16) use ($Applicative0_15_27, $Bind1_14_22) {
  $__num = \func_num_args();
  $__res = function($a_17) use ($Applicative0_15_27, $Bind1_14_22, $f_16) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_22)->{'bind'})($f_16))(function($f_prime__18) use ($Applicative0_15_27, $Bind1_14_22, $a_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_22)->{'bind'})($a_17))(function($a_prime__19) use ($Applicative0_15_27, $f_prime__18) {
  $__num = \func_num_args();
  $__res = (($Applicative0_15_27)->{'pure'})(($f_prime__18)($a_prime__19));
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
}, "Functor0" => function($_dollar___unused_14) use ($functorExceptT1_13_18) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_13_18;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_11_39 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_11) {
  $__num = \func_num_args();
  $__res = $x_11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_11) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_12_39 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_12_39 = (object)["map" => function($f_13) use ($__local_var_12_39) {
  $__num = \func_num_args();
  $__local_var_14_40 = (($__local_var_12_39)->{'map'})(function($m_14) use ($f_13) {
  $__num = \func_num_args();
  $__t40 = null;;
  if ($m_14 instanceof \Data\Either\Data_Either_Left) {
$__t40 = new \Data\Either\Data_Either_Left(($m_14)->{'value0'});
goto end_branch_40;;
};
  if ($m_14 instanceof \Data\Either\Data_Either_Right) {
$__t40 = new \Data\Either\Data_Either_Right(($f_13)(($m_14)->{'value0'}));
goto end_branch_40;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t40 = null;
  end_branch_40:;
  $__res = $__t40;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_15) use ($__local_var_14_40) {
  $__num = \func_num_args();
  $__res = ($__local_var_14_40)($v_15);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_13_43 = (($__local_var_2_1)->{'Bind1'})(null);
  $pure_14_44 = ((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'};
  $Bind1_13_43 = (object)["bind" => function($v_15) use ($Bind1_13_43, $pure_14_44) {
  $__num = \func_num_args();
  $__res = function($k_16) use ($Bind1_13_43, $pure_14_44, $v_15) {
  $__num = \func_num_args();
  $__local_var_17_45 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_14_44))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_13_43)->{'bind'})($v_15))(function($v2_18) use ($__local_var_17_45, $k_16) {
  $__num = \func_num_args();
  $__t46 = null;;
  if ($v2_18 instanceof \Data\Either\Data_Either_Left) {
$__t46 = ($__local_var_17_45)(($v2_18)->{'value0'});
goto end_branch_46;;
};
  if ($v2_18 instanceof \Data\Either\Data_Either_Right) {
$__t46 = ($k_16)(($v2_18)->{'value0'});
goto end_branch_46;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t46 = null;
  end_branch_46:;
  $__res = $__t46;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_15) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_16_47 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_16_47 = (object)["map" => function($f_17) use ($__local_var_16_47) {
  $__num = \func_num_args();
  $__local_var_18_48 = (($__local_var_16_47)->{'map'})(function($m_18) use ($f_17) {
  $__num = \func_num_args();
  $__t48 = null;;
  if ($m_18 instanceof \Data\Either\Data_Either_Left) {
$__t48 = new \Data\Either\Data_Either_Left(($m_18)->{'value0'});
goto end_branch_48;;
};
  if ($m_18 instanceof \Data\Either\Data_Either_Right) {
$__t48 = new \Data\Either\Data_Either_Right(($f_17)(($m_18)->{'value0'}));
goto end_branch_48;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t48 = null;
  end_branch_48:;
  $__res = $__t48;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_19) use ($__local_var_18_48) {
  $__num = \func_num_args();
  $__res = ($__local_var_18_48)($v_19);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_17_51 = (($__local_var_2_1)->{'Bind1'})(null);
  $pure_18_52 = ((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'};
  $Bind1_17_51 = (object)["bind" => function($v_19) use ($Bind1_17_51, $pure_18_52) {
  $__num = \func_num_args();
  $__res = function($k_20) use ($Bind1_17_51, $pure_18_52, $v_19) {
  $__num = \func_num_args();
  $__local_var_21_53 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_18_52))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_17_51)->{'bind'})($v_19))(function($v2_22) use ($__local_var_21_53, $k_20) {
  $__num = \func_num_args();
  $__t54 = null;;
  if ($v2_22 instanceof \Data\Either\Data_Either_Left) {
$__t54 = ($__local_var_21_53)(($v2_22)->{'value0'});
goto end_branch_54;;
};
  if ($v2_22 instanceof \Data\Either\Data_Either_Right) {
$__t54 = ($k_20)(($v2_22)->{'value0'});
goto end_branch_54;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t54 = null;
  end_branch_54:;
  $__res = $__t54;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_19) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($__local_var_2_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_18_56 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($__local_var_2_1);
  $__res = (object)["apply" => function($f_19) use ($Applicative0_18_56, $Bind1_17_51) {
  $__num = \func_num_args();
  $__res = function($a_20) use ($Applicative0_18_56, $Bind1_17_51, $f_19) {
  $__num = \func_num_args();
  $__res = ((($Bind1_17_51)->{'bind'})($f_19))(function($f_prime__21) use ($Applicative0_18_56, $Bind1_17_51, $a_20) {
  $__num = \func_num_args();
  $__res = ((($Bind1_17_51)->{'bind'})($a_20))(function($a_prime__22) use ($Applicative0_18_56, $f_prime__21) {
  $__num = \func_num_args();
  $__res = (($Applicative0_18_56)->{'pure'})(($f_prime__21)($a_prime__22));
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
}, "Functor0" => function($_dollar___unused_17) use ($functorExceptT1_16_47) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_16_47;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_14_58 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($__local_var_2_1);
  $__res = (object)["apply" => function($f_15) use ($Applicative0_14_58, $Bind1_13_43) {
  $__num = \func_num_args();
  $__res = function($a_16) use ($Applicative0_14_58, $Bind1_13_43, $f_15) {
  $__num = \func_num_args();
  $__res = ((($Bind1_13_43)->{'bind'})($f_15))(function($f_prime__17) use ($Applicative0_14_58, $Bind1_13_43, $a_16) {
  $__num = \func_num_args();
  $__res = ((($Bind1_13_43)->{'bind'})($a_16))(function($a_prime__18) use ($Applicative0_14_58, $f_prime__17) {
  $__num = \func_num_args();
  $__res = (($Applicative0_14_58)->{'pure'})(($f_prime__17)($a_prime__18));
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
}, "Functor0" => function($_dollar___unused_13) use ($functorExceptT1_12_39) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_12_39;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_12) use ($Applicative0_11_39, $Bind1_10_14) {
  $__num = \func_num_args();
  $__res = function($a_13) use ($Applicative0_11_39, $Bind1_10_14, $f_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_14)->{'bind'})($f_12))(function($f_prime__14) use ($Applicative0_11_39, $Bind1_10_14, $a_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_14)->{'bind'})($a_13))(function($a_prime__15) use ($Applicative0_11_39, $f_prime__14) {
  $__num = \func_num_args();
  $__res = (($Applicative0_11_39)->{'pure'})(($f_prime__14)($a_prime__15));
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
}, "Functor0" => function($_dollar___unused_10) use ($functorExceptT1_9_10) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_9_10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_7_61 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_7) {
  $__num = \func_num_args();
  $__res = $x_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_7) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_8_61 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_8_61 = (object)["map" => function($f_9) use ($__local_var_8_61) {
  $__num = \func_num_args();
  $__local_var_10_62 = (($__local_var_8_61)->{'map'})(function($m_10) use ($f_9) {
  $__num = \func_num_args();
  $__t62 = null;;
  if ($m_10 instanceof \Data\Either\Data_Either_Left) {
$__t62 = new \Data\Either\Data_Either_Left(($m_10)->{'value0'});
goto end_branch_62;;
};
  if ($m_10 instanceof \Data\Either\Data_Either_Right) {
$__t62 = new \Data\Either\Data_Either_Right(($f_9)(($m_10)->{'value0'}));
goto end_branch_62;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t62 = null;
  end_branch_62:;
  $__res = $__t62;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_11) use ($__local_var_10_62) {
  $__num = \func_num_args();
  $__res = ($__local_var_10_62)($v_11);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_9_65 = (($__local_var_2_1)->{'Bind1'})(null);
  $pure_10_66 = ((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'};
  $Bind1_9_65 = (object)["bind" => function($v_11) use ($Bind1_9_65, $pure_10_66) {
  $__num = \func_num_args();
  $__res = function($k_12) use ($Bind1_9_65, $pure_10_66, $v_11) {
  $__num = \func_num_args();
  $__local_var_13_67 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_10_66))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_9_65)->{'bind'})($v_11))(function($v2_14) use ($__local_var_13_67, $k_12) {
  $__num = \func_num_args();
  $__t68 = null;;
  if ($v2_14 instanceof \Data\Either\Data_Either_Left) {
$__t68 = ($__local_var_13_67)(($v2_14)->{'value0'});
goto end_branch_68;;
};
  if ($v2_14 instanceof \Data\Either\Data_Either_Right) {
$__t68 = ($k_12)(($v2_14)->{'value0'});
goto end_branch_68;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t68 = null;
  end_branch_68:;
  $__res = $__t68;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_11) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_12_69 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_12_69 = (object)["map" => function($f_13) use ($__local_var_12_69) {
  $__num = \func_num_args();
  $__local_var_14_70 = (($__local_var_12_69)->{'map'})(function($m_14) use ($f_13) {
  $__num = \func_num_args();
  $__t70 = null;;
  if ($m_14 instanceof \Data\Either\Data_Either_Left) {
$__t70 = new \Data\Either\Data_Either_Left(($m_14)->{'value0'});
goto end_branch_70;;
};
  if ($m_14 instanceof \Data\Either\Data_Either_Right) {
$__t70 = new \Data\Either\Data_Either_Right(($f_13)(($m_14)->{'value0'}));
goto end_branch_70;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t70 = null;
  end_branch_70:;
  $__res = $__t70;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_15) use ($__local_var_14_70) {
  $__num = \func_num_args();
  $__res = ($__local_var_14_70)($v_15);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_13_73 = (($__local_var_2_1)->{'Bind1'})(null);
  $pure_14_74 = ((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'};
  $Bind1_13_73 = (object)["bind" => function($v_15) use ($Bind1_13_73, $pure_14_74) {
  $__num = \func_num_args();
  $__res = function($k_16) use ($Bind1_13_73, $pure_14_74, $v_15) {
  $__num = \func_num_args();
  $__local_var_17_75 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_14_74))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_13_73)->{'bind'})($v_15))(function($v2_18) use ($__local_var_17_75, $k_16) {
  $__num = \func_num_args();
  $__t76 = null;;
  if ($v2_18 instanceof \Data\Either\Data_Either_Left) {
$__t76 = ($__local_var_17_75)(($v2_18)->{'value0'});
goto end_branch_76;;
};
  if ($v2_18 instanceof \Data\Either\Data_Either_Right) {
$__t76 = ($k_16)(($v2_18)->{'value0'});
goto end_branch_76;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t76 = null;
  end_branch_76:;
  $__res = $__t76;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_15) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($__local_var_2_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_14_78 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($__local_var_2_1);
  $__res = (object)["apply" => function($f_15) use ($Applicative0_14_78, $Bind1_13_73) {
  $__num = \func_num_args();
  $__res = function($a_16) use ($Applicative0_14_78, $Bind1_13_73, $f_15) {
  $__num = \func_num_args();
  $__res = ((($Bind1_13_73)->{'bind'})($f_15))(function($f_prime__17) use ($Applicative0_14_78, $Bind1_13_73, $a_16) {
  $__num = \func_num_args();
  $__res = ((($Bind1_13_73)->{'bind'})($a_16))(function($a_prime__18) use ($Applicative0_14_78, $f_prime__17) {
  $__num = \func_num_args();
  $__res = (($Applicative0_14_78)->{'pure'})(($f_prime__17)($a_prime__18));
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
}, "Functor0" => function($_dollar___unused_13) use ($functorExceptT1_12_69) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_12_69;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_10_80 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($__local_var_2_1);
  $__res = (object)["apply" => function($f_11) use ($Applicative0_10_80, $Bind1_9_65) {
  $__num = \func_num_args();
  $__res = function($a_12) use ($Applicative0_10_80, $Bind1_9_65, $f_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_65)->{'bind'})($f_11))(function($f_prime__13) use ($Applicative0_10_80, $Bind1_9_65, $a_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_65)->{'bind'})($a_12))(function($a_prime__14) use ($Applicative0_10_80, $f_prime__13) {
  $__num = \func_num_args();
  $__res = (($Applicative0_10_80)->{'pure'})(($f_prime__13)($a_prime__14));
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
}, "Functor0" => function($_dollar___unused_9) use ($functorExceptT1_8_61) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_8_61;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_8) use ($Applicative0_7_61, $Bind1_6_6) {
  $__num = \func_num_args();
  $__res = function($a_9) use ($Applicative0_7_61, $Bind1_6_6, $f_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_6)->{'bind'})($f_8))(function($f_prime__10) use ($Applicative0_7_61, $Bind1_6_6, $a_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_6)->{'bind'})($a_9))(function($a_prime__11) use ($Applicative0_7_61, $f_prime__10) {
  $__num = \func_num_args();
  $__res = (($Applicative0_7_61)->{'pure'})(($f_prime__10)($a_prime__11));
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
}, "Functor0" => function($_dollar___unused_6) use ($functorExceptT1_5_2) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_5_2;
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
  $Bind1_4_82 = (($__local_var_2_1)->{'Bind1'})(null);
  $pure_5_83 = ((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'};
  $__res = (object)["bind" => function($v_6) use ($Bind1_4_82, $pure_5_83) {
  $__num = \func_num_args();
  $__res = function($k_7) use ($Bind1_4_82, $pure_5_83, $v_6) {
  $__num = \func_num_args();
  $__local_var_8_84 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_5_83))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_4_82)->{'bind'})($v_6))(function($v2_9) use ($__local_var_8_84, $k_7) {
  $__num = \func_num_args();
  $__t85 = null;;
  if ($v2_9 instanceof \Data\Either\Data_Either_Left) {
$__t85 = ($__local_var_8_84)(($v2_9)->{'value0'});
goto end_branch_85;;
};
  if ($v2_9 instanceof \Data\Either\Data_Either_Right) {
$__t85 = ($k_7)(($v2_9)->{'value0'});
goto end_branch_85;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t85 = null;
  end_branch_85:;
  $__res = $__t85;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_6) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_7_86 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_7_86 = (object)["map" => function($f_8) use ($__local_var_7_86) {
  $__num = \func_num_args();
  $__local_var_9_87 = (($__local_var_7_86)->{'map'})(function($m_9) use ($f_8) {
  $__num = \func_num_args();
  $__t87 = null;;
  if ($m_9 instanceof \Data\Either\Data_Either_Left) {
$__t87 = new \Data\Either\Data_Either_Left(($m_9)->{'value0'});
goto end_branch_87;;
};
  if ($m_9 instanceof \Data\Either\Data_Either_Right) {
$__t87 = new \Data\Either\Data_Either_Right(($f_8)(($m_9)->{'value0'}));
goto end_branch_87;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t87 = null;
  end_branch_87:;
  $__res = $__t87;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_10) use ($__local_var_9_87) {
  $__num = \func_num_args();
  $__res = ($__local_var_9_87)($v_10);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_8_90 = (($__local_var_2_1)->{'Bind1'})(null);
  $pure_9_91 = ((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'};
  $Bind1_8_90 = (object)["bind" => function($v_10) use ($Bind1_8_90, $pure_9_91) {
  $__num = \func_num_args();
  $__res = function($k_11) use ($Bind1_8_90, $pure_9_91, $v_10) {
  $__num = \func_num_args();
  $__local_var_12_92 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_9_91))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_8_90)->{'bind'})($v_10))(function($v2_13) use ($__local_var_12_92, $k_11) {
  $__num = \func_num_args();
  $__t93 = null;;
  if ($v2_13 instanceof \Data\Either\Data_Either_Left) {
$__t93 = ($__local_var_12_92)(($v2_13)->{'value0'});
goto end_branch_93;;
};
  if ($v2_13 instanceof \Data\Either\Data_Either_Right) {
$__t93 = ($k_11)(($v2_13)->{'value0'});
goto end_branch_93;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t93 = null;
  end_branch_93:;
  $__res = $__t93;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_10) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_11_94 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_11_94 = (object)["map" => function($f_12) use ($__local_var_11_94) {
  $__num = \func_num_args();
  $__local_var_13_95 = (($__local_var_11_94)->{'map'})(function($m_13) use ($f_12) {
  $__num = \func_num_args();
  $__t95 = null;;
  if ($m_13 instanceof \Data\Either\Data_Either_Left) {
$__t95 = new \Data\Either\Data_Either_Left(($m_13)->{'value0'});
goto end_branch_95;;
};
  if ($m_13 instanceof \Data\Either\Data_Either_Right) {
$__t95 = new \Data\Either\Data_Either_Right(($f_12)(($m_13)->{'value0'}));
goto end_branch_95;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t95 = null;
  end_branch_95:;
  $__res = $__t95;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_14) use ($__local_var_13_95) {
  $__num = \func_num_args();
  $__res = ($__local_var_13_95)($v_14);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_12_98 = (($__local_var_2_1)->{'Bind1'})(null);
  $pure_13_99 = ((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'};
  $Bind1_12_98 = (object)["bind" => function($v_14) use ($Bind1_12_98, $pure_13_99) {
  $__num = \func_num_args();
  $__res = function($k_15) use ($Bind1_12_98, $pure_13_99, $v_14) {
  $__num = \func_num_args();
  $__local_var_16_100 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_13_99))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_12_98)->{'bind'})($v_14))(function($v2_17) use ($__local_var_16_100, $k_15) {
  $__num = \func_num_args();
  $__t101 = null;;
  if ($v2_17 instanceof \Data\Either\Data_Either_Left) {
$__t101 = ($__local_var_16_100)(($v2_17)->{'value0'});
goto end_branch_101;;
};
  if ($v2_17 instanceof \Data\Either\Data_Either_Right) {
$__t101 = ($k_15)(($v2_17)->{'value0'});
goto end_branch_101;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t101 = null;
  end_branch_101:;
  $__res = $__t101;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_14) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($__local_var_2_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_13_103 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_13) {
  $__num = \func_num_args();
  $__res = $x_13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_13) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_14_103 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_14_103 = (object)["map" => function($f_15) use ($__local_var_14_103) {
  $__num = \func_num_args();
  $__local_var_16_104 = (($__local_var_14_103)->{'map'})(function($m_16) use ($f_15) {
  $__num = \func_num_args();
  $__t104 = null;;
  if ($m_16 instanceof \Data\Either\Data_Either_Left) {
$__t104 = new \Data\Either\Data_Either_Left(($m_16)->{'value0'});
goto end_branch_104;;
};
  if ($m_16 instanceof \Data\Either\Data_Either_Right) {
$__t104 = new \Data\Either\Data_Either_Right(($f_15)(($m_16)->{'value0'}));
goto end_branch_104;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t104 = null;
  end_branch_104:;
  $__res = $__t104;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_17) use ($__local_var_16_104) {
  $__num = \func_num_args();
  $__res = ($__local_var_16_104)($v_17);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_15_107 = (($__local_var_2_1)->{'Bind1'})(null);
  $pure_16_108 = ((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'};
  $Bind1_15_107 = (object)["bind" => function($v_17) use ($Bind1_15_107, $pure_16_108) {
  $__num = \func_num_args();
  $__res = function($k_18) use ($Bind1_15_107, $pure_16_108, $v_17) {
  $__num = \func_num_args();
  $__local_var_19_109 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_16_108))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_15_107)->{'bind'})($v_17))(function($v2_20) use ($__local_var_19_109, $k_18) {
  $__num = \func_num_args();
  $__t110 = null;;
  if ($v2_20 instanceof \Data\Either\Data_Either_Left) {
$__t110 = ($__local_var_19_109)(($v2_20)->{'value0'});
goto end_branch_110;;
};
  if ($v2_20 instanceof \Data\Either\Data_Either_Right) {
$__t110 = ($k_18)(($v2_20)->{'value0'});
goto end_branch_110;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t110 = null;
  end_branch_110:;
  $__res = $__t110;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_17) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($__local_var_2_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_16_112 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($__local_var_2_1);
  $__res = (object)["apply" => function($f_17) use ($Applicative0_16_112, $Bind1_15_107) {
  $__num = \func_num_args();
  $__res = function($a_18) use ($Applicative0_16_112, $Bind1_15_107, $f_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_107)->{'bind'})($f_17))(function($f_prime__19) use ($Applicative0_16_112, $Bind1_15_107, $a_18) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_107)->{'bind'})($a_18))(function($a_prime__20) use ($Applicative0_16_112, $f_prime__19) {
  $__num = \func_num_args();
  $__res = (($Applicative0_16_112)->{'pure'})(($f_prime__19)($a_prime__20));
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
}, "Functor0" => function($_dollar___unused_15) use ($functorExceptT1_14_103) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_14_103;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_14) use ($Applicative0_13_103, $Bind1_12_98) {
  $__num = \func_num_args();
  $__res = function($a_15) use ($Applicative0_13_103, $Bind1_12_98, $f_14) {
  $__num = \func_num_args();
  $__res = ((($Bind1_12_98)->{'bind'})($f_14))(function($f_prime__16) use ($Applicative0_13_103, $Bind1_12_98, $a_15) {
  $__num = \func_num_args();
  $__res = ((($Bind1_12_98)->{'bind'})($a_15))(function($a_prime__17) use ($Applicative0_13_103, $f_prime__16) {
  $__num = \func_num_args();
  $__res = (($Applicative0_13_103)->{'pure'})(($f_prime__16)($a_prime__17));
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
}, "Functor0" => function($_dollar___unused_12) use ($functorExceptT1_11_94) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_11_94;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_9_115 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_9) {
  $__num = \func_num_args();
  $__res = $x_9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_9) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_10_115 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_10_115 = (object)["map" => function($f_11) use ($__local_var_10_115) {
  $__num = \func_num_args();
  $__local_var_12_116 = (($__local_var_10_115)->{'map'})(function($m_12) use ($f_11) {
  $__num = \func_num_args();
  $__t116 = null;;
  if ($m_12 instanceof \Data\Either\Data_Either_Left) {
$__t116 = new \Data\Either\Data_Either_Left(($m_12)->{'value0'});
goto end_branch_116;;
};
  if ($m_12 instanceof \Data\Either\Data_Either_Right) {
$__t116 = new \Data\Either\Data_Either_Right(($f_11)(($m_12)->{'value0'}));
goto end_branch_116;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t116 = null;
  end_branch_116:;
  $__res = $__t116;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_13) use ($__local_var_12_116) {
  $__num = \func_num_args();
  $__res = ($__local_var_12_116)($v_13);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_11_119 = (($__local_var_2_1)->{'Bind1'})(null);
  $pure_12_120 = ((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'};
  $Bind1_11_119 = (object)["bind" => function($v_13) use ($Bind1_11_119, $pure_12_120) {
  $__num = \func_num_args();
  $__res = function($k_14) use ($Bind1_11_119, $pure_12_120, $v_13) {
  $__num = \func_num_args();
  $__local_var_15_121 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_12_120))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_11_119)->{'bind'})($v_13))(function($v2_16) use ($__local_var_15_121, $k_14) {
  $__num = \func_num_args();
  $__t122 = null;;
  if ($v2_16 instanceof \Data\Either\Data_Either_Left) {
$__t122 = ($__local_var_15_121)(($v2_16)->{'value0'});
goto end_branch_122;;
};
  if ($v2_16 instanceof \Data\Either\Data_Either_Right) {
$__t122 = ($k_14)(($v2_16)->{'value0'});
goto end_branch_122;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t122 = null;
  end_branch_122:;
  $__res = $__t122;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_13) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_14_123 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_14_123 = (object)["map" => function($f_15) use ($__local_var_14_123) {
  $__num = \func_num_args();
  $__local_var_16_124 = (($__local_var_14_123)->{'map'})(function($m_16) use ($f_15) {
  $__num = \func_num_args();
  $__t124 = null;;
  if ($m_16 instanceof \Data\Either\Data_Either_Left) {
$__t124 = new \Data\Either\Data_Either_Left(($m_16)->{'value0'});
goto end_branch_124;;
};
  if ($m_16 instanceof \Data\Either\Data_Either_Right) {
$__t124 = new \Data\Either\Data_Either_Right(($f_15)(($m_16)->{'value0'}));
goto end_branch_124;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t124 = null;
  end_branch_124:;
  $__res = $__t124;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_17) use ($__local_var_16_124) {
  $__num = \func_num_args();
  $__res = ($__local_var_16_124)($v_17);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_15_127 = (($__local_var_2_1)->{'Bind1'})(null);
  $pure_16_128 = ((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'};
  $Bind1_15_127 = (object)["bind" => function($v_17) use ($Bind1_15_127, $pure_16_128) {
  $__num = \func_num_args();
  $__res = function($k_18) use ($Bind1_15_127, $pure_16_128, $v_17) {
  $__num = \func_num_args();
  $__local_var_19_129 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_16_128))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_15_127)->{'bind'})($v_17))(function($v2_20) use ($__local_var_19_129, $k_18) {
  $__num = \func_num_args();
  $__t130 = null;;
  if ($v2_20 instanceof \Data\Either\Data_Either_Left) {
$__t130 = ($__local_var_19_129)(($v2_20)->{'value0'});
goto end_branch_130;;
};
  if ($v2_20 instanceof \Data\Either\Data_Either_Right) {
$__t130 = ($k_18)(($v2_20)->{'value0'});
goto end_branch_130;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t130 = null;
  end_branch_130:;
  $__res = $__t130;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_17) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($__local_var_2_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_16_132 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_16) {
  $__num = \func_num_args();
  $__res = $x_16;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_16) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_17_132 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_17_132 = (object)["map" => function($f_18) use ($__local_var_17_132) {
  $__num = \func_num_args();
  $__local_var_19_133 = (($__local_var_17_132)->{'map'})(function($m_19) use ($f_18) {
  $__num = \func_num_args();
  $__t133 = null;;
  if ($m_19 instanceof \Data\Either\Data_Either_Left) {
$__t133 = new \Data\Either\Data_Either_Left(($m_19)->{'value0'});
goto end_branch_133;;
};
  if ($m_19 instanceof \Data\Either\Data_Either_Right) {
$__t133 = new \Data\Either\Data_Either_Right(($f_18)(($m_19)->{'value0'}));
goto end_branch_133;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t133 = null;
  end_branch_133:;
  $__res = $__t133;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_20) use ($__local_var_19_133) {
  $__num = \func_num_args();
  $__res = ($__local_var_19_133)($v_20);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_18_136 = (($__local_var_2_1)->{'Bind1'})(null);
  $pure_19_137 = ((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'};
  $Bind1_18_136 = (object)["bind" => function($v_20) use ($Bind1_18_136, $pure_19_137) {
  $__num = \func_num_args();
  $__res = function($k_21) use ($Bind1_18_136, $pure_19_137, $v_20) {
  $__num = \func_num_args();
  $__local_var_22_138 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_19_137))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_18_136)->{'bind'})($v_20))(function($v2_23) use ($__local_var_22_138, $k_21) {
  $__num = \func_num_args();
  $__t139 = null;;
  if ($v2_23 instanceof \Data\Either\Data_Either_Left) {
$__t139 = ($__local_var_22_138)(($v2_23)->{'value0'});
goto end_branch_139;;
};
  if ($v2_23 instanceof \Data\Either\Data_Either_Right) {
$__t139 = ($k_21)(($v2_23)->{'value0'});
goto end_branch_139;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t139 = null;
  end_branch_139:;
  $__res = $__t139;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_20) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($__local_var_2_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_19_141 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($__local_var_2_1);
  $__res = (object)["apply" => function($f_20) use ($Applicative0_19_141, $Bind1_18_136) {
  $__num = \func_num_args();
  $__res = function($a_21) use ($Applicative0_19_141, $Bind1_18_136, $f_20) {
  $__num = \func_num_args();
  $__res = ((($Bind1_18_136)->{'bind'})($f_20))(function($f_prime__22) use ($Applicative0_19_141, $Bind1_18_136, $a_21) {
  $__num = \func_num_args();
  $__res = ((($Bind1_18_136)->{'bind'})($a_21))(function($a_prime__23) use ($Applicative0_19_141, $f_prime__22) {
  $__num = \func_num_args();
  $__res = (($Applicative0_19_141)->{'pure'})(($f_prime__22)($a_prime__23));
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
}, "Functor0" => function($_dollar___unused_18) use ($functorExceptT1_17_132) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_17_132;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_17) use ($Applicative0_16_132, $Bind1_15_127) {
  $__num = \func_num_args();
  $__res = function($a_18) use ($Applicative0_16_132, $Bind1_15_127, $f_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_127)->{'bind'})($f_17))(function($f_prime__19) use ($Applicative0_16_132, $Bind1_15_127, $a_18) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_127)->{'bind'})($a_18))(function($a_prime__20) use ($Applicative0_16_132, $f_prime__19) {
  $__num = \func_num_args();
  $__res = (($Applicative0_16_132)->{'pure'})(($f_prime__19)($a_prime__20));
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
}, "Functor0" => function($_dollar___unused_15) use ($functorExceptT1_14_123) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_14_123;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_12_144 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_12) {
  $__num = \func_num_args();
  $__res = $x_12;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_12) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_13_144 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_13_144 = (object)["map" => function($f_14) use ($__local_var_13_144) {
  $__num = \func_num_args();
  $__local_var_15_145 = (($__local_var_13_144)->{'map'})(function($m_15) use ($f_14) {
  $__num = \func_num_args();
  $__t145 = null;;
  if ($m_15 instanceof \Data\Either\Data_Either_Left) {
$__t145 = new \Data\Either\Data_Either_Left(($m_15)->{'value0'});
goto end_branch_145;;
};
  if ($m_15 instanceof \Data\Either\Data_Either_Right) {
$__t145 = new \Data\Either\Data_Either_Right(($f_14)(($m_15)->{'value0'}));
goto end_branch_145;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t145 = null;
  end_branch_145:;
  $__res = $__t145;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_16) use ($__local_var_15_145) {
  $__num = \func_num_args();
  $__res = ($__local_var_15_145)($v_16);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_14_148 = (($__local_var_2_1)->{'Bind1'})(null);
  $pure_15_149 = ((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'};
  $Bind1_14_148 = (object)["bind" => function($v_16) use ($Bind1_14_148, $pure_15_149) {
  $__num = \func_num_args();
  $__res = function($k_17) use ($Bind1_14_148, $pure_15_149, $v_16) {
  $__num = \func_num_args();
  $__local_var_18_150 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_15_149))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_14_148)->{'bind'})($v_16))(function($v2_19) use ($__local_var_18_150, $k_17) {
  $__num = \func_num_args();
  $__t151 = null;;
  if ($v2_19 instanceof \Data\Either\Data_Either_Left) {
$__t151 = ($__local_var_18_150)(($v2_19)->{'value0'});
goto end_branch_151;;
};
  if ($v2_19 instanceof \Data\Either\Data_Either_Right) {
$__t151 = ($k_17)(($v2_19)->{'value0'});
goto end_branch_151;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t151 = null;
  end_branch_151:;
  $__res = $__t151;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_16) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_17_152 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_17_152 = (object)["map" => function($f_18) use ($__local_var_17_152) {
  $__num = \func_num_args();
  $__local_var_19_153 = (($__local_var_17_152)->{'map'})(function($m_19) use ($f_18) {
  $__num = \func_num_args();
  $__t153 = null;;
  if ($m_19 instanceof \Data\Either\Data_Either_Left) {
$__t153 = new \Data\Either\Data_Either_Left(($m_19)->{'value0'});
goto end_branch_153;;
};
  if ($m_19 instanceof \Data\Either\Data_Either_Right) {
$__t153 = new \Data\Either\Data_Either_Right(($f_18)(($m_19)->{'value0'}));
goto end_branch_153;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t153 = null;
  end_branch_153:;
  $__res = $__t153;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_20) use ($__local_var_19_153) {
  $__num = \func_num_args();
  $__res = ($__local_var_19_153)($v_20);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_18_156 = (($__local_var_2_1)->{'Bind1'})(null);
  $pure_19_157 = ((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'};
  $Bind1_18_156 = (object)["bind" => function($v_20) use ($Bind1_18_156, $pure_19_157) {
  $__num = \func_num_args();
  $__res = function($k_21) use ($Bind1_18_156, $pure_19_157, $v_20) {
  $__num = \func_num_args();
  $__local_var_22_158 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_19_157))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_18_156)->{'bind'})($v_20))(function($v2_23) use ($__local_var_22_158, $k_21) {
  $__num = \func_num_args();
  $__t159 = null;;
  if ($v2_23 instanceof \Data\Either\Data_Either_Left) {
$__t159 = ($__local_var_22_158)(($v2_23)->{'value0'});
goto end_branch_159;;
};
  if ($v2_23 instanceof \Data\Either\Data_Either_Right) {
$__t159 = ($k_21)(($v2_23)->{'value0'});
goto end_branch_159;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t159 = null;
  end_branch_159:;
  $__res = $__t159;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_20) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($__local_var_2_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_19_161 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($__local_var_2_1);
  $__res = (object)["apply" => function($f_20) use ($Applicative0_19_161, $Bind1_18_156) {
  $__num = \func_num_args();
  $__res = function($a_21) use ($Applicative0_19_161, $Bind1_18_156, $f_20) {
  $__num = \func_num_args();
  $__res = ((($Bind1_18_156)->{'bind'})($f_20))(function($f_prime__22) use ($Applicative0_19_161, $Bind1_18_156, $a_21) {
  $__num = \func_num_args();
  $__res = ((($Bind1_18_156)->{'bind'})($a_21))(function($a_prime__23) use ($Applicative0_19_161, $f_prime__22) {
  $__num = \func_num_args();
  $__res = (($Applicative0_19_161)->{'pure'})(($f_prime__22)($a_prime__23));
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
}, "Functor0" => function($_dollar___unused_18) use ($functorExceptT1_17_152) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_17_152;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_15_163 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($__local_var_2_1);
  $__res = (object)["apply" => function($f_16) use ($Applicative0_15_163, $Bind1_14_148) {
  $__num = \func_num_args();
  $__res = function($a_17) use ($Applicative0_15_163, $Bind1_14_148, $f_16) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_148)->{'bind'})($f_16))(function($f_prime__18) use ($Applicative0_15_163, $Bind1_14_148, $a_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_148)->{'bind'})($a_17))(function($a_prime__19) use ($Applicative0_15_163, $f_prime__18) {
  $__num = \func_num_args();
  $__res = (($Applicative0_15_163)->{'pure'})(($f_prime__18)($a_prime__19));
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
}, "Functor0" => function($_dollar___unused_14) use ($functorExceptT1_13_144) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_13_144;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_13) use ($Applicative0_12_144, $Bind1_11_119) {
  $__num = \func_num_args();
  $__res = function($a_14) use ($Applicative0_12_144, $Bind1_11_119, $f_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_119)->{'bind'})($f_13))(function($f_prime__15) use ($Applicative0_12_144, $Bind1_11_119, $a_14) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_119)->{'bind'})($a_14))(function($a_prime__16) use ($Applicative0_12_144, $f_prime__15) {
  $__num = \func_num_args();
  $__res = (($Applicative0_12_144)->{'pure'})(($f_prime__15)($a_prime__16));
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
}, "Functor0" => function($_dollar___unused_11) use ($functorExceptT1_10_115) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_10_115;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_10) use ($Applicative0_9_115, $Bind1_8_90) {
  $__num = \func_num_args();
  $__res = function($a_11) use ($Applicative0_9_115, $Bind1_8_90, $f_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_8_90)->{'bind'})($f_10))(function($f_prime__12) use ($Applicative0_9_115, $Bind1_8_90, $a_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_8_90)->{'bind'})($a_11))(function($a_prime__13) use ($Applicative0_9_115, $f_prime__12) {
  $__num = \func_num_args();
  $__res = (($Applicative0_9_115)->{'pure'})(($f_prime__12)($a_prime__13));
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
}, "Functor0" => function($_dollar___unused_8) use ($functorExceptT1_7_86) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_7_86;
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
  $__local_var_3_167 = (($__local_var_1_0)->{'Monad0'})(null);
  $pure_4_168 = ((($__local_var_3_167)->{'Applicative0'})(null))->{'pure'};
  $monadAskExceptT1_1_0 = (object)["ask" => ((((($__local_var_3_167)->{'Bind1'})(null))->{'bind'})(($__local_var_1_0)->{'ask'}))(function($a_5) use ($pure_4_168) {
  $__num = \func_num_args();
  $__res = ($pure_4_168)(new \Data\Either\Data_Either_Right($a_5));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), "Monad0" => function($_dollar___unused_3) use ($monadExceptT1_2_1) {
  $__num = \func_num_args();
  $__res = $monadExceptT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["local" => function($f_2) use ($dictMonadReader_0) {
  $__num = \func_num_args();
  $__local_var_3_170 = (($dictMonadReader_0)->{'local'})($f_2);
  $__res = function($v_4) use ($__local_var_3_170) {
  $__num = \func_num_args();
  $__res = ($__local_var_3_170)($v_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "MonadAsk0" => function($_dollar___unused_2) use ($monadAskExceptT1_1_0) {
  $__num = \func_num_args();
  $__res = $monadAskExceptT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Except_Trans_monadReaderExceptT'] = __NAMESPACE__ . '\\majControl_majMonad_majExcept_majTrans_monadmajReadermajExceptmajT';

// Control_Monad_Except_Trans_monadContExceptT
function majControl_majMonad_majExcept_majTrans_monadmajContmajExceptmajT($dictMonadCont_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majExcept_majTrans_monadmajContmajExceptmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictMonadCont_0)->{'Monad0'})(null);
  $monadExceptT1_1_0 = (object)["Applicative0" => function($_dollar___unused_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_3) {
  $__num = \func_num_args();
  $__res = $x_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($__local_var_1_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_3) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__local_var_4_1 = (((((($__local_var_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_4_1 = (object)["map" => function($f_5) use ($__local_var_4_1) {
  $__num = \func_num_args();
  $__local_var_6_2 = (($__local_var_4_1)->{'map'})(function($m_6) use ($f_5) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ($m_6 instanceof \Data\Either\Data_Either_Left) {
$__t2 = new \Data\Either\Data_Either_Left(($m_6)->{'value0'});
goto end_branch_2;;
};
  if ($m_6 instanceof \Data\Either\Data_Either_Right) {
$__t2 = new \Data\Either\Data_Either_Right(($f_5)(($m_6)->{'value0'}));
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_7) use ($__local_var_6_2) {
  $__num = \func_num_args();
  $__res = ($__local_var_6_2)($v_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_5_5 = (($__local_var_1_0)->{'Bind1'})(null);
  $pure_6_6 = ((($__local_var_1_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_5_5 = (object)["bind" => function($v_7) use ($Bind1_5_5, $pure_6_6) {
  $__num = \func_num_args();
  $__res = function($k_8) use ($Bind1_5_5, $pure_6_6, $v_7) {
  $__num = \func_num_args();
  $__local_var_9_7 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_6_6))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_5_5)->{'bind'})($v_7))(function($v2_10) use ($__local_var_9_7, $k_8) {
  $__num = \func_num_args();
  $__t8 = null;;
  if ($v2_10 instanceof \Data\Either\Data_Either_Left) {
$__t8 = ($__local_var_9_7)(($v2_10)->{'value0'});
goto end_branch_8;;
};
  if ($v2_10 instanceof \Data\Either\Data_Either_Right) {
$__t8 = ($k_8)(($v2_10)->{'value0'});
goto end_branch_8;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t8 = null;
  end_branch_8:;
  $__res = $__t8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_7) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__local_var_8_9 = (((((($__local_var_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_8_9 = (object)["map" => function($f_9) use ($__local_var_8_9) {
  $__num = \func_num_args();
  $__local_var_10_10 = (($__local_var_8_9)->{'map'})(function($m_10) use ($f_9) {
  $__num = \func_num_args();
  $__t10 = null;;
  if ($m_10 instanceof \Data\Either\Data_Either_Left) {
$__t10 = new \Data\Either\Data_Either_Left(($m_10)->{'value0'});
goto end_branch_10;;
};
  if ($m_10 instanceof \Data\Either\Data_Either_Right) {
$__t10 = new \Data\Either\Data_Either_Right(($f_9)(($m_10)->{'value0'}));
goto end_branch_10;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t10 = null;
  end_branch_10:;
  $__res = $__t10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_11) use ($__local_var_10_10) {
  $__num = \func_num_args();
  $__res = ($__local_var_10_10)($v_11);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_9_13 = (($__local_var_1_0)->{'Bind1'})(null);
  $pure_10_14 = ((($__local_var_1_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_9_13 = (object)["bind" => function($v_11) use ($Bind1_9_13, $pure_10_14) {
  $__num = \func_num_args();
  $__res = function($k_12) use ($Bind1_9_13, $pure_10_14, $v_11) {
  $__num = \func_num_args();
  $__local_var_13_15 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_10_14))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_9_13)->{'bind'})($v_11))(function($v2_14) use ($__local_var_13_15, $k_12) {
  $__num = \func_num_args();
  $__t16 = null;;
  if ($v2_14 instanceof \Data\Either\Data_Either_Left) {
$__t16 = ($__local_var_13_15)(($v2_14)->{'value0'});
goto end_branch_16;;
};
  if ($v2_14 instanceof \Data\Either\Data_Either_Right) {
$__t16 = ($k_12)(($v2_14)->{'value0'});
goto end_branch_16;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t16 = null;
  end_branch_16:;
  $__res = $__t16;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_11) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($__local_var_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_10_18 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($__local_var_1_0);
  $__res = (object)["apply" => function($f_11) use ($Applicative0_10_18, $Bind1_9_13) {
  $__num = \func_num_args();
  $__res = function($a_12) use ($Applicative0_10_18, $Bind1_9_13, $f_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_13)->{'bind'})($f_11))(function($f_prime__13) use ($Applicative0_10_18, $Bind1_9_13, $a_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_13)->{'bind'})($a_12))(function($a_prime__14) use ($Applicative0_10_18, $f_prime__13) {
  $__num = \func_num_args();
  $__res = (($Applicative0_10_18)->{'pure'})(($f_prime__13)($a_prime__14));
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
}, "Functor0" => function($_dollar___unused_9) use ($functorExceptT1_8_9) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_8_9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_6_20 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($__local_var_1_0);
  $__res = (object)["apply" => function($f_7) use ($Applicative0_6_20, $Bind1_5_5) {
  $__num = \func_num_args();
  $__res = function($a_8) use ($Applicative0_6_20, $Bind1_5_5, $f_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_5_5)->{'bind'})($f_7))(function($f_prime__9) use ($Applicative0_6_20, $Bind1_5_5, $a_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_5_5)->{'bind'})($a_8))(function($a_prime__10) use ($Applicative0_6_20, $f_prime__9) {
  $__num = \func_num_args();
  $__res = (($Applicative0_6_20)->{'pure'})(($f_prime__9)($a_prime__10));
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
}, "Functor0" => function($_dollar___unused_5) use ($functorExceptT1_4_1) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_4_1;
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
  $Bind1_3_21 = (($__local_var_1_0)->{'Bind1'})(null);
  $pure_4_22 = ((($__local_var_1_0)->{'Applicative0'})(null))->{'pure'};
  $__res = (object)["bind" => function($v_5) use ($Bind1_3_21, $pure_4_22) {
  $__num = \func_num_args();
  $__res = function($k_6) use ($Bind1_3_21, $pure_4_22, $v_5) {
  $__num = \func_num_args();
  $__local_var_7_23 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_4_22))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_3_21)->{'bind'})($v_5))(function($v2_8) use ($__local_var_7_23, $k_6) {
  $__num = \func_num_args();
  $__t24 = null;;
  if ($v2_8 instanceof \Data\Either\Data_Either_Left) {
$__t24 = ($__local_var_7_23)(($v2_8)->{'value0'});
goto end_branch_24;;
};
  if ($v2_8 instanceof \Data\Either\Data_Either_Right) {
$__t24 = ($k_6)(($v2_8)->{'value0'});
goto end_branch_24;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t24 = null;
  end_branch_24:;
  $__res = $__t24;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_5) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__local_var_6_25 = (((((($__local_var_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_6_25 = (object)["map" => function($f_7) use ($__local_var_6_25) {
  $__num = \func_num_args();
  $__local_var_8_26 = (($__local_var_6_25)->{'map'})(function($m_8) use ($f_7) {
  $__num = \func_num_args();
  $__t26 = null;;
  if ($m_8 instanceof \Data\Either\Data_Either_Left) {
$__t26 = new \Data\Either\Data_Either_Left(($m_8)->{'value0'});
goto end_branch_26;;
};
  if ($m_8 instanceof \Data\Either\Data_Either_Right) {
$__t26 = new \Data\Either\Data_Either_Right(($f_7)(($m_8)->{'value0'}));
goto end_branch_26;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t26 = null;
  end_branch_26:;
  $__res = $__t26;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_9) use ($__local_var_8_26) {
  $__num = \func_num_args();
  $__res = ($__local_var_8_26)($v_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_7_29 = (($__local_var_1_0)->{'Bind1'})(null);
  $pure_8_30 = ((($__local_var_1_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_7_29 = (object)["bind" => function($v_9) use ($Bind1_7_29, $pure_8_30) {
  $__num = \func_num_args();
  $__res = function($k_10) use ($Bind1_7_29, $pure_8_30, $v_9) {
  $__num = \func_num_args();
  $__local_var_11_31 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_8_30))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_7_29)->{'bind'})($v_9))(function($v2_12) use ($__local_var_11_31, $k_10) {
  $__num = \func_num_args();
  $__t32 = null;;
  if ($v2_12 instanceof \Data\Either\Data_Either_Left) {
$__t32 = ($__local_var_11_31)(($v2_12)->{'value0'});
goto end_branch_32;;
};
  if ($v2_12 instanceof \Data\Either\Data_Either_Right) {
$__t32 = ($k_10)(($v2_12)->{'value0'});
goto end_branch_32;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t32 = null;
  end_branch_32:;
  $__res = $__t32;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_9) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($__local_var_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_8_34 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_8) {
  $__num = \func_num_args();
  $__res = $x_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($__local_var_1_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_8) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__local_var_9_34 = (((((($__local_var_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_9_34 = (object)["map" => function($f_10) use ($__local_var_9_34) {
  $__num = \func_num_args();
  $__local_var_11_35 = (($__local_var_9_34)->{'map'})(function($m_11) use ($f_10) {
  $__num = \func_num_args();
  $__t35 = null;;
  if ($m_11 instanceof \Data\Either\Data_Either_Left) {
$__t35 = new \Data\Either\Data_Either_Left(($m_11)->{'value0'});
goto end_branch_35;;
};
  if ($m_11 instanceof \Data\Either\Data_Either_Right) {
$__t35 = new \Data\Either\Data_Either_Right(($f_10)(($m_11)->{'value0'}));
goto end_branch_35;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t35 = null;
  end_branch_35:;
  $__res = $__t35;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_12) use ($__local_var_11_35) {
  $__num = \func_num_args();
  $__res = ($__local_var_11_35)($v_12);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_10_38 = (($__local_var_1_0)->{'Bind1'})(null);
  $pure_11_39 = ((($__local_var_1_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_10_38 = (object)["bind" => function($v_12) use ($Bind1_10_38, $pure_11_39) {
  $__num = \func_num_args();
  $__res = function($k_13) use ($Bind1_10_38, $pure_11_39, $v_12) {
  $__num = \func_num_args();
  $__local_var_14_40 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_11_39))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_10_38)->{'bind'})($v_12))(function($v2_15) use ($__local_var_14_40, $k_13) {
  $__num = \func_num_args();
  $__t41 = null;;
  if ($v2_15 instanceof \Data\Either\Data_Either_Left) {
$__t41 = ($__local_var_14_40)(($v2_15)->{'value0'});
goto end_branch_41;;
};
  if ($v2_15 instanceof \Data\Either\Data_Either_Right) {
$__t41 = ($k_13)(($v2_15)->{'value0'});
goto end_branch_41;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t41 = null;
  end_branch_41:;
  $__res = $__t41;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_12) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($__local_var_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_11_43 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($__local_var_1_0);
  $__res = (object)["apply" => function($f_12) use ($Applicative0_11_43, $Bind1_10_38) {
  $__num = \func_num_args();
  $__res = function($a_13) use ($Applicative0_11_43, $Bind1_10_38, $f_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_38)->{'bind'})($f_12))(function($f_prime__14) use ($Applicative0_11_43, $Bind1_10_38, $a_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_38)->{'bind'})($a_13))(function($a_prime__15) use ($Applicative0_11_43, $f_prime__14) {
  $__num = \func_num_args();
  $__res = (($Applicative0_11_43)->{'pure'})(($f_prime__14)($a_prime__15));
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
}, "Functor0" => function($_dollar___unused_10) use ($functorExceptT1_9_34) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_9_34;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_9) use ($Applicative0_8_34, $Bind1_7_29) {
  $__num = \func_num_args();
  $__res = function($a_10) use ($Applicative0_8_34, $Bind1_7_29, $f_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_29)->{'bind'})($f_9))(function($f_prime__11) use ($Applicative0_8_34, $Bind1_7_29, $a_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_29)->{'bind'})($a_10))(function($a_prime__12) use ($Applicative0_8_34, $f_prime__11) {
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
}, "Functor0" => function($_dollar___unused_7) use ($functorExceptT1_6_25) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_6_25;
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
  $__res = (($dictMonadCont_0)->{'callCC'})(function($c_3) use ($f_2) {
  $__num = \func_num_args();
  $__res = ($f_2)(function($a_4) use ($c_3) {
  $__num = \func_num_args();
  $__res = ($c_3)(new \Data\Either\Data_Either_Right($a_4));
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
}, "Monad0" => function($_dollar___unused_2) use ($monadExceptT1_1_0) {
  $__num = \func_num_args();
  $__res = $monadExceptT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Except_Trans_monadContExceptT'] = __NAMESPACE__ . '\\majControl_majMonad_majExcept_majTrans_monadmajContmajExceptmajT';

// Control_Monad_Except_Trans_monadEffectExceptT
function majControl_majMonad_majExcept_majTrans_monadmajEffectmajExceptmajT($dictMonadEffect_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majExcept_majTrans_monadmajEffectmajExceptmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Monad0_1_0 = (($dictMonadEffect_0)->{'Monad0'})(null);
  $monadExceptT1_2_1 = (object)["Applicative0" => function($_dollar___unused_2) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_3) {
  $__num = \func_num_args();
  $__res = $x_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($Monad0_1_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_3) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__local_var_4_1 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_4_1 = (object)["map" => function($f_5) use ($__local_var_4_1) {
  $__num = \func_num_args();
  $__local_var_6_2 = (($__local_var_4_1)->{'map'})(function($m_6) use ($f_5) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ($m_6 instanceof \Data\Either\Data_Either_Left) {
$__t2 = new \Data\Either\Data_Either_Left(($m_6)->{'value0'});
goto end_branch_2;;
};
  if ($m_6 instanceof \Data\Either\Data_Either_Right) {
$__t2 = new \Data\Either\Data_Either_Right(($f_5)(($m_6)->{'value0'}));
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_7) use ($__local_var_6_2) {
  $__num = \func_num_args();
  $__res = ($__local_var_6_2)($v_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_5_5 = (($Monad0_1_0)->{'Bind1'})(null);
  $pure_6_6 = ((($Monad0_1_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_5_5 = (object)["bind" => function($v_7) use ($Bind1_5_5, $pure_6_6) {
  $__num = \func_num_args();
  $__res = function($k_8) use ($Bind1_5_5, $pure_6_6, $v_7) {
  $__num = \func_num_args();
  $__local_var_9_7 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_6_6))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_5_5)->{'bind'})($v_7))(function($v2_10) use ($__local_var_9_7, $k_8) {
  $__num = \func_num_args();
  $__t8 = null;;
  if ($v2_10 instanceof \Data\Either\Data_Either_Left) {
$__t8 = ($__local_var_9_7)(($v2_10)->{'value0'});
goto end_branch_8;;
};
  if ($v2_10 instanceof \Data\Either\Data_Either_Right) {
$__t8 = ($k_8)(($v2_10)->{'value0'});
goto end_branch_8;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t8 = null;
  end_branch_8:;
  $__res = $__t8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_7) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__local_var_8_9 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_8_9 = (object)["map" => function($f_9) use ($__local_var_8_9) {
  $__num = \func_num_args();
  $__local_var_10_10 = (($__local_var_8_9)->{'map'})(function($m_10) use ($f_9) {
  $__num = \func_num_args();
  $__t10 = null;;
  if ($m_10 instanceof \Data\Either\Data_Either_Left) {
$__t10 = new \Data\Either\Data_Either_Left(($m_10)->{'value0'});
goto end_branch_10;;
};
  if ($m_10 instanceof \Data\Either\Data_Either_Right) {
$__t10 = new \Data\Either\Data_Either_Right(($f_9)(($m_10)->{'value0'}));
goto end_branch_10;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t10 = null;
  end_branch_10:;
  $__res = $__t10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_11) use ($__local_var_10_10) {
  $__num = \func_num_args();
  $__res = ($__local_var_10_10)($v_11);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_9_13 = (($Monad0_1_0)->{'Bind1'})(null);
  $pure_10_14 = ((($Monad0_1_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_9_13 = (object)["bind" => function($v_11) use ($Bind1_9_13, $pure_10_14) {
  $__num = \func_num_args();
  $__res = function($k_12) use ($Bind1_9_13, $pure_10_14, $v_11) {
  $__num = \func_num_args();
  $__local_var_13_15 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_10_14))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_9_13)->{'bind'})($v_11))(function($v2_14) use ($__local_var_13_15, $k_12) {
  $__num = \func_num_args();
  $__t16 = null;;
  if ($v2_14 instanceof \Data\Either\Data_Either_Left) {
$__t16 = ($__local_var_13_15)(($v2_14)->{'value0'});
goto end_branch_16;;
};
  if ($v2_14 instanceof \Data\Either\Data_Either_Right) {
$__t16 = ($k_12)(($v2_14)->{'value0'});
goto end_branch_16;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t16 = null;
  end_branch_16:;
  $__res = $__t16;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_11) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_10_18 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($Monad0_1_0);
  $__res = (object)["apply" => function($f_11) use ($Applicative0_10_18, $Bind1_9_13) {
  $__num = \func_num_args();
  $__res = function($a_12) use ($Applicative0_10_18, $Bind1_9_13, $f_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_13)->{'bind'})($f_11))(function($f_prime__13) use ($Applicative0_10_18, $Bind1_9_13, $a_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_13)->{'bind'})($a_12))(function($a_prime__14) use ($Applicative0_10_18, $f_prime__13) {
  $__num = \func_num_args();
  $__res = (($Applicative0_10_18)->{'pure'})(($f_prime__13)($a_prime__14));
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
}, "Functor0" => function($_dollar___unused_9) use ($functorExceptT1_8_9) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_8_9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_6_20 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($Monad0_1_0);
  $__res = (object)["apply" => function($f_7) use ($Applicative0_6_20, $Bind1_5_5) {
  $__num = \func_num_args();
  $__res = function($a_8) use ($Applicative0_6_20, $Bind1_5_5, $f_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_5_5)->{'bind'})($f_7))(function($f_prime__9) use ($Applicative0_6_20, $Bind1_5_5, $a_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_5_5)->{'bind'})($a_8))(function($a_prime__10) use ($Applicative0_6_20, $f_prime__9) {
  $__num = \func_num_args();
  $__res = (($Applicative0_6_20)->{'pure'})(($f_prime__9)($a_prime__10));
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
}, "Functor0" => function($_dollar___unused_5) use ($functorExceptT1_4_1) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_4_1;
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
  $Bind1_3_21 = (($Monad0_1_0)->{'Bind1'})(null);
  $pure_4_22 = ((($Monad0_1_0)->{'Applicative0'})(null))->{'pure'};
  $__res = (object)["bind" => function($v_5) use ($Bind1_3_21, $pure_4_22) {
  $__num = \func_num_args();
  $__res = function($k_6) use ($Bind1_3_21, $pure_4_22, $v_5) {
  $__num = \func_num_args();
  $__local_var_7_23 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_4_22))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_3_21)->{'bind'})($v_5))(function($v2_8) use ($__local_var_7_23, $k_6) {
  $__num = \func_num_args();
  $__t24 = null;;
  if ($v2_8 instanceof \Data\Either\Data_Either_Left) {
$__t24 = ($__local_var_7_23)(($v2_8)->{'value0'});
goto end_branch_24;;
};
  if ($v2_8 instanceof \Data\Either\Data_Either_Right) {
$__t24 = ($k_6)(($v2_8)->{'value0'});
goto end_branch_24;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t24 = null;
  end_branch_24:;
  $__res = $__t24;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_5) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__local_var_6_25 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_6_25 = (object)["map" => function($f_7) use ($__local_var_6_25) {
  $__num = \func_num_args();
  $__local_var_8_26 = (($__local_var_6_25)->{'map'})(function($m_8) use ($f_7) {
  $__num = \func_num_args();
  $__t26 = null;;
  if ($m_8 instanceof \Data\Either\Data_Either_Left) {
$__t26 = new \Data\Either\Data_Either_Left(($m_8)->{'value0'});
goto end_branch_26;;
};
  if ($m_8 instanceof \Data\Either\Data_Either_Right) {
$__t26 = new \Data\Either\Data_Either_Right(($f_7)(($m_8)->{'value0'}));
goto end_branch_26;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t26 = null;
  end_branch_26:;
  $__res = $__t26;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_9) use ($__local_var_8_26) {
  $__num = \func_num_args();
  $__res = ($__local_var_8_26)($v_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_7_29 = (($Monad0_1_0)->{'Bind1'})(null);
  $pure_8_30 = ((($Monad0_1_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_7_29 = (object)["bind" => function($v_9) use ($Bind1_7_29, $pure_8_30) {
  $__num = \func_num_args();
  $__res = function($k_10) use ($Bind1_7_29, $pure_8_30, $v_9) {
  $__num = \func_num_args();
  $__local_var_11_31 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_8_30))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_7_29)->{'bind'})($v_9))(function($v2_12) use ($__local_var_11_31, $k_10) {
  $__num = \func_num_args();
  $__t32 = null;;
  if ($v2_12 instanceof \Data\Either\Data_Either_Left) {
$__t32 = ($__local_var_11_31)(($v2_12)->{'value0'});
goto end_branch_32;;
};
  if ($v2_12 instanceof \Data\Either\Data_Either_Right) {
$__t32 = ($k_10)(($v2_12)->{'value0'});
goto end_branch_32;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t32 = null;
  end_branch_32:;
  $__res = $__t32;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_9) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_8_34 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_8) {
  $__num = \func_num_args();
  $__res = $x_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($Monad0_1_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_8) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__local_var_9_34 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_9_34 = (object)["map" => function($f_10) use ($__local_var_9_34) {
  $__num = \func_num_args();
  $__local_var_11_35 = (($__local_var_9_34)->{'map'})(function($m_11) use ($f_10) {
  $__num = \func_num_args();
  $__t35 = null;;
  if ($m_11 instanceof \Data\Either\Data_Either_Left) {
$__t35 = new \Data\Either\Data_Either_Left(($m_11)->{'value0'});
goto end_branch_35;;
};
  if ($m_11 instanceof \Data\Either\Data_Either_Right) {
$__t35 = new \Data\Either\Data_Either_Right(($f_10)(($m_11)->{'value0'}));
goto end_branch_35;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t35 = null;
  end_branch_35:;
  $__res = $__t35;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_12) use ($__local_var_11_35) {
  $__num = \func_num_args();
  $__res = ($__local_var_11_35)($v_12);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_10_38 = (($Monad0_1_0)->{'Bind1'})(null);
  $pure_11_39 = ((($Monad0_1_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_10_38 = (object)["bind" => function($v_12) use ($Bind1_10_38, $pure_11_39) {
  $__num = \func_num_args();
  $__res = function($k_13) use ($Bind1_10_38, $pure_11_39, $v_12) {
  $__num = \func_num_args();
  $__local_var_14_40 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_11_39))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_10_38)->{'bind'})($v_12))(function($v2_15) use ($__local_var_14_40, $k_13) {
  $__num = \func_num_args();
  $__t41 = null;;
  if ($v2_15 instanceof \Data\Either\Data_Either_Left) {
$__t41 = ($__local_var_14_40)(($v2_15)->{'value0'});
goto end_branch_41;;
};
  if ($v2_15 instanceof \Data\Either\Data_Either_Right) {
$__t41 = ($k_13)(($v2_15)->{'value0'});
goto end_branch_41;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t41 = null;
  end_branch_41:;
  $__res = $__t41;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_12) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_11_43 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($Monad0_1_0);
  $__res = (object)["apply" => function($f_12) use ($Applicative0_11_43, $Bind1_10_38) {
  $__num = \func_num_args();
  $__res = function($a_13) use ($Applicative0_11_43, $Bind1_10_38, $f_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_38)->{'bind'})($f_12))(function($f_prime__14) use ($Applicative0_11_43, $Bind1_10_38, $a_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_38)->{'bind'})($a_13))(function($a_prime__15) use ($Applicative0_11_43, $f_prime__14) {
  $__num = \func_num_args();
  $__res = (($Applicative0_11_43)->{'pure'})(($f_prime__14)($a_prime__15));
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
}, "Functor0" => function($_dollar___unused_10) use ($functorExceptT1_9_34) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_9_34;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_9) use ($Applicative0_8_34, $Bind1_7_29) {
  $__num = \func_num_args();
  $__res = function($a_10) use ($Applicative0_8_34, $Bind1_7_29, $f_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_29)->{'bind'})($f_9))(function($f_prime__11) use ($Applicative0_8_34, $Bind1_7_29, $a_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_29)->{'bind'})($a_10))(function($a_prime__12) use ($Applicative0_8_34, $f_prime__11) {
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
}, "Functor0" => function($_dollar___unused_7) use ($functorExceptT1_6_25) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_6_25;
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
  $Bind1_3_46 = (($Monad0_1_0)->{'Bind1'})(null);
  $pure_4_47 = ((($Monad0_1_0)->{'Applicative0'})(null))->{'pure'};
  $__res = (object)["liftEffect" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($m_5) use ($Bind1_3_46, $pure_4_47) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_46)->{'bind'})($m_5))(function($a_6) use ($pure_4_47) {
  $__num = \func_num_args();
  $__res = ($pure_4_47)(new \Data\Either\Data_Either_Right($a_6));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($dictMonadEffect_0)->{'liftEffect'}), "Monad0" => function($_dollar___unused_3) use ($monadExceptT1_2_1) {
  $__num = \func_num_args();
  $__res = $monadExceptT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Except_Trans_monadEffectExceptT'] = __NAMESPACE__ . '\\majControl_majMonad_majExcept_majTrans_monadmajEffectmajExceptmajT';

// Control_Monad_Except_Trans_monadRecExceptT
function majControl_majMonad_majExcept_majTrans_monadmajRecmajExceptmajT($dictMonadRec_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majExcept_majTrans_monadmajRecmajExceptmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Monad0_1_0 = (($dictMonadRec_0)->{'Monad0'})(null);
  $Bind1_2_1 = (($Monad0_1_0)->{'Bind1'})(null);
  $Applicative0_3_2 = (($Monad0_1_0)->{'Applicative0'})(null);
  $monadExceptT1_4_3 = (object)["Applicative0" => function($_dollar___unused_4) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_5) {
  $__num = \func_num_args();
  $__res = $x_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($Monad0_1_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_5) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__local_var_6_3 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_6_3 = (object)["map" => function($f_7) use ($__local_var_6_3) {
  $__num = \func_num_args();
  $__local_var_8_4 = (($__local_var_6_3)->{'map'})(function($m_8) use ($f_7) {
  $__num = \func_num_args();
  $__t4 = null;;
  if ($m_8 instanceof \Data\Either\Data_Either_Left) {
$__t4 = new \Data\Either\Data_Either_Left(($m_8)->{'value0'});
goto end_branch_4;;
};
  if ($m_8 instanceof \Data\Either\Data_Either_Right) {
$__t4 = new \Data\Either\Data_Either_Right(($f_7)(($m_8)->{'value0'}));
goto end_branch_4;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t4 = null;
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_9) use ($__local_var_8_4) {
  $__num = \func_num_args();
  $__res = ($__local_var_8_4)($v_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_7_7 = (($Monad0_1_0)->{'Bind1'})(null);
  $pure_8_8 = ((($Monad0_1_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_7_7 = (object)["bind" => function($v_9) use ($Bind1_7_7, $pure_8_8) {
  $__num = \func_num_args();
  $__res = function($k_10) use ($Bind1_7_7, $pure_8_8, $v_9) {
  $__num = \func_num_args();
  $__local_var_11_9 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_8_8))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_7_7)->{'bind'})($v_9))(function($v2_12) use ($__local_var_11_9, $k_10) {
  $__num = \func_num_args();
  $__t10 = null;;
  if ($v2_12 instanceof \Data\Either\Data_Either_Left) {
$__t10 = ($__local_var_11_9)(($v2_12)->{'value0'});
goto end_branch_10;;
};
  if ($v2_12 instanceof \Data\Either\Data_Either_Right) {
$__t10 = ($k_10)(($v2_12)->{'value0'});
goto end_branch_10;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t10 = null;
  end_branch_10:;
  $__res = $__t10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_9) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__local_var_10_11 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_10_11 = (object)["map" => function($f_11) use ($__local_var_10_11) {
  $__num = \func_num_args();
  $__local_var_12_12 = (($__local_var_10_11)->{'map'})(function($m_12) use ($f_11) {
  $__num = \func_num_args();
  $__t12 = null;;
  if ($m_12 instanceof \Data\Either\Data_Either_Left) {
$__t12 = new \Data\Either\Data_Either_Left(($m_12)->{'value0'});
goto end_branch_12;;
};
  if ($m_12 instanceof \Data\Either\Data_Either_Right) {
$__t12 = new \Data\Either\Data_Either_Right(($f_11)(($m_12)->{'value0'}));
goto end_branch_12;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t12 = null;
  end_branch_12:;
  $__res = $__t12;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_13) use ($__local_var_12_12) {
  $__num = \func_num_args();
  $__res = ($__local_var_12_12)($v_13);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_11_15 = (($Monad0_1_0)->{'Bind1'})(null);
  $pure_12_16 = ((($Monad0_1_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_11_15 = (object)["bind" => function($v_13) use ($Bind1_11_15, $pure_12_16) {
  $__num = \func_num_args();
  $__res = function($k_14) use ($Bind1_11_15, $pure_12_16, $v_13) {
  $__num = \func_num_args();
  $__local_var_15_17 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_12_16))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_11_15)->{'bind'})($v_13))(function($v2_16) use ($__local_var_15_17, $k_14) {
  $__num = \func_num_args();
  $__t18 = null;;
  if ($v2_16 instanceof \Data\Either\Data_Either_Left) {
$__t18 = ($__local_var_15_17)(($v2_16)->{'value0'});
goto end_branch_18;;
};
  if ($v2_16 instanceof \Data\Either\Data_Either_Right) {
$__t18 = ($k_14)(($v2_16)->{'value0'});
goto end_branch_18;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t18 = null;
  end_branch_18:;
  $__res = $__t18;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_13) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_12_20 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($Monad0_1_0);
  $__res = (object)["apply" => function($f_13) use ($Applicative0_12_20, $Bind1_11_15) {
  $__num = \func_num_args();
  $__res = function($a_14) use ($Applicative0_12_20, $Bind1_11_15, $f_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_15)->{'bind'})($f_13))(function($f_prime__15) use ($Applicative0_12_20, $Bind1_11_15, $a_14) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_15)->{'bind'})($a_14))(function($a_prime__16) use ($Applicative0_12_20, $f_prime__15) {
  $__num = \func_num_args();
  $__res = (($Applicative0_12_20)->{'pure'})(($f_prime__15)($a_prime__16));
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
}, "Functor0" => function($_dollar___unused_11) use ($functorExceptT1_10_11) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_10_11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_8_22 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($Monad0_1_0);
  $__res = (object)["apply" => function($f_9) use ($Applicative0_8_22, $Bind1_7_7) {
  $__num = \func_num_args();
  $__res = function($a_10) use ($Applicative0_8_22, $Bind1_7_7, $f_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_7)->{'bind'})($f_9))(function($f_prime__11) use ($Applicative0_8_22, $Bind1_7_7, $a_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_7)->{'bind'})($a_10))(function($a_prime__12) use ($Applicative0_8_22, $f_prime__11) {
  $__num = \func_num_args();
  $__res = (($Applicative0_8_22)->{'pure'})(($f_prime__11)($a_prime__12));
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
}, "Functor0" => function($_dollar___unused_7) use ($functorExceptT1_6_3) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_6_3;
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
  $Bind1_5_23 = (($Monad0_1_0)->{'Bind1'})(null);
  $pure_6_24 = ((($Monad0_1_0)->{'Applicative0'})(null))->{'pure'};
  $__res = (object)["bind" => function($v_7) use ($Bind1_5_23, $pure_6_24) {
  $__num = \func_num_args();
  $__res = function($k_8) use ($Bind1_5_23, $pure_6_24, $v_7) {
  $__num = \func_num_args();
  $__local_var_9_25 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_6_24))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_5_23)->{'bind'})($v_7))(function($v2_10) use ($__local_var_9_25, $k_8) {
  $__num = \func_num_args();
  $__t26 = null;;
  if ($v2_10 instanceof \Data\Either\Data_Either_Left) {
$__t26 = ($__local_var_9_25)(($v2_10)->{'value0'});
goto end_branch_26;;
};
  if ($v2_10 instanceof \Data\Either\Data_Either_Right) {
$__t26 = ($k_8)(($v2_10)->{'value0'});
goto end_branch_26;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t26 = null;
  end_branch_26:;
  $__res = $__t26;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_7) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__local_var_8_27 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_8_27 = (object)["map" => function($f_9) use ($__local_var_8_27) {
  $__num = \func_num_args();
  $__local_var_10_28 = (($__local_var_8_27)->{'map'})(function($m_10) use ($f_9) {
  $__num = \func_num_args();
  $__t28 = null;;
  if ($m_10 instanceof \Data\Either\Data_Either_Left) {
$__t28 = new \Data\Either\Data_Either_Left(($m_10)->{'value0'});
goto end_branch_28;;
};
  if ($m_10 instanceof \Data\Either\Data_Either_Right) {
$__t28 = new \Data\Either\Data_Either_Right(($f_9)(($m_10)->{'value0'}));
goto end_branch_28;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t28 = null;
  end_branch_28:;
  $__res = $__t28;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_11) use ($__local_var_10_28) {
  $__num = \func_num_args();
  $__res = ($__local_var_10_28)($v_11);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_9_31 = (($Monad0_1_0)->{'Bind1'})(null);
  $pure_10_32 = ((($Monad0_1_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_9_31 = (object)["bind" => function($v_11) use ($Bind1_9_31, $pure_10_32) {
  $__num = \func_num_args();
  $__res = function($k_12) use ($Bind1_9_31, $pure_10_32, $v_11) {
  $__num = \func_num_args();
  $__local_var_13_33 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_10_32))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_9_31)->{'bind'})($v_11))(function($v2_14) use ($__local_var_13_33, $k_12) {
  $__num = \func_num_args();
  $__t34 = null;;
  if ($v2_14 instanceof \Data\Either\Data_Either_Left) {
$__t34 = ($__local_var_13_33)(($v2_14)->{'value0'});
goto end_branch_34;;
};
  if ($v2_14 instanceof \Data\Either\Data_Either_Right) {
$__t34 = ($k_12)(($v2_14)->{'value0'});
goto end_branch_34;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t34 = null;
  end_branch_34:;
  $__res = $__t34;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_11) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_10_36 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_10) {
  $__num = \func_num_args();
  $__res = $x_10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($Monad0_1_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_10) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__local_var_11_36 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_11_36 = (object)["map" => function($f_12) use ($__local_var_11_36) {
  $__num = \func_num_args();
  $__local_var_13_37 = (($__local_var_11_36)->{'map'})(function($m_13) use ($f_12) {
  $__num = \func_num_args();
  $__t37 = null;;
  if ($m_13 instanceof \Data\Either\Data_Either_Left) {
$__t37 = new \Data\Either\Data_Either_Left(($m_13)->{'value0'});
goto end_branch_37;;
};
  if ($m_13 instanceof \Data\Either\Data_Either_Right) {
$__t37 = new \Data\Either\Data_Either_Right(($f_12)(($m_13)->{'value0'}));
goto end_branch_37;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t37 = null;
  end_branch_37:;
  $__res = $__t37;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_14) use ($__local_var_13_37) {
  $__num = \func_num_args();
  $__res = ($__local_var_13_37)($v_14);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_12_40 = (($Monad0_1_0)->{'Bind1'})(null);
  $pure_13_41 = ((($Monad0_1_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_12_40 = (object)["bind" => function($v_14) use ($Bind1_12_40, $pure_13_41) {
  $__num = \func_num_args();
  $__res = function($k_15) use ($Bind1_12_40, $pure_13_41, $v_14) {
  $__num = \func_num_args();
  $__local_var_16_42 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_13_41))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_12_40)->{'bind'})($v_14))(function($v2_17) use ($__local_var_16_42, $k_15) {
  $__num = \func_num_args();
  $__t43 = null;;
  if ($v2_17 instanceof \Data\Either\Data_Either_Left) {
$__t43 = ($__local_var_16_42)(($v2_17)->{'value0'});
goto end_branch_43;;
};
  if ($v2_17 instanceof \Data\Either\Data_Either_Right) {
$__t43 = ($k_15)(($v2_17)->{'value0'});
goto end_branch_43;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t43 = null;
  end_branch_43:;
  $__res = $__t43;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_14) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_13_45 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($Monad0_1_0);
  $__res = (object)["apply" => function($f_14) use ($Applicative0_13_45, $Bind1_12_40) {
  $__num = \func_num_args();
  $__res = function($a_15) use ($Applicative0_13_45, $Bind1_12_40, $f_14) {
  $__num = \func_num_args();
  $__res = ((($Bind1_12_40)->{'bind'})($f_14))(function($f_prime__16) use ($Applicative0_13_45, $Bind1_12_40, $a_15) {
  $__num = \func_num_args();
  $__res = ((($Bind1_12_40)->{'bind'})($a_15))(function($a_prime__17) use ($Applicative0_13_45, $f_prime__16) {
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
}, "Functor0" => function($_dollar___unused_12) use ($functorExceptT1_11_36) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_11_36;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_11) use ($Applicative0_10_36, $Bind1_9_31) {
  $__num = \func_num_args();
  $__res = function($a_12) use ($Applicative0_10_36, $Bind1_9_31, $f_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_31)->{'bind'})($f_11))(function($f_prime__13) use ($Applicative0_10_36, $Bind1_9_31, $a_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_31)->{'bind'})($a_12))(function($a_prime__14) use ($Applicative0_10_36, $f_prime__13) {
  $__num = \func_num_args();
  $__res = (($Applicative0_10_36)->{'pure'})(($f_prime__13)($a_prime__14));
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
}, "Functor0" => function($_dollar___unused_9) use ($functorExceptT1_8_27) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_8_27;
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
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_6) {
  $__num = \func_num_args();
  $__res = $x_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($dictMonadRec_0)->{'tailRecM'})(function($a_6) use ($Applicative0_3_2, $Bind1_2_1, $f_5) {
  $__num = \func_num_args();
  $__res = ((($Bind1_2_1)->{'bind'})(($f_5)($a_6)))(function($m_prime__7) use ($Applicative0_3_2) {
  $__num = \func_num_args();
  $__t48 = null;;
  if ($m_prime__7 instanceof \Data\Either\Data_Either_Left) {
$__t48 = new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Done(new \Data\Either\Data_Either_Left(($m_prime__7)->{'value0'}));
goto end_branch_48;;
};
  if ($m_prime__7 instanceof \Data\Either\Data_Either_Right) {
$__t49 = null;;
if (($m_prime__7)->{'value0'} instanceof \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Loop) {
$__t49 = new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Loop((($m_prime__7)->{'value0'})->{'value0'});
goto end_branch_49;;
};
if (($m_prime__7)->{'value0'} instanceof \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Done) {
$__t49 = new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Done(new \Data\Either\Data_Either_Right((($m_prime__7)->{'value0'})->{'value0'}));
goto end_branch_49;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t49 = null;
end_branch_49:;
$__t48 = $__t49;
goto end_branch_48;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t48 = null;
  end_branch_48:;
  $__res = (($Applicative0_3_2)->{'pure'})($__t48);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad0" => function($_dollar___unused_5) use ($monadExceptT1_4_3) {
  $__num = \func_num_args();
  $__res = $monadExceptT1_4_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Except_Trans_monadRecExceptT'] = __NAMESPACE__ . '\\majControl_majMonad_majExcept_majTrans_monadmajRecmajExceptmajT';

// Control_Monad_Except_Trans_monadStateExceptT
function majControl_majMonad_majExcept_majTrans_monadmajStatemajExceptmajT($dictMonadState_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majExcept_majTrans_monadmajStatemajExceptmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Monad0_1_0 = (($dictMonadState_0)->{'Monad0'})(null);
  $__local_var_2_1 = (($dictMonadState_0)->{'Monad0'})(null);
  $monadExceptT1_2_1 = (object)["Applicative0" => function($_dollar___unused_3) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_4) {
  $__num = \func_num_args();
  $__res = $x_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_4) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_5_2 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_5_2 = (object)["map" => function($f_6) use ($__local_var_5_2) {
  $__num = \func_num_args();
  $__local_var_7_3 = (($__local_var_5_2)->{'map'})(function($m_7) use ($f_6) {
  $__num = \func_num_args();
  $__t3 = null;;
  if ($m_7 instanceof \Data\Either\Data_Either_Left) {
$__t3 = new \Data\Either\Data_Either_Left(($m_7)->{'value0'});
goto end_branch_3;;
};
  if ($m_7 instanceof \Data\Either\Data_Either_Right) {
$__t3 = new \Data\Either\Data_Either_Right(($f_6)(($m_7)->{'value0'}));
goto end_branch_3;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t3 = null;
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_8) use ($__local_var_7_3) {
  $__num = \func_num_args();
  $__res = ($__local_var_7_3)($v_8);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_6_6 = (($__local_var_2_1)->{'Bind1'})(null);
  $pure_7_7 = ((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'};
  $Bind1_6_6 = (object)["bind" => function($v_8) use ($Bind1_6_6, $pure_7_7) {
  $__num = \func_num_args();
  $__res = function($k_9) use ($Bind1_6_6, $pure_7_7, $v_8) {
  $__num = \func_num_args();
  $__local_var_10_8 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_7_7))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_6_6)->{'bind'})($v_8))(function($v2_11) use ($__local_var_10_8, $k_9) {
  $__num = \func_num_args();
  $__t9 = null;;
  if ($v2_11 instanceof \Data\Either\Data_Either_Left) {
$__t9 = ($__local_var_10_8)(($v2_11)->{'value0'});
goto end_branch_9;;
};
  if ($v2_11 instanceof \Data\Either\Data_Either_Right) {
$__t9 = ($k_9)(($v2_11)->{'value0'});
goto end_branch_9;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t9 = null;
  end_branch_9:;
  $__res = $__t9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_8) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_9_10 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_9_10 = (object)["map" => function($f_10) use ($__local_var_9_10) {
  $__num = \func_num_args();
  $__local_var_11_11 = (($__local_var_9_10)->{'map'})(function($m_11) use ($f_10) {
  $__num = \func_num_args();
  $__t11 = null;;
  if ($m_11 instanceof \Data\Either\Data_Either_Left) {
$__t11 = new \Data\Either\Data_Either_Left(($m_11)->{'value0'});
goto end_branch_11;;
};
  if ($m_11 instanceof \Data\Either\Data_Either_Right) {
$__t11 = new \Data\Either\Data_Either_Right(($f_10)(($m_11)->{'value0'}));
goto end_branch_11;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t11 = null;
  end_branch_11:;
  $__res = $__t11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_12) use ($__local_var_11_11) {
  $__num = \func_num_args();
  $__res = ($__local_var_11_11)($v_12);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_10_14 = (($__local_var_2_1)->{'Bind1'})(null);
  $pure_11_15 = ((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'};
  $Bind1_10_14 = (object)["bind" => function($v_12) use ($Bind1_10_14, $pure_11_15) {
  $__num = \func_num_args();
  $__res = function($k_13) use ($Bind1_10_14, $pure_11_15, $v_12) {
  $__num = \func_num_args();
  $__local_var_14_16 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_11_15))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_10_14)->{'bind'})($v_12))(function($v2_15) use ($__local_var_14_16, $k_13) {
  $__num = \func_num_args();
  $__t17 = null;;
  if ($v2_15 instanceof \Data\Either\Data_Either_Left) {
$__t17 = ($__local_var_14_16)(($v2_15)->{'value0'});
goto end_branch_17;;
};
  if ($v2_15 instanceof \Data\Either\Data_Either_Right) {
$__t17 = ($k_13)(($v2_15)->{'value0'});
goto end_branch_17;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t17 = null;
  end_branch_17:;
  $__res = $__t17;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_12) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($__local_var_2_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_11_19 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($__local_var_2_1);
  $__res = (object)["apply" => function($f_12) use ($Applicative0_11_19, $Bind1_10_14) {
  $__num = \func_num_args();
  $__res = function($a_13) use ($Applicative0_11_19, $Bind1_10_14, $f_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_14)->{'bind'})($f_12))(function($f_prime__14) use ($Applicative0_11_19, $Bind1_10_14, $a_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_14)->{'bind'})($a_13))(function($a_prime__15) use ($Applicative0_11_19, $f_prime__14) {
  $__num = \func_num_args();
  $__res = (($Applicative0_11_19)->{'pure'})(($f_prime__14)($a_prime__15));
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
}, "Functor0" => function($_dollar___unused_10) use ($functorExceptT1_9_10) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_9_10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_7_21 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($__local_var_2_1);
  $__res = (object)["apply" => function($f_8) use ($Applicative0_7_21, $Bind1_6_6) {
  $__num = \func_num_args();
  $__res = function($a_9) use ($Applicative0_7_21, $Bind1_6_6, $f_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_6)->{'bind'})($f_8))(function($f_prime__10) use ($Applicative0_7_21, $Bind1_6_6, $a_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_6)->{'bind'})($a_9))(function($a_prime__11) use ($Applicative0_7_21, $f_prime__10) {
  $__num = \func_num_args();
  $__res = (($Applicative0_7_21)->{'pure'})(($f_prime__10)($a_prime__11));
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
}, "Functor0" => function($_dollar___unused_6) use ($functorExceptT1_5_2) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_5_2;
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
  $Bind1_4_22 = (($__local_var_2_1)->{'Bind1'})(null);
  $pure_5_23 = ((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'};
  $__res = (object)["bind" => function($v_6) use ($Bind1_4_22, $pure_5_23) {
  $__num = \func_num_args();
  $__res = function($k_7) use ($Bind1_4_22, $pure_5_23, $v_6) {
  $__num = \func_num_args();
  $__local_var_8_24 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_5_23))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_4_22)->{'bind'})($v_6))(function($v2_9) use ($__local_var_8_24, $k_7) {
  $__num = \func_num_args();
  $__t25 = null;;
  if ($v2_9 instanceof \Data\Either\Data_Either_Left) {
$__t25 = ($__local_var_8_24)(($v2_9)->{'value0'});
goto end_branch_25;;
};
  if ($v2_9 instanceof \Data\Either\Data_Either_Right) {
$__t25 = ($k_7)(($v2_9)->{'value0'});
goto end_branch_25;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t25 = null;
  end_branch_25:;
  $__res = $__t25;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_6) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_7_26 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_7_26 = (object)["map" => function($f_8) use ($__local_var_7_26) {
  $__num = \func_num_args();
  $__local_var_9_27 = (($__local_var_7_26)->{'map'})(function($m_9) use ($f_8) {
  $__num = \func_num_args();
  $__t27 = null;;
  if ($m_9 instanceof \Data\Either\Data_Either_Left) {
$__t27 = new \Data\Either\Data_Either_Left(($m_9)->{'value0'});
goto end_branch_27;;
};
  if ($m_9 instanceof \Data\Either\Data_Either_Right) {
$__t27 = new \Data\Either\Data_Either_Right(($f_8)(($m_9)->{'value0'}));
goto end_branch_27;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t27 = null;
  end_branch_27:;
  $__res = $__t27;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_10) use ($__local_var_9_27) {
  $__num = \func_num_args();
  $__res = ($__local_var_9_27)($v_10);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_8_30 = (($__local_var_2_1)->{'Bind1'})(null);
  $pure_9_31 = ((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'};
  $Bind1_8_30 = (object)["bind" => function($v_10) use ($Bind1_8_30, $pure_9_31) {
  $__num = \func_num_args();
  $__res = function($k_11) use ($Bind1_8_30, $pure_9_31, $v_10) {
  $__num = \func_num_args();
  $__local_var_12_32 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_9_31))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_8_30)->{'bind'})($v_10))(function($v2_13) use ($__local_var_12_32, $k_11) {
  $__num = \func_num_args();
  $__t33 = null;;
  if ($v2_13 instanceof \Data\Either\Data_Either_Left) {
$__t33 = ($__local_var_12_32)(($v2_13)->{'value0'});
goto end_branch_33;;
};
  if ($v2_13 instanceof \Data\Either\Data_Either_Right) {
$__t33 = ($k_11)(($v2_13)->{'value0'});
goto end_branch_33;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t33 = null;
  end_branch_33:;
  $__res = $__t33;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_10) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($__local_var_2_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_9_35 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_9) {
  $__num = \func_num_args();
  $__res = $x_9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_9) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_10_35 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_10_35 = (object)["map" => function($f_11) use ($__local_var_10_35) {
  $__num = \func_num_args();
  $__local_var_12_36 = (($__local_var_10_35)->{'map'})(function($m_12) use ($f_11) {
  $__num = \func_num_args();
  $__t36 = null;;
  if ($m_12 instanceof \Data\Either\Data_Either_Left) {
$__t36 = new \Data\Either\Data_Either_Left(($m_12)->{'value0'});
goto end_branch_36;;
};
  if ($m_12 instanceof \Data\Either\Data_Either_Right) {
$__t36 = new \Data\Either\Data_Either_Right(($f_11)(($m_12)->{'value0'}));
goto end_branch_36;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t36 = null;
  end_branch_36:;
  $__res = $__t36;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_13) use ($__local_var_12_36) {
  $__num = \func_num_args();
  $__res = ($__local_var_12_36)($v_13);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_11_39 = (($__local_var_2_1)->{'Bind1'})(null);
  $pure_12_40 = ((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'};
  $Bind1_11_39 = (object)["bind" => function($v_13) use ($Bind1_11_39, $pure_12_40) {
  $__num = \func_num_args();
  $__res = function($k_14) use ($Bind1_11_39, $pure_12_40, $v_13) {
  $__num = \func_num_args();
  $__local_var_15_41 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_12_40))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_11_39)->{'bind'})($v_13))(function($v2_16) use ($__local_var_15_41, $k_14) {
  $__num = \func_num_args();
  $__t42 = null;;
  if ($v2_16 instanceof \Data\Either\Data_Either_Left) {
$__t42 = ($__local_var_15_41)(($v2_16)->{'value0'});
goto end_branch_42;;
};
  if ($v2_16 instanceof \Data\Either\Data_Either_Right) {
$__t42 = ($k_14)(($v2_16)->{'value0'});
goto end_branch_42;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t42 = null;
  end_branch_42:;
  $__res = $__t42;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_13) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($__local_var_2_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_12_44 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($__local_var_2_1);
  $__res = (object)["apply" => function($f_13) use ($Applicative0_12_44, $Bind1_11_39) {
  $__num = \func_num_args();
  $__res = function($a_14) use ($Applicative0_12_44, $Bind1_11_39, $f_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_39)->{'bind'})($f_13))(function($f_prime__15) use ($Applicative0_12_44, $Bind1_11_39, $a_14) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_39)->{'bind'})($a_14))(function($a_prime__16) use ($Applicative0_12_44, $f_prime__15) {
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
}, "Functor0" => function($_dollar___unused_11) use ($functorExceptT1_10_35) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_10_35;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_10) use ($Applicative0_9_35, $Bind1_8_30) {
  $__num = \func_num_args();
  $__res = function($a_11) use ($Applicative0_9_35, $Bind1_8_30, $f_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_8_30)->{'bind'})($f_10))(function($f_prime__12) use ($Applicative0_9_35, $Bind1_8_30, $a_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_8_30)->{'bind'})($a_11))(function($a_prime__13) use ($Applicative0_9_35, $f_prime__12) {
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
}, "Functor0" => function($_dollar___unused_8) use ($functorExceptT1_7_26) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_7_26;
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
  $__res = (object)["state" => function($f_3) use ($Monad0_1_0, $dictMonadState_0) {
  $__num = \func_num_args();
  $pure_4_47 = ((($Monad0_1_0)->{'Applicative0'})(null))->{'pure'};
  $__res = ((((($Monad0_1_0)->{'Bind1'})(null))->{'bind'})((($dictMonadState_0)->{'state'})($f_3)))(function($a_5) use ($pure_4_47) {
  $__num = \func_num_args();
  $__res = ($pure_4_47)(new \Data\Either\Data_Either_Right($a_5));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad0" => function($_dollar___unused_3) use ($monadExceptT1_2_1) {
  $__num = \func_num_args();
  $__res = $monadExceptT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Except_Trans_monadStateExceptT'] = __NAMESPACE__ . '\\majControl_majMonad_majExcept_majTrans_monadmajStatemajExceptmajT';

// Control_Monad_Except_Trans_monadTellExceptT
function majControl_majMonad_majExcept_majTrans_monadmajTellmajExceptmajT($dictMonadTell_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majExcept_majTrans_monadmajTellmajExceptmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Monad1_1_0 = (($dictMonadTell_0)->{'Monad1'})(null);
  $Semigroup0_2_1 = (($dictMonadTell_0)->{'Semigroup0'})(null);
  $monadExceptT1_3_2 = (object)["Applicative0" => function($_dollar___unused_3) use ($Monad1_1_0) {
  $__num = \func_num_args();
  $__res = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_4) {
  $__num = \func_num_args();
  $__res = $x_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($Monad1_1_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_4) use ($Monad1_1_0) {
  $__num = \func_num_args();
  $__local_var_5_2 = (((((($Monad1_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_5_2 = (object)["map" => function($f_6) use ($__local_var_5_2) {
  $__num = \func_num_args();
  $__local_var_7_3 = (($__local_var_5_2)->{'map'})(function($m_7) use ($f_6) {
  $__num = \func_num_args();
  $__t3 = null;;
  if ($m_7 instanceof \Data\Either\Data_Either_Left) {
$__t3 = new \Data\Either\Data_Either_Left(($m_7)->{'value0'});
goto end_branch_3;;
};
  if ($m_7 instanceof \Data\Either\Data_Either_Right) {
$__t3 = new \Data\Either\Data_Either_Right(($f_6)(($m_7)->{'value0'}));
goto end_branch_3;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t3 = null;
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_8) use ($__local_var_7_3) {
  $__num = \func_num_args();
  $__res = ($__local_var_7_3)($v_8);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_6_6 = (($Monad1_1_0)->{'Bind1'})(null);
  $pure_7_7 = ((($Monad1_1_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_6_6 = (object)["bind" => function($v_8) use ($Bind1_6_6, $pure_7_7) {
  $__num = \func_num_args();
  $__res = function($k_9) use ($Bind1_6_6, $pure_7_7, $v_8) {
  $__num = \func_num_args();
  $__local_var_10_8 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_7_7))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_6_6)->{'bind'})($v_8))(function($v2_11) use ($__local_var_10_8, $k_9) {
  $__num = \func_num_args();
  $__t9 = null;;
  if ($v2_11 instanceof \Data\Either\Data_Either_Left) {
$__t9 = ($__local_var_10_8)(($v2_11)->{'value0'});
goto end_branch_9;;
};
  if ($v2_11 instanceof \Data\Either\Data_Either_Right) {
$__t9 = ($k_9)(($v2_11)->{'value0'});
goto end_branch_9;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t9 = null;
  end_branch_9:;
  $__res = $__t9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_8) use ($Monad1_1_0) {
  $__num = \func_num_args();
  $__local_var_9_10 = (((((($Monad1_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_9_10 = (object)["map" => function($f_10) use ($__local_var_9_10) {
  $__num = \func_num_args();
  $__local_var_11_11 = (($__local_var_9_10)->{'map'})(function($m_11) use ($f_10) {
  $__num = \func_num_args();
  $__t11 = null;;
  if ($m_11 instanceof \Data\Either\Data_Either_Left) {
$__t11 = new \Data\Either\Data_Either_Left(($m_11)->{'value0'});
goto end_branch_11;;
};
  if ($m_11 instanceof \Data\Either\Data_Either_Right) {
$__t11 = new \Data\Either\Data_Either_Right(($f_10)(($m_11)->{'value0'}));
goto end_branch_11;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t11 = null;
  end_branch_11:;
  $__res = $__t11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_12) use ($__local_var_11_11) {
  $__num = \func_num_args();
  $__res = ($__local_var_11_11)($v_12);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_10_14 = (($Monad1_1_0)->{'Bind1'})(null);
  $pure_11_15 = ((($Monad1_1_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_10_14 = (object)["bind" => function($v_12) use ($Bind1_10_14, $pure_11_15) {
  $__num = \func_num_args();
  $__res = function($k_13) use ($Bind1_10_14, $pure_11_15, $v_12) {
  $__num = \func_num_args();
  $__local_var_14_16 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_11_15))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_10_14)->{'bind'})($v_12))(function($v2_15) use ($__local_var_14_16, $k_13) {
  $__num = \func_num_args();
  $__t17 = null;;
  if ($v2_15 instanceof \Data\Either\Data_Either_Left) {
$__t17 = ($__local_var_14_16)(($v2_15)->{'value0'});
goto end_branch_17;;
};
  if ($v2_15 instanceof \Data\Either\Data_Either_Right) {
$__t17 = ($k_13)(($v2_15)->{'value0'});
goto end_branch_17;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t17 = null;
  end_branch_17:;
  $__res = $__t17;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_12) use ($Monad1_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($Monad1_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_11_19 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($Monad1_1_0);
  $__res = (object)["apply" => function($f_12) use ($Applicative0_11_19, $Bind1_10_14) {
  $__num = \func_num_args();
  $__res = function($a_13) use ($Applicative0_11_19, $Bind1_10_14, $f_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_14)->{'bind'})($f_12))(function($f_prime__14) use ($Applicative0_11_19, $Bind1_10_14, $a_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_14)->{'bind'})($a_13))(function($a_prime__15) use ($Applicative0_11_19, $f_prime__14) {
  $__num = \func_num_args();
  $__res = (($Applicative0_11_19)->{'pure'})(($f_prime__14)($a_prime__15));
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
}, "Functor0" => function($_dollar___unused_10) use ($functorExceptT1_9_10) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_9_10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_7_21 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($Monad1_1_0);
  $__res = (object)["apply" => function($f_8) use ($Applicative0_7_21, $Bind1_6_6) {
  $__num = \func_num_args();
  $__res = function($a_9) use ($Applicative0_7_21, $Bind1_6_6, $f_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_6)->{'bind'})($f_8))(function($f_prime__10) use ($Applicative0_7_21, $Bind1_6_6, $a_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_6)->{'bind'})($a_9))(function($a_prime__11) use ($Applicative0_7_21, $f_prime__10) {
  $__num = \func_num_args();
  $__res = (($Applicative0_7_21)->{'pure'})(($f_prime__10)($a_prime__11));
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
}, "Functor0" => function($_dollar___unused_6) use ($functorExceptT1_5_2) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_5_2;
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
  $Bind1_4_22 = (($Monad1_1_0)->{'Bind1'})(null);
  $pure_5_23 = ((($Monad1_1_0)->{'Applicative0'})(null))->{'pure'};
  $__res = (object)["bind" => function($v_6) use ($Bind1_4_22, $pure_5_23) {
  $__num = \func_num_args();
  $__res = function($k_7) use ($Bind1_4_22, $pure_5_23, $v_6) {
  $__num = \func_num_args();
  $__local_var_8_24 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_5_23))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_4_22)->{'bind'})($v_6))(function($v2_9) use ($__local_var_8_24, $k_7) {
  $__num = \func_num_args();
  $__t25 = null;;
  if ($v2_9 instanceof \Data\Either\Data_Either_Left) {
$__t25 = ($__local_var_8_24)(($v2_9)->{'value0'});
goto end_branch_25;;
};
  if ($v2_9 instanceof \Data\Either\Data_Either_Right) {
$__t25 = ($k_7)(($v2_9)->{'value0'});
goto end_branch_25;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t25 = null;
  end_branch_25:;
  $__res = $__t25;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_6) use ($Monad1_1_0) {
  $__num = \func_num_args();
  $__local_var_7_26 = (((((($Monad1_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_7_26 = (object)["map" => function($f_8) use ($__local_var_7_26) {
  $__num = \func_num_args();
  $__local_var_9_27 = (($__local_var_7_26)->{'map'})(function($m_9) use ($f_8) {
  $__num = \func_num_args();
  $__t27 = null;;
  if ($m_9 instanceof \Data\Either\Data_Either_Left) {
$__t27 = new \Data\Either\Data_Either_Left(($m_9)->{'value0'});
goto end_branch_27;;
};
  if ($m_9 instanceof \Data\Either\Data_Either_Right) {
$__t27 = new \Data\Either\Data_Either_Right(($f_8)(($m_9)->{'value0'}));
goto end_branch_27;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t27 = null;
  end_branch_27:;
  $__res = $__t27;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_10) use ($__local_var_9_27) {
  $__num = \func_num_args();
  $__res = ($__local_var_9_27)($v_10);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_8_30 = (($Monad1_1_0)->{'Bind1'})(null);
  $pure_9_31 = ((($Monad1_1_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_8_30 = (object)["bind" => function($v_10) use ($Bind1_8_30, $pure_9_31) {
  $__num = \func_num_args();
  $__res = function($k_11) use ($Bind1_8_30, $pure_9_31, $v_10) {
  $__num = \func_num_args();
  $__local_var_12_32 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_9_31))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_8_30)->{'bind'})($v_10))(function($v2_13) use ($__local_var_12_32, $k_11) {
  $__num = \func_num_args();
  $__t33 = null;;
  if ($v2_13 instanceof \Data\Either\Data_Either_Left) {
$__t33 = ($__local_var_12_32)(($v2_13)->{'value0'});
goto end_branch_33;;
};
  if ($v2_13 instanceof \Data\Either\Data_Either_Right) {
$__t33 = ($k_11)(($v2_13)->{'value0'});
goto end_branch_33;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t33 = null;
  end_branch_33:;
  $__res = $__t33;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_10) use ($Monad1_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($Monad1_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_9_35 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_9) {
  $__num = \func_num_args();
  $__res = $x_9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($Monad1_1_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_9) use ($Monad1_1_0) {
  $__num = \func_num_args();
  $__local_var_10_35 = (((((($Monad1_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_10_35 = (object)["map" => function($f_11) use ($__local_var_10_35) {
  $__num = \func_num_args();
  $__local_var_12_36 = (($__local_var_10_35)->{'map'})(function($m_12) use ($f_11) {
  $__num = \func_num_args();
  $__t36 = null;;
  if ($m_12 instanceof \Data\Either\Data_Either_Left) {
$__t36 = new \Data\Either\Data_Either_Left(($m_12)->{'value0'});
goto end_branch_36;;
};
  if ($m_12 instanceof \Data\Either\Data_Either_Right) {
$__t36 = new \Data\Either\Data_Either_Right(($f_11)(($m_12)->{'value0'}));
goto end_branch_36;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t36 = null;
  end_branch_36:;
  $__res = $__t36;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_13) use ($__local_var_12_36) {
  $__num = \func_num_args();
  $__res = ($__local_var_12_36)($v_13);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_11_39 = (($Monad1_1_0)->{'Bind1'})(null);
  $pure_12_40 = ((($Monad1_1_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_11_39 = (object)["bind" => function($v_13) use ($Bind1_11_39, $pure_12_40) {
  $__num = \func_num_args();
  $__res = function($k_14) use ($Bind1_11_39, $pure_12_40, $v_13) {
  $__num = \func_num_args();
  $__local_var_15_41 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_12_40))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_11_39)->{'bind'})($v_13))(function($v2_16) use ($__local_var_15_41, $k_14) {
  $__num = \func_num_args();
  $__t42 = null;;
  if ($v2_16 instanceof \Data\Either\Data_Either_Left) {
$__t42 = ($__local_var_15_41)(($v2_16)->{'value0'});
goto end_branch_42;;
};
  if ($v2_16 instanceof \Data\Either\Data_Either_Right) {
$__t42 = ($k_14)(($v2_16)->{'value0'});
goto end_branch_42;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t42 = null;
  end_branch_42:;
  $__res = $__t42;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_13) use ($Monad1_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($Monad1_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_12_44 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($Monad1_1_0);
  $__res = (object)["apply" => function($f_13) use ($Applicative0_12_44, $Bind1_11_39) {
  $__num = \func_num_args();
  $__res = function($a_14) use ($Applicative0_12_44, $Bind1_11_39, $f_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_39)->{'bind'})($f_13))(function($f_prime__15) use ($Applicative0_12_44, $Bind1_11_39, $a_14) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_39)->{'bind'})($a_14))(function($a_prime__16) use ($Applicative0_12_44, $f_prime__15) {
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
}, "Functor0" => function($_dollar___unused_11) use ($functorExceptT1_10_35) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_10_35;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_10) use ($Applicative0_9_35, $Bind1_8_30) {
  $__num = \func_num_args();
  $__res = function($a_11) use ($Applicative0_9_35, $Bind1_8_30, $f_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_8_30)->{'bind'})($f_10))(function($f_prime__12) use ($Applicative0_9_35, $Bind1_8_30, $a_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_8_30)->{'bind'})($a_11))(function($a_prime__13) use ($Applicative0_9_35, $f_prime__12) {
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
}, "Functor0" => function($_dollar___unused_8) use ($functorExceptT1_7_26) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_7_26;
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
  $Bind1_4_47 = (($Monad1_1_0)->{'Bind1'})(null);
  $pure_5_48 = ((($Monad1_1_0)->{'Applicative0'})(null))->{'pure'};
  $__res = (object)["tell" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($m_6) use ($Bind1_4_47, $pure_5_48) {
  $__num = \func_num_args();
  $__res = ((($Bind1_4_47)->{'bind'})($m_6))(function($a_7) use ($pure_5_48) {
  $__num = \func_num_args();
  $__res = ($pure_5_48)(new \Data\Either\Data_Either_Right($a_7));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($dictMonadTell_0)->{'tell'}), "Semigroup0" => function($_dollar___unused_4) use ($Semigroup0_2_1) {
  $__num = \func_num_args();
  $__res = $Semigroup0_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad1" => function($_dollar___unused_4) use ($monadExceptT1_3_2) {
  $__num = \func_num_args();
  $__res = $monadExceptT1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Except_Trans_monadTellExceptT'] = __NAMESPACE__ . '\\majControl_majMonad_majExcept_majTrans_monadmajTellmajExceptmajT';

// Control_Monad_Except_Trans_monadWriterExceptT
function majControl_majMonad_majExcept_majTrans_monadmajWritermajExceptmajT($dictMonadWriter_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majExcept_majTrans_monadmajWritermajExceptmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $MonadTell1_1_0 = (($dictMonadWriter_0)->{'MonadTell1'})(null);
  $Monad1_2_1 = (($MonadTell1_1_0)->{'Monad1'})(null);
  $Bind1_3_2 = (($Monad1_2_1)->{'Bind1'})(null);
  $pure_4_3 = ((($Monad1_2_1)->{'Applicative0'})(null))->{'pure'};
  $Applicative0_5_4 = (($Monad1_2_1)->{'Applicative0'})(null);
  $Monoid0_6_5 = (($dictMonadWriter_0)->{'Monoid0'})(null);
  $Monad1_7_6 = (($MonadTell1_1_0)->{'Monad1'})(null);
  $Semigroup0_8_7 = (($MonadTell1_1_0)->{'Semigroup0'})(null);
  $monadExceptT1_9_8 = (object)["Applicative0" => function($_dollar___unused_9) use ($Monad1_7_6) {
  $__num = \func_num_args();
  $__res = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_10) {
  $__num = \func_num_args();
  $__res = $x_10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($Monad1_7_6)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_10) use ($Monad1_7_6) {
  $__num = \func_num_args();
  $__local_var_11_8 = (((((($Monad1_7_6)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_11_8 = (object)["map" => function($f_12) use ($__local_var_11_8) {
  $__num = \func_num_args();
  $__local_var_13_9 = (($__local_var_11_8)->{'map'})(function($m_13) use ($f_12) {
  $__num = \func_num_args();
  $__t9 = null;;
  if ($m_13 instanceof \Data\Either\Data_Either_Left) {
$__t9 = new \Data\Either\Data_Either_Left(($m_13)->{'value0'});
goto end_branch_9;;
};
  if ($m_13 instanceof \Data\Either\Data_Either_Right) {
$__t9 = new \Data\Either\Data_Either_Right(($f_12)(($m_13)->{'value0'}));
goto end_branch_9;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t9 = null;
  end_branch_9:;
  $__res = $__t9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_14) use ($__local_var_13_9) {
  $__num = \func_num_args();
  $__res = ($__local_var_13_9)($v_14);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_12_12 = (($Monad1_7_6)->{'Bind1'})(null);
  $pure_13_13 = ((($Monad1_7_6)->{'Applicative0'})(null))->{'pure'};
  $Bind1_12_12 = (object)["bind" => function($v_14) use ($Bind1_12_12, $pure_13_13) {
  $__num = \func_num_args();
  $__res = function($k_15) use ($Bind1_12_12, $pure_13_13, $v_14) {
  $__num = \func_num_args();
  $__local_var_16_14 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_13_13))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_12_12)->{'bind'})($v_14))(function($v2_17) use ($__local_var_16_14, $k_15) {
  $__num = \func_num_args();
  $__t15 = null;;
  if ($v2_17 instanceof \Data\Either\Data_Either_Left) {
$__t15 = ($__local_var_16_14)(($v2_17)->{'value0'});
goto end_branch_15;;
};
  if ($v2_17 instanceof \Data\Either\Data_Either_Right) {
$__t15 = ($k_15)(($v2_17)->{'value0'});
goto end_branch_15;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t15 = null;
  end_branch_15:;
  $__res = $__t15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_14) use ($Monad1_7_6) {
  $__num = \func_num_args();
  $__local_var_15_16 = (((((($Monad1_7_6)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_15_16 = (object)["map" => function($f_16) use ($__local_var_15_16) {
  $__num = \func_num_args();
  $__local_var_17_17 = (($__local_var_15_16)->{'map'})(function($m_17) use ($f_16) {
  $__num = \func_num_args();
  $__t17 = null;;
  if ($m_17 instanceof \Data\Either\Data_Either_Left) {
$__t17 = new \Data\Either\Data_Either_Left(($m_17)->{'value0'});
goto end_branch_17;;
};
  if ($m_17 instanceof \Data\Either\Data_Either_Right) {
$__t17 = new \Data\Either\Data_Either_Right(($f_16)(($m_17)->{'value0'}));
goto end_branch_17;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t17 = null;
  end_branch_17:;
  $__res = $__t17;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_18) use ($__local_var_17_17) {
  $__num = \func_num_args();
  $__res = ($__local_var_17_17)($v_18);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_16_20 = (($Monad1_7_6)->{'Bind1'})(null);
  $pure_17_21 = ((($Monad1_7_6)->{'Applicative0'})(null))->{'pure'};
  $Bind1_16_20 = (object)["bind" => function($v_18) use ($Bind1_16_20, $pure_17_21) {
  $__num = \func_num_args();
  $__res = function($k_19) use ($Bind1_16_20, $pure_17_21, $v_18) {
  $__num = \func_num_args();
  $__local_var_20_22 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_17_21))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_16_20)->{'bind'})($v_18))(function($v2_21) use ($__local_var_20_22, $k_19) {
  $__num = \func_num_args();
  $__t23 = null;;
  if ($v2_21 instanceof \Data\Either\Data_Either_Left) {
$__t23 = ($__local_var_20_22)(($v2_21)->{'value0'});
goto end_branch_23;;
};
  if ($v2_21 instanceof \Data\Either\Data_Either_Right) {
$__t23 = ($k_19)(($v2_21)->{'value0'});
goto end_branch_23;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t23 = null;
  end_branch_23:;
  $__res = $__t23;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_18) use ($Monad1_7_6) {
  $__num = \func_num_args();
  $__local_var_19_24 = (((((($Monad1_7_6)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_19_24 = (object)["map" => function($f_20) use ($__local_var_19_24) {
  $__num = \func_num_args();
  $__local_var_21_25 = (($__local_var_19_24)->{'map'})(function($m_21) use ($f_20) {
  $__num = \func_num_args();
  $__t25 = null;;
  if ($m_21 instanceof \Data\Either\Data_Either_Left) {
$__t25 = new \Data\Either\Data_Either_Left(($m_21)->{'value0'});
goto end_branch_25;;
};
  if ($m_21 instanceof \Data\Either\Data_Either_Right) {
$__t25 = new \Data\Either\Data_Either_Right(($f_20)(($m_21)->{'value0'}));
goto end_branch_25;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t25 = null;
  end_branch_25:;
  $__res = $__t25;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_22) use ($__local_var_21_25) {
  $__num = \func_num_args();
  $__res = ($__local_var_21_25)($v_22);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_20_28 = (($Monad1_7_6)->{'Bind1'})(null);
  $pure_21_29 = ((($Monad1_7_6)->{'Applicative0'})(null))->{'pure'};
  $Bind1_20_28 = (object)["bind" => function($v_22) use ($Bind1_20_28, $pure_21_29) {
  $__num = \func_num_args();
  $__res = function($k_23) use ($Bind1_20_28, $pure_21_29, $v_22) {
  $__num = \func_num_args();
  $__local_var_24_30 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_21_29))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_20_28)->{'bind'})($v_22))(function($v2_25) use ($__local_var_24_30, $k_23) {
  $__num = \func_num_args();
  $__t31 = null;;
  if ($v2_25 instanceof \Data\Either\Data_Either_Left) {
$__t31 = ($__local_var_24_30)(($v2_25)->{'value0'});
goto end_branch_31;;
};
  if ($v2_25 instanceof \Data\Either\Data_Either_Right) {
$__t31 = ($k_23)(($v2_25)->{'value0'});
goto end_branch_31;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t31 = null;
  end_branch_31:;
  $__res = $__t31;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_22) use ($Monad1_7_6) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($Monad1_7_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_21_33 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_21) {
  $__num = \func_num_args();
  $__res = $x_21;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($Monad1_7_6)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_21) use ($Monad1_7_6) {
  $__num = \func_num_args();
  $__local_var_22_33 = (((((($Monad1_7_6)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_22_33 = (object)["map" => function($f_23) use ($__local_var_22_33) {
  $__num = \func_num_args();
  $__local_var_24_34 = (($__local_var_22_33)->{'map'})(function($m_24) use ($f_23) {
  $__num = \func_num_args();
  $__t34 = null;;
  if ($m_24 instanceof \Data\Either\Data_Either_Left) {
$__t34 = new \Data\Either\Data_Either_Left(($m_24)->{'value0'});
goto end_branch_34;;
};
  if ($m_24 instanceof \Data\Either\Data_Either_Right) {
$__t34 = new \Data\Either\Data_Either_Right(($f_23)(($m_24)->{'value0'}));
goto end_branch_34;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t34 = null;
  end_branch_34:;
  $__res = $__t34;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_25) use ($__local_var_24_34) {
  $__num = \func_num_args();
  $__res = ($__local_var_24_34)($v_25);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_23_37 = (($Monad1_7_6)->{'Bind1'})(null);
  $pure_24_38 = ((($Monad1_7_6)->{'Applicative0'})(null))->{'pure'};
  $Bind1_23_37 = (object)["bind" => function($v_25) use ($Bind1_23_37, $pure_24_38) {
  $__num = \func_num_args();
  $__res = function($k_26) use ($Bind1_23_37, $pure_24_38, $v_25) {
  $__num = \func_num_args();
  $__local_var_27_39 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_24_38))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_23_37)->{'bind'})($v_25))(function($v2_28) use ($__local_var_27_39, $k_26) {
  $__num = \func_num_args();
  $__t40 = null;;
  if ($v2_28 instanceof \Data\Either\Data_Either_Left) {
$__t40 = ($__local_var_27_39)(($v2_28)->{'value0'});
goto end_branch_40;;
};
  if ($v2_28 instanceof \Data\Either\Data_Either_Right) {
$__t40 = ($k_26)(($v2_28)->{'value0'});
goto end_branch_40;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t40 = null;
  end_branch_40:;
  $__res = $__t40;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_25) use ($Monad1_7_6) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($Monad1_7_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_24_42 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($Monad1_7_6);
  $__res = (object)["apply" => function($f_25) use ($Applicative0_24_42, $Bind1_23_37) {
  $__num = \func_num_args();
  $__res = function($a_26) use ($Applicative0_24_42, $Bind1_23_37, $f_25) {
  $__num = \func_num_args();
  $__res = ((($Bind1_23_37)->{'bind'})($f_25))(function($f_prime__27) use ($Applicative0_24_42, $Bind1_23_37, $a_26) {
  $__num = \func_num_args();
  $__res = ((($Bind1_23_37)->{'bind'})($a_26))(function($a_prime__28) use ($Applicative0_24_42, $f_prime__27) {
  $__num = \func_num_args();
  $__res = (($Applicative0_24_42)->{'pure'})(($f_prime__27)($a_prime__28));
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
}, "Functor0" => function($_dollar___unused_23) use ($functorExceptT1_22_33) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_22_33;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_22) use ($Applicative0_21_33, $Bind1_20_28) {
  $__num = \func_num_args();
  $__res = function($a_23) use ($Applicative0_21_33, $Bind1_20_28, $f_22) {
  $__num = \func_num_args();
  $__res = ((($Bind1_20_28)->{'bind'})($f_22))(function($f_prime__24) use ($Applicative0_21_33, $Bind1_20_28, $a_23) {
  $__num = \func_num_args();
  $__res = ((($Bind1_20_28)->{'bind'})($a_23))(function($a_prime__25) use ($Applicative0_21_33, $f_prime__24) {
  $__num = \func_num_args();
  $__res = (($Applicative0_21_33)->{'pure'})(($f_prime__24)($a_prime__25));
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
}, "Functor0" => function($_dollar___unused_20) use ($functorExceptT1_19_24) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_19_24;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_17_45 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_17) {
  $__num = \func_num_args();
  $__res = $x_17;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($Monad1_7_6)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_17) use ($Monad1_7_6) {
  $__num = \func_num_args();
  $__local_var_18_45 = (((((($Monad1_7_6)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_18_45 = (object)["map" => function($f_19) use ($__local_var_18_45) {
  $__num = \func_num_args();
  $__local_var_20_46 = (($__local_var_18_45)->{'map'})(function($m_20) use ($f_19) {
  $__num = \func_num_args();
  $__t46 = null;;
  if ($m_20 instanceof \Data\Either\Data_Either_Left) {
$__t46 = new \Data\Either\Data_Either_Left(($m_20)->{'value0'});
goto end_branch_46;;
};
  if ($m_20 instanceof \Data\Either\Data_Either_Right) {
$__t46 = new \Data\Either\Data_Either_Right(($f_19)(($m_20)->{'value0'}));
goto end_branch_46;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t46 = null;
  end_branch_46:;
  $__res = $__t46;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_21) use ($__local_var_20_46) {
  $__num = \func_num_args();
  $__res = ($__local_var_20_46)($v_21);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_19_49 = (($Monad1_7_6)->{'Bind1'})(null);
  $pure_20_50 = ((($Monad1_7_6)->{'Applicative0'})(null))->{'pure'};
  $Bind1_19_49 = (object)["bind" => function($v_21) use ($Bind1_19_49, $pure_20_50) {
  $__num = \func_num_args();
  $__res = function($k_22) use ($Bind1_19_49, $pure_20_50, $v_21) {
  $__num = \func_num_args();
  $__local_var_23_51 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_20_50))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_19_49)->{'bind'})($v_21))(function($v2_24) use ($__local_var_23_51, $k_22) {
  $__num = \func_num_args();
  $__t52 = null;;
  if ($v2_24 instanceof \Data\Either\Data_Either_Left) {
$__t52 = ($__local_var_23_51)(($v2_24)->{'value0'});
goto end_branch_52;;
};
  if ($v2_24 instanceof \Data\Either\Data_Either_Right) {
$__t52 = ($k_22)(($v2_24)->{'value0'});
goto end_branch_52;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t52 = null;
  end_branch_52:;
  $__res = $__t52;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_21) use ($Monad1_7_6) {
  $__num = \func_num_args();
  $__local_var_22_53 = (((((($Monad1_7_6)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_22_53 = (object)["map" => function($f_23) use ($__local_var_22_53) {
  $__num = \func_num_args();
  $__local_var_24_54 = (($__local_var_22_53)->{'map'})(function($m_24) use ($f_23) {
  $__num = \func_num_args();
  $__t54 = null;;
  if ($m_24 instanceof \Data\Either\Data_Either_Left) {
$__t54 = new \Data\Either\Data_Either_Left(($m_24)->{'value0'});
goto end_branch_54;;
};
  if ($m_24 instanceof \Data\Either\Data_Either_Right) {
$__t54 = new \Data\Either\Data_Either_Right(($f_23)(($m_24)->{'value0'}));
goto end_branch_54;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t54 = null;
  end_branch_54:;
  $__res = $__t54;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_25) use ($__local_var_24_54) {
  $__num = \func_num_args();
  $__res = ($__local_var_24_54)($v_25);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_23_57 = (($Monad1_7_6)->{'Bind1'})(null);
  $pure_24_58 = ((($Monad1_7_6)->{'Applicative0'})(null))->{'pure'};
  $Bind1_23_57 = (object)["bind" => function($v_25) use ($Bind1_23_57, $pure_24_58) {
  $__num = \func_num_args();
  $__res = function($k_26) use ($Bind1_23_57, $pure_24_58, $v_25) {
  $__num = \func_num_args();
  $__local_var_27_59 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_24_58))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_23_57)->{'bind'})($v_25))(function($v2_28) use ($__local_var_27_59, $k_26) {
  $__num = \func_num_args();
  $__t60 = null;;
  if ($v2_28 instanceof \Data\Either\Data_Either_Left) {
$__t60 = ($__local_var_27_59)(($v2_28)->{'value0'});
goto end_branch_60;;
};
  if ($v2_28 instanceof \Data\Either\Data_Either_Right) {
$__t60 = ($k_26)(($v2_28)->{'value0'});
goto end_branch_60;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t60 = null;
  end_branch_60:;
  $__res = $__t60;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_25) use ($Monad1_7_6) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($Monad1_7_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_24_62 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($Monad1_7_6);
  $__res = (object)["apply" => function($f_25) use ($Applicative0_24_62, $Bind1_23_57) {
  $__num = \func_num_args();
  $__res = function($a_26) use ($Applicative0_24_62, $Bind1_23_57, $f_25) {
  $__num = \func_num_args();
  $__res = ((($Bind1_23_57)->{'bind'})($f_25))(function($f_prime__27) use ($Applicative0_24_62, $Bind1_23_57, $a_26) {
  $__num = \func_num_args();
  $__res = ((($Bind1_23_57)->{'bind'})($a_26))(function($a_prime__28) use ($Applicative0_24_62, $f_prime__27) {
  $__num = \func_num_args();
  $__res = (($Applicative0_24_62)->{'pure'})(($f_prime__27)($a_prime__28));
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
}, "Functor0" => function($_dollar___unused_23) use ($functorExceptT1_22_53) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_22_53;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_20_64 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($Monad1_7_6);
  $__res = (object)["apply" => function($f_21) use ($Applicative0_20_64, $Bind1_19_49) {
  $__num = \func_num_args();
  $__res = function($a_22) use ($Applicative0_20_64, $Bind1_19_49, $f_21) {
  $__num = \func_num_args();
  $__res = ((($Bind1_19_49)->{'bind'})($f_21))(function($f_prime__23) use ($Applicative0_20_64, $Bind1_19_49, $a_22) {
  $__num = \func_num_args();
  $__res = ((($Bind1_19_49)->{'bind'})($a_22))(function($a_prime__24) use ($Applicative0_20_64, $f_prime__23) {
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
}, "Functor0" => function($_dollar___unused_19) use ($functorExceptT1_18_45) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_18_45;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_18) use ($Applicative0_17_45, $Bind1_16_20) {
  $__num = \func_num_args();
  $__res = function($a_19) use ($Applicative0_17_45, $Bind1_16_20, $f_18) {
  $__num = \func_num_args();
  $__res = ((($Bind1_16_20)->{'bind'})($f_18))(function($f_prime__20) use ($Applicative0_17_45, $Bind1_16_20, $a_19) {
  $__num = \func_num_args();
  $__res = ((($Bind1_16_20)->{'bind'})($a_19))(function($a_prime__21) use ($Applicative0_17_45, $f_prime__20) {
  $__num = \func_num_args();
  $__res = (($Applicative0_17_45)->{'pure'})(($f_prime__20)($a_prime__21));
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
}, "Functor0" => function($_dollar___unused_16) use ($functorExceptT1_15_16) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_15_16;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_13_67 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_13) {
  $__num = \func_num_args();
  $__res = $x_13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($Monad1_7_6)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_13) use ($Monad1_7_6) {
  $__num = \func_num_args();
  $__local_var_14_67 = (((((($Monad1_7_6)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_14_67 = (object)["map" => function($f_15) use ($__local_var_14_67) {
  $__num = \func_num_args();
  $__local_var_16_68 = (($__local_var_14_67)->{'map'})(function($m_16) use ($f_15) {
  $__num = \func_num_args();
  $__t68 = null;;
  if ($m_16 instanceof \Data\Either\Data_Either_Left) {
$__t68 = new \Data\Either\Data_Either_Left(($m_16)->{'value0'});
goto end_branch_68;;
};
  if ($m_16 instanceof \Data\Either\Data_Either_Right) {
$__t68 = new \Data\Either\Data_Either_Right(($f_15)(($m_16)->{'value0'}));
goto end_branch_68;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t68 = null;
  end_branch_68:;
  $__res = $__t68;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_17) use ($__local_var_16_68) {
  $__num = \func_num_args();
  $__res = ($__local_var_16_68)($v_17);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_15_71 = (($Monad1_7_6)->{'Bind1'})(null);
  $pure_16_72 = ((($Monad1_7_6)->{'Applicative0'})(null))->{'pure'};
  $Bind1_15_71 = (object)["bind" => function($v_17) use ($Bind1_15_71, $pure_16_72) {
  $__num = \func_num_args();
  $__res = function($k_18) use ($Bind1_15_71, $pure_16_72, $v_17) {
  $__num = \func_num_args();
  $__local_var_19_73 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_16_72))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_15_71)->{'bind'})($v_17))(function($v2_20) use ($__local_var_19_73, $k_18) {
  $__num = \func_num_args();
  $__t74 = null;;
  if ($v2_20 instanceof \Data\Either\Data_Either_Left) {
$__t74 = ($__local_var_19_73)(($v2_20)->{'value0'});
goto end_branch_74;;
};
  if ($v2_20 instanceof \Data\Either\Data_Either_Right) {
$__t74 = ($k_18)(($v2_20)->{'value0'});
goto end_branch_74;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t74 = null;
  end_branch_74:;
  $__res = $__t74;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_17) use ($Monad1_7_6) {
  $__num = \func_num_args();
  $__local_var_18_75 = (((((($Monad1_7_6)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_18_75 = (object)["map" => function($f_19) use ($__local_var_18_75) {
  $__num = \func_num_args();
  $__local_var_20_76 = (($__local_var_18_75)->{'map'})(function($m_20) use ($f_19) {
  $__num = \func_num_args();
  $__t76 = null;;
  if ($m_20 instanceof \Data\Either\Data_Either_Left) {
$__t76 = new \Data\Either\Data_Either_Left(($m_20)->{'value0'});
goto end_branch_76;;
};
  if ($m_20 instanceof \Data\Either\Data_Either_Right) {
$__t76 = new \Data\Either\Data_Either_Right(($f_19)(($m_20)->{'value0'}));
goto end_branch_76;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t76 = null;
  end_branch_76:;
  $__res = $__t76;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_21) use ($__local_var_20_76) {
  $__num = \func_num_args();
  $__res = ($__local_var_20_76)($v_21);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_19_79 = (($Monad1_7_6)->{'Bind1'})(null);
  $pure_20_80 = ((($Monad1_7_6)->{'Applicative0'})(null))->{'pure'};
  $Bind1_19_79 = (object)["bind" => function($v_21) use ($Bind1_19_79, $pure_20_80) {
  $__num = \func_num_args();
  $__res = function($k_22) use ($Bind1_19_79, $pure_20_80, $v_21) {
  $__num = \func_num_args();
  $__local_var_23_81 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_20_80))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_19_79)->{'bind'})($v_21))(function($v2_24) use ($__local_var_23_81, $k_22) {
  $__num = \func_num_args();
  $__t82 = null;;
  if ($v2_24 instanceof \Data\Either\Data_Either_Left) {
$__t82 = ($__local_var_23_81)(($v2_24)->{'value0'});
goto end_branch_82;;
};
  if ($v2_24 instanceof \Data\Either\Data_Either_Right) {
$__t82 = ($k_22)(($v2_24)->{'value0'});
goto end_branch_82;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t82 = null;
  end_branch_82:;
  $__res = $__t82;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_21) use ($Monad1_7_6) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($Monad1_7_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_20_84 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($Monad1_7_6);
  $__res = (object)["apply" => function($f_21) use ($Applicative0_20_84, $Bind1_19_79) {
  $__num = \func_num_args();
  $__res = function($a_22) use ($Applicative0_20_84, $Bind1_19_79, $f_21) {
  $__num = \func_num_args();
  $__res = ((($Bind1_19_79)->{'bind'})($f_21))(function($f_prime__23) use ($Applicative0_20_84, $Bind1_19_79, $a_22) {
  $__num = \func_num_args();
  $__res = ((($Bind1_19_79)->{'bind'})($a_22))(function($a_prime__24) use ($Applicative0_20_84, $f_prime__23) {
  $__num = \func_num_args();
  $__res = (($Applicative0_20_84)->{'pure'})(($f_prime__23)($a_prime__24));
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
}, "Functor0" => function($_dollar___unused_19) use ($functorExceptT1_18_75) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_18_75;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_16_86 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($Monad1_7_6);
  $__res = (object)["apply" => function($f_17) use ($Applicative0_16_86, $Bind1_15_71) {
  $__num = \func_num_args();
  $__res = function($a_18) use ($Applicative0_16_86, $Bind1_15_71, $f_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_71)->{'bind'})($f_17))(function($f_prime__19) use ($Applicative0_16_86, $Bind1_15_71, $a_18) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_71)->{'bind'})($a_18))(function($a_prime__20) use ($Applicative0_16_86, $f_prime__19) {
  $__num = \func_num_args();
  $__res = (($Applicative0_16_86)->{'pure'})(($f_prime__19)($a_prime__20));
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
}, "Functor0" => function($_dollar___unused_15) use ($functorExceptT1_14_67) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_14_67;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_14) use ($Applicative0_13_67, $Bind1_12_12) {
  $__num = \func_num_args();
  $__res = function($a_15) use ($Applicative0_13_67, $Bind1_12_12, $f_14) {
  $__num = \func_num_args();
  $__res = ((($Bind1_12_12)->{'bind'})($f_14))(function($f_prime__16) use ($Applicative0_13_67, $Bind1_12_12, $a_15) {
  $__num = \func_num_args();
  $__res = ((($Bind1_12_12)->{'bind'})($a_15))(function($a_prime__17) use ($Applicative0_13_67, $f_prime__16) {
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
}, "Functor0" => function($_dollar___unused_12) use ($functorExceptT1_11_8) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_11_8;
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
}, "Bind1" => function($_dollar___unused_9) use ($Monad1_7_6) {
  $__num = \func_num_args();
  $Bind1_10_88 = (($Monad1_7_6)->{'Bind1'})(null);
  $pure_11_89 = ((($Monad1_7_6)->{'Applicative0'})(null))->{'pure'};
  $__res = (object)["bind" => function($v_12) use ($Bind1_10_88, $pure_11_89) {
  $__num = \func_num_args();
  $__res = function($k_13) use ($Bind1_10_88, $pure_11_89, $v_12) {
  $__num = \func_num_args();
  $__local_var_14_90 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_11_89))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_10_88)->{'bind'})($v_12))(function($v2_15) use ($__local_var_14_90, $k_13) {
  $__num = \func_num_args();
  $__t91 = null;;
  if ($v2_15 instanceof \Data\Either\Data_Either_Left) {
$__t91 = ($__local_var_14_90)(($v2_15)->{'value0'});
goto end_branch_91;;
};
  if ($v2_15 instanceof \Data\Either\Data_Either_Right) {
$__t91 = ($k_13)(($v2_15)->{'value0'});
goto end_branch_91;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t91 = null;
  end_branch_91:;
  $__res = $__t91;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_12) use ($Monad1_7_6) {
  $__num = \func_num_args();
  $__local_var_13_92 = (((((($Monad1_7_6)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_13_92 = (object)["map" => function($f_14) use ($__local_var_13_92) {
  $__num = \func_num_args();
  $__local_var_15_93 = (($__local_var_13_92)->{'map'})(function($m_15) use ($f_14) {
  $__num = \func_num_args();
  $__t93 = null;;
  if ($m_15 instanceof \Data\Either\Data_Either_Left) {
$__t93 = new \Data\Either\Data_Either_Left(($m_15)->{'value0'});
goto end_branch_93;;
};
  if ($m_15 instanceof \Data\Either\Data_Either_Right) {
$__t93 = new \Data\Either\Data_Either_Right(($f_14)(($m_15)->{'value0'}));
goto end_branch_93;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t93 = null;
  end_branch_93:;
  $__res = $__t93;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_16) use ($__local_var_15_93) {
  $__num = \func_num_args();
  $__res = ($__local_var_15_93)($v_16);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_14_96 = (($Monad1_7_6)->{'Bind1'})(null);
  $pure_15_97 = ((($Monad1_7_6)->{'Applicative0'})(null))->{'pure'};
  $Bind1_14_96 = (object)["bind" => function($v_16) use ($Bind1_14_96, $pure_15_97) {
  $__num = \func_num_args();
  $__res = function($k_17) use ($Bind1_14_96, $pure_15_97, $v_16) {
  $__num = \func_num_args();
  $__local_var_18_98 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_15_97))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_14_96)->{'bind'})($v_16))(function($v2_19) use ($__local_var_18_98, $k_17) {
  $__num = \func_num_args();
  $__t99 = null;;
  if ($v2_19 instanceof \Data\Either\Data_Either_Left) {
$__t99 = ($__local_var_18_98)(($v2_19)->{'value0'});
goto end_branch_99;;
};
  if ($v2_19 instanceof \Data\Either\Data_Either_Right) {
$__t99 = ($k_17)(($v2_19)->{'value0'});
goto end_branch_99;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t99 = null;
  end_branch_99:;
  $__res = $__t99;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_16) use ($Monad1_7_6) {
  $__num = \func_num_args();
  $__local_var_17_100 = (((((($Monad1_7_6)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_17_100 = (object)["map" => function($f_18) use ($__local_var_17_100) {
  $__num = \func_num_args();
  $__local_var_19_101 = (($__local_var_17_100)->{'map'})(function($m_19) use ($f_18) {
  $__num = \func_num_args();
  $__t101 = null;;
  if ($m_19 instanceof \Data\Either\Data_Either_Left) {
$__t101 = new \Data\Either\Data_Either_Left(($m_19)->{'value0'});
goto end_branch_101;;
};
  if ($m_19 instanceof \Data\Either\Data_Either_Right) {
$__t101 = new \Data\Either\Data_Either_Right(($f_18)(($m_19)->{'value0'}));
goto end_branch_101;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t101 = null;
  end_branch_101:;
  $__res = $__t101;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_20) use ($__local_var_19_101) {
  $__num = \func_num_args();
  $__res = ($__local_var_19_101)($v_20);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_18_104 = (($Monad1_7_6)->{'Bind1'})(null);
  $pure_19_105 = ((($Monad1_7_6)->{'Applicative0'})(null))->{'pure'};
  $Bind1_18_104 = (object)["bind" => function($v_20) use ($Bind1_18_104, $pure_19_105) {
  $__num = \func_num_args();
  $__res = function($k_21) use ($Bind1_18_104, $pure_19_105, $v_20) {
  $__num = \func_num_args();
  $__local_var_22_106 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_19_105))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_18_104)->{'bind'})($v_20))(function($v2_23) use ($__local_var_22_106, $k_21) {
  $__num = \func_num_args();
  $__t107 = null;;
  if ($v2_23 instanceof \Data\Either\Data_Either_Left) {
$__t107 = ($__local_var_22_106)(($v2_23)->{'value0'});
goto end_branch_107;;
};
  if ($v2_23 instanceof \Data\Either\Data_Either_Right) {
$__t107 = ($k_21)(($v2_23)->{'value0'});
goto end_branch_107;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t107 = null;
  end_branch_107:;
  $__res = $__t107;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_20) use ($Monad1_7_6) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($Monad1_7_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_19_109 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_19) {
  $__num = \func_num_args();
  $__res = $x_19;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($Monad1_7_6)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_19) use ($Monad1_7_6) {
  $__num = \func_num_args();
  $__local_var_20_109 = (((((($Monad1_7_6)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_20_109 = (object)["map" => function($f_21) use ($__local_var_20_109) {
  $__num = \func_num_args();
  $__local_var_22_110 = (($__local_var_20_109)->{'map'})(function($m_22) use ($f_21) {
  $__num = \func_num_args();
  $__t110 = null;;
  if ($m_22 instanceof \Data\Either\Data_Either_Left) {
$__t110 = new \Data\Either\Data_Either_Left(($m_22)->{'value0'});
goto end_branch_110;;
};
  if ($m_22 instanceof \Data\Either\Data_Either_Right) {
$__t110 = new \Data\Either\Data_Either_Right(($f_21)(($m_22)->{'value0'}));
goto end_branch_110;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t110 = null;
  end_branch_110:;
  $__res = $__t110;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_23) use ($__local_var_22_110) {
  $__num = \func_num_args();
  $__res = ($__local_var_22_110)($v_23);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_21_113 = (($Monad1_7_6)->{'Bind1'})(null);
  $pure_22_114 = ((($Monad1_7_6)->{'Applicative0'})(null))->{'pure'};
  $Bind1_21_113 = (object)["bind" => function($v_23) use ($Bind1_21_113, $pure_22_114) {
  $__num = \func_num_args();
  $__res = function($k_24) use ($Bind1_21_113, $pure_22_114, $v_23) {
  $__num = \func_num_args();
  $__local_var_25_115 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_22_114))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_21_113)->{'bind'})($v_23))(function($v2_26) use ($__local_var_25_115, $k_24) {
  $__num = \func_num_args();
  $__t116 = null;;
  if ($v2_26 instanceof \Data\Either\Data_Either_Left) {
$__t116 = ($__local_var_25_115)(($v2_26)->{'value0'});
goto end_branch_116;;
};
  if ($v2_26 instanceof \Data\Either\Data_Either_Right) {
$__t116 = ($k_24)(($v2_26)->{'value0'});
goto end_branch_116;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t116 = null;
  end_branch_116:;
  $__res = $__t116;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_23) use ($Monad1_7_6) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($Monad1_7_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_22_118 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($Monad1_7_6);
  $__res = (object)["apply" => function($f_23) use ($Applicative0_22_118, $Bind1_21_113) {
  $__num = \func_num_args();
  $__res = function($a_24) use ($Applicative0_22_118, $Bind1_21_113, $f_23) {
  $__num = \func_num_args();
  $__res = ((($Bind1_21_113)->{'bind'})($f_23))(function($f_prime__25) use ($Applicative0_22_118, $Bind1_21_113, $a_24) {
  $__num = \func_num_args();
  $__res = ((($Bind1_21_113)->{'bind'})($a_24))(function($a_prime__26) use ($Applicative0_22_118, $f_prime__25) {
  $__num = \func_num_args();
  $__res = (($Applicative0_22_118)->{'pure'})(($f_prime__25)($a_prime__26));
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
}, "Functor0" => function($_dollar___unused_21) use ($functorExceptT1_20_109) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_20_109;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_20) use ($Applicative0_19_109, $Bind1_18_104) {
  $__num = \func_num_args();
  $__res = function($a_21) use ($Applicative0_19_109, $Bind1_18_104, $f_20) {
  $__num = \func_num_args();
  $__res = ((($Bind1_18_104)->{'bind'})($f_20))(function($f_prime__22) use ($Applicative0_19_109, $Bind1_18_104, $a_21) {
  $__num = \func_num_args();
  $__res = ((($Bind1_18_104)->{'bind'})($a_21))(function($a_prime__23) use ($Applicative0_19_109, $f_prime__22) {
  $__num = \func_num_args();
  $__res = (($Applicative0_19_109)->{'pure'})(($f_prime__22)($a_prime__23));
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
}, "Functor0" => function($_dollar___unused_18) use ($functorExceptT1_17_100) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_17_100;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_15_121 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_15) {
  $__num = \func_num_args();
  $__res = $x_15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($Monad1_7_6)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_15) use ($Monad1_7_6) {
  $__num = \func_num_args();
  $__local_var_16_121 = (((((($Monad1_7_6)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_16_121 = (object)["map" => function($f_17) use ($__local_var_16_121) {
  $__num = \func_num_args();
  $__local_var_18_122 = (($__local_var_16_121)->{'map'})(function($m_18) use ($f_17) {
  $__num = \func_num_args();
  $__t122 = null;;
  if ($m_18 instanceof \Data\Either\Data_Either_Left) {
$__t122 = new \Data\Either\Data_Either_Left(($m_18)->{'value0'});
goto end_branch_122;;
};
  if ($m_18 instanceof \Data\Either\Data_Either_Right) {
$__t122 = new \Data\Either\Data_Either_Right(($f_17)(($m_18)->{'value0'}));
goto end_branch_122;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t122 = null;
  end_branch_122:;
  $__res = $__t122;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_19) use ($__local_var_18_122) {
  $__num = \func_num_args();
  $__res = ($__local_var_18_122)($v_19);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_17_125 = (($Monad1_7_6)->{'Bind1'})(null);
  $pure_18_126 = ((($Monad1_7_6)->{'Applicative0'})(null))->{'pure'};
  $Bind1_17_125 = (object)["bind" => function($v_19) use ($Bind1_17_125, $pure_18_126) {
  $__num = \func_num_args();
  $__res = function($k_20) use ($Bind1_17_125, $pure_18_126, $v_19) {
  $__num = \func_num_args();
  $__local_var_21_127 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_18_126))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_17_125)->{'bind'})($v_19))(function($v2_22) use ($__local_var_21_127, $k_20) {
  $__num = \func_num_args();
  $__t128 = null;;
  if ($v2_22 instanceof \Data\Either\Data_Either_Left) {
$__t128 = ($__local_var_21_127)(($v2_22)->{'value0'});
goto end_branch_128;;
};
  if ($v2_22 instanceof \Data\Either\Data_Either_Right) {
$__t128 = ($k_20)(($v2_22)->{'value0'});
goto end_branch_128;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t128 = null;
  end_branch_128:;
  $__res = $__t128;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_19) use ($Monad1_7_6) {
  $__num = \func_num_args();
  $__local_var_20_129 = (((((($Monad1_7_6)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_20_129 = (object)["map" => function($f_21) use ($__local_var_20_129) {
  $__num = \func_num_args();
  $__local_var_22_130 = (($__local_var_20_129)->{'map'})(function($m_22) use ($f_21) {
  $__num = \func_num_args();
  $__t130 = null;;
  if ($m_22 instanceof \Data\Either\Data_Either_Left) {
$__t130 = new \Data\Either\Data_Either_Left(($m_22)->{'value0'});
goto end_branch_130;;
};
  if ($m_22 instanceof \Data\Either\Data_Either_Right) {
$__t130 = new \Data\Either\Data_Either_Right(($f_21)(($m_22)->{'value0'}));
goto end_branch_130;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t130 = null;
  end_branch_130:;
  $__res = $__t130;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_23) use ($__local_var_22_130) {
  $__num = \func_num_args();
  $__res = ($__local_var_22_130)($v_23);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_21_133 = (($Monad1_7_6)->{'Bind1'})(null);
  $pure_22_134 = ((($Monad1_7_6)->{'Applicative0'})(null))->{'pure'};
  $Bind1_21_133 = (object)["bind" => function($v_23) use ($Bind1_21_133, $pure_22_134) {
  $__num = \func_num_args();
  $__res = function($k_24) use ($Bind1_21_133, $pure_22_134, $v_23) {
  $__num = \func_num_args();
  $__local_var_25_135 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_22_134))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_21_133)->{'bind'})($v_23))(function($v2_26) use ($__local_var_25_135, $k_24) {
  $__num = \func_num_args();
  $__t136 = null;;
  if ($v2_26 instanceof \Data\Either\Data_Either_Left) {
$__t136 = ($__local_var_25_135)(($v2_26)->{'value0'});
goto end_branch_136;;
};
  if ($v2_26 instanceof \Data\Either\Data_Either_Right) {
$__t136 = ($k_24)(($v2_26)->{'value0'});
goto end_branch_136;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t136 = null;
  end_branch_136:;
  $__res = $__t136;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_23) use ($Monad1_7_6) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($Monad1_7_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_22_138 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_22) {
  $__num = \func_num_args();
  $__res = $x_22;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($Monad1_7_6)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_22) use ($Monad1_7_6) {
  $__num = \func_num_args();
  $__local_var_23_138 = (((((($Monad1_7_6)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_23_138 = (object)["map" => function($f_24) use ($__local_var_23_138) {
  $__num = \func_num_args();
  $__local_var_25_139 = (($__local_var_23_138)->{'map'})(function($m_25) use ($f_24) {
  $__num = \func_num_args();
  $__t139 = null;;
  if ($m_25 instanceof \Data\Either\Data_Either_Left) {
$__t139 = new \Data\Either\Data_Either_Left(($m_25)->{'value0'});
goto end_branch_139;;
};
  if ($m_25 instanceof \Data\Either\Data_Either_Right) {
$__t139 = new \Data\Either\Data_Either_Right(($f_24)(($m_25)->{'value0'}));
goto end_branch_139;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t139 = null;
  end_branch_139:;
  $__res = $__t139;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_26) use ($__local_var_25_139) {
  $__num = \func_num_args();
  $__res = ($__local_var_25_139)($v_26);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_24_142 = (($Monad1_7_6)->{'Bind1'})(null);
  $pure_25_143 = ((($Monad1_7_6)->{'Applicative0'})(null))->{'pure'};
  $Bind1_24_142 = (object)["bind" => function($v_26) use ($Bind1_24_142, $pure_25_143) {
  $__num = \func_num_args();
  $__res = function($k_27) use ($Bind1_24_142, $pure_25_143, $v_26) {
  $__num = \func_num_args();
  $__local_var_28_144 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_25_143))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_24_142)->{'bind'})($v_26))(function($v2_29) use ($__local_var_28_144, $k_27) {
  $__num = \func_num_args();
  $__t145 = null;;
  if ($v2_29 instanceof \Data\Either\Data_Either_Left) {
$__t145 = ($__local_var_28_144)(($v2_29)->{'value0'});
goto end_branch_145;;
};
  if ($v2_29 instanceof \Data\Either\Data_Either_Right) {
$__t145 = ($k_27)(($v2_29)->{'value0'});
goto end_branch_145;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t145 = null;
  end_branch_145:;
  $__res = $__t145;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_26) use ($Monad1_7_6) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($Monad1_7_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_25_147 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($Monad1_7_6);
  $__res = (object)["apply" => function($f_26) use ($Applicative0_25_147, $Bind1_24_142) {
  $__num = \func_num_args();
  $__res = function($a_27) use ($Applicative0_25_147, $Bind1_24_142, $f_26) {
  $__num = \func_num_args();
  $__res = ((($Bind1_24_142)->{'bind'})($f_26))(function($f_prime__28) use ($Applicative0_25_147, $Bind1_24_142, $a_27) {
  $__num = \func_num_args();
  $__res = ((($Bind1_24_142)->{'bind'})($a_27))(function($a_prime__29) use ($Applicative0_25_147, $f_prime__28) {
  $__num = \func_num_args();
  $__res = (($Applicative0_25_147)->{'pure'})(($f_prime__28)($a_prime__29));
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
}, "Functor0" => function($_dollar___unused_24) use ($functorExceptT1_23_138) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_23_138;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_23) use ($Applicative0_22_138, $Bind1_21_133) {
  $__num = \func_num_args();
  $__res = function($a_24) use ($Applicative0_22_138, $Bind1_21_133, $f_23) {
  $__num = \func_num_args();
  $__res = ((($Bind1_21_133)->{'bind'})($f_23))(function($f_prime__25) use ($Applicative0_22_138, $Bind1_21_133, $a_24) {
  $__num = \func_num_args();
  $__res = ((($Bind1_21_133)->{'bind'})($a_24))(function($a_prime__26) use ($Applicative0_22_138, $f_prime__25) {
  $__num = \func_num_args();
  $__res = (($Applicative0_22_138)->{'pure'})(($f_prime__25)($a_prime__26));
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
}, "Functor0" => function($_dollar___unused_21) use ($functorExceptT1_20_129) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_20_129;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_18_150 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_18) {
  $__num = \func_num_args();
  $__res = $x_18;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($Monad1_7_6)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_18) use ($Monad1_7_6) {
  $__num = \func_num_args();
  $__local_var_19_150 = (((((($Monad1_7_6)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_19_150 = (object)["map" => function($f_20) use ($__local_var_19_150) {
  $__num = \func_num_args();
  $__local_var_21_151 = (($__local_var_19_150)->{'map'})(function($m_21) use ($f_20) {
  $__num = \func_num_args();
  $__t151 = null;;
  if ($m_21 instanceof \Data\Either\Data_Either_Left) {
$__t151 = new \Data\Either\Data_Either_Left(($m_21)->{'value0'});
goto end_branch_151;;
};
  if ($m_21 instanceof \Data\Either\Data_Either_Right) {
$__t151 = new \Data\Either\Data_Either_Right(($f_20)(($m_21)->{'value0'}));
goto end_branch_151;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t151 = null;
  end_branch_151:;
  $__res = $__t151;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_22) use ($__local_var_21_151) {
  $__num = \func_num_args();
  $__res = ($__local_var_21_151)($v_22);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_20_154 = (($Monad1_7_6)->{'Bind1'})(null);
  $pure_21_155 = ((($Monad1_7_6)->{'Applicative0'})(null))->{'pure'};
  $Bind1_20_154 = (object)["bind" => function($v_22) use ($Bind1_20_154, $pure_21_155) {
  $__num = \func_num_args();
  $__res = function($k_23) use ($Bind1_20_154, $pure_21_155, $v_22) {
  $__num = \func_num_args();
  $__local_var_24_156 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_21_155))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_20_154)->{'bind'})($v_22))(function($v2_25) use ($__local_var_24_156, $k_23) {
  $__num = \func_num_args();
  $__t157 = null;;
  if ($v2_25 instanceof \Data\Either\Data_Either_Left) {
$__t157 = ($__local_var_24_156)(($v2_25)->{'value0'});
goto end_branch_157;;
};
  if ($v2_25 instanceof \Data\Either\Data_Either_Right) {
$__t157 = ($k_23)(($v2_25)->{'value0'});
goto end_branch_157;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t157 = null;
  end_branch_157:;
  $__res = $__t157;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_22) use ($Monad1_7_6) {
  $__num = \func_num_args();
  $__local_var_23_158 = (((((($Monad1_7_6)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_23_158 = (object)["map" => function($f_24) use ($__local_var_23_158) {
  $__num = \func_num_args();
  $__local_var_25_159 = (($__local_var_23_158)->{'map'})(function($m_25) use ($f_24) {
  $__num = \func_num_args();
  $__t159 = null;;
  if ($m_25 instanceof \Data\Either\Data_Either_Left) {
$__t159 = new \Data\Either\Data_Either_Left(($m_25)->{'value0'});
goto end_branch_159;;
};
  if ($m_25 instanceof \Data\Either\Data_Either_Right) {
$__t159 = new \Data\Either\Data_Either_Right(($f_24)(($m_25)->{'value0'}));
goto end_branch_159;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t159 = null;
  end_branch_159:;
  $__res = $__t159;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_26) use ($__local_var_25_159) {
  $__num = \func_num_args();
  $__res = ($__local_var_25_159)($v_26);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_24_162 = (($Monad1_7_6)->{'Bind1'})(null);
  $pure_25_163 = ((($Monad1_7_6)->{'Applicative0'})(null))->{'pure'};
  $Bind1_24_162 = (object)["bind" => function($v_26) use ($Bind1_24_162, $pure_25_163) {
  $__num = \func_num_args();
  $__res = function($k_27) use ($Bind1_24_162, $pure_25_163, $v_26) {
  $__num = \func_num_args();
  $__local_var_28_164 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_25_163))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_24_162)->{'bind'})($v_26))(function($v2_29) use ($__local_var_28_164, $k_27) {
  $__num = \func_num_args();
  $__t165 = null;;
  if ($v2_29 instanceof \Data\Either\Data_Either_Left) {
$__t165 = ($__local_var_28_164)(($v2_29)->{'value0'});
goto end_branch_165;;
};
  if ($v2_29 instanceof \Data\Either\Data_Either_Right) {
$__t165 = ($k_27)(($v2_29)->{'value0'});
goto end_branch_165;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t165 = null;
  end_branch_165:;
  $__res = $__t165;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_26) use ($Monad1_7_6) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($Monad1_7_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_25_167 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($Monad1_7_6);
  $__res = (object)["apply" => function($f_26) use ($Applicative0_25_167, $Bind1_24_162) {
  $__num = \func_num_args();
  $__res = function($a_27) use ($Applicative0_25_167, $Bind1_24_162, $f_26) {
  $__num = \func_num_args();
  $__res = ((($Bind1_24_162)->{'bind'})($f_26))(function($f_prime__28) use ($Applicative0_25_167, $Bind1_24_162, $a_27) {
  $__num = \func_num_args();
  $__res = ((($Bind1_24_162)->{'bind'})($a_27))(function($a_prime__29) use ($Applicative0_25_167, $f_prime__28) {
  $__num = \func_num_args();
  $__res = (($Applicative0_25_167)->{'pure'})(($f_prime__28)($a_prime__29));
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
}, "Functor0" => function($_dollar___unused_24) use ($functorExceptT1_23_158) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_23_158;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_21_169 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($Monad1_7_6);
  $__res = (object)["apply" => function($f_22) use ($Applicative0_21_169, $Bind1_20_154) {
  $__num = \func_num_args();
  $__res = function($a_23) use ($Applicative0_21_169, $Bind1_20_154, $f_22) {
  $__num = \func_num_args();
  $__res = ((($Bind1_20_154)->{'bind'})($f_22))(function($f_prime__24) use ($Applicative0_21_169, $Bind1_20_154, $a_23) {
  $__num = \func_num_args();
  $__res = ((($Bind1_20_154)->{'bind'})($a_23))(function($a_prime__25) use ($Applicative0_21_169, $f_prime__24) {
  $__num = \func_num_args();
  $__res = (($Applicative0_21_169)->{'pure'})(($f_prime__24)($a_prime__25));
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
}, "Functor0" => function($_dollar___unused_20) use ($functorExceptT1_19_150) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_19_150;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_19) use ($Applicative0_18_150, $Bind1_17_125) {
  $__num = \func_num_args();
  $__res = function($a_20) use ($Applicative0_18_150, $Bind1_17_125, $f_19) {
  $__num = \func_num_args();
  $__res = ((($Bind1_17_125)->{'bind'})($f_19))(function($f_prime__21) use ($Applicative0_18_150, $Bind1_17_125, $a_20) {
  $__num = \func_num_args();
  $__res = ((($Bind1_17_125)->{'bind'})($a_20))(function($a_prime__22) use ($Applicative0_18_150, $f_prime__21) {
  $__num = \func_num_args();
  $__res = (($Applicative0_18_150)->{'pure'})(($f_prime__21)($a_prime__22));
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
}, "Functor0" => function($_dollar___unused_17) use ($functorExceptT1_16_121) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_16_121;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_16) use ($Applicative0_15_121, $Bind1_14_96) {
  $__num = \func_num_args();
  $__res = function($a_17) use ($Applicative0_15_121, $Bind1_14_96, $f_16) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_96)->{'bind'})($f_16))(function($f_prime__18) use ($Applicative0_15_121, $Bind1_14_96, $a_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_96)->{'bind'})($a_17))(function($a_prime__19) use ($Applicative0_15_121, $f_prime__18) {
  $__num = \func_num_args();
  $__res = (($Applicative0_15_121)->{'pure'})(($f_prime__18)($a_prime__19));
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
}, "Functor0" => function($_dollar___unused_14) use ($functorExceptT1_13_92) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_13_92;
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
  $Bind1_10_173 = (($Monad1_7_6)->{'Bind1'})(null);
  $pure_11_174 = ((($Monad1_7_6)->{'Applicative0'})(null))->{'pure'};
  $monadTellExceptT1_7_6 = (object)["tell" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($m_12) use ($Bind1_10_173, $pure_11_174) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_173)->{'bind'})($m_12))(function($a_13) use ($pure_11_174) {
  $__num = \func_num_args();
  $__res = ($pure_11_174)(new \Data\Either\Data_Either_Right($a_13));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($MonadTell1_1_0)->{'tell'}), "Semigroup0" => function($_dollar___unused_10) use ($Semigroup0_8_7) {
  $__num = \func_num_args();
  $__res = $Semigroup0_8_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad1" => function($_dollar___unused_10) use ($monadExceptT1_9_8) {
  $__num = \func_num_args();
  $__res = $monadExceptT1_9_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["listen" => function($v_8) use ($Bind1_3_2, $dictMonadWriter_0, $pure_4_3) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_2)->{'bind'})((($dictMonadWriter_0)->{'listen'})($v_8)))(function($v_9) use ($pure_4_3) {
  $__num = \func_num_args();
  $__t176 = null;;
  if (($v_9)->{'value0'} instanceof \Data\Either\Data_Either_Left) {
$__t176 = new \Data\Either\Data_Either_Left((($v_9)->{'value0'})->{'value0'});
goto end_branch_176;;
};
  if (($v_9)->{'value0'} instanceof \Data\Either\Data_Either_Right) {
$__t176 = new \Data\Either\Data_Either_Right(new \Data\Tuple\Data_Tuple_Tuple((($v_9)->{'value0'})->{'value0'}, ($v_9)->{'value1'}));
goto end_branch_176;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t176 = null;
  end_branch_176:;
  $__res = ($pure_4_3)($__t176);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "pass" => function($v_8) use ($Applicative0_5_4, $Bind1_3_2, $dictMonadWriter_0) {
  $__num = \func_num_args();
  $__res = (($dictMonadWriter_0)->{'pass'})(((($Bind1_3_2)->{'bind'})($v_8))(function($a_9) use ($Applicative0_5_4) {
  $__num = \func_num_args();
  $__t177 = null;;
  if ($a_9 instanceof \Data\Either\Data_Either_Left) {
$__t177 = new \Data\Tuple\Data_Tuple_Tuple(new \Data\Either\Data_Either_Left(($a_9)->{'value0'}), function($x_10) {
  $__num = \func_num_args();
  $__res = $x_10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
goto end_branch_177;;
};
  if ($a_9 instanceof \Data\Either\Data_Either_Right) {
$__t177 = new \Data\Tuple\Data_Tuple_Tuple(new \Data\Either\Data_Either_Right((($a_9)->{'value0'})->{'value0'}), (($a_9)->{'value0'})->{'value1'});
goto end_branch_177;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t177 = null;
  end_branch_177:;
  $__res = (($Applicative0_5_4)->{'pure'})($__t177);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monoid0" => function($_dollar___unused_8) use ($Monoid0_6_5) {
  $__num = \func_num_args();
  $__res = $Monoid0_6_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "MonadTell1" => function($_dollar___unused_8) use ($monadTellExceptT1_7_6) {
  $__num = \func_num_args();
  $__res = $monadTellExceptT1_7_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Except_Trans_monadWriterExceptT'] = __NAMESPACE__ . '\\majControl_majMonad_majExcept_majTrans_monadmajWritermajExceptmajT';

// Control_Monad_Except_Trans_monadThrowExceptT
function majControl_majMonad_majExcept_majTrans_monadmajThrowmajExceptmajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majExcept_majTrans_monadmajThrowmajExceptmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $monadExceptT1_1_0 = (object)["Applicative0" => function($_dollar___unused_1) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_2) {
  $__num = \func_num_args();
  $__res = $x_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_2) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_3_0 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_3_0 = (object)["map" => function($f_4) use ($__local_var_3_0) {
  $__num = \func_num_args();
  $__local_var_5_1 = (($__local_var_3_0)->{'map'})(function($m_5) use ($f_4) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($m_5 instanceof \Data\Either\Data_Either_Left) {
$__t1 = new \Data\Either\Data_Either_Left(($m_5)->{'value0'});
goto end_branch_1;;
};
  if ($m_5 instanceof \Data\Either\Data_Either_Right) {
$__t1 = new \Data\Either\Data_Either_Right(($f_4)(($m_5)->{'value0'}));
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_6) use ($__local_var_5_1) {
  $__num = \func_num_args();
  $__res = ($__local_var_5_1)($v_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_4_4 = (($dictMonad_0)->{'Bind1'})(null);
  $pure_5_5 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_4_4 = (object)["bind" => function($v_6) use ($Bind1_4_4, $pure_5_5) {
  $__num = \func_num_args();
  $__res = function($k_7) use ($Bind1_4_4, $pure_5_5, $v_6) {
  $__num = \func_num_args();
  $__local_var_8_6 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_5_5))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_4_4)->{'bind'})($v_6))(function($v2_9) use ($__local_var_8_6, $k_7) {
  $__num = \func_num_args();
  $__t7 = null;;
  if ($v2_9 instanceof \Data\Either\Data_Either_Left) {
$__t7 = ($__local_var_8_6)(($v2_9)->{'value0'});
goto end_branch_7;;
};
  if ($v2_9 instanceof \Data\Either\Data_Either_Right) {
$__t7 = ($k_7)(($v2_9)->{'value0'});
goto end_branch_7;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t7 = null;
  end_branch_7:;
  $__res = $__t7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_6) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_7_8 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_7_8 = (object)["map" => function($f_8) use ($__local_var_7_8) {
  $__num = \func_num_args();
  $__local_var_9_9 = (($__local_var_7_8)->{'map'})(function($m_9) use ($f_8) {
  $__num = \func_num_args();
  $__t9 = null;;
  if ($m_9 instanceof \Data\Either\Data_Either_Left) {
$__t9 = new \Data\Either\Data_Either_Left(($m_9)->{'value0'});
goto end_branch_9;;
};
  if ($m_9 instanceof \Data\Either\Data_Either_Right) {
$__t9 = new \Data\Either\Data_Either_Right(($f_8)(($m_9)->{'value0'}));
goto end_branch_9;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t9 = null;
  end_branch_9:;
  $__res = $__t9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_10) use ($__local_var_9_9) {
  $__num = \func_num_args();
  $__res = ($__local_var_9_9)($v_10);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_8_12 = (($dictMonad_0)->{'Bind1'})(null);
  $pure_9_13 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_8_12 = (object)["bind" => function($v_10) use ($Bind1_8_12, $pure_9_13) {
  $__num = \func_num_args();
  $__res = function($k_11) use ($Bind1_8_12, $pure_9_13, $v_10) {
  $__num = \func_num_args();
  $__local_var_12_14 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_9_13))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_8_12)->{'bind'})($v_10))(function($v2_13) use ($__local_var_12_14, $k_11) {
  $__num = \func_num_args();
  $__t15 = null;;
  if ($v2_13 instanceof \Data\Either\Data_Either_Left) {
$__t15 = ($__local_var_12_14)(($v2_13)->{'value0'});
goto end_branch_15;;
};
  if ($v2_13 instanceof \Data\Either\Data_Either_Right) {
$__t15 = ($k_11)(($v2_13)->{'value0'});
goto end_branch_15;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t15 = null;
  end_branch_15:;
  $__res = $__t15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_10) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_9_17 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_10) use ($Applicative0_9_17, $Bind1_8_12) {
  $__num = \func_num_args();
  $__res = function($a_11) use ($Applicative0_9_17, $Bind1_8_12, $f_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_8_12)->{'bind'})($f_10))(function($f_prime__12) use ($Applicative0_9_17, $Bind1_8_12, $a_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_8_12)->{'bind'})($a_11))(function($a_prime__13) use ($Applicative0_9_17, $f_prime__12) {
  $__num = \func_num_args();
  $__res = (($Applicative0_9_17)->{'pure'})(($f_prime__12)($a_prime__13));
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
}, "Functor0" => function($_dollar___unused_8) use ($functorExceptT1_7_8) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_7_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_5_19 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_6) use ($Applicative0_5_19, $Bind1_4_4) {
  $__num = \func_num_args();
  $__res = function($a_7) use ($Applicative0_5_19, $Bind1_4_4, $f_6) {
  $__num = \func_num_args();
  $__res = ((($Bind1_4_4)->{'bind'})($f_6))(function($f_prime__8) use ($Applicative0_5_19, $Bind1_4_4, $a_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_4_4)->{'bind'})($a_7))(function($a_prime__9) use ($Applicative0_5_19, $f_prime__8) {
  $__num = \func_num_args();
  $__res = (($Applicative0_5_19)->{'pure'})(($f_prime__8)($a_prime__9));
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
}, "Functor0" => function($_dollar___unused_4) use ($functorExceptT1_3_0) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_3_0;
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
}, "Bind1" => function($_dollar___unused_1) use ($dictMonad_0) {
  $__num = \func_num_args();
  $Bind1_2_20 = (($dictMonad_0)->{'Bind1'})(null);
  $pure_3_21 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $__res = (object)["bind" => function($v_4) use ($Bind1_2_20, $pure_3_21) {
  $__num = \func_num_args();
  $__res = function($k_5) use ($Bind1_2_20, $pure_3_21, $v_4) {
  $__num = \func_num_args();
  $__local_var_6_22 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_3_21))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_2_20)->{'bind'})($v_4))(function($v2_7) use ($__local_var_6_22, $k_5) {
  $__num = \func_num_args();
  $__t23 = null;;
  if ($v2_7 instanceof \Data\Either\Data_Either_Left) {
$__t23 = ($__local_var_6_22)(($v2_7)->{'value0'});
goto end_branch_23;;
};
  if ($v2_7 instanceof \Data\Either\Data_Either_Right) {
$__t23 = ($k_5)(($v2_7)->{'value0'});
goto end_branch_23;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t23 = null;
  end_branch_23:;
  $__res = $__t23;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_4) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_5_24 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_5_24 = (object)["map" => function($f_6) use ($__local_var_5_24) {
  $__num = \func_num_args();
  $__local_var_7_25 = (($__local_var_5_24)->{'map'})(function($m_7) use ($f_6) {
  $__num = \func_num_args();
  $__t25 = null;;
  if ($m_7 instanceof \Data\Either\Data_Either_Left) {
$__t25 = new \Data\Either\Data_Either_Left(($m_7)->{'value0'});
goto end_branch_25;;
};
  if ($m_7 instanceof \Data\Either\Data_Either_Right) {
$__t25 = new \Data\Either\Data_Either_Right(($f_6)(($m_7)->{'value0'}));
goto end_branch_25;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t25 = null;
  end_branch_25:;
  $__res = $__t25;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_8) use ($__local_var_7_25) {
  $__num = \func_num_args();
  $__res = ($__local_var_7_25)($v_8);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_6_28 = (($dictMonad_0)->{'Bind1'})(null);
  $pure_7_29 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_6_28 = (object)["bind" => function($v_8) use ($Bind1_6_28, $pure_7_29) {
  $__num = \func_num_args();
  $__res = function($k_9) use ($Bind1_6_28, $pure_7_29, $v_8) {
  $__num = \func_num_args();
  $__local_var_10_30 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_7_29))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_6_28)->{'bind'})($v_8))(function($v2_11) use ($__local_var_10_30, $k_9) {
  $__num = \func_num_args();
  $__t31 = null;;
  if ($v2_11 instanceof \Data\Either\Data_Either_Left) {
$__t31 = ($__local_var_10_30)(($v2_11)->{'value0'});
goto end_branch_31;;
};
  if ($v2_11 instanceof \Data\Either\Data_Either_Right) {
$__t31 = ($k_9)(($v2_11)->{'value0'});
goto end_branch_31;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t31 = null;
  end_branch_31:;
  $__res = $__t31;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_8) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_7_33 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_7) {
  $__num = \func_num_args();
  $__res = $x_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_7) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_8_33 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_8_33 = (object)["map" => function($f_9) use ($__local_var_8_33) {
  $__num = \func_num_args();
  $__local_var_10_34 = (($__local_var_8_33)->{'map'})(function($m_10) use ($f_9) {
  $__num = \func_num_args();
  $__t34 = null;;
  if ($m_10 instanceof \Data\Either\Data_Either_Left) {
$__t34 = new \Data\Either\Data_Either_Left(($m_10)->{'value0'});
goto end_branch_34;;
};
  if ($m_10 instanceof \Data\Either\Data_Either_Right) {
$__t34 = new \Data\Either\Data_Either_Right(($f_9)(($m_10)->{'value0'}));
goto end_branch_34;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t34 = null;
  end_branch_34:;
  $__res = $__t34;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_11) use ($__local_var_10_34) {
  $__num = \func_num_args();
  $__res = ($__local_var_10_34)($v_11);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_9_37 = (($dictMonad_0)->{'Bind1'})(null);
  $pure_10_38 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_9_37 = (object)["bind" => function($v_11) use ($Bind1_9_37, $pure_10_38) {
  $__num = \func_num_args();
  $__res = function($k_12) use ($Bind1_9_37, $pure_10_38, $v_11) {
  $__num = \func_num_args();
  $__local_var_13_39 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_10_38))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_9_37)->{'bind'})($v_11))(function($v2_14) use ($__local_var_13_39, $k_12) {
  $__num = \func_num_args();
  $__t40 = null;;
  if ($v2_14 instanceof \Data\Either\Data_Either_Left) {
$__t40 = ($__local_var_13_39)(($v2_14)->{'value0'});
goto end_branch_40;;
};
  if ($v2_14 instanceof \Data\Either\Data_Either_Right) {
$__t40 = ($k_12)(($v2_14)->{'value0'});
goto end_branch_40;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t40 = null;
  end_branch_40:;
  $__res = $__t40;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_11) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_10_42 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_11) use ($Applicative0_10_42, $Bind1_9_37) {
  $__num = \func_num_args();
  $__res = function($a_12) use ($Applicative0_10_42, $Bind1_9_37, $f_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_37)->{'bind'})($f_11))(function($f_prime__13) use ($Applicative0_10_42, $Bind1_9_37, $a_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_37)->{'bind'})($a_12))(function($a_prime__14) use ($Applicative0_10_42, $f_prime__13) {
  $__num = \func_num_args();
  $__res = (($Applicative0_10_42)->{'pure'})(($f_prime__13)($a_prime__14));
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
}, "Functor0" => function($_dollar___unused_9) use ($functorExceptT1_8_33) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_8_33;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_8) use ($Applicative0_7_33, $Bind1_6_28) {
  $__num = \func_num_args();
  $__res = function($a_9) use ($Applicative0_7_33, $Bind1_6_28, $f_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_28)->{'bind'})($f_8))(function($f_prime__10) use ($Applicative0_7_33, $Bind1_6_28, $a_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_28)->{'bind'})($a_9))(function($a_prime__11) use ($Applicative0_7_33, $f_prime__10) {
  $__num = \func_num_args();
  $__res = (($Applicative0_7_33)->{'pure'})(($f_prime__10)($a_prime__11));
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
}, "Functor0" => function($_dollar___unused_6) use ($functorExceptT1_5_24) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_5_24;
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
  $__res = (object)["throwError" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_2) {
  $__num = \func_num_args();
  $__res = $x_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Left'])), "Monad0" => function($_dollar___unused_2) use ($monadExceptT1_1_0) {
  $__num = \func_num_args();
  $__res = $monadExceptT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Except_Trans_monadThrowExceptT'] = __NAMESPACE__ . '\\majControl_majMonad_majExcept_majTrans_monadmajThrowmajExceptmajT';

// Control_Monad_Except_Trans_monadErrorExceptT
function majControl_majMonad_majExcept_majTrans_monadmajErrormajExceptmajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majExcept_majTrans_monadmajErrormajExceptmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Bind1_1_0 = (($dictMonad_0)->{'Bind1'})(null);
  $pure_2_1 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $monadExceptT1_3_2 = (object)["Applicative0" => function($_dollar___unused_3) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_4) {
  $__num = \func_num_args();
  $__res = $x_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_4) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_5_2 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_5_2 = (object)["map" => function($f_6) use ($__local_var_5_2) {
  $__num = \func_num_args();
  $__local_var_7_3 = (($__local_var_5_2)->{'map'})(function($m_7) use ($f_6) {
  $__num = \func_num_args();
  $__t3 = null;;
  if ($m_7 instanceof \Data\Either\Data_Either_Left) {
$__t3 = new \Data\Either\Data_Either_Left(($m_7)->{'value0'});
goto end_branch_3;;
};
  if ($m_7 instanceof \Data\Either\Data_Either_Right) {
$__t3 = new \Data\Either\Data_Either_Right(($f_6)(($m_7)->{'value0'}));
goto end_branch_3;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t3 = null;
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_8) use ($__local_var_7_3) {
  $__num = \func_num_args();
  $__res = ($__local_var_7_3)($v_8);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_6_6 = (($dictMonad_0)->{'Bind1'})(null);
  $pure_7_7 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_6_6 = (object)["bind" => function($v_8) use ($Bind1_6_6, $pure_7_7) {
  $__num = \func_num_args();
  $__res = function($k_9) use ($Bind1_6_6, $pure_7_7, $v_8) {
  $__num = \func_num_args();
  $__local_var_10_8 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_7_7))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_6_6)->{'bind'})($v_8))(function($v2_11) use ($__local_var_10_8, $k_9) {
  $__num = \func_num_args();
  $__t9 = null;;
  if ($v2_11 instanceof \Data\Either\Data_Either_Left) {
$__t9 = ($__local_var_10_8)(($v2_11)->{'value0'});
goto end_branch_9;;
};
  if ($v2_11 instanceof \Data\Either\Data_Either_Right) {
$__t9 = ($k_9)(($v2_11)->{'value0'});
goto end_branch_9;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t9 = null;
  end_branch_9:;
  $__res = $__t9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_8) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_9_10 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_9_10 = (object)["map" => function($f_10) use ($__local_var_9_10) {
  $__num = \func_num_args();
  $__local_var_11_11 = (($__local_var_9_10)->{'map'})(function($m_11) use ($f_10) {
  $__num = \func_num_args();
  $__t11 = null;;
  if ($m_11 instanceof \Data\Either\Data_Either_Left) {
$__t11 = new \Data\Either\Data_Either_Left(($m_11)->{'value0'});
goto end_branch_11;;
};
  if ($m_11 instanceof \Data\Either\Data_Either_Right) {
$__t11 = new \Data\Either\Data_Either_Right(($f_10)(($m_11)->{'value0'}));
goto end_branch_11;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t11 = null;
  end_branch_11:;
  $__res = $__t11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_12) use ($__local_var_11_11) {
  $__num = \func_num_args();
  $__res = ($__local_var_11_11)($v_12);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_10_14 = (($dictMonad_0)->{'Bind1'})(null);
  $pure_11_15 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_10_14 = (object)["bind" => function($v_12) use ($Bind1_10_14, $pure_11_15) {
  $__num = \func_num_args();
  $__res = function($k_13) use ($Bind1_10_14, $pure_11_15, $v_12) {
  $__num = \func_num_args();
  $__local_var_14_16 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_11_15))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_10_14)->{'bind'})($v_12))(function($v2_15) use ($__local_var_14_16, $k_13) {
  $__num = \func_num_args();
  $__t17 = null;;
  if ($v2_15 instanceof \Data\Either\Data_Either_Left) {
$__t17 = ($__local_var_14_16)(($v2_15)->{'value0'});
goto end_branch_17;;
};
  if ($v2_15 instanceof \Data\Either\Data_Either_Right) {
$__t17 = ($k_13)(($v2_15)->{'value0'});
goto end_branch_17;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t17 = null;
  end_branch_17:;
  $__res = $__t17;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_12) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_13_18 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_13_18 = (object)["map" => function($f_14) use ($__local_var_13_18) {
  $__num = \func_num_args();
  $__local_var_15_19 = (($__local_var_13_18)->{'map'})(function($m_15) use ($f_14) {
  $__num = \func_num_args();
  $__t19 = null;;
  if ($m_15 instanceof \Data\Either\Data_Either_Left) {
$__t19 = new \Data\Either\Data_Either_Left(($m_15)->{'value0'});
goto end_branch_19;;
};
  if ($m_15 instanceof \Data\Either\Data_Either_Right) {
$__t19 = new \Data\Either\Data_Either_Right(($f_14)(($m_15)->{'value0'}));
goto end_branch_19;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t19 = null;
  end_branch_19:;
  $__res = $__t19;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_16) use ($__local_var_15_19) {
  $__num = \func_num_args();
  $__res = ($__local_var_15_19)($v_16);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_14_22 = (($dictMonad_0)->{'Bind1'})(null);
  $pure_15_23 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_14_22 = (object)["bind" => function($v_16) use ($Bind1_14_22, $pure_15_23) {
  $__num = \func_num_args();
  $__res = function($k_17) use ($Bind1_14_22, $pure_15_23, $v_16) {
  $__num = \func_num_args();
  $__local_var_18_24 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_15_23))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_14_22)->{'bind'})($v_16))(function($v2_19) use ($__local_var_18_24, $k_17) {
  $__num = \func_num_args();
  $__t25 = null;;
  if ($v2_19 instanceof \Data\Either\Data_Either_Left) {
$__t25 = ($__local_var_18_24)(($v2_19)->{'value0'});
goto end_branch_25;;
};
  if ($v2_19 instanceof \Data\Either\Data_Either_Right) {
$__t25 = ($k_17)(($v2_19)->{'value0'});
goto end_branch_25;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t25 = null;
  end_branch_25:;
  $__res = $__t25;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_16) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_15_27 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_15) {
  $__num = \func_num_args();
  $__res = $x_15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_15) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_16_27 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_16_27 = (object)["map" => function($f_17) use ($__local_var_16_27) {
  $__num = \func_num_args();
  $__local_var_18_28 = (($__local_var_16_27)->{'map'})(function($m_18) use ($f_17) {
  $__num = \func_num_args();
  $__t28 = null;;
  if ($m_18 instanceof \Data\Either\Data_Either_Left) {
$__t28 = new \Data\Either\Data_Either_Left(($m_18)->{'value0'});
goto end_branch_28;;
};
  if ($m_18 instanceof \Data\Either\Data_Either_Right) {
$__t28 = new \Data\Either\Data_Either_Right(($f_17)(($m_18)->{'value0'}));
goto end_branch_28;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t28 = null;
  end_branch_28:;
  $__res = $__t28;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_19) use ($__local_var_18_28) {
  $__num = \func_num_args();
  $__res = ($__local_var_18_28)($v_19);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_17_31 = (($dictMonad_0)->{'Bind1'})(null);
  $pure_18_32 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_17_31 = (object)["bind" => function($v_19) use ($Bind1_17_31, $pure_18_32) {
  $__num = \func_num_args();
  $__res = function($k_20) use ($Bind1_17_31, $pure_18_32, $v_19) {
  $__num = \func_num_args();
  $__local_var_21_33 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_18_32))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_17_31)->{'bind'})($v_19))(function($v2_22) use ($__local_var_21_33, $k_20) {
  $__num = \func_num_args();
  $__t34 = null;;
  if ($v2_22 instanceof \Data\Either\Data_Either_Left) {
$__t34 = ($__local_var_21_33)(($v2_22)->{'value0'});
goto end_branch_34;;
};
  if ($v2_22 instanceof \Data\Either\Data_Either_Right) {
$__t34 = ($k_20)(($v2_22)->{'value0'});
goto end_branch_34;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t34 = null;
  end_branch_34:;
  $__res = $__t34;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_19) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_18_36 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_19) use ($Applicative0_18_36, $Bind1_17_31) {
  $__num = \func_num_args();
  $__res = function($a_20) use ($Applicative0_18_36, $Bind1_17_31, $f_19) {
  $__num = \func_num_args();
  $__res = ((($Bind1_17_31)->{'bind'})($f_19))(function($f_prime__21) use ($Applicative0_18_36, $Bind1_17_31, $a_20) {
  $__num = \func_num_args();
  $__res = ((($Bind1_17_31)->{'bind'})($a_20))(function($a_prime__22) use ($Applicative0_18_36, $f_prime__21) {
  $__num = \func_num_args();
  $__res = (($Applicative0_18_36)->{'pure'})(($f_prime__21)($a_prime__22));
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
}, "Functor0" => function($_dollar___unused_17) use ($functorExceptT1_16_27) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_16_27;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_16) use ($Applicative0_15_27, $Bind1_14_22) {
  $__num = \func_num_args();
  $__res = function($a_17) use ($Applicative0_15_27, $Bind1_14_22, $f_16) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_22)->{'bind'})($f_16))(function($f_prime__18) use ($Applicative0_15_27, $Bind1_14_22, $a_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_22)->{'bind'})($a_17))(function($a_prime__19) use ($Applicative0_15_27, $f_prime__18) {
  $__num = \func_num_args();
  $__res = (($Applicative0_15_27)->{'pure'})(($f_prime__18)($a_prime__19));
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
}, "Functor0" => function($_dollar___unused_14) use ($functorExceptT1_13_18) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_13_18;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_11_39 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_11) {
  $__num = \func_num_args();
  $__res = $x_11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_11) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_12_39 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_12_39 = (object)["map" => function($f_13) use ($__local_var_12_39) {
  $__num = \func_num_args();
  $__local_var_14_40 = (($__local_var_12_39)->{'map'})(function($m_14) use ($f_13) {
  $__num = \func_num_args();
  $__t40 = null;;
  if ($m_14 instanceof \Data\Either\Data_Either_Left) {
$__t40 = new \Data\Either\Data_Either_Left(($m_14)->{'value0'});
goto end_branch_40;;
};
  if ($m_14 instanceof \Data\Either\Data_Either_Right) {
$__t40 = new \Data\Either\Data_Either_Right(($f_13)(($m_14)->{'value0'}));
goto end_branch_40;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t40 = null;
  end_branch_40:;
  $__res = $__t40;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_15) use ($__local_var_14_40) {
  $__num = \func_num_args();
  $__res = ($__local_var_14_40)($v_15);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_13_43 = (($dictMonad_0)->{'Bind1'})(null);
  $pure_14_44 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_13_43 = (object)["bind" => function($v_15) use ($Bind1_13_43, $pure_14_44) {
  $__num = \func_num_args();
  $__res = function($k_16) use ($Bind1_13_43, $pure_14_44, $v_15) {
  $__num = \func_num_args();
  $__local_var_17_45 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_14_44))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_13_43)->{'bind'})($v_15))(function($v2_18) use ($__local_var_17_45, $k_16) {
  $__num = \func_num_args();
  $__t46 = null;;
  if ($v2_18 instanceof \Data\Either\Data_Either_Left) {
$__t46 = ($__local_var_17_45)(($v2_18)->{'value0'});
goto end_branch_46;;
};
  if ($v2_18 instanceof \Data\Either\Data_Either_Right) {
$__t46 = ($k_16)(($v2_18)->{'value0'});
goto end_branch_46;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t46 = null;
  end_branch_46:;
  $__res = $__t46;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_15) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_16_47 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_16_47 = (object)["map" => function($f_17) use ($__local_var_16_47) {
  $__num = \func_num_args();
  $__local_var_18_48 = (($__local_var_16_47)->{'map'})(function($m_18) use ($f_17) {
  $__num = \func_num_args();
  $__t48 = null;;
  if ($m_18 instanceof \Data\Either\Data_Either_Left) {
$__t48 = new \Data\Either\Data_Either_Left(($m_18)->{'value0'});
goto end_branch_48;;
};
  if ($m_18 instanceof \Data\Either\Data_Either_Right) {
$__t48 = new \Data\Either\Data_Either_Right(($f_17)(($m_18)->{'value0'}));
goto end_branch_48;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t48 = null;
  end_branch_48:;
  $__res = $__t48;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_19) use ($__local_var_18_48) {
  $__num = \func_num_args();
  $__res = ($__local_var_18_48)($v_19);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_17_51 = (($dictMonad_0)->{'Bind1'})(null);
  $pure_18_52 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_17_51 = (object)["bind" => function($v_19) use ($Bind1_17_51, $pure_18_52) {
  $__num = \func_num_args();
  $__res = function($k_20) use ($Bind1_17_51, $pure_18_52, $v_19) {
  $__num = \func_num_args();
  $__local_var_21_53 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_18_52))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_17_51)->{'bind'})($v_19))(function($v2_22) use ($__local_var_21_53, $k_20) {
  $__num = \func_num_args();
  $__t54 = null;;
  if ($v2_22 instanceof \Data\Either\Data_Either_Left) {
$__t54 = ($__local_var_21_53)(($v2_22)->{'value0'});
goto end_branch_54;;
};
  if ($v2_22 instanceof \Data\Either\Data_Either_Right) {
$__t54 = ($k_20)(($v2_22)->{'value0'});
goto end_branch_54;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t54 = null;
  end_branch_54:;
  $__res = $__t54;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_19) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_18_56 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_19) use ($Applicative0_18_56, $Bind1_17_51) {
  $__num = \func_num_args();
  $__res = function($a_20) use ($Applicative0_18_56, $Bind1_17_51, $f_19) {
  $__num = \func_num_args();
  $__res = ((($Bind1_17_51)->{'bind'})($f_19))(function($f_prime__21) use ($Applicative0_18_56, $Bind1_17_51, $a_20) {
  $__num = \func_num_args();
  $__res = ((($Bind1_17_51)->{'bind'})($a_20))(function($a_prime__22) use ($Applicative0_18_56, $f_prime__21) {
  $__num = \func_num_args();
  $__res = (($Applicative0_18_56)->{'pure'})(($f_prime__21)($a_prime__22));
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
}, "Functor0" => function($_dollar___unused_17) use ($functorExceptT1_16_47) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_16_47;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_14_58 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_15) use ($Applicative0_14_58, $Bind1_13_43) {
  $__num = \func_num_args();
  $__res = function($a_16) use ($Applicative0_14_58, $Bind1_13_43, $f_15) {
  $__num = \func_num_args();
  $__res = ((($Bind1_13_43)->{'bind'})($f_15))(function($f_prime__17) use ($Applicative0_14_58, $Bind1_13_43, $a_16) {
  $__num = \func_num_args();
  $__res = ((($Bind1_13_43)->{'bind'})($a_16))(function($a_prime__18) use ($Applicative0_14_58, $f_prime__17) {
  $__num = \func_num_args();
  $__res = (($Applicative0_14_58)->{'pure'})(($f_prime__17)($a_prime__18));
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
}, "Functor0" => function($_dollar___unused_13) use ($functorExceptT1_12_39) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_12_39;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_12) use ($Applicative0_11_39, $Bind1_10_14) {
  $__num = \func_num_args();
  $__res = function($a_13) use ($Applicative0_11_39, $Bind1_10_14, $f_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_14)->{'bind'})($f_12))(function($f_prime__14) use ($Applicative0_11_39, $Bind1_10_14, $a_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_14)->{'bind'})($a_13))(function($a_prime__15) use ($Applicative0_11_39, $f_prime__14) {
  $__num = \func_num_args();
  $__res = (($Applicative0_11_39)->{'pure'})(($f_prime__14)($a_prime__15));
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
}, "Functor0" => function($_dollar___unused_10) use ($functorExceptT1_9_10) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_9_10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_7_61 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_7) {
  $__num = \func_num_args();
  $__res = $x_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_7) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_8_61 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_8_61 = (object)["map" => function($f_9) use ($__local_var_8_61) {
  $__num = \func_num_args();
  $__local_var_10_62 = (($__local_var_8_61)->{'map'})(function($m_10) use ($f_9) {
  $__num = \func_num_args();
  $__t62 = null;;
  if ($m_10 instanceof \Data\Either\Data_Either_Left) {
$__t62 = new \Data\Either\Data_Either_Left(($m_10)->{'value0'});
goto end_branch_62;;
};
  if ($m_10 instanceof \Data\Either\Data_Either_Right) {
$__t62 = new \Data\Either\Data_Either_Right(($f_9)(($m_10)->{'value0'}));
goto end_branch_62;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t62 = null;
  end_branch_62:;
  $__res = $__t62;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_11) use ($__local_var_10_62) {
  $__num = \func_num_args();
  $__res = ($__local_var_10_62)($v_11);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_9_65 = (($dictMonad_0)->{'Bind1'})(null);
  $pure_10_66 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_9_65 = (object)["bind" => function($v_11) use ($Bind1_9_65, $pure_10_66) {
  $__num = \func_num_args();
  $__res = function($k_12) use ($Bind1_9_65, $pure_10_66, $v_11) {
  $__num = \func_num_args();
  $__local_var_13_67 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_10_66))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_9_65)->{'bind'})($v_11))(function($v2_14) use ($__local_var_13_67, $k_12) {
  $__num = \func_num_args();
  $__t68 = null;;
  if ($v2_14 instanceof \Data\Either\Data_Either_Left) {
$__t68 = ($__local_var_13_67)(($v2_14)->{'value0'});
goto end_branch_68;;
};
  if ($v2_14 instanceof \Data\Either\Data_Either_Right) {
$__t68 = ($k_12)(($v2_14)->{'value0'});
goto end_branch_68;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t68 = null;
  end_branch_68:;
  $__res = $__t68;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_11) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_12_69 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_12_69 = (object)["map" => function($f_13) use ($__local_var_12_69) {
  $__num = \func_num_args();
  $__local_var_14_70 = (($__local_var_12_69)->{'map'})(function($m_14) use ($f_13) {
  $__num = \func_num_args();
  $__t70 = null;;
  if ($m_14 instanceof \Data\Either\Data_Either_Left) {
$__t70 = new \Data\Either\Data_Either_Left(($m_14)->{'value0'});
goto end_branch_70;;
};
  if ($m_14 instanceof \Data\Either\Data_Either_Right) {
$__t70 = new \Data\Either\Data_Either_Right(($f_13)(($m_14)->{'value0'}));
goto end_branch_70;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t70 = null;
  end_branch_70:;
  $__res = $__t70;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_15) use ($__local_var_14_70) {
  $__num = \func_num_args();
  $__res = ($__local_var_14_70)($v_15);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_13_73 = (($dictMonad_0)->{'Bind1'})(null);
  $pure_14_74 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_13_73 = (object)["bind" => function($v_15) use ($Bind1_13_73, $pure_14_74) {
  $__num = \func_num_args();
  $__res = function($k_16) use ($Bind1_13_73, $pure_14_74, $v_15) {
  $__num = \func_num_args();
  $__local_var_17_75 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_14_74))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_13_73)->{'bind'})($v_15))(function($v2_18) use ($__local_var_17_75, $k_16) {
  $__num = \func_num_args();
  $__t76 = null;;
  if ($v2_18 instanceof \Data\Either\Data_Either_Left) {
$__t76 = ($__local_var_17_75)(($v2_18)->{'value0'});
goto end_branch_76;;
};
  if ($v2_18 instanceof \Data\Either\Data_Either_Right) {
$__t76 = ($k_16)(($v2_18)->{'value0'});
goto end_branch_76;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t76 = null;
  end_branch_76:;
  $__res = $__t76;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_15) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_14_78 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_15) use ($Applicative0_14_78, $Bind1_13_73) {
  $__num = \func_num_args();
  $__res = function($a_16) use ($Applicative0_14_78, $Bind1_13_73, $f_15) {
  $__num = \func_num_args();
  $__res = ((($Bind1_13_73)->{'bind'})($f_15))(function($f_prime__17) use ($Applicative0_14_78, $Bind1_13_73, $a_16) {
  $__num = \func_num_args();
  $__res = ((($Bind1_13_73)->{'bind'})($a_16))(function($a_prime__18) use ($Applicative0_14_78, $f_prime__17) {
  $__num = \func_num_args();
  $__res = (($Applicative0_14_78)->{'pure'})(($f_prime__17)($a_prime__18));
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
}, "Functor0" => function($_dollar___unused_13) use ($functorExceptT1_12_69) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_12_69;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_10_80 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_11) use ($Applicative0_10_80, $Bind1_9_65) {
  $__num = \func_num_args();
  $__res = function($a_12) use ($Applicative0_10_80, $Bind1_9_65, $f_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_65)->{'bind'})($f_11))(function($f_prime__13) use ($Applicative0_10_80, $Bind1_9_65, $a_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_65)->{'bind'})($a_12))(function($a_prime__14) use ($Applicative0_10_80, $f_prime__13) {
  $__num = \func_num_args();
  $__res = (($Applicative0_10_80)->{'pure'})(($f_prime__13)($a_prime__14));
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
}, "Functor0" => function($_dollar___unused_9) use ($functorExceptT1_8_61) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_8_61;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_8) use ($Applicative0_7_61, $Bind1_6_6) {
  $__num = \func_num_args();
  $__res = function($a_9) use ($Applicative0_7_61, $Bind1_6_6, $f_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_6)->{'bind'})($f_8))(function($f_prime__10) use ($Applicative0_7_61, $Bind1_6_6, $a_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_6)->{'bind'})($a_9))(function($a_prime__11) use ($Applicative0_7_61, $f_prime__10) {
  $__num = \func_num_args();
  $__res = (($Applicative0_7_61)->{'pure'})(($f_prime__10)($a_prime__11));
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
}, "Functor0" => function($_dollar___unused_6) use ($functorExceptT1_5_2) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_5_2;
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
}, "Bind1" => function($_dollar___unused_3) use ($dictMonad_0) {
  $__num = \func_num_args();
  $Bind1_4_82 = (($dictMonad_0)->{'Bind1'})(null);
  $pure_5_83 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $__res = (object)["bind" => function($v_6) use ($Bind1_4_82, $pure_5_83) {
  $__num = \func_num_args();
  $__res = function($k_7) use ($Bind1_4_82, $pure_5_83, $v_6) {
  $__num = \func_num_args();
  $__local_var_8_84 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_5_83))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_4_82)->{'bind'})($v_6))(function($v2_9) use ($__local_var_8_84, $k_7) {
  $__num = \func_num_args();
  $__t85 = null;;
  if ($v2_9 instanceof \Data\Either\Data_Either_Left) {
$__t85 = ($__local_var_8_84)(($v2_9)->{'value0'});
goto end_branch_85;;
};
  if ($v2_9 instanceof \Data\Either\Data_Either_Right) {
$__t85 = ($k_7)(($v2_9)->{'value0'});
goto end_branch_85;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t85 = null;
  end_branch_85:;
  $__res = $__t85;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_6) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_7_86 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_7_86 = (object)["map" => function($f_8) use ($__local_var_7_86) {
  $__num = \func_num_args();
  $__local_var_9_87 = (($__local_var_7_86)->{'map'})(function($m_9) use ($f_8) {
  $__num = \func_num_args();
  $__t87 = null;;
  if ($m_9 instanceof \Data\Either\Data_Either_Left) {
$__t87 = new \Data\Either\Data_Either_Left(($m_9)->{'value0'});
goto end_branch_87;;
};
  if ($m_9 instanceof \Data\Either\Data_Either_Right) {
$__t87 = new \Data\Either\Data_Either_Right(($f_8)(($m_9)->{'value0'}));
goto end_branch_87;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t87 = null;
  end_branch_87:;
  $__res = $__t87;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_10) use ($__local_var_9_87) {
  $__num = \func_num_args();
  $__res = ($__local_var_9_87)($v_10);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_8_90 = (($dictMonad_0)->{'Bind1'})(null);
  $pure_9_91 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_8_90 = (object)["bind" => function($v_10) use ($Bind1_8_90, $pure_9_91) {
  $__num = \func_num_args();
  $__res = function($k_11) use ($Bind1_8_90, $pure_9_91, $v_10) {
  $__num = \func_num_args();
  $__local_var_12_92 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_9_91))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_8_90)->{'bind'})($v_10))(function($v2_13) use ($__local_var_12_92, $k_11) {
  $__num = \func_num_args();
  $__t93 = null;;
  if ($v2_13 instanceof \Data\Either\Data_Either_Left) {
$__t93 = ($__local_var_12_92)(($v2_13)->{'value0'});
goto end_branch_93;;
};
  if ($v2_13 instanceof \Data\Either\Data_Either_Right) {
$__t93 = ($k_11)(($v2_13)->{'value0'});
goto end_branch_93;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t93 = null;
  end_branch_93:;
  $__res = $__t93;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_10) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_11_94 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_11_94 = (object)["map" => function($f_12) use ($__local_var_11_94) {
  $__num = \func_num_args();
  $__local_var_13_95 = (($__local_var_11_94)->{'map'})(function($m_13) use ($f_12) {
  $__num = \func_num_args();
  $__t95 = null;;
  if ($m_13 instanceof \Data\Either\Data_Either_Left) {
$__t95 = new \Data\Either\Data_Either_Left(($m_13)->{'value0'});
goto end_branch_95;;
};
  if ($m_13 instanceof \Data\Either\Data_Either_Right) {
$__t95 = new \Data\Either\Data_Either_Right(($f_12)(($m_13)->{'value0'}));
goto end_branch_95;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t95 = null;
  end_branch_95:;
  $__res = $__t95;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_14) use ($__local_var_13_95) {
  $__num = \func_num_args();
  $__res = ($__local_var_13_95)($v_14);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_12_98 = (($dictMonad_0)->{'Bind1'})(null);
  $pure_13_99 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_12_98 = (object)["bind" => function($v_14) use ($Bind1_12_98, $pure_13_99) {
  $__num = \func_num_args();
  $__res = function($k_15) use ($Bind1_12_98, $pure_13_99, $v_14) {
  $__num = \func_num_args();
  $__local_var_16_100 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_13_99))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_12_98)->{'bind'})($v_14))(function($v2_17) use ($__local_var_16_100, $k_15) {
  $__num = \func_num_args();
  $__t101 = null;;
  if ($v2_17 instanceof \Data\Either\Data_Either_Left) {
$__t101 = ($__local_var_16_100)(($v2_17)->{'value0'});
goto end_branch_101;;
};
  if ($v2_17 instanceof \Data\Either\Data_Either_Right) {
$__t101 = ($k_15)(($v2_17)->{'value0'});
goto end_branch_101;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t101 = null;
  end_branch_101:;
  $__res = $__t101;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_14) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_13_103 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_13) {
  $__num = \func_num_args();
  $__res = $x_13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_13) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_14_103 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_14_103 = (object)["map" => function($f_15) use ($__local_var_14_103) {
  $__num = \func_num_args();
  $__local_var_16_104 = (($__local_var_14_103)->{'map'})(function($m_16) use ($f_15) {
  $__num = \func_num_args();
  $__t104 = null;;
  if ($m_16 instanceof \Data\Either\Data_Either_Left) {
$__t104 = new \Data\Either\Data_Either_Left(($m_16)->{'value0'});
goto end_branch_104;;
};
  if ($m_16 instanceof \Data\Either\Data_Either_Right) {
$__t104 = new \Data\Either\Data_Either_Right(($f_15)(($m_16)->{'value0'}));
goto end_branch_104;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t104 = null;
  end_branch_104:;
  $__res = $__t104;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_17) use ($__local_var_16_104) {
  $__num = \func_num_args();
  $__res = ($__local_var_16_104)($v_17);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_15_107 = (($dictMonad_0)->{'Bind1'})(null);
  $pure_16_108 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_15_107 = (object)["bind" => function($v_17) use ($Bind1_15_107, $pure_16_108) {
  $__num = \func_num_args();
  $__res = function($k_18) use ($Bind1_15_107, $pure_16_108, $v_17) {
  $__num = \func_num_args();
  $__local_var_19_109 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_16_108))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_15_107)->{'bind'})($v_17))(function($v2_20) use ($__local_var_19_109, $k_18) {
  $__num = \func_num_args();
  $__t110 = null;;
  if ($v2_20 instanceof \Data\Either\Data_Either_Left) {
$__t110 = ($__local_var_19_109)(($v2_20)->{'value0'});
goto end_branch_110;;
};
  if ($v2_20 instanceof \Data\Either\Data_Either_Right) {
$__t110 = ($k_18)(($v2_20)->{'value0'});
goto end_branch_110;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t110 = null;
  end_branch_110:;
  $__res = $__t110;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_17) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_16_112 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_17) use ($Applicative0_16_112, $Bind1_15_107) {
  $__num = \func_num_args();
  $__res = function($a_18) use ($Applicative0_16_112, $Bind1_15_107, $f_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_107)->{'bind'})($f_17))(function($f_prime__19) use ($Applicative0_16_112, $Bind1_15_107, $a_18) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_107)->{'bind'})($a_18))(function($a_prime__20) use ($Applicative0_16_112, $f_prime__19) {
  $__num = \func_num_args();
  $__res = (($Applicative0_16_112)->{'pure'})(($f_prime__19)($a_prime__20));
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
}, "Functor0" => function($_dollar___unused_15) use ($functorExceptT1_14_103) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_14_103;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_14) use ($Applicative0_13_103, $Bind1_12_98) {
  $__num = \func_num_args();
  $__res = function($a_15) use ($Applicative0_13_103, $Bind1_12_98, $f_14) {
  $__num = \func_num_args();
  $__res = ((($Bind1_12_98)->{'bind'})($f_14))(function($f_prime__16) use ($Applicative0_13_103, $Bind1_12_98, $a_15) {
  $__num = \func_num_args();
  $__res = ((($Bind1_12_98)->{'bind'})($a_15))(function($a_prime__17) use ($Applicative0_13_103, $f_prime__16) {
  $__num = \func_num_args();
  $__res = (($Applicative0_13_103)->{'pure'})(($f_prime__16)($a_prime__17));
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
}, "Functor0" => function($_dollar___unused_12) use ($functorExceptT1_11_94) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_11_94;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_9_115 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_9) {
  $__num = \func_num_args();
  $__res = $x_9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_9) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_10_115 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_10_115 = (object)["map" => function($f_11) use ($__local_var_10_115) {
  $__num = \func_num_args();
  $__local_var_12_116 = (($__local_var_10_115)->{'map'})(function($m_12) use ($f_11) {
  $__num = \func_num_args();
  $__t116 = null;;
  if ($m_12 instanceof \Data\Either\Data_Either_Left) {
$__t116 = new \Data\Either\Data_Either_Left(($m_12)->{'value0'});
goto end_branch_116;;
};
  if ($m_12 instanceof \Data\Either\Data_Either_Right) {
$__t116 = new \Data\Either\Data_Either_Right(($f_11)(($m_12)->{'value0'}));
goto end_branch_116;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t116 = null;
  end_branch_116:;
  $__res = $__t116;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_13) use ($__local_var_12_116) {
  $__num = \func_num_args();
  $__res = ($__local_var_12_116)($v_13);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_11_119 = (($dictMonad_0)->{'Bind1'})(null);
  $pure_12_120 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_11_119 = (object)["bind" => function($v_13) use ($Bind1_11_119, $pure_12_120) {
  $__num = \func_num_args();
  $__res = function($k_14) use ($Bind1_11_119, $pure_12_120, $v_13) {
  $__num = \func_num_args();
  $__local_var_15_121 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_12_120))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_11_119)->{'bind'})($v_13))(function($v2_16) use ($__local_var_15_121, $k_14) {
  $__num = \func_num_args();
  $__t122 = null;;
  if ($v2_16 instanceof \Data\Either\Data_Either_Left) {
$__t122 = ($__local_var_15_121)(($v2_16)->{'value0'});
goto end_branch_122;;
};
  if ($v2_16 instanceof \Data\Either\Data_Either_Right) {
$__t122 = ($k_14)(($v2_16)->{'value0'});
goto end_branch_122;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t122 = null;
  end_branch_122:;
  $__res = $__t122;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_13) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_14_123 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_14_123 = (object)["map" => function($f_15) use ($__local_var_14_123) {
  $__num = \func_num_args();
  $__local_var_16_124 = (($__local_var_14_123)->{'map'})(function($m_16) use ($f_15) {
  $__num = \func_num_args();
  $__t124 = null;;
  if ($m_16 instanceof \Data\Either\Data_Either_Left) {
$__t124 = new \Data\Either\Data_Either_Left(($m_16)->{'value0'});
goto end_branch_124;;
};
  if ($m_16 instanceof \Data\Either\Data_Either_Right) {
$__t124 = new \Data\Either\Data_Either_Right(($f_15)(($m_16)->{'value0'}));
goto end_branch_124;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t124 = null;
  end_branch_124:;
  $__res = $__t124;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_17) use ($__local_var_16_124) {
  $__num = \func_num_args();
  $__res = ($__local_var_16_124)($v_17);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_15_127 = (($dictMonad_0)->{'Bind1'})(null);
  $pure_16_128 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_15_127 = (object)["bind" => function($v_17) use ($Bind1_15_127, $pure_16_128) {
  $__num = \func_num_args();
  $__res = function($k_18) use ($Bind1_15_127, $pure_16_128, $v_17) {
  $__num = \func_num_args();
  $__local_var_19_129 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_16_128))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_15_127)->{'bind'})($v_17))(function($v2_20) use ($__local_var_19_129, $k_18) {
  $__num = \func_num_args();
  $__t130 = null;;
  if ($v2_20 instanceof \Data\Either\Data_Either_Left) {
$__t130 = ($__local_var_19_129)(($v2_20)->{'value0'});
goto end_branch_130;;
};
  if ($v2_20 instanceof \Data\Either\Data_Either_Right) {
$__t130 = ($k_18)(($v2_20)->{'value0'});
goto end_branch_130;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t130 = null;
  end_branch_130:;
  $__res = $__t130;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_17) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_16_132 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_16) {
  $__num = \func_num_args();
  $__res = $x_16;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_16) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_17_132 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_17_132 = (object)["map" => function($f_18) use ($__local_var_17_132) {
  $__num = \func_num_args();
  $__local_var_19_133 = (($__local_var_17_132)->{'map'})(function($m_19) use ($f_18) {
  $__num = \func_num_args();
  $__t133 = null;;
  if ($m_19 instanceof \Data\Either\Data_Either_Left) {
$__t133 = new \Data\Either\Data_Either_Left(($m_19)->{'value0'});
goto end_branch_133;;
};
  if ($m_19 instanceof \Data\Either\Data_Either_Right) {
$__t133 = new \Data\Either\Data_Either_Right(($f_18)(($m_19)->{'value0'}));
goto end_branch_133;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t133 = null;
  end_branch_133:;
  $__res = $__t133;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_20) use ($__local_var_19_133) {
  $__num = \func_num_args();
  $__res = ($__local_var_19_133)($v_20);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_18_136 = (($dictMonad_0)->{'Bind1'})(null);
  $pure_19_137 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_18_136 = (object)["bind" => function($v_20) use ($Bind1_18_136, $pure_19_137) {
  $__num = \func_num_args();
  $__res = function($k_21) use ($Bind1_18_136, $pure_19_137, $v_20) {
  $__num = \func_num_args();
  $__local_var_22_138 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_19_137))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_18_136)->{'bind'})($v_20))(function($v2_23) use ($__local_var_22_138, $k_21) {
  $__num = \func_num_args();
  $__t139 = null;;
  if ($v2_23 instanceof \Data\Either\Data_Either_Left) {
$__t139 = ($__local_var_22_138)(($v2_23)->{'value0'});
goto end_branch_139;;
};
  if ($v2_23 instanceof \Data\Either\Data_Either_Right) {
$__t139 = ($k_21)(($v2_23)->{'value0'});
goto end_branch_139;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t139 = null;
  end_branch_139:;
  $__res = $__t139;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_20) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_19_141 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_20) use ($Applicative0_19_141, $Bind1_18_136) {
  $__num = \func_num_args();
  $__res = function($a_21) use ($Applicative0_19_141, $Bind1_18_136, $f_20) {
  $__num = \func_num_args();
  $__res = ((($Bind1_18_136)->{'bind'})($f_20))(function($f_prime__22) use ($Applicative0_19_141, $Bind1_18_136, $a_21) {
  $__num = \func_num_args();
  $__res = ((($Bind1_18_136)->{'bind'})($a_21))(function($a_prime__23) use ($Applicative0_19_141, $f_prime__22) {
  $__num = \func_num_args();
  $__res = (($Applicative0_19_141)->{'pure'})(($f_prime__22)($a_prime__23));
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
}, "Functor0" => function($_dollar___unused_18) use ($functorExceptT1_17_132) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_17_132;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_17) use ($Applicative0_16_132, $Bind1_15_127) {
  $__num = \func_num_args();
  $__res = function($a_18) use ($Applicative0_16_132, $Bind1_15_127, $f_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_127)->{'bind'})($f_17))(function($f_prime__19) use ($Applicative0_16_132, $Bind1_15_127, $a_18) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_127)->{'bind'})($a_18))(function($a_prime__20) use ($Applicative0_16_132, $f_prime__19) {
  $__num = \func_num_args();
  $__res = (($Applicative0_16_132)->{'pure'})(($f_prime__19)($a_prime__20));
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
}, "Functor0" => function($_dollar___unused_15) use ($functorExceptT1_14_123) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_14_123;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_12_144 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_12) {
  $__num = \func_num_args();
  $__res = $x_12;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_12) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_13_144 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_13_144 = (object)["map" => function($f_14) use ($__local_var_13_144) {
  $__num = \func_num_args();
  $__local_var_15_145 = (($__local_var_13_144)->{'map'})(function($m_15) use ($f_14) {
  $__num = \func_num_args();
  $__t145 = null;;
  if ($m_15 instanceof \Data\Either\Data_Either_Left) {
$__t145 = new \Data\Either\Data_Either_Left(($m_15)->{'value0'});
goto end_branch_145;;
};
  if ($m_15 instanceof \Data\Either\Data_Either_Right) {
$__t145 = new \Data\Either\Data_Either_Right(($f_14)(($m_15)->{'value0'}));
goto end_branch_145;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t145 = null;
  end_branch_145:;
  $__res = $__t145;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_16) use ($__local_var_15_145) {
  $__num = \func_num_args();
  $__res = ($__local_var_15_145)($v_16);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_14_148 = (($dictMonad_0)->{'Bind1'})(null);
  $pure_15_149 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_14_148 = (object)["bind" => function($v_16) use ($Bind1_14_148, $pure_15_149) {
  $__num = \func_num_args();
  $__res = function($k_17) use ($Bind1_14_148, $pure_15_149, $v_16) {
  $__num = \func_num_args();
  $__local_var_18_150 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_15_149))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_14_148)->{'bind'})($v_16))(function($v2_19) use ($__local_var_18_150, $k_17) {
  $__num = \func_num_args();
  $__t151 = null;;
  if ($v2_19 instanceof \Data\Either\Data_Either_Left) {
$__t151 = ($__local_var_18_150)(($v2_19)->{'value0'});
goto end_branch_151;;
};
  if ($v2_19 instanceof \Data\Either\Data_Either_Right) {
$__t151 = ($k_17)(($v2_19)->{'value0'});
goto end_branch_151;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t151 = null;
  end_branch_151:;
  $__res = $__t151;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_16) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_17_152 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_17_152 = (object)["map" => function($f_18) use ($__local_var_17_152) {
  $__num = \func_num_args();
  $__local_var_19_153 = (($__local_var_17_152)->{'map'})(function($m_19) use ($f_18) {
  $__num = \func_num_args();
  $__t153 = null;;
  if ($m_19 instanceof \Data\Either\Data_Either_Left) {
$__t153 = new \Data\Either\Data_Either_Left(($m_19)->{'value0'});
goto end_branch_153;;
};
  if ($m_19 instanceof \Data\Either\Data_Either_Right) {
$__t153 = new \Data\Either\Data_Either_Right(($f_18)(($m_19)->{'value0'}));
goto end_branch_153;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t153 = null;
  end_branch_153:;
  $__res = $__t153;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_20) use ($__local_var_19_153) {
  $__num = \func_num_args();
  $__res = ($__local_var_19_153)($v_20);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_18_156 = (($dictMonad_0)->{'Bind1'})(null);
  $pure_19_157 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_18_156 = (object)["bind" => function($v_20) use ($Bind1_18_156, $pure_19_157) {
  $__num = \func_num_args();
  $__res = function($k_21) use ($Bind1_18_156, $pure_19_157, $v_20) {
  $__num = \func_num_args();
  $__local_var_22_158 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_19_157))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_18_156)->{'bind'})($v_20))(function($v2_23) use ($__local_var_22_158, $k_21) {
  $__num = \func_num_args();
  $__t159 = null;;
  if ($v2_23 instanceof \Data\Either\Data_Either_Left) {
$__t159 = ($__local_var_22_158)(($v2_23)->{'value0'});
goto end_branch_159;;
};
  if ($v2_23 instanceof \Data\Either\Data_Either_Right) {
$__t159 = ($k_21)(($v2_23)->{'value0'});
goto end_branch_159;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t159 = null;
  end_branch_159:;
  $__res = $__t159;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_20) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_19_161 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_20) use ($Applicative0_19_161, $Bind1_18_156) {
  $__num = \func_num_args();
  $__res = function($a_21) use ($Applicative0_19_161, $Bind1_18_156, $f_20) {
  $__num = \func_num_args();
  $__res = ((($Bind1_18_156)->{'bind'})($f_20))(function($f_prime__22) use ($Applicative0_19_161, $Bind1_18_156, $a_21) {
  $__num = \func_num_args();
  $__res = ((($Bind1_18_156)->{'bind'})($a_21))(function($a_prime__23) use ($Applicative0_19_161, $f_prime__22) {
  $__num = \func_num_args();
  $__res = (($Applicative0_19_161)->{'pure'})(($f_prime__22)($a_prime__23));
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
}, "Functor0" => function($_dollar___unused_18) use ($functorExceptT1_17_152) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_17_152;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_15_163 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_16) use ($Applicative0_15_163, $Bind1_14_148) {
  $__num = \func_num_args();
  $__res = function($a_17) use ($Applicative0_15_163, $Bind1_14_148, $f_16) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_148)->{'bind'})($f_16))(function($f_prime__18) use ($Applicative0_15_163, $Bind1_14_148, $a_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_148)->{'bind'})($a_17))(function($a_prime__19) use ($Applicative0_15_163, $f_prime__18) {
  $__num = \func_num_args();
  $__res = (($Applicative0_15_163)->{'pure'})(($f_prime__18)($a_prime__19));
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
}, "Functor0" => function($_dollar___unused_14) use ($functorExceptT1_13_144) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_13_144;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_13) use ($Applicative0_12_144, $Bind1_11_119) {
  $__num = \func_num_args();
  $__res = function($a_14) use ($Applicative0_12_144, $Bind1_11_119, $f_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_119)->{'bind'})($f_13))(function($f_prime__15) use ($Applicative0_12_144, $Bind1_11_119, $a_14) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_119)->{'bind'})($a_14))(function($a_prime__16) use ($Applicative0_12_144, $f_prime__15) {
  $__num = \func_num_args();
  $__res = (($Applicative0_12_144)->{'pure'})(($f_prime__15)($a_prime__16));
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
}, "Functor0" => function($_dollar___unused_11) use ($functorExceptT1_10_115) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_10_115;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_10) use ($Applicative0_9_115, $Bind1_8_90) {
  $__num = \func_num_args();
  $__res = function($a_11) use ($Applicative0_9_115, $Bind1_8_90, $f_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_8_90)->{'bind'})($f_10))(function($f_prime__12) use ($Applicative0_9_115, $Bind1_8_90, $a_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_8_90)->{'bind'})($a_11))(function($a_prime__13) use ($Applicative0_9_115, $f_prime__12) {
  $__num = \func_num_args();
  $__res = (($Applicative0_9_115)->{'pure'})(($f_prime__12)($a_prime__13));
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
}, "Functor0" => function($_dollar___unused_8) use ($functorExceptT1_7_86) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_7_86;
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
  $monadThrowExceptT1_3_2 = (object)["throwError" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_4) {
  $__num = \func_num_args();
  $__res = $x_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Left'])), "Monad0" => function($_dollar___unused_4) use ($monadExceptT1_3_2) {
  $__num = \func_num_args();
  $__res = $monadExceptT1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["catchError" => function($v_4) use ($Bind1_1_0, $pure_2_1) {
  $__num = \func_num_args();
  $__res = function($k_5) use ($Bind1_1_0, $pure_2_1, $v_4) {
  $__num = \func_num_args();
  $__local_var_6_168 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_2_1))($GLOBALS['Data_Either_Right']);
  $__res = ((($Bind1_1_0)->{'bind'})($v_4))(function($v2_7) use ($__local_var_6_168, $k_5) {
  $__num = \func_num_args();
  $__t169 = null;;
  if ($v2_7 instanceof \Data\Either\Data_Either_Left) {
$__t169 = ($k_5)(($v2_7)->{'value0'});
goto end_branch_169;;
};
  if ($v2_7 instanceof \Data\Either\Data_Either_Right) {
$__t169 = ($__local_var_6_168)(($v2_7)->{'value0'});
goto end_branch_169;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t169 = null;
  end_branch_169:;
  $__res = $__t169;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "MonadThrow0" => function($_dollar___unused_4) use ($monadThrowExceptT1_3_2) {
  $__num = \func_num_args();
  $__res = $monadThrowExceptT1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Except_Trans_monadErrorExceptT'] = __NAMESPACE__ . '\\majControl_majMonad_majExcept_majTrans_monadmajErrormajExceptmajT';

// Control_Monad_Except_Trans_monadSTExceptT
function majControl_majMonad_majExcept_majTrans_monadmajSmajTmajExceptmajT($dictMonadST_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majExcept_majTrans_monadmajSmajTmajExceptmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Monad0_1_0 = (($dictMonadST_0)->{'Monad0'})(null);
  $monadExceptT1_2_1 = (object)["Applicative0" => function($_dollar___unused_2) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_3) {
  $__num = \func_num_args();
  $__res = $x_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($Monad0_1_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_3) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__local_var_4_1 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_4_1 = (object)["map" => function($f_5) use ($__local_var_4_1) {
  $__num = \func_num_args();
  $__local_var_6_2 = (($__local_var_4_1)->{'map'})(function($m_6) use ($f_5) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ($m_6 instanceof \Data\Either\Data_Either_Left) {
$__t2 = new \Data\Either\Data_Either_Left(($m_6)->{'value0'});
goto end_branch_2;;
};
  if ($m_6 instanceof \Data\Either\Data_Either_Right) {
$__t2 = new \Data\Either\Data_Either_Right(($f_5)(($m_6)->{'value0'}));
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_7) use ($__local_var_6_2) {
  $__num = \func_num_args();
  $__res = ($__local_var_6_2)($v_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_5_5 = (($Monad0_1_0)->{'Bind1'})(null);
  $pure_6_6 = ((($Monad0_1_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_5_5 = (object)["bind" => function($v_7) use ($Bind1_5_5, $pure_6_6) {
  $__num = \func_num_args();
  $__res = function($k_8) use ($Bind1_5_5, $pure_6_6, $v_7) {
  $__num = \func_num_args();
  $__local_var_9_7 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_6_6))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_5_5)->{'bind'})($v_7))(function($v2_10) use ($__local_var_9_7, $k_8) {
  $__num = \func_num_args();
  $__t8 = null;;
  if ($v2_10 instanceof \Data\Either\Data_Either_Left) {
$__t8 = ($__local_var_9_7)(($v2_10)->{'value0'});
goto end_branch_8;;
};
  if ($v2_10 instanceof \Data\Either\Data_Either_Right) {
$__t8 = ($k_8)(($v2_10)->{'value0'});
goto end_branch_8;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t8 = null;
  end_branch_8:;
  $__res = $__t8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_7) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__local_var_8_9 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_8_9 = (object)["map" => function($f_9) use ($__local_var_8_9) {
  $__num = \func_num_args();
  $__local_var_10_10 = (($__local_var_8_9)->{'map'})(function($m_10) use ($f_9) {
  $__num = \func_num_args();
  $__t10 = null;;
  if ($m_10 instanceof \Data\Either\Data_Either_Left) {
$__t10 = new \Data\Either\Data_Either_Left(($m_10)->{'value0'});
goto end_branch_10;;
};
  if ($m_10 instanceof \Data\Either\Data_Either_Right) {
$__t10 = new \Data\Either\Data_Either_Right(($f_9)(($m_10)->{'value0'}));
goto end_branch_10;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t10 = null;
  end_branch_10:;
  $__res = $__t10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_11) use ($__local_var_10_10) {
  $__num = \func_num_args();
  $__res = ($__local_var_10_10)($v_11);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_9_13 = (($Monad0_1_0)->{'Bind1'})(null);
  $pure_10_14 = ((($Monad0_1_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_9_13 = (object)["bind" => function($v_11) use ($Bind1_9_13, $pure_10_14) {
  $__num = \func_num_args();
  $__res = function($k_12) use ($Bind1_9_13, $pure_10_14, $v_11) {
  $__num = \func_num_args();
  $__local_var_13_15 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_10_14))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_9_13)->{'bind'})($v_11))(function($v2_14) use ($__local_var_13_15, $k_12) {
  $__num = \func_num_args();
  $__t16 = null;;
  if ($v2_14 instanceof \Data\Either\Data_Either_Left) {
$__t16 = ($__local_var_13_15)(($v2_14)->{'value0'});
goto end_branch_16;;
};
  if ($v2_14 instanceof \Data\Either\Data_Either_Right) {
$__t16 = ($k_12)(($v2_14)->{'value0'});
goto end_branch_16;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t16 = null;
  end_branch_16:;
  $__res = $__t16;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_11) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_10_18 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($Monad0_1_0);
  $__res = (object)["apply" => function($f_11) use ($Applicative0_10_18, $Bind1_9_13) {
  $__num = \func_num_args();
  $__res = function($a_12) use ($Applicative0_10_18, $Bind1_9_13, $f_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_13)->{'bind'})($f_11))(function($f_prime__13) use ($Applicative0_10_18, $Bind1_9_13, $a_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_13)->{'bind'})($a_12))(function($a_prime__14) use ($Applicative0_10_18, $f_prime__13) {
  $__num = \func_num_args();
  $__res = (($Applicative0_10_18)->{'pure'})(($f_prime__13)($a_prime__14));
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
}, "Functor0" => function($_dollar___unused_9) use ($functorExceptT1_8_9) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_8_9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_6_20 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($Monad0_1_0);
  $__res = (object)["apply" => function($f_7) use ($Applicative0_6_20, $Bind1_5_5) {
  $__num = \func_num_args();
  $__res = function($a_8) use ($Applicative0_6_20, $Bind1_5_5, $f_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_5_5)->{'bind'})($f_7))(function($f_prime__9) use ($Applicative0_6_20, $Bind1_5_5, $a_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_5_5)->{'bind'})($a_8))(function($a_prime__10) use ($Applicative0_6_20, $f_prime__9) {
  $__num = \func_num_args();
  $__res = (($Applicative0_6_20)->{'pure'})(($f_prime__9)($a_prime__10));
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
}, "Functor0" => function($_dollar___unused_5) use ($functorExceptT1_4_1) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_4_1;
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
  $Bind1_3_21 = (($Monad0_1_0)->{'Bind1'})(null);
  $pure_4_22 = ((($Monad0_1_0)->{'Applicative0'})(null))->{'pure'};
  $__res = (object)["bind" => function($v_5) use ($Bind1_3_21, $pure_4_22) {
  $__num = \func_num_args();
  $__res = function($k_6) use ($Bind1_3_21, $pure_4_22, $v_5) {
  $__num = \func_num_args();
  $__local_var_7_23 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_4_22))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_3_21)->{'bind'})($v_5))(function($v2_8) use ($__local_var_7_23, $k_6) {
  $__num = \func_num_args();
  $__t24 = null;;
  if ($v2_8 instanceof \Data\Either\Data_Either_Left) {
$__t24 = ($__local_var_7_23)(($v2_8)->{'value0'});
goto end_branch_24;;
};
  if ($v2_8 instanceof \Data\Either\Data_Either_Right) {
$__t24 = ($k_6)(($v2_8)->{'value0'});
goto end_branch_24;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t24 = null;
  end_branch_24:;
  $__res = $__t24;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_5) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__local_var_6_25 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_6_25 = (object)["map" => function($f_7) use ($__local_var_6_25) {
  $__num = \func_num_args();
  $__local_var_8_26 = (($__local_var_6_25)->{'map'})(function($m_8) use ($f_7) {
  $__num = \func_num_args();
  $__t26 = null;;
  if ($m_8 instanceof \Data\Either\Data_Either_Left) {
$__t26 = new \Data\Either\Data_Either_Left(($m_8)->{'value0'});
goto end_branch_26;;
};
  if ($m_8 instanceof \Data\Either\Data_Either_Right) {
$__t26 = new \Data\Either\Data_Either_Right(($f_7)(($m_8)->{'value0'}));
goto end_branch_26;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t26 = null;
  end_branch_26:;
  $__res = $__t26;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_9) use ($__local_var_8_26) {
  $__num = \func_num_args();
  $__res = ($__local_var_8_26)($v_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_7_29 = (($Monad0_1_0)->{'Bind1'})(null);
  $pure_8_30 = ((($Monad0_1_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_7_29 = (object)["bind" => function($v_9) use ($Bind1_7_29, $pure_8_30) {
  $__num = \func_num_args();
  $__res = function($k_10) use ($Bind1_7_29, $pure_8_30, $v_9) {
  $__num = \func_num_args();
  $__local_var_11_31 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_8_30))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_7_29)->{'bind'})($v_9))(function($v2_12) use ($__local_var_11_31, $k_10) {
  $__num = \func_num_args();
  $__t32 = null;;
  if ($v2_12 instanceof \Data\Either\Data_Either_Left) {
$__t32 = ($__local_var_11_31)(($v2_12)->{'value0'});
goto end_branch_32;;
};
  if ($v2_12 instanceof \Data\Either\Data_Either_Right) {
$__t32 = ($k_10)(($v2_12)->{'value0'});
goto end_branch_32;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t32 = null;
  end_branch_32:;
  $__res = $__t32;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_9) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_8_34 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_8) {
  $__num = \func_num_args();
  $__res = $x_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($Monad0_1_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_8) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__local_var_9_34 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_9_34 = (object)["map" => function($f_10) use ($__local_var_9_34) {
  $__num = \func_num_args();
  $__local_var_11_35 = (($__local_var_9_34)->{'map'})(function($m_11) use ($f_10) {
  $__num = \func_num_args();
  $__t35 = null;;
  if ($m_11 instanceof \Data\Either\Data_Either_Left) {
$__t35 = new \Data\Either\Data_Either_Left(($m_11)->{'value0'});
goto end_branch_35;;
};
  if ($m_11 instanceof \Data\Either\Data_Either_Right) {
$__t35 = new \Data\Either\Data_Either_Right(($f_10)(($m_11)->{'value0'}));
goto end_branch_35;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t35 = null;
  end_branch_35:;
  $__res = $__t35;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_12) use ($__local_var_11_35) {
  $__num = \func_num_args();
  $__res = ($__local_var_11_35)($v_12);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_10_38 = (($Monad0_1_0)->{'Bind1'})(null);
  $pure_11_39 = ((($Monad0_1_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_10_38 = (object)["bind" => function($v_12) use ($Bind1_10_38, $pure_11_39) {
  $__num = \func_num_args();
  $__res = function($k_13) use ($Bind1_10_38, $pure_11_39, $v_12) {
  $__num = \func_num_args();
  $__local_var_14_40 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_11_39))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_10_38)->{'bind'})($v_12))(function($v2_15) use ($__local_var_14_40, $k_13) {
  $__num = \func_num_args();
  $__t41 = null;;
  if ($v2_15 instanceof \Data\Either\Data_Either_Left) {
$__t41 = ($__local_var_14_40)(($v2_15)->{'value0'});
goto end_branch_41;;
};
  if ($v2_15 instanceof \Data\Either\Data_Either_Right) {
$__t41 = ($k_13)(($v2_15)->{'value0'});
goto end_branch_41;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t41 = null;
  end_branch_41:;
  $__res = $__t41;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_12) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_11_43 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($Monad0_1_0);
  $__res = (object)["apply" => function($f_12) use ($Applicative0_11_43, $Bind1_10_38) {
  $__num = \func_num_args();
  $__res = function($a_13) use ($Applicative0_11_43, $Bind1_10_38, $f_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_38)->{'bind'})($f_12))(function($f_prime__14) use ($Applicative0_11_43, $Bind1_10_38, $a_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_38)->{'bind'})($a_13))(function($a_prime__15) use ($Applicative0_11_43, $f_prime__14) {
  $__num = \func_num_args();
  $__res = (($Applicative0_11_43)->{'pure'})(($f_prime__14)($a_prime__15));
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
}, "Functor0" => function($_dollar___unused_10) use ($functorExceptT1_9_34) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_9_34;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_9) use ($Applicative0_8_34, $Bind1_7_29) {
  $__num = \func_num_args();
  $__res = function($a_10) use ($Applicative0_8_34, $Bind1_7_29, $f_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_29)->{'bind'})($f_9))(function($f_prime__11) use ($Applicative0_8_34, $Bind1_7_29, $a_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_29)->{'bind'})($a_10))(function($a_prime__12) use ($Applicative0_8_34, $f_prime__11) {
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
}, "Functor0" => function($_dollar___unused_7) use ($functorExceptT1_6_25) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_6_25;
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
  $Bind1_3_46 = (($Monad0_1_0)->{'Bind1'})(null);
  $pure_4_47 = ((($Monad0_1_0)->{'Applicative0'})(null))->{'pure'};
  $__res = (object)["liftST" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($m_5) use ($Bind1_3_46, $pure_4_47) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_46)->{'bind'})($m_5))(function($a_6) use ($pure_4_47) {
  $__num = \func_num_args();
  $__res = ($pure_4_47)(new \Data\Either\Data_Either_Right($a_6));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($dictMonadST_0)->{'liftST'}), "Monad0" => function($_dollar___unused_3) use ($monadExceptT1_2_1) {
  $__num = \func_num_args();
  $__res = $monadExceptT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Except_Trans_monadSTExceptT'] = __NAMESPACE__ . '\\majControl_majMonad_majExcept_majTrans_monadmajSmajTmajExceptmajT';

// Control_Monad_Except_Trans_monoidExceptT
function majControl_majMonad_majExcept_majTrans_monoidmajExceptmajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majExcept_majTrans_monoidmajExceptmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $applicativeExceptT1_1_0 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_1) {
  $__num = \func_num_args();
  $__res = $x_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_1) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_2_0 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_2_0 = (object)["map" => function($f_3) use ($__local_var_2_0) {
  $__num = \func_num_args();
  $__local_var_4_1 = (($__local_var_2_0)->{'map'})(function($m_4) use ($f_3) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($m_4 instanceof \Data\Either\Data_Either_Left) {
$__t1 = new \Data\Either\Data_Either_Left(($m_4)->{'value0'});
goto end_branch_1;;
};
  if ($m_4 instanceof \Data\Either\Data_Either_Right) {
$__t1 = new \Data\Either\Data_Either_Right(($f_3)(($m_4)->{'value0'}));
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_5) use ($__local_var_4_1) {
  $__num = \func_num_args();
  $__res = ($__local_var_4_1)($v_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_3_4 = (($dictMonad_0)->{'Bind1'})(null);
  $pure_4_5 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_3_4 = (object)["bind" => function($v_5) use ($Bind1_3_4, $pure_4_5) {
  $__num = \func_num_args();
  $__res = function($k_6) use ($Bind1_3_4, $pure_4_5, $v_5) {
  $__num = \func_num_args();
  $__local_var_7_6 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_4_5))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_3_4)->{'bind'})($v_5))(function($v2_8) use ($__local_var_7_6, $k_6) {
  $__num = \func_num_args();
  $__t7 = null;;
  if ($v2_8 instanceof \Data\Either\Data_Either_Left) {
$__t7 = ($__local_var_7_6)(($v2_8)->{'value0'});
goto end_branch_7;;
};
  if ($v2_8 instanceof \Data\Either\Data_Either_Right) {
$__t7 = ($k_6)(($v2_8)->{'value0'});
goto end_branch_7;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t7 = null;
  end_branch_7:;
  $__res = $__t7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_5) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_6_8 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_6_8 = (object)["map" => function($f_7) use ($__local_var_6_8) {
  $__num = \func_num_args();
  $__local_var_8_9 = (($__local_var_6_8)->{'map'})(function($m_8) use ($f_7) {
  $__num = \func_num_args();
  $__t9 = null;;
  if ($m_8 instanceof \Data\Either\Data_Either_Left) {
$__t9 = new \Data\Either\Data_Either_Left(($m_8)->{'value0'});
goto end_branch_9;;
};
  if ($m_8 instanceof \Data\Either\Data_Either_Right) {
$__t9 = new \Data\Either\Data_Either_Right(($f_7)(($m_8)->{'value0'}));
goto end_branch_9;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t9 = null;
  end_branch_9:;
  $__res = $__t9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_9) use ($__local_var_8_9) {
  $__num = \func_num_args();
  $__res = ($__local_var_8_9)($v_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_7_12 = (($dictMonad_0)->{'Bind1'})(null);
  $pure_8_13 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_7_12 = (object)["bind" => function($v_9) use ($Bind1_7_12, $pure_8_13) {
  $__num = \func_num_args();
  $__res = function($k_10) use ($Bind1_7_12, $pure_8_13, $v_9) {
  $__num = \func_num_args();
  $__local_var_11_14 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_8_13))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_7_12)->{'bind'})($v_9))(function($v2_12) use ($__local_var_11_14, $k_10) {
  $__num = \func_num_args();
  $__t15 = null;;
  if ($v2_12 instanceof \Data\Either\Data_Either_Left) {
$__t15 = ($__local_var_11_14)(($v2_12)->{'value0'});
goto end_branch_15;;
};
  if ($v2_12 instanceof \Data\Either\Data_Either_Right) {
$__t15 = ($k_10)(($v2_12)->{'value0'});
goto end_branch_15;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t15 = null;
  end_branch_15:;
  $__res = $__t15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_9) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_8_17 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_9) use ($Applicative0_8_17, $Bind1_7_12) {
  $__num = \func_num_args();
  $__res = function($a_10) use ($Applicative0_8_17, $Bind1_7_12, $f_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_12)->{'bind'})($f_9))(function($f_prime__11) use ($Applicative0_8_17, $Bind1_7_12, $a_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_12)->{'bind'})($a_10))(function($a_prime__12) use ($Applicative0_8_17, $f_prime__11) {
  $__num = \func_num_args();
  $__res = (($Applicative0_8_17)->{'pure'})(($f_prime__11)($a_prime__12));
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
}, "Functor0" => function($_dollar___unused_7) use ($functorExceptT1_6_8) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_6_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_4_19 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_5) use ($Applicative0_4_19, $Bind1_3_4) {
  $__num = \func_num_args();
  $__res = function($a_6) use ($Applicative0_4_19, $Bind1_3_4, $f_5) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_4)->{'bind'})($f_5))(function($f_prime__7) use ($Applicative0_4_19, $Bind1_3_4, $a_6) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_4)->{'bind'})($a_6))(function($a_prime__8) use ($Applicative0_4_19, $f_prime__7) {
  $__num = \func_num_args();
  $__res = (($Applicative0_4_19)->{'pure'})(($f_prime__7)($a_prime__8));
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
}, "Functor0" => function($_dollar___unused_3) use ($functorExceptT1_2_0) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_2_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_2_21 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_2_21 = (object)["map" => function($f_3) use ($__local_var_2_21) {
  $__num = \func_num_args();
  $__local_var_4_22 = (($__local_var_2_21)->{'map'})(function($m_4) use ($f_3) {
  $__num = \func_num_args();
  $__t22 = null;;
  if ($m_4 instanceof \Data\Either\Data_Either_Left) {
$__t22 = new \Data\Either\Data_Either_Left(($m_4)->{'value0'});
goto end_branch_22;;
};
  if ($m_4 instanceof \Data\Either\Data_Either_Right) {
$__t22 = new \Data\Either\Data_Either_Right(($f_3)(($m_4)->{'value0'}));
goto end_branch_22;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t22 = null;
  end_branch_22:;
  $__res = $__t22;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_5) use ($__local_var_4_22) {
  $__num = \func_num_args();
  $__res = ($__local_var_4_22)($v_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_3_25 = (($dictMonad_0)->{'Bind1'})(null);
  $pure_4_26 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_3_25 = (object)["bind" => function($v_5) use ($Bind1_3_25, $pure_4_26) {
  $__num = \func_num_args();
  $__res = function($k_6) use ($Bind1_3_25, $pure_4_26, $v_5) {
  $__num = \func_num_args();
  $__local_var_7_27 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_4_26))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_3_25)->{'bind'})($v_5))(function($v2_8) use ($__local_var_7_27, $k_6) {
  $__num = \func_num_args();
  $__t28 = null;;
  if ($v2_8 instanceof \Data\Either\Data_Either_Left) {
$__t28 = ($__local_var_7_27)(($v2_8)->{'value0'});
goto end_branch_28;;
};
  if ($v2_8 instanceof \Data\Either\Data_Either_Right) {
$__t28 = ($k_6)(($v2_8)->{'value0'});
goto end_branch_28;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t28 = null;
  end_branch_28:;
  $__res = $__t28;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_5) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_6_29 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_6_29 = (object)["map" => function($f_7) use ($__local_var_6_29) {
  $__num = \func_num_args();
  $__local_var_8_30 = (($__local_var_6_29)->{'map'})(function($m_8) use ($f_7) {
  $__num = \func_num_args();
  $__t30 = null;;
  if ($m_8 instanceof \Data\Either\Data_Either_Left) {
$__t30 = new \Data\Either\Data_Either_Left(($m_8)->{'value0'});
goto end_branch_30;;
};
  if ($m_8 instanceof \Data\Either\Data_Either_Right) {
$__t30 = new \Data\Either\Data_Either_Right(($f_7)(($m_8)->{'value0'}));
goto end_branch_30;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t30 = null;
  end_branch_30:;
  $__res = $__t30;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_9) use ($__local_var_8_30) {
  $__num = \func_num_args();
  $__res = ($__local_var_8_30)($v_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_7_33 = (($dictMonad_0)->{'Bind1'})(null);
  $pure_8_34 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_7_33 = (object)["bind" => function($v_9) use ($Bind1_7_33, $pure_8_34) {
  $__num = \func_num_args();
  $__res = function($k_10) use ($Bind1_7_33, $pure_8_34, $v_9) {
  $__num = \func_num_args();
  $__local_var_11_35 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_8_34))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_7_33)->{'bind'})($v_9))(function($v2_12) use ($__local_var_11_35, $k_10) {
  $__num = \func_num_args();
  $__t36 = null;;
  if ($v2_12 instanceof \Data\Either\Data_Either_Left) {
$__t36 = ($__local_var_11_35)(($v2_12)->{'value0'});
goto end_branch_36;;
};
  if ($v2_12 instanceof \Data\Either\Data_Either_Right) {
$__t36 = ($k_10)(($v2_12)->{'value0'});
goto end_branch_36;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t36 = null;
  end_branch_36:;
  $__res = $__t36;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_9) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_8_38 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_8) {
  $__num = \func_num_args();
  $__res = $x_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_8) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_9_38 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_9_38 = (object)["map" => function($f_10) use ($__local_var_9_38) {
  $__num = \func_num_args();
  $__local_var_11_39 = (($__local_var_9_38)->{'map'})(function($m_11) use ($f_10) {
  $__num = \func_num_args();
  $__t39 = null;;
  if ($m_11 instanceof \Data\Either\Data_Either_Left) {
$__t39 = new \Data\Either\Data_Either_Left(($m_11)->{'value0'});
goto end_branch_39;;
};
  if ($m_11 instanceof \Data\Either\Data_Either_Right) {
$__t39 = new \Data\Either\Data_Either_Right(($f_10)(($m_11)->{'value0'}));
goto end_branch_39;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t39 = null;
  end_branch_39:;
  $__res = $__t39;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_12) use ($__local_var_11_39) {
  $__num = \func_num_args();
  $__res = ($__local_var_11_39)($v_12);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_10_42 = (($dictMonad_0)->{'Bind1'})(null);
  $pure_11_43 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_10_42 = (object)["bind" => function($v_12) use ($Bind1_10_42, $pure_11_43) {
  $__num = \func_num_args();
  $__res = function($k_13) use ($Bind1_10_42, $pure_11_43, $v_12) {
  $__num = \func_num_args();
  $__local_var_14_44 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_11_43))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_10_42)->{'bind'})($v_12))(function($v2_15) use ($__local_var_14_44, $k_13) {
  $__num = \func_num_args();
  $__t45 = null;;
  if ($v2_15 instanceof \Data\Either\Data_Either_Left) {
$__t45 = ($__local_var_14_44)(($v2_15)->{'value0'});
goto end_branch_45;;
};
  if ($v2_15 instanceof \Data\Either\Data_Either_Right) {
$__t45 = ($k_13)(($v2_15)->{'value0'});
goto end_branch_45;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t45 = null;
  end_branch_45:;
  $__res = $__t45;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_12) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_11_47 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_12) use ($Applicative0_11_47, $Bind1_10_42) {
  $__num = \func_num_args();
  $__res = function($a_13) use ($Applicative0_11_47, $Bind1_10_42, $f_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_42)->{'bind'})($f_12))(function($f_prime__14) use ($Applicative0_11_47, $Bind1_10_42, $a_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_42)->{'bind'})($a_13))(function($a_prime__15) use ($Applicative0_11_47, $f_prime__14) {
  $__num = \func_num_args();
  $__res = (($Applicative0_11_47)->{'pure'})(($f_prime__14)($a_prime__15));
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
}, "Functor0" => function($_dollar___unused_10) use ($functorExceptT1_9_38) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_9_38;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_9) use ($Applicative0_8_38, $Bind1_7_33) {
  $__num = \func_num_args();
  $__res = function($a_10) use ($Applicative0_8_38, $Bind1_7_33, $f_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_33)->{'bind'})($f_9))(function($f_prime__11) use ($Applicative0_8_38, $Bind1_7_33, $a_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_33)->{'bind'})($a_10))(function($a_prime__12) use ($Applicative0_8_38, $f_prime__11) {
  $__num = \func_num_args();
  $__res = (($Applicative0_8_38)->{'pure'})(($f_prime__11)($a_prime__12));
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
}, "Functor0" => function($_dollar___unused_7) use ($functorExceptT1_6_29) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_6_29;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_4_50 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_4) {
  $__num = \func_num_args();
  $__res = $x_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_4) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_5_50 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_5_50 = (object)["map" => function($f_6) use ($__local_var_5_50) {
  $__num = \func_num_args();
  $__local_var_7_51 = (($__local_var_5_50)->{'map'})(function($m_7) use ($f_6) {
  $__num = \func_num_args();
  $__t51 = null;;
  if ($m_7 instanceof \Data\Either\Data_Either_Left) {
$__t51 = new \Data\Either\Data_Either_Left(($m_7)->{'value0'});
goto end_branch_51;;
};
  if ($m_7 instanceof \Data\Either\Data_Either_Right) {
$__t51 = new \Data\Either\Data_Either_Right(($f_6)(($m_7)->{'value0'}));
goto end_branch_51;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t51 = null;
  end_branch_51:;
  $__res = $__t51;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_8) use ($__local_var_7_51) {
  $__num = \func_num_args();
  $__res = ($__local_var_7_51)($v_8);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_6_54 = (($dictMonad_0)->{'Bind1'})(null);
  $pure_7_55 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_6_54 = (object)["bind" => function($v_8) use ($Bind1_6_54, $pure_7_55) {
  $__num = \func_num_args();
  $__res = function($k_9) use ($Bind1_6_54, $pure_7_55, $v_8) {
  $__num = \func_num_args();
  $__local_var_10_56 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_7_55))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_6_54)->{'bind'})($v_8))(function($v2_11) use ($__local_var_10_56, $k_9) {
  $__num = \func_num_args();
  $__t57 = null;;
  if ($v2_11 instanceof \Data\Either\Data_Either_Left) {
$__t57 = ($__local_var_10_56)(($v2_11)->{'value0'});
goto end_branch_57;;
};
  if ($v2_11 instanceof \Data\Either\Data_Either_Right) {
$__t57 = ($k_9)(($v2_11)->{'value0'});
goto end_branch_57;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t57 = null;
  end_branch_57:;
  $__res = $__t57;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_8) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_9_58 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_9_58 = (object)["map" => function($f_10) use ($__local_var_9_58) {
  $__num = \func_num_args();
  $__local_var_11_59 = (($__local_var_9_58)->{'map'})(function($m_11) use ($f_10) {
  $__num = \func_num_args();
  $__t59 = null;;
  if ($m_11 instanceof \Data\Either\Data_Either_Left) {
$__t59 = new \Data\Either\Data_Either_Left(($m_11)->{'value0'});
goto end_branch_59;;
};
  if ($m_11 instanceof \Data\Either\Data_Either_Right) {
$__t59 = new \Data\Either\Data_Either_Right(($f_10)(($m_11)->{'value0'}));
goto end_branch_59;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t59 = null;
  end_branch_59:;
  $__res = $__t59;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_12) use ($__local_var_11_59) {
  $__num = \func_num_args();
  $__res = ($__local_var_11_59)($v_12);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_10_62 = (($dictMonad_0)->{'Bind1'})(null);
  $pure_11_63 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_10_62 = (object)["bind" => function($v_12) use ($Bind1_10_62, $pure_11_63) {
  $__num = \func_num_args();
  $__res = function($k_13) use ($Bind1_10_62, $pure_11_63, $v_12) {
  $__num = \func_num_args();
  $__local_var_14_64 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_11_63))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_10_62)->{'bind'})($v_12))(function($v2_15) use ($__local_var_14_64, $k_13) {
  $__num = \func_num_args();
  $__t65 = null;;
  if ($v2_15 instanceof \Data\Either\Data_Either_Left) {
$__t65 = ($__local_var_14_64)(($v2_15)->{'value0'});
goto end_branch_65;;
};
  if ($v2_15 instanceof \Data\Either\Data_Either_Right) {
$__t65 = ($k_13)(($v2_15)->{'value0'});
goto end_branch_65;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t65 = null;
  end_branch_65:;
  $__res = $__t65;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_12) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_11_67 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_11) {
  $__num = \func_num_args();
  $__res = $x_11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_11) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_12_67 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_12_67 = (object)["map" => function($f_13) use ($__local_var_12_67) {
  $__num = \func_num_args();
  $__local_var_14_68 = (($__local_var_12_67)->{'map'})(function($m_14) use ($f_13) {
  $__num = \func_num_args();
  $__t68 = null;;
  if ($m_14 instanceof \Data\Either\Data_Either_Left) {
$__t68 = new \Data\Either\Data_Either_Left(($m_14)->{'value0'});
goto end_branch_68;;
};
  if ($m_14 instanceof \Data\Either\Data_Either_Right) {
$__t68 = new \Data\Either\Data_Either_Right(($f_13)(($m_14)->{'value0'}));
goto end_branch_68;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t68 = null;
  end_branch_68:;
  $__res = $__t68;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_15) use ($__local_var_14_68) {
  $__num = \func_num_args();
  $__res = ($__local_var_14_68)($v_15);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_13_71 = (($dictMonad_0)->{'Bind1'})(null);
  $pure_14_72 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_13_71 = (object)["bind" => function($v_15) use ($Bind1_13_71, $pure_14_72) {
  $__num = \func_num_args();
  $__res = function($k_16) use ($Bind1_13_71, $pure_14_72, $v_15) {
  $__num = \func_num_args();
  $__local_var_17_73 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_14_72))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_13_71)->{'bind'})($v_15))(function($v2_18) use ($__local_var_17_73, $k_16) {
  $__num = \func_num_args();
  $__t74 = null;;
  if ($v2_18 instanceof \Data\Either\Data_Either_Left) {
$__t74 = ($__local_var_17_73)(($v2_18)->{'value0'});
goto end_branch_74;;
};
  if ($v2_18 instanceof \Data\Either\Data_Either_Right) {
$__t74 = ($k_16)(($v2_18)->{'value0'});
goto end_branch_74;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t74 = null;
  end_branch_74:;
  $__res = $__t74;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_15) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_14_76 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_15) use ($Applicative0_14_76, $Bind1_13_71) {
  $__num = \func_num_args();
  $__res = function($a_16) use ($Applicative0_14_76, $Bind1_13_71, $f_15) {
  $__num = \func_num_args();
  $__res = ((($Bind1_13_71)->{'bind'})($f_15))(function($f_prime__17) use ($Applicative0_14_76, $Bind1_13_71, $a_16) {
  $__num = \func_num_args();
  $__res = ((($Bind1_13_71)->{'bind'})($a_16))(function($a_prime__18) use ($Applicative0_14_76, $f_prime__17) {
  $__num = \func_num_args();
  $__res = (($Applicative0_14_76)->{'pure'})(($f_prime__17)($a_prime__18));
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
}, "Functor0" => function($_dollar___unused_13) use ($functorExceptT1_12_67) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_12_67;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_12) use ($Applicative0_11_67, $Bind1_10_62) {
  $__num = \func_num_args();
  $__res = function($a_13) use ($Applicative0_11_67, $Bind1_10_62, $f_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_62)->{'bind'})($f_12))(function($f_prime__14) use ($Applicative0_11_67, $Bind1_10_62, $a_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_62)->{'bind'})($a_13))(function($a_prime__15) use ($Applicative0_11_67, $f_prime__14) {
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
}, "Functor0" => function($_dollar___unused_10) use ($functorExceptT1_9_58) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_9_58;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_7_79 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_7) {
  $__num = \func_num_args();
  $__res = $x_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_7) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_8_79 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_8_79 = (object)["map" => function($f_9) use ($__local_var_8_79) {
  $__num = \func_num_args();
  $__local_var_10_80 = (($__local_var_8_79)->{'map'})(function($m_10) use ($f_9) {
  $__num = \func_num_args();
  $__t80 = null;;
  if ($m_10 instanceof \Data\Either\Data_Either_Left) {
$__t80 = new \Data\Either\Data_Either_Left(($m_10)->{'value0'});
goto end_branch_80;;
};
  if ($m_10 instanceof \Data\Either\Data_Either_Right) {
$__t80 = new \Data\Either\Data_Either_Right(($f_9)(($m_10)->{'value0'}));
goto end_branch_80;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t80 = null;
  end_branch_80:;
  $__res = $__t80;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_11) use ($__local_var_10_80) {
  $__num = \func_num_args();
  $__res = ($__local_var_10_80)($v_11);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_9_83 = (($dictMonad_0)->{'Bind1'})(null);
  $pure_10_84 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_9_83 = (object)["bind" => function($v_11) use ($Bind1_9_83, $pure_10_84) {
  $__num = \func_num_args();
  $__res = function($k_12) use ($Bind1_9_83, $pure_10_84, $v_11) {
  $__num = \func_num_args();
  $__local_var_13_85 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_10_84))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_9_83)->{'bind'})($v_11))(function($v2_14) use ($__local_var_13_85, $k_12) {
  $__num = \func_num_args();
  $__t86 = null;;
  if ($v2_14 instanceof \Data\Either\Data_Either_Left) {
$__t86 = ($__local_var_13_85)(($v2_14)->{'value0'});
goto end_branch_86;;
};
  if ($v2_14 instanceof \Data\Either\Data_Either_Right) {
$__t86 = ($k_12)(($v2_14)->{'value0'});
goto end_branch_86;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t86 = null;
  end_branch_86:;
  $__res = $__t86;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_11) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_12_87 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_12_87 = (object)["map" => function($f_13) use ($__local_var_12_87) {
  $__num = \func_num_args();
  $__local_var_14_88 = (($__local_var_12_87)->{'map'})(function($m_14) use ($f_13) {
  $__num = \func_num_args();
  $__t88 = null;;
  if ($m_14 instanceof \Data\Either\Data_Either_Left) {
$__t88 = new \Data\Either\Data_Either_Left(($m_14)->{'value0'});
goto end_branch_88;;
};
  if ($m_14 instanceof \Data\Either\Data_Either_Right) {
$__t88 = new \Data\Either\Data_Either_Right(($f_13)(($m_14)->{'value0'}));
goto end_branch_88;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t88 = null;
  end_branch_88:;
  $__res = $__t88;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_15) use ($__local_var_14_88) {
  $__num = \func_num_args();
  $__res = ($__local_var_14_88)($v_15);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_13_91 = (($dictMonad_0)->{'Bind1'})(null);
  $pure_14_92 = ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'};
  $Bind1_13_91 = (object)["bind" => function($v_15) use ($Bind1_13_91, $pure_14_92) {
  $__num = \func_num_args();
  $__res = function($k_16) use ($Bind1_13_91, $pure_14_92, $v_15) {
  $__num = \func_num_args();
  $__local_var_17_93 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_14_92))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_13_91)->{'bind'})($v_15))(function($v2_18) use ($__local_var_17_93, $k_16) {
  $__num = \func_num_args();
  $__t94 = null;;
  if ($v2_18 instanceof \Data\Either\Data_Either_Left) {
$__t94 = ($__local_var_17_93)(($v2_18)->{'value0'});
goto end_branch_94;;
};
  if ($v2_18 instanceof \Data\Either\Data_Either_Right) {
$__t94 = ($k_16)(($v2_18)->{'value0'});
goto end_branch_94;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t94 = null;
  end_branch_94:;
  $__res = $__t94;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_15) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_14_96 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_15) use ($Applicative0_14_96, $Bind1_13_91) {
  $__num = \func_num_args();
  $__res = function($a_16) use ($Applicative0_14_96, $Bind1_13_91, $f_15) {
  $__num = \func_num_args();
  $__res = ((($Bind1_13_91)->{'bind'})($f_15))(function($f_prime__17) use ($Applicative0_14_96, $Bind1_13_91, $a_16) {
  $__num = \func_num_args();
  $__res = ((($Bind1_13_91)->{'bind'})($a_16))(function($a_prime__18) use ($Applicative0_14_96, $f_prime__17) {
  $__num = \func_num_args();
  $__res = (($Applicative0_14_96)->{'pure'})(($f_prime__17)($a_prime__18));
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
}, "Functor0" => function($_dollar___unused_13) use ($functorExceptT1_12_87) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_12_87;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_10_98 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_11) use ($Applicative0_10_98, $Bind1_9_83) {
  $__num = \func_num_args();
  $__res = function($a_12) use ($Applicative0_10_98, $Bind1_9_83, $f_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_83)->{'bind'})($f_11))(function($f_prime__13) use ($Applicative0_10_98, $Bind1_9_83, $a_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_83)->{'bind'})($a_12))(function($a_prime__14) use ($Applicative0_10_98, $f_prime__13) {
  $__num = \func_num_args();
  $__res = (($Applicative0_10_98)->{'pure'})(($f_prime__13)($a_prime__14));
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
}, "Functor0" => function($_dollar___unused_9) use ($functorExceptT1_8_79) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_8_79;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_8) use ($Applicative0_7_79, $Bind1_6_54) {
  $__num = \func_num_args();
  $__res = function($a_9) use ($Applicative0_7_79, $Bind1_6_54, $f_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_54)->{'bind'})($f_8))(function($f_prime__10) use ($Applicative0_7_79, $Bind1_6_54, $a_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_54)->{'bind'})($a_9))(function($a_prime__11) use ($Applicative0_7_79, $f_prime__10) {
  $__num = \func_num_args();
  $__res = (($Applicative0_7_79)->{'pure'})(($f_prime__10)($a_prime__11));
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
}, "Functor0" => function($_dollar___unused_6) use ($functorExceptT1_5_50) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_5_50;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyExceptT1_2_21 = (object)["apply" => function($f_5) use ($Applicative0_4_50, $Bind1_3_25) {
  $__num = \func_num_args();
  $__res = function($a_6) use ($Applicative0_4_50, $Bind1_3_25, $f_5) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_25)->{'bind'})($f_5))(function($f_prime__7) use ($Applicative0_4_50, $Bind1_3_25, $a_6) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_25)->{'bind'})($a_6))(function($a_prime__8) use ($Applicative0_4_50, $f_prime__7) {
  $__num = \func_num_args();
  $__res = (($Applicative0_4_50)->{'pure'})(($f_prime__7)($a_prime__8));
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
}, "Functor0" => function($_dollar___unused_3) use ($functorExceptT1_2_21) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_2_21;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = function($dictMonoid_3) use ($applicativeExceptT1_1_0, $applyExceptT1_2_21) {
  $__num = \func_num_args();
  $Functor0_4_102 = (($applyExceptT1_2_21)->{'Functor0'})(null);
  $__local_var_5_103 = ((($dictMonoid_3)->{'Semigroup0'})(null))->{'append'};
  $semigroupExceptT2_4_102 = (object)["append" => function($a_6) use ($Functor0_4_102, $__local_var_5_103, $applyExceptT1_2_21) {
  $__num = \func_num_args();
  $__res = function($b_7) use ($Functor0_4_102, $__local_var_5_103, $a_6, $applyExceptT1_2_21) {
  $__num = \func_num_args();
  $__res = ((($applyExceptT1_2_21)->{'apply'})(((($Functor0_4_102)->{'map'})($__local_var_5_103))($a_6)))($b_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["mempty" => (($applicativeExceptT1_1_0)->{'pure'})(($dictMonoid_3)->{'mempty'}), "Semigroup0" => function($_dollar___unused_5) use ($semigroupExceptT2_4_102) {
  $__num = \func_num_args();
  $__res = $semigroupExceptT2_4_102;
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
$GLOBALS['Control_Monad_Except_Trans_monoidExceptT'] = __NAMESPACE__ . '\\majControl_majMonad_majExcept_majTrans_monoidmajExceptmajT';

// Control_Monad_Except_Trans_altExceptT
function majControl_majMonad_majExcept_majTrans_altmajExceptmajT($dictSemigroup_0, $dictMonad_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majExcept_majTrans_altmajExceptmajT';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $Bind1_2_0 = (($dictMonad_1)->{'Bind1'})(null);
  $Applicative0_3_1 = (($dictMonad_1)->{'Applicative0'})(null);
  $__local_var_4_2 = (((((($dictMonad_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_4_2 = (object)["map" => function($f_5) use ($__local_var_4_2) {
  $__num = \func_num_args();
  $__local_var_6_3 = (($__local_var_4_2)->{'map'})(function($m_6) use ($f_5) {
  $__num = \func_num_args();
  $__t3 = null;;
  if ($m_6 instanceof \Data\Either\Data_Either_Left) {
$__t3 = new \Data\Either\Data_Either_Left(($m_6)->{'value0'});
goto end_branch_3;;
};
  if ($m_6 instanceof \Data\Either\Data_Either_Right) {
$__t3 = new \Data\Either\Data_Either_Right(($f_5)(($m_6)->{'value0'}));
goto end_branch_3;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t3 = null;
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_7) use ($__local_var_6_3) {
  $__num = \func_num_args();
  $__res = ($__local_var_6_3)($v_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["alt" => function($v_5) use ($Applicative0_3_1, $Bind1_2_0, $dictSemigroup_0) {
  $__num = \func_num_args();
  $__res = function($v1_6) use ($Applicative0_3_1, $Bind1_2_0, $dictSemigroup_0, $v_5) {
  $__num = \func_num_args();
  $__res = ((($Bind1_2_0)->{'bind'})($v_5))(function($rm_7) use ($Applicative0_3_1, $Bind1_2_0, $dictSemigroup_0, $v1_6) {
  $__num = \func_num_args();
  $__t6 = null;;
  if ($rm_7 instanceof \Data\Either\Data_Either_Right) {
$__t6 = (($Applicative0_3_1)->{'pure'})(new \Data\Either\Data_Either_Right(($rm_7)->{'value0'}));
goto end_branch_6;;
};
  if ($rm_7 instanceof \Data\Either\Data_Either_Left) {
$__local_var_8_7 = ($rm_7)->{'value0'};
$__t6 = ((($Bind1_2_0)->{'bind'})($v1_6))(function($rn_9) use ($Applicative0_3_1, $__local_var_8_7, $dictSemigroup_0) {
  $__num = \func_num_args();
  $__t8 = null;;
  if ($rn_9 instanceof \Data\Either\Data_Either_Right) {
$__t8 = (($Applicative0_3_1)->{'pure'})(new \Data\Either\Data_Either_Right(($rn_9)->{'value0'}));
goto end_branch_8;;
};
  if ($rn_9 instanceof \Data\Either\Data_Either_Left) {
$__t8 = (($Applicative0_3_1)->{'pure'})(new \Data\Either\Data_Either_Left(((($dictSemigroup_0)->{'append'})($__local_var_8_7))(($rn_9)->{'value0'})));
goto end_branch_8;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t8 = null;
  end_branch_8:;
  $__res = $__t8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
goto end_branch_6;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t6 = null;
  end_branch_6:;
  $__res = $__t6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_5) use ($functorExceptT1_4_2) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_4_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Control_Monad_Except_Trans_altExceptT'] = __NAMESPACE__ . '\\majControl_majMonad_majExcept_majTrans_altmajExceptmajT';

// Control_Monad_Except_Trans_plusExceptT
function majControl_majMonad_majExcept_majTrans_plusmajExceptmajT($dictMonoid_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majExcept_majTrans_plusmajExceptmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $mempty_1_0 = ($dictMonoid_0)->{'mempty'};
  $__local_var_2_1 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $__res = function($dictMonad_3) use ($__local_var_2_1, $mempty_1_0) {
  $__num = \func_num_args();
  $Bind1_4_2 = (($dictMonad_3)->{'Bind1'})(null);
  $Applicative0_5_3 = (($dictMonad_3)->{'Applicative0'})(null);
  $__local_var_6_4 = (((((($dictMonad_3)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_6_4 = (object)["map" => function($f_7) use ($__local_var_6_4) {
  $__num = \func_num_args();
  $__local_var_8_5 = (($__local_var_6_4)->{'map'})(function($m_8) use ($f_7) {
  $__num = \func_num_args();
  $__t5 = null;;
  if ($m_8 instanceof \Data\Either\Data_Either_Left) {
$__t5 = new \Data\Either\Data_Either_Left(($m_8)->{'value0'});
goto end_branch_5;;
};
  if ($m_8 instanceof \Data\Either\Data_Either_Right) {
$__t5 = new \Data\Either\Data_Either_Right(($f_7)(($m_8)->{'value0'}));
goto end_branch_5;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t5 = null;
  end_branch_5:;
  $__res = $__t5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_9) use ($__local_var_8_5) {
  $__num = \func_num_args();
  $__res = ($__local_var_8_5)($v_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $altExceptT2_4_2 = (object)["alt" => function($v_7) use ($Applicative0_5_3, $Bind1_4_2, $__local_var_2_1) {
  $__num = \func_num_args();
  $__res = function($v1_8) use ($Applicative0_5_3, $Bind1_4_2, $__local_var_2_1, $v_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_4_2)->{'bind'})($v_7))(function($rm_9) use ($Applicative0_5_3, $Bind1_4_2, $__local_var_2_1, $v1_8) {
  $__num = \func_num_args();
  $__t8 = null;;
  if ($rm_9 instanceof \Data\Either\Data_Either_Right) {
$__t8 = (($Applicative0_5_3)->{'pure'})(new \Data\Either\Data_Either_Right(($rm_9)->{'value0'}));
goto end_branch_8;;
};
  if ($rm_9 instanceof \Data\Either\Data_Either_Left) {
$__local_var_10_9 = ($rm_9)->{'value0'};
$__t8 = ((($Bind1_4_2)->{'bind'})($v1_8))(function($rn_11) use ($Applicative0_5_3, $__local_var_10_9, $__local_var_2_1) {
  $__num = \func_num_args();
  $__t10 = null;;
  if ($rn_11 instanceof \Data\Either\Data_Either_Right) {
$__t10 = (($Applicative0_5_3)->{'pure'})(new \Data\Either\Data_Either_Right(($rn_11)->{'value0'}));
goto end_branch_10;;
};
  if ($rn_11 instanceof \Data\Either\Data_Either_Left) {
$__t10 = (($Applicative0_5_3)->{'pure'})(new \Data\Either\Data_Either_Left(((($__local_var_2_1)->{'append'})($__local_var_10_9))(($rn_11)->{'value0'})));
goto end_branch_10;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t10 = null;
  end_branch_10:;
  $__res = $__t10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
goto end_branch_8;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t8 = null;
  end_branch_8:;
  $__res = $__t8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_7) use ($functorExceptT1_6_4) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_6_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["empty" => \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($x_5) {
  $__num = \func_num_args();
  $__res = $x_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_3)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Left']), $mempty_1_0), "Alt0" => function($_dollar___unused_5) use ($altExceptT2_4_2) {
  $__num = \func_num_args();
  $__res = $altExceptT2_4_2;
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
$GLOBALS['Control_Monad_Except_Trans_plusExceptT'] = __NAMESPACE__ . '\\majControl_majMonad_majExcept_majTrans_plusmajExceptmajT';

// Control_Monad_Except_Trans_alternativeExceptT
function majControl_majMonad_majExcept_majTrans_alternativemajExceptmajT($dictMonoid_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majExcept_majTrans_alternativemajExceptmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $mempty_1_0 = ($dictMonoid_0)->{'mempty'};
  $__local_var_2_1 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $__res = function($dictMonad_3) use ($__local_var_2_1, $mempty_1_0) {
  $__num = \func_num_args();
  $applicativeExceptT1_4_2 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_4) {
  $__num = \func_num_args();
  $__res = $x_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_3)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_4) use ($dictMonad_3) {
  $__num = \func_num_args();
  $__local_var_5_2 = (((((($dictMonad_3)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_5_2 = (object)["map" => function($f_6) use ($__local_var_5_2) {
  $__num = \func_num_args();
  $__local_var_7_3 = (($__local_var_5_2)->{'map'})(function($m_7) use ($f_6) {
  $__num = \func_num_args();
  $__t3 = null;;
  if ($m_7 instanceof \Data\Either\Data_Either_Left) {
$__t3 = new \Data\Either\Data_Either_Left(($m_7)->{'value0'});
goto end_branch_3;;
};
  if ($m_7 instanceof \Data\Either\Data_Either_Right) {
$__t3 = new \Data\Either\Data_Either_Right(($f_6)(($m_7)->{'value0'}));
goto end_branch_3;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t3 = null;
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_8) use ($__local_var_7_3) {
  $__num = \func_num_args();
  $__res = ($__local_var_7_3)($v_8);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_6_6 = (($dictMonad_3)->{'Bind1'})(null);
  $pure_7_7 = ((($dictMonad_3)->{'Applicative0'})(null))->{'pure'};
  $Bind1_6_6 = (object)["bind" => function($v_8) use ($Bind1_6_6, $pure_7_7) {
  $__num = \func_num_args();
  $__res = function($k_9) use ($Bind1_6_6, $pure_7_7, $v_8) {
  $__num = \func_num_args();
  $__local_var_10_8 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_7_7))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_6_6)->{'bind'})($v_8))(function($v2_11) use ($__local_var_10_8, $k_9) {
  $__num = \func_num_args();
  $__t9 = null;;
  if ($v2_11 instanceof \Data\Either\Data_Either_Left) {
$__t9 = ($__local_var_10_8)(($v2_11)->{'value0'});
goto end_branch_9;;
};
  if ($v2_11 instanceof \Data\Either\Data_Either_Right) {
$__t9 = ($k_9)(($v2_11)->{'value0'});
goto end_branch_9;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t9 = null;
  end_branch_9:;
  $__res = $__t9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_8) use ($dictMonad_3) {
  $__num = \func_num_args();
  $__local_var_9_10 = (((((($dictMonad_3)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_9_10 = (object)["map" => function($f_10) use ($__local_var_9_10) {
  $__num = \func_num_args();
  $__local_var_11_11 = (($__local_var_9_10)->{'map'})(function($m_11) use ($f_10) {
  $__num = \func_num_args();
  $__t11 = null;;
  if ($m_11 instanceof \Data\Either\Data_Either_Left) {
$__t11 = new \Data\Either\Data_Either_Left(($m_11)->{'value0'});
goto end_branch_11;;
};
  if ($m_11 instanceof \Data\Either\Data_Either_Right) {
$__t11 = new \Data\Either\Data_Either_Right(($f_10)(($m_11)->{'value0'}));
goto end_branch_11;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t11 = null;
  end_branch_11:;
  $__res = $__t11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_12) use ($__local_var_11_11) {
  $__num = \func_num_args();
  $__res = ($__local_var_11_11)($v_12);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_10_14 = (($dictMonad_3)->{'Bind1'})(null);
  $pure_11_15 = ((($dictMonad_3)->{'Applicative0'})(null))->{'pure'};
  $Bind1_10_14 = (object)["bind" => function($v_12) use ($Bind1_10_14, $pure_11_15) {
  $__num = \func_num_args();
  $__res = function($k_13) use ($Bind1_10_14, $pure_11_15, $v_12) {
  $__num = \func_num_args();
  $__local_var_14_16 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_11_15))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_10_14)->{'bind'})($v_12))(function($v2_15) use ($__local_var_14_16, $k_13) {
  $__num = \func_num_args();
  $__t17 = null;;
  if ($v2_15 instanceof \Data\Either\Data_Either_Left) {
$__t17 = ($__local_var_14_16)(($v2_15)->{'value0'});
goto end_branch_17;;
};
  if ($v2_15 instanceof \Data\Either\Data_Either_Right) {
$__t17 = ($k_13)(($v2_15)->{'value0'});
goto end_branch_17;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t17 = null;
  end_branch_17:;
  $__res = $__t17;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_12) use ($dictMonad_3) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($dictMonad_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_11_19 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($dictMonad_3);
  $__res = (object)["apply" => function($f_12) use ($Applicative0_11_19, $Bind1_10_14) {
  $__num = \func_num_args();
  $__res = function($a_13) use ($Applicative0_11_19, $Bind1_10_14, $f_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_14)->{'bind'})($f_12))(function($f_prime__14) use ($Applicative0_11_19, $Bind1_10_14, $a_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_14)->{'bind'})($a_13))(function($a_prime__15) use ($Applicative0_11_19, $f_prime__14) {
  $__num = \func_num_args();
  $__res = (($Applicative0_11_19)->{'pure'})(($f_prime__14)($a_prime__15));
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
}, "Functor0" => function($_dollar___unused_10) use ($functorExceptT1_9_10) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_9_10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_7_21 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($dictMonad_3);
  $__res = (object)["apply" => function($f_8) use ($Applicative0_7_21, $Bind1_6_6) {
  $__num = \func_num_args();
  $__res = function($a_9) use ($Applicative0_7_21, $Bind1_6_6, $f_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_6)->{'bind'})($f_8))(function($f_prime__10) use ($Applicative0_7_21, $Bind1_6_6, $a_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_6)->{'bind'})($a_9))(function($a_prime__11) use ($Applicative0_7_21, $f_prime__10) {
  $__num = \func_num_args();
  $__res = (($Applicative0_7_21)->{'pure'})(($f_prime__10)($a_prime__11));
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
}, "Functor0" => function($_dollar___unused_6) use ($functorExceptT1_5_2) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_5_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_5_23 = (($dictMonad_3)->{'Bind1'})(null);
  $Applicative0_6_24 = (($dictMonad_3)->{'Applicative0'})(null);
  $__local_var_7_25 = (((((($dictMonad_3)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_7_25 = (object)["map" => function($f_8) use ($__local_var_7_25) {
  $__num = \func_num_args();
  $__local_var_9_26 = (($__local_var_7_25)->{'map'})(function($m_9) use ($f_8) {
  $__num = \func_num_args();
  $__t26 = null;;
  if ($m_9 instanceof \Data\Either\Data_Either_Left) {
$__t26 = new \Data\Either\Data_Either_Left(($m_9)->{'value0'});
goto end_branch_26;;
};
  if ($m_9 instanceof \Data\Either\Data_Either_Right) {
$__t26 = new \Data\Either\Data_Either_Right(($f_8)(($m_9)->{'value0'}));
goto end_branch_26;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t26 = null;
  end_branch_26:;
  $__res = $__t26;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_10) use ($__local_var_9_26) {
  $__num = \func_num_args();
  $__res = ($__local_var_9_26)($v_10);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $altExceptT2_5_23 = (object)["alt" => function($v_8) use ($Applicative0_6_24, $Bind1_5_23, $__local_var_2_1) {
  $__num = \func_num_args();
  $__res = function($v1_9) use ($Applicative0_6_24, $Bind1_5_23, $__local_var_2_1, $v_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_5_23)->{'bind'})($v_8))(function($rm_10) use ($Applicative0_6_24, $Bind1_5_23, $__local_var_2_1, $v1_9) {
  $__num = \func_num_args();
  $__t29 = null;;
  if ($rm_10 instanceof \Data\Either\Data_Either_Right) {
$__t29 = (($Applicative0_6_24)->{'pure'})(new \Data\Either\Data_Either_Right(($rm_10)->{'value0'}));
goto end_branch_29;;
};
  if ($rm_10 instanceof \Data\Either\Data_Either_Left) {
$__local_var_11_30 = ($rm_10)->{'value0'};
$__t29 = ((($Bind1_5_23)->{'bind'})($v1_9))(function($rn_12) use ($Applicative0_6_24, $__local_var_11_30, $__local_var_2_1) {
  $__num = \func_num_args();
  $__t31 = null;;
  if ($rn_12 instanceof \Data\Either\Data_Either_Right) {
$__t31 = (($Applicative0_6_24)->{'pure'})(new \Data\Either\Data_Either_Right(($rn_12)->{'value0'}));
goto end_branch_31;;
};
  if ($rn_12 instanceof \Data\Either\Data_Either_Left) {
$__t31 = (($Applicative0_6_24)->{'pure'})(new \Data\Either\Data_Either_Left(((($__local_var_2_1)->{'append'})($__local_var_11_30))(($rn_12)->{'value0'})));
goto end_branch_31;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t31 = null;
  end_branch_31:;
  $__res = $__t31;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
goto end_branch_29;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t29 = null;
  end_branch_29:;
  $__res = $__t29;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_8) use ($functorExceptT1_7_25) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_7_25;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $plusExceptT2_5_23 = (object)["empty" => \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($x_6) {
  $__num = \func_num_args();
  $__res = $x_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_3)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Left']), $mempty_1_0), "Alt0" => function($_dollar___unused_6) use ($altExceptT2_5_23) {
  $__num = \func_num_args();
  $__res = $altExceptT2_5_23;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["Applicative0" => function($_dollar___unused_6) use ($applicativeExceptT1_4_2) {
  $__num = \func_num_args();
  $__res = $applicativeExceptT1_4_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Plus1" => function($_dollar___unused_6) use ($plusExceptT2_5_23) {
  $__num = \func_num_args();
  $__res = $plusExceptT2_5_23;
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
$GLOBALS['Control_Monad_Except_Trans_alternativeExceptT'] = __NAMESPACE__ . '\\majControl_majMonad_majExcept_majTrans_alternativemajExceptmajT';

// Control_Monad_Except_Trans_monadPlusExceptT
function majControl_majMonad_majExcept_majTrans_monadmajPlusmajExceptmajT($dictMonoid_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majExcept_majTrans_monadmajPlusmajExceptmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $mempty_1_0 = ($dictMonoid_0)->{'mempty'};
  $__local_var_2_1 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $__res = function($dictMonad_3) use ($__local_var_2_1, $mempty_1_0) {
  $__num = \func_num_args();
  $monadExceptT1_4_2 = (object)["Applicative0" => function($_dollar___unused_4) use ($dictMonad_3) {
  $__num = \func_num_args();
  $__res = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_5) {
  $__num = \func_num_args();
  $__res = $x_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_3)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_5) use ($dictMonad_3) {
  $__num = \func_num_args();
  $__local_var_6_2 = (((((($dictMonad_3)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_6_2 = (object)["map" => function($f_7) use ($__local_var_6_2) {
  $__num = \func_num_args();
  $__local_var_8_3 = (($__local_var_6_2)->{'map'})(function($m_8) use ($f_7) {
  $__num = \func_num_args();
  $__t3 = null;;
  if ($m_8 instanceof \Data\Either\Data_Either_Left) {
$__t3 = new \Data\Either\Data_Either_Left(($m_8)->{'value0'});
goto end_branch_3;;
};
  if ($m_8 instanceof \Data\Either\Data_Either_Right) {
$__t3 = new \Data\Either\Data_Either_Right(($f_7)(($m_8)->{'value0'}));
goto end_branch_3;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t3 = null;
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_9) use ($__local_var_8_3) {
  $__num = \func_num_args();
  $__res = ($__local_var_8_3)($v_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_7_6 = (($dictMonad_3)->{'Bind1'})(null);
  $pure_8_7 = ((($dictMonad_3)->{'Applicative0'})(null))->{'pure'};
  $Bind1_7_6 = (object)["bind" => function($v_9) use ($Bind1_7_6, $pure_8_7) {
  $__num = \func_num_args();
  $__res = function($k_10) use ($Bind1_7_6, $pure_8_7, $v_9) {
  $__num = \func_num_args();
  $__local_var_11_8 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_8_7))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_7_6)->{'bind'})($v_9))(function($v2_12) use ($__local_var_11_8, $k_10) {
  $__num = \func_num_args();
  $__t9 = null;;
  if ($v2_12 instanceof \Data\Either\Data_Either_Left) {
$__t9 = ($__local_var_11_8)(($v2_12)->{'value0'});
goto end_branch_9;;
};
  if ($v2_12 instanceof \Data\Either\Data_Either_Right) {
$__t9 = ($k_10)(($v2_12)->{'value0'});
goto end_branch_9;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t9 = null;
  end_branch_9:;
  $__res = $__t9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_9) use ($dictMonad_3) {
  $__num = \func_num_args();
  $__local_var_10_10 = (((((($dictMonad_3)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_10_10 = (object)["map" => function($f_11) use ($__local_var_10_10) {
  $__num = \func_num_args();
  $__local_var_12_11 = (($__local_var_10_10)->{'map'})(function($m_12) use ($f_11) {
  $__num = \func_num_args();
  $__t11 = null;;
  if ($m_12 instanceof \Data\Either\Data_Either_Left) {
$__t11 = new \Data\Either\Data_Either_Left(($m_12)->{'value0'});
goto end_branch_11;;
};
  if ($m_12 instanceof \Data\Either\Data_Either_Right) {
$__t11 = new \Data\Either\Data_Either_Right(($f_11)(($m_12)->{'value0'}));
goto end_branch_11;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t11 = null;
  end_branch_11:;
  $__res = $__t11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_13) use ($__local_var_12_11) {
  $__num = \func_num_args();
  $__res = ($__local_var_12_11)($v_13);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_11_14 = (($dictMonad_3)->{'Bind1'})(null);
  $pure_12_15 = ((($dictMonad_3)->{'Applicative0'})(null))->{'pure'};
  $Bind1_11_14 = (object)["bind" => function($v_13) use ($Bind1_11_14, $pure_12_15) {
  $__num = \func_num_args();
  $__res = function($k_14) use ($Bind1_11_14, $pure_12_15, $v_13) {
  $__num = \func_num_args();
  $__local_var_15_16 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_12_15))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_11_14)->{'bind'})($v_13))(function($v2_16) use ($__local_var_15_16, $k_14) {
  $__num = \func_num_args();
  $__t17 = null;;
  if ($v2_16 instanceof \Data\Either\Data_Either_Left) {
$__t17 = ($__local_var_15_16)(($v2_16)->{'value0'});
goto end_branch_17;;
};
  if ($v2_16 instanceof \Data\Either\Data_Either_Right) {
$__t17 = ($k_14)(($v2_16)->{'value0'});
goto end_branch_17;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t17 = null;
  end_branch_17:;
  $__res = $__t17;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_13) use ($dictMonad_3) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($dictMonad_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_12_19 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($dictMonad_3);
  $__res = (object)["apply" => function($f_13) use ($Applicative0_12_19, $Bind1_11_14) {
  $__num = \func_num_args();
  $__res = function($a_14) use ($Applicative0_12_19, $Bind1_11_14, $f_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_14)->{'bind'})($f_13))(function($f_prime__15) use ($Applicative0_12_19, $Bind1_11_14, $a_14) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_14)->{'bind'})($a_14))(function($a_prime__16) use ($Applicative0_12_19, $f_prime__15) {
  $__num = \func_num_args();
  $__res = (($Applicative0_12_19)->{'pure'})(($f_prime__15)($a_prime__16));
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
}, "Functor0" => function($_dollar___unused_11) use ($functorExceptT1_10_10) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_10_10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_8_21 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($dictMonad_3);
  $__res = (object)["apply" => function($f_9) use ($Applicative0_8_21, $Bind1_7_6) {
  $__num = \func_num_args();
  $__res = function($a_10) use ($Applicative0_8_21, $Bind1_7_6, $f_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_6)->{'bind'})($f_9))(function($f_prime__11) use ($Applicative0_8_21, $Bind1_7_6, $a_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_6)->{'bind'})($a_10))(function($a_prime__12) use ($Applicative0_8_21, $f_prime__11) {
  $__num = \func_num_args();
  $__res = (($Applicative0_8_21)->{'pure'})(($f_prime__11)($a_prime__12));
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
}, "Functor0" => function($_dollar___unused_7) use ($functorExceptT1_6_2) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_6_2;
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
}, "Bind1" => function($_dollar___unused_4) use ($dictMonad_3) {
  $__num = \func_num_args();
  $Bind1_5_22 = (($dictMonad_3)->{'Bind1'})(null);
  $pure_6_23 = ((($dictMonad_3)->{'Applicative0'})(null))->{'pure'};
  $__res = (object)["bind" => function($v_7) use ($Bind1_5_22, $pure_6_23) {
  $__num = \func_num_args();
  $__res = function($k_8) use ($Bind1_5_22, $pure_6_23, $v_7) {
  $__num = \func_num_args();
  $__local_var_9_24 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_6_23))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_5_22)->{'bind'})($v_7))(function($v2_10) use ($__local_var_9_24, $k_8) {
  $__num = \func_num_args();
  $__t25 = null;;
  if ($v2_10 instanceof \Data\Either\Data_Either_Left) {
$__t25 = ($__local_var_9_24)(($v2_10)->{'value0'});
goto end_branch_25;;
};
  if ($v2_10 instanceof \Data\Either\Data_Either_Right) {
$__t25 = ($k_8)(($v2_10)->{'value0'});
goto end_branch_25;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t25 = null;
  end_branch_25:;
  $__res = $__t25;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_7) use ($dictMonad_3) {
  $__num = \func_num_args();
  $__local_var_8_26 = (((((($dictMonad_3)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_8_26 = (object)["map" => function($f_9) use ($__local_var_8_26) {
  $__num = \func_num_args();
  $__local_var_10_27 = (($__local_var_8_26)->{'map'})(function($m_10) use ($f_9) {
  $__num = \func_num_args();
  $__t27 = null;;
  if ($m_10 instanceof \Data\Either\Data_Either_Left) {
$__t27 = new \Data\Either\Data_Either_Left(($m_10)->{'value0'});
goto end_branch_27;;
};
  if ($m_10 instanceof \Data\Either\Data_Either_Right) {
$__t27 = new \Data\Either\Data_Either_Right(($f_9)(($m_10)->{'value0'}));
goto end_branch_27;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t27 = null;
  end_branch_27:;
  $__res = $__t27;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_11) use ($__local_var_10_27) {
  $__num = \func_num_args();
  $__res = ($__local_var_10_27)($v_11);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_9_30 = (($dictMonad_3)->{'Bind1'})(null);
  $pure_10_31 = ((($dictMonad_3)->{'Applicative0'})(null))->{'pure'};
  $Bind1_9_30 = (object)["bind" => function($v_11) use ($Bind1_9_30, $pure_10_31) {
  $__num = \func_num_args();
  $__res = function($k_12) use ($Bind1_9_30, $pure_10_31, $v_11) {
  $__num = \func_num_args();
  $__local_var_13_32 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_10_31))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_9_30)->{'bind'})($v_11))(function($v2_14) use ($__local_var_13_32, $k_12) {
  $__num = \func_num_args();
  $__t33 = null;;
  if ($v2_14 instanceof \Data\Either\Data_Either_Left) {
$__t33 = ($__local_var_13_32)(($v2_14)->{'value0'});
goto end_branch_33;;
};
  if ($v2_14 instanceof \Data\Either\Data_Either_Right) {
$__t33 = ($k_12)(($v2_14)->{'value0'});
goto end_branch_33;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t33 = null;
  end_branch_33:;
  $__res = $__t33;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_11) use ($dictMonad_3) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($dictMonad_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_10_35 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_10) {
  $__num = \func_num_args();
  $__res = $x_10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_3)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_10) use ($dictMonad_3) {
  $__num = \func_num_args();
  $__local_var_11_35 = (((((($dictMonad_3)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_11_35 = (object)["map" => function($f_12) use ($__local_var_11_35) {
  $__num = \func_num_args();
  $__local_var_13_36 = (($__local_var_11_35)->{'map'})(function($m_13) use ($f_12) {
  $__num = \func_num_args();
  $__t36 = null;;
  if ($m_13 instanceof \Data\Either\Data_Either_Left) {
$__t36 = new \Data\Either\Data_Either_Left(($m_13)->{'value0'});
goto end_branch_36;;
};
  if ($m_13 instanceof \Data\Either\Data_Either_Right) {
$__t36 = new \Data\Either\Data_Either_Right(($f_12)(($m_13)->{'value0'}));
goto end_branch_36;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t36 = null;
  end_branch_36:;
  $__res = $__t36;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_14) use ($__local_var_13_36) {
  $__num = \func_num_args();
  $__res = ($__local_var_13_36)($v_14);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_12_39 = (($dictMonad_3)->{'Bind1'})(null);
  $pure_13_40 = ((($dictMonad_3)->{'Applicative0'})(null))->{'pure'};
  $Bind1_12_39 = (object)["bind" => function($v_14) use ($Bind1_12_39, $pure_13_40) {
  $__num = \func_num_args();
  $__res = function($k_15) use ($Bind1_12_39, $pure_13_40, $v_14) {
  $__num = \func_num_args();
  $__local_var_16_41 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_13_40))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_12_39)->{'bind'})($v_14))(function($v2_17) use ($__local_var_16_41, $k_15) {
  $__num = \func_num_args();
  $__t42 = null;;
  if ($v2_17 instanceof \Data\Either\Data_Either_Left) {
$__t42 = ($__local_var_16_41)(($v2_17)->{'value0'});
goto end_branch_42;;
};
  if ($v2_17 instanceof \Data\Either\Data_Either_Right) {
$__t42 = ($k_15)(($v2_17)->{'value0'});
goto end_branch_42;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t42 = null;
  end_branch_42:;
  $__res = $__t42;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_14) use ($dictMonad_3) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($dictMonad_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_13_44 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($dictMonad_3);
  $__res = (object)["apply" => function($f_14) use ($Applicative0_13_44, $Bind1_12_39) {
  $__num = \func_num_args();
  $__res = function($a_15) use ($Applicative0_13_44, $Bind1_12_39, $f_14) {
  $__num = \func_num_args();
  $__res = ((($Bind1_12_39)->{'bind'})($f_14))(function($f_prime__16) use ($Applicative0_13_44, $Bind1_12_39, $a_15) {
  $__num = \func_num_args();
  $__res = ((($Bind1_12_39)->{'bind'})($a_15))(function($a_prime__17) use ($Applicative0_13_44, $f_prime__16) {
  $__num = \func_num_args();
  $__res = (($Applicative0_13_44)->{'pure'})(($f_prime__16)($a_prime__17));
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
}, "Functor0" => function($_dollar___unused_12) use ($functorExceptT1_11_35) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_11_35;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_11) use ($Applicative0_10_35, $Bind1_9_30) {
  $__num = \func_num_args();
  $__res = function($a_12) use ($Applicative0_10_35, $Bind1_9_30, $f_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_30)->{'bind'})($f_11))(function($f_prime__13) use ($Applicative0_10_35, $Bind1_9_30, $a_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_30)->{'bind'})($a_12))(function($a_prime__14) use ($Applicative0_10_35, $f_prime__13) {
  $__num = \func_num_args();
  $__res = (($Applicative0_10_35)->{'pure'})(($f_prime__13)($a_prime__14));
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
}, "Functor0" => function($_dollar___unused_9) use ($functorExceptT1_8_26) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_8_26;
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
  $applicativeExceptT1_5_47 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_5) {
  $__num = \func_num_args();
  $__res = $x_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_3)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_5) use ($dictMonad_3) {
  $__num = \func_num_args();
  $__local_var_6_47 = (((((($dictMonad_3)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_6_47 = (object)["map" => function($f_7) use ($__local_var_6_47) {
  $__num = \func_num_args();
  $__local_var_8_48 = (($__local_var_6_47)->{'map'})(function($m_8) use ($f_7) {
  $__num = \func_num_args();
  $__t48 = null;;
  if ($m_8 instanceof \Data\Either\Data_Either_Left) {
$__t48 = new \Data\Either\Data_Either_Left(($m_8)->{'value0'});
goto end_branch_48;;
};
  if ($m_8 instanceof \Data\Either\Data_Either_Right) {
$__t48 = new \Data\Either\Data_Either_Right(($f_7)(($m_8)->{'value0'}));
goto end_branch_48;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t48 = null;
  end_branch_48:;
  $__res = $__t48;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_9) use ($__local_var_8_48) {
  $__num = \func_num_args();
  $__res = ($__local_var_8_48)($v_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_7_51 = (($dictMonad_3)->{'Bind1'})(null);
  $pure_8_52 = ((($dictMonad_3)->{'Applicative0'})(null))->{'pure'};
  $Bind1_7_51 = (object)["bind" => function($v_9) use ($Bind1_7_51, $pure_8_52) {
  $__num = \func_num_args();
  $__res = function($k_10) use ($Bind1_7_51, $pure_8_52, $v_9) {
  $__num = \func_num_args();
  $__local_var_11_53 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_8_52))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_7_51)->{'bind'})($v_9))(function($v2_12) use ($__local_var_11_53, $k_10) {
  $__num = \func_num_args();
  $__t54 = null;;
  if ($v2_12 instanceof \Data\Either\Data_Either_Left) {
$__t54 = ($__local_var_11_53)(($v2_12)->{'value0'});
goto end_branch_54;;
};
  if ($v2_12 instanceof \Data\Either\Data_Either_Right) {
$__t54 = ($k_10)(($v2_12)->{'value0'});
goto end_branch_54;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t54 = null;
  end_branch_54:;
  $__res = $__t54;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_9) use ($dictMonad_3) {
  $__num = \func_num_args();
  $__local_var_10_55 = (((((($dictMonad_3)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_10_55 = (object)["map" => function($f_11) use ($__local_var_10_55) {
  $__num = \func_num_args();
  $__local_var_12_56 = (($__local_var_10_55)->{'map'})(function($m_12) use ($f_11) {
  $__num = \func_num_args();
  $__t56 = null;;
  if ($m_12 instanceof \Data\Either\Data_Either_Left) {
$__t56 = new \Data\Either\Data_Either_Left(($m_12)->{'value0'});
goto end_branch_56;;
};
  if ($m_12 instanceof \Data\Either\Data_Either_Right) {
$__t56 = new \Data\Either\Data_Either_Right(($f_11)(($m_12)->{'value0'}));
goto end_branch_56;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t56 = null;
  end_branch_56:;
  $__res = $__t56;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_13) use ($__local_var_12_56) {
  $__num = \func_num_args();
  $__res = ($__local_var_12_56)($v_13);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_11_59 = (($dictMonad_3)->{'Bind1'})(null);
  $pure_12_60 = ((($dictMonad_3)->{'Applicative0'})(null))->{'pure'};
  $Bind1_11_59 = (object)["bind" => function($v_13) use ($Bind1_11_59, $pure_12_60) {
  $__num = \func_num_args();
  $__res = function($k_14) use ($Bind1_11_59, $pure_12_60, $v_13) {
  $__num = \func_num_args();
  $__local_var_15_61 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_12_60))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_11_59)->{'bind'})($v_13))(function($v2_16) use ($__local_var_15_61, $k_14) {
  $__num = \func_num_args();
  $__t62 = null;;
  if ($v2_16 instanceof \Data\Either\Data_Either_Left) {
$__t62 = ($__local_var_15_61)(($v2_16)->{'value0'});
goto end_branch_62;;
};
  if ($v2_16 instanceof \Data\Either\Data_Either_Right) {
$__t62 = ($k_14)(($v2_16)->{'value0'});
goto end_branch_62;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t62 = null;
  end_branch_62:;
  $__res = $__t62;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_13) use ($dictMonad_3) {
  $__num = \func_num_args();
  $__local_var_14_63 = (((((($dictMonad_3)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_14_63 = (object)["map" => function($f_15) use ($__local_var_14_63) {
  $__num = \func_num_args();
  $__local_var_16_64 = (($__local_var_14_63)->{'map'})(function($m_16) use ($f_15) {
  $__num = \func_num_args();
  $__t64 = null;;
  if ($m_16 instanceof \Data\Either\Data_Either_Left) {
$__t64 = new \Data\Either\Data_Either_Left(($m_16)->{'value0'});
goto end_branch_64;;
};
  if ($m_16 instanceof \Data\Either\Data_Either_Right) {
$__t64 = new \Data\Either\Data_Either_Right(($f_15)(($m_16)->{'value0'}));
goto end_branch_64;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t64 = null;
  end_branch_64:;
  $__res = $__t64;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_17) use ($__local_var_16_64) {
  $__num = \func_num_args();
  $__res = ($__local_var_16_64)($v_17);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_15_67 = (($dictMonad_3)->{'Bind1'})(null);
  $pure_16_68 = ((($dictMonad_3)->{'Applicative0'})(null))->{'pure'};
  $Bind1_15_67 = (object)["bind" => function($v_17) use ($Bind1_15_67, $pure_16_68) {
  $__num = \func_num_args();
  $__res = function($k_18) use ($Bind1_15_67, $pure_16_68, $v_17) {
  $__num = \func_num_args();
  $__local_var_19_69 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_16_68))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_15_67)->{'bind'})($v_17))(function($v2_20) use ($__local_var_19_69, $k_18) {
  $__num = \func_num_args();
  $__t70 = null;;
  if ($v2_20 instanceof \Data\Either\Data_Either_Left) {
$__t70 = ($__local_var_19_69)(($v2_20)->{'value0'});
goto end_branch_70;;
};
  if ($v2_20 instanceof \Data\Either\Data_Either_Right) {
$__t70 = ($k_18)(($v2_20)->{'value0'});
goto end_branch_70;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t70 = null;
  end_branch_70:;
  $__res = $__t70;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_17) use ($dictMonad_3) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($dictMonad_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_16_72 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_16) {
  $__num = \func_num_args();
  $__res = $x_16;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_3)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_16) use ($dictMonad_3) {
  $__num = \func_num_args();
  $__local_var_17_72 = (((((($dictMonad_3)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_17_72 = (object)["map" => function($f_18) use ($__local_var_17_72) {
  $__num = \func_num_args();
  $__local_var_19_73 = (($__local_var_17_72)->{'map'})(function($m_19) use ($f_18) {
  $__num = \func_num_args();
  $__t73 = null;;
  if ($m_19 instanceof \Data\Either\Data_Either_Left) {
$__t73 = new \Data\Either\Data_Either_Left(($m_19)->{'value0'});
goto end_branch_73;;
};
  if ($m_19 instanceof \Data\Either\Data_Either_Right) {
$__t73 = new \Data\Either\Data_Either_Right(($f_18)(($m_19)->{'value0'}));
goto end_branch_73;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t73 = null;
  end_branch_73:;
  $__res = $__t73;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_20) use ($__local_var_19_73) {
  $__num = \func_num_args();
  $__res = ($__local_var_19_73)($v_20);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_18_76 = (($dictMonad_3)->{'Bind1'})(null);
  $pure_19_77 = ((($dictMonad_3)->{'Applicative0'})(null))->{'pure'};
  $Bind1_18_76 = (object)["bind" => function($v_20) use ($Bind1_18_76, $pure_19_77) {
  $__num = \func_num_args();
  $__res = function($k_21) use ($Bind1_18_76, $pure_19_77, $v_20) {
  $__num = \func_num_args();
  $__local_var_22_78 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_19_77))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_18_76)->{'bind'})($v_20))(function($v2_23) use ($__local_var_22_78, $k_21) {
  $__num = \func_num_args();
  $__t79 = null;;
  if ($v2_23 instanceof \Data\Either\Data_Either_Left) {
$__t79 = ($__local_var_22_78)(($v2_23)->{'value0'});
goto end_branch_79;;
};
  if ($v2_23 instanceof \Data\Either\Data_Either_Right) {
$__t79 = ($k_21)(($v2_23)->{'value0'});
goto end_branch_79;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t79 = null;
  end_branch_79:;
  $__res = $__t79;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_20) use ($dictMonad_3) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($dictMonad_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_19_81 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($dictMonad_3);
  $__res = (object)["apply" => function($f_20) use ($Applicative0_19_81, $Bind1_18_76) {
  $__num = \func_num_args();
  $__res = function($a_21) use ($Applicative0_19_81, $Bind1_18_76, $f_20) {
  $__num = \func_num_args();
  $__res = ((($Bind1_18_76)->{'bind'})($f_20))(function($f_prime__22) use ($Applicative0_19_81, $Bind1_18_76, $a_21) {
  $__num = \func_num_args();
  $__res = ((($Bind1_18_76)->{'bind'})($a_21))(function($a_prime__23) use ($Applicative0_19_81, $f_prime__22) {
  $__num = \func_num_args();
  $__res = (($Applicative0_19_81)->{'pure'})(($f_prime__22)($a_prime__23));
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
}, "Functor0" => function($_dollar___unused_18) use ($functorExceptT1_17_72) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_17_72;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_17) use ($Applicative0_16_72, $Bind1_15_67) {
  $__num = \func_num_args();
  $__res = function($a_18) use ($Applicative0_16_72, $Bind1_15_67, $f_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_67)->{'bind'})($f_17))(function($f_prime__19) use ($Applicative0_16_72, $Bind1_15_67, $a_18) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_67)->{'bind'})($a_18))(function($a_prime__20) use ($Applicative0_16_72, $f_prime__19) {
  $__num = \func_num_args();
  $__res = (($Applicative0_16_72)->{'pure'})(($f_prime__19)($a_prime__20));
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
}, "Functor0" => function($_dollar___unused_15) use ($functorExceptT1_14_63) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_14_63;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_12_84 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_12) {
  $__num = \func_num_args();
  $__res = $x_12;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_3)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_12) use ($dictMonad_3) {
  $__num = \func_num_args();
  $__local_var_13_84 = (((((($dictMonad_3)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_13_84 = (object)["map" => function($f_14) use ($__local_var_13_84) {
  $__num = \func_num_args();
  $__local_var_15_85 = (($__local_var_13_84)->{'map'})(function($m_15) use ($f_14) {
  $__num = \func_num_args();
  $__t85 = null;;
  if ($m_15 instanceof \Data\Either\Data_Either_Left) {
$__t85 = new \Data\Either\Data_Either_Left(($m_15)->{'value0'});
goto end_branch_85;;
};
  if ($m_15 instanceof \Data\Either\Data_Either_Right) {
$__t85 = new \Data\Either\Data_Either_Right(($f_14)(($m_15)->{'value0'}));
goto end_branch_85;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t85 = null;
  end_branch_85:;
  $__res = $__t85;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_16) use ($__local_var_15_85) {
  $__num = \func_num_args();
  $__res = ($__local_var_15_85)($v_16);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_14_88 = (($dictMonad_3)->{'Bind1'})(null);
  $pure_15_89 = ((($dictMonad_3)->{'Applicative0'})(null))->{'pure'};
  $Bind1_14_88 = (object)["bind" => function($v_16) use ($Bind1_14_88, $pure_15_89) {
  $__num = \func_num_args();
  $__res = function($k_17) use ($Bind1_14_88, $pure_15_89, $v_16) {
  $__num = \func_num_args();
  $__local_var_18_90 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_15_89))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_14_88)->{'bind'})($v_16))(function($v2_19) use ($__local_var_18_90, $k_17) {
  $__num = \func_num_args();
  $__t91 = null;;
  if ($v2_19 instanceof \Data\Either\Data_Either_Left) {
$__t91 = ($__local_var_18_90)(($v2_19)->{'value0'});
goto end_branch_91;;
};
  if ($v2_19 instanceof \Data\Either\Data_Either_Right) {
$__t91 = ($k_17)(($v2_19)->{'value0'});
goto end_branch_91;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t91 = null;
  end_branch_91:;
  $__res = $__t91;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_16) use ($dictMonad_3) {
  $__num = \func_num_args();
  $__local_var_17_92 = (((((($dictMonad_3)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_17_92 = (object)["map" => function($f_18) use ($__local_var_17_92) {
  $__num = \func_num_args();
  $__local_var_19_93 = (($__local_var_17_92)->{'map'})(function($m_19) use ($f_18) {
  $__num = \func_num_args();
  $__t93 = null;;
  if ($m_19 instanceof \Data\Either\Data_Either_Left) {
$__t93 = new \Data\Either\Data_Either_Left(($m_19)->{'value0'});
goto end_branch_93;;
};
  if ($m_19 instanceof \Data\Either\Data_Either_Right) {
$__t93 = new \Data\Either\Data_Either_Right(($f_18)(($m_19)->{'value0'}));
goto end_branch_93;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t93 = null;
  end_branch_93:;
  $__res = $__t93;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_20) use ($__local_var_19_93) {
  $__num = \func_num_args();
  $__res = ($__local_var_19_93)($v_20);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_18_96 = (($dictMonad_3)->{'Bind1'})(null);
  $pure_19_97 = ((($dictMonad_3)->{'Applicative0'})(null))->{'pure'};
  $Bind1_18_96 = (object)["bind" => function($v_20) use ($Bind1_18_96, $pure_19_97) {
  $__num = \func_num_args();
  $__res = function($k_21) use ($Bind1_18_96, $pure_19_97, $v_20) {
  $__num = \func_num_args();
  $__local_var_22_98 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_19_97))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_18_96)->{'bind'})($v_20))(function($v2_23) use ($__local_var_22_98, $k_21) {
  $__num = \func_num_args();
  $__t99 = null;;
  if ($v2_23 instanceof \Data\Either\Data_Either_Left) {
$__t99 = ($__local_var_22_98)(($v2_23)->{'value0'});
goto end_branch_99;;
};
  if ($v2_23 instanceof \Data\Either\Data_Either_Right) {
$__t99 = ($k_21)(($v2_23)->{'value0'});
goto end_branch_99;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t99 = null;
  end_branch_99:;
  $__res = $__t99;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_20) use ($dictMonad_3) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($dictMonad_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_19_101 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($dictMonad_3);
  $__res = (object)["apply" => function($f_20) use ($Applicative0_19_101, $Bind1_18_96) {
  $__num = \func_num_args();
  $__res = function($a_21) use ($Applicative0_19_101, $Bind1_18_96, $f_20) {
  $__num = \func_num_args();
  $__res = ((($Bind1_18_96)->{'bind'})($f_20))(function($f_prime__22) use ($Applicative0_19_101, $Bind1_18_96, $a_21) {
  $__num = \func_num_args();
  $__res = ((($Bind1_18_96)->{'bind'})($a_21))(function($a_prime__23) use ($Applicative0_19_101, $f_prime__22) {
  $__num = \func_num_args();
  $__res = (($Applicative0_19_101)->{'pure'})(($f_prime__22)($a_prime__23));
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
}, "Functor0" => function($_dollar___unused_18) use ($functorExceptT1_17_92) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_17_92;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_15_103 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($dictMonad_3);
  $__res = (object)["apply" => function($f_16) use ($Applicative0_15_103, $Bind1_14_88) {
  $__num = \func_num_args();
  $__res = function($a_17) use ($Applicative0_15_103, $Bind1_14_88, $f_16) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_88)->{'bind'})($f_16))(function($f_prime__18) use ($Applicative0_15_103, $Bind1_14_88, $a_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_88)->{'bind'})($a_17))(function($a_prime__19) use ($Applicative0_15_103, $f_prime__18) {
  $__num = \func_num_args();
  $__res = (($Applicative0_15_103)->{'pure'})(($f_prime__18)($a_prime__19));
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
}, "Functor0" => function($_dollar___unused_14) use ($functorExceptT1_13_84) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_13_84;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_13) use ($Applicative0_12_84, $Bind1_11_59) {
  $__num = \func_num_args();
  $__res = function($a_14) use ($Applicative0_12_84, $Bind1_11_59, $f_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_59)->{'bind'})($f_13))(function($f_prime__15) use ($Applicative0_12_84, $Bind1_11_59, $a_14) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_59)->{'bind'})($a_14))(function($a_prime__16) use ($Applicative0_12_84, $f_prime__15) {
  $__num = \func_num_args();
  $__res = (($Applicative0_12_84)->{'pure'})(($f_prime__15)($a_prime__16));
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
}, "Functor0" => function($_dollar___unused_11) use ($functorExceptT1_10_55) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_10_55;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_8_106 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_8) {
  $__num = \func_num_args();
  $__res = $x_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_3)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_8) use ($dictMonad_3) {
  $__num = \func_num_args();
  $__local_var_9_106 = (((((($dictMonad_3)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_9_106 = (object)["map" => function($f_10) use ($__local_var_9_106) {
  $__num = \func_num_args();
  $__local_var_11_107 = (($__local_var_9_106)->{'map'})(function($m_11) use ($f_10) {
  $__num = \func_num_args();
  $__t107 = null;;
  if ($m_11 instanceof \Data\Either\Data_Either_Left) {
$__t107 = new \Data\Either\Data_Either_Left(($m_11)->{'value0'});
goto end_branch_107;;
};
  if ($m_11 instanceof \Data\Either\Data_Either_Right) {
$__t107 = new \Data\Either\Data_Either_Right(($f_10)(($m_11)->{'value0'}));
goto end_branch_107;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t107 = null;
  end_branch_107:;
  $__res = $__t107;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_12) use ($__local_var_11_107) {
  $__num = \func_num_args();
  $__res = ($__local_var_11_107)($v_12);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_10_110 = (($dictMonad_3)->{'Bind1'})(null);
  $pure_11_111 = ((($dictMonad_3)->{'Applicative0'})(null))->{'pure'};
  $Bind1_10_110 = (object)["bind" => function($v_12) use ($Bind1_10_110, $pure_11_111) {
  $__num = \func_num_args();
  $__res = function($k_13) use ($Bind1_10_110, $pure_11_111, $v_12) {
  $__num = \func_num_args();
  $__local_var_14_112 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_11_111))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_10_110)->{'bind'})($v_12))(function($v2_15) use ($__local_var_14_112, $k_13) {
  $__num = \func_num_args();
  $__t113 = null;;
  if ($v2_15 instanceof \Data\Either\Data_Either_Left) {
$__t113 = ($__local_var_14_112)(($v2_15)->{'value0'});
goto end_branch_113;;
};
  if ($v2_15 instanceof \Data\Either\Data_Either_Right) {
$__t113 = ($k_13)(($v2_15)->{'value0'});
goto end_branch_113;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t113 = null;
  end_branch_113:;
  $__res = $__t113;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_12) use ($dictMonad_3) {
  $__num = \func_num_args();
  $__local_var_13_114 = (((((($dictMonad_3)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_13_114 = (object)["map" => function($f_14) use ($__local_var_13_114) {
  $__num = \func_num_args();
  $__local_var_15_115 = (($__local_var_13_114)->{'map'})(function($m_15) use ($f_14) {
  $__num = \func_num_args();
  $__t115 = null;;
  if ($m_15 instanceof \Data\Either\Data_Either_Left) {
$__t115 = new \Data\Either\Data_Either_Left(($m_15)->{'value0'});
goto end_branch_115;;
};
  if ($m_15 instanceof \Data\Either\Data_Either_Right) {
$__t115 = new \Data\Either\Data_Either_Right(($f_14)(($m_15)->{'value0'}));
goto end_branch_115;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t115 = null;
  end_branch_115:;
  $__res = $__t115;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_16) use ($__local_var_15_115) {
  $__num = \func_num_args();
  $__res = ($__local_var_15_115)($v_16);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_14_118 = (($dictMonad_3)->{'Bind1'})(null);
  $pure_15_119 = ((($dictMonad_3)->{'Applicative0'})(null))->{'pure'};
  $Bind1_14_118 = (object)["bind" => function($v_16) use ($Bind1_14_118, $pure_15_119) {
  $__num = \func_num_args();
  $__res = function($k_17) use ($Bind1_14_118, $pure_15_119, $v_16) {
  $__num = \func_num_args();
  $__local_var_18_120 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_15_119))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_14_118)->{'bind'})($v_16))(function($v2_19) use ($__local_var_18_120, $k_17) {
  $__num = \func_num_args();
  $__t121 = null;;
  if ($v2_19 instanceof \Data\Either\Data_Either_Left) {
$__t121 = ($__local_var_18_120)(($v2_19)->{'value0'});
goto end_branch_121;;
};
  if ($v2_19 instanceof \Data\Either\Data_Either_Right) {
$__t121 = ($k_17)(($v2_19)->{'value0'});
goto end_branch_121;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t121 = null;
  end_branch_121:;
  $__res = $__t121;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_16) use ($dictMonad_3) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($dictMonad_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_15_123 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($dictMonad_3);
  $__res = (object)["apply" => function($f_16) use ($Applicative0_15_123, $Bind1_14_118) {
  $__num = \func_num_args();
  $__res = function($a_17) use ($Applicative0_15_123, $Bind1_14_118, $f_16) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_118)->{'bind'})($f_16))(function($f_prime__18) use ($Applicative0_15_123, $Bind1_14_118, $a_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_118)->{'bind'})($a_17))(function($a_prime__19) use ($Applicative0_15_123, $f_prime__18) {
  $__num = \func_num_args();
  $__res = (($Applicative0_15_123)->{'pure'})(($f_prime__18)($a_prime__19));
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
}, "Functor0" => function($_dollar___unused_14) use ($functorExceptT1_13_114) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_13_114;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_11_125 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($dictMonad_3);
  $__res = (object)["apply" => function($f_12) use ($Applicative0_11_125, $Bind1_10_110) {
  $__num = \func_num_args();
  $__res = function($a_13) use ($Applicative0_11_125, $Bind1_10_110, $f_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_110)->{'bind'})($f_12))(function($f_prime__14) use ($Applicative0_11_125, $Bind1_10_110, $a_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_110)->{'bind'})($a_13))(function($a_prime__15) use ($Applicative0_11_125, $f_prime__14) {
  $__num = \func_num_args();
  $__res = (($Applicative0_11_125)->{'pure'})(($f_prime__14)($a_prime__15));
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
}, "Functor0" => function($_dollar___unused_10) use ($functorExceptT1_9_106) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_9_106;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_9) use ($Applicative0_8_106, $Bind1_7_51) {
  $__num = \func_num_args();
  $__res = function($a_10) use ($Applicative0_8_106, $Bind1_7_51, $f_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_51)->{'bind'})($f_9))(function($f_prime__11) use ($Applicative0_8_106, $Bind1_7_51, $a_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_51)->{'bind'})($a_10))(function($a_prime__12) use ($Applicative0_8_106, $f_prime__11) {
  $__num = \func_num_args();
  $__res = (($Applicative0_8_106)->{'pure'})(($f_prime__11)($a_prime__12));
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
}, "Functor0" => function($_dollar___unused_7) use ($functorExceptT1_6_47) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_6_47;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_6_128 = (($dictMonad_3)->{'Bind1'})(null);
  $Applicative0_7_129 = (($dictMonad_3)->{'Applicative0'})(null);
  $__local_var_8_130 = (((((($dictMonad_3)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_8_130 = (object)["map" => function($f_9) use ($__local_var_8_130) {
  $__num = \func_num_args();
  $__local_var_10_131 = (($__local_var_8_130)->{'map'})(function($m_10) use ($f_9) {
  $__num = \func_num_args();
  $__t131 = null;;
  if ($m_10 instanceof \Data\Either\Data_Either_Left) {
$__t131 = new \Data\Either\Data_Either_Left(($m_10)->{'value0'});
goto end_branch_131;;
};
  if ($m_10 instanceof \Data\Either\Data_Either_Right) {
$__t131 = new \Data\Either\Data_Either_Right(($f_9)(($m_10)->{'value0'}));
goto end_branch_131;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t131 = null;
  end_branch_131:;
  $__res = $__t131;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_11) use ($__local_var_10_131) {
  $__num = \func_num_args();
  $__res = ($__local_var_10_131)($v_11);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $altExceptT2_6_128 = (object)["alt" => function($v_9) use ($Applicative0_7_129, $Bind1_6_128, $__local_var_2_1) {
  $__num = \func_num_args();
  $__res = function($v1_10) use ($Applicative0_7_129, $Bind1_6_128, $__local_var_2_1, $v_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_128)->{'bind'})($v_9))(function($rm_11) use ($Applicative0_7_129, $Bind1_6_128, $__local_var_2_1, $v1_10) {
  $__num = \func_num_args();
  $__t134 = null;;
  if ($rm_11 instanceof \Data\Either\Data_Either_Right) {
$__t134 = (($Applicative0_7_129)->{'pure'})(new \Data\Either\Data_Either_Right(($rm_11)->{'value0'}));
goto end_branch_134;;
};
  if ($rm_11 instanceof \Data\Either\Data_Either_Left) {
$__local_var_12_135 = ($rm_11)->{'value0'};
$__t134 = ((($Bind1_6_128)->{'bind'})($v1_10))(function($rn_13) use ($Applicative0_7_129, $__local_var_12_135, $__local_var_2_1) {
  $__num = \func_num_args();
  $__t136 = null;;
  if ($rn_13 instanceof \Data\Either\Data_Either_Right) {
$__t136 = (($Applicative0_7_129)->{'pure'})(new \Data\Either\Data_Either_Right(($rn_13)->{'value0'}));
goto end_branch_136;;
};
  if ($rn_13 instanceof \Data\Either\Data_Either_Left) {
$__t136 = (($Applicative0_7_129)->{'pure'})(new \Data\Either\Data_Either_Left(((($__local_var_2_1)->{'append'})($__local_var_12_135))(($rn_13)->{'value0'})));
goto end_branch_136;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t136 = null;
  end_branch_136:;
  $__res = $__t136;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
goto end_branch_134;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t134 = null;
  end_branch_134:;
  $__res = $__t134;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_9) use ($functorExceptT1_8_130) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_8_130;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $plusExceptT2_6_128 = (object)["empty" => \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($x_7) {
  $__num = \func_num_args();
  $__res = $x_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_3)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Left']), $mempty_1_0), "Alt0" => function($_dollar___unused_7) use ($altExceptT2_6_128) {
  $__num = \func_num_args();
  $__res = $altExceptT2_6_128;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $alternativeExceptT2_5_47 = (object)["Applicative0" => function($_dollar___unused_7) use ($applicativeExceptT1_5_47) {
  $__num = \func_num_args();
  $__res = $applicativeExceptT1_5_47;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Plus1" => function($_dollar___unused_7) use ($plusExceptT2_6_128) {
  $__num = \func_num_args();
  $__res = $plusExceptT2_6_128;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["Monad0" => function($_dollar___unused_6) use ($monadExceptT1_4_2) {
  $__num = \func_num_args();
  $__res = $monadExceptT1_4_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Alternative1" => function($_dollar___unused_6) use ($alternativeExceptT2_5_47) {
  $__num = \func_num_args();
  $__res = $alternativeExceptT2_5_47;
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
$GLOBALS['Control_Monad_Except_Trans_monadPlusExceptT'] = __NAMESPACE__ . '\\majControl_majMonad_majExcept_majTrans_monadmajPlusmajExceptmajT';

