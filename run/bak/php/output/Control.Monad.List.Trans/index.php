<?php

namespace Control\Monad\List\Trans;

// ALL IMPORTS: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Monad, Control.Monad.List.Trans, Control.Monad.Rec.Class, Control.Monad.ST.Class, Control.Monad.Trans.Class, Control.MonadPlus, Control.Plus, Control.Semigroupoid, Data.Function, Data.Functor, Data.Lazy, Data.Maybe, Data.Monoid, Data.Newtype, Data.Ring, Data.Semigroup, Data.Tuple, Data.Unfoldable, Data.Unfoldable1, Data.Unit, Effect.Class, Prelude, Prim
// TO REQUIRE: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Monad, Control.Monad.List.Trans, Control.Monad.Rec.Class, Control.Monad.ST.Class, Control.Monad.Trans.Class, Control.MonadPlus, Control.Plus, Control.Semigroupoid, Data.Function, Data.Functor, Data.Lazy, Data.Maybe, Data.Monoid, Data.Newtype, Data.Ring, Data.Semigroup, Data.Tuple, Data.Unfoldable, Data.Unfoldable1, Data.Unit, Effect.Class, Prelude
require_once __DIR__ . '/../Control.Alt/index.php';
require_once __DIR__ . '/../Control.Alternative/index.php';
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Category/index.php';
require_once __DIR__ . '/../Control.Monad/index.php';
require_once __DIR__ . '/../Control.Monad.List.Trans/index.php';
require_once __DIR__ . '/../Control.Monad.Rec.Class/index.php';
require_once __DIR__ . '/../Control.Monad.ST.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Trans.Class/index.php';
require_once __DIR__ . '/../Control.MonadPlus/index.php';
require_once __DIR__ . '/../Control.Plus/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Lazy/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Monoid/index.php';
require_once __DIR__ . '/../Data.Newtype/index.php';
require_once __DIR__ . '/../Data.Ring/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Tuple/index.php';
require_once __DIR__ . '/../Data.Unfoldable/index.php';
require_once __DIR__ . '/../Data.Unfoldable1/index.php';
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


final class Control_Monad_List_Trans_Yield { public $tag = 'Yield'; public function __construct(public  $value0, public  $value1) {} }
final class Control_Monad_List_Trans_Skip { public $tag = 'Skip'; public function __construct(public  $value0) {} }
final class Control_Monad_List_Trans_Done { public $tag = 'Done'; public function __construct() {} }

// Control_Monad_List_Trans_identity
function majControl_majMonad_majList_majTrans_identity($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majList_majTrans_identity';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_List_Trans_identity'] = __NAMESPACE__ . '\\majControl_majMonad_majList_majTrans_identity';

