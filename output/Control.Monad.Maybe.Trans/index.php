<?php

namespace Control\Monad\Maybe\Trans;

// ALL IMPORTS: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Monad, Control.Monad.Cont.Class, Control.Monad.Error.Class, Control.Monad.Maybe.Trans, Control.Monad.Reader.Class, Control.Monad.Rec.Class, Control.Monad.ST.Class, Control.Monad.State.Class, Control.Monad.Trans.Class, Control.Monad.Writer.Class, Control.MonadPlus, Control.Plus, Control.Semigroupoid, Data.Function, Data.Functor, Data.Maybe, Data.Monoid, Data.Newtype, Data.Semigroup, Data.Tuple, Effect.Class, Prelude, Prim
// TO REQUIRE: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Monad, Control.Monad.Cont.Class, Control.Monad.Error.Class, Control.Monad.Maybe.Trans, Control.Monad.Reader.Class, Control.Monad.Rec.Class, Control.Monad.ST.Class, Control.Monad.State.Class, Control.Monad.Trans.Class, Control.Monad.Writer.Class, Control.MonadPlus, Control.Plus, Control.Semigroupoid, Data.Function, Data.Functor, Data.Maybe, Data.Monoid, Data.Newtype, Data.Semigroup, Data.Tuple, Effect.Class, Prelude
require_once __DIR__ . '/../Control.Alt/index.php';
require_once __DIR__ . '/../Control.Alternative/index.php';
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Category/index.php';
require_once __DIR__ . '/../Control.Monad/index.php';
require_once __DIR__ . '/../Control.Monad.Cont.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Error.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Maybe.Trans/index.php';
require_once __DIR__ . '/../Control.Monad.Reader.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Rec.Class/index.php';
require_once __DIR__ . '/../Control.Monad.ST.Class/index.php';
require_once __DIR__ . '/../Control.Monad.State.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Trans.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Writer.Class/index.php';
require_once __DIR__ . '/../Control.MonadPlus/index.php';
require_once __DIR__ . '/../Control.Plus/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
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




