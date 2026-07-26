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


// Control_Monad_List_Trans_Yield
$GLOBALS['Control_Monad_List_Trans_Yield'] = (function() {
  $__fn = function($value0 = null, $value1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("Yield", $value0, $value1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Control_Monad_List_Trans_Skip
$GLOBALS['Control_Monad_List_Trans_Skip'] = function($value0 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data1("Skip", $value0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_List_Trans_Done
$GLOBALS['Control_Monad_List_Trans_Done'] = ($GLOBALS['__phpurs_data0_Done'] ??= new Phpurs_Data0("Done"));

// Control_Monad_List_Trans_ListT
$GLOBALS['Control_Monad_List_Trans_ListT'] = function($x_0 = null) {
  $__num = \func_num_args();
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_List_Trans_wrapLazy
$GLOBALS['Control_Monad_List_Trans_wrapLazy'] = (function() {
  $__fn = function($dictApplicative_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($dictApplicative_0)['pure'])(new Phpurs_Data1("Skip", $v_1));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Control_Monad_List_Trans_wrapEffect
$GLOBALS['Control_Monad_List_Trans_wrapEffect'] = (function() {
  $__fn = function($dictFunctor_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictFunctor_0)['map'])(((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])($GLOBALS['Control_Monad_List_Trans_Skip']))(((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])($GLOBALS['Data_Lazy_defer']))($GLOBALS['Data_Function_const']))))($v_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Control_Monad_List_Trans_unfold
$GLOBALS['Control_Monad_List_Trans_unfold'] = (function() {
  $__fn = function($dictMonad_0 = null, $f_1 = null, $z_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((((((($dictMonad_0)['Bind1'])(null))['Apply0'])(null))['Functor0'])(null))['map'])(function($v_3 = null) use ($dictMonad_0, $f_1) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object($v_3) && (($v_3)->{'tag'} === "Just"))) {
$__local_var_4_1 = (($v_3)->{'value0'})->{'value0'};
$__t0 = new Phpurs_Data2("Yield", (($v_3)->{'value0'})->{'value1'}, ($GLOBALS['Data_Lazy_defer'])(function($v1_5 = null) use ($__local_var_4_1, $dictMonad_0, $f_1) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Monad_List_Trans_unfold'])($dictMonad_0))($f_1))($__local_var_4_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
goto end_branch_0;;
};
  if ((is_object($v_3) && (($v_3)->{'tag'} === "Nothing"))) {
$__t0 = new Phpurs_Data0("Done");
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
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Control_Monad_List_Trans_uncons
$GLOBALS['Control_Monad_List_Trans_uncons'] = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($dictMonad_0)['Applicative0'])(null);
  $__res = function($v_2 = null) use ($__local_var_1_0, $dictMonad_0) {
  $__num = \func_num_args();
  $__res = ((((($dictMonad_0)['Bind1'])(null))['bind'])($v_2))(function($v1_3 = null) use ($__local_var_1_0, $dictMonad_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((is_object($v1_3) && (($v1_3)->{'tag'} === "Yield"))) {
$__t1 = (($__local_var_1_0)['pure'])(new Phpurs_Data1("Just", new Phpurs_Data2("Tuple", ($v1_3)->{'value0'}, ($GLOBALS['Data_Lazy_force'])(($v1_3)->{'value1'}))));
goto end_branch_1;;
};
  if ((is_object($v1_3) && (($v1_3)->{'tag'} === "Skip"))) {
$__t1 = (($GLOBALS['Control_Monad_List_Trans_uncons'])($dictMonad_0))(($GLOBALS['Data_Lazy_force'])(($v1_3)->{'value0'}));
goto end_branch_1;;
};
  if ((is_object($v1_3) && (($v1_3)->{'tag'} === "Done"))) {
$__t1 = (($__local_var_1_0)['pure'])(new Phpurs_Data0("Nothing"));
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_List_Trans_tail
$GLOBALS['Control_Monad_List_Trans_tail'] = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $uncons1_1_0 = ($GLOBALS['Control_Monad_List_Trans_uncons'])($dictMonad_0);
  $__res = function($l_2 = null) use ($dictMonad_0, $uncons1_1_0) {
  $__num = \func_num_args();
  $__res = ((((((((($dictMonad_0)['Bind1'])(null))['Apply0'])(null))['Functor0'])(null))['map'])((($GLOBALS['Data_Maybe_functorMaybe'])['map'])($GLOBALS['Data_Tuple_snd'])))(($uncons1_1_0)($l_2));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_List_Trans_takeWhile
$GLOBALS['Control_Monad_List_Trans_takeWhile'] = function($dictApplicative_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (((($dictApplicative_0)['Apply0'])(null))['Functor0'])(null);
  $__res = (function() use ($__local_var_1_0, $dictApplicative_0) {
  $__fn = function($f_2 = null, $v_3 = null) use ($__local_var_1_0, $dictApplicative_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_1_0)['map'])(function($v_4 = null) use ($dictApplicative_0, $f_2) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((is_object($v_4) && (($v_4)->{'tag'} === "Yield"))) {
$__t2 = null;;
if (($f_2)(($v_4)->{'value0'})) {
$__t2 = new Phpurs_Data2("Yield", ($v_4)->{'value0'}, ((($GLOBALS['Data_Lazy_functorLazy'])['map'])((($GLOBALS['Control_Monad_List_Trans_takeWhile'])($dictApplicative_0))($f_2)))(($v_4)->{'value1'}));
goto end_branch_2;;
};
$__t2 = new Phpurs_Data0("Done");
end_branch_2:;
$__t1 = $__t2;
goto end_branch_1;;
};
  if ((is_object($v_4) && (($v_4)->{'tag'} === "Skip"))) {
$__t1 = new Phpurs_Data1("Skip", ((($GLOBALS['Data_Lazy_functorLazy'])['map'])((($GLOBALS['Control_Monad_List_Trans_takeWhile'])($dictApplicative_0))($f_2)))(($v_4)->{'value0'}));
goto end_branch_1;;
};
  if ((is_object($v_4) && (($v_4)->{'tag'} === "Done"))) {
$__t1 = new Phpurs_Data0("Done");
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_List_Trans_scanl
$GLOBALS['Control_Monad_List_Trans_scanl'] = (function() {
  $__fn = function($dictMonad_0 = null, $f_1 = null, $b_2 = null, $l_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = ((($GLOBALS['Control_Monad_List_Trans_unfold'])($dictMonad_0))(function($v_4 = null) use ($dictMonad_0, $f_1) {
  $__num = \func_num_args();
  $__local_var_5_0 = ($v_4)->{'value0'};
  $__res = ((((((((($dictMonad_0)['Bind1'])(null))['Apply0'])(null))['Functor0'])(null))['map'])(function($v1_6 = null) use ($__local_var_5_0, $f_1) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((is_object($v1_6) && (($v1_6)->{'tag'} === "Yield"))) {
$__t1 = new Phpurs_Data1("Just", new Phpurs_Data2("Tuple", new Phpurs_Data2("Tuple", (($f_1)($__local_var_5_0))(($v1_6)->{'value0'}), ($GLOBALS['Data_Lazy_force'])(($v1_6)->{'value1'})), $__local_var_5_0));
goto end_branch_1;;
};
  if ((is_object($v1_6) && (($v1_6)->{'tag'} === "Skip"))) {
$__t1 = new Phpurs_Data1("Just", new Phpurs_Data2("Tuple", new Phpurs_Data2("Tuple", $__local_var_5_0, ($GLOBALS['Data_Lazy_force'])(($v1_6)->{'value0'})), $__local_var_5_0));
goto end_branch_1;;
};
  if ((is_object($v1_6) && (($v1_6)->{'tag'} === "Done"))) {
$__t1 = new Phpurs_Data0("Nothing");
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
}))(new Phpurs_Data2("Tuple", $b_2, $l_3));
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})();

// Control_Monad_List_Trans_prepend'
$GLOBALS['Control_Monad_List_Trans_prepend__prime__'] = (function() {
  $__fn = function($dictApplicative_0 = null, $h_1 = null, $t_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($dictApplicative_0)['pure'])(new Phpurs_Data2("Yield", $h_1, $t_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Control_Monad_List_Trans_prepend
$GLOBALS['Control_Monad_List_Trans_prepend'] = (function() {
  $__fn = function($dictApplicative_0 = null, $h_1 = null, $t_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($dictApplicative_0)['pure'])(new Phpurs_Data2("Yield", $h_1, ($GLOBALS['Data_Lazy_defer'])(function($v_3 = null) use ($t_2) {
  $__num = \func_num_args();
  $__res = $t_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Control_Monad_List_Trans_nil
$GLOBALS['Control_Monad_List_Trans_nil'] = function($dictApplicative_0 = null) {
  $__num = \func_num_args();
  $__res = (($dictApplicative_0)['pure'])(new Phpurs_Data0("Done"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_List_Trans_singleton
$GLOBALS['Control_Monad_List_Trans_singleton'] = function($dictApplicative_0 = null) {
  $__num = \func_num_args();
  $nil1_1_0 = (($dictApplicative_0)['pure'])(new Phpurs_Data0("Done"));
  $__res = function($a_2 = null) use ($dictApplicative_0, $nil1_1_0) {
  $__num = \func_num_args();
  $__res = (($dictApplicative_0)['pure'])(new Phpurs_Data2("Yield", $a_2, ($GLOBALS['Data_Lazy_defer'])(function($v_3 = null) use ($nil1_1_0) {
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_List_Trans_take
$GLOBALS['Control_Monad_List_Trans_take'] = function($dictApplicative_0 = null) {
  $__num = \func_num_args();
  $nil1_1_0 = (($dictApplicative_0)['pure'])(new Phpurs_Data0("Done"));
  $__local_var_2_1 = (((($dictApplicative_0)['Apply0'])(null))['Functor0'])(null);
  $__res = (function() use ($__local_var_2_1, $dictApplicative_0, $nil1_1_0) {
  $__fn = function($v_3 = null, $v1_4 = null) use ($__local_var_2_1, $dictApplicative_0, $nil1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = match ($v_3) { 0 => $nil1_1_0, default => ((($__local_var_2_1)['map'])(function($v2_5 = null) use ($dictApplicative_0, $v_3) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ((is_object($v2_5) && (($v2_5)->{'tag'} === "Yield"))) {
$__t2 = new Phpurs_Data2("Yield", ($v2_5)->{'value0'}, ((($GLOBALS['Data_Lazy_functorLazy'])['map'])((($GLOBALS['Control_Monad_List_Trans_take'])($dictApplicative_0))(((($GLOBALS['Data_Ring_ringInt'])['sub'])($v_3))(1))))(($v2_5)->{'value1'}));
goto end_branch_2;;
};
  if ((is_object($v2_5) && (($v2_5)->{'tag'} === "Skip"))) {
$__t2 = new Phpurs_Data1("Skip", ((($GLOBALS['Data_Lazy_functorLazy'])['map'])((($GLOBALS['Control_Monad_List_Trans_take'])($dictApplicative_0))($v_3)))(($v2_5)->{'value0'}));
goto end_branch_2;;
};
  if ((is_object($v2_5) && (($v2_5)->{'tag'} === "Done"))) {
$__t2 = new Phpurs_Data0("Done");
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_List_Trans_zipWith'
$GLOBALS['Control_Monad_List_Trans_zipWith__prime__'] = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $Applicative0_1_0 = (($dictMonad_0)['Applicative0'])(null);
  $nil1_2_1 = (($Applicative0_1_0)['pure'])(new Phpurs_Data0("Done"));
  $Bind1_3_2 = (($dictMonad_0)['Bind1'])(null);
  $Functor0_4_3 = (((($Bind1_3_2)['Apply0'])(null))['Functor0'])(null);
  $uncons1_5_4 = ($GLOBALS['Control_Monad_List_Trans_uncons'])($dictMonad_0);
  $__res = (function() use ($Applicative0_1_0, $Bind1_3_2, $Functor0_4_3, $dictMonad_0, $nil1_2_1, $uncons1_5_4) {
  $__fn = function($f_6 = null, $fa_7 = null, $fb_8 = null) use ($Applicative0_1_0, $Bind1_3_2, $Functor0_4_3, $dictMonad_0, $nil1_2_1, $uncons1_5_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($GLOBALS['Control_Monad_List_Trans_wrapEffect'])($Functor0_4_3))(((($Bind1_3_2)['bind'])(($uncons1_5_4)($fa_7)))(function($ua_9 = null) use ($Applicative0_1_0, $Bind1_3_2, $Functor0_4_3, $dictMonad_0, $f_6, $fb_8, $nil1_2_1, $uncons1_5_4) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_2)['bind'])(($uncons1_5_4)($fb_8)))(function($ub_10 = null) use ($Applicative0_1_0, $Functor0_4_3, $dictMonad_0, $f_6, $nil1_2_1, $ua_9) {
  $__num = \func_num_args();
  $__t5 = null;;
  if ((is_object($ub_10) && (($ub_10)->{'tag'} === "Nothing"))) {
$__t5 = (($Applicative0_1_0)['pure'])($nil1_2_1);
goto end_branch_5;;
};
  if ((is_object($ua_9) && (($ua_9)->{'tag'} === "Nothing"))) {
$__t5 = (($Applicative0_1_0)['pure'])($nil1_2_1);
goto end_branch_5;;
};
  if (((is_object($ua_9) && (($ua_9)->{'tag'} === "Just")) && (is_object($ub_10) && (($ub_10)->{'tag'} === "Just")))) {
$__local_var_11_6 = (($ua_9)->{'value0'})->{'value1'};
$__local_var_12_7 = (($ub_10)->{'value0'})->{'value1'};
$__local_var_13_8 = ($GLOBALS['Data_Lazy_defer'])(function($v2_13 = null) use ($__local_var_11_6, $__local_var_12_7, $dictMonad_0, $f_6) {
  $__num = \func_num_args();
  $__res = (((($GLOBALS['Control_Monad_List_Trans_zipWith__prime__'])($dictMonad_0))($f_6))($__local_var_11_6))($__local_var_12_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
$__t5 = ((($Functor0_4_3)['map'])(function($a_14 = null) use ($Applicative0_1_0, $__local_var_13_8) {
  $__num = \func_num_args();
  $__res = (($Applicative0_1_0)['pure'])(new Phpurs_Data2("Yield", $a_14, $__local_var_13_8));
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_List_Trans_zipWith
$GLOBALS['Control_Monad_List_Trans_zipWith'] = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $zipWith_prime1_1_0 = ($GLOBALS['Control_Monad_List_Trans_zipWith__prime__'])($dictMonad_0);
  $__res = function($f_2 = null) use ($dictMonad_0, $zipWith_prime1_1_0) {
  $__num = \func_num_args();
  $__res = ($zipWith_prime1_1_0)((function() use ($dictMonad_0, $f_2) {
  $__fn = function($a_3 = null, $b_4 = null) use ($dictMonad_0, $f_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($dictMonad_0)['Applicative0'])(null))['pure'])((($f_2)($a_3))($b_4));
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_List_Trans_newtypeListT
$GLOBALS['Control_Monad_List_Trans_newtypeListT'] = ["Coercible0" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Control_Monad_List_Trans_mapMaybe
$GLOBALS['Control_Monad_List_Trans_mapMaybe'] = (function() {
  $__fn = function($dictFunctor_0 = null, $f_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictFunctor_0)['map'])(function($v_3 = null) use ($dictFunctor_0, $f_1) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object($v_3) && (($v_3)->{'tag'} === "Yield"))) {
$__local_var_4_1 = ((($GLOBALS['Data_Maybe_functorMaybe'])['map'])($GLOBALS['Control_Monad_List_Trans_Yield']))(($f_1)(($v_3)->{'value0'}));
$__t2 = null;;
if ((is_object($__local_var_4_1) && (($__local_var_4_1)->{'tag'} === "Nothing"))) {
$__t2 = $GLOBALS['Control_Monad_List_Trans_Skip'];
goto end_branch_2;;
};
if ((is_object($__local_var_4_1) && (($__local_var_4_1)->{'tag'} === "Just"))) {
$__t2 = (($GLOBALS['Control_Category_categoryFn'])['identity'])(($__local_var_4_1)->{'value0'});
goto end_branch_2;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
end_branch_2:;
$__t0 = ($__t2)(((($GLOBALS['Data_Lazy_functorLazy'])['map'])((($GLOBALS['Control_Monad_List_Trans_mapMaybe'])($dictFunctor_0))($f_1)))(($v_3)->{'value1'}));
goto end_branch_0;;
};
  if ((is_object($v_3) && (($v_3)->{'tag'} === "Skip"))) {
$__t0 = new Phpurs_Data1("Skip", ((($GLOBALS['Data_Lazy_functorLazy'])['map'])((($GLOBALS['Control_Monad_List_Trans_mapMaybe'])($dictFunctor_0))($f_1)))(($v_3)->{'value0'}));
goto end_branch_0;;
};
  if ((is_object($v_3) && (($v_3)->{'tag'} === "Done"))) {
$__t0 = new Phpurs_Data0("Done");
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
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Control_Monad_List_Trans_iterate
$GLOBALS['Control_Monad_List_Trans_iterate'] = (function() {
  $__fn = function($dictMonad_0 = null, $f_1 = null, $a_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($GLOBALS['Control_Monad_List_Trans_unfold'])($dictMonad_0))(function($x_3 = null) use ($dictMonad_0, $f_1) {
  $__num = \func_num_args();
  $__res = (((($dictMonad_0)['Applicative0'])(null))['pure'])(new Phpurs_Data1("Just", new Phpurs_Data2("Tuple", ($f_1)($x_3), $x_3)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($a_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Control_Monad_List_Trans_repeat
$GLOBALS['Control_Monad_List_Trans_repeat'] = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_List_Trans_iterate'])($dictMonad_0))(($GLOBALS['Control_Category_categoryFn'])['identity']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_List_Trans_head
$GLOBALS['Control_Monad_List_Trans_head'] = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $uncons1_1_0 = ($GLOBALS['Control_Monad_List_Trans_uncons'])($dictMonad_0);
  $__res = function($l_2 = null) use ($dictMonad_0, $uncons1_1_0) {
  $__num = \func_num_args();
  $__res = ((((((((($dictMonad_0)['Bind1'])(null))['Apply0'])(null))['Functor0'])(null))['map'])((($GLOBALS['Data_Maybe_functorMaybe'])['map'])($GLOBALS['Data_Tuple_fst'])))(($uncons1_1_0)($l_2));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_List_Trans_functorListT
$GLOBALS['Control_Monad_List_Trans_functorListT'] = function($dictFunctor_0 = null) {
  $__num = \func_num_args();
  $__res = ["map" => (function() use ($dictFunctor_0) {
  $__fn = function($f_1 = null, $v_2 = null) use ($dictFunctor_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictFunctor_0)['map'])(function($v_3 = null) use ($dictFunctor_0, $f_1) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object($v_3) && (($v_3)->{'tag'} === "Yield"))) {
$__t0 = new Phpurs_Data2("Yield", ($f_1)(($v_3)->{'value0'}), ((($GLOBALS['Data_Lazy_functorLazy'])['map'])(((($GLOBALS['Control_Monad_List_Trans_functorListT'])($dictFunctor_0))['map'])($f_1)))(($v_3)->{'value1'}));
goto end_branch_0;;
};
  if ((is_object($v_3) && (($v_3)->{'tag'} === "Skip"))) {
$__t0 = new Phpurs_Data1("Skip", ((($GLOBALS['Data_Lazy_functorLazy'])['map'])(((($GLOBALS['Control_Monad_List_Trans_functorListT'])($dictFunctor_0))['map'])($f_1)))(($v_3)->{'value0'}));
goto end_branch_0;;
};
  if ((is_object($v_3) && (($v_3)->{'tag'} === "Done"))) {
$__t0 = new Phpurs_Data0("Done");
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_List_Trans_fromEffect
$GLOBALS['Control_Monad_List_Trans_fromEffect'] = function($dictApplicative_0 = null) {
  $__num = \func_num_args();
  $nil1_1_0 = (($dictApplicative_0)['pure'])(new Phpurs_Data0("Done"));
  $__res = function($fa_2 = null) use ($dictApplicative_0, $nil1_1_0) {
  $__num = \func_num_args();
  $__local_var_3_1 = ($GLOBALS['Data_Lazy_defer'])(function($v_3 = null) use ($nil1_1_0) {
  $__num = \func_num_args();
  $__res = $nil1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = ((((((($dictApplicative_0)['Apply0'])(null))['Functor0'])(null))['map'])(function($a_4 = null) use ($__local_var_3_1) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Yield", $a_4, $__local_var_3_1);
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_List_Trans_monadTransListT
$GLOBALS['Control_Monad_List_Trans_monadTransListT'] = ["lift" => function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_List_Trans_fromEffect'])((($dictMonad_0)['Applicative0'])(null));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Control_Monad_List_Trans_foldlRec'
$GLOBALS['Control_Monad_List_Trans_foldlRec__prime__'] = function($dictMonadRec_0 = null) {
  $__num = \func_num_args();
  $Monad0_1_0 = (($dictMonadRec_0)['Monad0'])(null);
  $__local_var_2_1 = (($Monad0_1_0)['Applicative0'])(null);
  $__local_var_3_2 = (($Monad0_1_0)['Bind1'])(null);
  $uncons1_4_3 = ($GLOBALS['Control_Monad_List_Trans_uncons'])($Monad0_1_0);
  $__res = (function() use ($__local_var_2_1, $__local_var_3_2, $dictMonadRec_0, $uncons1_4_3) {
  $__fn = function($f_5 = null, $a_6 = null, $b_7 = null) use ($__local_var_2_1, $__local_var_3_2, $dictMonadRec_0, $uncons1_4_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictMonadRec_0)['tailRecM'])(function($o_8 = null) use ($__local_var_2_1, $__local_var_3_2, $f_5, $uncons1_4_3) {
  $__num = \func_num_args();
  $__local_var_9_4 = ($o_8)['a'];
  $__res = ((($__local_var_3_2)['bind'])(($uncons1_4_3)(($o_8)['b'])))(function($v_10 = null) use ($__local_var_2_1, $__local_var_3_2, $__local_var_9_4, $f_5) {
  $__num = \func_num_args();
  $__t5 = null;;
  if ((is_object($v_10) && (($v_10)->{'tag'} === "Nothing"))) {
$__t5 = (($__local_var_2_1)['pure'])(new Phpurs_Data1("Done", $__local_var_9_4));
goto end_branch_5;;
};
  if ((is_object($v_10) && (($v_10)->{'tag'} === "Just"))) {
$__local_var_11_6 = (($v_10)->{'value0'})->{'value1'};
$__t5 = ((($__local_var_3_2)['bind'])((($f_5)($__local_var_9_4))((($v_10)->{'value0'})->{'value0'})))(function($b_prime_12 = null) use ($__local_var_11_6, $__local_var_2_1) {
  $__num = \func_num_args();
  $__res = (($__local_var_2_1)['pure'])(new Phpurs_Data1("Loop", ["a" => $b_prime_12, "b" => $__local_var_11_6]));
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
}))(["a" => $a_6, "b" => $b_7]);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_List_Trans_runListTRec
$GLOBALS['Control_Monad_List_Trans_runListTRec'] = function($dictMonadRec_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Monad_List_Trans_foldlRec__prime__'])($dictMonadRec_0))((function() use ($dictMonadRec_0) {
  $__fn = function($v_1 = null, $v1_2 = null) use ($dictMonadRec_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((((($dictMonadRec_0)['Monad0'])(null))['Applicative0'])(null))['pure'])($GLOBALS['Data_Unit_unit']);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($GLOBALS['Data_Unit_unit']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_List_Trans_foldlRec
$GLOBALS['Control_Monad_List_Trans_foldlRec'] = function($dictMonadRec_0 = null) {
  $__num = \func_num_args();
  $Monad0_1_0 = (($dictMonadRec_0)['Monad0'])(null);
  $__local_var_2_1 = (($Monad0_1_0)['Applicative0'])(null);
  $uncons1_3_2 = ($GLOBALS['Control_Monad_List_Trans_uncons'])($Monad0_1_0);
  $__res = (function() use ($Monad0_1_0, $__local_var_2_1, $dictMonadRec_0, $uncons1_3_2) {
  $__fn = function($f_4 = null, $a_5 = null, $b_6 = null) use ($Monad0_1_0, $__local_var_2_1, $dictMonadRec_0, $uncons1_3_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictMonadRec_0)['tailRecM'])(function($o_7 = null) use ($Monad0_1_0, $__local_var_2_1, $f_4, $uncons1_3_2) {
  $__num = \func_num_args();
  $__local_var_8_3 = ($o_7)['a'];
  $__res = ((((($Monad0_1_0)['Bind1'])(null))['bind'])(($uncons1_3_2)(($o_7)['b'])))(function($v_9 = null) use ($__local_var_2_1, $__local_var_8_3, $f_4) {
  $__num = \func_num_args();
  $__t4 = null;;
  if ((is_object($v_9) && (($v_9)->{'tag'} === "Nothing"))) {
$__t4 = (($__local_var_2_1)['pure'])(new Phpurs_Data1("Done", $__local_var_8_3));
goto end_branch_4;;
};
  if ((is_object($v_9) && (($v_9)->{'tag'} === "Just"))) {
$__t4 = (($__local_var_2_1)['pure'])(new Phpurs_Data1("Loop", ["a" => (($f_4)($__local_var_8_3))((($v_9)->{'value0'})->{'value0'}), "b" => (($v_9)->{'value0'})->{'value1'}]));
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
}))(["a" => $a_5, "b" => $b_6]);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_List_Trans_foldl'
$GLOBALS['Control_Monad_List_Trans_foldl__prime__'] = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($dictMonad_0)['Bind1'])(null);
  $uncons1_2_1 = ($GLOBALS['Control_Monad_List_Trans_uncons'])($dictMonad_0);
  $__res = function($f_3 = null) use ($__local_var_1_0, $dictMonad_0, $uncons1_2_1) {
  $__num = \func_num_args();
  $loop_4_2 = null;
  $loop_4_2 = (function() use ($__local_var_1_0, $dictMonad_0, $f_3, &$loop_4_2, $uncons1_2_1) {
  $__fn = function($b_5 = null, $l_6 = null) use ($__local_var_1_0, $dictMonad_0, $f_3, &$loop_4_2, $uncons1_2_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_1_0)['bind'])(($uncons1_2_1)($l_6)))(function($v_7 = null) use ($__local_var_1_0, $b_5, $dictMonad_0, $f_3, &$loop_4_2) {
  $__num = \func_num_args();
  $__t3 = null;;
  if ((is_object($v_7) && (($v_7)->{'tag'} === "Nothing"))) {
$__t3 = (((($dictMonad_0)['Applicative0'])(null))['pure'])($b_5);
goto end_branch_3;;
};
  if ((is_object($v_7) && (($v_7)->{'tag'} === "Just"))) {
$__local_var_8_4 = (($v_7)->{'value0'})->{'value1'};
$__t3 = ((($__local_var_1_0)['bind'])((($f_3)($b_5))((($v_7)->{'value0'})->{'value0'})))(function($a_9 = null) use ($__local_var_8_4, &$loop_4_2) {
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_List_Trans_runListT
$GLOBALS['Control_Monad_List_Trans_runListT'] = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Monad_List_Trans_foldl__prime__'])($dictMonad_0))((function() use ($dictMonad_0) {
  $__fn = function($v_1 = null, $v1_2 = null) use ($dictMonad_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($dictMonad_0)['Applicative0'])(null))['pure'])($GLOBALS['Data_Unit_unit']);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($GLOBALS['Data_Unit_unit']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_List_Trans_foldl
$GLOBALS['Control_Monad_List_Trans_foldl'] = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $uncons1_1_0 = ($GLOBALS['Control_Monad_List_Trans_uncons'])($dictMonad_0);
  $__res = function($f_2 = null) use ($dictMonad_0, $uncons1_1_0) {
  $__num = \func_num_args();
  $loop_3_1 = null;
  $loop_3_1 = (function() use ($dictMonad_0, $f_2, &$loop_3_1, $uncons1_1_0) {
  $__fn = function($b_4 = null, $l_5 = null) use ($dictMonad_0, $f_2, &$loop_3_1, $uncons1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((($dictMonad_0)['Bind1'])(null))['bind'])(($uncons1_1_0)($l_5)))(function($v_6 = null) use ($b_4, $dictMonad_0, $f_2, &$loop_3_1) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ((is_object($v_6) && (($v_6)->{'tag'} === "Nothing"))) {
$__t2 = (((($dictMonad_0)['Applicative0'])(null))['pure'])($b_4);
goto end_branch_2;;
};
  if ((is_object($v_6) && (($v_6)->{'tag'} === "Just"))) {
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_List_Trans_filter
$GLOBALS['Control_Monad_List_Trans_filter'] = (function() {
  $__fn = function($dictFunctor_0 = null, $f_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictFunctor_0)['map'])(function($v_3 = null) use ($dictFunctor_0, $f_1) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object($v_3) && (($v_3)->{'tag'} === "Yield"))) {
$s_prime_4_1 = ((($GLOBALS['Data_Lazy_functorLazy'])['map'])((($GLOBALS['Control_Monad_List_Trans_filter'])($dictFunctor_0))($f_1)))(($v_3)->{'value1'});
$__t2 = null;;
if (($f_1)(($v_3)->{'value0'})) {
$__t2 = new Phpurs_Data2("Yield", ($v_3)->{'value0'}, $s_prime_4_1);
goto end_branch_2;;
};
$__t2 = new Phpurs_Data1("Skip", $s_prime_4_1);
end_branch_2:;
$__t0 = $__t2;
goto end_branch_0;;
};
  if ((is_object($v_3) && (($v_3)->{'tag'} === "Skip"))) {
$__t0 = new Phpurs_Data1("Skip", ((($GLOBALS['Data_Lazy_functorLazy'])['map'])((($GLOBALS['Control_Monad_List_Trans_filter'])($dictFunctor_0))($f_1)))(($v_3)->{'value0'}));
goto end_branch_0;;
};
  if ((is_object($v_3) && (($v_3)->{'tag'} === "Done"))) {
$__t0 = new Phpurs_Data0("Done");
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
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Control_Monad_List_Trans_dropWhile
$GLOBALS['Control_Monad_List_Trans_dropWhile'] = function($dictApplicative_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (((($dictApplicative_0)['Apply0'])(null))['Functor0'])(null);
  $__res = (function() use ($__local_var_1_0, $dictApplicative_0) {
  $__fn = function($f_2 = null, $v_3 = null) use ($__local_var_1_0, $dictApplicative_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_1_0)['map'])(function($v_4 = null) use ($dictApplicative_0, $f_2) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((is_object($v_4) && (($v_4)->{'tag'} === "Yield"))) {
$__t2 = null;;
if (($f_2)(($v_4)->{'value0'})) {
$__t2 = new Phpurs_Data1("Skip", ((($GLOBALS['Data_Lazy_functorLazy'])['map'])((($GLOBALS['Control_Monad_List_Trans_dropWhile'])($dictApplicative_0))($f_2)))(($v_4)->{'value1'}));
goto end_branch_2;;
};
$__t2 = new Phpurs_Data2("Yield", ($v_4)->{'value0'}, ($v_4)->{'value1'});
end_branch_2:;
$__t1 = $__t2;
goto end_branch_1;;
};
  if ((is_object($v_4) && (($v_4)->{'tag'} === "Skip"))) {
$__t1 = new Phpurs_Data1("Skip", ((($GLOBALS['Data_Lazy_functorLazy'])['map'])((($GLOBALS['Control_Monad_List_Trans_dropWhile'])($dictApplicative_0))($f_2)))(($v_4)->{'value0'}));
goto end_branch_1;;
};
  if ((is_object($v_4) && (($v_4)->{'tag'} === "Done"))) {
$__t1 = new Phpurs_Data0("Done");
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_List_Trans_drop
$GLOBALS['Control_Monad_List_Trans_drop'] = function($dictApplicative_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (((($dictApplicative_0)['Apply0'])(null))['Functor0'])(null);
  $__res = (function() use ($__local_var_1_0, $dictApplicative_0) {
  $__fn = function($v_2 = null, $v1_3 = null) use ($__local_var_1_0, $dictApplicative_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = match ($v_2) { 0 => $v1_3, default => ((($__local_var_1_0)['map'])(function($v2_4 = null) use ($dictApplicative_0, $v_2) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((is_object($v2_4) && (($v2_4)->{'tag'} === "Yield"))) {
$__t1 = new Phpurs_Data1("Skip", ((($GLOBALS['Data_Lazy_functorLazy'])['map'])((($GLOBALS['Control_Monad_List_Trans_drop'])($dictApplicative_0))(((($GLOBALS['Data_Ring_ringInt'])['sub'])($v_2))(1))))(($v2_4)->{'value1'}));
goto end_branch_1;;
};
  if ((is_object($v2_4) && (($v2_4)->{'tag'} === "Skip"))) {
$__t1 = new Phpurs_Data1("Skip", ((($GLOBALS['Data_Lazy_functorLazy'])['map'])((($GLOBALS['Control_Monad_List_Trans_drop'])($dictApplicative_0))($v_2)))(($v2_4)->{'value0'}));
goto end_branch_1;;
};
  if ((is_object($v2_4) && (($v2_4)->{'tag'} === "Done"))) {
$__t1 = new Phpurs_Data0("Done");
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_List_Trans_cons
$GLOBALS['Control_Monad_List_Trans_cons'] = (function() {
  $__fn = function($dictApplicative_0 = null, $lh_1 = null, $t_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($dictApplicative_0)['pure'])(new Phpurs_Data2("Yield", ($GLOBALS['Data_Lazy_force'])($lh_1), $t_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Control_Monad_List_Trans_unfoldable1ListT
$GLOBALS['Control_Monad_List_Trans_unfoldable1ListT'] = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $Applicative0_1_0 = (($dictMonad_0)['Applicative0'])(null);
  $singleton1_2_1 = ($GLOBALS['Control_Monad_List_Trans_singleton'])($Applicative0_1_0);
  $__res = ["unfoldr1" => (function() use ($Applicative0_1_0, $singleton1_2_1) {
  $__fn = function($f_3 = null, $b_4 = null) use ($Applicative0_1_0, $singleton1_2_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__5_2 = null;
  $go__5_2 = function($v_6 = null) use ($Applicative0_1_0, $f_3, &$go__5_2, $singleton1_2_1) {
  $__num = \func_num_args();
  $__t3 = null;;
  if ((is_object(($v_6)->{'value1'}) && ((($v_6)->{'value1'})->{'tag'} === "Nothing"))) {
$__t3 = ($singleton1_2_1)(($v_6)->{'value0'});
goto end_branch_3;;
};
  if ((is_object(($v_6)->{'value1'}) && ((($v_6)->{'value1'})->{'tag'} === "Just"))) {
$__local_var_7_4 = (($v_6)->{'value1'})->{'value0'};
$__t3 = (($Applicative0_1_0)['pure'])(new Phpurs_Data2("Yield", ($GLOBALS['Data_Lazy_force'])((($GLOBALS['Data_Lazy_applicativeLazy'])['pure'])(($v_6)->{'value0'})), ($GLOBALS['Data_Lazy_defer'])(function($v1_8 = null) use ($__local_var_7_4, $f_3, &$go__5_2) {
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_List_Trans_unfoldableListT
$GLOBALS['Control_Monad_List_Trans_unfoldableListT'] = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $Applicative0_1_0 = (($dictMonad_0)['Applicative0'])(null);
  $nil1_2_1 = (($Applicative0_1_0)['pure'])(new Phpurs_Data0("Done"));
  $unfoldable1ListT1_3_2 = ($GLOBALS['Control_Monad_List_Trans_unfoldable1ListT'])($dictMonad_0);
  $__res = ["unfoldr" => (function() use ($Applicative0_1_0, $nil1_2_1) {
  $__fn = function($f_4 = null, $b_5 = null) use ($Applicative0_1_0, $nil1_2_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__6_3 = null;
  $go__6_3 = function($v_7 = null) use ($Applicative0_1_0, $f_4, &$go__6_3, $nil1_2_1) {
  $__num = \func_num_args();
  $__t4 = null;;
  if ((is_object($v_7) && (($v_7)->{'tag'} === "Nothing"))) {
$__t4 = $nil1_2_1;
goto end_branch_4;;
};
  if ((is_object($v_7) && (($v_7)->{'tag'} === "Just"))) {
$__local_var_8_5 = (($v_7)->{'value0'})->{'value1'};
$__t4 = (($Applicative0_1_0)['pure'])(new Phpurs_Data2("Yield", ($GLOBALS['Data_Lazy_force'])((($GLOBALS['Data_Lazy_applicativeLazy'])['pure'])((($v_7)->{'value0'})->{'value0'})), ($GLOBALS['Data_Lazy_defer'])(function($v1_9 = null) use ($__local_var_8_5, $f_4, &$go__6_3) {
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
})(), "Unfoldable10" => function($_dollar__unused_4 = null) use ($unfoldable1ListT1_3_2) {
  $__num = \func_num_args();
  $__res = $unfoldable1ListT1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_List_Trans_semigroupListT
$GLOBALS['Control_Monad_List_Trans_semigroupListT'] = function($dictApplicative_0 = null) {
  $__num = \func_num_args();
  $__res = ["append" => ($GLOBALS['Control_Monad_List_Trans_concat'])($dictApplicative_0)];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_List_Trans_concat
$GLOBALS['Control_Monad_List_Trans_concat'] = function($dictApplicative_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (((($dictApplicative_0)['Apply0'])(null))['Functor0'])(null);
  $__res = (function() use ($__local_var_1_0, $dictApplicative_0) {
  $__fn = function($x_2 = null, $y_3 = null) use ($__local_var_1_0, $dictApplicative_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_1_0)['map'])(function($v_4 = null) use ($dictApplicative_0, $y_3) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((is_object($v_4) && (($v_4)->{'tag'} === "Yield"))) {
$__t1 = new Phpurs_Data2("Yield", ($v_4)->{'value0'}, ((($GLOBALS['Data_Lazy_functorLazy'])['map'])(function($v1_5 = null) use ($dictApplicative_0, $y_3) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Monad_List_Trans_concat'])($dictApplicative_0))($v1_5))($y_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_4)->{'value1'}));
goto end_branch_1;;
};
  if ((is_object($v_4) && (($v_4)->{'tag'} === "Skip"))) {
$__t1 = new Phpurs_Data1("Skip", ((($GLOBALS['Data_Lazy_functorLazy'])['map'])(function($v1_5 = null) use ($dictApplicative_0, $y_3) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Monad_List_Trans_concat'])($dictApplicative_0))($v1_5))($y_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_4)->{'value0'}));
goto end_branch_1;;
};
  if ((is_object($v_4) && (($v_4)->{'tag'} === "Done"))) {
$__t1 = new Phpurs_Data1("Skip", ($GLOBALS['Data_Lazy_defer'])(function($v_5 = null) use ($y_3) {
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_List_Trans_monoidListT
$GLOBALS['Control_Monad_List_Trans_monoidListT'] = function($dictApplicative_0 = null) {
  $__num = \func_num_args();
  $semigroupListT1_1_0 = ["append" => ($GLOBALS['Control_Monad_List_Trans_concat'])($dictApplicative_0)];
  $__res = ["mempty" => (($dictApplicative_0)['pure'])(new Phpurs_Data0("Done")), "Semigroup0" => function($_dollar__unused_2 = null) use ($semigroupListT1_1_0) {
  $__num = \func_num_args();
  $__res = $semigroupListT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_List_Trans_catMaybes
$GLOBALS['Control_Monad_List_Trans_catMaybes'] = function($dictFunctor_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_List_Trans_mapMaybe'])($dictFunctor_0))(($GLOBALS['Control_Category_categoryFn'])['identity']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_List_Trans_monadListT
$GLOBALS['Control_Monad_List_Trans_monadListT'] = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $__res = ["Applicative0" => function($_dollar__unused_1 = null) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_List_Trans_applicativeListT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar__unused_1 = null) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_List_Trans_bindListT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_List_Trans_bindListT
$GLOBALS['Control_Monad_List_Trans_bindListT'] = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $append_1_0 = ($GLOBALS['Control_Monad_List_Trans_concat'])((($dictMonad_0)['Applicative0'])(null));
  $__local_var_2_1 = (((((($dictMonad_0)['Bind1'])(null))['Apply0'])(null))['Functor0'])(null);
  $__res = ["bind" => (function() use ($__local_var_2_1, $append_1_0, $dictMonad_0) {
  $__fn = function($fa_3 = null, $f_4 = null) use ($__local_var_2_1, $append_1_0, $dictMonad_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_2_1)['map'])(function($v_5 = null) use ($append_1_0, $dictMonad_0, $f_4) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ((is_object($v_5) && (($v_5)->{'tag'} === "Yield"))) {
$__local_var_6_3 = ($v_5)->{'value0'};
$__t2 = new Phpurs_Data1("Skip", ((($GLOBALS['Data_Lazy_functorLazy'])['map'])(function($s_prime_7 = null) use ($__local_var_6_3, $append_1_0, $dictMonad_0, $f_4) {
  $__num = \func_num_args();
  $__res = (($append_1_0)(($f_4)($__local_var_6_3)))((((($GLOBALS['Control_Monad_List_Trans_bindListT'])($dictMonad_0))['bind'])($s_prime_7))($f_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_5)->{'value1'}));
goto end_branch_2;;
};
  if ((is_object($v_5) && (($v_5)->{'tag'} === "Skip"))) {
$__t2 = new Phpurs_Data1("Skip", ((($GLOBALS['Data_Lazy_functorLazy'])['map'])(function($v1_6 = null) use ($dictMonad_0, $f_4) {
  $__num = \func_num_args();
  $__res = (((($GLOBALS['Control_Monad_List_Trans_bindListT'])($dictMonad_0))['bind'])($v1_6))($f_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_5)->{'value0'}));
goto end_branch_2;;
};
  if ((is_object($v_5) && (($v_5)->{'tag'} === "Done"))) {
$__t2 = new Phpurs_Data0("Done");
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
})(), "Apply0" => function($_dollar__unused_3 = null) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_List_Trans_applyListT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_List_Trans_applyListT
$GLOBALS['Control_Monad_List_Trans_applyListT'] = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $functorListT1_1_0 = ($GLOBALS['Control_Monad_List_Trans_functorListT'])((((((($dictMonad_0)['Bind1'])(null))['Apply0'])(null))['Functor0'])(null));
  $__local_var_2_1 = ($GLOBALS['Control_Monad_List_Trans_bindListT'])($dictMonad_0);
  $__res = ["apply" => (function() use ($__local_var_2_1, $dictMonad_0) {
  $__fn = function($f_3 = null, $a_4 = null) use ($__local_var_2_1, $dictMonad_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_2_1)['bind'])($f_3))(function($f_prime_5 = null) use ($__local_var_2_1, $a_4, $dictMonad_0) {
  $__num = \func_num_args();
  $__res = ((($__local_var_2_1)['bind'])($a_4))(function($a_prime_6 = null) use ($dictMonad_0, $f_prime_5) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Monad_List_Trans_applicativeListT'])($dictMonad_0))['pure'])(($f_prime_5)($a_prime_6));
  goto __end;;
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
})(), "Functor0" => function($_dollar__unused_2 = null) use ($functorListT1_1_0) {
  $__num = \func_num_args();
  $__res = $functorListT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_List_Trans_applicativeListT
$GLOBALS['Control_Monad_List_Trans_applicativeListT'] = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $__res = ["pure" => ($GLOBALS['Control_Monad_List_Trans_singleton'])((($dictMonad_0)['Applicative0'])(null)), "Apply0" => function($_dollar__unused_1 = null) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_List_Trans_applyListT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_List_Trans_monadEffectListT
$GLOBALS['Control_Monad_List_Trans_monadEffectListT'] = function($dictMonadEffect_0 = null) {
  $__num = \func_num_args();
  $Monad0_1_0 = (($dictMonadEffect_0)['Monad0'])(null);
  $monadListT1_2_1 = ["Applicative0" => function($_dollar__unused_2 = null) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ["pure" => ($GLOBALS['Control_Monad_List_Trans_singleton'])((($Monad0_1_0)['Applicative0'])(null)), "Apply0" => function($_dollar__unused_3 = null) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_List_Trans_applyListT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar__unused_2 = null) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_List_Trans_bindListT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ["liftEffect" => ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])((($GLOBALS['Control_Monad_List_Trans_monadTransListT'])['lift'])($Monad0_1_0)))(($dictMonadEffect_0)['liftEffect']), "Monad0" => function($_dollar__unused_3 = null) use ($monadListT1_2_1) {
  $__num = \func_num_args();
  $__res = $monadListT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_List_Trans_monadSTListT
$GLOBALS['Control_Monad_List_Trans_monadSTListT'] = function($dictMonadST_0 = null) {
  $__num = \func_num_args();
  $Monad0_1_0 = (($dictMonadST_0)['Monad0'])(null);
  $monadListT1_2_1 = ["Applicative0" => function($_dollar__unused_2 = null) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ["pure" => ($GLOBALS['Control_Monad_List_Trans_singleton'])((($Monad0_1_0)['Applicative0'])(null)), "Apply0" => function($_dollar__unused_3 = null) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_List_Trans_applyListT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar__unused_2 = null) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_List_Trans_bindListT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ["liftST" => ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])((($GLOBALS['Control_Monad_List_Trans_monadTransListT'])['lift'])($Monad0_1_0)))(($dictMonadST_0)['liftST']), "Monad0" => function($_dollar__unused_3 = null) use ($monadListT1_2_1) {
  $__num = \func_num_args();
  $__res = $monadListT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_List_Trans_altListT
$GLOBALS['Control_Monad_List_Trans_altListT'] = function($dictApplicative_0 = null) {
  $__num = \func_num_args();
  $functorListT1_1_0 = ($GLOBALS['Control_Monad_List_Trans_functorListT'])((((($dictApplicative_0)['Apply0'])(null))['Functor0'])(null));
  $__res = ["alt" => ($GLOBALS['Control_Monad_List_Trans_concat'])($dictApplicative_0), "Functor0" => function($_dollar__unused_2 = null) use ($functorListT1_1_0) {
  $__num = \func_num_args();
  $__res = $functorListT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_List_Trans_plusListT
$GLOBALS['Control_Monad_List_Trans_plusListT'] = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $Applicative0_1_0 = (($dictMonad_0)['Applicative0'])(null);
  $altListT1_2_1 = ($GLOBALS['Control_Monad_List_Trans_altListT'])($Applicative0_1_0);
  $__res = ["empty" => (($Applicative0_1_0)['pure'])(new Phpurs_Data0("Done")), "Alt0" => function($_dollar__unused_3 = null) use ($altListT1_2_1) {
  $__num = \func_num_args();
  $__res = $altListT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_List_Trans_alternativeListT
$GLOBALS['Control_Monad_List_Trans_alternativeListT'] = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $applicativeListT1_1_0 = ["pure" => ($GLOBALS['Control_Monad_List_Trans_singleton'])((($dictMonad_0)['Applicative0'])(null)), "Apply0" => function($_dollar__unused_1 = null) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_List_Trans_applyListT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $plusListT1_2_1 = ($GLOBALS['Control_Monad_List_Trans_plusListT'])($dictMonad_0);
  $__res = ["Applicative0" => function($_dollar__unused_3 = null) use ($applicativeListT1_1_0) {
  $__num = \func_num_args();
  $__res = $applicativeListT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Plus1" => function($_dollar__unused_3 = null) use ($plusListT1_2_1) {
  $__num = \func_num_args();
  $__res = $plusListT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_List_Trans_monadPlusListT
$GLOBALS['Control_Monad_List_Trans_monadPlusListT'] = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $monadListT1_1_0 = ["Applicative0" => function($_dollar__unused_1 = null) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ["pure" => ($GLOBALS['Control_Monad_List_Trans_singleton'])((($dictMonad_0)['Applicative0'])(null)), "Apply0" => function($_dollar__unused_2 = null) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_List_Trans_applyListT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar__unused_1 = null) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_List_Trans_bindListT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $alternativeListT1_2_1 = ($GLOBALS['Control_Monad_List_Trans_alternativeListT'])($dictMonad_0);
  $__res = ["Monad0" => function($_dollar__unused_3 = null) use ($monadListT1_1_0) {
  $__num = \func_num_args();
  $__res = $monadListT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Alternative1" => function($_dollar__unused_3 = null) use ($alternativeListT1_2_1) {
  $__num = \func_num_args();
  $__res = $alternativeListT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