// Control_Monad_List_Trans_Yield
$GLOBALS['Control_Monad_List_Trans_Yield'] = (function() {
  $__fn = function($value0, $value1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Control\Monad\List\Trans\Control_Monad_List_Trans_Yield($value0, $value1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Control_Monad_List_Trans_Skip
$GLOBALS['Control_Monad_List_Trans_Skip'] = function($value0) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\List\Trans\Control_Monad_List_Trans_Skip($value0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_List_Trans_Done
$GLOBALS['Control_Monad_List_Trans_Done'] = ($GLOBALS['__phpurs_data0_Done'] ??= new \Control\Monad\List\Trans\Control_Monad_List_Trans_Done());

// Control_Monad_List_Trans_ListT
function majControl_majMonad_majList_majTrans_majListmajT($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majList_majTrans_majListmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_List_Trans_ListT'] = __NAMESPACE__ . '\\majControl_majMonad_majList_majTrans_majListmajT';

// Control_Monad_List_Trans_wrapLazy
function majControl_majMonad_majList_majTrans_wrapmajLazy($dictApplicative_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majList_majTrans_wrapmajLazy';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($dictApplicative_0)->{'pure'})(new \Control\Monad\List\Trans\Control_Monad_List_Trans_Skip($v_1));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Control_Monad_List_Trans_wrapLazy'] = __NAMESPACE__ . '\\majControl_majMonad_majList_majTrans_wrapmajLazy';

// Control_Monad_List_Trans_wrapEffect
function majControl_majMonad_majList_majTrans_wrapmajEffect($dictFunctor_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majList_majTrans_wrapmajEffect';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictFunctor_0)->{'map'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_List_Trans_Skip']))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Lazy_defer']))($GLOBALS['Data_Function_const']))))($v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Control_Monad_List_Trans_wrapEffect'] = __NAMESPACE__ . '\\majControl_majMonad_majList_majTrans_wrapmajEffect';

// Control_Monad_List_Trans_unfold
function majControl_majMonad_majList_majTrans_unfold($dictMonad_0, $f_1 = null, $z_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majList_majTrans_unfold';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_Control_Monad_List_Trans_unfold_dictMonad_0 = $dictMonad_0;
  $__tco_var_Control_Monad_List_Trans_unfold_f_1 = $f_1;
  $__tco_var_Control_Monad_List_Trans_unfold_z_2 = $z_2;
  tco_loop_Control_Monad_List_Trans_unfold:;
  $dictMonad_0 = $__tco_var_Control_Monad_List_Trans_unfold_dictMonad_0;
  $f_1 = $__tco_var_Control_Monad_List_Trans_unfold_f_1;
  $z_2 = $__tco_var_Control_Monad_List_Trans_unfold_z_2;
  $__res = ((((((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})(function($v_3) use ($dictMonad_0, $f_1) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__local_var_4_1 = (($v_3)->{'value0'})->{'value0'};
$__t0 = new \Control\Monad\List\Trans\Control_Monad_List_Trans_Yield((($v_3)->{'value0'})->{'value1'}, \Data\Lazy\majData_majLazy_defer(function($v1_5) use ($__local_var_4_1, $dictMonad_0, $f_1) {
  $__num = \func_num_args();
  $__res = \Control\Monad\List\Trans\majControl_majMonad_majList_majTrans_unfold($dictMonad_0, $f_1, $__local_var_4_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
goto end_branch_0;;
};
  if ($v_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t0 = new \Control\Monad\List\Trans\Control_Monad_List_Trans_Done();
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($f_1)($z_2));
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Control_Monad_List_Trans_unfold'] = __NAMESPACE__ . '\\majControl_majMonad_majList_majTrans_unfold';

// Control_Monad_List_Trans_uncons
function majControl_majMonad_majList_majTrans_uncons($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majList_majTrans_uncons';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__tco_var_Control_Monad_List_Trans_uncons_dictMonad_0 = $dictMonad_0;
  tco_loop_Control_Monad_List_Trans_uncons:;
  $dictMonad_0 = $__tco_var_Control_Monad_List_Trans_uncons_dictMonad_0;
  $__local_var_1_0 = (($dictMonad_0)->{'Applicative0'})(null);
  $__res = function($v_2) use ($__local_var_1_0, $dictMonad_0) {
  $__num = \func_num_args();
  $__res = ((((($dictMonad_0)->{'Bind1'})(null))->{'bind'})($v_2))(function($v1_3) use ($__local_var_1_0, $dictMonad_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v1_3 instanceof \Control\Monad\List\Trans\Control_Monad_List_Trans_Yield) {
$__t1 = (($__local_var_1_0)->{'pure'})(new \Data\Maybe\Data_Maybe_Just(new \Data\Tuple\Data_Tuple_Tuple(($v1_3)->{'value0'}, \Data\Lazy\majData_majLazy_force(($v1_3)->{'value1'}))));
goto end_branch_1;;
};
  if ($v1_3 instanceof \Control\Monad\List\Trans\Control_Monad_List_Trans_Skip) {
$__t1 = \Control\Monad\List\Trans\majControl_majMonad_majList_majTrans_uncons($dictMonad_0, \Data\Lazy\majData_majLazy_force(($v1_3)->{'value0'}));
goto end_branch_1;;
};
  if ($v1_3 instanceof \Control\Monad\List\Trans\Control_Monad_List_Trans_Done) {
$__t1 = (($__local_var_1_0)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_List_Trans_uncons'] = __NAMESPACE__ . '\\majControl_majMonad_majList_majTrans_uncons';

// Control_Monad_List_Trans_tail
function majControl_majMonad_majList_majTrans_tail($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majList_majTrans_tail';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $uncons1_1_0 = ($GLOBALS['Control_Monad_List_Trans_uncons'])($dictMonad_0);
  $__res = function($l_2) use ($dictMonad_0, $uncons1_1_0) {
  $__num = \func_num_args();
  $__res = ((((((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})($GLOBALS['Data_Tuple_snd'])))(($uncons1_1_0)($l_2));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_List_Trans_tail'] = __NAMESPACE__ . '\\majControl_majMonad_majList_majTrans_tail';

// Control_Monad_List_Trans_takeWhile
function majControl_majMonad_majList_majTrans_takemajWhile($dictApplicative_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majList_majTrans_takemajWhile';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__tco_var_Control_Monad_List_Trans_takeWhile_dictApplicative_0 = $dictApplicative_0;
  tco_loop_Control_Monad_List_Trans_takeWhile:;
  $dictApplicative_0 = $__tco_var_Control_Monad_List_Trans_takeWhile_dictApplicative_0;
  $__local_var_1_0 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = (function() use ($__local_var_1_0, $dictApplicative_0) {
  $__fn = function($f_2, $v_3 = null) use ($__local_var_1_0, $dictApplicative_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_1_0)->{'map'})(function($v_4) use ($dictApplicative_0, $f_2) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v_4 instanceof \Control\Monad\List\Trans\Control_Monad_List_Trans_Yield) {
$__t2 = null;;
if (($f_2)(($v_4)->{'value0'})) {
$__t2 = new \Control\Monad\List\Trans\Control_Monad_List_Trans_Yield(($v_4)->{'value0'}, ((($GLOBALS['Data_Lazy_functorLazy'])->{'map'})((($GLOBALS['Control_Monad_List_Trans_takeWhile'])($dictApplicative_0))($f_2)))(($v_4)->{'value1'}));
goto end_branch_2;;
};
$__t2 = new \Control\Monad\List\Trans\Control_Monad_List_Trans_Done();
end_branch_2:;
$__t1 = $__t2;
goto end_branch_1;;
};
  if ($v_4 instanceof \Control\Monad\List\Trans\Control_Monad_List_Trans_Skip) {
$__t1 = new \Control\Monad\List\Trans\Control_Monad_List_Trans_Skip(((($GLOBALS['Data_Lazy_functorLazy'])->{'map'})((($GLOBALS['Control_Monad_List_Trans_takeWhile'])($dictApplicative_0))($f_2)))(($v_4)->{'value0'}));
goto end_branch_1;;
};
  if ($v_4 instanceof \Control\Monad\List\Trans\Control_Monad_List_Trans_Done) {
$__t1 = new \Control\Monad\List\Trans\Control_Monad_List_Trans_Done();
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_List_Trans_takeWhile'] = __NAMESPACE__ . '\\majControl_majMonad_majList_majTrans_takemajWhile';

// Control_Monad_List_Trans_scanl
function majControl_majMonad_majList_majTrans_scanl($dictMonad_0, $f_1 = null, $b_2 = null, $l_3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majList_majTrans_scanl';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = \Control\Monad\List\Trans\majControl_majMonad_majList_majTrans_unfold($dictMonad_0, function($v_4) use ($dictMonad_0, $f_1) {
  $__num = \func_num_args();
  $__local_var_5_0 = ($v_4)->{'value0'};
  $__res = ((((((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})(function($v1_6) use ($__local_var_5_0, $f_1) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v1_6 instanceof \Control\Monad\List\Trans\Control_Monad_List_Trans_Yield) {
$__t1 = new \Data\Maybe\Data_Maybe_Just(new \Data\Tuple\Data_Tuple_Tuple(new \Data\Tuple\Data_Tuple_Tuple((($f_1)($__local_var_5_0))(($v1_6)->{'value0'}), \Data\Lazy\majData_majLazy_force(($v1_6)->{'value1'})), $__local_var_5_0));
goto end_branch_1;;
};
  if ($v1_6 instanceof \Control\Monad\List\Trans\Control_Monad_List_Trans_Skip) {
$__t1 = new \Data\Maybe\Data_Maybe_Just(new \Data\Tuple\Data_Tuple_Tuple(new \Data\Tuple\Data_Tuple_Tuple($__local_var_5_0, \Data\Lazy\majData_majLazy_force(($v1_6)->{'value0'})), $__local_var_5_0));
goto end_branch_1;;
};
  if ($v1_6 instanceof \Control\Monad\List\Trans\Control_Monad_List_Trans_Done) {
$__t1 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_4)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, new \Data\Tuple\Data_Tuple_Tuple($b_2, $l_3));
  goto __end;;
  __end:
  return 4 < $__num ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
}
$GLOBALS['Control_Monad_List_Trans_scanl'] = __NAMESPACE__ . '\\majControl_majMonad_majList_majTrans_scanl';

// Control_Monad_List_Trans_prepend'
function majControl_majMonad_majList_majTrans_prepend__prime__($dictApplicative_0, $h_1 = null, $t_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majList_majTrans_prepend__prime__';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($dictApplicative_0)->{'pure'})(new \Control\Monad\List\Trans\Control_Monad_List_Trans_Yield($h_1, $t_2));
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Control_Monad_List_Trans_prepend__prime__'] = __NAMESPACE__ . '\\majControl_majMonad_majList_majTrans_prepend__prime__';

// Control_Monad_List_Trans_prepend
function majControl_majMonad_majList_majTrans_prepend($dictApplicative_0, $h_1 = null, $t_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majList_majTrans_prepend';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($dictApplicative_0)->{'pure'})(new \Control\Monad\List\Trans\Control_Monad_List_Trans_Yield($h_1, \Data\Lazy\majData_majLazy_defer(function($v_3) use ($t_2) {
  $__num = \func_num_args();
  $__res = $t_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})));
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Control_Monad_List_Trans_prepend'] = __NAMESPACE__ . '\\majControl_majMonad_majList_majTrans_prepend';

// Control_Monad_List_Trans_nil
function majControl_majMonad_majList_majTrans_nil($dictApplicative_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majList_majTrans_nil';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($dictApplicative_0)->{'pure'})(new \Control\Monad\List\Trans\Control_Monad_List_Trans_Done());
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_List_Trans_nil'] = __NAMESPACE__ . '\\majControl_majMonad_majList_majTrans_nil';

// Control_Monad_List_Trans_singleton
function majControl_majMonad_majList_majTrans_singleton($dictApplicative_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majList_majTrans_singleton';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $nil1_1_0 = (($dictApplicative_0)->{'pure'})(new \Control\Monad\List\Trans\Control_Monad_List_Trans_Done());
  $__res = function($a_2) use ($dictApplicative_0, $nil1_1_0) {
  $__num = \func_num_args();
  $__res = (($dictApplicative_0)->{'pure'})(new \Control\Monad\List\Trans\Control_Monad_List_Trans_Yield($a_2, \Data\Lazy\majData_majLazy_defer(function($v_3) use ($nil1_1_0) {
  $__num = \func_num_args();
  $__res = $nil1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_List_Trans_singleton'] = __NAMESPACE__ . '\\majControl_majMonad_majList_majTrans_singleton';

// Control_Monad_List_Trans_take
function majControl_majMonad_majList_majTrans_take($dictApplicative_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majList_majTrans_take';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__tco_var_Control_Monad_List_Trans_take_dictApplicative_0 = $dictApplicative_0;
  tco_loop_Control_Monad_List_Trans_take:;
  $dictApplicative_0 = $__tco_var_Control_Monad_List_Trans_take_dictApplicative_0;
  $nil1_1_0 = (($dictApplicative_0)->{'pure'})(new \Control\Monad\List\Trans\Control_Monad_List_Trans_Done());
  $__local_var_2_1 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = (function() use ($__local_var_2_1, $dictApplicative_0, $nil1_1_0) {
  $__fn = function($v_3, $v1_4 = null) use ($__local_var_2_1, $dictApplicative_0, $nil1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = match ($v_3) { 0 => $nil1_1_0, default => ((($__local_var_2_1)->{'map'})(function($v2_5) use ($dictApplicative_0, $v_3) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ($v2_5 instanceof \Control\Monad\List\Trans\Control_Monad_List_Trans_Yield) {
$__t2 = new \Control\Monad\List\Trans\Control_Monad_List_Trans_Yield(($v2_5)->{'value0'}, ((($GLOBALS['Data_Lazy_functorLazy'])->{'map'})((($GLOBALS['Control_Monad_List_Trans_take'])($dictApplicative_0))(($v_3 - 1))))(($v2_5)->{'value1'}));
goto end_branch_2;;
};
  if ($v2_5 instanceof \Control\Monad\List\Trans\Control_Monad_List_Trans_Skip) {
$__t2 = new \Control\Monad\List\Trans\Control_Monad_List_Trans_Skip(((($GLOBALS['Data_Lazy_functorLazy'])->{'map'})((($GLOBALS['Control_Monad_List_Trans_take'])($dictApplicative_0))($v_3)))(($v2_5)->{'value0'}));
goto end_branch_2;;
};
  if ($v2_5 instanceof \Control\Monad\List\Trans\Control_Monad_List_Trans_Done) {
$__t2 = new \Control\Monad\List\Trans\Control_Monad_List_Trans_Done();
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v1_4) };
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_List_Trans_take'] = __NAMESPACE__ . '\\majControl_majMonad_majList_majTrans_take';

// Control_Monad_List_Trans_zipWith'
function majControl_majMonad_majList_majTrans_zipmajWith__prime__($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majList_majTrans_zipmajWith__prime__';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__tco_var_Control_Monad_List_Trans_zipWith__prime___dictMonad_0 = $dictMonad_0;
  tco_loop_Control_Monad_List_Trans_zipWith__prime__:;
  $dictMonad_0 = $__tco_var_Control_Monad_List_Trans_zipWith__prime___dictMonad_0;
  $Applicative0_1_0 = (($dictMonad_0)->{'Applicative0'})(null);
  $nil1_2_1 = (($Applicative0_1_0)->{'pure'})(new \Control\Monad\List\Trans\Control_Monad_List_Trans_Done());
  $Bind1_3_2 = (($dictMonad_0)->{'Bind1'})(null);
  $Functor0_4_3 = (((($Bind1_3_2)->{'Apply0'})(null))->{'Functor0'})(null);
  $uncons1_5_4 = ($GLOBALS['Control_Monad_List_Trans_uncons'])($dictMonad_0);
  $__res = (function() use ($Applicative0_1_0, $Bind1_3_2, $Functor0_4_3, $dictMonad_0, $nil1_2_1, $uncons1_5_4) {
  $__fn = function($f_6, $fa_7 = null, $fb_8 = null) use ($Applicative0_1_0, $Bind1_3_2, $Functor0_4_3, $dictMonad_0, $nil1_2_1, $uncons1_5_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($Functor0_4_3)->{'map'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_List_Trans_Skip']))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Lazy_defer']))($GLOBALS['Data_Function_const']))))(((($Bind1_3_2)->{'bind'})(($uncons1_5_4)($fa_7)))(function($ua_9) use ($Applicative0_1_0, $Bind1_3_2, $Functor0_4_3, $dictMonad_0, $f_6, $fb_8, $nil1_2_1, $uncons1_5_4) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_2)->{'bind'})(($uncons1_5_4)($fb_8)))(function($ub_10) use ($Applicative0_1_0, $Functor0_4_3, $dictMonad_0, $f_6, $nil1_2_1, $ua_9) {
  $__num = \func_num_args();
  $__t5 = null;;
  if ($ub_10 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t5 = (($Applicative0_1_0)->{'pure'})($nil1_2_1);
goto end_branch_5;;
};
  if ($ua_9 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t5 = (($Applicative0_1_0)->{'pure'})($nil1_2_1);
goto end_branch_5;;
};
  if (($ua_9 instanceof \Data\Maybe\Data_Maybe_Just && $ub_10 instanceof \Data\Maybe\Data_Maybe_Just)) {
$__local_var_11_6 = (($ua_9)->{'value0'})->{'value1'};
$__local_var_12_7 = (($ub_10)->{'value0'})->{'value1'};
$__local_var_13_8 = \Data\Lazy\majData_majLazy_defer(function($v2_13) use ($__local_var_11_6, $__local_var_12_7, $dictMonad_0, $f_6) {
  $__num = \func_num_args();
  $__res = (((($GLOBALS['Control_Monad_List_Trans_zipWith__prime__'])($dictMonad_0))($f_6))($__local_var_11_6))($__local_var_12_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
$__t5 = ((($Functor0_4_3)->{'map'})(function($a_14) use ($Applicative0_1_0, $__local_var_13_8) {
  $__num = \func_num_args();
  $__res = (($Applicative0_1_0)->{'pure'})(new \Control\Monad\List\Trans\Control_Monad_List_Trans_Yield($a_14, $__local_var_13_8));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($f_6)((($ua_9)->{'value0'})->{'value0'}))((($ub_10)->{'value0'})->{'value0'}));
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
}));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_List_Trans_zipWith__prime__'] = __NAMESPACE__ . '\\majControl_majMonad_majList_majTrans_zipmajWith__prime__';

// Control_Monad_List_Trans_zipWith
function majControl_majMonad_majList_majTrans_zipmajWith($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majList_majTrans_zipmajWith';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $zipWith_prime1_1_0 = ($GLOBALS['Control_Monad_List_Trans_zipWith__prime__'])($dictMonad_0);
  $__res = function($f_2) use ($dictMonad_0, $zipWith_prime1_1_0) {
  $__num = \func_num_args();
  $__res = ($zipWith_prime1_1_0)((function() use ($dictMonad_0, $f_2) {
  $__fn = function($a_3, $b_4 = null) use ($dictMonad_0, $f_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($dictMonad_0)->{'Applicative0'})(null))->{'pure'})((($f_2)($a_3))($b_4));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})());
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_List_Trans_zipWith'] = __NAMESPACE__ . '\\majControl_majMonad_majList_majTrans_zipmajWith';

// Control_Monad_List_Trans_newtypeListT
$GLOBALS['Control_Monad_List_Trans_newtypeListT'] = (object)["Coercible0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Control_Monad_List_Trans_mapMaybe
function majControl_majMonad_majList_majTrans_mapmajMaybe($dictFunctor_0, $f_1 = null, $v_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majList_majTrans_mapmajMaybe';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_Control_Monad_List_Trans_mapMaybe_dictFunctor_0 = $dictFunctor_0;
  $__tco_var_Control_Monad_List_Trans_mapMaybe_f_1 = $f_1;
  $__tco_var_Control_Monad_List_Trans_mapMaybe_v_2 = $v_2;
  tco_loop_Control_Monad_List_Trans_mapMaybe:;
  $dictFunctor_0 = $__tco_var_Control_Monad_List_Trans_mapMaybe_dictFunctor_0;
  $f_1 = $__tco_var_Control_Monad_List_Trans_mapMaybe_f_1;
  $v_2 = $__tco_var_Control_Monad_List_Trans_mapMaybe_v_2;
  $__res = ((($dictFunctor_0)->{'map'})(function($v_3) use ($dictFunctor_0, $f_1) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v_3 instanceof \Control\Monad\List\Trans\Control_Monad_List_Trans_Yield) {
$__local_var_4_1 = ((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})($GLOBALS['Control_Monad_List_Trans_Yield']))(($f_1)(($v_3)->{'value0'}));
$__t2 = null;;
if ($__local_var_4_1 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t2 = $GLOBALS['Control_Monad_List_Trans_Skip'];
goto end_branch_2;;
};
if ($__local_var_4_1 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t2 = ($__local_var_4_1)->{'value0'};
goto end_branch_2;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
end_branch_2:;
$__t0 = ($__t2)(((($GLOBALS['Data_Lazy_functorLazy'])->{'map'})((($GLOBALS['Control_Monad_List_Trans_mapMaybe'])($dictFunctor_0))($f_1)))(($v_3)->{'value1'}));
goto end_branch_0;;
};
  if ($v_3 instanceof \Control\Monad\List\Trans\Control_Monad_List_Trans_Skip) {
$__t0 = new \Control\Monad\List\Trans\Control_Monad_List_Trans_Skip(((($GLOBALS['Data_Lazy_functorLazy'])->{'map'})((($GLOBALS['Control_Monad_List_Trans_mapMaybe'])($dictFunctor_0))($f_1)))(($v_3)->{'value0'}));
goto end_branch_0;;
};
  if ($v_3 instanceof \Control\Monad\List\Trans\Control_Monad_List_Trans_Done) {
$__t0 = new \Control\Monad\List\Trans\Control_Monad_List_Trans_Done();
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
$GLOBALS['Control_Monad_List_Trans_mapMaybe'] = __NAMESPACE__ . '\\majControl_majMonad_majList_majTrans_mapmajMaybe';

// Control_Monad_List_Trans_iterate
function majControl_majMonad_majList_majTrans_iterate($dictMonad_0, $f_1 = null, $a_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majList_majTrans_iterate';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = \Control\Monad\List\Trans\majControl_majMonad_majList_majTrans_unfold($dictMonad_0, function($x_3) use ($dictMonad_0, $f_1) {
  $__num = \func_num_args();
  $__res = (((($dictMonad_0)->{'Applicative0'})(null))->{'pure'})(new \Data\Maybe\Data_Maybe_Just(new \Data\Tuple\Data_Tuple_Tuple(($f_1)($x_3), $x_3)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $a_2);
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Control_Monad_List_Trans_iterate'] = __NAMESPACE__ . '\\majControl_majMonad_majList_majTrans_iterate';

// Control_Monad_List_Trans_repeat
function majControl_majMonad_majList_majTrans_repeat($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majList_majTrans_repeat';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Monad_List_Trans_iterate'])($dictMonad_0))($GLOBALS['Control_Monad_List_Trans_identity']);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_List_Trans_repeat'] = __NAMESPACE__ . '\\majControl_majMonad_majList_majTrans_repeat';

// Control_Monad_List_Trans_head
function majControl_majMonad_majList_majTrans_head($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majList_majTrans_head';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $uncons1_1_0 = ($GLOBALS['Control_Monad_List_Trans_uncons'])($dictMonad_0);
  $__res = function($l_2) use ($dictMonad_0, $uncons1_1_0) {
  $__num = \func_num_args();
  $__res = ((((((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})($GLOBALS['Data_Tuple_fst'])))(($uncons1_1_0)($l_2));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_List_Trans_head'] = __NAMESPACE__ . '\\majControl_majMonad_majList_majTrans_head';

// Control_Monad_List_Trans_functorListT
function majControl_majMonad_majList_majTrans_functormajListmajT($dictFunctor_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majList_majTrans_functormajListmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__tco_var_Control_Monad_List_Trans_functorListT_dictFunctor_0 = $dictFunctor_0;
  tco_loop_Control_Monad_List_Trans_functorListT:;
  $dictFunctor_0 = $__tco_var_Control_Monad_List_Trans_functorListT_dictFunctor_0;
  $__res = (object)["map" => (function() use ($dictFunctor_0) {
  $__fn = function($f_1, $v_2 = null) use ($dictFunctor_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictFunctor_0)->{'map'})(function($v_3) use ($dictFunctor_0, $f_1) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v_3 instanceof \Control\Monad\List\Trans\Control_Monad_List_Trans_Yield) {
$__t0 = new \Control\Monad\List\Trans\Control_Monad_List_Trans_Yield(($f_1)(($v_3)->{'value0'}), ((($GLOBALS['Data_Lazy_functorLazy'])->{'map'})(((\Control\Monad\List\Trans\majControl_majMonad_majList_majTrans_functormajListmajT($dictFunctor_0))->{'map'})($f_1)))(($v_3)->{'value1'}));
goto end_branch_0;;
};
  if ($v_3 instanceof \Control\Monad\List\Trans\Control_Monad_List_Trans_Skip) {
$__t0 = new \Control\Monad\List\Trans\Control_Monad_List_Trans_Skip(((($GLOBALS['Data_Lazy_functorLazy'])->{'map'})(((\Control\Monad\List\Trans\majControl_majMonad_majList_majTrans_functormajListmajT($dictFunctor_0))->{'map'})($f_1)))(($v_3)->{'value0'}));
goto end_branch_0;;
};
  if ($v_3 instanceof \Control\Monad\List\Trans\Control_Monad_List_Trans_Done) {
$__t0 = new \Control\Monad\List\Trans\Control_Monad_List_Trans_Done();
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
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_List_Trans_functorListT'] = __NAMESPACE__ . '\\majControl_majMonad_majList_majTrans_functormajListmajT';

// Control_Monad_List_Trans_fromEffect
function majControl_majMonad_majList_majTrans_frommajEffect($dictApplicative_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majList_majTrans_frommajEffect';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $nil1_1_0 = (($dictApplicative_0)->{'pure'})(new \Control\Monad\List\Trans\Control_Monad_List_Trans_Done());
  $__res = function($fa_2) use ($dictApplicative_0, $nil1_1_0) {
  $__num = \func_num_args();
  $__local_var_3_1 = \Data\Lazy\majData_majLazy_defer(function($v_3) use ($nil1_1_0) {
  $__num = \func_num_args();
  $__res = $nil1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = ((((((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})(function($a_4) use ($__local_var_3_1) {
  $__num = \func_num_args();
  $__res = new \Control\Monad\List\Trans\Control_Monad_List_Trans_Yield($a_4, $__local_var_3_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($fa_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_List_Trans_fromEffect'] = __NAMESPACE__ . '\\majControl_majMonad_majList_majTrans_frommajEffect';

// Control_Monad_List_Trans_monadTransListT
$GLOBALS['Control_Monad_List_Trans_monadTransListT'] = (object)["lift" => function($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_List_Trans_fromEffect'])((($dictMonad_0)->{'Applicative0'})(null));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Control_Monad_List_Trans_foldlRec'
function majControl_majMonad_majList_majTrans_foldlmajRec__prime__($dictMonadRec_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majList_majTrans_foldlmajRec__prime__';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Monad0_1_0 = (($dictMonadRec_0)->{'Monad0'})(null);
  $__local_var_2_1 = (($Monad0_1_0)->{'Applicative0'})(null);
  $__local_var_3_2 = (($Monad0_1_0)->{'Bind1'})(null);
  $uncons1_4_3 = ($GLOBALS['Control_Monad_List_Trans_uncons'])($Monad0_1_0);
  $__res = (function() use ($__local_var_2_1, $__local_var_3_2, $dictMonadRec_0, $uncons1_4_3) {
  $__fn = function($f_5, $a_6 = null, $b_7 = null) use ($__local_var_2_1, $__local_var_3_2, $dictMonadRec_0, $uncons1_4_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictMonadRec_0)->{'tailRecM'})(function($o_8) use ($__local_var_2_1, $__local_var_3_2, $f_5, $uncons1_4_3) {
  $__num = \func_num_args();
  $__local_var_9_4 = ($o_8)->{'a'};
  $__res = ((($__local_var_3_2)->{'bind'})(($uncons1_4_3)(($o_8)->{'b'})))(function($v_10) use ($__local_var_2_1, $__local_var_3_2, $__local_var_9_4, $f_5) {
  $__num = \func_num_args();
  $__t5 = null;;
  if ($v_10 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t5 = (($__local_var_2_1)->{'pure'})(new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Done($__local_var_9_4));
goto end_branch_5;;
};
  if ($v_10 instanceof \Data\Maybe\Data_Maybe_Just) {
$__local_var_11_6 = (($v_10)->{'value0'})->{'value1'};
$__t5 = ((($__local_var_3_2)->{'bind'})((($f_5)($__local_var_9_4))((($v_10)->{'value0'})->{'value0'})))(function($b_prime_12) use ($__local_var_11_6, $__local_var_2_1) {
  $__num = \func_num_args();
  $__res = (($__local_var_2_1)->{'pure'})(new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Loop((object)["a" => $b_prime_12, "b" => $__local_var_11_6]));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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
}))((object)["a" => $a_6, "b" => $b_7]);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_List_Trans_foldlRec__prime__'] = __NAMESPACE__ . '\\majControl_majMonad_majList_majTrans_foldlmajRec__prime__';

// Control_Monad_List_Trans_runListTRec
function majControl_majMonad_majList_majTrans_runmajListmajTmajRec($dictMonadRec_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majList_majTrans_runmajListmajTmajRec';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ((($GLOBALS['Control_Monad_List_Trans_foldlRec__prime__'])($dictMonadRec_0))((function() use ($dictMonadRec_0) {
  $__fn = function($v_1, $v1_2 = null) use ($dictMonadRec_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((((($dictMonadRec_0)->{'Monad0'})(null))->{'Applicative0'})(null))->{'pure'})($GLOBALS['Data_Unit_unit']);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($GLOBALS['Data_Unit_unit']);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_List_Trans_runListTRec'] = __NAMESPACE__ . '\\majControl_majMonad_majList_majTrans_runmajListmajTmajRec';

// Control_Monad_List_Trans_foldlRec
function majControl_majMonad_majList_majTrans_foldlmajRec($dictMonadRec_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majList_majTrans_foldlmajRec';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Monad0_1_0 = (($dictMonadRec_0)->{'Monad0'})(null);
  $__local_var_2_1 = (($Monad0_1_0)->{'Applicative0'})(null);
  $uncons1_3_2 = ($GLOBALS['Control_Monad_List_Trans_uncons'])($Monad0_1_0);
  $__res = (function() use ($Monad0_1_0, $__local_var_2_1, $dictMonadRec_0, $uncons1_3_2) {
  $__fn = function($f_4, $a_5 = null, $b_6 = null) use ($Monad0_1_0, $__local_var_2_1, $dictMonadRec_0, $uncons1_3_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictMonadRec_0)->{'tailRecM'})(function($o_7) use ($Monad0_1_0, $__local_var_2_1, $f_4, $uncons1_3_2) {
  $__num = \func_num_args();
  $__local_var_8_3 = ($o_7)->{'a'};
  $__res = ((((($Monad0_1_0)->{'Bind1'})(null))->{'bind'})(($uncons1_3_2)(($o_7)->{'b'})))(function($v_9) use ($__local_var_2_1, $__local_var_8_3, $f_4) {
  $__num = \func_num_args();
  $__t4 = null;;
  if ($v_9 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t4 = (($__local_var_2_1)->{'pure'})(new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Done($__local_var_8_3));
goto end_branch_4;;
};
  if ($v_9 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t4 = (($__local_var_2_1)->{'pure'})(new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Loop((object)["a" => (($f_4)($__local_var_8_3))((($v_9)->{'value0'})->{'value0'}), "b" => (($v_9)->{'value0'})->{'value1'}]));
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((object)["a" => $a_5, "b" => $b_6]);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_List_Trans_foldlRec'] = __NAMESPACE__ . '\\majControl_majMonad_majList_majTrans_foldlmajRec';

// Control_Monad_List_Trans_foldl'
function majControl_majMonad_majList_majTrans_foldl__prime__($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majList_majTrans_foldl__prime__';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictMonad_0)->{'Bind1'})(null);
  $uncons1_2_1 = ($GLOBALS['Control_Monad_List_Trans_uncons'])($dictMonad_0);
  $__res = function($f_3) use ($__local_var_1_0, $dictMonad_0, $uncons1_2_1) {
  $__num = \func_num_args();
  $loop_4_2 = null;
  $loop_4_2 = (function() use ($__local_var_1_0, $dictMonad_0, $f_3, &$loop_4_2, $uncons1_2_1) {
  $__fn = function($b_5, $l_6 = null) use ($__local_var_1_0, $dictMonad_0, $f_3, &$loop_4_2, $uncons1_2_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_1_0)->{'bind'})(($uncons1_2_1)($l_6)))(function($v_7) use ($__local_var_1_0, $b_5, $dictMonad_0, $f_3, &$loop_4_2) {
  $__num = \func_num_args();
  $__t3 = null;;
  if ($v_7 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t3 = (((($dictMonad_0)->{'Applicative0'})(null))->{'pure'})($b_5);
goto end_branch_3;;
};
  if ($v_7 instanceof \Data\Maybe\Data_Maybe_Just) {
$__local_var_8_4 = (($v_7)->{'value0'})->{'value1'};
$__t3 = ((($__local_var_1_0)->{'bind'})((($f_3)($b_5))((($v_7)->{'value0'})->{'value0'})))(function($a_9) use ($__local_var_8_4, &$loop_4_2) {
  $__num = \func_num_args();
  $__res = (($loop_4_2)($a_9))($__local_var_8_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = $loop_4_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_List_Trans_foldl__prime__'] = __NAMESPACE__ . '\\majControl_majMonad_majList_majTrans_foldl__prime__';

// Control_Monad_List_Trans_runListT
function majControl_majMonad_majList_majTrans_runmajListmajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majList_majTrans_runmajListmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ((($GLOBALS['Control_Monad_List_Trans_foldl__prime__'])($dictMonad_0))((function() use ($dictMonad_0) {
  $__fn = function($v_1, $v1_2 = null) use ($dictMonad_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($dictMonad_0)->{'Applicative0'})(null))->{'pure'})($GLOBALS['Data_Unit_unit']);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($GLOBALS['Data_Unit_unit']);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_List_Trans_runListT'] = __NAMESPACE__ . '\\majControl_majMonad_majList_majTrans_runmajListmajT';

// Control_Monad_List_Trans_foldl
function majControl_majMonad_majList_majTrans_foldl($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majList_majTrans_foldl';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $uncons1_1_0 = ($GLOBALS['Control_Monad_List_Trans_uncons'])($dictMonad_0);
  $__res = function($f_2) use ($dictMonad_0, $uncons1_1_0) {
  $__num = \func_num_args();
  $loop_3_1 = null;
  $loop_3_1 = (function() use ($dictMonad_0, $f_2, &$loop_3_1, $uncons1_1_0) {
  $__fn = function($b_4, $l_5 = null) use ($dictMonad_0, $f_2, &$loop_3_1, $uncons1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((($dictMonad_0)->{'Bind1'})(null))->{'bind'})(($uncons1_1_0)($l_5)))(function($v_6) use ($b_4, $dictMonad_0, $f_2, &$loop_3_1) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ($v_6 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t2 = (((($dictMonad_0)->{'Applicative0'})(null))->{'pure'})($b_4);
goto end_branch_2;;
};
  if ($v_6 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t2 = (($loop_3_1)((($f_2)($b_4))((($v_6)->{'value0'})->{'value0'})))((($v_6)->{'value0'})->{'value1'});
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
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = $loop_3_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_List_Trans_foldl'] = __NAMESPACE__ . '\\majControl_majMonad_majList_majTrans_foldl';

// Control_Monad_List_Trans_filter
function majControl_majMonad_majList_majTrans_filter($dictFunctor_0, $f_1 = null, $v_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majList_majTrans_filter';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_Control_Monad_List_Trans_filter_dictFunctor_0 = $dictFunctor_0;
  $__tco_var_Control_Monad_List_Trans_filter_f_1 = $f_1;
  $__tco_var_Control_Monad_List_Trans_filter_v_2 = $v_2;
  tco_loop_Control_Monad_List_Trans_filter:;
  $dictFunctor_0 = $__tco_var_Control_Monad_List_Trans_filter_dictFunctor_0;
  $f_1 = $__tco_var_Control_Monad_List_Trans_filter_f_1;
  $v_2 = $__tco_var_Control_Monad_List_Trans_filter_v_2;
  $__res = ((($dictFunctor_0)->{'map'})(function($v_3) use ($dictFunctor_0, $f_1) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v_3 instanceof \Control\Monad\List\Trans\Control_Monad_List_Trans_Yield) {
$s_prime_4_1 = ((($GLOBALS['Data_Lazy_functorLazy'])->{'map'})((($GLOBALS['Control_Monad_List_Trans_filter'])($dictFunctor_0))($f_1)))(($v_3)->{'value1'});
$__t2 = null;;
if (($f_1)(($v_3)->{'value0'})) {
$__t2 = new \Control\Monad\List\Trans\Control_Monad_List_Trans_Yield(($v_3)->{'value0'}, $s_prime_4_1);
goto end_branch_2;;
};
$__t2 = new \Control\Monad\List\Trans\Control_Monad_List_Trans_Skip($s_prime_4_1);
end_branch_2:;
$__t0 = $__t2;
goto end_branch_0;;
};
  if ($v_3 instanceof \Control\Monad\List\Trans\Control_Monad_List_Trans_Skip) {
$__t0 = new \Control\Monad\List\Trans\Control_Monad_List_Trans_Skip(((($GLOBALS['Data_Lazy_functorLazy'])->{'map'})((($GLOBALS['Control_Monad_List_Trans_filter'])($dictFunctor_0))($f_1)))(($v_3)->{'value0'}));
goto end_branch_0;;
};
  if ($v_3 instanceof \Control\Monad\List\Trans\Control_Monad_List_Trans_Done) {
$__t0 = new \Control\Monad\List\Trans\Control_Monad_List_Trans_Done();
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
$GLOBALS['Control_Monad_List_Trans_filter'] = __NAMESPACE__ . '\\majControl_majMonad_majList_majTrans_filter';

// Control_Monad_List_Trans_dropWhile
function majControl_majMonad_majList_majTrans_dropmajWhile($dictApplicative_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majList_majTrans_dropmajWhile';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__tco_var_Control_Monad_List_Trans_dropWhile_dictApplicative_0 = $dictApplicative_0;
  tco_loop_Control_Monad_List_Trans_dropWhile:;
  $dictApplicative_0 = $__tco_var_Control_Monad_List_Trans_dropWhile_dictApplicative_0;
  $__local_var_1_0 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = (function() use ($__local_var_1_0, $dictApplicative_0) {
  $__fn = function($f_2, $v_3 = null) use ($__local_var_1_0, $dictApplicative_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_1_0)->{'map'})(function($v_4) use ($dictApplicative_0, $f_2) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v_4 instanceof \Control\Monad\List\Trans\Control_Monad_List_Trans_Yield) {
$__t2 = null;;
if (($f_2)(($v_4)->{'value0'})) {
$__t2 = new \Control\Monad\List\Trans\Control_Monad_List_Trans_Skip(((($GLOBALS['Data_Lazy_functorLazy'])->{'map'})((($GLOBALS['Control_Monad_List_Trans_dropWhile'])($dictApplicative_0))($f_2)))(($v_4)->{'value1'}));
goto end_branch_2;;
};
$__t2 = new \Control\Monad\List\Trans\Control_Monad_List_Trans_Yield(($v_4)->{'value0'}, ($v_4)->{'value1'});
end_branch_2:;
$__t1 = $__t2;
goto end_branch_1;;
};
  if ($v_4 instanceof \Control\Monad\List\Trans\Control_Monad_List_Trans_Skip) {
$__t1 = new \Control\Monad\List\Trans\Control_Monad_List_Trans_Skip(((($GLOBALS['Data_Lazy_functorLazy'])->{'map'})((($GLOBALS['Control_Monad_List_Trans_dropWhile'])($dictApplicative_0))($f_2)))(($v_4)->{'value0'}));
goto end_branch_1;;
};
  if ($v_4 instanceof \Control\Monad\List\Trans\Control_Monad_List_Trans_Done) {
$__t1 = new \Control\Monad\List\Trans\Control_Monad_List_Trans_Done();
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_List_Trans_dropWhile'] = __NAMESPACE__ . '\\majControl_majMonad_majList_majTrans_dropmajWhile';

// Control_Monad_List_Trans_drop
function majControl_majMonad_majList_majTrans_drop($dictApplicative_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majList_majTrans_drop';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__tco_var_Control_Monad_List_Trans_drop_dictApplicative_0 = $dictApplicative_0;
  tco_loop_Control_Monad_List_Trans_drop:;
  $dictApplicative_0 = $__tco_var_Control_Monad_List_Trans_drop_dictApplicative_0;
  $__local_var_1_0 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = (function() use ($__local_var_1_0, $dictApplicative_0) {
  $__fn = function($v_2, $v1_3 = null) use ($__local_var_1_0, $dictApplicative_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = match ($v_2) { 0 => $v1_3, default => ((($__local_var_1_0)->{'map'})(function($v2_4) use ($dictApplicative_0, $v_2) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v2_4 instanceof \Control\Monad\List\Trans\Control_Monad_List_Trans_Yield) {
$__t1 = new \Control\Monad\List\Trans\Control_Monad_List_Trans_Skip(((($GLOBALS['Data_Lazy_functorLazy'])->{'map'})((($GLOBALS['Control_Monad_List_Trans_drop'])($dictApplicative_0))(($v_2 - 1))))(($v2_4)->{'value1'}));
goto end_branch_1;;
};
  if ($v2_4 instanceof \Control\Monad\List\Trans\Control_Monad_List_Trans_Skip) {
$__t1 = new \Control\Monad\List\Trans\Control_Monad_List_Trans_Skip(((($GLOBALS['Data_Lazy_functorLazy'])->{'map'})((($GLOBALS['Control_Monad_List_Trans_drop'])($dictApplicative_0))($v_2)))(($v2_4)->{'value0'}));
goto end_branch_1;;
};
  if ($v2_4 instanceof \Control\Monad\List\Trans\Control_Monad_List_Trans_Done) {
$__t1 = new \Control\Monad\List\Trans\Control_Monad_List_Trans_Done();
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v1_3) };
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_List_Trans_drop'] = __NAMESPACE__ . '\\majControl_majMonad_majList_majTrans_drop';

// Control_Monad_List_Trans_cons
function majControl_majMonad_majList_majTrans_cons($dictApplicative_0, $lh_1 = null, $t_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majList_majTrans_cons';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($dictApplicative_0)->{'pure'})(new \Control\Monad\List\Trans\Control_Monad_List_Trans_Yield(\Data\Lazy\majData_majLazy_force($lh_1), $t_2));
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Control_Monad_List_Trans_cons'] = __NAMESPACE__ . '\\majControl_majMonad_majList_majTrans_cons';

// Control_Monad_List_Trans_unfoldable1ListT
function majControl_majMonad_majList_majTrans_unfoldable1majListmajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majList_majTrans_unfoldable1majListmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Applicative0_1_0 = (($dictMonad_0)->{'Applicative0'})(null);
  $singleton1_2_1 = ($GLOBALS['Control_Monad_List_Trans_singleton'])($Applicative0_1_0);
  $__res = (object)["unfoldr1" => (function() use ($Applicative0_1_0, $singleton1_2_1) {
  $__fn = function($f_3, $b_4 = null) use ($Applicative0_1_0, $singleton1_2_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__5_2 = null;
  $go__5_2 = function($v_6) use ($Applicative0_1_0, $f_3, &$go__5_2, $singleton1_2_1) {
  $__num = \func_num_args();
  $__t3 = null;;
  if (($v_6)->{'value1'} instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t3 = ($singleton1_2_1)(($v_6)->{'value0'});
goto end_branch_3;;
};
  if (($v_6)->{'value1'} instanceof \Data\Maybe\Data_Maybe_Just) {
$__local_var_7_4 = (($v_6)->{'value1'})->{'value0'};
$__t3 = (($Applicative0_1_0)->{'pure'})(new \Control\Monad\List\Trans\Control_Monad_List_Trans_Yield(\Data\Lazy\majData_majLazy_force((($GLOBALS['Data_Lazy_applicativeLazy'])->{'pure'})(($v_6)->{'value0'})), \Data\Lazy\majData_majLazy_defer(function($v1_8) use ($__local_var_7_4, $f_3, &$go__5_2) {
  $__num = \func_num_args();
  $__res = ($go__5_2)(($f_3)($__local_var_7_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})));
goto end_branch_3;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t3 = null;
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = ($go__5_2)(($f_3)($b_4));
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
$GLOBALS['Control_Monad_List_Trans_unfoldable1ListT'] = __NAMESPACE__ . '\\majControl_majMonad_majList_majTrans_unfoldable1majListmajT';

// Control_Monad_List_Trans_unfoldableListT
function majControl_majMonad_majList_majTrans_unfoldablemajListmajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majList_majTrans_unfoldablemajListmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Applicative0_1_0 = (($dictMonad_0)->{'Applicative0'})(null);
  $nil1_2_1 = (($Applicative0_1_0)->{'pure'})(new \Control\Monad\List\Trans\Control_Monad_List_Trans_Done());
  $unfoldable1ListT1_3_2 = \Control\Monad\List\Trans\majControl_majMonad_majList_majTrans_unfoldable1majListmajT($dictMonad_0);
  $__res = (object)["unfoldr" => (function() use ($Applicative0_1_0, $nil1_2_1) {
  $__fn = function($f_4, $b_5 = null) use ($Applicative0_1_0, $nil1_2_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__6_3 = null;
  $go__6_3 = function($v_7) use ($Applicative0_1_0, $f_4, &$go__6_3, $nil1_2_1) {
  $__num = \func_num_args();
  $__t4 = null;;
  if ($v_7 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t4 = $nil1_2_1;
goto end_branch_4;;
};
  if ($v_7 instanceof \Data\Maybe\Data_Maybe_Just) {
$__local_var_8_5 = (($v_7)->{'value0'})->{'value1'};
$__t4 = (($Applicative0_1_0)->{'pure'})(new \Control\Monad\List\Trans\Control_Monad_List_Trans_Yield(\Data\Lazy\majData_majLazy_force((($GLOBALS['Data_Lazy_applicativeLazy'])->{'pure'})((($v_7)->{'value0'})->{'value0'})), \Data\Lazy\majData_majLazy_defer(function($v1_9) use ($__local_var_8_5, $f_4, &$go__6_3) {
  $__num = \func_num_args();
  $__res = ($go__6_3)(($f_4)($__local_var_8_5));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})));
goto end_branch_4;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t4 = null;
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = ($go__6_3)(($f_4)($b_5));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Unfoldable10" => function($_dollar__unused_4) use ($unfoldable1ListT1_3_2) {
  $__num = \func_num_args();
  $__res = $unfoldable1ListT1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_List_Trans_unfoldableListT'] = __NAMESPACE__ . '\\majControl_majMonad_majList_majTrans_unfoldablemajListmajT';

// Control_Monad_List_Trans_semigroupListT
function majControl_majMonad_majList_majTrans_semigroupmajListmajT($dictApplicative_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majList_majTrans_semigroupmajListmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["append" => ($GLOBALS['Control_Monad_List_Trans_concat'])($dictApplicative_0)];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_List_Trans_semigroupListT'] = __NAMESPACE__ . '\\majControl_majMonad_majList_majTrans_semigroupmajListmajT';

// Control_Monad_List_Trans_concat
function majControl_majMonad_majList_majTrans_concat($dictApplicative_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majList_majTrans_concat';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = (function() use ($__local_var_1_0, $dictApplicative_0) {
  $__fn = function($x_2, $y_3 = null) use ($__local_var_1_0, $dictApplicative_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_1_0)->{'map'})(function($v_4) use ($dictApplicative_0, $y_3) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v_4 instanceof \Control\Monad\List\Trans\Control_Monad_List_Trans_Yield) {
$__t1 = new \Control\Monad\List\Trans\Control_Monad_List_Trans_Yield(($v_4)->{'value0'}, ((($GLOBALS['Data_Lazy_functorLazy'])->{'map'})(function($v1_5) use ($dictApplicative_0, $y_3) {
  $__num = \func_num_args();
  $__res = \Control\Monad\List\Trans\majControl_majMonad_majList_majTrans_concat($dictApplicative_0, $v1_5, $y_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_4)->{'value1'}));
goto end_branch_1;;
};
  if ($v_4 instanceof \Control\Monad\List\Trans\Control_Monad_List_Trans_Skip) {
$__t1 = new \Control\Monad\List\Trans\Control_Monad_List_Trans_Skip(((($GLOBALS['Data_Lazy_functorLazy'])->{'map'})(function($v1_5) use ($dictApplicative_0, $y_3) {
  $__num = \func_num_args();
  $__res = \Control\Monad\List\Trans\majControl_majMonad_majList_majTrans_concat($dictApplicative_0, $v1_5, $y_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_4)->{'value0'}));
goto end_branch_1;;
};
  if ($v_4 instanceof \Control\Monad\List\Trans\Control_Monad_List_Trans_Done) {
$__t1 = new \Control\Monad\List\Trans\Control_Monad_List_Trans_Skip(\Data\Lazy\majData_majLazy_defer(function($v_5) use ($y_3) {
  $__num = \func_num_args();
  $__res = $y_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($x_2);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_List_Trans_concat'] = __NAMESPACE__ . '\\majControl_majMonad_majList_majTrans_concat';

// Control_Monad_List_Trans_monoidListT
function majControl_majMonad_majList_majTrans_monoidmajListmajT($dictApplicative_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majList_majTrans_monoidmajListmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $semigroupListT1_1_0 = (object)["append" => ($GLOBALS['Control_Monad_List_Trans_concat'])($dictApplicative_0)];
  $__res = (object)["mempty" => (($dictApplicative_0)->{'pure'})(new \Control\Monad\List\Trans\Control_Monad_List_Trans_Done()), "Semigroup0" => function($_dollar__unused_2) use ($semigroupListT1_1_0) {
  $__num = \func_num_args();
  $__res = $semigroupListT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_List_Trans_monoidListT'] = __NAMESPACE__ . '\\majControl_majMonad_majList_majTrans_monoidmajListmajT';

// Control_Monad_List_Trans_catMaybes
function majControl_majMonad_majList_majTrans_catmajMaybes($dictFunctor_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majList_majTrans_catmajMaybes';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Monad_List_Trans_mapMaybe'])($dictFunctor_0))($GLOBALS['Control_Monad_List_Trans_identity']);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_List_Trans_catMaybes'] = __NAMESPACE__ . '\\majControl_majMonad_majList_majTrans_catmajMaybes';

// Control_Monad_List_Trans_monadListT
function majControl_majMonad_majList_majTrans_monadmajListmajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majList_majTrans_monadmajListmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["Applicative0" => function($_dollar__unused_1) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = \Control\Monad\List\Trans\majControl_majMonad_majList_majTrans_applicativemajListmajT($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar__unused_1) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = \Control\Monad\List\Trans\majControl_majMonad_majList_majTrans_bindmajListmajT($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_List_Trans_monadListT'] = __NAMESPACE__ . '\\majControl_majMonad_majList_majTrans_monadmajListmajT';

// Control_Monad_List_Trans_bindListT
function majControl_majMonad_majList_majTrans_bindmajListmajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majList_majTrans_bindmajListmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $append_1_0 = ($GLOBALS['Control_Monad_List_Trans_concat'])((($dictMonad_0)->{'Applicative0'})(null));
  $__local_var_2_1 = (((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = (object)["bind" => (function() use ($__local_var_2_1, $append_1_0, $dictMonad_0) {
  $__fn = function($fa_3, $f_4 = null) use ($__local_var_2_1, $append_1_0, $dictMonad_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_2_1)->{'map'})(function($v_5) use ($append_1_0, $dictMonad_0, $f_4) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ($v_5 instanceof \Control\Monad\List\Trans\Control_Monad_List_Trans_Yield) {
$__local_var_6_3 = ($v_5)->{'value0'};
$__t2 = new \Control\Monad\List\Trans\Control_Monad_List_Trans_Skip(((($GLOBALS['Data_Lazy_functorLazy'])->{'map'})(function($s_prime_7) use ($__local_var_6_3, $append_1_0, $dictMonad_0, $f_4) {
  $__num = \func_num_args();
  $__res = (($append_1_0)(($f_4)($__local_var_6_3)))((((\Control\Monad\List\Trans\majControl_majMonad_majList_majTrans_bindmajListmajT($dictMonad_0))->{'bind'})($s_prime_7))($f_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_5)->{'value1'}));
goto end_branch_2;;
};
  if ($v_5 instanceof \Control\Monad\List\Trans\Control_Monad_List_Trans_Skip) {
$__t2 = new \Control\Monad\List\Trans\Control_Monad_List_Trans_Skip(((($GLOBALS['Data_Lazy_functorLazy'])->{'map'})(function($v1_6) use ($dictMonad_0, $f_4) {
  $__num = \func_num_args();
  $__res = (((\Control\Monad\List\Trans\majControl_majMonad_majList_majTrans_bindmajListmajT($dictMonad_0))->{'bind'})($v1_6))($f_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_5)->{'value0'}));
goto end_branch_2;;
};
  if ($v_5 instanceof \Control\Monad\List\Trans\Control_Monad_List_Trans_Done) {
$__t2 = new \Control\Monad\List\Trans\Control_Monad_List_Trans_Done();
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($fa_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Apply0" => function($_dollar__unused_3) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = \Control\Monad\List\Trans\majControl_majMonad_majList_majTrans_applymajListmajT($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_List_Trans_bindListT'] = __NAMESPACE__ . '\\majControl_majMonad_majList_majTrans_bindmajListmajT';

// Control_Monad_List_Trans_applyListT
function majControl_majMonad_majList_majTrans_applymajListmajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majList_majTrans_applymajListmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $functorListT1_1_0 = \Control\Monad\List\Trans\majControl_majMonad_majList_majTrans_functormajListmajT((((((($dictMonad_0)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null));
  $__local_var_2_1 = \Control\Monad\List\Trans\majControl_majMonad_majList_majTrans_bindmajListmajT($dictMonad_0);
  $__res = (object)["apply" => (function() use ($__local_var_2_1, $dictMonad_0) {
  $__fn = function($f_3, $a_4 = null) use ($__local_var_2_1, $dictMonad_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_2_1)->{'bind'})($f_3))(function($f_prime_5) use ($__local_var_2_1, $a_4, $dictMonad_0) {
  $__num = \func_num_args();
  $__res = ((($__local_var_2_1)->{'bind'})($a_4))(function($a_prime_6) use ($dictMonad_0, $f_prime_5) {
  $__num = \func_num_args();
  $__res = ((\Control\Monad\List\Trans\majControl_majMonad_majList_majTrans_applicativemajListmajT($dictMonad_0))->{'pure'})(($f_prime_5)($a_prime_6));
  goto __end;;
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
})(), "Functor0" => function($_dollar__unused_2) use ($functorListT1_1_0) {
  $__num = \func_num_args();
  $__res = $functorListT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_List_Trans_applyListT'] = __NAMESPACE__ . '\\majControl_majMonad_majList_majTrans_applymajListmajT';

// Control_Monad_List_Trans_applicativeListT
function majControl_majMonad_majList_majTrans_applicativemajListmajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majList_majTrans_applicativemajListmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["pure" => ($GLOBALS['Control_Monad_List_Trans_singleton'])((($dictMonad_0)->{'Applicative0'})(null)), "Apply0" => function($_dollar__unused_1) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = \Control\Monad\List\Trans\majControl_majMonad_majList_majTrans_applymajListmajT($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_List_Trans_applicativeListT'] = __NAMESPACE__ . '\\majControl_majMonad_majList_majTrans_applicativemajListmajT';

// Control_Monad_List_Trans_monadEffectListT
function majControl_majMonad_majList_majTrans_monadmajEffectmajListmajT($dictMonadEffect_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majList_majTrans_monadmajEffectmajListmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Monad0_1_0 = (($dictMonadEffect_0)->{'Monad0'})(null);
  $monadListT1_2_1 = (object)["Applicative0" => function($_dollar__unused_2) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = (object)["pure" => ($GLOBALS['Control_Monad_List_Trans_singleton'])((($Monad0_1_0)->{'Applicative0'})(null)), "Apply0" => function($_dollar__unused_3) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = \Control\Monad\List\Trans\majControl_majMonad_majList_majTrans_applymajListmajT($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar__unused_2) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = \Control\Monad\List\Trans\majControl_majMonad_majList_majTrans_bindmajListmajT($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["liftEffect" => (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Control_Monad_List_Trans_monadTransListT'])->{'lift'})($Monad0_1_0)))(($dictMonadEffect_0)->{'liftEffect'}), "Monad0" => function($_dollar__unused_3) use ($monadListT1_2_1) {
  $__num = \func_num_args();
  $__res = $monadListT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_List_Trans_monadEffectListT'] = __NAMESPACE__ . '\\majControl_majMonad_majList_majTrans_monadmajEffectmajListmajT';

// Control_Monad_List_Trans_monadSTListT
function majControl_majMonad_majList_majTrans_monadmajSmajTmajListmajT($dictMonadST_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majList_majTrans_monadmajSmajTmajListmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Monad0_1_0 = (($dictMonadST_0)->{'Monad0'})(null);
  $monadListT1_2_1 = (object)["Applicative0" => function($_dollar__unused_2) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = (object)["pure" => ($GLOBALS['Control_Monad_List_Trans_singleton'])((($Monad0_1_0)->{'Applicative0'})(null)), "Apply0" => function($_dollar__unused_3) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = \Control\Monad\List\Trans\majControl_majMonad_majList_majTrans_applymajListmajT($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar__unused_2) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = \Control\Monad\List\Trans\majControl_majMonad_majList_majTrans_bindmajListmajT($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["liftST" => (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Control_Monad_List_Trans_monadTransListT'])->{'lift'})($Monad0_1_0)))(($dictMonadST_0)->{'liftST'}), "Monad0" => function($_dollar__unused_3) use ($monadListT1_2_1) {
  $__num = \func_num_args();
  $__res = $monadListT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_List_Trans_monadSTListT'] = __NAMESPACE__ . '\\majControl_majMonad_majList_majTrans_monadmajSmajTmajListmajT';

// Control_Monad_List_Trans_altListT
function majControl_majMonad_majList_majTrans_altmajListmajT($dictApplicative_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majList_majTrans_altmajListmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $functorListT1_1_0 = \Control\Monad\List\Trans\majControl_majMonad_majList_majTrans_functormajListmajT((((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null));
  $__res = (object)["alt" => ($GLOBALS['Control_Monad_List_Trans_concat'])($dictApplicative_0), "Functor0" => function($_dollar__unused_2) use ($functorListT1_1_0) {
  $__num = \func_num_args();
  $__res = $functorListT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_List_Trans_altListT'] = __NAMESPACE__ . '\\majControl_majMonad_majList_majTrans_altmajListmajT';

// Control_Monad_List_Trans_plusListT
function majControl_majMonad_majList_majTrans_plusmajListmajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majList_majTrans_plusmajListmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Applicative0_1_0 = (($dictMonad_0)->{'Applicative0'})(null);
  $altListT1_2_1 = \Control\Monad\List\Trans\majControl_majMonad_majList_majTrans_altmajListmajT($Applicative0_1_0);
  $__res = (object)["empty" => (($Applicative0_1_0)->{'pure'})(new \Control\Monad\List\Trans\Control_Monad_List_Trans_Done()), "Alt0" => function($_dollar__unused_3) use ($altListT1_2_1) {
  $__num = \func_num_args();
  $__res = $altListT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_List_Trans_plusListT'] = __NAMESPACE__ . '\\majControl_majMonad_majList_majTrans_plusmajListmajT';

// Control_Monad_List_Trans_alternativeListT
function majControl_majMonad_majList_majTrans_alternativemajListmajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majList_majTrans_alternativemajListmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $applicativeListT1_1_0 = (object)["pure" => ($GLOBALS['Control_Monad_List_Trans_singleton'])((($dictMonad_0)->{'Applicative0'})(null)), "Apply0" => function($_dollar__unused_1) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = \Control\Monad\List\Trans\majControl_majMonad_majList_majTrans_applymajListmajT($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $plusListT1_2_1 = \Control\Monad\List\Trans\majControl_majMonad_majList_majTrans_plusmajListmajT($dictMonad_0);
  $__res = (object)["Applicative0" => function($_dollar__unused_3) use ($applicativeListT1_1_0) {
  $__num = \func_num_args();
  $__res = $applicativeListT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Plus1" => function($_dollar__unused_3) use ($plusListT1_2_1) {
  $__num = \func_num_args();
  $__res = $plusListT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_List_Trans_alternativeListT'] = __NAMESPACE__ . '\\majControl_majMonad_majList_majTrans_alternativemajListmajT';

// Control_Monad_List_Trans_monadPlusListT
function majControl_majMonad_majList_majTrans_monadmajPlusmajListmajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majList_majTrans_monadmajPlusmajListmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $monadListT1_1_0 = (object)["Applicative0" => function($_dollar__unused_1) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = (object)["pure" => ($GLOBALS['Control_Monad_List_Trans_singleton'])((($dictMonad_0)->{'Applicative0'})(null)), "Apply0" => function($_dollar__unused_2) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = \Control\Monad\List\Trans\majControl_majMonad_majList_majTrans_applymajListmajT($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar__unused_1) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = \Control\Monad\List\Trans\majControl_majMonad_majList_majTrans_bindmajListmajT($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $alternativeListT1_2_1 = \Control\Monad\List\Trans\majControl_majMonad_majList_majTrans_alternativemajListmajT($dictMonad_0);
  $__res = (object)["Monad0" => function($_dollar__unused_3) use ($monadListT1_1_0) {
  $__num = \func_num_args();
  $__res = $monadListT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Alternative1" => function($_dollar__unused_3) use ($alternativeListT1_2_1) {
  $__num = \func_num_args();
  $__res = $alternativeListT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_List_Trans_monadPlusListT'] = __NAMESPACE__ . '\\majControl_majMonad_majList_majTrans_monadmajPlusmajListmajT';

