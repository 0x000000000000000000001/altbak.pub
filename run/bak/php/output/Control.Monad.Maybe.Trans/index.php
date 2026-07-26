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

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };




// Control_Monad_Maybe_Trans_identity
function majControl_majMonad_majMaybe_majTrans_identity($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majMaybe_majTrans_identity';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Maybe_Trans_identity'] = __NAMESPACE__ . '\\majControl_majMonad_majMaybe_majTrans_identity';

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
$GLOBALS['Control_Monad_Maybe_Trans_newtypeMaybeT'] = ["Coercible0" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Control_Monad_Maybe_Trans_monadTransMaybeT
$GLOBALS['Control_Monad_Maybe_Trans_monadTransMaybeT'] = ["lift" => function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Maybe_Trans_MaybeT'], function($a_1 = null) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = (((($dictMonad_0)['Bind1'])(null))['bind'])($a_1, function($a_prime_2 = null) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = (((($dictMonad_0)['Applicative0'])(null))['pure'])(new \Data\Maybe\Data_Maybe_Just($a_prime_2));
  goto __end;;
  __end:
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
  $__res = ["map" => (function() use ($dictFunctor_0) {
  $__fn = function($f_1 = null, $v_2 = null) use ($dictFunctor_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($dictFunctor_0)['map'])((($GLOBALS['Data_Maybe_functorMaybe'])['map'])($f_1), $v_2);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
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
  $__res = ["Applicative0" => function($_dollar__unused_1 = null) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar__unused_1 = null) use ($dictMonad_0) {
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
  $__res = ["bind" => (function() use ($dictMonad_0) {
  $__fn = function($v_1 = null, $f_2 = null) use ($dictMonad_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($dictMonad_0)['Bind1'])(null))['bind'])($v_1, function($v1_3 = null) use ($dictMonad_0, $f_2) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v1_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t0 = (((($dictMonad_0)['Applicative0'])(null))['pure'])(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_0;;
};
  if ($v1_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t0 = ($f_2)(($v1_3)->{'value0'});
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
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Apply0" => function($_dollar__unused_1 = null) use ($dictMonad_0) {
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
  $__local_var_1_0 = (((((($dictMonad_0)['Bind1'])(null))['Apply0'])(null))['Functor0'])(null);
  $functorMaybeT1_2_1 = ["map" => (function() use ($__local_var_1_0) {
  $__fn = function($f_2 = null, $v_3 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($__local_var_1_0)['map'])((($GLOBALS['Data_Maybe_functorMaybe'])['map'])($f_2), $v_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $__local_var_3_2 = ($GLOBALS['Control_Monad_Maybe_Trans_bindMaybeT'])($dictMonad_0);
  $__res = ["apply" => (function() use ($__local_var_3_2, $dictMonad_0) {
  $__fn = function($f_4 = null, $a_5 = null) use ($__local_var_3_2, $dictMonad_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($__local_var_3_2)['bind'])($f_4, function($f_prime_6 = null) use ($__local_var_3_2, $a_5, $dictMonad_0) {
  $__num = \func_num_args();
  $__res = (($__local_var_3_2)['bind'])($a_5, function($a_prime_7 = null) use ($dictMonad_0, $f_prime_6) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($dictMonad_0))['pure'])(($f_prime_6)($a_prime_7));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_3 = null) use ($functorMaybeT1_2_1) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_2_1;
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
  $__res = ["pure" => ($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Maybe_Trans_MaybeT'], ($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_0)['Applicative0'])(null))['pure'], $GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar__unused_1 = null) use ($dictMonad_0) {
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
$GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'] = __NAMESPACE__ . '\\majControl_majMonad_majMaybe_majTrans_applicativemajMaybemajT';

// Control_Monad_Maybe_Trans_semigroupMaybeT
function majControl_majMonad_majMaybe_majTrans_semigroupmajMaybemajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majMaybe_majTrans_semigroupmajMaybemajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($dictMonad_0);
  $__res = function($dictSemigroup_2 = null) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__local_var_3_1 = ($dictSemigroup_2)['append'];
  $__res = ["append" => (function() use ($__local_var_1_0, $__local_var_3_1) {
  $__fn = function($a_4 = null, $b_5 = null) use ($__local_var_1_0, $__local_var_3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($__local_var_1_0)['apply'])((((($__local_var_1_0)['Functor0'])(null))['map'])($__local_var_3_1, $a_4), $b_5);
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
  $Monad0_1_0 = (($dictMonadAsk_0)['Monad0'])(null);
  $monadMaybeT1_2_1 = ["Applicative0" => function($_dollar__unused_2 = null) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar__unused_2 = null) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_bindMaybeT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ["ask" => (($GLOBALS['Control_Monad_Maybe_Trans_monadTransMaybeT'])['lift'])($Monad0_1_0, ($dictMonadAsk_0)['ask']), "Monad0" => function($_dollar__unused_3 = null) use ($monadMaybeT1_2_1) {
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
$GLOBALS['Control_Monad_Maybe_Trans_monadAskMaybeT'] = __NAMESPACE__ . '\\majControl_majMonad_majMaybe_majTrans_monadmajAskmajMaybemajT';

// Control_Monad_Maybe_Trans_monadReaderMaybeT
function majControl_majMonad_majMaybe_majTrans_monadmajReadermajMaybemajT($dictMonadReader_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majMaybe_majTrans_monadmajReadermajMaybemajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $monadAskMaybeT1_1_0 = ($GLOBALS['Control_Monad_Maybe_Trans_monadAskMaybeT'])((($dictMonadReader_0)['MonadAsk0'])(null));
  $__res = ["local" => function($f_2 = null) use ($dictMonadReader_0) {
  $__num = \func_num_args();
  $__res = (($dictMonadReader_0)['local'])($f_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "MonadAsk0" => function($_dollar__unused_2 = null) use ($monadAskMaybeT1_1_0) {
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
  $__local_var_1_0 = (($dictMonadCont_0)['Monad0'])(null);
  $monadMaybeT1_2_1 = ["Applicative0" => function($_dollar__unused_2 = null) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($__local_var_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar__unused_2 = null) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_bindMaybeT'])($__local_var_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ["callCC" => function($f_3 = null) use ($dictMonadCont_0) {
  $__num = \func_num_args();
  $__res = (($dictMonadCont_0)['callCC'])(function($c_4 = null) use ($f_3) {
  $__num = \func_num_args();
  $__res = ($f_3)(function($a_5 = null) use ($c_4) {
  $__num = \func_num_args();
  $__res = ($c_4)(new \Data\Maybe\Data_Maybe_Just($a_5));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad0" => function($_dollar__unused_3 = null) use ($monadMaybeT1_2_1) {
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
$GLOBALS['Control_Monad_Maybe_Trans_monadContMaybeT'] = __NAMESPACE__ . '\\majControl_majMonad_majMaybe_majTrans_monadmajContmajMaybemajT';

// Control_Monad_Maybe_Trans_monadEffectMaybe
function majControl_majMonad_majMaybe_majTrans_monadmajEffectmajMaybe($dictMonadEffect_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majMaybe_majTrans_monadmajEffectmajMaybe';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Monad0_1_0 = (($dictMonadEffect_0)['Monad0'])(null);
  $monadMaybeT1_2_1 = ["Applicative0" => function($_dollar__unused_2 = null) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar__unused_2 = null) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_bindMaybeT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ["liftEffect" => ($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Control_Monad_Maybe_Trans_monadTransMaybeT'])['lift'])($Monad0_1_0), ($dictMonadEffect_0)['liftEffect']), "Monad0" => function($_dollar__unused_3 = null) use ($monadMaybeT1_2_1) {
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
  $Monad0_1_0 = (($dictMonadRec_0)['Monad0'])(null);
  $monadMaybeT1_2_1 = ["Applicative0" => function($_dollar__unused_2 = null) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar__unused_2 = null) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_bindMaybeT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ["tailRecM" => function($f_3 = null) use ($Monad0_1_0, $dictMonadRec_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Maybe_Trans_MaybeT'], (($dictMonadRec_0)['tailRecM'])(function($a_4 = null) use ($Monad0_1_0, $f_3) {
  $__num = \func_num_args();
  $__res = (((($Monad0_1_0)['Bind1'])(null))['bind'])(($f_3)($a_4), function($m_prime_5 = null) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ($m_prime_5 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t2 = new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Done(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_2;;
};
  if ($m_prime_5 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t3 = null;;
if (($m_prime_5)->{'value0'} instanceof \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Loop) {
$__t3 = new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Loop((($m_prime_5)->{'value0'})->{'value0'});
goto end_branch_3;;
};
if (($m_prime_5)->{'value0'} instanceof \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Done) {
$__t3 = new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Done(new \Data\Maybe\Data_Maybe_Just((($m_prime_5)->{'value0'})->{'value0'}));
goto end_branch_3;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
end_branch_3:;
$__t2 = $__t3;
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = (((($Monad0_1_0)['Applicative0'])(null))['pure'])($__t2);
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
}, "Monad0" => function($_dollar__unused_3 = null) use ($monadMaybeT1_2_1) {
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
$GLOBALS['Control_Monad_Maybe_Trans_monadRecMaybeT'] = __NAMESPACE__ . '\\majControl_majMonad_majMaybe_majTrans_monadmajRecmajMaybemajT';

// Control_Monad_Maybe_Trans_monadStateMaybeT
function majControl_majMonad_majMaybe_majTrans_monadmajStatemajMaybemajT($dictMonadState_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majMaybe_majTrans_monadmajStatemajMaybemajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Monad0_1_0 = (($dictMonadState_0)['Monad0'])(null);
  $lift1_2_1 = (($GLOBALS['Control_Monad_Maybe_Trans_monadTransMaybeT'])['lift'])($Monad0_1_0);
  $monadMaybeT1_3_2 = ["Applicative0" => function($_dollar__unused_3 = null) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar__unused_3 = null) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_bindMaybeT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ["state" => function($f_4 = null) use ($dictMonadState_0, $lift1_2_1) {
  $__num = \func_num_args();
  $__res = ($lift1_2_1)((($dictMonadState_0)['state'])($f_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad0" => function($_dollar__unused_4 = null) use ($monadMaybeT1_3_2) {
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
$GLOBALS['Control_Monad_Maybe_Trans_monadStateMaybeT'] = __NAMESPACE__ . '\\majControl_majMonad_majMaybe_majTrans_monadmajStatemajMaybemajT';

// Control_Monad_Maybe_Trans_monadTellMaybeT
function majControl_majMonad_majMaybe_majTrans_monadmajTellmajMaybemajT($dictMonadTell_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majMaybe_majTrans_monadmajTellmajMaybemajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Monad1_1_0 = (($dictMonadTell_0)['Monad1'])(null);
  $Semigroup0_2_1 = (($dictMonadTell_0)['Semigroup0'])(null);
  $monadMaybeT1_3_2 = ["Applicative0" => function($_dollar__unused_3 = null) use ($Monad1_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($Monad1_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar__unused_3 = null) use ($Monad1_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_bindMaybeT'])($Monad1_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ["tell" => ($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Control_Monad_Maybe_Trans_monadTransMaybeT'])['lift'])($Monad1_1_0), ($dictMonadTell_0)['tell']), "Semigroup0" => function($_dollar__unused_4 = null) use ($Semigroup0_2_1) {
  $__num = \func_num_args();
  $__res = $Semigroup0_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad1" => function($_dollar__unused_4 = null) use ($monadMaybeT1_3_2) {
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
  $MonadTell1_1_0 = (($dictMonadWriter_0)['MonadTell1'])(null);
  $Monad1_2_1 = (($MonadTell1_1_0)['Monad1'])(null);
  $__local_var_3_2 = (($Monad1_2_1)['Bind1'])(null);
  $__local_var_4_3 = (($Monad1_2_1)['Applicative0'])(null);
  $Monoid0_5_4 = (($dictMonadWriter_0)['Monoid0'])(null);
  $monadTellMaybeT1_6_5 = ($GLOBALS['Control_Monad_Maybe_Trans_monadTellMaybeT'])($MonadTell1_1_0);
  $__res = ["listen" => function($v_7 = null) use ($__local_var_3_2, $__local_var_4_3, $dictMonadWriter_0) {
  $__num = \func_num_args();
  $__res = (($__local_var_3_2)['bind'])((($dictMonadWriter_0)['listen'])($v_7), function($v_8 = null) use ($__local_var_4_3) {
  $__num = \func_num_args();
  $__local_var_9_6 = ($v_8)->{'value1'};
  $__res = (($__local_var_4_3)['pure'])((($GLOBALS['Data_Maybe_functorMaybe'])['map'])(function($r_10 = null) use ($__local_var_9_6) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple($r_10, $__local_var_9_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($v_8)->{'value0'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "pass" => function($v_7 = null) use ($__local_var_3_2, $__local_var_4_3, $dictMonadWriter_0) {
  $__num = \func_num_args();
  $__res = (($dictMonadWriter_0)['pass'])((($__local_var_3_2)['bind'])($v_7, function($a_8 = null) use ($__local_var_4_3) {
  $__num = \func_num_args();
  $__t7 = null;;
  if ($a_8 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t7 = new \Data\Tuple\Data_Tuple_Tuple(new \Data\Maybe\Data_Maybe_Nothing(), $GLOBALS['Control_Monad_Maybe_Trans_identity']);
goto end_branch_7;;
};
  if ($a_8 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t7 = new \Data\Tuple\Data_Tuple_Tuple(new \Data\Maybe\Data_Maybe_Just((($a_8)->{'value0'})->{'value0'}), (($a_8)->{'value0'})->{'value1'});
goto end_branch_7;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t7 = null;
  end_branch_7:;
  $__res = (($__local_var_4_3)['pure'])($__t7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monoid0" => function($_dollar__unused_7 = null) use ($Monoid0_5_4) {
  $__num = \func_num_args();
  $__res = $Monoid0_5_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "MonadTell1" => function($_dollar__unused_7 = null) use ($monadTellMaybeT1_6_5) {
  $__num = \func_num_args();
  $__res = $monadTellMaybeT1_6_5;
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
  $Monad0_1_0 = (($dictMonadThrow_0)['Monad0'])(null);
  $lift1_2_1 = (($GLOBALS['Control_Monad_Maybe_Trans_monadTransMaybeT'])['lift'])($Monad0_1_0);
  $monadMaybeT1_3_2 = ["Applicative0" => function($_dollar__unused_3 = null) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar__unused_3 = null) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_bindMaybeT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ["throwError" => function($e_4 = null) use ($dictMonadThrow_0, $lift1_2_1) {
  $__num = \func_num_args();
  $__res = ($lift1_2_1)((($dictMonadThrow_0)['throwError'])($e_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad0" => function($_dollar__unused_4 = null) use ($monadMaybeT1_3_2) {
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
$GLOBALS['Control_Monad_Maybe_Trans_monadThrowMaybeT'] = __NAMESPACE__ . '\\majControl_majMonad_majMaybe_majTrans_monadmajThrowmajMaybemajT';

// Control_Monad_Maybe_Trans_monadErrorMaybeT
function majControl_majMonad_majMaybe_majTrans_monadmajErrormajMaybemajT($dictMonadError_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majMaybe_majTrans_monadmajErrormajMaybemajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $monadThrowMaybeT1_1_0 = ($GLOBALS['Control_Monad_Maybe_Trans_monadThrowMaybeT'])((($dictMonadError_0)['MonadThrow0'])(null));
  $__res = ["catchError" => (function() use ($dictMonadError_0) {
  $__fn = function($v_2 = null, $h_3 = null) use ($dictMonadError_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($dictMonadError_0)['catchError'])($v_2, function($a_4 = null) use ($h_3) {
  $__num = \func_num_args();
  $__res = ($h_3)($a_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "MonadThrow0" => function($_dollar__unused_2 = null) use ($monadThrowMaybeT1_1_0) {
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
  $Monad0_1_0 = (($dictMonadST_0)['Monad0'])(null);
  $monadMaybeT1_2_1 = ["Applicative0" => function($_dollar__unused_2 = null) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar__unused_2 = null) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_bindMaybeT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ["liftST" => ($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Control_Monad_Maybe_Trans_monadTransMaybeT'])['lift'])($Monad0_1_0), ($dictMonadST_0)['liftST']), "Monad0" => function($_dollar__unused_3 = null) use ($monadMaybeT1_2_1) {
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
  $semigroupMaybeT1_1_0 = ($GLOBALS['Control_Monad_Maybe_Trans_semigroupMaybeT'])($dictMonad_0);
  $__res = function($dictMonoid_2 = null) use ($dictMonad_0, $semigroupMaybeT1_1_0) {
  $__num = \func_num_args();
  $semigroupMaybeT2_3_1 = ($semigroupMaybeT1_1_0)((($dictMonoid_2)['Semigroup0'])(null));
  $__res = ["mempty" => ((($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($dictMonad_0))['pure'])(($dictMonoid_2)['mempty']), "Semigroup0" => function($_dollar__unused_4 = null) use ($semigroupMaybeT2_3_1) {
  $__num = \func_num_args();
  $__res = $semigroupMaybeT2_3_1;
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
  $Bind1_1_0 = (($dictMonad_0)['Bind1'])(null);
  $__local_var_2_1 = (((($Bind1_1_0)['Apply0'])(null))['Functor0'])(null);
  $functorMaybeT1_3_2 = ["map" => (function() use ($__local_var_2_1) {
  $__fn = function($f_3 = null, $v_4 = null) use ($__local_var_2_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($__local_var_2_1)['map'])((($GLOBALS['Data_Maybe_functorMaybe'])['map'])($f_3), $v_4);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $__res = ["alt" => (function() use ($Bind1_1_0, $dictMonad_0) {
  $__fn = function($v_4 = null, $v1_5 = null) use ($Bind1_1_0, $dictMonad_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($Bind1_1_0)['bind'])($v_4, function($m_6 = null) use ($dictMonad_0, $v1_5) {
  $__num = \func_num_args();
  $__t3 = null;;
  if ($m_6 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t3 = $v1_5;
goto end_branch_3;;
};
  $__t3 = (((($dictMonad_0)['Applicative0'])(null))['pure'])($m_6);
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_4 = null) use ($functorMaybeT1_3_2) {
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
  $altMaybeT1_1_0 = ($GLOBALS['Control_Monad_Maybe_Trans_altMaybeT'])($dictMonad_0);
  $__res = ["empty" => (((($dictMonad_0)['Applicative0'])(null))['pure'])(new \Data\Maybe\Data_Maybe_Nothing()), "Alt0" => function($_dollar__unused_2 = null) use ($altMaybeT1_1_0) {
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
  $applicativeMaybeT1_1_0 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($dictMonad_0);
  $plusMaybeT1_2_1 = ($GLOBALS['Control_Monad_Maybe_Trans_plusMaybeT'])($dictMonad_0);
  $__res = ["Applicative0" => function($_dollar__unused_3 = null) use ($applicativeMaybeT1_1_0) {
  $__num = \func_num_args();
  $__res = $applicativeMaybeT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Plus1" => function($_dollar__unused_3 = null) use ($plusMaybeT1_2_1) {
  $__num = \func_num_args();
  $__res = $plusMaybeT1_2_1;
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
  $monadMaybeT1_1_0 = ["Applicative0" => function($_dollar__unused_1 = null) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar__unused_1 = null) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_bindMaybeT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $alternativeMaybeT1_2_1 = ($GLOBALS['Control_Monad_Maybe_Trans_alternativeMaybeT'])($dictMonad_0);
  $__res = ["Monad0" => function($_dollar__unused_3 = null) use ($monadMaybeT1_1_0) {
  $__num = \func_num_args();
  $__res = $monadMaybeT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Alternative1" => function($_dollar__unused_3 = null) use ($alternativeMaybeT1_2_1) {
  $__num = \func_num_args();
  $__res = $alternativeMaybeT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Maybe_Trans_monadPlusMaybeT'] = __NAMESPACE__ . '\\majControl_majMonad_majMaybe_majTrans_monadmajPlusmajMaybemajT';