// Control_Monad_Maybe_Trans_MaybeT
function majControl_majMonad_majMaybe_majTrans_majMaybemajT($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majMaybe_majTrans_majMaybemajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Maybe_Trans_MaybeT'] = __NAMESPACE__ . '\\majControl_majMonad_majMaybe_majTrans_majMaybemajT';

// Control_Monad_Maybe_Trans_runMaybeT
function majControl_majMonad_majMaybe_majTrans_runmajMaybemajT($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majMaybe_majTrans_runmajMaybemajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $v_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Maybe_Trans_runMaybeT'] = __NAMESPACE__ . '\\majControl_majMonad_majMaybe_majTrans_runmajMaybemajT';

// Control_Monad_Maybe_Trans_newtypeMaybeT
$GLOBALS['Control_Monad_Maybe_Trans_newtypeMaybeT'] = (object)["Coercible0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Control_Monad_Maybe_Trans_monadTransMaybeT
$GLOBALS['Control_Monad_Maybe_Trans_monadTransMaybeT'] = (object)["lift" => function($dictMonad_0) {
  $__num = \func_num_args();
  $Bind1_1_0 = (($dictMonad_0)->{'Bind1'})(null);
  $Applicative0_2_1 = (($dictMonad_0)->{'Applicative0'})(null);
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_1) {
  $__num = \func_num_args();
  $__res = $x_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($a_3) use ($Applicative0_2_1, $Bind1_1_0) {
  $__num = \func_num_args();
  $__res = ((($Bind1_1_0)->{'bind'})($a_3))(function($a_prime__4) use ($Applicative0_2_1) {
  $__num = \func_num_args();
  $__res = (($Applicative0_2_1)->{'pure'})(new \Data\Maybe\Data_Maybe_Just($a_prime__4));
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
}];

// Control_Monad_Maybe_Trans_mapMaybeT
function majControl_majMonad_majMaybe_majTrans_mapmajMaybemajT($f_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majMaybe_majTrans_mapmajMaybemajT';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($f_0)($v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Control_Monad_Maybe_Trans_mapMaybeT'] = __NAMESPACE__ . '\\majControl_majMonad_majMaybe_majTrans_mapmajMaybemajT';

// Control_Monad_Maybe_Trans_functorMaybeT
function majControl_majMonad_majMaybe_majTrans_functormajMaybemajT($dictFunctor_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majMaybe_majTrans_functormajMaybemajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["map" => function($f_1) use ($dictFunctor_0) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($dictFunctor_0, $f_1) {
  $__num = \func_num_args();
  $__res = ((($dictFunctor_0)->{'map'})(function($v1_3) use ($f_1) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v1_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t0 = new \Data\Maybe\Data_Maybe_Just(($f_1)(($v1_3)->{'value0'}));
goto end_branch_0;;
};
  $__t0 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_2);
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
$GLOBALS['Control_Monad_Maybe_Trans_functorMaybeT'] = __NAMESPACE__ . '\\majControl_majMonad_majMaybe_majTrans_functormajMaybemajT';

// Control_Monad_Maybe_Trans_monadMaybeT
function majControl_majMonad_majMaybe_majTrans_monadmajMaybemajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majMaybe_majTrans_monadmajMaybemajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["Applicative0" => function($_dollar___unused_1) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_1) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_bindMaybeT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Maybe_Trans_monadMaybeT'] = __NAMESPACE__ . '\\majControl_majMonad_majMaybe_majTrans_monadmajMaybemajT';

// Control_Monad_Maybe_Trans_bindMaybeT
function majControl_majMonad_majMaybe_majTrans_bindmajMaybemajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majMaybe_majTrans_bindmajMaybemajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Bind1_1_0 = (($dictMonad_0)->{'Bind1'})(null);
  $Applicative0_2_1 = (($dictMonad_0)->{'Applicative0'})(null);
  $__res = (object)["bind" => function($v_3) use ($Applicative0_2_1, $Bind1_1_0) {
  $__num = \func_num_args();
  $__res = function($f_4) use ($Applicative0_2_1, $Bind1_1_0, $v_3) {
  $__num = \func_num_args();
  $__res = ((($Bind1_1_0)->{'bind'})($v_3))(function($v1_5) use ($Applicative0_2_1, $f_4) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ($v1_5 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t2 = (($Applicative0_2_1)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_2;;
};
  if ($v1_5 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t2 = ($f_4)(($v1_5)->{'value0'});
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_3) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Maybe_Trans_bindMaybeT'] = __NAMESPACE__ . '\\majControl_majMonad_majMaybe_majTrans_bindmajMaybemajT';

// Control_Monad_Maybe_Trans_applyMaybeT
function majControl_majMonad_majMaybe_majTrans_applymajMaybemajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majMaybe_majTrans_applymajMaybemajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_1_0 = (object)["map" => function($f_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($__local_var_1_0, $f_2) {
  $__num = \func_num_args();
  $__res = ((($__local_var_1_0)->{'map'})(function($v1_4) use ($f_2) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v1_4 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = new \Data\Maybe\Data_Maybe_Just(($f_2)(($v1_4)->{'value0'}));
goto end_branch_1;;
};
  $__t1 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_2_3 = (($dictMonad_0)->{'Bind1'})(null);
  $Applicative0_3_4 = (($dictMonad_0)->{'Applicative0'})(null);
  $Bind1_2_3 = (object)["bind" => function($v_4) use ($Applicative0_3_4, $Bind1_2_3) {
  $__num = \func_num_args();
  $__res = function($f_5) use ($Applicative0_3_4, $Bind1_2_3, $v_4) {
  $__num = \func_num_args();
  $__res = ((($Bind1_2_3)->{'bind'})($v_4))(function($v1_6) use ($Applicative0_3_4, $f_5) {
  $__num = \func_num_args();
  $__t5 = null;;
  if ($v1_6 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t5 = (($Applicative0_3_4)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_5;;
};
  if ($v1_6 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t5 = ($f_5)(($v1_6)->{'value0'});
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_4) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_3_7 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_4) use ($Applicative0_3_7, $Bind1_2_3) {
  $__num = \func_num_args();
  $__res = function($a_5) use ($Applicative0_3_7, $Bind1_2_3, $f_4) {
  $__num = \func_num_args();
  $__res = ((($Bind1_2_3)->{'bind'})($f_4))(function($f_prime__6) use ($Applicative0_3_7, $Bind1_2_3, $a_5) {
  $__num = \func_num_args();
  $__res = ((($Bind1_2_3)->{'bind'})($a_5))(function($a_prime__7) use ($Applicative0_3_7, $f_prime__6) {
  $__num = \func_num_args();
  $__res = (($Applicative0_3_7)->{'pure'})(($f_prime__6)($a_prime__7));
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
}, "Functor0" => function($_dollar___unused_2) use ($functorMaybeT1_1_0) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'] = __NAMESPACE__ . '\\majControl_majMonad_majMaybe_majTrans_applymajMaybemajT';

// Control_Monad_Maybe_Trans_applicativeMaybeT
function majControl_majMonad_majMaybe_majTrans_applicativemajMaybemajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majMaybe_majTrans_applicativemajMaybemajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_1) {
  $__num = \func_num_args();
  $__res = $x_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_1) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_2_0 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_2_0 = (object)["map" => function($f_3) use ($__local_var_2_0) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($__local_var_2_0, $f_3) {
  $__num = \func_num_args();
  $__res = ((($__local_var_2_0)->{'map'})(function($v1_5) use ($f_3) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v1_5 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = new \Data\Maybe\Data_Maybe_Just(($f_3)(($v1_5)->{'value0'}));
goto end_branch_1;;
};
  $__t1 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_3_3 = (($dictMonad_0)->{'Bind1'})(null);
  $Applicative0_4_4 = (($dictMonad_0)->{'Applicative0'})(null);
  $Bind1_3_3 = (object)["bind" => function($v_5) use ($Applicative0_4_4, $Bind1_3_3) {
  $__num = \func_num_args();
  $__res = function($f_6) use ($Applicative0_4_4, $Bind1_3_3, $v_5) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_3)->{'bind'})($v_5))(function($v1_7) use ($Applicative0_4_4, $f_6) {
  $__num = \func_num_args();
  $__t5 = null;;
  if ($v1_7 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t5 = (($Applicative0_4_4)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_5;;
};
  if ($v1_7 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t5 = ($f_6)(($v1_7)->{'value0'});
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_5) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_4_7 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_5) use ($Applicative0_4_7, $Bind1_3_3) {
  $__num = \func_num_args();
  $__res = function($a_6) use ($Applicative0_4_7, $Bind1_3_3, $f_5) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_3)->{'bind'})($f_5))(function($f_prime__7) use ($Applicative0_4_7, $Bind1_3_3, $a_6) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_3)->{'bind'})($a_6))(function($a_prime__8) use ($Applicative0_4_7, $f_prime__7) {
  $__num = \func_num_args();
  $__res = (($Applicative0_4_7)->{'pure'})(($f_prime__7)($a_prime__8));
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
}, "Functor0" => function($_dollar___unused_3) use ($functorMaybeT1_2_0) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_2_0;
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
$GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'] = __NAMESPACE__ . '\\majControl_majMonad_majMaybe_majTrans_applicativemajMaybemajT';

// Control_Monad_Maybe_Trans_semigroupMaybeT
function majControl_majMonad_majMaybe_majTrans_semigroupmajMaybemajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majMaybe_majTrans_semigroupmajMaybemajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_1_0 = (object)["map" => function($f_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($__local_var_1_0, $f_2) {
  $__num = \func_num_args();
  $__res = ((($__local_var_1_0)->{'map'})(function($v1_4) use ($f_2) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v1_4 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = new \Data\Maybe\Data_Maybe_Just(($f_2)(($v1_4)->{'value0'}));
goto end_branch_1;;
};
  $__t1 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_2_3 = (($dictMonad_0)->{'Bind1'})(null);
  $Applicative0_3_4 = (($dictMonad_0)->{'Applicative0'})(null);
  $Bind1_2_3 = (object)["bind" => function($v_4) use ($Applicative0_3_4, $Bind1_2_3) {
  $__num = \func_num_args();
  $__res = function($f_5) use ($Applicative0_3_4, $Bind1_2_3, $v_4) {
  $__num = \func_num_args();
  $__res = ((($Bind1_2_3)->{'bind'})($v_4))(function($v1_6) use ($Applicative0_3_4, $f_5) {
  $__num = \func_num_args();
  $__t5 = null;;
  if ($v1_6 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t5 = (($Applicative0_3_4)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_5;;
};
  if ($v1_6 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t5 = ($f_5)(($v1_6)->{'value0'});
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_4) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_3_7 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_3) {
  $__num = \func_num_args();
  $__res = $x_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_3) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_4_7 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_4_7 = (object)["map" => function($f_5) use ($__local_var_4_7) {
  $__num = \func_num_args();
  $__res = function($v_6) use ($__local_var_4_7, $f_5) {
  $__num = \func_num_args();
  $__res = ((($__local_var_4_7)->{'map'})(function($v1_7) use ($f_5) {
  $__num = \func_num_args();
  $__t8 = null;;
  if ($v1_7 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t8 = new \Data\Maybe\Data_Maybe_Just(($f_5)(($v1_7)->{'value0'}));
goto end_branch_8;;
};
  $__t8 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_8:;
  $__res = $__t8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_5_10 = (($dictMonad_0)->{'Bind1'})(null);
  $Applicative0_6_11 = (($dictMonad_0)->{'Applicative0'})(null);
  $Bind1_5_10 = (object)["bind" => function($v_7) use ($Applicative0_6_11, $Bind1_5_10) {
  $__num = \func_num_args();
  $__res = function($f_8) use ($Applicative0_6_11, $Bind1_5_10, $v_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_5_10)->{'bind'})($v_7))(function($v1_9) use ($Applicative0_6_11, $f_8) {
  $__num = \func_num_args();
  $__t12 = null;;
  if ($v1_9 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t12 = (($Applicative0_6_11)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_12;;
};
  if ($v1_9 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t12 = ($f_8)(($v1_9)->{'value0'});
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_7) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_6_14 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_7) use ($Applicative0_6_14, $Bind1_5_10) {
  $__num = \func_num_args();
  $__res = function($a_8) use ($Applicative0_6_14, $Bind1_5_10, $f_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_5_10)->{'bind'})($f_7))(function($f_prime__9) use ($Applicative0_6_14, $Bind1_5_10, $a_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_5_10)->{'bind'})($a_8))(function($a_prime__10) use ($Applicative0_6_14, $f_prime__9) {
  $__num = \func_num_args();
  $__res = (($Applicative0_6_14)->{'pure'})(($f_prime__9)($a_prime__10));
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
}, "Functor0" => function($_dollar___unused_5) use ($functorMaybeT1_4_7) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_4_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyMaybeT1_1_0 = (object)["apply" => function($f_4) use ($Applicative0_3_7, $Bind1_2_3) {
  $__num = \func_num_args();
  $__res = function($a_5) use ($Applicative0_3_7, $Bind1_2_3, $f_4) {
  $__num = \func_num_args();
  $__res = ((($Bind1_2_3)->{'bind'})($f_4))(function($f_prime__6) use ($Applicative0_3_7, $Bind1_2_3, $a_5) {
  $__num = \func_num_args();
  $__res = ((($Bind1_2_3)->{'bind'})($a_5))(function($a_prime__7) use ($Applicative0_3_7, $f_prime__6) {
  $__num = \func_num_args();
  $__res = (($Applicative0_3_7)->{'pure'})(($f_prime__6)($a_prime__7));
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
}, "Functor0" => function($_dollar___unused_2) use ($functorMaybeT1_1_0) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = function($dictSemigroup_2) use ($applyMaybeT1_1_0) {
  $__num = \func_num_args();
  $Functor0_3_17 = (($applyMaybeT1_1_0)->{'Functor0'})(null);
  $__local_var_4_18 = ($dictSemigroup_2)->{'append'};
  $__res = (object)["append" => function($a_5) use ($Functor0_3_17, $__local_var_4_18, $applyMaybeT1_1_0) {
  $__num = \func_num_args();
  $__res = function($b_6) use ($Functor0_3_17, $__local_var_4_18, $a_5, $applyMaybeT1_1_0) {
  $__num = \func_num_args();
  $__res = ((($applyMaybeT1_1_0)->{'apply'})(((($Functor0_3_17)->{'map'})($__local_var_4_18))($a_5)))($b_6);
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
$GLOBALS['Control_Monad_Maybe_Trans_semigroupMaybeT'] = __NAMESPACE__ . '\\majControl_majMonad_majMaybe_majTrans_semigroupmajMaybemajT';

// Control_Monad_Maybe_Trans_monadAskMaybeT
function majControl_majMonad_majMaybe_majTrans_monadmajAskmajMaybemajT($dictMonadAsk_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majMaybe_majTrans_monadmajAskmajMaybemajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictMonadAsk_0)->{'Monad0'})(null);
  $monadMaybeT1_1_0 = (object)["Applicative0" => function($_dollar___unused_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_3) {
  $__num = \func_num_args();
  $__res = $x_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($__local_var_1_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_3) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__local_var_4_1 = (((((($__local_var_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_4_1 = (object)["map" => function($f_5) use ($__local_var_4_1) {
  $__num = \func_num_args();
  $__res = function($v_6) use ($__local_var_4_1, $f_5) {
  $__num = \func_num_args();
  $__res = ((($__local_var_4_1)->{'map'})(function($v1_7) use ($f_5) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ($v1_7 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t2 = new \Data\Maybe\Data_Maybe_Just(($f_5)(($v1_7)->{'value0'}));
goto end_branch_2;;
};
  $__t2 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_5_4 = (($__local_var_1_0)->{'Bind1'})(null);
  $Applicative0_6_5 = (($__local_var_1_0)->{'Applicative0'})(null);
  $Bind1_5_4 = (object)["bind" => function($v_7) use ($Applicative0_6_5, $Bind1_5_4) {
  $__num = \func_num_args();
  $__res = function($f_8) use ($Applicative0_6_5, $Bind1_5_4, $v_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_5_4)->{'bind'})($v_7))(function($v1_9) use ($Applicative0_6_5, $f_8) {
  $__num = \func_num_args();
  $__t6 = null;;
  if ($v1_9 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t6 = (($Applicative0_6_5)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_6;;
};
  if ($v1_9 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t6 = ($f_8)(($v1_9)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_7) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__local_var_8_7 = (((((($__local_var_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_8_7 = (object)["map" => function($f_9) use ($__local_var_8_7) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($__local_var_8_7, $f_9) {
  $__num = \func_num_args();
  $__res = ((($__local_var_8_7)->{'map'})(function($v1_11) use ($f_9) {
  $__num = \func_num_args();
  $__t8 = null;;
  if ($v1_11 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t8 = new \Data\Maybe\Data_Maybe_Just(($f_9)(($v1_11)->{'value0'}));
goto end_branch_8;;
};
  $__t8 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_8:;
  $__res = $__t8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_10);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_9_10 = (($__local_var_1_0)->{'Bind1'})(null);
  $Applicative0_10_11 = (($__local_var_1_0)->{'Applicative0'})(null);
  $Bind1_9_10 = (object)["bind" => function($v_11) use ($Applicative0_10_11, $Bind1_9_10) {
  $__num = \func_num_args();
  $__res = function($f_12) use ($Applicative0_10_11, $Bind1_9_10, $v_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_10)->{'bind'})($v_11))(function($v1_13) use ($Applicative0_10_11, $f_12) {
  $__num = \func_num_args();
  $__t12 = null;;
  if ($v1_13 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t12 = (($Applicative0_10_11)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_12;;
};
  if ($v1_13 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t12 = ($f_12)(($v1_13)->{'value0'});
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_11) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($__local_var_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_10_14 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($__local_var_1_0);
  $__res = (object)["apply" => function($f_11) use ($Applicative0_10_14, $Bind1_9_10) {
  $__num = \func_num_args();
  $__res = function($a_12) use ($Applicative0_10_14, $Bind1_9_10, $f_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_10)->{'bind'})($f_11))(function($f_prime__13) use ($Applicative0_10_14, $Bind1_9_10, $a_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_10)->{'bind'})($a_12))(function($a_prime__14) use ($Applicative0_10_14, $f_prime__13) {
  $__num = \func_num_args();
  $__res = (($Applicative0_10_14)->{'pure'})(($f_prime__13)($a_prime__14));
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
}, "Functor0" => function($_dollar___unused_9) use ($functorMaybeT1_8_7) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_8_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_6_16 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($__local_var_1_0);
  $__res = (object)["apply" => function($f_7) use ($Applicative0_6_16, $Bind1_5_4) {
  $__num = \func_num_args();
  $__res = function($a_8) use ($Applicative0_6_16, $Bind1_5_4, $f_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_5_4)->{'bind'})($f_7))(function($f_prime__9) use ($Applicative0_6_16, $Bind1_5_4, $a_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_5_4)->{'bind'})($a_8))(function($a_prime__10) use ($Applicative0_6_16, $f_prime__9) {
  $__num = \func_num_args();
  $__res = (($Applicative0_6_16)->{'pure'})(($f_prime__9)($a_prime__10));
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
}, "Functor0" => function($_dollar___unused_5) use ($functorMaybeT1_4_1) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_4_1;
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
  $Bind1_3_17 = (($__local_var_1_0)->{'Bind1'})(null);
  $Applicative0_4_18 = (($__local_var_1_0)->{'Applicative0'})(null);
  $__res = (object)["bind" => function($v_5) use ($Applicative0_4_18, $Bind1_3_17) {
  $__num = \func_num_args();
  $__res = function($f_6) use ($Applicative0_4_18, $Bind1_3_17, $v_5) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_17)->{'bind'})($v_5))(function($v1_7) use ($Applicative0_4_18, $f_6) {
  $__num = \func_num_args();
  $__t19 = null;;
  if ($v1_7 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t19 = (($Applicative0_4_18)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_19;;
};
  if ($v1_7 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t19 = ($f_6)(($v1_7)->{'value0'});
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_5) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__local_var_6_20 = (((((($__local_var_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_6_20 = (object)["map" => function($f_7) use ($__local_var_6_20) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($__local_var_6_20, $f_7) {
  $__num = \func_num_args();
  $__res = ((($__local_var_6_20)->{'map'})(function($v1_9) use ($f_7) {
  $__num = \func_num_args();
  $__t21 = null;;
  if ($v1_9 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t21 = new \Data\Maybe\Data_Maybe_Just(($f_7)(($v1_9)->{'value0'}));
goto end_branch_21;;
};
  $__t21 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_21:;
  $__res = $__t21;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_8);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_7_23 = (($__local_var_1_0)->{'Bind1'})(null);
  $Applicative0_8_24 = (($__local_var_1_0)->{'Applicative0'})(null);
  $Bind1_7_23 = (object)["bind" => function($v_9) use ($Applicative0_8_24, $Bind1_7_23) {
  $__num = \func_num_args();
  $__res = function($f_10) use ($Applicative0_8_24, $Bind1_7_23, $v_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_23)->{'bind'})($v_9))(function($v1_11) use ($Applicative0_8_24, $f_10) {
  $__num = \func_num_args();
  $__t25 = null;;
  if ($v1_11 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t25 = (($Applicative0_8_24)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_25;;
};
  if ($v1_11 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t25 = ($f_10)(($v1_11)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_9) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($__local_var_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_8_27 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_8) {
  $__num = \func_num_args();
  $__res = $x_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($__local_var_1_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_8) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__local_var_9_27 = (((((($__local_var_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_9_27 = (object)["map" => function($f_10) use ($__local_var_9_27) {
  $__num = \func_num_args();
  $__res = function($v_11) use ($__local_var_9_27, $f_10) {
  $__num = \func_num_args();
  $__res = ((($__local_var_9_27)->{'map'})(function($v1_12) use ($f_10) {
  $__num = \func_num_args();
  $__t28 = null;;
  if ($v1_12 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t28 = new \Data\Maybe\Data_Maybe_Just(($f_10)(($v1_12)->{'value0'}));
goto end_branch_28;;
};
  $__t28 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_28:;
  $__res = $__t28;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_11);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_10_30 = (($__local_var_1_0)->{'Bind1'})(null);
  $Applicative0_11_31 = (($__local_var_1_0)->{'Applicative0'})(null);
  $Bind1_10_30 = (object)["bind" => function($v_12) use ($Applicative0_11_31, $Bind1_10_30) {
  $__num = \func_num_args();
  $__res = function($f_13) use ($Applicative0_11_31, $Bind1_10_30, $v_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_30)->{'bind'})($v_12))(function($v1_14) use ($Applicative0_11_31, $f_13) {
  $__num = \func_num_args();
  $__t32 = null;;
  if ($v1_14 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t32 = (($Applicative0_11_31)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_32;;
};
  if ($v1_14 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t32 = ($f_13)(($v1_14)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_12) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($__local_var_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_11_34 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($__local_var_1_0);
  $__res = (object)["apply" => function($f_12) use ($Applicative0_11_34, $Bind1_10_30) {
  $__num = \func_num_args();
  $__res = function($a_13) use ($Applicative0_11_34, $Bind1_10_30, $f_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_30)->{'bind'})($f_12))(function($f_prime__14) use ($Applicative0_11_34, $Bind1_10_30, $a_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_30)->{'bind'})($a_13))(function($a_prime__15) use ($Applicative0_11_34, $f_prime__14) {
  $__num = \func_num_args();
  $__res = (($Applicative0_11_34)->{'pure'})(($f_prime__14)($a_prime__15));
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
}, "Functor0" => function($_dollar___unused_10) use ($functorMaybeT1_9_27) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_9_27;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_9) use ($Applicative0_8_27, $Bind1_7_23) {
  $__num = \func_num_args();
  $__res = function($a_10) use ($Applicative0_8_27, $Bind1_7_23, $f_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_23)->{'bind'})($f_9))(function($f_prime__11) use ($Applicative0_8_27, $Bind1_7_23, $a_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_23)->{'bind'})($a_10))(function($a_prime__12) use ($Applicative0_8_27, $f_prime__11) {
  $__num = \func_num_args();
  $__res = (($Applicative0_8_27)->{'pure'})(($f_prime__11)($a_prime__12));
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
}, "Functor0" => function($_dollar___unused_7) use ($functorMaybeT1_6_20) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_6_20;
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
  $__local_var_2_37 = (($dictMonadAsk_0)->{'Monad0'})(null);
  $Bind1_3_38 = (($__local_var_2_37)->{'Bind1'})(null);
  $Applicative0_4_39 = (($__local_var_2_37)->{'Applicative0'})(null);
  $__res = (object)["ask" => \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($x_3) {
  $__num = \func_num_args();
  $__res = $x_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, function($a_5) use ($Applicative0_4_39, $Bind1_3_38) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_38)->{'bind'})($a_5))(function($a_prime__6) use ($Applicative0_4_39) {
  $__num = \func_num_args();
  $__res = (($Applicative0_4_39)->{'pure'})(new \Data\Maybe\Data_Maybe_Just($a_prime__6));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($dictMonadAsk_0)->{'ask'}), "Monad0" => function($_dollar___unused_2) use ($monadMaybeT1_1_0) {
  $__num = \func_num_args();
  $__res = $monadMaybeT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Maybe_Trans_monadAskMaybeT'] = __NAMESPACE__ . '\\majControl_majMonad_majMaybe_majTrans_monadmajAskmajMaybemajT';

// Control_Monad_Maybe_Trans_monadReaderMaybeT
function majControl_majMonad_majMaybe_majTrans_monadmajReadermajMaybemajT($dictMonadReader_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majMaybe_majTrans_monadmajReadermajMaybemajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictMonadReader_0)->{'MonadAsk0'})(null);
  $__local_var_2_1 = (($__local_var_1_0)->{'Monad0'})(null);
  $monadMaybeT1_2_1 = (object)["Applicative0" => function($_dollar___unused_3) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_4) {
  $__num = \func_num_args();
  $__res = $x_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_4) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_5_2 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_5_2 = (object)["map" => function($f_6) use ($__local_var_5_2) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($__local_var_5_2, $f_6) {
  $__num = \func_num_args();
  $__res = ((($__local_var_5_2)->{'map'})(function($v1_8) use ($f_6) {
  $__num = \func_num_args();
  $__t3 = null;;
  if ($v1_8 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t3 = new \Data\Maybe\Data_Maybe_Just(($f_6)(($v1_8)->{'value0'}));
goto end_branch_3;;
};
  $__t3 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_6_5 = (($__local_var_2_1)->{'Bind1'})(null);
  $Applicative0_7_6 = (($__local_var_2_1)->{'Applicative0'})(null);
  $Bind1_6_5 = (object)["bind" => function($v_8) use ($Applicative0_7_6, $Bind1_6_5) {
  $__num = \func_num_args();
  $__res = function($f_9) use ($Applicative0_7_6, $Bind1_6_5, $v_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_5)->{'bind'})($v_8))(function($v1_10) use ($Applicative0_7_6, $f_9) {
  $__num = \func_num_args();
  $__t7 = null;;
  if ($v1_10 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t7 = (($Applicative0_7_6)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_7;;
};
  if ($v1_10 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t7 = ($f_9)(($v1_10)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_8) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_9_8 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_9_8 = (object)["map" => function($f_10) use ($__local_var_9_8) {
  $__num = \func_num_args();
  $__res = function($v_11) use ($__local_var_9_8, $f_10) {
  $__num = \func_num_args();
  $__res = ((($__local_var_9_8)->{'map'})(function($v1_12) use ($f_10) {
  $__num = \func_num_args();
  $__t9 = null;;
  if ($v1_12 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t9 = new \Data\Maybe\Data_Maybe_Just(($f_10)(($v1_12)->{'value0'}));
goto end_branch_9;;
};
  $__t9 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_9:;
  $__res = $__t9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_11);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_10_11 = (($__local_var_2_1)->{'Bind1'})(null);
  $Applicative0_11_12 = (($__local_var_2_1)->{'Applicative0'})(null);
  $Bind1_10_11 = (object)["bind" => function($v_12) use ($Applicative0_11_12, $Bind1_10_11) {
  $__num = \func_num_args();
  $__res = function($f_13) use ($Applicative0_11_12, $Bind1_10_11, $v_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_11)->{'bind'})($v_12))(function($v1_14) use ($Applicative0_11_12, $f_13) {
  $__num = \func_num_args();
  $__t13 = null;;
  if ($v1_14 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t13 = (($Applicative0_11_12)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_13;;
};
  if ($v1_14 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t13 = ($f_13)(($v1_14)->{'value0'});
goto end_branch_13;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t13 = null;
  end_branch_13:;
  $__res = $__t13;
  goto __end;;
  __end:
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
  $__local_var_13_14 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_13_14 = (object)["map" => function($f_14) use ($__local_var_13_14) {
  $__num = \func_num_args();
  $__res = function($v_15) use ($__local_var_13_14, $f_14) {
  $__num = \func_num_args();
  $__res = ((($__local_var_13_14)->{'map'})(function($v1_16) use ($f_14) {
  $__num = \func_num_args();
  $__t15 = null;;
  if ($v1_16 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t15 = new \Data\Maybe\Data_Maybe_Just(($f_14)(($v1_16)->{'value0'}));
goto end_branch_15;;
};
  $__t15 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_15:;
  $__res = $__t15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_15);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_14_17 = (($__local_var_2_1)->{'Bind1'})(null);
  $Applicative0_15_18 = (($__local_var_2_1)->{'Applicative0'})(null);
  $Bind1_14_17 = (object)["bind" => function($v_16) use ($Applicative0_15_18, $Bind1_14_17) {
  $__num = \func_num_args();
  $__res = function($f_17) use ($Applicative0_15_18, $Bind1_14_17, $v_16) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_17)->{'bind'})($v_16))(function($v1_18) use ($Applicative0_15_18, $f_17) {
  $__num = \func_num_args();
  $__t19 = null;;
  if ($v1_18 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t19 = (($Applicative0_15_18)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_19;;
};
  if ($v1_18 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t19 = ($f_17)(($v1_18)->{'value0'});
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_16) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($__local_var_2_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_15_21 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_15) {
  $__num = \func_num_args();
  $__res = $x_15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_15) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_16_21 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_16_21 = (object)["map" => function($f_17) use ($__local_var_16_21) {
  $__num = \func_num_args();
  $__res = function($v_18) use ($__local_var_16_21, $f_17) {
  $__num = \func_num_args();
  $__res = ((($__local_var_16_21)->{'map'})(function($v1_19) use ($f_17) {
  $__num = \func_num_args();
  $__t22 = null;;
  if ($v1_19 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t22 = new \Data\Maybe\Data_Maybe_Just(($f_17)(($v1_19)->{'value0'}));
goto end_branch_22;;
};
  $__t22 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_22:;
  $__res = $__t22;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_18);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_17_24 = (($__local_var_2_1)->{'Bind1'})(null);
  $Applicative0_18_25 = (($__local_var_2_1)->{'Applicative0'})(null);
  $Bind1_17_24 = (object)["bind" => function($v_19) use ($Applicative0_18_25, $Bind1_17_24) {
  $__num = \func_num_args();
  $__res = function($f_20) use ($Applicative0_18_25, $Bind1_17_24, $v_19) {
  $__num = \func_num_args();
  $__res = ((($Bind1_17_24)->{'bind'})($v_19))(function($v1_21) use ($Applicative0_18_25, $f_20) {
  $__num = \func_num_args();
  $__t26 = null;;
  if ($v1_21 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t26 = (($Applicative0_18_25)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_26;;
};
  if ($v1_21 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t26 = ($f_20)(($v1_21)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_19) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($__local_var_2_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_18_28 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($__local_var_2_1);
  $__res = (object)["apply" => function($f_19) use ($Applicative0_18_28, $Bind1_17_24) {
  $__num = \func_num_args();
  $__res = function($a_20) use ($Applicative0_18_28, $Bind1_17_24, $f_19) {
  $__num = \func_num_args();
  $__res = ((($Bind1_17_24)->{'bind'})($f_19))(function($f_prime__21) use ($Applicative0_18_28, $Bind1_17_24, $a_20) {
  $__num = \func_num_args();
  $__res = ((($Bind1_17_24)->{'bind'})($a_20))(function($a_prime__22) use ($Applicative0_18_28, $f_prime__21) {
  $__num = \func_num_args();
  $__res = (($Applicative0_18_28)->{'pure'})(($f_prime__21)($a_prime__22));
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
}, "Functor0" => function($_dollar___unused_17) use ($functorMaybeT1_16_21) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_16_21;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_16) use ($Applicative0_15_21, $Bind1_14_17) {
  $__num = \func_num_args();
  $__res = function($a_17) use ($Applicative0_15_21, $Bind1_14_17, $f_16) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_17)->{'bind'})($f_16))(function($f_prime__18) use ($Applicative0_15_21, $Bind1_14_17, $a_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_17)->{'bind'})($a_17))(function($a_prime__19) use ($Applicative0_15_21, $f_prime__18) {
  $__num = \func_num_args();
  $__res = (($Applicative0_15_21)->{'pure'})(($f_prime__18)($a_prime__19));
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
}, "Functor0" => function($_dollar___unused_14) use ($functorMaybeT1_13_14) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_13_14;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_11_31 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_11) {
  $__num = \func_num_args();
  $__res = $x_11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_11) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_12_31 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_12_31 = (object)["map" => function($f_13) use ($__local_var_12_31) {
  $__num = \func_num_args();
  $__res = function($v_14) use ($__local_var_12_31, $f_13) {
  $__num = \func_num_args();
  $__res = ((($__local_var_12_31)->{'map'})(function($v1_15) use ($f_13) {
  $__num = \func_num_args();
  $__t32 = null;;
  if ($v1_15 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t32 = new \Data\Maybe\Data_Maybe_Just(($f_13)(($v1_15)->{'value0'}));
goto end_branch_32;;
};
  $__t32 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_32:;
  $__res = $__t32;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_14);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_13_34 = (($__local_var_2_1)->{'Bind1'})(null);
  $Applicative0_14_35 = (($__local_var_2_1)->{'Applicative0'})(null);
  $Bind1_13_34 = (object)["bind" => function($v_15) use ($Applicative0_14_35, $Bind1_13_34) {
  $__num = \func_num_args();
  $__res = function($f_16) use ($Applicative0_14_35, $Bind1_13_34, $v_15) {
  $__num = \func_num_args();
  $__res = ((($Bind1_13_34)->{'bind'})($v_15))(function($v1_17) use ($Applicative0_14_35, $f_16) {
  $__num = \func_num_args();
  $__t36 = null;;
  if ($v1_17 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t36 = (($Applicative0_14_35)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_36;;
};
  if ($v1_17 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t36 = ($f_16)(($v1_17)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_15) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_16_37 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_16_37 = (object)["map" => function($f_17) use ($__local_var_16_37) {
  $__num = \func_num_args();
  $__res = function($v_18) use ($__local_var_16_37, $f_17) {
  $__num = \func_num_args();
  $__res = ((($__local_var_16_37)->{'map'})(function($v1_19) use ($f_17) {
  $__num = \func_num_args();
  $__t38 = null;;
  if ($v1_19 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t38 = new \Data\Maybe\Data_Maybe_Just(($f_17)(($v1_19)->{'value0'}));
goto end_branch_38;;
};
  $__t38 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_38:;
  $__res = $__t38;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_18);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_17_40 = (($__local_var_2_1)->{'Bind1'})(null);
  $Applicative0_18_41 = (($__local_var_2_1)->{'Applicative0'})(null);
  $Bind1_17_40 = (object)["bind" => function($v_19) use ($Applicative0_18_41, $Bind1_17_40) {
  $__num = \func_num_args();
  $__res = function($f_20) use ($Applicative0_18_41, $Bind1_17_40, $v_19) {
  $__num = \func_num_args();
  $__res = ((($Bind1_17_40)->{'bind'})($v_19))(function($v1_21) use ($Applicative0_18_41, $f_20) {
  $__num = \func_num_args();
  $__t42 = null;;
  if ($v1_21 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t42 = (($Applicative0_18_41)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_42;;
};
  if ($v1_21 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t42 = ($f_20)(($v1_21)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_19) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($__local_var_2_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_18_44 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($__local_var_2_1);
  $__res = (object)["apply" => function($f_19) use ($Applicative0_18_44, $Bind1_17_40) {
  $__num = \func_num_args();
  $__res = function($a_20) use ($Applicative0_18_44, $Bind1_17_40, $f_19) {
  $__num = \func_num_args();
  $__res = ((($Bind1_17_40)->{'bind'})($f_19))(function($f_prime__21) use ($Applicative0_18_44, $Bind1_17_40, $a_20) {
  $__num = \func_num_args();
  $__res = ((($Bind1_17_40)->{'bind'})($a_20))(function($a_prime__22) use ($Applicative0_18_44, $f_prime__21) {
  $__num = \func_num_args();
  $__res = (($Applicative0_18_44)->{'pure'})(($f_prime__21)($a_prime__22));
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
}, "Functor0" => function($_dollar___unused_17) use ($functorMaybeT1_16_37) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_16_37;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_14_46 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($__local_var_2_1);
  $__res = (object)["apply" => function($f_15) use ($Applicative0_14_46, $Bind1_13_34) {
  $__num = \func_num_args();
  $__res = function($a_16) use ($Applicative0_14_46, $Bind1_13_34, $f_15) {
  $__num = \func_num_args();
  $__res = ((($Bind1_13_34)->{'bind'})($f_15))(function($f_prime__17) use ($Applicative0_14_46, $Bind1_13_34, $a_16) {
  $__num = \func_num_args();
  $__res = ((($Bind1_13_34)->{'bind'})($a_16))(function($a_prime__18) use ($Applicative0_14_46, $f_prime__17) {
  $__num = \func_num_args();
  $__res = (($Applicative0_14_46)->{'pure'})(($f_prime__17)($a_prime__18));
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
}, "Functor0" => function($_dollar___unused_13) use ($functorMaybeT1_12_31) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_12_31;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_12) use ($Applicative0_11_31, $Bind1_10_11) {
  $__num = \func_num_args();
  $__res = function($a_13) use ($Applicative0_11_31, $Bind1_10_11, $f_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_11)->{'bind'})($f_12))(function($f_prime__14) use ($Applicative0_11_31, $Bind1_10_11, $a_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_11)->{'bind'})($a_13))(function($a_prime__15) use ($Applicative0_11_31, $f_prime__14) {
  $__num = \func_num_args();
  $__res = (($Applicative0_11_31)->{'pure'})(($f_prime__14)($a_prime__15));
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
}, "Functor0" => function($_dollar___unused_10) use ($functorMaybeT1_9_8) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_9_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_7_49 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_7) {
  $__num = \func_num_args();
  $__res = $x_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_7) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_8_49 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_8_49 = (object)["map" => function($f_9) use ($__local_var_8_49) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($__local_var_8_49, $f_9) {
  $__num = \func_num_args();
  $__res = ((($__local_var_8_49)->{'map'})(function($v1_11) use ($f_9) {
  $__num = \func_num_args();
  $__t50 = null;;
  if ($v1_11 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t50 = new \Data\Maybe\Data_Maybe_Just(($f_9)(($v1_11)->{'value0'}));
goto end_branch_50;;
};
  $__t50 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_50:;
  $__res = $__t50;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_10);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_9_52 = (($__local_var_2_1)->{'Bind1'})(null);
  $Applicative0_10_53 = (($__local_var_2_1)->{'Applicative0'})(null);
  $Bind1_9_52 = (object)["bind" => function($v_11) use ($Applicative0_10_53, $Bind1_9_52) {
  $__num = \func_num_args();
  $__res = function($f_12) use ($Applicative0_10_53, $Bind1_9_52, $v_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_52)->{'bind'})($v_11))(function($v1_13) use ($Applicative0_10_53, $f_12) {
  $__num = \func_num_args();
  $__t54 = null;;
  if ($v1_13 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t54 = (($Applicative0_10_53)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_54;;
};
  if ($v1_13 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t54 = ($f_12)(($v1_13)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_11) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_12_55 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_12_55 = (object)["map" => function($f_13) use ($__local_var_12_55) {
  $__num = \func_num_args();
  $__res = function($v_14) use ($__local_var_12_55, $f_13) {
  $__num = \func_num_args();
  $__res = ((($__local_var_12_55)->{'map'})(function($v1_15) use ($f_13) {
  $__num = \func_num_args();
  $__t56 = null;;
  if ($v1_15 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t56 = new \Data\Maybe\Data_Maybe_Just(($f_13)(($v1_15)->{'value0'}));
goto end_branch_56;;
};
  $__t56 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_56:;
  $__res = $__t56;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_14);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_13_58 = (($__local_var_2_1)->{'Bind1'})(null);
  $Applicative0_14_59 = (($__local_var_2_1)->{'Applicative0'})(null);
  $Bind1_13_58 = (object)["bind" => function($v_15) use ($Applicative0_14_59, $Bind1_13_58) {
  $__num = \func_num_args();
  $__res = function($f_16) use ($Applicative0_14_59, $Bind1_13_58, $v_15) {
  $__num = \func_num_args();
  $__res = ((($Bind1_13_58)->{'bind'})($v_15))(function($v1_17) use ($Applicative0_14_59, $f_16) {
  $__num = \func_num_args();
  $__t60 = null;;
  if ($v1_17 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t60 = (($Applicative0_14_59)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_60;;
};
  if ($v1_17 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t60 = ($f_16)(($v1_17)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_15) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($__local_var_2_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_14_62 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($__local_var_2_1);
  $__res = (object)["apply" => function($f_15) use ($Applicative0_14_62, $Bind1_13_58) {
  $__num = \func_num_args();
  $__res = function($a_16) use ($Applicative0_14_62, $Bind1_13_58, $f_15) {
  $__num = \func_num_args();
  $__res = ((($Bind1_13_58)->{'bind'})($f_15))(function($f_prime__17) use ($Applicative0_14_62, $Bind1_13_58, $a_16) {
  $__num = \func_num_args();
  $__res = ((($Bind1_13_58)->{'bind'})($a_16))(function($a_prime__18) use ($Applicative0_14_62, $f_prime__17) {
  $__num = \func_num_args();
  $__res = (($Applicative0_14_62)->{'pure'})(($f_prime__17)($a_prime__18));
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
}, "Functor0" => function($_dollar___unused_13) use ($functorMaybeT1_12_55) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_12_55;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_10_64 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($__local_var_2_1);
  $__res = (object)["apply" => function($f_11) use ($Applicative0_10_64, $Bind1_9_52) {
  $__num = \func_num_args();
  $__res = function($a_12) use ($Applicative0_10_64, $Bind1_9_52, $f_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_52)->{'bind'})($f_11))(function($f_prime__13) use ($Applicative0_10_64, $Bind1_9_52, $a_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_52)->{'bind'})($a_12))(function($a_prime__14) use ($Applicative0_10_64, $f_prime__13) {
  $__num = \func_num_args();
  $__res = (($Applicative0_10_64)->{'pure'})(($f_prime__13)($a_prime__14));
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
}, "Functor0" => function($_dollar___unused_9) use ($functorMaybeT1_8_49) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_8_49;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_8) use ($Applicative0_7_49, $Bind1_6_5) {
  $__num = \func_num_args();
  $__res = function($a_9) use ($Applicative0_7_49, $Bind1_6_5, $f_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_5)->{'bind'})($f_8))(function($f_prime__10) use ($Applicative0_7_49, $Bind1_6_5, $a_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_5)->{'bind'})($a_9))(function($a_prime__11) use ($Applicative0_7_49, $f_prime__10) {
  $__num = \func_num_args();
  $__res = (($Applicative0_7_49)->{'pure'})(($f_prime__10)($a_prime__11));
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
}, "Functor0" => function($_dollar___unused_6) use ($functorMaybeT1_5_2) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_5_2;
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
  $Bind1_4_66 = (($__local_var_2_1)->{'Bind1'})(null);
  $Applicative0_5_67 = (($__local_var_2_1)->{'Applicative0'})(null);
  $__res = (object)["bind" => function($v_6) use ($Applicative0_5_67, $Bind1_4_66) {
  $__num = \func_num_args();
  $__res = function($f_7) use ($Applicative0_5_67, $Bind1_4_66, $v_6) {
  $__num = \func_num_args();
  $__res = ((($Bind1_4_66)->{'bind'})($v_6))(function($v1_8) use ($Applicative0_5_67, $f_7) {
  $__num = \func_num_args();
  $__t68 = null;;
  if ($v1_8 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t68 = (($Applicative0_5_67)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_68;;
};
  if ($v1_8 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t68 = ($f_7)(($v1_8)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_6) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_7_69 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_7_69 = (object)["map" => function($f_8) use ($__local_var_7_69) {
  $__num = \func_num_args();
  $__res = function($v_9) use ($__local_var_7_69, $f_8) {
  $__num = \func_num_args();
  $__res = ((($__local_var_7_69)->{'map'})(function($v1_10) use ($f_8) {
  $__num = \func_num_args();
  $__t70 = null;;
  if ($v1_10 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t70 = new \Data\Maybe\Data_Maybe_Just(($f_8)(($v1_10)->{'value0'}));
goto end_branch_70;;
};
  $__t70 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_70:;
  $__res = $__t70;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_8_72 = (($__local_var_2_1)->{'Bind1'})(null);
  $Applicative0_9_73 = (($__local_var_2_1)->{'Applicative0'})(null);
  $Bind1_8_72 = (object)["bind" => function($v_10) use ($Applicative0_9_73, $Bind1_8_72) {
  $__num = \func_num_args();
  $__res = function($f_11) use ($Applicative0_9_73, $Bind1_8_72, $v_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_8_72)->{'bind'})($v_10))(function($v1_12) use ($Applicative0_9_73, $f_11) {
  $__num = \func_num_args();
  $__t74 = null;;
  if ($v1_12 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t74 = (($Applicative0_9_73)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_74;;
};
  if ($v1_12 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t74 = ($f_11)(($v1_12)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_10) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_11_75 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_11_75 = (object)["map" => function($f_12) use ($__local_var_11_75) {
  $__num = \func_num_args();
  $__res = function($v_13) use ($__local_var_11_75, $f_12) {
  $__num = \func_num_args();
  $__res = ((($__local_var_11_75)->{'map'})(function($v1_14) use ($f_12) {
  $__num = \func_num_args();
  $__t76 = null;;
  if ($v1_14 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t76 = new \Data\Maybe\Data_Maybe_Just(($f_12)(($v1_14)->{'value0'}));
goto end_branch_76;;
};
  $__t76 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_76:;
  $__res = $__t76;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_13);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_12_78 = (($__local_var_2_1)->{'Bind1'})(null);
  $Applicative0_13_79 = (($__local_var_2_1)->{'Applicative0'})(null);
  $Bind1_12_78 = (object)["bind" => function($v_14) use ($Applicative0_13_79, $Bind1_12_78) {
  $__num = \func_num_args();
  $__res = function($f_15) use ($Applicative0_13_79, $Bind1_12_78, $v_14) {
  $__num = \func_num_args();
  $__res = ((($Bind1_12_78)->{'bind'})($v_14))(function($v1_16) use ($Applicative0_13_79, $f_15) {
  $__num = \func_num_args();
  $__t80 = null;;
  if ($v1_16 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t80 = (($Applicative0_13_79)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_80;;
};
  if ($v1_16 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t80 = ($f_15)(($v1_16)->{'value0'});
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_14) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($__local_var_2_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_13_82 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_13) {
  $__num = \func_num_args();
  $__res = $x_13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_13) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_14_82 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_14_82 = (object)["map" => function($f_15) use ($__local_var_14_82) {
  $__num = \func_num_args();
  $__res = function($v_16) use ($__local_var_14_82, $f_15) {
  $__num = \func_num_args();
  $__res = ((($__local_var_14_82)->{'map'})(function($v1_17) use ($f_15) {
  $__num = \func_num_args();
  $__t83 = null;;
  if ($v1_17 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t83 = new \Data\Maybe\Data_Maybe_Just(($f_15)(($v1_17)->{'value0'}));
goto end_branch_83;;
};
  $__t83 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_83:;
  $__res = $__t83;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_16);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_15_85 = (($__local_var_2_1)->{'Bind1'})(null);
  $Applicative0_16_86 = (($__local_var_2_1)->{'Applicative0'})(null);
  $Bind1_15_85 = (object)["bind" => function($v_17) use ($Applicative0_16_86, $Bind1_15_85) {
  $__num = \func_num_args();
  $__res = function($f_18) use ($Applicative0_16_86, $Bind1_15_85, $v_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_85)->{'bind'})($v_17))(function($v1_19) use ($Applicative0_16_86, $f_18) {
  $__num = \func_num_args();
  $__t87 = null;;
  if ($v1_19 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t87 = (($Applicative0_16_86)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_87;;
};
  if ($v1_19 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t87 = ($f_18)(($v1_19)->{'value0'});
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_17) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($__local_var_2_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_16_89 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($__local_var_2_1);
  $__res = (object)["apply" => function($f_17) use ($Applicative0_16_89, $Bind1_15_85) {
  $__num = \func_num_args();
  $__res = function($a_18) use ($Applicative0_16_89, $Bind1_15_85, $f_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_85)->{'bind'})($f_17))(function($f_prime__19) use ($Applicative0_16_89, $Bind1_15_85, $a_18) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_85)->{'bind'})($a_18))(function($a_prime__20) use ($Applicative0_16_89, $f_prime__19) {
  $__num = \func_num_args();
  $__res = (($Applicative0_16_89)->{'pure'})(($f_prime__19)($a_prime__20));
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
}, "Functor0" => function($_dollar___unused_15) use ($functorMaybeT1_14_82) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_14_82;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_14) use ($Applicative0_13_82, $Bind1_12_78) {
  $__num = \func_num_args();
  $__res = function($a_15) use ($Applicative0_13_82, $Bind1_12_78, $f_14) {
  $__num = \func_num_args();
  $__res = ((($Bind1_12_78)->{'bind'})($f_14))(function($f_prime__16) use ($Applicative0_13_82, $Bind1_12_78, $a_15) {
  $__num = \func_num_args();
  $__res = ((($Bind1_12_78)->{'bind'})($a_15))(function($a_prime__17) use ($Applicative0_13_82, $f_prime__16) {
  $__num = \func_num_args();
  $__res = (($Applicative0_13_82)->{'pure'})(($f_prime__16)($a_prime__17));
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
}, "Functor0" => function($_dollar___unused_12) use ($functorMaybeT1_11_75) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_11_75;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_9_92 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_9) {
  $__num = \func_num_args();
  $__res = $x_9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_9) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_10_92 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_10_92 = (object)["map" => function($f_11) use ($__local_var_10_92) {
  $__num = \func_num_args();
  $__res = function($v_12) use ($__local_var_10_92, $f_11) {
  $__num = \func_num_args();
  $__res = ((($__local_var_10_92)->{'map'})(function($v1_13) use ($f_11) {
  $__num = \func_num_args();
  $__t93 = null;;
  if ($v1_13 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t93 = new \Data\Maybe\Data_Maybe_Just(($f_11)(($v1_13)->{'value0'}));
goto end_branch_93;;
};
  $__t93 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_93:;
  $__res = $__t93;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_12);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_11_95 = (($__local_var_2_1)->{'Bind1'})(null);
  $Applicative0_12_96 = (($__local_var_2_1)->{'Applicative0'})(null);
  $Bind1_11_95 = (object)["bind" => function($v_13) use ($Applicative0_12_96, $Bind1_11_95) {
  $__num = \func_num_args();
  $__res = function($f_14) use ($Applicative0_12_96, $Bind1_11_95, $v_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_95)->{'bind'})($v_13))(function($v1_15) use ($Applicative0_12_96, $f_14) {
  $__num = \func_num_args();
  $__t97 = null;;
  if ($v1_15 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t97 = (($Applicative0_12_96)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_97;;
};
  if ($v1_15 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t97 = ($f_14)(($v1_15)->{'value0'});
goto end_branch_97;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t97 = null;
  end_branch_97:;
  $__res = $__t97;
  goto __end;;
  __end:
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
  $__local_var_14_98 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_14_98 = (object)["map" => function($f_15) use ($__local_var_14_98) {
  $__num = \func_num_args();
  $__res = function($v_16) use ($__local_var_14_98, $f_15) {
  $__num = \func_num_args();
  $__res = ((($__local_var_14_98)->{'map'})(function($v1_17) use ($f_15) {
  $__num = \func_num_args();
  $__t99 = null;;
  if ($v1_17 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t99 = new \Data\Maybe\Data_Maybe_Just(($f_15)(($v1_17)->{'value0'}));
goto end_branch_99;;
};
  $__t99 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_99:;
  $__res = $__t99;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_16);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_15_101 = (($__local_var_2_1)->{'Bind1'})(null);
  $Applicative0_16_102 = (($__local_var_2_1)->{'Applicative0'})(null);
  $Bind1_15_101 = (object)["bind" => function($v_17) use ($Applicative0_16_102, $Bind1_15_101) {
  $__num = \func_num_args();
  $__res = function($f_18) use ($Applicative0_16_102, $Bind1_15_101, $v_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_101)->{'bind'})($v_17))(function($v1_19) use ($Applicative0_16_102, $f_18) {
  $__num = \func_num_args();
  $__t103 = null;;
  if ($v1_19 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t103 = (($Applicative0_16_102)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_103;;
};
  if ($v1_19 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t103 = ($f_18)(($v1_19)->{'value0'});
goto end_branch_103;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t103 = null;
  end_branch_103:;
  $__res = $__t103;
  goto __end;;
  __end:
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
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($__local_var_2_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_16_105 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_16) {
  $__num = \func_num_args();
  $__res = $x_16;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_16) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_17_105 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_17_105 = (object)["map" => function($f_18) use ($__local_var_17_105) {
  $__num = \func_num_args();
  $__res = function($v_19) use ($__local_var_17_105, $f_18) {
  $__num = \func_num_args();
  $__res = ((($__local_var_17_105)->{'map'})(function($v1_20) use ($f_18) {
  $__num = \func_num_args();
  $__t106 = null;;
  if ($v1_20 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t106 = new \Data\Maybe\Data_Maybe_Just(($f_18)(($v1_20)->{'value0'}));
goto end_branch_106;;
};
  $__t106 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_106:;
  $__res = $__t106;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_19);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_18_108 = (($__local_var_2_1)->{'Bind1'})(null);
  $Applicative0_19_109 = (($__local_var_2_1)->{'Applicative0'})(null);
  $Bind1_18_108 = (object)["bind" => function($v_20) use ($Applicative0_19_109, $Bind1_18_108) {
  $__num = \func_num_args();
  $__res = function($f_21) use ($Applicative0_19_109, $Bind1_18_108, $v_20) {
  $__num = \func_num_args();
  $__res = ((($Bind1_18_108)->{'bind'})($v_20))(function($v1_22) use ($Applicative0_19_109, $f_21) {
  $__num = \func_num_args();
  $__t110 = null;;
  if ($v1_22 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t110 = (($Applicative0_19_109)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_110;;
};
  if ($v1_22 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t110 = ($f_21)(($v1_22)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_20) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($__local_var_2_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_19_112 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($__local_var_2_1);
  $__res = (object)["apply" => function($f_20) use ($Applicative0_19_112, $Bind1_18_108) {
  $__num = \func_num_args();
  $__res = function($a_21) use ($Applicative0_19_112, $Bind1_18_108, $f_20) {
  $__num = \func_num_args();
  $__res = ((($Bind1_18_108)->{'bind'})($f_20))(function($f_prime__22) use ($Applicative0_19_112, $Bind1_18_108, $a_21) {
  $__num = \func_num_args();
  $__res = ((($Bind1_18_108)->{'bind'})($a_21))(function($a_prime__23) use ($Applicative0_19_112, $f_prime__22) {
  $__num = \func_num_args();
  $__res = (($Applicative0_19_112)->{'pure'})(($f_prime__22)($a_prime__23));
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
}, "Functor0" => function($_dollar___unused_18) use ($functorMaybeT1_17_105) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_17_105;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_17) use ($Applicative0_16_105, $Bind1_15_101) {
  $__num = \func_num_args();
  $__res = function($a_18) use ($Applicative0_16_105, $Bind1_15_101, $f_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_101)->{'bind'})($f_17))(function($f_prime__19) use ($Applicative0_16_105, $Bind1_15_101, $a_18) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_101)->{'bind'})($a_18))(function($a_prime__20) use ($Applicative0_16_105, $f_prime__19) {
  $__num = \func_num_args();
  $__res = (($Applicative0_16_105)->{'pure'})(($f_prime__19)($a_prime__20));
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
}, "Functor0" => function($_dollar___unused_15) use ($functorMaybeT1_14_98) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_14_98;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_12_115 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_12) {
  $__num = \func_num_args();
  $__res = $x_12;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_12) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_13_115 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_13_115 = (object)["map" => function($f_14) use ($__local_var_13_115) {
  $__num = \func_num_args();
  $__res = function($v_15) use ($__local_var_13_115, $f_14) {
  $__num = \func_num_args();
  $__res = ((($__local_var_13_115)->{'map'})(function($v1_16) use ($f_14) {
  $__num = \func_num_args();
  $__t116 = null;;
  if ($v1_16 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t116 = new \Data\Maybe\Data_Maybe_Just(($f_14)(($v1_16)->{'value0'}));
goto end_branch_116;;
};
  $__t116 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_116:;
  $__res = $__t116;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_15);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_14_118 = (($__local_var_2_1)->{'Bind1'})(null);
  $Applicative0_15_119 = (($__local_var_2_1)->{'Applicative0'})(null);
  $Bind1_14_118 = (object)["bind" => function($v_16) use ($Applicative0_15_119, $Bind1_14_118) {
  $__num = \func_num_args();
  $__res = function($f_17) use ($Applicative0_15_119, $Bind1_14_118, $v_16) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_118)->{'bind'})($v_16))(function($v1_18) use ($Applicative0_15_119, $f_17) {
  $__num = \func_num_args();
  $__t120 = null;;
  if ($v1_18 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t120 = (($Applicative0_15_119)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_120;;
};
  if ($v1_18 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t120 = ($f_17)(($v1_18)->{'value0'});
goto end_branch_120;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t120 = null;
  end_branch_120:;
  $__res = $__t120;
  goto __end;;
  __end:
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
  $__local_var_17_121 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_17_121 = (object)["map" => function($f_18) use ($__local_var_17_121) {
  $__num = \func_num_args();
  $__res = function($v_19) use ($__local_var_17_121, $f_18) {
  $__num = \func_num_args();
  $__res = ((($__local_var_17_121)->{'map'})(function($v1_20) use ($f_18) {
  $__num = \func_num_args();
  $__t122 = null;;
  if ($v1_20 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t122 = new \Data\Maybe\Data_Maybe_Just(($f_18)(($v1_20)->{'value0'}));
goto end_branch_122;;
};
  $__t122 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_122:;
  $__res = $__t122;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_19);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_18_124 = (($__local_var_2_1)->{'Bind1'})(null);
  $Applicative0_19_125 = (($__local_var_2_1)->{'Applicative0'})(null);
  $Bind1_18_124 = (object)["bind" => function($v_20) use ($Applicative0_19_125, $Bind1_18_124) {
  $__num = \func_num_args();
  $__res = function($f_21) use ($Applicative0_19_125, $Bind1_18_124, $v_20) {
  $__num = \func_num_args();
  $__res = ((($Bind1_18_124)->{'bind'})($v_20))(function($v1_22) use ($Applicative0_19_125, $f_21) {
  $__num = \func_num_args();
  $__t126 = null;;
  if ($v1_22 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t126 = (($Applicative0_19_125)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_126;;
};
  if ($v1_22 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t126 = ($f_21)(($v1_22)->{'value0'});
goto end_branch_126;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t126 = null;
  end_branch_126:;
  $__res = $__t126;
  goto __end;;
  __end:
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
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($__local_var_2_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_19_128 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($__local_var_2_1);
  $__res = (object)["apply" => function($f_20) use ($Applicative0_19_128, $Bind1_18_124) {
  $__num = \func_num_args();
  $__res = function($a_21) use ($Applicative0_19_128, $Bind1_18_124, $f_20) {
  $__num = \func_num_args();
  $__res = ((($Bind1_18_124)->{'bind'})($f_20))(function($f_prime__22) use ($Applicative0_19_128, $Bind1_18_124, $a_21) {
  $__num = \func_num_args();
  $__res = ((($Bind1_18_124)->{'bind'})($a_21))(function($a_prime__23) use ($Applicative0_19_128, $f_prime__22) {
  $__num = \func_num_args();
  $__res = (($Applicative0_19_128)->{'pure'})(($f_prime__22)($a_prime__23));
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
}, "Functor0" => function($_dollar___unused_18) use ($functorMaybeT1_17_121) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_17_121;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_15_130 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($__local_var_2_1);
  $__res = (object)["apply" => function($f_16) use ($Applicative0_15_130, $Bind1_14_118) {
  $__num = \func_num_args();
  $__res = function($a_17) use ($Applicative0_15_130, $Bind1_14_118, $f_16) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_118)->{'bind'})($f_16))(function($f_prime__18) use ($Applicative0_15_130, $Bind1_14_118, $a_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_118)->{'bind'})($a_17))(function($a_prime__19) use ($Applicative0_15_130, $f_prime__18) {
  $__num = \func_num_args();
  $__res = (($Applicative0_15_130)->{'pure'})(($f_prime__18)($a_prime__19));
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
}, "Functor0" => function($_dollar___unused_14) use ($functorMaybeT1_13_115) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_13_115;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_13) use ($Applicative0_12_115, $Bind1_11_95) {
  $__num = \func_num_args();
  $__res = function($a_14) use ($Applicative0_12_115, $Bind1_11_95, $f_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_95)->{'bind'})($f_13))(function($f_prime__15) use ($Applicative0_12_115, $Bind1_11_95, $a_14) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_95)->{'bind'})($a_14))(function($a_prime__16) use ($Applicative0_12_115, $f_prime__15) {
  $__num = \func_num_args();
  $__res = (($Applicative0_12_115)->{'pure'})(($f_prime__15)($a_prime__16));
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
}, "Functor0" => function($_dollar___unused_11) use ($functorMaybeT1_10_92) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_10_92;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_10) use ($Applicative0_9_92, $Bind1_8_72) {
  $__num = \func_num_args();
  $__res = function($a_11) use ($Applicative0_9_92, $Bind1_8_72, $f_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_8_72)->{'bind'})($f_10))(function($f_prime__12) use ($Applicative0_9_92, $Bind1_8_72, $a_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_8_72)->{'bind'})($a_11))(function($a_prime__13) use ($Applicative0_9_92, $f_prime__12) {
  $__num = \func_num_args();
  $__res = (($Applicative0_9_92)->{'pure'})(($f_prime__12)($a_prime__13));
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
}, "Functor0" => function($_dollar___unused_8) use ($functorMaybeT1_7_69) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_7_69;
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
  $__local_var_3_134 = (($__local_var_1_0)->{'Monad0'})(null);
  $Bind1_4_135 = (($__local_var_3_134)->{'Bind1'})(null);
  $Applicative0_5_136 = (($__local_var_3_134)->{'Applicative0'})(null);
  $monadAskMaybeT1_1_0 = (object)["ask" => \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($x_4) {
  $__num = \func_num_args();
  $__res = $x_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, function($a_6) use ($Applicative0_5_136, $Bind1_4_135) {
  $__num = \func_num_args();
  $__res = ((($Bind1_4_135)->{'bind'})($a_6))(function($a_prime__7) use ($Applicative0_5_136) {
  $__num = \func_num_args();
  $__res = (($Applicative0_5_136)->{'pure'})(new \Data\Maybe\Data_Maybe_Just($a_prime__7));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($__local_var_1_0)->{'ask'}), "Monad0" => function($_dollar___unused_3) use ($monadMaybeT1_2_1) {
  $__num = \func_num_args();
  $__res = $monadMaybeT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["local" => function($f_2) use ($dictMonadReader_0) {
  $__num = \func_num_args();
  $__local_var_3_138 = (($dictMonadReader_0)->{'local'})($f_2);
  $__res = function($v_4) use ($__local_var_3_138) {
  $__num = \func_num_args();
  $__res = ($__local_var_3_138)($v_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "MonadAsk0" => function($_dollar___unused_2) use ($monadAskMaybeT1_1_0) {
  $__num = \func_num_args();
  $__res = $monadAskMaybeT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Maybe_Trans_monadReaderMaybeT'] = __NAMESPACE__ . '\\majControl_majMonad_majMaybe_majTrans_monadmajReadermajMaybemajT';

// Control_Monad_Maybe_Trans_monadContMaybeT
function majControl_majMonad_majMaybe_majTrans_monadmajContmajMaybemajT($dictMonadCont_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majMaybe_majTrans_monadmajContmajMaybemajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictMonadCont_0)->{'Monad0'})(null);
  $monadMaybeT1_1_0 = (object)["Applicative0" => function($_dollar___unused_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_3) {
  $__num = \func_num_args();
  $__res = $x_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($__local_var_1_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_3) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__local_var_4_1 = (((((($__local_var_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_4_1 = (object)["map" => function($f_5) use ($__local_var_4_1) {
  $__num = \func_num_args();
  $__res = function($v_6) use ($__local_var_4_1, $f_5) {
  $__num = \func_num_args();
  $__res = ((($__local_var_4_1)->{'map'})(function($v1_7) use ($f_5) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ($v1_7 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t2 = new \Data\Maybe\Data_Maybe_Just(($f_5)(($v1_7)->{'value0'}));
goto end_branch_2;;
};
  $__t2 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_5_4 = (($__local_var_1_0)->{'Bind1'})(null);
  $Applicative0_6_5 = (($__local_var_1_0)->{'Applicative0'})(null);
  $Bind1_5_4 = (object)["bind" => function($v_7) use ($Applicative0_6_5, $Bind1_5_4) {
  $__num = \func_num_args();
  $__res = function($f_8) use ($Applicative0_6_5, $Bind1_5_4, $v_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_5_4)->{'bind'})($v_7))(function($v1_9) use ($Applicative0_6_5, $f_8) {
  $__num = \func_num_args();
  $__t6 = null;;
  if ($v1_9 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t6 = (($Applicative0_6_5)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_6;;
};
  if ($v1_9 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t6 = ($f_8)(($v1_9)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_7) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__local_var_8_7 = (((((($__local_var_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_8_7 = (object)["map" => function($f_9) use ($__local_var_8_7) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($__local_var_8_7, $f_9) {
  $__num = \func_num_args();
  $__res = ((($__local_var_8_7)->{'map'})(function($v1_11) use ($f_9) {
  $__num = \func_num_args();
  $__t8 = null;;
  if ($v1_11 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t8 = new \Data\Maybe\Data_Maybe_Just(($f_9)(($v1_11)->{'value0'}));
goto end_branch_8;;
};
  $__t8 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_8:;
  $__res = $__t8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_10);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_9_10 = (($__local_var_1_0)->{'Bind1'})(null);
  $Applicative0_10_11 = (($__local_var_1_0)->{'Applicative0'})(null);
  $Bind1_9_10 = (object)["bind" => function($v_11) use ($Applicative0_10_11, $Bind1_9_10) {
  $__num = \func_num_args();
  $__res = function($f_12) use ($Applicative0_10_11, $Bind1_9_10, $v_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_10)->{'bind'})($v_11))(function($v1_13) use ($Applicative0_10_11, $f_12) {
  $__num = \func_num_args();
  $__t12 = null;;
  if ($v1_13 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t12 = (($Applicative0_10_11)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_12;;
};
  if ($v1_13 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t12 = ($f_12)(($v1_13)->{'value0'});
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_11) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($__local_var_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_10_14 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($__local_var_1_0);
  $__res = (object)["apply" => function($f_11) use ($Applicative0_10_14, $Bind1_9_10) {
  $__num = \func_num_args();
  $__res = function($a_12) use ($Applicative0_10_14, $Bind1_9_10, $f_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_10)->{'bind'})($f_11))(function($f_prime__13) use ($Applicative0_10_14, $Bind1_9_10, $a_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_10)->{'bind'})($a_12))(function($a_prime__14) use ($Applicative0_10_14, $f_prime__13) {
  $__num = \func_num_args();
  $__res = (($Applicative0_10_14)->{'pure'})(($f_prime__13)($a_prime__14));
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
}, "Functor0" => function($_dollar___unused_9) use ($functorMaybeT1_8_7) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_8_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_6_16 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($__local_var_1_0);
  $__res = (object)["apply" => function($f_7) use ($Applicative0_6_16, $Bind1_5_4) {
  $__num = \func_num_args();
  $__res = function($a_8) use ($Applicative0_6_16, $Bind1_5_4, $f_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_5_4)->{'bind'})($f_7))(function($f_prime__9) use ($Applicative0_6_16, $Bind1_5_4, $a_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_5_4)->{'bind'})($a_8))(function($a_prime__10) use ($Applicative0_6_16, $f_prime__9) {
  $__num = \func_num_args();
  $__res = (($Applicative0_6_16)->{'pure'})(($f_prime__9)($a_prime__10));
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
}, "Functor0" => function($_dollar___unused_5) use ($functorMaybeT1_4_1) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_4_1;
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
  $Bind1_3_17 = (($__local_var_1_0)->{'Bind1'})(null);
  $Applicative0_4_18 = (($__local_var_1_0)->{'Applicative0'})(null);
  $__res = (object)["bind" => function($v_5) use ($Applicative0_4_18, $Bind1_3_17) {
  $__num = \func_num_args();
  $__res = function($f_6) use ($Applicative0_4_18, $Bind1_3_17, $v_5) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_17)->{'bind'})($v_5))(function($v1_7) use ($Applicative0_4_18, $f_6) {
  $__num = \func_num_args();
  $__t19 = null;;
  if ($v1_7 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t19 = (($Applicative0_4_18)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_19;;
};
  if ($v1_7 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t19 = ($f_6)(($v1_7)->{'value0'});
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_5) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__local_var_6_20 = (((((($__local_var_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_6_20 = (object)["map" => function($f_7) use ($__local_var_6_20) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($__local_var_6_20, $f_7) {
  $__num = \func_num_args();
  $__res = ((($__local_var_6_20)->{'map'})(function($v1_9) use ($f_7) {
  $__num = \func_num_args();
  $__t21 = null;;
  if ($v1_9 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t21 = new \Data\Maybe\Data_Maybe_Just(($f_7)(($v1_9)->{'value0'}));
goto end_branch_21;;
};
  $__t21 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_21:;
  $__res = $__t21;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_8);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_7_23 = (($__local_var_1_0)->{'Bind1'})(null);
  $Applicative0_8_24 = (($__local_var_1_0)->{'Applicative0'})(null);
  $Bind1_7_23 = (object)["bind" => function($v_9) use ($Applicative0_8_24, $Bind1_7_23) {
  $__num = \func_num_args();
  $__res = function($f_10) use ($Applicative0_8_24, $Bind1_7_23, $v_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_23)->{'bind'})($v_9))(function($v1_11) use ($Applicative0_8_24, $f_10) {
  $__num = \func_num_args();
  $__t25 = null;;
  if ($v1_11 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t25 = (($Applicative0_8_24)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_25;;
};
  if ($v1_11 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t25 = ($f_10)(($v1_11)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_9) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($__local_var_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_8_27 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_8) {
  $__num = \func_num_args();
  $__res = $x_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($__local_var_1_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_8) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__local_var_9_27 = (((((($__local_var_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_9_27 = (object)["map" => function($f_10) use ($__local_var_9_27) {
  $__num = \func_num_args();
  $__res = function($v_11) use ($__local_var_9_27, $f_10) {
  $__num = \func_num_args();
  $__res = ((($__local_var_9_27)->{'map'})(function($v1_12) use ($f_10) {
  $__num = \func_num_args();
  $__t28 = null;;
  if ($v1_12 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t28 = new \Data\Maybe\Data_Maybe_Just(($f_10)(($v1_12)->{'value0'}));
goto end_branch_28;;
};
  $__t28 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_28:;
  $__res = $__t28;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_11);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_10_30 = (($__local_var_1_0)->{'Bind1'})(null);
  $Applicative0_11_31 = (($__local_var_1_0)->{'Applicative0'})(null);
  $Bind1_10_30 = (object)["bind" => function($v_12) use ($Applicative0_11_31, $Bind1_10_30) {
  $__num = \func_num_args();
  $__res = function($f_13) use ($Applicative0_11_31, $Bind1_10_30, $v_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_30)->{'bind'})($v_12))(function($v1_14) use ($Applicative0_11_31, $f_13) {
  $__num = \func_num_args();
  $__t32 = null;;
  if ($v1_14 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t32 = (($Applicative0_11_31)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_32;;
};
  if ($v1_14 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t32 = ($f_13)(($v1_14)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_12) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($__local_var_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_11_34 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($__local_var_1_0);
  $__res = (object)["apply" => function($f_12) use ($Applicative0_11_34, $Bind1_10_30) {
  $__num = \func_num_args();
  $__res = function($a_13) use ($Applicative0_11_34, $Bind1_10_30, $f_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_30)->{'bind'})($f_12))(function($f_prime__14) use ($Applicative0_11_34, $Bind1_10_30, $a_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_30)->{'bind'})($a_13))(function($a_prime__15) use ($Applicative0_11_34, $f_prime__14) {
  $__num = \func_num_args();
  $__res = (($Applicative0_11_34)->{'pure'})(($f_prime__14)($a_prime__15));
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
}, "Functor0" => function($_dollar___unused_10) use ($functorMaybeT1_9_27) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_9_27;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_9) use ($Applicative0_8_27, $Bind1_7_23) {
  $__num = \func_num_args();
  $__res = function($a_10) use ($Applicative0_8_27, $Bind1_7_23, $f_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_23)->{'bind'})($f_9))(function($f_prime__11) use ($Applicative0_8_27, $Bind1_7_23, $a_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_23)->{'bind'})($a_10))(function($a_prime__12) use ($Applicative0_8_27, $f_prime__11) {
  $__num = \func_num_args();
  $__res = (($Applicative0_8_27)->{'pure'})(($f_prime__11)($a_prime__12));
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
}, "Functor0" => function($_dollar___unused_7) use ($functorMaybeT1_6_20) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_6_20;
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
  $__res = ($c_3)(new \Data\Maybe\Data_Maybe_Just($a_4));
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
}, "Monad0" => function($_dollar___unused_2) use ($monadMaybeT1_1_0) {
  $__num = \func_num_args();
  $__res = $monadMaybeT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Maybe_Trans_monadContMaybeT'] = __NAMESPACE__ . '\\majControl_majMonad_majMaybe_majTrans_monadmajContmajMaybemajT';

// Control_Monad_Maybe_Trans_monadEffectMaybe
function majControl_majMonad_majMaybe_majTrans_monadmajEffectmajMaybe($dictMonadEffect_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majMaybe_majTrans_monadmajEffectmajMaybe';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Monad0_1_0 = (($dictMonadEffect_0)->{'Monad0'})(null);
  $monadMaybeT1_2_1 = (object)["Applicative0" => function($_dollar___unused_2) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_3) {
  $__num = \func_num_args();
  $__res = $x_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($Monad0_1_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_3) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__local_var_4_1 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_4_1 = (object)["map" => function($f_5) use ($__local_var_4_1) {
  $__num = \func_num_args();
  $__res = function($v_6) use ($__local_var_4_1, $f_5) {
  $__num = \func_num_args();
  $__res = ((($__local_var_4_1)->{'map'})(function($v1_7) use ($f_5) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ($v1_7 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t2 = new \Data\Maybe\Data_Maybe_Just(($f_5)(($v1_7)->{'value0'}));
goto end_branch_2;;
};
  $__t2 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_5_4 = (($Monad0_1_0)->{'Bind1'})(null);
  $Applicative0_6_5 = (($Monad0_1_0)->{'Applicative0'})(null);
  $Bind1_5_4 = (object)["bind" => function($v_7) use ($Applicative0_6_5, $Bind1_5_4) {
  $__num = \func_num_args();
  $__res = function($f_8) use ($Applicative0_6_5, $Bind1_5_4, $v_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_5_4)->{'bind'})($v_7))(function($v1_9) use ($Applicative0_6_5, $f_8) {
  $__num = \func_num_args();
  $__t6 = null;;
  if ($v1_9 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t6 = (($Applicative0_6_5)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_6;;
};
  if ($v1_9 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t6 = ($f_8)(($v1_9)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_7) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__local_var_8_7 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_8_7 = (object)["map" => function($f_9) use ($__local_var_8_7) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($__local_var_8_7, $f_9) {
  $__num = \func_num_args();
  $__res = ((($__local_var_8_7)->{'map'})(function($v1_11) use ($f_9) {
  $__num = \func_num_args();
  $__t8 = null;;
  if ($v1_11 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t8 = new \Data\Maybe\Data_Maybe_Just(($f_9)(($v1_11)->{'value0'}));
goto end_branch_8;;
};
  $__t8 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_8:;
  $__res = $__t8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_10);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_9_10 = (($Monad0_1_0)->{'Bind1'})(null);
  $Applicative0_10_11 = (($Monad0_1_0)->{'Applicative0'})(null);
  $Bind1_9_10 = (object)["bind" => function($v_11) use ($Applicative0_10_11, $Bind1_9_10) {
  $__num = \func_num_args();
  $__res = function($f_12) use ($Applicative0_10_11, $Bind1_9_10, $v_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_10)->{'bind'})($v_11))(function($v1_13) use ($Applicative0_10_11, $f_12) {
  $__num = \func_num_args();
  $__t12 = null;;
  if ($v1_13 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t12 = (($Applicative0_10_11)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_12;;
};
  if ($v1_13 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t12 = ($f_12)(($v1_13)->{'value0'});
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_11) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_10_14 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($Monad0_1_0);
  $__res = (object)["apply" => function($f_11) use ($Applicative0_10_14, $Bind1_9_10) {
  $__num = \func_num_args();
  $__res = function($a_12) use ($Applicative0_10_14, $Bind1_9_10, $f_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_10)->{'bind'})($f_11))(function($f_prime__13) use ($Applicative0_10_14, $Bind1_9_10, $a_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_10)->{'bind'})($a_12))(function($a_prime__14) use ($Applicative0_10_14, $f_prime__13) {
  $__num = \func_num_args();
  $__res = (($Applicative0_10_14)->{'pure'})(($f_prime__13)($a_prime__14));
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
}, "Functor0" => function($_dollar___unused_9) use ($functorMaybeT1_8_7) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_8_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_6_16 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($Monad0_1_0);
  $__res = (object)["apply" => function($f_7) use ($Applicative0_6_16, $Bind1_5_4) {
  $__num = \func_num_args();
  $__res = function($a_8) use ($Applicative0_6_16, $Bind1_5_4, $f_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_5_4)->{'bind'})($f_7))(function($f_prime__9) use ($Applicative0_6_16, $Bind1_5_4, $a_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_5_4)->{'bind'})($a_8))(function($a_prime__10) use ($Applicative0_6_16, $f_prime__9) {
  $__num = \func_num_args();
  $__res = (($Applicative0_6_16)->{'pure'})(($f_prime__9)($a_prime__10));
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
}, "Functor0" => function($_dollar___unused_5) use ($functorMaybeT1_4_1) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_4_1;
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
  $Bind1_3_17 = (($Monad0_1_0)->{'Bind1'})(null);
  $Applicative0_4_18 = (($Monad0_1_0)->{'Applicative0'})(null);
  $__res = (object)["bind" => function($v_5) use ($Applicative0_4_18, $Bind1_3_17) {
  $__num = \func_num_args();
  $__res = function($f_6) use ($Applicative0_4_18, $Bind1_3_17, $v_5) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_17)->{'bind'})($v_5))(function($v1_7) use ($Applicative0_4_18, $f_6) {
  $__num = \func_num_args();
  $__t19 = null;;
  if ($v1_7 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t19 = (($Applicative0_4_18)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_19;;
};
  if ($v1_7 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t19 = ($f_6)(($v1_7)->{'value0'});
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_5) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__local_var_6_20 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_6_20 = (object)["map" => function($f_7) use ($__local_var_6_20) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($__local_var_6_20, $f_7) {
  $__num = \func_num_args();
  $__res = ((($__local_var_6_20)->{'map'})(function($v1_9) use ($f_7) {
  $__num = \func_num_args();
  $__t21 = null;;
  if ($v1_9 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t21 = new \Data\Maybe\Data_Maybe_Just(($f_7)(($v1_9)->{'value0'}));
goto end_branch_21;;
};
  $__t21 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_21:;
  $__res = $__t21;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_8);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_7_23 = (($Monad0_1_0)->{'Bind1'})(null);
  $Applicative0_8_24 = (($Monad0_1_0)->{'Applicative0'})(null);
  $Bind1_7_23 = (object)["bind" => function($v_9) use ($Applicative0_8_24, $Bind1_7_23) {
  $__num = \func_num_args();
  $__res = function($f_10) use ($Applicative0_8_24, $Bind1_7_23, $v_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_23)->{'bind'})($v_9))(function($v1_11) use ($Applicative0_8_24, $f_10) {
  $__num = \func_num_args();
  $__t25 = null;;
  if ($v1_11 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t25 = (($Applicative0_8_24)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_25;;
};
  if ($v1_11 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t25 = ($f_10)(($v1_11)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_9) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_8_27 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_8) {
  $__num = \func_num_args();
  $__res = $x_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($Monad0_1_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_8) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__local_var_9_27 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_9_27 = (object)["map" => function($f_10) use ($__local_var_9_27) {
  $__num = \func_num_args();
  $__res = function($v_11) use ($__local_var_9_27, $f_10) {
  $__num = \func_num_args();
  $__res = ((($__local_var_9_27)->{'map'})(function($v1_12) use ($f_10) {
  $__num = \func_num_args();
  $__t28 = null;;
  if ($v1_12 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t28 = new \Data\Maybe\Data_Maybe_Just(($f_10)(($v1_12)->{'value0'}));
goto end_branch_28;;
};
  $__t28 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_28:;
  $__res = $__t28;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_11);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_10_30 = (($Monad0_1_0)->{'Bind1'})(null);
  $Applicative0_11_31 = (($Monad0_1_0)->{'Applicative0'})(null);
  $Bind1_10_30 = (object)["bind" => function($v_12) use ($Applicative0_11_31, $Bind1_10_30) {
  $__num = \func_num_args();
  $__res = function($f_13) use ($Applicative0_11_31, $Bind1_10_30, $v_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_30)->{'bind'})($v_12))(function($v1_14) use ($Applicative0_11_31, $f_13) {
  $__num = \func_num_args();
  $__t32 = null;;
  if ($v1_14 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t32 = (($Applicative0_11_31)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_32;;
};
  if ($v1_14 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t32 = ($f_13)(($v1_14)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_12) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_11_34 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($Monad0_1_0);
  $__res = (object)["apply" => function($f_12) use ($Applicative0_11_34, $Bind1_10_30) {
  $__num = \func_num_args();
  $__res = function($a_13) use ($Applicative0_11_34, $Bind1_10_30, $f_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_30)->{'bind'})($f_12))(function($f_prime__14) use ($Applicative0_11_34, $Bind1_10_30, $a_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_30)->{'bind'})($a_13))(function($a_prime__15) use ($Applicative0_11_34, $f_prime__14) {
  $__num = \func_num_args();
  $__res = (($Applicative0_11_34)->{'pure'})(($f_prime__14)($a_prime__15));
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
}, "Functor0" => function($_dollar___unused_10) use ($functorMaybeT1_9_27) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_9_27;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_9) use ($Applicative0_8_27, $Bind1_7_23) {
  $__num = \func_num_args();
  $__res = function($a_10) use ($Applicative0_8_27, $Bind1_7_23, $f_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_23)->{'bind'})($f_9))(function($f_prime__11) use ($Applicative0_8_27, $Bind1_7_23, $a_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_23)->{'bind'})($a_10))(function($a_prime__12) use ($Applicative0_8_27, $f_prime__11) {
  $__num = \func_num_args();
  $__res = (($Applicative0_8_27)->{'pure'})(($f_prime__11)($a_prime__12));
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
}, "Functor0" => function($_dollar___unused_7) use ($functorMaybeT1_6_20) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_6_20;
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
  $Bind1_3_37 = (($Monad0_1_0)->{'Bind1'})(null);
  $Applicative0_4_38 = (($Monad0_1_0)->{'Applicative0'})(null);
  $__res = (object)["liftEffect" => (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_3) {
  $__num = \func_num_args();
  $__res = $x_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($a_5) use ($Applicative0_4_38, $Bind1_3_37) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_37)->{'bind'})($a_5))(function($a_prime__6) use ($Applicative0_4_38) {
  $__num = \func_num_args();
  $__res = (($Applicative0_4_38)->{'pure'})(new \Data\Maybe\Data_Maybe_Just($a_prime__6));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(($dictMonadEffect_0)->{'liftEffect'}), "Monad0" => function($_dollar___unused_3) use ($monadMaybeT1_2_1) {
  $__num = \func_num_args();
  $__res = $monadMaybeT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Maybe_Trans_monadEffectMaybe'] = __NAMESPACE__ . '\\majControl_majMonad_majMaybe_majTrans_monadmajEffectmajMaybe';

// Control_Monad_Maybe_Trans_monadRecMaybeT
function majControl_majMonad_majMaybe_majTrans_monadmajRecmajMaybemajT($dictMonadRec_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majMaybe_majTrans_monadmajRecmajMaybemajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Monad0_1_0 = (($dictMonadRec_0)->{'Monad0'})(null);
  $Bind1_2_1 = (($Monad0_1_0)->{'Bind1'})(null);
  $Applicative0_3_2 = (($Monad0_1_0)->{'Applicative0'})(null);
  $monadMaybeT1_4_3 = (object)["Applicative0" => function($_dollar___unused_4) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_5) {
  $__num = \func_num_args();
  $__res = $x_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($Monad0_1_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_5) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__local_var_6_3 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_6_3 = (object)["map" => function($f_7) use ($__local_var_6_3) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($__local_var_6_3, $f_7) {
  $__num = \func_num_args();
  $__res = ((($__local_var_6_3)->{'map'})(function($v1_9) use ($f_7) {
  $__num = \func_num_args();
  $__t4 = null;;
  if ($v1_9 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t4 = new \Data\Maybe\Data_Maybe_Just(($f_7)(($v1_9)->{'value0'}));
goto end_branch_4;;
};
  $__t4 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_8);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_7_6 = (($Monad0_1_0)->{'Bind1'})(null);
  $Applicative0_8_7 = (($Monad0_1_0)->{'Applicative0'})(null);
  $Bind1_7_6 = (object)["bind" => function($v_9) use ($Applicative0_8_7, $Bind1_7_6) {
  $__num = \func_num_args();
  $__res = function($f_10) use ($Applicative0_8_7, $Bind1_7_6, $v_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_6)->{'bind'})($v_9))(function($v1_11) use ($Applicative0_8_7, $f_10) {
  $__num = \func_num_args();
  $__t8 = null;;
  if ($v1_11 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t8 = (($Applicative0_8_7)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_8;;
};
  if ($v1_11 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t8 = ($f_10)(($v1_11)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_9) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__local_var_10_9 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_10_9 = (object)["map" => function($f_11) use ($__local_var_10_9) {
  $__num = \func_num_args();
  $__res = function($v_12) use ($__local_var_10_9, $f_11) {
  $__num = \func_num_args();
  $__res = ((($__local_var_10_9)->{'map'})(function($v1_13) use ($f_11) {
  $__num = \func_num_args();
  $__t10 = null;;
  if ($v1_13 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t10 = new \Data\Maybe\Data_Maybe_Just(($f_11)(($v1_13)->{'value0'}));
goto end_branch_10;;
};
  $__t10 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_10:;
  $__res = $__t10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_12);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_11_12 = (($Monad0_1_0)->{'Bind1'})(null);
  $Applicative0_12_13 = (($Monad0_1_0)->{'Applicative0'})(null);
  $Bind1_11_12 = (object)["bind" => function($v_13) use ($Applicative0_12_13, $Bind1_11_12) {
  $__num = \func_num_args();
  $__res = function($f_14) use ($Applicative0_12_13, $Bind1_11_12, $v_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_12)->{'bind'})($v_13))(function($v1_15) use ($Applicative0_12_13, $f_14) {
  $__num = \func_num_args();
  $__t14 = null;;
  if ($v1_15 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t14 = (($Applicative0_12_13)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_14;;
};
  if ($v1_15 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t14 = ($f_14)(($v1_15)->{'value0'});
goto end_branch_14;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t14 = null;
  end_branch_14:;
  $__res = $__t14;
  goto __end;;
  __end:
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
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_12_16 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($Monad0_1_0);
  $__res = (object)["apply" => function($f_13) use ($Applicative0_12_16, $Bind1_11_12) {
  $__num = \func_num_args();
  $__res = function($a_14) use ($Applicative0_12_16, $Bind1_11_12, $f_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_12)->{'bind'})($f_13))(function($f_prime__15) use ($Applicative0_12_16, $Bind1_11_12, $a_14) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_12)->{'bind'})($a_14))(function($a_prime__16) use ($Applicative0_12_16, $f_prime__15) {
  $__num = \func_num_args();
  $__res = (($Applicative0_12_16)->{'pure'})(($f_prime__15)($a_prime__16));
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
}, "Functor0" => function($_dollar___unused_11) use ($functorMaybeT1_10_9) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_10_9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_8_18 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($Monad0_1_0);
  $__res = (object)["apply" => function($f_9) use ($Applicative0_8_18, $Bind1_7_6) {
  $__num = \func_num_args();
  $__res = function($a_10) use ($Applicative0_8_18, $Bind1_7_6, $f_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_6)->{'bind'})($f_9))(function($f_prime__11) use ($Applicative0_8_18, $Bind1_7_6, $a_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_6)->{'bind'})($a_10))(function($a_prime__12) use ($Applicative0_8_18, $f_prime__11) {
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
}, "Functor0" => function($_dollar___unused_7) use ($functorMaybeT1_6_3) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_6_3;
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
  $Bind1_5_19 = (($Monad0_1_0)->{'Bind1'})(null);
  $Applicative0_6_20 = (($Monad0_1_0)->{'Applicative0'})(null);
  $__res = (object)["bind" => function($v_7) use ($Applicative0_6_20, $Bind1_5_19) {
  $__num = \func_num_args();
  $__res = function($f_8) use ($Applicative0_6_20, $Bind1_5_19, $v_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_5_19)->{'bind'})($v_7))(function($v1_9) use ($Applicative0_6_20, $f_8) {
  $__num = \func_num_args();
  $__t21 = null;;
  if ($v1_9 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t21 = (($Applicative0_6_20)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_21;;
};
  if ($v1_9 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t21 = ($f_8)(($v1_9)->{'value0'});
goto end_branch_21;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t21 = null;
  end_branch_21:;
  $__res = $__t21;
  goto __end;;
  __end:
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
  $__local_var_8_22 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_8_22 = (object)["map" => function($f_9) use ($__local_var_8_22) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($__local_var_8_22, $f_9) {
  $__num = \func_num_args();
  $__res = ((($__local_var_8_22)->{'map'})(function($v1_11) use ($f_9) {
  $__num = \func_num_args();
  $__t23 = null;;
  if ($v1_11 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t23 = new \Data\Maybe\Data_Maybe_Just(($f_9)(($v1_11)->{'value0'}));
goto end_branch_23;;
};
  $__t23 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_23:;
  $__res = $__t23;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_10);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_9_25 = (($Monad0_1_0)->{'Bind1'})(null);
  $Applicative0_10_26 = (($Monad0_1_0)->{'Applicative0'})(null);
  $Bind1_9_25 = (object)["bind" => function($v_11) use ($Applicative0_10_26, $Bind1_9_25) {
  $__num = \func_num_args();
  $__res = function($f_12) use ($Applicative0_10_26, $Bind1_9_25, $v_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_25)->{'bind'})($v_11))(function($v1_13) use ($Applicative0_10_26, $f_12) {
  $__num = \func_num_args();
  $__t27 = null;;
  if ($v1_13 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t27 = (($Applicative0_10_26)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_27;;
};
  if ($v1_13 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t27 = ($f_12)(($v1_13)->{'value0'});
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_11) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_10_29 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_10) {
  $__num = \func_num_args();
  $__res = $x_10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($Monad0_1_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_10) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__local_var_11_29 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_11_29 = (object)["map" => function($f_12) use ($__local_var_11_29) {
  $__num = \func_num_args();
  $__res = function($v_13) use ($__local_var_11_29, $f_12) {
  $__num = \func_num_args();
  $__res = ((($__local_var_11_29)->{'map'})(function($v1_14) use ($f_12) {
  $__num = \func_num_args();
  $__t30 = null;;
  if ($v1_14 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t30 = new \Data\Maybe\Data_Maybe_Just(($f_12)(($v1_14)->{'value0'}));
goto end_branch_30;;
};
  $__t30 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_30:;
  $__res = $__t30;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_13);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_12_32 = (($Monad0_1_0)->{'Bind1'})(null);
  $Applicative0_13_33 = (($Monad0_1_0)->{'Applicative0'})(null);
  $Bind1_12_32 = (object)["bind" => function($v_14) use ($Applicative0_13_33, $Bind1_12_32) {
  $__num = \func_num_args();
  $__res = function($f_15) use ($Applicative0_13_33, $Bind1_12_32, $v_14) {
  $__num = \func_num_args();
  $__res = ((($Bind1_12_32)->{'bind'})($v_14))(function($v1_16) use ($Applicative0_13_33, $f_15) {
  $__num = \func_num_args();
  $__t34 = null;;
  if ($v1_16 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t34 = (($Applicative0_13_33)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_34;;
};
  if ($v1_16 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t34 = ($f_15)(($v1_16)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_14) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_13_36 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($Monad0_1_0);
  $__res = (object)["apply" => function($f_14) use ($Applicative0_13_36, $Bind1_12_32) {
  $__num = \func_num_args();
  $__res = function($a_15) use ($Applicative0_13_36, $Bind1_12_32, $f_14) {
  $__num = \func_num_args();
  $__res = ((($Bind1_12_32)->{'bind'})($f_14))(function($f_prime__16) use ($Applicative0_13_36, $Bind1_12_32, $a_15) {
  $__num = \func_num_args();
  $__res = ((($Bind1_12_32)->{'bind'})($a_15))(function($a_prime__17) use ($Applicative0_13_36, $f_prime__16) {
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
}, "Functor0" => function($_dollar___unused_12) use ($functorMaybeT1_11_29) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_11_29;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_11) use ($Applicative0_10_29, $Bind1_9_25) {
  $__num = \func_num_args();
  $__res = function($a_12) use ($Applicative0_10_29, $Bind1_9_25, $f_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_25)->{'bind'})($f_11))(function($f_prime__13) use ($Applicative0_10_29, $Bind1_9_25, $a_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_25)->{'bind'})($a_12))(function($a_prime__14) use ($Applicative0_10_29, $f_prime__13) {
  $__num = \func_num_args();
  $__res = (($Applicative0_10_29)->{'pure'})(($f_prime__13)($a_prime__14));
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
}, "Functor0" => function($_dollar___unused_9) use ($functorMaybeT1_8_22) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_8_22;
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
  $__t39 = null;;
  if ($m_prime__7 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t39 = new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Done(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_39;;
};
  if ($m_prime__7 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t40 = null;;
if (($m_prime__7)->{'value0'} instanceof \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Loop) {
$__t40 = new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Loop((($m_prime__7)->{'value0'})->{'value0'});
goto end_branch_40;;
};
if (($m_prime__7)->{'value0'} instanceof \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Done) {
$__t40 = new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Done(new \Data\Maybe\Data_Maybe_Just((($m_prime__7)->{'value0'})->{'value0'}));
goto end_branch_40;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t40 = null;
end_branch_40:;
$__t39 = $__t40;
goto end_branch_39;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t39 = null;
  end_branch_39:;
  $__res = (($Applicative0_3_2)->{'pure'})($__t39);
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
}, "Monad0" => function($_dollar___unused_5) use ($monadMaybeT1_4_3) {
  $__num = \func_num_args();
  $__res = $monadMaybeT1_4_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Maybe_Trans_monadRecMaybeT'] = __NAMESPACE__ . '\\majControl_majMonad_majMaybe_majTrans_monadmajRecmajMaybemajT';

// Control_Monad_Maybe_Trans_monadStateMaybeT
function majControl_majMonad_majMaybe_majTrans_monadmajStatemajMaybemajT($dictMonadState_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majMaybe_majTrans_monadmajStatemajMaybemajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Monad0_1_0 = (($dictMonadState_0)->{'Monad0'})(null);
  $__local_var_2_1 = (($dictMonadState_0)->{'Monad0'})(null);
  $monadMaybeT1_2_1 = (object)["Applicative0" => function($_dollar___unused_3) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_4) {
  $__num = \func_num_args();
  $__res = $x_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_4) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_5_2 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_5_2 = (object)["map" => function($f_6) use ($__local_var_5_2) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($__local_var_5_2, $f_6) {
  $__num = \func_num_args();
  $__res = ((($__local_var_5_2)->{'map'})(function($v1_8) use ($f_6) {
  $__num = \func_num_args();
  $__t3 = null;;
  if ($v1_8 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t3 = new \Data\Maybe\Data_Maybe_Just(($f_6)(($v1_8)->{'value0'}));
goto end_branch_3;;
};
  $__t3 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_6_5 = (($__local_var_2_1)->{'Bind1'})(null);
  $Applicative0_7_6 = (($__local_var_2_1)->{'Applicative0'})(null);
  $Bind1_6_5 = (object)["bind" => function($v_8) use ($Applicative0_7_6, $Bind1_6_5) {
  $__num = \func_num_args();
  $__res = function($f_9) use ($Applicative0_7_6, $Bind1_6_5, $v_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_5)->{'bind'})($v_8))(function($v1_10) use ($Applicative0_7_6, $f_9) {
  $__num = \func_num_args();
  $__t7 = null;;
  if ($v1_10 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t7 = (($Applicative0_7_6)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_7;;
};
  if ($v1_10 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t7 = ($f_9)(($v1_10)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_8) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_9_8 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_9_8 = (object)["map" => function($f_10) use ($__local_var_9_8) {
  $__num = \func_num_args();
  $__res = function($v_11) use ($__local_var_9_8, $f_10) {
  $__num = \func_num_args();
  $__res = ((($__local_var_9_8)->{'map'})(function($v1_12) use ($f_10) {
  $__num = \func_num_args();
  $__t9 = null;;
  if ($v1_12 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t9 = new \Data\Maybe\Data_Maybe_Just(($f_10)(($v1_12)->{'value0'}));
goto end_branch_9;;
};
  $__t9 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_9:;
  $__res = $__t9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_11);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_10_11 = (($__local_var_2_1)->{'Bind1'})(null);
  $Applicative0_11_12 = (($__local_var_2_1)->{'Applicative0'})(null);
  $Bind1_10_11 = (object)["bind" => function($v_12) use ($Applicative0_11_12, $Bind1_10_11) {
  $__num = \func_num_args();
  $__res = function($f_13) use ($Applicative0_11_12, $Bind1_10_11, $v_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_11)->{'bind'})($v_12))(function($v1_14) use ($Applicative0_11_12, $f_13) {
  $__num = \func_num_args();
  $__t13 = null;;
  if ($v1_14 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t13 = (($Applicative0_11_12)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_13;;
};
  if ($v1_14 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t13 = ($f_13)(($v1_14)->{'value0'});
goto end_branch_13;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t13 = null;
  end_branch_13:;
  $__res = $__t13;
  goto __end;;
  __end:
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
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($__local_var_2_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_11_15 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($__local_var_2_1);
  $__res = (object)["apply" => function($f_12) use ($Applicative0_11_15, $Bind1_10_11) {
  $__num = \func_num_args();
  $__res = function($a_13) use ($Applicative0_11_15, $Bind1_10_11, $f_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_11)->{'bind'})($f_12))(function($f_prime__14) use ($Applicative0_11_15, $Bind1_10_11, $a_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_11)->{'bind'})($a_13))(function($a_prime__15) use ($Applicative0_11_15, $f_prime__14) {
  $__num = \func_num_args();
  $__res = (($Applicative0_11_15)->{'pure'})(($f_prime__14)($a_prime__15));
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
}, "Functor0" => function($_dollar___unused_10) use ($functorMaybeT1_9_8) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_9_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_7_17 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($__local_var_2_1);
  $__res = (object)["apply" => function($f_8) use ($Applicative0_7_17, $Bind1_6_5) {
  $__num = \func_num_args();
  $__res = function($a_9) use ($Applicative0_7_17, $Bind1_6_5, $f_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_5)->{'bind'})($f_8))(function($f_prime__10) use ($Applicative0_7_17, $Bind1_6_5, $a_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_5)->{'bind'})($a_9))(function($a_prime__11) use ($Applicative0_7_17, $f_prime__10) {
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
}, "Functor0" => function($_dollar___unused_6) use ($functorMaybeT1_5_2) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_5_2;
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
  $Bind1_4_18 = (($__local_var_2_1)->{'Bind1'})(null);
  $Applicative0_5_19 = (($__local_var_2_1)->{'Applicative0'})(null);
  $__res = (object)["bind" => function($v_6) use ($Applicative0_5_19, $Bind1_4_18) {
  $__num = \func_num_args();
  $__res = function($f_7) use ($Applicative0_5_19, $Bind1_4_18, $v_6) {
  $__num = \func_num_args();
  $__res = ((($Bind1_4_18)->{'bind'})($v_6))(function($v1_8) use ($Applicative0_5_19, $f_7) {
  $__num = \func_num_args();
  $__t20 = null;;
  if ($v1_8 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t20 = (($Applicative0_5_19)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_20;;
};
  if ($v1_8 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t20 = ($f_7)(($v1_8)->{'value0'});
goto end_branch_20;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t20 = null;
  end_branch_20:;
  $__res = $__t20;
  goto __end;;
  __end:
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
  $__local_var_7_21 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_7_21 = (object)["map" => function($f_8) use ($__local_var_7_21) {
  $__num = \func_num_args();
  $__res = function($v_9) use ($__local_var_7_21, $f_8) {
  $__num = \func_num_args();
  $__res = ((($__local_var_7_21)->{'map'})(function($v1_10) use ($f_8) {
  $__num = \func_num_args();
  $__t22 = null;;
  if ($v1_10 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t22 = new \Data\Maybe\Data_Maybe_Just(($f_8)(($v1_10)->{'value0'}));
goto end_branch_22;;
};
  $__t22 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_22:;
  $__res = $__t22;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_8_24 = (($__local_var_2_1)->{'Bind1'})(null);
  $Applicative0_9_25 = (($__local_var_2_1)->{'Applicative0'})(null);
  $Bind1_8_24 = (object)["bind" => function($v_10) use ($Applicative0_9_25, $Bind1_8_24) {
  $__num = \func_num_args();
  $__res = function($f_11) use ($Applicative0_9_25, $Bind1_8_24, $v_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_8_24)->{'bind'})($v_10))(function($v1_12) use ($Applicative0_9_25, $f_11) {
  $__num = \func_num_args();
  $__t26 = null;;
  if ($v1_12 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t26 = (($Applicative0_9_25)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_26;;
};
  if ($v1_12 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t26 = ($f_11)(($v1_12)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_10) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($__local_var_2_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_9_28 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_9) {
  $__num = \func_num_args();
  $__res = $x_9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_9) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_10_28 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_10_28 = (object)["map" => function($f_11) use ($__local_var_10_28) {
  $__num = \func_num_args();
  $__res = function($v_12) use ($__local_var_10_28, $f_11) {
  $__num = \func_num_args();
  $__res = ((($__local_var_10_28)->{'map'})(function($v1_13) use ($f_11) {
  $__num = \func_num_args();
  $__t29 = null;;
  if ($v1_13 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t29 = new \Data\Maybe\Data_Maybe_Just(($f_11)(($v1_13)->{'value0'}));
goto end_branch_29;;
};
  $__t29 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_29:;
  $__res = $__t29;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_12);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_11_31 = (($__local_var_2_1)->{'Bind1'})(null);
  $Applicative0_12_32 = (($__local_var_2_1)->{'Applicative0'})(null);
  $Bind1_11_31 = (object)["bind" => function($v_13) use ($Applicative0_12_32, $Bind1_11_31) {
  $__num = \func_num_args();
  $__res = function($f_14) use ($Applicative0_12_32, $Bind1_11_31, $v_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_31)->{'bind'})($v_13))(function($v1_15) use ($Applicative0_12_32, $f_14) {
  $__num = \func_num_args();
  $__t33 = null;;
  if ($v1_15 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t33 = (($Applicative0_12_32)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_33;;
};
  if ($v1_15 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t33 = ($f_14)(($v1_15)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_13) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($__local_var_2_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_12_35 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($__local_var_2_1);
  $__res = (object)["apply" => function($f_13) use ($Applicative0_12_35, $Bind1_11_31) {
  $__num = \func_num_args();
  $__res = function($a_14) use ($Applicative0_12_35, $Bind1_11_31, $f_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_31)->{'bind'})($f_13))(function($f_prime__15) use ($Applicative0_12_35, $Bind1_11_31, $a_14) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_31)->{'bind'})($a_14))(function($a_prime__16) use ($Applicative0_12_35, $f_prime__15) {
  $__num = \func_num_args();
  $__res = (($Applicative0_12_35)->{'pure'})(($f_prime__15)($a_prime__16));
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
}, "Functor0" => function($_dollar___unused_11) use ($functorMaybeT1_10_28) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_10_28;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_10) use ($Applicative0_9_28, $Bind1_8_24) {
  $__num = \func_num_args();
  $__res = function($a_11) use ($Applicative0_9_28, $Bind1_8_24, $f_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_8_24)->{'bind'})($f_10))(function($f_prime__12) use ($Applicative0_9_28, $Bind1_8_24, $a_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_8_24)->{'bind'})($a_11))(function($a_prime__13) use ($Applicative0_9_28, $f_prime__12) {
  $__num = \func_num_args();
  $__res = (($Applicative0_9_28)->{'pure'})(($f_prime__12)($a_prime__13));
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
}, "Functor0" => function($_dollar___unused_8) use ($functorMaybeT1_7_21) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_7_21;
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
  $Bind1_4_38 = (($Monad0_1_0)->{'Bind1'})(null);
  $Applicative0_5_39 = (($Monad0_1_0)->{'Applicative0'})(null);
  $__res = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($x_4) {
  $__num = \func_num_args();
  $__res = $x_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, function($a_6) use ($Applicative0_5_39, $Bind1_4_38) {
  $__num = \func_num_args();
  $__res = ((($Bind1_4_38)->{'bind'})($a_6))(function($a_prime__7) use ($Applicative0_5_39) {
  $__num = \func_num_args();
  $__res = (($Applicative0_5_39)->{'pure'})(new \Data\Maybe\Data_Maybe_Just($a_prime__7));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($dictMonadState_0)->{'state'})($f_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad0" => function($_dollar___unused_3) use ($monadMaybeT1_2_1) {
  $__num = \func_num_args();
  $__res = $monadMaybeT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Maybe_Trans_monadStateMaybeT'] = __NAMESPACE__ . '\\majControl_majMonad_majMaybe_majTrans_monadmajStatemajMaybemajT';

// Control_Monad_Maybe_Trans_monadTellMaybeT
function majControl_majMonad_majMaybe_majTrans_monadmajTellmajMaybemajT($dictMonadTell_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majMaybe_majTrans_monadmajTellmajMaybemajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Monad1_1_0 = (($dictMonadTell_0)->{'Monad1'})(null);
  $Semigroup0_2_1 = (($dictMonadTell_0)->{'Semigroup0'})(null);
  $monadMaybeT1_3_2 = (object)["Applicative0" => function($_dollar___unused_3) use ($Monad1_1_0) {
  $__num = \func_num_args();
  $__res = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_4) {
  $__num = \func_num_args();
  $__res = $x_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($Monad1_1_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_4) use ($Monad1_1_0) {
  $__num = \func_num_args();
  $__local_var_5_2 = (((((($Monad1_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_5_2 = (object)["map" => function($f_6) use ($__local_var_5_2) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($__local_var_5_2, $f_6) {
  $__num = \func_num_args();
  $__res = ((($__local_var_5_2)->{'map'})(function($v1_8) use ($f_6) {
  $__num = \func_num_args();
  $__t3 = null;;
  if ($v1_8 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t3 = new \Data\Maybe\Data_Maybe_Just(($f_6)(($v1_8)->{'value0'}));
goto end_branch_3;;
};
  $__t3 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_6_5 = (($Monad1_1_0)->{'Bind1'})(null);
  $Applicative0_7_6 = (($Monad1_1_0)->{'Applicative0'})(null);
  $Bind1_6_5 = (object)["bind" => function($v_8) use ($Applicative0_7_6, $Bind1_6_5) {
  $__num = \func_num_args();
  $__res = function($f_9) use ($Applicative0_7_6, $Bind1_6_5, $v_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_5)->{'bind'})($v_8))(function($v1_10) use ($Applicative0_7_6, $f_9) {
  $__num = \func_num_args();
  $__t7 = null;;
  if ($v1_10 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t7 = (($Applicative0_7_6)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_7;;
};
  if ($v1_10 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t7 = ($f_9)(($v1_10)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_8) use ($Monad1_1_0) {
  $__num = \func_num_args();
  $__local_var_9_8 = (((((($Monad1_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_9_8 = (object)["map" => function($f_10) use ($__local_var_9_8) {
  $__num = \func_num_args();
  $__res = function($v_11) use ($__local_var_9_8, $f_10) {
  $__num = \func_num_args();
  $__res = ((($__local_var_9_8)->{'map'})(function($v1_12) use ($f_10) {
  $__num = \func_num_args();
  $__t9 = null;;
  if ($v1_12 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t9 = new \Data\Maybe\Data_Maybe_Just(($f_10)(($v1_12)->{'value0'}));
goto end_branch_9;;
};
  $__t9 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_9:;
  $__res = $__t9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_11);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_10_11 = (($Monad1_1_0)->{'Bind1'})(null);
  $Applicative0_11_12 = (($Monad1_1_0)->{'Applicative0'})(null);
  $Bind1_10_11 = (object)["bind" => function($v_12) use ($Applicative0_11_12, $Bind1_10_11) {
  $__num = \func_num_args();
  $__res = function($f_13) use ($Applicative0_11_12, $Bind1_10_11, $v_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_11)->{'bind'})($v_12))(function($v1_14) use ($Applicative0_11_12, $f_13) {
  $__num = \func_num_args();
  $__t13 = null;;
  if ($v1_14 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t13 = (($Applicative0_11_12)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_13;;
};
  if ($v1_14 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t13 = ($f_13)(($v1_14)->{'value0'});
goto end_branch_13;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t13 = null;
  end_branch_13:;
  $__res = $__t13;
  goto __end;;
  __end:
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
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($Monad1_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_11_15 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($Monad1_1_0);
  $__res = (object)["apply" => function($f_12) use ($Applicative0_11_15, $Bind1_10_11) {
  $__num = \func_num_args();
  $__res = function($a_13) use ($Applicative0_11_15, $Bind1_10_11, $f_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_11)->{'bind'})($f_12))(function($f_prime__14) use ($Applicative0_11_15, $Bind1_10_11, $a_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_11)->{'bind'})($a_13))(function($a_prime__15) use ($Applicative0_11_15, $f_prime__14) {
  $__num = \func_num_args();
  $__res = (($Applicative0_11_15)->{'pure'})(($f_prime__14)($a_prime__15));
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
}, "Functor0" => function($_dollar___unused_10) use ($functorMaybeT1_9_8) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_9_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_7_17 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($Monad1_1_0);
  $__res = (object)["apply" => function($f_8) use ($Applicative0_7_17, $Bind1_6_5) {
  $__num = \func_num_args();
  $__res = function($a_9) use ($Applicative0_7_17, $Bind1_6_5, $f_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_5)->{'bind'})($f_8))(function($f_prime__10) use ($Applicative0_7_17, $Bind1_6_5, $a_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_5)->{'bind'})($a_9))(function($a_prime__11) use ($Applicative0_7_17, $f_prime__10) {
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
}, "Functor0" => function($_dollar___unused_6) use ($functorMaybeT1_5_2) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_5_2;
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
  $Bind1_4_18 = (($Monad1_1_0)->{'Bind1'})(null);
  $Applicative0_5_19 = (($Monad1_1_0)->{'Applicative0'})(null);
  $__res = (object)["bind" => function($v_6) use ($Applicative0_5_19, $Bind1_4_18) {
  $__num = \func_num_args();
  $__res = function($f_7) use ($Applicative0_5_19, $Bind1_4_18, $v_6) {
  $__num = \func_num_args();
  $__res = ((($Bind1_4_18)->{'bind'})($v_6))(function($v1_8) use ($Applicative0_5_19, $f_7) {
  $__num = \func_num_args();
  $__t20 = null;;
  if ($v1_8 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t20 = (($Applicative0_5_19)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_20;;
};
  if ($v1_8 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t20 = ($f_7)(($v1_8)->{'value0'});
goto end_branch_20;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t20 = null;
  end_branch_20:;
  $__res = $__t20;
  goto __end;;
  __end:
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
  $__local_var_7_21 = (((((($Monad1_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_7_21 = (object)["map" => function($f_8) use ($__local_var_7_21) {
  $__num = \func_num_args();
  $__res = function($v_9) use ($__local_var_7_21, $f_8) {
  $__num = \func_num_args();
  $__res = ((($__local_var_7_21)->{'map'})(function($v1_10) use ($f_8) {
  $__num = \func_num_args();
  $__t22 = null;;
  if ($v1_10 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t22 = new \Data\Maybe\Data_Maybe_Just(($f_8)(($v1_10)->{'value0'}));
goto end_branch_22;;
};
  $__t22 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_22:;
  $__res = $__t22;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_8_24 = (($Monad1_1_0)->{'Bind1'})(null);
  $Applicative0_9_25 = (($Monad1_1_0)->{'Applicative0'})(null);
  $Bind1_8_24 = (object)["bind" => function($v_10) use ($Applicative0_9_25, $Bind1_8_24) {
  $__num = \func_num_args();
  $__res = function($f_11) use ($Applicative0_9_25, $Bind1_8_24, $v_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_8_24)->{'bind'})($v_10))(function($v1_12) use ($Applicative0_9_25, $f_11) {
  $__num = \func_num_args();
  $__t26 = null;;
  if ($v1_12 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t26 = (($Applicative0_9_25)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_26;;
};
  if ($v1_12 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t26 = ($f_11)(($v1_12)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_10) use ($Monad1_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($Monad1_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_9_28 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_9) {
  $__num = \func_num_args();
  $__res = $x_9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($Monad1_1_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_9) use ($Monad1_1_0) {
  $__num = \func_num_args();
  $__local_var_10_28 = (((((($Monad1_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_10_28 = (object)["map" => function($f_11) use ($__local_var_10_28) {
  $__num = \func_num_args();
  $__res = function($v_12) use ($__local_var_10_28, $f_11) {
  $__num = \func_num_args();
  $__res = ((($__local_var_10_28)->{'map'})(function($v1_13) use ($f_11) {
  $__num = \func_num_args();
  $__t29 = null;;
  if ($v1_13 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t29 = new \Data\Maybe\Data_Maybe_Just(($f_11)(($v1_13)->{'value0'}));
goto end_branch_29;;
};
  $__t29 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_29:;
  $__res = $__t29;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_12);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_11_31 = (($Monad1_1_0)->{'Bind1'})(null);
  $Applicative0_12_32 = (($Monad1_1_0)->{'Applicative0'})(null);
  $Bind1_11_31 = (object)["bind" => function($v_13) use ($Applicative0_12_32, $Bind1_11_31) {
  $__num = \func_num_args();
  $__res = function($f_14) use ($Applicative0_12_32, $Bind1_11_31, $v_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_31)->{'bind'})($v_13))(function($v1_15) use ($Applicative0_12_32, $f_14) {
  $__num = \func_num_args();
  $__t33 = null;;
  if ($v1_15 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t33 = (($Applicative0_12_32)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_33;;
};
  if ($v1_15 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t33 = ($f_14)(($v1_15)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_13) use ($Monad1_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($Monad1_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_12_35 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($Monad1_1_0);
  $__res = (object)["apply" => function($f_13) use ($Applicative0_12_35, $Bind1_11_31) {
  $__num = \func_num_args();
  $__res = function($a_14) use ($Applicative0_12_35, $Bind1_11_31, $f_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_31)->{'bind'})($f_13))(function($f_prime__15) use ($Applicative0_12_35, $Bind1_11_31, $a_14) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_31)->{'bind'})($a_14))(function($a_prime__16) use ($Applicative0_12_35, $f_prime__15) {
  $__num = \func_num_args();
  $__res = (($Applicative0_12_35)->{'pure'})(($f_prime__15)($a_prime__16));
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
}, "Functor0" => function($_dollar___unused_11) use ($functorMaybeT1_10_28) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_10_28;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_10) use ($Applicative0_9_28, $Bind1_8_24) {
  $__num = \func_num_args();
  $__res = function($a_11) use ($Applicative0_9_28, $Bind1_8_24, $f_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_8_24)->{'bind'})($f_10))(function($f_prime__12) use ($Applicative0_9_28, $Bind1_8_24, $a_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_8_24)->{'bind'})($a_11))(function($a_prime__13) use ($Applicative0_9_28, $f_prime__12) {
  $__num = \func_num_args();
  $__res = (($Applicative0_9_28)->{'pure'})(($f_prime__12)($a_prime__13));
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
}, "Functor0" => function($_dollar___unused_8) use ($functorMaybeT1_7_21) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_7_21;
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
  $Bind1_4_38 = (($Monad1_1_0)->{'Bind1'})(null);
  $Applicative0_5_39 = (($Monad1_1_0)->{'Applicative0'})(null);
  $__res = (object)["tell" => (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_4) {
  $__num = \func_num_args();
  $__res = $x_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($a_6) use ($Applicative0_5_39, $Bind1_4_38) {
  $__num = \func_num_args();
  $__res = ((($Bind1_4_38)->{'bind'})($a_6))(function($a_prime__7) use ($Applicative0_5_39) {
  $__num = \func_num_args();
  $__res = (($Applicative0_5_39)->{'pure'})(new \Data\Maybe\Data_Maybe_Just($a_prime__7));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(($dictMonadTell_0)->{'tell'}), "Semigroup0" => function($_dollar___unused_4) use ($Semigroup0_2_1) {
  $__num = \func_num_args();
  $__res = $Semigroup0_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad1" => function($_dollar___unused_4) use ($monadMaybeT1_3_2) {
  $__num = \func_num_args();
  $__res = $monadMaybeT1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Maybe_Trans_monadTellMaybeT'] = __NAMESPACE__ . '\\majControl_majMonad_majMaybe_majTrans_monadmajTellmajMaybemajT';

// Control_Monad_Maybe_Trans_monadWriterMaybeT
function majControl_majMonad_majMaybe_majTrans_monadmajWritermajMaybemajT($dictMonadWriter_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majMaybe_majTrans_monadmajWritermajMaybemajT';
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
  $monadMaybeT1_9_8 = (object)["Applicative0" => function($_dollar___unused_9) use ($Monad1_7_6) {
  $__num = \func_num_args();
  $__res = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_10) {
  $__num = \func_num_args();
  $__res = $x_10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($Monad1_7_6)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_10) use ($Monad1_7_6) {
  $__num = \func_num_args();
  $__local_var_11_8 = (((((($Monad1_7_6)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_11_8 = (object)["map" => function($f_12) use ($__local_var_11_8) {
  $__num = \func_num_args();
  $__res = function($v_13) use ($__local_var_11_8, $f_12) {
  $__num = \func_num_args();
  $__res = ((($__local_var_11_8)->{'map'})(function($v1_14) use ($f_12) {
  $__num = \func_num_args();
  $__t9 = null;;
  if ($v1_14 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t9 = new \Data\Maybe\Data_Maybe_Just(($f_12)(($v1_14)->{'value0'}));
goto end_branch_9;;
};
  $__t9 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_9:;
  $__res = $__t9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_13);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_12_11 = (($Monad1_7_6)->{'Bind1'})(null);
  $Applicative0_13_12 = (($Monad1_7_6)->{'Applicative0'})(null);
  $Bind1_12_11 = (object)["bind" => function($v_14) use ($Applicative0_13_12, $Bind1_12_11) {
  $__num = \func_num_args();
  $__res = function($f_15) use ($Applicative0_13_12, $Bind1_12_11, $v_14) {
  $__num = \func_num_args();
  $__res = ((($Bind1_12_11)->{'bind'})($v_14))(function($v1_16) use ($Applicative0_13_12, $f_15) {
  $__num = \func_num_args();
  $__t13 = null;;
  if ($v1_16 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t13 = (($Applicative0_13_12)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_13;;
};
  if ($v1_16 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t13 = ($f_15)(($v1_16)->{'value0'});
goto end_branch_13;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t13 = null;
  end_branch_13:;
  $__res = $__t13;
  goto __end;;
  __end:
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
  $__local_var_15_14 = (((((($Monad1_7_6)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_15_14 = (object)["map" => function($f_16) use ($__local_var_15_14) {
  $__num = \func_num_args();
  $__res = function($v_17) use ($__local_var_15_14, $f_16) {
  $__num = \func_num_args();
  $__res = ((($__local_var_15_14)->{'map'})(function($v1_18) use ($f_16) {
  $__num = \func_num_args();
  $__t15 = null;;
  if ($v1_18 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t15 = new \Data\Maybe\Data_Maybe_Just(($f_16)(($v1_18)->{'value0'}));
goto end_branch_15;;
};
  $__t15 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_15:;
  $__res = $__t15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_17);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_16_17 = (($Monad1_7_6)->{'Bind1'})(null);
  $Applicative0_17_18 = (($Monad1_7_6)->{'Applicative0'})(null);
  $Bind1_16_17 = (object)["bind" => function($v_18) use ($Applicative0_17_18, $Bind1_16_17) {
  $__num = \func_num_args();
  $__res = function($f_19) use ($Applicative0_17_18, $Bind1_16_17, $v_18) {
  $__num = \func_num_args();
  $__res = ((($Bind1_16_17)->{'bind'})($v_18))(function($v1_20) use ($Applicative0_17_18, $f_19) {
  $__num = \func_num_args();
  $__t19 = null;;
  if ($v1_20 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t19 = (($Applicative0_17_18)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_19;;
};
  if ($v1_20 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t19 = ($f_19)(($v1_20)->{'value0'});
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_18) use ($Monad1_7_6) {
  $__num = \func_num_args();
  $__local_var_19_20 = (((((($Monad1_7_6)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_19_20 = (object)["map" => function($f_20) use ($__local_var_19_20) {
  $__num = \func_num_args();
  $__res = function($v_21) use ($__local_var_19_20, $f_20) {
  $__num = \func_num_args();
  $__res = ((($__local_var_19_20)->{'map'})(function($v1_22) use ($f_20) {
  $__num = \func_num_args();
  $__t21 = null;;
  if ($v1_22 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t21 = new \Data\Maybe\Data_Maybe_Just(($f_20)(($v1_22)->{'value0'}));
goto end_branch_21;;
};
  $__t21 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_21:;
  $__res = $__t21;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_21);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_20_23 = (($Monad1_7_6)->{'Bind1'})(null);
  $Applicative0_21_24 = (($Monad1_7_6)->{'Applicative0'})(null);
  $Bind1_20_23 = (object)["bind" => function($v_22) use ($Applicative0_21_24, $Bind1_20_23) {
  $__num = \func_num_args();
  $__res = function($f_23) use ($Applicative0_21_24, $Bind1_20_23, $v_22) {
  $__num = \func_num_args();
  $__res = ((($Bind1_20_23)->{'bind'})($v_22))(function($v1_24) use ($Applicative0_21_24, $f_23) {
  $__num = \func_num_args();
  $__t25 = null;;
  if ($v1_24 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t25 = (($Applicative0_21_24)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_25;;
};
  if ($v1_24 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t25 = ($f_23)(($v1_24)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_22) use ($Monad1_7_6) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($Monad1_7_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_21_27 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_21) {
  $__num = \func_num_args();
  $__res = $x_21;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($Monad1_7_6)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_21) use ($Monad1_7_6) {
  $__num = \func_num_args();
  $__local_var_22_27 = (((((($Monad1_7_6)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_22_27 = (object)["map" => function($f_23) use ($__local_var_22_27) {
  $__num = \func_num_args();
  $__res = function($v_24) use ($__local_var_22_27, $f_23) {
  $__num = \func_num_args();
  $__res = ((($__local_var_22_27)->{'map'})(function($v1_25) use ($f_23) {
  $__num = \func_num_args();
  $__t28 = null;;
  if ($v1_25 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t28 = new \Data\Maybe\Data_Maybe_Just(($f_23)(($v1_25)->{'value0'}));
goto end_branch_28;;
};
  $__t28 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_28:;
  $__res = $__t28;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_24);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_23_30 = (($Monad1_7_6)->{'Bind1'})(null);
  $Applicative0_24_31 = (($Monad1_7_6)->{'Applicative0'})(null);
  $Bind1_23_30 = (object)["bind" => function($v_25) use ($Applicative0_24_31, $Bind1_23_30) {
  $__num = \func_num_args();
  $__res = function($f_26) use ($Applicative0_24_31, $Bind1_23_30, $v_25) {
  $__num = \func_num_args();
  $__res = ((($Bind1_23_30)->{'bind'})($v_25))(function($v1_27) use ($Applicative0_24_31, $f_26) {
  $__num = \func_num_args();
  $__t32 = null;;
  if ($v1_27 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t32 = (($Applicative0_24_31)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_32;;
};
  if ($v1_27 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t32 = ($f_26)(($v1_27)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_25) use ($Monad1_7_6) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($Monad1_7_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_24_34 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($Monad1_7_6);
  $__res = (object)["apply" => function($f_25) use ($Applicative0_24_34, $Bind1_23_30) {
  $__num = \func_num_args();
  $__res = function($a_26) use ($Applicative0_24_34, $Bind1_23_30, $f_25) {
  $__num = \func_num_args();
  $__res = ((($Bind1_23_30)->{'bind'})($f_25))(function($f_prime__27) use ($Applicative0_24_34, $Bind1_23_30, $a_26) {
  $__num = \func_num_args();
  $__res = ((($Bind1_23_30)->{'bind'})($a_26))(function($a_prime__28) use ($Applicative0_24_34, $f_prime__27) {
  $__num = \func_num_args();
  $__res = (($Applicative0_24_34)->{'pure'})(($f_prime__27)($a_prime__28));
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
}, "Functor0" => function($_dollar___unused_23) use ($functorMaybeT1_22_27) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_22_27;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_22) use ($Applicative0_21_27, $Bind1_20_23) {
  $__num = \func_num_args();
  $__res = function($a_23) use ($Applicative0_21_27, $Bind1_20_23, $f_22) {
  $__num = \func_num_args();
  $__res = ((($Bind1_20_23)->{'bind'})($f_22))(function($f_prime__24) use ($Applicative0_21_27, $Bind1_20_23, $a_23) {
  $__num = \func_num_args();
  $__res = ((($Bind1_20_23)->{'bind'})($a_23))(function($a_prime__25) use ($Applicative0_21_27, $f_prime__24) {
  $__num = \func_num_args();
  $__res = (($Applicative0_21_27)->{'pure'})(($f_prime__24)($a_prime__25));
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
}, "Functor0" => function($_dollar___unused_20) use ($functorMaybeT1_19_20) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_19_20;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_17_37 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_17) {
  $__num = \func_num_args();
  $__res = $x_17;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($Monad1_7_6)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_17) use ($Monad1_7_6) {
  $__num = \func_num_args();
  $__local_var_18_37 = (((((($Monad1_7_6)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_18_37 = (object)["map" => function($f_19) use ($__local_var_18_37) {
  $__num = \func_num_args();
  $__res = function($v_20) use ($__local_var_18_37, $f_19) {
  $__num = \func_num_args();
  $__res = ((($__local_var_18_37)->{'map'})(function($v1_21) use ($f_19) {
  $__num = \func_num_args();
  $__t38 = null;;
  if ($v1_21 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t38 = new \Data\Maybe\Data_Maybe_Just(($f_19)(($v1_21)->{'value0'}));
goto end_branch_38;;
};
  $__t38 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_38:;
  $__res = $__t38;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_20);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_19_40 = (($Monad1_7_6)->{'Bind1'})(null);
  $Applicative0_20_41 = (($Monad1_7_6)->{'Applicative0'})(null);
  $Bind1_19_40 = (object)["bind" => function($v_21) use ($Applicative0_20_41, $Bind1_19_40) {
  $__num = \func_num_args();
  $__res = function($f_22) use ($Applicative0_20_41, $Bind1_19_40, $v_21) {
  $__num = \func_num_args();
  $__res = ((($Bind1_19_40)->{'bind'})($v_21))(function($v1_23) use ($Applicative0_20_41, $f_22) {
  $__num = \func_num_args();
  $__t42 = null;;
  if ($v1_23 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t42 = (($Applicative0_20_41)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_42;;
};
  if ($v1_23 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t42 = ($f_22)(($v1_23)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_21) use ($Monad1_7_6) {
  $__num = \func_num_args();
  $__local_var_22_43 = (((((($Monad1_7_6)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_22_43 = (object)["map" => function($f_23) use ($__local_var_22_43) {
  $__num = \func_num_args();
  $__res = function($v_24) use ($__local_var_22_43, $f_23) {
  $__num = \func_num_args();
  $__res = ((($__local_var_22_43)->{'map'})(function($v1_25) use ($f_23) {
  $__num = \func_num_args();
  $__t44 = null;;
  if ($v1_25 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t44 = new \Data\Maybe\Data_Maybe_Just(($f_23)(($v1_25)->{'value0'}));
goto end_branch_44;;
};
  $__t44 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_44:;
  $__res = $__t44;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_24);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_23_46 = (($Monad1_7_6)->{'Bind1'})(null);
  $Applicative0_24_47 = (($Monad1_7_6)->{'Applicative0'})(null);
  $Bind1_23_46 = (object)["bind" => function($v_25) use ($Applicative0_24_47, $Bind1_23_46) {
  $__num = \func_num_args();
  $__res = function($f_26) use ($Applicative0_24_47, $Bind1_23_46, $v_25) {
  $__num = \func_num_args();
  $__res = ((($Bind1_23_46)->{'bind'})($v_25))(function($v1_27) use ($Applicative0_24_47, $f_26) {
  $__num = \func_num_args();
  $__t48 = null;;
  if ($v1_27 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t48 = (($Applicative0_24_47)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_48;;
};
  if ($v1_27 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t48 = ($f_26)(($v1_27)->{'value0'});
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_25) use ($Monad1_7_6) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($Monad1_7_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_24_50 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($Monad1_7_6);
  $__res = (object)["apply" => function($f_25) use ($Applicative0_24_50, $Bind1_23_46) {
  $__num = \func_num_args();
  $__res = function($a_26) use ($Applicative0_24_50, $Bind1_23_46, $f_25) {
  $__num = \func_num_args();
  $__res = ((($Bind1_23_46)->{'bind'})($f_25))(function($f_prime__27) use ($Applicative0_24_50, $Bind1_23_46, $a_26) {
  $__num = \func_num_args();
  $__res = ((($Bind1_23_46)->{'bind'})($a_26))(function($a_prime__28) use ($Applicative0_24_50, $f_prime__27) {
  $__num = \func_num_args();
  $__res = (($Applicative0_24_50)->{'pure'})(($f_prime__27)($a_prime__28));
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
}, "Functor0" => function($_dollar___unused_23) use ($functorMaybeT1_22_43) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_22_43;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_20_52 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($Monad1_7_6);
  $__res = (object)["apply" => function($f_21) use ($Applicative0_20_52, $Bind1_19_40) {
  $__num = \func_num_args();
  $__res = function($a_22) use ($Applicative0_20_52, $Bind1_19_40, $f_21) {
  $__num = \func_num_args();
  $__res = ((($Bind1_19_40)->{'bind'})($f_21))(function($f_prime__23) use ($Applicative0_20_52, $Bind1_19_40, $a_22) {
  $__num = \func_num_args();
  $__res = ((($Bind1_19_40)->{'bind'})($a_22))(function($a_prime__24) use ($Applicative0_20_52, $f_prime__23) {
  $__num = \func_num_args();
  $__res = (($Applicative0_20_52)->{'pure'})(($f_prime__23)($a_prime__24));
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
}, "Functor0" => function($_dollar___unused_19) use ($functorMaybeT1_18_37) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_18_37;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_18) use ($Applicative0_17_37, $Bind1_16_17) {
  $__num = \func_num_args();
  $__res = function($a_19) use ($Applicative0_17_37, $Bind1_16_17, $f_18) {
  $__num = \func_num_args();
  $__res = ((($Bind1_16_17)->{'bind'})($f_18))(function($f_prime__20) use ($Applicative0_17_37, $Bind1_16_17, $a_19) {
  $__num = \func_num_args();
  $__res = ((($Bind1_16_17)->{'bind'})($a_19))(function($a_prime__21) use ($Applicative0_17_37, $f_prime__20) {
  $__num = \func_num_args();
  $__res = (($Applicative0_17_37)->{'pure'})(($f_prime__20)($a_prime__21));
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
}, "Functor0" => function($_dollar___unused_16) use ($functorMaybeT1_15_14) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_15_14;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_13_55 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_13) {
  $__num = \func_num_args();
  $__res = $x_13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($Monad1_7_6)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_13) use ($Monad1_7_6) {
  $__num = \func_num_args();
  $__local_var_14_55 = (((((($Monad1_7_6)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_14_55 = (object)["map" => function($f_15) use ($__local_var_14_55) {
  $__num = \func_num_args();
  $__res = function($v_16) use ($__local_var_14_55, $f_15) {
  $__num = \func_num_args();
  $__res = ((($__local_var_14_55)->{'map'})(function($v1_17) use ($f_15) {
  $__num = \func_num_args();
  $__t56 = null;;
  if ($v1_17 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t56 = new \Data\Maybe\Data_Maybe_Just(($f_15)(($v1_17)->{'value0'}));
goto end_branch_56;;
};
  $__t56 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_56:;
  $__res = $__t56;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_16);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_15_58 = (($Monad1_7_6)->{'Bind1'})(null);
  $Applicative0_16_59 = (($Monad1_7_6)->{'Applicative0'})(null);
  $Bind1_15_58 = (object)["bind" => function($v_17) use ($Applicative0_16_59, $Bind1_15_58) {
  $__num = \func_num_args();
  $__res = function($f_18) use ($Applicative0_16_59, $Bind1_15_58, $v_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_58)->{'bind'})($v_17))(function($v1_19) use ($Applicative0_16_59, $f_18) {
  $__num = \func_num_args();
  $__t60 = null;;
  if ($v1_19 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t60 = (($Applicative0_16_59)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_60;;
};
  if ($v1_19 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t60 = ($f_18)(($v1_19)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_17) use ($Monad1_7_6) {
  $__num = \func_num_args();
  $__local_var_18_61 = (((((($Monad1_7_6)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_18_61 = (object)["map" => function($f_19) use ($__local_var_18_61) {
  $__num = \func_num_args();
  $__res = function($v_20) use ($__local_var_18_61, $f_19) {
  $__num = \func_num_args();
  $__res = ((($__local_var_18_61)->{'map'})(function($v1_21) use ($f_19) {
  $__num = \func_num_args();
  $__t62 = null;;
  if ($v1_21 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t62 = new \Data\Maybe\Data_Maybe_Just(($f_19)(($v1_21)->{'value0'}));
goto end_branch_62;;
};
  $__t62 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_62:;
  $__res = $__t62;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_20);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_19_64 = (($Monad1_7_6)->{'Bind1'})(null);
  $Applicative0_20_65 = (($Monad1_7_6)->{'Applicative0'})(null);
  $Bind1_19_64 = (object)["bind" => function($v_21) use ($Applicative0_20_65, $Bind1_19_64) {
  $__num = \func_num_args();
  $__res = function($f_22) use ($Applicative0_20_65, $Bind1_19_64, $v_21) {
  $__num = \func_num_args();
  $__res = ((($Bind1_19_64)->{'bind'})($v_21))(function($v1_23) use ($Applicative0_20_65, $f_22) {
  $__num = \func_num_args();
  $__t66 = null;;
  if ($v1_23 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t66 = (($Applicative0_20_65)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_66;;
};
  if ($v1_23 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t66 = ($f_22)(($v1_23)->{'value0'});
goto end_branch_66;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t66 = null;
  end_branch_66:;
  $__res = $__t66;
  goto __end;;
  __end:
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
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($Monad1_7_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_20_68 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($Monad1_7_6);
  $__res = (object)["apply" => function($f_21) use ($Applicative0_20_68, $Bind1_19_64) {
  $__num = \func_num_args();
  $__res = function($a_22) use ($Applicative0_20_68, $Bind1_19_64, $f_21) {
  $__num = \func_num_args();
  $__res = ((($Bind1_19_64)->{'bind'})($f_21))(function($f_prime__23) use ($Applicative0_20_68, $Bind1_19_64, $a_22) {
  $__num = \func_num_args();
  $__res = ((($Bind1_19_64)->{'bind'})($a_22))(function($a_prime__24) use ($Applicative0_20_68, $f_prime__23) {
  $__num = \func_num_args();
  $__res = (($Applicative0_20_68)->{'pure'})(($f_prime__23)($a_prime__24));
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
}, "Functor0" => function($_dollar___unused_19) use ($functorMaybeT1_18_61) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_18_61;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_16_70 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($Monad1_7_6);
  $__res = (object)["apply" => function($f_17) use ($Applicative0_16_70, $Bind1_15_58) {
  $__num = \func_num_args();
  $__res = function($a_18) use ($Applicative0_16_70, $Bind1_15_58, $f_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_58)->{'bind'})($f_17))(function($f_prime__19) use ($Applicative0_16_70, $Bind1_15_58, $a_18) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_58)->{'bind'})($a_18))(function($a_prime__20) use ($Applicative0_16_70, $f_prime__19) {
  $__num = \func_num_args();
  $__res = (($Applicative0_16_70)->{'pure'})(($f_prime__19)($a_prime__20));
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
}, "Functor0" => function($_dollar___unused_15) use ($functorMaybeT1_14_55) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_14_55;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_14) use ($Applicative0_13_55, $Bind1_12_11) {
  $__num = \func_num_args();
  $__res = function($a_15) use ($Applicative0_13_55, $Bind1_12_11, $f_14) {
  $__num = \func_num_args();
  $__res = ((($Bind1_12_11)->{'bind'})($f_14))(function($f_prime__16) use ($Applicative0_13_55, $Bind1_12_11, $a_15) {
  $__num = \func_num_args();
  $__res = ((($Bind1_12_11)->{'bind'})($a_15))(function($a_prime__17) use ($Applicative0_13_55, $f_prime__16) {
  $__num = \func_num_args();
  $__res = (($Applicative0_13_55)->{'pure'})(($f_prime__16)($a_prime__17));
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
}, "Functor0" => function($_dollar___unused_12) use ($functorMaybeT1_11_8) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_11_8;
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
  $Bind1_10_72 = (($Monad1_7_6)->{'Bind1'})(null);
  $Applicative0_11_73 = (($Monad1_7_6)->{'Applicative0'})(null);
  $__res = (object)["bind" => function($v_12) use ($Applicative0_11_73, $Bind1_10_72) {
  $__num = \func_num_args();
  $__res = function($f_13) use ($Applicative0_11_73, $Bind1_10_72, $v_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_72)->{'bind'})($v_12))(function($v1_14) use ($Applicative0_11_73, $f_13) {
  $__num = \func_num_args();
  $__t74 = null;;
  if ($v1_14 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t74 = (($Applicative0_11_73)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_74;;
};
  if ($v1_14 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t74 = ($f_13)(($v1_14)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_12) use ($Monad1_7_6) {
  $__num = \func_num_args();
  $__local_var_13_75 = (((((($Monad1_7_6)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_13_75 = (object)["map" => function($f_14) use ($__local_var_13_75) {
  $__num = \func_num_args();
  $__res = function($v_15) use ($__local_var_13_75, $f_14) {
  $__num = \func_num_args();
  $__res = ((($__local_var_13_75)->{'map'})(function($v1_16) use ($f_14) {
  $__num = \func_num_args();
  $__t76 = null;;
  if ($v1_16 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t76 = new \Data\Maybe\Data_Maybe_Just(($f_14)(($v1_16)->{'value0'}));
goto end_branch_76;;
};
  $__t76 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_76:;
  $__res = $__t76;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_15);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_14_78 = (($Monad1_7_6)->{'Bind1'})(null);
  $Applicative0_15_79 = (($Monad1_7_6)->{'Applicative0'})(null);
  $Bind1_14_78 = (object)["bind" => function($v_16) use ($Applicative0_15_79, $Bind1_14_78) {
  $__num = \func_num_args();
  $__res = function($f_17) use ($Applicative0_15_79, $Bind1_14_78, $v_16) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_78)->{'bind'})($v_16))(function($v1_18) use ($Applicative0_15_79, $f_17) {
  $__num = \func_num_args();
  $__t80 = null;;
  if ($v1_18 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t80 = (($Applicative0_15_79)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_80;;
};
  if ($v1_18 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t80 = ($f_17)(($v1_18)->{'value0'});
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_16) use ($Monad1_7_6) {
  $__num = \func_num_args();
  $__local_var_17_81 = (((((($Monad1_7_6)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_17_81 = (object)["map" => function($f_18) use ($__local_var_17_81) {
  $__num = \func_num_args();
  $__res = function($v_19) use ($__local_var_17_81, $f_18) {
  $__num = \func_num_args();
  $__res = ((($__local_var_17_81)->{'map'})(function($v1_20) use ($f_18) {
  $__num = \func_num_args();
  $__t82 = null;;
  if ($v1_20 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t82 = new \Data\Maybe\Data_Maybe_Just(($f_18)(($v1_20)->{'value0'}));
goto end_branch_82;;
};
  $__t82 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_82:;
  $__res = $__t82;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_19);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_18_84 = (($Monad1_7_6)->{'Bind1'})(null);
  $Applicative0_19_85 = (($Monad1_7_6)->{'Applicative0'})(null);
  $Bind1_18_84 = (object)["bind" => function($v_20) use ($Applicative0_19_85, $Bind1_18_84) {
  $__num = \func_num_args();
  $__res = function($f_21) use ($Applicative0_19_85, $Bind1_18_84, $v_20) {
  $__num = \func_num_args();
  $__res = ((($Bind1_18_84)->{'bind'})($v_20))(function($v1_22) use ($Applicative0_19_85, $f_21) {
  $__num = \func_num_args();
  $__t86 = null;;
  if ($v1_22 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t86 = (($Applicative0_19_85)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_86;;
};
  if ($v1_22 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t86 = ($f_21)(($v1_22)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_20) use ($Monad1_7_6) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($Monad1_7_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_19_88 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_19) {
  $__num = \func_num_args();
  $__res = $x_19;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($Monad1_7_6)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_19) use ($Monad1_7_6) {
  $__num = \func_num_args();
  $__local_var_20_88 = (((((($Monad1_7_6)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_20_88 = (object)["map" => function($f_21) use ($__local_var_20_88) {
  $__num = \func_num_args();
  $__res = function($v_22) use ($__local_var_20_88, $f_21) {
  $__num = \func_num_args();
  $__res = ((($__local_var_20_88)->{'map'})(function($v1_23) use ($f_21) {
  $__num = \func_num_args();
  $__t89 = null;;
  if ($v1_23 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t89 = new \Data\Maybe\Data_Maybe_Just(($f_21)(($v1_23)->{'value0'}));
goto end_branch_89;;
};
  $__t89 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_89:;
  $__res = $__t89;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_22);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_21_91 = (($Monad1_7_6)->{'Bind1'})(null);
  $Applicative0_22_92 = (($Monad1_7_6)->{'Applicative0'})(null);
  $Bind1_21_91 = (object)["bind" => function($v_23) use ($Applicative0_22_92, $Bind1_21_91) {
  $__num = \func_num_args();
  $__res = function($f_24) use ($Applicative0_22_92, $Bind1_21_91, $v_23) {
  $__num = \func_num_args();
  $__res = ((($Bind1_21_91)->{'bind'})($v_23))(function($v1_25) use ($Applicative0_22_92, $f_24) {
  $__num = \func_num_args();
  $__t93 = null;;
  if ($v1_25 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t93 = (($Applicative0_22_92)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_93;;
};
  if ($v1_25 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t93 = ($f_24)(($v1_25)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_23) use ($Monad1_7_6) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($Monad1_7_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_22_95 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($Monad1_7_6);
  $__res = (object)["apply" => function($f_23) use ($Applicative0_22_95, $Bind1_21_91) {
  $__num = \func_num_args();
  $__res = function($a_24) use ($Applicative0_22_95, $Bind1_21_91, $f_23) {
  $__num = \func_num_args();
  $__res = ((($Bind1_21_91)->{'bind'})($f_23))(function($f_prime__25) use ($Applicative0_22_95, $Bind1_21_91, $a_24) {
  $__num = \func_num_args();
  $__res = ((($Bind1_21_91)->{'bind'})($a_24))(function($a_prime__26) use ($Applicative0_22_95, $f_prime__25) {
  $__num = \func_num_args();
  $__res = (($Applicative0_22_95)->{'pure'})(($f_prime__25)($a_prime__26));
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
}, "Functor0" => function($_dollar___unused_21) use ($functorMaybeT1_20_88) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_20_88;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_20) use ($Applicative0_19_88, $Bind1_18_84) {
  $__num = \func_num_args();
  $__res = function($a_21) use ($Applicative0_19_88, $Bind1_18_84, $f_20) {
  $__num = \func_num_args();
  $__res = ((($Bind1_18_84)->{'bind'})($f_20))(function($f_prime__22) use ($Applicative0_19_88, $Bind1_18_84, $a_21) {
  $__num = \func_num_args();
  $__res = ((($Bind1_18_84)->{'bind'})($a_21))(function($a_prime__23) use ($Applicative0_19_88, $f_prime__22) {
  $__num = \func_num_args();
  $__res = (($Applicative0_19_88)->{'pure'})(($f_prime__22)($a_prime__23));
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
}, "Functor0" => function($_dollar___unused_18) use ($functorMaybeT1_17_81) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_17_81;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_15_98 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_15) {
  $__num = \func_num_args();
  $__res = $x_15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($Monad1_7_6)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_15) use ($Monad1_7_6) {
  $__num = \func_num_args();
  $__local_var_16_98 = (((((($Monad1_7_6)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_16_98 = (object)["map" => function($f_17) use ($__local_var_16_98) {
  $__num = \func_num_args();
  $__res = function($v_18) use ($__local_var_16_98, $f_17) {
  $__num = \func_num_args();
  $__res = ((($__local_var_16_98)->{'map'})(function($v1_19) use ($f_17) {
  $__num = \func_num_args();
  $__t99 = null;;
  if ($v1_19 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t99 = new \Data\Maybe\Data_Maybe_Just(($f_17)(($v1_19)->{'value0'}));
goto end_branch_99;;
};
  $__t99 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_99:;
  $__res = $__t99;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_18);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_17_101 = (($Monad1_7_6)->{'Bind1'})(null);
  $Applicative0_18_102 = (($Monad1_7_6)->{'Applicative0'})(null);
  $Bind1_17_101 = (object)["bind" => function($v_19) use ($Applicative0_18_102, $Bind1_17_101) {
  $__num = \func_num_args();
  $__res = function($f_20) use ($Applicative0_18_102, $Bind1_17_101, $v_19) {
  $__num = \func_num_args();
  $__res = ((($Bind1_17_101)->{'bind'})($v_19))(function($v1_21) use ($Applicative0_18_102, $f_20) {
  $__num = \func_num_args();
  $__t103 = null;;
  if ($v1_21 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t103 = (($Applicative0_18_102)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_103;;
};
  if ($v1_21 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t103 = ($f_20)(($v1_21)->{'value0'});
goto end_branch_103;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t103 = null;
  end_branch_103:;
  $__res = $__t103;
  goto __end;;
  __end:
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
  $__local_var_20_104 = (((((($Monad1_7_6)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_20_104 = (object)["map" => function($f_21) use ($__local_var_20_104) {
  $__num = \func_num_args();
  $__res = function($v_22) use ($__local_var_20_104, $f_21) {
  $__num = \func_num_args();
  $__res = ((($__local_var_20_104)->{'map'})(function($v1_23) use ($f_21) {
  $__num = \func_num_args();
  $__t105 = null;;
  if ($v1_23 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t105 = new \Data\Maybe\Data_Maybe_Just(($f_21)(($v1_23)->{'value0'}));
goto end_branch_105;;
};
  $__t105 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_105:;
  $__res = $__t105;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_22);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_21_107 = (($Monad1_7_6)->{'Bind1'})(null);
  $Applicative0_22_108 = (($Monad1_7_6)->{'Applicative0'})(null);
  $Bind1_21_107 = (object)["bind" => function($v_23) use ($Applicative0_22_108, $Bind1_21_107) {
  $__num = \func_num_args();
  $__res = function($f_24) use ($Applicative0_22_108, $Bind1_21_107, $v_23) {
  $__num = \func_num_args();
  $__res = ((($Bind1_21_107)->{'bind'})($v_23))(function($v1_25) use ($Applicative0_22_108, $f_24) {
  $__num = \func_num_args();
  $__t109 = null;;
  if ($v1_25 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t109 = (($Applicative0_22_108)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_109;;
};
  if ($v1_25 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t109 = ($f_24)(($v1_25)->{'value0'});
goto end_branch_109;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t109 = null;
  end_branch_109:;
  $__res = $__t109;
  goto __end;;
  __end:
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
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($Monad1_7_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_22_111 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_22) {
  $__num = \func_num_args();
  $__res = $x_22;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($Monad1_7_6)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_22) use ($Monad1_7_6) {
  $__num = \func_num_args();
  $__local_var_23_111 = (((((($Monad1_7_6)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_23_111 = (object)["map" => function($f_24) use ($__local_var_23_111) {
  $__num = \func_num_args();
  $__res = function($v_25) use ($__local_var_23_111, $f_24) {
  $__num = \func_num_args();
  $__res = ((($__local_var_23_111)->{'map'})(function($v1_26) use ($f_24) {
  $__num = \func_num_args();
  $__t112 = null;;
  if ($v1_26 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t112 = new \Data\Maybe\Data_Maybe_Just(($f_24)(($v1_26)->{'value0'}));
goto end_branch_112;;
};
  $__t112 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_112:;
  $__res = $__t112;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_25);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_24_114 = (($Monad1_7_6)->{'Bind1'})(null);
  $Applicative0_25_115 = (($Monad1_7_6)->{'Applicative0'})(null);
  $Bind1_24_114 = (object)["bind" => function($v_26) use ($Applicative0_25_115, $Bind1_24_114) {
  $__num = \func_num_args();
  $__res = function($f_27) use ($Applicative0_25_115, $Bind1_24_114, $v_26) {
  $__num = \func_num_args();
  $__res = ((($Bind1_24_114)->{'bind'})($v_26))(function($v1_28) use ($Applicative0_25_115, $f_27) {
  $__num = \func_num_args();
  $__t116 = null;;
  if ($v1_28 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t116 = (($Applicative0_25_115)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_116;;
};
  if ($v1_28 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t116 = ($f_27)(($v1_28)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_26) use ($Monad1_7_6) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($Monad1_7_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_25_118 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($Monad1_7_6);
  $__res = (object)["apply" => function($f_26) use ($Applicative0_25_118, $Bind1_24_114) {
  $__num = \func_num_args();
  $__res = function($a_27) use ($Applicative0_25_118, $Bind1_24_114, $f_26) {
  $__num = \func_num_args();
  $__res = ((($Bind1_24_114)->{'bind'})($f_26))(function($f_prime__28) use ($Applicative0_25_118, $Bind1_24_114, $a_27) {
  $__num = \func_num_args();
  $__res = ((($Bind1_24_114)->{'bind'})($a_27))(function($a_prime__29) use ($Applicative0_25_118, $f_prime__28) {
  $__num = \func_num_args();
  $__res = (($Applicative0_25_118)->{'pure'})(($f_prime__28)($a_prime__29));
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
}, "Functor0" => function($_dollar___unused_24) use ($functorMaybeT1_23_111) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_23_111;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_23) use ($Applicative0_22_111, $Bind1_21_107) {
  $__num = \func_num_args();
  $__res = function($a_24) use ($Applicative0_22_111, $Bind1_21_107, $f_23) {
  $__num = \func_num_args();
  $__res = ((($Bind1_21_107)->{'bind'})($f_23))(function($f_prime__25) use ($Applicative0_22_111, $Bind1_21_107, $a_24) {
  $__num = \func_num_args();
  $__res = ((($Bind1_21_107)->{'bind'})($a_24))(function($a_prime__26) use ($Applicative0_22_111, $f_prime__25) {
  $__num = \func_num_args();
  $__res = (($Applicative0_22_111)->{'pure'})(($f_prime__25)($a_prime__26));
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
}, "Functor0" => function($_dollar___unused_21) use ($functorMaybeT1_20_104) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_20_104;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_18_121 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_18) {
  $__num = \func_num_args();
  $__res = $x_18;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($Monad1_7_6)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_18) use ($Monad1_7_6) {
  $__num = \func_num_args();
  $__local_var_19_121 = (((((($Monad1_7_6)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_19_121 = (object)["map" => function($f_20) use ($__local_var_19_121) {
  $__num = \func_num_args();
  $__res = function($v_21) use ($__local_var_19_121, $f_20) {
  $__num = \func_num_args();
  $__res = ((($__local_var_19_121)->{'map'})(function($v1_22) use ($f_20) {
  $__num = \func_num_args();
  $__t122 = null;;
  if ($v1_22 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t122 = new \Data\Maybe\Data_Maybe_Just(($f_20)(($v1_22)->{'value0'}));
goto end_branch_122;;
};
  $__t122 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_122:;
  $__res = $__t122;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_21);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_20_124 = (($Monad1_7_6)->{'Bind1'})(null);
  $Applicative0_21_125 = (($Monad1_7_6)->{'Applicative0'})(null);
  $Bind1_20_124 = (object)["bind" => function($v_22) use ($Applicative0_21_125, $Bind1_20_124) {
  $__num = \func_num_args();
  $__res = function($f_23) use ($Applicative0_21_125, $Bind1_20_124, $v_22) {
  $__num = \func_num_args();
  $__res = ((($Bind1_20_124)->{'bind'})($v_22))(function($v1_24) use ($Applicative0_21_125, $f_23) {
  $__num = \func_num_args();
  $__t126 = null;;
  if ($v1_24 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t126 = (($Applicative0_21_125)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_126;;
};
  if ($v1_24 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t126 = ($f_23)(($v1_24)->{'value0'});
goto end_branch_126;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t126 = null;
  end_branch_126:;
  $__res = $__t126;
  goto __end;;
  __end:
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
  $__local_var_23_127 = (((((($Monad1_7_6)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_23_127 = (object)["map" => function($f_24) use ($__local_var_23_127) {
  $__num = \func_num_args();
  $__res = function($v_25) use ($__local_var_23_127, $f_24) {
  $__num = \func_num_args();
  $__res = ((($__local_var_23_127)->{'map'})(function($v1_26) use ($f_24) {
  $__num = \func_num_args();
  $__t128 = null;;
  if ($v1_26 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t128 = new \Data\Maybe\Data_Maybe_Just(($f_24)(($v1_26)->{'value0'}));
goto end_branch_128;;
};
  $__t128 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_128:;
  $__res = $__t128;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_25);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_24_130 = (($Monad1_7_6)->{'Bind1'})(null);
  $Applicative0_25_131 = (($Monad1_7_6)->{'Applicative0'})(null);
  $Bind1_24_130 = (object)["bind" => function($v_26) use ($Applicative0_25_131, $Bind1_24_130) {
  $__num = \func_num_args();
  $__res = function($f_27) use ($Applicative0_25_131, $Bind1_24_130, $v_26) {
  $__num = \func_num_args();
  $__res = ((($Bind1_24_130)->{'bind'})($v_26))(function($v1_28) use ($Applicative0_25_131, $f_27) {
  $__num = \func_num_args();
  $__t132 = null;;
  if ($v1_28 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t132 = (($Applicative0_25_131)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_132;;
};
  if ($v1_28 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t132 = ($f_27)(($v1_28)->{'value0'});
goto end_branch_132;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t132 = null;
  end_branch_132:;
  $__res = $__t132;
  goto __end;;
  __end:
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
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($Monad1_7_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_25_134 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($Monad1_7_6);
  $__res = (object)["apply" => function($f_26) use ($Applicative0_25_134, $Bind1_24_130) {
  $__num = \func_num_args();
  $__res = function($a_27) use ($Applicative0_25_134, $Bind1_24_130, $f_26) {
  $__num = \func_num_args();
  $__res = ((($Bind1_24_130)->{'bind'})($f_26))(function($f_prime__28) use ($Applicative0_25_134, $Bind1_24_130, $a_27) {
  $__num = \func_num_args();
  $__res = ((($Bind1_24_130)->{'bind'})($a_27))(function($a_prime__29) use ($Applicative0_25_134, $f_prime__28) {
  $__num = \func_num_args();
  $__res = (($Applicative0_25_134)->{'pure'})(($f_prime__28)($a_prime__29));
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
}, "Functor0" => function($_dollar___unused_24) use ($functorMaybeT1_23_127) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_23_127;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_21_136 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($Monad1_7_6);
  $__res = (object)["apply" => function($f_22) use ($Applicative0_21_136, $Bind1_20_124) {
  $__num = \func_num_args();
  $__res = function($a_23) use ($Applicative0_21_136, $Bind1_20_124, $f_22) {
  $__num = \func_num_args();
  $__res = ((($Bind1_20_124)->{'bind'})($f_22))(function($f_prime__24) use ($Applicative0_21_136, $Bind1_20_124, $a_23) {
  $__num = \func_num_args();
  $__res = ((($Bind1_20_124)->{'bind'})($a_23))(function($a_prime__25) use ($Applicative0_21_136, $f_prime__24) {
  $__num = \func_num_args();
  $__res = (($Applicative0_21_136)->{'pure'})(($f_prime__24)($a_prime__25));
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
}, "Functor0" => function($_dollar___unused_20) use ($functorMaybeT1_19_121) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_19_121;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_19) use ($Applicative0_18_121, $Bind1_17_101) {
  $__num = \func_num_args();
  $__res = function($a_20) use ($Applicative0_18_121, $Bind1_17_101, $f_19) {
  $__num = \func_num_args();
  $__res = ((($Bind1_17_101)->{'bind'})($f_19))(function($f_prime__21) use ($Applicative0_18_121, $Bind1_17_101, $a_20) {
  $__num = \func_num_args();
  $__res = ((($Bind1_17_101)->{'bind'})($a_20))(function($a_prime__22) use ($Applicative0_18_121, $f_prime__21) {
  $__num = \func_num_args();
  $__res = (($Applicative0_18_121)->{'pure'})(($f_prime__21)($a_prime__22));
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
}, "Functor0" => function($_dollar___unused_17) use ($functorMaybeT1_16_98) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_16_98;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_16) use ($Applicative0_15_98, $Bind1_14_78) {
  $__num = \func_num_args();
  $__res = function($a_17) use ($Applicative0_15_98, $Bind1_14_78, $f_16) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_78)->{'bind'})($f_16))(function($f_prime__18) use ($Applicative0_15_98, $Bind1_14_78, $a_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_78)->{'bind'})($a_17))(function($a_prime__19) use ($Applicative0_15_98, $f_prime__18) {
  $__num = \func_num_args();
  $__res = (($Applicative0_15_98)->{'pure'})(($f_prime__18)($a_prime__19));
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
}, "Functor0" => function($_dollar___unused_14) use ($functorMaybeT1_13_75) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_13_75;
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
  $Bind1_10_140 = (($Monad1_7_6)->{'Bind1'})(null);
  $Applicative0_11_141 = (($Monad1_7_6)->{'Applicative0'})(null);
  $monadTellMaybeT1_7_6 = (object)["tell" => (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_10) {
  $__num = \func_num_args();
  $__res = $x_10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($a_12) use ($Applicative0_11_141, $Bind1_10_140) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_140)->{'bind'})($a_12))(function($a_prime__13) use ($Applicative0_11_141) {
  $__num = \func_num_args();
  $__res = (($Applicative0_11_141)->{'pure'})(new \Data\Maybe\Data_Maybe_Just($a_prime__13));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(($MonadTell1_1_0)->{'tell'}), "Semigroup0" => function($_dollar___unused_10) use ($Semigroup0_8_7) {
  $__num = \func_num_args();
  $__res = $Semigroup0_8_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad1" => function($_dollar___unused_10) use ($monadMaybeT1_9_8) {
  $__num = \func_num_args();
  $__res = $monadMaybeT1_9_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["listen" => function($v_8) use ($Bind1_3_2, $dictMonadWriter_0, $pure_4_3) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_2)->{'bind'})((($dictMonadWriter_0)->{'listen'})($v_8)))(function($v_9) use ($pure_4_3) {
  $__num = \func_num_args();
  $__t143 = null;;
  if (($v_9)->{'value0'} instanceof \Data\Maybe\Data_Maybe_Just) {
$__t143 = new \Data\Maybe\Data_Maybe_Just(new \Data\Tuple\Data_Tuple_Tuple((($v_9)->{'value0'})->{'value0'}, ($v_9)->{'value1'}));
goto end_branch_143;;
};
  $__t143 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_143:;
  $__res = ($pure_4_3)($__t143);
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
  $__t144 = null;;
  if ($a_9 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t144 = new \Data\Tuple\Data_Tuple_Tuple(new \Data\Maybe\Data_Maybe_Nothing(), function($x_10) {
  $__num = \func_num_args();
  $__res = $x_10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
goto end_branch_144;;
};
  if ($a_9 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t144 = new \Data\Tuple\Data_Tuple_Tuple(new \Data\Maybe\Data_Maybe_Just((($a_9)->{'value0'})->{'value0'}), (($a_9)->{'value0'})->{'value1'});
goto end_branch_144;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t144 = null;
  end_branch_144:;
  $__res = (($Applicative0_5_4)->{'pure'})($__t144);
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
}, "MonadTell1" => function($_dollar___unused_8) use ($monadTellMaybeT1_7_6) {
  $__num = \func_num_args();
  $__res = $monadTellMaybeT1_7_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Maybe_Trans_monadWriterMaybeT'] = __NAMESPACE__ . '\\majControl_majMonad_majMaybe_majTrans_monadmajWritermajMaybemajT';

// Control_Monad_Maybe_Trans_monadThrowMaybeT
function majControl_majMonad_majMaybe_majTrans_monadmajThrowmajMaybemajT($dictMonadThrow_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majMaybe_majTrans_monadmajThrowmajMaybemajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Monad0_1_0 = (($dictMonadThrow_0)->{'Monad0'})(null);
  $__local_var_2_1 = (($dictMonadThrow_0)->{'Monad0'})(null);
  $monadMaybeT1_2_1 = (object)["Applicative0" => function($_dollar___unused_3) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_4) {
  $__num = \func_num_args();
  $__res = $x_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_4) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_5_2 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_5_2 = (object)["map" => function($f_6) use ($__local_var_5_2) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($__local_var_5_2, $f_6) {
  $__num = \func_num_args();
  $__res = ((($__local_var_5_2)->{'map'})(function($v1_8) use ($f_6) {
  $__num = \func_num_args();
  $__t3 = null;;
  if ($v1_8 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t3 = new \Data\Maybe\Data_Maybe_Just(($f_6)(($v1_8)->{'value0'}));
goto end_branch_3;;
};
  $__t3 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_6_5 = (($__local_var_2_1)->{'Bind1'})(null);
  $Applicative0_7_6 = (($__local_var_2_1)->{'Applicative0'})(null);
  $Bind1_6_5 = (object)["bind" => function($v_8) use ($Applicative0_7_6, $Bind1_6_5) {
  $__num = \func_num_args();
  $__res = function($f_9) use ($Applicative0_7_6, $Bind1_6_5, $v_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_5)->{'bind'})($v_8))(function($v1_10) use ($Applicative0_7_6, $f_9) {
  $__num = \func_num_args();
  $__t7 = null;;
  if ($v1_10 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t7 = (($Applicative0_7_6)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_7;;
};
  if ($v1_10 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t7 = ($f_9)(($v1_10)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_8) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_9_8 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_9_8 = (object)["map" => function($f_10) use ($__local_var_9_8) {
  $__num = \func_num_args();
  $__res = function($v_11) use ($__local_var_9_8, $f_10) {
  $__num = \func_num_args();
  $__res = ((($__local_var_9_8)->{'map'})(function($v1_12) use ($f_10) {
  $__num = \func_num_args();
  $__t9 = null;;
  if ($v1_12 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t9 = new \Data\Maybe\Data_Maybe_Just(($f_10)(($v1_12)->{'value0'}));
goto end_branch_9;;
};
  $__t9 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_9:;
  $__res = $__t9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_11);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_10_11 = (($__local_var_2_1)->{'Bind1'})(null);
  $Applicative0_11_12 = (($__local_var_2_1)->{'Applicative0'})(null);
  $Bind1_10_11 = (object)["bind" => function($v_12) use ($Applicative0_11_12, $Bind1_10_11) {
  $__num = \func_num_args();
  $__res = function($f_13) use ($Applicative0_11_12, $Bind1_10_11, $v_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_11)->{'bind'})($v_12))(function($v1_14) use ($Applicative0_11_12, $f_13) {
  $__num = \func_num_args();
  $__t13 = null;;
  if ($v1_14 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t13 = (($Applicative0_11_12)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_13;;
};
  if ($v1_14 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t13 = ($f_13)(($v1_14)->{'value0'});
goto end_branch_13;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t13 = null;
  end_branch_13:;
  $__res = $__t13;
  goto __end;;
  __end:
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
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($__local_var_2_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_11_15 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($__local_var_2_1);
  $__res = (object)["apply" => function($f_12) use ($Applicative0_11_15, $Bind1_10_11) {
  $__num = \func_num_args();
  $__res = function($a_13) use ($Applicative0_11_15, $Bind1_10_11, $f_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_11)->{'bind'})($f_12))(function($f_prime__14) use ($Applicative0_11_15, $Bind1_10_11, $a_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_11)->{'bind'})($a_13))(function($a_prime__15) use ($Applicative0_11_15, $f_prime__14) {
  $__num = \func_num_args();
  $__res = (($Applicative0_11_15)->{'pure'})(($f_prime__14)($a_prime__15));
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
}, "Functor0" => function($_dollar___unused_10) use ($functorMaybeT1_9_8) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_9_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_7_17 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($__local_var_2_1);
  $__res = (object)["apply" => function($f_8) use ($Applicative0_7_17, $Bind1_6_5) {
  $__num = \func_num_args();
  $__res = function($a_9) use ($Applicative0_7_17, $Bind1_6_5, $f_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_5)->{'bind'})($f_8))(function($f_prime__10) use ($Applicative0_7_17, $Bind1_6_5, $a_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_5)->{'bind'})($a_9))(function($a_prime__11) use ($Applicative0_7_17, $f_prime__10) {
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
}, "Functor0" => function($_dollar___unused_6) use ($functorMaybeT1_5_2) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_5_2;
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
  $Bind1_4_18 = (($__local_var_2_1)->{'Bind1'})(null);
  $Applicative0_5_19 = (($__local_var_2_1)->{'Applicative0'})(null);
  $__res = (object)["bind" => function($v_6) use ($Applicative0_5_19, $Bind1_4_18) {
  $__num = \func_num_args();
  $__res = function($f_7) use ($Applicative0_5_19, $Bind1_4_18, $v_6) {
  $__num = \func_num_args();
  $__res = ((($Bind1_4_18)->{'bind'})($v_6))(function($v1_8) use ($Applicative0_5_19, $f_7) {
  $__num = \func_num_args();
  $__t20 = null;;
  if ($v1_8 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t20 = (($Applicative0_5_19)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_20;;
};
  if ($v1_8 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t20 = ($f_7)(($v1_8)->{'value0'});
goto end_branch_20;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t20 = null;
  end_branch_20:;
  $__res = $__t20;
  goto __end;;
  __end:
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
  $__local_var_7_21 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_7_21 = (object)["map" => function($f_8) use ($__local_var_7_21) {
  $__num = \func_num_args();
  $__res = function($v_9) use ($__local_var_7_21, $f_8) {
  $__num = \func_num_args();
  $__res = ((($__local_var_7_21)->{'map'})(function($v1_10) use ($f_8) {
  $__num = \func_num_args();
  $__t22 = null;;
  if ($v1_10 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t22 = new \Data\Maybe\Data_Maybe_Just(($f_8)(($v1_10)->{'value0'}));
goto end_branch_22;;
};
  $__t22 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_22:;
  $__res = $__t22;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_8_24 = (($__local_var_2_1)->{'Bind1'})(null);
  $Applicative0_9_25 = (($__local_var_2_1)->{'Applicative0'})(null);
  $Bind1_8_24 = (object)["bind" => function($v_10) use ($Applicative0_9_25, $Bind1_8_24) {
  $__num = \func_num_args();
  $__res = function($f_11) use ($Applicative0_9_25, $Bind1_8_24, $v_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_8_24)->{'bind'})($v_10))(function($v1_12) use ($Applicative0_9_25, $f_11) {
  $__num = \func_num_args();
  $__t26 = null;;
  if ($v1_12 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t26 = (($Applicative0_9_25)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_26;;
};
  if ($v1_12 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t26 = ($f_11)(($v1_12)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_10) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($__local_var_2_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_9_28 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_9) {
  $__num = \func_num_args();
  $__res = $x_9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($__local_var_2_1)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_9) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_10_28 = (((((($__local_var_2_1)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_10_28 = (object)["map" => function($f_11) use ($__local_var_10_28) {
  $__num = \func_num_args();
  $__res = function($v_12) use ($__local_var_10_28, $f_11) {
  $__num = \func_num_args();
  $__res = ((($__local_var_10_28)->{'map'})(function($v1_13) use ($f_11) {
  $__num = \func_num_args();
  $__t29 = null;;
  if ($v1_13 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t29 = new \Data\Maybe\Data_Maybe_Just(($f_11)(($v1_13)->{'value0'}));
goto end_branch_29;;
};
  $__t29 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_29:;
  $__res = $__t29;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_12);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_11_31 = (($__local_var_2_1)->{'Bind1'})(null);
  $Applicative0_12_32 = (($__local_var_2_1)->{'Applicative0'})(null);
  $Bind1_11_31 = (object)["bind" => function($v_13) use ($Applicative0_12_32, $Bind1_11_31) {
  $__num = \func_num_args();
  $__res = function($f_14) use ($Applicative0_12_32, $Bind1_11_31, $v_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_31)->{'bind'})($v_13))(function($v1_15) use ($Applicative0_12_32, $f_14) {
  $__num = \func_num_args();
  $__t33 = null;;
  if ($v1_15 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t33 = (($Applicative0_12_32)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_33;;
};
  if ($v1_15 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t33 = ($f_14)(($v1_15)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_13) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($__local_var_2_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_12_35 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($__local_var_2_1);
  $__res = (object)["apply" => function($f_13) use ($Applicative0_12_35, $Bind1_11_31) {
  $__num = \func_num_args();
  $__res = function($a_14) use ($Applicative0_12_35, $Bind1_11_31, $f_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_31)->{'bind'})($f_13))(function($f_prime__15) use ($Applicative0_12_35, $Bind1_11_31, $a_14) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_31)->{'bind'})($a_14))(function($a_prime__16) use ($Applicative0_12_35, $f_prime__15) {
  $__num = \func_num_args();
  $__res = (($Applicative0_12_35)->{'pure'})(($f_prime__15)($a_prime__16));
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
}, "Functor0" => function($_dollar___unused_11) use ($functorMaybeT1_10_28) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_10_28;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_10) use ($Applicative0_9_28, $Bind1_8_24) {
  $__num = \func_num_args();
  $__res = function($a_11) use ($Applicative0_9_28, $Bind1_8_24, $f_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_8_24)->{'bind'})($f_10))(function($f_prime__12) use ($Applicative0_9_28, $Bind1_8_24, $a_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_8_24)->{'bind'})($a_11))(function($a_prime__13) use ($Applicative0_9_28, $f_prime__12) {
  $__num = \func_num_args();
  $__res = (($Applicative0_9_28)->{'pure'})(($f_prime__12)($a_prime__13));
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
}, "Functor0" => function($_dollar___unused_8) use ($functorMaybeT1_7_21) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_7_21;
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
  $Bind1_4_38 = (($Monad0_1_0)->{'Bind1'})(null);
  $Applicative0_5_39 = (($Monad0_1_0)->{'Applicative0'})(null);
  $__res = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($x_4) {
  $__num = \func_num_args();
  $__res = $x_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, function($a_6) use ($Applicative0_5_39, $Bind1_4_38) {
  $__num = \func_num_args();
  $__res = ((($Bind1_4_38)->{'bind'})($a_6))(function($a_prime__7) use ($Applicative0_5_39) {
  $__num = \func_num_args();
  $__res = (($Applicative0_5_39)->{'pure'})(new \Data\Maybe\Data_Maybe_Just($a_prime__7));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($dictMonadThrow_0)->{'throwError'})($e_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad0" => function($_dollar___unused_3) use ($monadMaybeT1_2_1) {
  $__num = \func_num_args();
  $__res = $monadMaybeT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Maybe_Trans_monadThrowMaybeT'] = __NAMESPACE__ . '\\majControl_majMonad_majMaybe_majTrans_monadmajThrowmajMaybemajT';

// Control_Monad_Maybe_Trans_monadErrorMaybeT
function majControl_majMonad_majMaybe_majTrans_monadmajErrormajMaybemajT($dictMonadError_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majMaybe_majTrans_monadmajErrormajMaybemajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictMonadError_0)->{'MonadThrow0'})(null);
  $Monad0_2_1 = (($__local_var_1_0)->{'Monad0'})(null);
  $__local_var_3_2 = (($__local_var_1_0)->{'Monad0'})(null);
  $monadMaybeT1_3_2 = (object)["Applicative0" => function($_dollar___unused_4) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__res = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_5) {
  $__num = \func_num_args();
  $__res = $x_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($__local_var_3_2)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_5) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__local_var_6_3 = (((((($__local_var_3_2)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_6_3 = (object)["map" => function($f_7) use ($__local_var_6_3) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($__local_var_6_3, $f_7) {
  $__num = \func_num_args();
  $__res = ((($__local_var_6_3)->{'map'})(function($v1_9) use ($f_7) {
  $__num = \func_num_args();
  $__t4 = null;;
  if ($v1_9 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t4 = new \Data\Maybe\Data_Maybe_Just(($f_7)(($v1_9)->{'value0'}));
goto end_branch_4;;
};
  $__t4 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_8);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_7_6 = (($__local_var_3_2)->{'Bind1'})(null);
  $Applicative0_8_7 = (($__local_var_3_2)->{'Applicative0'})(null);
  $Bind1_7_6 = (object)["bind" => function($v_9) use ($Applicative0_8_7, $Bind1_7_6) {
  $__num = \func_num_args();
  $__res = function($f_10) use ($Applicative0_8_7, $Bind1_7_6, $v_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_6)->{'bind'})($v_9))(function($v1_11) use ($Applicative0_8_7, $f_10) {
  $__num = \func_num_args();
  $__t8 = null;;
  if ($v1_11 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t8 = (($Applicative0_8_7)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_8;;
};
  if ($v1_11 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t8 = ($f_10)(($v1_11)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_9) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__local_var_10_9 = (((((($__local_var_3_2)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_10_9 = (object)["map" => function($f_11) use ($__local_var_10_9) {
  $__num = \func_num_args();
  $__res = function($v_12) use ($__local_var_10_9, $f_11) {
  $__num = \func_num_args();
  $__res = ((($__local_var_10_9)->{'map'})(function($v1_13) use ($f_11) {
  $__num = \func_num_args();
  $__t10 = null;;
  if ($v1_13 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t10 = new \Data\Maybe\Data_Maybe_Just(($f_11)(($v1_13)->{'value0'}));
goto end_branch_10;;
};
  $__t10 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_10:;
  $__res = $__t10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_12);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_11_12 = (($__local_var_3_2)->{'Bind1'})(null);
  $Applicative0_12_13 = (($__local_var_3_2)->{'Applicative0'})(null);
  $Bind1_11_12 = (object)["bind" => function($v_13) use ($Applicative0_12_13, $Bind1_11_12) {
  $__num = \func_num_args();
  $__res = function($f_14) use ($Applicative0_12_13, $Bind1_11_12, $v_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_12)->{'bind'})($v_13))(function($v1_15) use ($Applicative0_12_13, $f_14) {
  $__num = \func_num_args();
  $__t14 = null;;
  if ($v1_15 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t14 = (($Applicative0_12_13)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_14;;
};
  if ($v1_15 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t14 = ($f_14)(($v1_15)->{'value0'});
goto end_branch_14;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t14 = null;
  end_branch_14:;
  $__res = $__t14;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_13) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__local_var_14_15 = (((((($__local_var_3_2)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_14_15 = (object)["map" => function($f_15) use ($__local_var_14_15) {
  $__num = \func_num_args();
  $__res = function($v_16) use ($__local_var_14_15, $f_15) {
  $__num = \func_num_args();
  $__res = ((($__local_var_14_15)->{'map'})(function($v1_17) use ($f_15) {
  $__num = \func_num_args();
  $__t16 = null;;
  if ($v1_17 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t16 = new \Data\Maybe\Data_Maybe_Just(($f_15)(($v1_17)->{'value0'}));
goto end_branch_16;;
};
  $__t16 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_16:;
  $__res = $__t16;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_16);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_15_18 = (($__local_var_3_2)->{'Bind1'})(null);
  $Applicative0_16_19 = (($__local_var_3_2)->{'Applicative0'})(null);
  $Bind1_15_18 = (object)["bind" => function($v_17) use ($Applicative0_16_19, $Bind1_15_18) {
  $__num = \func_num_args();
  $__res = function($f_18) use ($Applicative0_16_19, $Bind1_15_18, $v_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_18)->{'bind'})($v_17))(function($v1_19) use ($Applicative0_16_19, $f_18) {
  $__num = \func_num_args();
  $__t20 = null;;
  if ($v1_19 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t20 = (($Applicative0_16_19)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_20;;
};
  if ($v1_19 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t20 = ($f_18)(($v1_19)->{'value0'});
goto end_branch_20;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t20 = null;
  end_branch_20:;
  $__res = $__t20;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_17) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($__local_var_3_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_16_22 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_16) {
  $__num = \func_num_args();
  $__res = $x_16;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($__local_var_3_2)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_16) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__local_var_17_22 = (((((($__local_var_3_2)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_17_22 = (object)["map" => function($f_18) use ($__local_var_17_22) {
  $__num = \func_num_args();
  $__res = function($v_19) use ($__local_var_17_22, $f_18) {
  $__num = \func_num_args();
  $__res = ((($__local_var_17_22)->{'map'})(function($v1_20) use ($f_18) {
  $__num = \func_num_args();
  $__t23 = null;;
  if ($v1_20 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t23 = new \Data\Maybe\Data_Maybe_Just(($f_18)(($v1_20)->{'value0'}));
goto end_branch_23;;
};
  $__t23 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_23:;
  $__res = $__t23;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_19);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_18_25 = (($__local_var_3_2)->{'Bind1'})(null);
  $Applicative0_19_26 = (($__local_var_3_2)->{'Applicative0'})(null);
  $Bind1_18_25 = (object)["bind" => function($v_20) use ($Applicative0_19_26, $Bind1_18_25) {
  $__num = \func_num_args();
  $__res = function($f_21) use ($Applicative0_19_26, $Bind1_18_25, $v_20) {
  $__num = \func_num_args();
  $__res = ((($Bind1_18_25)->{'bind'})($v_20))(function($v1_22) use ($Applicative0_19_26, $f_21) {
  $__num = \func_num_args();
  $__t27 = null;;
  if ($v1_22 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t27 = (($Applicative0_19_26)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_27;;
};
  if ($v1_22 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t27 = ($f_21)(($v1_22)->{'value0'});
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_20) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($__local_var_3_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_19_29 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($__local_var_3_2);
  $__res = (object)["apply" => function($f_20) use ($Applicative0_19_29, $Bind1_18_25) {
  $__num = \func_num_args();
  $__res = function($a_21) use ($Applicative0_19_29, $Bind1_18_25, $f_20) {
  $__num = \func_num_args();
  $__res = ((($Bind1_18_25)->{'bind'})($f_20))(function($f_prime__22) use ($Applicative0_19_29, $Bind1_18_25, $a_21) {
  $__num = \func_num_args();
  $__res = ((($Bind1_18_25)->{'bind'})($a_21))(function($a_prime__23) use ($Applicative0_19_29, $f_prime__22) {
  $__num = \func_num_args();
  $__res = (($Applicative0_19_29)->{'pure'})(($f_prime__22)($a_prime__23));
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
}, "Functor0" => function($_dollar___unused_18) use ($functorMaybeT1_17_22) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_17_22;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_17) use ($Applicative0_16_22, $Bind1_15_18) {
  $__num = \func_num_args();
  $__res = function($a_18) use ($Applicative0_16_22, $Bind1_15_18, $f_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_18)->{'bind'})($f_17))(function($f_prime__19) use ($Applicative0_16_22, $Bind1_15_18, $a_18) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_18)->{'bind'})($a_18))(function($a_prime__20) use ($Applicative0_16_22, $f_prime__19) {
  $__num = \func_num_args();
  $__res = (($Applicative0_16_22)->{'pure'})(($f_prime__19)($a_prime__20));
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
}, "Functor0" => function($_dollar___unused_15) use ($functorMaybeT1_14_15) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_14_15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_12_32 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_12) {
  $__num = \func_num_args();
  $__res = $x_12;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($__local_var_3_2)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_12) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__local_var_13_32 = (((((($__local_var_3_2)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_13_32 = (object)["map" => function($f_14) use ($__local_var_13_32) {
  $__num = \func_num_args();
  $__res = function($v_15) use ($__local_var_13_32, $f_14) {
  $__num = \func_num_args();
  $__res = ((($__local_var_13_32)->{'map'})(function($v1_16) use ($f_14) {
  $__num = \func_num_args();
  $__t33 = null;;
  if ($v1_16 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t33 = new \Data\Maybe\Data_Maybe_Just(($f_14)(($v1_16)->{'value0'}));
goto end_branch_33;;
};
  $__t33 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_33:;
  $__res = $__t33;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_15);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_14_35 = (($__local_var_3_2)->{'Bind1'})(null);
  $Applicative0_15_36 = (($__local_var_3_2)->{'Applicative0'})(null);
  $Bind1_14_35 = (object)["bind" => function($v_16) use ($Applicative0_15_36, $Bind1_14_35) {
  $__num = \func_num_args();
  $__res = function($f_17) use ($Applicative0_15_36, $Bind1_14_35, $v_16) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_35)->{'bind'})($v_16))(function($v1_18) use ($Applicative0_15_36, $f_17) {
  $__num = \func_num_args();
  $__t37 = null;;
  if ($v1_18 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t37 = (($Applicative0_15_36)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_37;;
};
  if ($v1_18 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t37 = ($f_17)(($v1_18)->{'value0'});
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_16) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__local_var_17_38 = (((((($__local_var_3_2)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_17_38 = (object)["map" => function($f_18) use ($__local_var_17_38) {
  $__num = \func_num_args();
  $__res = function($v_19) use ($__local_var_17_38, $f_18) {
  $__num = \func_num_args();
  $__res = ((($__local_var_17_38)->{'map'})(function($v1_20) use ($f_18) {
  $__num = \func_num_args();
  $__t39 = null;;
  if ($v1_20 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t39 = new \Data\Maybe\Data_Maybe_Just(($f_18)(($v1_20)->{'value0'}));
goto end_branch_39;;
};
  $__t39 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_39:;
  $__res = $__t39;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_19);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_18_41 = (($__local_var_3_2)->{'Bind1'})(null);
  $Applicative0_19_42 = (($__local_var_3_2)->{'Applicative0'})(null);
  $Bind1_18_41 = (object)["bind" => function($v_20) use ($Applicative0_19_42, $Bind1_18_41) {
  $__num = \func_num_args();
  $__res = function($f_21) use ($Applicative0_19_42, $Bind1_18_41, $v_20) {
  $__num = \func_num_args();
  $__res = ((($Bind1_18_41)->{'bind'})($v_20))(function($v1_22) use ($Applicative0_19_42, $f_21) {
  $__num = \func_num_args();
  $__t43 = null;;
  if ($v1_22 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t43 = (($Applicative0_19_42)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_43;;
};
  if ($v1_22 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t43 = ($f_21)(($v1_22)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_20) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($__local_var_3_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_19_45 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($__local_var_3_2);
  $__res = (object)["apply" => function($f_20) use ($Applicative0_19_45, $Bind1_18_41) {
  $__num = \func_num_args();
  $__res = function($a_21) use ($Applicative0_19_45, $Bind1_18_41, $f_20) {
  $__num = \func_num_args();
  $__res = ((($Bind1_18_41)->{'bind'})($f_20))(function($f_prime__22) use ($Applicative0_19_45, $Bind1_18_41, $a_21) {
  $__num = \func_num_args();
  $__res = ((($Bind1_18_41)->{'bind'})($a_21))(function($a_prime__23) use ($Applicative0_19_45, $f_prime__22) {
  $__num = \func_num_args();
  $__res = (($Applicative0_19_45)->{'pure'})(($f_prime__22)($a_prime__23));
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
}, "Functor0" => function($_dollar___unused_18) use ($functorMaybeT1_17_38) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_17_38;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_15_47 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($__local_var_3_2);
  $__res = (object)["apply" => function($f_16) use ($Applicative0_15_47, $Bind1_14_35) {
  $__num = \func_num_args();
  $__res = function($a_17) use ($Applicative0_15_47, $Bind1_14_35, $f_16) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_35)->{'bind'})($f_16))(function($f_prime__18) use ($Applicative0_15_47, $Bind1_14_35, $a_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_35)->{'bind'})($a_17))(function($a_prime__19) use ($Applicative0_15_47, $f_prime__18) {
  $__num = \func_num_args();
  $__res = (($Applicative0_15_47)->{'pure'})(($f_prime__18)($a_prime__19));
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
}, "Functor0" => function($_dollar___unused_14) use ($functorMaybeT1_13_32) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_13_32;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_13) use ($Applicative0_12_32, $Bind1_11_12) {
  $__num = \func_num_args();
  $__res = function($a_14) use ($Applicative0_12_32, $Bind1_11_12, $f_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_12)->{'bind'})($f_13))(function($f_prime__15) use ($Applicative0_12_32, $Bind1_11_12, $a_14) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_12)->{'bind'})($a_14))(function($a_prime__16) use ($Applicative0_12_32, $f_prime__15) {
  $__num = \func_num_args();
  $__res = (($Applicative0_12_32)->{'pure'})(($f_prime__15)($a_prime__16));
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
}, "Functor0" => function($_dollar___unused_11) use ($functorMaybeT1_10_9) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_10_9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_8_50 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_8) {
  $__num = \func_num_args();
  $__res = $x_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($__local_var_3_2)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_8) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__local_var_9_50 = (((((($__local_var_3_2)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_9_50 = (object)["map" => function($f_10) use ($__local_var_9_50) {
  $__num = \func_num_args();
  $__res = function($v_11) use ($__local_var_9_50, $f_10) {
  $__num = \func_num_args();
  $__res = ((($__local_var_9_50)->{'map'})(function($v1_12) use ($f_10) {
  $__num = \func_num_args();
  $__t51 = null;;
  if ($v1_12 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t51 = new \Data\Maybe\Data_Maybe_Just(($f_10)(($v1_12)->{'value0'}));
goto end_branch_51;;
};
  $__t51 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_51:;
  $__res = $__t51;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_11);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_10_53 = (($__local_var_3_2)->{'Bind1'})(null);
  $Applicative0_11_54 = (($__local_var_3_2)->{'Applicative0'})(null);
  $Bind1_10_53 = (object)["bind" => function($v_12) use ($Applicative0_11_54, $Bind1_10_53) {
  $__num = \func_num_args();
  $__res = function($f_13) use ($Applicative0_11_54, $Bind1_10_53, $v_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_53)->{'bind'})($v_12))(function($v1_14) use ($Applicative0_11_54, $f_13) {
  $__num = \func_num_args();
  $__t55 = null;;
  if ($v1_14 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t55 = (($Applicative0_11_54)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_55;;
};
  if ($v1_14 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t55 = ($f_13)(($v1_14)->{'value0'});
goto end_branch_55;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t55 = null;
  end_branch_55:;
  $__res = $__t55;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_12) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__local_var_13_56 = (((((($__local_var_3_2)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_13_56 = (object)["map" => function($f_14) use ($__local_var_13_56) {
  $__num = \func_num_args();
  $__res = function($v_15) use ($__local_var_13_56, $f_14) {
  $__num = \func_num_args();
  $__res = ((($__local_var_13_56)->{'map'})(function($v1_16) use ($f_14) {
  $__num = \func_num_args();
  $__t57 = null;;
  if ($v1_16 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t57 = new \Data\Maybe\Data_Maybe_Just(($f_14)(($v1_16)->{'value0'}));
goto end_branch_57;;
};
  $__t57 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_57:;
  $__res = $__t57;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_15);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_14_59 = (($__local_var_3_2)->{'Bind1'})(null);
  $Applicative0_15_60 = (($__local_var_3_2)->{'Applicative0'})(null);
  $Bind1_14_59 = (object)["bind" => function($v_16) use ($Applicative0_15_60, $Bind1_14_59) {
  $__num = \func_num_args();
  $__res = function($f_17) use ($Applicative0_15_60, $Bind1_14_59, $v_16) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_59)->{'bind'})($v_16))(function($v1_18) use ($Applicative0_15_60, $f_17) {
  $__num = \func_num_args();
  $__t61 = null;;
  if ($v1_18 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t61 = (($Applicative0_15_60)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_61;;
};
  if ($v1_18 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t61 = ($f_17)(($v1_18)->{'value0'});
goto end_branch_61;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t61 = null;
  end_branch_61:;
  $__res = $__t61;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_16) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($__local_var_3_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_15_63 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($__local_var_3_2);
  $__res = (object)["apply" => function($f_16) use ($Applicative0_15_63, $Bind1_14_59) {
  $__num = \func_num_args();
  $__res = function($a_17) use ($Applicative0_15_63, $Bind1_14_59, $f_16) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_59)->{'bind'})($f_16))(function($f_prime__18) use ($Applicative0_15_63, $Bind1_14_59, $a_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_14_59)->{'bind'})($a_17))(function($a_prime__19) use ($Applicative0_15_63, $f_prime__18) {
  $__num = \func_num_args();
  $__res = (($Applicative0_15_63)->{'pure'})(($f_prime__18)($a_prime__19));
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
}, "Functor0" => function($_dollar___unused_14) use ($functorMaybeT1_13_56) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_13_56;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_11_65 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($__local_var_3_2);
  $__res = (object)["apply" => function($f_12) use ($Applicative0_11_65, $Bind1_10_53) {
  $__num = \func_num_args();
  $__res = function($a_13) use ($Applicative0_11_65, $Bind1_10_53, $f_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_53)->{'bind'})($f_12))(function($f_prime__14) use ($Applicative0_11_65, $Bind1_10_53, $a_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_53)->{'bind'})($a_13))(function($a_prime__15) use ($Applicative0_11_65, $f_prime__14) {
  $__num = \func_num_args();
  $__res = (($Applicative0_11_65)->{'pure'})(($f_prime__14)($a_prime__15));
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
}, "Functor0" => function($_dollar___unused_10) use ($functorMaybeT1_9_50) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_9_50;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_9) use ($Applicative0_8_50, $Bind1_7_6) {
  $__num = \func_num_args();
  $__res = function($a_10) use ($Applicative0_8_50, $Bind1_7_6, $f_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_6)->{'bind'})($f_9))(function($f_prime__11) use ($Applicative0_8_50, $Bind1_7_6, $a_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_6)->{'bind'})($a_10))(function($a_prime__12) use ($Applicative0_8_50, $f_prime__11) {
  $__num = \func_num_args();
  $__res = (($Applicative0_8_50)->{'pure'})(($f_prime__11)($a_prime__12));
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
}, "Functor0" => function($_dollar___unused_7) use ($functorMaybeT1_6_3) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_6_3;
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
  $Bind1_5_67 = (($__local_var_3_2)->{'Bind1'})(null);
  $Applicative0_6_68 = (($__local_var_3_2)->{'Applicative0'})(null);
  $__res = (object)["bind" => function($v_7) use ($Applicative0_6_68, $Bind1_5_67) {
  $__num = \func_num_args();
  $__res = function($f_8) use ($Applicative0_6_68, $Bind1_5_67, $v_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_5_67)->{'bind'})($v_7))(function($v1_9) use ($Applicative0_6_68, $f_8) {
  $__num = \func_num_args();
  $__t69 = null;;
  if ($v1_9 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t69 = (($Applicative0_6_68)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_69;;
};
  if ($v1_9 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t69 = ($f_8)(($v1_9)->{'value0'});
goto end_branch_69;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t69 = null;
  end_branch_69:;
  $__res = $__t69;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_7) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__local_var_8_70 = (((((($__local_var_3_2)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_8_70 = (object)["map" => function($f_9) use ($__local_var_8_70) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($__local_var_8_70, $f_9) {
  $__num = \func_num_args();
  $__res = ((($__local_var_8_70)->{'map'})(function($v1_11) use ($f_9) {
  $__num = \func_num_args();
  $__t71 = null;;
  if ($v1_11 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t71 = new \Data\Maybe\Data_Maybe_Just(($f_9)(($v1_11)->{'value0'}));
goto end_branch_71;;
};
  $__t71 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_71:;
  $__res = $__t71;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_10);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_9_73 = (($__local_var_3_2)->{'Bind1'})(null);
  $Applicative0_10_74 = (($__local_var_3_2)->{'Applicative0'})(null);
  $Bind1_9_73 = (object)["bind" => function($v_11) use ($Applicative0_10_74, $Bind1_9_73) {
  $__num = \func_num_args();
  $__res = function($f_12) use ($Applicative0_10_74, $Bind1_9_73, $v_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_73)->{'bind'})($v_11))(function($v1_13) use ($Applicative0_10_74, $f_12) {
  $__num = \func_num_args();
  $__t75 = null;;
  if ($v1_13 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t75 = (($Applicative0_10_74)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_75;;
};
  if ($v1_13 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t75 = ($f_12)(($v1_13)->{'value0'});
goto end_branch_75;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t75 = null;
  end_branch_75:;
  $__res = $__t75;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_11) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__local_var_12_76 = (((((($__local_var_3_2)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_12_76 = (object)["map" => function($f_13) use ($__local_var_12_76) {
  $__num = \func_num_args();
  $__res = function($v_14) use ($__local_var_12_76, $f_13) {
  $__num = \func_num_args();
  $__res = ((($__local_var_12_76)->{'map'})(function($v1_15) use ($f_13) {
  $__num = \func_num_args();
  $__t77 = null;;
  if ($v1_15 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t77 = new \Data\Maybe\Data_Maybe_Just(($f_13)(($v1_15)->{'value0'}));
goto end_branch_77;;
};
  $__t77 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_77:;
  $__res = $__t77;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_14);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_13_79 = (($__local_var_3_2)->{'Bind1'})(null);
  $Applicative0_14_80 = (($__local_var_3_2)->{'Applicative0'})(null);
  $Bind1_13_79 = (object)["bind" => function($v_15) use ($Applicative0_14_80, $Bind1_13_79) {
  $__num = \func_num_args();
  $__res = function($f_16) use ($Applicative0_14_80, $Bind1_13_79, $v_15) {
  $__num = \func_num_args();
  $__res = ((($Bind1_13_79)->{'bind'})($v_15))(function($v1_17) use ($Applicative0_14_80, $f_16) {
  $__num = \func_num_args();
  $__t81 = null;;
  if ($v1_17 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t81 = (($Applicative0_14_80)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_81;;
};
  if ($v1_17 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t81 = ($f_16)(($v1_17)->{'value0'});
goto end_branch_81;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t81 = null;
  end_branch_81:;
  $__res = $__t81;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_15) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($__local_var_3_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_14_83 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_14) {
  $__num = \func_num_args();
  $__res = $x_14;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($__local_var_3_2)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_14) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__local_var_15_83 = (((((($__local_var_3_2)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_15_83 = (object)["map" => function($f_16) use ($__local_var_15_83) {
  $__num = \func_num_args();
  $__res = function($v_17) use ($__local_var_15_83, $f_16) {
  $__num = \func_num_args();
  $__res = ((($__local_var_15_83)->{'map'})(function($v1_18) use ($f_16) {
  $__num = \func_num_args();
  $__t84 = null;;
  if ($v1_18 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t84 = new \Data\Maybe\Data_Maybe_Just(($f_16)(($v1_18)->{'value0'}));
goto end_branch_84;;
};
  $__t84 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_84:;
  $__res = $__t84;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_17);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_16_86 = (($__local_var_3_2)->{'Bind1'})(null);
  $Applicative0_17_87 = (($__local_var_3_2)->{'Applicative0'})(null);
  $Bind1_16_86 = (object)["bind" => function($v_18) use ($Applicative0_17_87, $Bind1_16_86) {
  $__num = \func_num_args();
  $__res = function($f_19) use ($Applicative0_17_87, $Bind1_16_86, $v_18) {
  $__num = \func_num_args();
  $__res = ((($Bind1_16_86)->{'bind'})($v_18))(function($v1_20) use ($Applicative0_17_87, $f_19) {
  $__num = \func_num_args();
  $__t88 = null;;
  if ($v1_20 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t88 = (($Applicative0_17_87)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_88;;
};
  if ($v1_20 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t88 = ($f_19)(($v1_20)->{'value0'});
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_18) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($__local_var_3_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_17_90 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($__local_var_3_2);
  $__res = (object)["apply" => function($f_18) use ($Applicative0_17_90, $Bind1_16_86) {
  $__num = \func_num_args();
  $__res = function($a_19) use ($Applicative0_17_90, $Bind1_16_86, $f_18) {
  $__num = \func_num_args();
  $__res = ((($Bind1_16_86)->{'bind'})($f_18))(function($f_prime__20) use ($Applicative0_17_90, $Bind1_16_86, $a_19) {
  $__num = \func_num_args();
  $__res = ((($Bind1_16_86)->{'bind'})($a_19))(function($a_prime__21) use ($Applicative0_17_90, $f_prime__20) {
  $__num = \func_num_args();
  $__res = (($Applicative0_17_90)->{'pure'})(($f_prime__20)($a_prime__21));
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
}, "Functor0" => function($_dollar___unused_16) use ($functorMaybeT1_15_83) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_15_83;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_15) use ($Applicative0_14_83, $Bind1_13_79) {
  $__num = \func_num_args();
  $__res = function($a_16) use ($Applicative0_14_83, $Bind1_13_79, $f_15) {
  $__num = \func_num_args();
  $__res = ((($Bind1_13_79)->{'bind'})($f_15))(function($f_prime__17) use ($Applicative0_14_83, $Bind1_13_79, $a_16) {
  $__num = \func_num_args();
  $__res = ((($Bind1_13_79)->{'bind'})($a_16))(function($a_prime__18) use ($Applicative0_14_83, $f_prime__17) {
  $__num = \func_num_args();
  $__res = (($Applicative0_14_83)->{'pure'})(($f_prime__17)($a_prime__18));
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
}, "Functor0" => function($_dollar___unused_13) use ($functorMaybeT1_12_76) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_12_76;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_10_93 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_10) {
  $__num = \func_num_args();
  $__res = $x_10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($__local_var_3_2)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_10) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__local_var_11_93 = (((((($__local_var_3_2)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_11_93 = (object)["map" => function($f_12) use ($__local_var_11_93) {
  $__num = \func_num_args();
  $__res = function($v_13) use ($__local_var_11_93, $f_12) {
  $__num = \func_num_args();
  $__res = ((($__local_var_11_93)->{'map'})(function($v1_14) use ($f_12) {
  $__num = \func_num_args();
  $__t94 = null;;
  if ($v1_14 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t94 = new \Data\Maybe\Data_Maybe_Just(($f_12)(($v1_14)->{'value0'}));
goto end_branch_94;;
};
  $__t94 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_94:;
  $__res = $__t94;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_13);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_12_96 = (($__local_var_3_2)->{'Bind1'})(null);
  $Applicative0_13_97 = (($__local_var_3_2)->{'Applicative0'})(null);
  $Bind1_12_96 = (object)["bind" => function($v_14) use ($Applicative0_13_97, $Bind1_12_96) {
  $__num = \func_num_args();
  $__res = function($f_15) use ($Applicative0_13_97, $Bind1_12_96, $v_14) {
  $__num = \func_num_args();
  $__res = ((($Bind1_12_96)->{'bind'})($v_14))(function($v1_16) use ($Applicative0_13_97, $f_15) {
  $__num = \func_num_args();
  $__t98 = null;;
  if ($v1_16 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t98 = (($Applicative0_13_97)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_98;;
};
  if ($v1_16 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t98 = ($f_15)(($v1_16)->{'value0'});
goto end_branch_98;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t98 = null;
  end_branch_98:;
  $__res = $__t98;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_14) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__local_var_15_99 = (((((($__local_var_3_2)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_15_99 = (object)["map" => function($f_16) use ($__local_var_15_99) {
  $__num = \func_num_args();
  $__res = function($v_17) use ($__local_var_15_99, $f_16) {
  $__num = \func_num_args();
  $__res = ((($__local_var_15_99)->{'map'})(function($v1_18) use ($f_16) {
  $__num = \func_num_args();
  $__t100 = null;;
  if ($v1_18 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t100 = new \Data\Maybe\Data_Maybe_Just(($f_16)(($v1_18)->{'value0'}));
goto end_branch_100;;
};
  $__t100 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_100:;
  $__res = $__t100;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_17);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_16_102 = (($__local_var_3_2)->{'Bind1'})(null);
  $Applicative0_17_103 = (($__local_var_3_2)->{'Applicative0'})(null);
  $Bind1_16_102 = (object)["bind" => function($v_18) use ($Applicative0_17_103, $Bind1_16_102) {
  $__num = \func_num_args();
  $__res = function($f_19) use ($Applicative0_17_103, $Bind1_16_102, $v_18) {
  $__num = \func_num_args();
  $__res = ((($Bind1_16_102)->{'bind'})($v_18))(function($v1_20) use ($Applicative0_17_103, $f_19) {
  $__num = \func_num_args();
  $__t104 = null;;
  if ($v1_20 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t104 = (($Applicative0_17_103)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_104;;
};
  if ($v1_20 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t104 = ($f_19)(($v1_20)->{'value0'});
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_18) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($__local_var_3_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_17_106 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_17) {
  $__num = \func_num_args();
  $__res = $x_17;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($__local_var_3_2)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_17) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__local_var_18_106 = (((((($__local_var_3_2)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_18_106 = (object)["map" => function($f_19) use ($__local_var_18_106) {
  $__num = \func_num_args();
  $__res = function($v_20) use ($__local_var_18_106, $f_19) {
  $__num = \func_num_args();
  $__res = ((($__local_var_18_106)->{'map'})(function($v1_21) use ($f_19) {
  $__num = \func_num_args();
  $__t107 = null;;
  if ($v1_21 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t107 = new \Data\Maybe\Data_Maybe_Just(($f_19)(($v1_21)->{'value0'}));
goto end_branch_107;;
};
  $__t107 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_107:;
  $__res = $__t107;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_20);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_19_109 = (($__local_var_3_2)->{'Bind1'})(null);
  $Applicative0_20_110 = (($__local_var_3_2)->{'Applicative0'})(null);
  $Bind1_19_109 = (object)["bind" => function($v_21) use ($Applicative0_20_110, $Bind1_19_109) {
  $__num = \func_num_args();
  $__res = function($f_22) use ($Applicative0_20_110, $Bind1_19_109, $v_21) {
  $__num = \func_num_args();
  $__res = ((($Bind1_19_109)->{'bind'})($v_21))(function($v1_23) use ($Applicative0_20_110, $f_22) {
  $__num = \func_num_args();
  $__t111 = null;;
  if ($v1_23 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t111 = (($Applicative0_20_110)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_111;;
};
  if ($v1_23 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t111 = ($f_22)(($v1_23)->{'value0'});
goto end_branch_111;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t111 = null;
  end_branch_111:;
  $__res = $__t111;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_21) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($__local_var_3_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_20_113 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($__local_var_3_2);
  $__res = (object)["apply" => function($f_21) use ($Applicative0_20_113, $Bind1_19_109) {
  $__num = \func_num_args();
  $__res = function($a_22) use ($Applicative0_20_113, $Bind1_19_109, $f_21) {
  $__num = \func_num_args();
  $__res = ((($Bind1_19_109)->{'bind'})($f_21))(function($f_prime__23) use ($Applicative0_20_113, $Bind1_19_109, $a_22) {
  $__num = \func_num_args();
  $__res = ((($Bind1_19_109)->{'bind'})($a_22))(function($a_prime__24) use ($Applicative0_20_113, $f_prime__23) {
  $__num = \func_num_args();
  $__res = (($Applicative0_20_113)->{'pure'})(($f_prime__23)($a_prime__24));
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
}, "Functor0" => function($_dollar___unused_19) use ($functorMaybeT1_18_106) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_18_106;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_18) use ($Applicative0_17_106, $Bind1_16_102) {
  $__num = \func_num_args();
  $__res = function($a_19) use ($Applicative0_17_106, $Bind1_16_102, $f_18) {
  $__num = \func_num_args();
  $__res = ((($Bind1_16_102)->{'bind'})($f_18))(function($f_prime__20) use ($Applicative0_17_106, $Bind1_16_102, $a_19) {
  $__num = \func_num_args();
  $__res = ((($Bind1_16_102)->{'bind'})($a_19))(function($a_prime__21) use ($Applicative0_17_106, $f_prime__20) {
  $__num = \func_num_args();
  $__res = (($Applicative0_17_106)->{'pure'})(($f_prime__20)($a_prime__21));
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
}, "Functor0" => function($_dollar___unused_16) use ($functorMaybeT1_15_99) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_15_99;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_13_116 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_13) {
  $__num = \func_num_args();
  $__res = $x_13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($__local_var_3_2)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_13) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__local_var_14_116 = (((((($__local_var_3_2)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_14_116 = (object)["map" => function($f_15) use ($__local_var_14_116) {
  $__num = \func_num_args();
  $__res = function($v_16) use ($__local_var_14_116, $f_15) {
  $__num = \func_num_args();
  $__res = ((($__local_var_14_116)->{'map'})(function($v1_17) use ($f_15) {
  $__num = \func_num_args();
  $__t117 = null;;
  if ($v1_17 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t117 = new \Data\Maybe\Data_Maybe_Just(($f_15)(($v1_17)->{'value0'}));
goto end_branch_117;;
};
  $__t117 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_117:;
  $__res = $__t117;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_16);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_15_119 = (($__local_var_3_2)->{'Bind1'})(null);
  $Applicative0_16_120 = (($__local_var_3_2)->{'Applicative0'})(null);
  $Bind1_15_119 = (object)["bind" => function($v_17) use ($Applicative0_16_120, $Bind1_15_119) {
  $__num = \func_num_args();
  $__res = function($f_18) use ($Applicative0_16_120, $Bind1_15_119, $v_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_119)->{'bind'})($v_17))(function($v1_19) use ($Applicative0_16_120, $f_18) {
  $__num = \func_num_args();
  $__t121 = null;;
  if ($v1_19 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t121 = (($Applicative0_16_120)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_121;;
};
  if ($v1_19 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t121 = ($f_18)(($v1_19)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_17) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__local_var_18_122 = (((((($__local_var_3_2)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_18_122 = (object)["map" => function($f_19) use ($__local_var_18_122) {
  $__num = \func_num_args();
  $__res = function($v_20) use ($__local_var_18_122, $f_19) {
  $__num = \func_num_args();
  $__res = ((($__local_var_18_122)->{'map'})(function($v1_21) use ($f_19) {
  $__num = \func_num_args();
  $__t123 = null;;
  if ($v1_21 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t123 = new \Data\Maybe\Data_Maybe_Just(($f_19)(($v1_21)->{'value0'}));
goto end_branch_123;;
};
  $__t123 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_123:;
  $__res = $__t123;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_20);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_19_125 = (($__local_var_3_2)->{'Bind1'})(null);
  $Applicative0_20_126 = (($__local_var_3_2)->{'Applicative0'})(null);
  $Bind1_19_125 = (object)["bind" => function($v_21) use ($Applicative0_20_126, $Bind1_19_125) {
  $__num = \func_num_args();
  $__res = function($f_22) use ($Applicative0_20_126, $Bind1_19_125, $v_21) {
  $__num = \func_num_args();
  $__res = ((($Bind1_19_125)->{'bind'})($v_21))(function($v1_23) use ($Applicative0_20_126, $f_22) {
  $__num = \func_num_args();
  $__t127 = null;;
  if ($v1_23 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t127 = (($Applicative0_20_126)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_127;;
};
  if ($v1_23 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t127 = ($f_22)(($v1_23)->{'value0'});
goto end_branch_127;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t127 = null;
  end_branch_127:;
  $__res = $__t127;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_21) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($__local_var_3_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_20_129 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($__local_var_3_2);
  $__res = (object)["apply" => function($f_21) use ($Applicative0_20_129, $Bind1_19_125) {
  $__num = \func_num_args();
  $__res = function($a_22) use ($Applicative0_20_129, $Bind1_19_125, $f_21) {
  $__num = \func_num_args();
  $__res = ((($Bind1_19_125)->{'bind'})($f_21))(function($f_prime__23) use ($Applicative0_20_129, $Bind1_19_125, $a_22) {
  $__num = \func_num_args();
  $__res = ((($Bind1_19_125)->{'bind'})($a_22))(function($a_prime__24) use ($Applicative0_20_129, $f_prime__23) {
  $__num = \func_num_args();
  $__res = (($Applicative0_20_129)->{'pure'})(($f_prime__23)($a_prime__24));
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
}, "Functor0" => function($_dollar___unused_19) use ($functorMaybeT1_18_122) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_18_122;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_16_131 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($__local_var_3_2);
  $__res = (object)["apply" => function($f_17) use ($Applicative0_16_131, $Bind1_15_119) {
  $__num = \func_num_args();
  $__res = function($a_18) use ($Applicative0_16_131, $Bind1_15_119, $f_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_119)->{'bind'})($f_17))(function($f_prime__19) use ($Applicative0_16_131, $Bind1_15_119, $a_18) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_119)->{'bind'})($a_18))(function($a_prime__20) use ($Applicative0_16_131, $f_prime__19) {
  $__num = \func_num_args();
  $__res = (($Applicative0_16_131)->{'pure'})(($f_prime__19)($a_prime__20));
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
}, "Functor0" => function($_dollar___unused_15) use ($functorMaybeT1_14_116) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_14_116;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_14) use ($Applicative0_13_116, $Bind1_12_96) {
  $__num = \func_num_args();
  $__res = function($a_15) use ($Applicative0_13_116, $Bind1_12_96, $f_14) {
  $__num = \func_num_args();
  $__res = ((($Bind1_12_96)->{'bind'})($f_14))(function($f_prime__16) use ($Applicative0_13_116, $Bind1_12_96, $a_15) {
  $__num = \func_num_args();
  $__res = ((($Bind1_12_96)->{'bind'})($a_15))(function($a_prime__17) use ($Applicative0_13_116, $f_prime__16) {
  $__num = \func_num_args();
  $__res = (($Applicative0_13_116)->{'pure'})(($f_prime__16)($a_prime__17));
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
}, "Functor0" => function($_dollar___unused_12) use ($functorMaybeT1_11_93) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_11_93;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_11) use ($Applicative0_10_93, $Bind1_9_73) {
  $__num = \func_num_args();
  $__res = function($a_12) use ($Applicative0_10_93, $Bind1_9_73, $f_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_73)->{'bind'})($f_11))(function($f_prime__13) use ($Applicative0_10_93, $Bind1_9_73, $a_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_73)->{'bind'})($a_12))(function($a_prime__14) use ($Applicative0_10_93, $f_prime__13) {
  $__num = \func_num_args();
  $__res = (($Applicative0_10_93)->{'pure'})(($f_prime__13)($a_prime__14));
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
}, "Functor0" => function($_dollar___unused_9) use ($functorMaybeT1_8_70) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_8_70;
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
  $monadThrowMaybeT1_1_0 = (object)["throwError" => function($e_4) use ($Monad0_2_1, $__local_var_1_0) {
  $__num = \func_num_args();
  $Bind1_5_135 = (($Monad0_2_1)->{'Bind1'})(null);
  $Applicative0_6_136 = (($Monad0_2_1)->{'Applicative0'})(null);
  $__res = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($x_5) {
  $__num = \func_num_args();
  $__res = $x_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, function($a_7) use ($Applicative0_6_136, $Bind1_5_135) {
  $__num = \func_num_args();
  $__res = ((($Bind1_5_135)->{'bind'})($a_7))(function($a_prime__8) use ($Applicative0_6_136) {
  $__num = \func_num_args();
  $__res = (($Applicative0_6_136)->{'pure'})(new \Data\Maybe\Data_Maybe_Just($a_prime__8));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (($__local_var_1_0)->{'throwError'})($e_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad0" => function($_dollar___unused_4) use ($monadMaybeT1_3_2) {
  $__num = \func_num_args();
  $__res = $monadMaybeT1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["catchError" => function($v_2) use ($dictMonadError_0) {
  $__num = \func_num_args();
  $__res = function($h_3) use ($dictMonadError_0, $v_2) {
  $__num = \func_num_args();
  $__res = ((($dictMonadError_0)->{'catchError'})($v_2))(function($a_4) use ($h_3) {
  $__num = \func_num_args();
  $__res = ($h_3)($a_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "MonadThrow0" => function($_dollar___unused_2) use ($monadThrowMaybeT1_1_0) {
  $__num = \func_num_args();
  $__res = $monadThrowMaybeT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Maybe_Trans_monadErrorMaybeT'] = __NAMESPACE__ . '\\majControl_majMonad_majMaybe_majTrans_monadmajErrormajMaybemajT';

// Control_Monad_Maybe_Trans_monadSTMaybeT
function majControl_majMonad_majMaybe_majTrans_monadmajSmajTmajMaybemajT($dictMonadST_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majMaybe_majTrans_monadmajSmajTmajMaybemajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Monad0_1_0 = (($dictMonadST_0)->{'Monad0'})(null);
  $monadMaybeT1_2_1 = (object)["Applicative0" => function($_dollar___unused_2) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_3) {
  $__num = \func_num_args();
  $__res = $x_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($Monad0_1_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_3) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__local_var_4_1 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_4_1 = (object)["map" => function($f_5) use ($__local_var_4_1) {
  $__num = \func_num_args();
  $__res = function($v_6) use ($__local_var_4_1, $f_5) {
  $__num = \func_num_args();
  $__res = ((($__local_var_4_1)->{'map'})(function($v1_7) use ($f_5) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ($v1_7 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t2 = new \Data\Maybe\Data_Maybe_Just(($f_5)(($v1_7)->{'value0'}));
goto end_branch_2;;
};
  $__t2 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_5_4 = (($Monad0_1_0)->{'Bind1'})(null);
  $Applicative0_6_5 = (($Monad0_1_0)->{'Applicative0'})(null);
  $Bind1_5_4 = (object)["bind" => function($v_7) use ($Applicative0_6_5, $Bind1_5_4) {
  $__num = \func_num_args();
  $__res = function($f_8) use ($Applicative0_6_5, $Bind1_5_4, $v_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_5_4)->{'bind'})($v_7))(function($v1_9) use ($Applicative0_6_5, $f_8) {
  $__num = \func_num_args();
  $__t6 = null;;
  if ($v1_9 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t6 = (($Applicative0_6_5)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_6;;
};
  if ($v1_9 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t6 = ($f_8)(($v1_9)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_7) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__local_var_8_7 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_8_7 = (object)["map" => function($f_9) use ($__local_var_8_7) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($__local_var_8_7, $f_9) {
  $__num = \func_num_args();
  $__res = ((($__local_var_8_7)->{'map'})(function($v1_11) use ($f_9) {
  $__num = \func_num_args();
  $__t8 = null;;
  if ($v1_11 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t8 = new \Data\Maybe\Data_Maybe_Just(($f_9)(($v1_11)->{'value0'}));
goto end_branch_8;;
};
  $__t8 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_8:;
  $__res = $__t8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_10);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_9_10 = (($Monad0_1_0)->{'Bind1'})(null);
  $Applicative0_10_11 = (($Monad0_1_0)->{'Applicative0'})(null);
  $Bind1_9_10 = (object)["bind" => function($v_11) use ($Applicative0_10_11, $Bind1_9_10) {
  $__num = \func_num_args();
  $__res = function($f_12) use ($Applicative0_10_11, $Bind1_9_10, $v_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_10)->{'bind'})($v_11))(function($v1_13) use ($Applicative0_10_11, $f_12) {
  $__num = \func_num_args();
  $__t12 = null;;
  if ($v1_13 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t12 = (($Applicative0_10_11)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_12;;
};
  if ($v1_13 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t12 = ($f_12)(($v1_13)->{'value0'});
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_11) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_10_14 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($Monad0_1_0);
  $__res = (object)["apply" => function($f_11) use ($Applicative0_10_14, $Bind1_9_10) {
  $__num = \func_num_args();
  $__res = function($a_12) use ($Applicative0_10_14, $Bind1_9_10, $f_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_10)->{'bind'})($f_11))(function($f_prime__13) use ($Applicative0_10_14, $Bind1_9_10, $a_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_10)->{'bind'})($a_12))(function($a_prime__14) use ($Applicative0_10_14, $f_prime__13) {
  $__num = \func_num_args();
  $__res = (($Applicative0_10_14)->{'pure'})(($f_prime__13)($a_prime__14));
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
}, "Functor0" => function($_dollar___unused_9) use ($functorMaybeT1_8_7) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_8_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_6_16 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($Monad0_1_0);
  $__res = (object)["apply" => function($f_7) use ($Applicative0_6_16, $Bind1_5_4) {
  $__num = \func_num_args();
  $__res = function($a_8) use ($Applicative0_6_16, $Bind1_5_4, $f_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_5_4)->{'bind'})($f_7))(function($f_prime__9) use ($Applicative0_6_16, $Bind1_5_4, $a_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_5_4)->{'bind'})($a_8))(function($a_prime__10) use ($Applicative0_6_16, $f_prime__9) {
  $__num = \func_num_args();
  $__res = (($Applicative0_6_16)->{'pure'})(($f_prime__9)($a_prime__10));
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
}, "Functor0" => function($_dollar___unused_5) use ($functorMaybeT1_4_1) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_4_1;
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
  $Bind1_3_17 = (($Monad0_1_0)->{'Bind1'})(null);
  $Applicative0_4_18 = (($Monad0_1_0)->{'Applicative0'})(null);
  $__res = (object)["bind" => function($v_5) use ($Applicative0_4_18, $Bind1_3_17) {
  $__num = \func_num_args();
  $__res = function($f_6) use ($Applicative0_4_18, $Bind1_3_17, $v_5) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_17)->{'bind'})($v_5))(function($v1_7) use ($Applicative0_4_18, $f_6) {
  $__num = \func_num_args();
  $__t19 = null;;
  if ($v1_7 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t19 = (($Applicative0_4_18)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_19;;
};
  if ($v1_7 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t19 = ($f_6)(($v1_7)->{'value0'});
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_5) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__local_var_6_20 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_6_20 = (object)["map" => function($f_7) use ($__local_var_6_20) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($__local_var_6_20, $f_7) {
  $__num = \func_num_args();
  $__res = ((($__local_var_6_20)->{'map'})(function($v1_9) use ($f_7) {
  $__num = \func_num_args();
  $__t21 = null;;
  if ($v1_9 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t21 = new \Data\Maybe\Data_Maybe_Just(($f_7)(($v1_9)->{'value0'}));
goto end_branch_21;;
};
  $__t21 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_21:;
  $__res = $__t21;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_8);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_7_23 = (($Monad0_1_0)->{'Bind1'})(null);
  $Applicative0_8_24 = (($Monad0_1_0)->{'Applicative0'})(null);
  $Bind1_7_23 = (object)["bind" => function($v_9) use ($Applicative0_8_24, $Bind1_7_23) {
  $__num = \func_num_args();
  $__res = function($f_10) use ($Applicative0_8_24, $Bind1_7_23, $v_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_23)->{'bind'})($v_9))(function($v1_11) use ($Applicative0_8_24, $f_10) {
  $__num = \func_num_args();
  $__t25 = null;;
  if ($v1_11 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t25 = (($Applicative0_8_24)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_25;;
};
  if ($v1_11 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t25 = ($f_10)(($v1_11)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_9) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_8_27 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_8) {
  $__num = \func_num_args();
  $__res = $x_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($Monad0_1_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_8) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__local_var_9_27 = (((((($Monad0_1_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_9_27 = (object)["map" => function($f_10) use ($__local_var_9_27) {
  $__num = \func_num_args();
  $__res = function($v_11) use ($__local_var_9_27, $f_10) {
  $__num = \func_num_args();
  $__res = ((($__local_var_9_27)->{'map'})(function($v1_12) use ($f_10) {
  $__num = \func_num_args();
  $__t28 = null;;
  if ($v1_12 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t28 = new \Data\Maybe\Data_Maybe_Just(($f_10)(($v1_12)->{'value0'}));
goto end_branch_28;;
};
  $__t28 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_28:;
  $__res = $__t28;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_11);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_10_30 = (($Monad0_1_0)->{'Bind1'})(null);
  $Applicative0_11_31 = (($Monad0_1_0)->{'Applicative0'})(null);
  $Bind1_10_30 = (object)["bind" => function($v_12) use ($Applicative0_11_31, $Bind1_10_30) {
  $__num = \func_num_args();
  $__res = function($f_13) use ($Applicative0_11_31, $Bind1_10_30, $v_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_30)->{'bind'})($v_12))(function($v1_14) use ($Applicative0_11_31, $f_13) {
  $__num = \func_num_args();
  $__t32 = null;;
  if ($v1_14 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t32 = (($Applicative0_11_31)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_32;;
};
  if ($v1_14 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t32 = ($f_13)(($v1_14)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_12) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_11_34 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($Monad0_1_0);
  $__res = (object)["apply" => function($f_12) use ($Applicative0_11_34, $Bind1_10_30) {
  $__num = \func_num_args();
  $__res = function($a_13) use ($Applicative0_11_34, $Bind1_10_30, $f_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_30)->{'bind'})($f_12))(function($f_prime__14) use ($Applicative0_11_34, $Bind1_10_30, $a_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_30)->{'bind'})($a_13))(function($a_prime__15) use ($Applicative0_11_34, $f_prime__14) {
  $__num = \func_num_args();
  $__res = (($Applicative0_11_34)->{'pure'})(($f_prime__14)($a_prime__15));
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
}, "Functor0" => function($_dollar___unused_10) use ($functorMaybeT1_9_27) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_9_27;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_9) use ($Applicative0_8_27, $Bind1_7_23) {
  $__num = \func_num_args();
  $__res = function($a_10) use ($Applicative0_8_27, $Bind1_7_23, $f_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_23)->{'bind'})($f_9))(function($f_prime__11) use ($Applicative0_8_27, $Bind1_7_23, $a_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_23)->{'bind'})($a_10))(function($a_prime__12) use ($Applicative0_8_27, $f_prime__11) {
  $__num = \func_num_args();
  $__res = (($Applicative0_8_27)->{'pure'})(($f_prime__11)($a_prime__12));
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
}, "Functor0" => function($_dollar___unused_7) use ($functorMaybeT1_6_20) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_6_20;
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
  $Bind1_3_37 = (($Monad0_1_0)->{'Bind1'})(null);
  $Applicative0_4_38 = (($Monad0_1_0)->{'Applicative0'})(null);
  $__res = (object)["liftST" => (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_3) {
  $__num = \func_num_args();
  $__res = $x_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($a_5) use ($Applicative0_4_38, $Bind1_3_37) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_37)->{'bind'})($a_5))(function($a_prime__6) use ($Applicative0_4_38) {
  $__num = \func_num_args();
  $__res = (($Applicative0_4_38)->{'pure'})(new \Data\Maybe\Data_Maybe_Just($a_prime__6));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(($dictMonadST_0)->{'liftST'}), "Monad0" => function($_dollar___unused_3) use ($monadMaybeT1_2_1) {
  $__num = \func_num_args();
  $__res = $monadMaybeT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Maybe_Trans_monadSTMaybeT'] = __NAMESPACE__ . '\\majControl_majMonad_majMaybe_majTrans_monadmajSmajTmajMaybemajT';

// Control_Monad_Maybe_Trans_monoidMaybeT
function majControl_majMonad_majMaybe_majTrans_monoidmajMaybemajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majMaybe_majTrans_monoidmajMaybemajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $applicativeMaybeT1_1_0 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_1) {
  $__num = \func_num_args();
  $__res = $x_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_1) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_2_0 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_2_0 = (object)["map" => function($f_3) use ($__local_var_2_0) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($__local_var_2_0, $f_3) {
  $__num = \func_num_args();
  $__res = ((($__local_var_2_0)->{'map'})(function($v1_5) use ($f_3) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v1_5 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = new \Data\Maybe\Data_Maybe_Just(($f_3)(($v1_5)->{'value0'}));
goto end_branch_1;;
};
  $__t1 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_3_3 = (($dictMonad_0)->{'Bind1'})(null);
  $Applicative0_4_4 = (($dictMonad_0)->{'Applicative0'})(null);
  $Bind1_3_3 = (object)["bind" => function($v_5) use ($Applicative0_4_4, $Bind1_3_3) {
  $__num = \func_num_args();
  $__res = function($f_6) use ($Applicative0_4_4, $Bind1_3_3, $v_5) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_3)->{'bind'})($v_5))(function($v1_7) use ($Applicative0_4_4, $f_6) {
  $__num = \func_num_args();
  $__t5 = null;;
  if ($v1_7 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t5 = (($Applicative0_4_4)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_5;;
};
  if ($v1_7 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t5 = ($f_6)(($v1_7)->{'value0'});
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_5) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_6_6 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_6_6 = (object)["map" => function($f_7) use ($__local_var_6_6) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($__local_var_6_6, $f_7) {
  $__num = \func_num_args();
  $__res = ((($__local_var_6_6)->{'map'})(function($v1_9) use ($f_7) {
  $__num = \func_num_args();
  $__t7 = null;;
  if ($v1_9 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t7 = new \Data\Maybe\Data_Maybe_Just(($f_7)(($v1_9)->{'value0'}));
goto end_branch_7;;
};
  $__t7 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_7:;
  $__res = $__t7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_8);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_7_9 = (($dictMonad_0)->{'Bind1'})(null);
  $Applicative0_8_10 = (($dictMonad_0)->{'Applicative0'})(null);
  $Bind1_7_9 = (object)["bind" => function($v_9) use ($Applicative0_8_10, $Bind1_7_9) {
  $__num = \func_num_args();
  $__res = function($f_10) use ($Applicative0_8_10, $Bind1_7_9, $v_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_9)->{'bind'})($v_9))(function($v1_11) use ($Applicative0_8_10, $f_10) {
  $__num = \func_num_args();
  $__t11 = null;;
  if ($v1_11 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t11 = (($Applicative0_8_10)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_11;;
};
  if ($v1_11 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t11 = ($f_10)(($v1_11)->{'value0'});
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_9) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_8_13 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_9) use ($Applicative0_8_13, $Bind1_7_9) {
  $__num = \func_num_args();
  $__res = function($a_10) use ($Applicative0_8_13, $Bind1_7_9, $f_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_9)->{'bind'})($f_9))(function($f_prime__11) use ($Applicative0_8_13, $Bind1_7_9, $a_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_9)->{'bind'})($a_10))(function($a_prime__12) use ($Applicative0_8_13, $f_prime__11) {
  $__num = \func_num_args();
  $__res = (($Applicative0_8_13)->{'pure'})(($f_prime__11)($a_prime__12));
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
}, "Functor0" => function($_dollar___unused_7) use ($functorMaybeT1_6_6) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_6_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_4_15 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_5) use ($Applicative0_4_15, $Bind1_3_3) {
  $__num = \func_num_args();
  $__res = function($a_6) use ($Applicative0_4_15, $Bind1_3_3, $f_5) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_3)->{'bind'})($f_5))(function($f_prime__7) use ($Applicative0_4_15, $Bind1_3_3, $a_6) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_3)->{'bind'})($a_6))(function($a_prime__8) use ($Applicative0_4_15, $f_prime__7) {
  $__num = \func_num_args();
  $__res = (($Applicative0_4_15)->{'pure'})(($f_prime__7)($a_prime__8));
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
}, "Functor0" => function($_dollar___unused_3) use ($functorMaybeT1_2_0) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_2_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_2_17 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_2_17 = (object)["map" => function($f_3) use ($__local_var_2_17) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($__local_var_2_17, $f_3) {
  $__num = \func_num_args();
  $__res = ((($__local_var_2_17)->{'map'})(function($v1_5) use ($f_3) {
  $__num = \func_num_args();
  $__t18 = null;;
  if ($v1_5 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t18 = new \Data\Maybe\Data_Maybe_Just(($f_3)(($v1_5)->{'value0'}));
goto end_branch_18;;
};
  $__t18 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_18:;
  $__res = $__t18;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_3_20 = (($dictMonad_0)->{'Bind1'})(null);
  $Applicative0_4_21 = (($dictMonad_0)->{'Applicative0'})(null);
  $Bind1_3_20 = (object)["bind" => function($v_5) use ($Applicative0_4_21, $Bind1_3_20) {
  $__num = \func_num_args();
  $__res = function($f_6) use ($Applicative0_4_21, $Bind1_3_20, $v_5) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_20)->{'bind'})($v_5))(function($v1_7) use ($Applicative0_4_21, $f_6) {
  $__num = \func_num_args();
  $__t22 = null;;
  if ($v1_7 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t22 = (($Applicative0_4_21)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_22;;
};
  if ($v1_7 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t22 = ($f_6)(($v1_7)->{'value0'});
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_5) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_6_23 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_6_23 = (object)["map" => function($f_7) use ($__local_var_6_23) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($__local_var_6_23, $f_7) {
  $__num = \func_num_args();
  $__res = ((($__local_var_6_23)->{'map'})(function($v1_9) use ($f_7) {
  $__num = \func_num_args();
  $__t24 = null;;
  if ($v1_9 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t24 = new \Data\Maybe\Data_Maybe_Just(($f_7)(($v1_9)->{'value0'}));
goto end_branch_24;;
};
  $__t24 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_24:;
  $__res = $__t24;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_8);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_7_26 = (($dictMonad_0)->{'Bind1'})(null);
  $Applicative0_8_27 = (($dictMonad_0)->{'Applicative0'})(null);
  $Bind1_7_26 = (object)["bind" => function($v_9) use ($Applicative0_8_27, $Bind1_7_26) {
  $__num = \func_num_args();
  $__res = function($f_10) use ($Applicative0_8_27, $Bind1_7_26, $v_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_26)->{'bind'})($v_9))(function($v1_11) use ($Applicative0_8_27, $f_10) {
  $__num = \func_num_args();
  $__t28 = null;;
  if ($v1_11 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t28 = (($Applicative0_8_27)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_28;;
};
  if ($v1_11 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t28 = ($f_10)(($v1_11)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_9) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_8_30 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_8) {
  $__num = \func_num_args();
  $__res = $x_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_8) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_9_30 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_9_30 = (object)["map" => function($f_10) use ($__local_var_9_30) {
  $__num = \func_num_args();
  $__res = function($v_11) use ($__local_var_9_30, $f_10) {
  $__num = \func_num_args();
  $__res = ((($__local_var_9_30)->{'map'})(function($v1_12) use ($f_10) {
  $__num = \func_num_args();
  $__t31 = null;;
  if ($v1_12 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t31 = new \Data\Maybe\Data_Maybe_Just(($f_10)(($v1_12)->{'value0'}));
goto end_branch_31;;
};
  $__t31 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_31:;
  $__res = $__t31;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_11);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_10_33 = (($dictMonad_0)->{'Bind1'})(null);
  $Applicative0_11_34 = (($dictMonad_0)->{'Applicative0'})(null);
  $Bind1_10_33 = (object)["bind" => function($v_12) use ($Applicative0_11_34, $Bind1_10_33) {
  $__num = \func_num_args();
  $__res = function($f_13) use ($Applicative0_11_34, $Bind1_10_33, $v_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_33)->{'bind'})($v_12))(function($v1_14) use ($Applicative0_11_34, $f_13) {
  $__num = \func_num_args();
  $__t35 = null;;
  if ($v1_14 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t35 = (($Applicative0_11_34)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_35;;
};
  if ($v1_14 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t35 = ($f_13)(($v1_14)->{'value0'});
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_12) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_11_37 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_12) use ($Applicative0_11_37, $Bind1_10_33) {
  $__num = \func_num_args();
  $__res = function($a_13) use ($Applicative0_11_37, $Bind1_10_33, $f_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_33)->{'bind'})($f_12))(function($f_prime__14) use ($Applicative0_11_37, $Bind1_10_33, $a_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_33)->{'bind'})($a_13))(function($a_prime__15) use ($Applicative0_11_37, $f_prime__14) {
  $__num = \func_num_args();
  $__res = (($Applicative0_11_37)->{'pure'})(($f_prime__14)($a_prime__15));
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
}, "Functor0" => function($_dollar___unused_10) use ($functorMaybeT1_9_30) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_9_30;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_9) use ($Applicative0_8_30, $Bind1_7_26) {
  $__num = \func_num_args();
  $__res = function($a_10) use ($Applicative0_8_30, $Bind1_7_26, $f_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_26)->{'bind'})($f_9))(function($f_prime__11) use ($Applicative0_8_30, $Bind1_7_26, $a_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_26)->{'bind'})($a_10))(function($a_prime__12) use ($Applicative0_8_30, $f_prime__11) {
  $__num = \func_num_args();
  $__res = (($Applicative0_8_30)->{'pure'})(($f_prime__11)($a_prime__12));
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
}, "Functor0" => function($_dollar___unused_7) use ($functorMaybeT1_6_23) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_6_23;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_4_40 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_4) {
  $__num = \func_num_args();
  $__res = $x_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_4) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_5_40 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_5_40 = (object)["map" => function($f_6) use ($__local_var_5_40) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($__local_var_5_40, $f_6) {
  $__num = \func_num_args();
  $__res = ((($__local_var_5_40)->{'map'})(function($v1_8) use ($f_6) {
  $__num = \func_num_args();
  $__t41 = null;;
  if ($v1_8 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t41 = new \Data\Maybe\Data_Maybe_Just(($f_6)(($v1_8)->{'value0'}));
goto end_branch_41;;
};
  $__t41 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_41:;
  $__res = $__t41;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_6_43 = (($dictMonad_0)->{'Bind1'})(null);
  $Applicative0_7_44 = (($dictMonad_0)->{'Applicative0'})(null);
  $Bind1_6_43 = (object)["bind" => function($v_8) use ($Applicative0_7_44, $Bind1_6_43) {
  $__num = \func_num_args();
  $__res = function($f_9) use ($Applicative0_7_44, $Bind1_6_43, $v_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_43)->{'bind'})($v_8))(function($v1_10) use ($Applicative0_7_44, $f_9) {
  $__num = \func_num_args();
  $__t45 = null;;
  if ($v1_10 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t45 = (($Applicative0_7_44)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_45;;
};
  if ($v1_10 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t45 = ($f_9)(($v1_10)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_8) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_9_46 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_9_46 = (object)["map" => function($f_10) use ($__local_var_9_46) {
  $__num = \func_num_args();
  $__res = function($v_11) use ($__local_var_9_46, $f_10) {
  $__num = \func_num_args();
  $__res = ((($__local_var_9_46)->{'map'})(function($v1_12) use ($f_10) {
  $__num = \func_num_args();
  $__t47 = null;;
  if ($v1_12 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t47 = new \Data\Maybe\Data_Maybe_Just(($f_10)(($v1_12)->{'value0'}));
goto end_branch_47;;
};
  $__t47 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_47:;
  $__res = $__t47;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_11);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_10_49 = (($dictMonad_0)->{'Bind1'})(null);
  $Applicative0_11_50 = (($dictMonad_0)->{'Applicative0'})(null);
  $Bind1_10_49 = (object)["bind" => function($v_12) use ($Applicative0_11_50, $Bind1_10_49) {
  $__num = \func_num_args();
  $__res = function($f_13) use ($Applicative0_11_50, $Bind1_10_49, $v_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_49)->{'bind'})($v_12))(function($v1_14) use ($Applicative0_11_50, $f_13) {
  $__num = \func_num_args();
  $__t51 = null;;
  if ($v1_14 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t51 = (($Applicative0_11_50)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_51;;
};
  if ($v1_14 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t51 = ($f_13)(($v1_14)->{'value0'});
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_12) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_11_53 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_11) {
  $__num = \func_num_args();
  $__res = $x_11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_11) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_12_53 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_12_53 = (object)["map" => function($f_13) use ($__local_var_12_53) {
  $__num = \func_num_args();
  $__res = function($v_14) use ($__local_var_12_53, $f_13) {
  $__num = \func_num_args();
  $__res = ((($__local_var_12_53)->{'map'})(function($v1_15) use ($f_13) {
  $__num = \func_num_args();
  $__t54 = null;;
  if ($v1_15 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t54 = new \Data\Maybe\Data_Maybe_Just(($f_13)(($v1_15)->{'value0'}));
goto end_branch_54;;
};
  $__t54 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_54:;
  $__res = $__t54;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_14);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_13_56 = (($dictMonad_0)->{'Bind1'})(null);
  $Applicative0_14_57 = (($dictMonad_0)->{'Applicative0'})(null);
  $Bind1_13_56 = (object)["bind" => function($v_15) use ($Applicative0_14_57, $Bind1_13_56) {
  $__num = \func_num_args();
  $__res = function($f_16) use ($Applicative0_14_57, $Bind1_13_56, $v_15) {
  $__num = \func_num_args();
  $__res = ((($Bind1_13_56)->{'bind'})($v_15))(function($v1_17) use ($Applicative0_14_57, $f_16) {
  $__num = \func_num_args();
  $__t58 = null;;
  if ($v1_17 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t58 = (($Applicative0_14_57)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_58;;
};
  if ($v1_17 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t58 = ($f_16)(($v1_17)->{'value0'});
goto end_branch_58;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t58 = null;
  end_branch_58:;
  $__res = $__t58;
  goto __end;;
  __end:
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
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_14_60 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_15) use ($Applicative0_14_60, $Bind1_13_56) {
  $__num = \func_num_args();
  $__res = function($a_16) use ($Applicative0_14_60, $Bind1_13_56, $f_15) {
  $__num = \func_num_args();
  $__res = ((($Bind1_13_56)->{'bind'})($f_15))(function($f_prime__17) use ($Applicative0_14_60, $Bind1_13_56, $a_16) {
  $__num = \func_num_args();
  $__res = ((($Bind1_13_56)->{'bind'})($a_16))(function($a_prime__18) use ($Applicative0_14_60, $f_prime__17) {
  $__num = \func_num_args();
  $__res = (($Applicative0_14_60)->{'pure'})(($f_prime__17)($a_prime__18));
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
}, "Functor0" => function($_dollar___unused_13) use ($functorMaybeT1_12_53) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_12_53;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_12) use ($Applicative0_11_53, $Bind1_10_49) {
  $__num = \func_num_args();
  $__res = function($a_13) use ($Applicative0_11_53, $Bind1_10_49, $f_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_49)->{'bind'})($f_12))(function($f_prime__14) use ($Applicative0_11_53, $Bind1_10_49, $a_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_10_49)->{'bind'})($a_13))(function($a_prime__15) use ($Applicative0_11_53, $f_prime__14) {
  $__num = \func_num_args();
  $__res = (($Applicative0_11_53)->{'pure'})(($f_prime__14)($a_prime__15));
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
}, "Functor0" => function($_dollar___unused_10) use ($functorMaybeT1_9_46) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_9_46;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_7_63 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_7) {
  $__num = \func_num_args();
  $__res = $x_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_7) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_8_63 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_8_63 = (object)["map" => function($f_9) use ($__local_var_8_63) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($__local_var_8_63, $f_9) {
  $__num = \func_num_args();
  $__res = ((($__local_var_8_63)->{'map'})(function($v1_11) use ($f_9) {
  $__num = \func_num_args();
  $__t64 = null;;
  if ($v1_11 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t64 = new \Data\Maybe\Data_Maybe_Just(($f_9)(($v1_11)->{'value0'}));
goto end_branch_64;;
};
  $__t64 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_64:;
  $__res = $__t64;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_10);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_9_66 = (($dictMonad_0)->{'Bind1'})(null);
  $Applicative0_10_67 = (($dictMonad_0)->{'Applicative0'})(null);
  $Bind1_9_66 = (object)["bind" => function($v_11) use ($Applicative0_10_67, $Bind1_9_66) {
  $__num = \func_num_args();
  $__res = function($f_12) use ($Applicative0_10_67, $Bind1_9_66, $v_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_66)->{'bind'})($v_11))(function($v1_13) use ($Applicative0_10_67, $f_12) {
  $__num = \func_num_args();
  $__t68 = null;;
  if ($v1_13 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t68 = (($Applicative0_10_67)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_68;;
};
  if ($v1_13 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t68 = ($f_12)(($v1_13)->{'value0'});
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
  $functorMaybeT1_12_69 = (object)["map" => function($f_13) use ($__local_var_12_69) {
  $__num = \func_num_args();
  $__res = function($v_14) use ($__local_var_12_69, $f_13) {
  $__num = \func_num_args();
  $__res = ((($__local_var_12_69)->{'map'})(function($v1_15) use ($f_13) {
  $__num = \func_num_args();
  $__t70 = null;;
  if ($v1_15 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t70 = new \Data\Maybe\Data_Maybe_Just(($f_13)(($v1_15)->{'value0'}));
goto end_branch_70;;
};
  $__t70 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_70:;
  $__res = $__t70;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_14);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_13_72 = (($dictMonad_0)->{'Bind1'})(null);
  $Applicative0_14_73 = (($dictMonad_0)->{'Applicative0'})(null);
  $Bind1_13_72 = (object)["bind" => function($v_15) use ($Applicative0_14_73, $Bind1_13_72) {
  $__num = \func_num_args();
  $__res = function($f_16) use ($Applicative0_14_73, $Bind1_13_72, $v_15) {
  $__num = \func_num_args();
  $__res = ((($Bind1_13_72)->{'bind'})($v_15))(function($v1_17) use ($Applicative0_14_73, $f_16) {
  $__num = \func_num_args();
  $__t74 = null;;
  if ($v1_17 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t74 = (($Applicative0_14_73)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_74;;
};
  if ($v1_17 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t74 = ($f_16)(($v1_17)->{'value0'});
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
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_14_76 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_15) use ($Applicative0_14_76, $Bind1_13_72) {
  $__num = \func_num_args();
  $__res = function($a_16) use ($Applicative0_14_76, $Bind1_13_72, $f_15) {
  $__num = \func_num_args();
  $__res = ((($Bind1_13_72)->{'bind'})($f_15))(function($f_prime__17) use ($Applicative0_14_76, $Bind1_13_72, $a_16) {
  $__num = \func_num_args();
  $__res = ((($Bind1_13_72)->{'bind'})($a_16))(function($a_prime__18) use ($Applicative0_14_76, $f_prime__17) {
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
}, "Functor0" => function($_dollar___unused_13) use ($functorMaybeT1_12_69) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_12_69;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_10_78 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_11) use ($Applicative0_10_78, $Bind1_9_66) {
  $__num = \func_num_args();
  $__res = function($a_12) use ($Applicative0_10_78, $Bind1_9_66, $f_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_66)->{'bind'})($f_11))(function($f_prime__13) use ($Applicative0_10_78, $Bind1_9_66, $a_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_66)->{'bind'})($a_12))(function($a_prime__14) use ($Applicative0_10_78, $f_prime__13) {
  $__num = \func_num_args();
  $__res = (($Applicative0_10_78)->{'pure'})(($f_prime__13)($a_prime__14));
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
}, "Functor0" => function($_dollar___unused_9) use ($functorMaybeT1_8_63) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_8_63;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_8) use ($Applicative0_7_63, $Bind1_6_43) {
  $__num = \func_num_args();
  $__res = function($a_9) use ($Applicative0_7_63, $Bind1_6_43, $f_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_43)->{'bind'})($f_8))(function($f_prime__10) use ($Applicative0_7_63, $Bind1_6_43, $a_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_43)->{'bind'})($a_9))(function($a_prime__11) use ($Applicative0_7_63, $f_prime__10) {
  $__num = \func_num_args();
  $__res = (($Applicative0_7_63)->{'pure'})(($f_prime__10)($a_prime__11));
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
}, "Functor0" => function($_dollar___unused_6) use ($functorMaybeT1_5_40) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_5_40;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyMaybeT1_2_17 = (object)["apply" => function($f_5) use ($Applicative0_4_40, $Bind1_3_20) {
  $__num = \func_num_args();
  $__res = function($a_6) use ($Applicative0_4_40, $Bind1_3_20, $f_5) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_20)->{'bind'})($f_5))(function($f_prime__7) use ($Applicative0_4_40, $Bind1_3_20, $a_6) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_20)->{'bind'})($a_6))(function($a_prime__8) use ($Applicative0_4_40, $f_prime__7) {
  $__num = \func_num_args();
  $__res = (($Applicative0_4_40)->{'pure'})(($f_prime__7)($a_prime__8));
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
}, "Functor0" => function($_dollar___unused_3) use ($functorMaybeT1_2_17) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_2_17;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = function($dictMonoid_3) use ($applicativeMaybeT1_1_0, $applyMaybeT1_2_17) {
  $__num = \func_num_args();
  $Functor0_4_82 = (($applyMaybeT1_2_17)->{'Functor0'})(null);
  $__local_var_5_83 = ((($dictMonoid_3)->{'Semigroup0'})(null))->{'append'};
  $semigroupMaybeT2_4_82 = (object)["append" => function($a_6) use ($Functor0_4_82, $__local_var_5_83, $applyMaybeT1_2_17) {
  $__num = \func_num_args();
  $__res = function($b_7) use ($Functor0_4_82, $__local_var_5_83, $a_6, $applyMaybeT1_2_17) {
  $__num = \func_num_args();
  $__res = ((($applyMaybeT1_2_17)->{'apply'})(((($Functor0_4_82)->{'map'})($__local_var_5_83))($a_6)))($b_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["mempty" => (($applicativeMaybeT1_1_0)->{'pure'})(($dictMonoid_3)->{'mempty'}), "Semigroup0" => function($_dollar___unused_5) use ($semigroupMaybeT2_4_82) {
  $__num = \func_num_args();
  $__res = $semigroupMaybeT2_4_82;
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
$GLOBALS['Control_Monad_Maybe_Trans_monoidMaybeT'] = __NAMESPACE__ . '\\majControl_majMonad_majMaybe_majTrans_monoidmajMaybemajT';

// Control_Monad_Maybe_Trans_altMaybeT
function majControl_majMonad_majMaybe_majTrans_altmajMaybemajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majMaybe_majTrans_altmajMaybemajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Bind1_1_0 = (($dictMonad_0)->{'Bind1'})(null);
  $Applicative0_2_1 = (($dictMonad_0)->{'Applicative0'})(null);
  $__local_var_3_2 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_3_2 = (object)["map" => function($f_4) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__res = function($v_5) use ($__local_var_3_2, $f_4) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_2)->{'map'})(function($v1_6) use ($f_4) {
  $__num = \func_num_args();
  $__t3 = null;;
  if ($v1_6 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t3 = new \Data\Maybe\Data_Maybe_Just(($f_4)(($v1_6)->{'value0'}));
goto end_branch_3;;
};
  $__t3 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["alt" => function($v_4) use ($Applicative0_2_1, $Bind1_1_0) {
  $__num = \func_num_args();
  $__res = function($v1_5) use ($Applicative0_2_1, $Bind1_1_0, $v_4) {
  $__num = \func_num_args();
  $__res = ((($Bind1_1_0)->{'bind'})($v_4))(function($m_6) use ($Applicative0_2_1, $v1_5) {
  $__num = \func_num_args();
  $__t5 = null;;
  if ($m_6 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t5 = $v1_5;
goto end_branch_5;;
};
  $__t5 = (($Applicative0_2_1)->{'pure'})($m_6);
  end_branch_5:;
  $__res = $__t5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_4) use ($functorMaybeT1_3_2) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Maybe_Trans_altMaybeT'] = __NAMESPACE__ . '\\majControl_majMonad_majMaybe_majTrans_altmajMaybemajT';

// Control_Monad_Maybe_Trans_plusMaybeT
function majControl_majMonad_majMaybe_majTrans_plusmajMaybemajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majMaybe_majTrans_plusmajMaybemajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Bind1_1_0 = (($dictMonad_0)->{'Bind1'})(null);
  $Applicative0_2_1 = (($dictMonad_0)->{'Applicative0'})(null);
  $__local_var_3_2 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_3_2 = (object)["map" => function($f_4) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__res = function($v_5) use ($__local_var_3_2, $f_4) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_2)->{'map'})(function($v1_6) use ($f_4) {
  $__num = \func_num_args();
  $__t3 = null;;
  if ($v1_6 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t3 = new \Data\Maybe\Data_Maybe_Just(($f_4)(($v1_6)->{'value0'}));
goto end_branch_3;;
};
  $__t3 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $altMaybeT1_1_0 = (object)["alt" => function($v_4) use ($Applicative0_2_1, $Bind1_1_0) {
  $__num = \func_num_args();
  $__res = function($v1_5) use ($Applicative0_2_1, $Bind1_1_0, $v_4) {
  $__num = \func_num_args();
  $__res = ((($Bind1_1_0)->{'bind'})($v_4))(function($m_6) use ($Applicative0_2_1, $v1_5) {
  $__num = \func_num_args();
  $__t5 = null;;
  if ($m_6 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t5 = $v1_5;
goto end_branch_5;;
};
  $__t5 = (($Applicative0_2_1)->{'pure'})($m_6);
  end_branch_5:;
  $__res = $__t5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_4) use ($functorMaybeT1_3_2) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["empty" => (((($dictMonad_0)->{'Applicative0'})(null))->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing()), "Alt0" => function($_dollar___unused_2) use ($altMaybeT1_1_0) {
  $__num = \func_num_args();
  $__res = $altMaybeT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Maybe_Trans_plusMaybeT'] = __NAMESPACE__ . '\\majControl_majMonad_majMaybe_majTrans_plusmajMaybemajT';

// Control_Monad_Maybe_Trans_alternativeMaybeT
function majControl_majMonad_majMaybe_majTrans_alternativemajMaybemajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majMaybe_majTrans_alternativemajMaybemajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $applicativeMaybeT1_1_0 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_1) {
  $__num = \func_num_args();
  $__res = $x_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_1) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_2_0 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_2_0 = (object)["map" => function($f_3) use ($__local_var_2_0) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($__local_var_2_0, $f_3) {
  $__num = \func_num_args();
  $__res = ((($__local_var_2_0)->{'map'})(function($v1_5) use ($f_3) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v1_5 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = new \Data\Maybe\Data_Maybe_Just(($f_3)(($v1_5)->{'value0'}));
goto end_branch_1;;
};
  $__t1 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_3_3 = (($dictMonad_0)->{'Bind1'})(null);
  $Applicative0_4_4 = (($dictMonad_0)->{'Applicative0'})(null);
  $Bind1_3_3 = (object)["bind" => function($v_5) use ($Applicative0_4_4, $Bind1_3_3) {
  $__num = \func_num_args();
  $__res = function($f_6) use ($Applicative0_4_4, $Bind1_3_3, $v_5) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_3)->{'bind'})($v_5))(function($v1_7) use ($Applicative0_4_4, $f_6) {
  $__num = \func_num_args();
  $__t5 = null;;
  if ($v1_7 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t5 = (($Applicative0_4_4)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_5;;
};
  if ($v1_7 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t5 = ($f_6)(($v1_7)->{'value0'});
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_5) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_6_6 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_6_6 = (object)["map" => function($f_7) use ($__local_var_6_6) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($__local_var_6_6, $f_7) {
  $__num = \func_num_args();
  $__res = ((($__local_var_6_6)->{'map'})(function($v1_9) use ($f_7) {
  $__num = \func_num_args();
  $__t7 = null;;
  if ($v1_9 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t7 = new \Data\Maybe\Data_Maybe_Just(($f_7)(($v1_9)->{'value0'}));
goto end_branch_7;;
};
  $__t7 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_7:;
  $__res = $__t7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_8);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_7_9 = (($dictMonad_0)->{'Bind1'})(null);
  $Applicative0_8_10 = (($dictMonad_0)->{'Applicative0'})(null);
  $Bind1_7_9 = (object)["bind" => function($v_9) use ($Applicative0_8_10, $Bind1_7_9) {
  $__num = \func_num_args();
  $__res = function($f_10) use ($Applicative0_8_10, $Bind1_7_9, $v_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_9)->{'bind'})($v_9))(function($v1_11) use ($Applicative0_8_10, $f_10) {
  $__num = \func_num_args();
  $__t11 = null;;
  if ($v1_11 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t11 = (($Applicative0_8_10)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_11;;
};
  if ($v1_11 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t11 = ($f_10)(($v1_11)->{'value0'});
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_9) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_8_13 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_9) use ($Applicative0_8_13, $Bind1_7_9) {
  $__num = \func_num_args();
  $__res = function($a_10) use ($Applicative0_8_13, $Bind1_7_9, $f_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_9)->{'bind'})($f_9))(function($f_prime__11) use ($Applicative0_8_13, $Bind1_7_9, $a_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_9)->{'bind'})($a_10))(function($a_prime__12) use ($Applicative0_8_13, $f_prime__11) {
  $__num = \func_num_args();
  $__res = (($Applicative0_8_13)->{'pure'})(($f_prime__11)($a_prime__12));
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
}, "Functor0" => function($_dollar___unused_7) use ($functorMaybeT1_6_6) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_6_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_4_15 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_5) use ($Applicative0_4_15, $Bind1_3_3) {
  $__num = \func_num_args();
  $__res = function($a_6) use ($Applicative0_4_15, $Bind1_3_3, $f_5) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_3)->{'bind'})($f_5))(function($f_prime__7) use ($Applicative0_4_15, $Bind1_3_3, $a_6) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_3)->{'bind'})($a_6))(function($a_prime__8) use ($Applicative0_4_15, $f_prime__7) {
  $__num = \func_num_args();
  $__res = (($Applicative0_4_15)->{'pure'})(($f_prime__7)($a_prime__8));
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
}, "Functor0" => function($_dollar___unused_3) use ($functorMaybeT1_2_0) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_2_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_2_17 = (($dictMonad_0)->{'Bind1'})(null);
  $Applicative0_3_18 = (($dictMonad_0)->{'Applicative0'})(null);
  $__local_var_4_19 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_4_19 = (object)["map" => function($f_5) use ($__local_var_4_19) {
  $__num = \func_num_args();
  $__res = function($v_6) use ($__local_var_4_19, $f_5) {
  $__num = \func_num_args();
  $__res = ((($__local_var_4_19)->{'map'})(function($v1_7) use ($f_5) {
  $__num = \func_num_args();
  $__t20 = null;;
  if ($v1_7 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t20 = new \Data\Maybe\Data_Maybe_Just(($f_5)(($v1_7)->{'value0'}));
goto end_branch_20;;
};
  $__t20 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_20:;
  $__res = $__t20;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $altMaybeT1_2_17 = (object)["alt" => function($v_5) use ($Applicative0_3_18, $Bind1_2_17) {
  $__num = \func_num_args();
  $__res = function($v1_6) use ($Applicative0_3_18, $Bind1_2_17, $v_5) {
  $__num = \func_num_args();
  $__res = ((($Bind1_2_17)->{'bind'})($v_5))(function($m_7) use ($Applicative0_3_18, $v1_6) {
  $__num = \func_num_args();
  $__t22 = null;;
  if ($m_7 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t22 = $v1_6;
goto end_branch_22;;
};
  $__t22 = (($Applicative0_3_18)->{'pure'})($m_7);
  end_branch_22:;
  $__res = $__t22;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_5) use ($functorMaybeT1_4_19) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_4_19;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $plusMaybeT1_2_17 = (object)["empty" => (((($dictMonad_0)->{'Applicative0'})(null))->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing()), "Alt0" => function($_dollar___unused_3) use ($altMaybeT1_2_17) {
  $__num = \func_num_args();
  $__res = $altMaybeT1_2_17;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["Applicative0" => function($_dollar___unused_3) use ($applicativeMaybeT1_1_0) {
  $__num = \func_num_args();
  $__res = $applicativeMaybeT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Plus1" => function($_dollar___unused_3) use ($plusMaybeT1_2_17) {
  $__num = \func_num_args();
  $__res = $plusMaybeT1_2_17;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Maybe_Trans_alternativeMaybeT'] = __NAMESPACE__ . '\\majControl_majMonad_majMaybe_majTrans_alternativemajMaybemajT';

// Control_Monad_Maybe_Trans_monadPlusMaybeT
function majControl_majMonad_majMaybe_majTrans_monadmajPlusmajMaybemajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majMaybe_majTrans_monadmajPlusmajMaybemajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $monadMaybeT1_1_0 = (object)["Applicative0" => function($_dollar___unused_1) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_2) {
  $__num = \func_num_args();
  $__res = $x_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_2) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_3_0 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_3_0 = (object)["map" => function($f_4) use ($__local_var_3_0) {
  $__num = \func_num_args();
  $__res = function($v_5) use ($__local_var_3_0, $f_4) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_0)->{'map'})(function($v1_6) use ($f_4) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v1_6 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = new \Data\Maybe\Data_Maybe_Just(($f_4)(($v1_6)->{'value0'}));
goto end_branch_1;;
};
  $__t1 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_4_3 = (($dictMonad_0)->{'Bind1'})(null);
  $Applicative0_5_4 = (($dictMonad_0)->{'Applicative0'})(null);
  $Bind1_4_3 = (object)["bind" => function($v_6) use ($Applicative0_5_4, $Bind1_4_3) {
  $__num = \func_num_args();
  $__res = function($f_7) use ($Applicative0_5_4, $Bind1_4_3, $v_6) {
  $__num = \func_num_args();
  $__res = ((($Bind1_4_3)->{'bind'})($v_6))(function($v1_8) use ($Applicative0_5_4, $f_7) {
  $__num = \func_num_args();
  $__t5 = null;;
  if ($v1_8 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t5 = (($Applicative0_5_4)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_5;;
};
  if ($v1_8 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t5 = ($f_7)(($v1_8)->{'value0'});
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_6) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_7_6 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_7_6 = (object)["map" => function($f_8) use ($__local_var_7_6) {
  $__num = \func_num_args();
  $__res = function($v_9) use ($__local_var_7_6, $f_8) {
  $__num = \func_num_args();
  $__res = ((($__local_var_7_6)->{'map'})(function($v1_10) use ($f_8) {
  $__num = \func_num_args();
  $__t7 = null;;
  if ($v1_10 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t7 = new \Data\Maybe\Data_Maybe_Just(($f_8)(($v1_10)->{'value0'}));
goto end_branch_7;;
};
  $__t7 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_7:;
  $__res = $__t7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_8_9 = (($dictMonad_0)->{'Bind1'})(null);
  $Applicative0_9_10 = (($dictMonad_0)->{'Applicative0'})(null);
  $Bind1_8_9 = (object)["bind" => function($v_10) use ($Applicative0_9_10, $Bind1_8_9) {
  $__num = \func_num_args();
  $__res = function($f_11) use ($Applicative0_9_10, $Bind1_8_9, $v_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_8_9)->{'bind'})($v_10))(function($v1_12) use ($Applicative0_9_10, $f_11) {
  $__num = \func_num_args();
  $__t11 = null;;
  if ($v1_12 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t11 = (($Applicative0_9_10)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_11;;
};
  if ($v1_12 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t11 = ($f_11)(($v1_12)->{'value0'});
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_10) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_9_13 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_10) use ($Applicative0_9_13, $Bind1_8_9) {
  $__num = \func_num_args();
  $__res = function($a_11) use ($Applicative0_9_13, $Bind1_8_9, $f_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_8_9)->{'bind'})($f_10))(function($f_prime__12) use ($Applicative0_9_13, $Bind1_8_9, $a_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_8_9)->{'bind'})($a_11))(function($a_prime__13) use ($Applicative0_9_13, $f_prime__12) {
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
}, "Functor0" => function($_dollar___unused_8) use ($functorMaybeT1_7_6) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_7_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_5_15 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_6) use ($Applicative0_5_15, $Bind1_4_3) {
  $__num = \func_num_args();
  $__res = function($a_7) use ($Applicative0_5_15, $Bind1_4_3, $f_6) {
  $__num = \func_num_args();
  $__res = ((($Bind1_4_3)->{'bind'})($f_6))(function($f_prime__8) use ($Applicative0_5_15, $Bind1_4_3, $a_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_4_3)->{'bind'})($a_7))(function($a_prime__9) use ($Applicative0_5_15, $f_prime__8) {
  $__num = \func_num_args();
  $__res = (($Applicative0_5_15)->{'pure'})(($f_prime__8)($a_prime__9));
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
}, "Functor0" => function($_dollar___unused_4) use ($functorMaybeT1_3_0) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_3_0;
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
  $Bind1_2_16 = (($dictMonad_0)->{'Bind1'})(null);
  $Applicative0_3_17 = (($dictMonad_0)->{'Applicative0'})(null);
  $__res = (object)["bind" => function($v_4) use ($Applicative0_3_17, $Bind1_2_16) {
  $__num = \func_num_args();
  $__res = function($f_5) use ($Applicative0_3_17, $Bind1_2_16, $v_4) {
  $__num = \func_num_args();
  $__res = ((($Bind1_2_16)->{'bind'})($v_4))(function($v1_6) use ($Applicative0_3_17, $f_5) {
  $__num = \func_num_args();
  $__t18 = null;;
  if ($v1_6 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t18 = (($Applicative0_3_17)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_18;;
};
  if ($v1_6 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t18 = ($f_5)(($v1_6)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_4) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_5_19 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_5_19 = (object)["map" => function($f_6) use ($__local_var_5_19) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($__local_var_5_19, $f_6) {
  $__num = \func_num_args();
  $__res = ((($__local_var_5_19)->{'map'})(function($v1_8) use ($f_6) {
  $__num = \func_num_args();
  $__t20 = null;;
  if ($v1_8 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t20 = new \Data\Maybe\Data_Maybe_Just(($f_6)(($v1_8)->{'value0'}));
goto end_branch_20;;
};
  $__t20 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_20:;
  $__res = $__t20;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_6_22 = (($dictMonad_0)->{'Bind1'})(null);
  $Applicative0_7_23 = (($dictMonad_0)->{'Applicative0'})(null);
  $Bind1_6_22 = (object)["bind" => function($v_8) use ($Applicative0_7_23, $Bind1_6_22) {
  $__num = \func_num_args();
  $__res = function($f_9) use ($Applicative0_7_23, $Bind1_6_22, $v_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_22)->{'bind'})($v_8))(function($v1_10) use ($Applicative0_7_23, $f_9) {
  $__num = \func_num_args();
  $__t24 = null;;
  if ($v1_10 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t24 = (($Applicative0_7_23)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_24;;
};
  if ($v1_10 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t24 = ($f_9)(($v1_10)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_8) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_7_26 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_7) {
  $__num = \func_num_args();
  $__res = $x_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_7) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_8_26 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_8_26 = (object)["map" => function($f_9) use ($__local_var_8_26) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($__local_var_8_26, $f_9) {
  $__num = \func_num_args();
  $__res = ((($__local_var_8_26)->{'map'})(function($v1_11) use ($f_9) {
  $__num = \func_num_args();
  $__t27 = null;;
  if ($v1_11 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t27 = new \Data\Maybe\Data_Maybe_Just(($f_9)(($v1_11)->{'value0'}));
goto end_branch_27;;
};
  $__t27 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_27:;
  $__res = $__t27;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_10);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_9_29 = (($dictMonad_0)->{'Bind1'})(null);
  $Applicative0_10_30 = (($dictMonad_0)->{'Applicative0'})(null);
  $Bind1_9_29 = (object)["bind" => function($v_11) use ($Applicative0_10_30, $Bind1_9_29) {
  $__num = \func_num_args();
  $__res = function($f_12) use ($Applicative0_10_30, $Bind1_9_29, $v_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_29)->{'bind'})($v_11))(function($v1_13) use ($Applicative0_10_30, $f_12) {
  $__num = \func_num_args();
  $__t31 = null;;
  if ($v1_13 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t31 = (($Applicative0_10_30)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_31;;
};
  if ($v1_13 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t31 = ($f_12)(($v1_13)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_11) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_10_33 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_11) use ($Applicative0_10_33, $Bind1_9_29) {
  $__num = \func_num_args();
  $__res = function($a_12) use ($Applicative0_10_33, $Bind1_9_29, $f_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_29)->{'bind'})($f_11))(function($f_prime__13) use ($Applicative0_10_33, $Bind1_9_29, $a_12) {
  $__num = \func_num_args();
  $__res = ((($Bind1_9_29)->{'bind'})($a_12))(function($a_prime__14) use ($Applicative0_10_33, $f_prime__13) {
  $__num = \func_num_args();
  $__res = (($Applicative0_10_33)->{'pure'})(($f_prime__13)($a_prime__14));
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
}, "Functor0" => function($_dollar___unused_9) use ($functorMaybeT1_8_26) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_8_26;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_8) use ($Applicative0_7_26, $Bind1_6_22) {
  $__num = \func_num_args();
  $__res = function($a_9) use ($Applicative0_7_26, $Bind1_6_22, $f_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_22)->{'bind'})($f_8))(function($f_prime__10) use ($Applicative0_7_26, $Bind1_6_22, $a_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_22)->{'bind'})($a_9))(function($a_prime__11) use ($Applicative0_7_26, $f_prime__10) {
  $__num = \func_num_args();
  $__res = (($Applicative0_7_26)->{'pure'})(($f_prime__10)($a_prime__11));
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
}, "Functor0" => function($_dollar___unused_6) use ($functorMaybeT1_5_19) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_5_19;
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
  $applicativeMaybeT1_2_36 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_2) {
  $__num = \func_num_args();
  $__res = $x_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_2) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_3_36 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_3_36 = (object)["map" => function($f_4) use ($__local_var_3_36) {
  $__num = \func_num_args();
  $__res = function($v_5) use ($__local_var_3_36, $f_4) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_36)->{'map'})(function($v1_6) use ($f_4) {
  $__num = \func_num_args();
  $__t37 = null;;
  if ($v1_6 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t37 = new \Data\Maybe\Data_Maybe_Just(($f_4)(($v1_6)->{'value0'}));
goto end_branch_37;;
};
  $__t37 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_37:;
  $__res = $__t37;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_4_39 = (($dictMonad_0)->{'Bind1'})(null);
  $Applicative0_5_40 = (($dictMonad_0)->{'Applicative0'})(null);
  $Bind1_4_39 = (object)["bind" => function($v_6) use ($Applicative0_5_40, $Bind1_4_39) {
  $__num = \func_num_args();
  $__res = function($f_7) use ($Applicative0_5_40, $Bind1_4_39, $v_6) {
  $__num = \func_num_args();
  $__res = ((($Bind1_4_39)->{'bind'})($v_6))(function($v1_8) use ($Applicative0_5_40, $f_7) {
  $__num = \func_num_args();
  $__t41 = null;;
  if ($v1_8 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t41 = (($Applicative0_5_40)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_41;;
};
  if ($v1_8 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t41 = ($f_7)(($v1_8)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_6) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_7_42 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_7_42 = (object)["map" => function($f_8) use ($__local_var_7_42) {
  $__num = \func_num_args();
  $__res = function($v_9) use ($__local_var_7_42, $f_8) {
  $__num = \func_num_args();
  $__res = ((($__local_var_7_42)->{'map'})(function($v1_10) use ($f_8) {
  $__num = \func_num_args();
  $__t43 = null;;
  if ($v1_10 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t43 = new \Data\Maybe\Data_Maybe_Just(($f_8)(($v1_10)->{'value0'}));
goto end_branch_43;;
};
  $__t43 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_43:;
  $__res = $__t43;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_8_45 = (($dictMonad_0)->{'Bind1'})(null);
  $Applicative0_9_46 = (($dictMonad_0)->{'Applicative0'})(null);
  $Bind1_8_45 = (object)["bind" => function($v_10) use ($Applicative0_9_46, $Bind1_8_45) {
  $__num = \func_num_args();
  $__res = function($f_11) use ($Applicative0_9_46, $Bind1_8_45, $v_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_8_45)->{'bind'})($v_10))(function($v1_12) use ($Applicative0_9_46, $f_11) {
  $__num = \func_num_args();
  $__t47 = null;;
  if ($v1_12 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t47 = (($Applicative0_9_46)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_47;;
};
  if ($v1_12 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t47 = ($f_11)(($v1_12)->{'value0'});
goto end_branch_47;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t47 = null;
  end_branch_47:;
  $__res = $__t47;
  goto __end;;
  __end:
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
  $__local_var_11_48 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_11_48 = (object)["map" => function($f_12) use ($__local_var_11_48) {
  $__num = \func_num_args();
  $__res = function($v_13) use ($__local_var_11_48, $f_12) {
  $__num = \func_num_args();
  $__res = ((($__local_var_11_48)->{'map'})(function($v1_14) use ($f_12) {
  $__num = \func_num_args();
  $__t49 = null;;
  if ($v1_14 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t49 = new \Data\Maybe\Data_Maybe_Just(($f_12)(($v1_14)->{'value0'}));
goto end_branch_49;;
};
  $__t49 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_49:;
  $__res = $__t49;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_13);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_12_51 = (($dictMonad_0)->{'Bind1'})(null);
  $Applicative0_13_52 = (($dictMonad_0)->{'Applicative0'})(null);
  $Bind1_12_51 = (object)["bind" => function($v_14) use ($Applicative0_13_52, $Bind1_12_51) {
  $__num = \func_num_args();
  $__res = function($f_15) use ($Applicative0_13_52, $Bind1_12_51, $v_14) {
  $__num = \func_num_args();
  $__res = ((($Bind1_12_51)->{'bind'})($v_14))(function($v1_16) use ($Applicative0_13_52, $f_15) {
  $__num = \func_num_args();
  $__t53 = null;;
  if ($v1_16 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t53 = (($Applicative0_13_52)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_53;;
};
  if ($v1_16 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t53 = ($f_15)(($v1_16)->{'value0'});
goto end_branch_53;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t53 = null;
  end_branch_53:;
  $__res = $__t53;
  goto __end;;
  __end:
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
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_13_55 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_13) {
  $__num = \func_num_args();
  $__res = $x_13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_13) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_14_55 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_14_55 = (object)["map" => function($f_15) use ($__local_var_14_55) {
  $__num = \func_num_args();
  $__res = function($v_16) use ($__local_var_14_55, $f_15) {
  $__num = \func_num_args();
  $__res = ((($__local_var_14_55)->{'map'})(function($v1_17) use ($f_15) {
  $__num = \func_num_args();
  $__t56 = null;;
  if ($v1_17 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t56 = new \Data\Maybe\Data_Maybe_Just(($f_15)(($v1_17)->{'value0'}));
goto end_branch_56;;
};
  $__t56 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_56:;
  $__res = $__t56;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_16);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_15_58 = (($dictMonad_0)->{'Bind1'})(null);
  $Applicative0_16_59 = (($dictMonad_0)->{'Applicative0'})(null);
  $Bind1_15_58 = (object)["bind" => function($v_17) use ($Applicative0_16_59, $Bind1_15_58) {
  $__num = \func_num_args();
  $__res = function($f_18) use ($Applicative0_16_59, $Bind1_15_58, $v_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_58)->{'bind'})($v_17))(function($v1_19) use ($Applicative0_16_59, $f_18) {
  $__num = \func_num_args();
  $__t60 = null;;
  if ($v1_19 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t60 = (($Applicative0_16_59)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_60;;
};
  if ($v1_19 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t60 = ($f_18)(($v1_19)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_17) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_16_62 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_17) use ($Applicative0_16_62, $Bind1_15_58) {
  $__num = \func_num_args();
  $__res = function($a_18) use ($Applicative0_16_62, $Bind1_15_58, $f_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_58)->{'bind'})($f_17))(function($f_prime__19) use ($Applicative0_16_62, $Bind1_15_58, $a_18) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_58)->{'bind'})($a_18))(function($a_prime__20) use ($Applicative0_16_62, $f_prime__19) {
  $__num = \func_num_args();
  $__res = (($Applicative0_16_62)->{'pure'})(($f_prime__19)($a_prime__20));
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
}, "Functor0" => function($_dollar___unused_15) use ($functorMaybeT1_14_55) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_14_55;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_14) use ($Applicative0_13_55, $Bind1_12_51) {
  $__num = \func_num_args();
  $__res = function($a_15) use ($Applicative0_13_55, $Bind1_12_51, $f_14) {
  $__num = \func_num_args();
  $__res = ((($Bind1_12_51)->{'bind'})($f_14))(function($f_prime__16) use ($Applicative0_13_55, $Bind1_12_51, $a_15) {
  $__num = \func_num_args();
  $__res = ((($Bind1_12_51)->{'bind'})($a_15))(function($a_prime__17) use ($Applicative0_13_55, $f_prime__16) {
  $__num = \func_num_args();
  $__res = (($Applicative0_13_55)->{'pure'})(($f_prime__16)($a_prime__17));
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
}, "Functor0" => function($_dollar___unused_12) use ($functorMaybeT1_11_48) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_11_48;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_9_65 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_9) {
  $__num = \func_num_args();
  $__res = $x_9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_9) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_10_65 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_10_65 = (object)["map" => function($f_11) use ($__local_var_10_65) {
  $__num = \func_num_args();
  $__res = function($v_12) use ($__local_var_10_65, $f_11) {
  $__num = \func_num_args();
  $__res = ((($__local_var_10_65)->{'map'})(function($v1_13) use ($f_11) {
  $__num = \func_num_args();
  $__t66 = null;;
  if ($v1_13 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t66 = new \Data\Maybe\Data_Maybe_Just(($f_11)(($v1_13)->{'value0'}));
goto end_branch_66;;
};
  $__t66 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_66:;
  $__res = $__t66;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_12);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_11_68 = (($dictMonad_0)->{'Bind1'})(null);
  $Applicative0_12_69 = (($dictMonad_0)->{'Applicative0'})(null);
  $Bind1_11_68 = (object)["bind" => function($v_13) use ($Applicative0_12_69, $Bind1_11_68) {
  $__num = \func_num_args();
  $__res = function($f_14) use ($Applicative0_12_69, $Bind1_11_68, $v_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_68)->{'bind'})($v_13))(function($v1_15) use ($Applicative0_12_69, $f_14) {
  $__num = \func_num_args();
  $__t70 = null;;
  if ($v1_15 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t70 = (($Applicative0_12_69)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_70;;
};
  if ($v1_15 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t70 = ($f_14)(($v1_15)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_13) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_14_71 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_14_71 = (object)["map" => function($f_15) use ($__local_var_14_71) {
  $__num = \func_num_args();
  $__res = function($v_16) use ($__local_var_14_71, $f_15) {
  $__num = \func_num_args();
  $__res = ((($__local_var_14_71)->{'map'})(function($v1_17) use ($f_15) {
  $__num = \func_num_args();
  $__t72 = null;;
  if ($v1_17 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t72 = new \Data\Maybe\Data_Maybe_Just(($f_15)(($v1_17)->{'value0'}));
goto end_branch_72;;
};
  $__t72 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_72:;
  $__res = $__t72;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_16);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_15_74 = (($dictMonad_0)->{'Bind1'})(null);
  $Applicative0_16_75 = (($dictMonad_0)->{'Applicative0'})(null);
  $Bind1_15_74 = (object)["bind" => function($v_17) use ($Applicative0_16_75, $Bind1_15_74) {
  $__num = \func_num_args();
  $__res = function($f_18) use ($Applicative0_16_75, $Bind1_15_74, $v_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_74)->{'bind'})($v_17))(function($v1_19) use ($Applicative0_16_75, $f_18) {
  $__num = \func_num_args();
  $__t76 = null;;
  if ($v1_19 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t76 = (($Applicative0_16_75)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_76;;
};
  if ($v1_19 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t76 = ($f_18)(($v1_19)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_17) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_16_78 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_17) use ($Applicative0_16_78, $Bind1_15_74) {
  $__num = \func_num_args();
  $__res = function($a_18) use ($Applicative0_16_78, $Bind1_15_74, $f_17) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_74)->{'bind'})($f_17))(function($f_prime__19) use ($Applicative0_16_78, $Bind1_15_74, $a_18) {
  $__num = \func_num_args();
  $__res = ((($Bind1_15_74)->{'bind'})($a_18))(function($a_prime__20) use ($Applicative0_16_78, $f_prime__19) {
  $__num = \func_num_args();
  $__res = (($Applicative0_16_78)->{'pure'})(($f_prime__19)($a_prime__20));
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
}, "Functor0" => function($_dollar___unused_15) use ($functorMaybeT1_14_71) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_14_71;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_12_80 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_13) use ($Applicative0_12_80, $Bind1_11_68) {
  $__num = \func_num_args();
  $__res = function($a_14) use ($Applicative0_12_80, $Bind1_11_68, $f_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_68)->{'bind'})($f_13))(function($f_prime__15) use ($Applicative0_12_80, $Bind1_11_68, $a_14) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_68)->{'bind'})($a_14))(function($a_prime__16) use ($Applicative0_12_80, $f_prime__15) {
  $__num = \func_num_args();
  $__res = (($Applicative0_12_80)->{'pure'})(($f_prime__15)($a_prime__16));
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
}, "Functor0" => function($_dollar___unused_11) use ($functorMaybeT1_10_65) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_10_65;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_10) use ($Applicative0_9_65, $Bind1_8_45) {
  $__num = \func_num_args();
  $__res = function($a_11) use ($Applicative0_9_65, $Bind1_8_45, $f_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_8_45)->{'bind'})($f_10))(function($f_prime__12) use ($Applicative0_9_65, $Bind1_8_45, $a_11) {
  $__num = \func_num_args();
  $__res = ((($Bind1_8_45)->{'bind'})($a_11))(function($a_prime__13) use ($Applicative0_9_65, $f_prime__12) {
  $__num = \func_num_args();
  $__res = (($Applicative0_9_65)->{'pure'})(($f_prime__12)($a_prime__13));
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
}, "Functor0" => function($_dollar___unused_8) use ($functorMaybeT1_7_42) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_7_42;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_5_83 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_5) {
  $__num = \func_num_args();
  $__res = $x_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_0)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_5) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_6_83 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_6_83 = (object)["map" => function($f_7) use ($__local_var_6_83) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($__local_var_6_83, $f_7) {
  $__num = \func_num_args();
  $__res = ((($__local_var_6_83)->{'map'})(function($v1_9) use ($f_7) {
  $__num = \func_num_args();
  $__t84 = null;;
  if ($v1_9 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t84 = new \Data\Maybe\Data_Maybe_Just(($f_7)(($v1_9)->{'value0'}));
goto end_branch_84;;
};
  $__t84 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_84:;
  $__res = $__t84;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_8);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_7_86 = (($dictMonad_0)->{'Bind1'})(null);
  $Applicative0_8_87 = (($dictMonad_0)->{'Applicative0'})(null);
  $Bind1_7_86 = (object)["bind" => function($v_9) use ($Applicative0_8_87, $Bind1_7_86) {
  $__num = \func_num_args();
  $__res = function($f_10) use ($Applicative0_8_87, $Bind1_7_86, $v_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_86)->{'bind'})($v_9))(function($v1_11) use ($Applicative0_8_87, $f_10) {
  $__num = \func_num_args();
  $__t88 = null;;
  if ($v1_11 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t88 = (($Applicative0_8_87)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_88;;
};
  if ($v1_11 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t88 = ($f_10)(($v1_11)->{'value0'});
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_9) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__local_var_10_89 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_10_89 = (object)["map" => function($f_11) use ($__local_var_10_89) {
  $__num = \func_num_args();
  $__res = function($v_12) use ($__local_var_10_89, $f_11) {
  $__num = \func_num_args();
  $__res = ((($__local_var_10_89)->{'map'})(function($v1_13) use ($f_11) {
  $__num = \func_num_args();
  $__t90 = null;;
  if ($v1_13 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t90 = new \Data\Maybe\Data_Maybe_Just(($f_11)(($v1_13)->{'value0'}));
goto end_branch_90;;
};
  $__t90 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_90:;
  $__res = $__t90;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_12);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_11_92 = (($dictMonad_0)->{'Bind1'})(null);
  $Applicative0_12_93 = (($dictMonad_0)->{'Applicative0'})(null);
  $Bind1_11_92 = (object)["bind" => function($v_13) use ($Applicative0_12_93, $Bind1_11_92) {
  $__num = \func_num_args();
  $__res = function($f_14) use ($Applicative0_12_93, $Bind1_11_92, $v_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_92)->{'bind'})($v_13))(function($v1_15) use ($Applicative0_12_93, $f_14) {
  $__num = \func_num_args();
  $__t94 = null;;
  if ($v1_15 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t94 = (($Applicative0_12_93)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_94;;
};
  if ($v1_15 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t94 = ($f_14)(($v1_15)->{'value0'});
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
}, "Apply0" => function($_dollar___unused_13) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_12_96 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_13) use ($Applicative0_12_96, $Bind1_11_92) {
  $__num = \func_num_args();
  $__res = function($a_14) use ($Applicative0_12_96, $Bind1_11_92, $f_13) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_92)->{'bind'})($f_13))(function($f_prime__15) use ($Applicative0_12_96, $Bind1_11_92, $a_14) {
  $__num = \func_num_args();
  $__res = ((($Bind1_11_92)->{'bind'})($a_14))(function($a_prime__16) use ($Applicative0_12_96, $f_prime__15) {
  $__num = \func_num_args();
  $__res = (($Applicative0_12_96)->{'pure'})(($f_prime__15)($a_prime__16));
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
}, "Functor0" => function($_dollar___unused_11) use ($functorMaybeT1_10_89) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_10_89;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_8_98 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($dictMonad_0);
  $__res = (object)["apply" => function($f_9) use ($Applicative0_8_98, $Bind1_7_86) {
  $__num = \func_num_args();
  $__res = function($a_10) use ($Applicative0_8_98, $Bind1_7_86, $f_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_86)->{'bind'})($f_9))(function($f_prime__11) use ($Applicative0_8_98, $Bind1_7_86, $a_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_86)->{'bind'})($a_10))(function($a_prime__12) use ($Applicative0_8_98, $f_prime__11) {
  $__num = \func_num_args();
  $__res = (($Applicative0_8_98)->{'pure'})(($f_prime__11)($a_prime__12));
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
}, "Functor0" => function($_dollar___unused_7) use ($functorMaybeT1_6_83) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_6_83;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($f_6) use ($Applicative0_5_83, $Bind1_4_39) {
  $__num = \func_num_args();
  $__res = function($a_7) use ($Applicative0_5_83, $Bind1_4_39, $f_6) {
  $__num = \func_num_args();
  $__res = ((($Bind1_4_39)->{'bind'})($f_6))(function($f_prime__8) use ($Applicative0_5_83, $Bind1_4_39, $a_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_4_39)->{'bind'})($a_7))(function($a_prime__9) use ($Applicative0_5_83, $f_prime__8) {
  $__num = \func_num_args();
  $__res = (($Applicative0_5_83)->{'pure'})(($f_prime__8)($a_prime__9));
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
}, "Functor0" => function($_dollar___unused_4) use ($functorMaybeT1_3_36) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_3_36;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_3_101 = (($dictMonad_0)->{'Bind1'})(null);
  $Applicative0_4_102 = (($dictMonad_0)->{'Applicative0'})(null);
  $__local_var_5_103 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_5_103 = (object)["map" => function($f_6) use ($__local_var_5_103) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($__local_var_5_103, $f_6) {
  $__num = \func_num_args();
  $__res = ((($__local_var_5_103)->{'map'})(function($v1_8) use ($f_6) {
  $__num = \func_num_args();
  $__t104 = null;;
  if ($v1_8 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t104 = new \Data\Maybe\Data_Maybe_Just(($f_6)(($v1_8)->{'value0'}));
goto end_branch_104;;
};
  $__t104 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_104:;
  $__res = $__t104;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $altMaybeT1_3_101 = (object)["alt" => function($v_6) use ($Applicative0_4_102, $Bind1_3_101) {
  $__num = \func_num_args();
  $__res = function($v1_7) use ($Applicative0_4_102, $Bind1_3_101, $v_6) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_101)->{'bind'})($v_6))(function($m_8) use ($Applicative0_4_102, $v1_7) {
  $__num = \func_num_args();
  $__t106 = null;;
  if ($m_8 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t106 = $v1_7;
goto end_branch_106;;
};
  $__t106 = (($Applicative0_4_102)->{'pure'})($m_8);
  end_branch_106:;
  $__res = $__t106;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_6) use ($functorMaybeT1_5_103) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_5_103;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $plusMaybeT1_3_101 = (object)["empty" => (((($dictMonad_0)->{'Applicative0'})(null))->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing()), "Alt0" => function($_dollar___unused_4) use ($altMaybeT1_3_101) {
  $__num = \func_num_args();
  $__res = $altMaybeT1_3_101;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $alternativeMaybeT1_2_36 = (object)["Applicative0" => function($_dollar___unused_4) use ($applicativeMaybeT1_2_36) {
  $__num = \func_num_args();
  $__res = $applicativeMaybeT1_2_36;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Plus1" => function($_dollar___unused_4) use ($plusMaybeT1_3_101) {
  $__num = \func_num_args();
  $__res = $plusMaybeT1_3_101;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["Monad0" => function($_dollar___unused_3) use ($monadMaybeT1_1_0) {
  $__num = \func_num_args();
  $__res = $monadMaybeT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Alternative1" => function($_dollar___unused_3) use ($alternativeMaybeT1_2_36) {
  $__num = \func_num_args();
  $__res = $alternativeMaybeT1_2_36;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Maybe_Trans_monadPlusMaybeT'] = __NAMESPACE__ . '\\majControl_majMonad_majMaybe_majTrans_monadmajPlusmajMaybemajT';

